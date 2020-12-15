package biz

import (
	"context"
	"github.com/orvice/ab-job/pkg/mod"
	"github.com/weeon/log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	tgs = make([]string, 0)
)

func Init() {
	log.SetupStdoutLogger()

	tgstr := os.Getenv("TGS")
	tgs = strings.Split(tgstr, ",")
	log.Infow("",
		"tgs", tgs,
	)
	ctx := context.Background()

	for _, u := range tgs {
		startHttpJob(ctx, u, 50)
	}
}

func startHttpJob(ctx context.Context, u string, number int) {
	uri, err := url.Parse(u)
	if err != nil {
		log.Errorw("parse url error",
			"error", err)
		return
	}
	log.Infow("start http job",
		"url", uri.String(),
		"host", uri.Host,
	)
	var i = 0
	for i < number {
		go httpJob(ctx, u)
		i++
	}
}

func httpJob(ctx context.Context, u string) {
	for {
		select {
		case <-ctx.Done():
			break
		default:
			ret, err := httpGet(ctx, u)
			if err != nil {
				log.Errorw("http get error",
					"error", err)
				continue
			}
			log.Infow("http get ",
				"ret", ret)
		}
	}
}

func httpGet(ctx context.Context, u string) (*mod.HttpGetRet, error) {
	before := time.Now()
	resp, err := http.DefaultClient.Get(u)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	mill := time.Now().Sub(before).Milliseconds()

	return &mod.HttpGetRet{
		Milliseconds: mill,
	}, nil
}
