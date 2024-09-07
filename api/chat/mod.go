package chat

import "gorm.io/gorm"

// ContentChatTable 存储用户聊天内容
type ContentChatTable struct {
	gorm.Model
	Sender      string `json:"sender" gorm:"sender"`
	Receiver    string `json:"receiver" gorm:"receiver"`
	ContentType string `json:"contentType" gorm:"content_type"` // 内容类型，例如 "text", "image", "file"
	Content     string `json:"content" gorm:"content"`          // 存储内容的路径或 URL，具体内容依赖于 ContentType
	FileURL     string `json:"fileURL" gorm:"file_url"`         // 文件 URL（可选，适用于文件或图片）
}

func (receiver ContentChatTable) TableName() string {
	return "content_chat_table"
}
