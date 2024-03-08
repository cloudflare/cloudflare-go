// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// AddressingPrefixBGPBindingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAddressingPrefixBGPBindingService] method instead.
type AddressingPrefixBGPBindingService struct {
	Options []option.RequestOption
}

// NewAddressingPrefixBGPBindingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAddressingPrefixBGPBindingService(opts ...option.RequestOption) (r *AddressingPrefixBGPBindingService) {
	r = &AddressingPrefixBGPBindingService{}
	r.Options = opts
	return
}

// Creates a new Service Binding, routing traffic to IPs within the given CIDR to a
// service running on Cloudflare's network. **Note:** This API may only be used on
// prefixes currently configured with a Magic Transit service binding, and only
// allows creating service bindings for the Cloudflare CDN or Cloudflare Spectrum.
func (r *AddressingPrefixBGPBindingService) New(ctx context.Context, prefixID string, params AddressingPrefixBGPBindingNewParams, opts ...option.RequestOption) (res *AddressingPrefixBGPBindingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBGPBindingNewResponseEnvelope
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
func (r *AddressingPrefixBGPBindingService) List(ctx context.Context, prefixID string, query AddressingPrefixBGPBindingListParams, opts ...option.RequestOption) (res *[]AddressingPrefixBGPBindingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBGPBindingListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", query.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Service Binding
func (r *AddressingPrefixBGPBindingService) Delete(ctx context.Context, prefixID string, bindingID string, body AddressingPrefixBGPBindingDeleteParams, opts ...option.RequestOption) (res *AddressingPrefixBGPBindingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBGPBindingDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", body.AccountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single Service Binding
func (r *AddressingPrefixBGPBindingService) Get(ctx context.Context, prefixID string, bindingID string, query AddressingPrefixBGPBindingGetParams, opts ...option.RequestOption) (res *AddressingPrefixBGPBindingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBGPBindingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", query.AccountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressingPrefixBGPBindingNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AddressingPrefixBGPBindingNewResponseProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                    `json:"service_name"`
	JSON        addressingPrefixBGPBindingNewResponseJSON `json:"-"`
}

// addressingPrefixBGPBindingNewResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBGPBindingNewResponse]
type addressingPrefixBGPBindingNewResponseJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingNewResponseJSON) RawJSON() string {
	return r.raw
}

// Status of a Service Binding's deployment to the Cloudflare network
type AddressingPrefixBGPBindingNewResponseProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AddressingPrefixBGPBindingNewResponseProvisioningState `json:"state"`
	JSON  addressingPrefixBGPBindingNewResponseProvisioningJSON  `json:"-"`
}

// addressingPrefixBGPBindingNewResponseProvisioningJSON contains the JSON metadata
// for the struct [AddressingPrefixBGPBindingNewResponseProvisioning]
type addressingPrefixBGPBindingNewResponseProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingNewResponseProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingNewResponseProvisioningJSON) RawJSON() string {
	return r.raw
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AddressingPrefixBGPBindingNewResponseProvisioningState string

const (
	AddressingPrefixBGPBindingNewResponseProvisioningStateProvisioning AddressingPrefixBGPBindingNewResponseProvisioningState = "provisioning"
	AddressingPrefixBGPBindingNewResponseProvisioningStateActive       AddressingPrefixBGPBindingNewResponseProvisioningState = "active"
)

type AddressingPrefixBGPBindingListResponse struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AddressingPrefixBGPBindingListResponseProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                     `json:"service_name"`
	JSON        addressingPrefixBGPBindingListResponseJSON `json:"-"`
}

// addressingPrefixBGPBindingListResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBGPBindingListResponse]
type addressingPrefixBGPBindingListResponseJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingListResponseJSON) RawJSON() string {
	return r.raw
}

// Status of a Service Binding's deployment to the Cloudflare network
type AddressingPrefixBGPBindingListResponseProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AddressingPrefixBGPBindingListResponseProvisioningState `json:"state"`
	JSON  addressingPrefixBGPBindingListResponseProvisioningJSON  `json:"-"`
}

// addressingPrefixBGPBindingListResponseProvisioningJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPBindingListResponseProvisioning]
type addressingPrefixBGPBindingListResponseProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingListResponseProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingListResponseProvisioningJSON) RawJSON() string {
	return r.raw
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AddressingPrefixBGPBindingListResponseProvisioningState string

const (
	AddressingPrefixBGPBindingListResponseProvisioningStateProvisioning AddressingPrefixBGPBindingListResponseProvisioningState = "provisioning"
	AddressingPrefixBGPBindingListResponseProvisioningStateActive       AddressingPrefixBGPBindingListResponseProvisioningState = "active"
)

// Union satisfied by [AddressingPrefixBGPBindingDeleteResponseUnknown],
// [AddressingPrefixBGPBindingDeleteResponseArray] or [shared.UnionString].
type AddressingPrefixBGPBindingDeleteResponse interface {
	ImplementsAddressingPrefixBGPBindingDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressingPrefixBGPBindingDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressingPrefixBGPBindingDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressingPrefixBGPBindingDeleteResponseArray []interface{}

func (r AddressingPrefixBGPBindingDeleteResponseArray) ImplementsAddressingPrefixBGPBindingDeleteResponse() {
}

type AddressingPrefixBGPBindingGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AddressingPrefixBGPBindingGetResponseProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                                    `json:"service_name"`
	JSON        addressingPrefixBGPBindingGetResponseJSON `json:"-"`
}

// addressingPrefixBGPBindingGetResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBGPBindingGetResponse]
type addressingPrefixBGPBindingGetResponseJSON struct {
	ID           apijson.Field
	Cidr         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingGetResponseJSON) RawJSON() string {
	return r.raw
}

// Status of a Service Binding's deployment to the Cloudflare network
type AddressingPrefixBGPBindingGetResponseProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AddressingPrefixBGPBindingGetResponseProvisioningState `json:"state"`
	JSON  addressingPrefixBGPBindingGetResponseProvisioningJSON  `json:"-"`
}

// addressingPrefixBGPBindingGetResponseProvisioningJSON contains the JSON metadata
// for the struct [AddressingPrefixBGPBindingGetResponseProvisioning]
type addressingPrefixBGPBindingGetResponseProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingGetResponseProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingGetResponseProvisioningJSON) RawJSON() string {
	return r.raw
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AddressingPrefixBGPBindingGetResponseProvisioningState string

const (
	AddressingPrefixBGPBindingGetResponseProvisioningStateProvisioning AddressingPrefixBGPBindingGetResponseProvisioningState = "provisioning"
	AddressingPrefixBGPBindingGetResponseProvisioningStateActive       AddressingPrefixBGPBindingGetResponseProvisioningState = "active"
)

type AddressingPrefixBGPBindingNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr"`
	// Identifier
	ServiceID param.Field[string] `json:"service_id"`
}

func (r AddressingPrefixBGPBindingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressingPrefixBGPBindingNewResponseEnvelope struct {
	Errors   []AddressingPrefixBGPBindingNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixBGPBindingNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixBGPBindingNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixBGPBindingNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixBGPBindingNewResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBGPBindingNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixBGPBindingNewResponseEnvelope]
type addressingPrefixBGPBindingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPBindingNewResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressingPrefixBGPBindingNewResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBGPBindingNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPBindingNewResponseEnvelopeErrors]
type addressingPrefixBGPBindingNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPBindingNewResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    addressingPrefixBGPBindingNewResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBGPBindingNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPBindingNewResponseEnvelopeMessages]
type addressingPrefixBGPBindingNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingPrefixBGPBindingNewResponseEnvelopeSuccess bool

const (
	AddressingPrefixBGPBindingNewResponseEnvelopeSuccessTrue AddressingPrefixBGPBindingNewResponseEnvelopeSuccess = true
)

type AddressingPrefixBGPBindingListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingPrefixBGPBindingListResponseEnvelope struct {
	Errors   []AddressingPrefixBGPBindingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixBGPBindingListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressingPrefixBGPBindingListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixBGPBindingListResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixBGPBindingListResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBGPBindingListResponseEnvelopeJSON contains the JSON metadata
// for the struct [AddressingPrefixBGPBindingListResponseEnvelope]
type addressingPrefixBGPBindingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPBindingListResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    addressingPrefixBGPBindingListResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBGPBindingListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPBindingListResponseEnvelopeErrors]
type addressingPrefixBGPBindingListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPBindingListResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    addressingPrefixBGPBindingListResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBGPBindingListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPBindingListResponseEnvelopeMessages]
type addressingPrefixBGPBindingListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingPrefixBGPBindingListResponseEnvelopeSuccess bool

const (
	AddressingPrefixBGPBindingListResponseEnvelopeSuccessTrue AddressingPrefixBGPBindingListResponseEnvelopeSuccess = true
)

type AddressingPrefixBGPBindingDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingPrefixBGPBindingDeleteResponseEnvelope struct {
	Errors   []AddressingPrefixBGPBindingDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixBGPBindingDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixBGPBindingDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixBGPBindingDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixBGPBindingDeleteResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBGPBindingDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [AddressingPrefixBGPBindingDeleteResponseEnvelope]
type addressingPrefixBGPBindingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPBindingDeleteResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    addressingPrefixBGPBindingDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBGPBindingDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPBindingDeleteResponseEnvelopeErrors]
type addressingPrefixBGPBindingDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPBindingDeleteResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    addressingPrefixBGPBindingDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBGPBindingDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AddressingPrefixBGPBindingDeleteResponseEnvelopeMessages]
type addressingPrefixBGPBindingDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingPrefixBGPBindingDeleteResponseEnvelopeSuccess bool

const (
	AddressingPrefixBGPBindingDeleteResponseEnvelopeSuccessTrue AddressingPrefixBGPBindingDeleteResponseEnvelopeSuccess = true
)

type AddressingPrefixBGPBindingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingPrefixBGPBindingGetResponseEnvelope struct {
	Errors   []AddressingPrefixBGPBindingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixBGPBindingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixBGPBindingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixBGPBindingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixBGPBindingGetResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBGPBindingGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixBGPBindingGetResponseEnvelope]
type addressingPrefixBGPBindingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPBindingGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressingPrefixBGPBindingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBGPBindingGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPBindingGetResponseEnvelopeErrors]
type addressingPrefixBGPBindingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPBindingGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    addressingPrefixBGPBindingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBGPBindingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPBindingGetResponseEnvelopeMessages]
type addressingPrefixBGPBindingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPBindingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPBindingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingPrefixBGPBindingGetResponseEnvelopeSuccess bool

const (
	AddressingPrefixBGPBindingGetResponseEnvelopeSuccessTrue AddressingPrefixBGPBindingGetResponseEnvelopeSuccess = true
)
