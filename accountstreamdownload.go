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

// AccountStreamDownloadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamDownloadService]
// method instead.
type AccountStreamDownloadService struct {
	Options []option.RequestOption
}

// NewAccountStreamDownloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamDownloadService(opts ...option.RequestOption) (r *AccountStreamDownloadService) {
	r = &AccountStreamDownloadService{}
	r.Options = opts
	return
}

// Delete the downloads for a video.
func (r *AccountStreamDownloadService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountStreamDownloadDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a download for a video when a video is ready to view.
func (r *AccountStreamDownloadService) StreamMP4DownloadsNewDownloads(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Lists the downloads created for a video.
func (r *AccountStreamDownloadService) StreamMP4DownloadsListDownloads(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountStreamDownloadStreamMP4DownloadsListDownloadsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountStreamDownloadDeleteResponse struct {
	Errors   []AccountStreamDownloadDeleteResponseError   `json:"errors"`
	Messages []AccountStreamDownloadDeleteResponseMessage `json:"messages"`
	Result   string                                       `json:"result"`
	// Whether the API call was successful
	Success AccountStreamDownloadDeleteResponseSuccess `json:"success"`
	JSON    accountStreamDownloadDeleteResponseJSON    `json:"-"`
}

// accountStreamDownloadDeleteResponseJSON contains the JSON metadata for the
// struct [AccountStreamDownloadDeleteResponse]
type accountStreamDownloadDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamDownloadDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamDownloadDeleteResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountStreamDownloadDeleteResponseErrorJSON `json:"-"`
}

// accountStreamDownloadDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountStreamDownloadDeleteResponseError]
type accountStreamDownloadDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamDownloadDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamDownloadDeleteResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountStreamDownloadDeleteResponseMessageJSON `json:"-"`
}

// accountStreamDownloadDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountStreamDownloadDeleteResponseMessage]
type accountStreamDownloadDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamDownloadDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamDownloadDeleteResponseSuccess bool

const (
	AccountStreamDownloadDeleteResponseSuccessTrue AccountStreamDownloadDeleteResponseSuccess = true
)

type AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponse struct {
	Errors   []AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponseError   `json:"errors"`
	Messages []AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponseMessage `json:"messages"`
	Result   interface{}                                                          `json:"result"`
	// Whether the API call was successful
	Success AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponseSuccess `json:"success"`
	JSON    accountStreamDownloadStreamMP4DownloadsNewDownloadsResponseJSON    `json:"-"`
}

// accountStreamDownloadStreamMP4DownloadsNewDownloadsResponseJSON contains the
// JSON metadata for the struct
// [AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponse]
type accountStreamDownloadStreamMP4DownloadsNewDownloadsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponseError struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    accountStreamDownloadStreamMP4DownloadsNewDownloadsResponseErrorJSON `json:"-"`
}

// accountStreamDownloadStreamMP4DownloadsNewDownloadsResponseErrorJSON contains
// the JSON metadata for the struct
// [AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponseError]
type accountStreamDownloadStreamMP4DownloadsNewDownloadsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponseMessage struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    accountStreamDownloadStreamMP4DownloadsNewDownloadsResponseMessageJSON `json:"-"`
}

// accountStreamDownloadStreamMP4DownloadsNewDownloadsResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponseMessage]
type accountStreamDownloadStreamMP4DownloadsNewDownloadsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponseSuccess bool

const (
	AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponseSuccessTrue AccountStreamDownloadStreamMP4DownloadsNewDownloadsResponseSuccess = true
)

type AccountStreamDownloadStreamMP4DownloadsListDownloadsResponse struct {
	Errors   []AccountStreamDownloadStreamMP4DownloadsListDownloadsResponseError   `json:"errors"`
	Messages []AccountStreamDownloadStreamMP4DownloadsListDownloadsResponseMessage `json:"messages"`
	Result   interface{}                                                           `json:"result"`
	// Whether the API call was successful
	Success AccountStreamDownloadStreamMP4DownloadsListDownloadsResponseSuccess `json:"success"`
	JSON    accountStreamDownloadStreamMP4DownloadsListDownloadsResponseJSON    `json:"-"`
}

// accountStreamDownloadStreamMP4DownloadsListDownloadsResponseJSON contains the
// JSON metadata for the struct
// [AccountStreamDownloadStreamMP4DownloadsListDownloadsResponse]
type accountStreamDownloadStreamMP4DownloadsListDownloadsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamDownloadStreamMP4DownloadsListDownloadsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamDownloadStreamMP4DownloadsListDownloadsResponseError struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountStreamDownloadStreamMP4DownloadsListDownloadsResponseErrorJSON `json:"-"`
}

// accountStreamDownloadStreamMP4DownloadsListDownloadsResponseErrorJSON contains
// the JSON metadata for the struct
// [AccountStreamDownloadStreamMP4DownloadsListDownloadsResponseError]
type accountStreamDownloadStreamMP4DownloadsListDownloadsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamDownloadStreamMP4DownloadsListDownloadsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamDownloadStreamMP4DownloadsListDownloadsResponseMessage struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    accountStreamDownloadStreamMP4DownloadsListDownloadsResponseMessageJSON `json:"-"`
}

// accountStreamDownloadStreamMP4DownloadsListDownloadsResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountStreamDownloadStreamMP4DownloadsListDownloadsResponseMessage]
type accountStreamDownloadStreamMP4DownloadsListDownloadsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamDownloadStreamMP4DownloadsListDownloadsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamDownloadStreamMP4DownloadsListDownloadsResponseSuccess bool

const (
	AccountStreamDownloadStreamMP4DownloadsListDownloadsResponseSuccessTrue AccountStreamDownloadStreamMP4DownloadsListDownloadsResponseSuccess = true
)
