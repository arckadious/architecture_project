package client

import (
	"auth-api-crossfitlov/models/structs/out"
	"auth-api-crossfitlov/parameters"
	"bytes"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

//GetUserInfos read user infos
func GetUserInfos(ID string) (out.UserInfos, error) {
	url := parameters.Config.UserAPI.URL + "/users/" + ID

	logrus.Info(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return out.UserInfos{}, err
	}

	req.SetBasicAuth(parameters.Config.UserAPI.Username, parameters.Config.UserAPI.Password)

	httpTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	httpClient := &http.Client{
		Timeout:   10 * time.Minute,
		Transport: httpTransport,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return out.UserInfos{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return out.UserInfos{}, errors.New("Status code != 200 ( = " + strconv.Itoa(resp.StatusCode) + ")")
	}

	var userInfos out.UserInfos
	err = json.NewDecoder(resp.Body).Decode(&userInfos)
	if err != nil {
		return out.UserInfos{}, err
	}

	return userInfos, nil

}

// CreateUserInfos creates a new user with main informations
func CreateUserInfos(user out.UserInfos) error {
	url := parameters.Config.UserAPI.URL + "/users"

	buffer := bytes.NewBuffer([]byte{})
	json.NewEncoder(buffer).Encode(user)

	req, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		return err
	}

	req.SetBasicAuth(parameters.Config.UserAPI.Username, parameters.Config.UserAPI.Password)

	httpTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	httpClient := &http.Client{
		Timeout:   10 * time.Minute,
		Transport: httpTransport,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Status code != 200 ( = " + strconv.Itoa(resp.StatusCode) + ")")
	}

	return nil

}

//DeleteUserInfos read user infos
func DeleteUserInfos(ID string) error {
	url := parameters.Config.UserAPI.URL + "/users/" + ID

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(parameters.Config.UserAPI.Username, parameters.Config.UserAPI.Password)

	httpTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	httpClient := &http.Client{
		Timeout:   10 * time.Minute,
		Transport: httpTransport,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Status code != 200 ( = " + strconv.Itoa(resp.StatusCode) + ")")
	}

	return nil

}
