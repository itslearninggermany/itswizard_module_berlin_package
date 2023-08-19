package itswizard_module_berlin_package

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

/*
All Schools are delivered
*/
func (p *BerlinBsp) SearchSchools() (schools []School, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	u, err := url.Parse(p.fqdn + api + urlSchool)
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return schools, err
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return schools, err
		}

		err = json.Unmarshal(buf.Bytes(), &schools)
		if err != nil {
			return schools, err
		}
		return schools, nil
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return schools, err
	}

	return nil, errors.New(buf.String())

}

/*
One School is delivered
*/
func (p *BerlinBsp) Getschool(schulUid string) (school School, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	u, err := url.Parse(p.fqdn + api + urlSchool + "?schuleUID=" + schulUid)
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return school, err
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return school, err
		}
		var tmp []School
		err = json.Unmarshal(buf.Bytes(), &tmp)
		if err != nil {
			return school, err
		}
		if len(tmp) < 1 {
			return school, nil
		} else {
			school = tmp[0]
			return school, nil
		}

	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return school, err
	}

	return school, errors.New(buf.String())

}
