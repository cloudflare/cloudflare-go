// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/tidwall/gjson"
)

type ErrorData struct {
	Code    int64         `json:"code"`
	Message string        `json:"message"`
	JSON    errorDataJSON `json:"-"`
}

// errorDataJSON contains the JSON metadata for the struct [ErrorData]
type errorDataJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r errorDataJSON) RawJSON() string {
	return r.raw
}

type ResponseInfo struct {
	Code    int64            `json:"code,required"`
	Message string           `json:"message,required"`
	JSON    responseInfoJSON `json:"-"`
}

// responseInfoJSON contains the JSON metadata for the struct [ResponseInfo]
type responseInfoJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseInfoJSON) RawJSON() string {
	return r.raw
}

type ResponseInfoParam struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r ResponseInfoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                 `json:"enabled,required"`
	JSON    unnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cJSON `json:"-"`
}

// unnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cJSON contains the JSON metadata
// for the struct [UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c]
type unnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Union satisfied by
// [shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnknown] or
// [shared.UnionString].
type UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion interface {
	ImplementsSharedUnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
	)
}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type UnnamedSchemaRefEe1e79edcb234d14c4dd266880f2fd24Param struct {
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r UnnamedSchemaRefEe1e79edcb234d14c4dd266880f2fd24Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
