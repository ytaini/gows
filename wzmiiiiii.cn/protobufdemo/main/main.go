package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"wzmiiiiii.cn/protobufdemo/pb"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
			{Number: "123-1231", Type: pb.Person_WORK},
			{Number: "312-1233", Type: pb.Person_MOBILE},
		},
		LastUpdated: timestamppb.Now(),
	}
	data, err := proto.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	p1 := &pb.Person{}
	err = proto.Unmarshal(data, p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println(p1)
}
