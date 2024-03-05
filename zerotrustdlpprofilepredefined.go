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

// ZeroTrustDLPProfilePredefinedService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustDLPProfilePredefinedService] method instead.
type ZeroTrustDLPProfilePredefinedService struct {
	Options []option.RequestOption
}

// NewZeroTrustDLPProfilePredefinedService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustDLPProfilePredefinedService(opts ...option.RequestOption) (r *ZeroTrustDLPProfilePredefinedService) {
	r = &ZeroTrustDLPProfilePredefinedService{}
	r.Options = opts
	return
}

// Updates a DLP predefined profile. Only supports enabling/disabling entries.
func (r *ZeroTrustDLPProfilePredefinedService) Update(ctx context.Context, profileID string, params ZeroTrustDLPProfilePredefinedUpdateParams, opts ...option.RequestOption) (res *DLPPredefinedProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Fetches a predefined DLP profile.
func (r *ZeroTrustDLPProfilePredefinedService) Get(ctx context.Context, profileID string, query ZeroTrustDLPProfilePredefinedGetParams, opts ...option.RequestOption) (res *DLPPredefinedProfile, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPProfilePredefinedGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPPredefinedProfile struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness DLPPredefinedProfileContextAwareness `json:"context_awareness"`
	// The entries for this profile.
	Entries []DLPPredefinedProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type DLPPredefinedProfileType `json:"type"`
	JSON dlpPredefinedProfileJSON `json:"-"`
}

// dlpPredefinedProfileJSON contains the JSON metadata for the struct
// [DLPPredefinedProfile]
type dlpPredefinedProfileJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPPredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DLPPredefinedProfile) implementsDLPProfiles() {}

func (r DLPPredefinedProfile) implementsZeroTrustDLPProfileGetResponse() {}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPPredefinedProfileContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip DLPPredefinedProfileContextAwarenessSkip `json:"skip,required"`
	JSON dlpPredefinedProfileContextAwarenessJSON `json:"-"`
}

// dlpPredefinedProfileContextAwarenessJSON contains the JSON metadata for the
// struct [DLPPredefinedProfileContextAwareness]
type dlpPredefinedProfileContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPredefinedProfileContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Content types to exclude from context analysis and return all matches.
type DLPPredefinedProfileContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                         `json:"files,required"`
	JSON  dlpPredefinedProfileContextAwarenessSkipJSON `json:"-"`
}

// dlpPredefinedProfileContextAwarenessSkipJSON contains the JSON metadata for the
// struct [DLPPredefinedProfileContextAwarenessSkip]
type dlpPredefinedProfileContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPredefinedProfileContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A predefined entry that matches a profile
type DLPPredefinedProfileEntry struct {
	// The ID for this entry
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                   `json:"profile_id"`
	JSON      dlpPredefinedProfileEntryJSON `json:"-"`
}

// dlpPredefinedProfileEntryJSON contains the JSON metadata for the struct
// [DLPPredefinedProfileEntry]
type dlpPredefinedProfileEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the profile.
type DLPPredefinedProfileType string

const (
	DLPPredefinedProfileTypePredefined DLPPredefinedProfileType = "predefined"
)

type ZeroTrustDLPProfilePredefinedUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ZeroTrustDLPProfilePredefinedUpdateParamsContextAwareness] `json:"context_awareness"`
	// The entries for this profile.
	Entries param.Field[[]ZeroTrustDLPProfilePredefinedUpdateParamsEntry] `json:"entries"`
}

func (r ZeroTrustDLPProfilePredefinedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ZeroTrustDLPProfilePredefinedUpdateParamsContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip param.Field[ZeroTrustDLPProfilePredefinedUpdateParamsContextAwarenessSkip] `json:"skip,required"`
}

func (r ZeroTrustDLPProfilePredefinedUpdateParamsContextAwareness) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Content types to exclude from context analysis and return all matches.
type ZeroTrustDLPProfilePredefinedUpdateParamsContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files param.Field[bool] `json:"files,required"`
}

func (r ZeroTrustDLPProfilePredefinedUpdateParamsContextAwarenessSkip) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDLPProfilePredefinedUpdateParamsEntry struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustDLPProfilePredefinedUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDLPProfilePredefinedGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDLPProfilePredefinedGetResponseEnvelope struct {
	Errors   []ZeroTrustDLPProfilePredefinedGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDLPProfilePredefinedGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DLPPredefinedProfile                                       `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDLPProfilePredefinedGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDLPProfilePredefinedGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDLPProfilePredefinedGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustDLPProfilePredefinedGetResponseEnvelope]
type zeroTrustDLPProfilePredefinedGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfilePredefinedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPProfilePredefinedGetResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustDLPProfilePredefinedGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPProfilePredefinedGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfilePredefinedGetResponseEnvelopeErrors]
type zeroTrustDLPProfilePredefinedGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfilePredefinedGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPProfilePredefinedGetResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustDLPProfilePredefinedGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPProfilePredefinedGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustDLPProfilePredefinedGetResponseEnvelopeMessages]
type zeroTrustDLPProfilePredefinedGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfilePredefinedGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustDLPProfilePredefinedGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDLPProfilePredefinedGetResponseEnvelopeSuccessTrue ZeroTrustDLPProfilePredefinedGetResponseEnvelopeSuccess = true
)
