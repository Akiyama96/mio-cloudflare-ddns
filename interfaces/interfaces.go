package interfaces

import "ddns/models"

type CloudflareApi interface {
	ListDNSRecords() (*models.DNSRecordsResponse, error)
	OverwriteDNSRecord(*models.DDNSSetting, string, string) error
}

type IpifyApi interface {
	GetCurrentIP() (string, error)
}
