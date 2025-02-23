package goram

const (
	// New incoming message of any kind - text, photo, sticker, etc.
	UpdateMessage string = "message"

	// New version of a message that is known to the bot and was edited. This update may at times be triggered by changes to message fields that are either unavailable or not actively used by your bot.
	UpdateEditedMessage string = "edited_message"

	// New incoming channel post of any kind - text, photo, sticker, etc.
	UpdateChannelPost string = "channel_post"

	// New version of a channel post that is known to the bot and was edited. This update may at times be triggered by changes to message fields that are either unavailable or not actively used by your bot.
	UpdateEditedChannelPost string = "edited_channel_post"

	// The bot was connected to or disconnected from a business account, or a user edited an existing connection with the bot
	UpdateBusinessConnection string = "business_connection"

	// New message from a connected business account
	UpdateBusinessMessage string = "business_message"

	// New version of a message from a connected business account
	UpdateEditedBusinessMessage string = "edited_business_message"

	// Messages were deleted from a connected business account
	UpdateDeletedBusinessMessages string = "deleted_business_messages"

	// A reaction to a message was changed by a user. The bot must be an administrator in the chat and must explicitly specify "message_reaction" in the list of allowed_updates to receive these updates. The update isn't received for reactions set by bots.
	UpdateMessageReaction string = "message_reaction"

	// Reactions to a message with anonymous reactions were changed. The bot must be an administrator in the chat and must explicitly specify "message_reaction_count" in the list of allowed_updates to receive these updates. The updates are grouped and can be sent with delay up to a few minutes.
	UpdateMessageReactionCount string = "message_reaction_count"

	// New incoming inline query
	UpdateInlineQuery string = "inline_query"

	// The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	UpdateChosenInlineResult string = "chosen_inline_result"

	// New incoming callback query
	UpdateCallbackQuery string = "callback_query"

	// New incoming shipping query. Only for invoices with flexible price
	UpdateShippingQuery string = "shipping_query"

	// New incoming pre-checkout query. Contains full information about checkout
	UpdatePreCheckoutQuery string = "pre_checkout_query"

	// A user purchased paid media with a non-empty payload sent by the bot in a non-channel chat
	UpdatePurchasedPaidMedia string = "purchased_paid_media"

	// New poll state. Bots receive only updates about manually stopped polls and polls, which are sent by the bot
	UpdatePoll string = "poll"

	// A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	UpdatePollAnswer string = "poll_answer"

	// The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
	UpdateMyChatMember string = "my_chat_member"

	// A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
	UpdateChatMember string = "chat_member"

	// A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
	UpdateChatJoinRequest string = "chat_join_request"

	// A chat boost was added or changed. The bot must be an administrator in the chat to receive these updates.
	UpdateChatBoost string = "chat_boost"

	// A boost was removed from a chat. The bot must be an administrator in the chat to receive these updates.
	UpdateRemovedChatBoost string = "removed_chat_boost"
)
