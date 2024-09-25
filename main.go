package main

import (
	"fmt"
	pb "go-proto/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main(){
	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id : 1,
			    Name : "Baju",
			    Price : 20000.00,
			    Stock : 200,
		    	Category: &pb.Category{
			    	Id : 1,
				    Name: "T-Shirt",
			    },
	    	},
			{
				Id : 1,
			    Name : "assic",
			    Price : 20000.00,
			    Stock : 200,
		    	Category: &pb.Category{
			    	Id : 2,
				    Name: "shoe",
			    },
	    	},
		},
	}

	data, err := proto.Marshal(products)
	if err != nil{
		log.Fatal("marshal error", err)
	}

	fmt.Println(data)

	testproducts := &pb.Products{}
	if err = proto.Unmarshal(data, testproducts); err != nil{
		log.Fatal("unmarshal error",err)
	}

	fmt.Println(testproducts)


}