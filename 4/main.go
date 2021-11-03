package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/artemirk/yandex/4/grid"
)

type cellsAreaResult struct {
	Sum   int
	Cells []grid.Cell
}

func main() {
	fi, _ := os.Open("input.txt")
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Printf("area: %v\n", n)

	g := grid.NewGrid()

	iRow := 0

	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		rowValues := make([]int, len(vals))
		for i, v := range vals {
			rowValues[i], _ = strconv.Atoi(v)
		}
		fmt.Printf("#%v row: %v\n", iRow, rowValues)
		g.AddValues(iRow, rowValues)
		iRow++
	}

	areas := areas(n, g.Width, g.Height)
	fmt.Printf("areas: %v for width: %v heigth: %v\n", areas, g.Width, g.Height)
	if len(areas) == 0 {
		fmt.Printf("cannot find any areas: for width: %v heigth: %v with size: %v\n", g.Width, g.Height, n)
		return
	}

	fmt.Printf("grid: %v\n", g.Cells)

	resultsChan := make(chan cellsAreaResult, len(areas))

	var wg sync.WaitGroup
	wg.Add(len(areas))

	fmt.Printf("start\n")
	for _, area := range areas {
		go func(a grid.Area) {
			defer wg.Done()
			sum(a, g, resultsChan)
		}(area)
	}

	go func() {
		defer close(resultsChan)
		wg.Wait()
		fmt.Printf("done\n")
	}()

	results := make(map[int][]cellsAreaResult)
	var maxValue int
	isFirst := true
	for result := range resultsChan {
		if _, ok := results[maxValue]; !ok {
			results[result.Sum] = make([]cellsAreaResult, 0, 1)
		}
		results[result.Sum] = append(results[result.Sum], result)
		if isFirst {
			maxValue = result.Sum
			isFirst = false
		} else if maxValue < result.Sum {
			maxValue = result.Sum
		}
	}

	for _, result := range results[maxValue] {
		outputResult(result, g.Width, g.Height)
	}

	ioutil.WriteFile("output.txt", []byte(fmt.Sprint(maxValue)), 0644)

}

func areas(n int, width int, height int) []grid.Area {
	m := make([]grid.Area, 0, 2)
	for i := 1; float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			j := n / i
			if i <= width && j <= height {
				m = append(m, grid.Area{Width: i, Height: j})
			}
			if i != j && j <= width && i <= height {
				m = append(m, grid.Area{Width: j, Height: i})
			}
		}
	}
	return m
}

func sum(a grid.Area, g *grid.Grid, areaResultsChan chan cellsAreaResult) {
	i := 0
	for i < len(g.Cells) {
		i = processAreaFromCell(i, a, g, areaResultsChan)
	}
}

func processAreaFromCell(startIndex int, a grid.Area, g *grid.Grid, areaResultsChan chan cellsAreaResult) int {
	result := cellsAreaResult{
		Cells: make([]grid.Cell, a.Width*a.Height),
	}
	c := 0
	for i := 0; i < a.Width; i++ {
		for j := 0; j < a.Height; j++ {
			idx := startIndex + i + j*g.Width
			cell := g.Cells[idx]
			result.Cells[c] = cell
			result.Sum += cell.Value
			c++
		}
	}
	areaResultsChan <- result

	for startIndex < len(g.Cells) {
		startIndex++

		iCol := startIndex % g.Width
		iRow := (startIndex - iCol) / g.Width

		if (g.Width-iCol) >= a.Width && (g.Height-iRow) >= a.Height {
			break
		}
	}
	return startIndex
}

func outputResult(result cellsAreaResult, width int, height int) {
	fmt.Printf("sum : %v\n", result.Sum)

	buff := bytes.NewBufferString("")

	colValues := make([]string, width)
	gridCellsValues := make(map[int]map[int]int, width)

	for _, c := range result.Cells {
		if _, ok := gridCellsValues[c.Row]; !ok {
			gridCellsValues[c.Row] = make(map[int]int, 1)
		}
		gridCellsValues[c.Row][c.Col] = c.Value
	}

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			colValues[i] = "-"
			if v, ok := gridCellsValues[j][i]; ok {
				colValues[i] = fmt.Sprintf("%v", v)
			}
		}
		buff.WriteString(strings.Join(colValues, " "))
		buff.WriteString("\n")
	}
	fmt.Print(buff.String())

}
