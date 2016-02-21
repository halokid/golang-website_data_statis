/**
* Copyright 0x029 Inc. 
* License: MIT License
* Author: JJyy<82049406@qq.com>
* get the src route info for user 
**/

package main

import (
	"gopkg.in/redis.v3"
	// "strconv"
	"strings"
	// "fmt"
)

func routeUserIds(c *redis.Client) []string {
	//  res, err := c.Keys("*baidu*").Result()
	// res, err := c.Keys("*" + ad_channel + "*").Result()
	res, err := c.Keys("*").Result()
	checkErr(err)

	// fmt.Println(res)
	spl_sli := make([]string, 0)
	// spl_sli := new([]string)
	user_id_sli := make([]string, 0)
	//ip_sli := new([]string)

	get_start_str := make([]string, 0)
	for _, v := range res {
			//string like: li9t209jm7mc6m4vmn88o5a7j0:1454035403.8093:10.10.10.29:baidu:0
			spl_sli = strings.Split(v, ":")
			get_start_str = strings.Split(spl_sli[1], ".")
			// fmt.Println("get_start_str ", get_start_str)
			//format to int 32 type
			// get_start, _ := strconv.ParseInt(get_start_str[0], 10, 32)

			// fmt.Println("get_start ", get_start)
			user_id_sli = append(user_id_sli, get_start_str[0]+":"+spl_sli[4]+":"+spl_sli[0]+":"+spl_sli[3])
	}

	//fmt.Println(len(ip_sli))
	// removeDuplicate(&ip_sli)
	// fmt.Println(user_id_sli)
	return user_id_sli
}
