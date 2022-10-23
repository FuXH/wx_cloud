package parse

import (
	"wx_cloud/entity"
)

// ParseStudentLevel 解析学生作业等级
func ParseStudentLevel(fileName string) ([]*entity.TStudentLevel, error) {
	datas, err := readExcel(fileName)
	if err != nil {
		return nil, err
	}

	tstudentLevel := convertToStudentLevel(datas)

	return tstudentLevel, nil
}

func convertToStudentLevel(datas []string) []*entity.TStudentLevel {
	studentLevel := make([]*entity.TStudentLevel, 0)
	var school, grade, class string
	for i := 0; i < len(datas); i++ {
		if datas[i] == "学校" {
			school = datas[i+1]
			i += 1
		} else if datas[i] == "年级" {
			grade = datas[i+1]
			i += 1
		} else if datas[i] == "班级" {
			class = datas[i+1]
			i += 1
		} else if datas[i] == "等级1" {
			for j := 1; (i + j) < len(datas); j++ {
				if datas[i+j] == "等级1" || datas[i+j] == "等级2" ||
					datas[i+j] == "等级3" || datas[i+j] == "" {
					i += j
					break
				}
				studentLevel = append(studentLevel, &entity.TStudentLevel{
					School:  school,
					Grade:   grade,
					Class:   class,
					Level:   1,
					Student: datas[i+j],
				})
			}
		} else if datas[i] == "等级2" {
			for j := 1; (i + j) < len(datas); j++ {
				if datas[i+j] == "等级1" || datas[i+j] == "等级2" ||
					datas[i+j] == "等级3" || datas[i+j] == "" {
					i += j
					break
				}
				studentLevel = append(studentLevel, &entity.TStudentLevel{
					School:  school,
					Grade:   grade,
					Class:   class,
					Level:   2,
					Student: datas[i+j],
				})
			}
		} else if datas[i] == "等级3" {
			for j := 1; (i + j) < len(datas); j++ {
				if datas[i+j] == "等级1" || datas[i+j] == "等级2" ||
					datas[i+j] == "等级3" || datas[i+j] == "" {
					i += j
					break
				}
				studentLevel = append(studentLevel, &entity.TStudentLevel{
					School:  school,
					Grade:   grade,
					Class:   class,
					Level:   3,
					Student: datas[i+j],
				})
			}
		}
	}
	return studentLevel
}
