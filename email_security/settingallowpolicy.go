// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

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

// SettingAllowPolicyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingAllowPolicyService] method instead.
type SettingAllowPolicyService struct {
	Options []option.RequestOption
}

// NewSettingAllowPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingAllowPolicyService(opts ...option.RequestOption) (r *SettingAllowPolicyService) {
	r = &SettingAllowPolicyService{}
	r.Options = opts
	return
}

// Creates a new allow policy that exempts matching emails from security
// detections. Use with caution as this bypasses email security scanning. Policies
// can match on sender patterns and apply to specific detections or all detections.
func (r *SettingAllowPolicyService) New(ctx context.Context, params SettingAllowPolicyNewParams, opts ...option.RequestOption) (res *SettingAllowPolicyNewResponse, err error) {
	var env SettingAllowPolicyNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of email allow policies. These policies exempt matching
// emails from security detection, allowing them to bypass disposition actions.
// Supports filtering by pattern type and policy attributes.
func (r *SettingAllowPolicyService) List(ctx context.Context, params SettingAllowPolicyListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingAllowPolicyListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies", params.AccountID)
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

// Returns a paginated list of email allow policies. These policies exempt matching
// emails from security detection, allowing them to bypass disposition actions.
// Supports filtering by pattern type and policy attributes.
func (r *SettingAllowPolicyService) ListAutoPaging(ctx context.Context, params SettingAllowPolicyListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingAllowPolicyListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Removes an allow policy. After deletion, emails matching this pattern will be
// subject to normal security scanning and disposition actions.
func (r *SettingAllowPolicyService) Delete(ctx context.Context, policyID string, body SettingAllowPolicyDeleteParams, opts ...option.RequestOption) (res *SettingAllowPolicyDeleteResponse, err error) {
	var env SettingAllowPolicyDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies/%s", body.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Executes multiple operations atomically. All four operation arrays (deletes,
// patches, puts, posts) are required and executed in order. Send empty arrays for
// unused operations.
func (r *SettingAllowPolicyService) Batch(ctx context.Context, params SettingAllowPolicyBatchParams, opts ...option.RequestOption) (res *SettingAllowPolicyBatchResponse, err error) {
	var env SettingAllowPolicyBatchResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies/batch", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an existing allow policy. Only provided fields will be modified. Changes
// take effect for new emails matching the pattern.
func (r *SettingAllowPolicyService) Edit(ctx context.Context, policyID string, params SettingAllowPolicyEditParams, opts ...option.RequestOption) (res *SettingAllowPolicyEditResponse, err error) {
	var env SettingAllowPolicyEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves details for a specific allow policy including its pattern,
// dispositions that are exempted, and whether it applies to all detections.
func (r *SettingAllowPolicyService) Get(ctx context.Context, policyID string, query SettingAllowPolicyGetParams, opts ...option.RequestOption) (res *SettingAllowPolicyGetResponse, err error) {
	var env SettingAllowPolicyGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies/%s", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// An email allow policy.
type SettingAllowPolicyNewResponse struct {
	// Allow policy identifier.
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified time.Time `json:"last_modified" api:"required" format:"date-time"`
	Comments     string    `json:"comments" api:"nullable"`
	// Exempts messages from this sender from Spam, Spoof and Bulk dispositions only;
	// Malicious and Suspicious dispositions still apply.
	IsAcceptableSender bool `json:"is_acceptable_sender"`
	// Bypasses all detections for messages to this recipient.
	IsExemptRecipient bool `json:"is_exempt_recipient"`
	// Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_exempt_recipient` instead.
	IsRecipient bool `json:"is_recipient"`
	IsRegex     bool `json:"is_regex"`
	// Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_trusted_sender` instead.
	IsSender bool `json:"is_sender"`
	// Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_acceptable_sender` instead.
	IsSpoof bool `json:"is_spoof"`
	// Bypasses all detections and link following for messages from this sender.
	IsTrustedSender bool      `json:"is_trusted_sender"`
	ModifiedAt      time.Time `json:"modified_at" format:"date-time"`
	// The pattern value to match. The format depends on `pattern_type`: a valid email
	// address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN
	// (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g.
	// `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the
	// API rejects private or unique-local, loopback, link-local, unspecified, and IPv4
	// broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	Pattern string `json:"pattern"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
	//   `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
	//   `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
	//   link-local, unspecified, and IPv4 broadcast addresses, including their
	//   IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
	//   but it may appear on existing entries.
	PatternType SettingAllowPolicyNewResponsePatternType `json:"pattern_type"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool                              `json:"verify_sender"`
	JSON         settingAllowPolicyNewResponseJSON `json:"-"`
}

// settingAllowPolicyNewResponseJSON contains the JSON metadata for the struct
// [SettingAllowPolicyNewResponse]
type settingAllowPolicyNewResponseJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	LastModified       apijson.Field
	Comments           apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRecipient        apijson.Field
	IsRegex            apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	IsTrustedSender    apijson.Field
	ModifiedAt         apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingAllowPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

// Type of pattern matching.
//
//   - EMAIL: matches a full email address (e.g. `user@example.com`)
//   - DOMAIN: matches a domain name (e.g. `example.com`)
//   - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
//     `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
//     `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
//     link-local, unspecified, and IPv4 broadcast addresses, including their
//     IPv4-mapped IPv6 equivalents.
//   - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
//     but it may appear on existing entries.
type SettingAllowPolicyNewResponsePatternType string

const (
	SettingAllowPolicyNewResponsePatternTypeEmail   SettingAllowPolicyNewResponsePatternType = "EMAIL"
	SettingAllowPolicyNewResponsePatternTypeDomain  SettingAllowPolicyNewResponsePatternType = "DOMAIN"
	SettingAllowPolicyNewResponsePatternTypeIP      SettingAllowPolicyNewResponsePatternType = "IP"
	SettingAllowPolicyNewResponsePatternTypeUnknown SettingAllowPolicyNewResponsePatternType = "UNKNOWN"
)

func (r SettingAllowPolicyNewResponsePatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyNewResponsePatternTypeEmail, SettingAllowPolicyNewResponsePatternTypeDomain, SettingAllowPolicyNewResponsePatternTypeIP, SettingAllowPolicyNewResponsePatternTypeUnknown:
		return true
	}
	return false
}

// An email allow policy.
type SettingAllowPolicyListResponse struct {
	// Allow policy identifier.
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified time.Time `json:"last_modified" api:"required" format:"date-time"`
	Comments     string    `json:"comments" api:"nullable"`
	// Exempts messages from this sender from Spam, Spoof and Bulk dispositions only;
	// Malicious and Suspicious dispositions still apply.
	IsAcceptableSender bool `json:"is_acceptable_sender"`
	// Bypasses all detections for messages to this recipient.
	IsExemptRecipient bool `json:"is_exempt_recipient"`
	// Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_exempt_recipient` instead.
	IsRecipient bool `json:"is_recipient"`
	IsRegex     bool `json:"is_regex"`
	// Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_trusted_sender` instead.
	IsSender bool `json:"is_sender"`
	// Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_acceptable_sender` instead.
	IsSpoof bool `json:"is_spoof"`
	// Bypasses all detections and link following for messages from this sender.
	IsTrustedSender bool      `json:"is_trusted_sender"`
	ModifiedAt      time.Time `json:"modified_at" format:"date-time"`
	// The pattern value to match. The format depends on `pattern_type`: a valid email
	// address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN
	// (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g.
	// `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the
	// API rejects private or unique-local, loopback, link-local, unspecified, and IPv4
	// broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	Pattern string `json:"pattern"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
	//   `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
	//   `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
	//   link-local, unspecified, and IPv4 broadcast addresses, including their
	//   IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
	//   but it may appear on existing entries.
	PatternType SettingAllowPolicyListResponsePatternType `json:"pattern_type"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool                               `json:"verify_sender"`
	JSON         settingAllowPolicyListResponseJSON `json:"-"`
}

// settingAllowPolicyListResponseJSON contains the JSON metadata for the struct
// [SettingAllowPolicyListResponse]
type settingAllowPolicyListResponseJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	LastModified       apijson.Field
	Comments           apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRecipient        apijson.Field
	IsRegex            apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	IsTrustedSender    apijson.Field
	ModifiedAt         apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingAllowPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

// Type of pattern matching.
//
//   - EMAIL: matches a full email address (e.g. `user@example.com`)
//   - DOMAIN: matches a domain name (e.g. `example.com`)
//   - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
//     `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
//     `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
//     link-local, unspecified, and IPv4 broadcast addresses, including their
//     IPv4-mapped IPv6 equivalents.
//   - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
//     but it may appear on existing entries.
type SettingAllowPolicyListResponsePatternType string

const (
	SettingAllowPolicyListResponsePatternTypeEmail   SettingAllowPolicyListResponsePatternType = "EMAIL"
	SettingAllowPolicyListResponsePatternTypeDomain  SettingAllowPolicyListResponsePatternType = "DOMAIN"
	SettingAllowPolicyListResponsePatternTypeIP      SettingAllowPolicyListResponsePatternType = "IP"
	SettingAllowPolicyListResponsePatternTypeUnknown SettingAllowPolicyListResponsePatternType = "UNKNOWN"
)

func (r SettingAllowPolicyListResponsePatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyListResponsePatternTypeEmail, SettingAllowPolicyListResponsePatternTypeDomain, SettingAllowPolicyListResponsePatternTypeIP, SettingAllowPolicyListResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPolicyDeleteResponse struct {
	// Allow policy identifier.
	ID   string                               `json:"id" api:"required" format:"uuid"`
	JSON settingAllowPolicyDeleteResponseJSON `json:"-"`
}

// settingAllowPolicyDeleteResponseJSON contains the JSON metadata for the struct
// [SettingAllowPolicyDeleteResponse]
type settingAllowPolicyDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyBatchResponse struct {
	Deletes []SettingAllowPolicyBatchResponseDelete `json:"deletes"`
	Patches []SettingAllowPolicyBatchResponsePatch  `json:"patches"`
	Posts   []SettingAllowPolicyBatchResponsePost   `json:"posts"`
	Puts    []SettingAllowPolicyBatchResponsePut    `json:"puts"`
	JSON    settingAllowPolicyBatchResponseJSON     `json:"-"`
}

// settingAllowPolicyBatchResponseJSON contains the JSON metadata for the struct
// [SettingAllowPolicyBatchResponse]
type settingAllowPolicyBatchResponseJSON struct {
	Deletes     apijson.Field
	Patches     apijson.Field
	Posts       apijson.Field
	Puts        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyBatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyBatchResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyBatchResponseDelete struct {
	// Allow policy identifier.
	ID   string                                    `json:"id" api:"required" format:"uuid"`
	JSON settingAllowPolicyBatchResponseDeleteJSON `json:"-"`
}

// settingAllowPolicyBatchResponseDeleteJSON contains the JSON metadata for the
// struct [SettingAllowPolicyBatchResponseDelete]
type settingAllowPolicyBatchResponseDeleteJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyBatchResponseDelete) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyBatchResponseDeleteJSON) RawJSON() string {
	return r.raw
}

// An email allow policy.
type SettingAllowPolicyBatchResponsePatch struct {
	// Allow policy identifier.
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified time.Time `json:"last_modified" api:"required" format:"date-time"`
	Comments     string    `json:"comments" api:"nullable"`
	// Exempts messages from this sender from Spam, Spoof and Bulk dispositions only;
	// Malicious and Suspicious dispositions still apply.
	IsAcceptableSender bool `json:"is_acceptable_sender"`
	// Bypasses all detections for messages to this recipient.
	IsExemptRecipient bool `json:"is_exempt_recipient"`
	// Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_exempt_recipient` instead.
	IsRecipient bool `json:"is_recipient"`
	IsRegex     bool `json:"is_regex"`
	// Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_trusted_sender` instead.
	IsSender bool `json:"is_sender"`
	// Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_acceptable_sender` instead.
	IsSpoof bool `json:"is_spoof"`
	// Bypasses all detections and link following for messages from this sender.
	IsTrustedSender bool      `json:"is_trusted_sender"`
	ModifiedAt      time.Time `json:"modified_at" format:"date-time"`
	// The pattern value to match. The format depends on `pattern_type`: a valid email
	// address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN
	// (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g.
	// `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the
	// API rejects private or unique-local, loopback, link-local, unspecified, and IPv4
	// broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	Pattern string `json:"pattern"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
	//   `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
	//   `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
	//   link-local, unspecified, and IPv4 broadcast addresses, including their
	//   IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
	//   but it may appear on existing entries.
	PatternType SettingAllowPolicyBatchResponsePatchesPatternType `json:"pattern_type"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool                                     `json:"verify_sender"`
	JSON         settingAllowPolicyBatchResponsePatchJSON `json:"-"`
}

// settingAllowPolicyBatchResponsePatchJSON contains the JSON metadata for the
// struct [SettingAllowPolicyBatchResponsePatch]
type settingAllowPolicyBatchResponsePatchJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	LastModified       apijson.Field
	Comments           apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRecipient        apijson.Field
	IsRegex            apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	IsTrustedSender    apijson.Field
	ModifiedAt         apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingAllowPolicyBatchResponsePatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyBatchResponsePatchJSON) RawJSON() string {
	return r.raw
}

// Type of pattern matching.
//
//   - EMAIL: matches a full email address (e.g. `user@example.com`)
//   - DOMAIN: matches a domain name (e.g. `example.com`)
//   - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
//     `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
//     `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
//     link-local, unspecified, and IPv4 broadcast addresses, including their
//     IPv4-mapped IPv6 equivalents.
//   - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
//     but it may appear on existing entries.
type SettingAllowPolicyBatchResponsePatchesPatternType string

const (
	SettingAllowPolicyBatchResponsePatchesPatternTypeEmail   SettingAllowPolicyBatchResponsePatchesPatternType = "EMAIL"
	SettingAllowPolicyBatchResponsePatchesPatternTypeDomain  SettingAllowPolicyBatchResponsePatchesPatternType = "DOMAIN"
	SettingAllowPolicyBatchResponsePatchesPatternTypeIP      SettingAllowPolicyBatchResponsePatchesPatternType = "IP"
	SettingAllowPolicyBatchResponsePatchesPatternTypeUnknown SettingAllowPolicyBatchResponsePatchesPatternType = "UNKNOWN"
)

func (r SettingAllowPolicyBatchResponsePatchesPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyBatchResponsePatchesPatternTypeEmail, SettingAllowPolicyBatchResponsePatchesPatternTypeDomain, SettingAllowPolicyBatchResponsePatchesPatternTypeIP, SettingAllowPolicyBatchResponsePatchesPatternTypeUnknown:
		return true
	}
	return false
}

// An email allow policy.
type SettingAllowPolicyBatchResponsePost struct {
	// Allow policy identifier.
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified time.Time `json:"last_modified" api:"required" format:"date-time"`
	Comments     string    `json:"comments" api:"nullable"`
	// Exempts messages from this sender from Spam, Spoof and Bulk dispositions only;
	// Malicious and Suspicious dispositions still apply.
	IsAcceptableSender bool `json:"is_acceptable_sender"`
	// Bypasses all detections for messages to this recipient.
	IsExemptRecipient bool `json:"is_exempt_recipient"`
	// Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_exempt_recipient` instead.
	IsRecipient bool `json:"is_recipient"`
	IsRegex     bool `json:"is_regex"`
	// Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_trusted_sender` instead.
	IsSender bool `json:"is_sender"`
	// Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_acceptable_sender` instead.
	IsSpoof bool `json:"is_spoof"`
	// Bypasses all detections and link following for messages from this sender.
	IsTrustedSender bool      `json:"is_trusted_sender"`
	ModifiedAt      time.Time `json:"modified_at" format:"date-time"`
	// The pattern value to match. The format depends on `pattern_type`: a valid email
	// address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN
	// (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g.
	// `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the
	// API rejects private or unique-local, loopback, link-local, unspecified, and IPv4
	// broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	Pattern string `json:"pattern"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
	//   `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
	//   `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
	//   link-local, unspecified, and IPv4 broadcast addresses, including their
	//   IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
	//   but it may appear on existing entries.
	PatternType SettingAllowPolicyBatchResponsePostsPatternType `json:"pattern_type"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool                                    `json:"verify_sender"`
	JSON         settingAllowPolicyBatchResponsePostJSON `json:"-"`
}

// settingAllowPolicyBatchResponsePostJSON contains the JSON metadata for the
// struct [SettingAllowPolicyBatchResponsePost]
type settingAllowPolicyBatchResponsePostJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	LastModified       apijson.Field
	Comments           apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRecipient        apijson.Field
	IsRegex            apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	IsTrustedSender    apijson.Field
	ModifiedAt         apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingAllowPolicyBatchResponsePost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyBatchResponsePostJSON) RawJSON() string {
	return r.raw
}

// Type of pattern matching.
//
//   - EMAIL: matches a full email address (e.g. `user@example.com`)
//   - DOMAIN: matches a domain name (e.g. `example.com`)
//   - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
//     `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
//     `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
//     link-local, unspecified, and IPv4 broadcast addresses, including their
//     IPv4-mapped IPv6 equivalents.
//   - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
//     but it may appear on existing entries.
type SettingAllowPolicyBatchResponsePostsPatternType string

const (
	SettingAllowPolicyBatchResponsePostsPatternTypeEmail   SettingAllowPolicyBatchResponsePostsPatternType = "EMAIL"
	SettingAllowPolicyBatchResponsePostsPatternTypeDomain  SettingAllowPolicyBatchResponsePostsPatternType = "DOMAIN"
	SettingAllowPolicyBatchResponsePostsPatternTypeIP      SettingAllowPolicyBatchResponsePostsPatternType = "IP"
	SettingAllowPolicyBatchResponsePostsPatternTypeUnknown SettingAllowPolicyBatchResponsePostsPatternType = "UNKNOWN"
)

func (r SettingAllowPolicyBatchResponsePostsPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyBatchResponsePostsPatternTypeEmail, SettingAllowPolicyBatchResponsePostsPatternTypeDomain, SettingAllowPolicyBatchResponsePostsPatternTypeIP, SettingAllowPolicyBatchResponsePostsPatternTypeUnknown:
		return true
	}
	return false
}

// An email allow policy.
type SettingAllowPolicyBatchResponsePut struct {
	// Allow policy identifier.
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified time.Time `json:"last_modified" api:"required" format:"date-time"`
	Comments     string    `json:"comments" api:"nullable"`
	// Exempts messages from this sender from Spam, Spoof and Bulk dispositions only;
	// Malicious and Suspicious dispositions still apply.
	IsAcceptableSender bool `json:"is_acceptable_sender"`
	// Bypasses all detections for messages to this recipient.
	IsExemptRecipient bool `json:"is_exempt_recipient"`
	// Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_exempt_recipient` instead.
	IsRecipient bool `json:"is_recipient"`
	IsRegex     bool `json:"is_regex"`
	// Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_trusted_sender` instead.
	IsSender bool `json:"is_sender"`
	// Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_acceptable_sender` instead.
	IsSpoof bool `json:"is_spoof"`
	// Bypasses all detections and link following for messages from this sender.
	IsTrustedSender bool      `json:"is_trusted_sender"`
	ModifiedAt      time.Time `json:"modified_at" format:"date-time"`
	// The pattern value to match. The format depends on `pattern_type`: a valid email
	// address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN
	// (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g.
	// `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the
	// API rejects private or unique-local, loopback, link-local, unspecified, and IPv4
	// broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	Pattern string `json:"pattern"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
	//   `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
	//   `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
	//   link-local, unspecified, and IPv4 broadcast addresses, including their
	//   IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
	//   but it may appear on existing entries.
	PatternType SettingAllowPolicyBatchResponsePutsPatternType `json:"pattern_type"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool                                   `json:"verify_sender"`
	JSON         settingAllowPolicyBatchResponsePutJSON `json:"-"`
}

// settingAllowPolicyBatchResponsePutJSON contains the JSON metadata for the struct
// [SettingAllowPolicyBatchResponsePut]
type settingAllowPolicyBatchResponsePutJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	LastModified       apijson.Field
	Comments           apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRecipient        apijson.Field
	IsRegex            apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	IsTrustedSender    apijson.Field
	ModifiedAt         apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingAllowPolicyBatchResponsePut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyBatchResponsePutJSON) RawJSON() string {
	return r.raw
}

// Type of pattern matching.
//
//   - EMAIL: matches a full email address (e.g. `user@example.com`)
//   - DOMAIN: matches a domain name (e.g. `example.com`)
//   - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
//     `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
//     `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
//     link-local, unspecified, and IPv4 broadcast addresses, including their
//     IPv4-mapped IPv6 equivalents.
//   - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
//     but it may appear on existing entries.
type SettingAllowPolicyBatchResponsePutsPatternType string

const (
	SettingAllowPolicyBatchResponsePutsPatternTypeEmail   SettingAllowPolicyBatchResponsePutsPatternType = "EMAIL"
	SettingAllowPolicyBatchResponsePutsPatternTypeDomain  SettingAllowPolicyBatchResponsePutsPatternType = "DOMAIN"
	SettingAllowPolicyBatchResponsePutsPatternTypeIP      SettingAllowPolicyBatchResponsePutsPatternType = "IP"
	SettingAllowPolicyBatchResponsePutsPatternTypeUnknown SettingAllowPolicyBatchResponsePutsPatternType = "UNKNOWN"
)

func (r SettingAllowPolicyBatchResponsePutsPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyBatchResponsePutsPatternTypeEmail, SettingAllowPolicyBatchResponsePutsPatternTypeDomain, SettingAllowPolicyBatchResponsePutsPatternTypeIP, SettingAllowPolicyBatchResponsePutsPatternTypeUnknown:
		return true
	}
	return false
}

// An email allow policy.
type SettingAllowPolicyEditResponse struct {
	// Allow policy identifier.
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified time.Time `json:"last_modified" api:"required" format:"date-time"`
	Comments     string    `json:"comments" api:"nullable"`
	// Exempts messages from this sender from Spam, Spoof and Bulk dispositions only;
	// Malicious and Suspicious dispositions still apply.
	IsAcceptableSender bool `json:"is_acceptable_sender"`
	// Bypasses all detections for messages to this recipient.
	IsExemptRecipient bool `json:"is_exempt_recipient"`
	// Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_exempt_recipient` instead.
	IsRecipient bool `json:"is_recipient"`
	IsRegex     bool `json:"is_regex"`
	// Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_trusted_sender` instead.
	IsSender bool `json:"is_sender"`
	// Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_acceptable_sender` instead.
	IsSpoof bool `json:"is_spoof"`
	// Bypasses all detections and link following for messages from this sender.
	IsTrustedSender bool      `json:"is_trusted_sender"`
	ModifiedAt      time.Time `json:"modified_at" format:"date-time"`
	// The pattern value to match. The format depends on `pattern_type`: a valid email
	// address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN
	// (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g.
	// `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the
	// API rejects private or unique-local, loopback, link-local, unspecified, and IPv4
	// broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	Pattern string `json:"pattern"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
	//   `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
	//   `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
	//   link-local, unspecified, and IPv4 broadcast addresses, including their
	//   IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
	//   but it may appear on existing entries.
	PatternType SettingAllowPolicyEditResponsePatternType `json:"pattern_type"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool                               `json:"verify_sender"`
	JSON         settingAllowPolicyEditResponseJSON `json:"-"`
}

// settingAllowPolicyEditResponseJSON contains the JSON metadata for the struct
// [SettingAllowPolicyEditResponse]
type settingAllowPolicyEditResponseJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	LastModified       apijson.Field
	Comments           apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRecipient        apijson.Field
	IsRegex            apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	IsTrustedSender    apijson.Field
	ModifiedAt         apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingAllowPolicyEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyEditResponseJSON) RawJSON() string {
	return r.raw
}

// Type of pattern matching.
//
//   - EMAIL: matches a full email address (e.g. `user@example.com`)
//   - DOMAIN: matches a domain name (e.g. `example.com`)
//   - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
//     `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
//     `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
//     link-local, unspecified, and IPv4 broadcast addresses, including their
//     IPv4-mapped IPv6 equivalents.
//   - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
//     but it may appear on existing entries.
type SettingAllowPolicyEditResponsePatternType string

const (
	SettingAllowPolicyEditResponsePatternTypeEmail   SettingAllowPolicyEditResponsePatternType = "EMAIL"
	SettingAllowPolicyEditResponsePatternTypeDomain  SettingAllowPolicyEditResponsePatternType = "DOMAIN"
	SettingAllowPolicyEditResponsePatternTypeIP      SettingAllowPolicyEditResponsePatternType = "IP"
	SettingAllowPolicyEditResponsePatternTypeUnknown SettingAllowPolicyEditResponsePatternType = "UNKNOWN"
)

func (r SettingAllowPolicyEditResponsePatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyEditResponsePatternTypeEmail, SettingAllowPolicyEditResponsePatternTypeDomain, SettingAllowPolicyEditResponsePatternTypeIP, SettingAllowPolicyEditResponsePatternTypeUnknown:
		return true
	}
	return false
}

// An email allow policy.
type SettingAllowPolicyGetResponse struct {
	// Allow policy identifier.
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: Use `modified_at` instead.
	LastModified time.Time `json:"last_modified" api:"required" format:"date-time"`
	Comments     string    `json:"comments" api:"nullable"`
	// Exempts messages from this sender from Spam, Spoof and Bulk dispositions only;
	// Malicious and Suspicious dispositions still apply.
	IsAcceptableSender bool `json:"is_acceptable_sender"`
	// Bypasses all detections for messages to this recipient.
	IsExemptRecipient bool `json:"is_exempt_recipient"`
	// Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_exempt_recipient` instead.
	IsRecipient bool `json:"is_recipient"`
	IsRegex     bool `json:"is_regex"`
	// Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_trusted_sender` instead.
	IsSender bool `json:"is_sender"`
	// Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_acceptable_sender` instead.
	IsSpoof bool `json:"is_spoof"`
	// Bypasses all detections and link following for messages from this sender.
	IsTrustedSender bool      `json:"is_trusted_sender"`
	ModifiedAt      time.Time `json:"modified_at" format:"date-time"`
	// The pattern value to match. The format depends on `pattern_type`: a valid email
	// address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN
	// (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g.
	// `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the
	// API rejects private or unique-local, loopback, link-local, unspecified, and IPv4
	// broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	Pattern string `json:"pattern"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
	//   `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
	//   `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
	//   link-local, unspecified, and IPv4 broadcast addresses, including their
	//   IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
	//   but it may appear on existing entries.
	PatternType SettingAllowPolicyGetResponsePatternType `json:"pattern_type"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool                              `json:"verify_sender"`
	JSON         settingAllowPolicyGetResponseJSON `json:"-"`
}

// settingAllowPolicyGetResponseJSON contains the JSON metadata for the struct
// [SettingAllowPolicyGetResponse]
type settingAllowPolicyGetResponseJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	LastModified       apijson.Field
	Comments           apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRecipient        apijson.Field
	IsRegex            apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	IsTrustedSender    apijson.Field
	ModifiedAt         apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingAllowPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

// Type of pattern matching.
//
//   - EMAIL: matches a full email address (e.g. `user@example.com`)
//   - DOMAIN: matches a domain name (e.g. `example.com`)
//   - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
//     `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
//     `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
//     link-local, unspecified, and IPv4 broadcast addresses, including their
//     IPv4-mapped IPv6 equivalents.
//   - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
//     but it may appear on existing entries.
type SettingAllowPolicyGetResponsePatternType string

const (
	SettingAllowPolicyGetResponsePatternTypeEmail   SettingAllowPolicyGetResponsePatternType = "EMAIL"
	SettingAllowPolicyGetResponsePatternTypeDomain  SettingAllowPolicyGetResponsePatternType = "DOMAIN"
	SettingAllowPolicyGetResponsePatternTypeIP      SettingAllowPolicyGetResponsePatternType = "IP"
	SettingAllowPolicyGetResponsePatternTypeUnknown SettingAllowPolicyGetResponsePatternType = "UNKNOWN"
)

func (r SettingAllowPolicyGetResponsePatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyGetResponsePatternTypeEmail, SettingAllowPolicyGetResponsePatternTypeDomain, SettingAllowPolicyGetResponsePatternTypeIP, SettingAllowPolicyGetResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPolicyNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Exempts messages from this sender from Spam, Spoof and Bulk dispositions only;
	// Malicious and Suspicious dispositions still apply.
	IsAcceptableSender param.Field[bool] `json:"is_acceptable_sender" api:"required"`
	// Bypasses all detections for messages to this recipient.
	IsExemptRecipient param.Field[bool] `json:"is_exempt_recipient" api:"required"`
	IsRegex           param.Field[bool] `json:"is_regex" api:"required"`
	// Bypasses all detections and link following for messages from this sender.
	IsTrustedSender param.Field[bool] `json:"is_trusted_sender" api:"required"`
	// The pattern value to match. The format depends on `pattern_type`: a valid email
	// address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN
	// (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g.
	// `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the
	// API rejects private or unique-local, loopback, link-local, unspecified, and IPv4
	// broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	Pattern param.Field[string] `json:"pattern" api:"required"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
	//   `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
	//   `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
	//   link-local, unspecified, and IPv4 broadcast addresses, including their
	//   IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
	//   but it may appear on existing entries.
	PatternType param.Field[SettingAllowPolicyNewParamsPatternType] `json:"pattern_type" api:"required"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender param.Field[bool]   `json:"verify_sender" api:"required"`
	Comments     param.Field[string] `json:"comments"`
	// Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life:
	// July 1, 2026.
	IsRecipient param.Field[bool] `json:"is_recipient"`
	// Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life:
	// July 1, 2026.
	IsSender param.Field[bool] `json:"is_sender"`
	// Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life:
	// July 1, 2026.
	IsSpoof param.Field[bool] `json:"is_spoof"`
}

func (r SettingAllowPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of pattern matching.
//
//   - EMAIL: matches a full email address (e.g. `user@example.com`)
//   - DOMAIN: matches a domain name (e.g. `example.com`)
//   - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
//     `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
//     `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
//     link-local, unspecified, and IPv4 broadcast addresses, including their
//     IPv4-mapped IPv6 equivalents.
//   - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
//     but it may appear on existing entries.
type SettingAllowPolicyNewParamsPatternType string

const (
	SettingAllowPolicyNewParamsPatternTypeEmail   SettingAllowPolicyNewParamsPatternType = "EMAIL"
	SettingAllowPolicyNewParamsPatternTypeDomain  SettingAllowPolicyNewParamsPatternType = "DOMAIN"
	SettingAllowPolicyNewParamsPatternTypeIP      SettingAllowPolicyNewParamsPatternType = "IP"
	SettingAllowPolicyNewParamsPatternTypeUnknown SettingAllowPolicyNewParamsPatternType = "UNKNOWN"
)

func (r SettingAllowPolicyNewParamsPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyNewParamsPatternTypeEmail, SettingAllowPolicyNewParamsPatternTypeDomain, SettingAllowPolicyNewParamsPatternTypeIP, SettingAllowPolicyNewParamsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPolicyNewResponseEnvelope struct {
	Errors   []SettingAllowPolicyNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingAllowPolicyNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingAllowPolicyNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// An email allow policy.
	Result SettingAllowPolicyNewResponse             `json:"result"`
	JSON   settingAllowPolicyNewResponseEnvelopeJSON `json:"-"`
}

// settingAllowPolicyNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAllowPolicyNewResponseEnvelope]
type settingAllowPolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyNewResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           SettingAllowPolicyNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingAllowPolicyNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingAllowPolicyNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingAllowPolicyNewResponseEnvelopeErrors]
type settingAllowPolicyNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAllowPolicyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    settingAllowPolicyNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingAllowPolicyNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SettingAllowPolicyNewResponseEnvelopeErrorsSource]
type settingAllowPolicyNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyNewResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           SettingAllowPolicyNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingAllowPolicyNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingAllowPolicyNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingAllowPolicyNewResponseEnvelopeMessages]
type settingAllowPolicyNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAllowPolicyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    settingAllowPolicyNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingAllowPolicyNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingAllowPolicyNewResponseEnvelopeMessagesSource]
type settingAllowPolicyNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingAllowPolicyNewResponseEnvelopeSuccess bool

const (
	SettingAllowPolicyNewResponseEnvelopeSuccessTrue SettingAllowPolicyNewResponseEnvelopeSuccess = true
)

func (r SettingAllowPolicyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingAllowPolicyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingAllowPolicyListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The sorting direction.
	Direction param.Field[SettingAllowPolicyListParamsDirection] `query:"direction"`
	// Filter to show only policies where messages from the sender are exempted from
	// Spam, Spoof, and Bulk dispositions (not Malicious or Suspicious).
	IsAcceptableSender param.Field[bool] `query:"is_acceptable_sender"`
	// Filter to show only policies where messages to the recipient bypass all
	// detections.
	IsExemptRecipient param.Field[bool] `query:"is_exempt_recipient"`
	// Filter to show only policies where messages from the sender bypass all
	// detections and link following.
	IsTrustedSender param.Field[bool] `query:"is_trusted_sender"`
	// Field to sort by.
	Order param.Field[SettingAllowPolicyListParamsOrder] `query:"order"`
	// Current page within paginated list of results.
	Page    param.Field[int64]  `query:"page"`
	Pattern param.Field[string] `query:"pattern"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
	//   `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
	//   `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
	//   link-local, unspecified, and IPv4 broadcast addresses, including their
	//   IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
	//   but it may appear on existing entries.
	PatternType param.Field[SettingAllowPolicyListParamsPatternType] `query:"pattern_type"`
	// The number of results per page. Maximum value is 1000.
	PerPage param.Field[int64] `query:"per_page"`
	// Search term for filtering records. Behavior may change.
	Search param.Field[string] `query:"search"`
	// Filter to show only policies that enforce DMARC, SPF, or DKIM authentication.
	VerifySender param.Field[bool] `query:"verify_sender"`
}

// URLQuery serializes [SettingAllowPolicyListParams]'s query parameters as
// `url.Values`.
func (r SettingAllowPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The sorting direction.
type SettingAllowPolicyListParamsDirection string

const (
	SettingAllowPolicyListParamsDirectionAsc  SettingAllowPolicyListParamsDirection = "asc"
	SettingAllowPolicyListParamsDirectionDesc SettingAllowPolicyListParamsDirection = "desc"
)

func (r SettingAllowPolicyListParamsDirection) IsKnown() bool {
	switch r {
	case SettingAllowPolicyListParamsDirectionAsc, SettingAllowPolicyListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to sort by.
type SettingAllowPolicyListParamsOrder string

const (
	SettingAllowPolicyListParamsOrderPattern   SettingAllowPolicyListParamsOrder = "pattern"
	SettingAllowPolicyListParamsOrderCreatedAt SettingAllowPolicyListParamsOrder = "created_at"
)

func (r SettingAllowPolicyListParamsOrder) IsKnown() bool {
	switch r {
	case SettingAllowPolicyListParamsOrderPattern, SettingAllowPolicyListParamsOrderCreatedAt:
		return true
	}
	return false
}

// Type of pattern matching.
//
//   - EMAIL: matches a full email address (e.g. `user@example.com`)
//   - DOMAIN: matches a domain name (e.g. `example.com`)
//   - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
//     `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
//     `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
//     link-local, unspecified, and IPv4 broadcast addresses, including their
//     IPv4-mapped IPv6 equivalents.
//   - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
//     but it may appear on existing entries.
type SettingAllowPolicyListParamsPatternType string

const (
	SettingAllowPolicyListParamsPatternTypeEmail   SettingAllowPolicyListParamsPatternType = "EMAIL"
	SettingAllowPolicyListParamsPatternTypeDomain  SettingAllowPolicyListParamsPatternType = "DOMAIN"
	SettingAllowPolicyListParamsPatternTypeIP      SettingAllowPolicyListParamsPatternType = "IP"
	SettingAllowPolicyListParamsPatternTypeUnknown SettingAllowPolicyListParamsPatternType = "UNKNOWN"
)

func (r SettingAllowPolicyListParamsPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyListParamsPatternTypeEmail, SettingAllowPolicyListParamsPatternTypeDomain, SettingAllowPolicyListParamsPatternTypeIP, SettingAllowPolicyListParamsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPolicyDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingAllowPolicyDeleteResponseEnvelope struct {
	Errors   []SettingAllowPolicyDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingAllowPolicyDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingAllowPolicyDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingAllowPolicyDeleteResponse                `json:"result"`
	JSON    settingAllowPolicyDeleteResponseEnvelopeJSON    `json:"-"`
}

// settingAllowPolicyDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAllowPolicyDeleteResponseEnvelope]
type settingAllowPolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyDeleteResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           SettingAllowPolicyDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingAllowPolicyDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingAllowPolicyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingAllowPolicyDeleteResponseEnvelopeErrors]
type settingAllowPolicyDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAllowPolicyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    settingAllowPolicyDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingAllowPolicyDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingAllowPolicyDeleteResponseEnvelopeErrorsSource]
type settingAllowPolicyDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyDeleteResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           SettingAllowPolicyDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingAllowPolicyDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingAllowPolicyDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAllowPolicyDeleteResponseEnvelopeMessages]
type settingAllowPolicyDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAllowPolicyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    settingAllowPolicyDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingAllowPolicyDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingAllowPolicyDeleteResponseEnvelopeMessagesSource]
type settingAllowPolicyDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingAllowPolicyDeleteResponseEnvelopeSuccess bool

const (
	SettingAllowPolicyDeleteResponseEnvelopeSuccessTrue SettingAllowPolicyDeleteResponseEnvelopeSuccess = true
)

func (r SettingAllowPolicyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingAllowPolicyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingAllowPolicyBatchParams struct {
	// Identifier.
	AccountID param.Field[string]                                `path:"account_id" api:"required"`
	Deletes   param.Field[[]SettingAllowPolicyBatchParamsDelete] `json:"deletes" api:"required"`
	Patches   param.Field[[]SettingAllowPolicyBatchParamsPatch]  `json:"patches" api:"required"`
	Posts     param.Field[[]SettingAllowPolicyBatchParamsPost]   `json:"posts" api:"required"`
	Puts      param.Field[[]SettingAllowPolicyBatchParamsPut]    `json:"puts" api:"required"`
}

func (r SettingAllowPolicyBatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingAllowPolicyBatchParamsDelete struct {
	// Allow policy identifier.
	ID param.Field[string] `json:"id" api:"required" format:"uuid"`
}

func (r SettingAllowPolicyBatchParamsDelete) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An email allow policy.
type SettingAllowPolicyBatchParamsPatch struct {
	Comments param.Field[string] `json:"comments"`
	// Exempts messages from this sender from Spam, Spoof and Bulk dispositions only;
	// Malicious and Suspicious dispositions still apply.
	IsAcceptableSender param.Field[bool] `json:"is_acceptable_sender"`
	// Bypasses all detections for messages to this recipient.
	IsExemptRecipient param.Field[bool] `json:"is_exempt_recipient"`
	// Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_exempt_recipient` instead.
	IsRecipient param.Field[bool] `json:"is_recipient"`
	IsRegex     param.Field[bool] `json:"is_regex"`
	// Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_trusted_sender` instead.
	IsSender param.Field[bool] `json:"is_sender"`
	// Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_acceptable_sender` instead.
	IsSpoof param.Field[bool] `json:"is_spoof"`
	// Bypasses all detections and link following for messages from this sender.
	IsTrustedSender param.Field[bool] `json:"is_trusted_sender"`
	// The pattern value to match. The format depends on `pattern_type`: a valid email
	// address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN
	// (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g.
	// `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the
	// API rejects private or unique-local, loopback, link-local, unspecified, and IPv4
	// broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	Pattern param.Field[string] `json:"pattern"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
	//   `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
	//   `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
	//   link-local, unspecified, and IPv4 broadcast addresses, including their
	//   IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
	//   but it may appear on existing entries.
	PatternType param.Field[SettingAllowPolicyBatchParamsPatchesPatternType] `json:"pattern_type"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender param.Field[bool] `json:"verify_sender"`
}

func (r SettingAllowPolicyBatchParamsPatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of pattern matching.
//
//   - EMAIL: matches a full email address (e.g. `user@example.com`)
//   - DOMAIN: matches a domain name (e.g. `example.com`)
//   - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
//     `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
//     `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
//     link-local, unspecified, and IPv4 broadcast addresses, including their
//     IPv4-mapped IPv6 equivalents.
//   - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
//     but it may appear on existing entries.
type SettingAllowPolicyBatchParamsPatchesPatternType string

const (
	SettingAllowPolicyBatchParamsPatchesPatternTypeEmail   SettingAllowPolicyBatchParamsPatchesPatternType = "EMAIL"
	SettingAllowPolicyBatchParamsPatchesPatternTypeDomain  SettingAllowPolicyBatchParamsPatchesPatternType = "DOMAIN"
	SettingAllowPolicyBatchParamsPatchesPatternTypeIP      SettingAllowPolicyBatchParamsPatchesPatternType = "IP"
	SettingAllowPolicyBatchParamsPatchesPatternTypeUnknown SettingAllowPolicyBatchParamsPatchesPatternType = "UNKNOWN"
)

func (r SettingAllowPolicyBatchParamsPatchesPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyBatchParamsPatchesPatternTypeEmail, SettingAllowPolicyBatchParamsPatchesPatternTypeDomain, SettingAllowPolicyBatchParamsPatchesPatternTypeIP, SettingAllowPolicyBatchParamsPatchesPatternTypeUnknown:
		return true
	}
	return false
}

// Create an allow policy.
type SettingAllowPolicyBatchParamsPost struct {
	// Exempts messages from this sender from Spam, Spoof and Bulk dispositions only;
	// Malicious and Suspicious dispositions still apply.
	IsAcceptableSender param.Field[bool] `json:"is_acceptable_sender" api:"required"`
	// Bypasses all detections for messages to this recipient.
	IsExemptRecipient param.Field[bool] `json:"is_exempt_recipient" api:"required"`
	IsRegex           param.Field[bool] `json:"is_regex" api:"required"`
	// Bypasses all detections and link following for messages from this sender.
	IsTrustedSender param.Field[bool] `json:"is_trusted_sender" api:"required"`
	// The pattern value to match. The format depends on `pattern_type`: a valid email
	// address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN
	// (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g.
	// `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the
	// API rejects private or unique-local, loopback, link-local, unspecified, and IPv4
	// broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	Pattern param.Field[string] `json:"pattern" api:"required"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
	//   `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
	//   `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
	//   link-local, unspecified, and IPv4 broadcast addresses, including their
	//   IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
	//   but it may appear on existing entries.
	PatternType param.Field[SettingAllowPolicyBatchParamsPostsPatternType] `json:"pattern_type" api:"required"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender param.Field[bool]   `json:"verify_sender" api:"required"`
	Comments     param.Field[string] `json:"comments"`
	// Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_exempt_recipient` instead.
	IsRecipient param.Field[bool] `json:"is_recipient"`
	// Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_trusted_sender` instead.
	IsSender param.Field[bool] `json:"is_sender"`
	// Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_acceptable_sender` instead.
	IsSpoof param.Field[bool] `json:"is_spoof"`
}

func (r SettingAllowPolicyBatchParamsPost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of pattern matching.
//
//   - EMAIL: matches a full email address (e.g. `user@example.com`)
//   - DOMAIN: matches a domain name (e.g. `example.com`)
//   - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
//     `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
//     `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
//     link-local, unspecified, and IPv4 broadcast addresses, including their
//     IPv4-mapped IPv6 equivalents.
//   - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
//     but it may appear on existing entries.
type SettingAllowPolicyBatchParamsPostsPatternType string

const (
	SettingAllowPolicyBatchParamsPostsPatternTypeEmail   SettingAllowPolicyBatchParamsPostsPatternType = "EMAIL"
	SettingAllowPolicyBatchParamsPostsPatternTypeDomain  SettingAllowPolicyBatchParamsPostsPatternType = "DOMAIN"
	SettingAllowPolicyBatchParamsPostsPatternTypeIP      SettingAllowPolicyBatchParamsPostsPatternType = "IP"
	SettingAllowPolicyBatchParamsPostsPatternTypeUnknown SettingAllowPolicyBatchParamsPostsPatternType = "UNKNOWN"
)

func (r SettingAllowPolicyBatchParamsPostsPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyBatchParamsPostsPatternTypeEmail, SettingAllowPolicyBatchParamsPostsPatternTypeDomain, SettingAllowPolicyBatchParamsPostsPatternTypeIP, SettingAllowPolicyBatchParamsPostsPatternTypeUnknown:
		return true
	}
	return false
}

// An email allow policy.
type SettingAllowPolicyBatchParamsPut struct {
	Comments param.Field[string] `json:"comments"`
	// Exempts messages from this sender from Spam, Spoof and Bulk dispositions only;
	// Malicious and Suspicious dispositions still apply.
	IsAcceptableSender param.Field[bool] `json:"is_acceptable_sender"`
	// Bypasses all detections for messages to this recipient.
	IsExemptRecipient param.Field[bool] `json:"is_exempt_recipient"`
	// Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_exempt_recipient` instead.
	IsRecipient param.Field[bool] `json:"is_recipient"`
	IsRegex     param.Field[bool] `json:"is_regex"`
	// Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_trusted_sender` instead.
	IsSender param.Field[bool] `json:"is_sender"`
	// Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life:
	// July 1, 2026.
	//
	// Deprecated: Use `is_acceptable_sender` instead.
	IsSpoof param.Field[bool] `json:"is_spoof"`
	// Bypasses all detections and link following for messages from this sender.
	IsTrustedSender param.Field[bool] `json:"is_trusted_sender"`
	// The pattern value to match. The format depends on `pattern_type`: a valid email
	// address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN
	// (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g.
	// `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the
	// API rejects private or unique-local, loopback, link-local, unspecified, and IPv4
	// broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	Pattern param.Field[string] `json:"pattern"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
	//   `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
	//   `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
	//   link-local, unspecified, and IPv4 broadcast addresses, including their
	//   IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
	//   but it may appear on existing entries.
	PatternType param.Field[SettingAllowPolicyBatchParamsPutsPatternType] `json:"pattern_type"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender param.Field[bool] `json:"verify_sender"`
}

func (r SettingAllowPolicyBatchParamsPut) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of pattern matching.
//
//   - EMAIL: matches a full email address (e.g. `user@example.com`)
//   - DOMAIN: matches a domain name (e.g. `example.com`)
//   - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
//     `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
//     `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
//     link-local, unspecified, and IPv4 broadcast addresses, including their
//     IPv4-mapped IPv6 equivalents.
//   - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
//     but it may appear on existing entries.
type SettingAllowPolicyBatchParamsPutsPatternType string

const (
	SettingAllowPolicyBatchParamsPutsPatternTypeEmail   SettingAllowPolicyBatchParamsPutsPatternType = "EMAIL"
	SettingAllowPolicyBatchParamsPutsPatternTypeDomain  SettingAllowPolicyBatchParamsPutsPatternType = "DOMAIN"
	SettingAllowPolicyBatchParamsPutsPatternTypeIP      SettingAllowPolicyBatchParamsPutsPatternType = "IP"
	SettingAllowPolicyBatchParamsPutsPatternTypeUnknown SettingAllowPolicyBatchParamsPutsPatternType = "UNKNOWN"
)

func (r SettingAllowPolicyBatchParamsPutsPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyBatchParamsPutsPatternTypeEmail, SettingAllowPolicyBatchParamsPutsPatternTypeDomain, SettingAllowPolicyBatchParamsPutsPatternTypeIP, SettingAllowPolicyBatchParamsPutsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPolicyBatchResponseEnvelope struct {
	Errors   []SettingAllowPolicyBatchResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingAllowPolicyBatchResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingAllowPolicyBatchResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingAllowPolicyBatchResponse                `json:"result"`
	JSON    settingAllowPolicyBatchResponseEnvelopeJSON    `json:"-"`
}

// settingAllowPolicyBatchResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAllowPolicyBatchResponseEnvelope]
type settingAllowPolicyBatchResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyBatchResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyBatchResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyBatchResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           SettingAllowPolicyBatchResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingAllowPolicyBatchResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingAllowPolicyBatchResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingAllowPolicyBatchResponseEnvelopeErrors]
type settingAllowPolicyBatchResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAllowPolicyBatchResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyBatchResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyBatchResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    settingAllowPolicyBatchResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingAllowPolicyBatchResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingAllowPolicyBatchResponseEnvelopeErrorsSource]
type settingAllowPolicyBatchResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyBatchResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyBatchResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyBatchResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           SettingAllowPolicyBatchResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingAllowPolicyBatchResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingAllowPolicyBatchResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAllowPolicyBatchResponseEnvelopeMessages]
type settingAllowPolicyBatchResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAllowPolicyBatchResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyBatchResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyBatchResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    settingAllowPolicyBatchResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingAllowPolicyBatchResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingAllowPolicyBatchResponseEnvelopeMessagesSource]
type settingAllowPolicyBatchResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyBatchResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyBatchResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingAllowPolicyBatchResponseEnvelopeSuccess bool

const (
	SettingAllowPolicyBatchResponseEnvelopeSuccessTrue SettingAllowPolicyBatchResponseEnvelopeSuccess = true
)

func (r SettingAllowPolicyBatchResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingAllowPolicyBatchResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingAllowPolicyEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Comments  param.Field[string] `json:"comments"`
	// Exempts messages from this sender from Spam, Spoof and Bulk dispositions only;
	// Malicious and Suspicious dispositions still apply.
	IsAcceptableSender param.Field[bool] `json:"is_acceptable_sender"`
	// Bypasses all detections for messages to this recipient.
	IsExemptRecipient param.Field[bool] `json:"is_exempt_recipient"`
	// Deprecated as of July 1, 2025. Use `is_exempt_recipient` instead. End of life:
	// July 1, 2026.
	IsRecipient param.Field[bool] `json:"is_recipient"`
	IsRegex     param.Field[bool] `json:"is_regex"`
	// Deprecated as of July 1, 2025. Use `is_trusted_sender` instead. End of life:
	// July 1, 2026.
	IsSender param.Field[bool] `json:"is_sender"`
	// Deprecated as of July 1, 2025. Use `is_acceptable_sender` instead. End of life:
	// July 1, 2026.
	IsSpoof param.Field[bool] `json:"is_spoof"`
	// Bypasses all detections and link following for messages from this sender.
	IsTrustedSender param.Field[bool] `json:"is_trusted_sender"`
	// The pattern value to match. The format depends on `pattern_type`: a valid email
	// address for EMAIL (e.g. `user@example.com`), a valid domain name for DOMAIN
	// (e.g. `example.com`), or a plain IPv4 or IPv6 address or CIDR block for IP (e.g.
	// `1.2.3.4`, `1.2.3.0/24`, `2606:4700:4700::1111`, or `2606:4700:4700::/48`); the
	// API rejects private or unique-local, loopback, link-local, unspecified, and IPv4
	// broadcast addresses, including their IPv4-mapped IPv6 equivalents.
	Pattern param.Field[string] `json:"pattern"`
	// Type of pattern matching.
	//
	// - EMAIL: matches a full email address (e.g. `user@example.com`)
	// - DOMAIN: matches a domain name (e.g. `example.com`)
	// - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
	//   `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
	//   `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
	//   link-local, unspecified, and IPv4 broadcast addresses, including their
	//   IPv4-mapped IPv6 equivalents.
	// - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
	//   but it may appear on existing entries.
	PatternType param.Field[SettingAllowPolicyEditParamsPatternType] `json:"pattern_type"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender param.Field[bool] `json:"verify_sender"`
}

func (r SettingAllowPolicyEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of pattern matching.
//
//   - EMAIL: matches a full email address (e.g. `user@example.com`)
//   - DOMAIN: matches a domain name (e.g. `example.com`)
//   - IP: matches a plain IPv4 or IPv6 address (e.g. `1.2.3.4` or
//     `2606:4700:4700::1111`) or CIDR block (e.g. `1.2.3.0/24` or
//     `2606:4700:4700::/48`). The API rejects private or unique-local, loopback,
//     link-local, unspecified, and IPv4 broadcast addresses, including their
//     IPv4-mapped IPv6 equivalents.
//   - UNKNOWN: deprecated; you cannot use this when creating or updating policies,
//     but it may appear on existing entries.
type SettingAllowPolicyEditParamsPatternType string

const (
	SettingAllowPolicyEditParamsPatternTypeEmail   SettingAllowPolicyEditParamsPatternType = "EMAIL"
	SettingAllowPolicyEditParamsPatternTypeDomain  SettingAllowPolicyEditParamsPatternType = "DOMAIN"
	SettingAllowPolicyEditParamsPatternTypeIP      SettingAllowPolicyEditParamsPatternType = "IP"
	SettingAllowPolicyEditParamsPatternTypeUnknown SettingAllowPolicyEditParamsPatternType = "UNKNOWN"
)

func (r SettingAllowPolicyEditParamsPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyEditParamsPatternTypeEmail, SettingAllowPolicyEditParamsPatternTypeDomain, SettingAllowPolicyEditParamsPatternTypeIP, SettingAllowPolicyEditParamsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPolicyEditResponseEnvelope struct {
	Errors   []SettingAllowPolicyEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingAllowPolicyEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingAllowPolicyEditResponseEnvelopeSuccess `json:"success" api:"required"`
	// An email allow policy.
	Result SettingAllowPolicyEditResponse             `json:"result"`
	JSON   settingAllowPolicyEditResponseEnvelopeJSON `json:"-"`
}

// settingAllowPolicyEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAllowPolicyEditResponseEnvelope]
type settingAllowPolicyEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyEditResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code" api:"required"`
	Message          string                                             `json:"message" api:"required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           SettingAllowPolicyEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingAllowPolicyEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingAllowPolicyEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingAllowPolicyEditResponseEnvelopeErrors]
type settingAllowPolicyEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAllowPolicyEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    settingAllowPolicyEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingAllowPolicyEditResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingAllowPolicyEditResponseEnvelopeErrorsSource]
type settingAllowPolicyEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyEditResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           SettingAllowPolicyEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingAllowPolicyEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingAllowPolicyEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAllowPolicyEditResponseEnvelopeMessages]
type settingAllowPolicyEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAllowPolicyEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    settingAllowPolicyEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingAllowPolicyEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingAllowPolicyEditResponseEnvelopeMessagesSource]
type settingAllowPolicyEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingAllowPolicyEditResponseEnvelopeSuccess bool

const (
	SettingAllowPolicyEditResponseEnvelopeSuccessTrue SettingAllowPolicyEditResponseEnvelopeSuccess = true
)

func (r SettingAllowPolicyEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingAllowPolicyEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingAllowPolicyGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingAllowPolicyGetResponseEnvelope struct {
	Errors   []SettingAllowPolicyGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingAllowPolicyGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingAllowPolicyGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// An email allow policy.
	Result SettingAllowPolicyGetResponse             `json:"result"`
	JSON   settingAllowPolicyGetResponseEnvelopeJSON `json:"-"`
}

// settingAllowPolicyGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAllowPolicyGetResponseEnvelope]
type settingAllowPolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyGetResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           SettingAllowPolicyGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingAllowPolicyGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingAllowPolicyGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingAllowPolicyGetResponseEnvelopeErrors]
type settingAllowPolicyGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAllowPolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    settingAllowPolicyGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingAllowPolicyGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SettingAllowPolicyGetResponseEnvelopeErrorsSource]
type settingAllowPolicyGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyGetResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           SettingAllowPolicyGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingAllowPolicyGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingAllowPolicyGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingAllowPolicyGetResponseEnvelopeMessages]
type settingAllowPolicyGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAllowPolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    settingAllowPolicyGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingAllowPolicyGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingAllowPolicyGetResponseEnvelopeMessagesSource]
type settingAllowPolicyGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingAllowPolicyGetResponseEnvelopeSuccess bool

const (
	SettingAllowPolicyGetResponseEnvelopeSuccessTrue SettingAllowPolicyGetResponseEnvelopeSuccess = true
)

func (r SettingAllowPolicyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingAllowPolicyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
