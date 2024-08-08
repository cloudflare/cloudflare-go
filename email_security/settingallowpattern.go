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

// SettingAllowPatternService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingAllowPatternService] method instead.
type SettingAllowPatternService struct {
	Options []option.RequestOption
}

// NewSettingAllowPatternService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingAllowPatternService(opts ...option.RequestOption) (r *SettingAllowPatternService) {
	r = &SettingAllowPatternService{}
	r.Options = opts
	return
}

// Create an email allow pattern
func (r *SettingAllowPatternService) New(ctx context.Context, params SettingAllowPatternNewParams, opts ...option.RequestOption) (res *SettingAllowPatternNewResponseUnion, err error) {
	var env SettingAllowPatternNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_patterns", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, search, and sort an accounts's email allow patterns.
func (r *SettingAllowPatternService) List(ctx context.Context, params SettingAllowPatternListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingAllowPatternListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_patterns", params.AccountID)
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

// List, search, and sort an accounts's email allow patterns.
func (r *SettingAllowPatternService) ListAutoPaging(ctx context.Context, params SettingAllowPatternListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingAllowPatternListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete an email allow pattern
func (r *SettingAllowPatternService) Delete(ctx context.Context, patternID int64, body SettingAllowPatternDeleteParams, opts ...option.RequestOption) (res *SettingAllowPatternDeleteResponse, err error) {
	var env SettingAllowPatternDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_patterns/%v", body.AccountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an email allow pattern
func (r *SettingAllowPatternService) Edit(ctx context.Context, patternID int64, params SettingAllowPatternEditParams, opts ...option.RequestOption) (res *SettingAllowPatternEditResponse, err error) {
	var env SettingAllowPatternEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_patterns/%v", params.AccountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get an email allow pattern
func (r *SettingAllowPatternService) Get(ctx context.Context, patternID int64, query SettingAllowPatternGetParams, opts ...option.RequestOption) (res *SettingAllowPatternGetResponse, err error) {
	var env SettingAllowPatternGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_patterns/%v", query.AccountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [email_security.SettingAllowPatternNewResponseEmailSecurityAllowPattern] or
// [email_security.SettingAllowPatternNewResponseArray].
type SettingAllowPatternNewResponseUnion interface {
	implementsEmailSecuritySettingAllowPatternNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingAllowPatternNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingAllowPatternNewResponseEmailSecurityAllowPattern{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingAllowPatternNewResponseArray{}),
		},
	)
}

type SettingAllowPatternNewResponseEmailSecurityAllowPattern struct {
	ID           int64                                                              `json:"id,required"`
	CreatedAt    time.Time                                                          `json:"created_at,required" format:"date-time"`
	IsRecipient  bool                                                               `json:"is_recipient,required"`
	IsRegex      bool                                                               `json:"is_regex,required"`
	IsSender     bool                                                               `json:"is_sender,required"`
	IsSpoof      bool                                                               `json:"is_spoof,required"`
	LastModified time.Time                                                          `json:"last_modified,required" format:"date-time"`
	Pattern      string                                                             `json:"pattern,required"`
	PatternType  SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternType `json:"pattern_type,required"`
	VerifySender bool                                                               `json:"verify_sender,required"`
	Comments     string                                                             `json:"comments,nullable"`
	JSON         settingAllowPatternNewResponseEmailSecurityAllowPatternJSON        `json:"-"`
}

// settingAllowPatternNewResponseEmailSecurityAllowPatternJSON contains the JSON
// metadata for the struct
// [SettingAllowPatternNewResponseEmailSecurityAllowPattern]
type settingAllowPatternNewResponseEmailSecurityAllowPatternJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRecipient  apijson.Field
	IsRegex      apijson.Field
	IsSender     apijson.Field
	IsSpoof      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	VerifySender apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingAllowPatternNewResponseEmailSecurityAllowPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPatternNewResponseEmailSecurityAllowPatternJSON) RawJSON() string {
	return r.raw
}

func (r SettingAllowPatternNewResponseEmailSecurityAllowPattern) implementsEmailSecuritySettingAllowPatternNewResponseUnion() {
}

type SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternType string

const (
	SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternTypeEmail   SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternType = "EMAIL"
	SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternTypeDomain  SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternType = "DOMAIN"
	SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternTypeIP      SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternType = "IP"
	SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternTypeUnknown SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternType = "UNKNOWN"
)

func (r SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternTypeEmail, SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternTypeDomain, SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternTypeIP, SettingAllowPatternNewResponseEmailSecurityAllowPatternPatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPatternNewResponseArray []SettingAllowPatternNewResponseArrayItem

func (r SettingAllowPatternNewResponseArray) implementsEmailSecuritySettingAllowPatternNewResponseUnion() {
}

type SettingAllowPatternNewResponseArrayItem struct {
	ID           int64                                          `json:"id,required"`
	CreatedAt    time.Time                                      `json:"created_at,required" format:"date-time"`
	IsRecipient  bool                                           `json:"is_recipient,required"`
	IsRegex      bool                                           `json:"is_regex,required"`
	IsSender     bool                                           `json:"is_sender,required"`
	IsSpoof      bool                                           `json:"is_spoof,required"`
	LastModified time.Time                                      `json:"last_modified,required" format:"date-time"`
	Pattern      string                                         `json:"pattern,required"`
	PatternType  SettingAllowPatternNewResponseArrayPatternType `json:"pattern_type,required"`
	VerifySender bool                                           `json:"verify_sender,required"`
	Comments     string                                         `json:"comments,nullable"`
	JSON         settingAllowPatternNewResponseArrayItemJSON    `json:"-"`
}

// settingAllowPatternNewResponseArrayItemJSON contains the JSON metadata for the
// struct [SettingAllowPatternNewResponseArrayItem]
type settingAllowPatternNewResponseArrayItemJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRecipient  apijson.Field
	IsRegex      apijson.Field
	IsSender     apijson.Field
	IsSpoof      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	VerifySender apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingAllowPatternNewResponseArrayItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPatternNewResponseArrayItemJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPatternNewResponseArrayPatternType string

const (
	SettingAllowPatternNewResponseArrayPatternTypeEmail   SettingAllowPatternNewResponseArrayPatternType = "EMAIL"
	SettingAllowPatternNewResponseArrayPatternTypeDomain  SettingAllowPatternNewResponseArrayPatternType = "DOMAIN"
	SettingAllowPatternNewResponseArrayPatternTypeIP      SettingAllowPatternNewResponseArrayPatternType = "IP"
	SettingAllowPatternNewResponseArrayPatternTypeUnknown SettingAllowPatternNewResponseArrayPatternType = "UNKNOWN"
)

func (r SettingAllowPatternNewResponseArrayPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPatternNewResponseArrayPatternTypeEmail, SettingAllowPatternNewResponseArrayPatternTypeDomain, SettingAllowPatternNewResponseArrayPatternTypeIP, SettingAllowPatternNewResponseArrayPatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPatternListResponse struct {
	ID           int64                                      `json:"id,required"`
	CreatedAt    time.Time                                  `json:"created_at,required" format:"date-time"`
	IsRecipient  bool                                       `json:"is_recipient,required"`
	IsRegex      bool                                       `json:"is_regex,required"`
	IsSender     bool                                       `json:"is_sender,required"`
	IsSpoof      bool                                       `json:"is_spoof,required"`
	LastModified time.Time                                  `json:"last_modified,required" format:"date-time"`
	Pattern      string                                     `json:"pattern,required"`
	PatternType  SettingAllowPatternListResponsePatternType `json:"pattern_type,required"`
	VerifySender bool                                       `json:"verify_sender,required"`
	Comments     string                                     `json:"comments,nullable"`
	JSON         settingAllowPatternListResponseJSON        `json:"-"`
}

// settingAllowPatternListResponseJSON contains the JSON metadata for the struct
// [SettingAllowPatternListResponse]
type settingAllowPatternListResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRecipient  apijson.Field
	IsRegex      apijson.Field
	IsSender     apijson.Field
	IsSpoof      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	VerifySender apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingAllowPatternListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPatternListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPatternListResponsePatternType string

const (
	SettingAllowPatternListResponsePatternTypeEmail   SettingAllowPatternListResponsePatternType = "EMAIL"
	SettingAllowPatternListResponsePatternTypeDomain  SettingAllowPatternListResponsePatternType = "DOMAIN"
	SettingAllowPatternListResponsePatternTypeIP      SettingAllowPatternListResponsePatternType = "IP"
	SettingAllowPatternListResponsePatternTypeUnknown SettingAllowPatternListResponsePatternType = "UNKNOWN"
)

func (r SettingAllowPatternListResponsePatternType) IsKnown() bool {
	switch r {
	case SettingAllowPatternListResponsePatternTypeEmail, SettingAllowPatternListResponsePatternTypeDomain, SettingAllowPatternListResponsePatternTypeIP, SettingAllowPatternListResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPatternDeleteResponse struct {
	ID   int64                                 `json:"id,required"`
	JSON settingAllowPatternDeleteResponseJSON `json:"-"`
}

// settingAllowPatternDeleteResponseJSON contains the JSON metadata for the struct
// [SettingAllowPatternDeleteResponse]
type settingAllowPatternDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPatternDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPatternDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPatternEditResponse struct {
	ID           int64                                      `json:"id,required"`
	CreatedAt    time.Time                                  `json:"created_at,required" format:"date-time"`
	IsRecipient  bool                                       `json:"is_recipient,required"`
	IsRegex      bool                                       `json:"is_regex,required"`
	IsSender     bool                                       `json:"is_sender,required"`
	IsSpoof      bool                                       `json:"is_spoof,required"`
	LastModified time.Time                                  `json:"last_modified,required" format:"date-time"`
	Pattern      string                                     `json:"pattern,required"`
	PatternType  SettingAllowPatternEditResponsePatternType `json:"pattern_type,required"`
	VerifySender bool                                       `json:"verify_sender,required"`
	Comments     string                                     `json:"comments,nullable"`
	JSON         settingAllowPatternEditResponseJSON        `json:"-"`
}

// settingAllowPatternEditResponseJSON contains the JSON metadata for the struct
// [SettingAllowPatternEditResponse]
type settingAllowPatternEditResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRecipient  apijson.Field
	IsRegex      apijson.Field
	IsSender     apijson.Field
	IsSpoof      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	VerifySender apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingAllowPatternEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPatternEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPatternEditResponsePatternType string

const (
	SettingAllowPatternEditResponsePatternTypeEmail   SettingAllowPatternEditResponsePatternType = "EMAIL"
	SettingAllowPatternEditResponsePatternTypeDomain  SettingAllowPatternEditResponsePatternType = "DOMAIN"
	SettingAllowPatternEditResponsePatternTypeIP      SettingAllowPatternEditResponsePatternType = "IP"
	SettingAllowPatternEditResponsePatternTypeUnknown SettingAllowPatternEditResponsePatternType = "UNKNOWN"
)

func (r SettingAllowPatternEditResponsePatternType) IsKnown() bool {
	switch r {
	case SettingAllowPatternEditResponsePatternTypeEmail, SettingAllowPatternEditResponsePatternTypeDomain, SettingAllowPatternEditResponsePatternTypeIP, SettingAllowPatternEditResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPatternGetResponse struct {
	ID           int64                                     `json:"id,required"`
	CreatedAt    time.Time                                 `json:"created_at,required" format:"date-time"`
	IsRecipient  bool                                      `json:"is_recipient,required"`
	IsRegex      bool                                      `json:"is_regex,required"`
	IsSender     bool                                      `json:"is_sender,required"`
	IsSpoof      bool                                      `json:"is_spoof,required"`
	LastModified time.Time                                 `json:"last_modified,required" format:"date-time"`
	Pattern      string                                    `json:"pattern,required"`
	PatternType  SettingAllowPatternGetResponsePatternType `json:"pattern_type,required"`
	VerifySender bool                                      `json:"verify_sender,required"`
	Comments     string                                    `json:"comments,nullable"`
	JSON         settingAllowPatternGetResponseJSON        `json:"-"`
}

// settingAllowPatternGetResponseJSON contains the JSON metadata for the struct
// [SettingAllowPatternGetResponse]
type settingAllowPatternGetResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRecipient  apijson.Field
	IsRegex      apijson.Field
	IsSender     apijson.Field
	IsSpoof      apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	VerifySender apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingAllowPatternGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPatternGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPatternGetResponsePatternType string

const (
	SettingAllowPatternGetResponsePatternTypeEmail   SettingAllowPatternGetResponsePatternType = "EMAIL"
	SettingAllowPatternGetResponsePatternTypeDomain  SettingAllowPatternGetResponsePatternType = "DOMAIN"
	SettingAllowPatternGetResponsePatternTypeIP      SettingAllowPatternGetResponsePatternType = "IP"
	SettingAllowPatternGetResponsePatternTypeUnknown SettingAllowPatternGetResponsePatternType = "UNKNOWN"
)

func (r SettingAllowPatternGetResponsePatternType) IsKnown() bool {
	switch r {
	case SettingAllowPatternGetResponsePatternTypeEmail, SettingAllowPatternGetResponsePatternTypeDomain, SettingAllowPatternGetResponsePatternTypeIP, SettingAllowPatternGetResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPatternNewParams struct {
	// Account Identifier
	AccountID param.Field[string]                   `path:"account_id,required"`
	Body      SettingAllowPatternNewParamsBodyUnion `json:"body,required"`
}

func (r SettingAllowPatternNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Satisfied by
// [email_security.SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPattern],
// [email_security.SettingAllowPatternNewParamsBodyArray].
type SettingAllowPatternNewParamsBodyUnion interface {
	implementsEmailSecuritySettingAllowPatternNewParamsBodyUnion()
}

type SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPattern struct {
	IsRecipient  param.Field[bool]                                                                       `json:"is_recipient,required"`
	IsRegex      param.Field[bool]                                                                       `json:"is_regex,required"`
	IsSender     param.Field[bool]                                                                       `json:"is_sender,required"`
	IsSpoof      param.Field[bool]                                                                       `json:"is_spoof,required"`
	Pattern      param.Field[string]                                                                     `json:"pattern,required"`
	PatternType  param.Field[SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternType] `json:"pattern_type,required"`
	VerifySender param.Field[bool]                                                                       `json:"verify_sender,required"`
	Comments     param.Field[string]                                                                     `json:"comments"`
}

func (r SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPattern) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPattern) implementsEmailSecuritySettingAllowPatternNewParamsBodyUnion() {
}

type SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternType string

const (
	SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternTypeEmail   SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternType = "EMAIL"
	SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternTypeDomain  SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternType = "DOMAIN"
	SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternTypeIP      SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternType = "IP"
	SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternTypeUnknown SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternType = "UNKNOWN"
)

func (r SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternTypeEmail, SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternTypeDomain, SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternTypeIP, SettingAllowPatternNewParamsBodyEmailSecurityCreateAllowPatternPatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPatternNewParamsBodyArray []SettingAllowPatternNewParamsBodyArray

func (r SettingAllowPatternNewParamsBodyArray) implementsEmailSecuritySettingAllowPatternNewParamsBodyUnion() {
}

type SettingAllowPatternNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo                      `json:"errors,required"`
	Messages []shared.ResponseInfo                      `json:"messages,required"`
	Result   SettingAllowPatternNewResponseUnion        `json:"result,required"`
	Success  bool                                       `json:"success,required"`
	JSON     settingAllowPatternNewResponseEnvelopeJSON `json:"-"`
}

// settingAllowPatternNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAllowPatternNewResponseEnvelope]
type settingAllowPatternNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPatternNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPatternNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPatternListParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The sorting direction.
	Direction   param.Field[SettingAllowPatternListParamsDirection] `query:"direction"`
	IsRecipient param.Field[bool]                                   `query:"is_recipient"`
	IsSender    param.Field[bool]                                   `query:"is_sender"`
	IsSpoof     param.Field[bool]                                   `query:"is_spoof"`
	// The field to sort by.
	Order param.Field[SettingAllowPatternListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page        param.Field[int64]                                    `query:"page"`
	PatternType param.Field[SettingAllowPatternListParamsPatternType] `query:"pattern_type"`
	// Number of results to display.
	PerPage param.Field[int64] `query:"per_page"`
	// Allows searching in multiple properties of a record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future.
	Search       param.Field[string] `query:"search"`
	VerifySender param.Field[bool]   `query:"verify_sender"`
}

// URLQuery serializes [SettingAllowPatternListParams]'s query parameters as
// `url.Values`.
func (r SettingAllowPatternListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The sorting direction.
type SettingAllowPatternListParamsDirection string

const (
	SettingAllowPatternListParamsDirectionAsc  SettingAllowPatternListParamsDirection = "asc"
	SettingAllowPatternListParamsDirectionDesc SettingAllowPatternListParamsDirection = "desc"
)

func (r SettingAllowPatternListParamsDirection) IsKnown() bool {
	switch r {
	case SettingAllowPatternListParamsDirectionAsc, SettingAllowPatternListParamsDirectionDesc:
		return true
	}
	return false
}

// The field to sort by.
type SettingAllowPatternListParamsOrder string

const (
	SettingAllowPatternListParamsOrderPattern   SettingAllowPatternListParamsOrder = "pattern"
	SettingAllowPatternListParamsOrderCreatedAt SettingAllowPatternListParamsOrder = "created_at"
)

func (r SettingAllowPatternListParamsOrder) IsKnown() bool {
	switch r {
	case SettingAllowPatternListParamsOrderPattern, SettingAllowPatternListParamsOrderCreatedAt:
		return true
	}
	return false
}

type SettingAllowPatternListParamsPatternType string

const (
	SettingAllowPatternListParamsPatternTypeEmail   SettingAllowPatternListParamsPatternType = "EMAIL"
	SettingAllowPatternListParamsPatternTypeDomain  SettingAllowPatternListParamsPatternType = "DOMAIN"
	SettingAllowPatternListParamsPatternTypeIP      SettingAllowPatternListParamsPatternType = "IP"
	SettingAllowPatternListParamsPatternTypeUnknown SettingAllowPatternListParamsPatternType = "UNKNOWN"
)

func (r SettingAllowPatternListParamsPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPatternListParamsPatternTypeEmail, SettingAllowPatternListParamsPatternTypeDomain, SettingAllowPatternListParamsPatternTypeIP, SettingAllowPatternListParamsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPatternDeleteParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingAllowPatternDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                         `json:"errors,required"`
	Messages []shared.ResponseInfo                         `json:"messages,required"`
	Result   SettingAllowPatternDeleteResponse             `json:"result,required"`
	Success  bool                                          `json:"success,required"`
	JSON     settingAllowPatternDeleteResponseEnvelopeJSON `json:"-"`
}

// settingAllowPatternDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAllowPatternDeleteResponseEnvelope]
type settingAllowPatternDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPatternDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPatternDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPatternEditParams struct {
	// Account Identifier
	AccountID    param.Field[string]                                   `path:"account_id,required"`
	Comments     param.Field[string]                                   `json:"comments"`
	IsRecipient  param.Field[bool]                                     `json:"is_recipient"`
	IsRegex      param.Field[bool]                                     `json:"is_regex"`
	IsSender     param.Field[bool]                                     `json:"is_sender"`
	IsSpoof      param.Field[bool]                                     `json:"is_spoof"`
	Pattern      param.Field[string]                                   `json:"pattern"`
	PatternType  param.Field[SettingAllowPatternEditParamsPatternType] `json:"pattern_type"`
	VerifySender param.Field[bool]                                     `json:"verify_sender"`
}

func (r SettingAllowPatternEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingAllowPatternEditParamsPatternType string

const (
	SettingAllowPatternEditParamsPatternTypeEmail   SettingAllowPatternEditParamsPatternType = "EMAIL"
	SettingAllowPatternEditParamsPatternTypeDomain  SettingAllowPatternEditParamsPatternType = "DOMAIN"
	SettingAllowPatternEditParamsPatternTypeIP      SettingAllowPatternEditParamsPatternType = "IP"
	SettingAllowPatternEditParamsPatternTypeUnknown SettingAllowPatternEditParamsPatternType = "UNKNOWN"
)

func (r SettingAllowPatternEditParamsPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPatternEditParamsPatternTypeEmail, SettingAllowPatternEditParamsPatternTypeDomain, SettingAllowPatternEditParamsPatternTypeIP, SettingAllowPatternEditParamsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPatternEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo                       `json:"errors,required"`
	Messages []shared.ResponseInfo                       `json:"messages,required"`
	Result   SettingAllowPatternEditResponse             `json:"result,required"`
	Success  bool                                        `json:"success,required"`
	JSON     settingAllowPatternEditResponseEnvelopeJSON `json:"-"`
}

// settingAllowPatternEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAllowPatternEditResponseEnvelope]
type settingAllowPatternEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPatternEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPatternEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPatternGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingAllowPatternGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                      `json:"errors,required"`
	Messages []shared.ResponseInfo                      `json:"messages,required"`
	Result   SettingAllowPatternGetResponse             `json:"result,required"`
	Success  bool                                       `json:"success,required"`
	JSON     settingAllowPatternGetResponseEnvelopeJSON `json:"-"`
}

// settingAllowPatternGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAllowPatternGetResponseEnvelope]
type settingAllowPatternGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPatternGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPatternGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
