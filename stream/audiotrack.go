// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// AudioTrackService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAudioTrackService] method instead.
type AudioTrackService struct {
	Options []option.RequestOption
}

// NewAudioTrackService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAudioTrackService(opts ...option.RequestOption) (r *AudioTrackService) {
	r = &AudioTrackService{}
	r.Options = opts
	return
}

// Deletes additional audio tracks on a video. Deleting a default audio track is
// not allowed. You must assign another audio track as default prior to deletion.
func (r *AudioTrackService) Delete(ctx context.Context, identifier string, audioIdentifier string, body AudioTrackDeleteParams, opts ...option.RequestOption) (res *AudioTrackDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AudioTrackDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/%s", body.AccountID, identifier, audioIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Adds an additional audio track to a video using the provided audio track URL.
func (r *AudioTrackService) Copy(ctx context.Context, identifier string, params AudioTrackCopyParams, opts ...option.RequestOption) (res *StreamAdditionalAudio, err error) {
	opts = append(r.Options[:], opts...)
	var env AudioTrackCopyResponseEnvelope
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
func (r *AudioTrackService) Edit(ctx context.Context, identifier string, audioIdentifier string, params AudioTrackEditParams, opts ...option.RequestOption) (res *StreamAdditionalAudio, err error) {
	opts = append(r.Options[:], opts...)
	var env AudioTrackEditResponseEnvelope
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
func (r *AudioTrackService) Get(ctx context.Context, identifier string, query AudioTrackGetParams, opts ...option.RequestOption) (res *[]StreamAdditionalAudio, err error) {
	opts = append(r.Options[:], opts...)
	var env AudioTrackGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/audio", query.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamAdditionalAudio struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status StreamAdditionalAudioStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid  string                    `json:"uid"`
	JSON streamAdditionalAudioJSON `json:"-"`
}

// streamAdditionalAudioJSON contains the JSON metadata for the struct
// [StreamAdditionalAudio]
type streamAdditionalAudioJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamAdditionalAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamAdditionalAudioJSON) RawJSON() string {
	return r.raw
}

// Specifies the processing status of the video.
type StreamAdditionalAudioStatus string

const (
	StreamAdditionalAudioStatusQueued StreamAdditionalAudioStatus = "queued"
	StreamAdditionalAudioStatusReady  StreamAdditionalAudioStatus = "ready"
	StreamAdditionalAudioStatusError  StreamAdditionalAudioStatus = "error"
)

// Union satisfied by [stream.AudioTrackDeleteResponseUnknown] or
// [shared.UnionString].
type AudioTrackDeleteResponse interface {
	ImplementsStreamAudioTrackDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AudioTrackDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AudioTrackDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AudioTrackDeleteResponseEnvelope struct {
	Errors   []AudioTrackDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AudioTrackDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AudioTrackDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AudioTrackDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    audioTrackDeleteResponseEnvelopeJSON    `json:"-"`
}

// audioTrackDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AudioTrackDeleteResponseEnvelope]
type audioTrackDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AudioTrackDeleteResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    audioTrackDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// audioTrackDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AudioTrackDeleteResponseEnvelopeErrors]
type audioTrackDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AudioTrackDeleteResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    audioTrackDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// audioTrackDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AudioTrackDeleteResponseEnvelopeMessages]
type audioTrackDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AudioTrackDeleteResponseEnvelopeSuccess bool

const (
	AudioTrackDeleteResponseEnvelopeSuccessTrue AudioTrackDeleteResponseEnvelopeSuccess = true
)

type AudioTrackCopyParams struct {
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

func (r AudioTrackCopyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudioTrackCopyResponseEnvelope struct {
	Errors   []AudioTrackCopyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AudioTrackCopyResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamAdditionalAudio                    `json:"result,required"`
	// Whether the API call was successful
	Success AudioTrackCopyResponseEnvelopeSuccess `json:"success,required"`
	JSON    audioTrackCopyResponseEnvelopeJSON    `json:"-"`
}

// audioTrackCopyResponseEnvelopeJSON contains the JSON metadata for the struct
// [AudioTrackCopyResponseEnvelope]
type audioTrackCopyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackCopyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackCopyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AudioTrackCopyResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    audioTrackCopyResponseEnvelopeErrorsJSON `json:"-"`
}

// audioTrackCopyResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AudioTrackCopyResponseEnvelopeErrors]
type audioTrackCopyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackCopyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackCopyResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AudioTrackCopyResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    audioTrackCopyResponseEnvelopeMessagesJSON `json:"-"`
}

// audioTrackCopyResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AudioTrackCopyResponseEnvelopeMessages]
type audioTrackCopyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackCopyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackCopyResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AudioTrackCopyResponseEnvelopeSuccess bool

const (
	AudioTrackCopyResponseEnvelopeSuccessTrue AudioTrackCopyResponseEnvelopeSuccess = true
)

type AudioTrackEditParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Denotes whether the audio track will be played by default in a player.
	Default param.Field[bool] `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label param.Field[string] `json:"label"`
}

func (r AudioTrackEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudioTrackEditResponseEnvelope struct {
	Errors   []AudioTrackEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AudioTrackEditResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamAdditionalAudio                    `json:"result,required"`
	// Whether the API call was successful
	Success AudioTrackEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    audioTrackEditResponseEnvelopeJSON    `json:"-"`
}

// audioTrackEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [AudioTrackEditResponseEnvelope]
type audioTrackEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AudioTrackEditResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    audioTrackEditResponseEnvelopeErrorsJSON `json:"-"`
}

// audioTrackEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AudioTrackEditResponseEnvelopeErrors]
type audioTrackEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AudioTrackEditResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    audioTrackEditResponseEnvelopeMessagesJSON `json:"-"`
}

// audioTrackEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AudioTrackEditResponseEnvelopeMessages]
type audioTrackEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AudioTrackEditResponseEnvelopeSuccess bool

const (
	AudioTrackEditResponseEnvelopeSuccessTrue AudioTrackEditResponseEnvelopeSuccess = true
)

type AudioTrackGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AudioTrackGetResponseEnvelope struct {
	Errors   []AudioTrackGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AudioTrackGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []StreamAdditionalAudio                 `json:"result,required"`
	// Whether the API call was successful
	Success AudioTrackGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    audioTrackGetResponseEnvelopeJSON    `json:"-"`
}

// audioTrackGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AudioTrackGetResponseEnvelope]
type audioTrackGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AudioTrackGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    audioTrackGetResponseEnvelopeErrorsJSON `json:"-"`
}

// audioTrackGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AudioTrackGetResponseEnvelopeErrors]
type audioTrackGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AudioTrackGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    audioTrackGetResponseEnvelopeMessagesJSON `json:"-"`
}

// audioTrackGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AudioTrackGetResponseEnvelopeMessages]
type audioTrackGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AudioTrackGetResponseEnvelopeSuccess bool

const (
	AudioTrackGetResponseEnvelopeSuccessTrue AudioTrackGetResponseEnvelopeSuccess = true
)
