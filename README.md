# Advent of Code 2021

## Getting Started

- Clone the repository
- Each folder is a day's challenge. The challenge is listed in that folder's README.
- The input for the challenge is normally in the package at the bottom. This is just so I don't have to fiddle with reading from a file.

## Benchmarking

```bash
# Assumes you're in the project root

# Run all tests, including benchmarks
go test -bench . ./...

# Run tests for a specific subdirectory
DIR='20'
go test -v ./$DIR/...
```

## Results
<table>
  <thead>
    <tr>
      <th>Day</th>
      <th>Title</th>
      <th>Part One (ms)</th>
      <th>Part Two (ms)</th>
      <th>Notes</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>01</td>
      <td>Sonar Sweep</td>
      <td>N/A</td>
      <td>N/A</td>
      <td>Didn't measure benchmarks simply because I wasn't benchmarking at this stage.</td>
    </tr>
    <tr>
      <td>02</td>
      <td>Dive!</td>
      <td>N/A</td>
      <td>0..0014</td>
      <td>Easy peasy.</td>
    </tr>
    <tr>
      <td>03</td>
      <td>Binary Diagnostic</td>
      <td>0.116</td>
      <td>0.155</td>
      <td>Pretty easy, though it was definitely my first time bit-pushing with Go. Not altogether a bad experience.</td>
    </tr>
    <tr>
      <td>04</td>
      <td>Giant Squid (Bingo)</td>
      <td>0.227</td>
      <td>0.372</td>
      <td>Not too bad. I seem to remember that efficiency was annoying, but it still ran, well, faster than is ever necessary.</td>
    </tr>
    <tr>
      <td>05</td>
      <td>Hydrothermal Venture</td>
      <td>34.193</td>
      <td>57.359</td>
      <td>Not too difficult from an execution standpoint. Essentially, the interesting part was parsing the input into lines. Since the coordinate system only accepted integers, each line could be broken into a finite set of points and mapped to its count, making filtering for "Number of coordinates with count greater than X" pretty easy.</td>
    </tr>
    <tr>
      <td>06</td>
      <td>Lanternfish</td>
      <td>0.046</td>
      <td>0.129</td>
      <td>Not much to say about this one -- pretty trivial, as long as you're not trying to track it through a list or stacks. Maps are the best!</td>
    </tr>
    <tr>
      <td>07</td>
      <td>The Treachery of Whales</td>
      <td>0.065</td>
      <td>0.021</td>
      <td>This one was pretty interesting. The solution to use the median in the first part was obvious to me, but the mean in the second part took me a bit to figure out. I'm sure there's a better way to resolve split medians and figure out whether to floor or ceiling the result of mean, but brute forcing it was still super fast. This is the first one where I've seen faster times on Part 2 than Part 1, simply because calculating the mean doesn't require a list sort like the median does.</td>
    </tr>
    <tr>
      <td>09</td>
      <td>Smoke Basin</td>
      <td>6.155</td>
      <td>8.569</td>
      <td>This was a super fun depth-first search algorithm to implement. So fun, in fact, that I packaged this one with a super cool terminal visualization. It should be pretty self-explanatory how to view it -- just look in <code>main.go</code>.</td>
    </tr>
    <tr>
      <td>10</td>
      <td>Syntax Scoring</td>
      <td>0.607</td>
      <td>0.652</td>
      <td>I found this one to be interesting and fun. I didn't particularly prioritize bleeding fast performance; rather, I did my best from the outset to build a legit tag parsing system. I figured I'd have to do something with incomplete tags, so I went ahead and dealt with them during Part One. The design means that I pay the penalty of parsing lines that I could throw away for both Parts One and Two, but it also means that I <em>could</em> share a <code>LogDump</code> object between them if this were a real-world scenario. The looping/recursion combination I used for parsing everything in Part One translated over to Part Two <em>extremely</em> well -- for Part Two, I only had to add one line of code to the existing base, plus the scoring logic for incomplete lines.</td>
    </tr>
    <tr>
      <td>13</td>
      <td>Transparent Oragami</td>
      <td>0.113</td>
      <td>0.654</td>
      <td>I was really a fan of this one. Super fun to print each stage, adn the fact that the solution was a visual solution was really neat. Seems to be a pretty performant solution, too -- I used a map of a string representation of each dot to its "data structure" representation. This made merging the dots after a fold super easy, and it also made counting visible dots really inexpensive, because I can just <code>len(map)</code>.</td>
    </tr>
    <tr>
      <td>14</td>
      <td>Extended Polymerization</td>
      <td colspan="2">
      <details>
        <summary>Iterative benchmarks</summary>
        <ul>
          <li>10 iterations: 0.039, +/-0</li>
          <li>20 iterations: 0.078, +39</li>
          <li>30 iterations: 0.116, +38</li>
          <li>40 iterations: 0.154, +38</li>
          <li>50 iterations: 0.206, +52</li>
          <li>60 iterations: 0.230, +24</li>
          <li>70 iterations: 0.275, +45</li>
          <li>80 iterations: 0.305, +30</li>
          <li>90 iterations: 0.364, +59</li>
          <li>100 iterations: 0.391, +37</li>
          <li>...</li>
          <li>150 iterations: 0.610, +65</li>
          <li>...</li>
          <li>200 iterations: 0.796, +67</li>
        </ul>
      </details>
      </td>
      <td>This was unquestionably my favorite one so far. I was one of the chumps who brute forced it with a string at the beginning, only to be slapped in the face with the exponential inefficiencies in the next section. I rewrote the solution as a map, which was not too hard at all. Since Part One and Part Two are the same thing but with more iterations, I decided to bench this one a little differently. As you can see, it runs in roughly <code>O(N)</code> time, which makes sense, as there are only a few possible polymer pairs and that number is reached very early on.</td>
    </tr>
    <tr>
      <td>15</td>
      <td>Chiton</td>
      <td>8.523</td>
      <td>199.635</td>
      <td>This one was easy only because I cheated... sort of. I knew Dijkstra's or A* was pretty much the way to go for this, so I found a library for Dijkstra's. No point in implementing very well-defined functionality like that. As usual, parsing the input was the hardest part -- it's times like these when I <em>really</em> wish Go had <code>array.Map</code> and company. </td>
    </tr>
    <tr>
      <td>21</td>
      <td>Dirac Dice</td>
      <td>0.005</td>
      <td>413.700</td>
      <td>Well isn't that a jump from Part One to Part Two. Part one was fun and I had a bit too much fun with the solution. Part Two kicked my butt, as my previous design pretty much didn't support what I needed to do at all. I'm still not super happy with the performance, but it's memoized already and I really don't want to optimize more.</td>
    </tr>
  </tbody>
</table>
