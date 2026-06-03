// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package csam_scanner

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

// CsamScannerService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCsamScannerService] method instead.
type CsamScannerService struct {
	Options []option.RequestOption
}

// NewCsamScannerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCsamScannerService(opts ...option.RequestOption) (r *CsamScannerService) {
	r = &CsamScannerService{}
	r.Options = opts
	return
}

// Update the CSAM Scanner configuration for a zone. Allows enabling or disabling
// CSAM scanning, updating the notification email, and configuring scanning
// sources.
//
// When a new email is provided, email verification is triggered automatically. The
// `enabled` field is a toggle; the server may adjust it based on whether the
// notification email is verified.
//
// Returns 403 if the zone or account is locked by Trust & Safety.
func (r *CsamScannerService) Edit(ctx context.Context, params CsamScannerEditParams, opts ...option.RequestOption) (res *CsamScannerEditResponse, err error) {
	var env CsamScannerEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/csam_scanner_third_party", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve the current CSAM Scanner configuration for a zone.
//
// The notification email is masked by default in responses.
func (r *CsamScannerService) Get(ctx context.Context, query CsamScannerGetParams, opts ...option.RequestOption) (res *CsamScannerGetResponse, err error) {
	var env CsamScannerGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/csam_scanner_third_party", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// CSAM Scanner configuration for a zone.
type CsamScannerEditResponse struct {
	// The feature identifier.
	ID CsamScannerEditResponseID `json:"id"`
	// Whether the feature state can be changed. When false, the zone or account may be
	// locked by Trust & Safety.
	Editable bool `json:"editable"`
	// When the setting was last modified. Currently always null as the server does not
	// populate this field.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
	// The CSAM Scanner feature configuration values. Contains the notification email
	// and scanning enablement settings.
	Value CsamScannerEditResponseValue `json:"value"`
	JSON  csamScannerEditResponseJSON  `json:"-"`
}

// csamScannerEditResponseJSON contains the JSON metadata for the struct
// [CsamScannerEditResponse]
type csamScannerEditResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CsamScannerEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r csamScannerEditResponseJSON) RawJSON() string {
	return r.raw
}

// The feature identifier.
type CsamScannerEditResponseID string

const (
	CsamScannerEditResponseIDCsamScanner CsamScannerEditResponseID = "csam_scanner"
)

func (r CsamScannerEditResponseID) IsKnown() bool {
	switch r {
	case CsamScannerEditResponseIDCsamScanner:
		return true
	}
	return false
}

// The CSAM Scanner feature configuration values. Contains the notification email
// and scanning enablement settings.
type CsamScannerEditResponseValue struct {
	// Notification email address for CSAM scan results. Masked in responses unless
	// explicitly unmasked via admin endpoint.
	Email string `json:"email"`
	// Current verification state of the notification email.
	EmailState CsamScannerEditResponseValueEmailState `json:"email_state"`
	// Whether CSAM scanning is enabled for this zone.
	Enabled bool `json:"enabled"`
	// Map of scanning sources and their enabled state.
	Sources map[string]bool `json:"sources"`
	// The zone's plan level.
	ZonePlan string                           `json:"zone_plan"`
	JSON     csamScannerEditResponseValueJSON `json:"-"`
}

// csamScannerEditResponseValueJSON contains the JSON metadata for the struct
// [CsamScannerEditResponseValue]
type csamScannerEditResponseValueJSON struct {
	Email       apijson.Field
	EmailState  apijson.Field
	Enabled     apijson.Field
	Sources     apijson.Field
	ZonePlan    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CsamScannerEditResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r csamScannerEditResponseValueJSON) RawJSON() string {
	return r.raw
}

// Current verification state of the notification email.
type CsamScannerEditResponseValueEmailState string

const (
	CsamScannerEditResponseValueEmailStateValid      CsamScannerEditResponseValueEmailState = "valid"
	CsamScannerEditResponseValueEmailStatePending    CsamScannerEditResponseValueEmailState = "pending"
	CsamScannerEditResponseValueEmailStateUnverified CsamScannerEditResponseValueEmailState = "unverified"
)

func (r CsamScannerEditResponseValueEmailState) IsKnown() bool {
	switch r {
	case CsamScannerEditResponseValueEmailStateValid, CsamScannerEditResponseValueEmailStatePending, CsamScannerEditResponseValueEmailStateUnverified:
		return true
	}
	return false
}

// CSAM Scanner configuration for a zone.
type CsamScannerGetResponse struct {
	// The feature identifier.
	ID CsamScannerGetResponseID `json:"id"`
	// Whether the feature state can be changed. When false, the zone or account may be
	// locked by Trust & Safety.
	Editable bool `json:"editable"`
	// When the setting was last modified. Currently always null as the server does not
	// populate this field.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
	// The CSAM Scanner feature configuration values. Contains the notification email
	// and scanning enablement settings.
	Value CsamScannerGetResponseValue `json:"value"`
	JSON  csamScannerGetResponseJSON  `json:"-"`
}

// csamScannerGetResponseJSON contains the JSON metadata for the struct
// [CsamScannerGetResponse]
type csamScannerGetResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CsamScannerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r csamScannerGetResponseJSON) RawJSON() string {
	return r.raw
}

// The feature identifier.
type CsamScannerGetResponseID string

const (
	CsamScannerGetResponseIDCsamScanner CsamScannerGetResponseID = "csam_scanner"
)

func (r CsamScannerGetResponseID) IsKnown() bool {
	switch r {
	case CsamScannerGetResponseIDCsamScanner:
		return true
	}
	return false
}

// The CSAM Scanner feature configuration values. Contains the notification email
// and scanning enablement settings.
type CsamScannerGetResponseValue struct {
	// Notification email address for CSAM scan results. Masked in responses unless
	// explicitly unmasked via admin endpoint.
	Email string `json:"email"`
	// Current verification state of the notification email.
	EmailState CsamScannerGetResponseValueEmailState `json:"email_state"`
	// Whether CSAM scanning is enabled for this zone.
	Enabled bool `json:"enabled"`
	// Map of scanning sources and their enabled state.
	Sources map[string]bool `json:"sources"`
	// The zone's plan level.
	ZonePlan string                          `json:"zone_plan"`
	JSON     csamScannerGetResponseValueJSON `json:"-"`
}

// csamScannerGetResponseValueJSON contains the JSON metadata for the struct
// [CsamScannerGetResponseValue]
type csamScannerGetResponseValueJSON struct {
	Email       apijson.Field
	EmailState  apijson.Field
	Enabled     apijson.Field
	Sources     apijson.Field
	ZonePlan    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CsamScannerGetResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r csamScannerGetResponseValueJSON) RawJSON() string {
	return r.raw
}

// Current verification state of the notification email.
type CsamScannerGetResponseValueEmailState string

const (
	CsamScannerGetResponseValueEmailStateValid      CsamScannerGetResponseValueEmailState = "valid"
	CsamScannerGetResponseValueEmailStatePending    CsamScannerGetResponseValueEmailState = "pending"
	CsamScannerGetResponseValueEmailStateUnverified CsamScannerGetResponseValueEmailState = "unverified"
)

func (r CsamScannerGetResponseValueEmailState) IsKnown() bool {
	switch r {
	case CsamScannerGetResponseValueEmailStateValid, CsamScannerGetResponseValueEmailStatePending, CsamScannerGetResponseValueEmailStateUnverified:
		return true
	}
	return false
}

type CsamScannerEditParams struct {
	// Identifier for the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// The feature identifier.
	ID param.Field[CsamScannerEditParamsID] `json:"id"`
	// Writable CSAM Scanner feature configuration values.
	Value param.Field[CsamScannerEditParamsValue] `json:"value"`
}

func (r CsamScannerEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The feature identifier.
type CsamScannerEditParamsID string

const (
	CsamScannerEditParamsIDCsamScanner CsamScannerEditParamsID = "csam_scanner"
)

func (r CsamScannerEditParamsID) IsKnown() bool {
	switch r {
	case CsamScannerEditParamsIDCsamScanner:
		return true
	}
	return false
}

// Writable CSAM Scanner feature configuration values.
type CsamScannerEditParamsValue struct {
	// Notification email address for CSAM scan results. When changed, email
	// verification is triggered automatically.
	Email param.Field[string] `json:"email"`
	// Whether CSAM scanning is enabled for this zone.
	Enabled param.Field[bool] `json:"enabled"`
	// Set to true to trigger re-sending the email verification. Write-only; never
	// appears in responses (omitted when false).
	ResendEmail param.Field[bool] `json:"resend_email"`
	// Map of scanning sources and their enabled state.
	Sources param.Field[map[string]bool] `json:"sources"`
}

func (r CsamScannerEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CsamScannerEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// CSAM Scanner configuration for a zone.
	Result CsamScannerEditResponse             `json:"result"`
	JSON   csamScannerEditResponseEnvelopeJSON `json:"-"`
}

// csamScannerEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [CsamScannerEditResponseEnvelope]
type csamScannerEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CsamScannerEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r csamScannerEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CsamScannerGetParams struct {
	// Identifier for the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type CsamScannerGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// CSAM Scanner configuration for a zone.
	Result CsamScannerGetResponse             `json:"result"`
	JSON   csamScannerGetResponseEnvelopeJSON `json:"-"`
}

// csamScannerGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CsamScannerGetResponseEnvelope]
type csamScannerGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CsamScannerGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r csamScannerGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
