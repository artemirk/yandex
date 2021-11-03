package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
    fi, _ := os.Open("input.txt")
    defer fi.Close()

    scanner := bufio.NewScanner(fi)
    scanner.Scan()
    k, _ := strconv.Atoi(scanner.Text())

    scanner.Scan()
    m_x_n := strings.Split(scanner.Text(), " ")
    m, _ := strconv.Atoi(m_x_n[0])
    n, _ := strconv.Atoi(m_x_n[1])

    dVals := make([]int, 0, k)

    maxD := 0
    sumD := 0

    var dVal int

    for scanner.Scan() {
        d, _ := strconv.Atoi(scanner.Text())
        if len(dVals) != k {
            dVals = append(dVals, d)
            sumD += d
            if len(dVals) == k {
                maxD = sumD
            }
            continue
        }
        dVal, dVals = dVals[0], dVals[1:]
        sumD -= dVal
        sumD += d
        dVals = append(dVals, d)
        if maxD < sumD {
            maxD = sumD
        }
    }

    ioutil.WriteFile("output.txt", []byte(fmt.Sprint(maxD)), 0644)
}
