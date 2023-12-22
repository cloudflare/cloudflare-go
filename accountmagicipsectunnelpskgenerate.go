// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountMagicIpsecTunnelPskGenerateService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountMagicIpsecTunnelPskGenerateService] method instead.
type AccountMagicIpsecTunnelPskGenerateService struct {
	Options []option.RequestOption
}

// NewAccountMagicIpsecTunnelPskGenerateService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountMagicIpsecTunnelPskGenerateService(opts ...option.RequestOption) (r *AccountMagicIpsecTunnelPskGenerateService) {
	r = &AccountMagicIpsecTunnelPskGenerateService{}
	r.Options = opts
	return
}

// Generates a Pre Shared Key for a specific IPsec tunnel used in the IKE session.
// Use `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes. After a PSK is generated, the PSK is immediately
// persisted to Cloudflare's edge and cannot be retrieved later. Note the PSK in a
// safe place.
func (r *AccountMagicIpsecTunnelPskGenerateService) MagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnels(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *PskGenerationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s/psk_generate", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type PskGenerationResponse struct {
	Errors   []PskGenerationResponseError   `json:"errors"`
	Messages []PskGenerationResponseMessage `json:"messages"`
	Result   PskGenerationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success PskGenerationResponseSuccess `json:"success"`
	JSON    pskGenerationResponseJSON    `json:"-"`
}

// pskGenerationResponseJSON contains the JSON metadata for the struct
// [PskGenerationResponse]
type pskGenerationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PskGenerationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PskGenerationResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    pskGenerationResponseErrorJSON `json:"-"`
}

// pskGenerationResponseErrorJSON contains the JSON metadata for the struct
// [PskGenerationResponseError]
type pskGenerationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PskGenerationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PskGenerationResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    pskGenerationResponseMessageJSON `json:"-"`
}

// pskGenerationResponseMessageJSON contains the JSON metadata for the struct
// [PskGenerationResponseMessage]
type pskGenerationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PskGenerationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PskGenerationResponseResult struct {
	// Identifier
	IpsecTunnelID string `json:"ipsec_tunnel_id"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	Psk string `json:"psk"`
	// The PSK metadata that includes when the PSK was generated.
	PskMetadata PskGenerationResponseResultPskMetadata `json:"psk_metadata"`
	JSON        pskGenerationResponseResultJSON        `json:"-"`
}

// pskGenerationResponseResultJSON contains the JSON metadata for the struct
// [PskGenerationResponseResult]
type pskGenerationResponseResultJSON struct {
	IpsecTunnelID apijson.Field
	Psk           apijson.Field
	PskMetadata   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PskGenerationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type PskGenerationResponseResultPskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                  `json:"last_generated_on" format:"date-time"`
	JSON            pskGenerationResponseResultPskMetadataJSON `json:"-"`
}

// pskGenerationResponseResultPskMetadataJSON contains the JSON metadata for the
// struct [PskGenerationResponseResultPskMetadata]
type pskGenerationResponseResultPskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PskGenerationResponseResultPskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PskGenerationResponseSuccess bool

const (
	PskGenerationResponseSuccessTrue PskGenerationResponseSuccess = true
)
