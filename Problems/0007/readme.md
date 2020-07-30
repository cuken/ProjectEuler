# Problem 7

[Link to Problem](https://projecteuler.net/problem=7)

## 10001st prime

```
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10001st prime number?

```

### Thought Process

We need to calculate out to the 10001st prime number and return what it's value is.

We'll implement a new Sieve (SieveofAtkin) that's way more efficient than SieveofEratosthenes.
We'll provide an arbitrary start number (1000000) and count how many primes are returned; assuming that 
number is more than 10001 we'll have it return the 10000th index and get our answer.