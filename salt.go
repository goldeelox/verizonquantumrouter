package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type loginError struct {
	PasswordSalt string `json:"passwordSalt"`
}

func getRouterSalt() string {
	r, err := http.Get("http://" + router + loginAPI)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	var m loginError
	j := json.NewDecoder(r.Body)
	j.Decode(&m)

	return m.PasswordSalt
}
