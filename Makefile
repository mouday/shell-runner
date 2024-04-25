# 版本号
version = ''

# 编译到 Linux
.PHONY: build-linux
build-linux:
	mkdir -p ./build/linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/linux/shell-runner ./src/main.go 

# 编译到 macOS
# make build-darwin
.PHONY: build-darwin
build-darwin:
	mkdir -p ./build/darwin
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./build/darwin/shell-runner ./src/main.go

# 编译到 windows
.PHONY: build-windows
build-windows:
	mkdir -p ./build/windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/windows/shell-runner.exe ./src/main.go 

# 编译到 全部平台
# make build
.PHONY: build
build:
	make clean
	make build-linux
	make build-darwin
	make build-windows


# 发布linux
.PHONY: release-linux
release-linux:
	mkdir -p ./release/linux/shell-runner-$(version)-linux-amd64/
	cp ./env.example ./release/linux/shell-runner-$(version)-linux-amd64/
	cp -r ./scripts ./release/linux/shell-runner-$(version)-linux-amd64/
	cp -r ./config ./release/linux/shell-runner-$(version)-linux-amd64/
	cp ./build/linux/shell-runner ./release/linux/shell-runner-$(version)-linux-amd64/
	tar -zcvf release/shell-runner-$(version)-linux-amd64.tar.gz -C ./release/linux/ ./shell-runner-$(version)-linux-amd64

# 发布darwin
.PHONY: release-darwin
release-darwin:
	mkdir -p ./release/darwin/shell-runner-$(version)-darwin-amd64/
	cp ./env.example ./release/darwin/shell-runner-$(version)-darwin-amd64/
	cp -r ./scripts ./release/darwin/shell-runner-$(version)-darwin-amd64/
	cp -r ./config ./release/darwin/shell-runner-$(version)-darwin-amd64/
	cp ./build/darwin/shell-runner ./release/darwin/shell-runner-$(version)-darwin-amd64/
	tar -zcvf release/darwin/shell-runner-$(version)-darwin-amd64.tar.gz -C ./release/darwin/ ./shell-runner-$(version)-darwin-amd64

# 发布windows
.PHONY: release-windows
release-windows:
	mkdir -p ./release/windows/shell-runner-$(version)-windows-amd64/
	cp ./env.example ./release/windows/shell-runner-$(version)-windows-amd64/
	cp -r ./scripts ./release/windows/shell-runner-$(version)-windows-amd64/
	cp -r ./config ./release/windows/shell-runner-$(version)-windows-amd64/
	cp ./build/windows/shell-runner.exe ./release/windows/shell-runner-$(version)-windows-amd64/
	cd ./release/windows && zip -r ../shell-runner-$(version)-windows-amd64.zip shell-runner-$(version)-windows-amd64/; cd -

# 发布全平台
# make release
.PHONY: release
release:
	make release-linux
	make release-darwin
	make release-windows


.PHONY: clean
clean:
	rm -rf ./build ./release

.PHONY: dev
dev:
	# go run ./src/main.go
	air