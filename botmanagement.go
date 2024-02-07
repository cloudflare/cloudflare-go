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

// BotManagementService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewBotManagementService] method
// instead.
type BotManagementService struct {
	Options []option.RequestOption
}

// NewBotManagementService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBotManagementService(opts ...option.RequestOption) (r *BotManagementService) {
	r = &BotManagementService{}
	r.Options = opts
	return
}

// Retrieve a zone's Bot Management Config
func (r *BotManagementService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *BotManagementGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/bot_management", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *BotManagementService) Update(ctx context.Context, zoneID string, body BotManagementUpdateParams, opts ...option.RequestOption) (res *BotManagementUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/bot_management", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type BotManagementGetResponse struct {
	Errors   []BotManagementGetResponseError   `json:"errors"`
	Messages []BotManagementGetResponseMessage `json:"messages"`
	Result   BotManagementGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success BotManagementGetResponseSuccess `json:"success"`
	JSON    botManagementGetResponseJSON    `json:"-"`
}

// botManagementGetResponseJSON contains the JSON metadata for the struct
// [BotManagementGetResponse]
type botManagementGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BotManagementGetResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    botManagementGetResponseErrorJSON `json:"-"`
}

// botManagementGetResponseErrorJSON contains the JSON metadata for the struct
// [BotManagementGetResponseError]
type botManagementGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BotManagementGetResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    botManagementGetResponseMessageJSON `json:"-"`
}

// botManagementGetResponseMessageJSON contains the JSON metadata for the struct
// [BotManagementGetResponseMessage]
type botManagementGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [BotManagementGetResponseResultBotManagementBotFightModeConfig],
// [BotManagementGetResponseResultBotManagementSbfmDefinitelyConfig],
// [BotManagementGetResponseResultBotManagementSbfmLikelyConfig] or
// [BotManagementGetResponseResultBotManagementBmSubscriptionConfig].
type BotManagementGetResponseResult interface {
	implementsBotManagementGetResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*BotManagementGetResponseResult)(nil)).Elem(), "")
}

type BotManagementGetResponseResultBotManagementBotFightModeConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                              `json:"using_latest_model"`
	JSON             botManagementGetResponseResultBotManagementBotFightModeConfigJSON `json:"-"`
}

// botManagementGetResponseResultBotManagementBotFightModeConfigJSON contains the
// JSON metadata for the struct
// [BotManagementGetResponseResultBotManagementBotFightModeConfig]
type botManagementGetResponseResultBotManagementBotFightModeConfigJSON struct {
	EnableJs         apijson.Field
	FightMode        apijson.Field
	UsingLatestModel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BotManagementGetResponseResultBotManagementBotFightModeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementGetResponseResultBotManagementBotFightModeConfig) implementsBotManagementGetResponseResult() {
}

type BotManagementGetResponseResultBotManagementSbfmDefinitelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                                `json:"using_latest_model"`
	JSON             botManagementGetResponseResultBotManagementSbfmDefinitelyConfigJSON `json:"-"`
}

// botManagementGetResponseResultBotManagementSbfmDefinitelyConfigJSON contains the
// JSON metadata for the struct
// [BotManagementGetResponseResultBotManagementSbfmDefinitelyConfig]
type botManagementGetResponseResultBotManagementSbfmDefinitelyConfigJSON struct {
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *BotManagementGetResponseResultBotManagementSbfmDefinitelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementGetResponseResultBotManagementSbfmDefinitelyConfig) implementsBotManagementGetResponseResult() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedAllow            BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedBlock            BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmVerifiedBots string

const (
	BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsAllow BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "allow"
	BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsBlock BotManagementGetResponseResultBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "block"
)

type BotManagementGetResponseResultBotManagementSbfmLikelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                            `json:"using_latest_model"`
	JSON             botManagementGetResponseResultBotManagementSbfmLikelyConfigJSON `json:"-"`
}

// botManagementGetResponseResultBotManagementSbfmLikelyConfigJSON contains the
// JSON metadata for the struct
// [BotManagementGetResponseResultBotManagementSbfmLikelyConfig]
type botManagementGetResponseResultBotManagementSbfmLikelyConfigJSON struct {
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

func (r *BotManagementGetResponseResultBotManagementSbfmLikelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementGetResponseResultBotManagementSbfmLikelyConfig) implementsBotManagementGetResponseResult() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedAllow            BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedBlock            BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomated string

const (
	BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomatedAllow            BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "allow"
	BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomatedBlock            BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "block"
	BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmVerifiedBots string

const (
	BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmVerifiedBotsAllow BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmVerifiedBots = "allow"
	BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmVerifiedBotsBlock BotManagementGetResponseResultBotManagementSbfmLikelyConfigSbfmVerifiedBots = "block"
)

type BotManagementGetResponseResultBotManagementBmSubscriptionConfig struct {
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
	UsingLatestModel bool                                                                `json:"using_latest_model"`
	JSON             botManagementGetResponseResultBotManagementBmSubscriptionConfigJSON `json:"-"`
}

// botManagementGetResponseResultBotManagementBmSubscriptionConfigJSON contains the
// JSON metadata for the struct
// [BotManagementGetResponseResultBotManagementBmSubscriptionConfig]
type botManagementGetResponseResultBotManagementBmSubscriptionConfigJSON struct {
	AutoUpdateModel      apijson.Field
	EnableJs             apijson.Field
	SuppressSessionScore apijson.Field
	UsingLatestModel     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BotManagementGetResponseResultBotManagementBmSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementGetResponseResultBotManagementBmSubscriptionConfig) implementsBotManagementGetResponseResult() {
}

// Whether the API call was successful
type BotManagementGetResponseSuccess bool

const (
	BotManagementGetResponseSuccessTrue BotManagementGetResponseSuccess = true
)

type BotManagementUpdateResponse struct {
	Errors   []BotManagementUpdateResponseError   `json:"errors"`
	Messages []BotManagementUpdateResponseMessage `json:"messages"`
	Result   BotManagementUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success BotManagementUpdateResponseSuccess `json:"success"`
	JSON    botManagementUpdateResponseJSON    `json:"-"`
}

// botManagementUpdateResponseJSON contains the JSON metadata for the struct
// [BotManagementUpdateResponse]
type botManagementUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BotManagementUpdateResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    botManagementUpdateResponseErrorJSON `json:"-"`
}

// botManagementUpdateResponseErrorJSON contains the JSON metadata for the struct
// [BotManagementUpdateResponseError]
type botManagementUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BotManagementUpdateResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    botManagementUpdateResponseMessageJSON `json:"-"`
}

// botManagementUpdateResponseMessageJSON contains the JSON metadata for the struct
// [BotManagementUpdateResponseMessage]
type botManagementUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [BotManagementUpdateResponseResultBotManagementBotFightModeConfig],
// [BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfig],
// [BotManagementUpdateResponseResultBotManagementSbfmLikelyConfig] or
// [BotManagementUpdateResponseResultBotManagementBmSubscriptionConfig].
type BotManagementUpdateResponseResult interface {
	implementsBotManagementUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*BotManagementUpdateResponseResult)(nil)).Elem(), "")
}

type BotManagementUpdateResponseResultBotManagementBotFightModeConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                                 `json:"using_latest_model"`
	JSON             botManagementUpdateResponseResultBotManagementBotFightModeConfigJSON `json:"-"`
}

// botManagementUpdateResponseResultBotManagementBotFightModeConfigJSON contains
// the JSON metadata for the struct
// [BotManagementUpdateResponseResultBotManagementBotFightModeConfig]
type botManagementUpdateResponseResultBotManagementBotFightModeConfigJSON struct {
	EnableJs         apijson.Field
	FightMode        apijson.Field
	UsingLatestModel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BotManagementUpdateResponseResultBotManagementBotFightModeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementUpdateResponseResultBotManagementBotFightModeConfig) implementsBotManagementUpdateResponseResult() {
}

type BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                                   `json:"using_latest_model"`
	JSON             botManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigJSON `json:"-"`
}

// botManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigJSON contains
// the JSON metadata for the struct
// [BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfig]
type botManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigJSON struct {
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfig) implementsBotManagementUpdateResponseResult() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedAllow            BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedBlock            BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmVerifiedBots string

const (
	BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsAllow BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "allow"
	BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsBlock BotManagementUpdateResponseResultBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "block"
)

type BotManagementUpdateResponseResultBotManagementSbfmLikelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                               `json:"using_latest_model"`
	JSON             botManagementUpdateResponseResultBotManagementSbfmLikelyConfigJSON `json:"-"`
}

// botManagementUpdateResponseResultBotManagementSbfmLikelyConfigJSON contains the
// JSON metadata for the struct
// [BotManagementUpdateResponseResultBotManagementSbfmLikelyConfig]
type botManagementUpdateResponseResultBotManagementSbfmLikelyConfigJSON struct {
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

func (r *BotManagementUpdateResponseResultBotManagementSbfmLikelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementUpdateResponseResultBotManagementSbfmLikelyConfig) implementsBotManagementUpdateResponseResult() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedAllow            BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedBlock            BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomated string

const (
	BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomatedAllow            BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "allow"
	BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomatedBlock            BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "block"
	BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmVerifiedBots string

const (
	BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmVerifiedBotsAllow BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmVerifiedBots = "allow"
	BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmVerifiedBotsBlock BotManagementUpdateResponseResultBotManagementSbfmLikelyConfigSbfmVerifiedBots = "block"
)

type BotManagementUpdateResponseResultBotManagementBmSubscriptionConfig struct {
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
	UsingLatestModel bool                                                                   `json:"using_latest_model"`
	JSON             botManagementUpdateResponseResultBotManagementBmSubscriptionConfigJSON `json:"-"`
}

// botManagementUpdateResponseResultBotManagementBmSubscriptionConfigJSON contains
// the JSON metadata for the struct
// [BotManagementUpdateResponseResultBotManagementBmSubscriptionConfig]
type botManagementUpdateResponseResultBotManagementBmSubscriptionConfigJSON struct {
	AutoUpdateModel      apijson.Field
	EnableJs             apijson.Field
	SuppressSessionScore apijson.Field
	UsingLatestModel     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BotManagementUpdateResponseResultBotManagementBmSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementUpdateResponseResultBotManagementBmSubscriptionConfig) implementsBotManagementUpdateResponseResult() {
}

// Whether the API call was successful
type BotManagementUpdateResponseSuccess bool

const (
	BotManagementUpdateResponseSuccessTrue BotManagementUpdateResponseSuccess = true
)

// This interface is a union satisfied by one of the following:
// [BotManagementUpdateParamsBotManagementBotFightModeConfig],
// [BotManagementUpdateParamsBotManagementSbfmDefinitelyConfig],
// [BotManagementUpdateParamsBotManagementSbfmLikelyConfig],
// [BotManagementUpdateParamsBotManagementBmSubscriptionConfig].
type BotManagementUpdateParams interface {
	ImplementsBotManagementUpdateParams()
}

type BotManagementUpdateParamsBotManagementBotFightModeConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode param.Field[bool] `json:"fight_mode"`
}

func (r BotManagementUpdateParamsBotManagementBotFightModeConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (BotManagementUpdateParamsBotManagementBotFightModeConfig) ImplementsBotManagementUpdateParams() {

}

type BotManagementUpdateParamsBotManagementSbfmDefinitelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated param.Field[BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots param.Field[BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r BotManagementUpdateParamsBotManagementSbfmDefinitelyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (BotManagementUpdateParamsBotManagementSbfmDefinitelyConfig) ImplementsBotManagementUpdateParams() {

}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedAllow            BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedBlock            BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBots string

const (
	BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsAllow BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "allow"
	BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsBlock BotManagementUpdateParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "block"
)

type BotManagementUpdateParamsBotManagementSbfmLikelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated param.Field[BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated param.Field[BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated] `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots param.Field[BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r BotManagementUpdateParamsBotManagementSbfmLikelyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (BotManagementUpdateParamsBotManagementSbfmLikelyConfig) ImplementsBotManagementUpdateParams() {

}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedAllow            BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedBlock            BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated string

const (
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomatedAllow            BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "allow"
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomatedBlock            BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "block"
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmVerifiedBots string

const (
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmVerifiedBotsAllow BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmVerifiedBots = "allow"
	BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmVerifiedBotsBlock BotManagementUpdateParamsBotManagementSbfmLikelyConfigSbfmVerifiedBots = "block"
)

type BotManagementUpdateParamsBotManagementBmSubscriptionConfig struct {
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
}

func (r BotManagementUpdateParamsBotManagementBmSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (BotManagementUpdateParamsBotManagementBmSubscriptionConfig) ImplementsBotManagementUpdateParams() {

}
