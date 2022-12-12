package main

import (
	"fmt"

	pb "github.com/andres-root/proto-go/proto"
)

func do_simple() *pb.Simple {
	return &pb.Simple{
		Id:          43,
		Issimple:    true,
		Name:        "Michi",
		SampleLists: []int32{1, 2, 3, 4, 5, 6, 7},
	}
}

func main() {
	fmt.Println(do_simple())
}
