package main

// Message 结构体
type Message struct {
	MsgType   string         `json:"msg_type"`
	Sign      string         `json:"sign"`
	Content   MessageContent `json:"content"`
	Timestamp string         `json:"timestamp"`
}

// MessageContent 是一个接口，用于区分不同类型的消息内容
type MessageContent interface{}

// TextMessage 结构体用于文本消息
type TextMessage struct {
	Text string `json:"text"`
}

// RichTextMessage 结构体用于富文本消息
type RichTextMessage struct {
	Post struct {
		Cn RichSubBody `json:"zh_cn,omitempty"`
		//Us RichSubBody `json:"en_us,omitempty"`
	} `json:"post"`
}

// RichSubBody 结构体用于富文本的具体内容
type RichSubBody struct {
	Title   string      `json:"title"`
	Content [][]Segment `json:"content"`
}

type Segment struct {
	Tag     string `json:"tag"`
	Text    string `json:"text,omitempty"`
	Escape  bool   `json:"un_escape,omitempty"`
	Url     string `json:"href,omitempty"`
	UserId  string `json:"user_id,omitempty"`
	ImagUrl string `json:"image_key,omitempty"`
}
