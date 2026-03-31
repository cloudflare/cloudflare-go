// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/rules"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// EnvironmentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentService] method instead.
type EnvironmentService struct {
	Options []option.RequestOption
}

// NewEnvironmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnvironmentService(opts ...option.RequestOption) (r *EnvironmentService) {
	r = &EnvironmentService{}
	r.Options = opts
	return
}

// Create zone environments
func (r *EnvironmentService) New(ctx context.Context, params EnvironmentNewParams, opts ...option.RequestOption) (res *EnvironmentNewResponse, err error) {
	var env EnvironmentNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/environments", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Upsert zone environments
func (r *EnvironmentService) Update(ctx context.Context, params EnvironmentUpdateParams, opts ...option.RequestOption) (res *EnvironmentUpdateResponse, err error) {
	var env EnvironmentUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/environments", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List zone environments
func (r *EnvironmentService) List(ctx context.Context, query EnvironmentListParams, opts ...option.RequestOption) (res *EnvironmentListResponse, err error) {
	var env EnvironmentListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/environments", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Delete zone environment
func (r *EnvironmentService) Delete(ctx context.Context, environmentID string, body EnvironmentDeleteParams, opts ...option.RequestOption) (res *EnvironmentDeleteResponse, err error) {
	var env EnvironmentDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if environmentID == "" {
		err = errors.New("missing required environment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/environments/%s", body.ZoneID, environmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Partially update zone environments
func (r *EnvironmentService) Edit(ctx context.Context, params EnvironmentEditParams, opts ...option.RequestOption) (res *EnvironmentEditResponse, err error) {
	var env EnvironmentEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/environments", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Roll back zone environment
func (r *EnvironmentService) Rollback(ctx context.Context, environmentID string, body EnvironmentRollbackParams, opts ...option.RequestOption) (res *EnvironmentRollbackResponse, err error) {
	var env EnvironmentRollbackResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if environmentID == "" {
		err = errors.New("missing required environment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/environments/%s/rollback", body.ZoneID, environmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type EnvironmentNewResponse struct {
	Environments []EnvironmentNewResponseEnvironment `json:"environments" api:"required"`
	JSON         environmentNewResponseJSON          `json:"-"`
}

// environmentNewResponseJSON contains the JSON metadata for the struct
// [EnvironmentNewResponse]
type environmentNewResponseJSON struct {
	Environments apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewResponseEnvironment struct {
	Expression         string                                `json:"expression" api:"required"`
	LockedOnDeployment bool                                  `json:"locked_on_deployment" api:"required,nullable"`
	Name               string                                `json:"name" api:"required"`
	Position           rules.ListCursor                      `json:"position" api:"required"`
	Ref                string                                `json:"ref" api:"required"`
	Version            int64                                 `json:"version" api:"required,nullable"`
	HTTPApplicationID  string                                `json:"http_application_id" api:"nullable"`
	JSON               environmentNewResponseEnvironmentJSON `json:"-"`
}

// environmentNewResponseEnvironmentJSON contains the JSON metadata for the struct
// [EnvironmentNewResponseEnvironment]
type environmentNewResponseEnvironmentJSON struct {
	Expression         apijson.Field
	LockedOnDeployment apijson.Field
	Name               apijson.Field
	Position           apijson.Field
	Ref                apijson.Field
	Version            apijson.Field
	HTTPApplicationID  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvironmentJSON) RawJSON() string {
	return r.raw
}

type EnvironmentUpdateResponse struct {
	Environments []EnvironmentUpdateResponseEnvironment `json:"environments" api:"required"`
	JSON         environmentUpdateResponseJSON          `json:"-"`
}

// environmentUpdateResponseJSON contains the JSON metadata for the struct
// [EnvironmentUpdateResponse]
type environmentUpdateResponseJSON struct {
	Environments apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentUpdateResponseEnvironment struct {
	Expression         string                                   `json:"expression" api:"required"`
	LockedOnDeployment bool                                     `json:"locked_on_deployment" api:"required,nullable"`
	Name               string                                   `json:"name" api:"required"`
	Position           rules.ListCursor                         `json:"position" api:"required"`
	Ref                string                                   `json:"ref" api:"required"`
	Version            int64                                    `json:"version" api:"required,nullable"`
	HTTPApplicationID  string                                   `json:"http_application_id" api:"nullable"`
	JSON               environmentUpdateResponseEnvironmentJSON `json:"-"`
}

// environmentUpdateResponseEnvironmentJSON contains the JSON metadata for the
// struct [EnvironmentUpdateResponseEnvironment]
type environmentUpdateResponseEnvironmentJSON struct {
	Expression         apijson.Field
	LockedOnDeployment apijson.Field
	Name               apijson.Field
	Position           apijson.Field
	Ref                apijson.Field
	Version            apijson.Field
	HTTPApplicationID  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentUpdateResponseEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentUpdateResponseEnvironmentJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListResponse struct {
	Environments []EnvironmentListResponseEnvironment `json:"environments" api:"required"`
	JSON         environmentListResponseJSON          `json:"-"`
}

// environmentListResponseJSON contains the JSON metadata for the struct
// [EnvironmentListResponse]
type environmentListResponseJSON struct {
	Environments apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListResponseEnvironment struct {
	Expression         string                                 `json:"expression" api:"required"`
	LockedOnDeployment bool                                   `json:"locked_on_deployment" api:"required,nullable"`
	Name               string                                 `json:"name" api:"required"`
	Position           rules.ListCursor                       `json:"position" api:"required"`
	Ref                string                                 `json:"ref" api:"required"`
	Version            int64                                  `json:"version" api:"required,nullable"`
	HTTPApplicationID  string                                 `json:"http_application_id" api:"nullable"`
	JSON               environmentListResponseEnvironmentJSON `json:"-"`
}

// environmentListResponseEnvironmentJSON contains the JSON metadata for the struct
// [EnvironmentListResponseEnvironment]
type environmentListResponseEnvironmentJSON struct {
	Expression         apijson.Field
	LockedOnDeployment apijson.Field
	Name               apijson.Field
	Position           apijson.Field
	Ref                apijson.Field
	Version            apijson.Field
	HTTPApplicationID  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvironmentJSON) RawJSON() string {
	return r.raw
}

type EnvironmentDeleteResponse struct {
	Environments []EnvironmentDeleteResponseEnvironment `json:"environments" api:"required"`
	JSON         environmentDeleteResponseJSON          `json:"-"`
}

// environmentDeleteResponseJSON contains the JSON metadata for the struct
// [EnvironmentDeleteResponse]
type environmentDeleteResponseJSON struct {
	Environments apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentDeleteResponseEnvironment struct {
	Expression         string                                   `json:"expression" api:"required"`
	LockedOnDeployment bool                                     `json:"locked_on_deployment" api:"required,nullable"`
	Name               string                                   `json:"name" api:"required"`
	Position           rules.ListCursor                         `json:"position" api:"required"`
	Ref                string                                   `json:"ref" api:"required"`
	Version            int64                                    `json:"version" api:"required,nullable"`
	HTTPApplicationID  string                                   `json:"http_application_id" api:"nullable"`
	JSON               environmentDeleteResponseEnvironmentJSON `json:"-"`
}

// environmentDeleteResponseEnvironmentJSON contains the JSON metadata for the
// struct [EnvironmentDeleteResponseEnvironment]
type environmentDeleteResponseEnvironmentJSON struct {
	Expression         apijson.Field
	LockedOnDeployment apijson.Field
	Name               apijson.Field
	Position           apijson.Field
	Ref                apijson.Field
	Version            apijson.Field
	HTTPApplicationID  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentDeleteResponseEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentDeleteResponseEnvironmentJSON) RawJSON() string {
	return r.raw
}

type EnvironmentEditResponse struct {
	Environments []EnvironmentEditResponseEnvironment `json:"environments" api:"required"`
	JSON         environmentEditResponseJSON          `json:"-"`
}

// environmentEditResponseJSON contains the JSON metadata for the struct
// [EnvironmentEditResponse]
type environmentEditResponseJSON struct {
	Environments apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentEditResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentEditResponseEnvironment struct {
	Expression         string                                 `json:"expression" api:"required"`
	LockedOnDeployment bool                                   `json:"locked_on_deployment" api:"required,nullable"`
	Name               string                                 `json:"name" api:"required"`
	Position           rules.ListCursor                       `json:"position" api:"required"`
	Ref                string                                 `json:"ref" api:"required"`
	Version            int64                                  `json:"version" api:"required,nullable"`
	HTTPApplicationID  string                                 `json:"http_application_id" api:"nullable"`
	JSON               environmentEditResponseEnvironmentJSON `json:"-"`
}

// environmentEditResponseEnvironmentJSON contains the JSON metadata for the struct
// [EnvironmentEditResponseEnvironment]
type environmentEditResponseEnvironmentJSON struct {
	Expression         apijson.Field
	LockedOnDeployment apijson.Field
	Name               apijson.Field
	Position           apijson.Field
	Ref                apijson.Field
	Version            apijson.Field
	HTTPApplicationID  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentEditResponseEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentEditResponseEnvironmentJSON) RawJSON() string {
	return r.raw
}

type EnvironmentRollbackResponse struct {
	Environments []EnvironmentRollbackResponseEnvironment `json:"environments" api:"required"`
	JSON         environmentRollbackResponseJSON          `json:"-"`
}

// environmentRollbackResponseJSON contains the JSON metadata for the struct
// [EnvironmentRollbackResponse]
type environmentRollbackResponseJSON struct {
	Environments apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentRollbackResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentRollbackResponseJSON) RawJSON() string {
	return r.raw
}

type EnvironmentRollbackResponseEnvironment struct {
	Expression         string                                     `json:"expression" api:"required"`
	LockedOnDeployment bool                                       `json:"locked_on_deployment" api:"required,nullable"`
	Name               string                                     `json:"name" api:"required"`
	Position           rules.ListCursor                           `json:"position" api:"required"`
	Ref                string                                     `json:"ref" api:"required"`
	Version            int64                                      `json:"version" api:"required,nullable"`
	HTTPApplicationID  string                                     `json:"http_application_id" api:"nullable"`
	JSON               environmentRollbackResponseEnvironmentJSON `json:"-"`
}

// environmentRollbackResponseEnvironmentJSON contains the JSON metadata for the
// struct [EnvironmentRollbackResponseEnvironment]
type environmentRollbackResponseEnvironmentJSON struct {
	Expression         apijson.Field
	LockedOnDeployment apijson.Field
	Name               apijson.Field
	Position           apijson.Field
	Ref                apijson.Field
	Version            apijson.Field
	HTTPApplicationID  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EnvironmentRollbackResponseEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentRollbackResponseEnvironmentJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewParams struct {
	ZoneID       param.Field[string]                            `path:"zone_id" api:"required"`
	Environments param.Field[[]EnvironmentNewParamsEnvironment] `json:"environments" api:"required"`
}

func (r EnvironmentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentNewParamsEnvironment struct {
	Expression         param.Field[string]                `json:"expression" api:"required"`
	LockedOnDeployment param.Field[bool]                  `json:"locked_on_deployment" api:"required"`
	Name               param.Field[string]                `json:"name" api:"required"`
	Position           param.Field[rules.ListCursorParam] `json:"position" api:"required"`
	Ref                param.Field[string]                `json:"ref" api:"required"`
	Version            param.Field[int64]                 `json:"version" api:"required"`
	HTTPApplicationID  param.Field[string]                `json:"http_application_id"`
}

func (r EnvironmentNewParamsEnvironment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo              `json:"errors" api:"required"`
	Messages []shared.ResponseInfo              `json:"messages" api:"required"`
	Result   EnvironmentNewResponse             `json:"result" api:"required"`
	Success  bool                               `json:"success" api:"required"`
	JSON     environmentNewResponseEnvelopeJSON `json:"-"`
}

// environmentNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [EnvironmentNewResponseEnvelope]
type environmentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EnvironmentUpdateParams struct {
	ZoneID       param.Field[string]                               `path:"zone_id" api:"required"`
	Environments param.Field[[]EnvironmentUpdateParamsEnvironment] `json:"environments" api:"required"`
}

func (r EnvironmentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateParamsEnvironment struct {
	Expression         param.Field[string]                `json:"expression" api:"required"`
	LockedOnDeployment param.Field[bool]                  `json:"locked_on_deployment" api:"required"`
	Name               param.Field[string]                `json:"name" api:"required"`
	Position           param.Field[rules.ListCursorParam] `json:"position" api:"required"`
	Ref                param.Field[string]                `json:"ref" api:"required"`
	Version            param.Field[int64]                 `json:"version" api:"required"`
	HTTPApplicationID  param.Field[string]                `json:"http_application_id"`
}

func (r EnvironmentUpdateParamsEnvironment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo                 `json:"errors" api:"required"`
	Messages []shared.ResponseInfo                 `json:"messages" api:"required"`
	Result   EnvironmentUpdateResponse             `json:"result" api:"required"`
	Success  bool                                  `json:"success" api:"required"`
	JSON     environmentUpdateResponseEnvelopeJSON `json:"-"`
}

// environmentUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [EnvironmentUpdateResponseEnvelope]
type environmentUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EnvironmentListParams struct {
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type EnvironmentListResponseEnvelope struct {
	Errors   []shared.ResponseInfo               `json:"errors" api:"required"`
	Messages []shared.ResponseInfo               `json:"messages" api:"required"`
	Result   EnvironmentListResponse             `json:"result" api:"required"`
	Success  bool                                `json:"success" api:"required"`
	JSON     environmentListResponseEnvelopeJSON `json:"-"`
}

// environmentListResponseEnvelopeJSON contains the JSON metadata for the struct
// [EnvironmentListResponseEnvelope]
type environmentListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EnvironmentDeleteParams struct {
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type EnvironmentDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                 `json:"errors" api:"required"`
	Messages []shared.ResponseInfo                 `json:"messages" api:"required"`
	Result   EnvironmentDeleteResponse             `json:"result" api:"required"`
	Success  bool                                  `json:"success" api:"required"`
	JSON     environmentDeleteResponseEnvelopeJSON `json:"-"`
}

// environmentDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [EnvironmentDeleteResponseEnvelope]
type environmentDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EnvironmentEditParams struct {
	ZoneID       param.Field[string]                             `path:"zone_id" api:"required"`
	Environments param.Field[[]EnvironmentEditParamsEnvironment] `json:"environments" api:"required"`
}

func (r EnvironmentEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentEditParamsEnvironment struct {
	Expression         param.Field[string]                `json:"expression" api:"required"`
	LockedOnDeployment param.Field[bool]                  `json:"locked_on_deployment" api:"required"`
	Name               param.Field[string]                `json:"name" api:"required"`
	Position           param.Field[rules.ListCursorParam] `json:"position" api:"required"`
	Ref                param.Field[string]                `json:"ref" api:"required"`
	Version            param.Field[int64]                 `json:"version" api:"required"`
	HTTPApplicationID  param.Field[string]                `json:"http_application_id"`
}

func (r EnvironmentEditParamsEnvironment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnvironmentEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo               `json:"errors" api:"required"`
	Messages []shared.ResponseInfo               `json:"messages" api:"required"`
	Result   EnvironmentEditResponse             `json:"result" api:"required"`
	Success  bool                                `json:"success" api:"required"`
	JSON     environmentEditResponseEnvelopeJSON `json:"-"`
}

// environmentEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [EnvironmentEditResponseEnvelope]
type environmentEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EnvironmentRollbackParams struct {
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type EnvironmentRollbackResponseEnvelope struct {
	Errors   []shared.ResponseInfo                   `json:"errors" api:"required"`
	Messages []shared.ResponseInfo                   `json:"messages" api:"required"`
	Result   EnvironmentRollbackResponse             `json:"result" api:"required"`
	Success  bool                                    `json:"success" api:"required"`
	JSON     environmentRollbackResponseEnvelopeJSON `json:"-"`
}

// environmentRollbackResponseEnvelopeJSON contains the JSON metadata for the
// struct [EnvironmentRollbackResponseEnvelope]
type environmentRollbackResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnvironmentRollbackResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentRollbackResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
