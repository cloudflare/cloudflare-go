// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// DLPProfileService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPProfileService] method instead.
type DLPProfileService struct {
	Options    []option.RequestOption
	Custom     *DLPProfileCustomService
	Predefined *DLPProfilePredefinedService
}

// NewDLPProfileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPProfileService(opts ...option.RequestOption) (r *DLPProfileService) {
	r = &DLPProfileService{}
	r.Options = opts
	r.Custom = NewDLPProfileCustomService(opts...)
	r.Predefined = NewDLPProfilePredefinedService(opts...)
	return
}

// Lists all DLP profiles in an account.
func (r *DLPProfileService) List(ctx context.Context, query DLPProfileListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPProfileListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all DLP profiles in an account.
func (r *DLPProfileService) ListAutoPaging(ctx context.Context, query DLPProfileListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPProfileListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Fetches a DLP profile by ID
func (r *DLPProfileService) Get(ctx context.Context, profileID string, query DLPProfileGetParams, opts ...option.RequestOption) (res *DLPProfileGetResponse, err error) {
	var env DLPProfileGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip SkipConfiguration    `json:"skip,required"`
	JSON contextAwarenessJSON `json:"-"`
}

// contextAwarenessJSON contains the JSON metadata for the struct
// [ContextAwareness]
type contextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type SkipConfiguration struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                  `json:"files,required"`
	JSON  skipConfigurationJSON `json:"-"`
}

// skipConfigurationJSON contains the JSON metadata for the struct
// [SkipConfiguration]
type skipConfigurationJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkipConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skipConfigurationJSON) RawJSON() string {
	return r.raw
}

type DLPProfileListResponse struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// When the profile was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile
	Description string `json:"description,nullable"`
	// This field can have the runtime type of [[]DLPProfileListResponseCustomEntry],
	// [[]DLPProfileListResponsePredefinedEntry],
	// [[]DLPProfileListResponseIntegrationEntry].
	Entries interface{} `json:"entries"`
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// The name of the profile
	Name       string `json:"name,required"`
	OCREnabled bool   `json:"ocr_enabled"`
	// When the profile was lasted updated
	UpdatedAt  time.Time                  `json:"updated_at" format:"date-time"`
	Type       DLPProfileListResponseType `json:"type,required"`
	OpenAccess bool                       `json:"open_access"`
	JSON       dlpProfileListResponseJSON `json:"-"`
	union      DLPProfileListResponseUnion
}

// dlpProfileListResponseJSON contains the JSON metadata for the struct
// [DLPProfileListResponse]
type dlpProfileListResponseJSON struct {
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	CreatedAt         apijson.Field
	Description       apijson.Field
	Entries           apijson.Field
	ID                apijson.Field
	Name              apijson.Field
	OCREnabled        apijson.Field
	UpdatedAt         apijson.Field
	Type              apijson.Field
	OpenAccess        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r dlpProfileListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileListResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileListResponseCustom],
// [zero_trust.DLPProfileListResponsePredefined],
// [zero_trust.DLPProfileListResponseIntegration].
func (r DLPProfileListResponse) AsUnion() DLPProfileListResponseUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfileListResponseCustom],
// [zero_trust.DLPProfileListResponsePredefined] or
// [zero_trust.DLPProfileListResponseIntegration].
type DLPProfileListResponseUnion interface {
	implementsZeroTrustDLPProfileListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileListResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
	)
}

type DLPProfileListResponseCustom struct {
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness,required"`
	// When the profile was created
	CreatedAt time.Time                           `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileListResponseCustomEntry `json:"entries,required"`
	// The name of the profile
	Name       string                           `json:"name,required"`
	OCREnabled bool                             `json:"ocr_enabled,required"`
	Type       DLPProfileListResponseCustomType `json:"type,required"`
	// When the profile was lasted updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                           `json:"description,nullable"`
	JSON        dlpProfileListResponseCustomJSON `json:"-"`
}

// dlpProfileListResponseCustomJSON contains the JSON metadata for the struct
// [DLPProfileListResponseCustom]
type dlpProfileListResponseCustomJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	CreatedAt         apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	OCREnabled        apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	Description       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPProfileListResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseCustom) implementsZeroTrustDLPProfileListResponse() {}

type DLPProfileListResponseCustomEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileListResponseCustomEntriesCustomPattern].
	Pattern   interface{}                             `json:"pattern,required"`
	ProfileID string                                  `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	Type      DLPProfileListResponseCustomEntriesType `json:"type,required"`
	Secret    bool                                    `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                           `json:"word_list,required"`
	JSON     dlpProfileListResponseCustomEntryJSON `json:"-"`
	union    DLPProfileListResponseCustomEntriesUnion
}

// dlpProfileListResponseCustomEntryJSON contains the JSON metadata for the struct
// [DLPProfileListResponseCustomEntry]
type dlpProfileListResponseCustomEntryJSON struct {
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	Type        apijson.Field
	Secret      apijson.Field
	WordList    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dlpProfileListResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileListResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileListResponseCustomEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileListResponseCustomEntriesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileListResponseCustomEntriesCustom],
// [zero_trust.DLPProfileListResponseCustomEntriesPredefined],
// [zero_trust.DLPProfileListResponseCustomEntriesIntegration],
// [zero_trust.DLPProfileListResponseCustomEntriesExactData],
// [zero_trust.DLPProfileListResponseCustomEntriesWordList].
func (r DLPProfileListResponseCustomEntry) AsUnion() DLPProfileListResponseCustomEntriesUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfileListResponseCustomEntriesCustom],
// [zero_trust.DLPProfileListResponseCustomEntriesPredefined],
// [zero_trust.DLPProfileListResponseCustomEntriesIntegration],
// [zero_trust.DLPProfileListResponseCustomEntriesExactData] or
// [zero_trust.DLPProfileListResponseCustomEntriesWordList].
type DLPProfileListResponseCustomEntriesUnion interface {
	implementsZeroTrustDLPProfileListResponseCustomEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileListResponseCustomEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponseCustomEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponseCustomEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponseCustomEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponseCustomEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponseCustomEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileListResponseCustomEntriesCustom struct {
	ID        string                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                             `json:"enabled,required"`
	Name      string                                           `json:"name,required"`
	Pattern   DLPProfileListResponseCustomEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileListResponseCustomEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                        `json:"updated_at,required" format:"date-time"`
	ProfileID string                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileListResponseCustomEntriesCustomJSON    `json:"-"`
}

// dlpProfileListResponseCustomEntriesCustomJSON contains the JSON metadata for the
// struct [DLPProfileListResponseCustomEntriesCustom]
type dlpProfileListResponseCustomEntriesCustomJSON struct {
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

func (r *DLPProfileListResponseCustomEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseCustomEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseCustomEntriesCustom) implementsZeroTrustDLPProfileListResponseCustomEntry() {
}

type DLPProfileListResponseCustomEntriesCustomPattern struct {
	Regex      string                                                     `json:"regex,required"`
	Validation DLPProfileListResponseCustomEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileListResponseCustomEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileListResponseCustomEntriesCustomPatternJSON contains the JSON metadata
// for the struct [DLPProfileListResponseCustomEntriesCustomPattern]
type dlpProfileListResponseCustomEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseCustomEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseCustomEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileListResponseCustomEntriesCustomPatternValidation string

const (
	DLPProfileListResponseCustomEntriesCustomPatternValidationLuhn DLPProfileListResponseCustomEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileListResponseCustomEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileListResponseCustomEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileListResponseCustomEntriesCustomType string

const (
	DLPProfileListResponseCustomEntriesCustomTypeCustom DLPProfileListResponseCustomEntriesCustomType = "custom"
)

func (r DLPProfileListResponseCustomEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseCustomEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileListResponseCustomEntriesPredefined struct {
	ID        string                                            `json:"id,required" format:"uuid"`
	Enabled   bool                                              `json:"enabled,required"`
	Name      string                                            `json:"name,required"`
	Type      DLPProfileListResponseCustomEntriesPredefinedType `json:"type,required"`
	ProfileID string                                            `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileListResponseCustomEntriesPredefinedJSON `json:"-"`
}

// dlpProfileListResponseCustomEntriesPredefinedJSON contains the JSON metadata for
// the struct [DLPProfileListResponseCustomEntriesPredefined]
type dlpProfileListResponseCustomEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseCustomEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseCustomEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseCustomEntriesPredefined) implementsZeroTrustDLPProfileListResponseCustomEntry() {
}

type DLPProfileListResponseCustomEntriesPredefinedType string

const (
	DLPProfileListResponseCustomEntriesPredefinedTypePredefined DLPProfileListResponseCustomEntriesPredefinedType = "predefined"
)

func (r DLPProfileListResponseCustomEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseCustomEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileListResponseCustomEntriesIntegration struct {
	ID        string                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                               `json:"enabled,required"`
	Name      string                                             `json:"name,required"`
	Type      DLPProfileListResponseCustomEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                          `json:"updated_at,required" format:"date-time"`
	ProfileID string                                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileListResponseCustomEntriesIntegrationJSON `json:"-"`
}

// dlpProfileListResponseCustomEntriesIntegrationJSON contains the JSON metadata
// for the struct [DLPProfileListResponseCustomEntriesIntegration]
type dlpProfileListResponseCustomEntriesIntegrationJSON struct {
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

func (r *DLPProfileListResponseCustomEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseCustomEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseCustomEntriesIntegration) implementsZeroTrustDLPProfileListResponseCustomEntry() {
}

type DLPProfileListResponseCustomEntriesIntegrationType string

const (
	DLPProfileListResponseCustomEntriesIntegrationTypeIntegration DLPProfileListResponseCustomEntriesIntegrationType = "integration"
)

func (r DLPProfileListResponseCustomEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseCustomEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileListResponseCustomEntriesExactData struct {
	ID        string                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                             `json:"enabled,required"`
	Name      string                                           `json:"name,required"`
	Secret    bool                                             `json:"secret,required"`
	Type      DLPProfileListResponseCustomEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                        `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileListResponseCustomEntriesExactDataJSON `json:"-"`
}

// dlpProfileListResponseCustomEntriesExactDataJSON contains the JSON metadata for
// the struct [DLPProfileListResponseCustomEntriesExactData]
type dlpProfileListResponseCustomEntriesExactDataJSON struct {
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

func (r *DLPProfileListResponseCustomEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseCustomEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseCustomEntriesExactData) implementsZeroTrustDLPProfileListResponseCustomEntry() {
}

type DLPProfileListResponseCustomEntriesExactDataType string

const (
	DLPProfileListResponseCustomEntriesExactDataTypeExactData DLPProfileListResponseCustomEntriesExactDataType = "exact_data"
)

func (r DLPProfileListResponseCustomEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseCustomEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileListResponseCustomEntriesWordList struct {
	ID        string                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                            `json:"enabled,required"`
	Name      string                                          `json:"name,required"`
	Type      DLPProfileListResponseCustomEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                       `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                     `json:"word_list,required"`
	ProfileID string                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileListResponseCustomEntriesWordListJSON `json:"-"`
}

// dlpProfileListResponseCustomEntriesWordListJSON contains the JSON metadata for
// the struct [DLPProfileListResponseCustomEntriesWordList]
type dlpProfileListResponseCustomEntriesWordListJSON struct {
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

func (r *DLPProfileListResponseCustomEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseCustomEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseCustomEntriesWordList) implementsZeroTrustDLPProfileListResponseCustomEntry() {
}

type DLPProfileListResponseCustomEntriesWordListType string

const (
	DLPProfileListResponseCustomEntriesWordListTypeWordList DLPProfileListResponseCustomEntriesWordListType = "word_list"
)

func (r DLPProfileListResponseCustomEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseCustomEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileListResponseCustomEntriesType string

const (
	DLPProfileListResponseCustomEntriesTypeCustom      DLPProfileListResponseCustomEntriesType = "custom"
	DLPProfileListResponseCustomEntriesTypePredefined  DLPProfileListResponseCustomEntriesType = "predefined"
	DLPProfileListResponseCustomEntriesTypeIntegration DLPProfileListResponseCustomEntriesType = "integration"
	DLPProfileListResponseCustomEntriesTypeExactData   DLPProfileListResponseCustomEntriesType = "exact_data"
	DLPProfileListResponseCustomEntriesTypeWordList    DLPProfileListResponseCustomEntriesType = "word_list"
)

func (r DLPProfileListResponseCustomEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseCustomEntriesTypeCustom, DLPProfileListResponseCustomEntriesTypePredefined, DLPProfileListResponseCustomEntriesTypeIntegration, DLPProfileListResponseCustomEntriesTypeExactData, DLPProfileListResponseCustomEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileListResponseCustomType string

const (
	DLPProfileListResponseCustomTypeCustom DLPProfileListResponseCustomType = "custom"
)

func (r DLPProfileListResponseCustomType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileListResponsePredefined struct {
	// The id of the predefined profile (uuid)
	ID                string                                  `json:"id,required" format:"uuid"`
	AllowedMatchCount int64                                   `json:"allowed_match_count,required"`
	Entries           []DLPProfileListResponsePredefinedEntry `json:"entries,required"`
	// The name of the predefined profile
	Name string                               `json:"name,required"`
	Type DLPProfileListResponsePredefinedType `json:"type,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness                     `json:"context_awareness"`
	OCREnabled       bool                                 `json:"ocr_enabled"`
	OpenAccess       bool                                 `json:"open_access"`
	JSON             dlpProfileListResponsePredefinedJSON `json:"-"`
}

// dlpProfileListResponsePredefinedJSON contains the JSON metadata for the struct
// [DLPProfileListResponsePredefined]
type dlpProfileListResponsePredefinedJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	ContextAwareness  apijson.Field
	OCREnabled        apijson.Field
	OpenAccess        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPProfileListResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponsePredefined) implementsZeroTrustDLPProfileListResponse() {}

type DLPProfileListResponsePredefinedEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileListResponsePredefinedEntriesCustomPattern].
	Pattern   interface{}                                 `json:"pattern,required"`
	ProfileID string                                      `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	Type      DLPProfileListResponsePredefinedEntriesType `json:"type,required"`
	Secret    bool                                        `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                               `json:"word_list,required"`
	JSON     dlpProfileListResponsePredefinedEntryJSON `json:"-"`
	union    DLPProfileListResponsePredefinedEntriesUnion
}

// dlpProfileListResponsePredefinedEntryJSON contains the JSON metadata for the
// struct [DLPProfileListResponsePredefinedEntry]
type dlpProfileListResponsePredefinedEntryJSON struct {
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	Type        apijson.Field
	Secret      apijson.Field
	WordList    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dlpProfileListResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileListResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileListResponsePredefinedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileListResponsePredefinedEntriesUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileListResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfileListResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfileListResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfileListResponsePredefinedEntriesExactData],
// [zero_trust.DLPProfileListResponsePredefinedEntriesWordList].
func (r DLPProfileListResponsePredefinedEntry) AsUnion() DLPProfileListResponsePredefinedEntriesUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfileListResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfileListResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfileListResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfileListResponsePredefinedEntriesExactData] or
// [zero_trust.DLPProfileListResponsePredefinedEntriesWordList].
type DLPProfileListResponsePredefinedEntriesUnion interface {
	implementsZeroTrustDLPProfileListResponsePredefinedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileListResponsePredefinedEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponsePredefinedEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponsePredefinedEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponsePredefinedEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponsePredefinedEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponsePredefinedEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileListResponsePredefinedEntriesCustom struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                 `json:"enabled,required"`
	Name      string                                               `json:"name,required"`
	Pattern   DLPProfileListResponsePredefinedEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileListResponsePredefinedEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	ProfileID string                                               `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileListResponsePredefinedEntriesCustomJSON    `json:"-"`
}

// dlpProfileListResponsePredefinedEntriesCustomJSON contains the JSON metadata for
// the struct [DLPProfileListResponsePredefinedEntriesCustom]
type dlpProfileListResponsePredefinedEntriesCustomJSON struct {
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

func (r *DLPProfileListResponsePredefinedEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponsePredefinedEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponsePredefinedEntriesCustom) implementsZeroTrustDLPProfileListResponsePredefinedEntry() {
}

type DLPProfileListResponsePredefinedEntriesCustomPattern struct {
	Regex      string                                                         `json:"regex,required"`
	Validation DLPProfileListResponsePredefinedEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileListResponsePredefinedEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileListResponsePredefinedEntriesCustomPatternJSON contains the JSON
// metadata for the struct [DLPProfileListResponsePredefinedEntriesCustomPattern]
type dlpProfileListResponsePredefinedEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponsePredefinedEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponsePredefinedEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileListResponsePredefinedEntriesCustomPatternValidation string

const (
	DLPProfileListResponsePredefinedEntriesCustomPatternValidationLuhn DLPProfileListResponsePredefinedEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileListResponsePredefinedEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileListResponsePredefinedEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileListResponsePredefinedEntriesCustomType string

const (
	DLPProfileListResponsePredefinedEntriesCustomTypeCustom DLPProfileListResponsePredefinedEntriesCustomType = "custom"
)

func (r DLPProfileListResponsePredefinedEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileListResponsePredefinedEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileListResponsePredefinedEntriesPredefined struct {
	ID        string                                                `json:"id,required" format:"uuid"`
	Enabled   bool                                                  `json:"enabled,required"`
	Name      string                                                `json:"name,required"`
	Type      DLPProfileListResponsePredefinedEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileListResponsePredefinedEntriesPredefinedJSON `json:"-"`
}

// dlpProfileListResponsePredefinedEntriesPredefinedJSON contains the JSON metadata
// for the struct [DLPProfileListResponsePredefinedEntriesPredefined]
type dlpProfileListResponsePredefinedEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponsePredefinedEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponsePredefinedEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponsePredefinedEntriesPredefined) implementsZeroTrustDLPProfileListResponsePredefinedEntry() {
}

type DLPProfileListResponsePredefinedEntriesPredefinedType string

const (
	DLPProfileListResponsePredefinedEntriesPredefinedTypePredefined DLPProfileListResponsePredefinedEntriesPredefinedType = "predefined"
)

func (r DLPProfileListResponsePredefinedEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileListResponsePredefinedEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileListResponsePredefinedEntriesIntegration struct {
	ID        string                                                 `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                              `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                   `json:"enabled,required"`
	Name      string                                                 `json:"name,required"`
	Type      DLPProfileListResponsePredefinedEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                              `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                 `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileListResponsePredefinedEntriesIntegrationJSON `json:"-"`
}

// dlpProfileListResponsePredefinedEntriesIntegrationJSON contains the JSON
// metadata for the struct [DLPProfileListResponsePredefinedEntriesIntegration]
type dlpProfileListResponsePredefinedEntriesIntegrationJSON struct {
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

func (r *DLPProfileListResponsePredefinedEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponsePredefinedEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponsePredefinedEntriesIntegration) implementsZeroTrustDLPProfileListResponsePredefinedEntry() {
}

type DLPProfileListResponsePredefinedEntriesIntegrationType string

const (
	DLPProfileListResponsePredefinedEntriesIntegrationTypeIntegration DLPProfileListResponsePredefinedEntriesIntegrationType = "integration"
)

func (r DLPProfileListResponsePredefinedEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileListResponsePredefinedEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileListResponsePredefinedEntriesExactData struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                 `json:"enabled,required"`
	Name      string                                               `json:"name,required"`
	Secret    bool                                                 `json:"secret,required"`
	Type      DLPProfileListResponsePredefinedEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileListResponsePredefinedEntriesExactDataJSON `json:"-"`
}

// dlpProfileListResponsePredefinedEntriesExactDataJSON contains the JSON metadata
// for the struct [DLPProfileListResponsePredefinedEntriesExactData]
type dlpProfileListResponsePredefinedEntriesExactDataJSON struct {
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

func (r *DLPProfileListResponsePredefinedEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponsePredefinedEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponsePredefinedEntriesExactData) implementsZeroTrustDLPProfileListResponsePredefinedEntry() {
}

type DLPProfileListResponsePredefinedEntriesExactDataType string

const (
	DLPProfileListResponsePredefinedEntriesExactDataTypeExactData DLPProfileListResponsePredefinedEntriesExactDataType = "exact_data"
)

func (r DLPProfileListResponsePredefinedEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileListResponsePredefinedEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileListResponsePredefinedEntriesWordList struct {
	ID        string                                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                `json:"enabled,required"`
	Name      string                                              `json:"name,required"`
	Type      DLPProfileListResponsePredefinedEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                           `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                         `json:"word_list,required"`
	ProfileID string                                              `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileListResponsePredefinedEntriesWordListJSON `json:"-"`
}

// dlpProfileListResponsePredefinedEntriesWordListJSON contains the JSON metadata
// for the struct [DLPProfileListResponsePredefinedEntriesWordList]
type dlpProfileListResponsePredefinedEntriesWordListJSON struct {
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

func (r *DLPProfileListResponsePredefinedEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponsePredefinedEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponsePredefinedEntriesWordList) implementsZeroTrustDLPProfileListResponsePredefinedEntry() {
}

type DLPProfileListResponsePredefinedEntriesWordListType string

const (
	DLPProfileListResponsePredefinedEntriesWordListTypeWordList DLPProfileListResponsePredefinedEntriesWordListType = "word_list"
)

func (r DLPProfileListResponsePredefinedEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileListResponsePredefinedEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileListResponsePredefinedEntriesType string

const (
	DLPProfileListResponsePredefinedEntriesTypeCustom      DLPProfileListResponsePredefinedEntriesType = "custom"
	DLPProfileListResponsePredefinedEntriesTypePredefined  DLPProfileListResponsePredefinedEntriesType = "predefined"
	DLPProfileListResponsePredefinedEntriesTypeIntegration DLPProfileListResponsePredefinedEntriesType = "integration"
	DLPProfileListResponsePredefinedEntriesTypeExactData   DLPProfileListResponsePredefinedEntriesType = "exact_data"
	DLPProfileListResponsePredefinedEntriesTypeWordList    DLPProfileListResponsePredefinedEntriesType = "word_list"
)

func (r DLPProfileListResponsePredefinedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileListResponsePredefinedEntriesTypeCustom, DLPProfileListResponsePredefinedEntriesTypePredefined, DLPProfileListResponsePredefinedEntriesTypeIntegration, DLPProfileListResponsePredefinedEntriesTypeExactData, DLPProfileListResponsePredefinedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileListResponsePredefinedType string

const (
	DLPProfileListResponsePredefinedTypePredefined DLPProfileListResponsePredefinedType = "predefined"
)

func (r DLPProfileListResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileListResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileListResponseIntegration struct {
	ID        string                                   `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileListResponseIntegrationEntry `json:"entries,required"`
	Name      string                                   `json:"name,required"`
	Type      DLPProfileListResponseIntegrationType    `json:"type,required"`
	UpdatedAt time.Time                                `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                                `json:"description,nullable"`
	JSON        dlpProfileListResponseIntegrationJSON `json:"-"`
}

// dlpProfileListResponseIntegrationJSON contains the JSON metadata for the struct
// [DLPProfileListResponseIntegration]
type dlpProfileListResponseIntegrationJSON struct {
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

func (r *DLPProfileListResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseIntegration) implementsZeroTrustDLPProfileListResponse() {}

type DLPProfileListResponseIntegrationEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileListResponseIntegrationEntriesCustomPattern].
	Pattern   interface{}                                  `json:"pattern,required"`
	ProfileID string                                       `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	Type      DLPProfileListResponseIntegrationEntriesType `json:"type,required"`
	Secret    bool                                         `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                `json:"word_list,required"`
	JSON     dlpProfileListResponseIntegrationEntryJSON `json:"-"`
	union    DLPProfileListResponseIntegrationEntriesUnion
}

// dlpProfileListResponseIntegrationEntryJSON contains the JSON metadata for the
// struct [DLPProfileListResponseIntegrationEntry]
type dlpProfileListResponseIntegrationEntryJSON struct {
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	Type        apijson.Field
	Secret      apijson.Field
	WordList    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dlpProfileListResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileListResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileListResponseIntegrationEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileListResponseIntegrationEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileListResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfileListResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfileListResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfileListResponseIntegrationEntriesExactData],
// [zero_trust.DLPProfileListResponseIntegrationEntriesWordList].
func (r DLPProfileListResponseIntegrationEntry) AsUnion() DLPProfileListResponseIntegrationEntriesUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfileListResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfileListResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfileListResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfileListResponseIntegrationEntriesExactData] or
// [zero_trust.DLPProfileListResponseIntegrationEntriesWordList].
type DLPProfileListResponseIntegrationEntriesUnion interface {
	implementsZeroTrustDLPProfileListResponseIntegrationEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileListResponseIntegrationEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponseIntegrationEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponseIntegrationEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponseIntegrationEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponseIntegrationEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileListResponseIntegrationEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileListResponseIntegrationEntriesCustom struct {
	ID        string                                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                  `json:"enabled,required"`
	Name      string                                                `json:"name,required"`
	Pattern   DLPProfileListResponseIntegrationEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileListResponseIntegrationEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                             `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileListResponseIntegrationEntriesCustomJSON    `json:"-"`
}

// dlpProfileListResponseIntegrationEntriesCustomJSON contains the JSON metadata
// for the struct [DLPProfileListResponseIntegrationEntriesCustom]
type dlpProfileListResponseIntegrationEntriesCustomJSON struct {
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

func (r *DLPProfileListResponseIntegrationEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseIntegrationEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseIntegrationEntriesCustom) implementsZeroTrustDLPProfileListResponseIntegrationEntry() {
}

type DLPProfileListResponseIntegrationEntriesCustomPattern struct {
	Regex      string                                                          `json:"regex,required"`
	Validation DLPProfileListResponseIntegrationEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileListResponseIntegrationEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileListResponseIntegrationEntriesCustomPatternJSON contains the JSON
// metadata for the struct [DLPProfileListResponseIntegrationEntriesCustomPattern]
type dlpProfileListResponseIntegrationEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseIntegrationEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseIntegrationEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileListResponseIntegrationEntriesCustomPatternValidation string

const (
	DLPProfileListResponseIntegrationEntriesCustomPatternValidationLuhn DLPProfileListResponseIntegrationEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileListResponseIntegrationEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileListResponseIntegrationEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileListResponseIntegrationEntriesCustomType string

const (
	DLPProfileListResponseIntegrationEntriesCustomTypeCustom DLPProfileListResponseIntegrationEntriesCustomType = "custom"
)

func (r DLPProfileListResponseIntegrationEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseIntegrationEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileListResponseIntegrationEntriesPredefined struct {
	ID        string                                                 `json:"id,required" format:"uuid"`
	Enabled   bool                                                   `json:"enabled,required"`
	Name      string                                                 `json:"name,required"`
	Type      DLPProfileListResponseIntegrationEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                 `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileListResponseIntegrationEntriesPredefinedJSON `json:"-"`
}

// dlpProfileListResponseIntegrationEntriesPredefinedJSON contains the JSON
// metadata for the struct [DLPProfileListResponseIntegrationEntriesPredefined]
type dlpProfileListResponseIntegrationEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseIntegrationEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseIntegrationEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseIntegrationEntriesPredefined) implementsZeroTrustDLPProfileListResponseIntegrationEntry() {
}

type DLPProfileListResponseIntegrationEntriesPredefinedType string

const (
	DLPProfileListResponseIntegrationEntriesPredefinedTypePredefined DLPProfileListResponseIntegrationEntriesPredefinedType = "predefined"
)

func (r DLPProfileListResponseIntegrationEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseIntegrationEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileListResponseIntegrationEntriesIntegration struct {
	ID        string                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                    `json:"enabled,required"`
	Name      string                                                  `json:"name,required"`
	Type      DLPProfileListResponseIntegrationEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                               `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileListResponseIntegrationEntriesIntegrationJSON `json:"-"`
}

// dlpProfileListResponseIntegrationEntriesIntegrationJSON contains the JSON
// metadata for the struct [DLPProfileListResponseIntegrationEntriesIntegration]
type dlpProfileListResponseIntegrationEntriesIntegrationJSON struct {
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

func (r *DLPProfileListResponseIntegrationEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseIntegrationEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseIntegrationEntriesIntegration) implementsZeroTrustDLPProfileListResponseIntegrationEntry() {
}

type DLPProfileListResponseIntegrationEntriesIntegrationType string

const (
	DLPProfileListResponseIntegrationEntriesIntegrationTypeIntegration DLPProfileListResponseIntegrationEntriesIntegrationType = "integration"
)

func (r DLPProfileListResponseIntegrationEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseIntegrationEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileListResponseIntegrationEntriesExactData struct {
	ID        string                                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                  `json:"enabled,required"`
	Name      string                                                `json:"name,required"`
	Secret    bool                                                  `json:"secret,required"`
	Type      DLPProfileListResponseIntegrationEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                             `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileListResponseIntegrationEntriesExactDataJSON `json:"-"`
}

// dlpProfileListResponseIntegrationEntriesExactDataJSON contains the JSON metadata
// for the struct [DLPProfileListResponseIntegrationEntriesExactData]
type dlpProfileListResponseIntegrationEntriesExactDataJSON struct {
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

func (r *DLPProfileListResponseIntegrationEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseIntegrationEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseIntegrationEntriesExactData) implementsZeroTrustDLPProfileListResponseIntegrationEntry() {
}

type DLPProfileListResponseIntegrationEntriesExactDataType string

const (
	DLPProfileListResponseIntegrationEntriesExactDataTypeExactData DLPProfileListResponseIntegrationEntriesExactDataType = "exact_data"
)

func (r DLPProfileListResponseIntegrationEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseIntegrationEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileListResponseIntegrationEntriesWordList struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                 `json:"enabled,required"`
	Name      string                                               `json:"name,required"`
	Type      DLPProfileListResponseIntegrationEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                          `json:"word_list,required"`
	ProfileID string                                               `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileListResponseIntegrationEntriesWordListJSON `json:"-"`
}

// dlpProfileListResponseIntegrationEntriesWordListJSON contains the JSON metadata
// for the struct [DLPProfileListResponseIntegrationEntriesWordList]
type dlpProfileListResponseIntegrationEntriesWordListJSON struct {
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

func (r *DLPProfileListResponseIntegrationEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseIntegrationEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseIntegrationEntriesWordList) implementsZeroTrustDLPProfileListResponseIntegrationEntry() {
}

type DLPProfileListResponseIntegrationEntriesWordListType string

const (
	DLPProfileListResponseIntegrationEntriesWordListTypeWordList DLPProfileListResponseIntegrationEntriesWordListType = "word_list"
)

func (r DLPProfileListResponseIntegrationEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseIntegrationEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileListResponseIntegrationEntriesType string

const (
	DLPProfileListResponseIntegrationEntriesTypeCustom      DLPProfileListResponseIntegrationEntriesType = "custom"
	DLPProfileListResponseIntegrationEntriesTypePredefined  DLPProfileListResponseIntegrationEntriesType = "predefined"
	DLPProfileListResponseIntegrationEntriesTypeIntegration DLPProfileListResponseIntegrationEntriesType = "integration"
	DLPProfileListResponseIntegrationEntriesTypeExactData   DLPProfileListResponseIntegrationEntriesType = "exact_data"
	DLPProfileListResponseIntegrationEntriesTypeWordList    DLPProfileListResponseIntegrationEntriesType = "word_list"
)

func (r DLPProfileListResponseIntegrationEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseIntegrationEntriesTypeCustom, DLPProfileListResponseIntegrationEntriesTypePredefined, DLPProfileListResponseIntegrationEntriesTypeIntegration, DLPProfileListResponseIntegrationEntriesTypeExactData, DLPProfileListResponseIntegrationEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileListResponseIntegrationType string

const (
	DLPProfileListResponseIntegrationTypeIntegration DLPProfileListResponseIntegrationType = "integration"
)

func (r DLPProfileListResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileListResponseType string

const (
	DLPProfileListResponseTypeCustom      DLPProfileListResponseType = "custom"
	DLPProfileListResponseTypePredefined  DLPProfileListResponseType = "predefined"
	DLPProfileListResponseTypeIntegration DLPProfileListResponseType = "integration"
)

func (r DLPProfileListResponseType) IsKnown() bool {
	switch r {
	case DLPProfileListResponseTypeCustom, DLPProfileListResponseTypePredefined, DLPProfileListResponseTypeIntegration:
		return true
	}
	return false
}

type DLPProfileGetResponse struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// When the profile was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile
	Description string `json:"description,nullable"`
	// This field can have the runtime type of [[]DLPProfileGetResponseCustomEntry],
	// [[]DLPProfileGetResponsePredefinedEntry],
	// [[]DLPProfileGetResponseIntegrationEntry].
	Entries interface{} `json:"entries"`
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// The name of the profile
	Name       string `json:"name,required"`
	OCREnabled bool   `json:"ocr_enabled"`
	// When the profile was lasted updated
	UpdatedAt  time.Time                 `json:"updated_at" format:"date-time"`
	Type       DLPProfileGetResponseType `json:"type,required"`
	OpenAccess bool                      `json:"open_access"`
	JSON       dlpProfileGetResponseJSON `json:"-"`
	union      DLPProfileGetResponseUnion
}

// dlpProfileGetResponseJSON contains the JSON metadata for the struct
// [DLPProfileGetResponse]
type dlpProfileGetResponseJSON struct {
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	CreatedAt         apijson.Field
	Description       apijson.Field
	Entries           apijson.Field
	ID                apijson.Field
	Name              apijson.Field
	OCREnabled        apijson.Field
	UpdatedAt         apijson.Field
	Type              apijson.Field
	OpenAccess        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r dlpProfileGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileGetResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileGetResponseCustom],
// [zero_trust.DLPProfileGetResponsePredefined],
// [zero_trust.DLPProfileGetResponseIntegration].
func (r DLPProfileGetResponse) AsUnion() DLPProfileGetResponseUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfileGetResponseCustom],
// [zero_trust.DLPProfileGetResponsePredefined] or
// [zero_trust.DLPProfileGetResponseIntegration].
type DLPProfileGetResponseUnion interface {
	implementsZeroTrustDLPProfileGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
	)
}

type DLPProfileGetResponseCustom struct {
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness,required"`
	// When the profile was created
	CreatedAt time.Time                          `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileGetResponseCustomEntry `json:"entries,required"`
	// The name of the profile
	Name       string                          `json:"name,required"`
	OCREnabled bool                            `json:"ocr_enabled,required"`
	Type       DLPProfileGetResponseCustomType `json:"type,required"`
	// When the profile was lasted updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                          `json:"description,nullable"`
	JSON        dlpProfileGetResponseCustomJSON `json:"-"`
}

// dlpProfileGetResponseCustomJSON contains the JSON metadata for the struct
// [DLPProfileGetResponseCustom]
type dlpProfileGetResponseCustomJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	CreatedAt         apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	OCREnabled        apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	Description       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustom) implementsZeroTrustDLPProfileGetResponse() {}

type DLPProfileGetResponseCustomEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileGetResponseCustomEntriesCustomPattern].
	Pattern   interface{}                            `json:"pattern,required"`
	ProfileID string                                 `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                              `json:"updated_at" format:"date-time"`
	Type      DLPProfileGetResponseCustomEntriesType `json:"type,required"`
	Secret    bool                                   `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                          `json:"word_list,required"`
	JSON     dlpProfileGetResponseCustomEntryJSON `json:"-"`
	union    DLPProfileGetResponseCustomEntriesUnion
}

// dlpProfileGetResponseCustomEntryJSON contains the JSON metadata for the struct
// [DLPProfileGetResponseCustomEntry]
type dlpProfileGetResponseCustomEntryJSON struct {
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	Type        apijson.Field
	Secret      apijson.Field
	WordList    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dlpProfileGetResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileGetResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileGetResponseCustomEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileGetResponseCustomEntriesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileGetResponseCustomEntriesCustom],
// [zero_trust.DLPProfileGetResponseCustomEntriesPredefined],
// [zero_trust.DLPProfileGetResponseCustomEntriesIntegration],
// [zero_trust.DLPProfileGetResponseCustomEntriesExactData],
// [zero_trust.DLPProfileGetResponseCustomEntriesWordList].
func (r DLPProfileGetResponseCustomEntry) AsUnion() DLPProfileGetResponseCustomEntriesUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfileGetResponseCustomEntriesCustom],
// [zero_trust.DLPProfileGetResponseCustomEntriesPredefined],
// [zero_trust.DLPProfileGetResponseCustomEntriesIntegration],
// [zero_trust.DLPProfileGetResponseCustomEntriesExactData] or
// [zero_trust.DLPProfileGetResponseCustomEntriesWordList].
type DLPProfileGetResponseCustomEntriesUnion interface {
	implementsZeroTrustDLPProfileGetResponseCustomEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileGetResponseCustomEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseCustomEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseCustomEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseCustomEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseCustomEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseCustomEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileGetResponseCustomEntriesCustom struct {
	ID        string                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                            `json:"enabled,required"`
	Name      string                                          `json:"name,required"`
	Pattern   DLPProfileGetResponseCustomEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileGetResponseCustomEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                       `json:"updated_at,required" format:"date-time"`
	ProfileID string                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseCustomEntriesCustomJSON    `json:"-"`
}

// dlpProfileGetResponseCustomEntriesCustomJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseCustomEntriesCustom]
type dlpProfileGetResponseCustomEntriesCustomJSON struct {
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

func (r *DLPProfileGetResponseCustomEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomEntriesCustom) implementsZeroTrustDLPProfileGetResponseCustomEntry() {
}

type DLPProfileGetResponseCustomEntriesCustomPattern struct {
	Regex      string                                                    `json:"regex,required"`
	Validation DLPProfileGetResponseCustomEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileGetResponseCustomEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileGetResponseCustomEntriesCustomPatternJSON contains the JSON metadata
// for the struct [DLPProfileGetResponseCustomEntriesCustomPattern]
type dlpProfileGetResponseCustomEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseCustomEntriesCustomPatternValidation string

const (
	DLPProfileGetResponseCustomEntriesCustomPatternValidationLuhn DLPProfileGetResponseCustomEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileGetResponseCustomEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesCustomType string

const (
	DLPProfileGetResponseCustomEntriesCustomTypeCustom DLPProfileGetResponseCustomEntriesCustomType = "custom"
)

func (r DLPProfileGetResponseCustomEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesPredefined struct {
	ID        string                                           `json:"id,required" format:"uuid"`
	Enabled   bool                                             `json:"enabled,required"`
	Name      string                                           `json:"name,required"`
	Type      DLPProfileGetResponseCustomEntriesPredefinedType `json:"type,required"`
	ProfileID string                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseCustomEntriesPredefinedJSON `json:"-"`
}

// dlpProfileGetResponseCustomEntriesPredefinedJSON contains the JSON metadata for
// the struct [DLPProfileGetResponseCustomEntriesPredefined]
type dlpProfileGetResponseCustomEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomEntriesPredefined) implementsZeroTrustDLPProfileGetResponseCustomEntry() {
}

type DLPProfileGetResponseCustomEntriesPredefinedType string

const (
	DLPProfileGetResponseCustomEntriesPredefinedTypePredefined DLPProfileGetResponseCustomEntriesPredefinedType = "predefined"
)

func (r DLPProfileGetResponseCustomEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesIntegration struct {
	ID        string                                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                         `json:"created_at,required" format:"date-time"`
	Enabled   bool                                              `json:"enabled,required"`
	Name      string                                            `json:"name,required"`
	Type      DLPProfileGetResponseCustomEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                         `json:"updated_at,required" format:"date-time"`
	ProfileID string                                            `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseCustomEntriesIntegrationJSON `json:"-"`
}

// dlpProfileGetResponseCustomEntriesIntegrationJSON contains the JSON metadata for
// the struct [DLPProfileGetResponseCustomEntriesIntegration]
type dlpProfileGetResponseCustomEntriesIntegrationJSON struct {
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

func (r *DLPProfileGetResponseCustomEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomEntriesIntegration) implementsZeroTrustDLPProfileGetResponseCustomEntry() {
}

type DLPProfileGetResponseCustomEntriesIntegrationType string

const (
	DLPProfileGetResponseCustomEntriesIntegrationTypeIntegration DLPProfileGetResponseCustomEntriesIntegrationType = "integration"
)

func (r DLPProfileGetResponseCustomEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesExactData struct {
	ID        string                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                            `json:"enabled,required"`
	Name      string                                          `json:"name,required"`
	Secret    bool                                            `json:"secret,required"`
	Type      DLPProfileGetResponseCustomEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                       `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileGetResponseCustomEntriesExactDataJSON `json:"-"`
}

// dlpProfileGetResponseCustomEntriesExactDataJSON contains the JSON metadata for
// the struct [DLPProfileGetResponseCustomEntriesExactData]
type dlpProfileGetResponseCustomEntriesExactDataJSON struct {
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

func (r *DLPProfileGetResponseCustomEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomEntriesExactData) implementsZeroTrustDLPProfileGetResponseCustomEntry() {
}

type DLPProfileGetResponseCustomEntriesExactDataType string

const (
	DLPProfileGetResponseCustomEntriesExactDataTypeExactData DLPProfileGetResponseCustomEntriesExactDataType = "exact_data"
)

func (r DLPProfileGetResponseCustomEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesWordList struct {
	ID        string                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                           `json:"enabled,required"`
	Name      string                                         `json:"name,required"`
	Type      DLPProfileGetResponseCustomEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                      `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                    `json:"word_list,required"`
	ProfileID string                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseCustomEntriesWordListJSON `json:"-"`
}

// dlpProfileGetResponseCustomEntriesWordListJSON contains the JSON metadata for
// the struct [DLPProfileGetResponseCustomEntriesWordList]
type dlpProfileGetResponseCustomEntriesWordListJSON struct {
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

func (r *DLPProfileGetResponseCustomEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomEntriesWordList) implementsZeroTrustDLPProfileGetResponseCustomEntry() {
}

type DLPProfileGetResponseCustomEntriesWordListType string

const (
	DLPProfileGetResponseCustomEntriesWordListTypeWordList DLPProfileGetResponseCustomEntriesWordListType = "word_list"
)

func (r DLPProfileGetResponseCustomEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesType string

const (
	DLPProfileGetResponseCustomEntriesTypeCustom      DLPProfileGetResponseCustomEntriesType = "custom"
	DLPProfileGetResponseCustomEntriesTypePredefined  DLPProfileGetResponseCustomEntriesType = "predefined"
	DLPProfileGetResponseCustomEntriesTypeIntegration DLPProfileGetResponseCustomEntriesType = "integration"
	DLPProfileGetResponseCustomEntriesTypeExactData   DLPProfileGetResponseCustomEntriesType = "exact_data"
	DLPProfileGetResponseCustomEntriesTypeWordList    DLPProfileGetResponseCustomEntriesType = "word_list"
)

func (r DLPProfileGetResponseCustomEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesTypeCustom, DLPProfileGetResponseCustomEntriesTypePredefined, DLPProfileGetResponseCustomEntriesTypeIntegration, DLPProfileGetResponseCustomEntriesTypeExactData, DLPProfileGetResponseCustomEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomType string

const (
	DLPProfileGetResponseCustomTypeCustom DLPProfileGetResponseCustomType = "custom"
)

func (r DLPProfileGetResponseCustomType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefined struct {
	// The id of the predefined profile (uuid)
	ID                string                                 `json:"id,required" format:"uuid"`
	AllowedMatchCount int64                                  `json:"allowed_match_count,required"`
	Entries           []DLPProfileGetResponsePredefinedEntry `json:"entries,required"`
	// The name of the predefined profile
	Name string                              `json:"name,required"`
	Type DLPProfileGetResponsePredefinedType `json:"type,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness                    `json:"context_awareness"`
	OCREnabled       bool                                `json:"ocr_enabled"`
	OpenAccess       bool                                `json:"open_access"`
	JSON             dlpProfileGetResponsePredefinedJSON `json:"-"`
}

// dlpProfileGetResponsePredefinedJSON contains the JSON metadata for the struct
// [DLPProfileGetResponsePredefined]
type dlpProfileGetResponsePredefinedJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	ContextAwareness  apijson.Field
	OCREnabled        apijson.Field
	OpenAccess        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPProfileGetResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponsePredefined) implementsZeroTrustDLPProfileGetResponse() {}

type DLPProfileGetResponsePredefinedEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileGetResponsePredefinedEntriesCustomPattern].
	Pattern   interface{}                                `json:"pattern,required"`
	ProfileID string                                     `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	Type      DLPProfileGetResponsePredefinedEntriesType `json:"type,required"`
	Secret    bool                                       `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                              `json:"word_list,required"`
	JSON     dlpProfileGetResponsePredefinedEntryJSON `json:"-"`
	union    DLPProfileGetResponsePredefinedEntriesUnion
}

// dlpProfileGetResponsePredefinedEntryJSON contains the JSON metadata for the
// struct [DLPProfileGetResponsePredefinedEntry]
type dlpProfileGetResponsePredefinedEntryJSON struct {
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	Type        apijson.Field
	Secret      apijson.Field
	WordList    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dlpProfileGetResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileGetResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileGetResponsePredefinedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileGetResponsePredefinedEntriesUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileGetResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfileGetResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfileGetResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfileGetResponsePredefinedEntriesExactData],
// [zero_trust.DLPProfileGetResponsePredefinedEntriesWordList].
func (r DLPProfileGetResponsePredefinedEntry) AsUnion() DLPProfileGetResponsePredefinedEntriesUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfileGetResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfileGetResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfileGetResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfileGetResponsePredefinedEntriesExactData] or
// [zero_trust.DLPProfileGetResponsePredefinedEntriesWordList].
type DLPProfileGetResponsePredefinedEntriesUnion interface {
	implementsZeroTrustDLPProfileGetResponsePredefinedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileGetResponsePredefinedEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponsePredefinedEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponsePredefinedEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponsePredefinedEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponsePredefinedEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponsePredefinedEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileGetResponsePredefinedEntriesCustom struct {
	ID        string                                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                `json:"enabled,required"`
	Name      string                                              `json:"name,required"`
	Pattern   DLPProfileGetResponsePredefinedEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileGetResponsePredefinedEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                           `json:"updated_at,required" format:"date-time"`
	ProfileID string                                              `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponsePredefinedEntriesCustomJSON    `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesCustomJSON contains the JSON metadata for
// the struct [DLPProfileGetResponsePredefinedEntriesCustom]
type dlpProfileGetResponsePredefinedEntriesCustomJSON struct {
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

func (r *DLPProfileGetResponsePredefinedEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponsePredefinedEntriesCustom) implementsZeroTrustDLPProfileGetResponsePredefinedEntry() {
}

type DLPProfileGetResponsePredefinedEntriesCustomPattern struct {
	Regex      string                                                        `json:"regex,required"`
	Validation DLPProfileGetResponsePredefinedEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileGetResponsePredefinedEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesCustomPatternJSON contains the JSON
// metadata for the struct [DLPProfileGetResponsePredefinedEntriesCustomPattern]
type dlpProfileGetResponsePredefinedEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponsePredefinedEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponsePredefinedEntriesCustomPatternValidation string

const (
	DLPProfileGetResponsePredefinedEntriesCustomPatternValidationLuhn DLPProfileGetResponsePredefinedEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileGetResponsePredefinedEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesCustomType string

const (
	DLPProfileGetResponsePredefinedEntriesCustomTypeCustom DLPProfileGetResponsePredefinedEntriesCustomType = "custom"
)

func (r DLPProfileGetResponsePredefinedEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesPredefined struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	Enabled   bool                                                 `json:"enabled,required"`
	Name      string                                               `json:"name,required"`
	Type      DLPProfileGetResponsePredefinedEntriesPredefinedType `json:"type,required"`
	ProfileID string                                               `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponsePredefinedEntriesPredefinedJSON `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesPredefinedJSON contains the JSON metadata
// for the struct [DLPProfileGetResponsePredefinedEntriesPredefined]
type dlpProfileGetResponsePredefinedEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponsePredefinedEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponsePredefinedEntriesPredefined) implementsZeroTrustDLPProfileGetResponsePredefinedEntry() {
}

type DLPProfileGetResponsePredefinedEntriesPredefinedType string

const (
	DLPProfileGetResponsePredefinedEntriesPredefinedTypePredefined DLPProfileGetResponsePredefinedEntriesPredefinedType = "predefined"
)

func (r DLPProfileGetResponsePredefinedEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesIntegration struct {
	ID        string                                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                  `json:"enabled,required"`
	Name      string                                                `json:"name,required"`
	Type      DLPProfileGetResponsePredefinedEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                             `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponsePredefinedEntriesIntegrationJSON `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesIntegrationJSON contains the JSON metadata
// for the struct [DLPProfileGetResponsePredefinedEntriesIntegration]
type dlpProfileGetResponsePredefinedEntriesIntegrationJSON struct {
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

func (r *DLPProfileGetResponsePredefinedEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponsePredefinedEntriesIntegration) implementsZeroTrustDLPProfileGetResponsePredefinedEntry() {
}

type DLPProfileGetResponsePredefinedEntriesIntegrationType string

const (
	DLPProfileGetResponsePredefinedEntriesIntegrationTypeIntegration DLPProfileGetResponsePredefinedEntriesIntegrationType = "integration"
)

func (r DLPProfileGetResponsePredefinedEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesExactData struct {
	ID        string                                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                `json:"enabled,required"`
	Name      string                                              `json:"name,required"`
	Secret    bool                                                `json:"secret,required"`
	Type      DLPProfileGetResponsePredefinedEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                           `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileGetResponsePredefinedEntriesExactDataJSON `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesExactDataJSON contains the JSON metadata
// for the struct [DLPProfileGetResponsePredefinedEntriesExactData]
type dlpProfileGetResponsePredefinedEntriesExactDataJSON struct {
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

func (r *DLPProfileGetResponsePredefinedEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponsePredefinedEntriesExactData) implementsZeroTrustDLPProfileGetResponsePredefinedEntry() {
}

type DLPProfileGetResponsePredefinedEntriesExactDataType string

const (
	DLPProfileGetResponsePredefinedEntriesExactDataTypeExactData DLPProfileGetResponsePredefinedEntriesExactDataType = "exact_data"
)

func (r DLPProfileGetResponsePredefinedEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesWordList struct {
	ID        string                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                               `json:"enabled,required"`
	Name      string                                             `json:"name,required"`
	Type      DLPProfileGetResponsePredefinedEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                          `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                        `json:"word_list,required"`
	ProfileID string                                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponsePredefinedEntriesWordListJSON `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesWordListJSON contains the JSON metadata
// for the struct [DLPProfileGetResponsePredefinedEntriesWordList]
type dlpProfileGetResponsePredefinedEntriesWordListJSON struct {
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

func (r *DLPProfileGetResponsePredefinedEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponsePredefinedEntriesWordList) implementsZeroTrustDLPProfileGetResponsePredefinedEntry() {
}

type DLPProfileGetResponsePredefinedEntriesWordListType string

const (
	DLPProfileGetResponsePredefinedEntriesWordListTypeWordList DLPProfileGetResponsePredefinedEntriesWordListType = "word_list"
)

func (r DLPProfileGetResponsePredefinedEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesType string

const (
	DLPProfileGetResponsePredefinedEntriesTypeCustom      DLPProfileGetResponsePredefinedEntriesType = "custom"
	DLPProfileGetResponsePredefinedEntriesTypePredefined  DLPProfileGetResponsePredefinedEntriesType = "predefined"
	DLPProfileGetResponsePredefinedEntriesTypeIntegration DLPProfileGetResponsePredefinedEntriesType = "integration"
	DLPProfileGetResponsePredefinedEntriesTypeExactData   DLPProfileGetResponsePredefinedEntriesType = "exact_data"
	DLPProfileGetResponsePredefinedEntriesTypeWordList    DLPProfileGetResponsePredefinedEntriesType = "word_list"
)

func (r DLPProfileGetResponsePredefinedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesTypeCustom, DLPProfileGetResponsePredefinedEntriesTypePredefined, DLPProfileGetResponsePredefinedEntriesTypeIntegration, DLPProfileGetResponsePredefinedEntriesTypeExactData, DLPProfileGetResponsePredefinedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedType string

const (
	DLPProfileGetResponsePredefinedTypePredefined DLPProfileGetResponsePredefinedType = "predefined"
)

func (r DLPProfileGetResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegration struct {
	ID        string                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                               `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileGetResponseIntegrationEntry `json:"entries,required"`
	Name      string                                  `json:"name,required"`
	Type      DLPProfileGetResponseIntegrationType    `json:"type,required"`
	UpdatedAt time.Time                               `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                               `json:"description,nullable"`
	JSON        dlpProfileGetResponseIntegrationJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationJSON contains the JSON metadata for the struct
// [DLPProfileGetResponseIntegration]
type dlpProfileGetResponseIntegrationJSON struct {
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

func (r *DLPProfileGetResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegration) implementsZeroTrustDLPProfileGetResponse() {}

type DLPProfileGetResponseIntegrationEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileGetResponseIntegrationEntriesCustomPattern].
	Pattern   interface{}                                 `json:"pattern,required"`
	ProfileID string                                      `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	Type      DLPProfileGetResponseIntegrationEntriesType `json:"type,required"`
	Secret    bool                                        `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                               `json:"word_list,required"`
	JSON     dlpProfileGetResponseIntegrationEntryJSON `json:"-"`
	union    DLPProfileGetResponseIntegrationEntriesUnion
}

// dlpProfileGetResponseIntegrationEntryJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseIntegrationEntry]
type dlpProfileGetResponseIntegrationEntryJSON struct {
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	Type        apijson.Field
	Secret      apijson.Field
	WordList    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dlpProfileGetResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileGetResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileGetResponseIntegrationEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileGetResponseIntegrationEntriesUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileGetResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfileGetResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfileGetResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfileGetResponseIntegrationEntriesExactData],
// [zero_trust.DLPProfileGetResponseIntegrationEntriesWordList].
func (r DLPProfileGetResponseIntegrationEntry) AsUnion() DLPProfileGetResponseIntegrationEntriesUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfileGetResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfileGetResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfileGetResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfileGetResponseIntegrationEntriesExactData] or
// [zero_trust.DLPProfileGetResponseIntegrationEntriesWordList].
type DLPProfileGetResponseIntegrationEntriesUnion interface {
	implementsZeroTrustDLPProfileGetResponseIntegrationEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileGetResponseIntegrationEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseIntegrationEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseIntegrationEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseIntegrationEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseIntegrationEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseIntegrationEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileGetResponseIntegrationEntriesCustom struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                 `json:"enabled,required"`
	Name      string                                               `json:"name,required"`
	Pattern   DLPProfileGetResponseIntegrationEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileGetResponseIntegrationEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	ProfileID string                                               `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseIntegrationEntriesCustomJSON    `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesCustomJSON contains the JSON metadata for
// the struct [DLPProfileGetResponseIntegrationEntriesCustom]
type dlpProfileGetResponseIntegrationEntriesCustomJSON struct {
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

func (r *DLPProfileGetResponseIntegrationEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationEntriesCustom) implementsZeroTrustDLPProfileGetResponseIntegrationEntry() {
}

type DLPProfileGetResponseIntegrationEntriesCustomPattern struct {
	Regex      string                                                         `json:"regex,required"`
	Validation DLPProfileGetResponseIntegrationEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileGetResponseIntegrationEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesCustomPatternJSON contains the JSON
// metadata for the struct [DLPProfileGetResponseIntegrationEntriesCustomPattern]
type dlpProfileGetResponseIntegrationEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseIntegrationEntriesCustomPatternValidation string

const (
	DLPProfileGetResponseIntegrationEntriesCustomPatternValidationLuhn DLPProfileGetResponseIntegrationEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileGetResponseIntegrationEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesCustomType string

const (
	DLPProfileGetResponseIntegrationEntriesCustomTypeCustom DLPProfileGetResponseIntegrationEntriesCustomType = "custom"
)

func (r DLPProfileGetResponseIntegrationEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesPredefined struct {
	ID        string                                                `json:"id,required" format:"uuid"`
	Enabled   bool                                                  `json:"enabled,required"`
	Name      string                                                `json:"name,required"`
	Type      DLPProfileGetResponseIntegrationEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseIntegrationEntriesPredefinedJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesPredefinedJSON contains the JSON metadata
// for the struct [DLPProfileGetResponseIntegrationEntriesPredefined]
type dlpProfileGetResponseIntegrationEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationEntriesPredefined) implementsZeroTrustDLPProfileGetResponseIntegrationEntry() {
}

type DLPProfileGetResponseIntegrationEntriesPredefinedType string

const (
	DLPProfileGetResponseIntegrationEntriesPredefinedTypePredefined DLPProfileGetResponseIntegrationEntriesPredefinedType = "predefined"
)

func (r DLPProfileGetResponseIntegrationEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesIntegration struct {
	ID        string                                                 `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                              `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                   `json:"enabled,required"`
	Name      string                                                 `json:"name,required"`
	Type      DLPProfileGetResponseIntegrationEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                              `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                 `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseIntegrationEntriesIntegrationJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesIntegrationJSON contains the JSON
// metadata for the struct [DLPProfileGetResponseIntegrationEntriesIntegration]
type dlpProfileGetResponseIntegrationEntriesIntegrationJSON struct {
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

func (r *DLPProfileGetResponseIntegrationEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationEntriesIntegration) implementsZeroTrustDLPProfileGetResponseIntegrationEntry() {
}

type DLPProfileGetResponseIntegrationEntriesIntegrationType string

const (
	DLPProfileGetResponseIntegrationEntriesIntegrationTypeIntegration DLPProfileGetResponseIntegrationEntriesIntegrationType = "integration"
)

func (r DLPProfileGetResponseIntegrationEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesExactData struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                 `json:"enabled,required"`
	Name      string                                               `json:"name,required"`
	Secret    bool                                                 `json:"secret,required"`
	Type      DLPProfileGetResponseIntegrationEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileGetResponseIntegrationEntriesExactDataJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesExactDataJSON contains the JSON metadata
// for the struct [DLPProfileGetResponseIntegrationEntriesExactData]
type dlpProfileGetResponseIntegrationEntriesExactDataJSON struct {
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

func (r *DLPProfileGetResponseIntegrationEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationEntriesExactData) implementsZeroTrustDLPProfileGetResponseIntegrationEntry() {
}

type DLPProfileGetResponseIntegrationEntriesExactDataType string

const (
	DLPProfileGetResponseIntegrationEntriesExactDataTypeExactData DLPProfileGetResponseIntegrationEntriesExactDataType = "exact_data"
)

func (r DLPProfileGetResponseIntegrationEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesWordList struct {
	ID        string                                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                `json:"enabled,required"`
	Name      string                                              `json:"name,required"`
	Type      DLPProfileGetResponseIntegrationEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                           `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                         `json:"word_list,required"`
	ProfileID string                                              `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseIntegrationEntriesWordListJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesWordListJSON contains the JSON metadata
// for the struct [DLPProfileGetResponseIntegrationEntriesWordList]
type dlpProfileGetResponseIntegrationEntriesWordListJSON struct {
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

func (r *DLPProfileGetResponseIntegrationEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationEntriesWordList) implementsZeroTrustDLPProfileGetResponseIntegrationEntry() {
}

type DLPProfileGetResponseIntegrationEntriesWordListType string

const (
	DLPProfileGetResponseIntegrationEntriesWordListTypeWordList DLPProfileGetResponseIntegrationEntriesWordListType = "word_list"
)

func (r DLPProfileGetResponseIntegrationEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesType string

const (
	DLPProfileGetResponseIntegrationEntriesTypeCustom      DLPProfileGetResponseIntegrationEntriesType = "custom"
	DLPProfileGetResponseIntegrationEntriesTypePredefined  DLPProfileGetResponseIntegrationEntriesType = "predefined"
	DLPProfileGetResponseIntegrationEntriesTypeIntegration DLPProfileGetResponseIntegrationEntriesType = "integration"
	DLPProfileGetResponseIntegrationEntriesTypeExactData   DLPProfileGetResponseIntegrationEntriesType = "exact_data"
	DLPProfileGetResponseIntegrationEntriesTypeWordList    DLPProfileGetResponseIntegrationEntriesType = "word_list"
)

func (r DLPProfileGetResponseIntegrationEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesTypeCustom, DLPProfileGetResponseIntegrationEntriesTypePredefined, DLPProfileGetResponseIntegrationEntriesTypeIntegration, DLPProfileGetResponseIntegrationEntriesTypeExactData, DLPProfileGetResponseIntegrationEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationType string

const (
	DLPProfileGetResponseIntegrationTypeIntegration DLPProfileGetResponseIntegrationType = "integration"
)

func (r DLPProfileGetResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileGetResponseType string

const (
	DLPProfileGetResponseTypeCustom      DLPProfileGetResponseType = "custom"
	DLPProfileGetResponseTypePredefined  DLPProfileGetResponseType = "predefined"
	DLPProfileGetResponseTypeIntegration DLPProfileGetResponseType = "integration"
)

func (r DLPProfileGetResponseType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseTypeCustom, DLPProfileGetResponseTypePredefined, DLPProfileGetResponseTypeIntegration:
		return true
	}
	return false
}

type DLPProfileListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileGetResponseEnvelope struct {
	Errors     []shared.ResponseInfo                   `json:"errors,required"`
	Messages   []shared.ResponseInfo                   `json:"messages,required"`
	Success    bool                                    `json:"success,required"`
	Result     DLPProfileGetResponse                   `json:"result"`
	ResultInfo DLPProfileGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpProfileGetResponseEnvelopeJSON       `json:"-"`
}

// dlpProfileGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPProfileGetResponseEnvelope]
type dlpProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                       `json:"total_count,required"`
	JSON       dlpProfileGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpProfileGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseEnvelopeResultInfo]
type dlpProfileGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
