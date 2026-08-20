// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_sending

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// EmailSendingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSendingService] method instead.
type EmailSendingService struct {
	Options    []option.RequestOption
	Subdomains *SubdomainService
}

// NewEmailSendingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailSendingService(opts ...option.RequestOption) (r *EmailSendingService) {
	r = &EmailSendingService{}
	r.Options = opts
	r.Subdomains = NewSubdomainService(opts...)
	return
}

// Send an email for the specified account using the structured builder. Provide
// the sender, recipients, subject, and at least one of text or html; attachments
// are optional.
func (r *EmailSendingService) Send(ctx context.Context, params EmailSendingSendParams, opts ...option.RequestOption) (res *EmailSendingSendResponse, err error) {
	var env EmailSendingSendResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/send", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Send a raw RFC 5322 (MIME) email for the specified account. Provide the full
// MIME message plus the SMTP envelope (from and recipients).
func (r *EmailSendingService) SendRaw(ctx context.Context, params EmailSendingSendRawParams, opts ...option.RequestOption) (res *EmailSendingSendRawResponse, err error) {
	var env EmailSendingSendRawResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/send_raw", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type EmailSendingSendResponse struct {
	// Email addresses to which the message was delivered immediately.
	Delivered []string `json:"delivered" api:"required"`
	// Message ID of the sent email.
	MessageID string `json:"message_id" api:"required"`
	// Email addresses that permanently bounced.
	PermanentBounces []string `json:"permanent_bounces" api:"required"`
	// Email addresses for which delivery was queued for later.
	Queued []string                     `json:"queued" api:"required"`
	JSON   emailSendingSendResponseJSON `json:"-"`
}

// emailSendingSendResponseJSON contains the JSON metadata for the struct
// [EmailSendingSendResponse]
type emailSendingSendResponseJSON struct {
	Delivered        apijson.Field
	MessageID        apijson.Field
	PermanentBounces apijson.Field
	Queued           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EmailSendingSendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSendingSendResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSendingSendRawResponse struct {
	// Email addresses to which the message was delivered immediately.
	Delivered []string `json:"delivered" api:"required"`
	// Message ID of the sent email.
	MessageID string `json:"message_id" api:"required"`
	// Email addresses that permanently bounced.
	PermanentBounces []string `json:"permanent_bounces" api:"required"`
	// Email addresses for which delivery was queued for later.
	Queued []string                        `json:"queued" api:"required"`
	JSON   emailSendingSendRawResponseJSON `json:"-"`
}

// emailSendingSendRawResponseJSON contains the JSON metadata for the struct
// [EmailSendingSendRawResponse]
type emailSendingSendRawResponseJSON struct {
	Delivered        apijson.Field
	MessageID        apijson.Field
	PermanentBounces apijson.Field
	Queued           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EmailSendingSendRawResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSendingSendRawResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSendingSendParams struct {
	// Identifier of the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Sender email address. Either a plain string or an object with address and name.
	From param.Field[EmailSendingSendParamsFromUnion] `json:"from" api:"required"`
	// Email subject line.
	Subject param.Field[string] `json:"subject" api:"required"`
	// File attachments and inline images.
	Attachments param.Field[[]EmailSendingSendParamsAttachmentUnion] `json:"attachments"`
	// Recipient(s). Optional if cc or bcc is provided. A single email string, a named
	// address object, or an array of either.
	Bcc param.Field[EmailSendingSendParamsBccUnion] `json:"bcc"`
	// Recipient(s). Optional if cc or bcc is provided. A single email string, a named
	// address object, or an array of either.
	Cc param.Field[EmailSendingSendParamsCcUnion] `json:"cc"`
	// Custom email headers as key-value pairs.
	Headers param.Field[map[string]string] `json:"headers"`
	// HTML body of the email. Provide at least one of text or html (non-empty).
	HTML param.Field[string] `json:"html"`
	// Reply-to address. Either a plain string or an object with address and name.
	ReplyTo param.Field[EmailSendingSendParamsReplyToUnion] `json:"reply_to"`
	// Plain text body of the email. Provide at least one of text or html (non-empty).
	Text param.Field[string] `json:"text"`
	// Recipient(s). Optional if cc or bcc is provided. A single email string, a named
	// address object, or an array of either.
	To param.Field[EmailSendingSendParamsToUnion] `json:"to"`
}

func (r EmailSendingSendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sender email address. Either a plain string or an object with address and name.
//
// Satisfied by [shared.UnionString],
// [email_sending.EmailSendingSendParamsFromEmailSendingEmailAddressObject].
type EmailSendingSendParamsFromUnion interface {
	ImplementsEmailSendingSendParamsFromUnion()
}

type EmailSendingSendParamsFromEmailSendingEmailAddressObject struct {
	// Email address (e.g., 'user@example.com').
	Address param.Field[string] `json:"address" api:"required"`
	// Display name for the email address (e.g., 'John Doe'). Optional; set to null or
	// leave it unset to send the address on its own.
	Name param.Field[string] `json:"name"`
}

func (r EmailSendingSendParamsFromEmailSendingEmailAddressObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailSendingSendParamsFromEmailSendingEmailAddressObject) ImplementsEmailSendingSendParamsFromUnion() {
}

type EmailSendingSendParamsAttachment struct {
	// Base64-encoded content of the attachment.
	Content param.Field[string] `json:"content" api:"required"`
	// Use 'inline' to embed the attachment in the email body.
	Disposition param.Field[EmailSendingSendParamsAttachmentsDisposition] `json:"disposition" api:"required"`
	// Filename for the attachment.
	Filename param.Field[string] `json:"filename" api:"required"`
	// MIME type of the attachment (e.g., 'image/png', 'text/plain').
	Type param.Field[string] `json:"type" api:"required"`
	// Content ID used to reference this attachment in HTML via cid: URI (e.g.,
	// <img src="cid:logo">).
	ContentID param.Field[string] `json:"content_id"`
}

func (r EmailSendingSendParamsAttachment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailSendingSendParamsAttachment) implementsEmailSendingSendParamsAttachmentUnion() {}

// Satisfied by
// [email_sending.EmailSendingSendParamsAttachmentsEmailSendingEmailInlineAttachment],
// [email_sending.EmailSendingSendParamsAttachmentsEmailSendingEmailAttachment],
// [EmailSendingSendParamsAttachment].
type EmailSendingSendParamsAttachmentUnion interface {
	implementsEmailSendingSendParamsAttachmentUnion()
}

type EmailSendingSendParamsAttachmentsEmailSendingEmailInlineAttachment struct {
	// Base64-encoded content of the attachment.
	Content param.Field[string] `json:"content" api:"required"`
	// Content ID used to reference this attachment in HTML via cid: URI (e.g.,
	// <img src="cid:logo">).
	ContentID param.Field[string] `json:"content_id" api:"required"`
	// Use 'inline' to embed the attachment in the email body.
	Disposition param.Field[EmailSendingSendParamsAttachmentsEmailSendingEmailInlineAttachmentDisposition] `json:"disposition" api:"required"`
	// Filename for the attachment.
	Filename param.Field[string] `json:"filename" api:"required"`
	// MIME type of the attachment (e.g., 'image/png', 'text/plain').
	Type param.Field[string] `json:"type" api:"required"`
}

func (r EmailSendingSendParamsAttachmentsEmailSendingEmailInlineAttachment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailSendingSendParamsAttachmentsEmailSendingEmailInlineAttachment) implementsEmailSendingSendParamsAttachmentUnion() {
}

// Use 'inline' to embed the attachment in the email body.
type EmailSendingSendParamsAttachmentsEmailSendingEmailInlineAttachmentDisposition string

const (
	EmailSendingSendParamsAttachmentsEmailSendingEmailInlineAttachmentDispositionInline EmailSendingSendParamsAttachmentsEmailSendingEmailInlineAttachmentDisposition = "inline"
)

func (r EmailSendingSendParamsAttachmentsEmailSendingEmailInlineAttachmentDisposition) IsKnown() bool {
	switch r {
	case EmailSendingSendParamsAttachmentsEmailSendingEmailInlineAttachmentDispositionInline:
		return true
	}
	return false
}

type EmailSendingSendParamsAttachmentsEmailSendingEmailAttachment struct {
	// Base64-encoded content of the attachment.
	Content param.Field[string] `json:"content" api:"required"`
	// Use 'attachment' for a standard file attachment.
	Disposition param.Field[EmailSendingSendParamsAttachmentsEmailSendingEmailAttachmentDisposition] `json:"disposition" api:"required"`
	// Filename for the attachment.
	Filename param.Field[string] `json:"filename" api:"required"`
	// MIME type of the attachment (e.g., 'application/pdf', 'text/plain').
	Type param.Field[string] `json:"type" api:"required"`
}

func (r EmailSendingSendParamsAttachmentsEmailSendingEmailAttachment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailSendingSendParamsAttachmentsEmailSendingEmailAttachment) implementsEmailSendingSendParamsAttachmentUnion() {
}

// Use 'attachment' for a standard file attachment.
type EmailSendingSendParamsAttachmentsEmailSendingEmailAttachmentDisposition string

const (
	EmailSendingSendParamsAttachmentsEmailSendingEmailAttachmentDispositionAttachment EmailSendingSendParamsAttachmentsEmailSendingEmailAttachmentDisposition = "attachment"
)

func (r EmailSendingSendParamsAttachmentsEmailSendingEmailAttachmentDisposition) IsKnown() bool {
	switch r {
	case EmailSendingSendParamsAttachmentsEmailSendingEmailAttachmentDispositionAttachment:
		return true
	}
	return false
}

// Use 'inline' to embed the attachment in the email body.
type EmailSendingSendParamsAttachmentsDisposition string

const (
	EmailSendingSendParamsAttachmentsDispositionInline     EmailSendingSendParamsAttachmentsDisposition = "inline"
	EmailSendingSendParamsAttachmentsDispositionAttachment EmailSendingSendParamsAttachmentsDisposition = "attachment"
)

func (r EmailSendingSendParamsAttachmentsDisposition) IsKnown() bool {
	switch r {
	case EmailSendingSendParamsAttachmentsDispositionInline, EmailSendingSendParamsAttachmentsDispositionAttachment:
		return true
	}
	return false
}

// Recipient(s). Optional if cc or bcc is provided. A single email string, a named
// address object, or an array of either.
//
// Satisfied by [shared.UnionString],
// [email_sending.EmailSendingSendParamsBccEmailSendingEmailAddressObject],
// [email_sending.EmailSendingSendParamsBccArray].
type EmailSendingSendParamsBccUnion interface {
	ImplementsEmailSendingSendParamsBccUnion()
}

type EmailSendingSendParamsBccEmailSendingEmailAddressObject struct {
	// Email address (e.g., 'user@example.com').
	Address param.Field[string] `json:"address" api:"required"`
	// Display name for the email address (e.g., 'John Doe'). Optional; set to null or
	// leave it unset to send the address on its own.
	Name param.Field[string] `json:"name"`
}

func (r EmailSendingSendParamsBccEmailSendingEmailAddressObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailSendingSendParamsBccEmailSendingEmailAddressObject) ImplementsEmailSendingSendParamsBccUnion() {
}

type EmailSendingSendParamsBccArray []EmailSendingSendParamsBccArrayItemUnion

func (r EmailSendingSendParamsBccArray) ImplementsEmailSendingSendParamsBccUnion() {}

// An email address as a plain string.
//
// Satisfied by [shared.UnionString],
// [email_sending.EmailSendingSendParamsBccArrayEmailSendingEmailAddressObject].
type EmailSendingSendParamsBccArrayItemUnion interface {
	ImplementsEmailSendingSendParamsBccArrayItemUnion()
}

type EmailSendingSendParamsBccArrayEmailSendingEmailAddressObject struct {
	// Email address (e.g., 'user@example.com').
	Address param.Field[string] `json:"address" api:"required"`
	// Display name for the email address (e.g., 'John Doe'). Optional; set to null or
	// leave it unset to send the address on its own.
	Name param.Field[string] `json:"name"`
}

func (r EmailSendingSendParamsBccArrayEmailSendingEmailAddressObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailSendingSendParamsBccArrayEmailSendingEmailAddressObject) ImplementsEmailSendingSendParamsBccArrayItemUnion() {
}

// Recipient(s). Optional if cc or bcc is provided. A single email string, a named
// address object, or an array of either.
//
// Satisfied by [shared.UnionString],
// [email_sending.EmailSendingSendParamsCcEmailSendingEmailAddressObject],
// [email_sending.EmailSendingSendParamsCcArray].
type EmailSendingSendParamsCcUnion interface {
	ImplementsEmailSendingSendParamsCcUnion()
}

type EmailSendingSendParamsCcEmailSendingEmailAddressObject struct {
	// Email address (e.g., 'user@example.com').
	Address param.Field[string] `json:"address" api:"required"`
	// Display name for the email address (e.g., 'John Doe'). Optional; set to null or
	// leave it unset to send the address on its own.
	Name param.Field[string] `json:"name"`
}

func (r EmailSendingSendParamsCcEmailSendingEmailAddressObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailSendingSendParamsCcEmailSendingEmailAddressObject) ImplementsEmailSendingSendParamsCcUnion() {
}

type EmailSendingSendParamsCcArray []EmailSendingSendParamsCcArrayItemUnion

func (r EmailSendingSendParamsCcArray) ImplementsEmailSendingSendParamsCcUnion() {}

// An email address as a plain string.
//
// Satisfied by [shared.UnionString],
// [email_sending.EmailSendingSendParamsCcArrayEmailSendingEmailAddressObject].
type EmailSendingSendParamsCcArrayItemUnion interface {
	ImplementsEmailSendingSendParamsCcArrayItemUnion()
}

type EmailSendingSendParamsCcArrayEmailSendingEmailAddressObject struct {
	// Email address (e.g., 'user@example.com').
	Address param.Field[string] `json:"address" api:"required"`
	// Display name for the email address (e.g., 'John Doe'). Optional; set to null or
	// leave it unset to send the address on its own.
	Name param.Field[string] `json:"name"`
}

func (r EmailSendingSendParamsCcArrayEmailSendingEmailAddressObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailSendingSendParamsCcArrayEmailSendingEmailAddressObject) ImplementsEmailSendingSendParamsCcArrayItemUnion() {
}

// Reply-to address. Either a plain string or an object with address and name.
//
// Satisfied by [shared.UnionString],
// [email_sending.EmailSendingSendParamsReplyToEmailSendingEmailAddressObject].
type EmailSendingSendParamsReplyToUnion interface {
	ImplementsEmailSendingSendParamsReplyToUnion()
}

type EmailSendingSendParamsReplyToEmailSendingEmailAddressObject struct {
	// Email address (e.g., 'user@example.com').
	Address param.Field[string] `json:"address" api:"required"`
	// Display name for the email address (e.g., 'John Doe'). Optional; set to null or
	// leave it unset to send the address on its own.
	Name param.Field[string] `json:"name"`
}

func (r EmailSendingSendParamsReplyToEmailSendingEmailAddressObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailSendingSendParamsReplyToEmailSendingEmailAddressObject) ImplementsEmailSendingSendParamsReplyToUnion() {
}

// Recipient(s). Optional if cc or bcc is provided. A single email string, a named
// address object, or an array of either.
//
// Satisfied by [shared.UnionString],
// [email_sending.EmailSendingSendParamsToEmailSendingEmailAddressObject],
// [email_sending.EmailSendingSendParamsToArray].
type EmailSendingSendParamsToUnion interface {
	ImplementsEmailSendingSendParamsToUnion()
}

type EmailSendingSendParamsToEmailSendingEmailAddressObject struct {
	// Email address (e.g., 'user@example.com').
	Address param.Field[string] `json:"address" api:"required"`
	// Display name for the email address (e.g., 'John Doe'). Optional; set to null or
	// leave it unset to send the address on its own.
	Name param.Field[string] `json:"name"`
}

func (r EmailSendingSendParamsToEmailSendingEmailAddressObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailSendingSendParamsToEmailSendingEmailAddressObject) ImplementsEmailSendingSendParamsToUnion() {
}

type EmailSendingSendParamsToArray []EmailSendingSendParamsToArrayItemUnion

func (r EmailSendingSendParamsToArray) ImplementsEmailSendingSendParamsToUnion() {}

// An email address as a plain string.
//
// Satisfied by [shared.UnionString],
// [email_sending.EmailSendingSendParamsToArrayEmailSendingEmailAddressObject].
type EmailSendingSendParamsToArrayItemUnion interface {
	ImplementsEmailSendingSendParamsToArrayItemUnion()
}

type EmailSendingSendParamsToArrayEmailSendingEmailAddressObject struct {
	// Email address (e.g., 'user@example.com').
	Address param.Field[string] `json:"address" api:"required"`
	// Display name for the email address (e.g., 'John Doe'). Optional; set to null or
	// leave it unset to send the address on its own.
	Name param.Field[string] `json:"name"`
}

func (r EmailSendingSendParamsToArrayEmailSendingEmailAddressObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailSendingSendParamsToArrayEmailSendingEmailAddressObject) ImplementsEmailSendingSendParamsToArrayItemUnion() {
}

type EmailSendingSendResponseEnvelope struct {
	Errors     []EmailSendingSendResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []EmailSendingSendResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     EmailSendingSendResponse                   `json:"result" api:"required"`
	Success    EmailSendingSendResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo EmailSendingSendResponseEnvelopeResultInfo `json:"result_info"`
	JSON       emailSendingSendResponseEnvelopeJSON       `json:"-"`
}

// emailSendingSendResponseEnvelopeJSON contains the JSON metadata for the struct
// [EmailSendingSendResponseEnvelope]
type emailSendingSendResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSendingSendResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSendingSendResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSendingSendResponseEnvelopeErrors struct {
	Code    float64                                    `json:"code" api:"required"`
	Message string                                     `json:"message" api:"required"`
	JSON    emailSendingSendResponseEnvelopeErrorsJSON `json:"-"`
}

// emailSendingSendResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [EmailSendingSendResponseEnvelopeErrors]
type emailSendingSendResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSendingSendResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSendingSendResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EmailSendingSendResponseEnvelopeMessages struct {
	Code    float64                                      `json:"code" api:"required"`
	Message string                                       `json:"message" api:"required"`
	JSON    emailSendingSendResponseEnvelopeMessagesJSON `json:"-"`
}

// emailSendingSendResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [EmailSendingSendResponseEnvelopeMessages]
type emailSendingSendResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSendingSendResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSendingSendResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type EmailSendingSendResponseEnvelopeSuccess bool

const (
	EmailSendingSendResponseEnvelopeSuccessTrue EmailSendingSendResponseEnvelopeSuccess = true
)

func (r EmailSendingSendResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EmailSendingSendResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type EmailSendingSendResponseEnvelopeResultInfo struct {
	Count      float64                                        `json:"count" api:"required"`
	PerPage    float64                                        `json:"per_page" api:"required"`
	TotalCount float64                                        `json:"total_count" api:"required"`
	Cursor     string                                         `json:"cursor"`
	Page       float64                                        `json:"page"`
	JSON       emailSendingSendResponseEnvelopeResultInfoJSON `json:"-"`
}

// emailSendingSendResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [EmailSendingSendResponseEnvelopeResultInfo]
type emailSendingSendResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	Cursor      apijson.Field
	Page        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSendingSendResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSendingSendResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSendingSendRawParams struct {
	// Identifier of the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Sender email address.
	From param.Field[string] `json:"from" api:"required"`
	// The full MIME-encoded email message. Should include standard RFC 5322 headers
	// such as From, To, Subject, and Content-Type. The from and recipients fields in
	// the request body control SMTP envelope routing; the From and To headers in the
	// MIME message control what the recipient's email client displays.
	MimeMessage param.Field[string] `json:"mime_message" api:"required"`
	// List of recipient email addresses.
	Recipients param.Field[[]string] `json:"recipients" api:"required"`
}

func (r EmailSendingSendRawParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmailSendingSendRawResponseEnvelope struct {
	Errors     []EmailSendingSendRawResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []EmailSendingSendRawResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     EmailSendingSendRawResponse                   `json:"result" api:"required"`
	Success    EmailSendingSendRawResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo EmailSendingSendRawResponseEnvelopeResultInfo `json:"result_info"`
	JSON       emailSendingSendRawResponseEnvelopeJSON       `json:"-"`
}

// emailSendingSendRawResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailSendingSendRawResponseEnvelope]
type emailSendingSendRawResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSendingSendRawResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSendingSendRawResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSendingSendRawResponseEnvelopeErrors struct {
	Code    float64                                       `json:"code" api:"required"`
	Message string                                        `json:"message" api:"required"`
	JSON    emailSendingSendRawResponseEnvelopeErrorsJSON `json:"-"`
}

// emailSendingSendRawResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [EmailSendingSendRawResponseEnvelopeErrors]
type emailSendingSendRawResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSendingSendRawResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSendingSendRawResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EmailSendingSendRawResponseEnvelopeMessages struct {
	Code    float64                                         `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	JSON    emailSendingSendRawResponseEnvelopeMessagesJSON `json:"-"`
}

// emailSendingSendRawResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [EmailSendingSendRawResponseEnvelopeMessages]
type emailSendingSendRawResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSendingSendRawResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSendingSendRawResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type EmailSendingSendRawResponseEnvelopeSuccess bool

const (
	EmailSendingSendRawResponseEnvelopeSuccessTrue EmailSendingSendRawResponseEnvelopeSuccess = true
)

func (r EmailSendingSendRawResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EmailSendingSendRawResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type EmailSendingSendRawResponseEnvelopeResultInfo struct {
	Count      float64                                           `json:"count" api:"required"`
	PerPage    float64                                           `json:"per_page" api:"required"`
	TotalCount float64                                           `json:"total_count" api:"required"`
	Cursor     string                                            `json:"cursor"`
	Page       float64                                           `json:"page"`
	JSON       emailSendingSendRawResponseEnvelopeResultInfoJSON `json:"-"`
}

// emailSendingSendRawResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [EmailSendingSendRawResponseEnvelopeResultInfo]
type emailSendingSendRawResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	Cursor      apijson.Field
	Page        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSendingSendRawResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSendingSendRawResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
