package main

import (
	"gopkg.in/redis.v3"
	"strconv"
	"strings"
	// "fmt"
)

/**
*@author: JJyy
*@todo: get the ip number for some ad channel
*@param: redis client, channel_name
*
**/
func ip(c *redis.Client, ad_channel string, start_time int64, end_time int64) int {

	//    res, err := c.Keys("*baidu*").Result()
	res, err := c.Keys("*" + ad_channel + "*").Result()
	checkErr(err)

	//    fmt.Println(res)
	spl_sli := make([]string, 0)
	//    spl_sli := new([]string)
	ip_sli := make([]string, 0)
	//    ip_sli := new([]string)

	get_start_str := make([]string, 0)
	for _, v := range res {
		//string like: li9t209jm7mc6m4vmn88o5a7j0:1454035403.8093:10.10.10.29:baidu:0
		spl_sli = strings.Split(v, ":")
		//        ip_sli[k] = spl_sli[2]

		//check the timsset
		if start_time == 0 && end_time == 0 {
			ip_sli = append(ip_sli, spl_sli[2])
		} else if start_time != 0 && end_time != 0 {

			// fmt.Println("get_start_sli ", spl_sli[1])
			// get_start, _ := strconv.Atoi(spl_sli[1])
			// get_start, _ := strconv.ParseInt(spl_sli[1], 10, 32)

			// get_start_str := make([]string, 0)
			//timeset is microtime, like 1454035403.8093
			get_start_str = strings.Split(spl_sli[1], ".")
			// fmt.Println("get_start_str ", get_start_str)
			//format to int 32 type
			get_start, _ := strconv.ParseInt(get_start_str[0], 10, 32)

			// fmt.Println("get_start ", get_start)
			if get_start >= start_time && get_start <= end_time {
				ip_sli = append(ip_sli, spl_sli[2])
			}

		}

	}

	//    fmt.Println(len(ip_sli))

	removeDuplicate(&ip_sli)
	//    fmt.Println(len(ip_sli))

	return len(ip_sli)
}
