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
	"github.com/tidwall/gjson"
)

// RecordingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecordingService] method instead.
type RecordingService struct {
	Options []option.RequestOption
}

// NewRecordingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecordingService(opts ...option.RequestOption) (r *RecordingService) {
	r = &RecordingService{}
	r.Options = opts
	return
}

// Returns the active recording details for the given meeting ID.
func (r *RecordingService) GetActiveRecordings(ctx context.Context, appID string, meetingID string, query RecordingGetActiveRecordingsParams, opts ...option.RequestOption) (res *RecordingGetActiveRecordingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/recordings/active-recording/%s", query.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns details of a recording for the given recording ID.
func (r *RecordingService) GetOneRecording(ctx context.Context, appID string, recordingID string, query RecordingGetOneRecordingParams, opts ...option.RequestOption) (res *RecordingGetOneRecordingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if recordingID == "" {
		err = errors.New("missing required recording_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/recordings/%s", query.AccountID, appID, recordingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns all recordings for an App. If the `meeting_id` parameter is passed,
// returns all recordings for the given meeting ID.
func (r *RecordingService) GetRecordings(ctx context.Context, appID string, params RecordingGetRecordingsParams, opts ...option.RequestOption) (res *RecordingGetRecordingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/recordings", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Pause/Resume/Stop a given recording ID.
func (r *RecordingService) PauseResumeStopRecording(ctx context.Context, appID string, recordingID string, params RecordingPauseResumeStopRecordingParams, opts ...option.RequestOption) (res *RecordingPauseResumeStopRecordingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if recordingID == "" {
		err = errors.New("missing required recording_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/recordings/%s", params.AccountID, appID, recordingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Starts recording a meeting. The meeting can be started by an App admin directly,
// or a participant with permissions to start a recording, based on the type of
// authorization used.
func (r *RecordingService) StartRecordings(ctx context.Context, appID string, params RecordingStartRecordingsParams, opts ...option.RequestOption) (res *RecordingStartRecordingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/recordings", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Starts track recording for a meeting. Track recording currently records separate
// participant audio tracks as WebM files in the RealtimeKit bucket. Video track
// recording is in development. For more information, refer to
// [Track recording](/realtime/realtimekit/recording-guide/track-recording/).
func (r *RecordingService) StartTrackRecording(ctx context.Context, appID string, params RecordingStartTrackRecordingParams, opts ...option.RequestOption) (res *RecordingStartTrackRecordingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/recordings/track", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type RecordingGetActiveRecordingsResponse struct {
	// Data returned by the operation
	Data RecordingGetActiveRecordingsResponseData `json:"data" api:"required"`
	// Success status of the operation
	Success bool                                     `json:"success" api:"required"`
	JSON    recordingGetActiveRecordingsResponseJSON `json:"-"`
}

// recordingGetActiveRecordingsResponseJSON contains the JSON metadata for the
// struct [RecordingGetActiveRecordingsResponse]
type recordingGetActiveRecordingsResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetActiveRecordingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetActiveRecordingsResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type RecordingGetActiveRecordingsResponseData struct {
	// ID of the recording
	ID string `json:"id" api:"required" format:"uuid"`
	// If the audio_config is passed, the URL for downloading the audio recording is
	// returned.
	AudioDownloadURL string `json:"audio_download_url" api:"required,nullable" format:"uri"`
	// URL where the recording can be downloaded.
	DownloadURL string `json:"download_url" api:"required,nullable" format:"uri"`
	// Timestamp when the download URL expires.
	DownloadURLExpiry time.Time `json:"download_url_expiry" api:"required,nullable" format:"date-time"`
	// File size of the recording, in bytes.
	FileSize float64 `json:"file_size" api:"required,nullable"`
	// Timestamp when this recording was invoked.
	InvokedTime time.Time `json:"invoked_time" api:"required" format:"date-time"`
	// File name of the recording.
	OutputFileName string `json:"output_file_name" api:"required"`
	// ID of the meeting session this recording is for.
	SessionID string `json:"session_id" api:"required,nullable" format:"uuid"`
	// Timestamp when this recording actually started after being invoked. Usually a
	// few seconds after `invoked_time`.
	StartedTime time.Time `json:"started_time" api:"required,nullable" format:"date-time"`
	// Current status of the recording.
	Status RecordingGetActiveRecordingsResponseDataStatus `json:"status" api:"required"`
	// Timestamp when this recording was stopped. Optional; is present only when the
	// recording has actually been stopped.
	StoppedTime time.Time `json:"stopped_time" api:"required,nullable" format:"date-time"`
	// Total recording time in seconds.
	RecordingDuration int64                                        `json:"recording_duration"`
	JSON              recordingGetActiveRecordingsResponseDataJSON `json:"-"`
}

// recordingGetActiveRecordingsResponseDataJSON contains the JSON metadata for the
// struct [RecordingGetActiveRecordingsResponseData]
type recordingGetActiveRecordingsResponseDataJSON struct {
	ID                apijson.Field
	AudioDownloadURL  apijson.Field
	DownloadURL       apijson.Field
	DownloadURLExpiry apijson.Field
	FileSize          apijson.Field
	InvokedTime       apijson.Field
	OutputFileName    apijson.Field
	SessionID         apijson.Field
	StartedTime       apijson.Field
	Status            apijson.Field
	StoppedTime       apijson.Field
	RecordingDuration apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordingGetActiveRecordingsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetActiveRecordingsResponseDataJSON) RawJSON() string {
	return r.raw
}

// Current status of the recording.
type RecordingGetActiveRecordingsResponseDataStatus string

const (
	RecordingGetActiveRecordingsResponseDataStatusInvoked   RecordingGetActiveRecordingsResponseDataStatus = "INVOKED"
	RecordingGetActiveRecordingsResponseDataStatusRecording RecordingGetActiveRecordingsResponseDataStatus = "RECORDING"
	RecordingGetActiveRecordingsResponseDataStatusUploading RecordingGetActiveRecordingsResponseDataStatus = "UPLOADING"
	RecordingGetActiveRecordingsResponseDataStatusUploaded  RecordingGetActiveRecordingsResponseDataStatus = "UPLOADED"
	RecordingGetActiveRecordingsResponseDataStatusErrored   RecordingGetActiveRecordingsResponseDataStatus = "ERRORED"
	RecordingGetActiveRecordingsResponseDataStatusPaused    RecordingGetActiveRecordingsResponseDataStatus = "PAUSED"
)

func (r RecordingGetActiveRecordingsResponseDataStatus) IsKnown() bool {
	switch r {
	case RecordingGetActiveRecordingsResponseDataStatusInvoked, RecordingGetActiveRecordingsResponseDataStatusRecording, RecordingGetActiveRecordingsResponseDataStatusUploading, RecordingGetActiveRecordingsResponseDataStatusUploaded, RecordingGetActiveRecordingsResponseDataStatusErrored, RecordingGetActiveRecordingsResponseDataStatusPaused:
		return true
	}
	return false
}

type RecordingGetOneRecordingResponse struct {
	// Success status of the operation
	Success bool `json:"success" api:"required"`
	// Data returned by the operation
	Data RecordingGetOneRecordingResponseData `json:"data"`
	JSON recordingGetOneRecordingResponseJSON `json:"-"`
}

// recordingGetOneRecordingResponseJSON contains the JSON metadata for the struct
// [RecordingGetOneRecordingResponse]
type recordingGetOneRecordingResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetOneRecordingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetOneRecordingResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type RecordingGetOneRecordingResponseData struct {
	// ID of the recording
	ID string `json:"id" api:"required" format:"uuid"`
	// If the audio_config is passed, the URL for downloading the audio recording is
	// returned.
	AudioDownloadURL string `json:"audio_download_url" api:"required,nullable" format:"uri"`
	// URL where the recording can be downloaded.
	DownloadURL string `json:"download_url" api:"required,nullable" format:"uri"`
	// Timestamp when the download URL expires.
	DownloadURLExpiry time.Time `json:"download_url_expiry" api:"required,nullable" format:"date-time"`
	// File size of the recording, in bytes.
	FileSize float64 `json:"file_size" api:"required,nullable"`
	// Timestamp when this recording was invoked.
	InvokedTime time.Time `json:"invoked_time" api:"required" format:"date-time"`
	// File name of the recording.
	OutputFileName string `json:"output_file_name" api:"required"`
	// ID of the meeting session this recording is for.
	SessionID string `json:"session_id" api:"required,nullable" format:"uuid"`
	// Timestamp when this recording actually started after being invoked. Usually a
	// few seconds after `invoked_time`.
	StartedTime time.Time `json:"started_time" api:"required,nullable" format:"date-time"`
	// Current status of the recording.
	Status RecordingGetOneRecordingResponseDataStatus `json:"status" api:"required"`
	// Timestamp when this recording was stopped. Optional; is present only when the
	// recording has actually been stopped.
	StoppedTime time.Time `json:"stopped_time" api:"required,nullable" format:"date-time"`
	// Total recording time in seconds.
	RecordingDuration int64                                             `json:"recording_duration"`
	StartReason       RecordingGetOneRecordingResponseDataStartReason   `json:"start_reason"`
	StopReason        RecordingGetOneRecordingResponseDataStopReason    `json:"stop_reason"`
	StorageConfig     RecordingGetOneRecordingResponseDataStorageConfig `json:"storage_config" api:"nullable"`
	JSON              recordingGetOneRecordingResponseDataJSON          `json:"-"`
}

// recordingGetOneRecordingResponseDataJSON contains the JSON metadata for the
// struct [RecordingGetOneRecordingResponseData]
type recordingGetOneRecordingResponseDataJSON struct {
	ID                apijson.Field
	AudioDownloadURL  apijson.Field
	DownloadURL       apijson.Field
	DownloadURLExpiry apijson.Field
	FileSize          apijson.Field
	InvokedTime       apijson.Field
	OutputFileName    apijson.Field
	SessionID         apijson.Field
	StartedTime       apijson.Field
	Status            apijson.Field
	StoppedTime       apijson.Field
	RecordingDuration apijson.Field
	StartReason       apijson.Field
	StopReason        apijson.Field
	StorageConfig     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordingGetOneRecordingResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetOneRecordingResponseDataJSON) RawJSON() string {
	return r.raw
}

// Current status of the recording.
type RecordingGetOneRecordingResponseDataStatus string

const (
	RecordingGetOneRecordingResponseDataStatusInvoked   RecordingGetOneRecordingResponseDataStatus = "INVOKED"
	RecordingGetOneRecordingResponseDataStatusRecording RecordingGetOneRecordingResponseDataStatus = "RECORDING"
	RecordingGetOneRecordingResponseDataStatusUploading RecordingGetOneRecordingResponseDataStatus = "UPLOADING"
	RecordingGetOneRecordingResponseDataStatusUploaded  RecordingGetOneRecordingResponseDataStatus = "UPLOADED"
	RecordingGetOneRecordingResponseDataStatusErrored   RecordingGetOneRecordingResponseDataStatus = "ERRORED"
	RecordingGetOneRecordingResponseDataStatusPaused    RecordingGetOneRecordingResponseDataStatus = "PAUSED"
)

func (r RecordingGetOneRecordingResponseDataStatus) IsKnown() bool {
	switch r {
	case RecordingGetOneRecordingResponseDataStatusInvoked, RecordingGetOneRecordingResponseDataStatusRecording, RecordingGetOneRecordingResponseDataStatusUploading, RecordingGetOneRecordingResponseDataStatusUploaded, RecordingGetOneRecordingResponseDataStatusErrored, RecordingGetOneRecordingResponseDataStatusPaused:
		return true
	}
	return false
}

type RecordingGetOneRecordingResponseDataStartReason struct {
	Caller RecordingGetOneRecordingResponseDataStartReasonCaller `json:"caller"`
	// Specifies if the recording was started using the "Start a Recording"API or using
	// the parameter RECORD_ON_START in the "Create a meeting" API.
	//
	// If the recording is initiated using the "RECORD_ON_START" parameter, the user
	// details will not be populated.
	Reason RecordingGetOneRecordingResponseDataStartReasonReason `json:"reason"`
	JSON   recordingGetOneRecordingResponseDataStartReasonJSON   `json:"-"`
}

// recordingGetOneRecordingResponseDataStartReasonJSON contains the JSON metadata
// for the struct [RecordingGetOneRecordingResponseDataStartReason]
type recordingGetOneRecordingResponseDataStartReasonJSON struct {
	Caller      apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetOneRecordingResponseDataStartReason) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetOneRecordingResponseDataStartReasonJSON) RawJSON() string {
	return r.raw
}

type RecordingGetOneRecordingResponseDataStartReasonCaller struct {
	// Name of the user who started the recording.
	Name string `json:"name"`
	// The type can be an App or a user. If the type is `user`, then only the `user_Id`
	// and `name` are returned.
	Type RecordingGetOneRecordingResponseDataStartReasonCallerType `json:"type"`
	// The user ID of the person who started the recording.
	UserID string                                                    `json:"user_Id" format:"uuid"`
	JSON   recordingGetOneRecordingResponseDataStartReasonCallerJSON `json:"-"`
}

// recordingGetOneRecordingResponseDataStartReasonCallerJSON contains the JSON
// metadata for the struct [RecordingGetOneRecordingResponseDataStartReasonCaller]
type recordingGetOneRecordingResponseDataStartReasonCallerJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetOneRecordingResponseDataStartReasonCaller) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetOneRecordingResponseDataStartReasonCallerJSON) RawJSON() string {
	return r.raw
}

// The type can be an App or a user. If the type is `user`, then only the `user_Id`
// and `name` are returned.
type RecordingGetOneRecordingResponseDataStartReasonCallerType string

const (
	RecordingGetOneRecordingResponseDataStartReasonCallerTypeOrganization RecordingGetOneRecordingResponseDataStartReasonCallerType = "ORGANIZATION"
	RecordingGetOneRecordingResponseDataStartReasonCallerTypeUser         RecordingGetOneRecordingResponseDataStartReasonCallerType = "USER"
)

func (r RecordingGetOneRecordingResponseDataStartReasonCallerType) IsKnown() bool {
	switch r {
	case RecordingGetOneRecordingResponseDataStartReasonCallerTypeOrganization, RecordingGetOneRecordingResponseDataStartReasonCallerTypeUser:
		return true
	}
	return false
}

// Specifies if the recording was started using the "Start a Recording"API or using
// the parameter RECORD_ON_START in the "Create a meeting" API.
//
// If the recording is initiated using the "RECORD_ON_START" parameter, the user
// details will not be populated.
type RecordingGetOneRecordingResponseDataStartReasonReason string

const (
	RecordingGetOneRecordingResponseDataStartReasonReasonAPICall       RecordingGetOneRecordingResponseDataStartReasonReason = "API_CALL"
	RecordingGetOneRecordingResponseDataStartReasonReasonRecordOnStart RecordingGetOneRecordingResponseDataStartReasonReason = "RECORD_ON_START"
)

func (r RecordingGetOneRecordingResponseDataStartReasonReason) IsKnown() bool {
	switch r {
	case RecordingGetOneRecordingResponseDataStartReasonReasonAPICall, RecordingGetOneRecordingResponseDataStartReasonReasonRecordOnStart:
		return true
	}
	return false
}

type RecordingGetOneRecordingResponseDataStopReason struct {
	Caller RecordingGetOneRecordingResponseDataStopReasonCaller `json:"caller"`
	// Specifies the reason why the recording stopped.
	Reason RecordingGetOneRecordingResponseDataStopReasonReason `json:"reason"`
	JSON   recordingGetOneRecordingResponseDataStopReasonJSON   `json:"-"`
}

// recordingGetOneRecordingResponseDataStopReasonJSON contains the JSON metadata
// for the struct [RecordingGetOneRecordingResponseDataStopReason]
type recordingGetOneRecordingResponseDataStopReasonJSON struct {
	Caller      apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetOneRecordingResponseDataStopReason) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetOneRecordingResponseDataStopReasonJSON) RawJSON() string {
	return r.raw
}

type RecordingGetOneRecordingResponseDataStopReasonCaller struct {
	// Name of the user who stopped the recording.
	Name string `json:"name"`
	// The type can be an App or a user. If the type is `user`, then only the `user_Id`
	// and `name` are returned.
	Type RecordingGetOneRecordingResponseDataStopReasonCallerType `json:"type"`
	// The user ID of the person who stopped the recording.
	UserID string                                                   `json:"user_Id" format:"uuid"`
	JSON   recordingGetOneRecordingResponseDataStopReasonCallerJSON `json:"-"`
}

// recordingGetOneRecordingResponseDataStopReasonCallerJSON contains the JSON
// metadata for the struct [RecordingGetOneRecordingResponseDataStopReasonCaller]
type recordingGetOneRecordingResponseDataStopReasonCallerJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetOneRecordingResponseDataStopReasonCaller) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetOneRecordingResponseDataStopReasonCallerJSON) RawJSON() string {
	return r.raw
}

// The type can be an App or a user. If the type is `user`, then only the `user_Id`
// and `name` are returned.
type RecordingGetOneRecordingResponseDataStopReasonCallerType string

const (
	RecordingGetOneRecordingResponseDataStopReasonCallerTypeOrganization RecordingGetOneRecordingResponseDataStopReasonCallerType = "ORGANIZATION"
	RecordingGetOneRecordingResponseDataStopReasonCallerTypeUser         RecordingGetOneRecordingResponseDataStopReasonCallerType = "USER"
)

func (r RecordingGetOneRecordingResponseDataStopReasonCallerType) IsKnown() bool {
	switch r {
	case RecordingGetOneRecordingResponseDataStopReasonCallerTypeOrganization, RecordingGetOneRecordingResponseDataStopReasonCallerTypeUser:
		return true
	}
	return false
}

// Specifies the reason why the recording stopped.
type RecordingGetOneRecordingResponseDataStopReasonReason string

const (
	RecordingGetOneRecordingResponseDataStopReasonReasonAPICall       RecordingGetOneRecordingResponseDataStopReasonReason = "API_CALL"
	RecordingGetOneRecordingResponseDataStopReasonReasonInternalError RecordingGetOneRecordingResponseDataStopReasonReason = "INTERNAL_ERROR"
	RecordingGetOneRecordingResponseDataStopReasonReasonAllPeersLeft  RecordingGetOneRecordingResponseDataStopReasonReason = "ALL_PEERS_LEFT"
)

func (r RecordingGetOneRecordingResponseDataStopReasonReason) IsKnown() bool {
	switch r {
	case RecordingGetOneRecordingResponseDataStopReasonReasonAPICall, RecordingGetOneRecordingResponseDataStopReasonReasonInternalError, RecordingGetOneRecordingResponseDataStopReasonReasonAllPeersLeft:
		return true
	}
	return false
}

type RecordingGetOneRecordingResponseDataStorageConfig struct {
	// This field can have the runtime type of [string], [interface{}].
	AccessKey interface{} `json:"access_key"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingGetOneRecordingResponseDataStorageConfigAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// This field can have the runtime type of [string], [interface{}].
	Region interface{}                                           `json:"region"`
	Type   RecordingGetOneRecordingResponseDataStorageConfigType `json:"type"`
	// SSH destination server username for SFTP type storage medium
	Username string                                                `json:"username"`
	JSON     recordingGetOneRecordingResponseDataStorageConfigJSON `json:"-"`
	union    RecordingGetOneRecordingResponseDataStorageConfigUnion
}

// recordingGetOneRecordingResponseDataStorageConfigJSON contains the JSON metadata
// for the struct [RecordingGetOneRecordingResponseDataStorageConfig]
type recordingGetOneRecordingResponseDataStorageConfigJSON struct {
	AccessKey   apijson.Field
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r recordingGetOneRecordingResponseDataStorageConfigJSON) RawJSON() string {
	return r.raw
}

func (r *RecordingGetOneRecordingResponseDataStorageConfig) UnmarshalJSON(data []byte) (err error) {
	*r = RecordingGetOneRecordingResponseDataStorageConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecordingGetOneRecordingResponseDataStorageConfigUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RecordingGetOneRecordingResponseDataStorageConfigObject],
// [RecordingGetOneRecordingResponseDataStorageConfigObject],
// [RecordingGetOneRecordingResponseDataStorageConfigObject],
// [RecordingGetOneRecordingResponseDataStorageConfigObject].
func (r RecordingGetOneRecordingResponseDataStorageConfig) AsUnion() RecordingGetOneRecordingResponseDataStorageConfigUnion {
	return r.union
}

// Union satisfied by [RecordingGetOneRecordingResponseDataStorageConfigObject],
// [RecordingGetOneRecordingResponseDataStorageConfigObject],
// [RecordingGetOneRecordingResponseDataStorageConfigObject] or
// [RecordingGetOneRecordingResponseDataStorageConfigObject].
type RecordingGetOneRecordingResponseDataStorageConfigUnion interface {
	implementsRecordingGetOneRecordingResponseDataStorageConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordingGetOneRecordingResponseDataStorageConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingGetOneRecordingResponseDataStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingGetOneRecordingResponseDataStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingGetOneRecordingResponseDataStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingGetOneRecordingResponseDataStorageConfigObject{}),
		},
	)
}

type RecordingGetOneRecordingResponseDataStorageConfigObject struct {
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingGetOneRecordingResponseDataStorageConfigObjectAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// Region of the storage medium.
	Region string                                                      `json:"region"`
	Type   RecordingGetOneRecordingResponseDataStorageConfigObjectType `json:"type"`
	// SSH destination server username for SFTP type storage medium
	Username string                                                      `json:"username"`
	JSON     recordingGetOneRecordingResponseDataStorageConfigObjectJSON `json:"-"`
}

// recordingGetOneRecordingResponseDataStorageConfigObjectJSON contains the JSON
// metadata for the struct
// [RecordingGetOneRecordingResponseDataStorageConfigObject]
type recordingGetOneRecordingResponseDataStorageConfigObjectJSON struct {
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetOneRecordingResponseDataStorageConfigObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetOneRecordingResponseDataStorageConfigObjectJSON) RawJSON() string {
	return r.raw
}

func (r RecordingGetOneRecordingResponseDataStorageConfigObject) implementsRecordingGetOneRecordingResponseDataStorageConfig() {
}

// Authentication method used for "sftp" type storage medium
type RecordingGetOneRecordingResponseDataStorageConfigObjectAuthMethod string

const (
	RecordingGetOneRecordingResponseDataStorageConfigObjectAuthMethodKey      RecordingGetOneRecordingResponseDataStorageConfigObjectAuthMethod = "KEY"
	RecordingGetOneRecordingResponseDataStorageConfigObjectAuthMethodPassword RecordingGetOneRecordingResponseDataStorageConfigObjectAuthMethod = "PASSWORD"
)

func (r RecordingGetOneRecordingResponseDataStorageConfigObjectAuthMethod) IsKnown() bool {
	switch r {
	case RecordingGetOneRecordingResponseDataStorageConfigObjectAuthMethodKey, RecordingGetOneRecordingResponseDataStorageConfigObjectAuthMethodPassword:
		return true
	}
	return false
}

type RecordingGetOneRecordingResponseDataStorageConfigObjectType string

const (
	RecordingGetOneRecordingResponseDataStorageConfigObjectTypeGcs RecordingGetOneRecordingResponseDataStorageConfigObjectType = "gcs"
)

func (r RecordingGetOneRecordingResponseDataStorageConfigObjectType) IsKnown() bool {
	switch r {
	case RecordingGetOneRecordingResponseDataStorageConfigObjectTypeGcs:
		return true
	}
	return false
}

// Authentication method used for "sftp" type storage medium
type RecordingGetOneRecordingResponseDataStorageConfigAuthMethod string

const (
	RecordingGetOneRecordingResponseDataStorageConfigAuthMethodKey      RecordingGetOneRecordingResponseDataStorageConfigAuthMethod = "KEY"
	RecordingGetOneRecordingResponseDataStorageConfigAuthMethodPassword RecordingGetOneRecordingResponseDataStorageConfigAuthMethod = "PASSWORD"
)

func (r RecordingGetOneRecordingResponseDataStorageConfigAuthMethod) IsKnown() bool {
	switch r {
	case RecordingGetOneRecordingResponseDataStorageConfigAuthMethodKey, RecordingGetOneRecordingResponseDataStorageConfigAuthMethodPassword:
		return true
	}
	return false
}

type RecordingGetOneRecordingResponseDataStorageConfigType string

const (
	RecordingGetOneRecordingResponseDataStorageConfigTypeGcs          RecordingGetOneRecordingResponseDataStorageConfigType = "gcs"
	RecordingGetOneRecordingResponseDataStorageConfigTypeAws          RecordingGetOneRecordingResponseDataStorageConfigType = "aws"
	RecordingGetOneRecordingResponseDataStorageConfigTypeAzure        RecordingGetOneRecordingResponseDataStorageConfigType = "azure"
	RecordingGetOneRecordingResponseDataStorageConfigTypeDigitalocean RecordingGetOneRecordingResponseDataStorageConfigType = "digitalocean"
	RecordingGetOneRecordingResponseDataStorageConfigTypeSftp         RecordingGetOneRecordingResponseDataStorageConfigType = "sftp"
)

func (r RecordingGetOneRecordingResponseDataStorageConfigType) IsKnown() bool {
	switch r {
	case RecordingGetOneRecordingResponseDataStorageConfigTypeGcs, RecordingGetOneRecordingResponseDataStorageConfigTypeAws, RecordingGetOneRecordingResponseDataStorageConfigTypeAzure, RecordingGetOneRecordingResponseDataStorageConfigTypeDigitalocean, RecordingGetOneRecordingResponseDataStorageConfigTypeSftp:
		return true
	}
	return false
}

type RecordingGetRecordingsResponse struct {
	Data    []RecordingGetRecordingsResponseData `json:"data" api:"required"`
	Paging  RecordingGetRecordingsResponsePaging `json:"paging" api:"required"`
	Success bool                                 `json:"success" api:"required"`
	JSON    recordingGetRecordingsResponseJSON   `json:"-"`
}

// recordingGetRecordingsResponseJSON contains the JSON metadata for the struct
// [RecordingGetRecordingsResponse]
type recordingGetRecordingsResponseJSON struct {
	Data        apijson.Field
	Paging      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetRecordingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseJSON) RawJSON() string {
	return r.raw
}

type RecordingGetRecordingsResponseData struct {
	// ID of the recording
	ID string `json:"id" api:"required" format:"uuid"`
	// If the audio_config is passed, the URL for downloading the audio recording is
	// returned.
	AudioDownloadURL string `json:"audio_download_url" api:"required,nullable" format:"uri"`
	// URL where the recording can be downloaded.
	DownloadURL string `json:"download_url" api:"required,nullable" format:"uri"`
	// Timestamp when the download URL expires.
	DownloadURLExpiry time.Time `json:"download_url_expiry" api:"required,nullable" format:"date-time"`
	// File size of the recording, in bytes.
	FileSize float64 `json:"file_size" api:"required,nullable"`
	// Timestamp when this recording was invoked.
	InvokedTime time.Time `json:"invoked_time" api:"required" format:"date-time"`
	// File name of the recording.
	OutputFileName string `json:"output_file_name" api:"required"`
	// ID of the meeting session this recording is for.
	SessionID string `json:"session_id" api:"required,nullable" format:"uuid"`
	// Timestamp when this recording actually started after being invoked. Usually a
	// few seconds after `invoked_time`.
	StartedTime time.Time `json:"started_time" api:"required,nullable" format:"date-time"`
	// Current status of the recording.
	Status RecordingGetRecordingsResponseDataStatus `json:"status" api:"required"`
	// Timestamp when this recording was stopped. Optional; is present only when the
	// recording has actually been stopped.
	StoppedTime time.Time                                 `json:"stopped_time" api:"required,nullable" format:"date-time"`
	Meeting     RecordingGetRecordingsResponseDataMeeting `json:"meeting"`
	// Total recording time in seconds.
	RecordingDuration int64                                           `json:"recording_duration"`
	StorageConfig     RecordingGetRecordingsResponseDataStorageConfig `json:"storage_config" api:"nullable"`
	JSON              recordingGetRecordingsResponseDataJSON          `json:"-"`
}

// recordingGetRecordingsResponseDataJSON contains the JSON metadata for the struct
// [RecordingGetRecordingsResponseData]
type recordingGetRecordingsResponseDataJSON struct {
	ID                apijson.Field
	AudioDownloadURL  apijson.Field
	DownloadURL       apijson.Field
	DownloadURLExpiry apijson.Field
	FileSize          apijson.Field
	InvokedTime       apijson.Field
	OutputFileName    apijson.Field
	SessionID         apijson.Field
	StartedTime       apijson.Field
	Status            apijson.Field
	StoppedTime       apijson.Field
	Meeting           apijson.Field
	RecordingDuration apijson.Field
	StorageConfig     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordingGetRecordingsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseDataJSON) RawJSON() string {
	return r.raw
}

// Current status of the recording.
type RecordingGetRecordingsResponseDataStatus string

const (
	RecordingGetRecordingsResponseDataStatusInvoked   RecordingGetRecordingsResponseDataStatus = "INVOKED"
	RecordingGetRecordingsResponseDataStatusRecording RecordingGetRecordingsResponseDataStatus = "RECORDING"
	RecordingGetRecordingsResponseDataStatusUploading RecordingGetRecordingsResponseDataStatus = "UPLOADING"
	RecordingGetRecordingsResponseDataStatusUploaded  RecordingGetRecordingsResponseDataStatus = "UPLOADED"
	RecordingGetRecordingsResponseDataStatusErrored   RecordingGetRecordingsResponseDataStatus = "ERRORED"
	RecordingGetRecordingsResponseDataStatusPaused    RecordingGetRecordingsResponseDataStatus = "PAUSED"
)

func (r RecordingGetRecordingsResponseDataStatus) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataStatusInvoked, RecordingGetRecordingsResponseDataStatusRecording, RecordingGetRecordingsResponseDataStatusUploading, RecordingGetRecordingsResponseDataStatusUploaded, RecordingGetRecordingsResponseDataStatusErrored, RecordingGetRecordingsResponseDataStatusPaused:
		return true
	}
	return false
}

type RecordingGetRecordingsResponseDataMeeting struct {
	// ID of the meeting.
	ID string `json:"id" api:"required" format:"uuid"`
	// Timestamp the object was created at. The time is returned in ISO format.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Timestamp the object was updated at. The time is returned in ISO format.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
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
	RecordingConfig RecordingGetRecordingsResponseDataMeetingRecordingConfig `json:"recording_config"`
	// Time in seconds, for which a session remains active, after the last participant
	// has left the meeting.
	SessionKeepAliveTimeInSecs float64 `json:"session_keep_alive_time_in_secs"`
	// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
	// `INACTIVE` meeting.
	Status RecordingGetRecordingsResponseDataMeetingStatus `json:"status"`
	// Automatically generate summary of meetings using transcripts. Requires
	// Transcriptions to be enabled, and can be retrieved via Webhooks or summary API.
	SummarizeOnEnd bool `json:"summarize_on_end"`
	// Title of the meeting.
	Title string `json:"title"`
	// Automatically generate transcripts when the meeting ends.
	TranscribeOnEnd bool                                          `json:"transcribe_on_end"`
	JSON            recordingGetRecordingsResponseDataMeetingJSON `json:"-"`
}

// recordingGetRecordingsResponseDataMeetingJSON contains the JSON metadata for the
// struct [RecordingGetRecordingsResponseDataMeeting]
type recordingGetRecordingsResponseDataMeetingJSON struct {
	ID                         apijson.Field
	CreatedAt                  apijson.Field
	UpdatedAt                  apijson.Field
	LiveStreamOnStart          apijson.Field
	PersistChat                apijson.Field
	RecordOnStart              apijson.Field
	RecordingConfig            apijson.Field
	SessionKeepAliveTimeInSecs apijson.Field
	Status                     apijson.Field
	SummarizeOnEnd             apijson.Field
	Title                      apijson.Field
	TranscribeOnEnd            apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *RecordingGetRecordingsResponseDataMeeting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseDataMeetingJSON) RawJSON() string {
	return r.raw
}

// Recording Configurations to be used for this meeting. This level of configs
// takes higher preference over App level configs on the RealtimeKit developer
// portal.
type RecordingGetRecordingsResponseDataMeetingRecordingConfig struct {
	// Object containing configuration regarding the audio that is being recorded.
	AudioConfig RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfig `json:"audio_config"`
	// Adds a prefix to the beginning of the file name of the recording.
	FileNamePrefix      string                                                                      `json:"file_name_prefix"`
	LiveStreamingConfig RecordingGetRecordingsResponseDataMeetingRecordingConfigLiveStreamingConfig `json:"live_streaming_config"`
	// Specifies the maximum duration for recording in seconds, ranging from a minimum
	// of 60 seconds to a maximum of 24 hours.
	MaxSeconds              float64                                                                         `json:"max_seconds"`
	RealtimekitBucketConfig RecordingGetRecordingsResponseDataMeetingRecordingConfigRealtimekitBucketConfig `json:"realtimekit_bucket_config"`
	StorageConfig           RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfig           `json:"storage_config" api:"nullable"`
	VideoConfig             RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfig             `json:"video_config"`
	JSON                    recordingGetRecordingsResponseDataMeetingRecordingConfigJSON                    `json:"-"`
}

// recordingGetRecordingsResponseDataMeetingRecordingConfigJSON contains the JSON
// metadata for the struct
// [RecordingGetRecordingsResponseDataMeetingRecordingConfig]
type recordingGetRecordingsResponseDataMeetingRecordingConfigJSON struct {
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

func (r *RecordingGetRecordingsResponseDataMeetingRecordingConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseDataMeetingRecordingConfigJSON) RawJSON() string {
	return r.raw
}

// Object containing configuration regarding the audio that is being recorded.
type RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfig struct {
	// Audio signal pathway within an audio file that carries a specific sound source.
	Channel RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigChannel `json:"channel"`
	// Codec using which the recording will be encoded. If VP8/VP9 is selected for
	// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
	// audioConfig is automatically set to vorbis.
	Codec RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigCodec `json:"codec"`
	// Controls whether to export audio file seperately
	ExportFile bool                                                                    `json:"export_file"`
	JSON       recordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigJSON `json:"-"`
}

// recordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigJSON contains
// the JSON metadata for the struct
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfig]
type recordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigJSON struct {
	Channel     apijson.Field
	Codec       apijson.Field
	ExportFile  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigJSON) RawJSON() string {
	return r.raw
}

// Audio signal pathway within an audio file that carries a specific sound source.
type RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigChannel string

const (
	RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigChannelMono   RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigChannel = "mono"
	RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigChannelStereo RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigChannel = "stereo"
)

func (r RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigChannel) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigChannelMono, RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigChannelStereo:
		return true
	}
	return false
}

// Codec using which the recording will be encoded. If VP8/VP9 is selected for
// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
// audioConfig is automatically set to vorbis.
type RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigCodec string

const (
	RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigCodecMP3 RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigCodec = "MP3"
	RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigCodecAac RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigCodec = "AAC"
)

func (r RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigCodec) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigCodecMP3, RecordingGetRecordingsResponseDataMeetingRecordingConfigAudioConfigCodecAac:
		return true
	}
	return false
}

type RecordingGetRecordingsResponseDataMeetingRecordingConfigLiveStreamingConfig struct {
	// RTMP URL to stream to
	RtmpURL string                                                                          `json:"rtmp_url" format:"uri"`
	JSON    recordingGetRecordingsResponseDataMeetingRecordingConfigLiveStreamingConfigJSON `json:"-"`
}

// recordingGetRecordingsResponseDataMeetingRecordingConfigLiveStreamingConfigJSON
// contains the JSON metadata for the struct
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigLiveStreamingConfig]
type recordingGetRecordingsResponseDataMeetingRecordingConfigLiveStreamingConfigJSON struct {
	RtmpURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetRecordingsResponseDataMeetingRecordingConfigLiveStreamingConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseDataMeetingRecordingConfigLiveStreamingConfigJSON) RawJSON() string {
	return r.raw
}

type RecordingGetRecordingsResponseDataMeetingRecordingConfigRealtimekitBucketConfig struct {
	// Controls whether recordings are uploaded to RealtimeKit's bucket. If set to
	// false, `download_url`, `audio_download_url`, `download_url_expiry` won't be
	// generated for a recording.
	Enabled bool                                                                                `json:"enabled" api:"required"`
	JSON    recordingGetRecordingsResponseDataMeetingRecordingConfigRealtimekitBucketConfigJSON `json:"-"`
}

// recordingGetRecordingsResponseDataMeetingRecordingConfigRealtimekitBucketConfigJSON
// contains the JSON metadata for the struct
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigRealtimekitBucketConfig]
type recordingGetRecordingsResponseDataMeetingRecordingConfigRealtimekitBucketConfigJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetRecordingsResponseDataMeetingRecordingConfigRealtimekitBucketConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseDataMeetingRecordingConfigRealtimekitBucketConfigJSON) RawJSON() string {
	return r.raw
}

type RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfig struct {
	// This field can have the runtime type of [string], [interface{}].
	AccessKey interface{} `json:"access_key"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// This field can have the runtime type of [string], [interface{}].
	Region interface{}                                                               `json:"region"`
	Type   RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigType `json:"type"`
	// SSH destination server username for SFTP type storage medium
	Username string                                                                    `json:"username"`
	JSON     recordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigJSON `json:"-"`
	union    RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigUnion
}

// recordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigJSON
// contains the JSON metadata for the struct
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfig]
type recordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigJSON struct {
	AccessKey   apijson.Field
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r recordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigJSON) RawJSON() string {
	return r.raw
}

func (r *RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfig) UnmarshalJSON(data []byte) (err error) {
	*r = RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject],
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject],
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject],
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject].
func (r RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfig) AsUnion() RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigUnion {
	return r.union
}

// Union satisfied by
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject],
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject],
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject] or
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject].
type RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigUnion interface {
	implementsRecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject{}),
		},
	)
}

type RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject struct {
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// Region of the storage medium.
	Region string                                                                          `json:"region"`
	Type   RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectType `json:"type"`
	// SSH destination server username for SFTP type storage medium
	Username string                                                                          `json:"username"`
	JSON     recordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectJSON `json:"-"`
}

// recordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectJSON
// contains the JSON metadata for the struct
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject]
type recordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectJSON struct {
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectJSON) RawJSON() string {
	return r.raw
}

func (r RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObject) implementsRecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfig() {
}

// Authentication method used for "sftp" type storage medium
type RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectAuthMethod string

const (
	RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectAuthMethodKey      RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectAuthMethod = "KEY"
	RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectAuthMethodPassword RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectAuthMethod = "PASSWORD"
)

func (r RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectAuthMethod) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectAuthMethodKey, RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectAuthMethodPassword:
		return true
	}
	return false
}

type RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectType string

const (
	RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectTypeGcs RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectType = "gcs"
)

func (r RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectType) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigObjectTypeGcs:
		return true
	}
	return false
}

// Authentication method used for "sftp" type storage medium
type RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigAuthMethod string

const (
	RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigAuthMethodKey      RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigAuthMethod = "KEY"
	RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigAuthMethodPassword RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigAuthMethod = "PASSWORD"
)

func (r RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigAuthMethod) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigAuthMethodKey, RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigAuthMethodPassword:
		return true
	}
	return false
}

type RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigType string

const (
	RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigTypeGcs          RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigType = "gcs"
	RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigTypeAws          RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigType = "aws"
	RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigTypeAzure        RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigType = "azure"
	RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigTypeDigitalocean RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigType = "digitalocean"
	RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigTypeSftp         RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigType = "sftp"
)

func (r RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigType) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigTypeGcs, RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigTypeAws, RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigTypeAzure, RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigTypeDigitalocean, RecordingGetRecordingsResponseDataMeetingRecordingConfigStorageConfigTypeSftp:
		return true
	}
	return false
}

type RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfig struct {
	// Codec using which the recording will be encoded.
	Codec RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigCodec `json:"codec"`
	// Controls whether to export video file seperately
	ExportFile bool `json:"export_file"`
	// Height of the recording video in pixels
	Height int64 `json:"height"`
	// Watermark to be added to the recording
	Watermark RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermark `json:"watermark"`
	// Width of the recording video in pixels
	Width int64                                                                   `json:"width"`
	JSON  recordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigJSON `json:"-"`
}

// recordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigJSON contains
// the JSON metadata for the struct
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfig]
type recordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigJSON struct {
	Codec       apijson.Field
	ExportFile  apijson.Field
	Height      apijson.Field
	Watermark   apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigJSON) RawJSON() string {
	return r.raw
}

// Codec using which the recording will be encoded.
type RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigCodec string

const (
	RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigCodecH264 RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigCodec = "H264"
	RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigCodecVp8  RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigCodec = "VP8"
	RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigCodecVp9  RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigCodec = "VP9"
)

func (r RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigCodec) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigCodecH264, RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigCodecVp8, RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigCodecVp9:
		return true
	}
	return false
}

// Watermark to be added to the recording
type RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermark struct {
	// Position of the watermark
	Position RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPosition `json:"position"`
	// Size of the watermark
	Size RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkSize `json:"size"`
	// URL of the watermark image
	URL  string                                                                           `json:"url" format:"uri"`
	JSON recordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkJSON `json:"-"`
}

// recordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkJSON
// contains the JSON metadata for the struct
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermark]
type recordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkJSON struct {
	Position    apijson.Field
	Size        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkJSON) RawJSON() string {
	return r.raw
}

// Position of the watermark
type RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPosition string

const (
	RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPositionLeftTop     RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPosition = "left top"
	RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPositionRightTop    RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPosition = "right top"
	RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPositionLeftBottom  RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPosition = "left bottom"
	RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPositionRightBottom RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPosition = "right bottom"
)

func (r RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPosition) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPositionLeftTop, RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPositionRightTop, RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPositionLeftBottom, RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkPositionRightBottom:
		return true
	}
	return false
}

// Size of the watermark
type RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkSize struct {
	// Height of the watermark in px
	Height int64 `json:"height"`
	// Width of the watermark in px
	Width int64                                                                                `json:"width"`
	JSON  recordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkSizeJSON `json:"-"`
}

// recordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkSizeJSON
// contains the JSON metadata for the struct
// [RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkSize]
type recordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkSizeJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkSize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseDataMeetingRecordingConfigVideoConfigWatermarkSizeJSON) RawJSON() string {
	return r.raw
}

// Whether the meeting is `ACTIVE` or `INACTIVE`. Users will not be able to join an
// `INACTIVE` meeting.
type RecordingGetRecordingsResponseDataMeetingStatus string

const (
	RecordingGetRecordingsResponseDataMeetingStatusActive   RecordingGetRecordingsResponseDataMeetingStatus = "ACTIVE"
	RecordingGetRecordingsResponseDataMeetingStatusInactive RecordingGetRecordingsResponseDataMeetingStatus = "INACTIVE"
)

func (r RecordingGetRecordingsResponseDataMeetingStatus) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataMeetingStatusActive, RecordingGetRecordingsResponseDataMeetingStatusInactive:
		return true
	}
	return false
}

type RecordingGetRecordingsResponseDataStorageConfig struct {
	// This field can have the runtime type of [string], [interface{}].
	AccessKey interface{} `json:"access_key"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingGetRecordingsResponseDataStorageConfigAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// This field can have the runtime type of [string], [interface{}].
	Region interface{}                                         `json:"region"`
	Type   RecordingGetRecordingsResponseDataStorageConfigType `json:"type"`
	// SSH destination server username for SFTP type storage medium
	Username string                                              `json:"username"`
	JSON     recordingGetRecordingsResponseDataStorageConfigJSON `json:"-"`
	union    RecordingGetRecordingsResponseDataStorageConfigUnion
}

// recordingGetRecordingsResponseDataStorageConfigJSON contains the JSON metadata
// for the struct [RecordingGetRecordingsResponseDataStorageConfig]
type recordingGetRecordingsResponseDataStorageConfigJSON struct {
	AccessKey   apijson.Field
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r recordingGetRecordingsResponseDataStorageConfigJSON) RawJSON() string {
	return r.raw
}

func (r *RecordingGetRecordingsResponseDataStorageConfig) UnmarshalJSON(data []byte) (err error) {
	*r = RecordingGetRecordingsResponseDataStorageConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecordingGetRecordingsResponseDataStorageConfigUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RecordingGetRecordingsResponseDataStorageConfigObject],
// [RecordingGetRecordingsResponseDataStorageConfigObject],
// [RecordingGetRecordingsResponseDataStorageConfigObject],
// [RecordingGetRecordingsResponseDataStorageConfigObject].
func (r RecordingGetRecordingsResponseDataStorageConfig) AsUnion() RecordingGetRecordingsResponseDataStorageConfigUnion {
	return r.union
}

// Union satisfied by [RecordingGetRecordingsResponseDataStorageConfigObject],
// [RecordingGetRecordingsResponseDataStorageConfigObject],
// [RecordingGetRecordingsResponseDataStorageConfigObject] or
// [RecordingGetRecordingsResponseDataStorageConfigObject].
type RecordingGetRecordingsResponseDataStorageConfigUnion interface {
	implementsRecordingGetRecordingsResponseDataStorageConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordingGetRecordingsResponseDataStorageConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingGetRecordingsResponseDataStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingGetRecordingsResponseDataStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingGetRecordingsResponseDataStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingGetRecordingsResponseDataStorageConfigObject{}),
		},
	)
}

type RecordingGetRecordingsResponseDataStorageConfigObject struct {
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingGetRecordingsResponseDataStorageConfigObjectAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// Region of the storage medium.
	Region string                                                    `json:"region"`
	Type   RecordingGetRecordingsResponseDataStorageConfigObjectType `json:"type"`
	// SSH destination server username for SFTP type storage medium
	Username string                                                    `json:"username"`
	JSON     recordingGetRecordingsResponseDataStorageConfigObjectJSON `json:"-"`
}

// recordingGetRecordingsResponseDataStorageConfigObjectJSON contains the JSON
// metadata for the struct [RecordingGetRecordingsResponseDataStorageConfigObject]
type recordingGetRecordingsResponseDataStorageConfigObjectJSON struct {
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetRecordingsResponseDataStorageConfigObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseDataStorageConfigObjectJSON) RawJSON() string {
	return r.raw
}

func (r RecordingGetRecordingsResponseDataStorageConfigObject) implementsRecordingGetRecordingsResponseDataStorageConfig() {
}

// Authentication method used for "sftp" type storage medium
type RecordingGetRecordingsResponseDataStorageConfigObjectAuthMethod string

const (
	RecordingGetRecordingsResponseDataStorageConfigObjectAuthMethodKey      RecordingGetRecordingsResponseDataStorageConfigObjectAuthMethod = "KEY"
	RecordingGetRecordingsResponseDataStorageConfigObjectAuthMethodPassword RecordingGetRecordingsResponseDataStorageConfigObjectAuthMethod = "PASSWORD"
)

func (r RecordingGetRecordingsResponseDataStorageConfigObjectAuthMethod) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataStorageConfigObjectAuthMethodKey, RecordingGetRecordingsResponseDataStorageConfigObjectAuthMethodPassword:
		return true
	}
	return false
}

type RecordingGetRecordingsResponseDataStorageConfigObjectType string

const (
	RecordingGetRecordingsResponseDataStorageConfigObjectTypeGcs RecordingGetRecordingsResponseDataStorageConfigObjectType = "gcs"
)

func (r RecordingGetRecordingsResponseDataStorageConfigObjectType) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataStorageConfigObjectTypeGcs:
		return true
	}
	return false
}

// Authentication method used for "sftp" type storage medium
type RecordingGetRecordingsResponseDataStorageConfigAuthMethod string

const (
	RecordingGetRecordingsResponseDataStorageConfigAuthMethodKey      RecordingGetRecordingsResponseDataStorageConfigAuthMethod = "KEY"
	RecordingGetRecordingsResponseDataStorageConfigAuthMethodPassword RecordingGetRecordingsResponseDataStorageConfigAuthMethod = "PASSWORD"
)

func (r RecordingGetRecordingsResponseDataStorageConfigAuthMethod) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataStorageConfigAuthMethodKey, RecordingGetRecordingsResponseDataStorageConfigAuthMethodPassword:
		return true
	}
	return false
}

type RecordingGetRecordingsResponseDataStorageConfigType string

const (
	RecordingGetRecordingsResponseDataStorageConfigTypeGcs          RecordingGetRecordingsResponseDataStorageConfigType = "gcs"
	RecordingGetRecordingsResponseDataStorageConfigTypeAws          RecordingGetRecordingsResponseDataStorageConfigType = "aws"
	RecordingGetRecordingsResponseDataStorageConfigTypeAzure        RecordingGetRecordingsResponseDataStorageConfigType = "azure"
	RecordingGetRecordingsResponseDataStorageConfigTypeDigitalocean RecordingGetRecordingsResponseDataStorageConfigType = "digitalocean"
	RecordingGetRecordingsResponseDataStorageConfigTypeSftp         RecordingGetRecordingsResponseDataStorageConfigType = "sftp"
)

func (r RecordingGetRecordingsResponseDataStorageConfigType) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataStorageConfigTypeGcs, RecordingGetRecordingsResponseDataStorageConfigTypeAws, RecordingGetRecordingsResponseDataStorageConfigTypeAzure, RecordingGetRecordingsResponseDataStorageConfigTypeDigitalocean, RecordingGetRecordingsResponseDataStorageConfigTypeSftp:
		return true
	}
	return false
}

type RecordingGetRecordingsResponsePaging struct {
	EndOffset   float64                                  `json:"end_offset" api:"required"`
	StartOffset float64                                  `json:"start_offset" api:"required"`
	TotalCount  float64                                  `json:"total_count" api:"required"`
	JSON        recordingGetRecordingsResponsePagingJSON `json:"-"`
}

// recordingGetRecordingsResponsePagingJSON contains the JSON metadata for the
// struct [RecordingGetRecordingsResponsePaging]
type recordingGetRecordingsResponsePagingJSON struct {
	EndOffset   apijson.Field
	StartOffset apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingGetRecordingsResponsePaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponsePagingJSON) RawJSON() string {
	return r.raw
}

type RecordingPauseResumeStopRecordingResponse struct {
	// Success status of the operation
	Success bool `json:"success" api:"required"`
	// Data returned by the operation
	Data RecordingPauseResumeStopRecordingResponseData `json:"data"`
	JSON recordingPauseResumeStopRecordingResponseJSON `json:"-"`
}

// recordingPauseResumeStopRecordingResponseJSON contains the JSON metadata for the
// struct [RecordingPauseResumeStopRecordingResponse]
type recordingPauseResumeStopRecordingResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingPauseResumeStopRecordingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingPauseResumeStopRecordingResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type RecordingPauseResumeStopRecordingResponseData struct {
	// ID of the recording
	ID string `json:"id" api:"required" format:"uuid"`
	// If the audio_config is passed, the URL for downloading the audio recording is
	// returned.
	AudioDownloadURL string `json:"audio_download_url" api:"required,nullable" format:"uri"`
	// URL where the recording can be downloaded.
	DownloadURL string `json:"download_url" api:"required,nullable" format:"uri"`
	// Timestamp when the download URL expires.
	DownloadURLExpiry time.Time `json:"download_url_expiry" api:"required,nullable" format:"date-time"`
	// File size of the recording, in bytes.
	FileSize float64 `json:"file_size" api:"required,nullable"`
	// Timestamp when this recording was invoked.
	InvokedTime time.Time `json:"invoked_time" api:"required" format:"date-time"`
	// File name of the recording.
	OutputFileName string `json:"output_file_name" api:"required"`
	// ID of the meeting session this recording is for.
	SessionID string `json:"session_id" api:"required,nullable" format:"uuid"`
	// Timestamp when this recording actually started after being invoked. Usually a
	// few seconds after `invoked_time`.
	StartedTime time.Time `json:"started_time" api:"required,nullable" format:"date-time"`
	// Current status of the recording.
	Status RecordingPauseResumeStopRecordingResponseDataStatus `json:"status" api:"required"`
	// Timestamp when this recording was stopped. Optional; is present only when the
	// recording has actually been stopped.
	StoppedTime time.Time `json:"stopped_time" api:"required,nullable" format:"date-time"`
	// Total recording time in seconds.
	RecordingDuration int64                                                      `json:"recording_duration"`
	StartReason       RecordingPauseResumeStopRecordingResponseDataStartReason   `json:"start_reason"`
	StopReason        RecordingPauseResumeStopRecordingResponseDataStopReason    `json:"stop_reason"`
	StorageConfig     RecordingPauseResumeStopRecordingResponseDataStorageConfig `json:"storage_config" api:"nullable"`
	JSON              recordingPauseResumeStopRecordingResponseDataJSON          `json:"-"`
}

// recordingPauseResumeStopRecordingResponseDataJSON contains the JSON metadata for
// the struct [RecordingPauseResumeStopRecordingResponseData]
type recordingPauseResumeStopRecordingResponseDataJSON struct {
	ID                apijson.Field
	AudioDownloadURL  apijson.Field
	DownloadURL       apijson.Field
	DownloadURLExpiry apijson.Field
	FileSize          apijson.Field
	InvokedTime       apijson.Field
	OutputFileName    apijson.Field
	SessionID         apijson.Field
	StartedTime       apijson.Field
	Status            apijson.Field
	StoppedTime       apijson.Field
	RecordingDuration apijson.Field
	StartReason       apijson.Field
	StopReason        apijson.Field
	StorageConfig     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordingPauseResumeStopRecordingResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingPauseResumeStopRecordingResponseDataJSON) RawJSON() string {
	return r.raw
}

// Current status of the recording.
type RecordingPauseResumeStopRecordingResponseDataStatus string

const (
	RecordingPauseResumeStopRecordingResponseDataStatusInvoked   RecordingPauseResumeStopRecordingResponseDataStatus = "INVOKED"
	RecordingPauseResumeStopRecordingResponseDataStatusRecording RecordingPauseResumeStopRecordingResponseDataStatus = "RECORDING"
	RecordingPauseResumeStopRecordingResponseDataStatusUploading RecordingPauseResumeStopRecordingResponseDataStatus = "UPLOADING"
	RecordingPauseResumeStopRecordingResponseDataStatusUploaded  RecordingPauseResumeStopRecordingResponseDataStatus = "UPLOADED"
	RecordingPauseResumeStopRecordingResponseDataStatusErrored   RecordingPauseResumeStopRecordingResponseDataStatus = "ERRORED"
	RecordingPauseResumeStopRecordingResponseDataStatusPaused    RecordingPauseResumeStopRecordingResponseDataStatus = "PAUSED"
)

func (r RecordingPauseResumeStopRecordingResponseDataStatus) IsKnown() bool {
	switch r {
	case RecordingPauseResumeStopRecordingResponseDataStatusInvoked, RecordingPauseResumeStopRecordingResponseDataStatusRecording, RecordingPauseResumeStopRecordingResponseDataStatusUploading, RecordingPauseResumeStopRecordingResponseDataStatusUploaded, RecordingPauseResumeStopRecordingResponseDataStatusErrored, RecordingPauseResumeStopRecordingResponseDataStatusPaused:
		return true
	}
	return false
}

type RecordingPauseResumeStopRecordingResponseDataStartReason struct {
	Caller RecordingPauseResumeStopRecordingResponseDataStartReasonCaller `json:"caller"`
	// Specifies if the recording was started using the "Start a Recording"API or using
	// the parameter RECORD_ON_START in the "Create a meeting" API.
	//
	// If the recording is initiated using the "RECORD_ON_START" parameter, the user
	// details will not be populated.
	Reason RecordingPauseResumeStopRecordingResponseDataStartReasonReason `json:"reason"`
	JSON   recordingPauseResumeStopRecordingResponseDataStartReasonJSON   `json:"-"`
}

// recordingPauseResumeStopRecordingResponseDataStartReasonJSON contains the JSON
// metadata for the struct
// [RecordingPauseResumeStopRecordingResponseDataStartReason]
type recordingPauseResumeStopRecordingResponseDataStartReasonJSON struct {
	Caller      apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingPauseResumeStopRecordingResponseDataStartReason) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingPauseResumeStopRecordingResponseDataStartReasonJSON) RawJSON() string {
	return r.raw
}

type RecordingPauseResumeStopRecordingResponseDataStartReasonCaller struct {
	// Name of the user who started the recording.
	Name string `json:"name"`
	// The type can be an App or a user. If the type is `user`, then only the `user_Id`
	// and `name` are returned.
	Type RecordingPauseResumeStopRecordingResponseDataStartReasonCallerType `json:"type"`
	// The user ID of the person who started the recording.
	UserID string                                                             `json:"user_Id" format:"uuid"`
	JSON   recordingPauseResumeStopRecordingResponseDataStartReasonCallerJSON `json:"-"`
}

// recordingPauseResumeStopRecordingResponseDataStartReasonCallerJSON contains the
// JSON metadata for the struct
// [RecordingPauseResumeStopRecordingResponseDataStartReasonCaller]
type recordingPauseResumeStopRecordingResponseDataStartReasonCallerJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingPauseResumeStopRecordingResponseDataStartReasonCaller) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingPauseResumeStopRecordingResponseDataStartReasonCallerJSON) RawJSON() string {
	return r.raw
}

// The type can be an App or a user. If the type is `user`, then only the `user_Id`
// and `name` are returned.
type RecordingPauseResumeStopRecordingResponseDataStartReasonCallerType string

const (
	RecordingPauseResumeStopRecordingResponseDataStartReasonCallerTypeOrganization RecordingPauseResumeStopRecordingResponseDataStartReasonCallerType = "ORGANIZATION"
	RecordingPauseResumeStopRecordingResponseDataStartReasonCallerTypeUser         RecordingPauseResumeStopRecordingResponseDataStartReasonCallerType = "USER"
)

func (r RecordingPauseResumeStopRecordingResponseDataStartReasonCallerType) IsKnown() bool {
	switch r {
	case RecordingPauseResumeStopRecordingResponseDataStartReasonCallerTypeOrganization, RecordingPauseResumeStopRecordingResponseDataStartReasonCallerTypeUser:
		return true
	}
	return false
}

// Specifies if the recording was started using the "Start a Recording"API or using
// the parameter RECORD_ON_START in the "Create a meeting" API.
//
// If the recording is initiated using the "RECORD_ON_START" parameter, the user
// details will not be populated.
type RecordingPauseResumeStopRecordingResponseDataStartReasonReason string

const (
	RecordingPauseResumeStopRecordingResponseDataStartReasonReasonAPICall       RecordingPauseResumeStopRecordingResponseDataStartReasonReason = "API_CALL"
	RecordingPauseResumeStopRecordingResponseDataStartReasonReasonRecordOnStart RecordingPauseResumeStopRecordingResponseDataStartReasonReason = "RECORD_ON_START"
)

func (r RecordingPauseResumeStopRecordingResponseDataStartReasonReason) IsKnown() bool {
	switch r {
	case RecordingPauseResumeStopRecordingResponseDataStartReasonReasonAPICall, RecordingPauseResumeStopRecordingResponseDataStartReasonReasonRecordOnStart:
		return true
	}
	return false
}

type RecordingPauseResumeStopRecordingResponseDataStopReason struct {
	Caller RecordingPauseResumeStopRecordingResponseDataStopReasonCaller `json:"caller"`
	// Specifies the reason why the recording stopped.
	Reason RecordingPauseResumeStopRecordingResponseDataStopReasonReason `json:"reason"`
	JSON   recordingPauseResumeStopRecordingResponseDataStopReasonJSON   `json:"-"`
}

// recordingPauseResumeStopRecordingResponseDataStopReasonJSON contains the JSON
// metadata for the struct
// [RecordingPauseResumeStopRecordingResponseDataStopReason]
type recordingPauseResumeStopRecordingResponseDataStopReasonJSON struct {
	Caller      apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingPauseResumeStopRecordingResponseDataStopReason) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingPauseResumeStopRecordingResponseDataStopReasonJSON) RawJSON() string {
	return r.raw
}

type RecordingPauseResumeStopRecordingResponseDataStopReasonCaller struct {
	// Name of the user who stopped the recording.
	Name string `json:"name"`
	// The type can be an App or a user. If the type is `user`, then only the `user_Id`
	// and `name` are returned.
	Type RecordingPauseResumeStopRecordingResponseDataStopReasonCallerType `json:"type"`
	// The user ID of the person who stopped the recording.
	UserID string                                                            `json:"user_Id" format:"uuid"`
	JSON   recordingPauseResumeStopRecordingResponseDataStopReasonCallerJSON `json:"-"`
}

// recordingPauseResumeStopRecordingResponseDataStopReasonCallerJSON contains the
// JSON metadata for the struct
// [RecordingPauseResumeStopRecordingResponseDataStopReasonCaller]
type recordingPauseResumeStopRecordingResponseDataStopReasonCallerJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingPauseResumeStopRecordingResponseDataStopReasonCaller) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingPauseResumeStopRecordingResponseDataStopReasonCallerJSON) RawJSON() string {
	return r.raw
}

// The type can be an App or a user. If the type is `user`, then only the `user_Id`
// and `name` are returned.
type RecordingPauseResumeStopRecordingResponseDataStopReasonCallerType string

const (
	RecordingPauseResumeStopRecordingResponseDataStopReasonCallerTypeOrganization RecordingPauseResumeStopRecordingResponseDataStopReasonCallerType = "ORGANIZATION"
	RecordingPauseResumeStopRecordingResponseDataStopReasonCallerTypeUser         RecordingPauseResumeStopRecordingResponseDataStopReasonCallerType = "USER"
)

func (r RecordingPauseResumeStopRecordingResponseDataStopReasonCallerType) IsKnown() bool {
	switch r {
	case RecordingPauseResumeStopRecordingResponseDataStopReasonCallerTypeOrganization, RecordingPauseResumeStopRecordingResponseDataStopReasonCallerTypeUser:
		return true
	}
	return false
}

// Specifies the reason why the recording stopped.
type RecordingPauseResumeStopRecordingResponseDataStopReasonReason string

const (
	RecordingPauseResumeStopRecordingResponseDataStopReasonReasonAPICall       RecordingPauseResumeStopRecordingResponseDataStopReasonReason = "API_CALL"
	RecordingPauseResumeStopRecordingResponseDataStopReasonReasonInternalError RecordingPauseResumeStopRecordingResponseDataStopReasonReason = "INTERNAL_ERROR"
	RecordingPauseResumeStopRecordingResponseDataStopReasonReasonAllPeersLeft  RecordingPauseResumeStopRecordingResponseDataStopReasonReason = "ALL_PEERS_LEFT"
)

func (r RecordingPauseResumeStopRecordingResponseDataStopReasonReason) IsKnown() bool {
	switch r {
	case RecordingPauseResumeStopRecordingResponseDataStopReasonReasonAPICall, RecordingPauseResumeStopRecordingResponseDataStopReasonReasonInternalError, RecordingPauseResumeStopRecordingResponseDataStopReasonReasonAllPeersLeft:
		return true
	}
	return false
}

type RecordingPauseResumeStopRecordingResponseDataStorageConfig struct {
	// This field can have the runtime type of [string], [interface{}].
	AccessKey interface{} `json:"access_key"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingPauseResumeStopRecordingResponseDataStorageConfigAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// This field can have the runtime type of [string], [interface{}].
	Region interface{}                                                    `json:"region"`
	Type   RecordingPauseResumeStopRecordingResponseDataStorageConfigType `json:"type"`
	// SSH destination server username for SFTP type storage medium
	Username string                                                         `json:"username"`
	JSON     recordingPauseResumeStopRecordingResponseDataStorageConfigJSON `json:"-"`
	union    RecordingPauseResumeStopRecordingResponseDataStorageConfigUnion
}

// recordingPauseResumeStopRecordingResponseDataStorageConfigJSON contains the JSON
// metadata for the struct
// [RecordingPauseResumeStopRecordingResponseDataStorageConfig]
type recordingPauseResumeStopRecordingResponseDataStorageConfigJSON struct {
	AccessKey   apijson.Field
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r recordingPauseResumeStopRecordingResponseDataStorageConfigJSON) RawJSON() string {
	return r.raw
}

func (r *RecordingPauseResumeStopRecordingResponseDataStorageConfig) UnmarshalJSON(data []byte) (err error) {
	*r = RecordingPauseResumeStopRecordingResponseDataStorageConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RecordingPauseResumeStopRecordingResponseDataStorageConfigUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RecordingPauseResumeStopRecordingResponseDataStorageConfigObject],
// [RecordingPauseResumeStopRecordingResponseDataStorageConfigObject],
// [RecordingPauseResumeStopRecordingResponseDataStorageConfigObject],
// [RecordingPauseResumeStopRecordingResponseDataStorageConfigObject].
func (r RecordingPauseResumeStopRecordingResponseDataStorageConfig) AsUnion() RecordingPauseResumeStopRecordingResponseDataStorageConfigUnion {
	return r.union
}

// Union satisfied by
// [RecordingPauseResumeStopRecordingResponseDataStorageConfigObject],
// [RecordingPauseResumeStopRecordingResponseDataStorageConfigObject],
// [RecordingPauseResumeStopRecordingResponseDataStorageConfigObject] or
// [RecordingPauseResumeStopRecordingResponseDataStorageConfigObject].
type RecordingPauseResumeStopRecordingResponseDataStorageConfigUnion interface {
	implementsRecordingPauseResumeStopRecordingResponseDataStorageConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordingPauseResumeStopRecordingResponseDataStorageConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingPauseResumeStopRecordingResponseDataStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingPauseResumeStopRecordingResponseDataStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingPauseResumeStopRecordingResponseDataStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingPauseResumeStopRecordingResponseDataStorageConfigObject{}),
		},
	)
}

type RecordingPauseResumeStopRecordingResponseDataStorageConfigObject struct {
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// Region of the storage medium.
	Region string                                                               `json:"region"`
	Type   RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectType `json:"type"`
	// SSH destination server username for SFTP type storage medium
	Username string                                                               `json:"username"`
	JSON     recordingPauseResumeStopRecordingResponseDataStorageConfigObjectJSON `json:"-"`
}

// recordingPauseResumeStopRecordingResponseDataStorageConfigObjectJSON contains
// the JSON metadata for the struct
// [RecordingPauseResumeStopRecordingResponseDataStorageConfigObject]
type recordingPauseResumeStopRecordingResponseDataStorageConfigObjectJSON struct {
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingPauseResumeStopRecordingResponseDataStorageConfigObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingPauseResumeStopRecordingResponseDataStorageConfigObjectJSON) RawJSON() string {
	return r.raw
}

func (r RecordingPauseResumeStopRecordingResponseDataStorageConfigObject) implementsRecordingPauseResumeStopRecordingResponseDataStorageConfig() {
}

// Authentication method used for "sftp" type storage medium
type RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectAuthMethod string

const (
	RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectAuthMethodKey      RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectAuthMethod = "KEY"
	RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectAuthMethodPassword RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectAuthMethod = "PASSWORD"
)

func (r RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectAuthMethod) IsKnown() bool {
	switch r {
	case RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectAuthMethodKey, RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectAuthMethodPassword:
		return true
	}
	return false
}

type RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectType string

const (
	RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectTypeGcs RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectType = "gcs"
)

func (r RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectType) IsKnown() bool {
	switch r {
	case RecordingPauseResumeStopRecordingResponseDataStorageConfigObjectTypeGcs:
		return true
	}
	return false
}

// Authentication method used for "sftp" type storage medium
type RecordingPauseResumeStopRecordingResponseDataStorageConfigAuthMethod string

const (
	RecordingPauseResumeStopRecordingResponseDataStorageConfigAuthMethodKey      RecordingPauseResumeStopRecordingResponseDataStorageConfigAuthMethod = "KEY"
	RecordingPauseResumeStopRecordingResponseDataStorageConfigAuthMethodPassword RecordingPauseResumeStopRecordingResponseDataStorageConfigAuthMethod = "PASSWORD"
)

func (r RecordingPauseResumeStopRecordingResponseDataStorageConfigAuthMethod) IsKnown() bool {
	switch r {
	case RecordingPauseResumeStopRecordingResponseDataStorageConfigAuthMethodKey, RecordingPauseResumeStopRecordingResponseDataStorageConfigAuthMethodPassword:
		return true
	}
	return false
}

type RecordingPauseResumeStopRecordingResponseDataStorageConfigType string

const (
	RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeGcs          RecordingPauseResumeStopRecordingResponseDataStorageConfigType = "gcs"
	RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeAws          RecordingPauseResumeStopRecordingResponseDataStorageConfigType = "aws"
	RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeAzure        RecordingPauseResumeStopRecordingResponseDataStorageConfigType = "azure"
	RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeDigitalocean RecordingPauseResumeStopRecordingResponseDataStorageConfigType = "digitalocean"
	RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeSftp         RecordingPauseResumeStopRecordingResponseDataStorageConfigType = "sftp"
)

func (r RecordingPauseResumeStopRecordingResponseDataStorageConfigType) IsKnown() bool {
	switch r {
	case RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeGcs, RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeAws, RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeAzure, RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeDigitalocean, RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeSftp:
		return true
	}
	return false
}

type RecordingStartRecordingsResponse struct {
	// Success status of the operation
	Success bool `json:"success" api:"required"`
	// Data returned by the operation
	Data RecordingStartRecordingsResponseData `json:"data"`
	JSON recordingStartRecordingsResponseJSON `json:"-"`
}

// recordingStartRecordingsResponseJSON contains the JSON metadata for the struct
// [RecordingStartRecordingsResponse]
type recordingStartRecordingsResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingStartRecordingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingStartRecordingsResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type RecordingStartRecordingsResponseData struct {
	// ID of the recording
	ID string `json:"id" api:"required" format:"uuid"`
	// If the audio_config is passed, the URL for downloading the audio recording is
	// returned.
	AudioDownloadURL string `json:"audio_download_url" api:"required,nullable" format:"uri"`
	// URL where the recording can be downloaded.
	DownloadURL string `json:"download_url" api:"required,nullable" format:"uri"`
	// Timestamp when the download URL expires.
	DownloadURLExpiry time.Time `json:"download_url_expiry" api:"required,nullable" format:"date-time"`
	// File size of the recording, in bytes.
	FileSize float64 `json:"file_size" api:"required,nullable"`
	// Timestamp when this recording was invoked.
	InvokedTime time.Time `json:"invoked_time" api:"required" format:"date-time"`
	// File name of the recording.
	OutputFileName string `json:"output_file_name" api:"required"`
	// ID of the meeting session this recording is for.
	SessionID string `json:"session_id" api:"required,nullable" format:"uuid"`
	// Timestamp when this recording actually started after being invoked. Usually a
	// few seconds after `invoked_time`.
	StartedTime time.Time `json:"started_time" api:"required,nullable" format:"date-time"`
	// Current status of the recording.
	Status RecordingStartRecordingsResponseDataStatus `json:"status" api:"required"`
	// Timestamp when this recording was stopped. Optional; is present only when the
	// recording has actually been stopped.
	StoppedTime time.Time `json:"stopped_time" api:"required,nullable" format:"date-time"`
	// Total recording time in seconds.
	RecordingDuration int64                                             `json:"recording_duration"`
	StartReason       RecordingStartRecordingsResponseDataStartReason   `json:"start_reason"`
	StopReason        RecordingStartRecordingsResponseDataStopReason    `json:"stop_reason"`
	StorageConfig     RecordingStartRecordingsResponseDataStorageConfig `json:"storage_config" api:"nullable"`
	JSON              recordingStartRecordingsResponseDataJSON          `json:"-"`
}

// recordingStartRecordingsResponseDataJSON contains the JSON metadata for the
// struct [RecordingStartRecordingsResponseData]
type recordingStartRecordingsResponseDataJSON struct {
	ID                apijson.Field
	AudioDownloadURL  apijson.Field
	DownloadURL       apijson.Field
	DownloadURLExpiry apijson.Field
	FileSize          apijson.Field
	InvokedTime       apijson.Field
	OutputFileName    apijson.Field
	SessionID         apijson.Field
	StartedTime       apijson.Field
	Status            apijson.Field
	StoppedTime       apijson.Field
	RecordingDuration apijson.Field
	StartReason       apijson.Field
	StopReason        apijson.Field
	StorageConfig     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordingStartRecordingsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingStartRecordingsResponseDataJSON) RawJSON() string {
	return r.raw
}

// Current status of the recording.
type RecordingStartRecordingsResponseDataStatus string

const (
	RecordingStartRecordingsResponseDataStatusInvoked   RecordingStartRecordingsResponseDataStatus = "INVOKED"
	RecordingStartRecordingsResponseDataStatusRecording RecordingStartRecordingsResponseDataStatus = "RECORDING"
	RecordingStartRecordingsResponseDataStatusUploading RecordingStartRecordingsResponseDataStatus = "UPLOADING"
	RecordingStartRecordingsResponseDataStatusUploaded  RecordingStartRecordingsResponseDataStatus = "UPLOADED"
	RecordingStartRecordingsResponseDataStatusErrored   RecordingStartRecordingsResponseDataStatus = "ERRORED"
	RecordingStartRecordingsResponseDataStatusPaused    RecordingStartRecordingsResponseDataStatus = "PAUSED"
)

func (r RecordingStartRecordingsResponseDataStatus) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsResponseDataStatusInvoked, RecordingStartRecordingsResponseDataStatusRecording, RecordingStartRecordingsResponseDataStatusUploading, RecordingStartRecordingsResponseDataStatusUploaded, RecordingStartRecordingsResponseDataStatusErrored, RecordingStartRecordingsResponseDataStatusPaused:
		return true
	}
	return false
}

type RecordingStartRecordingsResponseDataStartReason struct {
	Caller RecordingStartRecordingsResponseDataStartReasonCaller `json:"caller"`
	// Specifies if the recording was started using the "Start a Recording"API or using
	// the parameter RECORD_ON_START in the "Create a meeting" API.
	//
	// If the recording is initiated using the "RECORD_ON_START" parameter, the user
	// details will not be populated.
	Reason RecordingStartRecordingsResponseDataStartReasonReason `json:"reason"`
	JSON   recordingStartRecordingsResponseDataStartReasonJSON   `json:"-"`
}

// recordingStartRecordingsResponseDataStartReasonJSON contains the JSON metadata
// for the struct [RecordingStartRecordingsResponseDataStartReason]
type recordingStartRecordingsResponseDataStartReasonJSON struct {
	Caller      apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingStartRecordingsResponseDataStartReason) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingStartRecordingsResponseDataStartReasonJSON) RawJSON() string {
	return r.raw
}

type RecordingStartRecordingsResponseDataStartReasonCaller struct {
	// Name of the user who started the recording.
	Name string `json:"name"`
	// The type can be an App or a user. If the type is `user`, then only the `user_Id`
	// and `name` are returned.
	Type RecordingStartRecordingsResponseDataStartReasonCallerType `json:"type"`
	// The user ID of the person who started the recording.
	UserID string                                                    `json:"user_Id" format:"uuid"`
	JSON   recordingStartRecordingsResponseDataStartReasonCallerJSON `json:"-"`
}

// recordingStartRecordingsResponseDataStartReasonCallerJSON contains the JSON
// metadata for the struct [RecordingStartRecordingsResponseDataStartReasonCaller]
type recordingStartRecordingsResponseDataStartReasonCallerJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingStartRecordingsResponseDataStartReasonCaller) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingStartRecordingsResponseDataStartReasonCallerJSON) RawJSON() string {
	return r.raw
}

// The type can be an App or a user. If the type is `user`, then only the `user_Id`
// and `name` are returned.
type RecordingStartRecordingsResponseDataStartReasonCallerType string

const (
	RecordingStartRecordingsResponseDataStartReasonCallerTypeOrganization RecordingStartRecordingsResponseDataStartReasonCallerType = "ORGANIZATION"
	RecordingStartRecordingsResponseDataStartReasonCallerTypeUser         RecordingStartRecordingsResponseDataStartReasonCallerType = "USER"
)

func (r RecordingStartRecordingsResponseDataStartReasonCallerType) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsResponseDataStartReasonCallerTypeOrganization, RecordingStartRecordingsResponseDataStartReasonCallerTypeUser:
		return true
	}
	return false
}

// Specifies if the recording was started using the "Start a Recording"API or using
// the parameter RECORD_ON_START in the "Create a meeting" API.
//
// If the recording is initiated using the "RECORD_ON_START" parameter, the user
// details will not be populated.
type RecordingStartRecordingsResponseDataStartReasonReason string

const (
	RecordingStartRecordingsResponseDataStartReasonReasonAPICall       RecordingStartRecordingsResponseDataStartReasonReason = "API_CALL"
	RecordingStartRecordingsResponseDataStartReasonReasonRecordOnStart RecordingStartRecordingsResponseDataStartReasonReason = "RECORD_ON_START"
)

func (r RecordingStartRecordingsResponseDataStartReasonReason) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsResponseDataStartReasonReasonAPICall, RecordingStartRecordingsResponseDataStartReasonReasonRecordOnStart:
		return true
	}
	return false
}

type RecordingStartRecordingsResponseDataStopReason struct {
	Caller RecordingStartRecordingsResponseDataStopReasonCaller `json:"caller"`
	// Specifies the reason why the recording stopped.
	Reason RecordingStartRecordingsResponseDataStopReasonReason `json:"reason"`
	JSON   recordingStartRecordingsResponseDataStopReasonJSON   `json:"-"`
}

// recordingStartRecordingsResponseDataStopReasonJSON contains the JSON metadata
// for the struct [RecordingStartRecordingsResponseDataStopReason]
type recordingStartRecordingsResponseDataStopReasonJSON struct {
	Caller      apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingStartRecordingsResponseDataStopReason) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingStartRecordingsResponseDataStopReasonJSON) RawJSON() string {
	return r.raw
}

type RecordingStartRecordingsResponseDataStopReasonCaller struct {
	// Name of the user who stopped the recording.
	Name string `json:"name"`
	// The type can be an App or a user. If the type is `user`, then only the `user_Id`
	// and `name` are returned.
	Type RecordingStartRecordingsResponseDataStopReasonCallerType `json:"type"`
	// The user ID of the person who stopped the recording.
	UserID string                                                   `json:"user_Id" format:"uuid"`
	JSON   recordingStartRecordingsResponseDataStopReasonCallerJSON `json:"-"`
}

// recordingStartRecordingsResponseDataStopReasonCallerJSON contains the JSON
// metadata for the struct [RecordingStartRecordingsResponseDataStopReasonCaller]
type recordingStartRecordingsResponseDataStopReasonCallerJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingStartRecordingsResponseDataStopReasonCaller) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingStartRecordingsResponseDataStopReasonCallerJSON) RawJSON() string {
	return r.raw
}

// The type can be an App or a user. If the type is `user`, then only the `user_Id`
// and `name` are returned.
type RecordingStartRecordingsResponseDataStopReasonCallerType string

const (
	RecordingStartRecordingsResponseDataStopReasonCallerTypeOrganization RecordingStartRecordingsResponseDataStopReasonCallerType = "ORGANIZATION"
	RecordingStartRecordingsResponseDataStopReasonCallerTypeUser         RecordingStartRecordingsResponseDataStopReasonCallerType = "USER"
)

func (r RecordingStartRecordingsResponseDataStopReasonCallerType) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsResponseDataStopReasonCallerTypeOrganization, RecordingStartRecordingsResponseDataStopReasonCallerTypeUser:
		return true
	}
	return false
}

// Specifies the reason why the recording stopped.
type RecordingStartRecordingsResponseDataStopReasonReason string

const (
	RecordingStartRecordingsResponseDataStopReasonReasonAPICall       RecordingStartRecordingsResponseDataStopReasonReason = "API_CALL"
	RecordingStartRecordingsResponseDataStopReasonReasonInternalError RecordingStartRecordingsResponseDataStopReasonReason = "INTERNAL_ERROR"
	RecordingStartRecordingsResponseDataStopReasonReasonAllPeersLeft  RecordingStartRecordingsResponseDataStopReasonReason = "ALL_PEERS_LEFT"
)

func (r RecordingStartRecordingsResponseDataStopReasonReason) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsResponseDataStopReasonReasonAPICall, RecordingStartRecordingsResponseDataStopReasonReasonInternalError, RecordingStartRecordingsResponseDataStopReasonReasonAllPeersLeft:
		return true
	}
	return false
}

type RecordingStartRecordingsResponseDataStorageConfig struct {
	// This field can have the runtime type of [string], [interface{}].
	AccessKey interface{} `json:"access_key"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingStartRecordingsResponseDataStorageConfigAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// This field can have the runtime type of [string], [interface{}].
	Region interface{}                                           `json:"region"`
	Type   RecordingStartRecordingsResponseDataStorageConfigType `json:"type"`
	// SSH destination server username for SFTP type storage medium
	Username string                                                `json:"username"`
	JSON     recordingStartRecordingsResponseDataStorageConfigJSON `json:"-"`
	union    RecordingStartRecordingsResponseDataStorageConfigUnion
}

// recordingStartRecordingsResponseDataStorageConfigJSON contains the JSON metadata
// for the struct [RecordingStartRecordingsResponseDataStorageConfig]
type recordingStartRecordingsResponseDataStorageConfigJSON struct {
	AccessKey   apijson.Field
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r recordingStartRecordingsResponseDataStorageConfigJSON) RawJSON() string {
	return r.raw
}

func (r *RecordingStartRecordingsResponseDataStorageConfig) UnmarshalJSON(data []byte) (err error) {
	*r = RecordingStartRecordingsResponseDataStorageConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecordingStartRecordingsResponseDataStorageConfigUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RecordingStartRecordingsResponseDataStorageConfigObject],
// [RecordingStartRecordingsResponseDataStorageConfigObject],
// [RecordingStartRecordingsResponseDataStorageConfigObject],
// [RecordingStartRecordingsResponseDataStorageConfigObject].
func (r RecordingStartRecordingsResponseDataStorageConfig) AsUnion() RecordingStartRecordingsResponseDataStorageConfigUnion {
	return r.union
}

// Union satisfied by [RecordingStartRecordingsResponseDataStorageConfigObject],
// [RecordingStartRecordingsResponseDataStorageConfigObject],
// [RecordingStartRecordingsResponseDataStorageConfigObject] or
// [RecordingStartRecordingsResponseDataStorageConfigObject].
type RecordingStartRecordingsResponseDataStorageConfigUnion interface {
	implementsRecordingStartRecordingsResponseDataStorageConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordingStartRecordingsResponseDataStorageConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingStartRecordingsResponseDataStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingStartRecordingsResponseDataStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingStartRecordingsResponseDataStorageConfigObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordingStartRecordingsResponseDataStorageConfigObject{}),
		},
	)
}

type RecordingStartRecordingsResponseDataStorageConfigObject struct {
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingStartRecordingsResponseDataStorageConfigObjectAuthMethod `json:"auth_method"`
	// Name of the storage medium's bucket.
	Bucket string `json:"bucket"`
	// SSH destination server host for SFTP type storage medium
	Host string `json:"host"`
	// Path relative to the bucket root at which the recording will be placed.
	Path string `json:"path"`
	// SSH destination server port for SFTP type storage medium
	Port float64 `json:"port"`
	// Region of the storage medium.
	Region string                                                      `json:"region"`
	Type   RecordingStartRecordingsResponseDataStorageConfigObjectType `json:"type"`
	// SSH destination server username for SFTP type storage medium
	Username string                                                      `json:"username"`
	JSON     recordingStartRecordingsResponseDataStorageConfigObjectJSON `json:"-"`
}

// recordingStartRecordingsResponseDataStorageConfigObjectJSON contains the JSON
// metadata for the struct
// [RecordingStartRecordingsResponseDataStorageConfigObject]
type recordingStartRecordingsResponseDataStorageConfigObjectJSON struct {
	AuthMethod  apijson.Field
	Bucket      apijson.Field
	Host        apijson.Field
	Path        apijson.Field
	Port        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingStartRecordingsResponseDataStorageConfigObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingStartRecordingsResponseDataStorageConfigObjectJSON) RawJSON() string {
	return r.raw
}

func (r RecordingStartRecordingsResponseDataStorageConfigObject) implementsRecordingStartRecordingsResponseDataStorageConfig() {
}

// Authentication method used for "sftp" type storage medium
type RecordingStartRecordingsResponseDataStorageConfigObjectAuthMethod string

const (
	RecordingStartRecordingsResponseDataStorageConfigObjectAuthMethodKey      RecordingStartRecordingsResponseDataStorageConfigObjectAuthMethod = "KEY"
	RecordingStartRecordingsResponseDataStorageConfigObjectAuthMethodPassword RecordingStartRecordingsResponseDataStorageConfigObjectAuthMethod = "PASSWORD"
)

func (r RecordingStartRecordingsResponseDataStorageConfigObjectAuthMethod) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsResponseDataStorageConfigObjectAuthMethodKey, RecordingStartRecordingsResponseDataStorageConfigObjectAuthMethodPassword:
		return true
	}
	return false
}

type RecordingStartRecordingsResponseDataStorageConfigObjectType string

const (
	RecordingStartRecordingsResponseDataStorageConfigObjectTypeGcs RecordingStartRecordingsResponseDataStorageConfigObjectType = "gcs"
)

func (r RecordingStartRecordingsResponseDataStorageConfigObjectType) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsResponseDataStorageConfigObjectTypeGcs:
		return true
	}
	return false
}

// Authentication method used for "sftp" type storage medium
type RecordingStartRecordingsResponseDataStorageConfigAuthMethod string

const (
	RecordingStartRecordingsResponseDataStorageConfigAuthMethodKey      RecordingStartRecordingsResponseDataStorageConfigAuthMethod = "KEY"
	RecordingStartRecordingsResponseDataStorageConfigAuthMethodPassword RecordingStartRecordingsResponseDataStorageConfigAuthMethod = "PASSWORD"
)

func (r RecordingStartRecordingsResponseDataStorageConfigAuthMethod) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsResponseDataStorageConfigAuthMethodKey, RecordingStartRecordingsResponseDataStorageConfigAuthMethodPassword:
		return true
	}
	return false
}

type RecordingStartRecordingsResponseDataStorageConfigType string

const (
	RecordingStartRecordingsResponseDataStorageConfigTypeGcs          RecordingStartRecordingsResponseDataStorageConfigType = "gcs"
	RecordingStartRecordingsResponseDataStorageConfigTypeAws          RecordingStartRecordingsResponseDataStorageConfigType = "aws"
	RecordingStartRecordingsResponseDataStorageConfigTypeAzure        RecordingStartRecordingsResponseDataStorageConfigType = "azure"
	RecordingStartRecordingsResponseDataStorageConfigTypeDigitalocean RecordingStartRecordingsResponseDataStorageConfigType = "digitalocean"
	RecordingStartRecordingsResponseDataStorageConfigTypeSftp         RecordingStartRecordingsResponseDataStorageConfigType = "sftp"
)

func (r RecordingStartRecordingsResponseDataStorageConfigType) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsResponseDataStorageConfigTypeGcs, RecordingStartRecordingsResponseDataStorageConfigTypeAws, RecordingStartRecordingsResponseDataStorageConfigTypeAzure, RecordingStartRecordingsResponseDataStorageConfigTypeDigitalocean, RecordingStartRecordingsResponseDataStorageConfigTypeSftp:
		return true
	}
	return false
}

type RecordingStartTrackRecordingResponse struct {
	// Success status of the operation
	Success bool `json:"success" api:"required"`
	// Data returned by the operation
	Data RecordingStartTrackRecordingResponseData `json:"data"`
	JSON recordingStartTrackRecordingResponseJSON `json:"-"`
}

// recordingStartTrackRecordingResponseJSON contains the JSON metadata for the
// struct [RecordingStartTrackRecordingResponse]
type recordingStartTrackRecordingResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingStartTrackRecordingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingStartTrackRecordingResponseJSON) RawJSON() string {
	return r.raw
}

// Data returned by the operation
type RecordingStartTrackRecordingResponseData struct {
	Recording RecordingStartTrackRecordingResponseDataRecording `json:"recording" api:"required"`
	JSON      recordingStartTrackRecordingResponseDataJSON      `json:"-"`
}

// recordingStartTrackRecordingResponseDataJSON contains the JSON metadata for the
// struct [RecordingStartTrackRecordingResponseData]
type recordingStartTrackRecordingResponseDataJSON struct {
	Recording   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordingStartTrackRecordingResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingStartTrackRecordingResponseDataJSON) RawJSON() string {
	return r.raw
}

type RecordingStartTrackRecordingResponseDataRecording struct {
	// ID of the recording
	ID string `json:"id" api:"required" format:"uuid"`
	// If the audio_config is passed, the URL for downloading the audio recording is
	// returned.
	AudioDownloadURL string `json:"audio_download_url" api:"required,nullable" format:"uri"`
	// URL where the recording can be downloaded.
	DownloadURL string `json:"download_url" api:"required,nullable" format:"uri"`
	// Timestamp when the download URL expires.
	DownloadURLExpiry time.Time `json:"download_url_expiry" api:"required,nullable" format:"date-time"`
	// File size of the recording, in bytes.
	FileSize float64 `json:"file_size" api:"required,nullable"`
	// Timestamp when this recording was invoked.
	InvokedTime time.Time `json:"invoked_time" api:"required" format:"date-time"`
	// File name of the recording.
	OutputFileName string `json:"output_file_name" api:"required"`
	// ID of the meeting session this recording is for.
	SessionID string `json:"session_id" api:"required,nullable" format:"uuid"`
	// Timestamp when this recording actually started after being invoked. Usually a
	// few seconds after `invoked_time`.
	StartedTime time.Time `json:"started_time" api:"required,nullable" format:"date-time"`
	// Current status of the recording.
	Status RecordingStartTrackRecordingResponseDataRecordingStatus `json:"status" api:"required"`
	// Timestamp when this recording was stopped. Optional; is present only when the
	// recording has actually been stopped.
	StoppedTime time.Time `json:"stopped_time" api:"required,nullable" format:"date-time"`
	// Total recording time in seconds.
	RecordingDuration int64                                                 `json:"recording_duration"`
	JSON              recordingStartTrackRecordingResponseDataRecordingJSON `json:"-"`
}

// recordingStartTrackRecordingResponseDataRecordingJSON contains the JSON metadata
// for the struct [RecordingStartTrackRecordingResponseDataRecording]
type recordingStartTrackRecordingResponseDataRecordingJSON struct {
	ID                apijson.Field
	AudioDownloadURL  apijson.Field
	DownloadURL       apijson.Field
	DownloadURLExpiry apijson.Field
	FileSize          apijson.Field
	InvokedTime       apijson.Field
	OutputFileName    apijson.Field
	SessionID         apijson.Field
	StartedTime       apijson.Field
	Status            apijson.Field
	StoppedTime       apijson.Field
	RecordingDuration apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordingStartTrackRecordingResponseDataRecording) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingStartTrackRecordingResponseDataRecordingJSON) RawJSON() string {
	return r.raw
}

// Current status of the recording.
type RecordingStartTrackRecordingResponseDataRecordingStatus string

const (
	RecordingStartTrackRecordingResponseDataRecordingStatusInvoked   RecordingStartTrackRecordingResponseDataRecordingStatus = "INVOKED"
	RecordingStartTrackRecordingResponseDataRecordingStatusRecording RecordingStartTrackRecordingResponseDataRecordingStatus = "RECORDING"
	RecordingStartTrackRecordingResponseDataRecordingStatusUploading RecordingStartTrackRecordingResponseDataRecordingStatus = "UPLOADING"
	RecordingStartTrackRecordingResponseDataRecordingStatusUploaded  RecordingStartTrackRecordingResponseDataRecordingStatus = "UPLOADED"
	RecordingStartTrackRecordingResponseDataRecordingStatusErrored   RecordingStartTrackRecordingResponseDataRecordingStatus = "ERRORED"
	RecordingStartTrackRecordingResponseDataRecordingStatusPaused    RecordingStartTrackRecordingResponseDataRecordingStatus = "PAUSED"
)

func (r RecordingStartTrackRecordingResponseDataRecordingStatus) IsKnown() bool {
	switch r {
	case RecordingStartTrackRecordingResponseDataRecordingStatusInvoked, RecordingStartTrackRecordingResponseDataRecordingStatusRecording, RecordingStartTrackRecordingResponseDataRecordingStatusUploading, RecordingStartTrackRecordingResponseDataRecordingStatusUploaded, RecordingStartTrackRecordingResponseDataRecordingStatusErrored, RecordingStartTrackRecordingResponseDataRecordingStatusPaused:
		return true
	}
	return false
}

type RecordingGetActiveRecordingsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RecordingGetOneRecordingParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RecordingGetRecordingsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The end time range for which you want to retrieve the meetings. The time must be
	// specified in ISO format.
	EndTime param.Field[time.Time] `query:"end_time" format:"date-time"`
	// If passed, only shows expired/non-expired recordings on RealtimeKit's bucket
	Expired param.Field[bool] `query:"expired"`
	// ID of a meeting. Optional. Will limit results to only this meeting if passed.
	MeetingID param.Field[string] `query:"meeting_id" format:"uuid"`
	// The page number from which you want your page search results to be displayed.
	PageNo param.Field[float64] `query:"page_no"`
	// Number of results per page
	PerPage param.Field[float64] `query:"per_page"`
	// The search query string. You can search using the meeting ID or title.
	Search    param.Field[string]                                `query:"search"`
	SortBy    param.Field[RecordingGetRecordingsParamsSortBy]    `query:"sort_by"`
	SortOrder param.Field[RecordingGetRecordingsParamsSortOrder] `query:"sort_order"`
	// The start time range for which you want to retrieve the meetings. The time must
	// be specified in ISO format.
	StartTime param.Field[time.Time] `query:"start_time" format:"date-time"`
	// Filter by one or more recording status
	Status param.Field[[]RecordingGetRecordingsParamsStatus] `query:"status"`
}

// URLQuery serializes [RecordingGetRecordingsParams]'s query parameters as
// `url.Values`.
func (r RecordingGetRecordingsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RecordingGetRecordingsParamsSortBy string

const (
	RecordingGetRecordingsParamsSortByInvokedTime RecordingGetRecordingsParamsSortBy = "invokedTime"
)

func (r RecordingGetRecordingsParamsSortBy) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsParamsSortByInvokedTime:
		return true
	}
	return false
}

type RecordingGetRecordingsParamsSortOrder string

const (
	RecordingGetRecordingsParamsSortOrderAsc  RecordingGetRecordingsParamsSortOrder = "ASC"
	RecordingGetRecordingsParamsSortOrderDesc RecordingGetRecordingsParamsSortOrder = "DESC"
)

func (r RecordingGetRecordingsParamsSortOrder) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsParamsSortOrderAsc, RecordingGetRecordingsParamsSortOrderDesc:
		return true
	}
	return false
}

type RecordingGetRecordingsParamsStatus string

const (
	RecordingGetRecordingsParamsStatusInvoked   RecordingGetRecordingsParamsStatus = "INVOKED"
	RecordingGetRecordingsParamsStatusRecording RecordingGetRecordingsParamsStatus = "RECORDING"
	RecordingGetRecordingsParamsStatusUploading RecordingGetRecordingsParamsStatus = "UPLOADING"
	RecordingGetRecordingsParamsStatusUploaded  RecordingGetRecordingsParamsStatus = "UPLOADED"
	RecordingGetRecordingsParamsStatusErrored   RecordingGetRecordingsParamsStatus = "ERRORED"
	RecordingGetRecordingsParamsStatusPaused    RecordingGetRecordingsParamsStatus = "PAUSED"
)

func (r RecordingGetRecordingsParamsStatus) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsParamsStatusInvoked, RecordingGetRecordingsParamsStatusRecording, RecordingGetRecordingsParamsStatusUploading, RecordingGetRecordingsParamsStatusUploaded, RecordingGetRecordingsParamsStatusErrored, RecordingGetRecordingsParamsStatusPaused:
		return true
	}
	return false
}

type RecordingPauseResumeStopRecordingParams struct {
	// The account identifier tag.
	AccountID param.Field[string]                                        `path:"account_id" api:"required"`
	Action    param.Field[RecordingPauseResumeStopRecordingParamsAction] `json:"action" api:"required"`
}

func (r RecordingPauseResumeStopRecordingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RecordingPauseResumeStopRecordingParamsAction string

const (
	RecordingPauseResumeStopRecordingParamsActionStop   RecordingPauseResumeStopRecordingParamsAction = "stop"
	RecordingPauseResumeStopRecordingParamsActionPause  RecordingPauseResumeStopRecordingParamsAction = "pause"
	RecordingPauseResumeStopRecordingParamsActionResume RecordingPauseResumeStopRecordingParamsAction = "resume"
)

func (r RecordingPauseResumeStopRecordingParamsAction) IsKnown() bool {
	switch r {
	case RecordingPauseResumeStopRecordingParamsActionStop, RecordingPauseResumeStopRecordingParamsActionPause, RecordingPauseResumeStopRecordingParamsActionResume:
		return true
	}
	return false
}

type RecordingStartRecordingsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// ID of the meeting to record.
	MeetingID param.Field[string] `json:"meeting_id" api:"required" format:"uuid"`
	// By default, a meeting allows only one recording to run at a time. Enabling the
	// `allow_multiple_recordings` parameter to true allows you to initiate multiple
	// recordings concurrently in the same meeting. This allows you to record separate
	// videos of the same meeting with different configurations, such as portrait mode
	// or landscape mode.
	AllowMultipleRecordings param.Field[bool] `json:"allow_multiple_recordings"`
	// Object containing configuration regarding the audio that is being recorded.
	AudioConfig param.Field[RecordingStartRecordingsParamsAudioConfig] `json:"audio_config"`
	// Update the recording file name.
	FileNamePrefix param.Field[string] `json:"file_name_prefix"`
	// Allows you to add timed metadata to your recordings, which are digital markers
	// inserted into a video file to provide contextual information at specific points
	// in the content range. The ID3 tags containing this information are available to
	// clients on the playback timeline in HLS format. The output files are generated
	// in a compressed .tar format.
	InteractiveConfig param.Field[RecordingStartRecordingsParamsInteractiveConfig] `json:"interactive_config"`
	// Specifies the maximum duration for recording in seconds, ranging from a minimum
	// of 60 seconds to a maximum of 24 hours.
	MaxSeconds              param.Field[int64]                                                 `json:"max_seconds"`
	RealtimekitBucketConfig param.Field[RecordingStartRecordingsParamsRealtimekitBucketConfig] `json:"realtimekit_bucket_config"`
	RtmpOutConfig           param.Field[RecordingStartRecordingsParamsRtmpOutConfig]           `json:"rtmp_out_config"`
	StorageConfig           param.Field[RecordingStartRecordingsParamsStorageConfigUnion]      `json:"storage_config"`
	// Pass a custom url to record arbitary screen
	URL         param.Field[string]                                    `json:"url" format:"uri"`
	VideoConfig param.Field[RecordingStartRecordingsParamsVideoConfig] `json:"video_config"`
}

func (r RecordingStartRecordingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object containing configuration regarding the audio that is being recorded.
type RecordingStartRecordingsParamsAudioConfig struct {
	// Audio signal pathway within an audio file that carries a specific sound source.
	Channel param.Field[RecordingStartRecordingsParamsAudioConfigChannel] `json:"channel"`
	// Codec using which the recording will be encoded. If VP8/VP9 is selected for
	// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
	// audioConfig is automatically set to vorbis.
	Codec param.Field[RecordingStartRecordingsParamsAudioConfigCodec] `json:"codec"`
	// Controls whether to export audio file seperately
	ExportFile param.Field[bool] `json:"export_file"`
}

func (r RecordingStartRecordingsParamsAudioConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Audio signal pathway within an audio file that carries a specific sound source.
type RecordingStartRecordingsParamsAudioConfigChannel string

const (
	RecordingStartRecordingsParamsAudioConfigChannelMono   RecordingStartRecordingsParamsAudioConfigChannel = "mono"
	RecordingStartRecordingsParamsAudioConfigChannelStereo RecordingStartRecordingsParamsAudioConfigChannel = "stereo"
)

func (r RecordingStartRecordingsParamsAudioConfigChannel) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsParamsAudioConfigChannelMono, RecordingStartRecordingsParamsAudioConfigChannelStereo:
		return true
	}
	return false
}

// Codec using which the recording will be encoded. If VP8/VP9 is selected for
// videoConfig, changing audioConfig is not allowed. In this case, the codec in the
// audioConfig is automatically set to vorbis.
type RecordingStartRecordingsParamsAudioConfigCodec string

const (
	RecordingStartRecordingsParamsAudioConfigCodecMP3 RecordingStartRecordingsParamsAudioConfigCodec = "MP3"
	RecordingStartRecordingsParamsAudioConfigCodecAac RecordingStartRecordingsParamsAudioConfigCodec = "AAC"
)

func (r RecordingStartRecordingsParamsAudioConfigCodec) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsParamsAudioConfigCodecMP3, RecordingStartRecordingsParamsAudioConfigCodecAac:
		return true
	}
	return false
}

// Allows you to add timed metadata to your recordings, which are digital markers
// inserted into a video file to provide contextual information at specific points
// in the content range. The ID3 tags containing this information are available to
// clients on the playback timeline in HLS format. The output files are generated
// in a compressed .tar format.
type RecordingStartRecordingsParamsInteractiveConfig struct {
	// The metadata is presented in the form of ID3 tags.
	Type param.Field[RecordingStartRecordingsParamsInteractiveConfigType] `json:"type"`
}

func (r RecordingStartRecordingsParamsInteractiveConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The metadata is presented in the form of ID3 tags.
type RecordingStartRecordingsParamsInteractiveConfigType string

const (
	RecordingStartRecordingsParamsInteractiveConfigTypeId3 RecordingStartRecordingsParamsInteractiveConfigType = "ID3"
)

func (r RecordingStartRecordingsParamsInteractiveConfigType) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsParamsInteractiveConfigTypeId3:
		return true
	}
	return false
}

type RecordingStartRecordingsParamsRealtimekitBucketConfig struct {
	// Controls whether recordings are uploaded to RealtimeKit's bucket. If set to
	// false, `download_url`, `audio_download_url`, `download_url_expiry` won't be
	// generated for a recording.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r RecordingStartRecordingsParamsRealtimekitBucketConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RecordingStartRecordingsParamsRtmpOutConfig struct {
	// RTMP URL to stream to
	RtmpURL param.Field[string] `json:"rtmp_url" format:"uri"`
}

func (r RecordingStartRecordingsParamsRtmpOutConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RecordingStartRecordingsParamsStorageConfig struct {
	AccessKey param.Field[interface{}] `json:"access_key"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod param.Field[RecordingStartRecordingsParamsStorageConfigAuthMethod] `json:"auth_method"`
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
	PrivateKey param.Field[string]      `json:"private_key"`
	Region     param.Field[interface{}] `json:"region"`
	// Secret key of the storage medium. Similar to `access_key`, it is only writeable
	// by clients, not readable.
	Secret param.Field[string]                                          `json:"secret"`
	Type   param.Field[RecordingStartRecordingsParamsStorageConfigType] `json:"type"`
	// SSH destination server username for SFTP type storage medium
	Username param.Field[string] `json:"username"`
}

func (r RecordingStartRecordingsParamsStorageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordingStartRecordingsParamsStorageConfig) implementsRecordingStartRecordingsParamsStorageConfigUnion() {
}

// Satisfied by [realtime_kit.RecordingStartRecordingsParamsStorageConfigObject],
// [realtime_kit.RecordingStartRecordingsParamsStorageConfigObject],
// [realtime_kit.RecordingStartRecordingsParamsStorageConfigObject],
// [realtime_kit.RecordingStartRecordingsParamsStorageConfigObject],
// [RecordingStartRecordingsParamsStorageConfig].
type RecordingStartRecordingsParamsStorageConfigUnion interface {
	implementsRecordingStartRecordingsParamsStorageConfigUnion()
}

type RecordingStartRecordingsParamsStorageConfigObject struct {
	// Access key of the storage medium. Access key is not required for the `gcs`
	// storage media type.
	//
	// Note that this field is not readable by clients, only writeable.
	AccessKey param.Field[string] `json:"access_key"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod param.Field[RecordingStartRecordingsParamsStorageConfigObjectAuthMethod] `json:"auth_method"`
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
	Secret param.Field[string]                                                `json:"secret"`
	Type   param.Field[RecordingStartRecordingsParamsStorageConfigObjectType] `json:"type"`
	// SSH destination server username for SFTP type storage medium
	Username param.Field[string] `json:"username"`
}

func (r RecordingStartRecordingsParamsStorageConfigObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordingStartRecordingsParamsStorageConfigObject) implementsRecordingStartRecordingsParamsStorageConfigUnion() {
}

// Authentication method used for "sftp" type storage medium
type RecordingStartRecordingsParamsStorageConfigObjectAuthMethod string

const (
	RecordingStartRecordingsParamsStorageConfigObjectAuthMethodKey      RecordingStartRecordingsParamsStorageConfigObjectAuthMethod = "KEY"
	RecordingStartRecordingsParamsStorageConfigObjectAuthMethodPassword RecordingStartRecordingsParamsStorageConfigObjectAuthMethod = "PASSWORD"
)

func (r RecordingStartRecordingsParamsStorageConfigObjectAuthMethod) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsParamsStorageConfigObjectAuthMethodKey, RecordingStartRecordingsParamsStorageConfigObjectAuthMethodPassword:
		return true
	}
	return false
}

type RecordingStartRecordingsParamsStorageConfigObjectType string

const (
	RecordingStartRecordingsParamsStorageConfigObjectTypeGcs RecordingStartRecordingsParamsStorageConfigObjectType = "gcs"
)

func (r RecordingStartRecordingsParamsStorageConfigObjectType) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsParamsStorageConfigObjectTypeGcs:
		return true
	}
	return false
}

// Authentication method used for "sftp" type storage medium
type RecordingStartRecordingsParamsStorageConfigAuthMethod string

const (
	RecordingStartRecordingsParamsStorageConfigAuthMethodKey      RecordingStartRecordingsParamsStorageConfigAuthMethod = "KEY"
	RecordingStartRecordingsParamsStorageConfigAuthMethodPassword RecordingStartRecordingsParamsStorageConfigAuthMethod = "PASSWORD"
)

func (r RecordingStartRecordingsParamsStorageConfigAuthMethod) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsParamsStorageConfigAuthMethodKey, RecordingStartRecordingsParamsStorageConfigAuthMethodPassword:
		return true
	}
	return false
}

type RecordingStartRecordingsParamsStorageConfigType string

const (
	RecordingStartRecordingsParamsStorageConfigTypeGcs          RecordingStartRecordingsParamsStorageConfigType = "gcs"
	RecordingStartRecordingsParamsStorageConfigTypeAws          RecordingStartRecordingsParamsStorageConfigType = "aws"
	RecordingStartRecordingsParamsStorageConfigTypeAzure        RecordingStartRecordingsParamsStorageConfigType = "azure"
	RecordingStartRecordingsParamsStorageConfigTypeDigitalocean RecordingStartRecordingsParamsStorageConfigType = "digitalocean"
	RecordingStartRecordingsParamsStorageConfigTypeSftp         RecordingStartRecordingsParamsStorageConfigType = "sftp"
)

func (r RecordingStartRecordingsParamsStorageConfigType) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsParamsStorageConfigTypeGcs, RecordingStartRecordingsParamsStorageConfigTypeAws, RecordingStartRecordingsParamsStorageConfigTypeAzure, RecordingStartRecordingsParamsStorageConfigTypeDigitalocean, RecordingStartRecordingsParamsStorageConfigTypeSftp:
		return true
	}
	return false
}

type RecordingStartRecordingsParamsVideoConfig struct {
	// Codec using which the recording will be encoded.
	Codec param.Field[RecordingStartRecordingsParamsVideoConfigCodec] `json:"codec"`
	// Controls whether to export video file seperately
	ExportFile param.Field[bool] `json:"export_file"`
	// Height of the recording video in pixels
	Height param.Field[int64] `json:"height"`
	// Watermark to be added to the recording
	Watermark param.Field[RecordingStartRecordingsParamsVideoConfigWatermark] `json:"watermark"`
	// Width of the recording video in pixels
	Width param.Field[int64] `json:"width"`
}

func (r RecordingStartRecordingsParamsVideoConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Codec using which the recording will be encoded.
type RecordingStartRecordingsParamsVideoConfigCodec string

const (
	RecordingStartRecordingsParamsVideoConfigCodecH264 RecordingStartRecordingsParamsVideoConfigCodec = "H264"
	RecordingStartRecordingsParamsVideoConfigCodecVp8  RecordingStartRecordingsParamsVideoConfigCodec = "VP8"
	RecordingStartRecordingsParamsVideoConfigCodecVp9  RecordingStartRecordingsParamsVideoConfigCodec = "VP9"
)

func (r RecordingStartRecordingsParamsVideoConfigCodec) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsParamsVideoConfigCodecH264, RecordingStartRecordingsParamsVideoConfigCodecVp8, RecordingStartRecordingsParamsVideoConfigCodecVp9:
		return true
	}
	return false
}

// Watermark to be added to the recording
type RecordingStartRecordingsParamsVideoConfigWatermark struct {
	// Position of the watermark
	Position param.Field[RecordingStartRecordingsParamsVideoConfigWatermarkPosition] `json:"position"`
	// Size of the watermark
	Size param.Field[RecordingStartRecordingsParamsVideoConfigWatermarkSize] `json:"size"`
	// URL of the watermark image
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r RecordingStartRecordingsParamsVideoConfigWatermark) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Position of the watermark
type RecordingStartRecordingsParamsVideoConfigWatermarkPosition string

const (
	RecordingStartRecordingsParamsVideoConfigWatermarkPositionLeftTop     RecordingStartRecordingsParamsVideoConfigWatermarkPosition = "left top"
	RecordingStartRecordingsParamsVideoConfigWatermarkPositionRightTop    RecordingStartRecordingsParamsVideoConfigWatermarkPosition = "right top"
	RecordingStartRecordingsParamsVideoConfigWatermarkPositionLeftBottom  RecordingStartRecordingsParamsVideoConfigWatermarkPosition = "left bottom"
	RecordingStartRecordingsParamsVideoConfigWatermarkPositionRightBottom RecordingStartRecordingsParamsVideoConfigWatermarkPosition = "right bottom"
)

func (r RecordingStartRecordingsParamsVideoConfigWatermarkPosition) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsParamsVideoConfigWatermarkPositionLeftTop, RecordingStartRecordingsParamsVideoConfigWatermarkPositionRightTop, RecordingStartRecordingsParamsVideoConfigWatermarkPositionLeftBottom, RecordingStartRecordingsParamsVideoConfigWatermarkPositionRightBottom:
		return true
	}
	return false
}

// Size of the watermark
type RecordingStartRecordingsParamsVideoConfigWatermarkSize struct {
	// Height of the watermark in px
	Height param.Field[int64] `json:"height"`
	// Width of the watermark in px
	Width param.Field[int64] `json:"width"`
}

func (r RecordingStartRecordingsParamsVideoConfigWatermarkSize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RecordingStartTrackRecordingParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// ID of the meeting to record.
	MeetingID param.Field[string] `json:"meeting_id" api:"required" format:"uuid"`
	// Optional audio layer configuration. If omitted, RealtimeKit records all
	// participant audio using the default file name prefix.
	Layers param.Field[map[string]RecordingStartTrackRecordingParamsLayers] `json:"layers"`
	// Optional list of participant user IDs to record. Selective track recording
	// (`user_ids`) is in early beta contact support to use this feature.
	UserIDs param.Field[[]string] `json:"user_ids"`
}

func (r RecordingStartTrackRecordingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RecordingStartTrackRecordingParamsLayers struct {
	// A file name prefix to apply for files generated from this layer
	FileNamePrefix param.Field[string] `json:"file_name_prefix"`
	// Media kind to record. Track recording currently supports audio only.
	MediaKind param.Field[RecordingStartTrackRecordingParamsLayersMediaKind] `json:"media_kind"`
}

func (r RecordingStartTrackRecordingParamsLayers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Media kind to record. Track recording currently supports audio only.
type RecordingStartTrackRecordingParamsLayersMediaKind string

const (
	RecordingStartTrackRecordingParamsLayersMediaKindAudio RecordingStartTrackRecordingParamsLayersMediaKind = "audio"
)

func (r RecordingStartTrackRecordingParamsLayersMediaKind) IsKnown() bool {
	switch r {
	case RecordingStartTrackRecordingParamsLayersMediaKindAudio:
		return true
	}
	return false
}
