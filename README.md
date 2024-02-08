# Calculator Kata

Write a calculator that reads mathematical terms containing
floating point numbers, +, -, * and / as well as parenthesis. 

Calculate the result by applying the "usual" mathematical rules:

* operators are used in _infix notation_: `2 + 3`
* multiply/div have higher precedence then add/sub
* parenthesis have higher precedence then operators
* numbers can contain a mix of integers and floating point numbers
* calculation should follow mathematical rules, i.e. `2 / 0` is an error and not
  something like `NaN` or `-Inf`

The calculator should

* calculate the result of a single mathematical expression (no matter how long)
* print the result in a human readable style
* report any error and refuse the calculation (error reporting style doesn't matter)
* handle arbitrary long terms (i.e. tens of megabytes of input) well (i.e. in 
  terms of memory usage as well as running time)

# Specification of the input language

The calculator processes input that is read as `UTF-8` encoded characters. A 
formal specification of the input language is defined by the following [EBNF]:

```ebnf
(* An expression defines the start of the production. *)
expr = number (* A bare number is a valid expression*)
     | ( expr, S, operator, S, number ) (* left-recursive operator chained expression *)
     | ( "(" S, expr, S, ")" ); (* parenthesis surround an expression *)

(* The list of available operators. *)
operator = "+" | "-" | "*" | "/";

(* Definition of a number - either 0.xyz or abc.xyz *)
number = zero_fraction 
       | non_zero_fraction;

non_zero_fraction = non_zero_digit { digit } { "." digit { digit } };

zero_fraction = "0" { "." digit { digit } };

non_zero_digit = "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9";

digit = "0" | non_zero_digit;

(* Any whitespace used as a separator including none. *)
S = "" | { " " | "\n" | "\t" | "\r" | "\f" | "\b" } ;
```

Note that expressions in parenthesis can also contain parenthesis.

[EBNF]: https://en.wikipedia.org/wiki/Extended_Backus–Naur_form

## Examples

The following lines are all valid input sequences:

```
0
0.123
807.1328
17 * 19
(21 - 3) * 8
```

# Implementation Restrictions

* Only use the standard library for the production code; do not rely on external
  libraries (except for testing/benchmarking)
* Apply an appropriate structure (i.e. packages, types, ...)
* Write _idiomatic_ code
* Write unit tests to


# Testcases

A couple of test files with increasing sizes are provided in [`testdata`](./testdata). 
The file names roughly tell the file's sizes. You can use them for functional
as well as for benchmark tests.

The results are

* `testdata/1k` -> 248253190541.03891
* `testdata/10k` -> 214512826488689376
* `testdata/100k` -> 481585283158357967896576
* `testdata/1m` -> 830415166156152287340265472
* `testdata/10m` -> 6780466519056739908798922614519103488
