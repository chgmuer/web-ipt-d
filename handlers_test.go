package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_getAllChains(t *testing.T) {
	app := initConfig("./testdata/config.json")
	app.debLog = log.New(ioutil.Discard, "", 0)
	app.infoLog = log.New(ioutil.Discard, "", 0)
	app.debLog = log.New(ioutil.Discard, "", 0)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getAllChains)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := ""
	count := 0
	for _, chain := range app.iptChains {
		expected += fmt.Sprintf(expected, "%s\t", chain)
		count++
	}
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// func Test_application_getIPsInChain(t *testing.T) {
// 	type fields struct {
// 		basicAuthUser string
// 		basicAuthPW   string
// 		errorLog      *log.Logger
// 		infoLog       *log.Logger
// 		debLog        *log.Logger
// 		logFile       *os.File
// 		debugEnabled  bool
// 		confFile      string
// 		srv           http.Server
// 		execPath      string
// 		certFile      string
// 		keyFile       string
// 		ipt           *iptables.IPTables
// 		iptChains     []string
// 		iptTbl        string
// 	}
// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			app := &application{
// 				basicAuthUser: tt.fields.basicAuthUser,
// 				basicAuthPW:   tt.fields.basicAuthPW,
// 				errorLog:      tt.fields.errorLog,
// 				infoLog:       tt.fields.infoLog,
// 				debLog:        tt.fields.debLog,
// 				logFile:       tt.fields.logFile,
// 				debugEnabled:  tt.fields.debugEnabled,
// 				confFile:      tt.fields.confFile,
// 				srv:           tt.fields.srv,
// 				execPath:      tt.fields.execPath,
// 				certFile:      tt.fields.certFile,
// 				keyFile:       tt.fields.keyFile,
// 				ipt:           tt.fields.ipt,
// 				iptChains:     tt.fields.iptChains,
// 				iptTbl:        tt.fields.iptTbl,
// 			}
// 			app.getIPsInChain(tt.args.w, tt.args.r)
// 		})
// 	}
// }

// func Test_application_getIPv4(t *testing.T) {
// 	type fields struct {
// 		basicAuthUser string
// 		basicAuthPW   string
// 		errorLog      *log.Logger
// 		infoLog       *log.Logger
// 		debLog        *log.Logger
// 		logFile       *os.File
// 		debugEnabled  bool
// 		confFile      string
// 		srv           http.Server
// 		execPath      string
// 		certFile      string
// 		keyFile       string
// 		ipt           *iptables.IPTables
// 		iptChains     []string
// 		iptTbl        string
// 	}
// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			app := &application{
// 				basicAuthUser: tt.fields.basicAuthUser,
// 				basicAuthPW:   tt.fields.basicAuthPW,
// 				errorLog:      tt.fields.errorLog,
// 				infoLog:       tt.fields.infoLog,
// 				debLog:        tt.fields.debLog,
// 				logFile:       tt.fields.logFile,
// 				debugEnabled:  tt.fields.debugEnabled,
// 				confFile:      tt.fields.confFile,
// 				srv:           tt.fields.srv,
// 				execPath:      tt.fields.execPath,
// 				certFile:      tt.fields.certFile,
// 				keyFile:       tt.fields.keyFile,
// 				ipt:           tt.fields.ipt,
// 				iptChains:     tt.fields.iptChains,
// 				iptTbl:        tt.fields.iptTbl,
// 			}
// 			app.getIPv4(tt.args.w, tt.args.r)
// 		})
// 	}
// }

// func Test_application_putIPv4(t *testing.T) {
// 	type fields struct {
// 		basicAuthUser string
// 		basicAuthPW   string
// 		errorLog      *log.Logger
// 		infoLog       *log.Logger
// 		debLog        *log.Logger
// 		logFile       *os.File
// 		debugEnabled  bool
// 		confFile      string
// 		srv           http.Server
// 		execPath      string
// 		certFile      string
// 		keyFile       string
// 		ipt           *iptables.IPTables
// 		iptChains     []string
// 		iptTbl        string
// 	}
// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			app := &application{
// 				basicAuthUser: tt.fields.basicAuthUser,
// 				basicAuthPW:   tt.fields.basicAuthPW,
// 				errorLog:      tt.fields.errorLog,
// 				infoLog:       tt.fields.infoLog,
// 				debLog:        tt.fields.debLog,
// 				logFile:       tt.fields.logFile,
// 				debugEnabled:  tt.fields.debugEnabled,
// 				confFile:      tt.fields.confFile,
// 				srv:           tt.fields.srv,
// 				execPath:      tt.fields.execPath,
// 				certFile:      tt.fields.certFile,
// 				keyFile:       tt.fields.keyFile,
// 				ipt:           tt.fields.ipt,
// 				iptChains:     tt.fields.iptChains,
// 				iptTbl:        tt.fields.iptTbl,
// 			}
// 			app.putIPv4(tt.args.w, tt.args.r)
// 		})
// 	}
// }

// func Test_application_delIPv4(t *testing.T) {
// 	type fields struct {
// 		basicAuthUser string
// 		basicAuthPW   string
// 		errorLog      *log.Logger
// 		infoLog       *log.Logger
// 		debLog        *log.Logger
// 		logFile       *os.File
// 		debugEnabled  bool
// 		confFile      string
// 		srv           http.Server
// 		execPath      string
// 		certFile      string
// 		keyFile       string
// 		ipt           *iptables.IPTables
// 		iptChains     []string
// 		iptTbl        string
// 	}
// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			app := &application{
// 				basicAuthUser: tt.fields.basicAuthUser,
// 				basicAuthPW:   tt.fields.basicAuthPW,
// 				errorLog:      tt.fields.errorLog,
// 				infoLog:       tt.fields.infoLog,
// 				debLog:        tt.fields.debLog,
// 				logFile:       tt.fields.logFile,
// 				debugEnabled:  tt.fields.debugEnabled,
// 				confFile:      tt.fields.confFile,
// 				srv:           tt.fields.srv,
// 				execPath:      tt.fields.execPath,
// 				certFile:      tt.fields.certFile,
// 				keyFile:       tt.fields.keyFile,
// 				ipt:           tt.fields.ipt,
// 				iptChains:     tt.fields.iptChains,
// 				iptTbl:        tt.fields.iptTbl,
// 			}
// 			app.delIPv4(tt.args.w, tt.args.r)
// 		})
// 	}
// }

// func Test_application_metrics(t *testing.T) {
// 	type fields struct {
// 		basicAuthUser string
// 		basicAuthPW   string
// 		errorLog      *log.Logger
// 		infoLog       *log.Logger
// 		debLog        *log.Logger
// 		logFile       *os.File
// 		debugEnabled  bool
// 		confFile      string
// 		srv           http.Server
// 		execPath      string
// 		certFile      string
// 		keyFile       string
// 		ipt           *iptables.IPTables
// 		iptChains     []string
// 		iptTbl        string
// 	}
// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			app := &application{
// 				basicAuthUser: tt.fields.basicAuthUser,
// 				basicAuthPW:   tt.fields.basicAuthPW,
// 				errorLog:      tt.fields.errorLog,
// 				infoLog:       tt.fields.infoLog,
// 				debLog:        tt.fields.debLog,
// 				logFile:       tt.fields.logFile,
// 				debugEnabled:  tt.fields.debugEnabled,
// 				confFile:      tt.fields.confFile,
// 				srv:           tt.fields.srv,
// 				execPath:      tt.fields.execPath,
// 				certFile:      tt.fields.certFile,
// 				keyFile:       tt.fields.keyFile,
// 				ipt:           tt.fields.ipt,
// 				iptChains:     tt.fields.iptChains,
// 				iptTbl:        tt.fields.iptTbl,
// 			}
// 			app.metrics(tt.args.w, tt.args.r)
// 		})
// 	}
// }
