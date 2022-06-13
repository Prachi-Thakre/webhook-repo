package controllers

import (
	"github.com/spf13/cast"
	"github.com/webhook-repo/models"
	"github.com/webhook-repo/utilities"
	"time"
)

func GetAllData() (returndata utilities.ResponseJSON) {
	resp, err := models.MongoDbConnection.FindAllDoc()
	if err != nil {
		utilities.ErrorResponse(&returndata, err.Error())
		return
	}
	utilities.SuccessResponse(&returndata, resp)
	return
}

func AddData(playload PlayLoad) (returndata utilities.ResponseJSON) {

	doc := models.Doc{}
	doc.Author = playload.Repository.Owner.Login
	doc.Time = time.Now()
	if playload.PullRequest.ID == 0 {
		doc.Action = "PULL_REQUEST"
		doc.FromBranch = playload.PullRequest.Head.Repo.Name
		doc.ToBranch = playload.PullRequest.Base.Repo.Name
		doc.RequestId = cast.ToString(playload.PullRequest.ID)
	} else {
		doc.Action = "PUSH"
		doc.RequestId = cast.ToString(playload.Ref)

	}
	err := models.MongoDbConnection.InserOneDoc(doc)
	if err != nil {
		utilities.ErrorResponse(&returndata, err.Error())
		return
	}
	utilities.SuccessResponse(&returndata, doc)
	return
}
