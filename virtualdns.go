package cloudflare

import (
    "encoding/json"
    "github.com/pkg/errors"
)

/*
Create a Virtual DNS cluster

API reference:
  https://api.cloudflare.com/#virtual-dns-users--create-a-virtual-dns-cluster
  POST /user/virtual_dns
*/
func (api *API) CreateVirtualDNS(v VirtualDNS) (VirtualDNS, error) {
	res, err := api.makeRequest("POST", "/user/virtual_dns", v)
	if err != nil {
		return VirtualDNS{}, errors.Wrap(err, errMakeRequestError)
	}

	var recordResp *VirtualDNSResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return VirtualDNS{}, errors.Wrap(err, errUnmarshalError)
	}

	return recordResp.Result, nil
}

/*
List Virtual DNS clusters

API reference:
  https://api.cloudflare.com/#virtual-dns-users--get-virtual-dns-clusters
  GET /user/virtual_dns
*/
func (api *API) ListVirtualDNS() ([]VirtualDNS, error) {
    var v VirtualDNSListResponse
    res, err := api.makeRequest("GET", "/user/virtual_dns", nil)
    if err != nil {
        return []VirtualDNS{}, errors.Wrap(err, errMakeRequestError)
    }
    err = json.Unmarshal(res, &v)
    if err != nil {
        return []VirtualDNS{}, errors.Wrap(err, errUnmarshalError)
    }
    return v.Result, nil
}

/*
Fetch a Virtual DNS cluster

API reference:
  https://api.cloudflare.com/#virtual-dns-users--get-a-virtual-dns-cluster
  GET /user/virtual_dns/:identifier
*/
func (api *API) VirtualDNS(virtualDNSID string) (VirtualDNS, error) {
	uri := "/user/virtual_dns/" + virtualDNSID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return VirtualDNS{}, errors.Wrap(err, errMakeRequestError)
	}
	var v VirtualDNSResponse
	err = json.Unmarshal(res, &v)
	if err != nil {
		return VirtualDNS{}, errors.Wrap(err, errUnmarshalError)
	}
	return v.Result, nil
}

/*
Update a Virtual DNS cluster

API reference:

  https://api.cloudflare.com/#virtual-dns-users--modify-a-virtual-dns-cluster
  PATCH /user/virtual_dns/:identifier
*/
func (api *API) UpdateVirtualDNS(virtualDNSID string, vv VirtualDNS) error {
	uri := "/user/virtual_dns/" + virtualDNSID
	res, err := api.makeRequest("PUT", uri, vv)
	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	var v VirtualDNSResponse
	err = json.Unmarshal(res, &v)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}
	return nil
}

/*
Delete a Virtual DNS cluster

API reference:
  https://api.cloudflare.com/#virtual-dns-users--delete-a-virtual-dns-cluster
  DELETE /user/virtual_dns/:identifier
*/
func (api *API) DeleteVirtualDNS(virtualDNSID string) error {
	uri := "/user/virtual_dns/" + virtualDNSID
	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	var v VirtualDNSResponse
	err = json.Unmarshal(res, &v)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}
	return nil
}
