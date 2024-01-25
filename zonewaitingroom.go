// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneWaitingRoomService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneWaitingRoomService] method
// instead.
type ZoneWaitingRoomService struct {
	Options  []option.RequestOption
	Previews *ZoneWaitingRoomPreviewService
	Events   *ZoneWaitingRoomEventService
	Rules    *ZoneWaitingRoomRuleService
	Statuses *ZoneWaitingRoomStatusService
	Settings *ZoneWaitingRoomSettingService
}

// NewZoneWaitingRoomService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneWaitingRoomService(opts ...option.RequestOption) (r *ZoneWaitingRoomService) {
	r = &ZoneWaitingRoomService{}
	r.Options = opts
	r.Previews = NewZoneWaitingRoomPreviewService(opts...)
	r.Events = NewZoneWaitingRoomEventService(opts...)
	r.Rules = NewZoneWaitingRoomRuleService(opts...)
	r.Statuses = NewZoneWaitingRoomStatusService(opts...)
	r.Settings = NewZoneWaitingRoomSettingService(opts...)
	return
}

// Fetches a single configured waiting room.
func (r *ZoneWaitingRoomService) Get(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, opts ...option.RequestOption) (res *SingleWaitingRoomResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured waiting room.
func (r *ZoneWaitingRoomService) Update(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body ZoneWaitingRoomUpdateParams, opts ...option.RequestOption) (res *SingleWaitingRoomResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a waiting room.
func (r *ZoneWaitingRoomService) Delete(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, opts ...option.RequestOption) (res *ZoneWaitingRoomDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Patches a configured waiting room.
func (r *ZoneWaitingRoomService) Patch(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body ZoneWaitingRoomPatchParams, opts ...option.RequestOption) (res *SingleWaitingRoomResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Creates a new waiting room.
func (r *ZoneWaitingRoomService) WaitingRoomNewWaitingRoom(ctx context.Context, zoneIdentifier string, body ZoneWaitingRoomWaitingRoomNewWaitingRoomParams, opts ...option.RequestOption) (res *SingleWaitingRoomResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists waiting rooms.
func (r *ZoneWaitingRoomService) WaitingRoomListWaitingRooms(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneWaitingRoomWaitingRoomListWaitingRoomsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SingleWaitingRoomResponse struct {
	Result SingleWaitingRoomResponseResult `json:"result"`
	JSON   singleWaitingRoomResponseJSON   `json:"-"`
}

// singleWaitingRoomResponseJSON contains the JSON metadata for the struct
// [SingleWaitingRoomResponse]
type singleWaitingRoomResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleWaitingRoomResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleWaitingRoomResponseResult struct {
	ID interface{} `json:"id"`
	// Only available for the Waiting Room Advanced subscription. Additional hostname
	// and path combinations to which this waiting room will be applied. There is an
	// implied wildcard at the end of the path. The hostname and path combination must
	// be unique to this and all other waiting rooms.
	AdditionalRoutes []SingleWaitingRoomResponseResultAdditionalRoute `json:"additional_routes"`
	// Configures cookie attributes for the waiting room cookie. This encrypted cookie
	// stores a user's status in the waiting room, such as queue position.
	CookieAttributes SingleWaitingRoomResponseResultCookieAttributes `json:"cookie_attributes"`
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
	DefaultTemplateLanguage SingleWaitingRoomResponseResultDefaultTemplateLanguage `json:"default_template_language"`
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
	QueueingMethod SingleWaitingRoomResponseResultQueueingMethod `json:"queueing_method"`
	// HTTP status code returned to a user while in the queue.
	QueueingStatusCode SingleWaitingRoomResponseResultQueueingStatusCode `json:"queueing_status_code"`
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
	TotalActiveUsers int64                               `json:"total_active_users"`
	JSON             singleWaitingRoomResponseResultJSON `json:"-"`
}

// singleWaitingRoomResponseResultJSON contains the JSON metadata for the struct
// [SingleWaitingRoomResponseResult]
type singleWaitingRoomResponseResultJSON struct {
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

func (r *SingleWaitingRoomResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleWaitingRoomResponseResultAdditionalRoute struct {
	// The hostname to which this waiting room will be applied (no wildcards). The
	// hostname must be the primary domain, subdomain, or custom hostname (if using SSL
	// for SaaS) of this zone. Please do not include the scheme (http:// or https://).
	Host string `json:"host"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path string                                             `json:"path"`
	JSON singleWaitingRoomResponseResultAdditionalRouteJSON `json:"-"`
}

// singleWaitingRoomResponseResultAdditionalRouteJSON contains the JSON metadata
// for the struct [SingleWaitingRoomResponseResultAdditionalRoute]
type singleWaitingRoomResponseResultAdditionalRouteJSON struct {
	Host        apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleWaitingRoomResponseResultAdditionalRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures cookie attributes for the waiting room cookie. This encrypted cookie
// stores a user's status in the waiting room, such as queue position.
type SingleWaitingRoomResponseResultCookieAttributes struct {
	// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
	// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
	// Note that when using value `none`, the secure attribute cannot be set to
	// `never`.
	Samesite SingleWaitingRoomResponseResultCookieAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on the waiting room cookie. Value `always`
	// indicates that the Secure attribute will be set in the Set-Cookie header,
	// `never` indicates that the Secure attribute will not be set, and `auto` will set
	// the Secure attribute depending if **Always Use HTTPS** is enabled.
	Secure SingleWaitingRoomResponseResultCookieAttributesSecure `json:"secure"`
	JSON   singleWaitingRoomResponseResultCookieAttributesJSON   `json:"-"`
}

// singleWaitingRoomResponseResultCookieAttributesJSON contains the JSON metadata
// for the struct [SingleWaitingRoomResponseResultCookieAttributes]
type singleWaitingRoomResponseResultCookieAttributesJSON struct {
	Samesite    apijson.Field
	Secure      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleWaitingRoomResponseResultCookieAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
// Note that when using value `none`, the secure attribute cannot be set to
// `never`.
type SingleWaitingRoomResponseResultCookieAttributesSamesite string

const (
	SingleWaitingRoomResponseResultCookieAttributesSamesiteAuto   SingleWaitingRoomResponseResultCookieAttributesSamesite = "auto"
	SingleWaitingRoomResponseResultCookieAttributesSamesiteLax    SingleWaitingRoomResponseResultCookieAttributesSamesite = "lax"
	SingleWaitingRoomResponseResultCookieAttributesSamesiteNone   SingleWaitingRoomResponseResultCookieAttributesSamesite = "none"
	SingleWaitingRoomResponseResultCookieAttributesSamesiteStrict SingleWaitingRoomResponseResultCookieAttributesSamesite = "strict"
)

// Configures the Secure attribute on the waiting room cookie. Value `always`
// indicates that the Secure attribute will be set in the Set-Cookie header,
// `never` indicates that the Secure attribute will not be set, and `auto` will set
// the Secure attribute depending if **Always Use HTTPS** is enabled.
type SingleWaitingRoomResponseResultCookieAttributesSecure string

const (
	SingleWaitingRoomResponseResultCookieAttributesSecureAuto   SingleWaitingRoomResponseResultCookieAttributesSecure = "auto"
	SingleWaitingRoomResponseResultCookieAttributesSecureAlways SingleWaitingRoomResponseResultCookieAttributesSecure = "always"
	SingleWaitingRoomResponseResultCookieAttributesSecureNever  SingleWaitingRoomResponseResultCookieAttributesSecure = "never"
)

// The language of the default page template. If no default_template_language is
// provided, then `en-US` (English) will be used.
type SingleWaitingRoomResponseResultDefaultTemplateLanguage string

const (
	SingleWaitingRoomResponseResultDefaultTemplateLanguageEnUs SingleWaitingRoomResponseResultDefaultTemplateLanguage = "en-US"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageEsEs SingleWaitingRoomResponseResultDefaultTemplateLanguage = "es-ES"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageDeDe SingleWaitingRoomResponseResultDefaultTemplateLanguage = "de-DE"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageFrFr SingleWaitingRoomResponseResultDefaultTemplateLanguage = "fr-FR"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageItIt SingleWaitingRoomResponseResultDefaultTemplateLanguage = "it-IT"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageJaJp SingleWaitingRoomResponseResultDefaultTemplateLanguage = "ja-JP"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageKoKr SingleWaitingRoomResponseResultDefaultTemplateLanguage = "ko-KR"
	SingleWaitingRoomResponseResultDefaultTemplateLanguagePtBr SingleWaitingRoomResponseResultDefaultTemplateLanguage = "pt-BR"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageZhCn SingleWaitingRoomResponseResultDefaultTemplateLanguage = "zh-CN"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageZhTw SingleWaitingRoomResponseResultDefaultTemplateLanguage = "zh-TW"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageNlNl SingleWaitingRoomResponseResultDefaultTemplateLanguage = "nl-NL"
	SingleWaitingRoomResponseResultDefaultTemplateLanguagePlPl SingleWaitingRoomResponseResultDefaultTemplateLanguage = "pl-PL"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageIDID SingleWaitingRoomResponseResultDefaultTemplateLanguage = "id-ID"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageTrTr SingleWaitingRoomResponseResultDefaultTemplateLanguage = "tr-TR"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageArEg SingleWaitingRoomResponseResultDefaultTemplateLanguage = "ar-EG"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageRuRu SingleWaitingRoomResponseResultDefaultTemplateLanguage = "ru-RU"
	SingleWaitingRoomResponseResultDefaultTemplateLanguageFaIr SingleWaitingRoomResponseResultDefaultTemplateLanguage = "fa-IR"
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
type SingleWaitingRoomResponseResultQueueingMethod string

const (
	SingleWaitingRoomResponseResultQueueingMethodFifo        SingleWaitingRoomResponseResultQueueingMethod = "fifo"
	SingleWaitingRoomResponseResultQueueingMethodRandom      SingleWaitingRoomResponseResultQueueingMethod = "random"
	SingleWaitingRoomResponseResultQueueingMethodPassthrough SingleWaitingRoomResponseResultQueueingMethod = "passthrough"
	SingleWaitingRoomResponseResultQueueingMethodReject      SingleWaitingRoomResponseResultQueueingMethod = "reject"
)

// HTTP status code returned to a user while in the queue.
type SingleWaitingRoomResponseResultQueueingStatusCode int64

const (
	SingleWaitingRoomResponseResultQueueingStatusCode200 SingleWaitingRoomResponseResultQueueingStatusCode = 200
	SingleWaitingRoomResponseResultQueueingStatusCode202 SingleWaitingRoomResponseResultQueueingStatusCode = 202
	SingleWaitingRoomResponseResultQueueingStatusCode429 SingleWaitingRoomResponseResultQueueingStatusCode = 429
)

type ZoneWaitingRoomDeleteResponse struct {
	Result ZoneWaitingRoomDeleteResponseResult `json:"result"`
	JSON   zoneWaitingRoomDeleteResponseJSON   `json:"-"`
}

// zoneWaitingRoomDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneWaitingRoomDeleteResponse]
type zoneWaitingRoomDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomDeleteResponseResult struct {
	ID   interface{}                             `json:"id"`
	JSON zoneWaitingRoomDeleteResponseResultJSON `json:"-"`
}

// zoneWaitingRoomDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomDeleteResponseResult]
type zoneWaitingRoomDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomWaitingRoomListWaitingRoomsResponse struct {
	Errors     []ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseError    `json:"errors"`
	Messages   []ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseMessage  `json:"messages"`
	Result     []ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResult   `json:"result"`
	ResultInfo ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseSuccess `json:"success"`
	JSON    zoneWaitingRoomWaitingRoomListWaitingRoomsResponseJSON    `json:"-"`
}

// zoneWaitingRoomWaitingRoomListWaitingRoomsResponseJSON contains the JSON
// metadata for the struct [ZoneWaitingRoomWaitingRoomListWaitingRoomsResponse]
type zoneWaitingRoomWaitingRoomListWaitingRoomsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomWaitingRoomListWaitingRoomsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseError struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zoneWaitingRoomWaitingRoomListWaitingRoomsResponseErrorJSON `json:"-"`
}

// zoneWaitingRoomWaitingRoomListWaitingRoomsResponseErrorJSON contains the JSON
// metadata for the struct
// [ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseError]
type zoneWaitingRoomWaitingRoomListWaitingRoomsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseMessage struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zoneWaitingRoomWaitingRoomListWaitingRoomsResponseMessageJSON `json:"-"`
}

// zoneWaitingRoomWaitingRoomListWaitingRoomsResponseMessageJSON contains the JSON
// metadata for the struct
// [ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseMessage]
type zoneWaitingRoomWaitingRoomListWaitingRoomsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResult struct {
	ID interface{} `json:"id"`
	// Only available for the Waiting Room Advanced subscription. Additional hostname
	// and path combinations to which this waiting room will be applied. There is an
	// implied wildcard at the end of the path. The hostname and path combination must
	// be unique to this and all other waiting rooms.
	AdditionalRoutes []ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultAdditionalRoute `json:"additional_routes"`
	// Configures cookie attributes for the waiting room cookie. This encrypted cookie
	// stores a user's status in the waiting room, such as queue position.
	CookieAttributes ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributes `json:"cookie_attributes"`
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
	DefaultTemplateLanguage ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage `json:"default_template_language"`
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
	QueueingMethod ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingMethod `json:"queueing_method"`
	// HTTP status code returned to a user while in the queue.
	QueueingStatusCode ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingStatusCode `json:"queueing_status_code"`
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
	TotalActiveUsers int64                                                        `json:"total_active_users"`
	JSON             zoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultJSON `json:"-"`
}

// zoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultJSON contains the JSON
// metadata for the struct
// [ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResult]
type zoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultJSON struct {
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

func (r *ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultAdditionalRoute struct {
	// The hostname to which this waiting room will be applied (no wildcards). The
	// hostname must be the primary domain, subdomain, or custom hostname (if using SSL
	// for SaaS) of this zone. Please do not include the scheme (http:// or https://).
	Host string `json:"host"`
	// Sets the path within the host to enable the waiting room on. The waiting room
	// will be enabled for all subpaths as well. If there are two waiting rooms on the
	// same subpath, the waiting room for the most specific path will be chosen.
	// Wildcards and query parameters are not supported.
	Path string                                                                      `json:"path"`
	JSON zoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultAdditionalRouteJSON `json:"-"`
}

// zoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultAdditionalRouteJSON
// contains the JSON metadata for the struct
// [ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultAdditionalRoute]
type zoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultAdditionalRouteJSON struct {
	Host        apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultAdditionalRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures cookie attributes for the waiting room cookie. This encrypted cookie
// stores a user's status in the waiting room, such as queue position.
type ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributes struct {
	// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
	// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
	// Note that when using value `none`, the secure attribute cannot be set to
	// `never`.
	Samesite ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on the waiting room cookie. Value `always`
	// indicates that the Secure attribute will be set in the Set-Cookie header,
	// `never` indicates that the Secure attribute will not be set, and `auto` will set
	// the Secure attribute depending if **Always Use HTTPS** is enabled.
	Secure ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSecure `json:"secure"`
	JSON   zoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesJSON   `json:"-"`
}

// zoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesJSON
// contains the JSON metadata for the struct
// [ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributes]
type zoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesJSON struct {
	Samesite    apijson.Field
	Secure      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
// Note that when using value `none`, the secure attribute cannot be set to
// `never`.
type ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSamesite string

const (
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSamesiteAuto   ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSamesite = "auto"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSamesiteLax    ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSamesite = "lax"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSamesiteNone   ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSamesite = "none"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSamesiteStrict ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSamesite = "strict"
)

// Configures the Secure attribute on the waiting room cookie. Value `always`
// indicates that the Secure attribute will be set in the Set-Cookie header,
// `never` indicates that the Secure attribute will not be set, and `auto` will set
// the Secure attribute depending if **Always Use HTTPS** is enabled.
type ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSecure string

const (
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSecureAuto   ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSecure = "auto"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSecureAlways ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSecure = "always"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSecureNever  ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultCookieAttributesSecure = "never"
)

// The language of the default page template. If no default_template_language is
// provided, then `en-US` (English) will be used.
type ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage string

const (
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageEnUs ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "en-US"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageEsEs ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "es-ES"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageDeDe ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "de-DE"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageFrFr ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "fr-FR"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageItIt ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "it-IT"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageJaJp ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "ja-JP"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageKoKr ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "ko-KR"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguagePtBr ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "pt-BR"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageZhCn ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "zh-CN"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageZhTw ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "zh-TW"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageNlNl ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "nl-NL"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguagePlPl ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "pl-PL"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageIDID ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "id-ID"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageTrTr ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "tr-TR"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageArEg ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "ar-EG"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageRuRu ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "ru-RU"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguageFaIr ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultDefaultTemplateLanguage = "fa-IR"
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
type ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingMethod string

const (
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingMethodFifo        ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingMethod = "fifo"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingMethodRandom      ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingMethod = "random"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingMethodPassthrough ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingMethod = "passthrough"
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingMethodReject      ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingMethod = "reject"
)

// HTTP status code returned to a user while in the queue.
type ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingStatusCode int64

const (
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingStatusCode200 ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingStatusCode = 200
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingStatusCode202 ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingStatusCode = 202
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingStatusCode429 ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultQueueingStatusCode = 429
)

type ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                          `json:"total_count"`
	JSON       zoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultInfoJSON `json:"-"`
}

// zoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultInfoJSON contains the
// JSON metadata for the struct
// [ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultInfo]
type zoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseSuccess bool

const (
	ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseSuccessTrue ZoneWaitingRoomWaitingRoomListWaitingRoomsResponseSuccess = true
)

type ZoneWaitingRoomUpdateParams struct {
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
	AdditionalRoutes param.Field[[]ZoneWaitingRoomUpdateParamsAdditionalRoute] `json:"additional_routes"`
	// Configures cookie attributes for the waiting room cookie. This encrypted cookie
	// stores a user's status in the waiting room, such as queue position.
	CookieAttributes param.Field[ZoneWaitingRoomUpdateParamsCookieAttributes] `json:"cookie_attributes"`
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
	DefaultTemplateLanguage param.Field[ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage] `json:"default_template_language"`
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
	QueueingMethod param.Field[ZoneWaitingRoomUpdateParamsQueueingMethod] `json:"queueing_method"`
	// HTTP status code returned to a user while in the queue.
	QueueingStatusCode param.Field[ZoneWaitingRoomUpdateParamsQueueingStatusCode] `json:"queueing_status_code"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to
	// the route. If a user is not seen by Cloudflare again in that time period, they
	// will be treated as a new user that visits the route.
	SessionDuration param.Field[int64] `json:"session_duration"`
	// Suspends or allows traffic going to the waiting room. If set to `true`, the
	// traffic will not go to the waiting room.
	Suspended param.Field[bool] `json:"suspended"`
}

func (r ZoneWaitingRoomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneWaitingRoomUpdateParamsAdditionalRoute struct {
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

func (r ZoneWaitingRoomUpdateParamsAdditionalRoute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures cookie attributes for the waiting room cookie. This encrypted cookie
// stores a user's status in the waiting room, such as queue position.
type ZoneWaitingRoomUpdateParamsCookieAttributes struct {
	// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
	// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
	// Note that when using value `none`, the secure attribute cannot be set to
	// `never`.
	Samesite param.Field[ZoneWaitingRoomUpdateParamsCookieAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on the waiting room cookie. Value `always`
	// indicates that the Secure attribute will be set in the Set-Cookie header,
	// `never` indicates that the Secure attribute will not be set, and `auto` will set
	// the Secure attribute depending if **Always Use HTTPS** is enabled.
	Secure param.Field[ZoneWaitingRoomUpdateParamsCookieAttributesSecure] `json:"secure"`
}

func (r ZoneWaitingRoomUpdateParamsCookieAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
// Note that when using value `none`, the secure attribute cannot be set to
// `never`.
type ZoneWaitingRoomUpdateParamsCookieAttributesSamesite string

const (
	ZoneWaitingRoomUpdateParamsCookieAttributesSamesiteAuto   ZoneWaitingRoomUpdateParamsCookieAttributesSamesite = "auto"
	ZoneWaitingRoomUpdateParamsCookieAttributesSamesiteLax    ZoneWaitingRoomUpdateParamsCookieAttributesSamesite = "lax"
	ZoneWaitingRoomUpdateParamsCookieAttributesSamesiteNone   ZoneWaitingRoomUpdateParamsCookieAttributesSamesite = "none"
	ZoneWaitingRoomUpdateParamsCookieAttributesSamesiteStrict ZoneWaitingRoomUpdateParamsCookieAttributesSamesite = "strict"
)

// Configures the Secure attribute on the waiting room cookie. Value `always`
// indicates that the Secure attribute will be set in the Set-Cookie header,
// `never` indicates that the Secure attribute will not be set, and `auto` will set
// the Secure attribute depending if **Always Use HTTPS** is enabled.
type ZoneWaitingRoomUpdateParamsCookieAttributesSecure string

const (
	ZoneWaitingRoomUpdateParamsCookieAttributesSecureAuto   ZoneWaitingRoomUpdateParamsCookieAttributesSecure = "auto"
	ZoneWaitingRoomUpdateParamsCookieAttributesSecureAlways ZoneWaitingRoomUpdateParamsCookieAttributesSecure = "always"
	ZoneWaitingRoomUpdateParamsCookieAttributesSecureNever  ZoneWaitingRoomUpdateParamsCookieAttributesSecure = "never"
)

// The language of the default page template. If no default_template_language is
// provided, then `en-US` (English) will be used.
type ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage string

const (
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageEnUs ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "en-US"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageEsEs ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "es-ES"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageDeDe ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "de-DE"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageFrFr ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "fr-FR"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageItIt ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "it-IT"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageJaJp ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "ja-JP"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageKoKr ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "ko-KR"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguagePtBr ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "pt-BR"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageZhCn ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "zh-CN"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageZhTw ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "zh-TW"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageNlNl ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "nl-NL"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguagePlPl ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "pl-PL"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageIDID ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "id-ID"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageTrTr ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "tr-TR"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageArEg ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "ar-EG"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageRuRu ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "ru-RU"
	ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageFaIr ZoneWaitingRoomUpdateParamsDefaultTemplateLanguage = "fa-IR"
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
type ZoneWaitingRoomUpdateParamsQueueingMethod string

const (
	ZoneWaitingRoomUpdateParamsQueueingMethodFifo        ZoneWaitingRoomUpdateParamsQueueingMethod = "fifo"
	ZoneWaitingRoomUpdateParamsQueueingMethodRandom      ZoneWaitingRoomUpdateParamsQueueingMethod = "random"
	ZoneWaitingRoomUpdateParamsQueueingMethodPassthrough ZoneWaitingRoomUpdateParamsQueueingMethod = "passthrough"
	ZoneWaitingRoomUpdateParamsQueueingMethodReject      ZoneWaitingRoomUpdateParamsQueueingMethod = "reject"
)

// HTTP status code returned to a user while in the queue.
type ZoneWaitingRoomUpdateParamsQueueingStatusCode int64

const (
	ZoneWaitingRoomUpdateParamsQueueingStatusCode200 ZoneWaitingRoomUpdateParamsQueueingStatusCode = 200
	ZoneWaitingRoomUpdateParamsQueueingStatusCode202 ZoneWaitingRoomUpdateParamsQueueingStatusCode = 202
	ZoneWaitingRoomUpdateParamsQueueingStatusCode429 ZoneWaitingRoomUpdateParamsQueueingStatusCode = 429
)

type ZoneWaitingRoomPatchParams struct {
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
	AdditionalRoutes param.Field[[]ZoneWaitingRoomPatchParamsAdditionalRoute] `json:"additional_routes"`
	// Configures cookie attributes for the waiting room cookie. This encrypted cookie
	// stores a user's status in the waiting room, such as queue position.
	CookieAttributes param.Field[ZoneWaitingRoomPatchParamsCookieAttributes] `json:"cookie_attributes"`
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
	DefaultTemplateLanguage param.Field[ZoneWaitingRoomPatchParamsDefaultTemplateLanguage] `json:"default_template_language"`
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
	QueueingMethod param.Field[ZoneWaitingRoomPatchParamsQueueingMethod] `json:"queueing_method"`
	// HTTP status code returned to a user while in the queue.
	QueueingStatusCode param.Field[ZoneWaitingRoomPatchParamsQueueingStatusCode] `json:"queueing_status_code"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to
	// the route. If a user is not seen by Cloudflare again in that time period, they
	// will be treated as a new user that visits the route.
	SessionDuration param.Field[int64] `json:"session_duration"`
	// Suspends or allows traffic going to the waiting room. If set to `true`, the
	// traffic will not go to the waiting room.
	Suspended param.Field[bool] `json:"suspended"`
}

func (r ZoneWaitingRoomPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneWaitingRoomPatchParamsAdditionalRoute struct {
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

func (r ZoneWaitingRoomPatchParamsAdditionalRoute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures cookie attributes for the waiting room cookie. This encrypted cookie
// stores a user's status in the waiting room, such as queue position.
type ZoneWaitingRoomPatchParamsCookieAttributes struct {
	// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
	// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
	// Note that when using value `none`, the secure attribute cannot be set to
	// `never`.
	Samesite param.Field[ZoneWaitingRoomPatchParamsCookieAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on the waiting room cookie. Value `always`
	// indicates that the Secure attribute will be set in the Set-Cookie header,
	// `never` indicates that the Secure attribute will not be set, and `auto` will set
	// the Secure attribute depending if **Always Use HTTPS** is enabled.
	Secure param.Field[ZoneWaitingRoomPatchParamsCookieAttributesSecure] `json:"secure"`
}

func (r ZoneWaitingRoomPatchParamsCookieAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
// Note that when using value `none`, the secure attribute cannot be set to
// `never`.
type ZoneWaitingRoomPatchParamsCookieAttributesSamesite string

const (
	ZoneWaitingRoomPatchParamsCookieAttributesSamesiteAuto   ZoneWaitingRoomPatchParamsCookieAttributesSamesite = "auto"
	ZoneWaitingRoomPatchParamsCookieAttributesSamesiteLax    ZoneWaitingRoomPatchParamsCookieAttributesSamesite = "lax"
	ZoneWaitingRoomPatchParamsCookieAttributesSamesiteNone   ZoneWaitingRoomPatchParamsCookieAttributesSamesite = "none"
	ZoneWaitingRoomPatchParamsCookieAttributesSamesiteStrict ZoneWaitingRoomPatchParamsCookieAttributesSamesite = "strict"
)

// Configures the Secure attribute on the waiting room cookie. Value `always`
// indicates that the Secure attribute will be set in the Set-Cookie header,
// `never` indicates that the Secure attribute will not be set, and `auto` will set
// the Secure attribute depending if **Always Use HTTPS** is enabled.
type ZoneWaitingRoomPatchParamsCookieAttributesSecure string

const (
	ZoneWaitingRoomPatchParamsCookieAttributesSecureAuto   ZoneWaitingRoomPatchParamsCookieAttributesSecure = "auto"
	ZoneWaitingRoomPatchParamsCookieAttributesSecureAlways ZoneWaitingRoomPatchParamsCookieAttributesSecure = "always"
	ZoneWaitingRoomPatchParamsCookieAttributesSecureNever  ZoneWaitingRoomPatchParamsCookieAttributesSecure = "never"
)

// The language of the default page template. If no default_template_language is
// provided, then `en-US` (English) will be used.
type ZoneWaitingRoomPatchParamsDefaultTemplateLanguage string

const (
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageEnUs ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "en-US"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageEsEs ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "es-ES"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageDeDe ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "de-DE"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageFrFr ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "fr-FR"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageItIt ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "it-IT"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageJaJp ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "ja-JP"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageKoKr ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "ko-KR"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguagePtBr ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "pt-BR"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageZhCn ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "zh-CN"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageZhTw ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "zh-TW"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageNlNl ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "nl-NL"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguagePlPl ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "pl-PL"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageIDID ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "id-ID"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageTrTr ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "tr-TR"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageArEg ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "ar-EG"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageRuRu ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "ru-RU"
	ZoneWaitingRoomPatchParamsDefaultTemplateLanguageFaIr ZoneWaitingRoomPatchParamsDefaultTemplateLanguage = "fa-IR"
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
type ZoneWaitingRoomPatchParamsQueueingMethod string

const (
	ZoneWaitingRoomPatchParamsQueueingMethodFifo        ZoneWaitingRoomPatchParamsQueueingMethod = "fifo"
	ZoneWaitingRoomPatchParamsQueueingMethodRandom      ZoneWaitingRoomPatchParamsQueueingMethod = "random"
	ZoneWaitingRoomPatchParamsQueueingMethodPassthrough ZoneWaitingRoomPatchParamsQueueingMethod = "passthrough"
	ZoneWaitingRoomPatchParamsQueueingMethodReject      ZoneWaitingRoomPatchParamsQueueingMethod = "reject"
)

// HTTP status code returned to a user while in the queue.
type ZoneWaitingRoomPatchParamsQueueingStatusCode int64

const (
	ZoneWaitingRoomPatchParamsQueueingStatusCode200 ZoneWaitingRoomPatchParamsQueueingStatusCode = 200
	ZoneWaitingRoomPatchParamsQueueingStatusCode202 ZoneWaitingRoomPatchParamsQueueingStatusCode = 202
	ZoneWaitingRoomPatchParamsQueueingStatusCode429 ZoneWaitingRoomPatchParamsQueueingStatusCode = 429
)

type ZoneWaitingRoomWaitingRoomNewWaitingRoomParams struct {
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
	AdditionalRoutes param.Field[[]ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsAdditionalRoute] `json:"additional_routes"`
	// Configures cookie attributes for the waiting room cookie. This encrypted cookie
	// stores a user's status in the waiting room, such as queue position.
	CookieAttributes param.Field[ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributes] `json:"cookie_attributes"`
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
	DefaultTemplateLanguage param.Field[ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage] `json:"default_template_language"`
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
	QueueingMethod param.Field[ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingMethod] `json:"queueing_method"`
	// HTTP status code returned to a user while in the queue.
	QueueingStatusCode param.Field[ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingStatusCode] `json:"queueing_status_code"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to
	// the route. If a user is not seen by Cloudflare again in that time period, they
	// will be treated as a new user that visits the route.
	SessionDuration param.Field[int64] `json:"session_duration"`
	// Suspends or allows traffic going to the waiting room. If set to `true`, the
	// traffic will not go to the waiting room.
	Suspended param.Field[bool] `json:"suspended"`
}

func (r ZoneWaitingRoomWaitingRoomNewWaitingRoomParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsAdditionalRoute struct {
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

func (r ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsAdditionalRoute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures cookie attributes for the waiting room cookie. This encrypted cookie
// stores a user's status in the waiting room, such as queue position.
type ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributes struct {
	// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
	// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
	// Note that when using value `none`, the secure attribute cannot be set to
	// `never`.
	Samesite param.Field[ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on the waiting room cookie. Value `always`
	// indicates that the Secure attribute will be set in the Set-Cookie header,
	// `never` indicates that the Secure attribute will not be set, and `auto` will set
	// the Secure attribute depending if **Always Use HTTPS** is enabled.
	Secure param.Field[ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSecure] `json:"secure"`
}

func (r ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on the waiting room cookie. Value `auto` will
// be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled.
// Note that when using value `none`, the secure attribute cannot be set to
// `never`.
type ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSamesite string

const (
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSamesiteAuto   ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSamesite = "auto"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSamesiteLax    ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSamesite = "lax"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSamesiteNone   ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSamesite = "none"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSamesiteStrict ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSamesite = "strict"
)

// Configures the Secure attribute on the waiting room cookie. Value `always`
// indicates that the Secure attribute will be set in the Set-Cookie header,
// `never` indicates that the Secure attribute will not be set, and `auto` will set
// the Secure attribute depending if **Always Use HTTPS** is enabled.
type ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSecure string

const (
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSecureAuto   ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSecure = "auto"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSecureAlways ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSecure = "always"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSecureNever  ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSecure = "never"
)

// The language of the default page template. If no default_template_language is
// provided, then `en-US` (English) will be used.
type ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage string

const (
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageEnUs ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "en-US"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageEsEs ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "es-ES"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageDeDe ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "de-DE"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageFrFr ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "fr-FR"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageItIt ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "it-IT"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageJaJp ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "ja-JP"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageKoKr ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "ko-KR"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguagePtBr ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "pt-BR"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageZhCn ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "zh-CN"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageZhTw ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "zh-TW"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageNlNl ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "nl-NL"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguagePlPl ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "pl-PL"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageIDID ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "id-ID"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageTrTr ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "tr-TR"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageArEg ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "ar-EG"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageRuRu ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "ru-RU"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageFaIr ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguage = "fa-IR"
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
type ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingMethod string

const (
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingMethodFifo        ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingMethod = "fifo"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingMethodRandom      ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingMethod = "random"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingMethodPassthrough ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingMethod = "passthrough"
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingMethodReject      ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingMethod = "reject"
)

// HTTP status code returned to a user while in the queue.
type ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingStatusCode int64

const (
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingStatusCode200 ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingStatusCode = 200
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingStatusCode202 ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingStatusCode = 202
	ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingStatusCode429 ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingStatusCode = 429
)
