// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package abuse_reports

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// MitigationService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMitigationService] method instead.
type MitigationService struct {
	Options []option.RequestOption
}

// NewMitigationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMitigationService(opts ...option.RequestOption) (r *MitigationService) {
	r = &MitigationService{}
	r.Options = opts
	return
}

// List mitigations done to remediate the abuse report.
func (r *MitigationService) List(ctx context.Context, reportID string, params MitigationListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[MitigationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if reportID == "" {
		err = errors.New("missing required report_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/abuse-reports/%s/mitigations", params.AccountID, reportID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List mitigations done to remediate the abuse report.
func (r *MitigationService) ListAutoPaging(ctx context.Context, reportID string, params MitigationListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[MitigationListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, reportID, params, opts...))
}

// Request a review for mitigations on an account.
func (r *MitigationService) Review(ctx context.Context, reportID string, params MitigationReviewParams, opts ...option.RequestOption) (res *pagination.SinglePage[MitigationReviewResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if reportID == "" {
		err = errors.New("missing required report_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/abuse-reports/%s/mitigations/appeal", params.AccountID, reportID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Request a review for mitigations on an account.
func (r *MitigationService) ReviewAutoPaging(ctx context.Context, reportID string, params MitigationReviewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[MitigationReviewResponse] {
	return pagination.NewSinglePageAutoPager(r.Review(ctx, reportID, params, opts...))
}

type MitigationListResponse struct {
	Mitigations []MitigationListResponseMitigation `json:"mitigations,required"`
	JSON        mitigationListResponseJSON         `json:"-"`
}

// mitigationListResponseJSON contains the JSON metadata for the struct
// [MitigationListResponse]
type mitigationListResponseJSON struct {
	Mitigations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MitigationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mitigationListResponseJSON) RawJSON() string {
	return r.raw
}

type MitigationListResponseMitigation struct {
	// ID of remediation.
	ID string `json:"id,required"`
	// Date when the mitigation will become active. Time in RFC 3339 format
	// (https://www.rfc-editor.org/rfc/rfc3339.html)
	EffectiveDate string                                      `json:"effective_date,required"`
	EntityID      string                                      `json:"entity_id,required"`
	EntityType    MitigationListResponseMitigationsEntityType `json:"entity_type,required"`
	// The status of a mitigation
	Status MitigationListResponseMitigationsStatus `json:"status,required"`
	// The type of mitigation
	Type MitigationListResponseMitigationsType `json:"type,required"`
	JSON mitigationListResponseMitigationJSON  `json:"-"`
}

// mitigationListResponseMitigationJSON contains the JSON metadata for the struct
// [MitigationListResponseMitigation]
type mitigationListResponseMitigationJSON struct {
	ID            apijson.Field
	EffectiveDate apijson.Field
	EntityID      apijson.Field
	EntityType    apijson.Field
	Status        apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MitigationListResponseMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mitigationListResponseMitigationJSON) RawJSON() string {
	return r.raw
}

type MitigationListResponseMitigationsEntityType string

const (
	MitigationListResponseMitigationsEntityTypeURLPattern MitigationListResponseMitigationsEntityType = "url_pattern"
	MitigationListResponseMitigationsEntityTypeAccount    MitigationListResponseMitigationsEntityType = "account"
	MitigationListResponseMitigationsEntityTypeZone       MitigationListResponseMitigationsEntityType = "zone"
)

func (r MitigationListResponseMitigationsEntityType) IsKnown() bool {
	switch r {
	case MitigationListResponseMitigationsEntityTypeURLPattern, MitigationListResponseMitigationsEntityTypeAccount, MitigationListResponseMitigationsEntityTypeZone:
		return true
	}
	return false
}

// The status of a mitigation
type MitigationListResponseMitigationsStatus string

const (
	MitigationListResponseMitigationsStatusPending   MitigationListResponseMitigationsStatus = "pending"
	MitigationListResponseMitigationsStatusActive    MitigationListResponseMitigationsStatus = "active"
	MitigationListResponseMitigationsStatusInReview  MitigationListResponseMitigationsStatus = "in_review"
	MitigationListResponseMitigationsStatusCancelled MitigationListResponseMitigationsStatus = "cancelled"
	MitigationListResponseMitigationsStatusRemoved   MitigationListResponseMitigationsStatus = "removed"
)

func (r MitigationListResponseMitigationsStatus) IsKnown() bool {
	switch r {
	case MitigationListResponseMitigationsStatusPending, MitigationListResponseMitigationsStatusActive, MitigationListResponseMitigationsStatusInReview, MitigationListResponseMitigationsStatusCancelled, MitigationListResponseMitigationsStatusRemoved:
		return true
	}
	return false
}

// The type of mitigation
type MitigationListResponseMitigationsType string

const (
	MitigationListResponseMitigationsTypeLegalBlock           MitigationListResponseMitigationsType = "legal_block"
	MitigationListResponseMitigationsTypePhishingInterstitial MitigationListResponseMitigationsType = "phishing_interstitial"
	MitigationListResponseMitigationsTypeNetworkBlock         MitigationListResponseMitigationsType = "network_block"
	MitigationListResponseMitigationsTypeRateLimitCache       MitigationListResponseMitigationsType = "rate_limit_cache"
	MitigationListResponseMitigationsTypeAccountSuspend       MitigationListResponseMitigationsType = "account_suspend"
	MitigationListResponseMitigationsTypeRedirectVideoStream  MitigationListResponseMitigationsType = "redirect_video_stream"
)

func (r MitigationListResponseMitigationsType) IsKnown() bool {
	switch r {
	case MitigationListResponseMitigationsTypeLegalBlock, MitigationListResponseMitigationsTypePhishingInterstitial, MitigationListResponseMitigationsTypeNetworkBlock, MitigationListResponseMitigationsTypeRateLimitCache, MitigationListResponseMitigationsTypeAccountSuspend, MitigationListResponseMitigationsTypeRedirectVideoStream:
		return true
	}
	return false
}

type MitigationReviewResponse struct {
	// ID of remediation.
	ID string `json:"id,required"`
	// Date when the mitigation will become active. Time in RFC 3339 format
	// (https://www.rfc-editor.org/rfc/rfc3339.html)
	EffectiveDate string                             `json:"effective_date,required"`
	EntityID      string                             `json:"entity_id,required"`
	EntityType    MitigationReviewResponseEntityType `json:"entity_type,required"`
	// The status of a mitigation
	Status MitigationReviewResponseStatus `json:"status,required"`
	// The type of mitigation
	Type MitigationReviewResponseType `json:"type,required"`
	JSON mitigationReviewResponseJSON `json:"-"`
}

// mitigationReviewResponseJSON contains the JSON metadata for the struct
// [MitigationReviewResponse]
type mitigationReviewResponseJSON struct {
	ID            apijson.Field
	EffectiveDate apijson.Field
	EntityID      apijson.Field
	EntityType    apijson.Field
	Status        apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MitigationReviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mitigationReviewResponseJSON) RawJSON() string {
	return r.raw
}

type MitigationReviewResponseEntityType string

const (
	MitigationReviewResponseEntityTypeURLPattern MitigationReviewResponseEntityType = "url_pattern"
	MitigationReviewResponseEntityTypeAccount    MitigationReviewResponseEntityType = "account"
	MitigationReviewResponseEntityTypeZone       MitigationReviewResponseEntityType = "zone"
)

func (r MitigationReviewResponseEntityType) IsKnown() bool {
	switch r {
	case MitigationReviewResponseEntityTypeURLPattern, MitigationReviewResponseEntityTypeAccount, MitigationReviewResponseEntityTypeZone:
		return true
	}
	return false
}

// The status of a mitigation
type MitigationReviewResponseStatus string

const (
	MitigationReviewResponseStatusPending   MitigationReviewResponseStatus = "pending"
	MitigationReviewResponseStatusActive    MitigationReviewResponseStatus = "active"
	MitigationReviewResponseStatusInReview  MitigationReviewResponseStatus = "in_review"
	MitigationReviewResponseStatusCancelled MitigationReviewResponseStatus = "cancelled"
	MitigationReviewResponseStatusRemoved   MitigationReviewResponseStatus = "removed"
)

func (r MitigationReviewResponseStatus) IsKnown() bool {
	switch r {
	case MitigationReviewResponseStatusPending, MitigationReviewResponseStatusActive, MitigationReviewResponseStatusInReview, MitigationReviewResponseStatusCancelled, MitigationReviewResponseStatusRemoved:
		return true
	}
	return false
}

// The type of mitigation
type MitigationReviewResponseType string

const (
	MitigationReviewResponseTypeLegalBlock           MitigationReviewResponseType = "legal_block"
	MitigationReviewResponseTypePhishingInterstitial MitigationReviewResponseType = "phishing_interstitial"
	MitigationReviewResponseTypeNetworkBlock         MitigationReviewResponseType = "network_block"
	MitigationReviewResponseTypeRateLimitCache       MitigationReviewResponseType = "rate_limit_cache"
	MitigationReviewResponseTypeAccountSuspend       MitigationReviewResponseType = "account_suspend"
	MitigationReviewResponseTypeRedirectVideoStream  MitigationReviewResponseType = "redirect_video_stream"
)

func (r MitigationReviewResponseType) IsKnown() bool {
	switch r {
	case MitigationReviewResponseTypeLegalBlock, MitigationReviewResponseTypePhishingInterstitial, MitigationReviewResponseTypeNetworkBlock, MitigationReviewResponseTypeRateLimitCache, MitigationReviewResponseTypeAccountSuspend, MitigationReviewResponseTypeRedirectVideoStream:
		return true
	}
	return false
}

type MitigationListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Returns mitigation that were dispatched after the given date
	EffectiveAfter param.Field[string] `query:"effective_after"`
	// Returns mitigations that were dispatched before the given date
	EffectiveBefore param.Field[string] `query:"effective_before"`
	// Filter by the type of entity the mitigation impacts.
	EntityType param.Field[MitigationListParamsEntityType] `query:"entity_type"`
	// Where in pagination to start listing abuse reports
	Page param.Field[int64] `query:"page"`
	// How many abuse reports per page to list
	PerPage param.Field[int64] `query:"per_page"`
	// A property to sort by, followed by the order
	Sort param.Field[MitigationListParamsSort] `query:"sort"`
	// Filter by the status of the mitigation.
	Status param.Field[MitigationListParamsStatus] `query:"status"`
	// Filter by the type of mitigation. This filter parameter can be specified
	// multiple times to include multiple types of mitigations in the result set, e.g.
	// ?type=rate_limit_cache&type=legal_block.
	Type param.Field[MitigationListParamsType] `query:"type"`
}

// URLQuery serializes [MitigationListParams]'s query parameters as `url.Values`.
func (r MitigationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by the type of entity the mitigation impacts.
type MitigationListParamsEntityType string

const (
	MitigationListParamsEntityTypeURLPattern MitigationListParamsEntityType = "url_pattern"
	MitigationListParamsEntityTypeAccount    MitigationListParamsEntityType = "account"
	MitigationListParamsEntityTypeZone       MitigationListParamsEntityType = "zone"
)

func (r MitigationListParamsEntityType) IsKnown() bool {
	switch r {
	case MitigationListParamsEntityTypeURLPattern, MitigationListParamsEntityTypeAccount, MitigationListParamsEntityTypeZone:
		return true
	}
	return false
}

// A property to sort by, followed by the order
type MitigationListParamsSort string

const (
	MitigationListParamsSortTypeAsc           MitigationListParamsSort = "type,asc"
	MitigationListParamsSortTypeDesc          MitigationListParamsSort = "type,desc"
	MitigationListParamsSortEffectiveDateAsc  MitigationListParamsSort = "effective_date,asc"
	MitigationListParamsSortEffectiveDateDesc MitigationListParamsSort = "effective_date,desc"
	MitigationListParamsSortStatusAsc         MitigationListParamsSort = "status,asc"
	MitigationListParamsSortStatusDesc        MitigationListParamsSort = "status,desc"
	MitigationListParamsSortEntityTypeAsc     MitigationListParamsSort = "entity_type,asc"
	MitigationListParamsSortEntityTypeDesc    MitigationListParamsSort = "entity_type,desc"
)

func (r MitigationListParamsSort) IsKnown() bool {
	switch r {
	case MitigationListParamsSortTypeAsc, MitigationListParamsSortTypeDesc, MitigationListParamsSortEffectiveDateAsc, MitigationListParamsSortEffectiveDateDesc, MitigationListParamsSortStatusAsc, MitigationListParamsSortStatusDesc, MitigationListParamsSortEntityTypeAsc, MitigationListParamsSortEntityTypeDesc:
		return true
	}
	return false
}

// Filter by the status of the mitigation.
type MitigationListParamsStatus string

const (
	MitigationListParamsStatusPending   MitigationListParamsStatus = "pending"
	MitigationListParamsStatusActive    MitigationListParamsStatus = "active"
	MitigationListParamsStatusInReview  MitigationListParamsStatus = "in_review"
	MitigationListParamsStatusCancelled MitigationListParamsStatus = "cancelled"
	MitigationListParamsStatusRemoved   MitigationListParamsStatus = "removed"
)

func (r MitigationListParamsStatus) IsKnown() bool {
	switch r {
	case MitigationListParamsStatusPending, MitigationListParamsStatusActive, MitigationListParamsStatusInReview, MitigationListParamsStatusCancelled, MitigationListParamsStatusRemoved:
		return true
	}
	return false
}

// Filter by the type of mitigation. This filter parameter can be specified
// multiple times to include multiple types of mitigations in the result set, e.g.
// ?type=rate_limit_cache&type=legal_block.
type MitigationListParamsType string

const (
	MitigationListParamsTypeLegalBlock           MitigationListParamsType = "legal_block"
	MitigationListParamsTypePhishingInterstitial MitigationListParamsType = "phishing_interstitial"
	MitigationListParamsTypeNetworkBlock         MitigationListParamsType = "network_block"
	MitigationListParamsTypeRateLimitCache       MitigationListParamsType = "rate_limit_cache"
	MitigationListParamsTypeAccountSuspend       MitigationListParamsType = "account_suspend"
	MitigationListParamsTypeRedirectVideoStream  MitigationListParamsType = "redirect_video_stream"
)

func (r MitigationListParamsType) IsKnown() bool {
	switch r {
	case MitigationListParamsTypeLegalBlock, MitigationListParamsTypePhishingInterstitial, MitigationListParamsTypeNetworkBlock, MitigationListParamsTypeRateLimitCache, MitigationListParamsTypeAccountSuspend, MitigationListParamsTypeRedirectVideoStream:
		return true
	}
	return false
}

type MitigationReviewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// List of mitigations to appeal.
	Appeals param.Field[[]MitigationReviewParamsAppeal] `json:"appeals,required"`
}

func (r MitigationReviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MitigationReviewParamsAppeal struct {
	// ID of the mitigation to appeal.
	ID param.Field[string] `json:"id,required"`
	// Reason why the customer is appealing.
	Reason param.Field[MitigationReviewParamsAppealsReason] `json:"reason,required"`
}

func (r MitigationReviewParamsAppeal) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Reason why the customer is appealing.
type MitigationReviewParamsAppealsReason string

const (
	MitigationReviewParamsAppealsReasonRemoved       MitigationReviewParamsAppealsReason = "removed"
	MitigationReviewParamsAppealsReasonMisclassified MitigationReviewParamsAppealsReason = "misclassified"
)

func (r MitigationReviewParamsAppealsReason) IsKnown() bool {
	switch r {
	case MitigationReviewParamsAppealsReasonRemoved, MitigationReviewParamsAppealsReasonMisclassified:
		return true
	}
	return false
}
