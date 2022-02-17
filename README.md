# Numeral Systems

> Warning: It is not allowed to use any external long maths / big numbers libraries.

## Task 1: Long Multiplication Base 10 (100%)

### Definition

Write a program (with tests) that implements "long" arithmetics multiplication. The program may contain
a `LongMul` function, which accepts two decimal numbers encoded as strings.

```go
func LongMul(a, b string) string { ... }
```

### Test examples

```go
{"55555", "55", "3055525"},
{"98765432109876543210987654321098765432109876543210",
  "22222222222222222222222222222222222222222222222222",
  "2194787380219478738021947873802194787380219478737978052126197805212619780521261978052126197805212620"},
```

## Extra Task: Long Addition Base N (+30%)

Write a program (with tests) that implements "long" arithmetics multiplication in arbitrary numeral system.
The program may contain a `LongMul` function, which accepts two numbers encoded as strings and the base.

```go
func LongMul(a, b string, base int) string { ... }
```

### Test Examples

```go
{"55555", "55", 10, "3055525"},
{"98765432109876543210987654321098765432109876543210",
  "22222222222222222222222222222222222222222222222222",
  16,
  "14540b39e014540b39e014540b39e014540b39e014540b39dfebabf4c61febabf4c61febabf4c61febabf4c61febabf4c620"},
