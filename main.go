package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type SiminnSMS struct {
	URL      string
	Username string
	Password string
	SendFrom string
}

func (s *SiminnSMS) SendSMS(text string, number string) (error, bool) {
	response, err := http.PostForm(s.URL+"push", url.Values{
		"l":      {s.Username},
		"p":      {s.Password},
		"msisdn": {number},
		"T":      {text},
		"A":      {s.SendFrom},
	},
	)

	if err != nil {
		return err, false
	}

	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err, false
	}

	if response.StatusCode != 200 {
		return nil, false
	}

	return nil, true
}
