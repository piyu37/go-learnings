package main

import "fmt"

type pos struct {
	x, y int
}

// Time: O(n^2), Space: O(n^2), where n = max(|x|, |y|), d = answer (moves).
//
// Why:
// 1. A knight moves at most 2 squares per axis per move, so d ~ n/2 -> O(n).
//
//  2. BFS stops at level d. In d moves the knight stays inside x, y in
//     -2d .. 2d, i.e. a box around the start.
//     - One side has 4d + 1 values: 2d negatives + zero + 2d positives.
//     - The board is a grid: x picks the column, y picks the row, so
//     squares = columns x rows = (4d+1) x (4d+1).
//     d = 1 (K = start, N = squares one move away):
//     x: -2 -1  0  1  2
//     y =  2    .  N  .  N  .
//     y =  1    N  .  .  .  N
//     y =  0    .  .  K  .  .     5 columns x 5 rows = 25 squares
//     y = -1    N  .  .  .  N
//     y = -2    .  N  .  N  .
//     The knight never leaves the box, so BFS visits at most
//     (4d+1)^2 ~ 16 * (n/2)^2 = 4 * n^2 -> O(n^2) squares.
//
//  3. Each square is enqueued once (visited map) with 8 moves tried -> O(n^2).
//     The fence (x, y >= -2) only cuts work ~4x, still O(n^2).
//
//  4. Space: the visited map holds every visited square -> O(n^2).
//     (The queue holds one level at a time, which is smaller.)
//
// Example: target (20, 0) -> d = 10, box 41 x 41 = 1681 squares at most.
// Target (40, 0) -> 81 x 81 = 6561, ~4x more. Double n -> 4x work = O(n^2).
//
// Interview one-liner: "The answer is ~n/2 moves, and BFS only visits squares
// within that many moves, an area of about n by n, so time and space are O(n^2)."
func minKnightMovesUsingBFS(x int, y int) int {
	x, y = abs(x), abs(y)

	if x == 0 && y == 0 {
		return 0
	}

	queue := make([]pos, 0)
	visited := make(map[pos]bool)

	queue = append(queue, pos{0, 0})
	visited[pos{0, 0}] = true

	coordinates := [][]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}}
	steps := 1

	for len(queue) > 0 {
		tmpQueue := make([]pos, 0)

		for i := range queue {
			curr := queue[i]

			for _, cor := range coordinates {
				next := pos{curr.x + cor[0], curr.y + cor[1]}

				if next.x == x && next.y == y {
					return steps
				}

				// Fence: skip squares too far into negative x or y.
				//
				// Why a fence at all? After abs(), the target is always in the
				// first quadrant (x >= 0, y >= 0). Wandering far into negative
				// coordinates can never give a shorter path, so cutting those
				// squares out saves most of the BFS work.
				//
				// Why -2 and not 0? Some shortest paths near the origin have to
				// step slightly below zero and come back.
				//
				// Example: target (1, 1). The shortest path is 2 moves:
				//     (0,0) -> (2,-1) -> (1,1)
				// The middle square (2,-1) has a negative y.
				//
				// If the fence were at 0 (only x >= 0 and y >= 0 allowed):
				//   - From (0,0), the only allowed moves are (1,2) and (2,1).
				//   - Neither of those can reach (1,1) in one knight move.
				//   - So a 2-move answer is impossible inside the fence.
				//   - A 3-move answer is also impossible: every knight move
				//     changes (x + y) by an odd number, and (0,0) and (1,1)
				//     both have an even sum, so it always takes an even
				//     number of moves.
				//   - The shortest path that stays inside the fence is 4:
				//     (0,0) -> (2,1) -> (0,2) -> (2,3) -> (1,1)
				// So BFS would wrongly return 4 instead of 2.
				//
				// A knight moves at most 2 squares in any direction per step,
				// so allowing coordinates down to -2 leaves room for these
				// short detours.
				if next.x < -2 || next.y < -2 {
					continue
				}

				if !visited[next] {
					tmpQueue = append(tmpQueue, next)
					visited[next] = true
				}
			}
		}

		queue = tmpQueue

		steps++
	}

	return -1
}

// Time: O(x * y), Space: O(x * y), where x, y are already abs().
// In terms of n = max(|x|, |y|) that is O(n^2) worst case, same as BFS.
//
// Why:
//  1. Each call goes to (x-1, y-2) or (x-2, y-1), so x and y only move
//     toward 0 (abs() can bounce them to 1 or 2 near the edges, no higher).
//     So every state (a, b) has a in 0..x and b in 0..y: a grid of
//     (x+1) columns * (y+1) rows -> O(x * y) distinct states.
//  2. The memo map solves each state once, and each solve is O(1) work
//     (2 memo lookups + min) -> time O(x * y).
//  3. Space: the memo map holds every state -> O(x * y). The recursion stack
//     is smaller: x + y drops by about 3 per call, so depth is O(x + y).
//
// Example: target (20, 20) -> at most 21 x 21 = 441 states.
// Target (40, 40) -> 41 x 41 = 1681, ~4x more. Double n -> 4x work = O(n^2).
//
// vs BFS: same O(n^2) worst case, but DFS tries only 2 moves per state, both
// heading toward (0, 0), and only touches the rectangle (0,0)..(x,y), so it
// does much less work in practice.
//
// Interview one-liner: "Memo solves each (x, y) once, and x, y only shrink
// toward 0, so there are at most about x * y states -> O(x * y) time and
// space, which is O(n^2) in the worst case."
func minKnightMovesUsingDFS(x int, y int) int {
	memo := map[pos]int{}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		// Symmetry: same idea as the abs() in the BFS version.
		x, y = abs(x), abs(y)

		// Base cases.
		if x+y == 0 {
			return 0 // already at (0,0)
		}
		if x+y == 2 {
			return 2 // (1,1), (2,0) and (0,2) each take exactly 2 moves
		}

		key := pos{x, y}
		if v, ok := memo[key]; ok {
			return v
		}

		// Only the two moves that head back toward (0,0).
		res := min(dfs(x-1, y-2), dfs(x-2, y-1)) + 1
		memo[key] = res
		return res
	}

	return dfs(x, y)
}

// https://leetcode.com/problems/minimum-knight-moves/description/
func minKnightMovesMain() {
	fmt.Println(minKnightMovesUsingBFS(2, 1))
	fmt.Println(minKnightMovesUsingDFS(2, 1))
}
