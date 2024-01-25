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
func (r *AccountMagicIpsecTunnelPskGenerateService) MagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnels(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s/psk_generate", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponse struct {
	Errors   []AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseError   `json:"errors"`
	Messages []AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseMessage `json:"messages"`
	Result   AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseSuccess `json:"success"`
	JSON    accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseJSON    `json:"-"`
}

// accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponse]
type accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseError struct {
	Code    int64                                                                                                      `json:"code,required"`
	Message string                                                                                                     `json:"message,required"`
	JSON    accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseErrorJSON `json:"-"`
}

// accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseError]
type accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseMessage struct {
	Code    int64                                                                                                        `json:"code,required"`
	Message string                                                                                                       `json:"message,required"`
	JSON    accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseMessageJSON `json:"-"`
}

// accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseMessage]
type accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResult struct {
	// Identifier
	IpsecTunnelID string `json:"ipsec_tunnel_id"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	Psk string `json:"psk"`
	// The PSK metadata that includes when the PSK was generated.
	PskMetadata AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResultPskMetadata `json:"psk_metadata"`
	JSON        accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResultJSON        `json:"-"`
}

// accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResult]
type accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResultJSON struct {
	IpsecTunnelID apijson.Field
	Psk           apijson.Field
	PskMetadata   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResultPskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                                                                              `json:"last_generated_on" format:"date-time"`
	JSON            accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResultPskMetadataJSON `json:"-"`
}

// accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResultPskMetadataJSON
// contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResultPskMetadata]
type accountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResultPskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseResultPskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseSuccess bool

const (
	AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseSuccessTrue AccountMagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseSuccess = true
)
