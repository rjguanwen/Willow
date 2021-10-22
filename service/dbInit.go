// @Title  dbinit
// @Description  标签数据库初始化
// @Author  zhengbin  2021/9/13 14:54
package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/rjguanwen/willow/public"
	"io"
	log "github.com/cihub/seelog"
	"os"
	"strings"
	"time"
)

var TagsDb *memdb.MemDB
var tagsFileName string
func init(){
	tagsFileName = "data/tags-json-sample.txt"
	log.Debugf("init() - 数据文件路径：%s\n", tagsFileName)
}

// 设置数据文件路径
func SetTagsFile(newTagsFileName string) {
	tagsFileName = newTagsFileName
}

// @title    InitDB
// @description   初始化数据库
// @auth      zhengbin     2021/9/13 14:54
func InitDB() *memdb.MemDB{
	log.Info("dbInit - InitDB() - begin ...")
	startTime := time.Now()
	// 获取Schema
	tagsSchema := public.GetTagsSchema()
	// Create a new data base
	db, err := memdb.NewMemDB(tagsSchema)
	if err != nil {
		panic(err)
	}

	initOverChan := make(chan bool, 1)

	// 建立一个标签数据管道，方便并行读文件与写数据库
	tagsRecordChan := make(chan *public.Tags, 1000)
	//// 暂时将文件位置写死
	//tagsFileName := "../data/tags-json.txt"
	log.Debugf("InitDB - 数据文件路径：%s\n", tagsFileName)
	// 启动协程，从文件中读取标签记录
	go readTagsFromFile(tagsFileName, tagsRecordChan)
	// 启动协程，将标签插入数据库
	go insertTagsIntoDbFromChan(tagsRecordChan, db, initOverChan)
	// 获取完成标志
	<- initOverChan
	duration := time.Since(startTime)
	log.Info("用时：", duration)
	log.Info("dbInit - InitDB() - end!")
	return db
}

// 从文件中按行读取标签，将读取到的标签放入管道
// tagsFileName: 标签数据文件路径
// trChan: 标签对象管道
func readTagsFromFile(tagsFileName string, trChan chan *public.Tags){
	log.Info("开始读取标签文件...")
	file, err := os.OpenFile(tagsFileName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("文件打开错误！", err)
		return
	}
	defer file.Close()

	// 记录数
	tagsNum := 0
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("读取文件错误！", err)
			}
		}
		// 将读取到的行数据转换为结构放入管道
		var tag public.Tags
		json.Unmarshal([]byte(line), &tag)
		trChan <- &tag
		tagsNum++
	}
	close(trChan)
	log.Infof("标签文件读取结束，共读取到标签 %d 条！", tagsNum)
}

// 从管道中获取标签并存入数据库
// trChan: 标签数据管道
// db: 标签内存数据库
func insertTagsIntoDbFromChan(trChan chan *public.Tags, db *memdb.MemDB, overChan chan bool){
	// 创建一个写入事务
	txn := db.Txn(true)
	recordNum := 0
	for {
		tag, ok := <- trChan
		if !ok {
			log.Infof("管道已经关闭，插入数据 %d 条！", recordNum)
			break
		}
		err := txn.Insert("memberTags", tag)
		recordNum++
		if err != nil {
			log.Error("数据插入错误！", err)
			panic(err)
		}
	}
	txn.Commit()
	overChan <- true
}