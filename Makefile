run: build
	cp -rf ./ /var/www/dev/menu/
	-docker container stop dev_menu 
	-docker container rm -f dev_menu 
	cd docker
	nohup docker-compose -f docker/docker-compose-dev.yml up > /dev/null 2>&1 &
build:
	go mod tidy
	rm -rf /var/www/dev/menu
	mkdir -p /var/www/dev/menu
	mkdir -p /var/log/dev/menu
	go build --ldflags "-extldflags -static" -o /var/www/dev/menu/go_menu_run cmd/main/main.go

stop:
	 ps aux|grep "config_dev"|grep -v grep|awk '{print $$2}'|xargs kill
	 echo "stop"



