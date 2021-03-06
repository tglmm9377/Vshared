package obj

import "time"

type Comment struct{
	Id string `xorm:"id"`
	Content string `xorm:"content"`
	Uid int `xorm:"content"`
	Ctime time.Time `xorm:"created"`
}

