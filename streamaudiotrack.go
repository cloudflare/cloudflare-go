// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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

// Lists additional audio tracks on a video. Note this API will not return
// information for audio attached to the video upload.
func (r *StreamAudioTrackService) List(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *StreamAudioTrackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/audio", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

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
