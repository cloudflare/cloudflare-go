// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_tls_compliance_modes

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
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// OriginTLSComplianceModeService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOriginTLSComplianceModeService] method instead.
type OriginTLSComplianceModeService struct {
	Options []option.RequestOption
}

// NewOriginTLSComplianceModeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOriginTLSComplianceModeService(opts ...option.RequestOption) (r *OriginTLSComplianceModeService) {
	r = &OriginTLSComplianceModeService{}
	r.Options = opts
	return
}

// Replace the entire set of TLS compliance modes for the zone with the list
// provided in the request body. PUT performs a full replace, not a merge — any
// modes not present in the request body are removed. The request body must be of
// the form `{"value": ["fips", "pqh"]}`. Currently supported modes are `fips` and
// `pqh`; an empty list clears the constraint. Future modes (e.g. `cnsa2`) may be
// added; clients should treat unknown values as opaque strings. Invalid mode
// values are rejected with a 4xx response.
func (r *OriginTLSComplianceModeService) Update(ctx context.Context, params OriginTLSComplianceModeUpdateParams, opts ...option.RequestOption) (res *OriginTLSComplianceModeUpdateResponse, err error) {
	var env OriginTLSComplianceModeUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/origin_tls_compliance_modes", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Delete the Origin TLS Compliance Modes setting for the zone, removing any
// configured compliance constraint. After deletion, Cloudflare's default behavior
// applies (no compliance filtering of the key-exchange algorithm list sent to the
// origin).
func (r *OriginTLSComplianceModeService) Delete(ctx context.Context, body OriginTLSComplianceModeDeleteParams, opts ...option.RequestOption) (res *OriginTLSComplianceModeDeleteResponse, err error) {
	var env OriginTLSComplianceModeDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/origin_tls_compliance_modes", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update the set of TLS compliance modes for the zone. PATCH performs a full
// replace of the modes list, not a merge — the request body is treated as the
// complete new list, and any modes not present in it are removed. (To remove a
// single mode from an existing configuration, send the updated list without it.)
// The request body must be of the form `{"value": ["fips", "pqh"]}`. Currently
// supported modes are `fips` and `pqh`; an empty list clears the constraint.
// Future modes (e.g. `cnsa2`) may be added; clients should treat unknown values as
// opaque strings. Invalid mode values are rejected with a 4xx response.
func (r *OriginTLSComplianceModeService) Edit(ctx context.Context, params OriginTLSComplianceModeEditParams, opts ...option.RequestOption) (res *OriginTLSComplianceModeEditResponse, err error) {
	var env OriginTLSComplianceModeEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/origin_tls_compliance_modes", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Origin TLS Compliance Modes constrains the set of TLS key-exchange algorithms
// Cloudflare may use when establishing the TLS connection to the zone's origin.
// The value is a list of named compliance modes (currently `fips` and `pqh`).
// Multiple modes are combined as the intersection of their permitted algorithm
// lists. An empty list (or no rule configured) means no compliance constraint is
// applied.
func (r *OriginTLSComplianceModeService) Get(ctx context.Context, query OriginTLSComplianceModeGetParams, opts ...option.RequestOption) (res *OriginTLSComplianceModeGetResponse, err error) {
	var env OriginTLSComplianceModeGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/origin_tls_compliance_modes", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type OriginTLSComplianceModeUpdateResponse struct {
	// The identifier of the caching setting.
	ID OriginTLSComplianceModeUpdateResponseID `json:"id" api:"required"`
	// Whether the setting is editable.
	Editable bool `json:"editable" api:"required"`
	// List of TLS compliance modes that constrain the key-exchange algorithms
	// Cloudflare may use when establishing the TLS connection to the zone's origin.
	// Currently supported values are `fips` (FIPS-approved curves) and `pqh`
	// (post-quantum hybrid). Future modes (e.g. `cnsa2`) may be added; clients should
	// treat unknown values as opaque strings. Multiple modes are combined as the
	// intersection of their permitted algorithm lists; selections whose intersection
	// is empty are rejected. An empty list clears the constraint.
	Value []string `json:"value" api:"required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       originTLSComplianceModeUpdateResponseJSON `json:"-"`
}

// originTLSComplianceModeUpdateResponseJSON contains the JSON metadata for the
// struct [OriginTLSComplianceModeUpdateResponse]
type originTLSComplianceModeUpdateResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSComplianceModeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSComplianceModeUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The identifier of the caching setting.
type OriginTLSComplianceModeUpdateResponseID string

const (
	OriginTLSComplianceModeUpdateResponseIDOriginTLSComplianceModes OriginTLSComplianceModeUpdateResponseID = "origin_tls_compliance_modes"
)

func (r OriginTLSComplianceModeUpdateResponseID) IsKnown() bool {
	switch r {
	case OriginTLSComplianceModeUpdateResponseIDOriginTLSComplianceModes:
		return true
	}
	return false
}

type OriginTLSComplianceModeDeleteResponse struct {
	// The identifier of the caching setting.
	ID OriginTLSComplianceModeDeleteResponseID `json:"id" api:"required"`
	// Whether the setting is editable.
	Editable bool `json:"editable" api:"required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       originTLSComplianceModeDeleteResponseJSON `json:"-"`
}

// originTLSComplianceModeDeleteResponseJSON contains the JSON metadata for the
// struct [OriginTLSComplianceModeDeleteResponse]
type originTLSComplianceModeDeleteResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSComplianceModeDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSComplianceModeDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The identifier of the caching setting.
type OriginTLSComplianceModeDeleteResponseID string

const (
	OriginTLSComplianceModeDeleteResponseIDOriginTLSComplianceModes OriginTLSComplianceModeDeleteResponseID = "origin_tls_compliance_modes"
)

func (r OriginTLSComplianceModeDeleteResponseID) IsKnown() bool {
	switch r {
	case OriginTLSComplianceModeDeleteResponseIDOriginTLSComplianceModes:
		return true
	}
	return false
}

type OriginTLSComplianceModeEditResponse struct {
	// The identifier of the caching setting.
	ID OriginTLSComplianceModeEditResponseID `json:"id" api:"required"`
	// Whether the setting is editable.
	Editable bool `json:"editable" api:"required"`
	// List of TLS compliance modes that constrain the key-exchange algorithms
	// Cloudflare may use when establishing the TLS connection to the zone's origin.
	// Currently supported values are `fips` (FIPS-approved curves) and `pqh`
	// (post-quantum hybrid). Future modes (e.g. `cnsa2`) may be added; clients should
	// treat unknown values as opaque strings. Multiple modes are combined as the
	// intersection of their permitted algorithm lists; selections whose intersection
	// is empty are rejected. An empty list clears the constraint.
	Value []string `json:"value" api:"required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       originTLSComplianceModeEditResponseJSON `json:"-"`
}

// originTLSComplianceModeEditResponseJSON contains the JSON metadata for the
// struct [OriginTLSComplianceModeEditResponse]
type originTLSComplianceModeEditResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSComplianceModeEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSComplianceModeEditResponseJSON) RawJSON() string {
	return r.raw
}

// The identifier of the caching setting.
type OriginTLSComplianceModeEditResponseID string

const (
	OriginTLSComplianceModeEditResponseIDOriginTLSComplianceModes OriginTLSComplianceModeEditResponseID = "origin_tls_compliance_modes"
)

func (r OriginTLSComplianceModeEditResponseID) IsKnown() bool {
	switch r {
	case OriginTLSComplianceModeEditResponseIDOriginTLSComplianceModes:
		return true
	}
	return false
}

type OriginTLSComplianceModeGetResponse struct {
	// The identifier of the caching setting.
	ID OriginTLSComplianceModeGetResponseID `json:"id" api:"required"`
	// Whether the setting is editable.
	Editable bool `json:"editable" api:"required"`
	// List of TLS compliance modes that constrain the key-exchange algorithms
	// Cloudflare may use when establishing the TLS connection to the zone's origin.
	// Currently supported values are `fips` (FIPS-approved curves) and `pqh`
	// (post-quantum hybrid). Future modes (e.g. `cnsa2`) may be added; clients should
	// treat unknown values as opaque strings. Multiple modes are combined as the
	// intersection of their permitted algorithm lists; selections whose intersection
	// is empty are rejected. An empty list clears the constraint.
	Value []string `json:"value" api:"required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       originTLSComplianceModeGetResponseJSON `json:"-"`
}

// originTLSComplianceModeGetResponseJSON contains the JSON metadata for the struct
// [OriginTLSComplianceModeGetResponse]
type originTLSComplianceModeGetResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSComplianceModeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSComplianceModeGetResponseJSON) RawJSON() string {
	return r.raw
}

// The identifier of the caching setting.
type OriginTLSComplianceModeGetResponseID string

const (
	OriginTLSComplianceModeGetResponseIDOriginTLSComplianceModes OriginTLSComplianceModeGetResponseID = "origin_tls_compliance_modes"
)

func (r OriginTLSComplianceModeGetResponseID) IsKnown() bool {
	switch r {
	case OriginTLSComplianceModeGetResponseIDOriginTLSComplianceModes:
		return true
	}
	return false
}

type OriginTLSComplianceModeUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// List of TLS compliance modes that constrain the key-exchange algorithms
	// Cloudflare may use when establishing the TLS connection to the zone's origin.
	// Currently supported values are `fips` (FIPS-approved curves) and `pqh`
	// (post-quantum hybrid). Future modes (e.g. `cnsa2`) may be added; clients should
	// treat unknown values as opaque strings. Multiple modes are combined as the
	// intersection of their permitted algorithm lists; selections whose intersection
	// is empty are rejected. An empty list clears the constraint.
	Value param.Field[[]string] `json:"value" api:"required"`
}

func (r OriginTLSComplianceModeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OriginTLSComplianceModeUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginTLSComplianceModeUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  OriginTLSComplianceModeUpdateResponse                `json:"result"`
	JSON    originTLSComplianceModeUpdateResponseEnvelopeJSON    `json:"-"`
}

// originTLSComplianceModeUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [OriginTLSComplianceModeUpdateResponseEnvelope]
type originTLSComplianceModeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSComplianceModeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSComplianceModeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginTLSComplianceModeUpdateResponseEnvelopeSuccess bool

const (
	OriginTLSComplianceModeUpdateResponseEnvelopeSuccessTrue OriginTLSComplianceModeUpdateResponseEnvelopeSuccess = true
)

func (r OriginTLSComplianceModeUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginTLSComplianceModeUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginTLSComplianceModeDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type OriginTLSComplianceModeDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginTLSComplianceModeDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  OriginTLSComplianceModeDeleteResponse                `json:"result"`
	JSON    originTLSComplianceModeDeleteResponseEnvelopeJSON    `json:"-"`
}

// originTLSComplianceModeDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [OriginTLSComplianceModeDeleteResponseEnvelope]
type originTLSComplianceModeDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSComplianceModeDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSComplianceModeDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginTLSComplianceModeDeleteResponseEnvelopeSuccess bool

const (
	OriginTLSComplianceModeDeleteResponseEnvelopeSuccessTrue OriginTLSComplianceModeDeleteResponseEnvelopeSuccess = true
)

func (r OriginTLSComplianceModeDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginTLSComplianceModeDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginTLSComplianceModeEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// List of TLS compliance modes that constrain the key-exchange algorithms
	// Cloudflare may use when establishing the TLS connection to the zone's origin.
	// Currently supported values are `fips` (FIPS-approved curves) and `pqh`
	// (post-quantum hybrid). Future modes (e.g. `cnsa2`) may be added; clients should
	// treat unknown values as opaque strings. Multiple modes are combined as the
	// intersection of their permitted algorithm lists; selections whose intersection
	// is empty are rejected. An empty list clears the constraint.
	Value param.Field[[]string] `json:"value" api:"required"`
}

func (r OriginTLSComplianceModeEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OriginTLSComplianceModeEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginTLSComplianceModeEditResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  OriginTLSComplianceModeEditResponse                `json:"result"`
	JSON    originTLSComplianceModeEditResponseEnvelopeJSON    `json:"-"`
}

// originTLSComplianceModeEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [OriginTLSComplianceModeEditResponseEnvelope]
type originTLSComplianceModeEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSComplianceModeEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSComplianceModeEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginTLSComplianceModeEditResponseEnvelopeSuccess bool

const (
	OriginTLSComplianceModeEditResponseEnvelopeSuccessTrue OriginTLSComplianceModeEditResponseEnvelopeSuccess = true
)

func (r OriginTLSComplianceModeEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginTLSComplianceModeEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginTLSComplianceModeGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type OriginTLSComplianceModeGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginTLSComplianceModeGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  OriginTLSComplianceModeGetResponse                `json:"result"`
	JSON    originTLSComplianceModeGetResponseEnvelopeJSON    `json:"-"`
}

// originTLSComplianceModeGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [OriginTLSComplianceModeGetResponseEnvelope]
type originTLSComplianceModeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSComplianceModeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSComplianceModeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginTLSComplianceModeGetResponseEnvelopeSuccess bool

const (
	OriginTLSComplianceModeGetResponseEnvelopeSuccessTrue OriginTLSComplianceModeGetResponseEnvelopeSuccess = true
)

func (r OriginTLSComplianceModeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginTLSComplianceModeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
