// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package realtime_kit

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
)

// ActiveSessionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveSessionService] method instead.
type ActiveSessionService struct {
	Options []option.RequestOption
}

// NewActiveSessionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveSessionService(opts ...option.RequestOption) (r *ActiveSessionService) {
	r = &ActiveSessionService{}
	r.Options = opts
	return
}

// Creates a new poll in an active session for the given meeting ID.
func (r *ActiveSessionService) NewPoll(ctx context.Context, appID string, meetingID string, params ActiveSessionNewPollParams, opts ...option.RequestOption) (res *ActiveSessionNewPollResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/active-session/poll", params.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns details of an ongoing active session for the given meeting ID.
func (r *ActiveSessionService) GetActiveSession(ctx context.Context, appID string, meetingID string, query ActiveSessionGetActiveSessionParams, opts ...option.RequestOption) (res *ActiveSessionGetActiveSessionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/active-session", query.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Kicks all participants from an active session for the given meeting ID.
func (r *ActiveSessionService) KickAllParticipants(ctx context.Context, appID string, meetingID string, body ActiveSessionKickAllParticipantsParams, opts ...option.RequestOption) (res *ActiveSessionKickAllParticipantsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/active-session/kick-all", body.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Kicks one or more participants from an active session using user ID or custom
// participant ID.
func (r *ActiveSessionService) KickParticipants(ctx context.Context, appID string, meetingID string, params ActiveSessionKickParticipantsParams, opts ...option.RequestOption) (res *ActiveSessionKickParticipantsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/%s/meetings/%s/active-session/kick", params.AccountID, appID, meetingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ActiveSessionNewPollResponse struct {
	Data    ActiveSessionNewPollResponseData `json:"data"`
	Success bool                             `json:"success"`
	JSON    activeSessionNewPollResponseJSON `json:"-"`
}

// activeSessionNewPollResponseJSON contains the JSON metadata for the struct
// [ActiveSessionNewPollResponse]
type activeSessionNewPollResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActiveSessionNewPollResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activeSessionNewPollResponseJSON) RawJSON() string {
	return r.raw
}

type ActiveSessionNewPollResponseData struct {
	Action string                               `json:"action"`
	Poll   ActiveSessionNewPollResponseDataPoll `json:"poll"`
	JSON   activeSessionNewPollResponseDataJSON `json:"-"`
}

// activeSessionNewPollResponseDataJSON contains the JSON metadata for the struct
// [ActiveSessionNewPollResponseData]
type activeSessionNewPollResponseDataJSON struct {
	Action      apijson.Field
	Poll        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActiveSessionNewPollResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activeSessionNewPollResponseDataJSON) RawJSON() string {
	return r.raw
}

type ActiveSessionNewPollResponseDataPoll struct {
	// ID of the poll
	ID string `json:"id,required"`
	// Answer options
	Options []ActiveSessionNewPollResponseDataPollOption `json:"options,required"`
	// Question asked by the poll
	Question  string                                   `json:"question,required"`
	Anonymous bool                                     `json:"anonymous"`
	CreatedBy string                                   `json:"created_by"`
	HideVotes bool                                     `json:"hide_votes"`
	Voted     []string                                 `json:"voted"`
	JSON      activeSessionNewPollResponseDataPollJSON `json:"-"`
}

// activeSessionNewPollResponseDataPollJSON contains the JSON metadata for the
// struct [ActiveSessionNewPollResponseDataPoll]
type activeSessionNewPollResponseDataPollJSON struct {
	ID          apijson.Field
	Options     apijson.Field
	Question    apijson.Field
	Anonymous   apijson.Field
	CreatedBy   apijson.Field
	HideVotes   apijson.Field
	Voted       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActiveSessionNewPollResponseDataPoll) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activeSessionNewPollResponseDataPollJSON) RawJSON() string {
	return r.raw
}

type ActiveSessionNewPollResponseDataPollOption struct {
	Count float64 `json:"count,required"`
	// Text of the answer option
	Text  string                                            `json:"text,required"`
	Votes []ActiveSessionNewPollResponseDataPollOptionsVote `json:"votes,required"`
	JSON  activeSessionNewPollResponseDataPollOptionJSON    `json:"-"`
}

// activeSessionNewPollResponseDataPollOptionJSON contains the JSON metadata for
// the struct [ActiveSessionNewPollResponseDataPollOption]
type activeSessionNewPollResponseDataPollOptionJSON struct {
	Count       apijson.Field
	Text        apijson.Field
	Votes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActiveSessionNewPollResponseDataPollOption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activeSessionNewPollResponseDataPollOptionJSON) RawJSON() string {
	return r.raw
}

type ActiveSessionNewPollResponseDataPollOptionsVote struct {
	ID   string                                              `json:"id,required"`
	Name string                                              `json:"name,required"`
	JSON activeSessionNewPollResponseDataPollOptionsVoteJSON `json:"-"`
}

// activeSessionNewPollResponseDataPollOptionsVoteJSON contains the JSON metadata
// for the struct [ActiveSessionNewPollResponseDataPollOptionsVote]
type activeSessionNewPollResponseDataPollOptionsVoteJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActiveSessionNewPollResponseDataPollOptionsVote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activeSessionNewPollResponseDataPollOptionsVoteJSON) RawJSON() string {
	return r.raw
}

type ActiveSessionGetActiveSessionResponse struct {
	Data    ActiveSessionGetActiveSessionResponseData `json:"data"`
	Success bool                                      `json:"success"`
	JSON    activeSessionGetActiveSessionResponseJSON `json:"-"`
}

// activeSessionGetActiveSessionResponseJSON contains the JSON metadata for the
// struct [ActiveSessionGetActiveSessionResponse]
type activeSessionGetActiveSessionResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActiveSessionGetActiveSessionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activeSessionGetActiveSessionResponseJSON) RawJSON() string {
	return r.raw
}

type ActiveSessionGetActiveSessionResponseData struct {
	// ID of the session
	ID string `json:"id,required"`
	// ID of the meeting this session is associated with. In the case of V2 meetings,
	// it is always a UUID. In V1 meetings, it is a room name of the form
	// `abcdef-ghijkl`
	AssociatedID string `json:"associated_id,required"`
	// timestamp when session created
	CreatedAt string `json:"created_at,required"`
	// number of participants currently in the session
	LiveParticipants float64 `json:"live_participants,required"`
	// number of maximum participants that were in the session
	MaxConcurrentParticipants float64 `json:"max_concurrent_participants,required"`
	// Title of the meeting this session belongs to
	MeetingDisplayName string `json:"meeting_display_name,required"`
	// number of minutes consumed since the session started
	MinutesConsumed float64 `json:"minutes_consumed,required"`
	// App id that hosted this session
	OrganizationID string `json:"organization_id,required"`
	// timestamp when session started
	StartedAt string `json:"started_at,required"`
	// current status of session
	Status ActiveSessionGetActiveSessionResponseDataStatus `json:"status,required"`
	// type of session
	Type ActiveSessionGetActiveSessionResponseDataType `json:"type,required"`
	// timestamp when session was last updated
	UpdatedAt     string        `json:"updated_at,required"`
	BreakoutRooms []interface{} `json:"breakout_rooms"`
	// timestamp when session ended
	EndedAt string `json:"ended_at"`
	// Any meta data about session.
	Meta interface{}                                   `json:"meta"`
	JSON activeSessionGetActiveSessionResponseDataJSON `json:"-"`
}

// activeSessionGetActiveSessionResponseDataJSON contains the JSON metadata for the
// struct [ActiveSessionGetActiveSessionResponseData]
type activeSessionGetActiveSessionResponseDataJSON struct {
	ID                        apijson.Field
	AssociatedID              apijson.Field
	CreatedAt                 apijson.Field
	LiveParticipants          apijson.Field
	MaxConcurrentParticipants apijson.Field
	MeetingDisplayName        apijson.Field
	MinutesConsumed           apijson.Field
	OrganizationID            apijson.Field
	StartedAt                 apijson.Field
	Status                    apijson.Field
	Type                      apijson.Field
	UpdatedAt                 apijson.Field
	BreakoutRooms             apijson.Field
	EndedAt                   apijson.Field
	Meta                      apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ActiveSessionGetActiveSessionResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activeSessionGetActiveSessionResponseDataJSON) RawJSON() string {
	return r.raw
}

// current status of session
type ActiveSessionGetActiveSessionResponseDataStatus string

const (
	ActiveSessionGetActiveSessionResponseDataStatusLive  ActiveSessionGetActiveSessionResponseDataStatus = "LIVE"
	ActiveSessionGetActiveSessionResponseDataStatusEnded ActiveSessionGetActiveSessionResponseDataStatus = "ENDED"
)

func (r ActiveSessionGetActiveSessionResponseDataStatus) IsKnown() bool {
	switch r {
	case ActiveSessionGetActiveSessionResponseDataStatusLive, ActiveSessionGetActiveSessionResponseDataStatusEnded:
		return true
	}
	return false
}

// type of session
type ActiveSessionGetActiveSessionResponseDataType string

const (
	ActiveSessionGetActiveSessionResponseDataTypeMeeting     ActiveSessionGetActiveSessionResponseDataType = "meeting"
	ActiveSessionGetActiveSessionResponseDataTypeLivestream  ActiveSessionGetActiveSessionResponseDataType = "livestream"
	ActiveSessionGetActiveSessionResponseDataTypeParticipant ActiveSessionGetActiveSessionResponseDataType = "participant"
)

func (r ActiveSessionGetActiveSessionResponseDataType) IsKnown() bool {
	switch r {
	case ActiveSessionGetActiveSessionResponseDataTypeMeeting, ActiveSessionGetActiveSessionResponseDataTypeLivestream, ActiveSessionGetActiveSessionResponseDataTypeParticipant:
		return true
	}
	return false
}

type ActiveSessionKickAllParticipantsResponse struct {
	Data    ActiveSessionKickAllParticipantsResponseData `json:"data"`
	Success bool                                         `json:"success"`
	JSON    activeSessionKickAllParticipantsResponseJSON `json:"-"`
}

// activeSessionKickAllParticipantsResponseJSON contains the JSON metadata for the
// struct [ActiveSessionKickAllParticipantsResponse]
type activeSessionKickAllParticipantsResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActiveSessionKickAllParticipantsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activeSessionKickAllParticipantsResponseJSON) RawJSON() string {
	return r.raw
}

type ActiveSessionKickAllParticipantsResponseData struct {
	Action                  string                                           `json:"action"`
	KickedParticipantsCount float64                                          `json:"kicked_participants_count"`
	JSON                    activeSessionKickAllParticipantsResponseDataJSON `json:"-"`
}

// activeSessionKickAllParticipantsResponseDataJSON contains the JSON metadata for
// the struct [ActiveSessionKickAllParticipantsResponseData]
type activeSessionKickAllParticipantsResponseDataJSON struct {
	Action                  apijson.Field
	KickedParticipantsCount apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ActiveSessionKickAllParticipantsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activeSessionKickAllParticipantsResponseDataJSON) RawJSON() string {
	return r.raw
}

type ActiveSessionKickParticipantsResponse struct {
	Data    ActiveSessionKickParticipantsResponseData `json:"data"`
	Success bool                                      `json:"success"`
	JSON    activeSessionKickParticipantsResponseJSON `json:"-"`
}

// activeSessionKickParticipantsResponseJSON contains the JSON metadata for the
// struct [ActiveSessionKickParticipantsResponse]
type activeSessionKickParticipantsResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActiveSessionKickParticipantsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activeSessionKickParticipantsResponseJSON) RawJSON() string {
	return r.raw
}

type ActiveSessionKickParticipantsResponseData struct {
	Action       string                                                 `json:"action"`
	Participants []ActiveSessionKickParticipantsResponseDataParticipant `json:"participants"`
	JSON         activeSessionKickParticipantsResponseDataJSON          `json:"-"`
}

// activeSessionKickParticipantsResponseDataJSON contains the JSON metadata for the
// struct [ActiveSessionKickParticipantsResponseData]
type activeSessionKickParticipantsResponseDataJSON struct {
	Action       apijson.Field
	Participants apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ActiveSessionKickParticipantsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activeSessionKickParticipantsResponseDataJSON) RawJSON() string {
	return r.raw
}

type ActiveSessionKickParticipantsResponseDataParticipant struct {
	// ID of the session participant
	ID        string `json:"id,required"`
	CreatedAt string `json:"created_at,required"`
	UpdatedAt string `json:"updated_at,required"`
	// Email of the session participant.
	Email string `json:"email"`
	// Name of the session participant.
	Name string `json:"name"`
	// A URL pointing to a picture of the participant.
	Picture string                                                   `json:"picture"`
	JSON    activeSessionKickParticipantsResponseDataParticipantJSON `json:"-"`
}

// activeSessionKickParticipantsResponseDataParticipantJSON contains the JSON
// metadata for the struct [ActiveSessionKickParticipantsResponseDataParticipant]
type activeSessionKickParticipantsResponseDataParticipantJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Picture     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActiveSessionKickParticipantsResponseDataParticipant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activeSessionKickParticipantsResponseDataParticipantJSON) RawJSON() string {
	return r.raw
}

type ActiveSessionNewPollParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Different options for the question
	Options param.Field[[]string] `json:"options,required"`
	// Question of the poll
	Question param.Field[string] `json:"question,required"`
	// if voters on a poll are anonymous
	Anonymous param.Field[bool] `json:"anonymous"`
	// if votes on an option are visible before a person votes
	HideVotes param.Field[bool] `json:"hide_votes"`
}

func (r ActiveSessionNewPollParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ActiveSessionGetActiveSessionParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ActiveSessionKickAllParticipantsParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ActiveSessionKickParticipantsParams struct {
	// The account identifier tag.
	AccountID            param.Field[string]   `path:"account_id,required"`
	CustomParticipantIDs param.Field[[]string] `json:"custom_participant_ids,required"`
	ParticipantIDs       param.Field[[]string] `json:"participant_ids,required"`
}

func (r ActiveSessionKickParticipantsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
