// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
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
	opts = append(r.Options[:], opts...)
	var env BotManagementUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/bot_management", params.getZoneID())
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve a zone's Bot Management Config
func (r *BotManagementService) Get(ctx context.Context, query BotManagementGetParams, opts ...option.RequestOption) (res *BotManagementGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BotManagementGetResponseEnvelope
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
	EnableJs bool `json:"enable_js"`
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
	EnableJs         apijson.Field
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

type SubscriptionConfiguration struct {
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
	UsingLatestModel bool                          `json:"using_latest_model"`
	JSON             subscriptionConfigurationJSON `json:"-"`
}

// subscriptionConfigurationJSON contains the JSON metadata for the struct
// [SubscriptionConfiguration]
type subscriptionConfigurationJSON struct {
	AutoUpdateModel      apijson.Field
	EnableJs             apijson.Field
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

type SuperBotFightModeDefinitelyConfiguration struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated SuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots SuperBotFightModeDefinitelyConfigurationSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                         `json:"using_latest_model"`
	JSON             superBotFightModeDefinitelyConfigurationJSON `json:"-"`
}

// superBotFightModeDefinitelyConfigurationJSON contains the JSON metadata for the
// struct [SuperBotFightModeDefinitelyConfiguration]
type superBotFightModeDefinitelyConfigurationJSON struct {
	EnableJs                     apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
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
type SuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomated string

const (
	SuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomatedAllow            SuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomated = "allow"
	SuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomatedBlock            SuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomated = "block"
	SuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomatedManagedChallenge SuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomated = "managed_challenge"
)

func (r SuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomated) IsKnown() bool {
	switch r {
	case SuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomatedAllow, SuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomatedBlock, SuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type SuperBotFightModeDefinitelyConfigurationSbfmVerifiedBots string

const (
	SuperBotFightModeDefinitelyConfigurationSbfmVerifiedBotsAllow SuperBotFightModeDefinitelyConfigurationSbfmVerifiedBots = "allow"
	SuperBotFightModeDefinitelyConfigurationSbfmVerifiedBotsBlock SuperBotFightModeDefinitelyConfigurationSbfmVerifiedBots = "block"
)

func (r SuperBotFightModeDefinitelyConfigurationSbfmVerifiedBots) IsKnown() bool {
	switch r {
	case SuperBotFightModeDefinitelyConfigurationSbfmVerifiedBotsAllow, SuperBotFightModeDefinitelyConfigurationSbfmVerifiedBotsBlock:
		return true
	}
	return false
}

type SuperBotFightModeLikelyConfiguration struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated SuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated SuperBotFightModeLikelyConfigurationSbfmLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots SuperBotFightModeLikelyConfigurationSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                     `json:"using_latest_model"`
	JSON             superBotFightModeLikelyConfigurationJSON `json:"-"`
}

// superBotFightModeLikelyConfigurationJSON contains the JSON metadata for the
// struct [SuperBotFightModeLikelyConfiguration]
type superBotFightModeLikelyConfigurationJSON struct {
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

func (r *SuperBotFightModeLikelyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superBotFightModeLikelyConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r SuperBotFightModeLikelyConfiguration) implementsBotManagementBotManagementUpdateResponse() {}

func (r SuperBotFightModeLikelyConfiguration) implementsBotManagementBotManagementGetResponse() {}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type SuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomated string

const (
	SuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomatedAllow            SuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomated = "allow"
	SuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomatedBlock            SuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomated = "block"
	SuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomatedManagedChallenge SuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomated = "managed_challenge"
)

func (r SuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomated) IsKnown() bool {
	switch r {
	case SuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomatedAllow, SuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomatedBlock, SuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type SuperBotFightModeLikelyConfigurationSbfmLikelyAutomated string

const (
	SuperBotFightModeLikelyConfigurationSbfmLikelyAutomatedAllow            SuperBotFightModeLikelyConfigurationSbfmLikelyAutomated = "allow"
	SuperBotFightModeLikelyConfigurationSbfmLikelyAutomatedBlock            SuperBotFightModeLikelyConfigurationSbfmLikelyAutomated = "block"
	SuperBotFightModeLikelyConfigurationSbfmLikelyAutomatedManagedChallenge SuperBotFightModeLikelyConfigurationSbfmLikelyAutomated = "managed_challenge"
)

func (r SuperBotFightModeLikelyConfigurationSbfmLikelyAutomated) IsKnown() bool {
	switch r {
	case SuperBotFightModeLikelyConfigurationSbfmLikelyAutomatedAllow, SuperBotFightModeLikelyConfigurationSbfmLikelyAutomatedBlock, SuperBotFightModeLikelyConfigurationSbfmLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type SuperBotFightModeLikelyConfigurationSbfmVerifiedBots string

const (
	SuperBotFightModeLikelyConfigurationSbfmVerifiedBotsAllow SuperBotFightModeLikelyConfigurationSbfmVerifiedBots = "allow"
	SuperBotFightModeLikelyConfigurationSbfmVerifiedBotsBlock SuperBotFightModeLikelyConfigurationSbfmVerifiedBots = "block"
)

func (r SuperBotFightModeLikelyConfigurationSbfmVerifiedBots) IsKnown() bool {
	switch r {
	case SuperBotFightModeLikelyConfigurationSbfmVerifiedBotsAllow, SuperBotFightModeLikelyConfigurationSbfmVerifiedBotsBlock:
		return true
	}
	return false
}

type BotManagementUpdateResponse struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool `json:"using_latest_model"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementUpdateResponseSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementUpdateResponseSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated BotManagementUpdateResponseSbfmLikelyAutomated `json:"sbfm_likely_automated"`
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
	EnableJs                     apijson.Field
	UsingLatestModel             apijson.Field
	FightMode                    apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	SbfmLikelyAutomated          apijson.Field
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
type BotManagementUpdateResponseSbfmDefinitelyAutomated string

const (
	BotManagementUpdateResponseSbfmDefinitelyAutomatedAllow            BotManagementUpdateResponseSbfmDefinitelyAutomated = "allow"
	BotManagementUpdateResponseSbfmDefinitelyAutomatedBlock            BotManagementUpdateResponseSbfmDefinitelyAutomated = "block"
	BotManagementUpdateResponseSbfmDefinitelyAutomatedManagedChallenge BotManagementUpdateResponseSbfmDefinitelyAutomated = "managed_challenge"
)

func (r BotManagementUpdateResponseSbfmDefinitelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementUpdateResponseSbfmDefinitelyAutomatedAllow, BotManagementUpdateResponseSbfmDefinitelyAutomatedBlock, BotManagementUpdateResponseSbfmDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateResponseSbfmVerifiedBots string

const (
	BotManagementUpdateResponseSbfmVerifiedBotsAllow BotManagementUpdateResponseSbfmVerifiedBots = "allow"
	BotManagementUpdateResponseSbfmVerifiedBotsBlock BotManagementUpdateResponseSbfmVerifiedBots = "block"
)

func (r BotManagementUpdateResponseSbfmVerifiedBots) IsKnown() bool {
	switch r {
	case BotManagementUpdateResponseSbfmVerifiedBotsAllow, BotManagementUpdateResponseSbfmVerifiedBotsBlock:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementUpdateResponseSbfmLikelyAutomated string

const (
	BotManagementUpdateResponseSbfmLikelyAutomatedAllow            BotManagementUpdateResponseSbfmLikelyAutomated = "allow"
	BotManagementUpdateResponseSbfmLikelyAutomatedBlock            BotManagementUpdateResponseSbfmLikelyAutomated = "block"
	BotManagementUpdateResponseSbfmLikelyAutomatedManagedChallenge BotManagementUpdateResponseSbfmLikelyAutomated = "managed_challenge"
)

func (r BotManagementUpdateResponseSbfmLikelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementUpdateResponseSbfmLikelyAutomatedAllow, BotManagementUpdateResponseSbfmLikelyAutomatedBlock, BotManagementUpdateResponseSbfmLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

type BotManagementGetResponse struct {
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs bool `json:"enable_js"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool `json:"using_latest_model"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated BotManagementGetResponseSbfmDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots BotManagementGetResponseSbfmVerifiedBots `json:"sbfm_verified_bots"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated BotManagementGetResponseSbfmLikelyAutomated `json:"sbfm_likely_automated"`
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
	EnableJs                     apijson.Field
	UsingLatestModel             apijson.Field
	FightMode                    apijson.Field
	OptimizeWordpress            apijson.Field
	SbfmDefinitelyAutomated      apijson.Field
	SbfmStaticResourceProtection apijson.Field
	SbfmVerifiedBots             apijson.Field
	SbfmLikelyAutomated          apijson.Field
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
type BotManagementGetResponseSbfmDefinitelyAutomated string

const (
	BotManagementGetResponseSbfmDefinitelyAutomatedAllow            BotManagementGetResponseSbfmDefinitelyAutomated = "allow"
	BotManagementGetResponseSbfmDefinitelyAutomatedBlock            BotManagementGetResponseSbfmDefinitelyAutomated = "block"
	BotManagementGetResponseSbfmDefinitelyAutomatedManagedChallenge BotManagementGetResponseSbfmDefinitelyAutomated = "managed_challenge"
)

func (r BotManagementGetResponseSbfmDefinitelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementGetResponseSbfmDefinitelyAutomatedAllow, BotManagementGetResponseSbfmDefinitelyAutomatedBlock, BotManagementGetResponseSbfmDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementGetResponseSbfmVerifiedBots string

const (
	BotManagementGetResponseSbfmVerifiedBotsAllow BotManagementGetResponseSbfmVerifiedBots = "allow"
	BotManagementGetResponseSbfmVerifiedBotsBlock BotManagementGetResponseSbfmVerifiedBots = "block"
)

func (r BotManagementGetResponseSbfmVerifiedBots) IsKnown() bool {
	switch r {
	case BotManagementGetResponseSbfmVerifiedBotsAllow, BotManagementGetResponseSbfmVerifiedBotsBlock:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementGetResponseSbfmLikelyAutomated string

const (
	BotManagementGetResponseSbfmLikelyAutomatedAllow            BotManagementGetResponseSbfmLikelyAutomated = "allow"
	BotManagementGetResponseSbfmLikelyAutomatedBlock            BotManagementGetResponseSbfmLikelyAutomated = "block"
	BotManagementGetResponseSbfmLikelyAutomatedManagedChallenge BotManagementGetResponseSbfmLikelyAutomated = "managed_challenge"
)

func (r BotManagementGetResponseSbfmLikelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementGetResponseSbfmLikelyAutomatedAllow, BotManagementGetResponseSbfmLikelyAutomatedBlock, BotManagementGetResponseSbfmLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// This interface is a union satisfied by one of the following:
// [BotManagementUpdateParamsBotFightModeConfiguration],
// [BotManagementUpdateParamsSuperBotFightModeDefinitelyConfiguration],
// [BotManagementUpdateParamsSuperBotFightModeLikelyConfiguration],
// [BotManagementUpdateParamsSubscriptionConfiguration].
type BotManagementUpdateParams interface {
	ImplementsBotManagementUpdateParams()

	getZoneID() param.Field[string]
}

type BotManagementUpdateParamsBotFightModeConfiguration struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode param.Field[bool] `json:"fight_mode"`
}

func (r BotManagementUpdateParamsBotFightModeConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BotManagementUpdateParamsBotFightModeConfiguration) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (BotManagementUpdateParamsBotFightModeConfiguration) ImplementsBotManagementUpdateParams() {

}

type BotManagementUpdateParamsSuperBotFightModeDefinitelyConfiguration struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated param.Field[BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots param.Field[BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r BotManagementUpdateParamsSuperBotFightModeDefinitelyConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BotManagementUpdateParamsSuperBotFightModeDefinitelyConfiguration) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (BotManagementUpdateParamsSuperBotFightModeDefinitelyConfiguration) ImplementsBotManagementUpdateParams() {

}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomated string

const (
	BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomatedAllow            BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomated = "allow"
	BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomatedBlock            BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomated = "block"
	BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomatedManagedChallenge BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomated = "managed_challenge"
)

func (r BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomatedAllow, BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomatedBlock, BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmVerifiedBots string

const (
	BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmVerifiedBotsAllow BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmVerifiedBots = "allow"
	BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmVerifiedBotsBlock BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmVerifiedBots = "block"
)

func (r BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmVerifiedBots) IsKnown() bool {
	switch r {
	case BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmVerifiedBotsAllow, BotManagementUpdateParamsSuperBotFightModeDefinitelyConfigurationSbfmVerifiedBotsBlock:
		return true
	}
	return false
}

type BotManagementUpdateParamsSuperBotFightModeLikelyConfiguration struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJs param.Field[bool] `json:"enable_js"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SbfmDefinitelyAutomated param.Field[BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SbfmLikelyAutomated param.Field[BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmLikelyAutomated] `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SbfmStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SbfmVerifiedBots param.Field[BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r BotManagementUpdateParamsSuperBotFightModeLikelyConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BotManagementUpdateParamsSuperBotFightModeLikelyConfiguration) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (BotManagementUpdateParamsSuperBotFightModeLikelyConfiguration) ImplementsBotManagementUpdateParams() {

}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomated string

const (
	BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomatedAllow            BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomated = "allow"
	BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomatedBlock            BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomated = "block"
	BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomatedManagedChallenge BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomated = "managed_challenge"
)

func (r BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomatedAllow, BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomatedBlock, BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmLikelyAutomated string

const (
	BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmLikelyAutomatedAllow            BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmLikelyAutomated = "allow"
	BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmLikelyAutomatedBlock            BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmLikelyAutomated = "block"
	BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmLikelyAutomatedManagedChallenge BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmLikelyAutomated = "managed_challenge"
)

func (r BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmLikelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmLikelyAutomatedAllow, BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmLikelyAutomatedBlock, BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmVerifiedBots string

const (
	BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmVerifiedBotsAllow BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmVerifiedBots = "allow"
	BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmVerifiedBotsBlock BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmVerifiedBots = "block"
)

func (r BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmVerifiedBots) IsKnown() bool {
	switch r {
	case BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmVerifiedBotsAllow, BotManagementUpdateParamsSuperBotFightModeLikelyConfigurationSbfmVerifiedBotsBlock:
		return true
	}
	return false
}

type BotManagementUpdateParamsSubscriptionConfiguration struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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

func (r BotManagementUpdateParamsSubscriptionConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BotManagementUpdateParamsSubscriptionConfiguration) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (BotManagementUpdateParamsSubscriptionConfiguration) ImplementsBotManagementUpdateParams() {

}

type BotManagementUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   BotManagementUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success BotManagementUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    botManagementUpdateResponseEnvelopeJSON    `json:"-"`
}

// botManagementUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [BotManagementUpdateResponseEnvelope]
type botManagementUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   BotManagementGetResponse `json:"result,required"`
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
