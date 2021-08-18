package test

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"testing"
	"unioj/utils/future/protobuf"
)

func TestProtoBuffer(t *testing.T) {
	text := &protobuf.PandaRequest{
		Name:   "panda",
		Weight: []int32{100, 101, 103, 104, 120},
		Height: 120,
		Motto:  "always head up high",
	}
	fmt.Println(text)
	data, err := proto.Marshal(text)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	newPanda := &protobuf.PandaRequest{}
	err = proto.Unmarshal(data, newPanda)
	if err != nil {
		return
	}
	fmt.Println(newPanda)
}
