// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// PrefixServiceBindingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrefixServiceBindingService] method instead.
type PrefixServiceBindingService struct {
	Options []option.RequestOption
}

// NewPrefixServiceBindingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrefixServiceBindingService(opts ...option.RequestOption) (r *PrefixServiceBindingService) {
	r = &PrefixServiceBindingService{}
	r.Options = opts
	return
}

// Creates a new Service Binding, routing traffic to IPs within the given CIDR to a
// service running on Cloudflare's network. **Note:** This API may only be used on
// prefixes currently configured with a Magic Transit service binding, and only
// allows creating service bindings for the Cloudflare CDN or Cloudflare Spectrum.
func (r *PrefixServiceBindingService) New(ctx context.Context, prefixID string, params PrefixServiceBindingNewParams, opts ...option.RequestOption) (res *ServiceBinding, err error) {
	var env PrefixServiceBindingNewResponseEnvelope
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
func (r *PrefixServiceBindingService) List(ctx context.Context, prefixID string, query PrefixServiceBindingListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ServiceBinding], err error) {
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
func (r *PrefixServiceBindingService) ListAutoPaging(ctx context.Context, prefixID string, query PrefixServiceBindingListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ServiceBinding] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, prefixID, query, opts...))
}

// Delete a Service Binding
func (r *PrefixServiceBindingService) Delete(ctx context.Context, prefixID string, bindingID string, body PrefixServiceBindingDeleteParams, opts ...option.RequestOption) (res *PrefixServiceBindingDeleteResponse, err error) {
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
func (r *PrefixServiceBindingService) Get(ctx context.Context, prefixID string, bindingID string, query PrefixServiceBindingGetParams, opts ...option.RequestOption) (res *ServiceBinding, err error) {
	var env PrefixServiceBindingGetResponseEnvelope
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
	// Identifier of a Service Binding.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning ServiceBindingProvisioning `json:"provisioning"`
	// Identifier of a Service on the Cloudflare network. Available services and their
	// IDs may be found in the **List Services** endpoint.
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

type PrefixServiceBindingDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixServiceBindingDeleteResponseSuccess `json:"success,required"`
	JSON    prefixServiceBindingDeleteResponseJSON    `json:"-"`
}

// prefixServiceBindingDeleteResponseJSON contains the JSON metadata for the struct
// [PrefixServiceBindingDeleteResponse]
type prefixServiceBindingDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixServiceBindingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixServiceBindingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixServiceBindingDeleteResponseSuccess bool

const (
	PrefixServiceBindingDeleteResponseSuccessTrue PrefixServiceBindingDeleteResponseSuccess = true
)

func (r PrefixServiceBindingDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case PrefixServiceBindingDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type PrefixServiceBindingNewParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR param.Field[string] `json:"cidr"`
	// Identifier of a Service on the Cloudflare network. Available services and their
	// IDs may be found in the **List Services** endpoint.
	ServiceID param.Field[string] `json:"service_id"`
}

func (r PrefixServiceBindingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixServiceBindingNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixServiceBindingNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ServiceBinding                                 `json:"result"`
	JSON    prefixServiceBindingNewResponseEnvelopeJSON    `json:"-"`
}

// prefixServiceBindingNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixServiceBindingNewResponseEnvelope]
type prefixServiceBindingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixServiceBindingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixServiceBindingNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixServiceBindingNewResponseEnvelopeSuccess bool

const (
	PrefixServiceBindingNewResponseEnvelopeSuccessTrue PrefixServiceBindingNewResponseEnvelopeSuccess = true
)

func (r PrefixServiceBindingNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixServiceBindingNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrefixServiceBindingListParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixServiceBindingDeleteParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixServiceBindingGetParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixServiceBindingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixServiceBindingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ServiceBinding                                 `json:"result"`
	JSON    prefixServiceBindingGetResponseEnvelopeJSON    `json:"-"`
}

// prefixServiceBindingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixServiceBindingGetResponseEnvelope]
type prefixServiceBindingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixServiceBindingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixServiceBindingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixServiceBindingGetResponseEnvelopeSuccess bool

const (
	PrefixServiceBindingGetResponseEnvelopeSuccessTrue PrefixServiceBindingGetResponseEnvelopeSuccess = true
)

func (r PrefixServiceBindingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixServiceBindingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
