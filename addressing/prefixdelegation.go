// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// PrefixDelegationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrefixDelegationService] method instead.
type PrefixDelegationService struct {
	Options []option.RequestOption
}

// NewPrefixDelegationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrefixDelegationService(opts ...option.RequestOption) (r *PrefixDelegationService) {
	r = &PrefixDelegationService{}
	r.Options = opts
	return
}

// Create a new account delegation for a given IP prefix.
func (r *PrefixDelegationService) New(ctx context.Context, prefixID string, params PrefixDelegationNewParams, opts ...option.RequestOption) (res *Delegations, err error) {
	var env PrefixDelegationNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all delegations for a given account IP prefix.
func (r *PrefixDelegationService) List(ctx context.Context, prefixID string, query PrefixDelegationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Delegations], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", query.AccountID, prefixID)
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

// List all delegations for a given account IP prefix.
func (r *PrefixDelegationService) ListAutoPaging(ctx context.Context, prefixID string, query PrefixDelegationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Delegations] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, prefixID, query, opts...))
}

// Delete an account delegation for a given IP prefix.
func (r *PrefixDelegationService) Delete(ctx context.Context, prefixID string, delegationID string, body PrefixDelegationDeleteParams, opts ...option.RequestOption) (res *PrefixDelegationDeleteResponse, err error) {
	var env PrefixDelegationDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	if delegationID == "" {
		err = errors.New("missing required delegation_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations/%s", body.AccountID, prefixID, delegationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Delegations struct {
	// Delegation identifier tag.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier
	ParentPrefixID string          `json:"parent_prefix_id"`
	JSON           delegationsJSON `json:"-"`
}

// delegationsJSON contains the JSON metadata for the struct [Delegations]
type delegationsJSON struct {
	ID                 apijson.Field
	CIDR               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Delegations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r delegationsJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationDeleteResponse struct {
	// Delegation identifier tag.
	ID   string                             `json:"id"`
	JSON prefixDelegationDeleteResponseJSON `json:"-"`
}

// prefixDelegationDeleteResponseJSON contains the JSON metadata for the struct
// [PrefixDelegationDeleteResponse]
type prefixDelegationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixDelegationNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR param.Field[string] `json:"cidr,required"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID param.Field[string] `json:"delegated_account_id,required"`
}

func (r PrefixDelegationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixDelegationNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixDelegationNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Delegations                                `json:"result"`
	JSON    prefixDelegationNewResponseEnvelopeJSON    `json:"-"`
}

// prefixDelegationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixDelegationNewResponseEnvelope]
type prefixDelegationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixDelegationNewResponseEnvelopeSuccess bool

const (
	PrefixDelegationNewResponseEnvelopeSuccessTrue PrefixDelegationNewResponseEnvelopeSuccess = true
)

func (r PrefixDelegationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixDelegationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrefixDelegationListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixDelegationDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixDelegationDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixDelegationDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  PrefixDelegationDeleteResponse                `json:"result"`
	JSON    prefixDelegationDeleteResponseEnvelopeJSON    `json:"-"`
}

// prefixDelegationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixDelegationDeleteResponseEnvelope]
type prefixDelegationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDelegationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDelegationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixDelegationDeleteResponseEnvelopeSuccess bool

const (
	PrefixDelegationDeleteResponseEnvelopeSuccessTrue PrefixDelegationDeleteResponseEnvelopeSuccess = true
)

func (r PrefixDelegationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixDelegationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
