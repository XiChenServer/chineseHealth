package article

import (
	"chineseHealthy/apps/article/rpc/types/article"
	"context"

	"chineseHealthy/apps/app/api/internal/svc"
	"chineseHealthy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookAllArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookAllArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookAllArticlesLogic {
	return &LookAllArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookAllArticlesLogic) LookAllArticles(req *types.LookAllArticleRequest) (resp *types.LookAllArticleResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.ArticleRpc.LookAllArticles(l.ctx, &article.LookAllArticlesRequest{})
	if err != nil {
		return nil, err
	}
	var articleInfo []*types.ArticleInfo
	for _, v := range res.ArticlesData {
		articleInfo = append(articleInfo, &types.ArticleInfo{
			Id:      v.Id,
			Title:   v.Title,
			Content: v.Content,
			Image:   v.Image,
			Author:  v.Author,
		})
	}
	return &types.LookAllArticleResponse{Data: articleInfo}, nil
}
