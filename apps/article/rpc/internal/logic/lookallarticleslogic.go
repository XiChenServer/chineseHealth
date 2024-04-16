package logic

import (
	"chineseHealthy/apps/article/model"
	"context"

	"chineseHealthy/apps/article/rpc/internal/svc"
	"chineseHealthy/apps/article/rpc/types/article"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookAllArticlesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookAllArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookAllArticlesLogic {
	return &LookAllArticlesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LookAllArticlesLogic) LookAllArticles(in *article.LookAllArticlesRequest) (*article.LookAllArticlesResponse, error) {
	// todo: add your logic here and delete this line
	res, err := (&model.Article{}).LookAllArticle(l.svcCtx.DB)
	if err != nil {
		return nil, err
	}
	var articles []*article.ArticleInfo
	for _, v := range *res {
		newArticle := article.ArticleInfo{
			Id:      uint32(v.ID),
			Title:   v.Title,
			Content: v.Content,
			Image:   v.Image,
			Author:  v.Author,
		}
		articles = append(articles, &newArticle)
	}
	return &article.LookAllArticlesResponse{
		ArticlesData: articles,
	}, nil
}
