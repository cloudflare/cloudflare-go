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
func (r *AddressingPrefixBindingService) New(ctx context.Context, accountIdentifier string, prefixIdentifier string, body AddressingPrefixBindingNewParams, opts ...option.RequestOption) (res *AddressingPrefixBindingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a single Service Binding
func (r *AddressingPrefixBindingService) Get(ctx context.Context, accountIdentifier string, prefixIdentifier string, bindingIdentifier string, opts ...option.RequestOption) (res *AddressingPrefixBindingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", accountIdentifier, prefixIdentifier, bindingIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the Cloudflare services this prefix is currently bound to. Traffic sent to
// an address within an IP prefix will be routed to the Cloudflare service of the
// most-specific Service Binding matching the address. **Example:** binding
// `192.0.2.0/24` to Cloudflare Magic Transit and `192.0.2.1/32` to the Cloudflare
// CDN would route traffic for `192.0.2.1` to the CDN, and traffic for all other
// IPs in the prefix to Cloudflare Magic Transit.
func (r *AddressingPrefixBindingService) List(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *AddressingPrefixBindingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Service Binding
func (r *AddressingPrefixBindingService) Delete(ctx context.Context, accountIdentifier string, prefixIdentifier string, bindingIdentifier string, opts ...option.RequestOption) (res *AddressingPrefixBindingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", accountIdentifier, prefixIdentifier, bindingIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AddressingPrefixBindingNewResponse struct {
	Errors   []AddressingPrefixBindingNewResponseError   `json:"errors"`
	Messages []AddressingPrefixBindingNewResponseMessage `json:"messages"`
	Result   AddressingPrefixBindingNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AddressingPrefixBindingNewResponseSuccess `json:"success"`
	JSON    addressingPrefixBindingNewResponseJSON    `json:"-"`
}

// addressingPrefixBindingNewResponseJSON contains the JSON metadata for the struct
// [AddressingPrefixBindingNewResponse]
type addressingPrefixBindingNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingNewResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    addressingPrefixBindingNewResponseErrorJSON `json:"-"`
}

// addressingPrefixBindingNewResponseErrorJSON contains the JSON metadata for the
// struct [AddressingPrefixBindingNewResponseError]
type addressingPrefixBindingNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingNewResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    addressingPrefixBindingNewResponseMessageJSON `json:"-"`
}

// addressingPrefixBindingNewResponseMessageJSON contains the JSON metadata for the
// struct [AddressingPrefixBindingNewResponseMessage]
type addressingPrefixBindingNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingNewResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AddressingPrefixBindingNewResponseResultProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                       `json:"service_name"`
	JSON        addressingPrefixBindingNewResponseResultJSON `json:"-"`
}

// addressingPrefixBindingNewResponseResultJSON contains the JSON metadata for the
// struct [AddressingPrefixBindingNewResponseResult]
type addressingPrefixBindingNewResponseResultJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressingPrefixBindingNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a Service Binding's deployment to the Cloudflare network
type AddressingPrefixBindingNewResponseResultProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AddressingPrefixBindingNewResponseResultProvisioningState `json:"state"`
	JSON  addressingPrefixBindingNewResponseResultProvisioningJSON  `json:"-"`
}

// addressingPrefixBindingNewResponseResultProvisioningJSON contains the JSON
// metadata for the struct [AddressingPrefixBindingNewResponseResultProvisioning]
type addressingPrefixBindingNewResponseResultProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingNewResponseResultProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AddressingPrefixBindingNewResponseResultProvisioningState string

const (
	AddressingPrefixBindingNewResponseResultProvisioningStateProvisioning AddressingPrefixBindingNewResponseResultProvisioningState = "provisioning"
	AddressingPrefixBindingNewResponseResultProvisioningStateActive       AddressingPrefixBindingNewResponseResultProvisioningState = "active"
)

// Whether the API call was successful
type AddressingPrefixBindingNewResponseSuccess bool

const (
	AddressingPrefixBindingNewResponseSuccessTrue AddressingPrefixBindingNewResponseSuccess = true
)

type AddressingPrefixBindingGetResponse struct {
	Errors   []AddressingPrefixBindingGetResponseError   `json:"errors"`
	Messages []AddressingPrefixBindingGetResponseMessage `json:"messages"`
	Result   AddressingPrefixBindingGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AddressingPrefixBindingGetResponseSuccess `json:"success"`
	JSON    addressingPrefixBindingGetResponseJSON    `json:"-"`
}

// addressingPrefixBindingGetResponseJSON contains the JSON metadata for the struct
// [AddressingPrefixBindingGetResponse]
type addressingPrefixBindingGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingGetResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    addressingPrefixBindingGetResponseErrorJSON `json:"-"`
}

// addressingPrefixBindingGetResponseErrorJSON contains the JSON metadata for the
// struct [AddressingPrefixBindingGetResponseError]
type addressingPrefixBindingGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingGetResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    addressingPrefixBindingGetResponseMessageJSON `json:"-"`
}

// addressingPrefixBindingGetResponseMessageJSON contains the JSON metadata for the
// struct [AddressingPrefixBindingGetResponseMessage]
type addressingPrefixBindingGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingGetResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AddressingPrefixBindingGetResponseResultProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                       `json:"service_name"`
	JSON        addressingPrefixBindingGetResponseResultJSON `json:"-"`
}

// addressingPrefixBindingGetResponseResultJSON contains the JSON metadata for the
// struct [AddressingPrefixBindingGetResponseResult]
type addressingPrefixBindingGetResponseResultJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressingPrefixBindingGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a Service Binding's deployment to the Cloudflare network
type AddressingPrefixBindingGetResponseResultProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AddressingPrefixBindingGetResponseResultProvisioningState `json:"state"`
	JSON  addressingPrefixBindingGetResponseResultProvisioningJSON  `json:"-"`
}

// addressingPrefixBindingGetResponseResultProvisioningJSON contains the JSON
// metadata for the struct [AddressingPrefixBindingGetResponseResultProvisioning]
type addressingPrefixBindingGetResponseResultProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingGetResponseResultProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AddressingPrefixBindingGetResponseResultProvisioningState string

const (
	AddressingPrefixBindingGetResponseResultProvisioningStateProvisioning AddressingPrefixBindingGetResponseResultProvisioningState = "provisioning"
	AddressingPrefixBindingGetResponseResultProvisioningStateActive       AddressingPrefixBindingGetResponseResultProvisioningState = "active"
)

// Whether the API call was successful
type AddressingPrefixBindingGetResponseSuccess bool

const (
	AddressingPrefixBindingGetResponseSuccessTrue AddressingPrefixBindingGetResponseSuccess = true
)

type AddressingPrefixBindingListResponse struct {
	Errors   []AddressingPrefixBindingListResponseError   `json:"errors"`
	Messages []AddressingPrefixBindingListResponseMessage `json:"messages"`
	Result   []AddressingPrefixBindingListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AddressingPrefixBindingListResponseSuccess `json:"success"`
	JSON    addressingPrefixBindingListResponseJSON    `json:"-"`
}

// addressingPrefixBindingListResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBindingListResponse]
type addressingPrefixBindingListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingListResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    addressingPrefixBindingListResponseErrorJSON `json:"-"`
}

// addressingPrefixBindingListResponseErrorJSON contains the JSON metadata for the
// struct [AddressingPrefixBindingListResponseError]
type addressingPrefixBindingListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingListResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    addressingPrefixBindingListResponseMessageJSON `json:"-"`
}

// addressingPrefixBindingListResponseMessageJSON contains the JSON metadata for
// the struct [AddressingPrefixBindingListResponseMessage]
type addressingPrefixBindingListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingListResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AddressingPrefixBindingListResponseResultProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                        `json:"service_name"`
	JSON        addressingPrefixBindingListResponseResultJSON `json:"-"`
}

// addressingPrefixBindingListResponseResultJSON contains the JSON metadata for the
// struct [AddressingPrefixBindingListResponseResult]
type addressingPrefixBindingListResponseResultJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressingPrefixBindingListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a Service Binding's deployment to the Cloudflare network
type AddressingPrefixBindingListResponseResultProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AddressingPrefixBindingListResponseResultProvisioningState `json:"state"`
	JSON  addressingPrefixBindingListResponseResultProvisioningJSON  `json:"-"`
}

// addressingPrefixBindingListResponseResultProvisioningJSON contains the JSON
// metadata for the struct [AddressingPrefixBindingListResponseResultProvisioning]
type addressingPrefixBindingListResponseResultProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingListResponseResultProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AddressingPrefixBindingListResponseResultProvisioningState string

const (
	AddressingPrefixBindingListResponseResultProvisioningStateProvisioning AddressingPrefixBindingListResponseResultProvisioningState = "provisioning"
	AddressingPrefixBindingListResponseResultProvisioningStateActive       AddressingPrefixBindingListResponseResultProvisioningState = "active"
)

// Whether the API call was successful
type AddressingPrefixBindingListResponseSuccess bool

const (
	AddressingPrefixBindingListResponseSuccessTrue AddressingPrefixBindingListResponseSuccess = true
)

type AddressingPrefixBindingDeleteResponse struct {
	Errors   []AddressingPrefixBindingDeleteResponseError   `json:"errors,required"`
	Messages []AddressingPrefixBindingDeleteResponseMessage `json:"messages,required"`
	Result   AddressingPrefixBindingDeleteResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixBindingDeleteResponseSuccess `json:"success,required"`
	JSON    addressingPrefixBindingDeleteResponseJSON    `json:"-"`
}

// addressingPrefixBindingDeleteResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBindingDeleteResponse]
type addressingPrefixBindingDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingDeleteResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    addressingPrefixBindingDeleteResponseErrorJSON `json:"-"`
}

// addressingPrefixBindingDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AddressingPrefixBindingDeleteResponseError]
type addressingPrefixBindingDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBindingDeleteResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    addressingPrefixBindingDeleteResponseMessageJSON `json:"-"`
}

// addressingPrefixBindingDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AddressingPrefixBindingDeleteResponseMessage]
type addressingPrefixBindingDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBindingDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AddressingPrefixBindingDeleteResponseResultUnknown],
// [AddressingPrefixBindingDeleteResponseResultArray] or [shared.UnionString].
type AddressingPrefixBindingDeleteResponseResult interface {
	ImplementsAddressingPrefixBindingDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressingPrefixBindingDeleteResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressingPrefixBindingDeleteResponseResultArray []interface{}

func (r AddressingPrefixBindingDeleteResponseResultArray) ImplementsAddressingPrefixBindingDeleteResponseResult() {
}

// Whether the API call was successful
type AddressingPrefixBindingDeleteResponseSuccess bool

const (
	AddressingPrefixBindingDeleteResponseSuccessTrue AddressingPrefixBindingDeleteResponseSuccess = true
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
