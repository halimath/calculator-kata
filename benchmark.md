# Benchmark Results

All benchmarks have been run on a 2021 MacBook using a M1 pro Apple Silicon
processor with 16Gb ram using go 1.22.0.

# Implementation Strategies

The results compare different implementation strategies:

# Results

Strategy | Test | Avg. duration | Avg. bytes/op | Avg. allocs/op
-- | -- | --: | --: | --:
RPN w/ Token struct | Simple | 1,938 ns/op | 7976 | 22
RPN w/ Token struct | 1k |  24,578 ns/op | 12464 | 394
RPN w/ Token struct | 10k | 258,822 ns/op | 71600 | 4,591
RPN w/ Token struct | 1m | 25,602,550 ns/op | 6731522 | 481,991
RPN w/ Token struct | 10m | 249,842,958 ns/op | 67243596 | 4,816,630
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