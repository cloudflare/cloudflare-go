package cloudflare

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	pkgErrors "github.com/pkg/errors"
)

/*
IPs gets a list of CloudFlare's IP ranges

This does not require logging in to the API.

API reference:
  https://api.cloudflare.com/#cloudflare-ips
  GET /client/v4/ips
*/
func IPs() (IPRanges, error) {
	resp, err := http.Get(apiURL + "/ips")
	if err != nil {
		return IPRanges{}, pkgErrors.Wrap(err, "HTTP request failed")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return IPRanges{}, pkgErrors.Wrap(err, "Response body could not be read")
	}
	var r IPsResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return IPRanges{}, pkgErrors.Wrap(err, "Error from unmarshal")
	}
	return r.Result, nil
}
