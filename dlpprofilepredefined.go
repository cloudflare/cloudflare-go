// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DLPProfilePredefinedService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDLPProfilePredefinedService]
// method instead.
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

// Fetches a predefined DLP profile.
func (r *DLPProfilePredefinedService) Get(ctx context.Context, accountID string, profileID string, opts ...option.RequestOption) (res *DLPProfilePredefinedGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfilePredefinedGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", accountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP predefined profile. Only supports enabling/disabling entries.
func (r *DLPProfilePredefinedService) Replace(ctx context.Context, accountID string, profileID string, body DLPProfilePredefinedReplaceParams, opts ...option.RequestOption) (res *DLPProfilePredefinedReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", accountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type DLPProfilePredefinedGetResponse struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// The entries for this profile.
	Entries []DLPProfilePredefinedGetResponseEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type DLPProfilePredefinedGetResponseType `json:"type"`
	JSON dlpProfilePredefinedGetResponseJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseJSON contains the JSON metadata for the struct
// [DLPProfilePredefinedGetResponse]
type dlpProfilePredefinedGetResponseJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A predefined entry that matches a profile
type DLPProfilePredefinedGetResponseEntry struct {
	// The ID for this entry
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                              `json:"profile_id"`
	JSON      dlpProfilePredefinedGetResponseEntryJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEntryJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedGetResponseEntry]
type dlpProfilePredefinedGetResponseEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the profile.
type DLPProfilePredefinedGetResponseType string

const (
	DLPProfilePredefinedGetResponseTypePredefined DLPProfilePredefinedGetResponseType = "predefined"
)

type DLPProfilePredefinedReplaceResponse struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// The entries for this profile.
	Entries []DLPProfilePredefinedReplaceResponseEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type DLPProfilePredefinedReplaceResponseType `json:"type"`
	JSON dlpProfilePredefinedReplaceResponseJSON `json:"-"`
}

// dlpProfilePredefinedReplaceResponseJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedReplaceResponse]
type dlpProfilePredefinedReplaceResponseJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPProfilePredefinedReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A predefined entry that matches a profile
type DLPProfilePredefinedReplaceResponseEntry struct {
	// The ID for this entry
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                  `json:"profile_id"`
	JSON      dlpProfilePredefinedReplaceResponseEntryJSON `json:"-"`
}

// dlpProfilePredefinedReplaceResponseEntryJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedReplaceResponseEntry]
type dlpProfilePredefinedReplaceResponseEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedReplaceResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the profile.
type DLPProfilePredefinedReplaceResponseType string

const (
	DLPProfilePredefinedReplaceResponseTypePredefined DLPProfilePredefinedReplaceResponseType = "predefined"
)

type DLPProfilePredefinedGetResponseEnvelope struct {
	Errors   []DLPProfilePredefinedGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfilePredefinedGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DLPProfilePredefinedGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DLPProfilePredefinedGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpProfilePredefinedGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedGetResponseEnvelope]
type dlpProfilePredefinedGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPProfilePredefinedGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    dlpProfilePredefinedGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedGetResponseEnvelopeErrors]
type dlpProfilePredefinedGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPProfilePredefinedGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    dlpProfilePredefinedGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedGetResponseEnvelopeMessages]
type dlpProfilePredefinedGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DLPProfilePredefinedGetResponseEnvelopeSuccess bool

const (
	DLPProfilePredefinedGetResponseEnvelopeSuccessTrue DLPProfilePredefinedGetResponseEnvelopeSuccess = true
)

type DLPProfilePredefinedReplaceParams struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// The entries for this profile.
	Entries param.Field[[]DLPProfilePredefinedReplaceParamsEntry] `json:"entries"`
}

func (r DLPProfilePredefinedReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfilePredefinedReplaceParamsEntry struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r DLPProfilePredefinedReplaceParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
