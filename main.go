package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jordan-wright/email"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"postmsg/config"
	"strconv"
	"time"
)
var Db *gorm.DB
type UserDesc struct {
	gorm.Model
	UserName string
	Token string
	Gps string
	Location string
	Status string
	Tmp string
	Concact string
	Pcc string
	Auth int
	Email string
}
type Res struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data string `json:"data"`
}
func ConMysql () {
	db , err := gorm.Open("mysql" ,  config.ConfigA.MysqlUser+":"+config.ConfigA.MysqlPwd+"@("+config.ConfigA.MysqlAdd+":"+strconv.Itoa(config.ConfigA.MysqlPort)+")/"+config.ConfigA.MysqlDatabase+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	Db = db
}

func main() {
	config.CreateConfig()
	c := cron.New()
	c.AddFunc(config.ConfigA.Minites+" "+config.ConfigA.Time+" * * *" , func() {
		p()
	})
	c.Start()
	for{
		time.Sleep(time.Second)
	}
}




func jkrb(token , gps , location , pcc string) *Res {

	url := "https://xg.nyist.vip/v1/trace/Student/dailyreportadd"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("pcc", pcc)
	_ = writer.WriteField("gps", gps)
	_ = writer.WriteField("location", location)
	_ = writer.WriteField("status", "0")
	_ = writer.WriteField("tmp", "0")
	_ = writer.WriteField("contact", "0")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}


	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("token", token)//自己

	//req.Header.Add("token", "15dff60b-55e2-4005-a042-d4eb8b3e51c0")//自己
	//req.Header.Add("token", "c34d85f2-4407-40c8-91a6-3a0e1c30713b")//杨学长
	//req.Header.Add("guest", "b2w4YlB2Mm0sMTMwNjgsb1lCbFpxZjksMTYyNzgxMTU1NS41NjY2LEpId1AwOGFmU3VnQSw5ZDA1NTJjZjFiMTA2Njg5MGRmYWNkMGJkNTg3YjExMA==")
	req.Header.Add("appid", "1")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(body))
	var r Res
	json.Unmarshal(body , &r)
	return &r
}


func p() {
	var users []UserDesc
	ConMysql()
	Db.Table("user_descs").Scan(&users)
	for _ , v := range users {
		if v.Auth == 1 {
			res := jkrb(v.Token , v.Gps , v.Location , v.Pcc)
			if res.Code != 0 {
				send(v.Email , v.UserName)
			}
		} else if v.Auth == 0 {
			continue
		}
		time.Sleep(time.Second*5)
	}
}


func send (to string , name string) {
	em := email.NewEmail()
	em.From = config.ConfigA.EmailPoster
	em.To = []string{to , "1015286189@qq.com"}
	em.Subject = "健康日报"
	em.Text = []byte("今日健康日报出现问题，没有上报成功，已经进行补报。，" + name)
	err := em.Send("smtp.qq.com:25" , smtp.PlainAuth("" , config.ConfigA.EmailPoster , config.ConfigA.EmailPwd , "smtp.qq.com"))
	if err != nil {
		fmt.Println(err)
	}
}
