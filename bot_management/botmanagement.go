// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// BotManagementService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBotManagementService] method instead.
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
func (r *BotManagementService) Update(ctx context.Context, params BotManagementUpdateParams, opts ...option.RequestOption) (res *BotManagementUpdateResponse, err error) {
	var env BotManagementUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/bot_management", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve a zone's Bot Management Config
func (r *BotManagementService) Get(ctx context.Context, query BotManagementGetParams, opts ...option.RequestOption) (res *BotManagementGetResponse, err error) {
	var env BotManagementGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/bot_management", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BotFightModeConfiguration struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                          `json:"using_latest_model"`
	JSON             botFightModeConfigurationJSON `json:"-"`
}

// botFightModeConfigurationJSON contains the JSON metadata for the struct
// [BotFightModeConfiguration]
type botFightModeConfigurationJSON struct {
	EnableJS         apijson.Field
	FightMode        apijson.Field
	UsingLatestModel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BotFightModeConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botFightModeConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r BotFightModeConfiguration) implementsBotManagementBotManagementUpdateResponse() {}

func (r BotFightModeConfiguration) implementsBotManagementBotManagementGetResponse() {}

type BotFightModeConfigurationParam struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS param.Field[bool] `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode param.Field[bool] `json:"fight_mode"`
}

func (r BotFightModeConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BotFightModeConfigurationParam) implementsBotManagementBotManagementUpdateParamsBodyUnion() {}

type SubscriptionConfiguration struct {
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel bool `json:"auto_update_model"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS bool `json:"enable_js"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore bool `json:"suppress_session_score"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                          `json:"using_latest_model"`
	JSON             subscriptionConfigurationJSON `json:"-"`
}

// subscriptionConfigurationJSON contains the JSON metadata for the struct
// [SubscriptionConfiguration]
type subscriptionConfigurationJSON struct {
	AutoUpdateModel      apijson.Field
	EnableJS             apijson.Field
	SuppressSessionScore apijson.Field
	UsingLatestModel     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SubscriptionConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionConfiguration) implementsBotManagementBotManagementUpdateResponse() {}

func (r SubscriptionConfiguration) implementsBotManagementBotManagementGetResponse() {}

type SubscriptionConfigurationParam struct {
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel param.Field[bool] `json:"auto_update_model"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS param.Field[bool] `json:"enable_js"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore param.Field[bool] `json:"suppress_session_score"`
}

func (r SubscriptionConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionConfigurationParam) implementsBotManagementBotManagementUpdateParamsBodyUnion() {}

type SuperBotFightModeDefinitelyConfiguration struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                         `json:"using_latest_model"`
	JSON             superBotFightModeDefinitelyConfigurationJSON `json:"-"`
}

// superBotFightModeDefinitelyConfigurationJSON contains the JSON metadata for the
// struct [SuperBotFightModeDefinitelyConfiguration]
type superBotFightModeDefinitelyConfigurationJSON struct {
	EnableJS                     apijson.Field
	OptimizeWordpress            apijson.Field
	SBFMDefinitelyAutomated      apijson.Field
	SBFMStaticResourceProtection apijson.Field
	SBFMVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SuperBotFightModeDefinitelyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superBotFightModeDefinitelyConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r SuperBotFightModeDefinitelyConfiguration) implementsBotManagementBotManagementUpdateResponse() {
}

func (r SuperBotFightModeDefinitelyConfiguration) implementsBotManagementBotManagementGetResponse() {}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated string

const (
	SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomatedAllow            SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated = "allow"
	SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomatedBlock            SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated = "block"
	SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomatedManagedChallenge SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated = "managed_challenge"
)

func (r SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated) IsKnown() bool {
	switch r {
	case SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomatedAllow, SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomatedBlock, SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBots string

const (
	SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBotsAllow SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBots = "allow"
	SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBotsBlock SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBots = "block"
)

func (r SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBots) IsKnown() bool {
	switch r {
	case SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBotsAllow, SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBotsBlock:
		return true
	}
	return false
}

type SuperBotFightModeDefinitelyConfigurationParam struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated param.Field[SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots param.Field[SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r SuperBotFightModeDefinitelyConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuperBotFightModeDefinitelyConfigurationParam) implementsBotManagementBotManagementUpdateParamsBodyUnion() {
}

type SuperBotFightModeLikelyConfiguration struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SBFMLikelyAutomated SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots SuperBotFightModeLikelyConfigurationSBFMVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                     `json:"using_latest_model"`
	JSON             superBotFightModeLikelyConfigurationJSON `json:"-"`
}

// superBotFightModeLikelyConfigurationJSON contains the JSON metadata for the
// struct [SuperBotFightModeLikelyConfiguration]
type superBotFightModeLikelyConfigurationJSON struct {
	EnableJS                     apijson.Field
	OptimizeWordpress            apijson.Field
	SBFMDefinitelyAutomated      apijson.Field
	SBFMLikelyAutomated          apijson.Field
	SBFMStaticResourceProtection apijson.Field
	SBFMVerifiedBots             apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SuperBotFightModeLikelyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superBotFightModeLikelyConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r SuperBotFightModeLikelyConfiguration) implementsBotManagementBotManagementUpdateResponse() {}

func (r SuperBotFightModeLikelyConfiguration) implementsBotManagementBotManagementGetResponse() {}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated string

const (
	SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomatedAllow            SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated = "allow"
	SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomatedBlock            SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated = "block"
	SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomatedManagedChallenge SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated = "managed_challenge"
)

func (r SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated) IsKnown() bool {
	switch r {
	case SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomatedAllow, SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomatedBlock, SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated string

const (
	SuperBotFightModeLikelyConfigurationSBFMLikelyAutomatedAllow            SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated = "allow"
	SuperBotFightModeLikelyConfigurationSBFMLikelyAutomatedBlock            SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated = "block"
	SuperBotFightModeLikelyConfigurationSBFMLikelyAutomatedManagedChallenge SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated = "managed_challenge"
)

func (r SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated) IsKnown() bool {
	switch r {
	case SuperBotFightModeLikelyConfigurationSBFMLikelyAutomatedAllow, SuperBotFightModeLikelyConfigurationSBFMLikelyAutomatedBlock, SuperBotFightModeLikelyConfigurationSBFMLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type SuperBotFightModeLikelyConfigurationSBFMVerifiedBots string

const (
	SuperBotFightModeLikelyConfigurationSBFMVerifiedBotsAllow SuperBotFightModeLikelyConfigurationSBFMVerifiedBots = "allow"
	SuperBotFightModeLikelyConfigurationSBFMVerifiedBotsBlock SuperBotFightModeLikelyConfigurationSBFMVerifiedBots = "block"
)

func (r SuperBotFightModeLikelyConfigurationSBFMVerifiedBots) IsKnown() bool {
	switch r {
	case SuperBotFightModeLikelyConfigurationSBFMVerifiedBotsAllow, SuperBotFightModeLikelyConfigurationSBFMVerifiedBotsBlock:
		return true
	}
	return false
}

type SuperBotFightModeLikelyConfigurationParam struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated param.Field[SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SBFMLikelyAutomated param.Field[SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated] `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots param.Field[SuperBotFightModeLikelyConfigurationSBFMVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r SuperBotFightModeLikelyConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuperBotFightModeLikelyConfigurationParam) implementsBotManagementBotManagementUpdateParamsBodyUnion() {
}

type BotManagementUpdateResponse struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS bool `json:"enable_js"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool `json:"using_latest_model"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated BotManagementUpdateResponseSBFMDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots BotManagementUpdateResponseSBFMVerifiedBots `json:"sbfm_verified_bots"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SBFMLikelyAutomated BotManagementUpdateResponseSBFMLikelyAutomated `json:"sbfm_likely_automated"`
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel bool `json:"auto_update_model"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore bool                            `json:"suppress_session_score"`
	JSON                 botManagementUpdateResponseJSON `json:"-"`
	union                BotManagementUpdateResponseUnion
}

// botManagementUpdateResponseJSON contains the JSON metadata for the struct
// [BotManagementUpdateResponse]
type botManagementUpdateResponseJSON struct {
	EnableJS                     apijson.Field
	UsingLatestModel             apijson.Field
	FightMode                    apijson.Field
	OptimizeWordpress            apijson.Field
	SBFMDefinitelyAutomated      apijson.Field
	SBFMStaticResourceProtection apijson.Field
	SBFMVerifiedBots             apijson.Field
	SBFMLikelyAutomated          apijson.Field
	AutoUpdateModel              apijson.Field
	SuppressSessionScore         apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r botManagementUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *BotManagementUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BotManagementUpdateResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [bot_management.BotFightModeConfiguration],
// [bot_management.SuperBotFightModeDefinitelyConfiguration],
// [bot_management.SuperBotFightModeLikelyConfiguration],
// [bot_management.SubscriptionConfiguration].
func (r BotManagementUpdateResponse) AsUnion() BotManagementUpdateResponseUnion {
	return r.union
}

// Union satisfied by [bot_management.BotFightModeConfiguration],
// [bot_management.SuperBotFightModeDefinitelyConfiguration],
// [bot_management.SuperBotFightModeLikelyConfiguration] or
// [bot_management.SubscriptionConfiguration].
type BotManagementUpdateResponseUnion interface {
	implementsBotManagementBotManagementUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BotManagementUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BotFightModeConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuperBotFightModeDefinitelyConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuperBotFightModeLikelyConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionConfiguration{}),
		},
	)
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateResponseSBFMDefinitelyAutomated string

const (
	BotManagementUpdateResponseSBFMDefinitelyAutomatedAllow            BotManagementUpdateResponseSBFMDefinitelyAutomated = "allow"
	BotManagementUpdateResponseSBFMDefinitelyAutomatedBlock            BotManagementUpdateResponseSBFMDefinitelyAutomated = "block"
	BotManagementUpdateResponseSBFMDefinitelyAutomatedManagedChallenge BotManagementUpdateResponseSBFMDefinitelyAutomated = "managed_challenge"
)

func (r BotManagementUpdateResponseSBFMDefinitelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementUpdateResponseSBFMDefinitelyAutomatedAllow, BotManagementUpdateResponseSBFMDefinitelyAutomatedBlock, BotManagementUpdateResponseSBFMDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateResponseSBFMVerifiedBots string

const (
	BotManagementUpdateResponseSBFMVerifiedBotsAllow BotManagementUpdateResponseSBFMVerifiedBots = "allow"
	BotManagementUpdateResponseSBFMVerifiedBotsBlock BotManagementUpdateResponseSBFMVerifiedBots = "block"
)

func (r BotManagementUpdateResponseSBFMVerifiedBots) IsKnown() bool {
	switch r {
	case BotManagementUpdateResponseSBFMVerifiedBotsAllow, BotManagementUpdateResponseSBFMVerifiedBotsBlock:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementUpdateResponseSBFMLikelyAutomated string

const (
	BotManagementUpdateResponseSBFMLikelyAutomatedAllow            BotManagementUpdateResponseSBFMLikelyAutomated = "allow"
	BotManagementUpdateResponseSBFMLikelyAutomatedBlock            BotManagementUpdateResponseSBFMLikelyAutomated = "block"
	BotManagementUpdateResponseSBFMLikelyAutomatedManagedChallenge BotManagementUpdateResponseSBFMLikelyAutomated = "managed_challenge"
)

func (r BotManagementUpdateResponseSBFMLikelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementUpdateResponseSBFMLikelyAutomatedAllow, BotManagementUpdateResponseSBFMLikelyAutomatedBlock, BotManagementUpdateResponseSBFMLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

type BotManagementGetResponse struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS bool `json:"enable_js"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool `json:"using_latest_model"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated BotManagementGetResponseSBFMDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots BotManagementGetResponseSBFMVerifiedBots `json:"sbfm_verified_bots"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SBFMLikelyAutomated BotManagementGetResponseSBFMLikelyAutomated `json:"sbfm_likely_automated"`
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel bool `json:"auto_update_model"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore bool                         `json:"suppress_session_score"`
	JSON                 botManagementGetResponseJSON `json:"-"`
	union                BotManagementGetResponseUnion
}

// botManagementGetResponseJSON contains the JSON metadata for the struct
// [BotManagementGetResponse]
type botManagementGetResponseJSON struct {
	EnableJS                     apijson.Field
	UsingLatestModel             apijson.Field
	FightMode                    apijson.Field
	OptimizeWordpress            apijson.Field
	SBFMDefinitelyAutomated      apijson.Field
	SBFMStaticResourceProtection apijson.Field
	SBFMVerifiedBots             apijson.Field
	SBFMLikelyAutomated          apijson.Field
	AutoUpdateModel              apijson.Field
	SuppressSessionScore         apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r botManagementGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *BotManagementGetResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BotManagementGetResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [bot_management.BotFightModeConfiguration],
// [bot_management.SuperBotFightModeDefinitelyConfiguration],
// [bot_management.SuperBotFightModeLikelyConfiguration],
// [bot_management.SubscriptionConfiguration].
func (r BotManagementGetResponse) AsUnion() BotManagementGetResponseUnion {
	return r.union
}

// Union satisfied by [bot_management.BotFightModeConfiguration],
// [bot_management.SuperBotFightModeDefinitelyConfiguration],
// [bot_management.SuperBotFightModeLikelyConfiguration] or
// [bot_management.SubscriptionConfiguration].
type BotManagementGetResponseUnion interface {
	implementsBotManagementBotManagementGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BotManagementGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BotFightModeConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuperBotFightModeDefinitelyConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuperBotFightModeLikelyConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionConfiguration{}),
		},
	)
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementGetResponseSBFMDefinitelyAutomated string

const (
	BotManagementGetResponseSBFMDefinitelyAutomatedAllow            BotManagementGetResponseSBFMDefinitelyAutomated = "allow"
	BotManagementGetResponseSBFMDefinitelyAutomatedBlock            BotManagementGetResponseSBFMDefinitelyAutomated = "block"
	BotManagementGetResponseSBFMDefinitelyAutomatedManagedChallenge BotManagementGetResponseSBFMDefinitelyAutomated = "managed_challenge"
)

func (r BotManagementGetResponseSBFMDefinitelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementGetResponseSBFMDefinitelyAutomatedAllow, BotManagementGetResponseSBFMDefinitelyAutomatedBlock, BotManagementGetResponseSBFMDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementGetResponseSBFMVerifiedBots string

const (
	BotManagementGetResponseSBFMVerifiedBotsAllow BotManagementGetResponseSBFMVerifiedBots = "allow"
	BotManagementGetResponseSBFMVerifiedBotsBlock BotManagementGetResponseSBFMVerifiedBots = "block"
)

func (r BotManagementGetResponseSBFMVerifiedBots) IsKnown() bool {
	switch r {
	case BotManagementGetResponseSBFMVerifiedBotsAllow, BotManagementGetResponseSBFMVerifiedBotsBlock:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementGetResponseSBFMLikelyAutomated string

const (
	BotManagementGetResponseSBFMLikelyAutomatedAllow            BotManagementGetResponseSBFMLikelyAutomated = "allow"
	BotManagementGetResponseSBFMLikelyAutomatedBlock            BotManagementGetResponseSBFMLikelyAutomated = "block"
	BotManagementGetResponseSBFMLikelyAutomatedManagedChallenge BotManagementGetResponseSBFMLikelyAutomated = "managed_challenge"
)

func (r BotManagementGetResponseSBFMLikelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementGetResponseSBFMLikelyAutomatedAllow, BotManagementGetResponseSBFMLikelyAutomatedBlock, BotManagementGetResponseSBFMLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

type BotManagementUpdateParams struct {
	// Identifier
	ZoneID param.Field[string]                `path:"zone_id,required"`
	Body   BotManagementUpdateParamsBodyUnion `json:"body,required"`
}

func (r BotManagementUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type BotManagementUpdateParamsBody struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS param.Field[bool] `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode param.Field[bool] `json:"fight_mode"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated param.Field[BotManagementUpdateParamsBodySBFMDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots param.Field[BotManagementUpdateParamsBodySBFMVerifiedBots] `json:"sbfm_verified_bots"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SBFMLikelyAutomated param.Field[BotManagementUpdateParamsBodySBFMLikelyAutomated] `json:"sbfm_likely_automated"`
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel param.Field[bool] `json:"auto_update_model"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore param.Field[bool] `json:"suppress_session_score"`
}

func (r BotManagementUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BotManagementUpdateParamsBody) implementsBotManagementBotManagementUpdateParamsBodyUnion() {}

// Satisfied by [bot_management.BotFightModeConfigurationParam],
// [bot_management.SuperBotFightModeDefinitelyConfigurationParam],
// [bot_management.SuperBotFightModeLikelyConfigurationParam],
// [bot_management.SubscriptionConfigurationParam],
// [BotManagementUpdateParamsBody].
type BotManagementUpdateParamsBodyUnion interface {
	implementsBotManagementBotManagementUpdateParamsBodyUnion()
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateParamsBodySBFMDefinitelyAutomated string

const (
	BotManagementUpdateParamsBodySBFMDefinitelyAutomatedAllow            BotManagementUpdateParamsBodySBFMDefinitelyAutomated = "allow"
	BotManagementUpdateParamsBodySBFMDefinitelyAutomatedBlock            BotManagementUpdateParamsBodySBFMDefinitelyAutomated = "block"
	BotManagementUpdateParamsBodySBFMDefinitelyAutomatedManagedChallenge BotManagementUpdateParamsBodySBFMDefinitelyAutomated = "managed_challenge"
)

func (r BotManagementUpdateParamsBodySBFMDefinitelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementUpdateParamsBodySBFMDefinitelyAutomatedAllow, BotManagementUpdateParamsBodySBFMDefinitelyAutomatedBlock, BotManagementUpdateParamsBodySBFMDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateParamsBodySBFMVerifiedBots string

const (
	BotManagementUpdateParamsBodySBFMVerifiedBotsAllow BotManagementUpdateParamsBodySBFMVerifiedBots = "allow"
	BotManagementUpdateParamsBodySBFMVerifiedBotsBlock BotManagementUpdateParamsBodySBFMVerifiedBots = "block"
)

func (r BotManagementUpdateParamsBodySBFMVerifiedBots) IsKnown() bool {
	switch r {
	case BotManagementUpdateParamsBodySBFMVerifiedBotsAllow, BotManagementUpdateParamsBodySBFMVerifiedBotsBlock:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementUpdateParamsBodySBFMLikelyAutomated string

const (
	BotManagementUpdateParamsBodySBFMLikelyAutomatedAllow            BotManagementUpdateParamsBodySBFMLikelyAutomated = "allow"
	BotManagementUpdateParamsBodySBFMLikelyAutomatedBlock            BotManagementUpdateParamsBodySBFMLikelyAutomated = "block"
	BotManagementUpdateParamsBodySBFMLikelyAutomatedManagedChallenge BotManagementUpdateParamsBodySBFMLikelyAutomated = "managed_challenge"
)

func (r BotManagementUpdateParamsBodySBFMLikelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementUpdateParamsBodySBFMLikelyAutomatedAllow, BotManagementUpdateParamsBodySBFMLikelyAutomatedBlock, BotManagementUpdateParamsBodySBFMLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

type BotManagementUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success BotManagementUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  BotManagementUpdateResponse                `json:"result"`
	JSON    botManagementUpdateResponseEnvelopeJSON    `json:"-"`
}

// botManagementUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [BotManagementUpdateResponseEnvelope]
type botManagementUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botManagementUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BotManagementUpdateResponseEnvelopeSuccess bool

const (
	BotManagementUpdateResponseEnvelopeSuccessTrue BotManagementUpdateResponseEnvelopeSuccess = true
)

func (r BotManagementUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BotManagementUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BotManagementGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type BotManagementGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success BotManagementGetResponseEnvelopeSuccess `json:"success,required"`
	Result  BotManagementGetResponse                `json:"result"`
	JSON    botManagementGetResponseEnvelopeJSON    `json:"-"`
}

// botManagementGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BotManagementGetResponseEnvelope]
type botManagementGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botManagementGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BotManagementGetResponseEnvelopeSuccess bool

const (
	BotManagementGetResponseEnvelopeSuccessTrue BotManagementGetResponseEnvelopeSuccess = true
)

func (r BotManagementGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BotManagementGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
