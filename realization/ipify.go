package realization

import (
	"bytes"
	"io"
	"net/http"
)

type Ipify struct{}

func (i *Ipify) GetCurrentIP() (string, error) {
	ip, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(ip.Body)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(ip.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
