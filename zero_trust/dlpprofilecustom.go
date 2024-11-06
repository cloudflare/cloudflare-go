// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
	"github.com/tidwall/gjson"
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

// Creates a DLP custom profile.
func (r *DLPProfileCustomService) New(ctx context.Context, params DLPProfileCustomNewParams, opts ...option.RequestOption) (res *DLPProfileCustomNewResponseUnion, err error) {
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

// Union satisfied by [zero_trust.DLPProfileCustomNewResponseCustomProfile],
// [zero_trust.DLPProfileCustomNewResponsePredefinedProfile],
// [zero_trust.DLPProfileCustomNewResponseIntegrationProfile] or
// [zero_trust.DLPProfileCustomNewResponseArray].
type DLPProfileCustomNewResponseUnion interface {
	implementsZeroTrustDLPProfileCustomNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponsePredefinedProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseArray{}),
		},
	)
}

type DLPProfileCustomNewResponseCustomProfile struct {
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness,required"`
	// When the profile was created
	CreatedAt time.Time                                       `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileCustomNewResponseCustomProfileEntry `json:"entries,required"`
	// The name of the profile
	Name       string                                       `json:"name,required"`
	OCREnabled bool                                         `json:"ocr_enabled,required"`
	Type       DLPProfileCustomNewResponseCustomProfileType `json:"type,required"`
	// When the profile was lasted updated
	UpdatedAt           time.Time                                                   `json:"updated_at,required" format:"date-time"`
	ConfidenceThreshold DLPProfileCustomNewResponseCustomProfileConfidenceThreshold `json:"confidence_threshold"`
	// The description of the profile
	Description string                                       `json:"description,nullable"`
	JSON        dlpProfileCustomNewResponseCustomProfileJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseCustomProfile]
type dlpProfileCustomNewResponseCustomProfileJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	ContextAwareness    apijson.Field
	CreatedAt           apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	OCREnabled          apijson.Field
	Type                apijson.Field
	UpdatedAt           apijson.Field
	ConfidenceThreshold apijson.Field
	Description         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomProfile) implementsZeroTrustDLPProfileCustomNewResponseUnion() {
}

type DLPProfileCustomNewResponseCustomProfileEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseCustomProfileEntriesPredefinedConfidence].
	Confidence interface{}                                         `json:"confidence,required"`
	Enabled    bool                                                `json:"enabled,required"`
	Name       string                                              `json:"name,required"`
	Type       DLPProfileCustomNewResponseCustomProfileEntriesType `json:"type,required"`
	// This field can have the runtime type of [interface{}].
	WordList  interface{}                                       `json:"word_list,required"`
	CreatedAt time.Time                                         `json:"created_at" format:"date-time"`
	Pattern   Pattern                                           `json:"pattern"`
	ProfileID string                                            `json:"profile_id,nullable" format:"uuid"`
	Secret    bool                                              `json:"secret"`
	UpdatedAt time.Time                                         `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomNewResponseCustomProfileEntryJSON `json:"-"`
	union     DLPProfileCustomNewResponseCustomProfileEntriesUnion
}

// dlpProfileCustomNewResponseCustomProfileEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseCustomProfileEntry]
type dlpProfileCustomNewResponseCustomProfileEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	WordList    apijson.Field
	CreatedAt   apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	Secret      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dlpProfileCustomNewResponseCustomProfileEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponseCustomProfileEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponseCustomProfileEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponseCustomProfileEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomNewResponseCustomProfileEntriesCustom],
// [zero_trust.DLPProfileCustomNewResponseCustomProfileEntriesPredefined],
// [zero_trust.DLPProfileCustomNewResponseCustomProfileEntriesIntegration],
// [zero_trust.DLPProfileCustomNewResponseCustomProfileEntriesExactData],
// [zero_trust.DLPProfileCustomNewResponseCustomProfileEntriesWordList].
func (r DLPProfileCustomNewResponseCustomProfileEntry) AsUnion() DLPProfileCustomNewResponseCustomProfileEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfileCustomNewResponseCustomProfileEntriesCustom],
// [zero_trust.DLPProfileCustomNewResponseCustomProfileEntriesPredefined],
// [zero_trust.DLPProfileCustomNewResponseCustomProfileEntriesIntegration],
// [zero_trust.DLPProfileCustomNewResponseCustomProfileEntriesExactData] or
// [zero_trust.DLPProfileCustomNewResponseCustomProfileEntriesWordList].
type DLPProfileCustomNewResponseCustomProfileEntriesUnion interface {
	implementsZeroTrustDLPProfileCustomNewResponseCustomProfileEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseCustomProfileEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseCustomProfileEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseCustomProfileEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseCustomProfileEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseCustomProfileEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseCustomProfileEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileCustomNewResponseCustomProfileEntriesCustom struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Pattern   Pattern                                                   `json:"pattern,required"`
	Type      DLPProfileCustomNewResponseCustomProfileEntriesCustomType `json:"type,required"`
	UpdatedAt time.Time                                                 `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                    `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomProfileEntriesCustomJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileEntriesCustomJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponseCustomProfileEntriesCustom]
type dlpProfileCustomNewResponseCustomProfileEntriesCustomJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfileEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomProfileEntriesCustom) implementsZeroTrustDLPProfileCustomNewResponseCustomProfileEntry() {
}

type DLPProfileCustomNewResponseCustomProfileEntriesCustomType string

const (
	DLPProfileCustomNewResponseCustomProfileEntriesCustomTypeCustom DLPProfileCustomNewResponseCustomProfileEntriesCustomType = "custom"
)

func (r DLPProfileCustomNewResponseCustomProfileEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileEntriesPredefined struct {
	ID         string                                                              `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomNewResponseCustomProfileEntriesPredefinedConfidence `json:"confidence,required"`
	Enabled    bool                                                                `json:"enabled,required"`
	Name       string                                                              `json:"name,required"`
	Type       DLPProfileCustomNewResponseCustomProfileEntriesPredefinedType       `json:"type,required"`
	ProfileID  string                                                              `json:"profile_id,nullable" format:"uuid"`
	JSON       dlpProfileCustomNewResponseCustomProfileEntriesPredefinedJSON       `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileEntriesPredefinedJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseCustomProfileEntriesPredefined]
type dlpProfileCustomNewResponseCustomProfileEntriesPredefinedJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfileEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomProfileEntriesPredefined) implementsZeroTrustDLPProfileCustomNewResponseCustomProfileEntry() {
}

type DLPProfileCustomNewResponseCustomProfileEntriesPredefinedConfidence struct {
	// Indicates whether this entry can be made more or less sensitive by setting a
	// confidence threshold. Profiles that use an entry with `available` set to true
	// can use confidence thresholds
	Available bool                                                                    `json:"available,required"`
	JSON      dlpProfileCustomNewResponseCustomProfileEntriesPredefinedConfidenceJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileEntriesPredefinedConfidenceJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomProfileEntriesPredefinedConfidence]
type dlpProfileCustomNewResponseCustomProfileEntriesPredefinedConfidenceJSON struct {
	Available   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfileEntriesPredefinedConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileEntriesPredefinedConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseCustomProfileEntriesPredefinedType string

const (
	DLPProfileCustomNewResponseCustomProfileEntriesPredefinedTypePredefined DLPProfileCustomNewResponseCustomProfileEntriesPredefinedType = "predefined"
)

func (r DLPProfileCustomNewResponseCustomProfileEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileEntriesIntegration struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                           `json:"enabled,required"`
	Name      string                                                         `json:"name,required"`
	Type      DLPProfileCustomNewResponseCustomProfileEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                      `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomProfileEntriesIntegrationJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileEntriesIntegrationJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseCustomProfileEntriesIntegration]
type dlpProfileCustomNewResponseCustomProfileEntriesIntegrationJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfileEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomProfileEntriesIntegration) implementsZeroTrustDLPProfileCustomNewResponseCustomProfileEntry() {
}

type DLPProfileCustomNewResponseCustomProfileEntriesIntegrationType string

const (
	DLPProfileCustomNewResponseCustomProfileEntriesIntegrationTypeIntegration DLPProfileCustomNewResponseCustomProfileEntriesIntegrationType = "integration"
)

func (r DLPProfileCustomNewResponseCustomProfileEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileEntriesExactData struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Secret    bool                                                         `json:"secret,required"`
	Type      DLPProfileCustomNewResponseCustomProfileEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomNewResponseCustomProfileEntriesExactDataJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileEntriesExactDataJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseCustomProfileEntriesExactData]
type dlpProfileCustomNewResponseCustomProfileEntriesExactDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfileEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomProfileEntriesExactData) implementsZeroTrustDLPProfileCustomNewResponseCustomProfileEntry() {
}

type DLPProfileCustomNewResponseCustomProfileEntriesExactDataType string

const (
	DLPProfileCustomNewResponseCustomProfileEntriesExactDataTypeExactData DLPProfileCustomNewResponseCustomProfileEntriesExactDataType = "exact_data"
)

func (r DLPProfileCustomNewResponseCustomProfileEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileEntriesWordList struct {
	ID        string                                                      `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                   `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                        `json:"enabled,required"`
	Name      string                                                      `json:"name,required"`
	Type      DLPProfileCustomNewResponseCustomProfileEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                   `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                 `json:"word_list,required"`
	ProfileID string                                                      `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomProfileEntriesWordListJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomProfileEntriesWordListJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseCustomProfileEntriesWordList]
type dlpProfileCustomNewResponseCustomProfileEntriesWordListJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomProfileEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomProfileEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomProfileEntriesWordList) implementsZeroTrustDLPProfileCustomNewResponseCustomProfileEntry() {
}

type DLPProfileCustomNewResponseCustomProfileEntriesWordListType string

const (
	DLPProfileCustomNewResponseCustomProfileEntriesWordListTypeWordList DLPProfileCustomNewResponseCustomProfileEntriesWordListType = "word_list"
)

func (r DLPProfileCustomNewResponseCustomProfileEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileEntriesType string

const (
	DLPProfileCustomNewResponseCustomProfileEntriesTypeCustom      DLPProfileCustomNewResponseCustomProfileEntriesType = "custom"
	DLPProfileCustomNewResponseCustomProfileEntriesTypePredefined  DLPProfileCustomNewResponseCustomProfileEntriesType = "predefined"
	DLPProfileCustomNewResponseCustomProfileEntriesTypeIntegration DLPProfileCustomNewResponseCustomProfileEntriesType = "integration"
	DLPProfileCustomNewResponseCustomProfileEntriesTypeExactData   DLPProfileCustomNewResponseCustomProfileEntriesType = "exact_data"
	DLPProfileCustomNewResponseCustomProfileEntriesTypeWordList    DLPProfileCustomNewResponseCustomProfileEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponseCustomProfileEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileEntriesTypeCustom, DLPProfileCustomNewResponseCustomProfileEntriesTypePredefined, DLPProfileCustomNewResponseCustomProfileEntriesTypeIntegration, DLPProfileCustomNewResponseCustomProfileEntriesTypeExactData, DLPProfileCustomNewResponseCustomProfileEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileType string

const (
	DLPProfileCustomNewResponseCustomProfileTypeCustom DLPProfileCustomNewResponseCustomProfileType = "custom"
)

func (r DLPProfileCustomNewResponseCustomProfileType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomProfileConfidenceThreshold string

const (
	DLPProfileCustomNewResponseCustomProfileConfidenceThresholdLow      DLPProfileCustomNewResponseCustomProfileConfidenceThreshold = "low"
	DLPProfileCustomNewResponseCustomProfileConfidenceThresholdMedium   DLPProfileCustomNewResponseCustomProfileConfidenceThreshold = "medium"
	DLPProfileCustomNewResponseCustomProfileConfidenceThresholdHigh     DLPProfileCustomNewResponseCustomProfileConfidenceThreshold = "high"
	DLPProfileCustomNewResponseCustomProfileConfidenceThresholdVeryHigh DLPProfileCustomNewResponseCustomProfileConfidenceThreshold = "very_high"
)

func (r DLPProfileCustomNewResponseCustomProfileConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomProfileConfidenceThresholdLow, DLPProfileCustomNewResponseCustomProfileConfidenceThresholdMedium, DLPProfileCustomNewResponseCustomProfileConfidenceThresholdHigh, DLPProfileCustomNewResponseCustomProfileConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfile struct {
	// The id of the predefined profile (uuid)
	ID                string                                              `json:"id,required" format:"uuid"`
	AllowedMatchCount int64                                               `json:"allowed_match_count,required"`
	Entries           []DLPProfileCustomNewResponsePredefinedProfileEntry `json:"entries,required"`
	// The name of the predefined profile
	Name                string                                                          `json:"name,required"`
	Type                DLPProfileCustomNewResponsePredefinedProfileType                `json:"type,required"`
	ConfidenceThreshold DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness"`
	OCREnabled       bool             `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone
	OpenAccess bool                                             `json:"open_access"`
	JSON       dlpProfileCustomNewResponsePredefinedProfileJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponsePredefinedProfile]
type dlpProfileCustomNewResponsePredefinedProfileJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	Type                apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	OCREnabled          apijson.Field
	OpenAccess          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedProfile) implementsZeroTrustDLPProfileCustomNewResponseUnion() {
}

type DLPProfileCustomNewResponsePredefinedProfileEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedConfidence].
	Confidence interface{}                                             `json:"confidence,required"`
	Enabled    bool                                                    `json:"enabled,required"`
	Name       string                                                  `json:"name,required"`
	Type       DLPProfileCustomNewResponsePredefinedProfileEntriesType `json:"type,required"`
	// This field can have the runtime type of [interface{}].
	WordList  interface{}                                           `json:"word_list,required"`
	CreatedAt time.Time                                             `json:"created_at" format:"date-time"`
	Pattern   Pattern                                               `json:"pattern"`
	ProfileID string                                                `json:"profile_id,nullable" format:"uuid"`
	Secret    bool                                                  `json:"secret"`
	UpdatedAt time.Time                                             `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomNewResponsePredefinedProfileEntryJSON `json:"-"`
	union     DLPProfileCustomNewResponsePredefinedProfileEntriesUnion
}

// dlpProfileCustomNewResponsePredefinedProfileEntryJSON contains the JSON metadata
// for the struct [DLPProfileCustomNewResponsePredefinedProfileEntry]
type dlpProfileCustomNewResponsePredefinedProfileEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	WordList    apijson.Field
	CreatedAt   apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	Secret      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponsePredefinedProfileEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponsePredefinedProfileEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomNewResponsePredefinedProfileEntriesCustom],
// [zero_trust.DLPProfileCustomNewResponsePredefinedProfileEntriesPredefined],
// [zero_trust.DLPProfileCustomNewResponsePredefinedProfileEntriesIntegration],
// [zero_trust.DLPProfileCustomNewResponsePredefinedProfileEntriesExactData],
// [zero_trust.DLPProfileCustomNewResponsePredefinedProfileEntriesWordList].
func (r DLPProfileCustomNewResponsePredefinedProfileEntry) AsUnion() DLPProfileCustomNewResponsePredefinedProfileEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfileCustomNewResponsePredefinedProfileEntriesCustom],
// [zero_trust.DLPProfileCustomNewResponsePredefinedProfileEntriesPredefined],
// [zero_trust.DLPProfileCustomNewResponsePredefinedProfileEntriesIntegration],
// [zero_trust.DLPProfileCustomNewResponsePredefinedProfileEntriesExactData] or
// [zero_trust.DLPProfileCustomNewResponsePredefinedProfileEntriesWordList].
type DLPProfileCustomNewResponsePredefinedProfileEntriesUnion interface {
	implementsZeroTrustDLPProfileCustomNewResponsePredefinedProfileEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponsePredefinedProfileEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponsePredefinedProfileEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponsePredefinedProfileEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponsePredefinedProfileEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponsePredefinedProfileEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponsePredefinedProfileEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesCustom struct {
	ID        string                                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                          `json:"enabled,required"`
	Name      string                                                        `json:"name,required"`
	Pattern   Pattern                                                       `json:"pattern,required"`
	Type      DLPProfileCustomNewResponsePredefinedProfileEntriesCustomType `json:"type,required"`
	UpdatedAt time.Time                                                     `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponsePredefinedProfileEntriesCustomJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileEntriesCustomJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponsePredefinedProfileEntriesCustom]
type dlpProfileCustomNewResponsePredefinedProfileEntriesCustomJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesCustom) implementsZeroTrustDLPProfileCustomNewResponsePredefinedProfileEntry() {
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesCustomType string

const (
	DLPProfileCustomNewResponsePredefinedProfileEntriesCustomTypeCustom DLPProfileCustomNewResponsePredefinedProfileEntriesCustomType = "custom"
)

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesPredefined struct {
	ID         string                                                                  `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedConfidence `json:"confidence,required"`
	Enabled    bool                                                                    `json:"enabled,required"`
	Name       string                                                                  `json:"name,required"`
	Type       DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedType       `json:"type,required"`
	ProfileID  string                                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON       dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedJSON       `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedProfileEntriesPredefined]
type dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesPredefined) implementsZeroTrustDLPProfileCustomNewResponsePredefinedProfileEntry() {
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedConfidence struct {
	// Indicates whether this entry can be made more or less sensitive by setting a
	// confidence threshold. Profiles that use an entry with `available` set to true
	// can use confidence thresholds
	Available bool                                                                        `json:"available,required"`
	JSON      dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedConfidenceJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedConfidence]
type dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedConfidenceJSON struct {
	Available   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntriesPredefinedConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedType string

const (
	DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedTypePredefined DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedType = "predefined"
)

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesIntegration struct {
	ID        string                                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                               `json:"enabled,required"`
	Name      string                                                             `json:"name,required"`
	Type      DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                          `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponsePredefinedProfileEntriesIntegrationJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileEntriesIntegrationJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedProfileEntriesIntegration]
type dlpProfileCustomNewResponsePredefinedProfileEntriesIntegrationJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesIntegration) implementsZeroTrustDLPProfileCustomNewResponsePredefinedProfileEntry() {
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationType string

const (
	DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationTypeIntegration DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationType = "integration"
)

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesExactData struct {
	ID        string                                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                             `json:"enabled,required"`
	Name      string                                                           `json:"name,required"`
	Secret    bool                                                             `json:"secret,required"`
	Type      DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                        `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomNewResponsePredefinedProfileEntriesExactDataJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileEntriesExactDataJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedProfileEntriesExactData]
type dlpProfileCustomNewResponsePredefinedProfileEntriesExactDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesExactData) implementsZeroTrustDLPProfileCustomNewResponsePredefinedProfileEntry() {
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataType string

const (
	DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataTypeExactData DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataType = "exact_data"
)

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesWordList struct {
	ID        string                                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                            `json:"enabled,required"`
	Name      string                                                          `json:"name,required"`
	Type      DLPProfileCustomNewResponsePredefinedProfileEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                       `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                     `json:"word_list,required"`
	ProfileID string                                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponsePredefinedProfileEntriesWordListJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedProfileEntriesWordListJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedProfileEntriesWordList]
type dlpProfileCustomNewResponsePredefinedProfileEntriesWordListJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedProfileEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedProfileEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesWordList) implementsZeroTrustDLPProfileCustomNewResponsePredefinedProfileEntry() {
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesWordListType string

const (
	DLPProfileCustomNewResponsePredefinedProfileEntriesWordListTypeWordList DLPProfileCustomNewResponsePredefinedProfileEntriesWordListType = "word_list"
)

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileEntriesType string

const (
	DLPProfileCustomNewResponsePredefinedProfileEntriesTypeCustom      DLPProfileCustomNewResponsePredefinedProfileEntriesType = "custom"
	DLPProfileCustomNewResponsePredefinedProfileEntriesTypePredefined  DLPProfileCustomNewResponsePredefinedProfileEntriesType = "predefined"
	DLPProfileCustomNewResponsePredefinedProfileEntriesTypeIntegration DLPProfileCustomNewResponsePredefinedProfileEntriesType = "integration"
	DLPProfileCustomNewResponsePredefinedProfileEntriesTypeExactData   DLPProfileCustomNewResponsePredefinedProfileEntriesType = "exact_data"
	DLPProfileCustomNewResponsePredefinedProfileEntriesTypeWordList    DLPProfileCustomNewResponsePredefinedProfileEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponsePredefinedProfileEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileEntriesTypeCustom, DLPProfileCustomNewResponsePredefinedProfileEntriesTypePredefined, DLPProfileCustomNewResponsePredefinedProfileEntriesTypeIntegration, DLPProfileCustomNewResponsePredefinedProfileEntriesTypeExactData, DLPProfileCustomNewResponsePredefinedProfileEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileType string

const (
	DLPProfileCustomNewResponsePredefinedProfileTypePredefined DLPProfileCustomNewResponsePredefinedProfileType = "predefined"
)

func (r DLPProfileCustomNewResponsePredefinedProfileType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold string

const (
	DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdLow      DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold = "low"
	DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdMedium   DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold = "medium"
	DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdHigh     DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold = "high"
	DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdVeryHigh DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold = "very_high"
)

func (r DLPProfileCustomNewResponsePredefinedProfileConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdLow, DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdMedium, DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdHigh, DLPProfileCustomNewResponsePredefinedProfileConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfile struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileCustomNewResponseIntegrationProfileEntry `json:"entries,required"`
	Name      string                                               `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationProfileType    `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                                            `json:"description,nullable"`
	JSON        dlpProfileCustomNewResponseIntegrationProfileJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseIntegrationProfile]
type dlpProfileCustomNewResponseIntegrationProfileJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Entries     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationProfile) implementsZeroTrustDLPProfileCustomNewResponseUnion() {
}

type DLPProfileCustomNewResponseIntegrationProfileEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedConfidence].
	Confidence interface{}                                              `json:"confidence,required"`
	Enabled    bool                                                     `json:"enabled,required"`
	Name       string                                                   `json:"name,required"`
	Type       DLPProfileCustomNewResponseIntegrationProfileEntriesType `json:"type,required"`
	// This field can have the runtime type of [interface{}].
	WordList  interface{}                                            `json:"word_list,required"`
	CreatedAt time.Time                                              `json:"created_at" format:"date-time"`
	Pattern   Pattern                                                `json:"pattern"`
	ProfileID string                                                 `json:"profile_id,nullable" format:"uuid"`
	Secret    bool                                                   `json:"secret"`
	UpdatedAt time.Time                                              `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomNewResponseIntegrationProfileEntryJSON `json:"-"`
	union     DLPProfileCustomNewResponseIntegrationProfileEntriesUnion
}

// dlpProfileCustomNewResponseIntegrationProfileEntryJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponseIntegrationProfileEntry]
type dlpProfileCustomNewResponseIntegrationProfileEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	WordList    apijson.Field
	CreatedAt   apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	Secret      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponseIntegrationProfileEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponseIntegrationProfileEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomNewResponseIntegrationProfileEntriesCustom],
// [zero_trust.DLPProfileCustomNewResponseIntegrationProfileEntriesPredefined],
// [zero_trust.DLPProfileCustomNewResponseIntegrationProfileEntriesIntegration],
// [zero_trust.DLPProfileCustomNewResponseIntegrationProfileEntriesExactData],
// [zero_trust.DLPProfileCustomNewResponseIntegrationProfileEntriesWordList].
func (r DLPProfileCustomNewResponseIntegrationProfileEntry) AsUnion() DLPProfileCustomNewResponseIntegrationProfileEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfileCustomNewResponseIntegrationProfileEntriesCustom],
// [zero_trust.DLPProfileCustomNewResponseIntegrationProfileEntriesPredefined],
// [zero_trust.DLPProfileCustomNewResponseIntegrationProfileEntriesIntegration],
// [zero_trust.DLPProfileCustomNewResponseIntegrationProfileEntriesExactData] or
// [zero_trust.DLPProfileCustomNewResponseIntegrationProfileEntriesWordList].
type DLPProfileCustomNewResponseIntegrationProfileEntriesUnion interface {
	implementsZeroTrustDLPProfileCustomNewResponseIntegrationProfileEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseIntegrationProfileEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseIntegrationProfileEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseIntegrationProfileEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseIntegrationProfileEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseIntegrationProfileEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseIntegrationProfileEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesCustom struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                           `json:"enabled,required"`
	Name      string                                                         `json:"name,required"`
	Pattern   Pattern                                                        `json:"pattern,required"`
	Type      DLPProfileCustomNewResponseIntegrationProfileEntriesCustomType `json:"type,required"`
	UpdatedAt time.Time                                                      `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationProfileEntriesCustomJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileEntriesCustomJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseIntegrationProfileEntriesCustom]
type dlpProfileCustomNewResponseIntegrationProfileEntriesCustomJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesCustom) implementsZeroTrustDLPProfileCustomNewResponseIntegrationProfileEntry() {
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesCustomType string

const (
	DLPProfileCustomNewResponseIntegrationProfileEntriesCustomTypeCustom DLPProfileCustomNewResponseIntegrationProfileEntriesCustomType = "custom"
)

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesPredefined struct {
	ID         string                                                                   `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedConfidence `json:"confidence,required"`
	Enabled    bool                                                                     `json:"enabled,required"`
	Name       string                                                                   `json:"name,required"`
	Type       DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedType       `json:"type,required"`
	ProfileID  string                                                                   `json:"profile_id,nullable" format:"uuid"`
	JSON       dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedJSON       `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationProfileEntriesPredefined]
type dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesPredefined) implementsZeroTrustDLPProfileCustomNewResponseIntegrationProfileEntry() {
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedConfidence struct {
	// Indicates whether this entry can be made more or less sensitive by setting a
	// confidence threshold. Profiles that use an entry with `available` set to true
	// can use confidence thresholds
	Available bool                                                                         `json:"available,required"`
	JSON      dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedConfidenceJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedConfidence]
type dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedConfidenceJSON struct {
	Available   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntriesPredefinedConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedType string

const (
	DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedTypePredefined DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedType = "predefined"
)

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesIntegration struct {
	ID        string                                                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                `json:"enabled,required"`
	Name      string                                                              `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                           `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                              `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationProfileEntriesIntegrationJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileEntriesIntegrationJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationProfileEntriesIntegration]
type dlpProfileCustomNewResponseIntegrationProfileEntriesIntegrationJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesIntegration) implementsZeroTrustDLPProfileCustomNewResponseIntegrationProfileEntry() {
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationType string

const (
	DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationTypeIntegration DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationType = "integration"
)

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesExactData struct {
	ID        string                                                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                         `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                              `json:"enabled,required"`
	Name      string                                                            `json:"name,required"`
	Secret    bool                                                              `json:"secret,required"`
	Type      DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                         `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomNewResponseIntegrationProfileEntriesExactDataJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileEntriesExactDataJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationProfileEntriesExactData]
type dlpProfileCustomNewResponseIntegrationProfileEntriesExactDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesExactData) implementsZeroTrustDLPProfileCustomNewResponseIntegrationProfileEntry() {
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataType string

const (
	DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataTypeExactData DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataType = "exact_data"
)

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesWordList struct {
	ID        string                                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                             `json:"enabled,required"`
	Name      string                                                           `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationProfileEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                        `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                      `json:"word_list,required"`
	ProfileID string                                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationProfileEntriesWordListJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationProfileEntriesWordListJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationProfileEntriesWordList]
type dlpProfileCustomNewResponseIntegrationProfileEntriesWordListJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationProfileEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationProfileEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesWordList) implementsZeroTrustDLPProfileCustomNewResponseIntegrationProfileEntry() {
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesWordListType string

const (
	DLPProfileCustomNewResponseIntegrationProfileEntriesWordListTypeWordList DLPProfileCustomNewResponseIntegrationProfileEntriesWordListType = "word_list"
)

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfileEntriesType string

const (
	DLPProfileCustomNewResponseIntegrationProfileEntriesTypeCustom      DLPProfileCustomNewResponseIntegrationProfileEntriesType = "custom"
	DLPProfileCustomNewResponseIntegrationProfileEntriesTypePredefined  DLPProfileCustomNewResponseIntegrationProfileEntriesType = "predefined"
	DLPProfileCustomNewResponseIntegrationProfileEntriesTypeIntegration DLPProfileCustomNewResponseIntegrationProfileEntriesType = "integration"
	DLPProfileCustomNewResponseIntegrationProfileEntriesTypeExactData   DLPProfileCustomNewResponseIntegrationProfileEntriesType = "exact_data"
	DLPProfileCustomNewResponseIntegrationProfileEntriesTypeWordList    DLPProfileCustomNewResponseIntegrationProfileEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponseIntegrationProfileEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileEntriesTypeCustom, DLPProfileCustomNewResponseIntegrationProfileEntriesTypePredefined, DLPProfileCustomNewResponseIntegrationProfileEntriesTypeIntegration, DLPProfileCustomNewResponseIntegrationProfileEntriesTypeExactData, DLPProfileCustomNewResponseIntegrationProfileEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationProfileType string

const (
	DLPProfileCustomNewResponseIntegrationProfileTypeIntegration DLPProfileCustomNewResponseIntegrationProfileType = "integration"
)

func (r DLPProfileCustomNewResponseIntegrationProfileType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationProfileTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseArray []Profile

func (r DLPProfileCustomNewResponseArray) implementsZeroTrustDLPProfileCustomNewResponseUnion() {}

type DLPProfileCustomDeleteResponse = interface{}

type DLPProfileCustomNewParams struct {
	AccountID param.Field[string]                `path:"account_id,required"`
	Body      DLPProfileCustomNewParamsBodyUnion `json:"body,required"`
}

func (r DLPProfileCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DLPProfileCustomNewParamsBody struct {
	Entries       param.Field[interface{}] `json:"entries,required"`
	Profiles      param.Field[interface{}] `json:"profiles,required"`
	SharedEntries param.Field[interface{}] `json:"shared_entries,required"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
	OCREnabled  param.Field[bool]   `json:"ocr_enabled"`
}

func (r DLPProfileCustomNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBody) implementsZeroTrustDLPProfileCustomNewParamsBodyUnion() {}

// Satisfied by [zero_trust.DLPProfileCustomNewParamsBodyProfiles],
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfile],
// [DLPProfileCustomNewParamsBody].
type DLPProfileCustomNewParamsBodyUnion interface {
	implementsZeroTrustDLPProfileCustomNewParamsBodyUnion()
}

type DLPProfileCustomNewParamsBodyProfiles struct {
	Profiles param.Field[[]DLPProfileCustomNewParamsBodyProfilesProfile] `json:"profiles,required"`
}

func (r DLPProfileCustomNewParamsBodyProfiles) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfiles) implementsZeroTrustDLPProfileCustomNewParamsBodyUnion() {
}

type DLPProfileCustomNewParamsBodyProfilesProfile struct {
	Entries param.Field[[]DLPProfileCustomNewParamsBodyProfilesProfilesEntryUnion] `json:"entries,required"`
	Name    param.Field[string]                                                    `json:"name,required"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile
	Description param.Field[string] `json:"description"`
	OCREnabled  param.Field[bool]   `json:"ocr_enabled"`
	// Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your
	// Microsoft Information Protection profiles).
	SharedEntries param.Field[[]DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion] `json:"shared_entries"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomNewParamsBodyProfilesProfilesEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Words   param.Field[interface{}]  `json:"words,required"`
	Pattern param.Field[PatternParam] `json:"pattern"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesEntry) implementsZeroTrustDLPProfileCustomNewParamsBodyProfilesProfilesEntryUnion() {
}

// Satisfied by
// [zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewCustomEntry],
// [zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewWordListEntry],
// [DLPProfileCustomNewParamsBodyProfilesProfilesEntry].
type DLPProfileCustomNewParamsBodyProfilesProfilesEntryUnion interface {
	implementsZeroTrustDLPProfileCustomNewParamsBodyProfilesProfilesEntryUnion()
}

type DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewCustomEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewCustomEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewCustomEntry) implementsZeroTrustDLPProfileCustomNewParamsBodyProfilesProfilesEntryUnion() {
}

type DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewWordListEntry struct {
	Enabled param.Field[bool]     `json:"enabled,required"`
	Name    param.Field[string]   `json:"name,required"`
	Words   param.Field[[]string] `json:"words,required"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewWordListEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesEntriesDLPNewWordListEntry) implementsZeroTrustDLPProfileCustomNewParamsBodyProfilesProfilesEntryUnion() {
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntry struct {
	Enabled   param.Field[bool]                                                                `json:"enabled,required"`
	EntryID   param.Field[string]                                                              `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntry) implementsZeroTrustDLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion() {
}

// Satisfied by
// [zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesCustom],
// [zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesPredefined],
// [zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesIntegration],
// [zero_trust.DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesExactData],
// [DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntry].
type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion interface {
	implementsZeroTrustDLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion()
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesCustom struct {
	Enabled   param.Field[bool]                                                                      `json:"enabled,required"`
	EntryID   param.Field[string]                                                                    `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesCustomEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesCustom) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesCustom) implementsZeroTrustDLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion() {
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesCustomEntryType string

const (
	DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesCustomEntryTypeCustom DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesPredefined struct {
	Enabled   param.Field[bool]                                                                          `json:"enabled,required"`
	EntryID   param.Field[string]                                                                        `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesPredefinedEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesPredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesPredefined) implementsZeroTrustDLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion() {
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesPredefinedEntryType string

const (
	DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesPredefinedEntryTypePredefined DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesIntegration struct {
	Enabled   param.Field[bool]                                                                           `json:"enabled,required"`
	EntryID   param.Field[string]                                                                         `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesIntegrationEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesIntegration) implementsZeroTrustDLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion() {
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesIntegrationEntryType string

const (
	DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesIntegrationEntryTypeIntegration DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesExactData struct {
	Enabled   param.Field[bool]                                                                         `json:"enabled,required"`
	EntryID   param.Field[string]                                                                       `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesExactDataEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesExactData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesExactData) implementsZeroTrustDLPProfileCustomNewParamsBodyProfilesProfilesSharedEntryUnion() {
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesExactDataEntryType string

const (
	DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesExactDataEntryTypeExactData DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType string

const (
	DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypeCustom      DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType = "custom"
	DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypePredefined  DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType = "predefined"
	DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypeIntegration DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType = "integration"
	DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypeExactData   DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType = "exact_data"
)

func (r DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypeCustom, DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypePredefined, DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypeIntegration, DLPProfileCustomNewParamsBodyProfilesProfilesSharedEntriesEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfile struct {
	Entries param.Field[[]DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntryUnion] `json:"entries,required"`
	Name    param.Field[string]                                                       `json:"name,required"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile
	Description param.Field[string] `json:"description"`
	OCREnabled  param.Field[bool]   `json:"ocr_enabled"`
	// Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your
	// Microsoft Information Protection profiles).
	SharedEntries param.Field[[]DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntryUnion] `json:"shared_entries"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfile) implementsZeroTrustDLPProfileCustomNewParamsBodyUnion() {
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Words   param.Field[interface{}]  `json:"words,required"`
	Pattern param.Field[PatternParam] `json:"pattern"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntry) implementsZeroTrustDLPProfileCustomNewParamsBodyDLPNewCustomProfileEntryUnion() {
}

// Satisfied by
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewCustomEntry],
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewWordListEntry],
// [DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntry].
type DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntryUnion interface {
	implementsZeroTrustDLPProfileCustomNewParamsBodyDLPNewCustomProfileEntryUnion()
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewCustomEntry struct {
	Enabled param.Field[bool]         `json:"enabled,required"`
	Name    param.Field[string]       `json:"name,required"`
	Pattern param.Field[PatternParam] `json:"pattern,required"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewCustomEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewCustomEntry) implementsZeroTrustDLPProfileCustomNewParamsBodyDLPNewCustomProfileEntryUnion() {
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewWordListEntry struct {
	Enabled param.Field[bool]     `json:"enabled,required"`
	Name    param.Field[string]   `json:"name,required"`
	Words   param.Field[[]string] `json:"words,required"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewWordListEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileEntriesDLPNewWordListEntry) implementsZeroTrustDLPProfileCustomNewParamsBodyDLPNewCustomProfileEntryUnion() {
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntry struct {
	Enabled   param.Field[bool]                                                                   `json:"enabled,required"`
	EntryID   param.Field[string]                                                                 `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntry) implementsZeroTrustDLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntryUnion() {
}

// Satisfied by
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesCustom],
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesPredefined],
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesIntegration],
// [zero_trust.DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesExactData],
// [DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntry].
type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntryUnion interface {
	implementsZeroTrustDLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntryUnion()
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesCustom struct {
	Enabled   param.Field[bool]                                                                         `json:"enabled,required"`
	EntryID   param.Field[string]                                                                       `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesCustomEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesCustom) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesCustom) implementsZeroTrustDLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntryUnion() {
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesCustomEntryType string

const (
	DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesCustomEntryTypeCustom DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesPredefined struct {
	Enabled   param.Field[bool]                                                                             `json:"enabled,required"`
	EntryID   param.Field[string]                                                                           `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesPredefinedEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesPredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesPredefined) implementsZeroTrustDLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntryUnion() {
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesPredefinedEntryType string

const (
	DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesPredefinedEntryTypePredefined DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesIntegration struct {
	Enabled   param.Field[bool]                                                                              `json:"enabled,required"`
	EntryID   param.Field[string]                                                                            `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesIntegrationEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesIntegration) implementsZeroTrustDLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntryUnion() {
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesIntegrationEntryType string

const (
	DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesIntegrationEntryTypeIntegration DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesExactData struct {
	Enabled   param.Field[bool]                                                                            `json:"enabled,required"`
	EntryID   param.Field[string]                                                                          `json:"entry_id,required" format:"uuid"`
	EntryType param.Field[DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesExactDataEntryType] `json:"entry_type,required"`
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesExactData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesExactData) implementsZeroTrustDLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntryUnion() {
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesExactDataEntryType string

const (
	DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesExactDataEntryTypeExactData DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType string

const (
	DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypeCustom      DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType = "custom"
	DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypePredefined  DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType = "predefined"
	DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypeIntegration DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType = "integration"
	DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypeExactData   DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType = "exact_data"
)

func (r DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypeCustom, DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypePredefined, DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypeIntegration, DLPProfileCustomNewParamsBodyDLPNewCustomProfileSharedEntriesEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DLPProfileCustomNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPProfileCustomNewResponseUnion           `json:"result"`
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
	AccountID           param.Field[string] `path:"account_id,required"`
	Name                param.Field[string] `json:"name,required"`
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile
	Description param.Field[string] `json:"description"`
	// Custom entries from this profile. If this field is omitted, entries owned by
	// this profile will not be changed.
	Entries    param.Field[[]DLPProfileCustomUpdateParamsEntryUnion] `json:"entries"`
	OCREnabled param.Field[bool]                                     `json:"ocr_enabled"`
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
	Result  DLPProfileCustomDeleteResponse                `json:"result,nullable"`
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
