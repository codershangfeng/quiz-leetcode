/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Benefit struct {
	Previous *Benefit
	Amount   int
}

func (n Benefit) String() string {
	return fmt.Sprintf("Node{Weight: %d}", n.Amount)
}

func maxProfit(prices []int) int {
	benefits := []Benefit{}

	benefits = append(benefits, Benefit{Amount: 0})

	for i := 1; i < len(prices); i++ {
		parent := &benefits[len(benefits)-1]
		benefits = append(benefits, Benefit{Amount: prices[i] - prices[i-1], Previous: parent})
	}

	var maxTotalBenefit int
	for _, benefit := range benefits {
		if benefit.Amount > 0 {
			maxTotalBenefit += benefit.Amount
		}
	}

	return maxTotalBenefit
}

func TestCase1(t *testing.T) {
	mocked := []int{7, 1, 5, 3, 6, 4}

	got := maxProfit(mocked)

	assert.Equal(t, 7, got)
}

func TestCase2(t *testing.T) {
	mocked := []int{1, 2, 3, 4, 5}

	got := maxProfit(mocked)

	assert.Equal(t, 4, got)
}

func TestCase3(t *testing.T) {
	mocked := []int{7, 6, 4, 3, 1}

	got := maxProfit(mocked)

	assert.Equal(t, 0, got)
}

func TestCase4(t *testing.T) {
	mocked := []int{2, 1, 3, 7, 12}

	got := maxProfit(mocked)

	assert.Equal(t, 11, got)
}
