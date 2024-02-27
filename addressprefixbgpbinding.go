// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AddressPrefixBGPBindingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAddressPrefixBGPBindingService] method instead.
type AddressPrefixBGPBindingService struct {
	Options []option.RequestOption
}

// NewAddressPrefixBGPBindingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressPrefixBGPBindingService(opts ...option.RequestOption) (r *AddressPrefixBGPBindingService) {
	r = &AddressPrefixBGPBindingService{}
	r.Options = opts
	return
}

// Creates a new Service Binding, routing traffic to IPs within the given CIDR to a
// service running on Cloudflare's network. **Note:** This API may only be used on
// prefixes currently configured with a Magic Transit service binding, and only
// allows creating service bindings for the Cloudflare CDN or Cloudflare Spectrum.
func (r *AddressPrefixBGPBindingService) New(ctx context.Context, prefixID string, params AddressPrefixBGPBindingNewParams, opts ...option.RequestOption) (res *AddressPrefixBGPBindingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixBGPBindingNewResponseEnvelope
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
func (r *AddressPrefixBGPBindingService) List(ctx context.Context, prefixID string, query AddressPrefixBGPBindingListParams, opts ...option.RequestOption) (res *[]AddressPrefixBGPBindingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixBGPBindingListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", query.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Service Binding
func (r *AddressPrefixBGPBindingService) Delete(ctx context.Context, prefixID string, bindingID string, body AddressPrefixBGPBindingDeleteParams, opts ...option.RequestOption) (res *AddressPrefixBGPBindingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixBGPBindingDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", body.AccountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single Service Binding
func (r *AddressPrefixBGPBindingService) Get(ctx context.Context, prefixID string, bindingID string, query AddressPrefixBGPBindingGetParams, opts ...option.RequestOption) (res *AddressPrefixBGPBindingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixBGPBindingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", query.AccountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressPrefixBGPBindingNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AddressPrefixBGPBindingNewResponseProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                 `json:"service_name"`
	JSON        addressPrefixBGPBindingNewResponseJSON `json:"-"`
}

// addressPrefixBGPBindingNewResponseJSON contains the JSON metadata for the struct
// [AddressPrefixBGPBindingNewResponse]
type addressPrefixBGPBindingNewResponseJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a Service Binding's deployment to the Cloudflare network
type AddressPrefixBGPBindingNewResponseProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AddressPrefixBGPBindingNewResponseProvisioningState `json:"state"`
	JSON  addressPrefixBGPBindingNewResponseProvisioningJSON  `json:"-"`
}

// addressPrefixBGPBindingNewResponseProvisioningJSON contains the JSON metadata
// for the struct [AddressPrefixBGPBindingNewResponseProvisioning]
type addressPrefixBGPBindingNewResponseProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingNewResponseProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AddressPrefixBGPBindingNewResponseProvisioningState string

const (
	AddressPrefixBGPBindingNewResponseProvisioningStateProvisioning AddressPrefixBGPBindingNewResponseProvisioningState = "provisioning"
	AddressPrefixBGPBindingNewResponseProvisioningStateActive       AddressPrefixBGPBindingNewResponseProvisioningState = "active"
)

type AddressPrefixBGPBindingListResponse struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AddressPrefixBGPBindingListResponseProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                  `json:"service_name"`
	JSON        addressPrefixBGPBindingListResponseJSON `json:"-"`
}

// addressPrefixBGPBindingListResponseJSON contains the JSON metadata for the
// struct [AddressPrefixBGPBindingListResponse]
type addressPrefixBGPBindingListResponseJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a Service Binding's deployment to the Cloudflare network
type AddressPrefixBGPBindingListResponseProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AddressPrefixBGPBindingListResponseProvisioningState `json:"state"`
	JSON  addressPrefixBGPBindingListResponseProvisioningJSON  `json:"-"`
}

// addressPrefixBGPBindingListResponseProvisioningJSON contains the JSON metadata
// for the struct [AddressPrefixBGPBindingListResponseProvisioning]
type addressPrefixBGPBindingListResponseProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingListResponseProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AddressPrefixBGPBindingListResponseProvisioningState string

const (
	AddressPrefixBGPBindingListResponseProvisioningStateProvisioning AddressPrefixBGPBindingListResponseProvisioningState = "provisioning"
	AddressPrefixBGPBindingListResponseProvisioningStateActive       AddressPrefixBGPBindingListResponseProvisioningState = "active"
)

// Union satisfied by [AddressPrefixBGPBindingDeleteResponseUnknown],
// [AddressPrefixBGPBindingDeleteResponseArray] or [shared.UnionString].
type AddressPrefixBGPBindingDeleteResponse interface {
	ImplementsAddressPrefixBGPBindingDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressPrefixBGPBindingDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressPrefixBGPBindingDeleteResponseArray []interface{}

func (r AddressPrefixBGPBindingDeleteResponseArray) ImplementsAddressPrefixBGPBindingDeleteResponse() {
}

type AddressPrefixBGPBindingGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AddressPrefixBGPBindingGetResponseProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                 `json:"service_name"`
	JSON        addressPrefixBGPBindingGetResponseJSON `json:"-"`
}

// addressPrefixBGPBindingGetResponseJSON contains the JSON metadata for the struct
// [AddressPrefixBGPBindingGetResponse]
type addressPrefixBGPBindingGetResponseJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a Service Binding's deployment to the Cloudflare network
type AddressPrefixBGPBindingGetResponseProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AddressPrefixBGPBindingGetResponseProvisioningState `json:"state"`
	JSON  addressPrefixBGPBindingGetResponseProvisioningJSON  `json:"-"`
}

// addressPrefixBGPBindingGetResponseProvisioningJSON contains the JSON metadata
// for the struct [AddressPrefixBGPBindingGetResponseProvisioning]
type addressPrefixBGPBindingGetResponseProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingGetResponseProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AddressPrefixBGPBindingGetResponseProvisioningState string

const (
	AddressPrefixBGPBindingGetResponseProvisioningStateProvisioning AddressPrefixBGPBindingGetResponseProvisioningState = "provisioning"
	AddressPrefixBGPBindingGetResponseProvisioningStateActive       AddressPrefixBGPBindingGetResponseProvisioningState = "active"
)

type AddressPrefixBGPBindingNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr"`
	// Identifier
	ServiceID param.Field[string] `json:"service_id"`
}

func (r AddressPrefixBGPBindingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressPrefixBGPBindingNewResponseEnvelope struct {
	Errors   []AddressPrefixBGPBindingNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixBGPBindingNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixBGPBindingNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixBGPBindingNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixBGPBindingNewResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixBGPBindingNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressPrefixBGPBindingNewResponseEnvelope]
type addressPrefixBGPBindingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPBindingNewResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    addressPrefixBGPBindingNewResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixBGPBindingNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressPrefixBGPBindingNewResponseEnvelopeErrors]
type addressPrefixBGPBindingNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPBindingNewResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressPrefixBGPBindingNewResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixBGPBindingNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressPrefixBGPBindingNewResponseEnvelopeMessages]
type addressPrefixBGPBindingNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixBGPBindingNewResponseEnvelopeSuccess bool

const (
	AddressPrefixBGPBindingNewResponseEnvelopeSuccessTrue AddressPrefixBGPBindingNewResponseEnvelopeSuccess = true
)

type AddressPrefixBGPBindingListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressPrefixBGPBindingListResponseEnvelope struct {
	Errors   []AddressPrefixBGPBindingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixBGPBindingListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressPrefixBGPBindingListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixBGPBindingListResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixBGPBindingListResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixBGPBindingListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressPrefixBGPBindingListResponseEnvelope]
type addressPrefixBGPBindingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPBindingListResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    addressPrefixBGPBindingListResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixBGPBindingListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressPrefixBGPBindingListResponseEnvelopeErrors]
type addressPrefixBGPBindingListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPBindingListResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressPrefixBGPBindingListResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixBGPBindingListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressPrefixBGPBindingListResponseEnvelopeMessages]
type addressPrefixBGPBindingListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixBGPBindingListResponseEnvelopeSuccess bool

const (
	AddressPrefixBGPBindingListResponseEnvelopeSuccessTrue AddressPrefixBGPBindingListResponseEnvelopeSuccess = true
)

type AddressPrefixBGPBindingDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressPrefixBGPBindingDeleteResponseEnvelope struct {
	Errors   []AddressPrefixBGPBindingDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixBGPBindingDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixBGPBindingDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixBGPBindingDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixBGPBindingDeleteResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixBGPBindingDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressPrefixBGPBindingDeleteResponseEnvelope]
type addressPrefixBGPBindingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPBindingDeleteResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressPrefixBGPBindingDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixBGPBindingDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressPrefixBGPBindingDeleteResponseEnvelopeErrors]
type addressPrefixBGPBindingDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPBindingDeleteResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    addressPrefixBGPBindingDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixBGPBindingDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressPrefixBGPBindingDeleteResponseEnvelopeMessages]
type addressPrefixBGPBindingDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixBGPBindingDeleteResponseEnvelopeSuccess bool

const (
	AddressPrefixBGPBindingDeleteResponseEnvelopeSuccessTrue AddressPrefixBGPBindingDeleteResponseEnvelopeSuccess = true
)

type AddressPrefixBGPBindingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressPrefixBGPBindingGetResponseEnvelope struct {
	Errors   []AddressPrefixBGPBindingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixBGPBindingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixBGPBindingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixBGPBindingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixBGPBindingGetResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixBGPBindingGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressPrefixBGPBindingGetResponseEnvelope]
type addressPrefixBGPBindingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPBindingGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    addressPrefixBGPBindingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixBGPBindingGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressPrefixBGPBindingGetResponseEnvelopeErrors]
type addressPrefixBGPBindingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPBindingGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressPrefixBGPBindingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixBGPBindingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressPrefixBGPBindingGetResponseEnvelopeMessages]
type addressPrefixBGPBindingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPBindingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixBGPBindingGetResponseEnvelopeSuccess bool

const (
	AddressPrefixBGPBindingGetResponseEnvelopeSuccessTrue AddressPrefixBGPBindingGetResponseEnvelopeSuccess = true
)
