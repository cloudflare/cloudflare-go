// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package abuse_reports

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
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
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
		return nil, err
	}
	if reportID == "" {
		err = errors.New("missing required report_id parameter")
		return nil, err
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
		return nil, err
	}
	if reportID == "" {
		err = errors.New("missing required report_id parameter")
		return nil, err
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
	Mitigations []MitigationListResponseMitigation `json:"mitigations" api:"required"`
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
	ID string `json:"id" api:"required"`
	// Date when the mitigation will become active. Time in RFC 3339 format
	// (https://www.rfc-editor.org/rfc/rfc3339.html)
	EffectiveDate string `json:"effective_date" api:"required"`
	EntityID      string `json:"entity_id" api:"required"`
	// The type of entity targeted by a mitigation.
	EntityType MitigationListResponseMitigationsEntityType `json:"entity_type" api:"required"`
	// The status of a mitigation
	Status MitigationListResponseMitigationsStatus `json:"status" api:"required"`
	// The type of mitigation applied to a reported entity.
	Type MitigationListResponseMitigationsType `json:"type" api:"required"`
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

// The type of entity targeted by a mitigation.
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

// The type of mitigation applied to a reported entity.
type MitigationListResponseMitigationsType string

const (
	MitigationListResponseMitigationsTypeAccountSuspend          MitigationListResponseMitigationsType = "account_suspend"
	MitigationListResponseMitigationsTypeCopyrightInterstitial   MitigationListResponseMitigationsType = "copyright_interstitial"
	MitigationListResponseMitigationsTypeGeoBlock                MitigationListResponseMitigationsType = "geo_block"
	MitigationListResponseMitigationsTypeLegalBlock              MitigationListResponseMitigationsType = "legal_block"
	MitigationListResponseMitigationsTypeMalwareInterstitial     MitigationListResponseMitigationsType = "malware_interstitial"
	MitigationListResponseMitigationsTypeMisleadingInterstitial  MitigationListResponseMitigationsType = "misleading_interstitial"
	MitigationListResponseMitigationsTypeNetworkBlock            MitigationListResponseMitigationsType = "network_block"
	MitigationListResponseMitigationsTypePhishingInterstitial    MitigationListResponseMitigationsType = "phishing_interstitial"
	MitigationListResponseMitigationsTypePlayfairiteEnforce      MitigationListResponseMitigationsType = "playfairite_enforce"
	MitigationListResponseMitigationsTypeR2TakedownAccount       MitigationListResponseMitigationsType = "r2_takedown_account"
	MitigationListResponseMitigationsTypeR2TakedownBucket        MitigationListResponseMitigationsType = "r2_takedown_bucket"
	MitigationListResponseMitigationsTypeR2TakedownObject        MitigationListResponseMitigationsType = "r2_takedown_object"
	MitigationListResponseMitigationsTypeRateLimitCache          MitigationListResponseMitigationsType = "rate_limit_cache"
	MitigationListResponseMitigationsTypeRedirectVideoStream     MitigationListResponseMitigationsType = "redirect_video_stream"
	MitigationListResponseMitigationsTypeRegistrarFreeze         MitigationListResponseMitigationsType = "registrar_freeze"
	MitigationListResponseMitigationsTypeRegistrarParking        MitigationListResponseMitigationsType = "registrar_parking"
	MitigationListResponseMitigationsTypeStreamBlockAccount      MitigationListResponseMitigationsType = "stream_block_account"
	MitigationListResponseMitigationsTypeUserSuspend             MitigationListResponseMitigationsType = "user_suspend"
	MitigationListResponseMitigationsTypeWorkersTakedownByZoneID MitigationListResponseMitigationsType = "workers_takedown_by_zone_id"
)

func (r MitigationListResponseMitigationsType) IsKnown() bool {
	switch r {
	case MitigationListResponseMitigationsTypeAccountSuspend, MitigationListResponseMitigationsTypeCopyrightInterstitial, MitigationListResponseMitigationsTypeGeoBlock, MitigationListResponseMitigationsTypeLegalBlock, MitigationListResponseMitigationsTypeMalwareInterstitial, MitigationListResponseMitigationsTypeMisleadingInterstitial, MitigationListResponseMitigationsTypeNetworkBlock, MitigationListResponseMitigationsTypePhishingInterstitial, MitigationListResponseMitigationsTypePlayfairiteEnforce, MitigationListResponseMitigationsTypeR2TakedownAccount, MitigationListResponseMitigationsTypeR2TakedownBucket, MitigationListResponseMitigationsTypeR2TakedownObject, MitigationListResponseMitigationsTypeRateLimitCache, MitigationListResponseMitigationsTypeRedirectVideoStream, MitigationListResponseMitigationsTypeRegistrarFreeze, MitigationListResponseMitigationsTypeRegistrarParking, MitigationListResponseMitigationsTypeStreamBlockAccount, MitigationListResponseMitigationsTypeUserSuspend, MitigationListResponseMitigationsTypeWorkersTakedownByZoneID:
		return true
	}
	return false
}

type MitigationReviewResponse struct {
	// ID of remediation.
	ID string `json:"id" api:"required"`
	// Date when the mitigation will become active. Time in RFC 3339 format
	// (https://www.rfc-editor.org/rfc/rfc3339.html)
	EffectiveDate string `json:"effective_date" api:"required"`
	EntityID      string `json:"entity_id" api:"required"`
	// The type of entity targeted by a mitigation.
	EntityType MitigationReviewResponseEntityType `json:"entity_type" api:"required"`
	// The status of a mitigation
	Status MitigationReviewResponseStatus `json:"status" api:"required"`
	// The type of mitigation applied to a reported entity.
	Type MitigationReviewResponseType `json:"type" api:"required"`
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

// The type of entity targeted by a mitigation.
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

// The type of mitigation applied to a reported entity.
type MitigationReviewResponseType string

const (
	MitigationReviewResponseTypeAccountSuspend          MitigationReviewResponseType = "account_suspend"
	MitigationReviewResponseTypeCopyrightInterstitial   MitigationReviewResponseType = "copyright_interstitial"
	MitigationReviewResponseTypeGeoBlock                MitigationReviewResponseType = "geo_block"
	MitigationReviewResponseTypeLegalBlock              MitigationReviewResponseType = "legal_block"
	MitigationReviewResponseTypeMalwareInterstitial     MitigationReviewResponseType = "malware_interstitial"
	MitigationReviewResponseTypeMisleadingInterstitial  MitigationReviewResponseType = "misleading_interstitial"
	MitigationReviewResponseTypeNetworkBlock            MitigationReviewResponseType = "network_block"
	MitigationReviewResponseTypePhishingInterstitial    MitigationReviewResponseType = "phishing_interstitial"
	MitigationReviewResponseTypePlayfairiteEnforce      MitigationReviewResponseType = "playfairite_enforce"
	MitigationReviewResponseTypeR2TakedownAccount       MitigationReviewResponseType = "r2_takedown_account"
	MitigationReviewResponseTypeR2TakedownBucket        MitigationReviewResponseType = "r2_takedown_bucket"
	MitigationReviewResponseTypeR2TakedownObject        MitigationReviewResponseType = "r2_takedown_object"
	MitigationReviewResponseTypeRateLimitCache          MitigationReviewResponseType = "rate_limit_cache"
	MitigationReviewResponseTypeRedirectVideoStream     MitigationReviewResponseType = "redirect_video_stream"
	MitigationReviewResponseTypeRegistrarFreeze         MitigationReviewResponseType = "registrar_freeze"
	MitigationReviewResponseTypeRegistrarParking        MitigationReviewResponseType = "registrar_parking"
	MitigationReviewResponseTypeStreamBlockAccount      MitigationReviewResponseType = "stream_block_account"
	MitigationReviewResponseTypeUserSuspend             MitigationReviewResponseType = "user_suspend"
	MitigationReviewResponseTypeWorkersTakedownByZoneID MitigationReviewResponseType = "workers_takedown_by_zone_id"
)

func (r MitigationReviewResponseType) IsKnown() bool {
	switch r {
	case MitigationReviewResponseTypeAccountSuspend, MitigationReviewResponseTypeCopyrightInterstitial, MitigationReviewResponseTypeGeoBlock, MitigationReviewResponseTypeLegalBlock, MitigationReviewResponseTypeMalwareInterstitial, MitigationReviewResponseTypeMisleadingInterstitial, MitigationReviewResponseTypeNetworkBlock, MitigationReviewResponseTypePhishingInterstitial, MitigationReviewResponseTypePlayfairiteEnforce, MitigationReviewResponseTypeR2TakedownAccount, MitigationReviewResponseTypeR2TakedownBucket, MitigationReviewResponseTypeR2TakedownObject, MitigationReviewResponseTypeRateLimitCache, MitigationReviewResponseTypeRedirectVideoStream, MitigationReviewResponseTypeRegistrarFreeze, MitigationReviewResponseTypeRegistrarParking, MitigationReviewResponseTypeStreamBlockAccount, MitigationReviewResponseTypeUserSuspend, MitigationReviewResponseTypeWorkersTakedownByZoneID:
		return true
	}
	return false
}

type MitigationListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
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
	MitigationListParamsTypeAccountSuspend          MitigationListParamsType = "account_suspend"
	MitigationListParamsTypeCopyrightInterstitial   MitigationListParamsType = "copyright_interstitial"
	MitigationListParamsTypeGeoBlock                MitigationListParamsType = "geo_block"
	MitigationListParamsTypeLegalBlock              MitigationListParamsType = "legal_block"
	MitigationListParamsTypeMalwareInterstitial     MitigationListParamsType = "malware_interstitial"
	MitigationListParamsTypeMisleadingInterstitial  MitigationListParamsType = "misleading_interstitial"
	MitigationListParamsTypeNetworkBlock            MitigationListParamsType = "network_block"
	MitigationListParamsTypePhishingInterstitial    MitigationListParamsType = "phishing_interstitial"
	MitigationListParamsTypePlayfairiteEnforce      MitigationListParamsType = "playfairite_enforce"
	MitigationListParamsTypeR2TakedownAccount       MitigationListParamsType = "r2_takedown_account"
	MitigationListParamsTypeR2TakedownBucket        MitigationListParamsType = "r2_takedown_bucket"
	MitigationListParamsTypeR2TakedownObject        MitigationListParamsType = "r2_takedown_object"
	MitigationListParamsTypeRateLimitCache          MitigationListParamsType = "rate_limit_cache"
	MitigationListParamsTypeRedirectVideoStream     MitigationListParamsType = "redirect_video_stream"
	MitigationListParamsTypeRegistrarFreeze         MitigationListParamsType = "registrar_freeze"
	MitigationListParamsTypeRegistrarParking        MitigationListParamsType = "registrar_parking"
	MitigationListParamsTypeStreamBlockAccount      MitigationListParamsType = "stream_block_account"
	MitigationListParamsTypeUserSuspend             MitigationListParamsType = "user_suspend"
	MitigationListParamsTypeWorkersTakedownByZoneID MitigationListParamsType = "workers_takedown_by_zone_id"
)

func (r MitigationListParamsType) IsKnown() bool {
	switch r {
	case MitigationListParamsTypeAccountSuspend, MitigationListParamsTypeCopyrightInterstitial, MitigationListParamsTypeGeoBlock, MitigationListParamsTypeLegalBlock, MitigationListParamsTypeMalwareInterstitial, MitigationListParamsTypeMisleadingInterstitial, MitigationListParamsTypeNetworkBlock, MitigationListParamsTypePhishingInterstitial, MitigationListParamsTypePlayfairiteEnforce, MitigationListParamsTypeR2TakedownAccount, MitigationListParamsTypeR2TakedownBucket, MitigationListParamsTypeR2TakedownObject, MitigationListParamsTypeRateLimitCache, MitigationListParamsTypeRedirectVideoStream, MitigationListParamsTypeRegistrarFreeze, MitigationListParamsTypeRegistrarParking, MitigationListParamsTypeStreamBlockAccount, MitigationListParamsTypeUserSuspend, MitigationListParamsTypeWorkersTakedownByZoneID:
		return true
	}
	return false
}

type MitigationReviewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// List of mitigations to appeal.
	Appeals param.Field[[]MitigationReviewParamsAppeal] `json:"appeals" api:"required"`
}

func (r MitigationReviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MitigationReviewParamsAppeal struct {
	// ID of the mitigation to appeal.
	ID param.Field[string] `json:"id" api:"required"`
	// Reason why the customer is appealing.
	Reason param.Field[MitigationReviewParamsAppealsReason] `json:"reason" api:"required"`
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
