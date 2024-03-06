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

// Deletes additional audio tracks on a video. Deleting a default audio track is
// not allowed. You must assign another audio track as default prior to deletion.
func (r *StreamAudioTrackService) Delete(ctx context.Context, identifier string, audioIdentifier string, body StreamAudioTrackDeleteParams, opts ...option.RequestOption) (res *StreamAudioTrackDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamAudioTrackDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/%s", body.AccountID, identifier, audioIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Adds an additional audio track to a video using the provided audio track URL.
func (r *StreamAudioTrackService) Copy(ctx context.Context, identifier string, params StreamAudioTrackCopyParams, opts ...option.RequestOption) (res *StreamAudioTrackCopyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamAudioTrackCopyResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/copy", params.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edits additional audio tracks on a video. Editing the default status of an audio
// track to `true` will mark all other audio tracks on the video default status to
// `false`.
func (r *StreamAudioTrackService) Edit(ctx context.Context, identifier string, audioIdentifier string, params StreamAudioTrackEditParams, opts ...option.RequestOption) (res *StreamAudioTrackEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamAudioTrackEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/%s", params.AccountID, identifier, audioIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists additional audio tracks on a video. Note this API will not return
// information for audio attached to the video upload.
func (r *StreamAudioTrackService) Get(ctx context.Context, identifier string, query StreamAudioTrackGetParams, opts ...option.RequestOption) (res *[]StreamAudioTrackGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamAudioTrackGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/audio", query.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

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

type StreamAudioTrackGetResponse struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status StreamAudioTrackGetResponseStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid  string                          `json:"uid"`
	JSON streamAudioTrackGetResponseJSON `json:"-"`
}

// streamAudioTrackGetResponseJSON contains the JSON metadata for the struct
// [StreamAudioTrackGetResponse]
type streamAudioTrackGetResponseJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type StreamAudioTrackGetResponseStatus string

const (
	StreamAudioTrackGetResponseStatusQueued StreamAudioTrackGetResponseStatus = "queued"
	StreamAudioTrackGetResponseStatusReady  StreamAudioTrackGetResponseStatus = "ready"
	StreamAudioTrackGetResponseStatusError  StreamAudioTrackGetResponseStatus = "error"
)

type StreamAudioTrackDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

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

type StreamAudioTrackCopyParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
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
	Errors   []StreamAudioTrackCopyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamAudioTrackCopyResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamAudioTrackCopyResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamAudioTrackCopyResponseEnvelopeSuccess `json:"success,required"`
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

type StreamAudioTrackEditParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
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

type StreamAudioTrackGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamAudioTrackGetResponseEnvelope struct {
	Errors   []StreamAudioTrackGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamAudioTrackGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []StreamAudioTrackGetResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success StreamAudioTrackGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamAudioTrackGetResponseEnvelopeJSON    `json:"-"`
}

// streamAudioTrackGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamAudioTrackGetResponseEnvelope]
type streamAudioTrackGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    streamAudioTrackGetResponseEnvelopeErrorsJSON `json:"-"`
}

// streamAudioTrackGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamAudioTrackGetResponseEnvelopeErrors]
type streamAudioTrackGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamAudioTrackGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    streamAudioTrackGetResponseEnvelopeMessagesJSON `json:"-"`
}

// streamAudioTrackGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamAudioTrackGetResponseEnvelopeMessages]
type streamAudioTrackGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAudioTrackGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamAudioTrackGetResponseEnvelopeSuccess bool

const (
	StreamAudioTrackGetResponseEnvelopeSuccessTrue StreamAudioTrackGetResponseEnvelopeSuccess = true
)
