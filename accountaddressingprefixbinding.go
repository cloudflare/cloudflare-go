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

// AccountAddressingPrefixBindingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAddressingPrefixBindingService] method instead.
type AccountAddressingPrefixBindingService struct {
	Options []option.RequestOption
}

// NewAccountAddressingPrefixBindingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressingPrefixBindingService(opts ...option.RequestOption) (r *AccountAddressingPrefixBindingService) {
	r = &AccountAddressingPrefixBindingService{}
	r.Options = opts
	return
}

// Creates a new Service Binding, routing traffic to IPs within the given CIDR to a
// service running on Cloudflare's network. **Note:** This API may only be used on
// prefixes currently configured with a Magic Transit service binding, and only
// allows creating service bindings for the Cloudflare CDN or Cloudflare Spectrum.
func (r *AccountAddressingPrefixBindingService) New(ctx context.Context, accountIdentifier string, prefixIdentifier string, body AccountAddressingPrefixBindingNewParams, opts ...option.RequestOption) (res *AccountAddressingPrefixBindingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a single Service Binding
func (r *AccountAddressingPrefixBindingService) Get(ctx context.Context, accountIdentifier string, prefixIdentifier string, bindingIdentifier string, opts ...option.RequestOption) (res *AccountAddressingPrefixBindingGetResponse, err error) {
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
func (r *AccountAddressingPrefixBindingService) List(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *AccountAddressingPrefixBindingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Service Binding
func (r *AccountAddressingPrefixBindingService) Delete(ctx context.Context, accountIdentifier string, prefixIdentifier string, bindingIdentifier string, opts ...option.RequestOption) (res *AccountAddressingPrefixBindingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", accountIdentifier, prefixIdentifier, bindingIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountAddressingPrefixBindingNewResponse struct {
	Errors   []AccountAddressingPrefixBindingNewResponseError   `json:"errors"`
	Messages []AccountAddressingPrefixBindingNewResponseMessage `json:"messages"`
	Result   AccountAddressingPrefixBindingNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAddressingPrefixBindingNewResponseSuccess `json:"success"`
	JSON    accountAddressingPrefixBindingNewResponseJSON    `json:"-"`
}

// accountAddressingPrefixBindingNewResponseJSON contains the JSON metadata for the
// struct [AccountAddressingPrefixBindingNewResponse]
type accountAddressingPrefixBindingNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingPrefixBindingNewResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accountAddressingPrefixBindingNewResponseErrorJSON `json:"-"`
}

// accountAddressingPrefixBindingNewResponseErrorJSON contains the JSON metadata
// for the struct [AccountAddressingPrefixBindingNewResponseError]
type accountAddressingPrefixBindingNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingPrefixBindingNewResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountAddressingPrefixBindingNewResponseMessageJSON `json:"-"`
}

// accountAddressingPrefixBindingNewResponseMessageJSON contains the JSON metadata
// for the struct [AccountAddressingPrefixBindingNewResponseMessage]
type accountAddressingPrefixBindingNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingPrefixBindingNewResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AccountAddressingPrefixBindingNewResponseResultProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                              `json:"service_name"`
	JSON        accountAddressingPrefixBindingNewResponseResultJSON `json:"-"`
}

// accountAddressingPrefixBindingNewResponseResultJSON contains the JSON metadata
// for the struct [AccountAddressingPrefixBindingNewResponseResult]
type accountAddressingPrefixBindingNewResponseResultJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a Service Binding's deployment to the Cloudflare network
type AccountAddressingPrefixBindingNewResponseResultProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AccountAddressingPrefixBindingNewResponseResultProvisioningState `json:"state"`
	JSON  accountAddressingPrefixBindingNewResponseResultProvisioningJSON  `json:"-"`
}

// accountAddressingPrefixBindingNewResponseResultProvisioningJSON contains the
// JSON metadata for the struct
// [AccountAddressingPrefixBindingNewResponseResultProvisioning]
type accountAddressingPrefixBindingNewResponseResultProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingNewResponseResultProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AccountAddressingPrefixBindingNewResponseResultProvisioningState string

const (
	AccountAddressingPrefixBindingNewResponseResultProvisioningStateProvisioning AccountAddressingPrefixBindingNewResponseResultProvisioningState = "provisioning"
	AccountAddressingPrefixBindingNewResponseResultProvisioningStateActive       AccountAddressingPrefixBindingNewResponseResultProvisioningState = "active"
)

// Whether the API call was successful
type AccountAddressingPrefixBindingNewResponseSuccess bool

const (
	AccountAddressingPrefixBindingNewResponseSuccessTrue AccountAddressingPrefixBindingNewResponseSuccess = true
)

type AccountAddressingPrefixBindingGetResponse struct {
	Errors   []AccountAddressingPrefixBindingGetResponseError   `json:"errors"`
	Messages []AccountAddressingPrefixBindingGetResponseMessage `json:"messages"`
	Result   AccountAddressingPrefixBindingGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAddressingPrefixBindingGetResponseSuccess `json:"success"`
	JSON    accountAddressingPrefixBindingGetResponseJSON    `json:"-"`
}

// accountAddressingPrefixBindingGetResponseJSON contains the JSON metadata for the
// struct [AccountAddressingPrefixBindingGetResponse]
type accountAddressingPrefixBindingGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingPrefixBindingGetResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accountAddressingPrefixBindingGetResponseErrorJSON `json:"-"`
}

// accountAddressingPrefixBindingGetResponseErrorJSON contains the JSON metadata
// for the struct [AccountAddressingPrefixBindingGetResponseError]
type accountAddressingPrefixBindingGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingPrefixBindingGetResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountAddressingPrefixBindingGetResponseMessageJSON `json:"-"`
}

// accountAddressingPrefixBindingGetResponseMessageJSON contains the JSON metadata
// for the struct [AccountAddressingPrefixBindingGetResponseMessage]
type accountAddressingPrefixBindingGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingPrefixBindingGetResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AccountAddressingPrefixBindingGetResponseResultProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                              `json:"service_name"`
	JSON        accountAddressingPrefixBindingGetResponseResultJSON `json:"-"`
}

// accountAddressingPrefixBindingGetResponseResultJSON contains the JSON metadata
// for the struct [AccountAddressingPrefixBindingGetResponseResult]
type accountAddressingPrefixBindingGetResponseResultJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a Service Binding's deployment to the Cloudflare network
type AccountAddressingPrefixBindingGetResponseResultProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AccountAddressingPrefixBindingGetResponseResultProvisioningState `json:"state"`
	JSON  accountAddressingPrefixBindingGetResponseResultProvisioningJSON  `json:"-"`
}

// accountAddressingPrefixBindingGetResponseResultProvisioningJSON contains the
// JSON metadata for the struct
// [AccountAddressingPrefixBindingGetResponseResultProvisioning]
type accountAddressingPrefixBindingGetResponseResultProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingGetResponseResultProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AccountAddressingPrefixBindingGetResponseResultProvisioningState string

const (
	AccountAddressingPrefixBindingGetResponseResultProvisioningStateProvisioning AccountAddressingPrefixBindingGetResponseResultProvisioningState = "provisioning"
	AccountAddressingPrefixBindingGetResponseResultProvisioningStateActive       AccountAddressingPrefixBindingGetResponseResultProvisioningState = "active"
)

// Whether the API call was successful
type AccountAddressingPrefixBindingGetResponseSuccess bool

const (
	AccountAddressingPrefixBindingGetResponseSuccessTrue AccountAddressingPrefixBindingGetResponseSuccess = true
)

type AccountAddressingPrefixBindingListResponse struct {
	Errors   []AccountAddressingPrefixBindingListResponseError   `json:"errors"`
	Messages []AccountAddressingPrefixBindingListResponseMessage `json:"messages"`
	Result   []AccountAddressingPrefixBindingListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountAddressingPrefixBindingListResponseSuccess `json:"success"`
	JSON    accountAddressingPrefixBindingListResponseJSON    `json:"-"`
}

// accountAddressingPrefixBindingListResponseJSON contains the JSON metadata for
// the struct [AccountAddressingPrefixBindingListResponse]
type accountAddressingPrefixBindingListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingPrefixBindingListResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountAddressingPrefixBindingListResponseErrorJSON `json:"-"`
}

// accountAddressingPrefixBindingListResponseErrorJSON contains the JSON metadata
// for the struct [AccountAddressingPrefixBindingListResponseError]
type accountAddressingPrefixBindingListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingPrefixBindingListResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountAddressingPrefixBindingListResponseMessageJSON `json:"-"`
}

// accountAddressingPrefixBindingListResponseMessageJSON contains the JSON metadata
// for the struct [AccountAddressingPrefixBindingListResponseMessage]
type accountAddressingPrefixBindingListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingPrefixBindingListResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AccountAddressingPrefixBindingListResponseResultProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                               `json:"service_name"`
	JSON        accountAddressingPrefixBindingListResponseResultJSON `json:"-"`
}

// accountAddressingPrefixBindingListResponseResultJSON contains the JSON metadata
// for the struct [AccountAddressingPrefixBindingListResponseResult]
type accountAddressingPrefixBindingListResponseResultJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a Service Binding's deployment to the Cloudflare network
type AccountAddressingPrefixBindingListResponseResultProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AccountAddressingPrefixBindingListResponseResultProvisioningState `json:"state"`
	JSON  accountAddressingPrefixBindingListResponseResultProvisioningJSON  `json:"-"`
}

// accountAddressingPrefixBindingListResponseResultProvisioningJSON contains the
// JSON metadata for the struct
// [AccountAddressingPrefixBindingListResponseResultProvisioning]
type accountAddressingPrefixBindingListResponseResultProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingListResponseResultProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AccountAddressingPrefixBindingListResponseResultProvisioningState string

const (
	AccountAddressingPrefixBindingListResponseResultProvisioningStateProvisioning AccountAddressingPrefixBindingListResponseResultProvisioningState = "provisioning"
	AccountAddressingPrefixBindingListResponseResultProvisioningStateActive       AccountAddressingPrefixBindingListResponseResultProvisioningState = "active"
)

// Whether the API call was successful
type AccountAddressingPrefixBindingListResponseSuccess bool

const (
	AccountAddressingPrefixBindingListResponseSuccessTrue AccountAddressingPrefixBindingListResponseSuccess = true
)

type AccountAddressingPrefixBindingDeleteResponse struct {
	Errors   []AccountAddressingPrefixBindingDeleteResponseError   `json:"errors,required"`
	Messages []AccountAddressingPrefixBindingDeleteResponseMessage `json:"messages,required"`
	Result   AccountAddressingPrefixBindingDeleteResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success AccountAddressingPrefixBindingDeleteResponseSuccess `json:"success,required"`
	JSON    accountAddressingPrefixBindingDeleteResponseJSON    `json:"-"`
}

// accountAddressingPrefixBindingDeleteResponseJSON contains the JSON metadata for
// the struct [AccountAddressingPrefixBindingDeleteResponse]
type accountAddressingPrefixBindingDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingPrefixBindingDeleteResponseError struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountAddressingPrefixBindingDeleteResponseErrorJSON `json:"-"`
}

// accountAddressingPrefixBindingDeleteResponseErrorJSON contains the JSON metadata
// for the struct [AccountAddressingPrefixBindingDeleteResponseError]
type accountAddressingPrefixBindingDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingPrefixBindingDeleteResponseMessage struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accountAddressingPrefixBindingDeleteResponseMessageJSON `json:"-"`
}

// accountAddressingPrefixBindingDeleteResponseMessageJSON contains the JSON
// metadata for the struct [AccountAddressingPrefixBindingDeleteResponseMessage]
type accountAddressingPrefixBindingDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBindingDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccountAddressingPrefixBindingDeleteResponseResultUnknown],
// [AccountAddressingPrefixBindingDeleteResponseResultArray] or
// [shared.UnionString].
type AccountAddressingPrefixBindingDeleteResponseResult interface {
	ImplementsAccountAddressingPrefixBindingDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAddressingPrefixBindingDeleteResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAddressingPrefixBindingDeleteResponseResultArray []interface{}

func (r AccountAddressingPrefixBindingDeleteResponseResultArray) ImplementsAccountAddressingPrefixBindingDeleteResponseResult() {
}

// Whether the API call was successful
type AccountAddressingPrefixBindingDeleteResponseSuccess bool

const (
	AccountAddressingPrefixBindingDeleteResponseSuccessTrue AccountAddressingPrefixBindingDeleteResponseSuccess = true
)

type AccountAddressingPrefixBindingNewParams struct {
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr"`
	// Identifier
	ServiceID param.Field[string] `json:"service_id"`
}

func (r AccountAddressingPrefixBindingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
