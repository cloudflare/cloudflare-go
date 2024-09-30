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

// HarService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHarService] method instead.
type HarService struct {
	Options []option.RequestOption
}

// NewHarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHarService(opts ...option.RequestOption) (r *HarService) {
	r = &HarService{}
	r.Options = opts
	return
}

// Get a URL scan's HAR file. See HAR spec at
// http://www.softwareishard.com/blog/har-12-spec/.
func (r *HarService) Get(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *HarGetResponse, err error) {
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

type HarGetResponse struct {
	Log  HarGetResponseLog  `json:"log,required"`
	JSON harGetResponseJSON `json:"-"`
}

// harGetResponseJSON contains the JSON metadata for the struct [HarGetResponse]
type harGetResponseJSON struct {
	Log         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HarGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseJSON) RawJSON() string {
	return r.raw
}

type HarGetResponseLog struct {
	Creator HarGetResponseLogCreator `json:"creator,required"`
	Entries []HarGetResponseLogEntry `json:"entries,required"`
	Pages   []HarGetResponseLogPage  `json:"pages,required"`
	Version string                   `json:"version,required"`
	JSON    harGetResponseLogJSON    `json:"-"`
}

// harGetResponseLogJSON contains the JSON metadata for the struct
// [HarGetResponseLog]
type harGetResponseLogJSON struct {
	Creator     apijson.Field
	Entries     apijson.Field
	Pages       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HarGetResponseLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogJSON) RawJSON() string {
	return r.raw
}

type HarGetResponseLogCreator struct {
	Comment string                       `json:"comment,required"`
	Name    string                       `json:"name,required"`
	Version string                       `json:"version,required"`
	JSON    harGetResponseLogCreatorJSON `json:"-"`
}

// harGetResponseLogCreatorJSON contains the JSON metadata for the struct
// [HarGetResponseLogCreator]
type harGetResponseLogCreatorJSON struct {
	Comment     apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HarGetResponseLogCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogCreatorJSON) RawJSON() string {
	return r.raw
}

type HarGetResponseLogEntry struct {
	InitialPriority string                           `json:"_initialPriority,required"`
	InitiatorType   string                           `json:"_initiator_type,required"`
	Priority        string                           `json:"_priority,required"`
	RequestID       string                           `json:"_requestId,required"`
	RequestTime     float64                          `json:"_requestTime,required"`
	ResourceType    string                           `json:"_resourceType,required"`
	Cache           interface{}                      `json:"cache,required"`
	Connection      string                           `json:"connection,required"`
	Pageref         string                           `json:"pageref,required"`
	Request         HarGetResponseLogEntriesRequest  `json:"request,required"`
	Response        HarGetResponseLogEntriesResponse `json:"response,required"`
	ServerIPAddress string                           `json:"serverIPAddress,required"`
	StartedDateTime string                           `json:"startedDateTime,required"`
	Time            float64                          `json:"time,required"`
	JSON            harGetResponseLogEntryJSON       `json:"-"`
}

// harGetResponseLogEntryJSON contains the JSON metadata for the struct
// [HarGetResponseLogEntry]
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

func (r *HarGetResponseLogEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogEntryJSON) RawJSON() string {
	return r.raw
}

type HarGetResponseLogEntriesRequest struct {
	BodySize    float64                                 `json:"bodySize,required"`
	Headers     []HarGetResponseLogEntriesRequestHeader `json:"headers,required"`
	HeadersSize float64                                 `json:"headersSize,required"`
	HTTPVersion string                                  `json:"httpVersion,required"`
	Method      string                                  `json:"method,required"`
	URL         string                                  `json:"url,required"`
	JSON        harGetResponseLogEntriesRequestJSON     `json:"-"`
}

// harGetResponseLogEntriesRequestJSON contains the JSON metadata for the struct
// [HarGetResponseLogEntriesRequest]
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

func (r *HarGetResponseLogEntriesRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogEntriesRequestJSON) RawJSON() string {
	return r.raw
}

type HarGetResponseLogEntriesRequestHeader struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  harGetResponseLogEntriesRequestHeaderJSON `json:"-"`
}

// harGetResponseLogEntriesRequestHeaderJSON contains the JSON metadata for the
// struct [HarGetResponseLogEntriesRequestHeader]
type harGetResponseLogEntriesRequestHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HarGetResponseLogEntriesRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogEntriesRequestHeaderJSON) RawJSON() string {
	return r.raw
}

type HarGetResponseLogEntriesResponse struct {
	TransferSize float64                                  `json:"_transferSize,required"`
	BodySize     float64                                  `json:"bodySize,required"`
	Content      HarGetResponseLogEntriesResponseContent  `json:"content,required"`
	Headers      []HarGetResponseLogEntriesResponseHeader `json:"headers,required"`
	HeadersSize  float64                                  `json:"headersSize,required"`
	HTTPVersion  string                                   `json:"httpVersion,required"`
	RedirectURL  string                                   `json:"redirectURL,required"`
	Status       float64                                  `json:"status,required"`
	StatusText   string                                   `json:"statusText,required"`
	JSON         harGetResponseLogEntriesResponseJSON     `json:"-"`
}

// harGetResponseLogEntriesResponseJSON contains the JSON metadata for the struct
// [HarGetResponseLogEntriesResponse]
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

func (r *HarGetResponseLogEntriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogEntriesResponseJSON) RawJSON() string {
	return r.raw
}

type HarGetResponseLogEntriesResponseContent struct {
	MimeType    string                                      `json:"mimeType,required"`
	Size        float64                                     `json:"size,required"`
	Compression int64                                       `json:"compression"`
	JSON        harGetResponseLogEntriesResponseContentJSON `json:"-"`
}

// harGetResponseLogEntriesResponseContentJSON contains the JSON metadata for the
// struct [HarGetResponseLogEntriesResponseContent]
type harGetResponseLogEntriesResponseContentJSON struct {
	MimeType    apijson.Field
	Size        apijson.Field
	Compression apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HarGetResponseLogEntriesResponseContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogEntriesResponseContentJSON) RawJSON() string {
	return r.raw
}

type HarGetResponseLogEntriesResponseHeader struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  harGetResponseLogEntriesResponseHeaderJSON `json:"-"`
}

// harGetResponseLogEntriesResponseHeaderJSON contains the JSON metadata for the
// struct [HarGetResponseLogEntriesResponseHeader]
type harGetResponseLogEntriesResponseHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HarGetResponseLogEntriesResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogEntriesResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type HarGetResponseLogPage struct {
	ID              string                            `json:"id,required"`
	PageTimings     HarGetResponseLogPagesPageTimings `json:"pageTimings,required"`
	StartedDateTime string                            `json:"startedDateTime,required"`
	Title           string                            `json:"title,required"`
	JSON            harGetResponseLogPageJSON         `json:"-"`
}

// harGetResponseLogPageJSON contains the JSON metadata for the struct
// [HarGetResponseLogPage]
type harGetResponseLogPageJSON struct {
	ID              apijson.Field
	PageTimings     apijson.Field
	StartedDateTime apijson.Field
	Title           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *HarGetResponseLogPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogPageJSON) RawJSON() string {
	return r.raw
}

type HarGetResponseLogPagesPageTimings struct {
	OnContentLoad float64                               `json:"onContentLoad,required"`
	OnLoad        float64                               `json:"onLoad,required"`
	JSON          harGetResponseLogPagesPageTimingsJSON `json:"-"`
}

// harGetResponseLogPagesPageTimingsJSON contains the JSON metadata for the struct
// [HarGetResponseLogPagesPageTimings]
type harGetResponseLogPagesPageTimingsJSON struct {
	OnContentLoad apijson.Field
	OnLoad        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *HarGetResponseLogPagesPageTimings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r harGetResponseLogPagesPageTimingsJSON) RawJSON() string {
	return r.raw
}
