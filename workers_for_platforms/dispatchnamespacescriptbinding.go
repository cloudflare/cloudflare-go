// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/tidwall/gjson"
)

// DispatchNamespaceScriptBindingService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptBindingService] method instead.
type DispatchNamespaceScriptBindingService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptBindingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptBindingService(opts ...option.RequestOption) (r *DispatchNamespaceScriptBindingService) {
	r = &DispatchNamespaceScriptBindingService{}
	r.Options = opts
	return
}

// Fetch script bindings from a script uploaded to a Workers for Platforms
// namespace.
func (r *DispatchNamespaceScriptBindingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptBindingGetParams, opts ...option.RequestOption) (res *[]DispatchNamespaceScriptBindingGetResponse, err error) {
	var env DispatchNamespaceScriptBindingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/bindings", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A binding to allow the Worker to communicate with resources
type DispatchNamespaceScriptBindingGetResponse struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID string `json:"id"`
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The dataset name to bind to.
	Dataset string `json:"dataset"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// JSON data to use.
	Json string `json:"json"`
	// Namespace to bind to.
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// This field can have the runtime type of
	// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound].
	Outbound interface{} `json:"outbound"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string `json:"script_name"`
	// Name of Worker to bind to.
	Service string `json:"service"`
	// The text value to use.
	Text  string                                        `json:"text"`
	JSON  dispatchNamespaceScriptBindingGetResponseJSON `json:"-"`
	union DispatchNamespaceScriptBindingGetResponseUnion
}

// dispatchNamespaceScriptBindingGetResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptBindingGetResponse]
type dispatchNamespaceScriptBindingGetResponseJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	ID            apijson.Field
	BucketName    apijson.Field
	CertificateID apijson.Field
	ClassName     apijson.Field
	Dataset       apijson.Field
	Environment   apijson.Field
	IndexName     apijson.Field
	Json          apijson.Field
	Namespace     apijson.Field
	NamespaceID   apijson.Field
	Outbound      apijson.Field
	QueueName     apijson.Field
	ScriptName    apijson.Field
	Service       apijson.Field
	Text          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dispatchNamespaceScriptBindingGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptBindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptBindingGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptBindingGetResponseUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAny],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDo],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCERT],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecret],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata].
func (r DispatchNamespaceScriptBindingGetResponse) AsUnion() DispatchNamespaceScriptBindingGetResponseUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAny],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDo],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCERT],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecret],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize]
// or
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata].
type DispatchNamespaceScriptBindingGetResponseUnion interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptBindingGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAny{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDo{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCERT{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecret{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
		},
	)
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAny struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type        string                                                             `json:"type,required"`
	ExtraFields map[string]interface{}                                             `json:"-,extras"`
	JSON        dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnyJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnyJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAny]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnyJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAny) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnyJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAny) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeAI DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "ai"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine struct {
	// The dataset name to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeAssets DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "assets"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeBrowserRendering DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "browser_rendering"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeBrowserRendering:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1JSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1JSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeD1 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "d1"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeDispatchNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                         `json:"service"`
	JSON    dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorker]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDo struct {
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDoType `json:"type,required"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                            `json:"script_name"`
	JSON       dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDoJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDoJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDo]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDoJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDoJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDo) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDoType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDoTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDoType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDoType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDoTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeHyperdrive DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeJson DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "json"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeKVNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCERT struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCERTType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindMtlscertJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindMtlscertJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCERT]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindMtlscertJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCERT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindMtlscertJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCERT) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCERTType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCERTTypeMTLSCertificate DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCERTType = "mtls_certificate"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCERTType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCERTTypeMTLSCertificate:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypePlainText DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "plain_text"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeQueue DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "queue"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2 struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Type `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2JSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2JSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2JSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2JSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Type string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2TypeR2Bucket DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Type = "r2_bucket"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2TypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The secret value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecret]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecret) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTypeSecretText DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretType = "secret_text"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeService DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "service"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeTailConsumer DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeVectorize DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "vectorize"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeVersionMetadata DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptBindingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptBindingGetResponseEnvelopeSuccess `json:"success,required"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Result []DispatchNamespaceScriptBindingGetResponse           `json:"result"`
	JSON   dispatchNamespaceScriptBindingGetResponseEnvelopeJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptBindingGetResponseEnvelope]
type dispatchNamespaceScriptBindingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptBindingGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptBindingGetResponseEnvelopeSuccessTrue DispatchNamespaceScriptBindingGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptBindingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
