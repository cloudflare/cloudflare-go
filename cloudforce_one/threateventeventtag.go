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

// ThreatEventEventTagService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventEventTagService] method instead.
type ThreatEventEventTagService struct {
	Options []option.RequestOption
}

// NewThreatEventEventTagService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventEventTagService(opts ...option.RequestOption) (r *ThreatEventEventTagService) {
	r = &ThreatEventEventTagService{}
	r.Options = opts
	return
}

// Removes a tag from an event
func (r *ThreatEventEventTagService) Delete(ctx context.Context, eventID string, body ThreatEventEventTagDeleteParams, opts ...option.RequestOption) (res *ThreatEventEventTagDeleteResponse, err error) {
	var env ThreatEventEventTagDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/event_tag/%s", body.AccountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ThreatEventEventTagDeleteResponse struct {
	Success bool                                  `json:"success,required"`
	JSON    threatEventEventTagDeleteResponseJSON `json:"-"`
}

// threatEventEventTagDeleteResponseJSON contains the JSON metadata for the struct
// [ThreatEventEventTagDeleteResponse]
type threatEventEventTagDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventEventTagDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventEventTagDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventEventTagDeleteParams struct {
	// Account ID
	AccountID param.Field[float64] `path:"account_id,required"`
}

type ThreatEventEventTagDeleteResponseEnvelope struct {
	Result  ThreatEventEventTagDeleteResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    threatEventEventTagDeleteResponseEnvelopeJSON `json:"-"`
}

// threatEventEventTagDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatEventEventTagDeleteResponseEnvelope]
type threatEventEventTagDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventEventTagDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventEventTagDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
