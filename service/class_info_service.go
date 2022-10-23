package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wx_cloud/entity"

	"wx_cloud/parse"
	"wx_cloud/resposity/wx_cloud"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./index.html")
	if err != nil {
		_, _ = fmt.Fprint(w, "内部错误")
		return
	}
	_, _ = fmt.Fprint(w, string(data))
}

// ParseExcelFile 接收并解析excel文件
func ParseExcelFile(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fileName := r.FormValue("filename")

	// 解析excel
	tclass, tclassInfo, err := parse.ParseExcelFile(fileName)
	if err != nil {
		fmt.Printf("InsertClass fail, err: %v\n", err)
		return
	}

	// 记录课堂信息
	if err := wx_cloud.InsertClass(tclass); err != nil {
		fmt.Printf("InsertClass fail, err: %v\n", err)
		return
	}

	// 记录题目列表
	if err := wx_cloud.InsertClassInfo(tclassInfo); err != nil {
		fmt.Printf("InsertClass fail, err: %v\n", err)
		return
	}
}

// QueryClassInfo 根据classId获取课堂题目
func QueryClassInfo(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	classID := r.FormValue("classId")
	fmt.Println("classId: ", classID)

	tclassInfo, err := wx_cloud.QueryClassInfo(classID)
	if err != nil {
		fmt.Printf("QueryClassInfo fail, err: %v\n", err)
		return
	}

	rsp, _ := json.Marshal(tclassInfo)
	_, _ = w.Write(rsp)
}

// RecordAnswer 记录错题
func RecordAnswer(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	classID := r.FormValue("classId")
	openID := r.FormValue("openId")
	wrongIDs := r.FormValue("wrongIds")
	wrongInfo := &entity.TWrongInfo{
		ClassID: classID,
		OpenID: openID,
		WrongIDs: wrongIDs,
	}

	if err := wx_cloud.DeleteWrongInfo(classID, openID); err != nil {
		fmt.Printf("RecordAnswer-DeleteWrongInfo fail, err: %v\n", err)
		return
	}
	if err := wx_cloud.InsertWrongInfo(wrongInfo); err != nil {
		fmt.Printf("RecordAnswer fail, err: %v\n", err)
		return
	}
}

// QueryClassWrongInfo 根据classId和openId获取课堂的错题
func QueryClassWrongInfo(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	classID := r.FormValue("classId")
	openID := r.FormValue("openId")

	twrongInfo, err := wx_cloud.QueryClassWrongInfo(classID, openID)
	if err != nil {
		fmt.Printf("QueryClassWrongInfo fail, err: %v\n", err)
		return
	}

	rsp, _ := json.Marshal(twrongInfo)
	_, _ = w.Write(rsp)
}
