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

// CWSPolicyService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCWSPolicyService] method instead.
type CWSPolicyService struct {
	Options []option.RequestOption
}

// NewCWSPolicyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCWSPolicyService(opts ...option.RequestOption) (r *CWSPolicyService) {
	r = &CWSPolicyService{}
	r.Options = opts
	return
}

// Creates a policy from inline Rego source, an entrypoint, and a filter
// expression.
func (r *CWSPolicyService) New(ctx context.Context, params CWSPolicyNewParams, opts ...option.RequestOption) (res *Policy, err error) {
	var env CWSPolicyNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/policies", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces the policy's configuration. Omit tags or use {} to clear them.
func (r *CWSPolicyService) Update(ctx context.Context, policyID string, params CWSPolicyUpdateParams, opts ...option.RequestOption) (res *Policy, err error) {
	var env CWSPolicyUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/policies/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of policies in the account.
func (r *CWSPolicyService) List(ctx context.Context, params CWSPolicyListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Policy], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/policies", params.AccountID)
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

// Returns a paginated list of policies in the account.
func (r *CWSPolicyService) ListAutoPaging(ctx context.Context, params CWSPolicyListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Policy] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes the policy and cascades its policy-set membership rows. The policy sets
// remain.
func (r *CWSPolicyService) Delete(ctx context.Context, policyID string, body CWSPolicyDeleteParams, opts ...option.RequestOption) (res *CWSPolicyDeleteResponse, err error) {
	var env CWSPolicyDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/policies/%s", body.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the policy with the specified ID.
func (r *CWSPolicyService) Get(ctx context.Context, policyID string, query CWSPolicyGetParams, opts ...option.RequestOption) (res *Policy, err error) {
	var env CWSPolicyGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/policies/%s", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type Policy struct {
	// Server-assigned UUID
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Non-function rule declared in rego_source. Evaluation uses
	// data.<package_path>.<entrypoint>.
	Entrypoint string `json:"entrypoint" api:"required"`
	// Filter expression that determines whether this policy applies. Maximum 4096
	// UTF-8 bytes.
	Filter string `json:"filter" api:"required"`
	// Human-readable name; unique per account
	Name string `json:"name" api:"required"`
	// Dotted path derived from rego_source. Begins with policy. and is unique per
	// account.
	PackagePath string `json:"package_path" api:"required"`
	// Inline Rego source with a package beginning with "policy." Maximum 16384 UTF-8
	// bytes.
	RegoSource string `json:"rego_source" api:"required"`
	// Tags attached to this policy
	Tags      map[string]string `json:"tags" api:"required"`
	UpdatedAt time.Time         `json:"updated_at" api:"required" format:"date-time"`
	// Optional free-form description.
	Description string     `json:"description"`
	JSON        policyJSON `json:"-"`
}

// policyJSON contains the JSON metadata for the struct [Policy]
type policyJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Entrypoint  apijson.Field
	Filter      apijson.Field
	Name        apijson.Field
	PackagePath apijson.Field
	RegoSource  apijson.Field
	Tags        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Policy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyDeleteResponse struct {
	ID   string                      `json:"id" api:"required" format:"uuid"`
	JSON cwsPolicyDeleteResponseJSON `json:"-"`
}

// cwsPolicyDeleteResponseJSON contains the JSON metadata for the struct
// [CWSPolicyDeleteResponse]
type cwsPolicyDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyNewParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Non-function rule declared in rego_source. Evaluation uses
	// data.<package_path>.<entrypoint>.
	Entrypoint param.Field[string] `json:"entrypoint" api:"required"`
	// Filter expression that determines whether this policy applies. Maximum 4096
	// UTF-8 bytes.
	Filter param.Field[string] `json:"filter" api:"required"`
	// Name of the policy
	Name param.Field[string] `json:"name" api:"required"`
	// Inline Rego source with a package beginning with "policy." Maximum 16384 UTF-8
	// bytes.
	RegoSource param.Field[string] `json:"rego_source" api:"required"`
	// Optional free-form description.
	Description param.Field[string] `json:"description"`
	// Tags to attach. Keys are limited to 256 UTF-8 bytes and values to 1024 UTF-8
	// bytes.
	Tags param.Field[map[string]string] `json:"tags"`
}

func (r CWSPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSPolicyNewResponseEnvelope struct {
	Errors   []CWSPolicyNewResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSPolicyNewResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   Policy                                 `json:"result" api:"required"`
	Success  bool                                   `json:"success" api:"required"`
	JSON     cwsPolicyNewResponseEnvelopeJSON       `json:"-"`
}

// cwsPolicyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSPolicyNewResponseEnvelope]
type cwsPolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code" api:"required"`
	Message string                                   `json:"message" api:"required"`
	Source  CWSPolicyNewResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsPolicyNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsPolicyNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CWSPolicyNewResponseEnvelopeErrors]
type cwsPolicyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyNewResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer" api:"required"`
	JSON    cwsPolicyNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsPolicyNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [CWSPolicyNewResponseEnvelopeErrorsSource]
type cwsPolicyNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code" api:"required"`
	Message string                                     `json:"message" api:"required"`
	Source  CWSPolicyNewResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsPolicyNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsPolicyNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSPolicyNewResponseEnvelopeMessages]
type cwsPolicyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyNewResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer" api:"required"`
	JSON    cwsPolicyNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsPolicyNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CWSPolicyNewResponseEnvelopeMessagesSource]
type cwsPolicyNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyUpdateParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Non-function rule declared in rego_source. Evaluation uses
	// data.<package_path>.<entrypoint>.
	Entrypoint param.Field[string] `json:"entrypoint" api:"required"`
	// Filter expression that determines whether this policy applies. Maximum 4096
	// UTF-8 bytes.
	Filter param.Field[string] `json:"filter" api:"required"`
	// Name of the policy
	Name param.Field[string] `json:"name" api:"required"`
	// Inline Rego source with a package beginning with "policy." Maximum 16384 UTF-8
	// bytes.
	RegoSource param.Field[string] `json:"rego_source" api:"required"`
	// Optional free-form description.
	Description param.Field[string] `json:"description"`
	// Replacement tag set. Omit or use {} to clear. Keys are limited to 256 UTF-8
	// bytes and values to 1024 UTF-8 bytes.
	Tags param.Field[map[string]string] `json:"tags"`
}

func (r CWSPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSPolicyUpdateResponseEnvelope struct {
	Errors   []CWSPolicyUpdateResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSPolicyUpdateResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   Policy                                    `json:"result" api:"required"`
	Success  bool                                      `json:"success" api:"required"`
	JSON     cwsPolicyUpdateResponseEnvelopeJSON       `json:"-"`
}

// cwsPolicyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSPolicyUpdateResponseEnvelope]
type cwsPolicyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code" api:"required"`
	Message string                                      `json:"message" api:"required"`
	Source  CWSPolicyUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsPolicyUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsPolicyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSPolicyUpdateResponseEnvelopeErrors]
type cwsPolicyUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    cwsPolicyUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsPolicyUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSPolicyUpdateResponseEnvelopeErrorsSource]
type cwsPolicyUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code" api:"required"`
	Message string                                        `json:"message" api:"required"`
	Source  CWSPolicyUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsPolicyUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsPolicyUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSPolicyUpdateResponseEnvelopeMessages]
type cwsPolicyUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    cwsPolicyUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsPolicyUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CWSPolicyUpdateResponseEnvelopeMessagesSource]
type cwsPolicyUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyListParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Sort direction (default: asc)
	Direction param.Field[CWSPolicyListParamsDirection] `query:"direction"`
	// Filter by policy name (case-insensitive substring match)
	Name param.Field[string] `query:"name"`
	// Field to order results by (default: created_at)
	Order param.Field[CWSPolicyListParamsOrder] `query:"order"`
	// Filter by package path (case-sensitive literal substring match)
	PackagePath param.Field[string] `query:"package_path"`
	// Number of items per page (1-100, default 20)
	PerPage param.Field[int64] `query:"per_page"`
	// Filter to policies attached to (members of) this policy set
	PolicySetID param.Field[string] `query:"policy_set_id" format:"uuid"`
}

// URLQuery serializes [CWSPolicyListParams]'s query parameters as `url.Values`.
func (r CWSPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction (default: asc)
type CWSPolicyListParamsDirection string

const (
	CWSPolicyListParamsDirectionAsc  CWSPolicyListParamsDirection = "asc"
	CWSPolicyListParamsDirectionDesc CWSPolicyListParamsDirection = "desc"
)

func (r CWSPolicyListParamsDirection) IsKnown() bool {
	switch r {
	case CWSPolicyListParamsDirectionAsc, CWSPolicyListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by (default: created_at)
type CWSPolicyListParamsOrder string

const (
	CWSPolicyListParamsOrderName      CWSPolicyListParamsOrder = "name"
	CWSPolicyListParamsOrderCreatedAt CWSPolicyListParamsOrder = "created_at"
	CWSPolicyListParamsOrderUpdatedAt CWSPolicyListParamsOrder = "updated_at"
)

func (r CWSPolicyListParamsOrder) IsKnown() bool {
	switch r {
	case CWSPolicyListParamsOrderName, CWSPolicyListParamsOrderCreatedAt, CWSPolicyListParamsOrderUpdatedAt:
		return true
	}
	return false
}

type CWSPolicyDeleteParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSPolicyDeleteResponseEnvelope struct {
	Errors   []CWSPolicyDeleteResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSPolicyDeleteResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   CWSPolicyDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                      `json:"success" api:"required"`
	JSON     cwsPolicyDeleteResponseEnvelopeJSON       `json:"-"`
}

// cwsPolicyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSPolicyDeleteResponseEnvelope]
type cwsPolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code" api:"required"`
	Message string                                      `json:"message" api:"required"`
	Source  CWSPolicyDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsPolicyDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsPolicyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSPolicyDeleteResponseEnvelopeErrors]
type cwsPolicyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    cwsPolicyDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsPolicyDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSPolicyDeleteResponseEnvelopeErrorsSource]
type cwsPolicyDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code" api:"required"`
	Message string                                        `json:"message" api:"required"`
	Source  CWSPolicyDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsPolicyDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsPolicyDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSPolicyDeleteResponseEnvelopeMessages]
type cwsPolicyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    cwsPolicyDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsPolicyDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CWSPolicyDeleteResponseEnvelopeMessagesSource]
type cwsPolicyDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyGetParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSPolicyGetResponseEnvelope struct {
	Errors   []CWSPolicyGetResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSPolicyGetResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   Policy                                 `json:"result" api:"required"`
	Success  bool                                   `json:"success" api:"required"`
	JSON     cwsPolicyGetResponseEnvelopeJSON       `json:"-"`
}

// cwsPolicyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSPolicyGetResponseEnvelope]
type cwsPolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code" api:"required"`
	Message string                                   `json:"message" api:"required"`
	Source  CWSPolicyGetResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsPolicyGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsPolicyGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CWSPolicyGetResponseEnvelopeErrors]
type cwsPolicyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyGetResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer" api:"required"`
	JSON    cwsPolicyGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsPolicyGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [CWSPolicyGetResponseEnvelopeErrorsSource]
type cwsPolicyGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code" api:"required"`
	Message string                                     `json:"message" api:"required"`
	Source  CWSPolicyGetResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsPolicyGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsPolicyGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSPolicyGetResponseEnvelopeMessages]
type cwsPolicyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSPolicyGetResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer" api:"required"`
	JSON    cwsPolicyGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsPolicyGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CWSPolicyGetResponseEnvelopeMessagesSource]
type cwsPolicyGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSPolicyGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsPolicyGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
