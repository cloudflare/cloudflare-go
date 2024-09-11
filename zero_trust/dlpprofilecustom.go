// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/shared"
)

// DLPProfileCustomService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPProfileCustomService] method instead.
type DLPProfileCustomService struct {
	Options []option.RequestOption
}

// NewDLPProfileCustomService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPProfileCustomService(opts ...option.RequestOption) (r *DLPProfileCustomService) {
	r = &DLPProfileCustomService{}
	r.Options = opts
	return
}

// Creates a set of DLP custom profiles.
func (r *DLPProfileCustomService) New(ctx context.Context, params DLPProfileCustomNewParams, opts ...option.RequestOption) (res *[]Profile, err error) {
	var env DLPProfileCustomNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP custom profile.
func (r *DLPProfileCustomService) Update(ctx context.Context, profileID string, params DLPProfileCustomUpdateParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfileCustomUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a DLP custom profile.
func (r *DLPProfileCustomService) Delete(ctx context.Context, profileID string, body DLPProfileCustomDeleteParams, opts ...option.RequestOption) (res *DLPProfileCustomDeleteResponse, err error) {
	var env DLPProfileCustomDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", body.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a custom DLP profile by id.
func (r *DLPProfileCustomService) Get(ctx context.Context, profileID string, query DLPProfileCustomGetParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfileCustomGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Pattern struct {
	Regex      string            `json:"regex,required"`
	Validation PatternValidation `json:"validation"`
	JSON       patternJSON       `json:"-"`
}

// patternJSON contains the JSON metadata for the struct [Pattern]
type patternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r patternJSON) RawJSON() string {
	return r.raw
}

type PatternValidation string

const (
	PatternValidationLuhn PatternValidation = "luhn"
)

func (r PatternValidation) IsKnown() bool {
	switch r {
	case PatternValidationLuhn:
		return true
	}
	return false
}

type PatternParam struct {
	Regex      param.Field[string]            `json:"regex,required"`
	Validation param.Field[PatternValidation] `json:"validation"`
}

func (r PatternParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomDeleteResponse = interface{}

type DLPProfileCustomNewParams struct {
	AccountID param.Field[string]                             `path:"account_id,required"`
	Profiles  param.Field[[]DLPProfileCustomNewParamsProfile] `json:"profiles,required"`
}

func (r DLPProfileCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomNewParamsProfile struct {
	Entries param.Field[[]DLPProfileCustomNewParamsProfilesEntryUnion] `json:"entries,required"`
	Name    param.Field[string]                                        `json:"name,required"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[int64] `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile
	Description param.Field[string] `json:"description"`
	OCREnabled  param.Field[bool]   `json:"ocr_enabled"`
	// Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your
	// Microsoft Information Protection profiles).
	SharedEntries param.Field[[]DLPProfileCustomNewParamsProfilesSharedEntryUnion] `json:"shared_entries"`
}

func (r DLPProfileCustomNewParamsProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomNewParamsProfilesEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern"`
	Words   param.Field[interface{}]  `json:"words,required"`
}

func (r DLPProfileCustomNewParamsProfilesEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsProfilesEntry) implementsZeroTrustDLPProfileCustomNewParamsProfilesEntryUnion() {
}

// Satisfied by
// [zero_trust.DLPProfileCustomNewParamsProfilesEntriesDLPNewCustomEntry],
// [zero_trust.DLPProfileCustomNewParamsProfilesEntriesDLPNewWordListEntry],
// [DLPProfileCustomNewParamsProfilesEntry].
type DLPProfileCustomNewParamsProfilesEntryUnion interface {
	implementsZeroTrustDLPProfileCustomNewParamsProfilesEntryUnion()
}

type DLPProfileCustomNewParamsProfilesEntriesDLPNewCustomEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
}

func (r DLPProfileCustomNewParamsProfilesEntriesDLPNewCustomEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsProfilesEntriesDLPNewCustomEntry) implementsZeroTrustDLPProfileCustomNewParamsProfilesEntryUnion() {
}

type DLPProfileCustomNewParamsProfilesEntriesDLPNewWordListEntry struct {
	Enabled param.Field[bool]     `json:"enabled,required"`
	Name    param.Field[string]   `json:"name,required"`
	Words   param.Field[[]string] `json:"words,required"`
}

func (r DLPProfileCustomNewParamsProfilesEntriesDLPNewWordListEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsProfilesEntriesDLPNewWordListEntry) implementsZeroTrustDLPProfileCustomNewParamsProfilesEntryUnion() {
}

type DLPProfileCustomNewParamsProfilesSharedEntry struct {
	Enabled   param.Field[bool]                                                    `json:"enabled,required"`
	EntryID   param.Field[string]                                                  `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsProfilesSharedEntriesEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsProfilesSharedEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsProfilesSharedEntry) implementsZeroTrustDLPProfileCustomNewParamsProfilesSharedEntryUnion() {
}

// Satisfied by [zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesCustom],
// [zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesPredefined],
// [zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesIntegration],
// [zero_trust.DLPProfileCustomNewParamsProfilesSharedEntriesExactData],
// [DLPProfileCustomNewParamsProfilesSharedEntry].
type DLPProfileCustomNewParamsProfilesSharedEntryUnion interface {
	implementsZeroTrustDLPProfileCustomNewParamsProfilesSharedEntryUnion()
}

type DLPProfileCustomNewParamsProfilesSharedEntriesCustom struct {
	Enabled   param.Field[bool]                                                          `json:"enabled,required"`
	EntryID   param.Field[string]                                                        `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsProfilesSharedEntriesCustom) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsProfilesSharedEntriesCustom) implementsZeroTrustDLPProfileCustomNewParamsProfilesSharedEntryUnion() {
}

type DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryType string

const (
	DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryTypeCustom DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsProfilesSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsProfilesSharedEntriesPredefined struct {
	Enabled   param.Field[bool]                                                              `json:"enabled,required"`
	EntryID   param.Field[string]                                                            `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsProfilesSharedEntriesPredefinedEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsProfilesSharedEntriesPredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsProfilesSharedEntriesPredefined) implementsZeroTrustDLPProfileCustomNewParamsProfilesSharedEntryUnion() {
}

type DLPProfileCustomNewParamsProfilesSharedEntriesPredefinedEntryType string

const (
	DLPProfileCustomNewParamsProfilesSharedEntriesPredefinedEntryTypePredefined DLPProfileCustomNewParamsProfilesSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomNewParamsProfilesSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsProfilesSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsProfilesSharedEntriesIntegration struct {
	Enabled   param.Field[bool]                                                               `json:"enabled,required"`
	EntryID   param.Field[string]                                                             `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsProfilesSharedEntriesIntegrationEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsProfilesSharedEntriesIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsProfilesSharedEntriesIntegration) implementsZeroTrustDLPProfileCustomNewParamsProfilesSharedEntryUnion() {
}

type DLPProfileCustomNewParamsProfilesSharedEntriesIntegrationEntryType string

const (
	DLPProfileCustomNewParamsProfilesSharedEntriesIntegrationEntryTypeIntegration DLPProfileCustomNewParamsProfilesSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomNewParamsProfilesSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsProfilesSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsProfilesSharedEntriesExactData struct {
	Enabled   param.Field[bool]                                                             `json:"enabled,required"`
	EntryID   param.Field[string]                                                           `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsProfilesSharedEntriesExactDataEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsProfilesSharedEntriesExactData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsProfilesSharedEntriesExactData) implementsZeroTrustDLPProfileCustomNewParamsProfilesSharedEntryUnion() {
}

type DLPProfileCustomNewParamsProfilesSharedEntriesExactDataEntryType string

const (
	DLPProfileCustomNewParamsProfilesSharedEntriesExactDataEntryTypeExactData DLPProfileCustomNewParamsProfilesSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomNewParamsProfilesSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsProfilesSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsProfilesSharedEntriesEntryType string

const (
	DLPProfileCustomNewParamsProfilesSharedEntriesEntryTypeCustom      DLPProfileCustomNewParamsProfilesSharedEntriesEntryType = "custom"
	DLPProfileCustomNewParamsProfilesSharedEntriesEntryTypePredefined  DLPProfileCustomNewParamsProfilesSharedEntriesEntryType = "predefined"
	DLPProfileCustomNewParamsProfilesSharedEntriesEntryTypeIntegration DLPProfileCustomNewParamsProfilesSharedEntriesEntryType = "integration"
	DLPProfileCustomNewParamsProfilesSharedEntriesEntryTypeExactData   DLPProfileCustomNewParamsProfilesSharedEntriesEntryType = "exact_data"
)

func (r DLPProfileCustomNewParamsProfilesSharedEntriesEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsProfilesSharedEntriesEntryTypeCustom, DLPProfileCustomNewParamsProfilesSharedEntriesEntryTypePredefined, DLPProfileCustomNewParamsProfilesSharedEntriesEntryTypeIntegration, DLPProfileCustomNewParamsProfilesSharedEntriesEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DLPProfileCustomNewResponseEnvelopeSuccess `json:"success,required"`
	Result  []Profile                                  `json:"result"`
	JSON    dlpProfileCustomNewResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseEnvelope]
type dlpProfileCustomNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPProfileCustomNewResponseEnvelopeSuccess bool

const (
	DLPProfileCustomNewResponseEnvelopeSuccessTrue DLPProfileCustomNewResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfileCustomUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Custom entries from this profile
	Entries           param.Field[[]DLPProfileCustomUpdateParamsEntryUnion] `json:"entries,required"`
	Name              param.Field[string]                                   `json:"name,required"`
	AllowedMatchCount param.Field[int64]                                    `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile
	Description param.Field[string] `json:"description"`
	OCREnabled  param.Field[bool]   `json:"ocr_enabled"`
	// Other entries, e.g. predefined or integration.
	SharedEntries param.Field[[]DLPProfileCustomUpdateParamsSharedEntryUnion] `json:"shared_entries"`
}

func (r DLPProfileCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomUpdateParamsEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
	EntryID param.Field[string]       `json:"entry_id" format:"uuid"`
}

func (r DLPProfileCustomUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsEntry) implementsZeroTrustDLPProfileCustomUpdateParamsEntryUnion() {
}

// Satisfied by
// [zero_trust.DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID],
// [zero_trust.DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry],
// [DLPProfileCustomUpdateParamsEntry].
type DLPProfileCustomUpdateParamsEntryUnion interface {
	implementsZeroTrustDLPProfileCustomUpdateParamsEntryUnion()
}

type DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	EntryID param.Field[string]       `json:"entry_id,required" format:"uuid"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID) implementsZeroTrustDLPProfileCustomUpdateParamsEntryUnion() {
}

type DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry) implementsZeroTrustDLPProfileCustomUpdateParamsEntryUnion() {
}

type DLPProfileCustomUpdateParamsSharedEntry struct {
	Enabled   param.Field[bool]                                               `json:"enabled,required"`
	EntryID   param.Field[string]                                             `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomUpdateParamsSharedEntriesEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomUpdateParamsSharedEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntry) implementsZeroTrustDLPProfileCustomUpdateParamsSharedEntryUnion() {
}

// Satisfied by [zero_trust.DLPProfileCustomUpdateParamsSharedEntriesPredefined],
// [zero_trust.DLPProfileCustomUpdateParamsSharedEntriesIntegration],
// [zero_trust.DLPProfileCustomUpdateParamsSharedEntriesExactData],
// [DLPProfileCustomUpdateParamsSharedEntry].
type DLPProfileCustomUpdateParamsSharedEntryUnion interface {
	implementsZeroTrustDLPProfileCustomUpdateParamsSharedEntryUnion()
}

type DLPProfileCustomUpdateParamsSharedEntriesPredefined struct {
	Enabled   param.Field[bool]                                                         `json:"enabled,required"`
	EntryID   param.Field[string]                                                       `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomUpdateParamsSharedEntriesPredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntriesPredefined) implementsZeroTrustDLPProfileCustomUpdateParamsSharedEntryUnion() {
}

type DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryType string

const (
	DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryTypePredefined DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateParamsSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomUpdateParamsSharedEntriesIntegration struct {
	Enabled   param.Field[bool]                                                          `json:"enabled,required"`
	EntryID   param.Field[string]                                                        `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomUpdateParamsSharedEntriesIntegrationEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomUpdateParamsSharedEntriesIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntriesIntegration) implementsZeroTrustDLPProfileCustomUpdateParamsSharedEntryUnion() {
}

type DLPProfileCustomUpdateParamsSharedEntriesIntegrationEntryType string

const (
	DLPProfileCustomUpdateParamsSharedEntriesIntegrationEntryTypeIntegration DLPProfileCustomUpdateParamsSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomUpdateParamsSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateParamsSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateParamsSharedEntriesExactData struct {
	Enabled   param.Field[bool]                                                        `json:"enabled,required"`
	EntryID   param.Field[string]                                                      `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomUpdateParamsSharedEntriesExactDataEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomUpdateParamsSharedEntriesExactData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntriesExactData) implementsZeroTrustDLPProfileCustomUpdateParamsSharedEntryUnion() {
}

type DLPProfileCustomUpdateParamsSharedEntriesExactDataEntryType string

const (
	DLPProfileCustomUpdateParamsSharedEntriesExactDataEntryTypeExactData DLPProfileCustomUpdateParamsSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomUpdateParamsSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateParamsSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomUpdateParamsSharedEntriesEntryType string

const (
	DLPProfileCustomUpdateParamsSharedEntriesEntryTypePredefined  DLPProfileCustomUpdateParamsSharedEntriesEntryType = "predefined"
	DLPProfileCustomUpdateParamsSharedEntriesEntryTypeIntegration DLPProfileCustomUpdateParamsSharedEntriesEntryType = "integration"
	DLPProfileCustomUpdateParamsSharedEntriesEntryTypeExactData   DLPProfileCustomUpdateParamsSharedEntriesEntryType = "exact_data"
)

func (r DLPProfileCustomUpdateParamsSharedEntriesEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateParamsSharedEntriesEntryTypePredefined, DLPProfileCustomUpdateParamsSharedEntriesEntryTypeIntegration, DLPProfileCustomUpdateParamsSharedEntriesEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DLPProfileCustomUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Profile                                       `json:"result"`
	JSON    dlpProfileCustomUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponseEnvelope]
type dlpProfileCustomUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPProfileCustomUpdateResponseEnvelopeSuccess bool

const (
	DLPProfileCustomUpdateResponseEnvelopeSuccessTrue DLPProfileCustomUpdateResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfileCustomDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileCustomDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DLPProfileCustomDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPProfileCustomDeleteResponse                `json:"result"`
	JSON    dlpProfileCustomDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomDeleteResponseEnvelope]
type dlpProfileCustomDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPProfileCustomDeleteResponseEnvelopeSuccess bool

const (
	DLPProfileCustomDeleteResponseEnvelopeSuccessTrue DLPProfileCustomDeleteResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfileCustomGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileCustomGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DLPProfileCustomGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Profile                                    `json:"result"`
	JSON    dlpProfileCustomGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseEnvelope]
type dlpProfileCustomGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPProfileCustomGetResponseEnvelopeSuccess bool

const (
	DLPProfileCustomGetResponseEnvelopeSuccessTrue DLPProfileCustomGetResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
