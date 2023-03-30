package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type OrangeToOrangeListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result fields.Field[OrangeToOrange] `json:"result"`
}

// MarshalJSON serializes OrangeToOrangeListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *OrangeToOrangeListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OrangeToOrangeListResponse) String() (result string) {
	return fmt.Sprintf("&OrangeToOrangeListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type OrangeToOrangeUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result fields.Field[OrangeToOrange] `json:"result"`
}

// MarshalJSON serializes OrangeToOrangeUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *OrangeToOrangeUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OrangeToOrangeUpdateResponse) String() (result string) {
	return fmt.Sprintf("&OrangeToOrangeUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type OrangeToOrangeUpdateParams struct {
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Value fields.Field[OrangeToOrange] `json:"value,required"`
}

// MarshalJSON serializes OrangeToOrangeUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *OrangeToOrangeUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OrangeToOrangeUpdateParams) String() (result string) {
	return fmt.Sprintf("&OrangeToOrangeUpdateParams{Value:%s}", r.Value)
}
