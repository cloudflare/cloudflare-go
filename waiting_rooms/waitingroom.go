// File generated from our OpenAPI spec by Stainless.

package waiting_rooms

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

// WaitingRoomService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWaitingRoomService] method
// instead.
type WaitingRoomService struct {
	Options  []option.RequestOption
	Page     *PageService
	Events   *EventService
	Rules    *RuleService
	Statuses *StatusService
	Settings *SettingService
}

// NewWaitingRoomService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWaitingRoomService(opts ...option.RequestOption) (r *WaitingRoomService) {
	r = &WaitingRoomService{}
	r.Options = opts
	r.Page = NewPageService(opts...)
	r.Events = NewEventService(opts...)
	r.Rules = NewRuleService(opts...)
	r.Statuses = NewStatusService(opts...)
	r.Settings = NewSettingService(opts...)
	return
}

// Creates a new waiting room.
func (r *WaitingRoomService) New(ctx context.Context, zoneIdentifier string, body WaitingRoomNewParams, opts ...option.RequestOption) (res *WaitingroomWaitingroom, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured waiting room.
func (r *WaitingRoomService) Update(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body WaitingRoomUpdateParams, opts ...option.RequestOption) (res *WaitingroomWaitingroom, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists waiting rooms.
func (r *WaitingRoomService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]WaitingroomWaitingroom, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomListResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a waiting room.
func (r *WaitingRoomService) Delete(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, opts ...option.RequestOption) (res *WaitingRoomDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patches a configured waiting room.
func (r *WaitingRoomService) Edit(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body WaitingRoomEditParams, opts ...option.RequestOption) (res *WaitingroomWaitingroom, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single configured waiting room.
func (r *WaitingRoomService) Get(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, opts ...option.RequestOption) (res *WaitingroomWaitingroom, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WaitingroomWaitingroom struct {
	ID interface{} `json:"id"`
	// Only available for the Waiting Room Advanced subscription. Additional hostname
	// and path combinations to which this waiting room will be applied. There is an
	// implied wildcard at the end of the path. The hostname and path combination must
	// be unique to this and all other waiting rooms.
	AdditionalRoutes []WaitingroomWaitingroomAdditionalRoute `json:"additional_routes"`
	// Configures cookie attributes for the waiting room cookie. This encrypted cookie
	// stores a user's status in the waiting room, such as queue position.
	CookieAttributes WaitingroomWaitingroomCookieAttributes `json:"cookie_attributes"`
	// Appends a '\_' + a custom suffix to the end of Cloudflare Waiting Room's cookie
	// name(**cf_waitingroom). If `cookie_suffix` is "abcd", the cookie name will be
	// `**cf_waitingroom_abcd`. This field is required if using `additional_routes`.
	CookieSuffix string    `json:"cookie_suffix"`
	CreatedOn    time.Time `json:"created_on" format:"date-time"`
	// Only available for the Waiting Room Advanced subscription. This is a template
	// html file that will be rendered at the edge. If no custom_page_html is provided,
	// the default waiting room will be used. The template is based on mustache (
	// https://mustache.github.io/ ). There are several variables that are evaluated by
	// the Cloudflare edge:
	//
	//  1. {{`waitTimeKnown`}} Acts like a boolean value that indicates the behavior to
	//     take when wait time is not available, for instance when queue_all is
	//     **true**.
	//  2. {{`waitTimeFormatted`}} Estimated wait time for the user. For example, five
	//     minutes. Alternatively, you can use:
	//  3. {{`waitTime`}} Number of minutes of estimated wait for a user.
	//  4. {{`waitTimeHours`}} Number of hours of estimated wait for a user
	//     (`Math.floor(waitTime/60)`).
	//  5. {{`waitTimeHourMinutes`}} Number of minutes above the `waitTimeHours` value
	//     (`waitTime%60`).
	//  6. {{`queueIsFull`}} Changes to **true** when no more people can be added to the
	//     queue.
	//
	// To view the full list of variables, look at the `cfWaitingRoom` object described
	// under the `json_response_enabled` property in other Waiting Room API calls.
	CustomPageHTML string `json:"custom_page_html"`
	// The language of the default page template. If no default_template_language is
	// provided, then `en-US` (English) will be used.
	DefaultTemplateLanguage WaitingroomWaitingroomDefaultTemplateLanguage `json:"default_template_language"`
	// A note that you can use to add more details about the waiting room.
	Description string `json:"description"`
	// Only available for the Waiting Room Advanced subscription. Disables automatic
	// renewal of session cookies. If `true`, an accepted user will have
	// session_duration minutes to browse the site. After that, they will have to go
	// through the waiting room again. If `false`, a user's session cookie will be
	// automatically renewed on every request.
	DisableSessionRenewal bool `json:"disable_session_renewal"`
	// The host name to which the waiting room will be applied (no wildcards). Please
	// do not include the scheme (http:// or https://). The host and path combination
	// must be unique.
	Host string `json:"host"`
	// Only available for the Waiting Room Advanced subscription. If `true`, requests
	// to the waiting room with the header `Accept: application/json` will receive a
	// JSON response object with information on the user's status in the waiting room
	// as opposed to the configured static HTML page. This JSON response object has one
	// property `cfWaitingRoom` which is an object containing the following fields:
	//
	//  1. `inWaitingRoom`: Boolean indicating if the user is in the waiting room
	//     (always **true**).
	//  2. `waitTimeKnown`: Boolean indicating if the current estimated wait times are
	//     accurate. If **false**, they are not available.
	//  3. `waitTime`: Valid only when `waitTimeKnown` is **true**. Integer indicating
	//     the current estimated time in minutes the user will wait in the waiting room.
	//     When `queueingMethod` is **random**, this is set to `waitTime50Percentile`.
	//  4. `waitTime25Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 25% of users that gain entry the fastest (25th percentile).
	//  5. `waitTime50Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 50% of users that gain entry the fastest (50th percentile).
	//     In other words, half of the queued users are expected to let into the origin
	//     website before `waitTime50Percentile` and half are expected to be let in
	//     after it.
	//  6. `waitTime75Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 75% of users that gain entry the fastest (75th percentile).
	//  7. `waitTimeFormatted`: String displaying the `waitTime` formatted in English
	//     for users. If `waitTimeKnown` is **false**, `waitTimeFormatted` will display
	//     **unavailable**.
	//  8. `queueIsFull`: Boolean indicating if the waiting room's queue is currently
	//     full and not accepting new users at the moment.
	//  9. `queueAll`: Boolean indicating if all users will be queued in the waiting
	//     room and no one will be let into the origin website.
	//  10. `lastUpdated`: String displaying the timestamp as an ISO 8601 string of the
	//     user's last attempt to leave the waiting room and be let into the origin
	//     website. The user is able to make another attempt after
	//     `refreshIntervalSeconds` past this time. If the user makes a request too
	//     soon, it will be ignored and `lastUpdated` will not change.
	//  11. `refreshIntervalSeconds`: Integer indicating the number of seconds after
	//     `lastUpdated` until the user is able to make another attempt to leave the
	//     waiting room and be let into the origin website. When the `queueingMethod`
	//     is `reject`, there is no specified refresh time — it will always be
	//     **zero**.
	//  12. `queueingMethod`: The queueing method currently used by the waiting room. It
	//     is either **fifo**, **random**, **passthrough**, or **reject**.
	//  13. `isFIFOQueue`: Boolean indicating if the waiting room uses a FIFO
	//     (First-In-First-Out) queue.
	//  14. `isRandomQueue`: Boolean indicating if the waiting room uses a Random queue
	//     where users gain access randomly.
	//  15. `isPassthroughQueue`: Boolean indicating if the waiting room uses a
	//     passthrough queue. Keep in mind that when passthrough is enabled, this JSON
	//     response will only exist when `queueAll` is **true** or `isEventPrequeueing`
	//     is **true** because in all other cases requests will go directly to the
	//     origin.
	//  16. `isRejectQueue`: Boolean indicating if the waiting room uses a reject queue.
	//  17. `isEventActive`: Boolean indicating if an event is currently occurring.
	//     Events are able to change a waiting room's behavior during a specified
	//     period of time. For additional information, look at the event properties
	//     `prequeue_start_time`, `event_start_time`, and `event_end_time` in the
	//     documentation for creating waiting room events. Events are considered active
	//     between these start and end times, as well as during the prequeueing period
	//     if it exists.
	//  18. `isEventPrequeueing`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if an event is currently prequeueing users before it starts.
	//  19. `timeUntilEventStart`: Valid only when `isEventPrequeueing` is **true**.
	//     Integer indicating the number of minutes until the event starts.
	//  20. `timeUntilEventStartFormatted`: String displaying the `timeUntilEventStart`
	//     formatted in English for users. If `isEventPrequeueing` is **false**,
	//     `timeUntilEventStartFormatted` will display **unavailable**.
	//  21. `timeUntilEventEnd`: Valid only when `isEventActive` is **true**. Integer
	//     indicating the number of minutes until the event ends.
	//  22. `timeUntilEventEndFormatted`: String displaying the `timeUntilEventEnd`
	//     formatted in English for users. If `isEventActive` is **false**,
	//     `timeUntilEventEndFormatted` will display **unavailable**.
	//  23. `shuffleAtEventStart`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if the users in the prequeue are shuffled randomly when the event
	//     starts.
	//
	// An example cURL to a waiting room could be:
	//
	//	curl -X GET "https://example.com/waitingroom" \
	//		-H "Accept: application/json"
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **fifo** and no event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 0,
	//			"waitTime50Percentile": 0,
	//			"waitTime75Percentile": 0,
	//			"waitTimeFormatted": "10 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "fifo",
	//			"isFIFOQueue": true,
	//			"isRandomQueue": false,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": false,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 0,
	//			"timeUntilEventEndFormatted": "unavailable",
	//			"shuffleAtEventStart": false
	//		}
	//	}
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **random** and an event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 5,
	//			"waitTime50Percentile": 10,
	//			"waitTime75Percentile": 15,
	//			"waitTimeFormatted": "5 minutes to 15 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "random",
	//			"isFIFOQueue": false,
	//			"isRandomQueue": true,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": true,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 15,
	//			"timeUntilEventEndFormatted": "15 minutes",
	//			"shuffleAtEventStart": true
	//		}
	//	}.
	JsonResponseEnabled bool      `json:"json_response_enabled"`
	ModifiedOn          time.Time `json:"modified_on" format:"date-time"`
	// A unique name to identify the waiting room. Only alphanumeric characters,
	// hyphens and underscores are allowed.
	Name string `json:"name"`
	// Sets the number of new users that will be let into the route every minute. This
	// value is used as baseline for the number of users that are let in per minute. So
	// it is possible that there is a little more or little less traffic coming to the
	// route based on the traffic patterns at that time around the world.
	NewUsersPerMinute int64 `json:"new_users_per_minute"`
	// An ISO 8601 timestamp that marks when the next event will begin queueing.
	NextEventPrequeueStartTime string `json:"next_event_prequeue_start_time,nullable"`
	// An ISO 8601 timestamp that marks when the next event will start.
	NextEventStartTime string `json:"next_event_start_time,nullable"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path string `json:"path"`
	// If queue_all is `true`, all the traffic that is coming to a route will be sent
	// to the waiting room. No new traffic can get to the route once this field is set
	// and estimated time will become unavailable.
	QueueAll bool `json:"queue_all"`
	// Sets the queueing method used by the waiting room. Changing this parameter from
	// the **default** queueing method is only available for the Waiting Room Advanced
	// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
	// event is prequeueing, users in the waiting room will not be accepted to the
	// origin. These users will always see a waiting room page that refreshes
	// automatically. The valid queueing methods are:
	//
	//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
	//     the order they arrived.
	//  2. `random`: Random queue where customers gain access randomly, regardless of
	//     arrival time.
	//  3. `passthrough`: Users will pass directly through the waiting room and into the
	//     origin website. As a result, any configured limits will not be respected
	//     while this is enabled. This method can be used as an alternative to disabling
	//     a waiting room (with `suspended`) so that analytics are still reported. This
	//     can be used if you wish to allow all traffic normally, but want to restrict
	//     traffic during a waiting room event, or vice versa.
	//  4. `reject`: Users will be immediately rejected from the waiting room. As a
	//     result, no users will reach the origin website while this is enabled. This
	//     can be used if you wish to reject all traffic while performing maintenance,
	//     block traffic during a specified period of time (an event), or block traffic
	//     while events are not occurring. Consider a waiting room used for vaccine
	//     distribution that only allows traffic during sign-up events, and otherwise
	//     blocks all traffic. For this case, the waiting room uses `reject`, and its
	//     events override this with `fifo`, `random`, or `passthrough`. When this
	//     queueing method is enabled and neither `queueAll` is enabled nor an event is
	//     prequeueing, the waiting room page **will not refresh automatically**.
	QueueingMethod WaitingroomWaitingroomQueueingMethod `json:"queueing_method"`
	// HTTP status code returned to a user while in the queue.
	QueueingStatusCode WaitingroomWaitingroomQueueingStatusCode `json:"queueing_status_code"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to
	// the route. If a user is not seen by Cloudflare again in that time period, they
	// will be treated as a new user that visits the route.
	SessionDuration int64 `json:"session_duration"`
	// Suspends or allows traffic going to the waiting room. If set to `true`, the
	// traffic will not go to the waiting room.
	Suspended bool `json:"suspended"`
	// Sets the total number of active user sessions on the route at a point in time. A
	// route is a combination of host and path on which a waiting room is available.
	// This value is used as a baseline for the total number of active user sessions on
	// the route. It is possible to have a situation where there are more or less
	// active users sessions on the route based on the traffic patterns at that time
	// around the world.
	TotalActiveUsers int64                      `json:"total_active_users"`
	JSON             waitingroomWaitingroomJSON `json:"-"`
}

// waitingroomWaitingroomJSON contains the JSON metadata for the struct
// [WaitingroomWaitingroom]
type waitingroomWaitingroomJSON struct {
	ID                         apijson.Field
	AdditionalRoutes           apijson.Field
	CookieAttributes           apijson.Field
	CookieSuffix               apijson.Field
	CreatedOn                  apijson.Field
	CustomPageHTML             apijson.Field
	DefaultTemplateLanguage    apijson.Field
	Description                apijson.Field
	DisableSessionRenewal      apijson.Field
	Host                       apijson.Field
	JsonResponseEnabled        apijson.Field
	ModifiedOn                 apijson.Field
	Name                       apijson.Field
	NewUsersPerMinute          apijson.Field
	NextEventPrequeueStartTime apijson.Field
	NextEventStartTime         apijson.Field
	Path                       apijson.Field
	QueueAll                   apijson.Field
	QueueingMethod             apijson.Field
	QueueingStatusCode         apijson.Field
	SessionDuration            apijson.Field
	Suspended                  apijson.Field
	TotalActiveUsers           apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *WaitingroomWaitingroom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingroomWaitingroomJSON) RawJSON() string {
	return r.raw
}

type WaitingroomWaitingroomAdditionalRoute struct {
	// The hostname to which this waiting room will be applied (no wildcards). The
	// hostname must be the primary domain, subdomain, or custom hostname (if using SSL
	// for SaaS) of this zone. Please do not include the scheme (http:// or https://).
	Host string `json:"host"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path string                                    `json:"path"`
	JSON waitingroomWaitingroomAdditionalRouteJSON `json:"-"`
}

// waitingroomWaitingroomAdditionalRouteJSON contains the JSON metadata for the
// struct [WaitingroomWaitingroomAdditionalRoute]
type waitingroomWaitingroomAdditionalRouteJSON struct {
	Host        apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingroomWaitingroomAdditionalRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingroomWaitingroomAdditionalRouteJSON) RawJSON() string {
	return r.raw
}

// Configures cookie attributes for the waiting room cookie. This encrypted cookie
// stores a user's status in the waiting room, such as queue position.
type WaitingroomWaitingroomCookieAttributes struct {
	// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
	// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
	// Note that when using value `none`, the secure attribute cannot be set to
	// `never`.
	Samesite WaitingroomWaitingroomCookieAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on the waiting room cookie. Value `always`
	// indicates that the Secure attribute will be set in the Set-Cookie header,
	// `never` indicates that the Secure attribute will not be set, and `auto` will set
	// the Secure attribute depending if **Always Use HTTPS** is enabled.
	Secure WaitingroomWaitingroomCookieAttributesSecure `json:"secure"`
	JSON   waitingroomWaitingroomCookieAttributesJSON   `json:"-"`
}

// waitingroomWaitingroomCookieAttributesJSON contains the JSON metadata for the
// struct [WaitingroomWaitingroomCookieAttributes]
type waitingroomWaitingroomCookieAttributesJSON struct {
	Samesite    apijson.Field
	Secure      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingroomWaitingroomCookieAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingroomWaitingroomCookieAttributesJSON) RawJSON() string {
	return r.raw
}

// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
// Note that when using value `none`, the secure attribute cannot be set to
// `never`.
type WaitingroomWaitingroomCookieAttributesSamesite string

const (
	WaitingroomWaitingroomCookieAttributesSamesiteAuto   WaitingroomWaitingroomCookieAttributesSamesite = "auto"
	WaitingroomWaitingroomCookieAttributesSamesiteLax    WaitingroomWaitingroomCookieAttributesSamesite = "lax"
	WaitingroomWaitingroomCookieAttributesSamesiteNone   WaitingroomWaitingroomCookieAttributesSamesite = "none"
	WaitingroomWaitingroomCookieAttributesSamesiteStrict WaitingroomWaitingroomCookieAttributesSamesite = "strict"
)

// Configures the Secure attribute on the waiting room cookie. Value `always`
// indicates that the Secure attribute will be set in the Set-Cookie header,
// `never` indicates that the Secure attribute will not be set, and `auto` will set
// the Secure attribute depending if **Always Use HTTPS** is enabled.
type WaitingroomWaitingroomCookieAttributesSecure string

const (
	WaitingroomWaitingroomCookieAttributesSecureAuto   WaitingroomWaitingroomCookieAttributesSecure = "auto"
	WaitingroomWaitingroomCookieAttributesSecureAlways WaitingroomWaitingroomCookieAttributesSecure = "always"
	WaitingroomWaitingroomCookieAttributesSecureNever  WaitingroomWaitingroomCookieAttributesSecure = "never"
)

// The language of the default page template. If no default_template_language is
// provided, then `en-US` (English) will be used.
type WaitingroomWaitingroomDefaultTemplateLanguage string

const (
	WaitingroomWaitingroomDefaultTemplateLanguageEnUs WaitingroomWaitingroomDefaultTemplateLanguage = "en-US"
	WaitingroomWaitingroomDefaultTemplateLanguageEsEs WaitingroomWaitingroomDefaultTemplateLanguage = "es-ES"
	WaitingroomWaitingroomDefaultTemplateLanguageDeDe WaitingroomWaitingroomDefaultTemplateLanguage = "de-DE"
	WaitingroomWaitingroomDefaultTemplateLanguageFrFr WaitingroomWaitingroomDefaultTemplateLanguage = "fr-FR"
	WaitingroomWaitingroomDefaultTemplateLanguageItIt WaitingroomWaitingroomDefaultTemplateLanguage = "it-IT"
	WaitingroomWaitingroomDefaultTemplateLanguageJaJp WaitingroomWaitingroomDefaultTemplateLanguage = "ja-JP"
	WaitingroomWaitingroomDefaultTemplateLanguageKoKr WaitingroomWaitingroomDefaultTemplateLanguage = "ko-KR"
	WaitingroomWaitingroomDefaultTemplateLanguagePtBr WaitingroomWaitingroomDefaultTemplateLanguage = "pt-BR"
	WaitingroomWaitingroomDefaultTemplateLanguageZhCn WaitingroomWaitingroomDefaultTemplateLanguage = "zh-CN"
	WaitingroomWaitingroomDefaultTemplateLanguageZhTw WaitingroomWaitingroomDefaultTemplateLanguage = "zh-TW"
	WaitingroomWaitingroomDefaultTemplateLanguageNlNl WaitingroomWaitingroomDefaultTemplateLanguage = "nl-NL"
	WaitingroomWaitingroomDefaultTemplateLanguagePlPl WaitingroomWaitingroomDefaultTemplateLanguage = "pl-PL"
	WaitingroomWaitingroomDefaultTemplateLanguageIDID WaitingroomWaitingroomDefaultTemplateLanguage = "id-ID"
	WaitingroomWaitingroomDefaultTemplateLanguageTrTr WaitingroomWaitingroomDefaultTemplateLanguage = "tr-TR"
	WaitingroomWaitingroomDefaultTemplateLanguageArEg WaitingroomWaitingroomDefaultTemplateLanguage = "ar-EG"
	WaitingroomWaitingroomDefaultTemplateLanguageRuRu WaitingroomWaitingroomDefaultTemplateLanguage = "ru-RU"
	WaitingroomWaitingroomDefaultTemplateLanguageFaIr WaitingroomWaitingroomDefaultTemplateLanguage = "fa-IR"
)

// Sets the queueing method used by the waiting room. Changing this parameter from
// the **default** queueing method is only available for the Waiting Room Advanced
// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
// event is prequeueing, users in the waiting room will not be accepted to the
// origin. These users will always see a waiting room page that refreshes
// automatically. The valid queueing methods are:
//
//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
//     the order they arrived.
//  2. `random`: Random queue where customers gain access randomly, regardless of
//     arrival time.
//  3. `passthrough`: Users will pass directly through the waiting room and into the
//     origin website. As a result, any configured limits will not be respected
//     while this is enabled. This method can be used as an alternative to disabling
//     a waiting room (with `suspended`) so that analytics are still reported. This
//     can be used if you wish to allow all traffic normally, but want to restrict
//     traffic during a waiting room event, or vice versa.
//  4. `reject`: Users will be immediately rejected from the waiting room. As a
//     result, no users will reach the origin website while this is enabled. This
//     can be used if you wish to reject all traffic while performing maintenance,
//     block traffic during a specified period of time (an event), or block traffic
//     while events are not occurring. Consider a waiting room used for vaccine
//     distribution that only allows traffic during sign-up events, and otherwise
//     blocks all traffic. For this case, the waiting room uses `reject`, and its
//     events override this with `fifo`, `random`, or `passthrough`. When this
//     queueing method is enabled and neither `queueAll` is enabled nor an event is
//     prequeueing, the waiting room page **will not refresh automatically**.
type WaitingroomWaitingroomQueueingMethod string

const (
	WaitingroomWaitingroomQueueingMethodFifo        WaitingroomWaitingroomQueueingMethod = "fifo"
	WaitingroomWaitingroomQueueingMethodRandom      WaitingroomWaitingroomQueueingMethod = "random"
	WaitingroomWaitingroomQueueingMethodPassthrough WaitingroomWaitingroomQueueingMethod = "passthrough"
	WaitingroomWaitingroomQueueingMethodReject      WaitingroomWaitingroomQueueingMethod = "reject"
)

// HTTP status code returned to a user while in the queue.
type WaitingroomWaitingroomQueueingStatusCode int64

const (
	WaitingroomWaitingroomQueueingStatusCode200 WaitingroomWaitingroomQueueingStatusCode = 200
	WaitingroomWaitingroomQueueingStatusCode202 WaitingroomWaitingroomQueueingStatusCode = 202
	WaitingroomWaitingroomQueueingStatusCode429 WaitingroomWaitingroomQueueingStatusCode = 429
)

type WaitingRoomDeleteResponse struct {
	ID   interface{}                   `json:"id"`
	JSON waitingRoomDeleteResponseJSON `json:"-"`
}

// waitingRoomDeleteResponseJSON contains the JSON metadata for the struct
// [WaitingRoomDeleteResponse]
type waitingRoomDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomNewParams struct {
	// The host name to which the waiting room will be applied (no wildcards). Please
	// do not include the scheme (http:// or https://). The host and path combination
	// must be unique.
	Host param.Field[string] `json:"host,required"`
	// A unique name to identify the waiting room. Only alphanumeric characters,
	// hyphens and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// Sets the number of new users that will be let into the route every minute. This
	// value is used as baseline for the number of users that are let in per minute. So
	// it is possible that there is a little more or little less traffic coming to the
	// route based on the traffic patterns at that time around the world.
	NewUsersPerMinute param.Field[int64] `json:"new_users_per_minute,required"`
	// Sets the total number of active user sessions on the route at a point in time. A
	// route is a combination of host and path on which a waiting room is available.
	// This value is used as a baseline for the total number of active user sessions on
	// the route. It is possible to have a situation where there are more or less
	// active users sessions on the route based on the traffic patterns at that time
	// around the world.
	TotalActiveUsers param.Field[int64] `json:"total_active_users,required"`
	// Only available for the Waiting Room Advanced subscription. Additional hostname
	// and path combinations to which this waiting room will be applied. There is an
	// implied wildcard at the end of the path. The hostname and path combination must
	// be unique to this and all other waiting rooms.
	AdditionalRoutes param.Field[[]WaitingRoomNewParamsAdditionalRoute] `json:"additional_routes"`
	// Configures cookie attributes for the waiting room cookie. This encrypted cookie
	// stores a user's status in the waiting room, such as queue position.
	CookieAttributes param.Field[WaitingRoomNewParamsCookieAttributes] `json:"cookie_attributes"`
	// Appends a '\_' + a custom suffix to the end of Cloudflare Waiting Room's cookie
	// name(**cf_waitingroom). If `cookie_suffix` is "abcd", the cookie name will be
	// `**cf_waitingroom_abcd`. This field is required if using `additional_routes`.
	CookieSuffix param.Field[string] `json:"cookie_suffix"`
	// Only available for the Waiting Room Advanced subscription. This is a template
	// html file that will be rendered at the edge. If no custom_page_html is provided,
	// the default waiting room will be used. The template is based on mustache (
	// https://mustache.github.io/ ). There are several variables that are evaluated by
	// the Cloudflare edge:
	//
	//  1. {{`waitTimeKnown`}} Acts like a boolean value that indicates the behavior to
	//     take when wait time is not available, for instance when queue_all is
	//     **true**.
	//  2. {{`waitTimeFormatted`}} Estimated wait time for the user. For example, five
	//     minutes. Alternatively, you can use:
	//  3. {{`waitTime`}} Number of minutes of estimated wait for a user.
	//  4. {{`waitTimeHours`}} Number of hours of estimated wait for a user
	//     (`Math.floor(waitTime/60)`).
	//  5. {{`waitTimeHourMinutes`}} Number of minutes above the `waitTimeHours` value
	//     (`waitTime%60`).
	//  6. {{`queueIsFull`}} Changes to **true** when no more people can be added to the
	//     queue.
	//
	// To view the full list of variables, look at the `cfWaitingRoom` object described
	// under the `json_response_enabled` property in other Waiting Room API calls.
	CustomPageHTML param.Field[string] `json:"custom_page_html"`
	// The language of the default page template. If no default_template_language is
	// provided, then `en-US` (English) will be used.
	DefaultTemplateLanguage param.Field[WaitingRoomNewParamsDefaultTemplateLanguage] `json:"default_template_language"`
	// A note that you can use to add more details about the waiting room.
	Description param.Field[string] `json:"description"`
	// Only available for the Waiting Room Advanced subscription. Disables automatic
	// renewal of session cookies. If `true`, an accepted user will have
	// session_duration minutes to browse the site. After that, they will have to go
	// through the waiting room again. If `false`, a user's session cookie will be
	// automatically renewed on every request.
	DisableSessionRenewal param.Field[bool] `json:"disable_session_renewal"`
	// Only available for the Waiting Room Advanced subscription. If `true`, requests
	// to the waiting room with the header `Accept: application/json` will receive a
	// JSON response object with information on the user's status in the waiting room
	// as opposed to the configured static HTML page. This JSON response object has one
	// property `cfWaitingRoom` which is an object containing the following fields:
	//
	//  1. `inWaitingRoom`: Boolean indicating if the user is in the waiting room
	//     (always **true**).
	//  2. `waitTimeKnown`: Boolean indicating if the current estimated wait times are
	//     accurate. If **false**, they are not available.
	//  3. `waitTime`: Valid only when `waitTimeKnown` is **true**. Integer indicating
	//     the current estimated time in minutes the user will wait in the waiting room.
	//     When `queueingMethod` is **random**, this is set to `waitTime50Percentile`.
	//  4. `waitTime25Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 25% of users that gain entry the fastest (25th percentile).
	//  5. `waitTime50Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 50% of users that gain entry the fastest (50th percentile).
	//     In other words, half of the queued users are expected to let into the origin
	//     website before `waitTime50Percentile` and half are expected to be let in
	//     after it.
	//  6. `waitTime75Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 75% of users that gain entry the fastest (75th percentile).
	//  7. `waitTimeFormatted`: String displaying the `waitTime` formatted in English
	//     for users. If `waitTimeKnown` is **false**, `waitTimeFormatted` will display
	//     **unavailable**.
	//  8. `queueIsFull`: Boolean indicating if the waiting room's queue is currently
	//     full and not accepting new users at the moment.
	//  9. `queueAll`: Boolean indicating if all users will be queued in the waiting
	//     room and no one will be let into the origin website.
	//  10. `lastUpdated`: String displaying the timestamp as an ISO 8601 string of the
	//     user's last attempt to leave the waiting room and be let into the origin
	//     website. The user is able to make another attempt after
	//     `refreshIntervalSeconds` past this time. If the user makes a request too
	//     soon, it will be ignored and `lastUpdated` will not change.
	//  11. `refreshIntervalSeconds`: Integer indicating the number of seconds after
	//     `lastUpdated` until the user is able to make another attempt to leave the
	//     waiting room and be let into the origin website. When the `queueingMethod`
	//     is `reject`, there is no specified refresh time — it will always be
	//     **zero**.
	//  12. `queueingMethod`: The queueing method currently used by the waiting room. It
	//     is either **fifo**, **random**, **passthrough**, or **reject**.
	//  13. `isFIFOQueue`: Boolean indicating if the waiting room uses a FIFO
	//     (First-In-First-Out) queue.
	//  14. `isRandomQueue`: Boolean indicating if the waiting room uses a Random queue
	//     where users gain access randomly.
	//  15. `isPassthroughQueue`: Boolean indicating if the waiting room uses a
	//     passthrough queue. Keep in mind that when passthrough is enabled, this JSON
	//     response will only exist when `queueAll` is **true** or `isEventPrequeueing`
	//     is **true** because in all other cases requests will go directly to the
	//     origin.
	//  16. `isRejectQueue`: Boolean indicating if the waiting room uses a reject queue.
	//  17. `isEventActive`: Boolean indicating if an event is currently occurring.
	//     Events are able to change a waiting room's behavior during a specified
	//     period of time. For additional information, look at the event properties
	//     `prequeue_start_time`, `event_start_time`, and `event_end_time` in the
	//     documentation for creating waiting room events. Events are considered active
	//     between these start and end times, as well as during the prequeueing period
	//     if it exists.
	//  18. `isEventPrequeueing`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if an event is currently prequeueing users before it starts.
	//  19. `timeUntilEventStart`: Valid only when `isEventPrequeueing` is **true**.
	//     Integer indicating the number of minutes until the event starts.
	//  20. `timeUntilEventStartFormatted`: String displaying the `timeUntilEventStart`
	//     formatted in English for users. If `isEventPrequeueing` is **false**,
	//     `timeUntilEventStartFormatted` will display **unavailable**.
	//  21. `timeUntilEventEnd`: Valid only when `isEventActive` is **true**. Integer
	//     indicating the number of minutes until the event ends.
	//  22. `timeUntilEventEndFormatted`: String displaying the `timeUntilEventEnd`
	//     formatted in English for users. If `isEventActive` is **false**,
	//     `timeUntilEventEndFormatted` will display **unavailable**.
	//  23. `shuffleAtEventStart`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if the users in the prequeue are shuffled randomly when the event
	//     starts.
	//
	// An example cURL to a waiting room could be:
	//
	//	curl -X GET "https://example.com/waitingroom" \
	//		-H "Accept: application/json"
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **fifo** and no event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 0,
	//			"waitTime50Percentile": 0,
	//			"waitTime75Percentile": 0,
	//			"waitTimeFormatted": "10 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "fifo",
	//			"isFIFOQueue": true,
	//			"isRandomQueue": false,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": false,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 0,
	//			"timeUntilEventEndFormatted": "unavailable",
	//			"shuffleAtEventStart": false
	//		}
	//	}
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **random** and an event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 5,
	//			"waitTime50Percentile": 10,
	//			"waitTime75Percentile": 15,
	//			"waitTimeFormatted": "5 minutes to 15 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "random",
	//			"isFIFOQueue": false,
	//			"isRandomQueue": true,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": true,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 15,
	//			"timeUntilEventEndFormatted": "15 minutes",
	//			"shuffleAtEventStart": true
	//		}
	//	}.
	JsonResponseEnabled param.Field[bool] `json:"json_response_enabled"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path param.Field[string] `json:"path"`
	// If queue_all is `true`, all the traffic that is coming to a route will be sent
	// to the waiting room. No new traffic can get to the route once this field is set
	// and estimated time will become unavailable.
	QueueAll param.Field[bool] `json:"queue_all"`
	// Sets the queueing method used by the waiting room. Changing this parameter from
	// the **default** queueing method is only available for the Waiting Room Advanced
	// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
	// event is prequeueing, users in the waiting room will not be accepted to the
	// origin. These users will always see a waiting room page that refreshes
	// automatically. The valid queueing methods are:
	//
	//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
	//     the order they arrived.
	//  2. `random`: Random queue where customers gain access randomly, regardless of
	//     arrival time.
	//  3. `passthrough`: Users will pass directly through the waiting room and into the
	//     origin website. As a result, any configured limits will not be respected
	//     while this is enabled. This method can be used as an alternative to disabling
	//     a waiting room (with `suspended`) so that analytics are still reported. This
	//     can be used if you wish to allow all traffic normally, but want to restrict
	//     traffic during a waiting room event, or vice versa.
	//  4. `reject`: Users will be immediately rejected from the waiting room. As a
	//     result, no users will reach the origin website while this is enabled. This
	//     can be used if you wish to reject all traffic while performing maintenance,
	//     block traffic during a specified period of time (an event), or block traffic
	//     while events are not occurring. Consider a waiting room used for vaccine
	//     distribution that only allows traffic during sign-up events, and otherwise
	//     blocks all traffic. For this case, the waiting room uses `reject`, and its
	//     events override this with `fifo`, `random`, or `passthrough`. When this
	//     queueing method is enabled and neither `queueAll` is enabled nor an event is
	//     prequeueing, the waiting room page **will not refresh automatically**.
	QueueingMethod param.Field[WaitingRoomNewParamsQueueingMethod] `json:"queueing_method"`
	// HTTP status code returned to a user while in the queue.
	QueueingStatusCode param.Field[WaitingRoomNewParamsQueueingStatusCode] `json:"queueing_status_code"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to
	// the route. If a user is not seen by Cloudflare again in that time period, they
	// will be treated as a new user that visits the route.
	SessionDuration param.Field[int64] `json:"session_duration"`
	// Suspends or allows traffic going to the waiting room. If set to `true`, the
	// traffic will not go to the waiting room.
	Suspended param.Field[bool] `json:"suspended"`
}

func (r WaitingRoomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaitingRoomNewParamsAdditionalRoute struct {
	// The hostname to which this waiting room will be applied (no wildcards). The
	// hostname must be the primary domain, subdomain, or custom hostname (if using SSL
	// for SaaS) of this zone. Please do not include the scheme (http:// or https://).
	Host param.Field[string] `json:"host"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path param.Field[string] `json:"path"`
}

func (r WaitingRoomNewParamsAdditionalRoute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures cookie attributes for the waiting room cookie. This encrypted cookie
// stores a user's status in the waiting room, such as queue position.
type WaitingRoomNewParamsCookieAttributes struct {
	// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
	// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
	// Note that when using value `none`, the secure attribute cannot be set to
	// `never`.
	Samesite param.Field[WaitingRoomNewParamsCookieAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on the waiting room cookie. Value `always`
	// indicates that the Secure attribute will be set in the Set-Cookie header,
	// `never` indicates that the Secure attribute will not be set, and `auto` will set
	// the Secure attribute depending if **Always Use HTTPS** is enabled.
	Secure param.Field[WaitingRoomNewParamsCookieAttributesSecure] `json:"secure"`
}

func (r WaitingRoomNewParamsCookieAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
// Note that when using value `none`, the secure attribute cannot be set to
// `never`.
type WaitingRoomNewParamsCookieAttributesSamesite string

const (
	WaitingRoomNewParamsCookieAttributesSamesiteAuto   WaitingRoomNewParamsCookieAttributesSamesite = "auto"
	WaitingRoomNewParamsCookieAttributesSamesiteLax    WaitingRoomNewParamsCookieAttributesSamesite = "lax"
	WaitingRoomNewParamsCookieAttributesSamesiteNone   WaitingRoomNewParamsCookieAttributesSamesite = "none"
	WaitingRoomNewParamsCookieAttributesSamesiteStrict WaitingRoomNewParamsCookieAttributesSamesite = "strict"
)

// Configures the Secure attribute on the waiting room cookie. Value `always`
// indicates that the Secure attribute will be set in the Set-Cookie header,
// `never` indicates that the Secure attribute will not be set, and `auto` will set
// the Secure attribute depending if **Always Use HTTPS** is enabled.
type WaitingRoomNewParamsCookieAttributesSecure string

const (
	WaitingRoomNewParamsCookieAttributesSecureAuto   WaitingRoomNewParamsCookieAttributesSecure = "auto"
	WaitingRoomNewParamsCookieAttributesSecureAlways WaitingRoomNewParamsCookieAttributesSecure = "always"
	WaitingRoomNewParamsCookieAttributesSecureNever  WaitingRoomNewParamsCookieAttributesSecure = "never"
)

// The language of the default page template. If no default_template_language is
// provided, then `en-US` (English) will be used.
type WaitingRoomNewParamsDefaultTemplateLanguage string

const (
	WaitingRoomNewParamsDefaultTemplateLanguageEnUs WaitingRoomNewParamsDefaultTemplateLanguage = "en-US"
	WaitingRoomNewParamsDefaultTemplateLanguageEsEs WaitingRoomNewParamsDefaultTemplateLanguage = "es-ES"
	WaitingRoomNewParamsDefaultTemplateLanguageDeDe WaitingRoomNewParamsDefaultTemplateLanguage = "de-DE"
	WaitingRoomNewParamsDefaultTemplateLanguageFrFr WaitingRoomNewParamsDefaultTemplateLanguage = "fr-FR"
	WaitingRoomNewParamsDefaultTemplateLanguageItIt WaitingRoomNewParamsDefaultTemplateLanguage = "it-IT"
	WaitingRoomNewParamsDefaultTemplateLanguageJaJp WaitingRoomNewParamsDefaultTemplateLanguage = "ja-JP"
	WaitingRoomNewParamsDefaultTemplateLanguageKoKr WaitingRoomNewParamsDefaultTemplateLanguage = "ko-KR"
	WaitingRoomNewParamsDefaultTemplateLanguagePtBr WaitingRoomNewParamsDefaultTemplateLanguage = "pt-BR"
	WaitingRoomNewParamsDefaultTemplateLanguageZhCn WaitingRoomNewParamsDefaultTemplateLanguage = "zh-CN"
	WaitingRoomNewParamsDefaultTemplateLanguageZhTw WaitingRoomNewParamsDefaultTemplateLanguage = "zh-TW"
	WaitingRoomNewParamsDefaultTemplateLanguageNlNl WaitingRoomNewParamsDefaultTemplateLanguage = "nl-NL"
	WaitingRoomNewParamsDefaultTemplateLanguagePlPl WaitingRoomNewParamsDefaultTemplateLanguage = "pl-PL"
	WaitingRoomNewParamsDefaultTemplateLanguageIDID WaitingRoomNewParamsDefaultTemplateLanguage = "id-ID"
	WaitingRoomNewParamsDefaultTemplateLanguageTrTr WaitingRoomNewParamsDefaultTemplateLanguage = "tr-TR"
	WaitingRoomNewParamsDefaultTemplateLanguageArEg WaitingRoomNewParamsDefaultTemplateLanguage = "ar-EG"
	WaitingRoomNewParamsDefaultTemplateLanguageRuRu WaitingRoomNewParamsDefaultTemplateLanguage = "ru-RU"
	WaitingRoomNewParamsDefaultTemplateLanguageFaIr WaitingRoomNewParamsDefaultTemplateLanguage = "fa-IR"
)

// Sets the queueing method used by the waiting room. Changing this parameter from
// the **default** queueing method is only available for the Waiting Room Advanced
// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
// event is prequeueing, users in the waiting room will not be accepted to the
// origin. These users will always see a waiting room page that refreshes
// automatically. The valid queueing methods are:
//
//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
//     the order they arrived.
//  2. `random`: Random queue where customers gain access randomly, regardless of
//     arrival time.
//  3. `passthrough`: Users will pass directly through the waiting room and into the
//     origin website. As a result, any configured limits will not be respected
//     while this is enabled. This method can be used as an alternative to disabling
//     a waiting room (with `suspended`) so that analytics are still reported. This
//     can be used if you wish to allow all traffic normally, but want to restrict
//     traffic during a waiting room event, or vice versa.
//  4. `reject`: Users will be immediately rejected from the waiting room. As a
//     result, no users will reach the origin website while this is enabled. This
//     can be used if you wish to reject all traffic while performing maintenance,
//     block traffic during a specified period of time (an event), or block traffic
//     while events are not occurring. Consider a waiting room used for vaccine
//     distribution that only allows traffic during sign-up events, and otherwise
//     blocks all traffic. For this case, the waiting room uses `reject`, and its
//     events override this with `fifo`, `random`, or `passthrough`. When this
//     queueing method is enabled and neither `queueAll` is enabled nor an event is
//     prequeueing, the waiting room page **will not refresh automatically**.
type WaitingRoomNewParamsQueueingMethod string

const (
	WaitingRoomNewParamsQueueingMethodFifo        WaitingRoomNewParamsQueueingMethod = "fifo"
	WaitingRoomNewParamsQueueingMethodRandom      WaitingRoomNewParamsQueueingMethod = "random"
	WaitingRoomNewParamsQueueingMethodPassthrough WaitingRoomNewParamsQueueingMethod = "passthrough"
	WaitingRoomNewParamsQueueingMethodReject      WaitingRoomNewParamsQueueingMethod = "reject"
)

// HTTP status code returned to a user while in the queue.
type WaitingRoomNewParamsQueueingStatusCode int64

const (
	WaitingRoomNewParamsQueueingStatusCode200 WaitingRoomNewParamsQueueingStatusCode = 200
	WaitingRoomNewParamsQueueingStatusCode202 WaitingRoomNewParamsQueueingStatusCode = 202
	WaitingRoomNewParamsQueueingStatusCode429 WaitingRoomNewParamsQueueingStatusCode = 429
)

type WaitingRoomNewResponseEnvelope struct {
	Result WaitingroomWaitingroom             `json:"result,required"`
	JSON   waitingRoomNewResponseEnvelopeJSON `json:"-"`
}

// waitingRoomNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [WaitingRoomNewResponseEnvelope]
type waitingRoomNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomUpdateParams struct {
	// The host name to which the waiting room will be applied (no wildcards). Please
	// do not include the scheme (http:// or https://). The host and path combination
	// must be unique.
	Host param.Field[string] `json:"host,required"`
	// A unique name to identify the waiting room. Only alphanumeric characters,
	// hyphens and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// Sets the number of new users that will be let into the route every minute. This
	// value is used as baseline for the number of users that are let in per minute. So
	// it is possible that there is a little more or little less traffic coming to the
	// route based on the traffic patterns at that time around the world.
	NewUsersPerMinute param.Field[int64] `json:"new_users_per_minute,required"`
	// Sets the total number of active user sessions on the route at a point in time. A
	// route is a combination of host and path on which a waiting room is available.
	// This value is used as a baseline for the total number of active user sessions on
	// the route. It is possible to have a situation where there are more or less
	// active users sessions on the route based on the traffic patterns at that time
	// around the world.
	TotalActiveUsers param.Field[int64] `json:"total_active_users,required"`
	// Only available for the Waiting Room Advanced subscription. Additional hostname
	// and path combinations to which this waiting room will be applied. There is an
	// implied wildcard at the end of the path. The hostname and path combination must
	// be unique to this and all other waiting rooms.
	AdditionalRoutes param.Field[[]WaitingRoomUpdateParamsAdditionalRoute] `json:"additional_routes"`
	// Configures cookie attributes for the waiting room cookie. This encrypted cookie
	// stores a user's status in the waiting room, such as queue position.
	CookieAttributes param.Field[WaitingRoomUpdateParamsCookieAttributes] `json:"cookie_attributes"`
	// Appends a '\_' + a custom suffix to the end of Cloudflare Waiting Room's cookie
	// name(**cf_waitingroom). If `cookie_suffix` is "abcd", the cookie name will be
	// `**cf_waitingroom_abcd`. This field is required if using `additional_routes`.
	CookieSuffix param.Field[string] `json:"cookie_suffix"`
	// Only available for the Waiting Room Advanced subscription. This is a template
	// html file that will be rendered at the edge. If no custom_page_html is provided,
	// the default waiting room will be used. The template is based on mustache (
	// https://mustache.github.io/ ). There are several variables that are evaluated by
	// the Cloudflare edge:
	//
	//  1. {{`waitTimeKnown`}} Acts like a boolean value that indicates the behavior to
	//     take when wait time is not available, for instance when queue_all is
	//     **true**.
	//  2. {{`waitTimeFormatted`}} Estimated wait time for the user. For example, five
	//     minutes. Alternatively, you can use:
	//  3. {{`waitTime`}} Number of minutes of estimated wait for a user.
	//  4. {{`waitTimeHours`}} Number of hours of estimated wait for a user
	//     (`Math.floor(waitTime/60)`).
	//  5. {{`waitTimeHourMinutes`}} Number of minutes above the `waitTimeHours` value
	//     (`waitTime%60`).
	//  6. {{`queueIsFull`}} Changes to **true** when no more people can be added to the
	//     queue.
	//
	// To view the full list of variables, look at the `cfWaitingRoom` object described
	// under the `json_response_enabled` property in other Waiting Room API calls.
	CustomPageHTML param.Field[string] `json:"custom_page_html"`
	// The language of the default page template. If no default_template_language is
	// provided, then `en-US` (English) will be used.
	DefaultTemplateLanguage param.Field[WaitingRoomUpdateParamsDefaultTemplateLanguage] `json:"default_template_language"`
	// A note that you can use to add more details about the waiting room.
	Description param.Field[string] `json:"description"`
	// Only available for the Waiting Room Advanced subscription. Disables automatic
	// renewal of session cookies. If `true`, an accepted user will have
	// session_duration minutes to browse the site. After that, they will have to go
	// through the waiting room again. If `false`, a user's session cookie will be
	// automatically renewed on every request.
	DisableSessionRenewal param.Field[bool] `json:"disable_session_renewal"`
	// Only available for the Waiting Room Advanced subscription. If `true`, requests
	// to the waiting room with the header `Accept: application/json` will receive a
	// JSON response object with information on the user's status in the waiting room
	// as opposed to the configured static HTML page. This JSON response object has one
	// property `cfWaitingRoom` which is an object containing the following fields:
	//
	//  1. `inWaitingRoom`: Boolean indicating if the user is in the waiting room
	//     (always **true**).
	//  2. `waitTimeKnown`: Boolean indicating if the current estimated wait times are
	//     accurate. If **false**, they are not available.
	//  3. `waitTime`: Valid only when `waitTimeKnown` is **true**. Integer indicating
	//     the current estimated time in minutes the user will wait in the waiting room.
	//     When `queueingMethod` is **random**, this is set to `waitTime50Percentile`.
	//  4. `waitTime25Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 25% of users that gain entry the fastest (25th percentile).
	//  5. `waitTime50Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 50% of users that gain entry the fastest (50th percentile).
	//     In other words, half of the queued users are expected to let into the origin
	//     website before `waitTime50Percentile` and half are expected to be let in
	//     after it.
	//  6. `waitTime75Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 75% of users that gain entry the fastest (75th percentile).
	//  7. `waitTimeFormatted`: String displaying the `waitTime` formatted in English
	//     for users. If `waitTimeKnown` is **false**, `waitTimeFormatted` will display
	//     **unavailable**.
	//  8. `queueIsFull`: Boolean indicating if the waiting room's queue is currently
	//     full and not accepting new users at the moment.
	//  9. `queueAll`: Boolean indicating if all users will be queued in the waiting
	//     room and no one will be let into the origin website.
	//  10. `lastUpdated`: String displaying the timestamp as an ISO 8601 string of the
	//     user's last attempt to leave the waiting room and be let into the origin
	//     website. The user is able to make another attempt after
	//     `refreshIntervalSeconds` past this time. If the user makes a request too
	//     soon, it will be ignored and `lastUpdated` will not change.
	//  11. `refreshIntervalSeconds`: Integer indicating the number of seconds after
	//     `lastUpdated` until the user is able to make another attempt to leave the
	//     waiting room and be let into the origin website. When the `queueingMethod`
	//     is `reject`, there is no specified refresh time — it will always be
	//     **zero**.
	//  12. `queueingMethod`: The queueing method currently used by the waiting room. It
	//     is either **fifo**, **random**, **passthrough**, or **reject**.
	//  13. `isFIFOQueue`: Boolean indicating if the waiting room uses a FIFO
	//     (First-In-First-Out) queue.
	//  14. `isRandomQueue`: Boolean indicating if the waiting room uses a Random queue
	//     where users gain access randomly.
	//  15. `isPassthroughQueue`: Boolean indicating if the waiting room uses a
	//     passthrough queue. Keep in mind that when passthrough is enabled, this JSON
	//     response will only exist when `queueAll` is **true** or `isEventPrequeueing`
	//     is **true** because in all other cases requests will go directly to the
	//     origin.
	//  16. `isRejectQueue`: Boolean indicating if the waiting room uses a reject queue.
	//  17. `isEventActive`: Boolean indicating if an event is currently occurring.
	//     Events are able to change a waiting room's behavior during a specified
	//     period of time. For additional information, look at the event properties
	//     `prequeue_start_time`, `event_start_time`, and `event_end_time` in the
	//     documentation for creating waiting room events. Events are considered active
	//     between these start and end times, as well as during the prequeueing period
	//     if it exists.
	//  18. `isEventPrequeueing`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if an event is currently prequeueing users before it starts.
	//  19. `timeUntilEventStart`: Valid only when `isEventPrequeueing` is **true**.
	//     Integer indicating the number of minutes until the event starts.
	//  20. `timeUntilEventStartFormatted`: String displaying the `timeUntilEventStart`
	//     formatted in English for users. If `isEventPrequeueing` is **false**,
	//     `timeUntilEventStartFormatted` will display **unavailable**.
	//  21. `timeUntilEventEnd`: Valid only when `isEventActive` is **true**. Integer
	//     indicating the number of minutes until the event ends.
	//  22. `timeUntilEventEndFormatted`: String displaying the `timeUntilEventEnd`
	//     formatted in English for users. If `isEventActive` is **false**,
	//     `timeUntilEventEndFormatted` will display **unavailable**.
	//  23. `shuffleAtEventStart`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if the users in the prequeue are shuffled randomly when the event
	//     starts.
	//
	// An example cURL to a waiting room could be:
	//
	//	curl -X GET "https://example.com/waitingroom" \
	//		-H "Accept: application/json"
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **fifo** and no event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 0,
	//			"waitTime50Percentile": 0,
	//			"waitTime75Percentile": 0,
	//			"waitTimeFormatted": "10 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "fifo",
	//			"isFIFOQueue": true,
	//			"isRandomQueue": false,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": false,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 0,
	//			"timeUntilEventEndFormatted": "unavailable",
	//			"shuffleAtEventStart": false
	//		}
	//	}
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **random** and an event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 5,
	//			"waitTime50Percentile": 10,
	//			"waitTime75Percentile": 15,
	//			"waitTimeFormatted": "5 minutes to 15 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "random",
	//			"isFIFOQueue": false,
	//			"isRandomQueue": true,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": true,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 15,
	//			"timeUntilEventEndFormatted": "15 minutes",
	//			"shuffleAtEventStart": true
	//		}
	//	}.
	JsonResponseEnabled param.Field[bool] `json:"json_response_enabled"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path param.Field[string] `json:"path"`
	// If queue_all is `true`, all the traffic that is coming to a route will be sent
	// to the waiting room. No new traffic can get to the route once this field is set
	// and estimated time will become unavailable.
	QueueAll param.Field[bool] `json:"queue_all"`
	// Sets the queueing method used by the waiting room. Changing this parameter from
	// the **default** queueing method is only available for the Waiting Room Advanced
	// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
	// event is prequeueing, users in the waiting room will not be accepted to the
	// origin. These users will always see a waiting room page that refreshes
	// automatically. The valid queueing methods are:
	//
	//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
	//     the order they arrived.
	//  2. `random`: Random queue where customers gain access randomly, regardless of
	//     arrival time.
	//  3. `passthrough`: Users will pass directly through the waiting room and into the
	//     origin website. As a result, any configured limits will not be respected
	//     while this is enabled. This method can be used as an alternative to disabling
	//     a waiting room (with `suspended`) so that analytics are still reported. This
	//     can be used if you wish to allow all traffic normally, but want to restrict
	//     traffic during a waiting room event, or vice versa.
	//  4. `reject`: Users will be immediately rejected from the waiting room. As a
	//     result, no users will reach the origin website while this is enabled. This
	//     can be used if you wish to reject all traffic while performing maintenance,
	//     block traffic during a specified period of time (an event), or block traffic
	//     while events are not occurring. Consider a waiting room used for vaccine
	//     distribution that only allows traffic during sign-up events, and otherwise
	//     blocks all traffic. For this case, the waiting room uses `reject`, and its
	//     events override this with `fifo`, `random`, or `passthrough`. When this
	//     queueing method is enabled and neither `queueAll` is enabled nor an event is
	//     prequeueing, the waiting room page **will not refresh automatically**.
	QueueingMethod param.Field[WaitingRoomUpdateParamsQueueingMethod] `json:"queueing_method"`
	// HTTP status code returned to a user while in the queue.
	QueueingStatusCode param.Field[WaitingRoomUpdateParamsQueueingStatusCode] `json:"queueing_status_code"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to
	// the route. If a user is not seen by Cloudflare again in that time period, they
	// will be treated as a new user that visits the route.
	SessionDuration param.Field[int64] `json:"session_duration"`
	// Suspends or allows traffic going to the waiting room. If set to `true`, the
	// traffic will not go to the waiting room.
	Suspended param.Field[bool] `json:"suspended"`
}

func (r WaitingRoomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaitingRoomUpdateParamsAdditionalRoute struct {
	// The hostname to which this waiting room will be applied (no wildcards). The
	// hostname must be the primary domain, subdomain, or custom hostname (if using SSL
	// for SaaS) of this zone. Please do not include the scheme (http:// or https://).
	Host param.Field[string] `json:"host"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path param.Field[string] `json:"path"`
}

func (r WaitingRoomUpdateParamsAdditionalRoute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures cookie attributes for the waiting room cookie. This encrypted cookie
// stores a user's status in the waiting room, such as queue position.
type WaitingRoomUpdateParamsCookieAttributes struct {
	// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
	// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
	// Note that when using value `none`, the secure attribute cannot be set to
	// `never`.
	Samesite param.Field[WaitingRoomUpdateParamsCookieAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on the waiting room cookie. Value `always`
	// indicates that the Secure attribute will be set in the Set-Cookie header,
	// `never` indicates that the Secure attribute will not be set, and `auto` will set
	// the Secure attribute depending if **Always Use HTTPS** is enabled.
	Secure param.Field[WaitingRoomUpdateParamsCookieAttributesSecure] `json:"secure"`
}

func (r WaitingRoomUpdateParamsCookieAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
// Note that when using value `none`, the secure attribute cannot be set to
// `never`.
type WaitingRoomUpdateParamsCookieAttributesSamesite string

const (
	WaitingRoomUpdateParamsCookieAttributesSamesiteAuto   WaitingRoomUpdateParamsCookieAttributesSamesite = "auto"
	WaitingRoomUpdateParamsCookieAttributesSamesiteLax    WaitingRoomUpdateParamsCookieAttributesSamesite = "lax"
	WaitingRoomUpdateParamsCookieAttributesSamesiteNone   WaitingRoomUpdateParamsCookieAttributesSamesite = "none"
	WaitingRoomUpdateParamsCookieAttributesSamesiteStrict WaitingRoomUpdateParamsCookieAttributesSamesite = "strict"
)

// Configures the Secure attribute on the waiting room cookie. Value `always`
// indicates that the Secure attribute will be set in the Set-Cookie header,
// `never` indicates that the Secure attribute will not be set, and `auto` will set
// the Secure attribute depending if **Always Use HTTPS** is enabled.
type WaitingRoomUpdateParamsCookieAttributesSecure string

const (
	WaitingRoomUpdateParamsCookieAttributesSecureAuto   WaitingRoomUpdateParamsCookieAttributesSecure = "auto"
	WaitingRoomUpdateParamsCookieAttributesSecureAlways WaitingRoomUpdateParamsCookieAttributesSecure = "always"
	WaitingRoomUpdateParamsCookieAttributesSecureNever  WaitingRoomUpdateParamsCookieAttributesSecure = "never"
)

// The language of the default page template. If no default_template_language is
// provided, then `en-US` (English) will be used.
type WaitingRoomUpdateParamsDefaultTemplateLanguage string

const (
	WaitingRoomUpdateParamsDefaultTemplateLanguageEnUs WaitingRoomUpdateParamsDefaultTemplateLanguage = "en-US"
	WaitingRoomUpdateParamsDefaultTemplateLanguageEsEs WaitingRoomUpdateParamsDefaultTemplateLanguage = "es-ES"
	WaitingRoomUpdateParamsDefaultTemplateLanguageDeDe WaitingRoomUpdateParamsDefaultTemplateLanguage = "de-DE"
	WaitingRoomUpdateParamsDefaultTemplateLanguageFrFr WaitingRoomUpdateParamsDefaultTemplateLanguage = "fr-FR"
	WaitingRoomUpdateParamsDefaultTemplateLanguageItIt WaitingRoomUpdateParamsDefaultTemplateLanguage = "it-IT"
	WaitingRoomUpdateParamsDefaultTemplateLanguageJaJp WaitingRoomUpdateParamsDefaultTemplateLanguage = "ja-JP"
	WaitingRoomUpdateParamsDefaultTemplateLanguageKoKr WaitingRoomUpdateParamsDefaultTemplateLanguage = "ko-KR"
	WaitingRoomUpdateParamsDefaultTemplateLanguagePtBr WaitingRoomUpdateParamsDefaultTemplateLanguage = "pt-BR"
	WaitingRoomUpdateParamsDefaultTemplateLanguageZhCn WaitingRoomUpdateParamsDefaultTemplateLanguage = "zh-CN"
	WaitingRoomUpdateParamsDefaultTemplateLanguageZhTw WaitingRoomUpdateParamsDefaultTemplateLanguage = "zh-TW"
	WaitingRoomUpdateParamsDefaultTemplateLanguageNlNl WaitingRoomUpdateParamsDefaultTemplateLanguage = "nl-NL"
	WaitingRoomUpdateParamsDefaultTemplateLanguagePlPl WaitingRoomUpdateParamsDefaultTemplateLanguage = "pl-PL"
	WaitingRoomUpdateParamsDefaultTemplateLanguageIDID WaitingRoomUpdateParamsDefaultTemplateLanguage = "id-ID"
	WaitingRoomUpdateParamsDefaultTemplateLanguageTrTr WaitingRoomUpdateParamsDefaultTemplateLanguage = "tr-TR"
	WaitingRoomUpdateParamsDefaultTemplateLanguageArEg WaitingRoomUpdateParamsDefaultTemplateLanguage = "ar-EG"
	WaitingRoomUpdateParamsDefaultTemplateLanguageRuRu WaitingRoomUpdateParamsDefaultTemplateLanguage = "ru-RU"
	WaitingRoomUpdateParamsDefaultTemplateLanguageFaIr WaitingRoomUpdateParamsDefaultTemplateLanguage = "fa-IR"
)

// Sets the queueing method used by the waiting room. Changing this parameter from
// the **default** queueing method is only available for the Waiting Room Advanced
// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
// event is prequeueing, users in the waiting room will not be accepted to the
// origin. These users will always see a waiting room page that refreshes
// automatically. The valid queueing methods are:
//
//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
//     the order they arrived.
//  2. `random`: Random queue where customers gain access randomly, regardless of
//     arrival time.
//  3. `passthrough`: Users will pass directly through the waiting room and into the
//     origin website. As a result, any configured limits will not be respected
//     while this is enabled. This method can be used as an alternative to disabling
//     a waiting room (with `suspended`) so that analytics are still reported. This
//     can be used if you wish to allow all traffic normally, but want to restrict
//     traffic during a waiting room event, or vice versa.
//  4. `reject`: Users will be immediately rejected from the waiting room. As a
//     result, no users will reach the origin website while this is enabled. This
//     can be used if you wish to reject all traffic while performing maintenance,
//     block traffic during a specified period of time (an event), or block traffic
//     while events are not occurring. Consider a waiting room used for vaccine
//     distribution that only allows traffic during sign-up events, and otherwise
//     blocks all traffic. For this case, the waiting room uses `reject`, and its
//     events override this with `fifo`, `random`, or `passthrough`. When this
//     queueing method is enabled and neither `queueAll` is enabled nor an event is
//     prequeueing, the waiting room page **will not refresh automatically**.
type WaitingRoomUpdateParamsQueueingMethod string

const (
	WaitingRoomUpdateParamsQueueingMethodFifo        WaitingRoomUpdateParamsQueueingMethod = "fifo"
	WaitingRoomUpdateParamsQueueingMethodRandom      WaitingRoomUpdateParamsQueueingMethod = "random"
	WaitingRoomUpdateParamsQueueingMethodPassthrough WaitingRoomUpdateParamsQueueingMethod = "passthrough"
	WaitingRoomUpdateParamsQueueingMethodReject      WaitingRoomUpdateParamsQueueingMethod = "reject"
)

// HTTP status code returned to a user while in the queue.
type WaitingRoomUpdateParamsQueueingStatusCode int64

const (
	WaitingRoomUpdateParamsQueueingStatusCode200 WaitingRoomUpdateParamsQueueingStatusCode = 200
	WaitingRoomUpdateParamsQueueingStatusCode202 WaitingRoomUpdateParamsQueueingStatusCode = 202
	WaitingRoomUpdateParamsQueueingStatusCode429 WaitingRoomUpdateParamsQueueingStatusCode = 429
)

type WaitingRoomUpdateResponseEnvelope struct {
	Result WaitingroomWaitingroom                `json:"result,required"`
	JSON   waitingRoomUpdateResponseEnvelopeJSON `json:"-"`
}

// waitingRoomUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WaitingRoomUpdateResponseEnvelope]
type waitingRoomUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomListResponseEnvelope struct {
	Errors   []WaitingRoomListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WaitingRoomListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WaitingroomWaitingroom                  `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WaitingRoomListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WaitingRoomListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       waitingRoomListResponseEnvelopeJSON       `json:"-"`
}

// waitingRoomListResponseEnvelopeJSON contains the JSON metadata for the struct
// [WaitingRoomListResponseEnvelope]
type waitingRoomListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    waitingRoomListResponseEnvelopeErrorsJSON `json:"-"`
}

// waitingRoomListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WaitingRoomListResponseEnvelopeErrors]
type waitingRoomListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    waitingRoomListResponseEnvelopeMessagesJSON `json:"-"`
}

// waitingRoomListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WaitingRoomListResponseEnvelopeMessages]
type waitingRoomListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WaitingRoomListResponseEnvelopeSuccess bool

const (
	WaitingRoomListResponseEnvelopeSuccessTrue WaitingRoomListResponseEnvelopeSuccess = true
)

type WaitingRoomListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       waitingRoomListResponseEnvelopeResultInfoJSON `json:"-"`
}

// waitingRoomListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [WaitingRoomListResponseEnvelopeResultInfo]
type waitingRoomListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomDeleteResponseEnvelope struct {
	Result WaitingRoomDeleteResponse             `json:"result,required"`
	JSON   waitingRoomDeleteResponseEnvelopeJSON `json:"-"`
}

// waitingRoomDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [WaitingRoomDeleteResponseEnvelope]
type waitingRoomDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomEditParams struct {
	// The host name to which the waiting room will be applied (no wildcards). Please
	// do not include the scheme (http:// or https://). The host and path combination
	// must be unique.
	Host param.Field[string] `json:"host,required"`
	// A unique name to identify the waiting room. Only alphanumeric characters,
	// hyphens and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// Sets the number of new users that will be let into the route every minute. This
	// value is used as baseline for the number of users that are let in per minute. So
	// it is possible that there is a little more or little less traffic coming to the
	// route based on the traffic patterns at that time around the world.
	NewUsersPerMinute param.Field[int64] `json:"new_users_per_minute,required"`
	// Sets the total number of active user sessions on the route at a point in time. A
	// route is a combination of host and path on which a waiting room is available.
	// This value is used as a baseline for the total number of active user sessions on
	// the route. It is possible to have a situation where there are more or less
	// active users sessions on the route based on the traffic patterns at that time
	// around the world.
	TotalActiveUsers param.Field[int64] `json:"total_active_users,required"`
	// Only available for the Waiting Room Advanced subscription. Additional hostname
	// and path combinations to which this waiting room will be applied. There is an
	// implied wildcard at the end of the path. The hostname and path combination must
	// be unique to this and all other waiting rooms.
	AdditionalRoutes param.Field[[]WaitingRoomEditParamsAdditionalRoute] `json:"additional_routes"`
	// Configures cookie attributes for the waiting room cookie. This encrypted cookie
	// stores a user's status in the waiting room, such as queue position.
	CookieAttributes param.Field[WaitingRoomEditParamsCookieAttributes] `json:"cookie_attributes"`
	// Appends a '\_' + a custom suffix to the end of Cloudflare Waiting Room's cookie
	// name(**cf_waitingroom). If `cookie_suffix` is "abcd", the cookie name will be
	// `**cf_waitingroom_abcd`. This field is required if using `additional_routes`.
	CookieSuffix param.Field[string] `json:"cookie_suffix"`
	// Only available for the Waiting Room Advanced subscription. This is a template
	// html file that will be rendered at the edge. If no custom_page_html is provided,
	// the default waiting room will be used. The template is based on mustache (
	// https://mustache.github.io/ ). There are several variables that are evaluated by
	// the Cloudflare edge:
	//
	//  1. {{`waitTimeKnown`}} Acts like a boolean value that indicates the behavior to
	//     take when wait time is not available, for instance when queue_all is
	//     **true**.
	//  2. {{`waitTimeFormatted`}} Estimated wait time for the user. For example, five
	//     minutes. Alternatively, you can use:
	//  3. {{`waitTime`}} Number of minutes of estimated wait for a user.
	//  4. {{`waitTimeHours`}} Number of hours of estimated wait for a user
	//     (`Math.floor(waitTime/60)`).
	//  5. {{`waitTimeHourMinutes`}} Number of minutes above the `waitTimeHours` value
	//     (`waitTime%60`).
	//  6. {{`queueIsFull`}} Changes to **true** when no more people can be added to the
	//     queue.
	//
	// To view the full list of variables, look at the `cfWaitingRoom` object described
	// under the `json_response_enabled` property in other Waiting Room API calls.
	CustomPageHTML param.Field[string] `json:"custom_page_html"`
	// The language of the default page template. If no default_template_language is
	// provided, then `en-US` (English) will be used.
	DefaultTemplateLanguage param.Field[WaitingRoomEditParamsDefaultTemplateLanguage] `json:"default_template_language"`
	// A note that you can use to add more details about the waiting room.
	Description param.Field[string] `json:"description"`
	// Only available for the Waiting Room Advanced subscription. Disables automatic
	// renewal of session cookies. If `true`, an accepted user will have
	// session_duration minutes to browse the site. After that, they will have to go
	// through the waiting room again. If `false`, a user's session cookie will be
	// automatically renewed on every request.
	DisableSessionRenewal param.Field[bool] `json:"disable_session_renewal"`
	// Only available for the Waiting Room Advanced subscription. If `true`, requests
	// to the waiting room with the header `Accept: application/json` will receive a
	// JSON response object with information on the user's status in the waiting room
	// as opposed to the configured static HTML page. This JSON response object has one
	// property `cfWaitingRoom` which is an object containing the following fields:
	//
	//  1. `inWaitingRoom`: Boolean indicating if the user is in the waiting room
	//     (always **true**).
	//  2. `waitTimeKnown`: Boolean indicating if the current estimated wait times are
	//     accurate. If **false**, they are not available.
	//  3. `waitTime`: Valid only when `waitTimeKnown` is **true**. Integer indicating
	//     the current estimated time in minutes the user will wait in the waiting room.
	//     When `queueingMethod` is **random**, this is set to `waitTime50Percentile`.
	//  4. `waitTime25Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 25% of users that gain entry the fastest (25th percentile).
	//  5. `waitTime50Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 50% of users that gain entry the fastest (50th percentile).
	//     In other words, half of the queued users are expected to let into the origin
	//     website before `waitTime50Percentile` and half are expected to be let in
	//     after it.
	//  6. `waitTime75Percentile`: Valid only when `queueingMethod` is **random** and
	//     `waitTimeKnown` is **true**. Integer indicating the current estimated maximum
	//     wait time for the 75% of users that gain entry the fastest (75th percentile).
	//  7. `waitTimeFormatted`: String displaying the `waitTime` formatted in English
	//     for users. If `waitTimeKnown` is **false**, `waitTimeFormatted` will display
	//     **unavailable**.
	//  8. `queueIsFull`: Boolean indicating if the waiting room's queue is currently
	//     full and not accepting new users at the moment.
	//  9. `queueAll`: Boolean indicating if all users will be queued in the waiting
	//     room and no one will be let into the origin website.
	//  10. `lastUpdated`: String displaying the timestamp as an ISO 8601 string of the
	//     user's last attempt to leave the waiting room and be let into the origin
	//     website. The user is able to make another attempt after
	//     `refreshIntervalSeconds` past this time. If the user makes a request too
	//     soon, it will be ignored and `lastUpdated` will not change.
	//  11. `refreshIntervalSeconds`: Integer indicating the number of seconds after
	//     `lastUpdated` until the user is able to make another attempt to leave the
	//     waiting room and be let into the origin website. When the `queueingMethod`
	//     is `reject`, there is no specified refresh time — it will always be
	//     **zero**.
	//  12. `queueingMethod`: The queueing method currently used by the waiting room. It
	//     is either **fifo**, **random**, **passthrough**, or **reject**.
	//  13. `isFIFOQueue`: Boolean indicating if the waiting room uses a FIFO
	//     (First-In-First-Out) queue.
	//  14. `isRandomQueue`: Boolean indicating if the waiting room uses a Random queue
	//     where users gain access randomly.
	//  15. `isPassthroughQueue`: Boolean indicating if the waiting room uses a
	//     passthrough queue. Keep in mind that when passthrough is enabled, this JSON
	//     response will only exist when `queueAll` is **true** or `isEventPrequeueing`
	//     is **true** because in all other cases requests will go directly to the
	//     origin.
	//  16. `isRejectQueue`: Boolean indicating if the waiting room uses a reject queue.
	//  17. `isEventActive`: Boolean indicating if an event is currently occurring.
	//     Events are able to change a waiting room's behavior during a specified
	//     period of time. For additional information, look at the event properties
	//     `prequeue_start_time`, `event_start_time`, and `event_end_time` in the
	//     documentation for creating waiting room events. Events are considered active
	//     between these start and end times, as well as during the prequeueing period
	//     if it exists.
	//  18. `isEventPrequeueing`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if an event is currently prequeueing users before it starts.
	//  19. `timeUntilEventStart`: Valid only when `isEventPrequeueing` is **true**.
	//     Integer indicating the number of minutes until the event starts.
	//  20. `timeUntilEventStartFormatted`: String displaying the `timeUntilEventStart`
	//     formatted in English for users. If `isEventPrequeueing` is **false**,
	//     `timeUntilEventStartFormatted` will display **unavailable**.
	//  21. `timeUntilEventEnd`: Valid only when `isEventActive` is **true**. Integer
	//     indicating the number of minutes until the event ends.
	//  22. `timeUntilEventEndFormatted`: String displaying the `timeUntilEventEnd`
	//     formatted in English for users. If `isEventActive` is **false**,
	//     `timeUntilEventEndFormatted` will display **unavailable**.
	//  23. `shuffleAtEventStart`: Valid only when `isEventActive` is **true**. Boolean
	//     indicating if the users in the prequeue are shuffled randomly when the event
	//     starts.
	//
	// An example cURL to a waiting room could be:
	//
	//	curl -X GET "https://example.com/waitingroom" \
	//		-H "Accept: application/json"
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **fifo** and no event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 0,
	//			"waitTime50Percentile": 0,
	//			"waitTime75Percentile": 0,
	//			"waitTimeFormatted": "10 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "fifo",
	//			"isFIFOQueue": true,
	//			"isRandomQueue": false,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": false,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 0,
	//			"timeUntilEventEndFormatted": "unavailable",
	//			"shuffleAtEventStart": false
	//		}
	//	}
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an
	// example JSON response when `queueingMethod` is **random** and an event is active
	// could be:
	//
	//	{
	//		"cfWaitingRoom": {
	//			"inWaitingRoom": true,
	//			"waitTimeKnown": true,
	//			"waitTime": 10,
	//			"waitTime25Percentile": 5,
	//			"waitTime50Percentile": 10,
	//			"waitTime75Percentile": 15,
	//			"waitTimeFormatted": "5 minutes to 15 minutes",
	//			"queueIsFull": false,
	//			"queueAll": false,
	//			"lastUpdated": "2020-08-03T23:46:00.000Z",
	//			"refreshIntervalSeconds": 20,
	//			"queueingMethod": "random",
	//			"isFIFOQueue": false,
	//			"isRandomQueue": true,
	//			"isPassthroughQueue": false,
	//			"isRejectQueue": false,
	//			"isEventActive": true,
	//			"isEventPrequeueing": false,
	//			"timeUntilEventStart": 0,
	//			"timeUntilEventStartFormatted": "unavailable",
	//			"timeUntilEventEnd": 15,
	//			"timeUntilEventEndFormatted": "15 minutes",
	//			"shuffleAtEventStart": true
	//		}
	//	}.
	JsonResponseEnabled param.Field[bool] `json:"json_response_enabled"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path param.Field[string] `json:"path"`
	// If queue_all is `true`, all the traffic that is coming to a route will be sent
	// to the waiting room. No new traffic can get to the route once this field is set
	// and estimated time will become unavailable.
	QueueAll param.Field[bool] `json:"queue_all"`
	// Sets the queueing method used by the waiting room. Changing this parameter from
	// the **default** queueing method is only available for the Waiting Room Advanced
	// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
	// event is prequeueing, users in the waiting room will not be accepted to the
	// origin. These users will always see a waiting room page that refreshes
	// automatically. The valid queueing methods are:
	//
	//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
	//     the order they arrived.
	//  2. `random`: Random queue where customers gain access randomly, regardless of
	//     arrival time.
	//  3. `passthrough`: Users will pass directly through the waiting room and into the
	//     origin website. As a result, any configured limits will not be respected
	//     while this is enabled. This method can be used as an alternative to disabling
	//     a waiting room (with `suspended`) so that analytics are still reported. This
	//     can be used if you wish to allow all traffic normally, but want to restrict
	//     traffic during a waiting room event, or vice versa.
	//  4. `reject`: Users will be immediately rejected from the waiting room. As a
	//     result, no users will reach the origin website while this is enabled. This
	//     can be used if you wish to reject all traffic while performing maintenance,
	//     block traffic during a specified period of time (an event), or block traffic
	//     while events are not occurring. Consider a waiting room used for vaccine
	//     distribution that only allows traffic during sign-up events, and otherwise
	//     blocks all traffic. For this case, the waiting room uses `reject`, and its
	//     events override this with `fifo`, `random`, or `passthrough`. When this
	//     queueing method is enabled and neither `queueAll` is enabled nor an event is
	//     prequeueing, the waiting room page **will not refresh automatically**.
	QueueingMethod param.Field[WaitingRoomEditParamsQueueingMethod] `json:"queueing_method"`
	// HTTP status code returned to a user while in the queue.
	QueueingStatusCode param.Field[WaitingRoomEditParamsQueueingStatusCode] `json:"queueing_status_code"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to
	// the route. If a user is not seen by Cloudflare again in that time period, they
	// will be treated as a new user that visits the route.
	SessionDuration param.Field[int64] `json:"session_duration"`
	// Suspends or allows traffic going to the waiting room. If set to `true`, the
	// traffic will not go to the waiting room.
	Suspended param.Field[bool] `json:"suspended"`
}

func (r WaitingRoomEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaitingRoomEditParamsAdditionalRoute struct {
	// The hostname to which this waiting room will be applied (no wildcards). The
	// hostname must be the primary domain, subdomain, or custom hostname (if using SSL
	// for SaaS) of this zone. Please do not include the scheme (http:// or https://).
	Host param.Field[string] `json:"host"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path param.Field[string] `json:"path"`
}

func (r WaitingRoomEditParamsAdditionalRoute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures cookie attributes for the waiting room cookie. This encrypted cookie
// stores a user's status in the waiting room, such as queue position.
type WaitingRoomEditParamsCookieAttributes struct {
	// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
	// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
	// Note that when using value `none`, the secure attribute cannot be set to
	// `never`.
	Samesite param.Field[WaitingRoomEditParamsCookieAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on the waiting room cookie. Value `always`
	// indicates that the Secure attribute will be set in the Set-Cookie header,
	// `never` indicates that the Secure attribute will not be set, and `auto` will set
	// the Secure attribute depending if **Always Use HTTPS** is enabled.
	Secure param.Field[WaitingRoomEditParamsCookieAttributesSecure] `json:"secure"`
}

func (r WaitingRoomEditParamsCookieAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
// Note that when using value `none`, the secure attribute cannot be set to
// `never`.
type WaitingRoomEditParamsCookieAttributesSamesite string

const (
	WaitingRoomEditParamsCookieAttributesSamesiteAuto   WaitingRoomEditParamsCookieAttributesSamesite = "auto"
	WaitingRoomEditParamsCookieAttributesSamesiteLax    WaitingRoomEditParamsCookieAttributesSamesite = "lax"
	WaitingRoomEditParamsCookieAttributesSamesiteNone   WaitingRoomEditParamsCookieAttributesSamesite = "none"
	WaitingRoomEditParamsCookieAttributesSamesiteStrict WaitingRoomEditParamsCookieAttributesSamesite = "strict"
)

// Configures the Secure attribute on the waiting room cookie. Value `always`
// indicates that the Secure attribute will be set in the Set-Cookie header,
// `never` indicates that the Secure attribute will not be set, and `auto` will set
// the Secure attribute depending if **Always Use HTTPS** is enabled.
type WaitingRoomEditParamsCookieAttributesSecure string

const (
	WaitingRoomEditParamsCookieAttributesSecureAuto   WaitingRoomEditParamsCookieAttributesSecure = "auto"
	WaitingRoomEditParamsCookieAttributesSecureAlways WaitingRoomEditParamsCookieAttributesSecure = "always"
	WaitingRoomEditParamsCookieAttributesSecureNever  WaitingRoomEditParamsCookieAttributesSecure = "never"
)

// The language of the default page template. If no default_template_language is
// provided, then `en-US` (English) will be used.
type WaitingRoomEditParamsDefaultTemplateLanguage string

const (
	WaitingRoomEditParamsDefaultTemplateLanguageEnUs WaitingRoomEditParamsDefaultTemplateLanguage = "en-US"
	WaitingRoomEditParamsDefaultTemplateLanguageEsEs WaitingRoomEditParamsDefaultTemplateLanguage = "es-ES"
	WaitingRoomEditParamsDefaultTemplateLanguageDeDe WaitingRoomEditParamsDefaultTemplateLanguage = "de-DE"
	WaitingRoomEditParamsDefaultTemplateLanguageFrFr WaitingRoomEditParamsDefaultTemplateLanguage = "fr-FR"
	WaitingRoomEditParamsDefaultTemplateLanguageItIt WaitingRoomEditParamsDefaultTemplateLanguage = "it-IT"
	WaitingRoomEditParamsDefaultTemplateLanguageJaJp WaitingRoomEditParamsDefaultTemplateLanguage = "ja-JP"
	WaitingRoomEditParamsDefaultTemplateLanguageKoKr WaitingRoomEditParamsDefaultTemplateLanguage = "ko-KR"
	WaitingRoomEditParamsDefaultTemplateLanguagePtBr WaitingRoomEditParamsDefaultTemplateLanguage = "pt-BR"
	WaitingRoomEditParamsDefaultTemplateLanguageZhCn WaitingRoomEditParamsDefaultTemplateLanguage = "zh-CN"
	WaitingRoomEditParamsDefaultTemplateLanguageZhTw WaitingRoomEditParamsDefaultTemplateLanguage = "zh-TW"
	WaitingRoomEditParamsDefaultTemplateLanguageNlNl WaitingRoomEditParamsDefaultTemplateLanguage = "nl-NL"
	WaitingRoomEditParamsDefaultTemplateLanguagePlPl WaitingRoomEditParamsDefaultTemplateLanguage = "pl-PL"
	WaitingRoomEditParamsDefaultTemplateLanguageIDID WaitingRoomEditParamsDefaultTemplateLanguage = "id-ID"
	WaitingRoomEditParamsDefaultTemplateLanguageTrTr WaitingRoomEditParamsDefaultTemplateLanguage = "tr-TR"
	WaitingRoomEditParamsDefaultTemplateLanguageArEg WaitingRoomEditParamsDefaultTemplateLanguage = "ar-EG"
	WaitingRoomEditParamsDefaultTemplateLanguageRuRu WaitingRoomEditParamsDefaultTemplateLanguage = "ru-RU"
	WaitingRoomEditParamsDefaultTemplateLanguageFaIr WaitingRoomEditParamsDefaultTemplateLanguage = "fa-IR"
)

// Sets the queueing method used by the waiting room. Changing this parameter from
// the **default** queueing method is only available for the Waiting Room Advanced
// subscription. Regardless of the queueing method, if `queue_all` is enabled or an
// event is prequeueing, users in the waiting room will not be accepted to the
// origin. These users will always see a waiting room page that refreshes
// automatically. The valid queueing methods are:
//
//  1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in
//     the order they arrived.
//  2. `random`: Random queue where customers gain access randomly, regardless of
//     arrival time.
//  3. `passthrough`: Users will pass directly through the waiting room and into the
//     origin website. As a result, any configured limits will not be respected
//     while this is enabled. This method can be used as an alternative to disabling
//     a waiting room (with `suspended`) so that analytics are still reported. This
//     can be used if you wish to allow all traffic normally, but want to restrict
//     traffic during a waiting room event, or vice versa.
//  4. `reject`: Users will be immediately rejected from the waiting room. As a
//     result, no users will reach the origin website while this is enabled. This
//     can be used if you wish to reject all traffic while performing maintenance,
//     block traffic during a specified period of time (an event), or block traffic
//     while events are not occurring. Consider a waiting room used for vaccine
//     distribution that only allows traffic during sign-up events, and otherwise
//     blocks all traffic. For this case, the waiting room uses `reject`, and its
//     events override this with `fifo`, `random`, or `passthrough`. When this
//     queueing method is enabled and neither `queueAll` is enabled nor an event is
//     prequeueing, the waiting room page **will not refresh automatically**.
type WaitingRoomEditParamsQueueingMethod string

const (
	WaitingRoomEditParamsQueueingMethodFifo        WaitingRoomEditParamsQueueingMethod = "fifo"
	WaitingRoomEditParamsQueueingMethodRandom      WaitingRoomEditParamsQueueingMethod = "random"
	WaitingRoomEditParamsQueueingMethodPassthrough WaitingRoomEditParamsQueueingMethod = "passthrough"
	WaitingRoomEditParamsQueueingMethodReject      WaitingRoomEditParamsQueueingMethod = "reject"
)

// HTTP status code returned to a user while in the queue.
type WaitingRoomEditParamsQueueingStatusCode int64

const (
	WaitingRoomEditParamsQueueingStatusCode200 WaitingRoomEditParamsQueueingStatusCode = 200
	WaitingRoomEditParamsQueueingStatusCode202 WaitingRoomEditParamsQueueingStatusCode = 202
	WaitingRoomEditParamsQueueingStatusCode429 WaitingRoomEditParamsQueueingStatusCode = 429
)

type WaitingRoomEditResponseEnvelope struct {
	Result WaitingroomWaitingroom              `json:"result,required"`
	JSON   waitingRoomEditResponseEnvelopeJSON `json:"-"`
}

// waitingRoomEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [WaitingRoomEditResponseEnvelope]
type waitingRoomEditResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomGetResponseEnvelope struct {
	Result WaitingroomWaitingroom             `json:"result,required"`
	JSON   waitingRoomGetResponseEnvelopeJSON `json:"-"`
}

// waitingRoomGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WaitingRoomGetResponseEnvelope]
type waitingRoomGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
