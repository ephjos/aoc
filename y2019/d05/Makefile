
.PHONY: p1 p2 clean

LIBS="IntCode.go"

p1:
	go run part1.go ${LIBS} < input | tee out1

p2:
	go run part2.go ${LIBS} < input | tee out2

clean:
	rm -rf out*
