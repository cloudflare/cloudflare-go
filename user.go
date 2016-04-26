package cloudflare

import (
	"encoding/json"

	pkgErrors "github.com/pkg/errors"
)

/*
Information about the logged-in user.

API reference: https://api.cloudflare.com/#user-user-details
*/
func (api API) UserDetails() (User, error) {
	var r UserResponse
	res, err := api.makeRequest("GET", "/user", nil)
	if err != nil {
		return User{}, pkgErrors.Wrap(err, "Error from makeRequest")
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return User{}, pkgErrors.Wrap(err, "Error from unmarshal")
	}
	return r.Result, nil
}

/*
Update user properties.

API reference: https://api.cloudflare.com/#user-update-user
*/
func (api API) UpdateUser() (User, error) {
	// api.makeRequest("PATCH", "/user", user)
	return User{}, nil
}
