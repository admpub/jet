package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type Replace struct {
	Old      string
	New      string
	Regexp   *regexp.Regexp
	FileRule *regexp.Regexp
}

var replaces = []*Replace{
	&Replace{`"github.com/CloudyKit/jet"`, `"github.com/admpub/jet"`, nil, nil},
	&Replace{`"github.com/CloudyKit/jet/`, `"github.com/admpub/jet/`, nil, nil},
}

func main() {
	root := filepath.Join(os.Getenv(`GOPATH`), `src`, `github.com/CloudyKit/jet`)
	save := filepath.Join(os.Getenv(`GOPATH`), `src`, `github.com/admpub/jet`)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() == `_tools` || strings.HasPrefix(info.Name(), `.`) {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasPrefix(info.Name(), `.`) {
			return nil
		}
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		content := string(b)
		for _, re := range replaces {
			if re.FileRule != nil {
				if re.FileRule.MatchString(path) == false {
					continue
				}
				//panic(path)
			}
			if re.Regexp == nil {
				content = strings.Replace(content, re.Old, re.New, -1)
			} else {
				fmt.Printf("%#v\n", re.Regexp.FindAllString(content, -1))
				content = re.Regexp.ReplaceAllString(content, re.New)
			}
		}
		saveAs := strings.TrimPrefix(path, root)
		saveAs = filepath.Join(save, saveAs)
		err = os.MkdirAll(filepath.Dir(saveAs), os.ModePerm)
		if err == nil {
			file, err := os.Create(saveAs)
			if err == nil {
				_, err = file.WriteString(content)
			}
		}
		if err != nil {
			return err
		}
		fmt.Println(`Autofix ` + path + `.`)
		return nil
	})
	defer time.Sleep(5 * time.Minute)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(`Autofix complete.`)
}
