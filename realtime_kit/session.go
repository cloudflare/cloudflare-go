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

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
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
func (r *SessionService) GenerateSummaryOfTranscripts(ctx context.Context, appID string, sessionID string, body SessionGenerateSummaryOfTranscriptsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s/summary", body.AccountID, appID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Returns details of the given peer ID along with call statistics for the given
// session ID.
func (r *SessionService) GetParticipantDataFromPeerID(ctx context.Context, appID string, peerID string, params SessionGetParticipantDataFromPeerIDParams, opts ...option.RequestOption) (res *SessionGetParticipantDataFromPeerIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if peerID == "" {
		err = errors.New("missing required peer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/peer-report/%s", params.AccountID, appID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Returns a URL to download all chat messages of the session ID in CSV format.
func (r *SessionService) GetSessionChat(ctx context.Context, appID string, sessionID string, query SessionGetSessionChatParams, opts ...option.RequestOption) (res *SessionGetSessionChatResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s/chat", query.AccountID, appID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns data of the given session ID including recording details.
func (r *SessionService) GetSessionDetails(ctx context.Context, appID string, sessionID string, params SessionGetSessionDetailsParams, opts ...option.RequestOption) (res *SessionGetSessionDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s", params.AccountID, appID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Returns details of the given participant ID along with call statistics for the
// given session ID.
func (r *SessionService) GetSessionParticipantDetails(ctx context.Context, appID string, sessionID string, participantID string, params SessionGetSessionParticipantDetailsParams, opts ...option.RequestOption) (res *SessionGetSessionParticipantDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	if participantID == "" {
		err = errors.New("missing required participant_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s/participants/%s", params.AccountID, appID, sessionID, participantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Returns a list of participants for the given session ID.
func (r *SessionService) GetSessionParticipants(ctx context.Context, appID string, sessionID string, params SessionGetSessionParticipantsParams, opts ...option.RequestOption) (res *SessionGetSessionParticipantsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s/participants", params.AccountID, appID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Returns a Summary URL to download the Summary of Transcripts for the session ID
// as plain text.
func (r *SessionService) GetSessionSummary(ctx context.Context, appID string, sessionID string, query SessionGetSessionSummaryParams, opts ...option.RequestOption) (res *SessionGetSessionSummaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s/summary", query.AccountID, appID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a URL to download the transcript for the session ID in CSV format.
func (r *SessionService) GetSessionTranscripts(ctx context.Context, appID string, sessionID string, query SessionGetSessionTranscriptsParams, opts ...option.RequestOption) (res *SessionGetSessionTranscriptsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions/%s/transcript", query.AccountID, appID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns details of all sessions of an App.
func (r *SessionService) GetSessions(ctx context.Context, appID string, params SessionGetSessionsParams, opts ...option.RequestOption) (res *SessionGetSessionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/sessions", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
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
	ID                  string                                                                 `json:"id"`
	CreatedAt           string                                                                 `json:"created_at"`
	CustomParticipantID string                                                                 `json:"custom_participant_id"`
	DisplayName         string                                                                 `json:"display_name"`
	Duration            float64                                                                `json:"duration"`
	JoinedAt            string                                                                 `json:"joined_at"`
	LeftAt              string                                                                 `json:"left_at"`
	PeerReport          SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReport   `json:"peer_report"`
	PeerStats           SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStats    `json:"peer_stats"`
	QualityStats        SessionGetParticipantDataFromPeerIDResponseDataParticipantQualityStats `json:"quality_stats"`
	Role                string                                                                 `json:"role"`
	UpdatedAt           string                                                                 `json:"updated_at"`
	UserID              string                                                                 `json:"user_id"`
	JSON                sessionGetParticipantDataFromPeerIDResponseDataParticipantJSON         `json:"-"`
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
	PeerReport          apijson.Field
	PeerStats           apijson.Field
	QualityStats        apijson.Field
	Role                apijson.Field
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

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReport struct {
	Metadata SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadata `json:"metadata"`
	Quality  SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQuality  `json:"quality"`
	JSON     sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportJSON     `json:"-"`
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

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadata struct {
	AudioDevicesUpdates   []interface{}                                                                               `json:"audio_devices_updates"`
	BrowserMetadata       SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataBrowserMetadata `json:"browser_metadata"`
	CandidatePairs        SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairs  `json:"candidate_pairs"`
	DeviceInfo            SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataDeviceInfo      `json:"device_info"`
	Events                []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEvent         `json:"events"`
	IPInformation         SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformation   `json:"ip_information"`
	PcMetadata            []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataPcMetadata    `json:"pc_metadata"`
	RoomViewType          string                                                                                      `json:"room_view_type"`
	SDKName               string                                                                                      `json:"sdk_name"`
	SDKVersion            string                                                                                      `json:"sdk_version"`
	SelectedDeviceUpdates []interface{}                                                                               `json:"selected_device_updates"`
	SpeakerDevicesUpdates []interface{}                                                                               `json:"speaker_devices_updates"`
	VideoDevicesUpdates   []interface{}                                                                               `json:"video_devices_updates"`
	JSON                  sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataJSON            `json:"-"`
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
	PcMetadata            apijson.Field
	RoomViewType          apijson.Field
	SDKName               apijson.Field
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

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataBrowserMetadata struct {
	Browser        string                                                                                          `json:"browser"`
	BrowserVersion string                                                                                          `json:"browser_version"`
	Engine         string                                                                                          `json:"engine"`
	UserAgent      string                                                                                          `json:"user_agent"`
	WebglSupport   string                                                                                          `json:"webgl_support"`
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
	ConsumingTransport []interface{}                                                                                                  `json:"consuming_transport"`
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

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsProducingTransport struct {
	AvailableOutgoingBitrate     int64                                                                                                            `json:"available_outgoing_bitrate"`
	BytesDiscardedOnSend         int64                                                                                                            `json:"bytes_discarded_on_send"`
	BytesReceived                int64                                                                                                            `json:"bytes_received"`
	BytesSent                    int64                                                                                                            `json:"bytes_sent"`
	CurrentRoundTripTime         float64                                                                                                          `json:"current_round_trip_time"`
	LastPacketReceivedTimestamp  int64                                                                                                            `json:"last_packet_received_timestamp"`
	LastPacketSentTimestamp      int64                                                                                                            `json:"last_packet_sent_timestamp"`
	LocalCandidateAddress        string                                                                                                           `json:"local_candidate_address"`
	LocalCandidateID             string                                                                                                           `json:"local_candidate_id"`
	LocalCandidateNetworkType    string                                                                                                           `json:"local_candidate_network_type"`
	LocalCandidatePort           int64                                                                                                            `json:"local_candidate_port"`
	LocalCandidateProtocol       string                                                                                                           `json:"local_candidate_protocol"`
	LocalCandidateRelatedAddress string                                                                                                           `json:"local_candidate_related_address"`
	LocalCandidateRelatedPort    int64                                                                                                            `json:"local_candidate_related_port"`
	LocalCandidateType           string                                                                                                           `json:"local_candidate_type"`
	Nominated                    bool                                                                                                             `json:"nominated"`
	PacketsDiscardedOnSend       int64                                                                                                            `json:"packets_discarded_on_send"`
	PacketsReceived              int64                                                                                                            `json:"packets_received"`
	PacketsSent                  int64                                                                                                            `json:"packets_sent"`
	RemoteCandidateAddress       string                                                                                                           `json:"remote_candidate_address"`
	RemoteCandidateID            string                                                                                                           `json:"remote_candidate_id"`
	RemoteCandidatePort          int64                                                                                                            `json:"remote_candidate_port"`
	RemoteCandidateProtocol      string                                                                                                           `json:"remote_candidate_protocol"`
	RemoteCandidateType          string                                                                                                           `json:"remote_candidate_type"`
	TotalRoundTripTime           float64                                                                                                          `json:"total_round_trip_time"`
	JSON                         sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsProducingTransportJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsProducingTransportJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsProducingTransport]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataCandidatePairsProducingTransportJSON struct {
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
	Nominated                    apijson.Field
	PacketsDiscardedOnSend       apijson.Field
	PacketsReceived              apijson.Field
	PacketsSent                  apijson.Field
	RemoteCandidateAddress       apijson.Field
	RemoteCandidateID            apijson.Field
	RemoteCandidatePort          apijson.Field
	RemoteCandidateProtocol      apijson.Field
	RemoteCandidateType          apijson.Field
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
	CPUs      int64                                                                                      `json:"cpus"`
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

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEvent struct {
	Name      string                                                                                `json:"name"`
	Timestamp string                                                                                `json:"timestamp"`
	JSON      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEvent]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataEventJSON struct {
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

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformation struct {
	ASN      SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASN  `json:"asn"`
	City     string                                                                                        `json:"city"`
	Country  string                                                                                        `json:"country"`
	IPV4     string                                                                                        `json:"ipv4"`
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
	ASN  string                                                                                           `json:"asn"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASNJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASNJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASN]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASNJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataIPInformationASNJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportMetadataPcMetadata struct {
	EffectiveNetworkType  string                                                                                     `json:"effective_network_type"`
	ReflexiveConnectivity bool                                                                                       `json:"reflexive_connectivity"`
	RelayConnectivity     bool                                                                                       `json:"relay_connectivity"`
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

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQuality struct {
	AudioConsumer                      []interface{}                                                                                      `json:"audio_consumer"`
	AudioConsumerCumulative            interface{}                                                                                        `json:"audio_consumer_cumulative"`
	AudioProducer                      []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducer         `json:"audio_producer"`
	AudioProducerCumulative            SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulative `json:"audio_producer_cumulative"`
	ScreenshareAudioConsumer           []interface{}                                                                                      `json:"screenshare_audio_consumer"`
	ScreenshareAudioConsumerCumulative interface{}                                                                                        `json:"screenshare_audio_consumer_cumulative"`
	ScreenshareAudioProducer           []interface{}                                                                                      `json:"screenshare_audio_producer"`
	ScreenshareAudioProducerCumulative interface{}                                                                                        `json:"screenshare_audio_producer_cumulative"`
	ScreenshareVideoConsumer           []interface{}                                                                                      `json:"screenshare_video_consumer"`
	ScreenshareVideoConsumerCumulative interface{}                                                                                        `json:"screenshare_video_consumer_cumulative"`
	ScreenshareVideoProducer           []interface{}                                                                                      `json:"screenshare_video_producer"`
	ScreenshareVideoProducerCumulative interface{}                                                                                        `json:"screenshare_video_producer_cumulative"`
	VideoConsumer                      []interface{}                                                                                      `json:"video_consumer"`
	VideoConsumerCumulative            interface{}                                                                                        `json:"video_consumer_cumulative"`
	VideoProducer                      []interface{}                                                                                      `json:"video_producer"`
	VideoProducerCumulative            interface{}                                                                                        `json:"video_producer_cumulative"`
	JSON                               sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityJSON                    `json:"-"`
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

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducer struct {
	BytesSent   int64                                                                                        `json:"bytes_sent"`
	Jitter      int64                                                                                        `json:"jitter"`
	Mid         string                                                                                       `json:"mid"`
	MosQuality  int64                                                                                        `json:"mos_quality"`
	PacketsLost int64                                                                                        `json:"packets_lost"`
	PacketsSent int64                                                                                        `json:"packets_sent"`
	ProducerID  string                                                                                       `json:"producer_id"`
	RTT         float64                                                                                      `json:"rtt"`
	Ssrc        int64                                                                                        `json:"ssrc"`
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

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulative struct {
	PacketLoss SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativePacketLoss `json:"packet_loss"`
	QualityMos SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeQualityMos `json:"quality_mos"`
	RTT        SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeRTT        `json:"rtt"`
	JSON       sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeJSON       `json:"-"`
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

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativePacketLoss struct {
	Number10OrGreaterEventFraction int64                                                                                                            `json:"10_or_greater_event_fraction"`
	Number25OrGreaterEventFraction int64                                                                                                            `json:"25_or_greater_event_fraction"`
	Number5OrGreaterEventFraction  int64                                                                                                            `json:"5_or_greater_event_fraction"`
	Number50OrGreaterEventFraction int64                                                                                                            `json:"50_or_greater_event_fraction"`
	Avg                            int64                                                                                                            `json:"avg"`
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

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerReportQualityAudioProducerCumulativeQualityMos struct {
	Avg  int64                                                                                                            `json:"avg"`
	P50  int64                                                                                                            `json:"p50"`
	P75  int64                                                                                                            `json:"p75"`
	P90  int64                                                                                                            `json:"p90"`
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

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStats struct {
	DeviceInfo                SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsDeviceInfo                `json:"device_info"`
	Events                    []SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEvent                   `json:"events"`
	IPInformation             SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformation             `json:"ip_information"`
	PrecallNetworkInformation SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsPrecallNetworkInformation `json:"precall_network_information"`
	JSON                      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsJSON                      `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsJSON contains
// the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStats]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsJSON struct {
	DeviceInfo                apijson.Field
	Events                    apijson.Field
	IPInformation             apijson.Field
	PrecallNetworkInformation apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsDeviceInfo struct {
	Browser        string                                                                            `json:"browser"`
	BrowserVersion string                                                                            `json:"browser_version"`
	CPUs           int64                                                                             `json:"cpus"`
	Engine         string                                                                            `json:"engine"`
	IsMobile       bool                                                                              `json:"is_mobile"`
	OS             string                                                                            `json:"os"`
	OSVersion      string                                                                            `json:"os_version"`
	SDKName        string                                                                            `json:"sdk_name"`
	SDKVersion     string                                                                            `json:"sdk_version"`
	UserAgent      string                                                                            `json:"user_agent"`
	WebglSupport   string                                                                            `json:"webgl_support"`
	JSON           sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsDeviceInfoJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsDeviceInfoJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsDeviceInfo]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsDeviceInfoJSON struct {
	Browser        apijson.Field
	BrowserVersion apijson.Field
	CPUs           apijson.Field
	Engine         apijson.Field
	IsMobile       apijson.Field
	OS             apijson.Field
	OSVersion      apijson.Field
	SDKName        apijson.Field
	SDKVersion     apijson.Field
	UserAgent      apijson.Field
	WebglSupport   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsDeviceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsDeviceInfoJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEvent struct {
	Metadata  SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadata `json:"metadata"`
	Timestamp string                                                                            `json:"timestamp"`
	Type      string                                                                            `json:"type"`
	JSON      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventJSON      `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEvent]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventJSON struct {
	Metadata    apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadata struct {
	ConnectionInfo SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfo `json:"connection_info"`
	JSON           sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataJSON           `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadata]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataJSON struct {
	ConnectionInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfo struct {
	BackendRTT           float64                                                                                                     `json:"backend_r_t_t"`
	Connectivity         SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoConnectivity `json:"connectivity"`
	EffectiveNetworkType string                                                                                                      `json:"effective_network_type"`
	FractionalLoss       int64                                                                                                       `json:"fractional_loss"`
	IPDetails            SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetails    `json:"ip_details"`
	Jitter               int64                                                                                                       `json:"jitter"`
	Location             SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocation     `json:"location"`
	RTT                  float64                                                                                                     `json:"r_t_t"`
	Throughput           int64                                                                                                       `json:"throughput"`
	TURNConnectivity     bool                                                                                                        `json:"turn_connectivity"`
	JSON                 sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoJSON         `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfo]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoJSON struct {
	BackendRTT           apijson.Field
	Connectivity         apijson.Field
	EffectiveNetworkType apijson.Field
	FractionalLoss       apijson.Field
	IPDetails            apijson.Field
	Jitter               apijson.Field
	Location             apijson.Field
	RTT                  apijson.Field
	Throughput           apijson.Field
	TURNConnectivity     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoConnectivity struct {
	Host      bool                                                                                                            `json:"host"`
	Reflexive bool                                                                                                            `json:"reflexive"`
	Relay     bool                                                                                                            `json:"relay"`
	JSON      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoConnectivityJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoConnectivityJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoConnectivity]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoConnectivityJSON struct {
	Host        apijson.Field
	Reflexive   apijson.Field
	Relay       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoConnectivity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoConnectivityJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetails struct {
	ASN      SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetailsASN  `json:"asn"`
	City     string                                                                                                       `json:"city"`
	Country  string                                                                                                       `json:"country"`
	IP       string                                                                                                       `json:"ip"`
	LOC      string                                                                                                       `json:"loc"`
	Postal   string                                                                                                       `json:"postal"`
	Region   string                                                                                                       `json:"region"`
	Timezone string                                                                                                       `json:"timezone"`
	JSON     sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetailsJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetailsJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetails]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetailsJSON struct {
	ASN         apijson.Field
	City        apijson.Field
	Country     apijson.Field
	IP          apijson.Field
	LOC         apijson.Field
	Postal      apijson.Field
	Region      apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetailsJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetailsASN struct {
	ASN  string                                                                                                          `json:"asn"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetailsASNJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetailsASNJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetailsASN]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetailsASNJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetailsASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoIPDetailsASNJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocation struct {
	Coords SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocationCoords `json:"coords"`
	JSON   sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocationJSON   `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocationJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocation]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocationJSON struct {
	Coords      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocationJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocationCoords struct {
	Latitude  float64                                                                                                           `json:"latitude"`
	Longitude float64                                                                                                           `json:"longitude"`
	JSON      sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocationCoordsJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocationCoordsJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocationCoords]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocationCoordsJSON struct {
	Latitude    apijson.Field
	Longitude   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocationCoords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsEventsMetadataConnectionInfoLocationCoordsJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformation struct {
	ASN        SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformationASN  `json:"asn"`
	City       string                                                                               `json:"city"`
	Country    string                                                                               `json:"country"`
	IPLocation string                                                                               `json:"ip_location"`
	IPV4       string                                                                               `json:"ipv4"`
	Org        string                                                                               `json:"org"`
	Region     string                                                                               `json:"region"`
	Timezone   string                                                                               `json:"timezone"`
	JSON       sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformationJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformationJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformation]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformationJSON struct {
	ASN         apijson.Field
	City        apijson.Field
	Country     apijson.Field
	IPLocation  apijson.Field
	IPV4        apijson.Field
	Org         apijson.Field
	Region      apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformationJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformationASN struct {
	ASN  string                                                                                  `json:"asn"`
	JSON sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformationASNJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformationASNJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformationASN]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformationASNJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformationASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsIPInformationASNJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsPrecallNetworkInformation struct {
	BackendRTT            float64                                                                                          `json:"backend_rtt"`
	EffectiveNetworktype  string                                                                                           `json:"effective_networktype"`
	FractionalLoss        int64                                                                                            `json:"fractional_loss"`
	Jitter                int64                                                                                            `json:"jitter"`
	ReflexiveConnectivity bool                                                                                             `json:"reflexive_connectivity"`
	RelayConnectivity     bool                                                                                             `json:"relay_connectivity"`
	RTT                   float64                                                                                          `json:"rtt"`
	Throughput            int64                                                                                            `json:"throughput"`
	TURNConnectivity      bool                                                                                             `json:"turn_connectivity"`
	JSON                  sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsPrecallNetworkInformationJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsPrecallNetworkInformationJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsPrecallNetworkInformation]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsPrecallNetworkInformationJSON struct {
	BackendRTT            apijson.Field
	EffectiveNetworktype  apijson.Field
	FractionalLoss        apijson.Field
	Jitter                apijson.Field
	ReflexiveConnectivity apijson.Field
	RelayConnectivity     apijson.Field
	RTT                   apijson.Field
	Throughput            apijson.Field
	TURNConnectivity      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsPrecallNetworkInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantPeerStatsPrecallNetworkInformationJSON) RawJSON() string {
	return r.raw
}

type SessionGetParticipantDataFromPeerIDResponseDataParticipantQualityStats struct {
	AudioBandwidth           int64                                                                      `json:"audio_bandwidth"`
	AudioStats               []interface{}                                                              `json:"audio_stats"`
	AverageQuality           int64                                                                      `json:"average_quality"`
	End                      string                                                                     `json:"end,nullable"`
	FirstAudioPacketReceived string                                                                     `json:"first_audio_packet_received"`
	FirstVideoPacketReceived string                                                                     `json:"first_video_packet_received"`
	LastAudioPacketReceived  string                                                                     `json:"last_audio_packet_received"`
	LastVideoPacketReceived  string                                                                     `json:"last_video_packet_received"`
	PeerIDs                  []string                                                                   `json:"peer_ids"`
	Start                    string                                                                     `json:"start,nullable"`
	TotalAudioPackets        int64                                                                      `json:"total_audio_packets"`
	TotalAudioPacketsLost    int64                                                                      `json:"total_audio_packets_lost"`
	TotalVideoPackets        int64                                                                      `json:"total_video_packets"`
	TotalVideoPacketsLost    int64                                                                      `json:"total_video_packets_lost"`
	VideoBandwidth           int64                                                                      `json:"video_bandwidth"`
	VideoStats               []interface{}                                                              `json:"video_stats"`
	JSON                     sessionGetParticipantDataFromPeerIDResponseDataParticipantQualityStatsJSON `json:"-"`
}

// sessionGetParticipantDataFromPeerIDResponseDataParticipantQualityStatsJSON
// contains the JSON metadata for the struct
// [SessionGetParticipantDataFromPeerIDResponseDataParticipantQualityStats]
type sessionGetParticipantDataFromPeerIDResponseDataParticipantQualityStatsJSON struct {
	AudioBandwidth           apijson.Field
	AudioStats               apijson.Field
	AverageQuality           apijson.Field
	End                      apijson.Field
	FirstAudioPacketReceived apijson.Field
	FirstVideoPacketReceived apijson.Field
	LastAudioPacketReceived  apijson.Field
	LastVideoPacketReceived  apijson.Field
	PeerIDs                  apijson.Field
	Start                    apijson.Field
	TotalAudioPackets        apijson.Field
	TotalAudioPacketsLost    apijson.Field
	TotalVideoPackets        apijson.Field
	TotalVideoPacketsLost    apijson.Field
	VideoBandwidth           apijson.Field
	VideoStats               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SessionGetParticipantDataFromPeerIDResponseDataParticipantQualityStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetParticipantDataFromPeerIDResponseDataParticipantQualityStatsJSON) RawJSON() string {
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
	ChatDownloadURL string `json:"chat_download_url,required"`
	// Time when the download URL will expire
	ChatDownloadURLExpiry string                                `json:"chat_download_url_expiry,required"`
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
	Session SessionGetSessionDetailsResponseDataSession `json:"session"`
	JSON    sessionGetSessionDetailsResponseDataJSON    `json:"-"`
}

// sessionGetSessionDetailsResponseDataJSON contains the JSON metadata for the
// struct [SessionGetSessionDetailsResponseData]
type sessionGetSessionDetailsResponseDataJSON struct {
	Session     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetSessionDetailsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionDetailsResponseDataJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionDetailsResponseDataSession struct {
	// ID of the session
	ID string `json:"id,required"`
	// ID of the meeting this session is associated with. In the case of V2 meetings,
	// it is always a UUID. In V1 meetings, it is a room name of the form
	// `abcdef-ghijkl`
	AssociatedID string `json:"associated_id,required"`
	// timestamp when session created
	CreatedAt string `json:"created_at,required"`
	// number of participants currently in the session
	LiveParticipants float64 `json:"live_participants,required"`
	// number of maximum participants that were in the session
	MaxConcurrentParticipants float64 `json:"max_concurrent_participants,required"`
	// Title of the meeting this session belongs to
	MeetingDisplayName string `json:"meeting_display_name,required"`
	// number of minutes consumed since the session started
	MinutesConsumed float64 `json:"minutes_consumed,required"`
	// App id that hosted this session
	OrganizationID string `json:"organization_id,required"`
	// timestamp when session started
	StartedAt string `json:"started_at,required"`
	// current status of session
	Status SessionGetSessionDetailsResponseDataSessionStatus `json:"status,required"`
	// type of session
	Type SessionGetSessionDetailsResponseDataSessionType `json:"type,required"`
	// timestamp when session was last updated
	UpdatedAt     string        `json:"updated_at,required"`
	BreakoutRooms []interface{} `json:"breakout_rooms"`
	// timestamp when session ended
	EndedAt string `json:"ended_at"`
	// Any meta data about session.
	Meta interface{}                                     `json:"meta"`
	JSON sessionGetSessionDetailsResponseDataSessionJSON `json:"-"`
}

// sessionGetSessionDetailsResponseDataSessionJSON contains the JSON metadata for
// the struct [SessionGetSessionDetailsResponseDataSession]
type sessionGetSessionDetailsResponseDataSessionJSON struct {
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

func (r *SessionGetSessionDetailsResponseDataSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionDetailsResponseDataSessionJSON) RawJSON() string {
	return r.raw
}

// current status of session
type SessionGetSessionDetailsResponseDataSessionStatus string

const (
	SessionGetSessionDetailsResponseDataSessionStatusLive  SessionGetSessionDetailsResponseDataSessionStatus = "LIVE"
	SessionGetSessionDetailsResponseDataSessionStatusEnded SessionGetSessionDetailsResponseDataSessionStatus = "ENDED"
)

func (r SessionGetSessionDetailsResponseDataSessionStatus) IsKnown() bool {
	switch r {
	case SessionGetSessionDetailsResponseDataSessionStatusLive, SessionGetSessionDetailsResponseDataSessionStatusEnded:
		return true
	}
	return false
}

// type of session
type SessionGetSessionDetailsResponseDataSessionType string

const (
	SessionGetSessionDetailsResponseDataSessionTypeMeeting     SessionGetSessionDetailsResponseDataSessionType = "meeting"
	SessionGetSessionDetailsResponseDataSessionTypeLivestream  SessionGetSessionDetailsResponseDataSessionType = "livestream"
	SessionGetSessionDetailsResponseDataSessionTypeParticipant SessionGetSessionDetailsResponseDataSessionType = "participant"
)

func (r SessionGetSessionDetailsResponseDataSessionType) IsKnown() bool {
	switch r {
	case SessionGetSessionDetailsResponseDataSessionTypeMeeting, SessionGetSessionDetailsResponseDataSessionTypeLivestream, SessionGetSessionDetailsResponseDataSessionTypeParticipant:
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
	LeftAt    string                                                              `json:"left_at"`
	PeerStats SessionGetSessionParticipantDetailsResponseDataParticipantPeerStats `json:"peer_stats"`
	// Name of the preset associated with the participant.
	PresetName   string                                                                  `json:"preset_name"`
	QualityStats []SessionGetSessionParticipantDetailsResponseDataParticipantQualityStat `json:"quality_stats"`
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
	PeerStats           apijson.Field
	PresetName          apijson.Field
	QualityStats        apijson.Field
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

type SessionGetSessionParticipantDetailsResponseDataParticipantPeerStats struct {
	Config                    string                                                                                       `json:"config"`
	DeviceInfo                SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsDeviceInfo                `json:"device_info"`
	Events                    []SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsEvent                   `json:"events"`
	IPInformation             SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsIPInformation             `json:"ip_information"`
	PrecallNetworkInformation SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsPrecallNetworkInformation `json:"precall_network_information"`
	Status                    string                                                                                       `json:"status"`
	JSON                      sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsJSON                      `json:"-"`
}

// sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsJSON contains
// the JSON metadata for the struct
// [SessionGetSessionParticipantDetailsResponseDataParticipantPeerStats]
type sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsJSON struct {
	Config                    apijson.Field
	DeviceInfo                apijson.Field
	Events                    apijson.Field
	IPInformation             apijson.Field
	PrecallNetworkInformation apijson.Field
	Status                    apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SessionGetSessionParticipantDetailsResponseDataParticipantPeerStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsDeviceInfo struct {
	Browser        string                                                                            `json:"browser"`
	BrowserVersion string                                                                            `json:"browser_version"`
	CPUs           float64                                                                           `json:"cpus"`
	Engine         string                                                                            `json:"engine"`
	IsMobile       bool                                                                              `json:"is_mobile"`
	Memory         float64                                                                           `json:"memory"`
	OS             string                                                                            `json:"os"`
	OSVersion      string                                                                            `json:"os_version"`
	SDKName        string                                                                            `json:"sdk_name"`
	SDKVersion     string                                                                            `json:"sdk_version"`
	UserAgent      string                                                                            `json:"user_agent"`
	WebglSupport   string                                                                            `json:"webgl_support"`
	JSON           sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsDeviceInfoJSON `json:"-"`
}

// sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsDeviceInfoJSON
// contains the JSON metadata for the struct
// [SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsDeviceInfo]
type sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsDeviceInfoJSON struct {
	Browser        apijson.Field
	BrowserVersion apijson.Field
	CPUs           apijson.Field
	Engine         apijson.Field
	IsMobile       apijson.Field
	Memory         apijson.Field
	OS             apijson.Field
	OSVersion      apijson.Field
	SDKName        apijson.Field
	SDKVersion     apijson.Field
	UserAgent      apijson.Field
	WebglSupport   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsDeviceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsDeviceInfoJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsEvent struct {
	Timestamp string                                                                       `json:"timestamp"`
	Type      string                                                                       `json:"type"`
	JSON      sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsEventJSON `json:"-"`
}

// sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsEventJSON
// contains the JSON metadata for the struct
// [SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsEvent]
type sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsEventJSON struct {
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsEventJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsIPInformation struct {
	City       string                                                                               `json:"city"`
	Country    string                                                                               `json:"country"`
	IPLocation string                                                                               `json:"ip_location"`
	IPV4       string                                                                               `json:"ipv4"`
	Org        string                                                                               `json:"org"`
	Portal     string                                                                               `json:"portal"`
	Region     string                                                                               `json:"region"`
	Timezone   string                                                                               `json:"timezone"`
	JSON       sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsIPInformationJSON `json:"-"`
}

// sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsIPInformationJSON
// contains the JSON metadata for the struct
// [SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsIPInformation]
type sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsIPInformationJSON struct {
	City        apijson.Field
	Country     apijson.Field
	IPLocation  apijson.Field
	IPV4        apijson.Field
	Org         apijson.Field
	Portal      apijson.Field
	Region      apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsIPInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsIPInformationJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsPrecallNetworkInformation struct {
	BackendRTT            float64                                                                                          `json:"backend_rtt"`
	EffectiveNetworktype  string                                                                                           `json:"effective_networktype"`
	FractionalLoss        float64                                                                                          `json:"fractional_loss"`
	Jitter                float64                                                                                          `json:"jitter"`
	ReflexiveConnectivity bool                                                                                             `json:"reflexive_connectivity"`
	RelayConnectivity     bool                                                                                             `json:"relay_connectivity"`
	RTT                   float64                                                                                          `json:"rtt"`
	Throughtput           float64                                                                                          `json:"throughtput"`
	TURNConnectivity      bool                                                                                             `json:"turn_connectivity"`
	JSON                  sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsPrecallNetworkInformationJSON `json:"-"`
}

// sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsPrecallNetworkInformationJSON
// contains the JSON metadata for the struct
// [SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsPrecallNetworkInformation]
type sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsPrecallNetworkInformationJSON struct {
	BackendRTT            apijson.Field
	EffectiveNetworktype  apijson.Field
	FractionalLoss        apijson.Field
	Jitter                apijson.Field
	ReflexiveConnectivity apijson.Field
	RelayConnectivity     apijson.Field
	RTT                   apijson.Field
	Throughtput           apijson.Field
	TURNConnectivity      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsPrecallNetworkInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantDetailsResponseDataParticipantPeerStatsPrecallNetworkInformationJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionParticipantDetailsResponseDataParticipantQualityStat struct {
	AudioBandwidth  float64                                                                           `json:"audio_bandwidth"`
	AudioPacketLoss float64                                                                           `json:"audio_packet_loss"`
	AudioStats      []SessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsAudioStat `json:"audio_stats"`
	AverageQuality  float64                                                                           `json:"average_quality"`
	End             string                                                                            `json:"end"`
	PeerID          string                                                                            `json:"peer_id"`
	Start           string                                                                            `json:"start"`
	VideoBandwidth  float64                                                                           `json:"video_bandwidth"`
	VideoPacketLoss float64                                                                           `json:"video_packet_loss"`
	VideoStats      []SessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsVideoStat `json:"video_stats"`
	JSON            sessionGetSessionParticipantDetailsResponseDataParticipantQualityStatJSON         `json:"-"`
}

// sessionGetSessionParticipantDetailsResponseDataParticipantQualityStatJSON
// contains the JSON metadata for the struct
// [SessionGetSessionParticipantDetailsResponseDataParticipantQualityStat]
type sessionGetSessionParticipantDetailsResponseDataParticipantQualityStatJSON struct {
	AudioBandwidth  apijson.Field
	AudioPacketLoss apijson.Field
	AudioStats      apijson.Field
	AverageQuality  apijson.Field
	End             apijson.Field
	PeerID          apijson.Field
	Start           apijson.Field
	VideoBandwidth  apijson.Field
	VideoPacketLoss apijson.Field
	VideoStats      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SessionGetSessionParticipantDetailsResponseDataParticipantQualityStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantDetailsResponseDataParticipantQualityStatJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsAudioStat struct {
	ConcealmentEvents float64                                                                             `json:"concealment_events"`
	Jitter            float64                                                                             `json:"jitter"`
	PacketsLost       float64                                                                             `json:"packets_lost"`
	Quality           float64                                                                             `json:"quality"`
	Timestamp         string                                                                              `json:"timestamp"`
	JSON              sessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsAudioStatJSON `json:"-"`
}

// sessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsAudioStatJSON
// contains the JSON metadata for the struct
// [SessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsAudioStat]
type sessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsAudioStatJSON struct {
	ConcealmentEvents apijson.Field
	Jitter            apijson.Field
	PacketsLost       apijson.Field
	Quality           apijson.Field
	Timestamp         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsAudioStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsAudioStatJSON) RawJSON() string {
	return r.raw
}

type SessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsVideoStat struct {
	FrameHeight     float64                                                                             `json:"frame_height"`
	FrameWidth      float64                                                                             `json:"frame_width"`
	FramesDropped   float64                                                                             `json:"frames_dropped"`
	FramesPerSecond float64                                                                             `json:"frames_per_second"`
	Jitter          float64                                                                             `json:"jitter"`
	PacketsLost     float64                                                                             `json:"packets_lost"`
	Quality         float64                                                                             `json:"quality"`
	Timestamp       string                                                                              `json:"timestamp"`
	JSON            sessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsVideoStatJSON `json:"-"`
}

// sessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsVideoStatJSON
// contains the JSON metadata for the struct
// [SessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsVideoStat]
type sessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsVideoStatJSON struct {
	FrameHeight     apijson.Field
	FrameWidth      apijson.Field
	FramesDropped   apijson.Field
	FramesPerSecond apijson.Field
	Jitter          apijson.Field
	PacketsLost     apijson.Field
	Quality         apijson.Field
	Timestamp       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsVideoStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetSessionParticipantDetailsResponseDataParticipantQualityStatsVideoStatJSON) RawJSON() string {
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
	SessionID string `json:"sessionId,required"`
	// URL where the summary of transcripts can be downloaded
	SummaryDownloadURL string `json:"summaryDownloadUrl,required"`
	// Time of Expiry before when you need to download the csv file.
	SummaryDownloadURLExpiry string                                   `json:"summaryDownloadUrlExpiry,required"`
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
	SessionID string `json:"sessionId,required"`
	// URL where the transcript can be downloaded
	TranscriptDownloadURL string `json:"transcript_download_url,required"`
	// Time when the download URL will expire
	TranscriptDownloadURLExpiry string                                       `json:"transcript_download_url_expiry,required"`
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
	Data    SessionGetSessionsResponseData `json:"data"`
	Success bool                           `json:"success"`
	JSON    sessionGetSessionsResponseJSON `json:"-"`
}

// sessionGetSessionsResponseJSON contains the JSON metadata for the struct
// [SessionGetSessionsResponse]
type sessionGetSessionsResponseJSON struct {
	Data        apijson.Field
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
	ID string `json:"id,required"`
	// ID of the meeting this session is associated with. In the case of V2 meetings,
	// it is always a UUID. In V1 meetings, it is a room name of the form
	// `abcdef-ghijkl`
	AssociatedID string `json:"associated_id,required"`
	// timestamp when session created
	CreatedAt string `json:"created_at,required"`
	// number of participants currently in the session
	LiveParticipants float64 `json:"live_participants,required"`
	// number of maximum participants that were in the session
	MaxConcurrentParticipants float64 `json:"max_concurrent_participants,required"`
	// Title of the meeting this session belongs to
	MeetingDisplayName string `json:"meeting_display_name,required"`
	// number of minutes consumed since the session started
	MinutesConsumed float64 `json:"minutes_consumed,required"`
	// App id that hosted this session
	OrganizationID string `json:"organization_id,required"`
	// timestamp when session started
	StartedAt string `json:"started_at,required"`
	// current status of session
	Status SessionGetSessionsResponseDataSessionsStatus `json:"status,required"`
	// type of session
	Type SessionGetSessionsResponseDataSessionsType `json:"type,required"`
	// timestamp when session was last updated
	UpdatedAt     string        `json:"updated_at,required"`
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

type SessionGenerateSummaryOfTranscriptsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SessionGetParticipantDataFromPeerIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Comma separated list of filters to apply. Note that there must be no spaces
	// between the filters.
	Filters param.Field[SessionGetParticipantDataFromPeerIDParamsFilters] `query:"filters"`
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type SessionGetSessionDetailsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
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
	AccountID param.Field[string] `path:"account_id,required"`
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
	AccountID param.Field[string] `path:"account_id,required"`
	// if true, response includes all the peer events of participants.
	IncludePeerEvents param.Field[bool] `query:"include_peer_events"`
	// The page number from which you want your page search results to be displayed.
	PageNo param.Field[float64] `query:"page_no"`
	// Number of results per page
	PerPage param.Field[float64] `query:"per_page"`
	// The search query string. You can search using the meeting ID or title.
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type SessionGetSessionTranscriptsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SessionGetSessionsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// ID of the meeting that sessions should be associated with
	AssociatedID param.Field[string] `query:"associated_id"`
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
