# Problem 9

[Link to Problem](https://projecteuler.net/problem=9)

## Special Pythagorean triplet

```
A Pythagorean triplet is a set of three natural numbers, <var>a</var> &lt; <var>b</var> &lt; <var>c</var>, for which,
<div style="text-align:center;"> <var>a</var><sup>2</sup> + <var>b</var><sup>2</sup> = <var>c</var><sup>2</sup></div>
For example, 3<sup>2</sup> + 4<sup>2</sup> = 9 + 16 = 25 = 5<sup>2</sup>.
There exists exactly one Pythagorean triplet for which <var>a</var> + <var>b</var> + <var>c</var> = 1000.<br/>Find the product <var>abc</var>.

```

### Thought Process

Brute Force: check if a,b,c is a Pythagorean triplet; if it is, check if a+b+c = 1000, if it does stop.

Use 3 loops to itterate through a b and c and check if it's a triplet and equal to 1000