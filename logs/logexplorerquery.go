// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// LogExplorerQueryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogExplorerQueryService] method instead.
type LogExplorerQueryService struct {
	Options []option.RequestOption
}

// NewLogExplorerQueryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogExplorerQueryService(opts ...option.RequestOption) (r *LogExplorerQueryService) {
	r = &LogExplorerQueryService{}
	r.Options = opts
	return
}

// Run a SQL query against account or zone-level datasets.
//
// Timestamp fields are RFC3339 strings. Filter with: WHERE {timestamp_field} >=
// now() - INTERVAL '30' DAY WHERE {timestamp_field} >= '2026-04-01T00:00:00Z'
// WHERE {timestamp_field} BETWEEN '2026-04-01T00:00:00Z' AND
// '2026-04-30T23:59:59Z'
//
// Check /account or zones/{account or zone_id}/logs/explorer/datasets to see
// enabled account or zone level datasets. Zone-level datasets will not appear
// here. Check /account or zones/{account or
// zone_id}/logs/explorer/datasets/available for the schemas, and the name of the
// timestamp fields.
//
// For zone-level datasets use the zone-scoped endpoint: POST
// /zones/{zone_id}/logs/explorer/query/sql
//
// For more information about the datasets, and the meaning of each field, check
// out https://developers.cloudflare.com/logs/logpush/logpush-job/datasets/
func (r *LogExplorerQueryService) Sql(ctx context.Context, body io.Reader, params LogExplorerQuerySqlParams, opts ...option.RequestOption) (res *pagination.SinglePage[LogExplorerQuerySqlResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithRequestBody("text/plain", body), option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logs/explorer/query/sql", accountOrZone, accountOrZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, nil, &res, opts...)
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

// Run a SQL query against account or zone-level datasets.
//
// Timestamp fields are RFC3339 strings. Filter with: WHERE {timestamp_field} >=
// now() - INTERVAL '30' DAY WHERE {timestamp_field} >= '2026-04-01T00:00:00Z'
// WHERE {timestamp_field} BETWEEN '2026-04-01T00:00:00Z' AND
// '2026-04-30T23:59:59Z'
//
// Check /account or zones/{account or zone_id}/logs/explorer/datasets to see
// enabled account or zone level datasets. Zone-level datasets will not appear
// here. Check /account or zones/{account or
// zone_id}/logs/explorer/datasets/available for the schemas, and the name of the
// timestamp fields.
//
// For zone-level datasets use the zone-scoped endpoint: POST
// /zones/{zone_id}/logs/explorer/query/sql
//
// For more information about the datasets, and the meaning of each field, check
// out https://developers.cloudflare.com/logs/logpush/logpush-job/datasets/
func (r *LogExplorerQueryService) SqlAutoPaging(ctx context.Context, body io.Reader, params LogExplorerQuerySqlParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LogExplorerQuerySqlResponse] {
	return pagination.NewSinglePageAutoPager(r.Sql(ctx, body, params, opts...))
}

type LogExplorerQuerySqlResponse map[string]interface{}

type LogExplorerQuerySqlParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r LogExplorerQuerySqlParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
