package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"tohrux.com/grpc/service"
)

func main() {
	user := &service.User{
		Username: "jojo", Age: 20,
	}

	//序列化
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	//反序列化
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}

	fmt.Println(newUser.String())

}
