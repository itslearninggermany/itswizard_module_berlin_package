package itswizard_module_berlin_package

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func (p *BerlinBsp) Authentification() error {
	httpposturl := p.fqdn + "/auth/realms/BSP/protocol/openid-connect/token"

	data := url.Values{}
	data.Set("grant_type", p.grant_type)
	data.Set("client_id", p.client_id)
	data.Add("client_secret", p.client_secret)
	encodedData := data.Encode()

	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer([]byte(encodedData)))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}

	if response.Status == "200 OK" {
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(response.Body)
		if err != nil {
			return err
		}
		var authentification Authentification
		err = json.Unmarshal(buf.Bytes(), &authentification)
		if err != nil {
			return err
		}
		p.access_token = authentification.AccessToken
		p.token_type = authentification.TokenType
		p.authentification = true
		p.createdTokenTime = time.Now()
		return nil
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	return errors.New(buf.String())
}

type Authentification struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	Scope            string `json:"scope"`
}
