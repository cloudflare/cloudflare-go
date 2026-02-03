// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// ToMarkdownService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToMarkdownService] method instead.
type ToMarkdownService struct {
	Options []option.RequestOption
}

// NewToMarkdownService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToMarkdownService(opts ...option.RequestOption) (r *ToMarkdownService) {
	r = &ToMarkdownService{}
	r.Options = opts
	return
}

// Get all converted formats supported
func (r *ToMarkdownService) Supported(ctx context.Context, query ToMarkdownSupportedParams, opts ...option.RequestOption) (res *pagination.SinglePage[ToMarkdownSupportedResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/tomarkdown/supported", query.AccountID)
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

// Get all converted formats supported
func (r *ToMarkdownService) SupportedAutoPaging(ctx context.Context, query ToMarkdownSupportedParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ToMarkdownSupportedResponse] {
	return pagination.NewSinglePageAutoPager(r.Supported(ctx, query, opts...))
}

// Convert Files into Markdown
func (r *ToMarkdownService) Transform(ctx context.Context, params ToMarkdownTransformParams, opts ...option.RequestOption) (res *pagination.SinglePage[ToMarkdownTransformResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/tomarkdown", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Convert Files into Markdown
func (r *ToMarkdownService) TransformAutoPaging(ctx context.Context, params ToMarkdownTransformParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ToMarkdownTransformResponse] {
	return pagination.NewSinglePageAutoPager(r.Transform(ctx, params, opts...))
}

type ToMarkdownSupportedResponse struct {
	Extension string                          `json:"extension,required"`
	MimeType  string                          `json:"mimeType,required"`
	JSON      toMarkdownSupportedResponseJSON `json:"-"`
}

// toMarkdownSupportedResponseJSON contains the JSON metadata for the struct
// [ToMarkdownSupportedResponse]
type toMarkdownSupportedResponseJSON struct {
	Extension   apijson.Field
	MimeType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToMarkdownSupportedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toMarkdownSupportedResponseJSON) RawJSON() string {
	return r.raw
}

type ToMarkdownTransformResponse struct {
	Data     string                          `json:"data,required"`
	Format   string                          `json:"format,required"`
	MimeType string                          `json:"mimeType,required"`
	Name     string                          `json:"name,required"`
	Tokens   string                          `json:"tokens,required"`
	JSON     toMarkdownTransformResponseJSON `json:"-"`
}

// toMarkdownTransformResponseJSON contains the JSON metadata for the struct
// [ToMarkdownTransformResponse]
type toMarkdownTransformResponseJSON struct {
	Data        apijson.Field
	Format      apijson.Field
	MimeType    apijson.Field
	Name        apijson.Field
	Tokens      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToMarkdownTransformResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toMarkdownTransformResponseJSON) RawJSON() string {
	return r.raw
}

type ToMarkdownSupportedParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ToMarkdownTransformParams struct {
	AccountID param.Field[string]           `path:"account_id,required"`
	File      ToMarkdownTransformParamsFile `json:"file,required"`
}

func (r ToMarkdownTransformParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r.File, writer)
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

type ToMarkdownTransformParamsFile struct {
	Files param.Field[[]io.Reader] `json:"files,required" format:"binary"`
}

func (r ToMarkdownTransformParamsFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
