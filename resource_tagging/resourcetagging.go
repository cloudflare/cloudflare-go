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
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
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
	Summary     *SummaryService
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
	r.Summary = NewSummaryService(opts...)
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
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time `json:"tags_updated_at" format:"date-time"`
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
	TagsUpdatedAt       apijson.Field
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
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRuleset],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGateway],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicy],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhook],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificate],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostname],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeployment],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicy],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySet],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkload],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1Database],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecord],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayList],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRule],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheck],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImage],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTarget],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespace],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancer],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitor],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPool],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificate],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProject],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueue],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2Bucket],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShare],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInput],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideo],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndex],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorker],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRoute],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersion],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZone],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRuleset].
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
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRuleset],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGateway],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicy],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhook],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificate],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostname],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeployment],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicy],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySet],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkload],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1Database],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecord],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayList],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRule],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheck],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImage],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTarget],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespace],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancer],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitor],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPool],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificate],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProject],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueue],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2Bucket],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShare],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInput],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideo],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndex],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorker],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRoute],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersion],
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZone] or
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRuleset].
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
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRuleset{}),
			DiscriminatorValue: "account_ruleset",
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
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeployment{}),
			DiscriminatorValue: "cws_deployment",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicy{}),
			DiscriminatorValue: "cws_policy",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySet{}),
			DiscriminatorValue: "cws_policy_set",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkload{}),
			DiscriminatorValue: "cws_workload",
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
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheck{}),
			DiscriminatorValue: "healthcheck",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImage{}),
			DiscriminatorValue: "image",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTarget{}),
			DiscriminatorValue: "infrastructure_target",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancer{}),
			DiscriminatorValue: "load_balancer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitor{}),
			DiscriminatorValue: "load_balancer_monitor",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPool{}),
			DiscriminatorValue: "load_balancer_pool",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificate{}),
			DiscriminatorValue: "managed_client_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProject{}),
			DiscriminatorValue: "pages_project",
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
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndex{}),
			DiscriminatorValue: "vectorize_index",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorker{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRoute{}),
			DiscriminatorValue: "worker_route",
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
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRuleset{}),
			DiscriminatorValue: "zone_ruleset",
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                   `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                           `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplication]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                         `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string `json:"zone_id" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                                 `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON `json:"-"`
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
	TagsUpdatedAt       apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                     `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroup]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccessGroupJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                         `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                 `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccount]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// Response for account_ruleset resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRuleset struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRulesetType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                        `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRulesetJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRulesetJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRuleset]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRulesetJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRulesetJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRuleset) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRulesetType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRulesetTypeAccountRuleset ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRulesetType = "account_ruleset"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRulesetType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAccountRulesetTypeAccountRuleset:
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                           `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                   `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGateway]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAIGatewayJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                        `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicy]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                 `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                         `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhook]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                     `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string `json:"zone_id" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                             `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	ZoneID        apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                   `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                           `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                   `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string `json:"zone_id" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                           `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificate]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	ZoneID        apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string `json:"zone_id" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                        `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostname]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	ZoneID        apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// Response for cws_deployment resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeployment struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                               `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeploymentType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                       `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeploymentJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeploymentJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeployment]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeploymentJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeploymentJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeployment) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeploymentType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeploymentTypeCwsDeployment ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeploymentType = "cws_deployment"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeploymentType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsDeploymentTypeCwsDeployment:
		return true
	}
	return false
}

// Response for cws_policy resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicy struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                           `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicyType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                   `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicyJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicyJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicy]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicyJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicyJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicy) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicyType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicyTypeCwsPolicy ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicyType = "cws_policy"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicyType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicyTypeCwsPolicy:
		return true
	}
	return false
}

// Response for cws_policy_set resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySet struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                              `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySetType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                      `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySetJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySetJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySet]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySetJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySetJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySet) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySetType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySetTypeCwsPolicySet ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySetType = "cws_policy_set"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySetType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsPolicySetTypeCwsPolicySet:
		return true
	}
	return false
}

// Response for cws_workload resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkload struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkloadType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                     `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkloadJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkloadJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkload]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkloadJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkloadJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkload) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkloadType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkloadTypeCwsWorkload ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkloadType = "cws_workload"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkloadType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectCwsWorkloadTypeCwsWorkload:
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                            `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                    `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectD1Database]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                           `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string `json:"zone_id" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                   `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecord]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectDNSRecordJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	ZoneID        apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                        `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                                `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                     `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayList]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayListJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                     `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRule]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// Response for healthcheck resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheck struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheckType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string `json:"zone_id" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                     `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheckJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheckJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheck]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheckJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	ZoneID        apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheckJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheck) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheckType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheckTypeHealthcheck ResourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheckType = "healthcheck"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheckType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectHealthcheckTypeHealthcheck:
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                       `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImageType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                               `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectImageJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectImageJSON contains
// the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectImage]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectImageJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// Response for infrastructure_target resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTarget struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                      `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTargetType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                              `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTargetJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTargetJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTarget]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTargetJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTargetJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTarget) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTargetType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTargetTypeInfrastructureTarget ResourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTargetType = "infrastructure_target"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTargetType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectInfrastructureTargetTypeInfrastructureTarget:
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                     `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespace]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// Response for load_balancer resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancer struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                              `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string `json:"zone_id" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                      `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancer]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	ZoneID        apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancer) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerTypeLoadBalancer ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerType = "load_balancer"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerTypeLoadBalancer:
		return true
	}
	return false
}

// Response for load_balancer_monitor resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitor struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                     `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitorType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                             `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitorJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitorJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitor]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitorJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitorJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitor) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitorType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitorTypeLoadBalancerMonitor ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitorType = "load_balancer_monitor"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitorType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerMonitorTypeLoadBalancerMonitor:
		return true
	}
	return false
}

// Response for load_balancer_pool resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPool struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                  `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPoolType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                          `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPoolJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPoolJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPool]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPoolJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPoolJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPool) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPoolType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPoolTypeLoadBalancerPool ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPoolType = "load_balancer_pool"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPoolType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectLoadBalancerPoolTypeLoadBalancerPool:
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                          `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string `json:"zone_id" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                                  `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificate]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	ZoneID        apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// Response for pages_project resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProject struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                              `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProjectType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                      `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProjectJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProjectJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProject]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProjectJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProjectJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProject) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProjectType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProjectTypePagesProject ResourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProjectType = "pages_project"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProjectType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectPagesProjectTypePagesProject:
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                       `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                               `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueJSON contains
// the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectQueue]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectQueueJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                          `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                  `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectR2Bucket]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectR2BucketJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                               `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                       `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShare]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectResourceShareJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                 `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                         `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInput]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                     `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideo]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectStreamVideoJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// Response for vectorize_index resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndex struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                                `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndexType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                        `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndexJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndexJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndex]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndexJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndexJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndex) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndexType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndexTypeVectorizeIndex ResourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndexType = "vectorize_index"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndexType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectVectorizeIndexTypeVectorizeIndex:
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                        `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerType `json:"type" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorker]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// Response for worker_route resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRoute struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRouteType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string `json:"zone_id" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                     `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRouteJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRouteJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRoute]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRouteJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	ZoneID        apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRouteJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRoute) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRouteType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRouteTypeWorkerRoute ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRouteType = "worker_route"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRouteType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerRouteTypeWorkerRoute:
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                               `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionType `json:"type" api:"required"`
	// Worker ID is required only for worker_version resources
	WorkerID string `json:"worker_id" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                       `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersion]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	WorkerID      apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                      `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string `json:"zone_id" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                              `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneJSON contains
// the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZone]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	ZoneID        apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// Response for zone_ruleset resources
type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRuleset struct {
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
	// Contains key-value pairs of tags. Keys may contain at most 256 characters.
	// Values may contain at most 1024 characters and may be empty for key-only tags.
	Tags map[string]string                                                             `json:"tags" api:"required"`
	Type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRulesetType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string `json:"zone_id" api:"required"`
	// Monotonic version of the resource's tags: the timestamp assigned when the tags
	// were last written. Returned by read endpoints, by 2PC prepare (the version that
	// will be assigned on commit, unless a concurrent write lands first, in which case
	// a newer version is assigned), and by 2PC commit (the authoritative committed
	// version). Omitted for untagged resources and delete commits: a deleted resource
	// has no current version, and deletions are ordered by event order rather than by
	// version.
	TagsUpdatedAt time.Time                                                                     `json:"tags_updated_at" format:"date-time"`
	JSON          resourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRulesetJSON `json:"-"`
}

// resourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRulesetJSON
// contains the JSON metadata for the struct
// [ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRuleset]
type resourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRulesetJSON struct {
	ID            apijson.Field
	Etag          apijson.Field
	Name          apijson.Field
	Tags          apijson.Field
	Type          apijson.Field
	ZoneID        apijson.Field
	TagsUpdatedAt apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRulesetJSON) RawJSON() string {
	return r.raw
}

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRuleset) implementsResourceTaggingListResponse() {
}

type ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRulesetType string

const (
	ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRulesetTypeZoneRuleset ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRulesetType = "zone_ruleset"
)

func (r ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRulesetType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseResourceTaggingTaggedResourceObjectZoneRulesetTypeZoneRuleset:
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
	ResourceTaggingListResponseTypeAccountRuleset           ResourceTaggingListResponseType = "account_ruleset"
	ResourceTaggingListResponseTypeAIGateway                ResourceTaggingListResponseType = "ai_gateway"
	ResourceTaggingListResponseTypeAlertingPolicy           ResourceTaggingListResponseType = "alerting_policy"
	ResourceTaggingListResponseTypeAlertingWebhook          ResourceTaggingListResponseType = "alerting_webhook"
	ResourceTaggingListResponseTypeAPIGatewayOperation      ResourceTaggingListResponseType = "api_gateway_operation"
	ResourceTaggingListResponseTypeCloudflaredTunnel        ResourceTaggingListResponseType = "cloudflared_tunnel"
	ResourceTaggingListResponseTypeCustomCertificate        ResourceTaggingListResponseType = "custom_certificate"
	ResourceTaggingListResponseTypeCustomHostname           ResourceTaggingListResponseType = "custom_hostname"
	ResourceTaggingListResponseTypeCwsDeployment            ResourceTaggingListResponseType = "cws_deployment"
	ResourceTaggingListResponseTypeCwsPolicy                ResourceTaggingListResponseType = "cws_policy"
	ResourceTaggingListResponseTypeCwsPolicySet             ResourceTaggingListResponseType = "cws_policy_set"
	ResourceTaggingListResponseTypeCwsWorkload              ResourceTaggingListResponseType = "cws_workload"
	ResourceTaggingListResponseTypeD1Database               ResourceTaggingListResponseType = "d1_database"
	ResourceTaggingListResponseTypeDNSRecord                ResourceTaggingListResponseType = "dns_record"
	ResourceTaggingListResponseTypeDurableObjectNamespace   ResourceTaggingListResponseType = "durable_object_namespace"
	ResourceTaggingListResponseTypeGatewayList              ResourceTaggingListResponseType = "gateway_list"
	ResourceTaggingListResponseTypeGatewayRule              ResourceTaggingListResponseType = "gateway_rule"
	ResourceTaggingListResponseTypeHealthcheck              ResourceTaggingListResponseType = "healthcheck"
	ResourceTaggingListResponseTypeImage                    ResourceTaggingListResponseType = "image"
	ResourceTaggingListResponseTypeInfrastructureTarget     ResourceTaggingListResponseType = "infrastructure_target"
	ResourceTaggingListResponseTypeKVNamespace              ResourceTaggingListResponseType = "kv_namespace"
	ResourceTaggingListResponseTypeLoadBalancer             ResourceTaggingListResponseType = "load_balancer"
	ResourceTaggingListResponseTypeLoadBalancerMonitor      ResourceTaggingListResponseType = "load_balancer_monitor"
	ResourceTaggingListResponseTypeLoadBalancerPool         ResourceTaggingListResponseType = "load_balancer_pool"
	ResourceTaggingListResponseTypeManagedClientCertificate ResourceTaggingListResponseType = "managed_client_certificate"
	ResourceTaggingListResponseTypePagesProject             ResourceTaggingListResponseType = "pages_project"
	ResourceTaggingListResponseTypeQueue                    ResourceTaggingListResponseType = "queue"
	ResourceTaggingListResponseTypeR2Bucket                 ResourceTaggingListResponseType = "r2_bucket"
	ResourceTaggingListResponseTypeResourceShare            ResourceTaggingListResponseType = "resource_share"
	ResourceTaggingListResponseTypeStreamLiveInput          ResourceTaggingListResponseType = "stream_live_input"
	ResourceTaggingListResponseTypeStreamVideo              ResourceTaggingListResponseType = "stream_video"
	ResourceTaggingListResponseTypeVectorizeIndex           ResourceTaggingListResponseType = "vectorize_index"
	ResourceTaggingListResponseTypeWorker                   ResourceTaggingListResponseType = "worker"
	ResourceTaggingListResponseTypeWorkerRoute              ResourceTaggingListResponseType = "worker_route"
	ResourceTaggingListResponseTypeWorkerVersion            ResourceTaggingListResponseType = "worker_version"
	ResourceTaggingListResponseTypeZone                     ResourceTaggingListResponseType = "zone"
	ResourceTaggingListResponseTypeZoneRuleset              ResourceTaggingListResponseType = "zone_ruleset"
)

func (r ResourceTaggingListResponseType) IsKnown() bool {
	switch r {
	case ResourceTaggingListResponseTypeAccessApplication, ResourceTaggingListResponseTypeAccessApplicationPolicy, ResourceTaggingListResponseTypeAccessGroup, ResourceTaggingListResponseTypeAccount, ResourceTaggingListResponseTypeAccountRuleset, ResourceTaggingListResponseTypeAIGateway, ResourceTaggingListResponseTypeAlertingPolicy, ResourceTaggingListResponseTypeAlertingWebhook, ResourceTaggingListResponseTypeAPIGatewayOperation, ResourceTaggingListResponseTypeCloudflaredTunnel, ResourceTaggingListResponseTypeCustomCertificate, ResourceTaggingListResponseTypeCustomHostname, ResourceTaggingListResponseTypeCwsDeployment, ResourceTaggingListResponseTypeCwsPolicy, ResourceTaggingListResponseTypeCwsPolicySet, ResourceTaggingListResponseTypeCwsWorkload, ResourceTaggingListResponseTypeD1Database, ResourceTaggingListResponseTypeDNSRecord, ResourceTaggingListResponseTypeDurableObjectNamespace, ResourceTaggingListResponseTypeGatewayList, ResourceTaggingListResponseTypeGatewayRule, ResourceTaggingListResponseTypeHealthcheck, ResourceTaggingListResponseTypeImage, ResourceTaggingListResponseTypeInfrastructureTarget, ResourceTaggingListResponseTypeKVNamespace, ResourceTaggingListResponseTypeLoadBalancer, ResourceTaggingListResponseTypeLoadBalancerMonitor, ResourceTaggingListResponseTypeLoadBalancerPool, ResourceTaggingListResponseTypeManagedClientCertificate, ResourceTaggingListResponseTypePagesProject, ResourceTaggingListResponseTypeQueue, ResourceTaggingListResponseTypeR2Bucket, ResourceTaggingListResponseTypeResourceShare, ResourceTaggingListResponseTypeStreamLiveInput, ResourceTaggingListResponseTypeStreamVideo, ResourceTaggingListResponseTypeVectorizeIndex, ResourceTaggingListResponseTypeWorker, ResourceTaggingListResponseTypeWorkerRoute, ResourceTaggingListResponseTypeWorkerVersion, ResourceTaggingListResponseTypeZone, ResourceTaggingListResponseTypeZoneRuleset:
		return true
	}
	return false
}

type ResourceTaggingListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter by resource ID. Can be repeated up to 50 times to filter by multiple IDs.
	// Example: ?id=abc&id=def
	ID param.Field[[]string] `query:"id"`
	// Match `tag` keys and values case-insensitively. Stored casing is unchanged.
	// Example: ?tag=environment=production&case_insensitive=true
	CaseInsensitive param.Field[bool] `query:"case_insensitive"`
	// Cursor for pagination.
	Cursor param.Field[string] `query:"cursor"`
	// Filter by resource name. Performs a case-insensitive substring match. Example:
	// ?name=my-zone
	Name param.Field[string] `query:"name"`
	// Filter resources by tag criteria. This parameter can be repeated multiple times,
	// with AND logic between parameters.
	//
	// Supported syntax:
	//
	// - **Key-only**: `tag=<key>` - Resource must have the tag key (e.g.,
	//   `tag=production`)
	// - **Key-value**: `tag=<key>=<value>` - Resource must have the tag with specific
	//   value (e.g., `tag=env=prod`)
	// - **Multiple values (OR)**: `tag=<key>=<v1>,<v2>` - Resource must have tag with
	//   any of the values (e.g., `tag=env=prod,staging`)
	// - **Negate key-only**: `tag=!<key>` - Resource must not have the tag key (e.g.,
	//   `tag=!archived`)
	// - **Negate key-value**: `tag=<key>!=<value>` - Resource must not have the tag
	//   with specific value (e.g., `tag=region!=us-west-1`)
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
	ResourceTaggingListParamsTypeAccountRuleset           ResourceTaggingListParamsType = "account_ruleset"
	ResourceTaggingListParamsTypeAIGateway                ResourceTaggingListParamsType = "ai_gateway"
	ResourceTaggingListParamsTypeAlertingPolicy           ResourceTaggingListParamsType = "alerting_policy"
	ResourceTaggingListParamsTypeAlertingWebhook          ResourceTaggingListParamsType = "alerting_webhook"
	ResourceTaggingListParamsTypeAPIGatewayOperation      ResourceTaggingListParamsType = "api_gateway_operation"
	ResourceTaggingListParamsTypeCloudflaredTunnel        ResourceTaggingListParamsType = "cloudflared_tunnel"
	ResourceTaggingListParamsTypeCustomCertificate        ResourceTaggingListParamsType = "custom_certificate"
	ResourceTaggingListParamsTypeCustomHostname           ResourceTaggingListParamsType = "custom_hostname"
	ResourceTaggingListParamsTypeCwsDeployment            ResourceTaggingListParamsType = "cws_deployment"
	ResourceTaggingListParamsTypeCwsPolicy                ResourceTaggingListParamsType = "cws_policy"
	ResourceTaggingListParamsTypeCwsPolicySet             ResourceTaggingListParamsType = "cws_policy_set"
	ResourceTaggingListParamsTypeCwsWorkload              ResourceTaggingListParamsType = "cws_workload"
	ResourceTaggingListParamsTypeD1Database               ResourceTaggingListParamsType = "d1_database"
	ResourceTaggingListParamsTypeDNSRecord                ResourceTaggingListParamsType = "dns_record"
	ResourceTaggingListParamsTypeDurableObjectNamespace   ResourceTaggingListParamsType = "durable_object_namespace"
	ResourceTaggingListParamsTypeGatewayList              ResourceTaggingListParamsType = "gateway_list"
	ResourceTaggingListParamsTypeGatewayRule              ResourceTaggingListParamsType = "gateway_rule"
	ResourceTaggingListParamsTypeHealthcheck              ResourceTaggingListParamsType = "healthcheck"
	ResourceTaggingListParamsTypeImage                    ResourceTaggingListParamsType = "image"
	ResourceTaggingListParamsTypeInfrastructureTarget     ResourceTaggingListParamsType = "infrastructure_target"
	ResourceTaggingListParamsTypeKVNamespace              ResourceTaggingListParamsType = "kv_namespace"
	ResourceTaggingListParamsTypeLoadBalancer             ResourceTaggingListParamsType = "load_balancer"
	ResourceTaggingListParamsTypeLoadBalancerMonitor      ResourceTaggingListParamsType = "load_balancer_monitor"
	ResourceTaggingListParamsTypeLoadBalancerPool         ResourceTaggingListParamsType = "load_balancer_pool"
	ResourceTaggingListParamsTypeManagedClientCertificate ResourceTaggingListParamsType = "managed_client_certificate"
	ResourceTaggingListParamsTypePagesProject             ResourceTaggingListParamsType = "pages_project"
	ResourceTaggingListParamsTypeQueue                    ResourceTaggingListParamsType = "queue"
	ResourceTaggingListParamsTypeR2Bucket                 ResourceTaggingListParamsType = "r2_bucket"
	ResourceTaggingListParamsTypeResourceShare            ResourceTaggingListParamsType = "resource_share"
	ResourceTaggingListParamsTypeStreamLiveInput          ResourceTaggingListParamsType = "stream_live_input"
	ResourceTaggingListParamsTypeStreamVideo              ResourceTaggingListParamsType = "stream_video"
	ResourceTaggingListParamsTypeVectorizeIndex           ResourceTaggingListParamsType = "vectorize_index"
	ResourceTaggingListParamsTypeWorker                   ResourceTaggingListParamsType = "worker"
	ResourceTaggingListParamsTypeWorkerRoute              ResourceTaggingListParamsType = "worker_route"
	ResourceTaggingListParamsTypeWorkerVersion            ResourceTaggingListParamsType = "worker_version"
	ResourceTaggingListParamsTypeZone                     ResourceTaggingListParamsType = "zone"
	ResourceTaggingListParamsTypeZoneRuleset              ResourceTaggingListParamsType = "zone_ruleset"
)

func (r ResourceTaggingListParamsType) IsKnown() bool {
	switch r {
	case ResourceTaggingListParamsTypeAccessApplication, ResourceTaggingListParamsTypeAccessApplicationPolicy, ResourceTaggingListParamsTypeAccessGroup, ResourceTaggingListParamsTypeAccount, ResourceTaggingListParamsTypeAccountRuleset, ResourceTaggingListParamsTypeAIGateway, ResourceTaggingListParamsTypeAlertingPolicy, ResourceTaggingListParamsTypeAlertingWebhook, ResourceTaggingListParamsTypeAPIGatewayOperation, ResourceTaggingListParamsTypeCloudflaredTunnel, ResourceTaggingListParamsTypeCustomCertificate, ResourceTaggingListParamsTypeCustomHostname, ResourceTaggingListParamsTypeCwsDeployment, ResourceTaggingListParamsTypeCwsPolicy, ResourceTaggingListParamsTypeCwsPolicySet, ResourceTaggingListParamsTypeCwsWorkload, ResourceTaggingListParamsTypeD1Database, ResourceTaggingListParamsTypeDNSRecord, ResourceTaggingListParamsTypeDurableObjectNamespace, ResourceTaggingListParamsTypeGatewayList, ResourceTaggingListParamsTypeGatewayRule, ResourceTaggingListParamsTypeHealthcheck, ResourceTaggingListParamsTypeImage, ResourceTaggingListParamsTypeInfrastructureTarget, ResourceTaggingListParamsTypeKVNamespace, ResourceTaggingListParamsTypeLoadBalancer, ResourceTaggingListParamsTypeLoadBalancerMonitor, ResourceTaggingListParamsTypeLoadBalancerPool, ResourceTaggingListParamsTypeManagedClientCertificate, ResourceTaggingListParamsTypePagesProject, ResourceTaggingListParamsTypeQueue, ResourceTaggingListParamsTypeR2Bucket, ResourceTaggingListParamsTypeResourceShare, ResourceTaggingListParamsTypeStreamLiveInput, ResourceTaggingListParamsTypeStreamVideo, ResourceTaggingListParamsTypeVectorizeIndex, ResourceTaggingListParamsTypeWorker, ResourceTaggingListParamsTypeWorkerRoute, ResourceTaggingListParamsTypeWorkerVersion, ResourceTaggingListParamsTypeZone, ResourceTaggingListParamsTypeZoneRuleset:
		return true
	}
	return false
}
