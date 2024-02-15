# Benchmark Results

# Language Comparison

These benchmark results compare the performance of a Go(lang) and Rust implementation
using the same implementation strategy (reading the input in chunks, scanning for
tokens, transforming the token stream to reverse-polish-notation, evaluating the
NPM stream using a stack based evaluator).

The results are obtained from running the resulting binaries (produced with
optimizations applied) and using the standard unix `time` tool to generate
insights.

Lang | Real time [s] | User time [s] | Max. resident set size [bytes] | Peak memory footprint [bytes]
-- | --: | --: | --: | --:
Go | 0.22 | 0.22 | 9,502,720 | 8,323,968
Rust | 0.15 | 0.15 | 1,409,024 | 967,168

# Go Implementation Benchmarks

All benchmarks have been run on a 2021 MacBook using a M1 pro Apple Silicon
processor with 16Gb ram using go 1.22.0.

## Implementation Strategies

The results compare different implementation strategies. The strategies are
explained below.

## Results

Strategy | Test | Avg. duration | Avg. bytes/op | Avg. allocs/op
-- | -- | --: | --: | --:
RPN w/ Token struct (string value) | Simple | 1,938 ns/op | 7,976 | 22
RPN w/ Token struct (string value) | 1k |  24,578 ns/op | 12,464 | 394
RPN w/ Token struct (string value) | 10k | 258,822 ns/op | 71,600 | 4,591
RPN w/ Token struct (string value) | 1m | 25,602,550 ns/op | 6,731,522 | 481,991
RPN w/ Token struct (string value) | 10m | 249,842,958 ns/op | 67,243,596 | 4,816,630
| | | | 
RPN w/ Token struct (float64 value) | Simple | 1,663 ns/op | 6,888 | 14
RPN w/ Token struct (float64 value) | 1k |  18,877 ns/op | 9,336 | 223
RPN w/ Token struct (float64 value) | 10k | 189,297 ns/op | 43,216 | 2,737
RPN w/ Token struct (float64 value) | 1m | 19,423,377 ns/op | 38,008,740 | 283,417
RPN w/ Token struct (float64 value) | 10m | 193,736,326 ns/op | 38,042,613 | 2,834,386
| | | | 
RPN w/ Token pointer + obj pool | Simple | 2,039 ns/op | 6,128 | 36
RPN w/ Token pointer + obj pool | 1k |  30,775 ns/op | 13,064 | 702
RPN w/ Token pointer + obj pool | 10k | 316,661 ns/op | 86,424 | 7,820
RPN w/ Token pointer + obj pool | 1m | 32,217,288 ns/op | 8,473,102 | 822,405
RPN w/ Token pointer + obj pool | 10m | 318,119,521 ns/op | 84,580,920 | 8,216,380
| | | |
RPN w/ Token interface | Simple | 2,027 ns/op | 7,256 | 21
RPN w/ Token interface | 1k |  24,532 ns/op | 11,848 | 364
RPN w/ Token interface | 10k | 250,494 ns/op | 65,840 | 4,135
RPN w/ Token interface | 1m | 25,718,884 ns/op | 6,076,852 | 425,155
RPN w/ Token interface | 10m | 255,909,386 ns/op | 60,718,428 | 4,251,608
| | | |
AST w/ Token interface | Simple | 1,757 ns/op | 4,840 | 31
AST w/ Token interface | 1k |  28,371 ns/op | 16,632 | 566
AST w/ Token interface | 10k | 283,676 ns/op | 127,249 | 5,594
AST w/ Token interface | 1m | 37,091,376 ns/op | 12,477,205 | 566,954
AST w/ Token interface | 10m | 353,062,292 ns/op | 124,719,765 | 5,668,890