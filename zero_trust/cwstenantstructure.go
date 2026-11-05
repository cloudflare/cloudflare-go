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

// CWSTenantStructureService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCWSTenantStructureService] method instead.
type CWSTenantStructureService struct {
	Options []option.RequestOption
	Nodes   *CWSTenantStructureNodeService
}

// NewCWSTenantStructureService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCWSTenantStructureService(opts ...option.RequestOption) (r *CWSTenantStructureService) {
	r = &CWSTenantStructureService{}
	r.Options = opts
	r.Nodes = NewCWSTenantStructureNodeService(opts...)
	return
}

// Creates an empty tenant hierarchy for the selected cloud application.
func (r *CWSTenantStructureService) New(ctx context.Context, params CWSTenantStructureNewParams, opts ...option.RequestOption) (res *TenantStructure, err error) {
	var env CWSTenantStructureNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/tenant-structures", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates the tenant structure's name and import path; its application remains
// immutable.
func (r *CWSTenantStructureService) Update(ctx context.Context, structureID string, params CWSTenantStructureUpdateParams, opts ...option.RequestOption) (res *TenantStructure, err error) {
	var env CWSTenantStructureUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if structureID == "" {
		err = errors.New("missing required structure_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/tenant-structures/%s", params.AccountID, structureID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of tenant structures in the account.
func (r *CWSTenantStructureService) List(ctx context.Context, params CWSTenantStructureListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[TenantStructure], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/tenant-structures", params.AccountID)
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

// Returns a paginated list of tenant structures in the account.
func (r *CWSTenantStructureService) ListAutoPaging(ctx context.Context, params CWSTenantStructureListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[TenantStructure] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes the tenant structure and all nodes in its hierarchy.
func (r *CWSTenantStructureService) Delete(ctx context.Context, structureID string, body CWSTenantStructureDeleteParams, opts ...option.RequestOption) (res *CWSTenantStructureDeleteResponse, err error) {
	var env CWSTenantStructureDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if structureID == "" {
		err = errors.New("missing required structure_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/tenant-structures/%s", body.AccountID, structureID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the tenant structure with the specified ID.
func (r *CWSTenantStructureService) Get(ctx context.Context, structureID string, query CWSTenantStructureGetParams, opts ...option.RequestOption) (res *TenantStructure, err error) {
	var env CWSTenantStructureGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if structureID == "" {
		err = errors.New("missing required structure_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/tenant-structures/%s", query.AccountID, structureID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type TenantStructure struct {
	// Server-assigned UUID
	ID string `json:"id" api:"required" format:"uuid"`
	// Platform this structure models. Immutable after creation; determines valid node
	// kinds and identifier formats.
	Application TenantStructureApplication `json:"application" api:"required"`
	CreatedAt   time.Time                  `json:"created_at" api:"required" format:"date-time"`
	// Human-readable name; unique per account
	Name      string    `json:"name" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Optional unique Rego mount path in the form tenant.<application>.<name>. The
	// application segment must match application; omitted when unmounted.
	ImportPath string              `json:"import_path"`
	JSON       tenantStructureJSON `json:"-"`
}

// tenantStructureJSON contains the JSON metadata for the struct [TenantStructure]
type tenantStructureJSON struct {
	ID          apijson.Field
	Application apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	ImportPath  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantStructure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantStructureJSON) RawJSON() string {
	return r.raw
}

// Platform this structure models. Immutable after creation; determines valid node
// kinds and identifier formats.
type TenantStructureApplication string

const (
	TenantStructureApplicationAws        TenantStructureApplication = "aws"
	TenantStructureApplicationCloudflare TenantStructureApplication = "cloudflare"
)

func (r TenantStructureApplication) IsKnown() bool {
	switch r {
	case TenantStructureApplicationAws, TenantStructureApplicationCloudflare:
		return true
	}
	return false
}

type CWSTenantStructureDeleteResponse struct {
	ID   string                               `json:"id" api:"required" format:"uuid"`
	JSON cwsTenantStructureDeleteResponseJSON `json:"-"`
}

// cwsTenantStructureDeleteResponseJSON contains the JSON metadata for the struct
// [CWSTenantStructureDeleteResponse]
type cwsTenantStructureDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNewParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Platform this structure models. Immutable after creation.
	Application param.Field[CWSTenantStructureNewParamsApplication] `json:"application" api:"required"`
	// Human-readable name; unique per account
	Name param.Field[string] `json:"name" api:"required"`
	// Optional unique Rego mount path in the form tenant.<application>.<name>;
	// application must match. Omit to leave unmounted.
	ImportPath param.Field[string] `json:"import_path"`
}

func (r CWSTenantStructureNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Platform this structure models. Immutable after creation.
type CWSTenantStructureNewParamsApplication string

const (
	CWSTenantStructureNewParamsApplicationAws        CWSTenantStructureNewParamsApplication = "aws"
	CWSTenantStructureNewParamsApplicationCloudflare CWSTenantStructureNewParamsApplication = "cloudflare"
)

func (r CWSTenantStructureNewParamsApplication) IsKnown() bool {
	switch r {
	case CWSTenantStructureNewParamsApplicationAws, CWSTenantStructureNewParamsApplicationCloudflare:
		return true
	}
	return false
}

type CWSTenantStructureNewResponseEnvelope struct {
	Errors   []CWSTenantStructureNewResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSTenantStructureNewResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   TenantStructure                                 `json:"result" api:"required"`
	Success  bool                                            `json:"success" api:"required"`
	JSON     cwsTenantStructureNewResponseEnvelopeJSON       `json:"-"`
}

// cwsTenantStructureNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSTenantStructureNewResponseEnvelope]
type cwsTenantStructureNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	Source  CWSTenantStructureNewResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsTenantStructureNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsTenantStructureNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CWSTenantStructureNewResponseEnvelopeErrors]
type cwsTenantStructureNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer" api:"required"`
	JSON    cwsTenantStructureNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsTenantStructureNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSTenantStructureNewResponseEnvelopeErrorsSource]
type cwsTenantStructureNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code" api:"required"`
	Message string                                              `json:"message" api:"required"`
	Source  CWSTenantStructureNewResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsTenantStructureNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsTenantStructureNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSTenantStructureNewResponseEnvelopeMessages]
type cwsTenantStructureNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer" api:"required"`
	JSON    cwsTenantStructureNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsTenantStructureNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CWSTenantStructureNewResponseEnvelopeMessagesSource]
type cwsTenantStructureNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureUpdateParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Platform this structure models. Must match the existing value.
	Application param.Field[CWSTenantStructureUpdateParamsApplication] `json:"application" api:"required"`
	// Human-readable name; unique per account
	Name param.Field[string] `json:"name" api:"required"`
	// Optional unique Rego mount path in the form tenant.<application>.<name>;
	// application must match. Omit to clear.
	ImportPath param.Field[string] `json:"import_path"`
}

func (r CWSTenantStructureUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Platform this structure models. Must match the existing value.
type CWSTenantStructureUpdateParamsApplication string

const (
	CWSTenantStructureUpdateParamsApplicationAws        CWSTenantStructureUpdateParamsApplication = "aws"
	CWSTenantStructureUpdateParamsApplicationCloudflare CWSTenantStructureUpdateParamsApplication = "cloudflare"
)

func (r CWSTenantStructureUpdateParamsApplication) IsKnown() bool {
	switch r {
	case CWSTenantStructureUpdateParamsApplicationAws, CWSTenantStructureUpdateParamsApplicationCloudflare:
		return true
	}
	return false
}

type CWSTenantStructureUpdateResponseEnvelope struct {
	Errors   []CWSTenantStructureUpdateResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSTenantStructureUpdateResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   TenantStructure                                    `json:"result" api:"required"`
	Success  bool                                               `json:"success" api:"required"`
	JSON     cwsTenantStructureUpdateResponseEnvelopeJSON       `json:"-"`
}

// cwsTenantStructureUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSTenantStructureUpdateResponseEnvelope]
type cwsTenantStructureUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code" api:"required"`
	Message string                                               `json:"message" api:"required"`
	Source  CWSTenantStructureUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsTenantStructureUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsTenantStructureUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CWSTenantStructureUpdateResponseEnvelopeErrors]
type cwsTenantStructureUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer" api:"required"`
	JSON    cwsTenantStructureUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsTenantStructureUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CWSTenantStructureUpdateResponseEnvelopeErrorsSource]
type cwsTenantStructureUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code" api:"required"`
	Message string                                                 `json:"message" api:"required"`
	Source  CWSTenantStructureUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsTenantStructureUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsTenantStructureUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CWSTenantStructureUpdateResponseEnvelopeMessages]
type cwsTenantStructureUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer" api:"required"`
	JSON    cwsTenantStructureUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsTenantStructureUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CWSTenantStructureUpdateResponseEnvelopeMessagesSource]
type cwsTenantStructureUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureListParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Sort direction (default: asc)
	Direction param.Field[CWSTenantStructureListParamsDirection] `query:"direction"`
	// Filter by import path (case-sensitive literal substring match)
	ImportPath param.Field[string] `query:"import_path"`
	// Filter by tenant structure name (case-insensitive substring match)
	Name param.Field[string] `query:"name"`
	// Field to order results by (default: created_at)
	Order param.Field[CWSTenantStructureListParamsOrder] `query:"order"`
	// Number of items per page (1-100, default 20)
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [CWSTenantStructureListParams]'s query parameters as
// `url.Values`.
func (r CWSTenantStructureListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction (default: asc)
type CWSTenantStructureListParamsDirection string

const (
	CWSTenantStructureListParamsDirectionAsc  CWSTenantStructureListParamsDirection = "asc"
	CWSTenantStructureListParamsDirectionDesc CWSTenantStructureListParamsDirection = "desc"
)

func (r CWSTenantStructureListParamsDirection) IsKnown() bool {
	switch r {
	case CWSTenantStructureListParamsDirectionAsc, CWSTenantStructureListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by (default: created_at)
type CWSTenantStructureListParamsOrder string

const (
	CWSTenantStructureListParamsOrderName      CWSTenantStructureListParamsOrder = "name"
	CWSTenantStructureListParamsOrderCreatedAt CWSTenantStructureListParamsOrder = "created_at"
	CWSTenantStructureListParamsOrderUpdatedAt CWSTenantStructureListParamsOrder = "updated_at"
)

func (r CWSTenantStructureListParamsOrder) IsKnown() bool {
	switch r {
	case CWSTenantStructureListParamsOrderName, CWSTenantStructureListParamsOrderCreatedAt, CWSTenantStructureListParamsOrderUpdatedAt:
		return true
	}
	return false
}

type CWSTenantStructureDeleteParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSTenantStructureDeleteResponseEnvelope struct {
	Errors   []CWSTenantStructureDeleteResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSTenantStructureDeleteResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   CWSTenantStructureDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                               `json:"success" api:"required"`
	JSON     cwsTenantStructureDeleteResponseEnvelopeJSON       `json:"-"`
}

// cwsTenantStructureDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSTenantStructureDeleteResponseEnvelope]
type cwsTenantStructureDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code" api:"required"`
	Message string                                               `json:"message" api:"required"`
	Source  CWSTenantStructureDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsTenantStructureDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsTenantStructureDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CWSTenantStructureDeleteResponseEnvelopeErrors]
type cwsTenantStructureDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer" api:"required"`
	JSON    cwsTenantStructureDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsTenantStructureDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CWSTenantStructureDeleteResponseEnvelopeErrorsSource]
type cwsTenantStructureDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code" api:"required"`
	Message string                                                 `json:"message" api:"required"`
	Source  CWSTenantStructureDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsTenantStructureDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsTenantStructureDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CWSTenantStructureDeleteResponseEnvelopeMessages]
type cwsTenantStructureDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer" api:"required"`
	JSON    cwsTenantStructureDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsTenantStructureDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CWSTenantStructureDeleteResponseEnvelopeMessagesSource]
type cwsTenantStructureDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureGetParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSTenantStructureGetResponseEnvelope struct {
	Errors   []CWSTenantStructureGetResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSTenantStructureGetResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   TenantStructure                                 `json:"result" api:"required"`
	Success  bool                                            `json:"success" api:"required"`
	JSON     cwsTenantStructureGetResponseEnvelopeJSON       `json:"-"`
}

// cwsTenantStructureGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSTenantStructureGetResponseEnvelope]
type cwsTenantStructureGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	Source  CWSTenantStructureGetResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsTenantStructureGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsTenantStructureGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CWSTenantStructureGetResponseEnvelopeErrors]
type cwsTenantStructureGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer" api:"required"`
	JSON    cwsTenantStructureGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsTenantStructureGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSTenantStructureGetResponseEnvelopeErrorsSource]
type cwsTenantStructureGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code" api:"required"`
	Message string                                              `json:"message" api:"required"`
	Source  CWSTenantStructureGetResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsTenantStructureGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsTenantStructureGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSTenantStructureGetResponseEnvelopeMessages]
type cwsTenantStructureGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSTenantStructureGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer" api:"required"`
	JSON    cwsTenantStructureGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsTenantStructureGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CWSTenantStructureGetResponseEnvelopeMessagesSource]
type cwsTenantStructureGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSTenantStructureGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsTenantStructureGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
