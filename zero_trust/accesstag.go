// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// AccessTagService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessTagService] method instead.
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
func (r *AccessTagService) New(ctx context.Context, params AccessTagNewParams, opts ...option.RequestOption) (res *Tag, err error) {
	var env AccessTagNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/tags", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a tag
func (r *AccessTagService) Update(ctx context.Context, tagName string, params AccessTagUpdateParams, opts ...option.RequestOption) (res *Tag, err error) {
	var env AccessTagUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/tags/%s", params.AccountID, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List tags
func (r *AccessTagService) List(ctx context.Context, query AccessTagListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Tag], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/tags", query.AccountID)
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
func (r *AccessTagService) ListAutoPaging(ctx context.Context, query AccessTagListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Tag] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a tag
func (r *AccessTagService) Delete(ctx context.Context, tagName string, body AccessTagDeleteParams, opts ...option.RequestOption) (res *AccessTagDeleteResponse, err error) {
	var env AccessTagDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/tags/%s", body.AccountID, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a tag
func (r *AccessTagService) Get(ctx context.Context, tagName string, query AccessTagGetParams, opts ...option.RequestOption) (res *Tag, err error) {
	var env AccessTagGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/tags/%s", query.AccountID, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A tag
type Tag struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	JSON      tagJSON   `json:"-"`
}

// tagJSON contains the JSON metadata for the struct [Tag]
type tagJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Tag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tagJSON) RawJSON() string {
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
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the tag
	Name param.Field[string] `json:"name"`
}

func (r AccessTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessTagNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessTagNewResponseEnvelopeSuccess `json:"success,required"`
	// A tag
	Result Tag                              `json:"result"`
	JSON   accessTagNewResponseEnvelopeJSON `json:"-"`
}

// accessTagNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagNewResponseEnvelope]
type accessTagNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagNewResponseEnvelopeJSON) RawJSON() string {
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
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the tag
	Name param.Field[string] `json:"name,required"`
}

func (r AccessTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessTagUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessTagUpdateResponseEnvelopeSuccess `json:"success,required"`
	// A tag
	Result Tag                                 `json:"result"`
	JSON   accessTagUpdateResponseEnvelopeJSON `json:"-"`
}

// accessTagUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagUpdateResponseEnvelope]
type accessTagUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagUpdateResponseEnvelopeJSON) RawJSON() string {
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

type AccessTagListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessTagDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessTagDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessTagDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessTagDeleteResponse                `json:"result"`
	JSON    accessTagDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessTagDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagDeleteResponseEnvelope]
type accessTagDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagDeleteResponseEnvelopeJSON) RawJSON() string {
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

type AccessTagGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessTagGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessTagGetResponseEnvelopeSuccess `json:"success,required"`
	// A tag
	Result Tag                              `json:"result"`
	JSON   accessTagGetResponseEnvelopeJSON `json:"-"`
}

// accessTagGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessTagGetResponseEnvelope]
type accessTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessTagGetResponseEnvelopeJSON) RawJSON() string {
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
