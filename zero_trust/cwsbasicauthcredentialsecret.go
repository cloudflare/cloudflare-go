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

// CWSBasicAuthCredentialSecretService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCWSBasicAuthCredentialSecretService] method instead.
type CWSBasicAuthCredentialSecretService struct {
	Options []option.RequestOption
}

// NewCWSBasicAuthCredentialSecretService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCWSBasicAuthCredentialSecretService(opts ...option.RequestOption) (r *CWSBasicAuthCredentialSecretService) {
	r = &CWSBasicAuthCredentialSecretService{}
	r.Options = opts
	return
}

// Adds a password secret to the credential.
func (r *CWSBasicAuthCredentialSecretService) New(ctx context.Context, credentialID string, params CWSBasicAuthCredentialSecretNewParams, opts ...option.RequestOption) (res *BasicAuthSecret, err error) {
	var env CWSBasicAuthCredentialSecretNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if credentialID == "" {
		err = errors.New("missing required credential_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/basic-auth-credentials/%s/secrets", params.AccountID, credentialID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of secrets belonging to the credential.
func (r *CWSBasicAuthCredentialSecretService) List(ctx context.Context, credentialID string, params CWSBasicAuthCredentialSecretListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[BasicAuthSecret], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if credentialID == "" {
		err = errors.New("missing required credential_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/basic-auth-credentials/%s/secrets", params.AccountID, credentialID)
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

// Returns a paginated list of secrets belonging to the credential.
func (r *CWSBasicAuthCredentialSecretService) ListAutoPaging(ctx context.Context, credentialID string, params CWSBasicAuthCredentialSecretListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[BasicAuthSecret] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, credentialID, params, opts...))
}

// Deletes the specified secret from the credential.
func (r *CWSBasicAuthCredentialSecretService) Delete(ctx context.Context, credentialID string, secretID string, body CWSBasicAuthCredentialSecretDeleteParams, opts ...option.RequestOption) (res *CWSBasicAuthCredentialSecretDeleteResponse, err error) {
	var env CWSBasicAuthCredentialSecretDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if credentialID == "" {
		err = errors.New("missing required credential_id parameter")
		return nil, err
	}
	if secretID == "" {
		err = errors.New("missing required secret_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/basic-auth-credentials/%s/secrets/%s", body.AccountID, credentialID, secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the specified secret's metadata without exposing its password or hash.
func (r *CWSBasicAuthCredentialSecretService) Get(ctx context.Context, credentialID string, secretID string, query CWSBasicAuthCredentialSecretGetParams, opts ...option.RequestOption) (res *BasicAuthSecret, err error) {
	var env CWSBasicAuthCredentialSecretGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if credentialID == "" {
		err = errors.New("missing required credential_id parameter")
		return nil, err
	}
	if secretID == "" {
		err = errors.New("missing required secret_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/basic-auth-credentials/%s/secrets/%s", query.AccountID, credentialID, secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type BasicAuthSecret struct {
	// Server-assigned UUID
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Optional RFC 3339 expiration time. Omit for no expiration.
	ExpiresAt time.Time           `json:"expires_at" format:"date-time"`
	JSON      basicAuthSecretJSON `json:"-"`
}

// basicAuthSecretJSON contains the JSON metadata for the struct [BasicAuthSecret]
type basicAuthSecretJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	ExpiresAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasicAuthSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basicAuthSecretJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretDeleteResponse struct {
	ID   string                                         `json:"id" api:"required" format:"uuid"`
	JSON cwsBasicAuthCredentialSecretDeleteResponseJSON `json:"-"`
}

// cwsBasicAuthCredentialSecretDeleteResponseJSON contains the JSON metadata for
// the struct [CWSBasicAuthCredentialSecretDeleteResponse]
type cwsBasicAuthCredentialSecretDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretNewParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Password stored only as a one-way hash and never returned. Maximum 72 UTF-8
	// bytes.
	Password param.Field[string] `json:"password" api:"required"`
	// Optional deactivation deadline (RFC 3339). Omit or null for no expiry.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
}

func (r CWSBasicAuthCredentialSecretNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSBasicAuthCredentialSecretNewResponseEnvelope struct {
	Errors   []CWSBasicAuthCredentialSecretNewResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSBasicAuthCredentialSecretNewResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   BasicAuthSecret                                           `json:"result" api:"required"`
	Success  bool                                                      `json:"success" api:"required"`
	JSON     cwsBasicAuthCredentialSecretNewResponseEnvelopeJSON       `json:"-"`
}

// cwsBasicAuthCredentialSecretNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [CWSBasicAuthCredentialSecretNewResponseEnvelope]
type cwsBasicAuthCredentialSecretNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretNewResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code" api:"required"`
	Message string                                                      `json:"message" api:"required"`
	Source  CWSBasicAuthCredentialSecretNewResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsBasicAuthCredentialSecretNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsBasicAuthCredentialSecretNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CWSBasicAuthCredentialSecretNewResponseEnvelopeErrors]
type cwsBasicAuthCredentialSecretNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer" api:"required"`
	JSON    cwsBasicAuthCredentialSecretNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsBasicAuthCredentialSecretNewResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [CWSBasicAuthCredentialSecretNewResponseEnvelopeErrorsSource]
type cwsBasicAuthCredentialSecretNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretNewResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code" api:"required"`
	Message string                                                        `json:"message" api:"required"`
	Source  CWSBasicAuthCredentialSecretNewResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsBasicAuthCredentialSecretNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsBasicAuthCredentialSecretNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CWSBasicAuthCredentialSecretNewResponseEnvelopeMessages]
type cwsBasicAuthCredentialSecretNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer" api:"required"`
	JSON    cwsBasicAuthCredentialSecretNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsBasicAuthCredentialSecretNewResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [CWSBasicAuthCredentialSecretNewResponseEnvelopeMessagesSource]
type cwsBasicAuthCredentialSecretNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretListParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Sort direction (default: asc)
	Direction param.Field[CWSBasicAuthCredentialSecretListParamsDirection] `query:"direction"`
	// Field to order results by (default: created_at)
	Order param.Field[CWSBasicAuthCredentialSecretListParamsOrder] `query:"order"`
	// Number of items per page (1-100, default 20)
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [CWSBasicAuthCredentialSecretListParams]'s query parameters
// as `url.Values`.
func (r CWSBasicAuthCredentialSecretListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction (default: asc)
type CWSBasicAuthCredentialSecretListParamsDirection string

const (
	CWSBasicAuthCredentialSecretListParamsDirectionAsc  CWSBasicAuthCredentialSecretListParamsDirection = "asc"
	CWSBasicAuthCredentialSecretListParamsDirectionDesc CWSBasicAuthCredentialSecretListParamsDirection = "desc"
)

func (r CWSBasicAuthCredentialSecretListParamsDirection) IsKnown() bool {
	switch r {
	case CWSBasicAuthCredentialSecretListParamsDirectionAsc, CWSBasicAuthCredentialSecretListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by (default: created_at)
type CWSBasicAuthCredentialSecretListParamsOrder string

const (
	CWSBasicAuthCredentialSecretListParamsOrderCreatedAt CWSBasicAuthCredentialSecretListParamsOrder = "created_at"
)

func (r CWSBasicAuthCredentialSecretListParamsOrder) IsKnown() bool {
	switch r {
	case CWSBasicAuthCredentialSecretListParamsOrderCreatedAt:
		return true
	}
	return false
}

type CWSBasicAuthCredentialSecretDeleteParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSBasicAuthCredentialSecretDeleteResponseEnvelope struct {
	Errors   []CWSBasicAuthCredentialSecretDeleteResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSBasicAuthCredentialSecretDeleteResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   CWSBasicAuthCredentialSecretDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                                         `json:"success" api:"required"`
	JSON     cwsBasicAuthCredentialSecretDeleteResponseEnvelopeJSON       `json:"-"`
}

// cwsBasicAuthCredentialSecretDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [CWSBasicAuthCredentialSecretDeleteResponseEnvelope]
type cwsBasicAuthCredentialSecretDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretDeleteResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code" api:"required"`
	Message string                                                         `json:"message" api:"required"`
	Source  CWSBasicAuthCredentialSecretDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsBasicAuthCredentialSecretDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsBasicAuthCredentialSecretDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CWSBasicAuthCredentialSecretDeleteResponseEnvelopeErrors]
type cwsBasicAuthCredentialSecretDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                             `json:"pointer" api:"required"`
	JSON    cwsBasicAuthCredentialSecretDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsBasicAuthCredentialSecretDeleteResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [CWSBasicAuthCredentialSecretDeleteResponseEnvelopeErrorsSource]
type cwsBasicAuthCredentialSecretDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretDeleteResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code" api:"required"`
	Message string                                                           `json:"message" api:"required"`
	Source  CWSBasicAuthCredentialSecretDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsBasicAuthCredentialSecretDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsBasicAuthCredentialSecretDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CWSBasicAuthCredentialSecretDeleteResponseEnvelopeMessages]
type cwsBasicAuthCredentialSecretDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                               `json:"pointer" api:"required"`
	JSON    cwsBasicAuthCredentialSecretDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsBasicAuthCredentialSecretDeleteResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [CWSBasicAuthCredentialSecretDeleteResponseEnvelopeMessagesSource]
type cwsBasicAuthCredentialSecretDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretGetParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSBasicAuthCredentialSecretGetResponseEnvelope struct {
	Errors   []CWSBasicAuthCredentialSecretGetResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSBasicAuthCredentialSecretGetResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   BasicAuthSecret                                           `json:"result" api:"required"`
	Success  bool                                                      `json:"success" api:"required"`
	JSON     cwsBasicAuthCredentialSecretGetResponseEnvelopeJSON       `json:"-"`
}

// cwsBasicAuthCredentialSecretGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [CWSBasicAuthCredentialSecretGetResponseEnvelope]
type cwsBasicAuthCredentialSecretGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretGetResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code" api:"required"`
	Message string                                                      `json:"message" api:"required"`
	Source  CWSBasicAuthCredentialSecretGetResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsBasicAuthCredentialSecretGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsBasicAuthCredentialSecretGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CWSBasicAuthCredentialSecretGetResponseEnvelopeErrors]
type cwsBasicAuthCredentialSecretGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer" api:"required"`
	JSON    cwsBasicAuthCredentialSecretGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsBasicAuthCredentialSecretGetResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [CWSBasicAuthCredentialSecretGetResponseEnvelopeErrorsSource]
type cwsBasicAuthCredentialSecretGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretGetResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code" api:"required"`
	Message string                                                        `json:"message" api:"required"`
	Source  CWSBasicAuthCredentialSecretGetResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsBasicAuthCredentialSecretGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsBasicAuthCredentialSecretGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CWSBasicAuthCredentialSecretGetResponseEnvelopeMessages]
type cwsBasicAuthCredentialSecretGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSBasicAuthCredentialSecretGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer" api:"required"`
	JSON    cwsBasicAuthCredentialSecretGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsBasicAuthCredentialSecretGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [CWSBasicAuthCredentialSecretGetResponseEnvelopeMessagesSource]
type cwsBasicAuthCredentialSecretGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSBasicAuthCredentialSecretGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsBasicAuthCredentialSecretGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
