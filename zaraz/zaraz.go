// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zaraz

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

// ZarazService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZarazService] method instead.
type ZarazService struct {
	Options  []option.RequestOption
	Config   *ConfigService
	Default  *DefaultService
	Export   *ExportService
	History  *HistoryService
	Publish  *PublishService
	Workflow *WorkflowService
}

// NewZarazService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZarazService(opts ...option.RequestOption) (r *ZarazService) {
	r = &ZarazService{}
	r.Options = opts
	r.Config = NewConfigService(opts...)
	r.Default = NewDefaultService(opts...)
	r.Export = NewExportService(opts...)
	r.History = NewHistoryService(opts...)
	r.Publish = NewPublishService(opts...)
	r.Workflow = NewWorkflowService(opts...)
	return
}

// Updates Zaraz workflow for a zone.
func (r *ZarazService) Update(ctx context.Context, params ZarazUpdateParams, opts ...option.RequestOption) (res *Workflow, err error) {
	var env ZarazUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/workflow", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZarazUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Zaraz workflow
	Workflow Workflow `json:"workflow,required"`
}

func (r ZarazUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Workflow)
}

type ZarazUpdateResponseEnvelope struct {
	Errors   []interface{} `json:"errors,required"`
	Messages []interface{} `json:"messages,required"`
	// Zaraz workflow
	Result Workflow `json:"result,required"`
	// Whether the API call was successful
	Success bool                            `json:"success,required"`
	JSON    zarazUpdateResponseEnvelopeJSON `json:"-"`
}

// zarazUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZarazUpdateResponseEnvelope]
type zarazUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
