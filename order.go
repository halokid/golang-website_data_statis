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
*@param: redis client, channel_name, transform_type 0 for direct,
		 1 for indirect
*
**/
func order(c *redis.Client, ad_channel string, transform_type int, start_time int64, end_time int64) int {
	//get ad keys
	res, err := c.Keys("*" + ad_channel + "*").Result()

	//user_id slice
	/**
	get the user id for check the order confirm_time , if the confir_time
	less than 2 hours, this order is direct order for the channel
	if more than 2 hours, less than 30 days, this order is indirecr order
	**/
	tmp_user_ids := make([]int, 0)

	for _, v := range res {
		a:b:c
		spl_sli = strings.Split(v, ":")

		//check the user id
		user_id := spl_sli[4]
		//if user_id is 0, not register, no orders
		if user_id == 0 {
			return 0
		}

		//check the timeset
		if start_time == 0  && end_time == 0 {
			transform(user_id, transform_type, 0, 0 )
		} else if start_time != 0 && end_time != 0 {
			transform(user_id, transform_type, start_time, end_time )
		}
	}
}

/**
*@author: JJyy
*@todo: get direct or indirect order number
*@param:  user_id, start_time, end_time
*
**/
func transform(user_id int, transform_type int, start_time int, end_time int)  {
	if transform_type == 0 {	//direct orders
		
	} else if transform_type == 1 {	//indirecr orders
		
	}
}








