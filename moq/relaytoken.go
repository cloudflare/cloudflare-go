// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moq

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// RelayTokenService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRelayTokenService] method instead.
type RelayTokenService struct {
	Options []option.RequestOption
}

// NewRelayTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRelayTokenService(opts ...option.RequestOption) (r *RelayTokenService) {
	r = &RelayTokenService{}
	r.Options = opts
	return
}

// Generates a new token for the specified type. The old token is immediately
// invalidated. Token value is shown once in the response.
func (r *RelayTokenService) Rotate(ctx context.Context, relayID string, params RelayTokenRotateParams, opts ...option.RequestOption) (res *RelayTokenRotateResponse, err error) {
	var env RelayTokenRotateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if relayID == "" {
		err = errors.New("missing required relay_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/moq/relays/%s/tokens/rotate", params.AccountID, relayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type RelayTokenRotateResponse struct {
	// New token value (shown once). Treat as sensitive.
	Token string                       `json:"token" api:"required"`
	Type  RelayTokenRotateResponseType `json:"type" api:"required"`
	JSON  relayTokenRotateResponseJSON `json:"-"`
}

// relayTokenRotateResponseJSON contains the JSON metadata for the struct
// [RelayTokenRotateResponse]
type relayTokenRotateResponseJSON struct {
	Token       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenRotateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenRotateResponseJSON) RawJSON() string {
	return r.raw
}

type RelayTokenRotateResponseType string

const (
	RelayTokenRotateResponseTypePublishSubscribe RelayTokenRotateResponseType = "publish_subscribe"
	RelayTokenRotateResponseTypeSubscribe        RelayTokenRotateResponseType = "subscribe"
)

func (r RelayTokenRotateResponseType) IsKnown() bool {
	switch r {
	case RelayTokenRotateResponseTypePublishSubscribe, RelayTokenRotateResponseTypeSubscribe:
		return true
	}
	return false
}

type RelayTokenRotateParams struct {
	// Cloudflare account identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Which token type to rotate.
	Type param.Field[RelayTokenRotateParamsType] `json:"type" api:"required"`
}

func (r RelayTokenRotateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Which token type to rotate.
type RelayTokenRotateParamsType string

const (
	RelayTokenRotateParamsTypePublishSubscribe RelayTokenRotateParamsType = "publish_subscribe"
	RelayTokenRotateParamsTypeSubscribe        RelayTokenRotateParamsType = "subscribe"
)

func (r RelayTokenRotateParamsType) IsKnown() bool {
	switch r {
	case RelayTokenRotateParamsTypePublishSubscribe, RelayTokenRotateParamsTypeSubscribe:
		return true
	}
	return false
}

type RelayTokenRotateResponseEnvelope struct {
	Errors   []RelayTokenRotateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RelayTokenRotateResponseEnvelopeMessages `json:"messages" api:"required"`
	Success  bool                                       `json:"success" api:"required"`
	Result   RelayTokenRotateResponse                   `json:"result"`
	JSON     relayTokenRotateResponseEnvelopeJSON       `json:"-"`
}

// relayTokenRotateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RelayTokenRotateResponseEnvelope]
type relayTokenRotateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenRotateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenRotateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RelayTokenRotateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code"`
	Message string                                     `json:"message"`
	JSON    relayTokenRotateResponseEnvelopeErrorsJSON `json:"-"`
}

// relayTokenRotateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RelayTokenRotateResponseEnvelopeErrors]
type relayTokenRotateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenRotateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenRotateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RelayTokenRotateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code"`
	Message string                                       `json:"message"`
	JSON    relayTokenRotateResponseEnvelopeMessagesJSON `json:"-"`
}

// relayTokenRotateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RelayTokenRotateResponseEnvelopeMessages]
type relayTokenRotateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RelayTokenRotateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r relayTokenRotateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
