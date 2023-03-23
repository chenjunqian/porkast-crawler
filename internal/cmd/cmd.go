package cmd

import (
	"context"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/cache"
	"guoshao-fm-crawler/internal/service/celery"
	"guoshao-fm-crawler/internal/service/celery/jobs"
	"guoshao-fm-crawler/internal/service/celery/worker"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start podcast crawler",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			initCache(ctx)
			initCelery(ctx)
			hold()
			return
		},
	}
)

func initCache(ctx context.Context) {
	cache.InitCache(ctx)
}

func initCelery(ctx context.Context) {
	celery.InitCeleryClient(ctx)
	RegisterCeleryWorker()
	celery.GetClient().StartWorker()
	jobs.StartXiMaLaYaJobs(ctx)
	jobs.StartLizhiJob(ctx)
}

func RegisterCeleryWorker() {
	// XIMALAYA
	celery.GetClient().Register(consts.XIMALAYA_PODCAST_WORKER, worker.ParseXiMaLaYaPodcast)
	celery.GetClient().Register(consts.XIMALAYA_ENTRY_WORKER, worker.ParseXiMaLaYaEntry)
	// LIZHI FM
	celery.GetClient().Register(consts.LIZHI_ENTRY_WORKER, worker.ParseLizhiAllCategories)
	celery.GetClient().Register(consts.LIZHI_CATEGORY_PARSE_WORKER, worker.ParseLizhiPodcastByCategoryPage)
	celery.GetClient().Register(consts.LIZHI_PODCAST_XML_WORKER, worker.ParseLizhiPodcastRSS)
}
func hold() {
	select {}
}
