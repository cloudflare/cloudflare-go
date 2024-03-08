// File generated from our OpenAPI spec by Stainless.

package images

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// V1KeyService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewV1KeyService] method instead.
type V1KeyService struct {
	Options []option.RequestOption
}

// NewV1KeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1KeyService(opts ...option.RequestOption) (r *V1KeyService) {
	r = &V1KeyService{}
	r.Options = opts
	return
}

// Lists your signing keys. These can be found on your Cloudflare Images dashboard.
func (r *V1KeyService) List(ctx context.Context, query V1KeyListParams, opts ...option.RequestOption) (res *V1KeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V1KeyListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/keys", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type V1KeyListResponse struct {
	Keys []V1KeyListResponseKey `json:"keys"`
	JSON v1KeyListResponseJSON  `json:"-"`
}

// v1KeyListResponseJSON contains the JSON metadata for the struct
// [V1KeyListResponse]
type v1KeyListResponseJSON struct {
	Keys        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyListResponseJSON) RawJSON() string {
	return r.raw
}

type V1KeyListResponseKey struct {
	// Key name.
	Name string `json:"name"`
	// Key value.
	Value string                   `json:"value"`
	JSON  v1KeyListResponseKeyJSON `json:"-"`
}

// v1KeyListResponseKeyJSON contains the JSON metadata for the struct
// [V1KeyListResponseKey]
type v1KeyListResponseKeyJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyListResponseKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyListResponseKeyJSON) RawJSON() string {
	return r.raw
}

type V1KeyListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1KeyListResponseEnvelope struct {
	Errors   []V1KeyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1KeyListResponseEnvelopeMessages `json:"messages,required"`
	Result   V1KeyListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V1KeyListResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1KeyListResponseEnvelopeJSON    `json:"-"`
}

// v1KeyListResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1KeyListResponseEnvelope]
type v1KeyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1KeyListResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    v1KeyListResponseEnvelopeErrorsJSON `json:"-"`
}

// v1KeyListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V1KeyListResponseEnvelopeErrors]
type v1KeyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1KeyListResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    v1KeyListResponseEnvelopeMessagesJSON `json:"-"`
}

// v1KeyListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [V1KeyListResponseEnvelopeMessages]
type v1KeyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1KeyListResponseEnvelopeSuccess bool

const (
	V1KeyListResponseEnvelopeSuccessTrue V1KeyListResponseEnvelopeSuccess = true
)
