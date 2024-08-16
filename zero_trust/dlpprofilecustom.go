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
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
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

// Creates a set of DLP custom profiles.
func (r *DLPProfileCustomService) New(ctx context.Context, body DLPProfileCustomNewParams, opts ...option.RequestOption) (res *[]DLPProfileCustomNewResponse, err error) {
	var env DLPProfileCustomNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP custom profile.
func (r *DLPProfileCustomService) Update(ctx context.Context, profileID string, body DLPProfileCustomUpdateParams, opts ...option.RequestOption) (res *DLPProfileCustomUpdateResponse, err error) {
	var env DLPProfileCustomUpdateResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
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
func (r *DLPProfileCustomService) Get(ctx context.Context, profileID string, query DLPProfileCustomGetParams, opts ...option.RequestOption) (res *DLPProfileCustomGetResponse, err error) {
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

type DLPProfileCustomNewResponse struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// When the profile was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile
	Description string `json:"description,nullable"`
	// This field can have the runtime type of
	// [[]DLPProfileCustomNewResponseCustomEntry],
	// [[]DLPProfileCustomNewResponsePredefinedEntry],
	// [[]DLPProfileCustomNewResponseIntegrationEntry].
	Entries interface{} `json:"entries"`
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// The name of the profile
	Name       string `json:"name,required"`
	OCREnabled bool   `json:"ocr_enabled"`
	// When the profile was lasted updated
	UpdatedAt  time.Time                       `json:"updated_at" format:"date-time"`
	Type       DLPProfileCustomNewResponseType `json:"type,required"`
	OpenAccess bool                            `json:"open_access"`
	JSON       dlpProfileCustomNewResponseJSON `json:"-"`
	union      DLPProfileCustomNewResponseUnion
}

// dlpProfileCustomNewResponseJSON contains the JSON metadata for the struct
// [DLPProfileCustomNewResponse]
type dlpProfileCustomNewResponseJSON struct {
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

func (r dlpProfileCustomNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomNewResponseCustom],
// [zero_trust.DLPProfileCustomNewResponsePredefined],
// [zero_trust.DLPProfileCustomNewResponseIntegration].
func (r DLPProfileCustomNewResponse) AsUnion() DLPProfileCustomNewResponseUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfileCustomNewResponseCustom],
// [zero_trust.DLPProfileCustomNewResponsePredefined] or
// [zero_trust.DLPProfileCustomNewResponseIntegration].
type DLPProfileCustomNewResponseUnion interface {
	implementsZeroTrustDLPProfileCustomNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
	)
}

type DLPProfileCustomNewResponseCustom struct {
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness,required"`
	// When the profile was created
	CreatedAt time.Time                                `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileCustomNewResponseCustomEntry `json:"entries,required"`
	// The name of the profile
	Name       string                                `json:"name,required"`
	OCREnabled bool                                  `json:"ocr_enabled,required"`
	Type       DLPProfileCustomNewResponseCustomType `json:"type,required"`
	// When the profile was lasted updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                                `json:"description,nullable"`
	JSON        dlpProfileCustomNewResponseCustomJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomJSON contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseCustom]
type dlpProfileCustomNewResponseCustomJSON struct {
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

func (r *DLPProfileCustomNewResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustom) implementsZeroTrustDLPProfileCustomNewResponse() {}

type DLPProfileCustomNewResponseCustomEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseCustomEntriesCustomPattern].
	Pattern   interface{}                                  `json:"pattern,required"`
	ProfileID string                                       `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	Type      DLPProfileCustomNewResponseCustomEntriesType `json:"type,required"`
	Secret    bool                                         `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                `json:"word_list,required"`
	JSON     dlpProfileCustomNewResponseCustomEntryJSON `json:"-"`
	union    DLPProfileCustomNewResponseCustomEntriesUnion
}

// dlpProfileCustomNewResponseCustomEntryJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseCustomEntry]
type dlpProfileCustomNewResponseCustomEntryJSON struct {
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

func (r dlpProfileCustomNewResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponseCustomEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponseCustomEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomNewResponseCustomEntriesCustom],
// [zero_trust.DLPProfileCustomNewResponseCustomEntriesPredefined],
// [zero_trust.DLPProfileCustomNewResponseCustomEntriesIntegration],
// [zero_trust.DLPProfileCustomNewResponseCustomEntriesExactData],
// [zero_trust.DLPProfileCustomNewResponseCustomEntriesWordList].
func (r DLPProfileCustomNewResponseCustomEntry) AsUnion() DLPProfileCustomNewResponseCustomEntriesUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfileCustomNewResponseCustomEntriesCustom],
// [zero_trust.DLPProfileCustomNewResponseCustomEntriesPredefined],
// [zero_trust.DLPProfileCustomNewResponseCustomEntriesIntegration],
// [zero_trust.DLPProfileCustomNewResponseCustomEntriesExactData] or
// [zero_trust.DLPProfileCustomNewResponseCustomEntriesWordList].
type DLPProfileCustomNewResponseCustomEntriesUnion interface {
	implementsZeroTrustDLPProfileCustomNewResponseCustomEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseCustomEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseCustomEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseCustomEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseCustomEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseCustomEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseCustomEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileCustomNewResponseCustomEntriesCustom struct {
	ID        string                                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                  `json:"enabled,required"`
	Name      string                                                `json:"name,required"`
	Pattern   DLPProfileCustomNewResponseCustomEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileCustomNewResponseCustomEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                             `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomEntriesCustomJSON    `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesCustomJSON contains the JSON metadata
// for the struct [DLPProfileCustomNewResponseCustomEntriesCustom]
type dlpProfileCustomNewResponseCustomEntriesCustomJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomEntriesCustom) implementsZeroTrustDLPProfileCustomNewResponseCustomEntry() {
}

type DLPProfileCustomNewResponseCustomEntriesCustomPattern struct {
	Regex      string                                                          `json:"regex,required"`
	Validation DLPProfileCustomNewResponseCustomEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileCustomNewResponseCustomEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesCustomPatternJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponseCustomEntriesCustomPattern]
type dlpProfileCustomNewResponseCustomEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseCustomEntriesCustomPatternValidation string

const (
	DLPProfileCustomNewResponseCustomEntriesCustomPatternValidationLuhn DLPProfileCustomNewResponseCustomEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileCustomNewResponseCustomEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesCustomType string

const (
	DLPProfileCustomNewResponseCustomEntriesCustomTypeCustom DLPProfileCustomNewResponseCustomEntriesCustomType = "custom"
)

func (r DLPProfileCustomNewResponseCustomEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesPredefined struct {
	ID        string                                                 `json:"id,required" format:"uuid"`
	Enabled   bool                                                   `json:"enabled,required"`
	Name      string                                                 `json:"name,required"`
	Type      DLPProfileCustomNewResponseCustomEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                 `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomEntriesPredefinedJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesPredefinedJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponseCustomEntriesPredefined]
type dlpProfileCustomNewResponseCustomEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomEntriesPredefined) implementsZeroTrustDLPProfileCustomNewResponseCustomEntry() {
}

type DLPProfileCustomNewResponseCustomEntriesPredefinedType string

const (
	DLPProfileCustomNewResponseCustomEntriesPredefinedTypePredefined DLPProfileCustomNewResponseCustomEntriesPredefinedType = "predefined"
)

func (r DLPProfileCustomNewResponseCustomEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesIntegration struct {
	ID        string                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                    `json:"enabled,required"`
	Name      string                                                  `json:"name,required"`
	Type      DLPProfileCustomNewResponseCustomEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                               `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomEntriesIntegrationJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesIntegrationJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponseCustomEntriesIntegration]
type dlpProfileCustomNewResponseCustomEntriesIntegrationJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomEntriesIntegration) implementsZeroTrustDLPProfileCustomNewResponseCustomEntry() {
}

type DLPProfileCustomNewResponseCustomEntriesIntegrationType string

const (
	DLPProfileCustomNewResponseCustomEntriesIntegrationTypeIntegration DLPProfileCustomNewResponseCustomEntriesIntegrationType = "integration"
)

func (r DLPProfileCustomNewResponseCustomEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesExactData struct {
	ID        string                                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                  `json:"enabled,required"`
	Name      string                                                `json:"name,required"`
	Secret    bool                                                  `json:"secret,required"`
	Type      DLPProfileCustomNewResponseCustomEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                             `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomNewResponseCustomEntriesExactDataJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesExactDataJSON contains the JSON metadata
// for the struct [DLPProfileCustomNewResponseCustomEntriesExactData]
type dlpProfileCustomNewResponseCustomEntriesExactDataJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomEntriesExactData) implementsZeroTrustDLPProfileCustomNewResponseCustomEntry() {
}

type DLPProfileCustomNewResponseCustomEntriesExactDataType string

const (
	DLPProfileCustomNewResponseCustomEntriesExactDataTypeExactData DLPProfileCustomNewResponseCustomEntriesExactDataType = "exact_data"
)

func (r DLPProfileCustomNewResponseCustomEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesWordList struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                 `json:"enabled,required"`
	Name      string                                               `json:"name,required"`
	Type      DLPProfileCustomNewResponseCustomEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                          `json:"word_list,required"`
	ProfileID string                                               `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomEntriesWordListJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesWordListJSON contains the JSON metadata
// for the struct [DLPProfileCustomNewResponseCustomEntriesWordList]
type dlpProfileCustomNewResponseCustomEntriesWordListJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomEntriesWordList) implementsZeroTrustDLPProfileCustomNewResponseCustomEntry() {
}

type DLPProfileCustomNewResponseCustomEntriesWordListType string

const (
	DLPProfileCustomNewResponseCustomEntriesWordListTypeWordList DLPProfileCustomNewResponseCustomEntriesWordListType = "word_list"
)

func (r DLPProfileCustomNewResponseCustomEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesType string

const (
	DLPProfileCustomNewResponseCustomEntriesTypeCustom      DLPProfileCustomNewResponseCustomEntriesType = "custom"
	DLPProfileCustomNewResponseCustomEntriesTypePredefined  DLPProfileCustomNewResponseCustomEntriesType = "predefined"
	DLPProfileCustomNewResponseCustomEntriesTypeIntegration DLPProfileCustomNewResponseCustomEntriesType = "integration"
	DLPProfileCustomNewResponseCustomEntriesTypeExactData   DLPProfileCustomNewResponseCustomEntriesType = "exact_data"
	DLPProfileCustomNewResponseCustomEntriesTypeWordList    DLPProfileCustomNewResponseCustomEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponseCustomEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesTypeCustom, DLPProfileCustomNewResponseCustomEntriesTypePredefined, DLPProfileCustomNewResponseCustomEntriesTypeIntegration, DLPProfileCustomNewResponseCustomEntriesTypeExactData, DLPProfileCustomNewResponseCustomEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomType string

const (
	DLPProfileCustomNewResponseCustomTypeCustom DLPProfileCustomNewResponseCustomType = "custom"
)

func (r DLPProfileCustomNewResponseCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefined struct {
	// The id of the predefined profile (uuid)
	ID                string                                       `json:"id,required" format:"uuid"`
	AllowedMatchCount int64                                        `json:"allowed_match_count,required"`
	Entries           []DLPProfileCustomNewResponsePredefinedEntry `json:"entries,required"`
	// The name of the predefined profile
	Name string                                    `json:"name,required"`
	Type DLPProfileCustomNewResponsePredefinedType `json:"type,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness                          `json:"context_awareness"`
	OCREnabled       bool                                      `json:"ocr_enabled"`
	OpenAccess       bool                                      `json:"open_access"`
	JSON             dlpProfileCustomNewResponsePredefinedJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponsePredefined]
type dlpProfileCustomNewResponsePredefinedJSON struct {
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

func (r *DLPProfileCustomNewResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefined) implementsZeroTrustDLPProfileCustomNewResponse() {}

type DLPProfileCustomNewResponsePredefinedEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponsePredefinedEntriesCustomPattern].
	Pattern   interface{}                                      `json:"pattern,required"`
	ProfileID string                                           `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	Type      DLPProfileCustomNewResponsePredefinedEntriesType `json:"type,required"`
	Secret    bool                                             `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                    `json:"word_list,required"`
	JSON     dlpProfileCustomNewResponsePredefinedEntryJSON `json:"-"`
	union    DLPProfileCustomNewResponsePredefinedEntriesUnion
}

// dlpProfileCustomNewResponsePredefinedEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponsePredefinedEntry]
type dlpProfileCustomNewResponsePredefinedEntryJSON struct {
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

func (r dlpProfileCustomNewResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponsePredefinedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponsePredefinedEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomNewResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfileCustomNewResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfileCustomNewResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfileCustomNewResponsePredefinedEntriesExactData],
// [zero_trust.DLPProfileCustomNewResponsePredefinedEntriesWordList].
func (r DLPProfileCustomNewResponsePredefinedEntry) AsUnion() DLPProfileCustomNewResponsePredefinedEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfileCustomNewResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfileCustomNewResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfileCustomNewResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfileCustomNewResponsePredefinedEntriesExactData] or
// [zero_trust.DLPProfileCustomNewResponsePredefinedEntriesWordList].
type DLPProfileCustomNewResponsePredefinedEntriesUnion interface {
	implementsZeroTrustDLPProfileCustomNewResponsePredefinedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponsePredefinedEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponsePredefinedEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponsePredefinedEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponsePredefinedEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponsePredefinedEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponsePredefinedEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileCustomNewResponsePredefinedEntriesCustom struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Pattern   DLPProfileCustomNewResponsePredefinedEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileCustomNewResponsePredefinedEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                                 `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                    `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponsePredefinedEntriesCustomJSON    `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesCustomJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponsePredefinedEntriesCustom]
type dlpProfileCustomNewResponsePredefinedEntriesCustomJSON struct {
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

func (r *DLPProfileCustomNewResponsePredefinedEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedEntriesCustom) implementsZeroTrustDLPProfileCustomNewResponsePredefinedEntry() {
}

type DLPProfileCustomNewResponsePredefinedEntriesCustomPattern struct {
	Regex      string                                                              `json:"regex,required"`
	Validation DLPProfileCustomNewResponsePredefinedEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileCustomNewResponsePredefinedEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesCustomPatternJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponsePredefinedEntriesCustomPattern]
type dlpProfileCustomNewResponsePredefinedEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponsePredefinedEntriesCustomPatternValidation string

const (
	DLPProfileCustomNewResponsePredefinedEntriesCustomPatternValidationLuhn DLPProfileCustomNewResponsePredefinedEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesCustomType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesCustomTypeCustom DLPProfileCustomNewResponsePredefinedEntriesCustomType = "custom"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesPredefined struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Type      DLPProfileCustomNewResponsePredefinedEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponsePredefinedEntriesPredefinedJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesPredefinedJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponsePredefinedEntriesPredefined]
type dlpProfileCustomNewResponsePredefinedEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedEntriesPredefined) implementsZeroTrustDLPProfileCustomNewResponsePredefinedEntry() {
}

type DLPProfileCustomNewResponsePredefinedEntriesPredefinedType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesPredefinedTypePredefined DLPProfileCustomNewResponsePredefinedEntriesPredefinedType = "predefined"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesIntegration struct {
	ID        string                                                      `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                   `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                        `json:"enabled,required"`
	Name      string                                                      `json:"name,required"`
	Type      DLPProfileCustomNewResponsePredefinedEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                   `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                      `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponsePredefinedEntriesIntegrationJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesIntegrationJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponsePredefinedEntriesIntegration]
type dlpProfileCustomNewResponsePredefinedEntriesIntegrationJSON struct {
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

func (r *DLPProfileCustomNewResponsePredefinedEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedEntriesIntegration) implementsZeroTrustDLPProfileCustomNewResponsePredefinedEntry() {
}

type DLPProfileCustomNewResponsePredefinedEntriesIntegrationType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesIntegrationTypeIntegration DLPProfileCustomNewResponsePredefinedEntriesIntegrationType = "integration"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesExactData struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Secret    bool                                                      `json:"secret,required"`
	Type      DLPProfileCustomNewResponsePredefinedEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                 `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomNewResponsePredefinedEntriesExactDataJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesExactDataJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponsePredefinedEntriesExactData]
type dlpProfileCustomNewResponsePredefinedEntriesExactDataJSON struct {
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

func (r *DLPProfileCustomNewResponsePredefinedEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedEntriesExactData) implementsZeroTrustDLPProfileCustomNewResponsePredefinedEntry() {
}

type DLPProfileCustomNewResponsePredefinedEntriesExactDataType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesExactDataTypeExactData DLPProfileCustomNewResponsePredefinedEntriesExactDataType = "exact_data"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesWordList struct {
	ID        string                                                   `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                     `json:"enabled,required"`
	Name      string                                                   `json:"name,required"`
	Type      DLPProfileCustomNewResponsePredefinedEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                              `json:"word_list,required"`
	ProfileID string                                                   `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponsePredefinedEntriesWordListJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesWordListJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponsePredefinedEntriesWordList]
type dlpProfileCustomNewResponsePredefinedEntriesWordListJSON struct {
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

func (r *DLPProfileCustomNewResponsePredefinedEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedEntriesWordList) implementsZeroTrustDLPProfileCustomNewResponsePredefinedEntry() {
}

type DLPProfileCustomNewResponsePredefinedEntriesWordListType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesWordListTypeWordList DLPProfileCustomNewResponsePredefinedEntriesWordListType = "word_list"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesTypeCustom      DLPProfileCustomNewResponsePredefinedEntriesType = "custom"
	DLPProfileCustomNewResponsePredefinedEntriesTypePredefined  DLPProfileCustomNewResponsePredefinedEntriesType = "predefined"
	DLPProfileCustomNewResponsePredefinedEntriesTypeIntegration DLPProfileCustomNewResponsePredefinedEntriesType = "integration"
	DLPProfileCustomNewResponsePredefinedEntriesTypeExactData   DLPProfileCustomNewResponsePredefinedEntriesType = "exact_data"
	DLPProfileCustomNewResponsePredefinedEntriesTypeWordList    DLPProfileCustomNewResponsePredefinedEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesTypeCustom, DLPProfileCustomNewResponsePredefinedEntriesTypePredefined, DLPProfileCustomNewResponsePredefinedEntriesTypeIntegration, DLPProfileCustomNewResponsePredefinedEntriesTypeExactData, DLPProfileCustomNewResponsePredefinedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedType string

const (
	DLPProfileCustomNewResponsePredefinedTypePredefined DLPProfileCustomNewResponsePredefinedType = "predefined"
)

func (r DLPProfileCustomNewResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegration struct {
	ID        string                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                     `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileCustomNewResponseIntegrationEntry `json:"entries,required"`
	Name      string                                        `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationType    `json:"type,required"`
	UpdatedAt time.Time                                     `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                                     `json:"description,nullable"`
	JSON        dlpProfileCustomNewResponseIntegrationJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseIntegration]
type dlpProfileCustomNewResponseIntegrationJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegration) implementsZeroTrustDLPProfileCustomNewResponse() {}

type DLPProfileCustomNewResponseIntegrationEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseIntegrationEntriesCustomPattern].
	Pattern   interface{}                                       `json:"pattern,required"`
	ProfileID string                                            `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                         `json:"updated_at" format:"date-time"`
	Type      DLPProfileCustomNewResponseIntegrationEntriesType `json:"type,required"`
	Secret    bool                                              `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                     `json:"word_list,required"`
	JSON     dlpProfileCustomNewResponseIntegrationEntryJSON `json:"-"`
	union    DLPProfileCustomNewResponseIntegrationEntriesUnion
}

// dlpProfileCustomNewResponseIntegrationEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseIntegrationEntry]
type dlpProfileCustomNewResponseIntegrationEntryJSON struct {
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

func (r dlpProfileCustomNewResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponseIntegrationEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponseIntegrationEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomNewResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfileCustomNewResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfileCustomNewResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfileCustomNewResponseIntegrationEntriesExactData],
// [zero_trust.DLPProfileCustomNewResponseIntegrationEntriesWordList].
func (r DLPProfileCustomNewResponseIntegrationEntry) AsUnion() DLPProfileCustomNewResponseIntegrationEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfileCustomNewResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfileCustomNewResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfileCustomNewResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfileCustomNewResponseIntegrationEntriesExactData] or
// [zero_trust.DLPProfileCustomNewResponseIntegrationEntriesWordList].
type DLPProfileCustomNewResponseIntegrationEntriesUnion interface {
	implementsZeroTrustDLPProfileCustomNewResponseIntegrationEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseIntegrationEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseIntegrationEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseIntegrationEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseIntegrationEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseIntegrationEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseIntegrationEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileCustomNewResponseIntegrationEntriesCustom struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Pattern   DLPProfileCustomNewResponseIntegrationEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileCustomNewResponseIntegrationEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                                  `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationEntriesCustomJSON    `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesCustomJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponseIntegrationEntriesCustom]
type dlpProfileCustomNewResponseIntegrationEntriesCustomJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationEntriesCustom) implementsZeroTrustDLPProfileCustomNewResponseIntegrationEntry() {
}

type DLPProfileCustomNewResponseIntegrationEntriesCustomPattern struct {
	Regex      string                                                               `json:"regex,required"`
	Validation DLPProfileCustomNewResponseIntegrationEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileCustomNewResponseIntegrationEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesCustomPatternJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseIntegrationEntriesCustomPattern]
type dlpProfileCustomNewResponseIntegrationEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseIntegrationEntriesCustomPatternValidation string

const (
	DLPProfileCustomNewResponseIntegrationEntriesCustomPatternValidationLuhn DLPProfileCustomNewResponseIntegrationEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesCustomType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesCustomTypeCustom DLPProfileCustomNewResponseIntegrationEntriesCustomType = "custom"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesPredefined struct {
	ID        string                                                      `json:"id,required" format:"uuid"`
	Enabled   bool                                                        `json:"enabled,required"`
	Name      string                                                      `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                      `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationEntriesPredefinedJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesPredefinedJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseIntegrationEntriesPredefined]
type dlpProfileCustomNewResponseIntegrationEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationEntriesPredefined) implementsZeroTrustDLPProfileCustomNewResponseIntegrationEntry() {
}

type DLPProfileCustomNewResponseIntegrationEntriesPredefinedType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesPredefinedTypePredefined DLPProfileCustomNewResponseIntegrationEntriesPredefinedType = "predefined"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesIntegration struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationEntriesIntegrationJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesIntegrationJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseIntegrationEntriesIntegration]
type dlpProfileCustomNewResponseIntegrationEntriesIntegrationJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationEntriesIntegration) implementsZeroTrustDLPProfileCustomNewResponseIntegrationEntry() {
}

type DLPProfileCustomNewResponseIntegrationEntriesIntegrationType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesIntegrationTypeIntegration DLPProfileCustomNewResponseIntegrationEntriesIntegrationType = "integration"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesExactData struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Secret    bool                                                       `json:"secret,required"`
	Type      DLPProfileCustomNewResponseIntegrationEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                  `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomNewResponseIntegrationEntriesExactDataJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesExactDataJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponseIntegrationEntriesExactData]
type dlpProfileCustomNewResponseIntegrationEntriesExactDataJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationEntriesExactData) implementsZeroTrustDLPProfileCustomNewResponseIntegrationEntry() {
}

type DLPProfileCustomNewResponseIntegrationEntriesExactDataType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesExactDataTypeExactData DLPProfileCustomNewResponseIntegrationEntriesExactDataType = "exact_data"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesWordList struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                 `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                               `json:"word_list,required"`
	ProfileID string                                                    `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationEntriesWordListJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesWordListJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponseIntegrationEntriesWordList]
type dlpProfileCustomNewResponseIntegrationEntriesWordListJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationEntriesWordList) implementsZeroTrustDLPProfileCustomNewResponseIntegrationEntry() {
}

type DLPProfileCustomNewResponseIntegrationEntriesWordListType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesWordListTypeWordList DLPProfileCustomNewResponseIntegrationEntriesWordListType = "word_list"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesTypeCustom      DLPProfileCustomNewResponseIntegrationEntriesType = "custom"
	DLPProfileCustomNewResponseIntegrationEntriesTypePredefined  DLPProfileCustomNewResponseIntegrationEntriesType = "predefined"
	DLPProfileCustomNewResponseIntegrationEntriesTypeIntegration DLPProfileCustomNewResponseIntegrationEntriesType = "integration"
	DLPProfileCustomNewResponseIntegrationEntriesTypeExactData   DLPProfileCustomNewResponseIntegrationEntriesType = "exact_data"
	DLPProfileCustomNewResponseIntegrationEntriesTypeWordList    DLPProfileCustomNewResponseIntegrationEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesTypeCustom, DLPProfileCustomNewResponseIntegrationEntriesTypePredefined, DLPProfileCustomNewResponseIntegrationEntriesTypeIntegration, DLPProfileCustomNewResponseIntegrationEntriesTypeExactData, DLPProfileCustomNewResponseIntegrationEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationType string

const (
	DLPProfileCustomNewResponseIntegrationTypeIntegration DLPProfileCustomNewResponseIntegrationType = "integration"
)

func (r DLPProfileCustomNewResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseType string

const (
	DLPProfileCustomNewResponseTypeCustom      DLPProfileCustomNewResponseType = "custom"
	DLPProfileCustomNewResponseTypePredefined  DLPProfileCustomNewResponseType = "predefined"
	DLPProfileCustomNewResponseTypeIntegration DLPProfileCustomNewResponseType = "integration"
)

func (r DLPProfileCustomNewResponseType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseTypeCustom, DLPProfileCustomNewResponseTypePredefined, DLPProfileCustomNewResponseTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponse struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// When the profile was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile
	Description string `json:"description,nullable"`
	// This field can have the runtime type of
	// [[]DLPProfileCustomUpdateResponseCustomEntry],
	// [[]DLPProfileCustomUpdateResponsePredefinedEntry],
	// [[]DLPProfileCustomUpdateResponseIntegrationEntry].
	Entries interface{} `json:"entries"`
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// The name of the profile
	Name       string `json:"name,required"`
	OCREnabled bool   `json:"ocr_enabled"`
	// When the profile was lasted updated
	UpdatedAt  time.Time                          `json:"updated_at" format:"date-time"`
	Type       DLPProfileCustomUpdateResponseType `json:"type,required"`
	OpenAccess bool                               `json:"open_access"`
	JSON       dlpProfileCustomUpdateResponseJSON `json:"-"`
	union      DLPProfileCustomUpdateResponseUnion
}

// dlpProfileCustomUpdateResponseJSON contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponse]
type dlpProfileCustomUpdateResponseJSON struct {
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

func (r dlpProfileCustomUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomUpdateResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomUpdateResponseCustom],
// [zero_trust.DLPProfileCustomUpdateResponsePredefined],
// [zero_trust.DLPProfileCustomUpdateResponseIntegration].
func (r DLPProfileCustomUpdateResponse) AsUnion() DLPProfileCustomUpdateResponseUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfileCustomUpdateResponseCustom],
// [zero_trust.DLPProfileCustomUpdateResponsePredefined] or
// [zero_trust.DLPProfileCustomUpdateResponseIntegration].
type DLPProfileCustomUpdateResponseUnion interface {
	implementsZeroTrustDLPProfileCustomUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomUpdateResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
	)
}

type DLPProfileCustomUpdateResponseCustom struct {
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness,required"`
	// When the profile was created
	CreatedAt time.Time                                   `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileCustomUpdateResponseCustomEntry `json:"entries,required"`
	// The name of the profile
	Name       string                                   `json:"name,required"`
	OCREnabled bool                                     `json:"ocr_enabled,required"`
	Type       DLPProfileCustomUpdateResponseCustomType `json:"type,required"`
	// When the profile was lasted updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                                   `json:"description,nullable"`
	JSON        dlpProfileCustomUpdateResponseCustomJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponseCustom]
type dlpProfileCustomUpdateResponseCustomJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustom) implementsZeroTrustDLPProfileCustomUpdateResponse() {}

type DLPProfileCustomUpdateResponseCustomEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileCustomUpdateResponseCustomEntriesCustomPattern].
	Pattern   interface{}                                     `json:"pattern,required"`
	ProfileID string                                          `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                       `json:"updated_at" format:"date-time"`
	Type      DLPProfileCustomUpdateResponseCustomEntriesType `json:"type,required"`
	Secret    bool                                            `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                   `json:"word_list,required"`
	JSON     dlpProfileCustomUpdateResponseCustomEntryJSON `json:"-"`
	union    DLPProfileCustomUpdateResponseCustomEntriesUnion
}

// dlpProfileCustomUpdateResponseCustomEntryJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponseCustomEntry]
type dlpProfileCustomUpdateResponseCustomEntryJSON struct {
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

func (r dlpProfileCustomUpdateResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomUpdateResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomUpdateResponseCustomEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomUpdateResponseCustomEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomUpdateResponseCustomEntriesCustom],
// [zero_trust.DLPProfileCustomUpdateResponseCustomEntriesPredefined],
// [zero_trust.DLPProfileCustomUpdateResponseCustomEntriesIntegration],
// [zero_trust.DLPProfileCustomUpdateResponseCustomEntriesExactData],
// [zero_trust.DLPProfileCustomUpdateResponseCustomEntriesWordList].
func (r DLPProfileCustomUpdateResponseCustomEntry) AsUnion() DLPProfileCustomUpdateResponseCustomEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfileCustomUpdateResponseCustomEntriesCustom],
// [zero_trust.DLPProfileCustomUpdateResponseCustomEntriesPredefined],
// [zero_trust.DLPProfileCustomUpdateResponseCustomEntriesIntegration],
// [zero_trust.DLPProfileCustomUpdateResponseCustomEntriesExactData] or
// [zero_trust.DLPProfileCustomUpdateResponseCustomEntriesWordList].
type DLPProfileCustomUpdateResponseCustomEntriesUnion interface {
	implementsZeroTrustDLPProfileCustomUpdateResponseCustomEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomUpdateResponseCustomEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseCustomEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseCustomEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseCustomEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseCustomEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseCustomEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileCustomUpdateResponseCustomEntriesCustom struct {
	ID        string                                                   `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                     `json:"enabled,required"`
	Name      string                                                   `json:"name,required"`
	Pattern   DLPProfileCustomUpdateResponseCustomEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileCustomUpdateResponseCustomEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                                `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                   `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseCustomEntriesCustomJSON    `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesCustomJSON contains the JSON metadata
// for the struct [DLPProfileCustomUpdateResponseCustomEntriesCustom]
type dlpProfileCustomUpdateResponseCustomEntriesCustomJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomEntriesCustom) implementsZeroTrustDLPProfileCustomUpdateResponseCustomEntry() {
}

type DLPProfileCustomUpdateResponseCustomEntriesCustomPattern struct {
	Regex      string                                                             `json:"regex,required"`
	Validation DLPProfileCustomUpdateResponseCustomEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileCustomUpdateResponseCustomEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesCustomPatternJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomUpdateResponseCustomEntriesCustomPattern]
type dlpProfileCustomUpdateResponseCustomEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseCustomEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseCustomEntriesCustomPatternValidation string

const (
	DLPProfileCustomUpdateResponseCustomEntriesCustomPatternValidationLuhn DLPProfileCustomUpdateResponseCustomEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesCustomType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesCustomTypeCustom DLPProfileCustomUpdateResponseCustomEntriesCustomType = "custom"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesPredefined struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseCustomEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                    `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseCustomEntriesPredefinedJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesPredefinedJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseCustomEntriesPredefined]
type dlpProfileCustomUpdateResponseCustomEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseCustomEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomEntriesPredefined) implementsZeroTrustDLPProfileCustomUpdateResponseCustomEntry() {
}

type DLPProfileCustomUpdateResponseCustomEntriesPredefinedType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesPredefinedTypePredefined DLPProfileCustomUpdateResponseCustomEntriesPredefinedType = "predefined"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesIntegration struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseCustomEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                  `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseCustomEntriesIntegrationJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesIntegrationJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseCustomEntriesIntegration]
type dlpProfileCustomUpdateResponseCustomEntriesIntegrationJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomEntriesIntegration) implementsZeroTrustDLPProfileCustomUpdateResponseCustomEntry() {
}

type DLPProfileCustomUpdateResponseCustomEntriesIntegrationType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesIntegrationTypeIntegration DLPProfileCustomUpdateResponseCustomEntriesIntegrationType = "integration"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesExactData struct {
	ID        string                                                   `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                     `json:"enabled,required"`
	Name      string                                                   `json:"name,required"`
	Secret    bool                                                     `json:"secret,required"`
	Type      DLPProfileCustomUpdateResponseCustomEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomUpdateResponseCustomEntriesExactDataJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesExactDataJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseCustomEntriesExactData]
type dlpProfileCustomUpdateResponseCustomEntriesExactDataJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomEntriesExactData) implementsZeroTrustDLPProfileCustomUpdateResponseCustomEntry() {
}

type DLPProfileCustomUpdateResponseCustomEntriesExactDataType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesExactDataTypeExactData DLPProfileCustomUpdateResponseCustomEntriesExactDataType = "exact_data"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesWordList struct {
	ID        string                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                    `json:"enabled,required"`
	Name      string                                                  `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseCustomEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                               `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                             `json:"word_list,required"`
	ProfileID string                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseCustomEntriesWordListJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesWordListJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseCustomEntriesWordList]
type dlpProfileCustomUpdateResponseCustomEntriesWordListJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomEntriesWordList) implementsZeroTrustDLPProfileCustomUpdateResponseCustomEntry() {
}

type DLPProfileCustomUpdateResponseCustomEntriesWordListType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesWordListTypeWordList DLPProfileCustomUpdateResponseCustomEntriesWordListType = "word_list"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesTypeCustom      DLPProfileCustomUpdateResponseCustomEntriesType = "custom"
	DLPProfileCustomUpdateResponseCustomEntriesTypePredefined  DLPProfileCustomUpdateResponseCustomEntriesType = "predefined"
	DLPProfileCustomUpdateResponseCustomEntriesTypeIntegration DLPProfileCustomUpdateResponseCustomEntriesType = "integration"
	DLPProfileCustomUpdateResponseCustomEntriesTypeExactData   DLPProfileCustomUpdateResponseCustomEntriesType = "exact_data"
	DLPProfileCustomUpdateResponseCustomEntriesTypeWordList    DLPProfileCustomUpdateResponseCustomEntriesType = "word_list"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesTypeCustom, DLPProfileCustomUpdateResponseCustomEntriesTypePredefined, DLPProfileCustomUpdateResponseCustomEntriesTypeIntegration, DLPProfileCustomUpdateResponseCustomEntriesTypeExactData, DLPProfileCustomUpdateResponseCustomEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomType string

const (
	DLPProfileCustomUpdateResponseCustomTypeCustom DLPProfileCustomUpdateResponseCustomType = "custom"
)

func (r DLPProfileCustomUpdateResponseCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefined struct {
	// The id of the predefined profile (uuid)
	ID                string                                          `json:"id,required" format:"uuid"`
	AllowedMatchCount int64                                           `json:"allowed_match_count,required"`
	Entries           []DLPProfileCustomUpdateResponsePredefinedEntry `json:"entries,required"`
	// The name of the predefined profile
	Name string                                       `json:"name,required"`
	Type DLPProfileCustomUpdateResponsePredefinedType `json:"type,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness                             `json:"context_awareness"`
	OCREnabled       bool                                         `json:"ocr_enabled"`
	OpenAccess       bool                                         `json:"open_access"`
	JSON             dlpProfileCustomUpdateResponsePredefinedJSON `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponsePredefined]
type dlpProfileCustomUpdateResponsePredefinedJSON struct {
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

func (r *DLPProfileCustomUpdateResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponsePredefined) implementsZeroTrustDLPProfileCustomUpdateResponse() {
}

type DLPProfileCustomUpdateResponsePredefinedEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileCustomUpdateResponsePredefinedEntriesCustomPattern].
	Pattern   interface{}                                         `json:"pattern,required"`
	ProfileID string                                              `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                           `json:"updated_at" format:"date-time"`
	Type      DLPProfileCustomUpdateResponsePredefinedEntriesType `json:"type,required"`
	Secret    bool                                                `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                       `json:"word_list,required"`
	JSON     dlpProfileCustomUpdateResponsePredefinedEntryJSON `json:"-"`
	union    DLPProfileCustomUpdateResponsePredefinedEntriesUnion
}

// dlpProfileCustomUpdateResponsePredefinedEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomUpdateResponsePredefinedEntry]
type dlpProfileCustomUpdateResponsePredefinedEntryJSON struct {
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

func (r dlpProfileCustomUpdateResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomUpdateResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomUpdateResponsePredefinedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomUpdateResponsePredefinedEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomUpdateResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfileCustomUpdateResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfileCustomUpdateResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfileCustomUpdateResponsePredefinedEntriesExactData],
// [zero_trust.DLPProfileCustomUpdateResponsePredefinedEntriesWordList].
func (r DLPProfileCustomUpdateResponsePredefinedEntry) AsUnion() DLPProfileCustomUpdateResponsePredefinedEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfileCustomUpdateResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfileCustomUpdateResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfileCustomUpdateResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfileCustomUpdateResponsePredefinedEntriesExactData] or
// [zero_trust.DLPProfileCustomUpdateResponsePredefinedEntriesWordList].
type DLPProfileCustomUpdateResponsePredefinedEntriesUnion interface {
	implementsZeroTrustDLPProfileCustomUpdateResponsePredefinedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomUpdateResponsePredefinedEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponsePredefinedEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponsePredefinedEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponsePredefinedEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponsePredefinedEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponsePredefinedEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileCustomUpdateResponsePredefinedEntriesCustom struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Pattern   DLPProfileCustomUpdateResponsePredefinedEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileCustomUpdateResponsePredefinedEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponsePredefinedEntriesCustomJSON    `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesCustomJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponsePredefinedEntriesCustom]
type dlpProfileCustomUpdateResponsePredefinedEntriesCustomJSON struct {
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

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponsePredefinedEntriesCustom) implementsZeroTrustDLPProfileCustomUpdateResponsePredefinedEntry() {
}

type DLPProfileCustomUpdateResponsePredefinedEntriesCustomPattern struct {
	Regex      string                                                                 `json:"regex,required"`
	Validation DLPProfileCustomUpdateResponsePredefinedEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileCustomUpdateResponsePredefinedEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesCustomPatternJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponsePredefinedEntriesCustomPattern]
type dlpProfileCustomUpdateResponsePredefinedEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponsePredefinedEntriesCustomPatternValidation string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesCustomPatternValidationLuhn DLPProfileCustomUpdateResponsePredefinedEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesCustomType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesCustomTypeCustom DLPProfileCustomUpdateResponsePredefinedEntriesCustomType = "custom"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesPredefined struct {
	ID        string                                                        `json:"id,required" format:"uuid"`
	Enabled   bool                                                          `json:"enabled,required"`
	Name      string                                                        `json:"name,required"`
	Type      DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedJSON `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomUpdateResponsePredefinedEntriesPredefined]
type dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponsePredefinedEntriesPredefined) implementsZeroTrustDLPProfileCustomUpdateResponsePredefinedEntry() {
}

type DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedTypePredefined DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedType = "predefined"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesIntegration struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                           `json:"enabled,required"`
	Name      string                                                         `json:"name,required"`
	Type      DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                      `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponsePredefinedEntriesIntegrationJSON `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesIntegrationJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomUpdateResponsePredefinedEntriesIntegration]
type dlpProfileCustomUpdateResponsePredefinedEntriesIntegrationJSON struct {
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

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponsePredefinedEntriesIntegration) implementsZeroTrustDLPProfileCustomUpdateResponsePredefinedEntry() {
}

type DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationTypeIntegration DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationType = "integration"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesExactData struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Secret    bool                                                         `json:"secret,required"`
	Type      DLPProfileCustomUpdateResponsePredefinedEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomUpdateResponsePredefinedEntriesExactDataJSON `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesExactDataJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomUpdateResponsePredefinedEntriesExactData]
type dlpProfileCustomUpdateResponsePredefinedEntriesExactDataJSON struct {
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

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponsePredefinedEntriesExactData) implementsZeroTrustDLPProfileCustomUpdateResponsePredefinedEntry() {
}

type DLPProfileCustomUpdateResponsePredefinedEntriesExactDataType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesExactDataTypeExactData DLPProfileCustomUpdateResponsePredefinedEntriesExactDataType = "exact_data"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesWordList struct {
	ID        string                                                      `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                   `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                        `json:"enabled,required"`
	Name      string                                                      `json:"name,required"`
	Type      DLPProfileCustomUpdateResponsePredefinedEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                   `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                 `json:"word_list,required"`
	ProfileID string                                                      `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponsePredefinedEntriesWordListJSON `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesWordListJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomUpdateResponsePredefinedEntriesWordList]
type dlpProfileCustomUpdateResponsePredefinedEntriesWordListJSON struct {
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

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponsePredefinedEntriesWordList) implementsZeroTrustDLPProfileCustomUpdateResponsePredefinedEntry() {
}

type DLPProfileCustomUpdateResponsePredefinedEntriesWordListType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesWordListTypeWordList DLPProfileCustomUpdateResponsePredefinedEntriesWordListType = "word_list"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesTypeCustom      DLPProfileCustomUpdateResponsePredefinedEntriesType = "custom"
	DLPProfileCustomUpdateResponsePredefinedEntriesTypePredefined  DLPProfileCustomUpdateResponsePredefinedEntriesType = "predefined"
	DLPProfileCustomUpdateResponsePredefinedEntriesTypeIntegration DLPProfileCustomUpdateResponsePredefinedEntriesType = "integration"
	DLPProfileCustomUpdateResponsePredefinedEntriesTypeExactData   DLPProfileCustomUpdateResponsePredefinedEntriesType = "exact_data"
	DLPProfileCustomUpdateResponsePredefinedEntriesTypeWordList    DLPProfileCustomUpdateResponsePredefinedEntriesType = "word_list"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesTypeCustom, DLPProfileCustomUpdateResponsePredefinedEntriesTypePredefined, DLPProfileCustomUpdateResponsePredefinedEntriesTypeIntegration, DLPProfileCustomUpdateResponsePredefinedEntriesTypeExactData, DLPProfileCustomUpdateResponsePredefinedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedType string

const (
	DLPProfileCustomUpdateResponsePredefinedTypePredefined DLPProfileCustomUpdateResponsePredefinedType = "predefined"
)

func (r DLPProfileCustomUpdateResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegration struct {
	ID        string                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                        `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileCustomUpdateResponseIntegrationEntry `json:"entries,required"`
	Name      string                                           `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseIntegrationType    `json:"type,required"`
	UpdatedAt time.Time                                        `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                                        `json:"description,nullable"`
	JSON        dlpProfileCustomUpdateResponseIntegrationJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponseIntegration]
type dlpProfileCustomUpdateResponseIntegrationJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegration) implementsZeroTrustDLPProfileCustomUpdateResponse() {
}

type DLPProfileCustomUpdateResponseIntegrationEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileCustomUpdateResponseIntegrationEntriesCustomPattern].
	Pattern   interface{}                                          `json:"pattern,required"`
	ProfileID string                                               `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                            `json:"updated_at" format:"date-time"`
	Type      DLPProfileCustomUpdateResponseIntegrationEntriesType `json:"type,required"`
	Secret    bool                                                 `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                        `json:"word_list,required"`
	JSON     dlpProfileCustomUpdateResponseIntegrationEntryJSON `json:"-"`
	union    DLPProfileCustomUpdateResponseIntegrationEntriesUnion
}

// dlpProfileCustomUpdateResponseIntegrationEntryJSON contains the JSON metadata
// for the struct [DLPProfileCustomUpdateResponseIntegrationEntry]
type dlpProfileCustomUpdateResponseIntegrationEntryJSON struct {
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

func (r dlpProfileCustomUpdateResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomUpdateResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomUpdateResponseIntegrationEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomUpdateResponseIntegrationEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomUpdateResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfileCustomUpdateResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfileCustomUpdateResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfileCustomUpdateResponseIntegrationEntriesExactData],
// [zero_trust.DLPProfileCustomUpdateResponseIntegrationEntriesWordList].
func (r DLPProfileCustomUpdateResponseIntegrationEntry) AsUnion() DLPProfileCustomUpdateResponseIntegrationEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfileCustomUpdateResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfileCustomUpdateResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfileCustomUpdateResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfileCustomUpdateResponseIntegrationEntriesExactData] or
// [zero_trust.DLPProfileCustomUpdateResponseIntegrationEntriesWordList].
type DLPProfileCustomUpdateResponseIntegrationEntriesUnion interface {
	implementsZeroTrustDLPProfileCustomUpdateResponseIntegrationEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomUpdateResponseIntegrationEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileCustomUpdateResponseIntegrationEntriesCustom struct {
	ID        string                                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                          `json:"enabled,required"`
	Name      string                                                        `json:"name,required"`
	Pattern   DLPProfileCustomUpdateResponseIntegrationEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileCustomUpdateResponseIntegrationEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                                     `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseIntegrationEntriesCustomJSON    `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesCustomJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseIntegrationEntriesCustom]
type dlpProfileCustomUpdateResponseIntegrationEntriesCustomJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationEntriesCustom) implementsZeroTrustDLPProfileCustomUpdateResponseIntegrationEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationEntriesCustomPattern struct {
	Regex      string                                                                  `json:"regex,required"`
	Validation DLPProfileCustomUpdateResponseIntegrationEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileCustomUpdateResponseIntegrationEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesCustomPatternJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationEntriesCustomPattern]
type dlpProfileCustomUpdateResponseIntegrationEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseIntegrationEntriesCustomPatternValidation string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesCustomPatternValidationLuhn DLPProfileCustomUpdateResponseIntegrationEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesCustomType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesCustomTypeCustom DLPProfileCustomUpdateResponseIntegrationEntriesCustomType = "custom"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesPredefined struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	Enabled   bool                                                           `json:"enabled,required"`
	Name      string                                                         `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationEntriesPredefined]
type dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationEntriesPredefined) implementsZeroTrustDLPProfileCustomUpdateResponseIntegrationEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedTypePredefined DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedType = "predefined"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesIntegration struct {
	ID        string                                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                            `json:"enabled,required"`
	Name      string                                                          `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                       `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseIntegrationEntriesIntegrationJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesIntegrationJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationEntriesIntegration]
type dlpProfileCustomUpdateResponseIntegrationEntriesIntegrationJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationEntriesIntegration) implementsZeroTrustDLPProfileCustomUpdateResponseIntegrationEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationTypeIntegration DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationType = "integration"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesExactData struct {
	ID        string                                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                          `json:"enabled,required"`
	Name      string                                                        `json:"name,required"`
	Secret    bool                                                          `json:"secret,required"`
	Type      DLPProfileCustomUpdateResponseIntegrationEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                     `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomUpdateResponseIntegrationEntriesExactDataJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesExactDataJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationEntriesExactData]
type dlpProfileCustomUpdateResponseIntegrationEntriesExactDataJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationEntriesExactData) implementsZeroTrustDLPProfileCustomUpdateResponseIntegrationEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationEntriesExactDataType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesExactDataTypeExactData DLPProfileCustomUpdateResponseIntegrationEntriesExactDataType = "exact_data"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesWordList struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseIntegrationEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                  `json:"word_list,required"`
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseIntegrationEntriesWordListJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesWordListJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationEntriesWordList]
type dlpProfileCustomUpdateResponseIntegrationEntriesWordListJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationEntriesWordList) implementsZeroTrustDLPProfileCustomUpdateResponseIntegrationEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationEntriesWordListType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesWordListTypeWordList DLPProfileCustomUpdateResponseIntegrationEntriesWordListType = "word_list"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesTypeCustom      DLPProfileCustomUpdateResponseIntegrationEntriesType = "custom"
	DLPProfileCustomUpdateResponseIntegrationEntriesTypePredefined  DLPProfileCustomUpdateResponseIntegrationEntriesType = "predefined"
	DLPProfileCustomUpdateResponseIntegrationEntriesTypeIntegration DLPProfileCustomUpdateResponseIntegrationEntriesType = "integration"
	DLPProfileCustomUpdateResponseIntegrationEntriesTypeExactData   DLPProfileCustomUpdateResponseIntegrationEntriesType = "exact_data"
	DLPProfileCustomUpdateResponseIntegrationEntriesTypeWordList    DLPProfileCustomUpdateResponseIntegrationEntriesType = "word_list"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesTypeCustom, DLPProfileCustomUpdateResponseIntegrationEntriesTypePredefined, DLPProfileCustomUpdateResponseIntegrationEntriesTypeIntegration, DLPProfileCustomUpdateResponseIntegrationEntriesTypeExactData, DLPProfileCustomUpdateResponseIntegrationEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationType string

const (
	DLPProfileCustomUpdateResponseIntegrationTypeIntegration DLPProfileCustomUpdateResponseIntegrationType = "integration"
)

func (r DLPProfileCustomUpdateResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseType string

const (
	DLPProfileCustomUpdateResponseTypeCustom      DLPProfileCustomUpdateResponseType = "custom"
	DLPProfileCustomUpdateResponseTypePredefined  DLPProfileCustomUpdateResponseType = "predefined"
	DLPProfileCustomUpdateResponseTypeIntegration DLPProfileCustomUpdateResponseType = "integration"
)

func (r DLPProfileCustomUpdateResponseType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseTypeCustom, DLPProfileCustomUpdateResponseTypePredefined, DLPProfileCustomUpdateResponseTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomDeleteResponse = interface{}

type DLPProfileCustomGetResponse struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// When the profile was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile
	Description string `json:"description,nullable"`
	// This field can have the runtime type of
	// [[]DLPProfileCustomGetResponseCustomEntry],
	// [[]DLPProfileCustomGetResponsePredefinedEntry],
	// [[]DLPProfileCustomGetResponseIntegrationEntry].
	Entries interface{} `json:"entries"`
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// The name of the profile
	Name       string `json:"name,required"`
	OCREnabled bool   `json:"ocr_enabled"`
	// When the profile was lasted updated
	UpdatedAt  time.Time                       `json:"updated_at" format:"date-time"`
	Type       DLPProfileCustomGetResponseType `json:"type,required"`
	OpenAccess bool                            `json:"open_access"`
	JSON       dlpProfileCustomGetResponseJSON `json:"-"`
	union      DLPProfileCustomGetResponseUnion
}

// dlpProfileCustomGetResponseJSON contains the JSON metadata for the struct
// [DLPProfileCustomGetResponse]
type dlpProfileCustomGetResponseJSON struct {
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

func (r dlpProfileCustomGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomGetResponseCustom],
// [zero_trust.DLPProfileCustomGetResponsePredefined],
// [zero_trust.DLPProfileCustomGetResponseIntegration].
func (r DLPProfileCustomGetResponse) AsUnion() DLPProfileCustomGetResponseUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfileCustomGetResponseCustom],
// [zero_trust.DLPProfileCustomGetResponsePredefined] or
// [zero_trust.DLPProfileCustomGetResponseIntegration].
type DLPProfileCustomGetResponseUnion interface {
	implementsZeroTrustDLPProfileCustomGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
	)
}

type DLPProfileCustomGetResponseCustom struct {
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness,required"`
	// When the profile was created
	CreatedAt time.Time                                `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileCustomGetResponseCustomEntry `json:"entries,required"`
	// The name of the profile
	Name       string                                `json:"name,required"`
	OCREnabled bool                                  `json:"ocr_enabled,required"`
	Type       DLPProfileCustomGetResponseCustomType `json:"type,required"`
	// When the profile was lasted updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                                `json:"description,nullable"`
	JSON        dlpProfileCustomGetResponseCustomJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomJSON contains the JSON metadata for the struct
// [DLPProfileCustomGetResponseCustom]
type dlpProfileCustomGetResponseCustomJSON struct {
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

func (r *DLPProfileCustomGetResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustom) implementsZeroTrustDLPProfileCustomGetResponse() {}

type DLPProfileCustomGetResponseCustomEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileCustomGetResponseCustomEntriesCustomPattern].
	Pattern   interface{}                                  `json:"pattern,required"`
	ProfileID string                                       `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	Type      DLPProfileCustomGetResponseCustomEntriesType `json:"type,required"`
	Secret    bool                                         `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                `json:"word_list,required"`
	JSON     dlpProfileCustomGetResponseCustomEntryJSON `json:"-"`
	union    DLPProfileCustomGetResponseCustomEntriesUnion
}

// dlpProfileCustomGetResponseCustomEntryJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseCustomEntry]
type dlpProfileCustomGetResponseCustomEntryJSON struct {
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

func (r dlpProfileCustomGetResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomGetResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomGetResponseCustomEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomGetResponseCustomEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomGetResponseCustomEntriesCustom],
// [zero_trust.DLPProfileCustomGetResponseCustomEntriesPredefined],
// [zero_trust.DLPProfileCustomGetResponseCustomEntriesIntegration],
// [zero_trust.DLPProfileCustomGetResponseCustomEntriesExactData],
// [zero_trust.DLPProfileCustomGetResponseCustomEntriesWordList].
func (r DLPProfileCustomGetResponseCustomEntry) AsUnion() DLPProfileCustomGetResponseCustomEntriesUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfileCustomGetResponseCustomEntriesCustom],
// [zero_trust.DLPProfileCustomGetResponseCustomEntriesPredefined],
// [zero_trust.DLPProfileCustomGetResponseCustomEntriesIntegration],
// [zero_trust.DLPProfileCustomGetResponseCustomEntriesExactData] or
// [zero_trust.DLPProfileCustomGetResponseCustomEntriesWordList].
type DLPProfileCustomGetResponseCustomEntriesUnion interface {
	implementsZeroTrustDLPProfileCustomGetResponseCustomEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomGetResponseCustomEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseCustomEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseCustomEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseCustomEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseCustomEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseCustomEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileCustomGetResponseCustomEntriesCustom struct {
	ID        string                                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                  `json:"enabled,required"`
	Name      string                                                `json:"name,required"`
	Pattern   DLPProfileCustomGetResponseCustomEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileCustomGetResponseCustomEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                             `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseCustomEntriesCustomJSON    `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesCustomJSON contains the JSON metadata
// for the struct [DLPProfileCustomGetResponseCustomEntriesCustom]
type dlpProfileCustomGetResponseCustomEntriesCustomJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomEntriesCustom) implementsZeroTrustDLPProfileCustomGetResponseCustomEntry() {
}

type DLPProfileCustomGetResponseCustomEntriesCustomPattern struct {
	Regex      string                                                          `json:"regex,required"`
	Validation DLPProfileCustomGetResponseCustomEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileCustomGetResponseCustomEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesCustomPatternJSON contains the JSON
// metadata for the struct [DLPProfileCustomGetResponseCustomEntriesCustomPattern]
type dlpProfileCustomGetResponseCustomEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseCustomEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseCustomEntriesCustomPatternValidation string

const (
	DLPProfileCustomGetResponseCustomEntriesCustomPatternValidationLuhn DLPProfileCustomGetResponseCustomEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileCustomGetResponseCustomEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesCustomType string

const (
	DLPProfileCustomGetResponseCustomEntriesCustomTypeCustom DLPProfileCustomGetResponseCustomEntriesCustomType = "custom"
)

func (r DLPProfileCustomGetResponseCustomEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesPredefined struct {
	ID        string                                                 `json:"id,required" format:"uuid"`
	Enabled   bool                                                   `json:"enabled,required"`
	Name      string                                                 `json:"name,required"`
	Type      DLPProfileCustomGetResponseCustomEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                 `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseCustomEntriesPredefinedJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesPredefinedJSON contains the JSON
// metadata for the struct [DLPProfileCustomGetResponseCustomEntriesPredefined]
type dlpProfileCustomGetResponseCustomEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseCustomEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomEntriesPredefined) implementsZeroTrustDLPProfileCustomGetResponseCustomEntry() {
}

type DLPProfileCustomGetResponseCustomEntriesPredefinedType string

const (
	DLPProfileCustomGetResponseCustomEntriesPredefinedTypePredefined DLPProfileCustomGetResponseCustomEntriesPredefinedType = "predefined"
)

func (r DLPProfileCustomGetResponseCustomEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesIntegration struct {
	ID        string                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                    `json:"enabled,required"`
	Name      string                                                  `json:"name,required"`
	Type      DLPProfileCustomGetResponseCustomEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                               `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseCustomEntriesIntegrationJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesIntegrationJSON contains the JSON
// metadata for the struct [DLPProfileCustomGetResponseCustomEntriesIntegration]
type dlpProfileCustomGetResponseCustomEntriesIntegrationJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomEntriesIntegration) implementsZeroTrustDLPProfileCustomGetResponseCustomEntry() {
}

type DLPProfileCustomGetResponseCustomEntriesIntegrationType string

const (
	DLPProfileCustomGetResponseCustomEntriesIntegrationTypeIntegration DLPProfileCustomGetResponseCustomEntriesIntegrationType = "integration"
)

func (r DLPProfileCustomGetResponseCustomEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesExactData struct {
	ID        string                                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                  `json:"enabled,required"`
	Name      string                                                `json:"name,required"`
	Secret    bool                                                  `json:"secret,required"`
	Type      DLPProfileCustomGetResponseCustomEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                             `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomGetResponseCustomEntriesExactDataJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesExactDataJSON contains the JSON metadata
// for the struct [DLPProfileCustomGetResponseCustomEntriesExactData]
type dlpProfileCustomGetResponseCustomEntriesExactDataJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomEntriesExactData) implementsZeroTrustDLPProfileCustomGetResponseCustomEntry() {
}

type DLPProfileCustomGetResponseCustomEntriesExactDataType string

const (
	DLPProfileCustomGetResponseCustomEntriesExactDataTypeExactData DLPProfileCustomGetResponseCustomEntriesExactDataType = "exact_data"
)

func (r DLPProfileCustomGetResponseCustomEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesWordList struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                 `json:"enabled,required"`
	Name      string                                               `json:"name,required"`
	Type      DLPProfileCustomGetResponseCustomEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                          `json:"word_list,required"`
	ProfileID string                                               `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseCustomEntriesWordListJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesWordListJSON contains the JSON metadata
// for the struct [DLPProfileCustomGetResponseCustomEntriesWordList]
type dlpProfileCustomGetResponseCustomEntriesWordListJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomEntriesWordList) implementsZeroTrustDLPProfileCustomGetResponseCustomEntry() {
}

type DLPProfileCustomGetResponseCustomEntriesWordListType string

const (
	DLPProfileCustomGetResponseCustomEntriesWordListTypeWordList DLPProfileCustomGetResponseCustomEntriesWordListType = "word_list"
)

func (r DLPProfileCustomGetResponseCustomEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesType string

const (
	DLPProfileCustomGetResponseCustomEntriesTypeCustom      DLPProfileCustomGetResponseCustomEntriesType = "custom"
	DLPProfileCustomGetResponseCustomEntriesTypePredefined  DLPProfileCustomGetResponseCustomEntriesType = "predefined"
	DLPProfileCustomGetResponseCustomEntriesTypeIntegration DLPProfileCustomGetResponseCustomEntriesType = "integration"
	DLPProfileCustomGetResponseCustomEntriesTypeExactData   DLPProfileCustomGetResponseCustomEntriesType = "exact_data"
	DLPProfileCustomGetResponseCustomEntriesTypeWordList    DLPProfileCustomGetResponseCustomEntriesType = "word_list"
)

func (r DLPProfileCustomGetResponseCustomEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesTypeCustom, DLPProfileCustomGetResponseCustomEntriesTypePredefined, DLPProfileCustomGetResponseCustomEntriesTypeIntegration, DLPProfileCustomGetResponseCustomEntriesTypeExactData, DLPProfileCustomGetResponseCustomEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomType string

const (
	DLPProfileCustomGetResponseCustomTypeCustom DLPProfileCustomGetResponseCustomType = "custom"
)

func (r DLPProfileCustomGetResponseCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefined struct {
	// The id of the predefined profile (uuid)
	ID                string                                       `json:"id,required" format:"uuid"`
	AllowedMatchCount int64                                        `json:"allowed_match_count,required"`
	Entries           []DLPProfileCustomGetResponsePredefinedEntry `json:"entries,required"`
	// The name of the predefined profile
	Name string                                    `json:"name,required"`
	Type DLPProfileCustomGetResponsePredefinedType `json:"type,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness                          `json:"context_awareness"`
	OCREnabled       bool                                      `json:"ocr_enabled"`
	OpenAccess       bool                                      `json:"open_access"`
	JSON             dlpProfileCustomGetResponsePredefinedJSON `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponsePredefined]
type dlpProfileCustomGetResponsePredefinedJSON struct {
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

func (r *DLPProfileCustomGetResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponsePredefined) implementsZeroTrustDLPProfileCustomGetResponse() {}

type DLPProfileCustomGetResponsePredefinedEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileCustomGetResponsePredefinedEntriesCustomPattern].
	Pattern   interface{}                                      `json:"pattern,required"`
	ProfileID string                                           `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	Type      DLPProfileCustomGetResponsePredefinedEntriesType `json:"type,required"`
	Secret    bool                                             `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                    `json:"word_list,required"`
	JSON     dlpProfileCustomGetResponsePredefinedEntryJSON `json:"-"`
	union    DLPProfileCustomGetResponsePredefinedEntriesUnion
}

// dlpProfileCustomGetResponsePredefinedEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomGetResponsePredefinedEntry]
type dlpProfileCustomGetResponsePredefinedEntryJSON struct {
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

func (r dlpProfileCustomGetResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomGetResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomGetResponsePredefinedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomGetResponsePredefinedEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomGetResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfileCustomGetResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfileCustomGetResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfileCustomGetResponsePredefinedEntriesExactData],
// [zero_trust.DLPProfileCustomGetResponsePredefinedEntriesWordList].
func (r DLPProfileCustomGetResponsePredefinedEntry) AsUnion() DLPProfileCustomGetResponsePredefinedEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfileCustomGetResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfileCustomGetResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfileCustomGetResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfileCustomGetResponsePredefinedEntriesExactData] or
// [zero_trust.DLPProfileCustomGetResponsePredefinedEntriesWordList].
type DLPProfileCustomGetResponsePredefinedEntriesUnion interface {
	implementsZeroTrustDLPProfileCustomGetResponsePredefinedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomGetResponsePredefinedEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponsePredefinedEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponsePredefinedEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponsePredefinedEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponsePredefinedEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponsePredefinedEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileCustomGetResponsePredefinedEntriesCustom struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Pattern   DLPProfileCustomGetResponsePredefinedEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileCustomGetResponsePredefinedEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                                 `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                    `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponsePredefinedEntriesCustomJSON    `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesCustomJSON contains the JSON
// metadata for the struct [DLPProfileCustomGetResponsePredefinedEntriesCustom]
type dlpProfileCustomGetResponsePredefinedEntriesCustomJSON struct {
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

func (r *DLPProfileCustomGetResponsePredefinedEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponsePredefinedEntriesCustom) implementsZeroTrustDLPProfileCustomGetResponsePredefinedEntry() {
}

type DLPProfileCustomGetResponsePredefinedEntriesCustomPattern struct {
	Regex      string                                                              `json:"regex,required"`
	Validation DLPProfileCustomGetResponsePredefinedEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileCustomGetResponsePredefinedEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesCustomPatternJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomGetResponsePredefinedEntriesCustomPattern]
type dlpProfileCustomGetResponsePredefinedEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponsePredefinedEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponsePredefinedEntriesCustomPatternValidation string

const (
	DLPProfileCustomGetResponsePredefinedEntriesCustomPatternValidationLuhn DLPProfileCustomGetResponsePredefinedEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesCustomType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesCustomTypeCustom DLPProfileCustomGetResponsePredefinedEntriesCustomType = "custom"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesPredefined struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Type      DLPProfileCustomGetResponsePredefinedEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponsePredefinedEntriesPredefinedJSON `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesPredefinedJSON contains the JSON
// metadata for the struct [DLPProfileCustomGetResponsePredefinedEntriesPredefined]
type dlpProfileCustomGetResponsePredefinedEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponsePredefinedEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponsePredefinedEntriesPredefined) implementsZeroTrustDLPProfileCustomGetResponsePredefinedEntry() {
}

type DLPProfileCustomGetResponsePredefinedEntriesPredefinedType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesPredefinedTypePredefined DLPProfileCustomGetResponsePredefinedEntriesPredefinedType = "predefined"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesIntegration struct {
	ID        string                                                      `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                   `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                        `json:"enabled,required"`
	Name      string                                                      `json:"name,required"`
	Type      DLPProfileCustomGetResponsePredefinedEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                   `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                      `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponsePredefinedEntriesIntegrationJSON `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesIntegrationJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomGetResponsePredefinedEntriesIntegration]
type dlpProfileCustomGetResponsePredefinedEntriesIntegrationJSON struct {
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

func (r *DLPProfileCustomGetResponsePredefinedEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponsePredefinedEntriesIntegration) implementsZeroTrustDLPProfileCustomGetResponsePredefinedEntry() {
}

type DLPProfileCustomGetResponsePredefinedEntriesIntegrationType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesIntegrationTypeIntegration DLPProfileCustomGetResponsePredefinedEntriesIntegrationType = "integration"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesExactData struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Secret    bool                                                      `json:"secret,required"`
	Type      DLPProfileCustomGetResponsePredefinedEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                 `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomGetResponsePredefinedEntriesExactDataJSON `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesExactDataJSON contains the JSON
// metadata for the struct [DLPProfileCustomGetResponsePredefinedEntriesExactData]
type dlpProfileCustomGetResponsePredefinedEntriesExactDataJSON struct {
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

func (r *DLPProfileCustomGetResponsePredefinedEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponsePredefinedEntriesExactData) implementsZeroTrustDLPProfileCustomGetResponsePredefinedEntry() {
}

type DLPProfileCustomGetResponsePredefinedEntriesExactDataType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesExactDataTypeExactData DLPProfileCustomGetResponsePredefinedEntriesExactDataType = "exact_data"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesWordList struct {
	ID        string                                                   `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                     `json:"enabled,required"`
	Name      string                                                   `json:"name,required"`
	Type      DLPProfileCustomGetResponsePredefinedEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                              `json:"word_list,required"`
	ProfileID string                                                   `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponsePredefinedEntriesWordListJSON `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesWordListJSON contains the JSON
// metadata for the struct [DLPProfileCustomGetResponsePredefinedEntriesWordList]
type dlpProfileCustomGetResponsePredefinedEntriesWordListJSON struct {
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

func (r *DLPProfileCustomGetResponsePredefinedEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponsePredefinedEntriesWordList) implementsZeroTrustDLPProfileCustomGetResponsePredefinedEntry() {
}

type DLPProfileCustomGetResponsePredefinedEntriesWordListType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesWordListTypeWordList DLPProfileCustomGetResponsePredefinedEntriesWordListType = "word_list"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesTypeCustom      DLPProfileCustomGetResponsePredefinedEntriesType = "custom"
	DLPProfileCustomGetResponsePredefinedEntriesTypePredefined  DLPProfileCustomGetResponsePredefinedEntriesType = "predefined"
	DLPProfileCustomGetResponsePredefinedEntriesTypeIntegration DLPProfileCustomGetResponsePredefinedEntriesType = "integration"
	DLPProfileCustomGetResponsePredefinedEntriesTypeExactData   DLPProfileCustomGetResponsePredefinedEntriesType = "exact_data"
	DLPProfileCustomGetResponsePredefinedEntriesTypeWordList    DLPProfileCustomGetResponsePredefinedEntriesType = "word_list"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesTypeCustom, DLPProfileCustomGetResponsePredefinedEntriesTypePredefined, DLPProfileCustomGetResponsePredefinedEntriesTypeIntegration, DLPProfileCustomGetResponsePredefinedEntriesTypeExactData, DLPProfileCustomGetResponsePredefinedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedType string

const (
	DLPProfileCustomGetResponsePredefinedTypePredefined DLPProfileCustomGetResponsePredefinedType = "predefined"
)

func (r DLPProfileCustomGetResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegration struct {
	ID        string                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                     `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfileCustomGetResponseIntegrationEntry `json:"entries,required"`
	Name      string                                        `json:"name,required"`
	Type      DLPProfileCustomGetResponseIntegrationType    `json:"type,required"`
	UpdatedAt time.Time                                     `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                                     `json:"description,nullable"`
	JSON        dlpProfileCustomGetResponseIntegrationJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseIntegration]
type dlpProfileCustomGetResponseIntegrationJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegration) implementsZeroTrustDLPProfileCustomGetResponse() {}

type DLPProfileCustomGetResponseIntegrationEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfileCustomGetResponseIntegrationEntriesCustomPattern].
	Pattern   interface{}                                       `json:"pattern,required"`
	ProfileID string                                            `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                         `json:"updated_at" format:"date-time"`
	Type      DLPProfileCustomGetResponseIntegrationEntriesType `json:"type,required"`
	Secret    bool                                              `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                     `json:"word_list,required"`
	JSON     dlpProfileCustomGetResponseIntegrationEntryJSON `json:"-"`
	union    DLPProfileCustomGetResponseIntegrationEntriesUnion
}

// dlpProfileCustomGetResponseIntegrationEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomGetResponseIntegrationEntry]
type dlpProfileCustomGetResponseIntegrationEntryJSON struct {
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

func (r dlpProfileCustomGetResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomGetResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomGetResponseIntegrationEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomGetResponseIntegrationEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfileCustomGetResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfileCustomGetResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfileCustomGetResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfileCustomGetResponseIntegrationEntriesExactData],
// [zero_trust.DLPProfileCustomGetResponseIntegrationEntriesWordList].
func (r DLPProfileCustomGetResponseIntegrationEntry) AsUnion() DLPProfileCustomGetResponseIntegrationEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfileCustomGetResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfileCustomGetResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfileCustomGetResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfileCustomGetResponseIntegrationEntriesExactData] or
// [zero_trust.DLPProfileCustomGetResponseIntegrationEntriesWordList].
type DLPProfileCustomGetResponseIntegrationEntriesUnion interface {
	implementsZeroTrustDLPProfileCustomGetResponseIntegrationEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomGetResponseIntegrationEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseIntegrationEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseIntegrationEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseIntegrationEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseIntegrationEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseIntegrationEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfileCustomGetResponseIntegrationEntriesCustom struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Pattern   DLPProfileCustomGetResponseIntegrationEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfileCustomGetResponseIntegrationEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                                  `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseIntegrationEntriesCustomJSON    `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesCustomJSON contains the JSON
// metadata for the struct [DLPProfileCustomGetResponseIntegrationEntriesCustom]
type dlpProfileCustomGetResponseIntegrationEntriesCustomJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationEntriesCustom) implementsZeroTrustDLPProfileCustomGetResponseIntegrationEntry() {
}

type DLPProfileCustomGetResponseIntegrationEntriesCustomPattern struct {
	Regex      string                                                               `json:"regex,required"`
	Validation DLPProfileCustomGetResponseIntegrationEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfileCustomGetResponseIntegrationEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesCustomPatternJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomGetResponseIntegrationEntriesCustomPattern]
type dlpProfileCustomGetResponseIntegrationEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseIntegrationEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseIntegrationEntriesCustomPatternValidation string

const (
	DLPProfileCustomGetResponseIntegrationEntriesCustomPatternValidationLuhn DLPProfileCustomGetResponseIntegrationEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesCustomType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesCustomTypeCustom DLPProfileCustomGetResponseIntegrationEntriesCustomType = "custom"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesPredefined struct {
	ID        string                                                      `json:"id,required" format:"uuid"`
	Enabled   bool                                                        `json:"enabled,required"`
	Name      string                                                      `json:"name,required"`
	Type      DLPProfileCustomGetResponseIntegrationEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                      `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseIntegrationEntriesPredefinedJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesPredefinedJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomGetResponseIntegrationEntriesPredefined]
type dlpProfileCustomGetResponseIntegrationEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseIntegrationEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationEntriesPredefined) implementsZeroTrustDLPProfileCustomGetResponseIntegrationEntry() {
}

type DLPProfileCustomGetResponseIntegrationEntriesPredefinedType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesPredefinedTypePredefined DLPProfileCustomGetResponseIntegrationEntriesPredefinedType = "predefined"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesIntegration struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Type      DLPProfileCustomGetResponseIntegrationEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseIntegrationEntriesIntegrationJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesIntegrationJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomGetResponseIntegrationEntriesIntegration]
type dlpProfileCustomGetResponseIntegrationEntriesIntegrationJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationEntriesIntegration) implementsZeroTrustDLPProfileCustomGetResponseIntegrationEntry() {
}

type DLPProfileCustomGetResponseIntegrationEntriesIntegrationType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesIntegrationTypeIntegration DLPProfileCustomGetResponseIntegrationEntriesIntegrationType = "integration"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesExactData struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Secret    bool                                                       `json:"secret,required"`
	Type      DLPProfileCustomGetResponseIntegrationEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                  `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomGetResponseIntegrationEntriesExactDataJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesExactDataJSON contains the JSON
// metadata for the struct [DLPProfileCustomGetResponseIntegrationEntriesExactData]
type dlpProfileCustomGetResponseIntegrationEntriesExactDataJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationEntriesExactData) implementsZeroTrustDLPProfileCustomGetResponseIntegrationEntry() {
}

type DLPProfileCustomGetResponseIntegrationEntriesExactDataType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesExactDataTypeExactData DLPProfileCustomGetResponseIntegrationEntriesExactDataType = "exact_data"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesWordList struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Type      DLPProfileCustomGetResponseIntegrationEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                 `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                               `json:"word_list,required"`
	ProfileID string                                                    `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseIntegrationEntriesWordListJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesWordListJSON contains the JSON
// metadata for the struct [DLPProfileCustomGetResponseIntegrationEntriesWordList]
type dlpProfileCustomGetResponseIntegrationEntriesWordListJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationEntriesWordList) implementsZeroTrustDLPProfileCustomGetResponseIntegrationEntry() {
}

type DLPProfileCustomGetResponseIntegrationEntriesWordListType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesWordListTypeWordList DLPProfileCustomGetResponseIntegrationEntriesWordListType = "word_list"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesTypeCustom      DLPProfileCustomGetResponseIntegrationEntriesType = "custom"
	DLPProfileCustomGetResponseIntegrationEntriesTypePredefined  DLPProfileCustomGetResponseIntegrationEntriesType = "predefined"
	DLPProfileCustomGetResponseIntegrationEntriesTypeIntegration DLPProfileCustomGetResponseIntegrationEntriesType = "integration"
	DLPProfileCustomGetResponseIntegrationEntriesTypeExactData   DLPProfileCustomGetResponseIntegrationEntriesType = "exact_data"
	DLPProfileCustomGetResponseIntegrationEntriesTypeWordList    DLPProfileCustomGetResponseIntegrationEntriesType = "word_list"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesTypeCustom, DLPProfileCustomGetResponseIntegrationEntriesTypePredefined, DLPProfileCustomGetResponseIntegrationEntriesTypeIntegration, DLPProfileCustomGetResponseIntegrationEntriesTypeExactData, DLPProfileCustomGetResponseIntegrationEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationType string

const (
	DLPProfileCustomGetResponseIntegrationTypeIntegration DLPProfileCustomGetResponseIntegrationType = "integration"
)

func (r DLPProfileCustomGetResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseType string

const (
	DLPProfileCustomGetResponseTypeCustom      DLPProfileCustomGetResponseType = "custom"
	DLPProfileCustomGetResponseTypePredefined  DLPProfileCustomGetResponseType = "predefined"
	DLPProfileCustomGetResponseTypeIntegration DLPProfileCustomGetResponseType = "integration"
)

func (r DLPProfileCustomGetResponseType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseTypeCustom, DLPProfileCustomGetResponseTypePredefined, DLPProfileCustomGetResponseTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileCustomNewResponseEnvelope struct {
	Errors     []shared.ResponseInfo                         `json:"errors,required"`
	Messages   []shared.ResponseInfo                         `json:"messages,required"`
	Success    bool                                          `json:"success,required"`
	Result     []DLPProfileCustomNewResponse                 `json:"result"`
	ResultInfo DLPProfileCustomNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpProfileCustomNewResponseEnvelopeJSON       `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseEnvelope]
type dlpProfileCustomNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                             `json:"total_count,required"`
	JSON       dlpProfileCustomNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseEnvelopeResultInfo]
type dlpProfileCustomNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileCustomUpdateResponseEnvelope struct {
	Errors     []shared.ResponseInfo                            `json:"errors,required"`
	Messages   []shared.ResponseInfo                            `json:"messages,required"`
	Success    bool                                             `json:"success,required"`
	Result     DLPProfileCustomUpdateResponse                   `json:"result"`
	ResultInfo DLPProfileCustomUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpProfileCustomUpdateResponseEnvelopeJSON       `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponseEnvelope]
type dlpProfileCustomUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                                `json:"total_count,required"`
	JSON       dlpProfileCustomUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DLPProfileCustomUpdateResponseEnvelopeResultInfo]
type dlpProfileCustomUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileCustomDeleteResponseEnvelope struct {
	Errors     []shared.ResponseInfo                            `json:"errors,required"`
	Messages   []shared.ResponseInfo                            `json:"messages,required"`
	Success    bool                                             `json:"success,required"`
	Result     DLPProfileCustomDeleteResponse                   `json:"result,nullable"`
	ResultInfo DLPProfileCustomDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpProfileCustomDeleteResponseEnvelopeJSON       `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomDeleteResponseEnvelope]
type dlpProfileCustomDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomDeleteResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                                `json:"total_count,required"`
	JSON       dlpProfileCustomDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DLPProfileCustomDeleteResponseEnvelopeResultInfo]
type dlpProfileCustomDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileCustomGetResponseEnvelope struct {
	Errors     []shared.ResponseInfo                         `json:"errors,required"`
	Messages   []shared.ResponseInfo                         `json:"messages,required"`
	Success    bool                                          `json:"success,required"`
	Result     DLPProfileCustomGetResponse                   `json:"result"`
	ResultInfo DLPProfileCustomGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpProfileCustomGetResponseEnvelopeJSON       `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseEnvelope]
type dlpProfileCustomGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                             `json:"total_count,required"`
	JSON       dlpProfileCustomGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DLPProfileCustomGetResponseEnvelopeResultInfo]
type dlpProfileCustomGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
