package services

import (
	"ddns/interfaces"
	"ddns/models"
	"ddns/realization"
	"encoding/json"
	"errors"
	"log"
	"os"
)

func autoOverwrite(autoOverwriteServerSetting *models.DDNSSetting) error {
	log.Println("Read account setting.")
	accountSetting, err := readAccountConfig()
	if err != nil {
		return err
	}

	ipifyEntry := &realization.Ipify{}
	ipifyApi := interfaces.IpifyApi(ipifyEntry)

	log.Println("Getting current ip.")
	currentIP, err := ipifyApi.GetCurrentIP()
	if err != nil {
		return err
	}

	cloudflareEntry := &realization.Account{
		Email:  accountSetting.Email,
		ApiKey: accountSetting.ApiKey,
		ZoneID: accountSetting.ZoneID,
	}
	cloudflareApi := interfaces.CloudflareApi(cloudflareEntry)

	log.Println("Finding dns record.")
	dnsRecordsResponse, err := cloudflareApi.ListDNSRecords()
	if err != nil {
		return err
	}

	recordID, err := findRecordIDByDomain(autoOverwriteServerSetting.ServerInfo.Name, dnsRecordsResponse)
	if err != nil {
		return err
	}

	log.Println("Overwriting dns record.")
	return cloudflareApi.OverwriteDNSRecord(autoOverwriteServerSetting, recordID, currentIP)
}

func findRecordIDByDomain(domain string, dnsRecordsResponse *models.DNSRecordsResponse) (string, error) {
	dnsRecords := dnsRecordsResponse.Result
	for _, record := range dnsRecords {
		if record.Name == domain {
			return record.Id, nil
		}
	}

	return "", errors.New("record not found")
}

func readAccountConfig() (*models.AccountSetting, error) {
	accountFileContent, err := os.ReadFile("./configs/account.json")
	if err != nil {
		return nil, err
	}

	accountSetting := &models.AccountSetting{}
	err = json.Unmarshal(accountFileContent, accountSetting)
	if err != nil {
		return nil, err
	}

	return accountSetting, nil
}

func readAutoOverwriteServerConfig() (*models.DDNSSetting, error) {
	autoOverwriteServerFileContent, err := os.ReadFile("./configs/auto_overwrite.json")
	if err != nil {
		return nil, err
	}

	autoOverwriteServerSetting := &models.DDNSSetting{}
	err = json.Unmarshal(autoOverwriteServerFileContent, autoOverwriteServerSetting)
	if err != nil {
		return nil, err
	}

	return autoOverwriteServerSetting, nil
}
