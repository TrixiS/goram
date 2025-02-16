package types

type ChatID struct {
	ID       int64
	Username string
}

// This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	UpdateID                int64                        `json:"update_id"`                 // The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially. This identifier becomes especially handy if you're using webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	Message                 *Message                     `json:"message"`                   // Optional. New incoming message of any kind - text, photo, sticker, etc.
	EditedMessage           *Message                     `json:"edited_message"`            // Optional. New version of a message that is known to the bot and was edited. This update may at times be triggered by changes to message fields that are either unavailable or not actively used by your bot.
	ChannelPost             *Message                     `json:"channel_post"`              // Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	EditedChannelPost       *Message                     `json:"edited_channel_post"`       // Optional. New version of a channel post that is known to the bot and was edited. This update may at times be triggered by changes to message fields that are either unavailable or not actively used by your bot.
	BusinessConnection      *BusinessConnection          `json:"business_connection"`       // Optional. The bot was connected to or disconnected from a business account, or a user edited an existing connection with the bot
	BusinessMessage         *Message                     `json:"business_message"`          // Optional. New message from a connected business account
	EditedBusinessMessage   *Message                     `json:"edited_business_message"`   // Optional. New version of a message from a connected business account
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages"` // Optional. Messages were deleted from a connected business account
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction"`          // Optional. A reaction to a message was changed by a user. The bot must be an administrator in the chat and must explicitly specify "message_reaction" in the list of allowed_updates to receive these updates. The update isn't received for reactions set by bots.
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count"`    // Optional. Reactions to a message with anonymous reactions were changed. The bot must be an administrator in the chat and must explicitly specify "message_reaction_count" in the list of allowed_updates to receive these updates. The updates are grouped and can be sent with delay up to a few minutes.
	InlineQuery             *InlineQuery                 `json:"inline_query"`              // Optional. New incoming inline query
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result"`      // Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	CallbackQuery           *CallbackQuery               `json:"callback_query"`            // Optional. New incoming callback query
	ShippingQuery           *ShippingQuery               `json:"shipping_query"`            // Optional. New incoming shipping query. Only for invoices with flexible price
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query"`        // Optional. New incoming pre-checkout query. Contains full information about checkout
	PurchasedPaIDMedia      *PaidMediaPurchased          `json:"purchased_paid_media"`      // Optional. A user purchased paid media with a non-empty payload sent by the bot in a non-channel chat
	Poll                    *Poll                        `json:"poll"`                      // Optional. New poll state. Bots receive only updates about manually stopped polls and polls, which are sent by the bot
	PollAnswer              *PollAnswer                  `json:"poll_answer"`               // Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member"`            // Optional. The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
	ChatMember              ChatMemberUpdated            `json:"chat_member"`               // Optional. A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request"`         // Optional. A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost"`                // Optional. A chat boost was added or changed. The bot must be an administrator in the chat to receive these updates.
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost"`        // Optional. A boost was removed from a chat. The bot must be an administrator in the chat to receive these updates.
}

// Describes the current status of a webhook.
type WebhookInfo struct {
	Url                          string   `json:"url"`                             // Webhook URL, may be empty if webhook is not set up
	HasCustomCertificate         bool     `json:"has_custom_certificate"`          // True, if a custom certificate was provided for webhook certificate checks
	PendingUpdateCount           int64    `json:"pending_update_count"`            // Number of updates awaiting delivery
	IpAddress                    string   `json:"ip_address"`                      // Optional. Currently used webhook IP address
	LastErrorDate                int64    `json:"last_error_date"`                 // Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorMessage             string   `json:"last_error_message"`              // Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
	LastSynchronizationErrorDate int64    `json:"last_synchronization_error_date"` // Optional. Unix time of the most recent error that happened when trying to synchronize available updates with Telegram datacenters
	MaxConnections               int64    `json:"max_connections"`                 // Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	AllowedUpdates               []string `json:"allowed_updates"`                 // Optional. A list of update types the bot is subscribed to. Defaults to all update types except chat_member
}

// This object represents a Telegram user or bot.
type User struct {
	ID                      int64  `json:"id"`                          // Unique identifier for this user or bot. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	IsBot                   bool   `json:"is_bot"`                      // True, if this user is a bot
	FirstName               string `json:"first_name"`                  // User's or bot's first name
	LastName                string `json:"last_name"`                   // Optional. User's or bot's last name
	Username                string `json:"username"`                    // Optional. User's or bot's username
	LanguageCode            string `json:"language_code"`               // Optional. IETF language tag of the user's language
	IsPremium               bool   `json:"is_premium"`                  // Optional. True, if this user is a Telegram Premium user
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu"`    // Optional. True, if this user added the bot to the attachment menu
	CanJoinGroups           bool   `json:"can_join_groups"`             // Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"` // Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`     // Optional. True, if the bot supports inline queries. Returned only in getMe.
	CanConnectToBusiness    bool   `json:"can_connect_to_business"`     // Optional. True, if the bot can be connected to a Telegram Business account to receive its messages. Returned only in getMe.
	HasMainWebApp           bool   `json:"has_main_web_app"`            // Optional. True, if the bot has a main Web App. Returned only in getMe.
}

// This object represents a chat.
type Chat struct {
	ID        int64    `json:"id"`         // Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Type      ChatType `json:"type"`       // Type of the chat, can be either "private", "group", "supergroup" or "channel"
	Title     string   `json:"title"`      // Optional. Title, for supergroups, channels and group chats
	Username  string   `json:"username"`   // Optional. Username, for private chats, supergroups and channels if available
	FirstName string   `json:"first_name"` // Optional. First name of the other party in a private chat
	LastName  string   `json:"last_name"`  // Optional. Last name of the other party in a private chat
	IsForum   bool     `json:"is_forum"`   // Optional. True, if the supergroup chat is a forum (has topics enabled)
}

// This object contains full information about a chat.
type ChatFullInfo struct {
	ID                                 int64                 `json:"id"`                                      // Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Type                               ChatType              `json:"type"`                                    // Type of the chat, can be either "private", "group", "supergroup" or "channel"
	Title                              string                `json:"title"`                                   // Optional. Title, for supergroups, channels and group chats
	Username                           string                `json:"username"`                                // Optional. Username, for private chats, supergroups and channels if available
	FirstName                          string                `json:"first_name"`                              // Optional. First name of the other party in a private chat
	LastName                           string                `json:"last_name"`                               // Optional. Last name of the other party in a private chat
	IsForum                            bool                  `json:"is_forum"`                                // Optional. True, if the supergroup chat is a forum (has topics enabled)
	AccentColorID                      int64                 `json:"accent_color_id"`                         // Identifier of the accent color for the chat name and backgrounds of the chat photo, reply header, and link preview. See accent colors for more details.
	MaxReactionCount                   int64                 `json:"max_reaction_count"`                      // The maximum number of reactions that can be set on a message in the chat
	Photo                              *ChatPhoto            `json:"photo"`                                   // Optional. Chat photo
	ActiveUsernames                    []string              `json:"active_usernames"`                        // Optional. If non-empty, the list of all active chat usernames; for private chats, supergroups and channels
	Birthdate                          *Birthdate            `json:"birthdate"`                               // Optional. For private chats, the date of birth of the user
	BusinessIntro                      *BusinessIntro        `json:"business_intro"`                          // Optional. For private chats with business accounts, the intro of the business
	BusinessLocation                   *BusinessLocation     `json:"business_location"`                       // Optional. For private chats with business accounts, the location of the business
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours"`                  // Optional. For private chats with business accounts, the opening hours of the business
	PersonalChat                       *Chat                 `json:"personal_chat"`                           // Optional. For private chats, the personal channel of the user
	AvailableReactions                 []ReactionType        `json:"available_reactions"`                     // Optional. List of available reactions allowed in the chat. If omitted, then all emoji reactions are allowed.
	BackgroundCustomEmojiID            string                `json:"background_custom_emoji_id"`              // Optional. Custom emoji identifier of the emoji chosen by the chat for the reply header and link preview background
	ProfileAccentColorID               int64                 `json:"profile_accent_color_id"`                 // Optional. Identifier of the accent color for the chat's profile background. See profile accent colors for more details.
	ProfileBackgroundCustomEmojiID     string                `json:"profile_background_custom_emoji_id"`      // Optional. Custom emoji identifier of the emoji chosen by the chat for its profile background
	EmojiStatusCustomEmojiID           string                `json:"emoji_status_custom_emoji_id"`            // Optional. Custom emoji identifier of the emoji status of the chat or the other party in a private chat
	EmojiStatusExpirationDate          int64                 `json:"emoji_status_expiration_date"`            // Optional. Expiration date of the emoji status of the chat or the other party in a private chat, in Unix time, if any
	Bio                                string                `json:"bio"`                                     // Optional. Bio of the other party in a private chat
	HasPrivateForwards                 bool                  `json:"has_private_forwards"`                    // Optional. True, if privacy settings of the other party in the private chat allows to use tg://user?id=<user_id> links only in chats with the user
	HasRestrictedVoiceAndVIDeoMessages bool                  `json:"has_restricted_voice_and_video_messages"` // Optional. True, if the privacy settings of the other party restrict sending voice and video note messages in the private chat
	JoinToSendMessages                 bool                  `json:"join_to_send_messages"`                   // Optional. True, if users need to join the supergroup before they can send messages
	JoinByRequest                      bool                  `json:"join_by_request"`                         // Optional. True, if all users directly joining the supergroup without using an invite link need to be approved by supergroup administrators
	Description                        string                `json:"description"`                             // Optional. Description, for groups, supergroups and channel chats
	InviteLink                         string                `json:"invite_link"`                             // Optional. Primary invite link, for groups, supergroups and channel chats
	PinnedMessage                      *Message              `json:"pinned_message"`                          // Optional. The most recent pinned message (by sending date)
	Permissions                        *ChatPermissions      `json:"permissions"`                             // Optional. Default chat member permissions, for groups and supergroups
	CanSendGift                        bool                  `json:"can_send_gift"`                           // Optional. True, if gifts can be sent to the chat
	CanSendPaIDMedia                   bool                  `json:"can_send_paid_media"`                     // Optional. True, if paid media messages can be sent or forwarded to the channel chat. The field is available only for channel chats.
	SlowModeDelay                      int64                 `json:"slow_mode_delay"`                         // Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unprivileged user; in seconds
	UnrestrictBoostCount               int64                 `json:"unrestrict_boost_count"`                  // Optional. For supergroups, the minimum number of boosts that a non-administrator user needs to add in order to ignore slow mode and chat permissions
	MessageAutoDeleteTime              int64                 `json:"message_auto_delete_time"`                // Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds
	HasAggressiveAntiSpamEnabled       bool                  `json:"has_aggressive_anti_spam_enabled"`        // Optional. True, if aggressive anti-spam checks are enabled in the supergroup. The field is only available to chat administrators.
	HasHIDdenMembers                   bool                  `json:"has_hidden_members"`                      // Optional. True, if non-administrators can only get the list of bots and administrators in the chat
	HasProtectedContent                bool                  `json:"has_protected_content"`                   // Optional. True, if messages from the chat can't be forwarded to other chats
	HasVisibleHistory                  bool                  `json:"has_visible_history"`                     // Optional. True, if new chat members will have access to old messages; available only to chat administrators
	StickerSetName                     string                `json:"sticker_set_name"`                        // Optional. For supergroups, name of the group sticker set
	CanSetStickerSet                   bool                  `json:"can_set_sticker_set"`                     // Optional. True, if the bot can change the group sticker set
	CustomEmojiStickerSetName          string                `json:"custom_emoji_sticker_set_name"`           // Optional. For supergroups, the name of the group's custom emoji sticker set. Custom emoji from this set can be used by all users and bots in the group.
	LinkedChatID                       ChatID                `json:"linked_chat_id"`                          // Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	Location                           *ChatLocation         `json:"location"`                                // Optional. For supergroups, the location to which the supergroup is connected
}

// This object represents a message.
type Message struct {
	MessageID                     int64                          `json:"message_id"`                        // Unique message identifier inside this chat. In specific instances (e.g., message containing a video sent to a big chat), the server might automatically schedule a message instead of sending it immediately. In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
	MessageThreadID               int64                          `json:"message_thread_id"`                 // Optional. Unique identifier of a message thread to which the message belongs; for supergroups only
	From                          *User                          `json:"from"`                              // Optional. Sender of the message; may be empty for messages sent to channels. For backward compatibility, if the message was sent on behalf of a chat, the field contains a fake sender user in non-channel chats
	SenderChat                    *Chat                          `json:"sender_chat"`                       // Optional. Sender of the message when sent on behalf of a chat. For example, the supergroup itself for messages sent by its anonymous administrators or a linked channel for messages automatically forwarded to the channel's discussion group. For backward compatibility, if the message was sent on behalf of a chat, the field from contains a fake sender user in non-channel chats.
	SenderBoostCount              int64                          `json:"sender_boost_count"`                // Optional. If the sender of the message boosted the chat, the number of boosts added by the user
	SenderBusinessBot             *User                          `json:"sender_business_bot"`               // Optional. The bot that actually sent the message on behalf of the business account. Available only for outgoing messages sent on behalf of the connected business account.
	Date                          int64                          `json:"date"`                              // Date the message was sent in Unix time. It is always a positive number, representing a valid date.
	BusinessConnectionID          string                         `json:"business_connection_id"`            // Optional. Unique identifier of the business connection from which the message was received. If non-empty, the message belongs to a chat of the corresponding business account that is independent from any potential bot chat which might share the same identifier.
	Chat                          *Chat                          `json:"chat"`                              // Chat the message belongs to
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin"`                    // Optional. Information about the original message for forwarded messages
	IsTopicMessage                bool                           `json:"is_topic_message"`                  // Optional. True, if the message is sent to a forum topic
	IsAutomaticForward            bool                           `json:"is_automatic_forward"`              // Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
	ReplyToMessage                *Message                       `json:"reply_to_message"`                  // Optional. For replies in the same chat and message thread, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply"`                    // Optional. Information about the message that is being replied to, which may come from another chat or forum topic
	Quote                         *TextQuote                     `json:"quote"`                             // Optional. For replies that quote part of the original message, the quoted part of the message
	ReplyToStory                  *Story                         `json:"reply_to_story"`                    // Optional. For replies to a story, the original story
	ViaBot                        *User                          `json:"via_bot"`                           // Optional. Bot through which the message was sent
	EditDate                      int64                          `json:"edit_date"`                         // Optional. Date the message was last edited in Unix time
	HasProtectedContent           bool                           `json:"has_protected_content"`             // Optional. True, if the message can't be forwarded
	IsFromOffline                 bool                           `json:"is_from_offline"`                   // Optional. True, if the message was sent by an implicit action, for example, as an away or a greeting business message, or as a scheduled message
	MediaGroupID                  string                         `json:"media_group_id"`                    // Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature               string                         `json:"author_signature"`                  // Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	Text                          string                         `json:"text"`                              // Optional. For text messages, the actual UTF-8 text of the message
	Entities                      []MessageEntity                `json:"entities"`                          // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options"`              // Optional. Options used for link preview generation for the message, if it is a text message and link preview options were changed
	EffectID                      string                         `json:"effect_id"`                         // Optional. Unique identifier of the message effect added to the message
	Animation                     *Animation                     `json:"animation"`                         // Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Audio                         *Audio                         `json:"audio"`                             // Optional. Message is an audio file, information about the file
	Document                      *Document                      `json:"document"`                          // Optional. Message is a general file, information about the file
	PaIDMedia                     *PaidMediaInfo                 `json:"paid_media"`                        // Optional. Message contains paid media; information about the paid media
	Photo                         []PhotoSize                    `json:"photo"`                             // Optional. Message is a photo, available sizes of the photo
	Sticker                       *Sticker                       `json:"sticker"`                           // Optional. Message is a sticker, information about the sticker
	Story                         *Story                         `json:"story"`                             // Optional. Message is a forwarded story
	VIDeo                         *Video                         `json:"video"`                             // Optional. Message is a video, information about the video
	VIDeoNote                     *VideoNote                     `json:"video_note"`                        // Optional. Message is a video note, information about the video message
	Voice                         *Voice                         `json:"voice"`                             // Optional. Message is a voice message, information about the file
	Caption                       string                         `json:"caption"`                           // Optional. Caption for the animation, audio, document, paid media, photo, video or voice
	CaptionEntities               []MessageEntity                `json:"caption_entities"`                  // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media"`          // Optional. True, if the caption must be shown above the message media
	HasMediaSpoiler               bool                           `json:"has_media_spoiler"`                 // Optional. True, if the message media is covered by a spoiler animation
	Contact                       *Contact                       `json:"contact"`                           // Optional. Message is a shared contact, information about the contact
	Dice                          *Dice                          `json:"dice"`                              // Optional. Message is a dice with random value
	Game                          *Game                          `json:"game"`                              // Optional. Message is a game, information about the game. More about games: https://core.telegram.org/bots/api#games
	Poll                          *Poll                          `json:"poll"`                              // Optional. Message is a native poll, information about the poll
	Venue                         *Venue                         `json:"venue"`                             // Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set, the location field will also be set
	Location                      *Location                      `json:"location"`                          // Optional. Message is a shared location, information about the location
	NewChatMembers                []User                         `json:"new_chat_members"`                  // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	LeftChatMember                *User                          `json:"left_chat_member"`                  // Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle                  string                         `json:"new_chat_title"`                    // Optional. A chat title was changed to this value
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo"`                    // Optional. A chat photo was change to this value
	DeleteChatPhoto               bool                           `json:"delete_chat_photo"`                 // Optional. Service message: the chat photo was deleted
	GroupChatCreated              bool                           `json:"group_chat_created"`                // Optional. Service message: the group has been created
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created"`           // Optional. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	ChannelChatCreated            bool                           `json:"channel_chat_created"`              // Optional. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"` // Optional. Service message: auto-delete timer settings changed in the chat
	MigrateToChatID               ChatID                         `json:"migrate_to_chat_id"`                // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatID             ChatID                         `json:"migrate_from_chat_id"`              // Optional. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	PinnedMessage                 *MaybeInaccessibleMessage      `json:"pinned_message"`                    // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	Invoice                       *Invoice                       `json:"invoice"`                           // Optional. Message is an invoice for a payment, information about the invoice. More about payments: https://core.telegram.org/bots/api#payments
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment"`                // Optional. Message is a service message about a successful payment, information about the payment. More about payments: https://core.telegram.org/bots/api#payments
	RefundedPayment               *RefundedPayment               `json:"refunded_payment"`                  // Optional. Message is a service message about a refunded payment, information about the payment. More about payments: https://core.telegram.org/bots/api#payments
	UsersShared                   *UsersShared                   `json:"users_shared"`                      // Optional. Service message: users were shared with the bot
	ChatShared                    *ChatShared                    `json:"chat_shared"`                       // Optional. Service message: a chat was shared with the bot
	ConnectedWebsite              string                         `json:"connected_website"`                 // Optional. The domain name of the website on which the user has logged in. More about Telegram Login: https://core.telegram.org/widgets/login
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed"`              // Optional. Service message: the user allowed the bot to write messages after adding it to the attachment or side menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess
	PassportData                  *PassportData                  `json:"passport_data"`                     // Optional. Telegram Passport data
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered"`         // Optional. Service message. A user in the chat triggered another user's proximity alert while sharing Live Location.
	BoostAdded                    *ChatBoostAdded                `json:"boost_added"`                       // Optional. Service message: user boosted the chat
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set"`               // Optional. Service message: chat background set
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created"`               // Optional. Service message: forum topic created
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited"`                // Optional. Service message: forum topic edited
	ForumTopicClosed              ForumTopicClosed               `json:"forum_topic_closed"`                // Optional. Service message: forum topic closed
	ForumTopicReopened            ForumTopicReopened             `json:"forum_topic_reopened"`              // Optional. Service message: forum topic reopened
	GeneralForumTopicHIDden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden"`        // Optional. Service message: the 'General' forum topic hidden
	GeneralForumTopicUnhIDden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden"`      // Optional. Service message: the 'General' forum topic unhidden
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created"`                  // Optional. Service message: a scheduled giveaway was created
	Giveaway                      *Giveaway                      `json:"giveaway"`                          // Optional. The message is a scheduled giveaway message
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners"`                  // Optional. A giveaway with public winners was completed
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed"`                // Optional. Service message: a giveaway without public winners was completed
	VIDeoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled"`              // Optional. Service message: video chat scheduled
	VIDeoChatStarted              *VideoChatStarted              `json:"video_chat_started"`                // Optional. Service message: video chat started
	VIDeoChatEnded                *VideoChatEnded                `json:"video_chat_ended"`                  // Optional. Service message: video chat ended
	VIDeoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited"`   // Optional. Service message: new participants invited to a video chat
	WebAppData                    *WebAppData                    `json:"web_app_data"`                      // Optional. Service message: data sent by a Web App
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup"`                      // Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
}

// This object represents a unique message identifier.
type MessageId struct {
	MessageID int64 `json:"message_id"` // Unique message identifier. In specific instances (e.g., message containing a video sent to a big chat), the server might automatically schedule a message instead of sending it immediately. In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
}

// This object describes a message that was deleted or is otherwise inaccessible to the bot.
type InaccessibleMessage struct {
	Chat      *Chat `json:"chat"`       // Chat the message belonged to
	MessageID int64 `json:"message_id"` // Unique message identifier inside the chat
	Date      int64 `json:"date"`       // Always 0. The field can be used to differentiate regular and inaccessible messages.
}

// This object describes a message that can be inaccessible to the bot. It can be one of
// - Message
// - InaccessibleMessage
type MaybeInaccessibleMessage interface{}

// This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type          MessageEntityType `json:"type"`            // Type of the entity. Currently, can be "mention" (@username), "hashtag" (#hashtag or #hashtag@chatusername), "cashtag" ($USD or $USD@chatusername), "bot_command" (/start@jobs_bot), "url" (https://telegram.org), "email" (do-not-reply@telegram.org), "phone_number" (+1-212-555-0123), "bold" (bold text), "italic" (italic text), "underline" (underlined text), "strikethrough" (strikethrough text), "spoiler" (spoiler message), "blockquote" (block quotation), "expandable_blockquote" (collapsed-by-default block quotation), "code" (monowidth string), "pre" (monowidth block), "text_link" (for clickable text URLs), "text_mention" (for users without usernames), "custom_emoji" (for inline custom emoji stickers)
	Offset        int64             `json:"offset"`          // Offset in UTF-16 code units to the start of the entity
	Length        int64             `json:"length"`          // Length of the entity in UTF-16 code units
	Url           string            `json:"url"`             // Optional. For "text_link" only, URL that will be opened after user taps on the text
	User          *User             `json:"user"`            // Optional. For "text_mention" only, the mentioned user
	Language      string            `json:"language"`        // Optional. For "pre" only, the programming language of the entity text
	CustomEmojiID string            `json:"custom_emoji_id"` // Optional. For "custom_emoji" only, unique identifier of the custom emoji. Use getCustomEmojiStickers to get full information about the sticker
}

// This object contains information about the quoted part of a message that is replied to by the given message.
type TextQuote struct {
	Text     string          `json:"text"`      // Text of the quoted part of a message that is replied to by the given message
	Entities []MessageEntity `json:"entities"`  // Optional. Special entities that appear in the quote. Currently, only bold, italic, underline, strikethrough, spoiler, and custom_emoji entities are kept in quotes.
	Position int64           `json:"position"`  // Approximate quote position in the original message in UTF-16 code units as specified by the sender
	IsManual bool            `json:"is_manual"` // Optional. True, if the quote was chosen manually by the message sender. Otherwise, the quote was added automatically by the server.
}

// This object contains information about a message that is being replied to, which may come from another chat or forum topic.
type ExternalReplyInfo struct {
	Origin             *MessageOrigin      `json:"origin"`               // Origin of the message replied to by the given message
	Chat               *Chat               `json:"chat"`                 // Optional. Chat the original message belongs to. Available only if the chat is a supergroup or a channel.
	MessageID          int64               `json:"message_id"`           // Optional. Unique message identifier inside the original chat. Available only if the original chat is a supergroup or a channel.
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options"` // Optional. Options used for link preview generation for the original message, if it is a text message
	Animation          *Animation          `json:"animation"`            // Optional. Message is an animation, information about the animation
	Audio              *Audio              `json:"audio"`                // Optional. Message is an audio file, information about the file
	Document           *Document           `json:"document"`             // Optional. Message is a general file, information about the file
	PaIDMedia          *PaidMediaInfo      `json:"paid_media"`           // Optional. Message contains paid media; information about the paid media
	Photo              []PhotoSize         `json:"photo"`                // Optional. Message is a photo, available sizes of the photo
	Sticker            *Sticker            `json:"sticker"`              // Optional. Message is a sticker, information about the sticker
	Story              *Story              `json:"story"`                // Optional. Message is a forwarded story
	VIDeo              *Video              `json:"video"`                // Optional. Message is a video, information about the video
	VIDeoNote          *VideoNote          `json:"video_note"`           // Optional. Message is a video note, information about the video message
	Voice              *Voice              `json:"voice"`                // Optional. Message is a voice message, information about the file
	HasMediaSpoiler    bool                `json:"has_media_spoiler"`    // Optional. True, if the message media is covered by a spoiler animation
	Contact            *Contact            `json:"contact"`              // Optional. Message is a shared contact, information about the contact
	Dice               *Dice               `json:"dice"`                 // Optional. Message is a dice with random value
	Game               *Game               `json:"game"`                 // Optional. Message is a game, information about the game. More about games: https://core.telegram.org/bots/api#games
	Giveaway           *Giveaway           `json:"giveaway"`             // Optional. Message is a scheduled giveaway, information about the giveaway
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners"`     // Optional. A giveaway with public winners was completed
	Invoice            *Invoice            `json:"invoice"`              // Optional. Message is an invoice for a payment, information about the invoice. More about payments: https://core.telegram.org/bots/api#payments
	Location           *Location           `json:"location"`             // Optional. Message is a shared location, information about the location
	Poll               *Poll               `json:"poll"`                 // Optional. Message is a native poll, information about the poll
	Venue              *Venue              `json:"venue"`                // Optional. Message is a venue, information about the venue
}

// Describes reply parameters for the message that is being sent.
type ReplyParameters struct {
	MessageID                int64           `json:"message_id"`                  // Identifier of the message that will be replied to in the current chat, or in the chat chat_id if it is specified
	ChatID                   ChatID          `json:"chat_id"`                     // Optional. If the message to be replied to is from a different chat, unique identifier for the chat or username of the channel (in the format @channelusername). Not supported for messages sent on behalf of a business account.
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"` // Optional. Pass True if the message should be sent even if the specified message to be replied to is not found. Always False for replies in another chat or forum topic. Always True for messages sent on behalf of a business account.
	Quote                    string          `json:"quote"`                       // Optional. Quoted part of the message to be replied to; 0-1024 characters after entities parsing. The quote must be an exact substring of the message to be replied to, including bold, italic, underline, strikethrough, spoiler, and custom_emoji entities. The message will fail to send if the quote isn't found in the original message.
	QuoteParseMode           string          `json:"quote_parse_mode"`            // Optional. Mode for parsing entities in the quote. See formatting options for more details.
	QuoteEntities            []MessageEntity `json:"quote_entities"`              // Optional. A JSON-serialized list of special entities that appear in the quote. It can be specified instead of quote_parse_mode.
	QuotePosition            int64           `json:"quote_position"`              // Optional. Position of the quote in the original message in UTF-16 code units
}

// This object describes the origin of a message. It can be one of
// - MessageOriginUser
// - MessageOriginHiddenUser
// - MessageOriginChat
// - MessageOriginChannel
type MessageOrigin interface{}

// The message was originally sent by a known user.
type MessageOriginUser struct {
	Type       string `json:"type"`        // Type of the message origin, always "user"
	Date       int64  `json:"date"`        // Date the message was sent originally in Unix time
	SenderUser *User  `json:"sender_user"` // User that sent the message originally
}

// The message was originally sent by an unknown user.
type MessageOriginHiddenUser struct {
	Type           string `json:"type"`             // Type of the message origin, always "hidden_user"
	Date           int64  `json:"date"`             // Date the message was sent originally in Unix time
	SenderUserName string `json:"sender_user_name"` // Name of the user that sent the message originally
}

// The message was originally sent on behalf of a chat to a group chat.
type MessageOriginChat struct {
	Type            string `json:"type"`             // Type of the message origin, always "chat"
	Date            int64  `json:"date"`             // Date the message was sent originally in Unix time
	SenderChat      *Chat  `json:"sender_chat"`      // Chat that sent the message originally
	AuthorSignature string `json:"author_signature"` // Optional. For messages originally sent by an anonymous chat administrator, original message author signature
}

// The message was originally sent to a channel chat.
type MessageOriginChannel struct {
	Type            string `json:"type"`             // Type of the message origin, always "channel"
	Date            int64  `json:"date"`             // Date the message was sent originally in Unix time
	Chat            *Chat  `json:"chat"`             // Channel chat to which the message was originally sent
	MessageID       int64  `json:"message_id"`       // Unique message identifier inside the chat
	AuthorSignature string `json:"author_signature"` // Optional. Signature of the original post author
}

// This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	WIDth        int64  `json:"width"`          // Photo width
	Height       int64  `json:"height"`         // Photo height
	FileSize     int64  `json:"file_size"`      // Optional. File size in bytes
}

// This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	FileID       string     `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	WIDth        int64      `json:"width"`          // Video width as defined by the sender
	Height       int64      `json:"height"`         // Video height as defined by the sender
	Duration     int64      `json:"duration"`       // Duration of the video in seconds as defined by the sender
	Thumbnail    *PhotoSize `json:"thumbnail"`      // Optional. Animation thumbnail as defined by the sender
	FileName     string     `json:"file_name"`      // Optional. Original animation filename as defined by the sender
	MimeType     string     `json:"mime_type"`      // Optional. MIME type of the file as defined by the sender
	FileSize     int64      `json:"file_size"`      // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}

// This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID       string     `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int64      `json:"duration"`       // Duration of the audio in seconds as defined by the sender
	Performer    string     `json:"performer"`      // Optional. Performer of the audio as defined by the sender or by audio tags
	Title        string     `json:"title"`          // Optional. Title of the audio as defined by the sender or by audio tags
	FileName     string     `json:"file_name"`      // Optional. Original filename as defined by the sender
	MimeType     string     `json:"mime_type"`      // Optional. MIME type of the file as defined by the sender
	FileSize     int64      `json:"file_size"`      // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	Thumbnail    *PhotoSize `json:"thumbnail"`      // Optional. Thumbnail of the album cover to which the music file belongs
}

// This object represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileID       string     `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Thumbnail    *PhotoSize `json:"thumbnail"`      // Optional. Document thumbnail as defined by the sender
	FileName     string     `json:"file_name"`      // Optional. Original filename as defined by the sender
	MimeType     string     `json:"mime_type"`      // Optional. MIME type of the file as defined by the sender
	FileSize     int64      `json:"file_size"`      // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}

// This object represents a story.
type Story struct {
	Chat *Chat `json:"chat"` // Chat that posted the story
	ID   int64 `json:"id"`   // Unique identifier for the story in the chat
}

// This object represents a video file.
type Video struct {
	FileID         string      `json:"file_id"`         // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID   string      `json:"file_unique_id"`  // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	WIDth          int64       `json:"width"`           // Video width as defined by the sender
	Height         int64       `json:"height"`          // Video height as defined by the sender
	Duration       int64       `json:"duration"`        // Duration of the video in seconds as defined by the sender
	Thumbnail      *PhotoSize  `json:"thumbnail"`       // Optional. Video thumbnail
	Cover          []PhotoSize `json:"cover"`           // Optional. Available sizes of the cover of the video in the message
	StartTimestamp int64       `json:"start_timestamp"` // Optional. Timestamp in seconds from which the video will play in the message
	FileName       string      `json:"file_name"`       // Optional. Original filename as defined by the sender
	MimeType       string      `json:"mime_type"`       // Optional. MIME type of the file as defined by the sender
	FileSize       int64       `json:"file_size"`       // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}

// This object represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	FileID       string     `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Length       int64      `json:"length"`         // Video width and height (diameter of the video message) as defined by the sender
	Duration     int64      `json:"duration"`       // Duration of the video in seconds as defined by the sender
	Thumbnail    *PhotoSize `json:"thumbnail"`      // Optional. Video thumbnail
	FileSize     int64      `json:"file_size"`      // Optional. File size in bytes
}

// This object represents a voice note.
type Voice struct {
	FileID       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int64  `json:"duration"`       // Duration of the audio in seconds as defined by the sender
	MimeType     string `json:"mime_type"`      // Optional. MIME type of the file as defined by the sender
	FileSize     int64  `json:"file_size"`      // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}

// Describes the paid media added to a message.
type PaidMediaInfo struct {
	StarCount int64       `json:"star_count"` // The number of Telegram Stars that must be paid to buy access to the media
	PaIDMedia []PaidMedia `json:"paid_media"` // Information about the paid media
}

// This object describes paid media. Currently, it can be one of
// - PaidMediaPreview
// - PaidMediaPhoto
// - PaidMediaVideo
type PaidMedia interface{}

// The paid media isn't available before the payment.
type PaidMediaPreview struct {
	Type     string `json:"type"`     // Type of the paid media, always "preview"
	WIDth    int64  `json:"width"`    // Optional. Media width as defined by the sender
	Height   int64  `json:"height"`   // Optional. Media height as defined by the sender
	Duration int64  `json:"duration"` // Optional. Duration of the media in seconds as defined by the sender
}

// The paid media is a photo.
type PaidMediaPhoto struct {
	Type  string      `json:"type"`  // Type of the paid media, always "photo"
	Photo []PhotoSize `json:"photo"` // The photo
}

// The paid media is a video.
type PaidMediaVideo struct {
	Type  string `json:"type"`  // Type of the paid media, always "video"
	VIDeo *Video `json:"video"` // The video
}

// This object represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"` // Contact's phone number
	FirstName   string `json:"first_name"`   // Contact's first name
	LastName    string `json:"last_name"`    // Optional. Contact's last name
	UserID      int64  `json:"user_id"`      // Optional. Contact's user identifier in Telegram. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	Vcard       string `json:"vcard"`        // Optional. Additional data about the contact in the form of a vCard
}

// This object represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji"` // Emoji on which the dice throw animation is based
	Value int64  `json:"value"` // Value of the dice, 1-6 for "🎲", "🎯" and "🎳" base emoji, 1-5 for "🏀" and "⚽" base emoji, 1-64 for "🎰" base emoji
}

// This object contains information about one answer option in a poll.
type PollOption struct {
	Text         string          `json:"text"`          // Option text, 1-100 characters
	TextEntities []MessageEntity `json:"text_entities"` // Optional. Special entities that appear in the option text. Currently, only custom emoji entities are allowed in poll option texts
	VoterCount   int64           `json:"voter_count"`   // Number of users that voted for this option
}

// This object contains information about one answer option in a poll to be sent.
type InputPollOption struct {
	Text          string          `json:"text"`            // Option text, 1-100 characters
	TextParseMode string          `json:"text_parse_mode"` // Optional. Mode for parsing entities in the text. See formatting options for more details. Currently, only custom emoji entities are allowed
	TextEntities  []MessageEntity `json:"text_entities"`   // Optional. A JSON-serialized list of special entities that appear in the poll option text. It can be specified instead of text_parse_mode
}

// This object represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	PollID    string  `json:"poll_id"`    // Unique poll identifier
	VoterChat *Chat   `json:"voter_chat"` // Optional. The chat that changed the answer to the poll, if the voter is anonymous
	User      *User   `json:"user"`       // Optional. The user that changed the answer to the poll, if the voter isn't anonymous
	OptionIDs []int64 `json:"option_ids"` // 0-based identifiers of chosen answer options. May be empty if the vote was retracted.
}

// This object contains information about a poll.
type Poll struct {
	ID                    string          `json:"id"`                      // Unique poll identifier
	Question              string          `json:"question"`                // Poll question, 1-300 characters
	QuestionEntities      []MessageEntity `json:"question_entities"`       // Optional. Special entities that appear in the question. Currently, only custom emoji entities are allowed in poll questions
	Options               []PollOption    `json:"options"`                 // List of poll options
	TotalVoterCount       int64           `json:"total_voter_count"`       // Total number of users that voted in the poll
	IsClosed              bool            `json:"is_closed"`               // True, if the poll is closed
	IsAnonymous           bool            `json:"is_anonymous"`            // True, if the poll is anonymous
	Type                  string          `json:"type"`                    // Poll type, currently can be "regular" or "quiz"
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"` // True, if the poll allows multiple answers
	CorrectOptionID       int64           `json:"correct_option_id"`       // Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	Explanation           string          `json:"explanation"`             // Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	ExplanationEntities   []MessageEntity `json:"explanation_entities"`    // Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
	OpenPeriod            int64           `json:"open_period"`             // Optional. Amount of time in seconds the poll will be active after creation
	CloseDate             int64           `json:"close_date"`              // Optional. Point in time (Unix timestamp) when the poll will be automatically closed
}

// This object represents a point on the map.
type Location struct {
	Latitude             float64 `json:"latitude"`               // Latitude as defined by the sender
	Longitude            float64 `json:"longitude"`              // Longitude as defined by the sender
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int64   `json:"live_period"`            // Optional. Time relative to the message sending date, during which the location can be updated; in seconds. For active live locations only.
	Heading              int64   `json:"heading"`                // Optional. The direction in which user is moving, in degrees; 1-360. For active live locations only.
	ProximityAlertRadius int64   `json:"proximity_alert_radius"` // Optional. The maximum distance for proximity alerts about approaching another chat member, in meters. For sent live locations only.
}

// This object represents a venue.
type Venue struct {
	Location        *Location `json:"location"`          // Venue location. Can't be a live location
	Title           string    `json:"title"`             // Name of the venue
	Address         string    `json:"address"`           // Address of the venue
	FoursquareID    string    `json:"foursquare_id"`     // Optional. Foursquare identifier of the venue
	FoursquareType  string    `json:"foursquare_type"`   // Optional. Foursquare type of the venue. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	GooglePlaceID   string    `json:"google_place_id"`   // Optional. Google Places identifier of the venue
	GooglePlaceType string    `json:"google_place_type"` // Optional. Google Places type of the venue. (See supported types.)
}

// Describes data sent from a Web App to the bot.
type WebAppData struct {
	Data       string `json:"data"`        // The data. Be aware that a bad client can send arbitrary data in this field.
	ButtonText string `json:"button_text"` // Text of the web_app keyboard button from which the Web App was opened. Be aware that a bad client can send arbitrary data in this field.
}

// This object represents the content of a service message, sent whenever a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"` // User that triggered the alert
	Watcher  *User `json:"watcher"`  // User that set the alert
	Distance int64 `json:"distance"` // The distance between the users
}

// This object represents a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int64 `json:"message_auto_delete_time"` // New auto-delete time for messages in the chat; in seconds
}

// This object represents a service message about a user boosting a chat.
type ChatBoostAdded struct {
	BoostCount int64 `json:"boost_count"` // Number of boosts added by the user
}

// This object describes the way a background is filled based on the selected colors. Currently, it can be one of
// - BackgroundFillSolid
// - BackgroundFillGradient
// - BackgroundFillFreeformGradient
type BackgroundFill interface{}

// The background is filled using the selected color.
type BackgroundFillSolid struct {
	Type  string `json:"type"`  // Type of the background fill, always "solid"
	Color int64  `json:"color"` // The color of the background fill in the RGB24 format
}

// The background is a gradient fill.
type BackgroundFillGradient struct {
	Type          string `json:"type"`           // Type of the background fill, always "gradient"
	TopColor      int64  `json:"top_color"`      // Top color of the gradient in the RGB24 format
	BottomColor   int64  `json:"bottom_color"`   // Bottom color of the gradient in the RGB24 format
	RotationAngle int64  `json:"rotation_angle"` // Clockwise rotation angle of the background fill in degrees; 0-359
}

// The background is a freeform gradient that rotates after every message in the chat.
type BackgroundFillFreeformGradient struct {
	Type   string  `json:"type"`   // Type of the background fill, always "freeform_gradient"
	Colors []int64 `json:"colors"` // A list of the 3 or 4 base colors that are used to generate the freeform gradient in the RGB24 format
}

// This object describes the type of a background. Currently, it can be one of
// - BackgroundTypeFill
// - BackgroundTypeWallpaper
// - BackgroundTypePattern
// - BackgroundTypeChatTheme
type BackgroundType interface{}

// The background is automatically filled based on the selected colors.
type BackgroundTypeFill struct {
	Type             string          `json:"type"`               // Type of the background, always "fill"
	Fill             *BackgroundFill `json:"fill"`               // The background fill
	DarkThemeDimming int64           `json:"dark_theme_dimming"` // Dimming of the background in dark themes, as a percentage; 0-100
}

// The background is a wallpaper in the JPEG format.
type BackgroundTypeWallpaper struct {
	Type             string    `json:"type"`               // Type of the background, always "wallpaper"
	Document         *Document `json:"document"`           // Document with the wallpaper
	DarkThemeDimming int64     `json:"dark_theme_dimming"` // Dimming of the background in dark themes, as a percentage; 0-100
	IsBlurred        bool      `json:"is_blurred"`         // Optional. True, if the wallpaper is downscaled to fit in a 450x450 square and then box-blurred with radius 12
	IsMoving         bool      `json:"is_moving"`          // Optional. True, if the background moves slightly when the device is tilted
}

// The background is a .PNG or .TGV (gzipped subset of SVG with MIME type "application/x-tgwallpattern") pattern to be combined with the background fill chosen by the user.
type BackgroundTypePattern struct {
	Type       string          `json:"type"`        // Type of the background, always "pattern"
	Document   *Document       `json:"document"`    // Document with the pattern
	Fill       *BackgroundFill `json:"fill"`        // The background fill that is combined with the pattern
	Intensity  int64           `json:"intensity"`   // Intensity of the pattern when it is shown above the filled background; 0-100
	IsInverted bool            `json:"is_inverted"` // Optional. True, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only
	IsMoving   bool            `json:"is_moving"`   // Optional. True, if the background moves slightly when the device is tilted
}

// The background is taken directly from a built-in chat theme.
type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`       // Type of the background, always "chat_theme"
	ThemeName string `json:"theme_name"` // Name of the chat theme, which is usually an emoji
}

// This object represents a chat background.
type ChatBackground struct {
	Type *BackgroundType `json:"type"` // Type of the background
}

// This object represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
	Name              string `json:"name"`                 // Name of the topic
	IconColor         int64  `json:"icon_color"`           // Color of the topic icon in RGB format
	IconCustomEmojiID string `json:"icon_custom_emoji_id"` // Optional. Unique identifier of the custom emoji shown as the topic icon
}

// This object represents a service message about a forum topic closed in the chat. Currently holds no information.
type ForumTopicClosed interface{}

// This object represents a service message about an edited forum topic.
type ForumTopicEdited struct {
	Name              string `json:"name"`                 // Optional. New name of the topic, if it was edited
	IconCustomEmojiID string `json:"icon_custom_emoji_id"` // Optional. New identifier of the custom emoji shown as the topic icon, if it was edited; an empty string if the icon was removed
}

// This object represents a service message about a forum topic reopened in the chat. Currently holds no information.
type ForumTopicReopened interface{}

// This object represents a service message about General forum topic hidden in the chat. Currently holds no information.
type GeneralForumTopicHidden interface{}

// This object represents a service message about General forum topic unhidden in the chat. Currently holds no information.
type GeneralForumTopicUnhidden interface{}

// This object contains information about a user that was shared with the bot using a KeyboardButtonRequestUsers button.
type SharedUser struct {
	UserID    int64       `json:"user_id"`    // Identifier of the shared user. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so 64-bit integers or double-precision float types are safe for storing these identifiers. The bot may not have access to the user and could be unable to use this identifier, unless the user is already known to the bot by some other means.
	FirstName string      `json:"first_name"` // Optional. First name of the user, if the name was requested by the bot
	LastName  string      `json:"last_name"`  // Optional. Last name of the user, if the name was requested by the bot
	Username  string      `json:"username"`   // Optional. Username of the user, if the username was requested by the bot
	Photo     []PhotoSize `json:"photo"`      // Optional. Available sizes of the chat photo, if the photo was requested by the bot
}

// This object contains information about the users whose identifiers were shared with the bot using a KeyboardButtonRequestUsers button.
type UsersShared struct {
	RequestID int64        `json:"request_id"` // Identifier of the request
	Users     []SharedUser `json:"users"`      // Information about users shared with the bot.
}

// This object contains information about a chat that was shared with the bot using a KeyboardButtonRequestChat button.
type ChatShared struct {
	RequestID int64       `json:"request_id"` // Identifier of the request
	ChatID    ChatID      `json:"chat_id"`    // Identifier of the shared chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot may not have access to the chat and could be unable to use this identifier, unless the chat is already known to the bot by some other means.
	Title     string      `json:"title"`      // Optional. Title of the chat, if the title was requested by the bot.
	Username  string      `json:"username"`   // Optional. Username of the chat, if the username was requested by the bot and available.
	Photo     []PhotoSize `json:"photo"`      // Optional. Available sizes of the chat photo, if the photo was requested by the bot
}

// This object represents a service message about a user allowing a bot to write messages after adding it to the attachment menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess.
type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request"`         // Optional. True, if the access was granted after the user accepted an explicit request from a Web App sent by the method requestWriteAccess
	WebAppName         string `json:"web_app_name"`         // Optional. Name of the Web App, if the access was granted when the Web App was launched from a link
	FromAttachmentMenu bool   `json:"from_attachment_menu"` // Optional. True, if the access was granted when the bot was added to the attachment or side menu
}

// This object represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
	StartDate int64 `json:"start_date"` // Point in time (Unix timestamp) when the video chat is supposed to be started by a chat administrator
}

// This object represents a service message about a video chat started in the chat. Currently holds no information.
type VideoChatStarted interface{}

// This object represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	Duration int64 `json:"duration"` // Video chat duration in seconds
}

// This object represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"` // New members that were invited to the video chat
}

// This object represents a service message about the creation of a scheduled giveaway.
type GiveawayCreated struct {
	PrizeStarCount int64 `json:"prize_star_count"` // Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
}

// This object represents a message about a scheduled giveaway.
type Giveaway struct {
	Chats                         []Chat   `json:"chats"`                            // The list of chats which the user must join to participate in the giveaway
	WinnersSelectionDate          int64    `json:"winners_selection_date"`           // Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnerCount                   int64    `json:"winner_count"`                     // The number of users which are supposed to be selected as winners of the giveaway
	OnlyNewMembers                bool     `json:"only_new_members"`                 // Optional. True, if only users who join the chats after the giveaway started should be eligible to win
	HasPublicWinners              bool     `json:"has_public_winners"`               // Optional. True, if the list of giveaway winners will be visible to everyone
	PrizeDescription              string   `json:"prize_description"`                // Optional. Description of additional giveaway prize
	CountryCodes                  []string `json:"country_codes"`                    // Optional. A list of two-letter ISO 3166-1 alpha-2 country codes indicating the countries from which eligible users for the giveaway must come. If empty, then all users can participate in the giveaway. Users with a phone number that was bought on Fragment can always participate in giveaways.
	PrizeStarCount                int64    `json:"prize_star_count"`                 // Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PremiumSubscriptionMonthCount int64    `json:"premium_subscription_month_count"` // Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only
}

// This object represents a message about the completion of a giveaway with public winners.
type GiveawayWinners struct {
	Chat                          *Chat  `json:"chat"`                             // The chat that created the giveaway
	GiveawayMessageID             int64  `json:"giveaway_message_id"`              // Identifier of the message with the giveaway in the chat
	WinnersSelectionDate          int64  `json:"winners_selection_date"`           // Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnerCount                   int64  `json:"winner_count"`                     // Total number of winners in the giveaway
	Winners                       []User `json:"winners"`                          // List of up to 100 winners of the giveaway
	AdditionalChatCount           int64  `json:"additional_chat_count"`            // Optional. The number of other chats the user had to join in order to be eligible for the giveaway
	PrizeStarCount                int64  `json:"prize_star_count"`                 // Optional. The number of Telegram Stars that were split between giveaway winners; for Telegram Star giveaways only
	PremiumSubscriptionMonthCount int64  `json:"premium_subscription_month_count"` // Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only
	UnclaimedPrizeCount           int64  `json:"unclaimed_prize_count"`            // Optional. Number of undistributed prizes
	OnlyNewMembers                bool   `json:"only_new_members"`                 // Optional. True, if only users who had joined the chats after the giveaway started were eligible to win
	WasRefunded                   bool   `json:"was_refunded"`                     // Optional. True, if the giveaway was canceled because the payment for it was refunded
	PrizeDescription              string `json:"prize_description"`                // Optional. Description of additional giveaway prize
}

// This object represents a service message about the completion of a giveaway without public winners.
type GiveawayCompleted struct {
	WinnerCount         int64    `json:"winner_count"`          // Number of winners in the giveaway
	UnclaimedPrizeCount int64    `json:"unclaimed_prize_count"` // Optional. Number of undistributed prizes
	GiveawayMessage     *Message `json:"giveaway_message"`      // Optional. Message with the giveaway that was completed, if it wasn't deleted
	IsStarGiveaway      bool     `json:"is_star_giveaway"`      // Optional. True, if the giveaway is a Telegram Star giveaway. Otherwise, currently, the giveaway is a Telegram Premium giveaway.
}

// Describes the options used for link preview generation.
type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled"`        // Optional. True, if the link preview is disabled
	Url              string `json:"url"`                // Optional. URL to use for the link preview. If empty, then the first URL found in the message text will be used
	PreferSmallMedia bool   `json:"prefer_small_media"` // Optional. True, if the media in the link preview is supposed to be shrunk; ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	PreferLargeMedia bool   `json:"prefer_large_media"` // Optional. True, if the media in the link preview is supposed to be enlarged; ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	ShowAboveText    bool   `json:"show_above_text"`    // Optional. True, if the link preview must be shown above the message text; otherwise, the link preview will be shown below the message text
}

// This object represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int64         `json:"total_count"` // Total number of profile pictures the target user has
	Photos     [][]PhotoSize `json:"photos"`      // Requested profile pictures (in up to 4 sizes each)
}

// This object represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
type File struct {
	FileID       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int64  `json:"file_size"`      // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	FilePath     string `json:"file_path"`      // Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}

// Describes a Web App.
type WebAppInfo struct {
	Url string `json:"url"` // An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
}

// This object represents a custom keyboard with reply options (see Introduction to bots for details and examples). Not supported in channels and for messages sent on behalf of a Telegram Business account.
type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`                // Array of button rows, each represented by an Array of KeyboardButton objects
	IsPersistent          bool               `json:"is_persistent"`           // Optional. Requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard icon.
	ResizeKeyboard        bool               `json:"resize_keyboard"`         // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard       bool               `json:"one_time_keyboard"`       // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	InputFieldPlaceholder string             `json:"input_field_placeholder"` // Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters
	Selective             bool               `json:"selective"`               // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message. Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select the new language. Other users in the group don't see the keyboard.
}

// This object represents one button of the reply keyboard. At most one of the optional fields must be used to specify type of the button. For simple text buttons, String can be used instead of this object to specify the button text.
// Note: request_users and request_chat options will only work in Telegram versions released after 3 February, 2023. Older clients will display unsupported message.
type KeyboardButton struct {
	Text            string                      `json:"text"`             // Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users"`    // Optional. If specified, pressing the button will open a list of suitable users. Identifiers of selected users will be sent to the bot in a "users_shared" service message. Available in private chats only.
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat"`     // Optional. If specified, pressing the button will open a list of suitable chats. Tapping on a chat will send its identifier to the bot in a "chat_shared" service message. Available in private chats only.
	RequestContact  bool                        `json:"request_contact"`  // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only.
	RequestLocation bool                        `json:"request_location"` // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only.
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll"`     // Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only.
	WebApp          *WebAppInfo                 `json:"web_app"`          // Optional. If specified, the described Web App will be launched when the button is pressed. The Web App will be able to send a "web_app_data" service message. Available in private chats only.
}

// This object defines the criteria used to request suitable users. Information about the selected users will be shared with the bot when the corresponding button is pressed. More about requesting users: https://core.telegram.org/bots/features#chat-and-user-selection
type KeyboardButtonRequestUsers struct {
	RequestID       int64 `json:"request_id"`       // Signed 32-bit identifier of the request that will be received back in the UsersShared object. Must be unique within the message
	UserIsBot       bool  `json:"user_is_bot"`      // Optional. Pass True to request bots, pass False to request regular users. If not specified, no additional restrictions are applied.
	UserIsPremium   bool  `json:"user_is_premium"`  // Optional. Pass True to request premium users, pass False to request non-premium users. If not specified, no additional restrictions are applied.
	MaxQuantity     int64 `json:"max_quantity"`     // Optional. The maximum number of users to be selected; 1-10. Defaults to 1.
	RequestName     bool  `json:"request_name"`     // Optional. Pass True to request the users' first and last names
	RequestUsername bool  `json:"request_username"` // Optional. Pass True to request the users' usernames
	RequestPhoto    bool  `json:"request_photo"`    // Optional. Pass True to request the users' photos
}

// This object defines the criteria used to request a suitable chat. Information about the selected chat will be shared with the bot when the corresponding button is pressed. The bot will be granted requested rights in the chat if appropriate. More about requesting chats: https://core.telegram.org/bots/features#chat-and-user-selection.
type KeyboardButtonRequestChat struct {
	RequestID               int64                    `json:"request_id"`                // Signed 32-bit identifier of the request, which will be received back in the ChatShared object. Must be unique within the message
	ChatIsChannel           bool                     `json:"chat_is_channel"`           // Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	ChatIsForum             bool                     `json:"chat_is_forum"`             // Optional. Pass True to request a forum supergroup, pass False to request a non-forum chat. If not specified, no additional restrictions are applied.
	ChatHasUsername         bool                     `json:"chat_has_username"`         // Optional. Pass True to request a supergroup or a channel with a username, pass False to request a chat without a username. If not specified, no additional restrictions are applied.
	ChatIsCreated           bool                     `json:"chat_is_created"`           // Optional. Pass True to request a chat owned by the user. Otherwise, no additional restrictions are applied.
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights"` // Optional. A JSON-serialized object listing the required administrator rights of the user in the chat. The rights must be a superset of bot_administrator_rights. If not specified, no additional restrictions are applied.
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights"`  // Optional. A JSON-serialized object listing the required administrator rights of the bot in the chat. The rights must be a subset of user_administrator_rights. If not specified, no additional restrictions are applied.
	BotIsMember             bool                     `json:"bot_is_member"`             // Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
	RequestTitle            bool                     `json:"request_title"`             // Optional. Pass True to request the chat's title
	RequestUsername         bool                     `json:"request_username"`          // Optional. Pass True to request the chat's username
	RequestPhoto            bool                     `json:"request_photo"`             // Optional. Pass True to request the chat's photo
}

// This object represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	Type string `json:"type"` // Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
}

// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup). Not supported in channels and for messages sent on behalf of a Telegram Business account.
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"` // Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	Selective      bool `json:"selective"`       // Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message. Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
}

// This object represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

// This object represents one button of an inline keyboard. Exactly one of the optional fields must be used to specify type of the button.
type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`                             // Label text on the button
	Url                          string                       `json:"url"`                              // Optional. HTTP or tg:// URL to be opened when the button is pressed. Links tg://user?id=<user_id> can be used to mention a user by their identifier without using a username, if this is allowed by their privacy settings.
	CallbackData                 string                       `json:"callback_data"`                    // Optional. Data to be sent in a callback query to the bot when the button is pressed, 1-64 bytes
	WebApp                       *WebAppInfo                  `json:"web_app"`                          // Optional. Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery. Available only in private chats between a user and the bot. Not supported for messages sent on behalf of a Telegram Business account.
	LoginUrl                     *LoginUrl                    `json:"login_url"`                        // Optional. An HTTPS URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	SwitchInlineQuery            string                       `json:"switch_inline_query"`              // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. May be empty, in which case just the bot's username will be inserted. Not supported for messages sent on behalf of a Telegram Business account.
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat"` // Optional. If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. May be empty, in which case only the bot's username will be inserted. This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options. Not supported in channels and for messages sent on behalf of a Telegram Business account.
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat"`  // Optional. If set, pressing the button will prompt the user to select one of their chats of the specified type, open that chat and insert the bot's username and the specified inline query in the input field. Not supported for messages sent on behalf of a Telegram Business account.
	CopyText                     *CopyTextButton              `json:"copy_text"`                        // Optional. Description of the button that copies the specified text to the clipboard.
	CallbackGame                 CallbackGame                 `json:"callback_game"`                    // Optional. Description of the game that will be launched when the user presses the button. NOTE: This type of button must always be the first button in the first row.
	Pay                          bool                         `json:"pay"`                              // Optional. Specify True, to send a Pay button. Substrings "⭐" and "XTR" in the buttons's text will be replaced with a Telegram Star icon. NOTE: This type of button must always be the first button in the first row and can only be used in invoice messages.
}

// This object represents a parameter of the inline keyboard button used to automatically authorize a user. Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram. All the user needs to do is tap/click a button and confirm that they want to log in:
// Telegram apps support these buttons as of version 5.7.
type LoginUrl struct {
	Url                string `json:"url"`                  // An HTTPS URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data. NOTE: You must always check the hash of the received data to verify the authentication and the integrity of the data as described in Checking authorization.
	ForwardText        string `json:"forward_text"`         // Optional. New text of the button in forwarded messages.
	BotUsername        string `json:"bot_username"`         // Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
	RequestWriteAccess bool   `json:"request_write_access"` // Optional. Pass True to request the permission for your bot to send messages to the user.
}

// This object represents an inline button that switches the current user to inline mode in a chosen chat, with an optional default inline query.
type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query"`               // Optional. The default inline query to be inserted in the input field. If left empty, only the bot's username will be inserted
	AllowUserChats    bool   `json:"allow_user_chats"`    // Optional. True, if private chats with users can be chosen
	AllowBotChats     bool   `json:"allow_bot_chats"`     // Optional. True, if private chats with bots can be chosen
	AllowGroupChats   bool   `json:"allow_group_chats"`   // Optional. True, if group and supergroup chats can be chosen
	AllowChannelChats bool   `json:"allow_channel_chats"` // Optional. True, if channel chats can be chosen
}

// This object represents an inline keyboard button that copies specified text to the clipboard.
type CopyTextButton struct {
	Text string `json:"text"` // The text to be copied to the clipboard; 1-256 characters
}

// This object represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	ID              string                    `json:"id"`                // Unique identifier for this query
	From            *User                     `json:"from"`              // Sender
	Message         *MaybeInaccessibleMessage `json:"message"`           // Optional. Message sent by the bot with the callback button that originated the query
	InlineMessageID string                    `json:"inline_message_id"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	ChatInstance    string                    `json:"chat_instance"`     // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	Data            string                    `json:"data"`              // Optional. Data associated with the callback button. Be aware that the message originated the query can contain no callback buttons with this data.
	GameShortName   string                    `json:"game_short_name"`   // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}

// Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode. Not supported in channels and for messages sent on behalf of a Telegram Business account.
type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`             // Shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'
	InputFieldPlaceholder string `json:"input_field_placeholder"` // Optional. The placeholder to be shown in the input field when the reply is active; 1-64 characters
	Selective             bool   `json:"selective"`               // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message.
}

// This object represents a chat photo.
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`        // File identifier of small (160x160) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	SmallFileUniqueID string `json:"small_file_unique_id"` // Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	BigFileID         string `json:"big_file_id"`          // File identifier of big (640x640) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileUniqueID   string `json:"big_file_unique_id"`   // Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
}

// Represents an invite link for a chat.
type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`                // The invite link. If the link was created by another chat administrator, then the second part of the link will be replaced with "...".
	Creator                 *User  `json:"creator"`                    // Creator of the link
	CreatesJoinRequest      bool   `json:"creates_join_request"`       // True, if users joining the chat via the link need to be approved by chat administrators
	IsPrimary               bool   `json:"is_primary"`                 // True, if the link is primary
	IsRevoked               bool   `json:"is_revoked"`                 // True, if the link is revoked
	Name                    string `json:"name"`                       // Optional. Invite link name
	ExpireDate              int64  `json:"expire_date"`                // Optional. Point in time (Unix timestamp) when the link will expire or has been expired
	MemberLimit             int64  `json:"member_limit"`               // Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	PendingJoinRequestCount int64  `json:"pending_join_request_count"` // Optional. Number of pending join requests created using this link
	SubscriptionPeriod      int64  `json:"subscription_period"`        // Optional. The number of seconds the subscription will be active for before the next payment
	SubscriptionPrice       int64  `json:"subscription_price"`         // Optional. The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat using the link
}

// Represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {
	IsAnonymous         bool `json:"is_anonymous"`           // True, if the user's presence in the chat is hidden
	CanManageChat       bool `json:"can_manage_chat"`        // True, if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege.
	CanDeleteMessages   bool `json:"can_delete_messages"`    // True, if the administrator can delete messages of other users
	CanManageVIDeoChats bool `json:"can_manage_video_chats"` // True, if the administrator can manage video chats
	CanRestrictMembers  bool `json:"can_restrict_members"`   // True, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanPromoteMembers   bool `json:"can_promote_members"`    // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo       bool `json:"can_change_info"`        // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers      bool `json:"can_invite_users"`       // True, if the user is allowed to invite new users to the chat
	CanPostStories      bool `json:"can_post_stories"`       // True, if the administrator can post stories to the chat
	CanEditStories      bool `json:"can_edit_stories"`       // True, if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanDeleteStories    bool `json:"can_delete_stories"`     // True, if the administrator can delete stories posted by other users
	CanPostMessages     bool `json:"can_post_messages"`      // Optional. True, if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanEditMessages     bool `json:"can_edit_messages"`      // Optional. True, if the administrator can edit messages of other users and can pin messages; for channels only
	CanPinMessages      bool `json:"can_pin_messages"`       // Optional. True, if the user is allowed to pin messages; for groups and supergroups only
	CanManageTopics     bool `json:"can_manage_topics"`      // Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
}

// This object represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	Chat                    *Chat           `json:"chat"`                        // Chat the user belongs to
	From                    *User           `json:"from"`                        // Performer of the action, which resulted in the change
	Date                    int64           `json:"date"`                        // Date the change was done in Unix time
	OldChatMember           *ChatMember     `json:"old_chat_member"`             // Previous information about the chat member
	NewChatMember           *ChatMember     `json:"new_chat_member"`             // New information about the chat member
	InviteLink              *ChatInviteLink `json:"invite_link"`                 // Optional. Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
	ViaJoinRequest          bool            `json:"via_join_request"`            // Optional. True, if the user joined the chat after sending a direct join request without using an invite link and being approved by an administrator
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link"` // Optional. True, if the user joined the chat via a chat folder invite link
}

// This object contains information about one member of a chat. Currently, the following 6 types of chat members are supported:
// - ChatMemberOwner
// - ChatMemberAdministrator
// - ChatMemberMember
// - ChatMemberRestricted
// - ChatMemberLeft
// - ChatMemberBanned
type ChatMember interface{}

// Represents a chat member that owns the chat and has all administrator privileges.
type ChatMemberOwner struct {
	Status      string `json:"status"`       // The member's status in the chat, always "creator"
	User        *User  `json:"user"`         // Information about the user
	IsAnonymous bool   `json:"is_anonymous"` // True, if the user's presence in the chat is hidden
	CustomTitle string `json:"custom_title"` // Optional. Custom title for this user
}

// Represents a chat member that has some additional privileges.
type ChatMemberAdministrator struct {
	Status              string `json:"status"`                 // The member's status in the chat, always "administrator"
	User                *User  `json:"user"`                   // Information about the user
	CanBeEdited         bool   `json:"can_be_edited"`          // True, if the bot is allowed to edit administrator privileges of that user
	IsAnonymous         bool   `json:"is_anonymous"`           // True, if the user's presence in the chat is hidden
	CanManageChat       bool   `json:"can_manage_chat"`        // True, if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege.
	CanDeleteMessages   bool   `json:"can_delete_messages"`    // True, if the administrator can delete messages of other users
	CanManageVIDeoChats bool   `json:"can_manage_video_chats"` // True, if the administrator can manage video chats
	CanRestrictMembers  bool   `json:"can_restrict_members"`   // True, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanPromoteMembers   bool   `json:"can_promote_members"`    // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo       bool   `json:"can_change_info"`        // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers      bool   `json:"can_invite_users"`       // True, if the user is allowed to invite new users to the chat
	CanPostStories      bool   `json:"can_post_stories"`       // True, if the administrator can post stories to the chat
	CanEditStories      bool   `json:"can_edit_stories"`       // True, if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanDeleteStories    bool   `json:"can_delete_stories"`     // True, if the administrator can delete stories posted by other users
	CanPostMessages     bool   `json:"can_post_messages"`      // Optional. True, if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanEditMessages     bool   `json:"can_edit_messages"`      // Optional. True, if the administrator can edit messages of other users and can pin messages; for channels only
	CanPinMessages      bool   `json:"can_pin_messages"`       // Optional. True, if the user is allowed to pin messages; for groups and supergroups only
	CanManageTopics     bool   `json:"can_manage_topics"`      // Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
	CustomTitle         string `json:"custom_title"`           // Optional. Custom title for this user
}

// Represents a chat member that has no additional privileges or restrictions.
type ChatMemberMember struct {
	Status    string `json:"status"`     // The member's status in the chat, always "member"
	User      *User  `json:"user"`       // Information about the user
	UntilDate int64  `json:"until_date"` // Optional. Date when the user's subscription will expire; Unix time
}

// Represents a chat member that is under certain restrictions in the chat. Supergroups only.
type ChatMemberRestricted struct {
	Status                string `json:"status"`                    // The member's status in the chat, always "restricted"
	User                  *User  `json:"user"`                      // Information about the user
	IsMember              bool   `json:"is_member"`                 // True, if the user is a member of the chat at the moment of the request
	CanSendMessages       bool   `json:"can_send_messages"`         // True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues
	CanSendAudios         bool   `json:"can_send_audios"`           // True, if the user is allowed to send audios
	CanSendDocuments      bool   `json:"can_send_documents"`        // True, if the user is allowed to send documents
	CanSendPhotos         bool   `json:"can_send_photos"`           // True, if the user is allowed to send photos
	CanSendVIDeos         bool   `json:"can_send_videos"`           // True, if the user is allowed to send videos
	CanSendVIDeoNotes     bool   `json:"can_send_video_notes"`      // True, if the user is allowed to send video notes
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes"`      // True, if the user is allowed to send voice notes
	CanSendPolls          bool   `json:"can_send_polls"`            // True, if the user is allowed to send polls
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`   // True, if the user is allowed to send animations, games, stickers and use inline bots
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"` // True, if the user is allowed to add web page previews to their messages
	CanChangeInfo         bool   `json:"can_change_info"`           // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers        bool   `json:"can_invite_users"`          // True, if the user is allowed to invite new users to the chat
	CanPinMessages        bool   `json:"can_pin_messages"`          // True, if the user is allowed to pin messages
	CanManageTopics       bool   `json:"can_manage_topics"`         // True, if the user is allowed to create forum topics
	UntilDate             int64  `json:"until_date"`                // Date when restrictions will be lifted for this user; Unix time. If 0, then the user is restricted forever
}

// Represents a chat member that isn't currently a member of the chat, but may join it themselves.
type ChatMemberLeft struct {
	Status string `json:"status"` // The member's status in the chat, always "left"
	User   *User  `json:"user"`   // Information about the user
}

// Represents a chat member that was banned in the chat and can't return to the chat or view chat messages.
type ChatMemberBanned struct {
	Status    string `json:"status"`     // The member's status in the chat, always "kicked"
	User      *User  `json:"user"`       // Information about the user
	UntilDate int64  `json:"until_date"` // Date when restrictions will be lifted for this user; Unix time. If 0, then the user is banned forever
}

// Represents a join request sent to a chat.
type ChatJoinRequest struct {
	Chat       *Chat           `json:"chat"`         // Chat to which the request was sent
	From       *User           `json:"from"`         // User that sent the join request
	UserChatID ChatID          `json:"user_chat_id"` // Identifier of a private chat with the user who sent the join request. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot can use this identifier for 5 minutes to send messages until the join request is processed, assuming no other administrator contacted the user.
	Date       int64           `json:"date"`         // Date the request was sent in Unix time
	Bio        string          `json:"bio"`          // Optional. Bio of the user.
	InviteLink *ChatInviteLink `json:"invite_link"`  // Optional. Chat invite link that was used by the user to send the join request
}

// Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`         // Optional. True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues
	CanSendAudios         bool `json:"can_send_audios"`           // Optional. True, if the user is allowed to send audios
	CanSendDocuments      bool `json:"can_send_documents"`        // Optional. True, if the user is allowed to send documents
	CanSendPhotos         bool `json:"can_send_photos"`           // Optional. True, if the user is allowed to send photos
	CanSendVIDeos         bool `json:"can_send_videos"`           // Optional. True, if the user is allowed to send videos
	CanSendVIDeoNotes     bool `json:"can_send_video_notes"`      // Optional. True, if the user is allowed to send video notes
	CanSendVoiceNotes     bool `json:"can_send_voice_notes"`      // Optional. True, if the user is allowed to send voice notes
	CanSendPolls          bool `json:"can_send_polls"`            // Optional. True, if the user is allowed to send polls
	CanSendOtherMessages  bool `json:"can_send_other_messages"`   // Optional. True, if the user is allowed to send animations, games, stickers and use inline bots
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"` // Optional. True, if the user is allowed to add web page previews to their messages
	CanChangeInfo         bool `json:"can_change_info"`           // Optional. True, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups
	CanInviteUsers        bool `json:"can_invite_users"`          // Optional. True, if the user is allowed to invite new users to the chat
	CanPinMessages        bool `json:"can_pin_messages"`          // Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
	CanManageTopics       bool `json:"can_manage_topics"`         // Optional. True, if the user is allowed to create forum topics. If omitted defaults to the value of can_pin_messages
}

// Describes the birthdate of a user.
type Birthdate struct {
	Day   int64 `json:"day"`   // Day of the user's birth; 1-31
	Month int64 `json:"month"` // Month of the user's birth; 1-12
	Year  int64 `json:"year"`  // Optional. Year of the user's birth
}

// Contains information about the start page settings of a Telegram Business account.
type BusinessIntro struct {
	Title   string   `json:"title"`   // Optional. Title text of the business intro
	Message string   `json:"message"` // Optional. Message text of the business intro
	Sticker *Sticker `json:"sticker"` // Optional. Sticker of the business intro
}

// Contains information about the location of a Telegram Business account.
type BusinessLocation struct {
	Address  string    `json:"address"`  // Address of the business
	Location *Location `json:"location"` // Optional. Location of the business
}

// Describes an interval of time during which a business is open.
type BusinessOpeningHoursInterval struct {
	OpeningMinute int64 `json:"opening_minute"` // The minute's sequence number in a week, starting on Monday, marking the start of the time interval during which the business is open; 0 - 7 * 24 * 60
	ClosingMinute int64 `json:"closing_minute"` // The minute's sequence number in a week, starting on Monday, marking the end of the time interval during which the business is open; 0 - 8 * 24 * 60
}

// Describes the opening hours of a business.
type BusinessOpeningHours struct {
	TimeZoneName string                         `json:"time_zone_name"` // Unique name of the time zone for which the opening hours are defined
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`  // List of time intervals describing business opening hours
}

// Represents a location to which a chat is connected.
type ChatLocation struct {
	Location *Location `json:"location"` // The location to which the supergroup is connected. Can't be a live location.
	Address  string    `json:"address"`  // Location address; 1-64 characters, as defined by the chat owner
}

// This object describes the type of a reaction. Currently, it can be one of
// - ReactionTypeEmoji
// - ReactionTypeCustomEmoji
// - ReactionTypePaid
type ReactionType interface{}

// The reaction is based on an emoji.
type ReactionTypeEmoji struct {
	Type  string `json:"type"`  // Type of the reaction, always "emoji"
	Emoji string `json:"emoji"` // Reaction emoji. Currently, it can be one of "👍", "👎", "❤", "🔥", "🥰", "👏", "😁", "🤔", "🤯", "😱", "🤬", "😢", "🎉", "🤩", "🤮", "💩", "🙏", "👌", "🕊", "🤡", "🥱", "🥴", "😍", "🐳", "❤‍🔥", "🌚", "🌭", "💯", "🤣", "⚡", "🍌", "🏆", "💔", "🤨", "😐", "🍓", "🍾", "💋", "🖕", "😈", "😴", "😭", "🤓", "👻", "👨‍💻", "👀", "🎃", "🙈", "😇", "😨", "🤝", "✍", "🤗", "🫡", "🎅", "🎄", "☃", "💅", "🤪", "🗿", "🆒", "💘", "🙉", "🦄", "😘", "💊", "🙊", "😎", "👾", "🤷‍♂", "🤷", "🤷‍♀", "😡"
}

// The reaction is based on a custom emoji.
type ReactionTypeCustomEmoji struct {
	Type          string `json:"type"`            // Type of the reaction, always "custom_emoji"
	CustomEmojiID string `json:"custom_emoji_id"` // Custom emoji identifier
}

// The reaction is paid.
type ReactionTypePaid struct {
	Type string `json:"type"` // Type of the reaction, always "paid"
}

// Represents a reaction added to a message along with the number of times it was added.
type ReactionCount struct {
	Type       *ReactionType `json:"type"`        // Type of the reaction
	TotalCount int64         `json:"total_count"` // Number of times the reaction was added
}

// This object represents a change of a reaction on a message performed by a user.
type MessageReactionUpdated struct {
	Chat        *Chat          `json:"chat"`         // The chat containing the message the user reacted to
	MessageID   int64          `json:"message_id"`   // Unique identifier of the message inside the chat
	User        *User          `json:"user"`         // Optional. The user that changed the reaction, if the user isn't anonymous
	ActorChat   *Chat          `json:"actor_chat"`   // Optional. The chat on behalf of which the reaction was changed, if the user is anonymous
	Date        int64          `json:"date"`         // Date of the change in Unix time
	OldReaction []ReactionType `json:"old_reaction"` // Previous list of reaction types that were set by the user
	NewReaction []ReactionType `json:"new_reaction"` // New list of reaction types that have been set by the user
}

// This object represents reaction changes on a message with anonymous reactions.
type MessageReactionCountUpdated struct {
	Chat      *Chat           `json:"chat"`       // The chat containing the message
	MessageID int64           `json:"message_id"` // Unique message identifier inside the chat
	Date      int64           `json:"date"`       // Date of the change in Unix time
	Reactions []ReactionCount `json:"reactions"`  // List of reactions that are present on the message
}

// This object represents a forum topic.
type ForumTopic struct {
	MessageThreadID   int64  `json:"message_thread_id"`    // Unique identifier of the forum topic
	Name              string `json:"name"`                 // Name of the topic
	IconColor         int64  `json:"icon_color"`           // Color of the topic icon in RGB format
	IconCustomEmojiID string `json:"icon_custom_emoji_id"` // Optional. Unique identifier of the custom emoji shown as the topic icon
}

// This object represents a bot command.
type BotCommand struct {
	Command     string `json:"command"`     // Text of the command; 1-32 characters. Can contain only lowercase English letters, digits and underscores.
	Description string `json:"description"` // Description of the command; 1-256 characters.
}

// This object represents the scope to which bot commands are applied. Currently, the following 7 scopes are supported:
// - BotCommandScopeDefault
// - BotCommandScopeAllPrivateChats
// - BotCommandScopeAllGroupChats
// - BotCommandScopeAllChatAdministrators
// - BotCommandScopeChat
// - BotCommandScopeChatAdministrators
// - BotCommandScopeChatMember
type BotCommandScope interface{}

// Represents the default scope of bot commands. Default commands are used if no commands with a narrower scope are specified for the user.
type BotCommandScopeDefault struct {
	Type string `json:"type"` // Scope type, must be default
}

// Represents the scope of bot commands, covering all private chats.
type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"` // Scope type, must be all_private_chats
}

// Represents the scope of bot commands, covering all group and supergroup chats.
type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"` // Scope type, must be all_group_chats
}

// Represents the scope of bot commands, covering all group and supergroup chat administrators.
type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type"` // Scope type, must be all_chat_administrators
}

// Represents the scope of bot commands, covering a specific chat.
type BotCommandScopeChat struct {
	Type   string `json:"type"`    // Scope type, must be chat
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// Represents the scope of bot commands, covering all administrators of a specific group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type"`    // Scope type, must be chat_administrators
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// Represents the scope of bot commands, covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
	Type   string `json:"type"`    // Scope type, must be chat_member
	ChatID ChatID `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserID int64  `json:"user_id"` // Unique identifier of the target user
}

// This object represents the bot's name.
type BotName struct {
	Name string `json:"name"` // The bot's name
}

// This object represents the bot's description.
type BotDescription struct {
	Description string `json:"description"` // The bot's description
}

// This object represents the bot's short description.
type BotShortDescription struct {
	ShortDescription string `json:"short_description"` // The bot's short description
}

// This object describes the bot's menu button in a private chat. It should be one of
// - MenuButtonCommands
// - MenuButtonWebApp
// - MenuButtonDefault
// If a menu button other than MenuButtonDefault is set for a private chat, then it is applied in the chat. Otherwise the default menu button is applied. By default, the menu button opens the list of bot commands.
type MenuButton interface{}

// Represents a menu button, which opens the bot's list of commands.
type MenuButtonCommands struct {
	Type string `json:"type"` // Type of the button, must be commands
}

// Represents a menu button, which launches a Web App.
type MenuButtonWebApp struct {
	Type   string      `json:"type"`    // Type of the button, must be web_app
	Text   string      `json:"text"`    // Text on the button
	WebApp *WebAppInfo `json:"web_app"` // Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery. Alternatively, a t.me link to a Web App of the bot can be specified in the object instead of the Web App's URL, in which case the Web App will be opened as if the user pressed the link.
}

// Describes that no specific value for the menu button was set.
type MenuButtonDefault struct {
	Type string `json:"type"` // Type of the button, must be default
}

// This object describes the source of a chat boost. It can be one of
// - ChatBoostSourcePremium
// - ChatBoostSourceGiftCode
// - ChatBoostSourceGiveaway
type ChatBoostSource interface{}

// The boost was obtained by subscribing to Telegram Premium or by gifting a Telegram Premium subscription to another user.
type ChatBoostSourcePremium struct {
	Source string `json:"source"` // Source of the boost, always "premium"
	User   *User  `json:"user"`   // User that boosted the chat
}

// The boost was obtained by the creation of Telegram Premium gift codes to boost a chat. Each such code boosts the chat 4 times for the duration of the corresponding Telegram Premium subscription.
type ChatBoostSourceGiftCode struct {
	Source string `json:"source"` // Source of the boost, always "gift_code"
	User   *User  `json:"user"`   // User for which the gift code was created
}

// The boost was obtained by the creation of a Telegram Premium or a Telegram Star giveaway. This boosts the chat 4 times for the duration of the corresponding Telegram Premium subscription for Telegram Premium giveaways and prize_star_count / 500 times for one year for Telegram Star giveaways.
type ChatBoostSourceGiveaway struct {
	Source            string `json:"source"`              // Source of the boost, always "giveaway"
	GiveawayMessageID int64  `json:"giveaway_message_id"` // Identifier of a message in the chat with the giveaway; the message could have been deleted already. May be 0 if the message isn't sent yet.
	User              *User  `json:"user"`                // Optional. User that won the prize in the giveaway if any; for Telegram Premium giveaways only
	PrizeStarCount    int64  `json:"prize_star_count"`    // Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	IsUnclaimed       bool   `json:"is_unclaimed"`        // Optional. True, if the giveaway was completed, but there was no user to win the prize
}

// This object contains information about a chat boost.
type ChatBoost struct {
	BoostID        string           `json:"boost_id"`        // Unique identifier of the boost
	AddDate        int64            `json:"add_date"`        // Point in time (Unix timestamp) when the chat was boosted
	ExpirationDate int64            `json:"expiration_date"` // Point in time (Unix timestamp) when the boost will automatically expire, unless the booster's Telegram Premium subscription is prolonged
	Source         *ChatBoostSource `json:"source"`          // Source of the added boost
}

// This object represents a boost added to a chat or changed.
type ChatBoostUpdated struct {
	Chat  *Chat      `json:"chat"`  // Chat which was boosted
	Boost *ChatBoost `json:"boost"` // Information about the chat boost
}

// This object represents a boost removed from a chat.
type ChatBoostRemoved struct {
	Chat       *Chat            `json:"chat"`        // Chat which was boosted
	BoostID    string           `json:"boost_id"`    // Unique identifier of the boost
	RemoveDate int64            `json:"remove_date"` // Point in time (Unix timestamp) when the boost was removed
	Source     *ChatBoostSource `json:"source"`      // Source of the removed boost
}

// This object represents a list of boosts added to a chat by a user.
type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"` // The list of boosts added to the chat by the user
}

// Describes the connection of the bot with a business account.
type BusinessConnection struct {
	ID         string `json:"id"`           // Unique identifier of the business connection
	User       *User  `json:"user"`         // Business account user that created the business connection
	UserChatID ChatID `json:"user_chat_id"` // Identifier of a private chat with the user who created the business connection. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	Date       int64  `json:"date"`         // Date the connection was established in Unix time
	CanReply   bool   `json:"can_reply"`    // True, if the bot can act on behalf of the business account in chats that were active in the last 24 hours
	IsEnabled  bool   `json:"is_enabled"`   // True, if the connection is active
}

// This object is received when messages are deleted from a connected business account.
type BusinessMessagesDeleted struct {
	BusinessConnectionID string  `json:"business_connection_id"` // Unique identifier of the business connection
	Chat                 *Chat   `json:"chat"`                   // Information about a chat in the business account. The bot may not have access to the chat or the corresponding user.
	MessageIDs           []int64 `json:"message_ids"`            // The list of identifiers of deleted messages in the chat of the business account
}

// Describes why a request was unsuccessful.
type ResponseParameters struct {
	MigrateToChatID ChatID `json:"migrate_to_chat_id"` // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	RetryAfter      int64  `json:"retry_after"`        // Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
}

// This object represents the content of a media message to be sent. It should be one of
// - InputMediaAnimation
// - InputMediaDocument
// - InputMediaAudio
// - InputMediaPhoto
// - InputMediaVideo
type InputMedia interface{}

// Represents a photo to be sent.
type InputMediaPhoto struct {
	Type                  string          `json:"type"`                     // Type of the result, must be photo
	Media                 string          `json:"media"`                    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption               string          `json:"caption"`                  // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode             string          `json:"parse_mode"`               // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities       []MessageEntity `json:"caption_entities"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media"` // Optional. Pass True, if the caption must be shown above the message media
	HasSpoiler            bool            `json:"has_spoiler"`              // Optional. Pass True if the photo needs to be covered with a spoiler animation
}

// Represents a video to be sent.
type InputMediaVideo struct {
	Type                  string          `json:"type"`                     // Type of the result, must be video
	Media                 string          `json:"media"`                    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail             string          `json:"thumbnail"`                // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Cover                 string          `json:"cover"`                    // Optional. Cover for the video in the message. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	StartTimestamp        int64           `json:"start_timestamp"`          // Optional. Start timestamp for the video in the message
	Caption               string          `json:"caption"`                  // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode             string          `json:"parse_mode"`               // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities       []MessageEntity `json:"caption_entities"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media"` // Optional. Pass True, if the caption must be shown above the message media
	WIDth                 int64           `json:"width"`                    // Optional. Video width
	Height                int64           `json:"height"`                   // Optional. Video height
	Duration              int64           `json:"duration"`                 // Optional. Video duration in seconds
	SupportsStreaming     bool            `json:"supports_streaming"`       // Optional. Pass True if the uploaded video is suitable for streaming
	HasSpoiler            bool            `json:"has_spoiler"`              // Optional. Pass True if the video needs to be covered with a spoiler animation
}

// Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	Type                  string          `json:"type"`                     // Type of the result, must be animation
	Media                 string          `json:"media"`                    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail             string          `json:"thumbnail"`                // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption               string          `json:"caption"`                  // Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	ParseMode             string          `json:"parse_mode"`               // Optional. Mode for parsing entities in the animation caption. See formatting options for more details.
	CaptionEntities       []MessageEntity `json:"caption_entities"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media"` // Optional. Pass True, if the caption must be shown above the message media
	WIDth                 int64           `json:"width"`                    // Optional. Animation width
	Height                int64           `json:"height"`                   // Optional. Animation height
	Duration              int64           `json:"duration"`                 // Optional. Animation duration in seconds
	HasSpoiler            bool            `json:"has_spoiler"`              // Optional. Pass True if the animation needs to be covered with a spoiler animation
}

// Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	Type            string          `json:"type"`             // Type of the result, must be audio
	Media           string          `json:"media"`            // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail       string          `json:"thumbnail"`        // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption         string          `json:"caption"`          // Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	ParseMode       string          `json:"parse_mode"`       // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities []MessageEntity `json:"caption_entities"` // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration        int64           `json:"duration"`         // Optional. Duration of the audio in seconds
	Performer       string          `json:"performer"`        // Optional. Performer of the audio
	Title           string          `json:"title"`            // Optional. Title of the audio
}

// Represents a general file to be sent.
type InputMediaDocument struct {
	Type                        string          `json:"type"`                           // Type of the result, must be document
	Media                       string          `json:"media"`                          // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail                   string          `json:"thumbnail"`                      // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Caption                     string          `json:"caption"`                        // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode                   string          `json:"parse_mode"`                     // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities             []MessageEntity `json:"caption_entities"`               // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableContentTypeDetection bool            `json:"disable_content_type_detection"` // Optional. Disables automatic server-side content type detection for files uploaded using multipart/form-data. Always True, if the document is sent as part of an album.
}

// This object represents the contents of a file to be uploaded. Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
type InputFile interface{}

// This object describes the paid media to be sent. Currently, it can be one of
// - InputPaidMediaPhoto
// - InputPaidMediaVideo
type InputPaidMedia interface{}

// The paid media to send is a photo.
type InputPaidMediaPhoto struct {
	Type  string `json:"type"`  // Type of the media, must be photo
	Media string `json:"media"` // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
}

// The paid media to send is a video.
type InputPaidMediaVideo struct {
	Type              string `json:"type"`               // Type of the media, must be video
	Media             string `json:"media"`              // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail         string `json:"thumbnail"`          // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Cover             string `json:"cover"`              // Optional. Cover for the video in the message. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	StartTimestamp    int64  `json:"start_timestamp"`    // Optional. Start timestamp for the video in the message
	WIDth             int64  `json:"width"`              // Optional. Video width
	Height            int64  `json:"height"`             // Optional. Video height
	Duration          int64  `json:"duration"`           // Optional. Video duration in seconds
	SupportsStreaming bool   `json:"supports_streaming"` // Optional. Pass True if the uploaded video is suitable for streaming
}

// This object represents a sticker.
type Sticker struct {
	FileID           string        `json:"file_id"`           // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID     string        `json:"file_unique_id"`    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Type             string        `json:"type"`              // Type of the sticker, currently one of "regular", "mask", "custom_emoji". The type of the sticker is independent from its format, which is determined by the fields is_animated and is_video.
	WIDth            int64         `json:"width"`             // Sticker width
	Height           int64         `json:"height"`            // Sticker height
	IsAnimated       bool          `json:"is_animated"`       // True, if the sticker is animated
	IsVIDeo          bool          `json:"is_video"`          // True, if the sticker is a video sticker
	Thumbnail        *PhotoSize    `json:"thumbnail"`         // Optional. Sticker thumbnail in the .WEBP or .JPG format
	Emoji            string        `json:"emoji"`             // Optional. Emoji associated with the sticker
	SetName          string        `json:"set_name"`          // Optional. Name of the sticker set to which the sticker belongs
	PremiumAnimation *File         `json:"premium_animation"` // Optional. For premium regular stickers, premium animation for the sticker
	MaskPosition     *MaskPosition `json:"mask_position"`     // Optional. For mask stickers, the position where the mask should be placed
	CustomEmojiID    string        `json:"custom_emoji_id"`   // Optional. For custom emoji stickers, unique identifier of the custom emoji
	NeedsRepainting  bool          `json:"needs_repainting"`  // Optional. True, if the sticker must be repainted to a text color in messages, the color of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in other places
	FileSize         int64         `json:"file_size"`         // Optional. File size in bytes
}

// This object represents a sticker set.
type StickerSet struct {
	Name        string     `json:"name"`         // Sticker set name
	Title       string     `json:"title"`        // Sticker set title
	StickerType string     `json:"sticker_type"` // Type of stickers in the set, currently one of "regular", "mask", "custom_emoji"
	Stickers    []Sticker  `json:"stickers"`     // List of all set stickers
	Thumbnail   *PhotoSize `json:"thumbnail"`    // Optional. Sticker set thumbnail in the .WEBP, .TGS, or .WEBM format
}

// This object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  string  `json:"point"`   // The part of the face relative to which the mask should be placed. One of "forehead", "eyes", "mouth", or "chin".
	XShift float64 `json:"x_shift"` // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
	YShift float64 `json:"y_shift"` // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
	Scale  float64 `json:"scale"`   // Mask scaling coefficient. For example, 2.0 means double size.
}

// This object describes a sticker to be added to a sticker set.
type InputSticker struct {
	Sticker      *InputFile    `json:"sticker"`       // The added sticker. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, upload a new one using multipart/form-data, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. Animated and video stickers can't be uploaded via HTTP URL. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Format       string        `json:"format"`        // Format of the added sticker, must be one of "static" for a .WEBP or .PNG image, "animated" for a .TGS animation, "video" for a .WEBM video
	EmojiList    []string      `json:"emoji_list"`    // List of 1-20 emoji associated with the sticker
	MaskPosition *MaskPosition `json:"mask_position"` // Optional. Position where the mask should be placed on faces. For "mask" stickers only.
	Keywords     []string      `json:"keywords"`      // Optional. List of 0-20 search keywords for the sticker with total length of up to 64 characters. For "regular" and "custom_emoji" stickers only.
}

// This object represents a gift that can be sent by the bot.
type Gift struct {
	ID               string   `json:"id"`                 // Unique identifier of the gift
	Sticker          *Sticker `json:"sticker"`            // The sticker that represents the gift
	StarCount        int64    `json:"star_count"`         // The number of Telegram Stars that must be paid to send the sticker
	UpgradeStarCount int64    `json:"upgrade_star_count"` // Optional. The number of Telegram Stars that must be paid to upgrade the gift to a unique one
	TotalCount       int64    `json:"total_count"`        // Optional. The total number of the gifts of this type that can be sent; for limited gifts only
	RemainingCount   int64    `json:"remaining_count"`    // Optional. The number of remaining gifts of this type that can be sent; for limited gifts only
}

// This object represent a list of gifts.
type Gifts struct {
	Gifts []Gift `json:"gifts"` // The list of gifts
}

// This object represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	ID       string              `json:"id"`        // Unique identifier for this query
	From     *User               `json:"from"`      // Sender
	Query    string              `json:"query"`     // Text of the query (up to 256 characters)
	Offset   string              `json:"offset"`    // Offset of the results to be returned, can be controlled by the bot
	ChatType InlineQueryChatType `json:"chat_type"` // Optional. Type of the chat from which the inline query was sent. Can be either "sender" for a private chat with the inline query sender, "private", "group", "supergroup", or "channel". The chat type should be always known for requests sent from official clients and most third-party clients, unless the request was sent from a secret chat
	Location *Location           `json:"location"`  // Optional. Sender location, only for bots that request user location
}

// This object represents a button to be shown above inline query results. You must use exactly one of the optional fields.
type InlineQueryResultsButton struct {
	Text           string      `json:"text"`            // Label text on the button
	WebApp         *WebAppInfo `json:"web_app"`         // Optional. Description of the Web App that will be launched when the user presses the button. The Web App will be able to switch back to the inline mode using the method switchInlineQuery inside the Web App.
	StartParameter string      `json:"start_parameter"` // Optional. Deep-linking parameter for the /start message sent to the bot when a user presses the button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed. Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a 'Connect your YouTube account' button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an OAuth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
}

// This object represents one result of an inline query. Telegram clients currently support results of the following 20 types:
// - InlineQueryResultCachedAudio
// - InlineQueryResultCachedDocument
// - InlineQueryResultCachedGif
// - InlineQueryResultCachedMpeg4Gif
// - InlineQueryResultCachedPhoto
// - InlineQueryResultCachedSticker
// - InlineQueryResultCachedVideo
// - InlineQueryResultCachedVoice
// - InlineQueryResultArticle
// - InlineQueryResultAudio
// - InlineQueryResultContact
// - InlineQueryResultGame
// - InlineQueryResultDocument
// - InlineQueryResultGif
// - InlineQueryResultLocation
// - InlineQueryResultMpeg4Gif
// - InlineQueryResultPhoto
// - InlineQueryResultVenue
// - InlineQueryResultVideo
// - InlineQueryResultVoice
// Note: All URLs passed in inline query results will be available to end users and therefore must be assumed to be public.
type InlineQueryResult interface{}

// Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	Type                string                `json:"type"`                  // Type of the result, must be article
	ID                  string                `json:"id"`                    // Unique identifier for this result, 1-64 Bytes
	Title               string                `json:"title"`                 // Title of the result
	InputMessageContent InputMessageContent   `json:"input_message_content"` // Content of the message to be sent
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	Url                 string                `json:"url"`                   // Optional. URL of the result
	Description         string                `json:"description"`           // Optional. Short description of the result
	ThumbnailUrl        string                `json:"thumbnail_url"`         // Optional. Url of the thumbnail for the result
	ThumbnailWIDth      int64                 `json:"thumbnail_width"`       // Optional. Thumbnail width
	ThumbnailHeight     int64                 `json:"thumbnail_height"`      // Optional. Thumbnail height
}

// Represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	Type                  string                `json:"type"`                     // Type of the result, must be photo
	ID                    string                `json:"id"`                       // Unique identifier for this result, 1-64 bytes
	PhotoUrl              string                `json:"photo_url"`                // A valid URL of the photo. Photo must be in JPEG format. Photo size must not exceed 5MB
	ThumbnailUrl          string                `json:"thumbnail_url"`            // URL of the thumbnail for the photo
	PhotoWIDth            int64                 `json:"photo_width"`              // Optional. Width of the photo
	PhotoHeight           int64                 `json:"photo_height"`             // Optional. Height of the photo
	Title                 string                `json:"title"`                    // Optional. Title for the result
	Description           string                `json:"description"`              // Optional. Short description of the result
	Caption               string                `json:"caption"`                  // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode"`               // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities       []MessageEntity       `json:"caption_entities"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content"`    // Optional. Content of the message to be sent instead of the photo
}

// Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	Type                  string                `json:"type"`                     // Type of the result, must be gif
	ID                    string                `json:"id"`                       // Unique identifier for this result, 1-64 bytes
	GifUrl                string                `json:"gif_url"`                  // A valid URL for the GIF file
	GifWIDth              int64                 `json:"gif_width"`                // Optional. Width of the GIF
	GifHeight             int64                 `json:"gif_height"`               // Optional. Height of the GIF
	GifDuration           int64                 `json:"gif_duration"`             // Optional. Duration of the GIF in seconds
	ThumbnailUrl          string                `json:"thumbnail_url"`            // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailMimeType     string                `json:"thumbnail_mime_type"`      // Optional. MIME type of the thumbnail, must be one of "image/jpeg", "image/gif", or "video/mp4". Defaults to "image/jpeg"
	Title                 string                `json:"title"`                    // Optional. Title for the result
	Caption               string                `json:"caption"`                  // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode"`               // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities       []MessageEntity       `json:"caption_entities"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content"`    // Optional. Content of the message to be sent instead of the GIF animation
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	Type                  string                `json:"type"`                     // Type of the result, must be mpeg4_gif
	ID                    string                `json:"id"`                       // Unique identifier for this result, 1-64 bytes
	Mpeg4Url              string                `json:"mpeg4_url"`                // A valid URL for the MPEG4 file
	Mpeg4WIDth            int64                 `json:"mpeg4_width"`              // Optional. Video width
	Mpeg4Height           int64                 `json:"mpeg4_height"`             // Optional. Video height
	Mpeg4Duration         int64                 `json:"mpeg4_duration"`           // Optional. Video duration in seconds
	ThumbnailUrl          string                `json:"thumbnail_url"`            // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailMimeType     string                `json:"thumbnail_mime_type"`      // Optional. MIME type of the thumbnail, must be one of "image/jpeg", "image/gif", or "video/mp4". Defaults to "image/jpeg"
	Title                 string                `json:"title"`                    // Optional. Title for the result
	Caption               string                `json:"caption"`                  // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode"`               // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities       []MessageEntity       `json:"caption_entities"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content"`    // Optional. Content of the message to be sent instead of the video animation
}

// Represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultVideo struct {
	Type                  string                `json:"type"`                     // Type of the result, must be video
	ID                    string                `json:"id"`                       // Unique identifier for this result, 1-64 bytes
	VIDeoUrl              string                `json:"video_url"`                // A valid URL for the embedded video player or video file
	MimeType              string                `json:"mime_type"`                // MIME type of the content of the video URL, "text/html" or "video/mp4"
	ThumbnailUrl          string                `json:"thumbnail_url"`            // URL of the thumbnail (JPEG only) for the video
	Title                 string                `json:"title"`                    // Title for the result
	Caption               string                `json:"caption"`                  // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode"`               // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities       []MessageEntity       `json:"caption_entities"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"` // Optional. Pass True, if the caption must be shown above the message media
	VIDeoWIDth            int64                 `json:"video_width"`              // Optional. Video width
	VIDeoHeight           int64                 `json:"video_height"`             // Optional. Video height
	VIDeoDuration         int64                 `json:"video_duration"`           // Optional. Video duration in seconds
	Description           string                `json:"description"`              // Optional. Short description of the result
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content"`    // Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
}

// Represents a link to an MP3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
	Type                string                `json:"type"`                  // Type of the result, must be audio
	ID                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	AudioUrl            string                `json:"audio_url"`             // A valid URL for the audio file
	Title               string                `json:"title"`                 // Title
	Caption             string                `json:"caption"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities     []MessageEntity       `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Performer           string                `json:"performer"`             // Optional. Performer
	AudioDuration       int64                 `json:"audio_duration"`        // Optional. Audio duration in seconds
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content"` // Optional. Content of the message to be sent instead of the audio
}

// Represents a link to a voice recording in an .OGG container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	Type                string                `json:"type"`                  // Type of the result, must be voice
	ID                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	VoiceUrl            string                `json:"voice_url"`             // A valid URL for the voice recording
	Title               string                `json:"title"`                 // Recording title
	Caption             string                `json:"caption"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities     []MessageEntity       `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	VoiceDuration       int64                 `json:"voice_duration"`        // Optional. Recording duration in seconds
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content"` // Optional. Content of the message to be sent instead of the voice recording
}

// Represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	Type                string                `json:"type"`                  // Type of the result, must be document
	ID                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	Title               string                `json:"title"`                 // Title for the result
	Caption             string                `json:"caption"`               // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities     []MessageEntity       `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	DocumentUrl         string                `json:"document_url"`          // A valid URL for the file
	MimeType            string                `json:"mime_type"`             // MIME type of the content of the file, either "application/pdf" or "application/zip"
	Description         string                `json:"description"`           // Optional. Short description of the result
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content"` // Optional. Content of the message to be sent instead of the file
	ThumbnailUrl        string                `json:"thumbnail_url"`         // Optional. URL of the thumbnail (JPEG only) for the file
	ThumbnailWIDth      int64                 `json:"thumbnail_width"`       // Optional. Thumbnail width
	ThumbnailHeight     int64                 `json:"thumbnail_height"`      // Optional. Thumbnail height
}

// Represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
	Type                 string                `json:"type"`                   // Type of the result, must be location
	ID                   string                `json:"id"`                     // Unique identifier for this result, 1-64 Bytes
	Latitude             float64               `json:"latitude"`               // Location latitude in degrees
	Longitude            float64               `json:"longitude"`              // Location longitude in degrees
	Title                string                `json:"title"`                  // Location title
	HorizontalAccuracy   float64               `json:"horizontal_accuracy"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int64                 `json:"live_period"`            // Optional. Period in seconds during which the location can be updated, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	Heading              int64                 `json:"heading"`                // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int64                 `json:"proximity_alert_radius"` // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup"`           // Optional. Inline keyboard attached to the message
	InputMessageContent  InputMessageContent   `json:"input_message_content"`  // Optional. Content of the message to be sent instead of the location
	ThumbnailUrl         string                `json:"thumbnail_url"`          // Optional. Url of the thumbnail for the result
	ThumbnailWIDth       int64                 `json:"thumbnail_width"`        // Optional. Thumbnail width
	ThumbnailHeight      int64                 `json:"thumbnail_height"`       // Optional. Thumbnail height
}

// Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	Type                string                `json:"type"`                  // Type of the result, must be venue
	ID                  string                `json:"id"`                    // Unique identifier for this result, 1-64 Bytes
	Latitude            float64               `json:"latitude"`              // Latitude of the venue location in degrees
	Longitude           float64               `json:"longitude"`             // Longitude of the venue location in degrees
	Title               string                `json:"title"`                 // Title of the venue
	Address             string                `json:"address"`               // Address of the venue
	FoursquareID        string                `json:"foursquare_id"`         // Optional. Foursquare identifier of the venue if known
	FoursquareType      string                `json:"foursquare_type"`       // Optional. Foursquare type of the venue, if known. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	GooglePlaceID       string                `json:"google_place_id"`       // Optional. Google Places identifier of the venue
	GooglePlaceType     string                `json:"google_place_type"`     // Optional. Google Places type of the venue. (See supported types.)
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content"` // Optional. Content of the message to be sent instead of the venue
	ThumbnailUrl        string                `json:"thumbnail_url"`         // Optional. Url of the thumbnail for the result
	ThumbnailWIDth      int64                 `json:"thumbnail_width"`       // Optional. Thumbnail width
	ThumbnailHeight     int64                 `json:"thumbnail_height"`      // Optional. Thumbnail height
}

// Represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	Type                string                `json:"type"`                  // Type of the result, must be contact
	ID                  string                `json:"id"`                    // Unique identifier for this result, 1-64 Bytes
	PhoneNumber         string                `json:"phone_number"`          // Contact's phone number
	FirstName           string                `json:"first_name"`            // Contact's first name
	LastName            string                `json:"last_name"`             // Optional. Contact's last name
	Vcard               string                `json:"vcard"`                 // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content"` // Optional. Content of the message to be sent instead of the contact
	ThumbnailUrl        string                `json:"thumbnail_url"`         // Optional. Url of the thumbnail for the result
	ThumbnailWIDth      int64                 `json:"thumbnail_width"`       // Optional. Thumbnail width
	ThumbnailHeight     int64                 `json:"thumbnail_height"`      // Optional. Thumbnail height
}

// Represents a Game.
type InlineQueryResultGame struct {
	Type          string                `json:"type"`            // Type of the result, must be game
	ID            string                `json:"id"`              // Unique identifier for this result, 1-64 bytes
	GameShortName string                `json:"game_short_name"` // Short name of the game
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup"`    // Optional. Inline keyboard attached to the message
}

// Represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	Type                  string                `json:"type"`                     // Type of the result, must be photo
	ID                    string                `json:"id"`                       // Unique identifier for this result, 1-64 bytes
	PhotoFileID           string                `json:"photo_file_id"`            // A valid file identifier of the photo
	Title                 string                `json:"title"`                    // Optional. Title for the result
	Description           string                `json:"description"`              // Optional. Short description of the result
	Caption               string                `json:"caption"`                  // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode"`               // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities       []MessageEntity       `json:"caption_entities"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content"`    // Optional. Content of the message to be sent instead of the photo
}

// Represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	Type                  string                `json:"type"`                     // Type of the result, must be gif
	ID                    string                `json:"id"`                       // Unique identifier for this result, 1-64 bytes
	GifFileID             string                `json:"gif_file_id"`              // A valid file identifier for the GIF file
	Title                 string                `json:"title"`                    // Optional. Title for the result
	Caption               string                `json:"caption"`                  // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode"`               // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities       []MessageEntity       `json:"caption_entities"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content"`    // Optional. Content of the message to be sent instead of the GIF animation
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	Type                  string                `json:"type"`                     // Type of the result, must be mpeg4_gif
	ID                    string                `json:"id"`                       // Unique identifier for this result, 1-64 bytes
	Mpeg4FileID           string                `json:"mpeg4_file_id"`            // A valid file identifier for the MPEG4 file
	Title                 string                `json:"title"`                    // Optional. Title for the result
	Caption               string                `json:"caption"`                  // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode"`               // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities       []MessageEntity       `json:"caption_entities"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content"`    // Optional. Content of the message to be sent instead of the video animation
}

// Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`                  // Type of the result, must be sticker
	ID                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	StickerFileID       string                `json:"sticker_file_id"`       // A valid file identifier of the sticker
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content"` // Optional. Content of the message to be sent instead of the sticker
}

// Represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`                  // Type of the result, must be document
	ID                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	Title               string                `json:"title"`                 // Title for the result
	DocumentFileID      string                `json:"document_file_id"`      // A valid file identifier for the file
	Description         string                `json:"description"`           // Optional. Short description of the result
	Caption             string                `json:"caption"`               // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities     []MessageEntity       `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content"` // Optional. Content of the message to be sent instead of the file
}

// Represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	Type                  string                `json:"type"`                     // Type of the result, must be video
	ID                    string                `json:"id"`                       // Unique identifier for this result, 1-64 bytes
	VIDeoFileID           string                `json:"video_file_id"`            // A valid file identifier for the video file
	Title                 string                `json:"title"`                    // Title for the result
	Description           string                `json:"description"`              // Optional. Short description of the result
	Caption               string                `json:"caption"`                  // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode             string                `json:"parse_mode"`               // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities       []MessageEntity       `json:"caption_entities"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media"` // Optional. Pass True, if the caption must be shown above the message media
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`             // Optional. Inline keyboard attached to the message
	InputMessageContent   InputMessageContent   `json:"input_message_content"`    // Optional. Content of the message to be sent instead of the video
}

// Represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`                  // Type of the result, must be voice
	ID                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	VoiceFileID         string                `json:"voice_file_id"`         // A valid file identifier for the voice message
	Title               string                `json:"title"`                 // Voice message title
	Caption             string                `json:"caption"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities     []MessageEntity       `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content"` // Optional. Content of the message to be sent instead of the voice message
}

// Represents a link to an MP3 audio file stored on the Telegram servers. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`                  // Type of the result, must be audio
	ID                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	AudioFileID         string                `json:"audio_file_id"`         // A valid file identifier for the audio file
	Caption             string                `json:"caption"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities     []MessageEntity       `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent InputMessageContent   `json:"input_message_content"` // Optional. Content of the message to be sent instead of the audio
}

// This object represents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 5 types:
// - InputTextMessageContent
// - InputLocationMessageContent
// - InputVenueMessageContent
// - InputContactMessageContent
// - InputInvoiceMessageContent
type InputMessageContent interface{}

// Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`         // Text of the message to be sent, 1-4096 characters
	ParseMode          string              `json:"parse_mode"`           // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	Entities           []MessageEntity     `json:"entities"`             // Optional. List of special entities that appear in message text, which can be specified instead of parse_mode
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options"` // Optional. Link preview generation options for the message
}

// Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`               // Latitude of the location in degrees
	Longitude            float64 `json:"longitude"`              // Longitude of the location in degrees
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int64   `json:"live_period"`            // Optional. Period in seconds during which the location can be updated, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	Heading              int64   `json:"heading"`                // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int64   `json:"proximity_alert_radius"` // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
}

// Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`          // Latitude of the venue in degrees
	Longitude       float64 `json:"longitude"`         // Longitude of the venue in degrees
	Title           string  `json:"title"`             // Name of the venue
	Address         string  `json:"address"`           // Address of the venue
	FoursquareID    string  `json:"foursquare_id"`     // Optional. Foursquare identifier of the venue, if known
	FoursquareType  string  `json:"foursquare_type"`   // Optional. Foursquare type of the venue, if known. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	GooglePlaceID   string  `json:"google_place_id"`   // Optional. Google Places identifier of the venue
	GooglePlaceType string  `json:"google_place_type"` // Optional. Google Places type of the venue. (See supported types.)
}

// Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"` // Contact's phone number
	FirstName   string `json:"first_name"`   // Contact's first name
	LastName    string `json:"last_name"`    // Optional. Contact's last name
	Vcard       string `json:"vcard"`        // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
}

// Represents the content of an invoice message to be sent as the result of an inline query.
type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`                         // Product name, 1-32 characters
	Description               string         `json:"description"`                   // Product description, 1-255 characters
	Payload                   string         `json:"payload"`                       // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	ProvIDerToken             string         `json:"provider_token"`                // Optional. Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	Currency                  string         `json:"currency"`                      // Three-letter ISO 4217 currency code, see more on currencies. Pass "XTR" for payments in Telegram Stars.
	Prices                    []LabeledPrice `json:"prices"`                        // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
	MaxTipAmount              int64          `json:"max_tip_amount"`                // Optional. The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
	SuggestedTipAmounts       []int64        `json:"suggested_tip_amounts"`         // Optional. A JSON-serialized array of suggested amounts of tip in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	ProvIDerData              string         `json:"provider_data"`                 // Optional. A JSON-serialized object for data about the invoice, which will be shared with the payment provider. A detailed description of the required fields should be provided by the payment provider.
	PhotoUrl                  string         `json:"photo_url"`                     // Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoSize                 int64          `json:"photo_size"`                    // Optional. Photo size in bytes
	PhotoWIDth                int64          `json:"photo_width"`                   // Optional. Photo width
	PhotoHeight               int64          `json:"photo_height"`                  // Optional. Photo height
	NeedName                  bool           `json:"need_name"`                     // Optional. Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber           bool           `json:"need_phone_number"`             // Optional. Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
	NeedEmail                 bool           `json:"need_email"`                    // Optional. Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress       bool           `json:"need_shipping_address"`         // Optional. Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvIDer bool           `json:"send_phone_number_to_provider"` // Optional. Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvIDer       bool           `json:"send_email_to_provider"`        // Optional. Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
	IsFlexible                bool           `json:"is_flexible"`                   // Optional. Pass True if the final price depends on the shipping method. Ignored for payments in Telegram Stars.
}

// Represents a result of an inline query that was chosen by the user and sent to their chat partner.
// Note: It is necessary to enable inline feedback via @BotFather in order to receive these objects in updates.
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`         // The unique identifier for the result that was chosen
	From            *User     `json:"from"`              // The user that chose the result
	Location        *Location `json:"location"`          // Optional. Sender location, only for bots that require user location
	InlineMessageID string    `json:"inline_message_id"` // Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	Query           string    `json:"query"`             // The query that was used to obtain the result
}

// Describes an inline message sent by a Web App on behalf of a user.
type SentWebAppMessage struct {
	InlineMessageID string `json:"inline_message_id"` // Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message.
}

// Describes an inline message to be sent by a user of a Mini App.
type PreparedInlineMessage struct {
	ID             string `json:"id"`              // Unique identifier of the prepared message
	ExpirationDate int64  `json:"expiration_date"` // Expiration date of the prepared message, in Unix time. Expired prepared messages can no longer be used
}

// This object represents a portion of the price for goods or services.
type LabeledPrice struct {
	Label  string `json:"label"`  // Portion label
	Amount int64  `json:"amount"` // Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

// This object contains basic information about an invoice.
type Invoice struct {
	Title          string `json:"title"`           // Product name
	Description    string `json:"description"`     // Product description
	StartParameter string `json:"start_parameter"` // Unique bot deep-linking parameter that can be used to generate this invoice
	Currency       string `json:"currency"`        // Three-letter ISO 4217 currency code, or "XTR" for payments in Telegram Stars
	TotalAmount    int64  `json:"total_amount"`    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

// This object represents a shipping address.
type ShippingAddress struct {
	CountryCode string `json:"country_code"` // Two-letter ISO 3166-1 alpha-2 country code
	State       string `json:"state"`        // State, if applicable
	City        string `json:"city"`         // City
	StreetLine1 string `json:"street_line1"` // First line for the address
	StreetLine2 string `json:"street_line2"` // Second line for the address
	PostCode    string `json:"post_code"`    // Address post code
}

// This object represents information about an order.
type OrderInfo struct {
	Name            string           `json:"name"`             // Optional. User name
	PhoneNumber     string           `json:"phone_number"`     // Optional. User's phone number
	Email           string           `json:"email"`            // Optional. User email
	ShippingAddress *ShippingAddress `json:"shipping_address"` // Optional. User shipping address
}

// This object represents one shipping option.
type ShippingOption struct {
	ID     string         `json:"id"`     // Shipping option identifier
	Title  string         `json:"title"`  // Option title
	Prices []LabeledPrice `json:"prices"` // List of price portions
}

// This object contains basic information about a successful payment. Note that if the buyer initiates a chargeback with the relevant payment provider following this transaction, the funds may be debited from your balance. This is outside of Telegram's control.
type SuccessfulPayment struct {
	Currency                   string     `json:"currency"`                     // Three-letter ISO 4217 currency code, or "XTR" for payments in Telegram Stars
	TotalAmount                int64      `json:"total_amount"`                 // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload             string     `json:"invoice_payload"`              // Bot-specified invoice payload
	SubscriptionExpirationDate int64      `json:"subscription_expiration_date"` // Optional. Expiration date of the subscription, in Unix time; for recurring payments only
	IsRecurring                bool       `json:"is_recurring"`                 // Optional. True, if the payment is a recurring payment for a subscription
	IsFirstRecurring           bool       `json:"is_first_recurring"`           // Optional. True, if the payment is the first payment for a subscription
	ShippingOptionID           string     `json:"shipping_option_id"`           // Optional. Identifier of the shipping option chosen by the user
	OrderInfo                  *OrderInfo `json:"order_info"`                   // Optional. Order information provided by the user
	TelegramPaymentChargeID    string     `json:"telegram_payment_charge_id"`   // Telegram payment identifier
	ProvIDerPaymentChargeID    string     `json:"provider_payment_charge_id"`   // Provider payment identifier
}

// This object contains basic information about a refunded payment.
type RefundedPayment struct {
	Currency                string `json:"currency"`                   // Three-letter ISO 4217 currency code, or "XTR" for payments in Telegram Stars. Currently, always "XTR"
	TotalAmount             int64  `json:"total_amount"`               // Total refunded price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45, total_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload          string `json:"invoice_payload"`            // Bot-specified invoice payload
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"` // Telegram payment identifier
	ProvIDerPaymentChargeID string `json:"provider_payment_charge_id"` // Optional. Provider payment identifier
}

// This object contains information about an incoming shipping query.
type ShippingQuery struct {
	ID              string           `json:"id"`               // Unique query identifier
	From            *User            `json:"from"`             // User who sent the query
	InvoicePayload  string           `json:"invoice_payload"`  // Bot-specified invoice payload
	ShippingAddress *ShippingAddress `json:"shipping_address"` // User specified shipping address
}

// This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	ID               string     `json:"id"`                 // Unique query identifier
	From             *User      `json:"from"`               // User who sent the query
	Currency         string     `json:"currency"`           // Three-letter ISO 4217 currency code, or "XTR" for payments in Telegram Stars
	TotalAmount      int64      `json:"total_amount"`       // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload   string     `json:"invoice_payload"`    // Bot-specified invoice payload
	ShippingOptionID string     `json:"shipping_option_id"` // Optional. Identifier of the shipping option chosen by the user
	OrderInfo        *OrderInfo `json:"order_info"`         // Optional. Order information provided by the user
}

// This object contains information about a paid media purchase.
type PaidMediaPurchased struct {
	From             *User  `json:"from"`               // User who purchased the media
	PaIDMediaPayload string `json:"paid_media_payload"` // Bot-specified paid media payload
}

// This object describes the state of a revenue withdrawal operation. Currently, it can be one of
// - RevenueWithdrawalStatePending
// - RevenueWithdrawalStateSucceeded
// - RevenueWithdrawalStateFailed
type RevenueWithdrawalState interface{}

// The withdrawal is in progress.
type RevenueWithdrawalStatePending struct {
	Type string `json:"type"` // Type of the state, always "pending"
}

// The withdrawal succeeded.
type RevenueWithdrawalStateSucceeded struct {
	Type string `json:"type"` // Type of the state, always "succeeded"
	Date int64  `json:"date"` // Date the withdrawal was completed in Unix time
	Url  string `json:"url"`  // An HTTPS URL that can be used to see transaction details
}

// The withdrawal failed and the transaction was refunded.
type RevenueWithdrawalStateFailed struct {
	Type string `json:"type"` // Type of the state, always "failed"
}

// Contains information about the affiliate that received a commission via this transaction.
type AffiliateInfo struct {
	AffiliateUser      *User `json:"affiliate_user"`       // Optional. The bot or the user that received an affiliate commission if it was received by a bot or a user
	AffiliateChat      *Chat `json:"affiliate_chat"`       // Optional. The chat that received an affiliate commission if it was received by a chat
	CommissionPerMille int64 `json:"commission_per_mille"` // The number of Telegram Stars received by the affiliate for each 1000 Telegram Stars received by the bot from referred users
	Amount             int64 `json:"amount"`               // Integer amount of Telegram Stars received by the affiliate from the transaction, rounded to 0; can be negative for refunds
	NanostarAmount     int64 `json:"nanostar_amount"`      // Optional. The number of 1/1000000000 shares of Telegram Stars received by the affiliate; from -999999999 to 999999999; can be negative for refunds
}

// This object describes the source of a transaction, or its recipient for outgoing transactions. Currently, it can be one of
// - TransactionPartnerUser
// - TransactionPartnerChat
// - TransactionPartnerAffiliateProgram
// - TransactionPartnerFragment
// - TransactionPartnerTelegramAds
// - TransactionPartnerTelegramApi
// - TransactionPartnerOther
type TransactionPartner interface{}

// Describes a transaction with a user.
type TransactionPartnerUser struct {
	Type               string         `json:"type"`                // Type of the transaction partner, always "user"
	User               *User          `json:"user"`                // Information about the user
	Affiliate          *AffiliateInfo `json:"affiliate"`           // Optional. Information about the affiliate that received a commission via this transaction
	InvoicePayload     string         `json:"invoice_payload"`     // Optional. Bot-specified invoice payload
	SubscriptionPeriod int64          `json:"subscription_period"` // Optional. The duration of the paid subscription
	PaIDMedia          []PaidMedia    `json:"paid_media"`          // Optional. Information about the paid media bought by the user
	PaIDMediaPayload   string         `json:"paid_media_payload"`  // Optional. Bot-specified paid media payload
	Gift               *Gift          `json:"gift"`                // Optional. The gift sent to the user by the bot
}

// Describes a transaction with a chat.
type TransactionPartnerChat struct {
	Type string `json:"type"` // Type of the transaction partner, always "chat"
	Chat *Chat  `json:"chat"` // Information about the chat
	Gift *Gift  `json:"gift"` // Optional. The gift sent to the chat by the bot
}

// Describes the affiliate program that issued the affiliate commission received via this transaction.
type TransactionPartnerAffiliateProgram struct {
	Type               string `json:"type"`                 // Type of the transaction partner, always "affiliate_program"
	SponsorUser        *User  `json:"sponsor_user"`         // Optional. Information about the bot that sponsored the affiliate program
	CommissionPerMille int64  `json:"commission_per_mille"` // The number of Telegram Stars received by the bot for each 1000 Telegram Stars received by the affiliate program sponsor from referred users
}

// Describes a withdrawal transaction with Fragment.
type TransactionPartnerFragment struct {
	Type            string                  `json:"type"`             // Type of the transaction partner, always "fragment"
	WithdrawalState *RevenueWithdrawalState `json:"withdrawal_state"` // Optional. State of the transaction if the transaction is outgoing
}

// Describes a withdrawal transaction to the Telegram Ads platform.
type TransactionPartnerTelegramAds struct {
	Type string `json:"type"` // Type of the transaction partner, always "telegram_ads"
}

// Describes a transaction with payment for paid broadcasting.
type TransactionPartnerTelegramApi struct {
	Type         string `json:"type"`          // Type of the transaction partner, always "telegram_api"
	RequestCount int64  `json:"request_count"` // The number of successful requests that exceeded regular limits and were therefore billed
}

// Describes a transaction with an unknown source or recipient.
type TransactionPartnerOther struct {
	Type string `json:"type"` // Type of the transaction partner, always "other"
}

// Describes a Telegram Star transaction. Note that if the buyer initiates a chargeback with the payment provider from whom they acquired Stars (e.g., Apple, Google) following this transaction, the refunded Stars will be deducted from the bot's balance. This is outside of Telegram's control.
type StarTransaction struct {
	ID             string              `json:"id"`              // Unique identifier of the transaction. Coincides with the identifier of the original transaction for refund transactions. Coincides with SuccessfulPayment.telegram_payment_charge_id for successful incoming payments from users.
	Amount         int64               `json:"amount"`          // Integer amount of Telegram Stars transferred by the transaction
	NanostarAmount int64               `json:"nanostar_amount"` // Optional. The number of 1/1000000000 shares of Telegram Stars transferred by the transaction; from 0 to 999999999
	Date           int64               `json:"date"`            // Date the transaction was created in Unix time
	Source         *TransactionPartner `json:"source"`          // Optional. Source of an incoming transaction (e.g., a user purchasing goods or services, Fragment refunding a failed withdrawal). Only for incoming transactions
	Receiver       *TransactionPartner `json:"receiver"`        // Optional. Receiver of an outgoing transaction (e.g., a user for a purchase refund, Fragment for a withdrawal). Only for outgoing transactions
}

// Contains a list of Telegram Star transactions.
type StarTransactions struct {
	Transactions []StarTransaction `json:"transactions"` // The list of transactions
}

// Describes Telegram Passport data shared with the bot by the user.
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`        // Array with information about documents and other Telegram Passport elements that was shared with the bot
	Credentials *EncryptedCredentials      `json:"credentials"` // Encrypted credentials required to decrypt the data
}

// This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	FileID       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int64  `json:"file_size"`      // File size in bytes
	FileDate     int64  `json:"file_date"`      // Unix time when the file was uploaded
}

// Describes documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	Type        string         `json:"type"`         // Element type. One of "personal_details", "passport", "driver_license", "identity_card", "internal_passport", "address", "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration", "phone_number", "email".
	Data        string         `json:"data"`         // Optional. Base64-encoded encrypted Telegram Passport element data provided by the user; available only for "personal_details", "passport", "driver_license", "identity_card", "internal_passport" and "address" types. Can be decrypted and verified using the accompanying EncryptedCredentials.
	PhoneNumber string         `json:"phone_number"` // Optional. User's verified phone number; available only for "phone_number" type
	Email       string         `json:"email"`        // Optional. User's verified email address; available only for "email" type
	Files       []PassportFile `json:"files"`        // Optional. Array of encrypted files with documents provided by the user; available only for "utility_bill", "bank_statement", "rental_agreement", "passport_registration" and "temporary_registration" types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	FrontSIDe   *PassportFile  `json:"front_side"`   // Optional. Encrypted file with the front side of the document, provided by the user; available only for "passport", "driver_license", "identity_card" and "internal_passport". The file can be decrypted and verified using the accompanying EncryptedCredentials.
	ReverseSIDe *PassportFile  `json:"reverse_side"` // Optional. Encrypted file with the reverse side of the document, provided by the user; available only for "driver_license" and "identity_card". The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie      *PassportFile  `json:"selfie"`       // Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available if requested for "passport", "driver_license", "identity_card" and "internal_passport". The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Translation []PassportFile `json:"translation"`  // Optional. Array of encrypted files with translated versions of documents provided by the user; available if requested for "passport", "driver_license", "identity_card", "internal_passport", "utility_bill", "bank_statement", "rental_agreement", "passport_registration" and "temporary_registration" types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Hash        string         `json:"hash"`         // Base64-encoded element hash for using in PassportElementErrorUnspecified
}

// Describes data required for decrypting and authenticating EncryptedPassportElement. See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
type EncryptedCredentials struct {
	Data   string `json:"data"`   // Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
	Hash   string `json:"hash"`   // Base64-encoded data hash for data authentication
	Secret string `json:"secret"` // Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
}

// This object represents an error in the Telegram Passport element which was submitted that should be resolved by the user. It should be one of:
// - PassportElementErrorDataField
// - PassportElementErrorFrontSide
// - PassportElementErrorReverseSide
// - PassportElementErrorSelfie
// - PassportElementErrorFile
// - PassportElementErrorFiles
// - PassportElementErrorTranslationFile
// - PassportElementErrorTranslationFiles
// - PassportElementErrorUnspecified
type PassportElementError interface{}

// Represents an issue in one of the data fields that was provided by the user. The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
	Source    string `json:"source"`     // Error source, must be data
	Type      string `json:"type"`       // The section of the user's Telegram Passport which has the error, one of "personal_details", "passport", "driver_license", "identity_card", "internal_passport", "address"
	FieldName string `json:"field_name"` // Name of the data field which has the error
	DataHash  string `json:"data_hash"`  // Base64-encoded data hash
	Message   string `json:"message"`    // Error message
}

// Represents an issue with the front side of a document. The error is considered resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`    // Error source, must be front_side
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of "passport", "driver_license", "identity_card", "internal_passport"
	FileHash string `json:"file_hash"` // Base64-encoded hash of the file with the front side of the document
	Message  string `json:"message"`   // Error message
}

// Represents an issue with the reverse side of a document. The error is considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
	Source   string `json:"source"`    // Error source, must be reverse_side
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of "driver_license", "identity_card"
	FileHash string `json:"file_hash"` // Base64-encoded hash of the file with the reverse side of the document
	Message  string `json:"message"`   // Error message
}

// Represents an issue with the selfie with a document. The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	Source   string `json:"source"`    // Error source, must be selfie
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of "passport", "driver_license", "identity_card", "internal_passport"
	FileHash string `json:"file_hash"` // Base64-encoded hash of the file with the selfie
	Message  string `json:"message"`   // Error message
}

// Represents an issue with a document scan. The error is considered resolved when the file with the document scan changes.
type PassportElementErrorFile struct {
	Source   string `json:"source"`    // Error source, must be file
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
	FileHash string `json:"file_hash"` // Base64-encoded file hash
	Message  string `json:"message"`   // Error message
}

// Represents an issue with a list of scans. The error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
	Source     string   `json:"source"`      // Error source, must be files
	Type       string   `json:"type"`        // The section of the user's Telegram Passport which has the issue, one of "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
	FileHashes []string `json:"file_hashes"` // List of base64-encoded file hashes
	Message    string   `json:"message"`     // Error message
}

// Represents an issue with one of the files that constitute the translation of a document. The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"`    // Error source, must be translation_file
	Type     string `json:"type"`      // Type of element of the user's Telegram Passport which has the issue, one of "passport", "driver_license", "identity_card", "internal_passport", "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
	FileHash string `json:"file_hash"` // Base64-encoded file hash
	Message  string `json:"message"`   // Error message
}

// Represents an issue with the translated version of a document. The error is considered resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source"`      // Error source, must be translation_files
	Type       string   `json:"type"`        // Type of element of the user's Telegram Passport which has the issue, one of "passport", "driver_license", "identity_card", "internal_passport", "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
	FileHashes []string `json:"file_hashes"` // List of base64-encoded file hashes
	Message    string   `json:"message"`     // Error message
}

// Represents an issue in an unspecified place. The error is considered resolved when new data is added.
type PassportElementErrorUnspecified struct {
	Source      string `json:"source"`       // Error source, must be unspecified
	Type        string `json:"type"`         // Type of element of the user's Telegram Passport which has the issue
	ElementHash string `json:"element_hash"` // Base64-encoded element hash
	Message     string `json:"message"`      // Error message
}

// This object represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	Title        string          `json:"title"`         // Title of the game
	Description  string          `json:"description"`   // Description of the game
	Photo        []PhotoSize     `json:"photo"`         // Photo that will be displayed in the game message in chats.
	Text         string          `json:"text"`          // Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
	TextEntities []MessageEntity `json:"text_entities"` // Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	Animation    *Animation      `json:"animation"`     // Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
}

// A placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame interface{}

// This object represents one row of the high scores table for a game.
type GameHighScore struct {
	Position int64 `json:"position"` // Position in high score table for the game
	User     *User `json:"user"`     // User
	Score    int64 `json:"score"`    // Score
}
