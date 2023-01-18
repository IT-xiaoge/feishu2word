package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetSpacesRes struct {
	Code              int64  `json:"code"`
	Msg               string `json:"msg"`
	Data 			  GetSpacesResData
}

type GetSpacesResData struct {
	Has_more		bool `json:"has_more"`
	Page_token		string `json:"page_token"`
	Items 			[]GetSpacesResDataItems
}

type GetSpacesResDataItems struct {
	Description		string `json:"description"`
	Name			string `json:"name"`
	Space_id		string `json:"space_id"`
	Space_type		string `json:"space_type"`
	Visibility		string `json:"visibility"`
}

func GetSpaces(usertoken,savepath string) {
	url := "https://open.feishu.cn/open-apis/wiki/v2/spaces"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("build get request error：%s", err.Error())
		return
	}
	req.Header.Set("Authorization", "Bearer " + usertoken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("get response error：%s", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read body error：%s", err.Error())
		return
	}

	var gsr GetSpacesRes
	err = json.Unmarshal(body, &gsr)
	if err != nil {
		fmt.Printf("Unmarshal body error：%s", err.Error())
		return
	}

	for _,val := range gsr.Data.Items {
		//fmt.Println(val.Name,val.Space_id)
		GetNodeList(val.Space_id,"","",usertoken,savepath)
	}
}