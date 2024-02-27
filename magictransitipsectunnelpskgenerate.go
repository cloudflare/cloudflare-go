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

// MagicTransitIPSECTunnelPSKGenerateService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewMagicTransitIPSECTunnelPSKGenerateService] method instead.
type MagicTransitIPSECTunnelPSKGenerateService struct {
	Options []option.RequestOption
}

// NewMagicTransitIPSECTunnelPSKGenerateService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewMagicTransitIPSECTunnelPSKGenerateService(opts ...option.RequestOption) (r *MagicTransitIPSECTunnelPSKGenerateService) {
	r = &MagicTransitIPSECTunnelPSKGenerateService{}
	r.Options = opts
	return
}

// Generates a Pre Shared Key for a specific IPsec tunnel used in the IKE session.
// Use `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes. After a PSK is generated, the PSK is immediately
// persisted to Cloudflare's edge and cannot be retrieved later. Note the PSK in a
// safe place.
func (r *MagicTransitIPSECTunnelPSKGenerateService) New(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicTransitIPSECTunnelPSKGenerateNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s/psk_generate", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicTransitIPSECTunnelPSKGenerateNewResponse struct {
	// Identifier
	IPSECTunnelID string `json:"ipsec_tunnel_id"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK string `json:"psk"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata MagicTransitIPSECTunnelPSKGenerateNewResponsePSKMetadata `json:"psk_metadata"`
	JSON        magicTransitIPSECTunnelPSKGenerateNewResponseJSON        `json:"-"`
}

// magicTransitIPSECTunnelPSKGenerateNewResponseJSON contains the JSON metadata for
// the struct [MagicTransitIPSECTunnelPSKGenerateNewResponse]
type magicTransitIPSECTunnelPSKGenerateNewResponseJSON struct {
	IPSECTunnelID apijson.Field
	PSK           apijson.Field
	PSKMetadata   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelPSKGenerateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type MagicTransitIPSECTunnelPSKGenerateNewResponsePSKMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                                    `json:"last_generated_on" format:"date-time"`
	JSON            magicTransitIPSECTunnelPSKGenerateNewResponsePSKMetadataJSON `json:"-"`
}

// magicTransitIPSECTunnelPSKGenerateNewResponsePSKMetadataJSON contains the JSON
// metadata for the struct
// [MagicTransitIPSECTunnelPSKGenerateNewResponsePSKMetadata]
type magicTransitIPSECTunnelPSKGenerateNewResponsePSKMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelPSKGenerateNewResponsePSKMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelope struct {
	Errors   []MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitIPSECTunnelPSKGenerateNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeJSON    `json:"-"`
}

// magicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeJSON contains the JSON
// metadata for the struct [MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelope]
type magicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    magicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeErrors]
type magicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    magicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeMessages]
type magicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeSuccess bool

const (
	MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeSuccessTrue MagicTransitIPSECTunnelPSKGenerateNewResponseEnvelopeSuccess = true
)
