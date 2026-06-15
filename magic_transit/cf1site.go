// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// Cf1SiteService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCf1SiteService] method instead.
type Cf1SiteService struct {
	Options []option.RequestOption
	Ramps   *Cf1SiteRampService
}

// NewCf1SiteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCf1SiteService(opts ...option.RequestOption) (r *Cf1SiteService) {
	r = &Cf1SiteService{}
	r.Options = opts
	r.Ramps = NewCf1SiteRampService(opts...)
	return
}

// Creates new CF1 Sites for an account. Each site must have a unique name within
// the account.
func (r *Cf1SiteService) New(ctx context.Context, params Cf1SiteNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[Cf1Site], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/cf1_sites", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Creates new CF1 Sites for an account. Each site must have a unique name within
// the account.
func (r *Cf1SiteService) NewAutoPaging(ctx context.Context, params Cf1SiteNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Cf1Site] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, params, opts...))
}

// Partially updates a specific CF1 Site for an account. Only the fields included
// in the request body are modified; omitted fields retain their existing values.
func (r *Cf1SiteService) Update(ctx context.Context, cf1SiteID string, params Cf1SiteUpdateParams, opts ...option.RequestOption) (res *Cf1Site, err error) {
	var env Cf1SiteUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if cf1SiteID == "" {
		err = errors.New("missing required cf1_site_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/cf1_sites/%s", params.AccountID, cf1SiteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists CF1 Sites associated with an account. A CF1 Site represents a physical
// customer network location with optional geographic coordinates.
func (r *Cf1SiteService) List(ctx context.Context, query Cf1SiteListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Cf1Site], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/cf1_sites", query.AccountID)
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

// Lists CF1 Sites associated with an account. A CF1 Site represents a physical
// customer network location with optional geographic coordinates.
func (r *Cf1SiteService) ListAutoPaging(ctx context.Context, query Cf1SiteListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Cf1Site] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a specific CF1 Site for an account.
func (r *Cf1SiteService) Delete(ctx context.Context, cf1SiteID string, body Cf1SiteDeleteParams, opts ...option.RequestOption) (res *Cf1Site, err error) {
	var env Cf1SiteDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if cf1SiteID == "" {
		err = errors.New("missing required cf1_site_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/cf1_sites/%s", body.AccountID, cf1SiteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets a specific CF1 Site for an account.
func (r *Cf1SiteService) Get(ctx context.Context, cf1SiteID string, query Cf1SiteGetParams, opts ...option.RequestOption) (res *Cf1Site, err error) {
	var env Cf1SiteGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if cf1SiteID == "" {
		err = errors.New("missing required cf1_site_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/cf1_sites/%s", query.AccountID, cf1SiteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type Cf1Site struct {
	// A human-provided name describing the CF1 Site that should be unique within the
	// account.
	Name string `json:"name" api:"required"`
	// Identifier
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// A human-provided description of the CF1 Site.
	Description string          `json:"description"`
	Location    Cf1SiteLocation `json:"location"`
	ModifiedOn  time.Time       `json:"modified_on" format:"date-time"`
	JSON        cf1SiteJSON     `json:"-"`
}

// cf1SiteJSON contains the JSON metadata for the struct [Cf1Site]
type cf1SiteJSON struct {
	Name        apijson.Field
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	Location    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Cf1Site) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cf1SiteJSON) RawJSON() string {
	return r.raw
}

type Cf1SiteParam struct {
	// A human-provided name describing the CF1 Site that should be unique within the
	// account.
	Name param.Field[string] `json:"name" api:"required"`
	// A human-provided description of the CF1 Site.
	Description param.Field[string]               `json:"description"`
	Location    param.Field[Cf1SiteLocationParam] `json:"location"`
}

func (r Cf1SiteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Cf1SiteLocation struct {
	// Latitude of the CF1 Site.
	Lat float64 `json:"lat"`
	// Longitude of the CF1 Site.
	Long float64 `json:"long"`
	// Name of nearest town, city, or village.
	Name string              `json:"name"`
	JSON cf1SiteLocationJSON `json:"-"`
}

// cf1SiteLocationJSON contains the JSON metadata for the struct [Cf1SiteLocation]
type cf1SiteLocationJSON struct {
	Lat         apijson.Field
	Long        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Cf1SiteLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cf1SiteLocationJSON) RawJSON() string {
	return r.raw
}

type Cf1SiteLocationParam struct {
	// Latitude of the CF1 Site.
	Lat param.Field[float64] `json:"lat"`
	// Longitude of the CF1 Site.
	Long param.Field[float64] `json:"long"`
	// Name of nearest town, city, or village.
	Name param.Field[string] `json:"name"`
}

func (r Cf1SiteLocationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Cf1SiteNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Body      []Cf1SiteParam      `json:"body" api:"required"`
}

func (r Cf1SiteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type Cf1SiteUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A human-provided description of the CF1 Site.
	Description param.Field[string]               `json:"description"`
	Location    param.Field[Cf1SiteLocationParam] `json:"location"`
	// A human-provided name describing the CF1 Site that should be unique within the
	// account.
	Name param.Field[string] `json:"name"`
}

func (r Cf1SiteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Cf1SiteUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   Cf1Site               `json:"result" api:"required"`
	// Whether the API call was successful
	Success Cf1SiteUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    cf1SiteUpdateResponseEnvelopeJSON    `json:"-"`
}

// cf1SiteUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [Cf1SiteUpdateResponseEnvelope]
type cf1SiteUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Cf1SiteUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cf1SiteUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type Cf1SiteUpdateResponseEnvelopeSuccess bool

const (
	Cf1SiteUpdateResponseEnvelopeSuccessTrue Cf1SiteUpdateResponseEnvelopeSuccess = true
)

func (r Cf1SiteUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case Cf1SiteUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type Cf1SiteListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type Cf1SiteDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type Cf1SiteDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   Cf1Site               `json:"result" api:"required"`
	// Whether the API call was successful
	Success Cf1SiteDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    cf1SiteDeleteResponseEnvelopeJSON    `json:"-"`
}

// cf1SiteDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [Cf1SiteDeleteResponseEnvelope]
type cf1SiteDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Cf1SiteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cf1SiteDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type Cf1SiteDeleteResponseEnvelopeSuccess bool

const (
	Cf1SiteDeleteResponseEnvelopeSuccessTrue Cf1SiteDeleteResponseEnvelopeSuccess = true
)

func (r Cf1SiteDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case Cf1SiteDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type Cf1SiteGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type Cf1SiteGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   Cf1Site               `json:"result" api:"required"`
	// Whether the API call was successful
	Success Cf1SiteGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    cf1SiteGetResponseEnvelopeJSON    `json:"-"`
}

// cf1SiteGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [Cf1SiteGetResponseEnvelope]
type cf1SiteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Cf1SiteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cf1SiteGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type Cf1SiteGetResponseEnvelopeSuccess bool

const (
	Cf1SiteGetResponseEnvelopeSuccessTrue Cf1SiteGetResponseEnvelopeSuccess = true
)

func (r Cf1SiteGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case Cf1SiteGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
