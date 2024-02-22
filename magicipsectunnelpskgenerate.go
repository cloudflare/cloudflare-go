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

// MagicIPSECTunnelPSKGenerateService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewMagicIPSECTunnelPSKGenerateService] method instead.
type MagicIPSECTunnelPSKGenerateService struct {
	Options []option.RequestOption
}

// NewMagicIPSECTunnelPSKGenerateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMagicIPSECTunnelPSKGenerateService(opts ...option.RequestOption) (r *MagicIPSECTunnelPSKGenerateService) {
	r = &MagicIPSECTunnelPSKGenerateService{}
	r.Options = opts
	return
}

// Generates a Pre Shared Key for a specific IPsec tunnel used in the IKE session.
// Use `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes. After a PSK is generated, the PSK is immediately
// persisted to Cloudflare's edge and cannot be retrieved later. Note the PSK in a
// safe place.
func (r *MagicIPSECTunnelPSKGenerateService) New(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicIPSECTunnelPSKGenerateNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIPSECTunnelPSKGenerateNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s/psk_generate", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicIPSECTunnelPSKGenerateNewResponse struct {
	// Identifier
	IPSECTunnelID string `json:"ipsec_tunnel_id"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK string `json:"psk"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata MagicIPSECTunnelPSKGenerateNewResponsePSKMetadata `json:"psk_metadata"`
	JSON        magicIPSECTunnelPSKGenerateNewResponseJSON        `json:"-"`
}

// magicIPSECTunnelPSKGenerateNewResponseJSON contains the JSON metadata for the
// struct [MagicIPSECTunnelPSKGenerateNewResponse]
type magicIPSECTunnelPSKGenerateNewResponseJSON struct {
	IPSECTunnelID apijson.Field
	PSK           apijson.Field
	PSKMetadata   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MagicIPSECTunnelPSKGenerateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type MagicIPSECTunnelPSKGenerateNewResponsePSKMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                             `json:"last_generated_on" format:"date-time"`
	JSON            magicIPSECTunnelPSKGenerateNewResponsePSKMetadataJSON `json:"-"`
}

// magicIPSECTunnelPSKGenerateNewResponsePSKMetadataJSON contains the JSON metadata
// for the struct [MagicIPSECTunnelPSKGenerateNewResponsePSKMetadata]
type magicIPSECTunnelPSKGenerateNewResponsePSKMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicIPSECTunnelPSKGenerateNewResponsePSKMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelPSKGenerateNewResponseEnvelope struct {
	Errors   []MagicIPSECTunnelPSKGenerateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIPSECTunnelPSKGenerateNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIPSECTunnelPSKGenerateNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIPSECTunnelPSKGenerateNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIPSECTunnelPSKGenerateNewResponseEnvelopeJSON    `json:"-"`
}

// magicIPSECTunnelPSKGenerateNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [MagicIPSECTunnelPSKGenerateNewResponseEnvelope]
type magicIPSECTunnelPSKGenerateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelPSKGenerateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelPSKGenerateNewResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    magicIPSECTunnelPSKGenerateNewResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIPSECTunnelPSKGenerateNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicIPSECTunnelPSKGenerateNewResponseEnvelopeErrors]
type magicIPSECTunnelPSKGenerateNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelPSKGenerateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIPSECTunnelPSKGenerateNewResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    magicIPSECTunnelPSKGenerateNewResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIPSECTunnelPSKGenerateNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicIPSECTunnelPSKGenerateNewResponseEnvelopeMessages]
type magicIPSECTunnelPSKGenerateNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIPSECTunnelPSKGenerateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIPSECTunnelPSKGenerateNewResponseEnvelopeSuccess bool

const (
	MagicIPSECTunnelPSKGenerateNewResponseEnvelopeSuccessTrue MagicIPSECTunnelPSKGenerateNewResponseEnvelopeSuccess = true
)
