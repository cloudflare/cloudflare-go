// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// StreamAudioTrackService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStreamAudioTrackService] method
// instead.
type StreamAudioTrackService struct {
	Options []option.RequestOption
}

// NewStreamAudioTrackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStreamAudioTrackService(opts ...option.RequestOption) (r *StreamAudioTrackService) {
	r = &StreamAudioTrackService{}
	r.Options = opts
	return
}

// Edits additional audio tracks on a video. Editing the default status of an audio
// track to `true` will mark all other audio tracks on the video default status to
// `false`.
func (r *StreamAudioTrackService) Update(ctx context.Context, accountID string, identifier string, audioIdentifier string, body StreamAudioTrackUpdateParams, opts ...option.RequestOption) (res *StreamAudioTrackUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/%s", accountID, identifier, audioIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists additional audio tracks on a video. Note this API will not return
// information for audio attached to the video upload.
func (r *StreamAudioTrackService) List(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *StreamAudioTrackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/audio", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes additional audio tracks on a video. Deleting a default audio track is
// not allowed. You must assign another audio track as default prior to deletion.
func (r *StreamAudioTrackService) Delete(ctx context.Context, accountID string, identifier string, audioIdentifier string, opts ...option.RequestOption) (res *StreamAudioTrackDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/%s", accountID, identifier, audioIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds an additional audio track to a video using the provided audio track URL.
func (r *StreamAudioTrackService) Copy(ctx context.Context, accountID string, identifier string, body StreamAudioTrackCopyParams, opts ...option.RequestOption) (res *StreamAudioTrackCopyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/copy", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type StreamAudioTrackUpdateResponse struct {
	Errors   []StreamAudioTrackUpdateResponseError   `json:"errors"`
	Messages []StreamAudioTrackUpdateResponseMessage `json:"messages"`
	Result   StreamAudioTrackUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success StreamAudioTrackUpdateResponseSuccess `json:"success"`
	JSON    streamAudioTrackUpdateResponseJSON    `json:"-"`
}

// streamAudioTrackUpdateResponseJSON contains the JSON metadata for the struct
// [StreamAudioTrackUpdateResponse]
type streamAudioTrackUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackUpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    streamAudioTrackUpdateResponseErrorJSON `json:"-"`
}

// streamAudioTrackUpdateResponseErrorJSON contains the JSON metadata for the
// struct [StreamAudioTrackUpdateResponseError]
type streamAudioTrackUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackUpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    streamAudioTrackUpdateResponseMessageJSON `json:"-"`
}

// streamAudioTrackUpdateResponseMessageJSON contains the JSON metadata for the
// struct [StreamAudioTrackUpdateResponseMessage]
type streamAudioTrackUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackUpdateResponseResult struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status StreamAudioTrackUpdateResponseResultStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid  string                                   `json:"uid"`
	JSON streamAudioTrackUpdateResponseResultJSON `json:"-"`
}

// streamAudioTrackUpdateResponseResultJSON contains the JSON metadata for the
// struct [StreamAudioTrackUpdateResponseResult]
type streamAudioTrackUpdateResponseResultJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type StreamAudioTrackUpdateResponseResultStatus string

const (
	StreamAudioTrackUpdateResponseResultStatusQueued StreamAudioTrackUpdateResponseResultStatus = "queued"
	StreamAudioTrackUpdateResponseResultStatusReady  StreamAudioTrackUpdateResponseResultStatus = "ready"
	StreamAudioTrackUpdateResponseResultStatusError  StreamAudioTrackUpdateResponseResultStatus = "error"
)

// Whether the API call was successful
type StreamAudioTrackUpdateResponseSuccess bool

const (
	StreamAudioTrackUpdateResponseSuccessTrue StreamAudioTrackUpdateResponseSuccess = true
)

type StreamAudioTrackListResponse struct {
	Errors   []StreamAudioTrackListResponseError   `json:"errors"`
	Messages []StreamAudioTrackListResponseMessage `json:"messages"`
	Result   []StreamAudioTrackListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success StreamAudioTrackListResponseSuccess `json:"success"`
	JSON    streamAudioTrackListResponseJSON    `json:"-"`
}

// streamAudioTrackListResponseJSON contains the JSON metadata for the struct
// [StreamAudioTrackListResponse]
type streamAudioTrackListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackListResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    streamAudioTrackListResponseErrorJSON `json:"-"`
}

// streamAudioTrackListResponseErrorJSON contains the JSON metadata for the struct
// [StreamAudioTrackListResponseError]
type streamAudioTrackListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackListResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    streamAudioTrackListResponseMessageJSON `json:"-"`
}

// streamAudioTrackListResponseMessageJSON contains the JSON metadata for the
// struct [StreamAudioTrackListResponseMessage]
type streamAudioTrackListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackListResponseResult struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status StreamAudioTrackListResponseResultStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid  string                                 `json:"uid"`
	JSON streamAudioTrackListResponseResultJSON `json:"-"`
}

// streamAudioTrackListResponseResultJSON contains the JSON metadata for the struct
// [StreamAudioTrackListResponseResult]
type streamAudioTrackListResponseResultJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type StreamAudioTrackListResponseResultStatus string

const (
	StreamAudioTrackListResponseResultStatusQueued StreamAudioTrackListResponseResultStatus = "queued"
	StreamAudioTrackListResponseResultStatusReady  StreamAudioTrackListResponseResultStatus = "ready"
	StreamAudioTrackListResponseResultStatusError  StreamAudioTrackListResponseResultStatus = "error"
)

// Whether the API call was successful
type StreamAudioTrackListResponseSuccess bool

const (
	StreamAudioTrackListResponseSuccessTrue StreamAudioTrackListResponseSuccess = true
)

type StreamAudioTrackDeleteResponse struct {
	Errors   []StreamAudioTrackDeleteResponseError   `json:"errors"`
	Messages []StreamAudioTrackDeleteResponseMessage `json:"messages"`
	Result   string                                  `json:"result"`
	// Whether the API call was successful
	Success StreamAudioTrackDeleteResponseSuccess `json:"success"`
	JSON    streamAudioTrackDeleteResponseJSON    `json:"-"`
}

// streamAudioTrackDeleteResponseJSON contains the JSON metadata for the struct
// [StreamAudioTrackDeleteResponse]
type streamAudioTrackDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    streamAudioTrackDeleteResponseErrorJSON `json:"-"`
}

// streamAudioTrackDeleteResponseErrorJSON contains the JSON metadata for the
// struct [StreamAudioTrackDeleteResponseError]
type streamAudioTrackDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    streamAudioTrackDeleteResponseMessageJSON `json:"-"`
}

// streamAudioTrackDeleteResponseMessageJSON contains the JSON metadata for the
// struct [StreamAudioTrackDeleteResponseMessage]
type streamAudioTrackDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamAudioTrackDeleteResponseSuccess bool

const (
	StreamAudioTrackDeleteResponseSuccessTrue StreamAudioTrackDeleteResponseSuccess = true
)

type StreamAudioTrackCopyResponse struct {
	Errors   []StreamAudioTrackCopyResponseError   `json:"errors"`
	Messages []StreamAudioTrackCopyResponseMessage `json:"messages"`
	Result   StreamAudioTrackCopyResponseResult    `json:"result"`
	// Whether the API call was successful
	Success StreamAudioTrackCopyResponseSuccess `json:"success"`
	JSON    streamAudioTrackCopyResponseJSON    `json:"-"`
}

// streamAudioTrackCopyResponseJSON contains the JSON metadata for the struct
// [StreamAudioTrackCopyResponse]
type streamAudioTrackCopyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackCopyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackCopyResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    streamAudioTrackCopyResponseErrorJSON `json:"-"`
}

// streamAudioTrackCopyResponseErrorJSON contains the JSON metadata for the struct
// [StreamAudioTrackCopyResponseError]
type streamAudioTrackCopyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackCopyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackCopyResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    streamAudioTrackCopyResponseMessageJSON `json:"-"`
}

// streamAudioTrackCopyResponseMessageJSON contains the JSON metadata for the
// struct [StreamAudioTrackCopyResponseMessage]
type streamAudioTrackCopyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackCopyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackCopyResponseResult struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status StreamAudioTrackCopyResponseResultStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid  string                                 `json:"uid"`
	JSON streamAudioTrackCopyResponseResultJSON `json:"-"`
}

// streamAudioTrackCopyResponseResultJSON contains the JSON metadata for the struct
// [StreamAudioTrackCopyResponseResult]
type streamAudioTrackCopyResponseResultJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackCopyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type StreamAudioTrackCopyResponseResultStatus string

const (
	StreamAudioTrackCopyResponseResultStatusQueued StreamAudioTrackCopyResponseResultStatus = "queued"
	StreamAudioTrackCopyResponseResultStatusReady  StreamAudioTrackCopyResponseResultStatus = "ready"
	StreamAudioTrackCopyResponseResultStatusError  StreamAudioTrackCopyResponseResultStatus = "error"
)

// Whether the API call was successful
type StreamAudioTrackCopyResponseSuccess bool

const (
	StreamAudioTrackCopyResponseSuccessTrue StreamAudioTrackCopyResponseSuccess = true
)

type StreamAudioTrackUpdateParams struct {
	// Denotes whether the audio track will be played by default in a player.
	Default param.Field[bool] `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label param.Field[string] `json:"label"`
}

func (r StreamAudioTrackUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamAudioTrackCopyParams struct {
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label param.Field[string] `json:"label,required"`
	// An audio track URL. The server must be publicly routable and support `HTTP HEAD`
	// requests and `HTTP GET` range requests. The server should respond to `HTTP HEAD`
	// requests with a `content-range` header that includes the size of the file.
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r StreamAudioTrackCopyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
