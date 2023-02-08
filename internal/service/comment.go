package service

import (
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
	"errors"
	"fmt"
)

func (svc *Service) CommentAction(params request.CommentRequest, userId int64) (res response.CommentResponse, err error) {

	/*
	   根据评论的操作类型进行相应的操作
	   当action_type=1 代表发布评论，即插入
	   当action_type=2 代表删除当前评论
	*/

	err = errors.New("操作类型参数错误！")

	if params.ActionType == 1 { //插入评论
		comment, user, err1 := svc.dao.AddComment(userId, params.VideoId, params.CommentText)
		err = err1

		//user1 := response.User{
		//	ID: user.ID,
		//	Name: user.Username,
		//	FollowCount: user.FollowCount,
		//	FollowerCount: user.FollowerCount,
		//}

		com := response.Comment{
			Id:         int64(comment.ID),
			Content:    comment.Content,
			User:       user,
			CreateDate: comment.CreatedAt.Format("01-02"), //按照mm-dd格式
		}

		res = response.CommentResponse{
			Response: errcode.NewResponse(errcode.OK),
			Comment:  com,
		}

	} else if params.ActionType == 2 { //删除当前评论

		fmt.Println("删除评论id：", params.CommentId)
		err = svc.dao.DeleteComment(uint(params.CommentId))

		res = response.CommentResponse{
			Response: errcode.NewResponse(errcode.OK),
		}
	}

	return
}
