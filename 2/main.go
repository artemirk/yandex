package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "bufio"
)

func main() {
    fi, _ := os.Open("input.txt")
    defer fi.Close()

    scanner := bufio.NewScanner(fi)
    scanner.Scan()
    j := scanner.Text()

    scanner.Scan()
    s := scanner.Text()

    jewells := make(map[rune]struct{}, len(s))

    for _, jw := range j {
        jewells[jw] = struct{}{}
    }

    jwCount := 0
    for _, st := range s {
        if _, ok := jewells[st]; ok {
            jwCount++
        }
    }

    ioutil.WriteFile("output.txt", []byte(fmt.Sprint(jwCount)), 0644)
}
