package parse

import (
	"fmt"
	"strings"

	"wx_cloud/entity"

	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

func ParseClassFile(fileName string) (*entity.TClass, *entity.TClassInfo, error) {
	datas, err := readExcel(fileName)
	if err != nil {
		return nil, nil, err
	}

	classID := genClassID()
	tclass, tclassInfo := convertToClassInfo(classID, datas)

	return tclass, tclassInfo, nil
}

func readExcel(fileName string) ([]string, error) {
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Printf("readExcel-OpenFile fail, err: %v", err)
		return nil, err
	}
	defer file.Close()

	// 获取 Sheet1. 子表格的数据
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		fmt.Printf("readExcel-GetRows fail, err: %v\n", err)
		return nil, err
	}

	res := make([]string, len(rows))
	for index, val := range rows {
		if len(val) == 0 {
			continue
		}
		res[index] = val[0]
	}
	return res, nil
}

func convertToClassInfo(classID string, datas []string) (*entity.TClass, *entity.TClassInfo) {
	tclass := &entity.TClass{
		ClassID: classID,
	}
	tclassInfo := &entity.TClassInfo{
		ClassID:   classID,
		Questions: make([]*entity.QuestionInfo, 0),
	}
	questionID := 0
	for i := 0; i < len(datas); i++ {
		if datas[i] == "学校" {
			tclass.School = datas[i+1]
			i += 1
		} else if datas[i] == "年级" {
			tclass.Grade = datas[i+1]
			i += 1
		} else if datas[i] == "班级" {
			tclass.Class = datas[i+1]
			i += 1
		} else if datas[i] == "老师" {
			tclass.Teacher = datas[i+1]
			i += 1
		} else if datas[i] == "名称" {
			tclass.ClassName = datas[i+1]
			i += 1
		} else if datas[i] == "题目" {
			tmp := &entity.QuestionInfo{
				QuestionID: questionID,
				Subject:    datas[i+1],
			}
			for j := 2; (i + j) < len(datas); j++ {
				if datas[i+j] == "答案" {
					tmp.Right = datas[i+j+1]
					i = i + j + 1
					break
				}
				option := strings.Split(datas[i+j], "：")
				tmp.Options = append(tmp.Options, &entity.OptionInfo{
					Value: option[0],
					Name:  option[1],
				})
			}
			tclassInfo.Questions = append(tclassInfo.Questions, tmp)
			questionID += 1
		}
	}
	return tclass, tclassInfo
}

func genClassID() string {
	return uuid.New().String()
}
