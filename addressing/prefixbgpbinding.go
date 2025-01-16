// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// PrefixBGPBindingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrefixBGPBindingService] method instead.
type PrefixBGPBindingService struct {
	Options []option.RequestOption
}

// NewPrefixBGPBindingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrefixBGPBindingService(opts ...option.RequestOption) (r *PrefixBGPBindingService) {
	r = &PrefixBGPBindingService{}
	r.Options = opts
	return
}

// Creates a new Service Binding, routing traffic to IPs within the given CIDR to a
// service running on Cloudflare's network. **Note:** This API may only be used on
// prefixes currently configured with a Magic Transit service binding, and only
// allows creating service bindings for the Cloudflare CDN or Cloudflare Spectrum.
func (r *PrefixBGPBindingService) New(ctx context.Context, prefixID string, params PrefixBGPBindingNewParams, opts ...option.RequestOption) (res *ServiceBinding, err error) {
	var env PrefixBGPBindingNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List the Cloudflare services this prefix is currently bound to. Traffic sent to
// an address within an IP prefix will be routed to the Cloudflare service of the
// most-specific Service Binding matching the address. **Example:** binding
// `192.0.2.0/24` to Cloudflare Magic Transit and `192.0.2.1/32` to the Cloudflare
// CDN would route traffic for `192.0.2.1` to the CDN, and traffic for all other
// IPs in the prefix to Cloudflare Magic Transit.
func (r *PrefixBGPBindingService) List(ctx context.Context, prefixID string, query PrefixBGPBindingListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ServiceBinding], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", query.AccountID, prefixID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// List the Cloudflare services this prefix is currently bound to. Traffic sent to
// an address within an IP prefix will be routed to the Cloudflare service of the
// most-specific Service Binding matching the address. **Example:** binding
// `192.0.2.0/24` to Cloudflare Magic Transit and `192.0.2.1/32` to the Cloudflare
// CDN would route traffic for `192.0.2.1` to the CDN, and traffic for all other
// IPs in the prefix to Cloudflare Magic Transit.
func (r *PrefixBGPBindingService) ListAutoPaging(ctx context.Context, prefixID string, query PrefixBGPBindingListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ServiceBinding] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, prefixID, query, opts...))
}

// Delete a Service Binding
func (r *PrefixBGPBindingService) Delete(ctx context.Context, prefixID string, bindingID string, body PrefixBGPBindingDeleteParams, opts ...option.RequestOption) (res *PrefixBGPBindingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	if bindingID == "" {
		err = errors.New("missing required binding_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", body.AccountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetch a single Service Binding
func (r *PrefixBGPBindingService) Get(ctx context.Context, prefixID string, bindingID string, query PrefixBGPBindingGetParams, opts ...option.RequestOption) (res *ServiceBinding, err error) {
	var env PrefixBGPBindingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	if bindingID == "" {
		err = errors.New("missing required binding_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", query.AccountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ServiceBinding struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning ServiceBindingProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string             `json:"service_name"`
	JSON        serviceBindingJSON `json:"-"`
}

// serviceBindingJSON contains the JSON metadata for the struct [ServiceBinding]
type serviceBindingJSON struct {
	ID           apijson.Field
	CIDR         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceBinding) implementsWorkersBinding() {}

// Status of a Service Binding's deployment to the Cloudflare network
type ServiceBindingProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State ServiceBindingProvisioningState `json:"state"`
	JSON  serviceBindingProvisioningJSON  `json:"-"`
}

// serviceBindingProvisioningJSON contains the JSON metadata for the struct
// [ServiceBindingProvisioning]
type serviceBindingProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceBindingProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceBindingProvisioningJSON) RawJSON() string {
	return r.raw
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type ServiceBindingProvisioningState string

const (
	ServiceBindingProvisioningStateProvisioning ServiceBindingProvisioningState = "provisioning"
	ServiceBindingProvisioningStateActive       ServiceBindingProvisioningState = "active"
)

func (r ServiceBindingProvisioningState) IsKnown() bool {
	switch r {
	case ServiceBindingProvisioningStateProvisioning, ServiceBindingProvisioningStateActive:
		return true
	}
	return false
}

type PrefixBGPBindingDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixBGPBindingDeleteResponseSuccess `json:"success,required"`
	JSON    prefixBGPBindingDeleteResponseJSON    `json:"-"`
}

// prefixBGPBindingDeleteResponseJSON contains the JSON metadata for the struct
// [PrefixBGPBindingDeleteResponse]
type prefixBGPBindingDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPBindingDeleteResponseSuccess bool

const (
	PrefixBGPBindingDeleteResponseSuccessTrue PrefixBGPBindingDeleteResponseSuccess = true
)

func (r PrefixBGPBindingDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case PrefixBGPBindingDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type PrefixBGPBindingNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR param.Field[string] `json:"cidr"`
	// Identifier
	ServiceID param.Field[string] `json:"service_id"`
}

func (r PrefixBGPBindingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixBGPBindingNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixBGPBindingNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ServiceBinding                             `json:"result"`
	JSON    prefixBGPBindingNewResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPBindingNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixBGPBindingNewResponseEnvelope]
type prefixBGPBindingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPBindingNewResponseEnvelopeSuccess bool

const (
	PrefixBGPBindingNewResponseEnvelopeSuccessTrue PrefixBGPBindingNewResponseEnvelopeSuccess = true
)

func (r PrefixBGPBindingNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixBGPBindingNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrefixBGPBindingListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixBGPBindingDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixBGPBindingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixBGPBindingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixBGPBindingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ServiceBinding                             `json:"result"`
	JSON    prefixBGPBindingGetResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPBindingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixBGPBindingGetResponseEnvelope]
type prefixBGPBindingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPBindingGetResponseEnvelopeSuccess bool

const (
	PrefixBGPBindingGetResponseEnvelopeSuccessTrue PrefixBGPBindingGetResponseEnvelopeSuccess = true
)

func (r PrefixBGPBindingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixBGPBindingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
