package main

import (
	"fmt"
	"math"
	"strconv"
)

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func castle(layer int, usedNumbers map[int]bool, solution [][]int) [][]int {
	if layer == 5 {
		return solution
	}

	layerSizes := []int{5, 4, 3, 2, 1}
	currentSize := layerSizes[layer]

	if layer == 0 {
		nums := make([]int, 15)
		for i := 0; i < 15; i++ {
			nums[i] = i + 1
		}
		firstLayerPermutations := permutations(nums)
		for _, firstLayer := range firstLayerPermutations {
			if len(firstLayer) >= currentSize {
				firstLayerSlice := firstLayer[:currentSize]
				newUsedNumbers := make(map[int]bool)
				for _, num := range firstLayerSlice {
					newUsedNumbers[num] = true
				}
				result := castle(layer+1, newUsedNumbers, append(solution, firstLayerSlice))
				if result != nil {
					return result
				}
			}
		}
		return nil
	}

	oldLayer := solution[len(solution)-1]
	nums := make([]int, 15)
	for i := 0; i < 15; i++ {
		nums[i] = i + 1
	}
	nextLayerPermutations := permutations(nums)

	for _, nextLayer := range nextLayerPermutations {
		if len(nextLayer) >= currentSize {
			nextLayerSlice := nextLayer[:currentSize]
			valid := true
			for d := 0; d < currentSize-1; d++ {
				if math.Abs(float64(oldLayer[d]-oldLayer[d+1])) != float64(nextLayerSlice[d]) {
					valid = false
					break
				}
				if usedNumbers[nextLayerSlice[d]] {
					valid = false
					break
				}
			}
			if currentSize > 0 && usedNumbers[nextLayerSlice[currentSize-1]] {
				valid = false
			}

			if valid {
				newUsedNumbers := make(map[int]bool)
				for k, v := range usedNumbers {
					newUsedNumbers[k] = v
				}
				for _, num := range nextLayerSlice {
					newUsedNumbers[num] = true
				}
				result := castle(layer+1, newUsedNumbers, append(solution, nextLayerSlice))
				if result != nil {
					return result
				}
			}
		}
	}

	return nil
}

func main() {
	solution := castle(0, make(map[int]bool), [][]int{})
	if solution != nil {
		for _, layer := range solution {
			fmt.Println(layer)
		}
	} else {
		fmt.Println("no solution")
	}
}