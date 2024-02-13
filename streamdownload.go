// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// StreamDownloadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStreamDownloadService] method
// instead.
type StreamDownloadService struct {
	Options []option.RequestOption
}

// NewStreamDownloadService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamDownloadService(opts ...option.RequestOption) (r *StreamDownloadService) {
	r = &StreamDownloadService{}
	r.Options = opts
	return
}

// Delete the downloads for a video.
func (r *StreamDownloadService) Delete(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *StreamDownloadDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamDownloadDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a download for a video when a video is ready to view.
func (r *StreamDownloadService) StreamMP4DownloadsNewDownloads(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *StreamDownloadStreamMP4DownloadsNewDownloadsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the downloads created for a video.
func (r *StreamDownloadService) StreamMP4DownloadsListDownloads(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *StreamDownloadStreamMP4DownloadsListDownloadsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [StreamDownloadDeleteResponseUnknown] or
// [shared.UnionString].
type StreamDownloadDeleteResponse interface {
	ImplementsStreamDownloadDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamDownloadDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StreamDownloadStreamMP4DownloadsNewDownloadsResponseUnknown]
// or [shared.UnionString].
type StreamDownloadStreamMP4DownloadsNewDownloadsResponse interface {
	ImplementsStreamDownloadStreamMP4DownloadsNewDownloadsResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamDownloadStreamMP4DownloadsNewDownloadsResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [StreamDownloadStreamMP4DownloadsListDownloadsResponseUnknown] or
// [shared.UnionString].
type StreamDownloadStreamMP4DownloadsListDownloadsResponse interface {
	ImplementsStreamDownloadStreamMP4DownloadsListDownloadsResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamDownloadStreamMP4DownloadsListDownloadsResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StreamDownloadDeleteResponseEnvelope struct {
	Errors   []StreamDownloadDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamDownloadDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamDownloadDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamDownloadDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamDownloadDeleteResponseEnvelopeJSON    `json:"-"`
}

// streamDownloadDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamDownloadDeleteResponseEnvelope]
type streamDownloadDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDownloadDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    streamDownloadDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// streamDownloadDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StreamDownloadDeleteResponseEnvelopeErrors]
type streamDownloadDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDownloadDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    streamDownloadDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// streamDownloadDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamDownloadDeleteResponseEnvelopeMessages]
type streamDownloadDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamDownloadDeleteResponseEnvelopeSuccess bool

const (
	StreamDownloadDeleteResponseEnvelopeSuccessTrue StreamDownloadDeleteResponseEnvelopeSuccess = true
)

type StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelope struct {
	Errors   []StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamDownloadStreamMP4DownloadsNewDownloadsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeJSON    `json:"-"`
}

// streamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelope]
type streamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    streamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeErrorsJSON `json:"-"`
}

// streamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeErrors]
type streamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    streamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeMessagesJSON `json:"-"`
}

// streamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeMessages]
type streamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeSuccess bool

const (
	StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeSuccessTrue StreamDownloadStreamMP4DownloadsNewDownloadsResponseEnvelopeSuccess = true
)

type StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelope struct {
	Errors   []StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamDownloadStreamMP4DownloadsListDownloadsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeJSON    `json:"-"`
}

// streamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelope]
type streamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeErrors struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    streamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeErrorsJSON `json:"-"`
}

// streamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeErrors]
type streamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeMessages struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    streamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeMessagesJSON `json:"-"`
}

// streamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeMessages]
type streamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeSuccess bool

const (
	StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeSuccessTrue StreamDownloadStreamMP4DownloadsListDownloadsResponseEnvelopeSuccess = true
)
