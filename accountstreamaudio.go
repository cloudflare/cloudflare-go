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

// AccountStreamAudioService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamAudioService] method
// instead.
type AccountStreamAudioService struct {
	Options []option.RequestOption
}

// NewAccountStreamAudioService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamAudioService(opts ...option.RequestOption) (r *AccountStreamAudioService) {
	r = &AccountStreamAudioService{}
	r.Options = opts
	return
}

// Edits additional audio tracks on a video. Editing the default status of an audio
// track to `true` will mark all other audio tracks on the video default status to
// `false`.
func (r *AccountStreamAudioService) Update(ctx context.Context, accountIdentifier string, identifier string, audioIdentifier string, body AccountStreamAudioUpdateParams, opts ...option.RequestOption) (res *AccountStreamAudioUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/%s", accountIdentifier, identifier, audioIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes additional audio tracks on a video. Deleting a default audio track is
// not allowed. You must assign another audio track as default prior to deletion.
func (r *AccountStreamAudioService) Delete(ctx context.Context, accountIdentifier string, identifier string, audioIdentifier string, opts ...option.RequestOption) (res *AccountStreamAudioDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/%s", accountIdentifier, identifier, audioIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds an additional audio track to a video using the provided audio track URL.
func (r *AccountStreamAudioService) Copy(ctx context.Context, accountIdentifier string, identifier string, body AccountStreamAudioCopyParams, opts ...option.RequestOption) (res *AccountStreamAudioCopyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/audio/copy", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountStreamAudioUpdateResponse struct {
	Errors   []AccountStreamAudioUpdateResponseError   `json:"errors"`
	Messages []AccountStreamAudioUpdateResponseMessage `json:"messages"`
	Result   AccountStreamAudioUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStreamAudioUpdateResponseSuccess `json:"success"`
	JSON    accountStreamAudioUpdateResponseJSON    `json:"-"`
}

// accountStreamAudioUpdateResponseJSON contains the JSON metadata for the struct
// [AccountStreamAudioUpdateResponse]
type accountStreamAudioUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamAudioUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamAudioUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountStreamAudioUpdateResponseErrorJSON `json:"-"`
}

// accountStreamAudioUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountStreamAudioUpdateResponseError]
type accountStreamAudioUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamAudioUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamAudioUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountStreamAudioUpdateResponseMessageJSON `json:"-"`
}

// accountStreamAudioUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountStreamAudioUpdateResponseMessage]
type accountStreamAudioUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamAudioUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamAudioUpdateResponseResult struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status AccountStreamAudioUpdateResponseResultStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid  string                                     `json:"uid"`
	JSON accountStreamAudioUpdateResponseResultJSON `json:"-"`
}

// accountStreamAudioUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountStreamAudioUpdateResponseResult]
type accountStreamAudioUpdateResponseResultJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamAudioUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type AccountStreamAudioUpdateResponseResultStatus string

const (
	AccountStreamAudioUpdateResponseResultStatusQueued AccountStreamAudioUpdateResponseResultStatus = "queued"
	AccountStreamAudioUpdateResponseResultStatusReady  AccountStreamAudioUpdateResponseResultStatus = "ready"
	AccountStreamAudioUpdateResponseResultStatusError  AccountStreamAudioUpdateResponseResultStatus = "error"
)

// Whether the API call was successful
type AccountStreamAudioUpdateResponseSuccess bool

const (
	AccountStreamAudioUpdateResponseSuccessTrue AccountStreamAudioUpdateResponseSuccess = true
)

type AccountStreamAudioDeleteResponse struct {
	Errors   []AccountStreamAudioDeleteResponseError   `json:"errors"`
	Messages []AccountStreamAudioDeleteResponseMessage `json:"messages"`
	Result   string                                    `json:"result"`
	// Whether the API call was successful
	Success AccountStreamAudioDeleteResponseSuccess `json:"success"`
	JSON    accountStreamAudioDeleteResponseJSON    `json:"-"`
}

// accountStreamAudioDeleteResponseJSON contains the JSON metadata for the struct
// [AccountStreamAudioDeleteResponse]
type accountStreamAudioDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamAudioDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamAudioDeleteResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountStreamAudioDeleteResponseErrorJSON `json:"-"`
}

// accountStreamAudioDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountStreamAudioDeleteResponseError]
type accountStreamAudioDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamAudioDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamAudioDeleteResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountStreamAudioDeleteResponseMessageJSON `json:"-"`
}

// accountStreamAudioDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountStreamAudioDeleteResponseMessage]
type accountStreamAudioDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamAudioDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamAudioDeleteResponseSuccess bool

const (
	AccountStreamAudioDeleteResponseSuccessTrue AccountStreamAudioDeleteResponseSuccess = true
)

type AccountStreamAudioCopyResponse struct {
	Errors   []AccountStreamAudioCopyResponseError   `json:"errors"`
	Messages []AccountStreamAudioCopyResponseMessage `json:"messages"`
	Result   AccountStreamAudioCopyResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStreamAudioCopyResponseSuccess `json:"success"`
	JSON    accountStreamAudioCopyResponseJSON    `json:"-"`
}

// accountStreamAudioCopyResponseJSON contains the JSON metadata for the struct
// [AccountStreamAudioCopyResponse]
type accountStreamAudioCopyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamAudioCopyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamAudioCopyResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountStreamAudioCopyResponseErrorJSON `json:"-"`
}

// accountStreamAudioCopyResponseErrorJSON contains the JSON metadata for the
// struct [AccountStreamAudioCopyResponseError]
type accountStreamAudioCopyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamAudioCopyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamAudioCopyResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountStreamAudioCopyResponseMessageJSON `json:"-"`
}

// accountStreamAudioCopyResponseMessageJSON contains the JSON metadata for the
// struct [AccountStreamAudioCopyResponseMessage]
type accountStreamAudioCopyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamAudioCopyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamAudioCopyResponseResult struct {
	// Denotes whether the audio track will be played by default in a player.
	Default bool `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label string `json:"label"`
	// Specifies the processing status of the video.
	Status AccountStreamAudioCopyResponseResultStatus `json:"status"`
	// A Cloudflare-generated unique identifier for a media item.
	Uid  string                                   `json:"uid"`
	JSON accountStreamAudioCopyResponseResultJSON `json:"-"`
}

// accountStreamAudioCopyResponseResultJSON contains the JSON metadata for the
// struct [AccountStreamAudioCopyResponseResult]
type accountStreamAudioCopyResponseResultJSON struct {
	Default     apijson.Field
	Label       apijson.Field
	Status      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamAudioCopyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing status of the video.
type AccountStreamAudioCopyResponseResultStatus string

const (
	AccountStreamAudioCopyResponseResultStatusQueued AccountStreamAudioCopyResponseResultStatus = "queued"
	AccountStreamAudioCopyResponseResultStatusReady  AccountStreamAudioCopyResponseResultStatus = "ready"
	AccountStreamAudioCopyResponseResultStatusError  AccountStreamAudioCopyResponseResultStatus = "error"
)

// Whether the API call was successful
type AccountStreamAudioCopyResponseSuccess bool

const (
	AccountStreamAudioCopyResponseSuccessTrue AccountStreamAudioCopyResponseSuccess = true
)

type AccountStreamAudioUpdateParams struct {
	// Denotes whether the audio track will be played by default in a player.
	Default param.Field[bool] `json:"default"`
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label param.Field[string] `json:"label"`
}

func (r AccountStreamAudioUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountStreamAudioCopyParams struct {
	// A string to uniquely identify the track amongst other audio track labels for the
	// specified video.
	Label param.Field[string] `json:"label,required"`
	// An audio track URL. The server must be publicly routable and support `HTTP HEAD`
	// requests and `HTTP GET` range requests. The server should respond to `HTTP HEAD`
	// requests with a `content-range` header that includes the size of the file.
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r AccountStreamAudioCopyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
