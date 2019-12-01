
.PHONY: build test clean

EXEC="script"

build:
	go build -o ${EXEC} .

test:
	make build
	./${EXEC} | tee output

clean:
	rm -rf ${EXEC} output

