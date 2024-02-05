# Calculator Kata

Write a calculator that reads an arbitrary complex mathematical term containing
floating point number, +, -, * and / as well as parenthesis. Calculate the result.

The calculator should be able to handle arbitrary long terms (i.e. megabytes of
input) well.

# Specification of the input language

The calculator processes input that is read as UTF-8 encoded characters. A 
formal specification of the input language is defined by the following [EBNF]:

```ebnf
(* An expression defines the start of the production. *)
expr = number | ( expr, S, operator, S, number ) | ( "(" S, expr, S, ")" );

(* The list of available operators. *)
operator = "+" | "-" | "*" | "/";

(* Definition of a number - either 0.xyz or abc.xyz *)
number = zero_fraction | non_zero_fraction;
non_zero_fraction = non_zero_digit { digit } { "." digit { digit } };
zero_fraction = "0" { "." digit { digit } };
non_zero_digit = "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9";
digit = "0" | non_zero_digit;

(* Any whitespace used as a separator including none. *)
S = "" | { " " | "\n" | "\t" | "\r" | "\f" | "\b" } ;
```

## Examples

The following lines are all valid input sequences:

```
0
0.123
807.1328
17 * 19
(21 - 3) * 8
```

[EBNF]: https://en.wikipedia.org/wiki/Extended_Backus–Naur_form