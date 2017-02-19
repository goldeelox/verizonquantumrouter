package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type session struct {
	Client  *http.Client
	Cookies []*http.Cookie
	Token   string
}

func NewSession() session {
	cookies := login(password)
	var token string
	for i := range cookies {
		if cookies[i].Name == "XSRF-TOKEN" {
			token = cookies[i].Value
		}
	}

	return session{Client: &http.Client{}, Cookies: cookies, Token: token}
}

func (c *session) Logout() {
	req, _ := http.NewRequest("GET", "http://"+router+logoutAPI, nil)
	req.Header.Add("X-XSRF-TOKEN", c.Token)
	for i := range c.Cookies {
		req.AddCookie(c.Cookies[i])
	}

	resp, _ := c.Client.Do(req)
	defer resp.Body.Close()

	// TODO remove log
	log.Println(resp.Status)
}

func (c *session) GetARPTable() {
	req, _ := http.NewRequest("GET", "http://"+router+arptableAPI, nil)
	req.Header.Add("X-XSRF-TOKEN", c.Token)
	for i := range c.Cookies {
		req.AddCookie(c.Cookies[i])
	}
	resp, _ := c.Client.Do(req)
	defer resp.Body.Close()

	// TODO remove log
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	log.Println(buf.String())
}

func (c *session) Reboot() {
	req, _ := http.NewRequest("POST", "http://"+router+rebootAPI, nil)
	req.Header.Add("X-XSRF-TOKEN", c.Token)
	for i := range c.Cookies {
		req.AddCookie(c.Cookies[i])
	}

	resp, _ := c.Client.Do(req)
	defer resp.Body.Close()

	log.Println(resp.Status)
}

type loginBody struct {
	Password string `json:"password"`
}

func login(password string) []*http.Cookie {
	// hash password
	salt := getRouterSalt()
	hash := sha512.New()
	hash.Write([]byte(password))
	hash.Write([]byte(salt))
	hashPass := fmt.Sprintf("%x", hash.Sum(nil))

	// login
	reqBody, _ := json.Marshal(loginBody{hashPass})
	r, err := http.Post("http://"+router+loginAPI, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	return r.Cookies()
}
