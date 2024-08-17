// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DLPProfilePredefinedService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPProfilePredefinedService] method instead.
type DLPProfilePredefinedService struct {
	Options []option.RequestOption
}

// NewDLPProfilePredefinedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPProfilePredefinedService(opts ...option.RequestOption) (r *DLPProfilePredefinedService) {
	r = &DLPProfilePredefinedService{}
	r.Options = opts
	return
}

// Updates a DLP predefined profile. Only supports enabling/disabling entries.
func (r *DLPProfilePredefinedService) Update(ctx context.Context, profileID string, body DLPProfilePredefinedUpdateParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfilePredefinedUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", body.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a predefined DLP profile by id.
func (r *DLPProfilePredefinedService) Get(ctx context.Context, profileID string, query DLPProfilePredefinedGetParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfilePredefinedGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPProfilePredefinedUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfilePredefinedUpdateResponseEnvelope struct {
	Errors     []shared.ResponseInfo                                `json:"errors,required"`
	Messages   []shared.ResponseInfo                                `json:"messages,required"`
	Success    bool                                                 `json:"success,required"`
	Result     Profile                                              `json:"result"`
	ResultInfo DLPProfilePredefinedUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpProfilePredefinedUpdateResponseEnvelopeJSON       `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedUpdateResponseEnvelope]
type dlpProfilePredefinedUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                                    `json:"total_count,required"`
	JSON       dlpProfilePredefinedUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedUpdateResponseEnvelopeResultInfo]
type dlpProfilePredefinedUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfilePredefinedGetResponseEnvelope struct {
	Errors     []shared.ResponseInfo                             `json:"errors,required"`
	Messages   []shared.ResponseInfo                             `json:"messages,required"`
	Success    bool                                              `json:"success,required"`
	Result     Profile                                           `json:"result"`
	ResultInfo DLPProfilePredefinedGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpProfilePredefinedGetResponseEnvelopeJSON       `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedGetResponseEnvelope]
type dlpProfilePredefinedGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                                 `json:"total_count,required"`
	JSON       dlpProfilePredefinedGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedGetResponseEnvelopeResultInfo]
type dlpProfilePredefinedGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
