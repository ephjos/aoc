package main

import (
	"crypto/md5"
	"fmt"
	"runtime"
	"slices"
	"strings"
	//"strconv"
)

type day05 struct{}

type a_HashByte struct {
	suffix int
	b      byte
}

// Directly convert relevant word to equivalent hex digit ascii value
var word_to_hex_char = [16]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func a_worker(results chan []a_HashByte, prefix string, start int, end int) {
	var hash_bytes []a_HashByte

	input := []byte(fmt.Sprintf("%s%d", prefix, start-1))
	input_len := len(input)
	prefix_len := len(prefix)
	for i := start; i < end; i++ {
		// Increment the input bytes directly, avoid fmt.Sprintf allocation on every loop
		if input[input_len-1] == '9' {
			j := input_len - 1
			for input[j] == '9' {
				input[j] = '0'
				j -= 1
			}

			if j < prefix_len {
				input = append(input, '0')
				input_len = len(input)
				input[j+1] = '1'
			} else {
				input[j] += 1
			}

		} else {
			input[input_len-1] += 1
		}

		hash := md5.Sum(input)
		if hash[0] == 0 && hash[1] == 0 && hash[2] <= 0x0F {
			hash_bytes = append(hash_bytes, a_HashByte{i, word_to_hex_char[hash[2]]})
		}
	}

	results <- hash_bytes
}

func day05_a_core(input string, block_size int) string {
	root := strings.Trim(input, " \n")
	num_workers := runtime.NumCPU()
	window_size := block_size * num_workers
	window := 0

	var hash_bytes []a_HashByte

	for len(hash_bytes) < 8 {
		results := make(chan []a_HashByte, num_workers)
		window_start := window * window_size

		for i := 0; i < num_workers; i++ {
			go a_worker(results, root, window_start+(block_size*i), window_start+(block_size*(i+1)))
		}

		for i := 0; i < num_workers; i++ {
			result := <-results
			hash_bytes = append(hash_bytes, result...)
		}

		window += 1
	}

	slices.SortFunc(hash_bytes, func(a, b a_HashByte) int {
		return a.suffix - b.suffix
	})

	out := [8]byte{}
	for i := 0; i < 8; i++ {
		out[i] = hash_bytes[i].b
	}

	return string(out[:])
}

func (_d day05) a(input string) string {
	return day05_a_core(input, 1<<14)
}

type b_HashByte struct {
	suffix int
	pos    int
	b      byte
}

func b_worker(results chan []b_HashByte, prefix string, start int, end int) {
	var hash_bytes []b_HashByte

	input := []byte(fmt.Sprintf("%s%d", prefix, start-1))
	input_len := len(input)
	prefix_len := len(prefix)
	for i := start; i < end; i++ {

		// Increment the input bytes directly, avoid fmt.Sprintf allocation on every loop
		if input[input_len-1] == '9' {
			j := input_len - 1
			for input[j] == '9' {
				input[j] = '0'
				j -= 1
			}

			if j < prefix_len {
				input = append(input, '0')
				input_len = len(input)
				input[j+1] = '1'
			} else {
				input[j] += 1
			}

		} else {
			input[input_len-1] += 1
		}

		hash := md5.Sum(input)
		if hash[0] == 0 && hash[1] == 0 && hash[2] <= 0x0F {
			j := hash[2]
			if j >= 0 && j <= 7 {
				hash_bytes = append(hash_bytes, b_HashByte{i, int(j), word_to_hex_char[hash[3]>>4]})
			}
		}
	}

	results <- hash_bytes
}

func day05_b_core(input string, block_size int) string {
	root := strings.Trim(input, " \n")
	num_workers := runtime.NumCPU()
	window_size := block_size * num_workers
	window := 0

	count := 0
	out := [8]byte{}
	set := [8]bool{}

	for count < 8 {
		var hash_bytes []b_HashByte

		results := make(chan []b_HashByte, num_workers)
		window_start := window * window_size

		for i := 0; i < num_workers; i++ {
			go b_worker(results, root, window_start+(block_size*i), window_start+(block_size*(i+1)))
		}

		for i := 0; i < num_workers; i++ {
			result := <-results
			hash_bytes = append(hash_bytes, result...)
		}

		window += 1

		slices.SortFunc(hash_bytes, func(a, b b_HashByte) int {
			return a.suffix - b.suffix
		})

		for i := 0; i < len(hash_bytes); i++ {
			hb := hash_bytes[i]
			if !set[hb.pos] {
				out[hb.pos] = hb.b
				set[hb.pos] = true
				count += 1
				if count == 8 {
					break
				}
			}
		}
	}

	return string(out[:])
}

func (_d day05) b(input string) string {
	return day05_b_core(input, 1<<16)
}
