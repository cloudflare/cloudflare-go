// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneDnssecService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZoneDnssecService] method instead.
type ZoneDnssecService struct {
	Options []option.RequestOption
}

// NewZoneDnssecService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneDnssecService(opts ...option.RequestOption) (r *ZoneDnssecService) {
	r = &ZoneDnssecService{}
	r.Options = opts
	return
}

// Details about DNSSEC status and configuration.
func (r *ZoneDnssecService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *DnssecSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dnssec", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable DNSSEC.
func (r *ZoneDnssecService) Update(ctx context.Context, zoneIdentifier string, body ZoneDnssecUpdateParams, opts ...option.RequestOption) (res *DnssecSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dnssec", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type DnssecSingle struct {
	Errors   []DnssecSingleError   `json:"errors"`
	Messages []DnssecSingleMessage `json:"messages"`
	Result   DnssecSingleResult    `json:"result"`
	// Whether the API call was successful
	Success DnssecSingleSuccess `json:"success"`
	JSON    dnssecSingleJSON
}

// dnssecSingleJSON contains the JSON metadata for the struct [DnssecSingle]
type dnssecSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DnssecSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DnssecSingleError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    dnssecSingleErrorJSON
}

// dnssecSingleErrorJSON contains the JSON metadata for the struct
// [DnssecSingleError]
type dnssecSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DnssecSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DnssecSingleMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    dnssecSingleMessageJSON
}

// dnssecSingleMessageJSON contains the JSON metadata for the struct
// [DnssecSingleMessage]
type dnssecSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DnssecSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DnssecSingleResult struct {
	// Algorithm key code.
	Algorithm string `json:"algorithm,nullable"`
	// Digest hash.
	Digest string `json:"digest,nullable"`
	// Type of digest algorithm.
	DigestAlgorithm string `json:"digest_algorithm,nullable"`
	// Coded type for digest algorithm.
	DigestType string `json:"digest_type,nullable"`
	// Full DS record.
	Ds string `json:"ds,nullable"`
	// Flag for DNSSEC record.
	Flags float64 `json:"flags,nullable"`
	// Code for key tag.
	KeyTag float64 `json:"key_tag,nullable"`
	// Algorithm key type.
	KeyType string `json:"key_type,nullable"`
	// When DNSSEC was last modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Public key for DS record.
	PublicKey string `json:"public_key,nullable"`
	// Status of DNSSEC, based on user-desired state and presence of necessary records.
	Status DnssecSingleResultStatus `json:"status"`
	JSON   dnssecSingleResultJSON
}

// dnssecSingleResultJSON contains the JSON metadata for the struct
// [DnssecSingleResult]
type dnssecSingleResultJSON struct {
	Algorithm       apijson.Field
	Digest          apijson.Field
	DigestAlgorithm apijson.Field
	DigestType      apijson.Field
	Ds              apijson.Field
	Flags           apijson.Field
	KeyTag          apijson.Field
	KeyType         apijson.Field
	ModifiedOn      apijson.Field
	PublicKey       apijson.Field
	Status          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DnssecSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of DNSSEC, based on user-desired state and presence of necessary records.
type DnssecSingleResultStatus string

const (
	DnssecSingleResultStatusActive          DnssecSingleResultStatus = "active"
	DnssecSingleResultStatusPending         DnssecSingleResultStatus = "pending"
	DnssecSingleResultStatusDisabled        DnssecSingleResultStatus = "disabled"
	DnssecSingleResultStatusPendingDisabled DnssecSingleResultStatus = "pending-disabled"
	DnssecSingleResultStatusError           DnssecSingleResultStatus = "error"
)

// Whether the API call was successful
type DnssecSingleSuccess bool

const (
	DnssecSingleSuccessTrue DnssecSingleSuccess = true
)

type ZoneDnssecUpdateParams struct {
	// Status of DNSSEC, based on user-desired state and presence of necessary records.
	Status param.Field[ZoneDnssecUpdateParamsStatus] `json:"status,required"`
}

func (r ZoneDnssecUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of DNSSEC, based on user-desired state and presence of necessary records.
type ZoneDnssecUpdateParamsStatus string

const (
	ZoneDnssecUpdateParamsStatusActive   ZoneDnssecUpdateParamsStatus = "active"
	ZoneDnssecUpdateParamsStatusDisabled ZoneDnssecUpdateParamsStatus = "disabled"
)
