package main

import (
	"fmt"
	"testProto/user"

	"google.golang.org/protobuf/proto"
)

func main() {
	userbyte := &user.Article{
		Aid:   1,
		Title: "hello",
		Views: 100,
	}

	bytes, err := proto.Marshal(userbyte)
	if err != nil {
		panic(err)
	}

	fmt.Printf("bytes: %v\n", bytes)

	userobg := &user.Article{}
	err = proto.Unmarshal(bytes, userobg)

	if err != nil {
		panic(err)
	}

	fmt.Printf("userbog: %v\n", userobg)
}
