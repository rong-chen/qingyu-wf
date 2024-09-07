package chat

import "qingyu-wf/global"

func Create(table ContentChatTable) error {
	return global.MySql.Create(&table).Error
}
