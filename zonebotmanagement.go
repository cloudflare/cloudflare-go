// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneBotManagementService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneBotManagementService] method
// instead.
type ZoneBotManagementService struct {
	Options []option.RequestOption
}

// NewZoneBotManagementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneBotManagementService(opts ...option.RequestOption) (r *ZoneBotManagementService) {
	r = &ZoneBotManagementService{}
	r.Options = opts
	return
}

// Retrieve a zone's Bot Management Config
func (r *ZoneBotManagementService) Get(ctx context.Context, zoneIdentifier string, query ZoneBotManagementGetParams, opts ...option.RequestOption) (res *ZoneBotManagementGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/bot_management", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates the Bot Management configuration for a zone.
//
// This API is used to update:
//
// - **Bot Fight Mode**
// - **Super Bot Fight Mode**
// - **Bot Management for Enterprise**
//
// See [Bot Plans](https://developers.cloudflare.com/bots/plans/) for more
// information on the different plans
func (r *ZoneBotManagementService) Update(ctx context.Context, zoneIdentifier string, params ZoneBotManagementUpdateParams, opts ...option.RequestOption) (res *ZoneBotManagementUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/bot_management", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type ZoneBotManagementGetResponse struct {
	Errors   []ZoneBotManagementGetResponseError   `json:"errors"`
	Messages []ZoneBotManagementGetResponseMessage `json:"messages"`
	Result   ZoneBotManagementGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneBotManagementGetResponseSuccess `json:"success"`
	JSON    zoneBotManagementGetResponseJSON    `json:"-"`
}

// zoneBotManagementGetResponseJSON contains the JSON metadata for the struct
// [ZoneBotManagementGetResponse]
type zoneBotManagementGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneBotManagementGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneBotManagementGetResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneBotManagementGetResponseErrorJSON `json:"-"`
}

// zoneBotManagementGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneBotManagementGetResponseError]
type zoneBotManagementGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneBotManagementGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneBotManagementGetResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneBotManagementGetResponseMessageJSON `json:"-"`
}

// zoneBotManagementGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneBotManagementGetResponseMessage]
type zoneBotManagementGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneBotManagementGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ZoneBotManagementGetResponseResultEsic0axpBotFightModeConfig],
// [ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfig],
// [ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfig] or
// [ZoneBotManagementGetResponseResultEsic0axpBmSubscriptionConfig].
type ZoneBotManagementGetResponseResult interface {
	implementsZoneBotManagementGetResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneBotManagementGetResponseResult)(nil)).Elem(), "")
}

type ZoneBotManagementGetResponseResultEsic0axpBotFightModeConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                             `json:"using_latest_model"`
	JSON             zoneBotManagementGetResponseResultEsic0axpBotFightModeConfigJSON `json:"-"`
}

// zoneBotManagementGetResponseResultEsic0axpBotFightModeConfigJSON contains the
// JSON metadata for the struct
// [ZoneBotManagementGetResponseResultEsic0axpBotFightModeConfig]
type zoneBotManagementGetResponseResultEsic0axpBotFightModeConfigJSON struct {
	EnableJs         apijson.Field
	FightMode        apijson.Field
	UsingLatestModel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneBotManagementGetResponseResultEsic0axpBotFightModeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneBotManagementGetResponseResultEsic0axpBotFightModeConfig) implementsZoneBotManagementGetResponseResult() {
}

type ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                               `json:"using_latest_model"`
	JSON             zoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigJSON `json:"-"`
}

// zoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigJSON contains the
// JSON metadata for the struct
// [ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfig]
type zoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigJSON struct {
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfig) implementsZoneBotManagementGetResponseResult() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated string

const (
	ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomatedAllow            ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated = "allow"
	ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomatedBlock            ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated = "block"
	ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomatedManagedChallenge ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmVerifiedBots string

const (
	ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmVerifiedBotsAllow ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmVerifiedBots = "allow"
	ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmVerifiedBotsBlock ZoneBotManagementGetResponseResultEsic0axpSbfmDefinitelyConfigSbfmVerifiedBots = "block"
)

type ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                           `json:"using_latest_model"`
	JSON             zoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigJSON `json:"-"`
}

// zoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigJSON contains the JSON
// metadata for the struct
// [ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfig]
type zoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigJSON struct {
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmLikelyAutomated          apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfig) implementsZoneBotManagementGetResponseResult() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated string

const (
	ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomatedAllow            ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated = "allow"
	ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomatedBlock            ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated = "block"
	ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomatedManagedChallenge ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomated string

const (
	ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomatedAllow            ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomated = "allow"
	ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomatedBlock            ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomated = "block"
	ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmVerifiedBots string

const (
	ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmVerifiedBotsAllow ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmVerifiedBots = "allow"
	ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmVerifiedBotsBlock ZoneBotManagementGetResponseResultEsic0axpSbfmLikelyConfigSbfmVerifiedBots = "block"
)

type ZoneBotManagementGetResponseResultEsic0axpBmSubscriptionConfig struct {
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel bool `json:"auto_update_model"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore bool `json:"suppress_session_score"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                               `json:"using_latest_model"`
	JSON             zoneBotManagementGetResponseResultEsic0axpBmSubscriptionConfigJSON `json:"-"`
}

// zoneBotManagementGetResponseResultEsic0axpBmSubscriptionConfigJSON contains the
// JSON metadata for the struct
// [ZoneBotManagementGetResponseResultEsic0axpBmSubscriptionConfig]
type zoneBotManagementGetResponseResultEsic0axpBmSubscriptionConfigJSON struct {
	AutoUpdateModel      apijson.Field
	EnableJs             apijson.Field
	SuppressSessionScore apijson.Field
	UsingLatestModel     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneBotManagementGetResponseResultEsic0axpBmSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneBotManagementGetResponseResultEsic0axpBmSubscriptionConfig) implementsZoneBotManagementGetResponseResult() {
}

// Whether the API call was successful
type ZoneBotManagementGetResponseSuccess bool

const (
	ZoneBotManagementGetResponseSuccessTrue ZoneBotManagementGetResponseSuccess = true
)

type ZoneBotManagementUpdateResponse struct {
	Errors   []ZoneBotManagementUpdateResponseError   `json:"errors"`
	Messages []ZoneBotManagementUpdateResponseMessage `json:"messages"`
	Result   ZoneBotManagementUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneBotManagementUpdateResponseSuccess `json:"success"`
	JSON    zoneBotManagementUpdateResponseJSON    `json:"-"`
}

// zoneBotManagementUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneBotManagementUpdateResponse]
type zoneBotManagementUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneBotManagementUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneBotManagementUpdateResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneBotManagementUpdateResponseErrorJSON `json:"-"`
}

// zoneBotManagementUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneBotManagementUpdateResponseError]
type zoneBotManagementUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneBotManagementUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneBotManagementUpdateResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneBotManagementUpdateResponseMessageJSON `json:"-"`
}

// zoneBotManagementUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneBotManagementUpdateResponseMessage]
type zoneBotManagementUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneBotManagementUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ZoneBotManagementUpdateResponseResultEsic0axpBotFightModeConfig],
// [ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfig],
// [ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfig] or
// [ZoneBotManagementUpdateResponseResultEsic0axpBmSubscriptionConfig].
type ZoneBotManagementUpdateResponseResult interface {
	implementsZoneBotManagementUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneBotManagementUpdateResponseResult)(nil)).Elem(), "")
}

type ZoneBotManagementUpdateResponseResultEsic0axpBotFightModeConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                                `json:"using_latest_model"`
	JSON             zoneBotManagementUpdateResponseResultEsic0axpBotFightModeConfigJSON `json:"-"`
}

// zoneBotManagementUpdateResponseResultEsic0axpBotFightModeConfigJSON contains the
// JSON metadata for the struct
// [ZoneBotManagementUpdateResponseResultEsic0axpBotFightModeConfig]
type zoneBotManagementUpdateResponseResultEsic0axpBotFightModeConfigJSON struct {
	EnableJs         apijson.Field
	FightMode        apijson.Field
	UsingLatestModel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneBotManagementUpdateResponseResultEsic0axpBotFightModeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneBotManagementUpdateResponseResultEsic0axpBotFightModeConfig) implementsZoneBotManagementUpdateResponseResult() {
}

type ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                                  `json:"using_latest_model"`
	JSON             zoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigJSON `json:"-"`
}

// zoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigJSON contains
// the JSON metadata for the struct
// [ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfig]
type zoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigJSON struct {
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfig) implementsZoneBotManagementUpdateResponseResult() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated string

const (
	ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomatedAllow            ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated = "allow"
	ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomatedBlock            ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated = "block"
	ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomatedManagedChallenge ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmVerifiedBots string

const (
	ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmVerifiedBotsAllow ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmVerifiedBots = "allow"
	ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmVerifiedBotsBlock ZoneBotManagementUpdateResponseResultEsic0axpSbfmDefinitelyConfigSbfmVerifiedBots = "block"
)

type ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                              `json:"using_latest_model"`
	JSON             zoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigJSON `json:"-"`
}

// zoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigJSON contains the
// JSON metadata for the struct
// [ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfig]
type zoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigJSON struct {
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmLikelyAutomated          apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfig) implementsZoneBotManagementUpdateResponseResult() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated string

const (
	ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomatedAllow            ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated = "allow"
	ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomatedBlock            ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated = "block"
	ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomatedManagedChallenge ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomated string

const (
	ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomatedAllow            ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomated = "allow"
	ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomatedBlock            ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomated = "block"
	ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmLikelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmVerifiedBots string

const (
	ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmVerifiedBotsAllow ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmVerifiedBots = "allow"
	ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmVerifiedBotsBlock ZoneBotManagementUpdateResponseResultEsic0axpSbfmLikelyConfigSbfmVerifiedBots = "block"
)

type ZoneBotManagementUpdateResponseResultEsic0axpBmSubscriptionConfig struct {
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel bool `json:"auto_update_model"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore bool `json:"suppress_session_score"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                                  `json:"using_latest_model"`
	JSON             zoneBotManagementUpdateResponseResultEsic0axpBmSubscriptionConfigJSON `json:"-"`
}

// zoneBotManagementUpdateResponseResultEsic0axpBmSubscriptionConfigJSON contains
// the JSON metadata for the struct
// [ZoneBotManagementUpdateResponseResultEsic0axpBmSubscriptionConfig]
type zoneBotManagementUpdateResponseResultEsic0axpBmSubscriptionConfigJSON struct {
	AutoUpdateModel      apijson.Field
	EnableJs             apijson.Field
	SuppressSessionScore apijson.Field
	UsingLatestModel     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneBotManagementUpdateResponseResultEsic0axpBmSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneBotManagementUpdateResponseResultEsic0axpBmSubscriptionConfig) implementsZoneBotManagementUpdateResponseResult() {
}

// Whether the API call was successful
type ZoneBotManagementUpdateResponseSuccess bool

const (
	ZoneBotManagementUpdateResponseSuccessTrue ZoneBotManagementUpdateResponseSuccess = true
)

type ZoneBotManagementGetParams struct {
	// Header used to control which version of the API to use. Note: Only the 2.0.0
	// version is documented. The older 1.0.0 version is deprecated and will soon be
	// removed.
	CloudflareVersion param.Field[ZoneBotManagementGetParamsCloudflareVersion] `header:"Cloudflare-Version"`
}

// Header used to control which version of the API to use. Note: Only the 2.0.0
// version is documented. The older 1.0.0 version is deprecated and will soon be
// removed.
type ZoneBotManagementGetParamsCloudflareVersion string

const (
	ZoneBotManagementGetParamsCloudflareVersion2_0_0 ZoneBotManagementGetParamsCloudflareVersion = "2.0.0"
	ZoneBotManagementGetParamsCloudflareVersion1_0_0 ZoneBotManagementGetParamsCloudflareVersion = "1.0.0"
)

// This interface is a union satisfied by one of the following:
// [ZoneBotManagementUpdateParamsEsic0axpBotFightModeConfig],
// [ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfig],
// [ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfig],
// [ZoneBotManagementUpdateParamsEsic0axpBmSubscriptionConfig].
type ZoneBotManagementUpdateParams interface {
	ImplementsZoneBotManagementUpdateParams()
}

type ZoneBotManagementUpdateParamsEsic0axpBotFightModeConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode param.Field[bool] `json:"fight_mode"`
	// Header used to control which version of the API to use. Note: Only the 2.0.0
	// version is documented. The older 1.0.0 version is deprecated and will soon be
	// removed.
	CloudflareVersion param.Field[ZoneBotManagementUpdateParamsEsic0axpBotFightModeConfigCloudflareVersion] `header:"Cloudflare-Version"`
}

func (r ZoneBotManagementUpdateParamsEsic0axpBotFightModeConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneBotManagementUpdateParamsEsic0axpBotFightModeConfig) ImplementsZoneBotManagementUpdateParams() {

}

// Header used to control which version of the API to use. Note: Only the 2.0.0
// version is documented. The older 1.0.0 version is deprecated and will soon be
// removed.
type ZoneBotManagementUpdateParamsEsic0axpBotFightModeConfigCloudflareVersion string

const (
	ZoneBotManagementUpdateParamsEsic0axpBotFightModeConfigCloudflareVersion2_0_0 ZoneBotManagementUpdateParamsEsic0axpBotFightModeConfigCloudflareVersion = "2.0.0"
	ZoneBotManagementUpdateParamsEsic0axpBotFightModeConfigCloudflareVersion1_0_0 ZoneBotManagementUpdateParamsEsic0axpBotFightModeConfigCloudflareVersion = "1.0.0"
)

type ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated param.Field[ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots param.Field[ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmVerifiedBots] `json:"sbfm_verified_bots"`
	// Header used to control which version of the API to use. Note: Only the 2.0.0
	// version is documented. The older 1.0.0 version is deprecated and will soon be
	// removed.
	CloudflareVersion param.Field[ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigCloudflareVersion] `header:"Cloudflare-Version"`
}

func (r ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfig) ImplementsZoneBotManagementUpdateParams() {

}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated string

const (
	ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomatedAllow            ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated = "allow"
	ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomatedBlock            ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated = "block"
	ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomatedManagedChallenge ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmVerifiedBots string

const (
	ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmVerifiedBotsAllow ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmVerifiedBots = "allow"
	ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmVerifiedBotsBlock ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigSbfmVerifiedBots = "block"
)

// Header used to control which version of the API to use. Note: Only the 2.0.0
// version is documented. The older 1.0.0 version is deprecated and will soon be
// removed.
type ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigCloudflareVersion string

const (
	ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigCloudflareVersion2_0_0 ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigCloudflareVersion = "2.0.0"
	ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigCloudflareVersion1_0_0 ZoneBotManagementUpdateParamsEsic0axpSbfmDefinitelyConfigCloudflareVersion = "1.0.0"
)

type ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated param.Field[ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated param.Field[ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmLikelyAutomated] `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots param.Field[ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmVerifiedBots] `json:"sbfm_verified_bots"`
	// Header used to control which version of the API to use. Note: Only the 2.0.0
	// version is documented. The older 1.0.0 version is deprecated and will soon be
	// removed.
	CloudflareVersion param.Field[ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigCloudflareVersion] `header:"Cloudflare-Version"`
}

func (r ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfig) ImplementsZoneBotManagementUpdateParams() {

}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated string

const (
	ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomatedAllow            ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated = "allow"
	ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomatedBlock            ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated = "block"
	ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomatedManagedChallenge ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmLikelyAutomated string

const (
	ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmLikelyAutomatedAllow            ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmLikelyAutomated = "allow"
	ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmLikelyAutomatedBlock            ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmLikelyAutomated = "block"
	ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmLikelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmVerifiedBots string

const (
	ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmVerifiedBotsAllow ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmVerifiedBots = "allow"
	ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmVerifiedBotsBlock ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigSbfmVerifiedBots = "block"
)

// Header used to control which version of the API to use. Note: Only the 2.0.0
// version is documented. The older 1.0.0 version is deprecated and will soon be
// removed.
type ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigCloudflareVersion string

const (
	ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigCloudflareVersion2_0_0 ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigCloudflareVersion = "2.0.0"
	ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigCloudflareVersion1_0_0 ZoneBotManagementUpdateParamsEsic0axpSbfmLikelyConfigCloudflareVersion = "1.0.0"
)

type ZoneBotManagementUpdateParamsEsic0axpBmSubscriptionConfig struct {
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel param.Field[bool] `json:"auto_update_model"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore param.Field[bool] `json:"suppress_session_score"`
	// Header used to control which version of the API to use. Note: Only the 2.0.0
	// version is documented. The older 1.0.0 version is deprecated and will soon be
	// removed.
	CloudflareVersion param.Field[ZoneBotManagementUpdateParamsEsic0axpBmSubscriptionConfigCloudflareVersion] `header:"Cloudflare-Version"`
}

func (r ZoneBotManagementUpdateParamsEsic0axpBmSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneBotManagementUpdateParamsEsic0axpBmSubscriptionConfig) ImplementsZoneBotManagementUpdateParams() {

}

// Header used to control which version of the API to use. Note: Only the 2.0.0
// version is documented. The older 1.0.0 version is deprecated and will soon be
// removed.
type ZoneBotManagementUpdateParamsEsic0axpBmSubscriptionConfigCloudflareVersion string

const (
	ZoneBotManagementUpdateParamsEsic0axpBmSubscriptionConfigCloudflareVersion2_0_0 ZoneBotManagementUpdateParamsEsic0axpBmSubscriptionConfigCloudflareVersion = "2.0.0"
	ZoneBotManagementUpdateParamsEsic0axpBmSubscriptionConfigCloudflareVersion1_0_0 ZoneBotManagementUpdateParamsEsic0axpBmSubscriptionConfigCloudflareVersion = "1.0.0"
)
