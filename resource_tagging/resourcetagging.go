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
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/tidwall/gjson"
)

// ResourceTaggingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceTaggingService] method instead.
type ResourceTaggingService struct {
	Options     []option.RequestOption
	AccountTags *AccountTagService
	ZoneTags    *ZoneTagService
	Keys        *KeyService
	Values      *ValueService
}

// NewResourceTaggingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResourceTaggingService(opts ...option.RequestOption) (r *ResourceTaggingService) {
	r = &ResourceTaggingService{}
	r.Options = opts
	r.AccountTags = NewAccountTagService(opts...)
	r.ZoneTags = NewZoneTagService(opts...)
	r.Keys = NewKeyService(opts...)
	r.Values = NewValueService(opts...)
	return
}

// Lists all tagged resources for an account.
func (r *ResourceTaggingService) List(ctx context.Context, params ResourceTaggingListParams, opts ...option.RequestOption) (res *pagination.CursorPaginationAfter[ResourceTaggingListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/tags/resources", params.AccountID)
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

// Lists all tagged resources for an account.
func (r *ResourceTaggingService) ListAutoPaging(ctx context.Context, params ResourceTaggingListParams, opts ...option.RequestOption) *pagination.CursorPaginationAfterAutoPager[ResourceTaggingListResponse] {
	return pagination.NewCursorPaginationAfterAutoPager(r.List(ctx, params, opts...))
}

// Response for access_application resources
type ResourceTaggingListResponse struct {
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
	Tags interface{}                     `json:"tags" api:"required"`
	Type ResourceTaggingListResponseType `json:"type" api:"required"`
	// Access application ID is required only for access_application_policy resources
	AccessApplicationID string `json:"access_application_id" format:"uuid"`
	// Worker ID is required only for worker_version resources
	WorkerID string `json:"worker_id"`
	// Zone ID is required only for zone-level resources
	ZoneID string                          `json:"zone_id"`
	JSON   resourceTaggingListResponseJSON `json:"-"`
	union  ResourceTaggingListResponseUnion
}

// resourceTaggingListResponseJSON contains the JSON metadata for the struct
// [ResourceTaggingListResponse]
type resourceTaggingListResponseJSON struct {
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

func (r resourceTaggingListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ResourceTaggingListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ResourceTaggingListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ResourceTaggingListResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplication],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroup],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccount],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGateway],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicy],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhook],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificate],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostname],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1Database],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecord],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayList],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRule],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImage],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespace],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificate],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueue],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2Bucket],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShare],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInput],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideo],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorker],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersion],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZone].
func (r ResourceTaggingListResponse) AsUnion() ResourceTaggingListResponseUnion {
	return r.union
}

// Response for access_application resources
//
// Union satisfied by
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplication],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroup],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccount],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGateway],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicy],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhook],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificate],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostname],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1Database],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecord],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayList],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRule],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImage],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespace],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificate],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueue],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2Bucket],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShare],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInput],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideo],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorker],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersion] or
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZone].
type ResourceTaggingListResponseUnion interface {
	implementsResourceTaggingListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ResourceTaggingListResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplication{}),
			DiscriminatorValue: "access_application",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy{}),
			DiscriminatorValue: "access_application_policy",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroup{}),
			DiscriminatorValue: "access_group",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccount{}),
			DiscriminatorValue: "account",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGateway{}),
			DiscriminatorValue: "ai_gateway",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicy{}),
			DiscriminatorValue: "alerting_policy",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhook{}),
			DiscriminatorValue: "alerting_webhook",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation{}),
			DiscriminatorValue: "api_gateway_operation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel{}),
			DiscriminatorValue: "cloudflared_tunnel",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificate{}),
			DiscriminatorValue: "custom_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostname{}),
			DiscriminatorValue: "custom_hostname",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1Database{}),
			DiscriminatorValue: "d1_database",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecord{}),
			DiscriminatorValue: "dns_record",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayList{}),
			DiscriminatorValue: "gateway_list",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRule{}),
			DiscriminatorValue: "gateway_rule",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImage{}),
			DiscriminatorValue: "image",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificate{}),
			DiscriminatorValue: "managed_client_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShare{}),
			DiscriminatorValue: "resource_share",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInput{}),
			DiscriminatorValue: "stream_live_input",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideo{}),
			DiscriminatorValue: "stream_video",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorker{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersion{}),
			DiscriminatorValue: "worker_version",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZone{}),
			DiscriminatorValue: "zone",
		},
	)
}

// Response for access_application resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplication struct {
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
	Tags map[string]string                                                                   `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplication]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplication) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationTypeAccessApplication ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationType = "access_application"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationTypeAccessApplication:
		return true
	}
	return false
}

// Response for access_application_policy resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy struct {
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
	Tags map[string]string                                                                         `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                                    `json:"zone_id" api:"required"`
	JSON   resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON struct {
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

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyTypeAccessApplicationPolicy ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType = "access_application_policy"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyTypeAccessApplicationPolicy:
		return true
	}
	return false
}

// Response for access_group resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroup struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroup]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroup) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupTypeAccessGroup ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupType = "access_group"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupTypeAccessGroup:
		return true
	}
	return false
}

// Response for account resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccount struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccount]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccount) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountTypeAccount ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountType = "account"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountTypeAccount:
		return true
	}
	return false
}

// Response for ai_gateway resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGateway struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGateway]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGateway) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGateway) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayTypeAIGateway ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayType = "ai_gateway"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayTypeAIGateway:
		return true
	}
	return false
}

// Response for alerting_policy resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicy struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicy]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicy) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyTypeAlertingPolicy ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyType = "alerting_policy"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyTypeAlertingPolicy:
		return true
	}
	return false
}

// Response for alerting_webhook resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhook struct {
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
	Tags map[string]string                                                                 `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhook]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhook) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookTypeAlertingWebhook ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookType = "alerting_webhook"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookTypeAlertingWebhook:
		return true
	}
	return false
}

// Response for api_gateway_operation resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                                `json:"zone_id" api:"required"`
	JSON   resourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationTypeAPIGatewayOperation ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType = "api_gateway_operation"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationTypeAPIGatewayOperation:
		return true
	}
	return false
}

// Response for cloudflared_tunnel resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel struct {
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
	Tags map[string]string                                                                   `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelTypeCloudflaredTunnel ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType = "cloudflared_tunnel"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelTypeCloudflaredTunnel:
		return true
	}
	return false
}

// Response for custom_certificate resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificate struct {
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
	Tags map[string]string                                                                   `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                              `json:"zone_id" api:"required"`
	JSON   resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificate]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificate) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateTypeCustomCertificate ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateType = "custom_certificate"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateTypeCustomCertificate:
		return true
	}
	return false
}

// Response for custom_hostname resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostname struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                           `json:"zone_id" api:"required"`
	JSON   resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostname]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostname) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameTypeCustomHostname ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameType = "custom_hostname"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameTypeCustomHostname:
		return true
	}
	return false
}

// Response for d1_database resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1Database struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1Database]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1Database) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseTypeD1Database ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseType = "d1_database"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseTypeD1Database:
		return true
	}
	return false
}

// Response for dns_record resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecord struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                      `json:"zone_id" api:"required"`
	JSON   resourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecord]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecord) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordTypeDNSRecord ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordType = "dns_record"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordTypeDNSRecord:
		return true
	}
	return false
}

// Response for durable_object_namespace resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace struct {
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
	Tags map[string]string                                                                        `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceTypeDurableObjectNamespace ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

// Response for gateway_list resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayList struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayList]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayList) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListTypeGatewayList ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListType = "gateway_list"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListTypeGatewayList:
		return true
	}
	return false
}

// Response for gateway_rule resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRule struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRule]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRule) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleTypeGatewayRule ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleType = "gateway_rule"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleTypeGatewayRule:
		return true
	}
	return false
}

// Response for image resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImage struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImageType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectImageJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectImageJSON contains
// the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImage]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectImageJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectImageJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImage) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImageType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImageTypeImage ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImageType = "image"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImageType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImageTypeImage:
		return true
	}
	return false
}

// Response for kv_namespace resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespace struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespace]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespace) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceTypeKVNamespace ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceType = "kv_namespace"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

// Response for managed_client_certificate resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificate struct {
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
	Tags map[string]string                                                                          `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                                     `json:"zone_id" api:"required"`
	JSON   resourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificate]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificate) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateTypeManagedClientCertificate ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType = "managed_client_certificate"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateTypeManagedClientCertificate:
		return true
	}
	return false
}

// Response for queue resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueue struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueJSON contains
// the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueue]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueue) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueTypeQueue ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueType = "queue"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueTypeQueue:
		return true
	}
	return false
}

// Response for r2_bucket resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2Bucket struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2Bucket]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2Bucket) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketTypeR2Bucket ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketType = "r2_bucket"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketTypeR2Bucket:
		return true
	}
	return false
}

// Response for resource_share resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShare struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShare]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShare) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareTypeResourceShare ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareType = "resource_share"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareTypeResourceShare:
		return true
	}
	return false
}

// Response for stream_live_input resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInput struct {
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
	Tags map[string]string                                                                 `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInput]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInput) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputTypeStreamLiveInput ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputType = "stream_live_input"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputTypeStreamLiveInput:
		return true
	}
	return false
}

// Response for stream_video resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideo struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideo]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideo) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoTypeStreamVideo ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoType = "stream_video"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoTypeStreamVideo:
		return true
	}
	return false
}

// Response for worker resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorker struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerType `json:"type" api:"required"`
	JSON resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorker]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorker) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerTypeWorker ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerType = "worker"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerTypeWorker:
		return true
	}
	return false
}

// Response for worker_version resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersion struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionType `json:"type" api:"required"`
	// Worker ID is required only for worker_version resources
	WorkerID string                                                                          `json:"worker_id" api:"required"`
	JSON     resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersion]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	WorkerID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersion) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionTypeWorkerVersion ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionType = "worker_version"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionTypeWorkerVersion:
		return true
	}
	return false
}

// Response for zone resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZone struct {
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
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                 `json:"zone_id" api:"required"`
	JSON   resourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneJSON contains
// the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZone]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZone) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneTypeZone ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneType = "zone"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneTypeZone:
		return true
	}
	return false
}

type ResourceTaggingListResponseType string

const (
	ResourceTaggingListResponseTypeAccessApplication        ResourceTaggingListResponseType = "access_application"
	ResourceTaggingListResponseTypeAccessApplicationPolicy  ResourceTaggingListResponseType = "access_application_policy"
	ResourceTaggingListResponseTypeAccessGroup              ResourceTaggingListResponseType = "access_group"
	ResourceTaggingListResponseTypeAccount                  ResourceTaggingListResponseType = "account"
	ResourceTaggingListResponseTypeAIGateway                ResourceTaggingListResponseType = "ai_gateway"
	ResourceTaggingListResponseTypeAlertingPolicy           ResourceTaggingListResponseType = "alerting_policy"
	ResourceTaggingListResponseTypeAlertingWebhook          ResourceTaggingListResponseType = "alerting_webhook"
	ResourceTaggingListResponseTypeAPIGatewayOperation      ResourceTaggingListResponseType = "api_gateway_operation"
	ResourceTaggingListResponseTypeCloudflaredTunnel        ResourceTaggingListResponseType = "cloudflared_tunnel"
	ResourceTaggingListResponseTypeCustomCertificate        ResourceTaggingListResponseType = "custom_certificate"
	ResourceTaggingListResponseTypeCustomHostname           ResourceTaggingListResponseType = "custom_hostname"
	ResourceTaggingListResponseTypeD1Database               ResourceTaggingListResponseType = "d1_database"
	ResourceTaggingListResponseTypeDNSRecord                ResourceTaggingListResponseType = "dns_record"
	ResourceTaggingListResponseTypeDurableObjectNamespace   ResourceTaggingListResponseType = "durable_object_namespace"
	ResourceTaggingListResponseTypeGatewayList              ResourceTaggingListResponseType = "gateway_list"
	ResourceTaggingListResponseTypeGatewayRule              ResourceTaggingListResponseType = "gateway_rule"
	ResourceTaggingListResponseTypeImage                    ResourceTaggingListResponseType = "image"
	ResourceTaggingListResponseTypeKVNamespace              ResourceTaggingListResponseType = "kv_namespace"
	ResourceTaggingListResponseTypeManagedClientCertificate ResourceTaggingListResponseType = "managed_client_certificate"
	ResourceTaggingListResponseTypeQueue                    ResourceTaggingListResponseType = "queue"
	ResourceTaggingListResponseTypeR2Bucket                 ResourceTaggingListResponseType = "r2_bucket"
	ResourceTaggingListResponseTypeResourceShare            ResourceTaggingListResponseType = "resource_share"
	ResourceTaggingListResponseTypeStreamLiveInput          ResourceTaggingListResponseType = "stream_live_input"
	ResourceTaggingListResponseTypeStreamVideo              ResourceTaggingListResponseType = "stream_video"
	ResourceTaggingListResponseTypeWorker                   ResourceTaggingListResponseType = "worker"
	ResourceTaggingListResponseTypeWorkerVersion            ResourceTaggingListResponseType = "worker_version"
	ResourceTaggingListResponseTypeZone                     ResourceTaggingListResponseType = "zone"
)

func (r ResourceTaggingListResponseType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseTypeAccessApplication, ResourceTaggingListResponseTypeAccessApplicationPolicy, ResourceTaggingListResponseTypeAccessGroup, ResourceTaggingListResponseTypeAccount, ResourceTaggingListResponseTypeAIGateway, ResourceTaggingListResponseTypeAlertingPolicy, ResourceTaggingListResponseTypeAlertingWebhook, ResourceTaggingListResponseTypeAPIGatewayOperation, ResourceTaggingListResponseTypeCloudflaredTunnel, ResourceTaggingListResponseTypeCustomCertificate, ResourceTaggingListResponseTypeCustomHostname, ResourceTaggingListResponseTypeD1Database, ResourceTaggingListResponseTypeDNSRecord, ResourceTaggingListResponseTypeDurableObjectNamespace, ResourceTaggingListResponseTypeGatewayList, ResourceTaggingListResponseTypeGatewayRule, ResourceTaggingListResponseTypeImage, ResourceTaggingListResponseTypeKVNamespace, ResourceTaggingListResponseTypeManagedClientCertificate, ResourceTaggingListResponseTypeQueue, ResourceTaggingListResponseTypeR2Bucket, ResourceTaggingListResponseTypeResourceShare, ResourceTaggingListResponseTypeStreamLiveInput, ResourceTaggingListResponseTypeStreamVideo, ResourceTaggingListResponseTypeWorker, ResourceTaggingListResponseTypeWorkerVersion, ResourceTaggingListResponseTypeZone:
		return true
	}
	return false
}

type ResourceTaggingListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Cursor for pagination.
	Cursor param.Field[string] `query:"cursor"`
	// Filter resources by tag criteria. This parameter can be repeated multiple times,
	// with AND logic between parameters.
	//
	// Supported syntax:
	//
	//   - **Key-only**: `tag=<key>` - Resource must have the tag key (e.g.,
	//     `tag=production`)
	//   - **Key-value**: `tag=<key>=<value>` - Resource must have the tag with specific
	//     value (e.g., `tag=env=prod`)
	//   - **Multiple values (OR)**: `tag=<key>=<v1>,<v2>` - Resource must have tag with
	//     any of the values (e.g., `tag=env=prod,staging`)
	//   - **Negate key-only**: `tag=!<key>` - Resource must not have the tag key (e.g.,
	//     `tag=!archived`)
	//   - **Negate key-value**: `tag=<key>!=<value>` - Resource must not have the tag
	//     with specific value (e.g., `tag=region!=us-west-1`)
	//
	// Multiple tag parameters are combined with AND logic.
	Tag param.Field[[]string] `query:"tag"`
	// Filter by resource type. Can be repeated to filter by multiple types (OR logic).
	// Example: ?type=zone&type=worker
	Type param.Field[[]ResourceTaggingListParamsType] `query:"type"`
}

// URLQuery serializes [ResourceTaggingListParams]'s query parameters as
// `url.Values`.
func (r ResourceTaggingListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Identifies the type of resource.
type ResourceTaggingListParamsType string

const (
	ResourceTaggingListParamsTypeAccessApplication        ResourceTaggingListParamsType = "access_application"
	ResourceTaggingListParamsTypeAccessApplicationPolicy  ResourceTaggingListParamsType = "access_application_policy"
	ResourceTaggingListParamsTypeAccessGroup              ResourceTaggingListParamsType = "access_group"
	ResourceTaggingListParamsTypeAccount                  ResourceTaggingListParamsType = "account"
	ResourceTaggingListParamsTypeAIGateway                ResourceTaggingListParamsType = "ai_gateway"
	ResourceTaggingListParamsTypeAlertingPolicy           ResourceTaggingListParamsType = "alerting_policy"
	ResourceTaggingListParamsTypeAlertingWebhook          ResourceTaggingListParamsType = "alerting_webhook"
	ResourceTaggingListParamsTypeAPIGatewayOperation      ResourceTaggingListParamsType = "api_gateway_operation"
	ResourceTaggingListParamsTypeCloudflaredTunnel        ResourceTaggingListParamsType = "cloudflared_tunnel"
	ResourceTaggingListParamsTypeCustomCertificate        ResourceTaggingListParamsType = "custom_certificate"
	ResourceTaggingListParamsTypeCustomHostname           ResourceTaggingListParamsType = "custom_hostname"
	ResourceTaggingListParamsTypeD1Database               ResourceTaggingListParamsType = "d1_database"
	ResourceTaggingListParamsTypeDNSRecord                ResourceTaggingListParamsType = "dns_record"
	ResourceTaggingListParamsTypeDurableObjectNamespace   ResourceTaggingListParamsType = "durable_object_namespace"
	ResourceTaggingListParamsTypeGatewayList              ResourceTaggingListParamsType = "gateway_list"
	ResourceTaggingListParamsTypeGatewayRule              ResourceTaggingListParamsType = "gateway_rule"
	ResourceTaggingListParamsTypeImage                    ResourceTaggingListParamsType = "image"
	ResourceTaggingListParamsTypeKVNamespace              ResourceTaggingListParamsType = "kv_namespace"
	ResourceTaggingListParamsTypeManagedClientCertificate ResourceTaggingListParamsType = "managed_client_certificate"
	ResourceTaggingListParamsTypeQueue                    ResourceTaggingListParamsType = "queue"
	ResourceTaggingListParamsTypeR2Bucket                 ResourceTaggingListParamsType = "r2_bucket"
	ResourceTaggingListParamsTypeResourceShare            ResourceTaggingListParamsType = "resource_share"
	ResourceTaggingListParamsTypeStreamLiveInput          ResourceTaggingListParamsType = "stream_live_input"
	ResourceTaggingListParamsTypeStreamVideo              ResourceTaggingListParamsType = "stream_video"
	ResourceTaggingListParamsTypeWorker                   ResourceTaggingListParamsType = "worker"
	ResourceTaggingListParamsTypeWorkerVersion            ResourceTaggingListParamsType = "worker_version"
	ResourceTaggingListParamsTypeZone                     ResourceTaggingListParamsType = "zone"
)

func (r ResourceTaggingListParamsType) IsKnown() bool {
	switch r {
	case ResourceTaggingListParamsTypeAccessApplication, ResourceTaggingListParamsTypeAccessApplicationPolicy, ResourceTaggingListParamsTypeAccessGroup, ResourceTaggingListParamsTypeAccount, ResourceTaggingListParamsTypeAIGateway, ResourceTaggingListParamsTypeAlertingPolicy, ResourceTaggingListParamsTypeAlertingWebhook, ResourceTaggingListParamsTypeAPIGatewayOperation, ResourceTaggingListParamsTypeCloudflaredTunnel, ResourceTaggingListParamsTypeCustomCertificate, ResourceTaggingListParamsTypeCustomHostname, ResourceTaggingListParamsTypeD1Database, ResourceTaggingListParamsTypeDNSRecord, ResourceTaggingListParamsTypeDurableObjectNamespace, ResourceTaggingListParamsTypeGatewayList, ResourceTaggingListParamsTypeGatewayRule, ResourceTaggingListParamsTypeImage, ResourceTaggingListParamsTypeKVNamespace, ResourceTaggingListParamsTypeManagedClientCertificate, ResourceTaggingListParamsTypeQueue, ResourceTaggingListParamsTypeR2Bucket, ResourceTaggingListParamsTypeResourceShare, ResourceTaggingListParamsTypeStreamLiveInput, ResourceTaggingListParamsTypeStreamVideo, ResourceTaggingListParamsTypeWorker, ResourceTaggingListParamsTypeWorkerVersion, ResourceTaggingListParamsTypeZone:
		return true
	}
	return false
}
