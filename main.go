package main

import (
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	p := Person{
		Name: "Jhon",
		Age:  30,
		Address: &Address{
			State: "SC",
			City:  "Florianópolis",
		},
	}

	dataEncoded, err := proto.Marshal(&p)
	if err != nil {
		log.Fatalf("fails attempting marshal: %s", err)
	}

	log.Print(dataEncoded)

	personDecoded := Person{}
	err = proto.Unmarshal(dataEncoded, &personDecoded)
	if err != nil {
		log.Fatalf("fails attempting unmarshal: %s", err)
	}

	log.Printf("decoded: %s - %d - %+v", personDecoded.GetName(), personDecoded.GetAge(), personDecoded.GetAddress())
}
