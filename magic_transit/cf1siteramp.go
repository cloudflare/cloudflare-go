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

// Cf1SiteRampService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCf1SiteRampService] method instead.
type Cf1SiteRampService struct {
	Options []option.RequestOption
}

// NewCf1SiteRampService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCf1SiteRampService(opts ...option.RequestOption) (r *Cf1SiteRampService) {
	r = &Cf1SiteRampService{}
	r.Options = opts
	return
}

// Creates ramps (network connections) for a CF1 Site.
func (r *Cf1SiteRampService) New(ctx context.Context, cf1SiteID string, params Cf1SiteRampNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[Ramp], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if cf1SiteID == "" {
		err = errors.New("missing required cf1_site_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/cf1_sites/%s/ramps", params.AccountID, cf1SiteID)
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

// Creates ramps (network connections) for a CF1 Site.
func (r *Cf1SiteRampService) NewAutoPaging(ctx context.Context, cf1SiteID string, params Cf1SiteRampNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Ramp] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, cf1SiteID, params, opts...))
}

// Lists ramps (network connections) associated with a CF1 Site. Ramps represent
// GRE tunnels, IPsec tunnels, interconnects, or MCONN links.
func (r *Cf1SiteRampService) List(ctx context.Context, cf1SiteID string, query Cf1SiteRampListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Ramp], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if cf1SiteID == "" {
		err = errors.New("missing required cf1_site_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/cf1_sites/%s/ramps", query.AccountID, cf1SiteID)
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

// Lists ramps (network connections) associated with a CF1 Site. Ramps represent
// GRE tunnels, IPsec tunnels, interconnects, or MCONN links.
func (r *Cf1SiteRampService) ListAutoPaging(ctx context.Context, cf1SiteID string, query Cf1SiteRampListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Ramp] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, cf1SiteID, query, opts...))
}

// Deletes a specific ramp from a CF1 Site.
func (r *Cf1SiteRampService) Delete(ctx context.Context, cf1SiteID string, rampID string, body Cf1SiteRampDeleteParams, opts ...option.RequestOption) (res *Ramp, err error) {
	var env Cf1SiteRampDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if cf1SiteID == "" {
		err = errors.New("missing required cf1_site_id parameter")
		return nil, err
	}
	if rampID == "" {
		err = errors.New("missing required ramp_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/cf1_sites/%s/ramps/%s", body.AccountID, cf1SiteID, rampID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets a specific ramp for a CF1 Site.
func (r *Cf1SiteRampService) Get(ctx context.Context, cf1SiteID string, rampID string, query Cf1SiteRampGetParams, opts ...option.RequestOption) (res *Ramp, err error) {
	var env Cf1SiteRampGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if cf1SiteID == "" {
		err = errors.New("missing required cf1_site_id parameter")
		return nil, err
	}
	if rampID == "" {
		err = errors.New("missing required ramp_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/cf1_sites/%s/ramps/%s", query.AccountID, cf1SiteID, rampID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type Ramp struct {
	// Identifier
	ID         string    `json:"id" api:"required"`
	CreatedOn  time.Time `json:"created_on" api:"required" format:"date-time"`
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// A human-provided name describing the ramp that should be unique within the CF1
	// Site.
	Name string `json:"name" api:"required"`
	// The type of network connection (ramp) linking a CF1 Site to Cloudflare's
	// network.
	Type RampType `json:"type" api:"required"`
	// A human-provided description of the ramp.
	Description      string               `json:"description"`
	GRE              RampGRE              `json:"gre"`
	GREInterconnect  RampGREInterconnect  `json:"gre_interconnect"`
	IPSEC            RampIPSEC            `json:"ipsec"`
	Mconn            RampMconn            `json:"mconn"`
	MplsInterconnect RampMplsInterconnect `json:"mpls_interconnect"`
	JSON             rampJSON             `json:"-"`
}

// rampJSON contains the JSON metadata for the struct [Ramp]
type rampJSON struct {
	ID               apijson.Field
	CreatedOn        apijson.Field
	ModifiedOn       apijson.Field
	Name             apijson.Field
	Type             apijson.Field
	Description      apijson.Field
	GRE              apijson.Field
	GREInterconnect  apijson.Field
	IPSEC            apijson.Field
	Mconn            apijson.Field
	MplsInterconnect apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Ramp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rampJSON) RawJSON() string {
	return r.raw
}

type RampGRE struct {
	// URL reference to the source network resource that this ramp is managed by.
	ManagedBy string      `json:"managed_by"`
	JSON      rampGREJSON `json:"-"`
}

// rampGREJSON contains the JSON metadata for the struct [RampGRE]
type rampGREJSON struct {
	ManagedBy   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RampGRE) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rampGREJSON) RawJSON() string {
	return r.raw
}

type RampGREInterconnect struct {
	// URL reference to the source network resource that this ramp is managed by.
	ManagedBy string                  `json:"managed_by"`
	JSON      rampGREInterconnectJSON `json:"-"`
}

// rampGREInterconnectJSON contains the JSON metadata for the struct
// [RampGREInterconnect]
type rampGREInterconnectJSON struct {
	ManagedBy   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RampGREInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rampGREInterconnectJSON) RawJSON() string {
	return r.raw
}

type RampIPSEC struct {
	// URL reference to the source network resource that this ramp is managed by.
	ManagedBy string        `json:"managed_by"`
	JSON      rampIPSECJSON `json:"-"`
}

// rampIPSECJSON contains the JSON metadata for the struct [RampIPSEC]
type rampIPSECJSON struct {
	ManagedBy   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RampIPSEC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rampIPSECJSON) RawJSON() string {
	return r.raw
}

type RampMconn struct {
	// URL reference to the source network resource that this ramp is managed by.
	ManagedBy string        `json:"managed_by"`
	JSON      rampMconnJSON `json:"-"`
}

// rampMconnJSON contains the JSON metadata for the struct [RampMconn]
type rampMconnJSON struct {
	ManagedBy   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RampMconn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rampMconnJSON) RawJSON() string {
	return r.raw
}

type RampMplsInterconnect struct {
	// URL reference to the source network resource that this ramp is managed by.
	ManagedBy string                   `json:"managed_by"`
	JSON      rampMplsInterconnectJSON `json:"-"`
}

// rampMplsInterconnectJSON contains the JSON metadata for the struct
// [RampMplsInterconnect]
type rampMplsInterconnectJSON struct {
	ManagedBy   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RampMplsInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rampMplsInterconnectJSON) RawJSON() string {
	return r.raw
}

// The type of network connection (ramp) linking a CF1 Site to Cloudflare's
// network.
type RampType string

const (
	RampTypeGRE              RampType = "gre"
	RampTypeGREInterconnect  RampType = "gre_interconnect"
	RampTypeMplsInterconnect RampType = "mpls_interconnect"
	RampTypeMconn            RampType = "mconn"
	RampTypeIPSEC            RampType = "ipsec"
)

func (r RampType) IsKnown() bool {
	switch r {
	case RampTypeGRE, RampTypeGREInterconnect, RampTypeMplsInterconnect, RampTypeMconn, RampTypeIPSEC:
		return true
	}
	return false
}

type Cf1SiteRampNewParams struct {
	// Identifier
	AccountID param.Field[string]        `path:"account_id" api:"required"`
	Body      []Cf1SiteRampNewParamsBody `json:"body" api:"required"`
}

func (r Cf1SiteRampNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type Cf1SiteRampNewParamsBody struct {
	// Identifier of the source network resource to associate as a ramp.
	SourceRampID param.Field[string] `json:"source_ramp_id" api:"required"`
	// The type of network connection (ramp) linking a CF1 Site to Cloudflare's
	// network.
	Type param.Field[RampType] `json:"type" api:"required"`
}

func (r Cf1SiteRampNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Cf1SiteRampListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type Cf1SiteRampDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type Cf1SiteRampDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   Ramp                  `json:"result" api:"required"`
	// Whether the API call was successful
	Success Cf1SiteRampDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    cf1SiteRampDeleteResponseEnvelopeJSON    `json:"-"`
}

// cf1SiteRampDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [Cf1SiteRampDeleteResponseEnvelope]
type cf1SiteRampDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Cf1SiteRampDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cf1SiteRampDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type Cf1SiteRampDeleteResponseEnvelopeSuccess bool

const (
	Cf1SiteRampDeleteResponseEnvelopeSuccessTrue Cf1SiteRampDeleteResponseEnvelopeSuccess = true
)

func (r Cf1SiteRampDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case Cf1SiteRampDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type Cf1SiteRampGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type Cf1SiteRampGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   Ramp                  `json:"result" api:"required"`
	// Whether the API call was successful
	Success Cf1SiteRampGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    cf1SiteRampGetResponseEnvelopeJSON    `json:"-"`
}

// cf1SiteRampGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [Cf1SiteRampGetResponseEnvelope]
type cf1SiteRampGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Cf1SiteRampGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cf1SiteRampGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type Cf1SiteRampGetResponseEnvelopeSuccess bool

const (
	Cf1SiteRampGetResponseEnvelopeSuccessTrue Cf1SiteRampGetResponseEnvelopeSuccess = true
)

func (r Cf1SiteRampGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case Cf1SiteRampGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
