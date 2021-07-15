OUT_FILE=cdn.out

build:
	go build -o ${OUT_FILE}

run:
	go run .

deps:
	go get -v

clean:
	go clean
	rm cdn* dist
