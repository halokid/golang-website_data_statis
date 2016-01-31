package main

import (
//	"strings"
    "gopkg.in/redis.v3"
//    "fmt"
)


/**
*@author: JJyy
*@todo: get the pv number for some ad channel
*@param: redis client, channel_name
*
**/
func pv( c *redis.Client, ad_channel string ) int {

    res, err := c.Keys("*"+ad_channel+"*").Result()
    checkErr(err)
    
//    spl_sli := make([]string, 0)
//    ip_sli := make([]string, 0)
    
    var page_num int = 0
    
    for _, v := range res {
//        spl_sli = strings.Split(v, ":")
//        ip_sli = append( ip_sli, spl_sli[2] )
        z_pages, err := c.ZRange(v, 0, -1).Result()
        checkErr(err)
        
        for _, v = range z_pages {
            page_num++
        }
        
    }
    
    
    return page_num
}





















