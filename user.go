package cloudflare

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// User describes a user account.
type User struct {
	ID            string         `json:"id,omitempty"`
	Email         string         `json:"email,omitempty"`
	FirstName     string         `json:"first_name,omitempty"`
	LastName      string         `json:"last_name,omitempty"`
	Username      string         `json:"username,omitempty"`
	Telephone     string         `json:"telephone,omitempty"`
	Country       string         `json:"country,omitempty"`
	Zipcode       string         `json:"zipcode,omitempty"`
	CreatedOn     *time.Time     `json:"created_on,omitempty"`
	ModifiedOn    *time.Time     `json:"modified_on,omitempty"`
	APIKey        string         `json:"api_key,omitempty"`
	TwoFA         bool           `json:"two_factor_authentication_enabled,omitempty"`
	Betas         []string       `json:"betas,omitempty"`
	Organizations []Organization `json:"organizations,omitempty"`
}

// UserResponse wraps a response containing User accounts.
type UserResponse struct {
	Response
	Result User `json:"result"`
}

// UserDetails provides information about the logged-in user.
// API reference:
// 	https://api.cloudflare.com/#user-user-details
//	GET /user
func (api *API) UserDetails() (User, error) {
	var r UserResponse
	res, err := api.makeRequest("GET", "/user", nil)
	if err != nil {
		return User{}, errors.Wrap(err, errMakeRequestError)
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return User{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}

// UpdateUser updates the properties of the given user.
// API reference:
// 	https://api.cloudflare.com/#user-update-user
//	PATCH /user
func (api *API) UpdateUser(user *User) (User, error) {
	var r UserResponse
	res, err := api.makeRequest("PATCH", "/user", user)
	if err != nil {
		return User{}, errors.Wrap(err, errMakeRequestError)
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return User{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}
