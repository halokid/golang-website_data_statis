package main

import (
	"gopkg.in/redis.v3"
	"strings"
	//"fmt"
	"strconv"
)

/**
*@author: JJyy
*@todo: get the uv number for some ad channel
*@param: redis client, channel_name
*
**/
func uv(c *redis.Client, ad_channel string, start_time int64, end_time int64) int {

	res, err := c.Keys("*" + ad_channel + "*").Result()
	checkErr(err)

	spl_sli := make([]string, 0)
	uv_sli := make([]string, 0)

	tmp_time_sli := make([]string, 0) //for timeset

	for _, v := range res {
		spl_sli = strings.Split(v, ":")
		if start_time == 0 && end_time == 0 {
			uv_sli = append(uv_sli, spl_sli[0])
		} else if start_time != 0 && end_time != 0 {

			//get timeset
			tmp_time_sli = strings.Split(spl_sli[1], ".")
			//format to int 32 type
			get_start, _ := strconv.ParseInt(tmp_time_sli[0], 10, 32)
			if get_start >= start_time && get_start <= end_time {
				uv_sli = append(uv_sli, spl_sli[0])
			}
		}
	}

	return len(uv_sli)
}
