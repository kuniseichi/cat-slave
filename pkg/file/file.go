package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type File struct {
	Title string
	Content string
}

var FileList []File

func walkFunc(path string, info os.FileInfo, err error) error {
	if info == nil {
		// 文件名称超过限定长度等其他问题也会导致info == nil
		// 如果此时return err 就会显示找不到路径，并停止查找。
		println("can't find:(" + path + ")")
		return nil
	}
	println(info.Name())
	if info.IsDir() {
		//println("This is folder:(" + path + ")")
		return nil
	} else {
		if filepath.Ext(path) == ".js" {
			// name处理
			name := strings.Replace(info.Name(), ".js", "", -1)
			// content处理

			FileList = append(FileList, File{
				Title: name,
				Content:ReadFile(path),
			})
		}
		return nil
	}
}

func ShowFileList(root string) {
	err := filepath.Walk(root, walkFunc)
	if err != nil {
		fmt.Printf("filepath.Walk() error: %v\n", err)
	}
	return
}

func ReadFile(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	//str := string(b)
	//re := regexp.MustCompile(fileRe)
	//maches := re.FindAllSubmatch(b, -1)
	//for _, m := range maches{
	//
	//}
	return trimHtml(string(b))
}

func trimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	// 去除回车
	src = strings.Replace(src, "\n", "", -1)

	s1 := strings.Split(src, "return (")
	s2 := s1[1]
	s := strings.Split(s2, ");")
	return strings.TrimSpace(s[0])
}

func Substring(str string, begin, end int) string {
	rs := []rune(str)
	//lth := len(rs)
	//fmt.Printf("begin=%d, end=%d, lth=%d\n", begin, length, lth)
	//if begin < 0 {
	//	begin = 0
	//}
	//if begin >= lth {
	//	begin = lth
	//}
	//end := begin + length

	//if end > lth {
	//	end = lth
	//}
	fmt.Printf("begin=%d, end=%d, lth=%d\n", begin, end, len(rs))
	fmt.Println(string(rs))
	fmt.Println(string(rs[begin:end]))
	//fmt.Println(string(rs[begin*4:end*4]))
	return string(rs[begin:end])
}