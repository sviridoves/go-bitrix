package RecentUserChats

import "github.com/nightwriter/go-bitrix/types"

type RecentUserChatsResponce struct {
	types.Response
	Result []RecentUserChats `json:"result"`
}
