package test

import (
	"encoding/json"
	"ohdada/g2gserver/internal/pkg/pb"

	"testing"

	"google.golang.org/protobuf/proto"
)

func BenchmarkProtobuf(b *testing.B) {
	b.Log("BenchmarkProtobuf start...")
	for i := 0; i < b.N; i++ {
		pbPlayer := &pb.Player{
			Sn:   1,
			Id:   "1234",
			Name: "1234",
		}

		data, err := proto.Marshal(pbPlayer)

		if err != nil {
			b.Fatal(err)
		}

		pbPlayer = &pb.Player{}

		proto.Unmarshal(data, pbPlayer)

		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJSON(b *testing.B) {
	b.Log("BenchmarkJSON start...")
	for i := 0; i < b.N; i++ {
		pbPlayer := &pb.Player{
			Sn:   1,
			Id:   "1234",
			Name: "1234",
		}

		data, err := json.Marshal(pbPlayer)

		if err != nil {
			b.Fatal(err)
		}

		pbPlayer = &pb.Player{}

		json.Unmarshal(data, pbPlayer)

		if err != nil {
			b.Fatal(err)
		}
	}
}
