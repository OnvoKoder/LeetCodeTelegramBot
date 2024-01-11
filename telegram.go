package main

import (
	"strconv"
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func isCallbackQuery(update *botAPI.Update) bool {
	return update.CallbackQuery != nil && update.CallbackQuery.Message.Text != ""
}

func isStartMessage(update *botAPI.Update) bool {
	return update.Message != nil && update.Message.Text == "/start"
}

func getKeyboardRow(buttonText string, buttonCode string) []botAPI.InlineKeyboardButton {
	return botAPI.NewInlineKeyboardRow(botAPI.NewInlineKeyboardButtonData(buttonText, buttonCode))
}

func renderingMenu(bot *botAPI.BotAPI, chatId int64) {
	msg := botAPI.NewMessage(chatId, "Select action")
	msg.ReplyMarkup = botAPI.NewInlineKeyboardMarkup(
		getKeyboardRow(left+" "+"Information about bot"+" "+right, "info"),
	)
	bot.Send(msg)
}

func updateProccessing(update *botAPI.Update, bot *botAPI.BotAPI) {
	var message string
	choice := update.CallbackQuery.Data
	if choice == "info" {
		message = "My functionality: \n Get username information from leetcode. You only need to send the username in a message and get the result immediately"
	} 
	msg := botAPI.NewMessage(update.CallbackQuery.From.ID, message)
	bot.Send(msg)
}

func startBot() {
	var message string
	bot, err := botAPI.NewBotAPI(token)
	check(err)
	updateConfig := botAPI.NewUpdate(0)
	updateConfig.Timeout = timeout
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if isCallbackQuery(&update) {
			updateProccessing(&update, bot)
		} else {
			if isStartMessage(&update) {
				message = "Hi everyone! \nIt's a telegram bot. You can get information from leetcode by username. If you need more information than click info."
				renderingMenu(bot, update.Message.Chat.ID)
			} else {
				username := getUsersInfo(update.Message.Text)
				if username.MatchedUser.Username == "" {
					message = cancel+"Information: \nusername:  user not found" 
				} else {
					message = "Information: \nusername: " + username.MatchedUser.Username
					message += "\n" + solved + "solved: " + strconv.Itoa(username.MatchedUser.SubmitStats.AcSubmissionNum[0].Count) + " / " +strconv.Itoa(username.AllQuestionsCount[0].Count)
					message += "\n" + easy + "Easy: " + strconv.Itoa(username.MatchedUser.SubmitStats.AcSubmissionNum[1].Count) + " / " + strconv.Itoa(username.AllQuestionsCount[1].Count)
					message += "\n" + middle + "Middle: " + strconv.Itoa(username.MatchedUser.SubmitStats.AcSubmissionNum[2].Count) + " / " + strconv.Itoa(username.AllQuestionsCount[2].Count)
					message += "\n" + hard + "Hard: " + strconv.Itoa(username.MatchedUser.SubmitStats.AcSubmissionNum[3].Count) + " / " + strconv.Itoa(username.AllQuestionsCount[3].Count)
				}
			}
			msg := botAPI.NewMessage(update.Message.Chat.ID, message)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
