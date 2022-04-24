package main

import (
	"fmt"
	pb "github.com/goharbor/personproto/protofiles"
	"google.golang.org/protobuf/proto"
)

func main(){
	p := &pb.Person{
		Id: 1234,
		Name: "Roger F",
		Email: "rf@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	body, _ := proto.Marshal(p)
	p1 := &pb.Person{}
	_ = proto.Unmarshal(body, p1)
	fmt.Println("Unmarshalled data ", body, "\n")
	fmt.Println("Unmashalled struct: ", p1)
}
