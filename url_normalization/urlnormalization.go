// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_normalization

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/tidwall/gjson"
)

// URLNormalizationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewURLNormalizationService] method instead.
type URLNormalizationService struct {
	Options []option.RequestOption
}

// NewURLNormalizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewURLNormalizationService(opts ...option.RequestOption) (r *URLNormalizationService) {
	r = &URLNormalizationService{}
	r.Options = opts
	return
}

// Updates the URL Normalization settings.
func (r *URLNormalizationService) Update(ctx context.Context, params URLNormalizationUpdateParams, opts ...option.RequestOption) (res *URLNormalizationUpdateResponse, err error) {
	var env URLNormalizationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/url_normalization", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Deletes the URL Normalization settings.
func (r *URLNormalizationService) Delete(ctx context.Context, body URLNormalizationDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/url_normalization", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Fetches the current URL Normalization settings.
func (r *URLNormalizationService) Get(ctx context.Context, query URLNormalizationGetParams, opts ...option.RequestOption) (res *URLNormalizationGetResponse, err error) {
	var env URLNormalizationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/url_normalization", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A result.
type URLNormalizationUpdateResponse struct {
	// The scope of the URL normalization.
	Scope URLNormalizationUpdateResponseScope `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type  URLNormalizationUpdateResponseType `json:"type"`
	JSON  urlNormalizationUpdateResponseJSON `json:"-"`
	union URLNormalizationUpdateResponseUnion
}

// urlNormalizationUpdateResponseJSON contains the JSON metadata for the struct
// [URLNormalizationUpdateResponse]
type urlNormalizationUpdateResponseJSON struct {
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r urlNormalizationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *URLNormalizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = URLNormalizationUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [URLNormalizationUpdateResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [URLNormalizationUpdateResponseObject],
// [URLNormalizationUpdateResponseObject], [URLNormalizationUpdateResponseObject],
// [URLNormalizationUpdateResponseObject], [URLNormalizationUpdateResponseObject].
func (r URLNormalizationUpdateResponse) AsUnion() URLNormalizationUpdateResponseUnion {
	return r.union
}

// A result.
//
// Union satisfied by [URLNormalizationUpdateResponseObject],
// [URLNormalizationUpdateResponseObject], [URLNormalizationUpdateResponseObject],
// [URLNormalizationUpdateResponseObject] or
// [URLNormalizationUpdateResponseObject].
type URLNormalizationUpdateResponseUnion interface {
	implementsURLNormalizationUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*URLNormalizationUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(URLNormalizationUpdateResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(URLNormalizationUpdateResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(URLNormalizationUpdateResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(URLNormalizationUpdateResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(URLNormalizationUpdateResponseObject{}),
		},
	)
}

type URLNormalizationUpdateResponseObject struct {
	// The scope of the URL normalization.
	Scope URLNormalizationUpdateResponseObjectScope `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type URLNormalizationUpdateResponseObjectType `json:"type"`
	JSON urlNormalizationUpdateResponseObjectJSON `json:"-"`
}

// urlNormalizationUpdateResponseObjectJSON contains the JSON metadata for the
// struct [URLNormalizationUpdateResponseObject]
type urlNormalizationUpdateResponseObjectJSON struct {
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationUpdateResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationUpdateResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r URLNormalizationUpdateResponseObject) implementsURLNormalizationUpdateResponse() {}

// The scope of the URL normalization.
type URLNormalizationUpdateResponseObjectScope string

const (
	URLNormalizationUpdateResponseObjectScopeIncoming URLNormalizationUpdateResponseObjectScope = "incoming"
	URLNormalizationUpdateResponseObjectScopeBoth     URLNormalizationUpdateResponseObjectScope = "both"
	URLNormalizationUpdateResponseObjectScopeNone     URLNormalizationUpdateResponseObjectScope = "none"
)

func (r URLNormalizationUpdateResponseObjectScope) IsKnown() bool {
	switch r {
	case URLNormalizationUpdateResponseObjectScopeIncoming, URLNormalizationUpdateResponseObjectScopeBoth, URLNormalizationUpdateResponseObjectScopeNone:
		return true
	}
	return false
}

// The type of URL normalization performed by Cloudflare.
type URLNormalizationUpdateResponseObjectType string

const (
	URLNormalizationUpdateResponseObjectTypeCloudflare URLNormalizationUpdateResponseObjectType = "cloudflare"
	URLNormalizationUpdateResponseObjectTypeRfc3986    URLNormalizationUpdateResponseObjectType = "rfc3986"
)

func (r URLNormalizationUpdateResponseObjectType) IsKnown() bool {
	switch r {
	case URLNormalizationUpdateResponseObjectTypeCloudflare, URLNormalizationUpdateResponseObjectTypeRfc3986:
		return true
	}
	return false
}

// The scope of the URL normalization.
type URLNormalizationUpdateResponseScope string

const (
	URLNormalizationUpdateResponseScopeIncoming URLNormalizationUpdateResponseScope = "incoming"
	URLNormalizationUpdateResponseScopeBoth     URLNormalizationUpdateResponseScope = "both"
	URLNormalizationUpdateResponseScopeNone     URLNormalizationUpdateResponseScope = "none"
)

func (r URLNormalizationUpdateResponseScope) IsKnown() bool {
	switch r {
	case URLNormalizationUpdateResponseScopeIncoming, URLNormalizationUpdateResponseScopeBoth, URLNormalizationUpdateResponseScopeNone:
		return true
	}
	return false
}

// The type of URL normalization performed by Cloudflare.
type URLNormalizationUpdateResponseType string

const (
	URLNormalizationUpdateResponseTypeCloudflare URLNormalizationUpdateResponseType = "cloudflare"
	URLNormalizationUpdateResponseTypeRfc3986    URLNormalizationUpdateResponseType = "rfc3986"
)

func (r URLNormalizationUpdateResponseType) IsKnown() bool {
	switch r {
	case URLNormalizationUpdateResponseTypeCloudflare, URLNormalizationUpdateResponseTypeRfc3986:
		return true
	}
	return false
}

// A result.
type URLNormalizationGetResponse struct {
	// The scope of the URL normalization.
	Scope URLNormalizationGetResponseScope `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type  URLNormalizationGetResponseType `json:"type"`
	JSON  urlNormalizationGetResponseJSON `json:"-"`
	union URLNormalizationGetResponseUnion
}

// urlNormalizationGetResponseJSON contains the JSON metadata for the struct
// [URLNormalizationGetResponse]
type urlNormalizationGetResponseJSON struct {
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r urlNormalizationGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *URLNormalizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = URLNormalizationGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [URLNormalizationGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [URLNormalizationGetResponseObject],
// [URLNormalizationGetResponseObject], [URLNormalizationGetResponseObject],
// [URLNormalizationGetResponseObject], [URLNormalizationGetResponseObject].
func (r URLNormalizationGetResponse) AsUnion() URLNormalizationGetResponseUnion {
	return r.union
}

// A result.
//
// Union satisfied by [URLNormalizationGetResponseObject],
// [URLNormalizationGetResponseObject], [URLNormalizationGetResponseObject],
// [URLNormalizationGetResponseObject] or [URLNormalizationGetResponseObject].
type URLNormalizationGetResponseUnion interface {
	implementsURLNormalizationGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*URLNormalizationGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(URLNormalizationGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(URLNormalizationGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(URLNormalizationGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(URLNormalizationGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(URLNormalizationGetResponseObject{}),
		},
	)
}

type URLNormalizationGetResponseObject struct {
	// The scope of the URL normalization.
	Scope URLNormalizationGetResponseObjectScope `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type URLNormalizationGetResponseObjectType `json:"type"`
	JSON urlNormalizationGetResponseObjectJSON `json:"-"`
}

// urlNormalizationGetResponseObjectJSON contains the JSON metadata for the struct
// [URLNormalizationGetResponseObject]
type urlNormalizationGetResponseObjectJSON struct {
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationGetResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r URLNormalizationGetResponseObject) implementsURLNormalizationGetResponse() {}

// The scope of the URL normalization.
type URLNormalizationGetResponseObjectScope string

const (
	URLNormalizationGetResponseObjectScopeIncoming URLNormalizationGetResponseObjectScope = "incoming"
	URLNormalizationGetResponseObjectScopeBoth     URLNormalizationGetResponseObjectScope = "both"
	URLNormalizationGetResponseObjectScopeNone     URLNormalizationGetResponseObjectScope = "none"
)

func (r URLNormalizationGetResponseObjectScope) IsKnown() bool {
	switch r {
	case URLNormalizationGetResponseObjectScopeIncoming, URLNormalizationGetResponseObjectScopeBoth, URLNormalizationGetResponseObjectScopeNone:
		return true
	}
	return false
}

// The type of URL normalization performed by Cloudflare.
type URLNormalizationGetResponseObjectType string

const (
	URLNormalizationGetResponseObjectTypeCloudflare URLNormalizationGetResponseObjectType = "cloudflare"
	URLNormalizationGetResponseObjectTypeRfc3986    URLNormalizationGetResponseObjectType = "rfc3986"
)

func (r URLNormalizationGetResponseObjectType) IsKnown() bool {
	switch r {
	case URLNormalizationGetResponseObjectTypeCloudflare, URLNormalizationGetResponseObjectTypeRfc3986:
		return true
	}
	return false
}

// The scope of the URL normalization.
type URLNormalizationGetResponseScope string

const (
	URLNormalizationGetResponseScopeIncoming URLNormalizationGetResponseScope = "incoming"
	URLNormalizationGetResponseScopeBoth     URLNormalizationGetResponseScope = "both"
	URLNormalizationGetResponseScopeNone     URLNormalizationGetResponseScope = "none"
)

func (r URLNormalizationGetResponseScope) IsKnown() bool {
	switch r {
	case URLNormalizationGetResponseScopeIncoming, URLNormalizationGetResponseScopeBoth, URLNormalizationGetResponseScopeNone:
		return true
	}
	return false
}

// The type of URL normalization performed by Cloudflare.
type URLNormalizationGetResponseType string

const (
	URLNormalizationGetResponseTypeCloudflare URLNormalizationGetResponseType = "cloudflare"
	URLNormalizationGetResponseTypeRfc3986    URLNormalizationGetResponseType = "rfc3986"
)

func (r URLNormalizationGetResponseType) IsKnown() bool {
	switch r {
	case URLNormalizationGetResponseTypeCloudflare, URLNormalizationGetResponseTypeRfc3986:
		return true
	}
	return false
}

type URLNormalizationUpdateParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// The scope of the URL normalization.
	Scope param.Field[URLNormalizationUpdateParamsScope] `json:"scope" api:"required"`
	// The type of URL normalization performed by Cloudflare.
	Type param.Field[URLNormalizationUpdateParamsType] `json:"type" api:"required"`
}

func (r URLNormalizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The scope of the URL normalization.
type URLNormalizationUpdateParamsScope string

const (
	URLNormalizationUpdateParamsScopeIncoming URLNormalizationUpdateParamsScope = "incoming"
	URLNormalizationUpdateParamsScopeBoth     URLNormalizationUpdateParamsScope = "both"
	URLNormalizationUpdateParamsScopeNone     URLNormalizationUpdateParamsScope = "none"
)

func (r URLNormalizationUpdateParamsScope) IsKnown() bool {
	switch r {
	case URLNormalizationUpdateParamsScopeIncoming, URLNormalizationUpdateParamsScopeBoth, URLNormalizationUpdateParamsScopeNone:
		return true
	}
	return false
}

// The type of URL normalization performed by Cloudflare.
type URLNormalizationUpdateParamsType string

const (
	URLNormalizationUpdateParamsTypeCloudflare URLNormalizationUpdateParamsType = "cloudflare"
	URLNormalizationUpdateParamsTypeRfc3986    URLNormalizationUpdateParamsType = "rfc3986"
)

func (r URLNormalizationUpdateParamsType) IsKnown() bool {
	switch r {
	case URLNormalizationUpdateParamsTypeCloudflare, URLNormalizationUpdateParamsTypeRfc3986:
		return true
	}
	return false
}

// A response object.
type URLNormalizationUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []URLNormalizationUpdateResponseEnvelopeErrors `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []URLNormalizationUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// A result.
	Result URLNormalizationUpdateResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success URLNormalizationUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    urlNormalizationUpdateResponseEnvelopeJSON    `json:"-"`
}

// urlNormalizationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [URLNormalizationUpdateResponseEnvelope]
type urlNormalizationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type URLNormalizationUpdateResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source URLNormalizationUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON   urlNormalizationUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// urlNormalizationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [URLNormalizationUpdateResponseEnvelopeErrors]
type urlNormalizationUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type URLNormalizationUpdateResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                 `json:"pointer" api:"required"`
	JSON    urlNormalizationUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// urlNormalizationUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [URLNormalizationUpdateResponseEnvelopeErrorsSource]
type urlNormalizationUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type URLNormalizationUpdateResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source URLNormalizationUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON   urlNormalizationUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// urlNormalizationUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [URLNormalizationUpdateResponseEnvelopeMessages]
type urlNormalizationUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type URLNormalizationUpdateResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                   `json:"pointer" api:"required"`
	JSON    urlNormalizationUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// urlNormalizationUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [URLNormalizationUpdateResponseEnvelopeMessagesSource]
type urlNormalizationUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type URLNormalizationUpdateResponseEnvelopeSuccess bool

const (
	URLNormalizationUpdateResponseEnvelopeSuccessTrue URLNormalizationUpdateResponseEnvelopeSuccess = true
)

func (r URLNormalizationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case URLNormalizationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type URLNormalizationDeleteParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type URLNormalizationGetParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

// A response object.
type URLNormalizationGetResponseEnvelope struct {
	// A list of error messages.
	Errors []URLNormalizationGetResponseEnvelopeErrors `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []URLNormalizationGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// A result.
	Result URLNormalizationGetResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success URLNormalizationGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    urlNormalizationGetResponseEnvelopeJSON    `json:"-"`
}

// urlNormalizationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [URLNormalizationGetResponseEnvelope]
type urlNormalizationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type URLNormalizationGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source URLNormalizationGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   urlNormalizationGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// urlNormalizationGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [URLNormalizationGetResponseEnvelopeErrors]
type urlNormalizationGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type URLNormalizationGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                              `json:"pointer" api:"required"`
	JSON    urlNormalizationGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// urlNormalizationGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [URLNormalizationGetResponseEnvelopeErrorsSource]
type urlNormalizationGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type URLNormalizationGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source URLNormalizationGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   urlNormalizationGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// urlNormalizationGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [URLNormalizationGetResponseEnvelopeMessages]
type urlNormalizationGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type URLNormalizationGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                `json:"pointer" api:"required"`
	JSON    urlNormalizationGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// urlNormalizationGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [URLNormalizationGetResponseEnvelopeMessagesSource]
type urlNormalizationGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type URLNormalizationGetResponseEnvelopeSuccess bool

const (
	URLNormalizationGetResponseEnvelopeSuccessTrue URLNormalizationGetResponseEnvelopeSuccess = true
)

func (r URLNormalizationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case URLNormalizationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
