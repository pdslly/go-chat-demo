package main

import "encoding/json"

type BroatMsg struct {
	client *Client
	chat *ChatMsg
}

type ChatMsg struct {
	Type string `json:"type"`
	Date string `json:"date"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Message string `json:"message"`
}

type UserMsg struct {
	Type string `json:"type"`
	Users []User `json:"users"`
}

type ResMsg struct {
	Type string `json:"type"`
	Date string `json:"date"`
	Name string `json:"name"`
	Message string `json:"message"`
	IsSys bool `json:"isSys"`
	IsSelf bool `json:"isSelf"`
}

func resMsgFormat(t, date, name, message string, isSys, isSelf bool) []byte {
	msg := ResMsg{t, date, name, message, isSys, isSelf}
	ret, err := json.Marshal(msg)
	checkErr(err)
	return ret
}
