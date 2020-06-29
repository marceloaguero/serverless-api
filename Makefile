#.PHONY: build clean deploy gomodgen run-local
.PHONY: run-local

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/users users/deliveries/lambda/main.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

run-local-mongo:
	docker run --name serverless-api-mongo -d -p 27017:27017 --mount type=volume,src=serverless-api-mongo,dst=/data/db mongo:4.2
	PORT=8005 DB_URI=mongodb://127.0.0.1:27017 DB_NAME=serverless TABLE_NAME=users go run cmd/server/main.go

run-local-mysql:
	docker run --name serverless-api-mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=supersecret --mount type=volume,src=serverless-api-mysql,dst=/var/lib/mysql mysql:8.0
	PORT=8005 DB_URI=mysql://root:supersecret@127.0.0.1:3306 DB_NAME=serverless TABLE_NAME=users go run cmd/server/main.go
	