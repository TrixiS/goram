package types

import "mime/multipart"
import "encoding/json"

// use Bot.GetUpdates(ctx, &GetUpdatesRequest{})
type GetUpdatesRequest struct {
	Offset         int64    // Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will be forgotten.
	Limit          int64    // Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Timeout        int64    // Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	AllowedUpdates []string // A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default). If not specified, the previous setting will be used. Please note that this parameter doesn't affect updates created before the call to getUpdates, so unwanted updates may be received for a short period of time.
}

func (r *GetUpdatesRequest) WriteMultipart(w *multipart.Writer) {
	offset, _ := json.Marshal(r.Offset)
	w.WriteField("offset", string(offset))
	limit, _ := json.Marshal(r.Limit)
	w.WriteField("limit", string(limit))
	timeout, _ := json.Marshal(r.Timeout)
	w.WriteField("timeout", string(timeout))
	if r.AllowedUpdates != nil {
		allowed_updates, _ := json.Marshal(r.AllowedUpdates)
		w.WriteField("allowed_updates", string(allowed_updates))
	}
}

// use Bot.SetWebhook(ctx, &SetWebhookRequest{})
type SetWebhookRequest struct {
	Url                string    // HTTPS URL to send updates to. Use an empty string to remove webhook integration
	Certificate        InputFile // Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	IpAddress          string    // The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	MaxConnections     int64     // The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	AllowedUpdates     []string  // A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default). If not specified, the previous setting will be used. Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	DropPendingUpdates bool      // Pass True to drop all pending updates
	SecretToken        string    // A secret token to be sent in a header "X-Telegram-Bot-Api-Secret-Token" in every webhook request, 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a webhook set by you.
}

func (r *SetWebhookRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("url", r.Url)

	if s, ok := r.Certificate.(string); ok {
		w.WriteField("certificate", s)
	} else {
		// fw, _ := w.CreateFormFile("certificate", "todo.jpeg")
	}

	w.WriteField("ip_address", r.IpAddress)
	max_connections, _ := json.Marshal(r.MaxConnections)
	w.WriteField("max_connections", string(max_connections))
	if r.AllowedUpdates != nil {
		allowed_updates, _ := json.Marshal(r.AllowedUpdates)
		w.WriteField("allowed_updates", string(allowed_updates))
	}
	drop_pending_updates, _ := json.Marshal(r.DropPendingUpdates)
	w.WriteField("drop_pending_updates", string(drop_pending_updates))

	w.WriteField("secret_token", r.SecretToken)
}

// use Bot.DeleteWebhook(ctx, &DeleteWebhookRequest{})
type DeleteWebhookRequest struct {
	DropPendingUpdates bool // Pass True to drop all pending updates
}

func (r *DeleteWebhookRequest) WriteMultipart(w *multipart.Writer) {
	drop_pending_updates, _ := json.Marshal(r.DropPendingUpdates)
	w.WriteField("drop_pending_updates", string(drop_pending_updates))
}

// use Bot.SendMessage(ctx, &SendMessageRequest{})
type SendMessageRequest struct {
	BusinessConnectionID string              // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID              // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int                 // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Text                 string              // Text of the message to be sent, 1-4096 characters after entities parsing
	ParseMode            ParseMode           // Mode for parsing entities in the message text. See formatting options for more details.
	Entities             []MessageEntity     // A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	LinkPreviewOptions   *LinkPreviewOptions // Link preview generation options for the message
	DisableNotification  bool                // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool                // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool                // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      int                 // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters    // Description of the message to reply to
	ReplyMarkup          Markup              // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendMessageRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	w.WriteField("text", r.Text)

	w.WriteField("parse_mode", string(r.ParseMode))
	if r.Entities != nil {
		entities, _ := json.Marshal(r.Entities)
		w.WriteField("entities", string(entities))
	}
	if r.LinkPreviewOptions != nil {
		link_preview_options, _ := json.Marshal(r.LinkPreviewOptions)
		w.WriteField("link_preview_options", string(link_preview_options))
	}
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.ForwardMessage(ctx, &ForwardMessageRequest{})
type ForwardMessageRequest struct {
	ChatID              ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID     int    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatID          int64  // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	VideoStartTimestamp int64  // New start timestamp for the forwarded video in the message
	DisableNotification bool   // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent      bool   // Protects the contents of the forwarded message from forwarding and saving
	MessageID           int    // Message identifier in the chat specified in from_chat_id
}

func (r *ForwardMessageRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))
	from_chat_id, _ := json.Marshal(r.FromChatID)
	w.WriteField("from_chat_id", string(from_chat_id))
	video_start_timestamp, _ := json.Marshal(r.VideoStartTimestamp)
	w.WriteField("video_start_timestamp", string(video_start_timestamp))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
}

// use Bot.ForwardMessages(ctx, &ForwardMessagesRequest{})
type ForwardMessagesRequest struct {
	ChatID              ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID     int    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatID          int64  // Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
	MessageIDs          int    // A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to forward. The identifiers must be specified in a strictly increasing order.
	DisableNotification bool   // Sends the messages silently. Users will receive a notification with no sound.
	ProtectContent      bool   // Protects the contents of the forwarded messages from forwarding and saving
}

func (r *ForwardMessagesRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))
	from_chat_id, _ := json.Marshal(r.FromChatID)
	w.WriteField("from_chat_id", string(from_chat_id))
	message_ids, _ := json.Marshal(r.MessageIDs)
	w.WriteField("message_ids", string(message_ids))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
}

// use Bot.CopyMessage(ctx, &CopyMessageRequest{})
type CopyMessageRequest struct {
	ChatID                ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID       int              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatID            int64            // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	MessageID             int              // Message identifier in the chat specified in from_chat_id
	VideoStartTimestamp   int64            // New start timestamp for the copied video in the message
	Caption               string           // New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
	ParseMode             ParseMode        // Mode for parsing entities in the new caption. See formatting options for more details.
	CaptionEntities       []MessageEntity  // A JSON-serialized list of special entities that appear in the new caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             // Pass True, if the caption must be shown above the message media. Ignored if a new caption isn't specified.
	DisableNotification   bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	ReplyParameters       *ReplyParameters // Description of the message to reply to
	ReplyMarkup           Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *CopyMessageRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))
	from_chat_id, _ := json.Marshal(r.FromChatID)
	w.WriteField("from_chat_id", string(from_chat_id))
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
	video_start_timestamp, _ := json.Marshal(r.VideoStartTimestamp)
	w.WriteField("video_start_timestamp", string(video_start_timestamp))

	w.WriteField("caption", r.Caption)

	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		caption_entities, _ := json.Marshal(r.CaptionEntities)
		w.WriteField("caption_entities", string(caption_entities))
	}
	show_caption_above_media, _ := json.Marshal(r.ShowCaptionAboveMedia)
	w.WriteField("show_caption_above_media", string(show_caption_above_media))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.CopyMessages(ctx, &CopyMessagesRequest{})
type CopyMessagesRequest struct {
	ChatID              ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID     int    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	FromChatID          int64  // Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
	MessageIDs          int    // A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to copy. The identifiers must be specified in a strictly increasing order.
	DisableNotification bool   // Sends the messages silently. Users will receive a notification with no sound.
	ProtectContent      bool   // Protects the contents of the sent messages from forwarding and saving
	RemoveCaption       bool   // Pass True to copy the messages without their captions
}

func (r *CopyMessagesRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))
	from_chat_id, _ := json.Marshal(r.FromChatID)
	w.WriteField("from_chat_id", string(from_chat_id))
	message_ids, _ := json.Marshal(r.MessageIDs)
	w.WriteField("message_ids", string(message_ids))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	remove_caption, _ := json.Marshal(r.RemoveCaption)
	w.WriteField("remove_caption", string(remove_caption))
}

// use Bot.SendPhoto(ctx, &SendPhotoRequest{})
type SendPhotoRequest struct {
	BusinessConnectionID  string           // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID                ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID       int              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Photo                 InputFile        // Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total. Width and height ratio must be at most 20. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption               string           // Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	ParseMode             ParseMode        // Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities       []MessageEntity  // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             // Pass True, if the caption must be shown above the message media
	HasSpoiler            bool             // Pass True if the photo needs to be covered with a spoiler animation
	DisableNotification   bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID       int              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters       *ReplyParameters // Description of the message to reply to
	ReplyMarkup           Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendPhotoRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	if s, ok := r.Photo.(string); ok {
		w.WriteField("photo", s)
	} else {
		// fw, _ := w.CreateFormFile("photo", "todo.jpeg")
	}

	w.WriteField("caption", r.Caption)

	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		caption_entities, _ := json.Marshal(r.CaptionEntities)
		w.WriteField("caption_entities", string(caption_entities))
	}
	show_caption_above_media, _ := json.Marshal(r.ShowCaptionAboveMedia)
	w.WriteField("show_caption_above_media", string(show_caption_above_media))
	has_spoiler, _ := json.Marshal(r.HasSpoiler)
	w.WriteField("has_spoiler", string(has_spoiler))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SendAudio(ctx, &SendAudioRequest{})
type SendAudioRequest struct {
	BusinessConnectionID string           // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Audio                InputFile        // Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption              string           // Audio caption, 0-1024 characters after entities parsing
	ParseMode            ParseMode        // Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities      []MessageEntity  // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration             int64            // Duration of the audio in seconds
	Performer            string           // Performer
	Title                string           // Track name
	Thumbnail            InputFile        // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	DisableNotification  bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      int              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters // Description of the message to reply to
	ReplyMarkup          Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendAudioRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	if s, ok := r.Audio.(string); ok {
		w.WriteField("audio", s)
	} else {
		// fw, _ := w.CreateFormFile("audio", "todo.jpeg")
	}

	w.WriteField("caption", r.Caption)

	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		caption_entities, _ := json.Marshal(r.CaptionEntities)
		w.WriteField("caption_entities", string(caption_entities))
	}
	duration, _ := json.Marshal(r.Duration)
	w.WriteField("duration", string(duration))

	w.WriteField("performer", r.Performer)

	w.WriteField("title", r.Title)

	if s, ok := r.Thumbnail.(string); ok {
		w.WriteField("thumbnail", s)
	} else {
		// fw, _ := w.CreateFormFile("thumbnail", "todo.jpeg")
	}
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SendDocument(ctx, &SendDocumentRequest{})
type SendDocumentRequest struct {
	BusinessConnectionID        string           // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID                      ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID             int              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Document                    InputFile        // File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail                   InputFile        // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption                     string           // Document caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	ParseMode                   ParseMode        // Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities             []MessageEntity  // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableContentTypeDetection bool             // Disables automatic server-side content type detection for files uploaded using multipart/form-data
	DisableNotification         bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent              bool             // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast          bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID             int              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters             *ReplyParameters // Description of the message to reply to
	ReplyMarkup                 Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendDocumentRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	if s, ok := r.Document.(string); ok {
		w.WriteField("document", s)
	} else {
		// fw, _ := w.CreateFormFile("document", "todo.jpeg")
	}

	if s, ok := r.Thumbnail.(string); ok {
		w.WriteField("thumbnail", s)
	} else {
		// fw, _ := w.CreateFormFile("thumbnail", "todo.jpeg")
	}

	w.WriteField("caption", r.Caption)

	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		caption_entities, _ := json.Marshal(r.CaptionEntities)
		w.WriteField("caption_entities", string(caption_entities))
	}
	disable_content_type_detection, _ := json.Marshal(r.DisableContentTypeDetection)
	w.WriteField("disable_content_type_detection", string(disable_content_type_detection))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SendVideo(ctx, &SendVideoRequest{})
type SendVideoRequest struct {
	BusinessConnectionID  string           // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID                ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID       int              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Video                 InputFile        // Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Duration              int64            // Duration of sent video in seconds
	Width                 int64            // Video width
	Height                int64            // Video height
	Thumbnail             InputFile        // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Cover                 InputFile        // Cover for the video in the message. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	StartTimestamp        int64            // Start timestamp for the video in the message
	Caption               string           // Video caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
	ParseMode             ParseMode        // Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities       []MessageEntity  // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             // Pass True, if the caption must be shown above the message media
	HasSpoiler            bool             // Pass True if the video needs to be covered with a spoiler animation
	SupportsStreaming     bool             // Pass True if the uploaded video is suitable for streaming
	DisableNotification   bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID       int              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters       *ReplyParameters // Description of the message to reply to
	ReplyMarkup           Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendVideoRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	if s, ok := r.Video.(string); ok {
		w.WriteField("video", s)
	} else {
		// fw, _ := w.CreateFormFile("video", "todo.jpeg")
	}
	duration, _ := json.Marshal(r.Duration)
	w.WriteField("duration", string(duration))
	width, _ := json.Marshal(r.Width)
	w.WriteField("width", string(width))
	height, _ := json.Marshal(r.Height)
	w.WriteField("height", string(height))

	if s, ok := r.Thumbnail.(string); ok {
		w.WriteField("thumbnail", s)
	} else {
		// fw, _ := w.CreateFormFile("thumbnail", "todo.jpeg")
	}

	if s, ok := r.Cover.(string); ok {
		w.WriteField("cover", s)
	} else {
		// fw, _ := w.CreateFormFile("cover", "todo.jpeg")
	}
	start_timestamp, _ := json.Marshal(r.StartTimestamp)
	w.WriteField("start_timestamp", string(start_timestamp))

	w.WriteField("caption", r.Caption)

	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		caption_entities, _ := json.Marshal(r.CaptionEntities)
		w.WriteField("caption_entities", string(caption_entities))
	}
	show_caption_above_media, _ := json.Marshal(r.ShowCaptionAboveMedia)
	w.WriteField("show_caption_above_media", string(show_caption_above_media))
	has_spoiler, _ := json.Marshal(r.HasSpoiler)
	w.WriteField("has_spoiler", string(has_spoiler))
	supports_streaming, _ := json.Marshal(r.SupportsStreaming)
	w.WriteField("supports_streaming", string(supports_streaming))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SendAnimation(ctx, &SendAnimationRequest{})
type SendAnimationRequest struct {
	BusinessConnectionID  string           // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID                ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID       int              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Animation             InputFile        // Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Duration              int64            // Duration of sent animation in seconds
	Width                 int64            // Animation width
	Height                int64            // Animation height
	Thumbnail             InputFile        // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption               string           // Animation caption (may also be used when resending animation by file_id), 0-1024 characters after entities parsing
	ParseMode             ParseMode        // Mode for parsing entities in the animation caption. See formatting options for more details.
	CaptionEntities       []MessageEntity  // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             // Pass True, if the caption must be shown above the message media
	HasSpoiler            bool             // Pass True if the animation needs to be covered with a spoiler animation
	DisableNotification   bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID       int              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters       *ReplyParameters // Description of the message to reply to
	ReplyMarkup           Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendAnimationRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	if s, ok := r.Animation.(string); ok {
		w.WriteField("animation", s)
	} else {
		// fw, _ := w.CreateFormFile("animation", "todo.jpeg")
	}
	duration, _ := json.Marshal(r.Duration)
	w.WriteField("duration", string(duration))
	width, _ := json.Marshal(r.Width)
	w.WriteField("width", string(width))
	height, _ := json.Marshal(r.Height)
	w.WriteField("height", string(height))

	if s, ok := r.Thumbnail.(string); ok {
		w.WriteField("thumbnail", s)
	} else {
		// fw, _ := w.CreateFormFile("thumbnail", "todo.jpeg")
	}

	w.WriteField("caption", r.Caption)

	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		caption_entities, _ := json.Marshal(r.CaptionEntities)
		w.WriteField("caption_entities", string(caption_entities))
	}
	show_caption_above_media, _ := json.Marshal(r.ShowCaptionAboveMedia)
	w.WriteField("show_caption_above_media", string(show_caption_above_media))
	has_spoiler, _ := json.Marshal(r.HasSpoiler)
	w.WriteField("has_spoiler", string(has_spoiler))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SendVoice(ctx, &SendVoiceRequest{})
type SendVoiceRequest struct {
	BusinessConnectionID string           // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Voice                InputFile        // Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption              string           // Voice message caption, 0-1024 characters after entities parsing
	ParseMode            ParseMode        // Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities      []MessageEntity  // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration             int64            // Duration of the voice message in seconds
	DisableNotification  bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      int              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters // Description of the message to reply to
	ReplyMarkup          Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendVoiceRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	if s, ok := r.Voice.(string); ok {
		w.WriteField("voice", s)
	} else {
		// fw, _ := w.CreateFormFile("voice", "todo.jpeg")
	}

	w.WriteField("caption", r.Caption)

	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		caption_entities, _ := json.Marshal(r.CaptionEntities)
		w.WriteField("caption_entities", string(caption_entities))
	}
	duration, _ := json.Marshal(r.Duration)
	w.WriteField("duration", string(duration))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SendVideoNote(ctx, &SendVideoNoteRequest{})
type SendVideoNoteRequest struct {
	BusinessConnectionID string           // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	VideoNote            InputFile        // Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files. Sending video notes by a URL is currently unsupported
	Duration             int64            // Duration of sent video in seconds
	Length               int64            // Video width and height, i.e. diameter of the video message
	Thumbnail            InputFile        // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	DisableNotification  bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      int              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters // Description of the message to reply to
	ReplyMarkup          Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendVideoNoteRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	if s, ok := r.VideoNote.(string); ok {
		w.WriteField("video_note", s)
	} else {
		// fw, _ := w.CreateFormFile("video_note", "todo.jpeg")
	}
	duration, _ := json.Marshal(r.Duration)
	w.WriteField("duration", string(duration))
	length, _ := json.Marshal(r.Length)
	w.WriteField("length", string(length))

	if s, ok := r.Thumbnail.(string); ok {
		w.WriteField("thumbnail", s)
	} else {
		// fw, _ := w.CreateFormFile("thumbnail", "todo.jpeg")
	}
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SendPaidMedia(ctx, &SendPaidMediaRequest{})
type SendPaidMediaRequest struct {
	BusinessConnectionID  string           // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID                ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername). If the chat is a channel, all Telegram Star proceeds from this media will be credited to the chat's balance. Otherwise, they will be credited to the bot's balance.
	StarCount             int64            // The number of Telegram Stars that must be paid to buy access to the media; 1-2500
	Media                 []InputPaidMedia // A JSON-serialized array describing the media to be sent; up to 10 items
	Payload               string           // Bot-defined paid media payload, 0-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Caption               string           // Media caption, 0-1024 characters after entities parsing
	ParseMode             ParseMode        // Mode for parsing entities in the media caption. See formatting options for more details.
	CaptionEntities       []MessageEntity  // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             // Pass True, if the caption must be shown above the message media
	DisableNotification   bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool             // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	ReplyParameters       *ReplyParameters // Description of the message to reply to
	ReplyMarkup           Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendPaidMediaRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	star_count, _ := json.Marshal(r.StarCount)
	w.WriteField("star_count", string(star_count))
	if r.Media != nil {
		media, _ := json.Marshal(r.Media)
		w.WriteField("media", string(media))
	}

	w.WriteField("payload", r.Payload)

	w.WriteField("caption", r.Caption)

	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		caption_entities, _ := json.Marshal(r.CaptionEntities)
		w.WriteField("caption_entities", string(caption_entities))
	}
	show_caption_above_media, _ := json.Marshal(r.ShowCaptionAboveMedia)
	w.WriteField("show_caption_above_media", string(show_caption_above_media))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SendMediaGroup(ctx, &SendMediaGroupRequest{})
type SendMediaGroupRequest struct {
	BusinessConnectionID string                 // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID                 // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int                    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Media                []MediaGroupInputMedia // A JSON-serialized array describing messages to be sent, must include 2-10 items
	DisableNotification  bool                   // Sends messages silently. Users will receive a notification with no sound.
	ProtectContent       bool                   // Protects the contents of the sent messages from forwarding and saving
	AllowPaidBroadcast   bool                   // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      int                    // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters       // Description of the message to reply to
}

func (r *SendMediaGroupRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))
	if r.Media != nil {
		media, _ := json.Marshal(r.Media)
		w.WriteField("media", string(media))
	}
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
}

// use Bot.SendLocation(ctx, &SendLocationRequest{})
type SendLocationRequest struct {
	BusinessConnectionID string           // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Latitude             float64          // Latitude of the location
	Longitude            float64          // Longitude of the location
	HorizontalAccuracy   float64          // The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int64            // Period in seconds during which the location will be updated (see Live Locations, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	Heading              int64            // For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int64            // For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	DisableNotification  bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      int              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters // Description of the message to reply to
	ReplyMarkup          Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendLocationRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))
	latitude, _ := json.Marshal(r.Latitude)
	w.WriteField("latitude", string(latitude))
	longitude, _ := json.Marshal(r.Longitude)
	w.WriteField("longitude", string(longitude))
	horizontal_accuracy, _ := json.Marshal(r.HorizontalAccuracy)
	w.WriteField("horizontal_accuracy", string(horizontal_accuracy))
	live_period, _ := json.Marshal(r.LivePeriod)
	w.WriteField("live_period", string(live_period))
	heading, _ := json.Marshal(r.Heading)
	w.WriteField("heading", string(heading))
	proximity_alert_radius, _ := json.Marshal(r.ProximityAlertRadius)
	w.WriteField("proximity_alert_radius", string(proximity_alert_radius))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SendVenue(ctx, &SendVenueRequest{})
type SendVenueRequest struct {
	BusinessConnectionID string           // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Latitude             float64          // Latitude of the venue
	Longitude            float64          // Longitude of the venue
	Title                string           // Name of the venue
	Address              string           // Address of the venue
	FoursquareID         string           // Foursquare identifier of the venue
	FoursquareType       string           // Foursquare type of the venue, if known. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	GooglePlaceID        string           // Google Places identifier of the venue
	GooglePlaceType      string           // Google Places type of the venue. (See supported types.)
	DisableNotification  bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      int              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters // Description of the message to reply to
	ReplyMarkup          Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendVenueRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))
	latitude, _ := json.Marshal(r.Latitude)
	w.WriteField("latitude", string(latitude))
	longitude, _ := json.Marshal(r.Longitude)
	w.WriteField("longitude", string(longitude))

	w.WriteField("title", r.Title)

	w.WriteField("address", r.Address)

	w.WriteField("foursquare_id", r.FoursquareID)

	w.WriteField("foursquare_type", r.FoursquareType)

	w.WriteField("google_place_id", r.GooglePlaceID)

	w.WriteField("google_place_type", r.GooglePlaceType)
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SendContact(ctx, &SendContactRequest{})
type SendContactRequest struct {
	BusinessConnectionID string           // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	PhoneNumber          string           // Contact's phone number
	FirstName            string           // Contact's first name
	LastName             string           // Contact's last name
	Vcard                string           // Additional data about the contact in the form of a vCard, 0-2048 bytes
	DisableNotification  bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      int              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters // Description of the message to reply to
	ReplyMarkup          Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendContactRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	w.WriteField("phone_number", r.PhoneNumber)

	w.WriteField("first_name", r.FirstName)

	w.WriteField("last_name", r.LastName)

	w.WriteField("vcard", r.Vcard)
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SendPoll(ctx, &SendPollRequest{})
type SendPollRequest struct {
	BusinessConnectionID  string            // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID                ChatID            // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID       int               // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Question              string            // Poll question, 1-300 characters
	QuestionParseMode     string            // Mode for parsing entities in the question. See formatting options for more details. Currently, only custom emoji entities are allowed
	QuestionEntities      []MessageEntity   // A JSON-serialized list of special entities that appear in the poll question. It can be specified instead of question_parse_mode
	Options               []InputPollOption // A JSON-serialized list of 2-10 answer options
	IsAnonymous           bool              // True, if the poll needs to be anonymous, defaults to True
	Type                  string            // Poll type, "quiz" or "regular", defaults to "regular"
	AllowsMultipleAnswers bool              // True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	CorrectOptionID       int64             // 0-based identifier of the correct answer option, required for polls in quiz mode
	Explanation           string            // Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	ExplanationParseMode  string            // Mode for parsing entities in the explanation. See formatting options for more details.
	ExplanationEntities   []MessageEntity   // A JSON-serialized list of special entities that appear in the poll explanation. It can be specified instead of explanation_parse_mode
	OpenPeriod            int64             // Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with close_date.
	CloseDate             int64             // Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with open_period.
	IsClosed              bool              // Pass True if the poll needs to be immediately closed. This can be useful for poll preview.
	DisableNotification   bool              // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent        bool              // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast    bool              // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID       int               // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters       *ReplyParameters  // Description of the message to reply to
	ReplyMarkup           Markup            // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendPollRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	w.WriteField("question", r.Question)

	w.WriteField("question_parse_mode", r.QuestionParseMode)
	if r.QuestionEntities != nil {
		question_entities, _ := json.Marshal(r.QuestionEntities)
		w.WriteField("question_entities", string(question_entities))
	}
	if r.Options != nil {
		options, _ := json.Marshal(r.Options)
		w.WriteField("options", string(options))
	}
	is_anonymous, _ := json.Marshal(r.IsAnonymous)
	w.WriteField("is_anonymous", string(is_anonymous))

	w.WriteField("type", r.Type)
	allows_multiple_answers, _ := json.Marshal(r.AllowsMultipleAnswers)
	w.WriteField("allows_multiple_answers", string(allows_multiple_answers))
	correct_option_id, _ := json.Marshal(r.CorrectOptionID)
	w.WriteField("correct_option_id", string(correct_option_id))

	w.WriteField("explanation", r.Explanation)

	w.WriteField("explanation_parse_mode", r.ExplanationParseMode)
	if r.ExplanationEntities != nil {
		explanation_entities, _ := json.Marshal(r.ExplanationEntities)
		w.WriteField("explanation_entities", string(explanation_entities))
	}
	open_period, _ := json.Marshal(r.OpenPeriod)
	w.WriteField("open_period", string(open_period))
	close_date, _ := json.Marshal(r.CloseDate)
	w.WriteField("close_date", string(close_date))
	is_closed, _ := json.Marshal(r.IsClosed)
	w.WriteField("is_closed", string(is_closed))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SendDice(ctx, &SendDiceRequest{})
type SendDiceRequest struct {
	BusinessConnectionID string           // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Emoji                string           // Emoji on which the dice throw animation is based. Currently, must be one of "🎲", "🎯", "🏀", "⚽", "🎳", or "🎰". Dice can have values 1-6 for "🎲", "🎯" and "🎳", values 1-5 for "🏀" and "⚽", and values 1-64 for "🎰". Defaults to "🎲"
	DisableNotification  bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             // Protects the contents of the sent message from forwarding
	AllowPaidBroadcast   bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      int              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters // Description of the message to reply to
	ReplyMarkup          Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendDiceRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	w.WriteField("emoji", r.Emoji)
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SendChatAction(ctx, &SendChatActionRequest{})
type SendChatActionRequest struct {
	BusinessConnectionID string     // Unique identifier of the business connection on behalf of which the action will be sent
	ChatID               ChatID     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int        // Unique identifier for the target message thread; for supergroups only
	Action               ChatAction // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers, find_location for location data, record_video_note or upload_video_note for video notes.
}

func (r *SendChatActionRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	w.WriteField("action", string(r.Action))
}

// use Bot.SetMessageReaction(ctx, &SetMessageReactionRequest{})
type SetMessageReactionRequest struct {
	ChatID    ChatID         // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID int            // Identifier of the target message. If the message belongs to a media group, the reaction is set to the first non-deleted message in the group instead.
	Reaction  []ReactionType // A JSON-serialized list of reaction types to set on the message. Currently, as non-premium users, bots can set up to one reaction per message. A custom emoji reaction can be used if it is either already present on the message or explicitly allowed by chat administrators. Paid reactions can't be used by bots.
	IsBig     bool           // Pass True to set the reaction with a big animation
}

func (r *SetMessageReactionRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
	if r.Reaction != nil {
		reaction, _ := json.Marshal(r.Reaction)
		w.WriteField("reaction", string(reaction))
	}
	is_big, _ := json.Marshal(r.IsBig)
	w.WriteField("is_big", string(is_big))
}

// use Bot.GetUserProfilePhotos(ctx, &GetUserProfilePhotosRequest{})
type GetUserProfilePhotosRequest struct {
	UserID int64 // Unique identifier of the target user
	Offset int64 // Sequential number of the first photo to be returned. By default, all photos are returned.
	Limit  int64 // Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.
}

func (r *GetUserProfilePhotosRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
	offset, _ := json.Marshal(r.Offset)
	w.WriteField("offset", string(offset))
	limit, _ := json.Marshal(r.Limit)
	w.WriteField("limit", string(limit))
}

// use Bot.SetUserEmojiStatus(ctx, &SetUserEmojiStatusRequest{})
type SetUserEmojiStatusRequest struct {
	UserID                    int64  // Unique identifier of the target user
	EmojiStatusCustomEmojiID  string // Custom emoji identifier of the emoji status to set. Pass an empty string to remove the status.
	EmojiStatusExpirationDate int64  // Expiration date of the emoji status, if any
}

func (r *SetUserEmojiStatusRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))

	w.WriteField("emoji_status_custom_emoji_id", r.EmojiStatusCustomEmojiID)
	emoji_status_expiration_date, _ := json.Marshal(r.EmojiStatusExpirationDate)
	w.WriteField("emoji_status_expiration_date", string(emoji_status_expiration_date))
}

// use Bot.GetFile(ctx, &GetFileRequest{})
type GetFileRequest struct {
	FileID string // File identifier to get information about
}

func (r *GetFileRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("file_id", r.FileID)
}

// use Bot.BanChatMember(ctx, &BanChatMemberRequest{})
type BanChatMemberRequest struct {
	ChatID         ChatID // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserID         int64  // Unique identifier of the target user
	UntilDate      int64  // Date when the user will be unbanned; Unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever. Applied for supergroups and channels only.
	RevokeMessages bool   // Pass True to delete all messages from the chat for the user that is being removed. If False, the user will be able to see messages in the group that were sent before the user was removed. Always True for supergroups and channels.
}

func (r *BanChatMemberRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
	until_date, _ := json.Marshal(r.UntilDate)
	w.WriteField("until_date", string(until_date))
	revoke_messages, _ := json.Marshal(r.RevokeMessages)
	w.WriteField("revoke_messages", string(revoke_messages))
}

// use Bot.UnbanChatMember(ctx, &UnbanChatMemberRequest{})
type UnbanChatMemberRequest struct {
	ChatID       ChatID // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserID       int64  // Unique identifier of the target user
	OnlyIfBanned bool   // Do nothing if the user is not banned
}

func (r *UnbanChatMemberRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
	only_if_banned, _ := json.Marshal(r.OnlyIfBanned)
	w.WriteField("only_if_banned", string(only_if_banned))
}

// use Bot.RestrictChatMember(ctx, &RestrictChatMemberRequest{})
type RestrictChatMemberRequest struct {
	ChatID                        ChatID           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserID                        int64            // Unique identifier of the target user
	Permissions                   *ChatPermissions // A JSON-serialized object for new user permissions
	UseIndependentChatPermissions bool             // Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
	UntilDate                     int64            // Date when restrictions will be lifted for the user; Unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
}

func (r *RestrictChatMemberRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
	if r.Permissions != nil {
		permissions, _ := json.Marshal(r.Permissions)
		w.WriteField("permissions", string(permissions))
	}
	use_independent_chat_permissions, _ := json.Marshal(r.UseIndependentChatPermissions)
	w.WriteField("use_independent_chat_permissions", string(use_independent_chat_permissions))
	until_date, _ := json.Marshal(r.UntilDate)
	w.WriteField("until_date", string(until_date))
}

// use Bot.PromoteChatMember(ctx, &PromoteChatMemberRequest{})
type PromoteChatMemberRequest struct {
	ChatID              ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserID              int64  // Unique identifier of the target user
	IsAnonymous         bool   // Pass True if the administrator's presence in the chat is hidden
	CanManageChat       bool   // Pass True if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege.
	CanDeleteMessages   bool   // Pass True if the administrator can delete messages of other users
	CanManageVideoChats bool   // Pass True if the administrator can manage video chats
	CanRestrictMembers  bool   // Pass True if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanPromoteMembers   bool   // Pass True if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him)
	CanChangeInfo       bool   // Pass True if the administrator can change chat title, photo and other settings
	CanInviteUsers      bool   // Pass True if the administrator can invite new users to the chat
	CanPostStories      bool   // Pass True if the administrator can post stories to the chat
	CanEditStories      bool   // Pass True if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanDeleteStories    bool   // Pass True if the administrator can delete stories posted by other users
	CanPostMessages     bool   // Pass True if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanEditMessages     bool   // Pass True if the administrator can edit messages of other users and can pin messages; for channels only
	CanPinMessages      bool   // Pass True if the administrator can pin messages; for supergroups only
	CanManageTopics     bool   // Pass True if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
}

func (r *PromoteChatMemberRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
	is_anonymous, _ := json.Marshal(r.IsAnonymous)
	w.WriteField("is_anonymous", string(is_anonymous))
	can_manage_chat, _ := json.Marshal(r.CanManageChat)
	w.WriteField("can_manage_chat", string(can_manage_chat))
	can_delete_messages, _ := json.Marshal(r.CanDeleteMessages)
	w.WriteField("can_delete_messages", string(can_delete_messages))
	can_manage_video_chats, _ := json.Marshal(r.CanManageVideoChats)
	w.WriteField("can_manage_video_chats", string(can_manage_video_chats))
	can_restrict_members, _ := json.Marshal(r.CanRestrictMembers)
	w.WriteField("can_restrict_members", string(can_restrict_members))
	can_promote_members, _ := json.Marshal(r.CanPromoteMembers)
	w.WriteField("can_promote_members", string(can_promote_members))
	can_change_info, _ := json.Marshal(r.CanChangeInfo)
	w.WriteField("can_change_info", string(can_change_info))
	can_invite_users, _ := json.Marshal(r.CanInviteUsers)
	w.WriteField("can_invite_users", string(can_invite_users))
	can_post_stories, _ := json.Marshal(r.CanPostStories)
	w.WriteField("can_post_stories", string(can_post_stories))
	can_edit_stories, _ := json.Marshal(r.CanEditStories)
	w.WriteField("can_edit_stories", string(can_edit_stories))
	can_delete_stories, _ := json.Marshal(r.CanDeleteStories)
	w.WriteField("can_delete_stories", string(can_delete_stories))
	can_post_messages, _ := json.Marshal(r.CanPostMessages)
	w.WriteField("can_post_messages", string(can_post_messages))
	can_edit_messages, _ := json.Marshal(r.CanEditMessages)
	w.WriteField("can_edit_messages", string(can_edit_messages))
	can_pin_messages, _ := json.Marshal(r.CanPinMessages)
	w.WriteField("can_pin_messages", string(can_pin_messages))
	can_manage_topics, _ := json.Marshal(r.CanManageTopics)
	w.WriteField("can_manage_topics", string(can_manage_topics))
}

// use Bot.SetChatAdministratorCustomTitle(ctx, &SetChatAdministratorCustomTitleRequest{})
type SetChatAdministratorCustomTitleRequest struct {
	ChatID      ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserID      int64  // Unique identifier of the target user
	CustomTitle string // New custom title for the administrator; 0-16 characters, emoji are not allowed
}

func (r *SetChatAdministratorCustomTitleRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))

	w.WriteField("custom_title", r.CustomTitle)
}

// use Bot.BanChatSenderChat(ctx, &BanChatSenderChatRequest{})
type BanChatSenderChatRequest struct {
	ChatID       ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	SenderChatID int64  // Unique identifier of the target sender chat
}

func (r *BanChatSenderChatRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	sender_chat_id, _ := json.Marshal(r.SenderChatID)
	w.WriteField("sender_chat_id", string(sender_chat_id))
}

// use Bot.UnbanChatSenderChat(ctx, &UnbanChatSenderChatRequest{})
type UnbanChatSenderChatRequest struct {
	ChatID       ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	SenderChatID int64  // Unique identifier of the target sender chat
}

func (r *UnbanChatSenderChatRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	sender_chat_id, _ := json.Marshal(r.SenderChatID)
	w.WriteField("sender_chat_id", string(sender_chat_id))
}

// use Bot.SetChatPermissions(ctx, &SetChatPermissionsRequest{})
type SetChatPermissionsRequest struct {
	ChatID                        ChatID           // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Permissions                   *ChatPermissions // A JSON-serialized object for new default chat permissions
	UseIndependentChatPermissions bool             // Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
}

func (r *SetChatPermissionsRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	if r.Permissions != nil {
		permissions, _ := json.Marshal(r.Permissions)
		w.WriteField("permissions", string(permissions))
	}
	use_independent_chat_permissions, _ := json.Marshal(r.UseIndependentChatPermissions)
	w.WriteField("use_independent_chat_permissions", string(use_independent_chat_permissions))
}

// use Bot.ExportChatInviteLink(ctx, &ExportChatInviteLinkRequest{})
type ExportChatInviteLinkRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

func (r *ExportChatInviteLinkRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.CreateChatInviteLink(ctx, &CreateChatInviteLinkRequest{})
type CreateChatInviteLinkRequest struct {
	ChatID             ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Name               string // Invite link name; 0-32 characters
	ExpireDate         int64  // Point in time (Unix timestamp) when the link will expire
	MemberLimit        int64  // The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	CreatesJoinRequest bool   // True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
}

func (r *CreateChatInviteLinkRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("name", r.Name)
	expire_date, _ := json.Marshal(r.ExpireDate)
	w.WriteField("expire_date", string(expire_date))
	member_limit, _ := json.Marshal(r.MemberLimit)
	w.WriteField("member_limit", string(member_limit))
	creates_join_request, _ := json.Marshal(r.CreatesJoinRequest)
	w.WriteField("creates_join_request", string(creates_join_request))
}

// use Bot.EditChatInviteLink(ctx, &EditChatInviteLinkRequest{})
type EditChatInviteLinkRequest struct {
	ChatID             ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	InviteLink         string // The invite link to edit
	Name               string // Invite link name; 0-32 characters
	ExpireDate         int64  // Point in time (Unix timestamp) when the link will expire
	MemberLimit        int64  // The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	CreatesJoinRequest bool   // True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
}

func (r *EditChatInviteLinkRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("invite_link", r.InviteLink)

	w.WriteField("name", r.Name)
	expire_date, _ := json.Marshal(r.ExpireDate)
	w.WriteField("expire_date", string(expire_date))
	member_limit, _ := json.Marshal(r.MemberLimit)
	w.WriteField("member_limit", string(member_limit))
	creates_join_request, _ := json.Marshal(r.CreatesJoinRequest)
	w.WriteField("creates_join_request", string(creates_join_request))
}

// use Bot.CreateChatSubscriptionInviteLink(ctx, &CreateChatSubscriptionInviteLinkRequest{})
type CreateChatSubscriptionInviteLinkRequest struct {
	ChatID             ChatID // Unique identifier for the target channel chat or username of the target channel (in the format @channelusername)
	Name               string // Invite link name; 0-32 characters
	SubscriptionPeriod int64  // The number of seconds the subscription will be active for before the next payment. Currently, it must always be 2592000 (30 days).
	SubscriptionPrice  int64  // The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat; 1-2500
}

func (r *CreateChatSubscriptionInviteLinkRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("name", r.Name)
	subscription_period, _ := json.Marshal(r.SubscriptionPeriod)
	w.WriteField("subscription_period", string(subscription_period))
	subscription_price, _ := json.Marshal(r.SubscriptionPrice)
	w.WriteField("subscription_price", string(subscription_price))
}

// use Bot.EditChatSubscriptionInviteLink(ctx, &EditChatSubscriptionInviteLinkRequest{})
type EditChatSubscriptionInviteLinkRequest struct {
	ChatID     ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	InviteLink string // The invite link to edit
	Name       string // Invite link name; 0-32 characters
}

func (r *EditChatSubscriptionInviteLinkRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("invite_link", r.InviteLink)

	w.WriteField("name", r.Name)
}

// use Bot.RevokeChatInviteLink(ctx, &RevokeChatInviteLinkRequest{})
type RevokeChatInviteLinkRequest struct {
	ChatID     ChatID // Unique identifier of the target chat or username of the target channel (in the format @channelusername)
	InviteLink string // The invite link to revoke
}

func (r *RevokeChatInviteLinkRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("invite_link", r.InviteLink)
}

// use Bot.ApproveChatJoinRequest(ctx, &ApproveChatJoinRequestRequest{})
type ApproveChatJoinRequestRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserID int64  // Unique identifier of the target user
}

func (r *ApproveChatJoinRequestRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
}

// use Bot.DeclineChatJoinRequest(ctx, &DeclineChatJoinRequestRequest{})
type DeclineChatJoinRequestRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserID int64  // Unique identifier of the target user
}

func (r *DeclineChatJoinRequestRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
}

// use Bot.SetChatPhoto(ctx, &SetChatPhotoRequest{})
type SetChatPhotoRequest struct {
	ChatID ChatID    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo  InputFile // New chat photo, uploaded using multipart/form-data
}

func (r *SetChatPhotoRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	if s, ok := r.Photo.(string); ok {
		w.WriteField("photo", s)
	} else {
		// fw, _ := w.CreateFormFile("photo", "todo.jpeg")
	}
}

// use Bot.DeleteChatPhoto(ctx, &DeleteChatPhotoRequest{})
type DeleteChatPhotoRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

func (r *DeleteChatPhotoRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.SetChatTitle(ctx, &SetChatTitleRequest{})
type SetChatTitleRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Title  string // New chat title, 1-128 characters
}

func (r *SetChatTitleRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("title", r.Title)
}

// use Bot.SetChatDescription(ctx, &SetChatDescriptionRequest{})
type SetChatDescriptionRequest struct {
	ChatID      ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Description string // New chat description, 0-255 characters
}

func (r *SetChatDescriptionRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("description", r.Description)
}

// use Bot.PinChatMessage(ctx, &PinChatMessageRequest{})
type PinChatMessageRequest struct {
	BusinessConnectionID string // Unique identifier of the business connection on behalf of which the message will be pinned
	ChatID               ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int    // Identifier of a message to pin
	DisableNotification  bool   // Pass True if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels and private chats.
}

func (r *PinChatMessageRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
}

// use Bot.UnpinChatMessage(ctx, &UnpinChatMessageRequest{})
type UnpinChatMessageRequest struct {
	BusinessConnectionID string // Unique identifier of the business connection on behalf of which the message will be unpinned
	ChatID               ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int    // Identifier of the message to unpin. Required if business_connection_id is specified. If not specified, the most recent pinned message (by sending date) will be unpinned.
}

func (r *UnpinChatMessageRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
}

// use Bot.UnpinAllChatMessages(ctx, &UnpinAllChatMessagesRequest{})
type UnpinAllChatMessagesRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

func (r *UnpinAllChatMessagesRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.LeaveChat(ctx, &LeaveChatRequest{})
type LeaveChatRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (r *LeaveChatRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.GetChat(ctx, &GetChatRequest{})
type GetChatRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (r *GetChatRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.GetChatAdministrators(ctx, &GetChatAdministratorsRequest{})
type GetChatAdministratorsRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (r *GetChatAdministratorsRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.GetChatMemberCount(ctx, &GetChatMemberCountRequest{})
type GetChatMemberCountRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (r *GetChatMemberCountRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.GetChatMember(ctx, &GetChatMemberRequest{})
type GetChatMemberRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	UserID int64  // Unique identifier of the target user
}

func (r *GetChatMemberRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
}

// use Bot.SetChatStickerSet(ctx, &SetChatStickerSetRequest{})
type SetChatStickerSetRequest struct {
	ChatID         ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	StickerSetName string // Name of the sticker set to be set as the group sticker set
}

func (r *SetChatStickerSetRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("sticker_set_name", r.StickerSetName)
}

// use Bot.DeleteChatStickerSet(ctx, &DeleteChatStickerSetRequest{})
type DeleteChatStickerSetRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (r *DeleteChatStickerSetRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.CreateForumTopic(ctx, &CreateForumTopicRequest{})
type CreateForumTopicRequest struct {
	ChatID            ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Name              string // Topic name, 1-128 characters
	IconColor         int64  // Color of the topic icon in RGB format. Currently, must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F)
	IconCustomEmojiID string // Unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers.
}

func (r *CreateForumTopicRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("name", r.Name)
	icon_color, _ := json.Marshal(r.IconColor)
	w.WriteField("icon_color", string(icon_color))

	w.WriteField("icon_custom_emoji_id", r.IconCustomEmojiID)
}

// use Bot.EditForumTopic(ctx, &EditForumTopicRequest{})
type EditForumTopicRequest struct {
	ChatID            ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadID   int    // Unique identifier for the target message thread of the forum topic
	Name              string // New topic name, 0-128 characters. If not specified or empty, the current name of the topic will be kept
	IconCustomEmojiID string // New unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers. Pass an empty string to remove the icon. If not specified, the current icon will be kept
}

func (r *EditForumTopicRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	w.WriteField("name", r.Name)

	w.WriteField("icon_custom_emoji_id", r.IconCustomEmojiID)
}

// use Bot.CloseForumTopic(ctx, &CloseForumTopicRequest{})
type CloseForumTopicRequest struct {
	ChatID          ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadID int    // Unique identifier for the target message thread of the forum topic
}

func (r *CloseForumTopicRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))
}

// use Bot.ReopenForumTopic(ctx, &ReopenForumTopicRequest{})
type ReopenForumTopicRequest struct {
	ChatID          ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadID int    // Unique identifier for the target message thread of the forum topic
}

func (r *ReopenForumTopicRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))
}

// use Bot.DeleteForumTopic(ctx, &DeleteForumTopicRequest{})
type DeleteForumTopicRequest struct {
	ChatID          ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadID int    // Unique identifier for the target message thread of the forum topic
}

func (r *DeleteForumTopicRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))
}

// use Bot.UnpinAllForumTopicMessages(ctx, &UnpinAllForumTopicMessagesRequest{})
type UnpinAllForumTopicMessagesRequest struct {
	ChatID          ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	MessageThreadID int    // Unique identifier for the target message thread of the forum topic
}

func (r *UnpinAllForumTopicMessagesRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))
}

// use Bot.EditGeneralForumTopic(ctx, &EditGeneralForumTopicRequest{})
type EditGeneralForumTopicRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Name   string // New topic name, 1-128 characters
}

func (r *EditGeneralForumTopicRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("name", r.Name)
}

// use Bot.CloseGeneralForumTopic(ctx, &CloseGeneralForumTopicRequest{})
type CloseGeneralForumTopicRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (r *CloseGeneralForumTopicRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.ReopenGeneralForumTopic(ctx, &ReopenGeneralForumTopicRequest{})
type ReopenGeneralForumTopicRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (r *ReopenGeneralForumTopicRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.HideGeneralForumTopic(ctx, &HideGeneralForumTopicRequest{})
type HideGeneralForumTopicRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (r *HideGeneralForumTopicRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.UnhideGeneralForumTopic(ctx, &UnhideGeneralForumTopicRequest{})
type UnhideGeneralForumTopicRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (r *UnhideGeneralForumTopicRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.UnpinAllGeneralForumTopicMessages(ctx, &UnpinAllGeneralForumTopicMessagesRequest{})
type UnpinAllGeneralForumTopicMessagesRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (r *UnpinAllGeneralForumTopicMessagesRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.AnswerCallbackQuery(ctx, &AnswerCallbackQueryRequest{})
type AnswerCallbackQueryRequest struct {
	CallbackQueryID string // Unique identifier for the query to be answered
	Text            string // Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	ShowAlert       bool   // If True, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	Url             string // URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @BotFather, specify the URL that opens your game - note that this will only work if the query comes from a callback_game button. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	CacheTime       int64  // The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
}

func (r *AnswerCallbackQueryRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("callback_query_id", r.CallbackQueryID)

	w.WriteField("text", r.Text)
	show_alert, _ := json.Marshal(r.ShowAlert)
	w.WriteField("show_alert", string(show_alert))

	w.WriteField("url", r.Url)
	cache_time, _ := json.Marshal(r.CacheTime)
	w.WriteField("cache_time", string(cache_time))
}

// use Bot.GetUserChatBoosts(ctx, &GetUserChatBoostsRequest{})
type GetUserChatBoostsRequest struct {
	ChatID ChatID // Unique identifier for the chat or username of the channel (in the format @channelusername)
	UserID int64  // Unique identifier of the target user
}

func (r *GetUserChatBoostsRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
}

// use Bot.GetBusinessConnection(ctx, &GetBusinessConnectionRequest{})
type GetBusinessConnectionRequest struct {
	BusinessConnectionID string // Unique identifier of the business connection
}

func (r *GetBusinessConnectionRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)
}

// use Bot.SetMyCommands(ctx, &SetMyCommandsRequest{})
type SetMyCommandsRequest struct {
	Commands     []BotCommand    // A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	Scope        BotCommandScope // A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string          // A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

func (r *SetMyCommandsRequest) WriteMultipart(w *multipart.Writer) {
	if r.Commands != nil {
		commands, _ := json.Marshal(r.Commands)
		w.WriteField("commands", string(commands))
	}
	if r.Scope != nil {
		scope, _ := json.Marshal(r.Scope)
		w.WriteField("scope", string(scope))
	}

	w.WriteField("language_code", r.LanguageCode)
}

// use Bot.DeleteMyCommands(ctx, &DeleteMyCommandsRequest{})
type DeleteMyCommandsRequest struct {
	Scope        BotCommandScope // A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string          // A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

func (r *DeleteMyCommandsRequest) WriteMultipart(w *multipart.Writer) {
	if r.Scope != nil {
		scope, _ := json.Marshal(r.Scope)
		w.WriteField("scope", string(scope))
	}

	w.WriteField("language_code", r.LanguageCode)
}

// use Bot.GetMyCommands(ctx, &GetMyCommandsRequest{})
type GetMyCommandsRequest struct {
	Scope        BotCommandScope // A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
	LanguageCode string          // A two-letter ISO 639-1 language code or an empty string
}

func (r *GetMyCommandsRequest) WriteMultipart(w *multipart.Writer) {
	if r.Scope != nil {
		scope, _ := json.Marshal(r.Scope)
		w.WriteField("scope", string(scope))
	}

	w.WriteField("language_code", r.LanguageCode)
}

// use Bot.SetMyName(ctx, &SetMyNameRequest{})
type SetMyNameRequest struct {
	Name         string // New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given language.
	LanguageCode string // A two-letter ISO 639-1 language code. If empty, the name will be shown to all users for whose language there is no dedicated name.
}

func (r *SetMyNameRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("name", r.Name)

	w.WriteField("language_code", r.LanguageCode)
}

// use Bot.GetMyName(ctx, &GetMyNameRequest{})
type GetMyNameRequest struct {
	LanguageCode string // A two-letter ISO 639-1 language code or an empty string
}

func (r *GetMyNameRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("language_code", r.LanguageCode)
}

// use Bot.SetMyDescription(ctx, &SetMyDescriptionRequest{})
type SetMyDescriptionRequest struct {
	Description  string // New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language.
	LanguageCode string // A two-letter ISO 639-1 language code. If empty, the description will be applied to all users for whose language there is no dedicated description.
}

func (r *SetMyDescriptionRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("description", r.Description)

	w.WriteField("language_code", r.LanguageCode)
}

// use Bot.GetMyDescription(ctx, &GetMyDescriptionRequest{})
type GetMyDescriptionRequest struct {
	LanguageCode string // A two-letter ISO 639-1 language code or an empty string
}

func (r *GetMyDescriptionRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("language_code", r.LanguageCode)
}

// use Bot.SetMyShortDescription(ctx, &SetMyShortDescriptionRequest{})
type SetMyShortDescriptionRequest struct {
	ShortDescription string // New short description for the bot; 0-120 characters. Pass an empty string to remove the dedicated short description for the given language.
	LanguageCode     string // A two-letter ISO 639-1 language code. If empty, the short description will be applied to all users for whose language there is no dedicated short description.
}

func (r *SetMyShortDescriptionRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("short_description", r.ShortDescription)

	w.WriteField("language_code", r.LanguageCode)
}

// use Bot.GetMyShortDescription(ctx, &GetMyShortDescriptionRequest{})
type GetMyShortDescriptionRequest struct {
	LanguageCode string // A two-letter ISO 639-1 language code or an empty string
}

func (r *GetMyShortDescriptionRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("language_code", r.LanguageCode)
}

// use Bot.SetChatMenuButton(ctx, &SetChatMenuButtonRequest{})
type SetChatMenuButtonRequest struct {
	ChatID     int64      // Unique identifier for the target private chat. If not specified, default bot's menu button will be changed
	MenuButton MenuButton // A JSON-serialized object for the bot's new menu button. Defaults to MenuButtonDefault
}

func (r *SetChatMenuButtonRequest) WriteMultipart(w *multipart.Writer) {
	chat_id, _ := json.Marshal(r.ChatID)
	w.WriteField("chat_id", string(chat_id))
	if r.MenuButton != nil {
		menu_button, _ := json.Marshal(r.MenuButton)
		w.WriteField("menu_button", string(menu_button))
	}
}

// use Bot.GetChatMenuButton(ctx, &GetChatMenuButtonRequest{})
type GetChatMenuButtonRequest struct {
	ChatID int64 // Unique identifier for the target private chat. If not specified, default bot's menu button will be returned
}

func (r *GetChatMenuButtonRequest) WriteMultipart(w *multipart.Writer) {
	chat_id, _ := json.Marshal(r.ChatID)
	w.WriteField("chat_id", string(chat_id))
}

// use Bot.SetMyDefaultAdministratorRights(ctx, &SetMyDefaultAdministratorRightsRequest{})
type SetMyDefaultAdministratorRightsRequest struct {
	Rights      *ChatAdministratorRights // A JSON-serialized object describing new default administrator rights. If not specified, the default administrator rights will be cleared.
	ForChannels bool                     // Pass True to change the default administrator rights of the bot in channels. Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
}

func (r *SetMyDefaultAdministratorRightsRequest) WriteMultipart(w *multipart.Writer) {
	if r.Rights != nil {
		rights, _ := json.Marshal(r.Rights)
		w.WriteField("rights", string(rights))
	}
	for_channels, _ := json.Marshal(r.ForChannels)
	w.WriteField("for_channels", string(for_channels))
}

// use Bot.GetMyDefaultAdministratorRights(ctx, &GetMyDefaultAdministratorRightsRequest{})
type GetMyDefaultAdministratorRightsRequest struct {
	ForChannels bool // Pass True to get default administrator rights of the bot in channels. Otherwise, default administrator rights of the bot for groups and supergroups will be returned.
}

func (r *GetMyDefaultAdministratorRightsRequest) WriteMultipart(w *multipart.Writer) {
	for_channels, _ := json.Marshal(r.ForChannels)
	w.WriteField("for_channels", string(for_channels))
}

// use Bot.EditMessageText(ctx, &EditMessageTextRequest{})
type EditMessageTextRequest struct {
	BusinessConnectionID string                // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID               ChatID                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int                   // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID      int                   // Required if chat_id and message_id are not specified. Identifier of the inline message
	Text                 string                // New text of the message, 1-4096 characters after entities parsing
	ParseMode            ParseMode             // Mode for parsing entities in the message text. See formatting options for more details.
	Entities             []MessageEntity       // A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	LinkPreviewOptions   *LinkPreviewOptions   // Link preview generation options for the message
	ReplyMarkup          *InlineKeyboardMarkup // A JSON-serialized object for an inline keyboard.
}

func (r *EditMessageTextRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
	inline_message_id, _ := json.Marshal(r.InlineMessageID)
	w.WriteField("inline_message_id", string(inline_message_id))

	w.WriteField("text", r.Text)

	w.WriteField("parse_mode", string(r.ParseMode))
	if r.Entities != nil {
		entities, _ := json.Marshal(r.Entities)
		w.WriteField("entities", string(entities))
	}
	if r.LinkPreviewOptions != nil {
		link_preview_options, _ := json.Marshal(r.LinkPreviewOptions)
		w.WriteField("link_preview_options", string(link_preview_options))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.EditMessageCaption(ctx, &EditMessageCaptionRequest{})
type EditMessageCaptionRequest struct {
	BusinessConnectionID  string                // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID                ChatID                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID             int                   // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID       int                   // Required if chat_id and message_id are not specified. Identifier of the inline message
	Caption               string                // New caption of the message, 0-1024 characters after entities parsing
	ParseMode             ParseMode             // Mode for parsing entities in the message caption. See formatting options for more details.
	CaptionEntities       []MessageEntity       // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  // Pass True, if the caption must be shown above the message media. Supported only for animation, photo and video messages.
	ReplyMarkup           *InlineKeyboardMarkup // A JSON-serialized object for an inline keyboard.
}

func (r *EditMessageCaptionRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
	inline_message_id, _ := json.Marshal(r.InlineMessageID)
	w.WriteField("inline_message_id", string(inline_message_id))

	w.WriteField("caption", r.Caption)

	w.WriteField("parse_mode", string(r.ParseMode))
	if r.CaptionEntities != nil {
		caption_entities, _ := json.Marshal(r.CaptionEntities)
		w.WriteField("caption_entities", string(caption_entities))
	}
	show_caption_above_media, _ := json.Marshal(r.ShowCaptionAboveMedia)
	w.WriteField("show_caption_above_media", string(show_caption_above_media))
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.EditMessageMedia(ctx, &EditMessageMediaRequest{})
type EditMessageMediaRequest struct {
	BusinessConnectionID string                // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID               ChatID                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int                   // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID      int                   // Required if chat_id and message_id are not specified. Identifier of the inline message
	Media                InputMedia            // A JSON-serialized object for a new media content of the message
	ReplyMarkup          *InlineKeyboardMarkup // A JSON-serialized object for a new inline keyboard.
}

func (r *EditMessageMediaRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
	inline_message_id, _ := json.Marshal(r.InlineMessageID)
	w.WriteField("inline_message_id", string(inline_message_id))
	if r.Media != nil {
		media, _ := json.Marshal(r.Media)
		w.WriteField("media", string(media))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.EditMessageLiveLocation(ctx, &EditMessageLiveLocationRequest{})
type EditMessageLiveLocationRequest struct {
	BusinessConnectionID string                // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID               ChatID                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int                   // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID      int                   // Required if chat_id and message_id are not specified. Identifier of the inline message
	Latitude             float64               // Latitude of new location
	Longitude            float64               // Longitude of new location
	LivePeriod           int64                 // New period in seconds during which the location can be updated, starting from the message send date. If 0x7FFFFFFF is specified, then the location can be updated forever. Otherwise, the new value must not exceed the current live_period by more than a day, and the live location expiration date must remain within the next 90 days. If not specified, then live_period remains unchanged
	HorizontalAccuracy   float64               // The radius of uncertainty for the location, measured in meters; 0-1500
	Heading              int64                 // Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int64                 // The maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ReplyMarkup          *InlineKeyboardMarkup // A JSON-serialized object for a new inline keyboard.
}

func (r *EditMessageLiveLocationRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
	inline_message_id, _ := json.Marshal(r.InlineMessageID)
	w.WriteField("inline_message_id", string(inline_message_id))
	latitude, _ := json.Marshal(r.Latitude)
	w.WriteField("latitude", string(latitude))
	longitude, _ := json.Marshal(r.Longitude)
	w.WriteField("longitude", string(longitude))
	live_period, _ := json.Marshal(r.LivePeriod)
	w.WriteField("live_period", string(live_period))
	horizontal_accuracy, _ := json.Marshal(r.HorizontalAccuracy)
	w.WriteField("horizontal_accuracy", string(horizontal_accuracy))
	heading, _ := json.Marshal(r.Heading)
	w.WriteField("heading", string(heading))
	proximity_alert_radius, _ := json.Marshal(r.ProximityAlertRadius)
	w.WriteField("proximity_alert_radius", string(proximity_alert_radius))
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.StopMessageLiveLocation(ctx, &StopMessageLiveLocationRequest{})
type StopMessageLiveLocationRequest struct {
	BusinessConnectionID string                // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID               ChatID                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int                   // Required if inline_message_id is not specified. Identifier of the message with live location to stop
	InlineMessageID      int                   // Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup          *InlineKeyboardMarkup // A JSON-serialized object for a new inline keyboard.
}

func (r *StopMessageLiveLocationRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
	inline_message_id, _ := json.Marshal(r.InlineMessageID)
	w.WriteField("inline_message_id", string(inline_message_id))
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.EditMessageReplyMarkup(ctx, &EditMessageReplyMarkupRequest{})
type EditMessageReplyMarkupRequest struct {
	BusinessConnectionID string                // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID               ChatID                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int                   // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageID      int                   // Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup          *InlineKeyboardMarkup // A JSON-serialized object for an inline keyboard.
}

func (r *EditMessageReplyMarkupRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
	inline_message_id, _ := json.Marshal(r.InlineMessageID)
	w.WriteField("inline_message_id", string(inline_message_id))
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.StopPoll(ctx, &StopPollRequest{})
type StopPollRequest struct {
	BusinessConnectionID string                // Unique identifier of the business connection on behalf of which the message to be edited was sent
	ChatID               ChatID                // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID            int                   // Identifier of the original message with the poll
	ReplyMarkup          *InlineKeyboardMarkup // A JSON-serialized object for a new message inline keyboard.
}

func (r *StopPollRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.DeleteMessage(ctx, &DeleteMessageRequest{})
type DeleteMessageRequest struct {
	ChatID    ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID int    // Identifier of the message to delete
}

func (r *DeleteMessageRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
}

// use Bot.DeleteMessages(ctx, &DeleteMessagesRequest{})
type DeleteMessagesRequest struct {
	ChatID     ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageIDs int    // A JSON-serialized list of 1-100 identifiers of messages to delete. See deleteMessage for limitations on which messages can be deleted
}

func (r *DeleteMessagesRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	message_ids, _ := json.Marshal(r.MessageIDs)
	w.WriteField("message_ids", string(message_ids))
}

// use Bot.SendSticker(ctx, &SendStickerRequest{})
type SendStickerRequest struct {
	BusinessConnectionID string           // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               ChatID           // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID      int              // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Sticker              InputFile        // Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP sticker from the Internet, or upload a new .WEBP, .TGS, or .WEBM sticker using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files. Video and animated stickers can't be sent via an HTTP URL.
	Emoji                string           // Emoji associated with the sticker; only for just uploaded stickers
	DisableNotification  bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool             // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool             // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      int              // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters // Description of the message to reply to
	ReplyMarkup          Markup           // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

func (r *SendStickerRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	if s, ok := r.Sticker.(string); ok {
		w.WriteField("sticker", s)
	} else {
		// fw, _ := w.CreateFormFile("sticker", "todo.jpeg")
	}

	w.WriteField("emoji", r.Emoji)
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.GetStickerSet(ctx, &GetStickerSetRequest{})
type GetStickerSetRequest struct {
	Name string // Name of the sticker set
}

func (r *GetStickerSetRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("name", r.Name)
}

// use Bot.GetCustomEmojiStickers(ctx, &GetCustomEmojiStickersRequest{})
type GetCustomEmojiStickersRequest struct {
	CustomEmojiIds []string // A JSON-serialized list of custom emoji identifiers. At most 200 custom emoji identifiers can be specified.
}

func (r *GetCustomEmojiStickersRequest) WriteMultipart(w *multipart.Writer) {
	if r.CustomEmojiIds != nil {
		custom_emoji_ids, _ := json.Marshal(r.CustomEmojiIds)
		w.WriteField("custom_emoji_ids", string(custom_emoji_ids))
	}
}

// use Bot.UploadStickerFile(ctx, &UploadStickerFileRequest{})
type UploadStickerFileRequest struct {
	UserID        int64     // User identifier of sticker file owner
	Sticker       InputFile // A file with the sticker in .WEBP, .PNG, .TGS, or .WEBM format. See https://core.telegram.org/stickers for technical requirements. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	StickerFormat string    // Format of the sticker, must be one of "static", "animated", "video"
}

func (r *UploadStickerFileRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))

	if s, ok := r.Sticker.(string); ok {
		w.WriteField("sticker", s)
	} else {
		// fw, _ := w.CreateFormFile("sticker", "todo.jpeg")
	}

	w.WriteField("sticker_format", r.StickerFormat)
}

// use Bot.CreateNewStickerSet(ctx, &CreateNewStickerSetRequest{})
type CreateNewStickerSetRequest struct {
	UserID          int64          // User identifier of created sticker set owner
	Name            string         // Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only English letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in "_by_<bot_username>". <bot_username> is case insensitive. 1-64 characters.
	Title           string         // Sticker set title, 1-64 characters
	Stickers        []InputSticker // A JSON-serialized list of 1-50 initial stickers to be added to the sticker set
	StickerType     string         // Type of stickers in the set, pass "regular", "mask", or "custom_emoji". By default, a regular sticker set is created.
	NeedsRepainting bool           // Pass True if stickers in the sticker set must be repainted to the color of text when used in messages, the accent color if used as emoji status, white on chat photos, or another appropriate color based on context; for custom emoji sticker sets only
}

func (r *CreateNewStickerSetRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))

	w.WriteField("name", r.Name)

	w.WriteField("title", r.Title)
	if r.Stickers != nil {
		stickers, _ := json.Marshal(r.Stickers)
		w.WriteField("stickers", string(stickers))
	}

	w.WriteField("sticker_type", r.StickerType)
	needs_repainting, _ := json.Marshal(r.NeedsRepainting)
	w.WriteField("needs_repainting", string(needs_repainting))
}

// use Bot.AddStickerToSet(ctx, &AddStickerToSetRequest{})
type AddStickerToSetRequest struct {
	UserID  int64         // User identifier of sticker set owner
	Name    string        // Sticker set name
	Sticker *InputSticker // A JSON-serialized object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set isn't changed.
}

func (r *AddStickerToSetRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))

	w.WriteField("name", r.Name)
	if r.Sticker != nil {
		sticker, _ := json.Marshal(r.Sticker)
		w.WriteField("sticker", string(sticker))
	}
}

// use Bot.SetStickerPositionInSet(ctx, &SetStickerPositionInSetRequest{})
type SetStickerPositionInSetRequest struct {
	Sticker  string // File identifier of the sticker
	Position int64  // New sticker position in the set, zero-based
}

func (r *SetStickerPositionInSetRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("sticker", r.Sticker)
	position, _ := json.Marshal(r.Position)
	w.WriteField("position", string(position))
}

// use Bot.DeleteStickerFromSet(ctx, &DeleteStickerFromSetRequest{})
type DeleteStickerFromSetRequest struct {
	Sticker string // File identifier of the sticker
}

func (r *DeleteStickerFromSetRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("sticker", r.Sticker)
}

// use Bot.ReplaceStickerInSet(ctx, &ReplaceStickerInSetRequest{})
type ReplaceStickerInSetRequest struct {
	UserID     int64         // User identifier of the sticker set owner
	Name       string        // Sticker set name
	OldSticker string        // File identifier of the replaced sticker
	Sticker    *InputSticker // A JSON-serialized object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set remains unchanged.
}

func (r *ReplaceStickerInSetRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))

	w.WriteField("name", r.Name)

	w.WriteField("old_sticker", r.OldSticker)
	if r.Sticker != nil {
		sticker, _ := json.Marshal(r.Sticker)
		w.WriteField("sticker", string(sticker))
	}
}

// use Bot.SetStickerEmojiList(ctx, &SetStickerEmojiListRequest{})
type SetStickerEmojiListRequest struct {
	Sticker   string   // File identifier of the sticker
	EmojiList []string // A JSON-serialized list of 1-20 emoji associated with the sticker
}

func (r *SetStickerEmojiListRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("sticker", r.Sticker)
	if r.EmojiList != nil {
		emoji_list, _ := json.Marshal(r.EmojiList)
		w.WriteField("emoji_list", string(emoji_list))
	}
}

// use Bot.SetStickerKeywords(ctx, &SetStickerKeywordsRequest{})
type SetStickerKeywordsRequest struct {
	Sticker  string   // File identifier of the sticker
	Keywords []string // A JSON-serialized list of 0-20 search keywords for the sticker with total length of up to 64 characters
}

func (r *SetStickerKeywordsRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("sticker", r.Sticker)
	if r.Keywords != nil {
		keywords, _ := json.Marshal(r.Keywords)
		w.WriteField("keywords", string(keywords))
	}
}

// use Bot.SetStickerMaskPosition(ctx, &SetStickerMaskPositionRequest{})
type SetStickerMaskPositionRequest struct {
	Sticker      string        // File identifier of the sticker
	MaskPosition *MaskPosition // A JSON-serialized object with the position where the mask should be placed on faces. Omit the parameter to remove the mask position.
}

func (r *SetStickerMaskPositionRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("sticker", r.Sticker)
	if r.MaskPosition != nil {
		mask_position, _ := json.Marshal(r.MaskPosition)
		w.WriteField("mask_position", string(mask_position))
	}
}

// use Bot.SetStickerSetTitle(ctx, &SetStickerSetTitleRequest{})
type SetStickerSetTitleRequest struct {
	Name  string // Sticker set name
	Title string // Sticker set title, 1-64 characters
}

func (r *SetStickerSetTitleRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("name", r.Name)

	w.WriteField("title", r.Title)
}

// use Bot.SetStickerSetThumbnail(ctx, &SetStickerSetThumbnailRequest{})
type SetStickerSetThumbnailRequest struct {
	Name      string    // Sticker set name
	UserID    int64     // User identifier of the sticker set owner
	Thumbnail InputFile // A .WEBP or .PNG image with the thumbnail, must be up to 128 kilobytes in size and have a width and height of exactly 100px, or a .TGS animation with a thumbnail up to 32 kilobytes in size (see https://core.telegram.org/stickers#animation-requirements for animated sticker technical requirements), or a .WEBM video with the thumbnail up to 32 kilobytes in size; see https://core.telegram.org/stickers#video-requirements for video sticker technical requirements. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files. Animated and video sticker set thumbnails can't be uploaded via HTTP URL. If omitted, then the thumbnail is dropped and the first sticker is used as the thumbnail.
	Format    string    // Format of the thumbnail, must be one of "static" for a .WEBP or .PNG image, "animated" for a .TGS animation, or "video" for a .WEBM video
}

func (r *SetStickerSetThumbnailRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("name", r.Name)
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))

	if s, ok := r.Thumbnail.(string); ok {
		w.WriteField("thumbnail", s)
	} else {
		// fw, _ := w.CreateFormFile("thumbnail", "todo.jpeg")
	}

	w.WriteField("format", r.Format)
}

// use Bot.SetCustomEmojiStickerSetThumbnail(ctx, &SetCustomEmojiStickerSetThumbnailRequest{})
type SetCustomEmojiStickerSetThumbnailRequest struct {
	Name          string // Sticker set name
	CustomEmojiID string // Custom emoji identifier of a sticker from the sticker set; pass an empty string to drop the thumbnail and use the first sticker as the thumbnail.
}

func (r *SetCustomEmojiStickerSetThumbnailRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("name", r.Name)

	w.WriteField("custom_emoji_id", r.CustomEmojiID)
}

// use Bot.DeleteStickerSet(ctx, &DeleteStickerSetRequest{})
type DeleteStickerSetRequest struct {
	Name string // Sticker set name
}

func (r *DeleteStickerSetRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("name", r.Name)
}

// use Bot.SendGift(ctx, &SendGiftRequest{})
type SendGiftRequest struct {
	UserID        int64           // Required if chat_id is not specified. Unique identifier of the target user who will receive the gift.
	ChatID        ChatID          // Required if user_id is not specified. Unique identifier for the chat or username of the channel (in the format @channelusername) that will receive the gift.
	GiftID        string          // Identifier of the gift
	PayForUpgrade bool            // Pass True to pay for the gift upgrade from the bot's balance, thereby making the upgrade free for the receiver
	Text          string          // Text that will be shown along with the gift; 0-128 characters
	TextParseMode string          // Mode for parsing entities in the text. See formatting options for more details. Entities other than "bold", "italic", "underline", "strikethrough", "spoiler", and "custom_emoji" are ignored.
	TextEntities  []MessageEntity // A JSON-serialized list of special entities that appear in the gift text. It can be specified instead of text_parse_mode. Entities other than "bold", "italic", "underline", "strikethrough", "spoiler", and "custom_emoji" are ignored.
}

func (r *SendGiftRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))

	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("gift_id", r.GiftID)
	pay_for_upgrade, _ := json.Marshal(r.PayForUpgrade)
	w.WriteField("pay_for_upgrade", string(pay_for_upgrade))

	w.WriteField("text", r.Text)

	w.WriteField("text_parse_mode", r.TextParseMode)
	if r.TextEntities != nil {
		text_entities, _ := json.Marshal(r.TextEntities)
		w.WriteField("text_entities", string(text_entities))
	}
}

// use Bot.VerifyUser(ctx, &VerifyUserRequest{})
type VerifyUserRequest struct {
	UserID            int64  // Unique identifier of the target user
	CustomDescription string // Custom description for the verification; 0-70 characters. Must be empty if the organization isn't allowed to provide a custom verification description.
}

func (r *VerifyUserRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))

	w.WriteField("custom_description", r.CustomDescription)
}

// use Bot.VerifyChat(ctx, &VerifyChatRequest{})
type VerifyChatRequest struct {
	ChatID            ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	CustomDescription string // Custom description for the verification; 0-70 characters. Must be empty if the organization isn't allowed to provide a custom verification description.
}

func (r *VerifyChatRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())

	w.WriteField("custom_description", r.CustomDescription)
}

// use Bot.RemoveUserVerification(ctx, &RemoveUserVerificationRequest{})
type RemoveUserVerificationRequest struct {
	UserID int64 // Unique identifier of the target user
}

func (r *RemoveUserVerificationRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
}

// use Bot.RemoveChatVerification(ctx, &RemoveChatVerificationRequest{})
type RemoveChatVerificationRequest struct {
	ChatID ChatID // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

func (r *RemoveChatVerificationRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
}

// use Bot.AnswerInlineQuery(ctx, &AnswerInlineQueryRequest{})
type AnswerInlineQueryRequest struct {
	InlineQueryID string                    // Unique identifier for the answered query
	Results       []InlineQueryResult       // A JSON-serialized array of results for the inline query
	CacheTime     int64                     // The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	IsPersonal    bool                      // Pass True if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query.
	NextOffset    string                    // Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don't support pagination. Offset length can't exceed 64 bytes.
	Button        *InlineQueryResultsButton // A JSON-serialized object describing a button to be shown above inline query results
}

func (r *AnswerInlineQueryRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("inline_query_id", r.InlineQueryID)
	if r.Results != nil {
		results, _ := json.Marshal(r.Results)
		w.WriteField("results", string(results))
	}
	cache_time, _ := json.Marshal(r.CacheTime)
	w.WriteField("cache_time", string(cache_time))
	is_personal, _ := json.Marshal(r.IsPersonal)
	w.WriteField("is_personal", string(is_personal))

	w.WriteField("next_offset", r.NextOffset)
	if r.Button != nil {
		button, _ := json.Marshal(r.Button)
		w.WriteField("button", string(button))
	}
}

// use Bot.AnswerWebAppQuery(ctx, &AnswerWebAppQueryRequest{})
type AnswerWebAppQueryRequest struct {
	WebAppQueryID string            // Unique identifier for the query to be answered
	Result        InlineQueryResult // A JSON-serialized object describing the message to be sent
}

func (r *AnswerWebAppQueryRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("web_app_query_id", r.WebAppQueryID)
	if r.Result != nil {
		result, _ := json.Marshal(r.Result)
		w.WriteField("result", string(result))
	}
}

// use Bot.SavePreparedInlineMessage(ctx, &SavePreparedInlineMessageRequest{})
type SavePreparedInlineMessageRequest struct {
	UserID            int64             // Unique identifier of the target user that can use the prepared message
	Result            InlineQueryResult // A JSON-serialized object describing the message to be sent
	AllowUserChats    bool              // Pass True if the message can be sent to private chats with users
	AllowBotChats     bool              // Pass True if the message can be sent to private chats with bots
	AllowGroupChats   bool              // Pass True if the message can be sent to group and supergroup chats
	AllowChannelChats bool              // Pass True if the message can be sent to channel chats
}

func (r *SavePreparedInlineMessageRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
	if r.Result != nil {
		result, _ := json.Marshal(r.Result)
		w.WriteField("result", string(result))
	}
	allow_user_chats, _ := json.Marshal(r.AllowUserChats)
	w.WriteField("allow_user_chats", string(allow_user_chats))
	allow_bot_chats, _ := json.Marshal(r.AllowBotChats)
	w.WriteField("allow_bot_chats", string(allow_bot_chats))
	allow_group_chats, _ := json.Marshal(r.AllowGroupChats)
	w.WriteField("allow_group_chats", string(allow_group_chats))
	allow_channel_chats, _ := json.Marshal(r.AllowChannelChats)
	w.WriteField("allow_channel_chats", string(allow_channel_chats))
}

// use Bot.SendInvoice(ctx, &SendInvoiceRequest{})
type SendInvoiceRequest struct {
	ChatID                    ChatID                // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadID           int                   // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Title                     string                // Product name, 1-32 characters
	Description               string                // Product description, 1-255 characters
	Payload                   string                // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	ProviderToken             string                // Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	Currency                  string                // Three-letter ISO 4217 currency code, see more on currencies. Pass "XTR" for payments in Telegram Stars.
	Prices                    []LabeledPrice        // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
	MaxTipAmount              int64                 // The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
	SuggestedTipAmounts       []int64               // A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	StartParameter            string                // Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a Pay button, allowing multiple users to pay directly from the forwarded message, using the same invoice. If non-empty, forwarded copies of the sent message will have a URL button with a deep link to the bot (instead of a Pay button), with the value used as the start parameter
	ProviderData              string                // JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoUrl                  string                // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoSize                 int64                 // Photo size in bytes
	PhotoWidth                int64                 // Photo width
	PhotoHeight               int64                 // Photo height
	NeedName                  bool                  // Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber           bool                  // Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
	NeedEmail                 bool                  // Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress       bool                  // Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider bool                  // Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvider       bool                  // Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
	IsFlexible                bool                  // Pass True if the final price depends on the shipping method. Ignored for payments in Telegram Stars.
	DisableNotification       bool                  // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent            bool                  // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast        bool                  // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID           int                   // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters           *ReplyParameters      // Description of the message to reply to
	ReplyMarkup               *InlineKeyboardMarkup // A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
}

func (r *SendInvoiceRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("chat_id", r.ChatID.String())
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	w.WriteField("title", r.Title)

	w.WriteField("description", r.Description)

	w.WriteField("payload", r.Payload)

	w.WriteField("provider_token", r.ProviderToken)

	w.WriteField("currency", r.Currency)
	if r.Prices != nil {
		prices, _ := json.Marshal(r.Prices)
		w.WriteField("prices", string(prices))
	}
	max_tip_amount, _ := json.Marshal(r.MaxTipAmount)
	w.WriteField("max_tip_amount", string(max_tip_amount))
	if r.SuggestedTipAmounts != nil {
		suggested_tip_amounts, _ := json.Marshal(r.SuggestedTipAmounts)
		w.WriteField("suggested_tip_amounts", string(suggested_tip_amounts))
	}

	w.WriteField("start_parameter", r.StartParameter)

	w.WriteField("provider_data", r.ProviderData)

	w.WriteField("photo_url", r.PhotoUrl)
	photo_size, _ := json.Marshal(r.PhotoSize)
	w.WriteField("photo_size", string(photo_size))
	photo_width, _ := json.Marshal(r.PhotoWidth)
	w.WriteField("photo_width", string(photo_width))
	photo_height, _ := json.Marshal(r.PhotoHeight)
	w.WriteField("photo_height", string(photo_height))
	need_name, _ := json.Marshal(r.NeedName)
	w.WriteField("need_name", string(need_name))
	need_phone_number, _ := json.Marshal(r.NeedPhoneNumber)
	w.WriteField("need_phone_number", string(need_phone_number))
	need_email, _ := json.Marshal(r.NeedEmail)
	w.WriteField("need_email", string(need_email))
	need_shipping_address, _ := json.Marshal(r.NeedShippingAddress)
	w.WriteField("need_shipping_address", string(need_shipping_address))
	send_phone_number_to_provider, _ := json.Marshal(r.SendPhoneNumberToProvider)
	w.WriteField("send_phone_number_to_provider", string(send_phone_number_to_provider))
	send_email_to_provider, _ := json.Marshal(r.SendEmailToProvider)
	w.WriteField("send_email_to_provider", string(send_email_to_provider))
	is_flexible, _ := json.Marshal(r.IsFlexible)
	w.WriteField("is_flexible", string(is_flexible))
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.CreateInvoiceLink(ctx, &CreateInvoiceLinkRequest{})
type CreateInvoiceLinkRequest struct {
	BusinessConnectionID      string         // Unique identifier of the business connection on behalf of which the link will be created. For payments in Telegram Stars only.
	Title                     string         // Product name, 1-32 characters
	Description               string         // Product description, 1-255 characters
	Payload                   string         // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	ProviderToken             string         // Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	Currency                  string         // Three-letter ISO 4217 currency code, see more on currencies. Pass "XTR" for payments in Telegram Stars.
	Prices                    []LabeledPrice // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
	SubscriptionPeriod        int64          // The number of seconds the subscription will be active for before the next payment. The currency must be set to "XTR" (Telegram Stars) if the parameter is used. Currently, it must always be 2592000 (30 days) if specified. Any number of subscriptions can be active for a given bot at the same time, including multiple concurrent subscriptions from the same user. Subscription price must no exceed 2500 Telegram Stars.
	MaxTipAmount              int64          // The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
	SuggestedTipAmounts       []int64        // A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	ProviderData              string         // JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoUrl                  string         // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoSize                 int64          // Photo size in bytes
	PhotoWidth                int64          // Photo width
	PhotoHeight               int64          // Photo height
	NeedName                  bool           // Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber           bool           // Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
	NeedEmail                 bool           // Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress       bool           // Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider bool           // Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvider       bool           // Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
	IsFlexible                bool           // Pass True if the final price depends on the shipping method. Ignored for payments in Telegram Stars.
}

func (r *CreateInvoiceLinkRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)

	w.WriteField("title", r.Title)

	w.WriteField("description", r.Description)

	w.WriteField("payload", r.Payload)

	w.WriteField("provider_token", r.ProviderToken)

	w.WriteField("currency", r.Currency)
	if r.Prices != nil {
		prices, _ := json.Marshal(r.Prices)
		w.WriteField("prices", string(prices))
	}
	subscription_period, _ := json.Marshal(r.SubscriptionPeriod)
	w.WriteField("subscription_period", string(subscription_period))
	max_tip_amount, _ := json.Marshal(r.MaxTipAmount)
	w.WriteField("max_tip_amount", string(max_tip_amount))
	if r.SuggestedTipAmounts != nil {
		suggested_tip_amounts, _ := json.Marshal(r.SuggestedTipAmounts)
		w.WriteField("suggested_tip_amounts", string(suggested_tip_amounts))
	}

	w.WriteField("provider_data", r.ProviderData)

	w.WriteField("photo_url", r.PhotoUrl)
	photo_size, _ := json.Marshal(r.PhotoSize)
	w.WriteField("photo_size", string(photo_size))
	photo_width, _ := json.Marshal(r.PhotoWidth)
	w.WriteField("photo_width", string(photo_width))
	photo_height, _ := json.Marshal(r.PhotoHeight)
	w.WriteField("photo_height", string(photo_height))
	need_name, _ := json.Marshal(r.NeedName)
	w.WriteField("need_name", string(need_name))
	need_phone_number, _ := json.Marshal(r.NeedPhoneNumber)
	w.WriteField("need_phone_number", string(need_phone_number))
	need_email, _ := json.Marshal(r.NeedEmail)
	w.WriteField("need_email", string(need_email))
	need_shipping_address, _ := json.Marshal(r.NeedShippingAddress)
	w.WriteField("need_shipping_address", string(need_shipping_address))
	send_phone_number_to_provider, _ := json.Marshal(r.SendPhoneNumberToProvider)
	w.WriteField("send_phone_number_to_provider", string(send_phone_number_to_provider))
	send_email_to_provider, _ := json.Marshal(r.SendEmailToProvider)
	w.WriteField("send_email_to_provider", string(send_email_to_provider))
	is_flexible, _ := json.Marshal(r.IsFlexible)
	w.WriteField("is_flexible", string(is_flexible))
}

// use Bot.AnswerShippingQuery(ctx, &AnswerShippingQueryRequest{})
type AnswerShippingQueryRequest struct {
	ShippingQueryID string           // Unique identifier for the query to be answered
	Ok              bool             // Pass True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
	ShippingOptions []ShippingOption // Required if ok is True. A JSON-serialized array of available shipping options.
	ErrorMessage    string           // Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable"). Telegram will display this message to the user.
}

func (r *AnswerShippingQueryRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("shipping_query_id", r.ShippingQueryID)
	ok, _ := json.Marshal(r.Ok)
	w.WriteField("ok", string(ok))
	if r.ShippingOptions != nil {
		shipping_options, _ := json.Marshal(r.ShippingOptions)
		w.WriteField("shipping_options", string(shipping_options))
	}

	w.WriteField("error_message", r.ErrorMessage)
}

// use Bot.AnswerPreCheckoutQuery(ctx, &AnswerPreCheckoutQueryRequest{})
type AnswerPreCheckoutQueryRequest struct {
	PreCheckoutQueryID string // Unique identifier for the query to be answered
	Ok                 bool   // Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
	ErrorMessage       string // Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
}

func (r *AnswerPreCheckoutQueryRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("pre_checkout_query_id", r.PreCheckoutQueryID)
	ok, _ := json.Marshal(r.Ok)
	w.WriteField("ok", string(ok))

	w.WriteField("error_message", r.ErrorMessage)
}

// use Bot.GetStarTransactions(ctx, &GetStarTransactionsRequest{})
type GetStarTransactionsRequest struct {
	Offset int64 // Number of transactions to skip in the response
	Limit  int64 // The maximum number of transactions to be retrieved. Values between 1-100 are accepted. Defaults to 100.
}

func (r *GetStarTransactionsRequest) WriteMultipart(w *multipart.Writer) {
	offset, _ := json.Marshal(r.Offset)
	w.WriteField("offset", string(offset))
	limit, _ := json.Marshal(r.Limit)
	w.WriteField("limit", string(limit))
}

// use Bot.RefundStarPayment(ctx, &RefundStarPaymentRequest{})
type RefundStarPaymentRequest struct {
	UserID                  int64  // Identifier of the user whose payment will be refunded
	TelegramPaymentChargeID string // Telegram payment identifier
}

func (r *RefundStarPaymentRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))

	w.WriteField("telegram_payment_charge_id", r.TelegramPaymentChargeID)
}

// use Bot.EditUserStarSubscription(ctx, &EditUserStarSubscriptionRequest{})
type EditUserStarSubscriptionRequest struct {
	UserID                  int64  // Identifier of the user whose subscription will be edited
	TelegramPaymentChargeID string // Telegram payment identifier for the subscription
	IsCanceled              bool   // Pass True to cancel extension of the user subscription; the subscription must be active up to the end of the current subscription period. Pass False to allow the user to re-enable a subscription that was previously canceled by the bot.
}

func (r *EditUserStarSubscriptionRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))

	w.WriteField("telegram_payment_charge_id", r.TelegramPaymentChargeID)
	is_canceled, _ := json.Marshal(r.IsCanceled)
	w.WriteField("is_canceled", string(is_canceled))
}

// use Bot.SetPassportDataErrors(ctx, &SetPassportDataErrorsRequest{})
type SetPassportDataErrorsRequest struct {
	UserID int64                  // User identifier
	Errors []PassportElementError // A JSON-serialized array describing the errors
}

func (r *SetPassportDataErrorsRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
	if r.Errors != nil {
		errors, _ := json.Marshal(r.Errors)
		w.WriteField("errors", string(errors))
	}
}

// use Bot.SendGame(ctx, &SendGameRequest{})
type SendGameRequest struct {
	BusinessConnectionID string                // Unique identifier of the business connection on behalf of which the message will be sent
	ChatID               int64                 // Unique identifier for the target chat
	MessageThreadID      int                   // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	GameShortName        string                // Short name of the game, serves as the unique identifier for the game. Set up your games via @BotFather.
	DisableNotification  bool                  // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       bool                  // Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   bool                  // Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectID      int                   // Unique identifier of the message effect to be added to the message; for private chats only
	ReplyParameters      *ReplyParameters      // Description of the message to reply to
	ReplyMarkup          *InlineKeyboardMarkup // A JSON-serialized object for an inline keyboard. If empty, one 'Play game_title' button will be shown. If not empty, the first button must launch the game.
}

func (r *SendGameRequest) WriteMultipart(w *multipart.Writer) {
	w.WriteField("business_connection_id", r.BusinessConnectionID)
	chat_id, _ := json.Marshal(r.ChatID)
	w.WriteField("chat_id", string(chat_id))
	message_thread_id, _ := json.Marshal(r.MessageThreadID)
	w.WriteField("message_thread_id", string(message_thread_id))

	w.WriteField("game_short_name", r.GameShortName)
	disable_notification, _ := json.Marshal(r.DisableNotification)
	w.WriteField("disable_notification", string(disable_notification))
	protect_content, _ := json.Marshal(r.ProtectContent)
	w.WriteField("protect_content", string(protect_content))
	allow_paid_broadcast, _ := json.Marshal(r.AllowPaidBroadcast)
	w.WriteField("allow_paid_broadcast", string(allow_paid_broadcast))
	message_effect_id, _ := json.Marshal(r.MessageEffectID)
	w.WriteField("message_effect_id", string(message_effect_id))
	if r.ReplyParameters != nil {
		reply_parameters, _ := json.Marshal(r.ReplyParameters)
		w.WriteField("reply_parameters", string(reply_parameters))
	}
	if r.ReplyMarkup != nil {
		reply_markup, _ := json.Marshal(r.ReplyMarkup)
		w.WriteField("reply_markup", string(reply_markup))
	}
}

// use Bot.SetGameScore(ctx, &SetGameScoreRequest{})
type SetGameScoreRequest struct {
	UserID             int64 // User identifier
	Score              int64 // New score, must be non-negative
	Force              bool  // Pass True if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	DisableEditMessage bool  // Pass True if the game message should not be automatically edited to include the current scoreboard
	ChatID             int64 // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageID          int   // Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageID    int   // Required if chat_id and message_id are not specified. Identifier of the inline message
}

func (r *SetGameScoreRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
	score, _ := json.Marshal(r.Score)
	w.WriteField("score", string(score))
	force, _ := json.Marshal(r.Force)
	w.WriteField("force", string(force))
	disable_edit_message, _ := json.Marshal(r.DisableEditMessage)
	w.WriteField("disable_edit_message", string(disable_edit_message))
	chat_id, _ := json.Marshal(r.ChatID)
	w.WriteField("chat_id", string(chat_id))
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
	inline_message_id, _ := json.Marshal(r.InlineMessageID)
	w.WriteField("inline_message_id", string(inline_message_id))
}

// use Bot.GetGameHighScores(ctx, &GetGameHighScoresRequest{})
type GetGameHighScoresRequest struct {
	UserID          int64 // Target user id
	ChatID          int64 // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageID       int   // Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageID int   // Required if chat_id and message_id are not specified. Identifier of the inline message
}

func (r *GetGameHighScoresRequest) WriteMultipart(w *multipart.Writer) {
	user_id, _ := json.Marshal(r.UserID)
	w.WriteField("user_id", string(user_id))
	chat_id, _ := json.Marshal(r.ChatID)
	w.WriteField("chat_id", string(chat_id))
	message_id, _ := json.Marshal(r.MessageID)
	w.WriteField("message_id", string(message_id))
	inline_message_id, _ := json.Marshal(r.InlineMessageID)
	w.WriteField("inline_message_id", string(inline_message_id))
}
