// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

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

// OhttpGatewayService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOhttpGatewayService] method instead.
type OhttpGatewayService struct {
	Options []option.RequestOption
}

// NewOhttpGatewayService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOhttpGatewayService(opts ...option.RequestOption) (r *OhttpGatewayService) {
	r = &OhttpGatewayService{}
	r.Options = opts
	return
}

// Idempotently enables OHTTP Gateway for the zone. Requires product subscription
// to call.
func (r *OhttpGatewayService) New(ctx context.Context, body OhttpGatewayNewParams, opts ...option.RequestOption) (res *OhttpGatewayNewResponse, err error) {
	var env OhttpGatewayNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/ohttp-gateway", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Idempotently disables Oblivious HTTP Gateway for the zone. Requires product
// subscription to call.
func (r *OhttpGatewayService) Delete(ctx context.Context, body OhttpGatewayDeleteParams, opts ...option.RequestOption) (res *OhttpGatewayDeleteResponse, err error) {
	var env OhttpGatewayDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/ohttp-gateway", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the state of Oblivious HTTP enablement for the zone. Requires product
// subscription to call.
func (r *OhttpGatewayService) Get(ctx context.Context, query OhttpGatewayGetParams, opts ...option.RequestOption) (res *OhttpGatewayGetResponse, err error) {
	var env OhttpGatewayGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/ohttp-gateway", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type OhttpGatewayNewResponse string

const (
	OhttpGatewayNewResponseOn  OhttpGatewayNewResponse = "on"
	OhttpGatewayNewResponseOff OhttpGatewayNewResponse = "off"
)

func (r OhttpGatewayNewResponse) IsKnown() bool {
	switch r {
	case OhttpGatewayNewResponseOn, OhttpGatewayNewResponseOff:
		return true
	}
	return false
}

type OhttpGatewayDeleteResponse string

const (
	OhttpGatewayDeleteResponseOn  OhttpGatewayDeleteResponse = "on"
	OhttpGatewayDeleteResponseOff OhttpGatewayDeleteResponse = "off"
)

func (r OhttpGatewayDeleteResponse) IsKnown() bool {
	switch r {
	case OhttpGatewayDeleteResponseOn, OhttpGatewayDeleteResponseOff:
		return true
	}
	return false
}

type OhttpGatewayGetResponse string

const (
	OhttpGatewayGetResponseOn  OhttpGatewayGetResponse = "on"
	OhttpGatewayGetResponseOff OhttpGatewayGetResponse = "off"
)

func (r OhttpGatewayGetResponse) IsKnown() bool {
	switch r {
	case OhttpGatewayGetResponseOn, OhttpGatewayGetResponseOff:
		return true
	}
	return false
}

type OhttpGatewayNewParams struct {
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

// Standard Cloudflare API response wrapping the result of an operation.
type OhttpGatewayNewResponseEnvelope struct {
	// List of errors, empty on success.
	Errors []OhttpGatewayNewResponseEnvelopeErrors `json:"errors" api:"required"`
	// List of informational messages, empty on success.
	Messages []OhttpGatewayNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   OhttpGatewayNewResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                `json:"success" api:"required"`
	JSON    ohttpGatewayNewResponseEnvelopeJSON `json:"-"`
}

// ohttpGatewayNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OhttpGatewayNewResponseEnvelope]
type ohttpGatewayNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OhttpGatewayNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ohttpGatewayNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an error that the API returns.
type OhttpGatewayNewResponseEnvelopeErrors struct {
	// Numeric error code.
	Code int64 `json:"code" api:"required"`
	// Human-readable error message.
	Message string                                    `json:"message" api:"required"`
	JSON    ohttpGatewayNewResponseEnvelopeErrorsJSON `json:"-"`
}

// ohttpGatewayNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OhttpGatewayNewResponseEnvelopeErrors]
type ohttpGatewayNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OhttpGatewayNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ohttpGatewayNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an informational message that the API returns.
type OhttpGatewayNewResponseEnvelopeMessages struct {
	// Numeric message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string                                      `json:"message" api:"required"`
	JSON    ohttpGatewayNewResponseEnvelopeMessagesJSON `json:"-"`
}

// ohttpGatewayNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OhttpGatewayNewResponseEnvelopeMessages]
type ohttpGatewayNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OhttpGatewayNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ohttpGatewayNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OhttpGatewayDeleteParams struct {
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

// Standard Cloudflare API response wrapping the result of an operation.
type OhttpGatewayDeleteResponseEnvelope struct {
	// List of errors, empty on success.
	Errors []OhttpGatewayDeleteResponseEnvelopeErrors `json:"errors" api:"required"`
	// List of informational messages, empty on success.
	Messages []OhttpGatewayDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   OhttpGatewayDeleteResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                   `json:"success" api:"required"`
	JSON    ohttpGatewayDeleteResponseEnvelopeJSON `json:"-"`
}

// ohttpGatewayDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [OhttpGatewayDeleteResponseEnvelope]
type ohttpGatewayDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OhttpGatewayDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ohttpGatewayDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an error that the API returns.
type OhttpGatewayDeleteResponseEnvelopeErrors struct {
	// Numeric error code.
	Code int64 `json:"code" api:"required"`
	// Human-readable error message.
	Message string                                       `json:"message" api:"required"`
	JSON    ohttpGatewayDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// ohttpGatewayDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OhttpGatewayDeleteResponseEnvelopeErrors]
type ohttpGatewayDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OhttpGatewayDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ohttpGatewayDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an informational message that the API returns.
type OhttpGatewayDeleteResponseEnvelopeMessages struct {
	// Numeric message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string                                         `json:"message" api:"required"`
	JSON    ohttpGatewayDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// ohttpGatewayDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [OhttpGatewayDeleteResponseEnvelopeMessages]
type ohttpGatewayDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OhttpGatewayDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ohttpGatewayDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OhttpGatewayGetParams struct {
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

// Standard Cloudflare API response wrapping the result of an operation.
type OhttpGatewayGetResponseEnvelope struct {
	// List of errors, empty on success.
	Errors []OhttpGatewayGetResponseEnvelopeErrors `json:"errors" api:"required"`
	// List of informational messages, empty on success.
	Messages []OhttpGatewayGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   OhttpGatewayGetResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                `json:"success" api:"required"`
	JSON    ohttpGatewayGetResponseEnvelopeJSON `json:"-"`
}

// ohttpGatewayGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [OhttpGatewayGetResponseEnvelope]
type ohttpGatewayGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OhttpGatewayGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ohttpGatewayGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Represents an error that the API returns.
type OhttpGatewayGetResponseEnvelopeErrors struct {
	// Numeric error code.
	Code int64 `json:"code" api:"required"`
	// Human-readable error message.
	Message string                                    `json:"message" api:"required"`
	JSON    ohttpGatewayGetResponseEnvelopeErrorsJSON `json:"-"`
}

// ohttpGatewayGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OhttpGatewayGetResponseEnvelopeErrors]
type ohttpGatewayGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OhttpGatewayGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ohttpGatewayGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Represents an informational message that the API returns.
type OhttpGatewayGetResponseEnvelopeMessages struct {
	// Numeric message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string                                      `json:"message" api:"required"`
	JSON    ohttpGatewayGetResponseEnvelopeMessagesJSON `json:"-"`
}

// ohttpGatewayGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OhttpGatewayGetResponseEnvelopeMessages]
type ohttpGatewayGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OhttpGatewayGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ohttpGatewayGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
