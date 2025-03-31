// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// RoleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoleService] method instead.
type RoleService struct {
	Options []option.RequestOption
}

// NewRoleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRoleService(opts ...option.RequestOption) (r *RoleService) {
	r = &RoleService{}
	r.Options = opts
	return
}

// Get all available roles for an account.
func (r *RoleService) List(ctx context.Context, params RoleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[shared.Role], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/roles", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get all available roles for an account.
func (r *RoleService) ListAutoPaging(ctx context.Context, params RoleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[shared.Role] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Get information about a specific role for an account.
func (r *RoleService) Get(ctx context.Context, roleID string, query RoleGetParams, opts ...option.RequestOption) (res *shared.Role, err error) {
	var env RoleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if roleID == "" {
		err = errors.New("missing required role_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/roles/%s", query.AccountID, roleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RoleListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of roles per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RoleListParams]'s query parameters as `url.Values`.
func (r RoleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RoleGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RoleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RoleGetResponseEnvelopeSuccess `json:"success,required"`
	Result  shared.Role                    `json:"result"`
	JSON    roleGetResponseEnvelopeJSON    `json:"-"`
}

// roleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoleGetResponseEnvelope]
type roleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoleGetResponseEnvelopeSuccess bool

const (
	RoleGetResponseEnvelopeSuccessTrue RoleGetResponseEnvelopeSuccess = true
)

func (r RoleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RoleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
