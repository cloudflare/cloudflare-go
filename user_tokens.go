package cloudflare

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// UserTokensPermissionGroup contains permissions for the resources.
type UserTokensPermissionGroup struct {
	ID     string   `json:"id,omitempty"`
	Name   string   `json:"name,omitempty"`
	Label  string   `json:"label,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

// UserTokensPermissionGroupsResponse wraps a response containing slice of UserTokensPermissionGroups.
type UserTokensPermissionGroupsResponse struct {
	Response
	Result []UserTokensPermissionGroup `json:"result"`
}

// UserToken describes an user token token with all the policies with permissions for given resources.
type UserToken struct {
	ID         string            `json:"id,omitempty"`
	Name       string            `json:"name,omitempty"`
	Value      string            `json:"value,omitempty"`
	Status     string            `json:"status,omitempty"`
	IssuedOn   *time.Time        `json:"issued_on,omitempty"`
	ModifiedOn *time.Time        `json:"modified_on,omitempty"`
	LastUsedOn *time.Time        `json:"last_used_on,omitempty"`
	Policies   []UserTokenPolicy `json:"policies,omitempty"`
}

// UserTokenPolicy contains policy with permissions for given resources.
type UserTokenPolicy struct {
	ID               string                      `json:"id,omitempty"`
	Effect           string                      `json:"effect,omitempty"`
	Resources        map[string]string           `json:"resources,omitempty"`
	PermissionGroups []UserTokensPermissionGroup `json:"permission_groups,omitempty"`
}

// UserTokenResponse wraps a response containing a UserToken:
type UserTokenResponse struct {
	Response
	Result UserToken `json:"result"`
}

// UserTokensResponse wraps a response containing slice of UserTokens.
type UserTokensResponse struct {
	Response
	Result []UserToken `json:"result"`
}

// UserTokenID contains id of an user token.
type UserTokenID struct {
	ID string `json:"id"`
}

// UserTokenIDResponse wraps a response containing UserTokenID.
type UserTokenIDResponse struct {
	Response
	Result UserTokenID `json:"result"`
}

// UserTokens returns slice of all created UserTokens.
func (api *API) ListUserTokens() ([]UserToken, error) {
	var r UserTokensResponse
	res, err := api.makeRequest("GET", "/user/tokens", nil)
	if err != nil {
		return []UserToken{}, errors.Wrap(err, errMakeRequestError)
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return []UserToken{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}

// UserToken returns UserToken for given ID.
func (api *API) UserToken(tokenID string) (UserToken, error) {
	var r UserTokenResponse
	res, err := api.makeRequest("GET", "/user/tokens/"+tokenID, nil)
	if err != nil {
		return UserToken{}, errors.Wrap(err, errMakeRequestError)
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return UserToken{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}

// CreateUserToken creates UserToken with given name and permission groups
func (api *API) CreateUserToken(name string, policies []UserTokenPolicy) (UserToken, error) {
	t := UserToken{
		Name:     name,
		Policies: policies,
	}

	var r UserTokenResponse
	res, err := api.makeRequest("POST", "/user/tokens", t)
	if err != nil {
		return UserToken{}, errors.Wrap(err, errMakeRequestError)
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return UserToken{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}

// UpdateUserToken can update name or permissions of given UserToken.
func (api *API) UpdateUserToken(t UserToken) (UserToken, error) {
	var r UserTokenResponse
	res, err := api.makeRequest("PUT", "/user/tokens/"+t.ID, t)
	if err != nil {
		return UserToken{}, errors.Wrap(err, errMakeRequestError)
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return UserToken{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}

// DeleteUserToken deletes UserToken.
func (api *API) DeleteUserToken(tokenID string) (UserTokenID, error) {
	var r UserTokenIDResponse
	res, err := api.makeRequest("DELETE", "/user/tokens/"+tokenID, nil)
	if err != nil {
		return UserTokenID{}, errors.Wrap(err, errMakeRequestError)
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return UserTokenID{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}

// UserTokensPermissionGroups returns slice of all the possible PermissionGroups.
func (api *API) UserTokensPermissionGroups() ([]UserTokensPermissionGroup, error) {
	var r UserTokensPermissionGroupsResponse
	res, err := api.makeRequest("GET", "/user/tokens/permission_groups", nil)
	if err != nil {
		return []UserTokensPermissionGroup{}, errors.Wrap(err, errMakeRequestError)
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return []UserTokensPermissionGroup{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}
