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

// CWSBasicAuthCredentialService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCWSBasicAuthCredentialService] method instead.
type CWSBasicAuthCredentialService struct {
	Options []option.RequestOption
	Secrets *CWSBasicAuthCredentialSecretService
}

// NewCWSBasicAuthCredentialService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCWSBasicAuthCredentialService(opts ...option.RequestOption) (r *CWSBasicAuthCredentialService) {
	r = &CWSBasicAuthCredentialService{}
	r.Options = opts
	r.Secrets = NewCWSBasicAuthCredentialSecretService(opts...)
	return
}

// Creates a basic-auth credential with the specified username.
func (r *CWSBasicAuthCredentialService) New(ctx context.Context, params CWSBasicAuthCredentialNewParams, opts ...option.RequestOption) (res *BasicAuthCredential, err error) {
	var env CWSBasicAuthCredentialNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/basic-auth-credentials", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of basic-auth credentials in the account.
func (r *CWSBasicAuthCredentialService) List(ctx context.Context, params CWSBasicAuthCredentialListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[BasicAuthCredential], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/basic-auth-credentials", params.AccountID)
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

// Returns a paginated list of basic-auth credentials in the account.
func (r *CWSBasicAuthCredentialService) ListAutoPaging(ctx context.Context, params CWSBasicAuthCredentialListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[BasicAuthCredential] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes the credential and its secrets if it is not attached to a deployment.
func (r *CWSBasicAuthCredentialService) Delete(ctx context.Context, credentialID string, body CWSBasicAuthCredentialDeleteParams, opts ...option.RequestOption) (res *CWSBasicAuthCredentialDeleteResponse, err error) {
	var env CWSBasicAuthCredentialDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if credentialID == "" {
		err = errors.New("missing required credential_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/basic-auth-credentials/%s", body.AccountID, credentialID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the basic-auth credential with the specified ID.
func (r *CWSBasicAuthCredentialService) Get(ctx context.Context, credentialID string, query CWSBasicAuthCredentialGetParams, opts ...option.RequestOption) (res *BasicAuthCredential, err error) {
	var env CWSBasicAuthCredentialGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if credentialID == "" {
		err = errors.New("missing required credential_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/basic-auth-credentials/%s", query.AccountID, credentialID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type BasicAuthCredential struct {
	// Server-assigned UUID
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Login presented by the client. Printable ASCII only, excluding space and colon
	// (RFC 7617 forbids a colon in the user-id).
	Username string                  `json:"username" api:"required"`
	JSON     basicAuthCredentialJSON `json:"-"`
}

// basicAuthCredentialJSON contains the JSON metadata for the struct
// [BasicAuthCredential]
type basicAuthCredentialJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasicAuthCredential) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basicAuthCredentialJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialDeleteResponse struct {
	ID   string                                   `json:"id" api:"required" format:"uuid"`
	JSON cwsBasicAuthCredentialDeleteResponseJSON `json:"-"`
}

// cwsBasicAuthCredentialDeleteResponseJSON contains the JSON metadata for the
// struct [CWSBasicAuthCredentialDeleteResponse]
type cwsBasicAuthCredentialDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialNewParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Login presented by the client. Printable ASCII only, excluding space and colon.
	Username param.Field[string] `json:"username" api:"required"`
}

func (r CWSBasicAuthCredentialNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSBasicAuthCredentialNewResponseEnvelope struct {
	Errors   []CWSBasicAuthCredentialNewResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSBasicAuthCredentialNewResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   BasicAuthCredential                                 `json:"result" api:"required"`
	Success  bool                                                `json:"success" api:"required"`
	JSON     cwsBasicAuthCredentialNewResponseEnvelopeJSON       `json:"-"`
}

// cwsBasicAuthCredentialNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSBasicAuthCredentialNewResponseEnvelope]
type cwsBasicAuthCredentialNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialNewResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code" api:"required"`
	Message string                                                `json:"message" api:"required"`
	Source  CWSBasicAuthCredentialNewResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsBasicAuthCredentialNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsBasicAuthCredentialNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CWSBasicAuthCredentialNewResponseEnvelopeErrors]
type cwsBasicAuthCredentialNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer" api:"required"`
	JSON    cwsBasicAuthCredentialNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsBasicAuthCredentialNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CWSBasicAuthCredentialNewResponseEnvelopeErrorsSource]
type cwsBasicAuthCredentialNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialNewResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code" api:"required"`
	Message string                                                  `json:"message" api:"required"`
	Source  CWSBasicAuthCredentialNewResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsBasicAuthCredentialNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsBasicAuthCredentialNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CWSBasicAuthCredentialNewResponseEnvelopeMessages]
type cwsBasicAuthCredentialNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer" api:"required"`
	JSON    cwsBasicAuthCredentialNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsBasicAuthCredentialNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CWSBasicAuthCredentialNewResponseEnvelopeMessagesSource]
type cwsBasicAuthCredentialNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialListParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Filter to credentials attached to this deployment
	DeploymentID param.Field[string] `query:"deployment_id" format:"uuid"`
	// Sort direction (default: asc)
	Direction param.Field[CWSBasicAuthCredentialListParamsDirection] `query:"direction"`
	// Field to order results by (default: created_at)
	Order param.Field[CWSBasicAuthCredentialListParamsOrder] `query:"order"`
	// Number of items per page (1-100, default 20)
	PerPage param.Field[int64] `query:"per_page"`
	// Filter by username (case-insensitive substring match)
	Username param.Field[string] `query:"username"`
}

// URLQuery serializes [CWSBasicAuthCredentialListParams]'s query parameters as
// `url.Values`.
func (r CWSBasicAuthCredentialListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction (default: asc)
type CWSBasicAuthCredentialListParamsDirection string

const (
	CWSBasicAuthCredentialListParamsDirectionAsc  CWSBasicAuthCredentialListParamsDirection = "asc"
	CWSBasicAuthCredentialListParamsDirectionDesc CWSBasicAuthCredentialListParamsDirection = "desc"
)

func (r CWSBasicAuthCredentialListParamsDirection) IsKnown() bool {
	switch r {
	case CWSBasicAuthCredentialListParamsDirectionAsc, CWSBasicAuthCredentialListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by (default: created_at)
type CWSBasicAuthCredentialListParamsOrder string

const (
	CWSBasicAuthCredentialListParamsOrderUsername  CWSBasicAuthCredentialListParamsOrder = "username"
	CWSBasicAuthCredentialListParamsOrderCreatedAt CWSBasicAuthCredentialListParamsOrder = "created_at"
	CWSBasicAuthCredentialListParamsOrderUpdatedAt CWSBasicAuthCredentialListParamsOrder = "updated_at"
)

func (r CWSBasicAuthCredentialListParamsOrder) IsKnown() bool {
	switch r {
	case CWSBasicAuthCredentialListParamsOrderUsername, CWSBasicAuthCredentialListParamsOrderCreatedAt, CWSBasicAuthCredentialListParamsOrderUpdatedAt:
		return true
	}
	return false
}

type CWSBasicAuthCredentialDeleteParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSBasicAuthCredentialDeleteResponseEnvelope struct {
	Errors   []CWSBasicAuthCredentialDeleteResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSBasicAuthCredentialDeleteResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   CWSBasicAuthCredentialDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                                   `json:"success" api:"required"`
	JSON     cwsBasicAuthCredentialDeleteResponseEnvelopeJSON       `json:"-"`
}

// cwsBasicAuthCredentialDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [CWSBasicAuthCredentialDeleteResponseEnvelope]
type cwsBasicAuthCredentialDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialDeleteResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code" api:"required"`
	Message string                                                   `json:"message" api:"required"`
	Source  CWSBasicAuthCredentialDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsBasicAuthCredentialDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsBasicAuthCredentialDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CWSBasicAuthCredentialDeleteResponseEnvelopeErrors]
type cwsBasicAuthCredentialDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer" api:"required"`
	JSON    cwsBasicAuthCredentialDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsBasicAuthCredentialDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [CWSBasicAuthCredentialDeleteResponseEnvelopeErrorsSource]
type cwsBasicAuthCredentialDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialDeleteResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code" api:"required"`
	Message string                                                     `json:"message" api:"required"`
	Source  CWSBasicAuthCredentialDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsBasicAuthCredentialDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsBasicAuthCredentialDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CWSBasicAuthCredentialDeleteResponseEnvelopeMessages]
type cwsBasicAuthCredentialDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer" api:"required"`
	JSON    cwsBasicAuthCredentialDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsBasicAuthCredentialDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CWSBasicAuthCredentialDeleteResponseEnvelopeMessagesSource]
type cwsBasicAuthCredentialDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialGetParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSBasicAuthCredentialGetResponseEnvelope struct {
	Errors   []CWSBasicAuthCredentialGetResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSBasicAuthCredentialGetResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   BasicAuthCredential                                 `json:"result" api:"required"`
	Success  bool                                                `json:"success" api:"required"`
	JSON     cwsBasicAuthCredentialGetResponseEnvelopeJSON       `json:"-"`
}

// cwsBasicAuthCredentialGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSBasicAuthCredentialGetResponseEnvelope]
type cwsBasicAuthCredentialGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code" api:"required"`
	Message string                                                `json:"message" api:"required"`
	Source  CWSBasicAuthCredentialGetResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsBasicAuthCredentialGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsBasicAuthCredentialGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CWSBasicAuthCredentialGetResponseEnvelopeErrors]
type cwsBasicAuthCredentialGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer" api:"required"`
	JSON    cwsBasicAuthCredentialGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsBasicAuthCredentialGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CWSBasicAuthCredentialGetResponseEnvelopeErrorsSource]
type cwsBasicAuthCredentialGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code" api:"required"`
	Message string                                                  `json:"message" api:"required"`
	Source  CWSBasicAuthCredentialGetResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsBasicAuthCredentialGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsBasicAuthCredentialGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CWSBasicAuthCredentialGetResponseEnvelopeMessages]
type cwsBasicAuthCredentialGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer" api:"required"`
	JSON    cwsBasicAuthCredentialGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsBasicAuthCredentialGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CWSBasicAuthCredentialGetResponseEnvelopeMessagesSource]
type cwsBasicAuthCredentialGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
