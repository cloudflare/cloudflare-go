// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_tagging

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/tidwall/gjson"
)

// AccountTagService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountTagService] method instead.
type AccountTagService struct {
	Options []option.RequestOption
}

// NewAccountTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountTagService(opts ...option.RequestOption) (r *AccountTagService) {
	r = &AccountTagService{}
	r.Options = opts
	return
}

// Creates or updates tags for a specific account-level resource.
func (r *AccountTagService) Update(ctx context.Context, params AccountTagUpdateParams, opts ...option.RequestOption) (res *AccountTagUpdateResponse, err error) {
	var env AccountTagUpdateResponseEnvelope
	if params.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/tags", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Removes all tags from a specific account-level resource.
func (r *AccountTagService) Delete(ctx context.Context, params AccountTagDeleteParams, opts ...option.RequestOption) (err error) {
	if params.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/tags", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieves tags for a specific account-level resource.
func (r *AccountTagService) Get(ctx context.Context, params AccountTagGetParams, opts ...option.RequestOption) (res *AccountTagGetResponse, err error) {
	var env AccountTagGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/tags", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Response for access_application resources
type AccountTagUpdateResponse struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// This field can have the runtime type of [map[string]string].
	Tags interface{}                  `json:"tags" api:"required"`
	Type AccountTagUpdateResponseType `json:"type" api:"required"`
	// Access application ID is required only for access_application_policy resources
	AccessApplicationID string `json:"access_application_id" format:"uuid"`
	// Worker ID is required only for worker_version resources
	WorkerID string `json:"worker_id"`
	// Zone ID is required only for zone-level resources
	ZoneID string                       `json:"zone_id"`
	JSON   accountTagUpdateResponseJSON `json:"-"`
	union  AccountTagUpdateResponseUnion
}

// accountTagUpdateResponseJSON contains the JSON metadata for the struct
// [AccountTagUpdateResponse]
type accountTagUpdateResponseJSON struct {
	ID                  apijson.Field
	Etag                apijson.Field
	Name                apijson.Field
	Tags                apijson.Field
	Type                apijson.Field
	AccessApplicationID apijson.Field
	WorkerID            apijson.Field
	ZoneID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r accountTagUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccountTagUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccountTagUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountTagUpdateResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccount],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectImage],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectQueue],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorker],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectZone].
func (r AccountTagUpdateResponse) AsUnion() AccountTagUpdateResponseUnion {
	return r.union
}

// Response for access_application resources
//
// Union satisfied by
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccount],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectImage],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectQueue],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorker],
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion] or
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectZone].
type AccountTagUpdateResponseUnion interface {
	implementsAccountTagUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountTagUpdateResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication{}),
			DiscriminatorValue: "access_application",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy{}),
			DiscriminatorValue: "access_application_policy",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup{}),
			DiscriminatorValue: "access_group",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccount{}),
			DiscriminatorValue: "account",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway{}),
			DiscriminatorValue: "ai_gateway",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy{}),
			DiscriminatorValue: "alerting_policy",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook{}),
			DiscriminatorValue: "alerting_webhook",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation{}),
			DiscriminatorValue: "api_gateway_operation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel{}),
			DiscriminatorValue: "cloudflared_tunnel",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate{}),
			DiscriminatorValue: "custom_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname{}),
			DiscriminatorValue: "custom_hostname",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database{}),
			DiscriminatorValue: "d1_database",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord{}),
			DiscriminatorValue: "dns_record",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList{}),
			DiscriminatorValue: "gateway_list",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule{}),
			DiscriminatorValue: "gateway_rule",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectImage{}),
			DiscriminatorValue: "image",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate{}),
			DiscriminatorValue: "managed_client_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare{}),
			DiscriminatorValue: "resource_share",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput{}),
			DiscriminatorValue: "stream_live_input",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo{}),
			DiscriminatorValue: "stream_video",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorker{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion{}),
			DiscriminatorValue: "worker_version",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagUpdateResponseResourceTaggingTaggedResourceObjectZone{}),
			DiscriminatorValue: "zone",
		},
	)
}

// Response for access_application resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                                `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationTypeAccessApplication AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationType = "access_application"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationTypeAccessApplication:
		return true
	}
	return false
}

// Response for access_application_policy resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// Access application ID is required only for access_application_policy resources
	AccessApplicationID string `json:"access_application_id" api:"required" format:"uuid"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                                      `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                                 `json:"zone_id" api:"required"`
	JSON   accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON struct {
	ID                  apijson.Field
	AccessApplicationID apijson.Field
	Etag                apijson.Field
	Name                apijson.Field
	Tags                apijson.Field
	Type                apijson.Field
	ZoneID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyTypeAccessApplicationPolicy AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType = "access_application_policy"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyTypeAccessApplicationPolicy:
		return true
	}
	return false
}

// Response for access_group resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                          `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupTypeAccessGroup AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupType = "access_group"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupTypeAccessGroup:
		return true
	}
	return false
}

// Response for account resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccount struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                      `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccountType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccountJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccountJSON contains
// the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccount]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccountJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectAccountJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccount) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccountType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccountTypeAccount AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccountType = "account"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccountType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAccountTypeAccount:
		return true
	}
	return false
}

// Response for ai_gateway resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                        `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayTypeAIGateway AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayType = "ai_gateway"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayTypeAIGateway:
		return true
	}
	return false
}

// Response for alerting_policy resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyTypeAlertingPolicy AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyType = "alerting_policy"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyTypeAlertingPolicy:
		return true
	}
	return false
}

// Response for alerting_webhook resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                              `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookTypeAlertingWebhook AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookType = "alerting_webhook"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookTypeAlertingWebhook:
		return true
	}
	return false
}

// Response for api_gateway_operation resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                                  `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                             `json:"zone_id" api:"required"`
	JSON   accountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationTypeAPIGatewayOperation AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType = "api_gateway_operation"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationTypeAPIGatewayOperation:
		return true
	}
	return false
}

// Response for cloudflared_tunnel resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                                `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelTypeCloudflaredTunnel AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType = "cloudflared_tunnel"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelTypeCloudflaredTunnel:
		return true
	}
	return false
}

// Response for custom_certificate resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                                `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                           `json:"zone_id" api:"required"`
	JSON   accountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateTypeCustomCertificate AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateType = "custom_certificate"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateTypeCustomCertificate:
		return true
	}
	return false
}

// Response for custom_hostname resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                        `json:"zone_id" api:"required"`
	JSON   accountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameTypeCustomHostname AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameType = "custom_hostname"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameTypeCustomHostname:
		return true
	}
	return false
}

// Response for d1_database resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                         `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseTypeD1Database AccountTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseType = "d1_database"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseTypeD1Database:
		return true
	}
	return false
}

// Response for dns_record resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                        `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                   `json:"zone_id" api:"required"`
	JSON   accountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordTypeDNSRecord AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordType = "dns_record"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordTypeDNSRecord:
		return true
	}
	return false
}

// Response for durable_object_namespace resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                                     `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceTypeDurableObjectNamespace AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType = "durable_object_namespace"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

// Response for gateway_list resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                          `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListTypeGatewayList AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListType = "gateway_list"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListTypeGatewayList:
		return true
	}
	return false
}

// Response for gateway_rule resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                          `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleTypeGatewayRule AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleType = "gateway_rule"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleTypeGatewayRule:
		return true
	}
	return false
}

// Response for image resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectImage struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                    `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectImageType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectImageJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectImageJSON contains
// the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectImage]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectImageJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectImageJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectImage) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectImageType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectImageTypeImage AccountTagUpdateResponseResourceTaggingTaggedResourceObjectImageType = "image"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectImageType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectImageTypeImage:
		return true
	}
	return false
}

// Response for kv_namespace resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                          `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceTypeKVNamespace AccountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceType = "kv_namespace"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

// Response for managed_client_certificate resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                                       `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                                  `json:"zone_id" api:"required"`
	JSON   accountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateTypeManagedClientCertificate AccountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType = "managed_client_certificate"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateTypeManagedClientCertificate:
		return true
	}
	return false
}

// Response for queue resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectQueue struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                    `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectQueueType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectQueueJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectQueueJSON contains
// the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectQueue]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectQueueJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectQueueJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectQueue) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectQueueType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectQueueTypeQueue AccountTagUpdateResponseResourceTaggingTaggedResourceObjectQueueType = "queue"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectQueueType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectQueueTypeQueue:
		return true
	}
	return false
}

// Response for r2_bucket resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                       `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketJSON contains
// the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketTypeR2Bucket AccountTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketType = "r2_bucket"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketTypeR2Bucket:
		return true
	}
	return false
}

// Response for resource_share resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                            `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareTypeResourceShare AccountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareType = "resource_share"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareTypeResourceShare:
		return true
	}
	return false
}

// Response for stream_live_input resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                              `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputTypeStreamLiveInput AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputType = "stream_live_input"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputTypeStreamLiveInput:
		return true
	}
	return false
}

// Response for stream_video resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                          `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoTypeStreamVideo AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoType = "stream_video"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoTypeStreamVideo:
		return true
	}
	return false
}

// Response for worker resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorker struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                     `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerType `json:"type" api:"required"`
	JSON accountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerJSON contains
// the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorker]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorker) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerTypeWorker AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerType = "worker"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerTypeWorker:
		return true
	}
	return false
}

// Response for worker_version resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                            `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionType `json:"type" api:"required"`
	// Worker ID is required only for worker_version resources
	WorkerID string                                                                       `json:"worker_id" api:"required"`
	JSON     accountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON
// contains the JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	WorkerID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionTypeWorkerVersion AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionType = "worker_version"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionTypeWorkerVersion:
		return true
	}
	return false
}

// Response for zone resources
type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectZone struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                   `json:"tags" api:"required"`
	Type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectZoneType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                              `json:"zone_id" api:"required"`
	JSON   accountTagUpdateResponseResourceTaggingTaggedResourceObjectZoneJSON `json:"-"`
}

// accountTagUpdateResponseResourceTaggingTaggedResourceObjectZoneJSON contains the
// JSON metadata for the struct
// [AccountTagUpdateResponseResourceTaggingTaggedResourceObjectZone]
type accountTagUpdateResponseResourceTaggingTaggedResourceObjectZoneJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseResourceTaggingTaggedResourceObjectZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseResourceTaggingTaggedResourceObjectZoneJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectZone) implementsAccountTagUpdateResponse() {
}

type AccountTagUpdateResponseResourceTaggingTaggedResourceObjectZoneType string

const (
	AccountTagUpdateResponseResourceTaggingTaggedResourceObjectZoneTypeZone AccountTagUpdateResponseResourceTaggingTaggedResourceObjectZoneType = "zone"
)

func (r AccountTagUpdateResponseResourceTaggingTaggedResourceObjectZoneType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseResourceTaggingTaggedResourceObjectZoneTypeZone:
		return true
	}
	return false
}

type AccountTagUpdateResponseType string

const (
	AccountTagUpdateResponseTypeAccessApplication        AccountTagUpdateResponseType = "access_application"
	AccountTagUpdateResponseTypeAccessApplicationPolicy  AccountTagUpdateResponseType = "access_application_policy"
	AccountTagUpdateResponseTypeAccessGroup              AccountTagUpdateResponseType = "access_group"
	AccountTagUpdateResponseTypeAccount                  AccountTagUpdateResponseType = "account"
	AccountTagUpdateResponseTypeAIGateway                AccountTagUpdateResponseType = "ai_gateway"
	AccountTagUpdateResponseTypeAlertingPolicy           AccountTagUpdateResponseType = "alerting_policy"
	AccountTagUpdateResponseTypeAlertingWebhook          AccountTagUpdateResponseType = "alerting_webhook"
	AccountTagUpdateResponseTypeAPIGatewayOperation      AccountTagUpdateResponseType = "api_gateway_operation"
	AccountTagUpdateResponseTypeCloudflaredTunnel        AccountTagUpdateResponseType = "cloudflared_tunnel"
	AccountTagUpdateResponseTypeCustomCertificate        AccountTagUpdateResponseType = "custom_certificate"
	AccountTagUpdateResponseTypeCustomHostname           AccountTagUpdateResponseType = "custom_hostname"
	AccountTagUpdateResponseTypeD1Database               AccountTagUpdateResponseType = "d1_database"
	AccountTagUpdateResponseTypeDNSRecord                AccountTagUpdateResponseType = "dns_record"
	AccountTagUpdateResponseTypeDurableObjectNamespace   AccountTagUpdateResponseType = "durable_object_namespace"
	AccountTagUpdateResponseTypeGatewayList              AccountTagUpdateResponseType = "gateway_list"
	AccountTagUpdateResponseTypeGatewayRule              AccountTagUpdateResponseType = "gateway_rule"
	AccountTagUpdateResponseTypeImage                    AccountTagUpdateResponseType = "image"
	AccountTagUpdateResponseTypeKVNamespace              AccountTagUpdateResponseType = "kv_namespace"
	AccountTagUpdateResponseTypeManagedClientCertificate AccountTagUpdateResponseType = "managed_client_certificate"
	AccountTagUpdateResponseTypeQueue                    AccountTagUpdateResponseType = "queue"
	AccountTagUpdateResponseTypeR2Bucket                 AccountTagUpdateResponseType = "r2_bucket"
	AccountTagUpdateResponseTypeResourceShare            AccountTagUpdateResponseType = "resource_share"
	AccountTagUpdateResponseTypeStreamLiveInput          AccountTagUpdateResponseType = "stream_live_input"
	AccountTagUpdateResponseTypeStreamVideo              AccountTagUpdateResponseType = "stream_video"
	AccountTagUpdateResponseTypeWorker                   AccountTagUpdateResponseType = "worker"
	AccountTagUpdateResponseTypeWorkerVersion            AccountTagUpdateResponseType = "worker_version"
	AccountTagUpdateResponseTypeZone                     AccountTagUpdateResponseType = "zone"
)

func (r AccountTagUpdateResponseType) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseTypeAccessApplication, AccountTagUpdateResponseTypeAccessApplicationPolicy, AccountTagUpdateResponseTypeAccessGroup, AccountTagUpdateResponseTypeAccount, AccountTagUpdateResponseTypeAIGateway, AccountTagUpdateResponseTypeAlertingPolicy, AccountTagUpdateResponseTypeAlertingWebhook, AccountTagUpdateResponseTypeAPIGatewayOperation, AccountTagUpdateResponseTypeCloudflaredTunnel, AccountTagUpdateResponseTypeCustomCertificate, AccountTagUpdateResponseTypeCustomHostname, AccountTagUpdateResponseTypeD1Database, AccountTagUpdateResponseTypeDNSRecord, AccountTagUpdateResponseTypeDurableObjectNamespace, AccountTagUpdateResponseTypeGatewayList, AccountTagUpdateResponseTypeGatewayRule, AccountTagUpdateResponseTypeImage, AccountTagUpdateResponseTypeKVNamespace, AccountTagUpdateResponseTypeManagedClientCertificate, AccountTagUpdateResponseTypeQueue, AccountTagUpdateResponseTypeR2Bucket, AccountTagUpdateResponseTypeResourceShare, AccountTagUpdateResponseTypeStreamLiveInput, AccountTagUpdateResponseTypeStreamVideo, AccountTagUpdateResponseTypeWorker, AccountTagUpdateResponseTypeWorkerVersion, AccountTagUpdateResponseTypeZone:
		return true
	}
	return false
}

// Response for access_application resources
type AccountTagGetResponse struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// This field can have the runtime type of [map[string]string].
	Tags interface{}               `json:"tags" api:"required"`
	Type AccountTagGetResponseType `json:"type" api:"required"`
	// Access application ID is required only for access_application_policy resources
	AccessApplicationID string `json:"access_application_id" format:"uuid"`
	// Worker ID is required only for worker_version resources
	WorkerID string `json:"worker_id"`
	// Zone ID is required only for zone-level resources
	ZoneID string                    `json:"zone_id"`
	JSON   accountTagGetResponseJSON `json:"-"`
	union  AccountTagGetResponseUnion
}

// accountTagGetResponseJSON contains the JSON metadata for the struct
// [AccountTagGetResponse]
type accountTagGetResponseJSON struct {
	ID                  apijson.Field
	Etag                apijson.Field
	Name                apijson.Field
	Tags                apijson.Field
	Type                apijson.Field
	AccessApplicationID apijson.Field
	WorkerID            apijson.Field
	ZoneID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r accountTagGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccountTagGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccountTagGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountTagGetResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAccount],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAIGateway],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectD1Database],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayList],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectImage],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectQueue],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectResourceShare],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectWorker],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectZone].
func (r AccountTagGetResponse) AsUnion() AccountTagGetResponseUnion {
	return r.union
}

// Response for access_application resources
//
// Union satisfied by
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAccount],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAIGateway],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectD1Database],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayList],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectImage],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectQueue],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectResourceShare],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectWorker],
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion] or
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectZone].
type AccountTagGetResponseUnion interface {
	implementsAccountTagGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountTagGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication{}),
			DiscriminatorValue: "access_application",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy{}),
			DiscriminatorValue: "access_application_policy",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup{}),
			DiscriminatorValue: "access_group",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectAccount{}),
			DiscriminatorValue: "account",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectAIGateway{}),
			DiscriminatorValue: "ai_gateway",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy{}),
			DiscriminatorValue: "alerting_policy",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook{}),
			DiscriminatorValue: "alerting_webhook",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation{}),
			DiscriminatorValue: "api_gateway_operation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel{}),
			DiscriminatorValue: "cloudflared_tunnel",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate{}),
			DiscriminatorValue: "custom_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname{}),
			DiscriminatorValue: "custom_hostname",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectD1Database{}),
			DiscriminatorValue: "d1_database",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord{}),
			DiscriminatorValue: "dns_record",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayList{}),
			DiscriminatorValue: "gateway_list",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule{}),
			DiscriminatorValue: "gateway_rule",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectImage{}),
			DiscriminatorValue: "image",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate{}),
			DiscriminatorValue: "managed_client_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectResourceShare{}),
			DiscriminatorValue: "resource_share",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput{}),
			DiscriminatorValue: "stream_live_input",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo{}),
			DiscriminatorValue: "stream_video",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectWorker{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion{}),
			DiscriminatorValue: "worker_version",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountTagGetResponseResourceTaggingTaggedResourceObjectZone{}),
			DiscriminatorValue: "zone",
		},
	)
}

// Response for access_application resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON
// contains the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication]
type accountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationTypeAccessApplication AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationType = "access_application"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationTypeAccessApplication:
		return true
	}
	return false
}

// Response for access_application_policy resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// Access application ID is required only for access_application_policy resources
	AccessApplicationID string `json:"access_application_id" api:"required" format:"uuid"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                                   `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                              `json:"zone_id" api:"required"`
	JSON   accountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON
// contains the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy]
type accountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON struct {
	ID                  apijson.Field
	AccessApplicationID apijson.Field
	Etag                apijson.Field
	Name                apijson.Field
	Tags                apijson.Field
	Type                apijson.Field
	ZoneID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyTypeAccessApplicationPolicy AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType = "access_application_policy"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyTypeAccessApplicationPolicy:
		return true
	}
	return false
}

// Response for access_group resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                       `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupJSON contains
// the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup]
type accountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupTypeAccessGroup AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupType = "access_group"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupTypeAccessGroup:
		return true
	}
	return false
}

// Response for account resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectAccount struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                   `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectAccountType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectAccountJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectAccountJSON contains the
// JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAccount]
type accountTagGetResponseResourceTaggingTaggedResourceObjectAccountJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectAccountJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAccount) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectAccountType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectAccountTypeAccount AccountTagGetResponseResourceTaggingTaggedResourceObjectAccountType = "account"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAccountType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectAccountTypeAccount:
		return true
	}
	return false
}

// Response for ai_gateway resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectAIGateway struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                     `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayJSON contains
// the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAIGateway]
type accountTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectAIGateway) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAIGateway) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayTypeAIGateway AccountTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayType = "ai_gateway"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayTypeAIGateway:
		return true
	}
	return false
}

// Response for alerting_policy resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                          `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON
// contains the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy]
type accountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyTypeAlertingPolicy AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyType = "alerting_policy"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyTypeAlertingPolicy:
		return true
	}
	return false
}

// Response for alerting_webhook resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                           `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON
// contains the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook]
type accountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookTypeAlertingWebhook AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookType = "alerting_webhook"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookTypeAlertingWebhook:
		return true
	}
	return false
}

// Response for api_gateway_operation resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                               `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                          `json:"zone_id" api:"required"`
	JSON   accountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON
// contains the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation]
type accountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationTypeAPIGatewayOperation AccountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType = "api_gateway_operation"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationTypeAPIGatewayOperation:
		return true
	}
	return false
}

// Response for cloudflared_tunnel resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON
// contains the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel]
type accountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelTypeCloudflaredTunnel AccountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType = "cloudflared_tunnel"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelTypeCloudflaredTunnel:
		return true
	}
	return false
}

// Response for custom_certificate resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                        `json:"zone_id" api:"required"`
	JSON   accountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON
// contains the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate]
type accountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateTypeCustomCertificate AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateType = "custom_certificate"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateTypeCustomCertificate:
		return true
	}
	return false
}

// Response for custom_hostname resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                          `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                     `json:"zone_id" api:"required"`
	JSON   accountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON
// contains the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname]
type accountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameTypeCustomHostname AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameType = "custom_hostname"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameTypeCustomHostname:
		return true
	}
	return false
}

// Response for d1_database resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectD1Database struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                      `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON contains
// the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectD1Database]
type accountTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectD1Database) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseTypeD1Database AccountTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseType = "d1_database"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseTypeD1Database:
		return true
	}
	return false
}

// Response for dns_record resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                     `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                `json:"zone_id" api:"required"`
	JSON   accountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordJSON contains
// the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord]
type accountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordTypeDNSRecord AccountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordType = "dns_record"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordTypeDNSRecord:
		return true
	}
	return false
}

// Response for durable_object_namespace resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                                  `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace]
type accountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceTypeDurableObjectNamespace AccountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType = "durable_object_namespace"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

// Response for gateway_list resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayList struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                       `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayListType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectGatewayListJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectGatewayListJSON contains
// the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayList]
type accountTagGetResponseResourceTaggingTaggedResourceObjectGatewayListJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectGatewayListJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayList) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayListType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayListTypeGatewayList AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayListType = "gateway_list"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayListType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayListTypeGatewayList:
		return true
	}
	return false
}

// Response for gateway_rule resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                       `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON contains
// the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule]
type accountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleTypeGatewayRule AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleType = "gateway_rule"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleTypeGatewayRule:
		return true
	}
	return false
}

// Response for image resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectImage struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                 `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectImageType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectImageJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectImageJSON contains the
// JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectImage]
type accountTagGetResponseResourceTaggingTaggedResourceObjectImageJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectImageJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectImage) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectImageType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectImageTypeImage AccountTagGetResponseResourceTaggingTaggedResourceObjectImageType = "image"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectImageType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectImageTypeImage:
		return true
	}
	return false
}

// Response for kv_namespace resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                       `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON contains
// the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace]
type accountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceTypeKVNamespace AccountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceType = "kv_namespace"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

// Response for managed_client_certificate resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                                    `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                               `json:"zone_id" api:"required"`
	JSON   accountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON
// contains the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate]
type accountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateTypeManagedClientCertificate AccountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType = "managed_client_certificate"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateTypeManagedClientCertificate:
		return true
	}
	return false
}

// Response for queue resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectQueue struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                 `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectQueueType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectQueueJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectQueueJSON contains the
// JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectQueue]
type accountTagGetResponseResourceTaggingTaggedResourceObjectQueueJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectQueueJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectQueue) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectQueueType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectQueueTypeQueue AccountTagGetResponseResourceTaggingTaggedResourceObjectQueueType = "queue"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectQueueType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectQueueTypeQueue:
		return true
	}
	return false
}

// Response for r2_bucket resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                    `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectR2BucketType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectR2BucketJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectR2BucketJSON contains
// the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket]
type accountTagGetResponseResourceTaggingTaggedResourceObjectR2BucketJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectR2BucketType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectR2BucketTypeR2Bucket AccountTagGetResponseResourceTaggingTaggedResourceObjectR2BucketType = "r2_bucket"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectR2BucketType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectR2BucketTypeR2Bucket:
		return true
	}
	return false
}

// Response for resource_share resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectResourceShare struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                         `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectResourceShareType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectResourceShareJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectResourceShareJSON
// contains the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectResourceShare]
type accountTagGetResponseResourceTaggingTaggedResourceObjectResourceShareJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectResourceShare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectResourceShareJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectResourceShare) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectResourceShareType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectResourceShareTypeResourceShare AccountTagGetResponseResourceTaggingTaggedResourceObjectResourceShareType = "resource_share"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectResourceShareType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectResourceShareTypeResourceShare:
		return true
	}
	return false
}

// Response for stream_live_input resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                           `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON
// contains the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput]
type accountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputTypeStreamLiveInput AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputType = "stream_live_input"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputTypeStreamLiveInput:
		return true
	}
	return false
}

// Response for stream_video resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                       `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoJSON contains
// the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo]
type accountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoTypeStreamVideo AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoType = "stream_video"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoTypeStreamVideo:
		return true
	}
	return false
}

// Response for worker resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectWorker struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                  `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerType `json:"type" api:"required"`
	JSON accountTagGetResponseResourceTaggingTaggedResourceObjectWorkerJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectWorkerJSON contains the
// JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectWorker]
type accountTagGetResponseResourceTaggingTaggedResourceObjectWorkerJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectWorkerJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectWorker) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerTypeWorker AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerType = "worker"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerTypeWorker:
		return true
	}
	return false
}

// Response for worker_version resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                         `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionType `json:"type" api:"required"`
	// Worker ID is required only for worker_version resources
	WorkerID string                                                                    `json:"worker_id" api:"required"`
	JSON     accountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON
// contains the JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion]
type accountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	WorkerID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionTypeWorkerVersion AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionType = "worker_version"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionTypeWorkerVersion:
		return true
	}
	return false
}

// Response for zone resources
type AccountTagGetResponseResourceTaggingTaggedResourceObjectZone struct {
	// Identifies the unique resource.
	ID string `json:"id" api:"required"`
	// ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>"
	// where the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the
	// tags map canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients
	// should treat ETags as opaque strings and pass them back via the If-Match header
	// on write operations.
	Etag string `json:"etag" api:"required"`
	// Human-readable name of the resource.
	Name string `json:"name" api:"required"`
	// Contains key-value pairs of tags.
	Tags map[string]string                                                `json:"tags" api:"required"`
	Type AccountTagGetResponseResourceTaggingTaggedResourceObjectZoneType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                           `json:"zone_id" api:"required"`
	JSON   accountTagGetResponseResourceTaggingTaggedResourceObjectZoneJSON `json:"-"`
}

// accountTagGetResponseResourceTaggingTaggedResourceObjectZoneJSON contains the
// JSON metadata for the struct
// [AccountTagGetResponseResourceTaggingTaggedResourceObjectZone]
type accountTagGetResponseResourceTaggingTaggedResourceObjectZoneJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseResourceTaggingTaggedResourceObjectZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseResourceTaggingTaggedResourceObjectZoneJSON) RawJSON() string {
	return r.raw
}

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectZone) implementsAccountTagGetResponse() {
}

type AccountTagGetResponseResourceTaggingTaggedResourceObjectZoneType string

const (
	AccountTagGetResponseResourceTaggingTaggedResourceObjectZoneTypeZone AccountTagGetResponseResourceTaggingTaggedResourceObjectZoneType = "zone"
)

func (r AccountTagGetResponseResourceTaggingTaggedResourceObjectZoneType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseResourceTaggingTaggedResourceObjectZoneTypeZone:
		return true
	}
	return false
}

type AccountTagGetResponseType string

const (
	AccountTagGetResponseTypeAccessApplication        AccountTagGetResponseType = "access_application"
	AccountTagGetResponseTypeAccessApplicationPolicy  AccountTagGetResponseType = "access_application_policy"
	AccountTagGetResponseTypeAccessGroup              AccountTagGetResponseType = "access_group"
	AccountTagGetResponseTypeAccount                  AccountTagGetResponseType = "account"
	AccountTagGetResponseTypeAIGateway                AccountTagGetResponseType = "ai_gateway"
	AccountTagGetResponseTypeAlertingPolicy           AccountTagGetResponseType = "alerting_policy"
	AccountTagGetResponseTypeAlertingWebhook          AccountTagGetResponseType = "alerting_webhook"
	AccountTagGetResponseTypeAPIGatewayOperation      AccountTagGetResponseType = "api_gateway_operation"
	AccountTagGetResponseTypeCloudflaredTunnel        AccountTagGetResponseType = "cloudflared_tunnel"
	AccountTagGetResponseTypeCustomCertificate        AccountTagGetResponseType = "custom_certificate"
	AccountTagGetResponseTypeCustomHostname           AccountTagGetResponseType = "custom_hostname"
	AccountTagGetResponseTypeD1Database               AccountTagGetResponseType = "d1_database"
	AccountTagGetResponseTypeDNSRecord                AccountTagGetResponseType = "dns_record"
	AccountTagGetResponseTypeDurableObjectNamespace   AccountTagGetResponseType = "durable_object_namespace"
	AccountTagGetResponseTypeGatewayList              AccountTagGetResponseType = "gateway_list"
	AccountTagGetResponseTypeGatewayRule              AccountTagGetResponseType = "gateway_rule"
	AccountTagGetResponseTypeImage                    AccountTagGetResponseType = "image"
	AccountTagGetResponseTypeKVNamespace              AccountTagGetResponseType = "kv_namespace"
	AccountTagGetResponseTypeManagedClientCertificate AccountTagGetResponseType = "managed_client_certificate"
	AccountTagGetResponseTypeQueue                    AccountTagGetResponseType = "queue"
	AccountTagGetResponseTypeR2Bucket                 AccountTagGetResponseType = "r2_bucket"
	AccountTagGetResponseTypeResourceShare            AccountTagGetResponseType = "resource_share"
	AccountTagGetResponseTypeStreamLiveInput          AccountTagGetResponseType = "stream_live_input"
	AccountTagGetResponseTypeStreamVideo              AccountTagGetResponseType = "stream_video"
	AccountTagGetResponseTypeWorker                   AccountTagGetResponseType = "worker"
	AccountTagGetResponseTypeWorkerVersion            AccountTagGetResponseType = "worker_version"
	AccountTagGetResponseTypeZone                     AccountTagGetResponseType = "zone"
)

func (r AccountTagGetResponseType) IsKnown() bool {
	switch r {
	case AccountTagGetResponseTypeAccessApplication, AccountTagGetResponseTypeAccessApplicationPolicy, AccountTagGetResponseTypeAccessGroup, AccountTagGetResponseTypeAccount, AccountTagGetResponseTypeAIGateway, AccountTagGetResponseTypeAlertingPolicy, AccountTagGetResponseTypeAlertingWebhook, AccountTagGetResponseTypeAPIGatewayOperation, AccountTagGetResponseTypeCloudflaredTunnel, AccountTagGetResponseTypeCustomCertificate, AccountTagGetResponseTypeCustomHostname, AccountTagGetResponseTypeD1Database, AccountTagGetResponseTypeDNSRecord, AccountTagGetResponseTypeDurableObjectNamespace, AccountTagGetResponseTypeGatewayList, AccountTagGetResponseTypeGatewayRule, AccountTagGetResponseTypeImage, AccountTagGetResponseTypeKVNamespace, AccountTagGetResponseTypeManagedClientCertificate, AccountTagGetResponseTypeQueue, AccountTagGetResponseTypeR2Bucket, AccountTagGetResponseTypeResourceShare, AccountTagGetResponseTypeStreamLiveInput, AccountTagGetResponseTypeStreamVideo, AccountTagGetResponseTypeWorker, AccountTagGetResponseTypeWorkerVersion, AccountTagGetResponseTypeZone:
		return true
	}
	return false
}

type AccountTagUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Request body schema for setting tags on account-level resources.
	Body    AccountTagUpdateParamsBodyUnion `json:"body" api:"required"`
	IfMatch param.Field[string]             `header:"If-Match"`
}

func (r AccountTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request body schema for setting tags on account-level resources.
type AccountTagUpdateParamsBody struct {
	// Identifies the unique resource.
	ResourceID param.Field[string] `json:"resource_id" api:"required"`
	// Enum for base account-level resource types (those with no extra required
	// fields).
	ResourceType param.Field[AccountTagUpdateParamsBodyResourceType] `json:"resource_type" api:"required"`
	Tags         param.Field[interface{}]                            `json:"tags"`
	// Worker ID is required only for worker_version resources
	WorkerID param.Field[string] `json:"worker_id"`
}

func (r AccountTagUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountTagUpdateParamsBody) implementsAccountTagUpdateParamsBodyUnion() {}

// Request body schema for setting tags on account-level resources.
//
// Satisfied by
// [resource_tagging.AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersion],
// [resource_tagging.AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBase],
// [AccountTagUpdateParamsBody].
type AccountTagUpdateParamsBodyUnion interface {
	implementsAccountTagUpdateParamsBodyUnion()
}

// Request body schema for deleting tags from account-level resources.
type AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersion struct {
	// Identifies the unique resource.
	ResourceID param.Field[string] `json:"resource_id" api:"required"`
	// Enum for base account-level resource types (those with no extra required
	// fields).
	ResourceType param.Field[AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType] `json:"resource_type" api:"required"`
	// Worker ID is required only for worker_version resources
	WorkerID param.Field[string] `json:"worker_id" api:"required"`
	// Contains key-value pairs of tags.
	Tags param.Field[map[string]string] `json:"tags"`
}

func (r AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersion) implementsAccountTagUpdateParamsBodyUnion() {
}

// Enum for base account-level resource types (those with no extra required
// fields).
type AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType string

const (
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeAccessApplication      AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "access_application"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeAccessGroup            AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "access_group"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeAccount                AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "account"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeAIGateway              AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "ai_gateway"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeAlertingPolicy         AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "alerting_policy"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeAlertingWebhook        AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "alerting_webhook"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeCloudflaredTunnel      AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "cloudflared_tunnel"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeD1Database             AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "d1_database"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeDurableObjectNamespace AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "durable_object_namespace"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeGatewayList            AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "gateway_list"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeGatewayRule            AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "gateway_rule"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeImage                  AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "image"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeKVNamespace            AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "kv_namespace"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeQueue                  AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "queue"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeR2Bucket               AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "r2_bucket"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeResourceShare          AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "resource_share"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeStreamLiveInput        AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "stream_live_input"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeStreamVideo            AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "stream_video"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeWorker                 AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "worker"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeWorkerVersion          AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType = "worker_version"
)

func (r AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceType) IsKnown() bool {
	switch r {
	case AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeAccessApplication, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeAccessGroup, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeAccount, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeAIGateway, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeAlertingPolicy, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeAlertingWebhook, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeCloudflaredTunnel, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeD1Database, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeDurableObjectNamespace, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeGatewayList, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeGatewayRule, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeImage, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeKVNamespace, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeQueue, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeR2Bucket, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeResourceShare, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeStreamLiveInput, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeStreamVideo, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeWorker, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelWorkerVersionResourceTypeWorkerVersion:
		return true
	}
	return false
}

// Request body schema for deleting tags from account-level resources.
type AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBase struct {
	// Identifies the unique resource.
	ResourceID param.Field[string] `json:"resource_id" api:"required"`
	// Enum for base account-level resource types (those with no extra required
	// fields).
	ResourceType param.Field[AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType] `json:"resource_type" api:"required"`
	// Contains key-value pairs of tags.
	Tags param.Field[map[string]string] `json:"tags"`
}

func (r AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBase) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBase) implementsAccountTagUpdateParamsBodyUnion() {
}

// Enum for base account-level resource types (those with no extra required
// fields).
type AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType string

const (
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeAccessApplication      AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "access_application"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeAccessGroup            AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "access_group"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeAccount                AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "account"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeAIGateway              AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "ai_gateway"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeAlertingPolicy         AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "alerting_policy"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeAlertingWebhook        AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "alerting_webhook"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeCloudflaredTunnel      AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "cloudflared_tunnel"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeD1Database             AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "d1_database"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeDurableObjectNamespace AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "durable_object_namespace"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeGatewayList            AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "gateway_list"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeGatewayRule            AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "gateway_rule"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeImage                  AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "image"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeKVNamespace            AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "kv_namespace"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeQueue                  AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "queue"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeR2Bucket               AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "r2_bucket"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeResourceShare          AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "resource_share"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeStreamLiveInput        AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "stream_live_input"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeStreamVideo            AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "stream_video"
	AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeWorker                 AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType = "worker"
)

func (r AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceType) IsKnown() bool {
	switch r {
	case AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeAccessApplication, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeAccessGroup, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeAccount, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeAIGateway, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeAlertingPolicy, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeAlertingWebhook, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeCloudflaredTunnel, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeD1Database, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeDurableObjectNamespace, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeGatewayList, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeGatewayRule, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeImage, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeKVNamespace, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeQueue, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeR2Bucket, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeResourceShare, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeStreamLiveInput, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeStreamVideo, AccountTagUpdateParamsBodyResourceTaggingSetTagsRequestAccountLevelBaseResourceTypeWorker:
		return true
	}
	return false
}

// Enum for base account-level resource types (those with no extra required
// fields).
type AccountTagUpdateParamsBodyResourceType string

const (
	AccountTagUpdateParamsBodyResourceTypeAccessApplication      AccountTagUpdateParamsBodyResourceType = "access_application"
	AccountTagUpdateParamsBodyResourceTypeAccessGroup            AccountTagUpdateParamsBodyResourceType = "access_group"
	AccountTagUpdateParamsBodyResourceTypeAccount                AccountTagUpdateParamsBodyResourceType = "account"
	AccountTagUpdateParamsBodyResourceTypeAIGateway              AccountTagUpdateParamsBodyResourceType = "ai_gateway"
	AccountTagUpdateParamsBodyResourceTypeAlertingPolicy         AccountTagUpdateParamsBodyResourceType = "alerting_policy"
	AccountTagUpdateParamsBodyResourceTypeAlertingWebhook        AccountTagUpdateParamsBodyResourceType = "alerting_webhook"
	AccountTagUpdateParamsBodyResourceTypeCloudflaredTunnel      AccountTagUpdateParamsBodyResourceType = "cloudflared_tunnel"
	AccountTagUpdateParamsBodyResourceTypeD1Database             AccountTagUpdateParamsBodyResourceType = "d1_database"
	AccountTagUpdateParamsBodyResourceTypeDurableObjectNamespace AccountTagUpdateParamsBodyResourceType = "durable_object_namespace"
	AccountTagUpdateParamsBodyResourceTypeGatewayList            AccountTagUpdateParamsBodyResourceType = "gateway_list"
	AccountTagUpdateParamsBodyResourceTypeGatewayRule            AccountTagUpdateParamsBodyResourceType = "gateway_rule"
	AccountTagUpdateParamsBodyResourceTypeImage                  AccountTagUpdateParamsBodyResourceType = "image"
	AccountTagUpdateParamsBodyResourceTypeKVNamespace            AccountTagUpdateParamsBodyResourceType = "kv_namespace"
	AccountTagUpdateParamsBodyResourceTypeQueue                  AccountTagUpdateParamsBodyResourceType = "queue"
	AccountTagUpdateParamsBodyResourceTypeR2Bucket               AccountTagUpdateParamsBodyResourceType = "r2_bucket"
	AccountTagUpdateParamsBodyResourceTypeResourceShare          AccountTagUpdateParamsBodyResourceType = "resource_share"
	AccountTagUpdateParamsBodyResourceTypeStreamLiveInput        AccountTagUpdateParamsBodyResourceType = "stream_live_input"
	AccountTagUpdateParamsBodyResourceTypeStreamVideo            AccountTagUpdateParamsBodyResourceType = "stream_video"
	AccountTagUpdateParamsBodyResourceTypeWorker                 AccountTagUpdateParamsBodyResourceType = "worker"
	AccountTagUpdateParamsBodyResourceTypeWorkerVersion          AccountTagUpdateParamsBodyResourceType = "worker_version"
)

func (r AccountTagUpdateParamsBodyResourceType) IsKnown() bool {
	switch r {
	case AccountTagUpdateParamsBodyResourceTypeAccessApplication, AccountTagUpdateParamsBodyResourceTypeAccessGroup, AccountTagUpdateParamsBodyResourceTypeAccount, AccountTagUpdateParamsBodyResourceTypeAIGateway, AccountTagUpdateParamsBodyResourceTypeAlertingPolicy, AccountTagUpdateParamsBodyResourceTypeAlertingWebhook, AccountTagUpdateParamsBodyResourceTypeCloudflaredTunnel, AccountTagUpdateParamsBodyResourceTypeD1Database, AccountTagUpdateParamsBodyResourceTypeDurableObjectNamespace, AccountTagUpdateParamsBodyResourceTypeGatewayList, AccountTagUpdateParamsBodyResourceTypeGatewayRule, AccountTagUpdateParamsBodyResourceTypeImage, AccountTagUpdateParamsBodyResourceTypeKVNamespace, AccountTagUpdateParamsBodyResourceTypeQueue, AccountTagUpdateParamsBodyResourceTypeR2Bucket, AccountTagUpdateParamsBodyResourceTypeResourceShare, AccountTagUpdateParamsBodyResourceTypeStreamLiveInput, AccountTagUpdateParamsBodyResourceTypeStreamVideo, AccountTagUpdateParamsBodyResourceTypeWorker, AccountTagUpdateParamsBodyResourceTypeWorkerVersion:
		return true
	}
	return false
}

type AccountTagUpdateResponseEnvelope struct {
	Errors   []AccountTagUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccountTagUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccountTagUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	// Response for access_application resources
	Result AccountTagUpdateResponse             `json:"result"`
	JSON   accountTagUpdateResponseEnvelopeJSON `json:"-"`
}

// accountTagUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountTagUpdateResponseEnvelope]
type accountTagUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccountTagUpdateResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           AccountTagUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accountTagUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accountTagUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccountTagUpdateResponseEnvelopeErrors]
type accountTagUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountTagUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccountTagUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    accountTagUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accountTagUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [AccountTagUpdateResponseEnvelopeErrorsSource]
type accountTagUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountTagUpdateResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           AccountTagUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accountTagUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accountTagUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccountTagUpdateResponseEnvelopeMessages]
type accountTagUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountTagUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccountTagUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    accountTagUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accountTagUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [AccountTagUpdateResponseEnvelopeMessagesSource]
type accountTagUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountTagUpdateResponseEnvelopeSuccess bool

const (
	AccountTagUpdateResponseEnvelopeSuccessTrue AccountTagUpdateResponseEnvelopeSuccess = true
)

func (r AccountTagUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccountTagUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccountTagDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	IfMatch   param.Field[string] `header:"If-Match"`
}

type AccountTagGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The ID of the resource to retrieve tags for.
	ResourceID param.Field[string] `query:"resource_id" api:"required"`
	// The type of the resource.
	ResourceType param.Field[AccountTagGetParamsResourceType] `query:"resource_type" api:"required"`
	// Worker identifier. Required for worker_version resources.
	WorkerID param.Field[string] `query:"worker_id"`
}

// URLQuery serializes [AccountTagGetParams]'s query parameters as `url.Values`.
func (r AccountTagGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The type of the resource.
type AccountTagGetParamsResourceType string

const (
	AccountTagGetParamsResourceTypeAccessApplication      AccountTagGetParamsResourceType = "access_application"
	AccountTagGetParamsResourceTypeAccessGroup            AccountTagGetParamsResourceType = "access_group"
	AccountTagGetParamsResourceTypeAccount                AccountTagGetParamsResourceType = "account"
	AccountTagGetParamsResourceTypeAIGateway              AccountTagGetParamsResourceType = "ai_gateway"
	AccountTagGetParamsResourceTypeAlertingPolicy         AccountTagGetParamsResourceType = "alerting_policy"
	AccountTagGetParamsResourceTypeAlertingWebhook        AccountTagGetParamsResourceType = "alerting_webhook"
	AccountTagGetParamsResourceTypeCloudflaredTunnel      AccountTagGetParamsResourceType = "cloudflared_tunnel"
	AccountTagGetParamsResourceTypeD1Database             AccountTagGetParamsResourceType = "d1_database"
	AccountTagGetParamsResourceTypeDurableObjectNamespace AccountTagGetParamsResourceType = "durable_object_namespace"
	AccountTagGetParamsResourceTypeGatewayList            AccountTagGetParamsResourceType = "gateway_list"
	AccountTagGetParamsResourceTypeGatewayRule            AccountTagGetParamsResourceType = "gateway_rule"
	AccountTagGetParamsResourceTypeImage                  AccountTagGetParamsResourceType = "image"
	AccountTagGetParamsResourceTypeKVNamespace            AccountTagGetParamsResourceType = "kv_namespace"
	AccountTagGetParamsResourceTypeQueue                  AccountTagGetParamsResourceType = "queue"
	AccountTagGetParamsResourceTypeR2Bucket               AccountTagGetParamsResourceType = "r2_bucket"
	AccountTagGetParamsResourceTypeResourceShare          AccountTagGetParamsResourceType = "resource_share"
	AccountTagGetParamsResourceTypeStreamLiveInput        AccountTagGetParamsResourceType = "stream_live_input"
	AccountTagGetParamsResourceTypeStreamVideo            AccountTagGetParamsResourceType = "stream_video"
	AccountTagGetParamsResourceTypeWorker                 AccountTagGetParamsResourceType = "worker"
	AccountTagGetParamsResourceTypeWorkerVersion          AccountTagGetParamsResourceType = "worker_version"
)

func (r AccountTagGetParamsResourceType) IsKnown() bool {
	switch r {
	case AccountTagGetParamsResourceTypeAccessApplication, AccountTagGetParamsResourceTypeAccessGroup, AccountTagGetParamsResourceTypeAccount, AccountTagGetParamsResourceTypeAIGateway, AccountTagGetParamsResourceTypeAlertingPolicy, AccountTagGetParamsResourceTypeAlertingWebhook, AccountTagGetParamsResourceTypeCloudflaredTunnel, AccountTagGetParamsResourceTypeD1Database, AccountTagGetParamsResourceTypeDurableObjectNamespace, AccountTagGetParamsResourceTypeGatewayList, AccountTagGetParamsResourceTypeGatewayRule, AccountTagGetParamsResourceTypeImage, AccountTagGetParamsResourceTypeKVNamespace, AccountTagGetParamsResourceTypeQueue, AccountTagGetParamsResourceTypeR2Bucket, AccountTagGetParamsResourceTypeResourceShare, AccountTagGetParamsResourceTypeStreamLiveInput, AccountTagGetParamsResourceTypeStreamVideo, AccountTagGetParamsResourceTypeWorker, AccountTagGetParamsResourceTypeWorkerVersion:
		return true
	}
	return false
}

type AccountTagGetResponseEnvelope struct {
	Errors   []AccountTagGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccountTagGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccountTagGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// Response for access_application resources
	Result AccountTagGetResponse             `json:"result"`
	JSON   accountTagGetResponseEnvelopeJSON `json:"-"`
}

// accountTagGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountTagGetResponseEnvelope]
type accountTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccountTagGetResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           AccountTagGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accountTagGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accountTagGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccountTagGetResponseEnvelopeErrors]
type accountTagGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountTagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccountTagGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    accountTagGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accountTagGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AccountTagGetResponseEnvelopeErrorsSource]
type accountTagGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountTagGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AccountTagGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accountTagGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accountTagGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccountTagGetResponseEnvelopeMessages]
type accountTagGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountTagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccountTagGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    accountTagGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accountTagGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [AccountTagGetResponseEnvelopeMessagesSource]
type accountTagGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTagGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTagGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountTagGetResponseEnvelopeSuccess bool

const (
	AccountTagGetResponseEnvelopeSuccessTrue AccountTagGetResponseEnvelopeSuccess = true
)

func (r AccountTagGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccountTagGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
