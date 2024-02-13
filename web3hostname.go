// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// Web3HostnameService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWeb3HostnameService] method
// instead.
type Web3HostnameService struct {
	Options            []option.RequestOption
	IpfsUniversalPaths *Web3HostnameIpfsUniversalPathService
}

// NewWeb3HostnameService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWeb3HostnameService(opts ...option.RequestOption) (r *Web3HostnameService) {
	r = &Web3HostnameService{}
	r.Options = opts
	r.IpfsUniversalPaths = NewWeb3HostnameIpfsUniversalPathService(opts...)
	return
}

// Edit Web3 Hostname
func (r *Web3HostnameService) Update(ctx context.Context, zoneIdentifier string, identifier string, body Web3HostnameUpdateParams, opts ...option.RequestOption) (res *Web3HostnameUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Web3 Hostname
func (r *Web3HostnameService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *Web3HostnameDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Web3 Hostname Details
func (r *Web3HostnameService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *Web3HostnameGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create Web3 Hostname
func (r *Web3HostnameService) Web3HostnameNewWeb3Hostname(ctx context.Context, zoneIdentifier string, body Web3HostnameWeb3HostnameNewWeb3HostnameParams, opts ...option.RequestOption) (res *Web3HostnameWeb3HostnameNewWeb3HostnameResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Web3 Hostnames
func (r *Web3HostnameService) Web3HostnameListWeb3Hostnames(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]Web3HostnameWeb3HostnameListWeb3HostnamesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Web3HostnameUpdateResponse struct {
	// Identifier
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the hostname.
	Description string `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink    string    `json:"dnslink"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The hostname that will point to the target gateway via CNAME.
	Name string `json:"name"`
	// Status of the hostname's activation.
	Status Web3HostnameUpdateResponseStatus `json:"status"`
	// Target gateway of the hostname.
	Target Web3HostnameUpdateResponseTarget `json:"target"`
	JSON   web3HostnameUpdateResponseJSON   `json:"-"`
}

// web3HostnameUpdateResponseJSON contains the JSON metadata for the struct
// [Web3HostnameUpdateResponse]
type web3HostnameUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	Dnslink     apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type Web3HostnameUpdateResponseStatus string

const (
	Web3HostnameUpdateResponseStatusActive   Web3HostnameUpdateResponseStatus = "active"
	Web3HostnameUpdateResponseStatusPending  Web3HostnameUpdateResponseStatus = "pending"
	Web3HostnameUpdateResponseStatusDeleting Web3HostnameUpdateResponseStatus = "deleting"
	Web3HostnameUpdateResponseStatusError    Web3HostnameUpdateResponseStatus = "error"
)

// Target gateway of the hostname.
type Web3HostnameUpdateResponseTarget string

const (
	Web3HostnameUpdateResponseTargetEthereum          Web3HostnameUpdateResponseTarget = "ethereum"
	Web3HostnameUpdateResponseTargetIpfs              Web3HostnameUpdateResponseTarget = "ipfs"
	Web3HostnameUpdateResponseTargetIpfsUniversalPath Web3HostnameUpdateResponseTarget = "ipfs_universal_path"
)

type Web3HostnameDeleteResponse struct {
	// Identifier
	ID   string                         `json:"id,required"`
	JSON web3HostnameDeleteResponseJSON `json:"-"`
}

// web3HostnameDeleteResponseJSON contains the JSON metadata for the struct
// [Web3HostnameDeleteResponse]
type web3HostnameDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameGetResponse struct {
	// Identifier
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the hostname.
	Description string `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink    string    `json:"dnslink"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The hostname that will point to the target gateway via CNAME.
	Name string `json:"name"`
	// Status of the hostname's activation.
	Status Web3HostnameGetResponseStatus `json:"status"`
	// Target gateway of the hostname.
	Target Web3HostnameGetResponseTarget `json:"target"`
	JSON   web3HostnameGetResponseJSON   `json:"-"`
}

// web3HostnameGetResponseJSON contains the JSON metadata for the struct
// [Web3HostnameGetResponse]
type web3HostnameGetResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	Dnslink     apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type Web3HostnameGetResponseStatus string

const (
	Web3HostnameGetResponseStatusActive   Web3HostnameGetResponseStatus = "active"
	Web3HostnameGetResponseStatusPending  Web3HostnameGetResponseStatus = "pending"
	Web3HostnameGetResponseStatusDeleting Web3HostnameGetResponseStatus = "deleting"
	Web3HostnameGetResponseStatusError    Web3HostnameGetResponseStatus = "error"
)

// Target gateway of the hostname.
type Web3HostnameGetResponseTarget string

const (
	Web3HostnameGetResponseTargetEthereum          Web3HostnameGetResponseTarget = "ethereum"
	Web3HostnameGetResponseTargetIpfs              Web3HostnameGetResponseTarget = "ipfs"
	Web3HostnameGetResponseTargetIpfsUniversalPath Web3HostnameGetResponseTarget = "ipfs_universal_path"
)

type Web3HostnameWeb3HostnameNewWeb3HostnameResponse struct {
	// Identifier
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the hostname.
	Description string `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink    string    `json:"dnslink"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The hostname that will point to the target gateway via CNAME.
	Name string `json:"name"`
	// Status of the hostname's activation.
	Status Web3HostnameWeb3HostnameNewWeb3HostnameResponseStatus `json:"status"`
	// Target gateway of the hostname.
	Target Web3HostnameWeb3HostnameNewWeb3HostnameResponseTarget `json:"target"`
	JSON   web3HostnameWeb3HostnameNewWeb3HostnameResponseJSON   `json:"-"`
}

// web3HostnameWeb3HostnameNewWeb3HostnameResponseJSON contains the JSON metadata
// for the struct [Web3HostnameWeb3HostnameNewWeb3HostnameResponse]
type web3HostnameWeb3HostnameNewWeb3HostnameResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	Dnslink     apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameWeb3HostnameNewWeb3HostnameResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type Web3HostnameWeb3HostnameNewWeb3HostnameResponseStatus string

const (
	Web3HostnameWeb3HostnameNewWeb3HostnameResponseStatusActive   Web3HostnameWeb3HostnameNewWeb3HostnameResponseStatus = "active"
	Web3HostnameWeb3HostnameNewWeb3HostnameResponseStatusPending  Web3HostnameWeb3HostnameNewWeb3HostnameResponseStatus = "pending"
	Web3HostnameWeb3HostnameNewWeb3HostnameResponseStatusDeleting Web3HostnameWeb3HostnameNewWeb3HostnameResponseStatus = "deleting"
	Web3HostnameWeb3HostnameNewWeb3HostnameResponseStatusError    Web3HostnameWeb3HostnameNewWeb3HostnameResponseStatus = "error"
)

// Target gateway of the hostname.
type Web3HostnameWeb3HostnameNewWeb3HostnameResponseTarget string

const (
	Web3HostnameWeb3HostnameNewWeb3HostnameResponseTargetEthereum          Web3HostnameWeb3HostnameNewWeb3HostnameResponseTarget = "ethereum"
	Web3HostnameWeb3HostnameNewWeb3HostnameResponseTargetIpfs              Web3HostnameWeb3HostnameNewWeb3HostnameResponseTarget = "ipfs"
	Web3HostnameWeb3HostnameNewWeb3HostnameResponseTargetIpfsUniversalPath Web3HostnameWeb3HostnameNewWeb3HostnameResponseTarget = "ipfs_universal_path"
)

type Web3HostnameWeb3HostnameListWeb3HostnamesResponse struct {
	// Identifier
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the hostname.
	Description string `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink    string    `json:"dnslink"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The hostname that will point to the target gateway via CNAME.
	Name string `json:"name"`
	// Status of the hostname's activation.
	Status Web3HostnameWeb3HostnameListWeb3HostnamesResponseStatus `json:"status"`
	// Target gateway of the hostname.
	Target Web3HostnameWeb3HostnameListWeb3HostnamesResponseTarget `json:"target"`
	JSON   web3HostnameWeb3HostnameListWeb3HostnamesResponseJSON   `json:"-"`
}

// web3HostnameWeb3HostnameListWeb3HostnamesResponseJSON contains the JSON metadata
// for the struct [Web3HostnameWeb3HostnameListWeb3HostnamesResponse]
type web3HostnameWeb3HostnameListWeb3HostnamesResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	Dnslink     apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameWeb3HostnameListWeb3HostnamesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type Web3HostnameWeb3HostnameListWeb3HostnamesResponseStatus string

const (
	Web3HostnameWeb3HostnameListWeb3HostnamesResponseStatusActive   Web3HostnameWeb3HostnameListWeb3HostnamesResponseStatus = "active"
	Web3HostnameWeb3HostnameListWeb3HostnamesResponseStatusPending  Web3HostnameWeb3HostnameListWeb3HostnamesResponseStatus = "pending"
	Web3HostnameWeb3HostnameListWeb3HostnamesResponseStatusDeleting Web3HostnameWeb3HostnameListWeb3HostnamesResponseStatus = "deleting"
	Web3HostnameWeb3HostnameListWeb3HostnamesResponseStatusError    Web3HostnameWeb3HostnameListWeb3HostnamesResponseStatus = "error"
)

// Target gateway of the hostname.
type Web3HostnameWeb3HostnameListWeb3HostnamesResponseTarget string

const (
	Web3HostnameWeb3HostnameListWeb3HostnamesResponseTargetEthereum          Web3HostnameWeb3HostnameListWeb3HostnamesResponseTarget = "ethereum"
	Web3HostnameWeb3HostnameListWeb3HostnamesResponseTargetIpfs              Web3HostnameWeb3HostnameListWeb3HostnamesResponseTarget = "ipfs"
	Web3HostnameWeb3HostnameListWeb3HostnamesResponseTargetIpfsUniversalPath Web3HostnameWeb3HostnameListWeb3HostnamesResponseTarget = "ipfs_universal_path"
)

type Web3HostnameUpdateParams struct {
	// An optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r Web3HostnameUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Web3HostnameUpdateResponseEnvelope struct {
	Errors   []Web3HostnameUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameUpdateResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [Web3HostnameUpdateResponseEnvelope]
type web3HostnameUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameUpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    web3HostnameUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [Web3HostnameUpdateResponseEnvelopeErrors]
type web3HostnameUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameUpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    web3HostnameUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [Web3HostnameUpdateResponseEnvelopeMessages]
type web3HostnameUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameUpdateResponseEnvelopeSuccess bool

const (
	Web3HostnameUpdateResponseEnvelopeSuccessTrue Web3HostnameUpdateResponseEnvelopeSuccess = true
)

type Web3HostnameDeleteResponseEnvelope struct {
	Errors   []Web3HostnameDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success Web3HostnameDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameDeleteResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [Web3HostnameDeleteResponseEnvelope]
type web3HostnameDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameDeleteResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    web3HostnameDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [Web3HostnameDeleteResponseEnvelopeErrors]
type web3HostnameDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameDeleteResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    web3HostnameDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [Web3HostnameDeleteResponseEnvelopeMessages]
type web3HostnameDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameDeleteResponseEnvelopeSuccess bool

const (
	Web3HostnameDeleteResponseEnvelopeSuccessTrue Web3HostnameDeleteResponseEnvelopeSuccess = true
)

type Web3HostnameGetResponseEnvelope struct {
	Errors   []Web3HostnameGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameGetResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameGetResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [Web3HostnameGetResponseEnvelope]
type web3HostnameGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    web3HostnameGetResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [Web3HostnameGetResponseEnvelopeErrors]
type web3HostnameGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    web3HostnameGetResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [Web3HostnameGetResponseEnvelopeMessages]
type web3HostnameGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameGetResponseEnvelopeSuccess bool

const (
	Web3HostnameGetResponseEnvelopeSuccessTrue Web3HostnameGetResponseEnvelopeSuccess = true
)

type Web3HostnameWeb3HostnameNewWeb3HostnameParams struct {
	// Target gateway of the hostname.
	Target param.Field[Web3HostnameWeb3HostnameNewWeb3HostnameParamsTarget] `json:"target,required"`
	// An optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r Web3HostnameWeb3HostnameNewWeb3HostnameParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Target gateway of the hostname.
type Web3HostnameWeb3HostnameNewWeb3HostnameParamsTarget string

const (
	Web3HostnameWeb3HostnameNewWeb3HostnameParamsTargetEthereum          Web3HostnameWeb3HostnameNewWeb3HostnameParamsTarget = "ethereum"
	Web3HostnameWeb3HostnameNewWeb3HostnameParamsTargetIpfs              Web3HostnameWeb3HostnameNewWeb3HostnameParamsTarget = "ipfs"
	Web3HostnameWeb3HostnameNewWeb3HostnameParamsTargetIpfsUniversalPath Web3HostnameWeb3HostnameNewWeb3HostnameParamsTarget = "ipfs_universal_path"
)

type Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelope struct {
	Errors   []Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameWeb3HostnameNewWeb3HostnameResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelope]
type web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeErrors]
type web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeMessages]
type web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeSuccess bool

const (
	Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeSuccessTrue Web3HostnameWeb3HostnameNewWeb3HostnameResponseEnvelopeSuccess = true
)

type Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelope struct {
	Errors   []Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeMessages `json:"messages,required"`
	Result   []Web3HostnameWeb3HostnameListWeb3HostnamesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeJSON       `json:"-"`
}

// web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelope]
type web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeErrors struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeErrors]
type web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeMessages struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeMessages]
type web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeSuccess bool

const (
	Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeSuccessTrue Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeSuccess = true
)

type Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                 `json:"total_count"`
	JSON       web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeResultInfoJSON `json:"-"`
}

// web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeResultInfo]
type web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameWeb3HostnameListWeb3HostnamesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
