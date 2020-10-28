package main

import (
	"flag"
	"fmt"
	"github.com/LindsayBradford/go-dbf/godbf"
	"log"
	"sort"
)

func main() {
	var filePath string
	var fieldName string
	var rowNumber int

	flag.StringVar(&filePath, "f", "example.dbf", "需要统计的dbf 文件信息")
	flag.StringVar(&fieldName, "k", "", "按照该字段统计的列名称")
	flag.IntVar(&rowNumber, "l", 10, "展示的行数")
	flag.PrintDefaults()
	flag.Parse()

	dbfTable, err := godbf.NewFromFile(filePath, "UTF8")
	if dbfTable == nil {
		log.Fatalln("找不到文件", filePath)
	}
	if err != nil {
		log.Fatalln("读取dbf文件失败", err)
	}
	statisticTimes := make(map[string]int)
	for i := 0; i < dbfTable.NumberOfRecords(); i++ {
		res, err := dbfTable.FieldValueByName(i, fieldName)
		if err != nil {
			log.Fatalln(err)
		}
		statisticTimes[res] = statisticTimes[res] + 1
	}
	for ix, pair := range rankByWordCount(statisticTimes) {
		if ix > rowNumber {
			break
		}
		fmt.Printf("%4v. %v\n", ix, pair)
	}
}
func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	key   string
	value int
}

func (p Pair) String() string { return fmt.Sprintf(" %-6v -> %-9v ", p.key, p.value) }

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].value < p[j].value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
