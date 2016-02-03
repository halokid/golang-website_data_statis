package main

import (
// "gopkg.in/redis.v3"
// "strconv"
// "strings"
// "fmt"
)

/**
*@author: JJyy
*@todo: get the order data for some ad channel
*@param: redis client, channel_name
*
**/
func order(c *redis.Client, ad_channel string, start_time int64, end_time int64) int {
	//get ad keys
	res, err := c.Keys("*" + ad_channel + "*").Result()

	//user_id slice
	tmp_user_ids := make([]int, 0)

	for _, v := range res {
		spl_sli = strings.Split(v, ":")

		if start_time == 0  && end_time == 0 {
			
		}
	}
}
