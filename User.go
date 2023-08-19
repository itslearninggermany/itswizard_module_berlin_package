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

/*
Result are all Users which has to be in the system.
*/
func (p *BerlinBsp) SearchUser() (users []User, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	u, err := url.Parse(p.fqdn + api + urlUser)
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return users, err
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return users, err
		}

		err = json.Unmarshal(buf.Bytes(), &users)
		if err != nil {
			return users, err
		}
		return users, nil
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return users, err
	}

	return nil, errors.New(buf.String())
}

/*
Minutes. Result is uodate, delete and import.
*/
func (p *BerlinBsp) SearchUserDelta(minutes uint) (userToUpdate, userToImport, userToDelete []User, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	data := url.Values{}
	data.Set("dtLetzteAenderung", time.Now().Add(-time.Minute*time.Duration(minutes)).Format(TimeLayout))
	encodedData := data.Encode()

	fmt.Println(encodedData)

	u, err := url.Parse(p.fqdn + api + urlUser + "?" + encodedData)
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, err := client.Do(request)
	var users []User

	if err != nil {
		return
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return
		}

		err = json.Unmarshal(buf.Bytes(), &users)
		if err != nil {
			return
		}

		for _, v := range users {
			if v.BenutzerAktionCode == "C" {
				userToImport = append(userToImport, v)
			}
			if v.BenutzerAktionCode == "D" {
				userToDelete = append(userToDelete, v)
			}
			if v.BenutzerAktionCode == "U" {
				userToUpdate = append(userToUpdate, v)
			}
		}

		return
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return
	}

	err = errors.New(buf.String())
	return
}

/*
Result are all Users which has to be in the system And Carer.
*/
func (p *BerlinBsp) SearchUserWithParent() (users []User, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	data := url.Values{}
	data.Set("einschluss", "Bezugspersonen")
	//	data.Set("dtLetzteAenderung", time.Now().Add(-time.Minute*500).Format(TimeLayout))
	encodedData := data.Encode()

	fmt.Println(encodedData)

	u, err := url.Parse(p.fqdn + api + urlUser + "?" + encodedData)

	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return users, err
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return users, err
		}

		err = json.Unmarshal(buf.Bytes(), &users)
		if err != nil {
			return users, err
		}
		return users, nil
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return users, err
	}

	return nil, errors.New(buf.String())
}

/*
Gives the changes of the last 5 Minutes
*/
func (p *BerlinBsp) SearchUserDeltaWithParent(minutes uint) (userToUpdate, userToImport, userToDelete []User, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	data := url.Values{}
	data.Set("einschluss", "Bezugspersonen")
	data.Set("dtLetzteAenderung", time.Now().Add(-time.Minute*time.Duration(minutes)).Format(TimeLayout))
	encodedData := data.Encode()

	fmt.Println(encodedData)

	u, err := url.Parse(p.fqdn + api + urlUser + "?" + encodedData)
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, err := client.Do(request)
	var users []User

	if err != nil {
		return
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return
		}

		err = json.Unmarshal(buf.Bytes(), &users)
		if err != nil {
			return
		}

		for _, v := range users {
			if v.BenutzerAktionCode == "C" {
				userToImport = append(userToImport, v)
			}
			if v.BenutzerAktionCode == "D" {
				userToDelete = append(userToDelete, v)
			}
			if v.BenutzerAktionCode == "U" {
				userToUpdate = append(userToUpdate, v)
			}
		}

		return
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return
	}

	err = errors.New(buf.String())
	return
}

/*
Result are all Users which has to be in the system.
*/
func (p *BerlinBsp) SearchUserBySchool(schulID string) (users []User, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	data := url.Values{}
	data.Set("schuleUID", schulID)
	encodedData := data.Encode()

	fmt.Println(encodedData)

	u, err := url.Parse(p.fqdn + api + urlUser + "?" + encodedData)

	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return users, err
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return users, err
		}

		err = json.Unmarshal(buf.Bytes(), &users)
		if err != nil {
			return users, err
		}
		return users, nil
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return users, err
	}

	return nil, errors.New(buf.String())
}

/*
Minutes. Result is uodate, delete and import.
*/
func (p *BerlinBsp) SearchUserDeltaBySchool(schulID string, minutes uint) (userToUpdate, userToImport, userToDelete []User, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	data := url.Values{}
	data.Set("schuleUID", schulID)
	data.Set("dtLetzteAenderung", time.Now().Add(-time.Minute*time.Duration(minutes)).Format(TimeLayout))
	encodedData := data.Encode()

	fmt.Println(encodedData)

	u, err := url.Parse(p.fqdn + api + urlUser + "?" + encodedData)
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, err := client.Do(request)
	var users []User

	if err != nil {
		return
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return
		}

		err = json.Unmarshal(buf.Bytes(), &users)
		if err != nil {
			return
		}

		for _, v := range users {
			if v.BenutzerAktionCode == "C" {
				userToImport = append(userToImport, v)
			}
			if v.BenutzerAktionCode == "D" {
				userToDelete = append(userToDelete, v)
			}
			if v.BenutzerAktionCode == "U" {
				userToUpdate = append(userToUpdate, v)
			}
		}

		return
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return
	}

	err = errors.New(buf.String())
	return
}

/*
Result are all Users which has to be in the system And Carer.
*/
func (p *BerlinBsp) SearchUserWithParentBySchool(schulID string) (users []User, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	data := url.Values{}
	data.Set("schuleUID", schulID)
	data.Set("einschluss", "Bezugspersonen")
	//	data.Set("dtLetzteAenderung", time.Now().Add(-time.Minute*500).Format(TimeLayout))
	encodedData := data.Encode()

	fmt.Println(encodedData)

	u, err := url.Parse(p.fqdn + api + urlUser + "?" + encodedData)

	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return users, err
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return users, err
		}

		err = json.Unmarshal(buf.Bytes(), &users)
		if err != nil {
			return users, err
		}
		return users, nil
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return users, err
	}

	return nil, errors.New(buf.String())
}

/*
Gives the changes of the last 5 Minutes
*/
func (p *BerlinBsp) SearchUserDeltaWithParentBySchool(schulID string, minutes uint) (userToUpdate, userToImport, userToDelete []User, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	data := url.Values{}
	data.Set("schuleUID", schulID)
	data.Set("einschluss", "Bezugspersonen")
	data.Set("dtLetzteAenderung", time.Now().Add(-time.Minute*time.Duration(minutes)).Format(TimeLayout))
	encodedData := data.Encode()

	fmt.Println(encodedData)

	u, err := url.Parse(p.fqdn + api + urlUser + "?" + encodedData)
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, err := client.Do(request)
	var users []User

	if err != nil {
		return
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return
		}

		err = json.Unmarshal(buf.Bytes(), &users)
		if err != nil {
			return
		}

		for _, v := range users {
			if v.BenutzerAktionCode == "C" {
				userToImport = append(userToImport, v)
			}
			if v.BenutzerAktionCode == "D" {
				userToDelete = append(userToDelete, v)
			}
			if v.BenutzerAktionCode == "U" {
				userToUpdate = append(userToUpdate, v)
			}
		}

		return
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return
	}

	err = errors.New(buf.String())
	return
}

/*
 */
func (p *BerlinBsp) GetUser(userID string) (user User, err error) {
	err = p.authentificate()
	if err != nil {
		return
	}

	data := url.Values{}
	data.Set("einschluss", "Bezugspersonen")
	encodedData := data.Encode()

	fmt.Println(encodedData)

	u, err := url.Parse(p.fqdn + api + urlUser + "/" + userID + "?" + encodedData)

	fmt.Println(u)

	request, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return
	}

	request.Header.Set("Authorization", authorizationPrefix+p.access_token)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return
	}

	if response.StatusCode == 200 {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return
		}
		var temp []User
		err = json.Unmarshal(buf.Bytes(), &temp)
		if err != nil {
			return
		}

		if len(temp) < 1 {
			err = errors.New("No User delivered")
		}
		user = temp[0]
		return
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return
	}

	err = errors.New(buf.String())
	return
}
