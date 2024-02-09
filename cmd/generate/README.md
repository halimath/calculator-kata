# generator

This directory contains a small tool used to generate valid mathematical
expressions with respect to the formal specs given in the repo's `README`.

# Usage

```
Usage of generator:
  -max-decimals int
        Max number of decimal places (default 2)
  -max-digits int
        Max number of digits per number (default 3)
  -min-size int
        Min file size in bytes (default 2048)
```

Note that `-min-size` is used to provide a lower bound for the resulting size.
The generator will generate _at least_ that much bytes will always generate valid
expressions, thus the resulting size will be slightly larger.

Also note that in order to avoid _division by zero_ errors the generator will
_never_ include '/' as an operator.

# Build

Note that the generator uses go 1.22 and the new random number generator package.

```shell
go build main.go
```