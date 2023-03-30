package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/core/query"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type APIResponseSingleID struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool]                      `json:"success"`
	Result  fields.Field[APIResponseSingleIDResult] `json:"result,nullable"`
}

// MarshalJSON serializes APIResponseSingleID into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *APIResponseSingleID) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r APIResponseSingleID) String() (result string) {
	return fmt.Sprintf("&APIResponseSingleID{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type Messages struct {
	Code    fields.Field[int64]  `json:"code,required"`
	Message fields.Field[string] `json:"message,required"`
}

// MarshalJSON serializes Messages into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Messages) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Messages) String() (result string) {
	return fmt.Sprintf("&Messages{Code:%s Message:%s}", r.Code, r.Message)
}

type APIResponseSingleIDResult struct {
	// Identifier
	ID fields.Field[string] `json:"id,required"`
}

// MarshalJSON serializes APIResponseSingleIDResult into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *APIResponseSingleIDResult) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r APIResponseSingleIDResult) String() (result string) {
	return fmt.Sprintf("&APIResponseSingleIDResult{ID:%s}", r.ID)
}

type SessionAffinity string

const (
	SessionAffinityNone     SessionAffinity = "none"
	SessionAffinityCookie   SessionAffinity = "cookie"
	SessionAffinityIpCookie SessionAffinity = "ip_cookie"
	SessionAffinityEmpty    SessionAffinity = "\"\""
)

type ZoneListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success    fields.Field[bool]       `json:"success"`
	ResultInfo fields.Field[ResultInfo] `json:"result_info"`
	Result     fields.Field[[]Zone]     `json:"result"`
}

// MarshalJSON serializes ZoneListResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ZoneListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ZoneListResponse) String() (result string) {
	return fmt.Sprintf("&ZoneListResponse{Errors:%s Messages:%s Success:%s ResultInfo:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.ResultInfo, core.Fmt(r.Result))
}

type ResultInfo struct {
	// Total number of results for the requested service
	Count fields.Field[float64] `json:"count"`
	// Current page within paginated list of results
	Page fields.Field[float64] `json:"page"`
	// Number of results per page of results
	PerPage fields.Field[float64] `json:"per_page"`
	// Total results available without any search parameters
	TotalCount fields.Field[float64] `json:"total_count"`
}

// MarshalJSON serializes ResultInfo into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ResultInfo) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ResultInfo) String() (result string) {
	return fmt.Sprintf("&ResultInfo{Count:%s Page:%s PerPage:%s TotalCount:%s}", r.Count, r.Page, r.PerPage, r.TotalCount)
}

type Zone struct {
	// The last time proof of ownership was detected and the zone was made active
	ActivatedOn fields.Field[time.Time] `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created
	CreatedOn fields.Field[time.Time] `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode fields.Field[float64] `json:"development_mode,required"`
	// Identifier
	ID fields.Field[string] `json:"id,required"`
	// When the zone was last modified
	ModifiedOn fields.Field[time.Time] `json:"modified_on,required" format:"date-time"`
	// The domain name
	Name fields.Field[string] `json:"name,required"`
	// DNS host at the time of switching to Cloudflare
	OriginalDnshost fields.Field[string] `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare Notes: Is this only available
	// for full zones?
	OriginalNameServers fields.Field[[]string] `json:"original_name_servers,required,nullable"`
	// Registrar for the domain at the time of switching to Cloudflare
	OriginalRegistrar fields.Field[string] `json:"original_registrar,required,nullable"`
}

// MarshalJSON serializes Zone into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Zone) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Zone) String() (result string) {
	return fmt.Sprintf("&Zone{ActivatedOn:%s CreatedOn:%s DevelopmentMode:%s ID:%s ModifiedOn:%s Name:%s OriginalDnshost:%s OriginalNameServers:%s OriginalRegistrar:%s}", r.ActivatedOn, r.CreatedOn, r.DevelopmentMode, r.ID, r.ModifiedOn, r.Name, r.OriginalDnshost, core.Fmt(r.OriginalNameServers), r.OriginalRegistrar)
}

type ZoneCreateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	Result  fields.Field[Zone] `json:"result"`
}

// MarshalJSON serializes ZoneCreateResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ZoneCreateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ZoneCreateResponse) String() (result string) {
	return fmt.Sprintf("&ZoneCreateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type ZoneRetrieveResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	Result  fields.Field[Zone] `json:"result"`
}

// MarshalJSON serializes ZoneRetrieveResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ZoneRetrieveResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ZoneRetrieveResponse) String() (result string) {
	return fmt.Sprintf("&ZoneRetrieveResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type ZoneUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	Result  fields.Field[Zone] `json:"result"`
}

// MarshalJSON serializes ZoneUpdateResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ZoneUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ZoneUpdateResponse) String() (result string) {
	return fmt.Sprintf("&ZoneUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type ZoneNewParams struct {
	Account fields.Field[ZoneNewParamsAccount] `json:"account,required"`
	// The domain name
	Name fields.Field[string] `json:"name,required"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup.
	Type fields.Field[TypeFo5LnpAt] `json:"type"`
}

// MarshalJSON serializes ZoneNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ZoneNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ZoneNewParams) String() (result string) {
	return fmt.Sprintf("&ZoneNewParams{Account:%s Name:%s Type:%s}", r.Account, r.Name, r.Type)
}

type ZoneNewParamsAccount struct {
	// Identifier
	ID fields.Field[string] `json:"id"`
}

func (r ZoneNewParamsAccount) String() (result string) {
	return fmt.Sprintf("&ZoneNewParamsAccount{ID:%s}", r.ID)
}

type TypeFo5LnpAt string

const (
	TypeFo5LnpAtFull    TypeFo5LnpAt = "full"
	TypeFo5LnpAtPartial TypeFo5LnpAt = "partial"
)

type ZoneUpdateParams struct {
	// Indicates whether the zone is only using Cloudflare DNS services. A true value
	// means the zone will not receive security or performance benefits.
	Paused fields.Field[bool] `json:"paused"`
	// (Deprecated) Please use the `/zones/{identifier}/subscription` API to update a
	// zone's plan. Changing this value will create/cancel associated subscriptions. To
	// view available plans for this zone, see Zone Plans.
	Plan fields.Field[ZoneUpdateParamsPlan] `json:"plan"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup. This parameter is only
	// available to Enterprise customers or if it has been explicitly enabled on a
	// zone.
	Type fields.Field[ZoneUpdateParamsType] `json:"type"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers fields.Field[[]string] `json:"vanity_name_servers"`
}

// MarshalJSON serializes ZoneUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ZoneUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ZoneUpdateParams) String() (result string) {
	return fmt.Sprintf("&ZoneUpdateParams{Paused:%s Plan:%s Type:%s VanityNameServers:%s}", r.Paused, r.Plan, r.Type, core.Fmt(r.VanityNameServers))
}

type ZoneUpdateParamsPlan struct {
	// Identifier
	ID fields.Field[string] `json:"id"`
}

func (r ZoneUpdateParamsPlan) String() (result string) {
	return fmt.Sprintf("&ZoneUpdateParamsPlan{ID:%s}", r.ID)
}

type ZoneUpdateParamsType string

const (
	ZoneUpdateParamsTypeFull    ZoneUpdateParamsType = "full"
	ZoneUpdateParamsTypePartial ZoneUpdateParamsType = "partial"
)

type ZoneListParams struct {
	// A domain name.
	Name   fields.Field[string] `query:"name" format:"hostname"`
	Status fields.Field[string] `query:"status"`
	// Page number of paginated results.
	Page fields.Field[float64] `query:"page"`
	// Number of zones per page.
	PerPage fields.Field[float64] `query:"per_page"`
	// Field to order zones by.
	Order fields.Field[ZoneListParamsOrder] `query:"order"`
	// Direction to order zones.
	Direction fields.Field[ZoneListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any).
	Match   fields.Field[ZoneListParamsMatch]   `query:"match"`
	Account fields.Field[ZoneListParamsAccount] `query:"account"`
}

// URLQuery serializes ZoneListParams into a url.Values of the query parameters
// associated with this value
func (r *ZoneListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r ZoneListParams) String() (result string) {
	return fmt.Sprintf("&ZoneListParams{Name:%s Status:%s Page:%s PerPage:%s Order:%s Direction:%s Match:%s Account:%s}", r.Name, r.Status, r.Page, r.PerPage, r.Order, r.Direction, r.Match, r.Account)
}

type ZoneListParamsOrder string

const (
	ZoneListParamsOrderName        ZoneListParamsOrder = "name"
	ZoneListParamsOrderStatus      ZoneListParamsOrder = "status"
	ZoneListParamsOrderAccountID   ZoneListParamsOrder = "account.id"
	ZoneListParamsOrderAccountName ZoneListParamsOrder = "account.name"
)

type ZoneListParamsDirection string

const (
	ZoneListParamsDirectionAsc  ZoneListParamsDirection = "asc"
	ZoneListParamsDirectionDesc ZoneListParamsDirection = "desc"
)

type ZoneListParamsMatch string

const (
	ZoneListParamsMatchAny ZoneListParamsMatch = "any"
	ZoneListParamsMatchAll ZoneListParamsMatch = "all"
)

type ZoneListParamsAccount struct {
	ID   fields.Field[string] `query:"id"`
	Name fields.Field[string] `query:"name"`
}

// URLQuery serializes ZoneListParamsAccount into a url.Values of the query
// parameters associated with this value
func (r *ZoneListParamsAccount) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r ZoneListParamsAccount) String() (result string) {
	return fmt.Sprintf("&ZoneListParamsAccount{ID:%s Name:%s}", r.ID, r.Name)
}
