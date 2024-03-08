// File generated from our OpenAPI spec by Stainless.

package logs

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// ControlCmbConfigService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewControlCmbConfigService] method
// instead.
type ControlCmbConfigService struct {
	Options []option.RequestOption
}

// NewControlCmbConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewControlCmbConfigService(opts ...option.RequestOption) (r *ControlCmbConfigService) {
	r = &ControlCmbConfigService{}
	r.Options = opts
	return
}

// Updates CMB config.
func (r *ControlCmbConfigService) New(ctx context.Context, params ControlCmbConfigNewParams, opts ...option.RequestOption) (res *ControlCmbConfigNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ControlCmbConfigNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes CMB config.
func (r *ControlCmbConfigService) Delete(ctx context.Context, body ControlCmbConfigDeleteParams, opts ...option.RequestOption) (res *ControlCmbConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ControlCmbConfigDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets CMB config.
func (r *ControlCmbConfigService) Get(ctx context.Context, query ControlCmbConfigGetParams, opts ...option.RequestOption) (res *ControlCmbConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ControlCmbConfigGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ControlCmbConfigNewResponse struct {
	// Comma-separated list of regions.
	Regions string                          `json:"regions"`
	JSON    controlCmbConfigNewResponseJSON `json:"-"`
}

// controlCmbConfigNewResponseJSON contains the JSON metadata for the struct
// [ControlCmbConfigNewResponse]
type controlCmbConfigNewResponseJSON struct {
	Regions     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigNewResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [logs.ControlCmbConfigDeleteResponseUnknown],
// [logs.ControlCmbConfigDeleteResponseArray] or [shared.UnionString].
type ControlCmbConfigDeleteResponse interface {
	ImplementsLogsControlCmbConfigDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ControlCmbConfigDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ControlCmbConfigDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ControlCmbConfigDeleteResponseArray []interface{}

func (r ControlCmbConfigDeleteResponseArray) ImplementsLogsControlCmbConfigDeleteResponse() {}

type ControlCmbConfigGetResponse struct {
	// Comma-separated list of regions.
	Regions string                          `json:"regions"`
	JSON    controlCmbConfigGetResponseJSON `json:"-"`
}

// controlCmbConfigGetResponseJSON contains the JSON metadata for the struct
// [ControlCmbConfigGetResponse]
type controlCmbConfigGetResponseJSON struct {
	Regions     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigGetResponseJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Comma-separated list of regions.
	Regions param.Field[string] `json:"regions"`
}

func (r ControlCmbConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ControlCmbConfigNewResponseEnvelope struct {
	Errors   []ControlCmbConfigNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ControlCmbConfigNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ControlCmbConfigNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ControlCmbConfigNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    controlCmbConfigNewResponseEnvelopeJSON    `json:"-"`
}

// controlCmbConfigNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ControlCmbConfigNewResponseEnvelope]
type controlCmbConfigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    controlCmbConfigNewResponseEnvelopeErrorsJSON `json:"-"`
}

// controlCmbConfigNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ControlCmbConfigNewResponseEnvelopeErrors]
type controlCmbConfigNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    controlCmbConfigNewResponseEnvelopeMessagesJSON `json:"-"`
}

// controlCmbConfigNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ControlCmbConfigNewResponseEnvelopeMessages]
type controlCmbConfigNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ControlCmbConfigNewResponseEnvelopeSuccess bool

const (
	ControlCmbConfigNewResponseEnvelopeSuccessTrue ControlCmbConfigNewResponseEnvelopeSuccess = true
)

type ControlCmbConfigDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ControlCmbConfigDeleteResponseEnvelope struct {
	Errors   []ControlCmbConfigDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ControlCmbConfigDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ControlCmbConfigDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ControlCmbConfigDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    controlCmbConfigDeleteResponseEnvelopeJSON    `json:"-"`
}

// controlCmbConfigDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ControlCmbConfigDeleteResponseEnvelope]
type controlCmbConfigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    controlCmbConfigDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// controlCmbConfigDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ControlCmbConfigDeleteResponseEnvelopeErrors]
type controlCmbConfigDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    controlCmbConfigDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// controlCmbConfigDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ControlCmbConfigDeleteResponseEnvelopeMessages]
type controlCmbConfigDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ControlCmbConfigDeleteResponseEnvelopeSuccess bool

const (
	ControlCmbConfigDeleteResponseEnvelopeSuccessTrue ControlCmbConfigDeleteResponseEnvelopeSuccess = true
)

type ControlCmbConfigGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ControlCmbConfigGetResponseEnvelope struct {
	Errors   []ControlCmbConfigGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ControlCmbConfigGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ControlCmbConfigGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ControlCmbConfigGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    controlCmbConfigGetResponseEnvelopeJSON    `json:"-"`
}

// controlCmbConfigGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ControlCmbConfigGetResponseEnvelope]
type controlCmbConfigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    controlCmbConfigGetResponseEnvelopeErrorsJSON `json:"-"`
}

// controlCmbConfigGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ControlCmbConfigGetResponseEnvelopeErrors]
type controlCmbConfigGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    controlCmbConfigGetResponseEnvelopeMessagesJSON `json:"-"`
}

// controlCmbConfigGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ControlCmbConfigGetResponseEnvelopeMessages]
type controlCmbConfigGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ControlCmbConfigGetResponseEnvelopeSuccess bool

const (
	ControlCmbConfigGetResponseEnvelopeSuccessTrue ControlCmbConfigGetResponseEnvelopeSuccess = true
)
