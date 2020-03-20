run: build
	-docker container stop dev_menu 
	-docker container rm -f dev_menu 
	cd docker
	nohup docker-compose -f docker/docker-compose-dev.yml up > /dev/null 2>&1 &
build:
	go mod tidy
	go build --ldflags "-extldflags -static" -o menu_dev cmd/main/main.go
	nohup ./menu_dev -config=config_dev.ini > /tmp/menu_dev 2>&1 &
stop:
	 ps aux|grep "config_dev"|grep -v grep|awk '{print $$2}'|xargs kill
	 echo "stop"



