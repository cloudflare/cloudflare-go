// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_center

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

// InsightClassificationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightClassificationService] method instead.
type InsightClassificationService struct {
	Options []option.RequestOption
}

// NewInsightClassificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInsightClassificationService(opts ...option.RequestOption) (r *InsightClassificationService) {
	r = &InsightClassificationService{}
	r.Options = opts
	return
}

// Updates the user classification for a Security Center insight. Valid values are
// 'false_positive' or 'accept_risk'. To reset, set classification to null. Cannot
// change directly between classification values - must reset to null first.
func (r *InsightClassificationService) Update(ctx context.Context, issueID string, params InsightClassificationUpdateParams, opts ...option.RequestOption) (res *InsightClassificationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if issueID == "" {
		err = errors.New("missing required issue_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/security-center/insights/%s/classification", accountOrZone, accountOrZoneID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

type InsightClassificationUpdateResponse struct {
	Errors   []InsightClassificationUpdateResponseError   `json:"errors" api:"required"`
	Messages []InsightClassificationUpdateResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success InsightClassificationUpdateResponseSuccess `json:"success" api:"required"`
	JSON    insightClassificationUpdateResponseJSON    `json:"-"`
}

// insightClassificationUpdateResponseJSON contains the JSON metadata for the
// struct [InsightClassificationUpdateResponse]
type insightClassificationUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightClassificationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightClassificationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type InsightClassificationUpdateResponseError struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           InsightClassificationUpdateResponseErrorsSource `json:"source"`
	JSON             insightClassificationUpdateResponseErrorJSON    `json:"-"`
}

// insightClassificationUpdateResponseErrorJSON contains the JSON metadata for the
// struct [InsightClassificationUpdateResponseError]
type insightClassificationUpdateResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InsightClassificationUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightClassificationUpdateResponseErrorJSON) RawJSON() string {
	return r.raw
}

type InsightClassificationUpdateResponseErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    insightClassificationUpdateResponseErrorsSourceJSON `json:"-"`
}

// insightClassificationUpdateResponseErrorsSourceJSON contains the JSON metadata
// for the struct [InsightClassificationUpdateResponseErrorsSource]
type insightClassificationUpdateResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightClassificationUpdateResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightClassificationUpdateResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InsightClassificationUpdateResponseMessage struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           InsightClassificationUpdateResponseMessagesSource `json:"source"`
	JSON             insightClassificationUpdateResponseMessageJSON    `json:"-"`
}

// insightClassificationUpdateResponseMessageJSON contains the JSON metadata for
// the struct [InsightClassificationUpdateResponseMessage]
type insightClassificationUpdateResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InsightClassificationUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightClassificationUpdateResponseMessageJSON) RawJSON() string {
	return r.raw
}

type InsightClassificationUpdateResponseMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    insightClassificationUpdateResponseMessagesSourceJSON `json:"-"`
}

// insightClassificationUpdateResponseMessagesSourceJSON contains the JSON metadata
// for the struct [InsightClassificationUpdateResponseMessagesSource]
type insightClassificationUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightClassificationUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightClassificationUpdateResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InsightClassificationUpdateResponseSuccess bool

const (
	InsightClassificationUpdateResponseSuccessTrue InsightClassificationUpdateResponseSuccess = true
)

func (r InsightClassificationUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case InsightClassificationUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type InsightClassificationUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// User-defined classification for the insight. Can be 'false_positive',
	// 'accept_risk', 'other', or null.
	Classification param.Field[InsightClassificationUpdateParamsClassification] `json:"classification"`
	// Rationale for the classification change. Required when classification is
	// 'accept_risk' or 'other'.
	Rationale param.Field[string] `json:"rationale"`
}

func (r InsightClassificationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// User-defined classification for the insight. Can be 'false_positive',
// 'accept_risk', 'other', or null.
type InsightClassificationUpdateParamsClassification string

const (
	InsightClassificationUpdateParamsClassificationFalsePositive InsightClassificationUpdateParamsClassification = "false_positive"
	InsightClassificationUpdateParamsClassificationAcceptRisk    InsightClassificationUpdateParamsClassification = "accept_risk"
	InsightClassificationUpdateParamsClassificationOther         InsightClassificationUpdateParamsClassification = "other"
)

func (r InsightClassificationUpdateParamsClassification) IsKnown() bool {
	switch r {
	case InsightClassificationUpdateParamsClassificationFalsePositive, InsightClassificationUpdateParamsClassificationAcceptRisk, InsightClassificationUpdateParamsClassificationOther:
		return true
	}
	return false
}
