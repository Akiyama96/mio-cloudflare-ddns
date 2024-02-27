package models

import "time"

type AccountSetting struct {
	Email  string `json:"email"`
	ApiKey string `json:"api_key"`
	ZoneID string `json:"zone_id"`
}

type DDNSSetting struct {
	ServerInfo    ServerInfo    `json:"server_info"`
	ServerSetting ServerSetting `json:"server_setting"`
}

type ServerInfo struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Ttl     int    `json:"ttl"`
	Proxied bool   `json:"proxied"`
}

type ServerSetting struct {
	IntervalsMin int `json:"intervals_min"`
}

type DNSRecordsResponse struct {
	Result []struct {
		Id        string `json:"id"`
		ZoneId    string `json:"zone_id"`
		ZoneName  string `json:"zone_name"`
		Name      string `json:"name"`
		Type      string `json:"type"`
		Content   string `json:"content"`
		Proxiable bool   `json:"proxiable"`
		Proxied   bool   `json:"proxied"`
		Ttl       int    `json:"ttl"`
		Locked    bool   `json:"locked"`
		Meta      struct {
			AutoAdded           bool   `json:"auto_added"`
			ManagedByApps       bool   `json:"managed_by_apps"`
			ManagedByArgoTunnel bool   `json:"managed_by_argo_tunnel"`
			Source              string `json:"source"`
		} `json:"meta"`
		Comment    interface{}   `json:"comment"`
		Tags       []interface{} `json:"tags"`
		CreatedOn  time.Time     `json:"created_on"`
		ModifiedOn time.Time     `json:"modified_on"`
	} `json:"result"`
	Success    bool          `json:"success"`
	Errors     []interface{} `json:"errors"`
	Messages   []interface{} `json:"messages"`
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		Count      int `json:"count"`
		TotalCount int `json:"total_count"`
		TotalPages int `json:"total_pages"`
	} `json:"result_info"`
}

type OverwriteDNSRecordResponse struct {
	Result struct {
		Id        string `json:"id"`
		ZoneId    string `json:"zone_id"`
		ZoneName  string `json:"zone_name"`
		Name      string `json:"name"`
		Type      string `json:"type"`
		Content   string `json:"content"`
		Proxiable bool   `json:"proxiable"`
		Proxied   bool   `json:"proxied"`
		Ttl       int    `json:"ttl"`
		Locked    bool   `json:"locked"`
		Meta      struct {
			AutoAdded           bool   `json:"auto_added"`
			ManagedByApps       bool   `json:"managed_by_apps"`
			ManagedByArgoTunnel bool   `json:"managed_by_argo_tunnel"`
			Source              string `json:"source"`
		} `json:"meta"`
		Comment    interface{}   `json:"comment"`
		Tags       []interface{} `json:"tags"`
		CreatedOn  time.Time     `json:"created_on"`
		ModifiedOn time.Time     `json:"modified_on"`
	} `json:"result"`
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
}
