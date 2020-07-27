# Problem #1

## Multiples of 3 and 5

```If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.```

### Thought process for solution

We can itterate through all the numbers starting at 1 and going to 1000; If the number is modular to
3 or 5 we can add it to sum.

### Ways to make this really cool!

We can use number theory to check if numbers are divisible without going through the modular operation.

*Number Theory for 3*: If the sum of all digits in a number is divisible by 3 then the number is divisible by 3.

*Number Thoery for 5*: For all intergers greater than 0, if the last digit is 5 or 0 then it is divisible by 5.
