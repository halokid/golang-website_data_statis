/**
* Copyright 0x029 Inc. 
* License: MIT License
* Author: JJyy<82049406@qq.com>
* This file is .... 
**/
package main

import (
	"fmt"
	"testing"
	// "strconv"
)

func TestIp(t *testing.T) {
	c := rds()
	num := ip(c, "baidu", 1454035404, 1454035410)
	fmt.Println("ip is ", num)

	// test_int, _ := strconv.ParseInt( "1454035403",  10, 32 )
	// test_int, _ := strconv.ParseUint( "1454035403.8097",  10, 64 )
	// test_int, _ := strconv.ParseFloat( "1454035403.8097", 64 )
	// test_int, _ := strconv.ParseFloat( "1454035403.8097", 64 )
	// aa := strconv.FormatFloat(test_int, 'f', 4, 32)
	// test_int, _ := strconv.FormatFloat( "1454035403.8097", 64 )
	// fmt.Println("test_int is ", test_int)
	// fmt.Println("aa ", aa)
}

func TestPv(t *testing.T) {
	c := rds()
	num := pv(c, "baidu", 1454035404, 1454035410)
	fmt.Println("pv is ", num)
}

func TestUv(t *testing.T) {
	c := rds()
	num := uv(c, "baidu", 1454035404, 1454035410)
	fmt.Println("uv is ", num)
}
