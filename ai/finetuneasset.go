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
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// FinetuneAssetService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinetuneAssetService] method instead.
type FinetuneAssetService struct {
	Options []option.RequestOption
}

// NewFinetuneAssetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFinetuneAssetService(opts ...option.RequestOption) (r *FinetuneAssetService) {
	r = &FinetuneAssetService{}
	r.Options = opts
	return
}

// Upload a Finetune Asset
func (r *FinetuneAssetService) New(ctx context.Context, finetuneID string, params FinetuneAssetNewParams, opts ...option.RequestOption) (res *FinetuneAssetNewResponse, err error) {
	var env FinetuneAssetNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if finetuneID == "" {
		err = errors.New("missing required finetune_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/finetunes/%s/finetune-assets", params.AccountID, finetuneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FinetuneAssetNewResponse struct {
	ID         string                       `json:"id,required" format:"uuid"`
	BucketName string                       `json:"bucket_name,required"`
	CreatedAt  time.Time                    `json:"created_at,required" format:"date-time"`
	FileName   string                       `json:"file_name,required"`
	FinetuneID string                       `json:"finetune_id,required"`
	ModifiedAt time.Time                    `json:"modified_at,required" format:"date-time"`
	JSON       finetuneAssetNewResponseJSON `json:"-"`
}

// finetuneAssetNewResponseJSON contains the JSON metadata for the struct
// [FinetuneAssetNewResponse]
type finetuneAssetNewResponseJSON struct {
	ID          apijson.Field
	BucketName  apijson.Field
	CreatedAt   apijson.Field
	FileName    apijson.Field
	FinetuneID  apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinetuneAssetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r finetuneAssetNewResponseJSON) RawJSON() string {
	return r.raw
}

type FinetuneAssetNewParams struct {
	AccountID param.Field[string]    `path:"account_id,required"`
	File      param.Field[io.Reader] `json:"file" format:"binary"`
	FileName  param.Field[string]    `json:"file_name"`
}

func (r FinetuneAssetNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type FinetuneAssetNewResponseEnvelope struct {
	Result  FinetuneAssetNewResponse             `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    finetuneAssetNewResponseEnvelopeJSON `json:"-"`
}

// finetuneAssetNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [FinetuneAssetNewResponseEnvelope]
type finetuneAssetNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinetuneAssetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r finetuneAssetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
