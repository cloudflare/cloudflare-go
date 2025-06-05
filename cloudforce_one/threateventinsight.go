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

// Adds an insight to an event
func (r *ThreatEventInsightService) New(ctx context.Context, eventID string, params ThreatEventInsightNewParams, opts ...option.RequestOption) (res *ThreatEventInsightNewResponse, err error) {
	var env ThreatEventInsightNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s/insight/create", params.AccountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an event insight
func (r *ThreatEventInsightService) Delete(ctx context.Context, eventID string, insightID string, body ThreatEventInsightDeleteParams, opts ...option.RequestOption) (res *ThreatEventInsightDeleteResponse, err error) {
	var env ThreatEventInsightDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s/insight/%s", body.AccountID, eventID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an event insight
func (r *ThreatEventInsightService) Edit(ctx context.Context, eventID string, insightID string, params ThreatEventInsightEditParams, opts ...option.RequestOption) (res *ThreatEventInsightEditResponse, err error) {
	var env ThreatEventInsightEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s/insight/%s", params.AccountID, eventID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Reads an event insight
func (r *ThreatEventInsightService) Get(ctx context.Context, eventID string, insightID string, query ThreatEventInsightGetParams, opts ...option.RequestOption) (res *ThreatEventInsightGetResponse, err error) {
	var env ThreatEventInsightGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s/insight/%s", query.AccountID, eventID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ThreatEventInsightNewResponse struct {
	Content string                            `json:"content,required"`
	UUID    string                            `json:"uuid,required"`
	JSON    threatEventInsightNewResponseJSON `json:"-"`
}

// threatEventInsightNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventInsightNewResponse]
type threatEventInsightNewResponseJSON struct {
	Content     apijson.Field
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventInsightNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventInsightNewResponseJSON) RawJSON() string {
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

type ThreatEventInsightEditResponse struct {
	Content string                             `json:"content,required"`
	UUID    string                             `json:"uuid,required"`
	JSON    threatEventInsightEditResponseJSON `json:"-"`
}

// threatEventInsightEditResponseJSON contains the JSON metadata for the struct
// [ThreatEventInsightEditResponse]
type threatEventInsightEditResponseJSON struct {
	Content     apijson.Field
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventInsightEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventInsightEditResponseJSON) RawJSON() string {
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

type ThreatEventInsightNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	Content   param.Field[string] `json:"content,required"`
}

func (r ThreatEventInsightNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventInsightNewResponseEnvelope struct {
	Result  ThreatEventInsightNewResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    threatEventInsightNewResponseEnvelopeJSON `json:"-"`
}

// threatEventInsightNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatEventInsightNewResponseEnvelope]
type threatEventInsightNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventInsightNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventInsightNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatEventInsightDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
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

type ThreatEventInsightEditParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	Content   param.Field[string] `json:"content,required"`
}

func (r ThreatEventInsightEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventInsightEditResponseEnvelope struct {
	Result  ThreatEventInsightEditResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    threatEventInsightEditResponseEnvelopeJSON `json:"-"`
}

// threatEventInsightEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatEventInsightEditResponseEnvelope]
type threatEventInsightEditResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventInsightEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventInsightEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatEventInsightGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
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
