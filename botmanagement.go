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
	var env BotManagementGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/bot_management", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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
func (r *BotManagementService) Replace(ctx context.Context, zoneID string, body BotManagementReplaceParams, opts ...option.RequestOption) (res *BotManagementReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BotManagementReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/bot_management", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [BotManagementGetResponseBotManagementBotFightModeConfig],
// [BotManagementGetResponseBotManagementSbfmDefinitelyConfig],
// [BotManagementGetResponseBotManagementSbfmLikelyConfig] or
// [BotManagementGetResponseBotManagementBmSubscriptionConfig].
type BotManagementGetResponse interface {
	implementsBotManagementGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*BotManagementGetResponse)(nil)).Elem(), "")
}

type BotManagementGetResponseBotManagementBotFightModeConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                        `json:"using_latest_model"`
	JSON             botManagementGetResponseBotManagementBotFightModeConfigJSON `json:"-"`
}

// botManagementGetResponseBotManagementBotFightModeConfigJSON contains the JSON
// metadata for the struct
// [BotManagementGetResponseBotManagementBotFightModeConfig]
type botManagementGetResponseBotManagementBotFightModeConfigJSON struct {
	EnableJs         apijson.Field
	FightMode        apijson.Field
	UsingLatestModel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BotManagementGetResponseBotManagementBotFightModeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementGetResponseBotManagementBotFightModeConfig) implementsBotManagementGetResponse() {
}

type BotManagementGetResponseBotManagementSbfmDefinitelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                          `json:"using_latest_model"`
	JSON             botManagementGetResponseBotManagementSbfmDefinitelyConfigJSON `json:"-"`
}

// botManagementGetResponseBotManagementSbfmDefinitelyConfigJSON contains the JSON
// metadata for the struct
// [BotManagementGetResponseBotManagementSbfmDefinitelyConfig]
type botManagementGetResponseBotManagementSbfmDefinitelyConfigJSON struct {
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *BotManagementGetResponseBotManagementSbfmDefinitelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementGetResponseBotManagementSbfmDefinitelyConfig) implementsBotManagementGetResponse() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedAllow            BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedBlock            BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots string

const (
	BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsAllow BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "allow"
	BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsBlock BotManagementGetResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "block"
)

type BotManagementGetResponseBotManagementSbfmLikelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                      `json:"using_latest_model"`
	JSON             botManagementGetResponseBotManagementSbfmLikelyConfigJSON `json:"-"`
}

// botManagementGetResponseBotManagementSbfmLikelyConfigJSON contains the JSON
// metadata for the struct [BotManagementGetResponseBotManagementSbfmLikelyConfig]
type botManagementGetResponseBotManagementSbfmLikelyConfigJSON struct {
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

func (r *BotManagementGetResponseBotManagementSbfmLikelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementGetResponseBotManagementSbfmLikelyConfig) implementsBotManagementGetResponse() {}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedAllow            BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedBlock            BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated string

const (
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomatedAllow            BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "allow"
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomatedBlock            BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "block"
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots string

const (
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmVerifiedBotsAllow BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots = "allow"
	BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmVerifiedBotsBlock BotManagementGetResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots = "block"
)

type BotManagementGetResponseBotManagementBmSubscriptionConfig struct {
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
	UsingLatestModel bool                                                          `json:"using_latest_model"`
	JSON             botManagementGetResponseBotManagementBmSubscriptionConfigJSON `json:"-"`
}

// botManagementGetResponseBotManagementBmSubscriptionConfigJSON contains the JSON
// metadata for the struct
// [BotManagementGetResponseBotManagementBmSubscriptionConfig]
type botManagementGetResponseBotManagementBmSubscriptionConfigJSON struct {
	AutoUpdateModel      apijson.Field
	EnableJs             apijson.Field
	SuppressSessionScore apijson.Field
	UsingLatestModel     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BotManagementGetResponseBotManagementBmSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementGetResponseBotManagementBmSubscriptionConfig) implementsBotManagementGetResponse() {
}

// Union satisfied by
// [BotManagementReplaceResponseBotManagementBotFightModeConfig],
// [BotManagementReplaceResponseBotManagementSbfmDefinitelyConfig],
// [BotManagementReplaceResponseBotManagementSbfmLikelyConfig] or
// [BotManagementReplaceResponseBotManagementBmSubscriptionConfig].
type BotManagementReplaceResponse interface {
	implementsBotManagementReplaceResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*BotManagementReplaceResponse)(nil)).Elem(), "")
}

type BotManagementReplaceResponseBotManagementBotFightModeConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                            `json:"using_latest_model"`
	JSON             botManagementReplaceResponseBotManagementBotFightModeConfigJSON `json:"-"`
}

// botManagementReplaceResponseBotManagementBotFightModeConfigJSON contains the
// JSON metadata for the struct
// [BotManagementReplaceResponseBotManagementBotFightModeConfig]
type botManagementReplaceResponseBotManagementBotFightModeConfigJSON struct {
	EnableJs         apijson.Field
	FightMode        apijson.Field
	UsingLatestModel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BotManagementReplaceResponseBotManagementBotFightModeConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementReplaceResponseBotManagementBotFightModeConfig) implementsBotManagementReplaceResponse() {
}

type BotManagementReplaceResponseBotManagementSbfmDefinitelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                              `json:"using_latest_model"`
	JSON             botManagementReplaceResponseBotManagementSbfmDefinitelyConfigJSON `json:"-"`
}

// botManagementReplaceResponseBotManagementSbfmDefinitelyConfigJSON contains the
// JSON metadata for the struct
// [BotManagementReplaceResponseBotManagementSbfmDefinitelyConfig]
type botManagementReplaceResponseBotManagementSbfmDefinitelyConfigJSON struct {
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *BotManagementReplaceResponseBotManagementSbfmDefinitelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementReplaceResponseBotManagementSbfmDefinitelyConfig) implementsBotManagementReplaceResponse() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedAllow            BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedBlock            BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots string

const (
	BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsAllow BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "allow"
	BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsBlock BotManagementReplaceResponseBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "block"
)

type BotManagementReplaceResponseBotManagementSbfmLikelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                                          `json:"using_latest_model"`
	JSON             botManagementReplaceResponseBotManagementSbfmLikelyConfigJSON `json:"-"`
}

// botManagementReplaceResponseBotManagementSbfmLikelyConfigJSON contains the JSON
// metadata for the struct
// [BotManagementReplaceResponseBotManagementSbfmLikelyConfig]
type botManagementReplaceResponseBotManagementSbfmLikelyConfigJSON struct {
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

func (r *BotManagementReplaceResponseBotManagementSbfmLikelyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementReplaceResponseBotManagementSbfmLikelyConfig) implementsBotManagementReplaceResponse() {
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedAllow            BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedBlock            BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated string

const (
	BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomatedAllow            BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "allow"
	BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomatedBlock            BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "block"
	BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots string

const (
	BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmVerifiedBotsAllow BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots = "allow"
	BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmVerifiedBotsBlock BotManagementReplaceResponseBotManagementSbfmLikelyConfigSbfmVerifiedBots = "block"
)

type BotManagementReplaceResponseBotManagementBmSubscriptionConfig struct {
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
	UsingLatestModel bool                                                              `json:"using_latest_model"`
	JSON             botManagementReplaceResponseBotManagementBmSubscriptionConfigJSON `json:"-"`
}

// botManagementReplaceResponseBotManagementBmSubscriptionConfigJSON contains the
// JSON metadata for the struct
// [BotManagementReplaceResponseBotManagementBmSubscriptionConfig]
type botManagementReplaceResponseBotManagementBmSubscriptionConfigJSON struct {
	AutoUpdateModel      apijson.Field
	EnableJs             apijson.Field
	SuppressSessionScore apijson.Field
	UsingLatestModel     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BotManagementReplaceResponseBotManagementBmSubscriptionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r BotManagementReplaceResponseBotManagementBmSubscriptionConfig) implementsBotManagementReplaceResponse() {
}

type BotManagementGetResponseEnvelope struct {
	Errors   []BotManagementGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BotManagementGetResponseEnvelopeMessages `json:"messages,required"`
	Result   BotManagementGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success BotManagementGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    botManagementGetResponseEnvelopeJSON    `json:"-"`
}

// botManagementGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BotManagementGetResponseEnvelope]
type botManagementGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BotManagementGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    botManagementGetResponseEnvelopeErrorsJSON `json:"-"`
}

// botManagementGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BotManagementGetResponseEnvelopeErrors]
type botManagementGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BotManagementGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    botManagementGetResponseEnvelopeMessagesJSON `json:"-"`
}

// botManagementGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [BotManagementGetResponseEnvelopeMessages]
type botManagementGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type BotManagementGetResponseEnvelopeSuccess bool

const (
	BotManagementGetResponseEnvelopeSuccessTrue BotManagementGetResponseEnvelopeSuccess = true
)

// This interface is a union satisfied by one of the following:
// [BotManagementReplaceParamsBotManagementBotFightModeConfig],
// [BotManagementReplaceParamsBotManagementSbfmDefinitelyConfig],
// [BotManagementReplaceParamsBotManagementSbfmLikelyConfig],
// [BotManagementReplaceParamsBotManagementBmSubscriptionConfig].
type BotManagementReplaceParams interface {
	ImplementsBotManagementReplaceParams()
}

type BotManagementReplaceParamsBotManagementBotFightModeConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode param.Field[bool] `json:"fight_mode"`
}

func (r BotManagementReplaceParamsBotManagementBotFightModeConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (BotManagementReplaceParamsBotManagementBotFightModeConfig) ImplementsBotManagementReplaceParams() {

}

type BotManagementReplaceParamsBotManagementSbfmDefinitelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated param.Field[BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots param.Field[BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r BotManagementReplaceParamsBotManagementSbfmDefinitelyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (BotManagementReplaceParamsBotManagementSbfmDefinitelyConfig) ImplementsBotManagementReplaceParams() {

}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedAllow            BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedBlock            BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBots string

const (
	BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsAllow BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "allow"
	BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBotsBlock BotManagementReplaceParamsBotManagementSbfmDefinitelyConfigSbfmVerifiedBots = "block"
)

type BotManagementReplaceParamsBotManagementSbfmLikelyConfig struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated param.Field[BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated param.Field[BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated] `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots param.Field[BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r BotManagementReplaceParamsBotManagementSbfmLikelyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (BotManagementReplaceParamsBotManagementSbfmLikelyConfig) ImplementsBotManagementReplaceParams() {

}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated string

const (
	BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedAllow            BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "allow"
	BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedBlock            BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "block"
	BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomatedManagedChallenge BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmDefinitelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated string

const (
	BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomatedAllow            BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "allow"
	BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomatedBlock            BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "block"
	BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomatedManagedChallenge BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmLikelyAutomated = "managed_challenge"
)

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmVerifiedBots string

const (
	BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmVerifiedBotsAllow BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmVerifiedBots = "allow"
	BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmVerifiedBotsBlock BotManagementReplaceParamsBotManagementSbfmLikelyConfigSbfmVerifiedBots = "block"
)

type BotManagementReplaceParamsBotManagementBmSubscriptionConfig struct {
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

func (r BotManagementReplaceParamsBotManagementBmSubscriptionConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (BotManagementReplaceParamsBotManagementBmSubscriptionConfig) ImplementsBotManagementReplaceParams() {

}

type BotManagementReplaceResponseEnvelope struct {
	Errors   []BotManagementReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BotManagementReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   BotManagementReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success BotManagementReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    botManagementReplaceResponseEnvelopeJSON    `json:"-"`
}

// botManagementReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [BotManagementReplaceResponseEnvelope]
type botManagementReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BotManagementReplaceResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    botManagementReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// botManagementReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BotManagementReplaceResponseEnvelopeErrors]
type botManagementReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BotManagementReplaceResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    botManagementReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// botManagementReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BotManagementReplaceResponseEnvelopeMessages]
type botManagementReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type BotManagementReplaceResponseEnvelopeSuccess bool

const (
	BotManagementReplaceResponseEnvelopeSuccessTrue BotManagementReplaceResponseEnvelopeSuccess = true
)
