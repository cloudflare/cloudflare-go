// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// SettingBlockSenderService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingBlockSenderService] method instead.
type SettingBlockSenderService struct {
	Options []option.RequestOption
}

// NewSettingBlockSenderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingBlockSenderService(opts ...option.RequestOption) (r *SettingBlockSenderService) {
	r = &SettingBlockSenderService{}
	r.Options = opts
	return
}

// Create a blocked email sender
func (r *SettingBlockSenderService) New(ctx context.Context, params SettingBlockSenderNewParams, opts ...option.RequestOption) (res *SettingBlockSenderNewResponseUnion, err error) {
	var env SettingBlockSenderNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List blocked email senders
func (r *SettingBlockSenderService) List(ctx context.Context, params SettingBlockSenderListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingBlockSenderListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders", params.AccountID)
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

// List blocked email senders
func (r *SettingBlockSenderService) ListAutoPaging(ctx context.Context, params SettingBlockSenderListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingBlockSenderListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a blocked email sender
func (r *SettingBlockSenderService) Delete(ctx context.Context, patternID int64, body SettingBlockSenderDeleteParams, opts ...option.RequestOption) (res *SettingBlockSenderDeleteResponse, err error) {
	var env SettingBlockSenderDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders/%v", body.AccountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a blocked email sender
func (r *SettingBlockSenderService) Edit(ctx context.Context, patternID int64, params SettingBlockSenderEditParams, opts ...option.RequestOption) (res *SettingBlockSenderEditResponse, err error) {
	var env SettingBlockSenderEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders/%v", params.AccountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a blocked email sender
func (r *SettingBlockSenderService) Get(ctx context.Context, patternID int64, query SettingBlockSenderGetParams, opts ...option.RequestOption) (res *SettingBlockSenderGetResponse, err error) {
	var env SettingBlockSenderGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders/%v", query.AccountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [email_security.SettingBlockSenderNewResponseEmailSecurityBlockedSender] or
// [email_security.SettingBlockSenderNewResponseArray].
type SettingBlockSenderNewResponseUnion interface {
	implementsEmailSecuritySettingBlockSenderNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingBlockSenderNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingBlockSenderNewResponseEmailSecurityBlockedSender{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingBlockSenderNewResponseArray{}),
		},
	)
}

type SettingBlockSenderNewResponseEmailSecurityBlockedSender struct {
	ID           int64                                                              `json:"id,required"`
	CreatedAt    time.Time                                                          `json:"created_at,required" format:"date-time"`
	IsRegex      bool                                                               `json:"is_regex,required"`
	LastModified time.Time                                                          `json:"last_modified,required" format:"date-time"`
	Pattern      string                                                             `json:"pattern,required"`
	PatternType  SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternType `json:"pattern_type,required"`
	Comments     string                                                             `json:"comments,nullable"`
	JSON         settingBlockSenderNewResponseEmailSecurityBlockedSenderJSON        `json:"-"`
}

// settingBlockSenderNewResponseEmailSecurityBlockedSenderJSON contains the JSON
// metadata for the struct
// [SettingBlockSenderNewResponseEmailSecurityBlockedSender]
type settingBlockSenderNewResponseEmailSecurityBlockedSenderJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRegex      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingBlockSenderNewResponseEmailSecurityBlockedSender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderNewResponseEmailSecurityBlockedSenderJSON) RawJSON() string {
	return r.raw
}

func (r SettingBlockSenderNewResponseEmailSecurityBlockedSender) implementsEmailSecuritySettingBlockSenderNewResponseUnion() {
}

type SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternType string

const (
	SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternTypeEmail   SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternType = "EMAIL"
	SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternTypeDomain  SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternType = "DOMAIN"
	SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternTypeIP      SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternType = "IP"
	SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternTypeUnknown SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternType = "UNKNOWN"
)

func (r SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternTypeEmail, SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternTypeDomain, SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternTypeIP, SettingBlockSenderNewResponseEmailSecurityBlockedSenderPatternTypeUnknown:
		return true
	}
	return false
}

type SettingBlockSenderNewResponseArray []SettingBlockSenderNewResponseArrayItem

func (r SettingBlockSenderNewResponseArray) implementsEmailSecuritySettingBlockSenderNewResponseUnion() {
}

type SettingBlockSenderNewResponseArrayItem struct {
	ID           int64                                         `json:"id,required"`
	CreatedAt    time.Time                                     `json:"created_at,required" format:"date-time"`
	IsRegex      bool                                          `json:"is_regex,required"`
	LastModified time.Time                                     `json:"last_modified,required" format:"date-time"`
	Pattern      string                                        `json:"pattern,required"`
	PatternType  SettingBlockSenderNewResponseArrayPatternType `json:"pattern_type,required"`
	Comments     string                                        `json:"comments,nullable"`
	JSON         settingBlockSenderNewResponseArrayItemJSON    `json:"-"`
}

// settingBlockSenderNewResponseArrayItemJSON contains the JSON metadata for the
// struct [SettingBlockSenderNewResponseArrayItem]
type settingBlockSenderNewResponseArrayItemJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRegex      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingBlockSenderNewResponseArrayItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderNewResponseArrayItemJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderNewResponseArrayPatternType string

const (
	SettingBlockSenderNewResponseArrayPatternTypeEmail   SettingBlockSenderNewResponseArrayPatternType = "EMAIL"
	SettingBlockSenderNewResponseArrayPatternTypeDomain  SettingBlockSenderNewResponseArrayPatternType = "DOMAIN"
	SettingBlockSenderNewResponseArrayPatternTypeIP      SettingBlockSenderNewResponseArrayPatternType = "IP"
	SettingBlockSenderNewResponseArrayPatternTypeUnknown SettingBlockSenderNewResponseArrayPatternType = "UNKNOWN"
)

func (r SettingBlockSenderNewResponseArrayPatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderNewResponseArrayPatternTypeEmail, SettingBlockSenderNewResponseArrayPatternTypeDomain, SettingBlockSenderNewResponseArrayPatternTypeIP, SettingBlockSenderNewResponseArrayPatternTypeUnknown:
		return true
	}
	return false
}

type SettingBlockSenderListResponse struct {
	ID           int64                                     `json:"id,required"`
	CreatedAt    time.Time                                 `json:"created_at,required" format:"date-time"`
	IsRegex      bool                                      `json:"is_regex,required"`
	LastModified time.Time                                 `json:"last_modified,required" format:"date-time"`
	Pattern      string                                    `json:"pattern,required"`
	PatternType  SettingBlockSenderListResponsePatternType `json:"pattern_type,required"`
	Comments     string                                    `json:"comments,nullable"`
	JSON         settingBlockSenderListResponseJSON        `json:"-"`
}

// settingBlockSenderListResponseJSON contains the JSON metadata for the struct
// [SettingBlockSenderListResponse]
type settingBlockSenderListResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRegex      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingBlockSenderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderListResponsePatternType string

const (
	SettingBlockSenderListResponsePatternTypeEmail   SettingBlockSenderListResponsePatternType = "EMAIL"
	SettingBlockSenderListResponsePatternTypeDomain  SettingBlockSenderListResponsePatternType = "DOMAIN"
	SettingBlockSenderListResponsePatternTypeIP      SettingBlockSenderListResponsePatternType = "IP"
	SettingBlockSenderListResponsePatternTypeUnknown SettingBlockSenderListResponsePatternType = "UNKNOWN"
)

func (r SettingBlockSenderListResponsePatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderListResponsePatternTypeEmail, SettingBlockSenderListResponsePatternTypeDomain, SettingBlockSenderListResponsePatternTypeIP, SettingBlockSenderListResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingBlockSenderDeleteResponse struct {
	ID   int64                                `json:"id,required"`
	JSON settingBlockSenderDeleteResponseJSON `json:"-"`
}

// settingBlockSenderDeleteResponseJSON contains the JSON metadata for the struct
// [SettingBlockSenderDeleteResponse]
type settingBlockSenderDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderEditResponse struct {
	ID           int64                                     `json:"id,required"`
	CreatedAt    time.Time                                 `json:"created_at,required" format:"date-time"`
	IsRegex      bool                                      `json:"is_regex,required"`
	LastModified time.Time                                 `json:"last_modified,required" format:"date-time"`
	Pattern      string                                    `json:"pattern,required"`
	PatternType  SettingBlockSenderEditResponsePatternType `json:"pattern_type,required"`
	Comments     string                                    `json:"comments,nullable"`
	JSON         settingBlockSenderEditResponseJSON        `json:"-"`
}

// settingBlockSenderEditResponseJSON contains the JSON metadata for the struct
// [SettingBlockSenderEditResponse]
type settingBlockSenderEditResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRegex      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingBlockSenderEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderEditResponsePatternType string

const (
	SettingBlockSenderEditResponsePatternTypeEmail   SettingBlockSenderEditResponsePatternType = "EMAIL"
	SettingBlockSenderEditResponsePatternTypeDomain  SettingBlockSenderEditResponsePatternType = "DOMAIN"
	SettingBlockSenderEditResponsePatternTypeIP      SettingBlockSenderEditResponsePatternType = "IP"
	SettingBlockSenderEditResponsePatternTypeUnknown SettingBlockSenderEditResponsePatternType = "UNKNOWN"
)

func (r SettingBlockSenderEditResponsePatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderEditResponsePatternTypeEmail, SettingBlockSenderEditResponsePatternTypeDomain, SettingBlockSenderEditResponsePatternTypeIP, SettingBlockSenderEditResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingBlockSenderGetResponse struct {
	ID           int64                                    `json:"id,required"`
	CreatedAt    time.Time                                `json:"created_at,required" format:"date-time"`
	IsRegex      bool                                     `json:"is_regex,required"`
	LastModified time.Time                                `json:"last_modified,required" format:"date-time"`
	Pattern      string                                   `json:"pattern,required"`
	PatternType  SettingBlockSenderGetResponsePatternType `json:"pattern_type,required"`
	Comments     string                                   `json:"comments,nullable"`
	JSON         settingBlockSenderGetResponseJSON        `json:"-"`
}

// settingBlockSenderGetResponseJSON contains the JSON metadata for the struct
// [SettingBlockSenderGetResponse]
type settingBlockSenderGetResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRegex      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingBlockSenderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderGetResponsePatternType string

const (
	SettingBlockSenderGetResponsePatternTypeEmail   SettingBlockSenderGetResponsePatternType = "EMAIL"
	SettingBlockSenderGetResponsePatternTypeDomain  SettingBlockSenderGetResponsePatternType = "DOMAIN"
	SettingBlockSenderGetResponsePatternTypeIP      SettingBlockSenderGetResponsePatternType = "IP"
	SettingBlockSenderGetResponsePatternTypeUnknown SettingBlockSenderGetResponsePatternType = "UNKNOWN"
)

func (r SettingBlockSenderGetResponsePatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderGetResponsePatternTypeEmail, SettingBlockSenderGetResponsePatternTypeDomain, SettingBlockSenderGetResponsePatternTypeIP, SettingBlockSenderGetResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingBlockSenderNewParams struct {
	// Account Identifier
	AccountID param.Field[string]                  `path:"account_id,required"`
	Body      SettingBlockSenderNewParamsBodyUnion `json:"body,required"`
}

func (r SettingBlockSenderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Satisfied by
// [email_security.SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSender],
// [email_security.SettingBlockSenderNewParamsBodyArray].
type SettingBlockSenderNewParamsBodyUnion interface {
	implementsEmailSecuritySettingBlockSenderNewParamsBodyUnion()
}

type SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSender struct {
	IsRegex     param.Field[bool]                                                                       `json:"is_regex,required"`
	Pattern     param.Field[string]                                                                     `json:"pattern,required"`
	PatternType param.Field[SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternType] `json:"pattern_type,required"`
	Comments    param.Field[string]                                                                     `json:"comments"`
}

func (r SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSender) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSender) implementsEmailSecuritySettingBlockSenderNewParamsBodyUnion() {
}

type SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternType string

const (
	SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternTypeEmail   SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternType = "EMAIL"
	SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternTypeDomain  SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternType = "DOMAIN"
	SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternTypeIP      SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternType = "IP"
	SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternTypeUnknown SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternType = "UNKNOWN"
)

func (r SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternTypeEmail, SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternTypeDomain, SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternTypeIP, SettingBlockSenderNewParamsBodyEmailSecurityCreateBlockedSenderPatternTypeUnknown:
		return true
	}
	return false
}

type SettingBlockSenderNewParamsBodyArray []SettingBlockSenderNewParamsBodyArray

func (r SettingBlockSenderNewParamsBodyArray) implementsEmailSecuritySettingBlockSenderNewParamsBodyUnion() {
}

type SettingBlockSenderNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo                     `json:"errors,required"`
	Messages []shared.ResponseInfo                     `json:"messages,required"`
	Result   SettingBlockSenderNewResponseUnion        `json:"result,required"`
	Success  bool                                      `json:"success,required"`
	JSON     settingBlockSenderNewResponseEnvelopeJSON `json:"-"`
}

// settingBlockSenderNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingBlockSenderNewResponseEnvelope]
type settingBlockSenderNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderListParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The sorting direction.
	Direction param.Field[SettingBlockSenderListParamsDirection] `query:"direction"`
	// The field to sort by.
	Order param.Field[SettingBlockSenderListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page        param.Field[int64]                                   `query:"page"`
	PatternType param.Field[SettingBlockSenderListParamsPatternType] `query:"pattern_type"`
	// Number of results to display.
	PerPage param.Field[int64] `query:"per_page"`
	// Allows searching in multiple properties of a record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [SettingBlockSenderListParams]'s query parameters as
// `url.Values`.
func (r SettingBlockSenderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The sorting direction.
type SettingBlockSenderListParamsDirection string

const (
	SettingBlockSenderListParamsDirectionAsc  SettingBlockSenderListParamsDirection = "asc"
	SettingBlockSenderListParamsDirectionDesc SettingBlockSenderListParamsDirection = "desc"
)

func (r SettingBlockSenderListParamsDirection) IsKnown() bool {
	switch r {
	case SettingBlockSenderListParamsDirectionAsc, SettingBlockSenderListParamsDirectionDesc:
		return true
	}
	return false
}

// The field to sort by.
type SettingBlockSenderListParamsOrder string

const (
	SettingBlockSenderListParamsOrderPattern   SettingBlockSenderListParamsOrder = "pattern"
	SettingBlockSenderListParamsOrderCreatedAt SettingBlockSenderListParamsOrder = "created_at"
)

func (r SettingBlockSenderListParamsOrder) IsKnown() bool {
	switch r {
	case SettingBlockSenderListParamsOrderPattern, SettingBlockSenderListParamsOrderCreatedAt:
		return true
	}
	return false
}

type SettingBlockSenderListParamsPatternType string

const (
	SettingBlockSenderListParamsPatternTypeEmail   SettingBlockSenderListParamsPatternType = "EMAIL"
	SettingBlockSenderListParamsPatternTypeDomain  SettingBlockSenderListParamsPatternType = "DOMAIN"
	SettingBlockSenderListParamsPatternTypeIP      SettingBlockSenderListParamsPatternType = "IP"
	SettingBlockSenderListParamsPatternTypeUnknown SettingBlockSenderListParamsPatternType = "UNKNOWN"
)

func (r SettingBlockSenderListParamsPatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderListParamsPatternTypeEmail, SettingBlockSenderListParamsPatternTypeDomain, SettingBlockSenderListParamsPatternTypeIP, SettingBlockSenderListParamsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingBlockSenderDeleteParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingBlockSenderDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                        `json:"errors,required"`
	Messages []shared.ResponseInfo                        `json:"messages,required"`
	Result   SettingBlockSenderDeleteResponse             `json:"result,required"`
	Success  bool                                         `json:"success,required"`
	JSON     settingBlockSenderDeleteResponseEnvelopeJSON `json:"-"`
}

// settingBlockSenderDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingBlockSenderDeleteResponseEnvelope]
type settingBlockSenderDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderEditParams struct {
	// Account Identifier
	AccountID   param.Field[string]                                  `path:"account_id,required"`
	Comments    param.Field[string]                                  `json:"comments"`
	IsRegex     param.Field[bool]                                    `json:"is_regex"`
	Pattern     param.Field[string]                                  `json:"pattern"`
	PatternType param.Field[SettingBlockSenderEditParamsPatternType] `json:"pattern_type"`
}

func (r SettingBlockSenderEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingBlockSenderEditParamsPatternType string

const (
	SettingBlockSenderEditParamsPatternTypeEmail   SettingBlockSenderEditParamsPatternType = "EMAIL"
	SettingBlockSenderEditParamsPatternTypeDomain  SettingBlockSenderEditParamsPatternType = "DOMAIN"
	SettingBlockSenderEditParamsPatternTypeIP      SettingBlockSenderEditParamsPatternType = "IP"
	SettingBlockSenderEditParamsPatternTypeUnknown SettingBlockSenderEditParamsPatternType = "UNKNOWN"
)

func (r SettingBlockSenderEditParamsPatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderEditParamsPatternTypeEmail, SettingBlockSenderEditParamsPatternTypeDomain, SettingBlockSenderEditParamsPatternTypeIP, SettingBlockSenderEditParamsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingBlockSenderEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo                      `json:"errors,required"`
	Messages []shared.ResponseInfo                      `json:"messages,required"`
	Result   SettingBlockSenderEditResponse             `json:"result,required"`
	Success  bool                                       `json:"success,required"`
	JSON     settingBlockSenderEditResponseEnvelopeJSON `json:"-"`
}

// settingBlockSenderEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingBlockSenderEditResponseEnvelope]
type settingBlockSenderEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingBlockSenderGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                     `json:"errors,required"`
	Messages []shared.ResponseInfo                     `json:"messages,required"`
	Result   SettingBlockSenderGetResponse             `json:"result,required"`
	Success  bool                                      `json:"success,required"`
	JSON     settingBlockSenderGetResponseEnvelopeJSON `json:"-"`
}

// settingBlockSenderGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingBlockSenderGetResponseEnvelope]
type settingBlockSenderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
