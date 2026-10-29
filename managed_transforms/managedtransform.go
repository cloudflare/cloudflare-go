// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

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

// ManagedTransformService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManagedTransformService] method instead.
type ManagedTransformService struct {
	Options []option.RequestOption
}

// NewManagedTransformService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewManagedTransformService(opts ...option.RequestOption) (r *ManagedTransformService) {
	r = &ManagedTransformService{}
	r.Options = opts
	return
}

// Fetches a list of all Managed Transforms.
func (r *ManagedTransformService) List(ctx context.Context, query ManagedTransformListParams, opts ...option.RequestOption) (res *ManagedTransformListResponse, err error) {
	var env ManagedTransformListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/managed_headers", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Disables all Managed Transforms.
func (r *ManagedTransformService) Delete(ctx context.Context, body ManagedTransformDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/managed_headers", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Updates the status of one or more Managed Transforms.
func (r *ManagedTransformService) Edit(ctx context.Context, params ManagedTransformEditParams, opts ...option.RequestOption) (res *ManagedTransformEditResponse, err error) {
	var env ManagedTransformEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/managed_headers", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A result.
type ManagedTransformListResponse struct {
	// This field can have the runtime type of
	// [[]ManagedTransformListResponseObjectManagedRequestHeader].
	ManagedRequestHeaders interface{} `json:"managed_request_headers"`
	// This field can have the runtime type of
	// [[]ManagedTransformListResponseObjectManagedResponseHeader].
	ManagedResponseHeaders interface{}                      `json:"managed_response_headers"`
	JSON                   managedTransformListResponseJSON `json:"-"`
	union                  ManagedTransformListResponseUnion
}

// managedTransformListResponseJSON contains the JSON metadata for the struct
// [ManagedTransformListResponse]
type managedTransformListResponseJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r managedTransformListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ManagedTransformListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ManagedTransformListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ManagedTransformListResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [ManagedTransformListResponseObject],
// [ManagedTransformListResponseObject], [ManagedTransformListResponseObject],
// [ManagedTransformListResponseObject], [ManagedTransformListResponseObject].
func (r ManagedTransformListResponse) AsUnion() ManagedTransformListResponseUnion {
	return r.union
}

// A result.
//
// Union satisfied by [ManagedTransformListResponseObject],
// [ManagedTransformListResponseObject], [ManagedTransformListResponseObject],
// [ManagedTransformListResponseObject] or [ManagedTransformListResponseObject].
type ManagedTransformListResponseUnion interface {
	implementsManagedTransformListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ManagedTransformListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ManagedTransformListResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ManagedTransformListResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ManagedTransformListResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ManagedTransformListResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ManagedTransformListResponseObject{}),
		},
	)
}

type ManagedTransformListResponseObject struct {
	// The list of Managed Request Transforms.
	ManagedRequestHeaders []ManagedTransformListResponseObjectManagedRequestHeader `json:"managed_request_headers"`
	// The list of Managed Response Transforms.
	ManagedResponseHeaders []ManagedTransformListResponseObjectManagedResponseHeader `json:"managed_response_headers"`
	JSON                   managedTransformListResponseObjectJSON                    `json:"-"`
}

// managedTransformListResponseObjectJSON contains the JSON metadata for the struct
// [ManagedTransformListResponseObject]
type managedTransformListResponseObjectJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ManagedTransformListResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r ManagedTransformListResponseObject) implementsManagedTransformListResponse() {}

// A Managed Transform object.
type ManagedTransformListResponseObjectManagedRequestHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID string `json:"id" api:"required"`
	// Whether the Managed Transform is enabled.
	Enabled bool `json:"enabled" api:"required"`
	// Whether the Managed Transform conflicts with the currently-enabled Managed
	// Transforms.
	HasConflict bool `json:"has_conflict" api:"required"`
	// The Managed Transforms that this Managed Transform conflicts with.
	ConflictsWith []string                                                   `json:"conflicts_with"`
	JSON          managedTransformListResponseObjectManagedRequestHeaderJSON `json:"-"`
}

// managedTransformListResponseObjectManagedRequestHeaderJSON contains the JSON
// metadata for the struct [ManagedTransformListResponseObjectManagedRequestHeader]
type managedTransformListResponseObjectManagedRequestHeaderJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	HasConflict   apijson.Field
	ConflictsWith apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ManagedTransformListResponseObjectManagedRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseObjectManagedRequestHeaderJSON) RawJSON() string {
	return r.raw
}

// A Managed Transform object.
type ManagedTransformListResponseObjectManagedResponseHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID string `json:"id" api:"required"`
	// Whether the Managed Transform is enabled.
	Enabled bool `json:"enabled" api:"required"`
	// Whether the Managed Transform conflicts with the currently-enabled Managed
	// Transforms.
	HasConflict bool `json:"has_conflict" api:"required"`
	// The Managed Transforms that this Managed Transform conflicts with.
	ConflictsWith []string                                                    `json:"conflicts_with"`
	JSON          managedTransformListResponseObjectManagedResponseHeaderJSON `json:"-"`
}

// managedTransformListResponseObjectManagedResponseHeaderJSON contains the JSON
// metadata for the struct
// [ManagedTransformListResponseObjectManagedResponseHeader]
type managedTransformListResponseObjectManagedResponseHeaderJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	HasConflict   apijson.Field
	ConflictsWith apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ManagedTransformListResponseObjectManagedResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseObjectManagedResponseHeaderJSON) RawJSON() string {
	return r.raw
}

// A result.
type ManagedTransformEditResponse struct {
	// This field can have the runtime type of
	// [[]ManagedTransformEditResponseObjectManagedRequestHeader].
	ManagedRequestHeaders interface{} `json:"managed_request_headers"`
	// This field can have the runtime type of
	// [[]ManagedTransformEditResponseObjectManagedResponseHeader].
	ManagedResponseHeaders interface{}                      `json:"managed_response_headers"`
	JSON                   managedTransformEditResponseJSON `json:"-"`
	union                  ManagedTransformEditResponseUnion
}

// managedTransformEditResponseJSON contains the JSON metadata for the struct
// [ManagedTransformEditResponse]
type managedTransformEditResponseJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r managedTransformEditResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ManagedTransformEditResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ManagedTransformEditResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ManagedTransformEditResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [ManagedTransformEditResponseObject],
// [ManagedTransformEditResponseObject], [ManagedTransformEditResponseObject],
// [ManagedTransformEditResponseObject], [ManagedTransformEditResponseObject].
func (r ManagedTransformEditResponse) AsUnion() ManagedTransformEditResponseUnion {
	return r.union
}

// A result.
//
// Union satisfied by [ManagedTransformEditResponseObject],
// [ManagedTransformEditResponseObject], [ManagedTransformEditResponseObject],
// [ManagedTransformEditResponseObject] or [ManagedTransformEditResponseObject].
type ManagedTransformEditResponseUnion interface {
	implementsManagedTransformEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ManagedTransformEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ManagedTransformEditResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ManagedTransformEditResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ManagedTransformEditResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ManagedTransformEditResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ManagedTransformEditResponseObject{}),
		},
	)
}

type ManagedTransformEditResponseObject struct {
	// The list of Managed Request Transforms.
	ManagedRequestHeaders []ManagedTransformEditResponseObjectManagedRequestHeader `json:"managed_request_headers"`
	// The list of Managed Response Transforms.
	ManagedResponseHeaders []ManagedTransformEditResponseObjectManagedResponseHeader `json:"managed_response_headers"`
	JSON                   managedTransformEditResponseObjectJSON                    `json:"-"`
}

// managedTransformEditResponseObjectJSON contains the JSON metadata for the struct
// [ManagedTransformEditResponseObject]
type managedTransformEditResponseObjectJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ManagedTransformEditResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r ManagedTransformEditResponseObject) implementsManagedTransformEditResponse() {}

// A Managed Transform object.
type ManagedTransformEditResponseObjectManagedRequestHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID string `json:"id" api:"required"`
	// Whether the Managed Transform is enabled.
	Enabled bool `json:"enabled" api:"required"`
	// Whether the Managed Transform conflicts with the currently-enabled Managed
	// Transforms.
	HasConflict bool `json:"has_conflict" api:"required"`
	// The Managed Transforms that this Managed Transform conflicts with.
	ConflictsWith []string                                                   `json:"conflicts_with"`
	JSON          managedTransformEditResponseObjectManagedRequestHeaderJSON `json:"-"`
}

// managedTransformEditResponseObjectManagedRequestHeaderJSON contains the JSON
// metadata for the struct [ManagedTransformEditResponseObjectManagedRequestHeader]
type managedTransformEditResponseObjectManagedRequestHeaderJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	HasConflict   apijson.Field
	ConflictsWith apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ManagedTransformEditResponseObjectManagedRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseObjectManagedRequestHeaderJSON) RawJSON() string {
	return r.raw
}

// A Managed Transform object.
type ManagedTransformEditResponseObjectManagedResponseHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID string `json:"id" api:"required"`
	// Whether the Managed Transform is enabled.
	Enabled bool `json:"enabled" api:"required"`
	// Whether the Managed Transform conflicts with the currently-enabled Managed
	// Transforms.
	HasConflict bool `json:"has_conflict" api:"required"`
	// The Managed Transforms that this Managed Transform conflicts with.
	ConflictsWith []string                                                    `json:"conflicts_with"`
	JSON          managedTransformEditResponseObjectManagedResponseHeaderJSON `json:"-"`
}

// managedTransformEditResponseObjectManagedResponseHeaderJSON contains the JSON
// metadata for the struct
// [ManagedTransformEditResponseObjectManagedResponseHeader]
type managedTransformEditResponseObjectManagedResponseHeaderJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	HasConflict   apijson.Field
	ConflictsWith apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ManagedTransformEditResponseObjectManagedResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseObjectManagedResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type ManagedTransformListParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

// A response object.
type ManagedTransformListResponseEnvelope struct {
	// A list of error messages.
	Errors []ManagedTransformListResponseEnvelopeErrors `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []ManagedTransformListResponseEnvelopeMessages `json:"messages" api:"required"`
	// A result.
	Result ManagedTransformListResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ManagedTransformListResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    managedTransformListResponseEnvelopeJSON    `json:"-"`
}

// managedTransformListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ManagedTransformListResponseEnvelope]
type managedTransformListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type ManagedTransformListResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ManagedTransformListResponseEnvelopeErrorsSource `json:"source"`
	JSON   managedTransformListResponseEnvelopeErrorsJSON   `json:"-"`
}

// managedTransformListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ManagedTransformListResponseEnvelopeErrors]
type managedTransformListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ManagedTransformListResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                               `json:"pointer" api:"required"`
	JSON    managedTransformListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// managedTransformListResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ManagedTransformListResponseEnvelopeErrorsSource]
type managedTransformListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type ManagedTransformListResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ManagedTransformListResponseEnvelopeMessagesSource `json:"source"`
	JSON   managedTransformListResponseEnvelopeMessagesJSON   `json:"-"`
}

// managedTransformListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ManagedTransformListResponseEnvelopeMessages]
type managedTransformListResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ManagedTransformListResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                 `json:"pointer" api:"required"`
	JSON    managedTransformListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// managedTransformListResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ManagedTransformListResponseEnvelopeMessagesSource]
type managedTransformListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ManagedTransformListResponseEnvelopeSuccess bool

const (
	ManagedTransformListResponseEnvelopeSuccessTrue ManagedTransformListResponseEnvelopeSuccess = true
)

func (r ManagedTransformListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ManagedTransformListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ManagedTransformDeleteParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type ManagedTransformEditParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// The list of Managed Request Transforms.
	ManagedRequestHeaders param.Field[[]ManagedTransformEditParamsManagedRequestHeader] `json:"managed_request_headers"`
	// The list of Managed Response Transforms.
	ManagedResponseHeaders param.Field[[]ManagedTransformEditParamsManagedResponseHeader] `json:"managed_response_headers"`
}

func (r ManagedTransformEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A Managed Transform object.
type ManagedTransformEditParamsManagedRequestHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID param.Field[string] `json:"id" api:"required"`
	// Whether the Managed Transform is enabled.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r ManagedTransformEditParamsManagedRequestHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A Managed Transform object.
type ManagedTransformEditParamsManagedResponseHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID param.Field[string] `json:"id" api:"required"`
	// Whether the Managed Transform is enabled.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r ManagedTransformEditParamsManagedResponseHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A response object.
type ManagedTransformEditResponseEnvelope struct {
	// A list of error messages.
	Errors []ManagedTransformEditResponseEnvelopeErrors `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []ManagedTransformEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// A result.
	Result ManagedTransformEditResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ManagedTransformEditResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    managedTransformEditResponseEnvelopeJSON    `json:"-"`
}

// managedTransformEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ManagedTransformEditResponseEnvelope]
type managedTransformEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type ManagedTransformEditResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ManagedTransformEditResponseEnvelopeErrorsSource `json:"source"`
	JSON   managedTransformEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// managedTransformEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ManagedTransformEditResponseEnvelopeErrors]
type managedTransformEditResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ManagedTransformEditResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                               `json:"pointer" api:"required"`
	JSON    managedTransformEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// managedTransformEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ManagedTransformEditResponseEnvelopeErrorsSource]
type managedTransformEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type ManagedTransformEditResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ManagedTransformEditResponseEnvelopeMessagesSource `json:"source"`
	JSON   managedTransformEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// managedTransformEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ManagedTransformEditResponseEnvelopeMessages]
type managedTransformEditResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ManagedTransformEditResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                 `json:"pointer" api:"required"`
	JSON    managedTransformEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// managedTransformEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ManagedTransformEditResponseEnvelopeMessagesSource]
type managedTransformEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ManagedTransformEditResponseEnvelopeSuccess bool

const (
	ManagedTransformEditResponseEnvelopeSuccessTrue ManagedTransformEditResponseEnvelopeSuccess = true
)

func (r ManagedTransformEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ManagedTransformEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
