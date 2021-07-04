package job

type DocumentsAwaitingPickupSMSTemplate int

const (
	Day1 DocumentsAwaitingPickupSMSTemplate = iota
	Day2
	Day3
	Day4
	Day5
	Day6
	Day7
	Day8
)

func (t DocumentsAwaitingPickupSMSTemplate) GetTemplate() string {
	return [...]string{
		"День прибытия: {{.StatusUpdateDate}}",
		"{{.RecipientName}}, ваш заказ доставлен: https://putivoditel.store/buy",
		"Ваша посылка (плакат) ожидает на отделении уже третий день",
		"{{.RecipientName}}, для получения своего заказа продиктуйте последнего 4 цифры ттн: {{.ReferenceId}}",
		"Ваш плакат прибыл в отделение 5 дней назад",
		"Успейте забрать свой заказ {{.ReferenceId}}, до автовозврата осталось всего 3 дня 📦♻️",
		"Если у вас не получаеться забрать плакат, пожалуйста, свяжитесь с нами: https://putivoditel.store/call",
		"Напоминаем, что сегодня последний день перед возвратом плаката 📦♻️\nЗаберите посылку или свяжитесь с нами, если не успеваете 📮📲",
	}[t]
}

const DispatchedDocumentSMSTemplate = "Ваш плакат отправлен ({{.ReferenceId}}) 🚚"
