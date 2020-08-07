package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coreos/go-systemd/activation"
	"github.com/coreos/go-systemd/daemon"
)

// func (app *application) runListenAndServe() {
// 	app.infoLog.Printf("Starting server on %s", app.srv.Addr)
// 	certFile := app.certFile // private
// 	keyFile := app.keyFile   // public
// 	err := app.srv.ListenAndServeTLS(certFile, keyFile)
// 	app.errorLog.Fatal(err)
// }

// func (app *application) runListenAndServeDaemon() { // needs .service file for systemd
// 	app.infoLog.Printf("Starting server on %s", app.srv.Addr)
// 	certFile := app.certFile // private
// 	keyFile := app.keyFile   // public

// 	l, err := net.Listen("tcp", app.srv.Addr)
// 	if err != nil {
// 		app.errorLog.Fatalf("Could not listen on addr=%s, err=%v", app.srv.Addr, err)
// 	}
// 	sdNotify, err := daemon.SdNotify(false, daemon.SdNotifyReady)
// 	app.debLog.Printf("SdNotify returned %v", sdNotify)
// 	if err != nil {
// 		app.errorLog.Printf("SdNotify returned err=%v", err)
// 	}

// 	go app.sendKeepAlive()

// 	http.ServeTLS(l, app.srv.Handler, certFile, keyFile)
// 	if err != nil {
// 		app.errorLog.Fatalf("Could not ServeTLS on addr=%s, err=%v", app.srv.Addr, err)
// 	}
// }

func (app *application) runListenAndServeDaemon() {

	certFile := app.certFile // private
	keyFile := app.keyFile   // public

	listeners, err := activation.Listeners()
	if err != nil {
		app.errorLog.Fatalf("Could not retrieve listeners err=%v", err)
	}
	if len(listeners) != 1 {
		app.errorLog.Fatalf("Can only handle 1 listener socket but got=%d", len(listeners))
	}
	app.srv.Addr = listeners[0].Addr().String() // used in keepalive.go and error messages
	app.infoLog.Printf("Starting server on listener[0] addr=%s", app.srv.Addr)

	sdNotify, err := daemon.SdNotify(false, daemon.SdNotifyReady)
	app.debLog.Printf("SdNotify returned %v, err=%v", sdNotify, err)
	if err != nil {
		app.errorLog.Printf("SdNotify returned %v, err=%v", sdNotify, err)
	}

	// prepare to receive signals
	// done := make(chan struct{})
	doShutdown := make(chan os.Signal, 1)

	signal.Notify(doShutdown,
		syscall.SIGHUP,
		syscall.SIGINT) // , syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for {
			sig := <-doShutdown
			app.infoLog.Printf("Server is shutting down due to signal=%v", sig)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // interval / 6 ?
			defer cancel()
			app.srv.SetKeepAlivesEnabled(false)
			if err := app.srv.Shutdown(ctx); err != nil {
				app.errorLog.Panicf("Cannot gracefully shut down, err=%v", err)
			}
			app.initConfFromFile()
			// close(done)
		}
	}()

	go app.sendKeepAlive()

	http.ServeTLS(listeners[0], app.srv.Handler, certFile, keyFile)
	if err != nil {
		app.errorLog.Fatalf("Could not ServeTLS on addr=%s, err=%v", app.srv.Addr, err)
	}
	// <-done
}
