package obj

import "Vshared/api/db"

type Video struct{
	Id string `xorm:"id pk"`
	Name string `xorm:"name"`
	Url string `xorm:"url"`
	Tag string `xorm:"tag"`
	Uid int `xorm:"uid"`
}
func init(){
	db.MysqlEng.Sync2(new(Video))
}

func UploadVideo(id , name , url , tag  string, uid int)error{
	var video = new(Video)
	video.Id = id
	video.Name = name
	video.Url = url
	video.Tag = tag
	video.Uid = uid
	_ , err := db.MysqlEng.Insert(video)
	if err != nil{
		return err
	}
	return nil
}


func DeleteVideo(id string)error{

	_ , err := db.MysqlEng.Delete(Video{Id: id})
	return err
}

func updateVideoInfo(id string)error{
	return nil
}


func GetVideoInfo(id string)(*Video , error){

}





