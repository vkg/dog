package bot

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dty1er/discord/emoji"
	"golang.org/x/text/width"
)

func (b *Bot) AddEmoji(text, guildID, channelID string) {
	if text == "" {
		b.sendMessage(channelID, "引数がないワン")
		return
	}

	// parse arguments
	addEmojiCommand := flag.NewFlagSet("add emoji", flag.ExitOnError)

	bgs := addEmojiCommand.String("bg", "", "background color")
	fgs := addEmojiCommand.String("fg", "", "foreground color")
	alias := addEmojiCommand.String("alias", "", "alias of the emoji")

	args := strings.Split(text, " ")
	if len(args) < 2 {
		b.sendMessage(channelID, "使い方が違うワン add emoji <文字> -alias <alias>")
		return
	}

	emojiText := args[0]
	addEmojiCommand.Parse(args[1:])

	if *alias == "" {
		b.sendMessage(channelID, "-aliasは必須だワン")
		return
	}

	bg := emoji.Colors[*bgs]
	fg := emoji.Colors[*fgs]

	emojiText = width.Widen.String(emojiText)

	path, err := emoji.Gen(emojiText, emoji.Config{Bg: bg, Fg: fg})
	if err != nil {
		b.sendMessage(channelID, "絵文字生成エラー: "+err.Error()+"だワン")
		return
	}

	file, err := os.Open(path)
	if err != nil {
		b.sendMessage(channelID, "絵文字画像Openエラー: "+err.Error()+"だワン")
		return
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		b.sendMessage(channelID, "絵文字画像Openエラー: "+err.Error()+"だワン")
		return
	}

	size := fi.Size()

	data := make([]byte, size)
	file.Read(data)

	base64Image := base64.StdEncoding.EncodeToString(data)
	base64Image = "data:image/png;base64," + base64Image

	_, err = b.session.GuildEmojiCreate(guildID, *alias, base64Image, nil)
	if err != nil {
		b.sendMessage(channelID, "絵文字登録エラー: "+err.Error()+"だワン")
		return
	}

	b.sendMessage(channelID, fmt.Sprintf("絵文字を作っといたワン: :%s:", *alias))
}
