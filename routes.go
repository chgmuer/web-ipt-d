package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, app.secureHeaders, app.basicAuth)
	mux := pat.New()

	mux.Get("/chain/:chain/ipv4/:ipv4", http.HandlerFunc(app.getIPv4))
	mux.Put("/chain/:chain/ipv4/:ipv4", http.HandlerFunc(app.putIPv4))
	mux.Del("/chain/:chain/ipv4/:ipv4", http.HandlerFunc(app.delIPv4))

	mux.Get("/chain/:chain", http.HandlerFunc(app.getIPsInChain))

	mux.Get("/metrics", http.HandlerFunc(app.metrics)) // oder metrics in ein File schreiben und via Fileserver

	mux.Get("/", http.HandlerFunc(app.getAllChains)) // pat needs HandlerFunc to add ServeHTTP

	app.infoLog.Println("Mux initialized")
	return standardMiddleware.Then(mux)
}
