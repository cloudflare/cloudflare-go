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
func (r *MagicIpsecTunnelPskGenerateService) New(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicIpsecTunnelPskGenerateNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicIpsecTunnelPskGenerateNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s/psk_generate", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicIpsecTunnelPskGenerateNewResponse struct {
	// Identifier
	IpsecTunnelID string `json:"ipsec_tunnel_id"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	Psk string `json:"psk"`
	// The PSK metadata that includes when the PSK was generated.
	PskMetadata MagicIpsecTunnelPskGenerateNewResponsePskMetadata `json:"psk_metadata"`
	JSON        magicIpsecTunnelPskGenerateNewResponseJSON        `json:"-"`
}

// magicIpsecTunnelPskGenerateNewResponseJSON contains the JSON metadata for the
// struct [MagicIpsecTunnelPskGenerateNewResponse]
type magicIpsecTunnelPskGenerateNewResponseJSON struct {
	IpsecTunnelID apijson.Field
	Psk           apijson.Field
	PskMetadata   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MagicIpsecTunnelPskGenerateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The PSK metadata that includes when the PSK was generated.
type MagicIpsecTunnelPskGenerateNewResponsePskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time                                             `json:"last_generated_on" format:"date-time"`
	JSON            magicIpsecTunnelPskGenerateNewResponsePskMetadataJSON `json:"-"`
}

// magicIpsecTunnelPskGenerateNewResponsePskMetadataJSON contains the JSON metadata
// for the struct [MagicIpsecTunnelPskGenerateNewResponsePskMetadata]
type magicIpsecTunnelPskGenerateNewResponsePskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicIpsecTunnelPskGenerateNewResponsePskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelPskGenerateNewResponseEnvelope struct {
	Errors   []MagicIpsecTunnelPskGenerateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicIpsecTunnelPskGenerateNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicIpsecTunnelPskGenerateNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicIpsecTunnelPskGenerateNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicIpsecTunnelPskGenerateNewResponseEnvelopeJSON    `json:"-"`
}

// magicIpsecTunnelPskGenerateNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [MagicIpsecTunnelPskGenerateNewResponseEnvelope]
type magicIpsecTunnelPskGenerateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelPskGenerateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelPskGenerateNewResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    magicIpsecTunnelPskGenerateNewResponseEnvelopeErrorsJSON `json:"-"`
}

// magicIpsecTunnelPskGenerateNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicIpsecTunnelPskGenerateNewResponseEnvelopeErrors]
type magicIpsecTunnelPskGenerateNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelPskGenerateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicIpsecTunnelPskGenerateNewResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    magicIpsecTunnelPskGenerateNewResponseEnvelopeMessagesJSON `json:"-"`
}

// magicIpsecTunnelPskGenerateNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicIpsecTunnelPskGenerateNewResponseEnvelopeMessages]
type magicIpsecTunnelPskGenerateNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicIpsecTunnelPskGenerateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicIpsecTunnelPskGenerateNewResponseEnvelopeSuccess bool

const (
	MagicIpsecTunnelPskGenerateNewResponseEnvelopeSuccessTrue MagicIpsecTunnelPskGenerateNewResponseEnvelopeSuccess = true
)
