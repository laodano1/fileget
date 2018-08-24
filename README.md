# fileget
This is a golang competency build project. http/goroutine/bufio/os... will be involved.

How to use:
[~]$ go run bigFileDownload.go -h
  -l string
        The Link to Download. For Single File Download
  -mm
        Multi-file Download Mode. Read from local "dl.txt"
  -p    Use Proxy or not (default true)
  -pl string
        Proxy Link

     Example:
     [~]$ go run bigFileDownload.go -l https://upload.wikimedia.org/wikipedia/commons/d/d6/Wp-w4-big.jpg -p=0

     [~]$ go run bigFileDownload.go -mm=1



Have implemented:
---2018-08-23
1. download single file with 'l' option
2. download multiple files with 'dl.txt' list file
3. use random user agent to create the http request
4. use proxy or not 



Plan to do:
1. download files concurrently with goroutine
2. support resuming from break point function, if it is supported by server side.
3. if file size exceeds 50Mb, use multi-goroutine to download it.
     1) lock machanism
     2) wait goroutines
4. ftp protocol support
5. ssh2 protocol support
  

