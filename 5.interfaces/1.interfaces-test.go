package main

import "fmt"

type IMessage interface {
	send() bool
}

type Email struct {
	email   string
	content string
}
type Wechat struct {
	wid     string
	content string
}

func (p Email) send() bool {
	fmt.Println("发送邮件提醒", p.email, p.content)
	return true
}

func (p Wechat) send() bool {
	fmt.Println("发送微信提醒", p.wid, p.content)
	return true
}
func something2(objectList []IMessage) {
	for _, item := range objectList {
		result := item.send()
		fmt.Println(result)
	}
}
func main() {
	messageObjectList := []IMessage{
		Email{"jojo@live.com", "注册成功"},
		Wechat{"jojo@live.com", "注册成功"},
	}
	something2(messageObjectList)
}
