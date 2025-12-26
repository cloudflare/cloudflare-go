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
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/recordings/active-recording/%s", query.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns details of a recording for the given recording ID.
func (r *RecordingService) GetOneRecording(ctx context.Context, appID string, recordingID string, query RecordingGetOneRecordingParams, opts ...option.RequestOption) (res *RecordingGetOneRecordingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if recordingID == "" {
		err = errors.New("missing required recording_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/recordings/%s", query.AccountID, appID, recordingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all recordings for an App. If the `meeting_id` parameter is passed,
// returns all recordings for the given meeting ID.
func (r *RecordingService) GetRecordings(ctx context.Context, appID string, params RecordingGetRecordingsParams, opts ...option.RequestOption) (res *RecordingGetRecordingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/recordings", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Pause/Resume/Stop a given recording ID.
func (r *RecordingService) PauseResumeStopRecording(ctx context.Context, appID string, recordingID string, params RecordingPauseResumeStopRecordingParams, opts ...option.RequestOption) (res *RecordingPauseResumeStopRecordingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if recordingID == "" {
		err = errors.New("missing required recording_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/recordings/%s", params.AccountID, appID, recordingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Starts recording a meeting. The meeting can be started by an App admin directly,
// or a participant with permissions to start a recording, based on the type of
// authorization used.
func (r *RecordingService) StartRecordings(ctx context.Context, appID string, params RecordingStartRecordingsParams, opts ...option.RequestOption) (res *RecordingStartRecordingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/recordings", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Starts a track recording in a meeting. Track recordings consist of "layers".
// Layers are used to map audio/video tracks in a meeting to output destinations.
// More information about track recordings is available in the
// [Track Recordings Guide Page](https://docs.realtime.cloudflare.com/guides/capabilities/recording/recording-overview).
func (r *RecordingService) StartTrackRecording(ctx context.Context, appID string, params RecordingStartTrackRecordingParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/recordings/track", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

type RecordingGetActiveRecordingsResponse struct {
	// Data returned by the operation
	Data RecordingGetActiveRecordingsResponseData `json:"data,required"`
	// Success status of the operation
	Success bool                                     `json:"success,required"`
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
	ID string `json:"id,required" format:"uuid"`
	// If the audio_config is passed, the URL for downloading the audio recording is
	// returned.
	AudioDownloadURL string `json:"audio_download_url,required,nullable" format:"uri"`
	// URL where the recording can be downloaded.
	DownloadURL string `json:"download_url,required,nullable" format:"uri"`
	// Timestamp when the download URL expires.
	DownloadURLExpiry time.Time `json:"download_url_expiry,required,nullable" format:"date-time"`
	// File size of the recording, in bytes.
	FileSize float64 `json:"file_size,required,nullable"`
	// Timestamp when this recording was invoked.
	InvokedTime time.Time `json:"invoked_time,required" format:"date-time"`
	// File name of the recording.
	OutputFileName string `json:"output_file_name,required"`
	// ID of the meeting session this recording is for.
	SessionID string `json:"session_id,required,nullable" format:"uuid"`
	// Timestamp when this recording actually started after being invoked. Usually a
	// few seconds after `invoked_time`.
	StartedTime time.Time `json:"started_time,required,nullable" format:"date-time"`
	// Current status of the recording.
	Status RecordingGetActiveRecordingsResponseDataStatus `json:"status,required"`
	// Timestamp when this recording was stopped. Optional; is present only when the
	// recording has actually been stopped.
	StoppedTime time.Time `json:"stopped_time,required,nullable" format:"date-time"`
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
	Success bool `json:"success,required"`
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
	ID string `json:"id,required" format:"uuid"`
	// If the audio_config is passed, the URL for downloading the audio recording is
	// returned.
	AudioDownloadURL string `json:"audio_download_url,required,nullable" format:"uri"`
	// URL where the recording can be downloaded.
	DownloadURL string `json:"download_url,required,nullable" format:"uri"`
	// Timestamp when the download URL expires.
	DownloadURLExpiry time.Time `json:"download_url_expiry,required,nullable" format:"date-time"`
	// File size of the recording, in bytes.
	FileSize float64 `json:"file_size,required,nullable"`
	// Timestamp when this recording was invoked.
	InvokedTime time.Time `json:"invoked_time,required" format:"date-time"`
	// File name of the recording.
	OutputFileName string `json:"output_file_name,required"`
	// ID of the meeting session this recording is for.
	SessionID string `json:"session_id,required,nullable" format:"uuid"`
	// Timestamp when this recording actually started after being invoked. Usually a
	// few seconds after `invoked_time`.
	StartedTime time.Time `json:"started_time,required,nullable" format:"date-time"`
	// Current status of the recording.
	Status RecordingGetOneRecordingResponseDataStatus `json:"status,required"`
	// Timestamp when this recording was stopped. Optional; is present only when the
	// recording has actually been stopped.
	StoppedTime time.Time `json:"stopped_time,required,nullable" format:"date-time"`
	// Total recording time in seconds.
	RecordingDuration int64                                             `json:"recording_duration"`
	StartReason       RecordingGetOneRecordingResponseDataStartReason   `json:"start_reason"`
	StopReason        RecordingGetOneRecordingResponseDataStopReason    `json:"stop_reason"`
	StorageConfig     RecordingGetOneRecordingResponseDataStorageConfig `json:"storage_config,nullable"`
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
	// Type of storage media.
	Type RecordingGetOneRecordingResponseDataStorageConfigType `json:"type,required"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingGetOneRecordingResponseDataStorageConfigAuthMethod `json:"auth_method"`
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
	Username string                                                `json:"username"`
	JSON     recordingGetOneRecordingResponseDataStorageConfigJSON `json:"-"`
}

// recordingGetOneRecordingResponseDataStorageConfigJSON contains the JSON metadata
// for the struct [RecordingGetOneRecordingResponseDataStorageConfig]
type recordingGetOneRecordingResponseDataStorageConfigJSON struct {
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

func (r *RecordingGetOneRecordingResponseDataStorageConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetOneRecordingResponseDataStorageConfigJSON) RawJSON() string {
	return r.raw
}

// Type of storage media.
type RecordingGetOneRecordingResponseDataStorageConfigType string

const (
	RecordingGetOneRecordingResponseDataStorageConfigTypeAws          RecordingGetOneRecordingResponseDataStorageConfigType = "aws"
	RecordingGetOneRecordingResponseDataStorageConfigTypeAzure        RecordingGetOneRecordingResponseDataStorageConfigType = "azure"
	RecordingGetOneRecordingResponseDataStorageConfigTypeDigitalocean RecordingGetOneRecordingResponseDataStorageConfigType = "digitalocean"
	RecordingGetOneRecordingResponseDataStorageConfigTypeGcs          RecordingGetOneRecordingResponseDataStorageConfigType = "gcs"
	RecordingGetOneRecordingResponseDataStorageConfigTypeSftp         RecordingGetOneRecordingResponseDataStorageConfigType = "sftp"
)

func (r RecordingGetOneRecordingResponseDataStorageConfigType) IsKnown() bool {
	switch r {
	case RecordingGetOneRecordingResponseDataStorageConfigTypeAws, RecordingGetOneRecordingResponseDataStorageConfigTypeAzure, RecordingGetOneRecordingResponseDataStorageConfigTypeDigitalocean, RecordingGetOneRecordingResponseDataStorageConfigTypeGcs, RecordingGetOneRecordingResponseDataStorageConfigTypeSftp:
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

type RecordingGetRecordingsResponse struct {
	Data    []RecordingGetRecordingsResponseData `json:"data,required"`
	Paging  RecordingGetRecordingsResponsePaging `json:"paging,required"`
	Success bool                                 `json:"success,required"`
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
	ID string `json:"id,required" format:"uuid"`
	// If the audio_config is passed, the URL for downloading the audio recording is
	// returned.
	AudioDownloadURL string `json:"audio_download_url,required,nullable" format:"uri"`
	// URL where the recording can be downloaded.
	DownloadURL string `json:"download_url,required,nullable" format:"uri"`
	// Timestamp when the download URL expires.
	DownloadURLExpiry time.Time `json:"download_url_expiry,required,nullable" format:"date-time"`
	// File size of the recording, in bytes.
	FileSize float64 `json:"file_size,required,nullable"`
	// Timestamp when this recording was invoked.
	InvokedTime time.Time `json:"invoked_time,required" format:"date-time"`
	// File name of the recording.
	OutputFileName string `json:"output_file_name,required"`
	// ID of the meeting session this recording is for.
	SessionID string `json:"session_id,required,nullable" format:"uuid"`
	// Timestamp when this recording actually started after being invoked. Usually a
	// few seconds after `invoked_time`.
	StartedTime time.Time `json:"started_time,required,nullable" format:"date-time"`
	// Current status of the recording.
	Status RecordingGetRecordingsResponseDataStatus `json:"status,required"`
	// Timestamp when this recording was stopped. Optional; is present only when the
	// recording has actually been stopped.
	StoppedTime time.Time                                 `json:"stopped_time,required,nullable" format:"date-time"`
	Meeting     RecordingGetRecordingsResponseDataMeeting `json:"meeting"`
	// Total recording time in seconds.
	RecordingDuration int64                                           `json:"recording_duration"`
	StorageConfig     RecordingGetRecordingsResponseDataStorageConfig `json:"storage_config,nullable"`
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
	Status RecordingGetRecordingsResponseDataMeetingStatus `json:"status"`
	// Automatically generate summary of meetings using transcripts. Requires
	// Transcriptions to be enabled, and can be retrieved via Webhooks or summary API.
	SummarizeOnEnd bool `json:"summarize_on_end"`
	// Title of the meeting.
	Title string                                        `json:"title"`
	JSON  recordingGetRecordingsResponseDataMeetingJSON `json:"-"`
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
	SessionKeepAliveTimeInSecs apijson.Field
	Status                     apijson.Field
	SummarizeOnEnd             apijson.Field
	Title                      apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *RecordingGetRecordingsResponseDataMeeting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseDataMeetingJSON) RawJSON() string {
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
	// Type of storage media.
	Type RecordingGetRecordingsResponseDataStorageConfigType `json:"type,required"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingGetRecordingsResponseDataStorageConfigAuthMethod `json:"auth_method"`
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
	Username string                                              `json:"username"`
	JSON     recordingGetRecordingsResponseDataStorageConfigJSON `json:"-"`
}

// recordingGetRecordingsResponseDataStorageConfigJSON contains the JSON metadata
// for the struct [RecordingGetRecordingsResponseDataStorageConfig]
type recordingGetRecordingsResponseDataStorageConfigJSON struct {
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

func (r *RecordingGetRecordingsResponseDataStorageConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingGetRecordingsResponseDataStorageConfigJSON) RawJSON() string {
	return r.raw
}

// Type of storage media.
type RecordingGetRecordingsResponseDataStorageConfigType string

const (
	RecordingGetRecordingsResponseDataStorageConfigTypeAws          RecordingGetRecordingsResponseDataStorageConfigType = "aws"
	RecordingGetRecordingsResponseDataStorageConfigTypeAzure        RecordingGetRecordingsResponseDataStorageConfigType = "azure"
	RecordingGetRecordingsResponseDataStorageConfigTypeDigitalocean RecordingGetRecordingsResponseDataStorageConfigType = "digitalocean"
	RecordingGetRecordingsResponseDataStorageConfigTypeGcs          RecordingGetRecordingsResponseDataStorageConfigType = "gcs"
	RecordingGetRecordingsResponseDataStorageConfigTypeSftp         RecordingGetRecordingsResponseDataStorageConfigType = "sftp"
)

func (r RecordingGetRecordingsResponseDataStorageConfigType) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsResponseDataStorageConfigTypeAws, RecordingGetRecordingsResponseDataStorageConfigTypeAzure, RecordingGetRecordingsResponseDataStorageConfigTypeDigitalocean, RecordingGetRecordingsResponseDataStorageConfigTypeGcs, RecordingGetRecordingsResponseDataStorageConfigTypeSftp:
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

type RecordingGetRecordingsResponsePaging struct {
	EndOffset   float64                                  `json:"end_offset,required"`
	StartOffset float64                                  `json:"start_offset,required"`
	TotalCount  float64                                  `json:"total_count,required"`
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
	Success bool `json:"success,required"`
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
	ID string `json:"id,required" format:"uuid"`
	// If the audio_config is passed, the URL for downloading the audio recording is
	// returned.
	AudioDownloadURL string `json:"audio_download_url,required,nullable" format:"uri"`
	// URL where the recording can be downloaded.
	DownloadURL string `json:"download_url,required,nullable" format:"uri"`
	// Timestamp when the download URL expires.
	DownloadURLExpiry time.Time `json:"download_url_expiry,required,nullable" format:"date-time"`
	// File size of the recording, in bytes.
	FileSize float64 `json:"file_size,required,nullable"`
	// Timestamp when this recording was invoked.
	InvokedTime time.Time `json:"invoked_time,required" format:"date-time"`
	// File name of the recording.
	OutputFileName string `json:"output_file_name,required"`
	// ID of the meeting session this recording is for.
	SessionID string `json:"session_id,required,nullable" format:"uuid"`
	// Timestamp when this recording actually started after being invoked. Usually a
	// few seconds after `invoked_time`.
	StartedTime time.Time `json:"started_time,required,nullable" format:"date-time"`
	// Current status of the recording.
	Status RecordingPauseResumeStopRecordingResponseDataStatus `json:"status,required"`
	// Timestamp when this recording was stopped. Optional; is present only when the
	// recording has actually been stopped.
	StoppedTime time.Time `json:"stopped_time,required,nullable" format:"date-time"`
	// Total recording time in seconds.
	RecordingDuration int64                                                      `json:"recording_duration"`
	StartReason       RecordingPauseResumeStopRecordingResponseDataStartReason   `json:"start_reason"`
	StopReason        RecordingPauseResumeStopRecordingResponseDataStopReason    `json:"stop_reason"`
	StorageConfig     RecordingPauseResumeStopRecordingResponseDataStorageConfig `json:"storage_config,nullable"`
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
	// Type of storage media.
	Type RecordingPauseResumeStopRecordingResponseDataStorageConfigType `json:"type,required"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingPauseResumeStopRecordingResponseDataStorageConfigAuthMethod `json:"auth_method"`
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
	Username string                                                         `json:"username"`
	JSON     recordingPauseResumeStopRecordingResponseDataStorageConfigJSON `json:"-"`
}

// recordingPauseResumeStopRecordingResponseDataStorageConfigJSON contains the JSON
// metadata for the struct
// [RecordingPauseResumeStopRecordingResponseDataStorageConfig]
type recordingPauseResumeStopRecordingResponseDataStorageConfigJSON struct {
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

func (r *RecordingPauseResumeStopRecordingResponseDataStorageConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingPauseResumeStopRecordingResponseDataStorageConfigJSON) RawJSON() string {
	return r.raw
}

// Type of storage media.
type RecordingPauseResumeStopRecordingResponseDataStorageConfigType string

const (
	RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeAws          RecordingPauseResumeStopRecordingResponseDataStorageConfigType = "aws"
	RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeAzure        RecordingPauseResumeStopRecordingResponseDataStorageConfigType = "azure"
	RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeDigitalocean RecordingPauseResumeStopRecordingResponseDataStorageConfigType = "digitalocean"
	RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeGcs          RecordingPauseResumeStopRecordingResponseDataStorageConfigType = "gcs"
	RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeSftp         RecordingPauseResumeStopRecordingResponseDataStorageConfigType = "sftp"
)

func (r RecordingPauseResumeStopRecordingResponseDataStorageConfigType) IsKnown() bool {
	switch r {
	case RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeAws, RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeAzure, RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeDigitalocean, RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeGcs, RecordingPauseResumeStopRecordingResponseDataStorageConfigTypeSftp:
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

type RecordingStartRecordingsResponse struct {
	// Success status of the operation
	Success bool `json:"success,required"`
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
	ID string `json:"id,required" format:"uuid"`
	// If the audio_config is passed, the URL for downloading the audio recording is
	// returned.
	AudioDownloadURL string `json:"audio_download_url,required,nullable" format:"uri"`
	// URL where the recording can be downloaded.
	DownloadURL string `json:"download_url,required,nullable" format:"uri"`
	// Timestamp when the download URL expires.
	DownloadURLExpiry time.Time `json:"download_url_expiry,required,nullable" format:"date-time"`
	// File size of the recording, in bytes.
	FileSize float64 `json:"file_size,required,nullable"`
	// Timestamp when this recording was invoked.
	InvokedTime time.Time `json:"invoked_time,required" format:"date-time"`
	// File name of the recording.
	OutputFileName string `json:"output_file_name,required"`
	// ID of the meeting session this recording is for.
	SessionID string `json:"session_id,required,nullable" format:"uuid"`
	// Timestamp when this recording actually started after being invoked. Usually a
	// few seconds after `invoked_time`.
	StartedTime time.Time `json:"started_time,required,nullable" format:"date-time"`
	// Current status of the recording.
	Status RecordingStartRecordingsResponseDataStatus `json:"status,required"`
	// Timestamp when this recording was stopped. Optional; is present only when the
	// recording has actually been stopped.
	StoppedTime time.Time `json:"stopped_time,required,nullable" format:"date-time"`
	// Total recording time in seconds.
	RecordingDuration int64                                             `json:"recording_duration"`
	StartReason       RecordingStartRecordingsResponseDataStartReason   `json:"start_reason"`
	StopReason        RecordingStartRecordingsResponseDataStopReason    `json:"stop_reason"`
	StorageConfig     RecordingStartRecordingsResponseDataStorageConfig `json:"storage_config,nullable"`
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
	// Type of storage media.
	Type RecordingStartRecordingsResponseDataStorageConfigType `json:"type,required"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod RecordingStartRecordingsResponseDataStorageConfigAuthMethod `json:"auth_method"`
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
	Username string                                                `json:"username"`
	JSON     recordingStartRecordingsResponseDataStorageConfigJSON `json:"-"`
}

// recordingStartRecordingsResponseDataStorageConfigJSON contains the JSON metadata
// for the struct [RecordingStartRecordingsResponseDataStorageConfig]
type recordingStartRecordingsResponseDataStorageConfigJSON struct {
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

func (r *RecordingStartRecordingsResponseDataStorageConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordingStartRecordingsResponseDataStorageConfigJSON) RawJSON() string {
	return r.raw
}

// Type of storage media.
type RecordingStartRecordingsResponseDataStorageConfigType string

const (
	RecordingStartRecordingsResponseDataStorageConfigTypeAws          RecordingStartRecordingsResponseDataStorageConfigType = "aws"
	RecordingStartRecordingsResponseDataStorageConfigTypeAzure        RecordingStartRecordingsResponseDataStorageConfigType = "azure"
	RecordingStartRecordingsResponseDataStorageConfigTypeDigitalocean RecordingStartRecordingsResponseDataStorageConfigType = "digitalocean"
	RecordingStartRecordingsResponseDataStorageConfigTypeGcs          RecordingStartRecordingsResponseDataStorageConfigType = "gcs"
	RecordingStartRecordingsResponseDataStorageConfigTypeSftp         RecordingStartRecordingsResponseDataStorageConfigType = "sftp"
)

func (r RecordingStartRecordingsResponseDataStorageConfigType) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsResponseDataStorageConfigTypeAws, RecordingStartRecordingsResponseDataStorageConfigTypeAzure, RecordingStartRecordingsResponseDataStorageConfigTypeDigitalocean, RecordingStartRecordingsResponseDataStorageConfigTypeGcs, RecordingStartRecordingsResponseDataStorageConfigTypeSftp:
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

type RecordingGetActiveRecordingsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RecordingGetOneRecordingParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RecordingGetRecordingsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
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
)

func (r RecordingGetRecordingsParamsStatus) IsKnown() bool {
	switch r {
	case RecordingGetRecordingsParamsStatusInvoked, RecordingGetRecordingsParamsStatusRecording, RecordingGetRecordingsParamsStatusUploading, RecordingGetRecordingsParamsStatusUploaded:
		return true
	}
	return false
}

type RecordingPauseResumeStopRecordingParams struct {
	// The account identifier tag.
	AccountID param.Field[string]                                        `path:"account_id,required"`
	Action    param.Field[RecordingPauseResumeStopRecordingParamsAction] `json:"action,required"`
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
	AccountID param.Field[string] `path:"account_id,required"`
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
	MaxSeconds param.Field[int64] `json:"max_seconds"`
	// ID of the meeting to record.
	MeetingID               param.Field[string]                                                `json:"meeting_id" format:"uuid"`
	RealtimekitBucketConfig param.Field[RecordingStartRecordingsParamsRealtimekitBucketConfig] `json:"realtimekit_bucket_config"`
	RtmpOutConfig           param.Field[RecordingStartRecordingsParamsRtmpOutConfig]           `json:"rtmp_out_config"`
	StorageConfig           param.Field[RecordingStartRecordingsParamsStorageConfig]           `json:"storage_config"`
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
	Enabled param.Field[bool] `json:"enabled,required"`
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
	// Type of storage media.
	Type param.Field[RecordingStartRecordingsParamsStorageConfigType] `json:"type,required"`
	// Access key of the storage medium. Access key is not required for the `gcs`
	// storage media type.
	//
	// Note that this field is not readable by clients, only writeable.
	AccessKey param.Field[string] `json:"access_key"`
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
	PrivateKey param.Field[string] `json:"private_key"`
	// Region of the storage medium.
	Region param.Field[string] `json:"region"`
	// Secret key of the storage medium. Similar to `access_key`, it is only writeable
	// by clients, not readable.
	Secret param.Field[string] `json:"secret"`
	// SSH destination server username for SFTP type storage medium
	Username param.Field[string] `json:"username"`
}

func (r RecordingStartRecordingsParamsStorageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of storage media.
type RecordingStartRecordingsParamsStorageConfigType string

const (
	RecordingStartRecordingsParamsStorageConfigTypeAws          RecordingStartRecordingsParamsStorageConfigType = "aws"
	RecordingStartRecordingsParamsStorageConfigTypeAzure        RecordingStartRecordingsParamsStorageConfigType = "azure"
	RecordingStartRecordingsParamsStorageConfigTypeDigitalocean RecordingStartRecordingsParamsStorageConfigType = "digitalocean"
	RecordingStartRecordingsParamsStorageConfigTypeGcs          RecordingStartRecordingsParamsStorageConfigType = "gcs"
	RecordingStartRecordingsParamsStorageConfigTypeSftp         RecordingStartRecordingsParamsStorageConfigType = "sftp"
)

func (r RecordingStartRecordingsParamsStorageConfigType) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsParamsStorageConfigTypeAws, RecordingStartRecordingsParamsStorageConfigTypeAzure, RecordingStartRecordingsParamsStorageConfigTypeDigitalocean, RecordingStartRecordingsParamsStorageConfigTypeGcs, RecordingStartRecordingsParamsStorageConfigTypeSftp:
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
)

func (r RecordingStartRecordingsParamsVideoConfigCodec) IsKnown() bool {
	switch r {
	case RecordingStartRecordingsParamsVideoConfigCodecH264, RecordingStartRecordingsParamsVideoConfigCodecVp8:
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
	AccountID param.Field[string]                                              `path:"account_id,required"`
	Layers    param.Field[map[string]RecordingStartTrackRecordingParamsLayers] `json:"layers,required"`
	// ID of the meeting to record.
	MeetingID param.Field[string] `json:"meeting_id,required"`
	// Maximum seconds this recording should be active for (beta)
	MaxSeconds param.Field[float64] `json:"max_seconds"`
}

func (r RecordingStartTrackRecordingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RecordingStartTrackRecordingParamsLayers struct {
	// A file name prefix to apply for files generated from this layer
	FileNamePrefix param.Field[string]                                           `json:"file_name_prefix"`
	Outputs        param.Field[[]RecordingStartTrackRecordingParamsLayersOutput] `json:"outputs"`
}

func (r RecordingStartTrackRecordingParamsLayers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RecordingStartTrackRecordingParamsLayersOutput struct {
	StorageConfig param.Field[RecordingStartTrackRecordingParamsLayersOutputsStorageConfig] `json:"storage_config"`
	// The type of output destination this layer is being exported to.
	Type param.Field[RecordingStartTrackRecordingParamsLayersOutputsType] `json:"type"`
}

func (r RecordingStartTrackRecordingParamsLayersOutput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RecordingStartTrackRecordingParamsLayersOutputsStorageConfig struct {
	// Type of storage media.
	Type param.Field[RecordingStartTrackRecordingParamsLayersOutputsStorageConfigType] `json:"type,required"`
	// Access key of the storage medium. Access key is not required for the `gcs`
	// storage media type.
	//
	// Note that this field is not readable by clients, only writeable.
	AccessKey param.Field[string] `json:"access_key"`
	// Authentication method used for "sftp" type storage medium
	AuthMethod param.Field[RecordingStartTrackRecordingParamsLayersOutputsStorageConfigAuthMethod] `json:"auth_method"`
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

func (r RecordingStartTrackRecordingParamsLayersOutputsStorageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of storage media.
type RecordingStartTrackRecordingParamsLayersOutputsStorageConfigType string

const (
	RecordingStartTrackRecordingParamsLayersOutputsStorageConfigTypeAws          RecordingStartTrackRecordingParamsLayersOutputsStorageConfigType = "aws"
	RecordingStartTrackRecordingParamsLayersOutputsStorageConfigTypeAzure        RecordingStartTrackRecordingParamsLayersOutputsStorageConfigType = "azure"
	RecordingStartTrackRecordingParamsLayersOutputsStorageConfigTypeDigitalocean RecordingStartTrackRecordingParamsLayersOutputsStorageConfigType = "digitalocean"
	RecordingStartTrackRecordingParamsLayersOutputsStorageConfigTypeGcs          RecordingStartTrackRecordingParamsLayersOutputsStorageConfigType = "gcs"
	RecordingStartTrackRecordingParamsLayersOutputsStorageConfigTypeSftp         RecordingStartTrackRecordingParamsLayersOutputsStorageConfigType = "sftp"
)

func (r RecordingStartTrackRecordingParamsLayersOutputsStorageConfigType) IsKnown() bool {
	switch r {
	case RecordingStartTrackRecordingParamsLayersOutputsStorageConfigTypeAws, RecordingStartTrackRecordingParamsLayersOutputsStorageConfigTypeAzure, RecordingStartTrackRecordingParamsLayersOutputsStorageConfigTypeDigitalocean, RecordingStartTrackRecordingParamsLayersOutputsStorageConfigTypeGcs, RecordingStartTrackRecordingParamsLayersOutputsStorageConfigTypeSftp:
		return true
	}
	return false
}

// Authentication method used for "sftp" type storage medium
type RecordingStartTrackRecordingParamsLayersOutputsStorageConfigAuthMethod string

const (
	RecordingStartTrackRecordingParamsLayersOutputsStorageConfigAuthMethodKey      RecordingStartTrackRecordingParamsLayersOutputsStorageConfigAuthMethod = "KEY"
	RecordingStartTrackRecordingParamsLayersOutputsStorageConfigAuthMethodPassword RecordingStartTrackRecordingParamsLayersOutputsStorageConfigAuthMethod = "PASSWORD"
)

func (r RecordingStartTrackRecordingParamsLayersOutputsStorageConfigAuthMethod) IsKnown() bool {
	switch r {
	case RecordingStartTrackRecordingParamsLayersOutputsStorageConfigAuthMethodKey, RecordingStartTrackRecordingParamsLayersOutputsStorageConfigAuthMethodPassword:
		return true
	}
	return false
}

// The type of output destination this layer is being exported to.
type RecordingStartTrackRecordingParamsLayersOutputsType string

const (
	RecordingStartTrackRecordingParamsLayersOutputsTypeRealtimekitBucket RecordingStartTrackRecordingParamsLayersOutputsType = "REALTIMEKIT_BUCKET"
	RecordingStartTrackRecordingParamsLayersOutputsTypeStorageConfig     RecordingStartTrackRecordingParamsLayersOutputsType = "STORAGE_CONFIG"
)

func (r RecordingStartTrackRecordingParamsLayersOutputsType) IsKnown() bool {
	switch r {
	case RecordingStartTrackRecordingParamsLayersOutputsTypeRealtimekitBucket, RecordingStartTrackRecordingParamsLayersOutputsTypeStorageConfig:
		return true
	}
	return false
}
