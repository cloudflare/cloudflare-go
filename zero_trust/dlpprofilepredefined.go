// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// Updates a DLP predefined profile. Only supports enabling/disabling entries.
func (r *DLPProfilePredefinedService) Update(ctx context.Context, profileID string, params DLPProfilePredefinedUpdateParams, opts ...option.RequestOption) (res *DLPPredefinedProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Fetches a predefined DLP profile.
func (r *DLPProfilePredefinedService) Get(ctx context.Context, profileID string, query DLPProfilePredefinedGetParams, opts ...option.RequestOption) (res *DLPPredefinedProfile, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfilePredefinedGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
	// If true, scan images via OCR to determine if any text present matches filters.
	OCREnabled bool `json:"ocr_enabled"`
	// The type of the profile.
	Type shared.UnnamedSchemaRef97 `json:"type"`
	JSON dlpPredefinedProfileJSON  `json:"-"`
}

// dlpPredefinedProfileJSON contains the JSON metadata for the struct
// [DLPPredefinedProfile]
type dlpPredefinedProfileJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	OCREnabled        apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPPredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPredefinedProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPPredefinedProfile) implementsZeroTrustDLPProfiles() {}

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

func (r dlpPredefinedProfileContextAwarenessJSON) RawJSON() string {
	return r.raw
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

func (r dlpPredefinedProfileContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
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

func (r dlpPredefinedProfileEntryJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[DLPProfilePredefinedUpdateParamsContextAwareness] `json:"context_awareness"`
	// The entries for this profile.
	Entries param.Field[[]DLPProfilePredefinedUpdateParamsEntry] `json:"entries"`
	// If true, scan images via OCR to determine if any text present matches filters.
	OCREnabled param.Field[bool] `json:"ocr_enabled"`
}

func (r DLPProfilePredefinedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPProfilePredefinedUpdateParamsContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip param.Field[DLPProfilePredefinedUpdateParamsContextAwarenessSkip] `json:"skip,required"`
}

func (r DLPProfilePredefinedUpdateParamsContextAwareness) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Content types to exclude from context analysis and return all matches.
type DLPProfilePredefinedUpdateParamsContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files param.Field[bool] `json:"files,required"`
}

func (r DLPProfilePredefinedUpdateParamsContextAwarenessSkip) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfilePredefinedUpdateParamsEntry struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r DLPProfilePredefinedUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfilePredefinedGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfilePredefinedGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   DLPPredefinedProfile         `json:"result,required"`
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

func (r dlpProfilePredefinedGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPProfilePredefinedGetResponseEnvelopeSuccess bool

const (
	DLPProfilePredefinedGetResponseEnvelopeSuccessTrue DLPProfilePredefinedGetResponseEnvelopeSuccess = true
)

func (r DLPProfilePredefinedGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
