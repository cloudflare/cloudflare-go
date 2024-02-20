// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
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

// Adds an additional audio track to a video using the provided audio track URL.
func (r *StreamAudioTrackService) New(ctx context.Context, accountID string, identifier string, body StreamAudioTrackNewParams, opts ...option.RequestOption) (res *StreamAudioTrackNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamAudioTrackNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/copy", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
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
func (r *StreamAudioTrackService) Delete(ctx context.Context, accountID string, identifier string, audioIdentifier string, opts ...option.RequestOption) (res *StreamAudioTrackDeleteResponse, err error) {
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

// Edits additional audio tracks on a video. Editing the default status of an audio
// track to `true` will mark all other audio tracks on the video default status to
// `false`.
func (r *StreamAudioTrackService) Edit(ctx context.Context, accountID string, identifier string, audioIdentifier string, body StreamAudioTrackEditParams, opts ...option.RequestOption) (res *StreamAudioTrackEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamAudioTrackEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/%s", accountID, identifier, audioIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamAudioTrackNewResponse struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status StreamAudioTrackNewResponseStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid  string                          `json:"uid"`
	JSON streamAudioTrackNewResponseJSON `json:"-"`
}

// streamAudioTrackNewResponseJSON contains the JSON metadata for the struct
// [StreamAudioTrackNewResponse]
type streamAudioTrackNewResponseJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type StreamAudioTrackNewResponseStatus string

const (
	StreamAudioTrackNewResponseStatusQueued StreamAudioTrackNewResponseStatus = "queued"
	StreamAudioTrackNewResponseStatusReady  StreamAudioTrackNewResponseStatus = "ready"
	StreamAudioTrackNewResponseStatusError  StreamAudioTrackNewResponseStatus = "error"
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

// Union satisfied by [StreamAudioTrackDeleteResponseUnknown] or
// [shared.UnionString].
type StreamAudioTrackDeleteResponse interface {
	ImplementsStreamAudioTrackDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamAudioTrackDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StreamAudioTrackEditResponse struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status StreamAudioTrackEditResponseStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid  string                           `json:"uid"`
	JSON streamAudioTrackEditResponseJSON `json:"-"`
}

// streamAudioTrackEditResponseJSON contains the JSON metadata for the struct
// [StreamAudioTrackEditResponse]
type streamAudioTrackEditResponseJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type StreamAudioTrackEditResponseStatus string

const (
	StreamAudioTrackEditResponseStatusQueued StreamAudioTrackEditResponseStatus = "queued"
	StreamAudioTrackEditResponseStatusReady  StreamAudioTrackEditResponseStatus = "ready"
	StreamAudioTrackEditResponseStatusError  StreamAudioTrackEditResponseStatus = "error"
)

type StreamAudioTrackNewParams struct {
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label param.Field[string] `json:"label,required"`
	// An audio track URL. The server must be publicly routable and support `HTTP HEAD`
	// requests and `HTTP GET` range requests. The server should respond to `HTTP HEAD`
	// requests with a `content-range` header that includes the size of the file.
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r StreamAudioTrackNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamAudioTrackNewResponseEnvelope struct {
	Errors   []StreamAudioTrackNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamAudioTrackNewResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamAudioTrackNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamAudioTrackNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamAudioTrackNewResponseEnvelopeJSON    `json:"-"`
}

// streamAudioTrackNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamAudioTrackNewResponseEnvelope]
type streamAudioTrackNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    streamAudioTrackNewResponseEnvelopeErrorsJSON `json:"-"`
}

// streamAudioTrackNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamAudioTrackNewResponseEnvelopeErrors]
type streamAudioTrackNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    streamAudioTrackNewResponseEnvelopeMessagesJSON `json:"-"`
}

// streamAudioTrackNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamAudioTrackNewResponseEnvelopeMessages]
type streamAudioTrackNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamAudioTrackNewResponseEnvelopeSuccess bool

const (
	StreamAudioTrackNewResponseEnvelopeSuccessTrue StreamAudioTrackNewResponseEnvelopeSuccess = true
)

type StreamAudioTrackListResponseEnvelope struct {
	Errors   []StreamAudioTrackListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamAudioTrackListResponseEnvelopeMessages `json:"messages,required"`
	Result   []StreamAudioTrackListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success StreamAudioTrackListResponseEnvelopeSuccess `json:"success,required"`
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
	Errors   []StreamAudioTrackDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamAudioTrackDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamAudioTrackDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamAudioTrackDeleteResponseEnvelopeSuccess `json:"success,required"`
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

type StreamAudioTrackEditParams struct {
	// Denotes whether the audio track will be played by default in a player.
	Default param.Field[bool] `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label param.Field[string] `json:"label"`
}

func (r StreamAudioTrackEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamAudioTrackEditResponseEnvelope struct {
	Errors   []StreamAudioTrackEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamAudioTrackEditResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamAudioTrackEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamAudioTrackEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamAudioTrackEditResponseEnvelopeJSON    `json:"-"`
}

// streamAudioTrackEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamAudioTrackEditResponseEnvelope]
type streamAudioTrackEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackEditResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    streamAudioTrackEditResponseEnvelopeErrorsJSON `json:"-"`
}

// streamAudioTrackEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StreamAudioTrackEditResponseEnvelopeErrors]
type streamAudioTrackEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackEditResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    streamAudioTrackEditResponseEnvelopeMessagesJSON `json:"-"`
}

// streamAudioTrackEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamAudioTrackEditResponseEnvelopeMessages]
type streamAudioTrackEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamAudioTrackEditResponseEnvelopeSuccess bool

const (
	StreamAudioTrackEditResponseEnvelopeSuccessTrue StreamAudioTrackEditResponseEnvelopeSuccess = true
)
