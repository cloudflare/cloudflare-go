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

// ZonePurgeCachService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZonePurgeCachService] method
// instead.
type ZonePurgeCachService struct {
	Options []option.RequestOption
}

// NewZonePurgeCachService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZonePurgeCachService(opts ...option.RequestOption) (r *ZonePurgeCachService) {
	r = &ZonePurgeCachService{}
	r.Options = opts
	return
}

// ### Purge All Cached Content
//
// Removes ALL files from Cloudflare's cache. All tiers can purge everything.
//
// ### Purge Cached Content by URL
//
// Granularly removes one or more files from Cloudflare's cache by specifying URLs.
// All tiers can purge by URL.
//
// To purge files with custom cache keys, include the headers used to compute the
// cache key as in the example. If you have a device type or geo in your cache key,
// you will need to include the CF-Device-Type or CF-IPCountry headers.
//
// **NB:** When including the Origin header, be sure to include the **scheme** and
// **hostname**. The port number can be omitted if it is the default port (80 for
// http, 443 for https), but must be included otherwise.
//
// ### Purge Cached Content by Tag, Host or Prefix
//
// Granularly removes one or more files from Cloudflare's cache either by
// specifying the host, the associated Cache-Tag, or a Prefix. Only Enterprise
// customers are permitted to purge by Tag, Host or Prefix.
//
// **NB:** Cache-Tag, host, and prefix purging each have a rate limit of 30,000
// purge API calls in every 24 hour period. You may purge up to 30 tags, hosts, or
// prefixes in one API call. This rate limit can be raised for customers who need
// to purge at higher volume.
func (r *ZonePurgeCachService) ZonePurge(ctx context.Context, identifier string, body ZonePurgeCachZonePurgeParams, opts ...option.RequestOption) (res *APIResponseSingleIDQZpLXoSu, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/purge_cache", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type APIResponseSingleIDQZpLXoSu struct {
	Errors   []APIResponseSingleIDQZpLXoSuError   `json:"errors"`
	Messages []APIResponseSingleIDQZpLXoSuMessage `json:"messages"`
	Result   APIResponseSingleIDQZpLXoSuResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success APIResponseSingleIDQZpLXoSuSuccess `json:"success"`
	JSON    apiResponseSingleIdqZpLXoSuJSON    `json:"-"`
}

// apiResponseSingleIdqZpLXoSuJSON contains the JSON metadata for the struct
// [APIResponseSingleIDQZpLXoSu]
type apiResponseSingleIdqZpLXoSuJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDQZpLXoSu) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIResponseSingleIDQZpLXoSuError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    apiResponseSingleIdqZpLXoSuErrorJSON `json:"-"`
}

// apiResponseSingleIdqZpLXoSuErrorJSON contains the JSON metadata for the struct
// [APIResponseSingleIDQZpLXoSuError]
type apiResponseSingleIdqZpLXoSuErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDQZpLXoSuError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIResponseSingleIDQZpLXoSuMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    apiResponseSingleIdqZpLXoSuMessageJSON `json:"-"`
}

// apiResponseSingleIdqZpLXoSuMessageJSON contains the JSON metadata for the struct
// [APIResponseSingleIDQZpLXoSuMessage]
type apiResponseSingleIdqZpLXoSuMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDQZpLXoSuMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIResponseSingleIDQZpLXoSuResult struct {
	// Identifier
	ID   string                                `json:"id,required"`
	JSON apiResponseSingleIdqZpLXoSuResultJSON `json:"-"`
}

// apiResponseSingleIdqZpLXoSuResultJSON contains the JSON metadata for the struct
// [APIResponseSingleIDQZpLXoSuResult]
type apiResponseSingleIdqZpLXoSuResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDQZpLXoSuResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type APIResponseSingleIDQZpLXoSuSuccess bool

const (
	APIResponseSingleIDQZpLXoSuSuccessTrue APIResponseSingleIDQZpLXoSuSuccess = true
)

// This interface is a union satisfied by one of the following:
// [ZonePurgeCachZonePurgeParamsFlex], [ZonePurgeCachZonePurgeParamsEverything],
// [ZonePurgeCachZonePurgeParamsFiles].
type ZonePurgeCachZonePurgeParams interface {
	ImplementsZonePurgeCachZonePurgeParams()
}

type ZonePurgeCachZonePurgeParamsFlex struct {
	Hosts    param.Field[[]string] `json:"hosts"`
	Prefixes param.Field[[]string] `json:"prefixes"`
	Tags     param.Field[[]string] `json:"tags"`
}

func (r ZonePurgeCachZonePurgeParamsFlex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZonePurgeCachZonePurgeParamsFlex) ImplementsZonePurgeCachZonePurgeParams() {

}

type ZonePurgeCachZonePurgeParamsEverything struct {
	PurgeEverything param.Field[bool] `json:"purge_everything"`
}

func (r ZonePurgeCachZonePurgeParamsEverything) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZonePurgeCachZonePurgeParamsEverything) ImplementsZonePurgeCachZonePurgeParams() {

}

type ZonePurgeCachZonePurgeParamsFiles struct {
	Files param.Field[[]ZonePurgeCachZonePurgeParamsFilesFile] `json:"files"`
}

func (r ZonePurgeCachZonePurgeParamsFiles) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZonePurgeCachZonePurgeParamsFiles) ImplementsZonePurgeCachZonePurgeParams() {

}

// Satisfied by [shared.UnionString],
// [ZonePurgeCachZonePurgeParamsFilesFilesURLAndHeaders].
type ZonePurgeCachZonePurgeParamsFilesFile interface {
	ImplementsZonePurgeCachZonePurgeParamsFilesFile()
}

type ZonePurgeCachZonePurgeParamsFilesFilesURLAndHeaders struct {
	Headers param.Field[interface{}] `json:"headers"`
	URL     param.Field[string]      `json:"url"`
}

func (r ZonePurgeCachZonePurgeParamsFilesFilesURLAndHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonePurgeCachZonePurgeParamsFilesFilesURLAndHeaders) ImplementsZonePurgeCachZonePurgeParamsFilesFile() {
}
