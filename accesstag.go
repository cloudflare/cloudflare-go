// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccessTagService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccessTagService] method instead.
type AccessTagService struct {
	Options []option.RequestOption
}

// NewAccessTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessTagService(opts ...option.RequestOption) (r *AccessTagService) {
	r = &AccessTagService{}
	r.Options = opts
	return
}

// List tags
func (r *AccessTagService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *[]AccessTagListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessTagListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A tag
type AccessTagListResponse struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64                     `json:"app_count"`
	CreatedAt time.Time                 `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      accessTagListResponseJSON `json:"-"`
}

// accessTagListResponseJSON contains the JSON metadata for the struct
// [AccessTagListResponse]
type accessTagListResponseJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagListResponseEnvelope struct {
	Errors     []AccessTagListResponseEnvelopeErrors   `json:"errors"`
	Messages   []AccessTagListResponseEnvelopeMessages `json:"messages"`
	Result     []AccessTagListResponse                 `json:"result"`
	ResultInfo AccessTagListResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessTagListResponseEnvelopeSuccess `json:"success"`
	JSON    accessTagListResponseEnvelopeJSON    `json:"-"`
}

// accessTagListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagListResponseEnvelope]
type accessTagListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accessTagListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessTagListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessTagListResponseEnvelopeErrors]
type accessTagListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessTagListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessTagListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessTagListResponseEnvelopeMessages]
type accessTagListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTagListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       accessTagListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessTagListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [AccessTagListResponseEnvelopeResultInfo]
type accessTagListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessTagListResponseEnvelopeSuccess bool

const (
	AccessTagListResponseEnvelopeSuccessTrue AccessTagListResponseEnvelopeSuccess = true
)
