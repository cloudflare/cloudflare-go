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

// ResultService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResultService] method instead.
type ResultService struct {
	Options []option.RequestOption
}

// NewResultService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewResultService(opts ...option.RequestOption) (r *ResultService) {
	r = &ResultService{}
	r.Options = opts
	return
}

// Get URL scan by uuid
func (r *ResultService) Get(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *ResultGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scanId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/result/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ResultGetResponse struct {
	Data     ResultGetResponseData     `json:"data,required"`
	Lists    ResultGetResponseLists    `json:"lists,required"`
	Meta     ResultGetResponseMeta     `json:"meta,required"`
	Page     ResultGetResponsePage     `json:"page,required"`
	Scanner  ResultGetResponseScanner  `json:"scanner,required"`
	Stats    ResultGetResponseStats    `json:"stats,required"`
	Task     ResultGetResponseTask     `json:"task,required"`
	Verdicts ResultGetResponseVerdicts `json:"verdicts,required"`
	JSON     resultGetResponseJSON     `json:"-"`
}

// resultGetResponseJSON contains the JSON metadata for the struct
// [ResultGetResponse]
type resultGetResponseJSON struct {
	Data        apijson.Field
	Lists       apijson.Field
	Meta        apijson.Field
	Page        apijson.Field
	Scanner     apijson.Field
	Stats       apijson.Field
	Task        apijson.Field
	Verdicts    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseData struct {
	Console     []ResultGetResponseDataConsole     `json:"console,required"`
	Cookies     []ResultGetResponseDataCookie      `json:"cookies,required"`
	Globals     []ResultGetResponseDataGlobal      `json:"globals,required"`
	Links       []ResultGetResponseDataLink        `json:"links,required"`
	Performance []ResultGetResponseDataPerformance `json:"performance,required"`
	Requests    []ResultGetResponseDataRequest     `json:"requests,required"`
	JSON        resultGetResponseDataJSON          `json:"-"`
}

// resultGetResponseDataJSON contains the JSON metadata for the struct
// [ResultGetResponseData]
type resultGetResponseDataJSON struct {
	Console     apijson.Field
	Cookies     apijson.Field
	Globals     apijson.Field
	Links       apijson.Field
	Performance apijson.Field
	Requests    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataConsole struct {
	Message ResultGetResponseDataConsoleMessage `json:"message,required"`
	JSON    resultGetResponseDataConsoleJSON    `json:"-"`
}

// resultGetResponseDataConsoleJSON contains the JSON metadata for the struct
// [ResultGetResponseDataConsole]
type resultGetResponseDataConsoleJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseDataConsole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataConsoleJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataConsoleMessage struct {
	Level  string                                  `json:"level,required"`
	Source string                                  `json:"source,required"`
	Text   string                                  `json:"text,required"`
	URL    string                                  `json:"url,required"`
	JSON   resultGetResponseDataConsoleMessageJSON `json:"-"`
}

// resultGetResponseDataConsoleMessageJSON contains the JSON metadata for the
// struct [ResultGetResponseDataConsoleMessage]
type resultGetResponseDataConsoleMessageJSON struct {
	Level       apijson.Field
	Source      apijson.Field
	Text        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseDataConsoleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataConsoleMessageJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataCookie struct {
	Domain       string                          `json:"domain,required"`
	Expires      float64                         `json:"expires,required"`
	HTTPOnly     bool                            `json:"httpOnly,required"`
	Name         string                          `json:"name,required"`
	Path         string                          `json:"path,required"`
	Priority     string                          `json:"priority,required"`
	SameParty    bool                            `json:"sameParty,required"`
	Secure       bool                            `json:"secure,required"`
	Session      bool                            `json:"session,required"`
	Size         float64                         `json:"size,required"`
	SourcePort   float64                         `json:"sourcePort,required"`
	SourceScheme string                          `json:"sourceScheme,required"`
	Value        string                          `json:"value,required"`
	JSON         resultGetResponseDataCookieJSON `json:"-"`
}

// resultGetResponseDataCookieJSON contains the JSON metadata for the struct
// [ResultGetResponseDataCookie]
type resultGetResponseDataCookieJSON struct {
	Domain       apijson.Field
	Expires      apijson.Field
	HTTPOnly     apijson.Field
	Name         apijson.Field
	Path         apijson.Field
	Priority     apijson.Field
	SameParty    apijson.Field
	Secure       apijson.Field
	Session      apijson.Field
	Size         apijson.Field
	SourcePort   apijson.Field
	SourceScheme apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResultGetResponseDataCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataCookieJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataGlobal struct {
	Prop string                          `json:"prop,required"`
	Type string                          `json:"type,required"`
	JSON resultGetResponseDataGlobalJSON `json:"-"`
}

// resultGetResponseDataGlobalJSON contains the JSON metadata for the struct
// [ResultGetResponseDataGlobal]
type resultGetResponseDataGlobalJSON struct {
	Prop        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseDataGlobal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataGlobalJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataLink struct {
	Href string                        `json:"href,required"`
	Text string                        `json:"text,required"`
	JSON resultGetResponseDataLinkJSON `json:"-"`
}

// resultGetResponseDataLinkJSON contains the JSON metadata for the struct
// [ResultGetResponseDataLink]
type resultGetResponseDataLinkJSON struct {
	Href        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseDataLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataLinkJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataPerformance struct {
	Duration  float64                              `json:"duration,required"`
	EntryType string                               `json:"entryType,required"`
	Name      string                               `json:"name,required"`
	StartTime float64                              `json:"startTime,required"`
	JSON      resultGetResponseDataPerformanceJSON `json:"-"`
}

// resultGetResponseDataPerformanceJSON contains the JSON metadata for the struct
// [ResultGetResponseDataPerformance]
type resultGetResponseDataPerformanceJSON struct {
	Duration    apijson.Field
	EntryType   apijson.Field
	Name        apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseDataPerformance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataPerformanceJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataRequest struct {
	Request  ResultGetResponseDataRequestsRequest   `json:"request,required"`
	Response ResultGetResponseDataRequestsResponse  `json:"response,required"`
	Requests []ResultGetResponseDataRequestsRequest `json:"requests"`
	JSON     resultGetResponseDataRequestJSON       `json:"-"`
}

// resultGetResponseDataRequestJSON contains the JSON metadata for the struct
// [ResultGetResponseDataRequest]
type resultGetResponseDataRequestJSON struct {
	Request     apijson.Field
	Response    apijson.Field
	Requests    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseDataRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataRequestJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataRequestsRequest struct {
	DocumentURL          string                                               `json:"documentURL,required"`
	HasUserGesture       bool                                                 `json:"hasUserGesture,required"`
	Initiator            ResultGetResponseDataRequestsRequestInitiator        `json:"initiator,required"`
	RedirectHasExtraInfo bool                                                 `json:"redirectHasExtraInfo,required"`
	Request              ResultGetResponseDataRequestsRequestRequest          `json:"request,required"`
	RequestID            string                                               `json:"requestId,required"`
	Type                 string                                               `json:"type,required"`
	WallTime             float64                                              `json:"wallTime,required"`
	FrameID              string                                               `json:"frameId"`
	LoaderID             string                                               `json:"loaderId"`
	PrimaryRequest       bool                                                 `json:"primaryRequest"`
	RedirectResponse     ResultGetResponseDataRequestsRequestRedirectResponse `json:"redirectResponse"`
	JSON                 resultGetResponseDataRequestsRequestJSON             `json:"-"`
}

// resultGetResponseDataRequestsRequestJSON contains the JSON metadata for the
// struct [ResultGetResponseDataRequestsRequest]
type resultGetResponseDataRequestsRequestJSON struct {
	DocumentURL          apijson.Field
	HasUserGesture       apijson.Field
	Initiator            apijson.Field
	RedirectHasExtraInfo apijson.Field
	Request              apijson.Field
	RequestID            apijson.Field
	Type                 apijson.Field
	WallTime             apijson.Field
	FrameID              apijson.Field
	LoaderID             apijson.Field
	PrimaryRequest       apijson.Field
	RedirectResponse     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ResultGetResponseDataRequestsRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataRequestsRequestJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataRequestsRequestInitiator struct {
	Host string                                            `json:"host,required"`
	Type string                                            `json:"type,required"`
	URL  string                                            `json:"url,required"`
	JSON resultGetResponseDataRequestsRequestInitiatorJSON `json:"-"`
}

// resultGetResponseDataRequestsRequestInitiatorJSON contains the JSON metadata for
// the struct [ResultGetResponseDataRequestsRequestInitiator]
type resultGetResponseDataRequestsRequestInitiatorJSON struct {
	Host        apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseDataRequestsRequestInitiator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataRequestsRequestInitiatorJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataRequestsRequestRequest struct {
	InitialPriority  string                                          `json:"initialPriority,required"`
	IsSameSite       bool                                            `json:"isSameSite,required"`
	Method           string                                          `json:"method,required"`
	MixedContentType string                                          `json:"mixedContentType,required"`
	ReferrerPolicy   string                                          `json:"referrerPolicy,required"`
	URL              string                                          `json:"url,required"`
	Headers          interface{}                                     `json:"headers"`
	JSON             resultGetResponseDataRequestsRequestRequestJSON `json:"-"`
}

// resultGetResponseDataRequestsRequestRequestJSON contains the JSON metadata for
// the struct [ResultGetResponseDataRequestsRequestRequest]
type resultGetResponseDataRequestsRequestRequestJSON struct {
	InitialPriority  apijson.Field
	IsSameSite       apijson.Field
	Method           apijson.Field
	MixedContentType apijson.Field
	ReferrerPolicy   apijson.Field
	URL              apijson.Field
	Headers          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResultGetResponseDataRequestsRequestRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataRequestsRequestRequestJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataRequestsRequestRedirectResponse struct {
	Charset         string                                                               `json:"charset,required"`
	MimeType        string                                                               `json:"mimeType,required"`
	Protocol        string                                                               `json:"protocol,required"`
	RemoteIPAddress string                                                               `json:"remoteIPAddress,required"`
	RemotePort      float64                                                              `json:"remotePort,required"`
	SecurityHeaders []ResultGetResponseDataRequestsRequestRedirectResponseSecurityHeader `json:"securityHeaders,required"`
	SecurityState   string                                                               `json:"securityState,required"`
	Status          float64                                                              `json:"status,required"`
	StatusText      string                                                               `json:"statusText,required"`
	URL             string                                                               `json:"url,required"`
	Headers         interface{}                                                          `json:"headers"`
	JSON            resultGetResponseDataRequestsRequestRedirectResponseJSON             `json:"-"`
}

// resultGetResponseDataRequestsRequestRedirectResponseJSON contains the JSON
// metadata for the struct [ResultGetResponseDataRequestsRequestRedirectResponse]
type resultGetResponseDataRequestsRequestRedirectResponseJSON struct {
	Charset         apijson.Field
	MimeType        apijson.Field
	Protocol        apijson.Field
	RemoteIPAddress apijson.Field
	RemotePort      apijson.Field
	SecurityHeaders apijson.Field
	SecurityState   apijson.Field
	Status          apijson.Field
	StatusText      apijson.Field
	URL             apijson.Field
	Headers         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResultGetResponseDataRequestsRequestRedirectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataRequestsRequestRedirectResponseJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataRequestsRequestRedirectResponseSecurityHeader struct {
	Name  string                                                                 `json:"name,required"`
	Value string                                                                 `json:"value,required"`
	JSON  resultGetResponseDataRequestsRequestRedirectResponseSecurityHeaderJSON `json:"-"`
}

// resultGetResponseDataRequestsRequestRedirectResponseSecurityHeaderJSON contains
// the JSON metadata for the struct
// [ResultGetResponseDataRequestsRequestRedirectResponseSecurityHeader]
type resultGetResponseDataRequestsRequestRedirectResponseSecurityHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseDataRequestsRequestRedirectResponseSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataRequestsRequestRedirectResponseSecurityHeaderJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataRequestsResponse struct {
	ASN               ResultGetResponseDataRequestsResponseASN      `json:"asn,required"`
	DataLength        float64                                       `json:"dataLength,required"`
	EncodedDataLength float64                                       `json:"encodedDataLength,required"`
	Geoip             ResultGetResponseDataRequestsResponseGeoip    `json:"geoip,required"`
	HasExtraInfo      bool                                          `json:"hasExtraInfo,required"`
	RequestID         string                                        `json:"requestId,required"`
	Response          ResultGetResponseDataRequestsResponseResponse `json:"response,required"`
	Size              float64                                       `json:"size,required"`
	Type              string                                        `json:"type,required"`
	ContentAvailable  bool                                          `json:"contentAvailable"`
	Hash              string                                        `json:"hash"`
	JSON              resultGetResponseDataRequestsResponseJSON     `json:"-"`
}

// resultGetResponseDataRequestsResponseJSON contains the JSON metadata for the
// struct [ResultGetResponseDataRequestsResponse]
type resultGetResponseDataRequestsResponseJSON struct {
	ASN               apijson.Field
	DataLength        apijson.Field
	EncodedDataLength apijson.Field
	Geoip             apijson.Field
	HasExtraInfo      apijson.Field
	RequestID         apijson.Field
	Response          apijson.Field
	Size              apijson.Field
	Type              apijson.Field
	ContentAvailable  apijson.Field
	Hash              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResultGetResponseDataRequestsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataRequestsResponseJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataRequestsResponseASN struct {
	ASN         string                                       `json:"asn,required"`
	Country     string                                       `json:"country,required"`
	Description string                                       `json:"description,required"`
	IP          string                                       `json:"ip,required"`
	Name        string                                       `json:"name,required"`
	Org         string                                       `json:"org,required"`
	JSON        resultGetResponseDataRequestsResponseASNJSON `json:"-"`
}

// resultGetResponseDataRequestsResponseASNJSON contains the JSON metadata for the
// struct [ResultGetResponseDataRequestsResponseASN]
type resultGetResponseDataRequestsResponseASNJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	Description apijson.Field
	IP          apijson.Field
	Name        apijson.Field
	Org         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseDataRequestsResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataRequestsResponseASNJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataRequestsResponseGeoip struct {
	City        string                                         `json:"city,required"`
	Country     string                                         `json:"country,required"`
	CountryName string                                         `json:"country_name,required"`
	GeonameID   string                                         `json:"geonameId,required"`
	Ll          []interface{}                                  `json:"ll,required"`
	Region      string                                         `json:"region,required"`
	JSON        resultGetResponseDataRequestsResponseGeoipJSON `json:"-"`
}

// resultGetResponseDataRequestsResponseGeoipJSON contains the JSON metadata for
// the struct [ResultGetResponseDataRequestsResponseGeoip]
type resultGetResponseDataRequestsResponseGeoipJSON struct {
	City        apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	GeonameID   apijson.Field
	Ll          apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseDataRequestsResponseGeoip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataRequestsResponseGeoipJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataRequestsResponseResponse struct {
	Charset         string                                                        `json:"charset,required"`
	MimeType        string                                                        `json:"mimeType,required"`
	Protocol        string                                                        `json:"protocol,required"`
	RemoteIPAddress string                                                        `json:"remoteIPAddress,required"`
	RemotePort      float64                                                       `json:"remotePort,required"`
	SecurityDetails ResultGetResponseDataRequestsResponseResponseSecurityDetails  `json:"securityDetails,required"`
	SecurityHeaders []ResultGetResponseDataRequestsResponseResponseSecurityHeader `json:"securityHeaders,required"`
	SecurityState   string                                                        `json:"securityState,required"`
	Status          float64                                                       `json:"status,required"`
	StatusText      string                                                        `json:"statusText,required"`
	URL             string                                                        `json:"url,required"`
	Headers         interface{}                                                   `json:"headers"`
	JSON            resultGetResponseDataRequestsResponseResponseJSON             `json:"-"`
}

// resultGetResponseDataRequestsResponseResponseJSON contains the JSON metadata for
// the struct [ResultGetResponseDataRequestsResponseResponse]
type resultGetResponseDataRequestsResponseResponseJSON struct {
	Charset         apijson.Field
	MimeType        apijson.Field
	Protocol        apijson.Field
	RemoteIPAddress apijson.Field
	RemotePort      apijson.Field
	SecurityDetails apijson.Field
	SecurityHeaders apijson.Field
	SecurityState   apijson.Field
	Status          apijson.Field
	StatusText      apijson.Field
	URL             apijson.Field
	Headers         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResultGetResponseDataRequestsResponseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataRequestsResponseResponseJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataRequestsResponseResponseSecurityDetails struct {
	CertificateID                     float64                                                          `json:"certificateId,required"`
	CertificateTransparencyCompliance string                                                           `json:"certificateTransparencyCompliance,required"`
	Cipher                            string                                                           `json:"cipher,required"`
	EncryptedClientHello              bool                                                             `json:"encryptedClientHello,required"`
	Issuer                            string                                                           `json:"issuer,required"`
	KeyExchange                       string                                                           `json:"keyExchange,required"`
	KeyExchangeGroup                  string                                                           `json:"keyExchangeGroup,required"`
	Protocol                          string                                                           `json:"protocol,required"`
	SanList                           []string                                                         `json:"sanList,required"`
	ServerSignatureAlgorithm          float64                                                          `json:"serverSignatureAlgorithm,required"`
	SubjectName                       string                                                           `json:"subjectName,required"`
	ValidFrom                         float64                                                          `json:"validFrom,required"`
	ValidTo                           float64                                                          `json:"validTo,required"`
	JSON                              resultGetResponseDataRequestsResponseResponseSecurityDetailsJSON `json:"-"`
}

// resultGetResponseDataRequestsResponseResponseSecurityDetailsJSON contains the
// JSON metadata for the struct
// [ResultGetResponseDataRequestsResponseResponseSecurityDetails]
type resultGetResponseDataRequestsResponseResponseSecurityDetailsJSON struct {
	CertificateID                     apijson.Field
	CertificateTransparencyCompliance apijson.Field
	Cipher                            apijson.Field
	EncryptedClientHello              apijson.Field
	Issuer                            apijson.Field
	KeyExchange                       apijson.Field
	KeyExchangeGroup                  apijson.Field
	Protocol                          apijson.Field
	SanList                           apijson.Field
	ServerSignatureAlgorithm          apijson.Field
	SubjectName                       apijson.Field
	ValidFrom                         apijson.Field
	ValidTo                           apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *ResultGetResponseDataRequestsResponseResponseSecurityDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataRequestsResponseResponseSecurityDetailsJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseDataRequestsResponseResponseSecurityHeader struct {
	Name  string                                                          `json:"name,required"`
	Value string                                                          `json:"value,required"`
	JSON  resultGetResponseDataRequestsResponseResponseSecurityHeaderJSON `json:"-"`
}

// resultGetResponseDataRequestsResponseResponseSecurityHeaderJSON contains the
// JSON metadata for the struct
// [ResultGetResponseDataRequestsResponseResponseSecurityHeader]
type resultGetResponseDataRequestsResponseResponseSecurityHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseDataRequestsResponseResponseSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseDataRequestsResponseResponseSecurityHeaderJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseLists struct {
	ASNs         []string                            `json:"asns,required"`
	Certificates []ResultGetResponseListsCertificate `json:"certificates,required"`
	Continents   []string                            `json:"continents,required"`
	Countries    []string                            `json:"countries,required"`
	Domains      []string                            `json:"domains,required"`
	Hashes       []string                            `json:"hashes,required"`
	IPs          []string                            `json:"ips,required"`
	LinkDomains  []string                            `json:"linkDomains,required"`
	Servers      []string                            `json:"servers,required"`
	URLs         []string                            `json:"urls,required"`
	JSON         resultGetResponseListsJSON          `json:"-"`
}

// resultGetResponseListsJSON contains the JSON metadata for the struct
// [ResultGetResponseLists]
type resultGetResponseListsJSON struct {
	ASNs         apijson.Field
	Certificates apijson.Field
	Continents   apijson.Field
	Countries    apijson.Field
	Domains      apijson.Field
	Hashes       apijson.Field
	IPs          apijson.Field
	LinkDomains  apijson.Field
	Servers      apijson.Field
	URLs         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResultGetResponseLists) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseListsJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseListsCertificate struct {
	Issuer      string                                `json:"issuer,required"`
	SubjectName string                                `json:"subjectName,required"`
	ValidFrom   float64                               `json:"validFrom,required"`
	ValidTo     float64                               `json:"validTo,required"`
	JSON        resultGetResponseListsCertificateJSON `json:"-"`
}

// resultGetResponseListsCertificateJSON contains the JSON metadata for the struct
// [ResultGetResponseListsCertificate]
type resultGetResponseListsCertificateJSON struct {
	Issuer      apijson.Field
	SubjectName apijson.Field
	ValidFrom   apijson.Field
	ValidTo     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseListsCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseListsCertificateJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMeta struct {
	Processors ResultGetResponseMetaProcessors `json:"processors,required"`
	JSON       resultGetResponseMetaJSON       `json:"-"`
}

// resultGetResponseMetaJSON contains the JSON metadata for the struct
// [ResultGetResponseMeta]
type resultGetResponseMetaJSON struct {
	Processors  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessors struct {
	ASN              ResultGetResponseMetaProcessorsASN              `json:"asn,required"`
	DNS              ResultGetResponseMetaProcessorsDNS              `json:"dns,required"`
	DomainCategories ResultGetResponseMetaProcessorsDomainCategories `json:"domainCategories,required"`
	Geoip            ResultGetResponseMetaProcessorsGeoip            `json:"geoip,required"`
	Phishing         ResultGetResponseMetaProcessorsPhishing         `json:"phishing,required"`
	RadarRank        ResultGetResponseMetaProcessorsRadarRank        `json:"radarRank,required"`
	Wappa            ResultGetResponseMetaProcessorsWappa            `json:"wappa,required"`
	URLCategories    ResultGetResponseMetaProcessorsURLCategories    `json:"urlCategories"`
	JSON             resultGetResponseMetaProcessorsJSON             `json:"-"`
}

// resultGetResponseMetaProcessorsJSON contains the JSON metadata for the struct
// [ResultGetResponseMetaProcessors]
type resultGetResponseMetaProcessorsJSON struct {
	ASN              apijson.Field
	DNS              apijson.Field
	DomainCategories apijson.Field
	Geoip            apijson.Field
	Phishing         apijson.Field
	RadarRank        apijson.Field
	Wappa            apijson.Field
	URLCategories    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsASN struct {
	Data []ResultGetResponseMetaProcessorsASNData `json:"data,required"`
	JSON resultGetResponseMetaProcessorsASNJSON   `json:"-"`
}

// resultGetResponseMetaProcessorsASNJSON contains the JSON metadata for the struct
// [ResultGetResponseMetaProcessorsASN]
type resultGetResponseMetaProcessorsASNJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsASNJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsASNData struct {
	ASN         string                                     `json:"asn,required"`
	Country     string                                     `json:"country,required"`
	Description string                                     `json:"description,required"`
	IP          string                                     `json:"ip,required"`
	Name        string                                     `json:"name,required"`
	JSON        resultGetResponseMetaProcessorsASNDataJSON `json:"-"`
}

// resultGetResponseMetaProcessorsASNDataJSON contains the JSON metadata for the
// struct [ResultGetResponseMetaProcessorsASNData]
type resultGetResponseMetaProcessorsASNDataJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	Description apijson.Field
	IP          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsASNData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsASNDataJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsDNS struct {
	Data []ResultGetResponseMetaProcessorsDNSData `json:"data,required"`
	JSON resultGetResponseMetaProcessorsDNSJSON   `json:"-"`
}

// resultGetResponseMetaProcessorsDNSJSON contains the JSON metadata for the struct
// [ResultGetResponseMetaProcessorsDNS]
type resultGetResponseMetaProcessorsDNSJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsDNSJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsDNSData struct {
	Address     string                                     `json:"address,required"`
	DNSSECValid bool                                       `json:"dnssec_valid,required"`
	Name        string                                     `json:"name,required"`
	Type        string                                     `json:"type,required"`
	JSON        resultGetResponseMetaProcessorsDNSDataJSON `json:"-"`
}

// resultGetResponseMetaProcessorsDNSDataJSON contains the JSON metadata for the
// struct [ResultGetResponseMetaProcessorsDNSData]
type resultGetResponseMetaProcessorsDNSDataJSON struct {
	Address     apijson.Field
	DNSSECValid apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsDNSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsDNSDataJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsDomainCategories struct {
	Data []ResultGetResponseMetaProcessorsDomainCategoriesData `json:"data,required"`
	JSON resultGetResponseMetaProcessorsDomainCategoriesJSON   `json:"-"`
}

// resultGetResponseMetaProcessorsDomainCategoriesJSON contains the JSON metadata
// for the struct [ResultGetResponseMetaProcessorsDomainCategories]
type resultGetResponseMetaProcessorsDomainCategoriesJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsDomainCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsDomainCategoriesJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsDomainCategoriesData struct {
	Inherited interface{}                                             `json:"inherited,required"`
	IsPrimary bool                                                    `json:"isPrimary,required"`
	Name      string                                                  `json:"name,required"`
	JSON      resultGetResponseMetaProcessorsDomainCategoriesDataJSON `json:"-"`
}

// resultGetResponseMetaProcessorsDomainCategoriesDataJSON contains the JSON
// metadata for the struct [ResultGetResponseMetaProcessorsDomainCategoriesData]
type resultGetResponseMetaProcessorsDomainCategoriesDataJSON struct {
	Inherited   apijson.Field
	IsPrimary   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsDomainCategoriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsDomainCategoriesDataJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsGeoip struct {
	Data []ResultGetResponseMetaProcessorsGeoipData `json:"data,required"`
	JSON resultGetResponseMetaProcessorsGeoipJSON   `json:"-"`
}

// resultGetResponseMetaProcessorsGeoipJSON contains the JSON metadata for the
// struct [ResultGetResponseMetaProcessorsGeoip]
type resultGetResponseMetaProcessorsGeoipJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsGeoip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsGeoipJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsGeoipData struct {
	Geoip ResultGetResponseMetaProcessorsGeoipDataGeoip `json:"geoip,required"`
	IP    string                                        `json:"ip,required"`
	JSON  resultGetResponseMetaProcessorsGeoipDataJSON  `json:"-"`
}

// resultGetResponseMetaProcessorsGeoipDataJSON contains the JSON metadata for the
// struct [ResultGetResponseMetaProcessorsGeoipData]
type resultGetResponseMetaProcessorsGeoipDataJSON struct {
	Geoip       apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsGeoipData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsGeoipDataJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsGeoipDataGeoip struct {
	City        string                                            `json:"city,required"`
	Country     string                                            `json:"country,required"`
	CountryName string                                            `json:"country_name,required"`
	Ll          []float64                                         `json:"ll,required"`
	Region      string                                            `json:"region,required"`
	JSON        resultGetResponseMetaProcessorsGeoipDataGeoipJSON `json:"-"`
}

// resultGetResponseMetaProcessorsGeoipDataGeoipJSON contains the JSON metadata for
// the struct [ResultGetResponseMetaProcessorsGeoipDataGeoip]
type resultGetResponseMetaProcessorsGeoipDataGeoipJSON struct {
	City        apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	Ll          apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsGeoipDataGeoip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsGeoipDataGeoipJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsPhishing struct {
	Data []string                                    `json:"data,required"`
	JSON resultGetResponseMetaProcessorsPhishingJSON `json:"-"`
}

// resultGetResponseMetaProcessorsPhishingJSON contains the JSON metadata for the
// struct [ResultGetResponseMetaProcessorsPhishing]
type resultGetResponseMetaProcessorsPhishingJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsPhishing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsPhishingJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsRadarRank struct {
	Data []ResultGetResponseMetaProcessorsRadarRankData `json:"data,required"`
	JSON resultGetResponseMetaProcessorsRadarRankJSON   `json:"-"`
}

// resultGetResponseMetaProcessorsRadarRankJSON contains the JSON metadata for the
// struct [ResultGetResponseMetaProcessorsRadarRank]
type resultGetResponseMetaProcessorsRadarRankJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsRadarRank) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsRadarRankJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsRadarRankData struct {
	Bucket   string                                           `json:"bucket,required"`
	Hostname string                                           `json:"hostname,required"`
	Rank     float64                                          `json:"rank"`
	JSON     resultGetResponseMetaProcessorsRadarRankDataJSON `json:"-"`
}

// resultGetResponseMetaProcessorsRadarRankDataJSON contains the JSON metadata for
// the struct [ResultGetResponseMetaProcessorsRadarRankData]
type resultGetResponseMetaProcessorsRadarRankDataJSON struct {
	Bucket      apijson.Field
	Hostname    apijson.Field
	Rank        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsRadarRankData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsRadarRankDataJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsWappa struct {
	Data []ResultGetResponseMetaProcessorsWappaData `json:"data,required"`
	JSON resultGetResponseMetaProcessorsWappaJSON   `json:"-"`
}

// resultGetResponseMetaProcessorsWappaJSON contains the JSON metadata for the
// struct [ResultGetResponseMetaProcessorsWappa]
type resultGetResponseMetaProcessorsWappaJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsWappa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsWappaJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsWappaData struct {
	App             string                                               `json:"app,required"`
	Categories      []ResultGetResponseMetaProcessorsWappaDataCategory   `json:"categories,required"`
	Confidence      []ResultGetResponseMetaProcessorsWappaDataConfidence `json:"confidence,required"`
	ConfidenceTotal float64                                              `json:"confidenceTotal,required"`
	Icon            string                                               `json:"icon,required"`
	Website         string                                               `json:"website,required"`
	JSON            resultGetResponseMetaProcessorsWappaDataJSON         `json:"-"`
}

// resultGetResponseMetaProcessorsWappaDataJSON contains the JSON metadata for the
// struct [ResultGetResponseMetaProcessorsWappaData]
type resultGetResponseMetaProcessorsWappaDataJSON struct {
	App             apijson.Field
	Categories      apijson.Field
	Confidence      apijson.Field
	ConfidenceTotal apijson.Field
	Icon            apijson.Field
	Website         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsWappaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsWappaDataJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsWappaDataCategory struct {
	Name     string                                               `json:"name,required"`
	Priority float64                                              `json:"priority,required"`
	JSON     resultGetResponseMetaProcessorsWappaDataCategoryJSON `json:"-"`
}

// resultGetResponseMetaProcessorsWappaDataCategoryJSON contains the JSON metadata
// for the struct [ResultGetResponseMetaProcessorsWappaDataCategory]
type resultGetResponseMetaProcessorsWappaDataCategoryJSON struct {
	Name        apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsWappaDataCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsWappaDataCategoryJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsWappaDataConfidence struct {
	Confidence  float64                                                `json:"confidence,required"`
	Name        string                                                 `json:"name,required"`
	Pattern     string                                                 `json:"pattern,required"`
	PatternType string                                                 `json:"patternType,required"`
	JSON        resultGetResponseMetaProcessorsWappaDataConfidenceJSON `json:"-"`
}

// resultGetResponseMetaProcessorsWappaDataConfidenceJSON contains the JSON
// metadata for the struct [ResultGetResponseMetaProcessorsWappaDataConfidence]
type resultGetResponseMetaProcessorsWappaDataConfidenceJSON struct {
	Confidence  apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	PatternType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsWappaDataConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsWappaDataConfidenceJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsURLCategories struct {
	Data []ResultGetResponseMetaProcessorsURLCategoriesData `json:"data,required"`
	JSON resultGetResponseMetaProcessorsURLCategoriesJSON   `json:"-"`
}

// resultGetResponseMetaProcessorsURLCategoriesJSON contains the JSON metadata for
// the struct [ResultGetResponseMetaProcessorsURLCategories]
type resultGetResponseMetaProcessorsURLCategoriesJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsURLCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsURLCategoriesJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsURLCategoriesData struct {
	Content   []ResultGetResponseMetaProcessorsURLCategoriesDataContent `json:"content,required"`
	Inherited ResultGetResponseMetaProcessorsURLCategoriesDataInherited `json:"inherited,required"`
	Name      string                                                    `json:"name,required"`
	Risks     []ResultGetResponseMetaProcessorsURLCategoriesDataRisk    `json:"risks,required"`
	JSON      resultGetResponseMetaProcessorsURLCategoriesDataJSON      `json:"-"`
}

// resultGetResponseMetaProcessorsURLCategoriesDataJSON contains the JSON metadata
// for the struct [ResultGetResponseMetaProcessorsURLCategoriesData]
type resultGetResponseMetaProcessorsURLCategoriesDataJSON struct {
	Content     apijson.Field
	Inherited   apijson.Field
	Name        apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsURLCategoriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsURLCategoriesDataJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsURLCategoriesDataContent struct {
	ID              float64                                                     `json:"id,required"`
	Name            string                                                      `json:"name,required"`
	SuperCategoryID float64                                                     `json:"super_category_id,required"`
	JSON            resultGetResponseMetaProcessorsURLCategoriesDataContentJSON `json:"-"`
}

// resultGetResponseMetaProcessorsURLCategoriesDataContentJSON contains the JSON
// metadata for the struct
// [ResultGetResponseMetaProcessorsURLCategoriesDataContent]
type resultGetResponseMetaProcessorsURLCategoriesDataContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsURLCategoriesDataContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsURLCategoriesDataContentJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsURLCategoriesDataInherited struct {
	Content []ResultGetResponseMetaProcessorsURLCategoriesDataInheritedContent `json:"content,required"`
	From    string                                                             `json:"from,required"`
	Risks   []ResultGetResponseMetaProcessorsURLCategoriesDataInheritedRisk    `json:"risks,required"`
	JSON    resultGetResponseMetaProcessorsURLCategoriesDataInheritedJSON      `json:"-"`
}

// resultGetResponseMetaProcessorsURLCategoriesDataInheritedJSON contains the JSON
// metadata for the struct
// [ResultGetResponseMetaProcessorsURLCategoriesDataInherited]
type resultGetResponseMetaProcessorsURLCategoriesDataInheritedJSON struct {
	Content     apijson.Field
	From        apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsURLCategoriesDataInherited) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsURLCategoriesDataInheritedJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsURLCategoriesDataInheritedContent struct {
	ID              float64                                                              `json:"id,required"`
	Name            string                                                               `json:"name,required"`
	SuperCategoryID float64                                                              `json:"super_category_id,required"`
	JSON            resultGetResponseMetaProcessorsURLCategoriesDataInheritedContentJSON `json:"-"`
}

// resultGetResponseMetaProcessorsURLCategoriesDataInheritedContentJSON contains
// the JSON metadata for the struct
// [ResultGetResponseMetaProcessorsURLCategoriesDataInheritedContent]
type resultGetResponseMetaProcessorsURLCategoriesDataInheritedContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsURLCategoriesDataInheritedContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsURLCategoriesDataInheritedContentJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsURLCategoriesDataInheritedRisk struct {
	ID              float64                                                           `json:"id,required"`
	Name            string                                                            `json:"name,required"`
	SuperCategoryID float64                                                           `json:"super_category_id,required"`
	JSON            resultGetResponseMetaProcessorsURLCategoriesDataInheritedRiskJSON `json:"-"`
}

// resultGetResponseMetaProcessorsURLCategoriesDataInheritedRiskJSON contains the
// JSON metadata for the struct
// [ResultGetResponseMetaProcessorsURLCategoriesDataInheritedRisk]
type resultGetResponseMetaProcessorsURLCategoriesDataInheritedRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsURLCategoriesDataInheritedRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsURLCategoriesDataInheritedRiskJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseMetaProcessorsURLCategoriesDataRisk struct {
	ID              float64                                                  `json:"id,required"`
	Name            string                                                   `json:"name,required"`
	SuperCategoryID float64                                                  `json:"super_category_id,required"`
	JSON            resultGetResponseMetaProcessorsURLCategoriesDataRiskJSON `json:"-"`
}

// resultGetResponseMetaProcessorsURLCategoriesDataRiskJSON contains the JSON
// metadata for the struct [ResultGetResponseMetaProcessorsURLCategoriesDataRisk]
type resultGetResponseMetaProcessorsURLCategoriesDataRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResultGetResponseMetaProcessorsURLCategoriesDataRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseMetaProcessorsURLCategoriesDataRiskJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponsePage struct {
	ApexDomain   string                          `json:"apexDomain,required"`
	ASN          string                          `json:"asn,required"`
	Asnname      string                          `json:"asnname,required"`
	City         string                          `json:"city,required"`
	Country      string                          `json:"country,required"`
	Domain       string                          `json:"domain,required"`
	IP           string                          `json:"ip,required"`
	MimeType     string                          `json:"mimeType,required"`
	Server       string                          `json:"server,required"`
	Status       string                          `json:"status,required"`
	Title        string                          `json:"title,required"`
	TLSAgeDays   float64                         `json:"tlsAgeDays,required"`
	TLSIssuer    string                          `json:"tlsIssuer,required"`
	TLSValidDays float64                         `json:"tlsValidDays,required"`
	TLSValidFrom string                          `json:"tlsValidFrom,required"`
	URL          string                          `json:"url,required"`
	Screenshot   ResultGetResponsePageScreenshot `json:"screenshot"`
	JSON         resultGetResponsePageJSON       `json:"-"`
}

// resultGetResponsePageJSON contains the JSON metadata for the struct
// [ResultGetResponsePage]
type resultGetResponsePageJSON struct {
	ApexDomain   apijson.Field
	ASN          apijson.Field
	Asnname      apijson.Field
	City         apijson.Field
	Country      apijson.Field
	Domain       apijson.Field
	IP           apijson.Field
	MimeType     apijson.Field
	Server       apijson.Field
	Status       apijson.Field
	Title        apijson.Field
	TLSAgeDays   apijson.Field
	TLSIssuer    apijson.Field
	TLSValidDays apijson.Field
	TLSValidFrom apijson.Field
	URL          apijson.Field
	Screenshot   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResultGetResponsePage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponsePageJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponsePageScreenshot struct {
	Dhash   string                              `json:"dhash,required"`
	Mm3Hash float64                             `json:"mm3Hash,required"`
	Name    string                              `json:"name,required"`
	Phash   string                              `json:"phash,required"`
	JSON    resultGetResponsePageScreenshotJSON `json:"-"`
}

// resultGetResponsePageScreenshotJSON contains the JSON metadata for the struct
// [ResultGetResponsePageScreenshot]
type resultGetResponsePageScreenshotJSON struct {
	Dhash       apijson.Field
	Mm3Hash     apijson.Field
	Name        apijson.Field
	Phash       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponsePageScreenshot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponsePageScreenshotJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseScanner struct {
	Colo    string                       `json:"colo,required"`
	Country string                       `json:"country,required"`
	JSON    resultGetResponseScannerJSON `json:"-"`
}

// resultGetResponseScannerJSON contains the JSON metadata for the struct
// [ResultGetResponseScanner]
type resultGetResponseScannerJSON struct {
	Colo        apijson.Field
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseScanner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseScannerJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseStats struct {
	DomainStats      []ResultGetResponseStatsDomainStat   `json:"domainStats,required"`
	IPStats          []ResultGetResponseStatsIPStat       `json:"ipStats,required"`
	IPv6Percentage   float64                              `json:"IPv6Percentage,required"`
	Malicious        float64                              `json:"malicious,required"`
	ProtocolStats    []ResultGetResponseStatsProtocolStat `json:"protocolStats,required"`
	ResourceStats    []ResultGetResponseStatsResourceStat `json:"resourceStats,required"`
	SecurePercentage float64                              `json:"securePercentage,required"`
	SecureRequests   float64                              `json:"secureRequests,required"`
	ServerStats      []ResultGetResponseStatsServerStat   `json:"serverStats,required"`
	TLSStats         []ResultGetResponseStatsTLSStat      `json:"tlsStats,required"`
	TotalLinks       float64                              `json:"totalLinks,required"`
	UniqASNs         float64                              `json:"uniqASNs,required"`
	UniqCountries    float64                              `json:"uniqCountries,required"`
	JSON             resultGetResponseStatsJSON           `json:"-"`
}

// resultGetResponseStatsJSON contains the JSON metadata for the struct
// [ResultGetResponseStats]
type resultGetResponseStatsJSON struct {
	DomainStats      apijson.Field
	IPStats          apijson.Field
	IPv6Percentage   apijson.Field
	Malicious        apijson.Field
	ProtocolStats    apijson.Field
	ResourceStats    apijson.Field
	SecurePercentage apijson.Field
	SecureRequests   apijson.Field
	ServerStats      apijson.Field
	TLSStats         apijson.Field
	TotalLinks       apijson.Field
	UniqASNs         apijson.Field
	UniqCountries    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResultGetResponseStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseStatsJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseStatsDomainStat struct {
	Count       float64                              `json:"count,required"`
	Countries   []string                             `json:"countries,required"`
	Domain      string                               `json:"domain,required"`
	EncodedSize float64                              `json:"encodedSize,required"`
	Index       float64                              `json:"index,required"`
	Initiators  []string                             `json:"initiators,required"`
	IPs         []string                             `json:"ips,required"`
	Redirects   float64                              `json:"redirects,required"`
	Size        float64                              `json:"size,required"`
	JSON        resultGetResponseStatsDomainStatJSON `json:"-"`
}

// resultGetResponseStatsDomainStatJSON contains the JSON metadata for the struct
// [ResultGetResponseStatsDomainStat]
type resultGetResponseStatsDomainStatJSON struct {
	Count       apijson.Field
	Countries   apijson.Field
	Domain      apijson.Field
	EncodedSize apijson.Field
	Index       apijson.Field
	Initiators  apijson.Field
	IPs         apijson.Field
	Redirects   apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseStatsDomainStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseStatsDomainStatJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseStatsIPStat struct {
	ASN         ResultGetResponseStatsIPStatsASN   `json:"asn,required"`
	Countries   []string                           `json:"countries,required"`
	Domains     []string                           `json:"domains,required"`
	EncodedSize float64                            `json:"encodedSize,required"`
	Geoip       ResultGetResponseStatsIPStatsGeoip `json:"geoip,required"`
	Index       float64                            `json:"index,required"`
	IP          string                             `json:"ip,required"`
	IPV6        bool                               `json:"ipv6,required"`
	Redirects   float64                            `json:"redirects,required"`
	Requests    float64                            `json:"requests,required"`
	Size        float64                            `json:"size,required"`
	Count       float64                            `json:"count"`
	JSON        resultGetResponseStatsIPStatJSON   `json:"-"`
}

// resultGetResponseStatsIPStatJSON contains the JSON metadata for the struct
// [ResultGetResponseStatsIPStat]
type resultGetResponseStatsIPStatJSON struct {
	ASN         apijson.Field
	Countries   apijson.Field
	Domains     apijson.Field
	EncodedSize apijson.Field
	Geoip       apijson.Field
	Index       apijson.Field
	IP          apijson.Field
	IPV6        apijson.Field
	Redirects   apijson.Field
	Requests    apijson.Field
	Size        apijson.Field
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseStatsIPStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseStatsIPStatJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseStatsIPStatsASN struct {
	ASN         string                               `json:"asn,required"`
	Country     string                               `json:"country,required"`
	Description string                               `json:"description,required"`
	IP          string                               `json:"ip,required"`
	Name        string                               `json:"name,required"`
	Org         string                               `json:"org,required"`
	JSON        resultGetResponseStatsIPStatsASNJSON `json:"-"`
}

// resultGetResponseStatsIPStatsASNJSON contains the JSON metadata for the struct
// [ResultGetResponseStatsIPStatsASN]
type resultGetResponseStatsIPStatsASNJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	Description apijson.Field
	IP          apijson.Field
	Name        apijson.Field
	Org         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseStatsIPStatsASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseStatsIPStatsASNJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseStatsIPStatsGeoip struct {
	City        string                                 `json:"city,required"`
	Country     string                                 `json:"country,required"`
	CountryName string                                 `json:"country_name,required"`
	Ll          []float64                              `json:"ll,required"`
	Region      string                                 `json:"region,required"`
	JSON        resultGetResponseStatsIPStatsGeoipJSON `json:"-"`
}

// resultGetResponseStatsIPStatsGeoipJSON contains the JSON metadata for the struct
// [ResultGetResponseStatsIPStatsGeoip]
type resultGetResponseStatsIPStatsGeoipJSON struct {
	City        apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	Ll          apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseStatsIPStatsGeoip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseStatsIPStatsGeoipJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseStatsProtocolStat struct {
	Count       float64                                `json:"count,required"`
	Countries   []string                               `json:"countries,required"`
	EncodedSize float64                                `json:"encodedSize,required"`
	IPs         []string                               `json:"ips,required"`
	Protocol    string                                 `json:"protocol,required"`
	Size        float64                                `json:"size,required"`
	JSON        resultGetResponseStatsProtocolStatJSON `json:"-"`
}

// resultGetResponseStatsProtocolStatJSON contains the JSON metadata for the struct
// [ResultGetResponseStatsProtocolStat]
type resultGetResponseStatsProtocolStatJSON struct {
	Count       apijson.Field
	Countries   apijson.Field
	EncodedSize apijson.Field
	IPs         apijson.Field
	Protocol    apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseStatsProtocolStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseStatsProtocolStatJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseStatsResourceStat struct {
	Compression float64                                `json:"compression,required"`
	Count       float64                                `json:"count,required"`
	Countries   []string                               `json:"countries,required"`
	EncodedSize float64                                `json:"encodedSize,required"`
	IPs         []string                               `json:"ips,required"`
	Percentage  float64                                `json:"percentage,required"`
	Size        float64                                `json:"size,required"`
	Type        string                                 `json:"type,required"`
	JSON        resultGetResponseStatsResourceStatJSON `json:"-"`
}

// resultGetResponseStatsResourceStatJSON contains the JSON metadata for the struct
// [ResultGetResponseStatsResourceStat]
type resultGetResponseStatsResourceStatJSON struct {
	Compression apijson.Field
	Count       apijson.Field
	Countries   apijson.Field
	EncodedSize apijson.Field
	IPs         apijson.Field
	Percentage  apijson.Field
	Size        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseStatsResourceStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseStatsResourceStatJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseStatsServerStat struct {
	Count       float64                              `json:"count,required"`
	Countries   []string                             `json:"countries,required"`
	EncodedSize float64                              `json:"encodedSize,required"`
	IPs         []string                             `json:"ips,required"`
	Server      string                               `json:"server,required"`
	Size        float64                              `json:"size,required"`
	JSON        resultGetResponseStatsServerStatJSON `json:"-"`
}

// resultGetResponseStatsServerStatJSON contains the JSON metadata for the struct
// [ResultGetResponseStatsServerStat]
type resultGetResponseStatsServerStatJSON struct {
	Count       apijson.Field
	Countries   apijson.Field
	EncodedSize apijson.Field
	IPs         apijson.Field
	Server      apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseStatsServerStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseStatsServerStatJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseStatsTLSStat struct {
	Count         float64                                 `json:"count,required"`
	Countries     []string                                `json:"countries,required"`
	EncodedSize   float64                                 `json:"encodedSize,required"`
	IPs           []string                                `json:"ips,required"`
	Protocols     ResultGetResponseStatsTLSStatsProtocols `json:"protocols,required"`
	SecurityState string                                  `json:"securityState,required"`
	Size          float64                                 `json:"size,required"`
	JSON          resultGetResponseStatsTLSStatJSON       `json:"-"`
}

// resultGetResponseStatsTLSStatJSON contains the JSON metadata for the struct
// [ResultGetResponseStatsTLSStat]
type resultGetResponseStatsTLSStatJSON struct {
	Count         apijson.Field
	Countries     apijson.Field
	EncodedSize   apijson.Field
	IPs           apijson.Field
	Protocols     apijson.Field
	SecurityState apijson.Field
	Size          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResultGetResponseStatsTLSStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseStatsTLSStatJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseStatsTLSStatsProtocols struct {
	TLS1_3Aes128Gcm float64                                     `json:"TLS 1.3 / AES_128_GCM,required"`
	JSON            resultGetResponseStatsTLSStatsProtocolsJSON `json:"-"`
}

// resultGetResponseStatsTLSStatsProtocolsJSON contains the JSON metadata for the
// struct [ResultGetResponseStatsTLSStatsProtocols]
type resultGetResponseStatsTLSStatsProtocolsJSON struct {
	TLS1_3Aes128Gcm apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResultGetResponseStatsTLSStatsProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseStatsTLSStatsProtocolsJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseTask struct {
	ApexDomain    string                            `json:"apexDomain,required"`
	Domain        string                            `json:"domain,required"`
	DomURL        string                            `json:"domURL,required"`
	ExtraOptions  ResultGetResponseTaskExtraOptions `json:"extraOptions,required"`
	Method        string                            `json:"method,required"`
	ReportURL     string                            `json:"reportURL,required"`
	ScreenshotURL string                            `json:"screenshotURL,required"`
	Source        string                            `json:"source,required"`
	Success       bool                              `json:"success,required"`
	Time          string                            `json:"time,required"`
	URL           string                            `json:"url,required"`
	UUID          string                            `json:"uuid,required"`
	Visibility    string                            `json:"visibility,required"`
	JSON          resultGetResponseTaskJSON         `json:"-"`
}

// resultGetResponseTaskJSON contains the JSON metadata for the struct
// [ResultGetResponseTask]
type resultGetResponseTaskJSON struct {
	ApexDomain    apijson.Field
	Domain        apijson.Field
	DomURL        apijson.Field
	ExtraOptions  apijson.Field
	Method        apijson.Field
	ReportURL     apijson.Field
	ScreenshotURL apijson.Field
	Source        apijson.Field
	Success       apijson.Field
	Time          apijson.Field
	URL           apijson.Field
	UUID          apijson.Field
	Visibility    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResultGetResponseTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseTaskJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseTaskExtraOptions struct {
	ScreenshotsResolutions []string                              `json:"screenshotsResolutions"`
	JSON                   resultGetResponseTaskExtraOptionsJSON `json:"-"`
}

// resultGetResponseTaskExtraOptionsJSON contains the JSON metadata for the struct
// [ResultGetResponseTaskExtraOptions]
type resultGetResponseTaskExtraOptionsJSON struct {
	ScreenshotsResolutions apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ResultGetResponseTaskExtraOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseTaskExtraOptionsJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseVerdicts struct {
	Overall ResultGetResponseVerdictsOverall `json:"overall,required"`
	JSON    resultGetResponseVerdictsJSON    `json:"-"`
}

// resultGetResponseVerdictsJSON contains the JSON metadata for the struct
// [ResultGetResponseVerdicts]
type resultGetResponseVerdictsJSON struct {
	Overall     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseVerdicts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseVerdictsJSON) RawJSON() string {
	return r.raw
}

type ResultGetResponseVerdictsOverall struct {
	Categories  []string                             `json:"categories,required"`
	HasVerdicts bool                                 `json:"hasVerdicts,required"`
	Malicious   bool                                 `json:"malicious,required"`
	Tags        []string                             `json:"tags,required"`
	JSON        resultGetResponseVerdictsOverallJSON `json:"-"`
}

// resultGetResponseVerdictsOverallJSON contains the JSON metadata for the struct
// [ResultGetResponseVerdictsOverall]
type resultGetResponseVerdictsOverallJSON struct {
	Categories  apijson.Field
	HasVerdicts apijson.Field
	Malicious   apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultGetResponseVerdictsOverall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultGetResponseVerdictsOverallJSON) RawJSON() string {
	return r.raw
}
