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
	var env StreamAudioTrackUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/%s", accountID, identifier, audioIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists additional audio tracks on a video. Note this API will not return
// information for audio attached to the video upload.
func (r *StreamAudioTrackService) List(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *[]StreamAudioTrackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamAudioTrackListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/audio", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes additional audio tracks on a video. Deleting a default audio track is
// not allowed. You must assign another audio track as default prior to deletion.
func (r *StreamAudioTrackService) Delete(ctx context.Context, accountID string, identifier string, audioIdentifier string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamAudioTrackDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/%s", accountID, identifier, audioIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Adds an additional audio track to a video using the provided audio track URL.
func (r *StreamAudioTrackService) Copy(ctx context.Context, accountID string, identifier string, body StreamAudioTrackCopyParams, opts ...option.RequestOption) (res *StreamAudioTrackCopyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamAudioTrackCopyResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/copy", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamAudioTrackUpdateResponse struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status StreamAudioTrackUpdateResponseStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid  string                             `json:"uid"`
	JSON streamAudioTrackUpdateResponseJSON `json:"-"`
}

// streamAudioTrackUpdateResponseJSON contains the JSON metadata for the struct
// [StreamAudioTrackUpdateResponse]
type streamAudioTrackUpdateResponseJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type StreamAudioTrackUpdateResponseStatus string

const (
	StreamAudioTrackUpdateResponseStatusQueued StreamAudioTrackUpdateResponseStatus = "queued"
	StreamAudioTrackUpdateResponseStatusReady  StreamAudioTrackUpdateResponseStatus = "ready"
	StreamAudioTrackUpdateResponseStatusError  StreamAudioTrackUpdateResponseStatus = "error"
)

type StreamAudioTrackListResponse struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status StreamAudioTrackListResponseStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid  string                           `json:"uid"`
	JSON streamAudioTrackListResponseJSON `json:"-"`
}

// streamAudioTrackListResponseJSON contains the JSON metadata for the struct
// [StreamAudioTrackListResponse]
type streamAudioTrackListResponseJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type StreamAudioTrackListResponseStatus string

const (
	StreamAudioTrackListResponseStatusQueued StreamAudioTrackListResponseStatus = "queued"
	StreamAudioTrackListResponseStatusReady  StreamAudioTrackListResponseStatus = "ready"
	StreamAudioTrackListResponseStatusError  StreamAudioTrackListResponseStatus = "error"
)

type StreamAudioTrackCopyResponse struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status StreamAudioTrackCopyResponseStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid  string                           `json:"uid"`
	JSON streamAudioTrackCopyResponseJSON `json:"-"`
}

// streamAudioTrackCopyResponseJSON contains the JSON metadata for the struct
// [StreamAudioTrackCopyResponse]
type streamAudioTrackCopyResponseJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackCopyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type StreamAudioTrackCopyResponseStatus string

const (
	StreamAudioTrackCopyResponseStatusQueued StreamAudioTrackCopyResponseStatus = "queued"
	StreamAudioTrackCopyResponseStatusReady  StreamAudioTrackCopyResponseStatus = "ready"
	StreamAudioTrackCopyResponseStatusError  StreamAudioTrackCopyResponseStatus = "error"
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

type StreamAudioTrackUpdateResponseEnvelope struct {
	Errors   []StreamAudioTrackUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []StreamAudioTrackUpdateResponseEnvelopeMessages `json:"messages"`
	Result   StreamAudioTrackUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success StreamAudioTrackUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    streamAudioTrackUpdateResponseEnvelopeJSON    `json:"-"`
}

// streamAudioTrackUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamAudioTrackUpdateResponseEnvelope]
type streamAudioTrackUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    streamAudioTrackUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// streamAudioTrackUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StreamAudioTrackUpdateResponseEnvelopeErrors]
type streamAudioTrackUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    streamAudioTrackUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// streamAudioTrackUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [StreamAudioTrackUpdateResponseEnvelopeMessages]
type streamAudioTrackUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamAudioTrackUpdateResponseEnvelopeSuccess bool

const (
	StreamAudioTrackUpdateResponseEnvelopeSuccessTrue StreamAudioTrackUpdateResponseEnvelopeSuccess = true
)

type StreamAudioTrackListResponseEnvelope struct {
	Errors   []StreamAudioTrackListResponseEnvelopeErrors   `json:"errors"`
	Messages []StreamAudioTrackListResponseEnvelopeMessages `json:"messages"`
	Result   []StreamAudioTrackListResponse                 `json:"result"`
	// Whether the API call was successful
	Success StreamAudioTrackListResponseEnvelopeSuccess `json:"success"`
	JSON    streamAudioTrackListResponseEnvelopeJSON    `json:"-"`
}

// streamAudioTrackListResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamAudioTrackListResponseEnvelope]
type streamAudioTrackListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    streamAudioTrackListResponseEnvelopeErrorsJSON `json:"-"`
}

// streamAudioTrackListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StreamAudioTrackListResponseEnvelopeErrors]
type streamAudioTrackListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    streamAudioTrackListResponseEnvelopeMessagesJSON `json:"-"`
}

// streamAudioTrackListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamAudioTrackListResponseEnvelopeMessages]
type streamAudioTrackListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamAudioTrackListResponseEnvelopeSuccess bool

const (
	StreamAudioTrackListResponseEnvelopeSuccessTrue StreamAudioTrackListResponseEnvelopeSuccess = true
)

type StreamAudioTrackDeleteResponseEnvelope struct {
	Errors   []StreamAudioTrackDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []StreamAudioTrackDeleteResponseEnvelopeMessages `json:"messages"`
	Result   string                                           `json:"result"`
	// Whether the API call was successful
	Success StreamAudioTrackDeleteResponseEnvelopeSuccess `json:"success"`
	JSON    streamAudioTrackDeleteResponseEnvelopeJSON    `json:"-"`
}

// streamAudioTrackDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamAudioTrackDeleteResponseEnvelope]
type streamAudioTrackDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    streamAudioTrackDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// streamAudioTrackDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StreamAudioTrackDeleteResponseEnvelopeErrors]
type streamAudioTrackDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    streamAudioTrackDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// streamAudioTrackDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [StreamAudioTrackDeleteResponseEnvelopeMessages]
type streamAudioTrackDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamAudioTrackDeleteResponseEnvelopeSuccess bool

const (
	StreamAudioTrackDeleteResponseEnvelopeSuccessTrue StreamAudioTrackDeleteResponseEnvelopeSuccess = true
)

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

type StreamAudioTrackCopyResponseEnvelope struct {
	Errors   []StreamAudioTrackCopyResponseEnvelopeErrors   `json:"errors"`
	Messages []StreamAudioTrackCopyResponseEnvelopeMessages `json:"messages"`
	Result   StreamAudioTrackCopyResponse                   `json:"result"`
	// Whether the API call was successful
	Success StreamAudioTrackCopyResponseEnvelopeSuccess `json:"success"`
	JSON    streamAudioTrackCopyResponseEnvelopeJSON    `json:"-"`
}

// streamAudioTrackCopyResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamAudioTrackCopyResponseEnvelope]
type streamAudioTrackCopyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackCopyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackCopyResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    streamAudioTrackCopyResponseEnvelopeErrorsJSON `json:"-"`
}

// streamAudioTrackCopyResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StreamAudioTrackCopyResponseEnvelopeErrors]
type streamAudioTrackCopyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackCopyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackCopyResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    streamAudioTrackCopyResponseEnvelopeMessagesJSON `json:"-"`
}

// streamAudioTrackCopyResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamAudioTrackCopyResponseEnvelopeMessages]
type streamAudioTrackCopyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackCopyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamAudioTrackCopyResponseEnvelopeSuccess bool

const (
	StreamAudioTrackCopyResponseEnvelopeSuccessTrue StreamAudioTrackCopyResponseEnvelopeSuccess = true
)
