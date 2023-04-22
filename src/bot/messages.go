package bot

const (
	greetingsMessage = "Привет! Я музыкальный бот на Go\n\n" +
		"Напиши мне название нужного тебе трека через команду /track, а я найду его и отправлю тебе\n" +
		"Например, \"/track Gunna Pushin P\""

	helpMessage = greetingsMessage

	waitingMessage = "Твой трек загружается. Это может занять несколько минут."

	errorMessage = "Произошла неизвестная ошибка. Попробуйте позже."

	emptyTrackMessage = "В команде /track нет названия. Дополните команду названием трека и попробуйте снова."
)
