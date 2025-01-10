// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

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

// ScriptSubdomainService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptSubdomainService] method instead.
type ScriptSubdomainService struct {
	Options []option.RequestOption
}

// NewScriptSubdomainService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptSubdomainService(opts ...option.RequestOption) (r *ScriptSubdomainService) {
	r = &ScriptSubdomainService{}
	r.Options = opts
	return
}

// Enable or disable the Worker on the workers.dev subdomain.
func (r *ScriptSubdomainService) New(ctx context.Context, scriptName string, params ScriptSubdomainNewParams, opts ...option.RequestOption) (res *ScriptSubdomainNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/subdomain", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get if the Worker is available on the workers.dev subdomain.
func (r *ScriptSubdomainService) Get(ctx context.Context, scriptName string, query ScriptSubdomainGetParams, opts ...option.RequestOption) (res *ScriptSubdomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/subdomain", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ScriptSubdomainNewResponse struct {
	// Whether the Worker is available on the workers.dev subdomain.
	Enabled bool `json:"enabled"`
	// Whether the Worker's Preview URLs should be available on the workers.dev
	// subdomain.
	PreviewsEnabled bool                           `json:"previews_enabled"`
	JSON            scriptSubdomainNewResponseJSON `json:"-"`
}

// scriptSubdomainNewResponseJSON contains the JSON metadata for the struct
// [ScriptSubdomainNewResponse]
type scriptSubdomainNewResponseJSON struct {
	Enabled         apijson.Field
	PreviewsEnabled apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScriptSubdomainNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainGetResponse struct {
	// Whether the Worker is available on the workers.dev subdomain.
	Enabled bool `json:"enabled"`
	// Whether the Worker's Preview URLs should be available on the workers.dev
	// subdomain.
	PreviewsEnabled bool                           `json:"previews_enabled"`
	JSON            scriptSubdomainGetResponseJSON `json:"-"`
}

// scriptSubdomainGetResponseJSON contains the JSON metadata for the struct
// [ScriptSubdomainGetResponse]
type scriptSubdomainGetResponseJSON struct {
	Enabled         apijson.Field
	PreviewsEnabled apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScriptSubdomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSubdomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptSubdomainNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Whether the Worker should be available on the workers.dev subdomain.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Whether the Worker's Preview URLs should be available on the workers.dev
	// subdomain.
	PreviewsEnabled param.Field[bool] `json:"previews_enabled"`
}

func (r ScriptSubdomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSubdomainGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
