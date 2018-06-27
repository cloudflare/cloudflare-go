package cloudflare

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// AccountMember is the definition of a member of an account.
type AccountMember struct {
	ID     string                   `json:"id"`
	Code   string                   `json:"code"`
	User   AccountMemberUserDetails `json:"user"`
	Status string                   `json:"status"`
	Roles  []AccountMemberRoles     `json:"roles"`
}

// AccountMemberUserDetails outlines all the personal information about
// a member.
type AccountMemberUserDetails struct {
	ID                             string `json:"id"`
	FirstName                      string `json:"first_name"`
	LastName                       string `json:"last_name"`
	Email                          string `json:"email"`
	TwoFactorAuthenticationEnabled bool
}

// AccountMemberRoles defines the roles that a member can have attached.
type AccountMemberRoles struct {
	ID          string                                  `json:"id"`
	Name        string                                  `json:"name"`
	Description string                                  `json:"description"`
	Permissions map[string]AccountMemberRolePermissions `json:"permissions"`
}

// AccountMemberRolePermissions is the shared structure for all permissions
// that can be assigned to a member.
type AccountMemberRolePermissions struct {
	Read bool `json:"read"`
	Edit bool `json:"edit"`
}

// AccountMembersListResponse represents the response from the list
// account members endpoint.
type AccountMembersListResponse struct {
	Result []AccountMember `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// AccountMemberDetailResponse is the API response, containing a single
// account member.
type AccountMemberDetailResponse struct {
	Success  bool          `json:"success"`
	Errors   []string      `json:"errors"`
	Messages []string      `json:"messages"`
	Result   AccountMember `json:"result"`
}

// AccountMemberInvitation represents the invitation for a new member to
// the account.
type AccountMemberInvitation struct {
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}

// ListAccountMembers returns all members of an account.
//
// API reference: https://api.cloudflare.com/#accounts-list-accounts
func (api *API) ListAccountMembers(accountID string, pageOpts PaginationOptions) ([]AccountMember, ResultInfo, error) {
	v := url.Values{}
	if pageOpts.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(pageOpts.PerPage))
	}
	if pageOpts.Page > 0 {
		v.Set("page", strconv.Itoa(pageOpts.Page))
	}

	uri := "/accounts/" + accountID + "/members"
	if len(v) > 0 {
		uri = uri + "?" + v.Encode()
	}

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []AccountMember{}, ResultInfo{}, errors.Wrap(err, errMakeRequestError)
	}

	var accountMemberListresponse AccountMembersListResponse
	err = json.Unmarshal(res, &accountMemberListresponse)
	if err != nil {
		return []AccountMember{}, ResultInfo{}, errors.Wrap(err, errUnmarshalError)
	}

	return accountMemberListresponse.Result, accountMemberListresponse.ResultInfo, nil
}

// AddAccountMember invites a new member to join an account.
//
// API reference: https://api.cloudflare.com/#account-members-add-member
func (api *API) AddAccountMember(accountID string, emailAddress string, roles []string) (AccountMember, error) {
	uri := "/accounts/" + accountID + "/members"

	var newMember = AccountMemberInvitation{
		Email: emailAddress,
		Roles: roles,
	}
	res, err := api.makeRequest("POST", uri, newMember)
	if err != nil {
		return AccountMember{}, errors.Wrap(err, errMakeRequestError)
	}

	var accountMemberListResponse AccountMemberDetailResponse
	err = json.Unmarshal(res, &accountMemberListResponse)
	if err != nil {
		return AccountMember{}, errors.Wrap(err, errUnmarshalError)
	}

	return accountMemberListResponse.Result, nil
}
// UpdateAccountMember modifies an existing account member.
//
// API reference: https://api.cloudflare.com/#account-members-update-member
func (api *API) UpdateAccountMember(accountID string, userID string, member AccountMember) (AccountMember, error) {
	uri := fmt.Sprintf("/accounts/%s/members/%s", accountID, userID)

	res, err := api.makeRequest("PUT", uri, member)
	if err != nil {
		return AccountMember{}, errors.Wrap(err, errMakeRequestError)
	}

	var accountMemberListResponse AccountMemberDetailResponse
	err = json.Unmarshal(res, &accountMemberListResponse)
	if err != nil {
		return AccountMember{}, errors.Wrap(err, errUnmarshalError)
	}

	return accountMemberListResponse.Result, nil
}
