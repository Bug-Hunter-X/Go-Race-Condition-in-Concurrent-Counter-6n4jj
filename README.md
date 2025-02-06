# Go Race Condition Example

This repository demonstrates a common race condition in Go when multiple goroutines access and modify a shared variable without proper synchronization.

## Bug Description
The `bug.go` file contains a program that launches multiple goroutines to increment a counter. Because the increment operation (`count += i`) is not atomic, data races can occur, resulting in an incorrect final count. 

## Solution
The `bugSolution.go` file shows how to fix the race condition using a mutex to protect the shared counter.  This ensures that only one goroutine can access and modify the counter at any given time, preventing data corruption and providing the correct final count.