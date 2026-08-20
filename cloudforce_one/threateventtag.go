// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
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
	Options []option.RequestOption
}

// NewThreatEventTagService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThreatEventTagService(opts ...option.RequestOption) (r *ThreatEventTagService) {
	r = &ThreatEventTagService{}
	r.Options = opts
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

type ThreatEventTagNewResponse struct {
	UUID           string `json:"uuid" api:"required"`
	Value          string `json:"value" api:"required"`
	ActiveDuration string `json:"activeDuration"`
	ActorCategory  string `json:"actorCategory"`
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
	// Internal structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: never
	// returned to non-CFONE accounts.
	InternalAliases        []ThreatEventTagNewResponseInternalAlias `json:"internalAliases"`
	InternalDescription    string                                   `json:"internalDescription"`
	Motive                 string                                   `json:"motive"`
	OpsecLevel             string                                   `json:"opsecLevel"`
	OriginCountryISO       string                                   `json:"originCountryISO"`
	OriginCountryISOAlpha3 string                                   `json:"originCountryISOAlpha3"`
	Priority               float64                                  `json:"priority"`
	SophisticationLevel    string                                   `json:"sophisticationLevel"`
	JSON                   threatEventTagNewResponseJSON            `json:"-"`
}

// threatEventTagNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagNewResponse]
type threatEventTagNewResponseJSON struct {
	UUID                       apijson.Field
	Value                      apijson.Field
	ActiveDuration             apijson.Field
	ActorCategory              apijson.Field
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
	InternalAliases            apijson.Field
	InternalDescription        apijson.Field
	Motive                     apijson.Field
	OpsecLevel                 apijson.Field
	OriginCountryISO           apijson.Field
	OriginCountryISOAlpha3     apijson.Field
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

type ThreatEventTagNewParams struct {
	// Account ID.
	AccountID      param.Field[string] `path:"account_id" api:"required"`
	Value          param.Field[string] `json:"value" api:"required"`
	ActiveDuration param.Field[string] `json:"activeDuration"`
	// Actor variety. Allowed values: Activist, Competitor, Customer, Crime Syndicate,
	// Former Employee, Nation State, Organized Crime, Nation State Affiliated,
	// Terrorist, Unaffiliated.
	ActorCategory param.Field[string] `json:"actorCategory"`
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
	// Internal structured aliases ({ value, confidence 1-10, tlp }). CFONE-only: never
	// returned to non-CFONE accounts.
	InternalAliases     param.Field[[]ThreatEventTagNewParamsInternalAlias] `json:"internalAliases"`
	InternalDescription param.Field[string]                                 `json:"internalDescription"`
	// Actor motive. Allowed values: Convenience, Fear, Fun, Financial, Grudge,
	// Ideology, Espionage.
	Motive              param.Field[string]  `json:"motive"`
	OpsecLevel          param.Field[string]  `json:"opsecLevel"`
	OriginCountryISO    param.Field[string]  `json:"originCountryISO"`
	Priority            param.Field[float64] `json:"priority"`
	SophisticationLevel param.Field[string]  `json:"sophisticationLevel"`
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
