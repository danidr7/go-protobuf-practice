package main

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"log"
	"testing"
)

func BenchmarkProtobufJson(b *testing.B) {
	p := Person{
		Name: "Jhon",
		Age:  30,
		Address: &Address{
			State: "SC",
			City:  "Florianópolis",
		},
	}

	b.Run("protobuf encoding", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			_, err := proto.Marshal(&p)
			if err != nil {
				log.Fatalf("fails attempting marshal: %s", err)
			}
		}

	})

	b.Run("protobuf decoding", func(b *testing.B) {
		dataEncoded, err := proto.Marshal(&p)
		if err != nil {
			log.Fatalf("fails attempting marshal: %s", err)
		}

		b.ReportAllocs()
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			personDecoded := Person{}
			err := proto.Unmarshal(dataEncoded, &personDecoded)
			if err != nil {
				log.Fatalf("fails attempting unmarshal: %s", err)
			}
		}
	})

	b.Run("json encoding", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			_, err := json.Marshal(&p)
			if err != nil {
				log.Fatalf("fails attempting marshal: %s", err)
			}
		}

	})

	b.Run("json decoding", func(b *testing.B) {
		dataEncoded, err := json.Marshal(&p)
		if err != nil {
			log.Fatalf("fails attempting marshal: %s", err)
		}

		b.ReportAllocs()
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			personDecoded := Person{}
			err := json.Unmarshal(dataEncoded, &personDecoded)
			if err != nil {
				log.Fatalf("fails attempting unmarshal: %s", err)
			}
		}
	})

}
