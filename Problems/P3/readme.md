# Problem 3

## Largest prime factor

```
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
```

## Thought Process

The number will not fit inside of int32 so we're going to need to use int64 to handle the target number.
I'll start by dividing the number in 2, as the largest factor at most should be target / 2.

I think the trick to this problem will be finding a fast way to discover primes up to a target;

Maybe we'll try the https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes method?
