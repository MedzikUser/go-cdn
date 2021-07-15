OUT_FILE=cdn.out

build:
	go build -o ${OUT_FILE}

run:
	go run .

clean:
	go clean
	rm ${OUT_FILE} dist
