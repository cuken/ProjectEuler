# Problem 5

## Smallest multiple

```
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
```

## Thought Process

ALL positive intergers are divisible by 1.
We know the number must be even as it has to cleanly divide by 2.
We know the number must end in a 0 or 5 to be divisible by 5.
We know the number must end in a 0 to be divisible by 10.
We know the numbers digits must add up to a divisible number of 3.

We're looking for the smallest, so we'll start at 20 (the minimum for the number's that have to cleanly divide) and increment by 2.

Initial Thoughts; We could do a bunch of if statements checking if it is %x == 0 and increment up through x. That would take an increasing amount of time as we wanted to check for more than just 20 digits.

We can take shortcuts using number theory above to limit the amount of itterations we go through;