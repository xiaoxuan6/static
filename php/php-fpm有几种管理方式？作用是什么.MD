# php-fpm有几种管理方式？作用是什么？

目前最新5.3.x的php-fpm，有两种管理进程的方式，分别是 static 和 dynamic。

参数：

pm.max_children：静态方式下开启的php-fpm进程数量。

pm.start_servers：动态方式下的起始php-fpm进程数量。

pm.min_spare_servers：动态方式下的最小php-fpm进程数量。

pm.max_spare_servers：动态方式下的最大php-fpm进程数量。

### static
    进程数自始至终都是pm.max_children指定的数量，pm.start_servers，pm.min_spare_servers，pm.max_spare_servers配置将没有作用。

### dynamic 
    则进程数是动态的，最开始是pm.start_servers指定的数量，如果请求较多，则会自动增加，但不超过 pm.max_children指定的数量，同时保证空闲的进程数不小于pm.min_spare_servers，如果进程数较多，也会进行相应清理， 保证多余的进程数不多于pm.max_spare_servers。

当php-fpm启动后，一个php-cgi进程约战3M内存，但是当它们处理过一些请求后，有些内存是释放不掉的，占用的内存能达到20M-30M不等。

对于内存比较吃紧，同时并发量不是很大的应用，可以考虑采用static的方式，这样可以很好的控制php-fpm的所消耗的总内存数，让系统更加 平稳运行。另外由于并发量很小，可以适当的把设置pm.max_requests小一些，以便让php-fpm进程有机会重启，从而释放其占用的内存。

# 相关知识点
[php-fpm的静态static和动态dynamic执行方式比较](https://www.cnblogs.com/taijun/p/4206770.htmlhttps://blog.csdn.net/joshua_317/article/details/68946667)


