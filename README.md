# AOC_2023

## Summary

Benchmarked on AMD Ryzen 7 5800HS.

| Day | Done on | Part 1 | Part 2 | Comments (short)                                                                                                                                                                                                                            |
|-----|---------|--------|--------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 01  | 01      | 1.0 ms | 1.0 ms | -  Used strings package  for the first time.<br/> - Solved by scanning from left to right.                                                                                                                                                  |
| 02  | 03      | 1.0 ms | 1.0 ms | Easy. Used split for the first time.                                                                                                                                                                                                        |
| 03  | 03      | 1.0 ms | 1.0 ms | - Used unicode package for the first time. <br/> - Found this one challenging when I did it...<br/> - Not optimized : resorted to checking the neighborhood of each digit, instead of checking the neighborhood of each symbol/star symbol. |
| 04  | 04      | 1.0 ms | 1.0 ms | Easy. Started writing tests and a utils package.                                                                                                                                                                                            |
| 05  | 05      | 1.0 ms | 1.0ms  | Time-consuming, solved with intervals.                                                                                                                                                                                                      |
| 06  | 06      | 500 µs | 500 µs | Easiest day, solved naïvely (part one) and through a quadratic equation (part two).<br/>                                                                                                                                                  |
| 07  | 07      | 1.0 ms | 1.0 ms | - Solved with the sort package. <br/> - Utils package finally usable.                                                                                                                                                                       |
| 08  | 08      | 1.0 ms | 2.5 ms | Solved via LCM.                                                                                                                                                                                                                             |
| 09  | 10      | 1.0 ms | 1.0 ms | Easy.                                                                                                                                                                                                                                       |
| 10  | 10      | 5.0 ms | 8.0 ms | Time-consuming, part two solved with ray-tracing.                                                                                                                                                                                           |
| 11  | 11      | 6.0 ms | 6.0 ms | Easy.                                                                                                                                                                                                                                       |
| 12  | 12      | N/A    | N/A    | N/A                                                                                                                                                                                                                                         |
| 13  | 13      | N/A    | N/A    | N/A                                                                                                                                                                                                                                         |
| 14  | 14      | 1.0ms  | 5.0 s  |  Solution neither elegant nor optimized.                                                                                                                                                             |
| 15  | 15      | 1.0 ms | 1.0 ms | Relaxing albeit difficult to understand at first.                                                                                                                                                                                           |
| 16  | 16      | 4.0 ms | 700 ms | Solved recursively.                                                                                                                                                                                                                         |
| 17  | 17      | 300 ms | 1.0 s  | - Time-consuming.<br/> - Solved with A* (RedGlobGames).  <br/> - Borrowed a priority queue implementation from the Internet.                                                                                                                |
| 18  | 18      | 1.0 ms | 1.0 ms | - Easy once the shoelace algorithm and Pick's theorem are applied. <br/> - Once again wasted time because I can't read instructions.                                                                                                        |
| 19  | 19      | 1.0 ms | 1.0 ms | - Took less time than expected. <br/> - Uses the Interval class implemented during day 5. <br/> - Part two solved recursively.                                                                                                              |
| 20  | 20      | 10 ms  | 150 ms | - Tried to implement a sum type, but not possible in go, so found a workaround on the Internet.<br/> - Only done part one, solved with a FIFO queue. <br/> - Time-consuming. <br/> - Part two done on day 23.                               |
| 21  | 21      | 50 ms  | N/A    | - Listened to my voice of reason and gave up on it. <br/> - Part one is not even elegant, disappointed in myself to say the least.                                                                                                          |
| 22  |         |        |        |                                                                                                                                                                                                                                             |
| 23  |         |        |        |                                                                                                                                                                                                                                             |
| 24  |         |        |        |                                                                                                                                                                                                                                             |
| 25  |         |        |        |                                         

## What I learned.

Learned about ray-tracing, cycle detection in integer sequences, A*, pathfinding, the shoelace algorithm.

## Extended remarks.

### Day 5.

Tried to brute force part 2 before realizing that the computing time was too long.
After that, it wasn't hard to find a theoretical way to answer part 2 per se, what took time was to implement my solution to the problem which required an Interval class.


### Day 7.

Had to reread the instructions multiple times. Made some mistakes in my less function and got confused while writing getType.
Had to look up another test input on reddit.

Sorting is done via the interface provided by the sort package.

Part two was quicker to implement, the best strategies of wildcard replacement were found by hand and eventually led to getType2.

### Day 8.

Answer given via LCM, as discussed with my classmates. 
Looking back after solving the problem, I realized this solution shouldn't work without additional hypotheses on the input.

### Day 10.

Time-consuming because I was unfamiliar with the solving techniques involved and had to do some research.
Part two solved by ray tracing.

### Day 12 and 13.

Day 13's first star done on day 15.
The rest is not done yet.

### Day 14.

Day 14 part 2 solves the problem under the assumption that a cycle in the rocks movement will appear within 1000 tilt cycles in about 5 seconds.

Meddled with Floyd's cycle detection algorithm before giving up on it as it is not really applicable here.

### Day 17.
Immediately thought of A* pathfinding, transposed RedBlobGames' A* in Go. 

Frustrated because I got confused by what to actually put in my priority queue (A state and its priority) and my sets (just the state WITHOUT its priority).
Part two was much quicker to solve.

### Day 19.
First time my code almost works on first try. I love recursion.

### Day 20.
Had issues with my antivirus software, which didn't allow go to run. Forgot to add a stop case if there are no cycles within the 1000 button presses.

All in all, solving part one was mostly spent on modelling the input in multiple aliases and struct types.
Found a workaround to make a Module sum type. 

Part 2 done on day 23.


