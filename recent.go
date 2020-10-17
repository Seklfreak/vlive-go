package vlive_go

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const recentURL = "https://www.vlive.tv/home/video/more?pageNo=1&pageSize=50&viewType=recent"

func (v *VLive) Recents() ([]*Video, error) {
	req, err := http.NewRequest(http.MethodGet, recentURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := v.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	mainDoc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	videos := make([]*Video, 0, 50)

	mainDoc.Find("li").Each(func(i int, selection *goquery.Selection) {
		video := &Video{}

		video.Title, _ = selection.Find("a.thumb_area").First().Attr("title")
		video.URL, _ = selection.Find("a.thumb_area").First().Attr("href")
		if strings.HasPrefix(video.URL, "/") {
			video.URL = "https://www.vlive.tv" + video.URL
		}
		videoType, _ := selection.Find("a.thumb_area").First().Attr("data-ga-type")
		video.Type = VideoType(videoType)
		video.Seq, _ = selection.Find("a.thumb_area").First().Attr("data-seq")
		videoProduct, _ := selection.Find("a.thumb_area").First().Attr("data-ga-product")
		video.Product = VideoProduct(videoProduct)
		video.Thumbnail, _ = selection.Find("a.thumb_area img").First().Attr("src")
		if strings.HasPrefix(video.URL, "/") {
			video.Thumbnail = "https://www.vlive.tv" + video.URL
		}
		video.ChannelName, _ = selection.Find("a.thumb_area").First().Attr("data-ga-cname")
		videoChannelType, _ := selection.Find("a.thumb_area").First().Attr("data-ga-ctype")
		video.ChannelType = ChannelType(videoChannelType)
		video.ChannelSeq, _ = selection.Find("a.thumb_area").First().Attr("data-ga-cseq")
		video.ChannelId, _ = selection.Find("div.video_date a.name").First().Attr("href")
		video.ChannelId = strings.TrimPrefix(video.ChannelId, "/channels/")

		// get live thumbnail if no thumbnail in html
		if video.Thumbnail == "" && video.Type == "LIVE" { // TODO: TYPE FOR LIVE
			video.Thumbnail = fmt.Sprintf("https://vlive-thumb.pstatic.net/live/%s/thumb?type=f228_128", video.Seq)
		}

		if video.Seq == "" || video.Title == "" || video.URL == "" {
			return
		}

		videos = append(videos, video)
	})

	if len(videos) <= 0 {
		return nil, errors.New("unable to find any videos on page")
	}

	return videos, nil
}
