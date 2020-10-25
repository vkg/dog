package bot

import (
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	b.session = s

	if !b.isMsgToBot(m.Content) {
		return
	}

	content := b.suppressMention(m.Content)

	switch {
	case b.isCmd(content, "add emoji"):
		b.AddEmoji(b.getArgs(content, "add emoji"), m.GuildID, m.ChannelID)

	case b.isCmd(content, "help"):
		b.Help(m.ChannelID)

	case b.isCmd(content, "who"):
		b.Who(m.ChannelID)

	case b.isCmd(content, "friend"), b.isCmd(content, "犬"):
		b.Dog(m.ChannelID)

	case b.isCmd(content, "天気"):
		b.Weather(m.ChannelID)

	case b.isCmd(content, "oncall-vkgtaro"):
		b.OncallVkgtaro(m.ChannelID)

	case b.isCmd(content, "dice"), b.isCmd(content, "サイコロ"):
		b.Dice(m.ChannelID)

	case b.isCmd(content, "wiki"):
		b.Wiki(m.ChannelID)

	case b.isCmd(content, "vkgtaro"):
		b.Vkgtaro(m.ChannelID)

	case b.isCmd(content, "vkgfake"):
		b.Vkgfake(m.ChannelID)

	default:
		b.sendMessage(m.ChannelID, "ワン")
	}
}
