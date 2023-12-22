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
func (r *AccountStreamCaptionService) Update(ctx context.Context, accountIdentifier string, identifier string, language string, body AccountStreamCaptionUpdateParams, opts ...option.RequestOption) (res *LanguageResponseSingle, err error) {
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
func (r *AccountStreamCaptionService) StreamSubtitlesCaptionsListCaptionsOrSubtitles(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *LanguageResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/captions", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LanguageResponseCollection struct {
	Errors     []LanguageResponseCollectionError    `json:"errors"`
	Messages   []LanguageResponseCollectionMessage  `json:"messages"`
	Result     []LanguageResponseCollectionResult   `json:"result"`
	ResultInfo LanguageResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success LanguageResponseCollectionSuccess `json:"success"`
	JSON    languageResponseCollectionJSON    `json:"-"`
}

// languageResponseCollectionJSON contains the JSON metadata for the struct
// [LanguageResponseCollection]
type languageResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LanguageResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LanguageResponseCollectionError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    languageResponseCollectionErrorJSON `json:"-"`
}

// languageResponseCollectionErrorJSON contains the JSON metadata for the struct
// [LanguageResponseCollectionError]
type languageResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LanguageResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LanguageResponseCollectionMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    languageResponseCollectionMessageJSON `json:"-"`
}

// languageResponseCollectionMessageJSON contains the JSON metadata for the struct
// [LanguageResponseCollectionMessage]
type languageResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LanguageResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LanguageResponseCollectionResult struct {
	// The language label displayed in the native language to users.
	Label string `json:"label"`
	// The language tag in BCP 47 format.
	Language string                               `json:"language"`
	JSON     languageResponseCollectionResultJSON `json:"-"`
}

// languageResponseCollectionResultJSON contains the JSON metadata for the struct
// [LanguageResponseCollectionResult]
type languageResponseCollectionResultJSON struct {
	Label       apijson.Field
	Language    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LanguageResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LanguageResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       languageResponseCollectionResultInfoJSON `json:"-"`
}

// languageResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [LanguageResponseCollectionResultInfo]
type languageResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LanguageResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LanguageResponseCollectionSuccess bool

const (
	LanguageResponseCollectionSuccessTrue LanguageResponseCollectionSuccess = true
)

type LanguageResponseSingle struct {
	Errors   []LanguageResponseSingleError   `json:"errors"`
	Messages []LanguageResponseSingleMessage `json:"messages"`
	Result   interface{}                     `json:"result"`
	// Whether the API call was successful
	Success LanguageResponseSingleSuccess `json:"success"`
	JSON    languageResponseSingleJSON    `json:"-"`
}

// languageResponseSingleJSON contains the JSON metadata for the struct
// [LanguageResponseSingle]
type languageResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LanguageResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LanguageResponseSingleError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    languageResponseSingleErrorJSON `json:"-"`
}

// languageResponseSingleErrorJSON contains the JSON metadata for the struct
// [LanguageResponseSingleError]
type languageResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LanguageResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LanguageResponseSingleMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    languageResponseSingleMessageJSON `json:"-"`
}

// languageResponseSingleMessageJSON contains the JSON metadata for the struct
// [LanguageResponseSingleMessage]
type languageResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LanguageResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LanguageResponseSingleSuccess bool

const (
	LanguageResponseSingleSuccessTrue LanguageResponseSingleSuccess = true
)

type AccountStreamCaptionDeleteResponse struct {
	Errors     []AccountStreamCaptionDeleteResponseError    `json:"errors"`
	Messages   []AccountStreamCaptionDeleteResponseMessage  `json:"messages"`
	Result     string                                       `json:"result"`
	ResultInfo AccountStreamCaptionDeleteResponseResultInfo `json:"result_info"`
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
	ResultInfo  apijson.Field
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

type AccountStreamCaptionDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       accountStreamCaptionDeleteResponseResultInfoJSON `json:"-"`
}

// accountStreamCaptionDeleteResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountStreamCaptionDeleteResponseResultInfo]
type accountStreamCaptionDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamCaptionDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamCaptionDeleteResponseSuccess bool

const (
	AccountStreamCaptionDeleteResponseSuccessTrue AccountStreamCaptionDeleteResponseSuccess = true
)

type AccountStreamCaptionUpdateParams struct {
	// The WebVTT file containing the caption or subtitle content.
	File param.Field[string] `json:"file,required"`
}

func (r AccountStreamCaptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
