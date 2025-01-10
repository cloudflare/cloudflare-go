// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// AccessApplicationPolicyTestService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessApplicationPolicyTestService] method instead.
type AccessApplicationPolicyTestService struct {
	Options []option.RequestOption
	Users   *AccessApplicationPolicyTestUserService
}

// NewAccessApplicationPolicyTestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessApplicationPolicyTestService(opts ...option.RequestOption) (r *AccessApplicationPolicyTestService) {
	r = &AccessApplicationPolicyTestService{}
	r.Options = opts
	r.Users = NewAccessApplicationPolicyTestUserService(opts...)
	return
}

// Starts an Access policy test.
func (r *AccessApplicationPolicyTestService) New(ctx context.Context, params AccessApplicationPolicyTestNewParams, opts ...option.RequestOption) (res *AccessApplicationPolicyTestNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policy-tests", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Fetches the current status of a given Access policy test.
func (r *AccessApplicationPolicyTestService) Get(ctx context.Context, policyTestID string, query AccessApplicationPolicyTestGetParams, opts ...option.RequestOption) (res *AccessApplicationPolicyTestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyTestID == "" {
		err = errors.New("missing required policy_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policy-tests/%s", query.AccountID, policyTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessApplicationPolicyTestNewResponse struct {
	// The UUID of the policy test.
	ID string `json:"id"`
	// The status of the policy test request.
	Status AccessApplicationPolicyTestNewResponseStatus `json:"status"`
	JSON   accessApplicationPolicyTestNewResponseJSON   `json:"-"`
}

// accessApplicationPolicyTestNewResponseJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyTestNewResponse]
type accessApplicationPolicyTestNewResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestNewResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the policy test request.
type AccessApplicationPolicyTestNewResponseStatus string

const (
	AccessApplicationPolicyTestNewResponseStatusSuccess AccessApplicationPolicyTestNewResponseStatus = "success"
)

func (r AccessApplicationPolicyTestNewResponseStatus) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyTestNewResponseStatusSuccess:
		return true
	}
	return false
}

type AccessApplicationPolicyTestGetResponse struct {
	// The UUID of the policy test.
	ID string `json:"id"`
	// The number of pages of (processed) users.
	PagesProcessed int64 `json:"pages_processed"`
	// The percentage of (processed) users approved based on policy evaluation results.
	PercentApproved int64 `json:"percent_approved"`
	// The percentage of (processed) users blocked based on policy evaluation results.
	PercentBlocked int64 `json:"percent_blocked"`
	// The percentage of users processed so far (of the entire user base).
	PercentUsersProcessed int64 `json:"percent_users_processed"`
	// The status of the policy test.
	Status AccessApplicationPolicyTestGetResponseStatus `json:"status"`
	// The total number of users in the user base.
	TotalUsers int64 `json:"total_users"`
	// The number of (processed) users approved based on policy evaluation results.
	UsersApproved int64 `json:"users_approved"`
	// The number of (processed) users blocked based on policy evaluation results.
	UsersBlocked int64                                      `json:"users_blocked"`
	JSON         accessApplicationPolicyTestGetResponseJSON `json:"-"`
}

// accessApplicationPolicyTestGetResponseJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyTestGetResponse]
type accessApplicationPolicyTestGetResponseJSON struct {
	ID                    apijson.Field
	PagesProcessed        apijson.Field
	PercentApproved       apijson.Field
	PercentBlocked        apijson.Field
	PercentUsersProcessed apijson.Field
	Status                apijson.Field
	TotalUsers            apijson.Field
	UsersApproved         apijson.Field
	UsersBlocked          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestGetResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the policy test.
type AccessApplicationPolicyTestGetResponseStatus string

const (
	AccessApplicationPolicyTestGetResponseStatusBlocked    AccessApplicationPolicyTestGetResponseStatus = "blocked"
	AccessApplicationPolicyTestGetResponseStatusProcessing AccessApplicationPolicyTestGetResponseStatus = "processing"
	AccessApplicationPolicyTestGetResponseStatusComplete   AccessApplicationPolicyTestGetResponseStatus = "complete"
)

func (r AccessApplicationPolicyTestGetResponseStatus) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyTestGetResponseStatusBlocked, AccessApplicationPolicyTestGetResponseStatusProcessing, AccessApplicationPolicyTestGetResponseStatusComplete:
		return true
	}
	return false
}

type AccessApplicationPolicyTestNewParams struct {
	// Identifier
	AccountID         param.Field[string]    `path:"account_id,required"`
	ApplicationPolicy ApplicationPolicyParam `json:"application_policy,required"`
}

func (r AccessApplicationPolicyTestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ApplicationPolicy)
}

type AccessApplicationPolicyTestGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
