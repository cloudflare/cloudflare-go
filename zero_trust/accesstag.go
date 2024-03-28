// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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

// Create a tag
func (r *AccessTagService) New(ctx context.Context, identifier string, body AccessTagNewParams, opts ...option.RequestOption) (res *ZeroTrustTag, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessTagNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a tag
func (r *AccessTagService) Update(ctx context.Context, identifier string, tagName string, body AccessTagUpdateParams, opts ...option.RequestOption) (res *ZeroTrustTag, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessTagUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags/%s", identifier, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List tags
func (r *AccessTagService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *shared.SinglePage[ZeroTrustTag], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/access/tags", identifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List tags
func (r *AccessTagService) ListAutoPaging(ctx context.Context, identifier string, opts ...option.RequestOption) *shared.SinglePageAutoPager[ZeroTrustTag] {
	return shared.NewSinglePageAutoPager(r.List(ctx, identifier, opts...))
}

// Delete a tag
func (r *AccessTagService) Delete(ctx context.Context, identifier string, name string, opts ...option.RequestOption) (res *AccessTagDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessTagDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags/%s", identifier, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a tag
func (r *AccessTagService) Get(ctx context.Context, identifier string, name string, opts ...option.RequestOption) (res *ZeroTrustTag, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessTagGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/tags/%s", identifier, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A tag
type ZeroTrustTag struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64            `json:"app_count"`
	CreatedAt time.Time        `json:"created_at" format:"date-time"`
	UpdatedAt time.Time        `json:"updated_at" format:"date-time"`
	JSON      zeroTrustTagJSON `json:"-"`
}

// zeroTrustTagJSON contains the JSON metadata for the struct [ZeroTrustTag]
type zeroTrustTagJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustTagJSON) RawJSON() string {
	return r.raw
}

type AccessTagDeleteResponse struct {
	// The name of the tag
	Name string                      `json:"name"`
	JSON accessTagDeleteResponseJSON `json:"-"`
}

// accessTagDeleteResponseJSON contains the JSON metadata for the struct
// [AccessTagDeleteResponse]
type accessTagDeleteResponseJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessTagNewParams struct {
	// The name of the tag
	Name param.Field[string] `json:"name,required"`
}

func (r AccessTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessTagNewResponseEnvelope struct {
	Errors   []AccessTagNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessTagNewResponseEnvelopeMessages `json:"messages,required"`
	// A tag
	Result ZeroTrustTag `json:"result,required"`
	// Whether the API call was successful
	Success AccessTagNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessTagNewResponseEnvelopeJSON    `json:"-"`
}

// accessTagNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagNewResponseEnvelope]
type accessTagNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessTagNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accessTagNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessTagNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AccessTagNewResponseEnvelopeErrors]
type accessTagNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessTagNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessTagNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessTagNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessTagNewResponseEnvelopeMessages]
type accessTagNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessTagNewResponseEnvelopeSuccess bool

const (
	AccessTagNewResponseEnvelopeSuccessTrue AccessTagNewResponseEnvelopeSuccess = true
)

func (r AccessTagNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessTagNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessTagUpdateParams struct {
	// The name of the tag
	Name param.Field[string] `json:"name,required"`
}

func (r AccessTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessTagUpdateResponseEnvelope struct {
	Errors   []AccessTagUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessTagUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A tag
	Result ZeroTrustTag `json:"result,required"`
	// Whether the API call was successful
	Success AccessTagUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessTagUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessTagUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagUpdateResponseEnvelope]
type accessTagUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessTagUpdateResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessTagUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessTagUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessTagUpdateResponseEnvelopeErrors]
type accessTagUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessTagUpdateResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessTagUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessTagUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessTagUpdateResponseEnvelopeMessages]
type accessTagUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessTagUpdateResponseEnvelopeSuccess bool

const (
	AccessTagUpdateResponseEnvelopeSuccessTrue AccessTagUpdateResponseEnvelopeSuccess = true
)

func (r AccessTagUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessTagUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessTagDeleteResponseEnvelope struct {
	Errors   []AccessTagDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessTagDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessTagDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessTagDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessTagDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessTagDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagDeleteResponseEnvelope]
type accessTagDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessTagDeleteResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessTagDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessTagDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessTagDeleteResponseEnvelopeErrors]
type accessTagDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessTagDeleteResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessTagDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessTagDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessTagDeleteResponseEnvelopeMessages]
type accessTagDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessTagDeleteResponseEnvelopeSuccess bool

const (
	AccessTagDeleteResponseEnvelopeSuccessTrue AccessTagDeleteResponseEnvelopeSuccess = true
)

func (r AccessTagDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessTagDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessTagGetResponseEnvelope struct {
	Errors   []AccessTagGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessTagGetResponseEnvelopeMessages `json:"messages,required"`
	// A tag
	Result ZeroTrustTag `json:"result,required"`
	// Whether the API call was successful
	Success AccessTagGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessTagGetResponseEnvelopeJSON    `json:"-"`
}

// accessTagGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagGetResponseEnvelope]
type accessTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessTagGetResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accessTagGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessTagGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AccessTagGetResponseEnvelopeErrors]
type accessTagGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessTagGetResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessTagGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessTagGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessTagGetResponseEnvelopeMessages]
type accessTagGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessTagGetResponseEnvelopeSuccess bool

const (
	AccessTagGetResponseEnvelopeSuccessTrue AccessTagGetResponseEnvelopeSuccess = true
)

func (r AccessTagGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessTagGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
