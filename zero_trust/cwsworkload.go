// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// CWSWorkloadService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCWSWorkloadService] method instead.
type CWSWorkloadService struct {
	Options []option.RequestOption
}

// NewCWSWorkloadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCWSWorkloadService(opts ...option.RequestOption) (r *CWSWorkloadService) {
	r = &CWSWorkloadService{}
	r.Options = opts
	return
}

// Creates a reusable workload selector from a filter expression.
func (r *CWSWorkloadService) New(ctx context.Context, params CWSWorkloadNewParams, opts ...option.RequestOption) (res *Workload, err error) {
	var env CWSWorkloadNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/workloads", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces the workload's configuration. Omit tags or use {} to clear them.
func (r *CWSWorkloadService) Update(ctx context.Context, workloadID string, params CWSWorkloadUpdateParams, opts ...option.RequestOption) (res *Workload, err error) {
	var env CWSWorkloadUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workloadID == "" {
		err = errors.New("missing required workload_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/workloads/%s", params.AccountID, workloadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of workloads in the account.
func (r *CWSWorkloadService) List(ctx context.Context, params CWSWorkloadListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Workload], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/workloads", params.AccountID)
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

// Returns a paginated list of workloads in the account.
func (r *CWSWorkloadService) ListAutoPaging(ctx context.Context, params CWSWorkloadListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Workload] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes the workload if it is not targeted by a policy set.
func (r *CWSWorkloadService) Delete(ctx context.Context, workloadID string, body CWSWorkloadDeleteParams, opts ...option.RequestOption) (res *CWSWorkloadDeleteResponse, err error) {
	var env CWSWorkloadDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workloadID == "" {
		err = errors.New("missing required workload_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/workloads/%s", body.AccountID, workloadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the workload with the specified ID.
func (r *CWSWorkloadService) Get(ctx context.Context, workloadID string, query CWSWorkloadGetParams, opts ...option.RequestOption) (res *Workload, err error) {
	var env CWSWorkloadGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workloadID == "" {
		err = errors.New("missing required workload_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/workloads/%s", query.AccountID, workloadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type Workload struct {
	// Server-assigned UUID
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Filter expression selecting the requests this workload matches. Maximum 4096
	// UTF-8 bytes.
	Filter string `json:"filter" api:"required"`
	// Human-readable handle; unique per account.
	Name string `json:"name" api:"required"`
	// Tags attached to this workload
	Tags      map[string]string `json:"tags" api:"required"`
	UpdatedAt time.Time         `json:"updated_at" api:"required" format:"date-time"`
	// Optional free-form description.
	Description string       `json:"description"`
	JSON        workloadJSON `json:"-"`
}

// workloadJSON contains the JSON metadata for the struct [Workload]
type workloadJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Filter      apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Workload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workloadJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadDeleteResponse struct {
	ID   string                        `json:"id" api:"required" format:"uuid"`
	JSON cwsWorkloadDeleteResponseJSON `json:"-"`
}

// cwsWorkloadDeleteResponseJSON contains the JSON metadata for the struct
// [CWSWorkloadDeleteResponse]
type cwsWorkloadDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadNewParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter expression selecting the requests this workload matches. Maximum 4096
	// UTF-8 bytes.
	Filter param.Field[string] `json:"filter" api:"required"`
	// Human-readable handle; unique per account.
	Name param.Field[string] `json:"name" api:"required"`
	// Optional free-form description.
	Description param.Field[string] `json:"description"`
	// Tags to attach. Keys are limited to 256 UTF-8 bytes and values to 1024 UTF-8
	// bytes.
	Tags param.Field[map[string]string] `json:"tags"`
}

func (r CWSWorkloadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSWorkloadNewResponseEnvelope struct {
	Errors   []CWSWorkloadNewResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSWorkloadNewResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   Workload                                 `json:"result" api:"required"`
	Success  bool                                     `json:"success" api:"required"`
	JSON     cwsWorkloadNewResponseEnvelopeJSON       `json:"-"`
}

// cwsWorkloadNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSWorkloadNewResponseEnvelope]
type cwsWorkloadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code" api:"required"`
	Message string                                     `json:"message" api:"required"`
	Source  CWSWorkloadNewResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsWorkloadNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsWorkloadNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSWorkloadNewResponseEnvelopeErrors]
type cwsWorkloadNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadNewResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer" api:"required"`
	JSON    cwsWorkloadNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsWorkloadNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSWorkloadNewResponseEnvelopeErrorsSource]
type cwsWorkloadNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code" api:"required"`
	Message string                                       `json:"message" api:"required"`
	Source  CWSWorkloadNewResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsWorkloadNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsWorkloadNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSWorkloadNewResponseEnvelopeMessages]
type cwsWorkloadNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadNewResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer" api:"required"`
	JSON    cwsWorkloadNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsWorkloadNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CWSWorkloadNewResponseEnvelopeMessagesSource]
type cwsWorkloadNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadUpdateParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter expression selecting the requests this workload matches. Maximum 4096
	// UTF-8 bytes.
	Filter param.Field[string] `json:"filter" api:"required"`
	// Human-readable handle; unique per account.
	Name param.Field[string] `json:"name" api:"required"`
	// Optional free-form description.
	Description param.Field[string] `json:"description"`
	// Replacement tag set. Omit or use {} to clear. Keys are limited to 256 UTF-8
	// bytes and values to 1024 UTF-8 bytes.
	Tags param.Field[map[string]string] `json:"tags"`
}

func (r CWSWorkloadUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSWorkloadUpdateResponseEnvelope struct {
	Errors   []CWSWorkloadUpdateResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSWorkloadUpdateResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   Workload                                    `json:"result" api:"required"`
	Success  bool                                        `json:"success" api:"required"`
	JSON     cwsWorkloadUpdateResponseEnvelopeJSON       `json:"-"`
}

// cwsWorkloadUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSWorkloadUpdateResponseEnvelope]
type cwsWorkloadUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code" api:"required"`
	Message string                                        `json:"message" api:"required"`
	Source  CWSWorkloadUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsWorkloadUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsWorkloadUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSWorkloadUpdateResponseEnvelopeErrors]
type cwsWorkloadUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    cwsWorkloadUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsWorkloadUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSWorkloadUpdateResponseEnvelopeErrorsSource]
type cwsWorkloadUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	Source  CWSWorkloadUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsWorkloadUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsWorkloadUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSWorkloadUpdateResponseEnvelopeMessages]
type cwsWorkloadUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer" api:"required"`
	JSON    cwsWorkloadUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsWorkloadUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSWorkloadUpdateResponseEnvelopeMessagesSource]
type cwsWorkloadUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadListParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Sort direction (default: asc)
	Direction param.Field[CWSWorkloadListParamsDirection] `query:"direction"`
	// Filter by workload name (case-insensitive substring match)
	Name param.Field[string] `query:"name"`
	// Field to order results by (default: created_at)
	Order param.Field[CWSWorkloadListParamsOrder] `query:"order"`
	// Number of items per page (1-100, default 20)
	PerPage param.Field[int64] `query:"per_page"`
	// Filter to workloads attached to this policy set
	PolicySetID param.Field[string] `query:"policy_set_id" format:"uuid"`
}

// URLQuery serializes [CWSWorkloadListParams]'s query parameters as `url.Values`.
func (r CWSWorkloadListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction (default: asc)
type CWSWorkloadListParamsDirection string

const (
	CWSWorkloadListParamsDirectionAsc  CWSWorkloadListParamsDirection = "asc"
	CWSWorkloadListParamsDirectionDesc CWSWorkloadListParamsDirection = "desc"
)

func (r CWSWorkloadListParamsDirection) IsKnown() bool {
	switch r {
	case CWSWorkloadListParamsDirectionAsc, CWSWorkloadListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by (default: created_at)
type CWSWorkloadListParamsOrder string

const (
	CWSWorkloadListParamsOrderName      CWSWorkloadListParamsOrder = "name"
	CWSWorkloadListParamsOrderCreatedAt CWSWorkloadListParamsOrder = "created_at"
	CWSWorkloadListParamsOrderUpdatedAt CWSWorkloadListParamsOrder = "updated_at"
)

func (r CWSWorkloadListParamsOrder) IsKnown() bool {
	switch r {
	case CWSWorkloadListParamsOrderName, CWSWorkloadListParamsOrderCreatedAt, CWSWorkloadListParamsOrderUpdatedAt:
		return true
	}
	return false
}

type CWSWorkloadDeleteParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSWorkloadDeleteResponseEnvelope struct {
	Errors   []CWSWorkloadDeleteResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSWorkloadDeleteResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   CWSWorkloadDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                        `json:"success" api:"required"`
	JSON     cwsWorkloadDeleteResponseEnvelopeJSON       `json:"-"`
}

// cwsWorkloadDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSWorkloadDeleteResponseEnvelope]
type cwsWorkloadDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadDeleteResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code" api:"required"`
	Message string                                        `json:"message" api:"required"`
	Source  CWSWorkloadDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsWorkloadDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsWorkloadDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSWorkloadDeleteResponseEnvelopeErrors]
type cwsWorkloadDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    cwsWorkloadDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsWorkloadDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSWorkloadDeleteResponseEnvelopeErrorsSource]
type cwsWorkloadDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadDeleteResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	Source  CWSWorkloadDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsWorkloadDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsWorkloadDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSWorkloadDeleteResponseEnvelopeMessages]
type cwsWorkloadDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer" api:"required"`
	JSON    cwsWorkloadDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsWorkloadDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSWorkloadDeleteResponseEnvelopeMessagesSource]
type cwsWorkloadDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadGetParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSWorkloadGetResponseEnvelope struct {
	Errors   []CWSWorkloadGetResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSWorkloadGetResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   Workload                                 `json:"result" api:"required"`
	Success  bool                                     `json:"success" api:"required"`
	JSON     cwsWorkloadGetResponseEnvelopeJSON       `json:"-"`
}

// cwsWorkloadGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSWorkloadGetResponseEnvelope]
type cwsWorkloadGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code" api:"required"`
	Message string                                     `json:"message" api:"required"`
	Source  CWSWorkloadGetResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsWorkloadGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsWorkloadGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSWorkloadGetResponseEnvelopeErrors]
type cwsWorkloadGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadGetResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer" api:"required"`
	JSON    cwsWorkloadGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsWorkloadGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSWorkloadGetResponseEnvelopeErrorsSource]
type cwsWorkloadGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code" api:"required"`
	Message string                                       `json:"message" api:"required"`
	Source  CWSWorkloadGetResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsWorkloadGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsWorkloadGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSWorkloadGetResponseEnvelopeMessages]
type cwsWorkloadGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSWorkloadGetResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer" api:"required"`
	JSON    cwsWorkloadGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsWorkloadGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CWSWorkloadGetResponseEnvelopeMessagesSource]
type cwsWorkloadGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSWorkloadGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsWorkloadGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
