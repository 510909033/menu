run: build
build:
	go mod tidy
	go build --ldflags "-extldflags -static" -o /gobuild/go_menu_run cmd/main/main.go
	nohup /gobuild/go_menu_run > /dev/null 2>&1 &

stop:
	 ps aux|grep go_menu_run|grep -v grep|awk '{print $2}'|xargs kill
	 echo "stop"
