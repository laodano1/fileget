package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"net/url"
	"flag"
	"errors"
	"path"
	"log"
	"path/filepath"
	"math/rand"
	"time"
	"bufio"
)

type Items interface {
	getIt() []string
}

type FileList struct {
	nameList []string
}

type FileItem struct {
	name string
}

func (fl FileList) getIt() []string {
	return fl.nameList
}

func (fi FileItem) getIt() []string {
	return []string{fi.name}
}

var dlFilePaths FileList

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	//fmt.Printf("\r%s", strings.Repeat(" ", 35))
	//fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
	fmt.Printf("\rDownloading... %dk complete ", wc.Total/1024)
}

func getUserAgent() string {
	userAgent := []string{
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
		"Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
		"Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)",
		"Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0",
		"Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5",
		"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52",
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(userAgent))
	return userAgent[index]
}

func DoHTTPRequest(urlStr string, useProxy bool, proxyLink string) (*http.Response, error) {

	var proxyStr string

	if useProxy {
		if proxyLink != "" {
			proxyStr = proxyLink
			fmt.Println(fmt.Sprintf("Customized Proxy: \"%s\"", proxyStr))
		} else {
			proxyStr = "http://87.254.212.120:8080"
			fmt.Println(fmt.Sprintf("Use Default Proxy!!!"))
		}
	} else {
		fmt.Println(fmt.Sprintf("No Proxy!!!"))
	}

	proxyURL, err := url.Parse(proxyStr)
	errChk(err)

	//urlStr := "https://www.baidu.com"
	//urlStr := "http://www.dytt8.net"
	urlLink, err := url.Parse(urlStr)
	errChk(err)

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{
		Transport: transport,
	}

	request, err := http.NewRequest("GET", urlLink.String(), nil)
	errChk(err)

	request.Header.Add("User-Agent", getUserAgent())

	resp, err := client.Do(request)

	return resp, err
}

func errChk(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func DownloadFile(files Items, useProxy bool, proxyLink string) error {
	for index, urlStr := range files.getIt() {
		if urlStr == "" {
			return errors.New("URL is empty! Please use 'l' option to specify url! ")
		}
		dlName := path.Base(urlStr)

		var fileName string
		if len(dlName) > 100 {
			fileName = "DownloadedFile." + string(index)
		} else {
			fileName = dlName
		}

		aPath, err := filepath.Abs(".")
		errChk(err)

		fmt.Println("Saved File At: " + filepath.Join(aPath + string(filepath.Separator) + fileName))

		resp, err := DoHTTPRequest(urlStr, useProxy, proxyLink)
		errChk(err)


		fmt.Println(fmt.Sprintf("StatusCode: %d \nContent-Length: %s", resp.StatusCode, resp.Header.Get("Content-Length")))

		out, err := os.Create(fileName + ".tmp")
		errChk(err)

		//resp, err := http.Get(url)
		//errChk(err)

		counter := &WriteCounter{}
		_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
		errChk(err)

		resp.Body.Close()
		fmt.Print("\n")
		out.Close()

		err = os.Rename(fileName + ".tmp", fileName)
		errChk(err)
	}

	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile )
	//fmt.Println(os.Args[1:])
	link       := flag.String("l", "", "The Link to Download. For Single File Download")
	//name       := flag.String("n", "DownloadFile", "The Downloaded File Name. ")
	useProxy   := flag.Bool("p", true, "Use Proxy or not")
	proxyLink  := flag.String("pl", "", "Proxy Link")
	mMode      := flag.Bool("mm", false, "Multi-file Download Mode. Read from local \"dl.txt\" ")
	flag.Parse()

	if *mMode {
		fileList := "dl.txt"
		_, err := os.Stat(fileList)
		//fmt.Println(err)
		if err == nil {
			fmt.Println("Download Files:")
			file, err := os.Open(fileList)
			errChk(err)
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				fmt.Println("\t" + scanner.Text())
				dlFilePaths.nameList = append(dlFilePaths.nameList, scanner.Text())
			}
			//fmt.Println(fmt.Sprintf("%d files will be downloaded!", len(dlFilePaths.nameList)))
		} else {
			fmt.Println("Download list file does not exist!")
			log.Fatal("Please use 'l' option to specify url")
		}

	}

	//fmt.Println(*useProxy)

	//flag.Usage()
	fmt.Println("Download Started")

	//fileUrl := "https://upload.wikimedia.org/wikipedia/commons/d/d6/Wp-w4-big.jpg"
	var fileUrl FileItem
	fileUrl.name = *link
	//err := DownloadFile("avatar.jpg", fileUrl)
	var err error
	if len(dlFilePaths.nameList) > 0 {
		//fmt.Println("in 1")
		err = DownloadFile(dlFilePaths, *useProxy, *proxyLink)
	} else {
		//fmt.Println("in 2")
		err = DownloadFile(fileUrl, *useProxy, *proxyLink)
	}
	errChk(err)
}



