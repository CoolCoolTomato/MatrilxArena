package api

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/CoolCoolTomato/MatrilxArena/Server/config"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strconv"
)

func GetAttachmentList(c *gin.Context) {
	var attachment model.Attachment
	err := c.ShouldBindJSON(&attachment)
	if err != nil {
		response.Fail(err, localizer.GetMessage("Attachment.InvalidArgument", c), c)
		return
	}

	attachmentList, err := model.GetAttachmentList(attachment)
	if err != nil {
		response.Fail(err, localizer.GetMessage("Attachment.GetAttachmentListFail", c), c)
		return
	}

	response.OK(attachmentList, localizer.GetMessage("Attachment.GetAttachmentListSuccess", c), c)
}

func GetAttachment(c *gin.Context) {
	var attachment model.Attachment
	err := c.ShouldBindJSON(&attachment)
	if err != nil || attachment.ID == 0 {
		response.Fail(err, localizer.GetMessage("Attachment.InvalidArgument", c), c)
		return
	}

	err = attachment.GetAttachment()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Attachment.GetAttachmentFail", c), c)
		return
	}

	response.OK(attachment, localizer.GetMessage("Attachment.GetAttachmentSuccess", c), c)
}

func CreateAttachment(c *gin.Context) {
	var attachment model.Attachment
	err := c.ShouldBindJSON(&attachment)
	if err != nil || attachment.FileName == "" || attachment.FilePath == "" {
		response.Fail(err, localizer.GetMessage("Attachment.InvalidArgument", c), c)
		return
	}

	err = attachment.CreateAttachment()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Attachment.CreateAttachmentFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Attachment.CreateAttachmentSuccess", c), c)
}

func UpdateAttachment(c *gin.Context) {
	var attachment model.Attachment
	err := c.ShouldBindJSON(&attachment)
	if err != nil || attachment.ID == 0 || attachment.FileName == "" || attachment.FilePath == "" {
		response.Fail(err, localizer.GetMessage("Attachment.InvalidArgument", c), c)
		return
	}

	err = attachment.UpdateAttachment()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Attachment.UpdateAttachmentFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Attachment.UpdateAttachmentSuccess", c), c)
}

func DeleteAttachment(c *gin.Context) {
	var attachment model.Attachment
	err := c.ShouldBindJSON(&attachment)
	if err != nil || attachment.ID == 0 {
		response.Fail(err, localizer.GetMessage("Attachment.InvalidArgument", c), c)
		return
	}

	err = attachment.DeleteAttachment()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Attachment.DeleteAttachmentFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Attachment.DeleteAttachmentSuccess", c), c)
}

func UploadAttachment(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Fail(err, localizer.GetMessage("Attachment.UploadFileFail", c), c)
		return
	}

	filePath := filepath.Join(config.GetConfig().GetString("AttachmentPath"), generateRandomFilename()+filepath.Ext(file.Filename))
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		response.Fail(err, localizer.GetMessage("Attachment.SaveFileFail", c), c)
		return
	}

	var fileName string
	if c.PostForm("fileName") == "" {
		fileName = filepath.Base(file.Filename)
	} else {
		fileName = c.PostForm("fileName")
	}

	var attachment model.Attachment
	attachment.FileName = fileName
	attachment.FilePath = filePath
	err = attachment.CreateAttachment()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Attachment.CreateAttachmentFail", c), c)
		return
	}

	response.OK(attachment, localizer.GetMessage("Attachment.UploadAttachmentSuccess", c), c)
}

func DownloadAttachment(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Fail(err, localizer.GetMessage("Attachment.InvalidArgument", c), c)
		return
	}

	var attachment model.Attachment
	attachment.ID = uint(ID)
	err = attachment.GetAttachment()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Attachment.GetAttachmentFail", c), c)
		return
	}

	if _, err := os.Stat(attachment.FilePath); os.IsNotExist(err) {
		response.Fail(err, localizer.GetMessage("Attachment.FileNotFound", c), c)
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+attachment.FileName)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")
	c.File(attachment.FilePath)
}

func generateRandomFilename() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
