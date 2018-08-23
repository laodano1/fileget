# fileget
This is a golang competency build project. http/goroutine/bufio/os... will be involved.

Have implemented:
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
  
