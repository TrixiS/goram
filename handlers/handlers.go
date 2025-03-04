package handlers

import "github.com/TrixiS/goram"
import "context"

type routerHandlers[T any] struct {
	filters  []Filter[T] // router-level filters for this update
	handlers []handler[T]
}

type handlers struct {
	message                 routerHandlers[*goram.Message]
	editedMessage           routerHandlers[*goram.Message]
	channelPost             routerHandlers[*goram.Message]
	editedChannelPost       routerHandlers[*goram.Message]
	businessConnection      routerHandlers[*goram.BusinessConnection]
	businessMessage         routerHandlers[*goram.Message]
	editedBusinessMessage   routerHandlers[*goram.Message]
	deletedBusinessMessages routerHandlers[*goram.BusinessMessagesDeleted]
	messageReaction         routerHandlers[*goram.MessageReactionUpdated]
	messageReactionCount    routerHandlers[*goram.MessageReactionCountUpdated]
	inlineQuery             routerHandlers[*goram.InlineQuery]
	chosenInlineResult      routerHandlers[*goram.ChosenInlineResult]
	callbackQuery           routerHandlers[*goram.CallbackQuery]
	shippingQuery           routerHandlers[*goram.ShippingQuery]
	preCheckoutQuery        routerHandlers[*goram.PreCheckoutQuery]
	purchasedPaidMedia      routerHandlers[*goram.PaidMediaPurchased]
	poll                    routerHandlers[*goram.Poll]
	pollAnswer              routerHandlers[*goram.PollAnswer]
	myChatMember            routerHandlers[*goram.ChatMemberUpdated]
	chatMember              routerHandlers[*goram.ChatMemberUpdated]
	chatJoinRequest         routerHandlers[*goram.ChatJoinRequest]
	chatBoost               routerHandlers[*goram.ChatBoostUpdated]
	removedChatBoost        routerHandlers[*goram.ChatBoostRemoved]
}

// Add Message handler with provided filters
func (r *Router) Message(handlerFunc Func[*goram.Message], filters ...Filter[*goram.Message]) *Router {
	h := handler[*goram.Message]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.message.handlers = append(r.handlers.message.handlers, h)
	return r
}

// Add EditedMessage handler with provided filters
func (r *Router) EditedMessage(handlerFunc Func[*goram.Message], filters ...Filter[*goram.Message]) *Router {
	h := handler[*goram.Message]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.editedMessage.handlers = append(r.handlers.editedMessage.handlers, h)
	return r
}

// Add ChannelPost handler with provided filters
func (r *Router) ChannelPost(handlerFunc Func[*goram.Message], filters ...Filter[*goram.Message]) *Router {
	h := handler[*goram.Message]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.channelPost.handlers = append(r.handlers.channelPost.handlers, h)
	return r
}

// Add EditedChannelPost handler with provided filters
func (r *Router) EditedChannelPost(handlerFunc Func[*goram.Message], filters ...Filter[*goram.Message]) *Router {
	h := handler[*goram.Message]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.editedChannelPost.handlers = append(r.handlers.editedChannelPost.handlers, h)
	return r
}

// Add BusinessConnection handler with provided filters
func (r *Router) BusinessConnection(handlerFunc Func[*goram.BusinessConnection], filters ...Filter[*goram.BusinessConnection]) *Router {
	h := handler[*goram.BusinessConnection]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.businessConnection.handlers = append(r.handlers.businessConnection.handlers, h)
	return r
}

// Add BusinessMessage handler with provided filters
func (r *Router) BusinessMessage(handlerFunc Func[*goram.Message], filters ...Filter[*goram.Message]) *Router {
	h := handler[*goram.Message]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.businessMessage.handlers = append(r.handlers.businessMessage.handlers, h)
	return r
}

// Add EditedBusinessMessage handler with provided filters
func (r *Router) EditedBusinessMessage(handlerFunc Func[*goram.Message], filters ...Filter[*goram.Message]) *Router {
	h := handler[*goram.Message]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.editedBusinessMessage.handlers = append(r.handlers.editedBusinessMessage.handlers, h)
	return r
}

// Add DeletedBusinessMessages handler with provided filters
func (r *Router) DeletedBusinessMessages(handlerFunc Func[*goram.BusinessMessagesDeleted], filters ...Filter[*goram.BusinessMessagesDeleted]) *Router {
	h := handler[*goram.BusinessMessagesDeleted]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.deletedBusinessMessages.handlers = append(r.handlers.deletedBusinessMessages.handlers, h)
	return r
}

// Add MessageReaction handler with provided filters
func (r *Router) MessageReaction(handlerFunc Func[*goram.MessageReactionUpdated], filters ...Filter[*goram.MessageReactionUpdated]) *Router {
	h := handler[*goram.MessageReactionUpdated]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.messageReaction.handlers = append(r.handlers.messageReaction.handlers, h)
	return r
}

// Add MessageReactionCount handler with provided filters
func (r *Router) MessageReactionCount(handlerFunc Func[*goram.MessageReactionCountUpdated], filters ...Filter[*goram.MessageReactionCountUpdated]) *Router {
	h := handler[*goram.MessageReactionCountUpdated]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.messageReactionCount.handlers = append(r.handlers.messageReactionCount.handlers, h)
	return r
}

// Add InlineQuery handler with provided filters
func (r *Router) InlineQuery(handlerFunc Func[*goram.InlineQuery], filters ...Filter[*goram.InlineQuery]) *Router {
	h := handler[*goram.InlineQuery]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.inlineQuery.handlers = append(r.handlers.inlineQuery.handlers, h)
	return r
}

// Add ChosenInlineResult handler with provided filters
func (r *Router) ChosenInlineResult(handlerFunc Func[*goram.ChosenInlineResult], filters ...Filter[*goram.ChosenInlineResult]) *Router {
	h := handler[*goram.ChosenInlineResult]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.chosenInlineResult.handlers = append(r.handlers.chosenInlineResult.handlers, h)
	return r
}

// Add CallbackQuery handler with provided filters
func (r *Router) CallbackQuery(handlerFunc Func[*goram.CallbackQuery], filters ...Filter[*goram.CallbackQuery]) *Router {
	h := handler[*goram.CallbackQuery]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.callbackQuery.handlers = append(r.handlers.callbackQuery.handlers, h)
	return r
}

// Add ShippingQuery handler with provided filters
func (r *Router) ShippingQuery(handlerFunc Func[*goram.ShippingQuery], filters ...Filter[*goram.ShippingQuery]) *Router {
	h := handler[*goram.ShippingQuery]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.shippingQuery.handlers = append(r.handlers.shippingQuery.handlers, h)
	return r
}

// Add PreCheckoutQuery handler with provided filters
func (r *Router) PreCheckoutQuery(handlerFunc Func[*goram.PreCheckoutQuery], filters ...Filter[*goram.PreCheckoutQuery]) *Router {
	h := handler[*goram.PreCheckoutQuery]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.preCheckoutQuery.handlers = append(r.handlers.preCheckoutQuery.handlers, h)
	return r
}

// Add PurchasedPaidMedia handler with provided filters
func (r *Router) PurchasedPaidMedia(handlerFunc Func[*goram.PaidMediaPurchased], filters ...Filter[*goram.PaidMediaPurchased]) *Router {
	h := handler[*goram.PaidMediaPurchased]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.purchasedPaidMedia.handlers = append(r.handlers.purchasedPaidMedia.handlers, h)
	return r
}

// Add Poll handler with provided filters
func (r *Router) Poll(handlerFunc Func[*goram.Poll], filters ...Filter[*goram.Poll]) *Router {
	h := handler[*goram.Poll]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.poll.handlers = append(r.handlers.poll.handlers, h)
	return r
}

// Add PollAnswer handler with provided filters
func (r *Router) PollAnswer(handlerFunc Func[*goram.PollAnswer], filters ...Filter[*goram.PollAnswer]) *Router {
	h := handler[*goram.PollAnswer]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.pollAnswer.handlers = append(r.handlers.pollAnswer.handlers, h)
	return r
}

// Add MyChatMember handler with provided filters
func (r *Router) MyChatMember(handlerFunc Func[*goram.ChatMemberUpdated], filters ...Filter[*goram.ChatMemberUpdated]) *Router {
	h := handler[*goram.ChatMemberUpdated]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.myChatMember.handlers = append(r.handlers.myChatMember.handlers, h)
	return r
}

// Add ChatMember handler with provided filters
func (r *Router) ChatMember(handlerFunc Func[*goram.ChatMemberUpdated], filters ...Filter[*goram.ChatMemberUpdated]) *Router {
	h := handler[*goram.ChatMemberUpdated]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.chatMember.handlers = append(r.handlers.chatMember.handlers, h)
	return r
}

// Add ChatJoinRequest handler with provided filters
func (r *Router) ChatJoinRequest(handlerFunc Func[*goram.ChatJoinRequest], filters ...Filter[*goram.ChatJoinRequest]) *Router {
	h := handler[*goram.ChatJoinRequest]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.chatJoinRequest.handlers = append(r.handlers.chatJoinRequest.handlers, h)
	return r
}

// Add ChatBoost handler with provided filters
func (r *Router) ChatBoost(handlerFunc Func[*goram.ChatBoostUpdated], filters ...Filter[*goram.ChatBoostUpdated]) *Router {
	h := handler[*goram.ChatBoostUpdated]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.chatBoost.handlers = append(r.handlers.chatBoost.handlers, h)
	return r
}

// Add RemovedChatBoost handler with provided filters
func (r *Router) RemovedChatBoost(handlerFunc Func[*goram.ChatBoostRemoved], filters ...Filter[*goram.ChatBoostRemoved]) *Router {
	h := handler[*goram.ChatBoostRemoved]{
		cb:      handlerFunc,
		filters: filters,
	}

	r.handlers.removedChatBoost.handlers = append(r.handlers.removedChatBoost.handlers, h)
	return r
}

// Add router-level filter(s) to Message update
func (r *Router) FilterMessage(filters ...Filter[*goram.Message]) *Router {
	r.handlers.message.filters = append(r.handlers.message.filters, filters...)
	return r
}

// Add router-level filter(s) to EditedMessage update
func (r *Router) FilterEditedMessage(filters ...Filter[*goram.Message]) *Router {
	r.handlers.editedMessage.filters = append(r.handlers.editedMessage.filters, filters...)
	return r
}

// Add router-level filter(s) to ChannelPost update
func (r *Router) FilterChannelPost(filters ...Filter[*goram.Message]) *Router {
	r.handlers.channelPost.filters = append(r.handlers.channelPost.filters, filters...)
	return r
}

// Add router-level filter(s) to EditedChannelPost update
func (r *Router) FilterEditedChannelPost(filters ...Filter[*goram.Message]) *Router {
	r.handlers.editedChannelPost.filters = append(r.handlers.editedChannelPost.filters, filters...)
	return r
}

// Add router-level filter(s) to BusinessConnection update
func (r *Router) FilterBusinessConnection(filters ...Filter[*goram.BusinessConnection]) *Router {
	r.handlers.businessConnection.filters = append(r.handlers.businessConnection.filters, filters...)
	return r
}

// Add router-level filter(s) to BusinessMessage update
func (r *Router) FilterBusinessMessage(filters ...Filter[*goram.Message]) *Router {
	r.handlers.businessMessage.filters = append(r.handlers.businessMessage.filters, filters...)
	return r
}

// Add router-level filter(s) to EditedBusinessMessage update
func (r *Router) FilterEditedBusinessMessage(filters ...Filter[*goram.Message]) *Router {
	r.handlers.editedBusinessMessage.filters = append(r.handlers.editedBusinessMessage.filters, filters...)
	return r
}

// Add router-level filter(s) to DeletedBusinessMessages update
func (r *Router) FilterDeletedBusinessMessages(filters ...Filter[*goram.BusinessMessagesDeleted]) *Router {
	r.handlers.deletedBusinessMessages.filters = append(r.handlers.deletedBusinessMessages.filters, filters...)
	return r
}

// Add router-level filter(s) to MessageReaction update
func (r *Router) FilterMessageReaction(filters ...Filter[*goram.MessageReactionUpdated]) *Router {
	r.handlers.messageReaction.filters = append(r.handlers.messageReaction.filters, filters...)
	return r
}

// Add router-level filter(s) to MessageReactionCount update
func (r *Router) FilterMessageReactionCount(filters ...Filter[*goram.MessageReactionCountUpdated]) *Router {
	r.handlers.messageReactionCount.filters = append(r.handlers.messageReactionCount.filters, filters...)
	return r
}

// Add router-level filter(s) to InlineQuery update
func (r *Router) FilterInlineQuery(filters ...Filter[*goram.InlineQuery]) *Router {
	r.handlers.inlineQuery.filters = append(r.handlers.inlineQuery.filters, filters...)
	return r
}

// Add router-level filter(s) to ChosenInlineResult update
func (r *Router) FilterChosenInlineResult(filters ...Filter[*goram.ChosenInlineResult]) *Router {
	r.handlers.chosenInlineResult.filters = append(r.handlers.chosenInlineResult.filters, filters...)
	return r
}

// Add router-level filter(s) to CallbackQuery update
func (r *Router) FilterCallbackQuery(filters ...Filter[*goram.CallbackQuery]) *Router {
	r.handlers.callbackQuery.filters = append(r.handlers.callbackQuery.filters, filters...)
	return r
}

// Add router-level filter(s) to ShippingQuery update
func (r *Router) FilterShippingQuery(filters ...Filter[*goram.ShippingQuery]) *Router {
	r.handlers.shippingQuery.filters = append(r.handlers.shippingQuery.filters, filters...)
	return r
}

// Add router-level filter(s) to PreCheckoutQuery update
func (r *Router) FilterPreCheckoutQuery(filters ...Filter[*goram.PreCheckoutQuery]) *Router {
	r.handlers.preCheckoutQuery.filters = append(r.handlers.preCheckoutQuery.filters, filters...)
	return r
}

// Add router-level filter(s) to PurchasedPaidMedia update
func (r *Router) FilterPurchasedPaidMedia(filters ...Filter[*goram.PaidMediaPurchased]) *Router {
	r.handlers.purchasedPaidMedia.filters = append(r.handlers.purchasedPaidMedia.filters, filters...)
	return r
}

// Add router-level filter(s) to Poll update
func (r *Router) FilterPoll(filters ...Filter[*goram.Poll]) *Router {
	r.handlers.poll.filters = append(r.handlers.poll.filters, filters...)
	return r
}

// Add router-level filter(s) to PollAnswer update
func (r *Router) FilterPollAnswer(filters ...Filter[*goram.PollAnswer]) *Router {
	r.handlers.pollAnswer.filters = append(r.handlers.pollAnswer.filters, filters...)
	return r
}

// Add router-level filter(s) to MyChatMember update
func (r *Router) FilterMyChatMember(filters ...Filter[*goram.ChatMemberUpdated]) *Router {
	r.handlers.myChatMember.filters = append(r.handlers.myChatMember.filters, filters...)
	return r
}

// Add router-level filter(s) to ChatMember update
func (r *Router) FilterChatMember(filters ...Filter[*goram.ChatMemberUpdated]) *Router {
	r.handlers.chatMember.filters = append(r.handlers.chatMember.filters, filters...)
	return r
}

// Add router-level filter(s) to ChatJoinRequest update
func (r *Router) FilterChatJoinRequest(filters ...Filter[*goram.ChatJoinRequest]) *Router {
	r.handlers.chatJoinRequest.filters = append(r.handlers.chatJoinRequest.filters, filters...)
	return r
}

// Add router-level filter(s) to ChatBoost update
func (r *Router) FilterChatBoost(filters ...Filter[*goram.ChatBoostUpdated]) *Router {
	r.handlers.chatBoost.filters = append(r.handlers.chatBoost.filters, filters...)
	return r
}

// Add router-level filter(s) to RemovedChatBoost update
func (r *Router) FilterRemovedChatBoost(filters ...Filter[*goram.ChatBoostRemoved]) *Router {
	r.handlers.removedChatBoost.filters = append(r.handlers.removedChatBoost.filters, filters...)
	return r
}

func (r *Router) callMessageHandlers(ctx context.Context, bot *goram.Bot, update *goram.Message, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.message.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.message.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callEditedMessageHandlers(ctx context.Context, bot *goram.Bot, update *goram.Message, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.editedMessage.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.editedMessage.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callChannelPostHandlers(ctx context.Context, bot *goram.Bot, update *goram.Message, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.channelPost.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.channelPost.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callEditedChannelPostHandlers(ctx context.Context, bot *goram.Bot, update *goram.Message, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.editedChannelPost.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.editedChannelPost.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callBusinessConnectionHandlers(ctx context.Context, bot *goram.Bot, update *goram.BusinessConnection, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.businessConnection.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.businessConnection.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callBusinessMessageHandlers(ctx context.Context, bot *goram.Bot, update *goram.Message, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.businessMessage.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.businessMessage.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callEditedBusinessMessageHandlers(ctx context.Context, bot *goram.Bot, update *goram.Message, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.editedBusinessMessage.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.editedBusinessMessage.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callDeletedBusinessMessagesHandlers(ctx context.Context, bot *goram.Bot, update *goram.BusinessMessagesDeleted, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.deletedBusinessMessages.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.deletedBusinessMessages.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callMessageReactionHandlers(ctx context.Context, bot *goram.Bot, update *goram.MessageReactionUpdated, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.messageReaction.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.messageReaction.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callMessageReactionCountHandlers(ctx context.Context, bot *goram.Bot, update *goram.MessageReactionCountUpdated, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.messageReactionCount.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.messageReactionCount.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callInlineQueryHandlers(ctx context.Context, bot *goram.Bot, update *goram.InlineQuery, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.inlineQuery.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.inlineQuery.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callChosenInlineResultHandlers(ctx context.Context, bot *goram.Bot, update *goram.ChosenInlineResult, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.chosenInlineResult.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.chosenInlineResult.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callCallbackQueryHandlers(ctx context.Context, bot *goram.Bot, update *goram.CallbackQuery, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.callbackQuery.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.callbackQuery.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callShippingQueryHandlers(ctx context.Context, bot *goram.Bot, update *goram.ShippingQuery, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.shippingQuery.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.shippingQuery.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callPreCheckoutQueryHandlers(ctx context.Context, bot *goram.Bot, update *goram.PreCheckoutQuery, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.preCheckoutQuery.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.preCheckoutQuery.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callPurchasedPaidMediaHandlers(ctx context.Context, bot *goram.Bot, update *goram.PaidMediaPurchased, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.purchasedPaidMedia.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.purchasedPaidMedia.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callPollHandlers(ctx context.Context, bot *goram.Bot, update *goram.Poll, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.poll.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.poll.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callPollAnswerHandlers(ctx context.Context, bot *goram.Bot, update *goram.PollAnswer, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.pollAnswer.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.pollAnswer.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callMyChatMemberHandlers(ctx context.Context, bot *goram.Bot, update *goram.ChatMemberUpdated, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.myChatMember.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.myChatMember.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callChatMemberHandlers(ctx context.Context, bot *goram.Bot, update *goram.ChatMemberUpdated, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.chatMember.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.chatMember.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callChatJoinRequestHandlers(ctx context.Context, bot *goram.Bot, update *goram.ChatJoinRequest, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.chatJoinRequest.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.chatJoinRequest.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callChatBoostHandlers(ctx context.Context, bot *goram.Bot, update *goram.ChatBoostUpdated, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.chatBoost.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.chatBoost.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}

func (r *Router) callRemovedChatBoostHandlers(ctx context.Context, bot *goram.Bot, update *goram.ChatBoostRemoved, data Data) (bool, error) {
	queue := make([]*Router, 0, len(r.children)+1)
	queue = append(queue, r)

queueLoop:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, filter := range current.handlers.removedChatBoost.filters {
			if !filter(ctx, bot, update, data) {
				continue queueLoop
			}
		}

		found, err := callHandlers(ctx, bot, current.handlers.removedChatBoost.handlers, update, data)

		if found {
			return found, err
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return false, nil
}
func (r *Router) feedUpdate(ctx context.Context, bot *goram.Bot, update *goram.Update, data Data) (bool, error) {
	if update.Message != nil {
		return r.callMessageHandlers(ctx, bot, update.Message, data)
	}
	if update.EditedMessage != nil {
		return r.callEditedMessageHandlers(ctx, bot, update.EditedMessage, data)
	}
	if update.ChannelPost != nil {
		return r.callChannelPostHandlers(ctx, bot, update.ChannelPost, data)
	}
	if update.EditedChannelPost != nil {
		return r.callEditedChannelPostHandlers(ctx, bot, update.EditedChannelPost, data)
	}
	if update.BusinessConnection != nil {
		return r.callBusinessConnectionHandlers(ctx, bot, update.BusinessConnection, data)
	}
	if update.BusinessMessage != nil {
		return r.callBusinessMessageHandlers(ctx, bot, update.BusinessMessage, data)
	}
	if update.EditedBusinessMessage != nil {
		return r.callEditedBusinessMessageHandlers(ctx, bot, update.EditedBusinessMessage, data)
	}
	if update.DeletedBusinessMessages != nil {
		return r.callDeletedBusinessMessagesHandlers(ctx, bot, update.DeletedBusinessMessages, data)
	}
	if update.MessageReaction != nil {
		return r.callMessageReactionHandlers(ctx, bot, update.MessageReaction, data)
	}
	if update.MessageReactionCount != nil {
		return r.callMessageReactionCountHandlers(ctx, bot, update.MessageReactionCount, data)
	}
	if update.InlineQuery != nil {
		return r.callInlineQueryHandlers(ctx, bot, update.InlineQuery, data)
	}
	if update.ChosenInlineResult != nil {
		return r.callChosenInlineResultHandlers(ctx, bot, update.ChosenInlineResult, data)
	}
	if update.CallbackQuery != nil {
		return r.callCallbackQueryHandlers(ctx, bot, update.CallbackQuery, data)
	}
	if update.ShippingQuery != nil {
		return r.callShippingQueryHandlers(ctx, bot, update.ShippingQuery, data)
	}
	if update.PreCheckoutQuery != nil {
		return r.callPreCheckoutQueryHandlers(ctx, bot, update.PreCheckoutQuery, data)
	}
	if update.PurchasedPaidMedia != nil {
		return r.callPurchasedPaidMediaHandlers(ctx, bot, update.PurchasedPaidMedia, data)
	}
	if update.Poll != nil {
		return r.callPollHandlers(ctx, bot, update.Poll, data)
	}
	if update.PollAnswer != nil {
		return r.callPollAnswerHandlers(ctx, bot, update.PollAnswer, data)
	}
	if update.MyChatMember != nil {
		return r.callMyChatMemberHandlers(ctx, bot, update.MyChatMember, data)
	}
	if update.ChatMember != nil {
		return r.callChatMemberHandlers(ctx, bot, update.ChatMember, data)
	}
	if update.ChatJoinRequest != nil {
		return r.callChatJoinRequestHandlers(ctx, bot, update.ChatJoinRequest, data)
	}
	if update.ChatBoost != nil {
		return r.callChatBoostHandlers(ctx, bot, update.ChatBoost, data)
	}
	if update.RemovedChatBoost != nil {
		return r.callRemovedChatBoostHandlers(ctx, bot, update.RemovedChatBoost, data)
	}
	return false, nil
}
