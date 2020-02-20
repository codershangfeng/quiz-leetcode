/* 1266
On a plane there are n points with integer coordinates points[i] = [xi, yi]. Your task is to find the minimum time in seconds to visit all points.

You can move according to the next rules:

In one second always you can either move vertically, horizontally by one unit or diagonally (it means to move one unit vertically and one unit horizontally in one second).
You have to visit the points in the same order as they appear in the array.


Example 1:
Input: points = [[1,1],[3,4],[-1,0]]
Output: 7
Explanation: One optimal path is [1,1] -> [2,2] -> [3,3] -> [3,4] -> [2,3] -> [1,2] -> [0,1] -> [-1,0]
Time from [1,1] to [3,4] = 3 seconds
Time from [3,4] to [-1,0] = 4 seconds
Total time = 7 seconds

Example 2:
Input: points = [[3,2],[-2,2]]
Output: 5

Constraints:

points.length == n
1 <= n <= 100
points[i].length == 2
-1000 <= points[i][0], points[i][1] <= 1000
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minTimeToVisitAllPoints(points [][]int) int {
	minTime := 0

	length := len(points) - 1

	for i := 0; i < length; i++ {
		start := points[i]
		end := points[i+1]

		dx := start[0] - end[0]
		if dx < 0 {
			dx *= -1
		}

		dy := start[1] - end[1]
		if dy < 0 {
			dy *= -1
		}

		if dx > dy {
			minTime += dx
		} else {
			minTime += dy
		}

	}

	return minTime
}

func TestCase1(t *testing.T) {
	got := minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {1, 0}})

	assert.Equal(t, 7, got)
}

func TestCase2(t *testing.T) {
	got := minTimeToVisitAllPoints([][]int{{3, 2}, {-2, 2}})

	assert.Equal(t, 5, got)
}
