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
	UUID           string `json:"uuid" api:"required"`
	Value          string `json:"value" api:"required"`
	ActiveDuration string `json:"activeDuration"`
	ActorCategory  string `json:"actorCategory"`
	// Confidence (1-10) in the actor variety (actorCategory). CFONE-only: stripped
	// from responses to non-CFONE accounts.
	ActorCategoryConfidence int64 `json:"actorCategoryConfidence"`
	// Structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: stripped from
	// responses to non-CFONE accounts.
	Aliases                    []ThreatEventTagNewResponseAlias `json:"aliases"`
	AliasGroupNames            []string                         `json:"aliasGroupNames"`
	AliasGroupNamesInternal    []string                         `json:"aliasGroupNamesInternal"`
	AnalyticPriority           float64                          `json:"analyticPriority"`
	AttributionConfidence      string                           `json:"attributionConfidence"`
	AttributionConfidenceScore int64                            `json:"attributionConfidenceScore"`
	AttributionOrganization    string                           `json:"attributionOrganization"`
	CategoryName               string                           `json:"categoryName"`
	CategoryUUID               string                           `json:"categoryUuid"`
	DateOfDiscovery            string                           `json:"dateOfDiscovery"`
	ExternalReferenceLinks     []string                         `json:"externalReferenceLinks"`
	// Structured external references ({ url, description }). Public: returned to all
	// accounts.
	ExternalReferences []ThreatEventTagNewResponseExternalReference `json:"externalReferences"`
	// Internal structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: never
	// returned to non-CFONE accounts.
	InternalAliases     []ThreatEventTagNewResponseInternalAlias `json:"internalAliases"`
	InternalDescription string                                   `json:"internalDescription"`
	Motive              string                                   `json:"motive"`
	// Confidence (1-10) in the actor motive. CFONE-only: stripped from responses to
	// non-CFONE accounts.
	MotiveConfidence int64  `json:"motiveConfidence"`
	OpsecLevel       string `json:"opsecLevel"`
	// Confidence (1-10) in the origin-country attribution. CFONE-only: stripped from
	// responses to non-CFONE accounts.
	OriginCountryConfidence int64  `json:"originCountryConfidence"`
	OriginCountryISO        string `json:"originCountryISO"`
	OriginCountryISOAlpha3  string `json:"originCountryISOAlpha3"`
	// TLP marking for the origin-country attribution. CFONE-only: stripped from
	// responses to non-CFONE accounts.
	OriginCountryTLP    ThreatEventTagNewResponseOriginCountryTLP `json:"originCountryTlp"`
	Priority            float64                                   `json:"priority"`
	SophisticationLevel string                                    `json:"sophisticationLevel"`
	JSON                threatEventTagNewResponseJSON             `json:"-"`
}

// threatEventTagNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagNewResponse]
type threatEventTagNewResponseJSON struct {
	UUID                       apijson.Field
	Value                      apijson.Field
	ActiveDuration             apijson.Field
	ActorCategory              apijson.Field
	ActorCategoryConfidence    apijson.Field
	Aliases                    apijson.Field
	AliasGroupNames            apijson.Field
	AliasGroupNamesInternal    apijson.Field
	AnalyticPriority           apijson.Field
	AttributionConfidence      apijson.Field
	AttributionConfidenceScore apijson.Field
	AttributionOrganization    apijson.Field
	CategoryName               apijson.Field
	CategoryUUID               apijson.Field
	DateOfDiscovery            apijson.Field
	ExternalReferenceLinks     apijson.Field
	ExternalReferences         apijson.Field
	InternalAliases            apijson.Field
	InternalDescription        apijson.Field
	Motive                     apijson.Field
	MotiveConfidence           apijson.Field
	OpsecLevel                 apijson.Field
	OriginCountryConfidence    apijson.Field
	OriginCountryISO           apijson.Field
	OriginCountryISOAlpha3     apijson.Field
	OriginCountryTLP           apijson.Field
	Priority                   apijson.Field
	SophisticationLevel        apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ThreatEventTagNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseJSON) RawJSON() string {
	return r.raw
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
	ThreatEventTagNewResponseAliasesTLPRed   ThreatEventTagNewResponseAliasesTLP = "red"
	ThreatEventTagNewResponseAliasesTLPAmber ThreatEventTagNewResponseAliasesTLP = "amber"
	ThreatEventTagNewResponseAliasesTLPGreen ThreatEventTagNewResponseAliasesTLP = "green"
	ThreatEventTagNewResponseAliasesTLPWhite ThreatEventTagNewResponseAliasesTLP = "white"
)

func (r ThreatEventTagNewResponseAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseAliasesTLPRed, ThreatEventTagNewResponseAliasesTLPAmber, ThreatEventTagNewResponseAliasesTLPGreen, ThreatEventTagNewResponseAliasesTLPWhite:
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
	ThreatEventTagNewResponseInternalAliasesTLPRed   ThreatEventTagNewResponseInternalAliasesTLP = "red"
	ThreatEventTagNewResponseInternalAliasesTLPAmber ThreatEventTagNewResponseInternalAliasesTLP = "amber"
	ThreatEventTagNewResponseInternalAliasesTLPGreen ThreatEventTagNewResponseInternalAliasesTLP = "green"
	ThreatEventTagNewResponseInternalAliasesTLPWhite ThreatEventTagNewResponseInternalAliasesTLP = "white"
)

func (r ThreatEventTagNewResponseInternalAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseInternalAliasesTLPRed, ThreatEventTagNewResponseInternalAliasesTLPAmber, ThreatEventTagNewResponseInternalAliasesTLPGreen, ThreatEventTagNewResponseInternalAliasesTLPWhite:
		return true
	}
	return false
}

// TLP marking for the origin-country attribution. CFONE-only: stripped from
// responses to non-CFONE accounts.
type ThreatEventTagNewResponseOriginCountryTLP string

const (
	ThreatEventTagNewResponseOriginCountryTLPRed   ThreatEventTagNewResponseOriginCountryTLP = "red"
	ThreatEventTagNewResponseOriginCountryTLPAmber ThreatEventTagNewResponseOriginCountryTLP = "amber"
	ThreatEventTagNewResponseOriginCountryTLPGreen ThreatEventTagNewResponseOriginCountryTLP = "green"
	ThreatEventTagNewResponseOriginCountryTLPWhite ThreatEventTagNewResponseOriginCountryTLP = "white"
)

func (r ThreatEventTagNewResponseOriginCountryTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewResponseOriginCountryTLPRed, ThreatEventTagNewResponseOriginCountryTLPAmber, ThreatEventTagNewResponseOriginCountryTLPGreen, ThreatEventTagNewResponseOriginCountryTLPWhite:
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
	UUID           string `json:"uuid" api:"required"`
	Value          string `json:"value" api:"required"`
	ActiveDuration string `json:"activeDuration"`
	ActorCategory  string `json:"actorCategory"`
	// Confidence (1-10) in the actor variety (actorCategory). CFONE-only: stripped
	// from responses to non-CFONE accounts.
	ActorCategoryConfidence int64 `json:"actorCategoryConfidence"`
	// Structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: stripped from
	// responses to non-CFONE accounts.
	Aliases                    []ThreatEventTagListResponseTagsAlias `json:"aliases"`
	AliasGroupNames            []string                              `json:"aliasGroupNames"`
	AliasGroupNamesInternal    []string                              `json:"aliasGroupNamesInternal"`
	AnalyticPriority           float64                               `json:"analyticPriority"`
	AttributionConfidence      string                                `json:"attributionConfidence"`
	AttributionConfidenceScore int64                                 `json:"attributionConfidenceScore"`
	AttributionOrganization    string                                `json:"attributionOrganization"`
	CategoryName               string                                `json:"categoryName"`
	CategoryUUID               string                                `json:"categoryUuid"`
	DateOfDiscovery            string                                `json:"dateOfDiscovery"`
	ExternalReferenceLinks     []string                              `json:"externalReferenceLinks"`
	// Structured external references ({ url, description }). Public: returned to all
	// accounts.
	ExternalReferences []ThreatEventTagListResponseTagsExternalReference `json:"externalReferences"`
	// Internal structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: never
	// returned to non-CFONE accounts.
	InternalAliases     []ThreatEventTagListResponseTagsInternalAlias `json:"internalAliases"`
	InternalDescription string                                        `json:"internalDescription"`
	Motive              string                                        `json:"motive"`
	// Confidence (1-10) in the actor motive. CFONE-only: stripped from responses to
	// non-CFONE accounts.
	MotiveConfidence int64  `json:"motiveConfidence"`
	OpsecLevel       string `json:"opsecLevel"`
	// Confidence (1-10) in the origin-country attribution. CFONE-only: stripped from
	// responses to non-CFONE accounts.
	OriginCountryConfidence int64  `json:"originCountryConfidence"`
	OriginCountryISO        string `json:"originCountryISO"`
	OriginCountryISOAlpha3  string `json:"originCountryISOAlpha3"`
	// TLP marking for the origin-country attribution. CFONE-only: stripped from
	// responses to non-CFONE accounts.
	OriginCountryTLP    ThreatEventTagListResponseTagsOriginCountryTLP `json:"originCountryTlp"`
	Priority            float64                                        `json:"priority"`
	SophisticationLevel string                                         `json:"sophisticationLevel"`
	JSON                threatEventTagListResponseTagJSON              `json:"-"`
}

// threatEventTagListResponseTagJSON contains the JSON metadata for the struct
// [ThreatEventTagListResponseTag]
type threatEventTagListResponseTagJSON struct {
	UUID                       apijson.Field
	Value                      apijson.Field
	ActiveDuration             apijson.Field
	ActorCategory              apijson.Field
	ActorCategoryConfidence    apijson.Field
	Aliases                    apijson.Field
	AliasGroupNames            apijson.Field
	AliasGroupNamesInternal    apijson.Field
	AnalyticPriority           apijson.Field
	AttributionConfidence      apijson.Field
	AttributionConfidenceScore apijson.Field
	AttributionOrganization    apijson.Field
	CategoryName               apijson.Field
	CategoryUUID               apijson.Field
	DateOfDiscovery            apijson.Field
	ExternalReferenceLinks     apijson.Field
	ExternalReferences         apijson.Field
	InternalAliases            apijson.Field
	InternalDescription        apijson.Field
	Motive                     apijson.Field
	MotiveConfidence           apijson.Field
	OpsecLevel                 apijson.Field
	OriginCountryConfidence    apijson.Field
	OriginCountryISO           apijson.Field
	OriginCountryISOAlpha3     apijson.Field
	OriginCountryTLP           apijson.Field
	Priority                   apijson.Field
	SophisticationLevel        apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ThreatEventTagListResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagListResponseTagJSON) RawJSON() string {
	return r.raw
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
	ThreatEventTagListResponseTagsAliasesTLPRed   ThreatEventTagListResponseTagsAliasesTLP = "red"
	ThreatEventTagListResponseTagsAliasesTLPAmber ThreatEventTagListResponseTagsAliasesTLP = "amber"
	ThreatEventTagListResponseTagsAliasesTLPGreen ThreatEventTagListResponseTagsAliasesTLP = "green"
	ThreatEventTagListResponseTagsAliasesTLPWhite ThreatEventTagListResponseTagsAliasesTLP = "white"
)

func (r ThreatEventTagListResponseTagsAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsAliasesTLPRed, ThreatEventTagListResponseTagsAliasesTLPAmber, ThreatEventTagListResponseTagsAliasesTLPGreen, ThreatEventTagListResponseTagsAliasesTLPWhite:
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
	ThreatEventTagListResponseTagsInternalAliasesTLPRed   ThreatEventTagListResponseTagsInternalAliasesTLP = "red"
	ThreatEventTagListResponseTagsInternalAliasesTLPAmber ThreatEventTagListResponseTagsInternalAliasesTLP = "amber"
	ThreatEventTagListResponseTagsInternalAliasesTLPGreen ThreatEventTagListResponseTagsInternalAliasesTLP = "green"
	ThreatEventTagListResponseTagsInternalAliasesTLPWhite ThreatEventTagListResponseTagsInternalAliasesTLP = "white"
)

func (r ThreatEventTagListResponseTagsInternalAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsInternalAliasesTLPRed, ThreatEventTagListResponseTagsInternalAliasesTLPAmber, ThreatEventTagListResponseTagsInternalAliasesTLPGreen, ThreatEventTagListResponseTagsInternalAliasesTLPWhite:
		return true
	}
	return false
}

// TLP marking for the origin-country attribution. CFONE-only: stripped from
// responses to non-CFONE accounts.
type ThreatEventTagListResponseTagsOriginCountryTLP string

const (
	ThreatEventTagListResponseTagsOriginCountryTLPRed   ThreatEventTagListResponseTagsOriginCountryTLP = "red"
	ThreatEventTagListResponseTagsOriginCountryTLPAmber ThreatEventTagListResponseTagsOriginCountryTLP = "amber"
	ThreatEventTagListResponseTagsOriginCountryTLPGreen ThreatEventTagListResponseTagsOriginCountryTLP = "green"
	ThreatEventTagListResponseTagsOriginCountryTLPWhite ThreatEventTagListResponseTagsOriginCountryTLP = "white"
)

func (r ThreatEventTagListResponseTagsOriginCountryTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagListResponseTagsOriginCountryTLPRed, ThreatEventTagListResponseTagsOriginCountryTLPAmber, ThreatEventTagListResponseTagsOriginCountryTLPGreen, ThreatEventTagListResponseTagsOriginCountryTLPWhite:
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
	UUID           string `json:"uuid" api:"required"`
	Value          string `json:"value" api:"required"`
	ActiveDuration string `json:"activeDuration"`
	ActorCategory  string `json:"actorCategory"`
	// Confidence (1-10) in the actor variety (actorCategory). CFONE-only: stripped
	// from responses to non-CFONE accounts.
	ActorCategoryConfidence int64 `json:"actorCategoryConfidence"`
	// Structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: stripped from
	// responses to non-CFONE accounts.
	Aliases                    []ThreatEventTagEditResponseAlias `json:"aliases"`
	AliasGroupNames            []string                          `json:"aliasGroupNames"`
	AliasGroupNamesInternal    []string                          `json:"aliasGroupNamesInternal"`
	AnalyticPriority           float64                           `json:"analyticPriority"`
	AttributionConfidence      string                            `json:"attributionConfidence"`
	AttributionConfidenceScore int64                             `json:"attributionConfidenceScore"`
	AttributionOrganization    string                            `json:"attributionOrganization"`
	CategoryName               string                            `json:"categoryName"`
	CategoryUUID               string                            `json:"categoryUuid"`
	DateOfDiscovery            string                            `json:"dateOfDiscovery"`
	ExternalReferenceLinks     []string                          `json:"externalReferenceLinks"`
	// Structured external references ({ url, description }). Public: returned to all
	// accounts.
	ExternalReferences []ThreatEventTagEditResponseExternalReference `json:"externalReferences"`
	// Internal structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: never
	// returned to non-CFONE accounts.
	InternalAliases     []ThreatEventTagEditResponseInternalAlias `json:"internalAliases"`
	InternalDescription string                                    `json:"internalDescription"`
	Motive              string                                    `json:"motive"`
	// Confidence (1-10) in the actor motive. CFONE-only: stripped from responses to
	// non-CFONE accounts.
	MotiveConfidence int64  `json:"motiveConfidence"`
	OpsecLevel       string `json:"opsecLevel"`
	// Confidence (1-10) in the origin-country attribution. CFONE-only: stripped from
	// responses to non-CFONE accounts.
	OriginCountryConfidence int64  `json:"originCountryConfidence"`
	OriginCountryISO        string `json:"originCountryISO"`
	OriginCountryISOAlpha3  string `json:"originCountryISOAlpha3"`
	// TLP marking for the origin-country attribution. CFONE-only: stripped from
	// responses to non-CFONE accounts.
	OriginCountryTLP    ThreatEventTagEditResponseOriginCountryTLP `json:"originCountryTlp"`
	Priority            float64                                    `json:"priority"`
	SophisticationLevel string                                     `json:"sophisticationLevel"`
	JSON                threatEventTagEditResponseJSON             `json:"-"`
}

// threatEventTagEditResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagEditResponse]
type threatEventTagEditResponseJSON struct {
	UUID                       apijson.Field
	Value                      apijson.Field
	ActiveDuration             apijson.Field
	ActorCategory              apijson.Field
	ActorCategoryConfidence    apijson.Field
	Aliases                    apijson.Field
	AliasGroupNames            apijson.Field
	AliasGroupNamesInternal    apijson.Field
	AnalyticPriority           apijson.Field
	AttributionConfidence      apijson.Field
	AttributionConfidenceScore apijson.Field
	AttributionOrganization    apijson.Field
	CategoryName               apijson.Field
	CategoryUUID               apijson.Field
	DateOfDiscovery            apijson.Field
	ExternalReferenceLinks     apijson.Field
	ExternalReferences         apijson.Field
	InternalAliases            apijson.Field
	InternalDescription        apijson.Field
	Motive                     apijson.Field
	MotiveConfidence           apijson.Field
	OpsecLevel                 apijson.Field
	OriginCountryConfidence    apijson.Field
	OriginCountryISO           apijson.Field
	OriginCountryISOAlpha3     apijson.Field
	OriginCountryTLP           apijson.Field
	Priority                   apijson.Field
	SophisticationLevel        apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ThreatEventTagEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagEditResponseJSON) RawJSON() string {
	return r.raw
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
	ThreatEventTagEditResponseAliasesTLPRed   ThreatEventTagEditResponseAliasesTLP = "red"
	ThreatEventTagEditResponseAliasesTLPAmber ThreatEventTagEditResponseAliasesTLP = "amber"
	ThreatEventTagEditResponseAliasesTLPGreen ThreatEventTagEditResponseAliasesTLP = "green"
	ThreatEventTagEditResponseAliasesTLPWhite ThreatEventTagEditResponseAliasesTLP = "white"
)

func (r ThreatEventTagEditResponseAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseAliasesTLPRed, ThreatEventTagEditResponseAliasesTLPAmber, ThreatEventTagEditResponseAliasesTLPGreen, ThreatEventTagEditResponseAliasesTLPWhite:
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
	ThreatEventTagEditResponseInternalAliasesTLPRed   ThreatEventTagEditResponseInternalAliasesTLP = "red"
	ThreatEventTagEditResponseInternalAliasesTLPAmber ThreatEventTagEditResponseInternalAliasesTLP = "amber"
	ThreatEventTagEditResponseInternalAliasesTLPGreen ThreatEventTagEditResponseInternalAliasesTLP = "green"
	ThreatEventTagEditResponseInternalAliasesTLPWhite ThreatEventTagEditResponseInternalAliasesTLP = "white"
)

func (r ThreatEventTagEditResponseInternalAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseInternalAliasesTLPRed, ThreatEventTagEditResponseInternalAliasesTLPAmber, ThreatEventTagEditResponseInternalAliasesTLPGreen, ThreatEventTagEditResponseInternalAliasesTLPWhite:
		return true
	}
	return false
}

// TLP marking for the origin-country attribution. CFONE-only: stripped from
// responses to non-CFONE accounts.
type ThreatEventTagEditResponseOriginCountryTLP string

const (
	ThreatEventTagEditResponseOriginCountryTLPRed   ThreatEventTagEditResponseOriginCountryTLP = "red"
	ThreatEventTagEditResponseOriginCountryTLPAmber ThreatEventTagEditResponseOriginCountryTLP = "amber"
	ThreatEventTagEditResponseOriginCountryTLPGreen ThreatEventTagEditResponseOriginCountryTLP = "green"
	ThreatEventTagEditResponseOriginCountryTLPWhite ThreatEventTagEditResponseOriginCountryTLP = "white"
)

func (r ThreatEventTagEditResponseOriginCountryTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditResponseOriginCountryTLPRed, ThreatEventTagEditResponseOriginCountryTLPAmber, ThreatEventTagEditResponseOriginCountryTLPGreen, ThreatEventTagEditResponseOriginCountryTLPWhite:
		return true
	}
	return false
}

type ThreatEventTagNewParams struct {
	// Account ID.
	AccountID      param.Field[string] `path:"account_id" api:"required"`
	Value          param.Field[string] `json:"value" api:"required"`
	ActiveDuration param.Field[string] `json:"activeDuration"`
	// Actor variety. Allowed values: Activist, Competitor, Customer, Crime Syndicate,
	// Former Employee, Nation State, Organized Crime, Nation State Affiliated,
	// Terrorist, Unaffiliated.
	ActorCategory param.Field[string] `json:"actorCategory"`
	// Confidence (1-10) in the actor variety (actorCategory). CFONE-only: stripped
	// from responses to non-CFONE accounts.
	ActorCategoryConfidence param.Field[int64] `json:"actorCategoryConfidence"`
	// Structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: stripped from
	// responses to non-CFONE accounts.
	Aliases                    param.Field[[]ThreatEventTagNewParamsAlias] `json:"aliases"`
	AliasGroupNames            param.Field[[]string]                       `json:"aliasGroupNames"`
	AliasGroupNamesInternal    param.Field[[]string]                       `json:"aliasGroupNamesInternal"`
	AnalyticPriority           param.Field[float64]                        `json:"analyticPriority"`
	AttributionConfidence      param.Field[string]                         `json:"attributionConfidence"`
	AttributionConfidenceScore param.Field[int64]                          `json:"attributionConfidenceScore"`
	AttributionOrganization    param.Field[string]                         `json:"attributionOrganization"`
	CategoryUUID               param.Field[string]                         `json:"categoryUuid"`
	// Date the actor was discovered (ISO YYYY-MM-DD).
	DateOfDiscovery        param.Field[string]   `json:"dateOfDiscovery"`
	ExternalReferenceLinks param.Field[[]string] `json:"externalReferenceLinks"`
	// Structured external references ({ url, description }). Public: returned to all
	// accounts.
	ExternalReferences param.Field[[]ThreatEventTagNewParamsExternalReference] `json:"externalReferences"`
	// Internal structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: never
	// returned to non-CFONE accounts.
	InternalAliases     param.Field[[]ThreatEventTagNewParamsInternalAlias] `json:"internalAliases"`
	InternalDescription param.Field[string]                                 `json:"internalDescription"`
	// Actor motive. Allowed values: Convenience, Fear, Fun, Financial, Grudge,
	// Ideology, Espionage.
	Motive param.Field[string] `json:"motive"`
	// Confidence (1-10) in the actor motive. CFONE-only: stripped from responses to
	// non-CFONE accounts.
	MotiveConfidence param.Field[int64]  `json:"motiveConfidence"`
	OpsecLevel       param.Field[string] `json:"opsecLevel"`
	// Confidence (1-10) in the origin-country attribution. CFONE-only: stripped from
	// responses to non-CFONE accounts.
	OriginCountryConfidence param.Field[int64]  `json:"originCountryConfidence"`
	OriginCountryISO        param.Field[string] `json:"originCountryISO"`
	// TLP marking for the origin-country attribution. CFONE-only: stripped from
	// responses to non-CFONE accounts.
	OriginCountryTLP    param.Field[ThreatEventTagNewParamsOriginCountryTLP] `json:"originCountryTlp"`
	Priority            param.Field[float64]                                 `json:"priority"`
	SophisticationLevel param.Field[string]                                  `json:"sophisticationLevel"`
}

func (r ThreatEventTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	ThreatEventTagNewParamsAliasesTLPRed   ThreatEventTagNewParamsAliasesTLP = "red"
	ThreatEventTagNewParamsAliasesTLPAmber ThreatEventTagNewParamsAliasesTLP = "amber"
	ThreatEventTagNewParamsAliasesTLPGreen ThreatEventTagNewParamsAliasesTLP = "green"
	ThreatEventTagNewParamsAliasesTLPWhite ThreatEventTagNewParamsAliasesTLP = "white"
)

func (r ThreatEventTagNewParamsAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsAliasesTLPRed, ThreatEventTagNewParamsAliasesTLPAmber, ThreatEventTagNewParamsAliasesTLPGreen, ThreatEventTagNewParamsAliasesTLPWhite:
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
	ThreatEventTagNewParamsInternalAliasesTLPRed   ThreatEventTagNewParamsInternalAliasesTLP = "red"
	ThreatEventTagNewParamsInternalAliasesTLPAmber ThreatEventTagNewParamsInternalAliasesTLP = "amber"
	ThreatEventTagNewParamsInternalAliasesTLPGreen ThreatEventTagNewParamsInternalAliasesTLP = "green"
	ThreatEventTagNewParamsInternalAliasesTLPWhite ThreatEventTagNewParamsInternalAliasesTLP = "white"
)

func (r ThreatEventTagNewParamsInternalAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsInternalAliasesTLPRed, ThreatEventTagNewParamsInternalAliasesTLPAmber, ThreatEventTagNewParamsInternalAliasesTLPGreen, ThreatEventTagNewParamsInternalAliasesTLPWhite:
		return true
	}
	return false
}

// TLP marking for the origin-country attribution. CFONE-only: stripped from
// responses to non-CFONE accounts.
type ThreatEventTagNewParamsOriginCountryTLP string

const (
	ThreatEventTagNewParamsOriginCountryTLPRed   ThreatEventTagNewParamsOriginCountryTLP = "red"
	ThreatEventTagNewParamsOriginCountryTLPAmber ThreatEventTagNewParamsOriginCountryTLP = "amber"
	ThreatEventTagNewParamsOriginCountryTLPGreen ThreatEventTagNewParamsOriginCountryTLP = "green"
	ThreatEventTagNewParamsOriginCountryTLPWhite ThreatEventTagNewParamsOriginCountryTLP = "white"
)

func (r ThreatEventTagNewParamsOriginCountryTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagNewParamsOriginCountryTLPRed, ThreatEventTagNewParamsOriginCountryTLPAmber, ThreatEventTagNewParamsOriginCountryTLPGreen, ThreatEventTagNewParamsOriginCountryTLPWhite:
		return true
	}
	return false
}

type ThreatEventTagListParams struct {
	// Account ID.
	AccountID    param.Field[string] `path:"account_id" api:"required"`
	CategoryUUID param.Field[string] `query:"categoryUuid"`
	// Structured filters as a JSON array of {field, op, value} objects. Searchable
	// fields: uuid, value, actorCategory, actorCategoryConfidence, aliasGroupNames,
	// attributionConfidence, attributionConfidenceScore, attributionOrganization,
	// categoryName, motive, motiveConfidence, opsecLevel, originCountryISO,
	// originCountryConfidence, sophisticationLevel, priority, analyticPriority.
	// Operators: equals, not, contains, startsWith, endsWith, gt, lt, gte, lte, like,
	// in, find. Use 'in' for bulk OR within a single field, e.g.
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
	// Legacy free-text substring match on tag value.
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

type ThreatEventTagListParamsFilter struct {
	// Tag field to search on. Allowed: uuid, value, actorCategory,
	// actorCategoryConfidence, aliasGroupNames, attributionConfidence,
	// attributionConfidenceScore, attributionOrganization, categoryName, motive,
	// motiveConfidence, opsecLevel, originCountryISO, originCountryConfidence,
	// sophisticationLevel, priority, analyticPriority.
	Field param.Field[ThreatEventTagListParamsFiltersField] `query:"field" api:"required"`
	// Search operator. Use 'in' for bulk OR within a single field, e.g.
	// {field:"originCountryISO", op:"in", value:["IR","CN"]}.
	Op param.Field[ThreatEventTagListParamsFiltersOp] `query:"op" api:"required"`
	// Search value. String or number for most operators. Array for 'in' (max 100
	// items). Country values may be passed as alpha-2, alpha-3, name, or common alias
	// (e.g. 'iran', 'IR', 'IRN') and are normalized to alpha-2 server-side.
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

// Tag field to search on. Allowed: uuid, value, actorCategory,
// actorCategoryConfidence, aliasGroupNames, attributionConfidence,
// attributionConfidenceScore, attributionOrganization, categoryName, motive,
// motiveConfidence, opsecLevel, originCountryISO, originCountryConfidence,
// sophisticationLevel, priority, analyticPriority.
type ThreatEventTagListParamsFiltersField string

const (
	ThreatEventTagListParamsFiltersFieldUUID                       ThreatEventTagListParamsFiltersField = "uuid"
	ThreatEventTagListParamsFiltersFieldValue                      ThreatEventTagListParamsFiltersField = "value"
	ThreatEventTagListParamsFiltersFieldActorCategory              ThreatEventTagListParamsFiltersField = "actorCategory"
	ThreatEventTagListParamsFiltersFieldActorCategoryConfidence    ThreatEventTagListParamsFiltersField = "actorCategoryConfidence"
	ThreatEventTagListParamsFiltersFieldAliasGroupNames            ThreatEventTagListParamsFiltersField = "aliasGroupNames"
	ThreatEventTagListParamsFiltersFieldAttributionConfidence      ThreatEventTagListParamsFiltersField = "attributionConfidence"
	ThreatEventTagListParamsFiltersFieldAttributionConfidenceScore ThreatEventTagListParamsFiltersField = "attributionConfidenceScore"
	ThreatEventTagListParamsFiltersFieldAttributionOrganization    ThreatEventTagListParamsFiltersField = "attributionOrganization"
	ThreatEventTagListParamsFiltersFieldCategoryName               ThreatEventTagListParamsFiltersField = "categoryName"
	ThreatEventTagListParamsFiltersFieldMotive                     ThreatEventTagListParamsFiltersField = "motive"
	ThreatEventTagListParamsFiltersFieldMotiveConfidence           ThreatEventTagListParamsFiltersField = "motiveConfidence"
	ThreatEventTagListParamsFiltersFieldOpsecLevel                 ThreatEventTagListParamsFiltersField = "opsecLevel"
	ThreatEventTagListParamsFiltersFieldOriginCountryISO           ThreatEventTagListParamsFiltersField = "originCountryISO"
	ThreatEventTagListParamsFiltersFieldOriginCountryConfidence    ThreatEventTagListParamsFiltersField = "originCountryConfidence"
	ThreatEventTagListParamsFiltersFieldSophisticationLevel        ThreatEventTagListParamsFiltersField = "sophisticationLevel"
	ThreatEventTagListParamsFiltersFieldPriority                   ThreatEventTagListParamsFiltersField = "priority"
	ThreatEventTagListParamsFiltersFieldAnalyticPriority           ThreatEventTagListParamsFiltersField = "analyticPriority"
)

func (r ThreatEventTagListParamsFiltersField) IsKnown() bool {
	switch r {
	case ThreatEventTagListParamsFiltersFieldUUID, ThreatEventTagListParamsFiltersFieldValue, ThreatEventTagListParamsFiltersFieldActorCategory, ThreatEventTagListParamsFiltersFieldActorCategoryConfidence, ThreatEventTagListParamsFiltersFieldAliasGroupNames, ThreatEventTagListParamsFiltersFieldAttributionConfidence, ThreatEventTagListParamsFiltersFieldAttributionConfidenceScore, ThreatEventTagListParamsFiltersFieldAttributionOrganization, ThreatEventTagListParamsFiltersFieldCategoryName, ThreatEventTagListParamsFiltersFieldMotive, ThreatEventTagListParamsFiltersFieldMotiveConfidence, ThreatEventTagListParamsFiltersFieldOpsecLevel, ThreatEventTagListParamsFiltersFieldOriginCountryISO, ThreatEventTagListParamsFiltersFieldOriginCountryConfidence, ThreatEventTagListParamsFiltersFieldSophisticationLevel, ThreatEventTagListParamsFiltersFieldPriority, ThreatEventTagListParamsFiltersFieldAnalyticPriority:
		return true
	}
	return false
}

// Search operator. Use 'in' for bulk OR within a single field, e.g.
// {field:"originCountryISO", op:"in", value:["IR","CN"]}.
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
// items). Country values may be passed as alpha-2, alpha-3, name, or common alias
// (e.g. 'iran', 'IR', 'IRN') and are normalized to alpha-2 server-side.
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
	AccountID      param.Field[string] `path:"account_id" api:"required"`
	ActiveDuration param.Field[string] `json:"activeDuration"`
	// Actor variety. Allowed values: Activist, Competitor, Customer, Crime Syndicate,
	// Former Employee, Nation State, Organized Crime, Nation State Affiliated,
	// Terrorist, Unaffiliated.
	ActorCategory param.Field[string] `json:"actorCategory"`
	// Confidence (1-10) in the actor variety (actorCategory). CFONE-only: stripped
	// from responses to non-CFONE accounts.
	ActorCategoryConfidence param.Field[int64] `json:"actorCategoryConfidence"`
	// Structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: stripped from
	// responses to non-CFONE accounts.
	Aliases                    param.Field[[]ThreatEventTagEditParamsAlias] `json:"aliases"`
	AliasGroupNames            param.Field[[]string]                        `json:"aliasGroupNames"`
	AliasGroupNamesInternal    param.Field[[]string]                        `json:"aliasGroupNamesInternal"`
	AnalyticPriority           param.Field[float64]                         `json:"analyticPriority"`
	AttributionConfidence      param.Field[string]                          `json:"attributionConfidence"`
	AttributionConfidenceScore param.Field[int64]                           `json:"attributionConfidenceScore"`
	AttributionOrganization    param.Field[string]                          `json:"attributionOrganization"`
	CategoryUUID               param.Field[string]                          `json:"categoryUuid"`
	// Date the actor was discovered (ISO YYYY-MM-DD).
	DateOfDiscovery        param.Field[string]   `json:"dateOfDiscovery"`
	ExternalReferenceLinks param.Field[[]string] `json:"externalReferenceLinks"`
	// Structured external references ({ url, description }). Public: returned to all
	// accounts.
	ExternalReferences param.Field[[]ThreatEventTagEditParamsExternalReference] `json:"externalReferences"`
	// Internal structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: never
	// returned to non-CFONE accounts.
	InternalAliases     param.Field[[]ThreatEventTagEditParamsInternalAlias] `json:"internalAliases"`
	InternalDescription param.Field[string]                                  `json:"internalDescription"`
	// Actor motive. Allowed values: Convenience, Fear, Fun, Financial, Grudge,
	// Ideology, Espionage.
	Motive param.Field[string] `json:"motive"`
	// Confidence (1-10) in the actor motive. CFONE-only: stripped from responses to
	// non-CFONE accounts.
	MotiveConfidence param.Field[int64]  `json:"motiveConfidence"`
	OpsecLevel       param.Field[string] `json:"opsecLevel"`
	// Confidence (1-10) in the origin-country attribution. CFONE-only: stripped from
	// responses to non-CFONE accounts.
	OriginCountryConfidence param.Field[int64]  `json:"originCountryConfidence"`
	OriginCountryISO        param.Field[string] `json:"originCountryISO"`
	// TLP marking for the origin-country attribution. CFONE-only: stripped from
	// responses to non-CFONE accounts.
	OriginCountryTLP    param.Field[ThreatEventTagEditParamsOriginCountryTLP] `json:"originCountryTlp"`
	Priority            param.Field[float64]                                  `json:"priority"`
	SophisticationLevel param.Field[string]                                   `json:"sophisticationLevel"`
	Value               param.Field[string]                                   `json:"value"`
}

func (r ThreatEventTagEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	ThreatEventTagEditParamsAliasesTLPRed   ThreatEventTagEditParamsAliasesTLP = "red"
	ThreatEventTagEditParamsAliasesTLPAmber ThreatEventTagEditParamsAliasesTLP = "amber"
	ThreatEventTagEditParamsAliasesTLPGreen ThreatEventTagEditParamsAliasesTLP = "green"
	ThreatEventTagEditParamsAliasesTLPWhite ThreatEventTagEditParamsAliasesTLP = "white"
)

func (r ThreatEventTagEditParamsAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsAliasesTLPRed, ThreatEventTagEditParamsAliasesTLPAmber, ThreatEventTagEditParamsAliasesTLPGreen, ThreatEventTagEditParamsAliasesTLPWhite:
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
	ThreatEventTagEditParamsInternalAliasesTLPRed   ThreatEventTagEditParamsInternalAliasesTLP = "red"
	ThreatEventTagEditParamsInternalAliasesTLPAmber ThreatEventTagEditParamsInternalAliasesTLP = "amber"
	ThreatEventTagEditParamsInternalAliasesTLPGreen ThreatEventTagEditParamsInternalAliasesTLP = "green"
	ThreatEventTagEditParamsInternalAliasesTLPWhite ThreatEventTagEditParamsInternalAliasesTLP = "white"
)

func (r ThreatEventTagEditParamsInternalAliasesTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsInternalAliasesTLPRed, ThreatEventTagEditParamsInternalAliasesTLPAmber, ThreatEventTagEditParamsInternalAliasesTLPGreen, ThreatEventTagEditParamsInternalAliasesTLPWhite:
		return true
	}
	return false
}

// TLP marking for the origin-country attribution. CFONE-only: stripped from
// responses to non-CFONE accounts.
type ThreatEventTagEditParamsOriginCountryTLP string

const (
	ThreatEventTagEditParamsOriginCountryTLPRed   ThreatEventTagEditParamsOriginCountryTLP = "red"
	ThreatEventTagEditParamsOriginCountryTLPAmber ThreatEventTagEditParamsOriginCountryTLP = "amber"
	ThreatEventTagEditParamsOriginCountryTLPGreen ThreatEventTagEditParamsOriginCountryTLP = "green"
	ThreatEventTagEditParamsOriginCountryTLPWhite ThreatEventTagEditParamsOriginCountryTLP = "white"
)

func (r ThreatEventTagEditParamsOriginCountryTLP) IsKnown() bool {
	switch r {
	case ThreatEventTagEditParamsOriginCountryTLPRed, ThreatEventTagEditParamsOriginCountryTLPAmber, ThreatEventTagEditParamsOriginCountryTLPGreen, ThreatEventTagEditParamsOriginCountryTLPWhite:
		return true
	}
	return false
}
