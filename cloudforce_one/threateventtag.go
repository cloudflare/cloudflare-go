// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ThreatEventTagService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventTagService] method instead.
type ThreatEventTagService struct {
	Options    []option.RequestOption
	Categories *ThreatEventTagCategoryService
	Indicators *ThreatEventTagIndicatorService
}

// NewThreatEventTagService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThreatEventTagService(opts ...option.RequestOption) (r *ThreatEventTagService) {
	r = &ThreatEventTagService{}
	r.Options = opts
	r.Categories = NewThreatEventTagCategoryService(opts...)
	r.Indicators = NewThreatEventTagIndicatorService(opts...)
	return
}

// Creates a new tag to be used accross threat events.
func (r *ThreatEventTagService) New(ctx context.Context, params ThreatEventTagNewParams, opts ...option.RequestOption) (res *ThreatEventTagNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/create", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns all Source-of-Truth tags for an account. Supports legacy free-text
// `search` on tag value and `categoryUuid` exact match, plus a structured
// `filters` JSON array for filtering by metadata fields (originCountryISO,
// actorCategory, motive, priority, etc.). Country values may be passed as alpha-2,
// alpha-3, name, or common alias.
func (r *ThreatEventTagService) List(ctx context.Context, params ThreatEventTagListParams, opts ...option.RequestOption) (res *ThreatEventTagListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Deletes a Source-of-Truth tag by UUID.
func (r *ThreatEventTagService) Delete(ctx context.Context, tagUUID string, body ThreatEventTagDeleteParams, opts ...option.RequestOption) (res *ThreatEventTagDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if tagUUID == "" {
		err = errors.New("missing required tag_uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/%s", body.AccountID, tagUUID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Updates a Source-of-Truth tag by UUID.
func (r *ThreatEventTagService) Edit(ctx context.Context, tagUUID string, params ThreatEventTagEditParams, opts ...option.RequestOption) (res *ThreatEventTagEditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if tagUUID == "" {
		err = errors.New("missing required tag_uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/%s", params.AccountID, tagUUID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

type ThreatEventTagNewResponse struct {
	UUID                    string                                           `json:"uuid" api:"required"`
	Value                   string                                           `json:"value" api:"required"`
	ActiveDuration          string                                           `json:"activeDuration"`
	ActiveDurationAnnotated ThreatEventTagNewResponseActiveDurationAnnotated `json:"activeDuration_annotated" api:"nullable"`
	ActorCategory           string                                           `json:"actorCategory"`
	ActorCategoryAnnotated  ThreatEventTagNewResponseActorCategoryAnnotated  `json:"actorCategory_annotated" api:"nullable"`
	// Structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: stripped from
	// responses to non-CFONE accounts.
	Aliases                          []ThreatEventTagNewResponseAlias                          `json:"aliases"`
	AliasGroupNames                  []string                                                  `json:"aliasGroupNames"`
	AliasGroupNamesInternal          []string                                                  `json:"aliasGroupNamesInternal"`
	AttributionOrganization          string                                                    `json:"attributionOrganization"`
	AttributionOrganizationAnnotated ThreatEventTagNewResponseAttributionOrganizationAnnotated `json:"attributionOrganization_annotated" api:"nullable"`
	CategoryName                     string                                                    `json:"categoryName"`
	CategoryUUID                     string                                                    `json:"categoryUuid"`
	// Overall tag confidence (1-10).
	Confidence             int64    `json:"confidence" api:"nullable"`
	CreatedAt              string   `json:"createdAt"`
	DateOfDiscovery        string   `json:"dateOfDiscovery"`
	Description            string   `json:"description"`
	ExternalReferenceLinks []string `json:"externalReferenceLinks"`
	// Structured external references ({ url, description }). Public: returned to all
	// accounts.
	ExternalReferences          []ThreatEventTagNewResponseExternalReference           `json:"externalReferences"`
	ExternalReferencesAnnotated []ThreatEventTagNewResponseExternalReferencesAnnotated `json:"externalReferences_annotated" api:"nullable"`
	// Internal structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: never
	// returned to non-CFONE accounts.
	InternalAliases     []ThreatEventTagNewResponseInternalAlias     `json:"internalAliases"`
	InternalDescription string                                       `json:"internalDescription"`
	LastSeen            string                                       `json:"lastSeen"`
	Motive              string                                       `json:"motive"`
	MotiveAnnotated     ThreatEventTagNewResponseMotiveAnnotated     `json:"motive_annotated" api:"nullable"`
	OpsecLevel          string                                       `json:"opsecLevel"`
	OpsecLevelAnnotated ThreatEventTagNewResponseOpsecLevelAnnotated `json:"opsecLevel_annotated" api:"nullable"`
	// ISO country code (alpha-2 or alpha-3). Normalized to uppercase on read. Null
	// when stored value is blank/whitespace.
	OriginCountryISO          string                                             `json:"originCountryISO" api:"nullable"`
	OriginCountryISOAnnotated ThreatEventTagNewResponseOriginCountryISOAnnotated `json:"originCountryISO_annotated" api:"nullable"`
	Priority                  float64                                            `json:"priority"`
	PriorityAnnotated         ThreatEventTagNewResponsePriorityAnnotated         `json:"priority_annotated" api:"nullable"`
	// Parsed custom field values. Null when the tag has no custom fields.
	Properties                   map[string]interface{}                                `json:"properties" api:"nullable"`
	SophisticationLevel          string                                                `json:"sophisticationLevel"`
	SophisticationLevelAnnotated ThreatEventTagNewResponseSophisticationLevelAnnotated `json:"sophisticationLevel_annotated" api:"nullable"`
	// Tag-level TLP handling marking.
	TLP       ThreatEventTagNewResponseTLP  `json:"tlp" api:"nullable"`
	UpdatedAt string                        `json:"updatedAt"`
	Version   float64                       `json:"version"`
	JSON      threatEventTagNewResponseJSON `json:"-"`
}

// threatEventTagNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagNewResponse]
type threatEventTagNewResponseJSON struct {
	UUID                             apijson.Field
	Value                            apijson.Field
	ActiveDuration                   apijson.Field
	ActiveDurationAnnotated          apijson.Field
	ActorCategory                    apijson.Field
	ActorCategoryAnnotated           apijson.Field
	Aliases                          apijson.Field
	AliasGroupNames                  apijson.Field
	AliasGroupNamesInternal          apijson.Field
	AttributionOrganization          apijson.Field
	AttributionOrganizationAnnotated apijson.Field
	CategoryName                     apijson.Field
	CategoryUUID                     apijson.Field
	Confidence                       apijson.Field
	CreatedAt                        apijson.Field
	DateOfDiscovery                  apijson.Field
	Description                      apijson.Field
	ExternalReferenceLinks           apijson.Field
	ExternalReferences               apijson.Field
	ExternalReferencesAnnotated      apijson.Field
	InternalAliases                  apijson.Field
	InternalDescription              apijson.Field
	LastSeen                         apijson.Field
	Motive                           apijson.Field
	MotiveAnnotated                  apijson.Field
	OpsecLevel                       apijson.Field
	OpsecLevelAnnotated              apijson.Field
	OriginCountryISO                 apijson.Field
	OriginCountryISOAnnotated        apijson.Field
	Priority                         apijson.Field
	PriorityAnnotated                apijson.Field
	Properties                       apijson.Field
	SophisticationLevel              apijson.Field
	SophisticationLevelAnnotated     apijson.Field
	TLP                              apijson.Field
	UpdatedAt                        apijson.Field
	Version                          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *ThreatEventTagNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewResponseActiveDurationAnnotated struct {
	Value string                                               `json:"value" api:"required"`
	TLP   ThreatEventTagNewResponseActiveDurationAnnotatedTLP  `json:"tlp"`
	JSON  threatEventTagNewResponseActiveDurationAnnotatedJSON `json:"-"`
}

// threatEventTagNewResponseActiveDurationAnnotatedJSON contains the JSON metadata
// for the struct [ThreatEventTagNewResponseActiveDurationAnnotated]
type threatEventTagNewResponseActiveDurationAnnotatedJSON struct {
	Value       apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponseActiveDurationAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseActiveDurationAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewResponseActiveDurationAnnotatedTLP string

const (
	ThreatEventTagNewResponseActiveDurationAnnotatedTLPRed         ThreatEventTagNewResponseActiveDurationAnnotatedTLP = "red"
	ThreatEventTagNewResponseActiveDurationAnnotatedTLPAmber       ThreatEventTagNewResponseActiveDurationAnnotatedTLP = "amber"
	ThreatEventTagNewResponseActiveDurationAnnotatedTLPAmberStrict ThreatEventTagNewResponseActiveDurationAnnotatedTLP = "amber+strict"
	ThreatEventTagNewResponseActiveDurationAnnotatedTLPGreen       ThreatEventTagNewResponseActiveDurationAnnotatedTLP = "green"
	ThreatEventTagNewResponseActiveDurationAnnotatedTLPClear       ThreatEventTagNewResponseActiveDurationAnnotatedTLP = "clear"
	ThreatEventTagNewResponseActiveDurationAnnotatedTLPPurple      ThreatEventTagNewResponseActiveDurationAnnotatedTLP = "purple"
)

func (r ThreatEventTagNewResponseActiveDurationAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseActiveDurationAnnotatedTLPRed, ThreatEventTagNewResponseActiveDurationAnnotatedTLPAmber, ThreatEventTagNewResponseActiveDurationAnnotatedTLPAmberStrict, ThreatEventTagNewResponseActiveDurationAnnotatedTLPGreen, ThreatEventTagNewResponseActiveDurationAnnotatedTLPClear, ThreatEventTagNewResponseActiveDurationAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagNewResponseActorCategoryAnnotated struct {
	Value      string                                              `json:"value" api:"required"`
	Confidence float64                                             `json:"confidence"`
	TLP        ThreatEventTagNewResponseActorCategoryAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagNewResponseActorCategoryAnnotatedJSON `json:"-"`
}

// threatEventTagNewResponseActorCategoryAnnotatedJSON contains the JSON metadata
// for the struct [ThreatEventTagNewResponseActorCategoryAnnotated]
type threatEventTagNewResponseActorCategoryAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponseActorCategoryAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseActorCategoryAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewResponseActorCategoryAnnotatedTLP string

const (
	ThreatEventTagNewResponseActorCategoryAnnotatedTLPRed         ThreatEventTagNewResponseActorCategoryAnnotatedTLP = "red"
	ThreatEventTagNewResponseActorCategoryAnnotatedTLPAmber       ThreatEventTagNewResponseActorCategoryAnnotatedTLP = "amber"
	ThreatEventTagNewResponseActorCategoryAnnotatedTLPAmberStrict ThreatEventTagNewResponseActorCategoryAnnotatedTLP = "amber+strict"
	ThreatEventTagNewResponseActorCategoryAnnotatedTLPGreen       ThreatEventTagNewResponseActorCategoryAnnotatedTLP = "green"
	ThreatEventTagNewResponseActorCategoryAnnotatedTLPClear       ThreatEventTagNewResponseActorCategoryAnnotatedTLP = "clear"
	ThreatEventTagNewResponseActorCategoryAnnotatedTLPPurple      ThreatEventTagNewResponseActorCategoryAnnotatedTLP = "purple"
)

func (r ThreatEventTagNewResponseActorCategoryAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseActorCategoryAnnotatedTLPRed, ThreatEventTagNewResponseActorCategoryAnnotatedTLPAmber, ThreatEventTagNewResponseActorCategoryAnnotatedTLPAmberStrict, ThreatEventTagNewResponseActorCategoryAnnotatedTLPGreen, ThreatEventTagNewResponseActorCategoryAnnotatedTLPClear, ThreatEventTagNewResponseActorCategoryAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagNewResponseAlias struct {
	Value      string                              `json:"value" api:"required"`
	Confidence int64                               `json:"confidence" api:"nullable"`
	TLP        ThreatEventTagNewResponseAliasesTLP `json:"tlp" api:"nullable"`
	JSON       threatEventTagNewResponseAliasJSON  `json:"-"`
}

// threatEventTagNewResponseAliasJSON contains the JSON metadata for the struct
// [ThreatEventTagNewResponseAlias]
type threatEventTagNewResponseAliasJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponseAlias) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseAliasJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewResponseAliasesTLP string

const (
	ThreatEventTagNewResponseAliasesTLPRed         ThreatEventTagNewResponseAliasesTLP = "red"
	ThreatEventTagNewResponseAliasesTLPAmber       ThreatEventTagNewResponseAliasesTLP = "amber"
	ThreatEventTagNewResponseAliasesTLPAmberStrict ThreatEventTagNewResponseAliasesTLP = "amber+strict"
	ThreatEventTagNewResponseAliasesTLPGreen       ThreatEventTagNewResponseAliasesTLP = "green"
	ThreatEventTagNewResponseAliasesTLPClear       ThreatEventTagNewResponseAliasesTLP = "clear"
	ThreatEventTagNewResponseAliasesTLPPurple      ThreatEventTagNewResponseAliasesTLP = "purple"
)

func (r ThreatEventTagNewResponseAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseAliasesTLPRed, ThreatEventTagNewResponseAliasesTLPAmber, ThreatEventTagNewResponseAliasesTLPAmberStrict, ThreatEventTagNewResponseAliasesTLPGreen, ThreatEventTagNewResponseAliasesTLPClear, ThreatEventTagNewResponseAliasesTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagNewResponseAttributionOrganizationAnnotated struct {
	Value      string                                                        `json:"value" api:"required"`
	Confidence float64                                                       `json:"confidence"`
	TLP        ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagNewResponseAttributionOrganizationAnnotatedJSON `json:"-"`
}

// threatEventTagNewResponseAttributionOrganizationAnnotatedJSON contains the JSON
// metadata for the struct
// [ThreatEventTagNewResponseAttributionOrganizationAnnotated]
type threatEventTagNewResponseAttributionOrganizationAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponseAttributionOrganizationAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseAttributionOrganizationAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLP string

const (
	ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLPRed         ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLP = "red"
	ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLPAmber       ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLP = "amber"
	ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLPAmberStrict ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLP = "amber+strict"
	ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLPGreen       ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLP = "green"
	ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLPClear       ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLP = "clear"
	ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLPPurple      ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLP = "purple"
)

func (r ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLPRed, ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLPAmber, ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLPAmberStrict, ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLPGreen, ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLPClear, ThreatEventTagNewResponseAttributionOrganizationAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagNewResponseExternalReference struct {
	URL         string                                         `json:"url" api:"required"`
	Description string                                         `json:"description" api:"nullable"`
	JSON        threatEventTagNewResponseExternalReferenceJSON `json:"-"`
}

// threatEventTagNewResponseExternalReferenceJSON contains the JSON metadata for
// the struct [ThreatEventTagNewResponseExternalReference]
type threatEventTagNewResponseExternalReferenceJSON struct {
	URL         apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponseExternalReference) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseExternalReferenceJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewResponseExternalReferencesAnnotated struct {
	Value string                                                   `json:"value" api:"required"`
	TLP   ThreatEventTagNewResponseExternalReferencesAnnotatedTLP  `json:"tlp"`
	JSON  threatEventTagNewResponseExternalReferencesAnnotatedJSON `json:"-"`
}

// threatEventTagNewResponseExternalReferencesAnnotatedJSON contains the JSON
// metadata for the struct [ThreatEventTagNewResponseExternalReferencesAnnotated]
type threatEventTagNewResponseExternalReferencesAnnotatedJSON struct {
	Value       apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponseExternalReferencesAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseExternalReferencesAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewResponseExternalReferencesAnnotatedTLP string

const (
	ThreatEventTagNewResponseExternalReferencesAnnotatedTLPRed         ThreatEventTagNewResponseExternalReferencesAnnotatedTLP = "red"
	ThreatEventTagNewResponseExternalReferencesAnnotatedTLPAmber       ThreatEventTagNewResponseExternalReferencesAnnotatedTLP = "amber"
	ThreatEventTagNewResponseExternalReferencesAnnotatedTLPAmberStrict ThreatEventTagNewResponseExternalReferencesAnnotatedTLP = "amber+strict"
	ThreatEventTagNewResponseExternalReferencesAnnotatedTLPGreen       ThreatEventTagNewResponseExternalReferencesAnnotatedTLP = "green"
	ThreatEventTagNewResponseExternalReferencesAnnotatedTLPClear       ThreatEventTagNewResponseExternalReferencesAnnotatedTLP = "clear"
	ThreatEventTagNewResponseExternalReferencesAnnotatedTLPPurple      ThreatEventTagNewResponseExternalReferencesAnnotatedTLP = "purple"
)

func (r ThreatEventTagNewResponseExternalReferencesAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseExternalReferencesAnnotatedTLPRed, ThreatEventTagNewResponseExternalReferencesAnnotatedTLPAmber, ThreatEventTagNewResponseExternalReferencesAnnotatedTLPAmberStrict, ThreatEventTagNewResponseExternalReferencesAnnotatedTLPGreen, ThreatEventTagNewResponseExternalReferencesAnnotatedTLPClear, ThreatEventTagNewResponseExternalReferencesAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagNewResponseInternalAlias struct {
	Value      string                                      `json:"value" api:"required"`
	Confidence int64                                       `json:"confidence" api:"nullable"`
	TLP        ThreatEventTagNewResponseInternalAliasesTLP `json:"tlp" api:"nullable"`
	JSON       threatEventTagNewResponseInternalAliasJSON  `json:"-"`
}

// threatEventTagNewResponseInternalAliasJSON contains the JSON metadata for the
// struct [ThreatEventTagNewResponseInternalAlias]
type threatEventTagNewResponseInternalAliasJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponseInternalAlias) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseInternalAliasJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewResponseInternalAliasesTLP string

const (
	ThreatEventTagNewResponseInternalAliasesTLPRed         ThreatEventTagNewResponseInternalAliasesTLP = "red"
	ThreatEventTagNewResponseInternalAliasesTLPAmber       ThreatEventTagNewResponseInternalAliasesTLP = "amber"
	ThreatEventTagNewResponseInternalAliasesTLPAmberStrict ThreatEventTagNewResponseInternalAliasesTLP = "amber+strict"
	ThreatEventTagNewResponseInternalAliasesTLPGreen       ThreatEventTagNewResponseInternalAliasesTLP = "green"
	ThreatEventTagNewResponseInternalAliasesTLPClear       ThreatEventTagNewResponseInternalAliasesTLP = "clear"
	ThreatEventTagNewResponseInternalAliasesTLPPurple      ThreatEventTagNewResponseInternalAliasesTLP = "purple"
)

func (r ThreatEventTagNewResponseInternalAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseInternalAliasesTLPRed, ThreatEventTagNewResponseInternalAliasesTLPAmber, ThreatEventTagNewResponseInternalAliasesTLPAmberStrict, ThreatEventTagNewResponseInternalAliasesTLPGreen, ThreatEventTagNewResponseInternalAliasesTLPClear, ThreatEventTagNewResponseInternalAliasesTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagNewResponseMotiveAnnotated struct {
	Value      string                                       `json:"value" api:"required"`
	Confidence float64                                      `json:"confidence"`
	TLP        ThreatEventTagNewResponseMotiveAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagNewResponseMotiveAnnotatedJSON `json:"-"`
}

// threatEventTagNewResponseMotiveAnnotatedJSON contains the JSON metadata for the
// struct [ThreatEventTagNewResponseMotiveAnnotated]
type threatEventTagNewResponseMotiveAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponseMotiveAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseMotiveAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewResponseMotiveAnnotatedTLP string

const (
	ThreatEventTagNewResponseMotiveAnnotatedTLPRed         ThreatEventTagNewResponseMotiveAnnotatedTLP = "red"
	ThreatEventTagNewResponseMotiveAnnotatedTLPAmber       ThreatEventTagNewResponseMotiveAnnotatedTLP = "amber"
	ThreatEventTagNewResponseMotiveAnnotatedTLPAmberStrict ThreatEventTagNewResponseMotiveAnnotatedTLP = "amber+strict"
	ThreatEventTagNewResponseMotiveAnnotatedTLPGreen       ThreatEventTagNewResponseMotiveAnnotatedTLP = "green"
	ThreatEventTagNewResponseMotiveAnnotatedTLPClear       ThreatEventTagNewResponseMotiveAnnotatedTLP = "clear"
	ThreatEventTagNewResponseMotiveAnnotatedTLPPurple      ThreatEventTagNewResponseMotiveAnnotatedTLP = "purple"
)

func (r ThreatEventTagNewResponseMotiveAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseMotiveAnnotatedTLPRed, ThreatEventTagNewResponseMotiveAnnotatedTLPAmber, ThreatEventTagNewResponseMotiveAnnotatedTLPAmberStrict, ThreatEventTagNewResponseMotiveAnnotatedTLPGreen, ThreatEventTagNewResponseMotiveAnnotatedTLPClear, ThreatEventTagNewResponseMotiveAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagNewResponseOpsecLevelAnnotated struct {
	Value      string                                           `json:"value" api:"required"`
	Confidence float64                                          `json:"confidence"`
	TLP        ThreatEventTagNewResponseOpsecLevelAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagNewResponseOpsecLevelAnnotatedJSON `json:"-"`
}

// threatEventTagNewResponseOpsecLevelAnnotatedJSON contains the JSON metadata for
// the struct [ThreatEventTagNewResponseOpsecLevelAnnotated]
type threatEventTagNewResponseOpsecLevelAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponseOpsecLevelAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseOpsecLevelAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewResponseOpsecLevelAnnotatedTLP string

const (
	ThreatEventTagNewResponseOpsecLevelAnnotatedTLPRed         ThreatEventTagNewResponseOpsecLevelAnnotatedTLP = "red"
	ThreatEventTagNewResponseOpsecLevelAnnotatedTLPAmber       ThreatEventTagNewResponseOpsecLevelAnnotatedTLP = "amber"
	ThreatEventTagNewResponseOpsecLevelAnnotatedTLPAmberStrict ThreatEventTagNewResponseOpsecLevelAnnotatedTLP = "amber+strict"
	ThreatEventTagNewResponseOpsecLevelAnnotatedTLPGreen       ThreatEventTagNewResponseOpsecLevelAnnotatedTLP = "green"
	ThreatEventTagNewResponseOpsecLevelAnnotatedTLPClear       ThreatEventTagNewResponseOpsecLevelAnnotatedTLP = "clear"
	ThreatEventTagNewResponseOpsecLevelAnnotatedTLPPurple      ThreatEventTagNewResponseOpsecLevelAnnotatedTLP = "purple"
)

func (r ThreatEventTagNewResponseOpsecLevelAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseOpsecLevelAnnotatedTLPRed, ThreatEventTagNewResponseOpsecLevelAnnotatedTLPAmber, ThreatEventTagNewResponseOpsecLevelAnnotatedTLPAmberStrict, ThreatEventTagNewResponseOpsecLevelAnnotatedTLPGreen, ThreatEventTagNewResponseOpsecLevelAnnotatedTLPClear, ThreatEventTagNewResponseOpsecLevelAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagNewResponseOriginCountryISOAnnotated struct {
	Value      string                                                 `json:"value" api:"required,nullable"`
	Confidence float64                                                `json:"confidence"`
	TLP        ThreatEventTagNewResponseOriginCountryISOAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagNewResponseOriginCountryISOAnnotatedJSON `json:"-"`
}

// threatEventTagNewResponseOriginCountryISOAnnotatedJSON contains the JSON
// metadata for the struct [ThreatEventTagNewResponseOriginCountryISOAnnotated]
type threatEventTagNewResponseOriginCountryISOAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponseOriginCountryISOAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseOriginCountryISOAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewResponseOriginCountryISOAnnotatedTLP string

const (
	ThreatEventTagNewResponseOriginCountryISOAnnotatedTLPRed         ThreatEventTagNewResponseOriginCountryISOAnnotatedTLP = "red"
	ThreatEventTagNewResponseOriginCountryISOAnnotatedTLPAmber       ThreatEventTagNewResponseOriginCountryISOAnnotatedTLP = "amber"
	ThreatEventTagNewResponseOriginCountryISOAnnotatedTLPAmberStrict ThreatEventTagNewResponseOriginCountryISOAnnotatedTLP = "amber+strict"
	ThreatEventTagNewResponseOriginCountryISOAnnotatedTLPGreen       ThreatEventTagNewResponseOriginCountryISOAnnotatedTLP = "green"
	ThreatEventTagNewResponseOriginCountryISOAnnotatedTLPClear       ThreatEventTagNewResponseOriginCountryISOAnnotatedTLP = "clear"
	ThreatEventTagNewResponseOriginCountryISOAnnotatedTLPPurple      ThreatEventTagNewResponseOriginCountryISOAnnotatedTLP = "purple"
)

func (r ThreatEventTagNewResponseOriginCountryISOAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseOriginCountryISOAnnotatedTLPRed, ThreatEventTagNewResponseOriginCountryISOAnnotatedTLPAmber, ThreatEventTagNewResponseOriginCountryISOAnnotatedTLPAmberStrict, ThreatEventTagNewResponseOriginCountryISOAnnotatedTLPGreen, ThreatEventTagNewResponseOriginCountryISOAnnotatedTLPClear, ThreatEventTagNewResponseOriginCountryISOAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagNewResponsePriorityAnnotated struct {
	Value float64                                        `json:"value" api:"required"`
	TLP   ThreatEventTagNewResponsePriorityAnnotatedTLP  `json:"tlp"`
	JSON  threatEventTagNewResponsePriorityAnnotatedJSON `json:"-"`
}

// threatEventTagNewResponsePriorityAnnotatedJSON contains the JSON metadata for
// the struct [ThreatEventTagNewResponsePriorityAnnotated]
type threatEventTagNewResponsePriorityAnnotatedJSON struct {
	Value       apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponsePriorityAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponsePriorityAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewResponsePriorityAnnotatedTLP string

const (
	ThreatEventTagNewResponsePriorityAnnotatedTLPRed         ThreatEventTagNewResponsePriorityAnnotatedTLP = "red"
	ThreatEventTagNewResponsePriorityAnnotatedTLPAmber       ThreatEventTagNewResponsePriorityAnnotatedTLP = "amber"
	ThreatEventTagNewResponsePriorityAnnotatedTLPAmberStrict ThreatEventTagNewResponsePriorityAnnotatedTLP = "amber+strict"
	ThreatEventTagNewResponsePriorityAnnotatedTLPGreen       ThreatEventTagNewResponsePriorityAnnotatedTLP = "green"
	ThreatEventTagNewResponsePriorityAnnotatedTLPClear       ThreatEventTagNewResponsePriorityAnnotatedTLP = "clear"
	ThreatEventTagNewResponsePriorityAnnotatedTLPPurple      ThreatEventTagNewResponsePriorityAnnotatedTLP = "purple"
)

func (r ThreatEventTagNewResponsePriorityAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponsePriorityAnnotatedTLPRed, ThreatEventTagNewResponsePriorityAnnotatedTLPAmber, ThreatEventTagNewResponsePriorityAnnotatedTLPAmberStrict, ThreatEventTagNewResponsePriorityAnnotatedTLPGreen, ThreatEventTagNewResponsePriorityAnnotatedTLPClear, ThreatEventTagNewResponsePriorityAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagNewResponseSophisticationLevelAnnotated struct {
	Value      string                                                    `json:"value" api:"required"`
	Confidence float64                                                   `json:"confidence"`
	TLP        ThreatEventTagNewResponseSophisticationLevelAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagNewResponseSophisticationLevelAnnotatedJSON `json:"-"`
}

// threatEventTagNewResponseSophisticationLevelAnnotatedJSON contains the JSON
// metadata for the struct [ThreatEventTagNewResponseSophisticationLevelAnnotated]
type threatEventTagNewResponseSophisticationLevelAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagNewResponseSophisticationLevelAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseSophisticationLevelAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewResponseSophisticationLevelAnnotatedTLP string

const (
	ThreatEventTagNewResponseSophisticationLevelAnnotatedTLPRed         ThreatEventTagNewResponseSophisticationLevelAnnotatedTLP = "red"
	ThreatEventTagNewResponseSophisticationLevelAnnotatedTLPAmber       ThreatEventTagNewResponseSophisticationLevelAnnotatedTLP = "amber"
	ThreatEventTagNewResponseSophisticationLevelAnnotatedTLPAmberStrict ThreatEventTagNewResponseSophisticationLevelAnnotatedTLP = "amber+strict"
	ThreatEventTagNewResponseSophisticationLevelAnnotatedTLPGreen       ThreatEventTagNewResponseSophisticationLevelAnnotatedTLP = "green"
	ThreatEventTagNewResponseSophisticationLevelAnnotatedTLPClear       ThreatEventTagNewResponseSophisticationLevelAnnotatedTLP = "clear"
	ThreatEventTagNewResponseSophisticationLevelAnnotatedTLPPurple      ThreatEventTagNewResponseSophisticationLevelAnnotatedTLP = "purple"
)

func (r ThreatEventTagNewResponseSophisticationLevelAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseSophisticationLevelAnnotatedTLPRed, ThreatEventTagNewResponseSophisticationLevelAnnotatedTLPAmber, ThreatEventTagNewResponseSophisticationLevelAnnotatedTLPAmberStrict, ThreatEventTagNewResponseSophisticationLevelAnnotatedTLPGreen, ThreatEventTagNewResponseSophisticationLevelAnnotatedTLPClear, ThreatEventTagNewResponseSophisticationLevelAnnotatedTLPPurple:
		return true
	}
	return false
}

// Tag-level TLP handling marking.
type ThreatEventTagNewResponseTLP string

const (
	ThreatEventTagNewResponseTLPRed         ThreatEventTagNewResponseTLP = "red"
	ThreatEventTagNewResponseTLPAmber       ThreatEventTagNewResponseTLP = "amber"
	ThreatEventTagNewResponseTLPAmberStrict ThreatEventTagNewResponseTLP = "amber+strict"
	ThreatEventTagNewResponseTLPGreen       ThreatEventTagNewResponseTLP = "green"
	ThreatEventTagNewResponseTLPClear       ThreatEventTagNewResponseTLP = "clear"
	ThreatEventTagNewResponseTLPPurple      ThreatEventTagNewResponseTLP = "purple"
)

func (r ThreatEventTagNewResponseTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseTLPRed, ThreatEventTagNewResponseTLPAmber, ThreatEventTagNewResponseTLPAmberStrict, ThreatEventTagNewResponseTLPGreen, ThreatEventTagNewResponseTLPClear, ThreatEventTagNewResponseTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagListResponse struct {
	Pagination ThreatEventTagListResponsePagination `json:"pagination" api:"required"`
	Tags       []ThreatEventTagListResponseTag      `json:"tags" api:"required"`
	JSON       threatEventTagListResponseJSON       `json:"-"`
}

// threatEventTagListResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagListResponse]
type threatEventTagListResponseJSON struct {
	Pagination  apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponsePagination struct {
	Page       float64                                  `json:"page" api:"required"`
	PageSize   float64                                  `json:"pageSize" api:"required"`
	TotalCount float64                                  `json:"totalCount" api:"required"`
	TotalPages float64                                  `json:"totalPages" api:"required"`
	JSON       threatEventTagListResponsePaginationJSON `json:"-"`
}

// threatEventTagListResponsePaginationJSON contains the JSON metadata for the
// struct [ThreatEventTagListResponsePagination]
type threatEventTagListResponsePaginationJSON struct {
	Page        apijson.Field
	PageSize    apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTag struct {
	UUID                    string                                                `json:"uuid" api:"required"`
	Value                   string                                                `json:"value" api:"required"`
	ActiveDuration          string                                                `json:"activeDuration"`
	ActiveDurationAnnotated ThreatEventTagListResponseTagsActiveDurationAnnotated `json:"activeDuration_annotated" api:"nullable"`
	ActorCategory           string                                                `json:"actorCategory"`
	ActorCategoryAnnotated  ThreatEventTagListResponseTagsActorCategoryAnnotated  `json:"actorCategory_annotated" api:"nullable"`
	// Structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: stripped from
	// responses to non-CFONE accounts.
	Aliases                          []ThreatEventTagListResponseTagsAlias                          `json:"aliases"`
	AliasGroupNames                  []string                                                       `json:"aliasGroupNames"`
	AliasGroupNamesInternal          []string                                                       `json:"aliasGroupNamesInternal"`
	AttributionOrganization          string                                                         `json:"attributionOrganization"`
	AttributionOrganizationAnnotated ThreatEventTagListResponseTagsAttributionOrganizationAnnotated `json:"attributionOrganization_annotated" api:"nullable"`
	CategoryName                     string                                                         `json:"categoryName"`
	CategoryUUID                     string                                                         `json:"categoryUuid"`
	// Overall tag confidence (1-10).
	Confidence             int64    `json:"confidence" api:"nullable"`
	CreatedAt              string   `json:"createdAt"`
	DateOfDiscovery        string   `json:"dateOfDiscovery"`
	Description            string   `json:"description"`
	ExternalReferenceLinks []string `json:"externalReferenceLinks"`
	// Structured external references ({ url, description }). Public: returned to all
	// accounts.
	ExternalReferences          []ThreatEventTagListResponseTagsExternalReference           `json:"externalReferences"`
	ExternalReferencesAnnotated []ThreatEventTagListResponseTagsExternalReferencesAnnotated `json:"externalReferences_annotated" api:"nullable"`
	// Internal structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: never
	// returned to non-CFONE accounts.
	InternalAliases     []ThreatEventTagListResponseTagsInternalAlias     `json:"internalAliases"`
	InternalDescription string                                            `json:"internalDescription"`
	LastSeen            string                                            `json:"lastSeen"`
	Motive              string                                            `json:"motive"`
	MotiveAnnotated     ThreatEventTagListResponseTagsMotiveAnnotated     `json:"motive_annotated" api:"nullable"`
	OpsecLevel          string                                            `json:"opsecLevel"`
	OpsecLevelAnnotated ThreatEventTagListResponseTagsOpsecLevelAnnotated `json:"opsecLevel_annotated" api:"nullable"`
	// ISO country code (alpha-2 or alpha-3). Normalized to uppercase on read. Null
	// when stored value is blank/whitespace.
	OriginCountryISO          string                                                  `json:"originCountryISO" api:"nullable"`
	OriginCountryISOAnnotated ThreatEventTagListResponseTagsOriginCountryISOAnnotated `json:"originCountryISO_annotated" api:"nullable"`
	Priority                  float64                                                 `json:"priority"`
	PriorityAnnotated         ThreatEventTagListResponseTagsPriorityAnnotated         `json:"priority_annotated" api:"nullable"`
	// Parsed custom field values. Null when the tag has no custom fields.
	Properties                   map[string]interface{}                                     `json:"properties" api:"nullable"`
	SophisticationLevel          string                                                     `json:"sophisticationLevel"`
	SophisticationLevelAnnotated ThreatEventTagListResponseTagsSophisticationLevelAnnotated `json:"sophisticationLevel_annotated" api:"nullable"`
	// Tag-level TLP handling marking.
	TLP       ThreatEventTagListResponseTagsTLP `json:"tlp" api:"nullable"`
	UpdatedAt string                            `json:"updatedAt"`
	Version   float64                           `json:"version"`
	JSON      threatEventTagListResponseTagJSON `json:"-"`
}

// threatEventTagListResponseTagJSON contains the JSON metadata for the struct
// [ThreatEventTagListResponseTag]
type threatEventTagListResponseTagJSON struct {
	UUID                             apijson.Field
	Value                            apijson.Field
	ActiveDuration                   apijson.Field
	ActiveDurationAnnotated          apijson.Field
	ActorCategory                    apijson.Field
	ActorCategoryAnnotated           apijson.Field
	Aliases                          apijson.Field
	AliasGroupNames                  apijson.Field
	AliasGroupNamesInternal          apijson.Field
	AttributionOrganization          apijson.Field
	AttributionOrganizationAnnotated apijson.Field
	CategoryName                     apijson.Field
	CategoryUUID                     apijson.Field
	Confidence                       apijson.Field
	CreatedAt                        apijson.Field
	DateOfDiscovery                  apijson.Field
	Description                      apijson.Field
	ExternalReferenceLinks           apijson.Field
	ExternalReferences               apijson.Field
	ExternalReferencesAnnotated      apijson.Field
	InternalAliases                  apijson.Field
	InternalDescription              apijson.Field
	LastSeen                         apijson.Field
	Motive                           apijson.Field
	MotiveAnnotated                  apijson.Field
	OpsecLevel                       apijson.Field
	OpsecLevelAnnotated              apijson.Field
	OriginCountryISO                 apijson.Field
	OriginCountryISOAnnotated        apijson.Field
	Priority                         apijson.Field
	PriorityAnnotated                apijson.Field
	Properties                       apijson.Field
	SophisticationLevel              apijson.Field
	SophisticationLevelAnnotated     apijson.Field
	TLP                              apijson.Field
	UpdatedAt                        apijson.Field
	Version                          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTagsActiveDurationAnnotated struct {
	Value string                                                    `json:"value" api:"required"`
	TLP   ThreatEventTagListResponseTagsActiveDurationAnnotatedTLP  `json:"tlp"`
	JSON  threatEventTagListResponseTagsActiveDurationAnnotatedJSON `json:"-"`
}

// threatEventTagListResponseTagsActiveDurationAnnotatedJSON contains the JSON
// metadata for the struct [ThreatEventTagListResponseTagsActiveDurationAnnotated]
type threatEventTagListResponseTagsActiveDurationAnnotatedJSON struct {
	Value       apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTagsActiveDurationAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagsActiveDurationAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTagsActiveDurationAnnotatedTLP string

const (
	ThreatEventTagListResponseTagsActiveDurationAnnotatedTLPRed         ThreatEventTagListResponseTagsActiveDurationAnnotatedTLP = "red"
	ThreatEventTagListResponseTagsActiveDurationAnnotatedTLPAmber       ThreatEventTagListResponseTagsActiveDurationAnnotatedTLP = "amber"
	ThreatEventTagListResponseTagsActiveDurationAnnotatedTLPAmberStrict ThreatEventTagListResponseTagsActiveDurationAnnotatedTLP = "amber+strict"
	ThreatEventTagListResponseTagsActiveDurationAnnotatedTLPGreen       ThreatEventTagListResponseTagsActiveDurationAnnotatedTLP = "green"
	ThreatEventTagListResponseTagsActiveDurationAnnotatedTLPClear       ThreatEventTagListResponseTagsActiveDurationAnnotatedTLP = "clear"
	ThreatEventTagListResponseTagsActiveDurationAnnotatedTLPPurple      ThreatEventTagListResponseTagsActiveDurationAnnotatedTLP = "purple"
)

func (r ThreatEventTagListResponseTagsActiveDurationAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsActiveDurationAnnotatedTLPRed, ThreatEventTagListResponseTagsActiveDurationAnnotatedTLPAmber, ThreatEventTagListResponseTagsActiveDurationAnnotatedTLPAmberStrict, ThreatEventTagListResponseTagsActiveDurationAnnotatedTLPGreen, ThreatEventTagListResponseTagsActiveDurationAnnotatedTLPClear, ThreatEventTagListResponseTagsActiveDurationAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagListResponseTagsActorCategoryAnnotated struct {
	Value      string                                                   `json:"value" api:"required"`
	Confidence float64                                                  `json:"confidence"`
	TLP        ThreatEventTagListResponseTagsActorCategoryAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagListResponseTagsActorCategoryAnnotatedJSON `json:"-"`
}

// threatEventTagListResponseTagsActorCategoryAnnotatedJSON contains the JSON
// metadata for the struct [ThreatEventTagListResponseTagsActorCategoryAnnotated]
type threatEventTagListResponseTagsActorCategoryAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTagsActorCategoryAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagsActorCategoryAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTagsActorCategoryAnnotatedTLP string

const (
	ThreatEventTagListResponseTagsActorCategoryAnnotatedTLPRed         ThreatEventTagListResponseTagsActorCategoryAnnotatedTLP = "red"
	ThreatEventTagListResponseTagsActorCategoryAnnotatedTLPAmber       ThreatEventTagListResponseTagsActorCategoryAnnotatedTLP = "amber"
	ThreatEventTagListResponseTagsActorCategoryAnnotatedTLPAmberStrict ThreatEventTagListResponseTagsActorCategoryAnnotatedTLP = "amber+strict"
	ThreatEventTagListResponseTagsActorCategoryAnnotatedTLPGreen       ThreatEventTagListResponseTagsActorCategoryAnnotatedTLP = "green"
	ThreatEventTagListResponseTagsActorCategoryAnnotatedTLPClear       ThreatEventTagListResponseTagsActorCategoryAnnotatedTLP = "clear"
	ThreatEventTagListResponseTagsActorCategoryAnnotatedTLPPurple      ThreatEventTagListResponseTagsActorCategoryAnnotatedTLP = "purple"
)

func (r ThreatEventTagListResponseTagsActorCategoryAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsActorCategoryAnnotatedTLPRed, ThreatEventTagListResponseTagsActorCategoryAnnotatedTLPAmber, ThreatEventTagListResponseTagsActorCategoryAnnotatedTLPAmberStrict, ThreatEventTagListResponseTagsActorCategoryAnnotatedTLPGreen, ThreatEventTagListResponseTagsActorCategoryAnnotatedTLPClear, ThreatEventTagListResponseTagsActorCategoryAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagListResponseTagsAlias struct {
	Value      string                                   `json:"value" api:"required"`
	Confidence int64                                    `json:"confidence" api:"nullable"`
	TLP        ThreatEventTagListResponseTagsAliasesTLP `json:"tlp" api:"nullable"`
	JSON       threatEventTagListResponseTagsAliasJSON  `json:"-"`
}

// threatEventTagListResponseTagsAliasJSON contains the JSON metadata for the
// struct [ThreatEventTagListResponseTagsAlias]
type threatEventTagListResponseTagsAliasJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTagsAlias) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagsAliasJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTagsAliasesTLP string

const (
	ThreatEventTagListResponseTagsAliasesTLPRed         ThreatEventTagListResponseTagsAliasesTLP = "red"
	ThreatEventTagListResponseTagsAliasesTLPAmber       ThreatEventTagListResponseTagsAliasesTLP = "amber"
	ThreatEventTagListResponseTagsAliasesTLPAmberStrict ThreatEventTagListResponseTagsAliasesTLP = "amber+strict"
	ThreatEventTagListResponseTagsAliasesTLPGreen       ThreatEventTagListResponseTagsAliasesTLP = "green"
	ThreatEventTagListResponseTagsAliasesTLPClear       ThreatEventTagListResponseTagsAliasesTLP = "clear"
	ThreatEventTagListResponseTagsAliasesTLPPurple      ThreatEventTagListResponseTagsAliasesTLP = "purple"
)

func (r ThreatEventTagListResponseTagsAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsAliasesTLPRed, ThreatEventTagListResponseTagsAliasesTLPAmber, ThreatEventTagListResponseTagsAliasesTLPAmberStrict, ThreatEventTagListResponseTagsAliasesTLPGreen, ThreatEventTagListResponseTagsAliasesTLPClear, ThreatEventTagListResponseTagsAliasesTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagListResponseTagsAttributionOrganizationAnnotated struct {
	Value      string                                                             `json:"value" api:"required"`
	Confidence float64                                                            `json:"confidence"`
	TLP        ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagListResponseTagsAttributionOrganizationAnnotatedJSON `json:"-"`
}

// threatEventTagListResponseTagsAttributionOrganizationAnnotatedJSON contains the
// JSON metadata for the struct
// [ThreatEventTagListResponseTagsAttributionOrganizationAnnotated]
type threatEventTagListResponseTagsAttributionOrganizationAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTagsAttributionOrganizationAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagsAttributionOrganizationAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLP string

const (
	ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLPRed         ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLP = "red"
	ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLPAmber       ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLP = "amber"
	ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLPAmberStrict ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLP = "amber+strict"
	ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLPGreen       ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLP = "green"
	ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLPClear       ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLP = "clear"
	ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLPPurple      ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLP = "purple"
)

func (r ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLPRed, ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLPAmber, ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLPAmberStrict, ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLPGreen, ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLPClear, ThreatEventTagListResponseTagsAttributionOrganizationAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagListResponseTagsExternalReference struct {
	URL         string                                              `json:"url" api:"required"`
	Description string                                              `json:"description" api:"nullable"`
	JSON        threatEventTagListResponseTagsExternalReferenceJSON `json:"-"`
}

// threatEventTagListResponseTagsExternalReferenceJSON contains the JSON metadata
// for the struct [ThreatEventTagListResponseTagsExternalReference]
type threatEventTagListResponseTagsExternalReferenceJSON struct {
	URL         apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTagsExternalReference) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagsExternalReferenceJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTagsExternalReferencesAnnotated struct {
	Value string                                                        `json:"value" api:"required"`
	TLP   ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLP  `json:"tlp"`
	JSON  threatEventTagListResponseTagsExternalReferencesAnnotatedJSON `json:"-"`
}

// threatEventTagListResponseTagsExternalReferencesAnnotatedJSON contains the JSON
// metadata for the struct
// [ThreatEventTagListResponseTagsExternalReferencesAnnotated]
type threatEventTagListResponseTagsExternalReferencesAnnotatedJSON struct {
	Value       apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTagsExternalReferencesAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagsExternalReferencesAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLP string

const (
	ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLPRed         ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLP = "red"
	ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLPAmber       ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLP = "amber"
	ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLPAmberStrict ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLP = "amber+strict"
	ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLPGreen       ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLP = "green"
	ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLPClear       ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLP = "clear"
	ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLPPurple      ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLP = "purple"
)

func (r ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLPRed, ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLPAmber, ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLPAmberStrict, ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLPGreen, ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLPClear, ThreatEventTagListResponseTagsExternalReferencesAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagListResponseTagsInternalAlias struct {
	Value      string                                           `json:"value" api:"required"`
	Confidence int64                                            `json:"confidence" api:"nullable"`
	TLP        ThreatEventTagListResponseTagsInternalAliasesTLP `json:"tlp" api:"nullable"`
	JSON       threatEventTagListResponseTagsInternalAliasJSON  `json:"-"`
}

// threatEventTagListResponseTagsInternalAliasJSON contains the JSON metadata for
// the struct [ThreatEventTagListResponseTagsInternalAlias]
type threatEventTagListResponseTagsInternalAliasJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTagsInternalAlias) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagsInternalAliasJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTagsInternalAliasesTLP string

const (
	ThreatEventTagListResponseTagsInternalAliasesTLPRed         ThreatEventTagListResponseTagsInternalAliasesTLP = "red"
	ThreatEventTagListResponseTagsInternalAliasesTLPAmber       ThreatEventTagListResponseTagsInternalAliasesTLP = "amber"
	ThreatEventTagListResponseTagsInternalAliasesTLPAmberStrict ThreatEventTagListResponseTagsInternalAliasesTLP = "amber+strict"
	ThreatEventTagListResponseTagsInternalAliasesTLPGreen       ThreatEventTagListResponseTagsInternalAliasesTLP = "green"
	ThreatEventTagListResponseTagsInternalAliasesTLPClear       ThreatEventTagListResponseTagsInternalAliasesTLP = "clear"
	ThreatEventTagListResponseTagsInternalAliasesTLPPurple      ThreatEventTagListResponseTagsInternalAliasesTLP = "purple"
)

func (r ThreatEventTagListResponseTagsInternalAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsInternalAliasesTLPRed, ThreatEventTagListResponseTagsInternalAliasesTLPAmber, ThreatEventTagListResponseTagsInternalAliasesTLPAmberStrict, ThreatEventTagListResponseTagsInternalAliasesTLPGreen, ThreatEventTagListResponseTagsInternalAliasesTLPClear, ThreatEventTagListResponseTagsInternalAliasesTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagListResponseTagsMotiveAnnotated struct {
	Value      string                                            `json:"value" api:"required"`
	Confidence float64                                           `json:"confidence"`
	TLP        ThreatEventTagListResponseTagsMotiveAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagListResponseTagsMotiveAnnotatedJSON `json:"-"`
}

// threatEventTagListResponseTagsMotiveAnnotatedJSON contains the JSON metadata for
// the struct [ThreatEventTagListResponseTagsMotiveAnnotated]
type threatEventTagListResponseTagsMotiveAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTagsMotiveAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagsMotiveAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTagsMotiveAnnotatedTLP string

const (
	ThreatEventTagListResponseTagsMotiveAnnotatedTLPRed         ThreatEventTagListResponseTagsMotiveAnnotatedTLP = "red"
	ThreatEventTagListResponseTagsMotiveAnnotatedTLPAmber       ThreatEventTagListResponseTagsMotiveAnnotatedTLP = "amber"
	ThreatEventTagListResponseTagsMotiveAnnotatedTLPAmberStrict ThreatEventTagListResponseTagsMotiveAnnotatedTLP = "amber+strict"
	ThreatEventTagListResponseTagsMotiveAnnotatedTLPGreen       ThreatEventTagListResponseTagsMotiveAnnotatedTLP = "green"
	ThreatEventTagListResponseTagsMotiveAnnotatedTLPClear       ThreatEventTagListResponseTagsMotiveAnnotatedTLP = "clear"
	ThreatEventTagListResponseTagsMotiveAnnotatedTLPPurple      ThreatEventTagListResponseTagsMotiveAnnotatedTLP = "purple"
)

func (r ThreatEventTagListResponseTagsMotiveAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsMotiveAnnotatedTLPRed, ThreatEventTagListResponseTagsMotiveAnnotatedTLPAmber, ThreatEventTagListResponseTagsMotiveAnnotatedTLPAmberStrict, ThreatEventTagListResponseTagsMotiveAnnotatedTLPGreen, ThreatEventTagListResponseTagsMotiveAnnotatedTLPClear, ThreatEventTagListResponseTagsMotiveAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagListResponseTagsOpsecLevelAnnotated struct {
	Value      string                                                `json:"value" api:"required"`
	Confidence float64                                               `json:"confidence"`
	TLP        ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagListResponseTagsOpsecLevelAnnotatedJSON `json:"-"`
}

// threatEventTagListResponseTagsOpsecLevelAnnotatedJSON contains the JSON metadata
// for the struct [ThreatEventTagListResponseTagsOpsecLevelAnnotated]
type threatEventTagListResponseTagsOpsecLevelAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTagsOpsecLevelAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagsOpsecLevelAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLP string

const (
	ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLPRed         ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLP = "red"
	ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLPAmber       ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLP = "amber"
	ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLPAmberStrict ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLP = "amber+strict"
	ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLPGreen       ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLP = "green"
	ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLPClear       ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLP = "clear"
	ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLPPurple      ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLP = "purple"
)

func (r ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLPRed, ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLPAmber, ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLPAmberStrict, ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLPGreen, ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLPClear, ThreatEventTagListResponseTagsOpsecLevelAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagListResponseTagsOriginCountryISOAnnotated struct {
	Value      string                                                      `json:"value" api:"required,nullable"`
	Confidence float64                                                     `json:"confidence"`
	TLP        ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagListResponseTagsOriginCountryISOAnnotatedJSON `json:"-"`
}

// threatEventTagListResponseTagsOriginCountryISOAnnotatedJSON contains the JSON
// metadata for the struct
// [ThreatEventTagListResponseTagsOriginCountryISOAnnotated]
type threatEventTagListResponseTagsOriginCountryISOAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTagsOriginCountryISOAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagsOriginCountryISOAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLP string

const (
	ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLPRed         ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLP = "red"
	ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLPAmber       ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLP = "amber"
	ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLPAmberStrict ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLP = "amber+strict"
	ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLPGreen       ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLP = "green"
	ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLPClear       ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLP = "clear"
	ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLPPurple      ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLP = "purple"
)

func (r ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLPRed, ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLPAmber, ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLPAmberStrict, ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLPGreen, ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLPClear, ThreatEventTagListResponseTagsOriginCountryISOAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagListResponseTagsPriorityAnnotated struct {
	Value float64                                             `json:"value" api:"required"`
	TLP   ThreatEventTagListResponseTagsPriorityAnnotatedTLP  `json:"tlp"`
	JSON  threatEventTagListResponseTagsPriorityAnnotatedJSON `json:"-"`
}

// threatEventTagListResponseTagsPriorityAnnotatedJSON contains the JSON metadata
// for the struct [ThreatEventTagListResponseTagsPriorityAnnotated]
type threatEventTagListResponseTagsPriorityAnnotatedJSON struct {
	Value       apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTagsPriorityAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagsPriorityAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTagsPriorityAnnotatedTLP string

const (
	ThreatEventTagListResponseTagsPriorityAnnotatedTLPRed         ThreatEventTagListResponseTagsPriorityAnnotatedTLP = "red"
	ThreatEventTagListResponseTagsPriorityAnnotatedTLPAmber       ThreatEventTagListResponseTagsPriorityAnnotatedTLP = "amber"
	ThreatEventTagListResponseTagsPriorityAnnotatedTLPAmberStrict ThreatEventTagListResponseTagsPriorityAnnotatedTLP = "amber+strict"
	ThreatEventTagListResponseTagsPriorityAnnotatedTLPGreen       ThreatEventTagListResponseTagsPriorityAnnotatedTLP = "green"
	ThreatEventTagListResponseTagsPriorityAnnotatedTLPClear       ThreatEventTagListResponseTagsPriorityAnnotatedTLP = "clear"
	ThreatEventTagListResponseTagsPriorityAnnotatedTLPPurple      ThreatEventTagListResponseTagsPriorityAnnotatedTLP = "purple"
)

func (r ThreatEventTagListResponseTagsPriorityAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsPriorityAnnotatedTLPRed, ThreatEventTagListResponseTagsPriorityAnnotatedTLPAmber, ThreatEventTagListResponseTagsPriorityAnnotatedTLPAmberStrict, ThreatEventTagListResponseTagsPriorityAnnotatedTLPGreen, ThreatEventTagListResponseTagsPriorityAnnotatedTLPClear, ThreatEventTagListResponseTagsPriorityAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagListResponseTagsSophisticationLevelAnnotated struct {
	Value      string                                                         `json:"value" api:"required"`
	Confidence float64                                                        `json:"confidence"`
	TLP        ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagListResponseTagsSophisticationLevelAnnotatedJSON `json:"-"`
}

// threatEventTagListResponseTagsSophisticationLevelAnnotatedJSON contains the JSON
// metadata for the struct
// [ThreatEventTagListResponseTagsSophisticationLevelAnnotated]
type threatEventTagListResponseTagsSophisticationLevelAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTagsSophisticationLevelAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagsSophisticationLevelAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLP string

const (
	ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLPRed         ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLP = "red"
	ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLPAmber       ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLP = "amber"
	ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLPAmberStrict ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLP = "amber+strict"
	ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLPGreen       ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLP = "green"
	ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLPClear       ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLP = "clear"
	ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLPPurple      ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLP = "purple"
)

func (r ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLPRed, ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLPAmber, ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLPAmberStrict, ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLPGreen, ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLPClear, ThreatEventTagListResponseTagsSophisticationLevelAnnotatedTLPPurple:
		return true
	}
	return false
}

// Tag-level TLP handling marking.
type ThreatEventTagListResponseTagsTLP string

const (
	ThreatEventTagListResponseTagsTLPRed         ThreatEventTagListResponseTagsTLP = "red"
	ThreatEventTagListResponseTagsTLPAmber       ThreatEventTagListResponseTagsTLP = "amber"
	ThreatEventTagListResponseTagsTLPAmberStrict ThreatEventTagListResponseTagsTLP = "amber+strict"
	ThreatEventTagListResponseTagsTLPGreen       ThreatEventTagListResponseTagsTLP = "green"
	ThreatEventTagListResponseTagsTLPClear       ThreatEventTagListResponseTagsTLP = "clear"
	ThreatEventTagListResponseTagsTLPPurple      ThreatEventTagListResponseTagsTLP = "purple"
)

func (r ThreatEventTagListResponseTagsTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsTLPRed, ThreatEventTagListResponseTagsTLPAmber, ThreatEventTagListResponseTagsTLPAmberStrict, ThreatEventTagListResponseTagsTLPGreen, ThreatEventTagListResponseTagsTLPClear, ThreatEventTagListResponseTagsTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagDeleteResponse struct {
	UUID string                           `json:"uuid" api:"required"`
	JSON threatEventTagDeleteResponseJSON `json:"-"`
}

// threatEventTagDeleteResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagDeleteResponse]
type threatEventTagDeleteResponseJSON struct {
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponse struct {
	UUID                    string                                            `json:"uuid" api:"required"`
	Value                   string                                            `json:"value" api:"required"`
	ActiveDuration          string                                            `json:"activeDuration"`
	ActiveDurationAnnotated ThreatEventTagEditResponseActiveDurationAnnotated `json:"activeDuration_annotated" api:"nullable"`
	ActorCategory           string                                            `json:"actorCategory"`
	ActorCategoryAnnotated  ThreatEventTagEditResponseActorCategoryAnnotated  `json:"actorCategory_annotated" api:"nullable"`
	// Structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: stripped from
	// responses to non-CFONE accounts.
	Aliases                          []ThreatEventTagEditResponseAlias                          `json:"aliases"`
	AliasGroupNames                  []string                                                   `json:"aliasGroupNames"`
	AliasGroupNamesInternal          []string                                                   `json:"aliasGroupNamesInternal"`
	AttributionOrganization          string                                                     `json:"attributionOrganization"`
	AttributionOrganizationAnnotated ThreatEventTagEditResponseAttributionOrganizationAnnotated `json:"attributionOrganization_annotated" api:"nullable"`
	CategoryName                     string                                                     `json:"categoryName"`
	CategoryUUID                     string                                                     `json:"categoryUuid"`
	// Overall tag confidence (1-10).
	Confidence             int64    `json:"confidence" api:"nullable"`
	CreatedAt              string   `json:"createdAt"`
	DateOfDiscovery        string   `json:"dateOfDiscovery"`
	Description            string   `json:"description"`
	ExternalReferenceLinks []string `json:"externalReferenceLinks"`
	// Structured external references ({ url, description }). Public: returned to all
	// accounts.
	ExternalReferences          []ThreatEventTagEditResponseExternalReference           `json:"externalReferences"`
	ExternalReferencesAnnotated []ThreatEventTagEditResponseExternalReferencesAnnotated `json:"externalReferences_annotated" api:"nullable"`
	// Internal structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: never
	// returned to non-CFONE accounts.
	InternalAliases     []ThreatEventTagEditResponseInternalAlias     `json:"internalAliases"`
	InternalDescription string                                        `json:"internalDescription"`
	LastSeen            string                                        `json:"lastSeen"`
	Motive              string                                        `json:"motive"`
	MotiveAnnotated     ThreatEventTagEditResponseMotiveAnnotated     `json:"motive_annotated" api:"nullable"`
	OpsecLevel          string                                        `json:"opsecLevel"`
	OpsecLevelAnnotated ThreatEventTagEditResponseOpsecLevelAnnotated `json:"opsecLevel_annotated" api:"nullable"`
	// ISO country code (alpha-2 or alpha-3). Normalized to uppercase on read. Null
	// when stored value is blank/whitespace.
	OriginCountryISO          string                                              `json:"originCountryISO" api:"nullable"`
	OriginCountryISOAnnotated ThreatEventTagEditResponseOriginCountryISOAnnotated `json:"originCountryISO_annotated" api:"nullable"`
	Priority                  float64                                             `json:"priority"`
	PriorityAnnotated         ThreatEventTagEditResponsePriorityAnnotated         `json:"priority_annotated" api:"nullable"`
	// Parsed custom field values. Null when the tag has no custom fields.
	Properties                   map[string]interface{}                                 `json:"properties" api:"nullable"`
	SophisticationLevel          string                                                 `json:"sophisticationLevel"`
	SophisticationLevelAnnotated ThreatEventTagEditResponseSophisticationLevelAnnotated `json:"sophisticationLevel_annotated" api:"nullable"`
	// Tag-level TLP handling marking.
	TLP       ThreatEventTagEditResponseTLP  `json:"tlp" api:"nullable"`
	UpdatedAt string                         `json:"updatedAt"`
	Version   float64                        `json:"version"`
	JSON      threatEventTagEditResponseJSON `json:"-"`
}

// threatEventTagEditResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagEditResponse]
type threatEventTagEditResponseJSON struct {
	UUID                             apijson.Field
	Value                            apijson.Field
	ActiveDuration                   apijson.Field
	ActiveDurationAnnotated          apijson.Field
	ActorCategory                    apijson.Field
	ActorCategoryAnnotated           apijson.Field
	Aliases                          apijson.Field
	AliasGroupNames                  apijson.Field
	AliasGroupNamesInternal          apijson.Field
	AttributionOrganization          apijson.Field
	AttributionOrganizationAnnotated apijson.Field
	CategoryName                     apijson.Field
	CategoryUUID                     apijson.Field
	Confidence                       apijson.Field
	CreatedAt                        apijson.Field
	DateOfDiscovery                  apijson.Field
	Description                      apijson.Field
	ExternalReferenceLinks           apijson.Field
	ExternalReferences               apijson.Field
	ExternalReferencesAnnotated      apijson.Field
	InternalAliases                  apijson.Field
	InternalDescription              apijson.Field
	LastSeen                         apijson.Field
	Motive                           apijson.Field
	MotiveAnnotated                  apijson.Field
	OpsecLevel                       apijson.Field
	OpsecLevelAnnotated              apijson.Field
	OriginCountryISO                 apijson.Field
	OriginCountryISOAnnotated        apijson.Field
	Priority                         apijson.Field
	PriorityAnnotated                apijson.Field
	Properties                       apijson.Field
	SophisticationLevel              apijson.Field
	SophisticationLevelAnnotated     apijson.Field
	TLP                              apijson.Field
	UpdatedAt                        apijson.Field
	Version                          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *ThreatEventTagEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponseActiveDurationAnnotated struct {
	Value string                                                `json:"value" api:"required"`
	TLP   ThreatEventTagEditResponseActiveDurationAnnotatedTLP  `json:"tlp"`
	JSON  threatEventTagEditResponseActiveDurationAnnotatedJSON `json:"-"`
}

// threatEventTagEditResponseActiveDurationAnnotatedJSON contains the JSON metadata
// for the struct [ThreatEventTagEditResponseActiveDurationAnnotated]
type threatEventTagEditResponseActiveDurationAnnotatedJSON struct {
	Value       apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagEditResponseActiveDurationAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponseActiveDurationAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponseActiveDurationAnnotatedTLP string

const (
	ThreatEventTagEditResponseActiveDurationAnnotatedTLPRed         ThreatEventTagEditResponseActiveDurationAnnotatedTLP = "red"
	ThreatEventTagEditResponseActiveDurationAnnotatedTLPAmber       ThreatEventTagEditResponseActiveDurationAnnotatedTLP = "amber"
	ThreatEventTagEditResponseActiveDurationAnnotatedTLPAmberStrict ThreatEventTagEditResponseActiveDurationAnnotatedTLP = "amber+strict"
	ThreatEventTagEditResponseActiveDurationAnnotatedTLPGreen       ThreatEventTagEditResponseActiveDurationAnnotatedTLP = "green"
	ThreatEventTagEditResponseActiveDurationAnnotatedTLPClear       ThreatEventTagEditResponseActiveDurationAnnotatedTLP = "clear"
	ThreatEventTagEditResponseActiveDurationAnnotatedTLPPurple      ThreatEventTagEditResponseActiveDurationAnnotatedTLP = "purple"
)

func (r ThreatEventTagEditResponseActiveDurationAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseActiveDurationAnnotatedTLPRed, ThreatEventTagEditResponseActiveDurationAnnotatedTLPAmber, ThreatEventTagEditResponseActiveDurationAnnotatedTLPAmberStrict, ThreatEventTagEditResponseActiveDurationAnnotatedTLPGreen, ThreatEventTagEditResponseActiveDurationAnnotatedTLPClear, ThreatEventTagEditResponseActiveDurationAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagEditResponseActorCategoryAnnotated struct {
	Value      string                                               `json:"value" api:"required"`
	Confidence float64                                              `json:"confidence"`
	TLP        ThreatEventTagEditResponseActorCategoryAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagEditResponseActorCategoryAnnotatedJSON `json:"-"`
}

// threatEventTagEditResponseActorCategoryAnnotatedJSON contains the JSON metadata
// for the struct [ThreatEventTagEditResponseActorCategoryAnnotated]
type threatEventTagEditResponseActorCategoryAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagEditResponseActorCategoryAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponseActorCategoryAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponseActorCategoryAnnotatedTLP string

const (
	ThreatEventTagEditResponseActorCategoryAnnotatedTLPRed         ThreatEventTagEditResponseActorCategoryAnnotatedTLP = "red"
	ThreatEventTagEditResponseActorCategoryAnnotatedTLPAmber       ThreatEventTagEditResponseActorCategoryAnnotatedTLP = "amber"
	ThreatEventTagEditResponseActorCategoryAnnotatedTLPAmberStrict ThreatEventTagEditResponseActorCategoryAnnotatedTLP = "amber+strict"
	ThreatEventTagEditResponseActorCategoryAnnotatedTLPGreen       ThreatEventTagEditResponseActorCategoryAnnotatedTLP = "green"
	ThreatEventTagEditResponseActorCategoryAnnotatedTLPClear       ThreatEventTagEditResponseActorCategoryAnnotatedTLP = "clear"
	ThreatEventTagEditResponseActorCategoryAnnotatedTLPPurple      ThreatEventTagEditResponseActorCategoryAnnotatedTLP = "purple"
)

func (r ThreatEventTagEditResponseActorCategoryAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseActorCategoryAnnotatedTLPRed, ThreatEventTagEditResponseActorCategoryAnnotatedTLPAmber, ThreatEventTagEditResponseActorCategoryAnnotatedTLPAmberStrict, ThreatEventTagEditResponseActorCategoryAnnotatedTLPGreen, ThreatEventTagEditResponseActorCategoryAnnotatedTLPClear, ThreatEventTagEditResponseActorCategoryAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagEditResponseAlias struct {
	Value      string                               `json:"value" api:"required"`
	Confidence int64                                `json:"confidence" api:"nullable"`
	TLP        ThreatEventTagEditResponseAliasesTLP `json:"tlp" api:"nullable"`
	JSON       threatEventTagEditResponseAliasJSON  `json:"-"`
}

// threatEventTagEditResponseAliasJSON contains the JSON metadata for the struct
// [ThreatEventTagEditResponseAlias]
type threatEventTagEditResponseAliasJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagEditResponseAlias) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponseAliasJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponseAliasesTLP string

const (
	ThreatEventTagEditResponseAliasesTLPRed         ThreatEventTagEditResponseAliasesTLP = "red"
	ThreatEventTagEditResponseAliasesTLPAmber       ThreatEventTagEditResponseAliasesTLP = "amber"
	ThreatEventTagEditResponseAliasesTLPAmberStrict ThreatEventTagEditResponseAliasesTLP = "amber+strict"
	ThreatEventTagEditResponseAliasesTLPGreen       ThreatEventTagEditResponseAliasesTLP = "green"
	ThreatEventTagEditResponseAliasesTLPClear       ThreatEventTagEditResponseAliasesTLP = "clear"
	ThreatEventTagEditResponseAliasesTLPPurple      ThreatEventTagEditResponseAliasesTLP = "purple"
)

func (r ThreatEventTagEditResponseAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseAliasesTLPRed, ThreatEventTagEditResponseAliasesTLPAmber, ThreatEventTagEditResponseAliasesTLPAmberStrict, ThreatEventTagEditResponseAliasesTLPGreen, ThreatEventTagEditResponseAliasesTLPClear, ThreatEventTagEditResponseAliasesTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagEditResponseAttributionOrganizationAnnotated struct {
	Value      string                                                         `json:"value" api:"required"`
	Confidence float64                                                        `json:"confidence"`
	TLP        ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagEditResponseAttributionOrganizationAnnotatedJSON `json:"-"`
}

// threatEventTagEditResponseAttributionOrganizationAnnotatedJSON contains the JSON
// metadata for the struct
// [ThreatEventTagEditResponseAttributionOrganizationAnnotated]
type threatEventTagEditResponseAttributionOrganizationAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagEditResponseAttributionOrganizationAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponseAttributionOrganizationAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLP string

const (
	ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLPRed         ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLP = "red"
	ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLPAmber       ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLP = "amber"
	ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLPAmberStrict ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLP = "amber+strict"
	ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLPGreen       ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLP = "green"
	ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLPClear       ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLP = "clear"
	ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLPPurple      ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLP = "purple"
)

func (r ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLPRed, ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLPAmber, ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLPAmberStrict, ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLPGreen, ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLPClear, ThreatEventTagEditResponseAttributionOrganizationAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagEditResponseExternalReference struct {
	URL         string                                          `json:"url" api:"required"`
	Description string                                          `json:"description" api:"nullable"`
	JSON        threatEventTagEditResponseExternalReferenceJSON `json:"-"`
}

// threatEventTagEditResponseExternalReferenceJSON contains the JSON metadata for
// the struct [ThreatEventTagEditResponseExternalReference]
type threatEventTagEditResponseExternalReferenceJSON struct {
	URL         apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagEditResponseExternalReference) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponseExternalReferenceJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponseExternalReferencesAnnotated struct {
	Value string                                                    `json:"value" api:"required"`
	TLP   ThreatEventTagEditResponseExternalReferencesAnnotatedTLP  `json:"tlp"`
	JSON  threatEventTagEditResponseExternalReferencesAnnotatedJSON `json:"-"`
}

// threatEventTagEditResponseExternalReferencesAnnotatedJSON contains the JSON
// metadata for the struct [ThreatEventTagEditResponseExternalReferencesAnnotated]
type threatEventTagEditResponseExternalReferencesAnnotatedJSON struct {
	Value       apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagEditResponseExternalReferencesAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponseExternalReferencesAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponseExternalReferencesAnnotatedTLP string

const (
	ThreatEventTagEditResponseExternalReferencesAnnotatedTLPRed         ThreatEventTagEditResponseExternalReferencesAnnotatedTLP = "red"
	ThreatEventTagEditResponseExternalReferencesAnnotatedTLPAmber       ThreatEventTagEditResponseExternalReferencesAnnotatedTLP = "amber"
	ThreatEventTagEditResponseExternalReferencesAnnotatedTLPAmberStrict ThreatEventTagEditResponseExternalReferencesAnnotatedTLP = "amber+strict"
	ThreatEventTagEditResponseExternalReferencesAnnotatedTLPGreen       ThreatEventTagEditResponseExternalReferencesAnnotatedTLP = "green"
	ThreatEventTagEditResponseExternalReferencesAnnotatedTLPClear       ThreatEventTagEditResponseExternalReferencesAnnotatedTLP = "clear"
	ThreatEventTagEditResponseExternalReferencesAnnotatedTLPPurple      ThreatEventTagEditResponseExternalReferencesAnnotatedTLP = "purple"
)

func (r ThreatEventTagEditResponseExternalReferencesAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseExternalReferencesAnnotatedTLPRed, ThreatEventTagEditResponseExternalReferencesAnnotatedTLPAmber, ThreatEventTagEditResponseExternalReferencesAnnotatedTLPAmberStrict, ThreatEventTagEditResponseExternalReferencesAnnotatedTLPGreen, ThreatEventTagEditResponseExternalReferencesAnnotatedTLPClear, ThreatEventTagEditResponseExternalReferencesAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagEditResponseInternalAlias struct {
	Value      string                                       `json:"value" api:"required"`
	Confidence int64                                        `json:"confidence" api:"nullable"`
	TLP        ThreatEventTagEditResponseInternalAliasesTLP `json:"tlp" api:"nullable"`
	JSON       threatEventTagEditResponseInternalAliasJSON  `json:"-"`
}

// threatEventTagEditResponseInternalAliasJSON contains the JSON metadata for the
// struct [ThreatEventTagEditResponseInternalAlias]
type threatEventTagEditResponseInternalAliasJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagEditResponseInternalAlias) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponseInternalAliasJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponseInternalAliasesTLP string

const (
	ThreatEventTagEditResponseInternalAliasesTLPRed         ThreatEventTagEditResponseInternalAliasesTLP = "red"
	ThreatEventTagEditResponseInternalAliasesTLPAmber       ThreatEventTagEditResponseInternalAliasesTLP = "amber"
	ThreatEventTagEditResponseInternalAliasesTLPAmberStrict ThreatEventTagEditResponseInternalAliasesTLP = "amber+strict"
	ThreatEventTagEditResponseInternalAliasesTLPGreen       ThreatEventTagEditResponseInternalAliasesTLP = "green"
	ThreatEventTagEditResponseInternalAliasesTLPClear       ThreatEventTagEditResponseInternalAliasesTLP = "clear"
	ThreatEventTagEditResponseInternalAliasesTLPPurple      ThreatEventTagEditResponseInternalAliasesTLP = "purple"
)

func (r ThreatEventTagEditResponseInternalAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseInternalAliasesTLPRed, ThreatEventTagEditResponseInternalAliasesTLPAmber, ThreatEventTagEditResponseInternalAliasesTLPAmberStrict, ThreatEventTagEditResponseInternalAliasesTLPGreen, ThreatEventTagEditResponseInternalAliasesTLPClear, ThreatEventTagEditResponseInternalAliasesTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagEditResponseMotiveAnnotated struct {
	Value      string                                        `json:"value" api:"required"`
	Confidence float64                                       `json:"confidence"`
	TLP        ThreatEventTagEditResponseMotiveAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagEditResponseMotiveAnnotatedJSON `json:"-"`
}

// threatEventTagEditResponseMotiveAnnotatedJSON contains the JSON metadata for the
// struct [ThreatEventTagEditResponseMotiveAnnotated]
type threatEventTagEditResponseMotiveAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagEditResponseMotiveAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponseMotiveAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponseMotiveAnnotatedTLP string

const (
	ThreatEventTagEditResponseMotiveAnnotatedTLPRed         ThreatEventTagEditResponseMotiveAnnotatedTLP = "red"
	ThreatEventTagEditResponseMotiveAnnotatedTLPAmber       ThreatEventTagEditResponseMotiveAnnotatedTLP = "amber"
	ThreatEventTagEditResponseMotiveAnnotatedTLPAmberStrict ThreatEventTagEditResponseMotiveAnnotatedTLP = "amber+strict"
	ThreatEventTagEditResponseMotiveAnnotatedTLPGreen       ThreatEventTagEditResponseMotiveAnnotatedTLP = "green"
	ThreatEventTagEditResponseMotiveAnnotatedTLPClear       ThreatEventTagEditResponseMotiveAnnotatedTLP = "clear"
	ThreatEventTagEditResponseMotiveAnnotatedTLPPurple      ThreatEventTagEditResponseMotiveAnnotatedTLP = "purple"
)

func (r ThreatEventTagEditResponseMotiveAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseMotiveAnnotatedTLPRed, ThreatEventTagEditResponseMotiveAnnotatedTLPAmber, ThreatEventTagEditResponseMotiveAnnotatedTLPAmberStrict, ThreatEventTagEditResponseMotiveAnnotatedTLPGreen, ThreatEventTagEditResponseMotiveAnnotatedTLPClear, ThreatEventTagEditResponseMotiveAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagEditResponseOpsecLevelAnnotated struct {
	Value      string                                            `json:"value" api:"required"`
	Confidence float64                                           `json:"confidence"`
	TLP        ThreatEventTagEditResponseOpsecLevelAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagEditResponseOpsecLevelAnnotatedJSON `json:"-"`
}

// threatEventTagEditResponseOpsecLevelAnnotatedJSON contains the JSON metadata for
// the struct [ThreatEventTagEditResponseOpsecLevelAnnotated]
type threatEventTagEditResponseOpsecLevelAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagEditResponseOpsecLevelAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponseOpsecLevelAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponseOpsecLevelAnnotatedTLP string

const (
	ThreatEventTagEditResponseOpsecLevelAnnotatedTLPRed         ThreatEventTagEditResponseOpsecLevelAnnotatedTLP = "red"
	ThreatEventTagEditResponseOpsecLevelAnnotatedTLPAmber       ThreatEventTagEditResponseOpsecLevelAnnotatedTLP = "amber"
	ThreatEventTagEditResponseOpsecLevelAnnotatedTLPAmberStrict ThreatEventTagEditResponseOpsecLevelAnnotatedTLP = "amber+strict"
	ThreatEventTagEditResponseOpsecLevelAnnotatedTLPGreen       ThreatEventTagEditResponseOpsecLevelAnnotatedTLP = "green"
	ThreatEventTagEditResponseOpsecLevelAnnotatedTLPClear       ThreatEventTagEditResponseOpsecLevelAnnotatedTLP = "clear"
	ThreatEventTagEditResponseOpsecLevelAnnotatedTLPPurple      ThreatEventTagEditResponseOpsecLevelAnnotatedTLP = "purple"
)

func (r ThreatEventTagEditResponseOpsecLevelAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseOpsecLevelAnnotatedTLPRed, ThreatEventTagEditResponseOpsecLevelAnnotatedTLPAmber, ThreatEventTagEditResponseOpsecLevelAnnotatedTLPAmberStrict, ThreatEventTagEditResponseOpsecLevelAnnotatedTLPGreen, ThreatEventTagEditResponseOpsecLevelAnnotatedTLPClear, ThreatEventTagEditResponseOpsecLevelAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagEditResponseOriginCountryISOAnnotated struct {
	Value      string                                                  `json:"value" api:"required,nullable"`
	Confidence float64                                                 `json:"confidence"`
	TLP        ThreatEventTagEditResponseOriginCountryISOAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagEditResponseOriginCountryISOAnnotatedJSON `json:"-"`
}

// threatEventTagEditResponseOriginCountryISOAnnotatedJSON contains the JSON
// metadata for the struct [ThreatEventTagEditResponseOriginCountryISOAnnotated]
type threatEventTagEditResponseOriginCountryISOAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagEditResponseOriginCountryISOAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponseOriginCountryISOAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponseOriginCountryISOAnnotatedTLP string

const (
	ThreatEventTagEditResponseOriginCountryISOAnnotatedTLPRed         ThreatEventTagEditResponseOriginCountryISOAnnotatedTLP = "red"
	ThreatEventTagEditResponseOriginCountryISOAnnotatedTLPAmber       ThreatEventTagEditResponseOriginCountryISOAnnotatedTLP = "amber"
	ThreatEventTagEditResponseOriginCountryISOAnnotatedTLPAmberStrict ThreatEventTagEditResponseOriginCountryISOAnnotatedTLP = "amber+strict"
	ThreatEventTagEditResponseOriginCountryISOAnnotatedTLPGreen       ThreatEventTagEditResponseOriginCountryISOAnnotatedTLP = "green"
	ThreatEventTagEditResponseOriginCountryISOAnnotatedTLPClear       ThreatEventTagEditResponseOriginCountryISOAnnotatedTLP = "clear"
	ThreatEventTagEditResponseOriginCountryISOAnnotatedTLPPurple      ThreatEventTagEditResponseOriginCountryISOAnnotatedTLP = "purple"
)

func (r ThreatEventTagEditResponseOriginCountryISOAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseOriginCountryISOAnnotatedTLPRed, ThreatEventTagEditResponseOriginCountryISOAnnotatedTLPAmber, ThreatEventTagEditResponseOriginCountryISOAnnotatedTLPAmberStrict, ThreatEventTagEditResponseOriginCountryISOAnnotatedTLPGreen, ThreatEventTagEditResponseOriginCountryISOAnnotatedTLPClear, ThreatEventTagEditResponseOriginCountryISOAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagEditResponsePriorityAnnotated struct {
	Value float64                                         `json:"value" api:"required"`
	TLP   ThreatEventTagEditResponsePriorityAnnotatedTLP  `json:"tlp"`
	JSON  threatEventTagEditResponsePriorityAnnotatedJSON `json:"-"`
}

// threatEventTagEditResponsePriorityAnnotatedJSON contains the JSON metadata for
// the struct [ThreatEventTagEditResponsePriorityAnnotated]
type threatEventTagEditResponsePriorityAnnotatedJSON struct {
	Value       apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagEditResponsePriorityAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponsePriorityAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponsePriorityAnnotatedTLP string

const (
	ThreatEventTagEditResponsePriorityAnnotatedTLPRed         ThreatEventTagEditResponsePriorityAnnotatedTLP = "red"
	ThreatEventTagEditResponsePriorityAnnotatedTLPAmber       ThreatEventTagEditResponsePriorityAnnotatedTLP = "amber"
	ThreatEventTagEditResponsePriorityAnnotatedTLPAmberStrict ThreatEventTagEditResponsePriorityAnnotatedTLP = "amber+strict"
	ThreatEventTagEditResponsePriorityAnnotatedTLPGreen       ThreatEventTagEditResponsePriorityAnnotatedTLP = "green"
	ThreatEventTagEditResponsePriorityAnnotatedTLPClear       ThreatEventTagEditResponsePriorityAnnotatedTLP = "clear"
	ThreatEventTagEditResponsePriorityAnnotatedTLPPurple      ThreatEventTagEditResponsePriorityAnnotatedTLP = "purple"
)

func (r ThreatEventTagEditResponsePriorityAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponsePriorityAnnotatedTLPRed, ThreatEventTagEditResponsePriorityAnnotatedTLPAmber, ThreatEventTagEditResponsePriorityAnnotatedTLPAmberStrict, ThreatEventTagEditResponsePriorityAnnotatedTLPGreen, ThreatEventTagEditResponsePriorityAnnotatedTLPClear, ThreatEventTagEditResponsePriorityAnnotatedTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagEditResponseSophisticationLevelAnnotated struct {
	Value      string                                                     `json:"value" api:"required"`
	Confidence float64                                                    `json:"confidence"`
	TLP        ThreatEventTagEditResponseSophisticationLevelAnnotatedTLP  `json:"tlp"`
	JSON       threatEventTagEditResponseSophisticationLevelAnnotatedJSON `json:"-"`
}

// threatEventTagEditResponseSophisticationLevelAnnotatedJSON contains the JSON
// metadata for the struct [ThreatEventTagEditResponseSophisticationLevelAnnotated]
type threatEventTagEditResponseSophisticationLevelAnnotatedJSON struct {
	Value       apijson.Field
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagEditResponseSophisticationLevelAnnotated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponseSophisticationLevelAnnotatedJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagEditResponseSophisticationLevelAnnotatedTLP string

const (
	ThreatEventTagEditResponseSophisticationLevelAnnotatedTLPRed         ThreatEventTagEditResponseSophisticationLevelAnnotatedTLP = "red"
	ThreatEventTagEditResponseSophisticationLevelAnnotatedTLPAmber       ThreatEventTagEditResponseSophisticationLevelAnnotatedTLP = "amber"
	ThreatEventTagEditResponseSophisticationLevelAnnotatedTLPAmberStrict ThreatEventTagEditResponseSophisticationLevelAnnotatedTLP = "amber+strict"
	ThreatEventTagEditResponseSophisticationLevelAnnotatedTLPGreen       ThreatEventTagEditResponseSophisticationLevelAnnotatedTLP = "green"
	ThreatEventTagEditResponseSophisticationLevelAnnotatedTLPClear       ThreatEventTagEditResponseSophisticationLevelAnnotatedTLP = "clear"
	ThreatEventTagEditResponseSophisticationLevelAnnotatedTLPPurple      ThreatEventTagEditResponseSophisticationLevelAnnotatedTLP = "purple"
)

func (r ThreatEventTagEditResponseSophisticationLevelAnnotatedTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseSophisticationLevelAnnotatedTLPRed, ThreatEventTagEditResponseSophisticationLevelAnnotatedTLPAmber, ThreatEventTagEditResponseSophisticationLevelAnnotatedTLPAmberStrict, ThreatEventTagEditResponseSophisticationLevelAnnotatedTLPGreen, ThreatEventTagEditResponseSophisticationLevelAnnotatedTLPClear, ThreatEventTagEditResponseSophisticationLevelAnnotatedTLPPurple:
		return true
	}
	return false
}

// Tag-level TLP handling marking.
type ThreatEventTagEditResponseTLP string

const (
	ThreatEventTagEditResponseTLPRed         ThreatEventTagEditResponseTLP = "red"
	ThreatEventTagEditResponseTLPAmber       ThreatEventTagEditResponseTLP = "amber"
	ThreatEventTagEditResponseTLPAmberStrict ThreatEventTagEditResponseTLP = "amber+strict"
	ThreatEventTagEditResponseTLPGreen       ThreatEventTagEditResponseTLP = "green"
	ThreatEventTagEditResponseTLPClear       ThreatEventTagEditResponseTLP = "clear"
	ThreatEventTagEditResponseTLPPurple      ThreatEventTagEditResponseTLP = "purple"
)

func (r ThreatEventTagEditResponseTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseTLPRed, ThreatEventTagEditResponseTLPAmber, ThreatEventTagEditResponseTLPAmberStrict, ThreatEventTagEditResponseTLPGreen, ThreatEventTagEditResponseTLPClear, ThreatEventTagEditResponseTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagNewParams struct {
	// Account ID.
	AccountID      param.Field[string]                                     `path:"account_id" api:"required"`
	Value          param.Field[string]                                     `json:"value" api:"required"`
	ActiveDuration param.Field[ThreatEventTagNewParamsActiveDurationUnion] `json:"activeDuration"`
	ActorCategory  param.Field[ThreatEventTagNewParamsActorCategoryUnion]  `json:"actorCategory"`
	// Structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: stripped from
	// responses to non-CFONE accounts.
	Aliases                 param.Field[[]ThreatEventTagNewParamsAlias]                      `json:"aliases"`
	AliasGroupNames         param.Field[[]string]                                            `json:"aliasGroupNames"`
	AliasGroupNamesInternal param.Field[[]string]                                            `json:"aliasGroupNamesInternal"`
	AttributionOrganization param.Field[ThreatEventTagNewParamsAttributionOrganizationUnion] `json:"attributionOrganization"`
	// Tag type (category) UUID. Optional — when present, `properties` is validated
	// against this category's schema. When absent, the tag is typeless and properties
	// are accepted free-form.
	CategoryUUID param.Field[string] `json:"categoryUuid"`
	// Overall tag confidence (1-10). Optional.
	Confidence param.Field[int64] `json:"confidence"`
	// Date of discovery (ISO YYYY-MM-DD). Optional.
	DateOfDiscovery        param.Field[string]   `json:"dateOfDiscovery"`
	Description            param.Field[string]   `json:"description"`
	ExternalReferenceLinks param.Field[[]string] `json:"externalReferenceLinks"`
	// Structured external references ({ url, description }). Public: returned to all
	// accounts.
	ExternalReferences param.Field[[]ThreatEventTagNewParamsExternalReference] `json:"externalReferences"`
	// Internal structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: never
	// returned to non-CFONE accounts.
	InternalAliases     param.Field[[]ThreatEventTagNewParamsInternalAlias]       `json:"internalAliases"`
	InternalDescription param.Field[string]                                       `json:"internalDescription"`
	LastSeen            param.Field[string]                                       `json:"lastSeen"`
	Motive              param.Field[ThreatEventTagNewParamsMotiveUnion]           `json:"motive"`
	OpsecLevel          param.Field[ThreatEventTagNewParamsOpsecLevelUnion]       `json:"opsecLevel"`
	OriginCountryISO    param.Field[ThreatEventTagNewParamsOriginCountryISOUnion] `json:"originCountryISO"`
	Priority            param.Field[ThreatEventTagNewParamsPriorityUnion]         `json:"priority"`
	// Structured metadata blob. Optional. When `categoryUuid` is given, validated
	// against this category's schema on write. When typeless, accepted free-form. Use
	// `{}` for a tag with no custom data.
	Properties          param.Field[map[string]interface{}]                          `json:"properties"`
	SophisticationLevel param.Field[ThreatEventTagNewParamsSophisticationLevelUnion] `json:"sophisticationLevel"`
	// Tag-level TLP handling marking. Optional. Allowed values: red, amber,
	// amber+strict, green, clear, purple.
	TLP param.Field[ThreatEventTagNewParamsTLP] `json:"tlp"`
}

func (r ThreatEventTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagNewParamsActiveDurationObject].
type ThreatEventTagNewParamsActiveDurationUnion interface {
	ImplementsThreatEventTagNewParamsActiveDurationUnion()
}

type ThreatEventTagNewParamsActiveDurationObject struct {
	Value      param.Field[string]                                         `json:"value" api:"required"`
	Confidence param.Field[int64]                                          `json:"confidence"`
	TLP        param.Field[ThreatEventTagNewParamsActiveDurationObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagNewParamsActiveDurationObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagNewParamsActiveDurationObject) ImplementsThreatEventTagNewParamsActiveDurationUnion() {
}

type ThreatEventTagNewParamsActiveDurationObjectTLP string

const (
	ThreatEventTagNewParamsActiveDurationObjectTLPRed         ThreatEventTagNewParamsActiveDurationObjectTLP = "red"
	ThreatEventTagNewParamsActiveDurationObjectTLPAmber       ThreatEventTagNewParamsActiveDurationObjectTLP = "amber"
	ThreatEventTagNewParamsActiveDurationObjectTLPAmberStrict ThreatEventTagNewParamsActiveDurationObjectTLP = "amber+strict"
	ThreatEventTagNewParamsActiveDurationObjectTLPGreen       ThreatEventTagNewParamsActiveDurationObjectTLP = "green"
	ThreatEventTagNewParamsActiveDurationObjectTLPClear       ThreatEventTagNewParamsActiveDurationObjectTLP = "clear"
	ThreatEventTagNewParamsActiveDurationObjectTLPPurple      ThreatEventTagNewParamsActiveDurationObjectTLP = "purple"
)

func (r ThreatEventTagNewParamsActiveDurationObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsActiveDurationObjectTLPRed, ThreatEventTagNewParamsActiveDurationObjectTLPAmber, ThreatEventTagNewParamsActiveDurationObjectTLPAmberStrict, ThreatEventTagNewParamsActiveDurationObjectTLPGreen, ThreatEventTagNewParamsActiveDurationObjectTLPClear, ThreatEventTagNewParamsActiveDurationObjectTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagNewParamsActorCategoryObject].
type ThreatEventTagNewParamsActorCategoryUnion interface {
	ImplementsThreatEventTagNewParamsActorCategoryUnion()
}

type ThreatEventTagNewParamsActorCategoryObject struct {
	Value      param.Field[string]                                        `json:"value" api:"required"`
	Confidence param.Field[int64]                                         `json:"confidence"`
	TLP        param.Field[ThreatEventTagNewParamsActorCategoryObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagNewParamsActorCategoryObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagNewParamsActorCategoryObject) ImplementsThreatEventTagNewParamsActorCategoryUnion() {
}

type ThreatEventTagNewParamsActorCategoryObjectTLP string

const (
	ThreatEventTagNewParamsActorCategoryObjectTLPRed         ThreatEventTagNewParamsActorCategoryObjectTLP = "red"
	ThreatEventTagNewParamsActorCategoryObjectTLPAmber       ThreatEventTagNewParamsActorCategoryObjectTLP = "amber"
	ThreatEventTagNewParamsActorCategoryObjectTLPAmberStrict ThreatEventTagNewParamsActorCategoryObjectTLP = "amber+strict"
	ThreatEventTagNewParamsActorCategoryObjectTLPGreen       ThreatEventTagNewParamsActorCategoryObjectTLP = "green"
	ThreatEventTagNewParamsActorCategoryObjectTLPClear       ThreatEventTagNewParamsActorCategoryObjectTLP = "clear"
	ThreatEventTagNewParamsActorCategoryObjectTLPPurple      ThreatEventTagNewParamsActorCategoryObjectTLP = "purple"
)

func (r ThreatEventTagNewParamsActorCategoryObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsActorCategoryObjectTLPRed, ThreatEventTagNewParamsActorCategoryObjectTLPAmber, ThreatEventTagNewParamsActorCategoryObjectTLPAmberStrict, ThreatEventTagNewParamsActorCategoryObjectTLPGreen, ThreatEventTagNewParamsActorCategoryObjectTLPClear, ThreatEventTagNewParamsActorCategoryObjectTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagNewParamsAlias struct {
	Value      param.Field[string]                            `json:"value" api:"required"`
	Confidence param.Field[int64]                             `json:"confidence"`
	TLP        param.Field[ThreatEventTagNewParamsAliasesTLP] `json:"tlp"`
}

func (r ThreatEventTagNewParamsAlias) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagNewParamsAliasesTLP string

const (
	ThreatEventTagNewParamsAliasesTLPRed         ThreatEventTagNewParamsAliasesTLP = "red"
	ThreatEventTagNewParamsAliasesTLPAmber       ThreatEventTagNewParamsAliasesTLP = "amber"
	ThreatEventTagNewParamsAliasesTLPAmberStrict ThreatEventTagNewParamsAliasesTLP = "amber+strict"
	ThreatEventTagNewParamsAliasesTLPGreen       ThreatEventTagNewParamsAliasesTLP = "green"
	ThreatEventTagNewParamsAliasesTLPClear       ThreatEventTagNewParamsAliasesTLP = "clear"
	ThreatEventTagNewParamsAliasesTLPPurple      ThreatEventTagNewParamsAliasesTLP = "purple"
)

func (r ThreatEventTagNewParamsAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsAliasesTLPRed, ThreatEventTagNewParamsAliasesTLPAmber, ThreatEventTagNewParamsAliasesTLPAmberStrict, ThreatEventTagNewParamsAliasesTLPGreen, ThreatEventTagNewParamsAliasesTLPClear, ThreatEventTagNewParamsAliasesTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagNewParamsAttributionOrganizationObject].
type ThreatEventTagNewParamsAttributionOrganizationUnion interface {
	ImplementsThreatEventTagNewParamsAttributionOrganizationUnion()
}

type ThreatEventTagNewParamsAttributionOrganizationObject struct {
	Value      param.Field[string]                                                  `json:"value" api:"required"`
	Confidence param.Field[int64]                                                   `json:"confidence"`
	TLP        param.Field[ThreatEventTagNewParamsAttributionOrganizationObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagNewParamsAttributionOrganizationObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagNewParamsAttributionOrganizationObject) ImplementsThreatEventTagNewParamsAttributionOrganizationUnion() {
}

type ThreatEventTagNewParamsAttributionOrganizationObjectTLP string

const (
	ThreatEventTagNewParamsAttributionOrganizationObjectTLPRed         ThreatEventTagNewParamsAttributionOrganizationObjectTLP = "red"
	ThreatEventTagNewParamsAttributionOrganizationObjectTLPAmber       ThreatEventTagNewParamsAttributionOrganizationObjectTLP = "amber"
	ThreatEventTagNewParamsAttributionOrganizationObjectTLPAmberStrict ThreatEventTagNewParamsAttributionOrganizationObjectTLP = "amber+strict"
	ThreatEventTagNewParamsAttributionOrganizationObjectTLPGreen       ThreatEventTagNewParamsAttributionOrganizationObjectTLP = "green"
	ThreatEventTagNewParamsAttributionOrganizationObjectTLPClear       ThreatEventTagNewParamsAttributionOrganizationObjectTLP = "clear"
	ThreatEventTagNewParamsAttributionOrganizationObjectTLPPurple      ThreatEventTagNewParamsAttributionOrganizationObjectTLP = "purple"
)

func (r ThreatEventTagNewParamsAttributionOrganizationObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsAttributionOrganizationObjectTLPRed, ThreatEventTagNewParamsAttributionOrganizationObjectTLPAmber, ThreatEventTagNewParamsAttributionOrganizationObjectTLPAmberStrict, ThreatEventTagNewParamsAttributionOrganizationObjectTLPGreen, ThreatEventTagNewParamsAttributionOrganizationObjectTLPClear, ThreatEventTagNewParamsAttributionOrganizationObjectTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagNewParamsExternalReference struct {
	URL         param.Field[string] `json:"url" api:"required"`
	Description param.Field[string] `json:"description"`
}

func (r ThreatEventTagNewParamsExternalReference) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagNewParamsInternalAlias struct {
	Value      param.Field[string]                                    `json:"value" api:"required"`
	Confidence param.Field[int64]                                     `json:"confidence"`
	TLP        param.Field[ThreatEventTagNewParamsInternalAliasesTLP] `json:"tlp"`
}

func (r ThreatEventTagNewParamsInternalAlias) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagNewParamsInternalAliasesTLP string

const (
	ThreatEventTagNewParamsInternalAliasesTLPRed         ThreatEventTagNewParamsInternalAliasesTLP = "red"
	ThreatEventTagNewParamsInternalAliasesTLPAmber       ThreatEventTagNewParamsInternalAliasesTLP = "amber"
	ThreatEventTagNewParamsInternalAliasesTLPAmberStrict ThreatEventTagNewParamsInternalAliasesTLP = "amber+strict"
	ThreatEventTagNewParamsInternalAliasesTLPGreen       ThreatEventTagNewParamsInternalAliasesTLP = "green"
	ThreatEventTagNewParamsInternalAliasesTLPClear       ThreatEventTagNewParamsInternalAliasesTLP = "clear"
	ThreatEventTagNewParamsInternalAliasesTLPPurple      ThreatEventTagNewParamsInternalAliasesTLP = "purple"
)

func (r ThreatEventTagNewParamsInternalAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsInternalAliasesTLPRed, ThreatEventTagNewParamsInternalAliasesTLPAmber, ThreatEventTagNewParamsInternalAliasesTLPAmberStrict, ThreatEventTagNewParamsInternalAliasesTLPGreen, ThreatEventTagNewParamsInternalAliasesTLPClear, ThreatEventTagNewParamsInternalAliasesTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagNewParamsMotiveObject].
type ThreatEventTagNewParamsMotiveUnion interface {
	ImplementsThreatEventTagNewParamsMotiveUnion()
}

type ThreatEventTagNewParamsMotiveObject struct {
	Value      param.Field[string]                                 `json:"value" api:"required"`
	Confidence param.Field[int64]                                  `json:"confidence"`
	TLP        param.Field[ThreatEventTagNewParamsMotiveObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagNewParamsMotiveObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagNewParamsMotiveObject) ImplementsThreatEventTagNewParamsMotiveUnion() {}

type ThreatEventTagNewParamsMotiveObjectTLP string

const (
	ThreatEventTagNewParamsMotiveObjectTLPRed         ThreatEventTagNewParamsMotiveObjectTLP = "red"
	ThreatEventTagNewParamsMotiveObjectTLPAmber       ThreatEventTagNewParamsMotiveObjectTLP = "amber"
	ThreatEventTagNewParamsMotiveObjectTLPAmberStrict ThreatEventTagNewParamsMotiveObjectTLP = "amber+strict"
	ThreatEventTagNewParamsMotiveObjectTLPGreen       ThreatEventTagNewParamsMotiveObjectTLP = "green"
	ThreatEventTagNewParamsMotiveObjectTLPClear       ThreatEventTagNewParamsMotiveObjectTLP = "clear"
	ThreatEventTagNewParamsMotiveObjectTLPPurple      ThreatEventTagNewParamsMotiveObjectTLP = "purple"
)

func (r ThreatEventTagNewParamsMotiveObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsMotiveObjectTLPRed, ThreatEventTagNewParamsMotiveObjectTLPAmber, ThreatEventTagNewParamsMotiveObjectTLPAmberStrict, ThreatEventTagNewParamsMotiveObjectTLPGreen, ThreatEventTagNewParamsMotiveObjectTLPClear, ThreatEventTagNewParamsMotiveObjectTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagNewParamsOpsecLevelObject].
type ThreatEventTagNewParamsOpsecLevelUnion interface {
	ImplementsThreatEventTagNewParamsOpsecLevelUnion()
}

type ThreatEventTagNewParamsOpsecLevelObject struct {
	Value      param.Field[string]                                     `json:"value" api:"required"`
	Confidence param.Field[int64]                                      `json:"confidence"`
	TLP        param.Field[ThreatEventTagNewParamsOpsecLevelObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagNewParamsOpsecLevelObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagNewParamsOpsecLevelObject) ImplementsThreatEventTagNewParamsOpsecLevelUnion() {}

type ThreatEventTagNewParamsOpsecLevelObjectTLP string

const (
	ThreatEventTagNewParamsOpsecLevelObjectTLPRed         ThreatEventTagNewParamsOpsecLevelObjectTLP = "red"
	ThreatEventTagNewParamsOpsecLevelObjectTLPAmber       ThreatEventTagNewParamsOpsecLevelObjectTLP = "amber"
	ThreatEventTagNewParamsOpsecLevelObjectTLPAmberStrict ThreatEventTagNewParamsOpsecLevelObjectTLP = "amber+strict"
	ThreatEventTagNewParamsOpsecLevelObjectTLPGreen       ThreatEventTagNewParamsOpsecLevelObjectTLP = "green"
	ThreatEventTagNewParamsOpsecLevelObjectTLPClear       ThreatEventTagNewParamsOpsecLevelObjectTLP = "clear"
	ThreatEventTagNewParamsOpsecLevelObjectTLPPurple      ThreatEventTagNewParamsOpsecLevelObjectTLP = "purple"
)

func (r ThreatEventTagNewParamsOpsecLevelObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsOpsecLevelObjectTLPRed, ThreatEventTagNewParamsOpsecLevelObjectTLPAmber, ThreatEventTagNewParamsOpsecLevelObjectTLPAmberStrict, ThreatEventTagNewParamsOpsecLevelObjectTLPGreen, ThreatEventTagNewParamsOpsecLevelObjectTLPClear, ThreatEventTagNewParamsOpsecLevelObjectTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagNewParamsOriginCountryISOObject].
type ThreatEventTagNewParamsOriginCountryISOUnion interface {
	ImplementsThreatEventTagNewParamsOriginCountryISOUnion()
}

type ThreatEventTagNewParamsOriginCountryISOObject struct {
	Value      param.Field[string]                                           `json:"value" api:"required"`
	Confidence param.Field[int64]                                            `json:"confidence"`
	TLP        param.Field[ThreatEventTagNewParamsOriginCountryISOObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagNewParamsOriginCountryISOObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagNewParamsOriginCountryISOObject) ImplementsThreatEventTagNewParamsOriginCountryISOUnion() {
}

type ThreatEventTagNewParamsOriginCountryISOObjectTLP string

const (
	ThreatEventTagNewParamsOriginCountryISOObjectTLPRed         ThreatEventTagNewParamsOriginCountryISOObjectTLP = "red"
	ThreatEventTagNewParamsOriginCountryISOObjectTLPAmber       ThreatEventTagNewParamsOriginCountryISOObjectTLP = "amber"
	ThreatEventTagNewParamsOriginCountryISOObjectTLPAmberStrict ThreatEventTagNewParamsOriginCountryISOObjectTLP = "amber+strict"
	ThreatEventTagNewParamsOriginCountryISOObjectTLPGreen       ThreatEventTagNewParamsOriginCountryISOObjectTLP = "green"
	ThreatEventTagNewParamsOriginCountryISOObjectTLPClear       ThreatEventTagNewParamsOriginCountryISOObjectTLP = "clear"
	ThreatEventTagNewParamsOriginCountryISOObjectTLPPurple      ThreatEventTagNewParamsOriginCountryISOObjectTLP = "purple"
)

func (r ThreatEventTagNewParamsOriginCountryISOObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsOriginCountryISOObjectTLPRed, ThreatEventTagNewParamsOriginCountryISOObjectTLPAmber, ThreatEventTagNewParamsOriginCountryISOObjectTLPAmberStrict, ThreatEventTagNewParamsOriginCountryISOObjectTLPGreen, ThreatEventTagNewParamsOriginCountryISOObjectTLPClear, ThreatEventTagNewParamsOriginCountryISOObjectTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionFloat],
// [cloudforce_one.ThreatEventTagNewParamsPriorityObject].
type ThreatEventTagNewParamsPriorityUnion interface {
	ImplementsThreatEventTagNewParamsPriorityUnion()
}

type ThreatEventTagNewParamsPriorityObject struct {
	Value      param.Field[float64]                                  `json:"value" api:"required"`
	Confidence param.Field[int64]                                    `json:"confidence"`
	TLP        param.Field[ThreatEventTagNewParamsPriorityObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagNewParamsPriorityObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagNewParamsPriorityObject) ImplementsThreatEventTagNewParamsPriorityUnion() {}

type ThreatEventTagNewParamsPriorityObjectTLP string

const (
	ThreatEventTagNewParamsPriorityObjectTLPRed         ThreatEventTagNewParamsPriorityObjectTLP = "red"
	ThreatEventTagNewParamsPriorityObjectTLPAmber       ThreatEventTagNewParamsPriorityObjectTLP = "amber"
	ThreatEventTagNewParamsPriorityObjectTLPAmberStrict ThreatEventTagNewParamsPriorityObjectTLP = "amber+strict"
	ThreatEventTagNewParamsPriorityObjectTLPGreen       ThreatEventTagNewParamsPriorityObjectTLP = "green"
	ThreatEventTagNewParamsPriorityObjectTLPClear       ThreatEventTagNewParamsPriorityObjectTLP = "clear"
	ThreatEventTagNewParamsPriorityObjectTLPPurple      ThreatEventTagNewParamsPriorityObjectTLP = "purple"
)

func (r ThreatEventTagNewParamsPriorityObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsPriorityObjectTLPRed, ThreatEventTagNewParamsPriorityObjectTLPAmber, ThreatEventTagNewParamsPriorityObjectTLPAmberStrict, ThreatEventTagNewParamsPriorityObjectTLPGreen, ThreatEventTagNewParamsPriorityObjectTLPClear, ThreatEventTagNewParamsPriorityObjectTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagNewParamsSophisticationLevelObject].
type ThreatEventTagNewParamsSophisticationLevelUnion interface {
	ImplementsThreatEventTagNewParamsSophisticationLevelUnion()
}

type ThreatEventTagNewParamsSophisticationLevelObject struct {
	Value      param.Field[string]                                              `json:"value" api:"required"`
	Confidence param.Field[int64]                                               `json:"confidence"`
	TLP        param.Field[ThreatEventTagNewParamsSophisticationLevelObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagNewParamsSophisticationLevelObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagNewParamsSophisticationLevelObject) ImplementsThreatEventTagNewParamsSophisticationLevelUnion() {
}

type ThreatEventTagNewParamsSophisticationLevelObjectTLP string

const (
	ThreatEventTagNewParamsSophisticationLevelObjectTLPRed         ThreatEventTagNewParamsSophisticationLevelObjectTLP = "red"
	ThreatEventTagNewParamsSophisticationLevelObjectTLPAmber       ThreatEventTagNewParamsSophisticationLevelObjectTLP = "amber"
	ThreatEventTagNewParamsSophisticationLevelObjectTLPAmberStrict ThreatEventTagNewParamsSophisticationLevelObjectTLP = "amber+strict"
	ThreatEventTagNewParamsSophisticationLevelObjectTLPGreen       ThreatEventTagNewParamsSophisticationLevelObjectTLP = "green"
	ThreatEventTagNewParamsSophisticationLevelObjectTLPClear       ThreatEventTagNewParamsSophisticationLevelObjectTLP = "clear"
	ThreatEventTagNewParamsSophisticationLevelObjectTLPPurple      ThreatEventTagNewParamsSophisticationLevelObjectTLP = "purple"
)

func (r ThreatEventTagNewParamsSophisticationLevelObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsSophisticationLevelObjectTLPRed, ThreatEventTagNewParamsSophisticationLevelObjectTLPAmber, ThreatEventTagNewParamsSophisticationLevelObjectTLPAmberStrict, ThreatEventTagNewParamsSophisticationLevelObjectTLPGreen, ThreatEventTagNewParamsSophisticationLevelObjectTLPClear, ThreatEventTagNewParamsSophisticationLevelObjectTLPPurple:
		return true
	}
	return false
}

// Tag-level TLP handling marking. Optional. Allowed values: red, amber,
// amber+strict, green, clear, purple.
type ThreatEventTagNewParamsTLP string

const (
	ThreatEventTagNewParamsTLPRed         ThreatEventTagNewParamsTLP = "red"
	ThreatEventTagNewParamsTLPAmber       ThreatEventTagNewParamsTLP = "amber"
	ThreatEventTagNewParamsTLPAmberStrict ThreatEventTagNewParamsTLP = "amber+strict"
	ThreatEventTagNewParamsTLPGreen       ThreatEventTagNewParamsTLP = "green"
	ThreatEventTagNewParamsTLPClear       ThreatEventTagNewParamsTLP = "clear"
	ThreatEventTagNewParamsTLPPurple      ThreatEventTagNewParamsTLP = "purple"
)

func (r ThreatEventTagNewParamsTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsTLPRed, ThreatEventTagNewParamsTLPAmber, ThreatEventTagNewParamsTLPAmberStrict, ThreatEventTagNewParamsTLPGreen, ThreatEventTagNewParamsTLPClear, ThreatEventTagNewParamsTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Cache strategy. 'from-graph' serves results from the graph-node KV cache when
	// all requested UUIDs are cached; falls back to normal path on partial/zero hit.
	Cache        param.Field[ThreatEventTagListParamsCache] `query:"cache"`
	CategoryUUID param.Field[string]                        `query:"categoryUuid"`
	// Structured filters as a JSON array of {field, op, value} objects. Searchable
	// fields: uuid, value, categoryName, description, dateOfDiscovery, tlp,
	// confidence, actorCategory, motive, attributionOrganization, originCountryISO,
	// aliases, externalReferences, opsecLevel, sophisticationLevel, activeDuration,
	// priority, lastSeen, aliasGroupNames. Operators: equals, not, contains,
	// startsWith, endsWith, gt, lt, gte, lte, like, in, find. Use 'in' for bulk OR
	// within a single field, e.g.
	// filters=[{"field":"originCountryISO","op":"in","value":["IR","CN"]}]. Multiple
	// entries are AND-joined. Max 10 entries per request, max 100 values per 'in'.
	// Per-field notes: `uuid` accepts only 'equals' and 'in' (other operators throw
	// ValidationError) — matched against the canonical lowercase storage but callers
	// may pass either case (the server lowercases before comparison); index-backed by
	// the column's UNIQUE constraint and intended for batched UUID → tag resolution.
	// `originCountryISO` uses its B-tree index for equals/not/in. `priority` uses its
	// B-tree index for numeric comparisons. Other string columns (`actorCategory`,
	// `motive`, etc.) are case-insensitive and unindexed; current catalog size makes
	// this a non-issue. `endsWith` and `aliasGroupNames` contains/like are
	// leading-wildcard scans and slow on large result sets. `aliasGroupNames` matches
	// on the JSON-encoded text, so substrings can cross alias boundaries (a search for
	// "apt28" will also match "apt280" if both appear in the same tag's alias list).
	Filters  param.Field[[]ThreatEventTagListParamsFilter] `query:"filters"`
	Page     param.Field[float64]                          `query:"page"`
	PageSize param.Field[float64]                          `query:"pageSize"`
	// Free-text substring match on tag value AND custom-field properties. Searches
	// case-insensitively inside both `Tag.value` and the serialized `Tag.properties`
	// JSON blob (keys, values, and annotation metadata like confidence/tlp are all
	// searchable). Same serialized-text tradeoff as `aliasGroupNames` — substrings can
	// cross JSON boundaries.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [ThreatEventTagListParams]'s query parameters as
// `url.Values`.
func (r ThreatEventTagListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Cache strategy. 'from-graph' serves results from the graph-node KV cache when
// all requested UUIDs are cached; falls back to normal path on partial/zero hit.
type ThreatEventTagListParamsCache string

const (
	ThreatEventTagListParamsCacheFromGraph ThreatEventTagListParamsCache = "from-graph"
)

func (r ThreatEventTagListParamsCache) IsKnown() bool {
	switch r {
	case ThreatEventTagListParamsCacheFromGraph:
		return true
	}
	return false
}

type ThreatEventTagListParamsFilter struct {
	// Tag field to search on. Allowed first-class fields: uuid, value, categoryName,
	// description, dateOfDiscovery, tlp, confidence, actorCategory, motive,
	// attributionOrganization, originCountryISO, aliases, externalReferences,
	// opsecLevel, sophisticationLevel, activeDuration, priority, lastSeen,
	// aliasGroupNames. Also supports properties.<key> to filter on custom field values
	// (matches both raw values and annotated {value,confidence,tlp} shapes via
	// COALESCE), and properties.<key>.tlp / properties.<key>.confidence to filter
	// directly on annotation sub-fields.
	Field param.Field[string] `query:"field" api:"required"`
	// Search operator. Use 'in' for bulk OR within a single field.
	Op param.Field[ThreatEventTagListParamsFiltersOp] `query:"op" api:"required"`
	// Search value. String or number for most operators. Array for 'in' (max 100
	// items).
	Value param.Field[ThreatEventTagListParamsFiltersValueUnion] `query:"value"`
}

// URLQuery serializes [ThreatEventTagListParamsFilter]'s query parameters as
// `url.Values`.
func (r ThreatEventTagListParamsFilter) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Search operator. Use 'in' for bulk OR within a single field.
type ThreatEventTagListParamsFiltersOp string

const (
	ThreatEventTagListParamsFiltersOpEquals     ThreatEventTagListParamsFiltersOp = "equals"
	ThreatEventTagListParamsFiltersOpNot        ThreatEventTagListParamsFiltersOp = "not"
	ThreatEventTagListParamsFiltersOpGt         ThreatEventTagListParamsFiltersOp = "gt"
	ThreatEventTagListParamsFiltersOpGte        ThreatEventTagListParamsFiltersOp = "gte"
	ThreatEventTagListParamsFiltersOpLt         ThreatEventTagListParamsFiltersOp = "lt"
	ThreatEventTagListParamsFiltersOpLte        ThreatEventTagListParamsFiltersOp = "lte"
	ThreatEventTagListParamsFiltersOpLike       ThreatEventTagListParamsFiltersOp = "like"
	ThreatEventTagListParamsFiltersOpContains   ThreatEventTagListParamsFiltersOp = "contains"
	ThreatEventTagListParamsFiltersOpStartsWith ThreatEventTagListParamsFiltersOp = "startsWith"
	ThreatEventTagListParamsFiltersOpEndsWith   ThreatEventTagListParamsFiltersOp = "endsWith"
	ThreatEventTagListParamsFiltersOpIn         ThreatEventTagListParamsFiltersOp = "in"
	ThreatEventTagListParamsFiltersOpFind       ThreatEventTagListParamsFiltersOp = "find"
)

func (r ThreatEventTagListParamsFiltersOp) IsKnown() bool {
	switch r {
	case ThreatEventTagListParamsFiltersOpEquals, ThreatEventTagListParamsFiltersOpNot, ThreatEventTagListParamsFiltersOpGt, ThreatEventTagListParamsFiltersOpGte, ThreatEventTagListParamsFiltersOpLt, ThreatEventTagListParamsFiltersOpLte, ThreatEventTagListParamsFiltersOpLike, ThreatEventTagListParamsFiltersOpContains, ThreatEventTagListParamsFiltersOpStartsWith, ThreatEventTagListParamsFiltersOpEndsWith, ThreatEventTagListParamsFiltersOpIn, ThreatEventTagListParamsFiltersOpFind:
		return true
	}
	return false
}

// Search value. String or number for most operators. Array for 'in' (max 100
// items).
//
// Satisfied by [shared.UnionString], [shared.UnionFloat],
// [cloudforce_one.ThreatEventTagListParamsFiltersValueArray].
type ThreatEventTagListParamsFiltersValueUnion interface {
	ImplementsThreatEventTagListParamsFiltersValueUnion()
}

type ThreatEventTagListParamsFiltersValueArray []ThreatEventTagListParamsFiltersValueArrayItemUnion

func (r ThreatEventTagListParamsFiltersValueArray) ImplementsThreatEventTagListParamsFiltersValueUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type ThreatEventTagListParamsFiltersValueArrayItemUnion interface {
	ImplementsThreatEventTagListParamsFiltersValueArrayItemUnion()
}

type ThreatEventTagDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatEventTagEditParams struct {
	// Account ID.
	AccountID      param.Field[string]                                      `path:"account_id" api:"required"`
	ActiveDuration param.Field[ThreatEventTagEditParamsActiveDurationUnion] `json:"activeDuration"`
	ActorCategory  param.Field[ThreatEventTagEditParamsActorCategoryUnion]  `json:"actorCategory"`
	// Structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: stripped from
	// responses to non-CFONE accounts.
	Aliases                 param.Field[[]ThreatEventTagEditParamsAlias]                      `json:"aliases"`
	AliasGroupNames         param.Field[[]string]                                             `json:"aliasGroupNames"`
	AliasGroupNamesInternal param.Field[[]string]                                             `json:"aliasGroupNamesInternal"`
	AttributionOrganization param.Field[ThreatEventTagEditParamsAttributionOrganizationUnion] `json:"attributionOrganization"`
	// Tag type (category) UUID. When changed, existing `properties` are re-validated
	// against the new category's schema (400 on mismatch). Set to null to unlink
	// (typeless; properties stop being validated).
	CategoryUUID param.Field[string] `json:"categoryUuid"`
	// Overall tag confidence (1-10). Omit to preserve existing.
	Confidence param.Field[int64] `json:"confidence"`
	// Date of discovery (ISO YYYY-MM-DD). Omit to preserve existing.
	DateOfDiscovery        param.Field[string]   `json:"dateOfDiscovery"`
	Description            param.Field[string]   `json:"description"`
	ExternalReferenceLinks param.Field[[]string] `json:"externalReferenceLinks"`
	// Structured external references ({ url, description }). Public: returned to all
	// accounts.
	ExternalReferences param.Field[[]ThreatEventTagEditParamsExternalReference] `json:"externalReferences"`
	// Internal structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: never
	// returned to non-CFONE accounts.
	InternalAliases     param.Field[[]ThreatEventTagEditParamsInternalAlias]       `json:"internalAliases"`
	InternalDescription param.Field[string]                                        `json:"internalDescription"`
	LastSeen            param.Field[string]                                        `json:"lastSeen"`
	Motive              param.Field[ThreatEventTagEditParamsMotiveUnion]           `json:"motive"`
	OpsecLevel          param.Field[ThreatEventTagEditParamsOpsecLevelUnion]       `json:"opsecLevel"`
	OriginCountryISO    param.Field[ThreatEventTagEditParamsOriginCountryISOUnion] `json:"originCountryISO"`
	Priority            param.Field[ThreatEventTagEditParamsPriorityUnion]         `json:"priority"`
	// Custom field values blob. When omitted, the existing value is preserved. When
	// provided, performs a shallow per-key merge over the stored value (unmentioned
	// keys are retained). Setting an individual key to null deletes that key.
	// Validation runs against the merged result, so a partial update may omit a
	// schema-required key if the stored value supplies it.
	Properties          param.Field[map[string]interface{}]                           `json:"properties"`
	SophisticationLevel param.Field[ThreatEventTagEditParamsSophisticationLevelUnion] `json:"sophisticationLevel"`
	// Tag-level TLP marking. Omit to preserve existing. Cannot be cleared to null.
	TLP   param.Field[ThreatEventTagEditParamsTLP] `json:"tlp"`
	Value param.Field[string]                      `json:"value"`
}

func (r ThreatEventTagEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagEditParamsActiveDurationObject].
type ThreatEventTagEditParamsActiveDurationUnion interface {
	ImplementsThreatEventTagEditParamsActiveDurationUnion()
}

type ThreatEventTagEditParamsActiveDurationObject struct {
	Value      param.Field[string]                                          `json:"value" api:"required"`
	Confidence param.Field[int64]                                           `json:"confidence"`
	TLP        param.Field[ThreatEventTagEditParamsActiveDurationObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagEditParamsActiveDurationObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagEditParamsActiveDurationObject) ImplementsThreatEventTagEditParamsActiveDurationUnion() {
}

type ThreatEventTagEditParamsActiveDurationObjectTLP string

const (
	ThreatEventTagEditParamsActiveDurationObjectTLPRed         ThreatEventTagEditParamsActiveDurationObjectTLP = "red"
	ThreatEventTagEditParamsActiveDurationObjectTLPAmber       ThreatEventTagEditParamsActiveDurationObjectTLP = "amber"
	ThreatEventTagEditParamsActiveDurationObjectTLPAmberStrict ThreatEventTagEditParamsActiveDurationObjectTLP = "amber+strict"
	ThreatEventTagEditParamsActiveDurationObjectTLPGreen       ThreatEventTagEditParamsActiveDurationObjectTLP = "green"
	ThreatEventTagEditParamsActiveDurationObjectTLPClear       ThreatEventTagEditParamsActiveDurationObjectTLP = "clear"
	ThreatEventTagEditParamsActiveDurationObjectTLPPurple      ThreatEventTagEditParamsActiveDurationObjectTLP = "purple"
)

func (r ThreatEventTagEditParamsActiveDurationObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsActiveDurationObjectTLPRed, ThreatEventTagEditParamsActiveDurationObjectTLPAmber, ThreatEventTagEditParamsActiveDurationObjectTLPAmberStrict, ThreatEventTagEditParamsActiveDurationObjectTLPGreen, ThreatEventTagEditParamsActiveDurationObjectTLPClear, ThreatEventTagEditParamsActiveDurationObjectTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagEditParamsActorCategoryObject].
type ThreatEventTagEditParamsActorCategoryUnion interface {
	ImplementsThreatEventTagEditParamsActorCategoryUnion()
}

type ThreatEventTagEditParamsActorCategoryObject struct {
	Value      param.Field[string]                                         `json:"value" api:"required"`
	Confidence param.Field[int64]                                          `json:"confidence"`
	TLP        param.Field[ThreatEventTagEditParamsActorCategoryObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagEditParamsActorCategoryObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagEditParamsActorCategoryObject) ImplementsThreatEventTagEditParamsActorCategoryUnion() {
}

type ThreatEventTagEditParamsActorCategoryObjectTLP string

const (
	ThreatEventTagEditParamsActorCategoryObjectTLPRed         ThreatEventTagEditParamsActorCategoryObjectTLP = "red"
	ThreatEventTagEditParamsActorCategoryObjectTLPAmber       ThreatEventTagEditParamsActorCategoryObjectTLP = "amber"
	ThreatEventTagEditParamsActorCategoryObjectTLPAmberStrict ThreatEventTagEditParamsActorCategoryObjectTLP = "amber+strict"
	ThreatEventTagEditParamsActorCategoryObjectTLPGreen       ThreatEventTagEditParamsActorCategoryObjectTLP = "green"
	ThreatEventTagEditParamsActorCategoryObjectTLPClear       ThreatEventTagEditParamsActorCategoryObjectTLP = "clear"
	ThreatEventTagEditParamsActorCategoryObjectTLPPurple      ThreatEventTagEditParamsActorCategoryObjectTLP = "purple"
)

func (r ThreatEventTagEditParamsActorCategoryObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsActorCategoryObjectTLPRed, ThreatEventTagEditParamsActorCategoryObjectTLPAmber, ThreatEventTagEditParamsActorCategoryObjectTLPAmberStrict, ThreatEventTagEditParamsActorCategoryObjectTLPGreen, ThreatEventTagEditParamsActorCategoryObjectTLPClear, ThreatEventTagEditParamsActorCategoryObjectTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagEditParamsAlias struct {
	Value      param.Field[string]                             `json:"value" api:"required"`
	Confidence param.Field[int64]                              `json:"confidence"`
	TLP        param.Field[ThreatEventTagEditParamsAliasesTLP] `json:"tlp"`
}

func (r ThreatEventTagEditParamsAlias) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagEditParamsAliasesTLP string

const (
	ThreatEventTagEditParamsAliasesTLPRed         ThreatEventTagEditParamsAliasesTLP = "red"
	ThreatEventTagEditParamsAliasesTLPAmber       ThreatEventTagEditParamsAliasesTLP = "amber"
	ThreatEventTagEditParamsAliasesTLPAmberStrict ThreatEventTagEditParamsAliasesTLP = "amber+strict"
	ThreatEventTagEditParamsAliasesTLPGreen       ThreatEventTagEditParamsAliasesTLP = "green"
	ThreatEventTagEditParamsAliasesTLPClear       ThreatEventTagEditParamsAliasesTLP = "clear"
	ThreatEventTagEditParamsAliasesTLPPurple      ThreatEventTagEditParamsAliasesTLP = "purple"
)

func (r ThreatEventTagEditParamsAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsAliasesTLPRed, ThreatEventTagEditParamsAliasesTLPAmber, ThreatEventTagEditParamsAliasesTLPAmberStrict, ThreatEventTagEditParamsAliasesTLPGreen, ThreatEventTagEditParamsAliasesTLPClear, ThreatEventTagEditParamsAliasesTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagEditParamsAttributionOrganizationObject].
type ThreatEventTagEditParamsAttributionOrganizationUnion interface {
	ImplementsThreatEventTagEditParamsAttributionOrganizationUnion()
}

type ThreatEventTagEditParamsAttributionOrganizationObject struct {
	Value      param.Field[string]                                                   `json:"value" api:"required"`
	Confidence param.Field[int64]                                                    `json:"confidence"`
	TLP        param.Field[ThreatEventTagEditParamsAttributionOrganizationObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagEditParamsAttributionOrganizationObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagEditParamsAttributionOrganizationObject) ImplementsThreatEventTagEditParamsAttributionOrganizationUnion() {
}

type ThreatEventTagEditParamsAttributionOrganizationObjectTLP string

const (
	ThreatEventTagEditParamsAttributionOrganizationObjectTLPRed         ThreatEventTagEditParamsAttributionOrganizationObjectTLP = "red"
	ThreatEventTagEditParamsAttributionOrganizationObjectTLPAmber       ThreatEventTagEditParamsAttributionOrganizationObjectTLP = "amber"
	ThreatEventTagEditParamsAttributionOrganizationObjectTLPAmberStrict ThreatEventTagEditParamsAttributionOrganizationObjectTLP = "amber+strict"
	ThreatEventTagEditParamsAttributionOrganizationObjectTLPGreen       ThreatEventTagEditParamsAttributionOrganizationObjectTLP = "green"
	ThreatEventTagEditParamsAttributionOrganizationObjectTLPClear       ThreatEventTagEditParamsAttributionOrganizationObjectTLP = "clear"
	ThreatEventTagEditParamsAttributionOrganizationObjectTLPPurple      ThreatEventTagEditParamsAttributionOrganizationObjectTLP = "purple"
)

func (r ThreatEventTagEditParamsAttributionOrganizationObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsAttributionOrganizationObjectTLPRed, ThreatEventTagEditParamsAttributionOrganizationObjectTLPAmber, ThreatEventTagEditParamsAttributionOrganizationObjectTLPAmberStrict, ThreatEventTagEditParamsAttributionOrganizationObjectTLPGreen, ThreatEventTagEditParamsAttributionOrganizationObjectTLPClear, ThreatEventTagEditParamsAttributionOrganizationObjectTLPPurple:
		return true
	}
	return false
}

type ThreatEventTagEditParamsExternalReference struct {
	URL         param.Field[string] `json:"url" api:"required"`
	Description param.Field[string] `json:"description"`
}

func (r ThreatEventTagEditParamsExternalReference) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagEditParamsInternalAlias struct {
	Value      param.Field[string]                                     `json:"value" api:"required"`
	Confidence param.Field[int64]                                      `json:"confidence"`
	TLP        param.Field[ThreatEventTagEditParamsInternalAliasesTLP] `json:"tlp"`
}

func (r ThreatEventTagEditParamsInternalAlias) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagEditParamsInternalAliasesTLP string

const (
	ThreatEventTagEditParamsInternalAliasesTLPRed         ThreatEventTagEditParamsInternalAliasesTLP = "red"
	ThreatEventTagEditParamsInternalAliasesTLPAmber       ThreatEventTagEditParamsInternalAliasesTLP = "amber"
	ThreatEventTagEditParamsInternalAliasesTLPAmberStrict ThreatEventTagEditParamsInternalAliasesTLP = "amber+strict"
	ThreatEventTagEditParamsInternalAliasesTLPGreen       ThreatEventTagEditParamsInternalAliasesTLP = "green"
	ThreatEventTagEditParamsInternalAliasesTLPClear       ThreatEventTagEditParamsInternalAliasesTLP = "clear"
	ThreatEventTagEditParamsInternalAliasesTLPPurple      ThreatEventTagEditParamsInternalAliasesTLP = "purple"
)

func (r ThreatEventTagEditParamsInternalAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsInternalAliasesTLPRed, ThreatEventTagEditParamsInternalAliasesTLPAmber, ThreatEventTagEditParamsInternalAliasesTLPAmberStrict, ThreatEventTagEditParamsInternalAliasesTLPGreen, ThreatEventTagEditParamsInternalAliasesTLPClear, ThreatEventTagEditParamsInternalAliasesTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagEditParamsMotiveObject].
type ThreatEventTagEditParamsMotiveUnion interface {
	ImplementsThreatEventTagEditParamsMotiveUnion()
}

type ThreatEventTagEditParamsMotiveObject struct {
	Value      param.Field[string]                                  `json:"value" api:"required"`
	Confidence param.Field[int64]                                   `json:"confidence"`
	TLP        param.Field[ThreatEventTagEditParamsMotiveObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagEditParamsMotiveObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagEditParamsMotiveObject) ImplementsThreatEventTagEditParamsMotiveUnion() {}

type ThreatEventTagEditParamsMotiveObjectTLP string

const (
	ThreatEventTagEditParamsMotiveObjectTLPRed         ThreatEventTagEditParamsMotiveObjectTLP = "red"
	ThreatEventTagEditParamsMotiveObjectTLPAmber       ThreatEventTagEditParamsMotiveObjectTLP = "amber"
	ThreatEventTagEditParamsMotiveObjectTLPAmberStrict ThreatEventTagEditParamsMotiveObjectTLP = "amber+strict"
	ThreatEventTagEditParamsMotiveObjectTLPGreen       ThreatEventTagEditParamsMotiveObjectTLP = "green"
	ThreatEventTagEditParamsMotiveObjectTLPClear       ThreatEventTagEditParamsMotiveObjectTLP = "clear"
	ThreatEventTagEditParamsMotiveObjectTLPPurple      ThreatEventTagEditParamsMotiveObjectTLP = "purple"
)

func (r ThreatEventTagEditParamsMotiveObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsMotiveObjectTLPRed, ThreatEventTagEditParamsMotiveObjectTLPAmber, ThreatEventTagEditParamsMotiveObjectTLPAmberStrict, ThreatEventTagEditParamsMotiveObjectTLPGreen, ThreatEventTagEditParamsMotiveObjectTLPClear, ThreatEventTagEditParamsMotiveObjectTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagEditParamsOpsecLevelObject].
type ThreatEventTagEditParamsOpsecLevelUnion interface {
	ImplementsThreatEventTagEditParamsOpsecLevelUnion()
}

type ThreatEventTagEditParamsOpsecLevelObject struct {
	Value      param.Field[string]                                      `json:"value" api:"required"`
	Confidence param.Field[int64]                                       `json:"confidence"`
	TLP        param.Field[ThreatEventTagEditParamsOpsecLevelObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagEditParamsOpsecLevelObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagEditParamsOpsecLevelObject) ImplementsThreatEventTagEditParamsOpsecLevelUnion() {
}

type ThreatEventTagEditParamsOpsecLevelObjectTLP string

const (
	ThreatEventTagEditParamsOpsecLevelObjectTLPRed         ThreatEventTagEditParamsOpsecLevelObjectTLP = "red"
	ThreatEventTagEditParamsOpsecLevelObjectTLPAmber       ThreatEventTagEditParamsOpsecLevelObjectTLP = "amber"
	ThreatEventTagEditParamsOpsecLevelObjectTLPAmberStrict ThreatEventTagEditParamsOpsecLevelObjectTLP = "amber+strict"
	ThreatEventTagEditParamsOpsecLevelObjectTLPGreen       ThreatEventTagEditParamsOpsecLevelObjectTLP = "green"
	ThreatEventTagEditParamsOpsecLevelObjectTLPClear       ThreatEventTagEditParamsOpsecLevelObjectTLP = "clear"
	ThreatEventTagEditParamsOpsecLevelObjectTLPPurple      ThreatEventTagEditParamsOpsecLevelObjectTLP = "purple"
)

func (r ThreatEventTagEditParamsOpsecLevelObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsOpsecLevelObjectTLPRed, ThreatEventTagEditParamsOpsecLevelObjectTLPAmber, ThreatEventTagEditParamsOpsecLevelObjectTLPAmberStrict, ThreatEventTagEditParamsOpsecLevelObjectTLPGreen, ThreatEventTagEditParamsOpsecLevelObjectTLPClear, ThreatEventTagEditParamsOpsecLevelObjectTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagEditParamsOriginCountryISOObject].
type ThreatEventTagEditParamsOriginCountryISOUnion interface {
	ImplementsThreatEventTagEditParamsOriginCountryISOUnion()
}

type ThreatEventTagEditParamsOriginCountryISOObject struct {
	Value      param.Field[string]                                            `json:"value" api:"required"`
	Confidence param.Field[int64]                                             `json:"confidence"`
	TLP        param.Field[ThreatEventTagEditParamsOriginCountryISOObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagEditParamsOriginCountryISOObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagEditParamsOriginCountryISOObject) ImplementsThreatEventTagEditParamsOriginCountryISOUnion() {
}

type ThreatEventTagEditParamsOriginCountryISOObjectTLP string

const (
	ThreatEventTagEditParamsOriginCountryISOObjectTLPRed         ThreatEventTagEditParamsOriginCountryISOObjectTLP = "red"
	ThreatEventTagEditParamsOriginCountryISOObjectTLPAmber       ThreatEventTagEditParamsOriginCountryISOObjectTLP = "amber"
	ThreatEventTagEditParamsOriginCountryISOObjectTLPAmberStrict ThreatEventTagEditParamsOriginCountryISOObjectTLP = "amber+strict"
	ThreatEventTagEditParamsOriginCountryISOObjectTLPGreen       ThreatEventTagEditParamsOriginCountryISOObjectTLP = "green"
	ThreatEventTagEditParamsOriginCountryISOObjectTLPClear       ThreatEventTagEditParamsOriginCountryISOObjectTLP = "clear"
	ThreatEventTagEditParamsOriginCountryISOObjectTLPPurple      ThreatEventTagEditParamsOriginCountryISOObjectTLP = "purple"
)

func (r ThreatEventTagEditParamsOriginCountryISOObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsOriginCountryISOObjectTLPRed, ThreatEventTagEditParamsOriginCountryISOObjectTLPAmber, ThreatEventTagEditParamsOriginCountryISOObjectTLPAmberStrict, ThreatEventTagEditParamsOriginCountryISOObjectTLPGreen, ThreatEventTagEditParamsOriginCountryISOObjectTLPClear, ThreatEventTagEditParamsOriginCountryISOObjectTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionFloat],
// [cloudforce_one.ThreatEventTagEditParamsPriorityObject].
type ThreatEventTagEditParamsPriorityUnion interface {
	ImplementsThreatEventTagEditParamsPriorityUnion()
}

type ThreatEventTagEditParamsPriorityObject struct {
	Value      param.Field[float64]                                   `json:"value" api:"required"`
	Confidence param.Field[int64]                                     `json:"confidence"`
	TLP        param.Field[ThreatEventTagEditParamsPriorityObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagEditParamsPriorityObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagEditParamsPriorityObject) ImplementsThreatEventTagEditParamsPriorityUnion() {}

type ThreatEventTagEditParamsPriorityObjectTLP string

const (
	ThreatEventTagEditParamsPriorityObjectTLPRed         ThreatEventTagEditParamsPriorityObjectTLP = "red"
	ThreatEventTagEditParamsPriorityObjectTLPAmber       ThreatEventTagEditParamsPriorityObjectTLP = "amber"
	ThreatEventTagEditParamsPriorityObjectTLPAmberStrict ThreatEventTagEditParamsPriorityObjectTLP = "amber+strict"
	ThreatEventTagEditParamsPriorityObjectTLPGreen       ThreatEventTagEditParamsPriorityObjectTLP = "green"
	ThreatEventTagEditParamsPriorityObjectTLPClear       ThreatEventTagEditParamsPriorityObjectTLP = "clear"
	ThreatEventTagEditParamsPriorityObjectTLPPurple      ThreatEventTagEditParamsPriorityObjectTLP = "purple"
)

func (r ThreatEventTagEditParamsPriorityObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsPriorityObjectTLPRed, ThreatEventTagEditParamsPriorityObjectTLPAmber, ThreatEventTagEditParamsPriorityObjectTLPAmberStrict, ThreatEventTagEditParamsPriorityObjectTLPGreen, ThreatEventTagEditParamsPriorityObjectTLPClear, ThreatEventTagEditParamsPriorityObjectTLPPurple:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagEditParamsSophisticationLevelObject].
type ThreatEventTagEditParamsSophisticationLevelUnion interface {
	ImplementsThreatEventTagEditParamsSophisticationLevelUnion()
}

type ThreatEventTagEditParamsSophisticationLevelObject struct {
	Value      param.Field[string]                                               `json:"value" api:"required"`
	Confidence param.Field[int64]                                                `json:"confidence"`
	TLP        param.Field[ThreatEventTagEditParamsSophisticationLevelObjectTLP] `json:"tlp"`
}

func (r ThreatEventTagEditParamsSophisticationLevelObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThreatEventTagEditParamsSophisticationLevelObject) ImplementsThreatEventTagEditParamsSophisticationLevelUnion() {
}

type ThreatEventTagEditParamsSophisticationLevelObjectTLP string

const (
	ThreatEventTagEditParamsSophisticationLevelObjectTLPRed         ThreatEventTagEditParamsSophisticationLevelObjectTLP = "red"
	ThreatEventTagEditParamsSophisticationLevelObjectTLPAmber       ThreatEventTagEditParamsSophisticationLevelObjectTLP = "amber"
	ThreatEventTagEditParamsSophisticationLevelObjectTLPAmberStrict ThreatEventTagEditParamsSophisticationLevelObjectTLP = "amber+strict"
	ThreatEventTagEditParamsSophisticationLevelObjectTLPGreen       ThreatEventTagEditParamsSophisticationLevelObjectTLP = "green"
	ThreatEventTagEditParamsSophisticationLevelObjectTLPClear       ThreatEventTagEditParamsSophisticationLevelObjectTLP = "clear"
	ThreatEventTagEditParamsSophisticationLevelObjectTLPPurple      ThreatEventTagEditParamsSophisticationLevelObjectTLP = "purple"
)

func (r ThreatEventTagEditParamsSophisticationLevelObjectTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsSophisticationLevelObjectTLPRed, ThreatEventTagEditParamsSophisticationLevelObjectTLPAmber, ThreatEventTagEditParamsSophisticationLevelObjectTLPAmberStrict, ThreatEventTagEditParamsSophisticationLevelObjectTLPGreen, ThreatEventTagEditParamsSophisticationLevelObjectTLPClear, ThreatEventTagEditParamsSophisticationLevelObjectTLPPurple:
		return true
	}
	return false
}

// Tag-level TLP marking. Omit to preserve existing. Cannot be cleared to null.
type ThreatEventTagEditParamsTLP string

const (
	ThreatEventTagEditParamsTLPRed         ThreatEventTagEditParamsTLP = "red"
	ThreatEventTagEditParamsTLPAmber       ThreatEventTagEditParamsTLP = "amber"
	ThreatEventTagEditParamsTLPAmberStrict ThreatEventTagEditParamsTLP = "amber+strict"
	ThreatEventTagEditParamsTLPGreen       ThreatEventTagEditParamsTLP = "green"
	ThreatEventTagEditParamsTLPClear       ThreatEventTagEditParamsTLP = "clear"
	ThreatEventTagEditParamsTLPPurple      ThreatEventTagEditParamsTLP = "purple"
)

func (r ThreatEventTagEditParamsTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsTLPRed, ThreatEventTagEditParamsTLPAmber, ThreatEventTagEditParamsTLPAmberStrict, ThreatEventTagEditParamsTLPGreen, ThreatEventTagEditParamsTLPClear, ThreatEventTagEditParamsTLPPurple:
		return true
	}
	return false
}
