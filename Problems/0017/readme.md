# Problem 17

[Link to Problem](https://projecteuler.net/problem=17)

## Number letter counts

```
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used? 
<br/><p class="note"><b>NOTE:</b> Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of &#34;and&#34; when writing out numbers is in compliance with British usage.
```

### Thought Process

We need to itterate from 1 to 1000, creating a function that will convert our value to "words" and then count the length of the word. Add all of the counts of words to a sum and return the sum.

The tricky part will be determining, the 10's place of the number and how it changes it, from one to ten, to one hundered, etc.