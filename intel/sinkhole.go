// File generated from our OpenAPI spec by Stainless.

package intel

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// SinkholeService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSinkholeService] method instead.
type SinkholeService struct {
	Options []option.RequestOption
}

// NewSinkholeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSinkholeService(opts ...option.RequestOption) (r *SinkholeService) {
	r = &SinkholeService{}
	r.Options = opts
	return
}

// List sinkholes owned by this account
func (r *SinkholeService) List(ctx context.Context, query SinkholeListParams, opts ...option.RequestOption) (res *[]SinkholeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SinkholeListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/sinkholes", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SinkholeListResponse struct {
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
	R2ID string                   `json:"r2_id"`
	JSON sinkholeListResponseJSON `json:"-"`
}

// sinkholeListResponseJSON contains the JSON metadata for the struct
// [SinkholeListResponse]
type sinkholeListResponseJSON struct {
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

func (r *SinkholeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeListResponseJSON) RawJSON() string {
	return r.raw
}

type SinkholeListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SinkholeListResponseEnvelope struct {
	Errors   []SinkholeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SinkholeListResponseEnvelopeMessages `json:"messages,required"`
	Result   []SinkholeListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success SinkholeListResponseEnvelopeSuccess `json:"success,required"`
	JSON    sinkholeListResponseEnvelopeJSON    `json:"-"`
}

// sinkholeListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SinkholeListResponseEnvelope]
type sinkholeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SinkholeListResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    sinkholeListResponseEnvelopeErrorsJSON `json:"-"`
}

// sinkholeListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SinkholeListResponseEnvelopeErrors]
type sinkholeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SinkholeListResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    sinkholeListResponseEnvelopeMessagesJSON `json:"-"`
}

// sinkholeListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SinkholeListResponseEnvelopeMessages]
type sinkholeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SinkholeListResponseEnvelopeSuccess bool

const (
	SinkholeListResponseEnvelopeSuccessTrue SinkholeListResponseEnvelopeSuccess = true
)
