package controllers

import (
	"fmt"
	"github.com/NWPU-Paper/Paper/models"
	"time"
	"os"
)

const STATUS_FILE_UPLOAD_SUCCESS = 200
const STATUS_FILE_TOO_LARGE = -10001
const STATUS_FILE_ERROR_UNDEFINED = -10002

type FileController struct {
	BaseController
	status int
	message int
	document models.Document
}

// 上传文件 目录为 /upload key:file value:<file> 会上传到 /static/upload 目录
func (c *FileController) Post() {
	f, h, err := c.GetFile("file")
	defer f.Close()

	c.ErrorJson(err)

	c.document = models.Document{Name:h.Filename,Path:"./static/upload/" + time.Now().Format("2006/01")}

	c.ErrorJson(os.MkdirAll(c.document.Path, 0644))

	path := fmt.Sprintf("%s/%s", c.document.Path, c.document.Name)

	c.ErrorJson(c.SaveToFile("file",path))

	c.document.Path = c.document.Path[1:len(c.document.Path)]
	c.ErrorJson(c.document.Add())

	c.Data["json"] = map[string]interface{}{"success": STATUS_FILE_UPLOAD_SUCCESS, "message": c.document}

	c.ServeJSON()
	c.StopRun()

}

func (c *FileController) ErrorJson(err error) {

	if err != nil {
		// 出错则输出错误信息
		c.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
		c.ServeJSON()
		c.StopRun()
	}
}

