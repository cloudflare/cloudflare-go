package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

var errMissingAccessSeatUID = errors.New("missing required access seat UID")

type AccessUserSeat struct {
	AccessSeat  bool   `json:"access_seat"`
	CreatedAt   string `json:"created_at"`
	GatewaySeat bool   `json:"gateway_seat"`
	SeatUID     string `json:"seat_uid"`
	UpdatedAt   string `json:"updated_at"`
}

// Removes a user from a Zero Trust seat when both access_seat and gateway_seat are set to false.
type UpdateAccessUserSeatParams struct {
	SeatUID     string `json:"seat_uid"`
	AccessSeat  bool   `json:"access_seat"`
	GatewaySeat bool   `json:"gateway_seat"`
}

type UpdateAccessUserSeatResponse struct {
	Success    bool           `json:"success"`
	Errors     []string       `json:"errors"`
	Messages   []string       `json:"messages"`
	Result     AccessUserSeat `json:"result"`
	ResultInfo ResultInfo     `json:"result_info"`
}

// UpdateAccessUserSeat updates a Access User Seat.
// UpdateAccessUserSeat Removes a user from a Zero Trust seat when both access_seat and gateway_seat are set to false.
//
// Access API reference: https://developers.cloudflare.com/api/operations/zero-trust-seats-update-a-user-seat
func (api *API) UpdateAccessUserSeat(ctx context.Context, rc *ResourceContainer, params UpdateAccessUserSeatParams) (AccessUserSeat, error) {
	if params.SeatUID == "" {
		return AccessUserSeat{}, errMissingAccessSeatUID
	}

	uri := fmt.Sprintf(
		"/%s/%s/access/seats",
		rc.Level,
		rc.Identifier,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return AccessUserSeat{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var updateAccessUserSeatResponse UpdateAccessUserSeatResponse
	err = json.Unmarshal(res, &updateAccessUserSeatResponse)
	if err != nil {
		return AccessUserSeat{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return updateAccessUserSeatResponse.Result, nil
}
