package main

import (
	"strings"
    "gopkg.in/redis.v3"
//    "fmt"
)


/**
*@author: JJyy
*@todo: get the uv number for some ad channel
*@param: redis client, channel_name
*
**/
func uv( c *redis.Client, ad_channel string ) int {

    res, err := c.Keys("*"+ad_channel+"*").Result()
    checkErr(err)
    
    spl_sli := make([]string, 0)
    uv_sli := make([]string, 0)
    
    for _, v := range res {
        spl_sli = strings.Split(v, ":")
        uv_sli = append( uv_sli, spl_sli[0] )
    }
    
    
    return len(uv_sli)
}





















