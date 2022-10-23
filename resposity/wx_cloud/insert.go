package wx_cloud

import (
	"encoding/json"
	"fmt"
	"wx_cloud/entity"
)

// InsertClass 记录课堂信息
func InsertClass(class *entity.TClass) error {
	url := fmt.Sprintf("https://api.weixin.qq.com/tcb/databaseadd?access_token=%s", GetAccessToken())
	body, _ := json.Marshal(class)
	req := map[string]interface{}{
		"env":   "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"t_class\").add({data:[%s]})`,
			string(body)),
	}
	rsp, _ := clientPost(url, req)
	fmt.Printf("InsertClass rsp: %v\n", string(rsp))

	//resData := &ResData{}
	//if err := json.Unmarshal(rsp, &resData); err != nil {
	//	fmt.Println("err: ", err)
	//	return err
	//}
	//if resData.Errcode != 0 {
	//	err := fmt.Errorf("InsertClass fail, rsp: %v", resData)
	//	return err
	//}

	return nil
}

// InsertClassInfo 记录课堂题目
func InsertClassInfo(classInfo *entity.TClassInfo) error {
	url := fmt.Sprintf("https://api.weixin.qq.com/tcb/databaseadd?access_token=%s", GetAccessToken())
	questions, err := json.Marshal(classInfo.Questions)
	fmt.Println("err: ", err)
	//tmp := &struct {
	//	ClassID   string `json:"classId"`
	//	Questions string `json:"questions"`
	//}{
	//	ClassID:   classInfo.ClassID,
	//	Questions: string(questions),
	//}
	//body, err := json.Marshal(tmp)
	//fmt.Println("err: ", err)
	sql := fmt.Sprintf(`db.collection(\"t_class_info\").add({data:[{
		'classId': '%s',
		'questions': '%s'
	}]})`,
		classInfo.ClassID, string(questions))
	fmt.Println("sql: ", sql)

	req := map[string]interface{}{
		"env":   "cloud1-4g2pzysxb452412a",
		"query": sql,
	}
	rsp, err := clientPost(url, req)
	fmt.Printf("InsertClassInfo rsp: %v, err: %v\n", string(rsp), err)

	//resData := &ResData{}
	//if err := json.Unmarshal(rsp, &resData); err != nil {
	//	fmt.Println("err: ", err)
	//	return err
	//}
	//if resData.Errcode != 0 {
	//	err := fmt.Errorf("InsertClass fail, rsp: %v", resData)
	//	return err
	//}

	return nil
}
