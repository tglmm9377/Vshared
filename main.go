package main


import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	p ,_:= bcrypt.GenerateFromPassword([]byte("admin") , bcrypt.DefaultCost)
	fmt.Println(string(p))
	err := bcrypt.CompareHashAndPassword(p , []byte("admin1"))
	if err == nil{
		fmt.Println("okok")
	}else {
		fmt.Println(err)
	}
}
