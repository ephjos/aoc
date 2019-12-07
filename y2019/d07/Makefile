
.PHONY: p1 p2 clean

LIB="IntCode.go"

p1:
	go run part1.go ${LIB} < input | tee out1

p2:
	go run part2.go ${LIB} < input | tee out2

clean:
	rm -rf out*
