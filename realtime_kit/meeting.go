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

// MeetingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeetingService] method instead.
type MeetingService struct {
	Options []option.RequestOption
}

// NewMeetingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeetingService(opts ...option.RequestOption) (r *MeetingService) {
	r = &MeetingService{}
	r.Options = opts
	return
}

// Create a meeting for the given App ID.
func (r *MeetingService) New(ctx context.Context, appID string, params MeetingNewParams, opts ...option.RequestOption) (res *MeetingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Adds a participant to the given meeting ID.
func (r *MeetingService) AddParticipant(ctx context.Context, appID string, meetingID string, params MeetingAddParticipantParams, opts ...option.RequestOption) (res *MeetingAddParticipantResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/participants", params.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Deletes a participant for the given meeting and participant ID.
func (r *MeetingService) DeleteMeetingParticipant(ctx context.Context, appID string, meetingID string, participantID string, body MeetingDeleteMeetingParticipantParams, opts ...option.RequestOption) (res *MeetingDeleteMeetingParticipantResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	if participantID == "" {
		err = errors.New("missing required participant_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/participants/%s", body.AccountID, appID, meetingID, participantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Updates a participant's details for the given meeting and participant ID.
func (r *MeetingService) EditParticipant(ctx context.Context, appID string, meetingID string, participantID string, params MeetingEditParticipantParams, opts ...option.RequestOption) (res *MeetingEditParticipantResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	if participantID == "" {
		err = errors.New("missing required participant_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/participants/%s", params.AccountID, appID, meetingID, participantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns all meetings for the given App ID.
func (r *MeetingService) Get(ctx context.Context, appID string, params MeetingGetParams, opts ...option.RequestOption) (res *MeetingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Returns a meeting details in an App for the given meeting ID.
func (r *MeetingService) GetMeetingByID(ctx context.Context, appID string, meetingID string, params MeetingGetMeetingByIDParams, opts ...option.RequestOption) (res *MeetingGetMeetingByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s", params.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Returns a participant details for the given meeting and participant ID.
func (r *MeetingService) GetMeetingParticipant(ctx context.Context, appID string, meetingID string, participantID string, query MeetingGetMeetingParticipantParams, opts ...option.RequestOption) (res *MeetingGetMeetingParticipantResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	if participantID == "" {
		err = errors.New("missing required participant_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/participants/%s", query.AccountID, appID, meetingID, participantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all participants detail for the given meeting ID.
func (r *MeetingService) GetMeetingParticipants(ctx context.Context, appID string, meetingID string, params MeetingGetMeetingParticipantsParams, opts ...option.RequestOption) (res *MeetingGetMeetingParticipantsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/participants", params.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Regenerates participant's authentication token for the given meeting and
// participant ID.
func (r *MeetingService) RefreshParticipantToken(ctx context.Context, appID string, meetingID string, participantID string, body MeetingRefreshParticipantTokenParams, opts ...option.RequestOption) (res *MeetingRefreshParticipantTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	if participantID == "" {
		err = errors.New("missing required participant_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/participants/%s/token", body.AccountID, appID, meetingID, participantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Replaces all the details for the given meeting ID.
func (r *MeetingService) ReplaceMeetingByID(ctx context.Context, appID string, meetingID string, params MeetingReplaceMeetingByIDParams, opts ...option.RequestOption) (res *MeetingReplaceMeetingByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s", params.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Updates a meeting in an App for the given meeting ID.
func (r *MeetingService) UpdateMeetingByID(ctx context.Context, appID string, meetingID string, params MeetingUpdateMeetingByIDParams, opts ...option.RequestOption) (res *MeetingUpdateMeetingByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s", params.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type MeetingNewResponse struct {
	// Success status of the operation
	Success bool `json:"success,required"`
	// Data returned by the operation
	Data MeetingNewResponseData `json:"data"`
	JSON meetingNewResponseJSON `json:"-"`
}

// meetingNewResponseJSON contains the JSON metadata for the struct
// [MeetingNewResponse]
type meetingNewResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingNewResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type MeetingNewResponseData struct {
	// ID of the meeting.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp the object was created at. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Timestamp the object was updated at. The time is returned in ISO format.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The AI Config allows you to customize the behavior of meeting transcriptions and
	// summaries
	AIConfig MeetingNewResponseDataAIConfig `json:"ai_config"`
	// Specifies if the meeting should start getting livestreamed on start.
	LiveStreamOnStart bool `json:"live_stream_on_start"`
	// Specifies if Chat within a meeting should persist for a week.
	PersistChat bool `json:"persist_chat"`
	// Specifies if the meeting should start getting recorded as soon as someone joins
	// the meeting.
	RecordOnStart bool `json:"record_on_start"`
	// Recording Configurations to be used for this meeting. This level of configs
	// takes higher preference over App level configs on the RealtimeKit developer
	// portal.
	RecordingConfig MeetingNewResponseDataRecordingConfig `json:"recording_config"`
	// Time in seconds, for which a session remains active, after the last participant
	// has left the meeting.
	SessionKeepAliveTimeInSecs float64 `json:"session_keep_alive_time_in_secs"`
	// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
	// `INACTIVE` meeting.
	Status MeetingNewResponseDataStatus `json:"status"`
	// Automatically generate summary of meetings using transcripts. Requires
	// Transcriptions to be enabled, and can be retrieved via Webhooks or summary API.
	SummarizeOnEnd bool `json:"summarize_on_end"`
	// Title of the meeting.
	Title string                     `json:"title"`
	JSON  meetingNewResponseDataJSON `json:"-"`
}

// meetingNewResponseDataJSON contains the JSON metadata for the struct
// [MeetingNewResponseData]
type meetingNewResponseDataJSON struct {
	ID                         apijson.Field
	CreatedAt                  apijson.Field
	UpdatedAt                  apijson.Field
	AIConfig                   apijson.Field
	LiveStreamOnStart          apijson.Field
	PersistChat                apijson.Field
	RecordOnStart              apijson.Field
	RecordingConfig            apijson.Field
	SessionKeepAliveTimeInSecs apijson.Field
	Status                     apijson.Field
	SummarizeOnEnd             apijson.Field
	Title                      apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *MeetingNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingNewResponseDataJSON) RawJSON() string {
	return r.raw
}

// The AI Config allows you to customize the behavior of meeting transcriptions and
// summaries
type MeetingNewResponseDataAIConfig struct {
	// Summary Config
	Summarization MeetingNewResponseDataAIConfigSummarization `json:"summarization"`
	// Transcription Configurations
	Transcription MeetingNewResponseDataAIConfigTranscription `json:"transcription"`
	JSON          meetingNewResponseDataAIConfigJSON          `json:"-"`
}

// meetingNewResponseDataAIConfigJSON contains the JSON metadata for the struct
// [MeetingNewResponseDataAIConfig]
type meetingNewResponseDataAIConfigJSON struct {
	Summarization apijson.Field
	Transcription apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MeetingNewResponseDataAIConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingNewResponseDataAIConfigJSON) RawJSON() string {
	return r.raw
}

// Summary Config
type MeetingNewResponseDataAIConfigSummarization struct {
	// Defines the style of the summary, such as general, team meeting, or sales call.
	SummaryType MeetingNewResponseDataAIConfigSummarizationSummaryType `json:"summary_type"`
	// Determines the text format of the summary, such as plain text or markdown.
	TextFormat MeetingNewResponseDataAIConfigSummarizationTextFormat `json:"text_format"`
	// Sets the maximum number of words in the meeting summary.
	WordLimit int64                                           `json:"word_limit"`
	JSON      meetingNewResponseDataAIConfigSummarizationJSON `json:"-"`
}

// meetingNewResponseDataAIConfigSummarizationJSON contains the JSON metadata for
// the struct [MeetingNewResponseDataAIConfigSummarization]
type meetingNewResponseDataAIConfigSummarizationJSON struct {
	SummaryType apijson.Field
	TextFormat  apijson.Field
	WordLimit   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingNewResponseDataAIConfigSummarization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingNewResponseDataAIConfigSummarizationJSON) RawJSON() string {
	return r.raw
}

// Defines the style of the summary, such as general, team meeting, or sales call.
type MeetingNewResponseDataAIConfigSummarizationSummaryType string

const (
	MeetingNewResponseDataAIConfigSummarizationSummaryTypeGeneral         MeetingNewResponseDataAIConfigSummarizationSummaryType = "general"
	MeetingNewResponseDataAIConfigSummarizationSummaryTypeTeamMeeting     MeetingNewResponseDataAIConfigSummarizationSummaryType = "team_meeting"
	MeetingNewResponseDataAIConfigSummarizationSummaryTypeSalesCall       MeetingNewResponseDataAIConfigSummarizationSummaryType = "sales_call"
	MeetingNewResponseDataAIConfigSummarizationSummaryTypeClientCheckIn   MeetingNewResponseDataAIConfigSummarizationSummaryType = "client_check_in"
	MeetingNewResponseDataAIConfigSummarizationSummaryTypeInterview       MeetingNewResponseDataAIConfigSummarizationSummaryType = "interview"
	MeetingNewResponseDataAIConfigSummarizationSummaryTypeDailyStandup    MeetingNewResponseDataAIConfigSummarizationSummaryType = "daily_standup"
	MeetingNewResponseDataAIConfigSummarizationSummaryTypeOneOnOneMeeting MeetingNewResponseDataAIConfigSummarizationSummaryType = "one_on_one_meeting"
	MeetingNewResponseDataAIConfigSummarizationSummaryTypeLecture         MeetingNewResponseDataAIConfigSummarizationSummaryType = "lecture"
	MeetingNewResponseDataAIConfigSummarizationSummaryTypeCodeReview      MeetingNewResponseDataAIConfigSummarizationSummaryType = "code_review"
)

func (r MeetingNewResponseDataAIConfigSummarizationSummaryType) IsKnown() bool {
	switch r {
	case MeetingNewResponseDataAIConfigSummarizationSummaryTypeGeneral, MeetingNewResponseDataAIConfigSummarizationSummaryTypeTeamMeeting, MeetingNewResponseDataAIConfigSummarizationSummaryTypeSalesCall, MeetingNewResponseDataAIConfigSummarizationSummaryTypeClientCheckIn, MeetingNewResponseDataAIConfigSummarizationSummaryTypeInterview, MeetingNewResponseDataAIConfigSummarizationSummaryTypeDailyStandup, MeetingNewResponseDataAIConfigSummarizationSummaryTypeOneOnOneMeeting, MeetingNewResponseDataAIConfigSummarizationSummaryTypeLecture, MeetingNewResponseDataAIConfigSummarizationSummaryTypeCodeReview:
		return true
	}
	return false
}

// Determines the text format of the summary, such as plain text or markdown.
type MeetingNewResponseDataAIConfigSummarizationTextFormat string

const (
	MeetingNewResponseDataAIConfigSummarizationTextFormatPlainText MeetingNewResponseDataAIConfigSummarizationTextFormat = "plain_text"
	MeetingNewResponseDataAIConfigSummarizationTextFormatMarkdown  MeetingNewResponseDataAIConfigSummarizationTextFormat = "markdown"
)

func (r MeetingNewResponseDataAIConfigSummarizationTextFormat) IsKnown() bool {
	switch r {
	case MeetingNewResponseDataAIConfigSummarizationTextFormatPlainText, MeetingNewResponseDataAIConfigSummarizationTextFormatMarkdown:
		return true
	}
	return false
}

// Transcription Configurations
type MeetingNewResponseDataAIConfigTranscription struct {
	// Adds specific terms to improve accurate detection during transcription.
	Keywords []string `json:"keywords"`
	// Specifies the language code for transcription to ensure accurate results.
	Language MeetingNewResponseDataAIConfigTranscriptionLanguage `json:"language"`
	// Control the inclusion of offensive language in transcriptions.
	ProfanityFilter bool                                            `json:"profanity_filter"`
	JSON            meetingNewResponseDataAIConfigTranscriptionJSON `json:"-"`
}

// meetingNewResponseDataAIConfigTranscriptionJSON contains the JSON metadata for
// the struct [MeetingNewResponseDataAIConfigTranscription]
type meetingNewResponseDataAIConfigTranscriptionJSON struct {
	Keywords        apijson.Field
	Language        apijson.Field
	ProfanityFilter apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MeetingNewResponseDataAIConfigTranscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingNewResponseDataAIConfigTranscriptionJSON) RawJSON() string {
	return r.raw
}

// Specifies the language code for transcription to ensure accurate results.
type MeetingNewResponseDataAIConfigTranscriptionLanguage string

const (
	MeetingNewResponseDataAIConfigTranscriptionLanguageEnUs MeetingNewResponseDataAIConfigTranscriptionLanguage = "en-US"
	MeetingNewResponseDataAIConfigTranscriptionLanguageEnIn MeetingNewResponseDataAIConfigTranscriptionLanguage = "en-IN"
	MeetingNewResponseDataAIConfigTranscriptionLanguageDe   MeetingNewResponseDataAIConfigTranscriptionLanguage = "de"
	MeetingNewResponseDataAIConfigTranscriptionLanguageHi   MeetingNewResponseDataAIConfigTranscriptionLanguage = "hi"
	MeetingNewResponseDataAIConfigTranscriptionLanguageSv   MeetingNewResponseDataAIConfigTranscriptionLanguage = "sv"
	MeetingNewResponseDataAIConfigTranscriptionLanguageRu   MeetingNewResponseDataAIConfigTranscriptionLanguage = "ru"
	MeetingNewResponseDataAIConfigTranscriptionLanguagePl   MeetingNewResponseDataAIConfigTranscriptionLanguage = "pl"
	MeetingNewResponseDataAIConfigTranscriptionLanguageEl   MeetingNewResponseDataAIConfigTranscriptionLanguage = "el"
	MeetingNewResponseDataAIConfigTranscriptionLanguageFr   MeetingNewResponseDataAIConfigTranscriptionLanguage = "fr"
	MeetingNewResponseDataAIConfigTranscriptionLanguageNl   MeetingNewResponseDataAIConfigTranscriptionLanguage = "nl"
)

func (r MeetingNewResponseDataAIConfigTranscriptionLanguage) IsKnown() bool {
	switch r {
	case MeetingNewResponseDataAIConfigTranscriptionLanguageEnUs, MeetingNewResponseDataAIConfigTranscriptionLanguageEnIn, MeetingNewResponseDataAIConfigTranscriptionLanguageDe, MeetingNewResponseDataAIConfigTranscriptionLanguageHi, MeetingNewResponseDataAIConfigTranscriptionLanguageSv, MeetingNewResponseDataAIConfigTranscriptionLanguageRu, MeetingNewResponseDataAIConfigTranscriptionLanguagePl, MeetingNewResponseDataAIConfigTranscriptionLanguageEl, MeetingNewResponseDataAIConfigTranscriptionLanguageFr, MeetingNewResponseDataAIConfigTranscriptionLanguageNl:
		return true
	}
	return false
}

// Recording Configurations to be used for this meeting. This level of configs
// takes higher preference over App level configs on the RealtimeKit developer
// portal.
type MeetingNewResponseDataRecordingConfig struct {
	// Object containing configuration regarding the audio that is being recorded.
	AudioConfig MeetingNewResponseDataRecordingConfigAudioConfig `json:"audio_config"`
	// Adds a prefix to the beginning of the file name of the recording.
	FileNamePrefix      string                                                   `json:"file_name_prefix"`
	LiveStreamingConfig MeetingNewResponseDataRecordingConfigLiveStreamingConfig `json:"live_streaming_config"`
	// Specifies the maximum duration for recording in seconds, ranging from a minimum
	// of 60 seconds to a maximum of 24 hours.
	MaxSeconds              float64                                                      `json:"max_seconds"`
	RealtimekitBucketConfig MeetingNewResponseDataRecordingConfigRealtimekitBucketConfig `json:"realtimekit_bucket_config"`
	StorageConfig           MeetingNewResponseDataRecordingConfigStorageConfig           `json:"storage_config,nullable"`
	VideoConfig             MeetingNewResponseDataRecordingConfigVideoConfig             `json:"video_config"`
	JSON                    meetingNewResponseDataRecordingConfigJSON                    `json:"-"`
}

// meetingNewResponseDataRecordingConfigJSON contains the JSON metadata for the
// struct [MeetingNewResponseDataRecordingConfig]
type meetingNewResponseDataRecordingConfigJSON struct {
	AudioConfig             apijson.Field
	FileNamePrefix          apijson.Field
	LiveStreamingConfig     apijson.Field
	MaxSeconds              apijson.Field
	RealtimekitBucketConfig apijson.Field
	StorageConfig           apijson.Field
	VideoConfig             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *MeetingNewResponseDataRecordingConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingNewResponseDataRecordingConfigJSON) RawJSON() string {
	return r.raw
}

// Object containing configuration regarding the audio that is being recorded.
type MeetingNewResponseDataRecordingConfigAudioConfig struct {
	// Audio signal pathway within an audio file that carries a specific sound source.
	Channel MeetingNewResponseDataRecordingConfigAudioConfigChannel `json:"channel"`
	// Codec using which the recording will be encoded. If VP8/VP9 is selected for
	// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
	// audioConfig is automatically set to vorbis.
	Codec MeetingNewResponseDataRecordingConfigAudioConfigCodec `json:"codec"`
	// Controls whether to export audio file seperately
	ExportFile bool                                                 `json:"export_file"`
	JSON       meetingNewResponseDataRecordingConfigAudioConfigJSON `json:"-"`
}

// meetingNewResponseDataRecordingConfigAudioConfigJSON contains the JSON metadata
// for the struct [MeetingNewResponseDataRecordingConfigAudioConfig]
type meetingNewResponseDataRecordingConfigAudioConfigJSON struct {
	Channel     apijson.Field
	Codec       apijson.Field
	ExportFile  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingNewResponseDataRecordingConfigAudioConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingNewResponseDataRecordingConfigAudioConfigJSON) RawJSON() string {
	return r.raw
}

// Audio signal pathway within an audio file that carries a specific sound source.
type MeetingNewResponseDataRecordingConfigAudioConfigChannel string

const (
	MeetingNewResponseDataRecordingConfigAudioConfigChannelMono   MeetingNewResponseDataRecordingConfigAudioConfigChannel = "mono"
	MeetingNewResponseDataRecordingConfigAudioConfigChannelStereo MeetingNewResponseDataRecordingConfigAudioConfigChannel = "stereo"
)

func (r MeetingNewResponseDataRecordingConfigAudioConfigChannel) IsKnown() bool {
	switch r {
	case MeetingNewResponseDataRecordingConfigAudioConfigChannelMono, MeetingNewResponseDataRecordingConfigAudioConfigChannelStereo:
		return true
	}
	return false
}

// Codec using which the recording will be encoded. If VP8/VP9 is selected for
// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
// audioConfig is automatically set to vorbis.
type MeetingNewResponseDataRecordingConfigAudioConfigCodec string

const (
	MeetingNewResponseDataRecordingConfigAudioConfigCodecMP3 MeetingNewResponseDataRecordingConfigAudioConfigCodec = "MP3"
	MeetingNewResponseDataRecordingConfigAudioConfigCodecAac MeetingNewResponseDataRecordingConfigAudioConfigCodec = "AAC"
)

func (r MeetingNewResponseDataRecordingConfigAudioConfigCodec) IsKnown() bool {
	switch r {
	case MeetingNewResponseDataRecordingConfigAudioConfigCodecMP3, MeetingNewResponseDataRecordingConfigAudioConfigCodecAac:
		return true
	}
	return false
}

type MeetingNewResponseDataRecordingConfigLiveStreamingConfig struct {
	// RTMP URL to stream to
	RtmpURL string                                                       `json:"rtmp_url" format:"uri"`
	JSON    meetingNewResponseDataRecordingConfigLiveStreamingConfigJSON `json:"-"`
}

// meetingNewResponseDataRecordingConfigLiveStreamingConfigJSON contains the JSON
// metadata for the struct
// [MeetingNewResponseDataRecordingConfigLiveStreamingConfig]
type meetingNewResponseDataRecordingConfigLiveStreamingConfigJSON struct {
	RtmpURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingNewResponseDataRecordingConfigLiveStreamingConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingNewResponseDataRecordingConfigLiveStreamingConfigJSON) RawJSON() string {
	return r.raw
}

type MeetingNewResponseDataRecordingConfigRealtimekitBucketConfig struct {
	// Controls whether recordings are uploaded to RealtimeKit's bucket. If set to
	// false, `download_url`, `audio_download_url`, `download_url_expiry` won't be
	// generated for a recording.
	Enabled bool                                                             `json:"enabled,required"`
	JSON    meetingNewResponseDataRecordingConfigRealtimekitBucketConfigJSON `json:"-"`
}

// meetingNewResponseDataRecordingConfigRealtimekitBucketConfigJSON contains the
// JSON metadata for the struct
// [MeetingNewResponseDataRecordingConfigRealtimekitBucketConfig]
type meetingNewResponseDataRecordingConfigRealtimekitBucketConfigJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingNewResponseDataRecordingConfigRealtimekitBucketConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingNewResponseDataRecordingConfigRealtimekitBucketConfigJSON) RawJSON() string {
	return r.raw
}

type MeetingNewResponseDataRecordingConfigStorageConfig struct {
	// Type of storage media.
	Type MeetingNewResponseDataRecordingConfigStorageConfigType `json:"type,required"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod MeetingNewResponseDataRecordingConfigStorageConfigAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// SSH destination server password for SFTP type storage medium when auth_method is
	// "PASSWORD". If auth_method is "KEY", this specifies the password for the ssh
	// private key.
	Password string `json:"password"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// Private key used to login to destination SSH server for SFTP type storage
	// medium, when auth_method used is "KEY"
	PrivateKey string `json:"private_key"`
	// Region of the storage medium.
	Region string `json:"region"`
	// Secret key of the storage medium. Similar to `access_key`, it is only writeable
	// by clients, not readable.
	Secret string `json:"secret"`
	// SSH destination server username for SFTP type storage medium
	Username string                                                 `json:"username"`
	JSON     meetingNewResponseDataRecordingConfigStorageConfigJSON `json:"-"`
}

// meetingNewResponseDataRecordingConfigStorageConfigJSON contains the JSON
// metadata for the struct [MeetingNewResponseDataRecordingConfigStorageConfig]
type meetingNewResponseDataRecordingConfigStorageConfigJSON struct {
	Type        apijson.Field
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Password    apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	PrivateKey  apijson.Field
	Region      apijson.Field
	Secret      apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingNewResponseDataRecordingConfigStorageConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingNewResponseDataRecordingConfigStorageConfigJSON) RawJSON() string {
	return r.raw
}

// Type of storage media.
type MeetingNewResponseDataRecordingConfigStorageConfigType string

const (
	MeetingNewResponseDataRecordingConfigStorageConfigTypeAws          MeetingNewResponseDataRecordingConfigStorageConfigType = "aws"
	MeetingNewResponseDataRecordingConfigStorageConfigTypeAzure        MeetingNewResponseDataRecordingConfigStorageConfigType = "azure"
	MeetingNewResponseDataRecordingConfigStorageConfigTypeDigitalocean MeetingNewResponseDataRecordingConfigStorageConfigType = "digitalocean"
	MeetingNewResponseDataRecordingConfigStorageConfigTypeGcs          MeetingNewResponseDataRecordingConfigStorageConfigType = "gcs"
	MeetingNewResponseDataRecordingConfigStorageConfigTypeSftp         MeetingNewResponseDataRecordingConfigStorageConfigType = "sftp"
)

func (r MeetingNewResponseDataRecordingConfigStorageConfigType) IsKnown() bool {
	switch r {
	case MeetingNewResponseDataRecordingConfigStorageConfigTypeAws, MeetingNewResponseDataRecordingConfigStorageConfigTypeAzure, MeetingNewResponseDataRecordingConfigStorageConfigTypeDigitalocean, MeetingNewResponseDataRecordingConfigStorageConfigTypeGcs, MeetingNewResponseDataRecordingConfigStorageConfigTypeSftp:
		return true
	}
	return false
}

// Authentication method used for "sftp" type storage medium
type MeetingNewResponseDataRecordingConfigStorageConfigAuthMethod string

const (
	MeetingNewResponseDataRecordingConfigStorageConfigAuthMethodKey      MeetingNewResponseDataRecordingConfigStorageConfigAuthMethod = "KEY"
	MeetingNewResponseDataRecordingConfigStorageConfigAuthMethodPassword MeetingNewResponseDataRecordingConfigStorageConfigAuthMethod = "PASSWORD"
)

func (r MeetingNewResponseDataRecordingConfigStorageConfigAuthMethod) IsKnown() bool {
	switch r {
	case MeetingNewResponseDataRecordingConfigStorageConfigAuthMethodKey, MeetingNewResponseDataRecordingConfigStorageConfigAuthMethodPassword:
		return true
	}
	return false
}

type MeetingNewResponseDataRecordingConfigVideoConfig struct {
	// Codec using which the recording will be encoded.
	Codec MeetingNewResponseDataRecordingConfigVideoConfigCodec `json:"codec"`
	// Controls whether to export video file seperately
	ExportFile bool `json:"export_file"`
	// Height of the recording video in pixels
	Height int64 `json:"height"`
	// Watermark to be added to the recording
	Watermark MeetingNewResponseDataRecordingConfigVideoConfigWatermark `json:"watermark"`
	// Width of the recording video in pixels
	Width int64                                                `json:"width"`
	JSON  meetingNewResponseDataRecordingConfigVideoConfigJSON `json:"-"`
}

// meetingNewResponseDataRecordingConfigVideoConfigJSON contains the JSON metadata
// for the struct [MeetingNewResponseDataRecordingConfigVideoConfig]
type meetingNewResponseDataRecordingConfigVideoConfigJSON struct {
	Codec       apijson.Field
	ExportFile  apijson.Field
	Height      apijson.Field
	Watermark   apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingNewResponseDataRecordingConfigVideoConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingNewResponseDataRecordingConfigVideoConfigJSON) RawJSON() string {
	return r.raw
}

// Codec using which the recording will be encoded.
type MeetingNewResponseDataRecordingConfigVideoConfigCodec string

const (
	MeetingNewResponseDataRecordingConfigVideoConfigCodecH264 MeetingNewResponseDataRecordingConfigVideoConfigCodec = "H264"
	MeetingNewResponseDataRecordingConfigVideoConfigCodecVp8  MeetingNewResponseDataRecordingConfigVideoConfigCodec = "VP8"
)

func (r MeetingNewResponseDataRecordingConfigVideoConfigCodec) IsKnown() bool {
	switch r {
	case MeetingNewResponseDataRecordingConfigVideoConfigCodecH264, MeetingNewResponseDataRecordingConfigVideoConfigCodecVp8:
		return true
	}
	return false
}

// Watermark to be added to the recording
type MeetingNewResponseDataRecordingConfigVideoConfigWatermark struct {
	// Position of the watermark
	Position MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPosition `json:"position"`
	// Size of the watermark
	Size MeetingNewResponseDataRecordingConfigVideoConfigWatermarkSize `json:"size"`
	// URL of the watermark image
	URL  string                                                        `json:"url" format:"uri"`
	JSON meetingNewResponseDataRecordingConfigVideoConfigWatermarkJSON `json:"-"`
}

// meetingNewResponseDataRecordingConfigVideoConfigWatermarkJSON contains the JSON
// metadata for the struct
// [MeetingNewResponseDataRecordingConfigVideoConfigWatermark]
type meetingNewResponseDataRecordingConfigVideoConfigWatermarkJSON struct {
	Position    apijson.Field
	Size        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingNewResponseDataRecordingConfigVideoConfigWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingNewResponseDataRecordingConfigVideoConfigWatermarkJSON) RawJSON() string {
	return r.raw
}

// Position of the watermark
type MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPosition string

const (
	MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPositionLeftTop     MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPosition = "left top"
	MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPositionRightTop    MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPosition = "right top"
	MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPositionLeftBottom  MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPosition = "left bottom"
	MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPositionRightBottom MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPosition = "right bottom"
)

func (r MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPosition) IsKnown() bool {
	switch r {
	case MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPositionLeftTop, MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPositionRightTop, MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPositionLeftBottom, MeetingNewResponseDataRecordingConfigVideoConfigWatermarkPositionRightBottom:
		return true
	}
	return false
}

// Size of the watermark
type MeetingNewResponseDataRecordingConfigVideoConfigWatermarkSize struct {
	// Height of the watermark in px
	Height int64 `json:"height"`
	// Width of the watermark in px
	Width int64                                                             `json:"width"`
	JSON  meetingNewResponseDataRecordingConfigVideoConfigWatermarkSizeJSON `json:"-"`
}

// meetingNewResponseDataRecordingConfigVideoConfigWatermarkSizeJSON contains the
// JSON metadata for the struct
// [MeetingNewResponseDataRecordingConfigVideoConfigWatermarkSize]
type meetingNewResponseDataRecordingConfigVideoConfigWatermarkSizeJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingNewResponseDataRecordingConfigVideoConfigWatermarkSize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingNewResponseDataRecordingConfigVideoConfigWatermarkSizeJSON) RawJSON() string {
	return r.raw
}

// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
// `INACTIVE` meeting.
type MeetingNewResponseDataStatus string

const (
	MeetingNewResponseDataStatusActive   MeetingNewResponseDataStatus = "ACTIVE"
	MeetingNewResponseDataStatusInactive MeetingNewResponseDataStatus = "INACTIVE"
)

func (r MeetingNewResponseDataStatus) IsKnown() bool {
	switch r {
	case MeetingNewResponseDataStatusActive, MeetingNewResponseDataStatusInactive:
		return true
	}
	return false
}

type MeetingAddParticipantResponse struct {
	// Success status of the operation
	Success bool `json:"success,required"`
	// Represents a participant.
	Data MeetingAddParticipantResponseData `json:"data"`
	JSON meetingAddParticipantResponseJSON `json:"-"`
}

// meetingAddParticipantResponseJSON contains the JSON metadata for the struct
// [MeetingAddParticipantResponse]
type meetingAddParticipantResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingAddParticipantResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingAddParticipantResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a participant.
type MeetingAddParticipantResponseData struct {
	// ID of the participant.
	ID string `json:"id,required" format:"uuid"`
	// The participant's auth token that can be used for joining a meeting from the
	// client side.
	Token string `json:"token,required"`
	// When this object was created. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A unique participant ID generated by the client.
	CustomParticipantID string `json:"custom_participant_id,required"`
	// Preset applied to the participant.
	PresetName string `json:"preset_name,required"`
	// When this object was updated. The time is returned in ISO format.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Name of the participant.
	Name string `json:"name,nullable"`
	// URL to a picture of the participant.
	Picture string                                `json:"picture,nullable" format:"uri"`
	JSON    meetingAddParticipantResponseDataJSON `json:"-"`
}

// meetingAddParticipantResponseDataJSON contains the JSON metadata for the struct
// [MeetingAddParticipantResponseData]
type meetingAddParticipantResponseDataJSON struct {
	ID                  apijson.Field
	Token               apijson.Field
	CreatedAt           apijson.Field
	CustomParticipantID apijson.Field
	PresetName          apijson.Field
	UpdatedAt           apijson.Field
	Name                apijson.Field
	Picture             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MeetingAddParticipantResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingAddParticipantResponseDataJSON) RawJSON() string {
	return r.raw
}

type MeetingDeleteMeetingParticipantResponse struct {
	// Success status of the operation
	Success bool `json:"success,required"`
	// Data returned by the operation
	Data MeetingDeleteMeetingParticipantResponseData `json:"data"`
	JSON meetingDeleteMeetingParticipantResponseJSON `json:"-"`
}

// meetingDeleteMeetingParticipantResponseJSON contains the JSON metadata for the
// struct [MeetingDeleteMeetingParticipantResponse]
type meetingDeleteMeetingParticipantResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingDeleteMeetingParticipantResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingDeleteMeetingParticipantResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type MeetingDeleteMeetingParticipantResponseData struct {
	// Timestamp this object was created at. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A unique participant ID generated by the client.
	CustomParticipantID string `json:"custom_participant_id,required"`
	// ID of the preset applied to this participant.
	PresetID string `json:"preset_id,required" format:"uuid"`
	// Timestamp this object was updated at. The time is returned in ISO format.
	UpdatedAt time.Time                                       `json:"updated_at,required" format:"date-time"`
	JSON      meetingDeleteMeetingParticipantResponseDataJSON `json:"-"`
}

// meetingDeleteMeetingParticipantResponseDataJSON contains the JSON metadata for
// the struct [MeetingDeleteMeetingParticipantResponseData]
type meetingDeleteMeetingParticipantResponseDataJSON struct {
	CreatedAt           apijson.Field
	CustomParticipantID apijson.Field
	PresetID            apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MeetingDeleteMeetingParticipantResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingDeleteMeetingParticipantResponseDataJSON) RawJSON() string {
	return r.raw
}

type MeetingEditParticipantResponse struct {
	// Success status of the operation
	Success bool `json:"success,required"`
	// Represents a participant.
	Data MeetingEditParticipantResponseData `json:"data"`
	JSON meetingEditParticipantResponseJSON `json:"-"`
}

// meetingEditParticipantResponseJSON contains the JSON metadata for the struct
// [MeetingEditParticipantResponse]
type meetingEditParticipantResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingEditParticipantResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingEditParticipantResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a participant.
type MeetingEditParticipantResponseData struct {
	// ID of the participant.
	ID string `json:"id,required" format:"uuid"`
	// The participant's auth token that can be used for joining a meeting from the
	// client side.
	Token string `json:"token,required"`
	// When this object was created. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A unique participant ID generated by the client.
	CustomParticipantID string `json:"custom_participant_id,required"`
	// Preset applied to the participant.
	PresetName string `json:"preset_name,required"`
	// When this object was updated. The time is returned in ISO format.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Name of the participant.
	Name string `json:"name,nullable"`
	// URL to a picture of the participant.
	Picture string                                 `json:"picture,nullable" format:"uri"`
	JSON    meetingEditParticipantResponseDataJSON `json:"-"`
}

// meetingEditParticipantResponseDataJSON contains the JSON metadata for the struct
// [MeetingEditParticipantResponseData]
type meetingEditParticipantResponseDataJSON struct {
	ID                  apijson.Field
	Token               apijson.Field
	CreatedAt           apijson.Field
	CustomParticipantID apijson.Field
	PresetName          apijson.Field
	UpdatedAt           apijson.Field
	Name                apijson.Field
	Picture             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MeetingEditParticipantResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingEditParticipantResponseDataJSON) RawJSON() string {
	return r.raw
}

type MeetingGetResponse struct {
	Data    []MeetingGetResponseData `json:"data,required"`
	Paging  MeetingGetResponsePaging `json:"paging,required"`
	Success bool                     `json:"success,required"`
	JSON    meetingGetResponseJSON   `json:"-"`
}

// meetingGetResponseJSON contains the JSON metadata for the struct
// [MeetingGetResponse]
type meetingGetResponseJSON struct {
	Data        apijson.Field
	Paging      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetResponseJSON) RawJSON() string {
	return r.raw
}

type MeetingGetResponseData struct {
	// ID of the meeting.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp the object was created at. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Timestamp the object was updated at. The time is returned in ISO format.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Specifies if the meeting should start getting livestreamed on start.
	LiveStreamOnStart bool `json:"live_stream_on_start"`
	// Specifies if Chat within a meeting should persist for a week.
	PersistChat bool `json:"persist_chat"`
	// Specifies if the meeting should start getting recorded as soon as someone joins
	// the meeting.
	RecordOnStart bool `json:"record_on_start"`
	// Time in seconds, for which a session remains active, after the last participant
	// has left the meeting.
	SessionKeepAliveTimeInSecs float64 `json:"session_keep_alive_time_in_secs"`
	// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
	// `INACTIVE` meeting.
	Status MeetingGetResponseDataStatus `json:"status"`
	// Automatically generate summary of meetings using transcripts. Requires
	// Transcriptions to be enabled, and can be retrieved via Webhooks or summary API.
	SummarizeOnEnd bool `json:"summarize_on_end"`
	// Title of the meeting.
	Title string                     `json:"title"`
	JSON  meetingGetResponseDataJSON `json:"-"`
}

// meetingGetResponseDataJSON contains the JSON metadata for the struct
// [MeetingGetResponseData]
type meetingGetResponseDataJSON struct {
	ID                         apijson.Field
	CreatedAt                  apijson.Field
	UpdatedAt                  apijson.Field
	LiveStreamOnStart          apijson.Field
	PersistChat                apijson.Field
	RecordOnStart              apijson.Field
	SessionKeepAliveTimeInSecs apijson.Field
	Status                     apijson.Field
	SummarizeOnEnd             apijson.Field
	Title                      apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *MeetingGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
// `INACTIVE` meeting.
type MeetingGetResponseDataStatus string

const (
	MeetingGetResponseDataStatusActive   MeetingGetResponseDataStatus = "ACTIVE"
	MeetingGetResponseDataStatusInactive MeetingGetResponseDataStatus = "INACTIVE"
)

func (r MeetingGetResponseDataStatus) IsKnown() bool {
	switch r {
	case MeetingGetResponseDataStatusActive, MeetingGetResponseDataStatusInactive:
		return true
	}
	return false
}

type MeetingGetResponsePaging struct {
	EndOffset   float64                      `json:"end_offset,required"`
	StartOffset float64                      `json:"start_offset,required"`
	TotalCount  float64                      `json:"total_count,required"`
	JSON        meetingGetResponsePagingJSON `json:"-"`
}

// meetingGetResponsePagingJSON contains the JSON metadata for the struct
// [MeetingGetResponsePaging]
type meetingGetResponsePagingJSON struct {
	EndOffset   apijson.Field
	StartOffset apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetResponsePaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetResponsePagingJSON) RawJSON() string {
	return r.raw
}

type MeetingGetMeetingByIDResponse struct {
	// Success status of the operation
	Success bool `json:"success,required"`
	// Data returned by the operation
	Data MeetingGetMeetingByIDResponseData `json:"data"`
	JSON meetingGetMeetingByIDResponseJSON `json:"-"`
}

// meetingGetMeetingByIDResponseJSON contains the JSON metadata for the struct
// [MeetingGetMeetingByIDResponse]
type meetingGetMeetingByIDResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetMeetingByIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingByIDResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type MeetingGetMeetingByIDResponseData struct {
	// ID of the meeting.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp the object was created at. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Timestamp the object was updated at. The time is returned in ISO format.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The AI Config allows you to customize the behavior of meeting transcriptions and
	// summaries
	AIConfig MeetingGetMeetingByIDResponseDataAIConfig `json:"ai_config"`
	// Specifies if the meeting should start getting livestreamed on start.
	LiveStreamOnStart bool `json:"live_stream_on_start"`
	// Specifies if Chat within a meeting should persist for a week.
	PersistChat bool `json:"persist_chat"`
	// Specifies if the meeting should start getting recorded as soon as someone joins
	// the meeting.
	RecordOnStart bool `json:"record_on_start"`
	// Recording Configurations to be used for this meeting. This level of configs
	// takes higher preference over App level configs on the RealtimeKit developer
	// portal.
	RecordingConfig MeetingGetMeetingByIDResponseDataRecordingConfig `json:"recording_config"`
	// Time in seconds, for which a session remains active, after the last participant
	// has left the meeting.
	SessionKeepAliveTimeInSecs float64 `json:"session_keep_alive_time_in_secs"`
	// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
	// `INACTIVE` meeting.
	Status MeetingGetMeetingByIDResponseDataStatus `json:"status"`
	// Automatically generate summary of meetings using transcripts. Requires
	// Transcriptions to be enabled, and can be retrieved via Webhooks or summary API.
	SummarizeOnEnd bool `json:"summarize_on_end"`
	// Title of the meeting.
	Title string                                `json:"title"`
	JSON  meetingGetMeetingByIDResponseDataJSON `json:"-"`
}

// meetingGetMeetingByIDResponseDataJSON contains the JSON metadata for the struct
// [MeetingGetMeetingByIDResponseData]
type meetingGetMeetingByIDResponseDataJSON struct {
	ID                         apijson.Field
	CreatedAt                  apijson.Field
	UpdatedAt                  apijson.Field
	AIConfig                   apijson.Field
	LiveStreamOnStart          apijson.Field
	PersistChat                apijson.Field
	RecordOnStart              apijson.Field
	RecordingConfig            apijson.Field
	SessionKeepAliveTimeInSecs apijson.Field
	Status                     apijson.Field
	SummarizeOnEnd             apijson.Field
	Title                      apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *MeetingGetMeetingByIDResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingByIDResponseDataJSON) RawJSON() string {
	return r.raw
}

// The AI Config allows you to customize the behavior of meeting transcriptions and
// summaries
type MeetingGetMeetingByIDResponseDataAIConfig struct {
	// Summary Config
	Summarization MeetingGetMeetingByIDResponseDataAIConfigSummarization `json:"summarization"`
	// Transcription Configurations
	Transcription MeetingGetMeetingByIDResponseDataAIConfigTranscription `json:"transcription"`
	JSON          meetingGetMeetingByIDResponseDataAIConfigJSON          `json:"-"`
}

// meetingGetMeetingByIDResponseDataAIConfigJSON contains the JSON metadata for the
// struct [MeetingGetMeetingByIDResponseDataAIConfig]
type meetingGetMeetingByIDResponseDataAIConfigJSON struct {
	Summarization apijson.Field
	Transcription apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MeetingGetMeetingByIDResponseDataAIConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingByIDResponseDataAIConfigJSON) RawJSON() string {
	return r.raw
}

// Summary Config
type MeetingGetMeetingByIDResponseDataAIConfigSummarization struct {
	// Defines the style of the summary, such as general, team meeting, or sales call.
	SummaryType MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryType `json:"summary_type"`
	// Determines the text format of the summary, such as plain text or markdown.
	TextFormat MeetingGetMeetingByIDResponseDataAIConfigSummarizationTextFormat `json:"text_format"`
	// Sets the maximum number of words in the meeting summary.
	WordLimit int64                                                      `json:"word_limit"`
	JSON      meetingGetMeetingByIDResponseDataAIConfigSummarizationJSON `json:"-"`
}

// meetingGetMeetingByIDResponseDataAIConfigSummarizationJSON contains the JSON
// metadata for the struct [MeetingGetMeetingByIDResponseDataAIConfigSummarization]
type meetingGetMeetingByIDResponseDataAIConfigSummarizationJSON struct {
	SummaryType apijson.Field
	TextFormat  apijson.Field
	WordLimit   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetMeetingByIDResponseDataAIConfigSummarization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingByIDResponseDataAIConfigSummarizationJSON) RawJSON() string {
	return r.raw
}

// Defines the style of the summary, such as general, team meeting, or sales call.
type MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryType string

const (
	MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeGeneral         MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryType = "general"
	MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeTeamMeeting     MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryType = "team_meeting"
	MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeSalesCall       MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryType = "sales_call"
	MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeClientCheckIn   MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryType = "client_check_in"
	MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeInterview       MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryType = "interview"
	MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeDailyStandup    MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryType = "daily_standup"
	MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeOneOnOneMeeting MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryType = "one_on_one_meeting"
	MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeLecture         MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryType = "lecture"
	MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeCodeReview      MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryType = "code_review"
)

func (r MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryType) IsKnown() bool {
	switch r {
	case MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeGeneral, MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeTeamMeeting, MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeSalesCall, MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeClientCheckIn, MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeInterview, MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeDailyStandup, MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeOneOnOneMeeting, MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeLecture, MeetingGetMeetingByIDResponseDataAIConfigSummarizationSummaryTypeCodeReview:
		return true
	}
	return false
}

// Determines the text format of the summary, such as plain text or markdown.
type MeetingGetMeetingByIDResponseDataAIConfigSummarizationTextFormat string

const (
	MeetingGetMeetingByIDResponseDataAIConfigSummarizationTextFormatPlainText MeetingGetMeetingByIDResponseDataAIConfigSummarizationTextFormat = "plain_text"
	MeetingGetMeetingByIDResponseDataAIConfigSummarizationTextFormatMarkdown  MeetingGetMeetingByIDResponseDataAIConfigSummarizationTextFormat = "markdown"
)

func (r MeetingGetMeetingByIDResponseDataAIConfigSummarizationTextFormat) IsKnown() bool {
	switch r {
	case MeetingGetMeetingByIDResponseDataAIConfigSummarizationTextFormatPlainText, MeetingGetMeetingByIDResponseDataAIConfigSummarizationTextFormatMarkdown:
		return true
	}
	return false
}

// Transcription Configurations
type MeetingGetMeetingByIDResponseDataAIConfigTranscription struct {
	// Adds specific terms to improve accurate detection during transcription.
	Keywords []string `json:"keywords"`
	// Specifies the language code for transcription to ensure accurate results.
	Language MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguage `json:"language"`
	// Control the inclusion of offensive language in transcriptions.
	ProfanityFilter bool                                                       `json:"profanity_filter"`
	JSON            meetingGetMeetingByIDResponseDataAIConfigTranscriptionJSON `json:"-"`
}

// meetingGetMeetingByIDResponseDataAIConfigTranscriptionJSON contains the JSON
// metadata for the struct [MeetingGetMeetingByIDResponseDataAIConfigTranscription]
type meetingGetMeetingByIDResponseDataAIConfigTranscriptionJSON struct {
	Keywords        apijson.Field
	Language        apijson.Field
	ProfanityFilter apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MeetingGetMeetingByIDResponseDataAIConfigTranscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingByIDResponseDataAIConfigTranscriptionJSON) RawJSON() string {
	return r.raw
}

// Specifies the language code for transcription to ensure accurate results.
type MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguage string

const (
	MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageEnUs MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguage = "en-US"
	MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageEnIn MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguage = "en-IN"
	MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageDe   MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguage = "de"
	MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageHi   MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguage = "hi"
	MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageSv   MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguage = "sv"
	MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageRu   MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguage = "ru"
	MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguagePl   MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguage = "pl"
	MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageEl   MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguage = "el"
	MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageFr   MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguage = "fr"
	MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageNl   MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguage = "nl"
)

func (r MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguage) IsKnown() bool {
	switch r {
	case MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageEnUs, MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageEnIn, MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageDe, MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageHi, MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageSv, MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageRu, MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguagePl, MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageEl, MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageFr, MeetingGetMeetingByIDResponseDataAIConfigTranscriptionLanguageNl:
		return true
	}
	return false
}

// Recording Configurations to be used for this meeting. This level of configs
// takes higher preference over App level configs on the RealtimeKit developer
// portal.
type MeetingGetMeetingByIDResponseDataRecordingConfig struct {
	// Object containing configuration regarding the audio that is being recorded.
	AudioConfig MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfig `json:"audio_config"`
	// Adds a prefix to the beginning of the file name of the recording.
	FileNamePrefix      string                                                              `json:"file_name_prefix"`
	LiveStreamingConfig MeetingGetMeetingByIDResponseDataRecordingConfigLiveStreamingConfig `json:"live_streaming_config"`
	// Specifies the maximum duration for recording in seconds, ranging from a minimum
	// of 60 seconds to a maximum of 24 hours.
	MaxSeconds              float64                                                                 `json:"max_seconds"`
	RealtimekitBucketConfig MeetingGetMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfig `json:"realtimekit_bucket_config"`
	StorageConfig           MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfig           `json:"storage_config,nullable"`
	VideoConfig             MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfig             `json:"video_config"`
	JSON                    meetingGetMeetingByIDResponseDataRecordingConfigJSON                    `json:"-"`
}

// meetingGetMeetingByIDResponseDataRecordingConfigJSON contains the JSON metadata
// for the struct [MeetingGetMeetingByIDResponseDataRecordingConfig]
type meetingGetMeetingByIDResponseDataRecordingConfigJSON struct {
	AudioConfig             apijson.Field
	FileNamePrefix          apijson.Field
	LiveStreamingConfig     apijson.Field
	MaxSeconds              apijson.Field
	RealtimekitBucketConfig apijson.Field
	StorageConfig           apijson.Field
	VideoConfig             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *MeetingGetMeetingByIDResponseDataRecordingConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingByIDResponseDataRecordingConfigJSON) RawJSON() string {
	return r.raw
}

// Object containing configuration regarding the audio that is being recorded.
type MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfig struct {
	// Audio signal pathway within an audio file that carries a specific sound source.
	Channel MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigChannel `json:"channel"`
	// Codec using which the recording will be encoded. If VP8/VP9 is selected for
	// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
	// audioConfig is automatically set to vorbis.
	Codec MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigCodec `json:"codec"`
	// Controls whether to export audio file seperately
	ExportFile bool                                                            `json:"export_file"`
	JSON       meetingGetMeetingByIDResponseDataRecordingConfigAudioConfigJSON `json:"-"`
}

// meetingGetMeetingByIDResponseDataRecordingConfigAudioConfigJSON contains the
// JSON metadata for the struct
// [MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfig]
type meetingGetMeetingByIDResponseDataRecordingConfigAudioConfigJSON struct {
	Channel     apijson.Field
	Codec       apijson.Field
	ExportFile  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingByIDResponseDataRecordingConfigAudioConfigJSON) RawJSON() string {
	return r.raw
}

// Audio signal pathway within an audio file that carries a specific sound source.
type MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigChannel string

const (
	MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigChannelMono   MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigChannel = "mono"
	MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigChannelStereo MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigChannel = "stereo"
)

func (r MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigChannel) IsKnown() bool {
	switch r {
	case MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigChannelMono, MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigChannelStereo:
		return true
	}
	return false
}

// Codec using which the recording will be encoded. If VP8/VP9 is selected for
// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
// audioConfig is automatically set to vorbis.
type MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigCodec string

const (
	MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigCodecMP3 MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigCodec = "MP3"
	MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigCodecAac MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigCodec = "AAC"
)

func (r MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigCodec) IsKnown() bool {
	switch r {
	case MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigCodecMP3, MeetingGetMeetingByIDResponseDataRecordingConfigAudioConfigCodecAac:
		return true
	}
	return false
}

type MeetingGetMeetingByIDResponseDataRecordingConfigLiveStreamingConfig struct {
	// RTMP URL to stream to
	RtmpURL string                                                                  `json:"rtmp_url" format:"uri"`
	JSON    meetingGetMeetingByIDResponseDataRecordingConfigLiveStreamingConfigJSON `json:"-"`
}

// meetingGetMeetingByIDResponseDataRecordingConfigLiveStreamingConfigJSON contains
// the JSON metadata for the struct
// [MeetingGetMeetingByIDResponseDataRecordingConfigLiveStreamingConfig]
type meetingGetMeetingByIDResponseDataRecordingConfigLiveStreamingConfigJSON struct {
	RtmpURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetMeetingByIDResponseDataRecordingConfigLiveStreamingConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingByIDResponseDataRecordingConfigLiveStreamingConfigJSON) RawJSON() string {
	return r.raw
}

type MeetingGetMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfig struct {
	// Controls whether recordings are uploaded to RealtimeKit's bucket. If set to
	// false, `download_url`, `audio_download_url`, `download_url_expiry` won't be
	// generated for a recording.
	Enabled bool                                                                        `json:"enabled,required"`
	JSON    meetingGetMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfigJSON `json:"-"`
}

// meetingGetMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfigJSON
// contains the JSON metadata for the struct
// [MeetingGetMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfig]
type meetingGetMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfigJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfigJSON) RawJSON() string {
	return r.raw
}

type MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfig struct {
	// Type of storage media.
	Type MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigType `json:"type,required"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// SSH destination server password for SFTP type storage medium when auth_method is
	// "PASSWORD". If auth_method is "KEY", this specifies the password for the ssh
	// private key.
	Password string `json:"password"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// Private key used to login to destination SSH server for SFTP type storage
	// medium, when auth_method used is "KEY"
	PrivateKey string `json:"private_key"`
	// Region of the storage medium.
	Region string `json:"region"`
	// Secret key of the storage medium. Similar to `access_key`, it is only writeable
	// by clients, not readable.
	Secret string `json:"secret"`
	// SSH destination server username for SFTP type storage medium
	Username string                                                            `json:"username"`
	JSON     meetingGetMeetingByIDResponseDataRecordingConfigStorageConfigJSON `json:"-"`
}

// meetingGetMeetingByIDResponseDataRecordingConfigStorageConfigJSON contains the
// JSON metadata for the struct
// [MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfig]
type meetingGetMeetingByIDResponseDataRecordingConfigStorageConfigJSON struct {
	Type        apijson.Field
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Password    apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	PrivateKey  apijson.Field
	Region      apijson.Field
	Secret      apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingByIDResponseDataRecordingConfigStorageConfigJSON) RawJSON() string {
	return r.raw
}

// Type of storage media.
type MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigType string

const (
	MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigTypeAws          MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigType = "aws"
	MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigTypeAzure        MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigType = "azure"
	MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigTypeDigitalocean MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigType = "digitalocean"
	MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigTypeGcs          MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigType = "gcs"
	MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigTypeSftp         MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigType = "sftp"
)

func (r MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigType) IsKnown() bool {
	switch r {
	case MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigTypeAws, MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigTypeAzure, MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigTypeDigitalocean, MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigTypeGcs, MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigTypeSftp:
		return true
	}
	return false
}

// Authentication method used for "sftp" type storage medium
type MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod string

const (
	MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethodKey      MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod = "KEY"
	MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethodPassword MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod = "PASSWORD"
)

func (r MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod) IsKnown() bool {
	switch r {
	case MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethodKey, MeetingGetMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethodPassword:
		return true
	}
	return false
}

type MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfig struct {
	// Codec using which the recording will be encoded.
	Codec MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigCodec `json:"codec"`
	// Controls whether to export video file seperately
	ExportFile bool `json:"export_file"`
	// Height of the recording video in pixels
	Height int64 `json:"height"`
	// Watermark to be added to the recording
	Watermark MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermark `json:"watermark"`
	// Width of the recording video in pixels
	Width int64                                                           `json:"width"`
	JSON  meetingGetMeetingByIDResponseDataRecordingConfigVideoConfigJSON `json:"-"`
}

// meetingGetMeetingByIDResponseDataRecordingConfigVideoConfigJSON contains the
// JSON metadata for the struct
// [MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfig]
type meetingGetMeetingByIDResponseDataRecordingConfigVideoConfigJSON struct {
	Codec       apijson.Field
	ExportFile  apijson.Field
	Height      apijson.Field
	Watermark   apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingByIDResponseDataRecordingConfigVideoConfigJSON) RawJSON() string {
	return r.raw
}

// Codec using which the recording will be encoded.
type MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigCodec string

const (
	MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigCodecH264 MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigCodec = "H264"
	MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigCodecVp8  MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigCodec = "VP8"
)

func (r MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigCodec) IsKnown() bool {
	switch r {
	case MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigCodecH264, MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigCodecVp8:
		return true
	}
	return false
}

// Watermark to be added to the recording
type MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermark struct {
	// Position of the watermark
	Position MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition `json:"position"`
	// Size of the watermark
	Size MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSize `json:"size"`
	// URL of the watermark image
	URL  string                                                                   `json:"url" format:"uri"`
	JSON meetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkJSON `json:"-"`
}

// meetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkJSON
// contains the JSON metadata for the struct
// [MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermark]
type meetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkJSON struct {
	Position    apijson.Field
	Size        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkJSON) RawJSON() string {
	return r.raw
}

// Position of the watermark
type MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition string

const (
	MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionLeftTop     MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition = "left top"
	MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionRightTop    MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition = "right top"
	MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionLeftBottom  MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition = "left bottom"
	MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionRightBottom MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition = "right bottom"
)

func (r MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition) IsKnown() bool {
	switch r {
	case MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionLeftTop, MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionRightTop, MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionLeftBottom, MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionRightBottom:
		return true
	}
	return false
}

// Size of the watermark
type MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSize struct {
	// Height of the watermark in px
	Height int64 `json:"height"`
	// Width of the watermark in px
	Width int64                                                                        `json:"width"`
	JSON  meetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSizeJSON `json:"-"`
}

// meetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSizeJSON
// contains the JSON metadata for the struct
// [MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSize]
type meetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSizeJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSizeJSON) RawJSON() string {
	return r.raw
}

// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
// `INACTIVE` meeting.
type MeetingGetMeetingByIDResponseDataStatus string

const (
	MeetingGetMeetingByIDResponseDataStatusActive   MeetingGetMeetingByIDResponseDataStatus = "ACTIVE"
	MeetingGetMeetingByIDResponseDataStatusInactive MeetingGetMeetingByIDResponseDataStatus = "INACTIVE"
)

func (r MeetingGetMeetingByIDResponseDataStatus) IsKnown() bool {
	switch r {
	case MeetingGetMeetingByIDResponseDataStatusActive, MeetingGetMeetingByIDResponseDataStatusInactive:
		return true
	}
	return false
}

type MeetingGetMeetingParticipantResponse struct {
	// Data returned by the operation
	Data MeetingGetMeetingParticipantResponseData `json:"data,required"`
	// Success status of the operation
	Success bool                                     `json:"success,required"`
	JSON    meetingGetMeetingParticipantResponseJSON `json:"-"`
}

// meetingGetMeetingParticipantResponseJSON contains the JSON metadata for the
// struct [MeetingGetMeetingParticipantResponse]
type meetingGetMeetingParticipantResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetMeetingParticipantResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingParticipantResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type MeetingGetMeetingParticipantResponseData struct {
	// ID of the participant.
	ID string `json:"id,required" format:"uuid"`
	// When this object was created. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A unique participant ID generated by the client.
	CustomParticipantID string `json:"custom_participant_id,required"`
	// Preset applied to the participant.
	PresetName string `json:"preset_name,required"`
	// When this object was updated. The time is returned in ISO format.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Name of the participant.
	Name string `json:"name,nullable"`
	// URL to a picture of the participant.
	Picture string                                       `json:"picture,nullable" format:"uri"`
	JSON    meetingGetMeetingParticipantResponseDataJSON `json:"-"`
}

// meetingGetMeetingParticipantResponseDataJSON contains the JSON metadata for the
// struct [MeetingGetMeetingParticipantResponseData]
type meetingGetMeetingParticipantResponseDataJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	CustomParticipantID apijson.Field
	PresetName          apijson.Field
	UpdatedAt           apijson.Field
	Name                apijson.Field
	Picture             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MeetingGetMeetingParticipantResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingParticipantResponseDataJSON) RawJSON() string {
	return r.raw
}

type MeetingGetMeetingParticipantsResponse struct {
	Data    []MeetingGetMeetingParticipantsResponseData `json:"data,required"`
	Paging  MeetingGetMeetingParticipantsResponsePaging `json:"paging,required"`
	Success bool                                        `json:"success,required"`
	JSON    meetingGetMeetingParticipantsResponseJSON   `json:"-"`
}

// meetingGetMeetingParticipantsResponseJSON contains the JSON metadata for the
// struct [MeetingGetMeetingParticipantsResponse]
type meetingGetMeetingParticipantsResponseJSON struct {
	Data        apijson.Field
	Paging      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetMeetingParticipantsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingParticipantsResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a participant.
type MeetingGetMeetingParticipantsResponseData struct {
	// ID of the participant.
	ID string `json:"id,required" format:"uuid"`
	// When this object was created. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A unique participant ID generated by the client.
	CustomParticipantID string `json:"custom_participant_id,required"`
	// Preset applied to the participant.
	PresetName string `json:"preset_name,required"`
	// When this object was updated. The time is returned in ISO format.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Name of the participant.
	Name string `json:"name,nullable"`
	// URL to a picture of the participant.
	Picture string                                        `json:"picture,nullable" format:"uri"`
	JSON    meetingGetMeetingParticipantsResponseDataJSON `json:"-"`
}

// meetingGetMeetingParticipantsResponseDataJSON contains the JSON metadata for the
// struct [MeetingGetMeetingParticipantsResponseData]
type meetingGetMeetingParticipantsResponseDataJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	CustomParticipantID apijson.Field
	PresetName          apijson.Field
	UpdatedAt           apijson.Field
	Name                apijson.Field
	Picture             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MeetingGetMeetingParticipantsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingParticipantsResponseDataJSON) RawJSON() string {
	return r.raw
}

type MeetingGetMeetingParticipantsResponsePaging struct {
	EndOffset   float64                                         `json:"end_offset,required"`
	StartOffset float64                                         `json:"start_offset,required"`
	TotalCount  float64                                         `json:"total_count,required"`
	JSON        meetingGetMeetingParticipantsResponsePagingJSON `json:"-"`
}

// meetingGetMeetingParticipantsResponsePagingJSON contains the JSON metadata for
// the struct [MeetingGetMeetingParticipantsResponsePaging]
type meetingGetMeetingParticipantsResponsePagingJSON struct {
	EndOffset   apijson.Field
	StartOffset apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingGetMeetingParticipantsResponsePaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingGetMeetingParticipantsResponsePagingJSON) RawJSON() string {
	return r.raw
}

type MeetingRefreshParticipantTokenResponse struct {
	// Data returned by the operation
	Data MeetingRefreshParticipantTokenResponseData `json:"data,required"`
	// Success status of the operation
	Success bool                                       `json:"success,required"`
	JSON    meetingRefreshParticipantTokenResponseJSON `json:"-"`
}

// meetingRefreshParticipantTokenResponseJSON contains the JSON metadata for the
// struct [MeetingRefreshParticipantTokenResponse]
type meetingRefreshParticipantTokenResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingRefreshParticipantTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingRefreshParticipantTokenResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type MeetingRefreshParticipantTokenResponseData struct {
	// Regenerated participant's authentication token.
	Token string                                         `json:"token,required"`
	JSON  meetingRefreshParticipantTokenResponseDataJSON `json:"-"`
}

// meetingRefreshParticipantTokenResponseDataJSON contains the JSON metadata for
// the struct [MeetingRefreshParticipantTokenResponseData]
type meetingRefreshParticipantTokenResponseDataJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingRefreshParticipantTokenResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingRefreshParticipantTokenResponseDataJSON) RawJSON() string {
	return r.raw
}

type MeetingReplaceMeetingByIDResponse struct {
	// Success status of the operation
	Success bool `json:"success,required"`
	// Data returned by the operation
	Data MeetingReplaceMeetingByIDResponseData `json:"data"`
	JSON meetingReplaceMeetingByIDResponseJSON `json:"-"`
}

// meetingReplaceMeetingByIDResponseJSON contains the JSON metadata for the struct
// [MeetingReplaceMeetingByIDResponse]
type meetingReplaceMeetingByIDResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingReplaceMeetingByIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingReplaceMeetingByIDResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type MeetingReplaceMeetingByIDResponseData struct {
	// ID of the meeting.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp the object was created at. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Timestamp the object was updated at. The time is returned in ISO format.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The AI Config allows you to customize the behavior of meeting transcriptions and
	// summaries
	AIConfig MeetingReplaceMeetingByIDResponseDataAIConfig `json:"ai_config"`
	// Specifies if the meeting should start getting livestreamed on start.
	LiveStreamOnStart bool `json:"live_stream_on_start"`
	// Specifies if Chat within a meeting should persist for a week.
	PersistChat bool `json:"persist_chat"`
	// Specifies if the meeting should start getting recorded as soon as someone joins
	// the meeting.
	RecordOnStart bool `json:"record_on_start"`
	// Recording Configurations to be used for this meeting. This level of configs
	// takes higher preference over App level configs on the RealtimeKit developer
	// portal.
	RecordingConfig MeetingReplaceMeetingByIDResponseDataRecordingConfig `json:"recording_config"`
	// Time in seconds, for which a session remains active, after the last participant
	// has left the meeting.
	SessionKeepAliveTimeInSecs float64 `json:"session_keep_alive_time_in_secs"`
	// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
	// `INACTIVE` meeting.
	Status MeetingReplaceMeetingByIDResponseDataStatus `json:"status"`
	// Automatically generate summary of meetings using transcripts. Requires
	// Transcriptions to be enabled, and can be retrieved via Webhooks or summary API.
	SummarizeOnEnd bool `json:"summarize_on_end"`
	// Title of the meeting.
	Title string                                    `json:"title"`
	JSON  meetingReplaceMeetingByIDResponseDataJSON `json:"-"`
}

// meetingReplaceMeetingByIDResponseDataJSON contains the JSON metadata for the
// struct [MeetingReplaceMeetingByIDResponseData]
type meetingReplaceMeetingByIDResponseDataJSON struct {
	ID                         apijson.Field
	CreatedAt                  apijson.Field
	UpdatedAt                  apijson.Field
	AIConfig                   apijson.Field
	LiveStreamOnStart          apijson.Field
	PersistChat                apijson.Field
	RecordOnStart              apijson.Field
	RecordingConfig            apijson.Field
	SessionKeepAliveTimeInSecs apijson.Field
	Status                     apijson.Field
	SummarizeOnEnd             apijson.Field
	Title                      apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *MeetingReplaceMeetingByIDResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingReplaceMeetingByIDResponseDataJSON) RawJSON() string {
	return r.raw
}

// The AI Config allows you to customize the behavior of meeting transcriptions and
// summaries
type MeetingReplaceMeetingByIDResponseDataAIConfig struct {
	// Summary Config
	Summarization MeetingReplaceMeetingByIDResponseDataAIConfigSummarization `json:"summarization"`
	// Transcription Configurations
	Transcription MeetingReplaceMeetingByIDResponseDataAIConfigTranscription `json:"transcription"`
	JSON          meetingReplaceMeetingByIDResponseDataAIConfigJSON          `json:"-"`
}

// meetingReplaceMeetingByIDResponseDataAIConfigJSON contains the JSON metadata for
// the struct [MeetingReplaceMeetingByIDResponseDataAIConfig]
type meetingReplaceMeetingByIDResponseDataAIConfigJSON struct {
	Summarization apijson.Field
	Transcription apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MeetingReplaceMeetingByIDResponseDataAIConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingReplaceMeetingByIDResponseDataAIConfigJSON) RawJSON() string {
	return r.raw
}

// Summary Config
type MeetingReplaceMeetingByIDResponseDataAIConfigSummarization struct {
	// Defines the style of the summary, such as general, team meeting, or sales call.
	SummaryType MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryType `json:"summary_type"`
	// Determines the text format of the summary, such as plain text or markdown.
	TextFormat MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationTextFormat `json:"text_format"`
	// Sets the maximum number of words in the meeting summary.
	WordLimit int64                                                          `json:"word_limit"`
	JSON      meetingReplaceMeetingByIDResponseDataAIConfigSummarizationJSON `json:"-"`
}

// meetingReplaceMeetingByIDResponseDataAIConfigSummarizationJSON contains the JSON
// metadata for the struct
// [MeetingReplaceMeetingByIDResponseDataAIConfigSummarization]
type meetingReplaceMeetingByIDResponseDataAIConfigSummarizationJSON struct {
	SummaryType apijson.Field
	TextFormat  apijson.Field
	WordLimit   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingReplaceMeetingByIDResponseDataAIConfigSummarization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingReplaceMeetingByIDResponseDataAIConfigSummarizationJSON) RawJSON() string {
	return r.raw
}

// Defines the style of the summary, such as general, team meeting, or sales call.
type MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryType string

const (
	MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeGeneral         MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryType = "general"
	MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeTeamMeeting     MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryType = "team_meeting"
	MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeSalesCall       MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryType = "sales_call"
	MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeClientCheckIn   MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryType = "client_check_in"
	MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeInterview       MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryType = "interview"
	MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeDailyStandup    MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryType = "daily_standup"
	MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeOneOnOneMeeting MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryType = "one_on_one_meeting"
	MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeLecture         MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryType = "lecture"
	MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeCodeReview      MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryType = "code_review"
)

func (r MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryType) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeGeneral, MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeTeamMeeting, MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeSalesCall, MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeClientCheckIn, MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeInterview, MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeDailyStandup, MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeOneOnOneMeeting, MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeLecture, MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationSummaryTypeCodeReview:
		return true
	}
	return false
}

// Determines the text format of the summary, such as plain text or markdown.
type MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationTextFormat string

const (
	MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationTextFormatPlainText MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationTextFormat = "plain_text"
	MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationTextFormatMarkdown  MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationTextFormat = "markdown"
)

func (r MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationTextFormat) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationTextFormatPlainText, MeetingReplaceMeetingByIDResponseDataAIConfigSummarizationTextFormatMarkdown:
		return true
	}
	return false
}

// Transcription Configurations
type MeetingReplaceMeetingByIDResponseDataAIConfigTranscription struct {
	// Adds specific terms to improve accurate detection during transcription.
	Keywords []string `json:"keywords"`
	// Specifies the language code for transcription to ensure accurate results.
	Language MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguage `json:"language"`
	// Control the inclusion of offensive language in transcriptions.
	ProfanityFilter bool                                                           `json:"profanity_filter"`
	JSON            meetingReplaceMeetingByIDResponseDataAIConfigTranscriptionJSON `json:"-"`
}

// meetingReplaceMeetingByIDResponseDataAIConfigTranscriptionJSON contains the JSON
// metadata for the struct
// [MeetingReplaceMeetingByIDResponseDataAIConfigTranscription]
type meetingReplaceMeetingByIDResponseDataAIConfigTranscriptionJSON struct {
	Keywords        apijson.Field
	Language        apijson.Field
	ProfanityFilter apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MeetingReplaceMeetingByIDResponseDataAIConfigTranscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingReplaceMeetingByIDResponseDataAIConfigTranscriptionJSON) RawJSON() string {
	return r.raw
}

// Specifies the language code for transcription to ensure accurate results.
type MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguage string

const (
	MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageEnUs MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguage = "en-US"
	MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageEnIn MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguage = "en-IN"
	MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageDe   MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguage = "de"
	MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageHi   MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguage = "hi"
	MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageSv   MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguage = "sv"
	MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageRu   MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguage = "ru"
	MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguagePl   MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguage = "pl"
	MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageEl   MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguage = "el"
	MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageFr   MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguage = "fr"
	MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageNl   MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguage = "nl"
)

func (r MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguage) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageEnUs, MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageEnIn, MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageDe, MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageHi, MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageSv, MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageRu, MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguagePl, MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageEl, MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageFr, MeetingReplaceMeetingByIDResponseDataAIConfigTranscriptionLanguageNl:
		return true
	}
	return false
}

// Recording Configurations to be used for this meeting. This level of configs
// takes higher preference over App level configs on the RealtimeKit developer
// portal.
type MeetingReplaceMeetingByIDResponseDataRecordingConfig struct {
	// Object containing configuration regarding the audio that is being recorded.
	AudioConfig MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfig `json:"audio_config"`
	// Adds a prefix to the beginning of the file name of the recording.
	FileNamePrefix      string                                                                  `json:"file_name_prefix"`
	LiveStreamingConfig MeetingReplaceMeetingByIDResponseDataRecordingConfigLiveStreamingConfig `json:"live_streaming_config"`
	// Specifies the maximum duration for recording in seconds, ranging from a minimum
	// of 60 seconds to a maximum of 24 hours.
	MaxSeconds              float64                                                                     `json:"max_seconds"`
	RealtimekitBucketConfig MeetingReplaceMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfig `json:"realtimekit_bucket_config"`
	StorageConfig           MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfig           `json:"storage_config,nullable"`
	VideoConfig             MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfig             `json:"video_config"`
	JSON                    meetingReplaceMeetingByIDResponseDataRecordingConfigJSON                    `json:"-"`
}

// meetingReplaceMeetingByIDResponseDataRecordingConfigJSON contains the JSON
// metadata for the struct [MeetingReplaceMeetingByIDResponseDataRecordingConfig]
type meetingReplaceMeetingByIDResponseDataRecordingConfigJSON struct {
	AudioConfig             apijson.Field
	FileNamePrefix          apijson.Field
	LiveStreamingConfig     apijson.Field
	MaxSeconds              apijson.Field
	RealtimekitBucketConfig apijson.Field
	StorageConfig           apijson.Field
	VideoConfig             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *MeetingReplaceMeetingByIDResponseDataRecordingConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingReplaceMeetingByIDResponseDataRecordingConfigJSON) RawJSON() string {
	return r.raw
}

// Object containing configuration regarding the audio that is being recorded.
type MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfig struct {
	// Audio signal pathway within an audio file that carries a specific sound source.
	Channel MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigChannel `json:"channel"`
	// Codec using which the recording will be encoded. If VP8/VP9 is selected for
	// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
	// audioConfig is automatically set to vorbis.
	Codec MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigCodec `json:"codec"`
	// Controls whether to export audio file seperately
	ExportFile bool                                                                `json:"export_file"`
	JSON       meetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigJSON `json:"-"`
}

// meetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigJSON contains the
// JSON metadata for the struct
// [MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfig]
type meetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigJSON struct {
	Channel     apijson.Field
	Codec       apijson.Field
	ExportFile  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigJSON) RawJSON() string {
	return r.raw
}

// Audio signal pathway within an audio file that carries a specific sound source.
type MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigChannel string

const (
	MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigChannelMono   MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigChannel = "mono"
	MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigChannelStereo MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigChannel = "stereo"
)

func (r MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigChannel) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigChannelMono, MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigChannelStereo:
		return true
	}
	return false
}

// Codec using which the recording will be encoded. If VP8/VP9 is selected for
// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
// audioConfig is automatically set to vorbis.
type MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigCodec string

const (
	MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigCodecMP3 MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigCodec = "MP3"
	MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigCodecAac MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigCodec = "AAC"
)

func (r MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigCodec) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigCodecMP3, MeetingReplaceMeetingByIDResponseDataRecordingConfigAudioConfigCodecAac:
		return true
	}
	return false
}

type MeetingReplaceMeetingByIDResponseDataRecordingConfigLiveStreamingConfig struct {
	// RTMP URL to stream to
	RtmpURL string                                                                      `json:"rtmp_url" format:"uri"`
	JSON    meetingReplaceMeetingByIDResponseDataRecordingConfigLiveStreamingConfigJSON `json:"-"`
}

// meetingReplaceMeetingByIDResponseDataRecordingConfigLiveStreamingConfigJSON
// contains the JSON metadata for the struct
// [MeetingReplaceMeetingByIDResponseDataRecordingConfigLiveStreamingConfig]
type meetingReplaceMeetingByIDResponseDataRecordingConfigLiveStreamingConfigJSON struct {
	RtmpURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingReplaceMeetingByIDResponseDataRecordingConfigLiveStreamingConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingReplaceMeetingByIDResponseDataRecordingConfigLiveStreamingConfigJSON) RawJSON() string {
	return r.raw
}

type MeetingReplaceMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfig struct {
	// Controls whether recordings are uploaded to RealtimeKit's bucket. If set to
	// false, `download_url`, `audio_download_url`, `download_url_expiry` won't be
	// generated for a recording.
	Enabled bool                                                                            `json:"enabled,required"`
	JSON    meetingReplaceMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfigJSON `json:"-"`
}

// meetingReplaceMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfigJSON
// contains the JSON metadata for the struct
// [MeetingReplaceMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfig]
type meetingReplaceMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfigJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingReplaceMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingReplaceMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfigJSON) RawJSON() string {
	return r.raw
}

type MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfig struct {
	// Type of storage media.
	Type MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigType `json:"type,required"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// SSH destination server password for SFTP type storage medium when auth_method is
	// "PASSWORD". If auth_method is "KEY", this specifies the password for the ssh
	// private key.
	Password string `json:"password"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// Private key used to login to destination SSH server for SFTP type storage
	// medium, when auth_method used is "KEY"
	PrivateKey string `json:"private_key"`
	// Region of the storage medium.
	Region string `json:"region"`
	// Secret key of the storage medium. Similar to `access_key`, it is only writeable
	// by clients, not readable.
	Secret string `json:"secret"`
	// SSH destination server username for SFTP type storage medium
	Username string                                                                `json:"username"`
	JSON     meetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigJSON `json:"-"`
}

// meetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigJSON contains
// the JSON metadata for the struct
// [MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfig]
type meetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigJSON struct {
	Type        apijson.Field
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Password    apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	PrivateKey  apijson.Field
	Region      apijson.Field
	Secret      apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigJSON) RawJSON() string {
	return r.raw
}

// Type of storage media.
type MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigType string

const (
	MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigTypeAws          MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigType = "aws"
	MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigTypeAzure        MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigType = "azure"
	MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigTypeDigitalocean MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigType = "digitalocean"
	MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigTypeGcs          MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigType = "gcs"
	MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigTypeSftp         MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigType = "sftp"
)

func (r MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigType) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigTypeAws, MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigTypeAzure, MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigTypeDigitalocean, MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigTypeGcs, MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigTypeSftp:
		return true
	}
	return false
}

// Authentication method used for "sftp" type storage medium
type MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod string

const (
	MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethodKey      MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod = "KEY"
	MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethodPassword MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod = "PASSWORD"
)

func (r MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethodKey, MeetingReplaceMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethodPassword:
		return true
	}
	return false
}

type MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfig struct {
	// Codec using which the recording will be encoded.
	Codec MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigCodec `json:"codec"`
	// Controls whether to export video file seperately
	ExportFile bool `json:"export_file"`
	// Height of the recording video in pixels
	Height int64 `json:"height"`
	// Watermark to be added to the recording
	Watermark MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermark `json:"watermark"`
	// Width of the recording video in pixels
	Width int64                                                               `json:"width"`
	JSON  meetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigJSON `json:"-"`
}

// meetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigJSON contains the
// JSON metadata for the struct
// [MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfig]
type meetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigJSON struct {
	Codec       apijson.Field
	ExportFile  apijson.Field
	Height      apijson.Field
	Watermark   apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigJSON) RawJSON() string {
	return r.raw
}

// Codec using which the recording will be encoded.
type MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigCodec string

const (
	MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigCodecH264 MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigCodec = "H264"
	MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigCodecVp8  MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigCodec = "VP8"
)

func (r MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigCodec) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigCodecH264, MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigCodecVp8:
		return true
	}
	return false
}

// Watermark to be added to the recording
type MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermark struct {
	// Position of the watermark
	Position MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition `json:"position"`
	// Size of the watermark
	Size MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSize `json:"size"`
	// URL of the watermark image
	URL  string                                                                       `json:"url" format:"uri"`
	JSON meetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkJSON `json:"-"`
}

// meetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkJSON
// contains the JSON metadata for the struct
// [MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermark]
type meetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkJSON struct {
	Position    apijson.Field
	Size        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkJSON) RawJSON() string {
	return r.raw
}

// Position of the watermark
type MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition string

const (
	MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionLeftTop     MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition = "left top"
	MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionRightTop    MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition = "right top"
	MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionLeftBottom  MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition = "left bottom"
	MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionRightBottom MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition = "right bottom"
)

func (r MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionLeftTop, MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionRightTop, MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionLeftBottom, MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionRightBottom:
		return true
	}
	return false
}

// Size of the watermark
type MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSize struct {
	// Height of the watermark in px
	Height int64 `json:"height"`
	// Width of the watermark in px
	Width int64                                                                            `json:"width"`
	JSON  meetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSizeJSON `json:"-"`
}

// meetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSizeJSON
// contains the JSON metadata for the struct
// [MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSize]
type meetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSizeJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingReplaceMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSizeJSON) RawJSON() string {
	return r.raw
}

// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
// `INACTIVE` meeting.
type MeetingReplaceMeetingByIDResponseDataStatus string

const (
	MeetingReplaceMeetingByIDResponseDataStatusActive   MeetingReplaceMeetingByIDResponseDataStatus = "ACTIVE"
	MeetingReplaceMeetingByIDResponseDataStatusInactive MeetingReplaceMeetingByIDResponseDataStatus = "INACTIVE"
)

func (r MeetingReplaceMeetingByIDResponseDataStatus) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDResponseDataStatusActive, MeetingReplaceMeetingByIDResponseDataStatusInactive:
		return true
	}
	return false
}

type MeetingUpdateMeetingByIDResponse struct {
	// Success status of the operation
	Success bool `json:"success,required"`
	// Data returned by the operation
	Data MeetingUpdateMeetingByIDResponseData `json:"data"`
	JSON meetingUpdateMeetingByIDResponseJSON `json:"-"`
}

// meetingUpdateMeetingByIDResponseJSON contains the JSON metadata for the struct
// [MeetingUpdateMeetingByIDResponse]
type meetingUpdateMeetingByIDResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingUpdateMeetingByIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingUpdateMeetingByIDResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type MeetingUpdateMeetingByIDResponseData struct {
	// ID of the meeting.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp the object was created at. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Timestamp the object was updated at. The time is returned in ISO format.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The AI Config allows you to customize the behavior of meeting transcriptions and
	// summaries
	AIConfig MeetingUpdateMeetingByIDResponseDataAIConfig `json:"ai_config"`
	// Specifies if the meeting should start getting livestreamed on start.
	LiveStreamOnStart bool `json:"live_stream_on_start"`
	// Specifies if Chat within a meeting should persist for a week.
	PersistChat bool `json:"persist_chat"`
	// Specifies if the meeting should start getting recorded as soon as someone joins
	// the meeting.
	RecordOnStart bool `json:"record_on_start"`
	// Recording Configurations to be used for this meeting. This level of configs
	// takes higher preference over App level configs on the RealtimeKit developer
	// portal.
	RecordingConfig MeetingUpdateMeetingByIDResponseDataRecordingConfig `json:"recording_config"`
	// Time in seconds, for which a session remains active, after the last participant
	// has left the meeting.
	SessionKeepAliveTimeInSecs float64 `json:"session_keep_alive_time_in_secs"`
	// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
	// `INACTIVE` meeting.
	Status MeetingUpdateMeetingByIDResponseDataStatus `json:"status"`
	// Automatically generate summary of meetings using transcripts. Requires
	// Transcriptions to be enabled, and can be retrieved via Webhooks or summary API.
	SummarizeOnEnd bool `json:"summarize_on_end"`
	// Title of the meeting.
	Title string                                   `json:"title"`
	JSON  meetingUpdateMeetingByIDResponseDataJSON `json:"-"`
}

// meetingUpdateMeetingByIDResponseDataJSON contains the JSON metadata for the
// struct [MeetingUpdateMeetingByIDResponseData]
type meetingUpdateMeetingByIDResponseDataJSON struct {
	ID                         apijson.Field
	CreatedAt                  apijson.Field
	UpdatedAt                  apijson.Field
	AIConfig                   apijson.Field
	LiveStreamOnStart          apijson.Field
	PersistChat                apijson.Field
	RecordOnStart              apijson.Field
	RecordingConfig            apijson.Field
	SessionKeepAliveTimeInSecs apijson.Field
	Status                     apijson.Field
	SummarizeOnEnd             apijson.Field
	Title                      apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *MeetingUpdateMeetingByIDResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingUpdateMeetingByIDResponseDataJSON) RawJSON() string {
	return r.raw
}

// The AI Config allows you to customize the behavior of meeting transcriptions and
// summaries
type MeetingUpdateMeetingByIDResponseDataAIConfig struct {
	// Summary Config
	Summarization MeetingUpdateMeetingByIDResponseDataAIConfigSummarization `json:"summarization"`
	// Transcription Configurations
	Transcription MeetingUpdateMeetingByIDResponseDataAIConfigTranscription `json:"transcription"`
	JSON          meetingUpdateMeetingByIDResponseDataAIConfigJSON          `json:"-"`
}

// meetingUpdateMeetingByIDResponseDataAIConfigJSON contains the JSON metadata for
// the struct [MeetingUpdateMeetingByIDResponseDataAIConfig]
type meetingUpdateMeetingByIDResponseDataAIConfigJSON struct {
	Summarization apijson.Field
	Transcription apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MeetingUpdateMeetingByIDResponseDataAIConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingUpdateMeetingByIDResponseDataAIConfigJSON) RawJSON() string {
	return r.raw
}

// Summary Config
type MeetingUpdateMeetingByIDResponseDataAIConfigSummarization struct {
	// Defines the style of the summary, such as general, team meeting, or sales call.
	SummaryType MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryType `json:"summary_type"`
	// Determines the text format of the summary, such as plain text or markdown.
	TextFormat MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationTextFormat `json:"text_format"`
	// Sets the maximum number of words in the meeting summary.
	WordLimit int64                                                         `json:"word_limit"`
	JSON      meetingUpdateMeetingByIDResponseDataAIConfigSummarizationJSON `json:"-"`
}

// meetingUpdateMeetingByIDResponseDataAIConfigSummarizationJSON contains the JSON
// metadata for the struct
// [MeetingUpdateMeetingByIDResponseDataAIConfigSummarization]
type meetingUpdateMeetingByIDResponseDataAIConfigSummarizationJSON struct {
	SummaryType apijson.Field
	TextFormat  apijson.Field
	WordLimit   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingUpdateMeetingByIDResponseDataAIConfigSummarization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingUpdateMeetingByIDResponseDataAIConfigSummarizationJSON) RawJSON() string {
	return r.raw
}

// Defines the style of the summary, such as general, team meeting, or sales call.
type MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryType string

const (
	MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeGeneral         MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryType = "general"
	MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeTeamMeeting     MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryType = "team_meeting"
	MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeSalesCall       MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryType = "sales_call"
	MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeClientCheckIn   MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryType = "client_check_in"
	MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeInterview       MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryType = "interview"
	MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeDailyStandup    MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryType = "daily_standup"
	MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeOneOnOneMeeting MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryType = "one_on_one_meeting"
	MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeLecture         MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryType = "lecture"
	MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeCodeReview      MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryType = "code_review"
)

func (r MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryType) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeGeneral, MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeTeamMeeting, MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeSalesCall, MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeClientCheckIn, MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeInterview, MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeDailyStandup, MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeOneOnOneMeeting, MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeLecture, MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationSummaryTypeCodeReview:
		return true
	}
	return false
}

// Determines the text format of the summary, such as plain text or markdown.
type MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationTextFormat string

const (
	MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationTextFormatPlainText MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationTextFormat = "plain_text"
	MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationTextFormatMarkdown  MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationTextFormat = "markdown"
)

func (r MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationTextFormat) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationTextFormatPlainText, MeetingUpdateMeetingByIDResponseDataAIConfigSummarizationTextFormatMarkdown:
		return true
	}
	return false
}

// Transcription Configurations
type MeetingUpdateMeetingByIDResponseDataAIConfigTranscription struct {
	// Adds specific terms to improve accurate detection during transcription.
	Keywords []string `json:"keywords"`
	// Specifies the language code for transcription to ensure accurate results.
	Language MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguage `json:"language"`
	// Control the inclusion of offensive language in transcriptions.
	ProfanityFilter bool                                                          `json:"profanity_filter"`
	JSON            meetingUpdateMeetingByIDResponseDataAIConfigTranscriptionJSON `json:"-"`
}

// meetingUpdateMeetingByIDResponseDataAIConfigTranscriptionJSON contains the JSON
// metadata for the struct
// [MeetingUpdateMeetingByIDResponseDataAIConfigTranscription]
type meetingUpdateMeetingByIDResponseDataAIConfigTranscriptionJSON struct {
	Keywords        apijson.Field
	Language        apijson.Field
	ProfanityFilter apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MeetingUpdateMeetingByIDResponseDataAIConfigTranscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingUpdateMeetingByIDResponseDataAIConfigTranscriptionJSON) RawJSON() string {
	return r.raw
}

// Specifies the language code for transcription to ensure accurate results.
type MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguage string

const (
	MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageEnUs MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguage = "en-US"
	MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageEnIn MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguage = "en-IN"
	MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageDe   MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguage = "de"
	MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageHi   MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguage = "hi"
	MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageSv   MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguage = "sv"
	MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageRu   MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguage = "ru"
	MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguagePl   MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguage = "pl"
	MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageEl   MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguage = "el"
	MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageFr   MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguage = "fr"
	MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageNl   MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguage = "nl"
)

func (r MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguage) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageEnUs, MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageEnIn, MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageDe, MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageHi, MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageSv, MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageRu, MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguagePl, MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageEl, MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageFr, MeetingUpdateMeetingByIDResponseDataAIConfigTranscriptionLanguageNl:
		return true
	}
	return false
}

// Recording Configurations to be used for this meeting. This level of configs
// takes higher preference over App level configs on the RealtimeKit developer
// portal.
type MeetingUpdateMeetingByIDResponseDataRecordingConfig struct {
	// Object containing configuration regarding the audio that is being recorded.
	AudioConfig MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfig `json:"audio_config"`
	// Adds a prefix to the beginning of the file name of the recording.
	FileNamePrefix      string                                                                 `json:"file_name_prefix"`
	LiveStreamingConfig MeetingUpdateMeetingByIDResponseDataRecordingConfigLiveStreamingConfig `json:"live_streaming_config"`
	// Specifies the maximum duration for recording in seconds, ranging from a minimum
	// of 60 seconds to a maximum of 24 hours.
	MaxSeconds              float64                                                                    `json:"max_seconds"`
	RealtimekitBucketConfig MeetingUpdateMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfig `json:"realtimekit_bucket_config"`
	StorageConfig           MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfig           `json:"storage_config,nullable"`
	VideoConfig             MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfig             `json:"video_config"`
	JSON                    meetingUpdateMeetingByIDResponseDataRecordingConfigJSON                    `json:"-"`
}

// meetingUpdateMeetingByIDResponseDataRecordingConfigJSON contains the JSON
// metadata for the struct [MeetingUpdateMeetingByIDResponseDataRecordingConfig]
type meetingUpdateMeetingByIDResponseDataRecordingConfigJSON struct {
	AudioConfig             apijson.Field
	FileNamePrefix          apijson.Field
	LiveStreamingConfig     apijson.Field
	MaxSeconds              apijson.Field
	RealtimekitBucketConfig apijson.Field
	StorageConfig           apijson.Field
	VideoConfig             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *MeetingUpdateMeetingByIDResponseDataRecordingConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingUpdateMeetingByIDResponseDataRecordingConfigJSON) RawJSON() string {
	return r.raw
}

// Object containing configuration regarding the audio that is being recorded.
type MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfig struct {
	// Audio signal pathway within an audio file that carries a specific sound source.
	Channel MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigChannel `json:"channel"`
	// Codec using which the recording will be encoded. If VP8/VP9 is selected for
	// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
	// audioConfig is automatically set to vorbis.
	Codec MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigCodec `json:"codec"`
	// Controls whether to export audio file seperately
	ExportFile bool                                                               `json:"export_file"`
	JSON       meetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigJSON `json:"-"`
}

// meetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigJSON contains the
// JSON metadata for the struct
// [MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfig]
type meetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigJSON struct {
	Channel     apijson.Field
	Codec       apijson.Field
	ExportFile  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigJSON) RawJSON() string {
	return r.raw
}

// Audio signal pathway within an audio file that carries a specific sound source.
type MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigChannel string

const (
	MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigChannelMono   MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigChannel = "mono"
	MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigChannelStereo MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigChannel = "stereo"
)

func (r MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigChannel) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigChannelMono, MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigChannelStereo:
		return true
	}
	return false
}

// Codec using which the recording will be encoded. If VP8/VP9 is selected for
// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
// audioConfig is automatically set to vorbis.
type MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigCodec string

const (
	MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigCodecMP3 MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigCodec = "MP3"
	MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigCodecAac MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigCodec = "AAC"
)

func (r MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigCodec) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigCodecMP3, MeetingUpdateMeetingByIDResponseDataRecordingConfigAudioConfigCodecAac:
		return true
	}
	return false
}

type MeetingUpdateMeetingByIDResponseDataRecordingConfigLiveStreamingConfig struct {
	// RTMP URL to stream to
	RtmpURL string                                                                     `json:"rtmp_url" format:"uri"`
	JSON    meetingUpdateMeetingByIDResponseDataRecordingConfigLiveStreamingConfigJSON `json:"-"`
}

// meetingUpdateMeetingByIDResponseDataRecordingConfigLiveStreamingConfigJSON
// contains the JSON metadata for the struct
// [MeetingUpdateMeetingByIDResponseDataRecordingConfigLiveStreamingConfig]
type meetingUpdateMeetingByIDResponseDataRecordingConfigLiveStreamingConfigJSON struct {
	RtmpURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingUpdateMeetingByIDResponseDataRecordingConfigLiveStreamingConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingUpdateMeetingByIDResponseDataRecordingConfigLiveStreamingConfigJSON) RawJSON() string {
	return r.raw
}

type MeetingUpdateMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfig struct {
	// Controls whether recordings are uploaded to RealtimeKit's bucket. If set to
	// false, `download_url`, `audio_download_url`, `download_url_expiry` won't be
	// generated for a recording.
	Enabled bool                                                                           `json:"enabled,required"`
	JSON    meetingUpdateMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfigJSON `json:"-"`
}

// meetingUpdateMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfigJSON
// contains the JSON metadata for the struct
// [MeetingUpdateMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfig]
type meetingUpdateMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfigJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingUpdateMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingUpdateMeetingByIDResponseDataRecordingConfigRealtimekitBucketConfigJSON) RawJSON() string {
	return r.raw
}

type MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfig struct {
	// Type of storage media.
	Type MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigType `json:"type,required"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// SSH destination server password for SFTP type storage medium when auth_method is
	// "PASSWORD". If auth_method is "KEY", this specifies the password for the ssh
	// private key.
	Password string `json:"password"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// Private key used to login to destination SSH server for SFTP type storage
	// medium, when auth_method used is "KEY"
	PrivateKey string `json:"private_key"`
	// Region of the storage medium.
	Region string `json:"region"`
	// Secret key of the storage medium. Similar to `access_key`, it is only writeable
	// by clients, not readable.
	Secret string `json:"secret"`
	// SSH destination server username for SFTP type storage medium
	Username string                                                               `json:"username"`
	JSON     meetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigJSON `json:"-"`
}

// meetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigJSON contains
// the JSON metadata for the struct
// [MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfig]
type meetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigJSON struct {
	Type        apijson.Field
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Password    apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	PrivateKey  apijson.Field
	Region      apijson.Field
	Secret      apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigJSON) RawJSON() string {
	return r.raw
}

// Type of storage media.
type MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigType string

const (
	MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigTypeAws          MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigType = "aws"
	MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigTypeAzure        MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigType = "azure"
	MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigTypeDigitalocean MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigType = "digitalocean"
	MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigTypeGcs          MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigType = "gcs"
	MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigTypeSftp         MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigType = "sftp"
)

func (r MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigType) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigTypeAws, MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigTypeAzure, MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigTypeDigitalocean, MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigTypeGcs, MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigTypeSftp:
		return true
	}
	return false
}

// Authentication method used for "sftp" type storage medium
type MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod string

const (
	MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethodKey      MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod = "KEY"
	MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethodPassword MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod = "PASSWORD"
)

func (r MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethod) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethodKey, MeetingUpdateMeetingByIDResponseDataRecordingConfigStorageConfigAuthMethodPassword:
		return true
	}
	return false
}

type MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfig struct {
	// Codec using which the recording will be encoded.
	Codec MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigCodec `json:"codec"`
	// Controls whether to export video file seperately
	ExportFile bool `json:"export_file"`
	// Height of the recording video in pixels
	Height int64 `json:"height"`
	// Watermark to be added to the recording
	Watermark MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermark `json:"watermark"`
	// Width of the recording video in pixels
	Width int64                                                              `json:"width"`
	JSON  meetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigJSON `json:"-"`
}

// meetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigJSON contains the
// JSON metadata for the struct
// [MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfig]
type meetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigJSON struct {
	Codec       apijson.Field
	ExportFile  apijson.Field
	Height      apijson.Field
	Watermark   apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigJSON) RawJSON() string {
	return r.raw
}

// Codec using which the recording will be encoded.
type MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigCodec string

const (
	MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigCodecH264 MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigCodec = "H264"
	MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigCodecVp8  MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigCodec = "VP8"
)

func (r MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigCodec) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigCodecH264, MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigCodecVp8:
		return true
	}
	return false
}

// Watermark to be added to the recording
type MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermark struct {
	// Position of the watermark
	Position MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition `json:"position"`
	// Size of the watermark
	Size MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSize `json:"size"`
	// URL of the watermark image
	URL  string                                                                      `json:"url" format:"uri"`
	JSON meetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkJSON `json:"-"`
}

// meetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkJSON
// contains the JSON metadata for the struct
// [MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermark]
type meetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkJSON struct {
	Position    apijson.Field
	Size        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkJSON) RawJSON() string {
	return r.raw
}

// Position of the watermark
type MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition string

const (
	MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionLeftTop     MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition = "left top"
	MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionRightTop    MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition = "right top"
	MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionLeftBottom  MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition = "left bottom"
	MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionRightBottom MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition = "right bottom"
)

func (r MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPosition) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionLeftTop, MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionRightTop, MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionLeftBottom, MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkPositionRightBottom:
		return true
	}
	return false
}

// Size of the watermark
type MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSize struct {
	// Height of the watermark in px
	Height int64 `json:"height"`
	// Width of the watermark in px
	Width int64                                                                           `json:"width"`
	JSON  meetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSizeJSON `json:"-"`
}

// meetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSizeJSON
// contains the JSON metadata for the struct
// [MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSize]
type meetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSizeJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meetingUpdateMeetingByIDResponseDataRecordingConfigVideoConfigWatermarkSizeJSON) RawJSON() string {
	return r.raw
}

// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
// `INACTIVE` meeting.
type MeetingUpdateMeetingByIDResponseDataStatus string

const (
	MeetingUpdateMeetingByIDResponseDataStatusActive   MeetingUpdateMeetingByIDResponseDataStatus = "ACTIVE"
	MeetingUpdateMeetingByIDResponseDataStatusInactive MeetingUpdateMeetingByIDResponseDataStatus = "INACTIVE"
)

func (r MeetingUpdateMeetingByIDResponseDataStatus) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDResponseDataStatusActive, MeetingUpdateMeetingByIDResponseDataStatusInactive:
		return true
	}
	return false
}

type MeetingNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The AI Config allows you to customize the behavior of meeting transcriptions and
	// summaries
	AIConfig param.Field[MeetingNewParamsAIConfig] `json:"ai_config"`
	// Specifies if the meeting should start getting livestreamed on start.
	LiveStreamOnStart param.Field[bool] `json:"live_stream_on_start"`
	// If a meeting is set to persist_chat, meeting chat would remain for a week within
	// the meeting space.
	PersistChat param.Field[bool] `json:"persist_chat"`
	// Specifies if the meeting should start getting recorded as soon as someone joins
	// the meeting.
	RecordOnStart param.Field[bool] `json:"record_on_start"`
	// Recording Configurations to be used for this meeting. This level of configs
	// takes higher preference over App level configs on the RealtimeKit developer
	// portal.
	RecordingConfig param.Field[MeetingNewParamsRecordingConfig] `json:"recording_config"`
	// Time in seconds, for which a session remains active, after the last participant
	// has left the meeting.
	SessionKeepAliveTimeInSecs param.Field[float64] `json:"session_keep_alive_time_in_secs"`
	// Automatically generate summary of meetings using transcripts. Requires
	// Transcriptions to be enabled, and can be retrieved via Webhooks or summary API.
	SummarizeOnEnd param.Field[bool] `json:"summarize_on_end"`
	// Title of the meeting
	Title param.Field[string] `json:"title"`
}

func (r MeetingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The AI Config allows you to customize the behavior of meeting transcriptions and
// summaries
type MeetingNewParamsAIConfig struct {
	// Summary Config
	Summarization param.Field[MeetingNewParamsAIConfigSummarization] `json:"summarization"`
	// Transcription Configurations
	Transcription param.Field[MeetingNewParamsAIConfigTranscription] `json:"transcription"`
}

func (r MeetingNewParamsAIConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Summary Config
type MeetingNewParamsAIConfigSummarization struct {
	// Defines the style of the summary, such as general, team meeting, or sales call.
	SummaryType param.Field[MeetingNewParamsAIConfigSummarizationSummaryType] `json:"summary_type"`
	// Determines the text format of the summary, such as plain text or markdown.
	TextFormat param.Field[MeetingNewParamsAIConfigSummarizationTextFormat] `json:"text_format"`
	// Sets the maximum number of words in the meeting summary.
	WordLimit param.Field[int64] `json:"word_limit"`
}

func (r MeetingNewParamsAIConfigSummarization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines the style of the summary, such as general, team meeting, or sales call.
type MeetingNewParamsAIConfigSummarizationSummaryType string

const (
	MeetingNewParamsAIConfigSummarizationSummaryTypeGeneral         MeetingNewParamsAIConfigSummarizationSummaryType = "general"
	MeetingNewParamsAIConfigSummarizationSummaryTypeTeamMeeting     MeetingNewParamsAIConfigSummarizationSummaryType = "team_meeting"
	MeetingNewParamsAIConfigSummarizationSummaryTypeSalesCall       MeetingNewParamsAIConfigSummarizationSummaryType = "sales_call"
	MeetingNewParamsAIConfigSummarizationSummaryTypeClientCheckIn   MeetingNewParamsAIConfigSummarizationSummaryType = "client_check_in"
	MeetingNewParamsAIConfigSummarizationSummaryTypeInterview       MeetingNewParamsAIConfigSummarizationSummaryType = "interview"
	MeetingNewParamsAIConfigSummarizationSummaryTypeDailyStandup    MeetingNewParamsAIConfigSummarizationSummaryType = "daily_standup"
	MeetingNewParamsAIConfigSummarizationSummaryTypeOneOnOneMeeting MeetingNewParamsAIConfigSummarizationSummaryType = "one_on_one_meeting"
	MeetingNewParamsAIConfigSummarizationSummaryTypeLecture         MeetingNewParamsAIConfigSummarizationSummaryType = "lecture"
	MeetingNewParamsAIConfigSummarizationSummaryTypeCodeReview      MeetingNewParamsAIConfigSummarizationSummaryType = "code_review"
)

func (r MeetingNewParamsAIConfigSummarizationSummaryType) IsKnown() bool {
	switch r {
	case MeetingNewParamsAIConfigSummarizationSummaryTypeGeneral, MeetingNewParamsAIConfigSummarizationSummaryTypeTeamMeeting, MeetingNewParamsAIConfigSummarizationSummaryTypeSalesCall, MeetingNewParamsAIConfigSummarizationSummaryTypeClientCheckIn, MeetingNewParamsAIConfigSummarizationSummaryTypeInterview, MeetingNewParamsAIConfigSummarizationSummaryTypeDailyStandup, MeetingNewParamsAIConfigSummarizationSummaryTypeOneOnOneMeeting, MeetingNewParamsAIConfigSummarizationSummaryTypeLecture, MeetingNewParamsAIConfigSummarizationSummaryTypeCodeReview:
		return true
	}
	return false
}

// Determines the text format of the summary, such as plain text or markdown.
type MeetingNewParamsAIConfigSummarizationTextFormat string

const (
	MeetingNewParamsAIConfigSummarizationTextFormatPlainText MeetingNewParamsAIConfigSummarizationTextFormat = "plain_text"
	MeetingNewParamsAIConfigSummarizationTextFormatMarkdown  MeetingNewParamsAIConfigSummarizationTextFormat = "markdown"
)

func (r MeetingNewParamsAIConfigSummarizationTextFormat) IsKnown() bool {
	switch r {
	case MeetingNewParamsAIConfigSummarizationTextFormatPlainText, MeetingNewParamsAIConfigSummarizationTextFormatMarkdown:
		return true
	}
	return false
}

// Transcription Configurations
type MeetingNewParamsAIConfigTranscription struct {
	// Adds specific terms to improve accurate detection during transcription.
	Keywords param.Field[[]string] `json:"keywords"`
	// Specifies the language code for transcription to ensure accurate results.
	Language param.Field[MeetingNewParamsAIConfigTranscriptionLanguage] `json:"language"`
	// Control the inclusion of offensive language in transcriptions.
	ProfanityFilter param.Field[bool] `json:"profanity_filter"`
}

func (r MeetingNewParamsAIConfigTranscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the language code for transcription to ensure accurate results.
type MeetingNewParamsAIConfigTranscriptionLanguage string

const (
	MeetingNewParamsAIConfigTranscriptionLanguageEnUs MeetingNewParamsAIConfigTranscriptionLanguage = "en-US"
	MeetingNewParamsAIConfigTranscriptionLanguageEnIn MeetingNewParamsAIConfigTranscriptionLanguage = "en-IN"
	MeetingNewParamsAIConfigTranscriptionLanguageDe   MeetingNewParamsAIConfigTranscriptionLanguage = "de"
	MeetingNewParamsAIConfigTranscriptionLanguageHi   MeetingNewParamsAIConfigTranscriptionLanguage = "hi"
	MeetingNewParamsAIConfigTranscriptionLanguageSv   MeetingNewParamsAIConfigTranscriptionLanguage = "sv"
	MeetingNewParamsAIConfigTranscriptionLanguageRu   MeetingNewParamsAIConfigTranscriptionLanguage = "ru"
	MeetingNewParamsAIConfigTranscriptionLanguagePl   MeetingNewParamsAIConfigTranscriptionLanguage = "pl"
	MeetingNewParamsAIConfigTranscriptionLanguageEl   MeetingNewParamsAIConfigTranscriptionLanguage = "el"
	MeetingNewParamsAIConfigTranscriptionLanguageFr   MeetingNewParamsAIConfigTranscriptionLanguage = "fr"
	MeetingNewParamsAIConfigTranscriptionLanguageNl   MeetingNewParamsAIConfigTranscriptionLanguage = "nl"
)

func (r MeetingNewParamsAIConfigTranscriptionLanguage) IsKnown() bool {
	switch r {
	case MeetingNewParamsAIConfigTranscriptionLanguageEnUs, MeetingNewParamsAIConfigTranscriptionLanguageEnIn, MeetingNewParamsAIConfigTranscriptionLanguageDe, MeetingNewParamsAIConfigTranscriptionLanguageHi, MeetingNewParamsAIConfigTranscriptionLanguageSv, MeetingNewParamsAIConfigTranscriptionLanguageRu, MeetingNewParamsAIConfigTranscriptionLanguagePl, MeetingNewParamsAIConfigTranscriptionLanguageEl, MeetingNewParamsAIConfigTranscriptionLanguageFr, MeetingNewParamsAIConfigTranscriptionLanguageNl:
		return true
	}
	return false
}

// Recording Configurations to be used for this meeting. This level of configs
// takes higher preference over App level configs on the RealtimeKit developer
// portal.
type MeetingNewParamsRecordingConfig struct {
	// Object containing configuration regarding the audio that is being recorded.
	AudioConfig param.Field[MeetingNewParamsRecordingConfigAudioConfig] `json:"audio_config"`
	// Adds a prefix to the beginning of the file name of the recording.
	FileNamePrefix      param.Field[string]                                             `json:"file_name_prefix"`
	LiveStreamingConfig param.Field[MeetingNewParamsRecordingConfigLiveStreamingConfig] `json:"live_streaming_config"`
	// Specifies the maximum duration for recording in seconds, ranging from a minimum
	// of 60 seconds to a maximum of 24 hours.
	MaxSeconds              param.Field[float64]                                                `json:"max_seconds"`
	RealtimekitBucketConfig param.Field[MeetingNewParamsRecordingConfigRealtimekitBucketConfig] `json:"realtimekit_bucket_config"`
	StorageConfig           param.Field[MeetingNewParamsRecordingConfigStorageConfig]           `json:"storage_config"`
	VideoConfig             param.Field[MeetingNewParamsRecordingConfigVideoConfig]             `json:"video_config"`
}

func (r MeetingNewParamsRecordingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object containing configuration regarding the audio that is being recorded.
type MeetingNewParamsRecordingConfigAudioConfig struct {
	// Audio signal pathway within an audio file that carries a specific sound source.
	Channel param.Field[MeetingNewParamsRecordingConfigAudioConfigChannel] `json:"channel"`
	// Codec using which the recording will be encoded. If VP8/VP9 is selected for
	// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
	// audioConfig is automatically set to vorbis.
	Codec param.Field[MeetingNewParamsRecordingConfigAudioConfigCodec] `json:"codec"`
	// Controls whether to export audio file seperately
	ExportFile param.Field[bool] `json:"export_file"`
}

func (r MeetingNewParamsRecordingConfigAudioConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Audio signal pathway within an audio file that carries a specific sound source.
type MeetingNewParamsRecordingConfigAudioConfigChannel string

const (
	MeetingNewParamsRecordingConfigAudioConfigChannelMono   MeetingNewParamsRecordingConfigAudioConfigChannel = "mono"
	MeetingNewParamsRecordingConfigAudioConfigChannelStereo MeetingNewParamsRecordingConfigAudioConfigChannel = "stereo"
)

func (r MeetingNewParamsRecordingConfigAudioConfigChannel) IsKnown() bool {
	switch r {
	case MeetingNewParamsRecordingConfigAudioConfigChannelMono, MeetingNewParamsRecordingConfigAudioConfigChannelStereo:
		return true
	}
	return false
}

// Codec using which the recording will be encoded. If VP8/VP9 is selected for
// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
// audioConfig is automatically set to vorbis.
type MeetingNewParamsRecordingConfigAudioConfigCodec string

const (
	MeetingNewParamsRecordingConfigAudioConfigCodecMP3 MeetingNewParamsRecordingConfigAudioConfigCodec = "MP3"
	MeetingNewParamsRecordingConfigAudioConfigCodecAac MeetingNewParamsRecordingConfigAudioConfigCodec = "AAC"
)

func (r MeetingNewParamsRecordingConfigAudioConfigCodec) IsKnown() bool {
	switch r {
	case MeetingNewParamsRecordingConfigAudioConfigCodecMP3, MeetingNewParamsRecordingConfigAudioConfigCodecAac:
		return true
	}
	return false
}

type MeetingNewParamsRecordingConfigLiveStreamingConfig struct {
	// RTMP URL to stream to
	RtmpURL param.Field[string] `json:"rtmp_url" format:"uri"`
}

func (r MeetingNewParamsRecordingConfigLiveStreamingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MeetingNewParamsRecordingConfigRealtimekitBucketConfig struct {
	// Controls whether recordings are uploaded to RealtimeKit's bucket. If set to
	// false, `download_url`, `audio_download_url`, `download_url_expiry` won't be
	// generated for a recording.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r MeetingNewParamsRecordingConfigRealtimekitBucketConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MeetingNewParamsRecordingConfigStorageConfig struct {
	// Type of storage media.
	Type param.Field[MeetingNewParamsRecordingConfigStorageConfigType] `json:"type,required"`
	// Access key of the storage medium. Access key is not required for the `gcs`
	// storage media type.
	//
	// Note that this field is not readable by clients, only writeable.
	AccessKey param.Field[string] `json:"access_key"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod param.Field[MeetingNewParamsRecordingConfigStorageConfigAuthMethod] `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket param.Field[string] `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host param.Field[string] `json:"host"`
	// SSH destination server password for SFTP type storage medium when auth_method is
	// "PASSWORD". If auth_method is "KEY", this specifies the password for the ssh
	// private key.
	Password param.Field[string] `json:"password"`
	// Path relative to the bucket root at which the recording will be placed.
	Path param.Field[string] `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port param.Field[float64] `json:"port"`
	// Private key used to login to destination SSH server for SFTP type storage
	// medium, when auth_method used is "KEY"
	PrivateKey param.Field[string] `json:"private_key"`
	// Region of the storage medium.
	Region param.Field[string] `json:"region"`
	// Secret key of the storage medium. Similar to `access_key`, it is only writeable
	// by clients, not readable.
	Secret param.Field[string] `json:"secret"`
	// SSH destination server username for SFTP type storage medium
	Username param.Field[string] `json:"username"`
}

func (r MeetingNewParamsRecordingConfigStorageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of storage media.
type MeetingNewParamsRecordingConfigStorageConfigType string

const (
	MeetingNewParamsRecordingConfigStorageConfigTypeAws          MeetingNewParamsRecordingConfigStorageConfigType = "aws"
	MeetingNewParamsRecordingConfigStorageConfigTypeAzure        MeetingNewParamsRecordingConfigStorageConfigType = "azure"
	MeetingNewParamsRecordingConfigStorageConfigTypeDigitalocean MeetingNewParamsRecordingConfigStorageConfigType = "digitalocean"
	MeetingNewParamsRecordingConfigStorageConfigTypeGcs          MeetingNewParamsRecordingConfigStorageConfigType = "gcs"
	MeetingNewParamsRecordingConfigStorageConfigTypeSftp         MeetingNewParamsRecordingConfigStorageConfigType = "sftp"
)

func (r MeetingNewParamsRecordingConfigStorageConfigType) IsKnown() bool {
	switch r {
	case MeetingNewParamsRecordingConfigStorageConfigTypeAws, MeetingNewParamsRecordingConfigStorageConfigTypeAzure, MeetingNewParamsRecordingConfigStorageConfigTypeDigitalocean, MeetingNewParamsRecordingConfigStorageConfigTypeGcs, MeetingNewParamsRecordingConfigStorageConfigTypeSftp:
		return true
	}
	return false
}

// Authentication method used for "sftp" type storage medium
type MeetingNewParamsRecordingConfigStorageConfigAuthMethod string

const (
	MeetingNewParamsRecordingConfigStorageConfigAuthMethodKey      MeetingNewParamsRecordingConfigStorageConfigAuthMethod = "KEY"
	MeetingNewParamsRecordingConfigStorageConfigAuthMethodPassword MeetingNewParamsRecordingConfigStorageConfigAuthMethod = "PASSWORD"
)

func (r MeetingNewParamsRecordingConfigStorageConfigAuthMethod) IsKnown() bool {
	switch r {
	case MeetingNewParamsRecordingConfigStorageConfigAuthMethodKey, MeetingNewParamsRecordingConfigStorageConfigAuthMethodPassword:
		return true
	}
	return false
}

type MeetingNewParamsRecordingConfigVideoConfig struct {
	// Codec using which the recording will be encoded.
	Codec param.Field[MeetingNewParamsRecordingConfigVideoConfigCodec] `json:"codec"`
	// Controls whether to export video file seperately
	ExportFile param.Field[bool] `json:"export_file"`
	// Height of the recording video in pixels
	Height param.Field[int64] `json:"height"`
	// Watermark to be added to the recording
	Watermark param.Field[MeetingNewParamsRecordingConfigVideoConfigWatermark] `json:"watermark"`
	// Width of the recording video in pixels
	Width param.Field[int64] `json:"width"`
}

func (r MeetingNewParamsRecordingConfigVideoConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Codec using which the recording will be encoded.
type MeetingNewParamsRecordingConfigVideoConfigCodec string

const (
	MeetingNewParamsRecordingConfigVideoConfigCodecH264 MeetingNewParamsRecordingConfigVideoConfigCodec = "H264"
	MeetingNewParamsRecordingConfigVideoConfigCodecVp8  MeetingNewParamsRecordingConfigVideoConfigCodec = "VP8"
)

func (r MeetingNewParamsRecordingConfigVideoConfigCodec) IsKnown() bool {
	switch r {
	case MeetingNewParamsRecordingConfigVideoConfigCodecH264, MeetingNewParamsRecordingConfigVideoConfigCodecVp8:
		return true
	}
	return false
}

// Watermark to be added to the recording
type MeetingNewParamsRecordingConfigVideoConfigWatermark struct {
	// Position of the watermark
	Position param.Field[MeetingNewParamsRecordingConfigVideoConfigWatermarkPosition] `json:"position"`
	// Size of the watermark
	Size param.Field[MeetingNewParamsRecordingConfigVideoConfigWatermarkSize] `json:"size"`
	// URL of the watermark image
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r MeetingNewParamsRecordingConfigVideoConfigWatermark) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Position of the watermark
type MeetingNewParamsRecordingConfigVideoConfigWatermarkPosition string

const (
	MeetingNewParamsRecordingConfigVideoConfigWatermarkPositionLeftTop     MeetingNewParamsRecordingConfigVideoConfigWatermarkPosition = "left top"
	MeetingNewParamsRecordingConfigVideoConfigWatermarkPositionRightTop    MeetingNewParamsRecordingConfigVideoConfigWatermarkPosition = "right top"
	MeetingNewParamsRecordingConfigVideoConfigWatermarkPositionLeftBottom  MeetingNewParamsRecordingConfigVideoConfigWatermarkPosition = "left bottom"
	MeetingNewParamsRecordingConfigVideoConfigWatermarkPositionRightBottom MeetingNewParamsRecordingConfigVideoConfigWatermarkPosition = "right bottom"
)

func (r MeetingNewParamsRecordingConfigVideoConfigWatermarkPosition) IsKnown() bool {
	switch r {
	case MeetingNewParamsRecordingConfigVideoConfigWatermarkPositionLeftTop, MeetingNewParamsRecordingConfigVideoConfigWatermarkPositionRightTop, MeetingNewParamsRecordingConfigVideoConfigWatermarkPositionLeftBottom, MeetingNewParamsRecordingConfigVideoConfigWatermarkPositionRightBottom:
		return true
	}
	return false
}

// Size of the watermark
type MeetingNewParamsRecordingConfigVideoConfigWatermarkSize struct {
	// Height of the watermark in px
	Height param.Field[int64] `json:"height"`
	// Width of the watermark in px
	Width param.Field[int64] `json:"width"`
}

func (r MeetingNewParamsRecordingConfigVideoConfigWatermarkSize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MeetingAddParticipantParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A unique participant ID. You must specify a unique ID for the participant, for
	// example, UUID, email address, and so on.
	CustomParticipantID param.Field[string] `json:"custom_participant_id,required"`
	// Name of the preset to apply to this participant.
	PresetName param.Field[string] `json:"preset_name,required"`
	// (Optional) Name of the participant.
	Name param.Field[string] `json:"name"`
	// (Optional) A URL to a picture to be used for the participant.
	Picture param.Field[string] `json:"picture" format:"uri"`
}

func (r MeetingAddParticipantParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MeetingDeleteMeetingParticipantParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type MeetingEditParticipantParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// (Optional) Name of the participant.
	Name param.Field[string] `json:"name"`
	// (Optional) A URL to a picture to be used for the participant.
	Picture param.Field[string] `json:"picture" format:"uri"`
	// (Optional) Name of the preset to apply to this participant.
	PresetName param.Field[string] `json:"preset_name"`
}

func (r MeetingEditParticipantParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MeetingGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The end time range for which you want to retrieve the meetings. The time must be
	// specified in ISO format.
	EndTime param.Field[time.Time] `query:"end_time" format:"date-time"`
	// The page number from which you want your page search results to be displayed.
	PageNo param.Field[float64] `query:"page_no"`
	// Number of results per page
	PerPage param.Field[float64] `query:"per_page"`
	// The search query string. You can search using the meeting ID or title.
	Search param.Field[string] `query:"search"`
	// The start time range for which you want to retrieve the meetings. The time must
	// be specified in ISO format.
	StartTime param.Field[time.Time] `query:"start_time" format:"date-time"`
}

// URLQuery serializes [MeetingGetParams]'s query parameters as `url.Values`.
func (r MeetingGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type MeetingGetMeetingByIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	Name      param.Field[string] `query:"name"`
}

// URLQuery serializes [MeetingGetMeetingByIDParams]'s query parameters as
// `url.Values`.
func (r MeetingGetMeetingByIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type MeetingGetMeetingParticipantParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type MeetingGetMeetingParticipantsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The page number from which you want your page search results to be displayed.
	PageNo param.Field[float64] `query:"page_no"`
	// Number of results per page
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [MeetingGetMeetingParticipantsParams]'s query parameters as
// `url.Values`.
func (r MeetingGetMeetingParticipantsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type MeetingRefreshParticipantTokenParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type MeetingReplaceMeetingByIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The AI Config allows you to customize the behavior of meeting transcriptions and
	// summaries
	AIConfig param.Field[MeetingReplaceMeetingByIDParamsAIConfig] `json:"ai_config"`
	// Specifies if the meeting should start getting livestreamed on start.
	LiveStreamOnStart param.Field[bool] `json:"live_stream_on_start"`
	// If a meeting is set to persist_chat, meeting chat would remain for a week within
	// the meeting space.
	PersistChat param.Field[bool] `json:"persist_chat"`
	// Specifies if the meeting should start getting recorded as soon as someone joins
	// the meeting.
	RecordOnStart param.Field[bool] `json:"record_on_start"`
	// Recording Configurations to be used for this meeting. This level of configs
	// takes higher preference over App level configs on the RealtimeKit developer
	// portal.
	RecordingConfig param.Field[MeetingReplaceMeetingByIDParamsRecordingConfig] `json:"recording_config"`
	// Time in seconds, for which a session remains active, after the last participant
	// has left the meeting.
	SessionKeepAliveTimeInSecs param.Field[float64] `json:"session_keep_alive_time_in_secs"`
	// Automatically generate summary of meetings using transcripts. Requires
	// Transcriptions to be enabled, and can be retrieved via Webhooks or summary API.
	SummarizeOnEnd param.Field[bool] `json:"summarize_on_end"`
	// Title of the meeting
	Title param.Field[string] `json:"title"`
}

func (r MeetingReplaceMeetingByIDParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The AI Config allows you to customize the behavior of meeting transcriptions and
// summaries
type MeetingReplaceMeetingByIDParamsAIConfig struct {
	// Summary Config
	Summarization param.Field[MeetingReplaceMeetingByIDParamsAIConfigSummarization] `json:"summarization"`
	// Transcription Configurations
	Transcription param.Field[MeetingReplaceMeetingByIDParamsAIConfigTranscription] `json:"transcription"`
}

func (r MeetingReplaceMeetingByIDParamsAIConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Summary Config
type MeetingReplaceMeetingByIDParamsAIConfigSummarization struct {
	// Defines the style of the summary, such as general, team meeting, or sales call.
	SummaryType param.Field[MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryType] `json:"summary_type"`
	// Determines the text format of the summary, such as plain text or markdown.
	TextFormat param.Field[MeetingReplaceMeetingByIDParamsAIConfigSummarizationTextFormat] `json:"text_format"`
	// Sets the maximum number of words in the meeting summary.
	WordLimit param.Field[int64] `json:"word_limit"`
}

func (r MeetingReplaceMeetingByIDParamsAIConfigSummarization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines the style of the summary, such as general, team meeting, or sales call.
type MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryType string

const (
	MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeGeneral         MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryType = "general"
	MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeTeamMeeting     MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryType = "team_meeting"
	MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeSalesCall       MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryType = "sales_call"
	MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeClientCheckIn   MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryType = "client_check_in"
	MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeInterview       MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryType = "interview"
	MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeDailyStandup    MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryType = "daily_standup"
	MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeOneOnOneMeeting MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryType = "one_on_one_meeting"
	MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeLecture         MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryType = "lecture"
	MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeCodeReview      MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryType = "code_review"
)

func (r MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryType) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeGeneral, MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeTeamMeeting, MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeSalesCall, MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeClientCheckIn, MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeInterview, MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeDailyStandup, MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeOneOnOneMeeting, MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeLecture, MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeCodeReview:
		return true
	}
	return false
}

// Determines the text format of the summary, such as plain text or markdown.
type MeetingReplaceMeetingByIDParamsAIConfigSummarizationTextFormat string

const (
	MeetingReplaceMeetingByIDParamsAIConfigSummarizationTextFormatPlainText MeetingReplaceMeetingByIDParamsAIConfigSummarizationTextFormat = "plain_text"
	MeetingReplaceMeetingByIDParamsAIConfigSummarizationTextFormatMarkdown  MeetingReplaceMeetingByIDParamsAIConfigSummarizationTextFormat = "markdown"
)

func (r MeetingReplaceMeetingByIDParamsAIConfigSummarizationTextFormat) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDParamsAIConfigSummarizationTextFormatPlainText, MeetingReplaceMeetingByIDParamsAIConfigSummarizationTextFormatMarkdown:
		return true
	}
	return false
}

// Transcription Configurations
type MeetingReplaceMeetingByIDParamsAIConfigTranscription struct {
	// Adds specific terms to improve accurate detection during transcription.
	Keywords param.Field[[]string] `json:"keywords"`
	// Specifies the language code for transcription to ensure accurate results.
	Language param.Field[MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguage] `json:"language"`
	// Control the inclusion of offensive language in transcriptions.
	ProfanityFilter param.Field[bool] `json:"profanity_filter"`
}

func (r MeetingReplaceMeetingByIDParamsAIConfigTranscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the language code for transcription to ensure accurate results.
type MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguage string

const (
	MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageEnUs MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguage = "en-US"
	MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageEnIn MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguage = "en-IN"
	MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageDe   MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguage = "de"
	MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageHi   MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguage = "hi"
	MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageSv   MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguage = "sv"
	MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageRu   MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguage = "ru"
	MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguagePl   MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguage = "pl"
	MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageEl   MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguage = "el"
	MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageFr   MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguage = "fr"
	MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageNl   MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguage = "nl"
)

func (r MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguage) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageEnUs, MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageEnIn, MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageDe, MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageHi, MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageSv, MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageRu, MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguagePl, MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageEl, MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageFr, MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageNl:
		return true
	}
	return false
}

// Recording Configurations to be used for this meeting. This level of configs
// takes higher preference over App level configs on the RealtimeKit developer
// portal.
type MeetingReplaceMeetingByIDParamsRecordingConfig struct {
	// Object containing configuration regarding the audio that is being recorded.
	AudioConfig param.Field[MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfig] `json:"audio_config"`
	// Adds a prefix to the beginning of the file name of the recording.
	FileNamePrefix      param.Field[string]                                                            `json:"file_name_prefix"`
	LiveStreamingConfig param.Field[MeetingReplaceMeetingByIDParamsRecordingConfigLiveStreamingConfig] `json:"live_streaming_config"`
	// Specifies the maximum duration for recording in seconds, ranging from a minimum
	// of 60 seconds to a maximum of 24 hours.
	MaxSeconds              param.Field[float64]                                                               `json:"max_seconds"`
	RealtimekitBucketConfig param.Field[MeetingReplaceMeetingByIDParamsRecordingConfigRealtimekitBucketConfig] `json:"realtimekit_bucket_config"`
	StorageConfig           param.Field[MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfig]           `json:"storage_config"`
	VideoConfig             param.Field[MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfig]             `json:"video_config"`
}

func (r MeetingReplaceMeetingByIDParamsRecordingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object containing configuration regarding the audio that is being recorded.
type MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfig struct {
	// Audio signal pathway within an audio file that carries a specific sound source.
	Channel param.Field[MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigChannel] `json:"channel"`
	// Codec using which the recording will be encoded. If VP8/VP9 is selected for
	// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
	// audioConfig is automatically set to vorbis.
	Codec param.Field[MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigCodec] `json:"codec"`
	// Controls whether to export audio file seperately
	ExportFile param.Field[bool] `json:"export_file"`
}

func (r MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Audio signal pathway within an audio file that carries a specific sound source.
type MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigChannel string

const (
	MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigChannelMono   MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigChannel = "mono"
	MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigChannelStereo MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigChannel = "stereo"
)

func (r MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigChannel) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigChannelMono, MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigChannelStereo:
		return true
	}
	return false
}

// Codec using which the recording will be encoded. If VP8/VP9 is selected for
// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
// audioConfig is automatically set to vorbis.
type MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigCodec string

const (
	MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigCodecMP3 MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigCodec = "MP3"
	MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigCodecAac MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigCodec = "AAC"
)

func (r MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigCodec) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigCodecMP3, MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigCodecAac:
		return true
	}
	return false
}

type MeetingReplaceMeetingByIDParamsRecordingConfigLiveStreamingConfig struct {
	// RTMP URL to stream to
	RtmpURL param.Field[string] `json:"rtmp_url" format:"uri"`
}

func (r MeetingReplaceMeetingByIDParamsRecordingConfigLiveStreamingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MeetingReplaceMeetingByIDParamsRecordingConfigRealtimekitBucketConfig struct {
	// Controls whether recordings are uploaded to RealtimeKit's bucket. If set to
	// false, `download_url`, `audio_download_url`, `download_url_expiry` won't be
	// generated for a recording.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r MeetingReplaceMeetingByIDParamsRecordingConfigRealtimekitBucketConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfig struct {
	// Type of storage media.
	Type param.Field[MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigType] `json:"type,required"`
	// Access key of the storage medium. Access key is not required for the `gcs`
	// storage media type.
	//
	// Note that this field is not readable by clients, only writeable.
	AccessKey param.Field[string] `json:"access_key"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod param.Field[MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigAuthMethod] `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket param.Field[string] `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host param.Field[string] `json:"host"`
	// SSH destination server password for SFTP type storage medium when auth_method is
	// "PASSWORD". If auth_method is "KEY", this specifies the password for the ssh
	// private key.
	Password param.Field[string] `json:"password"`
	// Path relative to the bucket root at which the recording will be placed.
	Path param.Field[string] `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port param.Field[float64] `json:"port"`
	// Private key used to login to destination SSH server for SFTP type storage
	// medium, when auth_method used is "KEY"
	PrivateKey param.Field[string] `json:"private_key"`
	// Region of the storage medium.
	Region param.Field[string] `json:"region"`
	// Secret key of the storage medium. Similar to `access_key`, it is only writeable
	// by clients, not readable.
	Secret param.Field[string] `json:"secret"`
	// SSH destination server username for SFTP type storage medium
	Username param.Field[string] `json:"username"`
}

func (r MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of storage media.
type MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigType string

const (
	MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigTypeAws          MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigType = "aws"
	MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigTypeAzure        MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigType = "azure"
	MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigTypeDigitalocean MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigType = "digitalocean"
	MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigTypeGcs          MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigType = "gcs"
	MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigTypeSftp         MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigType = "sftp"
)

func (r MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigType) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigTypeAws, MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigTypeAzure, MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigTypeDigitalocean, MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigTypeGcs, MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigTypeSftp:
		return true
	}
	return false
}

// Authentication method used for "sftp" type storage medium
type MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigAuthMethod string

const (
	MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigAuthMethodKey      MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigAuthMethod = "KEY"
	MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigAuthMethodPassword MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigAuthMethod = "PASSWORD"
)

func (r MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigAuthMethod) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigAuthMethodKey, MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigAuthMethodPassword:
		return true
	}
	return false
}

type MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfig struct {
	// Codec using which the recording will be encoded.
	Codec param.Field[MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigCodec] `json:"codec"`
	// Controls whether to export video file seperately
	ExportFile param.Field[bool] `json:"export_file"`
	// Height of the recording video in pixels
	Height param.Field[int64] `json:"height"`
	// Watermark to be added to the recording
	Watermark param.Field[MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermark] `json:"watermark"`
	// Width of the recording video in pixels
	Width param.Field[int64] `json:"width"`
}

func (r MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Codec using which the recording will be encoded.
type MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigCodec string

const (
	MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigCodecH264 MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigCodec = "H264"
	MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigCodecVp8  MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigCodec = "VP8"
)

func (r MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigCodec) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigCodecH264, MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigCodecVp8:
		return true
	}
	return false
}

// Watermark to be added to the recording
type MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermark struct {
	// Position of the watermark
	Position param.Field[MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPosition] `json:"position"`
	// Size of the watermark
	Size param.Field[MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkSize] `json:"size"`
	// URL of the watermark image
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermark) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Position of the watermark
type MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPosition string

const (
	MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPositionLeftTop     MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPosition = "left top"
	MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPositionRightTop    MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPosition = "right top"
	MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPositionLeftBottom  MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPosition = "left bottom"
	MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPositionRightBottom MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPosition = "right bottom"
)

func (r MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPosition) IsKnown() bool {
	switch r {
	case MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPositionLeftTop, MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPositionRightTop, MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPositionLeftBottom, MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPositionRightBottom:
		return true
	}
	return false
}

// Size of the watermark
type MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkSize struct {
	// Height of the watermark in px
	Height param.Field[int64] `json:"height"`
	// Width of the watermark in px
	Width param.Field[int64] `json:"width"`
}

func (r MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkSize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MeetingUpdateMeetingByIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The AI Config allows you to customize the behavior of meeting transcriptions and
	// summaries
	AIConfig param.Field[MeetingUpdateMeetingByIDParamsAIConfig] `json:"ai_config"`
	// Specifies if the meeting should start getting livestreamed on start.
	LiveStreamOnStart param.Field[bool] `json:"live_stream_on_start"`
	// If a meeting is updated to persist_chat, meeting chat would remain for a week
	// within the meeting space.
	PersistChat param.Field[bool] `json:"persist_chat"`
	// Specifies if the meeting should start getting recorded as soon as someone joins
	// the meeting.
	RecordOnStart param.Field[bool] `json:"record_on_start"`
	// Time in seconds, for which a session remains active, after the last participant
	// has left the meeting.
	SessionKeepAliveTimeInSecs param.Field[float64] `json:"session_keep_alive_time_in_secs"`
	// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
	// `INACTIVE` meeting.
	Status param.Field[MeetingUpdateMeetingByIDParamsStatus] `json:"status"`
	// Automatically generate summary of meetings using transcripts. Requires
	// Transcriptions to be enabled, and can be retrieved via Webhooks or summary API.
	SummarizeOnEnd param.Field[bool] `json:"summarize_on_end"`
	// Title of the meeting
	Title param.Field[string] `json:"title"`
}

func (r MeetingUpdateMeetingByIDParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The AI Config allows you to customize the behavior of meeting transcriptions and
// summaries
type MeetingUpdateMeetingByIDParamsAIConfig struct {
	// Summary Config
	Summarization param.Field[MeetingUpdateMeetingByIDParamsAIConfigSummarization] `json:"summarization"`
	// Transcription Configurations
	Transcription param.Field[MeetingUpdateMeetingByIDParamsAIConfigTranscription] `json:"transcription"`
}

func (r MeetingUpdateMeetingByIDParamsAIConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Summary Config
type MeetingUpdateMeetingByIDParamsAIConfigSummarization struct {
	// Defines the style of the summary, such as general, team meeting, or sales call.
	SummaryType param.Field[MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryType] `json:"summary_type"`
	// Determines the text format of the summary, such as plain text or markdown.
	TextFormat param.Field[MeetingUpdateMeetingByIDParamsAIConfigSummarizationTextFormat] `json:"text_format"`
	// Sets the maximum number of words in the meeting summary.
	WordLimit param.Field[int64] `json:"word_limit"`
}

func (r MeetingUpdateMeetingByIDParamsAIConfigSummarization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines the style of the summary, such as general, team meeting, or sales call.
type MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryType string

const (
	MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeGeneral         MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryType = "general"
	MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeTeamMeeting     MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryType = "team_meeting"
	MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeSalesCall       MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryType = "sales_call"
	MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeClientCheckIn   MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryType = "client_check_in"
	MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeInterview       MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryType = "interview"
	MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeDailyStandup    MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryType = "daily_standup"
	MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeOneOnOneMeeting MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryType = "one_on_one_meeting"
	MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeLecture         MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryType = "lecture"
	MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeCodeReview      MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryType = "code_review"
)

func (r MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryType) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeGeneral, MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeTeamMeeting, MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeSalesCall, MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeClientCheckIn, MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeInterview, MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeDailyStandup, MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeOneOnOneMeeting, MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeLecture, MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeCodeReview:
		return true
	}
	return false
}

// Determines the text format of the summary, such as plain text or markdown.
type MeetingUpdateMeetingByIDParamsAIConfigSummarizationTextFormat string

const (
	MeetingUpdateMeetingByIDParamsAIConfigSummarizationTextFormatPlainText MeetingUpdateMeetingByIDParamsAIConfigSummarizationTextFormat = "plain_text"
	MeetingUpdateMeetingByIDParamsAIConfigSummarizationTextFormatMarkdown  MeetingUpdateMeetingByIDParamsAIConfigSummarizationTextFormat = "markdown"
)

func (r MeetingUpdateMeetingByIDParamsAIConfigSummarizationTextFormat) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDParamsAIConfigSummarizationTextFormatPlainText, MeetingUpdateMeetingByIDParamsAIConfigSummarizationTextFormatMarkdown:
		return true
	}
	return false
}

// Transcription Configurations
type MeetingUpdateMeetingByIDParamsAIConfigTranscription struct {
	// Adds specific terms to improve accurate detection during transcription.
	Keywords param.Field[[]string] `json:"keywords"`
	// Specifies the language code for transcription to ensure accurate results.
	Language param.Field[MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguage] `json:"language"`
	// Control the inclusion of offensive language in transcriptions.
	ProfanityFilter param.Field[bool] `json:"profanity_filter"`
}

func (r MeetingUpdateMeetingByIDParamsAIConfigTranscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the language code for transcription to ensure accurate results.
type MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguage string

const (
	MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageEnUs MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguage = "en-US"
	MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageEnIn MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguage = "en-IN"
	MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageDe   MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguage = "de"
	MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageHi   MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguage = "hi"
	MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageSv   MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguage = "sv"
	MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageRu   MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguage = "ru"
	MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguagePl   MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguage = "pl"
	MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageEl   MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguage = "el"
	MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageFr   MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguage = "fr"
	MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageNl   MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguage = "nl"
)

func (r MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguage) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageEnUs, MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageEnIn, MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageDe, MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageHi, MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageSv, MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageRu, MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguagePl, MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageEl, MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageFr, MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageNl:
		return true
	}
	return false
}

// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
// `INACTIVE` meeting.
type MeetingUpdateMeetingByIDParamsStatus string

const (
	MeetingUpdateMeetingByIDParamsStatusActive   MeetingUpdateMeetingByIDParamsStatus = "ACTIVE"
	MeetingUpdateMeetingByIDParamsStatusInactive MeetingUpdateMeetingByIDParamsStatus = "INACTIVE"
)

func (r MeetingUpdateMeetingByIDParamsStatus) IsKnown() bool {
	switch r {
	case MeetingUpdateMeetingByIDParamsStatusActive, MeetingUpdateMeetingByIDParamsStatusInactive:
		return true
	}
	return false
}
