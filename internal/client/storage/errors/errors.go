package errors

var (
	ErrLogin              = "Неудачная попытка аутентификации"
	ErrRegistration       = "Неудачная попытка регистрации"
	ErrUserExist          = "Пользователь с таким username зарегистрирован"
	ErrUsernameIncorrect  = "Длинна username должна быть не менее шести символов"
	ErrPasswordIncorrect  = "Пароль должен из минимум 8 букв, минимум 1 цифра, по крайней мере 1 верхний регистр, не менее 1 специального символа (пример: Password-123)"
	ErrPasswordDifferent  = "Пароли не совпали"
	ErrNameEmpty          = "Name не заполнен"
	ErrTextEmpty          = "Text не заполнен"
	ErrTextExist          = "Текст с таким name уже существует"
	ErrCartExist          = "Карта с таким name уже существует"
	ErrPaymentSystemEmpty = "Payment System не заполнен"
	ErrNumberEmpty        = "Number не заполнен"
	ErrHolderEmpty        = "Holder не заполнен"
	ErrEndDateEmpty       = "End date не заполнен"
	ErrEndDataIncorrect   = "End Date некорректный (пример: 01/02/2006)"
	ErrCvcEmpty           = "CVC не заполнен"
	ErrCvcIncorrect       = "CVC некорректный (пример: 123)"
)