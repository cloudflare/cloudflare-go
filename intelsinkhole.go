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

// IntelSinkholeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIntelSinkholeService] method
// instead.
type IntelSinkholeService struct {
	Options []option.RequestOption
}

// NewIntelSinkholeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntelSinkholeService(opts ...option.RequestOption) (r *IntelSinkholeService) {
	r = &IntelSinkholeService{}
	r.Options = opts
	return
}

// List sinkholes owned by this account
func (r *IntelSinkholeService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]IntelSinkholeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelSinkholeListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/sinkholes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelSinkholeListResponse struct {
	// The unique identifier for the sinkhole
	ID int64 `json:"id"`
	// The account tag that owns this sinkhole
	AccountTag string `json:"account_tag"`
	// The date and time when the sinkhole was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The date and time when the sinkhole was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the sinkhole
	Name string `json:"name"`
	// The name of the R2 bucket to store results
	R2Bucket string `json:"r2_bucket"`
	// The id of the R2 instance
	R2ID string                        `json:"r2_id"`
	JSON intelSinkholeListResponseJSON `json:"-"`
}

// intelSinkholeListResponseJSON contains the JSON metadata for the struct
// [IntelSinkholeListResponse]
type intelSinkholeListResponseJSON struct {
	ID          apijson.Field
	AccountTag  apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	R2Bucket    apijson.Field
	R2ID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelSinkholeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelSinkholeListResponseEnvelope struct {
	Errors   []IntelSinkholeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelSinkholeListResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelSinkholeListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success IntelSinkholeListResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelSinkholeListResponseEnvelopeJSON    `json:"-"`
}

// intelSinkholeListResponseEnvelopeJSON contains the JSON metadata for the struct
// [IntelSinkholeListResponseEnvelope]
type intelSinkholeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelSinkholeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelSinkholeListResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    intelSinkholeListResponseEnvelopeErrorsJSON `json:"-"`
}

// intelSinkholeListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IntelSinkholeListResponseEnvelopeErrors]
type intelSinkholeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelSinkholeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelSinkholeListResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    intelSinkholeListResponseEnvelopeMessagesJSON `json:"-"`
}

// intelSinkholeListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IntelSinkholeListResponseEnvelopeMessages]
type intelSinkholeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelSinkholeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelSinkholeListResponseEnvelopeSuccess bool

const (
	IntelSinkholeListResponseEnvelopeSuccessTrue IntelSinkholeListResponseEnvelopeSuccess = true
)
