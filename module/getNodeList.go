package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetNodeListRes struct {
	Code              int64  `json:"code"`
	Msg               string `json:"msg"`
	Data 			  GetNodeListResDatas
}

type GetNodeListResDatas struct {
	Has_more		bool `json:"has_more"`
	Page_token		string `json:"page_token"`
	Items 			[]GetNodeListResDatasItmes
}

type GetNodeListResDatasItmes struct {
	Space_id			string `json:"space_id"`
	Node_token			string `json:"node_token"`
	Obj_token			string `json:"obj_token"`
	Obj_type			string `json:"obj_type"`
	Parent_node_token	string `json:"parent_node_token"`
	Node_type			string `json:"node_type"`
	Origin_node_token	string `json:"origin_node_token"`
	Origin_space_id		string `json:"origin_space_id"`
	Has_child			bool 	`json:"has_child"`
	Title				string `json:"title"`
	Obj_create_time		string `json:"obj_create_time"`
	Obj_edit_time		string `json:"obj_edit_time"`
	Node_create_time	string `json:"node_create_time"`
	Creator				string `json:"creator"`
	Owner				string `json:"owner"`
}

func GetNodeList(space_id,page_token,parent_node_token,usertoken,savepath string) {
	url := "https://open.feishu.cn/open-apis/wiki/v2/spaces/" + space_id +"/nodes?page_size=50&page_token=" + page_token + "&parent_node_token=" + parent_node_token
	client := &http.Client{}

	req, err := http.NewRequest("GET", url,  nil)
	if err != nil {
		fmt.Printf("build get req error：%s", err.Error())
		return
	}
	req.Header.Set("Authorization", "Bearer " + usertoken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("get response error：%s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read body error：%s", err.Error())
	}

	var gnlr GetNodeListRes
	err = json.Unmarshal(body, &gnlr)
	if err != nil {
		fmt.Printf("Unmarshal body error：%s", err.Error())
	}

	//节点超过50个，需要多次遍历
	if gnlr.Data.Has_more {
		GetNodeList(space_id,gnlr.Data.Page_token,parent_node_token,usertoken,savepath)
	}

	for _,val := range gnlr.Data.Items {
		//判断当前节点是否有子节点
		if val.Has_child {
			GetNodeList(val.Space_id,"",val.Node_token,usertoken,savepath)
		}
		//如果没有子节点，获取当前文档的元信息
		GetWikiMeta(val.Node_token,usertoken,savepath)
	}
}

