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

// MagicIpsecTunnelPskGenerateService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewMagicIpsecTunnelPskGenerateService] method instead.
type MagicIpsecTunnelPskGenerateService struct {
	Options []option.RequestOption
}

// NewMagicIpsecTunnelPskGenerateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMagicIpsecTunnelPskGenerateService(opts ...option.RequestOption) (r *MagicIpsecTunnelPskGenerateService) {
	r = &MagicIpsecTunnelPskGenerateService{}
	r.Options = opts
	return
}

// Generates a Pre Shared Key for a specific IPsec tunnel used in the IKE session.
// Use `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes. After a PSK is generated, the PSK is immediately
// persisted to Cloudflare's edge and cannot be retrieved later. Note the PSK in a
// safe place.
func (r *MagicIpsecTunnelPskGenerateService) MagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnels(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s/psk_generate", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponse struct {
	// Identifier
	IpsecTunnelID string `json:"ipsec_tunnel_id"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	Psk string `json:"psk"`
	// The PSK metadata that includes when the PSK was generated.
	PskMetadata MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponsePskMetadata `json:"psk_metadata"`
	JSON        magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseJSON        `json:"-"`
}

// magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponse]
type magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseJSON struct {
	IpsecTunnelID apijson.Field
	Psk           apijson.Field
	PskMetadata   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponsePskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                                                                 `json:"last_generated_on" format:"date-time"`
	JSON            magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponsePskMetadataJSON `json:"-"`
}

// magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponsePskMetadataJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponsePskMetadata]
type magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponsePskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponsePskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelope struct {
	Errors   []MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeJSON    `json:"-"`
}

// magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelope]
type magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeErrors struct {
	Code    int64                                                                                                        `json:"code,required"`
	Message string                                                                                                       `json:"message,required"`
	JSON    magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeErrors]
type magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeMessages struct {
	Code    int64                                                                                                          `json:"code,required"`
	Message string                                                                                                         `json:"message,required"`
	JSON    magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeMessages]
type magicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeSuccess bool

const (
	MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeSuccessTrue MagicIpsecTunnelPskGenerateMagicIPsecTunnelsGeneratePreSharedKeyPskForIPsecTunnelsResponseEnvelopeSuccess = true
)
