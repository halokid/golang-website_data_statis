# what is this?

this is some tools for website data statis use golang.
like ip, pv, uv, web browser client url info clooect...

we can do much statis use the data if what we want :), this tools just
some example for statis the data.


# how use?   cxcxcx
the RAW DATA KEY like:
>li9t209jm7mc6m4vmn88o5a7j0:1433138813.8056:55.55.55.22:yahoo:0
>li9t209jm7mc6m4vmn88o5a7j0:1433138813.8056:55.55.55.22:google:0
>li9t209jm7mc6m4vmn88o5a7j0:1433138813.8056:55.55.55.22:baidu:0

this porject we store in redis

you can put them into to redis or some other nosql.
the KEY is zset type, the values of the KEY just like
>shell> ZRANGE li9t27mskklc769vmn88o5appp:1433138814.0488:10.10.10.29:360:0 0 -1
><br />"www.xxxxx.com/help/999"
><br />"www.xxxxx.com"
><br />"www.xxxxx.com/member/xxx"
><br />"www.xxxxx./com/kkki"
><br />"www.xxxxxx..com/goods/list/1020"
><br />"www.xxxxx.com/yyyy"



these values is record the url view


go run test:
>go run main.go ip.go pv.go uv.go -ip yahoo 1454035404(start time) 1454035490(end time)

if no time choose, just make the start_time=0,  end_time=0

-ip is for ip data statis, you can change to pv, uv...
yahoo is for the ad channel, you can change to google, baidu or something you have put into RAW DATA
