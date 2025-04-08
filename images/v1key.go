// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package images

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// V1KeyService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1KeyService] method instead.
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

// Create a new signing key with specified name. Returns all keys available.
func (r *V1KeyService) Update(ctx context.Context, signingKeyName string, body V1KeyUpdateParams, opts ...option.RequestOption) (res *V1KeyUpdateResponse, err error) {
	var env V1KeyUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if signingKeyName == "" {
		err = errors.New("missing required signing_key_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/keys/%s", body.AccountID, signingKeyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists your signing keys. These can be found on your Cloudflare Images dashboard.
func (r *V1KeyService) List(ctx context.Context, query V1KeyListParams, opts ...option.RequestOption) (res *V1KeyListResponse, err error) {
	var env V1KeyListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/keys", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete signing key with specified name. Returns all keys available. When last
// key is removed, a new default signing key will be generated.
func (r *V1KeyService) Delete(ctx context.Context, signingKeyName string, body V1KeyDeleteParams, opts ...option.RequestOption) (res *V1KeyDeleteResponse, err error) {
	var env V1KeyDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if signingKeyName == "" {
		err = errors.New("missing required signing_key_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/keys/%s", body.AccountID, signingKeyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Key struct {
	// Key name.
	Name string `json:"name"`
	// Key value.
	Value string  `json:"value"`
	JSON  keyJSON `json:"-"`
}

// keyJSON contains the JSON metadata for the struct [Key]
type keyJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Key) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyJSON) RawJSON() string {
	return r.raw
}

type V1KeyUpdateResponse struct {
	Keys []Key                   `json:"keys"`
	JSON v1KeyUpdateResponseJSON `json:"-"`
}

// v1KeyUpdateResponseJSON contains the JSON metadata for the struct
// [V1KeyUpdateResponse]
type v1KeyUpdateResponseJSON struct {
	Keys        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type V1KeyListResponse struct {
	Keys []Key                 `json:"keys"`
	JSON v1KeyListResponseJSON `json:"-"`
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

type V1KeyDeleteResponse struct {
	Keys []Key                   `json:"keys"`
	JSON v1KeyDeleteResponseJSON `json:"-"`
}

// v1KeyDeleteResponseJSON contains the JSON metadata for the struct
// [V1KeyDeleteResponse]
type v1KeyDeleteResponseJSON struct {
	Keys        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type V1KeyUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1KeyUpdateResponseEnvelope struct {
	Errors   []V1KeyUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1KeyUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   V1KeyUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V1KeyUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1KeyUpdateResponseEnvelopeJSON    `json:"-"`
}

// v1KeyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1KeyUpdateResponseEnvelope]
type v1KeyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1KeyUpdateResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    v1KeyUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// v1KeyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V1KeyUpdateResponseEnvelopeErrors]
type v1KeyUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1KeyUpdateResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    v1KeyUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// v1KeyUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V1KeyUpdateResponseEnvelopeMessages]
type v1KeyUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1KeyUpdateResponseEnvelopeSuccess bool

const (
	V1KeyUpdateResponseEnvelopeSuccessTrue V1KeyUpdateResponseEnvelopeSuccess = true
)

func (r V1KeyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1KeyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
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

func (r V1KeyListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1KeyListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V1KeyDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1KeyDeleteResponseEnvelope struct {
	Errors   []V1KeyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1KeyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   V1KeyDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V1KeyDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1KeyDeleteResponseEnvelopeJSON    `json:"-"`
}

// v1KeyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1KeyDeleteResponseEnvelope]
type v1KeyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1KeyDeleteResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    v1KeyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// v1KeyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V1KeyDeleteResponseEnvelopeErrors]
type v1KeyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1KeyDeleteResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    v1KeyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// v1KeyDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V1KeyDeleteResponseEnvelopeMessages]
type v1KeyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1KeyDeleteResponseEnvelopeSuccess bool

const (
	V1KeyDeleteResponseEnvelopeSuccessTrue V1KeyDeleteResponseEnvelopeSuccess = true
)

func (r V1KeyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1KeyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
