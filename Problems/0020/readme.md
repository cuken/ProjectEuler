# Problem 20

[Link to Problem](https://projecteuler.net/problem=20)

## Factorial digit sum

```
<i>n</i>! means <i>n</i> × (<i>n</i> − 1) × ... × 3 × 2 × 1
For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,<br/>and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
Find the sum of the digits in the number 100!

```

### Thought Process

We can probably side step the issue of working with large numbers by using the big int option in math.big for golang. This is going to be similar to problem 16.
