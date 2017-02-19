package main

import (
	"flag"
	"os"
)

const (
	loginAPI    = "/api/login"
	arptableAPI = "/api/settings/arptable"
	logoutAPI   = "/api/logout"
	rebootAPI   = "/api/settings/reboot"
)

var (
	router   string
	username string
	password string
	salt     string
)

func init() {
	flag.StringVar(&router, "router", "192.168.1.1", "router ip address")
	flag.StringVar(&username, "user", "admin", "router admin username")
	flag.StringVar(&password, "password", "", "router admin password")
	flag.Parse()
}

func main() {
	if password == "" {
		flag.Usage()
		os.Exit(0)
	}

	session := NewSession()
	session.GetARPTable()
	session.Logout()
}
