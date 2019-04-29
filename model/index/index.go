package index

import (
	"cat-slave/model"
	"cat-slave/model/passage"
	"cat-slave/pkg/file"
	"github.com/go-ego/riot/types"
	"log"
	"strconv"
)



func TotalIndex(keyword string) interface{} {
	result := model.DB.Roit.Search(types.SearchReq{Text:keyword})
	passageList := []passage.Passage{}
	for _, value := range result.Docs.(types.ScoredDocs) {
		//fmt.Println(value.ScoredID.DocId)
		id, err := strconv.Atoi(value.ScoredID.DocId)
		if err != nil {
			log.Panic(err)
		}
		content := value.Content
		//
		////rs := rune(keyword)
		//index := strings.Index(value.Content, keyword)
		//
		//var start int
		//var end int
		//if index - 10 < 0 {
		//	start = 0
		//} else {
		//	start = index - 10
		//}
		//if index + 10 > len(content) {
		//	end = len(content)
		//} else {
		//	end = index + 10
		//}
		//
		//content = file.Substring(content, start, end)
		passageList = append(passageList, passage.Passage{
			Title: file.FileList[id].Title,
			Content: content,
		})
	}
	return passageList
}