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
	IPFSUniversalPaths *Web3HostnameIPFSUniversalPathService
}

// NewWeb3HostnameService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWeb3HostnameService(opts ...option.RequestOption) (r *Web3HostnameService) {
	r = &Web3HostnameService{}
	r.Options = opts
	r.IPFSUniversalPaths = NewWeb3HostnameIPFSUniversalPathService(opts...)
	return
}

// Create Web3 Hostname
func (r *Web3HostnameService) New(ctx context.Context, zoneIdentifier string, body Web3HostnameNewParams, opts ...option.RequestOption) (res *DwebConfigWeb3Hostname, err error) {
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

// List Web3 Hostnames
func (r *Web3HostnameService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]DwebConfigWeb3Hostname, err error) {
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

// Edit Web3 Hostname
func (r *Web3HostnameService) Edit(ctx context.Context, zoneIdentifier string, identifier string, body Web3HostnameEditParams, opts ...option.RequestOption) (res *DwebConfigWeb3Hostname, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Web3 Hostname Details
func (r *Web3HostnameService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *DwebConfigWeb3Hostname, err error) {
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

type DwebConfigWeb3Hostname struct {
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
	Status DwebConfigWeb3HostnameStatus `json:"status"`
	// Target gateway of the hostname.
	Target DwebConfigWeb3HostnameTarget `json:"target"`
	JSON   dwebConfigWeb3HostnameJSON   `json:"-"`
}

// dwebConfigWeb3HostnameJSON contains the JSON metadata for the struct
// [DwebConfigWeb3Hostname]
type dwebConfigWeb3HostnameJSON struct {
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

func (r *DwebConfigWeb3Hostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type DwebConfigWeb3HostnameStatus string

const (
	DwebConfigWeb3HostnameStatusActive   DwebConfigWeb3HostnameStatus = "active"
	DwebConfigWeb3HostnameStatusPending  DwebConfigWeb3HostnameStatus = "pending"
	DwebConfigWeb3HostnameStatusDeleting DwebConfigWeb3HostnameStatus = "deleting"
	DwebConfigWeb3HostnameStatusError    DwebConfigWeb3HostnameStatus = "error"
)

// Target gateway of the hostname.
type DwebConfigWeb3HostnameTarget string

const (
	DwebConfigWeb3HostnameTargetEthereum          DwebConfigWeb3HostnameTarget = "ethereum"
	DwebConfigWeb3HostnameTargetIPFS              DwebConfigWeb3HostnameTarget = "ipfs"
	DwebConfigWeb3HostnameTargetIPFSUniversalPath DwebConfigWeb3HostnameTarget = "ipfs_universal_path"
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
	Web3HostnameNewParamsTargetIPFS              Web3HostnameNewParamsTarget = "ipfs"
	Web3HostnameNewParamsTargetIPFSUniversalPath Web3HostnameNewParamsTarget = "ipfs_universal_path"
)

type Web3HostnameNewResponseEnvelope struct {
	Errors   []Web3HostnameNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DwebConfigWeb3Hostname                    `json:"result,required"`
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

type Web3HostnameListResponseEnvelope struct {
	Errors   []Web3HostnameListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DwebConfigWeb3Hostname                   `json:"result,required,nullable"`
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

type Web3HostnameEditParams struct {
	// An optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r Web3HostnameEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Web3HostnameEditResponseEnvelope struct {
	Errors   []Web3HostnameEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameEditResponseEnvelopeMessages `json:"messages,required"`
	Result   DwebConfigWeb3Hostname                     `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameEditResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [Web3HostnameEditResponseEnvelope]
type web3HostnameEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameEditResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    web3HostnameEditResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [Web3HostnameEditResponseEnvelopeErrors]
type web3HostnameEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameEditResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    web3HostnameEditResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [Web3HostnameEditResponseEnvelopeMessages]
type web3HostnameEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameEditResponseEnvelopeSuccess bool

const (
	Web3HostnameEditResponseEnvelopeSuccessTrue Web3HostnameEditResponseEnvelopeSuccess = true
)

type Web3HostnameGetResponseEnvelope struct {
	Errors   []Web3HostnameGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DwebConfigWeb3Hostname                    `json:"result,required"`
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
