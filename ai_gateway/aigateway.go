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

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
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
	Billing         *BillingService
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
	r.Billing = NewBillingService(opts...)
	return
}

// Creates a new AI Gateway.
func (r *AIGatewayService) New(ctx context.Context, params AIGatewayNewParams, opts ...option.RequestOption) (res *AIGatewayNewResponse, err error) {
	var env AIGatewayNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an existing AI Gateway dataset.
func (r *AIGatewayService) Update(ctx context.Context, id string, params AIGatewayUpdateParams, opts ...option.RequestOption) (res *AIGatewayUpdateResponse, err error) {
	var env AIGatewayUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all AI Gateway evaluator types configured for the account.
func (r *AIGatewayService) List(ctx context.Context, params AIGatewayListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AIGatewayListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
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

// Lists all AI Gateway evaluator types configured for the account.
func (r *AIGatewayService) ListAutoPaging(ctx context.Context, params AIGatewayListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AIGatewayListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an AI Gateway dataset.
func (r *AIGatewayService) Delete(ctx context.Context, id string, body AIGatewayDeleteParams, opts ...option.RequestOption) (res *AIGatewayDeleteResponse, err error) {
	var env AIGatewayDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves details for a specific AI Gateway dataset.
func (r *AIGatewayService) Get(ctx context.Context, id string, query AIGatewayGetParams, opts ...option.RequestOption) (res *AIGatewayGetResponse, err error) {
	var env AIGatewayGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AIGatewayNewResponse struct {
	// gateway id
	ID                      string                                    `json:"id" api:"required"`
	CacheInvalidateOnUpdate bool                                      `json:"cache_invalidate_on_update" api:"required"`
	CacheTTL                int64                                     `json:"cache_ttl" api:"required,nullable"`
	CollectLogs             bool                                      `json:"collect_logs" api:"required"`
	CreatedAt               time.Time                                 `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt              time.Time                                 `json:"modified_at" api:"required" format:"date-time"`
	RateLimitingInterval    int64                                     `json:"rate_limiting_interval" api:"required,nullable"`
	RateLimitingLimit       int64                                     `json:"rate_limiting_limit" api:"required,nullable"`
	Authentication          bool                                      `json:"authentication"`
	DLP                     AIGatewayNewResponseDLP                   `json:"dlp"`
	Guardrails              AIGatewayNewResponseGuardrails            `json:"guardrails" api:"nullable"`
	IsDefault               bool                                      `json:"is_default"`
	LogManagement           int64                                     `json:"log_management" api:"nullable"`
	LogManagementStrategy   AIGatewayNewResponseLogManagementStrategy `json:"log_management_strategy" api:"nullable"`
	Logpush                 bool                                      `json:"logpush"`
	LogpushPublicKey        string                                    `json:"logpush_public_key" api:"nullable"`
	Otel                    []AIGatewayNewResponseOtel                `json:"otel" api:"nullable"`
	RateLimitingTechnique   AIGatewayNewResponseRateLimitingTechnique `json:"rate_limiting_technique" api:"nullable"`
	// Backoff strategy for retry delays
	RetryBackoff AIGatewayNewResponseRetryBackoff `json:"retry_backoff" api:"nullable"`
	// Delay between retry attempts in milliseconds (0-5000)
	RetryDelay int64 `json:"retry_delay" api:"nullable"`
	// Maximum number of retry attempts for failed requests (1-5)
	RetryMaxAttempts int64                      `json:"retry_max_attempts" api:"nullable"`
	StoreID          string                     `json:"store_id" api:"nullable"`
	Stripe           AIGatewayNewResponseStripe `json:"stripe" api:"nullable"`
	// Controls how Workers AI inference calls routed through this gateway are billed.
	// Only 'postpaid' is currently supported.
	WorkersAIBillingMode AIGatewayNewResponseWorkersAIBillingMode `json:"workers_ai_billing_mode"`
	Zdr                  bool                                     `json:"zdr"`
	JSON                 aiGatewayNewResponseJSON                 `json:"-"`
}

// aiGatewayNewResponseJSON contains the JSON metadata for the struct
// [AIGatewayNewResponse]
type aiGatewayNewResponseJSON struct {
	ID                      apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	Authentication          apijson.Field
	DLP                     apijson.Field
	Guardrails              apijson.Field
	IsDefault               apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	Otel                    apijson.Field
	RateLimitingTechnique   apijson.Field
	RetryBackoff            apijson.Field
	RetryDelay              apijson.Field
	RetryMaxAttempts        apijson.Field
	StoreID                 apijson.Field
	Stripe                  apijson.Field
	WorkersAIBillingMode    apijson.Field
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

type AIGatewayNewResponseDLP struct {
	Enabled bool                          `json:"enabled" api:"required"`
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
	Action   AIGatewayNewResponseDLPObjectAction `json:"action" api:"required"`
	Enabled  bool                                `json:"enabled" api:"required"`
	Profiles []string                            `json:"profiles" api:"required"`
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

type AIGatewayNewResponseGuardrails struct {
	Prompt   AIGatewayNewResponseGuardrailsPrompt   `json:"prompt" api:"required"`
	Response AIGatewayNewResponseGuardrailsResponse `json:"response" api:"required"`
	JSON     aiGatewayNewResponseGuardrailsJSON     `json:"-"`
}

// aiGatewayNewResponseGuardrailsJSON contains the JSON metadata for the struct
// [AIGatewayNewResponseGuardrails]
type aiGatewayNewResponseGuardrailsJSON struct {
	Prompt      apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayNewResponseGuardrails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseGuardrailsJSON) RawJSON() string {
	return r.raw
}

type AIGatewayNewResponseGuardrailsPrompt struct {
	P1   AIGatewayNewResponseGuardrailsPromptP1   `json:"P1"`
	S1   AIGatewayNewResponseGuardrailsPromptS1   `json:"S1"`
	S10  AIGatewayNewResponseGuardrailsPromptS10  `json:"S10"`
	S11  AIGatewayNewResponseGuardrailsPromptS11  `json:"S11"`
	S12  AIGatewayNewResponseGuardrailsPromptS12  `json:"S12"`
	S13  AIGatewayNewResponseGuardrailsPromptS13  `json:"S13"`
	S2   AIGatewayNewResponseGuardrailsPromptS2   `json:"S2"`
	S3   AIGatewayNewResponseGuardrailsPromptS3   `json:"S3"`
	S4   AIGatewayNewResponseGuardrailsPromptS4   `json:"S4"`
	S5   AIGatewayNewResponseGuardrailsPromptS5   `json:"S5"`
	S6   AIGatewayNewResponseGuardrailsPromptS6   `json:"S6"`
	S7   AIGatewayNewResponseGuardrailsPromptS7   `json:"S7"`
	S8   AIGatewayNewResponseGuardrailsPromptS8   `json:"S8"`
	S9   AIGatewayNewResponseGuardrailsPromptS9   `json:"S9"`
	JSON aiGatewayNewResponseGuardrailsPromptJSON `json:"-"`
}

// aiGatewayNewResponseGuardrailsPromptJSON contains the JSON metadata for the
// struct [AIGatewayNewResponseGuardrailsPrompt]
type aiGatewayNewResponseGuardrailsPromptJSON struct {
	P1          apijson.Field
	S1          apijson.Field
	S10         apijson.Field
	S11         apijson.Field
	S12         apijson.Field
	S13         apijson.Field
	S2          apijson.Field
	S3          apijson.Field
	S4          apijson.Field
	S5          apijson.Field
	S6          apijson.Field
	S7          apijson.Field
	S8          apijson.Field
	S9          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayNewResponseGuardrailsPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseGuardrailsPromptJSON) RawJSON() string {
	return r.raw
}

type AIGatewayNewResponseGuardrailsPromptP1 string

const (
	AIGatewayNewResponseGuardrailsPromptP1Flag  AIGatewayNewResponseGuardrailsPromptP1 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptP1Block AIGatewayNewResponseGuardrailsPromptP1 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptP1) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptP1Flag, AIGatewayNewResponseGuardrailsPromptP1Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsPromptS1 string

const (
	AIGatewayNewResponseGuardrailsPromptS1Flag  AIGatewayNewResponseGuardrailsPromptS1 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptS1Block AIGatewayNewResponseGuardrailsPromptS1 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptS1) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptS1Flag, AIGatewayNewResponseGuardrailsPromptS1Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsPromptS10 string

const (
	AIGatewayNewResponseGuardrailsPromptS10Flag  AIGatewayNewResponseGuardrailsPromptS10 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptS10Block AIGatewayNewResponseGuardrailsPromptS10 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptS10) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptS10Flag, AIGatewayNewResponseGuardrailsPromptS10Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsPromptS11 string

const (
	AIGatewayNewResponseGuardrailsPromptS11Flag  AIGatewayNewResponseGuardrailsPromptS11 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptS11Block AIGatewayNewResponseGuardrailsPromptS11 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptS11) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptS11Flag, AIGatewayNewResponseGuardrailsPromptS11Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsPromptS12 string

const (
	AIGatewayNewResponseGuardrailsPromptS12Flag  AIGatewayNewResponseGuardrailsPromptS12 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptS12Block AIGatewayNewResponseGuardrailsPromptS12 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptS12) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptS12Flag, AIGatewayNewResponseGuardrailsPromptS12Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsPromptS13 string

const (
	AIGatewayNewResponseGuardrailsPromptS13Flag  AIGatewayNewResponseGuardrailsPromptS13 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptS13Block AIGatewayNewResponseGuardrailsPromptS13 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptS13) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptS13Flag, AIGatewayNewResponseGuardrailsPromptS13Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsPromptS2 string

const (
	AIGatewayNewResponseGuardrailsPromptS2Flag  AIGatewayNewResponseGuardrailsPromptS2 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptS2Block AIGatewayNewResponseGuardrailsPromptS2 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptS2) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptS2Flag, AIGatewayNewResponseGuardrailsPromptS2Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsPromptS3 string

const (
	AIGatewayNewResponseGuardrailsPromptS3Flag  AIGatewayNewResponseGuardrailsPromptS3 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptS3Block AIGatewayNewResponseGuardrailsPromptS3 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptS3) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptS3Flag, AIGatewayNewResponseGuardrailsPromptS3Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsPromptS4 string

const (
	AIGatewayNewResponseGuardrailsPromptS4Flag  AIGatewayNewResponseGuardrailsPromptS4 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptS4Block AIGatewayNewResponseGuardrailsPromptS4 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptS4) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptS4Flag, AIGatewayNewResponseGuardrailsPromptS4Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsPromptS5 string

const (
	AIGatewayNewResponseGuardrailsPromptS5Flag  AIGatewayNewResponseGuardrailsPromptS5 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptS5Block AIGatewayNewResponseGuardrailsPromptS5 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptS5) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptS5Flag, AIGatewayNewResponseGuardrailsPromptS5Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsPromptS6 string

const (
	AIGatewayNewResponseGuardrailsPromptS6Flag  AIGatewayNewResponseGuardrailsPromptS6 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptS6Block AIGatewayNewResponseGuardrailsPromptS6 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptS6) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptS6Flag, AIGatewayNewResponseGuardrailsPromptS6Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsPromptS7 string

const (
	AIGatewayNewResponseGuardrailsPromptS7Flag  AIGatewayNewResponseGuardrailsPromptS7 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptS7Block AIGatewayNewResponseGuardrailsPromptS7 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptS7) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptS7Flag, AIGatewayNewResponseGuardrailsPromptS7Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsPromptS8 string

const (
	AIGatewayNewResponseGuardrailsPromptS8Flag  AIGatewayNewResponseGuardrailsPromptS8 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptS8Block AIGatewayNewResponseGuardrailsPromptS8 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptS8) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptS8Flag, AIGatewayNewResponseGuardrailsPromptS8Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsPromptS9 string

const (
	AIGatewayNewResponseGuardrailsPromptS9Flag  AIGatewayNewResponseGuardrailsPromptS9 = "FLAG"
	AIGatewayNewResponseGuardrailsPromptS9Block AIGatewayNewResponseGuardrailsPromptS9 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsPromptS9) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsPromptS9Flag, AIGatewayNewResponseGuardrailsPromptS9Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponse struct {
	P1   AIGatewayNewResponseGuardrailsResponseP1   `json:"P1"`
	S1   AIGatewayNewResponseGuardrailsResponseS1   `json:"S1"`
	S10  AIGatewayNewResponseGuardrailsResponseS10  `json:"S10"`
	S11  AIGatewayNewResponseGuardrailsResponseS11  `json:"S11"`
	S12  AIGatewayNewResponseGuardrailsResponseS12  `json:"S12"`
	S13  AIGatewayNewResponseGuardrailsResponseS13  `json:"S13"`
	S2   AIGatewayNewResponseGuardrailsResponseS2   `json:"S2"`
	S3   AIGatewayNewResponseGuardrailsResponseS3   `json:"S3"`
	S4   AIGatewayNewResponseGuardrailsResponseS4   `json:"S4"`
	S5   AIGatewayNewResponseGuardrailsResponseS5   `json:"S5"`
	S6   AIGatewayNewResponseGuardrailsResponseS6   `json:"S6"`
	S7   AIGatewayNewResponseGuardrailsResponseS7   `json:"S7"`
	S8   AIGatewayNewResponseGuardrailsResponseS8   `json:"S8"`
	S9   AIGatewayNewResponseGuardrailsResponseS9   `json:"S9"`
	JSON aiGatewayNewResponseGuardrailsResponseJSON `json:"-"`
}

// aiGatewayNewResponseGuardrailsResponseJSON contains the JSON metadata for the
// struct [AIGatewayNewResponseGuardrailsResponse]
type aiGatewayNewResponseGuardrailsResponseJSON struct {
	P1          apijson.Field
	S1          apijson.Field
	S10         apijson.Field
	S11         apijson.Field
	S12         apijson.Field
	S13         apijson.Field
	S2          apijson.Field
	S3          apijson.Field
	S4          apijson.Field
	S5          apijson.Field
	S6          apijson.Field
	S7          apijson.Field
	S8          apijson.Field
	S9          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayNewResponseGuardrailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseGuardrailsResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayNewResponseGuardrailsResponseP1 string

const (
	AIGatewayNewResponseGuardrailsResponseP1Flag  AIGatewayNewResponseGuardrailsResponseP1 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseP1Block AIGatewayNewResponseGuardrailsResponseP1 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseP1) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseP1Flag, AIGatewayNewResponseGuardrailsResponseP1Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponseS1 string

const (
	AIGatewayNewResponseGuardrailsResponseS1Flag  AIGatewayNewResponseGuardrailsResponseS1 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseS1Block AIGatewayNewResponseGuardrailsResponseS1 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseS1) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseS1Flag, AIGatewayNewResponseGuardrailsResponseS1Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponseS10 string

const (
	AIGatewayNewResponseGuardrailsResponseS10Flag  AIGatewayNewResponseGuardrailsResponseS10 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseS10Block AIGatewayNewResponseGuardrailsResponseS10 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseS10) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseS10Flag, AIGatewayNewResponseGuardrailsResponseS10Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponseS11 string

const (
	AIGatewayNewResponseGuardrailsResponseS11Flag  AIGatewayNewResponseGuardrailsResponseS11 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseS11Block AIGatewayNewResponseGuardrailsResponseS11 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseS11) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseS11Flag, AIGatewayNewResponseGuardrailsResponseS11Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponseS12 string

const (
	AIGatewayNewResponseGuardrailsResponseS12Flag  AIGatewayNewResponseGuardrailsResponseS12 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseS12Block AIGatewayNewResponseGuardrailsResponseS12 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseS12) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseS12Flag, AIGatewayNewResponseGuardrailsResponseS12Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponseS13 string

const (
	AIGatewayNewResponseGuardrailsResponseS13Flag  AIGatewayNewResponseGuardrailsResponseS13 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseS13Block AIGatewayNewResponseGuardrailsResponseS13 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseS13) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseS13Flag, AIGatewayNewResponseGuardrailsResponseS13Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponseS2 string

const (
	AIGatewayNewResponseGuardrailsResponseS2Flag  AIGatewayNewResponseGuardrailsResponseS2 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseS2Block AIGatewayNewResponseGuardrailsResponseS2 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseS2) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseS2Flag, AIGatewayNewResponseGuardrailsResponseS2Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponseS3 string

const (
	AIGatewayNewResponseGuardrailsResponseS3Flag  AIGatewayNewResponseGuardrailsResponseS3 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseS3Block AIGatewayNewResponseGuardrailsResponseS3 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseS3) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseS3Flag, AIGatewayNewResponseGuardrailsResponseS3Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponseS4 string

const (
	AIGatewayNewResponseGuardrailsResponseS4Flag  AIGatewayNewResponseGuardrailsResponseS4 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseS4Block AIGatewayNewResponseGuardrailsResponseS4 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseS4) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseS4Flag, AIGatewayNewResponseGuardrailsResponseS4Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponseS5 string

const (
	AIGatewayNewResponseGuardrailsResponseS5Flag  AIGatewayNewResponseGuardrailsResponseS5 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseS5Block AIGatewayNewResponseGuardrailsResponseS5 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseS5) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseS5Flag, AIGatewayNewResponseGuardrailsResponseS5Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponseS6 string

const (
	AIGatewayNewResponseGuardrailsResponseS6Flag  AIGatewayNewResponseGuardrailsResponseS6 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseS6Block AIGatewayNewResponseGuardrailsResponseS6 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseS6) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseS6Flag, AIGatewayNewResponseGuardrailsResponseS6Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponseS7 string

const (
	AIGatewayNewResponseGuardrailsResponseS7Flag  AIGatewayNewResponseGuardrailsResponseS7 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseS7Block AIGatewayNewResponseGuardrailsResponseS7 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseS7) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseS7Flag, AIGatewayNewResponseGuardrailsResponseS7Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponseS8 string

const (
	AIGatewayNewResponseGuardrailsResponseS8Flag  AIGatewayNewResponseGuardrailsResponseS8 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseS8Block AIGatewayNewResponseGuardrailsResponseS8 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseS8) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseS8Flag, AIGatewayNewResponseGuardrailsResponseS8Block:
		return true
	}
	return false
}

type AIGatewayNewResponseGuardrailsResponseS9 string

const (
	AIGatewayNewResponseGuardrailsResponseS9Flag  AIGatewayNewResponseGuardrailsResponseS9 = "FLAG"
	AIGatewayNewResponseGuardrailsResponseS9Block AIGatewayNewResponseGuardrailsResponseS9 = "BLOCK"
)

func (r AIGatewayNewResponseGuardrailsResponseS9) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseGuardrailsResponseS9Flag, AIGatewayNewResponseGuardrailsResponseS9Block:
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
	Authorization string                              `json:"authorization" api:"required"`
	Headers       map[string]string                   `json:"headers" api:"required"`
	URL           string                              `json:"url" api:"required" format:"uri"`
	ContentType   AIGatewayNewResponseOtelContentType `json:"content_type"`
	JSON          aiGatewayNewResponseOtelJSON        `json:"-"`
}

// aiGatewayNewResponseOtelJSON contains the JSON metadata for the struct
// [AIGatewayNewResponseOtel]
type aiGatewayNewResponseOtelJSON struct {
	Authorization apijson.Field
	Headers       apijson.Field
	URL           apijson.Field
	ContentType   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayNewResponseOtel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseOtelJSON) RawJSON() string {
	return r.raw
}

type AIGatewayNewResponseOtelContentType string

const (
	AIGatewayNewResponseOtelContentTypeJson     AIGatewayNewResponseOtelContentType = "json"
	AIGatewayNewResponseOtelContentTypeProtobuf AIGatewayNewResponseOtelContentType = "protobuf"
)

func (r AIGatewayNewResponseOtelContentType) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseOtelContentTypeJson, AIGatewayNewResponseOtelContentTypeProtobuf:
		return true
	}
	return false
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

// Backoff strategy for retry delays
type AIGatewayNewResponseRetryBackoff string

const (
	AIGatewayNewResponseRetryBackoffConstant    AIGatewayNewResponseRetryBackoff = "constant"
	AIGatewayNewResponseRetryBackoffLinear      AIGatewayNewResponseRetryBackoff = "linear"
	AIGatewayNewResponseRetryBackoffExponential AIGatewayNewResponseRetryBackoff = "exponential"
)

func (r AIGatewayNewResponseRetryBackoff) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseRetryBackoffConstant, AIGatewayNewResponseRetryBackoffLinear, AIGatewayNewResponseRetryBackoffExponential:
		return true
	}
	return false
}

type AIGatewayNewResponseStripe struct {
	Authorization string                                 `json:"authorization" api:"required"`
	UsageEvents   []AIGatewayNewResponseStripeUsageEvent `json:"usage_events" api:"required"`
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
	Payload string                                   `json:"payload" api:"required"`
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

// Controls how Workers AI inference calls routed through this gateway are billed.
// Only 'postpaid' is currently supported.
type AIGatewayNewResponseWorkersAIBillingMode string

const (
	AIGatewayNewResponseWorkersAIBillingModePostpaid AIGatewayNewResponseWorkersAIBillingMode = "postpaid"
)

func (r AIGatewayNewResponseWorkersAIBillingMode) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseWorkersAIBillingModePostpaid:
		return true
	}
	return false
}

type AIGatewayUpdateResponse struct {
	// gateway id
	ID                      string                                       `json:"id" api:"required"`
	CacheInvalidateOnUpdate bool                                         `json:"cache_invalidate_on_update" api:"required"`
	CacheTTL                int64                                        `json:"cache_ttl" api:"required,nullable"`
	CollectLogs             bool                                         `json:"collect_logs" api:"required"`
	CreatedAt               time.Time                                    `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt              time.Time                                    `json:"modified_at" api:"required" format:"date-time"`
	RateLimitingInterval    int64                                        `json:"rate_limiting_interval" api:"required,nullable"`
	RateLimitingLimit       int64                                        `json:"rate_limiting_limit" api:"required,nullable"`
	Authentication          bool                                         `json:"authentication"`
	DLP                     AIGatewayUpdateResponseDLP                   `json:"dlp"`
	Guardrails              AIGatewayUpdateResponseGuardrails            `json:"guardrails" api:"nullable"`
	IsDefault               bool                                         `json:"is_default"`
	LogManagement           int64                                        `json:"log_management" api:"nullable"`
	LogManagementStrategy   AIGatewayUpdateResponseLogManagementStrategy `json:"log_management_strategy" api:"nullable"`
	Logpush                 bool                                         `json:"logpush"`
	LogpushPublicKey        string                                       `json:"logpush_public_key" api:"nullable"`
	Otel                    []AIGatewayUpdateResponseOtel                `json:"otel" api:"nullable"`
	RateLimitingTechnique   AIGatewayUpdateResponseRateLimitingTechnique `json:"rate_limiting_technique" api:"nullable"`
	// Backoff strategy for retry delays
	RetryBackoff AIGatewayUpdateResponseRetryBackoff `json:"retry_backoff" api:"nullable"`
	// Delay between retry attempts in milliseconds (0-5000)
	RetryDelay int64 `json:"retry_delay" api:"nullable"`
	// Maximum number of retry attempts for failed requests (1-5)
	RetryMaxAttempts int64                         `json:"retry_max_attempts" api:"nullable"`
	StoreID          string                        `json:"store_id" api:"nullable"`
	Stripe           AIGatewayUpdateResponseStripe `json:"stripe" api:"nullable"`
	// Controls how Workers AI inference calls routed through this gateway are billed.
	// Only 'postpaid' is currently supported.
	WorkersAIBillingMode AIGatewayUpdateResponseWorkersAIBillingMode `json:"workers_ai_billing_mode"`
	Zdr                  bool                                        `json:"zdr"`
	JSON                 aiGatewayUpdateResponseJSON                 `json:"-"`
}

// aiGatewayUpdateResponseJSON contains the JSON metadata for the struct
// [AIGatewayUpdateResponse]
type aiGatewayUpdateResponseJSON struct {
	ID                      apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	Authentication          apijson.Field
	DLP                     apijson.Field
	Guardrails              apijson.Field
	IsDefault               apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	Otel                    apijson.Field
	RateLimitingTechnique   apijson.Field
	RetryBackoff            apijson.Field
	RetryDelay              apijson.Field
	RetryMaxAttempts        apijson.Field
	StoreID                 apijson.Field
	Stripe                  apijson.Field
	WorkersAIBillingMode    apijson.Field
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

type AIGatewayUpdateResponseDLP struct {
	Enabled bool                             `json:"enabled" api:"required"`
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
	Action   AIGatewayUpdateResponseDLPObjectAction `json:"action" api:"required"`
	Enabled  bool                                   `json:"enabled" api:"required"`
	Profiles []string                               `json:"profiles" api:"required"`
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

type AIGatewayUpdateResponseGuardrails struct {
	Prompt   AIGatewayUpdateResponseGuardrailsPrompt   `json:"prompt" api:"required"`
	Response AIGatewayUpdateResponseGuardrailsResponse `json:"response" api:"required"`
	JSON     aiGatewayUpdateResponseGuardrailsJSON     `json:"-"`
}

// aiGatewayUpdateResponseGuardrailsJSON contains the JSON metadata for the struct
// [AIGatewayUpdateResponseGuardrails]
type aiGatewayUpdateResponseGuardrailsJSON struct {
	Prompt      apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayUpdateResponseGuardrails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseGuardrailsJSON) RawJSON() string {
	return r.raw
}

type AIGatewayUpdateResponseGuardrailsPrompt struct {
	P1   AIGatewayUpdateResponseGuardrailsPromptP1   `json:"P1"`
	S1   AIGatewayUpdateResponseGuardrailsPromptS1   `json:"S1"`
	S10  AIGatewayUpdateResponseGuardrailsPromptS10  `json:"S10"`
	S11  AIGatewayUpdateResponseGuardrailsPromptS11  `json:"S11"`
	S12  AIGatewayUpdateResponseGuardrailsPromptS12  `json:"S12"`
	S13  AIGatewayUpdateResponseGuardrailsPromptS13  `json:"S13"`
	S2   AIGatewayUpdateResponseGuardrailsPromptS2   `json:"S2"`
	S3   AIGatewayUpdateResponseGuardrailsPromptS3   `json:"S3"`
	S4   AIGatewayUpdateResponseGuardrailsPromptS4   `json:"S4"`
	S5   AIGatewayUpdateResponseGuardrailsPromptS5   `json:"S5"`
	S6   AIGatewayUpdateResponseGuardrailsPromptS6   `json:"S6"`
	S7   AIGatewayUpdateResponseGuardrailsPromptS7   `json:"S7"`
	S8   AIGatewayUpdateResponseGuardrailsPromptS8   `json:"S8"`
	S9   AIGatewayUpdateResponseGuardrailsPromptS9   `json:"S9"`
	JSON aiGatewayUpdateResponseGuardrailsPromptJSON `json:"-"`
}

// aiGatewayUpdateResponseGuardrailsPromptJSON contains the JSON metadata for the
// struct [AIGatewayUpdateResponseGuardrailsPrompt]
type aiGatewayUpdateResponseGuardrailsPromptJSON struct {
	P1          apijson.Field
	S1          apijson.Field
	S10         apijson.Field
	S11         apijson.Field
	S12         apijson.Field
	S13         apijson.Field
	S2          apijson.Field
	S3          apijson.Field
	S4          apijson.Field
	S5          apijson.Field
	S6          apijson.Field
	S7          apijson.Field
	S8          apijson.Field
	S9          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayUpdateResponseGuardrailsPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseGuardrailsPromptJSON) RawJSON() string {
	return r.raw
}

type AIGatewayUpdateResponseGuardrailsPromptP1 string

const (
	AIGatewayUpdateResponseGuardrailsPromptP1Flag  AIGatewayUpdateResponseGuardrailsPromptP1 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptP1Block AIGatewayUpdateResponseGuardrailsPromptP1 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptP1) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptP1Flag, AIGatewayUpdateResponseGuardrailsPromptP1Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsPromptS1 string

const (
	AIGatewayUpdateResponseGuardrailsPromptS1Flag  AIGatewayUpdateResponseGuardrailsPromptS1 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptS1Block AIGatewayUpdateResponseGuardrailsPromptS1 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptS1) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptS1Flag, AIGatewayUpdateResponseGuardrailsPromptS1Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsPromptS10 string

const (
	AIGatewayUpdateResponseGuardrailsPromptS10Flag  AIGatewayUpdateResponseGuardrailsPromptS10 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptS10Block AIGatewayUpdateResponseGuardrailsPromptS10 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptS10) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptS10Flag, AIGatewayUpdateResponseGuardrailsPromptS10Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsPromptS11 string

const (
	AIGatewayUpdateResponseGuardrailsPromptS11Flag  AIGatewayUpdateResponseGuardrailsPromptS11 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptS11Block AIGatewayUpdateResponseGuardrailsPromptS11 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptS11) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptS11Flag, AIGatewayUpdateResponseGuardrailsPromptS11Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsPromptS12 string

const (
	AIGatewayUpdateResponseGuardrailsPromptS12Flag  AIGatewayUpdateResponseGuardrailsPromptS12 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptS12Block AIGatewayUpdateResponseGuardrailsPromptS12 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptS12) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptS12Flag, AIGatewayUpdateResponseGuardrailsPromptS12Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsPromptS13 string

const (
	AIGatewayUpdateResponseGuardrailsPromptS13Flag  AIGatewayUpdateResponseGuardrailsPromptS13 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptS13Block AIGatewayUpdateResponseGuardrailsPromptS13 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptS13) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptS13Flag, AIGatewayUpdateResponseGuardrailsPromptS13Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsPromptS2 string

const (
	AIGatewayUpdateResponseGuardrailsPromptS2Flag  AIGatewayUpdateResponseGuardrailsPromptS2 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptS2Block AIGatewayUpdateResponseGuardrailsPromptS2 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptS2) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptS2Flag, AIGatewayUpdateResponseGuardrailsPromptS2Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsPromptS3 string

const (
	AIGatewayUpdateResponseGuardrailsPromptS3Flag  AIGatewayUpdateResponseGuardrailsPromptS3 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptS3Block AIGatewayUpdateResponseGuardrailsPromptS3 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptS3) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptS3Flag, AIGatewayUpdateResponseGuardrailsPromptS3Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsPromptS4 string

const (
	AIGatewayUpdateResponseGuardrailsPromptS4Flag  AIGatewayUpdateResponseGuardrailsPromptS4 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptS4Block AIGatewayUpdateResponseGuardrailsPromptS4 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptS4) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptS4Flag, AIGatewayUpdateResponseGuardrailsPromptS4Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsPromptS5 string

const (
	AIGatewayUpdateResponseGuardrailsPromptS5Flag  AIGatewayUpdateResponseGuardrailsPromptS5 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptS5Block AIGatewayUpdateResponseGuardrailsPromptS5 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptS5) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptS5Flag, AIGatewayUpdateResponseGuardrailsPromptS5Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsPromptS6 string

const (
	AIGatewayUpdateResponseGuardrailsPromptS6Flag  AIGatewayUpdateResponseGuardrailsPromptS6 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptS6Block AIGatewayUpdateResponseGuardrailsPromptS6 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptS6) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptS6Flag, AIGatewayUpdateResponseGuardrailsPromptS6Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsPromptS7 string

const (
	AIGatewayUpdateResponseGuardrailsPromptS7Flag  AIGatewayUpdateResponseGuardrailsPromptS7 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptS7Block AIGatewayUpdateResponseGuardrailsPromptS7 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptS7) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptS7Flag, AIGatewayUpdateResponseGuardrailsPromptS7Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsPromptS8 string

const (
	AIGatewayUpdateResponseGuardrailsPromptS8Flag  AIGatewayUpdateResponseGuardrailsPromptS8 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptS8Block AIGatewayUpdateResponseGuardrailsPromptS8 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptS8) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptS8Flag, AIGatewayUpdateResponseGuardrailsPromptS8Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsPromptS9 string

const (
	AIGatewayUpdateResponseGuardrailsPromptS9Flag  AIGatewayUpdateResponseGuardrailsPromptS9 = "FLAG"
	AIGatewayUpdateResponseGuardrailsPromptS9Block AIGatewayUpdateResponseGuardrailsPromptS9 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsPromptS9) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsPromptS9Flag, AIGatewayUpdateResponseGuardrailsPromptS9Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponse struct {
	P1   AIGatewayUpdateResponseGuardrailsResponseP1   `json:"P1"`
	S1   AIGatewayUpdateResponseGuardrailsResponseS1   `json:"S1"`
	S10  AIGatewayUpdateResponseGuardrailsResponseS10  `json:"S10"`
	S11  AIGatewayUpdateResponseGuardrailsResponseS11  `json:"S11"`
	S12  AIGatewayUpdateResponseGuardrailsResponseS12  `json:"S12"`
	S13  AIGatewayUpdateResponseGuardrailsResponseS13  `json:"S13"`
	S2   AIGatewayUpdateResponseGuardrailsResponseS2   `json:"S2"`
	S3   AIGatewayUpdateResponseGuardrailsResponseS3   `json:"S3"`
	S4   AIGatewayUpdateResponseGuardrailsResponseS4   `json:"S4"`
	S5   AIGatewayUpdateResponseGuardrailsResponseS5   `json:"S5"`
	S6   AIGatewayUpdateResponseGuardrailsResponseS6   `json:"S6"`
	S7   AIGatewayUpdateResponseGuardrailsResponseS7   `json:"S7"`
	S8   AIGatewayUpdateResponseGuardrailsResponseS8   `json:"S8"`
	S9   AIGatewayUpdateResponseGuardrailsResponseS9   `json:"S9"`
	JSON aiGatewayUpdateResponseGuardrailsResponseJSON `json:"-"`
}

// aiGatewayUpdateResponseGuardrailsResponseJSON contains the JSON metadata for the
// struct [AIGatewayUpdateResponseGuardrailsResponse]
type aiGatewayUpdateResponseGuardrailsResponseJSON struct {
	P1          apijson.Field
	S1          apijson.Field
	S10         apijson.Field
	S11         apijson.Field
	S12         apijson.Field
	S13         apijson.Field
	S2          apijson.Field
	S3          apijson.Field
	S4          apijson.Field
	S5          apijson.Field
	S6          apijson.Field
	S7          apijson.Field
	S8          apijson.Field
	S9          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayUpdateResponseGuardrailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseGuardrailsResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayUpdateResponseGuardrailsResponseP1 string

const (
	AIGatewayUpdateResponseGuardrailsResponseP1Flag  AIGatewayUpdateResponseGuardrailsResponseP1 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseP1Block AIGatewayUpdateResponseGuardrailsResponseP1 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseP1) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseP1Flag, AIGatewayUpdateResponseGuardrailsResponseP1Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponseS1 string

const (
	AIGatewayUpdateResponseGuardrailsResponseS1Flag  AIGatewayUpdateResponseGuardrailsResponseS1 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseS1Block AIGatewayUpdateResponseGuardrailsResponseS1 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseS1) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseS1Flag, AIGatewayUpdateResponseGuardrailsResponseS1Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponseS10 string

const (
	AIGatewayUpdateResponseGuardrailsResponseS10Flag  AIGatewayUpdateResponseGuardrailsResponseS10 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseS10Block AIGatewayUpdateResponseGuardrailsResponseS10 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseS10) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseS10Flag, AIGatewayUpdateResponseGuardrailsResponseS10Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponseS11 string

const (
	AIGatewayUpdateResponseGuardrailsResponseS11Flag  AIGatewayUpdateResponseGuardrailsResponseS11 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseS11Block AIGatewayUpdateResponseGuardrailsResponseS11 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseS11) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseS11Flag, AIGatewayUpdateResponseGuardrailsResponseS11Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponseS12 string

const (
	AIGatewayUpdateResponseGuardrailsResponseS12Flag  AIGatewayUpdateResponseGuardrailsResponseS12 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseS12Block AIGatewayUpdateResponseGuardrailsResponseS12 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseS12) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseS12Flag, AIGatewayUpdateResponseGuardrailsResponseS12Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponseS13 string

const (
	AIGatewayUpdateResponseGuardrailsResponseS13Flag  AIGatewayUpdateResponseGuardrailsResponseS13 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseS13Block AIGatewayUpdateResponseGuardrailsResponseS13 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseS13) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseS13Flag, AIGatewayUpdateResponseGuardrailsResponseS13Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponseS2 string

const (
	AIGatewayUpdateResponseGuardrailsResponseS2Flag  AIGatewayUpdateResponseGuardrailsResponseS2 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseS2Block AIGatewayUpdateResponseGuardrailsResponseS2 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseS2) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseS2Flag, AIGatewayUpdateResponseGuardrailsResponseS2Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponseS3 string

const (
	AIGatewayUpdateResponseGuardrailsResponseS3Flag  AIGatewayUpdateResponseGuardrailsResponseS3 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseS3Block AIGatewayUpdateResponseGuardrailsResponseS3 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseS3) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseS3Flag, AIGatewayUpdateResponseGuardrailsResponseS3Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponseS4 string

const (
	AIGatewayUpdateResponseGuardrailsResponseS4Flag  AIGatewayUpdateResponseGuardrailsResponseS4 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseS4Block AIGatewayUpdateResponseGuardrailsResponseS4 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseS4) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseS4Flag, AIGatewayUpdateResponseGuardrailsResponseS4Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponseS5 string

const (
	AIGatewayUpdateResponseGuardrailsResponseS5Flag  AIGatewayUpdateResponseGuardrailsResponseS5 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseS5Block AIGatewayUpdateResponseGuardrailsResponseS5 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseS5) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseS5Flag, AIGatewayUpdateResponseGuardrailsResponseS5Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponseS6 string

const (
	AIGatewayUpdateResponseGuardrailsResponseS6Flag  AIGatewayUpdateResponseGuardrailsResponseS6 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseS6Block AIGatewayUpdateResponseGuardrailsResponseS6 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseS6) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseS6Flag, AIGatewayUpdateResponseGuardrailsResponseS6Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponseS7 string

const (
	AIGatewayUpdateResponseGuardrailsResponseS7Flag  AIGatewayUpdateResponseGuardrailsResponseS7 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseS7Block AIGatewayUpdateResponseGuardrailsResponseS7 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseS7) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseS7Flag, AIGatewayUpdateResponseGuardrailsResponseS7Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponseS8 string

const (
	AIGatewayUpdateResponseGuardrailsResponseS8Flag  AIGatewayUpdateResponseGuardrailsResponseS8 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseS8Block AIGatewayUpdateResponseGuardrailsResponseS8 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseS8) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseS8Flag, AIGatewayUpdateResponseGuardrailsResponseS8Block:
		return true
	}
	return false
}

type AIGatewayUpdateResponseGuardrailsResponseS9 string

const (
	AIGatewayUpdateResponseGuardrailsResponseS9Flag  AIGatewayUpdateResponseGuardrailsResponseS9 = "FLAG"
	AIGatewayUpdateResponseGuardrailsResponseS9Block AIGatewayUpdateResponseGuardrailsResponseS9 = "BLOCK"
)

func (r AIGatewayUpdateResponseGuardrailsResponseS9) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseGuardrailsResponseS9Flag, AIGatewayUpdateResponseGuardrailsResponseS9Block:
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
	Authorization string                                 `json:"authorization" api:"required"`
	Headers       map[string]string                      `json:"headers" api:"required"`
	URL           string                                 `json:"url" api:"required" format:"uri"`
	ContentType   AIGatewayUpdateResponseOtelContentType `json:"content_type"`
	JSON          aiGatewayUpdateResponseOtelJSON        `json:"-"`
}

// aiGatewayUpdateResponseOtelJSON contains the JSON metadata for the struct
// [AIGatewayUpdateResponseOtel]
type aiGatewayUpdateResponseOtelJSON struct {
	Authorization apijson.Field
	Headers       apijson.Field
	URL           apijson.Field
	ContentType   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayUpdateResponseOtel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseOtelJSON) RawJSON() string {
	return r.raw
}

type AIGatewayUpdateResponseOtelContentType string

const (
	AIGatewayUpdateResponseOtelContentTypeJson     AIGatewayUpdateResponseOtelContentType = "json"
	AIGatewayUpdateResponseOtelContentTypeProtobuf AIGatewayUpdateResponseOtelContentType = "protobuf"
)

func (r AIGatewayUpdateResponseOtelContentType) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseOtelContentTypeJson, AIGatewayUpdateResponseOtelContentTypeProtobuf:
		return true
	}
	return false
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

// Backoff strategy for retry delays
type AIGatewayUpdateResponseRetryBackoff string

const (
	AIGatewayUpdateResponseRetryBackoffConstant    AIGatewayUpdateResponseRetryBackoff = "constant"
	AIGatewayUpdateResponseRetryBackoffLinear      AIGatewayUpdateResponseRetryBackoff = "linear"
	AIGatewayUpdateResponseRetryBackoffExponential AIGatewayUpdateResponseRetryBackoff = "exponential"
)

func (r AIGatewayUpdateResponseRetryBackoff) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseRetryBackoffConstant, AIGatewayUpdateResponseRetryBackoffLinear, AIGatewayUpdateResponseRetryBackoffExponential:
		return true
	}
	return false
}

type AIGatewayUpdateResponseStripe struct {
	Authorization string                                    `json:"authorization" api:"required"`
	UsageEvents   []AIGatewayUpdateResponseStripeUsageEvent `json:"usage_events" api:"required"`
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
	Payload string                                      `json:"payload" api:"required"`
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

// Controls how Workers AI inference calls routed through this gateway are billed.
// Only 'postpaid' is currently supported.
type AIGatewayUpdateResponseWorkersAIBillingMode string

const (
	AIGatewayUpdateResponseWorkersAIBillingModePostpaid AIGatewayUpdateResponseWorkersAIBillingMode = "postpaid"
)

func (r AIGatewayUpdateResponseWorkersAIBillingMode) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseWorkersAIBillingModePostpaid:
		return true
	}
	return false
}

type AIGatewayListResponse struct {
	// gateway id
	ID                      string                                     `json:"id" api:"required"`
	CacheInvalidateOnUpdate bool                                       `json:"cache_invalidate_on_update" api:"required"`
	CacheTTL                int64                                      `json:"cache_ttl" api:"required,nullable"`
	CollectLogs             bool                                       `json:"collect_logs" api:"required"`
	CreatedAt               time.Time                                  `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt              time.Time                                  `json:"modified_at" api:"required" format:"date-time"`
	RateLimitingInterval    int64                                      `json:"rate_limiting_interval" api:"required,nullable"`
	RateLimitingLimit       int64                                      `json:"rate_limiting_limit" api:"required,nullable"`
	Authentication          bool                                       `json:"authentication"`
	DLP                     AIGatewayListResponseDLP                   `json:"dlp"`
	Guardrails              AIGatewayListResponseGuardrails            `json:"guardrails" api:"nullable"`
	IsDefault               bool                                       `json:"is_default"`
	LogManagement           int64                                      `json:"log_management" api:"nullable"`
	LogManagementStrategy   AIGatewayListResponseLogManagementStrategy `json:"log_management_strategy" api:"nullable"`
	Logpush                 bool                                       `json:"logpush"`
	LogpushPublicKey        string                                     `json:"logpush_public_key" api:"nullable"`
	Otel                    []AIGatewayListResponseOtel                `json:"otel" api:"nullable"`
	RateLimitingTechnique   AIGatewayListResponseRateLimitingTechnique `json:"rate_limiting_technique" api:"nullable"`
	// Backoff strategy for retry delays
	RetryBackoff AIGatewayListResponseRetryBackoff `json:"retry_backoff" api:"nullable"`
	// Delay between retry attempts in milliseconds (0-5000)
	RetryDelay int64 `json:"retry_delay" api:"nullable"`
	// Maximum number of retry attempts for failed requests (1-5)
	RetryMaxAttempts int64                       `json:"retry_max_attempts" api:"nullable"`
	StoreID          string                      `json:"store_id" api:"nullable"`
	Stripe           AIGatewayListResponseStripe `json:"stripe" api:"nullable"`
	// Controls how Workers AI inference calls routed through this gateway are billed.
	// Only 'postpaid' is currently supported.
	WorkersAIBillingMode AIGatewayListResponseWorkersAIBillingMode `json:"workers_ai_billing_mode"`
	Zdr                  bool                                      `json:"zdr"`
	JSON                 aiGatewayListResponseJSON                 `json:"-"`
}

// aiGatewayListResponseJSON contains the JSON metadata for the struct
// [AIGatewayListResponse]
type aiGatewayListResponseJSON struct {
	ID                      apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	Authentication          apijson.Field
	DLP                     apijson.Field
	Guardrails              apijson.Field
	IsDefault               apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	Otel                    apijson.Field
	RateLimitingTechnique   apijson.Field
	RetryBackoff            apijson.Field
	RetryDelay              apijson.Field
	RetryMaxAttempts        apijson.Field
	StoreID                 apijson.Field
	Stripe                  apijson.Field
	WorkersAIBillingMode    apijson.Field
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

type AIGatewayListResponseDLP struct {
	Enabled bool                           `json:"enabled" api:"required"`
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
	Action   AIGatewayListResponseDLPObjectAction `json:"action" api:"required"`
	Enabled  bool                                 `json:"enabled" api:"required"`
	Profiles []string                             `json:"profiles" api:"required"`
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

type AIGatewayListResponseGuardrails struct {
	Prompt   AIGatewayListResponseGuardrailsPrompt   `json:"prompt" api:"required"`
	Response AIGatewayListResponseGuardrailsResponse `json:"response" api:"required"`
	JSON     aiGatewayListResponseGuardrailsJSON     `json:"-"`
}

// aiGatewayListResponseGuardrailsJSON contains the JSON metadata for the struct
// [AIGatewayListResponseGuardrails]
type aiGatewayListResponseGuardrailsJSON struct {
	Prompt      apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayListResponseGuardrails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayListResponseGuardrailsJSON) RawJSON() string {
	return r.raw
}

type AIGatewayListResponseGuardrailsPrompt struct {
	P1   AIGatewayListResponseGuardrailsPromptP1   `json:"P1"`
	S1   AIGatewayListResponseGuardrailsPromptS1   `json:"S1"`
	S10  AIGatewayListResponseGuardrailsPromptS10  `json:"S10"`
	S11  AIGatewayListResponseGuardrailsPromptS11  `json:"S11"`
	S12  AIGatewayListResponseGuardrailsPromptS12  `json:"S12"`
	S13  AIGatewayListResponseGuardrailsPromptS13  `json:"S13"`
	S2   AIGatewayListResponseGuardrailsPromptS2   `json:"S2"`
	S3   AIGatewayListResponseGuardrailsPromptS3   `json:"S3"`
	S4   AIGatewayListResponseGuardrailsPromptS4   `json:"S4"`
	S5   AIGatewayListResponseGuardrailsPromptS5   `json:"S5"`
	S6   AIGatewayListResponseGuardrailsPromptS6   `json:"S6"`
	S7   AIGatewayListResponseGuardrailsPromptS7   `json:"S7"`
	S8   AIGatewayListResponseGuardrailsPromptS8   `json:"S8"`
	S9   AIGatewayListResponseGuardrailsPromptS9   `json:"S9"`
	JSON aiGatewayListResponseGuardrailsPromptJSON `json:"-"`
}

// aiGatewayListResponseGuardrailsPromptJSON contains the JSON metadata for the
// struct [AIGatewayListResponseGuardrailsPrompt]
type aiGatewayListResponseGuardrailsPromptJSON struct {
	P1          apijson.Field
	S1          apijson.Field
	S10         apijson.Field
	S11         apijson.Field
	S12         apijson.Field
	S13         apijson.Field
	S2          apijson.Field
	S3          apijson.Field
	S4          apijson.Field
	S5          apijson.Field
	S6          apijson.Field
	S7          apijson.Field
	S8          apijson.Field
	S9          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayListResponseGuardrailsPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayListResponseGuardrailsPromptJSON) RawJSON() string {
	return r.raw
}

type AIGatewayListResponseGuardrailsPromptP1 string

const (
	AIGatewayListResponseGuardrailsPromptP1Flag  AIGatewayListResponseGuardrailsPromptP1 = "FLAG"
	AIGatewayListResponseGuardrailsPromptP1Block AIGatewayListResponseGuardrailsPromptP1 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptP1) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptP1Flag, AIGatewayListResponseGuardrailsPromptP1Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsPromptS1 string

const (
	AIGatewayListResponseGuardrailsPromptS1Flag  AIGatewayListResponseGuardrailsPromptS1 = "FLAG"
	AIGatewayListResponseGuardrailsPromptS1Block AIGatewayListResponseGuardrailsPromptS1 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptS1) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptS1Flag, AIGatewayListResponseGuardrailsPromptS1Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsPromptS10 string

const (
	AIGatewayListResponseGuardrailsPromptS10Flag  AIGatewayListResponseGuardrailsPromptS10 = "FLAG"
	AIGatewayListResponseGuardrailsPromptS10Block AIGatewayListResponseGuardrailsPromptS10 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptS10) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptS10Flag, AIGatewayListResponseGuardrailsPromptS10Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsPromptS11 string

const (
	AIGatewayListResponseGuardrailsPromptS11Flag  AIGatewayListResponseGuardrailsPromptS11 = "FLAG"
	AIGatewayListResponseGuardrailsPromptS11Block AIGatewayListResponseGuardrailsPromptS11 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptS11) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptS11Flag, AIGatewayListResponseGuardrailsPromptS11Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsPromptS12 string

const (
	AIGatewayListResponseGuardrailsPromptS12Flag  AIGatewayListResponseGuardrailsPromptS12 = "FLAG"
	AIGatewayListResponseGuardrailsPromptS12Block AIGatewayListResponseGuardrailsPromptS12 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptS12) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptS12Flag, AIGatewayListResponseGuardrailsPromptS12Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsPromptS13 string

const (
	AIGatewayListResponseGuardrailsPromptS13Flag  AIGatewayListResponseGuardrailsPromptS13 = "FLAG"
	AIGatewayListResponseGuardrailsPromptS13Block AIGatewayListResponseGuardrailsPromptS13 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptS13) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptS13Flag, AIGatewayListResponseGuardrailsPromptS13Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsPromptS2 string

const (
	AIGatewayListResponseGuardrailsPromptS2Flag  AIGatewayListResponseGuardrailsPromptS2 = "FLAG"
	AIGatewayListResponseGuardrailsPromptS2Block AIGatewayListResponseGuardrailsPromptS2 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptS2) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptS2Flag, AIGatewayListResponseGuardrailsPromptS2Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsPromptS3 string

const (
	AIGatewayListResponseGuardrailsPromptS3Flag  AIGatewayListResponseGuardrailsPromptS3 = "FLAG"
	AIGatewayListResponseGuardrailsPromptS3Block AIGatewayListResponseGuardrailsPromptS3 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptS3) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptS3Flag, AIGatewayListResponseGuardrailsPromptS3Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsPromptS4 string

const (
	AIGatewayListResponseGuardrailsPromptS4Flag  AIGatewayListResponseGuardrailsPromptS4 = "FLAG"
	AIGatewayListResponseGuardrailsPromptS4Block AIGatewayListResponseGuardrailsPromptS4 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptS4) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptS4Flag, AIGatewayListResponseGuardrailsPromptS4Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsPromptS5 string

const (
	AIGatewayListResponseGuardrailsPromptS5Flag  AIGatewayListResponseGuardrailsPromptS5 = "FLAG"
	AIGatewayListResponseGuardrailsPromptS5Block AIGatewayListResponseGuardrailsPromptS5 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptS5) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptS5Flag, AIGatewayListResponseGuardrailsPromptS5Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsPromptS6 string

const (
	AIGatewayListResponseGuardrailsPromptS6Flag  AIGatewayListResponseGuardrailsPromptS6 = "FLAG"
	AIGatewayListResponseGuardrailsPromptS6Block AIGatewayListResponseGuardrailsPromptS6 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptS6) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptS6Flag, AIGatewayListResponseGuardrailsPromptS6Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsPromptS7 string

const (
	AIGatewayListResponseGuardrailsPromptS7Flag  AIGatewayListResponseGuardrailsPromptS7 = "FLAG"
	AIGatewayListResponseGuardrailsPromptS7Block AIGatewayListResponseGuardrailsPromptS7 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptS7) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptS7Flag, AIGatewayListResponseGuardrailsPromptS7Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsPromptS8 string

const (
	AIGatewayListResponseGuardrailsPromptS8Flag  AIGatewayListResponseGuardrailsPromptS8 = "FLAG"
	AIGatewayListResponseGuardrailsPromptS8Block AIGatewayListResponseGuardrailsPromptS8 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptS8) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptS8Flag, AIGatewayListResponseGuardrailsPromptS8Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsPromptS9 string

const (
	AIGatewayListResponseGuardrailsPromptS9Flag  AIGatewayListResponseGuardrailsPromptS9 = "FLAG"
	AIGatewayListResponseGuardrailsPromptS9Block AIGatewayListResponseGuardrailsPromptS9 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsPromptS9) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsPromptS9Flag, AIGatewayListResponseGuardrailsPromptS9Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponse struct {
	P1   AIGatewayListResponseGuardrailsResponseP1   `json:"P1"`
	S1   AIGatewayListResponseGuardrailsResponseS1   `json:"S1"`
	S10  AIGatewayListResponseGuardrailsResponseS10  `json:"S10"`
	S11  AIGatewayListResponseGuardrailsResponseS11  `json:"S11"`
	S12  AIGatewayListResponseGuardrailsResponseS12  `json:"S12"`
	S13  AIGatewayListResponseGuardrailsResponseS13  `json:"S13"`
	S2   AIGatewayListResponseGuardrailsResponseS2   `json:"S2"`
	S3   AIGatewayListResponseGuardrailsResponseS3   `json:"S3"`
	S4   AIGatewayListResponseGuardrailsResponseS4   `json:"S4"`
	S5   AIGatewayListResponseGuardrailsResponseS5   `json:"S5"`
	S6   AIGatewayListResponseGuardrailsResponseS6   `json:"S6"`
	S7   AIGatewayListResponseGuardrailsResponseS7   `json:"S7"`
	S8   AIGatewayListResponseGuardrailsResponseS8   `json:"S8"`
	S9   AIGatewayListResponseGuardrailsResponseS9   `json:"S9"`
	JSON aiGatewayListResponseGuardrailsResponseJSON `json:"-"`
}

// aiGatewayListResponseGuardrailsResponseJSON contains the JSON metadata for the
// struct [AIGatewayListResponseGuardrailsResponse]
type aiGatewayListResponseGuardrailsResponseJSON struct {
	P1          apijson.Field
	S1          apijson.Field
	S10         apijson.Field
	S11         apijson.Field
	S12         apijson.Field
	S13         apijson.Field
	S2          apijson.Field
	S3          apijson.Field
	S4          apijson.Field
	S5          apijson.Field
	S6          apijson.Field
	S7          apijson.Field
	S8          apijson.Field
	S9          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayListResponseGuardrailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayListResponseGuardrailsResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayListResponseGuardrailsResponseP1 string

const (
	AIGatewayListResponseGuardrailsResponseP1Flag  AIGatewayListResponseGuardrailsResponseP1 = "FLAG"
	AIGatewayListResponseGuardrailsResponseP1Block AIGatewayListResponseGuardrailsResponseP1 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseP1) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseP1Flag, AIGatewayListResponseGuardrailsResponseP1Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponseS1 string

const (
	AIGatewayListResponseGuardrailsResponseS1Flag  AIGatewayListResponseGuardrailsResponseS1 = "FLAG"
	AIGatewayListResponseGuardrailsResponseS1Block AIGatewayListResponseGuardrailsResponseS1 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseS1) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseS1Flag, AIGatewayListResponseGuardrailsResponseS1Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponseS10 string

const (
	AIGatewayListResponseGuardrailsResponseS10Flag  AIGatewayListResponseGuardrailsResponseS10 = "FLAG"
	AIGatewayListResponseGuardrailsResponseS10Block AIGatewayListResponseGuardrailsResponseS10 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseS10) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseS10Flag, AIGatewayListResponseGuardrailsResponseS10Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponseS11 string

const (
	AIGatewayListResponseGuardrailsResponseS11Flag  AIGatewayListResponseGuardrailsResponseS11 = "FLAG"
	AIGatewayListResponseGuardrailsResponseS11Block AIGatewayListResponseGuardrailsResponseS11 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseS11) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseS11Flag, AIGatewayListResponseGuardrailsResponseS11Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponseS12 string

const (
	AIGatewayListResponseGuardrailsResponseS12Flag  AIGatewayListResponseGuardrailsResponseS12 = "FLAG"
	AIGatewayListResponseGuardrailsResponseS12Block AIGatewayListResponseGuardrailsResponseS12 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseS12) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseS12Flag, AIGatewayListResponseGuardrailsResponseS12Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponseS13 string

const (
	AIGatewayListResponseGuardrailsResponseS13Flag  AIGatewayListResponseGuardrailsResponseS13 = "FLAG"
	AIGatewayListResponseGuardrailsResponseS13Block AIGatewayListResponseGuardrailsResponseS13 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseS13) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseS13Flag, AIGatewayListResponseGuardrailsResponseS13Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponseS2 string

const (
	AIGatewayListResponseGuardrailsResponseS2Flag  AIGatewayListResponseGuardrailsResponseS2 = "FLAG"
	AIGatewayListResponseGuardrailsResponseS2Block AIGatewayListResponseGuardrailsResponseS2 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseS2) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseS2Flag, AIGatewayListResponseGuardrailsResponseS2Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponseS3 string

const (
	AIGatewayListResponseGuardrailsResponseS3Flag  AIGatewayListResponseGuardrailsResponseS3 = "FLAG"
	AIGatewayListResponseGuardrailsResponseS3Block AIGatewayListResponseGuardrailsResponseS3 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseS3) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseS3Flag, AIGatewayListResponseGuardrailsResponseS3Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponseS4 string

const (
	AIGatewayListResponseGuardrailsResponseS4Flag  AIGatewayListResponseGuardrailsResponseS4 = "FLAG"
	AIGatewayListResponseGuardrailsResponseS4Block AIGatewayListResponseGuardrailsResponseS4 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseS4) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseS4Flag, AIGatewayListResponseGuardrailsResponseS4Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponseS5 string

const (
	AIGatewayListResponseGuardrailsResponseS5Flag  AIGatewayListResponseGuardrailsResponseS5 = "FLAG"
	AIGatewayListResponseGuardrailsResponseS5Block AIGatewayListResponseGuardrailsResponseS5 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseS5) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseS5Flag, AIGatewayListResponseGuardrailsResponseS5Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponseS6 string

const (
	AIGatewayListResponseGuardrailsResponseS6Flag  AIGatewayListResponseGuardrailsResponseS6 = "FLAG"
	AIGatewayListResponseGuardrailsResponseS6Block AIGatewayListResponseGuardrailsResponseS6 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseS6) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseS6Flag, AIGatewayListResponseGuardrailsResponseS6Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponseS7 string

const (
	AIGatewayListResponseGuardrailsResponseS7Flag  AIGatewayListResponseGuardrailsResponseS7 = "FLAG"
	AIGatewayListResponseGuardrailsResponseS7Block AIGatewayListResponseGuardrailsResponseS7 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseS7) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseS7Flag, AIGatewayListResponseGuardrailsResponseS7Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponseS8 string

const (
	AIGatewayListResponseGuardrailsResponseS8Flag  AIGatewayListResponseGuardrailsResponseS8 = "FLAG"
	AIGatewayListResponseGuardrailsResponseS8Block AIGatewayListResponseGuardrailsResponseS8 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseS8) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseS8Flag, AIGatewayListResponseGuardrailsResponseS8Block:
		return true
	}
	return false
}

type AIGatewayListResponseGuardrailsResponseS9 string

const (
	AIGatewayListResponseGuardrailsResponseS9Flag  AIGatewayListResponseGuardrailsResponseS9 = "FLAG"
	AIGatewayListResponseGuardrailsResponseS9Block AIGatewayListResponseGuardrailsResponseS9 = "BLOCK"
)

func (r AIGatewayListResponseGuardrailsResponseS9) IsKnown() bool {
	switch r {
	case AIGatewayListResponseGuardrailsResponseS9Flag, AIGatewayListResponseGuardrailsResponseS9Block:
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
	Authorization string                               `json:"authorization" api:"required"`
	Headers       map[string]string                    `json:"headers" api:"required"`
	URL           string                               `json:"url" api:"required" format:"uri"`
	ContentType   AIGatewayListResponseOtelContentType `json:"content_type"`
	JSON          aiGatewayListResponseOtelJSON        `json:"-"`
}

// aiGatewayListResponseOtelJSON contains the JSON metadata for the struct
// [AIGatewayListResponseOtel]
type aiGatewayListResponseOtelJSON struct {
	Authorization apijson.Field
	Headers       apijson.Field
	URL           apijson.Field
	ContentType   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayListResponseOtel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayListResponseOtelJSON) RawJSON() string {
	return r.raw
}

type AIGatewayListResponseOtelContentType string

const (
	AIGatewayListResponseOtelContentTypeJson     AIGatewayListResponseOtelContentType = "json"
	AIGatewayListResponseOtelContentTypeProtobuf AIGatewayListResponseOtelContentType = "protobuf"
)

func (r AIGatewayListResponseOtelContentType) IsKnown() bool {
	switch r {
	case AIGatewayListResponseOtelContentTypeJson, AIGatewayListResponseOtelContentTypeProtobuf:
		return true
	}
	return false
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

// Backoff strategy for retry delays
type AIGatewayListResponseRetryBackoff string

const (
	AIGatewayListResponseRetryBackoffConstant    AIGatewayListResponseRetryBackoff = "constant"
	AIGatewayListResponseRetryBackoffLinear      AIGatewayListResponseRetryBackoff = "linear"
	AIGatewayListResponseRetryBackoffExponential AIGatewayListResponseRetryBackoff = "exponential"
)

func (r AIGatewayListResponseRetryBackoff) IsKnown() bool {
	switch r {
	case AIGatewayListResponseRetryBackoffConstant, AIGatewayListResponseRetryBackoffLinear, AIGatewayListResponseRetryBackoffExponential:
		return true
	}
	return false
}

type AIGatewayListResponseStripe struct {
	Authorization string                                  `json:"authorization" api:"required"`
	UsageEvents   []AIGatewayListResponseStripeUsageEvent `json:"usage_events" api:"required"`
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
	Payload string                                    `json:"payload" api:"required"`
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

// Controls how Workers AI inference calls routed through this gateway are billed.
// Only 'postpaid' is currently supported.
type AIGatewayListResponseWorkersAIBillingMode string

const (
	AIGatewayListResponseWorkersAIBillingModePostpaid AIGatewayListResponseWorkersAIBillingMode = "postpaid"
)

func (r AIGatewayListResponseWorkersAIBillingMode) IsKnown() bool {
	switch r {
	case AIGatewayListResponseWorkersAIBillingModePostpaid:
		return true
	}
	return false
}

type AIGatewayDeleteResponse struct {
	// gateway id
	ID                      string                                       `json:"id" api:"required"`
	CacheInvalidateOnUpdate bool                                         `json:"cache_invalidate_on_update" api:"required"`
	CacheTTL                int64                                        `json:"cache_ttl" api:"required,nullable"`
	CollectLogs             bool                                         `json:"collect_logs" api:"required"`
	CreatedAt               time.Time                                    `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt              time.Time                                    `json:"modified_at" api:"required" format:"date-time"`
	RateLimitingInterval    int64                                        `json:"rate_limiting_interval" api:"required,nullable"`
	RateLimitingLimit       int64                                        `json:"rate_limiting_limit" api:"required,nullable"`
	Authentication          bool                                         `json:"authentication"`
	DLP                     AIGatewayDeleteResponseDLP                   `json:"dlp"`
	Guardrails              AIGatewayDeleteResponseGuardrails            `json:"guardrails" api:"nullable"`
	IsDefault               bool                                         `json:"is_default"`
	LogManagement           int64                                        `json:"log_management" api:"nullable"`
	LogManagementStrategy   AIGatewayDeleteResponseLogManagementStrategy `json:"log_management_strategy" api:"nullable"`
	Logpush                 bool                                         `json:"logpush"`
	LogpushPublicKey        string                                       `json:"logpush_public_key" api:"nullable"`
	Otel                    []AIGatewayDeleteResponseOtel                `json:"otel" api:"nullable"`
	RateLimitingTechnique   AIGatewayDeleteResponseRateLimitingTechnique `json:"rate_limiting_technique" api:"nullable"`
	// Backoff strategy for retry delays
	RetryBackoff AIGatewayDeleteResponseRetryBackoff `json:"retry_backoff" api:"nullable"`
	// Delay between retry attempts in milliseconds (0-5000)
	RetryDelay int64 `json:"retry_delay" api:"nullable"`
	// Maximum number of retry attempts for failed requests (1-5)
	RetryMaxAttempts int64                         `json:"retry_max_attempts" api:"nullable"`
	StoreID          string                        `json:"store_id" api:"nullable"`
	Stripe           AIGatewayDeleteResponseStripe `json:"stripe" api:"nullable"`
	// Controls how Workers AI inference calls routed through this gateway are billed.
	// Only 'postpaid' is currently supported.
	WorkersAIBillingMode AIGatewayDeleteResponseWorkersAIBillingMode `json:"workers_ai_billing_mode"`
	Zdr                  bool                                        `json:"zdr"`
	JSON                 aiGatewayDeleteResponseJSON                 `json:"-"`
}

// aiGatewayDeleteResponseJSON contains the JSON metadata for the struct
// [AIGatewayDeleteResponse]
type aiGatewayDeleteResponseJSON struct {
	ID                      apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	Authentication          apijson.Field
	DLP                     apijson.Field
	Guardrails              apijson.Field
	IsDefault               apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	Otel                    apijson.Field
	RateLimitingTechnique   apijson.Field
	RetryBackoff            apijson.Field
	RetryDelay              apijson.Field
	RetryMaxAttempts        apijson.Field
	StoreID                 apijson.Field
	Stripe                  apijson.Field
	WorkersAIBillingMode    apijson.Field
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

type AIGatewayDeleteResponseDLP struct {
	Enabled bool                             `json:"enabled" api:"required"`
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
	Action   AIGatewayDeleteResponseDLPObjectAction `json:"action" api:"required"`
	Enabled  bool                                   `json:"enabled" api:"required"`
	Profiles []string                               `json:"profiles" api:"required"`
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

type AIGatewayDeleteResponseGuardrails struct {
	Prompt   AIGatewayDeleteResponseGuardrailsPrompt   `json:"prompt" api:"required"`
	Response AIGatewayDeleteResponseGuardrailsResponse `json:"response" api:"required"`
	JSON     aiGatewayDeleteResponseGuardrailsJSON     `json:"-"`
}

// aiGatewayDeleteResponseGuardrailsJSON contains the JSON metadata for the struct
// [AIGatewayDeleteResponseGuardrails]
type aiGatewayDeleteResponseGuardrailsJSON struct {
	Prompt      apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayDeleteResponseGuardrails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseGuardrailsJSON) RawJSON() string {
	return r.raw
}

type AIGatewayDeleteResponseGuardrailsPrompt struct {
	P1   AIGatewayDeleteResponseGuardrailsPromptP1   `json:"P1"`
	S1   AIGatewayDeleteResponseGuardrailsPromptS1   `json:"S1"`
	S10  AIGatewayDeleteResponseGuardrailsPromptS10  `json:"S10"`
	S11  AIGatewayDeleteResponseGuardrailsPromptS11  `json:"S11"`
	S12  AIGatewayDeleteResponseGuardrailsPromptS12  `json:"S12"`
	S13  AIGatewayDeleteResponseGuardrailsPromptS13  `json:"S13"`
	S2   AIGatewayDeleteResponseGuardrailsPromptS2   `json:"S2"`
	S3   AIGatewayDeleteResponseGuardrailsPromptS3   `json:"S3"`
	S4   AIGatewayDeleteResponseGuardrailsPromptS4   `json:"S4"`
	S5   AIGatewayDeleteResponseGuardrailsPromptS5   `json:"S5"`
	S6   AIGatewayDeleteResponseGuardrailsPromptS6   `json:"S6"`
	S7   AIGatewayDeleteResponseGuardrailsPromptS7   `json:"S7"`
	S8   AIGatewayDeleteResponseGuardrailsPromptS8   `json:"S8"`
	S9   AIGatewayDeleteResponseGuardrailsPromptS9   `json:"S9"`
	JSON aiGatewayDeleteResponseGuardrailsPromptJSON `json:"-"`
}

// aiGatewayDeleteResponseGuardrailsPromptJSON contains the JSON metadata for the
// struct [AIGatewayDeleteResponseGuardrailsPrompt]
type aiGatewayDeleteResponseGuardrailsPromptJSON struct {
	P1          apijson.Field
	S1          apijson.Field
	S10         apijson.Field
	S11         apijson.Field
	S12         apijson.Field
	S13         apijson.Field
	S2          apijson.Field
	S3          apijson.Field
	S4          apijson.Field
	S5          apijson.Field
	S6          apijson.Field
	S7          apijson.Field
	S8          apijson.Field
	S9          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayDeleteResponseGuardrailsPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseGuardrailsPromptJSON) RawJSON() string {
	return r.raw
}

type AIGatewayDeleteResponseGuardrailsPromptP1 string

const (
	AIGatewayDeleteResponseGuardrailsPromptP1Flag  AIGatewayDeleteResponseGuardrailsPromptP1 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptP1Block AIGatewayDeleteResponseGuardrailsPromptP1 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptP1) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptP1Flag, AIGatewayDeleteResponseGuardrailsPromptP1Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsPromptS1 string

const (
	AIGatewayDeleteResponseGuardrailsPromptS1Flag  AIGatewayDeleteResponseGuardrailsPromptS1 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptS1Block AIGatewayDeleteResponseGuardrailsPromptS1 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptS1) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptS1Flag, AIGatewayDeleteResponseGuardrailsPromptS1Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsPromptS10 string

const (
	AIGatewayDeleteResponseGuardrailsPromptS10Flag  AIGatewayDeleteResponseGuardrailsPromptS10 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptS10Block AIGatewayDeleteResponseGuardrailsPromptS10 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptS10) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptS10Flag, AIGatewayDeleteResponseGuardrailsPromptS10Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsPromptS11 string

const (
	AIGatewayDeleteResponseGuardrailsPromptS11Flag  AIGatewayDeleteResponseGuardrailsPromptS11 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptS11Block AIGatewayDeleteResponseGuardrailsPromptS11 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptS11) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptS11Flag, AIGatewayDeleteResponseGuardrailsPromptS11Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsPromptS12 string

const (
	AIGatewayDeleteResponseGuardrailsPromptS12Flag  AIGatewayDeleteResponseGuardrailsPromptS12 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptS12Block AIGatewayDeleteResponseGuardrailsPromptS12 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptS12) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptS12Flag, AIGatewayDeleteResponseGuardrailsPromptS12Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsPromptS13 string

const (
	AIGatewayDeleteResponseGuardrailsPromptS13Flag  AIGatewayDeleteResponseGuardrailsPromptS13 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptS13Block AIGatewayDeleteResponseGuardrailsPromptS13 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptS13) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptS13Flag, AIGatewayDeleteResponseGuardrailsPromptS13Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsPromptS2 string

const (
	AIGatewayDeleteResponseGuardrailsPromptS2Flag  AIGatewayDeleteResponseGuardrailsPromptS2 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptS2Block AIGatewayDeleteResponseGuardrailsPromptS2 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptS2) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptS2Flag, AIGatewayDeleteResponseGuardrailsPromptS2Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsPromptS3 string

const (
	AIGatewayDeleteResponseGuardrailsPromptS3Flag  AIGatewayDeleteResponseGuardrailsPromptS3 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptS3Block AIGatewayDeleteResponseGuardrailsPromptS3 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptS3) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptS3Flag, AIGatewayDeleteResponseGuardrailsPromptS3Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsPromptS4 string

const (
	AIGatewayDeleteResponseGuardrailsPromptS4Flag  AIGatewayDeleteResponseGuardrailsPromptS4 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptS4Block AIGatewayDeleteResponseGuardrailsPromptS4 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptS4) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptS4Flag, AIGatewayDeleteResponseGuardrailsPromptS4Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsPromptS5 string

const (
	AIGatewayDeleteResponseGuardrailsPromptS5Flag  AIGatewayDeleteResponseGuardrailsPromptS5 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptS5Block AIGatewayDeleteResponseGuardrailsPromptS5 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptS5) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptS5Flag, AIGatewayDeleteResponseGuardrailsPromptS5Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsPromptS6 string

const (
	AIGatewayDeleteResponseGuardrailsPromptS6Flag  AIGatewayDeleteResponseGuardrailsPromptS6 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptS6Block AIGatewayDeleteResponseGuardrailsPromptS6 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptS6) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptS6Flag, AIGatewayDeleteResponseGuardrailsPromptS6Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsPromptS7 string

const (
	AIGatewayDeleteResponseGuardrailsPromptS7Flag  AIGatewayDeleteResponseGuardrailsPromptS7 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptS7Block AIGatewayDeleteResponseGuardrailsPromptS7 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptS7) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptS7Flag, AIGatewayDeleteResponseGuardrailsPromptS7Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsPromptS8 string

const (
	AIGatewayDeleteResponseGuardrailsPromptS8Flag  AIGatewayDeleteResponseGuardrailsPromptS8 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptS8Block AIGatewayDeleteResponseGuardrailsPromptS8 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptS8) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptS8Flag, AIGatewayDeleteResponseGuardrailsPromptS8Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsPromptS9 string

const (
	AIGatewayDeleteResponseGuardrailsPromptS9Flag  AIGatewayDeleteResponseGuardrailsPromptS9 = "FLAG"
	AIGatewayDeleteResponseGuardrailsPromptS9Block AIGatewayDeleteResponseGuardrailsPromptS9 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsPromptS9) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsPromptS9Flag, AIGatewayDeleteResponseGuardrailsPromptS9Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponse struct {
	P1   AIGatewayDeleteResponseGuardrailsResponseP1   `json:"P1"`
	S1   AIGatewayDeleteResponseGuardrailsResponseS1   `json:"S1"`
	S10  AIGatewayDeleteResponseGuardrailsResponseS10  `json:"S10"`
	S11  AIGatewayDeleteResponseGuardrailsResponseS11  `json:"S11"`
	S12  AIGatewayDeleteResponseGuardrailsResponseS12  `json:"S12"`
	S13  AIGatewayDeleteResponseGuardrailsResponseS13  `json:"S13"`
	S2   AIGatewayDeleteResponseGuardrailsResponseS2   `json:"S2"`
	S3   AIGatewayDeleteResponseGuardrailsResponseS3   `json:"S3"`
	S4   AIGatewayDeleteResponseGuardrailsResponseS4   `json:"S4"`
	S5   AIGatewayDeleteResponseGuardrailsResponseS5   `json:"S5"`
	S6   AIGatewayDeleteResponseGuardrailsResponseS6   `json:"S6"`
	S7   AIGatewayDeleteResponseGuardrailsResponseS7   `json:"S7"`
	S8   AIGatewayDeleteResponseGuardrailsResponseS8   `json:"S8"`
	S9   AIGatewayDeleteResponseGuardrailsResponseS9   `json:"S9"`
	JSON aiGatewayDeleteResponseGuardrailsResponseJSON `json:"-"`
}

// aiGatewayDeleteResponseGuardrailsResponseJSON contains the JSON metadata for the
// struct [AIGatewayDeleteResponseGuardrailsResponse]
type aiGatewayDeleteResponseGuardrailsResponseJSON struct {
	P1          apijson.Field
	S1          apijson.Field
	S10         apijson.Field
	S11         apijson.Field
	S12         apijson.Field
	S13         apijson.Field
	S2          apijson.Field
	S3          apijson.Field
	S4          apijson.Field
	S5          apijson.Field
	S6          apijson.Field
	S7          apijson.Field
	S8          apijson.Field
	S9          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayDeleteResponseGuardrailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseGuardrailsResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayDeleteResponseGuardrailsResponseP1 string

const (
	AIGatewayDeleteResponseGuardrailsResponseP1Flag  AIGatewayDeleteResponseGuardrailsResponseP1 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseP1Block AIGatewayDeleteResponseGuardrailsResponseP1 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseP1) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseP1Flag, AIGatewayDeleteResponseGuardrailsResponseP1Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponseS1 string

const (
	AIGatewayDeleteResponseGuardrailsResponseS1Flag  AIGatewayDeleteResponseGuardrailsResponseS1 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseS1Block AIGatewayDeleteResponseGuardrailsResponseS1 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseS1) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseS1Flag, AIGatewayDeleteResponseGuardrailsResponseS1Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponseS10 string

const (
	AIGatewayDeleteResponseGuardrailsResponseS10Flag  AIGatewayDeleteResponseGuardrailsResponseS10 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseS10Block AIGatewayDeleteResponseGuardrailsResponseS10 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseS10) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseS10Flag, AIGatewayDeleteResponseGuardrailsResponseS10Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponseS11 string

const (
	AIGatewayDeleteResponseGuardrailsResponseS11Flag  AIGatewayDeleteResponseGuardrailsResponseS11 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseS11Block AIGatewayDeleteResponseGuardrailsResponseS11 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseS11) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseS11Flag, AIGatewayDeleteResponseGuardrailsResponseS11Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponseS12 string

const (
	AIGatewayDeleteResponseGuardrailsResponseS12Flag  AIGatewayDeleteResponseGuardrailsResponseS12 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseS12Block AIGatewayDeleteResponseGuardrailsResponseS12 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseS12) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseS12Flag, AIGatewayDeleteResponseGuardrailsResponseS12Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponseS13 string

const (
	AIGatewayDeleteResponseGuardrailsResponseS13Flag  AIGatewayDeleteResponseGuardrailsResponseS13 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseS13Block AIGatewayDeleteResponseGuardrailsResponseS13 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseS13) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseS13Flag, AIGatewayDeleteResponseGuardrailsResponseS13Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponseS2 string

const (
	AIGatewayDeleteResponseGuardrailsResponseS2Flag  AIGatewayDeleteResponseGuardrailsResponseS2 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseS2Block AIGatewayDeleteResponseGuardrailsResponseS2 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseS2) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseS2Flag, AIGatewayDeleteResponseGuardrailsResponseS2Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponseS3 string

const (
	AIGatewayDeleteResponseGuardrailsResponseS3Flag  AIGatewayDeleteResponseGuardrailsResponseS3 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseS3Block AIGatewayDeleteResponseGuardrailsResponseS3 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseS3) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseS3Flag, AIGatewayDeleteResponseGuardrailsResponseS3Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponseS4 string

const (
	AIGatewayDeleteResponseGuardrailsResponseS4Flag  AIGatewayDeleteResponseGuardrailsResponseS4 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseS4Block AIGatewayDeleteResponseGuardrailsResponseS4 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseS4) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseS4Flag, AIGatewayDeleteResponseGuardrailsResponseS4Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponseS5 string

const (
	AIGatewayDeleteResponseGuardrailsResponseS5Flag  AIGatewayDeleteResponseGuardrailsResponseS5 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseS5Block AIGatewayDeleteResponseGuardrailsResponseS5 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseS5) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseS5Flag, AIGatewayDeleteResponseGuardrailsResponseS5Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponseS6 string

const (
	AIGatewayDeleteResponseGuardrailsResponseS6Flag  AIGatewayDeleteResponseGuardrailsResponseS6 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseS6Block AIGatewayDeleteResponseGuardrailsResponseS6 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseS6) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseS6Flag, AIGatewayDeleteResponseGuardrailsResponseS6Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponseS7 string

const (
	AIGatewayDeleteResponseGuardrailsResponseS7Flag  AIGatewayDeleteResponseGuardrailsResponseS7 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseS7Block AIGatewayDeleteResponseGuardrailsResponseS7 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseS7) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseS7Flag, AIGatewayDeleteResponseGuardrailsResponseS7Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponseS8 string

const (
	AIGatewayDeleteResponseGuardrailsResponseS8Flag  AIGatewayDeleteResponseGuardrailsResponseS8 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseS8Block AIGatewayDeleteResponseGuardrailsResponseS8 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseS8) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseS8Flag, AIGatewayDeleteResponseGuardrailsResponseS8Block:
		return true
	}
	return false
}

type AIGatewayDeleteResponseGuardrailsResponseS9 string

const (
	AIGatewayDeleteResponseGuardrailsResponseS9Flag  AIGatewayDeleteResponseGuardrailsResponseS9 = "FLAG"
	AIGatewayDeleteResponseGuardrailsResponseS9Block AIGatewayDeleteResponseGuardrailsResponseS9 = "BLOCK"
)

func (r AIGatewayDeleteResponseGuardrailsResponseS9) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseGuardrailsResponseS9Flag, AIGatewayDeleteResponseGuardrailsResponseS9Block:
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
	Authorization string                                 `json:"authorization" api:"required"`
	Headers       map[string]string                      `json:"headers" api:"required"`
	URL           string                                 `json:"url" api:"required" format:"uri"`
	ContentType   AIGatewayDeleteResponseOtelContentType `json:"content_type"`
	JSON          aiGatewayDeleteResponseOtelJSON        `json:"-"`
}

// aiGatewayDeleteResponseOtelJSON contains the JSON metadata for the struct
// [AIGatewayDeleteResponseOtel]
type aiGatewayDeleteResponseOtelJSON struct {
	Authorization apijson.Field
	Headers       apijson.Field
	URL           apijson.Field
	ContentType   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayDeleteResponseOtel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseOtelJSON) RawJSON() string {
	return r.raw
}

type AIGatewayDeleteResponseOtelContentType string

const (
	AIGatewayDeleteResponseOtelContentTypeJson     AIGatewayDeleteResponseOtelContentType = "json"
	AIGatewayDeleteResponseOtelContentTypeProtobuf AIGatewayDeleteResponseOtelContentType = "protobuf"
)

func (r AIGatewayDeleteResponseOtelContentType) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseOtelContentTypeJson, AIGatewayDeleteResponseOtelContentTypeProtobuf:
		return true
	}
	return false
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

// Backoff strategy for retry delays
type AIGatewayDeleteResponseRetryBackoff string

const (
	AIGatewayDeleteResponseRetryBackoffConstant    AIGatewayDeleteResponseRetryBackoff = "constant"
	AIGatewayDeleteResponseRetryBackoffLinear      AIGatewayDeleteResponseRetryBackoff = "linear"
	AIGatewayDeleteResponseRetryBackoffExponential AIGatewayDeleteResponseRetryBackoff = "exponential"
)

func (r AIGatewayDeleteResponseRetryBackoff) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseRetryBackoffConstant, AIGatewayDeleteResponseRetryBackoffLinear, AIGatewayDeleteResponseRetryBackoffExponential:
		return true
	}
	return false
}

type AIGatewayDeleteResponseStripe struct {
	Authorization string                                    `json:"authorization" api:"required"`
	UsageEvents   []AIGatewayDeleteResponseStripeUsageEvent `json:"usage_events" api:"required"`
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
	Payload string                                      `json:"payload" api:"required"`
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

// Controls how Workers AI inference calls routed through this gateway are billed.
// Only 'postpaid' is currently supported.
type AIGatewayDeleteResponseWorkersAIBillingMode string

const (
	AIGatewayDeleteResponseWorkersAIBillingModePostpaid AIGatewayDeleteResponseWorkersAIBillingMode = "postpaid"
)

func (r AIGatewayDeleteResponseWorkersAIBillingMode) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseWorkersAIBillingModePostpaid:
		return true
	}
	return false
}

type AIGatewayGetResponse struct {
	// gateway id
	ID                      string                                    `json:"id" api:"required"`
	CacheInvalidateOnUpdate bool                                      `json:"cache_invalidate_on_update" api:"required"`
	CacheTTL                int64                                     `json:"cache_ttl" api:"required,nullable"`
	CollectLogs             bool                                      `json:"collect_logs" api:"required"`
	CreatedAt               time.Time                                 `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt              time.Time                                 `json:"modified_at" api:"required" format:"date-time"`
	RateLimitingInterval    int64                                     `json:"rate_limiting_interval" api:"required,nullable"`
	RateLimitingLimit       int64                                     `json:"rate_limiting_limit" api:"required,nullable"`
	Authentication          bool                                      `json:"authentication"`
	DLP                     AIGatewayGetResponseDLP                   `json:"dlp"`
	Guardrails              AIGatewayGetResponseGuardrails            `json:"guardrails" api:"nullable"`
	IsDefault               bool                                      `json:"is_default"`
	LogManagement           int64                                     `json:"log_management" api:"nullable"`
	LogManagementStrategy   AIGatewayGetResponseLogManagementStrategy `json:"log_management_strategy" api:"nullable"`
	Logpush                 bool                                      `json:"logpush"`
	LogpushPublicKey        string                                    `json:"logpush_public_key" api:"nullable"`
	Otel                    []AIGatewayGetResponseOtel                `json:"otel" api:"nullable"`
	RateLimitingTechnique   AIGatewayGetResponseRateLimitingTechnique `json:"rate_limiting_technique" api:"nullable"`
	// Backoff strategy for retry delays
	RetryBackoff AIGatewayGetResponseRetryBackoff `json:"retry_backoff" api:"nullable"`
	// Delay between retry attempts in milliseconds (0-5000)
	RetryDelay int64 `json:"retry_delay" api:"nullable"`
	// Maximum number of retry attempts for failed requests (1-5)
	RetryMaxAttempts int64                      `json:"retry_max_attempts" api:"nullable"`
	StoreID          string                     `json:"store_id" api:"nullable"`
	Stripe           AIGatewayGetResponseStripe `json:"stripe" api:"nullable"`
	// Controls how Workers AI inference calls routed through this gateway are billed.
	// Only 'postpaid' is currently supported.
	WorkersAIBillingMode AIGatewayGetResponseWorkersAIBillingMode `json:"workers_ai_billing_mode"`
	Zdr                  bool                                     `json:"zdr"`
	JSON                 aiGatewayGetResponseJSON                 `json:"-"`
}

// aiGatewayGetResponseJSON contains the JSON metadata for the struct
// [AIGatewayGetResponse]
type aiGatewayGetResponseJSON struct {
	ID                      apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	Authentication          apijson.Field
	DLP                     apijson.Field
	Guardrails              apijson.Field
	IsDefault               apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	Otel                    apijson.Field
	RateLimitingTechnique   apijson.Field
	RetryBackoff            apijson.Field
	RetryDelay              apijson.Field
	RetryMaxAttempts        apijson.Field
	StoreID                 apijson.Field
	Stripe                  apijson.Field
	WorkersAIBillingMode    apijson.Field
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

type AIGatewayGetResponseDLP struct {
	Enabled bool                          `json:"enabled" api:"required"`
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
	Action   AIGatewayGetResponseDLPObjectAction `json:"action" api:"required"`
	Enabled  bool                                `json:"enabled" api:"required"`
	Profiles []string                            `json:"profiles" api:"required"`
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

type AIGatewayGetResponseGuardrails struct {
	Prompt   AIGatewayGetResponseGuardrailsPrompt   `json:"prompt" api:"required"`
	Response AIGatewayGetResponseGuardrailsResponse `json:"response" api:"required"`
	JSON     aiGatewayGetResponseGuardrailsJSON     `json:"-"`
}

// aiGatewayGetResponseGuardrailsJSON contains the JSON metadata for the struct
// [AIGatewayGetResponseGuardrails]
type aiGatewayGetResponseGuardrailsJSON struct {
	Prompt      apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayGetResponseGuardrails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseGuardrailsJSON) RawJSON() string {
	return r.raw
}

type AIGatewayGetResponseGuardrailsPrompt struct {
	P1   AIGatewayGetResponseGuardrailsPromptP1   `json:"P1"`
	S1   AIGatewayGetResponseGuardrailsPromptS1   `json:"S1"`
	S10  AIGatewayGetResponseGuardrailsPromptS10  `json:"S10"`
	S11  AIGatewayGetResponseGuardrailsPromptS11  `json:"S11"`
	S12  AIGatewayGetResponseGuardrailsPromptS12  `json:"S12"`
	S13  AIGatewayGetResponseGuardrailsPromptS13  `json:"S13"`
	S2   AIGatewayGetResponseGuardrailsPromptS2   `json:"S2"`
	S3   AIGatewayGetResponseGuardrailsPromptS3   `json:"S3"`
	S4   AIGatewayGetResponseGuardrailsPromptS4   `json:"S4"`
	S5   AIGatewayGetResponseGuardrailsPromptS5   `json:"S5"`
	S6   AIGatewayGetResponseGuardrailsPromptS6   `json:"S6"`
	S7   AIGatewayGetResponseGuardrailsPromptS7   `json:"S7"`
	S8   AIGatewayGetResponseGuardrailsPromptS8   `json:"S8"`
	S9   AIGatewayGetResponseGuardrailsPromptS9   `json:"S9"`
	JSON aiGatewayGetResponseGuardrailsPromptJSON `json:"-"`
}

// aiGatewayGetResponseGuardrailsPromptJSON contains the JSON metadata for the
// struct [AIGatewayGetResponseGuardrailsPrompt]
type aiGatewayGetResponseGuardrailsPromptJSON struct {
	P1          apijson.Field
	S1          apijson.Field
	S10         apijson.Field
	S11         apijson.Field
	S12         apijson.Field
	S13         apijson.Field
	S2          apijson.Field
	S3          apijson.Field
	S4          apijson.Field
	S5          apijson.Field
	S6          apijson.Field
	S7          apijson.Field
	S8          apijson.Field
	S9          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayGetResponseGuardrailsPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseGuardrailsPromptJSON) RawJSON() string {
	return r.raw
}

type AIGatewayGetResponseGuardrailsPromptP1 string

const (
	AIGatewayGetResponseGuardrailsPromptP1Flag  AIGatewayGetResponseGuardrailsPromptP1 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptP1Block AIGatewayGetResponseGuardrailsPromptP1 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptP1) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptP1Flag, AIGatewayGetResponseGuardrailsPromptP1Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsPromptS1 string

const (
	AIGatewayGetResponseGuardrailsPromptS1Flag  AIGatewayGetResponseGuardrailsPromptS1 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptS1Block AIGatewayGetResponseGuardrailsPromptS1 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptS1) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptS1Flag, AIGatewayGetResponseGuardrailsPromptS1Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsPromptS10 string

const (
	AIGatewayGetResponseGuardrailsPromptS10Flag  AIGatewayGetResponseGuardrailsPromptS10 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptS10Block AIGatewayGetResponseGuardrailsPromptS10 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptS10) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptS10Flag, AIGatewayGetResponseGuardrailsPromptS10Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsPromptS11 string

const (
	AIGatewayGetResponseGuardrailsPromptS11Flag  AIGatewayGetResponseGuardrailsPromptS11 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptS11Block AIGatewayGetResponseGuardrailsPromptS11 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptS11) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptS11Flag, AIGatewayGetResponseGuardrailsPromptS11Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsPromptS12 string

const (
	AIGatewayGetResponseGuardrailsPromptS12Flag  AIGatewayGetResponseGuardrailsPromptS12 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptS12Block AIGatewayGetResponseGuardrailsPromptS12 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptS12) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptS12Flag, AIGatewayGetResponseGuardrailsPromptS12Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsPromptS13 string

const (
	AIGatewayGetResponseGuardrailsPromptS13Flag  AIGatewayGetResponseGuardrailsPromptS13 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptS13Block AIGatewayGetResponseGuardrailsPromptS13 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptS13) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptS13Flag, AIGatewayGetResponseGuardrailsPromptS13Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsPromptS2 string

const (
	AIGatewayGetResponseGuardrailsPromptS2Flag  AIGatewayGetResponseGuardrailsPromptS2 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptS2Block AIGatewayGetResponseGuardrailsPromptS2 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptS2) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptS2Flag, AIGatewayGetResponseGuardrailsPromptS2Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsPromptS3 string

const (
	AIGatewayGetResponseGuardrailsPromptS3Flag  AIGatewayGetResponseGuardrailsPromptS3 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptS3Block AIGatewayGetResponseGuardrailsPromptS3 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptS3) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptS3Flag, AIGatewayGetResponseGuardrailsPromptS3Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsPromptS4 string

const (
	AIGatewayGetResponseGuardrailsPromptS4Flag  AIGatewayGetResponseGuardrailsPromptS4 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptS4Block AIGatewayGetResponseGuardrailsPromptS4 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptS4) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptS4Flag, AIGatewayGetResponseGuardrailsPromptS4Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsPromptS5 string

const (
	AIGatewayGetResponseGuardrailsPromptS5Flag  AIGatewayGetResponseGuardrailsPromptS5 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptS5Block AIGatewayGetResponseGuardrailsPromptS5 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptS5) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptS5Flag, AIGatewayGetResponseGuardrailsPromptS5Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsPromptS6 string

const (
	AIGatewayGetResponseGuardrailsPromptS6Flag  AIGatewayGetResponseGuardrailsPromptS6 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptS6Block AIGatewayGetResponseGuardrailsPromptS6 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptS6) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptS6Flag, AIGatewayGetResponseGuardrailsPromptS6Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsPromptS7 string

const (
	AIGatewayGetResponseGuardrailsPromptS7Flag  AIGatewayGetResponseGuardrailsPromptS7 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptS7Block AIGatewayGetResponseGuardrailsPromptS7 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptS7) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptS7Flag, AIGatewayGetResponseGuardrailsPromptS7Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsPromptS8 string

const (
	AIGatewayGetResponseGuardrailsPromptS8Flag  AIGatewayGetResponseGuardrailsPromptS8 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptS8Block AIGatewayGetResponseGuardrailsPromptS8 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptS8) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptS8Flag, AIGatewayGetResponseGuardrailsPromptS8Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsPromptS9 string

const (
	AIGatewayGetResponseGuardrailsPromptS9Flag  AIGatewayGetResponseGuardrailsPromptS9 = "FLAG"
	AIGatewayGetResponseGuardrailsPromptS9Block AIGatewayGetResponseGuardrailsPromptS9 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsPromptS9) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsPromptS9Flag, AIGatewayGetResponseGuardrailsPromptS9Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponse struct {
	P1   AIGatewayGetResponseGuardrailsResponseP1   `json:"P1"`
	S1   AIGatewayGetResponseGuardrailsResponseS1   `json:"S1"`
	S10  AIGatewayGetResponseGuardrailsResponseS10  `json:"S10"`
	S11  AIGatewayGetResponseGuardrailsResponseS11  `json:"S11"`
	S12  AIGatewayGetResponseGuardrailsResponseS12  `json:"S12"`
	S13  AIGatewayGetResponseGuardrailsResponseS13  `json:"S13"`
	S2   AIGatewayGetResponseGuardrailsResponseS2   `json:"S2"`
	S3   AIGatewayGetResponseGuardrailsResponseS3   `json:"S3"`
	S4   AIGatewayGetResponseGuardrailsResponseS4   `json:"S4"`
	S5   AIGatewayGetResponseGuardrailsResponseS5   `json:"S5"`
	S6   AIGatewayGetResponseGuardrailsResponseS6   `json:"S6"`
	S7   AIGatewayGetResponseGuardrailsResponseS7   `json:"S7"`
	S8   AIGatewayGetResponseGuardrailsResponseS8   `json:"S8"`
	S9   AIGatewayGetResponseGuardrailsResponseS9   `json:"S9"`
	JSON aiGatewayGetResponseGuardrailsResponseJSON `json:"-"`
}

// aiGatewayGetResponseGuardrailsResponseJSON contains the JSON metadata for the
// struct [AIGatewayGetResponseGuardrailsResponse]
type aiGatewayGetResponseGuardrailsResponseJSON struct {
	P1          apijson.Field
	S1          apijson.Field
	S10         apijson.Field
	S11         apijson.Field
	S12         apijson.Field
	S13         apijson.Field
	S2          apijson.Field
	S3          apijson.Field
	S4          apijson.Field
	S5          apijson.Field
	S6          apijson.Field
	S7          apijson.Field
	S8          apijson.Field
	S9          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayGetResponseGuardrailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseGuardrailsResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayGetResponseGuardrailsResponseP1 string

const (
	AIGatewayGetResponseGuardrailsResponseP1Flag  AIGatewayGetResponseGuardrailsResponseP1 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseP1Block AIGatewayGetResponseGuardrailsResponseP1 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseP1) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseP1Flag, AIGatewayGetResponseGuardrailsResponseP1Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponseS1 string

const (
	AIGatewayGetResponseGuardrailsResponseS1Flag  AIGatewayGetResponseGuardrailsResponseS1 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseS1Block AIGatewayGetResponseGuardrailsResponseS1 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseS1) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseS1Flag, AIGatewayGetResponseGuardrailsResponseS1Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponseS10 string

const (
	AIGatewayGetResponseGuardrailsResponseS10Flag  AIGatewayGetResponseGuardrailsResponseS10 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseS10Block AIGatewayGetResponseGuardrailsResponseS10 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseS10) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseS10Flag, AIGatewayGetResponseGuardrailsResponseS10Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponseS11 string

const (
	AIGatewayGetResponseGuardrailsResponseS11Flag  AIGatewayGetResponseGuardrailsResponseS11 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseS11Block AIGatewayGetResponseGuardrailsResponseS11 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseS11) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseS11Flag, AIGatewayGetResponseGuardrailsResponseS11Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponseS12 string

const (
	AIGatewayGetResponseGuardrailsResponseS12Flag  AIGatewayGetResponseGuardrailsResponseS12 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseS12Block AIGatewayGetResponseGuardrailsResponseS12 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseS12) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseS12Flag, AIGatewayGetResponseGuardrailsResponseS12Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponseS13 string

const (
	AIGatewayGetResponseGuardrailsResponseS13Flag  AIGatewayGetResponseGuardrailsResponseS13 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseS13Block AIGatewayGetResponseGuardrailsResponseS13 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseS13) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseS13Flag, AIGatewayGetResponseGuardrailsResponseS13Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponseS2 string

const (
	AIGatewayGetResponseGuardrailsResponseS2Flag  AIGatewayGetResponseGuardrailsResponseS2 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseS2Block AIGatewayGetResponseGuardrailsResponseS2 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseS2) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseS2Flag, AIGatewayGetResponseGuardrailsResponseS2Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponseS3 string

const (
	AIGatewayGetResponseGuardrailsResponseS3Flag  AIGatewayGetResponseGuardrailsResponseS3 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseS3Block AIGatewayGetResponseGuardrailsResponseS3 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseS3) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseS3Flag, AIGatewayGetResponseGuardrailsResponseS3Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponseS4 string

const (
	AIGatewayGetResponseGuardrailsResponseS4Flag  AIGatewayGetResponseGuardrailsResponseS4 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseS4Block AIGatewayGetResponseGuardrailsResponseS4 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseS4) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseS4Flag, AIGatewayGetResponseGuardrailsResponseS4Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponseS5 string

const (
	AIGatewayGetResponseGuardrailsResponseS5Flag  AIGatewayGetResponseGuardrailsResponseS5 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseS5Block AIGatewayGetResponseGuardrailsResponseS5 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseS5) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseS5Flag, AIGatewayGetResponseGuardrailsResponseS5Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponseS6 string

const (
	AIGatewayGetResponseGuardrailsResponseS6Flag  AIGatewayGetResponseGuardrailsResponseS6 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseS6Block AIGatewayGetResponseGuardrailsResponseS6 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseS6) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseS6Flag, AIGatewayGetResponseGuardrailsResponseS6Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponseS7 string

const (
	AIGatewayGetResponseGuardrailsResponseS7Flag  AIGatewayGetResponseGuardrailsResponseS7 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseS7Block AIGatewayGetResponseGuardrailsResponseS7 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseS7) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseS7Flag, AIGatewayGetResponseGuardrailsResponseS7Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponseS8 string

const (
	AIGatewayGetResponseGuardrailsResponseS8Flag  AIGatewayGetResponseGuardrailsResponseS8 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseS8Block AIGatewayGetResponseGuardrailsResponseS8 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseS8) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseS8Flag, AIGatewayGetResponseGuardrailsResponseS8Block:
		return true
	}
	return false
}

type AIGatewayGetResponseGuardrailsResponseS9 string

const (
	AIGatewayGetResponseGuardrailsResponseS9Flag  AIGatewayGetResponseGuardrailsResponseS9 = "FLAG"
	AIGatewayGetResponseGuardrailsResponseS9Block AIGatewayGetResponseGuardrailsResponseS9 = "BLOCK"
)

func (r AIGatewayGetResponseGuardrailsResponseS9) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseGuardrailsResponseS9Flag, AIGatewayGetResponseGuardrailsResponseS9Block:
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
	Authorization string                              `json:"authorization" api:"required"`
	Headers       map[string]string                   `json:"headers" api:"required"`
	URL           string                              `json:"url" api:"required" format:"uri"`
	ContentType   AIGatewayGetResponseOtelContentType `json:"content_type"`
	JSON          aiGatewayGetResponseOtelJSON        `json:"-"`
}

// aiGatewayGetResponseOtelJSON contains the JSON metadata for the struct
// [AIGatewayGetResponseOtel]
type aiGatewayGetResponseOtelJSON struct {
	Authorization apijson.Field
	Headers       apijson.Field
	URL           apijson.Field
	ContentType   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIGatewayGetResponseOtel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseOtelJSON) RawJSON() string {
	return r.raw
}

type AIGatewayGetResponseOtelContentType string

const (
	AIGatewayGetResponseOtelContentTypeJson     AIGatewayGetResponseOtelContentType = "json"
	AIGatewayGetResponseOtelContentTypeProtobuf AIGatewayGetResponseOtelContentType = "protobuf"
)

func (r AIGatewayGetResponseOtelContentType) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseOtelContentTypeJson, AIGatewayGetResponseOtelContentTypeProtobuf:
		return true
	}
	return false
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

// Backoff strategy for retry delays
type AIGatewayGetResponseRetryBackoff string

const (
	AIGatewayGetResponseRetryBackoffConstant    AIGatewayGetResponseRetryBackoff = "constant"
	AIGatewayGetResponseRetryBackoffLinear      AIGatewayGetResponseRetryBackoff = "linear"
	AIGatewayGetResponseRetryBackoffExponential AIGatewayGetResponseRetryBackoff = "exponential"
)

func (r AIGatewayGetResponseRetryBackoff) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseRetryBackoffConstant, AIGatewayGetResponseRetryBackoffLinear, AIGatewayGetResponseRetryBackoffExponential:
		return true
	}
	return false
}

type AIGatewayGetResponseStripe struct {
	Authorization string                                 `json:"authorization" api:"required"`
	UsageEvents   []AIGatewayGetResponseStripeUsageEvent `json:"usage_events" api:"required"`
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
	Payload string                                   `json:"payload" api:"required"`
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

// Controls how Workers AI inference calls routed through this gateway are billed.
// Only 'postpaid' is currently supported.
type AIGatewayGetResponseWorkersAIBillingMode string

const (
	AIGatewayGetResponseWorkersAIBillingModePostpaid AIGatewayGetResponseWorkersAIBillingMode = "postpaid"
)

func (r AIGatewayGetResponseWorkersAIBillingMode) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseWorkersAIBillingModePostpaid:
		return true
	}
	return false
}

type AIGatewayNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// gateway id
	ID                      param.Field[string]                                  `json:"id" api:"required"`
	CacheInvalidateOnUpdate param.Field[bool]                                    `json:"cache_invalidate_on_update" api:"required"`
	CacheTTL                param.Field[int64]                                   `json:"cache_ttl" api:"required"`
	CollectLogs             param.Field[bool]                                    `json:"collect_logs" api:"required"`
	RateLimitingInterval    param.Field[int64]                                   `json:"rate_limiting_interval" api:"required"`
	RateLimitingLimit       param.Field[int64]                                   `json:"rate_limiting_limit" api:"required"`
	Authentication          param.Field[bool]                                    `json:"authentication"`
	LogManagement           param.Field[int64]                                   `json:"log_management"`
	LogManagementStrategy   param.Field[AIGatewayNewParamsLogManagementStrategy] `json:"log_management_strategy"`
	Logpush                 param.Field[bool]                                    `json:"logpush"`
	LogpushPublicKey        param.Field[string]                                  `json:"logpush_public_key"`
	RateLimitingTechnique   param.Field[AIGatewayNewParamsRateLimitingTechnique] `json:"rate_limiting_technique"`
	// Backoff strategy for retry delays
	RetryBackoff param.Field[AIGatewayNewParamsRetryBackoff] `json:"retry_backoff"`
	// Delay between retry attempts in milliseconds (0-5000)
	RetryDelay param.Field[int64] `json:"retry_delay"`
	// Maximum number of retry attempts for failed requests (1-5)
	RetryMaxAttempts param.Field[int64] `json:"retry_max_attempts"`
	// Controls how Workers AI inference calls routed through this gateway are billed.
	// Only 'postpaid' is currently supported.
	WorkersAIBillingMode param.Field[AIGatewayNewParamsWorkersAIBillingMode] `json:"workers_ai_billing_mode"`
	Zdr                  param.Field[bool]                                   `json:"zdr"`
}

func (r AIGatewayNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// Backoff strategy for retry delays
type AIGatewayNewParamsRetryBackoff string

const (
	AIGatewayNewParamsRetryBackoffConstant    AIGatewayNewParamsRetryBackoff = "constant"
	AIGatewayNewParamsRetryBackoffLinear      AIGatewayNewParamsRetryBackoff = "linear"
	AIGatewayNewParamsRetryBackoffExponential AIGatewayNewParamsRetryBackoff = "exponential"
)

func (r AIGatewayNewParamsRetryBackoff) IsKnown() bool {
	switch r {
	case AIGatewayNewParamsRetryBackoffConstant, AIGatewayNewParamsRetryBackoffLinear, AIGatewayNewParamsRetryBackoffExponential:
		return true
	}
	return false
}

// Controls how Workers AI inference calls routed through this gateway are billed.
// Only 'postpaid' is currently supported.
type AIGatewayNewParamsWorkersAIBillingMode string

const (
	AIGatewayNewParamsWorkersAIBillingModePostpaid AIGatewayNewParamsWorkersAIBillingMode = "postpaid"
)

func (r AIGatewayNewParamsWorkersAIBillingMode) IsKnown() bool {
	switch r {
	case AIGatewayNewParamsWorkersAIBillingModePostpaid:
		return true
	}
	return false
}

type AIGatewayNewResponseEnvelope struct {
	Result  AIGatewayNewResponse             `json:"result" api:"required"`
	Success bool                             `json:"success" api:"required"`
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
	AccountID               param.Field[string]                                     `path:"account_id" api:"required"`
	CacheInvalidateOnUpdate param.Field[bool]                                       `json:"cache_invalidate_on_update" api:"required"`
	CacheTTL                param.Field[int64]                                      `json:"cache_ttl" api:"required"`
	CollectLogs             param.Field[bool]                                       `json:"collect_logs" api:"required"`
	RateLimitingInterval    param.Field[int64]                                      `json:"rate_limiting_interval" api:"required"`
	RateLimitingLimit       param.Field[int64]                                      `json:"rate_limiting_limit" api:"required"`
	Authentication          param.Field[bool]                                       `json:"authentication"`
	DLP                     param.Field[AIGatewayUpdateParamsDLPUnion]              `json:"dlp"`
	Guardrails              param.Field[AIGatewayUpdateParamsGuardrails]            `json:"guardrails"`
	LogManagement           param.Field[int64]                                      `json:"log_management"`
	LogManagementStrategy   param.Field[AIGatewayUpdateParamsLogManagementStrategy] `json:"log_management_strategy"`
	Logpush                 param.Field[bool]                                       `json:"logpush"`
	LogpushPublicKey        param.Field[string]                                     `json:"logpush_public_key"`
	Otel                    param.Field[[]AIGatewayUpdateParamsOtel]                `json:"otel"`
	RateLimitingTechnique   param.Field[AIGatewayUpdateParamsRateLimitingTechnique] `json:"rate_limiting_technique"`
	// Backoff strategy for retry delays
	RetryBackoff param.Field[AIGatewayUpdateParamsRetryBackoff] `json:"retry_backoff"`
	// Delay between retry attempts in milliseconds (0-5000)
	RetryDelay param.Field[int64] `json:"retry_delay"`
	// Maximum number of retry attempts for failed requests (1-5)
	RetryMaxAttempts param.Field[int64]                       `json:"retry_max_attempts"`
	StoreID          param.Field[string]                      `json:"store_id"`
	Stripe           param.Field[AIGatewayUpdateParamsStripe] `json:"stripe"`
	// Controls how Workers AI inference calls routed through this gateway are billed.
	// Only 'postpaid' is currently supported.
	WorkersAIBillingMode param.Field[AIGatewayUpdateParamsWorkersAIBillingMode] `json:"workers_ai_billing_mode"`
	Zdr                  param.Field[bool]                                      `json:"zdr"`
}

func (r AIGatewayUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayUpdateParamsDLP struct {
	Enabled  param.Field[bool]                           `json:"enabled" api:"required"`
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
	Action   param.Field[AIGatewayUpdateParamsDLPObjectAction] `json:"action" api:"required"`
	Enabled  param.Field[bool]                                 `json:"enabled" api:"required"`
	Profiles param.Field[[]string]                             `json:"profiles" api:"required"`
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

type AIGatewayUpdateParamsGuardrails struct {
	Prompt   param.Field[AIGatewayUpdateParamsGuardrailsPrompt]   `json:"prompt" api:"required"`
	Response param.Field[AIGatewayUpdateParamsGuardrailsResponse] `json:"response" api:"required"`
}

func (r AIGatewayUpdateParamsGuardrails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayUpdateParamsGuardrailsPrompt struct {
	P1  param.Field[AIGatewayUpdateParamsGuardrailsPromptP1]  `json:"P1"`
	S1  param.Field[AIGatewayUpdateParamsGuardrailsPromptS1]  `json:"S1"`
	S10 param.Field[AIGatewayUpdateParamsGuardrailsPromptS10] `json:"S10"`
	S11 param.Field[AIGatewayUpdateParamsGuardrailsPromptS11] `json:"S11"`
	S12 param.Field[AIGatewayUpdateParamsGuardrailsPromptS12] `json:"S12"`
	S13 param.Field[AIGatewayUpdateParamsGuardrailsPromptS13] `json:"S13"`
	S2  param.Field[AIGatewayUpdateParamsGuardrailsPromptS2]  `json:"S2"`
	S3  param.Field[AIGatewayUpdateParamsGuardrailsPromptS3]  `json:"S3"`
	S4  param.Field[AIGatewayUpdateParamsGuardrailsPromptS4]  `json:"S4"`
	S5  param.Field[AIGatewayUpdateParamsGuardrailsPromptS5]  `json:"S5"`
	S6  param.Field[AIGatewayUpdateParamsGuardrailsPromptS6]  `json:"S6"`
	S7  param.Field[AIGatewayUpdateParamsGuardrailsPromptS7]  `json:"S7"`
	S8  param.Field[AIGatewayUpdateParamsGuardrailsPromptS8]  `json:"S8"`
	S9  param.Field[AIGatewayUpdateParamsGuardrailsPromptS9]  `json:"S9"`
}

func (r AIGatewayUpdateParamsGuardrailsPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayUpdateParamsGuardrailsPromptP1 string

const (
	AIGatewayUpdateParamsGuardrailsPromptP1Flag  AIGatewayUpdateParamsGuardrailsPromptP1 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptP1Block AIGatewayUpdateParamsGuardrailsPromptP1 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptP1) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptP1Flag, AIGatewayUpdateParamsGuardrailsPromptP1Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsPromptS1 string

const (
	AIGatewayUpdateParamsGuardrailsPromptS1Flag  AIGatewayUpdateParamsGuardrailsPromptS1 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptS1Block AIGatewayUpdateParamsGuardrailsPromptS1 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptS1) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptS1Flag, AIGatewayUpdateParamsGuardrailsPromptS1Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsPromptS10 string

const (
	AIGatewayUpdateParamsGuardrailsPromptS10Flag  AIGatewayUpdateParamsGuardrailsPromptS10 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptS10Block AIGatewayUpdateParamsGuardrailsPromptS10 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptS10) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptS10Flag, AIGatewayUpdateParamsGuardrailsPromptS10Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsPromptS11 string

const (
	AIGatewayUpdateParamsGuardrailsPromptS11Flag  AIGatewayUpdateParamsGuardrailsPromptS11 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptS11Block AIGatewayUpdateParamsGuardrailsPromptS11 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptS11) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptS11Flag, AIGatewayUpdateParamsGuardrailsPromptS11Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsPromptS12 string

const (
	AIGatewayUpdateParamsGuardrailsPromptS12Flag  AIGatewayUpdateParamsGuardrailsPromptS12 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptS12Block AIGatewayUpdateParamsGuardrailsPromptS12 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptS12) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptS12Flag, AIGatewayUpdateParamsGuardrailsPromptS12Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsPromptS13 string

const (
	AIGatewayUpdateParamsGuardrailsPromptS13Flag  AIGatewayUpdateParamsGuardrailsPromptS13 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptS13Block AIGatewayUpdateParamsGuardrailsPromptS13 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptS13) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptS13Flag, AIGatewayUpdateParamsGuardrailsPromptS13Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsPromptS2 string

const (
	AIGatewayUpdateParamsGuardrailsPromptS2Flag  AIGatewayUpdateParamsGuardrailsPromptS2 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptS2Block AIGatewayUpdateParamsGuardrailsPromptS2 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptS2) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptS2Flag, AIGatewayUpdateParamsGuardrailsPromptS2Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsPromptS3 string

const (
	AIGatewayUpdateParamsGuardrailsPromptS3Flag  AIGatewayUpdateParamsGuardrailsPromptS3 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptS3Block AIGatewayUpdateParamsGuardrailsPromptS3 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptS3) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptS3Flag, AIGatewayUpdateParamsGuardrailsPromptS3Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsPromptS4 string

const (
	AIGatewayUpdateParamsGuardrailsPromptS4Flag  AIGatewayUpdateParamsGuardrailsPromptS4 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptS4Block AIGatewayUpdateParamsGuardrailsPromptS4 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptS4) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptS4Flag, AIGatewayUpdateParamsGuardrailsPromptS4Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsPromptS5 string

const (
	AIGatewayUpdateParamsGuardrailsPromptS5Flag  AIGatewayUpdateParamsGuardrailsPromptS5 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptS5Block AIGatewayUpdateParamsGuardrailsPromptS5 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptS5) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptS5Flag, AIGatewayUpdateParamsGuardrailsPromptS5Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsPromptS6 string

const (
	AIGatewayUpdateParamsGuardrailsPromptS6Flag  AIGatewayUpdateParamsGuardrailsPromptS6 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptS6Block AIGatewayUpdateParamsGuardrailsPromptS6 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptS6) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptS6Flag, AIGatewayUpdateParamsGuardrailsPromptS6Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsPromptS7 string

const (
	AIGatewayUpdateParamsGuardrailsPromptS7Flag  AIGatewayUpdateParamsGuardrailsPromptS7 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptS7Block AIGatewayUpdateParamsGuardrailsPromptS7 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptS7) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptS7Flag, AIGatewayUpdateParamsGuardrailsPromptS7Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsPromptS8 string

const (
	AIGatewayUpdateParamsGuardrailsPromptS8Flag  AIGatewayUpdateParamsGuardrailsPromptS8 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptS8Block AIGatewayUpdateParamsGuardrailsPromptS8 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptS8) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptS8Flag, AIGatewayUpdateParamsGuardrailsPromptS8Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsPromptS9 string

const (
	AIGatewayUpdateParamsGuardrailsPromptS9Flag  AIGatewayUpdateParamsGuardrailsPromptS9 = "FLAG"
	AIGatewayUpdateParamsGuardrailsPromptS9Block AIGatewayUpdateParamsGuardrailsPromptS9 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsPromptS9) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsPromptS9Flag, AIGatewayUpdateParamsGuardrailsPromptS9Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponse struct {
	P1  param.Field[AIGatewayUpdateParamsGuardrailsResponseP1]  `json:"P1"`
	S1  param.Field[AIGatewayUpdateParamsGuardrailsResponseS1]  `json:"S1"`
	S10 param.Field[AIGatewayUpdateParamsGuardrailsResponseS10] `json:"S10"`
	S11 param.Field[AIGatewayUpdateParamsGuardrailsResponseS11] `json:"S11"`
	S12 param.Field[AIGatewayUpdateParamsGuardrailsResponseS12] `json:"S12"`
	S13 param.Field[AIGatewayUpdateParamsGuardrailsResponseS13] `json:"S13"`
	S2  param.Field[AIGatewayUpdateParamsGuardrailsResponseS2]  `json:"S2"`
	S3  param.Field[AIGatewayUpdateParamsGuardrailsResponseS3]  `json:"S3"`
	S4  param.Field[AIGatewayUpdateParamsGuardrailsResponseS4]  `json:"S4"`
	S5  param.Field[AIGatewayUpdateParamsGuardrailsResponseS5]  `json:"S5"`
	S6  param.Field[AIGatewayUpdateParamsGuardrailsResponseS6]  `json:"S6"`
	S7  param.Field[AIGatewayUpdateParamsGuardrailsResponseS7]  `json:"S7"`
	S8  param.Field[AIGatewayUpdateParamsGuardrailsResponseS8]  `json:"S8"`
	S9  param.Field[AIGatewayUpdateParamsGuardrailsResponseS9]  `json:"S9"`
}

func (r AIGatewayUpdateParamsGuardrailsResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayUpdateParamsGuardrailsResponseP1 string

const (
	AIGatewayUpdateParamsGuardrailsResponseP1Flag  AIGatewayUpdateParamsGuardrailsResponseP1 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseP1Block AIGatewayUpdateParamsGuardrailsResponseP1 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseP1) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseP1Flag, AIGatewayUpdateParamsGuardrailsResponseP1Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponseS1 string

const (
	AIGatewayUpdateParamsGuardrailsResponseS1Flag  AIGatewayUpdateParamsGuardrailsResponseS1 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseS1Block AIGatewayUpdateParamsGuardrailsResponseS1 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseS1) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseS1Flag, AIGatewayUpdateParamsGuardrailsResponseS1Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponseS10 string

const (
	AIGatewayUpdateParamsGuardrailsResponseS10Flag  AIGatewayUpdateParamsGuardrailsResponseS10 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseS10Block AIGatewayUpdateParamsGuardrailsResponseS10 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseS10) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseS10Flag, AIGatewayUpdateParamsGuardrailsResponseS10Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponseS11 string

const (
	AIGatewayUpdateParamsGuardrailsResponseS11Flag  AIGatewayUpdateParamsGuardrailsResponseS11 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseS11Block AIGatewayUpdateParamsGuardrailsResponseS11 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseS11) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseS11Flag, AIGatewayUpdateParamsGuardrailsResponseS11Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponseS12 string

const (
	AIGatewayUpdateParamsGuardrailsResponseS12Flag  AIGatewayUpdateParamsGuardrailsResponseS12 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseS12Block AIGatewayUpdateParamsGuardrailsResponseS12 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseS12) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseS12Flag, AIGatewayUpdateParamsGuardrailsResponseS12Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponseS13 string

const (
	AIGatewayUpdateParamsGuardrailsResponseS13Flag  AIGatewayUpdateParamsGuardrailsResponseS13 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseS13Block AIGatewayUpdateParamsGuardrailsResponseS13 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseS13) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseS13Flag, AIGatewayUpdateParamsGuardrailsResponseS13Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponseS2 string

const (
	AIGatewayUpdateParamsGuardrailsResponseS2Flag  AIGatewayUpdateParamsGuardrailsResponseS2 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseS2Block AIGatewayUpdateParamsGuardrailsResponseS2 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseS2) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseS2Flag, AIGatewayUpdateParamsGuardrailsResponseS2Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponseS3 string

const (
	AIGatewayUpdateParamsGuardrailsResponseS3Flag  AIGatewayUpdateParamsGuardrailsResponseS3 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseS3Block AIGatewayUpdateParamsGuardrailsResponseS3 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseS3) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseS3Flag, AIGatewayUpdateParamsGuardrailsResponseS3Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponseS4 string

const (
	AIGatewayUpdateParamsGuardrailsResponseS4Flag  AIGatewayUpdateParamsGuardrailsResponseS4 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseS4Block AIGatewayUpdateParamsGuardrailsResponseS4 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseS4) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseS4Flag, AIGatewayUpdateParamsGuardrailsResponseS4Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponseS5 string

const (
	AIGatewayUpdateParamsGuardrailsResponseS5Flag  AIGatewayUpdateParamsGuardrailsResponseS5 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseS5Block AIGatewayUpdateParamsGuardrailsResponseS5 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseS5) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseS5Flag, AIGatewayUpdateParamsGuardrailsResponseS5Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponseS6 string

const (
	AIGatewayUpdateParamsGuardrailsResponseS6Flag  AIGatewayUpdateParamsGuardrailsResponseS6 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseS6Block AIGatewayUpdateParamsGuardrailsResponseS6 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseS6) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseS6Flag, AIGatewayUpdateParamsGuardrailsResponseS6Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponseS7 string

const (
	AIGatewayUpdateParamsGuardrailsResponseS7Flag  AIGatewayUpdateParamsGuardrailsResponseS7 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseS7Block AIGatewayUpdateParamsGuardrailsResponseS7 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseS7) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseS7Flag, AIGatewayUpdateParamsGuardrailsResponseS7Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponseS8 string

const (
	AIGatewayUpdateParamsGuardrailsResponseS8Flag  AIGatewayUpdateParamsGuardrailsResponseS8 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseS8Block AIGatewayUpdateParamsGuardrailsResponseS8 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseS8) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseS8Flag, AIGatewayUpdateParamsGuardrailsResponseS8Block:
		return true
	}
	return false
}

type AIGatewayUpdateParamsGuardrailsResponseS9 string

const (
	AIGatewayUpdateParamsGuardrailsResponseS9Flag  AIGatewayUpdateParamsGuardrailsResponseS9 = "FLAG"
	AIGatewayUpdateParamsGuardrailsResponseS9Block AIGatewayUpdateParamsGuardrailsResponseS9 = "BLOCK"
)

func (r AIGatewayUpdateParamsGuardrailsResponseS9) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsGuardrailsResponseS9Flag, AIGatewayUpdateParamsGuardrailsResponseS9Block:
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
	Authorization param.Field[string]                               `json:"authorization" api:"required"`
	Headers       param.Field[map[string]string]                    `json:"headers" api:"required"`
	URL           param.Field[string]                               `json:"url" api:"required" format:"uri"`
	ContentType   param.Field[AIGatewayUpdateParamsOtelContentType] `json:"content_type"`
}

func (r AIGatewayUpdateParamsOtel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayUpdateParamsOtelContentType string

const (
	AIGatewayUpdateParamsOtelContentTypeJson     AIGatewayUpdateParamsOtelContentType = "json"
	AIGatewayUpdateParamsOtelContentTypeProtobuf AIGatewayUpdateParamsOtelContentType = "protobuf"
)

func (r AIGatewayUpdateParamsOtelContentType) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsOtelContentTypeJson, AIGatewayUpdateParamsOtelContentTypeProtobuf:
		return true
	}
	return false
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

// Backoff strategy for retry delays
type AIGatewayUpdateParamsRetryBackoff string

const (
	AIGatewayUpdateParamsRetryBackoffConstant    AIGatewayUpdateParamsRetryBackoff = "constant"
	AIGatewayUpdateParamsRetryBackoffLinear      AIGatewayUpdateParamsRetryBackoff = "linear"
	AIGatewayUpdateParamsRetryBackoffExponential AIGatewayUpdateParamsRetryBackoff = "exponential"
)

func (r AIGatewayUpdateParamsRetryBackoff) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsRetryBackoffConstant, AIGatewayUpdateParamsRetryBackoffLinear, AIGatewayUpdateParamsRetryBackoffExponential:
		return true
	}
	return false
}

type AIGatewayUpdateParamsStripe struct {
	Authorization param.Field[string]                                  `json:"authorization" api:"required"`
	UsageEvents   param.Field[[]AIGatewayUpdateParamsStripeUsageEvent] `json:"usage_events" api:"required"`
}

func (r AIGatewayUpdateParamsStripe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayUpdateParamsStripeUsageEvent struct {
	Payload param.Field[string] `json:"payload" api:"required"`
}

func (r AIGatewayUpdateParamsStripeUsageEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls how Workers AI inference calls routed through this gateway are billed.
// Only 'postpaid' is currently supported.
type AIGatewayUpdateParamsWorkersAIBillingMode string

const (
	AIGatewayUpdateParamsWorkersAIBillingModePostpaid AIGatewayUpdateParamsWorkersAIBillingMode = "postpaid"
)

func (r AIGatewayUpdateParamsWorkersAIBillingMode) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsWorkersAIBillingModePostpaid:
		return true
	}
	return false
}

type AIGatewayUpdateResponseEnvelope struct {
	Result  AIGatewayUpdateResponse             `json:"result" api:"required"`
	Success bool                                `json:"success" api:"required"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AIGatewayDeleteResponseEnvelope struct {
	Result  AIGatewayDeleteResponse             `json:"result" api:"required"`
	Success bool                                `json:"success" api:"required"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AIGatewayGetResponseEnvelope struct {
	Result  AIGatewayGetResponse             `json:"result" api:"required"`
	Success bool                             `json:"success" api:"required"`
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
