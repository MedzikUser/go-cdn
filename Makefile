OUT_FILE=cdn.out

build:
	go build -o ${OUT_FILE}

all:
	./build.sh

run:
	go run .

deps:
	go get -v

clean:
	go clean
	rm cdn* MD5* SHA256* VERSION
