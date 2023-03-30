package responses

import (
	"time"

	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type APIResponseSingleID struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool                      `json:"success"`
	Result  APIResponseSingleIDResult `json:"result,nullable"`
	JSON    APIResponseSingleIDJSON
}

type APIResponseSingleIDJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into APIResponseSingleID using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *APIResponseSingleID) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Messages struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    MessagesJSON
}

type MessagesJSON struct {
	Code    pjson.Metadata
	Message pjson.Metadata
	Raw     []byte
	Extras  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Messages using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Messages) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type APIResponseSingleIDResult struct {
	// Identifier
	ID   string `json:"id,required"`
	JSON APIResponseSingleIDResultJSON
}

type APIResponseSingleIDResultJSON struct {
	ID     pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into APIResponseSingleIDResult
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *APIResponseSingleIDResult) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SessionAffinity string

const (
	SessionAffinityNone     SessionAffinity = "none"
	SessionAffinityCookie   SessionAffinity = "cookie"
	SessionAffinityIpCookie SessionAffinity = "ip_cookie"
	SessionAffinityEmpty    SessionAffinity = "\"\""
)

type ZoneListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success    bool       `json:"success"`
	ResultInfo ResultInfo `json:"result_info"`
	Result     []Zone     `json:"result"`
	JSON       ZoneListResponseJSON
}

type ZoneListResponseJSON struct {
	Errors     pjson.Metadata
	Messages   pjson.Metadata
	Success    pjson.Metadata
	ResultInfo pjson.Metadata
	Result     pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ZoneListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ZoneListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64 `json:"total_count"`
	JSON       ResultInfoJSON
}

type ResultInfoJSON struct {
	Count      pjson.Metadata
	Page       pjson.Metadata
	PerPage    pjson.Metadata
	TotalCount pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ResultInfo using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ResultInfo) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Zone struct {
	// The last time proof of ownership was detected and the zone was made active
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// Identifier
	ID string `json:"id,required"`
	// When the zone was last modified
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name
	Name string `json:"name,required"`
	// DNS host at the time of switching to Cloudflare
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare Notes: Is this only available
	// for full zones?
	OriginalNameServers []string `json:"original_name_servers,required,nullable"`
	// Registrar for the domain at the time of switching to Cloudflare
	OriginalRegistrar string `json:"original_registrar,required,nullable"`
	JSON              ZoneJSON
}

type ZoneJSON struct {
	ActivatedOn         pjson.Metadata
	CreatedOn           pjson.Metadata
	DevelopmentMode     pjson.Metadata
	ID                  pjson.Metadata
	ModifiedOn          pjson.Metadata
	Name                pjson.Metadata
	OriginalDnshost     pjson.Metadata
	OriginalNameServers pjson.Metadata
	OriginalRegistrar   pjson.Metadata
	Raw                 []byte
	Extras              map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Zone using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Zone) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ZoneCreateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	Result  Zone `json:"result"`
	JSON    ZoneCreateResponseJSON
}

type ZoneCreateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ZoneCreateResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ZoneCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ZoneRetrieveResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	Result  Zone `json:"result"`
	JSON    ZoneRetrieveResponseJSON
}

type ZoneRetrieveResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ZoneRetrieveResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ZoneRetrieveResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ZoneUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	Result  Zone `json:"result"`
	JSON    ZoneUpdateResponseJSON
}

type ZoneUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ZoneUpdateResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ZoneUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
