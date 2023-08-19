package itswizard_module_berlin_package

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func (p *BerlinBsp) SearchClasses() (classes []Class, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	u, err := url.Parse(p.fqdn + api + urlClass)
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return classes, err
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return classes, err
		}

		err = json.Unmarshal(buf.Bytes(), &classes)
		if err != nil {
			return classes, err
		}
		return classes, nil
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return classes, err
	}

	return nil, errors.New(buf.String())
}

func (p *BerlinBsp) SearchClassesBySchool(schulID string) (classes []Class, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	data := url.Values{}
	data.Set("schuleUID", schulID)
	encodedData := data.Encode()

	fmt.Println(encodedData)

	u, err := url.Parse(p.fqdn + api + urlClass + "?" + encodedData)

	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return classes, err
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return classes, err
		}

		err = json.Unmarshal(buf.Bytes(), &classes)
		if err != nil {
			return classes, err
		}
		return classes, nil
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return classes, err
	}

	return nil, errors.New(buf.String())
}
