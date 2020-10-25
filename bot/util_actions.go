package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func (b *Bot) Help(channelID string) {
	b.sendMessage(channelID, "@vkgdog subcommand\n"+
		"サブコマンド\n"+
		"  add emoji <文字> -alias <name>\t絵文字を追加するワン(文字の長さは16文字までだワン)\n"+
		"  who\t自己紹介するワン\n"+
		"  天気\t天気をランダムで教えるワン\n"+
		"  friend|犬\t友達の犬をお届けするワン\n"+
		"  oncall-vkgtaro\tvkgtaroを呼ぶワン\n"+
		"  dice|サイコロ\tサイコロを振るワン\n"+
		"  help\thelpを表示するワン\n"+
		"  wiki\twikipediaからランダムでなんか出すワン\n"+
		"  vkgtaro\tvkgtaroをお届けするワン\n"+
		"  vkgfake\tvkgtaroの偽物をお届けするワン\n"+
		"  training\tぼくのトレーニング場をご案内するワン\n")
}

func (b *Bot) Dice(channelID string) {
	randomVal := time.Now().Unix() % 6
	b.sendMessage(channelID, fmt.Sprintf(`%dが出たワン`, randomVal+1))
}

func (b *Bot) OncallVkgtaro(channelID string) {
	b.sendMessage(channelID, `<@768330139679326220> 呼んでるよ`)
}

func (b *Bot) Vkgtaro(channelID string) {
	vkgPhotos := []string{
		"https://www.flickr.com/photos/vkgtaro/5942816964/",
		"https://www.flickr.com/photos/vkgtaro/5942228821/",
		"https://www.flickr.com/photos/vkgtaro/5089308191/",
		"https://www.flickr.com/photos/vkgtaro/4611059803/",
		"https://www.flickr.com/photos/vkgtaro/4571642918/",
		"https://www.flickr.com/photos/vkgtaro/4571618624/",
		"https://www.flickr.com/photos/vkgtaro/4392371252/",
		"https://www.flickr.com/photos/vkgtaro/4392355042/",
		"https://www.flickr.com/photos/vkgtaro/3912240511/",
		"https://www.flickr.com/photos/vkgtaro/3912205989/",
		"https://www.flickr.com/photos/vkgtaro/3088737669/",
		// keep this image at the bottom
		"https://user-images.githubusercontent.com/16610193/97103910-ebdda400-16f2-11eb-91d5-4749ef23b151.png", // fake taro
	}
	randomVal := int(time.Now().Unix()) % len(vkgPhotos)
	prefix := "どうぞ"
	if randomVal == len(vkgPhotos)-1 {
		prefix = "大当たり どうぞ"
	}
	b.sendMessage(channelID, fmt.Sprintf(`%s %s`, prefix, vkgPhotos[randomVal]))
}

func (b *Bot) Vkgfake(channelID string) {
	image := "https://user-images.githubusercontent.com/16610193/97103910-ebdda400-16f2-11eb-91d5-4749ef23b151.png"
	b.sendMessage(channelID, fmt.Sprintf(`偽物だワン %s`, image))
}

func (b *Bot) Training(channelID string) {
	url := "https://www.google.com/maps/place/VKG+Dog+Training+Centre/@11.7234491,76.3007049,15z/data=!4m5!3m4!1s0x0:0x821414fa47a0ac8b!8m2!3d11.7234491!4d76.3007049"
	b.sendMessage(channelID, fmt.Sprintf(`vkg dog training centreで遊びたいワン %s`, url))
}

func (b *Bot) Weather(channelID string) {
	randomVal := time.Now().Unix() % 5
	weathers := map[int64]string{
		0: "晴れ",
		1: "雨",
		2: "雪",
		3: "雷",
		4: "曇り",
	}
	b.sendMessage(channelID, fmt.Sprintf(`明日の天気は%sだワン`, weathers[randomVal]))
}

func (b *Bot) Who(channelID string) {
	b.sendMessage(channelID, "vkgdogはvkgtaroの身の回りのお世話をする高度に訓練された人工知能だワン")
}

func (b *Bot) Dog(channelID string) {
	type Response struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}
	url := "https://dog.ceo/api/breeds/image/random"
	resp, err := http.Get(url)
	if err != nil {
		b.sendMessage(channelID, "犬が見つからなかったワン")
		return
	}

	defer resp.Body.Close()

	var res Response
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		b.sendMessage(channelID, "深刻なエラーが発生したようだワン")
		return
	}

	b.sendMessage(channelID, res.Message)
}

func (b *Bot) Wiki(channelID string) {
	type Response struct {
		Query struct {
			Random []struct {
				Title string `json:"title"`
			} `json:"random"`
		} `json:"query"`
	}
	u := "https://ja.wikipedia.org/w/api.php?action=query&list=random&rnnamespace=0&rnlimit=10&format=json"
	resp, err := http.Get(u)
	if err != nil {
		b.sendMessage(channelID, "犬が見つからなかったワン")
		return
	}

	defer resp.Body.Close()

	var res Response
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		b.sendMessage(channelID, "深刻なエラーが発生したようだワン")
		return
	}

	article := fmt.Sprintf("https://ja.wikipedia.org/wiki/%s", url.QueryEscape(res.Query.Random[0].Title))
	b.sendMessage(channelID, fmt.Sprintf("これを読むワン %s", article))
}
