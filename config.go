package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/coreos/go-iptables/iptables"
)

func initConfig(confFile string) *application {
	app := new(application)

	app.createLogger()
	// may be overwritten later with initConfFromFile()
	app.setSrvDefaults()

	execDir, err := os.Executable()
	if err != nil {
		app.errorLog.Fatalln("Failed to get absolute executable name")
	}
	app.execPath = path.Dir(execDir)

	// iptables initialisation
	ipt, err := iptables.New()
	if err != nil {
		app.errorLog.Fatalln("Could not create struct for iptables")
	}
	app.ipt = ipt
	app.iptTbl = "filter"
	app.confFile = confFile
	app.initConfFromFile()
	return app
}

// initConfFromFile reads conf file and updates configuration
// TODO lock during write
func (app *application) initConfFromFile() {

	var fp confFileContent

	app.infoLog.Printf("Reading config parameters from %s", app.confFile)
	fdata, err := ioutil.ReadFile(app.confFile)
	if err != nil {
		app.errorLog.Printf("Reading file %s failed err=%v\n", app.confFile, err)
		return
	}
	err = json.Unmarshal(fdata, &fp)
	if err != nil {
		app.errorLog.Printf("Unmarshal err=%v\n", err)
		return
	}

	if (app.basicAuthUser != fp.BasicAuthUser) || (app.basicAuthPW != fp.BasicAuthPW) {
		app.changeBasicAuth(fp.BasicAuthUser, fp.BasicAuthPW)
	}

	if app.debugEnabled != fp.DebugEnabled {
		app.debugEnabled = fp.DebugEnabled
		app.infoLog.Printf("Changing debug to %v", app.debugEnabled)
		app.debLog.SetOutput(ioutil.Discard)
		if app.debugEnabled {
			app.debLog.SetOutput(app.logFile)
		}
	}

	// if (app.iptTbl != fp.IptTbl) && (fp.IptTbl != "filter") {
	// 	app.errorLog.Printf("WARNING Change of iptables table value is not supported and this will break things - old table=%v, new table=%v\n", app.iptTbl, fp.IptTbl)
	// }
	// app.iptTbl = fp.IptTbl

	del, crea := checkDiffChains(app.iptChains, fp.IptChains)
	if len(crea) > 0 {
		for _, chain := range crea {
			err = app.ipt.NewChain(app.iptTbl, chain)
			if err != nil {
				// app.infoLog.Printf("Could not create chain=%s in table=%s; trying ClearChain", chain, app.iptTbl)
				err = app.ipt.ClearChain(app.iptTbl, chain)
				if err != nil {
					app.errorLog.Printf("ClearChain of chain=%s in table=%s failed: %v\n", chain, app.iptTbl, err)
					app.errorLog.Println("Could not create and clear chain in table")
				}
			}
			app.infoLog.Printf("Created (or cleared) chain=%s in table=%s", chain, app.iptTbl)
		}
	}
	if len(del) > 0 {
		for _, chain := range del {
			err = app.ipt.ClearChain(app.iptTbl, chain)
			if err != nil {
				app.errorLog.Printf("ClearChain of chain=%s in table=%s failed: %v\n", chain, app.iptTbl, err)
				app.errorLog.Println("Cannot delete chain in table")
			}
			err = app.ipt.DeleteChain(app.iptTbl, chain)
			if err != nil {
				app.errorLog.Printf("DeleteChain of chain=%s in table=%s failed: %v\n", chain, app.iptTbl, err)
			}
			app.infoLog.Printf("Deleted chain=%s in table=%s", chain, app.iptTbl)
		}
	}
	app.iptChains = fp.IptChains

	app.certFile = fp.CertFile
	app.keyFile = fp.KeyFile
	// app.srv.Addr = fp.SrvAddr
	app.srv.MaxHeaderBytes = fp.SrvMaxHeaderBytes
	app.srv.IdleTimeout = fp.SrvIdleTimeout.Duration
	app.srv.ReadHeaderTimeout = fp.SrvReadTimeout.Duration
	app.srv.WriteTimeout = fp.SrvWriteTimeout.Duration
	app.srv.TLSConfig.MinVersion = tls.VersionTLS12 // defaults to 1.2
	if fp.TLSMinVersion == "1.3" {
		app.srv.TLSConfig.MinVersion = tls.VersionTLS13
	}
}

func (app *application) createLogger() {
	var mw io.Writer
	fn := "/var/log/" + filepath.Base(os.Args[0]) + ".log"
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Could not create log File - setting logFile to os.Stdout")
		f = os.Stdout
		mw = io.Writer(f)
	} else {
		mw = io.MultiWriter(os.Stdout, f)
	}
	app.logFile = f // used for switching debug on and off
	app.errorLog = log.New(mw, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if err != nil {
		app.errorLog.Printf("Could not create/append log file err=%v", err)
	}

	app.infoLog = log.New(app.logFile, "INFO\t", log.Ldate|log.Ltime)
	app.debLog = log.New(app.logFile, "DEBUG\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func (app *application) setSrvDefaults() {
	var TLSConfig = tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		MinVersion: tls.VersionTLS12,
	}

	app.srv = http.Server{
		// Addr:           "0.0.0.0:8820", // This is read from the socket unit file
		ErrorLog:       app.errorLog,
		Handler:        app.routes(),     // http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(`Hello, world!`)) }),
		TLSConfig:      &TLSConfig,       // ignord for TLS1.3
		MaxHeaderBytes: 4096,             // defaults to 1MByte
		IdleTimeout:    time.Minute,      // close connection after 1 - default 3
		ReadTimeout:    5 * time.Second,  // connection close without any client msg
		WriteTimeout:   10 * time.Second, // should be bigger than read timeout for https
	}
}

func writeConfigFile() {
	var fp confFileContent
	execDir, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get absolute executable name")
	}
	execPath := path.Dir(execDir)
	fp.CertFile = execPath + "/tls/cert.pem"
	fp.KeyFile = execPath + "/tls/key.pem"
	fp.IptChains = []string{"httpd", "sshd"}
	// fp.IptTbl = "filter"

	fp.BasicAuthUser = ""
	fp.BasicAuthPW = ""
	fp.DebugEnabled = true

	fp.SrvMaxHeaderBytes = 4096
	fp.SrvIdleTimeout = Duration{time.Minute}
	fp.SrvReadTimeout = Duration{5 * time.Second}
	fp.SrvWriteTimeout = Duration{10 * time.Second}
	fp.TLSMinVersion = "1.2"

	jdata, err := json.MarshalIndent(fp, "", "    ")
	if err != nil {
		fmt.Printf("Error - MarshalIndent err=%v\n", err)
		return
	}
	cf := execPath + "/config.json"
	err = ioutil.WriteFile(cf, jdata, 0600) // -rw------- 1 root    root        332 Jun 19 16:22 config.json
	if err != nil {
		fmt.Printf("Error - Could not write to file %s err=%v\n", cf, err)
		return
	}
	fmt.Printf("Config file written to %s - check owner, permission and edit content\n", cf)
	return
}

func (app *application) changeBasicAuth(user string, pw string) {
	user = strings.TrimSpace(user)
	pw = strings.TrimSpace(pw)
	if len(user) < 4 || len(pw) < 4 {
		app.errorLog.Fatalln("Basic Auth user and/or pw are too short")
		return
	}
	app.basicAuthUser = user
	app.basicAuthPW = pw
}

func checkDiffChains(old []string, new []string) ([]string, []string) {
	var del, crea []string
	for i, v := range old {
		old[i] = strings.TrimSpace(v)
	}
	for i, v := range new {
		new[i] = strings.TrimSpace(v)
	}
	// create if in new but not in old
	for i := len(new) - 1; i >= 0; i-- {
		if !exists(old, new[i]) {
			if new[i] != "" {
				crea = append(crea, new[i])
			}
		}
	}
	// delete if in old but not in new
	for i := len(old) - 1; i >= 0; i-- {
		if !exists(new, old[i]) {
			if old[i] != "" {
				del = append(del, old[i])
			}
		}
	}
	return del, crea
}

func exists(slice []string, str string) bool {
	for _, s := range slice {
		if str == s {
			return true
		}
	}
	return false
}
