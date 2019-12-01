
.PHONY: build test clean

EXEC="01"

build:
	go build -o ${EXEC} .

test:
	make build
	./${EXEC} | tee output

clean:
	rm -rf ${EXEC} output

