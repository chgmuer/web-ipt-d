package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// standard middleware pattern:
// func myMiddleware(next http.Handler) http.Handler {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		// TODO: Execute our middleware logic here...
// 		next.ServeHTTP(w, r)
// 	}
//
// 	return http.HandlerFunc(fn)
// }
// same with anonymous function
// func myMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // TODO: Execute our middleware logic here...
//         next.ServeHTTP(w, r)
//     })
// }

func (app *application) secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := newLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)
		statusCode := lrw.statusCode

		// must be after next as time.Since is calulated immediately even if deferred.
		// Defer necessary - if panic occured will unwind with stack
		defer app.infoLog.Printf("RemoteAddr=%s Proto=%s Method=%s Code=%d Duration=%v URL=%s", r.RemoteAddr, r.Proto, r.Method, statusCode, time.Since(start), r.URL.RequestURI())
	})

}

func (app *application) basicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the user isn't authorized send a 403 Forbidden status and
		// return to stop executing the chain.
		if !app.isAuthorized(r) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Authentication required"`)
			w.WriteHeader(401)
			w.Write([]byte("401 Unauthorized\n"))
			return
		}

		// Otherwise, call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}

func (app *application) isAuthorized(r *http.Request) bool {
	authorized := false
	u, p, ok := r.BasicAuth()
	if !ok || len(strings.TrimSpace(u)) < 1 || len(strings.TrimSpace(p)) < 1 {
		return authorized
	}
	// This is a dummy check for credentials.
	if u == app.basicAuthUser && p == app.basicAuthPW {
		authorized = true
	}
	return authorized
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// run as go unwinds the stack).
		defer func() {
			// recover from panic
			if err := recover(); err != nil {
				// Set a "Connection: close" header on the response.
				w.Header().Set("Connection", "close") // HTTP/2, Go will automatically strip the Connection: Close header from the response (so it is not malformed) and send a GOAWAY frame.
				// Call the app.serverError helper method to return a 500
				// Internal Server response.
				app.serverError(w, fmt.Errorf("%v", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
