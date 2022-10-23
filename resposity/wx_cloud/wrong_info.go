package wx_cloud

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"wx_cloud/entity"
)

const (
	// 错题集信息
	wrongInfoTable = "t_wrong_info"
)

// InsertWrongInfo 记录错题信息
func InsertWrongInfo(wrongInfo *entity.TWrongInfo) error {
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databaseadd?access_token=%s", GetAccessToken())
	body, _ := json.Marshal(wrongInfo)
	req := map[string]interface{}{
		"env": "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"%s\").add({data:[%s]})`,
			wrongInfoTable, string(body)),
	}
	rsp, _ := clientPost(url, req)
	fmt.Printf("InsertWrongInfo rsp: %v\n", string(rsp))

	return nil
}

// DeleteWrongInfo 删除错题记录
func DeleteWrongInfo(classID, openID string) error {
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databaseadd?access_token=%s", GetAccessToken())
	req := map[string]interface{}{
		"env": "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"%s\").where({classId:%s,openId:%s}).remove()`,
			wrongInfoTable, classID, openID),
	}
	rsp, _ := clientPost(url, req)
	fmt.Printf("DeleteWrongInfo rsp: %v\n", string(rsp))

	return nil
}

// QueryClassWrongInfo 查询错题集
func QueryClassWrongInfo(classID, openID string) (*entity.TClassInfo, error) {
	classInfo, _ := QueryClassInfo(classID)
	wrongInfo, _ := queryWrongInfoByOpenID(classID, openID)

	res := &entity.TClassInfo{
		ClassID:   classID,
		Questions: []*entity.QuestionInfo{},
	}
	wrongIDs := strings.Split(wrongInfo.WrongIDs, ",")
	for _, val := range wrongIDs {
		questionIndex, _ := strconv.Atoi(val)
		res.Questions = append(res.Questions, classInfo.Questions[questionIndex])
	}
	for _, val := range res.Questions {
		for _, option := range val.Options {
			if option.Value == val.Right {
				option.Checked = "true"
			}
		}
	}
	return res, nil
}

func queryWrongInfoByOpenID(classID, openID string) (*entity.TWrongInfo, error) {
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databasequery?access_token=%s", GetAccessToken())
	req := map[string]interface{}{
		"env": "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"%s\").where({classId:\"%s\",openId:\"%s\"}).get()`,
			wrongInfoTable, classID, openID),
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
	res := &entity.TWrongInfo{}
	_ = json.Unmarshal([]byte(resData.Data[0]), res)

	return res, nil
}
