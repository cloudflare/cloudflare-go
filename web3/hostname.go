// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// HostnameService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewHostnameService] method instead.
type HostnameService struct {
	Options            []option.RequestOption
	IPFSUniversalPaths *HostnameIPFSUniversalPathService
}

// NewHostnameService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHostnameService(opts ...option.RequestOption) (r *HostnameService) {
	r = &HostnameService{}
	r.Options = opts
	r.IPFSUniversalPaths = NewHostnameIPFSUniversalPathService(opts...)
	return
}

// Create Web3 Hostname
func (r *HostnameService) New(ctx context.Context, zoneIdentifier string, body HostnameNewParams, opts ...option.RequestOption) (res *DwebConfigWeb3Hostname, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Web3 Hostnames
func (r *HostnameService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]DwebConfigWeb3Hostname, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameListResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Web3 Hostname
func (r *HostnameService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *HostnameDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit Web3 Hostname
func (r *HostnameService) Edit(ctx context.Context, zoneIdentifier string, identifier string, body HostnameEditParams, opts ...option.RequestOption) (res *DwebConfigWeb3Hostname, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Web3 Hostname Details
func (r *HostnameService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *DwebConfigWeb3Hostname, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameGetResponseEnvelope
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

func (r dwebConfigWeb3HostnameJSON) RawJSON() string {
	return r.raw
}

// Status of the hostname's activation.
type DwebConfigWeb3HostnameStatus string

const (
	DwebConfigWeb3HostnameStatusActive   DwebConfigWeb3HostnameStatus = "active"
	DwebConfigWeb3HostnameStatusPending  DwebConfigWeb3HostnameStatus = "pending"
	DwebConfigWeb3HostnameStatusDeleting DwebConfigWeb3HostnameStatus = "deleting"
	DwebConfigWeb3HostnameStatusError    DwebConfigWeb3HostnameStatus = "error"
)

func (r DwebConfigWeb3HostnameStatus) IsKnown() bool {
	switch r {
	case DwebConfigWeb3HostnameStatusActive, DwebConfigWeb3HostnameStatusPending, DwebConfigWeb3HostnameStatusDeleting, DwebConfigWeb3HostnameStatusError:
		return true
	}
	return false
}

// Target gateway of the hostname.
type DwebConfigWeb3HostnameTarget string

const (
	DwebConfigWeb3HostnameTargetEthereum          DwebConfigWeb3HostnameTarget = "ethereum"
	DwebConfigWeb3HostnameTargetIPFS              DwebConfigWeb3HostnameTarget = "ipfs"
	DwebConfigWeb3HostnameTargetIPFSUniversalPath DwebConfigWeb3HostnameTarget = "ipfs_universal_path"
)

func (r DwebConfigWeb3HostnameTarget) IsKnown() bool {
	switch r {
	case DwebConfigWeb3HostnameTargetEthereum, DwebConfigWeb3HostnameTargetIPFS, DwebConfigWeb3HostnameTargetIPFSUniversalPath:
		return true
	}
	return false
}

type HostnameDeleteResponse struct {
	// Identifier
	ID   string                     `json:"id,required"`
	JSON hostnameDeleteResponseJSON `json:"-"`
}

// hostnameDeleteResponseJSON contains the JSON metadata for the struct
// [HostnameDeleteResponse]
type hostnameDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type HostnameNewParams struct {
	// Target gateway of the hostname.
	Target param.Field[HostnameNewParamsTarget] `json:"target,required"`
	// An optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r HostnameNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Target gateway of the hostname.
type HostnameNewParamsTarget string

const (
	HostnameNewParamsTargetEthereum          HostnameNewParamsTarget = "ethereum"
	HostnameNewParamsTargetIPFS              HostnameNewParamsTarget = "ipfs"
	HostnameNewParamsTargetIPFSUniversalPath HostnameNewParamsTarget = "ipfs_universal_path"
)

func (r HostnameNewParamsTarget) IsKnown() bool {
	switch r {
	case HostnameNewParamsTargetEthereum, HostnameNewParamsTargetIPFS, HostnameNewParamsTargetIPFSUniversalPath:
		return true
	}
	return false
}

type HostnameNewResponseEnvelope struct {
	Errors   []HostnameNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DwebConfigWeb3Hostname                `json:"result,required"`
	// Whether the API call was successful
	Success HostnameNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameNewResponseEnvelopeJSON    `json:"-"`
}

// hostnameNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [HostnameNewResponseEnvelope]
type hostnameNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameNewResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    hostnameNewResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [HostnameNewResponseEnvelopeErrors]
type hostnameNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameNewResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    hostnameNewResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HostnameNewResponseEnvelopeMessages]
type hostnameNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameNewResponseEnvelopeSuccess bool

const (
	HostnameNewResponseEnvelopeSuccessTrue HostnameNewResponseEnvelopeSuccess = true
)

func (r HostnameNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameListResponseEnvelope struct {
	Errors   []HostnameListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DwebConfigWeb3Hostname               `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    HostnameListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo HostnameListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       hostnameListResponseEnvelopeJSON       `json:"-"`
}

// hostnameListResponseEnvelopeJSON contains the JSON metadata for the struct
// [HostnameListResponseEnvelope]
type hostnameListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameListResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    hostnameListResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [HostnameListResponseEnvelopeErrors]
type hostnameListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameListResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    hostnameListResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HostnameListResponseEnvelopeMessages]
type hostnameListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameListResponseEnvelopeSuccess bool

const (
	HostnameListResponseEnvelopeSuccessTrue HostnameListResponseEnvelopeSuccess = true
)

func (r HostnameListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       hostnameListResponseEnvelopeResultInfoJSON `json:"-"`
}

// hostnameListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [HostnameListResponseEnvelopeResultInfo]
type hostnameListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type HostnameDeleteResponseEnvelope struct {
	Errors   []HostnameDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   HostnameDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success HostnameDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameDeleteResponseEnvelopeJSON    `json:"-"`
}

// hostnameDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [HostnameDeleteResponseEnvelope]
type hostnameDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameDeleteResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    hostnameDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HostnameDeleteResponseEnvelopeErrors]
type hostnameDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameDeleteResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    hostnameDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HostnameDeleteResponseEnvelopeMessages]
type hostnameDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameDeleteResponseEnvelopeSuccess bool

const (
	HostnameDeleteResponseEnvelopeSuccessTrue HostnameDeleteResponseEnvelopeSuccess = true
)

func (r HostnameDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameEditParams struct {
	// An optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r HostnameEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HostnameEditResponseEnvelope struct {
	Errors   []HostnameEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameEditResponseEnvelopeMessages `json:"messages,required"`
	Result   DwebConfigWeb3Hostname                 `json:"result,required"`
	// Whether the API call was successful
	Success HostnameEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameEditResponseEnvelopeJSON    `json:"-"`
}

// hostnameEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [HostnameEditResponseEnvelope]
type hostnameEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameEditResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    hostnameEditResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [HostnameEditResponseEnvelopeErrors]
type hostnameEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameEditResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    hostnameEditResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HostnameEditResponseEnvelopeMessages]
type hostnameEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameEditResponseEnvelopeSuccess bool

const (
	HostnameEditResponseEnvelopeSuccessTrue HostnameEditResponseEnvelopeSuccess = true
)

func (r HostnameEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameGetResponseEnvelope struct {
	Errors   []HostnameGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DwebConfigWeb3Hostname                `json:"result,required"`
	// Whether the API call was successful
	Success HostnameGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameGetResponseEnvelopeJSON    `json:"-"`
}

// hostnameGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HostnameGetResponseEnvelope]
type hostnameGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    hostnameGetResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [HostnameGetResponseEnvelopeErrors]
type hostnameGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameGetResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    hostnameGetResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HostnameGetResponseEnvelopeMessages]
type hostnameGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameGetResponseEnvelopeSuccess bool

const (
	HostnameGetResponseEnvelopeSuccessTrue HostnameGetResponseEnvelopeSuccess = true
)

func (r HostnameGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
