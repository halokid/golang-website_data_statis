package main

import (
    "gopkg.in/redis.v3"
	"fmt"
//    "strings"
    "reflect"
)


func ExampleNewClient() {
    client := redis.NewClient(&redis.Options {
        Addr:    "172.16.2.29:6379",
        Password:    "",
        DB:    0,
        
    })
    
    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    
}




func main() {
//    ExampleNewClient()
      client := redis.NewClient(&redis.Options {
        Addr:    "172.16.2.29:6379",
        Password:    "",
        DB:    1,
        
    })
    
    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    
    s, err := client.Get("a").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(s)
    
    z, err := client.ZRange("li9t209jm7mc6m4vmn88o5a7j0:1454035404.5141:87.56.45.59:yahoo:0", 0, -1).Result()
    
    if err != nil{
        panic(err)
    }
    
//    fmt.Print(z)
    fmt.Print( reflect.TypeOf(z) )
    
    
    for k, v := range z {
        fmt.Println( k, v)
    }
    
}






















