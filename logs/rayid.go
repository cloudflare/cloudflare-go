// File generated from our OpenAPI spec by Stainless.

package logs

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// RayidService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRayidService] method instead.
type RayidService struct {
	Options []option.RequestOption
}

// NewRayidService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRayidService(opts ...option.RequestOption) (r *RayidService) {
	r = &RayidService{}
	r.Options = opts
	return
}

// The `/rayids` api route allows lookups by specific rayid. The rayids route will
// return zero, one, or more records (ray ids are not unique).
func (r *RayidService) Get(ctx context.Context, zoneIdentifier string, rayIdentifier string, query RayidGetParams, opts ...option.RequestOption) (res *RayidGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logs/rayids/%s", zoneIdentifier, rayIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Union satisfied by [shared.UnionString] or [logs.RayidGetResponseUnknown].
type RayidGetResponse interface {
	ImplementsLogsRayidGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RayidGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RayidGetParams struct {
	// The `/received` route by default returns a limited set of fields, and allows
	// customers to override the default field set by specifying individual fields. The
	// reasons for this are: 1. Most customers require only a small subset of fields,
	// but that subset varies from customer to customer; 2. Flat schema is much easier
	// to work with downstream (importing into BigTable etc); 3. Performance (time to
	// process, file size). If `?fields=` is not specified, default field set is
	// returned. This default field set may change at any time. When `?fields=` is
	// provided, each record is returned with the specified fields. `fields` must be
	// specified as a comma separated list without any whitespaces, and all fields must
	// exist. The order in which fields are specified does not matter, and the order of
	// fields in the response is not specified.
	Fields param.Field[string] `query:"fields"`
	// By default, timestamps in responses are returned as Unix nanosecond integers.
	// The `?timestamps=` argument can be set to change the format in which response
	// timestamps are returned. Possible values are: `unix`, `unixnano`, `rfc3339`.
	// Note that `unix` and `unixnano` return timestamps as integers; `rfc3339` returns
	// timestamps as strings.
	Timestamps param.Field[RayidGetParamsTimestamps] `query:"timestamps"`
}

// URLQuery serializes [RayidGetParams]'s query parameters as `url.Values`.
func (r RayidGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// By default, timestamps in responses are returned as Unix nanosecond integers.
// The `?timestamps=` argument can be set to change the format in which response
// timestamps are returned. Possible values are: `unix`, `unixnano`, `rfc3339`.
// Note that `unix` and `unixnano` return timestamps as integers; `rfc3339` returns
// timestamps as strings.
type RayidGetParamsTimestamps string

const (
	RayidGetParamsTimestampsUnix     RayidGetParamsTimestamps = "unix"
	RayidGetParamsTimestampsUnixnano RayidGetParamsTimestamps = "unixnano"
	RayidGetParamsTimestampsRfc3339  RayidGetParamsTimestamps = "rfc3339"
)
