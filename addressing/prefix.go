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

// PrefixService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrefixService] method instead.
type PrefixService struct {
	Options     []option.RequestOption
	BGP         *PrefixBGPService
	Delegations *PrefixDelegationService
}

// NewPrefixService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPrefixService(opts ...option.RequestOption) (r *PrefixService) {
	r = &PrefixService{}
	r.Options = opts
	r.BGP = NewPrefixBGPService(opts...)
	r.Delegations = NewPrefixDelegationService(opts...)
	return
}

// Add a new prefix under the account.
func (r *PrefixService) New(ctx context.Context, params PrefixNewParams, opts ...option.RequestOption) (res *Prefix, err error) {
	var env PrefixNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all prefixes owned by the account.
func (r *PrefixService) List(ctx context.Context, query PrefixListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Prefix], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", query.AccountID)
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

// List all prefixes owned by the account.
func (r *PrefixService) ListAutoPaging(ctx context.Context, query PrefixListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Prefix] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete an unapproved prefix owned by the account.
func (r *PrefixService) Delete(ctx context.Context, prefixID string, body PrefixDeleteParams, opts ...option.RequestOption) (res *PrefixDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", body.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Modify the description for a prefix owned by the account.
func (r *PrefixService) Edit(ctx context.Context, prefixID string, params PrefixEditParams, opts ...option.RequestOption) (res *Prefix, err error) {
	var env PrefixEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List a particular prefix owned by the account.
func (r *PrefixService) Get(ctx context.Context, prefixID string, query PrefixGetParams, opts ...option.RequestOption) (res *Prefix, err error) {
	var env PrefixGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", query.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Prefix struct {
	// Identifier
	ID string `json:"id"`
	// Identifier
	AccountID string `json:"account_id"`
	// Prefix advertisement status to the Internet. This field is only not 'null' if on
	// demand is enabled.
	Advertised bool `json:"advertised,nullable"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time `json:"advertised_modified_at,nullable" format:"date-time"`
	// Approval state of the prefix (P = pending, V = active).
	Approved string `json:"approved"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN int64 `json:"asn,nullable"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Description of the prefix.
	Description string `json:"description"`
	// Identifier for the uploaded LOA document.
	LOADocumentID string    `json:"loa_document_id,nullable"`
	ModifiedAt    time.Time `json:"modified_at" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool       `json:"on_demand_locked"`
	JSON           prefixJSON `json:"-"`
}

// prefixJSON contains the JSON metadata for the struct [Prefix]
type prefixJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	ASN                  apijson.Field
	CIDR                 apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	LOADocumentID        apijson.Field
	ModifiedAt           apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Prefix) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixJSON) RawJSON() string {
	return r.raw
}

type PrefixDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    PrefixDeleteResponseSuccess    `json:"success,required"`
	ResultInfo PrefixDeleteResponseResultInfo `json:"result_info"`
	JSON       prefixDeleteResponseJSON       `json:"-"`
}

// prefixDeleteResponseJSON contains the JSON metadata for the struct
// [PrefixDeleteResponse]
type prefixDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixDeleteResponseSuccess bool

const (
	PrefixDeleteResponseSuccessTrue PrefixDeleteResponseSuccess = true
)

func (r PrefixDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case PrefixDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type PrefixDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                            `json:"total_count"`
	JSON       prefixDeleteResponseResultInfoJSON `json:"-"`
}

// prefixDeleteResponseResultInfoJSON contains the JSON metadata for the struct
// [PrefixDeleteResponseResultInfo]
type prefixDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixDeleteResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type PrefixNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN param.Field[int64] `json:"asn,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR param.Field[string] `json:"cidr,required"`
	// Identifier for the uploaded LOA document.
	LOADocumentID param.Field[string] `json:"loa_document_id,required"`
}

func (r PrefixNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Prefix                           `json:"result"`
	JSON    prefixNewResponseEnvelopeJSON    `json:"-"`
}

// prefixNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PrefixNewResponseEnvelope]
type prefixNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixNewResponseEnvelopeSuccess bool

const (
	PrefixNewResponseEnvelopeSuccessTrue PrefixNewResponseEnvelopeSuccess = true
)

func (r PrefixNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrefixListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Description of the prefix.
	Description param.Field[string] `json:"description,required"`
}

func (r PrefixEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixEditResponseEnvelopeSuccess `json:"success,required"`
	Result  Prefix                            `json:"result"`
	JSON    prefixEditResponseEnvelopeJSON    `json:"-"`
}

// prefixEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [PrefixEditResponseEnvelope]
type prefixEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixEditResponseEnvelopeSuccess bool

const (
	PrefixEditResponseEnvelopeSuccessTrue PrefixEditResponseEnvelopeSuccess = true
)

func (r PrefixEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrefixGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Prefix                           `json:"result"`
	JSON    prefixGetResponseEnvelopeJSON    `json:"-"`
}

// prefixGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PrefixGetResponseEnvelope]
type prefixGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixGetResponseEnvelopeSuccess bool

const (
	PrefixGetResponseEnvelopeSuccessTrue PrefixGetResponseEnvelopeSuccess = true
)

func (r PrefixGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
