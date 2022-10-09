build: 
	GOOS=linux GOARCH=amd64 go build -o code/bootstrap code/main.go
	# upx -v code/bootstrap # 有 upx 工具可以取消注释

deploy:
	s deploy