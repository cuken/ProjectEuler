# Problem 22

[Link to Problem](https://projecteuler.net/problem=22)

## Names scores

```
Using <a href="project/resources/p022_names.txt">names.txt</a> (right click and &#39;Save Link/Target As...&#39;), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.
For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.
What is the total of all the name scores in the file?

```

### Thought Process

We need to ingest the names text file and then sort it alphabetically.
Then we need to run each name through a function to determine its "point" value. 
Once it's passed through that function we'll iterate through the list and multiply it's position by it's total score.

