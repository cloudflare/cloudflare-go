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

// AccessApplicationPolicyTestUserService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessApplicationPolicyTestUserService] method instead.
type AccessApplicationPolicyTestUserService struct {
	Options []option.RequestOption
}

// NewAccessApplicationPolicyTestUserService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessApplicationPolicyTestUserService(opts ...option.RequestOption) (r *AccessApplicationPolicyTestUserService) {
	r = &AccessApplicationPolicyTestUserService{}
	r.Options = opts
	return
}

// Fetches a single page of user results from an Access policy test.
func (r *AccessApplicationPolicyTestUserService) List(ctx context.Context, policyTestID string, query AccessApplicationPolicyTestUserListParams, opts ...option.RequestOption) (res *[]AccessApplicationPolicyTestUserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyTestID == "" {
		err = errors.New("missing required policy_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policy-tests/%s/users", query.AccountID, policyTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessApplicationPolicyTestUserListResponse struct {
	// UUID
	ID string `json:"id"`
	// The email of the user.
	Email string `json:"email" format:"email"`
	// The name of the user.
	Name string `json:"name"`
	// Policy evaluation result for an individual user.
	Status AccessApplicationPolicyTestUserListResponseStatus `json:"status"`
	JSON   accessApplicationPolicyTestUserListResponseJSON   `json:"-"`
}

// accessApplicationPolicyTestUserListResponseJSON contains the JSON metadata for
// the struct [AccessApplicationPolicyTestUserListResponse]
type accessApplicationPolicyTestUserListResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestUserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestUserListResponseJSON) RawJSON() string {
	return r.raw
}

// Policy evaluation result for an individual user.
type AccessApplicationPolicyTestUserListResponseStatus string

const (
	AccessApplicationPolicyTestUserListResponseStatusApproved AccessApplicationPolicyTestUserListResponseStatus = "approved"
	AccessApplicationPolicyTestUserListResponseStatusBlocked  AccessApplicationPolicyTestUserListResponseStatus = "blocked"
)

func (r AccessApplicationPolicyTestUserListResponseStatus) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyTestUserListResponseStatusApproved, AccessApplicationPolicyTestUserListResponseStatusBlocked:
		return true
	}
	return false
}

type AccessApplicationPolicyTestUserListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
