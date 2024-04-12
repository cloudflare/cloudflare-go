// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// DLPProfileService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDLPProfileService] method instead.
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
func (r *DLPProfileService) List(ctx context.Context, query DLPProfileListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Profile], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *DLPProfileService) ListAutoPaging(ctx context.Context, query DLPProfileListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Profile] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Fetches a DLP profile by ID. Supports both predefined and custom profiles
func (r *DLPProfileService) Get(ctx context.Context, profileID string, query DLPProfileGetParams, opts ...option.RequestOption) (res *DLPProfileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfileGetResponseEnvelope
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

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ContextAwarenessParam struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip param.Field[SkipConfigurationParam] `json:"skip,required"`
}

func (r ContextAwarenessParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Profile struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness"`
	Entries          interface{}      `json:"entries,required"`
	// The ID for this profile
	ID string `json:"id"`
	// The name of the profile.
	Name string `json:"name"`
	// If true, scan images via OCR to determine if any text present matches filters.
	OCREnabled bool `json:"ocr_enabled"`
	// The type of the profile.
	Type      ProfileType `json:"type"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string      `json:"description"`
	UpdatedAt   time.Time   `json:"updated_at" format:"date-time"`
	JSON        profileJSON `json:"-"`
	union       ProfileUnion
}

// profileJSON contains the JSON metadata for the struct [Profile]
type profileJSON struct {
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	Entries           apijson.Field
	ID                apijson.Field
	Name              apijson.Field
	OCREnabled        apijson.Field
	Type              apijson.Field
	CreatedAt         apijson.Field
	Description       apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r profileJSON) RawJSON() string {
	return r.raw
}

func (r *Profile) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r Profile) AsUnion() ProfileUnion {
	return r.union
}

// Union satisfied by [zero_trust.PredefinedProfile], [zero_trust.CustomProfile] or
// [zero_trust.ProfileDLPIntegrationProfile].
type ProfileUnion interface {
	implementsZeroTrustProfile()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PredefinedProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileDLPIntegrationProfile{}),
		},
	)
}

type ProfileDLPIntegrationProfile struct {
	// The ID for this profile
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []ProfileDLPIntegrationProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      ProfileDLPIntegrationProfileType `json:"type"`
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	JSON      profileDLPIntegrationProfileJSON `json:"-"`
}

// profileDLPIntegrationProfileJSON contains the JSON metadata for the struct
// [ProfileDLPIntegrationProfile]
type profileDLPIntegrationProfileJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Entries     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileDLPIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileDLPIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r ProfileDLPIntegrationProfile) implementsZeroTrustProfile() {}

// An entry derived from an integration
type ProfileDLPIntegrationProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                           `json:"profile_id"`
	UpdatedAt time.Time                             `json:"updated_at" format:"date-time"`
	JSON      profileDLPIntegrationProfileEntryJSON `json:"-"`
}

// profileDLPIntegrationProfileEntryJSON contains the JSON metadata for the struct
// [ProfileDLPIntegrationProfileEntry]
type profileDLPIntegrationProfileEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileDLPIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileDLPIntegrationProfileEntryJSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type ProfileDLPIntegrationProfileType string

const (
	ProfileDLPIntegrationProfileTypeIntegration ProfileDLPIntegrationProfileType = "integration"
)

func (r ProfileDLPIntegrationProfileType) IsKnown() bool {
	switch r {
	case ProfileDLPIntegrationProfileTypeIntegration:
		return true
	}
	return false
}

// The type of the profile.
type ProfileType string

const (
	ProfileTypePredefined  ProfileType = "predefined"
	ProfileTypeCustom      ProfileType = "custom"
	ProfileTypeIntegration ProfileType = "integration"
)

func (r ProfileType) IsKnown() bool {
	switch r {
	case ProfileTypePredefined, ProfileTypeCustom, ProfileTypeIntegration:
		return true
	}
	return false
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

// Content types to exclude from context analysis and return all matches.
type SkipConfigurationParam struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files param.Field[bool] `json:"files,required"`
}

func (r SkipConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileGetResponse struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness"`
	Entries          interface{}      `json:"entries,required"`
	// The ID for this profile
	ID string `json:"id"`
	// The name of the profile.
	Name string `json:"name"`
	// If true, scan images via OCR to determine if any text present matches filters.
	OCREnabled bool `json:"ocr_enabled"`
	// The type of the profile.
	Type      DLPProfileGetResponseType `json:"type"`
	CreatedAt time.Time                 `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string                    `json:"description"`
	UpdatedAt   time.Time                 `json:"updated_at" format:"date-time"`
	JSON        dlpProfileGetResponseJSON `json:"-"`
	union       DLPProfileGetResponseUnion
}

// dlpProfileGetResponseJSON contains the JSON metadata for the struct
// [DLPProfileGetResponse]
type dlpProfileGetResponseJSON struct {
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	Entries           apijson.Field
	ID                apijson.Field
	Name              apijson.Field
	OCREnabled        apijson.Field
	Type              apijson.Field
	CreatedAt         apijson.Field
	Description       apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r dlpProfileGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileGetResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r DLPProfileGetResponse) AsUnion() DLPProfileGetResponseUnion {
	return r.union
}

// Union satisfied by [zero_trust.PredefinedProfile], [zero_trust.CustomProfile] or
// [zero_trust.DLPProfileGetResponseDLPIntegrationProfile].
type DLPProfileGetResponseUnion interface {
	implementsZeroTrustDLPProfileGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PredefinedProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseDLPIntegrationProfile{}),
		},
	)
}

type DLPProfileGetResponseDLPIntegrationProfile struct {
	// The ID for this profile
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []DLPProfileGetResponseDLPIntegrationProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      DLPProfileGetResponseDLPIntegrationProfileType `json:"type"`
	UpdatedAt time.Time                                      `json:"updated_at" format:"date-time"`
	JSON      dlpProfileGetResponseDLPIntegrationProfileJSON `json:"-"`
}

// dlpProfileGetResponseDLPIntegrationProfileJSON contains the JSON metadata for
// the struct [DLPProfileGetResponseDLPIntegrationProfile]
type dlpProfileGetResponseDLPIntegrationProfileJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Entries     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseDLPIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseDLPIntegrationProfile) implementsZeroTrustDLPProfileGetResponse() {}

// An entry derived from an integration
type DLPProfileGetResponseDLPIntegrationProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                         `json:"profile_id"`
	UpdatedAt time.Time                                           `json:"updated_at" format:"date-time"`
	JSON      dlpProfileGetResponseDLPIntegrationProfileEntryJSON `json:"-"`
}

// dlpProfileGetResponseDLPIntegrationProfileEntryJSON contains the JSON metadata
// for the struct [DLPProfileGetResponseDLPIntegrationProfileEntry]
type dlpProfileGetResponseDLPIntegrationProfileEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseDLPIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPIntegrationProfileEntryJSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type DLPProfileGetResponseDLPIntegrationProfileType string

const (
	DLPProfileGetResponseDLPIntegrationProfileTypeIntegration DLPProfileGetResponseDLPIntegrationProfileType = "integration"
)

func (r DLPProfileGetResponseDLPIntegrationProfileType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseDLPIntegrationProfileTypeIntegration:
		return true
	}
	return false
}

// The type of the profile.
type DLPProfileGetResponseType string

const (
	DLPProfileGetResponseTypePredefined  DLPProfileGetResponseType = "predefined"
	DLPProfileGetResponseTypeCustom      DLPProfileGetResponseType = "custom"
	DLPProfileGetResponseTypeIntegration DLPProfileGetResponseType = "integration"
)

func (r DLPProfileGetResponseType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseTypePredefined, DLPProfileGetResponseTypeCustom, DLPProfileGetResponseTypeIntegration:
		return true
	}
	return false
}

type DLPProfileListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DLPProfileGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success DLPProfileGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpProfileGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPProfileGetResponseEnvelope]
type dlpProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPProfileGetResponseEnvelopeSuccess bool

const (
	DLPProfileGetResponseEnvelopeSuccessTrue DLPProfileGetResponseEnvelopeSuccess = true
)

func (r DLPProfileGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
