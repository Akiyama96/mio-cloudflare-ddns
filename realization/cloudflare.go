package realization

import (
	"bytes"
	"ddns/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Account struct {
	Email  string `json:"email"`
	ApiKey string `json:"api_key"`
	ZoneID string `json:"zone_id"`
}

const (
	CLOUDFLARE_API_URL = "https://api.cloudflare.com/client/v4"
)

func (a *Account) ListDNSRecords() (*models.DNSRecordsResponse, error) {
	url := fmt.Sprintf("%s/zones/%s/dns_records", CLOUDFLARE_API_URL, a.ZoneID)

	putRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	putRequest.Header.Set("X-Auth-Email", a.Email)
	putRequest.Header.Set("X-Auth-Key", a.ApiKey)

	client := http.DefaultClient
	client.Timeout = time.Second * 3

	response, err := client.Do(putRequest)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("request list DNS records failed, response status code not ok")
	}

	body := new(bytes.Buffer)
	_, err = body.ReadFrom(response.Body)
	if err != nil {
		return nil, err
	}

	DNSRecordsResponse := &models.DNSRecordsResponse{}
	err = json.Unmarshal(body.Bytes(), DNSRecordsResponse)
	if err != nil {
		return nil, err
	}

	if !DNSRecordsResponse.Success {
		return nil, errors.New("request list DNS records failed, response success==false")
	}

	return DNSRecordsResponse, nil
}

func (a *Account) OverwriteDNSRecord(ddnsSetting *models.DDNSSetting, recordId, ip string) error {
	url := fmt.Sprintf("%s/zones/%s/dns_records/%s", CLOUDFLARE_API_URL, a.ZoneID, recordId)

	payload := fmt.Sprintf(`{"type":"%s","name":"%s","content":"%s","ttl":%d,"proxied":%t}`,
		ddnsSetting.ServerInfo.Type,
		ddnsSetting.ServerInfo.Name,
		ip,
		ddnsSetting.ServerInfo.Ttl,
		ddnsSetting.ServerInfo.Proxied,
	)

	putRequest, err := http.NewRequest("PUT", url, bytes.NewBufferString(payload))
	if err != nil {
		return err
	}

	putRequest.Header.Set("X-Auth-Email", a.Email)
	putRequest.Header.Set("X-Auth-Key", a.ApiKey)
	putRequest.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	client.Timeout = time.Second * 3

	response, err := client.Do(putRequest)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return errors.New("request overwrite DNS record failed, response status code not ok")
	}

	body := new(bytes.Buffer)
	_, err = body.ReadFrom(response.Body)
	if err != nil {
		return err
	}

	log.Println("Response: ", body.String())

	return nil
}
