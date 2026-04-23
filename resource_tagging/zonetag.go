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

// ZoneTagService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneTagService] method instead.
type ZoneTagService struct {
	Options []option.RequestOption
}

// NewZoneTagService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZoneTagService(opts ...option.RequestOption) (r *ZoneTagService) {
	r = &ZoneTagService{}
	r.Options = opts
	return
}

// Creates or updates tags for a specific zone-level resource. Replaces all
// existing tags for the resource.
func (r *ZoneTagService) Update(ctx context.Context, params ZoneTagUpdateParams, opts ...option.RequestOption) (res *ZoneTagUpdateResponse, err error) {
	var env ZoneTagUpdateResponseEnvelope
	if params.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/tags", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Removes all tags from a specific zone-level resource.
func (r *ZoneTagService) Delete(ctx context.Context, params ZoneTagDeleteParams, opts ...option.RequestOption) (err error) {
	if params.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/tags", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieves tags for a specific zone-level resource.
func (r *ZoneTagService) Get(ctx context.Context, params ZoneTagGetParams, opts ...option.RequestOption) (res *ZoneTagGetResponse, err error) {
	var env ZoneTagGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/tags", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Response for access_application resources
type ZoneTagUpdateResponse struct {
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
	Type ZoneTagUpdateResponseType `json:"type" api:"required"`
	// Access application ID is required only for access_application_policy resources
	AccessApplicationID string `json:"access_application_id" format:"uuid"`
	// Worker ID is required only for worker_version resources
	WorkerID string `json:"worker_id"`
	// Zone ID is required only for zone-level resources
	ZoneID string                    `json:"zone_id"`
	JSON   zoneTagUpdateResponseJSON `json:"-"`
	union  ZoneTagUpdateResponseUnion
}

// zoneTagUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneTagUpdateResponse]
type zoneTagUpdateResponseJSON struct {
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

func (r zoneTagUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ZoneTagUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ZoneTagUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZoneTagUpdateResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccount],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectImage],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueue],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorker],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectZone].
func (r ZoneTagUpdateResponse) AsUnion() ZoneTagUpdateResponseUnion {
	return r.union
}

// Response for access_application resources
//
// Union satisfied by
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccount],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectImage],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueue],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorker],
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion] or
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectZone].
type ZoneTagUpdateResponseUnion interface {
	implementsZoneTagUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneTagUpdateResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication{}),
			DiscriminatorValue: "access_application",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy{}),
			DiscriminatorValue: "access_application_policy",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup{}),
			DiscriminatorValue: "access_group",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccount{}),
			DiscriminatorValue: "account",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway{}),
			DiscriminatorValue: "ai_gateway",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy{}),
			DiscriminatorValue: "alerting_policy",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook{}),
			DiscriminatorValue: "alerting_webhook",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation{}),
			DiscriminatorValue: "api_gateway_operation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel{}),
			DiscriminatorValue: "cloudflared_tunnel",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate{}),
			DiscriminatorValue: "custom_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname{}),
			DiscriminatorValue: "custom_hostname",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database{}),
			DiscriminatorValue: "d1_database",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord{}),
			DiscriminatorValue: "dns_record",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList{}),
			DiscriminatorValue: "gateway_list",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule{}),
			DiscriminatorValue: "gateway_rule",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectImage{}),
			DiscriminatorValue: "image",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate{}),
			DiscriminatorValue: "managed_client_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare{}),
			DiscriminatorValue: "resource_share",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput{}),
			DiscriminatorValue: "stream_live_input",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo{}),
			DiscriminatorValue: "stream_video",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorker{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion{}),
			DiscriminatorValue: "worker_version",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectZone{}),
			DiscriminatorValue: "zone",
		},
	)
}

// Response for access_application resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON
// contains the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplication) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationTypeAccessApplication ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationType = "access_application"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationTypeAccessApplication:
		return true
	}
	return false
}

// Response for access_application_policy resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                              `json:"zone_id" api:"required"`
	JSON   zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON
// contains the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON struct {
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

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyTypeAccessApplicationPolicy ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType = "access_application_policy"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyTypeAccessApplicationPolicy:
		return true
	}
	return false
}

// Response for access_group resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupJSON contains
// the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroup) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupTypeAccessGroup ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupType = "access_group"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccessGroupTypeAccessGroup:
		return true
	}
	return false
}

// Response for account resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccount struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccountType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccountJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccountJSON contains the
// JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccount]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccountJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccountJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccount) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccountType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccountTypeAccount ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccountType = "account"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccountType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAccountTypeAccount:
		return true
	}
	return false
}

// Response for ai_gateway resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayJSON contains
// the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGateway) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayTypeAIGateway ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayType = "ai_gateway"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAIGatewayTypeAIGateway:
		return true
	}
	return false
}

// Response for alerting_policy resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON
// contains the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicy) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyTypeAlertingPolicy ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyType = "alerting_policy"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingPolicyTypeAlertingPolicy:
		return true
	}
	return false
}

// Response for alerting_webhook resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON
// contains the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhook) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookTypeAlertingWebhook ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookType = "alerting_webhook"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAlertingWebhookTypeAlertingWebhook:
		return true
	}
	return false
}

// Response for api_gateway_operation resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                          `json:"zone_id" api:"required"`
	JSON   zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON
// contains the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationTypeAPIGatewayOperation ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType = "api_gateway_operation"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationTypeAPIGatewayOperation:
		return true
	}
	return false
}

// Response for cloudflared_tunnel resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON
// contains the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelTypeCloudflaredTunnel ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType = "cloudflared_tunnel"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelTypeCloudflaredTunnel:
		return true
	}
	return false
}

// Response for custom_certificate resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                        `json:"zone_id" api:"required"`
	JSON   zoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON
// contains the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificate) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateTypeCustomCertificate ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateType = "custom_certificate"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomCertificateTypeCustomCertificate:
		return true
	}
	return false
}

// Response for custom_hostname resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                     `json:"zone_id" api:"required"`
	JSON   zoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON
// contains the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostname) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameTypeCustomHostname ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameType = "custom_hostname"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectCustomHostnameTypeCustomHostname:
		return true
	}
	return false
}

// Response for d1_database resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON contains
// the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1Database) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseTypeD1Database ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseType = "d1_database"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectD1DatabaseTypeD1Database:
		return true
	}
	return false
}

// Response for dns_record resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                `json:"zone_id" api:"required"`
	JSON   zoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordJSON contains
// the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecord) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordTypeDNSRecord ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordType = "dns_record"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDNSRecordTypeDNSRecord:
		return true
	}
	return false
}

// Response for durable_object_namespace resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceTypeDurableObjectNamespace ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

// Response for gateway_list resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListJSON contains
// the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayList) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListTypeGatewayList ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListType = "gateway_list"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayListTypeGatewayList:
		return true
	}
	return false
}

// Response for gateway_rule resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON contains
// the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRule) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleTypeGatewayRule ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleType = "gateway_rule"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectGatewayRuleTypeGatewayRule:
		return true
	}
	return false
}

// Response for image resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectImage struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectImageType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectImageJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectImageJSON contains the
// JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectImage]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectImageJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectImageJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectImage) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectImageType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectImageTypeImage ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectImageType = "image"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectImageType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectImageTypeImage:
		return true
	}
	return false
}

// Response for kv_namespace resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON contains
// the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespace) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceTypeKVNamespace ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceType = "kv_namespace"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

// Response for managed_client_certificate resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                               `json:"zone_id" api:"required"`
	JSON   zoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON
// contains the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificate) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateTypeManagedClientCertificate ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType = "managed_client_certificate"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectManagedClientCertificateTypeManagedClientCertificate:
		return true
	}
	return false
}

// Response for queue resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueue struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueueType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueueJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueueJSON contains the
// JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueue]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueueJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueueJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueue) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueueType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueueTypeQueue ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueueType = "queue"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueueType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectQueueTypeQueue:
		return true
	}
	return false
}

// Response for r2_bucket resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketJSON contains
// the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2Bucket) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketTypeR2Bucket ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketType = "r2_bucket"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectR2BucketTypeR2Bucket:
		return true
	}
	return false
}

// Response for resource_share resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareJSON
// contains the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShare) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareTypeResourceShare ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareType = "resource_share"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectResourceShareTypeResourceShare:
		return true
	}
	return false
}

// Response for stream_live_input resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON
// contains the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInput) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputTypeStreamLiveInput ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputType = "stream_live_input"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamLiveInputTypeStreamLiveInput:
		return true
	}
	return false
}

// Response for stream_video resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoJSON contains
// the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideo) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoTypeStreamVideo ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoType = "stream_video"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectStreamVideoTypeStreamVideo:
		return true
	}
	return false
}

// Response for worker resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorker struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerType `json:"type" api:"required"`
	JSON zoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerJSON contains the
// JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorker]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorker) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerTypeWorker ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerType = "worker"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerTypeWorker:
		return true
	}
	return false
}

// Response for worker_version resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionType `json:"type" api:"required"`
	// Worker ID is required only for worker_version resources
	WorkerID string                                                                    `json:"worker_id" api:"required"`
	JSON     zoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON
// contains the JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	WorkerID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersion) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionTypeWorkerVersion ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionType = "worker_version"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectWorkerVersionTypeWorkerVersion:
		return true
	}
	return false
}

// Response for zone resources
type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectZone struct {
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
	Type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectZoneType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                           `json:"zone_id" api:"required"`
	JSON   zoneTagUpdateResponseResourceTaggingTaggedResourceObjectZoneJSON `json:"-"`
}

// zoneTagUpdateResponseResourceTaggingTaggedResourceObjectZoneJSON contains the
// JSON metadata for the struct
// [ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectZone]
type zoneTagUpdateResponseResourceTaggingTaggedResourceObjectZoneJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseResourceTaggingTaggedResourceObjectZoneJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectZone) implementsZoneTagUpdateResponse() {
}

type ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectZoneType string

const (
	ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectZoneTypeZone ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectZoneType = "zone"
)

func (r ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectZoneType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseResourceTaggingTaggedResourceObjectZoneTypeZone:
		return true
	}
	return false
}

type ZoneTagUpdateResponseType string

const (
	ZoneTagUpdateResponseTypeAccessApplication        ZoneTagUpdateResponseType = "access_application"
	ZoneTagUpdateResponseTypeAccessApplicationPolicy  ZoneTagUpdateResponseType = "access_application_policy"
	ZoneTagUpdateResponseTypeAccessGroup              ZoneTagUpdateResponseType = "access_group"
	ZoneTagUpdateResponseTypeAccount                  ZoneTagUpdateResponseType = "account"
	ZoneTagUpdateResponseTypeAIGateway                ZoneTagUpdateResponseType = "ai_gateway"
	ZoneTagUpdateResponseTypeAlertingPolicy           ZoneTagUpdateResponseType = "alerting_policy"
	ZoneTagUpdateResponseTypeAlertingWebhook          ZoneTagUpdateResponseType = "alerting_webhook"
	ZoneTagUpdateResponseTypeAPIGatewayOperation      ZoneTagUpdateResponseType = "api_gateway_operation"
	ZoneTagUpdateResponseTypeCloudflaredTunnel        ZoneTagUpdateResponseType = "cloudflared_tunnel"
	ZoneTagUpdateResponseTypeCustomCertificate        ZoneTagUpdateResponseType = "custom_certificate"
	ZoneTagUpdateResponseTypeCustomHostname           ZoneTagUpdateResponseType = "custom_hostname"
	ZoneTagUpdateResponseTypeD1Database               ZoneTagUpdateResponseType = "d1_database"
	ZoneTagUpdateResponseTypeDNSRecord                ZoneTagUpdateResponseType = "dns_record"
	ZoneTagUpdateResponseTypeDurableObjectNamespace   ZoneTagUpdateResponseType = "durable_object_namespace"
	ZoneTagUpdateResponseTypeGatewayList              ZoneTagUpdateResponseType = "gateway_list"
	ZoneTagUpdateResponseTypeGatewayRule              ZoneTagUpdateResponseType = "gateway_rule"
	ZoneTagUpdateResponseTypeImage                    ZoneTagUpdateResponseType = "image"
	ZoneTagUpdateResponseTypeKVNamespace              ZoneTagUpdateResponseType = "kv_namespace"
	ZoneTagUpdateResponseTypeManagedClientCertificate ZoneTagUpdateResponseType = "managed_client_certificate"
	ZoneTagUpdateResponseTypeQueue                    ZoneTagUpdateResponseType = "queue"
	ZoneTagUpdateResponseTypeR2Bucket                 ZoneTagUpdateResponseType = "r2_bucket"
	ZoneTagUpdateResponseTypeResourceShare            ZoneTagUpdateResponseType = "resource_share"
	ZoneTagUpdateResponseTypeStreamLiveInput          ZoneTagUpdateResponseType = "stream_live_input"
	ZoneTagUpdateResponseTypeStreamVideo              ZoneTagUpdateResponseType = "stream_video"
	ZoneTagUpdateResponseTypeWorker                   ZoneTagUpdateResponseType = "worker"
	ZoneTagUpdateResponseTypeWorkerVersion            ZoneTagUpdateResponseType = "worker_version"
	ZoneTagUpdateResponseTypeZone                     ZoneTagUpdateResponseType = "zone"
)

func (r ZoneTagUpdateResponseType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseTypeAccessApplication, ZoneTagUpdateResponseTypeAccessApplicationPolicy, ZoneTagUpdateResponseTypeAccessGroup, ZoneTagUpdateResponseTypeAccount, ZoneTagUpdateResponseTypeAIGateway, ZoneTagUpdateResponseTypeAlertingPolicy, ZoneTagUpdateResponseTypeAlertingWebhook, ZoneTagUpdateResponseTypeAPIGatewayOperation, ZoneTagUpdateResponseTypeCloudflaredTunnel, ZoneTagUpdateResponseTypeCustomCertificate, ZoneTagUpdateResponseTypeCustomHostname, ZoneTagUpdateResponseTypeD1Database, ZoneTagUpdateResponseTypeDNSRecord, ZoneTagUpdateResponseTypeDurableObjectNamespace, ZoneTagUpdateResponseTypeGatewayList, ZoneTagUpdateResponseTypeGatewayRule, ZoneTagUpdateResponseTypeImage, ZoneTagUpdateResponseTypeKVNamespace, ZoneTagUpdateResponseTypeManagedClientCertificate, ZoneTagUpdateResponseTypeQueue, ZoneTagUpdateResponseTypeR2Bucket, ZoneTagUpdateResponseTypeResourceShare, ZoneTagUpdateResponseTypeStreamLiveInput, ZoneTagUpdateResponseTypeStreamVideo, ZoneTagUpdateResponseTypeWorker, ZoneTagUpdateResponseTypeWorkerVersion, ZoneTagUpdateResponseTypeZone:
		return true
	}
	return false
}

// Response for access_application resources
type ZoneTagGetResponse struct {
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
	Tags interface{}            `json:"tags" api:"required"`
	Type ZoneTagGetResponseType `json:"type" api:"required"`
	// Access application ID is required only for access_application_policy resources
	AccessApplicationID string `json:"access_application_id" format:"uuid"`
	// Worker ID is required only for worker_version resources
	WorkerID string `json:"worker_id"`
	// Zone ID is required only for zone-level resources
	ZoneID string                 `json:"zone_id"`
	JSON   zoneTagGetResponseJSON `json:"-"`
	union  ZoneTagGetResponseUnion
}

// zoneTagGetResponseJSON contains the JSON metadata for the struct
// [ZoneTagGetResponse]
type zoneTagGetResponseJSON struct {
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

func (r zoneTagGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ZoneTagGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ZoneTagGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZoneTagGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccount],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAIGateway],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectD1Database],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayList],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectImage],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectQueue],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShare],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorker],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectZone].
func (r ZoneTagGetResponse) AsUnion() ZoneTagGetResponseUnion {
	return r.union
}

// Response for access_application resources
//
// Union satisfied by
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccount],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAIGateway],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectD1Database],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayList],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectImage],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectQueue],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShare],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorker],
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion] or
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectZone].
type ZoneTagGetResponseUnion interface {
	implementsZoneTagGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneTagGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication{}),
			DiscriminatorValue: "access_application",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy{}),
			DiscriminatorValue: "access_application_policy",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup{}),
			DiscriminatorValue: "access_group",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccount{}),
			DiscriminatorValue: "account",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectAIGateway{}),
			DiscriminatorValue: "ai_gateway",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy{}),
			DiscriminatorValue: "alerting_policy",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook{}),
			DiscriminatorValue: "alerting_webhook",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation{}),
			DiscriminatorValue: "api_gateway_operation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel{}),
			DiscriminatorValue: "cloudflared_tunnel",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate{}),
			DiscriminatorValue: "custom_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname{}),
			DiscriminatorValue: "custom_hostname",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectD1Database{}),
			DiscriminatorValue: "d1_database",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord{}),
			DiscriminatorValue: "dns_record",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayList{}),
			DiscriminatorValue: "gateway_list",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule{}),
			DiscriminatorValue: "gateway_rule",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectImage{}),
			DiscriminatorValue: "image",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate{}),
			DiscriminatorValue: "managed_client_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShare{}),
			DiscriminatorValue: "resource_share",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput{}),
			DiscriminatorValue: "stream_live_input",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo{}),
			DiscriminatorValue: "stream_video",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorker{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion{}),
			DiscriminatorValue: "worker_version",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ZoneTagGetResponseResourceTaggingTaggedResourceObjectZone{}),
			DiscriminatorValue: "zone",
		},
	)
}

// Response for access_application resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON
// contains the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplication) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationTypeAccessApplication ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationType = "access_application"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationTypeAccessApplication:
		return true
	}
	return false
}

// Response for access_application_policy resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy struct {
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
	Tags map[string]string                                                                `json:"tags" api:"required"`
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                           `json:"zone_id" api:"required"`
	JSON   zoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON
// contains the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON struct {
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

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicy) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyTypeAccessApplicationPolicy ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType = "access_application_policy"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessApplicationPolicyTypeAccessApplicationPolicy:
		return true
	}
	return false
}

// Response for access_group resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupJSON contains
// the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroup) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupTypeAccessGroup ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupType = "access_group"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccessGroupTypeAccessGroup:
		return true
	}
	return false
}

// Response for account resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccount struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccountType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectAccountJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectAccountJSON contains the
// JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccount]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectAccountJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectAccountJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccount) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccountType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccountTypeAccount ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccountType = "account"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccountType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectAccountTypeAccount:
		return true
	}
	return false
}

// Response for ai_gateway resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAIGateway struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayJSON contains the
// JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAIGateway]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectAIGateway) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAIGateway) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayTypeAIGateway ZoneTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayType = "ai_gateway"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectAIGatewayTypeAIGateway:
		return true
	}
	return false
}

// Response for alerting_policy resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON contains
// the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicy) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyTypeAlertingPolicy ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyType = "alerting_policy"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingPolicyTypeAlertingPolicy:
		return true
	}
	return false
}

// Response for alerting_webhook resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON
// contains the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhook) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookTypeAlertingWebhook ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookType = "alerting_webhook"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectAlertingWebhookTypeAlertingWebhook:
		return true
	}
	return false
}

// Response for api_gateway_operation resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                       `json:"zone_id" api:"required"`
	JSON   zoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON
// contains the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperation) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationTypeAPIGatewayOperation ZoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType = "api_gateway_operation"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectAPIGatewayOperationTypeAPIGatewayOperation:
		return true
	}
	return false
}

// Response for cloudflared_tunnel resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON
// contains the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnel) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelTypeCloudflaredTunnel ZoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType = "cloudflared_tunnel"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectCloudflaredTunnelTypeCloudflaredTunnel:
		return true
	}
	return false
}

// Response for custom_certificate resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                     `json:"zone_id" api:"required"`
	JSON   zoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON
// contains the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificate) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateTypeCustomCertificate ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateType = "custom_certificate"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomCertificateTypeCustomCertificate:
		return true
	}
	return false
}

// Response for custom_hostname resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                  `json:"zone_id" api:"required"`
	JSON   zoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON contains
// the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostname) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameTypeCustomHostname ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameType = "custom_hostname"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectCustomHostnameTypeCustomHostname:
		return true
	}
	return false
}

// Response for d1_database resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectD1Database struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON contains the
// JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectD1Database]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectD1Database) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseTypeD1Database ZoneTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseType = "d1_database"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectD1DatabaseTypeD1Database:
		return true
	}
	return false
}

// Response for dns_record resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                             `json:"zone_id" api:"required"`
	JSON   zoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordJSON contains the
// JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecord) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordTypeDNSRecord ZoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordType = "dns_record"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectDNSRecordTypeDNSRecord:
		return true
	}
	return false
}

// Response for durable_object_namespace resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespace) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceTypeDurableObjectNamespace ZoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

// Response for gateway_list resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayList struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayListType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayListJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayListJSON contains
// the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayList]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayListJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayListJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayList) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayListType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayListTypeGatewayList ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayListType = "gateway_list"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayListType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayListTypeGatewayList:
		return true
	}
	return false
}

// Response for gateway_rule resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON contains
// the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRule) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleTypeGatewayRule ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleType = "gateway_rule"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectGatewayRuleTypeGatewayRule:
		return true
	}
	return false
}

// Response for image resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectImage struct {
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
	Tags map[string]string                                              `json:"tags" api:"required"`
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectImageType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectImageJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectImageJSON contains the JSON
// metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectImage]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectImageJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectImageJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectImage) implementsZoneTagGetResponse() {}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectImageType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectImageTypeImage ZoneTagGetResponseResourceTaggingTaggedResourceObjectImageType = "image"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectImageType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectImageTypeImage:
		return true
	}
	return false
}

// Response for kv_namespace resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON contains
// the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespace) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceTypeKVNamespace ZoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceType = "kv_namespace"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

// Response for managed_client_certificate resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                                            `json:"zone_id" api:"required"`
	JSON   zoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON
// contains the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificate) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateTypeManagedClientCertificate ZoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType = "managed_client_certificate"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectManagedClientCertificateTypeManagedClientCertificate:
		return true
	}
	return false
}

// Response for queue resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectQueue struct {
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
	Tags map[string]string                                              `json:"tags" api:"required"`
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectQueueType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectQueueJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectQueueJSON contains the JSON
// metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectQueue]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectQueueJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectQueueJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectQueue) implementsZoneTagGetResponse() {}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectQueueType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectQueueTypeQueue ZoneTagGetResponseResourceTaggingTaggedResourceObjectQueueType = "queue"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectQueueType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectQueueTypeQueue:
		return true
	}
	return false
}

// Response for r2_bucket resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectR2BucketType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectR2BucketJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectR2BucketJSON contains the
// JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectR2BucketJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectR2Bucket) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectR2BucketType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectR2BucketTypeR2Bucket ZoneTagGetResponseResourceTaggingTaggedResourceObjectR2BucketType = "r2_bucket"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectR2BucketType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectR2BucketTypeR2Bucket:
		return true
	}
	return false
}

// Response for resource_share resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShare struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShareType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShareJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShareJSON contains
// the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShare]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShareJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShareJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShare) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShareType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShareTypeResourceShare ZoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShareType = "resource_share"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShareType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectResourceShareTypeResourceShare:
		return true
	}
	return false
}

// Response for stream_live_input resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON
// contains the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInput) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputTypeStreamLiveInput ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputType = "stream_live_input"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamLiveInputTypeStreamLiveInput:
		return true
	}
	return false
}

// Response for stream_video resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoJSON contains
// the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideo) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoTypeStreamVideo ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoType = "stream_video"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectStreamVideoTypeStreamVideo:
		return true
	}
	return false
}

// Response for worker resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorker struct {
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
	Tags map[string]string                                               `json:"tags" api:"required"`
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerType `json:"type" api:"required"`
	JSON zoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerJSON contains the
// JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorker]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorker) implementsZoneTagGetResponse() {}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerTypeWorker ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerType = "worker"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerTypeWorker:
		return true
	}
	return false
}

// Response for worker_version resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion struct {
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
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionType `json:"type" api:"required"`
	// Worker ID is required only for worker_version resources
	WorkerID string                                                                 `json:"worker_id" api:"required"`
	JSON     zoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON contains
// the JSON metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	WorkerID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersion) implementsZoneTagGetResponse() {
}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionTypeWorkerVersion ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionType = "worker_version"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectWorkerVersionTypeWorkerVersion:
		return true
	}
	return false
}

// Response for zone resources
type ZoneTagGetResponseResourceTaggingTaggedResourceObjectZone struct {
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
	Tags map[string]string                                             `json:"tags" api:"required"`
	Type ZoneTagGetResponseResourceTaggingTaggedResourceObjectZoneType `json:"type" api:"required"`
	// Zone ID is required only for zone-level resources
	ZoneID string                                                        `json:"zone_id" api:"required"`
	JSON   zoneTagGetResponseResourceTaggingTaggedResourceObjectZoneJSON `json:"-"`
}

// zoneTagGetResponseResourceTaggingTaggedResourceObjectZoneJSON contains the JSON
// metadata for the struct
// [ZoneTagGetResponseResourceTaggingTaggedResourceObjectZone]
type zoneTagGetResponseResourceTaggingTaggedResourceObjectZoneJSON struct {
	ID          apijson.Field
	Etag        apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	ZoneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseResourceTaggingTaggedResourceObjectZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseResourceTaggingTaggedResourceObjectZoneJSON) RawJSON() string {
	return r.raw
}

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectZone) implementsZoneTagGetResponse() {}

type ZoneTagGetResponseResourceTaggingTaggedResourceObjectZoneType string

const (
	ZoneTagGetResponseResourceTaggingTaggedResourceObjectZoneTypeZone ZoneTagGetResponseResourceTaggingTaggedResourceObjectZoneType = "zone"
)

func (r ZoneTagGetResponseResourceTaggingTaggedResourceObjectZoneType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseResourceTaggingTaggedResourceObjectZoneTypeZone:
		return true
	}
	return false
}

type ZoneTagGetResponseType string

const (
	ZoneTagGetResponseTypeAccessApplication        ZoneTagGetResponseType = "access_application"
	ZoneTagGetResponseTypeAccessApplicationPolicy  ZoneTagGetResponseType = "access_application_policy"
	ZoneTagGetResponseTypeAccessGroup              ZoneTagGetResponseType = "access_group"
	ZoneTagGetResponseTypeAccount                  ZoneTagGetResponseType = "account"
	ZoneTagGetResponseTypeAIGateway                ZoneTagGetResponseType = "ai_gateway"
	ZoneTagGetResponseTypeAlertingPolicy           ZoneTagGetResponseType = "alerting_policy"
	ZoneTagGetResponseTypeAlertingWebhook          ZoneTagGetResponseType = "alerting_webhook"
	ZoneTagGetResponseTypeAPIGatewayOperation      ZoneTagGetResponseType = "api_gateway_operation"
	ZoneTagGetResponseTypeCloudflaredTunnel        ZoneTagGetResponseType = "cloudflared_tunnel"
	ZoneTagGetResponseTypeCustomCertificate        ZoneTagGetResponseType = "custom_certificate"
	ZoneTagGetResponseTypeCustomHostname           ZoneTagGetResponseType = "custom_hostname"
	ZoneTagGetResponseTypeD1Database               ZoneTagGetResponseType = "d1_database"
	ZoneTagGetResponseTypeDNSRecord                ZoneTagGetResponseType = "dns_record"
	ZoneTagGetResponseTypeDurableObjectNamespace   ZoneTagGetResponseType = "durable_object_namespace"
	ZoneTagGetResponseTypeGatewayList              ZoneTagGetResponseType = "gateway_list"
	ZoneTagGetResponseTypeGatewayRule              ZoneTagGetResponseType = "gateway_rule"
	ZoneTagGetResponseTypeImage                    ZoneTagGetResponseType = "image"
	ZoneTagGetResponseTypeKVNamespace              ZoneTagGetResponseType = "kv_namespace"
	ZoneTagGetResponseTypeManagedClientCertificate ZoneTagGetResponseType = "managed_client_certificate"
	ZoneTagGetResponseTypeQueue                    ZoneTagGetResponseType = "queue"
	ZoneTagGetResponseTypeR2Bucket                 ZoneTagGetResponseType = "r2_bucket"
	ZoneTagGetResponseTypeResourceShare            ZoneTagGetResponseType = "resource_share"
	ZoneTagGetResponseTypeStreamLiveInput          ZoneTagGetResponseType = "stream_live_input"
	ZoneTagGetResponseTypeStreamVideo              ZoneTagGetResponseType = "stream_video"
	ZoneTagGetResponseTypeWorker                   ZoneTagGetResponseType = "worker"
	ZoneTagGetResponseTypeWorkerVersion            ZoneTagGetResponseType = "worker_version"
	ZoneTagGetResponseTypeZone                     ZoneTagGetResponseType = "zone"
)

func (r ZoneTagGetResponseType) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseTypeAccessApplication, ZoneTagGetResponseTypeAccessApplicationPolicy, ZoneTagGetResponseTypeAccessGroup, ZoneTagGetResponseTypeAccount, ZoneTagGetResponseTypeAIGateway, ZoneTagGetResponseTypeAlertingPolicy, ZoneTagGetResponseTypeAlertingWebhook, ZoneTagGetResponseTypeAPIGatewayOperation, ZoneTagGetResponseTypeCloudflaredTunnel, ZoneTagGetResponseTypeCustomCertificate, ZoneTagGetResponseTypeCustomHostname, ZoneTagGetResponseTypeD1Database, ZoneTagGetResponseTypeDNSRecord, ZoneTagGetResponseTypeDurableObjectNamespace, ZoneTagGetResponseTypeGatewayList, ZoneTagGetResponseTypeGatewayRule, ZoneTagGetResponseTypeImage, ZoneTagGetResponseTypeKVNamespace, ZoneTagGetResponseTypeManagedClientCertificate, ZoneTagGetResponseTypeQueue, ZoneTagGetResponseTypeR2Bucket, ZoneTagGetResponseTypeResourceShare, ZoneTagGetResponseTypeStreamLiveInput, ZoneTagGetResponseTypeStreamVideo, ZoneTagGetResponseTypeWorker, ZoneTagGetResponseTypeWorkerVersion, ZoneTagGetResponseTypeZone:
		return true
	}
	return false
}

type ZoneTagUpdateParams struct {
	// Zone ID is required only for zone-level resources
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Request body schema for setting tags on zone-level resources.
	Body    ZoneTagUpdateParamsBodyUnion `json:"body" api:"required"`
	IfMatch param.Field[string]          `header:"If-Match"`
}

func (r ZoneTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request body schema for setting tags on zone-level resources.
type ZoneTagUpdateParamsBody struct {
	// Identifies the unique resource.
	ResourceID param.Field[string] `json:"resource_id" api:"required"`
	// Enum for base zone-level resource types (those with no extra required fields).
	ResourceType param.Field[ZoneTagUpdateParamsBodyResourceType] `json:"resource_type" api:"required"`
	// Access application ID is required only for access_application_policy resources
	AccessApplicationID param.Field[string]      `json:"access_application_id" format:"uuid"`
	Tags                param.Field[interface{}] `json:"tags"`
}

func (r ZoneTagUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneTagUpdateParamsBody) implementsZoneTagUpdateParamsBodyUnion() {}

// Request body schema for setting tags on zone-level resources.
//
// Satisfied by
// [resource_tagging.ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBase],
// [resource_tagging.ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicy],
// [ZoneTagUpdateParamsBody].
type ZoneTagUpdateParamsBodyUnion interface {
	implementsZoneTagUpdateParamsBodyUnion()
}

// Request body schema for deleting tags from zone-level resources. Zone ID comes
// from URL path.
type ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBase struct {
	// Identifies the unique resource.
	ResourceID param.Field[string] `json:"resource_id" api:"required"`
	// Enum for base zone-level resource types (those with no extra required fields).
	ResourceType param.Field[ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceType] `json:"resource_type" api:"required"`
	// Contains key-value pairs of tags.
	Tags param.Field[map[string]string] `json:"tags"`
}

func (r ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBase) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBase) implementsZoneTagUpdateParamsBodyUnion() {
}

// Enum for base zone-level resource types (those with no extra required fields).
type ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceType string

const (
	ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceTypeAPIGatewayOperation      ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceType = "api_gateway_operation"
	ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceTypeCustomCertificate        ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceType = "custom_certificate"
	ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceTypeCustomHostname           ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceType = "custom_hostname"
	ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceTypeDNSRecord                ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceType = "dns_record"
	ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceTypeManagedClientCertificate ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceType = "managed_client_certificate"
	ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceTypeZone                     ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceType = "zone"
)

func (r ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceTypeAPIGatewayOperation, ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceTypeCustomCertificate, ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceTypeCustomHostname, ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceTypeDNSRecord, ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceTypeManagedClientCertificate, ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelBaseResourceTypeZone:
		return true
	}
	return false
}

// Request body schema for deleting tags from zone-level resources. Zone ID comes
// from URL path.
type ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicy struct {
	// Access application ID is required only for access_application_policy resources
	AccessApplicationID param.Field[string] `json:"access_application_id" api:"required" format:"uuid"`
	// Identifies the unique resource.
	ResourceID param.Field[string] `json:"resource_id" api:"required"`
	// Enum for base zone-level resource types (those with no extra required fields).
	ResourceType param.Field[ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceType] `json:"resource_type" api:"required"`
	// Contains key-value pairs of tags.
	Tags param.Field[map[string]string] `json:"tags"`
}

func (r ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicy) implementsZoneTagUpdateParamsBodyUnion() {
}

// Enum for base zone-level resource types (those with no extra required fields).
type ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceType string

const (
	ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeAPIGatewayOperation      ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceType = "api_gateway_operation"
	ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeCustomCertificate        ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceType = "custom_certificate"
	ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeCustomHostname           ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceType = "custom_hostname"
	ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeDNSRecord                ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceType = "dns_record"
	ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeManagedClientCertificate ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceType = "managed_client_certificate"
	ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeZone                     ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceType = "zone"
	ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeAccessApplicationPolicy  ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceType = "access_application_policy"
)

func (r ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeAPIGatewayOperation, ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeCustomCertificate, ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeCustomHostname, ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeDNSRecord, ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeManagedClientCertificate, ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeZone, ZoneTagUpdateParamsBodyResourceTaggingSetTagsRequestZoneLevelAccessApplicationPolicyResourceTypeAccessApplicationPolicy:
		return true
	}
	return false
}

// Enum for base zone-level resource types (those with no extra required fields).
type ZoneTagUpdateParamsBodyResourceType string

const (
	ZoneTagUpdateParamsBodyResourceTypeAPIGatewayOperation      ZoneTagUpdateParamsBodyResourceType = "api_gateway_operation"
	ZoneTagUpdateParamsBodyResourceTypeCustomCertificate        ZoneTagUpdateParamsBodyResourceType = "custom_certificate"
	ZoneTagUpdateParamsBodyResourceTypeCustomHostname           ZoneTagUpdateParamsBodyResourceType = "custom_hostname"
	ZoneTagUpdateParamsBodyResourceTypeDNSRecord                ZoneTagUpdateParamsBodyResourceType = "dns_record"
	ZoneTagUpdateParamsBodyResourceTypeManagedClientCertificate ZoneTagUpdateParamsBodyResourceType = "managed_client_certificate"
	ZoneTagUpdateParamsBodyResourceTypeZone                     ZoneTagUpdateParamsBodyResourceType = "zone"
	ZoneTagUpdateParamsBodyResourceTypeAccessApplicationPolicy  ZoneTagUpdateParamsBodyResourceType = "access_application_policy"
)

func (r ZoneTagUpdateParamsBodyResourceType) IsKnown() bool {
	switch r {
	case ZoneTagUpdateParamsBodyResourceTypeAPIGatewayOperation, ZoneTagUpdateParamsBodyResourceTypeCustomCertificate, ZoneTagUpdateParamsBodyResourceTypeCustomHostname, ZoneTagUpdateParamsBodyResourceTypeDNSRecord, ZoneTagUpdateParamsBodyResourceTypeManagedClientCertificate, ZoneTagUpdateParamsBodyResourceTypeZone, ZoneTagUpdateParamsBodyResourceTypeAccessApplicationPolicy:
		return true
	}
	return false
}

type ZoneTagUpdateResponseEnvelope struct {
	Errors   []ZoneTagUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ZoneTagUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success ZoneTagUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	// Response for access_application resources
	Result ZoneTagUpdateResponse             `json:"result"`
	JSON   zoneTagUpdateResponseEnvelopeJSON `json:"-"`
}

// zoneTagUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneTagUpdateResponseEnvelope]
type zoneTagUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTagUpdateResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           ZoneTagUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTagUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTagUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneTagUpdateResponseEnvelopeErrors]
type zoneTagUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTagUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    zoneTagUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTagUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ZoneTagUpdateResponseEnvelopeErrorsSource]
type zoneTagUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTagUpdateResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           ZoneTagUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTagUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTagUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZoneTagUpdateResponseEnvelopeMessages]
type zoneTagUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTagUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    zoneTagUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTagUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneTagUpdateResponseEnvelopeMessagesSource]
type zoneTagUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTagUpdateResponseEnvelopeSuccess bool

const (
	ZoneTagUpdateResponseEnvelopeSuccessTrue ZoneTagUpdateResponseEnvelopeSuccess = true
)

func (r ZoneTagUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTagUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTagDeleteParams struct {
	// Zone ID is required only for zone-level resources
	ZoneID  param.Field[string] `path:"zone_id" api:"required"`
	IfMatch param.Field[string] `header:"If-Match"`
}

type ZoneTagGetParams struct {
	// Zone ID is required only for zone-level resources
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// The ID of the resource to retrieve tags for.
	ResourceID param.Field[string] `query:"resource_id" api:"required"`
	// The type of the resource.
	ResourceType param.Field[ZoneTagGetParamsResourceType] `query:"resource_type" api:"required"`
	// Access application ID identifier. Required for access_application_policy
	// resources.
	AccessApplicationID param.Field[string] `query:"access_application_id" format:"uuid"`
}

// URLQuery serializes [ZoneTagGetParams]'s query parameters as `url.Values`.
func (r ZoneTagGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The type of the resource.
type ZoneTagGetParamsResourceType string

const (
	ZoneTagGetParamsResourceTypeAccessApplicationPolicy  ZoneTagGetParamsResourceType = "access_application_policy"
	ZoneTagGetParamsResourceTypeAPIGatewayOperation      ZoneTagGetParamsResourceType = "api_gateway_operation"
	ZoneTagGetParamsResourceTypeCustomCertificate        ZoneTagGetParamsResourceType = "custom_certificate"
	ZoneTagGetParamsResourceTypeCustomHostname           ZoneTagGetParamsResourceType = "custom_hostname"
	ZoneTagGetParamsResourceTypeDNSRecord                ZoneTagGetParamsResourceType = "dns_record"
	ZoneTagGetParamsResourceTypeManagedClientCertificate ZoneTagGetParamsResourceType = "managed_client_certificate"
	ZoneTagGetParamsResourceTypeZone                     ZoneTagGetParamsResourceType = "zone"
)

func (r ZoneTagGetParamsResourceType) IsKnown() bool {
	switch r {
	case ZoneTagGetParamsResourceTypeAccessApplicationPolicy, ZoneTagGetParamsResourceTypeAPIGatewayOperation, ZoneTagGetParamsResourceTypeCustomCertificate, ZoneTagGetParamsResourceTypeCustomHostname, ZoneTagGetParamsResourceTypeDNSRecord, ZoneTagGetParamsResourceTypeManagedClientCertificate, ZoneTagGetParamsResourceTypeZone:
		return true
	}
	return false
}

type ZoneTagGetResponseEnvelope struct {
	Errors   []ZoneTagGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ZoneTagGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success ZoneTagGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// Response for access_application resources
	Result ZoneTagGetResponse             `json:"result"`
	JSON   zoneTagGetResponseEnvelopeJSON `json:"-"`
}

// zoneTagGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneTagGetResponseEnvelope]
type zoneTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTagGetResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code" api:"required"`
	Message          string                                 `json:"message" api:"required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           ZoneTagGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTagGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTagGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ZoneTagGetResponseEnvelopeErrors]
type zoneTagGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTagGetResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    zoneTagGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTagGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ZoneTagGetResponseEnvelopeErrorsSource]
type zoneTagGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTagGetResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code" api:"required"`
	Message          string                                   `json:"message" api:"required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ZoneTagGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTagGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTagGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ZoneTagGetResponseEnvelopeMessages]
type zoneTagGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTagGetResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    zoneTagGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTagGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [ZoneTagGetResponseEnvelopeMessagesSource]
type zoneTagGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTagGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTagGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTagGetResponseEnvelopeSuccess bool

const (
	ZoneTagGetResponseEnvelopeSuccessTrue ZoneTagGetResponseEnvelopeSuccess = true
)

func (r ZoneTagGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTagGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
