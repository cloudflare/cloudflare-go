// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// SSOService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSSOService] method instead.
type SSOService struct {
	Options []option.RequestOption
}

// NewSSOService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSSOService(opts ...option.RequestOption) (r *SSOService) {
	r = &SSOService{}
	r.Options = opts
	return
}

// Initialize new SSO connector
func (r *SSOService) New(ctx context.Context, params SSONewParams, opts ...option.RequestOption) (res *SSONewResponse, err error) {
	var env SSONewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/sso_connectors", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update SSO connector state
func (r *SSOService) Update(ctx context.Context, ssoConnectorID string, params SSOUpdateParams, opts ...option.RequestOption) (res *SSOUpdateResponse, err error) {
	var env SSOUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ssoConnectorID == "" {
		err = errors.New("missing required sso_connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/sso_connectors/%s", params.AccountID, ssoConnectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get all SSO connectors
func (r *SSOService) List(ctx context.Context, query SSOListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SSOListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/sso_connectors", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Get all SSO connectors
func (r *SSOService) ListAutoPaging(ctx context.Context, query SSOListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SSOListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete SSO connector
func (r *SSOService) Delete(ctx context.Context, ssoConnectorID string, body SSODeleteParams, opts ...option.RequestOption) (res *SSODeleteResponse, err error) {
	var env SSODeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ssoConnectorID == "" {
		err = errors.New("missing required sso_connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/sso_connectors/%s", body.AccountID, ssoConnectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Begin SSO connector verification
func (r *SSOService) BeginVerification(ctx context.Context, ssoConnectorID string, body SSOBeginVerificationParams, opts ...option.RequestOption) (res *SSOBeginVerificationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ssoConnectorID == "" {
		err = errors.New("missing required sso_connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/sso_connectors/%s/begin_verification", body.AccountID, ssoConnectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get single SSO connector
func (r *SSOService) Get(ctx context.Context, ssoConnectorID string, query SSOGetParams, opts ...option.RequestOption) (res *SSOGetResponse, err error) {
	var env SSOGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ssoConnectorID == "" {
		err = errors.New("missing required sso_connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/sso_connectors/%s", query.AccountID, ssoConnectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SSONewResponse struct {
	// SSO Connector identifier tag.
	ID string `json:"id"`
	// Timestamp for the creation of the SSO connector
	CreatedOn   time.Time `json:"created_on" format:"date-time"`
	EmailDomain string    `json:"email_domain"`
	Enabled     bool      `json:"enabled"`
	// Timestamp for the last update of the SSO connector
	UpdatedOn time.Time `json:"updated_on" format:"date-time"`
	// Controls the display of FedRAMP language to the user during SSO login
	UseFedrampLanguage bool                       `json:"use_fedramp_language"`
	Verification       SSONewResponseVerification `json:"verification"`
	JSON               ssoNewResponseJSON         `json:"-"`
}

// ssoNewResponseJSON contains the JSON metadata for the struct [SSONewResponse]
type ssoNewResponseJSON struct {
	ID                 apijson.Field
	CreatedOn          apijson.Field
	EmailDomain        apijson.Field
	Enabled            apijson.Field
	UpdatedOn          apijson.Field
	UseFedrampLanguage apijson.Field
	Verification       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SSONewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoNewResponseJSON) RawJSON() string {
	return r.raw
}

type SSONewResponseVerification struct {
	// DNS verification code. Add this entire string to the DNS TXT record of the email
	// domain to validate ownership.
	Code string `json:"code"`
	// The status of the verification code from the verification process.
	Status SSONewResponseVerificationStatus `json:"status"`
	JSON   ssoNewResponseVerificationJSON   `json:"-"`
}

// ssoNewResponseVerificationJSON contains the JSON metadata for the struct
// [SSONewResponseVerification]
type ssoNewResponseVerificationJSON struct {
	Code        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSONewResponseVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoNewResponseVerificationJSON) RawJSON() string {
	return r.raw
}

// The status of the verification code from the verification process.
type SSONewResponseVerificationStatus string

const (
	SSONewResponseVerificationStatusAwaiting SSONewResponseVerificationStatus = "awaiting"
	SSONewResponseVerificationStatusPending  SSONewResponseVerificationStatus = "pending"
	SSONewResponseVerificationStatusFailed   SSONewResponseVerificationStatus = "failed"
	SSONewResponseVerificationStatusVerified SSONewResponseVerificationStatus = "verified"
)

func (r SSONewResponseVerificationStatus) IsKnown() bool {
	switch r {
	case SSONewResponseVerificationStatusAwaiting, SSONewResponseVerificationStatusPending, SSONewResponseVerificationStatusFailed, SSONewResponseVerificationStatusVerified:
		return true
	}
	return false
}

type SSOUpdateResponse struct {
	// SSO Connector identifier tag.
	ID string `json:"id"`
	// Timestamp for the creation of the SSO connector
	CreatedOn   time.Time `json:"created_on" format:"date-time"`
	EmailDomain string    `json:"email_domain"`
	Enabled     bool      `json:"enabled"`
	// Timestamp for the last update of the SSO connector
	UpdatedOn time.Time `json:"updated_on" format:"date-time"`
	// Controls the display of FedRAMP language to the user during SSO login
	UseFedrampLanguage bool                          `json:"use_fedramp_language"`
	Verification       SSOUpdateResponseVerification `json:"verification"`
	JSON               ssoUpdateResponseJSON         `json:"-"`
}

// ssoUpdateResponseJSON contains the JSON metadata for the struct
// [SSOUpdateResponse]
type ssoUpdateResponseJSON struct {
	ID                 apijson.Field
	CreatedOn          apijson.Field
	EmailDomain        apijson.Field
	Enabled            apijson.Field
	UpdatedOn          apijson.Field
	UseFedrampLanguage apijson.Field
	Verification       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SSOUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SSOUpdateResponseVerification struct {
	// DNS verification code. Add this entire string to the DNS TXT record of the email
	// domain to validate ownership.
	Code string `json:"code"`
	// The status of the verification code from the verification process.
	Status SSOUpdateResponseVerificationStatus `json:"status"`
	JSON   ssoUpdateResponseVerificationJSON   `json:"-"`
}

// ssoUpdateResponseVerificationJSON contains the JSON metadata for the struct
// [SSOUpdateResponseVerification]
type ssoUpdateResponseVerificationJSON struct {
	Code        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSOUpdateResponseVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoUpdateResponseVerificationJSON) RawJSON() string {
	return r.raw
}

// The status of the verification code from the verification process.
type SSOUpdateResponseVerificationStatus string

const (
	SSOUpdateResponseVerificationStatusAwaiting SSOUpdateResponseVerificationStatus = "awaiting"
	SSOUpdateResponseVerificationStatusPending  SSOUpdateResponseVerificationStatus = "pending"
	SSOUpdateResponseVerificationStatusFailed   SSOUpdateResponseVerificationStatus = "failed"
	SSOUpdateResponseVerificationStatusVerified SSOUpdateResponseVerificationStatus = "verified"
)

func (r SSOUpdateResponseVerificationStatus) IsKnown() bool {
	switch r {
	case SSOUpdateResponseVerificationStatusAwaiting, SSOUpdateResponseVerificationStatusPending, SSOUpdateResponseVerificationStatusFailed, SSOUpdateResponseVerificationStatusVerified:
		return true
	}
	return false
}

type SSOListResponse struct {
	// SSO Connector identifier tag.
	ID string `json:"id"`
	// Timestamp for the creation of the SSO connector
	CreatedOn   time.Time `json:"created_on" format:"date-time"`
	EmailDomain string    `json:"email_domain"`
	Enabled     bool      `json:"enabled"`
	// Timestamp for the last update of the SSO connector
	UpdatedOn time.Time `json:"updated_on" format:"date-time"`
	// Controls the display of FedRAMP language to the user during SSO login
	UseFedrampLanguage bool                        `json:"use_fedramp_language"`
	Verification       SSOListResponseVerification `json:"verification"`
	JSON               ssoListResponseJSON         `json:"-"`
}

// ssoListResponseJSON contains the JSON metadata for the struct [SSOListResponse]
type ssoListResponseJSON struct {
	ID                 apijson.Field
	CreatedOn          apijson.Field
	EmailDomain        apijson.Field
	Enabled            apijson.Field
	UpdatedOn          apijson.Field
	UseFedrampLanguage apijson.Field
	Verification       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SSOListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoListResponseJSON) RawJSON() string {
	return r.raw
}

type SSOListResponseVerification struct {
	// DNS verification code. Add this entire string to the DNS TXT record of the email
	// domain to validate ownership.
	Code string `json:"code"`
	// The status of the verification code from the verification process.
	Status SSOListResponseVerificationStatus `json:"status"`
	JSON   ssoListResponseVerificationJSON   `json:"-"`
}

// ssoListResponseVerificationJSON contains the JSON metadata for the struct
// [SSOListResponseVerification]
type ssoListResponseVerificationJSON struct {
	Code        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSOListResponseVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoListResponseVerificationJSON) RawJSON() string {
	return r.raw
}

// The status of the verification code from the verification process.
type SSOListResponseVerificationStatus string

const (
	SSOListResponseVerificationStatusAwaiting SSOListResponseVerificationStatus = "awaiting"
	SSOListResponseVerificationStatusPending  SSOListResponseVerificationStatus = "pending"
	SSOListResponseVerificationStatusFailed   SSOListResponseVerificationStatus = "failed"
	SSOListResponseVerificationStatusVerified SSOListResponseVerificationStatus = "verified"
)

func (r SSOListResponseVerificationStatus) IsKnown() bool {
	switch r {
	case SSOListResponseVerificationStatusAwaiting, SSOListResponseVerificationStatusPending, SSOListResponseVerificationStatusFailed, SSOListResponseVerificationStatusVerified:
		return true
	}
	return false
}

type SSODeleteResponse struct {
	// Identifier
	ID   string                `json:"id,required"`
	JSON ssoDeleteResponseJSON `json:"-"`
}

// ssoDeleteResponseJSON contains the JSON metadata for the struct
// [SSODeleteResponse]
type ssoDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSODeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SSOBeginVerificationResponse struct {
	Errors   []SSOBeginVerificationResponseError   `json:"errors,required"`
	Messages []SSOBeginVerificationResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success SSOBeginVerificationResponseSuccess `json:"success,required"`
	JSON    ssoBeginVerificationResponseJSON    `json:"-"`
}

// ssoBeginVerificationResponseJSON contains the JSON metadata for the struct
// [SSOBeginVerificationResponse]
type ssoBeginVerificationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSOBeginVerificationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoBeginVerificationResponseJSON) RawJSON() string {
	return r.raw
}

type SSOBeginVerificationResponseError struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           SSOBeginVerificationResponseErrorsSource `json:"source"`
	JSON             ssoBeginVerificationResponseErrorJSON    `json:"-"`
}

// ssoBeginVerificationResponseErrorJSON contains the JSON metadata for the struct
// [SSOBeginVerificationResponseError]
type ssoBeginVerificationResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SSOBeginVerificationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoBeginVerificationResponseErrorJSON) RawJSON() string {
	return r.raw
}

type SSOBeginVerificationResponseErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    ssoBeginVerificationResponseErrorsSourceJSON `json:"-"`
}

// ssoBeginVerificationResponseErrorsSourceJSON contains the JSON metadata for the
// struct [SSOBeginVerificationResponseErrorsSource]
type ssoBeginVerificationResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSOBeginVerificationResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoBeginVerificationResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SSOBeginVerificationResponseMessage struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           SSOBeginVerificationResponseMessagesSource `json:"source"`
	JSON             ssoBeginVerificationResponseMessageJSON    `json:"-"`
}

// ssoBeginVerificationResponseMessageJSON contains the JSON metadata for the
// struct [SSOBeginVerificationResponseMessage]
type ssoBeginVerificationResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SSOBeginVerificationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoBeginVerificationResponseMessageJSON) RawJSON() string {
	return r.raw
}

type SSOBeginVerificationResponseMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    ssoBeginVerificationResponseMessagesSourceJSON `json:"-"`
}

// ssoBeginVerificationResponseMessagesSourceJSON contains the JSON metadata for
// the struct [SSOBeginVerificationResponseMessagesSource]
type ssoBeginVerificationResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSOBeginVerificationResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoBeginVerificationResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SSOBeginVerificationResponseSuccess bool

const (
	SSOBeginVerificationResponseSuccessTrue SSOBeginVerificationResponseSuccess = true
)

func (r SSOBeginVerificationResponseSuccess) IsKnown() bool {
	switch r {
	case SSOBeginVerificationResponseSuccessTrue:
		return true
	}
	return false
}

type SSOGetResponse struct {
	// SSO Connector identifier tag.
	ID string `json:"id"`
	// Timestamp for the creation of the SSO connector
	CreatedOn   time.Time `json:"created_on" format:"date-time"`
	EmailDomain string    `json:"email_domain"`
	Enabled     bool      `json:"enabled"`
	// Timestamp for the last update of the SSO connector
	UpdatedOn time.Time `json:"updated_on" format:"date-time"`
	// Controls the display of FedRAMP language to the user during SSO login
	UseFedrampLanguage bool                       `json:"use_fedramp_language"`
	Verification       SSOGetResponseVerification `json:"verification"`
	JSON               ssoGetResponseJSON         `json:"-"`
}

// ssoGetResponseJSON contains the JSON metadata for the struct [SSOGetResponse]
type ssoGetResponseJSON struct {
	ID                 apijson.Field
	CreatedOn          apijson.Field
	EmailDomain        apijson.Field
	Enabled            apijson.Field
	UpdatedOn          apijson.Field
	UseFedrampLanguage apijson.Field
	Verification       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SSOGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoGetResponseJSON) RawJSON() string {
	return r.raw
}

type SSOGetResponseVerification struct {
	// DNS verification code. Add this entire string to the DNS TXT record of the email
	// domain to validate ownership.
	Code string `json:"code"`
	// The status of the verification code from the verification process.
	Status SSOGetResponseVerificationStatus `json:"status"`
	JSON   ssoGetResponseVerificationJSON   `json:"-"`
}

// ssoGetResponseVerificationJSON contains the JSON metadata for the struct
// [SSOGetResponseVerification]
type ssoGetResponseVerificationJSON struct {
	Code        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSOGetResponseVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoGetResponseVerificationJSON) RawJSON() string {
	return r.raw
}

// The status of the verification code from the verification process.
type SSOGetResponseVerificationStatus string

const (
	SSOGetResponseVerificationStatusAwaiting SSOGetResponseVerificationStatus = "awaiting"
	SSOGetResponseVerificationStatusPending  SSOGetResponseVerificationStatus = "pending"
	SSOGetResponseVerificationStatusFailed   SSOGetResponseVerificationStatus = "failed"
	SSOGetResponseVerificationStatusVerified SSOGetResponseVerificationStatus = "verified"
)

func (r SSOGetResponseVerificationStatus) IsKnown() bool {
	switch r {
	case SSOGetResponseVerificationStatusAwaiting, SSOGetResponseVerificationStatusPending, SSOGetResponseVerificationStatusFailed, SSOGetResponseVerificationStatusVerified:
		return true
	}
	return false
}

type SSONewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Email domain of the new SSO connector
	EmailDomain param.Field[string] `json:"email_domain,required"`
	// Begin the verification process after creation
	BeginVerification param.Field[bool] `json:"begin_verification"`
	// Controls the display of FedRAMP language to the user during SSO login
	UseFedrampLanguage param.Field[bool] `json:"use_fedramp_language"`
}

func (r SSONewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SSONewResponseEnvelope struct {
	Errors   []SSONewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSONewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SSONewResponseEnvelopeSuccess `json:"success,required"`
	Result  SSONewResponse                `json:"result"`
	JSON    ssoNewResponseEnvelopeJSON    `json:"-"`
}

// ssoNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SSONewResponseEnvelope]
type ssoNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSONewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SSONewResponseEnvelopeErrors struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           SSONewResponseEnvelopeErrorsSource `json:"source"`
	JSON             ssoNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// ssoNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SSONewResponseEnvelopeErrors]
type ssoNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SSONewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SSONewResponseEnvelopeErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    ssoNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ssoNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the struct
// [SSONewResponseEnvelopeErrorsSource]
type ssoNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSONewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SSONewResponseEnvelopeMessages struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           SSONewResponseEnvelopeMessagesSource `json:"source"`
	JSON             ssoNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// ssoNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SSONewResponseEnvelopeMessages]
type ssoNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SSONewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SSONewResponseEnvelopeMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    ssoNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ssoNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [SSONewResponseEnvelopeMessagesSource]
type ssoNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSONewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SSONewResponseEnvelopeSuccess bool

const (
	SSONewResponseEnvelopeSuccessTrue SSONewResponseEnvelopeSuccess = true
)

func (r SSONewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SSONewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SSOUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// SSO Connector enabled state
	Enabled param.Field[bool] `json:"enabled"`
	// Controls the display of FedRAMP language to the user during SSO login
	UseFedrampLanguage param.Field[bool] `json:"use_fedramp_language"`
}

func (r SSOUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SSOUpdateResponseEnvelope struct {
	Errors   []SSOUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSOUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SSOUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  SSOUpdateResponse                `json:"result"`
	JSON    ssoUpdateResponseEnvelopeJSON    `json:"-"`
}

// ssoUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SSOUpdateResponseEnvelope]
type ssoUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSOUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SSOUpdateResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           SSOUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             ssoUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// ssoUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SSOUpdateResponseEnvelopeErrors]
type ssoUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SSOUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SSOUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    ssoUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ssoUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [SSOUpdateResponseEnvelopeErrorsSource]
type ssoUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSOUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SSOUpdateResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           SSOUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             ssoUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// ssoUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SSOUpdateResponseEnvelopeMessages]
type ssoUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SSOUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SSOUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    ssoUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ssoUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [SSOUpdateResponseEnvelopeMessagesSource]
type ssoUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSOUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SSOUpdateResponseEnvelopeSuccess bool

const (
	SSOUpdateResponseEnvelopeSuccessTrue SSOUpdateResponseEnvelopeSuccess = true
)

func (r SSOUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SSOUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SSOListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SSODeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SSODeleteResponseEnvelope struct {
	Errors   []SSODeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSODeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SSODeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  SSODeleteResponse                `json:"result,nullable"`
	JSON    ssoDeleteResponseEnvelopeJSON    `json:"-"`
}

// ssoDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SSODeleteResponseEnvelope]
type ssoDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSODeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SSODeleteResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           SSODeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             ssoDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// ssoDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SSODeleteResponseEnvelopeErrors]
type ssoDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SSODeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SSODeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    ssoDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ssoDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [SSODeleteResponseEnvelopeErrorsSource]
type ssoDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSODeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SSODeleteResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           SSODeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             ssoDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// ssoDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SSODeleteResponseEnvelopeMessages]
type ssoDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SSODeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SSODeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    ssoDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ssoDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [SSODeleteResponseEnvelopeMessagesSource]
type ssoDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSODeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SSODeleteResponseEnvelopeSuccess bool

const (
	SSODeleteResponseEnvelopeSuccessTrue SSODeleteResponseEnvelopeSuccess = true
)

func (r SSODeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SSODeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SSOBeginVerificationParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SSOGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SSOGetResponseEnvelope struct {
	Errors   []SSOGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSOGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SSOGetResponseEnvelopeSuccess `json:"success,required"`
	Result  SSOGetResponse                `json:"result"`
	JSON    ssoGetResponseEnvelopeJSON    `json:"-"`
}

// ssoGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SSOGetResponseEnvelope]
type ssoGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSOGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SSOGetResponseEnvelopeErrors struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           SSOGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             ssoGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// ssoGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SSOGetResponseEnvelopeErrors]
type ssoGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SSOGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SSOGetResponseEnvelopeErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    ssoGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ssoGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the struct
// [SSOGetResponseEnvelopeErrorsSource]
type ssoGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSOGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SSOGetResponseEnvelopeMessages struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           SSOGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             ssoGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// ssoGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SSOGetResponseEnvelopeMessages]
type ssoGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SSOGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SSOGetResponseEnvelopeMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    ssoGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ssoGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [SSOGetResponseEnvelopeMessagesSource]
type ssoGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSOGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ssoGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SSOGetResponseEnvelopeSuccess bool

const (
	SSOGetResponseEnvelopeSuccessTrue SSOGetResponseEnvelopeSuccess = true
)

func (r SSOGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SSOGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
