package main

import (
	"fmt"
	"github.com/EvgenyOvsov/DiscordMicroservice/integrations"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)
var(
	FuncMap = map[string]func(m *discordgo.MessageCreate, args []string){
		"help":	help,
		"clean": clean,
		"delete": deleteobject,
		"renew": emergencydroplet,
	}
)


func delete(m *discordgo.Message, interval time.Duration){
	time.Sleep(interval)
	err := Discord.client.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		fmt.Println(err)
	}
}

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate){
	text := strings.Split(m.Content, " ")
	if len(text[0])<1 || text[0][:1]!="/"{return}
	text[0] = text[0][1:]
	for k,v := range FuncMap{
		if text[0]==k{
			fmt.Printf("%v in %v used command: %v\n", m.Author.Username, Discord.GetChName(m.ChannelID), text)
			v(m, text[1:])
			return
		}
	}
	return
}

func help(m *discordgo.MessageCreate, args []string){
	text := fmt.Sprintf("VQBot ver %v\n" +
		"Вот что я умею: \n\n" +
		"/help - Показать это сообщение\n" +
		"\n" +
		"# Канал portal:\n" +
		"/clean - Очистить тестовую базу данных портала\n" +
		"/delete - Удалить лицензионный ключ\n" +
		"/renew - Перезапустить облачный сервер\n",
		Version)
	message := Discord.MessageUser(m.Author.ID, text)
	Discord.client.ChannelMessageDelete(m.ChannelID, m.ID)
	go delete(message, 1*time.Minute)
}

func clean(m *discordgo.MessageCreate, args[]string) {
	defer Discord.client.ChannelMessageDelete(m.ChannelID, m.ID)
	help := "Для канала portal - очистка тестовой базы данных\n"
	switch Discord.GetChName(m.ChannelID) {
		case "portal":
			Discord.ChannelMessage(Discord.GetChID("portal"), fmt.Sprintf("%v очистил(а) базу данных.", m.Author.Username))
			if integrations.CleanDatabase()!=nil{
				Discord.ChannelMessage(Discord.GetChID("portal"), "Не удалось очистить базу")
			}
		default:
			Discord.MessageUser(m.Author.ID, "Запросить команду /clean можно в соответсвующем канале.\n"+help)
			return
	}
}

func deleteobject(m *discordgo.MessageCreate, args[]string) {
	defer Discord.client.ChannelMessageDelete(m.ChannelID, m.ID)
	help := "Для канала portal - удаление лицензионного ключа \n" +
		"Для использования необходимо после команды указать ключ.\n" +
		"Например: __/delete W87J-UUM[DELETED]QSZ-2X36-UNTA__"
	if len(args)==0 || args[0]==""{
		Discord.MessageUser(m.Author.ID, "Нужно что-то указать после команды (один объект)\n"+help)
		return
	}
	switch Discord.GetChName(m.ChannelID) {
		case "portal":
			Discord.ChannelMessage(Discord.GetChID("portal"), fmt.Sprintf("%v удалил(а) лицензионный ключ %v", m.Author.Username, args[0]))
			if integrations.DeleteKey(args[0])!=nil{
				Discord.ChannelMessage(Discord.GetChID("portal"), "Не удалось удалить ключ")
			}
		default:
			Discord.MessageUser(m.Author.ID, "Запросить команду /delete можно в соответсвующем канале.\n"+help)
			return
	}
}

func emergencydroplet(m *discordgo.MessageCreate, args[]string) {
	defer Discord.client.ChannelMessageDelete(m.ChannelID, m.ID)
	help := "Для канала portal - переустановка облачного сервера \n" +
		"Для использования необходимо после команды указать адрес.\n" +
		"Например: __/renew 18.[DELETED].100"
	if len(args)==0 || args[0]==""{
		Discord.MessageUser(m.Author.ID, "Нужно что-то указать после команды (один объект)\n"+help)
		return
	}
	switch Discord.GetChName(m.ChannelID) {
	case "portal":
		Discord.ChannelMessage(Discord.GetChID("portal"), fmt.Sprintf("%v запросил(а) перезапуск сервера %v", m.Author.Username, args[0]))
		integrations.RenewDroplet(args[0])
	default:
		Discord.MessageUser(m.Author.ID, "Запросить команду /renew можно в соответсвующем канале.\n"+help)
		return
	}
}
