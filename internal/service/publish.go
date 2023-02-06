package service

import (
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
	"douyin/pkg/upload"
	"douyin/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mime/multipart"
)

func (svc *Service) PublishAction(c *gin.Context, file *multipart.FileHeader) error {
	if !upload.CheckFileExt(file.Filename) {
		return errors.New("请上传.mp4格式的视频")
	}

	videoTitle := upload.GetIntactFileName(file.Filename)
	filepath := upload.GetSavePath()             //保存的文件夹
	addr := viper.GetString("upload.server_url") //静态服务地址
	videoUrl := filepath + videoTitle            //存放视频的相对路径
	// 上传文件至指定的完整文件路径
	err := c.SaveUploadedFile(file, videoUrl)
	if err != nil {
		return err
	}

	coverTitle := upload.GetFileName(videoTitle) + ".jpeg"
	coverUrl := filepath + coverTitle //存放封面的相对路径

	utils.FfmpegCoverJpeg(videoUrl, coverUrl, 5)

	return svc.dao.PublishAction(c.GetInt64("UserID"), c.PostForm("title"), addr+videoTitle, addr+coverTitle)
}

func (svc *Service) PublishList(id string) (*response.PublishListResponse, error) {
	videoList, err := svc.dao.GetPublishList(id)
	if err != nil {
		return nil, err
	}

	fmt.Println("视频")
	fmt.Println(videoList)

	data := response.PublishListResponse{
		Response:  errcode.NewResponse(errcode.OK),
		VideoList: videoList,
	}
	return &data, nil
}
