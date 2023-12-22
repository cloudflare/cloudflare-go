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
func (r *AccountStreamDownloadService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *DeletedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a download for a video when a video is ready to view.
func (r *AccountStreamDownloadService) StreamMP4DownloadsNewDownloads(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *DownloadsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Lists the downloads created for a video.
func (r *AccountStreamDownloadService) StreamMP4DownloadsListDownloads(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *DownloadsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DownloadsResponse struct {
	Errors   []DownloadsResponseError   `json:"errors"`
	Messages []DownloadsResponseMessage `json:"messages"`
	Result   interface{}                `json:"result"`
	// Whether the API call was successful
	Success DownloadsResponseSuccess `json:"success"`
	JSON    downloadsResponseJSON    `json:"-"`
}

// downloadsResponseJSON contains the JSON metadata for the struct
// [DownloadsResponse]
type downloadsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DownloadsResponseError struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    downloadsResponseErrorJSON `json:"-"`
}

// downloadsResponseErrorJSON contains the JSON metadata for the struct
// [DownloadsResponseError]
type downloadsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DownloadsResponseMessage struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    downloadsResponseMessageJSON `json:"-"`
}

// downloadsResponseMessageJSON contains the JSON metadata for the struct
// [DownloadsResponseMessage]
type downloadsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DownloadsResponseSuccess bool

const (
	DownloadsResponseSuccessTrue DownloadsResponseSuccess = true
)
