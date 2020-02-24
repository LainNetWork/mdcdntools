package processor

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mdcdntools/common"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type ResourceLink struct {
	origin   string
	tag      string
	url      string
	filename string
	filepath string
}

func downloadImgToLocal(path string) (bool, error) {
	links := getAllImgLinkInMdFile(path)
	srcFile, err := os.Open(path)
	if err != nil {
		return false, errors.New("file not exit")
	}
	dirPath := strings.TrimSuffix(path, filepath.Ext(path)) + "/"
	os.Mkdir(dirPath, os.ModePerm)
	destFile, _ := os.Create(dirPath + filepath.Base(path))
	defer destFile.Close()
	defer srcFile.Close()
	content, err := ioutil.ReadAll(srcFile)
	if err != nil {
		return false, errors.New("faild to read markdown file")
	}
	contentString := string(content)
	if links != nil {
		//创建资源文件夹

		for i, link := range links {
			imgPath := dirPath
			ext := filepath.Ext(link.url)
			filename := strconv.Itoa(i) + ext
			imgPath += filename
			link.filename = filename
			link.filepath = imgPath
			downloadImg(link.filepath, link.url, common.Config.Refer)
			contentString = strings.Replace(contentString, link.origin, fmt.Sprintf("![%s](%s)", link.tag, filename), -1)
		}

	}
	ioutil.WriteFile(dirPath+filepath.Base(path), []byte(contentString), os.ModePerm)
	return true, nil
}

var client = new(http.Client)

func downloadImg(path string, url string, referer string) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("referer", referer)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile(path, body, os.ModePerm)
}

func getAllImgLinkInMdFile(file string) []ResourceLink {
	mdFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("failed load md file", err)
	}
	article := string(mdFile)
	reg, err := regexp.Compile("!\\[\\S*\\]\\(\\S*\\)")
	//links := make([]string,5)
	urls := reg.FindAllString(article, -1)
	urlList := make([]ResourceLink, 0)
	for i := range urls {
		url := urls[i]
		temp := strings.Split(url, "](")
		link := new(ResourceLink)
		link.origin = url
		link.url = strings.TrimSuffix(temp[1], ")")
		link.tag = strings.TrimPrefix(temp[0], "![")
		urlList = append(urlList, *link)
	}
	return urlList
}

func Execute(config common.ArgsConfig) {
	filepath.Walk(config.Path, walk)
}

func walk(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return err
	}
	//取文件后缀名
	if strings.ToLower(filepath.Ext(path)) == ".md" {
		result, err := downloadImgToLocal(path)
		if result == false {
			log.Println(err.Error())
		}
	}
	return err
}
