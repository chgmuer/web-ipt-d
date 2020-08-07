package main

import (
	"fmt"
	"net"
	"net/http"
)

// create standardised rule based on ipv4 string
func createRuleString(ipv4 string) []string {
	r := []string{"-s", ipv4, "-j", "DROP"}
	// r = append(r, "-s")
	// r = append(r, ipv4)
	// r = append(r, "-j")
	// r = append(r, "DROP")
	// r = append(r, "--reject-with icmp-host-unreachable") // ? gibt es nicht mehr
	return r
}

// path handling functions
// getChainFromURL returns chain from path and checks with default chain
func (app *application) getChainFromURL(r *http.Request) (string, error) {
	chain := r.URL.Query().Get(":chain")
	if !contains(app.iptChains, chain) {
		err := fmt.Errorf("Unsupported chain=%s", chain)
		app.errorLog.Printf("getChains chain=%s failed, only configured chain=%s is supported\n", chain, app.iptChains)
		return chain, err
	}
	app.debLog.Printf("Extracted and formatted from URL=%s chain=%s", r.URL.EscapedPath(), chain)
	return chain, nil
}

// contains checks if string is in []string helper func
func contains(ss []string, s string) bool {
	for _, c := range ss {
		if c == s {
			return true
		}
	}
	return false
}

// getChainAndIPFromURL returns chain and ip from path, ip is validated and formatted
func (app *application) getChainAndIPFromURL(r *http.Request) (string, string, error) {
	chain, err := app.getChainFromURL(r)
	if err != nil {
		return chain, "", err
	}
	ipv4 := r.URL.Query().Get(":ipv4")
	n := net.ParseIP(ipv4)
	if n == nil {
		err = fmt.Errorf("The addr=%s is not recognised as a valid ipv4 address", ipv4)
		// app.errorLog.Printf("The addr=%s could not be parsed\n", ipv4)
		return chain, ipv4, err
	}
	ipv4 = fmt.Sprintf("%v", n) // to get formatted ip eg 003 -> 3
	app.debLog.Printf("Extracted and formatted from URL=%s chain=%s ipv4=%s", r.URL.EscapedPath(), chain, ipv4)
	return chain, ipv4, nil
}

// chain handling functions
// listChains
func (app *application) listChains() ([]string, error) {
	var listChain = []string{}
	listChain, err := app.ipt.ListChains(app.iptTbl)
	app.debLog.Printf("ListChain:\n%v\n", listChain)
	return listChain, err
}

// listIPsInChain
func (app *application) listIPsInChain(chain string) ([][]string, error) {
	var list = [][]string{}
	list, err := app.ipt.Stats(app.iptTbl, chain)
	app.debLog.Printf("Returning from listIPsInChains (Stats) of chain=%s:\n%v\n", chain, list)
	return list, err
}

// rule handling functions
// getRule checks if rule exists
func (app *application) getRule(chain string, ipv4 string) (bool, error) {
	// var exists = false
	ruleString := createRuleString(ipv4)
	exists, err := app.ipt.Exists(app.iptTbl, chain, ruleString...)
	app.debLog.Printf("Returning from check if rule with ipv4=%s in chain=%s exists=%v", ipv4, chain, exists)
	return exists, err
}

// putRule appends rule if rule is unique
func (app *application) putRule(chain string, ipv4 string) error {
	ruleString := createRuleString(ipv4)
	err := app.ipt.AppendUnique(app.iptTbl, chain, ruleString...)

	app.debLog.Printf("Returning from putRule with chain=%s, ipv4=%s", chain, ipv4)
	return err
}

// deleteRule deletes a rule from the chain
func (app *application) delRule(chain string, ipv4 string) error {
	ruleString := createRuleString(ipv4)
	err := app.ipt.Delete(app.iptTbl, chain, ruleString...)
	app.debLog.Printf("Returning from delRule with chain=%s, ipv4=%s", chain, ipv4)
	return err
}

// Stats
func (app *application) getMetrics() ([][]string, error) {
	var list [][]string
	list, err := app.ipt.Stats(app.iptTbl, app.iptChains[0]) // TODO
	app.debLog.Println("Returning from getMetrics")
	return list, err
}
