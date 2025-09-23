// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
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
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/create", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ThreatEventTagNewResponse struct {
	UUID                    string                        `json:"uuid,required"`
	Value                   string                        `json:"value,required"`
	ActiveDuration          string                        `json:"activeDuration"`
	ActorCategory           string                        `json:"actorCategory"`
	AliasGroupNames         []string                      `json:"aliasGroupNames"`
	AliasGroupNamesInternal []string                      `json:"aliasGroupNamesInternal"`
	AnalyticPriority        float64                       `json:"analyticPriority"`
	AttributionConfidence   string                        `json:"attributionConfidence"`
	AttributionOrganization string                        `json:"attributionOrganization"`
	CategoryName            string                        `json:"categoryName"`
	ExternalReferenceLinks  []string                      `json:"externalReferenceLinks"`
	InternalDescription     string                        `json:"internalDescription"`
	Motive                  string                        `json:"motive"`
	OpsecLevel              string                        `json:"opsecLevel"`
	OriginCountryISO        string                        `json:"originCountryISO"`
	Priority                float64                       `json:"priority"`
	SophisticationLevel     string                        `json:"sophisticationLevel"`
	JSON                    threatEventTagNewResponseJSON `json:"-"`
}

// threatEventTagNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagNewResponse]
type threatEventTagNewResponseJSON struct {
	UUID                    apijson.Field
	Value                   apijson.Field
	ActiveDuration          apijson.Field
	ActorCategory           apijson.Field
	AliasGroupNames         apijson.Field
	AliasGroupNamesInternal apijson.Field
	AnalyticPriority        apijson.Field
	AttributionConfidence   apijson.Field
	AttributionOrganization apijson.Field
	CategoryName            apijson.Field
	ExternalReferenceLinks  apijson.Field
	InternalDescription     apijson.Field
	Motive                  apijson.Field
	OpsecLevel              apijson.Field
	OriginCountryISO        apijson.Field
	Priority                apijson.Field
	SophisticationLevel     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ThreatEventTagNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagNewParams struct {
	// Account ID.
	AccountID               param.Field[string]   `path:"account_id,required"`
	Value                   param.Field[string]   `json:"value,required"`
	ActiveDuration          param.Field[string]   `json:"activeDuration"`
	ActorCategory           param.Field[string]   `json:"actorCategory"`
	AliasGroupNames         param.Field[[]string] `json:"aliasGroupNames"`
	AliasGroupNamesInternal param.Field[[]string] `json:"aliasGroupNamesInternal"`
	AnalyticPriority        param.Field[float64]  `json:"analyticPriority"`
	AttributionConfidence   param.Field[string]   `json:"attributionConfidence"`
	AttributionOrganization param.Field[string]   `json:"attributionOrganization"`
	CategoryUUID            param.Field[string]   `json:"categoryUuid"`
	ExternalReferenceLinks  param.Field[[]string] `json:"externalReferenceLinks"`
	InternalDescription     param.Field[string]   `json:"internalDescription"`
	Motive                  param.Field[string]   `json:"motive"`
	OpsecLevel              param.Field[string]   `json:"opsecLevel"`
	OriginCountryISO        param.Field[string]   `json:"originCountryISO"`
	Priority                param.Field[float64]  `json:"priority"`
	SophisticationLevel     param.Field[string]   `json:"sophisticationLevel"`
}

func (r ThreatEventTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
