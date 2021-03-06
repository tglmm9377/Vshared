package obj

import (
	"Vshared/api/db"
	"Vshared/api/utils"
	"fmt"
)
type User struct {
	Id       int    `xorm:"id pk autoincr"`
	Name     string `xorm:"name notnull unique"`
	Password string `xorm:"password notnull"`
	Email string `xorm:"email notnull unique"`
	Vip int `xorm:"vip"`
}
func init(){
	db.MysqlEng.Sync2(new(User))
}

func Register(username , password , email string)error{
	//注册校验操作
	var err error
	var user = new(User)
	// 验证用户是否存在
	name_exist , err := UserExistByName(username)
	if err != nil{
		return err
	}
	if name_exist{
		return fmt.Errorf("user %s already exist",username)
	}


	// 验证email 是否存在
	email_exist  ,err := UserExistByEmail(email)
	if err != nil{
		return err
	}
	if email_exist{
		return fmt.Errorf("email %s is exist" , email)
	}




	user.Password , err = utils.GeneratePassword(password)
	if err != nil{
		return err
	}
	user.Name = username
	user.Email = email
	_ , err = db.MysqlEng.Insert(user)
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}


func Login(name , password string)error{
	// 登陆输入的校验操作
	// 邮箱不用于登陆只用于忘记密码使用

	//判断用户是否存在
	exist , err := UserExistByName(name)
	if err != nil {
		return err
	}else if !exist{
		return fmt.Errorf("user %s not exist",name)
	}

	var user = new(User)
	db.MysqlEng.Where("name=?",name).Get(user)
	err = utils.ComparePassword(user.Password , password)
	return err
}


func GetUserIdByName(name string)(int , error){
	var user = new(User)
	bool , err := db.MysqlEng.Where("name=?" ,name).Get(user)
	if err != nil{
		return 0 , err
	}
	if !bool{
		return 0 ,fmt.Errorf("user %s Not found" , name)
	}
	return user.Id , nil
}

func DeleteUserByName(name string)error{
	_ , err := db.MysqlEng.Where(name).Delete(new(User))
	if err != nil{
		return err
	}
	return nil
}

// 邮箱 ， 密码
func UpdateUserInfo(id  int,email ,password string)error{
	//email 校验
	var user = new(User)
	user.Email = email
	if password != ""{
		password ,err := utils.GeneratePassword(password)
		if err != nil{
			return err
		}
		user.Password = password
	}
	_ , err := db.MysqlEng.ID(id).Update(user)
	if err != nil{
		return err
	}
	return nil
}

func UserExistByName(name string)(bool , error){
	name_exist , err := db.MysqlEng.Exist(&User{Name: name})
	return name_exist , err
}

func UserExistByEmail(email string)(bool , error){
	email_exist , err := db.MysqlEng.Exist(&User{Email: email})
	return email_exist , err
}