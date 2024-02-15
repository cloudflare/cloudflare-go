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

// PurgeCachService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPurgeCachService] method instead.
type PurgeCachService struct {
	Options []option.RequestOption
}

// NewPurgeCachService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPurgeCachService(opts ...option.RequestOption) (r *PurgeCachService) {
	r = &PurgeCachService{}
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
// you will need to include the CF-Device-Type or CF-IPCountry headers. If you have
// lang in your cache key, you will need to include the Accept-Language header.
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
func (r *PurgeCachService) ZonePurge(ctx context.Context, identifier string, body PurgeCachZonePurgeParams, opts ...option.RequestOption) (res *PurgeCachZonePurgeResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PurgeCachZonePurgeResponseEnvelope
	path := fmt.Sprintf("zones/%s/purge_cache", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PurgeCachZonePurgeResponse struct {
	// Identifier
	ID   string                         `json:"id,required"`
	JSON purgeCachZonePurgeResponseJSON `json:"-"`
}

// purgeCachZonePurgeResponseJSON contains the JSON metadata for the struct
// [PurgeCachZonePurgeResponse]
type purgeCachZonePurgeResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PurgeCachZonePurgeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This interface is a union satisfied by one of the following:
// [PurgeCachZonePurgeParamsONoTspwWFlex],
// [PurgeCachZonePurgeParamsONoTspwWEverything],
// [PurgeCachZonePurgeParamsONoTspwWFiles].
type PurgeCachZonePurgeParams interface {
	ImplementsPurgeCachZonePurgeParams()
}

type PurgeCachZonePurgeParamsONoTspwWFlex struct {
	Hosts    param.Field[[]string] `json:"hosts"`
	Prefixes param.Field[[]string] `json:"prefixes"`
	Tags     param.Field[[]string] `json:"tags"`
}

func (r PurgeCachZonePurgeParamsONoTspwWFlex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PurgeCachZonePurgeParamsONoTspwWFlex) ImplementsPurgeCachZonePurgeParams() {

}

type PurgeCachZonePurgeParamsONoTspwWEverything struct {
	PurgeEverything param.Field[bool] `json:"purge_everything"`
}

func (r PurgeCachZonePurgeParamsONoTspwWEverything) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PurgeCachZonePurgeParamsONoTspwWEverything) ImplementsPurgeCachZonePurgeParams() {

}

type PurgeCachZonePurgeParamsONoTspwWFiles struct {
	Files param.Field[[]PurgeCachZonePurgeParamsONoTspwWFilesFile] `json:"files"`
}

func (r PurgeCachZonePurgeParamsONoTspwWFiles) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PurgeCachZonePurgeParamsONoTspwWFiles) ImplementsPurgeCachZonePurgeParams() {

}

// Satisfied by [shared.UnionString],
// [PurgeCachZonePurgeParamsONoTspwWFilesFilesONoTspwWURLAndHeaders].
type PurgeCachZonePurgeParamsONoTspwWFilesFile interface {
	ImplementsPurgeCachZonePurgeParamsONoTspwWFilesFile()
}

type PurgeCachZonePurgeParamsONoTspwWFilesFilesONoTspwWURLAndHeaders struct {
	Headers param.Field[interface{}] `json:"headers"`
	URL     param.Field[string]      `json:"url"`
}

func (r PurgeCachZonePurgeParamsONoTspwWFilesFilesONoTspwWURLAndHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PurgeCachZonePurgeParamsONoTspwWFilesFilesONoTspwWURLAndHeaders) ImplementsPurgeCachZonePurgeParamsONoTspwWFilesFile() {
}

type PurgeCachZonePurgeResponseEnvelope struct {
	Errors   []PurgeCachZonePurgeResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PurgeCachZonePurgeResponseEnvelopeMessages `json:"messages,required"`
	Result   PurgeCachZonePurgeResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success PurgeCachZonePurgeResponseEnvelopeSuccess `json:"success,required"`
	JSON    purgeCachZonePurgeResponseEnvelopeJSON    `json:"-"`
}

// purgeCachZonePurgeResponseEnvelopeJSON contains the JSON metadata for the struct
// [PurgeCachZonePurgeResponseEnvelope]
type purgeCachZonePurgeResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PurgeCachZonePurgeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PurgeCachZonePurgeResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    purgeCachZonePurgeResponseEnvelopeErrorsJSON `json:"-"`
}

// purgeCachZonePurgeResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PurgeCachZonePurgeResponseEnvelopeErrors]
type purgeCachZonePurgeResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PurgeCachZonePurgeResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PurgeCachZonePurgeResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    purgeCachZonePurgeResponseEnvelopeMessagesJSON `json:"-"`
}

// purgeCachZonePurgeResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PurgeCachZonePurgeResponseEnvelopeMessages]
type purgeCachZonePurgeResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PurgeCachZonePurgeResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PurgeCachZonePurgeResponseEnvelopeSuccess bool

const (
	PurgeCachZonePurgeResponseEnvelopeSuccessTrue PurgeCachZonePurgeResponseEnvelopeSuccess = true
)
