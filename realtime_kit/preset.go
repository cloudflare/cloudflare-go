// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package realtime_kit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
	"github.com/tidwall/gjson"
)

// PresetService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPresetService] method instead.
type PresetService struct {
	Options []option.RequestOption
}

// NewPresetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPresetService(opts ...option.RequestOption) (r *PresetService) {
	r = &PresetService{}
	r.Options = opts
	return
}

// Creates a preset belonging to the current App
func (r *PresetService) New(ctx context.Context, appID string, params PresetNewParams, opts ...option.RequestOption) (res *PresetNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/presets", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update a preset by the provided preset ID
func (r *PresetService) Update(ctx context.Context, appID string, presetID string, params PresetUpdateParams, opts ...option.RequestOption) (res *PresetUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if presetID == "" {
		err = errors.New("missing required preset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/presets/%s", params.AccountID, appID, presetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Deletes a preset using the provided preset ID
func (r *PresetService) Delete(ctx context.Context, appID string, presetID string, body PresetDeleteParams, opts ...option.RequestOption) (res *PresetDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if presetID == "" {
		err = errors.New("missing required preset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/presets/%s", body.AccountID, appID, presetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetches all the presets belonging to an App.
func (r *PresetService) Get(ctx context.Context, appID string, params PresetGetParams, opts ...option.RequestOption) (res *PresetGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/presets", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Fetches details of a preset using the provided preset ID
func (r *PresetService) GetPresetByID(ctx context.Context, appID string, presetID string, query PresetGetPresetByIDParams, opts ...option.RequestOption) (res *PresetGetPresetByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if presetID == "" {
		err = errors.New("missing required preset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/presets/%s", query.AccountID, appID, presetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PresetNewResponse struct {
	// Data returned by the operation
	Data PresetNewResponseData `json:"data,required"`
	// Success status of the operation
	Success bool                  `json:"success,required"`
	JSON    presetNewResponseJSON `json:"-"`
}

// presetNewResponseJSON contains the JSON metadata for the struct
// [PresetNewResponse]
type presetNewResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type PresetNewResponseData struct {
	// ID of the preset
	ID     string                      `json:"id,required" format:"uuid"`
	Config PresetNewResponseDataConfig `json:"config,required"`
	// Name of the preset
	Name        string                           `json:"name,required"`
	UI          PresetNewResponseDataUI          `json:"ui,required"`
	Permissions PresetNewResponseDataPermissions `json:"permissions"`
	JSON        presetNewResponseDataJSON        `json:"-"`
}

// presetNewResponseDataJSON contains the JSON metadata for the struct
// [PresetNewResponseData]
type presetNewResponseDataJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	UI          apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type PresetNewResponseDataConfig struct {
	// Maximum number of screen shares that can be active at a given time
	MaxScreenshareCount int64 `json:"max_screenshare_count,required"`
	// Maximum number of streams that are visible on a device
	MaxVideoStreams PresetNewResponseDataConfigMaxVideoStreams `json:"max_video_streams,required"`
	// Media configuration options. eg: Video quality
	Media PresetNewResponseDataConfigMedia `json:"media,required"`
	// Type of the meeting
	ViewType PresetNewResponseDataConfigViewType `json:"view_type,required"`
	JSON     presetNewResponseDataConfigJSON     `json:"-"`
}

// presetNewResponseDataConfigJSON contains the JSON metadata for the struct
// [PresetNewResponseDataConfig]
type presetNewResponseDataConfigJSON struct {
	MaxScreenshareCount apijson.Field
	MaxVideoStreams     apijson.Field
	Media               apijson.Field
	ViewType            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PresetNewResponseDataConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataConfigJSON) RawJSON() string {
	return r.raw
}

// Maximum number of streams that are visible on a device
type PresetNewResponseDataConfigMaxVideoStreams struct {
	// Maximum number of video streams visible on desktop devices
	Desktop int64 `json:"desktop,required"`
	// Maximum number of streams visible on mobile devices
	Mobile int64                                          `json:"mobile,required"`
	JSON   presetNewResponseDataConfigMaxVideoStreamsJSON `json:"-"`
}

// presetNewResponseDataConfigMaxVideoStreamsJSON contains the JSON metadata for
// the struct [PresetNewResponseDataConfigMaxVideoStreams]
type presetNewResponseDataConfigMaxVideoStreamsJSON struct {
	Desktop     apijson.Field
	Mobile      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataConfigMaxVideoStreams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataConfigMaxVideoStreamsJSON) RawJSON() string {
	return r.raw
}

// Media configuration options. eg: Video quality
type PresetNewResponseDataConfigMedia struct {
	// Configuration options for participant screen shares
	Screenshare PresetNewResponseDataConfigMediaScreenshare `json:"screenshare,required"`
	// Configuration options for participant videos
	Video PresetNewResponseDataConfigMediaVideo `json:"video,required"`
	// Control options for Audio quality.
	Audio PresetNewResponseDataConfigMediaAudio `json:"audio"`
	JSON  presetNewResponseDataConfigMediaJSON  `json:"-"`
}

// presetNewResponseDataConfigMediaJSON contains the JSON metadata for the struct
// [PresetNewResponseDataConfigMedia]
type presetNewResponseDataConfigMediaJSON struct {
	Screenshare apijson.Field
	Video       apijson.Field
	Audio       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataConfigMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataConfigMediaJSON) RawJSON() string {
	return r.raw
}

// Configuration options for participant screen shares
type PresetNewResponseDataConfigMediaScreenshare struct {
	// Frame rate of screen share
	FrameRate int64 `json:"frame_rate,required"`
	// Quality of screen share
	Quality PresetNewResponseDataConfigMediaScreenshareQuality `json:"quality,required"`
	JSON    presetNewResponseDataConfigMediaScreenshareJSON    `json:"-"`
}

// presetNewResponseDataConfigMediaScreenshareJSON contains the JSON metadata for
// the struct [PresetNewResponseDataConfigMediaScreenshare]
type presetNewResponseDataConfigMediaScreenshareJSON struct {
	FrameRate   apijson.Field
	Quality     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataConfigMediaScreenshare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataConfigMediaScreenshareJSON) RawJSON() string {
	return r.raw
}

// Quality of screen share
type PresetNewResponseDataConfigMediaScreenshareQuality string

const (
	PresetNewResponseDataConfigMediaScreenshareQualityHD   PresetNewResponseDataConfigMediaScreenshareQuality = "hd"
	PresetNewResponseDataConfigMediaScreenshareQualityVga  PresetNewResponseDataConfigMediaScreenshareQuality = "vga"
	PresetNewResponseDataConfigMediaScreenshareQualityQvga PresetNewResponseDataConfigMediaScreenshareQuality = "qvga"
)

func (r PresetNewResponseDataConfigMediaScreenshareQuality) IsKnown() bool {
	switch r {
	case PresetNewResponseDataConfigMediaScreenshareQualityHD, PresetNewResponseDataConfigMediaScreenshareQualityVga, PresetNewResponseDataConfigMediaScreenshareQualityQvga:
		return true
	}
	return false
}

// Configuration options for participant videos
type PresetNewResponseDataConfigMediaVideo struct {
	// Frame rate of participants' video
	FrameRate int64 `json:"frame_rate,required"`
	// Video quality of participants
	Quality PresetNewResponseDataConfigMediaVideoQuality `json:"quality,required"`
	JSON    presetNewResponseDataConfigMediaVideoJSON    `json:"-"`
}

// presetNewResponseDataConfigMediaVideoJSON contains the JSON metadata for the
// struct [PresetNewResponseDataConfigMediaVideo]
type presetNewResponseDataConfigMediaVideoJSON struct {
	FrameRate   apijson.Field
	Quality     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataConfigMediaVideo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataConfigMediaVideoJSON) RawJSON() string {
	return r.raw
}

// Video quality of participants
type PresetNewResponseDataConfigMediaVideoQuality string

const (
	PresetNewResponseDataConfigMediaVideoQualityHD   PresetNewResponseDataConfigMediaVideoQuality = "hd"
	PresetNewResponseDataConfigMediaVideoQualityVga  PresetNewResponseDataConfigMediaVideoQuality = "vga"
	PresetNewResponseDataConfigMediaVideoQualityQvga PresetNewResponseDataConfigMediaVideoQuality = "qvga"
)

func (r PresetNewResponseDataConfigMediaVideoQuality) IsKnown() bool {
	switch r {
	case PresetNewResponseDataConfigMediaVideoQualityHD, PresetNewResponseDataConfigMediaVideoQualityVga, PresetNewResponseDataConfigMediaVideoQualityQvga:
		return true
	}
	return false
}

// Control options for Audio quality.
type PresetNewResponseDataConfigMediaAudio struct {
	// Enable High Quality Audio for your meetings
	EnableHighBitrate bool `json:"enable_high_bitrate"`
	// Enable Stereo for your meetings
	EnableStereo bool                                      `json:"enable_stereo"`
	JSON         presetNewResponseDataConfigMediaAudioJSON `json:"-"`
}

// presetNewResponseDataConfigMediaAudioJSON contains the JSON metadata for the
// struct [PresetNewResponseDataConfigMediaAudio]
type presetNewResponseDataConfigMediaAudioJSON struct {
	EnableHighBitrate apijson.Field
	EnableStereo      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PresetNewResponseDataConfigMediaAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataConfigMediaAudioJSON) RawJSON() string {
	return r.raw
}

// Type of the meeting
type PresetNewResponseDataConfigViewType string

const (
	PresetNewResponseDataConfigViewTypeGroupCall PresetNewResponseDataConfigViewType = "GROUP_CALL"
	PresetNewResponseDataConfigViewTypeWebinar   PresetNewResponseDataConfigViewType = "WEBINAR"
	PresetNewResponseDataConfigViewTypeAudioRoom PresetNewResponseDataConfigViewType = "AUDIO_ROOM"
)

func (r PresetNewResponseDataConfigViewType) IsKnown() bool {
	switch r {
	case PresetNewResponseDataConfigViewTypeGroupCall, PresetNewResponseDataConfigViewTypeWebinar, PresetNewResponseDataConfigViewTypeAudioRoom:
		return true
	}
	return false
}

type PresetNewResponseDataUI struct {
	DesignTokens PresetNewResponseDataUIDesignTokens `json:"design_tokens,required"`
	ConfigDiff   interface{}                         `json:"config_diff"`
	JSON         presetNewResponseDataUIJSON         `json:"-"`
}

// presetNewResponseDataUIJSON contains the JSON metadata for the struct
// [PresetNewResponseDataUI]
type presetNewResponseDataUIJSON struct {
	DesignTokens apijson.Field
	ConfigDiff   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PresetNewResponseDataUI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataUIJSON) RawJSON() string {
	return r.raw
}

type PresetNewResponseDataUIDesignTokens struct {
	BorderRadius PresetNewResponseDataUIDesignTokensBorderRadius `json:"border_radius,required"`
	BorderWidth  PresetNewResponseDataUIDesignTokensBorderWidth  `json:"border_width,required"`
	Colors       PresetNewResponseDataUIDesignTokensColors       `json:"colors,required"`
	Logo         string                                          `json:"logo,required"`
	SpacingBase  float64                                         `json:"spacing_base,required"`
	Theme        PresetNewResponseDataUIDesignTokensTheme        `json:"theme,required"`
	JSON         presetNewResponseDataUIDesignTokensJSON         `json:"-"`
}

// presetNewResponseDataUIDesignTokensJSON contains the JSON metadata for the
// struct [PresetNewResponseDataUIDesignTokens]
type presetNewResponseDataUIDesignTokensJSON struct {
	BorderRadius apijson.Field
	BorderWidth  apijson.Field
	Colors       apijson.Field
	Logo         apijson.Field
	SpacingBase  apijson.Field
	Theme        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PresetNewResponseDataUIDesignTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataUIDesignTokensJSON) RawJSON() string {
	return r.raw
}

type PresetNewResponseDataUIDesignTokensBorderRadius string

const (
	PresetNewResponseDataUIDesignTokensBorderRadiusRounded PresetNewResponseDataUIDesignTokensBorderRadius = "rounded"
)

func (r PresetNewResponseDataUIDesignTokensBorderRadius) IsKnown() bool {
	switch r {
	case PresetNewResponseDataUIDesignTokensBorderRadiusRounded:
		return true
	}
	return false
}

type PresetNewResponseDataUIDesignTokensBorderWidth string

const (
	PresetNewResponseDataUIDesignTokensBorderWidthThin PresetNewResponseDataUIDesignTokensBorderWidth = "thin"
)

func (r PresetNewResponseDataUIDesignTokensBorderWidth) IsKnown() bool {
	switch r {
	case PresetNewResponseDataUIDesignTokensBorderWidthThin:
		return true
	}
	return false
}

type PresetNewResponseDataUIDesignTokensColors struct {
	Background  PresetNewResponseDataUIDesignTokensColorsBackground `json:"background,required"`
	Brand       PresetNewResponseDataUIDesignTokensColorsBrand      `json:"brand,required"`
	Danger      string                                              `json:"danger,required"`
	Success     string                                              `json:"success,required"`
	Text        string                                              `json:"text,required"`
	TextOnBrand string                                              `json:"text_on_brand,required"`
	VideoBg     string                                              `json:"video_bg,required"`
	Warning     string                                              `json:"warning,required"`
	JSON        presetNewResponseDataUIDesignTokensColorsJSON       `json:"-"`
}

// presetNewResponseDataUIDesignTokensColorsJSON contains the JSON metadata for the
// struct [PresetNewResponseDataUIDesignTokensColors]
type presetNewResponseDataUIDesignTokensColorsJSON struct {
	Background  apijson.Field
	Brand       apijson.Field
	Danger      apijson.Field
	Success     apijson.Field
	Text        apijson.Field
	TextOnBrand apijson.Field
	VideoBg     apijson.Field
	Warning     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataUIDesignTokensColors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataUIDesignTokensColorsJSON) RawJSON() string {
	return r.raw
}

type PresetNewResponseDataUIDesignTokensColorsBackground struct {
	Number1000 string                                                  `json:"1000,required"`
	Number600  string                                                  `json:"600,required"`
	Number700  string                                                  `json:"700,required"`
	Number800  string                                                  `json:"800,required"`
	Number900  string                                                  `json:"900,required"`
	JSON       presetNewResponseDataUIDesignTokensColorsBackgroundJSON `json:"-"`
}

// presetNewResponseDataUIDesignTokensColorsBackgroundJSON contains the JSON
// metadata for the struct [PresetNewResponseDataUIDesignTokensColorsBackground]
type presetNewResponseDataUIDesignTokensColorsBackgroundJSON struct {
	Number1000  apijson.Field
	Number600   apijson.Field
	Number700   apijson.Field
	Number800   apijson.Field
	Number900   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataUIDesignTokensColorsBackground) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataUIDesignTokensColorsBackgroundJSON) RawJSON() string {
	return r.raw
}

type PresetNewResponseDataUIDesignTokensColorsBrand struct {
	Number300 string                                             `json:"300,required"`
	Number400 string                                             `json:"400,required"`
	Number500 string                                             `json:"500,required"`
	Number600 string                                             `json:"600,required"`
	Number700 string                                             `json:"700,required"`
	JSON      presetNewResponseDataUIDesignTokensColorsBrandJSON `json:"-"`
}

// presetNewResponseDataUIDesignTokensColorsBrandJSON contains the JSON metadata
// for the struct [PresetNewResponseDataUIDesignTokensColorsBrand]
type presetNewResponseDataUIDesignTokensColorsBrandJSON struct {
	Number300   apijson.Field
	Number400   apijson.Field
	Number500   apijson.Field
	Number600   apijson.Field
	Number700   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataUIDesignTokensColorsBrand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataUIDesignTokensColorsBrandJSON) RawJSON() string {
	return r.raw
}

type PresetNewResponseDataUIDesignTokensTheme string

const (
	PresetNewResponseDataUIDesignTokensThemeDark PresetNewResponseDataUIDesignTokensTheme = "dark"
)

func (r PresetNewResponseDataUIDesignTokensTheme) IsKnown() bool {
	switch r {
	case PresetNewResponseDataUIDesignTokensThemeDark:
		return true
	}
	return false
}

type PresetNewResponseDataPermissions struct {
	// Whether this participant can accept waiting requests
	AcceptWaitingRequests           bool `json:"accept_waiting_requests,required"`
	CanAcceptProductionRequests     bool `json:"can_accept_production_requests,required"`
	CanChangeParticipantPermissions bool `json:"can_change_participant_permissions,required"`
	CanEditDisplayName              bool `json:"can_edit_display_name,required"`
	CanLivestream                   bool `json:"can_livestream,required"`
	CanRecord                       bool `json:"can_record,required"`
	CanSpotlight                    bool `json:"can_spotlight,required"`
	// Chat permissions
	Chat                            PresetNewResponseDataPermissionsChat              `json:"chat,required"`
	ConnectedMeetings               PresetNewResponseDataPermissionsConnectedMeetings `json:"connected_meetings,required"`
	DisableParticipantAudio         bool                                              `json:"disable_participant_audio,required"`
	DisableParticipantScreensharing bool                                              `json:"disable_participant_screensharing,required"`
	DisableParticipantVideo         bool                                              `json:"disable_participant_video,required"`
	// Whether this participant is visible to others or not
	HiddenParticipant bool `json:"hidden_participant,required"`
	KickParticipant   bool `json:"kick_participant,required"`
	// Media permissions
	Media          PresetNewResponseDataPermissionsMedia `json:"media,required"`
	PinParticipant bool                                  `json:"pin_participant,required"`
	// Plugin permissions
	Plugins PresetNewResponseDataPermissionsPlugins `json:"plugins,required"`
	// Poll permissions
	Polls PresetNewResponseDataPermissionsPolls `json:"polls,required"`
	// Type of the recording peer
	RecorderType        PresetNewResponseDataPermissionsRecorderType `json:"recorder_type,required"`
	ShowParticipantList bool                                         `json:"show_participant_list,required"`
	// Waiting room type
	WaitingRoomType PresetNewResponseDataPermissionsWaitingRoomType `json:"waiting_room_type,required"`
	IsRecorder      bool                                            `json:"is_recorder"`
	JSON            presetNewResponseDataPermissionsJSON            `json:"-"`
}

// presetNewResponseDataPermissionsJSON contains the JSON metadata for the struct
// [PresetNewResponseDataPermissions]
type presetNewResponseDataPermissionsJSON struct {
	AcceptWaitingRequests           apijson.Field
	CanAcceptProductionRequests     apijson.Field
	CanChangeParticipantPermissions apijson.Field
	CanEditDisplayName              apijson.Field
	CanLivestream                   apijson.Field
	CanRecord                       apijson.Field
	CanSpotlight                    apijson.Field
	Chat                            apijson.Field
	ConnectedMeetings               apijson.Field
	DisableParticipantAudio         apijson.Field
	DisableParticipantScreensharing apijson.Field
	DisableParticipantVideo         apijson.Field
	HiddenParticipant               apijson.Field
	KickParticipant                 apijson.Field
	Media                           apijson.Field
	PinParticipant                  apijson.Field
	Plugins                         apijson.Field
	Polls                           apijson.Field
	RecorderType                    apijson.Field
	ShowParticipantList             apijson.Field
	WaitingRoomType                 apijson.Field
	IsRecorder                      apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *PresetNewResponseDataPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataPermissionsJSON) RawJSON() string {
	return r.raw
}

// Chat permissions
type PresetNewResponseDataPermissionsChat struct {
	Private PresetNewResponseDataPermissionsChatPrivate `json:"private,required"`
	Public  PresetNewResponseDataPermissionsChatPublic  `json:"public,required"`
	JSON    presetNewResponseDataPermissionsChatJSON    `json:"-"`
}

// presetNewResponseDataPermissionsChatJSON contains the JSON metadata for the
// struct [PresetNewResponseDataPermissionsChat]
type presetNewResponseDataPermissionsChatJSON struct {
	Private     apijson.Field
	Public      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataPermissionsChat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataPermissionsChatJSON) RawJSON() string {
	return r.raw
}

type PresetNewResponseDataPermissionsChatPrivate struct {
	CanReceive bool                                            `json:"can_receive,required"`
	CanSend    bool                                            `json:"can_send,required"`
	Files      bool                                            `json:"files,required"`
	Text       bool                                            `json:"text,required"`
	JSON       presetNewResponseDataPermissionsChatPrivateJSON `json:"-"`
}

// presetNewResponseDataPermissionsChatPrivateJSON contains the JSON metadata for
// the struct [PresetNewResponseDataPermissionsChatPrivate]
type presetNewResponseDataPermissionsChatPrivateJSON struct {
	CanReceive  apijson.Field
	CanSend     apijson.Field
	Files       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataPermissionsChatPrivate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataPermissionsChatPrivateJSON) RawJSON() string {
	return r.raw
}

type PresetNewResponseDataPermissionsChatPublic struct {
	// Can send messages in general
	CanSend bool `json:"can_send,required"`
	// Can send file messages
	Files bool `json:"files,required"`
	// Can send text messages
	Text bool                                           `json:"text,required"`
	JSON presetNewResponseDataPermissionsChatPublicJSON `json:"-"`
}

// presetNewResponseDataPermissionsChatPublicJSON contains the JSON metadata for
// the struct [PresetNewResponseDataPermissionsChatPublic]
type presetNewResponseDataPermissionsChatPublicJSON struct {
	CanSend     apijson.Field
	Files       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataPermissionsChatPublic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataPermissionsChatPublicJSON) RawJSON() string {
	return r.raw
}

type PresetNewResponseDataPermissionsConnectedMeetings struct {
	CanAlterConnectedMeetings  bool                                                  `json:"can_alter_connected_meetings,required"`
	CanSwitchConnectedMeetings bool                                                  `json:"can_switch_connected_meetings,required"`
	CanSwitchToParentMeeting   bool                                                  `json:"can_switch_to_parent_meeting,required"`
	JSON                       presetNewResponseDataPermissionsConnectedMeetingsJSON `json:"-"`
}

// presetNewResponseDataPermissionsConnectedMeetingsJSON contains the JSON metadata
// for the struct [PresetNewResponseDataPermissionsConnectedMeetings]
type presetNewResponseDataPermissionsConnectedMeetingsJSON struct {
	CanAlterConnectedMeetings  apijson.Field
	CanSwitchConnectedMeetings apijson.Field
	CanSwitchToParentMeeting   apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *PresetNewResponseDataPermissionsConnectedMeetings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataPermissionsConnectedMeetingsJSON) RawJSON() string {
	return r.raw
}

// Media permissions
type PresetNewResponseDataPermissionsMedia struct {
	// Audio permissions
	Audio PresetNewResponseDataPermissionsMediaAudio `json:"audio,required"`
	// Screenshare permissions
	Screenshare PresetNewResponseDataPermissionsMediaScreenshare `json:"screenshare,required"`
	// Video permissions
	Video PresetNewResponseDataPermissionsMediaVideo `json:"video,required"`
	JSON  presetNewResponseDataPermissionsMediaJSON  `json:"-"`
}

// presetNewResponseDataPermissionsMediaJSON contains the JSON metadata for the
// struct [PresetNewResponseDataPermissionsMedia]
type presetNewResponseDataPermissionsMediaJSON struct {
	Audio       apijson.Field
	Screenshare apijson.Field
	Video       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataPermissionsMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataPermissionsMediaJSON) RawJSON() string {
	return r.raw
}

// Audio permissions
type PresetNewResponseDataPermissionsMediaAudio struct {
	// Can produce audio
	CanProduce PresetNewResponseDataPermissionsMediaAudioCanProduce `json:"can_produce,required"`
	JSON       presetNewResponseDataPermissionsMediaAudioJSON       `json:"-"`
}

// presetNewResponseDataPermissionsMediaAudioJSON contains the JSON metadata for
// the struct [PresetNewResponseDataPermissionsMediaAudio]
type presetNewResponseDataPermissionsMediaAudioJSON struct {
	CanProduce  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataPermissionsMediaAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataPermissionsMediaAudioJSON) RawJSON() string {
	return r.raw
}

// Can produce audio
type PresetNewResponseDataPermissionsMediaAudioCanProduce string

const (
	PresetNewResponseDataPermissionsMediaAudioCanProduceAllowed    PresetNewResponseDataPermissionsMediaAudioCanProduce = "ALLOWED"
	PresetNewResponseDataPermissionsMediaAudioCanProduceNotAllowed PresetNewResponseDataPermissionsMediaAudioCanProduce = "NOT_ALLOWED"
	PresetNewResponseDataPermissionsMediaAudioCanProduceCanRequest PresetNewResponseDataPermissionsMediaAudioCanProduce = "CAN_REQUEST"
)

func (r PresetNewResponseDataPermissionsMediaAudioCanProduce) IsKnown() bool {
	switch r {
	case PresetNewResponseDataPermissionsMediaAudioCanProduceAllowed, PresetNewResponseDataPermissionsMediaAudioCanProduceNotAllowed, PresetNewResponseDataPermissionsMediaAudioCanProduceCanRequest:
		return true
	}
	return false
}

// Screenshare permissions
type PresetNewResponseDataPermissionsMediaScreenshare struct {
	// Can produce screen share video
	CanProduce PresetNewResponseDataPermissionsMediaScreenshareCanProduce `json:"can_produce,required"`
	JSON       presetNewResponseDataPermissionsMediaScreenshareJSON       `json:"-"`
}

// presetNewResponseDataPermissionsMediaScreenshareJSON contains the JSON metadata
// for the struct [PresetNewResponseDataPermissionsMediaScreenshare]
type presetNewResponseDataPermissionsMediaScreenshareJSON struct {
	CanProduce  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataPermissionsMediaScreenshare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataPermissionsMediaScreenshareJSON) RawJSON() string {
	return r.raw
}

// Can produce screen share video
type PresetNewResponseDataPermissionsMediaScreenshareCanProduce string

const (
	PresetNewResponseDataPermissionsMediaScreenshareCanProduceAllowed    PresetNewResponseDataPermissionsMediaScreenshareCanProduce = "ALLOWED"
	PresetNewResponseDataPermissionsMediaScreenshareCanProduceNotAllowed PresetNewResponseDataPermissionsMediaScreenshareCanProduce = "NOT_ALLOWED"
	PresetNewResponseDataPermissionsMediaScreenshareCanProduceCanRequest PresetNewResponseDataPermissionsMediaScreenshareCanProduce = "CAN_REQUEST"
)

func (r PresetNewResponseDataPermissionsMediaScreenshareCanProduce) IsKnown() bool {
	switch r {
	case PresetNewResponseDataPermissionsMediaScreenshareCanProduceAllowed, PresetNewResponseDataPermissionsMediaScreenshareCanProduceNotAllowed, PresetNewResponseDataPermissionsMediaScreenshareCanProduceCanRequest:
		return true
	}
	return false
}

// Video permissions
type PresetNewResponseDataPermissionsMediaVideo struct {
	// Can produce video
	CanProduce PresetNewResponseDataPermissionsMediaVideoCanProduce `json:"can_produce,required"`
	JSON       presetNewResponseDataPermissionsMediaVideoJSON       `json:"-"`
}

// presetNewResponseDataPermissionsMediaVideoJSON contains the JSON metadata for
// the struct [PresetNewResponseDataPermissionsMediaVideo]
type presetNewResponseDataPermissionsMediaVideoJSON struct {
	CanProduce  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataPermissionsMediaVideo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataPermissionsMediaVideoJSON) RawJSON() string {
	return r.raw
}

// Can produce video
type PresetNewResponseDataPermissionsMediaVideoCanProduce string

const (
	PresetNewResponseDataPermissionsMediaVideoCanProduceAllowed    PresetNewResponseDataPermissionsMediaVideoCanProduce = "ALLOWED"
	PresetNewResponseDataPermissionsMediaVideoCanProduceNotAllowed PresetNewResponseDataPermissionsMediaVideoCanProduce = "NOT_ALLOWED"
	PresetNewResponseDataPermissionsMediaVideoCanProduceCanRequest PresetNewResponseDataPermissionsMediaVideoCanProduce = "CAN_REQUEST"
)

func (r PresetNewResponseDataPermissionsMediaVideoCanProduce) IsKnown() bool {
	switch r {
	case PresetNewResponseDataPermissionsMediaVideoCanProduceAllowed, PresetNewResponseDataPermissionsMediaVideoCanProduceNotAllowed, PresetNewResponseDataPermissionsMediaVideoCanProduceCanRequest:
		return true
	}
	return false
}

// Plugin permissions
type PresetNewResponseDataPermissionsPlugins struct {
	// Can close plugins that are already open
	CanClose bool `json:"can_close,required"`
	// Can edit plugin config
	CanEditConfig bool `json:"can_edit_config,required"`
	// Can start plugins
	CanStart bool                                               `json:"can_start,required"`
	Config   PresetNewResponseDataPermissionsPluginsConfigUnion `json:"config,required" format:"uuid"`
	JSON     presetNewResponseDataPermissionsPluginsJSON        `json:"-"`
}

// presetNewResponseDataPermissionsPluginsJSON contains the JSON metadata for the
// struct [PresetNewResponseDataPermissionsPlugins]
type presetNewResponseDataPermissionsPluginsJSON struct {
	CanClose      apijson.Field
	CanEditConfig apijson.Field
	CanStart      apijson.Field
	Config        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PresetNewResponseDataPermissionsPlugins) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataPermissionsPluginsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or
// [PresetNewResponseDataPermissionsPluginsConfigObject].
type PresetNewResponseDataPermissionsPluginsConfigUnion interface {
	ImplementsPresetNewResponseDataPermissionsPluginsConfigUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PresetNewResponseDataPermissionsPluginsConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PresetNewResponseDataPermissionsPluginsConfigObject{}),
		},
	)
}

type PresetNewResponseDataPermissionsPluginsConfigObject struct {
	AccessControl   PresetNewResponseDataPermissionsPluginsConfigObjectAccessControl `json:"access_control,required"`
	HandlesViewOnly bool                                                             `json:"handles_view_only,required"`
	JSON            presetNewResponseDataPermissionsPluginsConfigObjectJSON          `json:"-"`
}

// presetNewResponseDataPermissionsPluginsConfigObjectJSON contains the JSON
// metadata for the struct [PresetNewResponseDataPermissionsPluginsConfigObject]
type presetNewResponseDataPermissionsPluginsConfigObjectJSON struct {
	AccessControl   apijson.Field
	HandlesViewOnly apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PresetNewResponseDataPermissionsPluginsConfigObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataPermissionsPluginsConfigObjectJSON) RawJSON() string {
	return r.raw
}

func (r PresetNewResponseDataPermissionsPluginsConfigObject) ImplementsPresetNewResponseDataPermissionsPluginsConfigUnion() {
}

type PresetNewResponseDataPermissionsPluginsConfigObjectAccessControl string

const (
	PresetNewResponseDataPermissionsPluginsConfigObjectAccessControlFullAccess PresetNewResponseDataPermissionsPluginsConfigObjectAccessControl = "FULL_ACCESS"
	PresetNewResponseDataPermissionsPluginsConfigObjectAccessControlViewOnly   PresetNewResponseDataPermissionsPluginsConfigObjectAccessControl = "VIEW_ONLY"
)

func (r PresetNewResponseDataPermissionsPluginsConfigObjectAccessControl) IsKnown() bool {
	switch r {
	case PresetNewResponseDataPermissionsPluginsConfigObjectAccessControlFullAccess, PresetNewResponseDataPermissionsPluginsConfigObjectAccessControlViewOnly:
		return true
	}
	return false
}

// Poll permissions
type PresetNewResponseDataPermissionsPolls struct {
	// Can create polls
	CanCreate bool `json:"can_create,required"`
	// Can view polls
	CanView bool `json:"can_view,required"`
	// Can vote on polls
	CanVote bool                                      `json:"can_vote,required"`
	JSON    presetNewResponseDataPermissionsPollsJSON `json:"-"`
}

// presetNewResponseDataPermissionsPollsJSON contains the JSON metadata for the
// struct [PresetNewResponseDataPermissionsPolls]
type presetNewResponseDataPermissionsPollsJSON struct {
	CanCreate   apijson.Field
	CanView     apijson.Field
	CanVote     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetNewResponseDataPermissionsPolls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetNewResponseDataPermissionsPollsJSON) RawJSON() string {
	return r.raw
}

// Type of the recording peer
type PresetNewResponseDataPermissionsRecorderType string

const (
	PresetNewResponseDataPermissionsRecorderTypeRecorder     PresetNewResponseDataPermissionsRecorderType = "RECORDER"
	PresetNewResponseDataPermissionsRecorderTypeLivestreamer PresetNewResponseDataPermissionsRecorderType = "LIVESTREAMER"
	PresetNewResponseDataPermissionsRecorderTypeNone         PresetNewResponseDataPermissionsRecorderType = "NONE"
)

func (r PresetNewResponseDataPermissionsRecorderType) IsKnown() bool {
	switch r {
	case PresetNewResponseDataPermissionsRecorderTypeRecorder, PresetNewResponseDataPermissionsRecorderTypeLivestreamer, PresetNewResponseDataPermissionsRecorderTypeNone:
		return true
	}
	return false
}

// Waiting room type
type PresetNewResponseDataPermissionsWaitingRoomType string

const (
	PresetNewResponseDataPermissionsWaitingRoomTypeSkip                  PresetNewResponseDataPermissionsWaitingRoomType = "SKIP"
	PresetNewResponseDataPermissionsWaitingRoomTypeOnPrivilegedUserEntry PresetNewResponseDataPermissionsWaitingRoomType = "ON_PRIVILEGED_USER_ENTRY"
	PresetNewResponseDataPermissionsWaitingRoomTypeSkipOnAccept          PresetNewResponseDataPermissionsWaitingRoomType = "SKIP_ON_ACCEPT"
)

func (r PresetNewResponseDataPermissionsWaitingRoomType) IsKnown() bool {
	switch r {
	case PresetNewResponseDataPermissionsWaitingRoomTypeSkip, PresetNewResponseDataPermissionsWaitingRoomTypeOnPrivilegedUserEntry, PresetNewResponseDataPermissionsWaitingRoomTypeSkipOnAccept:
		return true
	}
	return false
}

type PresetUpdateResponse struct {
	// Data returned by the operation
	Data PresetUpdateResponseData `json:"data,required"`
	// Success status of the operation
	Success bool                     `json:"success,required"`
	JSON    presetUpdateResponseJSON `json:"-"`
}

// presetUpdateResponseJSON contains the JSON metadata for the struct
// [PresetUpdateResponse]
type presetUpdateResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type PresetUpdateResponseData struct {
	// ID of the preset
	ID     string                         `json:"id,required" format:"uuid"`
	Config PresetUpdateResponseDataConfig `json:"config,required"`
	// Name of the preset
	Name        string                              `json:"name,required"`
	UI          PresetUpdateResponseDataUI          `json:"ui,required"`
	Permissions PresetUpdateResponseDataPermissions `json:"permissions"`
	JSON        presetUpdateResponseDataJSON        `json:"-"`
}

// presetUpdateResponseDataJSON contains the JSON metadata for the struct
// [PresetUpdateResponseData]
type presetUpdateResponseDataJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	UI          apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataJSON) RawJSON() string {
	return r.raw
}

type PresetUpdateResponseDataConfig struct {
	// Maximum number of screen shares that can be active at a given time
	MaxScreenshareCount int64 `json:"max_screenshare_count,required"`
	// Maximum number of streams that are visible on a device
	MaxVideoStreams PresetUpdateResponseDataConfigMaxVideoStreams `json:"max_video_streams,required"`
	// Media configuration options. eg: Video quality
	Media PresetUpdateResponseDataConfigMedia `json:"media,required"`
	// Type of the meeting
	ViewType PresetUpdateResponseDataConfigViewType `json:"view_type,required"`
	JSON     presetUpdateResponseDataConfigJSON     `json:"-"`
}

// presetUpdateResponseDataConfigJSON contains the JSON metadata for the struct
// [PresetUpdateResponseDataConfig]
type presetUpdateResponseDataConfigJSON struct {
	MaxScreenshareCount apijson.Field
	MaxVideoStreams     apijson.Field
	Media               apijson.Field
	ViewType            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PresetUpdateResponseDataConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataConfigJSON) RawJSON() string {
	return r.raw
}

// Maximum number of streams that are visible on a device
type PresetUpdateResponseDataConfigMaxVideoStreams struct {
	// Maximum number of video streams visible on desktop devices
	Desktop int64 `json:"desktop,required"`
	// Maximum number of streams visible on mobile devices
	Mobile int64                                             `json:"mobile,required"`
	JSON   presetUpdateResponseDataConfigMaxVideoStreamsJSON `json:"-"`
}

// presetUpdateResponseDataConfigMaxVideoStreamsJSON contains the JSON metadata for
// the struct [PresetUpdateResponseDataConfigMaxVideoStreams]
type presetUpdateResponseDataConfigMaxVideoStreamsJSON struct {
	Desktop     apijson.Field
	Mobile      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataConfigMaxVideoStreams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataConfigMaxVideoStreamsJSON) RawJSON() string {
	return r.raw
}

// Media configuration options. eg: Video quality
type PresetUpdateResponseDataConfigMedia struct {
	// Configuration options for participant screen shares
	Screenshare PresetUpdateResponseDataConfigMediaScreenshare `json:"screenshare,required"`
	// Configuration options for participant videos
	Video PresetUpdateResponseDataConfigMediaVideo `json:"video,required"`
	// Control options for Audio quality.
	Audio PresetUpdateResponseDataConfigMediaAudio `json:"audio"`
	JSON  presetUpdateResponseDataConfigMediaJSON  `json:"-"`
}

// presetUpdateResponseDataConfigMediaJSON contains the JSON metadata for the
// struct [PresetUpdateResponseDataConfigMedia]
type presetUpdateResponseDataConfigMediaJSON struct {
	Screenshare apijson.Field
	Video       apijson.Field
	Audio       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataConfigMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataConfigMediaJSON) RawJSON() string {
	return r.raw
}

// Configuration options for participant screen shares
type PresetUpdateResponseDataConfigMediaScreenshare struct {
	// Frame rate of screen share
	FrameRate int64 `json:"frame_rate,required"`
	// Quality of screen share
	Quality PresetUpdateResponseDataConfigMediaScreenshareQuality `json:"quality,required"`
	JSON    presetUpdateResponseDataConfigMediaScreenshareJSON    `json:"-"`
}

// presetUpdateResponseDataConfigMediaScreenshareJSON contains the JSON metadata
// for the struct [PresetUpdateResponseDataConfigMediaScreenshare]
type presetUpdateResponseDataConfigMediaScreenshareJSON struct {
	FrameRate   apijson.Field
	Quality     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataConfigMediaScreenshare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataConfigMediaScreenshareJSON) RawJSON() string {
	return r.raw
}

// Quality of screen share
type PresetUpdateResponseDataConfigMediaScreenshareQuality string

const (
	PresetUpdateResponseDataConfigMediaScreenshareQualityHD   PresetUpdateResponseDataConfigMediaScreenshareQuality = "hd"
	PresetUpdateResponseDataConfigMediaScreenshareQualityVga  PresetUpdateResponseDataConfigMediaScreenshareQuality = "vga"
	PresetUpdateResponseDataConfigMediaScreenshareQualityQvga PresetUpdateResponseDataConfigMediaScreenshareQuality = "qvga"
)

func (r PresetUpdateResponseDataConfigMediaScreenshareQuality) IsKnown() bool {
	switch r {
	case PresetUpdateResponseDataConfigMediaScreenshareQualityHD, PresetUpdateResponseDataConfigMediaScreenshareQualityVga, PresetUpdateResponseDataConfigMediaScreenshareQualityQvga:
		return true
	}
	return false
}

// Configuration options for participant videos
type PresetUpdateResponseDataConfigMediaVideo struct {
	// Frame rate of participants' video
	FrameRate int64 `json:"frame_rate,required"`
	// Video quality of participants
	Quality PresetUpdateResponseDataConfigMediaVideoQuality `json:"quality,required"`
	JSON    presetUpdateResponseDataConfigMediaVideoJSON    `json:"-"`
}

// presetUpdateResponseDataConfigMediaVideoJSON contains the JSON metadata for the
// struct [PresetUpdateResponseDataConfigMediaVideo]
type presetUpdateResponseDataConfigMediaVideoJSON struct {
	FrameRate   apijson.Field
	Quality     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataConfigMediaVideo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataConfigMediaVideoJSON) RawJSON() string {
	return r.raw
}

// Video quality of participants
type PresetUpdateResponseDataConfigMediaVideoQuality string

const (
	PresetUpdateResponseDataConfigMediaVideoQualityHD   PresetUpdateResponseDataConfigMediaVideoQuality = "hd"
	PresetUpdateResponseDataConfigMediaVideoQualityVga  PresetUpdateResponseDataConfigMediaVideoQuality = "vga"
	PresetUpdateResponseDataConfigMediaVideoQualityQvga PresetUpdateResponseDataConfigMediaVideoQuality = "qvga"
)

func (r PresetUpdateResponseDataConfigMediaVideoQuality) IsKnown() bool {
	switch r {
	case PresetUpdateResponseDataConfigMediaVideoQualityHD, PresetUpdateResponseDataConfigMediaVideoQualityVga, PresetUpdateResponseDataConfigMediaVideoQualityQvga:
		return true
	}
	return false
}

// Control options for Audio quality.
type PresetUpdateResponseDataConfigMediaAudio struct {
	// Enable High Quality Audio for your meetings
	EnableHighBitrate bool `json:"enable_high_bitrate"`
	// Enable Stereo for your meetings
	EnableStereo bool                                         `json:"enable_stereo"`
	JSON         presetUpdateResponseDataConfigMediaAudioJSON `json:"-"`
}

// presetUpdateResponseDataConfigMediaAudioJSON contains the JSON metadata for the
// struct [PresetUpdateResponseDataConfigMediaAudio]
type presetUpdateResponseDataConfigMediaAudioJSON struct {
	EnableHighBitrate apijson.Field
	EnableStereo      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PresetUpdateResponseDataConfigMediaAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataConfigMediaAudioJSON) RawJSON() string {
	return r.raw
}

// Type of the meeting
type PresetUpdateResponseDataConfigViewType string

const (
	PresetUpdateResponseDataConfigViewTypeGroupCall PresetUpdateResponseDataConfigViewType = "GROUP_CALL"
	PresetUpdateResponseDataConfigViewTypeWebinar   PresetUpdateResponseDataConfigViewType = "WEBINAR"
	PresetUpdateResponseDataConfigViewTypeAudioRoom PresetUpdateResponseDataConfigViewType = "AUDIO_ROOM"
)

func (r PresetUpdateResponseDataConfigViewType) IsKnown() bool {
	switch r {
	case PresetUpdateResponseDataConfigViewTypeGroupCall, PresetUpdateResponseDataConfigViewTypeWebinar, PresetUpdateResponseDataConfigViewTypeAudioRoom:
		return true
	}
	return false
}

type PresetUpdateResponseDataUI struct {
	DesignTokens PresetUpdateResponseDataUIDesignTokens `json:"design_tokens,required"`
	ConfigDiff   interface{}                            `json:"config_diff"`
	JSON         presetUpdateResponseDataUIJSON         `json:"-"`
}

// presetUpdateResponseDataUIJSON contains the JSON metadata for the struct
// [PresetUpdateResponseDataUI]
type presetUpdateResponseDataUIJSON struct {
	DesignTokens apijson.Field
	ConfigDiff   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PresetUpdateResponseDataUI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataUIJSON) RawJSON() string {
	return r.raw
}

type PresetUpdateResponseDataUIDesignTokens struct {
	BorderRadius PresetUpdateResponseDataUIDesignTokensBorderRadius `json:"border_radius,required"`
	BorderWidth  PresetUpdateResponseDataUIDesignTokensBorderWidth  `json:"border_width,required"`
	Colors       PresetUpdateResponseDataUIDesignTokensColors       `json:"colors,required"`
	Logo         string                                             `json:"logo,required"`
	SpacingBase  float64                                            `json:"spacing_base,required"`
	Theme        PresetUpdateResponseDataUIDesignTokensTheme        `json:"theme,required"`
	JSON         presetUpdateResponseDataUIDesignTokensJSON         `json:"-"`
}

// presetUpdateResponseDataUIDesignTokensJSON contains the JSON metadata for the
// struct [PresetUpdateResponseDataUIDesignTokens]
type presetUpdateResponseDataUIDesignTokensJSON struct {
	BorderRadius apijson.Field
	BorderWidth  apijson.Field
	Colors       apijson.Field
	Logo         apijson.Field
	SpacingBase  apijson.Field
	Theme        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PresetUpdateResponseDataUIDesignTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataUIDesignTokensJSON) RawJSON() string {
	return r.raw
}

type PresetUpdateResponseDataUIDesignTokensBorderRadius string

const (
	PresetUpdateResponseDataUIDesignTokensBorderRadiusRounded PresetUpdateResponseDataUIDesignTokensBorderRadius = "rounded"
)

func (r PresetUpdateResponseDataUIDesignTokensBorderRadius) IsKnown() bool {
	switch r {
	case PresetUpdateResponseDataUIDesignTokensBorderRadiusRounded:
		return true
	}
	return false
}

type PresetUpdateResponseDataUIDesignTokensBorderWidth string

const (
	PresetUpdateResponseDataUIDesignTokensBorderWidthThin PresetUpdateResponseDataUIDesignTokensBorderWidth = "thin"
)

func (r PresetUpdateResponseDataUIDesignTokensBorderWidth) IsKnown() bool {
	switch r {
	case PresetUpdateResponseDataUIDesignTokensBorderWidthThin:
		return true
	}
	return false
}

type PresetUpdateResponseDataUIDesignTokensColors struct {
	Background  PresetUpdateResponseDataUIDesignTokensColorsBackground `json:"background,required"`
	Brand       PresetUpdateResponseDataUIDesignTokensColorsBrand      `json:"brand,required"`
	Danger      string                                                 `json:"danger,required"`
	Success     string                                                 `json:"success,required"`
	Text        string                                                 `json:"text,required"`
	TextOnBrand string                                                 `json:"text_on_brand,required"`
	VideoBg     string                                                 `json:"video_bg,required"`
	Warning     string                                                 `json:"warning,required"`
	JSON        presetUpdateResponseDataUIDesignTokensColorsJSON       `json:"-"`
}

// presetUpdateResponseDataUIDesignTokensColorsJSON contains the JSON metadata for
// the struct [PresetUpdateResponseDataUIDesignTokensColors]
type presetUpdateResponseDataUIDesignTokensColorsJSON struct {
	Background  apijson.Field
	Brand       apijson.Field
	Danger      apijson.Field
	Success     apijson.Field
	Text        apijson.Field
	TextOnBrand apijson.Field
	VideoBg     apijson.Field
	Warning     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataUIDesignTokensColors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataUIDesignTokensColorsJSON) RawJSON() string {
	return r.raw
}

type PresetUpdateResponseDataUIDesignTokensColorsBackground struct {
	Number1000 string                                                     `json:"1000,required"`
	Number600  string                                                     `json:"600,required"`
	Number700  string                                                     `json:"700,required"`
	Number800  string                                                     `json:"800,required"`
	Number900  string                                                     `json:"900,required"`
	JSON       presetUpdateResponseDataUIDesignTokensColorsBackgroundJSON `json:"-"`
}

// presetUpdateResponseDataUIDesignTokensColorsBackgroundJSON contains the JSON
// metadata for the struct [PresetUpdateResponseDataUIDesignTokensColorsBackground]
type presetUpdateResponseDataUIDesignTokensColorsBackgroundJSON struct {
	Number1000  apijson.Field
	Number600   apijson.Field
	Number700   apijson.Field
	Number800   apijson.Field
	Number900   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataUIDesignTokensColorsBackground) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataUIDesignTokensColorsBackgroundJSON) RawJSON() string {
	return r.raw
}

type PresetUpdateResponseDataUIDesignTokensColorsBrand struct {
	Number300 string                                                `json:"300,required"`
	Number400 string                                                `json:"400,required"`
	Number500 string                                                `json:"500,required"`
	Number600 string                                                `json:"600,required"`
	Number700 string                                                `json:"700,required"`
	JSON      presetUpdateResponseDataUIDesignTokensColorsBrandJSON `json:"-"`
}

// presetUpdateResponseDataUIDesignTokensColorsBrandJSON contains the JSON metadata
// for the struct [PresetUpdateResponseDataUIDesignTokensColorsBrand]
type presetUpdateResponseDataUIDesignTokensColorsBrandJSON struct {
	Number300   apijson.Field
	Number400   apijson.Field
	Number500   apijson.Field
	Number600   apijson.Field
	Number700   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataUIDesignTokensColorsBrand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataUIDesignTokensColorsBrandJSON) RawJSON() string {
	return r.raw
}

type PresetUpdateResponseDataUIDesignTokensTheme string

const (
	PresetUpdateResponseDataUIDesignTokensThemeDark PresetUpdateResponseDataUIDesignTokensTheme = "dark"
)

func (r PresetUpdateResponseDataUIDesignTokensTheme) IsKnown() bool {
	switch r {
	case PresetUpdateResponseDataUIDesignTokensThemeDark:
		return true
	}
	return false
}

type PresetUpdateResponseDataPermissions struct {
	// Whether this participant can accept waiting requests
	AcceptWaitingRequests           bool `json:"accept_waiting_requests,required"`
	CanAcceptProductionRequests     bool `json:"can_accept_production_requests,required"`
	CanChangeParticipantPermissions bool `json:"can_change_participant_permissions,required"`
	CanEditDisplayName              bool `json:"can_edit_display_name,required"`
	CanLivestream                   bool `json:"can_livestream,required"`
	CanRecord                       bool `json:"can_record,required"`
	CanSpotlight                    bool `json:"can_spotlight,required"`
	// Chat permissions
	Chat                            PresetUpdateResponseDataPermissionsChat              `json:"chat,required"`
	ConnectedMeetings               PresetUpdateResponseDataPermissionsConnectedMeetings `json:"connected_meetings,required"`
	DisableParticipantAudio         bool                                                 `json:"disable_participant_audio,required"`
	DisableParticipantScreensharing bool                                                 `json:"disable_participant_screensharing,required"`
	DisableParticipantVideo         bool                                                 `json:"disable_participant_video,required"`
	// Whether this participant is visible to others or not
	HiddenParticipant bool `json:"hidden_participant,required"`
	KickParticipant   bool `json:"kick_participant,required"`
	// Media permissions
	Media          PresetUpdateResponseDataPermissionsMedia `json:"media,required"`
	PinParticipant bool                                     `json:"pin_participant,required"`
	// Plugin permissions
	Plugins PresetUpdateResponseDataPermissionsPlugins `json:"plugins,required"`
	// Poll permissions
	Polls PresetUpdateResponseDataPermissionsPolls `json:"polls,required"`
	// Type of the recording peer
	RecorderType        PresetUpdateResponseDataPermissionsRecorderType `json:"recorder_type,required"`
	ShowParticipantList bool                                            `json:"show_participant_list,required"`
	// Waiting room type
	WaitingRoomType PresetUpdateResponseDataPermissionsWaitingRoomType `json:"waiting_room_type,required"`
	IsRecorder      bool                                               `json:"is_recorder"`
	JSON            presetUpdateResponseDataPermissionsJSON            `json:"-"`
}

// presetUpdateResponseDataPermissionsJSON contains the JSON metadata for the
// struct [PresetUpdateResponseDataPermissions]
type presetUpdateResponseDataPermissionsJSON struct {
	AcceptWaitingRequests           apijson.Field
	CanAcceptProductionRequests     apijson.Field
	CanChangeParticipantPermissions apijson.Field
	CanEditDisplayName              apijson.Field
	CanLivestream                   apijson.Field
	CanRecord                       apijson.Field
	CanSpotlight                    apijson.Field
	Chat                            apijson.Field
	ConnectedMeetings               apijson.Field
	DisableParticipantAudio         apijson.Field
	DisableParticipantScreensharing apijson.Field
	DisableParticipantVideo         apijson.Field
	HiddenParticipant               apijson.Field
	KickParticipant                 apijson.Field
	Media                           apijson.Field
	PinParticipant                  apijson.Field
	Plugins                         apijson.Field
	Polls                           apijson.Field
	RecorderType                    apijson.Field
	ShowParticipantList             apijson.Field
	WaitingRoomType                 apijson.Field
	IsRecorder                      apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *PresetUpdateResponseDataPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataPermissionsJSON) RawJSON() string {
	return r.raw
}

// Chat permissions
type PresetUpdateResponseDataPermissionsChat struct {
	Private PresetUpdateResponseDataPermissionsChatPrivate `json:"private,required"`
	Public  PresetUpdateResponseDataPermissionsChatPublic  `json:"public,required"`
	JSON    presetUpdateResponseDataPermissionsChatJSON    `json:"-"`
}

// presetUpdateResponseDataPermissionsChatJSON contains the JSON metadata for the
// struct [PresetUpdateResponseDataPermissionsChat]
type presetUpdateResponseDataPermissionsChatJSON struct {
	Private     apijson.Field
	Public      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataPermissionsChat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataPermissionsChatJSON) RawJSON() string {
	return r.raw
}

type PresetUpdateResponseDataPermissionsChatPrivate struct {
	CanReceive bool                                               `json:"can_receive,required"`
	CanSend    bool                                               `json:"can_send,required"`
	Files      bool                                               `json:"files,required"`
	Text       bool                                               `json:"text,required"`
	JSON       presetUpdateResponseDataPermissionsChatPrivateJSON `json:"-"`
}

// presetUpdateResponseDataPermissionsChatPrivateJSON contains the JSON metadata
// for the struct [PresetUpdateResponseDataPermissionsChatPrivate]
type presetUpdateResponseDataPermissionsChatPrivateJSON struct {
	CanReceive  apijson.Field
	CanSend     apijson.Field
	Files       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataPermissionsChatPrivate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataPermissionsChatPrivateJSON) RawJSON() string {
	return r.raw
}

type PresetUpdateResponseDataPermissionsChatPublic struct {
	// Can send messages in general
	CanSend bool `json:"can_send,required"`
	// Can send file messages
	Files bool `json:"files,required"`
	// Can send text messages
	Text bool                                              `json:"text,required"`
	JSON presetUpdateResponseDataPermissionsChatPublicJSON `json:"-"`
}

// presetUpdateResponseDataPermissionsChatPublicJSON contains the JSON metadata for
// the struct [PresetUpdateResponseDataPermissionsChatPublic]
type presetUpdateResponseDataPermissionsChatPublicJSON struct {
	CanSend     apijson.Field
	Files       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataPermissionsChatPublic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataPermissionsChatPublicJSON) RawJSON() string {
	return r.raw
}

type PresetUpdateResponseDataPermissionsConnectedMeetings struct {
	CanAlterConnectedMeetings  bool                                                     `json:"can_alter_connected_meetings,required"`
	CanSwitchConnectedMeetings bool                                                     `json:"can_switch_connected_meetings,required"`
	CanSwitchToParentMeeting   bool                                                     `json:"can_switch_to_parent_meeting,required"`
	JSON                       presetUpdateResponseDataPermissionsConnectedMeetingsJSON `json:"-"`
}

// presetUpdateResponseDataPermissionsConnectedMeetingsJSON contains the JSON
// metadata for the struct [PresetUpdateResponseDataPermissionsConnectedMeetings]
type presetUpdateResponseDataPermissionsConnectedMeetingsJSON struct {
	CanAlterConnectedMeetings  apijson.Field
	CanSwitchConnectedMeetings apijson.Field
	CanSwitchToParentMeeting   apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *PresetUpdateResponseDataPermissionsConnectedMeetings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataPermissionsConnectedMeetingsJSON) RawJSON() string {
	return r.raw
}

// Media permissions
type PresetUpdateResponseDataPermissionsMedia struct {
	// Audio permissions
	Audio PresetUpdateResponseDataPermissionsMediaAudio `json:"audio,required"`
	// Screenshare permissions
	Screenshare PresetUpdateResponseDataPermissionsMediaScreenshare `json:"screenshare,required"`
	// Video permissions
	Video PresetUpdateResponseDataPermissionsMediaVideo `json:"video,required"`
	JSON  presetUpdateResponseDataPermissionsMediaJSON  `json:"-"`
}

// presetUpdateResponseDataPermissionsMediaJSON contains the JSON metadata for the
// struct [PresetUpdateResponseDataPermissionsMedia]
type presetUpdateResponseDataPermissionsMediaJSON struct {
	Audio       apijson.Field
	Screenshare apijson.Field
	Video       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataPermissionsMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataPermissionsMediaJSON) RawJSON() string {
	return r.raw
}

// Audio permissions
type PresetUpdateResponseDataPermissionsMediaAudio struct {
	// Can produce audio
	CanProduce PresetUpdateResponseDataPermissionsMediaAudioCanProduce `json:"can_produce,required"`
	JSON       presetUpdateResponseDataPermissionsMediaAudioJSON       `json:"-"`
}

// presetUpdateResponseDataPermissionsMediaAudioJSON contains the JSON metadata for
// the struct [PresetUpdateResponseDataPermissionsMediaAudio]
type presetUpdateResponseDataPermissionsMediaAudioJSON struct {
	CanProduce  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataPermissionsMediaAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataPermissionsMediaAudioJSON) RawJSON() string {
	return r.raw
}

// Can produce audio
type PresetUpdateResponseDataPermissionsMediaAudioCanProduce string

const (
	PresetUpdateResponseDataPermissionsMediaAudioCanProduceAllowed    PresetUpdateResponseDataPermissionsMediaAudioCanProduce = "ALLOWED"
	PresetUpdateResponseDataPermissionsMediaAudioCanProduceNotAllowed PresetUpdateResponseDataPermissionsMediaAudioCanProduce = "NOT_ALLOWED"
	PresetUpdateResponseDataPermissionsMediaAudioCanProduceCanRequest PresetUpdateResponseDataPermissionsMediaAudioCanProduce = "CAN_REQUEST"
)

func (r PresetUpdateResponseDataPermissionsMediaAudioCanProduce) IsKnown() bool {
	switch r {
	case PresetUpdateResponseDataPermissionsMediaAudioCanProduceAllowed, PresetUpdateResponseDataPermissionsMediaAudioCanProduceNotAllowed, PresetUpdateResponseDataPermissionsMediaAudioCanProduceCanRequest:
		return true
	}
	return false
}

// Screenshare permissions
type PresetUpdateResponseDataPermissionsMediaScreenshare struct {
	// Can produce screen share video
	CanProduce PresetUpdateResponseDataPermissionsMediaScreenshareCanProduce `json:"can_produce,required"`
	JSON       presetUpdateResponseDataPermissionsMediaScreenshareJSON       `json:"-"`
}

// presetUpdateResponseDataPermissionsMediaScreenshareJSON contains the JSON
// metadata for the struct [PresetUpdateResponseDataPermissionsMediaScreenshare]
type presetUpdateResponseDataPermissionsMediaScreenshareJSON struct {
	CanProduce  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataPermissionsMediaScreenshare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataPermissionsMediaScreenshareJSON) RawJSON() string {
	return r.raw
}

// Can produce screen share video
type PresetUpdateResponseDataPermissionsMediaScreenshareCanProduce string

const (
	PresetUpdateResponseDataPermissionsMediaScreenshareCanProduceAllowed    PresetUpdateResponseDataPermissionsMediaScreenshareCanProduce = "ALLOWED"
	PresetUpdateResponseDataPermissionsMediaScreenshareCanProduceNotAllowed PresetUpdateResponseDataPermissionsMediaScreenshareCanProduce = "NOT_ALLOWED"
	PresetUpdateResponseDataPermissionsMediaScreenshareCanProduceCanRequest PresetUpdateResponseDataPermissionsMediaScreenshareCanProduce = "CAN_REQUEST"
)

func (r PresetUpdateResponseDataPermissionsMediaScreenshareCanProduce) IsKnown() bool {
	switch r {
	case PresetUpdateResponseDataPermissionsMediaScreenshareCanProduceAllowed, PresetUpdateResponseDataPermissionsMediaScreenshareCanProduceNotAllowed, PresetUpdateResponseDataPermissionsMediaScreenshareCanProduceCanRequest:
		return true
	}
	return false
}

// Video permissions
type PresetUpdateResponseDataPermissionsMediaVideo struct {
	// Can produce video
	CanProduce PresetUpdateResponseDataPermissionsMediaVideoCanProduce `json:"can_produce,required"`
	JSON       presetUpdateResponseDataPermissionsMediaVideoJSON       `json:"-"`
}

// presetUpdateResponseDataPermissionsMediaVideoJSON contains the JSON metadata for
// the struct [PresetUpdateResponseDataPermissionsMediaVideo]
type presetUpdateResponseDataPermissionsMediaVideoJSON struct {
	CanProduce  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataPermissionsMediaVideo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataPermissionsMediaVideoJSON) RawJSON() string {
	return r.raw
}

// Can produce video
type PresetUpdateResponseDataPermissionsMediaVideoCanProduce string

const (
	PresetUpdateResponseDataPermissionsMediaVideoCanProduceAllowed    PresetUpdateResponseDataPermissionsMediaVideoCanProduce = "ALLOWED"
	PresetUpdateResponseDataPermissionsMediaVideoCanProduceNotAllowed PresetUpdateResponseDataPermissionsMediaVideoCanProduce = "NOT_ALLOWED"
	PresetUpdateResponseDataPermissionsMediaVideoCanProduceCanRequest PresetUpdateResponseDataPermissionsMediaVideoCanProduce = "CAN_REQUEST"
)

func (r PresetUpdateResponseDataPermissionsMediaVideoCanProduce) IsKnown() bool {
	switch r {
	case PresetUpdateResponseDataPermissionsMediaVideoCanProduceAllowed, PresetUpdateResponseDataPermissionsMediaVideoCanProduceNotAllowed, PresetUpdateResponseDataPermissionsMediaVideoCanProduceCanRequest:
		return true
	}
	return false
}

// Plugin permissions
type PresetUpdateResponseDataPermissionsPlugins struct {
	// Can close plugins that are already open
	CanClose bool `json:"can_close,required"`
	// Can edit plugin config
	CanEditConfig bool `json:"can_edit_config,required"`
	// Can start plugins
	CanStart bool                                                  `json:"can_start,required"`
	Config   PresetUpdateResponseDataPermissionsPluginsConfigUnion `json:"config,required" format:"uuid"`
	JSON     presetUpdateResponseDataPermissionsPluginsJSON        `json:"-"`
}

// presetUpdateResponseDataPermissionsPluginsJSON contains the JSON metadata for
// the struct [PresetUpdateResponseDataPermissionsPlugins]
type presetUpdateResponseDataPermissionsPluginsJSON struct {
	CanClose      apijson.Field
	CanEditConfig apijson.Field
	CanStart      apijson.Field
	Config        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PresetUpdateResponseDataPermissionsPlugins) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataPermissionsPluginsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or
// [PresetUpdateResponseDataPermissionsPluginsConfigObject].
type PresetUpdateResponseDataPermissionsPluginsConfigUnion interface {
	ImplementsPresetUpdateResponseDataPermissionsPluginsConfigUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PresetUpdateResponseDataPermissionsPluginsConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PresetUpdateResponseDataPermissionsPluginsConfigObject{}),
		},
	)
}

type PresetUpdateResponseDataPermissionsPluginsConfigObject struct {
	AccessControl   PresetUpdateResponseDataPermissionsPluginsConfigObjectAccessControl `json:"access_control,required"`
	HandlesViewOnly bool                                                                `json:"handles_view_only,required"`
	JSON            presetUpdateResponseDataPermissionsPluginsConfigObjectJSON          `json:"-"`
}

// presetUpdateResponseDataPermissionsPluginsConfigObjectJSON contains the JSON
// metadata for the struct [PresetUpdateResponseDataPermissionsPluginsConfigObject]
type presetUpdateResponseDataPermissionsPluginsConfigObjectJSON struct {
	AccessControl   apijson.Field
	HandlesViewOnly apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PresetUpdateResponseDataPermissionsPluginsConfigObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataPermissionsPluginsConfigObjectJSON) RawJSON() string {
	return r.raw
}

func (r PresetUpdateResponseDataPermissionsPluginsConfigObject) ImplementsPresetUpdateResponseDataPermissionsPluginsConfigUnion() {
}

type PresetUpdateResponseDataPermissionsPluginsConfigObjectAccessControl string

const (
	PresetUpdateResponseDataPermissionsPluginsConfigObjectAccessControlFullAccess PresetUpdateResponseDataPermissionsPluginsConfigObjectAccessControl = "FULL_ACCESS"
	PresetUpdateResponseDataPermissionsPluginsConfigObjectAccessControlViewOnly   PresetUpdateResponseDataPermissionsPluginsConfigObjectAccessControl = "VIEW_ONLY"
)

func (r PresetUpdateResponseDataPermissionsPluginsConfigObjectAccessControl) IsKnown() bool {
	switch r {
	case PresetUpdateResponseDataPermissionsPluginsConfigObjectAccessControlFullAccess, PresetUpdateResponseDataPermissionsPluginsConfigObjectAccessControlViewOnly:
		return true
	}
	return false
}

// Poll permissions
type PresetUpdateResponseDataPermissionsPolls struct {
	// Can create polls
	CanCreate bool `json:"can_create,required"`
	// Can view polls
	CanView bool `json:"can_view,required"`
	// Can vote on polls
	CanVote bool                                         `json:"can_vote,required"`
	JSON    presetUpdateResponseDataPermissionsPollsJSON `json:"-"`
}

// presetUpdateResponseDataPermissionsPollsJSON contains the JSON metadata for the
// struct [PresetUpdateResponseDataPermissionsPolls]
type presetUpdateResponseDataPermissionsPollsJSON struct {
	CanCreate   apijson.Field
	CanView     apijson.Field
	CanVote     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetUpdateResponseDataPermissionsPolls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetUpdateResponseDataPermissionsPollsJSON) RawJSON() string {
	return r.raw
}

// Type of the recording peer
type PresetUpdateResponseDataPermissionsRecorderType string

const (
	PresetUpdateResponseDataPermissionsRecorderTypeRecorder     PresetUpdateResponseDataPermissionsRecorderType = "RECORDER"
	PresetUpdateResponseDataPermissionsRecorderTypeLivestreamer PresetUpdateResponseDataPermissionsRecorderType = "LIVESTREAMER"
	PresetUpdateResponseDataPermissionsRecorderTypeNone         PresetUpdateResponseDataPermissionsRecorderType = "NONE"
)

func (r PresetUpdateResponseDataPermissionsRecorderType) IsKnown() bool {
	switch r {
	case PresetUpdateResponseDataPermissionsRecorderTypeRecorder, PresetUpdateResponseDataPermissionsRecorderTypeLivestreamer, PresetUpdateResponseDataPermissionsRecorderTypeNone:
		return true
	}
	return false
}

// Waiting room type
type PresetUpdateResponseDataPermissionsWaitingRoomType string

const (
	PresetUpdateResponseDataPermissionsWaitingRoomTypeSkip                  PresetUpdateResponseDataPermissionsWaitingRoomType = "SKIP"
	PresetUpdateResponseDataPermissionsWaitingRoomTypeOnPrivilegedUserEntry PresetUpdateResponseDataPermissionsWaitingRoomType = "ON_PRIVILEGED_USER_ENTRY"
	PresetUpdateResponseDataPermissionsWaitingRoomTypeSkipOnAccept          PresetUpdateResponseDataPermissionsWaitingRoomType = "SKIP_ON_ACCEPT"
)

func (r PresetUpdateResponseDataPermissionsWaitingRoomType) IsKnown() bool {
	switch r {
	case PresetUpdateResponseDataPermissionsWaitingRoomTypeSkip, PresetUpdateResponseDataPermissionsWaitingRoomTypeOnPrivilegedUserEntry, PresetUpdateResponseDataPermissionsWaitingRoomTypeSkipOnAccept:
		return true
	}
	return false
}

type PresetDeleteResponse struct {
	// Data returned by the operation
	Data PresetDeleteResponseData `json:"data,required"`
	// Success status of the operation
	Success bool                     `json:"success,required"`
	JSON    presetDeleteResponseJSON `json:"-"`
}

// presetDeleteResponseJSON contains the JSON metadata for the struct
// [PresetDeleteResponse]
type presetDeleteResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type PresetDeleteResponseData struct {
	// ID of the preset
	ID     string                         `json:"id,required" format:"uuid"`
	Config PresetDeleteResponseDataConfig `json:"config,required"`
	// Name of the preset
	Name        string                              `json:"name,required"`
	UI          PresetDeleteResponseDataUI          `json:"ui,required"`
	Permissions PresetDeleteResponseDataPermissions `json:"permissions"`
	JSON        presetDeleteResponseDataJSON        `json:"-"`
}

// presetDeleteResponseDataJSON contains the JSON metadata for the struct
// [PresetDeleteResponseData]
type presetDeleteResponseDataJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	UI          apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataJSON) RawJSON() string {
	return r.raw
}

type PresetDeleteResponseDataConfig struct {
	// Maximum number of screen shares that can be active at a given time
	MaxScreenshareCount int64 `json:"max_screenshare_count,required"`
	// Maximum number of streams that are visible on a device
	MaxVideoStreams PresetDeleteResponseDataConfigMaxVideoStreams `json:"max_video_streams,required"`
	// Media configuration options. eg: Video quality
	Media PresetDeleteResponseDataConfigMedia `json:"media,required"`
	// Type of the meeting
	ViewType PresetDeleteResponseDataConfigViewType `json:"view_type,required"`
	JSON     presetDeleteResponseDataConfigJSON     `json:"-"`
}

// presetDeleteResponseDataConfigJSON contains the JSON metadata for the struct
// [PresetDeleteResponseDataConfig]
type presetDeleteResponseDataConfigJSON struct {
	MaxScreenshareCount apijson.Field
	MaxVideoStreams     apijson.Field
	Media               apijson.Field
	ViewType            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PresetDeleteResponseDataConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataConfigJSON) RawJSON() string {
	return r.raw
}

// Maximum number of streams that are visible on a device
type PresetDeleteResponseDataConfigMaxVideoStreams struct {
	// Maximum number of video streams visible on desktop devices
	Desktop int64 `json:"desktop,required"`
	// Maximum number of streams visible on mobile devices
	Mobile int64                                             `json:"mobile,required"`
	JSON   presetDeleteResponseDataConfigMaxVideoStreamsJSON `json:"-"`
}

// presetDeleteResponseDataConfigMaxVideoStreamsJSON contains the JSON metadata for
// the struct [PresetDeleteResponseDataConfigMaxVideoStreams]
type presetDeleteResponseDataConfigMaxVideoStreamsJSON struct {
	Desktop     apijson.Field
	Mobile      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataConfigMaxVideoStreams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataConfigMaxVideoStreamsJSON) RawJSON() string {
	return r.raw
}

// Media configuration options. eg: Video quality
type PresetDeleteResponseDataConfigMedia struct {
	// Configuration options for participant screen shares
	Screenshare PresetDeleteResponseDataConfigMediaScreenshare `json:"screenshare,required"`
	// Configuration options for participant videos
	Video PresetDeleteResponseDataConfigMediaVideo `json:"video,required"`
	// Control options for Audio quality.
	Audio PresetDeleteResponseDataConfigMediaAudio `json:"audio"`
	JSON  presetDeleteResponseDataConfigMediaJSON  `json:"-"`
}

// presetDeleteResponseDataConfigMediaJSON contains the JSON metadata for the
// struct [PresetDeleteResponseDataConfigMedia]
type presetDeleteResponseDataConfigMediaJSON struct {
	Screenshare apijson.Field
	Video       apijson.Field
	Audio       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataConfigMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataConfigMediaJSON) RawJSON() string {
	return r.raw
}

// Configuration options for participant screen shares
type PresetDeleteResponseDataConfigMediaScreenshare struct {
	// Frame rate of screen share
	FrameRate int64 `json:"frame_rate,required"`
	// Quality of screen share
	Quality PresetDeleteResponseDataConfigMediaScreenshareQuality `json:"quality,required"`
	JSON    presetDeleteResponseDataConfigMediaScreenshareJSON    `json:"-"`
}

// presetDeleteResponseDataConfigMediaScreenshareJSON contains the JSON metadata
// for the struct [PresetDeleteResponseDataConfigMediaScreenshare]
type presetDeleteResponseDataConfigMediaScreenshareJSON struct {
	FrameRate   apijson.Field
	Quality     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataConfigMediaScreenshare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataConfigMediaScreenshareJSON) RawJSON() string {
	return r.raw
}

// Quality of screen share
type PresetDeleteResponseDataConfigMediaScreenshareQuality string

const (
	PresetDeleteResponseDataConfigMediaScreenshareQualityHD   PresetDeleteResponseDataConfigMediaScreenshareQuality = "hd"
	PresetDeleteResponseDataConfigMediaScreenshareQualityVga  PresetDeleteResponseDataConfigMediaScreenshareQuality = "vga"
	PresetDeleteResponseDataConfigMediaScreenshareQualityQvga PresetDeleteResponseDataConfigMediaScreenshareQuality = "qvga"
)

func (r PresetDeleteResponseDataConfigMediaScreenshareQuality) IsKnown() bool {
	switch r {
	case PresetDeleteResponseDataConfigMediaScreenshareQualityHD, PresetDeleteResponseDataConfigMediaScreenshareQualityVga, PresetDeleteResponseDataConfigMediaScreenshareQualityQvga:
		return true
	}
	return false
}

// Configuration options for participant videos
type PresetDeleteResponseDataConfigMediaVideo struct {
	// Frame rate of participants' video
	FrameRate int64 `json:"frame_rate,required"`
	// Video quality of participants
	Quality PresetDeleteResponseDataConfigMediaVideoQuality `json:"quality,required"`
	JSON    presetDeleteResponseDataConfigMediaVideoJSON    `json:"-"`
}

// presetDeleteResponseDataConfigMediaVideoJSON contains the JSON metadata for the
// struct [PresetDeleteResponseDataConfigMediaVideo]
type presetDeleteResponseDataConfigMediaVideoJSON struct {
	FrameRate   apijson.Field
	Quality     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataConfigMediaVideo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataConfigMediaVideoJSON) RawJSON() string {
	return r.raw
}

// Video quality of participants
type PresetDeleteResponseDataConfigMediaVideoQuality string

const (
	PresetDeleteResponseDataConfigMediaVideoQualityHD   PresetDeleteResponseDataConfigMediaVideoQuality = "hd"
	PresetDeleteResponseDataConfigMediaVideoQualityVga  PresetDeleteResponseDataConfigMediaVideoQuality = "vga"
	PresetDeleteResponseDataConfigMediaVideoQualityQvga PresetDeleteResponseDataConfigMediaVideoQuality = "qvga"
)

func (r PresetDeleteResponseDataConfigMediaVideoQuality) IsKnown() bool {
	switch r {
	case PresetDeleteResponseDataConfigMediaVideoQualityHD, PresetDeleteResponseDataConfigMediaVideoQualityVga, PresetDeleteResponseDataConfigMediaVideoQualityQvga:
		return true
	}
	return false
}

// Control options for Audio quality.
type PresetDeleteResponseDataConfigMediaAudio struct {
	// Enable High Quality Audio for your meetings
	EnableHighBitrate bool `json:"enable_high_bitrate"`
	// Enable Stereo for your meetings
	EnableStereo bool                                         `json:"enable_stereo"`
	JSON         presetDeleteResponseDataConfigMediaAudioJSON `json:"-"`
}

// presetDeleteResponseDataConfigMediaAudioJSON contains the JSON metadata for the
// struct [PresetDeleteResponseDataConfigMediaAudio]
type presetDeleteResponseDataConfigMediaAudioJSON struct {
	EnableHighBitrate apijson.Field
	EnableStereo      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PresetDeleteResponseDataConfigMediaAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataConfigMediaAudioJSON) RawJSON() string {
	return r.raw
}

// Type of the meeting
type PresetDeleteResponseDataConfigViewType string

const (
	PresetDeleteResponseDataConfigViewTypeGroupCall PresetDeleteResponseDataConfigViewType = "GROUP_CALL"
	PresetDeleteResponseDataConfigViewTypeWebinar   PresetDeleteResponseDataConfigViewType = "WEBINAR"
	PresetDeleteResponseDataConfigViewTypeAudioRoom PresetDeleteResponseDataConfigViewType = "AUDIO_ROOM"
)

func (r PresetDeleteResponseDataConfigViewType) IsKnown() bool {
	switch r {
	case PresetDeleteResponseDataConfigViewTypeGroupCall, PresetDeleteResponseDataConfigViewTypeWebinar, PresetDeleteResponseDataConfigViewTypeAudioRoom:
		return true
	}
	return false
}

type PresetDeleteResponseDataUI struct {
	DesignTokens PresetDeleteResponseDataUIDesignTokens `json:"design_tokens,required"`
	ConfigDiff   interface{}                            `json:"config_diff"`
	JSON         presetDeleteResponseDataUIJSON         `json:"-"`
}

// presetDeleteResponseDataUIJSON contains the JSON metadata for the struct
// [PresetDeleteResponseDataUI]
type presetDeleteResponseDataUIJSON struct {
	DesignTokens apijson.Field
	ConfigDiff   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PresetDeleteResponseDataUI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataUIJSON) RawJSON() string {
	return r.raw
}

type PresetDeleteResponseDataUIDesignTokens struct {
	BorderRadius PresetDeleteResponseDataUIDesignTokensBorderRadius `json:"border_radius,required"`
	BorderWidth  PresetDeleteResponseDataUIDesignTokensBorderWidth  `json:"border_width,required"`
	Colors       PresetDeleteResponseDataUIDesignTokensColors       `json:"colors,required"`
	Logo         string                                             `json:"logo,required"`
	SpacingBase  float64                                            `json:"spacing_base,required"`
	Theme        PresetDeleteResponseDataUIDesignTokensTheme        `json:"theme,required"`
	JSON         presetDeleteResponseDataUIDesignTokensJSON         `json:"-"`
}

// presetDeleteResponseDataUIDesignTokensJSON contains the JSON metadata for the
// struct [PresetDeleteResponseDataUIDesignTokens]
type presetDeleteResponseDataUIDesignTokensJSON struct {
	BorderRadius apijson.Field
	BorderWidth  apijson.Field
	Colors       apijson.Field
	Logo         apijson.Field
	SpacingBase  apijson.Field
	Theme        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PresetDeleteResponseDataUIDesignTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataUIDesignTokensJSON) RawJSON() string {
	return r.raw
}

type PresetDeleteResponseDataUIDesignTokensBorderRadius string

const (
	PresetDeleteResponseDataUIDesignTokensBorderRadiusRounded PresetDeleteResponseDataUIDesignTokensBorderRadius = "rounded"
)

func (r PresetDeleteResponseDataUIDesignTokensBorderRadius) IsKnown() bool {
	switch r {
	case PresetDeleteResponseDataUIDesignTokensBorderRadiusRounded:
		return true
	}
	return false
}

type PresetDeleteResponseDataUIDesignTokensBorderWidth string

const (
	PresetDeleteResponseDataUIDesignTokensBorderWidthThin PresetDeleteResponseDataUIDesignTokensBorderWidth = "thin"
)

func (r PresetDeleteResponseDataUIDesignTokensBorderWidth) IsKnown() bool {
	switch r {
	case PresetDeleteResponseDataUIDesignTokensBorderWidthThin:
		return true
	}
	return false
}

type PresetDeleteResponseDataUIDesignTokensColors struct {
	Background  PresetDeleteResponseDataUIDesignTokensColorsBackground `json:"background,required"`
	Brand       PresetDeleteResponseDataUIDesignTokensColorsBrand      `json:"brand,required"`
	Danger      string                                                 `json:"danger,required"`
	Success     string                                                 `json:"success,required"`
	Text        string                                                 `json:"text,required"`
	TextOnBrand string                                                 `json:"text_on_brand,required"`
	VideoBg     string                                                 `json:"video_bg,required"`
	Warning     string                                                 `json:"warning,required"`
	JSON        presetDeleteResponseDataUIDesignTokensColorsJSON       `json:"-"`
}

// presetDeleteResponseDataUIDesignTokensColorsJSON contains the JSON metadata for
// the struct [PresetDeleteResponseDataUIDesignTokensColors]
type presetDeleteResponseDataUIDesignTokensColorsJSON struct {
	Background  apijson.Field
	Brand       apijson.Field
	Danger      apijson.Field
	Success     apijson.Field
	Text        apijson.Field
	TextOnBrand apijson.Field
	VideoBg     apijson.Field
	Warning     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataUIDesignTokensColors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataUIDesignTokensColorsJSON) RawJSON() string {
	return r.raw
}

type PresetDeleteResponseDataUIDesignTokensColorsBackground struct {
	Number1000 string                                                     `json:"1000,required"`
	Number600  string                                                     `json:"600,required"`
	Number700  string                                                     `json:"700,required"`
	Number800  string                                                     `json:"800,required"`
	Number900  string                                                     `json:"900,required"`
	JSON       presetDeleteResponseDataUIDesignTokensColorsBackgroundJSON `json:"-"`
}

// presetDeleteResponseDataUIDesignTokensColorsBackgroundJSON contains the JSON
// metadata for the struct [PresetDeleteResponseDataUIDesignTokensColorsBackground]
type presetDeleteResponseDataUIDesignTokensColorsBackgroundJSON struct {
	Number1000  apijson.Field
	Number600   apijson.Field
	Number700   apijson.Field
	Number800   apijson.Field
	Number900   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataUIDesignTokensColorsBackground) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataUIDesignTokensColorsBackgroundJSON) RawJSON() string {
	return r.raw
}

type PresetDeleteResponseDataUIDesignTokensColorsBrand struct {
	Number300 string                                                `json:"300,required"`
	Number400 string                                                `json:"400,required"`
	Number500 string                                                `json:"500,required"`
	Number600 string                                                `json:"600,required"`
	Number700 string                                                `json:"700,required"`
	JSON      presetDeleteResponseDataUIDesignTokensColorsBrandJSON `json:"-"`
}

// presetDeleteResponseDataUIDesignTokensColorsBrandJSON contains the JSON metadata
// for the struct [PresetDeleteResponseDataUIDesignTokensColorsBrand]
type presetDeleteResponseDataUIDesignTokensColorsBrandJSON struct {
	Number300   apijson.Field
	Number400   apijson.Field
	Number500   apijson.Field
	Number600   apijson.Field
	Number700   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataUIDesignTokensColorsBrand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataUIDesignTokensColorsBrandJSON) RawJSON() string {
	return r.raw
}

type PresetDeleteResponseDataUIDesignTokensTheme string

const (
	PresetDeleteResponseDataUIDesignTokensThemeDark PresetDeleteResponseDataUIDesignTokensTheme = "dark"
)

func (r PresetDeleteResponseDataUIDesignTokensTheme) IsKnown() bool {
	switch r {
	case PresetDeleteResponseDataUIDesignTokensThemeDark:
		return true
	}
	return false
}

type PresetDeleteResponseDataPermissions struct {
	// Whether this participant can accept waiting requests
	AcceptWaitingRequests           bool `json:"accept_waiting_requests,required"`
	CanAcceptProductionRequests     bool `json:"can_accept_production_requests,required"`
	CanChangeParticipantPermissions bool `json:"can_change_participant_permissions,required"`
	CanEditDisplayName              bool `json:"can_edit_display_name,required"`
	CanLivestream                   bool `json:"can_livestream,required"`
	CanRecord                       bool `json:"can_record,required"`
	CanSpotlight                    bool `json:"can_spotlight,required"`
	// Chat permissions
	Chat                            PresetDeleteResponseDataPermissionsChat              `json:"chat,required"`
	ConnectedMeetings               PresetDeleteResponseDataPermissionsConnectedMeetings `json:"connected_meetings,required"`
	DisableParticipantAudio         bool                                                 `json:"disable_participant_audio,required"`
	DisableParticipantScreensharing bool                                                 `json:"disable_participant_screensharing,required"`
	DisableParticipantVideo         bool                                                 `json:"disable_participant_video,required"`
	// Whether this participant is visible to others or not
	HiddenParticipant bool `json:"hidden_participant,required"`
	KickParticipant   bool `json:"kick_participant,required"`
	// Media permissions
	Media          PresetDeleteResponseDataPermissionsMedia `json:"media,required"`
	PinParticipant bool                                     `json:"pin_participant,required"`
	// Plugin permissions
	Plugins PresetDeleteResponseDataPermissionsPlugins `json:"plugins,required"`
	// Poll permissions
	Polls PresetDeleteResponseDataPermissionsPolls `json:"polls,required"`
	// Type of the recording peer
	RecorderType        PresetDeleteResponseDataPermissionsRecorderType `json:"recorder_type,required"`
	ShowParticipantList bool                                            `json:"show_participant_list,required"`
	// Waiting room type
	WaitingRoomType PresetDeleteResponseDataPermissionsWaitingRoomType `json:"waiting_room_type,required"`
	IsRecorder      bool                                               `json:"is_recorder"`
	JSON            presetDeleteResponseDataPermissionsJSON            `json:"-"`
}

// presetDeleteResponseDataPermissionsJSON contains the JSON metadata for the
// struct [PresetDeleteResponseDataPermissions]
type presetDeleteResponseDataPermissionsJSON struct {
	AcceptWaitingRequests           apijson.Field
	CanAcceptProductionRequests     apijson.Field
	CanChangeParticipantPermissions apijson.Field
	CanEditDisplayName              apijson.Field
	CanLivestream                   apijson.Field
	CanRecord                       apijson.Field
	CanSpotlight                    apijson.Field
	Chat                            apijson.Field
	ConnectedMeetings               apijson.Field
	DisableParticipantAudio         apijson.Field
	DisableParticipantScreensharing apijson.Field
	DisableParticipantVideo         apijson.Field
	HiddenParticipant               apijson.Field
	KickParticipant                 apijson.Field
	Media                           apijson.Field
	PinParticipant                  apijson.Field
	Plugins                         apijson.Field
	Polls                           apijson.Field
	RecorderType                    apijson.Field
	ShowParticipantList             apijson.Field
	WaitingRoomType                 apijson.Field
	IsRecorder                      apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *PresetDeleteResponseDataPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataPermissionsJSON) RawJSON() string {
	return r.raw
}

// Chat permissions
type PresetDeleteResponseDataPermissionsChat struct {
	Private PresetDeleteResponseDataPermissionsChatPrivate `json:"private,required"`
	Public  PresetDeleteResponseDataPermissionsChatPublic  `json:"public,required"`
	JSON    presetDeleteResponseDataPermissionsChatJSON    `json:"-"`
}

// presetDeleteResponseDataPermissionsChatJSON contains the JSON metadata for the
// struct [PresetDeleteResponseDataPermissionsChat]
type presetDeleteResponseDataPermissionsChatJSON struct {
	Private     apijson.Field
	Public      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataPermissionsChat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataPermissionsChatJSON) RawJSON() string {
	return r.raw
}

type PresetDeleteResponseDataPermissionsChatPrivate struct {
	CanReceive bool                                               `json:"can_receive,required"`
	CanSend    bool                                               `json:"can_send,required"`
	Files      bool                                               `json:"files,required"`
	Text       bool                                               `json:"text,required"`
	JSON       presetDeleteResponseDataPermissionsChatPrivateJSON `json:"-"`
}

// presetDeleteResponseDataPermissionsChatPrivateJSON contains the JSON metadata
// for the struct [PresetDeleteResponseDataPermissionsChatPrivate]
type presetDeleteResponseDataPermissionsChatPrivateJSON struct {
	CanReceive  apijson.Field
	CanSend     apijson.Field
	Files       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataPermissionsChatPrivate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataPermissionsChatPrivateJSON) RawJSON() string {
	return r.raw
}

type PresetDeleteResponseDataPermissionsChatPublic struct {
	// Can send messages in general
	CanSend bool `json:"can_send,required"`
	// Can send file messages
	Files bool `json:"files,required"`
	// Can send text messages
	Text bool                                              `json:"text,required"`
	JSON presetDeleteResponseDataPermissionsChatPublicJSON `json:"-"`
}

// presetDeleteResponseDataPermissionsChatPublicJSON contains the JSON metadata for
// the struct [PresetDeleteResponseDataPermissionsChatPublic]
type presetDeleteResponseDataPermissionsChatPublicJSON struct {
	CanSend     apijson.Field
	Files       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataPermissionsChatPublic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataPermissionsChatPublicJSON) RawJSON() string {
	return r.raw
}

type PresetDeleteResponseDataPermissionsConnectedMeetings struct {
	CanAlterConnectedMeetings  bool                                                     `json:"can_alter_connected_meetings,required"`
	CanSwitchConnectedMeetings bool                                                     `json:"can_switch_connected_meetings,required"`
	CanSwitchToParentMeeting   bool                                                     `json:"can_switch_to_parent_meeting,required"`
	JSON                       presetDeleteResponseDataPermissionsConnectedMeetingsJSON `json:"-"`
}

// presetDeleteResponseDataPermissionsConnectedMeetingsJSON contains the JSON
// metadata for the struct [PresetDeleteResponseDataPermissionsConnectedMeetings]
type presetDeleteResponseDataPermissionsConnectedMeetingsJSON struct {
	CanAlterConnectedMeetings  apijson.Field
	CanSwitchConnectedMeetings apijson.Field
	CanSwitchToParentMeeting   apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *PresetDeleteResponseDataPermissionsConnectedMeetings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataPermissionsConnectedMeetingsJSON) RawJSON() string {
	return r.raw
}

// Media permissions
type PresetDeleteResponseDataPermissionsMedia struct {
	// Audio permissions
	Audio PresetDeleteResponseDataPermissionsMediaAudio `json:"audio,required"`
	// Screenshare permissions
	Screenshare PresetDeleteResponseDataPermissionsMediaScreenshare `json:"screenshare,required"`
	// Video permissions
	Video PresetDeleteResponseDataPermissionsMediaVideo `json:"video,required"`
	JSON  presetDeleteResponseDataPermissionsMediaJSON  `json:"-"`
}

// presetDeleteResponseDataPermissionsMediaJSON contains the JSON metadata for the
// struct [PresetDeleteResponseDataPermissionsMedia]
type presetDeleteResponseDataPermissionsMediaJSON struct {
	Audio       apijson.Field
	Screenshare apijson.Field
	Video       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataPermissionsMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataPermissionsMediaJSON) RawJSON() string {
	return r.raw
}

// Audio permissions
type PresetDeleteResponseDataPermissionsMediaAudio struct {
	// Can produce audio
	CanProduce PresetDeleteResponseDataPermissionsMediaAudioCanProduce `json:"can_produce,required"`
	JSON       presetDeleteResponseDataPermissionsMediaAudioJSON       `json:"-"`
}

// presetDeleteResponseDataPermissionsMediaAudioJSON contains the JSON metadata for
// the struct [PresetDeleteResponseDataPermissionsMediaAudio]
type presetDeleteResponseDataPermissionsMediaAudioJSON struct {
	CanProduce  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataPermissionsMediaAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataPermissionsMediaAudioJSON) RawJSON() string {
	return r.raw
}

// Can produce audio
type PresetDeleteResponseDataPermissionsMediaAudioCanProduce string

const (
	PresetDeleteResponseDataPermissionsMediaAudioCanProduceAllowed    PresetDeleteResponseDataPermissionsMediaAudioCanProduce = "ALLOWED"
	PresetDeleteResponseDataPermissionsMediaAudioCanProduceNotAllowed PresetDeleteResponseDataPermissionsMediaAudioCanProduce = "NOT_ALLOWED"
	PresetDeleteResponseDataPermissionsMediaAudioCanProduceCanRequest PresetDeleteResponseDataPermissionsMediaAudioCanProduce = "CAN_REQUEST"
)

func (r PresetDeleteResponseDataPermissionsMediaAudioCanProduce) IsKnown() bool {
	switch r {
	case PresetDeleteResponseDataPermissionsMediaAudioCanProduceAllowed, PresetDeleteResponseDataPermissionsMediaAudioCanProduceNotAllowed, PresetDeleteResponseDataPermissionsMediaAudioCanProduceCanRequest:
		return true
	}
	return false
}

// Screenshare permissions
type PresetDeleteResponseDataPermissionsMediaScreenshare struct {
	// Can produce screen share video
	CanProduce PresetDeleteResponseDataPermissionsMediaScreenshareCanProduce `json:"can_produce,required"`
	JSON       presetDeleteResponseDataPermissionsMediaScreenshareJSON       `json:"-"`
}

// presetDeleteResponseDataPermissionsMediaScreenshareJSON contains the JSON
// metadata for the struct [PresetDeleteResponseDataPermissionsMediaScreenshare]
type presetDeleteResponseDataPermissionsMediaScreenshareJSON struct {
	CanProduce  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataPermissionsMediaScreenshare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataPermissionsMediaScreenshareJSON) RawJSON() string {
	return r.raw
}

// Can produce screen share video
type PresetDeleteResponseDataPermissionsMediaScreenshareCanProduce string

const (
	PresetDeleteResponseDataPermissionsMediaScreenshareCanProduceAllowed    PresetDeleteResponseDataPermissionsMediaScreenshareCanProduce = "ALLOWED"
	PresetDeleteResponseDataPermissionsMediaScreenshareCanProduceNotAllowed PresetDeleteResponseDataPermissionsMediaScreenshareCanProduce = "NOT_ALLOWED"
	PresetDeleteResponseDataPermissionsMediaScreenshareCanProduceCanRequest PresetDeleteResponseDataPermissionsMediaScreenshareCanProduce = "CAN_REQUEST"
)

func (r PresetDeleteResponseDataPermissionsMediaScreenshareCanProduce) IsKnown() bool {
	switch r {
	case PresetDeleteResponseDataPermissionsMediaScreenshareCanProduceAllowed, PresetDeleteResponseDataPermissionsMediaScreenshareCanProduceNotAllowed, PresetDeleteResponseDataPermissionsMediaScreenshareCanProduceCanRequest:
		return true
	}
	return false
}

// Video permissions
type PresetDeleteResponseDataPermissionsMediaVideo struct {
	// Can produce video
	CanProduce PresetDeleteResponseDataPermissionsMediaVideoCanProduce `json:"can_produce,required"`
	JSON       presetDeleteResponseDataPermissionsMediaVideoJSON       `json:"-"`
}

// presetDeleteResponseDataPermissionsMediaVideoJSON contains the JSON metadata for
// the struct [PresetDeleteResponseDataPermissionsMediaVideo]
type presetDeleteResponseDataPermissionsMediaVideoJSON struct {
	CanProduce  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataPermissionsMediaVideo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataPermissionsMediaVideoJSON) RawJSON() string {
	return r.raw
}

// Can produce video
type PresetDeleteResponseDataPermissionsMediaVideoCanProduce string

const (
	PresetDeleteResponseDataPermissionsMediaVideoCanProduceAllowed    PresetDeleteResponseDataPermissionsMediaVideoCanProduce = "ALLOWED"
	PresetDeleteResponseDataPermissionsMediaVideoCanProduceNotAllowed PresetDeleteResponseDataPermissionsMediaVideoCanProduce = "NOT_ALLOWED"
	PresetDeleteResponseDataPermissionsMediaVideoCanProduceCanRequest PresetDeleteResponseDataPermissionsMediaVideoCanProduce = "CAN_REQUEST"
)

func (r PresetDeleteResponseDataPermissionsMediaVideoCanProduce) IsKnown() bool {
	switch r {
	case PresetDeleteResponseDataPermissionsMediaVideoCanProduceAllowed, PresetDeleteResponseDataPermissionsMediaVideoCanProduceNotAllowed, PresetDeleteResponseDataPermissionsMediaVideoCanProduceCanRequest:
		return true
	}
	return false
}

// Plugin permissions
type PresetDeleteResponseDataPermissionsPlugins struct {
	// Can close plugins that are already open
	CanClose bool `json:"can_close,required"`
	// Can edit plugin config
	CanEditConfig bool `json:"can_edit_config,required"`
	// Can start plugins
	CanStart bool                                                  `json:"can_start,required"`
	Config   PresetDeleteResponseDataPermissionsPluginsConfigUnion `json:"config,required" format:"uuid"`
	JSON     presetDeleteResponseDataPermissionsPluginsJSON        `json:"-"`
}

// presetDeleteResponseDataPermissionsPluginsJSON contains the JSON metadata for
// the struct [PresetDeleteResponseDataPermissionsPlugins]
type presetDeleteResponseDataPermissionsPluginsJSON struct {
	CanClose      apijson.Field
	CanEditConfig apijson.Field
	CanStart      apijson.Field
	Config        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PresetDeleteResponseDataPermissionsPlugins) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataPermissionsPluginsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or
// [PresetDeleteResponseDataPermissionsPluginsConfigObject].
type PresetDeleteResponseDataPermissionsPluginsConfigUnion interface {
	ImplementsPresetDeleteResponseDataPermissionsPluginsConfigUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PresetDeleteResponseDataPermissionsPluginsConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PresetDeleteResponseDataPermissionsPluginsConfigObject{}),
		},
	)
}

type PresetDeleteResponseDataPermissionsPluginsConfigObject struct {
	AccessControl   PresetDeleteResponseDataPermissionsPluginsConfigObjectAccessControl `json:"access_control,required"`
	HandlesViewOnly bool                                                                `json:"handles_view_only,required"`
	JSON            presetDeleteResponseDataPermissionsPluginsConfigObjectJSON          `json:"-"`
}

// presetDeleteResponseDataPermissionsPluginsConfigObjectJSON contains the JSON
// metadata for the struct [PresetDeleteResponseDataPermissionsPluginsConfigObject]
type presetDeleteResponseDataPermissionsPluginsConfigObjectJSON struct {
	AccessControl   apijson.Field
	HandlesViewOnly apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PresetDeleteResponseDataPermissionsPluginsConfigObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataPermissionsPluginsConfigObjectJSON) RawJSON() string {
	return r.raw
}

func (r PresetDeleteResponseDataPermissionsPluginsConfigObject) ImplementsPresetDeleteResponseDataPermissionsPluginsConfigUnion() {
}

type PresetDeleteResponseDataPermissionsPluginsConfigObjectAccessControl string

const (
	PresetDeleteResponseDataPermissionsPluginsConfigObjectAccessControlFullAccess PresetDeleteResponseDataPermissionsPluginsConfigObjectAccessControl = "FULL_ACCESS"
	PresetDeleteResponseDataPermissionsPluginsConfigObjectAccessControlViewOnly   PresetDeleteResponseDataPermissionsPluginsConfigObjectAccessControl = "VIEW_ONLY"
)

func (r PresetDeleteResponseDataPermissionsPluginsConfigObjectAccessControl) IsKnown() bool {
	switch r {
	case PresetDeleteResponseDataPermissionsPluginsConfigObjectAccessControlFullAccess, PresetDeleteResponseDataPermissionsPluginsConfigObjectAccessControlViewOnly:
		return true
	}
	return false
}

// Poll permissions
type PresetDeleteResponseDataPermissionsPolls struct {
	// Can create polls
	CanCreate bool `json:"can_create,required"`
	// Can view polls
	CanView bool `json:"can_view,required"`
	// Can vote on polls
	CanVote bool                                         `json:"can_vote,required"`
	JSON    presetDeleteResponseDataPermissionsPollsJSON `json:"-"`
}

// presetDeleteResponseDataPermissionsPollsJSON contains the JSON metadata for the
// struct [PresetDeleteResponseDataPermissionsPolls]
type presetDeleteResponseDataPermissionsPollsJSON struct {
	CanCreate   apijson.Field
	CanView     apijson.Field
	CanVote     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetDeleteResponseDataPermissionsPolls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetDeleteResponseDataPermissionsPollsJSON) RawJSON() string {
	return r.raw
}

// Type of the recording peer
type PresetDeleteResponseDataPermissionsRecorderType string

const (
	PresetDeleteResponseDataPermissionsRecorderTypeRecorder     PresetDeleteResponseDataPermissionsRecorderType = "RECORDER"
	PresetDeleteResponseDataPermissionsRecorderTypeLivestreamer PresetDeleteResponseDataPermissionsRecorderType = "LIVESTREAMER"
	PresetDeleteResponseDataPermissionsRecorderTypeNone         PresetDeleteResponseDataPermissionsRecorderType = "NONE"
)

func (r PresetDeleteResponseDataPermissionsRecorderType) IsKnown() bool {
	switch r {
	case PresetDeleteResponseDataPermissionsRecorderTypeRecorder, PresetDeleteResponseDataPermissionsRecorderTypeLivestreamer, PresetDeleteResponseDataPermissionsRecorderTypeNone:
		return true
	}
	return false
}

// Waiting room type
type PresetDeleteResponseDataPermissionsWaitingRoomType string

const (
	PresetDeleteResponseDataPermissionsWaitingRoomTypeSkip                  PresetDeleteResponseDataPermissionsWaitingRoomType = "SKIP"
	PresetDeleteResponseDataPermissionsWaitingRoomTypeOnPrivilegedUserEntry PresetDeleteResponseDataPermissionsWaitingRoomType = "ON_PRIVILEGED_USER_ENTRY"
	PresetDeleteResponseDataPermissionsWaitingRoomTypeSkipOnAccept          PresetDeleteResponseDataPermissionsWaitingRoomType = "SKIP_ON_ACCEPT"
)

func (r PresetDeleteResponseDataPermissionsWaitingRoomType) IsKnown() bool {
	switch r {
	case PresetDeleteResponseDataPermissionsWaitingRoomTypeSkip, PresetDeleteResponseDataPermissionsWaitingRoomTypeOnPrivilegedUserEntry, PresetDeleteResponseDataPermissionsWaitingRoomTypeSkipOnAccept:
		return true
	}
	return false
}

type PresetGetResponse struct {
	Data    []PresetGetResponseData `json:"data,required"`
	Paging  PresetGetResponsePaging `json:"paging,required"`
	Success bool                    `json:"success,required"`
	JSON    presetGetResponseJSON   `json:"-"`
}

// presetGetResponseJSON contains the JSON metadata for the struct
// [PresetGetResponse]
type presetGetResponseJSON struct {
	Data        apijson.Field
	Paging      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetResponseJSON) RawJSON() string {
	return r.raw
}

// Returned by Get All Presets route
type PresetGetResponseData struct {
	// ID of the preset
	ID string `json:"id" format:"uuid"`
	// Timestamp this preset was created at
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Name of the preset
	Name string `json:"name"`
	// Timestamp this preset was last updated
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      presetGetResponseDataJSON `json:"-"`
}

// presetGetResponseDataJSON contains the JSON metadata for the struct
// [PresetGetResponseData]
type presetGetResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type PresetGetResponsePaging struct {
	EndOffset   float64                     `json:"end_offset,required"`
	StartOffset float64                     `json:"start_offset,required"`
	TotalCount  float64                     `json:"total_count,required"`
	JSON        presetGetResponsePagingJSON `json:"-"`
}

// presetGetResponsePagingJSON contains the JSON metadata for the struct
// [PresetGetResponsePaging]
type presetGetResponsePagingJSON struct {
	EndOffset   apijson.Field
	StartOffset apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetResponsePaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetResponsePagingJSON) RawJSON() string {
	return r.raw
}

type PresetGetPresetByIDResponse struct {
	// Data returned by the operation
	Data PresetGetPresetByIDResponseData `json:"data,required"`
	// Success status of the operation
	Success bool                            `json:"success,required"`
	JSON    presetGetPresetByIDResponseJSON `json:"-"`
}

// presetGetPresetByIDResponseJSON contains the JSON metadata for the struct
// [PresetGetPresetByIDResponse]
type presetGetPresetByIDResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type PresetGetPresetByIDResponseData struct {
	// ID of the preset
	ID     string                                `json:"id,required" format:"uuid"`
	Config PresetGetPresetByIDResponseDataConfig `json:"config,required"`
	// Name of the preset
	Name        string                                     `json:"name,required"`
	UI          PresetGetPresetByIDResponseDataUI          `json:"ui,required"`
	Permissions PresetGetPresetByIDResponseDataPermissions `json:"permissions"`
	JSON        presetGetPresetByIDResponseDataJSON        `json:"-"`
}

// presetGetPresetByIDResponseDataJSON contains the JSON metadata for the struct
// [PresetGetPresetByIDResponseData]
type presetGetPresetByIDResponseDataJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	UI          apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataJSON) RawJSON() string {
	return r.raw
}

type PresetGetPresetByIDResponseDataConfig struct {
	// Maximum number of screen shares that can be active at a given time
	MaxScreenshareCount int64 `json:"max_screenshare_count,required"`
	// Maximum number of streams that are visible on a device
	MaxVideoStreams PresetGetPresetByIDResponseDataConfigMaxVideoStreams `json:"max_video_streams,required"`
	// Media configuration options. eg: Video quality
	Media PresetGetPresetByIDResponseDataConfigMedia `json:"media,required"`
	// Type of the meeting
	ViewType PresetGetPresetByIDResponseDataConfigViewType `json:"view_type,required"`
	JSON     presetGetPresetByIDResponseDataConfigJSON     `json:"-"`
}

// presetGetPresetByIDResponseDataConfigJSON contains the JSON metadata for the
// struct [PresetGetPresetByIDResponseDataConfig]
type presetGetPresetByIDResponseDataConfigJSON struct {
	MaxScreenshareCount apijson.Field
	MaxVideoStreams     apijson.Field
	Media               apijson.Field
	ViewType            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataConfigJSON) RawJSON() string {
	return r.raw
}

// Maximum number of streams that are visible on a device
type PresetGetPresetByIDResponseDataConfigMaxVideoStreams struct {
	// Maximum number of video streams visible on desktop devices
	Desktop int64 `json:"desktop,required"`
	// Maximum number of streams visible on mobile devices
	Mobile int64                                                    `json:"mobile,required"`
	JSON   presetGetPresetByIDResponseDataConfigMaxVideoStreamsJSON `json:"-"`
}

// presetGetPresetByIDResponseDataConfigMaxVideoStreamsJSON contains the JSON
// metadata for the struct [PresetGetPresetByIDResponseDataConfigMaxVideoStreams]
type presetGetPresetByIDResponseDataConfigMaxVideoStreamsJSON struct {
	Desktop     apijson.Field
	Mobile      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataConfigMaxVideoStreams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataConfigMaxVideoStreamsJSON) RawJSON() string {
	return r.raw
}

// Media configuration options. eg: Video quality
type PresetGetPresetByIDResponseDataConfigMedia struct {
	// Configuration options for participant screen shares
	Screenshare PresetGetPresetByIDResponseDataConfigMediaScreenshare `json:"screenshare,required"`
	// Configuration options for participant videos
	Video PresetGetPresetByIDResponseDataConfigMediaVideo `json:"video,required"`
	// Control options for Audio quality.
	Audio PresetGetPresetByIDResponseDataConfigMediaAudio `json:"audio"`
	JSON  presetGetPresetByIDResponseDataConfigMediaJSON  `json:"-"`
}

// presetGetPresetByIDResponseDataConfigMediaJSON contains the JSON metadata for
// the struct [PresetGetPresetByIDResponseDataConfigMedia]
type presetGetPresetByIDResponseDataConfigMediaJSON struct {
	Screenshare apijson.Field
	Video       apijson.Field
	Audio       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataConfigMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataConfigMediaJSON) RawJSON() string {
	return r.raw
}

// Configuration options for participant screen shares
type PresetGetPresetByIDResponseDataConfigMediaScreenshare struct {
	// Frame rate of screen share
	FrameRate int64 `json:"frame_rate,required"`
	// Quality of screen share
	Quality PresetGetPresetByIDResponseDataConfigMediaScreenshareQuality `json:"quality,required"`
	JSON    presetGetPresetByIDResponseDataConfigMediaScreenshareJSON    `json:"-"`
}

// presetGetPresetByIDResponseDataConfigMediaScreenshareJSON contains the JSON
// metadata for the struct [PresetGetPresetByIDResponseDataConfigMediaScreenshare]
type presetGetPresetByIDResponseDataConfigMediaScreenshareJSON struct {
	FrameRate   apijson.Field
	Quality     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataConfigMediaScreenshare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataConfigMediaScreenshareJSON) RawJSON() string {
	return r.raw
}

// Quality of screen share
type PresetGetPresetByIDResponseDataConfigMediaScreenshareQuality string

const (
	PresetGetPresetByIDResponseDataConfigMediaScreenshareQualityHD   PresetGetPresetByIDResponseDataConfigMediaScreenshareQuality = "hd"
	PresetGetPresetByIDResponseDataConfigMediaScreenshareQualityVga  PresetGetPresetByIDResponseDataConfigMediaScreenshareQuality = "vga"
	PresetGetPresetByIDResponseDataConfigMediaScreenshareQualityQvga PresetGetPresetByIDResponseDataConfigMediaScreenshareQuality = "qvga"
)

func (r PresetGetPresetByIDResponseDataConfigMediaScreenshareQuality) IsKnown() bool {
	switch r {
	case PresetGetPresetByIDResponseDataConfigMediaScreenshareQualityHD, PresetGetPresetByIDResponseDataConfigMediaScreenshareQualityVga, PresetGetPresetByIDResponseDataConfigMediaScreenshareQualityQvga:
		return true
	}
	return false
}

// Configuration options for participant videos
type PresetGetPresetByIDResponseDataConfigMediaVideo struct {
	// Frame rate of participants' video
	FrameRate int64 `json:"frame_rate,required"`
	// Video quality of participants
	Quality PresetGetPresetByIDResponseDataConfigMediaVideoQuality `json:"quality,required"`
	JSON    presetGetPresetByIDResponseDataConfigMediaVideoJSON    `json:"-"`
}

// presetGetPresetByIDResponseDataConfigMediaVideoJSON contains the JSON metadata
// for the struct [PresetGetPresetByIDResponseDataConfigMediaVideo]
type presetGetPresetByIDResponseDataConfigMediaVideoJSON struct {
	FrameRate   apijson.Field
	Quality     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataConfigMediaVideo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataConfigMediaVideoJSON) RawJSON() string {
	return r.raw
}

// Video quality of participants
type PresetGetPresetByIDResponseDataConfigMediaVideoQuality string

const (
	PresetGetPresetByIDResponseDataConfigMediaVideoQualityHD   PresetGetPresetByIDResponseDataConfigMediaVideoQuality = "hd"
	PresetGetPresetByIDResponseDataConfigMediaVideoQualityVga  PresetGetPresetByIDResponseDataConfigMediaVideoQuality = "vga"
	PresetGetPresetByIDResponseDataConfigMediaVideoQualityQvga PresetGetPresetByIDResponseDataConfigMediaVideoQuality = "qvga"
)

func (r PresetGetPresetByIDResponseDataConfigMediaVideoQuality) IsKnown() bool {
	switch r {
	case PresetGetPresetByIDResponseDataConfigMediaVideoQualityHD, PresetGetPresetByIDResponseDataConfigMediaVideoQualityVga, PresetGetPresetByIDResponseDataConfigMediaVideoQualityQvga:
		return true
	}
	return false
}

// Control options for Audio quality.
type PresetGetPresetByIDResponseDataConfigMediaAudio struct {
	// Enable High Quality Audio for your meetings
	EnableHighBitrate bool `json:"enable_high_bitrate"`
	// Enable Stereo for your meetings
	EnableStereo bool                                                `json:"enable_stereo"`
	JSON         presetGetPresetByIDResponseDataConfigMediaAudioJSON `json:"-"`
}

// presetGetPresetByIDResponseDataConfigMediaAudioJSON contains the JSON metadata
// for the struct [PresetGetPresetByIDResponseDataConfigMediaAudio]
type presetGetPresetByIDResponseDataConfigMediaAudioJSON struct {
	EnableHighBitrate apijson.Field
	EnableStereo      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataConfigMediaAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataConfigMediaAudioJSON) RawJSON() string {
	return r.raw
}

// Type of the meeting
type PresetGetPresetByIDResponseDataConfigViewType string

const (
	PresetGetPresetByIDResponseDataConfigViewTypeGroupCall PresetGetPresetByIDResponseDataConfigViewType = "GROUP_CALL"
	PresetGetPresetByIDResponseDataConfigViewTypeWebinar   PresetGetPresetByIDResponseDataConfigViewType = "WEBINAR"
	PresetGetPresetByIDResponseDataConfigViewTypeAudioRoom PresetGetPresetByIDResponseDataConfigViewType = "AUDIO_ROOM"
)

func (r PresetGetPresetByIDResponseDataConfigViewType) IsKnown() bool {
	switch r {
	case PresetGetPresetByIDResponseDataConfigViewTypeGroupCall, PresetGetPresetByIDResponseDataConfigViewTypeWebinar, PresetGetPresetByIDResponseDataConfigViewTypeAudioRoom:
		return true
	}
	return false
}

type PresetGetPresetByIDResponseDataUI struct {
	DesignTokens PresetGetPresetByIDResponseDataUIDesignTokens `json:"design_tokens,required"`
	ConfigDiff   interface{}                                   `json:"config_diff"`
	JSON         presetGetPresetByIDResponseDataUIJSON         `json:"-"`
}

// presetGetPresetByIDResponseDataUIJSON contains the JSON metadata for the struct
// [PresetGetPresetByIDResponseDataUI]
type presetGetPresetByIDResponseDataUIJSON struct {
	DesignTokens apijson.Field
	ConfigDiff   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataUI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataUIJSON) RawJSON() string {
	return r.raw
}

type PresetGetPresetByIDResponseDataUIDesignTokens struct {
	BorderRadius PresetGetPresetByIDResponseDataUIDesignTokensBorderRadius `json:"border_radius,required"`
	BorderWidth  PresetGetPresetByIDResponseDataUIDesignTokensBorderWidth  `json:"border_width,required"`
	Colors       PresetGetPresetByIDResponseDataUIDesignTokensColors       `json:"colors,required"`
	Logo         string                                                    `json:"logo,required"`
	SpacingBase  float64                                                   `json:"spacing_base,required"`
	Theme        PresetGetPresetByIDResponseDataUIDesignTokensTheme        `json:"theme,required"`
	JSON         presetGetPresetByIDResponseDataUIDesignTokensJSON         `json:"-"`
}

// presetGetPresetByIDResponseDataUIDesignTokensJSON contains the JSON metadata for
// the struct [PresetGetPresetByIDResponseDataUIDesignTokens]
type presetGetPresetByIDResponseDataUIDesignTokensJSON struct {
	BorderRadius apijson.Field
	BorderWidth  apijson.Field
	Colors       apijson.Field
	Logo         apijson.Field
	SpacingBase  apijson.Field
	Theme        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataUIDesignTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataUIDesignTokensJSON) RawJSON() string {
	return r.raw
}

type PresetGetPresetByIDResponseDataUIDesignTokensBorderRadius string

const (
	PresetGetPresetByIDResponseDataUIDesignTokensBorderRadiusRounded PresetGetPresetByIDResponseDataUIDesignTokensBorderRadius = "rounded"
)

func (r PresetGetPresetByIDResponseDataUIDesignTokensBorderRadius) IsKnown() bool {
	switch r {
	case PresetGetPresetByIDResponseDataUIDesignTokensBorderRadiusRounded:
		return true
	}
	return false
}

type PresetGetPresetByIDResponseDataUIDesignTokensBorderWidth string

const (
	PresetGetPresetByIDResponseDataUIDesignTokensBorderWidthThin PresetGetPresetByIDResponseDataUIDesignTokensBorderWidth = "thin"
)

func (r PresetGetPresetByIDResponseDataUIDesignTokensBorderWidth) IsKnown() bool {
	switch r {
	case PresetGetPresetByIDResponseDataUIDesignTokensBorderWidthThin:
		return true
	}
	return false
}

type PresetGetPresetByIDResponseDataUIDesignTokensColors struct {
	Background  PresetGetPresetByIDResponseDataUIDesignTokensColorsBackground `json:"background,required"`
	Brand       PresetGetPresetByIDResponseDataUIDesignTokensColorsBrand      `json:"brand,required"`
	Danger      string                                                        `json:"danger,required"`
	Success     string                                                        `json:"success,required"`
	Text        string                                                        `json:"text,required"`
	TextOnBrand string                                                        `json:"text_on_brand,required"`
	VideoBg     string                                                        `json:"video_bg,required"`
	Warning     string                                                        `json:"warning,required"`
	JSON        presetGetPresetByIDResponseDataUIDesignTokensColorsJSON       `json:"-"`
}

// presetGetPresetByIDResponseDataUIDesignTokensColorsJSON contains the JSON
// metadata for the struct [PresetGetPresetByIDResponseDataUIDesignTokensColors]
type presetGetPresetByIDResponseDataUIDesignTokensColorsJSON struct {
	Background  apijson.Field
	Brand       apijson.Field
	Danger      apijson.Field
	Success     apijson.Field
	Text        apijson.Field
	TextOnBrand apijson.Field
	VideoBg     apijson.Field
	Warning     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataUIDesignTokensColors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataUIDesignTokensColorsJSON) RawJSON() string {
	return r.raw
}

type PresetGetPresetByIDResponseDataUIDesignTokensColorsBackground struct {
	Number1000 string                                                            `json:"1000,required"`
	Number600  string                                                            `json:"600,required"`
	Number700  string                                                            `json:"700,required"`
	Number800  string                                                            `json:"800,required"`
	Number900  string                                                            `json:"900,required"`
	JSON       presetGetPresetByIDResponseDataUIDesignTokensColorsBackgroundJSON `json:"-"`
}

// presetGetPresetByIDResponseDataUIDesignTokensColorsBackgroundJSON contains the
// JSON metadata for the struct
// [PresetGetPresetByIDResponseDataUIDesignTokensColorsBackground]
type presetGetPresetByIDResponseDataUIDesignTokensColorsBackgroundJSON struct {
	Number1000  apijson.Field
	Number600   apijson.Field
	Number700   apijson.Field
	Number800   apijson.Field
	Number900   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataUIDesignTokensColorsBackground) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataUIDesignTokensColorsBackgroundJSON) RawJSON() string {
	return r.raw
}

type PresetGetPresetByIDResponseDataUIDesignTokensColorsBrand struct {
	Number300 string                                                       `json:"300,required"`
	Number400 string                                                       `json:"400,required"`
	Number500 string                                                       `json:"500,required"`
	Number600 string                                                       `json:"600,required"`
	Number700 string                                                       `json:"700,required"`
	JSON      presetGetPresetByIDResponseDataUIDesignTokensColorsBrandJSON `json:"-"`
}

// presetGetPresetByIDResponseDataUIDesignTokensColorsBrandJSON contains the JSON
// metadata for the struct
// [PresetGetPresetByIDResponseDataUIDesignTokensColorsBrand]
type presetGetPresetByIDResponseDataUIDesignTokensColorsBrandJSON struct {
	Number300   apijson.Field
	Number400   apijson.Field
	Number500   apijson.Field
	Number600   apijson.Field
	Number700   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataUIDesignTokensColorsBrand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataUIDesignTokensColorsBrandJSON) RawJSON() string {
	return r.raw
}

type PresetGetPresetByIDResponseDataUIDesignTokensTheme string

const (
	PresetGetPresetByIDResponseDataUIDesignTokensThemeDark PresetGetPresetByIDResponseDataUIDesignTokensTheme = "dark"
)

func (r PresetGetPresetByIDResponseDataUIDesignTokensTheme) IsKnown() bool {
	switch r {
	case PresetGetPresetByIDResponseDataUIDesignTokensThemeDark:
		return true
	}
	return false
}

type PresetGetPresetByIDResponseDataPermissions struct {
	// Whether this participant can accept waiting requests
	AcceptWaitingRequests           bool `json:"accept_waiting_requests,required"`
	CanAcceptProductionRequests     bool `json:"can_accept_production_requests,required"`
	CanChangeParticipantPermissions bool `json:"can_change_participant_permissions,required"`
	CanEditDisplayName              bool `json:"can_edit_display_name,required"`
	CanLivestream                   bool `json:"can_livestream,required"`
	CanRecord                       bool `json:"can_record,required"`
	CanSpotlight                    bool `json:"can_spotlight,required"`
	// Chat permissions
	Chat                            PresetGetPresetByIDResponseDataPermissionsChat              `json:"chat,required"`
	ConnectedMeetings               PresetGetPresetByIDResponseDataPermissionsConnectedMeetings `json:"connected_meetings,required"`
	DisableParticipantAudio         bool                                                        `json:"disable_participant_audio,required"`
	DisableParticipantScreensharing bool                                                        `json:"disable_participant_screensharing,required"`
	DisableParticipantVideo         bool                                                        `json:"disable_participant_video,required"`
	// Whether this participant is visible to others or not
	HiddenParticipant bool `json:"hidden_participant,required"`
	KickParticipant   bool `json:"kick_participant,required"`
	// Media permissions
	Media          PresetGetPresetByIDResponseDataPermissionsMedia `json:"media,required"`
	PinParticipant bool                                            `json:"pin_participant,required"`
	// Plugin permissions
	Plugins PresetGetPresetByIDResponseDataPermissionsPlugins `json:"plugins,required"`
	// Poll permissions
	Polls PresetGetPresetByIDResponseDataPermissionsPolls `json:"polls,required"`
	// Type of the recording peer
	RecorderType        PresetGetPresetByIDResponseDataPermissionsRecorderType `json:"recorder_type,required"`
	ShowParticipantList bool                                                   `json:"show_participant_list,required"`
	// Waiting room type
	WaitingRoomType PresetGetPresetByIDResponseDataPermissionsWaitingRoomType `json:"waiting_room_type,required"`
	IsRecorder      bool                                                      `json:"is_recorder"`
	JSON            presetGetPresetByIDResponseDataPermissionsJSON            `json:"-"`
}

// presetGetPresetByIDResponseDataPermissionsJSON contains the JSON metadata for
// the struct [PresetGetPresetByIDResponseDataPermissions]
type presetGetPresetByIDResponseDataPermissionsJSON struct {
	AcceptWaitingRequests           apijson.Field
	CanAcceptProductionRequests     apijson.Field
	CanChangeParticipantPermissions apijson.Field
	CanEditDisplayName              apijson.Field
	CanLivestream                   apijson.Field
	CanRecord                       apijson.Field
	CanSpotlight                    apijson.Field
	Chat                            apijson.Field
	ConnectedMeetings               apijson.Field
	DisableParticipantAudio         apijson.Field
	DisableParticipantScreensharing apijson.Field
	DisableParticipantVideo         apijson.Field
	HiddenParticipant               apijson.Field
	KickParticipant                 apijson.Field
	Media                           apijson.Field
	PinParticipant                  apijson.Field
	Plugins                         apijson.Field
	Polls                           apijson.Field
	RecorderType                    apijson.Field
	ShowParticipantList             apijson.Field
	WaitingRoomType                 apijson.Field
	IsRecorder                      apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataPermissionsJSON) RawJSON() string {
	return r.raw
}

// Chat permissions
type PresetGetPresetByIDResponseDataPermissionsChat struct {
	Private PresetGetPresetByIDResponseDataPermissionsChatPrivate `json:"private,required"`
	Public  PresetGetPresetByIDResponseDataPermissionsChatPublic  `json:"public,required"`
	JSON    presetGetPresetByIDResponseDataPermissionsChatJSON    `json:"-"`
}

// presetGetPresetByIDResponseDataPermissionsChatJSON contains the JSON metadata
// for the struct [PresetGetPresetByIDResponseDataPermissionsChat]
type presetGetPresetByIDResponseDataPermissionsChatJSON struct {
	Private     apijson.Field
	Public      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataPermissionsChat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataPermissionsChatJSON) RawJSON() string {
	return r.raw
}

type PresetGetPresetByIDResponseDataPermissionsChatPrivate struct {
	CanReceive bool                                                      `json:"can_receive,required"`
	CanSend    bool                                                      `json:"can_send,required"`
	Files      bool                                                      `json:"files,required"`
	Text       bool                                                      `json:"text,required"`
	JSON       presetGetPresetByIDResponseDataPermissionsChatPrivateJSON `json:"-"`
}

// presetGetPresetByIDResponseDataPermissionsChatPrivateJSON contains the JSON
// metadata for the struct [PresetGetPresetByIDResponseDataPermissionsChatPrivate]
type presetGetPresetByIDResponseDataPermissionsChatPrivateJSON struct {
	CanReceive  apijson.Field
	CanSend     apijson.Field
	Files       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataPermissionsChatPrivate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataPermissionsChatPrivateJSON) RawJSON() string {
	return r.raw
}

type PresetGetPresetByIDResponseDataPermissionsChatPublic struct {
	// Can send messages in general
	CanSend bool `json:"can_send,required"`
	// Can send file messages
	Files bool `json:"files,required"`
	// Can send text messages
	Text bool                                                     `json:"text,required"`
	JSON presetGetPresetByIDResponseDataPermissionsChatPublicJSON `json:"-"`
}

// presetGetPresetByIDResponseDataPermissionsChatPublicJSON contains the JSON
// metadata for the struct [PresetGetPresetByIDResponseDataPermissionsChatPublic]
type presetGetPresetByIDResponseDataPermissionsChatPublicJSON struct {
	CanSend     apijson.Field
	Files       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataPermissionsChatPublic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataPermissionsChatPublicJSON) RawJSON() string {
	return r.raw
}

type PresetGetPresetByIDResponseDataPermissionsConnectedMeetings struct {
	CanAlterConnectedMeetings  bool                                                            `json:"can_alter_connected_meetings,required"`
	CanSwitchConnectedMeetings bool                                                            `json:"can_switch_connected_meetings,required"`
	CanSwitchToParentMeeting   bool                                                            `json:"can_switch_to_parent_meeting,required"`
	JSON                       presetGetPresetByIDResponseDataPermissionsConnectedMeetingsJSON `json:"-"`
}

// presetGetPresetByIDResponseDataPermissionsConnectedMeetingsJSON contains the
// JSON metadata for the struct
// [PresetGetPresetByIDResponseDataPermissionsConnectedMeetings]
type presetGetPresetByIDResponseDataPermissionsConnectedMeetingsJSON struct {
	CanAlterConnectedMeetings  apijson.Field
	CanSwitchConnectedMeetings apijson.Field
	CanSwitchToParentMeeting   apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataPermissionsConnectedMeetings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataPermissionsConnectedMeetingsJSON) RawJSON() string {
	return r.raw
}

// Media permissions
type PresetGetPresetByIDResponseDataPermissionsMedia struct {
	// Audio permissions
	Audio PresetGetPresetByIDResponseDataPermissionsMediaAudio `json:"audio,required"`
	// Screenshare permissions
	Screenshare PresetGetPresetByIDResponseDataPermissionsMediaScreenshare `json:"screenshare,required"`
	// Video permissions
	Video PresetGetPresetByIDResponseDataPermissionsMediaVideo `json:"video,required"`
	JSON  presetGetPresetByIDResponseDataPermissionsMediaJSON  `json:"-"`
}

// presetGetPresetByIDResponseDataPermissionsMediaJSON contains the JSON metadata
// for the struct [PresetGetPresetByIDResponseDataPermissionsMedia]
type presetGetPresetByIDResponseDataPermissionsMediaJSON struct {
	Audio       apijson.Field
	Screenshare apijson.Field
	Video       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataPermissionsMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataPermissionsMediaJSON) RawJSON() string {
	return r.raw
}

// Audio permissions
type PresetGetPresetByIDResponseDataPermissionsMediaAudio struct {
	// Can produce audio
	CanProduce PresetGetPresetByIDResponseDataPermissionsMediaAudioCanProduce `json:"can_produce,required"`
	JSON       presetGetPresetByIDResponseDataPermissionsMediaAudioJSON       `json:"-"`
}

// presetGetPresetByIDResponseDataPermissionsMediaAudioJSON contains the JSON
// metadata for the struct [PresetGetPresetByIDResponseDataPermissionsMediaAudio]
type presetGetPresetByIDResponseDataPermissionsMediaAudioJSON struct {
	CanProduce  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataPermissionsMediaAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataPermissionsMediaAudioJSON) RawJSON() string {
	return r.raw
}

// Can produce audio
type PresetGetPresetByIDResponseDataPermissionsMediaAudioCanProduce string

const (
	PresetGetPresetByIDResponseDataPermissionsMediaAudioCanProduceAllowed    PresetGetPresetByIDResponseDataPermissionsMediaAudioCanProduce = "ALLOWED"
	PresetGetPresetByIDResponseDataPermissionsMediaAudioCanProduceNotAllowed PresetGetPresetByIDResponseDataPermissionsMediaAudioCanProduce = "NOT_ALLOWED"
	PresetGetPresetByIDResponseDataPermissionsMediaAudioCanProduceCanRequest PresetGetPresetByIDResponseDataPermissionsMediaAudioCanProduce = "CAN_REQUEST"
)

func (r PresetGetPresetByIDResponseDataPermissionsMediaAudioCanProduce) IsKnown() bool {
	switch r {
	case PresetGetPresetByIDResponseDataPermissionsMediaAudioCanProduceAllowed, PresetGetPresetByIDResponseDataPermissionsMediaAudioCanProduceNotAllowed, PresetGetPresetByIDResponseDataPermissionsMediaAudioCanProduceCanRequest:
		return true
	}
	return false
}

// Screenshare permissions
type PresetGetPresetByIDResponseDataPermissionsMediaScreenshare struct {
	// Can produce screen share video
	CanProduce PresetGetPresetByIDResponseDataPermissionsMediaScreenshareCanProduce `json:"can_produce,required"`
	JSON       presetGetPresetByIDResponseDataPermissionsMediaScreenshareJSON       `json:"-"`
}

// presetGetPresetByIDResponseDataPermissionsMediaScreenshareJSON contains the JSON
// metadata for the struct
// [PresetGetPresetByIDResponseDataPermissionsMediaScreenshare]
type presetGetPresetByIDResponseDataPermissionsMediaScreenshareJSON struct {
	CanProduce  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataPermissionsMediaScreenshare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataPermissionsMediaScreenshareJSON) RawJSON() string {
	return r.raw
}

// Can produce screen share video
type PresetGetPresetByIDResponseDataPermissionsMediaScreenshareCanProduce string

const (
	PresetGetPresetByIDResponseDataPermissionsMediaScreenshareCanProduceAllowed    PresetGetPresetByIDResponseDataPermissionsMediaScreenshareCanProduce = "ALLOWED"
	PresetGetPresetByIDResponseDataPermissionsMediaScreenshareCanProduceNotAllowed PresetGetPresetByIDResponseDataPermissionsMediaScreenshareCanProduce = "NOT_ALLOWED"
	PresetGetPresetByIDResponseDataPermissionsMediaScreenshareCanProduceCanRequest PresetGetPresetByIDResponseDataPermissionsMediaScreenshareCanProduce = "CAN_REQUEST"
)

func (r PresetGetPresetByIDResponseDataPermissionsMediaScreenshareCanProduce) IsKnown() bool {
	switch r {
	case PresetGetPresetByIDResponseDataPermissionsMediaScreenshareCanProduceAllowed, PresetGetPresetByIDResponseDataPermissionsMediaScreenshareCanProduceNotAllowed, PresetGetPresetByIDResponseDataPermissionsMediaScreenshareCanProduceCanRequest:
		return true
	}
	return false
}

// Video permissions
type PresetGetPresetByIDResponseDataPermissionsMediaVideo struct {
	// Can produce video
	CanProduce PresetGetPresetByIDResponseDataPermissionsMediaVideoCanProduce `json:"can_produce,required"`
	JSON       presetGetPresetByIDResponseDataPermissionsMediaVideoJSON       `json:"-"`
}

// presetGetPresetByIDResponseDataPermissionsMediaVideoJSON contains the JSON
// metadata for the struct [PresetGetPresetByIDResponseDataPermissionsMediaVideo]
type presetGetPresetByIDResponseDataPermissionsMediaVideoJSON struct {
	CanProduce  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataPermissionsMediaVideo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataPermissionsMediaVideoJSON) RawJSON() string {
	return r.raw
}

// Can produce video
type PresetGetPresetByIDResponseDataPermissionsMediaVideoCanProduce string

const (
	PresetGetPresetByIDResponseDataPermissionsMediaVideoCanProduceAllowed    PresetGetPresetByIDResponseDataPermissionsMediaVideoCanProduce = "ALLOWED"
	PresetGetPresetByIDResponseDataPermissionsMediaVideoCanProduceNotAllowed PresetGetPresetByIDResponseDataPermissionsMediaVideoCanProduce = "NOT_ALLOWED"
	PresetGetPresetByIDResponseDataPermissionsMediaVideoCanProduceCanRequest PresetGetPresetByIDResponseDataPermissionsMediaVideoCanProduce = "CAN_REQUEST"
)

func (r PresetGetPresetByIDResponseDataPermissionsMediaVideoCanProduce) IsKnown() bool {
	switch r {
	case PresetGetPresetByIDResponseDataPermissionsMediaVideoCanProduceAllowed, PresetGetPresetByIDResponseDataPermissionsMediaVideoCanProduceNotAllowed, PresetGetPresetByIDResponseDataPermissionsMediaVideoCanProduceCanRequest:
		return true
	}
	return false
}

// Plugin permissions
type PresetGetPresetByIDResponseDataPermissionsPlugins struct {
	// Can close plugins that are already open
	CanClose bool `json:"can_close,required"`
	// Can edit plugin config
	CanEditConfig bool `json:"can_edit_config,required"`
	// Can start plugins
	CanStart bool                                                         `json:"can_start,required"`
	Config   PresetGetPresetByIDResponseDataPermissionsPluginsConfigUnion `json:"config,required" format:"uuid"`
	JSON     presetGetPresetByIDResponseDataPermissionsPluginsJSON        `json:"-"`
}

// presetGetPresetByIDResponseDataPermissionsPluginsJSON contains the JSON metadata
// for the struct [PresetGetPresetByIDResponseDataPermissionsPlugins]
type presetGetPresetByIDResponseDataPermissionsPluginsJSON struct {
	CanClose      apijson.Field
	CanEditConfig apijson.Field
	CanStart      apijson.Field
	Config        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataPermissionsPlugins) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataPermissionsPluginsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or
// [PresetGetPresetByIDResponseDataPermissionsPluginsConfigObject].
type PresetGetPresetByIDResponseDataPermissionsPluginsConfigUnion interface {
	ImplementsPresetGetPresetByIDResponseDataPermissionsPluginsConfigUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PresetGetPresetByIDResponseDataPermissionsPluginsConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PresetGetPresetByIDResponseDataPermissionsPluginsConfigObject{}),
		},
	)
}

type PresetGetPresetByIDResponseDataPermissionsPluginsConfigObject struct {
	AccessControl   PresetGetPresetByIDResponseDataPermissionsPluginsConfigObjectAccessControl `json:"access_control,required"`
	HandlesViewOnly bool                                                                       `json:"handles_view_only,required"`
	JSON            presetGetPresetByIDResponseDataPermissionsPluginsConfigObjectJSON          `json:"-"`
}

// presetGetPresetByIDResponseDataPermissionsPluginsConfigObjectJSON contains the
// JSON metadata for the struct
// [PresetGetPresetByIDResponseDataPermissionsPluginsConfigObject]
type presetGetPresetByIDResponseDataPermissionsPluginsConfigObjectJSON struct {
	AccessControl   apijson.Field
	HandlesViewOnly apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataPermissionsPluginsConfigObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataPermissionsPluginsConfigObjectJSON) RawJSON() string {
	return r.raw
}

func (r PresetGetPresetByIDResponseDataPermissionsPluginsConfigObject) ImplementsPresetGetPresetByIDResponseDataPermissionsPluginsConfigUnion() {
}

type PresetGetPresetByIDResponseDataPermissionsPluginsConfigObjectAccessControl string

const (
	PresetGetPresetByIDResponseDataPermissionsPluginsConfigObjectAccessControlFullAccess PresetGetPresetByIDResponseDataPermissionsPluginsConfigObjectAccessControl = "FULL_ACCESS"
	PresetGetPresetByIDResponseDataPermissionsPluginsConfigObjectAccessControlViewOnly   PresetGetPresetByIDResponseDataPermissionsPluginsConfigObjectAccessControl = "VIEW_ONLY"
)

func (r PresetGetPresetByIDResponseDataPermissionsPluginsConfigObjectAccessControl) IsKnown() bool {
	switch r {
	case PresetGetPresetByIDResponseDataPermissionsPluginsConfigObjectAccessControlFullAccess, PresetGetPresetByIDResponseDataPermissionsPluginsConfigObjectAccessControlViewOnly:
		return true
	}
	return false
}

// Poll permissions
type PresetGetPresetByIDResponseDataPermissionsPolls struct {
	// Can create polls
	CanCreate bool `json:"can_create,required"`
	// Can view polls
	CanView bool `json:"can_view,required"`
	// Can vote on polls
	CanVote bool                                                `json:"can_vote,required"`
	JSON    presetGetPresetByIDResponseDataPermissionsPollsJSON `json:"-"`
}

// presetGetPresetByIDResponseDataPermissionsPollsJSON contains the JSON metadata
// for the struct [PresetGetPresetByIDResponseDataPermissionsPolls]
type presetGetPresetByIDResponseDataPermissionsPollsJSON struct {
	CanCreate   apijson.Field
	CanView     apijson.Field
	CanVote     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresetGetPresetByIDResponseDataPermissionsPolls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presetGetPresetByIDResponseDataPermissionsPollsJSON) RawJSON() string {
	return r.raw
}

// Type of the recording peer
type PresetGetPresetByIDResponseDataPermissionsRecorderType string

const (
	PresetGetPresetByIDResponseDataPermissionsRecorderTypeRecorder     PresetGetPresetByIDResponseDataPermissionsRecorderType = "RECORDER"
	PresetGetPresetByIDResponseDataPermissionsRecorderTypeLivestreamer PresetGetPresetByIDResponseDataPermissionsRecorderType = "LIVESTREAMER"
	PresetGetPresetByIDResponseDataPermissionsRecorderTypeNone         PresetGetPresetByIDResponseDataPermissionsRecorderType = "NONE"
)

func (r PresetGetPresetByIDResponseDataPermissionsRecorderType) IsKnown() bool {
	switch r {
	case PresetGetPresetByIDResponseDataPermissionsRecorderTypeRecorder, PresetGetPresetByIDResponseDataPermissionsRecorderTypeLivestreamer, PresetGetPresetByIDResponseDataPermissionsRecorderTypeNone:
		return true
	}
	return false
}

// Waiting room type
type PresetGetPresetByIDResponseDataPermissionsWaitingRoomType string

const (
	PresetGetPresetByIDResponseDataPermissionsWaitingRoomTypeSkip                  PresetGetPresetByIDResponseDataPermissionsWaitingRoomType = "SKIP"
	PresetGetPresetByIDResponseDataPermissionsWaitingRoomTypeOnPrivilegedUserEntry PresetGetPresetByIDResponseDataPermissionsWaitingRoomType = "ON_PRIVILEGED_USER_ENTRY"
	PresetGetPresetByIDResponseDataPermissionsWaitingRoomTypeSkipOnAccept          PresetGetPresetByIDResponseDataPermissionsWaitingRoomType = "SKIP_ON_ACCEPT"
)

func (r PresetGetPresetByIDResponseDataPermissionsWaitingRoomType) IsKnown() bool {
	switch r {
	case PresetGetPresetByIDResponseDataPermissionsWaitingRoomTypeSkip, PresetGetPresetByIDResponseDataPermissionsWaitingRoomTypeOnPrivilegedUserEntry, PresetGetPresetByIDResponseDataPermissionsWaitingRoomTypeSkipOnAccept:
		return true
	}
	return false
}

type PresetNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string]                `path:"account_id,required"`
	Config    param.Field[PresetNewParamsConfig] `json:"config,required"`
	// Name of the preset
	Name        param.Field[string]                     `json:"name,required"`
	UI          param.Field[PresetNewParamsUI]          `json:"ui,required"`
	Permissions param.Field[PresetNewParamsPermissions] `json:"permissions"`
}

func (r PresetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetNewParamsConfig struct {
	// Maximum number of screen shares that can be active at a given time
	MaxScreenshareCount param.Field[int64] `json:"max_screenshare_count,required"`
	// Maximum number of streams that are visible on a device
	MaxVideoStreams param.Field[PresetNewParamsConfigMaxVideoStreams] `json:"max_video_streams,required"`
	// Media configuration options. eg: Video quality
	Media param.Field[PresetNewParamsConfigMedia] `json:"media,required"`
	// Type of the meeting
	ViewType param.Field[PresetNewParamsConfigViewType] `json:"view_type,required"`
}

func (r PresetNewParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Maximum number of streams that are visible on a device
type PresetNewParamsConfigMaxVideoStreams struct {
	// Maximum number of video streams visible on desktop devices
	Desktop param.Field[int64] `json:"desktop,required"`
	// Maximum number of streams visible on mobile devices
	Mobile param.Field[int64] `json:"mobile,required"`
}

func (r PresetNewParamsConfigMaxVideoStreams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Media configuration options. eg: Video quality
type PresetNewParamsConfigMedia struct {
	// Configuration options for participant screen shares
	Screenshare param.Field[PresetNewParamsConfigMediaScreenshare] `json:"screenshare,required"`
	// Configuration options for participant videos
	Video param.Field[PresetNewParamsConfigMediaVideo] `json:"video,required"`
	// Control options for Audio quality.
	Audio param.Field[PresetNewParamsConfigMediaAudio] `json:"audio"`
}

func (r PresetNewParamsConfigMedia) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration options for participant screen shares
type PresetNewParamsConfigMediaScreenshare struct {
	// Frame rate of screen share
	FrameRate param.Field[int64] `json:"frame_rate,required"`
	// Quality of screen share
	Quality param.Field[PresetNewParamsConfigMediaScreenshareQuality] `json:"quality,required"`
}

func (r PresetNewParamsConfigMediaScreenshare) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Quality of screen share
type PresetNewParamsConfigMediaScreenshareQuality string

const (
	PresetNewParamsConfigMediaScreenshareQualityHD   PresetNewParamsConfigMediaScreenshareQuality = "hd"
	PresetNewParamsConfigMediaScreenshareQualityVga  PresetNewParamsConfigMediaScreenshareQuality = "vga"
	PresetNewParamsConfigMediaScreenshareQualityQvga PresetNewParamsConfigMediaScreenshareQuality = "qvga"
)

func (r PresetNewParamsConfigMediaScreenshareQuality) IsKnown() bool {
	switch r {
	case PresetNewParamsConfigMediaScreenshareQualityHD, PresetNewParamsConfigMediaScreenshareQualityVga, PresetNewParamsConfigMediaScreenshareQualityQvga:
		return true
	}
	return false
}

// Configuration options for participant videos
type PresetNewParamsConfigMediaVideo struct {
	// Frame rate of participants' video
	FrameRate param.Field[int64] `json:"frame_rate,required"`
	// Video quality of participants
	Quality param.Field[PresetNewParamsConfigMediaVideoQuality] `json:"quality,required"`
}

func (r PresetNewParamsConfigMediaVideo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Video quality of participants
type PresetNewParamsConfigMediaVideoQuality string

const (
	PresetNewParamsConfigMediaVideoQualityHD   PresetNewParamsConfigMediaVideoQuality = "hd"
	PresetNewParamsConfigMediaVideoQualityVga  PresetNewParamsConfigMediaVideoQuality = "vga"
	PresetNewParamsConfigMediaVideoQualityQvga PresetNewParamsConfigMediaVideoQuality = "qvga"
)

func (r PresetNewParamsConfigMediaVideoQuality) IsKnown() bool {
	switch r {
	case PresetNewParamsConfigMediaVideoQualityHD, PresetNewParamsConfigMediaVideoQualityVga, PresetNewParamsConfigMediaVideoQualityQvga:
		return true
	}
	return false
}

// Control options for Audio quality.
type PresetNewParamsConfigMediaAudio struct {
	// Enable High Quality Audio for your meetings
	EnableHighBitrate param.Field[bool] `json:"enable_high_bitrate"`
	// Enable Stereo for your meetings
	EnableStereo param.Field[bool] `json:"enable_stereo"`
}

func (r PresetNewParamsConfigMediaAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the meeting
type PresetNewParamsConfigViewType string

const (
	PresetNewParamsConfigViewTypeGroupCall PresetNewParamsConfigViewType = "GROUP_CALL"
	PresetNewParamsConfigViewTypeWebinar   PresetNewParamsConfigViewType = "WEBINAR"
	PresetNewParamsConfigViewTypeAudioRoom PresetNewParamsConfigViewType = "AUDIO_ROOM"
)

func (r PresetNewParamsConfigViewType) IsKnown() bool {
	switch r {
	case PresetNewParamsConfigViewTypeGroupCall, PresetNewParamsConfigViewTypeWebinar, PresetNewParamsConfigViewTypeAudioRoom:
		return true
	}
	return false
}

type PresetNewParamsUI struct {
	DesignTokens param.Field[PresetNewParamsUIDesignTokens] `json:"design_tokens,required"`
	ConfigDiff   param.Field[interface{}]                   `json:"config_diff"`
}

func (r PresetNewParamsUI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetNewParamsUIDesignTokens struct {
	BorderRadius param.Field[PresetNewParamsUIDesignTokensBorderRadius] `json:"border_radius,required"`
	BorderWidth  param.Field[PresetNewParamsUIDesignTokensBorderWidth]  `json:"border_width,required"`
	Colors       param.Field[PresetNewParamsUIDesignTokensColors]       `json:"colors,required"`
	Logo         param.Field[string]                                    `json:"logo,required"`
	SpacingBase  param.Field[float64]                                   `json:"spacing_base,required"`
	Theme        param.Field[PresetNewParamsUIDesignTokensTheme]        `json:"theme,required"`
}

func (r PresetNewParamsUIDesignTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetNewParamsUIDesignTokensBorderRadius string

const (
	PresetNewParamsUIDesignTokensBorderRadiusRounded PresetNewParamsUIDesignTokensBorderRadius = "rounded"
)

func (r PresetNewParamsUIDesignTokensBorderRadius) IsKnown() bool {
	switch r {
	case PresetNewParamsUIDesignTokensBorderRadiusRounded:
		return true
	}
	return false
}

type PresetNewParamsUIDesignTokensBorderWidth string

const (
	PresetNewParamsUIDesignTokensBorderWidthThin PresetNewParamsUIDesignTokensBorderWidth = "thin"
)

func (r PresetNewParamsUIDesignTokensBorderWidth) IsKnown() bool {
	switch r {
	case PresetNewParamsUIDesignTokensBorderWidthThin:
		return true
	}
	return false
}

type PresetNewParamsUIDesignTokensColors struct {
	Background  param.Field[PresetNewParamsUIDesignTokensColorsBackground] `json:"background,required"`
	Brand       param.Field[PresetNewParamsUIDesignTokensColorsBrand]      `json:"brand,required"`
	Danger      param.Field[string]                                        `json:"danger,required"`
	Success     param.Field[string]                                        `json:"success,required"`
	Text        param.Field[string]                                        `json:"text,required"`
	TextOnBrand param.Field[string]                                        `json:"text_on_brand,required"`
	VideoBg     param.Field[string]                                        `json:"video_bg,required"`
	Warning     param.Field[string]                                        `json:"warning,required"`
}

func (r PresetNewParamsUIDesignTokensColors) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetNewParamsUIDesignTokensColorsBackground struct {
	Number1000 param.Field[string] `json:"1000,required"`
	Number600  param.Field[string] `json:"600,required"`
	Number700  param.Field[string] `json:"700,required"`
	Number800  param.Field[string] `json:"800,required"`
	Number900  param.Field[string] `json:"900,required"`
}

func (r PresetNewParamsUIDesignTokensColorsBackground) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetNewParamsUIDesignTokensColorsBrand struct {
	Number300 param.Field[string] `json:"300,required"`
	Number400 param.Field[string] `json:"400,required"`
	Number500 param.Field[string] `json:"500,required"`
	Number600 param.Field[string] `json:"600,required"`
	Number700 param.Field[string] `json:"700,required"`
}

func (r PresetNewParamsUIDesignTokensColorsBrand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetNewParamsUIDesignTokensTheme string

const (
	PresetNewParamsUIDesignTokensThemeDark PresetNewParamsUIDesignTokensTheme = "dark"
)

func (r PresetNewParamsUIDesignTokensTheme) IsKnown() bool {
	switch r {
	case PresetNewParamsUIDesignTokensThemeDark:
		return true
	}
	return false
}

type PresetNewParamsPermissions struct {
	// Whether this participant can accept waiting requests
	AcceptWaitingRequests           param.Field[bool] `json:"accept_waiting_requests,required"`
	CanAcceptProductionRequests     param.Field[bool] `json:"can_accept_production_requests,required"`
	CanChangeParticipantPermissions param.Field[bool] `json:"can_change_participant_permissions,required"`
	CanEditDisplayName              param.Field[bool] `json:"can_edit_display_name,required"`
	CanLivestream                   param.Field[bool] `json:"can_livestream,required"`
	CanRecord                       param.Field[bool] `json:"can_record,required"`
	CanSpotlight                    param.Field[bool] `json:"can_spotlight,required"`
	// Chat permissions
	Chat                            param.Field[PresetNewParamsPermissionsChat]              `json:"chat,required"`
	ConnectedMeetings               param.Field[PresetNewParamsPermissionsConnectedMeetings] `json:"connected_meetings,required"`
	DisableParticipantAudio         param.Field[bool]                                        `json:"disable_participant_audio,required"`
	DisableParticipantScreensharing param.Field[bool]                                        `json:"disable_participant_screensharing,required"`
	DisableParticipantVideo         param.Field[bool]                                        `json:"disable_participant_video,required"`
	// Whether this participant is visible to others or not
	HiddenParticipant param.Field[bool] `json:"hidden_participant,required"`
	KickParticipant   param.Field[bool] `json:"kick_participant,required"`
	// Media permissions
	Media          param.Field[PresetNewParamsPermissionsMedia] `json:"media,required"`
	PinParticipant param.Field[bool]                            `json:"pin_participant,required"`
	// Plugin permissions
	Plugins param.Field[PresetNewParamsPermissionsPlugins] `json:"plugins,required"`
	// Poll permissions
	Polls param.Field[PresetNewParamsPermissionsPolls] `json:"polls,required"`
	// Type of the recording peer
	RecorderType        param.Field[PresetNewParamsPermissionsRecorderType] `json:"recorder_type,required"`
	ShowParticipantList param.Field[bool]                                   `json:"show_participant_list,required"`
	// Waiting room type
	WaitingRoomType param.Field[PresetNewParamsPermissionsWaitingRoomType] `json:"waiting_room_type,required"`
	IsRecorder      param.Field[bool]                                      `json:"is_recorder"`
}

func (r PresetNewParamsPermissions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Chat permissions
type PresetNewParamsPermissionsChat struct {
	Private param.Field[PresetNewParamsPermissionsChatPrivate] `json:"private,required"`
	Public  param.Field[PresetNewParamsPermissionsChatPublic]  `json:"public,required"`
}

func (r PresetNewParamsPermissionsChat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetNewParamsPermissionsChatPrivate struct {
	CanReceive param.Field[bool] `json:"can_receive,required"`
	CanSend    param.Field[bool] `json:"can_send,required"`
	Files      param.Field[bool] `json:"files,required"`
	Text       param.Field[bool] `json:"text,required"`
}

func (r PresetNewParamsPermissionsChatPrivate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetNewParamsPermissionsChatPublic struct {
	// Can send messages in general
	CanSend param.Field[bool] `json:"can_send,required"`
	// Can send file messages
	Files param.Field[bool] `json:"files,required"`
	// Can send text messages
	Text param.Field[bool] `json:"text,required"`
}

func (r PresetNewParamsPermissionsChatPublic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetNewParamsPermissionsConnectedMeetings struct {
	CanAlterConnectedMeetings  param.Field[bool] `json:"can_alter_connected_meetings,required"`
	CanSwitchConnectedMeetings param.Field[bool] `json:"can_switch_connected_meetings,required"`
	CanSwitchToParentMeeting   param.Field[bool] `json:"can_switch_to_parent_meeting,required"`
}

func (r PresetNewParamsPermissionsConnectedMeetings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Media permissions
type PresetNewParamsPermissionsMedia struct {
	// Audio permissions
	Audio param.Field[PresetNewParamsPermissionsMediaAudio] `json:"audio,required"`
	// Screenshare permissions
	Screenshare param.Field[PresetNewParamsPermissionsMediaScreenshare] `json:"screenshare,required"`
	// Video permissions
	Video param.Field[PresetNewParamsPermissionsMediaVideo] `json:"video,required"`
}

func (r PresetNewParamsPermissionsMedia) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Audio permissions
type PresetNewParamsPermissionsMediaAudio struct {
	// Can produce audio
	CanProduce param.Field[PresetNewParamsPermissionsMediaAudioCanProduce] `json:"can_produce,required"`
}

func (r PresetNewParamsPermissionsMediaAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Can produce audio
type PresetNewParamsPermissionsMediaAudioCanProduce string

const (
	PresetNewParamsPermissionsMediaAudioCanProduceAllowed    PresetNewParamsPermissionsMediaAudioCanProduce = "ALLOWED"
	PresetNewParamsPermissionsMediaAudioCanProduceNotAllowed PresetNewParamsPermissionsMediaAudioCanProduce = "NOT_ALLOWED"
	PresetNewParamsPermissionsMediaAudioCanProduceCanRequest PresetNewParamsPermissionsMediaAudioCanProduce = "CAN_REQUEST"
)

func (r PresetNewParamsPermissionsMediaAudioCanProduce) IsKnown() bool {
	switch r {
	case PresetNewParamsPermissionsMediaAudioCanProduceAllowed, PresetNewParamsPermissionsMediaAudioCanProduceNotAllowed, PresetNewParamsPermissionsMediaAudioCanProduceCanRequest:
		return true
	}
	return false
}

// Screenshare permissions
type PresetNewParamsPermissionsMediaScreenshare struct {
	// Can produce screen share video
	CanProduce param.Field[PresetNewParamsPermissionsMediaScreenshareCanProduce] `json:"can_produce,required"`
}

func (r PresetNewParamsPermissionsMediaScreenshare) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Can produce screen share video
type PresetNewParamsPermissionsMediaScreenshareCanProduce string

const (
	PresetNewParamsPermissionsMediaScreenshareCanProduceAllowed    PresetNewParamsPermissionsMediaScreenshareCanProduce = "ALLOWED"
	PresetNewParamsPermissionsMediaScreenshareCanProduceNotAllowed PresetNewParamsPermissionsMediaScreenshareCanProduce = "NOT_ALLOWED"
	PresetNewParamsPermissionsMediaScreenshareCanProduceCanRequest PresetNewParamsPermissionsMediaScreenshareCanProduce = "CAN_REQUEST"
)

func (r PresetNewParamsPermissionsMediaScreenshareCanProduce) IsKnown() bool {
	switch r {
	case PresetNewParamsPermissionsMediaScreenshareCanProduceAllowed, PresetNewParamsPermissionsMediaScreenshareCanProduceNotAllowed, PresetNewParamsPermissionsMediaScreenshareCanProduceCanRequest:
		return true
	}
	return false
}

// Video permissions
type PresetNewParamsPermissionsMediaVideo struct {
	// Can produce video
	CanProduce param.Field[PresetNewParamsPermissionsMediaVideoCanProduce] `json:"can_produce,required"`
}

func (r PresetNewParamsPermissionsMediaVideo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Can produce video
type PresetNewParamsPermissionsMediaVideoCanProduce string

const (
	PresetNewParamsPermissionsMediaVideoCanProduceAllowed    PresetNewParamsPermissionsMediaVideoCanProduce = "ALLOWED"
	PresetNewParamsPermissionsMediaVideoCanProduceNotAllowed PresetNewParamsPermissionsMediaVideoCanProduce = "NOT_ALLOWED"
	PresetNewParamsPermissionsMediaVideoCanProduceCanRequest PresetNewParamsPermissionsMediaVideoCanProduce = "CAN_REQUEST"
)

func (r PresetNewParamsPermissionsMediaVideoCanProduce) IsKnown() bool {
	switch r {
	case PresetNewParamsPermissionsMediaVideoCanProduceAllowed, PresetNewParamsPermissionsMediaVideoCanProduceNotAllowed, PresetNewParamsPermissionsMediaVideoCanProduceCanRequest:
		return true
	}
	return false
}

// Plugin permissions
type PresetNewParamsPermissionsPlugins struct {
	// Can close plugins that are already open
	CanClose param.Field[bool] `json:"can_close,required"`
	// Can edit plugin config
	CanEditConfig param.Field[bool] `json:"can_edit_config,required"`
	// Can start plugins
	CanStart param.Field[bool]                                         `json:"can_start,required"`
	Config   param.Field[PresetNewParamsPermissionsPluginsConfigUnion] `json:"config,required" format:"uuid"`
}

func (r PresetNewParamsPermissionsPlugins) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [realtime_kit.PresetNewParamsPermissionsPluginsConfigObject].
type PresetNewParamsPermissionsPluginsConfigUnion interface {
	ImplementsPresetNewParamsPermissionsPluginsConfigUnion()
}

type PresetNewParamsPermissionsPluginsConfigObject struct {
	AccessControl   param.Field[PresetNewParamsPermissionsPluginsConfigObjectAccessControl] `json:"access_control,required"`
	HandlesViewOnly param.Field[bool]                                                       `json:"handles_view_only,required"`
}

func (r PresetNewParamsPermissionsPluginsConfigObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PresetNewParamsPermissionsPluginsConfigObject) ImplementsPresetNewParamsPermissionsPluginsConfigUnion() {
}

type PresetNewParamsPermissionsPluginsConfigObjectAccessControl string

const (
	PresetNewParamsPermissionsPluginsConfigObjectAccessControlFullAccess PresetNewParamsPermissionsPluginsConfigObjectAccessControl = "FULL_ACCESS"
	PresetNewParamsPermissionsPluginsConfigObjectAccessControlViewOnly   PresetNewParamsPermissionsPluginsConfigObjectAccessControl = "VIEW_ONLY"
)

func (r PresetNewParamsPermissionsPluginsConfigObjectAccessControl) IsKnown() bool {
	switch r {
	case PresetNewParamsPermissionsPluginsConfigObjectAccessControlFullAccess, PresetNewParamsPermissionsPluginsConfigObjectAccessControlViewOnly:
		return true
	}
	return false
}

// Poll permissions
type PresetNewParamsPermissionsPolls struct {
	// Can create polls
	CanCreate param.Field[bool] `json:"can_create,required"`
	// Can view polls
	CanView param.Field[bool] `json:"can_view,required"`
	// Can vote on polls
	CanVote param.Field[bool] `json:"can_vote,required"`
}

func (r PresetNewParamsPermissionsPolls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the recording peer
type PresetNewParamsPermissionsRecorderType string

const (
	PresetNewParamsPermissionsRecorderTypeRecorder     PresetNewParamsPermissionsRecorderType = "RECORDER"
	PresetNewParamsPermissionsRecorderTypeLivestreamer PresetNewParamsPermissionsRecorderType = "LIVESTREAMER"
	PresetNewParamsPermissionsRecorderTypeNone         PresetNewParamsPermissionsRecorderType = "NONE"
)

func (r PresetNewParamsPermissionsRecorderType) IsKnown() bool {
	switch r {
	case PresetNewParamsPermissionsRecorderTypeRecorder, PresetNewParamsPermissionsRecorderTypeLivestreamer, PresetNewParamsPermissionsRecorderTypeNone:
		return true
	}
	return false
}

// Waiting room type
type PresetNewParamsPermissionsWaitingRoomType string

const (
	PresetNewParamsPermissionsWaitingRoomTypeSkip                  PresetNewParamsPermissionsWaitingRoomType = "SKIP"
	PresetNewParamsPermissionsWaitingRoomTypeOnPrivilegedUserEntry PresetNewParamsPermissionsWaitingRoomType = "ON_PRIVILEGED_USER_ENTRY"
	PresetNewParamsPermissionsWaitingRoomTypeSkipOnAccept          PresetNewParamsPermissionsWaitingRoomType = "SKIP_ON_ACCEPT"
)

func (r PresetNewParamsPermissionsWaitingRoomType) IsKnown() bool {
	switch r {
	case PresetNewParamsPermissionsWaitingRoomTypeSkip, PresetNewParamsPermissionsWaitingRoomTypeOnPrivilegedUserEntry, PresetNewParamsPermissionsWaitingRoomTypeSkipOnAccept:
		return true
	}
	return false
}

type PresetUpdateParams struct {
	// The account identifier tag.
	AccountID param.Field[string]                   `path:"account_id,required"`
	Config    param.Field[PresetUpdateParamsConfig] `json:"config"`
	// Name of the preset
	Name        param.Field[string]                        `json:"name"`
	Permissions param.Field[PresetUpdateParamsPermissions] `json:"permissions"`
	UI          param.Field[PresetUpdateParamsUI]          `json:"ui"`
}

func (r PresetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetUpdateParamsConfig struct {
	// Maximum number of screen shares that can be active at a given time
	MaxScreenshareCount param.Field[int64] `json:"max_screenshare_count"`
	// Maximum number of streams that are visible on a device
	MaxVideoStreams param.Field[PresetUpdateParamsConfigMaxVideoStreams] `json:"max_video_streams"`
	// Media configuration options. eg: Video quality
	Media param.Field[PresetUpdateParamsConfigMedia] `json:"media"`
	// Type of the meeting
	ViewType param.Field[PresetUpdateParamsConfigViewType] `json:"view_type"`
}

func (r PresetUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Maximum number of streams that are visible on a device
type PresetUpdateParamsConfigMaxVideoStreams struct {
	// Maximum number of video streams visible on desktop devices
	Desktop param.Field[int64] `json:"desktop"`
	// Maximum number of streams visible on mobile devices
	Mobile param.Field[int64] `json:"mobile"`
}

func (r PresetUpdateParamsConfigMaxVideoStreams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Media configuration options. eg: Video quality
type PresetUpdateParamsConfigMedia struct {
	// Configuration options for participant screen shares
	Screenshare param.Field[PresetUpdateParamsConfigMediaScreenshare] `json:"screenshare"`
	// Configuration options for participant videos
	Video param.Field[PresetUpdateParamsConfigMediaVideo] `json:"video"`
}

func (r PresetUpdateParamsConfigMedia) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration options for participant screen shares
type PresetUpdateParamsConfigMediaScreenshare struct {
	// Frame rate of screen share
	FrameRate param.Field[int64] `json:"frame_rate"`
	// Quality of screen share
	Quality param.Field[PresetUpdateParamsConfigMediaScreenshareQuality] `json:"quality"`
}

func (r PresetUpdateParamsConfigMediaScreenshare) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Quality of screen share
type PresetUpdateParamsConfigMediaScreenshareQuality string

const (
	PresetUpdateParamsConfigMediaScreenshareQualityHD   PresetUpdateParamsConfigMediaScreenshareQuality = "hd"
	PresetUpdateParamsConfigMediaScreenshareQualityVga  PresetUpdateParamsConfigMediaScreenshareQuality = "vga"
	PresetUpdateParamsConfigMediaScreenshareQualityQvga PresetUpdateParamsConfigMediaScreenshareQuality = "qvga"
)

func (r PresetUpdateParamsConfigMediaScreenshareQuality) IsKnown() bool {
	switch r {
	case PresetUpdateParamsConfigMediaScreenshareQualityHD, PresetUpdateParamsConfigMediaScreenshareQualityVga, PresetUpdateParamsConfigMediaScreenshareQualityQvga:
		return true
	}
	return false
}

// Configuration options for participant videos
type PresetUpdateParamsConfigMediaVideo struct {
	// Frame rate of participants' video
	FrameRate param.Field[int64] `json:"frame_rate"`
	// Video quality of participants
	Quality param.Field[PresetUpdateParamsConfigMediaVideoQuality] `json:"quality"`
}

func (r PresetUpdateParamsConfigMediaVideo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Video quality of participants
type PresetUpdateParamsConfigMediaVideoQuality string

const (
	PresetUpdateParamsConfigMediaVideoQualityHD   PresetUpdateParamsConfigMediaVideoQuality = "hd"
	PresetUpdateParamsConfigMediaVideoQualityVga  PresetUpdateParamsConfigMediaVideoQuality = "vga"
	PresetUpdateParamsConfigMediaVideoQualityQvga PresetUpdateParamsConfigMediaVideoQuality = "qvga"
)

func (r PresetUpdateParamsConfigMediaVideoQuality) IsKnown() bool {
	switch r {
	case PresetUpdateParamsConfigMediaVideoQualityHD, PresetUpdateParamsConfigMediaVideoQualityVga, PresetUpdateParamsConfigMediaVideoQualityQvga:
		return true
	}
	return false
}

// Type of the meeting
type PresetUpdateParamsConfigViewType string

const (
	PresetUpdateParamsConfigViewTypeGroupCall PresetUpdateParamsConfigViewType = "GROUP_CALL"
	PresetUpdateParamsConfigViewTypeWebinar   PresetUpdateParamsConfigViewType = "WEBINAR"
	PresetUpdateParamsConfigViewTypeAudioRoom PresetUpdateParamsConfigViewType = "AUDIO_ROOM"
)

func (r PresetUpdateParamsConfigViewType) IsKnown() bool {
	switch r {
	case PresetUpdateParamsConfigViewTypeGroupCall, PresetUpdateParamsConfigViewTypeWebinar, PresetUpdateParamsConfigViewTypeAudioRoom:
		return true
	}
	return false
}

type PresetUpdateParamsPermissions struct {
	// Whether this participant can accept waiting requests
	AcceptWaitingRequests           param.Field[bool] `json:"accept_waiting_requests"`
	CanAcceptProductionRequests     param.Field[bool] `json:"can_accept_production_requests"`
	CanChangeParticipantPermissions param.Field[bool] `json:"can_change_participant_permissions"`
	CanEditDisplayName              param.Field[bool] `json:"can_edit_display_name"`
	CanLivestream                   param.Field[bool] `json:"can_livestream"`
	CanRecord                       param.Field[bool] `json:"can_record"`
	CanSpotlight                    param.Field[bool] `json:"can_spotlight"`
	// Chat permissions
	Chat                            param.Field[PresetUpdateParamsPermissionsChat]              `json:"chat"`
	ConnectedMeetings               param.Field[PresetUpdateParamsPermissionsConnectedMeetings] `json:"connected_meetings"`
	DisableParticipantAudio         param.Field[bool]                                           `json:"disable_participant_audio"`
	DisableParticipantScreensharing param.Field[bool]                                           `json:"disable_participant_screensharing"`
	DisableParticipantVideo         param.Field[bool]                                           `json:"disable_participant_video"`
	// Whether this participant is visible to others or not
	HiddenParticipant param.Field[bool] `json:"hidden_participant"`
	IsRecorder        param.Field[bool] `json:"is_recorder"`
	KickParticipant   param.Field[bool] `json:"kick_participant"`
	// Media permissions
	Media          param.Field[PresetUpdateParamsPermissionsMedia] `json:"media"`
	PinParticipant param.Field[bool]                               `json:"pin_participant"`
	// Plugin permissions
	Plugins param.Field[PresetUpdateParamsPermissionsPlugins] `json:"plugins"`
	// Poll permissions
	Polls param.Field[PresetUpdateParamsPermissionsPolls] `json:"polls"`
	// Type of the recording peer
	RecorderType        param.Field[PresetUpdateParamsPermissionsRecorderType] `json:"recorder_type"`
	ShowParticipantList param.Field[bool]                                      `json:"show_participant_list"`
	// Waiting room type
	WaitingRoomType param.Field[PresetUpdateParamsPermissionsWaitingRoomType] `json:"waiting_room_type"`
}

func (r PresetUpdateParamsPermissions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Chat permissions
type PresetUpdateParamsPermissionsChat struct {
	Private param.Field[PresetUpdateParamsPermissionsChatPrivate] `json:"private"`
	Public  param.Field[PresetUpdateParamsPermissionsChatPublic]  `json:"public"`
}

func (r PresetUpdateParamsPermissionsChat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetUpdateParamsPermissionsChatPrivate struct {
	CanReceive param.Field[bool] `json:"can_receive"`
	CanSend    param.Field[bool] `json:"can_send"`
	Files      param.Field[bool] `json:"files"`
	Text       param.Field[bool] `json:"text"`
}

func (r PresetUpdateParamsPermissionsChatPrivate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetUpdateParamsPermissionsChatPublic struct {
	// Can send messages in general
	CanSend param.Field[bool] `json:"can_send"`
	// Can send file messages
	Files param.Field[bool] `json:"files"`
	// Can send text messages
	Text param.Field[bool] `json:"text"`
}

func (r PresetUpdateParamsPermissionsChatPublic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetUpdateParamsPermissionsConnectedMeetings struct {
	CanAlterConnectedMeetings  param.Field[bool] `json:"can_alter_connected_meetings"`
	CanSwitchConnectedMeetings param.Field[bool] `json:"can_switch_connected_meetings"`
	CanSwitchToParentMeeting   param.Field[bool] `json:"can_switch_to_parent_meeting"`
}

func (r PresetUpdateParamsPermissionsConnectedMeetings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Media permissions
type PresetUpdateParamsPermissionsMedia struct {
	// Audio permissions
	Audio param.Field[PresetUpdateParamsPermissionsMediaAudio] `json:"audio"`
	// Screenshare permissions
	Screenshare param.Field[PresetUpdateParamsPermissionsMediaScreenshare] `json:"screenshare"`
	// Video permissions
	Video param.Field[PresetUpdateParamsPermissionsMediaVideo] `json:"video"`
}

func (r PresetUpdateParamsPermissionsMedia) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Audio permissions
type PresetUpdateParamsPermissionsMediaAudio struct {
	// Can produce audio
	CanProduce param.Field[PresetUpdateParamsPermissionsMediaAudioCanProduce] `json:"can_produce"`
}

func (r PresetUpdateParamsPermissionsMediaAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Can produce audio
type PresetUpdateParamsPermissionsMediaAudioCanProduce string

const (
	PresetUpdateParamsPermissionsMediaAudioCanProduceAllowed    PresetUpdateParamsPermissionsMediaAudioCanProduce = "ALLOWED"
	PresetUpdateParamsPermissionsMediaAudioCanProduceNotAllowed PresetUpdateParamsPermissionsMediaAudioCanProduce = "NOT_ALLOWED"
	PresetUpdateParamsPermissionsMediaAudioCanProduceCanRequest PresetUpdateParamsPermissionsMediaAudioCanProduce = "CAN_REQUEST"
)

func (r PresetUpdateParamsPermissionsMediaAudioCanProduce) IsKnown() bool {
	switch r {
	case PresetUpdateParamsPermissionsMediaAudioCanProduceAllowed, PresetUpdateParamsPermissionsMediaAudioCanProduceNotAllowed, PresetUpdateParamsPermissionsMediaAudioCanProduceCanRequest:
		return true
	}
	return false
}

// Screenshare permissions
type PresetUpdateParamsPermissionsMediaScreenshare struct {
	// Can produce screen share video
	CanProduce param.Field[PresetUpdateParamsPermissionsMediaScreenshareCanProduce] `json:"can_produce"`
}

func (r PresetUpdateParamsPermissionsMediaScreenshare) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Can produce screen share video
type PresetUpdateParamsPermissionsMediaScreenshareCanProduce string

const (
	PresetUpdateParamsPermissionsMediaScreenshareCanProduceAllowed    PresetUpdateParamsPermissionsMediaScreenshareCanProduce = "ALLOWED"
	PresetUpdateParamsPermissionsMediaScreenshareCanProduceNotAllowed PresetUpdateParamsPermissionsMediaScreenshareCanProduce = "NOT_ALLOWED"
	PresetUpdateParamsPermissionsMediaScreenshareCanProduceCanRequest PresetUpdateParamsPermissionsMediaScreenshareCanProduce = "CAN_REQUEST"
)

func (r PresetUpdateParamsPermissionsMediaScreenshareCanProduce) IsKnown() bool {
	switch r {
	case PresetUpdateParamsPermissionsMediaScreenshareCanProduceAllowed, PresetUpdateParamsPermissionsMediaScreenshareCanProduceNotAllowed, PresetUpdateParamsPermissionsMediaScreenshareCanProduceCanRequest:
		return true
	}
	return false
}

// Video permissions
type PresetUpdateParamsPermissionsMediaVideo struct {
	// Can produce video
	CanProduce param.Field[PresetUpdateParamsPermissionsMediaVideoCanProduce] `json:"can_produce"`
}

func (r PresetUpdateParamsPermissionsMediaVideo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Can produce video
type PresetUpdateParamsPermissionsMediaVideoCanProduce string

const (
	PresetUpdateParamsPermissionsMediaVideoCanProduceAllowed    PresetUpdateParamsPermissionsMediaVideoCanProduce = "ALLOWED"
	PresetUpdateParamsPermissionsMediaVideoCanProduceNotAllowed PresetUpdateParamsPermissionsMediaVideoCanProduce = "NOT_ALLOWED"
	PresetUpdateParamsPermissionsMediaVideoCanProduceCanRequest PresetUpdateParamsPermissionsMediaVideoCanProduce = "CAN_REQUEST"
)

func (r PresetUpdateParamsPermissionsMediaVideoCanProduce) IsKnown() bool {
	switch r {
	case PresetUpdateParamsPermissionsMediaVideoCanProduceAllowed, PresetUpdateParamsPermissionsMediaVideoCanProduceNotAllowed, PresetUpdateParamsPermissionsMediaVideoCanProduceCanRequest:
		return true
	}
	return false
}

// Plugin permissions
type PresetUpdateParamsPermissionsPlugins struct {
	// Can close plugins that are already open
	CanClose param.Field[bool] `json:"can_close"`
	// Can edit plugin config
	CanEditConfig param.Field[bool] `json:"can_edit_config"`
	// Can start plugins
	CanStart param.Field[bool]                                            `json:"can_start"`
	Config   param.Field[PresetUpdateParamsPermissionsPluginsConfigUnion] `json:"config" format:"uuid"`
}

func (r PresetUpdateParamsPermissionsPlugins) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [realtime_kit.PresetUpdateParamsPermissionsPluginsConfigObject].
type PresetUpdateParamsPermissionsPluginsConfigUnion interface {
	ImplementsPresetUpdateParamsPermissionsPluginsConfigUnion()
}

type PresetUpdateParamsPermissionsPluginsConfigObject struct {
	AccessControl   param.Field[PresetUpdateParamsPermissionsPluginsConfigObjectAccessControl] `json:"access_control"`
	HandlesViewOnly param.Field[bool]                                                          `json:"handles_view_only"`
}

func (r PresetUpdateParamsPermissionsPluginsConfigObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PresetUpdateParamsPermissionsPluginsConfigObject) ImplementsPresetUpdateParamsPermissionsPluginsConfigUnion() {
}

type PresetUpdateParamsPermissionsPluginsConfigObjectAccessControl string

const (
	PresetUpdateParamsPermissionsPluginsConfigObjectAccessControlFullAccess PresetUpdateParamsPermissionsPluginsConfigObjectAccessControl = "FULL_ACCESS"
	PresetUpdateParamsPermissionsPluginsConfigObjectAccessControlViewOnly   PresetUpdateParamsPermissionsPluginsConfigObjectAccessControl = "VIEW_ONLY"
)

func (r PresetUpdateParamsPermissionsPluginsConfigObjectAccessControl) IsKnown() bool {
	switch r {
	case PresetUpdateParamsPermissionsPluginsConfigObjectAccessControlFullAccess, PresetUpdateParamsPermissionsPluginsConfigObjectAccessControlViewOnly:
		return true
	}
	return false
}

// Poll permissions
type PresetUpdateParamsPermissionsPolls struct {
	// Can create polls
	CanCreate param.Field[bool] `json:"can_create"`
	// Can view polls
	CanView param.Field[bool] `json:"can_view"`
	// Can vote on polls
	CanVote param.Field[bool] `json:"can_vote"`
}

func (r PresetUpdateParamsPermissionsPolls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the recording peer
type PresetUpdateParamsPermissionsRecorderType string

const (
	PresetUpdateParamsPermissionsRecorderTypeRecorder     PresetUpdateParamsPermissionsRecorderType = "RECORDER"
	PresetUpdateParamsPermissionsRecorderTypeLivestreamer PresetUpdateParamsPermissionsRecorderType = "LIVESTREAMER"
	PresetUpdateParamsPermissionsRecorderTypeNone         PresetUpdateParamsPermissionsRecorderType = "NONE"
)

func (r PresetUpdateParamsPermissionsRecorderType) IsKnown() bool {
	switch r {
	case PresetUpdateParamsPermissionsRecorderTypeRecorder, PresetUpdateParamsPermissionsRecorderTypeLivestreamer, PresetUpdateParamsPermissionsRecorderTypeNone:
		return true
	}
	return false
}

// Waiting room type
type PresetUpdateParamsPermissionsWaitingRoomType string

const (
	PresetUpdateParamsPermissionsWaitingRoomTypeSkip                  PresetUpdateParamsPermissionsWaitingRoomType = "SKIP"
	PresetUpdateParamsPermissionsWaitingRoomTypeOnPrivilegedUserEntry PresetUpdateParamsPermissionsWaitingRoomType = "ON_PRIVILEGED_USER_ENTRY"
	PresetUpdateParamsPermissionsWaitingRoomTypeSkipOnAccept          PresetUpdateParamsPermissionsWaitingRoomType = "SKIP_ON_ACCEPT"
)

func (r PresetUpdateParamsPermissionsWaitingRoomType) IsKnown() bool {
	switch r {
	case PresetUpdateParamsPermissionsWaitingRoomTypeSkip, PresetUpdateParamsPermissionsWaitingRoomTypeOnPrivilegedUserEntry, PresetUpdateParamsPermissionsWaitingRoomTypeSkipOnAccept:
		return true
	}
	return false
}

type PresetUpdateParamsUI struct {
	ConfigDiff   param.Field[interface{}]                      `json:"config_diff"`
	DesignTokens param.Field[PresetUpdateParamsUIDesignTokens] `json:"design_tokens"`
}

func (r PresetUpdateParamsUI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetUpdateParamsUIDesignTokens struct {
	BorderRadius param.Field[PresetUpdateParamsUIDesignTokensBorderRadius] `json:"border_radius"`
	BorderWidth  param.Field[PresetUpdateParamsUIDesignTokensBorderWidth]  `json:"border_width"`
	Colors       param.Field[PresetUpdateParamsUIDesignTokensColors]       `json:"colors"`
	Logo         param.Field[string]                                       `json:"logo"`
	SpacingBase  param.Field[float64]                                      `json:"spacing_base"`
	Theme        param.Field[PresetUpdateParamsUIDesignTokensTheme]        `json:"theme"`
}

func (r PresetUpdateParamsUIDesignTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetUpdateParamsUIDesignTokensBorderRadius string

const (
	PresetUpdateParamsUIDesignTokensBorderRadiusRounded PresetUpdateParamsUIDesignTokensBorderRadius = "rounded"
)

func (r PresetUpdateParamsUIDesignTokensBorderRadius) IsKnown() bool {
	switch r {
	case PresetUpdateParamsUIDesignTokensBorderRadiusRounded:
		return true
	}
	return false
}

type PresetUpdateParamsUIDesignTokensBorderWidth string

const (
	PresetUpdateParamsUIDesignTokensBorderWidthThin PresetUpdateParamsUIDesignTokensBorderWidth = "thin"
)

func (r PresetUpdateParamsUIDesignTokensBorderWidth) IsKnown() bool {
	switch r {
	case PresetUpdateParamsUIDesignTokensBorderWidthThin:
		return true
	}
	return false
}

type PresetUpdateParamsUIDesignTokensColors struct {
	Background  param.Field[PresetUpdateParamsUIDesignTokensColorsBackground] `json:"background"`
	Brand       param.Field[PresetUpdateParamsUIDesignTokensColorsBrand]      `json:"brand"`
	Danger      param.Field[string]                                           `json:"danger"`
	Success     param.Field[string]                                           `json:"success"`
	Text        param.Field[string]                                           `json:"text"`
	TextOnBrand param.Field[string]                                           `json:"text_on_brand"`
	VideoBg     param.Field[string]                                           `json:"video_bg"`
	Warning     param.Field[string]                                           `json:"warning"`
}

func (r PresetUpdateParamsUIDesignTokensColors) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetUpdateParamsUIDesignTokensColorsBackground struct {
	Number1000 param.Field[string] `json:"1000"`
	Number600  param.Field[string] `json:"600"`
	Number700  param.Field[string] `json:"700"`
	Number800  param.Field[string] `json:"800"`
	Number900  param.Field[string] `json:"900"`
}

func (r PresetUpdateParamsUIDesignTokensColorsBackground) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetUpdateParamsUIDesignTokensColorsBrand struct {
	Number300 param.Field[string] `json:"300"`
	Number400 param.Field[string] `json:"400"`
	Number500 param.Field[string] `json:"500"`
	Number600 param.Field[string] `json:"600"`
	Number700 param.Field[string] `json:"700"`
}

func (r PresetUpdateParamsUIDesignTokensColorsBrand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PresetUpdateParamsUIDesignTokensTheme string

const (
	PresetUpdateParamsUIDesignTokensThemeDark PresetUpdateParamsUIDesignTokensTheme = "dark"
)

func (r PresetUpdateParamsUIDesignTokensTheme) IsKnown() bool {
	switch r {
	case PresetUpdateParamsUIDesignTokensThemeDark:
		return true
	}
	return false
}

type PresetDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PresetGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The page number from which you want your page search results to be displayed.
	PageNo param.Field[float64] `query:"page_no"`
	// Number of results per page
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [PresetGetParams]'s query parameters as `url.Values`.
func (r PresetGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PresetGetPresetByIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}
