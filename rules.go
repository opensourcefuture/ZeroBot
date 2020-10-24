package ZeroBot

import "strings"

// 是否为消息事件
func IsMessage() func(event Event) bool {
	return func(event Event) bool {
		return event.PostType == "message"
	}
}

// 是否含有前缀,使用时请确保该事件为消息事件
func IsPrefix(prefixs ...string) func(event Event) bool {
	return func(event Event) bool {
		for _, prefix := range prefixs {
			if strings.HasPrefix(event.RawMessage, prefix) { // 只要有一个前缀就行了
				return true
			}
		}
		return false
	}
}

// 是否为通知事件
func IsNotice() func(event Event) bool {
	return func(event Event) bool {
		return event.PostType == "notice"
	}
}

// 是否为请求事件
func IsRequest() func(event Event) bool {
	return func(event Event) bool {
		return event.PostType == "request"
	}
}

// 是否为元事件
func IsMetaEvent() func(event Event) bool {
	return func(event Event) bool {
		return event.PostType == "meta_event"
	}
}

func CheckUser(userId int64) func(event Event) bool {
	return func(event Event) bool {
		return event.UserID == userId
	}
}