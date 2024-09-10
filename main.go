/*
此代码由Bing Ai 生成
*/
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var config = ""

func genRichMessage(message *Message, title string, contentStr string, url string) {
	message.MsgType = "post" // 改写为post 飞书这里就像妥屎一样烂.
	segmentTitle := Segment{
		Tag:  "text",
		Text: title,
	}

	segmentHref := Segment{
		Tag:  "a",
		Text: "链接",
		Url:  url,
	}

	segmentList := []Segment{segmentTitle, segmentHref}
	segArray := [][]Segment{segmentList}

	richBody := RichSubBody{
		Title:   contentStr,
		Content: segArray,
	}

	newBody := struct {
		Cn RichSubBody `json:"zh_cn,omitempty"`
	}{
		Cn: richBody,
	}

	message.Content = RichTextMessage{
		Post: newBody,
	}
}

func main() {
	msgType := flag.String("type", "text", "消息类型: text, rich")
	configd := flag.String("conf", "config.yaml", "配置文件路径")
	content := flag.String("content", "", "消息内容")
	title := flag.String("title", "", "消息标题")
	messageURL := flag.String("url", "", "消息链接")
	//picURL := flag.String("pic", "", "图片链接")
	flag.Parse()
	config = *configd

	if *content == "" && !isInputFromPipe() {
		fmt.Println("请提供消息内容")
		flag.PrintDefaults()
		return
	}

	if *content == "" && isInputFromPipe() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			*content += scanner.Text()
		}
	}

	message := Message{MsgType: *msgType}
	switch *msgType {
	case "text":
		message.Content = TextMessage{Text: *content}
	case "rich":
		genRichMessage(&message, *title, *content, *messageURL)
	default:
		fmt.Println("无效的消息类型")
		return
	}

	sendMessage(message)
}

func sendMessage(message Message) {
	config := getConfig()
	timestamp := time.Now().Unix()
	signature, err := getSignature(config.Secret, timestamp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	message.Sign = signature
	message.Timestamp = strconv.Itoa(int(timestamp))

	data, err := json.MarshalIndent(message, "", "  ")
	fmt.Println(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", config.Webhook, bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
