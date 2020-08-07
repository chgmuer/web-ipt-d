package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/coreos/go-systemd/daemon"
)

func (app *application) sendKeepAlive() {
	url := "https://" + app.srv.Addr + "/"
	httpTransport := http.Transport{
		TLSClientConfig:    app.tlsConfig(),
		DisableCompression: true,
	}
	client := &http.Client{Transport: &httpTransport}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		app.errorLog.Printf("Could not create request err=%v", err)
	}
	req.SetBasicAuth(app.basicAuthUser, app.basicAuthPW)

	_, getErr := client.Do(req)
	if getErr != nil {
		app.errorLog.Printf("--- WARNING --- WatchDog GET returned err=%v", getErr)
	}

	interval, err := daemon.SdWatchdogEnabled(false)
	if err != nil || interval == 0 {
		app.errorLog.Printf("--- WARNING --- systemd watchdog not working, interval=%v, err=%v\n", interval, err)
		return
	}

	tick := time.NewTicker(interval / 3)
	for {
		select {
		case <-tick.C:
			{
				_, getErr := client.Do(req)
				if getErr == nil {
					notified, notifyErr := daemon.SdNotify(false, daemon.SdNotifyWatchdog)
					app.debLog.Printf("SdNotify expected (true,nil) got (%v,%v)", notified, notifyErr) // true,nil - notification supported, data has been sent
				}
			}
		}
	}
	// for {
	// 	_, getErr := client.Do(req)
	// 	if getErr == nil {
	// 		daemon.SdNotify(false, daemon.SdNotifyWatchdog)
	// 	}
	// 	time.Sleep(interval / 3)
	// }
}

func (app *application) tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile(app.certFile)
	if err != nil {
		app.errorLog.Fatalf("Could not read certificate err=%v", err)
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(crt)

	return &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: false,
		ServerName:         "localhost",
	}
}
