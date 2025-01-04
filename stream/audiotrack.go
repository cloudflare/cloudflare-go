// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// AudioTrackService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioTrackService] method instead.
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
func (r *AudioTrackService) Delete(ctx context.Context, identifier string, audioIdentifier string, body AudioTrackDeleteParams, opts ...option.RequestOption) (res *string, err error) {
	var env AudioTrackDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if audioIdentifier == "" {
		err = errors.New("missing required audio_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/%s", body.AccountID, identifier, audioIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Adds an additional audio track to a video using the provided audio track URL.
func (r *AudioTrackService) Copy(ctx context.Context, identifier string, params AudioTrackCopyParams, opts ...option.RequestOption) (res *Audio, err error) {
	var env AudioTrackCopyResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
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
func (r *AudioTrackService) Edit(ctx context.Context, identifier string, audioIdentifier string, params AudioTrackEditParams, opts ...option.RequestOption) (res *Audio, err error) {
	var env AudioTrackEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if audioIdentifier == "" {
		err = errors.New("missing required audio_identifier parameter")
		return
	}
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
func (r *AudioTrackService) Get(ctx context.Context, identifier string, query AudioTrackGetParams, opts ...option.RequestOption) (res *[]Audio, err error) {
	var env AudioTrackGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/audio", query.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Audio struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status AudioStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	UID  string    `json:"uid"`
	JSON audioJSON `json:"-"`
}

// audioJSON contains the JSON metadata for the struct [Audio]
type audioJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Audio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioJSON) RawJSON() string {
	return r.raw
}

// Specifies the processing status of the video.
type AudioStatus string

const (
	AudioStatusQueued AudioStatus = "queued"
	AudioStatusReady  AudioStatus = "ready"
	AudioStatusError  AudioStatus = "error"
)

func (r AudioStatus) IsKnown() bool {
	switch r {
	case AudioStatusQueued, AudioStatusReady, AudioStatusError:
		return true
	}
	return false
}

type AudioTrackDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AudioTrackDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AudioTrackDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  string                                  `json:"result"`
	JSON    audioTrackDeleteResponseEnvelopeJSON    `json:"-"`
}

// audioTrackDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AudioTrackDeleteResponseEnvelope]
type audioTrackDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AudioTrackDeleteResponseEnvelopeSuccess bool

const (
	AudioTrackDeleteResponseEnvelopeSuccessTrue AudioTrackDeleteResponseEnvelopeSuccess = true
)

func (r AudioTrackDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AudioTrackDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AudioTrackCopyResponseEnvelopeSuccess `json:"success,required"`
	Result  Audio                                 `json:"result"`
	JSON    audioTrackCopyResponseEnvelopeJSON    `json:"-"`
}

// audioTrackCopyResponseEnvelopeJSON contains the JSON metadata for the struct
// [AudioTrackCopyResponseEnvelope]
type audioTrackCopyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackCopyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackCopyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AudioTrackCopyResponseEnvelopeSuccess bool

const (
	AudioTrackCopyResponseEnvelopeSuccessTrue AudioTrackCopyResponseEnvelopeSuccess = true
)

func (r AudioTrackCopyResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AudioTrackCopyResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AudioTrackEditResponseEnvelopeSuccess `json:"success,required"`
	Result  Audio                                 `json:"result"`
	JSON    audioTrackEditResponseEnvelopeJSON    `json:"-"`
}

// audioTrackEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [AudioTrackEditResponseEnvelope]
type audioTrackEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AudioTrackEditResponseEnvelopeSuccess bool

const (
	AudioTrackEditResponseEnvelopeSuccessTrue AudioTrackEditResponseEnvelopeSuccess = true
)

func (r AudioTrackEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AudioTrackEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AudioTrackGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AudioTrackGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AudioTrackGetResponseEnvelopeSuccess `json:"success,required"`
	Result  []Audio                              `json:"result"`
	JSON    audioTrackGetResponseEnvelopeJSON    `json:"-"`
}

// audioTrackGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AudioTrackGetResponseEnvelope]
type audioTrackGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTrackGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTrackGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AudioTrackGetResponseEnvelopeSuccess bool

const (
	AudioTrackGetResponseEnvelopeSuccessTrue AudioTrackGetResponseEnvelopeSuccess = true
)

func (r AudioTrackGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AudioTrackGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
