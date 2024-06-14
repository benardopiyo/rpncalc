# RPN Calculator

* The rpncalc program is designed to evaluate equations written in Reverse Polish Notation (RPN). 
* RPN is a mathematical notation where every operator follows all of its operands. This format eliminates the need for parentheses to denote precedence in expressions, as the order of operations is inherent in the notation.


## Features

* Evaluates expressions written in Reverse Polish Notation.
* Supports basic arithmetic operations: addition (+), subtraction (-), multiplication (*), division (/), and modulo (%).
* Handles invalid inputs gracefully by printing "Error".
* Accommodates extra spaces within the input, considering them as valid.

## Usage

* To use the rpncalc program, run it with a string containing the RPN expression as its argument. Here's how to execute the program:

```bash
$ go run . "1 2 * 3 * 4 +"
```

This command will evaluate the RPN expression 1 2 * 3 * 4 + and print the result:

```bash
10$
```

## Error Handling

* If the number of arguments is not exactly one, the program will output Error.
* If the input string contains invalid tokens that are not recognized as numbers or operators, the program will also output Error.
* The program is designed to handle input with extra spaces, which should still be considered valid.

### Examples

* Valid RPN Expression:
```bash
$ go run . "1 2 * 3 * 4 +"
10$
```

* Invalid Expression (not enough operands):
```bash
$ go run . "1 2 3 4 +"
Error$
```

* No Arguments:
```bash
$ go run .
Error$
```

* Expression with Extra Spaces:
```bash
$ go run . "  1   3 * 2 -"
1$
```

* Invalid Token:
```bash
$ go run . "  1   3 * ksd 2 -"
Error$
```