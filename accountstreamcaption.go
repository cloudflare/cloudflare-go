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

// AccountStreamCaptionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamCaptionService]
// method instead.
type AccountStreamCaptionService struct {
	Options []option.RequestOption
}

// NewAccountStreamCaptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamCaptionService(opts ...option.RequestOption) (r *AccountStreamCaptionService) {
	r = &AccountStreamCaptionService{}
	r.Options = opts
	return
}

// Uploads the caption or subtitle file to the endpoint for a specific BCP47
// language. One caption or subtitle file per language is allowed.
func (r *AccountStreamCaptionService) Update(ctx context.Context, accountIdentifier string, identifier string, language string, body AccountStreamCaptionUpdateParams, opts ...option.RequestOption) (res *AccountStreamCaptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s", accountIdentifier, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Removes the captions or subtitles from a video.
func (r *AccountStreamCaptionService) Delete(ctx context.Context, accountIdentifier string, identifier string, language string, opts ...option.RequestOption) (res *AccountStreamCaptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s", accountIdentifier, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Lists the available captions or subtitles for a specific video.
func (r *AccountStreamCaptionService) StreamSubtitlesCaptionsListCaptionsOrSubtitles(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/captions", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountStreamCaptionUpdateResponse struct {
	Errors   []AccountStreamCaptionUpdateResponseError   `json:"errors"`
	Messages []AccountStreamCaptionUpdateResponseMessage `json:"messages"`
	Result   interface{}                                 `json:"result"`
	// Whether the API call was successful
	Success AccountStreamCaptionUpdateResponseSuccess `json:"success"`
	JSON    accountStreamCaptionUpdateResponseJSON    `json:"-"`
}

// accountStreamCaptionUpdateResponseJSON contains the JSON metadata for the struct
// [AccountStreamCaptionUpdateResponse]
type accountStreamCaptionUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCaptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamCaptionUpdateResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountStreamCaptionUpdateResponseErrorJSON `json:"-"`
}

// accountStreamCaptionUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountStreamCaptionUpdateResponseError]
type accountStreamCaptionUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCaptionUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamCaptionUpdateResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountStreamCaptionUpdateResponseMessageJSON `json:"-"`
}

// accountStreamCaptionUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountStreamCaptionUpdateResponseMessage]
type accountStreamCaptionUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCaptionUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamCaptionUpdateResponseSuccess bool

const (
	AccountStreamCaptionUpdateResponseSuccessTrue AccountStreamCaptionUpdateResponseSuccess = true
)

type AccountStreamCaptionDeleteResponse struct {
	Errors   []AccountStreamCaptionDeleteResponseError   `json:"errors"`
	Messages []AccountStreamCaptionDeleteResponseMessage `json:"messages"`
	Result   string                                      `json:"result"`
	// Whether the API call was successful
	Success AccountStreamCaptionDeleteResponseSuccess `json:"success"`
	JSON    accountStreamCaptionDeleteResponseJSON    `json:"-"`
}

// accountStreamCaptionDeleteResponseJSON contains the JSON metadata for the struct
// [AccountStreamCaptionDeleteResponse]
type accountStreamCaptionDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCaptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamCaptionDeleteResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountStreamCaptionDeleteResponseErrorJSON `json:"-"`
}

// accountStreamCaptionDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountStreamCaptionDeleteResponseError]
type accountStreamCaptionDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCaptionDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamCaptionDeleteResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountStreamCaptionDeleteResponseMessageJSON `json:"-"`
}

// accountStreamCaptionDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountStreamCaptionDeleteResponseMessage]
type accountStreamCaptionDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCaptionDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamCaptionDeleteResponseSuccess bool

const (
	AccountStreamCaptionDeleteResponseSuccessTrue AccountStreamCaptionDeleteResponseSuccess = true
)

type AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponse struct {
	Errors   []AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseError   `json:"errors"`
	Messages []AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseMessage `json:"messages"`
	Result   []AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseSuccess `json:"success"`
	JSON    accountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseJSON    `json:"-"`
}

// accountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseJSON
// contains the JSON metadata for the struct
// [AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponse]
type accountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseError struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    accountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseErrorJSON `json:"-"`
}

// accountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseError]
type accountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseMessage struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseMessageJSON `json:"-"`
}

// accountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseMessage]
type accountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseResult struct {
	// The language label displayed in the native language to users.
	Label string `json:"label"`
	// The language tag in BCP 47 format.
	Language string                                                                               `json:"language"`
	JSON     accountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseResultJSON `json:"-"`
}

// accountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseResultJSON
// contains the JSON metadata for the struct
// [AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseResult]
type accountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseResultJSON struct {
	Label       apijson.Field
	Language    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseSuccess bool

const (
	AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseSuccessTrue AccountStreamCaptionStreamSubtitlesCaptionsListCaptionsOrSubtitlesResponseSuccess = true
)

type AccountStreamCaptionUpdateParams struct {
	// The WebVTT file containing the caption or subtitle content.
	File param.Field[string] `json:"file,required"`
}

func (r AccountStreamCaptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
