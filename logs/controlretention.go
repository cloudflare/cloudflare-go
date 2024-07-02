// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// ControlRetentionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewControlRetentionService] method instead.
type ControlRetentionService struct {
	Options []option.RequestOption
}

// NewControlRetentionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewControlRetentionService(opts ...option.RequestOption) (r *ControlRetentionService) {
	r = &ControlRetentionService{}
	r.Options = opts
	return
}

// Updates log retention flag for Logpull API.
func (r *ControlRetentionService) New(ctx context.Context, zoneIdentifier string, body ControlRetentionNewParams, opts ...option.RequestOption) (res *ControlRetentionNewResponse, err error) {
	var env ControlRetentionNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logs/control/retention/flag", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets log retention flag for Logpull API.
func (r *ControlRetentionService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ControlRetentionGetResponse, err error) {
	var env ControlRetentionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logs/control/retention/flag", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ControlRetentionNewResponse struct {
	Flag bool                            `json:"flag"`
	JSON controlRetentionNewResponseJSON `json:"-"`
}

// controlRetentionNewResponseJSON contains the JSON metadata for the struct
// [ControlRetentionNewResponse]
type controlRetentionNewResponseJSON struct {
	Flag        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlRetentionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlRetentionNewResponseJSON) RawJSON() string {
	return r.raw
}

type ControlRetentionGetResponse struct {
	Flag bool                            `json:"flag"`
	JSON controlRetentionGetResponseJSON `json:"-"`
}

// controlRetentionGetResponseJSON contains the JSON metadata for the struct
// [ControlRetentionGetResponse]
type controlRetentionGetResponseJSON struct {
	Flag        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlRetentionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlRetentionGetResponseJSON) RawJSON() string {
	return r.raw
}

type ControlRetentionNewParams struct {
	// The log retention flag for Logpull API.
	Flag param.Field[bool] `json:"flag,required"`
}

func (r ControlRetentionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ControlRetentionNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   ControlRetentionNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success ControlRetentionNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    controlRetentionNewResponseEnvelopeJSON    `json:"-"`
}

// controlRetentionNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ControlRetentionNewResponseEnvelope]
type controlRetentionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlRetentionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlRetentionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ControlRetentionNewResponseEnvelopeSuccess bool

const (
	ControlRetentionNewResponseEnvelopeSuccessTrue ControlRetentionNewResponseEnvelopeSuccess = true
)

func (r ControlRetentionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ControlRetentionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ControlRetentionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   ControlRetentionGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success ControlRetentionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    controlRetentionGetResponseEnvelopeJSON    `json:"-"`
}

// controlRetentionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ControlRetentionGetResponseEnvelope]
type controlRetentionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlRetentionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlRetentionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ControlRetentionGetResponseEnvelopeSuccess bool

const (
	ControlRetentionGetResponseEnvelopeSuccessTrue ControlRetentionGetResponseEnvelopeSuccess = true
)

func (r ControlRetentionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ControlRetentionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
