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
	"github.com/sssvip/qrterminal"
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
	fileNames := listDir(currentPath)
	for _, filename := range fileNames {
		fmt.Fprintf(w, "<a href=\"%s%s\">%s</a></br>", staticPath, string(filename), string(filename))
	}
	fmt.Fprintf(w, "</br></br></br>See the new version of SimpleFileServer :<a href=\"%s\">%s</a>", repoUrl, repoUrl)
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
	var defaultPort = "80"
	port := defaultPort
	showQRCode := false
	const qr = "qr"
	showQRCodeIpFilter := qr
	if len(os.Args) > 1 {
		tmpPort, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Println("http port must a num, eg: simplefileserver.exe 8080")
			log.Println("use defualut http port:" + port)
		} else {
			port = strconv.Itoa(tmpPort)
		}
		//check qr
		for _, arg := range os.Args {
			if strings.Contains(arg, qr) {
				showQRCode = true
				showQRCodeIpFilter = arg
				break
			}
		}
	}

	http.HandleFunc("/", home)
	http.HandleFunc(staticPath, func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, "=>", r.URL.Path)
		http.ServeFile(w, r, r.URL.Path[len(staticPath):])
	})
	var startEcho = "started file server http://127.0.0.1"
	if port != defaultPort {
		startEcho = fmt.Sprintf("%s:%s", startEcho, port)
	}
	log.Println(startEcho)
	var optPreAddr = "Optional addr: "
	for _, localIp := range getLocalIPs() {
		urlAddr := fmt.Sprintf("http://%s", localIp)
		if port != defaultPort {
			urlAddr = fmt.Sprintf("%s:%s", urlAddr, port)
		}
		log.Println(fmt.Sprintf("%s%s", optPreAddr, urlAddr))
		if showQRCode {
			if showQRCodeIpFilter != qr {
				if strings.Contains(urlAddr, strings.Replace(showQRCodeIpFilter, qr, "", -1)) {
					qrterminal.Generate(urlAddr, qrterminal.L, os.Stdout)
					fmt.Println()
					break
				}
			} else {
				qrterminal.Generate(urlAddr, qrterminal.L, os.Stdout)
				fmt.Println()
			}
		}

	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	entry()
}
