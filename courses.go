package itswizard_module_berlin_package

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func (p *BerlinBsp) SearchCourses() (courses []Course, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	u, err := url.Parse(p.fqdn + api + urlCourse)
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return courses, err
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return courses, err
		}

		err = json.Unmarshal(buf.Bytes(), &courses)
		if err != nil {
			return courses, err
		}
		return courses, nil
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return courses, err
	}

	return nil, errors.New(buf.String())
}

func (p *BerlinBsp) SearchCoursesBySchool(schulID string) (courses []Course, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	data := url.Values{}
	data.Set("schuleUID", schulID)
	encodedData := data.Encode()

	fmt.Println(encodedData)

	u, err := url.Parse(p.fqdn + api + urlCourse + "?" + encodedData)

	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return courses, err
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return courses, err
		}

		err = json.Unmarshal(buf.Bytes(), &courses)
		if err != nil {
			return courses, err
		}
		return courses, nil
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return courses, err
	}

	return nil, errors.New(buf.String())
}
