package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	clientID string
	session  *discordgo.Session
}

func New(clientID string) *Bot {
	return &Bot{clientID: clientID}
}

func (b *Bot) isMsgToBot(s string) bool {
	// from PC
	return strings.HasPrefix(s, fmt.Sprintf("<@!%s>", b.clientID)) ||
		// from mobile app
		strings.HasPrefix(s, fmt.Sprintf("<@%s>", b.clientID))
}

func (b *Bot) suppressMention(s string) string {
	s = strings.TrimPrefix(s, fmt.Sprintf("<@!%s> ", b.clientID))
	return strings.TrimPrefix(s, fmt.Sprintf("<@%s> ", b.clientID))
}

func (b *Bot) isCmd(content, cmd string) bool {
	return strings.HasPrefix(content, cmd)
}

func (b *Bot) getArgs(content, cmd string) string {
	content = strings.TrimLeft(content, cmd)
	return strings.Trim(content, " ")
}

func (b *Bot) sendMessage(channelID, msg string) {
	if _, err := b.session.ChannelMessageSend(channelID, msg); err != nil {
		log.Println("Error sending message: ", err)
	}
}
