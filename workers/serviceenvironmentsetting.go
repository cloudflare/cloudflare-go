// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ServiceEnvironmentSettingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewServiceEnvironmentSettingService] method instead.
type ServiceEnvironmentSettingService struct {
	Options []option.RequestOption
}

// NewServiceEnvironmentSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewServiceEnvironmentSettingService(opts ...option.RequestOption) (r *ServiceEnvironmentSettingService) {
	r = &ServiceEnvironmentSettingService{}
	r.Options = opts
	return
}

// Patch script metadata, such as bindings
func (r *ServiceEnvironmentSettingService) Edit(ctx context.Context, serviceName string, environmentName string, params ServiceEnvironmentSettingEditParams, opts ...option.RequestOption) (res *ServiceEnvironmentSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ServiceEnvironmentSettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", params.AccountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script settings from a worker with an environment
func (r *ServiceEnvironmentSettingService) Get(ctx context.Context, serviceName string, environmentName string, query ServiceEnvironmentSettingGetParams, opts ...option.RequestOption) (res *ServiceEnvironmentSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ServiceEnvironmentSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", query.AccountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ServiceEnvironmentSettingEditResponse struct {
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ServiceEnvironmentSettingEditResponseTailConsumer `json:"tail_consumers"`
	JSON          serviceEnvironmentSettingEditResponseJSON           `json:"-"`
}

// serviceEnvironmentSettingEditResponseJSON contains the JSON metadata for the
// struct [ServiceEnvironmentSettingEditResponse]
type serviceEnvironmentSettingEditResponseJSON struct {
	Logpush       apijson.Field
	TailConsumers apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a script that will consume logs from the attached Worker.
type ServiceEnvironmentSettingEditResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                `json:"namespace"`
	JSON      serviceEnvironmentSettingEditResponseTailConsumerJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseTailConsumerJSON contains the JSON metadata
// for the struct [ServiceEnvironmentSettingEditResponseTailConsumer]
type serviceEnvironmentSettingEditResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingGetResponse struct {
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ServiceEnvironmentSettingGetResponseTailConsumer `json:"tail_consumers"`
	JSON          serviceEnvironmentSettingGetResponseJSON           `json:"-"`
}

// serviceEnvironmentSettingGetResponseJSON contains the JSON metadata for the
// struct [ServiceEnvironmentSettingGetResponse]
type serviceEnvironmentSettingGetResponseJSON struct {
	Logpush       apijson.Field
	TailConsumers apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a script that will consume logs from the attached Worker.
type ServiceEnvironmentSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                               `json:"namespace"`
	JSON      serviceEnvironmentSettingGetResponseTailConsumerJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseTailConsumerJSON contains the JSON metadata
// for the struct [ServiceEnvironmentSettingGetResponseTailConsumer]
type serviceEnvironmentSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]                                    `path:"account_id,required"`
	Errors    param.Field[[]shared.UnnamedSchemaRef172Param]         `json:"errors,required"`
	Messages  param.Field[[]shared.UnnamedSchemaRef172Param]         `json:"messages,required"`
	Result    param.Field[ServiceEnvironmentSettingEditParamsResult] `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[ServiceEnvironmentSettingEditParamsSuccess] `json:"success,required"`
}

func (r ServiceEnvironmentSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceEnvironmentSettingEditParamsResult struct {
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ServiceEnvironmentSettingEditParamsResultTailConsumer] `json:"tail_consumers"`
}

func (r ServiceEnvironmentSettingEditParamsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A reference to a script that will consume logs from the attached Worker.
type ServiceEnvironmentSettingEditParamsResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r ServiceEnvironmentSettingEditParamsResultTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type ServiceEnvironmentSettingEditParamsSuccess bool

const (
	ServiceEnvironmentSettingEditParamsSuccessTrue ServiceEnvironmentSettingEditParamsSuccess = true
)

func (r ServiceEnvironmentSettingEditParamsSuccess) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditParamsSuccessTrue:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172          `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172          `json:"messages,required"`
	Result   ServiceEnvironmentSettingEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success ServiceEnvironmentSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    serviceEnvironmentSettingEditResponseEnvelopeJSON    `json:"-"`
}

// serviceEnvironmentSettingEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ServiceEnvironmentSettingEditResponseEnvelope]
type serviceEnvironmentSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ServiceEnvironmentSettingEditResponseEnvelopeSuccess bool

const (
	ServiceEnvironmentSettingEditResponseEnvelopeSuccessTrue ServiceEnvironmentSettingEditResponseEnvelopeSuccess = true
)

func (r ServiceEnvironmentSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ServiceEnvironmentSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ServiceEnvironmentSettingGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172         `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172         `json:"messages,required"`
	Result   ServiceEnvironmentSettingGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success ServiceEnvironmentSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    serviceEnvironmentSettingGetResponseEnvelopeJSON    `json:"-"`
}

// serviceEnvironmentSettingGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ServiceEnvironmentSettingGetResponseEnvelope]
type serviceEnvironmentSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ServiceEnvironmentSettingGetResponseEnvelopeSuccess bool

const (
	ServiceEnvironmentSettingGetResponseEnvelopeSuccessTrue ServiceEnvironmentSettingGetResponseEnvelopeSuccess = true
)

func (r ServiceEnvironmentSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
