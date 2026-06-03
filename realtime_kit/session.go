// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package realtime_kit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
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

// Returns details of the given peer ID along with call statistics for the given
// session ID.
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

// Returns details of the given participant ID along with call statistics for the
// given session ID.
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
	LeftAt     string                   `json:"left_at"`
	PeerEvents []map[string]interface{} `json:"peer_events"`
	// Peer call statistics report.
	PeerReport SessionGetParticipantDataFromPeerIDResponseDataPeerReport `json:"peer_report"`
	// Name of the preset associated with the participant.
	PresetName string `json:"preset_name"`
	SessionID  string `json:"session_id" format:"uuid"`
	// timestamp when this participant's data was last updated.
	UpdatedAt string `json:"updated_at"`
	// User id for this participant.
	UserID string                                              `json:"user_id"`
	JSON   sessionGetParticipantDataFromPeerIDResponseDataJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataJSON contains the JSON metadata
// for the struct [SessionGetParticipantDataFromPeerIDResponseData]
type sessionGetParticipantDataFromPeerIDResponseDataJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	CustomParticipantID apijson.Field
	DisplayName         apijson.Field
	Duration            apijson.Field
	JoinedAt            apijson.Field
	LeftAt              apijson.Field
	PeerEvents          apijson.Field
	PeerReport          apijson.Field
	PresetName          apijson.Field
	SessionID           apijson.Field
	UpdatedAt           apijson.Field
	UserID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataJSON) RawJSON() string {
	return r.raw
}

// Peer call statistics report.
type SessionGetParticipantDataFromPeerIDResponseDataPeerReport struct {
	Metadata    map[string]interface{}                                        `json:"metadata"`
	Quality     map[string]interface{}                                        `json:"quality"`
	ExtraFields map[string]interface{}                                        `json:"-" api:"extrafields"`
	JSON        sessionGetParticipantDataFromPeerIDResponseDataPeerReportJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataPeerReportJSON contains the JSON
// metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataPeerReport]
type sessionGetParticipantDataFromPeerIDResponseDataPeerReportJSON struct {
	Metadata    apijson.Field
	Quality     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataPeerReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataPeerReportJSON) RawJSON() string {
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
	EndedAt string `json:"ended_at"`
	// Any meta data about session.
	Meta interface{}                              `json:"meta"`
	JSON sessionGetSessionDetailsResponseDataJSON `json:"-"`
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
	Meta                      apijson.Field
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
	EndedAt string `json:"ended_at"`
	// Any meta data about session.
	Meta interface{}                               `json:"meta"`
	JSON sessionGetSessionsResponseDataSessionJSON `json:"-"`
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
	Meta                      apijson.Field
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
	// Comma separated list of filters to apply. Note that there must be no spaces
	// between the filters.
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

// Comma separated list of filters to apply. Note that there must be no spaces
// between the filters.
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
	// Comma separated list of filters to apply. Note that there must be no spaces
	// between the filters.
	Filters param.Field[SessionGetSessionParticipantDetailsParamsFilters] `query:"filters"`
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

// Comma separated list of filters to apply. Note that there must be no spaces
// between the filters.
type SessionGetSessionParticipantDetailsParamsFilters string

const (
	SessionGetSessionParticipantDetailsParamsFiltersDeviceInfo                SessionGetSessionParticipantDetailsParamsFilters = "device_info"
	SessionGetSessionParticipantDetailsParamsFiltersIPInformation             SessionGetSessionParticipantDetailsParamsFilters = "ip_information"
	SessionGetSessionParticipantDetailsParamsFiltersPrecallNetworkInformation SessionGetSessionParticipantDetailsParamsFilters = "precall_network_information"
	SessionGetSessionParticipantDetailsParamsFiltersEvents                    SessionGetSessionParticipantDetailsParamsFilters = "events"
	SessionGetSessionParticipantDetailsParamsFiltersQualityStats              SessionGetSessionParticipantDetailsParamsFilters = "quality_stats"
)

func (r SessionGetSessionParticipantDetailsParamsFilters) IsKnown() bool {
	switch r {
	case SessionGetSessionParticipantDetailsParamsFiltersDeviceInfo, SessionGetSessionParticipantDetailsParamsFiltersIPInformation, SessionGetSessionParticipantDetailsParamsFiltersPrecallNetworkInformation, SessionGetSessionParticipantDetailsParamsFiltersEvents, SessionGetSessionParticipantDetailsParamsFiltersQualityStats:
		return true
	}
	return false
}

type SessionGetSessionParticipantsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// if true, response includes all the peer events of participants.
	IncludePeerEvents param.Field[bool] `query:"include_peer_events"`
	// The page number from which you want your page search results to be displayed.
	PageNo param.Field[float64] `query:"page_no"`
	// Number of results per page
	PerPage param.Field[float64] `query:"per_page"`
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
