# Concurrent Web Crawler

The crawler concurrently crawls the web printing links to standard output using breadth-first traversal.

### Compile
Compile the program using command below.
```sh
go build web_crawler
```

### Usage
The program accepts 1 flag parameter "depth" and arbitrary number of links to start crawling from.
The flag is not required. Its default value is 1.

#### Example
```sh
web_crawler -depth=3 https://google.com
```
You should see the following output in your terminal.
> https://google.com
> 
> https://www.google.com/imghp?hl=kk&tab=wi
> 
> https://maps.google.kz/maps?hl=kk&tab=wl
> 
> https://play.google.com/?hl=kk&tab=w8
> 
> ...
