// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// CloudforceOneRequestPriorityService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewCloudforceOneRequestPriorityService] method instead.
type CloudforceOneRequestPriorityService struct {
	Options []option.RequestOption
}

// NewCloudforceOneRequestPriorityService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudforceOneRequestPriorityService(opts ...option.RequestOption) (r *CloudforceOneRequestPriorityService) {
	r = &CloudforceOneRequestPriorityService{}
	r.Options = opts
	return
}

// Create a New Priority Requirement
func (r *CloudforceOneRequestPriorityService) New(ctx context.Context, accountIdentifier string, body CloudforceOneRequestPriorityNewParams, opts ...option.RequestOption) (res *CloudforceOnePriorityItem, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestPriorityNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/new", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Priority Intelligence Requirement
func (r *CloudforceOneRequestPriorityService) Update(ctx context.Context, accountIdentifier string, priorityIdentifer string, body CloudforceOneRequestPriorityUpdateParams, opts ...option.RequestOption) (res *CloudforceOneRequestItem, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestPriorityUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Priority Intelligence Report
func (r *CloudforceOneRequestPriorityService) Delete(ctx context.Context, accountIdentifier string, priorityIdentifer string, opts ...option.RequestOption) (res *CloudforceOneRequestPriorityDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestPriorityDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a Priority Intelligence Requirement
func (r *CloudforceOneRequestPriorityService) Get(ctx context.Context, accountIdentifier string, priorityIdentifer string, opts ...option.RequestOption) (res *CloudforceOneRequestItem, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestPriorityGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Priority Intelligence Requirement Quota
func (r *CloudforceOneRequestPriorityService) Quota(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *CloudforceOneQuota, err error) {
	opts = append(r.Options[:], opts...)
	var env CloudforceOneRequestPriorityQuotaResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/quota", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CloudforceOnePriorityItem struct {
	// UUID
	ID string `json:"id,required"`
	// Priority creation time
	Created time.Time `json:"created,required" format:"date-time"`
	// List of labels
	Labels []string `json:"labels,required"`
	// Priority
	Priority int64 `json:"priority,required"`
	// Requirement
	Requirement string `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp CloudforceOnePriorityItemTlp `json:"tlp,required"`
	// Priority last updated time
	Updated time.Time                     `json:"updated,required" format:"date-time"`
	JSON    cloudforceOnePriorityItemJSON `json:"-"`
}

// cloudforceOnePriorityItemJSON contains the JSON metadata for the struct
// [CloudforceOnePriorityItem]
type cloudforceOnePriorityItemJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Labels      apijson.Field
	Priority    apijson.Field
	Requirement apijson.Field
	Tlp         apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOnePriorityItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOnePriorityItemTlp string

const (
	CloudforceOnePriorityItemTlpClear       CloudforceOnePriorityItemTlp = "clear"
	CloudforceOnePriorityItemTlpAmber       CloudforceOnePriorityItemTlp = "amber"
	CloudforceOnePriorityItemTlpAmberStrict CloudforceOnePriorityItemTlp = "amber-strict"
	CloudforceOnePriorityItemTlpGreen       CloudforceOnePriorityItemTlp = "green"
	CloudforceOnePriorityItemTlpRed         CloudforceOnePriorityItemTlp = "red"
)

// Union satisfied by [CloudforceOneRequestPriorityDeleteResponseUnknown],
// [CloudforceOneRequestPriorityDeleteResponseArray] or [shared.UnionString].
type CloudforceOneRequestPriorityDeleteResponse interface {
	ImplementsCloudforceOneRequestPriorityDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudforceOneRequestPriorityDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudforceOneRequestPriorityDeleteResponseArray []interface{}

func (r CloudforceOneRequestPriorityDeleteResponseArray) ImplementsCloudforceOneRequestPriorityDeleteResponse() {
}

type CloudforceOneRequestPriorityNewParams struct {
	// List of labels
	Labels param.Field[[]string] `json:"labels,required"`
	// Priority
	Priority param.Field[int64] `json:"priority,required"`
	// Requirement
	Requirement param.Field[string] `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[CloudforceOneRequestPriorityNewParamsTlp] `json:"tlp,required"`
}

func (r CloudforceOneRequestPriorityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestPriorityNewParamsTlp string

const (
	CloudforceOneRequestPriorityNewParamsTlpClear       CloudforceOneRequestPriorityNewParamsTlp = "clear"
	CloudforceOneRequestPriorityNewParamsTlpAmber       CloudforceOneRequestPriorityNewParamsTlp = "amber"
	CloudforceOneRequestPriorityNewParamsTlpAmberStrict CloudforceOneRequestPriorityNewParamsTlp = "amber-strict"
	CloudforceOneRequestPriorityNewParamsTlpGreen       CloudforceOneRequestPriorityNewParamsTlp = "green"
	CloudforceOneRequestPriorityNewParamsTlpRed         CloudforceOneRequestPriorityNewParamsTlp = "red"
)

type CloudforceOneRequestPriorityNewResponseEnvelope struct {
	Errors   []CloudforceOneRequestPriorityNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestPriorityNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOnePriorityItem                                 `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestPriorityNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestPriorityNewResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestPriorityNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [CloudforceOneRequestPriorityNewResponseEnvelope]
type cloudforceOneRequestPriorityNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityNewResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    cloudforceOneRequestPriorityNewResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestPriorityNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CloudforceOneRequestPriorityNewResponseEnvelopeErrors]
type cloudforceOneRequestPriorityNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityNewResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    cloudforceOneRequestPriorityNewResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestPriorityNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityNewResponseEnvelopeMessages]
type cloudforceOneRequestPriorityNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestPriorityNewResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestPriorityNewResponseEnvelopeSuccessTrue CloudforceOneRequestPriorityNewResponseEnvelopeSuccess = true
)

type CloudforceOneRequestPriorityUpdateParams struct {
	// List of labels
	Labels param.Field[[]string] `json:"labels,required"`
	// Priority
	Priority param.Field[int64] `json:"priority,required"`
	// Requirement
	Requirement param.Field[string] `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[CloudforceOneRequestPriorityUpdateParamsTlp] `json:"tlp,required"`
}

func (r CloudforceOneRequestPriorityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type CloudforceOneRequestPriorityUpdateParamsTlp string

const (
	CloudforceOneRequestPriorityUpdateParamsTlpClear       CloudforceOneRequestPriorityUpdateParamsTlp = "clear"
	CloudforceOneRequestPriorityUpdateParamsTlpAmber       CloudforceOneRequestPriorityUpdateParamsTlp = "amber"
	CloudforceOneRequestPriorityUpdateParamsTlpAmberStrict CloudforceOneRequestPriorityUpdateParamsTlp = "amber-strict"
	CloudforceOneRequestPriorityUpdateParamsTlpGreen       CloudforceOneRequestPriorityUpdateParamsTlp = "green"
	CloudforceOneRequestPriorityUpdateParamsTlpRed         CloudforceOneRequestPriorityUpdateParamsTlp = "red"
)

type CloudforceOneRequestPriorityUpdateResponseEnvelope struct {
	Errors   []CloudforceOneRequestPriorityUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestPriorityUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestItem                                     `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestPriorityUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestPriorityUpdateResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestPriorityUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [CloudforceOneRequestPriorityUpdateResponseEnvelope]
type cloudforceOneRequestPriorityUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityUpdateResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    cloudforceOneRequestPriorityUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestPriorityUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityUpdateResponseEnvelopeErrors]
type cloudforceOneRequestPriorityUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityUpdateResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    cloudforceOneRequestPriorityUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestPriorityUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityUpdateResponseEnvelopeMessages]
type cloudforceOneRequestPriorityUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestPriorityUpdateResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestPriorityUpdateResponseEnvelopeSuccessTrue CloudforceOneRequestPriorityUpdateResponseEnvelopeSuccess = true
)

type CloudforceOneRequestPriorityDeleteResponseEnvelope struct {
	Errors   []CloudforceOneRequestPriorityDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestPriorityDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestPriorityDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestPriorityDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestPriorityDeleteResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestPriorityDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [CloudforceOneRequestPriorityDeleteResponseEnvelope]
type cloudforceOneRequestPriorityDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityDeleteResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    cloudforceOneRequestPriorityDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestPriorityDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityDeleteResponseEnvelopeErrors]
type cloudforceOneRequestPriorityDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityDeleteResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    cloudforceOneRequestPriorityDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestPriorityDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityDeleteResponseEnvelopeMessages]
type cloudforceOneRequestPriorityDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestPriorityDeleteResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestPriorityDeleteResponseEnvelopeSuccessTrue CloudforceOneRequestPriorityDeleteResponseEnvelopeSuccess = true
)

type CloudforceOneRequestPriorityGetResponseEnvelope struct {
	Errors   []CloudforceOneRequestPriorityGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestPriorityGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneRequestItem                                  `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestPriorityGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestPriorityGetResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestPriorityGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [CloudforceOneRequestPriorityGetResponseEnvelope]
type cloudforceOneRequestPriorityGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    cloudforceOneRequestPriorityGetResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestPriorityGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CloudforceOneRequestPriorityGetResponseEnvelopeErrors]
type cloudforceOneRequestPriorityGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    cloudforceOneRequestPriorityGetResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestPriorityGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityGetResponseEnvelopeMessages]
type cloudforceOneRequestPriorityGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestPriorityGetResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestPriorityGetResponseEnvelopeSuccessTrue CloudforceOneRequestPriorityGetResponseEnvelopeSuccess = true
)

type CloudforceOneRequestPriorityQuotaResponseEnvelope struct {
	Errors   []CloudforceOneRequestPriorityQuotaResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CloudforceOneRequestPriorityQuotaResponseEnvelopeMessages `json:"messages,required"`
	Result   CloudforceOneQuota                                          `json:"result,required"`
	// Whether the API call was successful
	Success CloudforceOneRequestPriorityQuotaResponseEnvelopeSuccess `json:"success,required"`
	JSON    cloudforceOneRequestPriorityQuotaResponseEnvelopeJSON    `json:"-"`
}

// cloudforceOneRequestPriorityQuotaResponseEnvelopeJSON contains the JSON metadata
// for the struct [CloudforceOneRequestPriorityQuotaResponseEnvelope]
type cloudforceOneRequestPriorityQuotaResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityQuotaResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityQuotaResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    cloudforceOneRequestPriorityQuotaResponseEnvelopeErrorsJSON `json:"-"`
}

// cloudforceOneRequestPriorityQuotaResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityQuotaResponseEnvelopeErrors]
type cloudforceOneRequestPriorityQuotaResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityQuotaResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CloudforceOneRequestPriorityQuotaResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    cloudforceOneRequestPriorityQuotaResponseEnvelopeMessagesJSON `json:"-"`
}

// cloudforceOneRequestPriorityQuotaResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CloudforceOneRequestPriorityQuotaResponseEnvelopeMessages]
type cloudforceOneRequestPriorityQuotaResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudforceOneRequestPriorityQuotaResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CloudforceOneRequestPriorityQuotaResponseEnvelopeSuccess bool

const (
	CloudforceOneRequestPriorityQuotaResponseEnvelopeSuccessTrue CloudforceOneRequestPriorityQuotaResponseEnvelopeSuccess = true
)
