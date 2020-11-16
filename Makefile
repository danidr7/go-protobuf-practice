compile:
	protoc --go_opt=paths=source_relative --go_out=. person.proto

run:
	go run main.go person.pb.go

benchmark:
	go test -test.run=XXX -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out