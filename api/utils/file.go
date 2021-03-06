package utils

import(
	"github.com/toolkits/pkg/file"
)

func FileIsExist(path string)bool{
	return file.IsExist(path)
}
