# Problem 14

[Link to Problem](https://projecteuler.net/problem=14)

## Longest Collatz sequence

```
The following iterative sequence is defined for the set of positive integers:
<p style="margin-left:50px;"><var>n</var> → <var>n</var>/2 (<var>n</var> is even)<br/><var>n</var> → 3<var>n</var> + 1 (<var>n</var> is odd)
Using the rule above and starting with 13, we generate the following sequence:
<div style="text-align:center;">13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1</div>
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
Which starting number, under one million, produces the longest chain?
<p class="note"><b>NOTE:</b> Once the chain starts the terms are allowed to go above one million.

```

### Thought Process

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Brute Force: start at either 1 million -1 or 1 and itterate through every permiatation; counting the chain length and finding the maximum chain, recording the number and returning that number.

Slightly less brutish method: We can check the first 1000, terms and see what the largest value is...

2 -> 1
4 -> 2, 1

20 -> 10, 5, -> 
40 -> 20, 10, 5 -> 
50 -> 25, 76, 38, 14, 7 -> 

find all primes up to 1mil; 
calcualte chain length for each prime.
