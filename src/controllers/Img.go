package controllers

import (
	"net/http"

	"fmt"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"io"
	"os"
	"strings"

	. "../log"
	. "../models"
	"../utils"
)

const UPLOAD_PATH string = "./../image/"

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 判断并创建image文件夹
func CreateDir() {
	exist, err := PathExists(UPLOAD_PATH)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}

	if !exist {
		// 创建文件夹
		err := os.Mkdir(UPLOAD_PATH, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		}
	}
}

func Entrance(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./uploadImg.html")
	t.Execute(w, nil)
}

func UploadImg(w http.ResponseWriter, r *http.Request) {
	CreateDir()

	var img Img
	img.Id = bson.NewObjectId()

	r.ParseMultipartForm(1024)
	imgFile, imgHead, imgErr := r.FormFile("img")
	if imgErr != nil {
		fmt.Println(imgErr)
		return
	}
	defer imgFile.Close()

	imgFormat := strings.Split(imgHead.Filename, ".")
	img.ImgUrl = img.Id.Hex() + "." + imgFormat[len(imgFormat)-1]

	image, err := os.Create(UPLOAD_PATH + img.ImgUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer image.Close()

	_, err = io.Copy(image, imgFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Db["Imgs"].Insert(img)
	if err != nil {
		Log.Error("upload image failed: insert into db failed, ", err)
		utils.FailureResponse(&w, 500, "新建图片失败", "")
		return
	}

	Log.Notice("upload image successfully")
	utils.SuccessResponse(&w, 200, "上传图片成功", "./image/"+img.ImgUrl)
}

var ImgRoutes Routes = Routes{
	Route{"UploadImg", "POST", "/uploadImg", UploadImg},
	Route{"Entrance", "GET", "/entrance", Entrance},
}
