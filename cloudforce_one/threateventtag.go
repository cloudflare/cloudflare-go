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

// ThreatEventTagService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventTagService] method instead.
type ThreatEventTagService struct {
	Options []option.RequestOption
}

// NewThreatEventTagService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThreatEventTagService(opts ...option.RequestOption) (r *ThreatEventTagService) {
	r = &ThreatEventTagService{}
	r.Options = opts
	return
}

// Creates a new tag
func (r *ThreatEventTagService) New(ctx context.Context, accountID float64, body ThreatEventTagNewParams, opts ...option.RequestOption) (res *ThreatEventTagNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/tags/create", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Adds a tag to an event
func (r *ThreatEventTagService) Update(ctx context.Context, accountID float64, eventID string, body ThreatEventTagUpdateParams, opts ...option.RequestOption) (res *ThreatEventTagUpdateResponse, err error) {
	var env ThreatEventTagUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/tag/%s", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Removes a tag from an event
func (r *ThreatEventTagService) Delete(ctx context.Context, accountID float64, eventID string, opts ...option.RequestOption) (res *ThreatEventTagDeleteResponse, err error) {
	var env ThreatEventTagDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/tag/%s", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ThreatEventTagNewResponse struct {
	Name string                        `json:"name,required"`
	UUID string                        `json:"uuid,required"`
	JSON threatEventTagNewResponseJSON `json:"-"`
}

// threatEventTagNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagNewResponse]
type threatEventTagNewResponseJSON struct {
	Name        apijson.Field
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagUpdateResponse struct {
	Success bool                             `json:"success,required"`
	JSON    threatEventTagUpdateResponseJSON `json:"-"`
}

// threatEventTagUpdateResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagUpdateResponse]
type threatEventTagUpdateResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagDeleteResponse struct {
	Success bool                             `json:"success,required"`
	JSON    threatEventTagDeleteResponseJSON `json:"-"`
}

// threatEventTagDeleteResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagDeleteResponse]
type threatEventTagDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewParams struct {
	Name param.Field[string] `json:"name,required"`
}

func (r ThreatEventTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagUpdateParams struct {
	Tags param.Field[[]string] `json:"tags,required"`
}

func (r ThreatEventTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagUpdateResponseEnvelope struct {
	Result  ThreatEventTagUpdateResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    threatEventTagUpdateResponseEnvelopeJSON `json:"-"`
}

// threatEventTagUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatEventTagUpdateResponseEnvelope]
type threatEventTagUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagDeleteResponseEnvelope struct {
	Result  ThreatEventTagDeleteResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    threatEventTagDeleteResponseEnvelopeJSON `json:"-"`
}

// threatEventTagDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatEventTagDeleteResponseEnvelope]
type threatEventTagDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
