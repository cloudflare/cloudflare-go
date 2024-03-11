// File generated from our OpenAPI spec by Stainless.

package addressing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
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
func (r *PrefixBGPBindingService) New(ctx context.Context, prefixID string, params PrefixBGPBindingNewParams, opts ...option.RequestOption) (res *AddressingServiceBinding, err error) {
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
func (r *PrefixBGPBindingService) List(ctx context.Context, prefixID string, query PrefixBGPBindingListParams, opts ...option.RequestOption) (res *[]AddressingServiceBinding, err error) {
	opts = append(r.Options[:], opts...)
	var env PrefixBGPBindingListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings", query.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Service Binding
func (r *PrefixBGPBindingService) Delete(ctx context.Context, prefixID string, bindingID string, body PrefixBGPBindingDeleteParams, opts ...option.RequestOption) (res *PrefixBGPBindingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PrefixBGPBindingDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", body.AccountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single Service Binding
func (r *PrefixBGPBindingService) Get(ctx context.Context, prefixID string, bindingID string, query PrefixBGPBindingGetParams, opts ...option.RequestOption) (res *AddressingServiceBinding, err error) {
	opts = append(r.Options[:], opts...)
	var env PrefixBGPBindingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bindings/%s", query.AccountID, prefixID, bindingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressingServiceBinding struct {
	// Identifier
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR string `json:"cidr"`
	// Status of a Service Binding's deployment to the Cloudflare network
	Provisioning AddressingServiceBindingProvisioning `json:"provisioning"`
	// Identifier
	ServiceID string `json:"service_id"`
	// Name of a service running on the Cloudflare network
	ServiceName string                       `json:"service_name"`
	JSON        addressingServiceBindingJSON `json:"-"`
}

// addressingServiceBindingJSON contains the JSON metadata for the struct
// [AddressingServiceBinding]
type addressingServiceBindingJSON struct {
	ID           apijson.Field
	CIDR         apijson.Field
	Provisioning apijson.Field
	ServiceID    apijson.Field
	ServiceName  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressingServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingServiceBindingJSON) RawJSON() string {
	return r.raw
}

// Status of a Service Binding's deployment to the Cloudflare network
type AddressingServiceBindingProvisioning struct {
	// When a binding has been deployed to a majority of Cloudflare datacenters, the
	// binding will become active and can be used with its associated service.
	State AddressingServiceBindingProvisioningState `json:"state"`
	JSON  addressingServiceBindingProvisioningJSON  `json:"-"`
}

// addressingServiceBindingProvisioningJSON contains the JSON metadata for the
// struct [AddressingServiceBindingProvisioning]
type addressingServiceBindingProvisioningJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingServiceBindingProvisioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingServiceBindingProvisioningJSON) RawJSON() string {
	return r.raw
}

// When a binding has been deployed to a majority of Cloudflare datacenters, the
// binding will become active and can be used with its associated service.
type AddressingServiceBindingProvisioningState string

const (
	AddressingServiceBindingProvisioningStateProvisioning AddressingServiceBindingProvisioningState = "provisioning"
	AddressingServiceBindingProvisioningStateActive       AddressingServiceBindingProvisioningState = "active"
)

// Union satisfied by [addressing.PrefixBGPBindingDeleteResponseUnknown],
// [addressing.PrefixBGPBindingDeleteResponseArray] or [shared.UnionString].
type PrefixBGPBindingDeleteResponse interface {
	ImplementsAddressingPrefixBGPBindingDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PrefixBGPBindingDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PrefixBGPBindingDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PrefixBGPBindingDeleteResponseArray []interface{}

func (r PrefixBGPBindingDeleteResponseArray) ImplementsAddressingPrefixBGPBindingDeleteResponse() {}

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
	Errors   []PrefixBGPBindingNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixBGPBindingNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingServiceBinding                      `json:"result,required"`
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

type PrefixBGPBindingNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    prefixBGPBindingNewResponseEnvelopeErrorsJSON `json:"-"`
}

// prefixBGPBindingNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrefixBGPBindingNewResponseEnvelopeErrors]
type prefixBGPBindingNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPBindingNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    prefixBGPBindingNewResponseEnvelopeMessagesJSON `json:"-"`
}

// prefixBGPBindingNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PrefixBGPBindingNewResponseEnvelopeMessages]
type prefixBGPBindingNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPBindingNewResponseEnvelopeSuccess bool

const (
	PrefixBGPBindingNewResponseEnvelopeSuccessTrue PrefixBGPBindingNewResponseEnvelopeSuccess = true
)

type PrefixBGPBindingListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixBGPBindingListResponseEnvelope struct {
	Errors   []PrefixBGPBindingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixBGPBindingListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressingServiceBinding                     `json:"result,required"`
	// Whether the API call was successful
	Success PrefixBGPBindingListResponseEnvelopeSuccess `json:"success,required"`
	JSON    prefixBGPBindingListResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPBindingListResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixBGPBindingListResponseEnvelope]
type prefixBGPBindingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPBindingListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    prefixBGPBindingListResponseEnvelopeErrorsJSON `json:"-"`
}

// prefixBGPBindingListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PrefixBGPBindingListResponseEnvelopeErrors]
type prefixBGPBindingListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPBindingListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    prefixBGPBindingListResponseEnvelopeMessagesJSON `json:"-"`
}

// prefixBGPBindingListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PrefixBGPBindingListResponseEnvelopeMessages]
type prefixBGPBindingListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPBindingListResponseEnvelopeSuccess bool

const (
	PrefixBGPBindingListResponseEnvelopeSuccessTrue PrefixBGPBindingListResponseEnvelopeSuccess = true
)

type PrefixBGPBindingDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixBGPBindingDeleteResponseEnvelope struct {
	Errors   []PrefixBGPBindingDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixBGPBindingDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   PrefixBGPBindingDeleteResponse                   `json:"result,required"`
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

type PrefixBGPBindingDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    prefixBGPBindingDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// prefixBGPBindingDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PrefixBGPBindingDeleteResponseEnvelopeErrors]
type prefixBGPBindingDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPBindingDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    prefixBGPBindingDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// prefixBGPBindingDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [PrefixBGPBindingDeleteResponseEnvelopeMessages]
type prefixBGPBindingDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPBindingDeleteResponseEnvelopeSuccess bool

const (
	PrefixBGPBindingDeleteResponseEnvelopeSuccessTrue PrefixBGPBindingDeleteResponseEnvelopeSuccess = true
)

type PrefixBGPBindingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixBGPBindingGetResponseEnvelope struct {
	Errors   []PrefixBGPBindingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixBGPBindingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingServiceBinding                      `json:"result,required"`
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

type PrefixBGPBindingGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    prefixBGPBindingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// prefixBGPBindingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrefixBGPBindingGetResponseEnvelopeErrors]
type prefixBGPBindingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPBindingGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    prefixBGPBindingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// prefixBGPBindingGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PrefixBGPBindingGetResponseEnvelopeMessages]
type prefixBGPBindingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPBindingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPBindingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPBindingGetResponseEnvelopeSuccess bool

const (
	PrefixBGPBindingGetResponseEnvelopeSuccessTrue PrefixBGPBindingGetResponseEnvelopeSuccess = true
)
