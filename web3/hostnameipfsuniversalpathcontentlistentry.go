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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *HostnameIPFSUniversalPathContentListEntryService) New(ctx context.Context, zoneIdentifier string, identifier string, body HostnameIPFSUniversalPathContentListEntryNewParams, opts ...option.RequestOption) (res *ContentListItem, err error) {
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
func (r *HostnameIPFSUniversalPathContentListEntryService) Update(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, body HostnameIPFSUniversalPathContentListEntryUpdateParams, opts ...option.RequestOption) (res *ContentListItem, err error) {
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
func (r *HostnameIPFSUniversalPathContentListEntryService) Delete(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, body HostnameIPFSUniversalPathContentListEntryDeleteParams, opts ...option.RequestOption) (res *HostnameIPFSUniversalPathContentListEntryDeleteResponse, err error) {
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
func (r *HostnameIPFSUniversalPathContentListEntryService) Get(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, opts ...option.RequestOption) (res *ContentListItem, err error) {
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
type ContentListItem struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type ContentListItemType `json:"type"`
	JSON contentListItemJSON `json:"-"`
}

// contentListItemJSON contains the JSON metadata for the struct [ContentListItem]
type contentListItemJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentListItemJSON) RawJSON() string {
	return r.raw
}

// Type of content list entry to block.
type ContentListItemType string

const (
	ContentListItemTypeCid         ContentListItemType = "cid"
	ContentListItemTypeContentPath ContentListItemType = "content_path"
)

func (r ContentListItemType) IsKnown() bool {
	switch r {
	case ContentListItemTypeCid, ContentListItemTypeContentPath:
		return true
	}
	return false
}

// Content list entry to be blocked.
type ContentListItemParam struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
	// Type of content list entry to block.
	Type param.Field[ContentListItemType] `json:"type"`
}

func (r ContentListItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HostnameIPFSUniversalPathContentListEntryListResponse struct {
	// Content list entries.
	Entries []ContentListItem                                         `json:"entries"`
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

func (r HostnameIPFSUniversalPathContentListEntryNewParamsType) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryNewParamsTypeCid, HostnameIPFSUniversalPathContentListEntryNewParamsTypeContentPath:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListEntryNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Content list entry to be blocked.
	Result ContentListItem `json:"result,required"`
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

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r HostnameIPFSUniversalPathContentListEntryUpdateParamsType) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryUpdateParamsTypeCid, HostnameIPFSUniversalPathContentListEntryUpdateParamsTypeContentPath:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Content list entry to be blocked.
	Result ContentListItem `json:"result,required"`
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

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListEntryListResponseEnvelope struct {
	Errors   []shared.ResponseInfo                                 `json:"errors,required"`
	Messages []shared.ResponseInfo                                 `json:"messages,required"`
	Result   HostnameIPFSUniversalPathContentListEntryListResponse `json:"result,required,nullable"`
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

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

type HostnameIPFSUniversalPathContentListEntryDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r HostnameIPFSUniversalPathContentListEntryDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                                   `json:"errors,required"`
	Messages []shared.ResponseInfo                                   `json:"messages,required"`
	Result   HostnameIPFSUniversalPathContentListEntryDeleteResponse `json:"result,required,nullable"`
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

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListEntryGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Content list entry to be blocked.
	Result ContentListItem `json:"result,required"`
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

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
