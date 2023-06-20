setup-air:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

dev:
	clear && ./bin/air

restartBuild:
	rm -rf build && mkdir build
	
buildGo:
	GOOS=linux GOARCH=amd64  go build -a -ldflags="-s -w" -o build/main main.go 
	
compress:
	zip -j build/main.zip build/main

migrate:
	go run github.com/prisma/prisma-client-go migrate dev --schema=./src/prisma/schema.prisma

migrate-deploy:
	go run github.com/prisma/prisma-client-go migrate deploy --schema=./src/prisma/schema.prisma

pcg:
	go run github.com/prisma/prisma-client-go generate --schema=./src/prisma/schema.prisma

glg:
	go run github.com/arsmn/fastgql generate

mainBuild:
	go build -o main main.go

	
getfast:
	go get github.com/arsmn/fastgql 