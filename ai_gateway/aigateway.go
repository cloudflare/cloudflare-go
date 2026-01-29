// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/tidwall/gjson"
)

// AIGatewayService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIGatewayService] method instead.
type AIGatewayService struct {
	Options         []option.RequestOption
	EvaluationTypes *EvaluationTypeService
	Logs            *LogService
	Datasets        *DatasetService
	Evaluations     *EvaluationService
	DynamicRouting  *DynamicRoutingService
	ProviderConfigs *ProviderConfigService
	URLs            *URLService
}

// NewAIGatewayService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIGatewayService(opts ...option.RequestOption) (r *AIGatewayService) {
	r = &AIGatewayService{}
	r.Options = opts
	r.EvaluationTypes = NewEvaluationTypeService(opts...)
	r.Logs = NewLogService(opts...)
	r.Datasets = NewDatasetService(opts...)
	r.Evaluations = NewEvaluationService(opts...)
	r.DynamicRouting = NewDynamicRoutingService(opts...)
	r.ProviderConfigs = NewProviderConfigService(opts...)
	r.URLs = NewURLService(opts...)
	return
}

// Create a new Gateway
func (r *AIGatewayService) New(ctx context.Context, params AIGatewayNewParams, opts ...option.RequestOption) (res *AIGatewayNewResponse, err error) {
	var env AIGatewayNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Gateway
func (r *AIGatewayService) Update(ctx context.Context, id string, params AIGatewayUpdateParams, opts ...option.RequestOption) (res *AIGatewayUpdateResponse, err error) {
	var env AIGatewayUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Gateways
func (r *AIGatewayService) List(ctx context.Context, params AIGatewayListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AIGatewayListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways", params.AccountID)
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

// List Gateways
func (r *AIGatewayService) ListAutoPaging(ctx context.Context, params AIGatewayListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AIGatewayListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a Gateway
func (r *AIGatewayService) Delete(ctx context.Context, id string, body AIGatewayDeleteParams, opts ...option.RequestOption) (res *AIGatewayDeleteResponse, err error) {
	var env AIGatewayDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a Gateway
func (r *AIGatewayService) Get(ctx context.Context, id string, query AIGatewayGetParams, opts ...option.RequestOption) (res *AIGatewayGetResponse, err error) {
	var env AIGatewayGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIGatewayNewResponse struct {
	// gateway id
	ID                      string                                    `json:"id,required"`
	AccountID               string                                    `json:"account_id,required"`
	AccountTag              string                                    `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                      `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                                     `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                      `json:"collect_logs,required"`
	CreatedAt               time.Time                                 `json:"created_at,required" format:"date-time"`
	InternalID              string                                    `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                 `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                     `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                     `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AIGatewayNewResponseRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                      `json:"authentication"`
	DLP                     AIGatewayNewResponseDLP                   `json:"dlp"`
	IsDefault               bool                                      `json:"is_default"`
	LogManagement           int64                                     `json:"log_management,nullable"`
	LogManagementStrategy   AIGatewayNewResponseLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                      `json:"logpush"`
	LogpushPublicKey        string                                    `json:"logpush_public_key,nullable"`
	Otel                    []AIGatewayNewResponseOtel                `json:"otel,nullable"`
	StoreID                 string                                    `json:"store_id,nullable"`
	Stripe                  AIGatewayNewResponseStripe                `json:"stripe,nullable"`
	Zdr                     bool                                      `json:"zdr"`
	JSON                    aiGatewayNewResponseJSON                  `json:"-"`
}

// aiGatewayNewResponseJSON contains the JSON metadata for the struct
// [AIGatewayNewResponse]
type aiGatewayNewResponseJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	DLP                     apijson.Field
	IsDefault               apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	Otel                    apijson.Field
	StoreID                 apijson.Field
	Stripe                  apijson.Field
	Zdr                     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayNewResponseRateLimitingTechnique string

const (
	AIGatewayNewResponseRateLimitingTechniqueFixed   AIGatewayNewResponseRateLimitingTechnique = "fixed"
	AIGatewayNewResponseRateLimitingTechniqueSliding AIGatewayNewResponseRateLimitingTechnique = "sliding"
)

func (r AIGatewayNewResponseRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseRateLimitingTechniqueFixed, AIGatewayNewResponseRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayNewResponseDLP struct {
	Enabled bool                          `json:"enabled,required"`
	Action  AIGatewayNewResponseDLPAction `json:"action"`
	// This field can have the runtime type of [[]AIGatewayNewResponseDLPObjectPolicy].
	Policies interface{} `json:"policies"`
	// This field can have the runtime type of [[]string].
	Profiles interface{}                 `json:"profiles"`
	JSON     aiGatewayNewResponseDLPJSON `json:"-"`
	union    AIGatewayNewResponseDLPUnion
}

// aiGatewayNewResponseDLPJSON contains the JSON metadata for the struct
// [AIGatewayNewResponseDLP]
type aiGatewayNewResponseDLPJSON struct {
	Enabled     apijson.Field
	Action      apijson.Field
	Policies    apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r aiGatewayNewResponseDLPJSON) RawJSON() string {
	return r.raw
}

func (r *AIGatewayNewResponseDLP) UnmarshalJSON(data []byte) (err error) {
	*r = AIGatewayNewResponseDLP{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AIGatewayNewResponseDLPUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [AIGatewayNewResponseDLPObject],
// [AIGatewayNewResponseDLPObject].
func (r AIGatewayNewResponseDLP) AsUnion() AIGatewayNewResponseDLPUnion {
	return r.union
}

// Union satisfied by [AIGatewayNewResponseDLPObject] or
// [AIGatewayNewResponseDLPObject].
type AIGatewayNewResponseDLPUnion interface {
	implementsAIGatewayNewResponseDLP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIGatewayNewResponseDLPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIGatewayNewResponseDLPObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIGatewayNewResponseDLPObject{}),
		},
	)
}

type AIGatewayNewResponseDLPObject struct {
	Action   AIGatewayNewResponseDLPObjectAction `json:"action,required"`
	Enabled  bool                                `json:"enabled,required"`
	Profiles []string                            `json:"profiles,required"`
	JSON     aiGatewayNewResponseDLPObjectJSON   `json:"-"`
}

// aiGatewayNewResponseDLPObjectJSON contains the JSON metadata for the struct
// [AIGatewayNewResponseDLPObject]
type aiGatewayNewResponseDLPObjectJSON struct {
	Action      apijson.Field
	Enabled     apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayNewResponseDLPObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseDLPObjectJSON) RawJSON() string {
	return r.raw
}

func (r AIGatewayNewResponseDLPObject) implementsAIGatewayNewResponseDLP() {}

type AIGatewayNewResponseDLPObjectAction string

const (
	AIGatewayNewResponseDLPObjectActionBlock AIGatewayNewResponseDLPObjectAction = "BLOCK"
	AIGatewayNewResponseDLPObjectActionFlag  AIGatewayNewResponseDLPObjectAction = "FLAG"
)

func (r AIGatewayNewResponseDLPObjectAction) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseDLPObjectActionBlock, AIGatewayNewResponseDLPObjectActionFlag:
		return true
	}
	return false
}

type AIGatewayNewResponseDLPAction string

const (
	AIGatewayNewResponseDLPActionBlock AIGatewayNewResponseDLPAction = "BLOCK"
	AIGatewayNewResponseDLPActionFlag  AIGatewayNewResponseDLPAction = "FLAG"
)

func (r AIGatewayNewResponseDLPAction) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseDLPActionBlock, AIGatewayNewResponseDLPActionFlag:
		return true
	}
	return false
}

type AIGatewayNewResponseLogManagementStrategy string

const (
	AIGatewayNewResponseLogManagementStrategyStopInserting AIGatewayNewResponseLogManagementStrategy = "STOP_INSERTING"
	AIGatewayNewResponseLogManagementStrategyDeleteOldest  AIGatewayNewResponseLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayNewResponseLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseLogManagementStrategyStopInserting, AIGatewayNewResponseLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayNewResponseOtel struct {
	Authorization string                       `json:"authorization,required"`
	Headers       map[string]string            `json:"headers,required"`
	URL           string                       `json:"url,required"`
	JSON          aiGatewayNewResponseOtelJSON `json:"-"`
}

// aiGatewayNewResponseOtelJSON contains the JSON metadata for the struct
// [AIGatewayNewResponseOtel]
type aiGatewayNewResponseOtelJSON struct {
	Authorization apijson.Field
	Headers       apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayNewResponseOtel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseOtelJSON) RawJSON() string {
	return r.raw
}

type AIGatewayNewResponseStripe struct {
	Authorization string                                 `json:"authorization,required"`
	UsageEvents   []AIGatewayNewResponseStripeUsageEvent `json:"usage_events,required"`
	JSON          aiGatewayNewResponseStripeJSON         `json:"-"`
}

// aiGatewayNewResponseStripeJSON contains the JSON metadata for the struct
// [AIGatewayNewResponseStripe]
type aiGatewayNewResponseStripeJSON struct {
	Authorization apijson.Field
	UsageEvents   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayNewResponseStripe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseStripeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayNewResponseStripeUsageEvent struct {
	Payload string                                   `json:"payload,required"`
	JSON    aiGatewayNewResponseStripeUsageEventJSON `json:"-"`
}

// aiGatewayNewResponseStripeUsageEventJSON contains the JSON metadata for the
// struct [AIGatewayNewResponseStripeUsageEvent]
type aiGatewayNewResponseStripeUsageEventJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayNewResponseStripeUsageEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseStripeUsageEventJSON) RawJSON() string {
	return r.raw
}

type AIGatewayUpdateResponse struct {
	// gateway id
	ID                      string                                       `json:"id,required"`
	AccountID               string                                       `json:"account_id,required"`
	AccountTag              string                                       `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                         `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                                        `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                         `json:"collect_logs,required"`
	CreatedAt               time.Time                                    `json:"created_at,required" format:"date-time"`
	InternalID              string                                       `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                    `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                        `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                        `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AIGatewayUpdateResponseRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                         `json:"authentication"`
	DLP                     AIGatewayUpdateResponseDLP                   `json:"dlp"`
	IsDefault               bool                                         `json:"is_default"`
	LogManagement           int64                                        `json:"log_management,nullable"`
	LogManagementStrategy   AIGatewayUpdateResponseLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                         `json:"logpush"`
	LogpushPublicKey        string                                       `json:"logpush_public_key,nullable"`
	Otel                    []AIGatewayUpdateResponseOtel                `json:"otel,nullable"`
	StoreID                 string                                       `json:"store_id,nullable"`
	Stripe                  AIGatewayUpdateResponseStripe                `json:"stripe,nullable"`
	Zdr                     bool                                         `json:"zdr"`
	JSON                    aiGatewayUpdateResponseJSON                  `json:"-"`
}

// aiGatewayUpdateResponseJSON contains the JSON metadata for the struct
// [AIGatewayUpdateResponse]
type aiGatewayUpdateResponseJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	DLP                     apijson.Field
	IsDefault               apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	Otel                    apijson.Field
	StoreID                 apijson.Field
	Stripe                  apijson.Field
	Zdr                     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayUpdateResponseRateLimitingTechnique string

const (
	AIGatewayUpdateResponseRateLimitingTechniqueFixed   AIGatewayUpdateResponseRateLimitingTechnique = "fixed"
	AIGatewayUpdateResponseRateLimitingTechniqueSliding AIGatewayUpdateResponseRateLimitingTechnique = "sliding"
)

func (r AIGatewayUpdateResponseRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseRateLimitingTechniqueFixed, AIGatewayUpdateResponseRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayUpdateResponseDLP struct {
	Enabled bool                             `json:"enabled,required"`
	Action  AIGatewayUpdateResponseDLPAction `json:"action"`
	// This field can have the runtime type of
	// [[]AIGatewayUpdateResponseDLPObjectPolicy].
	Policies interface{} `json:"policies"`
	// This field can have the runtime type of [[]string].
	Profiles interface{}                    `json:"profiles"`
	JSON     aiGatewayUpdateResponseDLPJSON `json:"-"`
	union    AIGatewayUpdateResponseDLPUnion
}

// aiGatewayUpdateResponseDLPJSON contains the JSON metadata for the struct
// [AIGatewayUpdateResponseDLP]
type aiGatewayUpdateResponseDLPJSON struct {
	Enabled     apijson.Field
	Action      apijson.Field
	Policies    apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r aiGatewayUpdateResponseDLPJSON) RawJSON() string {
	return r.raw
}

func (r *AIGatewayUpdateResponseDLP) UnmarshalJSON(data []byte) (err error) {
	*r = AIGatewayUpdateResponseDLP{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AIGatewayUpdateResponseDLPUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [AIGatewayUpdateResponseDLPObject],
// [AIGatewayUpdateResponseDLPObject].
func (r AIGatewayUpdateResponseDLP) AsUnion() AIGatewayUpdateResponseDLPUnion {
	return r.union
}

// Union satisfied by [AIGatewayUpdateResponseDLPObject] or
// [AIGatewayUpdateResponseDLPObject].
type AIGatewayUpdateResponseDLPUnion interface {
	implementsAIGatewayUpdateResponseDLP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIGatewayUpdateResponseDLPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIGatewayUpdateResponseDLPObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIGatewayUpdateResponseDLPObject{}),
		},
	)
}

type AIGatewayUpdateResponseDLPObject struct {
	Action   AIGatewayUpdateResponseDLPObjectAction `json:"action,required"`
	Enabled  bool                                   `json:"enabled,required"`
	Profiles []string                               `json:"profiles,required"`
	JSON     aiGatewayUpdateResponseDLPObjectJSON   `json:"-"`
}

// aiGatewayUpdateResponseDLPObjectJSON contains the JSON metadata for the struct
// [AIGatewayUpdateResponseDLPObject]
type aiGatewayUpdateResponseDLPObjectJSON struct {
	Action      apijson.Field
	Enabled     apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayUpdateResponseDLPObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseDLPObjectJSON) RawJSON() string {
	return r.raw
}

func (r AIGatewayUpdateResponseDLPObject) implementsAIGatewayUpdateResponseDLP() {}

type AIGatewayUpdateResponseDLPObjectAction string

const (
	AIGatewayUpdateResponseDLPObjectActionBlock AIGatewayUpdateResponseDLPObjectAction = "BLOCK"
	AIGatewayUpdateResponseDLPObjectActionFlag  AIGatewayUpdateResponseDLPObjectAction = "FLAG"
)

func (r AIGatewayUpdateResponseDLPObjectAction) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseDLPObjectActionBlock, AIGatewayUpdateResponseDLPObjectActionFlag:
		return true
	}
	return false
}

type AIGatewayUpdateResponseDLPAction string

const (
	AIGatewayUpdateResponseDLPActionBlock AIGatewayUpdateResponseDLPAction = "BLOCK"
	AIGatewayUpdateResponseDLPActionFlag  AIGatewayUpdateResponseDLPAction = "FLAG"
)

func (r AIGatewayUpdateResponseDLPAction) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseDLPActionBlock, AIGatewayUpdateResponseDLPActionFlag:
		return true
	}
	return false
}

type AIGatewayUpdateResponseLogManagementStrategy string

const (
	AIGatewayUpdateResponseLogManagementStrategyStopInserting AIGatewayUpdateResponseLogManagementStrategy = "STOP_INSERTING"
	AIGatewayUpdateResponseLogManagementStrategyDeleteOldest  AIGatewayUpdateResponseLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayUpdateResponseLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseLogManagementStrategyStopInserting, AIGatewayUpdateResponseLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayUpdateResponseOtel struct {
	Authorization string                          `json:"authorization,required"`
	Headers       map[string]string               `json:"headers,required"`
	URL           string                          `json:"url,required"`
	JSON          aiGatewayUpdateResponseOtelJSON `json:"-"`
}

// aiGatewayUpdateResponseOtelJSON contains the JSON metadata for the struct
// [AIGatewayUpdateResponseOtel]
type aiGatewayUpdateResponseOtelJSON struct {
	Authorization apijson.Field
	Headers       apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayUpdateResponseOtel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseOtelJSON) RawJSON() string {
	return r.raw
}

type AIGatewayUpdateResponseStripe struct {
	Authorization string                                    `json:"authorization,required"`
	UsageEvents   []AIGatewayUpdateResponseStripeUsageEvent `json:"usage_events,required"`
	JSON          aiGatewayUpdateResponseStripeJSON         `json:"-"`
}

// aiGatewayUpdateResponseStripeJSON contains the JSON metadata for the struct
// [AIGatewayUpdateResponseStripe]
type aiGatewayUpdateResponseStripeJSON struct {
	Authorization apijson.Field
	UsageEvents   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayUpdateResponseStripe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseStripeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayUpdateResponseStripeUsageEvent struct {
	Payload string                                      `json:"payload,required"`
	JSON    aiGatewayUpdateResponseStripeUsageEventJSON `json:"-"`
}

// aiGatewayUpdateResponseStripeUsageEventJSON contains the JSON metadata for the
// struct [AIGatewayUpdateResponseStripeUsageEvent]
type aiGatewayUpdateResponseStripeUsageEventJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayUpdateResponseStripeUsageEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseStripeUsageEventJSON) RawJSON() string {
	return r.raw
}

type AIGatewayListResponse struct {
	// gateway id
	ID                      string                                     `json:"id,required"`
	AccountID               string                                     `json:"account_id,required"`
	AccountTag              string                                     `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                       `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                                      `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                       `json:"collect_logs,required"`
	CreatedAt               time.Time                                  `json:"created_at,required" format:"date-time"`
	InternalID              string                                     `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                  `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                      `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                      `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AIGatewayListResponseRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                       `json:"authentication"`
	DLP                     AIGatewayListResponseDLP                   `json:"dlp"`
	IsDefault               bool                                       `json:"is_default"`
	LogManagement           int64                                      `json:"log_management,nullable"`
	LogManagementStrategy   AIGatewayListResponseLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                       `json:"logpush"`
	LogpushPublicKey        string                                     `json:"logpush_public_key,nullable"`
	Otel                    []AIGatewayListResponseOtel                `json:"otel,nullable"`
	StoreID                 string                                     `json:"store_id,nullable"`
	Stripe                  AIGatewayListResponseStripe                `json:"stripe,nullable"`
	Zdr                     bool                                       `json:"zdr"`
	JSON                    aiGatewayListResponseJSON                  `json:"-"`
}

// aiGatewayListResponseJSON contains the JSON metadata for the struct
// [AIGatewayListResponse]
type aiGatewayListResponseJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	DLP                     apijson.Field
	IsDefault               apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	Otel                    apijson.Field
	StoreID                 apijson.Field
	Stripe                  apijson.Field
	Zdr                     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayListResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayListResponseRateLimitingTechnique string

const (
	AIGatewayListResponseRateLimitingTechniqueFixed   AIGatewayListResponseRateLimitingTechnique = "fixed"
	AIGatewayListResponseRateLimitingTechniqueSliding AIGatewayListResponseRateLimitingTechnique = "sliding"
)

func (r AIGatewayListResponseRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayListResponseRateLimitingTechniqueFixed, AIGatewayListResponseRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayListResponseDLP struct {
	Enabled bool                           `json:"enabled,required"`
	Action  AIGatewayListResponseDLPAction `json:"action"`
	// This field can have the runtime type of
	// [[]AIGatewayListResponseDLPObjectPolicy].
	Policies interface{} `json:"policies"`
	// This field can have the runtime type of [[]string].
	Profiles interface{}                  `json:"profiles"`
	JSON     aiGatewayListResponseDLPJSON `json:"-"`
	union    AIGatewayListResponseDLPUnion
}

// aiGatewayListResponseDLPJSON contains the JSON metadata for the struct
// [AIGatewayListResponseDLP]
type aiGatewayListResponseDLPJSON struct {
	Enabled     apijson.Field
	Action      apijson.Field
	Policies    apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r aiGatewayListResponseDLPJSON) RawJSON() string {
	return r.raw
}

func (r *AIGatewayListResponseDLP) UnmarshalJSON(data []byte) (err error) {
	*r = AIGatewayListResponseDLP{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AIGatewayListResponseDLPUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [AIGatewayListResponseDLPObject],
// [AIGatewayListResponseDLPObject].
func (r AIGatewayListResponseDLP) AsUnion() AIGatewayListResponseDLPUnion {
	return r.union
}

// Union satisfied by [AIGatewayListResponseDLPObject] or
// [AIGatewayListResponseDLPObject].
type AIGatewayListResponseDLPUnion interface {
	implementsAIGatewayListResponseDLP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIGatewayListResponseDLPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIGatewayListResponseDLPObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIGatewayListResponseDLPObject{}),
		},
	)
}

type AIGatewayListResponseDLPObject struct {
	Action   AIGatewayListResponseDLPObjectAction `json:"action,required"`
	Enabled  bool                                 `json:"enabled,required"`
	Profiles []string                             `json:"profiles,required"`
	JSON     aiGatewayListResponseDLPObjectJSON   `json:"-"`
}

// aiGatewayListResponseDLPObjectJSON contains the JSON metadata for the struct
// [AIGatewayListResponseDLPObject]
type aiGatewayListResponseDLPObjectJSON struct {
	Action      apijson.Field
	Enabled     apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayListResponseDLPObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayListResponseDLPObjectJSON) RawJSON() string {
	return r.raw
}

func (r AIGatewayListResponseDLPObject) implementsAIGatewayListResponseDLP() {}

type AIGatewayListResponseDLPObjectAction string

const (
	AIGatewayListResponseDLPObjectActionBlock AIGatewayListResponseDLPObjectAction = "BLOCK"
	AIGatewayListResponseDLPObjectActionFlag  AIGatewayListResponseDLPObjectAction = "FLAG"
)

func (r AIGatewayListResponseDLPObjectAction) IsKnown() bool {
	switch r {
	case AIGatewayListResponseDLPObjectActionBlock, AIGatewayListResponseDLPObjectActionFlag:
		return true
	}
	return false
}

type AIGatewayListResponseDLPAction string

const (
	AIGatewayListResponseDLPActionBlock AIGatewayListResponseDLPAction = "BLOCK"
	AIGatewayListResponseDLPActionFlag  AIGatewayListResponseDLPAction = "FLAG"
)

func (r AIGatewayListResponseDLPAction) IsKnown() bool {
	switch r {
	case AIGatewayListResponseDLPActionBlock, AIGatewayListResponseDLPActionFlag:
		return true
	}
	return false
}

type AIGatewayListResponseLogManagementStrategy string

const (
	AIGatewayListResponseLogManagementStrategyStopInserting AIGatewayListResponseLogManagementStrategy = "STOP_INSERTING"
	AIGatewayListResponseLogManagementStrategyDeleteOldest  AIGatewayListResponseLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayListResponseLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayListResponseLogManagementStrategyStopInserting, AIGatewayListResponseLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayListResponseOtel struct {
	Authorization string                        `json:"authorization,required"`
	Headers       map[string]string             `json:"headers,required"`
	URL           string                        `json:"url,required"`
	JSON          aiGatewayListResponseOtelJSON `json:"-"`
}

// aiGatewayListResponseOtelJSON contains the JSON metadata for the struct
// [AIGatewayListResponseOtel]
type aiGatewayListResponseOtelJSON struct {
	Authorization apijson.Field
	Headers       apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayListResponseOtel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayListResponseOtelJSON) RawJSON() string {
	return r.raw
}

type AIGatewayListResponseStripe struct {
	Authorization string                                  `json:"authorization,required"`
	UsageEvents   []AIGatewayListResponseStripeUsageEvent `json:"usage_events,required"`
	JSON          aiGatewayListResponseStripeJSON         `json:"-"`
}

// aiGatewayListResponseStripeJSON contains the JSON metadata for the struct
// [AIGatewayListResponseStripe]
type aiGatewayListResponseStripeJSON struct {
	Authorization apijson.Field
	UsageEvents   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayListResponseStripe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayListResponseStripeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayListResponseStripeUsageEvent struct {
	Payload string                                    `json:"payload,required"`
	JSON    aiGatewayListResponseStripeUsageEventJSON `json:"-"`
}

// aiGatewayListResponseStripeUsageEventJSON contains the JSON metadata for the
// struct [AIGatewayListResponseStripeUsageEvent]
type aiGatewayListResponseStripeUsageEventJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayListResponseStripeUsageEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayListResponseStripeUsageEventJSON) RawJSON() string {
	return r.raw
}

type AIGatewayDeleteResponse struct {
	// gateway id
	ID                      string                                       `json:"id,required"`
	AccountID               string                                       `json:"account_id,required"`
	AccountTag              string                                       `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                         `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                                        `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                         `json:"collect_logs,required"`
	CreatedAt               time.Time                                    `json:"created_at,required" format:"date-time"`
	InternalID              string                                       `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                    `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                        `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                        `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AIGatewayDeleteResponseRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                         `json:"authentication"`
	DLP                     AIGatewayDeleteResponseDLP                   `json:"dlp"`
	IsDefault               bool                                         `json:"is_default"`
	LogManagement           int64                                        `json:"log_management,nullable"`
	LogManagementStrategy   AIGatewayDeleteResponseLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                         `json:"logpush"`
	LogpushPublicKey        string                                       `json:"logpush_public_key,nullable"`
	Otel                    []AIGatewayDeleteResponseOtel                `json:"otel,nullable"`
	StoreID                 string                                       `json:"store_id,nullable"`
	Stripe                  AIGatewayDeleteResponseStripe                `json:"stripe,nullable"`
	Zdr                     bool                                         `json:"zdr"`
	JSON                    aiGatewayDeleteResponseJSON                  `json:"-"`
}

// aiGatewayDeleteResponseJSON contains the JSON metadata for the struct
// [AIGatewayDeleteResponse]
type aiGatewayDeleteResponseJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	DLP                     apijson.Field
	IsDefault               apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	Otel                    apijson.Field
	StoreID                 apijson.Field
	Stripe                  apijson.Field
	Zdr                     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayDeleteResponseRateLimitingTechnique string

const (
	AIGatewayDeleteResponseRateLimitingTechniqueFixed   AIGatewayDeleteResponseRateLimitingTechnique = "fixed"
	AIGatewayDeleteResponseRateLimitingTechniqueSliding AIGatewayDeleteResponseRateLimitingTechnique = "sliding"
)

func (r AIGatewayDeleteResponseRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseRateLimitingTechniqueFixed, AIGatewayDeleteResponseRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayDeleteResponseDLP struct {
	Enabled bool                             `json:"enabled,required"`
	Action  AIGatewayDeleteResponseDLPAction `json:"action"`
	// This field can have the runtime type of
	// [[]AIGatewayDeleteResponseDLPObjectPolicy].
	Policies interface{} `json:"policies"`
	// This field can have the runtime type of [[]string].
	Profiles interface{}                    `json:"profiles"`
	JSON     aiGatewayDeleteResponseDLPJSON `json:"-"`
	union    AIGatewayDeleteResponseDLPUnion
}

// aiGatewayDeleteResponseDLPJSON contains the JSON metadata for the struct
// [AIGatewayDeleteResponseDLP]
type aiGatewayDeleteResponseDLPJSON struct {
	Enabled     apijson.Field
	Action      apijson.Field
	Policies    apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r aiGatewayDeleteResponseDLPJSON) RawJSON() string {
	return r.raw
}

func (r *AIGatewayDeleteResponseDLP) UnmarshalJSON(data []byte) (err error) {
	*r = AIGatewayDeleteResponseDLP{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AIGatewayDeleteResponseDLPUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [AIGatewayDeleteResponseDLPObject],
// [AIGatewayDeleteResponseDLPObject].
func (r AIGatewayDeleteResponseDLP) AsUnion() AIGatewayDeleteResponseDLPUnion {
	return r.union
}

// Union satisfied by [AIGatewayDeleteResponseDLPObject] or
// [AIGatewayDeleteResponseDLPObject].
type AIGatewayDeleteResponseDLPUnion interface {
	implementsAIGatewayDeleteResponseDLP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIGatewayDeleteResponseDLPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIGatewayDeleteResponseDLPObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIGatewayDeleteResponseDLPObject{}),
		},
	)
}

type AIGatewayDeleteResponseDLPObject struct {
	Action   AIGatewayDeleteResponseDLPObjectAction `json:"action,required"`
	Enabled  bool                                   `json:"enabled,required"`
	Profiles []string                               `json:"profiles,required"`
	JSON     aiGatewayDeleteResponseDLPObjectJSON   `json:"-"`
}

// aiGatewayDeleteResponseDLPObjectJSON contains the JSON metadata for the struct
// [AIGatewayDeleteResponseDLPObject]
type aiGatewayDeleteResponseDLPObjectJSON struct {
	Action      apijson.Field
	Enabled     apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayDeleteResponseDLPObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseDLPObjectJSON) RawJSON() string {
	return r.raw
}

func (r AIGatewayDeleteResponseDLPObject) implementsAIGatewayDeleteResponseDLP() {}

type AIGatewayDeleteResponseDLPObjectAction string

const (
	AIGatewayDeleteResponseDLPObjectActionBlock AIGatewayDeleteResponseDLPObjectAction = "BLOCK"
	AIGatewayDeleteResponseDLPObjectActionFlag  AIGatewayDeleteResponseDLPObjectAction = "FLAG"
)

func (r AIGatewayDeleteResponseDLPObjectAction) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseDLPObjectActionBlock, AIGatewayDeleteResponseDLPObjectActionFlag:
		return true
	}
	return false
}

type AIGatewayDeleteResponseDLPAction string

const (
	AIGatewayDeleteResponseDLPActionBlock AIGatewayDeleteResponseDLPAction = "BLOCK"
	AIGatewayDeleteResponseDLPActionFlag  AIGatewayDeleteResponseDLPAction = "FLAG"
)

func (r AIGatewayDeleteResponseDLPAction) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseDLPActionBlock, AIGatewayDeleteResponseDLPActionFlag:
		return true
	}
	return false
}

type AIGatewayDeleteResponseLogManagementStrategy string

const (
	AIGatewayDeleteResponseLogManagementStrategyStopInserting AIGatewayDeleteResponseLogManagementStrategy = "STOP_INSERTING"
	AIGatewayDeleteResponseLogManagementStrategyDeleteOldest  AIGatewayDeleteResponseLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayDeleteResponseLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseLogManagementStrategyStopInserting, AIGatewayDeleteResponseLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayDeleteResponseOtel struct {
	Authorization string                          `json:"authorization,required"`
	Headers       map[string]string               `json:"headers,required"`
	URL           string                          `json:"url,required"`
	JSON          aiGatewayDeleteResponseOtelJSON `json:"-"`
}

// aiGatewayDeleteResponseOtelJSON contains the JSON metadata for the struct
// [AIGatewayDeleteResponseOtel]
type aiGatewayDeleteResponseOtelJSON struct {
	Authorization apijson.Field
	Headers       apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayDeleteResponseOtel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseOtelJSON) RawJSON() string {
	return r.raw
}

type AIGatewayDeleteResponseStripe struct {
	Authorization string                                    `json:"authorization,required"`
	UsageEvents   []AIGatewayDeleteResponseStripeUsageEvent `json:"usage_events,required"`
	JSON          aiGatewayDeleteResponseStripeJSON         `json:"-"`
}

// aiGatewayDeleteResponseStripeJSON contains the JSON metadata for the struct
// [AIGatewayDeleteResponseStripe]
type aiGatewayDeleteResponseStripeJSON struct {
	Authorization apijson.Field
	UsageEvents   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayDeleteResponseStripe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseStripeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayDeleteResponseStripeUsageEvent struct {
	Payload string                                      `json:"payload,required"`
	JSON    aiGatewayDeleteResponseStripeUsageEventJSON `json:"-"`
}

// aiGatewayDeleteResponseStripeUsageEventJSON contains the JSON metadata for the
// struct [AIGatewayDeleteResponseStripeUsageEvent]
type aiGatewayDeleteResponseStripeUsageEventJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayDeleteResponseStripeUsageEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseStripeUsageEventJSON) RawJSON() string {
	return r.raw
}

type AIGatewayGetResponse struct {
	// gateway id
	ID                      string                                    `json:"id,required"`
	AccountID               string                                    `json:"account_id,required"`
	AccountTag              string                                    `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                      `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                                     `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                      `json:"collect_logs,required"`
	CreatedAt               time.Time                                 `json:"created_at,required" format:"date-time"`
	InternalID              string                                    `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                 `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                     `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                     `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AIGatewayGetResponseRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                      `json:"authentication"`
	DLP                     AIGatewayGetResponseDLP                   `json:"dlp"`
	IsDefault               bool                                      `json:"is_default"`
	LogManagement           int64                                     `json:"log_management,nullable"`
	LogManagementStrategy   AIGatewayGetResponseLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                      `json:"logpush"`
	LogpushPublicKey        string                                    `json:"logpush_public_key,nullable"`
	Otel                    []AIGatewayGetResponseOtel                `json:"otel,nullable"`
	StoreID                 string                                    `json:"store_id,nullable"`
	Stripe                  AIGatewayGetResponseStripe                `json:"stripe,nullable"`
	Zdr                     bool                                      `json:"zdr"`
	JSON                    aiGatewayGetResponseJSON                  `json:"-"`
}

// aiGatewayGetResponseJSON contains the JSON metadata for the struct
// [AIGatewayGetResponse]
type aiGatewayGetResponseJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	DLP                     apijson.Field
	IsDefault               apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	Otel                    apijson.Field
	StoreID                 apijson.Field
	Stripe                  apijson.Field
	Zdr                     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayGetResponseRateLimitingTechnique string

const (
	AIGatewayGetResponseRateLimitingTechniqueFixed   AIGatewayGetResponseRateLimitingTechnique = "fixed"
	AIGatewayGetResponseRateLimitingTechniqueSliding AIGatewayGetResponseRateLimitingTechnique = "sliding"
)

func (r AIGatewayGetResponseRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseRateLimitingTechniqueFixed, AIGatewayGetResponseRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayGetResponseDLP struct {
	Enabled bool                          `json:"enabled,required"`
	Action  AIGatewayGetResponseDLPAction `json:"action"`
	// This field can have the runtime type of [[]AIGatewayGetResponseDLPObjectPolicy].
	Policies interface{} `json:"policies"`
	// This field can have the runtime type of [[]string].
	Profiles interface{}                 `json:"profiles"`
	JSON     aiGatewayGetResponseDLPJSON `json:"-"`
	union    AIGatewayGetResponseDLPUnion
}

// aiGatewayGetResponseDLPJSON contains the JSON metadata for the struct
// [AIGatewayGetResponseDLP]
type aiGatewayGetResponseDLPJSON struct {
	Enabled     apijson.Field
	Action      apijson.Field
	Policies    apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r aiGatewayGetResponseDLPJSON) RawJSON() string {
	return r.raw
}

func (r *AIGatewayGetResponseDLP) UnmarshalJSON(data []byte) (err error) {
	*r = AIGatewayGetResponseDLP{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AIGatewayGetResponseDLPUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [AIGatewayGetResponseDLPObject],
// [AIGatewayGetResponseDLPObject].
func (r AIGatewayGetResponseDLP) AsUnion() AIGatewayGetResponseDLPUnion {
	return r.union
}

// Union satisfied by [AIGatewayGetResponseDLPObject] or
// [AIGatewayGetResponseDLPObject].
type AIGatewayGetResponseDLPUnion interface {
	implementsAIGatewayGetResponseDLP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIGatewayGetResponseDLPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIGatewayGetResponseDLPObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIGatewayGetResponseDLPObject{}),
		},
	)
}

type AIGatewayGetResponseDLPObject struct {
	Action   AIGatewayGetResponseDLPObjectAction `json:"action,required"`
	Enabled  bool                                `json:"enabled,required"`
	Profiles []string                            `json:"profiles,required"`
	JSON     aiGatewayGetResponseDLPObjectJSON   `json:"-"`
}

// aiGatewayGetResponseDLPObjectJSON contains the JSON metadata for the struct
// [AIGatewayGetResponseDLPObject]
type aiGatewayGetResponseDLPObjectJSON struct {
	Action      apijson.Field
	Enabled     apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayGetResponseDLPObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseDLPObjectJSON) RawJSON() string {
	return r.raw
}

func (r AIGatewayGetResponseDLPObject) implementsAIGatewayGetResponseDLP() {}

type AIGatewayGetResponseDLPObjectAction string

const (
	AIGatewayGetResponseDLPObjectActionBlock AIGatewayGetResponseDLPObjectAction = "BLOCK"
	AIGatewayGetResponseDLPObjectActionFlag  AIGatewayGetResponseDLPObjectAction = "FLAG"
)

func (r AIGatewayGetResponseDLPObjectAction) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseDLPObjectActionBlock, AIGatewayGetResponseDLPObjectActionFlag:
		return true
	}
	return false
}

type AIGatewayGetResponseDLPAction string

const (
	AIGatewayGetResponseDLPActionBlock AIGatewayGetResponseDLPAction = "BLOCK"
	AIGatewayGetResponseDLPActionFlag  AIGatewayGetResponseDLPAction = "FLAG"
)

func (r AIGatewayGetResponseDLPAction) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseDLPActionBlock, AIGatewayGetResponseDLPActionFlag:
		return true
	}
	return false
}

type AIGatewayGetResponseLogManagementStrategy string

const (
	AIGatewayGetResponseLogManagementStrategyStopInserting AIGatewayGetResponseLogManagementStrategy = "STOP_INSERTING"
	AIGatewayGetResponseLogManagementStrategyDeleteOldest  AIGatewayGetResponseLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayGetResponseLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseLogManagementStrategyStopInserting, AIGatewayGetResponseLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayGetResponseOtel struct {
	Authorization string                       `json:"authorization,required"`
	Headers       map[string]string            `json:"headers,required"`
	URL           string                       `json:"url,required"`
	JSON          aiGatewayGetResponseOtelJSON `json:"-"`
}

// aiGatewayGetResponseOtelJSON contains the JSON metadata for the struct
// [AIGatewayGetResponseOtel]
type aiGatewayGetResponseOtelJSON struct {
	Authorization apijson.Field
	Headers       apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayGetResponseOtel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseOtelJSON) RawJSON() string {
	return r.raw
}

type AIGatewayGetResponseStripe struct {
	Authorization string                                 `json:"authorization,required"`
	UsageEvents   []AIGatewayGetResponseStripeUsageEvent `json:"usage_events,required"`
	JSON          aiGatewayGetResponseStripeJSON         `json:"-"`
}

// aiGatewayGetResponseStripeJSON contains the JSON metadata for the struct
// [AIGatewayGetResponseStripe]
type aiGatewayGetResponseStripeJSON struct {
	Authorization apijson.Field
	UsageEvents   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayGetResponseStripe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseStripeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayGetResponseStripeUsageEvent struct {
	Payload string                                   `json:"payload,required"`
	JSON    aiGatewayGetResponseStripeUsageEventJSON `json:"-"`
}

// aiGatewayGetResponseStripeUsageEventJSON contains the JSON metadata for the
// struct [AIGatewayGetResponseStripeUsageEvent]
type aiGatewayGetResponseStripeUsageEventJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayGetResponseStripeUsageEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseStripeUsageEventJSON) RawJSON() string {
	return r.raw
}

type AIGatewayNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// gateway id
	ID                      param.Field[string]                                  `json:"id,required"`
	CacheInvalidateOnUpdate param.Field[bool]                                    `json:"cache_invalidate_on_update,required"`
	CacheTTL                param.Field[int64]                                   `json:"cache_ttl,required"`
	CollectLogs             param.Field[bool]                                    `json:"collect_logs,required"`
	RateLimitingInterval    param.Field[int64]                                   `json:"rate_limiting_interval,required"`
	RateLimitingLimit       param.Field[int64]                                   `json:"rate_limiting_limit,required"`
	RateLimitingTechnique   param.Field[AIGatewayNewParamsRateLimitingTechnique] `json:"rate_limiting_technique,required"`
	Authentication          param.Field[bool]                                    `json:"authentication"`
	IsDefault               param.Field[bool]                                    `json:"is_default"`
	LogManagement           param.Field[int64]                                   `json:"log_management"`
	LogManagementStrategy   param.Field[AIGatewayNewParamsLogManagementStrategy] `json:"log_management_strategy"`
	Logpush                 param.Field[bool]                                    `json:"logpush"`
	LogpushPublicKey        param.Field[string]                                  `json:"logpush_public_key"`
	Zdr                     param.Field[bool]                                    `json:"zdr"`
}

func (r AIGatewayNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayNewParamsRateLimitingTechnique string

const (
	AIGatewayNewParamsRateLimitingTechniqueFixed   AIGatewayNewParamsRateLimitingTechnique = "fixed"
	AIGatewayNewParamsRateLimitingTechniqueSliding AIGatewayNewParamsRateLimitingTechnique = "sliding"
)

func (r AIGatewayNewParamsRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayNewParamsRateLimitingTechniqueFixed, AIGatewayNewParamsRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayNewParamsLogManagementStrategy string

const (
	AIGatewayNewParamsLogManagementStrategyStopInserting AIGatewayNewParamsLogManagementStrategy = "STOP_INSERTING"
	AIGatewayNewParamsLogManagementStrategyDeleteOldest  AIGatewayNewParamsLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayNewParamsLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayNewParamsLogManagementStrategyStopInserting, AIGatewayNewParamsLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayNewResponseEnvelope struct {
	Result  AIGatewayNewResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    aiGatewayNewResponseEnvelopeJSON `json:"-"`
}

// aiGatewayNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIGatewayNewResponseEnvelope]
type aiGatewayNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayUpdateParams struct {
	AccountID               param.Field[string]                                     `path:"account_id,required"`
	CacheInvalidateOnUpdate param.Field[bool]                                       `json:"cache_invalidate_on_update,required"`
	CacheTTL                param.Field[int64]                                      `json:"cache_ttl,required"`
	CollectLogs             param.Field[bool]                                       `json:"collect_logs,required"`
	RateLimitingInterval    param.Field[int64]                                      `json:"rate_limiting_interval,required"`
	RateLimitingLimit       param.Field[int64]                                      `json:"rate_limiting_limit,required"`
	RateLimitingTechnique   param.Field[AIGatewayUpdateParamsRateLimitingTechnique] `json:"rate_limiting_technique,required"`
	Authentication          param.Field[bool]                                       `json:"authentication"`
	DLP                     param.Field[AIGatewayUpdateParamsDLPUnion]              `json:"dlp"`
	IsDefault               param.Field[bool]                                       `json:"is_default"`
	LogManagement           param.Field[int64]                                      `json:"log_management"`
	LogManagementStrategy   param.Field[AIGatewayUpdateParamsLogManagementStrategy] `json:"log_management_strategy"`
	Logpush                 param.Field[bool]                                       `json:"logpush"`
	LogpushPublicKey        param.Field[string]                                     `json:"logpush_public_key"`
	Otel                    param.Field[[]AIGatewayUpdateParamsOtel]                `json:"otel"`
	StoreID                 param.Field[string]                                     `json:"store_id"`
	Stripe                  param.Field[AIGatewayUpdateParamsStripe]                `json:"stripe"`
	Zdr                     param.Field[bool]                                       `json:"zdr"`
}

func (r AIGatewayUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayUpdateParamsRateLimitingTechnique string

const (
	AIGatewayUpdateParamsRateLimitingTechniqueFixed   AIGatewayUpdateParamsRateLimitingTechnique = "fixed"
	AIGatewayUpdateParamsRateLimitingTechniqueSliding AIGatewayUpdateParamsRateLimitingTechnique = "sliding"
)

func (r AIGatewayUpdateParamsRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsRateLimitingTechniqueFixed, AIGatewayUpdateParamsRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayUpdateParamsDLP struct {
	Enabled  param.Field[bool]                           `json:"enabled,required"`
	Action   param.Field[AIGatewayUpdateParamsDLPAction] `json:"action"`
	Policies param.Field[interface{}]                    `json:"policies"`
	Profiles param.Field[interface{}]                    `json:"profiles"`
}

func (r AIGatewayUpdateParamsDLP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIGatewayUpdateParamsDLP) implementsAIGatewayUpdateParamsDLPUnion() {}

// Satisfied by [ai_gateway.AIGatewayUpdateParamsDLPObject],
// [ai_gateway.AIGatewayUpdateParamsDLPObject], [AIGatewayUpdateParamsDLP].
type AIGatewayUpdateParamsDLPUnion interface {
	implementsAIGatewayUpdateParamsDLPUnion()
}

type AIGatewayUpdateParamsDLPObject struct {
	Action   param.Field[AIGatewayUpdateParamsDLPObjectAction] `json:"action,required"`
	Enabled  param.Field[bool]                                 `json:"enabled,required"`
	Profiles param.Field[[]string]                             `json:"profiles,required"`
}

func (r AIGatewayUpdateParamsDLPObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIGatewayUpdateParamsDLPObject) implementsAIGatewayUpdateParamsDLPUnion() {}

type AIGatewayUpdateParamsDLPObjectAction string

const (
	AIGatewayUpdateParamsDLPObjectActionBlock AIGatewayUpdateParamsDLPObjectAction = "BLOCK"
	AIGatewayUpdateParamsDLPObjectActionFlag  AIGatewayUpdateParamsDLPObjectAction = "FLAG"
)

func (r AIGatewayUpdateParamsDLPObjectAction) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsDLPObjectActionBlock, AIGatewayUpdateParamsDLPObjectActionFlag:
		return true
	}
	return false
}

type AIGatewayUpdateParamsDLPAction string

const (
	AIGatewayUpdateParamsDLPActionBlock AIGatewayUpdateParamsDLPAction = "BLOCK"
	AIGatewayUpdateParamsDLPActionFlag  AIGatewayUpdateParamsDLPAction = "FLAG"
)

func (r AIGatewayUpdateParamsDLPAction) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsDLPActionBlock, AIGatewayUpdateParamsDLPActionFlag:
		return true
	}
	return false
}

type AIGatewayUpdateParamsLogManagementStrategy string

const (
	AIGatewayUpdateParamsLogManagementStrategyStopInserting AIGatewayUpdateParamsLogManagementStrategy = "STOP_INSERTING"
	AIGatewayUpdateParamsLogManagementStrategyDeleteOldest  AIGatewayUpdateParamsLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayUpdateParamsLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsLogManagementStrategyStopInserting, AIGatewayUpdateParamsLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayUpdateParamsOtel struct {
	Authorization param.Field[string]            `json:"authorization,required"`
	Headers       param.Field[map[string]string] `json:"headers,required"`
	URL           param.Field[string]            `json:"url,required"`
}

func (r AIGatewayUpdateParamsOtel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayUpdateParamsStripe struct {
	Authorization param.Field[string]                                  `json:"authorization,required"`
	UsageEvents   param.Field[[]AIGatewayUpdateParamsStripeUsageEvent] `json:"usage_events,required"`
}

func (r AIGatewayUpdateParamsStripe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayUpdateParamsStripeUsageEvent struct {
	Payload param.Field[string] `json:"payload,required"`
}

func (r AIGatewayUpdateParamsStripeUsageEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayUpdateResponseEnvelope struct {
	Result  AIGatewayUpdateResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    aiGatewayUpdateResponseEnvelopeJSON `json:"-"`
}

// aiGatewayUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIGatewayUpdateResponseEnvelope]
type aiGatewayUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
	// Search by id
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AIGatewayListParams]'s query parameters as `url.Values`.
func (r AIGatewayListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AIGatewayDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AIGatewayDeleteResponseEnvelope struct {
	Result  AIGatewayDeleteResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    aiGatewayDeleteResponseEnvelopeJSON `json:"-"`
}

// aiGatewayDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIGatewayDeleteResponseEnvelope]
type aiGatewayDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AIGatewayGetResponseEnvelope struct {
	Result  AIGatewayGetResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    aiGatewayGetResponseEnvelopeJSON `json:"-"`
}

// aiGatewayGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIGatewayGetResponseEnvelope]
type aiGatewayGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
