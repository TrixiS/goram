package types

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
	ChatActionRecordVIDeo     ChatAction = "record_video"
	ChatActionUploadVIDeo     ChatAction = "upload_video"
	ChatActionRecordVoice     ChatAction = "record_voice"
	ChatActionUploadVoice     ChatAction = "upload_voice"
	ChatActionUploadDocument  ChatAction = "upload_document"
	ChatActionFindLocation    ChatAction = "find_location"
	ChatActionRecordVIDeoNote ChatAction = "record_video_note"
	ChatActionUploadVIDeoNote ChatAction = "upload_video_note"
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
