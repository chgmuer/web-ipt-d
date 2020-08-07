package main

import (
	"testing"
)

// func Test_initConfig(t *testing.T) {
// 	type args struct {
// 		confFile string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *application
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := initConfig(tt.args.confFile); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("initConfig() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_application_initConfFromFile(t *testing.T) {
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
// 	tests := []struct {
// 		name   string
// 		fields fields
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
// 			app.initConfFromFile()
// 		})
// 	}
// }

// func Test_application_createLogger(t *testing.T) {
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
// 	tests := []struct {
// 		name   string
// 		fields fields
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
// 			app.createLogger()
// 		})
// 	}
// }

// func Test_application_setSrvDefaults(t *testing.T) {
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
// 	tests := []struct {
// 		name   string
// 		fields fields
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
// 			app.setSrvDefaults()
// 		})
// 	}
// }

// func Test_writeConfigFile(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			writeConfigFile()
// 		})
// 	}
// }

// func Test_application_changeBasicAuth(t *testing.T) {
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
// 		user string
// 		pw   string
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
// 			app.changeBasicAuth(tt.args.user, tt.args.pw)
// 		})
// 	}
// }

func Test_checkDiffChains(t *testing.T) {
	type args struct {
		old []string
		new []string
	}
	tests := []struct {
		name     string
		args     args
		wantDel  []string
		wantCrea []string
	}{
		{"old==new 1", args{[]string{"a", "b", "c"}, []string{"a", "b", "c"}}, []string{}, []string{}},
		{"old==new 2", args{[]string{"c", "b", "a"}, []string{"b", "c", "a"}}, []string{}, []string{}},
		{"no old in new", args{[]string{"a", "b", "c"}, []string{"d", "e", "f"}}, []string{"a", "b", "c"}, []string{"d", "e", "f"}},
		{"empty old new ", args{[]string{}, []string{"a", "b", "c"}}, []string{}, []string{"a", "b", "c"}},
		{"old empty new ", args{[]string{"a", "b", "c"}, []string{}}, []string{"a", "b", "c"}, []string{}},
		{"two old - new 1", args{[]string{"a", "b", "c"}, []string{"a", "b", "f"}}, []string{"c"}, []string{"f"}},
		{"two old - new 2", args{[]string{"a", "b", "c"}, []string{"a", "f", "c"}}, []string{"b"}, []string{"f"}},
		{"two old - new 3", args{[]string{"a", "b", "c"}, []string{"f", "b", "c"}}, []string{"a"}, []string{"f"}},
		{"one old - new 1", args{[]string{"a", "b", "c"}, []string{"a", "f", "g"}}, []string{"b", "c"}, []string{"f", "g"}},
		{"one old - new 2", args{[]string{"a", "b", "c"}, []string{"f", "a", "g"}}, []string{"b", "c"}, []string{"f", "g"}},
		{"one old - new 3", args{[]string{"a", "b", "c"}, []string{"f", "g", "a"}}, []string{"b", "c"}, []string{"f", "g"}},
		{"one empty old - new", args{[]string{"a", "b", ""}, []string{"a", "b", "c"}}, []string{}, []string{"c"}},
		{"old - one empty new", args{[]string{"a", "b", "c"}, []string{"a", "", "c"}}, []string{"b"}, []string{}},
		{"nixdrin", args{[]string{"", "b", ""}, []string{"", "", "b"}}, []string{}, []string{}},
		{"nixdrin-really", args{[]string{"", "", ""}, []string{"", "", ""}}, []string{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDel, gotCrea := checkDiffChains(tt.args.old, tt.args.new)
			if len(gotDel) != len(tt.wantDel) {
				t.Errorf("checkDiffChains(), not the same length got=%v, want=%v", gotDel, tt.wantDel)
			}
			if len(gotCrea) != len(tt.wantCrea) {
				t.Errorf("checkDiffChains(), not the same length got=%v, want=%v", gotDel, tt.wantDel)
			}
			// test exists()
			for _, d := range gotDel {
				if !exists(tt.wantDel, d) {
					t.Errorf("exists() in checkDiffChains(), not found in got=%v, expected to find want=%v", gotDel, d)
				}
			}
			for _, c := range gotCrea {
				if !exists(tt.wantCrea, c) {
					t.Errorf("exists() in checkDiffChains(), not found in got=%v, expected to find want=%v", gotCrea, c)
				}
			}
		})
	}
}
