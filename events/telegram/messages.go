package telegram

const msgHelp = `/help - чтобы помочь. Сохранять и отправлять ссылки. Чтобы сохранить ссылку - просто отправьте её в чат.
Чтобы получить рандомную ссылку - /rnd
Если вы получили рандомную ссылку, она будет удалена`

const msgHello = `Привет! ` + msgHelp

const (
	msgUnknownCommand = "Неизвестная команда"
	msgNoSavedPages   = "Нет сохранённых страниц"
	msgSaved          = "Сохранено"
	msgAlreadyExists  = "Такая страница уже сохранена"
)
