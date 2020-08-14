package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"cursmedia.com/rakuten/model"
)

const udemySID = 39197

var tokenRequest = os.Getenv("TOKEN_REQUEST")
var myUser = os.Getenv("LINKSHARE_USER")
var myPwd = os.Getenv("LINKSHARE_PWD")
var mySID = os.Getenv("LINKSHARE_SID_ID")

// RakutenClient rc
type RakutenClient struct {
	token  model.Token
	client http.Client
}

// NewRakutenClient creates a rakutenClient instance
func NewRakutenClient() RakutenClient {
	rc := RakutenClient{client: http.Client{
		Timeout: time.Second * 20, // Maximum of 20 secs
	}}
	rc.token = rc.GetToken()
	return rc
}

// MakeRequest makes a request to the rakuten api
func (rc RakutenClient) MakeRequest(pagenumber int) (string, error) {
	productURL := "https://api.rakutenmarketing.com/productsearch/1.0?mid=" + string(udemySID) + "&max=100&pagenumber=" + strconv.Itoa(pagenumber)

	req, err := http.NewRequest(http.MethodGet, productURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json, application/xml, */*;q=0.01")
	req.Header.Add("Authorization", "Bearer "+rc.token.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	res, getErr := rc.client.Do(req)
	if getErr != nil {
		return "", getErr
	}
	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			rc.token = rc.GetToken()
			return rc.MakeRequest(pagenumber)
		}
		body, _ := ioutil.ReadAll(res.Body)
		log.Fatal(string(body))
	}
	defer res.Body.Close()
	encoding := res.Header.Get("Content-Type")
	log.Println(encoding)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", getErr
	}
	return string(body), nil
}

// GetToken returns a new token
func (rc RakutenClient) GetToken() model.Token {
	url := "https://api.rakutenmarketing.com/token"

	reqBody := "grant_type=password&username=" + myUser + "&password=" + myPwd + "&scope=" + string(mySID)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json, text/javascript, */*;q=0.01")
	req.Header.Add("Authorization", tokenRequest)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, getErr := rc.client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var token model.Token
	json.Unmarshal(body, &token)
	log.Print(token.AccessToken)
	log.Print(string(body))
	return token
}
