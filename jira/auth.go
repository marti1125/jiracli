package jira

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Auth struct {
	SiteUrl string `json:"site_url"`
	Email   string `json:"email"`
	Token   string `json:"token"`
}

func readFile() *Auth {
	configFile, err := os.Stat("config.json")
	if err != nil {
		fmt.Println(err)
	}
	if configFile.Size() > 0 {
		content, err := ioutil.ReadFile("config.json")
		if err != nil {
			fmt.Println(err)
		}
		c := string(content)
		var config Auth
		err = json.Unmarshal([]byte(c), &config)
		if err != nil {
			fmt.Println(err)
		}
		return &config
	}
	return nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func Request(httpMethod string, resource string, body io.Reader) string {
	auth := readFile()
	if auth == nil {
		return errors.New("can not parse config.json").Error()
	}
	client := http.Client{}
	req, err := http.NewRequest(httpMethod, auth.SiteUrl+resource, body)
	req.Header.Add("X-Atlassian-Token", "no-check")
	req.Header.Add("Authorization", "Basic "+basicAuth(auth.Email, auth.Token))
	if err != nil {
		return err.Error()
	}
	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	return string(b)
}
