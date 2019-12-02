
.PHONY: build b1 b2 test t1 t2 clean

build:
	make b1
	make b2

b1:
	go build -o part1 part1.go

b2:
	go build -o part2 part2.go

test:
	make t1
	make t2

t1:
	make b1
	./part1 < input | tee out1

t2:
	make b2
	./part2 < input | tee out2

clean:
	rm -rf part1 part2 out*
