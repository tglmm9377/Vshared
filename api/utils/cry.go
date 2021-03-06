package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string)(string ,error){
	pwd , err := bcrypt.GenerateFromPassword([]byte(password) , bcrypt.DefaultCost)
	if err != nil{
		fmt.Println("GeneratePassword error:",err)
		return "" , err
	}
	return string(pwd) , nil
}


func ComparePassword(hash , password string)error{
	err := bcrypt.CompareHashAndPassword([]byte(hash) , []byte(password))
	return err
}



