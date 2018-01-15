package main

import (
	"fmt"
	"path/filepath"
	"os"
	"log"
	"strings"
	"io/ioutil"
	"net/http"
	"strconv"
	"net"
)

var staticPath = "/static/"
var repoUrl = "https://github.com/sssvip/simple-file-server"

func listDir(dirPth string) (a []string) {
	dir, err := ioutil.ReadDir(dirPth)
	var files = []string{}
	if err != nil {
		return nil
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		files = append(files, PthSep+fi.Name())
	}
	return files
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
func home(w http.ResponseWriter, r *http.Request) {
	currentPath := getCurrentDirectory()
	fmt.Fprintf(w, "<h1>Files in [%s]</h1></br></br></br>", currentPath)
	filenames := listDir(currentPath)
	for _, filename := range filenames {
		fmt.Fprintf(w, "<a href=\"%s%s\">%s</a></br>", staticPath, string(filename), string(filename))
	}
	fmt.Fprintf(w, "</br></br></br>See the new version of SimpleFileServer :<a href=\"%s\">%s</a>", repoUrl, repoUrl)
	//if len(filenames)<1{
	//	fmt.Fprintln(w,"no file in this directory")
	//}
}

func getLocalIPs() []string {
	result := []string{"localhost"}
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		log.Println("Get local IP addr failed!!!")
		return result
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				result = append(result, ipnet.IP.String())
			}
		}
	}
	return result
}

func entry() {
	port := "80"
	if len(os.Args) > 1 {
		tmpPort, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Println("http port must a num, eg: simplefileserver.exe 80")
			log.Println("use defualut http port:" + port)
		} else {
			port = strconv.Itoa(tmpPort)
		}
	}
	http.HandleFunc("/", home)
	http.HandleFunc(staticPath, func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		http.ServeFile(w, r, r.URL.Path[len(staticPath):])
	})
	log.Println("starttd file server http://127.0.0.1:" + port)
	for _, localIp := range getLocalIPs() {
		log.Println("Optional addr: http://" + localIp + ":" + port)
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	entry()
}
