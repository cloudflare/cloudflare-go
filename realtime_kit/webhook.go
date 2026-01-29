// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package realtime_kit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// WebhookService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	return
}

// Adds a new webhook to an App.
func (r *WebhookService) NewWebhook(ctx context.Context, appID string, params WebhookNewWebhookParams, opts ...option.RequestOption) (res *WebhookNewWebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/webhooks", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Removes a webhook for the given webhook ID.
func (r *WebhookService) DeleteWebhook(ctx context.Context, appID string, webhookID string, body WebhookDeleteWebhookParams, opts ...option.RequestOption) (res *WebhookDeleteWebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/webhooks/%s", body.AccountID, appID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Edits the webhook details for the given webhook ID.
func (r *WebhookService) EditWebhook(ctx context.Context, appID string, webhookID string, params WebhookEditWebhookParams, opts ...option.RequestOption) (res *WebhookEditWebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/webhooks/%s", params.AccountID, appID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns webhook details for the given webhook ID.
func (r *WebhookService) GetWebhookByID(ctx context.Context, appID string, webhookID string, query WebhookGetWebhookByIDParams, opts ...option.RequestOption) (res *WebhookGetWebhookByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/webhooks/%s", query.AccountID, appID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns details of all webhooks for an App.
func (r *WebhookService) GetWebhooks(ctx context.Context, appID string, query WebhookGetWebhooksParams, opts ...option.RequestOption) (res *WebhookGetWebhooksResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/webhooks", query.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Replace all details for the given webhook ID.
func (r *WebhookService) ReplaceWebhook(ctx context.Context, appID string, webhookID string, params WebhookReplaceWebhookParams, opts ...option.RequestOption) (res *WebhookReplaceWebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/webhooks/%s", params.AccountID, appID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type WebhookNewWebhookResponse struct {
	Data    WebhookNewWebhookResponseData `json:"data,required"`
	Success bool                          `json:"success,required"`
	JSON    webhookNewWebhookResponseJSON `json:"-"`
}

// webhookNewWebhookResponseJSON contains the JSON metadata for the struct
// [WebhookNewWebhookResponse]
type webhookNewWebhookResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookNewWebhookResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookNewWebhookResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookNewWebhookResponseData struct {
	// ID of the webhook
	ID string `json:"id,required" format:"uuid"`
	// Timestamp when this webhook was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Set to true if the webhook is active
	Enabled bool `json:"enabled,required"`
	// Events this webhook will send updates for
	Events []WebhookNewWebhookResponseDataEvent `json:"events,required"`
	// Name of the webhook
	Name string `json:"name,required"`
	// Timestamp when this webhook was updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// URL the webhook will send events to
	URL  string                            `json:"url,required" format:"uri"`
	JSON webhookNewWebhookResponseDataJSON `json:"-"`
}

// webhookNewWebhookResponseDataJSON contains the JSON metadata for the struct
// [WebhookNewWebhookResponseData]
type webhookNewWebhookResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookNewWebhookResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookNewWebhookResponseDataJSON) RawJSON() string {
	return r.raw
}

type WebhookNewWebhookResponseDataEvent string

const (
	WebhookNewWebhookResponseDataEventMeetingStarted            WebhookNewWebhookResponseDataEvent = "meeting.started"
	WebhookNewWebhookResponseDataEventMeetingEnded              WebhookNewWebhookResponseDataEvent = "meeting.ended"
	WebhookNewWebhookResponseDataEventMeetingParticipantJoined  WebhookNewWebhookResponseDataEvent = "meeting.participantJoined"
	WebhookNewWebhookResponseDataEventMeetingParticipantLeft    WebhookNewWebhookResponseDataEvent = "meeting.participantLeft"
	WebhookNewWebhookResponseDataEventMeetingChatSynced         WebhookNewWebhookResponseDataEvent = "meeting.chatSynced"
	WebhookNewWebhookResponseDataEventRecordingStatusUpdate     WebhookNewWebhookResponseDataEvent = "recording.statusUpdate"
	WebhookNewWebhookResponseDataEventLivestreamingStatusUpdate WebhookNewWebhookResponseDataEvent = "livestreaming.statusUpdate"
	WebhookNewWebhookResponseDataEventMeetingTranscript         WebhookNewWebhookResponseDataEvent = "meeting.transcript"
	WebhookNewWebhookResponseDataEventMeetingSummary            WebhookNewWebhookResponseDataEvent = "meeting.summary"
)

func (r WebhookNewWebhookResponseDataEvent) IsKnown() bool {
	switch r {
	case WebhookNewWebhookResponseDataEventMeetingStarted, WebhookNewWebhookResponseDataEventMeetingEnded, WebhookNewWebhookResponseDataEventMeetingParticipantJoined, WebhookNewWebhookResponseDataEventMeetingParticipantLeft, WebhookNewWebhookResponseDataEventMeetingChatSynced, WebhookNewWebhookResponseDataEventRecordingStatusUpdate, WebhookNewWebhookResponseDataEventLivestreamingStatusUpdate, WebhookNewWebhookResponseDataEventMeetingTranscript, WebhookNewWebhookResponseDataEventMeetingSummary:
		return true
	}
	return false
}

type WebhookDeleteWebhookResponse struct {
	Data    WebhookDeleteWebhookResponseData `json:"data,required"`
	Success bool                             `json:"success,required"`
	JSON    webhookDeleteWebhookResponseJSON `json:"-"`
}

// webhookDeleteWebhookResponseJSON contains the JSON metadata for the struct
// [WebhookDeleteWebhookResponse]
type webhookDeleteWebhookResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookDeleteWebhookResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookDeleteWebhookResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookDeleteWebhookResponseData struct {
	// ID of the webhook
	ID string `json:"id,required" format:"uuid"`
	// Timestamp when this webhook was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Set to true if the webhook is active
	Enabled bool `json:"enabled,required"`
	// Events this webhook will send updates for
	Events []WebhookDeleteWebhookResponseDataEvent `json:"events,required"`
	// Name of the webhook
	Name string `json:"name,required"`
	// Timestamp when this webhook was updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// URL the webhook will send events to
	URL  string                               `json:"url,required" format:"uri"`
	JSON webhookDeleteWebhookResponseDataJSON `json:"-"`
}

// webhookDeleteWebhookResponseDataJSON contains the JSON metadata for the struct
// [WebhookDeleteWebhookResponseData]
type webhookDeleteWebhookResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookDeleteWebhookResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookDeleteWebhookResponseDataJSON) RawJSON() string {
	return r.raw
}

type WebhookDeleteWebhookResponseDataEvent string

const (
	WebhookDeleteWebhookResponseDataEventMeetingStarted            WebhookDeleteWebhookResponseDataEvent = "meeting.started"
	WebhookDeleteWebhookResponseDataEventMeetingEnded              WebhookDeleteWebhookResponseDataEvent = "meeting.ended"
	WebhookDeleteWebhookResponseDataEventMeetingParticipantJoined  WebhookDeleteWebhookResponseDataEvent = "meeting.participantJoined"
	WebhookDeleteWebhookResponseDataEventMeetingParticipantLeft    WebhookDeleteWebhookResponseDataEvent = "meeting.participantLeft"
	WebhookDeleteWebhookResponseDataEventMeetingChatSynced         WebhookDeleteWebhookResponseDataEvent = "meeting.chatSynced"
	WebhookDeleteWebhookResponseDataEventRecordingStatusUpdate     WebhookDeleteWebhookResponseDataEvent = "recording.statusUpdate"
	WebhookDeleteWebhookResponseDataEventLivestreamingStatusUpdate WebhookDeleteWebhookResponseDataEvent = "livestreaming.statusUpdate"
	WebhookDeleteWebhookResponseDataEventMeetingTranscript         WebhookDeleteWebhookResponseDataEvent = "meeting.transcript"
	WebhookDeleteWebhookResponseDataEventMeetingSummary            WebhookDeleteWebhookResponseDataEvent = "meeting.summary"
)

func (r WebhookDeleteWebhookResponseDataEvent) IsKnown() bool {
	switch r {
	case WebhookDeleteWebhookResponseDataEventMeetingStarted, WebhookDeleteWebhookResponseDataEventMeetingEnded, WebhookDeleteWebhookResponseDataEventMeetingParticipantJoined, WebhookDeleteWebhookResponseDataEventMeetingParticipantLeft, WebhookDeleteWebhookResponseDataEventMeetingChatSynced, WebhookDeleteWebhookResponseDataEventRecordingStatusUpdate, WebhookDeleteWebhookResponseDataEventLivestreamingStatusUpdate, WebhookDeleteWebhookResponseDataEventMeetingTranscript, WebhookDeleteWebhookResponseDataEventMeetingSummary:
		return true
	}
	return false
}

type WebhookEditWebhookResponse struct {
	Data    WebhookEditWebhookResponseData `json:"data,required"`
	Success bool                           `json:"success,required"`
	JSON    webhookEditWebhookResponseJSON `json:"-"`
}

// webhookEditWebhookResponseJSON contains the JSON metadata for the struct
// [WebhookEditWebhookResponse]
type webhookEditWebhookResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookEditWebhookResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookEditWebhookResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookEditWebhookResponseData struct {
	// ID of the webhook
	ID string `json:"id,required" format:"uuid"`
	// Timestamp when this webhook was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Set to true if the webhook is active
	Enabled bool `json:"enabled,required"`
	// Events this webhook will send updates for
	Events []WebhookEditWebhookResponseDataEvent `json:"events,required"`
	// Name of the webhook
	Name string `json:"name,required"`
	// Timestamp when this webhook was updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// URL the webhook will send events to
	URL  string                             `json:"url,required" format:"uri"`
	JSON webhookEditWebhookResponseDataJSON `json:"-"`
}

// webhookEditWebhookResponseDataJSON contains the JSON metadata for the struct
// [WebhookEditWebhookResponseData]
type webhookEditWebhookResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookEditWebhookResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookEditWebhookResponseDataJSON) RawJSON() string {
	return r.raw
}

type WebhookEditWebhookResponseDataEvent string

const (
	WebhookEditWebhookResponseDataEventMeetingStarted            WebhookEditWebhookResponseDataEvent = "meeting.started"
	WebhookEditWebhookResponseDataEventMeetingEnded              WebhookEditWebhookResponseDataEvent = "meeting.ended"
	WebhookEditWebhookResponseDataEventMeetingParticipantJoined  WebhookEditWebhookResponseDataEvent = "meeting.participantJoined"
	WebhookEditWebhookResponseDataEventMeetingParticipantLeft    WebhookEditWebhookResponseDataEvent = "meeting.participantLeft"
	WebhookEditWebhookResponseDataEventMeetingChatSynced         WebhookEditWebhookResponseDataEvent = "meeting.chatSynced"
	WebhookEditWebhookResponseDataEventRecordingStatusUpdate     WebhookEditWebhookResponseDataEvent = "recording.statusUpdate"
	WebhookEditWebhookResponseDataEventLivestreamingStatusUpdate WebhookEditWebhookResponseDataEvent = "livestreaming.statusUpdate"
	WebhookEditWebhookResponseDataEventMeetingTranscript         WebhookEditWebhookResponseDataEvent = "meeting.transcript"
	WebhookEditWebhookResponseDataEventMeetingSummary            WebhookEditWebhookResponseDataEvent = "meeting.summary"
)

func (r WebhookEditWebhookResponseDataEvent) IsKnown() bool {
	switch r {
	case WebhookEditWebhookResponseDataEventMeetingStarted, WebhookEditWebhookResponseDataEventMeetingEnded, WebhookEditWebhookResponseDataEventMeetingParticipantJoined, WebhookEditWebhookResponseDataEventMeetingParticipantLeft, WebhookEditWebhookResponseDataEventMeetingChatSynced, WebhookEditWebhookResponseDataEventRecordingStatusUpdate, WebhookEditWebhookResponseDataEventLivestreamingStatusUpdate, WebhookEditWebhookResponseDataEventMeetingTranscript, WebhookEditWebhookResponseDataEventMeetingSummary:
		return true
	}
	return false
}

type WebhookGetWebhookByIDResponse struct {
	Data    WebhookGetWebhookByIDResponseData `json:"data,required"`
	Success bool                              `json:"success,required"`
	JSON    webhookGetWebhookByIDResponseJSON `json:"-"`
}

// webhookGetWebhookByIDResponseJSON contains the JSON metadata for the struct
// [WebhookGetWebhookByIDResponse]
type webhookGetWebhookByIDResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookGetWebhookByIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookGetWebhookByIDResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookGetWebhookByIDResponseData struct {
	// ID of the webhook
	ID string `json:"id,required" format:"uuid"`
	// Timestamp when this webhook was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Set to true if the webhook is active
	Enabled bool `json:"enabled,required"`
	// Events this webhook will send updates for
	Events []WebhookGetWebhookByIDResponseDataEvent `json:"events,required"`
	// Name of the webhook
	Name string `json:"name,required"`
	// Timestamp when this webhook was updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// URL the webhook will send events to
	URL  string                                `json:"url,required" format:"uri"`
	JSON webhookGetWebhookByIDResponseDataJSON `json:"-"`
}

// webhookGetWebhookByIDResponseDataJSON contains the JSON metadata for the struct
// [WebhookGetWebhookByIDResponseData]
type webhookGetWebhookByIDResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookGetWebhookByIDResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookGetWebhookByIDResponseDataJSON) RawJSON() string {
	return r.raw
}

type WebhookGetWebhookByIDResponseDataEvent string

const (
	WebhookGetWebhookByIDResponseDataEventMeetingStarted            WebhookGetWebhookByIDResponseDataEvent = "meeting.started"
	WebhookGetWebhookByIDResponseDataEventMeetingEnded              WebhookGetWebhookByIDResponseDataEvent = "meeting.ended"
	WebhookGetWebhookByIDResponseDataEventMeetingParticipantJoined  WebhookGetWebhookByIDResponseDataEvent = "meeting.participantJoined"
	WebhookGetWebhookByIDResponseDataEventMeetingParticipantLeft    WebhookGetWebhookByIDResponseDataEvent = "meeting.participantLeft"
	WebhookGetWebhookByIDResponseDataEventMeetingChatSynced         WebhookGetWebhookByIDResponseDataEvent = "meeting.chatSynced"
	WebhookGetWebhookByIDResponseDataEventRecordingStatusUpdate     WebhookGetWebhookByIDResponseDataEvent = "recording.statusUpdate"
	WebhookGetWebhookByIDResponseDataEventLivestreamingStatusUpdate WebhookGetWebhookByIDResponseDataEvent = "livestreaming.statusUpdate"
	WebhookGetWebhookByIDResponseDataEventMeetingTranscript         WebhookGetWebhookByIDResponseDataEvent = "meeting.transcript"
	WebhookGetWebhookByIDResponseDataEventMeetingSummary            WebhookGetWebhookByIDResponseDataEvent = "meeting.summary"
)

func (r WebhookGetWebhookByIDResponseDataEvent) IsKnown() bool {
	switch r {
	case WebhookGetWebhookByIDResponseDataEventMeetingStarted, WebhookGetWebhookByIDResponseDataEventMeetingEnded, WebhookGetWebhookByIDResponseDataEventMeetingParticipantJoined, WebhookGetWebhookByIDResponseDataEventMeetingParticipantLeft, WebhookGetWebhookByIDResponseDataEventMeetingChatSynced, WebhookGetWebhookByIDResponseDataEventRecordingStatusUpdate, WebhookGetWebhookByIDResponseDataEventLivestreamingStatusUpdate, WebhookGetWebhookByIDResponseDataEventMeetingTranscript, WebhookGetWebhookByIDResponseDataEventMeetingSummary:
		return true
	}
	return false
}

type WebhookGetWebhooksResponse struct {
	Data    []WebhookGetWebhooksResponseData `json:"data,required"`
	Success bool                             `json:"success,required"`
	JSON    webhookGetWebhooksResponseJSON   `json:"-"`
}

// webhookGetWebhooksResponseJSON contains the JSON metadata for the struct
// [WebhookGetWebhooksResponse]
type webhookGetWebhooksResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookGetWebhooksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookGetWebhooksResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookGetWebhooksResponseData struct {
	// ID of the webhook
	ID string `json:"id,required" format:"uuid"`
	// Timestamp when this webhook was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Set to true if the webhook is active
	Enabled bool `json:"enabled,required"`
	// Events this webhook will send updates for
	Events []WebhookGetWebhooksResponseDataEvent `json:"events,required"`
	// Name of the webhook
	Name string `json:"name,required"`
	// Timestamp when this webhook was updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// URL the webhook will send events to
	URL  string                             `json:"url,required" format:"uri"`
	JSON webhookGetWebhooksResponseDataJSON `json:"-"`
}

// webhookGetWebhooksResponseDataJSON contains the JSON metadata for the struct
// [WebhookGetWebhooksResponseData]
type webhookGetWebhooksResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookGetWebhooksResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookGetWebhooksResponseDataJSON) RawJSON() string {
	return r.raw
}

type WebhookGetWebhooksResponseDataEvent string

const (
	WebhookGetWebhooksResponseDataEventMeetingStarted            WebhookGetWebhooksResponseDataEvent = "meeting.started"
	WebhookGetWebhooksResponseDataEventMeetingEnded              WebhookGetWebhooksResponseDataEvent = "meeting.ended"
	WebhookGetWebhooksResponseDataEventMeetingParticipantJoined  WebhookGetWebhooksResponseDataEvent = "meeting.participantJoined"
	WebhookGetWebhooksResponseDataEventMeetingParticipantLeft    WebhookGetWebhooksResponseDataEvent = "meeting.participantLeft"
	WebhookGetWebhooksResponseDataEventMeetingChatSynced         WebhookGetWebhooksResponseDataEvent = "meeting.chatSynced"
	WebhookGetWebhooksResponseDataEventRecordingStatusUpdate     WebhookGetWebhooksResponseDataEvent = "recording.statusUpdate"
	WebhookGetWebhooksResponseDataEventLivestreamingStatusUpdate WebhookGetWebhooksResponseDataEvent = "livestreaming.statusUpdate"
	WebhookGetWebhooksResponseDataEventMeetingTranscript         WebhookGetWebhooksResponseDataEvent = "meeting.transcript"
	WebhookGetWebhooksResponseDataEventMeetingSummary            WebhookGetWebhooksResponseDataEvent = "meeting.summary"
)

func (r WebhookGetWebhooksResponseDataEvent) IsKnown() bool {
	switch r {
	case WebhookGetWebhooksResponseDataEventMeetingStarted, WebhookGetWebhooksResponseDataEventMeetingEnded, WebhookGetWebhooksResponseDataEventMeetingParticipantJoined, WebhookGetWebhooksResponseDataEventMeetingParticipantLeft, WebhookGetWebhooksResponseDataEventMeetingChatSynced, WebhookGetWebhooksResponseDataEventRecordingStatusUpdate, WebhookGetWebhooksResponseDataEventLivestreamingStatusUpdate, WebhookGetWebhooksResponseDataEventMeetingTranscript, WebhookGetWebhooksResponseDataEventMeetingSummary:
		return true
	}
	return false
}

type WebhookReplaceWebhookResponse struct {
	Data    WebhookReplaceWebhookResponseData `json:"data,required"`
	Success bool                              `json:"success,required"`
	JSON    webhookReplaceWebhookResponseJSON `json:"-"`
}

// webhookReplaceWebhookResponseJSON contains the JSON metadata for the struct
// [WebhookReplaceWebhookResponse]
type webhookReplaceWebhookResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookReplaceWebhookResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookReplaceWebhookResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookReplaceWebhookResponseData struct {
	// ID of the webhook
	ID string `json:"id,required" format:"uuid"`
	// Timestamp when this webhook was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Set to true if the webhook is active
	Enabled bool `json:"enabled,required"`
	// Events this webhook will send updates for
	Events []WebhookReplaceWebhookResponseDataEvent `json:"events,required"`
	// Name of the webhook
	Name string `json:"name,required"`
	// Timestamp when this webhook was updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// URL the webhook will send events to
	URL  string                                `json:"url,required" format:"uri"`
	JSON webhookReplaceWebhookResponseDataJSON `json:"-"`
}

// webhookReplaceWebhookResponseDataJSON contains the JSON metadata for the struct
// [WebhookReplaceWebhookResponseData]
type webhookReplaceWebhookResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Events      apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookReplaceWebhookResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookReplaceWebhookResponseDataJSON) RawJSON() string {
	return r.raw
}

type WebhookReplaceWebhookResponseDataEvent string

const (
	WebhookReplaceWebhookResponseDataEventMeetingStarted            WebhookReplaceWebhookResponseDataEvent = "meeting.started"
	WebhookReplaceWebhookResponseDataEventMeetingEnded              WebhookReplaceWebhookResponseDataEvent = "meeting.ended"
	WebhookReplaceWebhookResponseDataEventMeetingParticipantJoined  WebhookReplaceWebhookResponseDataEvent = "meeting.participantJoined"
	WebhookReplaceWebhookResponseDataEventMeetingParticipantLeft    WebhookReplaceWebhookResponseDataEvent = "meeting.participantLeft"
	WebhookReplaceWebhookResponseDataEventMeetingChatSynced         WebhookReplaceWebhookResponseDataEvent = "meeting.chatSynced"
	WebhookReplaceWebhookResponseDataEventRecordingStatusUpdate     WebhookReplaceWebhookResponseDataEvent = "recording.statusUpdate"
	WebhookReplaceWebhookResponseDataEventLivestreamingStatusUpdate WebhookReplaceWebhookResponseDataEvent = "livestreaming.statusUpdate"
	WebhookReplaceWebhookResponseDataEventMeetingTranscript         WebhookReplaceWebhookResponseDataEvent = "meeting.transcript"
	WebhookReplaceWebhookResponseDataEventMeetingSummary            WebhookReplaceWebhookResponseDataEvent = "meeting.summary"
)

func (r WebhookReplaceWebhookResponseDataEvent) IsKnown() bool {
	switch r {
	case WebhookReplaceWebhookResponseDataEventMeetingStarted, WebhookReplaceWebhookResponseDataEventMeetingEnded, WebhookReplaceWebhookResponseDataEventMeetingParticipantJoined, WebhookReplaceWebhookResponseDataEventMeetingParticipantLeft, WebhookReplaceWebhookResponseDataEventMeetingChatSynced, WebhookReplaceWebhookResponseDataEventRecordingStatusUpdate, WebhookReplaceWebhookResponseDataEventLivestreamingStatusUpdate, WebhookReplaceWebhookResponseDataEventMeetingTranscript, WebhookReplaceWebhookResponseDataEventMeetingSummary:
		return true
	}
	return false
}

type WebhookNewWebhookParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Events that this webhook will get triggered by
	Events param.Field[[]WebhookNewWebhookParamsEvent] `json:"events,required"`
	// Name of the webhook
	Name param.Field[string] `json:"name,required"`
	// URL this webhook will send events to
	URL param.Field[string] `json:"url,required" format:"uri"`
	// Set whether or not the webhook should be active when created
	Enabled param.Field[bool] `json:"enabled"`
}

func (r WebhookNewWebhookParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookNewWebhookParamsEvent string

const (
	WebhookNewWebhookParamsEventMeetingStarted            WebhookNewWebhookParamsEvent = "meeting.started"
	WebhookNewWebhookParamsEventMeetingEnded              WebhookNewWebhookParamsEvent = "meeting.ended"
	WebhookNewWebhookParamsEventMeetingParticipantJoined  WebhookNewWebhookParamsEvent = "meeting.participantJoined"
	WebhookNewWebhookParamsEventMeetingParticipantLeft    WebhookNewWebhookParamsEvent = "meeting.participantLeft"
	WebhookNewWebhookParamsEventMeetingChatSynced         WebhookNewWebhookParamsEvent = "meeting.chatSynced"
	WebhookNewWebhookParamsEventRecordingStatusUpdate     WebhookNewWebhookParamsEvent = "recording.statusUpdate"
	WebhookNewWebhookParamsEventLivestreamingStatusUpdate WebhookNewWebhookParamsEvent = "livestreaming.statusUpdate"
	WebhookNewWebhookParamsEventMeetingTranscript         WebhookNewWebhookParamsEvent = "meeting.transcript"
	WebhookNewWebhookParamsEventMeetingSummary            WebhookNewWebhookParamsEvent = "meeting.summary"
)

func (r WebhookNewWebhookParamsEvent) IsKnown() bool {
	switch r {
	case WebhookNewWebhookParamsEventMeetingStarted, WebhookNewWebhookParamsEventMeetingEnded, WebhookNewWebhookParamsEventMeetingParticipantJoined, WebhookNewWebhookParamsEventMeetingParticipantLeft, WebhookNewWebhookParamsEventMeetingChatSynced, WebhookNewWebhookParamsEventRecordingStatusUpdate, WebhookNewWebhookParamsEventLivestreamingStatusUpdate, WebhookNewWebhookParamsEventMeetingTranscript, WebhookNewWebhookParamsEventMeetingSummary:
		return true
	}
	return false
}

type WebhookDeleteWebhookParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type WebhookEditWebhookParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	Enabled   param.Field[bool]   `json:"enabled"`
	// Events that the webhook will get triggered by
	Events param.Field[[]WebhookEditWebhookParamsEvent] `json:"events"`
	// Name of the webhook
	Name param.Field[string] `json:"name"`
	// URL the webhook will send events to
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r WebhookEditWebhookParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookEditWebhookParamsEvent string

const (
	WebhookEditWebhookParamsEventMeetingStarted            WebhookEditWebhookParamsEvent = "meeting.started"
	WebhookEditWebhookParamsEventMeetingEnded              WebhookEditWebhookParamsEvent = "meeting.ended"
	WebhookEditWebhookParamsEventMeetingParticipantJoined  WebhookEditWebhookParamsEvent = "meeting.participantJoined"
	WebhookEditWebhookParamsEventMeetingParticipantLeft    WebhookEditWebhookParamsEvent = "meeting.participantLeft"
	WebhookEditWebhookParamsEventRecordingStatusUpdate     WebhookEditWebhookParamsEvent = "recording.statusUpdate"
	WebhookEditWebhookParamsEventLivestreamingStatusUpdate WebhookEditWebhookParamsEvent = "livestreaming.statusUpdate"
	WebhookEditWebhookParamsEventMeetingChatSynced         WebhookEditWebhookParamsEvent = "meeting.chatSynced"
	WebhookEditWebhookParamsEventMeetingTranscript         WebhookEditWebhookParamsEvent = "meeting.transcript"
	WebhookEditWebhookParamsEventMeetingSummary            WebhookEditWebhookParamsEvent = "meeting.summary"
)

func (r WebhookEditWebhookParamsEvent) IsKnown() bool {
	switch r {
	case WebhookEditWebhookParamsEventMeetingStarted, WebhookEditWebhookParamsEventMeetingEnded, WebhookEditWebhookParamsEventMeetingParticipantJoined, WebhookEditWebhookParamsEventMeetingParticipantLeft, WebhookEditWebhookParamsEventRecordingStatusUpdate, WebhookEditWebhookParamsEventLivestreamingStatusUpdate, WebhookEditWebhookParamsEventMeetingChatSynced, WebhookEditWebhookParamsEventMeetingTranscript, WebhookEditWebhookParamsEventMeetingSummary:
		return true
	}
	return false
}

type WebhookGetWebhookByIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type WebhookGetWebhooksParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type WebhookReplaceWebhookParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Events that this webhook will get triggered by
	Events param.Field[[]WebhookReplaceWebhookParamsEvent] `json:"events,required"`
	// Name of the webhook
	Name param.Field[string] `json:"name,required"`
	// URL this webhook will send events to
	URL param.Field[string] `json:"url,required" format:"uri"`
	// Set whether or not the webhook should be active when created
	Enabled param.Field[bool] `json:"enabled"`
}

func (r WebhookReplaceWebhookParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookReplaceWebhookParamsEvent string

const (
	WebhookReplaceWebhookParamsEventMeetingStarted            WebhookReplaceWebhookParamsEvent = "meeting.started"
	WebhookReplaceWebhookParamsEventMeetingEnded              WebhookReplaceWebhookParamsEvent = "meeting.ended"
	WebhookReplaceWebhookParamsEventMeetingParticipantJoined  WebhookReplaceWebhookParamsEvent = "meeting.participantJoined"
	WebhookReplaceWebhookParamsEventMeetingParticipantLeft    WebhookReplaceWebhookParamsEvent = "meeting.participantLeft"
	WebhookReplaceWebhookParamsEventMeetingChatSynced         WebhookReplaceWebhookParamsEvent = "meeting.chatSynced"
	WebhookReplaceWebhookParamsEventRecordingStatusUpdate     WebhookReplaceWebhookParamsEvent = "recording.statusUpdate"
	WebhookReplaceWebhookParamsEventLivestreamingStatusUpdate WebhookReplaceWebhookParamsEvent = "livestreaming.statusUpdate"
	WebhookReplaceWebhookParamsEventMeetingTranscript         WebhookReplaceWebhookParamsEvent = "meeting.transcript"
	WebhookReplaceWebhookParamsEventMeetingSummary            WebhookReplaceWebhookParamsEvent = "meeting.summary"
)

func (r WebhookReplaceWebhookParamsEvent) IsKnown() bool {
	switch r {
	case WebhookReplaceWebhookParamsEventMeetingStarted, WebhookReplaceWebhookParamsEventMeetingEnded, WebhookReplaceWebhookParamsEventMeetingParticipantJoined, WebhookReplaceWebhookParamsEventMeetingParticipantLeft, WebhookReplaceWebhookParamsEventMeetingChatSynced, WebhookReplaceWebhookParamsEventRecordingStatusUpdate, WebhookReplaceWebhookParamsEventLivestreamingStatusUpdate, WebhookReplaceWebhookParamsEventMeetingTranscript, WebhookReplaceWebhookParamsEventMeetingSummary:
		return true
	}
	return false
}
