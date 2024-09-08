package chat

import "qingyu-wf/global"

func Create(table ContentChatTable) error {
	return global.MySql.Create(&table).Error
}
func FindList(id string) (c []ContentChatTable) {
	global.MySql.Where("sender = ? or receiver = ?", id, id).Find(&c)
	return
}
