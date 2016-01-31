package main

import (
	"testing"
	"fmt"
)

func TestIp(t *testing.T) {
    c := rds()
    num := ip(c, "baidu")
    fmt.Println("ip is ", num)
}

func TestPv(t *testing.T) {
    c := rds()
    num := pv(c, "baidu")
    fmt.Println("pv is ", num)
}

func TestUv(t *testing.T) {
    c := rds()
    num := uv(c, "baidu")
    fmt.Println("uv is ", num)
}