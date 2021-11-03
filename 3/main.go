package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
    fi, _ := os.Open("input.txt")
    defer fi.Close()

    scanner := bufio.NewScanner(fi)
    scanner.Scan()
    n_x_k := strings.Split(scanner.Text(), " ")
    n, _ := strconv.Atoi(n_x_k[0])
    x, _ := strconv.Atoi(n_x_k[1])
    k, _ := strconv.Atoi(n_x_k[2])

    scanner.Scan()
    d := strings.Split(scanner.Text(), " ")
    deads := make([]int, n)
    for i, dl := range d {
        deads[i], _ = strconv.Atoi(dl)
    }

    sort.Ints(deads)

    run := true
    numIterations := 0

    reportResultsLookup := make(map[int]struct{}, n)
    reportResults := make([]int, 0, n)

    for run {
        for j := range deads {
            if _, ok := reportResultsLookup[deads[j]]; !ok {
                reportResultsLookup[deads[j]] = struct{}{}
                reportResults = append(reportResults, deads[j])
                numIterations++
                if numIterations == k {
                    run = false
                    break
                }
            }
            deads[j] += x
        }
    }
    sort.Ints(reportResults)
    fmt.Println(reportResults)

    ioutil.WriteFile("output.txt", []byte(fmt.Sprint(reportResults[len(reportResults)-1])), 0644)
}
