package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type OpportunisticEncryptionListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result fields.Field[OpportunisticEncryption] `json:"result"`
}

// MarshalJSON serializes OpportunisticEncryptionListResponse into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *OpportunisticEncryptionListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OpportunisticEncryptionListResponse) String() (result string) {
	return fmt.Sprintf("&OpportunisticEncryptionListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type OpportunisticEncryptionUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result fields.Field[OpportunisticEncryption] `json:"result"`
}

// MarshalJSON serializes OpportunisticEncryptionUpdateResponse into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *OpportunisticEncryptionUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OpportunisticEncryptionUpdateResponse) String() (result string) {
	return fmt.Sprintf("&OpportunisticEncryptionUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type OpportunisticEncryptionUpdateParams struct {
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value fields.Field[OpportunisticEncryptionValue] `json:"value,required"`
}

// MarshalJSON serializes OpportunisticEncryptionUpdateParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *OpportunisticEncryptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OpportunisticEncryptionUpdateParams) String() (result string) {
	return fmt.Sprintf("&OpportunisticEncryptionUpdateParams{Value:%s}", r.Value)
}
