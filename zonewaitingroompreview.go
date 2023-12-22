// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneWaitingRoomPreviewService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneWaitingRoomPreviewService]
// method instead.
type ZoneWaitingRoomPreviewService struct {
	Options []option.RequestOption
}

// NewZoneWaitingRoomPreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWaitingRoomPreviewService(opts ...option.RequestOption) (r *ZoneWaitingRoomPreviewService) {
	r = &ZoneWaitingRoomPreviewService{}
	r.Options = opts
	return
}

// Creates a waiting room page preview. Upload a custom waiting room page for
// preview. You will receive a preview URL in the form
// `http://waitingrooms.dev/preview/<uuid>`. You can use the following query
// parameters to change the state of the preview:
//
//  1. `force_queue`: Boolean indicating if all users will be queued in the waiting
//     room and no one will be let into the origin website (also known as queueAll).
//  2. `queue_is_full`: Boolean indicating if the waiting room's queue is currently
//     full and not accepting new users at the moment.
//  3. `queueing_method`: The queueing method currently used by the waiting room.
//     - **fifo** indicates a FIFO queue.
//     - **random** indicates a Random queue.
//     - **passthrough** indicates a Passthrough queue. Keep in mind that the
//     waiting room page will only be displayed if `force_queue=true` or
//     `event=prequeueing` — for other cases the request will pass through to the
//     origin. For our preview, this will be a fake origin website returning
//     "Welcome".
//     - **reject** indicates a Reject queue.
//  4. `event`: Used to preview a waiting room event.
//     - **none** indicates no event is occurring.
//     - **prequeueing** indicates that an event is prequeueing (between
//     `prequeue_start_time` and `event_start_time`).
//     - **started** indicates that an event has started (between `event_start_time`
//     and `event_end_time`).
//  5. `shuffle_at_event_start`: Boolean indicating if the event will shuffle users
//     in the prequeue when it starts. This can only be set to **true** if an event
//     is active (`event` is not **none**).
//
// For example, you can make a request to
// `http://waitingrooms.dev/preview/<uuid>?force_queue=false&queue_is_full=false&queueing_method=random&event=started&shuffle_at_event_start=true` 6.
// `waitTime`: Non-zero, positive integer indicating the estimated wait time in
// minutes. The default value is 10 minutes.
//
// For example, you can make a request to
// `http://waitingrooms.dev/preview/<uuid>?waitTime=50` to configure the estimated
// wait time as 50 minutes.
func (r *ZoneWaitingRoomPreviewService) WaitingRoomNewACustomWaitingRoomPagePreview(ctx context.Context, zoneIdentifier string, body ZoneWaitingRoomPreviewWaitingRoomNewACustomWaitingRoomPagePreviewParams, opts ...option.RequestOption) (res *PreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/preview", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PreviewResponse struct {
	Errors   []PreviewResponseError   `json:"errors"`
	Messages []PreviewResponseMessage `json:"messages"`
	Result   PreviewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success PreviewResponseSuccess `json:"success"`
	JSON    previewResponseJSON    `json:"-"`
}

// previewResponseJSON contains the JSON metadata for the struct [PreviewResponse]
type previewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewResponseError struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    previewResponseErrorJSON `json:"-"`
}

// previewResponseErrorJSON contains the JSON metadata for the struct
// [PreviewResponseError]
type previewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewResponseMessage struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    previewResponseMessageJSON `json:"-"`
}

// previewResponseMessageJSON contains the JSON metadata for the struct
// [PreviewResponseMessage]
type previewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewResponseResult struct {
	// URL where the custom waiting room page can temporarily be previewed.
	PreviewURL string                    `json:"preview_url"`
	JSON       previewResponseResultJSON `json:"-"`
}

// previewResponseResultJSON contains the JSON metadata for the struct
// [PreviewResponseResult]
type previewResponseResultJSON struct {
	PreviewURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PreviewResponseSuccess bool

const (
	PreviewResponseSuccessTrue PreviewResponseSuccess = true
)

type ZoneWaitingRoomPreviewWaitingRoomNewACustomWaitingRoomPagePreviewParams struct {
	// Only available for the Waiting Room Advanced subscription. This is a template
	// html file that will be rendered at the edge. If no custom_page_html is provided,
	// the default waiting room will be used. The template is based on mustache (
	// https://mustache.github.io/ ). There are several variables that are evaluated by
	// the Cloudflare edge:
	//
	//  1. {{`waitTimeKnown`}} Acts like a boolean value that indicates the behavior to
	//     take when wait time is not available, for instance when queue_all is
	//     **true**.
	//  2. {{`waitTimeFormatted`}} Estimated wait time for the user. For example, five
	//     minutes. Alternatively, you can use:
	//  3. {{`waitTime`}} Number of minutes of estimated wait for a user.
	//  4. {{`waitTimeHours`}} Number of hours of estimated wait for a user
	//     (`Math.floor(waitTime/60)`).
	//  5. {{`waitTimeHourMinutes`}} Number of minutes above the `waitTimeHours` value
	//     (`waitTime%60`).
	//  6. {{`queueIsFull`}} Changes to **true** when no more people can be added to the
	//     queue.
	//
	// To view the full list of variables, look at the `cfWaitingRoom` object described
	// under the `json_response_enabled` property in other Waiting Room API calls.
	CustomHTML param.Field[string] `json:"custom_html,required"`
}

func (r ZoneWaitingRoomPreviewWaitingRoomNewACustomWaitingRoomPagePreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
