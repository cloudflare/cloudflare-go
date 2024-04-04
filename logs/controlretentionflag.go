// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ControlRetentionFlagService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewControlRetentionFlagService]
// method instead.
type ControlRetentionFlagService struct {
	Options []option.RequestOption
}

// NewControlRetentionFlagService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewControlRetentionFlagService(opts ...option.RequestOption) (r *ControlRetentionFlagService) {
	r = &ControlRetentionFlagService{}
	r.Options = opts
	return
}

// Updates log retention flag for Logpull API.
func (r *ControlRetentionFlagService) New(ctx context.Context, zoneIdentifier string, body ControlRetentionFlagNewParams, opts ...option.RequestOption) (res *ControlRetentionFlagNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ControlRetentionFlagNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/logs/control/retention/flag", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets log retention flag for Logpull API.
func (r *ControlRetentionFlagService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ControlRetentionFlagGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ControlRetentionFlagGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/logs/control/retention/flag", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ControlRetentionFlagNewResponse struct {
	Flag bool                                `json:"flag"`
	JSON controlRetentionFlagNewResponseJSON `json:"-"`
}

// controlRetentionFlagNewResponseJSON contains the JSON metadata for the struct
// [ControlRetentionFlagNewResponse]
type controlRetentionFlagNewResponseJSON struct {
	Flag        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlRetentionFlagNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlRetentionFlagNewResponseJSON) RawJSON() string {
	return r.raw
}

type ControlRetentionFlagGetResponse struct {
	Flag bool                                `json:"flag"`
	JSON controlRetentionFlagGetResponseJSON `json:"-"`
}

// controlRetentionFlagGetResponseJSON contains the JSON metadata for the struct
// [ControlRetentionFlagGetResponse]
type controlRetentionFlagGetResponseJSON struct {
	Flag        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlRetentionFlagGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlRetentionFlagGetResponseJSON) RawJSON() string {
	return r.raw
}

type ControlRetentionFlagNewParams struct {
	// The log retention flag for Logpull API.
	Flag param.Field[bool] `json:"flag,required"`
}

func (r ControlRetentionFlagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ControlRetentionFlagNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ControlRetentionFlagNewResponse                           `json:"result,required"`
	// Whether the API call was successful
	Success ControlRetentionFlagNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    controlRetentionFlagNewResponseEnvelopeJSON    `json:"-"`
}

// controlRetentionFlagNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ControlRetentionFlagNewResponseEnvelope]
type controlRetentionFlagNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlRetentionFlagNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlRetentionFlagNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ControlRetentionFlagNewResponseEnvelopeSuccess bool

const (
	ControlRetentionFlagNewResponseEnvelopeSuccessTrue ControlRetentionFlagNewResponseEnvelopeSuccess = true
)

func (r ControlRetentionFlagNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ControlRetentionFlagNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ControlRetentionFlagGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ControlRetentionFlagGetResponse                           `json:"result,required"`
	// Whether the API call was successful
	Success ControlRetentionFlagGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    controlRetentionFlagGetResponseEnvelopeJSON    `json:"-"`
}

// controlRetentionFlagGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ControlRetentionFlagGetResponseEnvelope]
type controlRetentionFlagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlRetentionFlagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlRetentionFlagGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ControlRetentionFlagGetResponseEnvelopeSuccess bool

const (
	ControlRetentionFlagGetResponseEnvelopeSuccessTrue ControlRetentionFlagGetResponseEnvelopeSuccess = true
)

func (r ControlRetentionFlagGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ControlRetentionFlagGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
