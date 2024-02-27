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

// Updates a DLP predefined profile. Only supports enabling/disabling entries.
func (r *DLPProfilePredefinedService) Update(ctx context.Context, profileID string, params DLPProfilePredefinedUpdateParams, opts ...option.RequestOption) (res *DLPProfilePredefinedUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Fetches a predefined DLP profile.
func (r *DLPProfilePredefinedService) Get(ctx context.Context, profileID string, query DLPProfilePredefinedGetParams, opts ...option.RequestOption) (res *DLPProfilePredefinedGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfilePredefinedGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPProfilePredefinedUpdateResponse struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness DLPProfilePredefinedUpdateResponseContextAwareness `json:"context_awareness"`
	// The entries for this profile.
	Entries []DLPProfilePredefinedUpdateResponseEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type DLPProfilePredefinedUpdateResponseType `json:"type"`
	JSON dlpProfilePredefinedUpdateResponseJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseJSON contains the JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponse]
type dlpProfilePredefinedUpdateResponseJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPProfilePredefinedUpdateResponseContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip DLPProfilePredefinedUpdateResponseContextAwarenessSkip `json:"skip,required"`
	JSON dlpProfilePredefinedUpdateResponseContextAwarenessJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseContextAwarenessJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedUpdateResponseContextAwareness]
type dlpProfilePredefinedUpdateResponseContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Content types to exclude from context analysis and return all matches.
type DLPProfilePredefinedUpdateResponseContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                       `json:"files,required"`
	JSON  dlpProfilePredefinedUpdateResponseContextAwarenessSkipJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseContextAwarenessSkipJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedUpdateResponseContextAwarenessSkip]
type dlpProfilePredefinedUpdateResponseContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A predefined entry that matches a profile
type DLPProfilePredefinedUpdateResponseEntry struct {
	// The ID for this entry
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                 `json:"profile_id"`
	JSON      dlpProfilePredefinedUpdateResponseEntryJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEntryJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedUpdateResponseEntry]
type dlpProfilePredefinedUpdateResponseEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the profile.
type DLPProfilePredefinedUpdateResponseType string

const (
	DLPProfilePredefinedUpdateResponseTypePredefined DLPProfilePredefinedUpdateResponseType = "predefined"
)

type DLPProfilePredefinedGetResponse struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness DLPProfilePredefinedGetResponseContextAwareness `json:"context_awareness"`
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
	ContextAwareness  apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPProfilePredefinedGetResponseContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip DLPProfilePredefinedGetResponseContextAwarenessSkip `json:"skip,required"`
	JSON dlpProfilePredefinedGetResponseContextAwarenessJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseContextAwarenessJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedGetResponseContextAwareness]
type dlpProfilePredefinedGetResponseContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Content types to exclude from context analysis and return all matches.
type DLPProfilePredefinedGetResponseContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                    `json:"files,required"`
	JSON  dlpProfilePredefinedGetResponseContextAwarenessSkipJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseContextAwarenessSkipJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponseContextAwarenessSkip]
type dlpProfilePredefinedGetResponseContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
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
