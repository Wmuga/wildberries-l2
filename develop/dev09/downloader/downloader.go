package downloader

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

var (
	reHref = regexp.MustCompile(`href=['"](?P<url>/.*?)['"]`)

	stdPermissons fs.FileMode = 0766
)

func downloadPage(link string) (data []byte, ctype string, err error) {
	resp, err := http.Get(link)
	if err != nil {
		return
	}
	ctype = resp.Header.Get("Content-Type")
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	return
}

func createFileDir(prefix string, link string) (fullname string, err error) {
	fullname = path.Join(prefix, link[1:])
	dir := path.Dir(fullname)
	err = os.MkdirAll(dir, stdPermissons)
	return
}

func saveFile(directory, filename string, data []byte) error {
	filename, err := createFileDir(directory, filename)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, stdPermissons)
}

func DownloadPageToDir(base string, urlPath string, directory string) (links []string, err error) {
	filename := urlPath
	// Если основная страница - добавить index.html
	if urlPath == "/" {
		filename = path.Join(filename, "/index.html")
	}

	data, ctype, err := downloadPage(base + urlPath)
	if err != nil {
		return
	}

	// Добавить .html к страничке если ее нет
	if strings.Contains(ctype, "html") && !strings.HasSuffix(filename, ".html") {
		filename += ".html"
	}
	// Сохранить файл
	err = saveFile(directory, filename, data)
	if err != nil {
		return
	}

	if strings.Contains(ctype, "html") {
		// Поиск всего, что похоже на ссылку в html-ке
		regRes := reHref.FindAllStringSubmatch(string(data), -1)
		res := make([]string, len(regRes))
		for i := range res {
			res[i] = regRes[i][1]
		}
		return res, nil
	}
	// Если статика - то искать нечего
	return nil, nil
}

func DownloadPageToDirRecursive(base string, urlPath string, directory string, timeoutSec, cur, max int) {
	if cur >= max {
		return
	}
	fmt.Println(base, urlPath)
	links, err := DownloadPageToDir(base, urlPath, directory)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for i := range links {
		if strings.HasPrefix(links[i], "http") || links[i] == "" {
			continue
		}
		time.Sleep(time.Second * time.Duration(timeoutSec))
		DownloadPageToDirRecursive(base, links[i], directory, timeoutSec, cur+1, max)
	}
}
