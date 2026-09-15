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

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// SessionService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionService] method instead.
type SessionService struct {
	Options []option.RequestOption
}

// NewSessionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSessionService(opts ...option.RequestOption) (r *SessionService) {
	r = &SessionService{}
	r.Options = opts
	return
}

// Trigger Summary generation of Transcripts for the session ID.
func (r *SessionService) GenerateSummaryOfTranscripts(ctx context.Context, appID string, sessionID string, body SessionGenerateSummaryOfTranscriptsParams, opts ...option.RequestOption) (res *SessionGenerateSummaryOfTranscriptsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s/summary", body.AccountID, appID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns participant details for the given peer ID along with call statistics.
func (r *SessionService) GetParticipantDataFromPeerID(ctx context.Context, appID string, peerID string, params SessionGetParticipantDataFromPeerIDParams, opts ...option.RequestOption) (res *SessionGetParticipantDataFromPeerIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if peerID == "" {
		err = errors.New("missing required peer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/peer-report/%s", params.AccountID, appID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns a URL to download all chat messages of the session ID in CSV format.
func (r *SessionService) GetSessionChat(ctx context.Context, appID string, sessionID string, query SessionGetSessionChatParams, opts ...option.RequestOption) (res *SessionGetSessionChatResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s/chat", query.AccountID, appID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns data of the given session ID including recording details.
func (r *SessionService) GetSessionDetails(ctx context.Context, appID string, sessionID string, params SessionGetSessionDetailsParams, opts ...option.RequestOption) (res *SessionGetSessionDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s", params.AccountID, appID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns details of the given participant ID for the given session ID. Use the
// peer report endpoint to retrieve call statistics.
func (r *SessionService) GetSessionParticipantDetails(ctx context.Context, appID string, sessionID string, participantID string, params SessionGetSessionParticipantDetailsParams, opts ...option.RequestOption) (res *SessionGetSessionParticipantDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	if participantID == "" {
		err = errors.New("missing required participant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s/participants/%s", params.AccountID, appID, sessionID, participantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns a list of participants for the given session ID.
func (r *SessionService) GetSessionParticipants(ctx context.Context, appID string, sessionID string, params SessionGetSessionParticipantsParams, opts ...option.RequestOption) (res *SessionGetSessionParticipantsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s/participants", params.AccountID, appID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns a Summary URL to download the Summary of Transcripts for the session ID
// as plain text.
func (r *SessionService) GetSessionSummary(ctx context.Context, appID string, sessionID string, query SessionGetSessionSummaryParams, opts ...option.RequestOption) (res *SessionGetSessionSummaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s/summary", query.AccountID, appID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a URL to download the transcript for the session ID in CSV format.
func (r *SessionService) GetSessionTranscripts(ctx context.Context, appID string, sessionID string, params SessionGetSessionTranscriptsParams, opts ...option.RequestOption) (res *SessionGetSessionTranscriptsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s/transcript", params.AccountID, appID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns details of all sessions of an App.
func (r *SessionService) GetSessions(ctx context.Context, appID string, params SessionGetSessionsParams, opts ...option.RequestOption) (res *SessionGetSessionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type SessionGenerateSummaryOfTranscriptsResponse struct {
	Data    SessionGenerateSummaryOfTranscriptsResponseData `json:"data"`
	Success bool                                            `json:"success"`
	JSON    sessionGenerateSummaryOfTranscriptsResponseJSON `json:"-"`
}

// sessionGenerateSummaryOfTranscriptsResponseJSON contains the JSON metadata for
// the struct [SessionGenerateSummaryOfTranscriptsResponse]
type sessionGenerateSummaryOfTranscriptsResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGenerateSummaryOfTranscriptsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGenerateSummaryOfTranscriptsResponseJSON) RawJSON() string {
	return r.raw
}

type SessionGenerateSummaryOfTranscriptsResponseData struct {
	SessionID string                                              `json:"session_id" format:"uuid"`
	Status    string                                              `json:"status"`
	JSON      sessionGenerateSummaryOfTranscriptsResponseDataJSON `json:"-"`
}

// sessionGenerateSummaryOfTranscriptsResponseDataJSON contains the JSON metadata
// for the struct [SessionGenerateSummaryOfTranscriptsResponseData]
type sessionGenerateSummaryOfTranscriptsResponseDataJSON struct {
	SessionID   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGenerateSummaryOfTranscriptsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGenerateSummaryOfTranscriptsResponseDataJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponse struct {
	Data    SessionGetParticipantDataFromPeerIDResponseData `json:"data"`
	Success bool                                            `json:"success"`
	JSON    sessionGetParticipantDataFromPeerIDResponseJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseJSON contains the JSON metadata for
// the struct [SessionGetParticipantDataFromPeerIDResponse]
type sessionGetParticipantDataFromPeerIDResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseData struct {
	Participant SessionGetParticipantDataFromPeerIDResponseDataParticipant `json:"participant"`
	JSON        sessionGetParticipantDataFromPeerIDResponseDataJSON        `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataJSON contains the JSON metadata
// for the struct [SessionGetParticipantDataFromPeerIDResponseData]
type sessionGetParticipantDataFromPeerIDResponseDataJSON struct {
	Participant apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipant struct {
	// ID of the participant.
	ID string `json:"id" format:"uuid"`
	// timestamp when this participant was created.
	CreatedAt string `json:"created_at"`
	// ID passed by client to create this participant.
	CustomParticipantID string `json:"custom_participant_id"`
	// Display name of participant when joining the session.
	DisplayName string `json:"display_name"`
	// number of minutes for which the participant was in the session.
	Duration float64 `json:"duration"`
	// timestamp at which participant joined the session.
	JoinedAt string `json:"joined_at"`
	// timestamp at which participant left the session.
	LeftAt string `json:"left_at"`
	// Connection lifecycle events for the participant's peer.
	PeerEvents []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEvent `json:"peer_events"`
	// Peer call statistics report.
	PeerReport SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReport `json:"peer_report"`
	// Name of the preset associated with the participant.
	Role      string `json:"role"`
	SessionID string `json:"session_id" format:"uuid"`
	// timestamp when this participant's data was last updated.
	UpdatedAt string `json:"updated_at"`
	// User id for this participant.
	UserID string                                                         `json:"user_id"`
	JSON   sessionGetParticipantDataFromPeerIDResponseDataParticipantJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantJSON contains the JSON
// metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipant]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	CustomParticipantID apijson.Field
	DisplayName         apijson.Field
	Duration            apijson.Field
	JoinedAt            apijson.Field
	LeftAt              apijson.Field
	PeerEvents          apijson.Field
	PeerReport          apijson.Field
	Role                apijson.Field
	SessionID           apijson.Field
	UpdatedAt           apijson.Field
	UserID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantJSON) RawJSON() string {
	return r.raw
}

// A connection lifecycle event recorded for a participant's peer.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEvent struct {
	// ID of the peer event.
	ID string `json:"id"`
	// Timestamp when this peer event was created.
	CreatedAt string `json:"created_at"`
	// Name of the peer event.
	EventName SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsEventName `json:"event_name"`
	// Minutes consumed attributed to this event.
	MinutesConsumed float64 `json:"minutes_consumed"`
	// ID of the participant this event belongs to.
	ParticipantID string `json:"participant_id" api:"nullable"`
	// Peer ID this event belongs to.
	PeerID string `json:"peer_id"`
	// View type of the preset associated with the peer.
	PresetViewType SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewType `json:"preset_view_type" api:"nullable"`
	// ID of the session this event belongs to.
	SessionID string `json:"session_id" api:"nullable"`
	// ID of the socket session associated with this event.
	SocketSessionID string `json:"socket_session_id" api:"nullable"`
	// Timestamp when this peer event was last updated.
	UpdatedAt string                                                                  `json:"updated_at"`
	JSON      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventJSON contains
// the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEvent]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	EventName       apijson.Field
	MinutesConsumed apijson.Field
	ParticipantID   apijson.Field
	PeerID          apijson.Field
	PresetViewType  apijson.Field
	SessionID       apijson.Field
	SocketSessionID apijson.Field
	UpdatedAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventJSON) RawJSON() string {
	return r.raw
}

// Name of the peer event.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsEventName string

const (
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsEventNamePeerCreated SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsEventName = "PEER_CREATED"
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsEventNamePeerJoining SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsEventName = "PEER_JOINING"
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsEventNamePeerLeaving SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsEventName = "PEER_LEAVING"
)

func (r SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsEventName) IsKnown() bool {
	switch r {
	case SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsEventNamePeerCreated, SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsEventNamePeerJoining, SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsEventNamePeerLeaving:
		return true
	}
	return false
}

// View type of the preset associated with the peer.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewType string

const (
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewTypeGroupCall  SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewType = "GROUP_CALL"
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewTypeWebinar    SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewType = "WEBINAR"
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewTypeAudioRoom  SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewType = "AUDIO_ROOM"
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewTypeLivestream SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewType = "LIVESTREAM"
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewTypeChat       SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewType = "CHAT"
)

func (r SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewType) IsKnown() bool {
	switch r {
	case SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewTypeGroupCall, SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewTypeWebinar, SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewTypeAudioRoom, SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewTypeLivestream, SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerEventsPresetViewTypeChat:
		return true
	}
	return false
}

// Peer call statistics report.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReport struct {
	// Connection and device metadata for the participant.
	Metadata SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadata `json:"metadata"`
	// Media quality statistics for the participant.
	Quality SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQuality `json:"quality"`
	JSON    sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportJSON    `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReport]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportJSON struct {
	Metadata    apijson.Field
	Quality     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportJSON) RawJSON() string {
	return r.raw
}

// Connection and device metadata for the participant.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadata struct {
	AudioDevicesUpdates   []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdate   `json:"audio_devices_updates"`
	BrowserMetadata       SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataBrowserMetadata        `json:"browser_metadata"`
	CandidatePairs        SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairs         `json:"candidate_pairs"`
	DeviceInfo            SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataDeviceInfo             `json:"device_info"`
	Events                []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEvent                `json:"events"`
	IPInformation         SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformation          `json:"ip_information"`
	NativeMetadata        SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataNativeMetadata         `json:"native_metadata"`
	PcMetadata            []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataPcMetadata           `json:"pc_metadata"`
	RoomViewType          string                                                                                             `json:"room_view_type"`
	SDKName               string                                                                                             `json:"sdk_name"`
	SDKType               string                                                                                             `json:"sdk_type"`
	SDKVersion            string                                                                                             `json:"sdk_version"`
	SelectedDeviceUpdates []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdate `json:"selected_device_updates"`
	SpeakerDevicesUpdates []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdate `json:"speaker_devices_updates"`
	VideoDevicesUpdates   []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdate   `json:"video_devices_updates"`
	JSON                  sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataJSON                   `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadata]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataJSON struct {
	AudioDevicesUpdates   apijson.Field
	BrowserMetadata       apijson.Field
	CandidatePairs        apijson.Field
	DeviceInfo            apijson.Field
	Events                apijson.Field
	IPInformation         apijson.Field
	NativeMetadata        apijson.Field
	PcMetadata            apijson.Field
	RoomViewType          apijson.Field
	SDKName               apijson.Field
	SDKType               apijson.Field
	SDKVersion            apijson.Field
	SelectedDeviceUpdates apijson.Field
	SpeakerDevicesUpdates apijson.Field
	VideoDevicesUpdates   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataJSON) RawJSON() string {
	return r.raw
}

// A change to the set of available devices at a point in time.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdate struct {
	// Devices that became available.
	Added []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesAdded `json:"added"`
	// Devices that became unavailable.
	Removed []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesRemoved `json:"removed"`
	// Timestamp of the device update.
	Timestamp string                                                                                             `json:"timestamp"`
	JSON      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdateJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdateJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdate]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdateJSON struct {
	Added       apijson.Field
	Removed     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdateJSON) RawJSON() string {
	return r.raw
}

// A media device (camera, microphone, or speaker).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesAdded struct {
	// ID of the device.
	DeviceID string `json:"device_id"`
	// Kind of device, for example audioinput or videoinput.
	Kind string `json:"kind"`
	// Human-readable label of the device.
	Label string                                                                                                   `json:"label"`
	JSON  sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesAddedJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesAddedJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesAdded]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesAddedJSON struct {
	DeviceID    apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesAdded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesAddedJSON) RawJSON() string {
	return r.raw
}

// A media device (camera, microphone, or speaker).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesRemoved struct {
	// ID of the device.
	DeviceID string `json:"device_id"`
	// Kind of device, for example audioinput or videoinput.
	Kind string `json:"kind"`
	// Human-readable label of the device.
	Label string                                                                                                     `json:"label"`
	JSON  sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesRemovedJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesRemovedJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesRemoved]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesRemovedJSON struct {
	DeviceID    apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesRemoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataAudioDevicesUpdatesRemovedJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataBrowserMetadata struct {
	Browser        string                                                                                          `json:"browser"`
	BrowserVersion string                                                                                          `json:"browser_version"`
	Engine         string                                                                                          `json:"engine"`
	UserAgent      string                                                                                          `json:"user_agent"`
	WebglSupport   bool                                                                                            `json:"webgl_support" api:"nullable"`
	JSON           sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataBrowserMetadataJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataBrowserMetadataJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataBrowserMetadata]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataBrowserMetadataJSON struct {
	Browser        apijson.Field
	BrowserVersion apijson.Field
	Engine         apijson.Field
	UserAgent      apijson.Field
	WebglSupport   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataBrowserMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataBrowserMetadataJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairs struct {
	ConsumingTransport []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsConsumingTransport `json:"consuming_transport"`
	ProducingTransport []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsProducingTransport `json:"producing_transport"`
	JSON               sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsJSON                 `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairs]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsJSON struct {
	ConsumingTransport apijson.Field
	ProducingTransport apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsJSON) RawJSON() string {
	return r.raw
}

// ICE candidate pair statistics for a transport.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsConsumingTransport struct {
	AvailableIncomingBitrate float64 `json:"available_incoming_bitrate"`
	AvailableOutgoingBitrate float64 `json:"available_outgoing_bitrate"`
	BytesDiscardedOnSend     float64 `json:"bytes_discarded_on_send"`
	BytesReceived            float64 `json:"bytes_received"`
	BytesSent                float64 `json:"bytes_sent"`
	CurrentRoundTripTime     float64 `json:"current_round_trip_time"`
	// Epoch milliseconds when the last packet was received.
	LastPacketReceivedTimestamp float64 `json:"last_packet_received_timestamp"`
	// Epoch milliseconds when the last packet was sent.
	LastPacketSentTimestamp      float64                                                                                                          `json:"last_packet_sent_timestamp"`
	LocalCandidateAddress        string                                                                                                           `json:"local_candidate_address"`
	LocalCandidateID             string                                                                                                           `json:"local_candidate_id"`
	LocalCandidateNetworkType    string                                                                                                           `json:"local_candidate_network_type"`
	LocalCandidatePort           float64                                                                                                          `json:"local_candidate_port"`
	LocalCandidateProtocol       string                                                                                                           `json:"local_candidate_protocol"`
	LocalCandidateRelatedAddress string                                                                                                           `json:"local_candidate_related_address"`
	LocalCandidateRelatedPort    float64                                                                                                          `json:"local_candidate_related_port"`
	LocalCandidateType           string                                                                                                           `json:"local_candidate_type"`
	LocalCandidateURL            string                                                                                                           `json:"local_candidate_url"`
	Nominated                    bool                                                                                                             `json:"nominated"`
	PacketsDiscardedOnSend       float64                                                                                                          `json:"packets_discarded_on_send"`
	PacketsReceived              float64                                                                                                          `json:"packets_received"`
	PacketsSent                  float64                                                                                                          `json:"packets_sent"`
	RemoteCandidateAddress       string                                                                                                           `json:"remote_candidate_address"`
	RemoteCandidateID            string                                                                                                           `json:"remote_candidate_id"`
	RemoteCandidatePort          float64                                                                                                          `json:"remote_candidate_port"`
	RemoteCandidateProtocol      string                                                                                                           `json:"remote_candidate_protocol"`
	RemoteCandidateType          string                                                                                                           `json:"remote_candidate_type"`
	RemoteCandidateURL           string                                                                                                           `json:"remote_candidate_url"`
	TotalRoundTripTime           float64                                                                                                          `json:"total_round_trip_time"`
	JSON                         sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsConsumingTransportJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsConsumingTransportJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsConsumingTransport]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsConsumingTransportJSON struct {
	AvailableIncomingBitrate     apijson.Field
	AvailableOutgoingBitrate     apijson.Field
	BytesDiscardedOnSend         apijson.Field
	BytesReceived                apijson.Field
	BytesSent                    apijson.Field
	CurrentRoundTripTime         apijson.Field
	LastPacketReceivedTimestamp  apijson.Field
	LastPacketSentTimestamp      apijson.Field
	LocalCandidateAddress        apijson.Field
	LocalCandidateID             apijson.Field
	LocalCandidateNetworkType    apijson.Field
	LocalCandidatePort           apijson.Field
	LocalCandidateProtocol       apijson.Field
	LocalCandidateRelatedAddress apijson.Field
	LocalCandidateRelatedPort    apijson.Field
	LocalCandidateType           apijson.Field
	LocalCandidateURL            apijson.Field
	Nominated                    apijson.Field
	PacketsDiscardedOnSend       apijson.Field
	PacketsReceived              apijson.Field
	PacketsSent                  apijson.Field
	RemoteCandidateAddress       apijson.Field
	RemoteCandidateID            apijson.Field
	RemoteCandidatePort          apijson.Field
	RemoteCandidateProtocol      apijson.Field
	RemoteCandidateType          apijson.Field
	RemoteCandidateURL           apijson.Field
	TotalRoundTripTime           apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsConsumingTransport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsConsumingTransportJSON) RawJSON() string {
	return r.raw
}

// ICE candidate pair statistics for a transport.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsProducingTransport struct {
	AvailableIncomingBitrate float64 `json:"available_incoming_bitrate"`
	AvailableOutgoingBitrate float64 `json:"available_outgoing_bitrate"`
	BytesDiscardedOnSend     float64 `json:"bytes_discarded_on_send"`
	BytesReceived            float64 `json:"bytes_received"`
	BytesSent                float64 `json:"bytes_sent"`
	CurrentRoundTripTime     float64 `json:"current_round_trip_time"`
	// Epoch milliseconds when the last packet was received.
	LastPacketReceivedTimestamp float64 `json:"last_packet_received_timestamp"`
	// Epoch milliseconds when the last packet was sent.
	LastPacketSentTimestamp      float64                                                                                                          `json:"last_packet_sent_timestamp"`
	LocalCandidateAddress        string                                                                                                           `json:"local_candidate_address"`
	LocalCandidateID             string                                                                                                           `json:"local_candidate_id"`
	LocalCandidateNetworkType    string                                                                                                           `json:"local_candidate_network_type"`
	LocalCandidatePort           float64                                                                                                          `json:"local_candidate_port"`
	LocalCandidateProtocol       string                                                                                                           `json:"local_candidate_protocol"`
	LocalCandidateRelatedAddress string                                                                                                           `json:"local_candidate_related_address"`
	LocalCandidateRelatedPort    float64                                                                                                          `json:"local_candidate_related_port"`
	LocalCandidateType           string                                                                                                           `json:"local_candidate_type"`
	LocalCandidateURL            string                                                                                                           `json:"local_candidate_url"`
	Nominated                    bool                                                                                                             `json:"nominated"`
	PacketsDiscardedOnSend       float64                                                                                                          `json:"packets_discarded_on_send"`
	PacketsReceived              float64                                                                                                          `json:"packets_received"`
	PacketsSent                  float64                                                                                                          `json:"packets_sent"`
	RemoteCandidateAddress       string                                                                                                           `json:"remote_candidate_address"`
	RemoteCandidateID            string                                                                                                           `json:"remote_candidate_id"`
	RemoteCandidatePort          float64                                                                                                          `json:"remote_candidate_port"`
	RemoteCandidateProtocol      string                                                                                                           `json:"remote_candidate_protocol"`
	RemoteCandidateType          string                                                                                                           `json:"remote_candidate_type"`
	RemoteCandidateURL           string                                                                                                           `json:"remote_candidate_url"`
	TotalRoundTripTime           float64                                                                                                          `json:"total_round_trip_time"`
	JSON                         sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsProducingTransportJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsProducingTransportJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsProducingTransport]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsProducingTransportJSON struct {
	AvailableIncomingBitrate     apijson.Field
	AvailableOutgoingBitrate     apijson.Field
	BytesDiscardedOnSend         apijson.Field
	BytesReceived                apijson.Field
	BytesSent                    apijson.Field
	CurrentRoundTripTime         apijson.Field
	LastPacketReceivedTimestamp  apijson.Field
	LastPacketSentTimestamp      apijson.Field
	LocalCandidateAddress        apijson.Field
	LocalCandidateID             apijson.Field
	LocalCandidateNetworkType    apijson.Field
	LocalCandidatePort           apijson.Field
	LocalCandidateProtocol       apijson.Field
	LocalCandidateRelatedAddress apijson.Field
	LocalCandidateRelatedPort    apijson.Field
	LocalCandidateType           apijson.Field
	LocalCandidateURL            apijson.Field
	Nominated                    apijson.Field
	PacketsDiscardedOnSend       apijson.Field
	PacketsReceived              apijson.Field
	PacketsSent                  apijson.Field
	RemoteCandidateAddress       apijson.Field
	RemoteCandidateID            apijson.Field
	RemoteCandidatePort          apijson.Field
	RemoteCandidateProtocol      apijson.Field
	RemoteCandidateType          apijson.Field
	RemoteCandidateURL           apijson.Field
	TotalRoundTripTime           apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsProducingTransport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsProducingTransportJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataDeviceInfo struct {
	CPUs      float64                                                                                    `json:"cpus"`
	IsMobile  bool                                                                                       `json:"is_mobile"`
	OS        string                                                                                     `json:"os"`
	OSVersion string                                                                                     `json:"os_version"`
	JSON      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataDeviceInfoJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataDeviceInfoJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataDeviceInfo]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataDeviceInfoJSON struct {
	CPUs        apijson.Field
	IsMobile    apijson.Field
	OS          apijson.Field
	OSVersion   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataDeviceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataDeviceInfoJSON) RawJSON() string {
	return r.raw
}

// A timestamped event recorded during the participant's connection.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEvent struct {
	// Event-specific metadata. Keys vary per event; values are primitive scalars
	// (string, number, boolean, or null).
	Metadata map[string]SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventsMetadataUnion `json:"metadata"`
	// Name of the event.
	Name string `json:"name"`
	// Timestamp when the event occurred.
	Timestamp string                                                                                `json:"timestamp"`
	JSON      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEvent]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventJSON struct {
	Metadata    apijson.Field
	Name        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventsMetadataUnion interface {
	ImplementsSessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventsMetadataUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventsMetadataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformation struct {
	ASN      SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASN  `json:"asn"`
	City     string                                                                                        `json:"city"`
	Country  string                                                                                        `json:"country"`
	IPV4     string                                                                                        `json:"ipv4"`
	Org      string                                                                                        `json:"org"`
	Region   string                                                                                        `json:"region"`
	Timezone string                                                                                        `json:"timezone"`
	JSON     sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformation]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationJSON struct {
	ASN         apijson.Field
	City        apijson.Field
	Country     apijson.Field
	IPV4        apijson.Field
	Org         apijson.Field
	Region      apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASN struct {
	ASN    string                                                                                           `json:"asn"`
	Domain string                                                                                           `json:"domain"`
	Name   string                                                                                           `json:"name"`
	Route  string                                                                                           `json:"route"`
	Type   string                                                                                           `json:"type"`
	JSON   sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASNJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASNJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASN]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASNJSON struct {
	ASN         apijson.Field
	Domain      apijson.Field
	Name        apijson.Field
	Route       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASNJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataNativeMetadata struct {
	AudioEncoder string                                                                                         `json:"audio_encoder"`
	VideoEncoder string                                                                                         `json:"video_encoder"`
	JSON         sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataNativeMetadataJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataNativeMetadataJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataNativeMetadata]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataNativeMetadataJSON struct {
	AudioEncoder apijson.Field
	VideoEncoder apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataNativeMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataNativeMetadataJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataPcMetadata struct {
	EffectiveNetworkType  string                                                                                     `json:"effective_network_type"`
	ReflexiveConnectivity bool                                                                                       `json:"reflexive_connectivity"`
	RelayConnectivity     bool                                                                                       `json:"relay_connectivity"`
	Sdp                   []string                                                                                   `json:"sdp"`
	Timestamp             string                                                                                     `json:"timestamp"`
	TURNConnectivity      bool                                                                                       `json:"turn_connectivity"`
	JSON                  sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataPcMetadataJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataPcMetadataJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataPcMetadata]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataPcMetadataJSON struct {
	EffectiveNetworkType  apijson.Field
	ReflexiveConnectivity apijson.Field
	RelayConnectivity     apijson.Field
	Sdp                   apijson.Field
	Timestamp             apijson.Field
	TURNConnectivity      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataPcMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataPcMetadataJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdate struct {
	// A media device (camera, microphone, or speaker).
	Device    SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdatesDevice `json:"device"`
	Timestamp string                                                                                                  `json:"timestamp"`
	JSON      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdateJSON    `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdateJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdate]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdateJSON struct {
	Device      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdateJSON) RawJSON() string {
	return r.raw
}

// A media device (camera, microphone, or speaker).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdatesDevice struct {
	// ID of the device.
	DeviceID string `json:"device_id"`
	// Kind of device, for example audioinput or videoinput.
	Kind string `json:"kind"`
	// Human-readable label of the device.
	Label string                                                                                                      `json:"label"`
	JSON  sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdatesDeviceJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdatesDeviceJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdatesDevice]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdatesDeviceJSON struct {
	DeviceID    apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdatesDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSelectedDeviceUpdatesDeviceJSON) RawJSON() string {
	return r.raw
}

// A change to the set of available devices at a point in time.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdate struct {
	// Devices that became available.
	Added []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesAdded `json:"added"`
	// Devices that became unavailable.
	Removed []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesRemoved `json:"removed"`
	// Timestamp of the device update.
	Timestamp string                                                                                               `json:"timestamp"`
	JSON      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdateJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdateJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdate]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdateJSON struct {
	Added       apijson.Field
	Removed     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdateJSON) RawJSON() string {
	return r.raw
}

// A media device (camera, microphone, or speaker).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesAdded struct {
	// ID of the device.
	DeviceID string `json:"device_id"`
	// Kind of device, for example audioinput or videoinput.
	Kind string `json:"kind"`
	// Human-readable label of the device.
	Label string                                                                                                     `json:"label"`
	JSON  sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesAddedJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesAddedJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesAdded]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesAddedJSON struct {
	DeviceID    apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesAdded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesAddedJSON) RawJSON() string {
	return r.raw
}

// A media device (camera, microphone, or speaker).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesRemoved struct {
	// ID of the device.
	DeviceID string `json:"device_id"`
	// Kind of device, for example audioinput or videoinput.
	Kind string `json:"kind"`
	// Human-readable label of the device.
	Label string                                                                                                       `json:"label"`
	JSON  sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesRemovedJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesRemovedJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesRemoved]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesRemovedJSON struct {
	DeviceID    apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesRemoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataSpeakerDevicesUpdatesRemovedJSON) RawJSON() string {
	return r.raw
}

// A change to the set of available devices at a point in time.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdate struct {
	// Devices that became available.
	Added []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesAdded `json:"added"`
	// Devices that became unavailable.
	Removed []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesRemoved `json:"removed"`
	// Timestamp of the device update.
	Timestamp string                                                                                             `json:"timestamp"`
	JSON      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdateJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdateJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdate]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdateJSON struct {
	Added       apijson.Field
	Removed     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdateJSON) RawJSON() string {
	return r.raw
}

// A media device (camera, microphone, or speaker).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesAdded struct {
	// ID of the device.
	DeviceID string `json:"device_id"`
	// Kind of device, for example audioinput or videoinput.
	Kind string `json:"kind"`
	// Human-readable label of the device.
	Label string                                                                                                   `json:"label"`
	JSON  sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesAddedJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesAddedJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesAdded]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesAddedJSON struct {
	DeviceID    apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesAdded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesAddedJSON) RawJSON() string {
	return r.raw
}

// A media device (camera, microphone, or speaker).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesRemoved struct {
	// ID of the device.
	DeviceID string `json:"device_id"`
	// Kind of device, for example audioinput or videoinput.
	Kind string `json:"kind"`
	// Human-readable label of the device.
	Label string                                                                                                     `json:"label"`
	JSON  sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesRemovedJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesRemovedJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesRemoved]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesRemovedJSON struct {
	DeviceID    apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesRemoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataVideoDevicesUpdatesRemovedJSON) RawJSON() string {
	return r.raw
}

// Media quality statistics for the participant.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQuality struct {
	AudioConsumer []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumer `json:"audio_consumer"`
	// Aggregated inbound (consumer) audio statistics for the session.
	AudioConsumerCumulative SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulative `json:"audio_consumer_cumulative"`
	AudioProducer           []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducer         `json:"audio_producer"`
	// Aggregated outbound (producer) audio statistics for the session.
	AudioProducerCumulative  SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulative    `json:"audio_producer_cumulative"`
	ScreenshareAudioConsumer []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumer `json:"screenshare_audio_consumer"`
	// Aggregated inbound (consumer) audio statistics for the session.
	ScreenshareAudioConsumerCumulative SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulative `json:"screenshare_audio_consumer_cumulative"`
	ScreenshareAudioProducer           []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducer         `json:"screenshare_audio_producer"`
	// Aggregated outbound (producer) audio statistics for the session.
	ScreenshareAudioProducerCumulative SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulative `json:"screenshare_audio_producer_cumulative"`
	ScreenshareVideoConsumer           []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumer         `json:"screenshare_video_consumer"`
	// Aggregated inbound (consumer) video statistics for the session.
	ScreenshareVideoConsumerCumulative SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulative `json:"screenshare_video_consumer_cumulative"`
	ScreenshareVideoProducer           []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducer         `json:"screenshare_video_producer"`
	// Aggregated outbound (producer) video statistics for the session.
	ScreenshareVideoProducerCumulative SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulative `json:"screenshare_video_producer_cumulative"`
	VideoConsumer                      []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumer                    `json:"video_consumer"`
	// Aggregated inbound (consumer) video statistics for the session.
	VideoConsumerCumulative SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulative `json:"video_consumer_cumulative"`
	VideoProducer           []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducer         `json:"video_producer"`
	// Aggregated outbound (producer) video statistics for the session.
	VideoProducerCumulative SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulative `json:"video_producer_cumulative"`
	JSON                    sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityJSON                    `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQuality]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityJSON struct {
	AudioConsumer                      apijson.Field
	AudioConsumerCumulative            apijson.Field
	AudioProducer                      apijson.Field
	AudioProducerCumulative            apijson.Field
	ScreenshareAudioConsumer           apijson.Field
	ScreenshareAudioConsumerCumulative apijson.Field
	ScreenshareAudioProducer           apijson.Field
	ScreenshareAudioProducerCumulative apijson.Field
	ScreenshareVideoConsumer           apijson.Field
	ScreenshareVideoConsumerCumulative apijson.Field
	ScreenshareVideoProducer           apijson.Field
	ScreenshareVideoProducerCumulative apijson.Field
	VideoConsumer                      apijson.Field
	VideoConsumerCumulative            apijson.Field
	VideoProducer                      apijson.Field
	VideoProducerCumulative            apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQuality) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityJSON) RawJSON() string {
	return r.raw
}

// Per-sample inbound (consumer) audio statistics.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumer struct {
	BytesReceived            float64                                                                                      `json:"bytes_received"`
	ConcealmentEvents        float64                                                                                      `json:"concealment_events"`
	ConsumerID               string                                                                                       `json:"consumer_id"`
	Jitter                   float64                                                                                      `json:"jitter"`
	JitterBufferDelay        float64                                                                                      `json:"jitter_buffer_delay"`
	JitterBufferEmittedCount float64                                                                                      `json:"jitter_buffer_emitted_count"`
	Mid                      string                                                                                       `json:"mid"`
	MosQuality               float64                                                                                      `json:"mos_quality"`
	PacketsLost              float64                                                                                      `json:"packets_lost"`
	PacketsReceived          float64                                                                                      `json:"packets_received"`
	PeerID                   string                                                                                       `json:"peer_id"`
	ProducerID               string                                                                                       `json:"producer_id"`
	Ssrc                     float64                                                                                      `json:"ssrc"`
	Timestamp                string                                                                                       `json:"timestamp"`
	JSON                     sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumer]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerJSON struct {
	BytesReceived            apijson.Field
	ConcealmentEvents        apijson.Field
	ConsumerID               apijson.Field
	Jitter                   apijson.Field
	JitterBufferDelay        apijson.Field
	JitterBufferEmittedCount apijson.Field
	Mid                      apijson.Field
	MosQuality               apijson.Field
	PacketsLost              apijson.Field
	PacketsReceived          apijson.Field
	PeerID                   apijson.Field
	ProducerID               apijson.Field
	Ssrc                     apijson.Field
	Timestamp                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerJSON) RawJSON() string {
	return r.raw
}

// Aggregated inbound (consumer) audio statistics for the session.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulative struct {
	// Cumulative latency distribution (milliseconds-based thresholds).
	JitterBufferDelay SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeJitterBufferDelay `json:"jitter_buffer_delay"`
	// Cumulative packet loss distribution.
	PacketLoss SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativePacketLoss `json:"packet_loss"`
	// Distribution summary with average and percentiles.
	QualityMos SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeQualityMos `json:"quality_mos"`
	JSON       sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeJSON       `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulative]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeJSON struct {
	JitterBufferDelay apijson.Field
	PacketLoss        apijson.Field
	QualityMos        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeJSON) RawJSON() string {
	return r.raw
}

// Cumulative latency distribution (milliseconds-based thresholds).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeJitterBufferDelay struct {
	Number100msOrGreaterEventFraction float64                                                                                                                 `json:"100ms_or_greater_event_fraction"`
	Number250msOrGreaterEventFraction float64                                                                                                                 `json:"250ms_or_greater_event_fraction"`
	Number500msOrGreaterEventFraction float64                                                                                                                 `json:"500ms_or_greater_event_fraction"`
	Avg                               float64                                                                                                                 `json:"avg"`
	JSON                              sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeJitterBufferDelayJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeJitterBufferDelayJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeJitterBufferDelay]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeJitterBufferDelayJSON struct {
	Number100msOrGreaterEventFraction apijson.Field
	Number250msOrGreaterEventFraction apijson.Field
	Number500msOrGreaterEventFraction apijson.Field
	Avg                               apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeJitterBufferDelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeJitterBufferDelayJSON) RawJSON() string {
	return r.raw
}

// Cumulative packet loss distribution.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativePacketLoss struct {
	Number10OrGreaterEventFraction float64                                                                                                          `json:"10_or_greater_event_fraction"`
	Number25OrGreaterEventFraction float64                                                                                                          `json:"25_or_greater_event_fraction"`
	Number5OrGreaterEventFraction  float64                                                                                                          `json:"5_or_greater_event_fraction"`
	Number50OrGreaterEventFraction float64                                                                                                          `json:"50_or_greater_event_fraction"`
	Avg                            float64                                                                                                          `json:"avg"`
	JSON                           sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativePacketLossJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativePacketLossJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativePacketLoss]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativePacketLossJSON struct {
	Number10OrGreaterEventFraction apijson.Field
	Number25OrGreaterEventFraction apijson.Field
	Number5OrGreaterEventFraction  apijson.Field
	Number50OrGreaterEventFraction apijson.Field
	Avg                            apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativePacketLoss) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativePacketLossJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeQualityMos struct {
	Avg  float64                                                                                                          `json:"avg"`
	P50  float64                                                                                                          `json:"p50"`
	P75  float64                                                                                                          `json:"p75"`
	P90  float64                                                                                                          `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeQualityMosJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeQualityMosJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeQualityMos]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeQualityMosJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeQualityMos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioConsumerCumulativeQualityMosJSON) RawJSON() string {
	return r.raw
}

// Per-sample outbound (producer) audio statistics.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducer struct {
	BytesSent   float64                                                                                      `json:"bytes_sent"`
	Jitter      float64                                                                                      `json:"jitter"`
	Mid         string                                                                                       `json:"mid"`
	MosQuality  float64                                                                                      `json:"mos_quality"`
	PacketsLost float64                                                                                      `json:"packets_lost"`
	PacketsSent float64                                                                                      `json:"packets_sent"`
	ProducerID  string                                                                                       `json:"producer_id"`
	RTT         float64                                                                                      `json:"rtt"`
	Ssrc        float64                                                                                      `json:"ssrc"`
	Timestamp   string                                                                                       `json:"timestamp"`
	JSON        sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducer]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerJSON struct {
	BytesSent   apijson.Field
	Jitter      apijson.Field
	Mid         apijson.Field
	MosQuality  apijson.Field
	PacketsLost apijson.Field
	PacketsSent apijson.Field
	ProducerID  apijson.Field
	RTT         apijson.Field
	Ssrc        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerJSON) RawJSON() string {
	return r.raw
}

// Aggregated outbound (producer) audio statistics for the session.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulative struct {
	// Cumulative packet loss distribution.
	PacketLoss SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativePacketLoss `json:"packet_loss"`
	// Distribution summary with average and percentiles.
	QualityMos SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeQualityMos `json:"quality_mos"`
	// Cumulative latency distribution (milliseconds-based thresholds).
	RTT  SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeRTT  `json:"rtt"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulative]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeJSON struct {
	PacketLoss  apijson.Field
	QualityMos  apijson.Field
	RTT         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeJSON) RawJSON() string {
	return r.raw
}

// Cumulative packet loss distribution.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativePacketLoss struct {
	Number10OrGreaterEventFraction float64                                                                                                          `json:"10_or_greater_event_fraction"`
	Number25OrGreaterEventFraction float64                                                                                                          `json:"25_or_greater_event_fraction"`
	Number5OrGreaterEventFraction  float64                                                                                                          `json:"5_or_greater_event_fraction"`
	Number50OrGreaterEventFraction float64                                                                                                          `json:"50_or_greater_event_fraction"`
	Avg                            float64                                                                                                          `json:"avg"`
	JSON                           sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativePacketLossJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativePacketLossJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativePacketLoss]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativePacketLossJSON struct {
	Number10OrGreaterEventFraction apijson.Field
	Number25OrGreaterEventFraction apijson.Field
	Number5OrGreaterEventFraction  apijson.Field
	Number50OrGreaterEventFraction apijson.Field
	Avg                            apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativePacketLoss) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativePacketLossJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeQualityMos struct {
	Avg  float64                                                                                                          `json:"avg"`
	P50  float64                                                                                                          `json:"p50"`
	P75  float64                                                                                                          `json:"p75"`
	P90  float64                                                                                                          `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeQualityMosJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeQualityMosJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeQualityMos]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeQualityMosJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeQualityMos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeQualityMosJSON) RawJSON() string {
	return r.raw
}

// Cumulative latency distribution (milliseconds-based thresholds).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeRTT struct {
	Number100msOrGreaterEventFraction float64                                                                                                   `json:"100ms_or_greater_event_fraction"`
	Number250msOrGreaterEventFraction float64                                                                                                   `json:"250ms_or_greater_event_fraction"`
	Number500msOrGreaterEventFraction float64                                                                                                   `json:"500ms_or_greater_event_fraction"`
	Avg                               float64                                                                                                   `json:"avg"`
	JSON                              sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeRTTJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeRTTJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeRTT]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeRTTJSON struct {
	Number100msOrGreaterEventFraction apijson.Field
	Number250msOrGreaterEventFraction apijson.Field
	Number500msOrGreaterEventFraction apijson.Field
	Avg                               apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeRTT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeRTTJSON) RawJSON() string {
	return r.raw
}

// Per-sample inbound (consumer) audio statistics.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumer struct {
	BytesReceived            float64                                                                                                 `json:"bytes_received"`
	ConcealmentEvents        float64                                                                                                 `json:"concealment_events"`
	ConsumerID               string                                                                                                  `json:"consumer_id"`
	Jitter                   float64                                                                                                 `json:"jitter"`
	JitterBufferDelay        float64                                                                                                 `json:"jitter_buffer_delay"`
	JitterBufferEmittedCount float64                                                                                                 `json:"jitter_buffer_emitted_count"`
	Mid                      string                                                                                                  `json:"mid"`
	MosQuality               float64                                                                                                 `json:"mos_quality"`
	PacketsLost              float64                                                                                                 `json:"packets_lost"`
	PacketsReceived          float64                                                                                                 `json:"packets_received"`
	PeerID                   string                                                                                                  `json:"peer_id"`
	ProducerID               string                                                                                                  `json:"producer_id"`
	Ssrc                     float64                                                                                                 `json:"ssrc"`
	Timestamp                string                                                                                                  `json:"timestamp"`
	JSON                     sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumer]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerJSON struct {
	BytesReceived            apijson.Field
	ConcealmentEvents        apijson.Field
	ConsumerID               apijson.Field
	Jitter                   apijson.Field
	JitterBufferDelay        apijson.Field
	JitterBufferEmittedCount apijson.Field
	Mid                      apijson.Field
	MosQuality               apijson.Field
	PacketsLost              apijson.Field
	PacketsReceived          apijson.Field
	PeerID                   apijson.Field
	ProducerID               apijson.Field
	Ssrc                     apijson.Field
	Timestamp                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerJSON) RawJSON() string {
	return r.raw
}

// Aggregated inbound (consumer) audio statistics for the session.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulative struct {
	// Cumulative latency distribution (milliseconds-based thresholds).
	JitterBufferDelay SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeJitterBufferDelay `json:"jitter_buffer_delay"`
	// Cumulative packet loss distribution.
	PacketLoss SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativePacketLoss `json:"packet_loss"`
	// Distribution summary with average and percentiles.
	QualityMos SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeQualityMos `json:"quality_mos"`
	JSON       sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeJSON       `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulative]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeJSON struct {
	JitterBufferDelay apijson.Field
	PacketLoss        apijson.Field
	QualityMos        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeJSON) RawJSON() string {
	return r.raw
}

// Cumulative latency distribution (milliseconds-based thresholds).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeJitterBufferDelay struct {
	Number100msOrGreaterEventFraction float64                                                                                                                            `json:"100ms_or_greater_event_fraction"`
	Number250msOrGreaterEventFraction float64                                                                                                                            `json:"250ms_or_greater_event_fraction"`
	Number500msOrGreaterEventFraction float64                                                                                                                            `json:"500ms_or_greater_event_fraction"`
	Avg                               float64                                                                                                                            `json:"avg"`
	JSON                              sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeJitterBufferDelayJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeJitterBufferDelayJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeJitterBufferDelay]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeJitterBufferDelayJSON struct {
	Number100msOrGreaterEventFraction apijson.Field
	Number250msOrGreaterEventFraction apijson.Field
	Number500msOrGreaterEventFraction apijson.Field
	Avg                               apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeJitterBufferDelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeJitterBufferDelayJSON) RawJSON() string {
	return r.raw
}

// Cumulative packet loss distribution.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativePacketLoss struct {
	Number10OrGreaterEventFraction float64                                                                                                                     `json:"10_or_greater_event_fraction"`
	Number25OrGreaterEventFraction float64                                                                                                                     `json:"25_or_greater_event_fraction"`
	Number5OrGreaterEventFraction  float64                                                                                                                     `json:"5_or_greater_event_fraction"`
	Number50OrGreaterEventFraction float64                                                                                                                     `json:"50_or_greater_event_fraction"`
	Avg                            float64                                                                                                                     `json:"avg"`
	JSON                           sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativePacketLossJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativePacketLossJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativePacketLoss]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativePacketLossJSON struct {
	Number10OrGreaterEventFraction apijson.Field
	Number25OrGreaterEventFraction apijson.Field
	Number5OrGreaterEventFraction  apijson.Field
	Number50OrGreaterEventFraction apijson.Field
	Avg                            apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativePacketLoss) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativePacketLossJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeQualityMos struct {
	Avg  float64                                                                                                                     `json:"avg"`
	P50  float64                                                                                                                     `json:"p50"`
	P75  float64                                                                                                                     `json:"p75"`
	P90  float64                                                                                                                     `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeQualityMosJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeQualityMosJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeQualityMos]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeQualityMosJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeQualityMos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioConsumerCumulativeQualityMosJSON) RawJSON() string {
	return r.raw
}

// Per-sample outbound (producer) audio statistics.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducer struct {
	BytesSent   float64                                                                                                 `json:"bytes_sent"`
	Jitter      float64                                                                                                 `json:"jitter"`
	Mid         string                                                                                                  `json:"mid"`
	MosQuality  float64                                                                                                 `json:"mos_quality"`
	PacketsLost float64                                                                                                 `json:"packets_lost"`
	PacketsSent float64                                                                                                 `json:"packets_sent"`
	ProducerID  string                                                                                                  `json:"producer_id"`
	RTT         float64                                                                                                 `json:"rtt"`
	Ssrc        float64                                                                                                 `json:"ssrc"`
	Timestamp   string                                                                                                  `json:"timestamp"`
	JSON        sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducer]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerJSON struct {
	BytesSent   apijson.Field
	Jitter      apijson.Field
	Mid         apijson.Field
	MosQuality  apijson.Field
	PacketsLost apijson.Field
	PacketsSent apijson.Field
	ProducerID  apijson.Field
	RTT         apijson.Field
	Ssrc        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerJSON) RawJSON() string {
	return r.raw
}

// Aggregated outbound (producer) audio statistics for the session.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulative struct {
	// Cumulative packet loss distribution.
	PacketLoss SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativePacketLoss `json:"packet_loss"`
	// Distribution summary with average and percentiles.
	QualityMos SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeQualityMos `json:"quality_mos"`
	// Cumulative latency distribution (milliseconds-based thresholds).
	RTT  SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeRTT  `json:"rtt"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulative]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeJSON struct {
	PacketLoss  apijson.Field
	QualityMos  apijson.Field
	RTT         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeJSON) RawJSON() string {
	return r.raw
}

// Cumulative packet loss distribution.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativePacketLoss struct {
	Number10OrGreaterEventFraction float64                                                                                                                     `json:"10_or_greater_event_fraction"`
	Number25OrGreaterEventFraction float64                                                                                                                     `json:"25_or_greater_event_fraction"`
	Number5OrGreaterEventFraction  float64                                                                                                                     `json:"5_or_greater_event_fraction"`
	Number50OrGreaterEventFraction float64                                                                                                                     `json:"50_or_greater_event_fraction"`
	Avg                            float64                                                                                                                     `json:"avg"`
	JSON                           sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativePacketLossJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativePacketLossJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativePacketLoss]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativePacketLossJSON struct {
	Number10OrGreaterEventFraction apijson.Field
	Number25OrGreaterEventFraction apijson.Field
	Number5OrGreaterEventFraction  apijson.Field
	Number50OrGreaterEventFraction apijson.Field
	Avg                            apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativePacketLoss) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativePacketLossJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeQualityMos struct {
	Avg  float64                                                                                                                     `json:"avg"`
	P50  float64                                                                                                                     `json:"p50"`
	P75  float64                                                                                                                     `json:"p75"`
	P90  float64                                                                                                                     `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeQualityMosJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeQualityMosJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeQualityMos]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeQualityMosJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeQualityMos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeQualityMosJSON) RawJSON() string {
	return r.raw
}

// Cumulative latency distribution (milliseconds-based thresholds).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeRTT struct {
	Number100msOrGreaterEventFraction float64                                                                                                              `json:"100ms_or_greater_event_fraction"`
	Number250msOrGreaterEventFraction float64                                                                                                              `json:"250ms_or_greater_event_fraction"`
	Number500msOrGreaterEventFraction float64                                                                                                              `json:"500ms_or_greater_event_fraction"`
	Avg                               float64                                                                                                              `json:"avg"`
	JSON                              sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeRTTJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeRTTJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeRTT]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeRTTJSON struct {
	Number100msOrGreaterEventFraction apijson.Field
	Number250msOrGreaterEventFraction apijson.Field
	Number500msOrGreaterEventFraction apijson.Field
	Avg                               apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeRTT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareAudioProducerCumulativeRTTJSON) RawJSON() string {
	return r.raw
}

// Per-sample inbound (consumer) video statistics.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumer struct {
	BytesReceived            float64                                                                                                 `json:"bytes_received"`
	ConsumerID               string                                                                                                  `json:"consumer_id"`
	FirCount                 float64                                                                                                 `json:"fir_count"`
	FrameHeight              float64                                                                                                 `json:"frame_height"`
	FrameWidth               float64                                                                                                 `json:"frame_width"`
	FramesDecoded            float64                                                                                                 `json:"frames_decoded"`
	FramesDropped            float64                                                                                                 `json:"frames_dropped"`
	FramesPerSecond          float64                                                                                                 `json:"frames_per_second"`
	Jitter                   float64                                                                                                 `json:"jitter"`
	JitterBufferDelay        float64                                                                                                 `json:"jitter_buffer_delay"`
	JitterBufferEmittedCount float64                                                                                                 `json:"jitter_buffer_emitted_count"`
	KeyFramesDecoded         float64                                                                                                 `json:"key_frames_decoded"`
	Mid                      string                                                                                                  `json:"mid"`
	MosQuality               float64                                                                                                 `json:"mos_quality"`
	PacketsLost              float64                                                                                                 `json:"packets_lost"`
	PacketsReceived          float64                                                                                                 `json:"packets_received"`
	PeerID                   string                                                                                                  `json:"peer_id"`
	ProducerID               string                                                                                                  `json:"producer_id"`
	Ssrc                     float64                                                                                                 `json:"ssrc"`
	Timestamp                string                                                                                                  `json:"timestamp"`
	JSON                     sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumer]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerJSON struct {
	BytesReceived            apijson.Field
	ConsumerID               apijson.Field
	FirCount                 apijson.Field
	FrameHeight              apijson.Field
	FrameWidth               apijson.Field
	FramesDecoded            apijson.Field
	FramesDropped            apijson.Field
	FramesPerSecond          apijson.Field
	Jitter                   apijson.Field
	JitterBufferDelay        apijson.Field
	JitterBufferEmittedCount apijson.Field
	KeyFramesDecoded         apijson.Field
	Mid                      apijson.Field
	MosQuality               apijson.Field
	PacketsLost              apijson.Field
	PacketsReceived          apijson.Field
	PeerID                   apijson.Field
	ProducerID               apijson.Field
	Ssrc                     apijson.Field
	Timestamp                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerJSON) RawJSON() string {
	return r.raw
}

// Aggregated inbound (consumer) video statistics for the session.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulative struct {
	// Distribution summary with average and percentiles.
	FramePerSecond SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFramePerSecond `json:"frame_per_second"`
	// Distribution summary with average and percentiles.
	FrameWidth SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFrameWidth `json:"frame_width"`
	Issues     SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeIssues     `json:"issues"`
	// Cumulative latency distribution (milliseconds-based thresholds).
	JitterBufferDelay        SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeJitterBufferDelay `json:"jitter_buffer_delay"`
	KeyFramesDecodedFraction float64                                                                                                                        `json:"key_frames_decoded_fraction"`
	// Cumulative packet loss distribution.
	PacketLoss SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativePacketLoss `json:"packet_loss"`
	// Distribution summary with average and percentiles.
	QualityMos SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeQualityMos `json:"quality_mos"`
	JSON       sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeJSON       `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulative]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeJSON struct {
	FramePerSecond           apijson.Field
	FrameWidth               apijson.Field
	Issues                   apijson.Field
	JitterBufferDelay        apijson.Field
	KeyFramesDecodedFraction apijson.Field
	PacketLoss               apijson.Field
	QualityMos               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFramePerSecond struct {
	Avg  float64                                                                                                                         `json:"avg"`
	P50  float64                                                                                                                         `json:"p50"`
	P75  float64                                                                                                                         `json:"p75"`
	P90  float64                                                                                                                         `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFramePerSecondJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFramePerSecondJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFramePerSecond]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFramePerSecondJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFramePerSecond) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFramePerSecondJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFrameWidth struct {
	Avg  float64                                                                                                                     `json:"avg"`
	P50  float64                                                                                                                     `json:"p50"`
	P75  float64                                                                                                                     `json:"p75"`
	P90  float64                                                                                                                     `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFrameWidthJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFrameWidthJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFrameWidth]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFrameWidthJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFrameWidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeFrameWidthJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeIssues struct {
	LagFraction            float64                                                                                                                 `json:"lag_fraction"`
	NoVideoFraction        float64                                                                                                                 `json:"no_video_fraction"`
	PoorResolutionFraction float64                                                                                                                 `json:"poor_resolution_fraction"`
	JSON                   sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeIssuesJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeIssuesJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeIssues]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeIssuesJSON struct {
	LagFraction            apijson.Field
	NoVideoFraction        apijson.Field
	PoorResolutionFraction apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeIssues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeIssuesJSON) RawJSON() string {
	return r.raw
}

// Cumulative latency distribution (milliseconds-based thresholds).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeJitterBufferDelay struct {
	Number100msOrGreaterEventFraction float64                                                                                                                            `json:"100ms_or_greater_event_fraction"`
	Number250msOrGreaterEventFraction float64                                                                                                                            `json:"250ms_or_greater_event_fraction"`
	Number500msOrGreaterEventFraction float64                                                                                                                            `json:"500ms_or_greater_event_fraction"`
	Avg                               float64                                                                                                                            `json:"avg"`
	JSON                              sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeJitterBufferDelayJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeJitterBufferDelayJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeJitterBufferDelay]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeJitterBufferDelayJSON struct {
	Number100msOrGreaterEventFraction apijson.Field
	Number250msOrGreaterEventFraction apijson.Field
	Number500msOrGreaterEventFraction apijson.Field
	Avg                               apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeJitterBufferDelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeJitterBufferDelayJSON) RawJSON() string {
	return r.raw
}

// Cumulative packet loss distribution.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativePacketLoss struct {
	Number10OrGreaterEventFraction float64                                                                                                                     `json:"10_or_greater_event_fraction"`
	Number25OrGreaterEventFraction float64                                                                                                                     `json:"25_or_greater_event_fraction"`
	Number5OrGreaterEventFraction  float64                                                                                                                     `json:"5_or_greater_event_fraction"`
	Number50OrGreaterEventFraction float64                                                                                                                     `json:"50_or_greater_event_fraction"`
	Avg                            float64                                                                                                                     `json:"avg"`
	JSON                           sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativePacketLossJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativePacketLossJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativePacketLoss]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativePacketLossJSON struct {
	Number10OrGreaterEventFraction apijson.Field
	Number25OrGreaterEventFraction apijson.Field
	Number5OrGreaterEventFraction  apijson.Field
	Number50OrGreaterEventFraction apijson.Field
	Avg                            apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativePacketLoss) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativePacketLossJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeQualityMos struct {
	Avg  float64                                                                                                                     `json:"avg"`
	P50  float64                                                                                                                     `json:"p50"`
	P75  float64                                                                                                                     `json:"p75"`
	P90  float64                                                                                                                     `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeQualityMosJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeQualityMosJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeQualityMos]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeQualityMosJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeQualityMos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoConsumerCumulativeQualityMosJSON) RawJSON() string {
	return r.raw
}

// Per-sample outbound (producer) video statistics.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducer struct {
	BytesSent                          float64                                                                                                                       `json:"bytes_sent"`
	FirCount                           float64                                                                                                                       `json:"fir_count"`
	FrameHeight                        float64                                                                                                                       `json:"frame_height"`
	FrameWidth                         float64                                                                                                                       `json:"frame_width"`
	FramesEncoded                      float64                                                                                                                       `json:"frames_encoded"`
	FramesPerSecond                    float64                                                                                                                       `json:"frames_per_second"`
	Jitter                             float64                                                                                                                       `json:"jitter"`
	KeyFramesEncoded                   float64                                                                                                                       `json:"key_frames_encoded"`
	Mid                                string                                                                                                                        `json:"mid"`
	MosQuality                         float64                                                                                                                       `json:"mos_quality"`
	PacketsLost                        float64                                                                                                                       `json:"packets_lost"`
	PacketsSent                        float64                                                                                                                       `json:"packets_sent"`
	PliCount                           float64                                                                                                                       `json:"pli_count"`
	ProducerID                         string                                                                                                                        `json:"producer_id"`
	QualityLimitationDurations         SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationDurations `json:"quality_limitation_durations"`
	QualityLimitationReason            SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReason    `json:"quality_limitation_reason"`
	QualityLimitationResolutionChanges float64                                                                                                                       `json:"quality_limitation_resolution_changes"`
	RTT                                float64                                                                                                                       `json:"rtt"`
	Ssrc                               float64                                                                                                                       `json:"ssrc"`
	Timestamp                          string                                                                                                                        `json:"timestamp"`
	JSON                               sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerJSON                       `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducer]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerJSON struct {
	BytesSent                          apijson.Field
	FirCount                           apijson.Field
	FrameHeight                        apijson.Field
	FrameWidth                         apijson.Field
	FramesEncoded                      apijson.Field
	FramesPerSecond                    apijson.Field
	Jitter                             apijson.Field
	KeyFramesEncoded                   apijson.Field
	Mid                                apijson.Field
	MosQuality                         apijson.Field
	PacketsLost                        apijson.Field
	PacketsSent                        apijson.Field
	PliCount                           apijson.Field
	ProducerID                         apijson.Field
	QualityLimitationDurations         apijson.Field
	QualityLimitationReason            apijson.Field
	QualityLimitationResolutionChanges apijson.Field
	RTT                                apijson.Field
	Ssrc                               apijson.Field
	Timestamp                          apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationDurations struct {
	Bandwidth float64                                                                                                                           `json:"bandwidth"`
	CPU       float64                                                                                                                           `json:"cpu"`
	None      float64                                                                                                                           `json:"none"`
	Other     float64                                                                                                                           `json:"other"`
	JSON      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationDurationsJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationDurationsJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationDurations]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationDurationsJSON struct {
	Bandwidth   apijson.Field
	CPU         apijson.Field
	None        apijson.Field
	Other       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationDurations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationDurationsJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReason string

const (
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReasonCPU       SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReason = "cpu"
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReasonBandwidth SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReason = "bandwidth"
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReasonNone      SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReason = "none"
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReasonOther     SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReason = "other"
)

func (r SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReason) IsKnown() bool {
	switch r {
	case SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReasonCPU, SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReasonBandwidth, SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReasonNone, SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerQualityLimitationReasonOther:
		return true
	}
	return false
}

// Aggregated outbound (producer) video statistics for the session.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulative struct {
	// Distribution summary with average and percentiles.
	FramePerSecond SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFramePerSecond `json:"frame_per_second"`
	// Distribution summary with average and percentiles.
	FrameWidth                   SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFrameWidth `json:"frame_width"`
	HighNegativeFeedbackFraction float64                                                                                                                 `json:"high_negative_feedback_fraction"`
	Issues                       SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeIssues     `json:"issues"`
	KeyFramesEncodedFraction     float64                                                                                                                 `json:"key_frames_encoded_fraction"`
	// Cumulative packet loss distribution.
	PacketLoss SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativePacketLoss `json:"packet_loss"`
	// Distribution summary with average and percentiles.
	QualityMos SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeQualityMos `json:"quality_mos"`
	// Cumulative latency distribution (milliseconds-based thresholds).
	RTT  SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeRTT  `json:"rtt"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulative]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeJSON struct {
	FramePerSecond               apijson.Field
	FrameWidth                   apijson.Field
	HighNegativeFeedbackFraction apijson.Field
	Issues                       apijson.Field
	KeyFramesEncodedFraction     apijson.Field
	PacketLoss                   apijson.Field
	QualityMos                   apijson.Field
	RTT                          apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFramePerSecond struct {
	Avg  float64                                                                                                                         `json:"avg"`
	P50  float64                                                                                                                         `json:"p50"`
	P75  float64                                                                                                                         `json:"p75"`
	P90  float64                                                                                                                         `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFramePerSecondJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFramePerSecondJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFramePerSecond]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFramePerSecondJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFramePerSecond) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFramePerSecondJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFrameWidth struct {
	Avg  float64                                                                                                                     `json:"avg"`
	P50  float64                                                                                                                     `json:"p50"`
	P75  float64                                                                                                                     `json:"p75"`
	P90  float64                                                                                                                     `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFrameWidthJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFrameWidthJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFrameWidth]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFrameWidthJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFrameWidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeFrameWidthJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeIssues struct {
	BandwidthQualityLimitationFraction float64                                                                                                                 `json:"bandwidth_quality_limitation_fraction"`
	CPUQualityLimitationFraction       float64                                                                                                                 `json:"cpu_quality_limitation_fraction"`
	NoVideoFraction                    float64                                                                                                                 `json:"no_video_fraction"`
	PoorResolutionFraction             float64                                                                                                                 `json:"poor_resolution_fraction"`
	QualityLimitationFraction          float64                                                                                                                 `json:"quality_limitation_fraction"`
	JSON                               sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeIssuesJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeIssuesJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeIssues]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeIssuesJSON struct {
	BandwidthQualityLimitationFraction apijson.Field
	CPUQualityLimitationFraction       apijson.Field
	NoVideoFraction                    apijson.Field
	PoorResolutionFraction             apijson.Field
	QualityLimitationFraction          apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeIssues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeIssuesJSON) RawJSON() string {
	return r.raw
}

// Cumulative packet loss distribution.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativePacketLoss struct {
	Number10OrGreaterEventFraction float64                                                                                                                     `json:"10_or_greater_event_fraction"`
	Number25OrGreaterEventFraction float64                                                                                                                     `json:"25_or_greater_event_fraction"`
	Number5OrGreaterEventFraction  float64                                                                                                                     `json:"5_or_greater_event_fraction"`
	Number50OrGreaterEventFraction float64                                                                                                                     `json:"50_or_greater_event_fraction"`
	Avg                            float64                                                                                                                     `json:"avg"`
	JSON                           sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativePacketLossJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativePacketLossJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativePacketLoss]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativePacketLossJSON struct {
	Number10OrGreaterEventFraction apijson.Field
	Number25OrGreaterEventFraction apijson.Field
	Number5OrGreaterEventFraction  apijson.Field
	Number50OrGreaterEventFraction apijson.Field
	Avg                            apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativePacketLoss) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativePacketLossJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeQualityMos struct {
	Avg  float64                                                                                                                     `json:"avg"`
	P50  float64                                                                                                                     `json:"p50"`
	P75  float64                                                                                                                     `json:"p75"`
	P90  float64                                                                                                                     `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeQualityMosJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeQualityMosJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeQualityMos]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeQualityMosJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeQualityMos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeQualityMosJSON) RawJSON() string {
	return r.raw
}

// Cumulative latency distribution (milliseconds-based thresholds).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeRTT struct {
	Number100msOrGreaterEventFraction float64                                                                                                              `json:"100ms_or_greater_event_fraction"`
	Number250msOrGreaterEventFraction float64                                                                                                              `json:"250ms_or_greater_event_fraction"`
	Number500msOrGreaterEventFraction float64                                                                                                              `json:"500ms_or_greater_event_fraction"`
	Avg                               float64                                                                                                              `json:"avg"`
	JSON                              sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeRTTJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeRTTJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeRTT]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeRTTJSON struct {
	Number100msOrGreaterEventFraction apijson.Field
	Number250msOrGreaterEventFraction apijson.Field
	Number500msOrGreaterEventFraction apijson.Field
	Avg                               apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeRTT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityScreenshareVideoProducerCumulativeRTTJSON) RawJSON() string {
	return r.raw
}

// Per-sample inbound (consumer) video statistics.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumer struct {
	BytesReceived            float64                                                                                      `json:"bytes_received"`
	ConsumerID               string                                                                                       `json:"consumer_id"`
	FirCount                 float64                                                                                      `json:"fir_count"`
	FrameHeight              float64                                                                                      `json:"frame_height"`
	FrameWidth               float64                                                                                      `json:"frame_width"`
	FramesDecoded            float64                                                                                      `json:"frames_decoded"`
	FramesDropped            float64                                                                                      `json:"frames_dropped"`
	FramesPerSecond          float64                                                                                      `json:"frames_per_second"`
	Jitter                   float64                                                                                      `json:"jitter"`
	JitterBufferDelay        float64                                                                                      `json:"jitter_buffer_delay"`
	JitterBufferEmittedCount float64                                                                                      `json:"jitter_buffer_emitted_count"`
	KeyFramesDecoded         float64                                                                                      `json:"key_frames_decoded"`
	Mid                      string                                                                                       `json:"mid"`
	MosQuality               float64                                                                                      `json:"mos_quality"`
	PacketsLost              float64                                                                                      `json:"packets_lost"`
	PacketsReceived          float64                                                                                      `json:"packets_received"`
	PeerID                   string                                                                                       `json:"peer_id"`
	ProducerID               string                                                                                       `json:"producer_id"`
	Ssrc                     float64                                                                                      `json:"ssrc"`
	Timestamp                string                                                                                       `json:"timestamp"`
	JSON                     sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumer]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerJSON struct {
	BytesReceived            apijson.Field
	ConsumerID               apijson.Field
	FirCount                 apijson.Field
	FrameHeight              apijson.Field
	FrameWidth               apijson.Field
	FramesDecoded            apijson.Field
	FramesDropped            apijson.Field
	FramesPerSecond          apijson.Field
	Jitter                   apijson.Field
	JitterBufferDelay        apijson.Field
	JitterBufferEmittedCount apijson.Field
	KeyFramesDecoded         apijson.Field
	Mid                      apijson.Field
	MosQuality               apijson.Field
	PacketsLost              apijson.Field
	PacketsReceived          apijson.Field
	PeerID                   apijson.Field
	ProducerID               apijson.Field
	Ssrc                     apijson.Field
	Timestamp                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerJSON) RawJSON() string {
	return r.raw
}

// Aggregated inbound (consumer) video statistics for the session.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulative struct {
	// Distribution summary with average and percentiles.
	FramePerSecond SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFramePerSecond `json:"frame_per_second"`
	// Distribution summary with average and percentiles.
	FrameWidth SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFrameWidth `json:"frame_width"`
	Issues     SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeIssues     `json:"issues"`
	// Cumulative latency distribution (milliseconds-based thresholds).
	JitterBufferDelay        SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeJitterBufferDelay `json:"jitter_buffer_delay"`
	KeyFramesDecodedFraction float64                                                                                                             `json:"key_frames_decoded_fraction"`
	// Cumulative packet loss distribution.
	PacketLoss SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativePacketLoss `json:"packet_loss"`
	// Distribution summary with average and percentiles.
	QualityMos SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeQualityMos `json:"quality_mos"`
	JSON       sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeJSON       `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulative]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeJSON struct {
	FramePerSecond           apijson.Field
	FrameWidth               apijson.Field
	Issues                   apijson.Field
	JitterBufferDelay        apijson.Field
	KeyFramesDecodedFraction apijson.Field
	PacketLoss               apijson.Field
	QualityMos               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFramePerSecond struct {
	Avg  float64                                                                                                              `json:"avg"`
	P50  float64                                                                                                              `json:"p50"`
	P75  float64                                                                                                              `json:"p75"`
	P90  float64                                                                                                              `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFramePerSecondJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFramePerSecondJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFramePerSecond]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFramePerSecondJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFramePerSecond) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFramePerSecondJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFrameWidth struct {
	Avg  float64                                                                                                          `json:"avg"`
	P50  float64                                                                                                          `json:"p50"`
	P75  float64                                                                                                          `json:"p75"`
	P90  float64                                                                                                          `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFrameWidthJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFrameWidthJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFrameWidth]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFrameWidthJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFrameWidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeFrameWidthJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeIssues struct {
	LagFraction            float64                                                                                                      `json:"lag_fraction"`
	NoVideoFraction        float64                                                                                                      `json:"no_video_fraction"`
	PoorResolutionFraction float64                                                                                                      `json:"poor_resolution_fraction"`
	JSON                   sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeIssuesJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeIssuesJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeIssues]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeIssuesJSON struct {
	LagFraction            apijson.Field
	NoVideoFraction        apijson.Field
	PoorResolutionFraction apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeIssues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeIssuesJSON) RawJSON() string {
	return r.raw
}

// Cumulative latency distribution (milliseconds-based thresholds).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeJitterBufferDelay struct {
	Number100msOrGreaterEventFraction float64                                                                                                                 `json:"100ms_or_greater_event_fraction"`
	Number250msOrGreaterEventFraction float64                                                                                                                 `json:"250ms_or_greater_event_fraction"`
	Number500msOrGreaterEventFraction float64                                                                                                                 `json:"500ms_or_greater_event_fraction"`
	Avg                               float64                                                                                                                 `json:"avg"`
	JSON                              sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeJitterBufferDelayJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeJitterBufferDelayJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeJitterBufferDelay]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeJitterBufferDelayJSON struct {
	Number100msOrGreaterEventFraction apijson.Field
	Number250msOrGreaterEventFraction apijson.Field
	Number500msOrGreaterEventFraction apijson.Field
	Avg                               apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeJitterBufferDelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeJitterBufferDelayJSON) RawJSON() string {
	return r.raw
}

// Cumulative packet loss distribution.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativePacketLoss struct {
	Number10OrGreaterEventFraction float64                                                                                                          `json:"10_or_greater_event_fraction"`
	Number25OrGreaterEventFraction float64                                                                                                          `json:"25_or_greater_event_fraction"`
	Number5OrGreaterEventFraction  float64                                                                                                          `json:"5_or_greater_event_fraction"`
	Number50OrGreaterEventFraction float64                                                                                                          `json:"50_or_greater_event_fraction"`
	Avg                            float64                                                                                                          `json:"avg"`
	JSON                           sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativePacketLossJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativePacketLossJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativePacketLoss]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativePacketLossJSON struct {
	Number10OrGreaterEventFraction apijson.Field
	Number25OrGreaterEventFraction apijson.Field
	Number5OrGreaterEventFraction  apijson.Field
	Number50OrGreaterEventFraction apijson.Field
	Avg                            apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativePacketLoss) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativePacketLossJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeQualityMos struct {
	Avg  float64                                                                                                          `json:"avg"`
	P50  float64                                                                                                          `json:"p50"`
	P75  float64                                                                                                          `json:"p75"`
	P90  float64                                                                                                          `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeQualityMosJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeQualityMosJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeQualityMos]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeQualityMosJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeQualityMos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoConsumerCumulativeQualityMosJSON) RawJSON() string {
	return r.raw
}

// Per-sample outbound (producer) video statistics.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducer struct {
	BytesSent                          float64                                                                                                            `json:"bytes_sent"`
	FirCount                           float64                                                                                                            `json:"fir_count"`
	FrameHeight                        float64                                                                                                            `json:"frame_height"`
	FrameWidth                         float64                                                                                                            `json:"frame_width"`
	FramesEncoded                      float64                                                                                                            `json:"frames_encoded"`
	FramesPerSecond                    float64                                                                                                            `json:"frames_per_second"`
	Jitter                             float64                                                                                                            `json:"jitter"`
	KeyFramesEncoded                   float64                                                                                                            `json:"key_frames_encoded"`
	Mid                                string                                                                                                             `json:"mid"`
	MosQuality                         float64                                                                                                            `json:"mos_quality"`
	PacketsLost                        float64                                                                                                            `json:"packets_lost"`
	PacketsSent                        float64                                                                                                            `json:"packets_sent"`
	PliCount                           float64                                                                                                            `json:"pli_count"`
	ProducerID                         string                                                                                                             `json:"producer_id"`
	QualityLimitationDurations         SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationDurations `json:"quality_limitation_durations"`
	QualityLimitationReason            SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReason    `json:"quality_limitation_reason"`
	QualityLimitationResolutionChanges float64                                                                                                            `json:"quality_limitation_resolution_changes"`
	RTT                                float64                                                                                                            `json:"rtt"`
	Ssrc                               float64                                                                                                            `json:"ssrc"`
	Timestamp                          string                                                                                                             `json:"timestamp"`
	JSON                               sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerJSON                       `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducer]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerJSON struct {
	BytesSent                          apijson.Field
	FirCount                           apijson.Field
	FrameHeight                        apijson.Field
	FrameWidth                         apijson.Field
	FramesEncoded                      apijson.Field
	FramesPerSecond                    apijson.Field
	Jitter                             apijson.Field
	KeyFramesEncoded                   apijson.Field
	Mid                                apijson.Field
	MosQuality                         apijson.Field
	PacketsLost                        apijson.Field
	PacketsSent                        apijson.Field
	PliCount                           apijson.Field
	ProducerID                         apijson.Field
	QualityLimitationDurations         apijson.Field
	QualityLimitationReason            apijson.Field
	QualityLimitationResolutionChanges apijson.Field
	RTT                                apijson.Field
	Ssrc                               apijson.Field
	Timestamp                          apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationDurations struct {
	Bandwidth float64                                                                                                                `json:"bandwidth"`
	CPU       float64                                                                                                                `json:"cpu"`
	None      float64                                                                                                                `json:"none"`
	Other     float64                                                                                                                `json:"other"`
	JSON      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationDurationsJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationDurationsJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationDurations]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationDurationsJSON struct {
	Bandwidth   apijson.Field
	CPU         apijson.Field
	None        apijson.Field
	Other       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationDurations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationDurationsJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReason string

const (
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReasonCPU       SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReason = "cpu"
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReasonBandwidth SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReason = "bandwidth"
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReasonNone      SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReason = "none"
	SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReasonOther     SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReason = "other"
)

func (r SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReason) IsKnown() bool {
	switch r {
	case SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReasonCPU, SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReasonBandwidth, SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReasonNone, SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerQualityLimitationReasonOther:
		return true
	}
	return false
}

// Aggregated outbound (producer) video statistics for the session.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulative struct {
	// Distribution summary with average and percentiles.
	FramePerSecond SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFramePerSecond `json:"frame_per_second"`
	// Distribution summary with average and percentiles.
	FrameWidth                   SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFrameWidth `json:"frame_width"`
	HighNegativeFeedbackFraction float64                                                                                                      `json:"high_negative_feedback_fraction"`
	Issues                       SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeIssues     `json:"issues"`
	KeyFramesEncodedFraction     float64                                                                                                      `json:"key_frames_encoded_fraction"`
	// Cumulative packet loss distribution.
	PacketLoss SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativePacketLoss `json:"packet_loss"`
	// Distribution summary with average and percentiles.
	QualityMos SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeQualityMos `json:"quality_mos"`
	// Cumulative latency distribution (milliseconds-based thresholds).
	RTT  SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeRTT  `json:"rtt"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulative]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeJSON struct {
	FramePerSecond               apijson.Field
	FrameWidth                   apijson.Field
	HighNegativeFeedbackFraction apijson.Field
	Issues                       apijson.Field
	KeyFramesEncodedFraction     apijson.Field
	PacketLoss                   apijson.Field
	QualityMos                   apijson.Field
	RTT                          apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFramePerSecond struct {
	Avg  float64                                                                                                              `json:"avg"`
	P50  float64                                                                                                              `json:"p50"`
	P75  float64                                                                                                              `json:"p75"`
	P90  float64                                                                                                              `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFramePerSecondJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFramePerSecondJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFramePerSecond]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFramePerSecondJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFramePerSecond) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFramePerSecondJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFrameWidth struct {
	Avg  float64                                                                                                          `json:"avg"`
	P50  float64                                                                                                          `json:"p50"`
	P75  float64                                                                                                          `json:"p75"`
	P90  float64                                                                                                          `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFrameWidthJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFrameWidthJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFrameWidth]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFrameWidthJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFrameWidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeFrameWidthJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeIssues struct {
	BandwidthQualityLimitationFraction float64                                                                                                      `json:"bandwidth_quality_limitation_fraction"`
	CPUQualityLimitationFraction       float64                                                                                                      `json:"cpu_quality_limitation_fraction"`
	NoVideoFraction                    float64                                                                                                      `json:"no_video_fraction"`
	PoorResolutionFraction             float64                                                                                                      `json:"poor_resolution_fraction"`
	QualityLimitationFraction          float64                                                                                                      `json:"quality_limitation_fraction"`
	JSON                               sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeIssuesJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeIssuesJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeIssues]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeIssuesJSON struct {
	BandwidthQualityLimitationFraction apijson.Field
	CPUQualityLimitationFraction       apijson.Field
	NoVideoFraction                    apijson.Field
	PoorResolutionFraction             apijson.Field
	QualityLimitationFraction          apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeIssues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeIssuesJSON) RawJSON() string {
	return r.raw
}

// Cumulative packet loss distribution.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativePacketLoss struct {
	Number10OrGreaterEventFraction float64                                                                                                          `json:"10_or_greater_event_fraction"`
	Number25OrGreaterEventFraction float64                                                                                                          `json:"25_or_greater_event_fraction"`
	Number5OrGreaterEventFraction  float64                                                                                                          `json:"5_or_greater_event_fraction"`
	Number50OrGreaterEventFraction float64                                                                                                          `json:"50_or_greater_event_fraction"`
	Avg                            float64                                                                                                          `json:"avg"`
	JSON                           sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativePacketLossJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativePacketLossJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativePacketLoss]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativePacketLossJSON struct {
	Number10OrGreaterEventFraction apijson.Field
	Number25OrGreaterEventFraction apijson.Field
	Number5OrGreaterEventFraction  apijson.Field
	Number50OrGreaterEventFraction apijson.Field
	Avg                            apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativePacketLoss) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativePacketLossJSON) RawJSON() string {
	return r.raw
}

// Distribution summary with average and percentiles.
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeQualityMos struct {
	Avg  float64                                                                                                          `json:"avg"`
	P50  float64                                                                                                          `json:"p50"`
	P75  float64                                                                                                          `json:"p75"`
	P90  float64                                                                                                          `json:"p90"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeQualityMosJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeQualityMosJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeQualityMos]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeQualityMosJSON struct {
	Avg         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	P90         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeQualityMos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeQualityMosJSON) RawJSON() string {
	return r.raw
}

// Cumulative latency distribution (milliseconds-based thresholds).
type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeRTT struct {
	Number100msOrGreaterEventFraction float64                                                                                                   `json:"100ms_or_greater_event_fraction"`
	Number250msOrGreaterEventFraction float64                                                                                                   `json:"250ms_or_greater_event_fraction"`
	Number500msOrGreaterEventFraction float64                                                                                                   `json:"500ms_or_greater_event_fraction"`
	Avg                               float64                                                                                                   `json:"avg"`
	JSON                              sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeRTTJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeRTTJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeRTT]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeRTTJSON struct {
	Number100msOrGreaterEventFraction apijson.Field
	Number250msOrGreaterEventFraction apijson.Field
	Number500msOrGreaterEventFraction apijson.Field
	Avg                               apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeRTT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityVideoProducerCumulativeRTTJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionChatResponse struct {
	Data    SessionGetSessionChatResponseData `json:"data"`
	Success bool                              `json:"success"`
	JSON    sessionGetSessionChatResponseJSON `json:"-"`
}

// sessionGetSessionChatResponseJSON contains the JSON metadata for the struct
// [SessionGetSessionChatResponse]
type sessionGetSessionChatResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetSessionChatResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionChatResponseJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionChatResponseData struct {
	// URL where the chat logs can be downloaded
	ChatDownloadURL string `json:"chat_download_url" api:"required"`
	// Time when the download URL will expire
	ChatDownloadURLExpiry string                                `json:"chat_download_url_expiry" api:"required"`
	JSON                  sessionGetSessionChatResponseDataJSON `json:"-"`
}

// sessionGetSessionChatResponseDataJSON contains the JSON metadata for the struct
// [SessionGetSessionChatResponseData]
type sessionGetSessionChatResponseDataJSON struct {
	ChatDownloadURL       apijson.Field
	ChatDownloadURLExpiry apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SessionGetSessionChatResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionChatResponseDataJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionDetailsResponse struct {
	Data    SessionGetSessionDetailsResponseData `json:"data"`
	Success bool                                 `json:"success"`
	JSON    sessionGetSessionDetailsResponseJSON `json:"-"`
}

// sessionGetSessionDetailsResponseJSON contains the JSON metadata for the struct
// [SessionGetSessionDetailsResponse]
type sessionGetSessionDetailsResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetSessionDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionDetailsResponseData struct {
	// ID of the session
	ID string `json:"id" api:"required"`
	// ID of the meeting this session is associated with. In the case of V2 meetings,
	// it is always a UUID. In V1 meetings, it is a room name of the form
	// `abcdef-ghijkl`
	AssociatedID string `json:"associated_id" api:"required"`
	// timestamp when session created
	CreatedAt string `json:"created_at" api:"required"`
	// number of participants currently in the session
	LiveParticipants float64 `json:"live_participants" api:"required"`
	// number of maximum participants that were in the session
	MaxConcurrentParticipants float64 `json:"max_concurrent_participants" api:"required"`
	// Title of the meeting this session belongs to
	MeetingDisplayName string `json:"meeting_display_name" api:"required"`
	// number of minutes consumed since the session started
	MinutesConsumed float64 `json:"minutes_consumed" api:"required"`
	// App id that hosted this session
	OrganizationID string `json:"organization_id" api:"required"`
	// timestamp when session started
	StartedAt string `json:"started_at" api:"required"`
	// current status of session
	Status SessionGetSessionDetailsResponseDataStatus `json:"status" api:"required"`
	// type of session
	Type SessionGetSessionDetailsResponseDataType `json:"type" api:"required"`
	// timestamp when session was last updated
	UpdatedAt     string        `json:"updated_at" api:"required"`
	BreakoutRooms []interface{} `json:"breakout_rooms"`
	// timestamp when session ended
	EndedAt string                                   `json:"ended_at"`
	JSON    sessionGetSessionDetailsResponseDataJSON `json:"-"`
}

// sessionGetSessionDetailsResponseDataJSON contains the JSON metadata for the
// struct [SessionGetSessionDetailsResponseData]
type sessionGetSessionDetailsResponseDataJSON struct {
	ID                        apijson.Field
	AssociatedID              apijson.Field
	CreatedAt                 apijson.Field
	LiveParticipants          apijson.Field
	MaxConcurrentParticipants apijson.Field
	MeetingDisplayName        apijson.Field
	MinutesConsumed           apijson.Field
	OrganizationID            apijson.Field
	StartedAt                 apijson.Field
	Status                    apijson.Field
	Type                      apijson.Field
	UpdatedAt                 apijson.Field
	BreakoutRooms             apijson.Field
	EndedAt                   apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SessionGetSessionDetailsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionDetailsResponseDataJSON) RawJSON() string {
	return r.raw
}

// current status of session
type SessionGetSessionDetailsResponseDataStatus string

const (
	SessionGetSessionDetailsResponseDataStatusLive  SessionGetSessionDetailsResponseDataStatus = "LIVE"
	SessionGetSessionDetailsResponseDataStatusEnded SessionGetSessionDetailsResponseDataStatus = "ENDED"
)

func (r SessionGetSessionDetailsResponseDataStatus) IsKnown() bool {
	switch r {
	case SessionGetSessionDetailsResponseDataStatusLive, SessionGetSessionDetailsResponseDataStatusEnded:
		return true
	}
	return false
}

// type of session
type SessionGetSessionDetailsResponseDataType string

const (
	SessionGetSessionDetailsResponseDataTypeMeeting     SessionGetSessionDetailsResponseDataType = "meeting"
	SessionGetSessionDetailsResponseDataTypeLivestream  SessionGetSessionDetailsResponseDataType = "livestream"
	SessionGetSessionDetailsResponseDataTypeParticipant SessionGetSessionDetailsResponseDataType = "participant"
)

func (r SessionGetSessionDetailsResponseDataType) IsKnown() bool {
	switch r {
	case SessionGetSessionDetailsResponseDataTypeMeeting, SessionGetSessionDetailsResponseDataTypeLivestream, SessionGetSessionDetailsResponseDataTypeParticipant:
		return true
	}
	return false
}

type SessionGetSessionParticipantDetailsResponse struct {
	Data    SessionGetSessionParticipantDetailsResponseData `json:"data"`
	Success bool                                            `json:"success"`
	JSON    sessionGetSessionParticipantDetailsResponseJSON `json:"-"`
}

// sessionGetSessionParticipantDetailsResponseJSON contains the JSON metadata for
// the struct [SessionGetSessionParticipantDetailsResponse]
type sessionGetSessionParticipantDetailsResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetSessionParticipantDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionParticipantDetailsResponseData struct {
	Participant SessionGetSessionParticipantDetailsResponseDataParticipant `json:"participant"`
	JSON        sessionGetSessionParticipantDetailsResponseDataJSON        `json:"-"`
}

// sessionGetSessionParticipantDetailsResponseDataJSON contains the JSON metadata
// for the struct [SessionGetSessionParticipantDetailsResponseData]
type sessionGetSessionParticipantDetailsResponseDataJSON struct {
	Participant apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetSessionParticipantDetailsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantDetailsResponseDataJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionParticipantDetailsResponseDataParticipant struct {
	// Participant ID. This maps to the corresponding peerId.
	ID string `json:"id"`
	// timestamp when this participant was created.
	CreatedAt string `json:"created_at"`
	// ID passed by client to create this participant.
	CustomParticipantID string `json:"custom_participant_id"`
	// Display name of participant when joining the session.
	DisplayName string `json:"display_name"`
	// number of minutes for which the participant was in the session.
	Duration float64 `json:"duration"`
	// timestamp at which participant joined the session.
	JoinedAt string `json:"joined_at"`
	// timestamp at which participant left the session.
	LeftAt string `json:"left_at"`
	// Connection lifecycle events for the participant's peer. Only included when
	// `include_peer_events` is true.
	PeerEvents []SessionGetSessionParticipantDetailsResponseDataParticipantPeerEvent `json:"peer_events"`
	// Name of the preset associated with the participant.
	PresetName string `json:"preset_name"`
	// timestamp when this participant's data was last updated.
	UpdatedAt string `json:"updated_at"`
	// User id for this participant.
	UserID string                                                         `json:"user_id"`
	JSON   sessionGetSessionParticipantDetailsResponseDataParticipantJSON `json:"-"`
}

// sessionGetSessionParticipantDetailsResponseDataParticipantJSON contains the JSON
// metadata for the struct
// [SessionGetSessionParticipantDetailsResponseDataParticipant]
type sessionGetSessionParticipantDetailsResponseDataParticipantJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	CustomParticipantID apijson.Field
	DisplayName         apijson.Field
	Duration            apijson.Field
	JoinedAt            apijson.Field
	LeftAt              apijson.Field
	PeerEvents          apijson.Field
	PresetName          apijson.Field
	UpdatedAt           apijson.Field
	UserID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SessionGetSessionParticipantDetailsResponseDataParticipant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantDetailsResponseDataParticipantJSON) RawJSON() string {
	return r.raw
}

// A connection lifecycle event recorded for a participant's peer.
type SessionGetSessionParticipantDetailsResponseDataParticipantPeerEvent struct {
	// ID of the peer event.
	ID string `json:"id"`
	// Timestamp when this peer event was created.
	CreatedAt string `json:"created_at"`
	// Name of the peer event.
	EventName SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsEventName `json:"event_name"`
	// Minutes consumed attributed to this event.
	MinutesConsumed float64 `json:"minutes_consumed"`
	// ID of the participant this event belongs to.
	ParticipantID string `json:"participant_id" api:"nullable"`
	// Peer ID this event belongs to.
	PeerID string `json:"peer_id"`
	// View type of the preset associated with the peer.
	PresetViewType SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewType `json:"preset_view_type" api:"nullable"`
	// ID of the session this event belongs to.
	SessionID string `json:"session_id" api:"nullable"`
	// ID of the socket session associated with this event.
	SocketSessionID string `json:"socket_session_id" api:"nullable"`
	// Timestamp when this peer event was last updated.
	UpdatedAt string                                                                  `json:"updated_at"`
	JSON      sessionGetSessionParticipantDetailsResponseDataParticipantPeerEventJSON `json:"-"`
}

// sessionGetSessionParticipantDetailsResponseDataParticipantPeerEventJSON contains
// the JSON metadata for the struct
// [SessionGetSessionParticipantDetailsResponseDataParticipantPeerEvent]
type sessionGetSessionParticipantDetailsResponseDataParticipantPeerEventJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	EventName       apijson.Field
	MinutesConsumed apijson.Field
	ParticipantID   apijson.Field
	PeerID          apijson.Field
	PresetViewType  apijson.Field
	SessionID       apijson.Field
	SocketSessionID apijson.Field
	UpdatedAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SessionGetSessionParticipantDetailsResponseDataParticipantPeerEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantDetailsResponseDataParticipantPeerEventJSON) RawJSON() string {
	return r.raw
}

// Name of the peer event.
type SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsEventName string

const (
	SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsEventNamePeerCreated SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsEventName = "PEER_CREATED"
	SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsEventNamePeerJoining SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsEventName = "PEER_JOINING"
	SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsEventNamePeerLeaving SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsEventName = "PEER_LEAVING"
)

func (r SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsEventName) IsKnown() bool {
	switch r {
	case SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsEventNamePeerCreated, SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsEventNamePeerJoining, SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsEventNamePeerLeaving:
		return true
	}
	return false
}

// View type of the preset associated with the peer.
type SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewType string

const (
	SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewTypeGroupCall  SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewType = "GROUP_CALL"
	SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewTypeWebinar    SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewType = "WEBINAR"
	SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewTypeAudioRoom  SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewType = "AUDIO_ROOM"
	SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewTypeLivestream SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewType = "LIVESTREAM"
	SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewTypeChat       SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewType = "CHAT"
)

func (r SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewType) IsKnown() bool {
	switch r {
	case SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewTypeGroupCall, SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewTypeWebinar, SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewTypeAudioRoom, SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewTypeLivestream, SessionGetSessionParticipantDetailsResponseDataParticipantPeerEventsPresetViewTypeChat:
		return true
	}
	return false
}

type SessionGetSessionParticipantsResponse struct {
	Data    SessionGetSessionParticipantsResponseData `json:"data"`
	Success bool                                      `json:"success"`
	JSON    sessionGetSessionParticipantsResponseJSON `json:"-"`
}

// sessionGetSessionParticipantsResponseJSON contains the JSON metadata for the
// struct [SessionGetSessionParticipantsResponse]
type sessionGetSessionParticipantsResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetSessionParticipantsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantsResponseJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionParticipantsResponseData struct {
	Participants []SessionGetSessionParticipantsResponseDataParticipant `json:"participants"`
	JSON         sessionGetSessionParticipantsResponseDataJSON          `json:"-"`
}

// sessionGetSessionParticipantsResponseDataJSON contains the JSON metadata for the
// struct [SessionGetSessionParticipantsResponseData]
type sessionGetSessionParticipantsResponseDataJSON struct {
	Participants apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SessionGetSessionParticipantsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantsResponseDataJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionParticipantsResponseDataParticipant struct {
	// Participant ID. This maps to the corresponding peerId.
	ID string `json:"id"`
	// timestamp when this participant was created.
	CreatedAt string `json:"created_at"`
	// ID passed by client to create this participant.
	CustomParticipantID string `json:"custom_participant_id"`
	// Display name of participant when joining the session.
	DisplayName string `json:"display_name"`
	// number of minutes for which the participant was in the session.
	Duration float64 `json:"duration"`
	// timestamp at which participant joined the session.
	JoinedAt string `json:"joined_at"`
	// timestamp at which participant left the session.
	LeftAt string `json:"left_at"`
	// Connection lifecycle events for the participant's peer. Only included when
	// `include_peer_events` is true.
	PeerEvents []SessionGetSessionParticipantsResponseDataParticipantsPeerEvent `json:"peer_events"`
	// Name of the preset associated with the participant.
	PresetName string `json:"preset_name"`
	// timestamp when this participant's data was last updated.
	UpdatedAt string `json:"updated_at"`
	// User id for this participant.
	UserID string                                                   `json:"user_id"`
	JSON   sessionGetSessionParticipantsResponseDataParticipantJSON `json:"-"`
}

// sessionGetSessionParticipantsResponseDataParticipantJSON contains the JSON
// metadata for the struct [SessionGetSessionParticipantsResponseDataParticipant]
type sessionGetSessionParticipantsResponseDataParticipantJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	CustomParticipantID apijson.Field
	DisplayName         apijson.Field
	Duration            apijson.Field
	JoinedAt            apijson.Field
	LeftAt              apijson.Field
	PeerEvents          apijson.Field
	PresetName          apijson.Field
	UpdatedAt           apijson.Field
	UserID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SessionGetSessionParticipantsResponseDataParticipant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantsResponseDataParticipantJSON) RawJSON() string {
	return r.raw
}

// A connection lifecycle event recorded for a participant's peer.
type SessionGetSessionParticipantsResponseDataParticipantsPeerEvent struct {
	// ID of the peer event.
	ID string `json:"id"`
	// Timestamp when this peer event was created.
	CreatedAt string `json:"created_at"`
	// Name of the peer event.
	EventName SessionGetSessionParticipantsResponseDataParticipantsPeerEventsEventName `json:"event_name"`
	// Minutes consumed attributed to this event.
	MinutesConsumed float64 `json:"minutes_consumed"`
	// ID of the participant this event belongs to.
	ParticipantID string `json:"participant_id" api:"nullable"`
	// Peer ID this event belongs to.
	PeerID string `json:"peer_id"`
	// View type of the preset associated with the peer.
	PresetViewType SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewType `json:"preset_view_type" api:"nullable"`
	// ID of the session this event belongs to.
	SessionID string `json:"session_id" api:"nullable"`
	// ID of the socket session associated with this event.
	SocketSessionID string `json:"socket_session_id" api:"nullable"`
	// Timestamp when this peer event was last updated.
	UpdatedAt string                                                             `json:"updated_at"`
	JSON      sessionGetSessionParticipantsResponseDataParticipantsPeerEventJSON `json:"-"`
}

// sessionGetSessionParticipantsResponseDataParticipantsPeerEventJSON contains the
// JSON metadata for the struct
// [SessionGetSessionParticipantsResponseDataParticipantsPeerEvent]
type sessionGetSessionParticipantsResponseDataParticipantsPeerEventJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	EventName       apijson.Field
	MinutesConsumed apijson.Field
	ParticipantID   apijson.Field
	PeerID          apijson.Field
	PresetViewType  apijson.Field
	SessionID       apijson.Field
	SocketSessionID apijson.Field
	UpdatedAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SessionGetSessionParticipantsResponseDataParticipantsPeerEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantsResponseDataParticipantsPeerEventJSON) RawJSON() string {
	return r.raw
}

// Name of the peer event.
type SessionGetSessionParticipantsResponseDataParticipantsPeerEventsEventName string

const (
	SessionGetSessionParticipantsResponseDataParticipantsPeerEventsEventNamePeerCreated SessionGetSessionParticipantsResponseDataParticipantsPeerEventsEventName = "PEER_CREATED"
	SessionGetSessionParticipantsResponseDataParticipantsPeerEventsEventNamePeerJoining SessionGetSessionParticipantsResponseDataParticipantsPeerEventsEventName = "PEER_JOINING"
	SessionGetSessionParticipantsResponseDataParticipantsPeerEventsEventNamePeerLeaving SessionGetSessionParticipantsResponseDataParticipantsPeerEventsEventName = "PEER_LEAVING"
)

func (r SessionGetSessionParticipantsResponseDataParticipantsPeerEventsEventName) IsKnown() bool {
	switch r {
	case SessionGetSessionParticipantsResponseDataParticipantsPeerEventsEventNamePeerCreated, SessionGetSessionParticipantsResponseDataParticipantsPeerEventsEventNamePeerJoining, SessionGetSessionParticipantsResponseDataParticipantsPeerEventsEventNamePeerLeaving:
		return true
	}
	return false
}

// View type of the preset associated with the peer.
type SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewType string

const (
	SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewTypeGroupCall  SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewType = "GROUP_CALL"
	SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewTypeWebinar    SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewType = "WEBINAR"
	SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewTypeAudioRoom  SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewType = "AUDIO_ROOM"
	SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewTypeLivestream SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewType = "LIVESTREAM"
	SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewTypeChat       SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewType = "CHAT"
)

func (r SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewType) IsKnown() bool {
	switch r {
	case SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewTypeGroupCall, SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewTypeWebinar, SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewTypeAudioRoom, SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewTypeLivestream, SessionGetSessionParticipantsResponseDataParticipantsPeerEventsPresetViewTypeChat:
		return true
	}
	return false
}

type SessionGetSessionSummaryResponse struct {
	Data    SessionGetSessionSummaryResponseData `json:"data"`
	Success bool                                 `json:"success"`
	JSON    sessionGetSessionSummaryResponseJSON `json:"-"`
}

// sessionGetSessionSummaryResponseJSON contains the JSON metadata for the struct
// [SessionGetSessionSummaryResponse]
type sessionGetSessionSummaryResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetSessionSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionSummaryResponseData struct {
	SessionID string `json:"sessionId" api:"required"`
	// URL where the summary of transcripts can be downloaded
	SummaryDownloadURL string `json:"summaryDownloadUrl" api:"required"`
	// Time of Expiry before when you need to download the csv file.
	SummaryDownloadURLExpiry string                                   `json:"summaryDownloadUrlExpiry" api:"required"`
	JSON                     sessionGetSessionSummaryResponseDataJSON `json:"-"`
}

// sessionGetSessionSummaryResponseDataJSON contains the JSON metadata for the
// struct [SessionGetSessionSummaryResponseData]
type sessionGetSessionSummaryResponseDataJSON struct {
	SessionID                apijson.Field
	SummaryDownloadURL       apijson.Field
	SummaryDownloadURLExpiry apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SessionGetSessionSummaryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionSummaryResponseDataJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionTranscriptsResponse struct {
	Data    SessionGetSessionTranscriptsResponseData `json:"data"`
	Success bool                                     `json:"success"`
	JSON    sessionGetSessionTranscriptsResponseJSON `json:"-"`
}

// sessionGetSessionTranscriptsResponseJSON contains the JSON metadata for the
// struct [SessionGetSessionTranscriptsResponse]
type sessionGetSessionTranscriptsResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetSessionTranscriptsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionTranscriptsResponseJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionTranscriptsResponseData struct {
	SessionID string `json:"sessionId" api:"required"`
	// URL where the transcript can be downloaded
	TranscriptDownloadURL string `json:"transcript_download_url" api:"required"`
	// Time when the download URL will expire
	TranscriptDownloadURLExpiry string                                       `json:"transcript_download_url_expiry" api:"required"`
	JSON                        sessionGetSessionTranscriptsResponseDataJSON `json:"-"`
}

// sessionGetSessionTranscriptsResponseDataJSON contains the JSON metadata for the
// struct [SessionGetSessionTranscriptsResponseData]
type sessionGetSessionTranscriptsResponseDataJSON struct {
	SessionID                   apijson.Field
	TranscriptDownloadURL       apijson.Field
	TranscriptDownloadURLExpiry apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *SessionGetSessionTranscriptsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionTranscriptsResponseDataJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionsResponse struct {
	Data    SessionGetSessionsResponseData   `json:"data"`
	Paging  SessionGetSessionsResponsePaging `json:"paging"`
	Success bool                             `json:"success"`
	JSON    sessionGetSessionsResponseJSON   `json:"-"`
}

// sessionGetSessionsResponseJSON contains the JSON metadata for the struct
// [SessionGetSessionsResponse]
type sessionGetSessionsResponseJSON struct {
	Data        apijson.Field
	Paging      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetSessionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionsResponseJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionsResponseData struct {
	Sessions []SessionGetSessionsResponseDataSession `json:"sessions"`
	JSON     sessionGetSessionsResponseDataJSON      `json:"-"`
}

// sessionGetSessionsResponseDataJSON contains the JSON metadata for the struct
// [SessionGetSessionsResponseData]
type sessionGetSessionsResponseDataJSON struct {
	Sessions    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetSessionsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionsResponseDataJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionsResponseDataSession struct {
	// ID of the session
	ID string `json:"id" api:"required"`
	// ID of the meeting this session is associated with. In the case of V2 meetings,
	// it is always a UUID. In V1 meetings, it is a room name of the form
	// `abcdef-ghijkl`
	AssociatedID string `json:"associated_id" api:"required"`
	// timestamp when session created
	CreatedAt string `json:"created_at" api:"required"`
	// number of participants currently in the session
	LiveParticipants float64 `json:"live_participants" api:"required"`
	// number of maximum participants that were in the session
	MaxConcurrentParticipants float64 `json:"max_concurrent_participants" api:"required"`
	// Title of the meeting this session belongs to
	MeetingDisplayName string `json:"meeting_display_name" api:"required"`
	// number of minutes consumed since the session started
	MinutesConsumed float64 `json:"minutes_consumed" api:"required"`
	// App id that hosted this session
	OrganizationID string `json:"organization_id" api:"required"`
	// timestamp when session started
	StartedAt string `json:"started_at" api:"required"`
	// current status of session
	Status SessionGetSessionsResponseDataSessionsStatus `json:"status" api:"required"`
	// type of session
	Type SessionGetSessionsResponseDataSessionsType `json:"type" api:"required"`
	// timestamp when session was last updated
	UpdatedAt     string        `json:"updated_at" api:"required"`
	BreakoutRooms []interface{} `json:"breakout_rooms"`
	// timestamp when session ended
	EndedAt string                                    `json:"ended_at"`
	JSON    sessionGetSessionsResponseDataSessionJSON `json:"-"`
}

// sessionGetSessionsResponseDataSessionJSON contains the JSON metadata for the
// struct [SessionGetSessionsResponseDataSession]
type sessionGetSessionsResponseDataSessionJSON struct {
	ID                        apijson.Field
	AssociatedID              apijson.Field
	CreatedAt                 apijson.Field
	LiveParticipants          apijson.Field
	MaxConcurrentParticipants apijson.Field
	MeetingDisplayName        apijson.Field
	MinutesConsumed           apijson.Field
	OrganizationID            apijson.Field
	StartedAt                 apijson.Field
	Status                    apijson.Field
	Type                      apijson.Field
	UpdatedAt                 apijson.Field
	BreakoutRooms             apijson.Field
	EndedAt                   apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SessionGetSessionsResponseDataSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionsResponseDataSessionJSON) RawJSON() string {
	return r.raw
}

// current status of session
type SessionGetSessionsResponseDataSessionsStatus string

const (
	SessionGetSessionsResponseDataSessionsStatusLive  SessionGetSessionsResponseDataSessionsStatus = "LIVE"
	SessionGetSessionsResponseDataSessionsStatusEnded SessionGetSessionsResponseDataSessionsStatus = "ENDED"
)

func (r SessionGetSessionsResponseDataSessionsStatus) IsKnown() bool {
	switch r {
	case SessionGetSessionsResponseDataSessionsStatusLive, SessionGetSessionsResponseDataSessionsStatusEnded:
		return true
	}
	return false
}

// type of session
type SessionGetSessionsResponseDataSessionsType string

const (
	SessionGetSessionsResponseDataSessionsTypeMeeting     SessionGetSessionsResponseDataSessionsType = "meeting"
	SessionGetSessionsResponseDataSessionsTypeLivestream  SessionGetSessionsResponseDataSessionsType = "livestream"
	SessionGetSessionsResponseDataSessionsTypeParticipant SessionGetSessionsResponseDataSessionsType = "participant"
)

func (r SessionGetSessionsResponseDataSessionsType) IsKnown() bool {
	switch r {
	case SessionGetSessionsResponseDataSessionsTypeMeeting, SessionGetSessionsResponseDataSessionsTypeLivestream, SessionGetSessionsResponseDataSessionsTypeParticipant:
		return true
	}
	return false
}

type SessionGetSessionsResponsePaging struct {
	EndOffset   float64                              `json:"end_offset"`
	StartOffset float64                              `json:"start_offset"`
	TotalCount  float64                              `json:"total_count"`
	JSON        sessionGetSessionsResponsePagingJSON `json:"-"`
}

// sessionGetSessionsResponsePagingJSON contains the JSON metadata for the struct
// [SessionGetSessionsResponsePaging]
type sessionGetSessionsResponsePagingJSON struct {
	EndOffset   apijson.Field
	StartOffset apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetSessionsResponsePaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionsResponsePagingJSON) RawJSON() string {
	return r.raw
}

type SessionGenerateSummaryOfTranscriptsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SessionGetParticipantDataFromPeerIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter to apply to the peer report.
	Filters param.Field[SessionGetParticipantDataFromPeerIDParamsFilters] `query:"filters"`
	// if true, response includes all the peer events of participant.
	IncludePeerEvents param.Field[bool] `query:"include_peer_events"`
}

// URLQuery serializes [SessionGetParticipantDataFromPeerIDParams]'s query
// parameters as `url.Values`.
func (r SessionGetParticipantDataFromPeerIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter to apply to the peer report.
type SessionGetParticipantDataFromPeerIDParamsFilters string

const (
	SessionGetParticipantDataFromPeerIDParamsFiltersDeviceInfo                SessionGetParticipantDataFromPeerIDParamsFilters = "device_info"
	SessionGetParticipantDataFromPeerIDParamsFiltersIPInformation             SessionGetParticipantDataFromPeerIDParamsFilters = "ip_information"
	SessionGetParticipantDataFromPeerIDParamsFiltersPrecallNetworkInformation SessionGetParticipantDataFromPeerIDParamsFilters = "precall_network_information"
	SessionGetParticipantDataFromPeerIDParamsFiltersEvents                    SessionGetParticipantDataFromPeerIDParamsFilters = "events"
	SessionGetParticipantDataFromPeerIDParamsFiltersQualityStats              SessionGetParticipantDataFromPeerIDParamsFilters = "quality_stats"
)

func (r SessionGetParticipantDataFromPeerIDParamsFilters) IsKnown() bool {
	switch r {
	case SessionGetParticipantDataFromPeerIDParamsFiltersDeviceInfo, SessionGetParticipantDataFromPeerIDParamsFiltersIPInformation, SessionGetParticipantDataFromPeerIDParamsFiltersPrecallNetworkInformation, SessionGetParticipantDataFromPeerIDParamsFiltersEvents, SessionGetParticipantDataFromPeerIDParamsFiltersQualityStats:
		return true
	}
	return false
}

type SessionGetSessionChatParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SessionGetSessionDetailsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// List all breakout rooms
	IncludeBreakoutRooms param.Field[bool] `query:"include_breakout_rooms"`
}

// URLQuery serializes [SessionGetSessionDetailsParams]'s query parameters as
// `url.Values`.
func (r SessionGetSessionDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SessionGetSessionParticipantDetailsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// if true, response includes all the peer events of participant.
	IncludePeerEvents param.Field[bool] `query:"include_peer_events"`
}

// URLQuery serializes [SessionGetSessionParticipantDetailsParams]'s query
// parameters as `url.Values`.
func (r SessionGetSessionParticipantDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SessionGetSessionParticipantsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// if true, response includes all the peer events of participants.
	IncludePeerEvents param.Field[bool] `query:"include_peer_events"`
	// The page number from which you want your page search results to be displayed.
	PageNo param.Field[float64] `query:"page_no"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// The search query string. You can search using participant ID, custom participant
	// ID, or display name.
	Search    param.Field[string]                                       `query:"search"`
	SortBy    param.Field[SessionGetSessionParticipantsParamsSortBy]    `query:"sort_by"`
	SortOrder param.Field[SessionGetSessionParticipantsParamsSortOrder] `query:"sort_order"`
	// In breakout room sessions, the view parameter can be set to `raw` for session
	// specific duration for participants or `consolidated` to accumulate breakout room
	// durations.
	View param.Field[SessionGetSessionParticipantsParamsView] `query:"view"`
}

// URLQuery serializes [SessionGetSessionParticipantsParams]'s query parameters as
// `url.Values`.
func (r SessionGetSessionParticipantsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SessionGetSessionParticipantsParamsSortBy string

const (
	SessionGetSessionParticipantsParamsSortByJoinedAt SessionGetSessionParticipantsParamsSortBy = "joinedAt"
	SessionGetSessionParticipantsParamsSortByDuration SessionGetSessionParticipantsParamsSortBy = "duration"
)

func (r SessionGetSessionParticipantsParamsSortBy) IsKnown() bool {
	switch r {
	case SessionGetSessionParticipantsParamsSortByJoinedAt, SessionGetSessionParticipantsParamsSortByDuration:
		return true
	}
	return false
}

type SessionGetSessionParticipantsParamsSortOrder string

const (
	SessionGetSessionParticipantsParamsSortOrderAsc  SessionGetSessionParticipantsParamsSortOrder = "ASC"
	SessionGetSessionParticipantsParamsSortOrderDesc SessionGetSessionParticipantsParamsSortOrder = "DESC"
)

func (r SessionGetSessionParticipantsParamsSortOrder) IsKnown() bool {
	switch r {
	case SessionGetSessionParticipantsParamsSortOrderAsc, SessionGetSessionParticipantsParamsSortOrderDesc:
		return true
	}
	return false
}

// In breakout room sessions, the view parameter can be set to `raw` for session
// specific duration for participants or `consolidated` to accumulate breakout room
// durations.
type SessionGetSessionParticipantsParamsView string

const (
	SessionGetSessionParticipantsParamsViewRaw          SessionGetSessionParticipantsParamsView = "raw"
	SessionGetSessionParticipantsParamsViewConsolidated SessionGetSessionParticipantsParamsView = "consolidated"
)

func (r SessionGetSessionParticipantsParamsView) IsKnown() bool {
	switch r {
	case SessionGetSessionParticipantsParamsViewRaw, SessionGetSessionParticipantsParamsViewConsolidated:
		return true
	}
	return false
}

type SessionGetSessionSummaryParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SessionGetSessionTranscriptsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Transcript file format to fetch.
	Format param.Field[SessionGetSessionTranscriptsParamsFormat] `query:"format"`
}

// URLQuery serializes [SessionGetSessionTranscriptsParams]'s query parameters as
// `url.Values`.
func (r SessionGetSessionTranscriptsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Transcript file format to fetch.
type SessionGetSessionTranscriptsParamsFormat string

const (
	SessionGetSessionTranscriptsParamsFormatSrt  SessionGetSessionTranscriptsParamsFormat = "SRT"
	SessionGetSessionTranscriptsParamsFormatVtt  SessionGetSessionTranscriptsParamsFormat = "VTT"
	SessionGetSessionTranscriptsParamsFormatJson SessionGetSessionTranscriptsParamsFormat = "JSON"
	SessionGetSessionTranscriptsParamsFormatCsv  SessionGetSessionTranscriptsParamsFormat = "CSV"
)

func (r SessionGetSessionTranscriptsParamsFormat) IsKnown() bool {
	switch r {
	case SessionGetSessionTranscriptsParamsFormatSrt, SessionGetSessionTranscriptsParamsFormatVtt, SessionGetSessionTranscriptsParamsFormatJson, SessionGetSessionTranscriptsParamsFormatCsv:
		return true
	}
	return false
}

type SessionGetSessionsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// ID of the meeting that sessions should be associated with
	AssociatedID param.Field[string] `query:"associated_id" format:"uuid"`
	// The end time range for which you want to retrieve the meetings. The time must be
	// specified in ISO format.
	EndTime param.Field[time.Time] `query:"end_time" format:"date-time"`
	// The page number from which you want your page search results to be displayed.
	PageNo       param.Field[float64] `query:"page_no"`
	Participants param.Field[string]  `query:"participants"`
	// Number of results per page
	PerPage param.Field[float64] `query:"per_page"`
	// Search string that matches sessions based on meeting title, meeting ID, and
	// session ID
	Search    param.Field[string]                            `query:"search"`
	SortBy    param.Field[SessionGetSessionsParamsSortBy]    `query:"sort_by"`
	SortOrder param.Field[SessionGetSessionsParamsSortOrder] `query:"sort_order"`
	// The start time range for which you want to retrieve the meetings. The time must
	// be specified in ISO format.
	StartTime param.Field[time.Time]                      `query:"start_time" format:"date-time"`
	Status    param.Field[SessionGetSessionsParamsStatus] `query:"status"`
}

// URLQuery serializes [SessionGetSessionsParams]'s query parameters as
// `url.Values`.
func (r SessionGetSessionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SessionGetSessionsParamsSortBy string

const (
	SessionGetSessionsParamsSortByMinutesConsumed SessionGetSessionsParamsSortBy = "minutesConsumed"
	SessionGetSessionsParamsSortByCreatedAt       SessionGetSessionsParamsSortBy = "createdAt"
)

func (r SessionGetSessionsParamsSortBy) IsKnown() bool {
	switch r {
	case SessionGetSessionsParamsSortByMinutesConsumed, SessionGetSessionsParamsSortByCreatedAt:
		return true
	}
	return false
}

type SessionGetSessionsParamsSortOrder string

const (
	SessionGetSessionsParamsSortOrderAsc  SessionGetSessionsParamsSortOrder = "ASC"
	SessionGetSessionsParamsSortOrderDesc SessionGetSessionsParamsSortOrder = "DESC"
)

func (r SessionGetSessionsParamsSortOrder) IsKnown() bool {
	switch r {
	case SessionGetSessionsParamsSortOrderAsc, SessionGetSessionsParamsSortOrderDesc:
		return true
	}
	return false
}

type SessionGetSessionsParamsStatus string

const (
	SessionGetSessionsParamsStatusLive  SessionGetSessionsParamsStatus = "LIVE"
	SessionGetSessionsParamsStatusEnded SessionGetSessionsParamsStatus = "ENDED"
)

func (r SessionGetSessionsParamsStatus) IsKnown() bool {
	switch r {
	case SessionGetSessionsParamsStatusLive, SessionGetSessionsParamsStatusEnded:
		return true
	}
	return false
}
