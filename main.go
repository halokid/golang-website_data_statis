/*
how to use?
sample:
if you want to get the ip numbwer data for baidu channel, just run the command like that  
#shell:  ./statis_data -ip baidu

*/

package main

import (
	"os"
    "gopkg.in/redis.v3"
	"fmt"
//    "strings"
//    "reflect"
)


/**
*@author: JJyy
*@todo:  return the redis client handle
*@param: 
*
**/
func rds() *redis.Client {
    client := redis.NewClient(&redis.Options {
        Addr:    "192.168.0.139:6379",
        Password:    "",
        DB:    1,
        
    })
    
    return client
}


/**
*@author: JJyy
*@todo: check error
*@param: 
**/
func checkErr( err error ) {
    if err !=nil {
        panic(err)
    }
}


/**
*@author: JJyy
*@todo: del repeative element for slice
*@param: slice point
*
**/
func removeDuplicate(slis *[]string) {
  found := make(map[string]bool)
    j := 0
    for i,val := range *slis {
      if _,ok := found[val]; !ok {
          found[val] = true
          (*slis)[j] = (*slis)[i]
          j++
      }
    }
  *slis = (*slis)[:j]
}



func main() {
    //get the data type
    var data_type string
    data_type = os.Args[1]
    
    var channel_name string
    channel_name = os.Args[2]
//    fmt.Println(data_type)
    
    c := rds()    //redis client struct
    
    if data_type == "-ip" {    //get the ip data
        num := ip(c, channel_name)
        fmt.Println(num)
    } else if data_type == "-pv" {
        num := pv(c, channel_name)
        fmt.Println(num)
    } else if data_type == "-uv" {
        num := uv(c, channel_name)
        fmt.Println(num)
    }
    
    
    fmt.Println("main process")
}






















