VERSION?=dev
GOOS?=linux
BUILD_PATH?=./target
NAME=mich-echo-bot
build:
	rm -f $(BUILD_PATH)/$(NAME)
	go build $(GO_TAGS) $(GO_LDFLAGS) -o $(BUILD_PATH)/$(NAME) ./
	cp -v config.yml $(BUILD_PATH)/

clean:
	rm -rf $(BUILD_PATH)
