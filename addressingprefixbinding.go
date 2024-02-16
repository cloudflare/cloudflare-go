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

// AddressingPrefixBindingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAddressingPrefixBindingService] method instead.
type AddressingPrefixBindingService struct {
	Options []option.RequestOption
}

// NewAddressingPrefixBindingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressingPrefixBindingService(opts ...option.RequestOption) (r *AddressingPrefixBindingService) {
	r = &AddressingPrefixBindingService{}
	r.Options = opts
	return
}

// Creates a new Service Binding, routing traffic to IPs within the given CIDR to a
// service running on Cloudflare's network. **Note:** This API may only be used on
// prefixes currently configured with a Magic Transit service binding, and only
// allows creating service bindings for the Cloudflare CDN or Cloudflare Spectrum.
func (r *AddressingPrefixBindingService) New(ctx context.Context, accountID string, prefixID string, body AddressingPrefixBindingNewParams, opts ...option.RequestOption) (res *AddressingPrefixBindingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBindingNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
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
func (r *AddressingPrefixBindingService) List(ctx context.Context, accountID string, prefixID string, opts ...option.RequestOption) (res *[]AddressingPrefixBindingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBindingListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Service Binding
func (r *AddressingPrefixBindingService) Delete(ctx context.Context, accountID string, prefixID string, bindingID string, opts ...option.RequestOption) (res *AddressingPrefixBindingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBindingDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", accountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single Service Binding
func (r *AddressingPrefixBindingService) Get(ctx context.Context, accountID string, prefixID string, bindingID string, opts ...option.RequestOption) (res *AddressingPrefixBindingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBindingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", accountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressingPrefixBindingNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AddressingPrefixBindingNewResponseProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                 `json:"service_name"`
	JSON        addressingPrefixBindingNewResponseJSON `json:"-"`
}

// addressingPrefixBindingNewResponseJSON contains the JSON metadata for the struct
// [AddressingPrefixBindingNewResponse]
type addressingPrefixBindingNewResponseJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressingPrefixBindingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a Service Binding's deployment to the Cloudflare network
type AddressingPrefixBindingNewResponseProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AddressingPrefixBindingNewResponseProvisioningState `json:"state"`
	JSON  addressingPrefixBindingNewResponseProvisioningJSON  `json:"-"`
}

// addressingPrefixBindingNewResponseProvisioningJSON contains the JSON metadata
// for the struct [AddressingPrefixBindingNewResponseProvisioning]
type addressingPrefixBindingNewResponseProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingNewResponseProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AddressingPrefixBindingNewResponseProvisioningState string

const (
	AddressingPrefixBindingNewResponseProvisioningStateProvisioning AddressingPrefixBindingNewResponseProvisioningState = "provisioning"
	AddressingPrefixBindingNewResponseProvisioningStateActive       AddressingPrefixBindingNewResponseProvisioningState = "active"
)

type AddressingPrefixBindingListResponse struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AddressingPrefixBindingListResponseProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                  `json:"service_name"`
	JSON        addressingPrefixBindingListResponseJSON `json:"-"`
}

// addressingPrefixBindingListResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBindingListResponse]
type addressingPrefixBindingListResponseJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressingPrefixBindingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a Service Binding's deployment to the Cloudflare network
type AddressingPrefixBindingListResponseProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AddressingPrefixBindingListResponseProvisioningState `json:"state"`
	JSON  addressingPrefixBindingListResponseProvisioningJSON  `json:"-"`
}

// addressingPrefixBindingListResponseProvisioningJSON contains the JSON metadata
// for the struct [AddressingPrefixBindingListResponseProvisioning]
type addressingPrefixBindingListResponseProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingListResponseProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AddressingPrefixBindingListResponseProvisioningState string

const (
	AddressingPrefixBindingListResponseProvisioningStateProvisioning AddressingPrefixBindingListResponseProvisioningState = "provisioning"
	AddressingPrefixBindingListResponseProvisioningStateActive       AddressingPrefixBindingListResponseProvisioningState = "active"
)

// Union satisfied by [AddressingPrefixBindingDeleteResponseUnknown],
// [AddressingPrefixBindingDeleteResponseArray] or [shared.UnionString].
type AddressingPrefixBindingDeleteResponse interface {
	ImplementsAddressingPrefixBindingDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressingPrefixBindingDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressingPrefixBindingDeleteResponseArray []interface{}

func (r AddressingPrefixBindingDeleteResponseArray) ImplementsAddressingPrefixBindingDeleteResponse() {
}

type AddressingPrefixBindingGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AddressingPrefixBindingGetResponseProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                 `json:"service_name"`
	JSON        addressingPrefixBindingGetResponseJSON `json:"-"`
}

// addressingPrefixBindingGetResponseJSON contains the JSON metadata for the struct
// [AddressingPrefixBindingGetResponse]
type addressingPrefixBindingGetResponseJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressingPrefixBindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a Service Binding's deployment to the Cloudflare network
type AddressingPrefixBindingGetResponseProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AddressingPrefixBindingGetResponseProvisioningState `json:"state"`
	JSON  addressingPrefixBindingGetResponseProvisioningJSON  `json:"-"`
}

// addressingPrefixBindingGetResponseProvisioningJSON contains the JSON metadata
// for the struct [AddressingPrefixBindingGetResponseProvisioning]
type addressingPrefixBindingGetResponseProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingGetResponseProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AddressingPrefixBindingGetResponseProvisioningState string

const (
	AddressingPrefixBindingGetResponseProvisioningStateProvisioning AddressingPrefixBindingGetResponseProvisioningState = "provisioning"
	AddressingPrefixBindingGetResponseProvisioningStateActive       AddressingPrefixBindingGetResponseProvisioningState = "active"
)

type AddressingPrefixBindingNewParams struct {
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr"`
	// Identifier
	ServiceID param.Field[string] `json:"service_id"`
}

func (r AddressingPrefixBindingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressingPrefixBindingNewResponseEnvelope struct {
	Errors   []AddressingPrefixBindingNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixBindingNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixBindingNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixBindingNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixBindingNewResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBindingNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixBindingNewResponseEnvelope]
type addressingPrefixBindingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingNewResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    addressingPrefixBindingNewResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBindingNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressingPrefixBindingNewResponseEnvelopeErrors]
type addressingPrefixBindingNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingNewResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressingPrefixBindingNewResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBindingNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBindingNewResponseEnvelopeMessages]
type addressingPrefixBindingNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixBindingNewResponseEnvelopeSuccess bool

const (
	AddressingPrefixBindingNewResponseEnvelopeSuccessTrue AddressingPrefixBindingNewResponseEnvelopeSuccess = true
)

type AddressingPrefixBindingListResponseEnvelope struct {
	Errors   []AddressingPrefixBindingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixBindingListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressingPrefixBindingListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixBindingListResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixBindingListResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBindingListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixBindingListResponseEnvelope]
type addressingPrefixBindingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingListResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    addressingPrefixBindingListResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBindingListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressingPrefixBindingListResponseEnvelopeErrors]
type addressingPrefixBindingListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingListResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressingPrefixBindingListResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBindingListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBindingListResponseEnvelopeMessages]
type addressingPrefixBindingListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixBindingListResponseEnvelopeSuccess bool

const (
	AddressingPrefixBindingListResponseEnvelopeSuccessTrue AddressingPrefixBindingListResponseEnvelopeSuccess = true
)

type AddressingPrefixBindingDeleteResponseEnvelope struct {
	Errors   []AddressingPrefixBindingDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixBindingDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixBindingDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixBindingDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixBindingDeleteResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBindingDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixBindingDeleteResponseEnvelope]
type addressingPrefixBindingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingDeleteResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressingPrefixBindingDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBindingDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixBindingDeleteResponseEnvelopeErrors]
type addressingPrefixBindingDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingDeleteResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    addressingPrefixBindingDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBindingDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBindingDeleteResponseEnvelopeMessages]
type addressingPrefixBindingDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixBindingDeleteResponseEnvelopeSuccess bool

const (
	AddressingPrefixBindingDeleteResponseEnvelopeSuccessTrue AddressingPrefixBindingDeleteResponseEnvelopeSuccess = true
)

type AddressingPrefixBindingGetResponseEnvelope struct {
	Errors   []AddressingPrefixBindingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixBindingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixBindingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixBindingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixBindingGetResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBindingGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixBindingGetResponseEnvelope]
type addressingPrefixBindingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    addressingPrefixBindingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBindingGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressingPrefixBindingGetResponseEnvelopeErrors]
type addressingPrefixBindingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressingPrefixBindingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBindingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBindingGetResponseEnvelopeMessages]
type addressingPrefixBindingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixBindingGetResponseEnvelopeSuccess bool

const (
	AddressingPrefixBindingGetResponseEnvelopeSuccessTrue AddressingPrefixBindingGetResponseEnvelopeSuccess = true
)
