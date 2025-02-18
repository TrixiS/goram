package bot

import (
	"context"
	"github.com/TrixiS/goram/pkg/types"
)

// Use this method to receive incoming updates using long polling (wiki). Returns an Array of Update objects.
//
// https://core.telegram.org/bots/api#getupdates
func (b *Bot) GetUpdates(ctx context.Context, request *types.GetUpdatesRequest) (r []types.Update, err error) {
	res, err := makeRequest[[]types.Update](ctx, b.options.Client, b.baseURL, "getUpdates", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request (a request with response HTTP status code different from 2XY), we will repeat the request and give up after a reasonable amount of attempts. Returns True on success.
//
// If you'd like to make sure that the webhook was set by you, you can specify secret data in the parameter secret_token. If specified, the request will contain a header "X-Telegram-Bot-Api-Secret-Token" with the secret token as content.
//
// https://core.telegram.org/bots/api#setwebhook
func (b *Bot) SetWebhook(ctx context.Context, request *types.SetWebhookRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setWebhook", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
//
// https://core.telegram.org/bots/api#deletewebhook
func (b *Bot) DeleteWebhook(ctx context.Context, request *types.DeleteWebhookRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "deleteWebhook", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
//
// https://core.telegram.org/bots/api#getwebhookinfo
func (b *Bot) GetWebhookInfo(ctx context.Context) (r *types.WebhookInfo, err error) {
	res, err := makeRequest[*types.WebhookInfo](ctx, b.options.Client, b.baseURL, "getWebhookInfo", nil)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// A simple method for testing your bot's authentication token. Requires no parameters. Returns basic information about the bot in form of a User object.
//
// https://core.telegram.org/bots/api#getme
func (b *Bot) GetMe(ctx context.Context) (r *types.User, err error) {
	res, err := makeRequest[*types.User](ctx, b.options.Client, b.baseURL, "getMe", nil)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to log out from the cloud Bot API server before launching the bot locally. You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates. After a successful call, you can immediately log in on a local server, but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
//
// https://core.telegram.org/bots/api#logout
func (b *Bot) LogOut(ctx context.Context) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "logOut", nil)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to close the bot instance before moving it from one local server to another. You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns True on success. Requires no parameters.
//
// https://core.telegram.org/bots/api#close
func (b *Bot) Close(ctx context.Context) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "close", nil)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send text messages. On success, the sent Message is returned.
//
// https://core.telegram.org/bots/api#sendmessage
func (b *Bot) SendMessage(ctx context.Context, request *types.SendMessageRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendMessage", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to forward messages of any kind. Service messages and messages with protected content can't be forwarded. On success, the sent Message is returned.
//
// https://core.telegram.org/bots/api#forwardmessage
func (b *Bot) ForwardMessage(ctx context.Context, request *types.ForwardMessageRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "forwardMessage", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to forward multiple messages of any kind. If some of the specified messages can't be found or forwarded, they are skipped. Service messages and messages with protected content can't be forwarded. Album grouping is kept for forwarded messages. On success, an array of MessageId of the sent messages is returned.
//
// https://core.telegram.org/bots/api#forwardmessages
func (b *Bot) ForwardMessages(ctx context.Context, request *types.ForwardMessagesRequest) (r []types.MessageID, err error) {
	res, err := makeRequest[[]types.MessageID](ctx, b.options.Client, b.baseURL, "forwardMessages", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to copy messages of any kind. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
//
// https://core.telegram.org/bots/api#copymessage
func (b *Bot) CopyMessage(ctx context.Context, request *types.CopyMessageRequest) (r *types.MessageID, err error) {
	res, err := makeRequest[*types.MessageID](ctx, b.options.Client, b.baseURL, "copyMessage", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to copy messages of any kind. If some of the specified messages can't be found or copied, they are skipped. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessages, but the copied messages don't have a link to the original message. Album grouping is kept for copied messages. On success, an array of MessageId of the sent messages is returned.
//
// https://core.telegram.org/bots/api#copymessages
func (b *Bot) CopyMessages(ctx context.Context, request *types.CopyMessagesRequest) (r []types.MessageID, err error) {
	res, err := makeRequest[[]types.MessageID](ctx, b.options.Client, b.baseURL, "copyMessages", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send photos. On success, the sent Message is returned.
//
// https://core.telegram.org/bots/api#sendphoto
func (b *Bot) SendPhoto(ctx context.Context, request *types.SendPhotoRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendPhoto", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
//
// For sending voice messages, use the sendVoice method instead.
//
// https://core.telegram.org/bots/api#sendaudio
func (b *Bot) SendAudio(ctx context.Context, request *types.SendAudioRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendAudio", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
//
// https://core.telegram.org/bots/api#senddocument
func (b *Bot) SendDocument(ctx context.Context, request *types.SendDocumentRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendDocument", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
//
// https://core.telegram.org/bots/api#sendvideo
func (b *Bot) SendVideo(ctx context.Context, request *types.SendVideoRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendVideo", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
//
// https://core.telegram.org/bots/api#sendanimation
func (b *Bot) SendAnimation(ctx context.Context, request *types.SendAnimationRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendAnimation", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A format (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
//
// https://core.telegram.org/bots/api#sendvoice
func (b *Bot) SendVoice(ctx context.Context, request *types.SendVoiceRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendVoice", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
//
// https://core.telegram.org/bots/api#sendvideonote
func (b *Bot) SendVideoNote(ctx context.Context, request *types.SendVideoNoteRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendVideoNote", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send paid media. On success, the sent Message is returned.
//
// https://core.telegram.org/bots/api#sendpaidmedia
func (b *Bot) SendPaidMedia(ctx context.Context, request *types.SendPaidMediaRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendPaidMedia", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send a group of photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Messages that were sent is returned.
//
// https://core.telegram.org/bots/api#sendmediagroup
func (b *Bot) SendMediaGroup(ctx context.Context, request *types.SendMediaGroupRequest) (r []types.Message, err error) {
	res, err := makeRequest[[]types.Message](ctx, b.options.Client, b.baseURL, "sendMediaGroup", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send point on the map. On success, the sent Message is returned.
//
// https://core.telegram.org/bots/api#sendlocation
func (b *Bot) SendLocation(ctx context.Context, request *types.SendLocationRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendLocation", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send information about a venue. On success, the sent Message is returned.
//
// https://core.telegram.org/bots/api#sendvenue
func (b *Bot) SendVenue(ctx context.Context, request *types.SendVenueRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendVenue", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send phone contacts. On success, the sent Message is returned.
//
// https://core.telegram.org/bots/api#sendcontact
func (b *Bot) SendContact(ctx context.Context, request *types.SendContactRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendContact", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send a native poll. On success, the sent Message is returned.
//
// https://core.telegram.org/bots/api#sendpoll
func (b *Bot) SendPoll(ctx context.Context, request *types.SendPollRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendPoll", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned.
//
// https://core.telegram.org/bots/api#senddice
func (b *Bot) SendDice(ctx context.Context, request *types.SendDiceRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendDice", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
//
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
//
// https://core.telegram.org/bots/api#sendchataction
func (b *Bot) SendChatAction(ctx context.Context, request *types.SendChatActionRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "sendChatAction", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to change the chosen reactions on a message. Service messages of some types can't be reacted to. Automatically forwarded messages from a channel to its discussion group have the same available reactions as messages in the channel. Bots can't use paid reactions. Returns True on success.
//
// https://core.telegram.org/bots/api#setmessagereaction
func (b *Bot) SetMessageReaction(ctx context.Context, request *types.SetMessageReactionRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setMessageReaction", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
//
// https://core.telegram.org/bots/api#getuserprofilephotos
func (b *Bot) GetUserProfilePhotos(ctx context.Context, request *types.GetUserProfilePhotosRequest) (r *types.UserProfilePhotos, err error) {
	res, err := makeRequest[*types.UserProfilePhotos](ctx, b.options.Client, b.baseURL, "getUserProfilePhotos", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Changes the emoji status for a given user that previously allowed the bot to manage their emoji status via the Mini App method requestEmojiStatusAccess. Returns True on success.
//
// https://core.telegram.org/bots/api#setuseremojistatus
func (b *Bot) SetUserEmojiStatus(ctx context.Context, request *types.SetUserEmojiStatusRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setUserEmojiStatus", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
//
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
//
// https://core.telegram.org/bots/api#getfile
func (b *Bot) GetFile(ctx context.Context, request *types.GetFileRequest) (r *types.File, err error) {
	res, err := makeRequest[*types.File](ctx, b.options.Client, b.baseURL, "getFile", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#banchatmember
func (b *Bot) BanChatMember(ctx context.Context, request *types.BanChatMemberRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "banChatMember", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to unban a previously banned user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
//
// https://core.telegram.org/bots/api#unbanchatmember
func (b *Bot) UnbanChatMember(ctx context.Context, request *types.UnbanChatMemberRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "unbanChatMember", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
//
// https://core.telegram.org/bots/api#restrictchatmember
func (b *Bot) RestrictChatMember(ctx context.Context, request *types.RestrictChatMemberRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "restrictChatMember", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass False for all boolean parameters to demote a user. Returns True on success.
//
// https://core.telegram.org/bots/api#promotechatmember
func (b *Bot) PromoteChatMember(ctx context.Context, request *types.PromoteChatMemberRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "promoteChatMember", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func (b *Bot) SetChatAdministratorCustomTitle(ctx context.Context, request *types.SetChatAdministratorCustomTitleRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setChatAdministratorCustomTitle", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#banchatsenderchat
func (b *Bot) BanChatSenderChat(ctx context.Context, request *types.BanChatSenderChatRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "banChatSenderChat", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator for this to work and must have the appropriate administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#unbanchatsenderchat
func (b *Bot) UnbanChatSenderChat(ctx context.Context, request *types.UnbanChatSenderChatRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "unbanChatSenderChat", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#setchatpermissions
func (b *Bot) SetChatPermissions(ctx context.Context, request *types.SetChatPermissionsRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setChatPermissions", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.
//
// https://core.telegram.org/bots/api#exportchatinvitelink
func (b *Bot) ExportChatInviteLink(ctx context.Context, request *types.ExportChatInviteLinkRequest) (r string, err error) {
	res, err := makeRequest[string](ctx, b.options.Client, b.baseURL, "exportChatInviteLink", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
//
// https://core.telegram.org/bots/api#createchatinvitelink
func (b *Bot) CreateChatInviteLink(ctx context.Context, request *types.CreateChatInviteLinkRequest) (r *types.ChatInviteLink, err error) {
	res, err := makeRequest[*types.ChatInviteLink](ctx, b.options.Client, b.baseURL, "createChatInviteLink", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the edited invite link as a ChatInviteLink object.
//
// https://core.telegram.org/bots/api#editchatinvitelink
func (b *Bot) EditChatInviteLink(ctx context.Context, request *types.EditChatInviteLinkRequest) (r *types.ChatInviteLink, err error) {
	res, err := makeRequest[*types.ChatInviteLink](ctx, b.options.Client, b.baseURL, "editChatInviteLink", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to create a subscription invite link for a channel chat. The bot must have the can_invite_users administrator rights. The link can be edited using the method editChatSubscriptionInviteLink or revoked using the method revokeChatInviteLink. Returns the new invite link as a ChatInviteLink object.
//
// https://core.telegram.org/bots/api#createchatsubscriptioninvitelink
func (b *Bot) CreateChatSubscriptionInviteLink(ctx context.Context, request *types.CreateChatSubscriptionInviteLinkRequest) (r *types.ChatInviteLink, err error) {
	res, err := makeRequest[*types.ChatInviteLink](ctx, b.options.Client, b.baseURL, "createChatSubscriptionInviteLink", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to edit a subscription invite link created by the bot. The bot must have the can_invite_users administrator rights. Returns the edited invite link as a ChatInviteLink object.
//
// https://core.telegram.org/bots/api#editchatsubscriptioninvitelink
func (b *Bot) EditChatSubscriptionInviteLink(ctx context.Context, request *types.EditChatSubscriptionInviteLinkRequest) (r *types.ChatInviteLink, err error) {
	res, err := makeRequest[*types.ChatInviteLink](ctx, b.options.Client, b.baseURL, "editChatSubscriptionInviteLink", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the revoked invite link as ChatInviteLink object.
//
// https://core.telegram.org/bots/api#revokechatinvitelink
func (b *Bot) RevokeChatInviteLink(ctx context.Context, request *types.RevokeChatInviteLinkRequest) (r *types.ChatInviteLink, err error) {
	res, err := makeRequest[*types.ChatInviteLink](ctx, b.options.Client, b.baseURL, "revokeChatInviteLink", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to approve a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
//
// https://core.telegram.org/bots/api#approvechatjoinrequest
func (b *Bot) ApproveChatJoinRequest(ctx context.Context, request *types.ApproveChatJoinRequestRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "approveChatJoinRequest", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to decline a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
//
// https://core.telegram.org/bots/api#declinechatjoinrequest
func (b *Bot) DeclineChatJoinRequest(ctx context.Context, request *types.DeclineChatJoinRequestRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "declineChatJoinRequest", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#setchatphoto
func (b *Bot) SetChatPhoto(ctx context.Context, request *types.SetChatPhotoRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setChatPhoto", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#deletechatphoto
func (b *Bot) DeleteChatPhoto(ctx context.Context, request *types.DeleteChatPhotoRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "deleteChatPhoto", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#setchattitle
func (b *Bot) SetChatTitle(ctx context.Context, request *types.SetChatTitleRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setChatTitle", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#setchatdescription
func (b *Bot) SetChatDescription(ctx context.Context, request *types.SetChatDescriptionRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setChatDescription", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to add a message to the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
//
// https://core.telegram.org/bots/api#pinchatmessage
func (b *Bot) PinChatMessage(ctx context.Context, request *types.PinChatMessageRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "pinChatMessage", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to remove a message from the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
//
// https://core.telegram.org/bots/api#unpinchatmessage
func (b *Bot) UnpinChatMessage(ctx context.Context, request *types.UnpinChatMessageRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "unpinChatMessage", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to clear the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
//
// https://core.telegram.org/bots/api#unpinallchatmessages
func (b *Bot) UnpinAllChatMessages(ctx context.Context, request *types.UnpinAllChatMessagesRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "unpinAllChatMessages", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
//
// https://core.telegram.org/bots/api#leavechat
func (b *Bot) LeaveChat(ctx context.Context, request *types.LeaveChatRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "leaveChat", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get up-to-date information about the chat. Returns a ChatFullInfo object on success.
//
// https://core.telegram.org/bots/api#getchat
func (b *Bot) GetChat(ctx context.Context, request *types.GetChatRequest) (r *types.ChatFullInfo, err error) {
	res, err := makeRequest[*types.ChatFullInfo](ctx, b.options.Client, b.baseURL, "getChat", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get a list of administrators in a chat, which aren't bots. Returns an Array of ChatMember objects.
//
// https://core.telegram.org/bots/api#getchatadministrators
func (b *Bot) GetChatAdministrators(ctx context.Context, request *types.GetChatAdministratorsRequest) (r []types.ChatMember, err error) {
	res, err := makeRequest[[]types.ChatMember](ctx, b.options.Client, b.baseURL, "getChatAdministrators", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get the number of members in a chat. Returns Int on success.
//
// https://core.telegram.org/bots/api#getchatmembercount
func (b *Bot) GetChatMemberCount(ctx context.Context, request *types.GetChatMemberCountRequest) (r int, err error) {
	res, err := makeRequest[int](ctx, b.options.Client, b.baseURL, "getChatMemberCount", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get information about a member of a chat. The method is only guaranteed to work for other users if the bot is an administrator in the chat. Returns a ChatMember object on success.
//
// https://core.telegram.org/bots/api#getchatmember
func (b *Bot) GetChatMember(ctx context.Context, request *types.GetChatMemberRequest) (r types.ChatMember, err error) {
	res, err := makeRequest[types.ChatMember](ctx, b.options.Client, b.baseURL, "getChatMember", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
//
// https://core.telegram.org/bots/api#setchatstickerset
func (b *Bot) SetChatStickerSet(ctx context.Context, request *types.SetChatStickerSetRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setChatStickerSet", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
//
// https://core.telegram.org/bots/api#deletechatstickerset
func (b *Bot) DeleteChatStickerSet(ctx context.Context, request *types.DeleteChatStickerSetRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "deleteChatStickerSet", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get custom emoji stickers, which can be used as a forum topic icon by any user. Requires no parameters. Returns an Array of Sticker objects.
//
// https://core.telegram.org/bots/api#getforumtopiciconstickers
func (b *Bot) GetForumTopicIconStickers(ctx context.Context) (r []types.Sticker, err error) {
	res, err := makeRequest[[]types.Sticker](ctx, b.options.Client, b.baseURL, "getForumTopicIconStickers", nil)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to create a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns information about the created topic as a ForumTopic object.
//
// https://core.telegram.org/bots/api#createforumtopic
func (b *Bot) CreateForumTopic(ctx context.Context, request *types.CreateForumTopicRequest) (r *types.ForumTopic, err error) {
	res, err := makeRequest[*types.ForumTopic](ctx, b.options.Client, b.baseURL, "createForumTopic", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to edit name and icon of a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
//
// https://core.telegram.org/bots/api#editforumtopic
func (b *Bot) EditForumTopic(ctx context.Context, request *types.EditForumTopicRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "editForumTopic", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
//
// https://core.telegram.org/bots/api#closeforumtopic
func (b *Bot) CloseForumTopic(ctx context.Context, request *types.CloseForumTopicRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "closeForumTopic", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
//
// https://core.telegram.org/bots/api#reopenforumtopic
func (b *Bot) ReopenForumTopic(ctx context.Context, request *types.ReopenForumTopicRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "reopenForumTopic", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to delete a forum topic along with all its messages in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#deleteforumtopic
func (b *Bot) DeleteForumTopic(ctx context.Context, request *types.DeleteForumTopicRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "deleteForumTopic", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to clear the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
//
// https://core.telegram.org/bots/api#unpinallforumtopicmessages
func (b *Bot) UnpinAllForumTopicMessages(ctx context.Context, request *types.UnpinAllForumTopicMessagesRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "unpinAllForumTopicMessages", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#editgeneralforumtopic
func (b *Bot) EditGeneralForumTopic(ctx context.Context, request *types.EditGeneralForumTopicRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "editGeneralForumTopic", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to close an open 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#closegeneralforumtopic
func (b *Bot) CloseGeneralForumTopic(ctx context.Context, request *types.CloseGeneralForumTopicRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "closeGeneralForumTopic", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to reopen a closed 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically unhidden if it was hidden. Returns True on success.
//
// https://core.telegram.org/bots/api#reopengeneralforumtopic
func (b *Bot) ReopenGeneralForumTopic(ctx context.Context, request *types.ReopenGeneralForumTopicRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "reopenGeneralForumTopic", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to hide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically closed if it was open. Returns True on success.
//
// https://core.telegram.org/bots/api#hidegeneralforumtopic
func (b *Bot) HideGeneralForumTopic(ctx context.Context, request *types.HideGeneralForumTopicRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "hideGeneralForumTopic", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to unhide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#unhidegeneralforumtopic
func (b *Bot) UnhideGeneralForumTopic(ctx context.Context, request *types.UnhideGeneralForumTopicRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "unhideGeneralForumTopic", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to clear the list of pinned messages in a General forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
//
// https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
func (b *Bot) UnpinAllGeneralForumTopicMessages(ctx context.Context, request *types.UnpinAllGeneralForumTopicMessagesRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "unpinAllGeneralForumTopicMessages", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
//
// https://core.telegram.org/bots/api#answercallbackquery
func (b *Bot) AnswerCallbackQuery(ctx context.Context, request *types.AnswerCallbackQueryRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "answerCallbackQuery", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Returns a UserChatBoosts object.
//
// https://core.telegram.org/bots/api#getuserchatboosts
func (b *Bot) GetUserChatBoosts(ctx context.Context, request *types.GetUserChatBoostsRequest) (r *types.UserChatBoosts, err error) {
	res, err := makeRequest[*types.UserChatBoosts](ctx, b.options.Client, b.baseURL, "getUserChatBoosts", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get information about the connection of the bot with a business account. Returns a BusinessConnection object on success.
//
// https://core.telegram.org/bots/api#getbusinessconnection
func (b *Bot) GetBusinessConnection(ctx context.Context, request *types.GetBusinessConnectionRequest) (r *types.BusinessConnection, err error) {
	res, err := makeRequest[*types.BusinessConnection](ctx, b.options.Client, b.baseURL, "getBusinessConnection", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to change the list of the bot's commands. See this manual for more details about bot commands. Returns True on success.
//
// https://core.telegram.org/bots/api#setmycommands
func (b *Bot) SetMyCommands(ctx context.Context, request *types.SetMyCommandsRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setMyCommands", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to delete the list of the bot's commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
//
// https://core.telegram.org/bots/api#deletemycommands
func (b *Bot) DeleteMyCommands(ctx context.Context, request *types.DeleteMyCommandsRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "deleteMyCommands", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get the current list of the bot's commands for the given scope and user language. Returns an Array of BotCommand objects. If commands aren't set, an empty list is returned.
//
// https://core.telegram.org/bots/api#getmycommands
func (b *Bot) GetMyCommands(ctx context.Context, request *types.GetMyCommandsRequest) (r []types.BotCommand, err error) {
	res, err := makeRequest[[]types.BotCommand](ctx, b.options.Client, b.baseURL, "getMyCommands", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to change the bot's name. Returns True on success.
//
// https://core.telegram.org/bots/api#setmyname
func (b *Bot) SetMyName(ctx context.Context, request *types.SetMyNameRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setMyName", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get the current bot name for the given user language. Returns BotName on success.
//
// https://core.telegram.org/bots/api#getmyname
func (b *Bot) GetMyName(ctx context.Context, request *types.GetMyNameRequest) (r *types.BotName, err error) {
	res, err := makeRequest[*types.BotName](ctx, b.options.Client, b.baseURL, "getMyName", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to change the bot's description, which is shown in the chat with the bot if the chat is empty. Returns True on success.
//
// https://core.telegram.org/bots/api#setmydescription
func (b *Bot) SetMyDescription(ctx context.Context, request *types.SetMyDescriptionRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setMyDescription", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get the current bot description for the given user language. Returns BotDescription on success.
//
// https://core.telegram.org/bots/api#getmydescription
func (b *Bot) GetMyDescription(ctx context.Context, request *types.GetMyDescriptionRequest) (r *types.BotDescription, err error) {
	res, err := makeRequest[*types.BotDescription](ctx, b.options.Client, b.baseURL, "getMyDescription", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to change the bot's short description, which is shown on the bot's profile page and is sent together with the link when users share the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#setmyshortdescription
func (b *Bot) SetMyShortDescription(ctx context.Context, request *types.SetMyShortDescriptionRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setMyShortDescription", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get the current bot short description for the given user language. Returns BotShortDescription on success.
//
// https://core.telegram.org/bots/api#getmyshortdescription
func (b *Bot) GetMyShortDescription(ctx context.Context, request *types.GetMyShortDescriptionRequest) (r *types.BotShortDescription, err error) {
	res, err := makeRequest[*types.BotShortDescription](ctx, b.options.Client, b.baseURL, "getMyShortDescription", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to change the bot's menu button in a private chat, or the default menu button. Returns True on success.
//
// https://core.telegram.org/bots/api#setchatmenubutton
func (b *Bot) SetChatMenuButton(ctx context.Context, request *types.SetChatMenuButtonRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setChatMenuButton", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get the current value of the bot's menu button in a private chat, or the default menu button. Returns MenuButton on success.
//
// https://core.telegram.org/bots/api#getchatmenubutton
func (b *Bot) GetChatMenuButton(ctx context.Context, request *types.GetChatMenuButtonRequest) (r types.MenuButton, err error) {
	res, err := makeRequest[types.MenuButton](ctx, b.options.Client, b.baseURL, "getChatMenuButton", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels. These rights will be suggested to users, but they are free to modify the list before adding the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#setmydefaultadministratorrights
func (b *Bot) SetMyDefaultAdministratorRights(ctx context.Context, request *types.SetMyDefaultAdministratorRightsRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setMyDefaultAdministratorRights", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get the current default administrator rights of the bot. Returns ChatAdministratorRights on success.
//
// https://core.telegram.org/bots/api#getmydefaultadministratorrights
func (b *Bot) GetMyDefaultAdministratorRights(ctx context.Context, request *types.GetMyDefaultAdministratorRightsRequest) (r *types.ChatAdministratorRights, err error) {
	res, err := makeRequest[*types.ChatAdministratorRights](ctx, b.options.Client, b.baseURL, "getMyDefaultAdministratorRights", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
//
// https://core.telegram.org/bots/api#editmessagetext
func (b *Bot) EditMessageText(ctx context.Context, request *types.EditMessageTextRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "editMessageText", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
//
// https://core.telegram.org/bots/api#editmessagecaption
func (b *Bot) EditMessageCaption(ctx context.Context, request *types.EditMessageCaptionRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "editMessageCaption", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to edit animation, audio, document, photo, or video messages, or to add media to text messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo or a video otherwise. When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
//
// https://core.telegram.org/bots/api#editmessagemedia
func (b *Bot) EditMessageMedia(ctx context.Context, request *types.EditMessageMediaRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "editMessageMedia", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
//
// https://core.telegram.org/bots/api#editmessagelivelocation
func (b *Bot) EditMessageLiveLocation(ctx context.Context, request *types.EditMessageLiveLocationRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "editMessageLiveLocation", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
//
// https://core.telegram.org/bots/api#stopmessagelivelocation
func (b *Bot) StopMessageLiveLocation(ctx context.Context, request *types.StopMessageLiveLocationRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "stopMessageLiveLocation", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
//
// https://core.telegram.org/bots/api#editmessagereplymarkup
func (b *Bot) EditMessageReplyMarkup(ctx context.Context, request *types.EditMessageReplyMarkupRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "editMessageReplyMarkup", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to stop a poll which was sent by the bot. On success, the stopped Poll is returned.
//
// https://core.telegram.org/bots/api#stoppoll
func (b *Bot) StopPoll(ctx context.Context, request *types.StopPollRequest) (r *types.Poll, err error) {
	res, err := makeRequest[*types.Poll](ctx, b.options.Client, b.baseURL, "stopPoll", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to delete a message, including service messages, with the following limitations:
//
// - A message can only be deleted if it was sent less than 48 hours ago.
//
// - Service messages about a supergroup, channel, or forum topic creation can't be deleted.
//
// - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
//
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
//
// - Bots can delete incoming messages in private chats.
//
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
//
// - If the bot is an administrator of a group, it can delete any message there.
//
// - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
//
// Returns True on success.
//
// https://core.telegram.org/bots/api#deletemessage
func (b *Bot) DeleteMessage(ctx context.Context, request *types.DeleteMessageRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "deleteMessage", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to delete multiple messages simultaneously. If some of the specified messages can't be found, they are skipped. Returns True on success.
//
// https://core.telegram.org/bots/api#deletemessages
func (b *Bot) DeleteMessages(ctx context.Context, request *types.DeleteMessagesRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "deleteMessages", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
//
// https://core.telegram.org/bots/api#sendsticker
func (b *Bot) SendSticker(ctx context.Context, request *types.SendStickerRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendSticker", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get a sticker set. On success, a StickerSet object is returned.
//
// https://core.telegram.org/bots/api#getstickerset
func (b *Bot) GetStickerSet(ctx context.Context, request *types.GetStickerSetRequest) (r *types.StickerSet, err error) {
	res, err := makeRequest[*types.StickerSet](ctx, b.options.Client, b.baseURL, "getStickerSet", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.
//
// https://core.telegram.org/bots/api#getcustomemojistickers
func (b *Bot) GetCustomEmojiStickers(ctx context.Context, request *types.GetCustomEmojiStickersRequest) (r []types.Sticker, err error) {
	res, err := makeRequest[[]types.Sticker](ctx, b.options.Client, b.baseURL, "getCustomEmojiStickers", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to upload a file with a sticker for later use in the createNewStickerSet, addStickerToSet, or replaceStickerInSet methods (the file can be used multiple times). Returns the uploaded File on success.
//
// https://core.telegram.org/bots/api#uploadstickerfile
func (b *Bot) UploadStickerFile(ctx context.Context, request *types.UploadStickerFileRequest) (r *types.File, err error) {
	res, err := makeRequest[*types.File](ctx, b.options.Client, b.baseURL, "uploadStickerFile", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. Returns True on success.
//
// https://core.telegram.org/bots/api#createnewstickerset
func (b *Bot) CreateNewStickerSet(ctx context.Context, request *types.CreateNewStickerSetRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "createNewStickerSet", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to add a new sticker to a set created by the bot. Emoji sticker sets can have up to 200 stickers. Other sticker sets can have up to 120 stickers. Returns True on success.
//
// https://core.telegram.org/bots/api#addstickertoset
func (b *Bot) AddStickerToSet(ctx context.Context, request *types.AddStickerToSetRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "addStickerToSet", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to move a sticker in a set created by the bot to a specific position. Returns True on success.
//
// https://core.telegram.org/bots/api#setstickerpositioninset
func (b *Bot) SetStickerPositionInSet(ctx context.Context, request *types.SetStickerPositionInSetRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setStickerPositionInSet", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to delete a sticker from a set created by the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#deletestickerfromset
func (b *Bot) DeleteStickerFromSet(ctx context.Context, request *types.DeleteStickerFromSetRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "deleteStickerFromSet", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to replace an existing sticker in a sticker set with a new one. The method is equivalent to calling deleteStickerFromSet, then addStickerToSet, then setStickerPositionInSet. Returns True on success.
//
// https://core.telegram.org/bots/api#replacestickerinset
func (b *Bot) ReplaceStickerInSet(ctx context.Context, request *types.ReplaceStickerInSetRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "replaceStickerInSet", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to change the list of emoji assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#setstickeremojilist
func (b *Bot) SetStickerEmojiList(ctx context.Context, request *types.SetStickerEmojiListRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setStickerEmojiList", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to change search keywords assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#setstickerkeywords
func (b *Bot) SetStickerKeywords(ctx context.Context, request *types.SetStickerKeywordsRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setStickerKeywords", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to change the mask position of a mask sticker. The sticker must belong to a sticker set that was created by the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#setstickermaskposition
func (b *Bot) SetStickerMaskPosition(ctx context.Context, request *types.SetStickerMaskPositionRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setStickerMaskPosition", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to set the title of a created sticker set. Returns True on success.
//
// https://core.telegram.org/bots/api#setstickersettitle
func (b *Bot) SetStickerSetTitle(ctx context.Context, request *types.SetStickerSetTitleRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setStickerSetTitle", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to set the thumbnail of a regular or mask sticker set. The format of the thumbnail file must match the format of the stickers in the set. Returns True on success.
//
// https://core.telegram.org/bots/api#setstickersetthumbnail
func (b *Bot) SetStickerSetThumbnail(ctx context.Context, request *types.SetStickerSetThumbnailRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setStickerSetThumbnail", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to set the thumbnail of a custom emoji sticker set. Returns True on success.
//
// https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
func (b *Bot) SetCustomEmojiStickerSetThumbnail(ctx context.Context, request *types.SetCustomEmojiStickerSetThumbnailRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setCustomEmojiStickerSetThumbnail", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to delete a sticker set that was created by the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#deletestickerset
func (b *Bot) DeleteStickerSet(ctx context.Context, request *types.DeleteStickerSetRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "deleteStickerSet", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Returns the list of gifts that can be sent by the bot to users and channel chats. Requires no parameters. Returns a Gifts object.
//
// https://core.telegram.org/bots/api#getavailablegifts
func (b *Bot) GetAvailableGifts(ctx context.Context) (r *types.Gifts, err error) {
	res, err := makeRequest[*types.Gifts](ctx, b.options.Client, b.baseURL, "getAvailableGifts", nil)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Sends a gift to the given user or channel chat. The gift can't be converted to Telegram Stars by the receiver. Returns True on success.
//
// https://core.telegram.org/bots/api#sendgift
func (b *Bot) SendGift(ctx context.Context, request *types.SendGiftRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "sendGift", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Verifies a user on behalf of the organization which is represented by the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#verifyuser
func (b *Bot) VerifyUser(ctx context.Context, request *types.VerifyUserRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "verifyUser", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Verifies a chat on behalf of the organization which is represented by the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#verifychat
func (b *Bot) VerifyChat(ctx context.Context, request *types.VerifyChatRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "verifyChat", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Removes verification from a user who is currently verified on behalf of the organization represented by the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#removeuserverification
func (b *Bot) RemoveUserVerification(ctx context.Context, request *types.RemoveUserVerificationRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "removeUserVerification", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Removes verification from a chat that is currently verified on behalf of the organization represented by the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#removechatverification
func (b *Bot) RemoveChatVerification(ctx context.Context, request *types.RemoveChatVerificationRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "removeChatVerification", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send answers to an inline query. On success, True is returned.
//
// No more than 50 results per query are allowed.
//
// https://core.telegram.org/bots/api#answerinlinequery
func (b *Bot) AnswerInlineQuery(ctx context.Context, request *types.AnswerInlineQueryRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "answerInlineQuery", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated. On success, a SentWebAppMessage object is returned.
//
// https://core.telegram.org/bots/api#answerwebappquery
func (b *Bot) AnswerWebAppQuery(ctx context.Context, request *types.AnswerWebAppQueryRequest) (r *types.SentWebAppMessage, err error) {
	res, err := makeRequest[*types.SentWebAppMessage](ctx, b.options.Client, b.baseURL, "answerWebAppQuery", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Stores a message that can be sent by a user of a Mini App. Returns a PreparedInlineMessage object.
//
// https://core.telegram.org/bots/api#savepreparedinlinemessage
func (b *Bot) SavePreparedInlineMessage(ctx context.Context, request *types.SavePreparedInlineMessageRequest) (r *types.PreparedInlineMessage, err error) {
	res, err := makeRequest[*types.PreparedInlineMessage](ctx, b.options.Client, b.baseURL, "savePreparedInlineMessage", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send invoices. On success, the sent Message is returned.
//
// https://core.telegram.org/bots/api#sendinvoice
func (b *Bot) SendInvoice(ctx context.Context, request *types.SendInvoiceRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendInvoice", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to create a link for an invoice. Returns the created invoice link as String on success.
//
// https://core.telegram.org/bots/api#createinvoicelink
func (b *Bot) CreateInvoiceLink(ctx context.Context, request *types.CreateInvoiceLinkRequest) (r string, err error) {
	res, err := makeRequest[string](ctx, b.options.Client, b.baseURL, "createInvoiceLink", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
//
// https://core.telegram.org/bots/api#answershippingquery
func (b *Bot) AnswerShippingQuery(ctx context.Context, request *types.AnswerShippingQueryRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "answerShippingQuery", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
//
// https://core.telegram.org/bots/api#answerprecheckoutquery
func (b *Bot) AnswerPreCheckoutQuery(ctx context.Context, request *types.AnswerPreCheckoutQueryRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "answerPreCheckoutQuery", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Returns the bot's Telegram Star transactions in chronological order. On success, returns a StarTransactions object.
//
// https://core.telegram.org/bots/api#getstartransactions
func (b *Bot) GetStarTransactions(ctx context.Context, request *types.GetStarTransactionsRequest) (r *types.StarTransactions, err error) {
	res, err := makeRequest[*types.StarTransactions](ctx, b.options.Client, b.baseURL, "getStarTransactions", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Refunds a successful payment in Telegram Stars. Returns True on success.
//
// https://core.telegram.org/bots/api#refundstarpayment
func (b *Bot) RefundStarPayment(ctx context.Context, request *types.RefundStarPaymentRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "refundStarPayment", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Allows the bot to cancel or re-enable extension of a subscription paid in Telegram Stars. Returns True on success.
//
// https://core.telegram.org/bots/api#edituserstarsubscription
func (b *Bot) EditUserStarSubscription(ctx context.Context, request *types.EditUserStarSubscriptionRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "editUserStarSubscription", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
//
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
//
// https://core.telegram.org/bots/api#setpassportdataerrors
func (b *Bot) SetPassportDataErrors(ctx context.Context, request *types.SetPassportDataErrorsRequest) (r bool, err error) {
	res, err := makeRequest[bool](ctx, b.options.Client, b.baseURL, "setPassportDataErrors", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to send a game. On success, the sent Message is returned.
//
// https://core.telegram.org/bots/api#sendgame
func (b *Bot) SendGame(ctx context.Context, request *types.SendGameRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "sendGame", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to set the score of the specified user in a game message. On success, if the message is not an inline message, the Message is returned, otherwise True is returned. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
//
// https://core.telegram.org/bots/api#setgamescore
func (b *Bot) SetGameScore(ctx context.Context, request *types.SetGameScoreRequest) (r *types.Message, err error) {
	res, err := makeRequest[*types.Message](ctx, b.options.Client, b.baseURL, "setGameScore", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}

// Use this method to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of GameHighScore objects.
//
// https://core.telegram.org/bots/api#getgamehighscores
func (b *Bot) GetGameHighScores(ctx context.Context, request *types.GetGameHighScoresRequest) (r []types.GameHighScore, err error) {
	res, err := makeRequest[[]types.GameHighScore](ctx, b.options.Client, b.baseURL, "getGameHighScores", request)

	if err != nil {
		return r, err
	}

	return res.Result, nil
}
