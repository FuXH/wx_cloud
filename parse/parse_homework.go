package parse

import (
	"fmt"
	"strings"
	"wx_cloud/entity"
)

// ParseHomeworkFile 解析布置的作业
func ParseHomeworkFile(fileName string) (*entity.THomework, error) {
	datas, err := readExcel(fileName)
	if err != nil {
		return nil, err
	}

	thomework := convertToHomework(datas)
	fmt.Println(thomework)
	for _, val := range thomework.Datas {
		fmt.Println(val)
	}

	return thomework, nil
}

func convertToHomework(datas []string) *entity.THomework {
	homework := &entity.THomework{}
	for i := 0; i < len(datas); i++ {
		if datas[i] == "学校" {
			homework.School = datas[i+1]
			i += 1
		} else if datas[i] == "年级" {
			homework.Grade = datas[i+1]
			i += 1
		} else if datas[i] == "班级" {
			homework.Class = datas[i+1]
			i += 1
		} else if datas[i] == "日期" {
			homework.Date = datas[i+1]
			i += 1
		} else if datas[i] == "作业" {
			for j := 1; (i + j) < len(datas); j++ {
				if datas[i+j] == "等级1" {
					tmp := &entity.HomeInfo{
						Level: 1,
					}
					content := []string{}
					for z := 1; (i + j + z) < len(datas); z++ {
						if datas[i+j+z] == "等级1" || datas[i+j+z] == "等级2" ||
							datas[i+j+z] == "等级3" || datas[i+j+z] == "" {
							j = j + z - 1
							break
						}
						content = append(content, datas[i+j+z])
					}
					tmp.Content = strings.Join(content, ",")
					homework.Datas = append(homework.Datas, tmp)

				} else if datas[i+j] == "等级2" {
					tmp := &entity.HomeInfo{
						Level: 2,
					}
					content := []string{}
					for z := 1; (i + j + z) < len(datas); z++ {
						if datas[i+j+z] == "等级1" || datas[i+j+z] == "等级2" ||
							datas[i+j+z] == "等级3" || datas[i+j+z] == "" {
							j = j + z - 1
							break
						}
						content = append(content, datas[i+j+z])
					}
					tmp.Content = strings.Join(content, ",")
					homework.Datas = append(homework.Datas, tmp)

				} else if datas[i+j] == "等级3" {
					tmp := &entity.HomeInfo{
						Level: 3,
					}
					content := []string{}
					for z := 1; (i + j + z) < len(datas); z++ {
						if datas[i+j+z] == "等级1" || datas[i+j+z] == "等级2" ||
							datas[i+j+z] == "等级3" || datas[i+j+z] == "" {
							j = j + z - 1
							break
						}
						content = append(content, datas[i+j+z])
					}
					tmp.Content = strings.Join(content, ",")
					homework.Datas = append(homework.Datas, tmp)
				}
			}
		}
	}
	return homework
}
