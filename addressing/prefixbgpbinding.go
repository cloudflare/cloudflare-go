// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// PrefixBGPBindingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPrefixBGPBindingService] method
// instead.
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
func (r *PrefixBGPBindingService) New(ctx context.Context, prefixID string, params PrefixBGPBindingNewParams, opts ...option.RequestOption) (res *PrefixBGPBindingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PrefixBGPBindingNewResponseEnvelope
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
func (r *PrefixBGPBindingService) List(ctx context.Context, prefixID string, query PrefixBGPBindingListParams, opts ...option.RequestOption) (res *pagination.SinglePage[PrefixBGPBindingListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", query.AccountID, prefixID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *PrefixBGPBindingService) ListAutoPaging(ctx context.Context, prefixID string, query PrefixBGPBindingListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PrefixBGPBindingListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, prefixID, query, opts...))
}

// Delete a Service Binding
func (r *PrefixBGPBindingService) Delete(ctx context.Context, prefixID string, bindingID string, body PrefixBGPBindingDeleteParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRefEc4d85c3d1bcc6b3b7e99c199ae99846Union, err error) {
	opts = append(r.Options[:], opts...)
	var env PrefixBGPBindingDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", body.AccountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single Service Binding
func (r *PrefixBGPBindingService) Get(ctx context.Context, prefixID string, bindingID string, query PrefixBGPBindingGetParams, opts ...option.RequestOption) (res *PrefixBGPBindingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PrefixBGPBindingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", query.AccountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PrefixBGPBindingNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning PrefixBGPBindingNewResponseProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                          `json:"service_name"`
	JSON        prefixBGPBindingNewResponseJSON `json:"-"`
}

// prefixBGPBindingNewResponseJSON contains the JSON metadata for the struct
// [PrefixBGPBindingNewResponse]
type prefixBGPBindingNewResponseJSON struct {
	ID           apijson.Field
	CIDR         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrefixBGPBindingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingNewResponseJSON) RawJSON() string {
	return r.raw
}

// Status of a Service Binding's deployment to the Cloudflare network
type PrefixBGPBindingNewResponseProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State PrefixBGPBindingNewResponseProvisioningState `json:"state"`
	JSON  prefixBGPBindingNewResponseProvisioningJSON  `json:"-"`
}

// prefixBGPBindingNewResponseProvisioningJSON contains the JSON metadata for the
// struct [PrefixBGPBindingNewResponseProvisioning]
type prefixBGPBindingNewResponseProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingNewResponseProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingNewResponseProvisioningJSON) RawJSON() string {
	return r.raw
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type PrefixBGPBindingNewResponseProvisioningState string

const (
	PrefixBGPBindingNewResponseProvisioningStateProvisioning PrefixBGPBindingNewResponseProvisioningState = "provisioning"
	PrefixBGPBindingNewResponseProvisioningStateActive       PrefixBGPBindingNewResponseProvisioningState = "active"
)

func (r PrefixBGPBindingNewResponseProvisioningState) IsKnown() bool {
	switch r {
	case PrefixBGPBindingNewResponseProvisioningStateProvisioning, PrefixBGPBindingNewResponseProvisioningStateActive:
		return true
	}
	return false
}

type PrefixBGPBindingListResponse struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning PrefixBGPBindingListResponseProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                           `json:"service_name"`
	JSON        prefixBGPBindingListResponseJSON `json:"-"`
}

// prefixBGPBindingListResponseJSON contains the JSON metadata for the struct
// [PrefixBGPBindingListResponse]
type prefixBGPBindingListResponseJSON struct {
	ID           apijson.Field
	CIDR         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrefixBGPBindingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingListResponseJSON) RawJSON() string {
	return r.raw
}

// Status of a Service Binding's deployment to the Cloudflare network
type PrefixBGPBindingListResponseProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State PrefixBGPBindingListResponseProvisioningState `json:"state"`
	JSON  prefixBGPBindingListResponseProvisioningJSON  `json:"-"`
}

// prefixBGPBindingListResponseProvisioningJSON contains the JSON metadata for the
// struct [PrefixBGPBindingListResponseProvisioning]
type prefixBGPBindingListResponseProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingListResponseProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingListResponseProvisioningJSON) RawJSON() string {
	return r.raw
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type PrefixBGPBindingListResponseProvisioningState string

const (
	PrefixBGPBindingListResponseProvisioningStateProvisioning PrefixBGPBindingListResponseProvisioningState = "provisioning"
	PrefixBGPBindingListResponseProvisioningStateActive       PrefixBGPBindingListResponseProvisioningState = "active"
)

func (r PrefixBGPBindingListResponseProvisioningState) IsKnown() bool {
	switch r {
	case PrefixBGPBindingListResponseProvisioningStateProvisioning, PrefixBGPBindingListResponseProvisioningStateActive:
		return true
	}
	return false
}

type PrefixBGPBindingGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning PrefixBGPBindingGetResponseProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                          `json:"service_name"`
	JSON        prefixBGPBindingGetResponseJSON `json:"-"`
}

// prefixBGPBindingGetResponseJSON contains the JSON metadata for the struct
// [PrefixBGPBindingGetResponse]
type prefixBGPBindingGetResponseJSON struct {
	ID           apijson.Field
	CIDR         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrefixBGPBindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingGetResponseJSON) RawJSON() string {
	return r.raw
}

// Status of a Service Binding's deployment to the Cloudflare network
type PrefixBGPBindingGetResponseProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State PrefixBGPBindingGetResponseProvisioningState `json:"state"`
	JSON  prefixBGPBindingGetResponseProvisioningJSON  `json:"-"`
}

// prefixBGPBindingGetResponseProvisioningJSON contains the JSON metadata for the
// struct [PrefixBGPBindingGetResponseProvisioning]
type prefixBGPBindingGetResponseProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingGetResponseProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingGetResponseProvisioningJSON) RawJSON() string {
	return r.raw
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type PrefixBGPBindingGetResponseProvisioningState string

const (
	PrefixBGPBindingGetResponseProvisioningStateProvisioning PrefixBGPBindingGetResponseProvisioningState = "provisioning"
	PrefixBGPBindingGetResponseProvisioningStateActive       PrefixBGPBindingGetResponseProvisioningState = "active"
)

func (r PrefixBGPBindingGetResponseProvisioningState) IsKnown() bool {
	switch r {
	case PrefixBGPBindingGetResponseProvisioningStateProvisioning, PrefixBGPBindingGetResponseProvisioningStateActive:
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   PrefixBGPBindingNewResponse                               `json:"result,required"`
	// Whether the API call was successful
	Success PrefixBGPBindingNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    prefixBGPBindingNewResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPBindingNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixBGPBindingNewResponseEnvelope]
type prefixBGPBindingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

type PrefixBGPBindingDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRefEc4d85c3d1bcc6b3b7e99c199ae99846Union `json:"result,required"`
	// Whether the API call was successful
	Success PrefixBGPBindingDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    prefixBGPBindingDeleteResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPBindingDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixBGPBindingDeleteResponseEnvelope]
type prefixBGPBindingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPBindingDeleteResponseEnvelopeSuccess bool

const (
	PrefixBGPBindingDeleteResponseEnvelopeSuccessTrue PrefixBGPBindingDeleteResponseEnvelopeSuccess = true
)

func (r PrefixBGPBindingDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixBGPBindingDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrefixBGPBindingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixBGPBindingGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   PrefixBGPBindingGetResponse                               `json:"result,required"`
	// Whether the API call was successful
	Success PrefixBGPBindingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    prefixBGPBindingGetResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPBindingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixBGPBindingGetResponseEnvelope]
type prefixBGPBindingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
