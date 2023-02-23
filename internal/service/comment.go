package service

import (
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
	"errors"
	"fmt"
)

func (svc *Service) CommentAction(params *request.CommentRequest, userId uint) (resp response.CommentResponse, err error) {

	/*
	   根据评论的操作类型进行相应的操作
	   当action_type=1 代表发布评论，即插入
	   当action_type=2 代表删除当前评论
	*/

	err = errors.New("操作类型参数错误！")

	if params.ActionType == 1 { //插入评论
		comment, user, err1 := svc.dao.AddComment(userId, params.VideoId, params.CommentText)
		err = err1

		uResp, _ := svc.dao.GetUserInfo(user.ID)

		comResp := response.Comment{
			Id:         int64(comment.ID),
			Content:    comment.Content,
			User:       uResp,
			CreateDate: comment.CreatedAt.Format("01-02"), //按照mm-dd格式
		}

		resp = response.CommentResponse{
			Response: errcode.NewResponse(errcode.OK),
			Comment:  comResp,
		}

	} else if params.ActionType == 2 { //删除当前评论

		err = svc.dao.DeleteComment(uint(params.CommentId))

		resp = response.CommentResponse{
			Response: errcode.NewResponse(errcode.OK),
		}
	}

	return
}
func (svc *Service) CommentListAction(params *request.CommentListRequest, userId uint) (respData *response.CommentListResponse, err error) {

	/*
	   获取当前视频下方的所有评论
	*/

	var commentsRsp []response.Comment

	comments, err := svc.dao.GetCommentsByVideoId(params.VideoId)
	if err != nil {
		return respData, err
	}

	//存储所有作者的id
	idList := make([]uint, len(comments))
	for i := 0; i < len(comments); i++ {
		idList = append(idList, comments[i].UserID)
	}

	// 通过 in 查询 获取视频作者信息
	authorMap, err := svc.dao.GetUserInfoList(idList)
	if err != nil {
		return respData, err
	}

	for i := 0; i < len(comments); i++ {
		//评论的作者信息
		userRsp := authorMap[comments[i].UserID]
		if userId != 0 {
			isFollow, err := svc.dao.IsFollow(userId, comments[i].UserID)
			if err != nil {
				fmt.Println("CommentListAction IsFollow error:", err)
			}
			userRsp.IsFollow = isFollow
		}

		cRsp := response.Comment{
			Id:         int64(comments[i].ID),
			Content:    comments[i].Content,
			User:       userRsp,
			CreateDate: comments[i].CreatedAt.Format("01-02"), //按照mm-dd格式
		}

		commentsRsp = append(commentsRsp, cRsp)
	}

	respData = &response.CommentListResponse{
		Response:    errcode.NewResponse(errcode.OK),
		CommentList: commentsRsp,
	}

	return respData, err
}
