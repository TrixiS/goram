package goram

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"strconv"
)

// see Bot.GetUpdates(ctx, &GetUpdatesRequest{})
type GetUpdatesRequest struct {
	Offset         int64        `json:"offset,omitempty"`          // Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will be forgotten.
	Limit          int          `json:"limit,omitempty"`           // Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Timeout        int          `json:"timeout,omitempty"`         // Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	AllowedUpdates []UpdateType `json:"allowed_updates,omitempty"` // A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default). If not specified, the previous setting will be used. Please note that this parameter doesn't affect updates created before the call to getUpdates, so unwanted updates may be received for a short period of time.
}

func (r *GetUpdatesRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.Offset)
		fw, _ := w.CreateFormField("offset")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Limit)
		fw, _ := w.CreateFormField("limit")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Timeout)
		fw, _ := w.CreateFormField("timeout")
		fw.Write(b)
	}
	if r.AllowedUpdates != nil {
		{
			b, _ := json.Marshal(r.AllowedUpdates)
			fw, _ := w.CreateFormField("allowed_updates")
			fw.Write(b)
		}
	}
}

// see Bot.SetWebhook(ctx, &SetWebhookRequest{})
type SetWebhookRequest struct {
	Url                string       `json:"url"`                            // HTTPS URL to send updates to. Use an empty string to remove webhook integration
	Certificate        InputFile    `json:"certificate,omitempty"`          // Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	IpAddress          string       `json:"ip_address,omitempty"`           // The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	MaxConnections     int          `json:"max_connections,omitempty"`      // The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	AllowedUpdates     []UpdateType `json:"allowed_updates,omitempty"`      // A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default). If not specified, the previous setting will be used. Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	DropPendingUpdates bool         `json:"drop_pending_updates,omitempty"` // Pass True to drop all pending updates
	SecretToken        string       `json:"secret_token,omitempty"`         // A secret token to be sent in a header "X-Telegram-Bot-Api-Secret-Token" in every webhook request, 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a webhook set by you.
}

func (r *SetWebhookRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("url", r.Url)

	if r.Certificate.FileID != "" {
		w.WriteField("certificate", r.Certificate.FileID)
	} else if r.Certificate.Reader != nil {
		fw, _ := w.CreateFormFile("certificate", r.Certificate.Reader.Name())
		io.Copy(fw, r.Certificate.Reader)
	}
	if r.IpAddress != "" {
		w.WriteField("ip_address", r.IpAddress)
	}

	{
		b, _ := json.Marshal(r.MaxConnections)
		fw, _ := w.CreateFormField("max_connections")
		fw.Write(b)
	}
	if r.AllowedUpdates != nil {
		{
			b, _ := json.Marshal(r.AllowedUpdates)
			fw, _ := w.CreateFormField("allowed_updates")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.DropPendingUpdates)
		fw, _ := w.CreateFormField("drop_pending_updates")
		fw.Write(b)
	}
	if r.SecretToken != "" {
		w.WriteField("secret_token", r.SecretToken)
	}
}

// see Bot.DeleteWebhook(ctx, &DeleteWebhookRequest{})
type DeleteWebhookRequest struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"` // Pass True to drop all pending updates
}

func (r *DeleteWebhookRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.DropPendingUpdates)
		fw, _ := w.CreateFormField("drop_pending_updates")
		fw.Write(b)
	}
}

// see Bot.SendMessage(ctx, &SendMessageRequest{})
type SendMessageRequest struct {
	BusinessConnectionID string              `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID              `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int64               `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Text                 string              `json:"text"`                             // Text of the message to be sent, 1-4096 characters after entities parsing
	ParseMode            ParseMode           `json:"parse_mode,omitempty"`             // Mode for parsing entities in the message text. See formatting options for more details.
	Entities             []MessageEntity     `json:"entities,omitempty"`               // A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	LinkPreviewOptions   *LinkPreviewOptions `json:"link_preview_options,omitempty"`   // Link preview generation options for the message
	DisableNotification  bool                `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool                `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool                `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      string              `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters    `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          Markup              `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendMessageRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	w.WriteField("text", r.Text)
	w.WriteField("parse_mode", string(r.ParseMode))
	if r.Entities != nil {
		{
			b, _ := json.Marshal(r.Entities)
			fw, _ := w.CreateFormField("entities")
			fw.Write(b)
		}
	}
	if r.LinkPreviewOptions != nil {
		{
			b, _ := json.Marshal(r.LinkPreviewOptions)
			fw, _ := w.CreateFormField("link_preview_options")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.ForwardMessage(ctx, &ForwardMessageRequest{})
type ForwardMessageRequest struct {
	ChatID              ChatID `json:"chat_id"`                         // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID     int64  `json:"message_thread_id,omitempty"`     // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatID          ChatID `json:"from_chat_id"`                    // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	VideoStartTimestamp int    `json:"video_start_timestamp,omitempty"` // New start timestamp for the forwarded video in the message
	DisableNotification bool   `json:"disable_notification,omitempty"`  // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent      bool   `json:"protect_content,omitempty"`       // Protects the contents of the forwarded message from forwarding and saving
	MessageID           int    `json:"message_id"`                      // Message identifier in the chat specified in from_chat_id
}

func (r *ForwardMessageRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}
	w.WriteField("from_chat_id", r.FromChatID.String())

	{
		b, _ := json.Marshal(r.VideoStartTimestamp)
		fw, _ := w.CreateFormField("video_start_timestamp")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}
}

// see Bot.ForwardMessages(ctx, &ForwardMessagesRequest{})
type ForwardMessagesRequest struct {
	ChatID              ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID     int64  `json:"message_thread_id,omitempty"`    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatID          ChatID `json:"from_chat_id"`                   // Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
	MessageIds          []int  `json:"message_ids"`                    // A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to forward. The identifiers must be specified in a strictly increasing order.
	DisableNotification bool   `json:"disable_notification,omitempty"` // Sends the messages silently. Users will receive a notification with no sound.
	ProtectContent      bool   `json:"protect_content,omitempty"`      // Protects the contents of the forwarded messages from forwarding and saving
}

func (r *ForwardMessagesRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}
	w.WriteField("from_chat_id", r.FromChatID.String())

	{
		b, _ := json.Marshal(r.MessageIds)
		fw, _ := w.CreateFormField("message_ids")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}
}

// see Bot.CopyMessage(ctx, &CopyMessageRequest{})
type CopyMessageRequest struct {
	ChatID                ChatID           `json:"chat_id"`                            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID       int64            `json:"message_thread_id,omitempty"`        // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatID            ChatID           `json:"from_chat_id"`                       // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	MessageID             int              `json:"message_id"`                         // Message identifier in the chat specified in from_chat_id
	VideoStartTimestamp   int              `json:"video_start_timestamp,omitempty"`    // New start timestamp for the copied video in the message
	Caption               string           `json:"caption,omitempty"`                  // New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Mode for parsing entities in the new caption. See formatting options for more details.
	CaptionEntities       []MessageEntity  `json:"caption_entities,omitempty"`         // A JSON-serialized list of special entities that appear in the new caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Pass True, if the caption must be shown above the message media. Ignored if a new caption isn't specified.
	DisableNotification   bool             `json:"disable_notification,omitempty"`     // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             `json:"protect_content,omitempty"`          // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             `json:"allow_paid_broadcast,omitempty"`     // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	ReplyParameters       *ReplyParameters `json:"reply_parameters,omitempty"`         // Description of the message to reply to
	ReplyMarkup           Markup           `json:"reply_markup,omitempty"`             // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *CopyMessageRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}
	w.WriteField("from_chat_id", r.FromChatID.String())

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.VideoStartTimestamp)
		fw, _ := w.CreateFormField("video_start_timestamp")
		fw.Write(b)
	}
	if r.Caption != "" {
		w.WriteField("caption", r.Caption)
	}
	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		{
			b, _ := json.Marshal(r.CaptionEntities)
			fw, _ := w.CreateFormField("caption_entities")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.ShowCaptionAboveMedia)
		fw, _ := w.CreateFormField("show_caption_above_media")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.CopyMessages(ctx, &CopyMessagesRequest{})
type CopyMessagesRequest struct {
	ChatID              ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID     int64  `json:"message_thread_id,omitempty"`    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatID          ChatID `json:"from_chat_id"`                   // Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
	MessageIds          []int  `json:"message_ids"`                    // A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to copy. The identifiers must be specified in a strictly increasing order.
	DisableNotification bool   `json:"disable_notification,omitempty"` // Sends the messages silently. Users will receive a notification with no sound.
	ProtectContent      bool   `json:"protect_content,omitempty"`      // Protects the contents of the sent messages from forwarding and saving
	RemoveCaption       bool   `json:"remove_caption,omitempty"`       // Pass True to copy the messages without their captions
}

func (r *CopyMessagesRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}
	w.WriteField("from_chat_id", r.FromChatID.String())

	{
		b, _ := json.Marshal(r.MessageIds)
		fw, _ := w.CreateFormField("message_ids")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.RemoveCaption)
		fw, _ := w.CreateFormField("remove_caption")
		fw.Write(b)
	}
}

// see Bot.SendPhoto(ctx, &SendPhotoRequest{})
type SendPhotoRequest struct {
	BusinessConnectionID  string           `json:"business_connection_id,omitempty"`   // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID                ChatID           `json:"chat_id"`                            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID       int64            `json:"message_thread_id,omitempty"`        // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Photo                 InputFile        `json:"photo"`                              // Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total. Width and height ratio must be at most 20. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption               string           `json:"caption,omitempty"`                  // Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities       []MessageEntity  `json:"caption_entities,omitempty"`         // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Pass True, if the caption must be shown above the message media
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Pass True if the photo needs to be covered with a spoiler animation
	DisableNotification   bool             `json:"disable_notification,omitempty"`     // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             `json:"protect_content,omitempty"`          // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             `json:"allow_paid_broadcast,omitempty"`     // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID       string           `json:"message_effect_id,omitempty"`        // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters       *ReplyParameters `json:"reply_parameters,omitempty"`         // Description of the message to reply to
	ReplyMarkup           Markup           `json:"reply_markup,omitempty"`             // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendPhotoRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	if r.Photo.FileID != "" {
		w.WriteField("photo", r.Photo.FileID)
	} else if r.Photo.Reader != nil {
		fw, _ := w.CreateFormFile("photo", r.Photo.Reader.Name())
		io.Copy(fw, r.Photo.Reader)
	}
	if r.Caption != "" {
		w.WriteField("caption", r.Caption)
	}
	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		{
			b, _ := json.Marshal(r.CaptionEntities)
			fw, _ := w.CreateFormField("caption_entities")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.ShowCaptionAboveMedia)
		fw, _ := w.CreateFormField("show_caption_above_media")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.HasSpoiler)
		fw, _ := w.CreateFormField("has_spoiler")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SendAudio(ctx, &SendAudioRequest{})
type SendAudioRequest struct {
	BusinessConnectionID string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Audio                InputFile        `json:"audio"`                            // Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption              string           `json:"caption,omitempty"`                // Audio caption, 0-1024 characters after entities parsing
	ParseMode            ParseMode        `json:"parse_mode,omitempty"`             // Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities      []MessageEntity  `json:"caption_entities,omitempty"`       // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration             int              `json:"duration,omitempty"`               // Duration of the audio in seconds
	Performer            string           `json:"performer,omitempty"`              // Performer
	Title                string           `json:"title,omitempty"`                  // Track name
	Thumbnail            InputFile        `json:"thumbnail,omitempty"`              // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          Markup           `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendAudioRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	if r.Audio.FileID != "" {
		w.WriteField("audio", r.Audio.FileID)
	} else if r.Audio.Reader != nil {
		fw, _ := w.CreateFormFile("audio", r.Audio.Reader.Name())
		io.Copy(fw, r.Audio.Reader)
	}
	if r.Caption != "" {
		w.WriteField("caption", r.Caption)
	}
	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		{
			b, _ := json.Marshal(r.CaptionEntities)
			fw, _ := w.CreateFormField("caption_entities")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.Duration)
		fw, _ := w.CreateFormField("duration")
		fw.Write(b)
	}
	if r.Performer != "" {
		w.WriteField("performer", r.Performer)
	}
	if r.Title != "" {
		w.WriteField("title", r.Title)
	}

	if r.Thumbnail.FileID != "" {
		w.WriteField("thumbnail", r.Thumbnail.FileID)
	} else if r.Thumbnail.Reader != nil {
		fw, _ := w.CreateFormFile("thumbnail", r.Thumbnail.Reader.Name())
		io.Copy(fw, r.Thumbnail.Reader)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SendDocument(ctx, &SendDocumentRequest{})
type SendDocumentRequest struct {
	BusinessConnectionID        string           `json:"business_connection_id,omitempty"`         // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID                      ChatID           `json:"chat_id"`                                  // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID             int64            `json:"message_thread_id,omitempty"`              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Document                    InputFile        `json:"document"`                                 // File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail                   InputFile        `json:"thumbnail,omitempty"`                      // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption                     string           `json:"caption,omitempty"`                        // Document caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	ParseMode                   ParseMode        `json:"parse_mode,omitempty"`                     // Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities             []MessageEntity  `json:"caption_entities,omitempty"`               // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"` // Disables automatic server-side content type detection for files uploaded using multipart/form-data
	DisableNotification         bool             `json:"disable_notification,omitempty"`           // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent              bool             `json:"protect_content,omitempty"`                // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast          bool             `json:"allow_paid_broadcast,omitempty"`           // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID             string           `json:"message_effect_id,omitempty"`              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters             *ReplyParameters `json:"reply_parameters,omitempty"`               // Description of the message to reply to
	ReplyMarkup                 Markup           `json:"reply_markup,omitempty"`                   // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendDocumentRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	if r.Document.FileID != "" {
		w.WriteField("document", r.Document.FileID)
	} else if r.Document.Reader != nil {
		fw, _ := w.CreateFormFile("document", r.Document.Reader.Name())
		io.Copy(fw, r.Document.Reader)
	}

	if r.Thumbnail.FileID != "" {
		w.WriteField("thumbnail", r.Thumbnail.FileID)
	} else if r.Thumbnail.Reader != nil {
		fw, _ := w.CreateFormFile("thumbnail", r.Thumbnail.Reader.Name())
		io.Copy(fw, r.Thumbnail.Reader)
	}
	if r.Caption != "" {
		w.WriteField("caption", r.Caption)
	}
	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		{
			b, _ := json.Marshal(r.CaptionEntities)
			fw, _ := w.CreateFormField("caption_entities")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.DisableContentTypeDetection)
		fw, _ := w.CreateFormField("disable_content_type_detection")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SendVideo(ctx, &SendVideoRequest{})
type SendVideoRequest struct {
	BusinessConnectionID  string           `json:"business_connection_id,omitempty"`   // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID                ChatID           `json:"chat_id"`                            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID       int64            `json:"message_thread_id,omitempty"`        // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Video                 InputFile        `json:"video"`                              // Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Duration              int              `json:"duration,omitempty"`                 // Duration of sent video in seconds
	Width                 int              `json:"width,omitempty"`                    // Video width
	Height                int              `json:"height,omitempty"`                   // Video height
	Thumbnail             InputFile        `json:"thumbnail,omitempty"`                // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Cover                 InputFile        `json:"cover,omitempty"`                    // Cover for the video in the message. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	StartTimestamp        int              `json:"start_timestamp,omitempty"`          // Start timestamp for the video in the message
	Caption               string           `json:"caption,omitempty"`                  // Video caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities       []MessageEntity  `json:"caption_entities,omitempty"`         // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Pass True, if the caption must be shown above the message media
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Pass True if the video needs to be covered with a spoiler animation
	SupportsStreaming     bool             `json:"supports_streaming,omitempty"`       // Pass True if the uploaded video is suitable for streaming
	DisableNotification   bool             `json:"disable_notification,omitempty"`     // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             `json:"protect_content,omitempty"`          // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             `json:"allow_paid_broadcast,omitempty"`     // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID       string           `json:"message_effect_id,omitempty"`        // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters       *ReplyParameters `json:"reply_parameters,omitempty"`         // Description of the message to reply to
	ReplyMarkup           Markup           `json:"reply_markup,omitempty"`             // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendVideoRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	if r.Video.FileID != "" {
		w.WriteField("video", r.Video.FileID)
	} else if r.Video.Reader != nil {
		fw, _ := w.CreateFormFile("video", r.Video.Reader.Name())
		io.Copy(fw, r.Video.Reader)
	}

	{
		b, _ := json.Marshal(r.Duration)
		fw, _ := w.CreateFormField("duration")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Width)
		fw, _ := w.CreateFormField("width")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Height)
		fw, _ := w.CreateFormField("height")
		fw.Write(b)
	}

	if r.Thumbnail.FileID != "" {
		w.WriteField("thumbnail", r.Thumbnail.FileID)
	} else if r.Thumbnail.Reader != nil {
		fw, _ := w.CreateFormFile("thumbnail", r.Thumbnail.Reader.Name())
		io.Copy(fw, r.Thumbnail.Reader)
	}

	if r.Cover.FileID != "" {
		w.WriteField("cover", r.Cover.FileID)
	} else if r.Cover.Reader != nil {
		fw, _ := w.CreateFormFile("cover", r.Cover.Reader.Name())
		io.Copy(fw, r.Cover.Reader)
	}

	{
		b, _ := json.Marshal(r.StartTimestamp)
		fw, _ := w.CreateFormField("start_timestamp")
		fw.Write(b)
	}
	if r.Caption != "" {
		w.WriteField("caption", r.Caption)
	}
	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		{
			b, _ := json.Marshal(r.CaptionEntities)
			fw, _ := w.CreateFormField("caption_entities")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.ShowCaptionAboveMedia)
		fw, _ := w.CreateFormField("show_caption_above_media")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.HasSpoiler)
		fw, _ := w.CreateFormField("has_spoiler")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.SupportsStreaming)
		fw, _ := w.CreateFormField("supports_streaming")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SendAnimation(ctx, &SendAnimationRequest{})
type SendAnimationRequest struct {
	BusinessConnectionID  string           `json:"business_connection_id,omitempty"`   // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID                ChatID           `json:"chat_id"`                            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID       int64            `json:"message_thread_id,omitempty"`        // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Animation             InputFile        `json:"animation"`                          // Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Duration              int              `json:"duration,omitempty"`                 // Duration of sent animation in seconds
	Width                 int              `json:"width,omitempty"`                    // Animation width
	Height                int              `json:"height,omitempty"`                   // Animation height
	Thumbnail             InputFile        `json:"thumbnail,omitempty"`                // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption               string           `json:"caption,omitempty"`                  // Animation caption (may also be used when resending animation by file_id), 0-1024 characters after entities parsing
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Mode for parsing entities in the animation caption. See formatting options for more details.
	CaptionEntities       []MessageEntity  `json:"caption_entities,omitempty"`         // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Pass True, if the caption must be shown above the message media
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Pass True if the animation needs to be covered with a spoiler animation
	DisableNotification   bool             `json:"disable_notification,omitempty"`     // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             `json:"protect_content,omitempty"`          // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             `json:"allow_paid_broadcast,omitempty"`     // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID       string           `json:"message_effect_id,omitempty"`        // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters       *ReplyParameters `json:"reply_parameters,omitempty"`         // Description of the message to reply to
	ReplyMarkup           Markup           `json:"reply_markup,omitempty"`             // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendAnimationRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	if r.Animation.FileID != "" {
		w.WriteField("animation", r.Animation.FileID)
	} else if r.Animation.Reader != nil {
		fw, _ := w.CreateFormFile("animation", r.Animation.Reader.Name())
		io.Copy(fw, r.Animation.Reader)
	}

	{
		b, _ := json.Marshal(r.Duration)
		fw, _ := w.CreateFormField("duration")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Width)
		fw, _ := w.CreateFormField("width")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Height)
		fw, _ := w.CreateFormField("height")
		fw.Write(b)
	}

	if r.Thumbnail.FileID != "" {
		w.WriteField("thumbnail", r.Thumbnail.FileID)
	} else if r.Thumbnail.Reader != nil {
		fw, _ := w.CreateFormFile("thumbnail", r.Thumbnail.Reader.Name())
		io.Copy(fw, r.Thumbnail.Reader)
	}
	if r.Caption != "" {
		w.WriteField("caption", r.Caption)
	}
	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		{
			b, _ := json.Marshal(r.CaptionEntities)
			fw, _ := w.CreateFormField("caption_entities")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.ShowCaptionAboveMedia)
		fw, _ := w.CreateFormField("show_caption_above_media")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.HasSpoiler)
		fw, _ := w.CreateFormField("has_spoiler")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SendVoice(ctx, &SendVoiceRequest{})
type SendVoiceRequest struct {
	BusinessConnectionID string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Voice                InputFile        `json:"voice"`                            // Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption              string           `json:"caption,omitempty"`                // Voice message caption, 0-1024 characters after entities parsing
	ParseMode            ParseMode        `json:"parse_mode,omitempty"`             // Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities      []MessageEntity  `json:"caption_entities,omitempty"`       // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration             int              `json:"duration,omitempty"`               // Duration of the voice message in seconds
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          Markup           `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendVoiceRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	if r.Voice.FileID != "" {
		w.WriteField("voice", r.Voice.FileID)
	} else if r.Voice.Reader != nil {
		fw, _ := w.CreateFormFile("voice", r.Voice.Reader.Name())
		io.Copy(fw, r.Voice.Reader)
	}
	if r.Caption != "" {
		w.WriteField("caption", r.Caption)
	}
	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		{
			b, _ := json.Marshal(r.CaptionEntities)
			fw, _ := w.CreateFormField("caption_entities")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.Duration)
		fw, _ := w.CreateFormField("duration")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SendVideoNote(ctx, &SendVideoNoteRequest{})
type SendVideoNoteRequest struct {
	BusinessConnectionID string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	VideoNote            InputFile        `json:"video_note"`                       // Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files. Sending video notes by a URL is currently unsupported
	Duration             int              `json:"duration,omitempty"`               // Duration of sent video in seconds
	Length               int              `json:"length,omitempty"`                 // Video width and height, i.e. diameter of the video message
	Thumbnail            InputFile        `json:"thumbnail,omitempty"`              // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          Markup           `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendVideoNoteRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	if r.VideoNote.FileID != "" {
		w.WriteField("video_note", r.VideoNote.FileID)
	} else if r.VideoNote.Reader != nil {
		fw, _ := w.CreateFormFile("video_note", r.VideoNote.Reader.Name())
		io.Copy(fw, r.VideoNote.Reader)
	}

	{
		b, _ := json.Marshal(r.Duration)
		fw, _ := w.CreateFormField("duration")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Length)
		fw, _ := w.CreateFormField("length")
		fw.Write(b)
	}

	if r.Thumbnail.FileID != "" {
		w.WriteField("thumbnail", r.Thumbnail.FileID)
	} else if r.Thumbnail.Reader != nil {
		fw, _ := w.CreateFormFile("thumbnail", r.Thumbnail.Reader.Name())
		io.Copy(fw, r.Thumbnail.Reader)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SendPaidMedia(ctx, &SendPaidMediaRequest{})
type SendPaidMediaRequest struct {
	BusinessConnectionID  string           `json:"business_connection_id,omitempty"`   // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID                ChatID           `json:"chat_id"`                            // Unique identifier for the target chat or username of the target channel (in the format @channelusername). If the chat is a channel, all Telegram Star proceeds from this media will be credited to the chat's balance. Otherwise, they will be credited to the bot's balance.
	StarCount             int              `json:"star_count"`                         // The number of Telegram Stars that must be paid to buy access to the media; 1-2500
	Media                 []InputPaidMedia `json:"media"`                              // A JSON-serialized array describing the media to be sent; up to 10 items
	Payload               string           `json:"payload,omitempty"`                  // Bot-defined paid media payload, 0-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Caption               string           `json:"caption,omitempty"`                  // Media caption, 0-1024 characters after entities parsing
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`               // Mode for parsing entities in the media caption. See formatting options for more details.
	CaptionEntities       []MessageEntity  `json:"caption_entities,omitempty"`         // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Pass True, if the caption must be shown above the message media
	DisableNotification   bool             `json:"disable_notification,omitempty"`     // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             `json:"protect_content,omitempty"`          // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             `json:"allow_paid_broadcast,omitempty"`     // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	ReplyParameters       *ReplyParameters `json:"reply_parameters,omitempty"`         // Description of the message to reply to
	ReplyMarkup           Markup           `json:"reply_markup,omitempty"`             // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendPaidMediaRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.StarCount)
		fw, _ := w.CreateFormField("star_count")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Media)
		fw, _ := w.CreateFormField("media")
		fw.Write(b)
	}
	if r.Payload != "" {
		w.WriteField("payload", r.Payload)
	}
	if r.Caption != "" {
		w.WriteField("caption", r.Caption)
	}
	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		{
			b, _ := json.Marshal(r.CaptionEntities)
			fw, _ := w.CreateFormField("caption_entities")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.ShowCaptionAboveMedia)
		fw, _ := w.CreateFormField("show_caption_above_media")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SendMediaGroup(ctx, &SendMediaGroupRequest{})
type SendMediaGroupRequest struct {
	BusinessConnectionID string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Media                []InputMedia     `json:"media"`                            // A JSON-serialized array describing messages to be sent, must include 2-10 items
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends messages silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent messages from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
}

func (r *SendMediaGroupRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	for i := 0; i < len(r.Media); i++ {
		inputMedia := r.Media[i]
		fieldName := "media" + strconv.Itoa(i)
		inputFile := inputMedia.getMedia()
		if inputFile.Reader != nil {
			fw, _ := w.CreateFormFile(fieldName, inputFile.Reader.Name())
			io.Copy(fw, inputFile.Reader)
			inputMedia.setMedia("attach://" + fieldName)
		}
	}
	{
		b, _ := json.Marshal(r.Media)
		fw, _ := w.CreateFormField("media")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
}

// see Bot.SendLocation(ctx, &SendLocationRequest{})
type SendLocationRequest struct {
	BusinessConnectionID string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Latitude             float64          `json:"latitude"`                         // Latitude of the location
	Longitude            float64          `json:"longitude"`                        // Longitude of the location
	HorizontalAccuracy   float64          `json:"horizontal_accuracy,omitempty"`    // The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int              `json:"live_period,omitempty"`            // Period in seconds during which the location will be updated (see Live Locations, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	Heading              int              `json:"heading,omitempty"`                // For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int              `json:"proximity_alert_radius,omitempty"` // For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          Markup           `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendLocationRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Latitude)
		fw, _ := w.CreateFormField("latitude")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Longitude)
		fw, _ := w.CreateFormField("longitude")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.HorizontalAccuracy)
		fw, _ := w.CreateFormField("horizontal_accuracy")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.LivePeriod)
		fw, _ := w.CreateFormField("live_period")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Heading)
		fw, _ := w.CreateFormField("heading")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProximityAlertRadius)
		fw, _ := w.CreateFormField("proximity_alert_radius")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SendVenue(ctx, &SendVenueRequest{})
type SendVenueRequest struct {
	BusinessConnectionID string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Latitude             float64          `json:"latitude"`                         // Latitude of the venue
	Longitude            float64          `json:"longitude"`                        // Longitude of the venue
	Title                string           `json:"title"`                            // Name of the venue
	Address              string           `json:"address"`                          // Address of the venue
	FoursquareID         string           `json:"foursquare_id,omitempty"`          // Foursquare identifier of the venue
	FoursquareType       string           `json:"foursquare_type,omitempty"`        // Foursquare type of the venue, if known. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	GooglePlaceID        string           `json:"google_place_id,omitempty"`        // Google Places identifier of the venue
	GooglePlaceType      string           `json:"google_place_type,omitempty"`      // Google Places type of the venue. (See supported types.)
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          Markup           `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendVenueRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Latitude)
		fw, _ := w.CreateFormField("latitude")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Longitude)
		fw, _ := w.CreateFormField("longitude")
		fw.Write(b)
	}

	w.WriteField("title", r.Title)

	w.WriteField("address", r.Address)
	if r.FoursquareID != "" {
		w.WriteField("foursquare_id", r.FoursquareID)
	}
	if r.FoursquareType != "" {
		w.WriteField("foursquare_type", r.FoursquareType)
	}
	if r.GooglePlaceID != "" {
		w.WriteField("google_place_id", r.GooglePlaceID)
	}
	if r.GooglePlaceType != "" {
		w.WriteField("google_place_type", r.GooglePlaceType)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SendContact(ctx, &SendContactRequest{})
type SendContactRequest struct {
	BusinessConnectionID string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	PhoneNumber          string           `json:"phone_number"`                     // Contact's phone number
	FirstName            string           `json:"first_name"`                       // Contact's first name
	LastName             string           `json:"last_name,omitempty"`              // Contact's last name
	Vcard                string           `json:"vcard,omitempty"`                  // Additional data about the contact in the form of a vCard, 0-2048 bytes
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          Markup           `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendContactRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	w.WriteField("phone_number", r.PhoneNumber)

	w.WriteField("first_name", r.FirstName)
	if r.LastName != "" {
		w.WriteField("last_name", r.LastName)
	}
	if r.Vcard != "" {
		w.WriteField("vcard", r.Vcard)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SendPoll(ctx, &SendPollRequest{})
type SendPollRequest struct {
	BusinessConnectionID  string            `json:"business_connection_id,omitempty"`  // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID                ChatID            `json:"chat_id"`                           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID       int64             `json:"message_thread_id,omitempty"`       // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Question              string            `json:"question"`                          // Poll question, 1-300 characters
	QuestionParseMode     string            `json:"question_parse_mode,omitempty"`     // Mode for parsing entities in the question. See formatting options for more details. Currently, only custom emoji entities are allowed
	QuestionEntities      []MessageEntity   `json:"question_entities,omitempty"`       // A JSON-serialized list of special entities that appear in the poll question. It can be specified instead of question_parse_mode
	Options               []InputPollOption `json:"options"`                           // A JSON-serialized list of 2-10 answer options
	IsAnonymous           bool              `json:"is_anonymous,omitempty"`            // True, if the poll needs to be anonymous, defaults to True
	Type                  string            `json:"type,omitempty"`                    // Poll type, "quiz" or "regular", defaults to "regular"
	AllowsMultipleAnswers bool              `json:"allows_multiple_answers,omitempty"` // True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	CorrectOptionID       int64             `json:"correct_option_id,omitempty"`       // 0-based identifier of the correct answer option, required for polls in quiz mode
	Explanation           string            `json:"explanation,omitempty"`             // Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	ExplanationParseMode  string            `json:"explanation_parse_mode,omitempty"`  // Mode for parsing entities in the explanation. See formatting options for more details.
	ExplanationEntities   []MessageEntity   `json:"explanation_entities,omitempty"`    // A JSON-serialized list of special entities that appear in the poll explanation. It can be specified instead of explanation_parse_mode
	OpenPeriod            int               `json:"open_period,omitempty"`             // Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with close_date.
	CloseDate             int               `json:"close_date,omitempty"`              // Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with open_period.
	IsClosed              bool              `json:"is_closed,omitempty"`               // Pass True if the poll needs to be immediately closed. This can be useful for poll preview.
	DisableNotification   bool              `json:"disable_notification,omitempty"`    // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool              `json:"protect_content,omitempty"`         // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool              `json:"allow_paid_broadcast,omitempty"`    // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID       string            `json:"message_effect_id,omitempty"`       // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters       *ReplyParameters  `json:"reply_parameters,omitempty"`        // Description of the message to reply to
	ReplyMarkup           Markup            `json:"reply_markup,omitempty"`            // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendPollRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	w.WriteField("question", r.Question)
	if r.QuestionParseMode != "" {
		w.WriteField("question_parse_mode", r.QuestionParseMode)
	}
	if r.QuestionEntities != nil {
		{
			b, _ := json.Marshal(r.QuestionEntities)
			fw, _ := w.CreateFormField("question_entities")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.Options)
		fw, _ := w.CreateFormField("options")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.IsAnonymous)
		fw, _ := w.CreateFormField("is_anonymous")
		fw.Write(b)
	}
	if r.Type != "" {
		w.WriteField("type", r.Type)
	}

	{
		b, _ := json.Marshal(r.AllowsMultipleAnswers)
		fw, _ := w.CreateFormField("allows_multiple_answers")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CorrectOptionID)
		fw, _ := w.CreateFormField("correct_option_id")
		fw.Write(b)
	}
	if r.Explanation != "" {
		w.WriteField("explanation", r.Explanation)
	}
	if r.ExplanationParseMode != "" {
		w.WriteField("explanation_parse_mode", r.ExplanationParseMode)
	}
	if r.ExplanationEntities != nil {
		{
			b, _ := json.Marshal(r.ExplanationEntities)
			fw, _ := w.CreateFormField("explanation_entities")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.OpenPeriod)
		fw, _ := w.CreateFormField("open_period")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CloseDate)
		fw, _ := w.CreateFormField("close_date")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.IsClosed)
		fw, _ := w.CreateFormField("is_closed")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SendDice(ctx, &SendDiceRequest{})
type SendDiceRequest struct {
	BusinessConnectionID string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Emoji                string           `json:"emoji,omitempty"`                  // Emoji on which the dice throw animation is based. Currently, must be one of "🎲", "🎯", "🏀", "⚽", "🎳", or "🎰". Dice can have values 1-6 for "🎲", "🎯" and "🎳", values 1-5 for "🏀" and "⚽", and values 1-64 for "🎰". Defaults to "🎲"
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          Markup           `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendDiceRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}
	if r.Emoji != "" {
		w.WriteField("emoji", r.Emoji)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SendChatAction(ctx, &SendChatActionRequest{})
type SendChatActionRequest struct {
	BusinessConnectionID string     `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the action will be sent
	ChatID               ChatID     `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int64      `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread; for supergroups only
	Action               ChatAction `json:"action"`                           // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers, find_location for location data, record_video_note or upload_video_note for video notes.
}

func (r *SendChatActionRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}
	w.WriteField("action", string(r.Action))
}

// see Bot.SetMessageReaction(ctx, &SetMessageReactionRequest{})
type SetMessageReactionRequest struct {
	ChatID    ChatID         `json:"chat_id"`            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID int            `json:"message_id"`         // Identifier of the target message. If the message belongs to a media group, the reaction is set to the first non-deleted message in the group instead.
	Reaction  []ReactionType `json:"reaction,omitempty"` // A JSON-serialized list of reaction types to set on the message. Currently, as non-premium users, bots can set up to one reaction per message. A custom emoji reaction can be used if it is either already present on the message or explicitly allowed by chat administrators. Paid reactions can't be used by bots.
	IsBig     bool           `json:"is_big,omitempty"`   // Pass True to set the reaction with a big animation
}

func (r *SetMessageReactionRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}
	if r.Reaction != nil {
		{
			b, _ := json.Marshal(r.Reaction)
			fw, _ := w.CreateFormField("reaction")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.IsBig)
		fw, _ := w.CreateFormField("is_big")
		fw.Write(b)
	}
}

// see Bot.GetUserProfilePhotos(ctx, &GetUserProfilePhotosRequest{})
type GetUserProfilePhotosRequest struct {
	UserID int64 `json:"user_id"`          // Unique identifier of the target user
	Offset int64 `json:"offset,omitempty"` // Sequential number of the first photo to be returned. By default, all photos are returned.
	Limit  int   `json:"limit,omitempty"`  // Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.
}

func (r *GetUserProfilePhotosRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Offset)
		fw, _ := w.CreateFormField("offset")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Limit)
		fw, _ := w.CreateFormField("limit")
		fw.Write(b)
	}
}

// see Bot.SetUserEmojiStatus(ctx, &SetUserEmojiStatusRequest{})
type SetUserEmojiStatusRequest struct {
	UserID                    int64  `json:"user_id"`                                // Unique identifier of the target user
	EmojiStatusCustomEmojiID  string `json:"emoji_status_custom_emoji_id,omitempty"` // Custom emoji identifier of the emoji status to set. Pass an empty string to remove the status.
	EmojiStatusExpirationDate int    `json:"emoji_status_expiration_date,omitempty"` // Expiration date of the emoji status, if any
}

func (r *SetUserEmojiStatusRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}
	if r.EmojiStatusCustomEmojiID != "" {
		w.WriteField("emoji_status_custom_emoji_id", r.EmojiStatusCustomEmojiID)
	}

	{
		b, _ := json.Marshal(r.EmojiStatusExpirationDate)
		fw, _ := w.CreateFormField("emoji_status_expiration_date")
		fw.Write(b)
	}
}

// see Bot.GetFile(ctx, &GetFileRequest{})
type GetFileRequest struct {
	FileID string `json:"file_id"` // File identifier to get information about
}

func (r *GetFileRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("file_id", r.FileID)
}

// see Bot.BanChatMember(ctx, &BanChatMemberRequest{})
type BanChatMemberRequest struct {
	ChatID         ChatID `json:"chat_id"`                   // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserID         int64  `json:"user_id"`                   // Unique identifier of the target user
	UntilDate      int    `json:"until_date,omitempty"`      // Date when the user will be unbanned; Unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever. Applied for supergroups and channels only.
	RevokeMessages bool   `json:"revoke_messages,omitempty"` // Pass True to delete all messages from the chat for the user that is being removed. If False, the user will be able to see messages in the group that were sent before the user was removed. Always True for supergroups and channels.
}

func (r *BanChatMemberRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.UntilDate)
		fw, _ := w.CreateFormField("until_date")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.RevokeMessages)
		fw, _ := w.CreateFormField("revoke_messages")
		fw.Write(b)
	}
}

// see Bot.UnbanChatMember(ctx, &UnbanChatMemberRequest{})
type UnbanChatMemberRequest struct {
	ChatID       ChatID `json:"chat_id"`                  // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserID       int64  `json:"user_id"`                  // Unique identifier of the target user
	OnlyIfBanned bool   `json:"only_if_banned,omitempty"` // Do nothing if the user is not banned
}

func (r *UnbanChatMemberRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.OnlyIfBanned)
		fw, _ := w.CreateFormField("only_if_banned")
		fw.Write(b)
	}
}

// see Bot.RestrictChatMember(ctx, &RestrictChatMemberRequest{})
type RestrictChatMemberRequest struct {
	ChatID                        ChatID           `json:"chat_id"`                                    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserID                        int64            `json:"user_id"`                                    // Unique identifier of the target user
	Permissions                   *ChatPermissions `json:"permissions"`                                // A JSON-serialized object for new user permissions
	UseIndependentChatPermissions bool             `json:"use_independent_chat_permissions,omitempty"` // Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
	UntilDate                     int              `json:"until_date,omitempty"`                       // Date when restrictions will be lifted for the user; Unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
}

func (r *RestrictChatMemberRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Permissions)
		fw, _ := w.CreateFormField("permissions")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.UseIndependentChatPermissions)
		fw, _ := w.CreateFormField("use_independent_chat_permissions")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.UntilDate)
		fw, _ := w.CreateFormField("until_date")
		fw.Write(b)
	}
}

// see Bot.PromoteChatMember(ctx, &PromoteChatMemberRequest{})
type PromoteChatMemberRequest struct {
	ChatID              ChatID `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserID              int64  `json:"user_id"`                          // Unique identifier of the target user
	IsAnonymous         bool   `json:"is_anonymous,omitempty"`           // Pass True if the administrator's presence in the chat is hidden
	CanManageChat       bool   `json:"can_manage_chat,omitempty"`        // Pass True if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege.
	CanDeleteMessages   bool   `json:"can_delete_messages,omitempty"`    // Pass True if the administrator can delete messages of other users
	CanManageVideoChats bool   `json:"can_manage_video_chats,omitempty"` // Pass True if the administrator can manage video chats
	CanRestrictMembers  bool   `json:"can_restrict_members,omitempty"`   // Pass True if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanPromoteMembers   bool   `json:"can_promote_members,omitempty"`    // Pass True if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him)
	CanChangeInfo       bool   `json:"can_change_info,omitempty"`        // Pass True if the administrator can change chat title, photo and other settings
	CanInviteUsers      bool   `json:"can_invite_users,omitempty"`       // Pass True if the administrator can invite new users to the chat
	CanPostStories      bool   `json:"can_post_stories,omitempty"`       // Pass True if the administrator can post stories to the chat
	CanEditStories      bool   `json:"can_edit_stories,omitempty"`       // Pass True if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanDeleteStories    bool   `json:"can_delete_stories,omitempty"`     // Pass True if the administrator can delete stories posted by other users
	CanPostMessages     bool   `json:"can_post_messages,omitempty"`      // Pass True if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanEditMessages     bool   `json:"can_edit_messages,omitempty"`      // Pass True if the administrator can edit messages of other users and can pin messages; for channels only
	CanPinMessages      bool   `json:"can_pin_messages,omitempty"`       // Pass True if the administrator can pin messages; for supergroups only
	CanManageTopics     bool   `json:"can_manage_topics,omitempty"`      // Pass True if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
}

func (r *PromoteChatMemberRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.IsAnonymous)
		fw, _ := w.CreateFormField("is_anonymous")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanManageChat)
		fw, _ := w.CreateFormField("can_manage_chat")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanDeleteMessages)
		fw, _ := w.CreateFormField("can_delete_messages")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanManageVideoChats)
		fw, _ := w.CreateFormField("can_manage_video_chats")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanRestrictMembers)
		fw, _ := w.CreateFormField("can_restrict_members")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanPromoteMembers)
		fw, _ := w.CreateFormField("can_promote_members")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanChangeInfo)
		fw, _ := w.CreateFormField("can_change_info")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanInviteUsers)
		fw, _ := w.CreateFormField("can_invite_users")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanPostStories)
		fw, _ := w.CreateFormField("can_post_stories")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanEditStories)
		fw, _ := w.CreateFormField("can_edit_stories")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanDeleteStories)
		fw, _ := w.CreateFormField("can_delete_stories")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanPostMessages)
		fw, _ := w.CreateFormField("can_post_messages")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanEditMessages)
		fw, _ := w.CreateFormField("can_edit_messages")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanPinMessages)
		fw, _ := w.CreateFormField("can_pin_messages")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CanManageTopics)
		fw, _ := w.CreateFormField("can_manage_topics")
		fw.Write(b)
	}
}

// see Bot.SetChatAdministratorCustomTitle(ctx, &SetChatAdministratorCustomTitleRequest{})
type SetChatAdministratorCustomTitleRequest struct {
	ChatID      ChatID `json:"chat_id"`      // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserID      int64  `json:"user_id"`      // Unique identifier of the target user
	CustomTitle string `json:"custom_title"` // New custom title for the administrator; 0-16 characters, emoji are not allowed
}

func (r *SetChatAdministratorCustomTitleRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	w.WriteField("custom_title", r.CustomTitle)
}

// see Bot.BanChatSenderChat(ctx, &BanChatSenderChatRequest{})
type BanChatSenderChatRequest struct {
	ChatID       ChatID `json:"chat_id"`        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	SenderChatID int64  `json:"sender_chat_id"` // Unique identifier of the target sender chat
}

func (r *BanChatSenderChatRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.SenderChatID)
		fw, _ := w.CreateFormField("sender_chat_id")
		fw.Write(b)
	}
}

// see Bot.UnbanChatSenderChat(ctx, &UnbanChatSenderChatRequest{})
type UnbanChatSenderChatRequest struct {
	ChatID       ChatID `json:"chat_id"`        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	SenderChatID int64  `json:"sender_chat_id"` // Unique identifier of the target sender chat
}

func (r *UnbanChatSenderChatRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.SenderChatID)
		fw, _ := w.CreateFormField("sender_chat_id")
		fw.Write(b)
	}
}

// see Bot.SetChatPermissions(ctx, &SetChatPermissionsRequest{})
type SetChatPermissionsRequest struct {
	ChatID                        ChatID           `json:"chat_id"`                                    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Permissions                   *ChatPermissions `json:"permissions"`                                // A JSON-serialized object for new default chat permissions
	UseIndependentChatPermissions bool             `json:"use_independent_chat_permissions,omitempty"` // Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
}

func (r *SetChatPermissionsRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.Permissions)
		fw, _ := w.CreateFormField("permissions")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.UseIndependentChatPermissions)
		fw, _ := w.CreateFormField("use_independent_chat_permissions")
		fw.Write(b)
	}
}

// see Bot.ExportChatInviteLink(ctx, &ExportChatInviteLinkRequest{})
type ExportChatInviteLinkRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

func (r *ExportChatInviteLinkRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.CreateChatInviteLink(ctx, &CreateChatInviteLinkRequest{})
type CreateChatInviteLinkRequest struct {
	ChatID             ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Name               string `json:"name,omitempty"`                 // Invite link name; 0-32 characters
	ExpireDate         int    `json:"expire_date,omitempty"`          // Point in time (Unix timestamp) when the link will expire
	MemberLimit        int    `json:"member_limit,omitempty"`         // The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"` // True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
}

func (r *CreateChatInviteLinkRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	if r.Name != "" {
		w.WriteField("name", r.Name)
	}

	{
		b, _ := json.Marshal(r.ExpireDate)
		fw, _ := w.CreateFormField("expire_date")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.MemberLimit)
		fw, _ := w.CreateFormField("member_limit")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CreatesJoinRequest)
		fw, _ := w.CreateFormField("creates_join_request")
		fw.Write(b)
	}
}

// see Bot.EditChatInviteLink(ctx, &EditChatInviteLinkRequest{})
type EditChatInviteLinkRequest struct {
	ChatID             ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	InviteLink         string `json:"invite_link"`                    // The invite link to edit
	Name               string `json:"name,omitempty"`                 // Invite link name; 0-32 characters
	ExpireDate         int    `json:"expire_date,omitempty"`          // Point in time (Unix timestamp) when the link will expire
	MemberLimit        int    `json:"member_limit,omitempty"`         // The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"` // True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
}

func (r *EditChatInviteLinkRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("invite_link", r.InviteLink)
	if r.Name != "" {
		w.WriteField("name", r.Name)
	}

	{
		b, _ := json.Marshal(r.ExpireDate)
		fw, _ := w.CreateFormField("expire_date")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.MemberLimit)
		fw, _ := w.CreateFormField("member_limit")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CreatesJoinRequest)
		fw, _ := w.CreateFormField("creates_join_request")
		fw.Write(b)
	}
}

// see Bot.CreateChatSubscriptionInviteLink(ctx, &CreateChatSubscriptionInviteLinkRequest{})
type CreateChatSubscriptionInviteLinkRequest struct {
	ChatID             ChatID `json:"chat_id"`             // Unique identifier for the target channel chat or username of the target channel (in the format @channelusername)
	Name               string `json:"name,omitempty"`      // Invite link name; 0-32 characters
	SubscriptionPeriod int    `json:"subscription_period"` // The number of seconds the subscription will be active for before the next payment. Currently, it must always be 2592000 (30 days).
	SubscriptionPrice  int    `json:"subscription_price"`  // The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat; 1-2500
}

func (r *CreateChatSubscriptionInviteLinkRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	if r.Name != "" {
		w.WriteField("name", r.Name)
	}

	{
		b, _ := json.Marshal(r.SubscriptionPeriod)
		fw, _ := w.CreateFormField("subscription_period")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.SubscriptionPrice)
		fw, _ := w.CreateFormField("subscription_price")
		fw.Write(b)
	}
}

// see Bot.EditChatSubscriptionInviteLink(ctx, &EditChatSubscriptionInviteLinkRequest{})
type EditChatSubscriptionInviteLinkRequest struct {
	ChatID     ChatID `json:"chat_id"`        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	InviteLink string `json:"invite_link"`    // The invite link to edit
	Name       string `json:"name,omitempty"` // Invite link name; 0-32 characters
}

func (r *EditChatSubscriptionInviteLinkRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("invite_link", r.InviteLink)
	if r.Name != "" {
		w.WriteField("name", r.Name)
	}
}

// see Bot.RevokeChatInviteLink(ctx, &RevokeChatInviteLinkRequest{})
type RevokeChatInviteLinkRequest struct {
	ChatID     ChatID `json:"chat_id"`     // Unique identifier of the target chat or username of the target channel (in the format @channelusername)
	InviteLink string `json:"invite_link"` // The invite link to revoke
}

func (r *RevokeChatInviteLinkRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("invite_link", r.InviteLink)
}

// see Bot.ApproveChatJoinRequest(ctx, &ApproveChatJoinRequest{})
type ApproveChatJoinRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserID int64  `json:"user_id"` // Unique identifier of the target user
}

func (r *ApproveChatJoinRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}
}

// see Bot.DeclineChatJoinRequest(ctx, &DeclineChatJoinRequest{})
type DeclineChatJoinRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserID int64  `json:"user_id"` // Unique identifier of the target user
}

func (r *DeclineChatJoinRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}
}

// see Bot.SetChatPhoto(ctx, &SetChatPhotoRequest{})
type SetChatPhotoRequest struct {
	ChatID ChatID    `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo  InputFile `json:"photo"`   // New chat photo, uploaded using multipart/form-data
}

func (r *SetChatPhotoRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	if r.Photo.FileID != "" {
		w.WriteField("photo", r.Photo.FileID)
	} else if r.Photo.Reader != nil {
		fw, _ := w.CreateFormFile("photo", r.Photo.Reader.Name())
		io.Copy(fw, r.Photo.Reader)
	}
}

// see Bot.DeleteChatPhoto(ctx, &DeleteChatPhotoRequest{})
type DeleteChatPhotoRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

func (r *DeleteChatPhotoRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.SetChatTitle(ctx, &SetChatTitleRequest{})
type SetChatTitleRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Title  string `json:"title"`   // New chat title, 1-128 characters
}

func (r *SetChatTitleRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("title", r.Title)
}

// see Bot.SetChatDescription(ctx, &SetChatDescriptionRequest{})
type SetChatDescriptionRequest struct {
	ChatID      ChatID `json:"chat_id"`               // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Description string `json:"description,omitempty"` // New chat description, 0-255 characters
}

func (r *SetChatDescriptionRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	if r.Description != "" {
		w.WriteField("description", r.Description)
	}
}

// see Bot.PinChatMessage(ctx, &PinChatMessageRequest{})
type PinChatMessageRequest struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be pinned
	ChatID               ChatID `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int    `json:"message_id"`                       // Identifier of a message to pin
	DisableNotification  bool   `json:"disable_notification,omitempty"`   // Pass True if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels and private chats.
}

func (r *PinChatMessageRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}
}

// see Bot.UnpinChatMessage(ctx, &UnpinChatMessageRequest{})
type UnpinChatMessageRequest struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be unpinned
	ChatID               ChatID `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int    `json:"message_id,omitempty"`             // Identifier of the message to unpin. Required if business_connection_id is specified. If not specified, the most recent pinned message (by sending date) will be unpinned.
}

func (r *UnpinChatMessageRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}
}

// see Bot.UnpinAllChatMessages(ctx, &UnpinAllChatMessagesRequest{})
type UnpinAllChatMessagesRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

func (r *UnpinAllChatMessagesRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.LeaveChat(ctx, &LeaveChatRequest{})
type LeaveChatRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (r *LeaveChatRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.GetChat(ctx, &GetChatRequest{})
type GetChatRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (r *GetChatRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.GetChatAdministrators(ctx, &GetChatAdministratorsRequest{})
type GetChatAdministratorsRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (r *GetChatAdministratorsRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.GetChatMemberCount(ctx, &GetChatMemberCountRequest{})
type GetChatMemberCountRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (r *GetChatMemberCountRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.GetChatMember(ctx, &GetChatMemberRequest{})
type GetChatMemberRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	UserID int64  `json:"user_id"` // Unique identifier of the target user
}

func (r *GetChatMemberRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}
}

// see Bot.SetChatStickerSet(ctx, &SetChatStickerSetRequest{})
type SetChatStickerSetRequest struct {
	ChatID         ChatID `json:"chat_id"`          // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	StickerSetName string `json:"sticker_set_name"` // Name of the sticker set to be set as the group sticker set
}

func (r *SetChatStickerSetRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("sticker_set_name", r.StickerSetName)
}

// see Bot.DeleteChatStickerSet(ctx, &DeleteChatStickerSetRequest{})
type DeleteChatStickerSetRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (r *DeleteChatStickerSetRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.CreateForumTopic(ctx, &CreateForumTopicRequest{})
type CreateForumTopicRequest struct {
	ChatID            ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Name              string `json:"name"`                           // Topic name, 1-128 characters
	IconColor         int    `json:"icon_color,omitempty"`           // Color of the topic icon in RGB format. Currently, must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F)
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"` // Unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers.
}

func (r *CreateForumTopicRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("name", r.Name)

	{
		b, _ := json.Marshal(r.IconColor)
		fw, _ := w.CreateFormField("icon_color")
		fw.Write(b)
	}
	if r.IconCustomEmojiID != "" {
		w.WriteField("icon_custom_emoji_id", r.IconCustomEmojiID)
	}
}

// see Bot.EditForumTopic(ctx, &EditForumTopicRequest{})
type EditForumTopicRequest struct {
	ChatID            ChatID `json:"chat_id"`                        // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadID   int64  `json:"message_thread_id"`              // Unique identifier for the target message thread of the forum topic
	Name              string `json:"name,omitempty"`                 // New topic name, 0-128 characters. If not specified or empty, the current name of the topic will be kept
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"` // New unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers. Pass an empty string to remove the icon. If not specified, the current icon will be kept
}

func (r *EditForumTopicRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}
	if r.Name != "" {
		w.WriteField("name", r.Name)
	}
	if r.IconCustomEmojiID != "" {
		w.WriteField("icon_custom_emoji_id", r.IconCustomEmojiID)
	}
}

// see Bot.CloseForumTopic(ctx, &CloseForumTopicRequest{})
type CloseForumTopicRequest struct {
	ChatID          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadID int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

func (r *CloseForumTopicRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}
}

// see Bot.ReopenForumTopic(ctx, &ReopenForumTopicRequest{})
type ReopenForumTopicRequest struct {
	ChatID          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadID int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

func (r *ReopenForumTopicRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}
}

// see Bot.DeleteForumTopic(ctx, &DeleteForumTopicRequest{})
type DeleteForumTopicRequest struct {
	ChatID          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadID int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

func (r *DeleteForumTopicRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}
}

// see Bot.UnpinAllForumTopicMessages(ctx, &UnpinAllForumTopicMessagesRequest{})
type UnpinAllForumTopicMessagesRequest struct {
	ChatID          ChatID `json:"chat_id"`           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadID int64  `json:"message_thread_id"` // Unique identifier for the target message thread of the forum topic
}

func (r *UnpinAllForumTopicMessagesRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}
}

// see Bot.EditGeneralForumTopic(ctx, &EditGeneralForumTopicRequest{})
type EditGeneralForumTopicRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Name   string `json:"name"`    // New topic name, 1-128 characters
}

func (r *EditGeneralForumTopicRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("name", r.Name)
}

// see Bot.CloseGeneralForumTopic(ctx, &CloseGeneralForumTopicRequest{})
type CloseGeneralForumTopicRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (r *CloseGeneralForumTopicRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.ReopenGeneralForumTopic(ctx, &ReopenGeneralForumTopicRequest{})
type ReopenGeneralForumTopicRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (r *ReopenGeneralForumTopicRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.HideGeneralForumTopic(ctx, &HideGeneralForumTopicRequest{})
type HideGeneralForumTopicRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (r *HideGeneralForumTopicRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.UnhideGeneralForumTopic(ctx, &UnhideGeneralForumTopicRequest{})
type UnhideGeneralForumTopicRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (r *UnhideGeneralForumTopicRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.UnpinAllGeneralForumTopicMessages(ctx, &UnpinAllGeneralForumTopicMessagesRequest{})
type UnpinAllGeneralForumTopicMessagesRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (r *UnpinAllGeneralForumTopicMessagesRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.AnswerCallbackQuery(ctx, &AnswerCallbackQueryRequest{})
type AnswerCallbackQueryRequest struct {
	CallbackQueryID string `json:"callback_query_id"`    // Unique identifier for the query to be answered
	Text            string `json:"text,omitempty"`       // Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	ShowAlert       bool   `json:"show_alert,omitempty"` // If True, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	Url             string `json:"url,omitempty"`        // URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @BotFather, specify the URL that opens your game - note that this will only work if the query comes from a callback_game button. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	CacheTime       int    `json:"cache_time,omitempty"` // The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
}

func (r *AnswerCallbackQueryRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("callback_query_id", r.CallbackQueryID)
	if r.Text != "" {
		w.WriteField("text", r.Text)
	}

	{
		b, _ := json.Marshal(r.ShowAlert)
		fw, _ := w.CreateFormField("show_alert")
		fw.Write(b)
	}
	if r.Url != "" {
		w.WriteField("url", r.Url)
	}

	{
		b, _ := json.Marshal(r.CacheTime)
		fw, _ := w.CreateFormField("cache_time")
		fw.Write(b)
	}
}

// see Bot.GetUserChatBoosts(ctx, &GetUserChatBoostsRequest{})
type GetUserChatBoostsRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the chat or username of the channel (in the format @channelusername)
	UserID int64  `json:"user_id"` // Unique identifier of the target user
}

func (r *GetUserChatBoostsRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}
}

// see Bot.GetBusinessConnection(ctx, &GetBusinessConnectionRequest{})
type GetBusinessConnectionRequest struct {
	BusinessConnectionID string `json:"business_connection_id"` // Unique identifier of the business connection
}

func (r *GetBusinessConnectionRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)
}

// see Bot.SetMyCommands(ctx, &SetMyCommandsRequest{})
type SetMyCommandsRequest struct {
	Commands     []BotCommand     `json:"commands"`                // A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	Scope        *BotCommandScope `json:"scope,omitempty"`         // A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string           `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

func (r *SetMyCommandsRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.Commands)
		fw, _ := w.CreateFormField("commands")
		fw.Write(b)
	}
	if r.Scope != nil {
		{
			b, _ := json.Marshal(r.Scope)
			fw, _ := w.CreateFormField("scope")
			fw.Write(b)
		}
	}
	if r.LanguageCode != "" {
		w.WriteField("language_code", r.LanguageCode)
	}
}

// see Bot.DeleteMyCommands(ctx, &DeleteMyCommandsRequest{})
type DeleteMyCommandsRequest struct {
	Scope        *BotCommandScope `json:"scope,omitempty"`         // A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string           `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

func (r *DeleteMyCommandsRequest) writeMultipart(w *multipart.Writer) {
	if r.Scope != nil {
		{
			b, _ := json.Marshal(r.Scope)
			fw, _ := w.CreateFormField("scope")
			fw.Write(b)
		}
	}
	if r.LanguageCode != "" {
		w.WriteField("language_code", r.LanguageCode)
	}
}

// see Bot.GetMyCommands(ctx, &GetMyCommandsRequest{})
type GetMyCommandsRequest struct {
	Scope        *BotCommandScope `json:"scope,omitempty"`         // A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
	LanguageCode string           `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code or an empty string
}

func (r *GetMyCommandsRequest) writeMultipart(w *multipart.Writer) {
	if r.Scope != nil {
		{
			b, _ := json.Marshal(r.Scope)
			fw, _ := w.CreateFormField("scope")
			fw.Write(b)
		}
	}
	if r.LanguageCode != "" {
		w.WriteField("language_code", r.LanguageCode)
	}
}

// see Bot.SetMyName(ctx, &SetMyNameRequest{})
type SetMyNameRequest struct {
	Name         string `json:"name,omitempty"`          // New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given language.
	LanguageCode string `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code. If empty, the name will be shown to all users for whose language there is no dedicated name.
}

func (r *SetMyNameRequest) writeMultipart(w *multipart.Writer) {
	if r.Name != "" {
		w.WriteField("name", r.Name)
	}
	if r.LanguageCode != "" {
		w.WriteField("language_code", r.LanguageCode)
	}
}

// see Bot.GetMyName(ctx, &GetMyNameRequest{})
type GetMyNameRequest struct {
	LanguageCode string `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code or an empty string
}

func (r *GetMyNameRequest) writeMultipart(w *multipart.Writer) {
	if r.LanguageCode != "" {
		w.WriteField("language_code", r.LanguageCode)
	}
}

// see Bot.SetMyDescription(ctx, &SetMyDescriptionRequest{})
type SetMyDescriptionRequest struct {
	Description  string `json:"description,omitempty"`   // New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language.
	LanguageCode string `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code. If empty, the description will be applied to all users for whose language there is no dedicated description.
}

func (r *SetMyDescriptionRequest) writeMultipart(w *multipart.Writer) {
	if r.Description != "" {
		w.WriteField("description", r.Description)
	}
	if r.LanguageCode != "" {
		w.WriteField("language_code", r.LanguageCode)
	}
}

// see Bot.GetMyDescription(ctx, &GetMyDescriptionRequest{})
type GetMyDescriptionRequest struct {
	LanguageCode string `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code or an empty string
}

func (r *GetMyDescriptionRequest) writeMultipart(w *multipart.Writer) {
	if r.LanguageCode != "" {
		w.WriteField("language_code", r.LanguageCode)
	}
}

// see Bot.SetMyShortDescription(ctx, &SetMyShortDescriptionRequest{})
type SetMyShortDescriptionRequest struct {
	ShortDescription string `json:"short_description,omitempty"` // New short description for the bot; 0-120 characters. Pass an empty string to remove the dedicated short description for the given language.
	LanguageCode     string `json:"language_code,omitempty"`     // A two-letter ISO 639-1 language code. If empty, the short description will be applied to all users for whose language there is no dedicated short description.
}

func (r *SetMyShortDescriptionRequest) writeMultipart(w *multipart.Writer) {
	if r.ShortDescription != "" {
		w.WriteField("short_description", r.ShortDescription)
	}
	if r.LanguageCode != "" {
		w.WriteField("language_code", r.LanguageCode)
	}
}

// see Bot.GetMyShortDescription(ctx, &GetMyShortDescriptionRequest{})
type GetMyShortDescriptionRequest struct {
	LanguageCode string `json:"language_code,omitempty"` // A two-letter ISO 639-1 language code or an empty string
}

func (r *GetMyShortDescriptionRequest) writeMultipart(w *multipart.Writer) {
	if r.LanguageCode != "" {
		w.WriteField("language_code", r.LanguageCode)
	}
}

// see Bot.SetChatMenuButton(ctx, &SetChatMenuButtonRequest{})
type SetChatMenuButtonRequest struct {
	ChatID     int64       `json:"chat_id,omitempty"`     // Unique identifier for the target private chat. If not specified, default bot's menu button will be changed
	MenuButton *MenuButton `json:"menu_button,omitempty"` // A JSON-serialized object for the bot's new menu button. Defaults to MenuButtonDefault
}

func (r *SetChatMenuButtonRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.ChatID)
		fw, _ := w.CreateFormField("chat_id")
		fw.Write(b)
	}
	if r.MenuButton != nil {
		{
			b, _ := json.Marshal(r.MenuButton)
			fw, _ := w.CreateFormField("menu_button")
			fw.Write(b)
		}
	}
}

// see Bot.GetChatMenuButton(ctx, &GetChatMenuButtonRequest{})
type GetChatMenuButtonRequest struct {
	ChatID int64 `json:"chat_id,omitempty"` // Unique identifier for the target private chat. If not specified, default bot's menu button will be returned
}

func (r *GetChatMenuButtonRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.ChatID)
		fw, _ := w.CreateFormField("chat_id")
		fw.Write(b)
	}
}

// see Bot.SetMyDefaultAdministratorRights(ctx, &SetMyDefaultAdministratorRightsRequest{})
type SetMyDefaultAdministratorRightsRequest struct {
	Rights      *ChatAdministratorRights `json:"rights,omitempty"`       // A JSON-serialized object describing new default administrator rights. If not specified, the default administrator rights will be cleared.
	ForChannels bool                     `json:"for_channels,omitempty"` // Pass True to change the default administrator rights of the bot in channels. Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
}

func (r *SetMyDefaultAdministratorRightsRequest) writeMultipart(w *multipart.Writer) {
	if r.Rights != nil {
		{
			b, _ := json.Marshal(r.Rights)
			fw, _ := w.CreateFormField("rights")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.ForChannels)
		fw, _ := w.CreateFormField("for_channels")
		fw.Write(b)
	}
}

// see Bot.GetMyDefaultAdministratorRights(ctx, &GetMyDefaultAdministratorRightsRequest{})
type GetMyDefaultAdministratorRightsRequest struct {
	ForChannels bool `json:"for_channels,omitempty"` // Pass True to get default administrator rights of the bot in channels. Otherwise, default administrator rights of the bot for groups and supergroups will be returned.
}

func (r *GetMyDefaultAdministratorRightsRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.ForChannels)
		fw, _ := w.CreateFormField("for_channels")
		fw.Write(b)
	}
}

// see Bot.EditMessageText(ctx, &EditMessageTextRequest{})
type EditMessageTextRequest struct {
	BusinessConnectionID string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID               ChatID                `json:"chat_id,omitempty"`                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int                   `json:"message_id,omitempty"`             // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID      string                `json:"inline_message_id,omitempty"`      // Required if chat_id and message_id are not specified. Identifier of the inline message
	Text                 string                `json:"text"`                             // New text of the message, 1-4096 characters after entities parsing
	ParseMode            ParseMode             `json:"parse_mode,omitempty"`             // Mode for parsing entities in the message text. See formatting options for more details.
	Entities             []MessageEntity       `json:"entities,omitempty"`               // A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	LinkPreviewOptions   *LinkPreviewOptions   `json:"link_preview_options,omitempty"`   // Link preview generation options for the message
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for an inline keyboard.
}

func (r *EditMessageTextRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}
	if r.InlineMessageID != "" {
		w.WriteField("inline_message_id", r.InlineMessageID)
	}

	w.WriteField("text", r.Text)
	w.WriteField("parse_mode", string(r.ParseMode))
	if r.Entities != nil {
		{
			b, _ := json.Marshal(r.Entities)
			fw, _ := w.CreateFormField("entities")
			fw.Write(b)
		}
	}
	if r.LinkPreviewOptions != nil {
		{
			b, _ := json.Marshal(r.LinkPreviewOptions)
			fw, _ := w.CreateFormField("link_preview_options")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.EditMessageCaption(ctx, &EditMessageCaptionRequest{})
type EditMessageCaptionRequest struct {
	BusinessConnectionID  string                `json:"business_connection_id,omitempty"`   // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID                ChatID                `json:"chat_id,omitempty"`                  // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID             int                   `json:"message_id,omitempty"`               // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID       string                `json:"inline_message_id,omitempty"`        // Required if chat_id and message_id are not specified. Identifier of the inline message
	Caption               string                `json:"caption,omitempty"`                  // New caption of the message, 0-1024 characters after entities parsing
	ParseMode             ParseMode             `json:"parse_mode,omitempty"`               // Mode for parsing entities in the message caption. See formatting options for more details.
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`         // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"` // Pass True, if the caption must be shown above the message media. Supported only for animation, photo and video messages.
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`             // A JSON-serialized object for an inline keyboard.
}

func (r *EditMessageCaptionRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}
	if r.InlineMessageID != "" {
		w.WriteField("inline_message_id", r.InlineMessageID)
	}
	if r.Caption != "" {
		w.WriteField("caption", r.Caption)
	}
	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		{
			b, _ := json.Marshal(r.CaptionEntities)
			fw, _ := w.CreateFormField("caption_entities")
			fw.Write(b)
		}
	}

	{
		b, _ := json.Marshal(r.ShowCaptionAboveMedia)
		fw, _ := w.CreateFormField("show_caption_above_media")
		fw.Write(b)
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.EditMessageMedia(ctx, &EditMessageMediaRequest{})
type EditMessageMediaRequest struct {
	BusinessConnectionID string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID               ChatID                `json:"chat_id,omitempty"`                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int                   `json:"message_id,omitempty"`             // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID      string                `json:"inline_message_id,omitempty"`      // Required if chat_id and message_id are not specified. Identifier of the inline message
	Media                InputMedia            `json:"media"`                            // A JSON-serialized object for a new media content of the message
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for a new inline keyboard.
}

func (r *EditMessageMediaRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}
	if r.InlineMessageID != "" {
		w.WriteField("inline_message_id", r.InlineMessageID)
	}

	{
		inputFile := r.Media.getMedia()

		if inputFile.Reader != nil {
			fw, _ := w.CreateFormFile("media", inputFile.Reader.Name())
			io.Copy(fw, inputFile.Reader)
			r.Media.setMedia("attach://media")
		}

		b, _ := json.Marshal(r.Media)
		fw, _ := w.CreateFormField("media")
		fw.Write(b)
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.EditMessageLiveLocation(ctx, &EditMessageLiveLocationRequest{})
type EditMessageLiveLocationRequest struct {
	BusinessConnectionID string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID               ChatID                `json:"chat_id,omitempty"`                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int                   `json:"message_id,omitempty"`             // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID      string                `json:"inline_message_id,omitempty"`      // Required if chat_id and message_id are not specified. Identifier of the inline message
	Latitude             float64               `json:"latitude"`                         // Latitude of new location
	Longitude            float64               `json:"longitude"`                        // Longitude of new location
	LivePeriod           int                   `json:"live_period,omitempty"`            // New period in seconds during which the location can be updated, starting from the message send date. If 0x7FFFFFFF is specified, then the location can be updated forever. Otherwise, the new value must not exceed the current live_period by more than a day, and the live location expiration date must remain within the next 90 days. If not specified, then live_period remains unchanged
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`    // The radius of uncertainty for the location, measured in meters; 0-1500
	Heading              int                   `json:"heading,omitempty"`                // Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int                   `json:"proximity_alert_radius,omitempty"` // The maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for a new inline keyboard.
}

func (r *EditMessageLiveLocationRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}
	if r.InlineMessageID != "" {
		w.WriteField("inline_message_id", r.InlineMessageID)
	}

	{
		b, _ := json.Marshal(r.Latitude)
		fw, _ := w.CreateFormField("latitude")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Longitude)
		fw, _ := w.CreateFormField("longitude")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.LivePeriod)
		fw, _ := w.CreateFormField("live_period")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.HorizontalAccuracy)
		fw, _ := w.CreateFormField("horizontal_accuracy")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Heading)
		fw, _ := w.CreateFormField("heading")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProximityAlertRadius)
		fw, _ := w.CreateFormField("proximity_alert_radius")
		fw.Write(b)
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.StopMessageLiveLocation(ctx, &StopMessageLiveLocationRequest{})
type StopMessageLiveLocationRequest struct {
	BusinessConnectionID string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID               ChatID                `json:"chat_id,omitempty"`                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int                   `json:"message_id,omitempty"`             // Required if inline_message_id is not specified. Identifier of the message with live location to stop
	InlineMessageID      string                `json:"inline_message_id,omitempty"`      // Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for a new inline keyboard.
}

func (r *StopMessageLiveLocationRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}
	if r.InlineMessageID != "" {
		w.WriteField("inline_message_id", r.InlineMessageID)
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.EditMessageReplyMarkup(ctx, &EditMessageReplyMarkupRequest{})
type EditMessageReplyMarkupRequest struct {
	BusinessConnectionID string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID               ChatID                `json:"chat_id,omitempty"`                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int                   `json:"message_id,omitempty"`             // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID      string                `json:"inline_message_id,omitempty"`      // Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for an inline keyboard.
}

func (r *EditMessageReplyMarkupRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}
	if r.InlineMessageID != "" {
		w.WriteField("inline_message_id", r.InlineMessageID)
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.StopPoll(ctx, &StopPollRequest{})
type StopPollRequest struct {
	BusinessConnectionID string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID               ChatID                `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int                   `json:"message_id"`                       // Identifier of the original message with the poll
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for a new message inline keyboard.
}

func (r *StopPollRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.DeleteMessage(ctx, &DeleteMessageRequest{})
type DeleteMessageRequest struct {
	ChatID    ChatID `json:"chat_id"`    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID int    `json:"message_id"` // Identifier of the message to delete
}

func (r *DeleteMessageRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}
}

// see Bot.DeleteMessages(ctx, &DeleteMessagesRequest{})
type DeleteMessagesRequest struct {
	ChatID     ChatID `json:"chat_id"`     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageIds []int  `json:"message_ids"` // A JSON-serialized list of 1-100 identifiers of messages to delete. See deleteMessage for limitations on which messages can be deleted
}

func (r *DeleteMessagesRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageIds)
		fw, _ := w.CreateFormField("message_ids")
		fw.Write(b)
	}
}

// see Bot.SendSticker(ctx, &SendStickerRequest{})
type SendStickerRequest struct {
	BusinessConnectionID string           `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           `json:"chat_id"`                          // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int64            `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Sticker              InputFile        `json:"sticker"`                          // Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP sticker from the Internet, or upload a new .WEBP, .TGS, or .WEBM sticker using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files. Video and animated stickers can't be sent via an HTTP URL.
	Emoji                string           `json:"emoji,omitempty"`                  // Emoji associated with the sticker; only for just uploaded stickers
	DisableNotification  bool             `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      string           `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          Markup           `json:"reply_markup,omitempty"`           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendStickerRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	if r.Sticker.FileID != "" {
		w.WriteField("sticker", r.Sticker.FileID)
	} else if r.Sticker.Reader != nil {
		fw, _ := w.CreateFormFile("sticker", r.Sticker.Reader.Name())
		io.Copy(fw, r.Sticker.Reader)
	}
	if r.Emoji != "" {
		w.WriteField("emoji", r.Emoji)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.GetStickerSet(ctx, &GetStickerSetRequest{})
type GetStickerSetRequest struct {
	Name string `json:"name"` // Name of the sticker set
}

func (r *GetStickerSetRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("name", r.Name)
}

// see Bot.GetCustomEmojiStickers(ctx, &GetCustomEmojiStickersRequest{})
type GetCustomEmojiStickersRequest struct {
	CustomEmojiIds []string `json:"custom_emoji_ids"` // A JSON-serialized list of custom emoji identifiers. At most 200 custom emoji identifiers can be specified.
}

func (r *GetCustomEmojiStickersRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.CustomEmojiIds)
		fw, _ := w.CreateFormField("custom_emoji_ids")
		fw.Write(b)
	}
}

// see Bot.UploadStickerFile(ctx, &UploadStickerFileRequest{})
type UploadStickerFileRequest struct {
	UserID        int64     `json:"user_id"`        // User identifier of sticker file owner
	Sticker       InputFile `json:"sticker"`        // A file with the sticker in .WEBP, .PNG, .TGS, or .WEBM format. See https://core.telegram.org/stickers for technical requirements. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	StickerFormat string    `json:"sticker_format"` // Format of the sticker, must be one of "static", "animated", "video"
}

func (r *UploadStickerFileRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	if r.Sticker.FileID != "" {
		w.WriteField("sticker", r.Sticker.FileID)
	} else if r.Sticker.Reader != nil {
		fw, _ := w.CreateFormFile("sticker", r.Sticker.Reader.Name())
		io.Copy(fw, r.Sticker.Reader)
	}

	w.WriteField("sticker_format", r.StickerFormat)
}

// see Bot.CreateNewStickerSet(ctx, &CreateNewStickerSetRequest{})
type CreateNewStickerSetRequest struct {
	UserID          int64          `json:"user_id"`                    // User identifier of created sticker set owner
	Name            string         `json:"name"`                       // Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only English letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in "_by_<bot_username>". <bot_username> is case insensitive. 1-64 characters.
	Title           string         `json:"title"`                      // Sticker set title, 1-64 characters
	Stickers        []InputSticker `json:"stickers"`                   // A JSON-serialized list of 1-50 initial stickers to be added to the sticker set
	StickerType     string         `json:"sticker_type,omitempty"`     // Type of stickers in the set, pass "regular", "mask", or "custom_emoji". By default, a regular sticker set is created.
	NeedsRepainting bool           `json:"needs_repainting,omitempty"` // Pass True if stickers in the sticker set must be repainted to the color of text when used in messages, the accent color if used as emoji status, white on chat photos, or another appropriate color based on context; for custom emoji sticker sets only
}

func (r *CreateNewStickerSetRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	w.WriteField("name", r.Name)

	w.WriteField("title", r.Title)

	{
		b, _ := json.Marshal(r.Stickers)
		fw, _ := w.CreateFormField("stickers")
		fw.Write(b)
	}
	if r.StickerType != "" {
		w.WriteField("sticker_type", r.StickerType)
	}

	{
		b, _ := json.Marshal(r.NeedsRepainting)
		fw, _ := w.CreateFormField("needs_repainting")
		fw.Write(b)
	}
}

// see Bot.AddStickerToSet(ctx, &AddStickerToSetRequest{})
type AddStickerToSetRequest struct {
	UserID  int64         `json:"user_id"` // User identifier of sticker set owner
	Name    string        `json:"name"`    // Sticker set name
	Sticker *InputSticker `json:"sticker"` // A JSON-serialized object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set isn't changed.
}

func (r *AddStickerToSetRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	w.WriteField("name", r.Name)

	{
		b, _ := json.Marshal(r.Sticker)
		fw, _ := w.CreateFormField("sticker")
		fw.Write(b)
	}
}

// see Bot.SetStickerPositionInSet(ctx, &SetStickerPositionInSetRequest{})
type SetStickerPositionInSetRequest struct {
	Sticker  string `json:"sticker"`  // File identifier of the sticker
	Position int    `json:"position"` // New sticker position in the set, zero-based
}

func (r *SetStickerPositionInSetRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("sticker", r.Sticker)

	{
		b, _ := json.Marshal(r.Position)
		fw, _ := w.CreateFormField("position")
		fw.Write(b)
	}
}

// see Bot.DeleteStickerFromSet(ctx, &DeleteStickerFromSetRequest{})
type DeleteStickerFromSetRequest struct {
	Sticker string `json:"sticker"` // File identifier of the sticker
}

func (r *DeleteStickerFromSetRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("sticker", r.Sticker)
}

// see Bot.ReplaceStickerInSet(ctx, &ReplaceStickerInSetRequest{})
type ReplaceStickerInSetRequest struct {
	UserID     int64         `json:"user_id"`     // User identifier of the sticker set owner
	Name       string        `json:"name"`        // Sticker set name
	OldSticker string        `json:"old_sticker"` // File identifier of the replaced sticker
	Sticker    *InputSticker `json:"sticker"`     // A JSON-serialized object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set remains unchanged.
}

func (r *ReplaceStickerInSetRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	w.WriteField("name", r.Name)

	w.WriteField("old_sticker", r.OldSticker)

	{
		b, _ := json.Marshal(r.Sticker)
		fw, _ := w.CreateFormField("sticker")
		fw.Write(b)
	}
}

// see Bot.SetStickerEmojiList(ctx, &SetStickerEmojiListRequest{})
type SetStickerEmojiListRequest struct {
	Sticker   string   `json:"sticker"`    // File identifier of the sticker
	EmojiList []string `json:"emoji_list"` // A JSON-serialized list of 1-20 emoji associated with the sticker
}

func (r *SetStickerEmojiListRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("sticker", r.Sticker)

	{
		b, _ := json.Marshal(r.EmojiList)
		fw, _ := w.CreateFormField("emoji_list")
		fw.Write(b)
	}
}

// see Bot.SetStickerKeywords(ctx, &SetStickerKeywordsRequest{})
type SetStickerKeywordsRequest struct {
	Sticker  string   `json:"sticker"`            // File identifier of the sticker
	Keywords []string `json:"keywords,omitempty"` // A JSON-serialized list of 0-20 search keywords for the sticker with total length of up to 64 characters
}

func (r *SetStickerKeywordsRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("sticker", r.Sticker)
	if r.Keywords != nil {
		{
			b, _ := json.Marshal(r.Keywords)
			fw, _ := w.CreateFormField("keywords")
			fw.Write(b)
		}
	}
}

// see Bot.SetStickerMaskPosition(ctx, &SetStickerMaskPositionRequest{})
type SetStickerMaskPositionRequest struct {
	Sticker      string        `json:"sticker"`                 // File identifier of the sticker
	MaskPosition *MaskPosition `json:"mask_position,omitempty"` // A JSON-serialized object with the position where the mask should be placed on faces. Omit the parameter to remove the mask position.
}

func (r *SetStickerMaskPositionRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("sticker", r.Sticker)
	if r.MaskPosition != nil {
		{
			b, _ := json.Marshal(r.MaskPosition)
			fw, _ := w.CreateFormField("mask_position")
			fw.Write(b)
		}
	}
}

// see Bot.SetStickerSetTitle(ctx, &SetStickerSetTitleRequest{})
type SetStickerSetTitleRequest struct {
	Name  string `json:"name"`  // Sticker set name
	Title string `json:"title"` // Sticker set title, 1-64 characters
}

func (r *SetStickerSetTitleRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("name", r.Name)

	w.WriteField("title", r.Title)
}

// see Bot.SetStickerSetThumbnail(ctx, &SetStickerSetThumbnailRequest{})
type SetStickerSetThumbnailRequest struct {
	Name      string    `json:"name"`                // Sticker set name
	UserID    int64     `json:"user_id"`             // User identifier of the sticker set owner
	Thumbnail InputFile `json:"thumbnail,omitempty"` // A .WEBP or .PNG image with the thumbnail, must be up to 128 kilobytes in size and have a width and height of exactly 100px, or a .TGS animation with a thumbnail up to 32 kilobytes in size (see https://core.telegram.org/stickers#animation-requirements for animated sticker technical requirements), or a .WEBM video with the thumbnail up to 32 kilobytes in size; see https://core.telegram.org/stickers#video-requirements for video sticker technical requirements. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files. Animated and video sticker set thumbnails can't be uploaded via HTTP URL. If omitted, then the thumbnail is dropped and the first sticker is used as the thumbnail.
	Format    string    `json:"format"`              // Format of the thumbnail, must be one of "static" for a .WEBP or .PNG image, "animated" for a .TGS animation, or "video" for a .WEBM video
}

func (r *SetStickerSetThumbnailRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("name", r.Name)

	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	if r.Thumbnail.FileID != "" {
		w.WriteField("thumbnail", r.Thumbnail.FileID)
	} else if r.Thumbnail.Reader != nil {
		fw, _ := w.CreateFormFile("thumbnail", r.Thumbnail.Reader.Name())
		io.Copy(fw, r.Thumbnail.Reader)
	}

	w.WriteField("format", r.Format)
}

// see Bot.SetCustomEmojiStickerSetThumbnail(ctx, &SetCustomEmojiStickerSetThumbnailRequest{})
type SetCustomEmojiStickerSetThumbnailRequest struct {
	Name          string `json:"name"`                      // Sticker set name
	CustomEmojiID string `json:"custom_emoji_id,omitempty"` // Custom emoji identifier of a sticker from the sticker set; pass an empty string to drop the thumbnail and use the first sticker as the thumbnail.
}

func (r *SetCustomEmojiStickerSetThumbnailRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("name", r.Name)
	if r.CustomEmojiID != "" {
		w.WriteField("custom_emoji_id", r.CustomEmojiID)
	}
}

// see Bot.DeleteStickerSet(ctx, &DeleteStickerSetRequest{})
type DeleteStickerSetRequest struct {
	Name string `json:"name"` // Sticker set name
}

func (r *DeleteStickerSetRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("name", r.Name)
}

// see Bot.SendGift(ctx, &SendGiftRequest{})
type SendGiftRequest struct {
	UserID        int64           `json:"user_id,omitempty"`         // Required if chat_id is not specified. Unique identifier of the target user who will receive the gift.
	ChatID        ChatID          `json:"chat_id,omitempty"`         // Required if user_id is not specified. Unique identifier for the chat or username of the channel (in the format @channelusername) that will receive the gift.
	GiftID        string          `json:"gift_id"`                   // Identifier of the gift
	PayForUpgrade bool            `json:"pay_for_upgrade,omitempty"` // Pass True to pay for the gift upgrade from the bot's balance, thereby making the upgrade free for the receiver
	Text          string          `json:"text,omitempty"`            // Text that will be shown along with the gift; 0-128 characters
	TextParseMode string          `json:"text_parse_mode,omitempty"` // Mode for parsing entities in the text. See formatting options for more details. Entities other than "bold", "italic", "underline", "strikethrough", "spoiler", and "custom_emoji" are ignored.
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`   // A JSON-serialized list of special entities that appear in the gift text. It can be specified instead of text_parse_mode. Entities other than "bold", "italic", "underline", "strikethrough", "spoiler", and "custom_emoji" are ignored.
}

func (r *SendGiftRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("gift_id", r.GiftID)

	{
		b, _ := json.Marshal(r.PayForUpgrade)
		fw, _ := w.CreateFormField("pay_for_upgrade")
		fw.Write(b)
	}
	if r.Text != "" {
		w.WriteField("text", r.Text)
	}
	if r.TextParseMode != "" {
		w.WriteField("text_parse_mode", r.TextParseMode)
	}
	if r.TextEntities != nil {
		{
			b, _ := json.Marshal(r.TextEntities)
			fw, _ := w.CreateFormField("text_entities")
			fw.Write(b)
		}
	}
}

// see Bot.VerifyUser(ctx, &VerifyUserRequest{})
type VerifyUserRequest struct {
	UserID            int64  `json:"user_id"`                      // Unique identifier of the target user
	CustomDescription string `json:"custom_description,omitempty"` // Custom description for the verification; 0-70 characters. Must be empty if the organization isn't allowed to provide a custom verification description.
}

func (r *VerifyUserRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}
	if r.CustomDescription != "" {
		w.WriteField("custom_description", r.CustomDescription)
	}
}

// see Bot.VerifyChat(ctx, &VerifyChatRequest{})
type VerifyChatRequest struct {
	ChatID            ChatID `json:"chat_id"`                      // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	CustomDescription string `json:"custom_description,omitempty"` // Custom description for the verification; 0-70 characters. Must be empty if the organization isn't allowed to provide a custom verification description.
}

func (r *VerifyChatRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	if r.CustomDescription != "" {
		w.WriteField("custom_description", r.CustomDescription)
	}
}

// see Bot.RemoveUserVerification(ctx, &RemoveUserVerificationRequest{})
type RemoveUserVerificationRequest struct {
	UserID int64 `json:"user_id"` // Unique identifier of the target user
}

func (r *RemoveUserVerificationRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}
}

// see Bot.RemoveChatVerification(ctx, &RemoveChatVerificationRequest{})
type RemoveChatVerificationRequest struct {
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

func (r *RemoveChatVerificationRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// see Bot.AnswerInlineQuery(ctx, &AnswerInlineQueryRequest{})
type AnswerInlineQueryRequest struct {
	InlineQueryID string                    `json:"inline_query_id"`       // Unique identifier for the answered query
	Results       []InlineQueryResult       `json:"results"`               // A JSON-serialized array of results for the inline query
	CacheTime     int                       `json:"cache_time,omitempty"`  // The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	IsPersonal    bool                      `json:"is_personal,omitempty"` // Pass True if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query.
	NextOffset    string                    `json:"next_offset,omitempty"` // Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don't support pagination. Offset length can't exceed 64 bytes.
	Button        *InlineQueryResultsButton `json:"button,omitempty"`      // A JSON-serialized object describing a button to be shown above inline query results
}

func (r *AnswerInlineQueryRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("inline_query_id", r.InlineQueryID)

	{
		b, _ := json.Marshal(r.Results)
		fw, _ := w.CreateFormField("results")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.CacheTime)
		fw, _ := w.CreateFormField("cache_time")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.IsPersonal)
		fw, _ := w.CreateFormField("is_personal")
		fw.Write(b)
	}
	if r.NextOffset != "" {
		w.WriteField("next_offset", r.NextOffset)
	}
	if r.Button != nil {
		{
			b, _ := json.Marshal(r.Button)
			fw, _ := w.CreateFormField("button")
			fw.Write(b)
		}
	}
}

// see Bot.AnswerWebAppQuery(ctx, &AnswerWebAppQueryRequest{})
type AnswerWebAppQueryRequest struct {
	WebAppQueryID string            `json:"web_app_query_id"` // Unique identifier for the query to be answered
	Result        InlineQueryResult `json:"result"`           // A JSON-serialized object describing the message to be sent
}

func (r *AnswerWebAppQueryRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("web_app_query_id", r.WebAppQueryID)

	{
		b, _ := json.Marshal(r.Result)
		fw, _ := w.CreateFormField("result")
		fw.Write(b)
	}
}

// see Bot.SavePreparedInlineMessage(ctx, &SavePreparedInlineMessageRequest{})
type SavePreparedInlineMessageRequest struct {
	UserID            int64             `json:"user_id"`                       // Unique identifier of the target user that can use the prepared message
	Result            InlineQueryResult `json:"result"`                        // A JSON-serialized object describing the message to be sent
	AllowUserChats    bool              `json:"allow_user_chats,omitempty"`    // Pass True if the message can be sent to private chats with users
	AllowBotChats     bool              `json:"allow_bot_chats,omitempty"`     // Pass True if the message can be sent to private chats with bots
	AllowGroupChats   bool              `json:"allow_group_chats,omitempty"`   // Pass True if the message can be sent to group and supergroup chats
	AllowChannelChats bool              `json:"allow_channel_chats,omitempty"` // Pass True if the message can be sent to channel chats
}

func (r *SavePreparedInlineMessageRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Result)
		fw, _ := w.CreateFormField("result")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowUserChats)
		fw, _ := w.CreateFormField("allow_user_chats")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowBotChats)
		fw, _ := w.CreateFormField("allow_bot_chats")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowGroupChats)
		fw, _ := w.CreateFormField("allow_group_chats")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowChannelChats)
		fw, _ := w.CreateFormField("allow_channel_chats")
		fw.Write(b)
	}
}

// see Bot.SendInvoice(ctx, &SendInvoiceRequest{})
type SendInvoiceRequest struct {
	ChatID                    ChatID                `json:"chat_id"`                                 // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID           int64                 `json:"message_thread_id,omitempty"`             // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Title                     string                `json:"title"`                                   // Product name, 1-32 characters
	Description               string                `json:"description"`                             // Product description, 1-255 characters
	Payload                   string                `json:"payload"`                                 // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	ProviderToken             string                `json:"provider_token,omitempty"`                // Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	Currency                  string                `json:"currency"`                                // Three-letter ISO 4217 currency code, see more on currencies. Pass "XTR" for payments in Telegram Stars.
	Prices                    []LabeledPrice        `json:"prices"`                                  // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
	MaxTipAmount              int                   `json:"max_tip_amount,omitempty"`                // The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
	SuggestedTipAmounts       []int                 `json:"suggested_tip_amounts,omitempty"`         // A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	StartParameter            string                `json:"start_parameter,omitempty"`               // Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a Pay button, allowing multiple users to pay directly from the forwarded message, using the same invoice. If non-empty, forwarded copies of the sent message will have a URL button with a deep link to the bot (instead of a Pay button), with the value used as the start parameter
	ProviderData              string                `json:"provider_data,omitempty"`                 // JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoUrl                  string                `json:"photo_url,omitempty"`                     // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoSize                 int                   `json:"photo_size,omitempty"`                    // Photo size in bytes
	PhotoWidth                int                   `json:"photo_width,omitempty"`                   // Photo width
	PhotoHeight               int                   `json:"photo_height,omitempty"`                  // Photo height
	NeedName                  bool                  `json:"need_name,omitempty"`                     // Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber           bool                  `json:"need_phone_number,omitempty"`             // Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
	NeedEmail                 bool                  `json:"need_email,omitempty"`                    // Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress       bool                  `json:"need_shipping_address,omitempty"`         // Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider,omitempty"` // Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvider       bool                  `json:"send_email_to_provider,omitempty"`        // Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
	IsFlexible                bool                  `json:"is_flexible,omitempty"`                   // Pass True if the final price depends on the shipping method. Ignored for payments in Telegram Stars.
	DisableNotification       bool                  `json:"disable_notification,omitempty"`          // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent            bool                  `json:"protect_content,omitempty"`               // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast        bool                  `json:"allow_paid_broadcast,omitempty"`          // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID           string                `json:"message_effect_id,omitempty"`             // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters           *ReplyParameters      `json:"reply_parameters,omitempty"`              // Description of the message to reply to
	ReplyMarkup               *InlineKeyboardMarkup `json:"reply_markup,omitempty"`                  // A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
}

func (r *SendInvoiceRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	w.WriteField("title", r.Title)

	w.WriteField("description", r.Description)

	w.WriteField("payload", r.Payload)
	if r.ProviderToken != "" {
		w.WriteField("provider_token", r.ProviderToken)
	}

	w.WriteField("currency", r.Currency)

	{
		b, _ := json.Marshal(r.Prices)
		fw, _ := w.CreateFormField("prices")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.MaxTipAmount)
		fw, _ := w.CreateFormField("max_tip_amount")
		fw.Write(b)
	}
	if r.SuggestedTipAmounts != nil {
		{
			b, _ := json.Marshal(r.SuggestedTipAmounts)
			fw, _ := w.CreateFormField("suggested_tip_amounts")
			fw.Write(b)
		}
	}
	if r.StartParameter != "" {
		w.WriteField("start_parameter", r.StartParameter)
	}
	if r.ProviderData != "" {
		w.WriteField("provider_data", r.ProviderData)
	}
	if r.PhotoUrl != "" {
		w.WriteField("photo_url", r.PhotoUrl)
	}

	{
		b, _ := json.Marshal(r.PhotoSize)
		fw, _ := w.CreateFormField("photo_size")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.PhotoWidth)
		fw, _ := w.CreateFormField("photo_width")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.PhotoHeight)
		fw, _ := w.CreateFormField("photo_height")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.NeedName)
		fw, _ := w.CreateFormField("need_name")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.NeedPhoneNumber)
		fw, _ := w.CreateFormField("need_phone_number")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.NeedEmail)
		fw, _ := w.CreateFormField("need_email")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.NeedShippingAddress)
		fw, _ := w.CreateFormField("need_shipping_address")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.SendPhoneNumberToProvider)
		fw, _ := w.CreateFormField("send_phone_number_to_provider")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.SendEmailToProvider)
		fw, _ := w.CreateFormField("send_email_to_provider")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.IsFlexible)
		fw, _ := w.CreateFormField("is_flexible")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.CreateInvoiceLink(ctx, &CreateInvoiceLinkRequest{})
type CreateInvoiceLinkRequest struct {
	BusinessConnectionID      string         `json:"business_connection_id,omitempty"`        // Unique identifier of the business connection on behalf of which the link will be created. For payments in Telegram Stars only.
	Title                     string         `json:"title"`                                   // Product name, 1-32 characters
	Description               string         `json:"description"`                             // Product description, 1-255 characters
	Payload                   string         `json:"payload"`                                 // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	ProviderToken             string         `json:"provider_token,omitempty"`                // Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	Currency                  string         `json:"currency"`                                // Three-letter ISO 4217 currency code, see more on currencies. Pass "XTR" for payments in Telegram Stars.
	Prices                    []LabeledPrice `json:"prices"`                                  // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
	SubscriptionPeriod        int            `json:"subscription_period,omitempty"`           // The number of seconds the subscription will be active for before the next payment. The currency must be set to "XTR" (Telegram Stars) if the parameter is used. Currently, it must always be 2592000 (30 days) if specified. Any number of subscriptions can be active for a given bot at the same time, including multiple concurrent subscriptions from the same user. Subscription price must no exceed 2500 Telegram Stars.
	MaxTipAmount              int            `json:"max_tip_amount,omitempty"`                // The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`         // A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	ProviderData              string         `json:"provider_data,omitempty"`                 // JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoUrl                  string         `json:"photo_url,omitempty"`                     // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoSize                 int            `json:"photo_size,omitempty"`                    // Photo size in bytes
	PhotoWidth                int            `json:"photo_width,omitempty"`                   // Photo width
	PhotoHeight               int            `json:"photo_height,omitempty"`                  // Photo height
	NeedName                  bool           `json:"need_name,omitempty"`                     // Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`             // Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
	NeedEmail                 bool           `json:"need_email,omitempty"`                    // Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`         // Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"` // Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`        // Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
	IsFlexible                bool           `json:"is_flexible,omitempty"`                   // Pass True if the final price depends on the shipping method. Ignored for payments in Telegram Stars.
}

func (r *CreateInvoiceLinkRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}

	w.WriteField("title", r.Title)

	w.WriteField("description", r.Description)

	w.WriteField("payload", r.Payload)
	if r.ProviderToken != "" {
		w.WriteField("provider_token", r.ProviderToken)
	}

	w.WriteField("currency", r.Currency)

	{
		b, _ := json.Marshal(r.Prices)
		fw, _ := w.CreateFormField("prices")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.SubscriptionPeriod)
		fw, _ := w.CreateFormField("subscription_period")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.MaxTipAmount)
		fw, _ := w.CreateFormField("max_tip_amount")
		fw.Write(b)
	}
	if r.SuggestedTipAmounts != nil {
		{
			b, _ := json.Marshal(r.SuggestedTipAmounts)
			fw, _ := w.CreateFormField("suggested_tip_amounts")
			fw.Write(b)
		}
	}
	if r.ProviderData != "" {
		w.WriteField("provider_data", r.ProviderData)
	}
	if r.PhotoUrl != "" {
		w.WriteField("photo_url", r.PhotoUrl)
	}

	{
		b, _ := json.Marshal(r.PhotoSize)
		fw, _ := w.CreateFormField("photo_size")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.PhotoWidth)
		fw, _ := w.CreateFormField("photo_width")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.PhotoHeight)
		fw, _ := w.CreateFormField("photo_height")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.NeedName)
		fw, _ := w.CreateFormField("need_name")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.NeedPhoneNumber)
		fw, _ := w.CreateFormField("need_phone_number")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.NeedEmail)
		fw, _ := w.CreateFormField("need_email")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.NeedShippingAddress)
		fw, _ := w.CreateFormField("need_shipping_address")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.SendPhoneNumberToProvider)
		fw, _ := w.CreateFormField("send_phone_number_to_provider")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.SendEmailToProvider)
		fw, _ := w.CreateFormField("send_email_to_provider")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.IsFlexible)
		fw, _ := w.CreateFormField("is_flexible")
		fw.Write(b)
	}
}

// see Bot.AnswerShippingQuery(ctx, &AnswerShippingQueryRequest{})
type AnswerShippingQueryRequest struct {
	ShippingQueryID string           `json:"shipping_query_id"`          // Unique identifier for the query to be answered
	Ok              bool             `json:"ok"`                         // Pass True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
	ShippingOptions []ShippingOption `json:"shipping_options,omitempty"` // Required if ok is True. A JSON-serialized array of available shipping options.
	ErrorMessage    string           `json:"error_message,omitempty"`    // Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable"). Telegram will display this message to the user.
}

func (r *AnswerShippingQueryRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("shipping_query_id", r.ShippingQueryID)

	{
		b, _ := json.Marshal(r.Ok)
		fw, _ := w.CreateFormField("ok")
		fw.Write(b)
	}
	if r.ShippingOptions != nil {
		{
			b, _ := json.Marshal(r.ShippingOptions)
			fw, _ := w.CreateFormField("shipping_options")
			fw.Write(b)
		}
	}
	if r.ErrorMessage != "" {
		w.WriteField("error_message", r.ErrorMessage)
	}
}

// see Bot.AnswerPreCheckoutQuery(ctx, &AnswerPreCheckoutQueryRequest{})
type AnswerPreCheckoutQueryRequest struct {
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`   // Unique identifier for the query to be answered
	Ok                 bool   `json:"ok"`                      // Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
	ErrorMessage       string `json:"error_message,omitempty"` // Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
}

func (r *AnswerPreCheckoutQueryRequest) writeMultipart(w *multipart.Writer) {
	w.WriteField("pre_checkout_query_id", r.PreCheckoutQueryID)

	{
		b, _ := json.Marshal(r.Ok)
		fw, _ := w.CreateFormField("ok")
		fw.Write(b)
	}
	if r.ErrorMessage != "" {
		w.WriteField("error_message", r.ErrorMessage)
	}
}

// see Bot.GetStarTransactions(ctx, &GetStarTransactionsRequest{})
type GetStarTransactionsRequest struct {
	Offset int64 `json:"offset,omitempty"` // Number of transactions to skip in the response
	Limit  int   `json:"limit,omitempty"`  // The maximum number of transactions to be retrieved. Values between 1-100 are accepted. Defaults to 100.
}

func (r *GetStarTransactionsRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.Offset)
		fw, _ := w.CreateFormField("offset")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Limit)
		fw, _ := w.CreateFormField("limit")
		fw.Write(b)
	}
}

// see Bot.RefundStarPayment(ctx, &RefundStarPaymentRequest{})
type RefundStarPaymentRequest struct {
	UserID                  int64  `json:"user_id"`                    // Identifier of the user whose payment will be refunded
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"` // Telegram payment identifier
}

func (r *RefundStarPaymentRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	w.WriteField("telegram_payment_charge_id", r.TelegramPaymentChargeID)
}

// see Bot.EditUserStarSubscription(ctx, &EditUserStarSubscriptionRequest{})
type EditUserStarSubscriptionRequest struct {
	UserID                  int64  `json:"user_id"`                    // Identifier of the user whose subscription will be edited
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"` // Telegram payment identifier for the subscription
	IsCanceled              bool   `json:"is_canceled"`                // Pass True to cancel extension of the user subscription; the subscription must be active up to the end of the current subscription period. Pass False to allow the user to re-enable a subscription that was previously canceled by the bot.
}

func (r *EditUserStarSubscriptionRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	w.WriteField("telegram_payment_charge_id", r.TelegramPaymentChargeID)

	{
		b, _ := json.Marshal(r.IsCanceled)
		fw, _ := w.CreateFormField("is_canceled")
		fw.Write(b)
	}
}

// see Bot.SetPassportDataErrors(ctx, &SetPassportDataErrorsRequest{})
type SetPassportDataErrorsRequest struct {
	UserID int64                  `json:"user_id"` // User identifier
	Errors []PassportElementError `json:"errors"`  // A JSON-serialized array describing the errors
}

func (r *SetPassportDataErrorsRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Errors)
		fw, _ := w.CreateFormField("errors")
		fw.Write(b)
	}
}

// see Bot.SendGame(ctx, &SendGameRequest{})
type SendGameRequest struct {
	BusinessConnectionID string                `json:"business_connection_id,omitempty"` // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               int64                 `json:"chat_id"`                          // Unique identifier for the target chat
	MessageThreadID      int64                 `json:"message_thread_id,omitempty"`      // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	GameShortName        string                `json:"game_short_name"`                  // Short name of the game, serves as the unique identifier for the game. Set up your games via @BotFather.
	DisableNotification  bool                  `json:"disable_notification,omitempty"`   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool                  `json:"protect_content,omitempty"`        // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool                  `json:"allow_paid_broadcast,omitempty"`   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      string                `json:"message_effect_id,omitempty"`      // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters      `json:"reply_parameters,omitempty"`       // Description of the message to reply to
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // A JSON-serialized object for an inline keyboard. If empty, one 'Play game_title' button will be shown. If not empty, the first button must launch the game.
}

func (r *SendGameRequest) writeMultipart(w *multipart.Writer) {
	if r.BusinessConnectionID != "" {
		w.WriteField("business_connection_id", r.BusinessConnectionID)
	}

	{
		b, _ := json.Marshal(r.ChatID)
		fw, _ := w.CreateFormField("chat_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.MessageThreadID)
		fw, _ := w.CreateFormField("message_thread_id")
		fw.Write(b)
	}

	w.WriteField("game_short_name", r.GameShortName)

	{
		b, _ := json.Marshal(r.DisableNotification)
		fw, _ := w.CreateFormField("disable_notification")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ProtectContent)
		fw, _ := w.CreateFormField("protect_content")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.AllowPaidBroadcast)
		fw, _ := w.CreateFormField("allow_paid_broadcast")
		fw.Write(b)
	}
	if r.MessageEffectID != "" {
		w.WriteField("message_effect_id", r.MessageEffectID)
	}
	if r.ReplyParameters != nil {
		{
			b, _ := json.Marshal(r.ReplyParameters)
			fw, _ := w.CreateFormField("reply_parameters")
			fw.Write(b)
		}
	}
	if r.ReplyMarkup != nil {
		{
			b, _ := json.Marshal(r.ReplyMarkup)
			fw, _ := w.CreateFormField("reply_markup")
			fw.Write(b)
		}
	}
}

// see Bot.SetGameScore(ctx, &SetGameScoreRequest{})
type SetGameScoreRequest struct {
	UserID             int64  `json:"user_id"`                        // User identifier
	Score              int    `json:"score"`                          // New score, must be non-negative
	Force              bool   `json:"force,omitempty"`                // Pass True if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"` // Pass True if the game message should not be automatically edited to include the current scoreboard
	ChatID             int64  `json:"chat_id,omitempty"`              // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageID          int    `json:"message_id,omitempty"`           // Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageID    string `json:"inline_message_id,omitempty"`    // Required if chat_id and message_id are not specified. Identifier of the inline message
}

func (r *SetGameScoreRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Score)
		fw, _ := w.CreateFormField("score")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.Force)
		fw, _ := w.CreateFormField("force")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.DisableEditMessage)
		fw, _ := w.CreateFormField("disable_edit_message")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ChatID)
		fw, _ := w.CreateFormField("chat_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}
	if r.InlineMessageID != "" {
		w.WriteField("inline_message_id", r.InlineMessageID)
	}
}

// see Bot.GetGameHighScores(ctx, &GetGameHighScoresRequest{})
type GetGameHighScoresRequest struct {
	UserID          int64  `json:"user_id"`                     // Target user id
	ChatID          int64  `json:"chat_id,omitempty"`           // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageID       int    `json:"message_id,omitempty"`        // Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageID string `json:"inline_message_id,omitempty"` // Required if chat_id and message_id are not specified. Identifier of the inline message
}

func (r *GetGameHighScoresRequest) writeMultipart(w *multipart.Writer) {
	{
		b, _ := json.Marshal(r.UserID)
		fw, _ := w.CreateFormField("user_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.ChatID)
		fw, _ := w.CreateFormField("chat_id")
		fw.Write(b)
	}

	{
		b, _ := json.Marshal(r.MessageID)
		fw, _ := w.CreateFormField("message_id")
		fw.Write(b)
	}
	if r.InlineMessageID != "" {
		w.WriteField("inline_message_id", r.InlineMessageID)
	}
}
