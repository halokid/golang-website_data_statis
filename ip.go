package main

import (
	"strings"
    "gopkg.in/redis.v3"
//    "fmt"
)


/**
*@author: JJyy
*@todo: get the ip number for some ad channel
*@param: redis client, channel_name
*
**/
func ip( c *redis.Client, ad_channel string ) int {

//    res, err := c.Keys("*baidu*").Result()
    res, err := c.Keys("*"+ad_channel+"*").Result()
    checkErr(err)
    
//    fmt.Println(res)
    spl_sli := make([]string, 0)
//    spl_sli := new([]string)
    ip_sli := make([]string, 0)
//    ip_sli := new([]string)
    for _, v := range res {
        spl_sli = strings.Split(v, ":")
//        ip_sli[k] = spl_sli[2]
        ip_sli = append( ip_sli, spl_sli[2] )
    }
    
//    fmt.Println(len(ip_sli))
    
    removeDuplicate(&ip_sli)
//    fmt.Println(len(ip_sli))
    
    return len(ip_sli)
}





















