package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/coreos/go-iptables/iptables"
)

// loggingResponseWriter captures statusCode of response
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK} // StatusOK is the default
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

type application struct {
	basicAuthUser string
	basicAuthPW   string
	errorLog      *log.Logger
	infoLog       *log.Logger
	debLog        *log.Logger
	logFile       *os.File
	debugEnabled  bool
	confFile      string
	srv           http.Server
	execPath      string
	certFile      string
	keyFile       string
	ipt           *iptables.IPTables
	iptChains     []string
	iptTbl        string
}

// used to read configuration file in json format
type confFileContent struct {
	BasicAuthUser string   `json:"basic_auth_user"`
	BasicAuthPW   string   `json:"basic_auth_pw"`
	DebugEnabled  bool     `json:"debug_enabled"`
	IptChains     []string `json:"ipt_chains"`
	// IptTbl        string   `json:"ipt_tbl"`
	// SrvAddr           string   `json:"srv_addr"`
	CertFile          string   `json:"cert_file"`
	KeyFile           string   `json:"key_file"`
	SrvMaxHeaderBytes int      `json:"srv_max_header_bytes"`
	SrvIdleTimeout    Duration `json:"srv_idle_timeout"`
	SrvReadTimeout    Duration `json:"srv_read_timeout"`
	SrvWriteTimeout   Duration `json:"srv_write_timeout"`
	TLSMinVersion     string   `json:"tls_min_version"` //uint16 // tls.VersionTLS12
}

// Duration wraps time.Duration for the use in json files
type Duration struct {
	time.Duration
}

// MarshalJSON marshals the Duration type
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON unmarshals the Duration type
func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}
