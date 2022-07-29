package client

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
)

func NewDHAClient() *Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	client := http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}

	return &Client{
		HTTPClient: client,
	}
}

func (c *Client) DHARequest(url string) (models.DHAResponse, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.DHAResponse{}, err
	}

	response, err := c.HTTPClient.Do(request)
	if err != nil || response.StatusCode >= 300 {
		if response != nil {
			log.DHAError.Printf("DHA Request failed: %s", response.Status)
			if body, err := ioutil.ReadAll(response.Body); err != nil {
				log.DHAError.Printf("DHA Response body: %s", body)
			}
			return models.DHAResponse{}, errors.New("DHA Request failed " + err.Error())
		}
		return models.DHAResponse{}, err
	}
	defer response.Body.Close()

	var dhaData models.DHAResponse
	err = xml.NewDecoder(response.Body).Decode(&dhaData)
	if err != nil {
		log.DHAError.Printf("Failed to decode dha data: %s", err)
		return models.DHAResponse{}, err
	}

	return dhaData, nil
}

func (c *Client) DHASimulatorRequest(url string) (models.DHASimulatorResponse, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.DHASimulatorResponse{}, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := c.HTTPClient.Do(request)
	if err != nil || response.StatusCode >= 300 {
		if response != nil {
			log.DHAError.Printf("DHA Request failed: %s", response.Status)
			if body, err := ioutil.ReadAll(response.Body); err != nil {
				log.DHAError.Printf("DHA Response body: %s", body)
			}
			return models.DHASimulatorResponse{}, errors.New("DHA Request failed " + err.Error())
		}
		return models.DHASimulatorResponse{}, err
	}
	defer response.Body.Close()

	var dhaData models.DHASimulatorResponse
	err = json.NewDecoder(response.Body).Decode(&dhaData)
	if err != nil {
		log.DHAError.Printf("Failed to decode dha data: %s", err)
		return models.DHASimulatorResponse{}, err
	}

	return dhaData, nil
}

func (c *Client) DHAQueryRequest(url string) (models.DHAUser, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.DHAUser{}, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := c.HTTPClient.Do(request)
	if err != nil || response.StatusCode >= 300 {
		if response != nil {
			log.DHAError.Printf("DHA Request failed: %s", response.Status)
			if body, err := ioutil.ReadAll(response.Body); err != nil {
				log.DHAError.Printf("DHA Response body: %s", body)
			}
			return models.DHAUser{}, errors.New("DHA Request failed " + err.Error())
		}
		return models.DHAUser{}, err
	}
	defer response.Body.Close()

	var dhaData models.DHAUser
	err = json.NewDecoder(response.Body).Decode(&dhaData)
	if err != nil {
		log.DHAError.Printf("Failed to decode dha data: %s", err)
		return models.DHAUser{}, err
	}

	return dhaData, nil
}
