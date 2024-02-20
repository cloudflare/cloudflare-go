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

// Create Web3 Hostname
func (r *Web3HostnameService) New(ctx context.Context, zoneIdentifier string, body Web3HostnameNewParams, opts ...option.RequestOption) (res *Web3HostnameNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

// List Web3 Hostnames
func (r *Web3HostnameService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]Web3HostnameListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameListResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

type Web3HostnameNewResponse struct {
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
	Status Web3HostnameNewResponseStatus `json:"status"`
	// Target gateway of the hostname.
	Target Web3HostnameNewResponseTarget `json:"target"`
	JSON   web3HostnameNewResponseJSON   `json:"-"`
}

// web3HostnameNewResponseJSON contains the JSON metadata for the struct
// [Web3HostnameNewResponse]
type web3HostnameNewResponseJSON struct {
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

func (r *Web3HostnameNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type Web3HostnameNewResponseStatus string

const (
	Web3HostnameNewResponseStatusActive   Web3HostnameNewResponseStatus = "active"
	Web3HostnameNewResponseStatusPending  Web3HostnameNewResponseStatus = "pending"
	Web3HostnameNewResponseStatusDeleting Web3HostnameNewResponseStatus = "deleting"
	Web3HostnameNewResponseStatusError    Web3HostnameNewResponseStatus = "error"
)

// Target gateway of the hostname.
type Web3HostnameNewResponseTarget string

const (
	Web3HostnameNewResponseTargetEthereum          Web3HostnameNewResponseTarget = "ethereum"
	Web3HostnameNewResponseTargetIpfs              Web3HostnameNewResponseTarget = "ipfs"
	Web3HostnameNewResponseTargetIpfsUniversalPath Web3HostnameNewResponseTarget = "ipfs_universal_path"
)

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

type Web3HostnameListResponse struct {
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
	Status Web3HostnameListResponseStatus `json:"status"`
	// Target gateway of the hostname.
	Target Web3HostnameListResponseTarget `json:"target"`
	JSON   web3HostnameListResponseJSON   `json:"-"`
}

// web3HostnameListResponseJSON contains the JSON metadata for the struct
// [Web3HostnameListResponse]
type web3HostnameListResponseJSON struct {
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

func (r *Web3HostnameListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type Web3HostnameListResponseStatus string

const (
	Web3HostnameListResponseStatusActive   Web3HostnameListResponseStatus = "active"
	Web3HostnameListResponseStatusPending  Web3HostnameListResponseStatus = "pending"
	Web3HostnameListResponseStatusDeleting Web3HostnameListResponseStatus = "deleting"
	Web3HostnameListResponseStatusError    Web3HostnameListResponseStatus = "error"
)

// Target gateway of the hostname.
type Web3HostnameListResponseTarget string

const (
	Web3HostnameListResponseTargetEthereum          Web3HostnameListResponseTarget = "ethereum"
	Web3HostnameListResponseTargetIpfs              Web3HostnameListResponseTarget = "ipfs"
	Web3HostnameListResponseTargetIpfsUniversalPath Web3HostnameListResponseTarget = "ipfs_universal_path"
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

type Web3HostnameNewParams struct {
	// Target gateway of the hostname.
	Target param.Field[Web3HostnameNewParamsTarget] `json:"target,required"`
	// An optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r Web3HostnameNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Target gateway of the hostname.
type Web3HostnameNewParamsTarget string

const (
	Web3HostnameNewParamsTargetEthereum          Web3HostnameNewParamsTarget = "ethereum"
	Web3HostnameNewParamsTargetIpfs              Web3HostnameNewParamsTarget = "ipfs"
	Web3HostnameNewParamsTargetIpfsUniversalPath Web3HostnameNewParamsTarget = "ipfs_universal_path"
)

type Web3HostnameNewResponseEnvelope struct {
	Errors   []Web3HostnameNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameNewResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameNewResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [Web3HostnameNewResponseEnvelope]
type web3HostnameNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    web3HostnameNewResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [Web3HostnameNewResponseEnvelopeErrors]
type web3HostnameNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    web3HostnameNewResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [Web3HostnameNewResponseEnvelopeMessages]
type web3HostnameNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameNewResponseEnvelopeSuccess bool

const (
	Web3HostnameNewResponseEnvelopeSuccessTrue Web3HostnameNewResponseEnvelopeSuccess = true
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

type Web3HostnameListResponseEnvelope struct {
	Errors   []Web3HostnameListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameListResponseEnvelopeMessages `json:"messages,required"`
	Result   []Web3HostnameListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    Web3HostnameListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo Web3HostnameListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       web3HostnameListResponseEnvelopeJSON       `json:"-"`
}

// web3HostnameListResponseEnvelopeJSON contains the JSON metadata for the struct
// [Web3HostnameListResponseEnvelope]
type web3HostnameListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameListResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    web3HostnameListResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [Web3HostnameListResponseEnvelopeErrors]
type web3HostnameListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameListResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    web3HostnameListResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [Web3HostnameListResponseEnvelopeMessages]
type web3HostnameListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameListResponseEnvelopeSuccess bool

const (
	Web3HostnameListResponseEnvelopeSuccessTrue Web3HostnameListResponseEnvelopeSuccess = true
)

type Web3HostnameListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       web3HostnameListResponseEnvelopeResultInfoJSON `json:"-"`
}

// web3HostnameListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [Web3HostnameListResponseEnvelopeResultInfo]
type web3HostnameListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
