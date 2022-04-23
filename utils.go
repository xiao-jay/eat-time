package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type feishuMsg struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Post struct {
			ZhCn struct {
				Title   string `json:"title"`
				Content [][]struct {
					Tag    string `json:"tag"`
					Text   string `json:"text,omitempty"`
					Href   string `json:"href,omitempty"`
					UserId string `json:"user_id,omitempty"`
				} `json:"content"`
			} `json:"zh_cn"`
		} `json:"post"`
	} `json:"content"`
}

// SendMessage send Message to feishu group when request spend time a lot or other error
func SendMessage(msg string) {
	client := &http.Client{}
	feishuMsg := feishuMsg{MsgType: "post"}
	feishuMsg.Content.Post.ZhCn.Title = "吃饭 notice"
	feishuMsg.Content.Post.ZhCn.Content = make([][]struct {
		Tag    string `json:"tag"`
		Text   string `json:"text,omitempty"`
		Href   string `json:"href,omitempty"`
		UserId string `json:"user_id,omitempty"`
	}, 1)
	for i := 0; i < 1; i++ {
		feishuMsg.Content.Post.ZhCn.Content[i] = make([]struct {
			Tag    string "json:\"tag\""
			Text   string "json:\"text,omitempty\""
			Href   string "json:\"href,omitempty\""
			UserId string "json:\"user_id,omitempty\""
		}, 3)
	}
	feishuMsg.Content.Post.ZhCn.Content[0][0].Tag = "text"
	feishuMsg.Content.Post.ZhCn.Content[0][0].Text = msg + "\n"
	feishuMsg.Content.Post.ZhCn.Content[0][1].Tag = "text"
	feishuMsg.Content.Post.ZhCn.Content[0][1].Text = " Enviorment:"
	feishuMsg.Content.Post.ZhCn.Content[0][2].Tag = "at"
	feishuMsg.Content.Post.ZhCn.Content[0][2].UserId = "all"
	bytesData, _ := json.Marshal(feishuMsg)

	req, _ := http.NewRequest("POST", "", bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}
