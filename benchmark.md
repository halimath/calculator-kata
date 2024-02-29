# Benchmark Results

# Language Comparison

These benchmark results compare the performance of a Go(lang) and Rust implementation
using the same implementation strategy (reading the input in chunks, scanning for
tokens, transforming the token stream to reverse-polish-notation, evaluating the
NPM stream using a stack based evaluator).

The results are obtained from running the resulting binaries (produced with
optimizations applied) and using the standard unix `time` tool to generate
insights.

Lang | Real time [s] | User time [s] | Max. resident set size [bytes] | Peak memory footprint [bytes] | LoC
-- | --: | --: | --: | --: | --:
Go | 0.22 | 0.22 | 9,502,720 | 8,323,968 | 566
Rust | 0.15 | 0.15 | 1,409,024 | 967,168 | 393
Python | 3.82 | 3.78 | 13,352,960 | 8,963,328 | 164

# Go Implementation Benchmarks

The folllowing table compares different implementation strategies all using the
Go(lang) programming language.

All benchmarks have been run on a 2021 MacBook using a M1 pro Apple Silicon
processor with 16Gb ram using go 1.22.0.

## Results

Strategy | Test | Avg. duration | Avg. bytes/op | Avg. allocs/op
-- | -- | --: | --: | --:
RPN w/ Token struct (string value) | Simple | 1,938 ns/op | 7,976 | 22
| | 1k |  24,578 ns/op | 12,464 | 394
| | 10k | 258,822 ns/op | 71,600 | 4,591
| | 1m | 25,602,550 ns/op | 6,731,522 | 481,991
| | 10m | 249,842,958 ns/op | 67,243,596 | 4,816,630
| | | | 
RPN w/ Token struct (float64 value) | Simple | 1,663 ns/op | 6,888 | 14
| | 1k |  18,877 ns/op | 9,336 | 223
| | 10k | 189,297 ns/op | 43,216 | 2,737
| | 1m | 19,423,377 ns/op | 38,008,740 | 283,417
| | 10m | 193,736,326 ns/op | 38,042,613 | 2,834,386
| | | | 
RPN w/ Token pointer + obj pool | Simple | 2,039 ns/op | 6,128 | 36
| | 1k |  30,775 ns/op | 13,064 | 702
| | 10k | 316,661 ns/op | 86,424 | 7,820
| | 1m | 32,217,288 ns/op | 8,473,102 | 822,405
| | 10m | 318,119,521 ns/op | 84,580,920 | 8,216,380
| | | |
RPN w/ Token interface | Simple | 2,027 ns/op | 7,256 | 21
| | 1k |  24,532 ns/op | 11,848 | 364
| | 10k | 250,494 ns/op | 65,840 | 4,135
| | 1m | 25,718,884 ns/op | 6,076,852 | 425,155
| | 10m | 255,909,386 ns/op | 60,718,428 | 4,251,608
| | | |
AST w/ Token interface | Simple | 1,757 ns/op | 4,840 | 31
| | 1k |  28,371 ns/op | 16,632 | 566
| | 10k | 283,676 ns/op | 127,249 | 5,594
| | 1m | 37,091,376 ns/op | 12,477,205 | 566,954
| | 10m | 353,062,292 ns/op | 124,719,765 | 5,668,890

## Strategies

### RPN w/ Token struct (string value)

* Input stream is scanned for _Tokens_
* `Token` is a struct with a type discriminator (`int`) and an optional `string`
  value (for numbers)
* Token stream is reordered to _reverse-polish notation_ (RPN) using an
  implementation of the [shunting yard algortithm].
* RPN token stream is processed by a stack-based evaluator
* Result is pop'ed of the stack

### RPN w/ Token struct (float64 value)

* Input stream is scanned for _Tokens_
* `Token` is a struct with a type discriminator (`int`) and an optional `float64`
  value (for numbers)
* Token stream is reordered to _reverse-polish notation_ (RPN) using an
  implementation of the [shunting yard algortithm].
* RPN token stream is processed by a stack-based evaluator
* Result is pop'ed of the stack

### RPN w/ Token pointer + obj pool

* Input stream is scanned for _Tokens_
* `Token` is a struct with a type discriminator (`int`) and an optional `string`
  value (for numbers)
* Tokens are passed as references (pointers) with a global object pool used
  for allocating and returning tokens so that tokens can be reused
* Token stream is reordered to _reverse-polish notation_ (RPN) using an
  implementation of the [shunting yard algortithm].
* RPN token stream is processed by a stack-based evaluator
* Result is pop'ed of the stack

### RPN w/ Token interface

* Input stream is scanned for _Tokens_
* `Token` is an interface
* Each token type is a separate type satisfying the `Token` interface
* Token stream is reordered to _reverse-polish notation_ (RPN) using an
  implementation of the [shunting yard algortithm].
* RPN token stream is processed by a stack-based evaluator
* Result is pop'ed of the stack

### AST w/ Token interface

* Input stream is scanned for _Tokens_
* `Token` is an interface
* Each token type is a separate type satisfying the `Token` interface
* Token stream is parsed into an [abstract syntax tree] using an [LL(1) parser]
* The AST is evaluated using a depth first, left to right traversal

[shunting yard algortithm]: https://en.wikipedia.org/wiki/Shunting_yard_algorithm
[LL(1) parser]: https://en.wikipedia.org/wiki/LL_parser
[abstract syntax tree]: https://en.wikipedia.org/wiki/Abstract_syntax_tree