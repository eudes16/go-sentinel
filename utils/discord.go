package utils

import (
	"eudes16/go-sentinel/entities"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var discord *discordgo.Session
var err error
var BotId string

func InitBot() {
	_ = godotenv.Load(".env")

	discord, err = discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))


	if err != nil {
		log.Println(err)
		return
	}

	u, err := discord.User("@me")

	if err != nil {
		log.Println(err.Error())
		return
	}

	BotId = u.ID

	err = discord.Open()

	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Print("Bot is running!")
}

func SendMessage(l *entities.Log) {

	embed := NewEmbed().
		SetTitle("Houve um novo erro no " + l.Application.Name).
		SetDescription(l.Message).
		AddField("Usuário", l.UserID).
		AddField("Imobiliária", l.ClientID).
		AddField("Arquivo", l.File).
		AddField("StackTrace", l.StackTrace).
		SetColor(0xff0000).MessageEmbed

	discord.ChannelMessageSendEmbed(os.Getenv("DISCORD_CHANNEL_ID"), embed)
}
