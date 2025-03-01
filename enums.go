package goram

type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSupergroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

type InlineQueryChatType string

const (
	InlineQueryChatTypeSender     InlineQueryChatType = "sender"
	InlineQueryChatTypePrivate    InlineQueryChatType = "private"
	InlineQueryChatTypeGroup      InlineQueryChatType = "group"
	InlineQueryChatTypeSupergroup InlineQueryChatType = "supergroup"
	InlineQueryChatTypeChannel    InlineQueryChatType = "channel"
)

type ChatAction string

const (
	ChatActionTyping          ChatAction = "typing"
	ChatActionUploadPhoto     ChatAction = "upload_photo"
	ChatActionRecordVideo     ChatAction = "record_video"
	ChatActionUploadVideo     ChatAction = "upload_video"
	ChatActionRecordVoice     ChatAction = "record_voice"
	ChatActionUploadVoice     ChatAction = "upload_voice"
	ChatActionUploadDocument  ChatAction = "upload_document"
	ChatActionFindLocation    ChatAction = "find_location"
	ChatActionRecordVideoNote ChatAction = "record_video_note"
	ChatActionUploadVideoNote ChatAction = "upload_video_note"
)

type MessageEntityType string

const (
	MessageEntityTypeMention              MessageEntityType = "mention"
	MessageEntityTypeHashtag              MessageEntityType = "hashtag"
	MessageEntityTypeCashtag              MessageEntityType = "cashtag"
	MessageEntityTypeBotCommand           MessageEntityType = "bot_command"
	MessageEntityTypeUrl                  MessageEntityType = "url"
	MessageEntityTypeEmail                MessageEntityType = "email"
	MessageEntityTypePhoneNumber          MessageEntityType = "phone_number"
	MessageEntityTypeBold                 MessageEntityType = "bold"
	MessageEntityTypeItalic               MessageEntityType = "italic"
	MessageEntityTypeUnderline            MessageEntityType = "underline"
	MessageEntityTypeStrikethrough        MessageEntityType = "strikethrough"
	MessageEntityTypeSpoiler              MessageEntityType = "spoiler"
	MessageEntityTypeBlockquote           MessageEntityType = "blockquote"
	MessageEntityTypeExpandableBlockquote MessageEntityType = "expandable_blockquote"
	MessageEntityTypeCode                 MessageEntityType = "code"
	MessageEntityTypePre                  MessageEntityType = "pre"
	MessageEntityTypeTextLink             MessageEntityType = "text_link"
	MessageEntityTypeTextMention          MessageEntityType = "text_mention"
	MessageEntityTypeCustomEmoji          MessageEntityType = "custom_emoji"
)

type ParseMode string

const (
	ParseModeMarkdown   ParseMode = "Markdown"
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
	ParseModeHTML       ParseMode = "HTML"
)

type UpdateType string

const (
	// New incoming message of any kind - text, photo, sticker, etc.
	UpdateMessage UpdateType = "message"

	// New version of a message that is known to the bot and was edited. This update may at times be triggered by changes to message fields that are either unavailable or not actively used by your bot.
	UpdateEditedMessage UpdateType = "edited_message"

	// New incoming channel post of any kind - text, photo, sticker, etc.
	UpdateChannelPost UpdateType = "channel_post"

	// New version of a channel post that is known to the bot and was edited. This update may at times be triggered by changes to message fields that are either unavailable or not actively used by your bot.
	UpdateEditedChannelPost UpdateType = "edited_channel_post"

	// The bot was connected to or disconnected from a business account, or a user edited an existing connection with the bot
	UpdateBusinessConnection UpdateType = "business_connection"

	// New message from a connected business account
	UpdateBusinessMessage UpdateType = "business_message"

	// New version of a message from a connected business account
	UpdateEditedBusinessMessage UpdateType = "edited_business_message"

	// Messages were deleted from a connected business account
	UpdateDeletedBusinessMessages UpdateType = "deleted_business_messages"

	// A reaction to a message was changed by a user. The bot must be an administrator in the chat and must explicitly specify "message_reaction" in the list of allowed_updates to receive these updates. The update isn't received for reactions set by bots.
	UpdateMessageReaction UpdateType = "message_reaction"

	// Reactions to a message with anonymous reactions were changed. The bot must be an administrator in the chat and must explicitly specify "message_reaction_count" in the list of allowed_updates to receive these updates. The updates are grouped and can be sent with delay up to a few minutes.
	UpdateMessageReactionCount UpdateType = "message_reaction_count"

	// New incoming inline query
	UpdateInlineQuery UpdateType = "inline_query"

	// The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	UpdateChosenInlineResult UpdateType = "chosen_inline_result"

	// New incoming callback query
	UpdateCallbackQuery UpdateType = "callback_query"

	// New incoming shipping query. Only for invoices with flexible price
	UpdateShippingQuery UpdateType = "shipping_query"

	// New incoming pre-checkout query. Contains full information about checkout
	UpdatePreCheckoutQuery UpdateType = "pre_checkout_query"

	// A user purchased paid media with a non-empty payload sent by the bot in a non-channel chat
	UpdatePurchasedPaidMedia UpdateType = "purchased_paid_media"

	// New poll state. Bots receive only updates about manually stopped polls and polls, which are sent by the bot
	UpdatePoll UpdateType = "poll"

	// A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	UpdatePollAnswer UpdateType = "poll_answer"

	// The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
	UpdateMyChatMember UpdateType = "my_chat_member"

	// A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
	UpdateChatMember UpdateType = "chat_member"

	// A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
	UpdateChatJoinRequest UpdateType = "chat_join_request"

	// A chat boost was added or changed. The bot must be an administrator in the chat to receive these updates.
	UpdateChatBoost UpdateType = "chat_boost"

	// A boost was removed from a chat. The bot must be an administrator in the chat to receive these updates.
	UpdateRemovedChatBoost UpdateType = "removed_chat_boost"
)
