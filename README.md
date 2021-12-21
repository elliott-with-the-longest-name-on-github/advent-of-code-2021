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
  </tbody>
</table>
