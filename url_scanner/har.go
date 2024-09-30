// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_scanner

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// HARService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHARService] method instead.
type HARService struct {
	Options []option.RequestOption
}

// NewHARService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHARService(opts ...option.RequestOption) (r *HARService) {
	r = &HARService{}
	r.Options = opts
	return
}

// Get a URL scan's HAR file. See HAR spec at
// http://www.softwareishard.com/blog/har-12-spec/.
func (r *HARService) Get(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *HARGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scanId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/har/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type HARGetResponse struct {
	Log  HARGetResponseLog  `json:"log,required"`
	JSON harGetResponseJSON `json:"-"`
}

// harGetResponseJSON contains the JSON metadata for the struct [HARGetResponse]
type harGetResponseJSON struct {
	Log         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HARGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseJSON) RawJSON() string {
	return r.raw
}

type HARGetResponseLog struct {
	Creator HARGetResponseLogCreator `json:"creator,required"`
	Entries []HARGetResponseLogEntry `json:"entries,required"`
	Pages   []HARGetResponseLogPage  `json:"pages,required"`
	Version string                   `json:"version,required"`
	JSON    harGetResponseLogJSON    `json:"-"`
}

// harGetResponseLogJSON contains the JSON metadata for the struct
// [HARGetResponseLog]
type harGetResponseLogJSON struct {
	Creator     apijson.Field
	Entries     apijson.Field
	Pages       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HARGetResponseLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogJSON) RawJSON() string {
	return r.raw
}

type HARGetResponseLogCreator struct {
	Comment string                       `json:"comment,required"`
	Name    string                       `json:"name,required"`
	Version string                       `json:"version,required"`
	JSON    harGetResponseLogCreatorJSON `json:"-"`
}

// harGetResponseLogCreatorJSON contains the JSON metadata for the struct
// [HARGetResponseLogCreator]
type harGetResponseLogCreatorJSON struct {
	Comment     apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HARGetResponseLogCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogCreatorJSON) RawJSON() string {
	return r.raw
}

type HARGetResponseLogEntry struct {
	InitialPriority string                           `json:"_initialPriority,required"`
	InitiatorType   string                           `json:"_initiator_type,required"`
	Priority        string                           `json:"_priority,required"`
	RequestID       string                           `json:"_requestId,required"`
	RequestTime     float64                          `json:"_requestTime,required"`
	ResourceType    string                           `json:"_resourceType,required"`
	Cache           interface{}                      `json:"cache,required"`
	Connection      string                           `json:"connection,required"`
	Pageref         string                           `json:"pageref,required"`
	Request         HARGetResponseLogEntriesRequest  `json:"request,required"`
	Response        HARGetResponseLogEntriesResponse `json:"response,required"`
	ServerIPAddress string                           `json:"serverIPAddress,required"`
	StartedDateTime string                           `json:"startedDateTime,required"`
	Time            float64                          `json:"time,required"`
	JSON            harGetResponseLogEntryJSON       `json:"-"`
}

// harGetResponseLogEntryJSON contains the JSON metadata for the struct
// [HARGetResponseLogEntry]
type harGetResponseLogEntryJSON struct {
	InitialPriority apijson.Field
	InitiatorType   apijson.Field
	Priority        apijson.Field
	RequestID       apijson.Field
	RequestTime     apijson.Field
	ResourceType    apijson.Field
	Cache           apijson.Field
	Connection      apijson.Field
	Pageref         apijson.Field
	Request         apijson.Field
	Response        apijson.Field
	ServerIPAddress apijson.Field
	StartedDateTime apijson.Field
	Time            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *HARGetResponseLogEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogEntryJSON) RawJSON() string {
	return r.raw
}

type HARGetResponseLogEntriesRequest struct {
	BodySize    float64                                 `json:"bodySize,required"`
	Headers     []HARGetResponseLogEntriesRequestHeader `json:"headers,required"`
	HeadersSize float64                                 `json:"headersSize,required"`
	HTTPVersion string                                  `json:"httpVersion,required"`
	Method      string                                  `json:"method,required"`
	URL         string                                  `json:"url,required"`
	JSON        harGetResponseLogEntriesRequestJSON     `json:"-"`
}

// harGetResponseLogEntriesRequestJSON contains the JSON metadata for the struct
// [HARGetResponseLogEntriesRequest]
type harGetResponseLogEntriesRequestJSON struct {
	BodySize    apijson.Field
	Headers     apijson.Field
	HeadersSize apijson.Field
	HTTPVersion apijson.Field
	Method      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HARGetResponseLogEntriesRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogEntriesRequestJSON) RawJSON() string {
	return r.raw
}

type HARGetResponseLogEntriesRequestHeader struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  harGetResponseLogEntriesRequestHeaderJSON `json:"-"`
}

// harGetResponseLogEntriesRequestHeaderJSON contains the JSON metadata for the
// struct [HARGetResponseLogEntriesRequestHeader]
type harGetResponseLogEntriesRequestHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HARGetResponseLogEntriesRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogEntriesRequestHeaderJSON) RawJSON() string {
	return r.raw
}

type HARGetResponseLogEntriesResponse struct {
	TransferSize float64                                  `json:"_transferSize,required"`
	BodySize     float64                                  `json:"bodySize,required"`
	Content      HARGetResponseLogEntriesResponseContent  `json:"content,required"`
	Headers      []HARGetResponseLogEntriesResponseHeader `json:"headers,required"`
	HeadersSize  float64                                  `json:"headersSize,required"`
	HTTPVersion  string                                   `json:"httpVersion,required"`
	RedirectURL  string                                   `json:"redirectURL,required"`
	Status       float64                                  `json:"status,required"`
	StatusText   string                                   `json:"statusText,required"`
	JSON         harGetResponseLogEntriesResponseJSON     `json:"-"`
}

// harGetResponseLogEntriesResponseJSON contains the JSON metadata for the struct
// [HARGetResponseLogEntriesResponse]
type harGetResponseLogEntriesResponseJSON struct {
	TransferSize apijson.Field
	BodySize     apijson.Field
	Content      apijson.Field
	Headers      apijson.Field
	HeadersSize  apijson.Field
	HTTPVersion  apijson.Field
	RedirectURL  apijson.Field
	Status       apijson.Field
	StatusText   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HARGetResponseLogEntriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogEntriesResponseJSON) RawJSON() string {
	return r.raw
}

type HARGetResponseLogEntriesResponseContent struct {
	MimeType    string                                      `json:"mimeType,required"`
	Size        float64                                     `json:"size,required"`
	Compression int64                                       `json:"compression"`
	JSON        harGetResponseLogEntriesResponseContentJSON `json:"-"`
}

// harGetResponseLogEntriesResponseContentJSON contains the JSON metadata for the
// struct [HARGetResponseLogEntriesResponseContent]
type harGetResponseLogEntriesResponseContentJSON struct {
	MimeType    apijson.Field
	Size        apijson.Field
	Compression apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HARGetResponseLogEntriesResponseContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogEntriesResponseContentJSON) RawJSON() string {
	return r.raw
}

type HARGetResponseLogEntriesResponseHeader struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  harGetResponseLogEntriesResponseHeaderJSON `json:"-"`
}

// harGetResponseLogEntriesResponseHeaderJSON contains the JSON metadata for the
// struct [HARGetResponseLogEntriesResponseHeader]
type harGetResponseLogEntriesResponseHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HARGetResponseLogEntriesResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogEntriesResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type HARGetResponseLogPage struct {
	ID              string                            `json:"id,required"`
	PageTimings     HARGetResponseLogPagesPageTimings `json:"pageTimings,required"`
	StartedDateTime string                            `json:"startedDateTime,required"`
	Title           string                            `json:"title,required"`
	JSON            harGetResponseLogPageJSON         `json:"-"`
}

// harGetResponseLogPageJSON contains the JSON metadata for the struct
// [HARGetResponseLogPage]
type harGetResponseLogPageJSON struct {
	ID              apijson.Field
	PageTimings     apijson.Field
	StartedDateTime apijson.Field
	Title           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *HARGetResponseLogPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogPageJSON) RawJSON() string {
	return r.raw
}

type HARGetResponseLogPagesPageTimings struct {
	OnContentLoad float64                               `json:"onContentLoad,required"`
	OnLoad        float64                               `json:"onLoad,required"`
	JSON          harGetResponseLogPagesPageTimingsJSON `json:"-"`
}

// harGetResponseLogPagesPageTimingsJSON contains the JSON metadata for the struct
// [HARGetResponseLogPagesPageTimings]
type harGetResponseLogPagesPageTimingsJSON struct {
	OnContentLoad apijson.Field
	OnLoad        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *HARGetResponseLogPagesPageTimings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogPagesPageTimingsJSON) RawJSON() string {
	return r.raw
}
