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

// HostnameIPFSUniversalPathContentListEntryService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHostnameIPFSUniversalPathContentListEntryService] method instead.
type HostnameIPFSUniversalPathContentListEntryService struct {
	Options []option.RequestOption
}

// NewHostnameIPFSUniversalPathContentListEntryService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewHostnameIPFSUniversalPathContentListEntryService(opts ...option.RequestOption) (r *HostnameIPFSUniversalPathContentListEntryService) {
	r = &HostnameIPFSUniversalPathContentListEntryService{}
	r.Options = opts
	return
}

// Create IPFS Universal Path Gateway Content List Entry
func (r *HostnameIPFSUniversalPathContentListEntryService) New(ctx context.Context, zoneIdentifier string, identifier string, body HostnameIPFSUniversalPathContentListEntryNewParams, opts ...option.RequestOption) (res *DwebConfigContentListEntry, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameIPFSUniversalPathContentListEntryNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit IPFS Universal Path Gateway Content List Entry
func (r *HostnameIPFSUniversalPathContentListEntryService) Update(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, body HostnameIPFSUniversalPathContentListEntryUpdateParams, opts ...option.RequestOption) (res *DwebConfigContentListEntry, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List IPFS Universal Path Gateway Content List Entries
func (r *HostnameIPFSUniversalPathContentListEntryService) List(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *HostnameIPFSUniversalPathContentListEntryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameIPFSUniversalPathContentListEntryListResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete IPFS Universal Path Gateway Content List Entry
func (r *HostnameIPFSUniversalPathContentListEntryService) Delete(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, opts ...option.RequestOption) (res *HostnameIPFSUniversalPathContentListEntryDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// IPFS Universal Path Gateway Content List Entry Details
func (r *HostnameIPFSUniversalPathContentListEntryService) Get(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, opts ...option.RequestOption) (res *DwebConfigContentListEntry, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameIPFSUniversalPathContentListEntryGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Content list entry to be blocked.
type DwebConfigContentListEntry struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type DwebConfigContentListEntryType `json:"type"`
	JSON dwebConfigContentListEntryJSON `json:"-"`
}

// dwebConfigContentListEntryJSON contains the JSON metadata for the struct
// [DwebConfigContentListEntry]
type dwebConfigContentListEntryJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DwebConfigContentListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dwebConfigContentListEntryJSON) RawJSON() string {
	return r.raw
}

// Type of content list entry to block.
type DwebConfigContentListEntryType string

const (
	DwebConfigContentListEntryTypeCid         DwebConfigContentListEntryType = "cid"
	DwebConfigContentListEntryTypeContentPath DwebConfigContentListEntryType = "content_path"
)

// Content list entry to be blocked.
type DwebConfigContentListEntryParam struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
	// Type of content list entry to block.
	Type param.Field[DwebConfigContentListEntryType] `json:"type"`
}

func (r DwebConfigContentListEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HostnameIPFSUniversalPathContentListEntryListResponse struct {
	// Content list entries.
	Entries []DwebConfigContentListEntry                              `json:"entries"`
	JSON    hostnameIPFSUniversalPathContentListEntryListResponseJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryListResponseJSON contains the JSON
// metadata for the struct [HostnameIPFSUniversalPathContentListEntryListResponse]
type hostnameIPFSUniversalPathContentListEntryListResponseJSON struct {
	Entries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryListResponseJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryDeleteResponse struct {
	// Identifier
	ID   string                                                      `json:"id,required"`
	JSON hostnameIPFSUniversalPathContentListEntryDeleteResponseJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryDeleteResponseJSON contains the JSON
// metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryDeleteResponse]
type hostnameIPFSUniversalPathContentListEntryDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryNewParams struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content,required"`
	// Type of content list entry to block.
	Type param.Field[HostnameIPFSUniversalPathContentListEntryNewParamsType] `json:"type,required"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
}

func (r HostnameIPFSUniversalPathContentListEntryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of content list entry to block.
type HostnameIPFSUniversalPathContentListEntryNewParamsType string

const (
	HostnameIPFSUniversalPathContentListEntryNewParamsTypeCid         HostnameIPFSUniversalPathContentListEntryNewParamsType = "cid"
	HostnameIPFSUniversalPathContentListEntryNewParamsTypeContentPath HostnameIPFSUniversalPathContentListEntryNewParamsType = "content_path"
)

type HostnameIPFSUniversalPathContentListEntryNewResponseEnvelope struct {
	Errors   []HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessages `json:"messages,required"`
	// Content list entry to be blocked.
	Result DwebConfigContentListEntry `json:"result,required"`
	// Whether the API call was successful
	Success HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeJSON    `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryNewResponseEnvelope]
type hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrors]
type hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessages]
type hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccess = true
)

type HostnameIPFSUniversalPathContentListEntryUpdateParams struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content,required"`
	// Type of content list entry to block.
	Type param.Field[HostnameIPFSUniversalPathContentListEntryUpdateParamsType] `json:"type,required"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
}

func (r HostnameIPFSUniversalPathContentListEntryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of content list entry to block.
type HostnameIPFSUniversalPathContentListEntryUpdateParamsType string

const (
	HostnameIPFSUniversalPathContentListEntryUpdateParamsTypeCid         HostnameIPFSUniversalPathContentListEntryUpdateParamsType = "cid"
	HostnameIPFSUniversalPathContentListEntryUpdateParamsTypeContentPath HostnameIPFSUniversalPathContentListEntryUpdateParamsType = "content_path"
)

type HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelope struct {
	Errors   []HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Content list entry to be blocked.
	Result DwebConfigContentListEntry `json:"result,required"`
	// Whether the API call was successful
	Success HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeJSON    `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelope]
type hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrors]
type hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessages]
type hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccess = true
)

type HostnameIPFSUniversalPathContentListEntryListResponseEnvelope struct {
	Errors   []HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessages `json:"messages,required"`
	Result   HostnameIPFSUniversalPathContentListEntryListResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeJSON       `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryListResponseEnvelope]
type hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrors struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrors]
type hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessages struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessages]
type hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccess = true
)

type HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                     `json:"total_count"`
	JSON       hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfo]
type hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelope struct {
	Errors   []HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   HostnameIPFSUniversalPathContentListEntryDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeJSON    `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelope]
type hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrors]
type hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessages]
type hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccess = true
)

type HostnameIPFSUniversalPathContentListEntryGetResponseEnvelope struct {
	Errors   []HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessages `json:"messages,required"`
	// Content list entry to be blocked.
	Result DwebConfigContentListEntry `json:"result,required"`
	// Whether the API call was successful
	Success HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeJSON    `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryGetResponseEnvelope]
type hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrors]
type hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessages]
type hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccess = true
)
