package obj

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	fmt.Println(Register("tglmm3" , "123" , "tglmm@gmail.com"))
}


//func TestLogin(t *testing.T) {
//	err := Login("tglmm" , "1234")
//	if err == nil{
//		fmt.Println("login success")
//	}else{
//		fmt.Println(err)
//	}
//}

//func TestUploadVideo(t *testing.T) {
//	UploadVideo(utils.NewUuid() , "rap","https://youtube.com" , "youtube,rap",1)
//}


//func TestUpdateUserInfo(t *testing.T) {
//	UpdateUserInfo(1 , "www.baidu.com" , "123")
//}