// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// RoleService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRoleService] method instead.
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
func (r *RoleService) List(ctx context.Context, query RoleListParams, opts ...option.RequestOption) (res *pagination.SinglePage[shared.Role], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/roles", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *RoleService) ListAutoPaging(ctx context.Context, query RoleListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[shared.Role] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Get information about a specific role for an account.
func (r *RoleService) Get(ctx context.Context, roleID interface{}, query RoleGetParams, opts ...option.RequestOption) (res *RoleGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env RoleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/roles/%v", query.AccountID, roleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [accounts.RoleGetResponseUnknown] or [shared.UnionString].
type RoleGetResponseUnion interface {
	ImplementsAccountsRoleGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RoleGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RoleListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RoleGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RoleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   RoleGetResponseUnion  `json:"result,required"`
	// Whether the API call was successful
	Success RoleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    roleGetResponseEnvelopeJSON    `json:"-"`
}

// roleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoleGetResponseEnvelope]
type roleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
