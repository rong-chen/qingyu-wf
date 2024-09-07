package Websocket

// MessageParams ->    Type: 	  <text>普通文本消息,<file>文件,<notice>通知,<noticeService>通知服务器,<voice>语音,<video>视频
// MessageParams -> Description: Type为<voice>语音,<video>视频，这两个使用字段
type MessageParams struct {
	Type        string `json:"type"`
	Message     string `json:"message"`
	Description string `json:"description"`
	Sender      string `json:"sender"`
	Receiver    string `json:"receiver"`
}
