clean:
	rm -rf gen gen-*

gen-go:
	thrift --gen go *.thrift

test: gen-go
	go test . -v
