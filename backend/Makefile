GO = go

TARGET = mm-wiki

all:$(TARGET)

$(TARGET): *.go
	$(GO) build -o $@ $^

test:
	$(GO) test $(M)  -v -gcflags=all=-l -coverpkg=./... -coverprofile=test.out ./...

clean:
	rm -f $(TARGET)

linux:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(TARGET)

windows:
	GOOS=windows GOARCH=amd64 $(GO) build -o $(TARGET)

darwin:
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(TARGET)