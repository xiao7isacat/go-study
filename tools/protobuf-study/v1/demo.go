package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"main/v1/protoc"
)

func main() {
	var mobile = &protoc.MobileInfo{Brand: "wty"}
	data, err := proto.Marshal(mobile)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	var newMobile protoc.MobileInfo
	err = proto.Unmarshal(data, &newMobile)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	if mobile.GetBrand() != newMobile.GetBrand() {
		log.Fatalf("data mismatch %q != %q", mobile.GetBrand(), newMobile.GetBrand())
	}
	fmt.Printf("原始brand: %s \n转码后解码的brand: %s\n", mobile.GetBrand(), newMobile.GetBrand())
}
