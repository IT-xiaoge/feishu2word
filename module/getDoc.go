package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetDocRes struct {
	Code 	int		`json:"code"`
	Msg		string	`json:"msg"`
	Data	GetDocResdata
}

type GetDocResdata struct {
	Files				[]File `json:"files"`
	Next_page_token		string	`json:"next_page_token"`
	Has_more			bool	`json:"has_more"`
}

type File struct {
	Token			string	`json:"token"`
	Name			string	`json:"name"`
	Type			string	`json:"type"`
	Parent_token	string	`json:"parent_token"`
	Url				string		`json:"url"`
	Shortcut_Info	Shortcut_info
	//Extra			string	`json:"extra"`
}

type Shortcut_info struct {
	Target_type		string	`json:"target_type"`
	Target_token	string	`json:"target_token"`
}

//获取我的隐私空间下面的文档
func GetDoc(folder_token,usertoken,savepath string) {
	url := "https://open.feishu.cn/open-apis/drive/v1/files?folder_token=" + folder_token
	//fmt.Println(url)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url,  nil)
	if err != nil {
		fmt.Printf("build get req error %s", err.Error())
		return
	}
	req.Header.Set("Authorization", "Bearer " + usertoken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("get msg request error %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read body error %s", err.Error())
	}

	var gdr GetDocRes
	err = json.Unmarshal(body, &gdr)
	if err != nil {
		fmt.Println(err)
	}

	for _,val := range gdr.Data.Files {
		//fmt.Println(val.Type,val.Url,val.Name)
		if val.Type == "doc" || val.Type == "sheet" || val.Type == "bitable" || val.Type == "mindnote" || val.Type == "file" || val.Type == "wiki" || val.Type == "docx" {
			DownLoadWiki(val.Token,val.Name,val.Url,val.Type,usertoken,savepath)
		} else if val.Type == "folder" {
			GetDoc(val.Token,usertoken,savepath)
		}
		//DownLoadWiki(val.Doc_token,val.Title,val.Url)
	}
}
