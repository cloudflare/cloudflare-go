// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_sending

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// SuppressionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuppressionService] method instead.
type SuppressionService struct {
	Options []option.RequestOption
}

// NewSuppressionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSuppressionService(opts ...option.RequestOption) (r *SuppressionService) {
	r = &SuppressionService{}
	r.Options = opts
	return
}

// Creates an account-wide suppression. If a mutable legacy zone-linked row already
// exists, it is promoted without changing its identifier.
func (r *SuppressionService) New(ctx context.Context, params SuppressionNewParams, opts ...option.RequestOption) (res *SuppressionNewResponse, err error) {
	var env SuppressionNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/suppressions", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists every active Email Sending suppression owned by the account, including
// legacy rows with internal zone memberships.
func (r *SuppressionService) List(ctx context.Context, params SuppressionListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[SuppressionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/suppressions", params.AccountID)
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

// Lists every active Email Sending suppression owned by the account, including
// legacy rows with internal zone memberships.
func (r *SuppressionService) ListAutoPaging(ctx context.Context, params SuppressionListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[SuppressionListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes the suppression, its note, and every legacy internal zone membership,
// allowing future delivery attempts to the address.
func (r *SuppressionService) Delete(ctx context.Context, suppressionID string, body SuppressionDeleteParams, opts ...option.RequestOption) (res *SuppressionDeleteResponse, err error) {
	var env SuppressionDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if suppressionID == "" {
		err = errors.New("missing required suppression_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/suppressions/%s", body.AccountID, suppressionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates expiry or advisory note fields without changing legacy internal zone
// memberships.
func (r *SuppressionService) Edit(ctx context.Context, suppressionID string, params SuppressionEditParams, opts ...option.RequestOption) (res *SuppressionEditResponse, err error) {
	var env SuppressionEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if suppressionID == "" {
		err = errors.New("missing required suppression_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/suppressions/%s", params.AccountID, suppressionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets an Email Sending suppression owned by the account.
func (r *SuppressionService) Get(ctx context.Context, suppressionID string, query SuppressionGetParams, opts ...option.RequestOption) (res *SuppressionGetResponse, err error) {
	var env SuppressionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if suppressionID == "" {
		err = errors.New("missing required suppression_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/suppressions/%s", query.AccountID, suppressionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Imports up to 1,000 account-level Email Sending suppressions in one request.
func (r *SuppressionService) Import(ctx context.Context, params SuppressionImportParams, opts ...option.RequestOption) (res *SuppressionImportResponse, err error) {
	var env SuppressionImportResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/suppressions/bulk", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SuppressionNewResponse struct {
	ID   string                     `json:"id" api:"required" format:"uuid"`
	JSON suppressionNewResponseJSON `json:"-"`
}

// suppressionNewResponseJSON contains the JSON metadata for the struct
// [SuppressionNewResponse]
type suppressionNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionNewResponseJSON) RawJSON() string {
	return r.raw
}

type SuppressionListResponse struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Email     string    `json:"email" api:"required" format:"email"`
	ExpiresAt time.Time `json:"expires_at" api:"required,nullable" format:"date-time"`
	// Whether clients may mutate this suppression. This is determined by the server
	// and must not be inferred from `reason`.
	ReadOnly bool                        `json:"read_only" api:"required"`
	Reason   string                      `json:"reason" api:"required"`
	Note     string                      `json:"note" api:"nullable"`
	JSON     suppressionListResponseJSON `json:"-"`
}

// suppressionListResponseJSON contains the JSON metadata for the struct
// [SuppressionListResponse]
type suppressionListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	ExpiresAt   apijson.Field
	ReadOnly    apijson.Field
	Reason      apijson.Field
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionListResponseJSON) RawJSON() string {
	return r.raw
}

type SuppressionDeleteResponse struct {
	ID   string                        `json:"id" api:"required" format:"uuid"`
	JSON suppressionDeleteResponseJSON `json:"-"`
}

// suppressionDeleteResponseJSON contains the JSON metadata for the struct
// [SuppressionDeleteResponse]
type suppressionDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SuppressionEditResponse struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Email     string    `json:"email" api:"required" format:"email"`
	ExpiresAt time.Time `json:"expires_at" api:"required,nullable" format:"date-time"`
	// Whether clients may mutate this suppression. This is determined by the server
	// and must not be inferred from `reason`.
	ReadOnly bool                        `json:"read_only" api:"required"`
	Reason   string                      `json:"reason" api:"required"`
	Note     string                      `json:"note" api:"nullable"`
	JSON     suppressionEditResponseJSON `json:"-"`
}

// suppressionEditResponseJSON contains the JSON metadata for the struct
// [SuppressionEditResponse]
type suppressionEditResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	ExpiresAt   apijson.Field
	ReadOnly    apijson.Field
	Reason      apijson.Field
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionEditResponseJSON) RawJSON() string {
	return r.raw
}

type SuppressionGetResponse struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Email     string    `json:"email" api:"required" format:"email"`
	ExpiresAt time.Time `json:"expires_at" api:"required,nullable" format:"date-time"`
	// Whether clients may mutate this suppression. This is determined by the server
	// and must not be inferred from `reason`.
	ReadOnly bool                       `json:"read_only" api:"required"`
	Reason   string                     `json:"reason" api:"required"`
	Note     string                     `json:"note" api:"nullable"`
	JSON     suppressionGetResponseJSON `json:"-"`
}

// suppressionGetResponseJSON contains the JSON metadata for the struct
// [SuppressionGetResponse]
type suppressionGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	ExpiresAt   apijson.Field
	ReadOnly    apijson.Field
	Reason      apijson.Field
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionGetResponseJSON) RawJSON() string {
	return r.raw
}

type SuppressionImportResponse struct {
	Deduplicated int64                           `json:"deduplicated" api:"required"`
	Errors       int64                           `json:"errors" api:"required"`
	Invalid      int64                           `json:"invalid" api:"required"`
	Items        []SuppressionImportResponseItem `json:"items" api:"required"`
	Processed    int64                           `json:"processed" api:"required"`
	Skipped      int64                           `json:"skipped" api:"required"`
	Total        int64                           `json:"total" api:"required"`
	JSON         suppressionImportResponseJSON   `json:"-"`
}

// suppressionImportResponseJSON contains the JSON metadata for the struct
// [SuppressionImportResponse]
type suppressionImportResponseJSON struct {
	Deduplicated apijson.Field
	Errors       apijson.Field
	Invalid      apijson.Field
	Items        apijson.Field
	Processed    apijson.Field
	Skipped      apijson.Field
	Total        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SuppressionImportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionImportResponseJSON) RawJSON() string {
	return r.raw
}

type SuppressionImportResponseItem struct {
	Index  int64                                `json:"index" api:"required"`
	Status SuppressionImportResponseItemsStatus `json:"status" api:"required"`
	ID     string                               `json:"id" format:"uuid"`
	Email  string                               `json:"email" format:"email"`
	Error  string                               `json:"error"`
	JSON   suppressionImportResponseItemJSON    `json:"-"`
}

// suppressionImportResponseItemJSON contains the JSON metadata for the struct
// [SuppressionImportResponseItem]
type suppressionImportResponseItemJSON struct {
	Index       apijson.Field
	Status      apijson.Field
	ID          apijson.Field
	Email       apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionImportResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionImportResponseItemJSON) RawJSON() string {
	return r.raw
}

type SuppressionImportResponseItemsStatus string

const (
	SuppressionImportResponseItemsStatusProcessed SuppressionImportResponseItemsStatus = "processed"
	SuppressionImportResponseItemsStatusInvalid   SuppressionImportResponseItemsStatus = "invalid"
	SuppressionImportResponseItemsStatusError     SuppressionImportResponseItemsStatus = "error"
	SuppressionImportResponseItemsStatusSkipped   SuppressionImportResponseItemsStatus = "skipped"
)

func (r SuppressionImportResponseItemsStatus) IsKnown() bool {
	switch r {
	case SuppressionImportResponseItemsStatusProcessed, SuppressionImportResponseItemsStatusInvalid, SuppressionImportResponseItemsStatusError, SuppressionImportResponseItemsStatusSkipped:
		return true
	}
	return false
}

type SuppressionNewParams struct {
	AccountID param.Field[string]    `path:"account_id" api:"required"`
	Email     param.Field[string]    `json:"email" api:"required" format:"email"`
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	Note      param.Field[string]    `json:"note"`
}

func (r SuppressionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuppressionNewResponseEnvelope struct {
	Errors   []interface{}                      `json:"errors" api:"required"`
	Messages []interface{}                      `json:"messages" api:"required"`
	Result   SuppressionNewResponse             `json:"result" api:"required"`
	Success  bool                               `json:"success" api:"required"`
	JSON     suppressionNewResponseEnvelopeJSON `json:"-"`
}

// suppressionNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SuppressionNewResponseEnvelope]
type suppressionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SuppressionListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Opaque pagination cursor returned as `result_info.next_cursor`. It carries the
	// filters that produced it.
	Cursor param.Field[string] `query:"cursor"`
	// Exact email-address filter.
	Email   param.Field[string]                      `query:"email" format:"email"`
	PerPage param.Field[int64]                       `query:"per_page"`
	Reason  param.Field[SuppressionListParamsReason] `query:"reason"`
	// A complete address is an exact match; a value ending in `@` matches that
	// username across every domain. Prefix searches may return short intermediate
	// pages while the bounded account scan advances.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [SuppressionListParams]'s query parameters as `url.Values`.
func (r SuppressionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SuppressionListParamsReason string

const (
	SuppressionListParamsReasonManual     SuppressionListParamsReason = "manual"
	SuppressionListParamsReasonComplaint  SuppressionListParamsReason = "complaint"
	SuppressionListParamsReasonHardBounce SuppressionListParamsReason = "hard_bounce"
	SuppressionListParamsReasonSoftBounce SuppressionListParamsReason = "soft_bounce"
	SuppressionListParamsReasonPolicy     SuppressionListParamsReason = "policy"
)

func (r SuppressionListParamsReason) IsKnown() bool {
	switch r {
	case SuppressionListParamsReasonManual, SuppressionListParamsReasonComplaint, SuppressionListParamsReasonHardBounce, SuppressionListParamsReasonSoftBounce, SuppressionListParamsReasonPolicy:
		return true
	}
	return false
}

type SuppressionDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SuppressionDeleteResponseEnvelope struct {
	Errors   []interface{}                         `json:"errors" api:"required"`
	Messages []interface{}                         `json:"messages" api:"required"`
	Result   SuppressionDeleteResponse             `json:"result" api:"required"`
	Success  bool                                  `json:"success" api:"required"`
	JSON     suppressionDeleteResponseEnvelopeJSON `json:"-"`
}

// suppressionDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SuppressionDeleteResponseEnvelope]
type suppressionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SuppressionEditParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// New expiry. Send `null` to make the suppression permanent; omit to leave it
	// unchanged.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// Replacement advisory note. Send an empty string to clear it; omit to leave it
	// unchanged.
	Note param.Field[string] `json:"note"`
}

func (r SuppressionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuppressionEditResponseEnvelope struct {
	Errors   []interface{}                       `json:"errors" api:"required"`
	Messages []interface{}                       `json:"messages" api:"required"`
	Result   SuppressionEditResponse             `json:"result" api:"required"`
	Success  bool                                `json:"success" api:"required"`
	JSON     suppressionEditResponseEnvelopeJSON `json:"-"`
}

// suppressionEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SuppressionEditResponseEnvelope]
type suppressionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SuppressionGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SuppressionGetResponseEnvelope struct {
	Errors   []interface{}                      `json:"errors" api:"required"`
	Messages []interface{}                      `json:"messages" api:"required"`
	Result   SuppressionGetResponse             `json:"result" api:"required"`
	Success  bool                               `json:"success" api:"required"`
	JSON     suppressionGetResponseEnvelopeJSON `json:"-"`
}

// suppressionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SuppressionGetResponseEnvelope]
type suppressionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SuppressionImportParams struct {
	AccountID param.Field[string]                        `path:"account_id" api:"required"`
	Items     param.Field[[]SuppressionImportParamsItem] `json:"items" api:"required"`
}

func (r SuppressionImportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuppressionImportParamsItem struct {
	Email     param.Field[string]    `json:"email" api:"required"`
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	Note      param.Field[string]    `json:"note"`
}

func (r SuppressionImportParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuppressionImportResponseEnvelope struct {
	Errors   []interface{}                         `json:"errors" api:"required"`
	Messages []interface{}                         `json:"messages" api:"required"`
	Result   SuppressionImportResponse             `json:"result" api:"required"`
	Success  bool                                  `json:"success" api:"required"`
	JSON     suppressionImportResponseEnvelopeJSON `json:"-"`
}

// suppressionImportResponseEnvelopeJSON contains the JSON metadata for the struct
// [SuppressionImportResponseEnvelope]
type suppressionImportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionImportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionImportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
