// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// ThreatEventInsightService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventInsightService] method instead.
type ThreatEventInsightService struct {
	Options []option.RequestOption
}

// NewThreatEventInsightService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventInsightService(opts ...option.RequestOption) (r *ThreatEventInsightService) {
	r = &ThreatEventInsightService{}
	r.Options = opts
	return
}

// Updates an event insight
func (r *ThreatEventInsightService) Update(ctx context.Context, accountID float64, eventID string, insightID string, body ThreatEventInsightUpdateParams, opts ...option.RequestOption) (res *ThreatEventInsightUpdateResponse, err error) {
	var env ThreatEventInsightUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return
	}
	if insightID == "" {
		err = errors.New("missing required insightId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s/insight/%s", accountID, eventID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an event insight
func (r *ThreatEventInsightService) Delete(ctx context.Context, accountID float64, eventID string, insightID string, opts ...option.RequestOption) (res *ThreatEventInsightDeleteResponse, err error) {
	var env ThreatEventInsightDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return
	}
	if insightID == "" {
		err = errors.New("missing required insightId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s/insight/%s", accountID, eventID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Adds an insight to an event
func (r *ThreatEventInsightService) Creat(ctx context.Context, accountID float64, eventID string, body ThreatEventInsightCreatParams, opts ...option.RequestOption) (res *ThreatEventInsightCreatResponse, err error) {
	var env ThreatEventInsightCreatResponseEnvelope
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s/insight/create", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Reads an event insight
func (r *ThreatEventInsightService) Get(ctx context.Context, accountID float64, eventID string, insightID string, opts ...option.RequestOption) (res *ThreatEventInsightGetResponse, err error) {
	var env ThreatEventInsightGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return
	}
	if insightID == "" {
		err = errors.New("missing required insightId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s/insight/%s", accountID, eventID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ThreatEventInsightUpdateResponse struct {
	Content string                               `json:"content,required"`
	UUID    string                               `json:"uuid,required"`
	JSON    threatEventInsightUpdateResponseJSON `json:"-"`
}

// threatEventInsightUpdateResponseJSON contains the JSON metadata for the struct
// [ThreatEventInsightUpdateResponse]
type threatEventInsightUpdateResponseJSON struct {
	Content     apijson.Field
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventInsightUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventInsightUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventInsightDeleteResponse struct {
	Success bool                                 `json:"success,required"`
	JSON    threatEventInsightDeleteResponseJSON `json:"-"`
}

// threatEventInsightDeleteResponseJSON contains the JSON metadata for the struct
// [ThreatEventInsightDeleteResponse]
type threatEventInsightDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventInsightDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventInsightDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventInsightCreatResponse struct {
	Content string                              `json:"content,required"`
	UUID    string                              `json:"uuid,required"`
	JSON    threatEventInsightCreatResponseJSON `json:"-"`
}

// threatEventInsightCreatResponseJSON contains the JSON metadata for the struct
// [ThreatEventInsightCreatResponse]
type threatEventInsightCreatResponseJSON struct {
	Content     apijson.Field
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventInsightCreatResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventInsightCreatResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventInsightGetResponse struct {
	Content string                            `json:"content,required"`
	UUID    string                            `json:"uuid,required"`
	JSON    threatEventInsightGetResponseJSON `json:"-"`
}

// threatEventInsightGetResponseJSON contains the JSON metadata for the struct
// [ThreatEventInsightGetResponse]
type threatEventInsightGetResponseJSON struct {
	Content     apijson.Field
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventInsightGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventInsightGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventInsightUpdateParams struct {
	Content param.Field[string] `json:"content,required"`
}

func (r ThreatEventInsightUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventInsightUpdateResponseEnvelope struct {
	Result  ThreatEventInsightUpdateResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    threatEventInsightUpdateResponseEnvelopeJSON `json:"-"`
}

// threatEventInsightUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatEventInsightUpdateResponseEnvelope]
type threatEventInsightUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventInsightUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventInsightUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatEventInsightDeleteResponseEnvelope struct {
	Result  ThreatEventInsightDeleteResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    threatEventInsightDeleteResponseEnvelopeJSON `json:"-"`
}

// threatEventInsightDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatEventInsightDeleteResponseEnvelope]
type threatEventInsightDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventInsightDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventInsightDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatEventInsightCreatParams struct {
	Content param.Field[string] `json:"content,required"`
}

func (r ThreatEventInsightCreatParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventInsightCreatResponseEnvelope struct {
	Result  ThreatEventInsightCreatResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    threatEventInsightCreatResponseEnvelopeJSON `json:"-"`
}

// threatEventInsightCreatResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatEventInsightCreatResponseEnvelope]
type threatEventInsightCreatResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventInsightCreatResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventInsightCreatResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatEventInsightGetResponseEnvelope struct {
	Result  ThreatEventInsightGetResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    threatEventInsightGetResponseEnvelopeJSON `json:"-"`
}

// threatEventInsightGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatEventInsightGetResponseEnvelope]
type threatEventInsightGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventInsightGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventInsightGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
