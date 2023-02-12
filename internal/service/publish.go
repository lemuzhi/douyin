package service

import (
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
	"douyin/pkg/upload"
	"douyin/pkg/utils"
	"errors"
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

	//utils.FfmpegCoverJpeg(videoUrl, coverUrl, 5)
	utils.RunFfmpegCoverJpeg(videoUrl, coverUrl)

	return svc.dao.PublishAction(c.GetUint("UserID"), c.PostForm("title"), addr+videoTitle, addr+coverTitle)
}

func (svc *Service) PublishList(params *request.PublishListRequest) (*response.PublishListResponse, error) {
	var beUserID uint
	if params.UserID > 0 {
		claims, err := utils.ParseToken(params.Token)
		if err != nil {
			return nil, err
		}
		beUserID = claims.UserID
	}

	videoList, err := svc.dao.GetPublishList(params.UserID, beUserID)
	if err != nil {
		return nil, err
	}

	data := response.PublishListResponse{
		Response:  errcode.NewResponse(errcode.OK),
		VideoList: videoList,
	}
	return &data, nil
}
