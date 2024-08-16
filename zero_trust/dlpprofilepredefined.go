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
func (r *DLPProfilePredefinedService) Update(ctx context.Context, profileID string, body DLPProfilePredefinedUpdateParams, opts ...option.RequestOption) (res *DLPProfilePredefinedUpdateResponse, err error) {
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
func (r *DLPProfilePredefinedService) Get(ctx context.Context, profileID string, query DLPProfilePredefinedGetParams, opts ...option.RequestOption) (res *DLPProfilePredefinedGetResponse, err error) {
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

type DLPProfilePredefinedUpdateResponse struct {
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
	// [[]DLPProfilePredefinedUpdateResponseCustomEntry],
	// [[]DLPProfilePredefinedUpdateResponsePredefinedEntry],
	// [[]DLPProfilePredefinedUpdateResponseIntegrationEntry].
	Entries interface{} `json:"entries"`
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// The name of the profile
	Name       string `json:"name,required"`
	OCREnabled bool   `json:"ocr_enabled"`
	// When the profile was lasted updated
	UpdatedAt  time.Time                              `json:"updated_at" format:"date-time"`
	Type       DLPProfilePredefinedUpdateResponseType `json:"type,required"`
	OpenAccess bool                                   `json:"open_access"`
	JSON       dlpProfilePredefinedUpdateResponseJSON `json:"-"`
	union      DLPProfilePredefinedUpdateResponseUnion
}

// dlpProfilePredefinedUpdateResponseJSON contains the JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponse]
type dlpProfilePredefinedUpdateResponseJSON struct {
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

func (r dlpProfilePredefinedUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfilePredefinedUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfilePredefinedUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfilePredefinedUpdateResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfilePredefinedUpdateResponseCustom],
// [zero_trust.DLPProfilePredefinedUpdateResponsePredefined],
// [zero_trust.DLPProfilePredefinedUpdateResponseIntegration].
func (r DLPProfilePredefinedUpdateResponse) AsUnion() DLPProfilePredefinedUpdateResponseUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfilePredefinedUpdateResponseCustom],
// [zero_trust.DLPProfilePredefinedUpdateResponsePredefined] or
// [zero_trust.DLPProfilePredefinedUpdateResponseIntegration].
type DLPProfilePredefinedUpdateResponseUnion interface {
	implementsZeroTrustDLPProfilePredefinedUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfilePredefinedUpdateResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
	)
}

type DLPProfilePredefinedUpdateResponseCustom struct {
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness,required"`
	// When the profile was created
	CreatedAt time.Time                                       `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfilePredefinedUpdateResponseCustomEntry `json:"entries,required"`
	// The name of the profile
	Name       string                                       `json:"name,required"`
	OCREnabled bool                                         `json:"ocr_enabled,required"`
	Type       DLPProfilePredefinedUpdateResponseCustomType `json:"type,required"`
	// When the profile was lasted updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                                       `json:"description,nullable"`
	JSON        dlpProfilePredefinedUpdateResponseCustomJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseCustomJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedUpdateResponseCustom]
type dlpProfilePredefinedUpdateResponseCustomJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseCustom) implementsZeroTrustDLPProfilePredefinedUpdateResponse() {
}

type DLPProfilePredefinedUpdateResponseCustomEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfilePredefinedUpdateResponseCustomEntriesCustomPattern].
	Pattern   interface{}                                         `json:"pattern,required"`
	ProfileID string                                              `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                           `json:"updated_at" format:"date-time"`
	Type      DLPProfilePredefinedUpdateResponseCustomEntriesType `json:"type,required"`
	Secret    bool                                                `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                       `json:"word_list,required"`
	JSON     dlpProfilePredefinedUpdateResponseCustomEntryJSON `json:"-"`
	union    DLPProfilePredefinedUpdateResponseCustomEntriesUnion
}

// dlpProfilePredefinedUpdateResponseCustomEntryJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedUpdateResponseCustomEntry]
type dlpProfilePredefinedUpdateResponseCustomEntryJSON struct {
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

func (r dlpProfilePredefinedUpdateResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfilePredefinedUpdateResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfilePredefinedUpdateResponseCustomEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfilePredefinedUpdateResponseCustomEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfilePredefinedUpdateResponseCustomEntriesCustom],
// [zero_trust.DLPProfilePredefinedUpdateResponseCustomEntriesPredefined],
// [zero_trust.DLPProfilePredefinedUpdateResponseCustomEntriesIntegration],
// [zero_trust.DLPProfilePredefinedUpdateResponseCustomEntriesExactData],
// [zero_trust.DLPProfilePredefinedUpdateResponseCustomEntriesWordList].
func (r DLPProfilePredefinedUpdateResponseCustomEntry) AsUnion() DLPProfilePredefinedUpdateResponseCustomEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfilePredefinedUpdateResponseCustomEntriesCustom],
// [zero_trust.DLPProfilePredefinedUpdateResponseCustomEntriesPredefined],
// [zero_trust.DLPProfilePredefinedUpdateResponseCustomEntriesIntegration],
// [zero_trust.DLPProfilePredefinedUpdateResponseCustomEntriesExactData] or
// [zero_trust.DLPProfilePredefinedUpdateResponseCustomEntriesWordList].
type DLPProfilePredefinedUpdateResponseCustomEntriesUnion interface {
	implementsZeroTrustDLPProfilePredefinedUpdateResponseCustomEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfilePredefinedUpdateResponseCustomEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponseCustomEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponseCustomEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponseCustomEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponseCustomEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponseCustomEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfilePredefinedUpdateResponseCustomEntriesCustom struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Pattern   DLPProfilePredefinedUpdateResponseCustomEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfilePredefinedUpdateResponseCustomEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponseCustomEntriesCustomJSON    `json:"-"`
}

// dlpProfilePredefinedUpdateResponseCustomEntriesCustomJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedUpdateResponseCustomEntriesCustom]
type dlpProfilePredefinedUpdateResponseCustomEntriesCustomJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseCustomEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseCustomEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseCustomEntriesCustom) implementsZeroTrustDLPProfilePredefinedUpdateResponseCustomEntry() {
}

type DLPProfilePredefinedUpdateResponseCustomEntriesCustomPattern struct {
	Regex      string                                                                 `json:"regex,required"`
	Validation DLPProfilePredefinedUpdateResponseCustomEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfilePredefinedUpdateResponseCustomEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfilePredefinedUpdateResponseCustomEntriesCustomPatternJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponseCustomEntriesCustomPattern]
type dlpProfilePredefinedUpdateResponseCustomEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseCustomEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseCustomEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseCustomEntriesCustomPatternValidation string

const (
	DLPProfilePredefinedUpdateResponseCustomEntriesCustomPatternValidationLuhn DLPProfilePredefinedUpdateResponseCustomEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfilePredefinedUpdateResponseCustomEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseCustomEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseCustomEntriesCustomType string

const (
	DLPProfilePredefinedUpdateResponseCustomEntriesCustomTypeCustom DLPProfilePredefinedUpdateResponseCustomEntriesCustomType = "custom"
)

func (r DLPProfilePredefinedUpdateResponseCustomEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseCustomEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseCustomEntriesPredefined struct {
	ID        string                                                        `json:"id,required" format:"uuid"`
	Enabled   bool                                                          `json:"enabled,required"`
	Name      string                                                        `json:"name,required"`
	Type      DLPProfilePredefinedUpdateResponseCustomEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponseCustomEntriesPredefinedJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseCustomEntriesPredefinedJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedUpdateResponseCustomEntriesPredefined]
type dlpProfilePredefinedUpdateResponseCustomEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseCustomEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseCustomEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseCustomEntriesPredefined) implementsZeroTrustDLPProfilePredefinedUpdateResponseCustomEntry() {
}

type DLPProfilePredefinedUpdateResponseCustomEntriesPredefinedType string

const (
	DLPProfilePredefinedUpdateResponseCustomEntriesPredefinedTypePredefined DLPProfilePredefinedUpdateResponseCustomEntriesPredefinedType = "predefined"
)

func (r DLPProfilePredefinedUpdateResponseCustomEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseCustomEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseCustomEntriesIntegration struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                           `json:"enabled,required"`
	Name      string                                                         `json:"name,required"`
	Type      DLPProfilePredefinedUpdateResponseCustomEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                      `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponseCustomEntriesIntegrationJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseCustomEntriesIntegrationJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedUpdateResponseCustomEntriesIntegration]
type dlpProfilePredefinedUpdateResponseCustomEntriesIntegrationJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseCustomEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseCustomEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseCustomEntriesIntegration) implementsZeroTrustDLPProfilePredefinedUpdateResponseCustomEntry() {
}

type DLPProfilePredefinedUpdateResponseCustomEntriesIntegrationType string

const (
	DLPProfilePredefinedUpdateResponseCustomEntriesIntegrationTypeIntegration DLPProfilePredefinedUpdateResponseCustomEntriesIntegrationType = "integration"
)

func (r DLPProfilePredefinedUpdateResponseCustomEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseCustomEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseCustomEntriesExactData struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Secret    bool                                                         `json:"secret,required"`
	Type      DLPProfilePredefinedUpdateResponseCustomEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfilePredefinedUpdateResponseCustomEntriesExactDataJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseCustomEntriesExactDataJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedUpdateResponseCustomEntriesExactData]
type dlpProfilePredefinedUpdateResponseCustomEntriesExactDataJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseCustomEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseCustomEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseCustomEntriesExactData) implementsZeroTrustDLPProfilePredefinedUpdateResponseCustomEntry() {
}

type DLPProfilePredefinedUpdateResponseCustomEntriesExactDataType string

const (
	DLPProfilePredefinedUpdateResponseCustomEntriesExactDataTypeExactData DLPProfilePredefinedUpdateResponseCustomEntriesExactDataType = "exact_data"
)

func (r DLPProfilePredefinedUpdateResponseCustomEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseCustomEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseCustomEntriesWordList struct {
	ID        string                                                      `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                   `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                        `json:"enabled,required"`
	Name      string                                                      `json:"name,required"`
	Type      DLPProfilePredefinedUpdateResponseCustomEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                   `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                 `json:"word_list,required"`
	ProfileID string                                                      `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponseCustomEntriesWordListJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseCustomEntriesWordListJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedUpdateResponseCustomEntriesWordList]
type dlpProfilePredefinedUpdateResponseCustomEntriesWordListJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseCustomEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseCustomEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseCustomEntriesWordList) implementsZeroTrustDLPProfilePredefinedUpdateResponseCustomEntry() {
}

type DLPProfilePredefinedUpdateResponseCustomEntriesWordListType string

const (
	DLPProfilePredefinedUpdateResponseCustomEntriesWordListTypeWordList DLPProfilePredefinedUpdateResponseCustomEntriesWordListType = "word_list"
)

func (r DLPProfilePredefinedUpdateResponseCustomEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseCustomEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseCustomEntriesType string

const (
	DLPProfilePredefinedUpdateResponseCustomEntriesTypeCustom      DLPProfilePredefinedUpdateResponseCustomEntriesType = "custom"
	DLPProfilePredefinedUpdateResponseCustomEntriesTypePredefined  DLPProfilePredefinedUpdateResponseCustomEntriesType = "predefined"
	DLPProfilePredefinedUpdateResponseCustomEntriesTypeIntegration DLPProfilePredefinedUpdateResponseCustomEntriesType = "integration"
	DLPProfilePredefinedUpdateResponseCustomEntriesTypeExactData   DLPProfilePredefinedUpdateResponseCustomEntriesType = "exact_data"
	DLPProfilePredefinedUpdateResponseCustomEntriesTypeWordList    DLPProfilePredefinedUpdateResponseCustomEntriesType = "word_list"
)

func (r DLPProfilePredefinedUpdateResponseCustomEntriesType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseCustomEntriesTypeCustom, DLPProfilePredefinedUpdateResponseCustomEntriesTypePredefined, DLPProfilePredefinedUpdateResponseCustomEntriesTypeIntegration, DLPProfilePredefinedUpdateResponseCustomEntriesTypeExactData, DLPProfilePredefinedUpdateResponseCustomEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseCustomType string

const (
	DLPProfilePredefinedUpdateResponseCustomTypeCustom DLPProfilePredefinedUpdateResponseCustomType = "custom"
)

func (r DLPProfilePredefinedUpdateResponseCustomType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponsePredefined struct {
	// The id of the predefined profile (uuid)
	ID                string                                              `json:"id,required" format:"uuid"`
	AllowedMatchCount int64                                               `json:"allowed_match_count,required"`
	Entries           []DLPProfilePredefinedUpdateResponsePredefinedEntry `json:"entries,required"`
	// The name of the predefined profile
	Name string                                           `json:"name,required"`
	Type DLPProfilePredefinedUpdateResponsePredefinedType `json:"type,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness                                 `json:"context_awareness"`
	OCREnabled       bool                                             `json:"ocr_enabled"`
	OpenAccess       bool                                             `json:"open_access"`
	JSON             dlpProfilePredefinedUpdateResponsePredefinedJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponsePredefinedJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedUpdateResponsePredefined]
type dlpProfilePredefinedUpdateResponsePredefinedJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponsePredefined) implementsZeroTrustDLPProfilePredefinedUpdateResponse() {
}

type DLPProfilePredefinedUpdateResponsePredefinedEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomPattern].
	Pattern   interface{}                                             `json:"pattern,required"`
	ProfileID string                                                  `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                               `json:"updated_at" format:"date-time"`
	Type      DLPProfilePredefinedUpdateResponsePredefinedEntriesType `json:"type,required"`
	Secret    bool                                                    `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                           `json:"word_list,required"`
	JSON     dlpProfilePredefinedUpdateResponsePredefinedEntryJSON `json:"-"`
	union    DLPProfilePredefinedUpdateResponsePredefinedEntriesUnion
}

// dlpProfilePredefinedUpdateResponsePredefinedEntryJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedUpdateResponsePredefinedEntry]
type dlpProfilePredefinedUpdateResponsePredefinedEntryJSON struct {
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

func (r dlpProfilePredefinedUpdateResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfilePredefinedUpdateResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfilePredefinedUpdateResponsePredefinedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfilePredefinedUpdateResponsePredefinedEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfilePredefinedUpdateResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfilePredefinedUpdateResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfilePredefinedUpdateResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfilePredefinedUpdateResponsePredefinedEntriesExactData],
// [zero_trust.DLPProfilePredefinedUpdateResponsePredefinedEntriesWordList].
func (r DLPProfilePredefinedUpdateResponsePredefinedEntry) AsUnion() DLPProfilePredefinedUpdateResponsePredefinedEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfilePredefinedUpdateResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfilePredefinedUpdateResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfilePredefinedUpdateResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfilePredefinedUpdateResponsePredefinedEntriesExactData] or
// [zero_trust.DLPProfilePredefinedUpdateResponsePredefinedEntriesWordList].
type DLPProfilePredefinedUpdateResponsePredefinedEntriesUnion interface {
	implementsZeroTrustDLPProfilePredefinedUpdateResponsePredefinedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfilePredefinedUpdateResponsePredefinedEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponsePredefinedEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponsePredefinedEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponsePredefinedEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponsePredefinedEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponsePredefinedEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfilePredefinedUpdateResponsePredefinedEntriesCustom struct {
	ID        string                                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                             `json:"enabled,required"`
	Name      string                                                           `json:"name,required"`
	Pattern   DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                                        `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponsePredefinedEntriesCustomJSON    `json:"-"`
}

// dlpProfilePredefinedUpdateResponsePredefinedEntriesCustomJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedUpdateResponsePredefinedEntriesCustom]
type dlpProfilePredefinedUpdateResponsePredefinedEntriesCustomJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponsePredefinedEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponsePredefinedEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponsePredefinedEntriesCustom) implementsZeroTrustDLPProfilePredefinedUpdateResponsePredefinedEntry() {
}

type DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomPattern struct {
	Regex      string                                                                     `json:"regex,required"`
	Validation DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfilePredefinedUpdateResponsePredefinedEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfilePredefinedUpdateResponsePredefinedEntriesCustomPatternJSON contains
// the JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomPattern]
type dlpProfilePredefinedUpdateResponsePredefinedEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponsePredefinedEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomPatternValidation string

const (
	DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomPatternValidationLuhn DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomType string

const (
	DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomTypeCustom DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomType = "custom"
)

func (r DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponsePredefinedEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponsePredefinedEntriesPredefined struct {
	ID        string                                                            `json:"id,required" format:"uuid"`
	Enabled   bool                                                              `json:"enabled,required"`
	Name      string                                                            `json:"name,required"`
	Type      DLPProfilePredefinedUpdateResponsePredefinedEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                            `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponsePredefinedEntriesPredefinedJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponsePredefinedEntriesPredefinedJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponsePredefinedEntriesPredefined]
type dlpProfilePredefinedUpdateResponsePredefinedEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponsePredefinedEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponsePredefinedEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponsePredefinedEntriesPredefined) implementsZeroTrustDLPProfilePredefinedUpdateResponsePredefinedEntry() {
}

type DLPProfilePredefinedUpdateResponsePredefinedEntriesPredefinedType string

const (
	DLPProfilePredefinedUpdateResponsePredefinedEntriesPredefinedTypePredefined DLPProfilePredefinedUpdateResponsePredefinedEntriesPredefinedType = "predefined"
)

func (r DLPProfilePredefinedUpdateResponsePredefinedEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponsePredefinedEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponsePredefinedEntriesIntegration struct {
	ID        string                                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                               `json:"enabled,required"`
	Name      string                                                             `json:"name,required"`
	Type      DLPProfilePredefinedUpdateResponsePredefinedEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                          `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponsePredefinedEntriesIntegrationJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponsePredefinedEntriesIntegrationJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponsePredefinedEntriesIntegration]
type dlpProfilePredefinedUpdateResponsePredefinedEntriesIntegrationJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponsePredefinedEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponsePredefinedEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponsePredefinedEntriesIntegration) implementsZeroTrustDLPProfilePredefinedUpdateResponsePredefinedEntry() {
}

type DLPProfilePredefinedUpdateResponsePredefinedEntriesIntegrationType string

const (
	DLPProfilePredefinedUpdateResponsePredefinedEntriesIntegrationTypeIntegration DLPProfilePredefinedUpdateResponsePredefinedEntriesIntegrationType = "integration"
)

func (r DLPProfilePredefinedUpdateResponsePredefinedEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponsePredefinedEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponsePredefinedEntriesExactData struct {
	ID        string                                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                             `json:"enabled,required"`
	Name      string                                                           `json:"name,required"`
	Secret    bool                                                             `json:"secret,required"`
	Type      DLPProfilePredefinedUpdateResponsePredefinedEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                        `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfilePredefinedUpdateResponsePredefinedEntriesExactDataJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponsePredefinedEntriesExactDataJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponsePredefinedEntriesExactData]
type dlpProfilePredefinedUpdateResponsePredefinedEntriesExactDataJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponsePredefinedEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponsePredefinedEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponsePredefinedEntriesExactData) implementsZeroTrustDLPProfilePredefinedUpdateResponsePredefinedEntry() {
}

type DLPProfilePredefinedUpdateResponsePredefinedEntriesExactDataType string

const (
	DLPProfilePredefinedUpdateResponsePredefinedEntriesExactDataTypeExactData DLPProfilePredefinedUpdateResponsePredefinedEntriesExactDataType = "exact_data"
)

func (r DLPProfilePredefinedUpdateResponsePredefinedEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponsePredefinedEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponsePredefinedEntriesWordList struct {
	ID        string                                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                            `json:"enabled,required"`
	Name      string                                                          `json:"name,required"`
	Type      DLPProfilePredefinedUpdateResponsePredefinedEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                       `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                     `json:"word_list,required"`
	ProfileID string                                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponsePredefinedEntriesWordListJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponsePredefinedEntriesWordListJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponsePredefinedEntriesWordList]
type dlpProfilePredefinedUpdateResponsePredefinedEntriesWordListJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponsePredefinedEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponsePredefinedEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponsePredefinedEntriesWordList) implementsZeroTrustDLPProfilePredefinedUpdateResponsePredefinedEntry() {
}

type DLPProfilePredefinedUpdateResponsePredefinedEntriesWordListType string

const (
	DLPProfilePredefinedUpdateResponsePredefinedEntriesWordListTypeWordList DLPProfilePredefinedUpdateResponsePredefinedEntriesWordListType = "word_list"
)

func (r DLPProfilePredefinedUpdateResponsePredefinedEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponsePredefinedEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponsePredefinedEntriesType string

const (
	DLPProfilePredefinedUpdateResponsePredefinedEntriesTypeCustom      DLPProfilePredefinedUpdateResponsePredefinedEntriesType = "custom"
	DLPProfilePredefinedUpdateResponsePredefinedEntriesTypePredefined  DLPProfilePredefinedUpdateResponsePredefinedEntriesType = "predefined"
	DLPProfilePredefinedUpdateResponsePredefinedEntriesTypeIntegration DLPProfilePredefinedUpdateResponsePredefinedEntriesType = "integration"
	DLPProfilePredefinedUpdateResponsePredefinedEntriesTypeExactData   DLPProfilePredefinedUpdateResponsePredefinedEntriesType = "exact_data"
	DLPProfilePredefinedUpdateResponsePredefinedEntriesTypeWordList    DLPProfilePredefinedUpdateResponsePredefinedEntriesType = "word_list"
)

func (r DLPProfilePredefinedUpdateResponsePredefinedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponsePredefinedEntriesTypeCustom, DLPProfilePredefinedUpdateResponsePredefinedEntriesTypePredefined, DLPProfilePredefinedUpdateResponsePredefinedEntriesTypeIntegration, DLPProfilePredefinedUpdateResponsePredefinedEntriesTypeExactData, DLPProfilePredefinedUpdateResponsePredefinedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponsePredefinedType string

const (
	DLPProfilePredefinedUpdateResponsePredefinedTypePredefined DLPProfilePredefinedUpdateResponsePredefinedType = "predefined"
)

func (r DLPProfilePredefinedUpdateResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseIntegration struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfilePredefinedUpdateResponseIntegrationEntry `json:"entries,required"`
	Name      string                                               `json:"name,required"`
	Type      DLPProfilePredefinedUpdateResponseIntegrationType    `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                                            `json:"description,nullable"`
	JSON        dlpProfilePredefinedUpdateResponseIntegrationJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseIntegrationJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedUpdateResponseIntegration]
type dlpProfilePredefinedUpdateResponseIntegrationJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseIntegration) implementsZeroTrustDLPProfilePredefinedUpdateResponse() {
}

type DLPProfilePredefinedUpdateResponseIntegrationEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomPattern].
	Pattern   interface{}                                              `json:"pattern,required"`
	ProfileID string                                                   `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                                `json:"updated_at" format:"date-time"`
	Type      DLPProfilePredefinedUpdateResponseIntegrationEntriesType `json:"type,required"`
	Secret    bool                                                     `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                            `json:"word_list,required"`
	JSON     dlpProfilePredefinedUpdateResponseIntegrationEntryJSON `json:"-"`
	union    DLPProfilePredefinedUpdateResponseIntegrationEntriesUnion
}

// dlpProfilePredefinedUpdateResponseIntegrationEntryJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedUpdateResponseIntegrationEntry]
type dlpProfilePredefinedUpdateResponseIntegrationEntryJSON struct {
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

func (r dlpProfilePredefinedUpdateResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfilePredefinedUpdateResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfilePredefinedUpdateResponseIntegrationEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfilePredefinedUpdateResponseIntegrationEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfilePredefinedUpdateResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfilePredefinedUpdateResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfilePredefinedUpdateResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfilePredefinedUpdateResponseIntegrationEntriesExactData],
// [zero_trust.DLPProfilePredefinedUpdateResponseIntegrationEntriesWordList].
func (r DLPProfilePredefinedUpdateResponseIntegrationEntry) AsUnion() DLPProfilePredefinedUpdateResponseIntegrationEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfilePredefinedUpdateResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfilePredefinedUpdateResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfilePredefinedUpdateResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfilePredefinedUpdateResponseIntegrationEntriesExactData] or
// [zero_trust.DLPProfilePredefinedUpdateResponseIntegrationEntriesWordList].
type DLPProfilePredefinedUpdateResponseIntegrationEntriesUnion interface {
	implementsZeroTrustDLPProfilePredefinedUpdateResponseIntegrationEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfilePredefinedUpdateResponseIntegrationEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponseIntegrationEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponseIntegrationEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponseIntegrationEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponseIntegrationEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedUpdateResponseIntegrationEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfilePredefinedUpdateResponseIntegrationEntriesCustom struct {
	ID        string                                                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                         `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                              `json:"enabled,required"`
	Name      string                                                            `json:"name,required"`
	Pattern   DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                                         `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                            `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponseIntegrationEntriesCustomJSON    `json:"-"`
}

// dlpProfilePredefinedUpdateResponseIntegrationEntriesCustomJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedUpdateResponseIntegrationEntriesCustom]
type dlpProfilePredefinedUpdateResponseIntegrationEntriesCustomJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseIntegrationEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseIntegrationEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseIntegrationEntriesCustom) implementsZeroTrustDLPProfilePredefinedUpdateResponseIntegrationEntry() {
}

type DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomPattern struct {
	Regex      string                                                                      `json:"regex,required"`
	Validation DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfilePredefinedUpdateResponseIntegrationEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfilePredefinedUpdateResponseIntegrationEntriesCustomPatternJSON contains
// the JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomPattern]
type dlpProfilePredefinedUpdateResponseIntegrationEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseIntegrationEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomPatternValidation string

const (
	DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomPatternValidationLuhn DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomType string

const (
	DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomTypeCustom DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomType = "custom"
)

func (r DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseIntegrationEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseIntegrationEntriesPredefined struct {
	ID        string                                                             `json:"id,required" format:"uuid"`
	Enabled   bool                                                               `json:"enabled,required"`
	Name      string                                                             `json:"name,required"`
	Type      DLPProfilePredefinedUpdateResponseIntegrationEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponseIntegrationEntriesPredefinedJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseIntegrationEntriesPredefinedJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponseIntegrationEntriesPredefined]
type dlpProfilePredefinedUpdateResponseIntegrationEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseIntegrationEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseIntegrationEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseIntegrationEntriesPredefined) implementsZeroTrustDLPProfilePredefinedUpdateResponseIntegrationEntry() {
}

type DLPProfilePredefinedUpdateResponseIntegrationEntriesPredefinedType string

const (
	DLPProfilePredefinedUpdateResponseIntegrationEntriesPredefinedTypePredefined DLPProfilePredefinedUpdateResponseIntegrationEntriesPredefinedType = "predefined"
)

func (r DLPProfilePredefinedUpdateResponseIntegrationEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseIntegrationEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseIntegrationEntriesIntegration struct {
	ID        string                                                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                `json:"enabled,required"`
	Name      string                                                              `json:"name,required"`
	Type      DLPProfilePredefinedUpdateResponseIntegrationEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                           `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                              `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponseIntegrationEntriesIntegrationJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseIntegrationEntriesIntegrationJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponseIntegrationEntriesIntegration]
type dlpProfilePredefinedUpdateResponseIntegrationEntriesIntegrationJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseIntegrationEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseIntegrationEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseIntegrationEntriesIntegration) implementsZeroTrustDLPProfilePredefinedUpdateResponseIntegrationEntry() {
}

type DLPProfilePredefinedUpdateResponseIntegrationEntriesIntegrationType string

const (
	DLPProfilePredefinedUpdateResponseIntegrationEntriesIntegrationTypeIntegration DLPProfilePredefinedUpdateResponseIntegrationEntriesIntegrationType = "integration"
)

func (r DLPProfilePredefinedUpdateResponseIntegrationEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseIntegrationEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseIntegrationEntriesExactData struct {
	ID        string                                                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                         `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                              `json:"enabled,required"`
	Name      string                                                            `json:"name,required"`
	Secret    bool                                                              `json:"secret,required"`
	Type      DLPProfilePredefinedUpdateResponseIntegrationEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                         `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfilePredefinedUpdateResponseIntegrationEntriesExactDataJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseIntegrationEntriesExactDataJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponseIntegrationEntriesExactData]
type dlpProfilePredefinedUpdateResponseIntegrationEntriesExactDataJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseIntegrationEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseIntegrationEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseIntegrationEntriesExactData) implementsZeroTrustDLPProfilePredefinedUpdateResponseIntegrationEntry() {
}

type DLPProfilePredefinedUpdateResponseIntegrationEntriesExactDataType string

const (
	DLPProfilePredefinedUpdateResponseIntegrationEntriesExactDataTypeExactData DLPProfilePredefinedUpdateResponseIntegrationEntriesExactDataType = "exact_data"
)

func (r DLPProfilePredefinedUpdateResponseIntegrationEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseIntegrationEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseIntegrationEntriesWordList struct {
	ID        string                                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                             `json:"enabled,required"`
	Name      string                                                           `json:"name,required"`
	Type      DLPProfilePredefinedUpdateResponseIntegrationEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                        `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                      `json:"word_list,required"`
	ProfileID string                                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponseIntegrationEntriesWordListJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseIntegrationEntriesWordListJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponseIntegrationEntriesWordList]
type dlpProfilePredefinedUpdateResponseIntegrationEntriesWordListJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseIntegrationEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseIntegrationEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseIntegrationEntriesWordList) implementsZeroTrustDLPProfilePredefinedUpdateResponseIntegrationEntry() {
}

type DLPProfilePredefinedUpdateResponseIntegrationEntriesWordListType string

const (
	DLPProfilePredefinedUpdateResponseIntegrationEntriesWordListTypeWordList DLPProfilePredefinedUpdateResponseIntegrationEntriesWordListType = "word_list"
)

func (r DLPProfilePredefinedUpdateResponseIntegrationEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseIntegrationEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseIntegrationEntriesType string

const (
	DLPProfilePredefinedUpdateResponseIntegrationEntriesTypeCustom      DLPProfilePredefinedUpdateResponseIntegrationEntriesType = "custom"
	DLPProfilePredefinedUpdateResponseIntegrationEntriesTypePredefined  DLPProfilePredefinedUpdateResponseIntegrationEntriesType = "predefined"
	DLPProfilePredefinedUpdateResponseIntegrationEntriesTypeIntegration DLPProfilePredefinedUpdateResponseIntegrationEntriesType = "integration"
	DLPProfilePredefinedUpdateResponseIntegrationEntriesTypeExactData   DLPProfilePredefinedUpdateResponseIntegrationEntriesType = "exact_data"
	DLPProfilePredefinedUpdateResponseIntegrationEntriesTypeWordList    DLPProfilePredefinedUpdateResponseIntegrationEntriesType = "word_list"
)

func (r DLPProfilePredefinedUpdateResponseIntegrationEntriesType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseIntegrationEntriesTypeCustom, DLPProfilePredefinedUpdateResponseIntegrationEntriesTypePredefined, DLPProfilePredefinedUpdateResponseIntegrationEntriesTypeIntegration, DLPProfilePredefinedUpdateResponseIntegrationEntriesTypeExactData, DLPProfilePredefinedUpdateResponseIntegrationEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseIntegrationType string

const (
	DLPProfilePredefinedUpdateResponseIntegrationTypeIntegration DLPProfilePredefinedUpdateResponseIntegrationType = "integration"
)

func (r DLPProfilePredefinedUpdateResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseType string

const (
	DLPProfilePredefinedUpdateResponseTypeCustom      DLPProfilePredefinedUpdateResponseType = "custom"
	DLPProfilePredefinedUpdateResponseTypePredefined  DLPProfilePredefinedUpdateResponseType = "predefined"
	DLPProfilePredefinedUpdateResponseTypeIntegration DLPProfilePredefinedUpdateResponseType = "integration"
)

func (r DLPProfilePredefinedUpdateResponseType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseTypeCustom, DLPProfilePredefinedUpdateResponseTypePredefined, DLPProfilePredefinedUpdateResponseTypeIntegration:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponse struct {
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
	// [[]DLPProfilePredefinedGetResponseCustomEntry],
	// [[]DLPProfilePredefinedGetResponsePredefinedEntry],
	// [[]DLPProfilePredefinedGetResponseIntegrationEntry].
	Entries interface{} `json:"entries"`
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// The name of the profile
	Name       string `json:"name,required"`
	OCREnabled bool   `json:"ocr_enabled"`
	// When the profile was lasted updated
	UpdatedAt  time.Time                           `json:"updated_at" format:"date-time"`
	Type       DLPProfilePredefinedGetResponseType `json:"type,required"`
	OpenAccess bool                                `json:"open_access"`
	JSON       dlpProfilePredefinedGetResponseJSON `json:"-"`
	union      DLPProfilePredefinedGetResponseUnion
}

// dlpProfilePredefinedGetResponseJSON contains the JSON metadata for the struct
// [DLPProfilePredefinedGetResponse]
type dlpProfilePredefinedGetResponseJSON struct {
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

func (r dlpProfilePredefinedGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfilePredefinedGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfilePredefinedGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfilePredefinedGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfilePredefinedGetResponseCustom],
// [zero_trust.DLPProfilePredefinedGetResponsePredefined],
// [zero_trust.DLPProfilePredefinedGetResponseIntegration].
func (r DLPProfilePredefinedGetResponse) AsUnion() DLPProfilePredefinedGetResponseUnion {
	return r.union
}

// Union satisfied by [zero_trust.DLPProfilePredefinedGetResponseCustom],
// [zero_trust.DLPProfilePredefinedGetResponsePredefined] or
// [zero_trust.DLPProfilePredefinedGetResponseIntegration].
type DLPProfilePredefinedGetResponseUnion interface {
	implementsZeroTrustDLPProfilePredefinedGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfilePredefinedGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
	)
}

type DLPProfilePredefinedGetResponseCustom struct {
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness,required"`
	// When the profile was created
	CreatedAt time.Time                                    `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfilePredefinedGetResponseCustomEntry `json:"entries,required"`
	// The name of the profile
	Name       string                                    `json:"name,required"`
	OCREnabled bool                                      `json:"ocr_enabled,required"`
	Type       DLPProfilePredefinedGetResponseCustomType `json:"type,required"`
	// When the profile was lasted updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                                    `json:"description,nullable"`
	JSON        dlpProfilePredefinedGetResponseCustomJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseCustomJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedGetResponseCustom]
type dlpProfilePredefinedGetResponseCustomJSON struct {
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

func (r *DLPProfilePredefinedGetResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseCustom) implementsZeroTrustDLPProfilePredefinedGetResponse() {}

type DLPProfilePredefinedGetResponseCustomEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfilePredefinedGetResponseCustomEntriesCustomPattern].
	Pattern   interface{}                                      `json:"pattern,required"`
	ProfileID string                                           `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	Type      DLPProfilePredefinedGetResponseCustomEntriesType `json:"type,required"`
	Secret    bool                                             `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                    `json:"word_list,required"`
	JSON     dlpProfilePredefinedGetResponseCustomEntryJSON `json:"-"`
	union    DLPProfilePredefinedGetResponseCustomEntriesUnion
}

// dlpProfilePredefinedGetResponseCustomEntryJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedGetResponseCustomEntry]
type dlpProfilePredefinedGetResponseCustomEntryJSON struct {
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

func (r dlpProfilePredefinedGetResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfilePredefinedGetResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfilePredefinedGetResponseCustomEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfilePredefinedGetResponseCustomEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfilePredefinedGetResponseCustomEntriesCustom],
// [zero_trust.DLPProfilePredefinedGetResponseCustomEntriesPredefined],
// [zero_trust.DLPProfilePredefinedGetResponseCustomEntriesIntegration],
// [zero_trust.DLPProfilePredefinedGetResponseCustomEntriesExactData],
// [zero_trust.DLPProfilePredefinedGetResponseCustomEntriesWordList].
func (r DLPProfilePredefinedGetResponseCustomEntry) AsUnion() DLPProfilePredefinedGetResponseCustomEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfilePredefinedGetResponseCustomEntriesCustom],
// [zero_trust.DLPProfilePredefinedGetResponseCustomEntriesPredefined],
// [zero_trust.DLPProfilePredefinedGetResponseCustomEntriesIntegration],
// [zero_trust.DLPProfilePredefinedGetResponseCustomEntriesExactData] or
// [zero_trust.DLPProfilePredefinedGetResponseCustomEntriesWordList].
type DLPProfilePredefinedGetResponseCustomEntriesUnion interface {
	implementsZeroTrustDLPProfilePredefinedGetResponseCustomEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfilePredefinedGetResponseCustomEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponseCustomEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponseCustomEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponseCustomEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponseCustomEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponseCustomEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfilePredefinedGetResponseCustomEntriesCustom struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Pattern   DLPProfilePredefinedGetResponseCustomEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfilePredefinedGetResponseCustomEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                                 `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                    `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponseCustomEntriesCustomJSON    `json:"-"`
}

// dlpProfilePredefinedGetResponseCustomEntriesCustomJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponseCustomEntriesCustom]
type dlpProfilePredefinedGetResponseCustomEntriesCustomJSON struct {
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

func (r *DLPProfilePredefinedGetResponseCustomEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseCustomEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseCustomEntriesCustom) implementsZeroTrustDLPProfilePredefinedGetResponseCustomEntry() {
}

type DLPProfilePredefinedGetResponseCustomEntriesCustomPattern struct {
	Regex      string                                                              `json:"regex,required"`
	Validation DLPProfilePredefinedGetResponseCustomEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfilePredefinedGetResponseCustomEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfilePredefinedGetResponseCustomEntriesCustomPatternJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedGetResponseCustomEntriesCustomPattern]
type dlpProfilePredefinedGetResponseCustomEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseCustomEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseCustomEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseCustomEntriesCustomPatternValidation string

const (
	DLPProfilePredefinedGetResponseCustomEntriesCustomPatternValidationLuhn DLPProfilePredefinedGetResponseCustomEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfilePredefinedGetResponseCustomEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseCustomEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseCustomEntriesCustomType string

const (
	DLPProfilePredefinedGetResponseCustomEntriesCustomTypeCustom DLPProfilePredefinedGetResponseCustomEntriesCustomType = "custom"
)

func (r DLPProfilePredefinedGetResponseCustomEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseCustomEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseCustomEntriesPredefined struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Type      DLPProfilePredefinedGetResponseCustomEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponseCustomEntriesPredefinedJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseCustomEntriesPredefinedJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponseCustomEntriesPredefined]
type dlpProfilePredefinedGetResponseCustomEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseCustomEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseCustomEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseCustomEntriesPredefined) implementsZeroTrustDLPProfilePredefinedGetResponseCustomEntry() {
}

type DLPProfilePredefinedGetResponseCustomEntriesPredefinedType string

const (
	DLPProfilePredefinedGetResponseCustomEntriesPredefinedTypePredefined DLPProfilePredefinedGetResponseCustomEntriesPredefinedType = "predefined"
)

func (r DLPProfilePredefinedGetResponseCustomEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseCustomEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseCustomEntriesIntegration struct {
	ID        string                                                      `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                   `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                        `json:"enabled,required"`
	Name      string                                                      `json:"name,required"`
	Type      DLPProfilePredefinedGetResponseCustomEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                   `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                      `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponseCustomEntriesIntegrationJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseCustomEntriesIntegrationJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedGetResponseCustomEntriesIntegration]
type dlpProfilePredefinedGetResponseCustomEntriesIntegrationJSON struct {
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

func (r *DLPProfilePredefinedGetResponseCustomEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseCustomEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseCustomEntriesIntegration) implementsZeroTrustDLPProfilePredefinedGetResponseCustomEntry() {
}

type DLPProfilePredefinedGetResponseCustomEntriesIntegrationType string

const (
	DLPProfilePredefinedGetResponseCustomEntriesIntegrationTypeIntegration DLPProfilePredefinedGetResponseCustomEntriesIntegrationType = "integration"
)

func (r DLPProfilePredefinedGetResponseCustomEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseCustomEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseCustomEntriesExactData struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Secret    bool                                                      `json:"secret,required"`
	Type      DLPProfilePredefinedGetResponseCustomEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                 `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfilePredefinedGetResponseCustomEntriesExactDataJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseCustomEntriesExactDataJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponseCustomEntriesExactData]
type dlpProfilePredefinedGetResponseCustomEntriesExactDataJSON struct {
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

func (r *DLPProfilePredefinedGetResponseCustomEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseCustomEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseCustomEntriesExactData) implementsZeroTrustDLPProfilePredefinedGetResponseCustomEntry() {
}

type DLPProfilePredefinedGetResponseCustomEntriesExactDataType string

const (
	DLPProfilePredefinedGetResponseCustomEntriesExactDataTypeExactData DLPProfilePredefinedGetResponseCustomEntriesExactDataType = "exact_data"
)

func (r DLPProfilePredefinedGetResponseCustomEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseCustomEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseCustomEntriesWordList struct {
	ID        string                                                   `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                     `json:"enabled,required"`
	Name      string                                                   `json:"name,required"`
	Type      DLPProfilePredefinedGetResponseCustomEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                              `json:"word_list,required"`
	ProfileID string                                                   `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponseCustomEntriesWordListJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseCustomEntriesWordListJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponseCustomEntriesWordList]
type dlpProfilePredefinedGetResponseCustomEntriesWordListJSON struct {
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

func (r *DLPProfilePredefinedGetResponseCustomEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseCustomEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseCustomEntriesWordList) implementsZeroTrustDLPProfilePredefinedGetResponseCustomEntry() {
}

type DLPProfilePredefinedGetResponseCustomEntriesWordListType string

const (
	DLPProfilePredefinedGetResponseCustomEntriesWordListTypeWordList DLPProfilePredefinedGetResponseCustomEntriesWordListType = "word_list"
)

func (r DLPProfilePredefinedGetResponseCustomEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseCustomEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseCustomEntriesType string

const (
	DLPProfilePredefinedGetResponseCustomEntriesTypeCustom      DLPProfilePredefinedGetResponseCustomEntriesType = "custom"
	DLPProfilePredefinedGetResponseCustomEntriesTypePredefined  DLPProfilePredefinedGetResponseCustomEntriesType = "predefined"
	DLPProfilePredefinedGetResponseCustomEntriesTypeIntegration DLPProfilePredefinedGetResponseCustomEntriesType = "integration"
	DLPProfilePredefinedGetResponseCustomEntriesTypeExactData   DLPProfilePredefinedGetResponseCustomEntriesType = "exact_data"
	DLPProfilePredefinedGetResponseCustomEntriesTypeWordList    DLPProfilePredefinedGetResponseCustomEntriesType = "word_list"
)

func (r DLPProfilePredefinedGetResponseCustomEntriesType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseCustomEntriesTypeCustom, DLPProfilePredefinedGetResponseCustomEntriesTypePredefined, DLPProfilePredefinedGetResponseCustomEntriesTypeIntegration, DLPProfilePredefinedGetResponseCustomEntriesTypeExactData, DLPProfilePredefinedGetResponseCustomEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseCustomType string

const (
	DLPProfilePredefinedGetResponseCustomTypeCustom DLPProfilePredefinedGetResponseCustomType = "custom"
)

func (r DLPProfilePredefinedGetResponseCustomType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponsePredefined struct {
	// The id of the predefined profile (uuid)
	ID                string                                           `json:"id,required" format:"uuid"`
	AllowedMatchCount int64                                            `json:"allowed_match_count,required"`
	Entries           []DLPProfilePredefinedGetResponsePredefinedEntry `json:"entries,required"`
	// The name of the predefined profile
	Name string                                        `json:"name,required"`
	Type DLPProfilePredefinedGetResponsePredefinedType `json:"type,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness                              `json:"context_awareness"`
	OCREnabled       bool                                          `json:"ocr_enabled"`
	OpenAccess       bool                                          `json:"open_access"`
	JSON             dlpProfilePredefinedGetResponsePredefinedJSON `json:"-"`
}

// dlpProfilePredefinedGetResponsePredefinedJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedGetResponsePredefined]
type dlpProfilePredefinedGetResponsePredefinedJSON struct {
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

func (r *DLPProfilePredefinedGetResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponsePredefined) implementsZeroTrustDLPProfilePredefinedGetResponse() {
}

type DLPProfilePredefinedGetResponsePredefinedEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfilePredefinedGetResponsePredefinedEntriesCustomPattern].
	Pattern   interface{}                                          `json:"pattern,required"`
	ProfileID string                                               `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                            `json:"updated_at" format:"date-time"`
	Type      DLPProfilePredefinedGetResponsePredefinedEntriesType `json:"type,required"`
	Secret    bool                                                 `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                        `json:"word_list,required"`
	JSON     dlpProfilePredefinedGetResponsePredefinedEntryJSON `json:"-"`
	union    DLPProfilePredefinedGetResponsePredefinedEntriesUnion
}

// dlpProfilePredefinedGetResponsePredefinedEntryJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedGetResponsePredefinedEntry]
type dlpProfilePredefinedGetResponsePredefinedEntryJSON struct {
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

func (r dlpProfilePredefinedGetResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfilePredefinedGetResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfilePredefinedGetResponsePredefinedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfilePredefinedGetResponsePredefinedEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfilePredefinedGetResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfilePredefinedGetResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfilePredefinedGetResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfilePredefinedGetResponsePredefinedEntriesExactData],
// [zero_trust.DLPProfilePredefinedGetResponsePredefinedEntriesWordList].
func (r DLPProfilePredefinedGetResponsePredefinedEntry) AsUnion() DLPProfilePredefinedGetResponsePredefinedEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfilePredefinedGetResponsePredefinedEntriesCustom],
// [zero_trust.DLPProfilePredefinedGetResponsePredefinedEntriesPredefined],
// [zero_trust.DLPProfilePredefinedGetResponsePredefinedEntriesIntegration],
// [zero_trust.DLPProfilePredefinedGetResponsePredefinedEntriesExactData] or
// [zero_trust.DLPProfilePredefinedGetResponsePredefinedEntriesWordList].
type DLPProfilePredefinedGetResponsePredefinedEntriesUnion interface {
	implementsZeroTrustDLPProfilePredefinedGetResponsePredefinedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfilePredefinedGetResponsePredefinedEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponsePredefinedEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponsePredefinedEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponsePredefinedEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponsePredefinedEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponsePredefinedEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfilePredefinedGetResponsePredefinedEntriesCustom struct {
	ID        string                                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                          `json:"enabled,required"`
	Name      string                                                        `json:"name,required"`
	Pattern   DLPProfilePredefinedGetResponsePredefinedEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfilePredefinedGetResponsePredefinedEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                                     `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponsePredefinedEntriesCustomJSON    `json:"-"`
}

// dlpProfilePredefinedGetResponsePredefinedEntriesCustomJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponsePredefinedEntriesCustom]
type dlpProfilePredefinedGetResponsePredefinedEntriesCustomJSON struct {
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

func (r *DLPProfilePredefinedGetResponsePredefinedEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponsePredefinedEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponsePredefinedEntriesCustom) implementsZeroTrustDLPProfilePredefinedGetResponsePredefinedEntry() {
}

type DLPProfilePredefinedGetResponsePredefinedEntriesCustomPattern struct {
	Regex      string                                                                  `json:"regex,required"`
	Validation DLPProfilePredefinedGetResponsePredefinedEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfilePredefinedGetResponsePredefinedEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfilePredefinedGetResponsePredefinedEntriesCustomPatternJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedGetResponsePredefinedEntriesCustomPattern]
type dlpProfilePredefinedGetResponsePredefinedEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponsePredefinedEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponsePredefinedEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponsePredefinedEntriesCustomPatternValidation string

const (
	DLPProfilePredefinedGetResponsePredefinedEntriesCustomPatternValidationLuhn DLPProfilePredefinedGetResponsePredefinedEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfilePredefinedGetResponsePredefinedEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponsePredefinedEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponsePredefinedEntriesCustomType string

const (
	DLPProfilePredefinedGetResponsePredefinedEntriesCustomTypeCustom DLPProfilePredefinedGetResponsePredefinedEntriesCustomType = "custom"
)

func (r DLPProfilePredefinedGetResponsePredefinedEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponsePredefinedEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponsePredefinedEntriesPredefined struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	Enabled   bool                                                           `json:"enabled,required"`
	Name      string                                                         `json:"name,required"`
	Type      DLPProfilePredefinedGetResponsePredefinedEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponsePredefinedEntriesPredefinedJSON `json:"-"`
}

// dlpProfilePredefinedGetResponsePredefinedEntriesPredefinedJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedGetResponsePredefinedEntriesPredefined]
type dlpProfilePredefinedGetResponsePredefinedEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponsePredefinedEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponsePredefinedEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponsePredefinedEntriesPredefined) implementsZeroTrustDLPProfilePredefinedGetResponsePredefinedEntry() {
}

type DLPProfilePredefinedGetResponsePredefinedEntriesPredefinedType string

const (
	DLPProfilePredefinedGetResponsePredefinedEntriesPredefinedTypePredefined DLPProfilePredefinedGetResponsePredefinedEntriesPredefinedType = "predefined"
)

func (r DLPProfilePredefinedGetResponsePredefinedEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponsePredefinedEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponsePredefinedEntriesIntegration struct {
	ID        string                                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                            `json:"enabled,required"`
	Name      string                                                          `json:"name,required"`
	Type      DLPProfilePredefinedGetResponsePredefinedEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                       `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponsePredefinedEntriesIntegrationJSON `json:"-"`
}

// dlpProfilePredefinedGetResponsePredefinedEntriesIntegrationJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedGetResponsePredefinedEntriesIntegration]
type dlpProfilePredefinedGetResponsePredefinedEntriesIntegrationJSON struct {
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

func (r *DLPProfilePredefinedGetResponsePredefinedEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponsePredefinedEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponsePredefinedEntriesIntegration) implementsZeroTrustDLPProfilePredefinedGetResponsePredefinedEntry() {
}

type DLPProfilePredefinedGetResponsePredefinedEntriesIntegrationType string

const (
	DLPProfilePredefinedGetResponsePredefinedEntriesIntegrationTypeIntegration DLPProfilePredefinedGetResponsePredefinedEntriesIntegrationType = "integration"
)

func (r DLPProfilePredefinedGetResponsePredefinedEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponsePredefinedEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponsePredefinedEntriesExactData struct {
	ID        string                                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                          `json:"enabled,required"`
	Name      string                                                        `json:"name,required"`
	Secret    bool                                                          `json:"secret,required"`
	Type      DLPProfilePredefinedGetResponsePredefinedEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                     `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfilePredefinedGetResponsePredefinedEntriesExactDataJSON `json:"-"`
}

// dlpProfilePredefinedGetResponsePredefinedEntriesExactDataJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedGetResponsePredefinedEntriesExactData]
type dlpProfilePredefinedGetResponsePredefinedEntriesExactDataJSON struct {
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

func (r *DLPProfilePredefinedGetResponsePredefinedEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponsePredefinedEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponsePredefinedEntriesExactData) implementsZeroTrustDLPProfilePredefinedGetResponsePredefinedEntry() {
}

type DLPProfilePredefinedGetResponsePredefinedEntriesExactDataType string

const (
	DLPProfilePredefinedGetResponsePredefinedEntriesExactDataTypeExactData DLPProfilePredefinedGetResponsePredefinedEntriesExactDataType = "exact_data"
)

func (r DLPProfilePredefinedGetResponsePredefinedEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponsePredefinedEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponsePredefinedEntriesWordList struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Type      DLPProfilePredefinedGetResponsePredefinedEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                  `json:"word_list,required"`
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponsePredefinedEntriesWordListJSON `json:"-"`
}

// dlpProfilePredefinedGetResponsePredefinedEntriesWordListJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedGetResponsePredefinedEntriesWordList]
type dlpProfilePredefinedGetResponsePredefinedEntriesWordListJSON struct {
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

func (r *DLPProfilePredefinedGetResponsePredefinedEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponsePredefinedEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponsePredefinedEntriesWordList) implementsZeroTrustDLPProfilePredefinedGetResponsePredefinedEntry() {
}

type DLPProfilePredefinedGetResponsePredefinedEntriesWordListType string

const (
	DLPProfilePredefinedGetResponsePredefinedEntriesWordListTypeWordList DLPProfilePredefinedGetResponsePredefinedEntriesWordListType = "word_list"
)

func (r DLPProfilePredefinedGetResponsePredefinedEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponsePredefinedEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponsePredefinedEntriesType string

const (
	DLPProfilePredefinedGetResponsePredefinedEntriesTypeCustom      DLPProfilePredefinedGetResponsePredefinedEntriesType = "custom"
	DLPProfilePredefinedGetResponsePredefinedEntriesTypePredefined  DLPProfilePredefinedGetResponsePredefinedEntriesType = "predefined"
	DLPProfilePredefinedGetResponsePredefinedEntriesTypeIntegration DLPProfilePredefinedGetResponsePredefinedEntriesType = "integration"
	DLPProfilePredefinedGetResponsePredefinedEntriesTypeExactData   DLPProfilePredefinedGetResponsePredefinedEntriesType = "exact_data"
	DLPProfilePredefinedGetResponsePredefinedEntriesTypeWordList    DLPProfilePredefinedGetResponsePredefinedEntriesType = "word_list"
)

func (r DLPProfilePredefinedGetResponsePredefinedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponsePredefinedEntriesTypeCustom, DLPProfilePredefinedGetResponsePredefinedEntriesTypePredefined, DLPProfilePredefinedGetResponsePredefinedEntriesTypeIntegration, DLPProfilePredefinedGetResponsePredefinedEntriesTypeExactData, DLPProfilePredefinedGetResponsePredefinedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponsePredefinedType string

const (
	DLPProfilePredefinedGetResponsePredefinedTypePredefined DLPProfilePredefinedGetResponsePredefinedType = "predefined"
)

func (r DLPProfilePredefinedGetResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseIntegration struct {
	ID        string                                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                         `json:"created_at,required" format:"date-time"`
	Entries   []DLPProfilePredefinedGetResponseIntegrationEntry `json:"entries,required"`
	Name      string                                            `json:"name,required"`
	Type      DLPProfilePredefinedGetResponseIntegrationType    `json:"type,required"`
	UpdatedAt time.Time                                         `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                                         `json:"description,nullable"`
	JSON        dlpProfilePredefinedGetResponseIntegrationJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseIntegrationJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedGetResponseIntegration]
type dlpProfilePredefinedGetResponseIntegrationJSON struct {
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

func (r *DLPProfilePredefinedGetResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseIntegration) implementsZeroTrustDLPProfilePredefinedGetResponse() {
}

type DLPProfilePredefinedGetResponseIntegrationEntry struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Enabled   bool      `json:"enabled,required"`
	ID        string    `json:"id,required" format:"uuid"`
	Name      string    `json:"name,required"`
	// This field can have the runtime type of
	// [DLPProfilePredefinedGetResponseIntegrationEntriesCustomPattern].
	Pattern   interface{}                                           `json:"pattern,required"`
	ProfileID string                                                `json:"profile_id,nullable" format:"uuid"`
	UpdatedAt time.Time                                             `json:"updated_at" format:"date-time"`
	Type      DLPProfilePredefinedGetResponseIntegrationEntriesType `json:"type,required"`
	Secret    bool                                                  `json:"secret"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                         `json:"word_list,required"`
	JSON     dlpProfilePredefinedGetResponseIntegrationEntryJSON `json:"-"`
	union    DLPProfilePredefinedGetResponseIntegrationEntriesUnion
}

// dlpProfilePredefinedGetResponseIntegrationEntryJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedGetResponseIntegrationEntry]
type dlpProfilePredefinedGetResponseIntegrationEntryJSON struct {
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

func (r dlpProfilePredefinedGetResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfilePredefinedGetResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfilePredefinedGetResponseIntegrationEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfilePredefinedGetResponseIntegrationEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.DLPProfilePredefinedGetResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfilePredefinedGetResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfilePredefinedGetResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfilePredefinedGetResponseIntegrationEntriesExactData],
// [zero_trust.DLPProfilePredefinedGetResponseIntegrationEntriesWordList].
func (r DLPProfilePredefinedGetResponseIntegrationEntry) AsUnion() DLPProfilePredefinedGetResponseIntegrationEntriesUnion {
	return r.union
}

// Union satisfied by
// [zero_trust.DLPProfilePredefinedGetResponseIntegrationEntriesCustom],
// [zero_trust.DLPProfilePredefinedGetResponseIntegrationEntriesPredefined],
// [zero_trust.DLPProfilePredefinedGetResponseIntegrationEntriesIntegration],
// [zero_trust.DLPProfilePredefinedGetResponseIntegrationEntriesExactData] or
// [zero_trust.DLPProfilePredefinedGetResponseIntegrationEntriesWordList].
type DLPProfilePredefinedGetResponseIntegrationEntriesUnion interface {
	implementsZeroTrustDLPProfilePredefinedGetResponseIntegrationEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfilePredefinedGetResponseIntegrationEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponseIntegrationEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponseIntegrationEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponseIntegrationEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponseIntegrationEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfilePredefinedGetResponseIntegrationEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPProfilePredefinedGetResponseIntegrationEntriesCustom struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                           `json:"enabled,required"`
	Name      string                                                         `json:"name,required"`
	Pattern   DLPProfilePredefinedGetResponseIntegrationEntriesCustomPattern `json:"pattern,required"`
	Type      DLPProfilePredefinedGetResponseIntegrationEntriesCustomType    `json:"type,required"`
	UpdatedAt time.Time                                                      `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponseIntegrationEntriesCustomJSON    `json:"-"`
}

// dlpProfilePredefinedGetResponseIntegrationEntriesCustomJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedGetResponseIntegrationEntriesCustom]
type dlpProfilePredefinedGetResponseIntegrationEntriesCustomJSON struct {
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

func (r *DLPProfilePredefinedGetResponseIntegrationEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseIntegrationEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseIntegrationEntriesCustom) implementsZeroTrustDLPProfilePredefinedGetResponseIntegrationEntry() {
}

type DLPProfilePredefinedGetResponseIntegrationEntriesCustomPattern struct {
	Regex      string                                                                   `json:"regex,required"`
	Validation DLPProfilePredefinedGetResponseIntegrationEntriesCustomPatternValidation `json:"validation"`
	JSON       dlpProfilePredefinedGetResponseIntegrationEntriesCustomPatternJSON       `json:"-"`
}

// dlpProfilePredefinedGetResponseIntegrationEntriesCustomPatternJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedGetResponseIntegrationEntriesCustomPattern]
type dlpProfilePredefinedGetResponseIntegrationEntriesCustomPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseIntegrationEntriesCustomPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseIntegrationEntriesCustomPatternJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseIntegrationEntriesCustomPatternValidation string

const (
	DLPProfilePredefinedGetResponseIntegrationEntriesCustomPatternValidationLuhn DLPProfilePredefinedGetResponseIntegrationEntriesCustomPatternValidation = "luhn"
)

func (r DLPProfilePredefinedGetResponseIntegrationEntriesCustomPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseIntegrationEntriesCustomPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseIntegrationEntriesCustomType string

const (
	DLPProfilePredefinedGetResponseIntegrationEntriesCustomTypeCustom DLPProfilePredefinedGetResponseIntegrationEntriesCustomType = "custom"
)

func (r DLPProfilePredefinedGetResponseIntegrationEntriesCustomType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseIntegrationEntriesCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseIntegrationEntriesPredefined struct {
	ID        string                                                          `json:"id,required" format:"uuid"`
	Enabled   bool                                                            `json:"enabled,required"`
	Name      string                                                          `json:"name,required"`
	Type      DLPProfilePredefinedGetResponseIntegrationEntriesPredefinedType `json:"type,required"`
	ProfileID string                                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponseIntegrationEntriesPredefinedJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseIntegrationEntriesPredefinedJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedGetResponseIntegrationEntriesPredefined]
type dlpProfilePredefinedGetResponseIntegrationEntriesPredefinedJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseIntegrationEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseIntegrationEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseIntegrationEntriesPredefined) implementsZeroTrustDLPProfilePredefinedGetResponseIntegrationEntry() {
}

type DLPProfilePredefinedGetResponseIntegrationEntriesPredefinedType string

const (
	DLPProfilePredefinedGetResponseIntegrationEntriesPredefinedTypePredefined DLPProfilePredefinedGetResponseIntegrationEntriesPredefinedType = "predefined"
)

func (r DLPProfilePredefinedGetResponseIntegrationEntriesPredefinedType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseIntegrationEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseIntegrationEntriesIntegration struct {
	ID        string                                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                             `json:"enabled,required"`
	Name      string                                                           `json:"name,required"`
	Type      DLPProfilePredefinedGetResponseIntegrationEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                                        `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponseIntegrationEntriesIntegrationJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseIntegrationEntriesIntegrationJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedGetResponseIntegrationEntriesIntegration]
type dlpProfilePredefinedGetResponseIntegrationEntriesIntegrationJSON struct {
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

func (r *DLPProfilePredefinedGetResponseIntegrationEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseIntegrationEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseIntegrationEntriesIntegration) implementsZeroTrustDLPProfilePredefinedGetResponseIntegrationEntry() {
}

type DLPProfilePredefinedGetResponseIntegrationEntriesIntegrationType string

const (
	DLPProfilePredefinedGetResponseIntegrationEntriesIntegrationTypeIntegration DLPProfilePredefinedGetResponseIntegrationEntriesIntegrationType = "integration"
)

func (r DLPProfilePredefinedGetResponseIntegrationEntriesIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseIntegrationEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseIntegrationEntriesExactData struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                           `json:"enabled,required"`
	Name      string                                                         `json:"name,required"`
	Secret    bool                                                           `json:"secret,required"`
	Type      DLPProfilePredefinedGetResponseIntegrationEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                                                      `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfilePredefinedGetResponseIntegrationEntriesExactDataJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseIntegrationEntriesExactDataJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedGetResponseIntegrationEntriesExactData]
type dlpProfilePredefinedGetResponseIntegrationEntriesExactDataJSON struct {
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

func (r *DLPProfilePredefinedGetResponseIntegrationEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseIntegrationEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseIntegrationEntriesExactData) implementsZeroTrustDLPProfilePredefinedGetResponseIntegrationEntry() {
}

type DLPProfilePredefinedGetResponseIntegrationEntriesExactDataType string

const (
	DLPProfilePredefinedGetResponseIntegrationEntriesExactDataTypeExactData DLPProfilePredefinedGetResponseIntegrationEntriesExactDataType = "exact_data"
)

func (r DLPProfilePredefinedGetResponseIntegrationEntriesExactDataType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseIntegrationEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseIntegrationEntriesWordList struct {
	ID        string                                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                          `json:"enabled,required"`
	Name      string                                                        `json:"name,required"`
	Type      DLPProfilePredefinedGetResponseIntegrationEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                                                     `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                   `json:"word_list,required"`
	ProfileID string                                                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponseIntegrationEntriesWordListJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseIntegrationEntriesWordListJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedGetResponseIntegrationEntriesWordList]
type dlpProfilePredefinedGetResponseIntegrationEntriesWordListJSON struct {
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

func (r *DLPProfilePredefinedGetResponseIntegrationEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseIntegrationEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseIntegrationEntriesWordList) implementsZeroTrustDLPProfilePredefinedGetResponseIntegrationEntry() {
}

type DLPProfilePredefinedGetResponseIntegrationEntriesWordListType string

const (
	DLPProfilePredefinedGetResponseIntegrationEntriesWordListTypeWordList DLPProfilePredefinedGetResponseIntegrationEntriesWordListType = "word_list"
)

func (r DLPProfilePredefinedGetResponseIntegrationEntriesWordListType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseIntegrationEntriesWordListTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseIntegrationEntriesType string

const (
	DLPProfilePredefinedGetResponseIntegrationEntriesTypeCustom      DLPProfilePredefinedGetResponseIntegrationEntriesType = "custom"
	DLPProfilePredefinedGetResponseIntegrationEntriesTypePredefined  DLPProfilePredefinedGetResponseIntegrationEntriesType = "predefined"
	DLPProfilePredefinedGetResponseIntegrationEntriesTypeIntegration DLPProfilePredefinedGetResponseIntegrationEntriesType = "integration"
	DLPProfilePredefinedGetResponseIntegrationEntriesTypeExactData   DLPProfilePredefinedGetResponseIntegrationEntriesType = "exact_data"
	DLPProfilePredefinedGetResponseIntegrationEntriesTypeWordList    DLPProfilePredefinedGetResponseIntegrationEntriesType = "word_list"
)

func (r DLPProfilePredefinedGetResponseIntegrationEntriesType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseIntegrationEntriesTypeCustom, DLPProfilePredefinedGetResponseIntegrationEntriesTypePredefined, DLPProfilePredefinedGetResponseIntegrationEntriesTypeIntegration, DLPProfilePredefinedGetResponseIntegrationEntriesTypeExactData, DLPProfilePredefinedGetResponseIntegrationEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseIntegrationType string

const (
	DLPProfilePredefinedGetResponseIntegrationTypeIntegration DLPProfilePredefinedGetResponseIntegrationType = "integration"
)

func (r DLPProfilePredefinedGetResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseType string

const (
	DLPProfilePredefinedGetResponseTypeCustom      DLPProfilePredefinedGetResponseType = "custom"
	DLPProfilePredefinedGetResponseTypePredefined  DLPProfilePredefinedGetResponseType = "predefined"
	DLPProfilePredefinedGetResponseTypeIntegration DLPProfilePredefinedGetResponseType = "integration"
)

func (r DLPProfilePredefinedGetResponseType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseTypeCustom, DLPProfilePredefinedGetResponseTypePredefined, DLPProfilePredefinedGetResponseTypeIntegration:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfilePredefinedUpdateResponseEnvelope struct {
	Errors     []shared.ResponseInfo                                `json:"errors,required"`
	Messages   []shared.ResponseInfo                                `json:"messages,required"`
	Success    bool                                                 `json:"success,required"`
	Result     DLPProfilePredefinedUpdateResponse                   `json:"result"`
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
	Result     DLPProfilePredefinedGetResponse                   `json:"result"`
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
