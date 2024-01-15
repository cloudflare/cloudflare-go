// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountUrlscannerHarService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountUrlscannerHarService]
// method instead.
type AccountUrlscannerHarService struct {
	Options []option.RequestOption
}

// NewAccountUrlscannerHarService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountUrlscannerHarService(opts ...option.RequestOption) (r *AccountUrlscannerHarService) {
	r = &AccountUrlscannerHarService{}
	r.Options = opts
	return
}

// Get a URL scan's HAR file. See HAR spec at
// http://www.softwareishard.com/blog/har-12-spec/.
func (r *AccountUrlscannerHarService) Get(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *AccountUrlscannerHarGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s/har", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountUrlscannerHarGetResponse struct {
	Errors   []AccountUrlscannerHarGetResponseError   `json:"errors,required"`
	Messages []AccountUrlscannerHarGetResponseMessage `json:"messages,required"`
	Result   AccountUrlscannerHarGetResponseResult    `json:"result,required"`
	// Whether search request was successful or not
	Success bool                                `json:"success,required"`
	JSON    accountUrlscannerHarGetResponseJSON `json:"-"`
}

// accountUrlscannerHarGetResponseJSON contains the JSON metadata for the struct
// [AccountUrlscannerHarGetResponse]
type accountUrlscannerHarGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseError struct {
	Message string                                   `json:"message,required"`
	JSON    accountUrlscannerHarGetResponseErrorJSON `json:"-"`
}

// accountUrlscannerHarGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountUrlscannerHarGetResponseError]
type accountUrlscannerHarGetResponseErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseMessage struct {
	Message string                                     `json:"message,required"`
	JSON    accountUrlscannerHarGetResponseMessageJSON `json:"-"`
}

// accountUrlscannerHarGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountUrlscannerHarGetResponseMessage]
type accountUrlscannerHarGetResponseMessageJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseResult struct {
	Har  AccountUrlscannerHarGetResponseResultHar  `json:"har,required"`
	JSON accountUrlscannerHarGetResponseResultJSON `json:"-"`
}

// accountUrlscannerHarGetResponseResultJSON contains the JSON metadata for the
// struct [AccountUrlscannerHarGetResponseResult]
type accountUrlscannerHarGetResponseResultJSON struct {
	Har         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseResultHar struct {
	Log  AccountUrlscannerHarGetResponseResultHarLog  `json:"log,required"`
	JSON accountUrlscannerHarGetResponseResultHarJSON `json:"-"`
}

// accountUrlscannerHarGetResponseResultHarJSON contains the JSON metadata for the
// struct [AccountUrlscannerHarGetResponseResultHar]
type accountUrlscannerHarGetResponseResultHarJSON struct {
	Log         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseResultHar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseResultHarLog struct {
	Creator AccountUrlscannerHarGetResponseResultHarLogCreator `json:"creator,required"`
	Entries []AccountUrlscannerHarGetResponseResultHarLogEntry `json:"entries,required"`
	Pages   []AccountUrlscannerHarGetResponseResultHarLogPage  `json:"pages,required"`
	Version string                                             `json:"version,required"`
	JSON    accountUrlscannerHarGetResponseResultHarLogJSON    `json:"-"`
}

// accountUrlscannerHarGetResponseResultHarLogJSON contains the JSON metadata for
// the struct [AccountUrlscannerHarGetResponseResultHarLog]
type accountUrlscannerHarGetResponseResultHarLogJSON struct {
	Creator     apijson.Field
	Entries     apijson.Field
	Pages       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseResultHarLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseResultHarLogCreator struct {
	Comment string                                                 `json:"comment,required"`
	Name    string                                                 `json:"name,required"`
	Version string                                                 `json:"version,required"`
	JSON    accountUrlscannerHarGetResponseResultHarLogCreatorJSON `json:"-"`
}

// accountUrlscannerHarGetResponseResultHarLogCreatorJSON contains the JSON
// metadata for the struct [AccountUrlscannerHarGetResponseResultHarLogCreator]
type accountUrlscannerHarGetResponseResultHarLogCreatorJSON struct {
	Comment     apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseResultHarLogCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseResultHarLogEntry struct {
	_InitialPriority string                                                     `json:"_initialPriority,required"`
	_InitiatorType   string                                                     `json:"_initiator_type,required"`
	_Priority        string                                                     `json:"_priority,required"`
	_RequestID       string                                                     `json:"_requestId,required"`
	_RequestTime     float64                                                    `json:"_requestTime,required"`
	_ResourceType    string                                                     `json:"_resourceType,required"`
	Cache            interface{}                                                `json:"cache,required"`
	Connection       string                                                     `json:"connection,required"`
	Pageref          string                                                     `json:"pageref,required"`
	Request          AccountUrlscannerHarGetResponseResultHarLogEntriesRequest  `json:"request,required"`
	Response         AccountUrlscannerHarGetResponseResultHarLogEntriesResponse `json:"response,required"`
	ServerIPAddress  string                                                     `json:"serverIPAddress,required"`
	StartedDateTime  string                                                     `json:"startedDateTime,required"`
	Time             float64                                                    `json:"time,required"`
	JSON             accountUrlscannerHarGetResponseResultHarLogEntryJSON       `json:"-"`
}

// accountUrlscannerHarGetResponseResultHarLogEntryJSON contains the JSON metadata
// for the struct [AccountUrlscannerHarGetResponseResultHarLogEntry]
type accountUrlscannerHarGetResponseResultHarLogEntryJSON struct {
	_InitialPriority apijson.Field
	_InitiatorType   apijson.Field
	_Priority        apijson.Field
	_RequestID       apijson.Field
	_RequestTime     apijson.Field
	_ResourceType    apijson.Field
	Cache            apijson.Field
	Connection       apijson.Field
	Pageref          apijson.Field
	Request          apijson.Field
	Response         apijson.Field
	ServerIPAddress  apijson.Field
	StartedDateTime  apijson.Field
	Time             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseResultHarLogEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseResultHarLogEntriesRequest struct {
	BodySize    float64                                                           `json:"bodySize,required"`
	Headers     []AccountUrlscannerHarGetResponseResultHarLogEntriesRequestHeader `json:"headers,required"`
	HeadersSize float64                                                           `json:"headersSize,required"`
	HTTPVersion string                                                            `json:"httpVersion,required"`
	Method      string                                                            `json:"method,required"`
	URL         string                                                            `json:"url,required"`
	JSON        accountUrlscannerHarGetResponseResultHarLogEntriesRequestJSON     `json:"-"`
}

// accountUrlscannerHarGetResponseResultHarLogEntriesRequestJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerHarGetResponseResultHarLogEntriesRequest]
type accountUrlscannerHarGetResponseResultHarLogEntriesRequestJSON struct {
	BodySize    apijson.Field
	Headers     apijson.Field
	HeadersSize apijson.Field
	HTTPVersion apijson.Field
	Method      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseResultHarLogEntriesRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseResultHarLogEntriesRequestHeader struct {
	Name  string                                                              `json:"name,required"`
	Value string                                                              `json:"value,required"`
	JSON  accountUrlscannerHarGetResponseResultHarLogEntriesRequestHeaderJSON `json:"-"`
}

// accountUrlscannerHarGetResponseResultHarLogEntriesRequestHeaderJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerHarGetResponseResultHarLogEntriesRequestHeader]
type accountUrlscannerHarGetResponseResultHarLogEntriesRequestHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseResultHarLogEntriesRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseResultHarLogEntriesResponse struct {
	_TransferSize float64                                                            `json:"_transferSize,required"`
	BodySize      float64                                                            `json:"bodySize,required"`
	Content       AccountUrlscannerHarGetResponseResultHarLogEntriesResponseContent  `json:"content,required"`
	Headers       []AccountUrlscannerHarGetResponseResultHarLogEntriesResponseHeader `json:"headers,required"`
	HeadersSize   float64                                                            `json:"headersSize,required"`
	HTTPVersion   string                                                             `json:"httpVersion,required"`
	RedirectURL   string                                                             `json:"redirectURL,required"`
	Status        float64                                                            `json:"status,required"`
	StatusText    string                                                             `json:"statusText,required"`
	JSON          accountUrlscannerHarGetResponseResultHarLogEntriesResponseJSON     `json:"-"`
}

// accountUrlscannerHarGetResponseResultHarLogEntriesResponseJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerHarGetResponseResultHarLogEntriesResponse]
type accountUrlscannerHarGetResponseResultHarLogEntriesResponseJSON struct {
	_TransferSize apijson.Field
	BodySize      apijson.Field
	Content       apijson.Field
	Headers       apijson.Field
	HeadersSize   apijson.Field
	HTTPVersion   apijson.Field
	RedirectURL   apijson.Field
	Status        apijson.Field
	StatusText    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseResultHarLogEntriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseResultHarLogEntriesResponseContent struct {
	MimeType    string                                                                `json:"mimeType,required"`
	Size        float64                                                               `json:"size,required"`
	Compression int64                                                                 `json:"compression"`
	JSON        accountUrlscannerHarGetResponseResultHarLogEntriesResponseContentJSON `json:"-"`
}

// accountUrlscannerHarGetResponseResultHarLogEntriesResponseContentJSON contains
// the JSON metadata for the struct
// [AccountUrlscannerHarGetResponseResultHarLogEntriesResponseContent]
type accountUrlscannerHarGetResponseResultHarLogEntriesResponseContentJSON struct {
	MimeType    apijson.Field
	Size        apijson.Field
	Compression apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseResultHarLogEntriesResponseContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseResultHarLogEntriesResponseHeader struct {
	Name  string                                                               `json:"name,required"`
	Value string                                                               `json:"value,required"`
	JSON  accountUrlscannerHarGetResponseResultHarLogEntriesResponseHeaderJSON `json:"-"`
}

// accountUrlscannerHarGetResponseResultHarLogEntriesResponseHeaderJSON contains
// the JSON metadata for the struct
// [AccountUrlscannerHarGetResponseResultHarLogEntriesResponseHeader]
type accountUrlscannerHarGetResponseResultHarLogEntriesResponseHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseResultHarLogEntriesResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseResultHarLogPage struct {
	ID              string                                                      `json:"id,required"`
	PageTimings     AccountUrlscannerHarGetResponseResultHarLogPagesPageTimings `json:"pageTimings,required"`
	StartedDateTime string                                                      `json:"startedDateTime,required"`
	Title           string                                                      `json:"title,required"`
	JSON            accountUrlscannerHarGetResponseResultHarLogPageJSON         `json:"-"`
}

// accountUrlscannerHarGetResponseResultHarLogPageJSON contains the JSON metadata
// for the struct [AccountUrlscannerHarGetResponseResultHarLogPage]
type accountUrlscannerHarGetResponseResultHarLogPageJSON struct {
	ID              apijson.Field
	PageTimings     apijson.Field
	StartedDateTime apijson.Field
	Title           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseResultHarLogPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUrlscannerHarGetResponseResultHarLogPagesPageTimings struct {
	OnContentLoad float64                                                         `json:"onContentLoad,required"`
	OnLoad        float64                                                         `json:"onLoad,required"`
	JSON          accountUrlscannerHarGetResponseResultHarLogPagesPageTimingsJSON `json:"-"`
}

// accountUrlscannerHarGetResponseResultHarLogPagesPageTimingsJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerHarGetResponseResultHarLogPagesPageTimings]
type accountUrlscannerHarGetResponseResultHarLogPagesPageTimingsJSON struct {
	OnContentLoad apijson.Field
	OnLoad        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountUrlscannerHarGetResponseResultHarLogPagesPageTimings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
