package main

import (
	"fmt"
	"net/http"
)

// getAllChains
func (app *application) getAllChains(w http.ResponseWriter, r *http.Request) {
	listChains, err := app.listChains()
	if err != nil {
		app.errorLog.Printf("Failed during listChains err=%v\n", err)
		app.serverError(w, err)
		return
	}
	fmt.Fprintf(w, "List of chains in table=%s:", app.iptTbl)
	count := 0
	for _, chain := range listChains {
		fmt.Fprintf(w, "%s\t", chain)
		count++
	}
	fmt.Fprintf(w, "\nNumber of chains in table=%s: count=%d\n", app.iptTbl, count)
}

// getIPsInChain - get all rules in chain
func (app *application) getIPsInChain(w http.ResponseWriter, r *http.Request) {
	chain, err := app.getChainFromURL(r)
	if err != nil {
		app.errorLog.Printf("URL parse error, chain=%s err=%v\n", chain, err)
		app.clientError(w, http.StatusBadRequest)
		return
	}
	list, err := app.listIPsInChain(chain)
	if err != nil {
		app.errorLog.Printf("Failed during listIPsInChain err=%v\n", err)
		app.serverError(w, err)
		return
	}
	fmt.Fprintf(w, "list of entries in chain=%s:\n", chain)
	count := 0
	for _, entry := range list {
		fmt.Fprintf(w, "%v\n", entry) // entry is []string
		count++
	}
	fmt.Fprintf(w, "Number of entries in chain=%s: count=%d\n", chain, count)
}

// getIPv4 - check one rule for one given ip
func (app *application) getIPv4(w http.ResponseWriter, r *http.Request) {
	chain, ipv4, err := app.getChainAndIPFromURL(r)
	if err != nil {
		app.errorLog.Printf("URL parse error, chain=%s ipv4=%s err=%v\n", chain, ipv4, err)
		app.clientError(w, http.StatusBadRequest)
		return
	}
	exists, err := app.getRule(chain, ipv4)
	if err != nil {
		app.errorLog.Printf("Failed during getRule err=%v\n", err)
		app.serverError(w, err)
		return
	}
	output := fmt.Sprintf("Check if IP=%s exists in chain=%s returns %v\n", ipv4, chain, exists)
	w.Write([]byte(output))
}

func (app *application) putIPv4(w http.ResponseWriter, r *http.Request) {
	chain, ipv4, err := app.getChainAndIPFromURL(r)
	if err != nil {
		app.errorLog.Printf("URL parse error, chain=%s ipv4=%s err=%v\n", chain, ipv4, err)
		app.clientError(w, http.StatusBadRequest)
		return
	}
	err = app.putRule(chain, ipv4)
	if err != nil {
		app.errorLog.Printf("Failed during putRule with ip=%s to chain=%s err=%v\n", ipv4, chain, err)
		app.serverError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (app *application) delIPv4(w http.ResponseWriter, r *http.Request) {
	chain, ipv4, err := app.getChainAndIPFromURL(r)
	if err != nil {
		app.errorLog.Printf("URL parse error, chain=%s ipv4=%s err=%v\n", chain, ipv4, err)
		app.clientError(w, http.StatusBadRequest)
		return
	}
	exists, err := app.getRule(chain, ipv4)
	if err != nil {
		app.errorLog.Printf("Failed during getRule in delRule err=%v\n", err)
		app.serverError(w, err)
		return
	}
	if exists {
		err = app.delRule(chain, ipv4)
		if err != nil {
			app.errorLog.Printf("Failed during delRule with ip=%s in chain=%s err=%v\n", ipv4, chain, err)
			app.serverError(w, err)
			return
		}
		app.debLog.Printf("Returning from delIPv4 with ip=%s in chain=%s rule deleted\n", ipv4, chain)
	} else {
		app.infoLog.Printf("Could not delete rule with ip=%s in chain=%s as rule does not exist\n", ipv4, chain)
		app.clientError(w, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (app *application) metrics(w http.ResponseWriter, r *http.Request) {
	// parseUrl
	// getResult put, del whatever
	// printResult
	list, err := app.getMetrics()
	if err != nil {
		app.errorLog.Printf("Failed during getMetrics of default chain=%s err=%v\n", app.iptChains, err)
		app.serverError(w, err)
		return
	}
	fmt.Fprintln(w, "list of entries in default chain:")
	count := 0
	for _, entry := range list {
		fmt.Fprintf(w, "%v\n", entry) // entry is []string
		count++
	}
	fmt.Fprintf(w, "Number of entries: count=%d\n", count)
}
