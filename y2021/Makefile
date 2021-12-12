
target/release/y2021: src/*.rs
	cargo build --release

.PHONY: bench
bench: target/release/y2021
	du -h target/release/y2021
	hyperfine --warmup 100 --min-runs 300 ./target/release/y2021

