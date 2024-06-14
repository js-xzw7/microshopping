package handler

import (
	pb "adservice/proto"
	"context"
	"math/rand"
)

// 最大展示广告数
const MAX_ADS_TO_SERVER = 2

// 广告map
var adsMap = createAdsMap()

// 广告服务结构体
type Adservice struct{}

func (a *Adservice) GetAds(ctx context.Context, req *pb.AdRequest) (res *pb.AdResponse, err error) {
	allAds := make([]*pb.Ad, 0)

	if len(req.ContextKeys) == 0 {
		ads := getRandomAds()
		allAds = append(allAds, ads...)
	} else {
		for _, key := range req.ContextKeys {
			ads := getAdsByCategory(key)
			allAds = append(allAds, ads...)
		}

		if len(allAds) == 0 {
			ads := getRandomAds()
			allAds = append(allAds, ads...)
		}
	}

	res = &pb.AdResponse{Ads: allAds}
	return res, nil
}

func createAdsMap() map[string][]*pb.Ad {
	hairdryer := &pb.Ad{RedirectUrl: "/product/2ZYFJ3GM2N", Text: "出风机，5折热销"}
	tankTop := &pb.Ad{RedirectUrl: "/product/66VCHSJNUP", Text: "背心8折热销"}
	candleHolder := &pb.Ad{RedirectUrl: "/product/0PUK6V6EV0", Text: "烛台7折热销"}
	bambooGlassJar := &pb.Ad{RedirectUrl: "/product/9SIQT8TOJO", Text: "竹玻璃罐9折"}
	watch := &pb.Ad{RedirectUrl: "/product/1YMWWN1N4O", Text: "手表买一送一"}
	mug := &pb.Ad{RedirectUrl: "/product/6E92ZMYYFZ", Text: "马克杯买二送一"}
	loafers := &pb.Ad{RedirectUrl: "/product/L9ECAV7KIM", Text: "平底鞋，买一送二"}
	return map[string][]*pb.Ad{
		"clothing":    {tankTop},
		"accessories": {watch},
		"footwear":    {loafers},
		"hair":        {hairdryer},
		"decor":       {candleHolder},
		"kitchen":     {bambooGlassJar, mug},
	}
}

func getAdsByCategory(category string) []*pb.Ad {
	return adsMap[category]
}

func getRandomAds() []*pb.Ad {
	ads := make([]*pb.Ad, 0, MAX_ADS_TO_SERVER)
	allAds := make([]*pb.Ad, 0, len(adsMap))

	for _, ads := range adsMap {
		allAds = append(allAds, ads...)
	}

	for i := 0; i < MAX_ADS_TO_SERVER; i++ {
		ads = append(ads, allAds[rand.Intn(len(allAds))])
	}

	return ads
}
