// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// CasbPostureWebhookService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbPostureWebhookService] method instead.
type CasbPostureWebhookService struct {
	Options []option.RequestOption
	Jobs    *CasbPostureWebhookJobService
}

// NewCasbPostureWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCasbPostureWebhookService(opts ...option.RequestOption) (r *CasbPostureWebhookService) {
	r = &CasbPostureWebhookService{}
	r.Options = opts
	r.Jobs = NewCasbPostureWebhookJobService(opts...)
	return
}

// Creates a new webhook configuration for sending finding notifications to
// external endpoints.
func (r *CasbPostureWebhookService) New(ctx context.Context, params CasbPostureWebhookNewParams, opts ...option.RequestOption) (res *CasbPostureWebhookNewResponse, err error) {
	var env CasbPostureWebhookNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/webhooks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an existing webhook configuration with new settings.
func (r *CasbPostureWebhookService) Update(ctx context.Context, webhookID string, params CasbPostureWebhookUpdateParams, opts ...option.RequestOption) (res *CasbPostureWebhookUpdateResponse, err error) {
	var env CasbPostureWebhookUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/webhooks/%s", params.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves all webhook configurations for the authenticated account. Returns an
// array of webhook configurations that can be used to send finding notifications.
func (r *CasbPostureWebhookService) List(ctx context.Context, query CasbPostureWebhookListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CasbPostureWebhookListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/webhooks", query.AccountID)
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

// Retrieves all webhook configurations for the authenticated account. Returns an
// array of webhook configurations that can be used to send finding notifications.
func (r *CasbPostureWebhookService) ListAutoPaging(ctx context.Context, query CasbPostureWebhookListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CasbPostureWebhookListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Soft deletes a webhook configuration by its unique identifier. The webhook will
// be marked as deleted and will no longer be available for use.
func (r *CasbPostureWebhookService) Delete(ctx context.Context, webhookID string, body CasbPostureWebhookDeleteParams, opts ...option.RequestOption) (res *CasbPostureWebhookDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/webhooks/%s", body.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Sends a test webhook event to the specified destination URL to verify the
// webhook endpoint is reachable and properly configured. This allows customers to
// validate their webhook configuration before creating the actual webhook
// resource.
//
// The test payload includes:
//
// - event_type: "webhook.test"
// - timestamp: Current UTC timestamp
// - message: Test message indicating this is from Cloudflare CASB
// - data: Object with test: true
func (r *CasbPostureWebhookService) Evaluate(ctx context.Context, params CasbPostureWebhookEvaluateParams, opts ...option.RequestOption) (res *CasbPostureWebhookEvaluateResponse, err error) {
	var env CasbPostureWebhookEvaluateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/webhooks/evaluate", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Sends a test webhook event using an existing webhook configuration. This allows
// customers to verify their webhook endpoint is still reachable and properly
// configured after creating the webhook resource.
//
// The test payload includes:
//
// - event_type: "webhook.test"
// - timestamp: Current UTC timestamp
// - message: Test message indicating this is from Cloudflare CASB
// - data: Object with test: true
func (r *CasbPostureWebhookService) EvaluateExisting(ctx context.Context, webhookID string, body CasbPostureWebhookEvaluateExistingParams, opts ...option.RequestOption) (res *CasbPostureWebhookEvaluateExistingResponse, err error) {
	var env CasbPostureWebhookEvaluateExistingResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/webhooks/%s/evaluate", body.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves a specific webhook configuration by its unique identifier.
func (r *CasbPostureWebhookService) Get(ctx context.Context, webhookID string, query CasbPostureWebhookGetParams, opts ...option.RequestOption) (res *CasbPostureWebhookGetResponse, err error) {
	var env CasbPostureWebhookGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/webhooks/%s", query.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Webhook configuration for sending finding notifications.
type CasbPostureWebhookNewResponse struct {
	// Unique identifier for the specific webhook configuration.
	ID string `json:"id" api:"required" format:"uuid"`
	// Type of authentication used for the webhook.
	AuthenticationType CasbPostureWebhookNewResponseAuthenticationType `json:"authentication_type" api:"required"`
	// Timestamp when the webhook configuration was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Target URL for the webhook configuration. Where resulting data will be sent.
	DestinationURL string `json:"destination_url" api:"required" format:"uri"`
	// Account-specified display label for the webhook configuration.
	Label string `json:"label" api:"required"`
	// Current status of the webhook configuration. If disabled, data cannot be sent
	// through this configuration.
	Status CasbPostureWebhookNewResponseStatus `json:"status" api:"required"`
	// Timestamp when the webhook configuration was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Version number of the configuration.
	Version int64 `json:"version" api:"required"`
	// List of header keys configured for this webhook. Values are not included for
	// security reasons.
	Headers []CasbPostureWebhookNewResponseHeader `json:"headers"`
	JSON    casbPostureWebhookNewResponseJSON     `json:"-"`
}

// casbPostureWebhookNewResponseJSON contains the JSON metadata for the struct
// [CasbPostureWebhookNewResponse]
type casbPostureWebhookNewResponseJSON struct {
	ID                 apijson.Field
	AuthenticationType apijson.Field
	CreatedAt          apijson.Field
	DestinationURL     apijson.Field
	Label              apijson.Field
	Status             apijson.Field
	UpdatedAt          apijson.Field
	Version            apijson.Field
	Headers            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CasbPostureWebhookNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookNewResponseJSON) RawJSON() string {
	return r.raw
}

// Type of authentication used for the webhook.
type CasbPostureWebhookNewResponseAuthenticationType string

const (
	CasbPostureWebhookNewResponseAuthenticationTypeBasicAuth     CasbPostureWebhookNewResponseAuthenticationType = "Basic Auth"
	CasbPostureWebhookNewResponseAuthenticationTypeNone          CasbPostureWebhookNewResponseAuthenticationType = "None"
	CasbPostureWebhookNewResponseAuthenticationTypeBearerAuth    CasbPostureWebhookNewResponseAuthenticationType = "Bearer Auth"
	CasbPostureWebhookNewResponseAuthenticationTypeStaticHeaders CasbPostureWebhookNewResponseAuthenticationType = "Static Headers"
	CasbPostureWebhookNewResponseAuthenticationTypeHmacSigning   CasbPostureWebhookNewResponseAuthenticationType = "HMAC-Signing"
)

func (r CasbPostureWebhookNewResponseAuthenticationType) IsKnown() bool {
	switch r {
	case CasbPostureWebhookNewResponseAuthenticationTypeBasicAuth, CasbPostureWebhookNewResponseAuthenticationTypeNone, CasbPostureWebhookNewResponseAuthenticationTypeBearerAuth, CasbPostureWebhookNewResponseAuthenticationTypeStaticHeaders, CasbPostureWebhookNewResponseAuthenticationTypeHmacSigning:
		return true
	}
	return false
}

// Current status of the webhook configuration. If disabled, data cannot be sent
// through this configuration.
type CasbPostureWebhookNewResponseStatus string

const (
	CasbPostureWebhookNewResponseStatusEnabled  CasbPostureWebhookNewResponseStatus = "enabled"
	CasbPostureWebhookNewResponseStatusDisabled CasbPostureWebhookNewResponseStatus = "disabled"
)

func (r CasbPostureWebhookNewResponseStatus) IsKnown() bool {
	switch r {
	case CasbPostureWebhookNewResponseStatusEnabled, CasbPostureWebhookNewResponseStatusDisabled:
		return true
	}
	return false
}

type CasbPostureWebhookNewResponseHeader struct {
	// Header key name (lowercase).
	Key  string                                  `json:"key"`
	JSON casbPostureWebhookNewResponseHeaderJSON `json:"-"`
}

// casbPostureWebhookNewResponseHeaderJSON contains the JSON metadata for the
// struct [CasbPostureWebhookNewResponseHeader]
type casbPostureWebhookNewResponseHeaderJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookNewResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookNewResponseHeaderJSON) RawJSON() string {
	return r.raw
}

// Webhook configuration for sending finding notifications.
type CasbPostureWebhookUpdateResponse struct {
	// Unique identifier for the specific webhook configuration.
	ID string `json:"id" api:"required" format:"uuid"`
	// Type of authentication used for the webhook.
	AuthenticationType CasbPostureWebhookUpdateResponseAuthenticationType `json:"authentication_type" api:"required"`
	// Timestamp when the webhook configuration was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Target URL for the webhook configuration. Where resulting data will be sent.
	DestinationURL string `json:"destination_url" api:"required" format:"uri"`
	// Account-specified display label for the webhook configuration.
	Label string `json:"label" api:"required"`
	// Current status of the webhook configuration. If disabled, data cannot be sent
	// through this configuration.
	Status CasbPostureWebhookUpdateResponseStatus `json:"status" api:"required"`
	// Timestamp when the webhook configuration was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Version number of the configuration.
	Version int64 `json:"version" api:"required"`
	// List of header keys configured for this webhook. Values are not included for
	// security reasons.
	Headers []CasbPostureWebhookUpdateResponseHeader `json:"headers"`
	JSON    casbPostureWebhookUpdateResponseJSON     `json:"-"`
}

// casbPostureWebhookUpdateResponseJSON contains the JSON metadata for the struct
// [CasbPostureWebhookUpdateResponse]
type casbPostureWebhookUpdateResponseJSON struct {
	ID                 apijson.Field
	AuthenticationType apijson.Field
	CreatedAt          apijson.Field
	DestinationURL     apijson.Field
	Label              apijson.Field
	Status             apijson.Field
	UpdatedAt          apijson.Field
	Version            apijson.Field
	Headers            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CasbPostureWebhookUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Type of authentication used for the webhook.
type CasbPostureWebhookUpdateResponseAuthenticationType string

const (
	CasbPostureWebhookUpdateResponseAuthenticationTypeBasicAuth     CasbPostureWebhookUpdateResponseAuthenticationType = "Basic Auth"
	CasbPostureWebhookUpdateResponseAuthenticationTypeNone          CasbPostureWebhookUpdateResponseAuthenticationType = "None"
	CasbPostureWebhookUpdateResponseAuthenticationTypeBearerAuth    CasbPostureWebhookUpdateResponseAuthenticationType = "Bearer Auth"
	CasbPostureWebhookUpdateResponseAuthenticationTypeStaticHeaders CasbPostureWebhookUpdateResponseAuthenticationType = "Static Headers"
	CasbPostureWebhookUpdateResponseAuthenticationTypeHmacSigning   CasbPostureWebhookUpdateResponseAuthenticationType = "HMAC-Signing"
)

func (r CasbPostureWebhookUpdateResponseAuthenticationType) IsKnown() bool {
	switch r {
	case CasbPostureWebhookUpdateResponseAuthenticationTypeBasicAuth, CasbPostureWebhookUpdateResponseAuthenticationTypeNone, CasbPostureWebhookUpdateResponseAuthenticationTypeBearerAuth, CasbPostureWebhookUpdateResponseAuthenticationTypeStaticHeaders, CasbPostureWebhookUpdateResponseAuthenticationTypeHmacSigning:
		return true
	}
	return false
}

// Current status of the webhook configuration. If disabled, data cannot be sent
// through this configuration.
type CasbPostureWebhookUpdateResponseStatus string

const (
	CasbPostureWebhookUpdateResponseStatusEnabled  CasbPostureWebhookUpdateResponseStatus = "enabled"
	CasbPostureWebhookUpdateResponseStatusDisabled CasbPostureWebhookUpdateResponseStatus = "disabled"
)

func (r CasbPostureWebhookUpdateResponseStatus) IsKnown() bool {
	switch r {
	case CasbPostureWebhookUpdateResponseStatusEnabled, CasbPostureWebhookUpdateResponseStatusDisabled:
		return true
	}
	return false
}

type CasbPostureWebhookUpdateResponseHeader struct {
	// Header key name (lowercase).
	Key  string                                     `json:"key"`
	JSON casbPostureWebhookUpdateResponseHeaderJSON `json:"-"`
}

// casbPostureWebhookUpdateResponseHeaderJSON contains the JSON metadata for the
// struct [CasbPostureWebhookUpdateResponseHeader]
type casbPostureWebhookUpdateResponseHeaderJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookUpdateResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookUpdateResponseHeaderJSON) RawJSON() string {
	return r.raw
}

// Webhook configuration for sending finding notifications.
type CasbPostureWebhookListResponse struct {
	// Unique identifier for the specific webhook configuration.
	ID string `json:"id" api:"required" format:"uuid"`
	// Type of authentication used for the webhook.
	AuthenticationType CasbPostureWebhookListResponseAuthenticationType `json:"authentication_type" api:"required"`
	// Timestamp when the webhook configuration was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Target URL for the webhook configuration. Where resulting data will be sent.
	DestinationURL string `json:"destination_url" api:"required" format:"uri"`
	// Account-specified display label for the webhook configuration.
	Label string `json:"label" api:"required"`
	// Current status of the webhook configuration. If disabled, data cannot be sent
	// through this configuration.
	Status CasbPostureWebhookListResponseStatus `json:"status" api:"required"`
	// Timestamp when the webhook configuration was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Version number of the configuration.
	Version int64 `json:"version" api:"required"`
	// List of header keys configured for this webhook. Values are not included for
	// security reasons.
	Headers []CasbPostureWebhookListResponseHeader `json:"headers"`
	JSON    casbPostureWebhookListResponseJSON     `json:"-"`
}

// casbPostureWebhookListResponseJSON contains the JSON metadata for the struct
// [CasbPostureWebhookListResponse]
type casbPostureWebhookListResponseJSON struct {
	ID                 apijson.Field
	AuthenticationType apijson.Field
	CreatedAt          apijson.Field
	DestinationURL     apijson.Field
	Label              apijson.Field
	Status             apijson.Field
	UpdatedAt          apijson.Field
	Version            apijson.Field
	Headers            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CasbPostureWebhookListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookListResponseJSON) RawJSON() string {
	return r.raw
}

// Type of authentication used for the webhook.
type CasbPostureWebhookListResponseAuthenticationType string

const (
	CasbPostureWebhookListResponseAuthenticationTypeBasicAuth     CasbPostureWebhookListResponseAuthenticationType = "Basic Auth"
	CasbPostureWebhookListResponseAuthenticationTypeNone          CasbPostureWebhookListResponseAuthenticationType = "None"
	CasbPostureWebhookListResponseAuthenticationTypeBearerAuth    CasbPostureWebhookListResponseAuthenticationType = "Bearer Auth"
	CasbPostureWebhookListResponseAuthenticationTypeStaticHeaders CasbPostureWebhookListResponseAuthenticationType = "Static Headers"
	CasbPostureWebhookListResponseAuthenticationTypeHmacSigning   CasbPostureWebhookListResponseAuthenticationType = "HMAC-Signing"
)

func (r CasbPostureWebhookListResponseAuthenticationType) IsKnown() bool {
	switch r {
	case CasbPostureWebhookListResponseAuthenticationTypeBasicAuth, CasbPostureWebhookListResponseAuthenticationTypeNone, CasbPostureWebhookListResponseAuthenticationTypeBearerAuth, CasbPostureWebhookListResponseAuthenticationTypeStaticHeaders, CasbPostureWebhookListResponseAuthenticationTypeHmacSigning:
		return true
	}
	return false
}

// Current status of the webhook configuration. If disabled, data cannot be sent
// through this configuration.
type CasbPostureWebhookListResponseStatus string

const (
	CasbPostureWebhookListResponseStatusEnabled  CasbPostureWebhookListResponseStatus = "enabled"
	CasbPostureWebhookListResponseStatusDisabled CasbPostureWebhookListResponseStatus = "disabled"
)

func (r CasbPostureWebhookListResponseStatus) IsKnown() bool {
	switch r {
	case CasbPostureWebhookListResponseStatusEnabled, CasbPostureWebhookListResponseStatusDisabled:
		return true
	}
	return false
}

type CasbPostureWebhookListResponseHeader struct {
	// Header key name (lowercase).
	Key  string                                   `json:"key"`
	JSON casbPostureWebhookListResponseHeaderJSON `json:"-"`
}

// casbPostureWebhookListResponseHeaderJSON contains the JSON metadata for the
// struct [CasbPostureWebhookListResponseHeader]
type casbPostureWebhookListResponseHeaderJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookListResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookListResponseHeaderJSON) RawJSON() string {
	return r.raw
}

// Common response structure for all API endpoints.
type CasbPostureWebhookDeleteResponse struct {
	Errors   []CasbPostureWebhookDeleteResponseError   `json:"errors" api:"required"`
	Messages []CasbPostureWebhookDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool                                 `json:"success" api:"required"`
	JSON    casbPostureWebhookDeleteResponseJSON `json:"-"`
}

// casbPostureWebhookDeleteResponseJSON contains the JSON metadata for the struct
// [CasbPostureWebhookDeleteResponse]
type casbPostureWebhookDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookDeleteResponseError struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                       `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookDeleteResponseErrorsSource `json:"source"`
	JSON             casbPostureWebhookDeleteResponseErrorJSON    `json:"-"`
}

// casbPostureWebhookDeleteResponseErrorJSON contains the JSON metadata for the
// struct [CasbPostureWebhookDeleteResponseError]
type casbPostureWebhookDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookDeleteResponseErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                           `json:"pointer"`
	JSON    casbPostureWebhookDeleteResponseErrorsSourceJSON `json:"-"`
}

// casbPostureWebhookDeleteResponseErrorsSourceJSON contains the JSON metadata for
// the struct [CasbPostureWebhookDeleteResponseErrorsSource]
type casbPostureWebhookDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookDeleteResponseMessage struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                         `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookDeleteResponseMessagesSource `json:"source"`
	JSON             casbPostureWebhookDeleteResponseMessageJSON    `json:"-"`
}

// casbPostureWebhookDeleteResponseMessageJSON contains the JSON metadata for the
// struct [CasbPostureWebhookDeleteResponseMessage]
type casbPostureWebhookDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookDeleteResponseMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                             `json:"pointer"`
	JSON    casbPostureWebhookDeleteResponseMessagesSourceJSON `json:"-"`
}

// casbPostureWebhookDeleteResponseMessagesSourceJSON contains the JSON metadata
// for the struct [CasbPostureWebhookDeleteResponseMessagesSource]
type casbPostureWebhookDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Response body for webhook evaluation test results.
type CasbPostureWebhookEvaluateResponse struct {
	// Human-readable message describing the test result.
	Message string `json:"message" api:"required"`
	// HTTP status code returned by the webhook endpoint. 0 if connection failed.
	StatusCode int64 `json:"status_code" api:"required"`
	// Whether the webhook test was successful (received 2xx response).
	Success bool                                   `json:"success" api:"required"`
	JSON    casbPostureWebhookEvaluateResponseJSON `json:"-"`
}

// casbPostureWebhookEvaluateResponseJSON contains the JSON metadata for the struct
// [CasbPostureWebhookEvaluateResponse]
type casbPostureWebhookEvaluateResponseJSON struct {
	Message     apijson.Field
	StatusCode  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookEvaluateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookEvaluateResponseJSON) RawJSON() string {
	return r.raw
}

// Response body for webhook evaluation test results.
type CasbPostureWebhookEvaluateExistingResponse struct {
	// Human-readable message describing the test result.
	Message string `json:"message" api:"required"`
	// HTTP status code returned by the webhook endpoint. 0 if connection failed.
	StatusCode int64 `json:"status_code" api:"required"`
	// Whether the webhook test was successful (received 2xx response).
	Success bool                                           `json:"success" api:"required"`
	JSON    casbPostureWebhookEvaluateExistingResponseJSON `json:"-"`
}

// casbPostureWebhookEvaluateExistingResponseJSON contains the JSON metadata for
// the struct [CasbPostureWebhookEvaluateExistingResponse]
type casbPostureWebhookEvaluateExistingResponseJSON struct {
	Message     apijson.Field
	StatusCode  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookEvaluateExistingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookEvaluateExistingResponseJSON) RawJSON() string {
	return r.raw
}

// Webhook configuration for sending finding notifications.
type CasbPostureWebhookGetResponse struct {
	// Unique identifier for the specific webhook configuration.
	ID string `json:"id" api:"required" format:"uuid"`
	// Type of authentication used for the webhook.
	AuthenticationType CasbPostureWebhookGetResponseAuthenticationType `json:"authentication_type" api:"required"`
	// Timestamp when the webhook configuration was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Target URL for the webhook configuration. Where resulting data will be sent.
	DestinationURL string `json:"destination_url" api:"required" format:"uri"`
	// Account-specified display label for the webhook configuration.
	Label string `json:"label" api:"required"`
	// Current status of the webhook configuration. If disabled, data cannot be sent
	// through this configuration.
	Status CasbPostureWebhookGetResponseStatus `json:"status" api:"required"`
	// Timestamp when the webhook configuration was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Version number of the configuration.
	Version int64 `json:"version" api:"required"`
	// List of header keys configured for this webhook. Values are not included for
	// security reasons.
	Headers []CasbPostureWebhookGetResponseHeader `json:"headers"`
	JSON    casbPostureWebhookGetResponseJSON     `json:"-"`
}

// casbPostureWebhookGetResponseJSON contains the JSON metadata for the struct
// [CasbPostureWebhookGetResponse]
type casbPostureWebhookGetResponseJSON struct {
	ID                 apijson.Field
	AuthenticationType apijson.Field
	CreatedAt          apijson.Field
	DestinationURL     apijson.Field
	Label              apijson.Field
	Status             apijson.Field
	UpdatedAt          apijson.Field
	Version            apijson.Field
	Headers            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CasbPostureWebhookGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookGetResponseJSON) RawJSON() string {
	return r.raw
}

// Type of authentication used for the webhook.
type CasbPostureWebhookGetResponseAuthenticationType string

const (
	CasbPostureWebhookGetResponseAuthenticationTypeBasicAuth     CasbPostureWebhookGetResponseAuthenticationType = "Basic Auth"
	CasbPostureWebhookGetResponseAuthenticationTypeNone          CasbPostureWebhookGetResponseAuthenticationType = "None"
	CasbPostureWebhookGetResponseAuthenticationTypeBearerAuth    CasbPostureWebhookGetResponseAuthenticationType = "Bearer Auth"
	CasbPostureWebhookGetResponseAuthenticationTypeStaticHeaders CasbPostureWebhookGetResponseAuthenticationType = "Static Headers"
	CasbPostureWebhookGetResponseAuthenticationTypeHmacSigning   CasbPostureWebhookGetResponseAuthenticationType = "HMAC-Signing"
)

func (r CasbPostureWebhookGetResponseAuthenticationType) IsKnown() bool {
	switch r {
	case CasbPostureWebhookGetResponseAuthenticationTypeBasicAuth, CasbPostureWebhookGetResponseAuthenticationTypeNone, CasbPostureWebhookGetResponseAuthenticationTypeBearerAuth, CasbPostureWebhookGetResponseAuthenticationTypeStaticHeaders, CasbPostureWebhookGetResponseAuthenticationTypeHmacSigning:
		return true
	}
	return false
}

// Current status of the webhook configuration. If disabled, data cannot be sent
// through this configuration.
type CasbPostureWebhookGetResponseStatus string

const (
	CasbPostureWebhookGetResponseStatusEnabled  CasbPostureWebhookGetResponseStatus = "enabled"
	CasbPostureWebhookGetResponseStatusDisabled CasbPostureWebhookGetResponseStatus = "disabled"
)

func (r CasbPostureWebhookGetResponseStatus) IsKnown() bool {
	switch r {
	case CasbPostureWebhookGetResponseStatusEnabled, CasbPostureWebhookGetResponseStatusDisabled:
		return true
	}
	return false
}

type CasbPostureWebhookGetResponseHeader struct {
	// Header key name (lowercase).
	Key  string                                  `json:"key"`
	JSON casbPostureWebhookGetResponseHeaderJSON `json:"-"`
}

// casbPostureWebhookGetResponseHeaderJSON contains the JSON metadata for the
// struct [CasbPostureWebhookGetResponseHeader]
type casbPostureWebhookGetResponseHeaderJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookGetResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookGetResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Type of authentication used for the webhook.
	AuthenticationType param.Field[CasbPostureWebhookNewParamsAuthenticationType] `json:"authentication_type" api:"required"`
	// Target URL for the webhook configuration. Where resulting data will be sent.
	DestinationURL param.Field[string] `json:"destination_url" api:"required" format:"uri"`
	// Account-specified display label for the webhook configuration.
	Label param.Field[string] `json:"label" api:"required"`
	// List of custom headers to include in webhook requests.
	Headers param.Field[[]CasbPostureWebhookNewParamsHeader] `json:"headers"`
	// Secret key used for HMAC signing when authentication_type is "HMAC-Signing".
	SigningSecret param.Field[string] `json:"signing_secret"`
}

func (r CasbPostureWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of authentication used for the webhook.
type CasbPostureWebhookNewParamsAuthenticationType string

const (
	CasbPostureWebhookNewParamsAuthenticationTypeBasicAuth     CasbPostureWebhookNewParamsAuthenticationType = "Basic Auth"
	CasbPostureWebhookNewParamsAuthenticationTypeNone          CasbPostureWebhookNewParamsAuthenticationType = "None"
	CasbPostureWebhookNewParamsAuthenticationTypeBearerAuth    CasbPostureWebhookNewParamsAuthenticationType = "Bearer Auth"
	CasbPostureWebhookNewParamsAuthenticationTypeStaticHeaders CasbPostureWebhookNewParamsAuthenticationType = "Static Headers"
	CasbPostureWebhookNewParamsAuthenticationTypeHmacSigning   CasbPostureWebhookNewParamsAuthenticationType = "HMAC-Signing"
)

func (r CasbPostureWebhookNewParamsAuthenticationType) IsKnown() bool {
	switch r {
	case CasbPostureWebhookNewParamsAuthenticationTypeBasicAuth, CasbPostureWebhookNewParamsAuthenticationTypeNone, CasbPostureWebhookNewParamsAuthenticationTypeBearerAuth, CasbPostureWebhookNewParamsAuthenticationTypeStaticHeaders, CasbPostureWebhookNewParamsAuthenticationTypeHmacSigning:
		return true
	}
	return false
}

// A header to include in webhook requests. On Create/Evaluate, both key and value
// are required. On Update, omitting value means "keep existing value".
type CasbPostureWebhookNewParamsHeader struct {
	// Header key name.
	Key param.Field[string] `json:"key" api:"required"`
	// Header value. Required on Create and Evaluate. On Update, omit or set to null to
	// keep existing value.
	Value param.Field[string] `json:"value"`
}

func (r CasbPostureWebhookNewParamsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Common response structure for all API endpoints.
type CasbPostureWebhookNewResponseEnvelope struct {
	Errors   []CasbPostureWebhookNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureWebhookNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Webhook configuration for sending finding notifications.
	Result CasbPostureWebhookNewResponse             `json:"result"`
	JSON   casbPostureWebhookNewResponseEnvelopeJSON `json:"-"`
}

// casbPostureWebhookNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPostureWebhookNewResponseEnvelope]
type casbPostureWebhookNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookNewResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                            `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureWebhookNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureWebhookNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CasbPostureWebhookNewResponseEnvelopeErrors]
type casbPostureWebhookNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookNewResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                `json:"pointer"`
	JSON    casbPostureWebhookNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureWebhookNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CasbPostureWebhookNewResponseEnvelopeErrorsSource]
type casbPostureWebhookNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookNewResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                              `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureWebhookNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureWebhookNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CasbPostureWebhookNewResponseEnvelopeMessages]
type casbPostureWebhookNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookNewResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                  `json:"pointer"`
	JSON    casbPostureWebhookNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureWebhookNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CasbPostureWebhookNewResponseEnvelopeMessagesSource]
type casbPostureWebhookNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookUpdateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Type of authentication used for the webhook.
	AuthenticationType param.Field[CasbPostureWebhookUpdateParamsAuthenticationType] `json:"authentication_type" api:"required"`
	// Target URL for the webhook configuration. Where resulting data will be sent.
	DestinationURL param.Field[string] `json:"destination_url" api:"required" format:"uri"`
	// Account-specified display label for the webhook configuration.
	Label param.Field[string] `json:"label" api:"required"`
	// Status of the webhook configuration.
	Status param.Field[CasbPostureWebhookUpdateParamsStatus] `json:"status" api:"required"`
	// List of custom headers to include in webhook requests.
	Headers param.Field[[]CasbPostureWebhookUpdateParamsHeader] `json:"headers"`
	// Secret key used for HMAC signing when authentication_type is "HMAC-Signing".
	SigningSecret param.Field[string] `json:"signing_secret"`
}

func (r CasbPostureWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of authentication used for the webhook.
type CasbPostureWebhookUpdateParamsAuthenticationType string

const (
	CasbPostureWebhookUpdateParamsAuthenticationTypeBasicAuth     CasbPostureWebhookUpdateParamsAuthenticationType = "Basic Auth"
	CasbPostureWebhookUpdateParamsAuthenticationTypeNone          CasbPostureWebhookUpdateParamsAuthenticationType = "None"
	CasbPostureWebhookUpdateParamsAuthenticationTypeBearerAuth    CasbPostureWebhookUpdateParamsAuthenticationType = "Bearer Auth"
	CasbPostureWebhookUpdateParamsAuthenticationTypeStaticHeaders CasbPostureWebhookUpdateParamsAuthenticationType = "Static Headers"
	CasbPostureWebhookUpdateParamsAuthenticationTypeHmacSigning   CasbPostureWebhookUpdateParamsAuthenticationType = "HMAC-Signing"
)

func (r CasbPostureWebhookUpdateParamsAuthenticationType) IsKnown() bool {
	switch r {
	case CasbPostureWebhookUpdateParamsAuthenticationTypeBasicAuth, CasbPostureWebhookUpdateParamsAuthenticationTypeNone, CasbPostureWebhookUpdateParamsAuthenticationTypeBearerAuth, CasbPostureWebhookUpdateParamsAuthenticationTypeStaticHeaders, CasbPostureWebhookUpdateParamsAuthenticationTypeHmacSigning:
		return true
	}
	return false
}

// Status of the webhook configuration.
type CasbPostureWebhookUpdateParamsStatus string

const (
	CasbPostureWebhookUpdateParamsStatusEnabled  CasbPostureWebhookUpdateParamsStatus = "enabled"
	CasbPostureWebhookUpdateParamsStatusDisabled CasbPostureWebhookUpdateParamsStatus = "disabled"
)

func (r CasbPostureWebhookUpdateParamsStatus) IsKnown() bool {
	switch r {
	case CasbPostureWebhookUpdateParamsStatusEnabled, CasbPostureWebhookUpdateParamsStatusDisabled:
		return true
	}
	return false
}

// A header to include in webhook requests. On Create/Evaluate, both key and value
// are required. On Update, omitting value means "keep existing value".
type CasbPostureWebhookUpdateParamsHeader struct {
	// Header key name.
	Key param.Field[string] `json:"key" api:"required"`
	// Header value. Required on Create and Evaluate. On Update, omit or set to null to
	// keep existing value.
	Value param.Field[string] `json:"value"`
}

func (r CasbPostureWebhookUpdateParamsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Common response structure for all API endpoints.
type CasbPostureWebhookUpdateResponseEnvelope struct {
	Errors   []CasbPostureWebhookUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureWebhookUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Webhook configuration for sending finding notifications.
	Result CasbPostureWebhookUpdateResponse             `json:"result"`
	JSON   casbPostureWebhookUpdateResponseEnvelopeJSON `json:"-"`
}

// casbPostureWebhookUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPostureWebhookUpdateResponseEnvelope]
type casbPostureWebhookUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookUpdateResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                               `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureWebhookUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureWebhookUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CasbPostureWebhookUpdateResponseEnvelopeErrors]
type casbPostureWebhookUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookUpdateResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                   `json:"pointer"`
	JSON    casbPostureWebhookUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureWebhookUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CasbPostureWebhookUpdateResponseEnvelopeErrorsSource]
type casbPostureWebhookUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookUpdateResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                 `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureWebhookUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureWebhookUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CasbPostureWebhookUpdateResponseEnvelopeMessages]
type casbPostureWebhookUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookUpdateResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                     `json:"pointer"`
	JSON    casbPostureWebhookUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureWebhookUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CasbPostureWebhookUpdateResponseEnvelopeMessagesSource]
type casbPostureWebhookUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CasbPostureWebhookDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CasbPostureWebhookEvaluateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Type of authentication to use for the test webhook request.
	AuthenticationType param.Field[CasbPostureWebhookEvaluateParamsAuthenticationType] `json:"authentication_type" api:"required"`
	// Target URL to send the test webhook event to.
	DestinationURL param.Field[string] `json:"destination_url" api:"required" format:"uri"`
	// List of custom headers to include in the test webhook request.
	Headers param.Field[[]CasbPostureWebhookEvaluateParamsHeader] `json:"headers"`
	// Secret key used for HMAC signing when authentication_type is "HMAC-Signing".
	SigningSecret param.Field[string] `json:"signing_secret"`
}

func (r CasbPostureWebhookEvaluateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of authentication to use for the test webhook request.
type CasbPostureWebhookEvaluateParamsAuthenticationType string

const (
	CasbPostureWebhookEvaluateParamsAuthenticationTypeBasicAuth     CasbPostureWebhookEvaluateParamsAuthenticationType = "Basic Auth"
	CasbPostureWebhookEvaluateParamsAuthenticationTypeNone          CasbPostureWebhookEvaluateParamsAuthenticationType = "None"
	CasbPostureWebhookEvaluateParamsAuthenticationTypeBearerAuth    CasbPostureWebhookEvaluateParamsAuthenticationType = "Bearer Auth"
	CasbPostureWebhookEvaluateParamsAuthenticationTypeStaticHeaders CasbPostureWebhookEvaluateParamsAuthenticationType = "Static Headers"
	CasbPostureWebhookEvaluateParamsAuthenticationTypeHmacSigning   CasbPostureWebhookEvaluateParamsAuthenticationType = "HMAC-Signing"
)

func (r CasbPostureWebhookEvaluateParamsAuthenticationType) IsKnown() bool {
	switch r {
	case CasbPostureWebhookEvaluateParamsAuthenticationTypeBasicAuth, CasbPostureWebhookEvaluateParamsAuthenticationTypeNone, CasbPostureWebhookEvaluateParamsAuthenticationTypeBearerAuth, CasbPostureWebhookEvaluateParamsAuthenticationTypeStaticHeaders, CasbPostureWebhookEvaluateParamsAuthenticationTypeHmacSigning:
		return true
	}
	return false
}

// A header to include in webhook requests. On Create/Evaluate, both key and value
// are required. On Update, omitting value means "keep existing value".
type CasbPostureWebhookEvaluateParamsHeader struct {
	// Header key name.
	Key param.Field[string] `json:"key" api:"required"`
	// Header value. Required on Create and Evaluate. On Update, omit or set to null to
	// keep existing value.
	Value param.Field[string] `json:"value"`
}

func (r CasbPostureWebhookEvaluateParamsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Common response structure for all API endpoints.
type CasbPostureWebhookEvaluateResponseEnvelope struct {
	Errors   []CasbPostureWebhookEvaluateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureWebhookEvaluateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Response body for webhook evaluation test results.
	Result CasbPostureWebhookEvaluateResponse             `json:"result"`
	JSON   casbPostureWebhookEvaluateResponseEnvelopeJSON `json:"-"`
}

// casbPostureWebhookEvaluateResponseEnvelopeJSON contains the JSON metadata for
// the struct [CasbPostureWebhookEvaluateResponseEnvelope]
type casbPostureWebhookEvaluateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookEvaluateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookEvaluateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookEvaluateResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                 `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookEvaluateResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureWebhookEvaluateResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureWebhookEvaluateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CasbPostureWebhookEvaluateResponseEnvelopeErrors]
type casbPostureWebhookEvaluateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookEvaluateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookEvaluateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookEvaluateResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                     `json:"pointer"`
	JSON    casbPostureWebhookEvaluateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureWebhookEvaluateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CasbPostureWebhookEvaluateResponseEnvelopeErrorsSource]
type casbPostureWebhookEvaluateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookEvaluateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookEvaluateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookEvaluateResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                   `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookEvaluateResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureWebhookEvaluateResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureWebhookEvaluateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CasbPostureWebhookEvaluateResponseEnvelopeMessages]
type casbPostureWebhookEvaluateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookEvaluateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookEvaluateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookEvaluateResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                       `json:"pointer"`
	JSON    casbPostureWebhookEvaluateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureWebhookEvaluateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [CasbPostureWebhookEvaluateResponseEnvelopeMessagesSource]
type casbPostureWebhookEvaluateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookEvaluateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookEvaluateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookEvaluateExistingParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

// Common response structure for all API endpoints.
type CasbPostureWebhookEvaluateExistingResponseEnvelope struct {
	Errors   []CasbPostureWebhookEvaluateExistingResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureWebhookEvaluateExistingResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Response body for webhook evaluation test results.
	Result CasbPostureWebhookEvaluateExistingResponse             `json:"result"`
	JSON   casbPostureWebhookEvaluateExistingResponseEnvelopeJSON `json:"-"`
}

// casbPostureWebhookEvaluateExistingResponseEnvelopeJSON contains the JSON
// metadata for the struct [CasbPostureWebhookEvaluateExistingResponseEnvelope]
type casbPostureWebhookEvaluateExistingResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookEvaluateExistingResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookEvaluateExistingResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookEvaluateExistingResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                         `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookEvaluateExistingResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureWebhookEvaluateExistingResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureWebhookEvaluateExistingResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CasbPostureWebhookEvaluateExistingResponseEnvelopeErrors]
type casbPostureWebhookEvaluateExistingResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookEvaluateExistingResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookEvaluateExistingResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookEvaluateExistingResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                             `json:"pointer"`
	JSON    casbPostureWebhookEvaluateExistingResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureWebhookEvaluateExistingResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [CasbPostureWebhookEvaluateExistingResponseEnvelopeErrorsSource]
type casbPostureWebhookEvaluateExistingResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookEvaluateExistingResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookEvaluateExistingResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookEvaluateExistingResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                           `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookEvaluateExistingResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureWebhookEvaluateExistingResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureWebhookEvaluateExistingResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CasbPostureWebhookEvaluateExistingResponseEnvelopeMessages]
type casbPostureWebhookEvaluateExistingResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookEvaluateExistingResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookEvaluateExistingResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookEvaluateExistingResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                               `json:"pointer"`
	JSON    casbPostureWebhookEvaluateExistingResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureWebhookEvaluateExistingResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [CasbPostureWebhookEvaluateExistingResponseEnvelopeMessagesSource]
type casbPostureWebhookEvaluateExistingResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookEvaluateExistingResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookEvaluateExistingResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

// Common response structure for all API endpoints.
type CasbPostureWebhookGetResponseEnvelope struct {
	Errors   []CasbPostureWebhookGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureWebhookGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Webhook configuration for sending finding notifications.
	Result CasbPostureWebhookGetResponse             `json:"result"`
	JSON   casbPostureWebhookGetResponseEnvelopeJSON `json:"-"`
}

// casbPostureWebhookGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPostureWebhookGetResponseEnvelope]
type casbPostureWebhookGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookGetResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                            `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureWebhookGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureWebhookGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CasbPostureWebhookGetResponseEnvelopeErrors]
type casbPostureWebhookGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookGetResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                `json:"pointer"`
	JSON    casbPostureWebhookGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureWebhookGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CasbPostureWebhookGetResponseEnvelopeErrorsSource]
type casbPostureWebhookGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookGetResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                              `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureWebhookGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureWebhookGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CasbPostureWebhookGetResponseEnvelopeMessages]
type casbPostureWebhookGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookGetResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                  `json:"pointer"`
	JSON    casbPostureWebhookGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureWebhookGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CasbPostureWebhookGetResponseEnvelopeMessagesSource]
type casbPostureWebhookGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
