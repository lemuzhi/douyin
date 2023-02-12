package service

import (
	model "douyin/internal/model"
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
	"errors"
)

func (svc *Service) CommentAction(params request.CommentRequest, userId uint) (resp response.CommentResponse, err error) {

	/*
	   根据评论的操作类型进行相应的操作
	   当action_type=1 代表发布评论，即插入
	   当action_type=2 代表删除当前评论
	*/

	err = errors.New("操作类型参数错误！")

	if params.ActionType == 1 { //插入评论
		comment, user, err1 := svc.dao.AddComment(userId, params.VideoId, params.CommentText)
		err = err1

		uResp := response.User{
			ID:            user.ID,
			Name:          user.Username,
			FollowCount:   svc.dao.FollowCount(user.ID),
			FollowerCount: svc.dao.FollowerCount(user.ID),
			IsFollow:      true,
		}

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
func (svc *Service) CommentListAction(params request.CommentListRequest, userId uint) (response.CommentListResponse, error) {

	/*
	   获取当前视频下方的所有评论
	*/

	var commentsRsp []response.Comment

	comments, err := svc.dao.GetCommentsByVideoId(params.VideoId)

	//存储所有作者的id
	idList := make([]uint, len(comments))
	for i := 0; i < len(comments); i++ {
		idList = append(idList, comments[i].UserID)
	}

	// 通过 in 查询 获取视频作者信息
	users, err0 := svc.dao.FindUserIdByIdList(idList)
	err = err0

	//封装查询到的结果
	authorMap := make(map[uint]model.User, len(users))
	for i := 0; i < len(users); i++ {
		authorMap[users[i].ID] = users[i]
	}

	for i := 0; i < len(comments); i++ {

		//评论的作者信息
		//user, err1 := svc.dao.FindUserByID(uint(comments[i].UserID))
		//
		//if err1 != nil {
		//	return response.CommentListResponse{}, err1
		//}

		user := authorMap[uint(comments[i].UserID)]

		followFlag, err2 := svc.dao.IsFollow(userId, comments[i].UserID)

		if err2 != nil {
			return response.CommentListResponse{}, err2
		}

		userRsp := response.User{
			ID:            user.ID,
			Name:          user.Username,
			FollowCount:   svc.dao.FollowCount(user.ID),
			FollowerCount: svc.dao.FollowerCount(user.ID),
			IsFollow:      followFlag,
		}

		cRsp := response.Comment{
			Id:         int64(comments[i].ID),
			Content:    comments[i].Content,
			User:       userRsp,
			CreateDate: comments[i].CreatedAt.Format("01-02"), //按照mm-dd格式
		}

		commentsRsp = append(commentsRsp, cRsp)
	}

	return response.CommentListResponse{
		Response:    errcode.NewResponse(errcode.OK),
		CommentList: commentsRsp,
	}, err
}
