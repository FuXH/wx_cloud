package wx_cloud

import (
	"encoding/json"
	"fmt"
	"wx_cloud/entity"
)

const (
	// 课堂信息
	classTable = "t_class"
	// 课堂题目信息
	classInfoTable = "t_class_info"
)

// InsertClass 记录课堂信息
func InsertClass(class *entity.TClass) error {
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databaseadd?access_token=%s", GetAccessToken())
	body, _ := json.Marshal(class)
	req := map[string]interface{}{
		"env": "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"%s\").add({data:[%s]})`,
			classTable, string(body)),
	}
	rsp, _ := clientPost(url, req)
	fmt.Printf("InsertClass rsp: %v\n", string(rsp))

	return nil
}

// InsertClassInfo 记录课堂题目
func InsertClassInfo(classInfo *entity.TClassInfo) error {
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databaseadd?access_token=%s", GetAccessToken())
	questions, _ := json.Marshal(classInfo.Questions)
	sql := fmt.Sprintf(`db.collection(\"%s\").add({data:[{
		'classId': '%s',
		'questions': '%s'
	}]})`,
		classInfoTable, classInfo.ClassID, string(questions))
	fmt.Println("sql: ", sql)

	req := map[string]interface{}{
		"env":   "cloud1-4g2pzysxb452412a",
		"query": sql,
	}
	rsp, err := clientPost(url, req)
	fmt.Printf("InsertClassInfo rsp: %v, err: %v\n", string(rsp), err)

	return nil
}

// QueryClassInfo 查询课堂题目
func QueryClassInfo(classID string) (*entity.TClassInfo, error) {
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databasequery?access_token=%s", GetAccessToken())
	req := map[string]interface{}{
		"env":   "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"%s\").where({classId:\"%s\"}).get()`, classInfoTable, classID),
	}
	body, _ := clientPost(url, req)

	resData := &ResData{}
	if err := json.Unmarshal(body, &resData); err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	if len(resData.Data) == 0 {
		return nil, nil
	}

	fmt.Println("resData: ", resData.Data[0])
	tmp := &struct {
		ClassID   string `json:"classId"`
		Questions string `json:"questions"`
	}{}
	_ = json.Unmarshal([]byte(resData.Data[0]), tmp)
	res := &entity.TClassInfo{
		ClassID:   classID,
		Questions: []*entity.QuestionInfo{},
	}
	_ = json.Unmarshal([]byte(tmp.Questions), &res.Questions)

	return res, nil
}

// QueryClassList 查询课堂信息
func QueryClassList(school, grade, class string) ([]*entity.TClass, error) {
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databasequery?access_token=%s", GetAccessToken())
	req := map[string]interface{}{
		"env": "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"%s\").where({school:\"%s\",grade:\"%s\",class:\"%s\"}).get()`,
			classTable, school, grade, class),
	}
	body, _ := clientPost(url, req)

	resData := &ResData{}
	if err := json.Unmarshal(body, &resData); err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	res := make([]*entity.TClass, len(resData.Data))
	for i, val := range resData.Data {
		tmpClass := &entity.TClass{}
		_ = json.Unmarshal([]byte(val), tmpClass)
		res[i] = tmpClass
	}

	return res, nil
}
