package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type AdvancedDdoListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
	// website. This is an uneditable value that is 'on' in the case of Business and
	// Enterprise zones.
	Result fields.Field[AdvancedDdos] `json:"result"`
}

// MarshalJSON serializes AdvancedDdoListResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AdvancedDdoListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AdvancedDdoListResponse) String() (result string) {
	return fmt.Sprintf("&AdvancedDdoListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}
