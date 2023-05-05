# A game of Pig
Pig is an cli based program that simulates dice game between 2 players

## Features

- Supports single-single, single multi, multi multi strategy



## How to use
- clone this repository [git clone https://github.com/PratikJethe/go-a-game-of-pig.git]
- run [go build .] in root of project
- execute binary with appropriate input

## Examples 

### 1. Single Single Strategy
 
```sh
$ ./pig 10 15
Holding at 10 vs Holding at  15: wins: 1/10 (10.00%), losses: 9/10 (90.00%)
```

### 2. Single Multi Strategy
 
```sh
$ ./pig 21 1-100
Holding at 21 vs Holding at  1: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  2: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  3: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  4: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  5: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  6: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  7: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  8: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  9: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  10: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  11: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  12: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  13: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  14: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  15: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  16: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  17: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  18: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  19: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  20: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
Holding at 21 vs Holding at  22: wins: 10/10 (100.00%), losses: 0/10 (0.00%)
.
.
Holding at 21 vs Holding at  99: wins: 0/10 (0.00%), losses: 10/10 (100.00%)
Holding at 21 vs Holding at  100: wins: 0/10 (0.00%), losses: 10/10 (100.00%)
```
### 3. Multi Multi Strategy
 
```sh
$ ./pig 1-100 1-100
Result: Wins, losses staying at k = 1: 27/990 (2.73%), 963/990 (97.27%)
Result: Wins, losses staying at k = 2: 23/990 (2.32%), 967/990 (97.68%)
Result: Wins, losses staying at k = 3: 38/990 (3.84%), 952/990 (96.16%)
.
.
.
Result: Wins, losses staying at k = 97: 968/990 (97.78%), 22/990 (2.22%)
Result: Wins, losses staying at k = 98: 970/990 (97.98%), 20/990 (2.02%)
Result: Wins, losses staying at k = 99: 975/990 (98.48%), 15/990 (1.52%)
Result: Wins, losses staying at k = 100: 979/990 (98.89%), 11/990 (1.11%)