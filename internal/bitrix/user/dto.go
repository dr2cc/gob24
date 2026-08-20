package user

// Структура пользователя
type BitrixUserFields struct {
	Email        string `json:"EMAIL" yaml:"email"`
	Name         string `json:"NAME" yaml:"name"`
	LastName     string `json:"LAST_NAME" yaml:"last_name"`
	WorkPosition string `json:"WORK_POSITION,omitempty" yaml:"work_position"` // Должность
	UFDepartment []int  `json:"UF_DEPARTMENT" yaml:"uf_department"`           // Номер в срезе с ID отделов (например, [1])
}

// Структура для разбора ЛЮБОГО ответа от Битрикс24
type BitrixResponse struct {
	Result           int    `json:"result"`            // Сюда запишется ID, если всё ок
	ErrorType        string `json:"error"`             // Код ошибки (например, ERROR_USER_EMAIL_ALREADY_EXISTS)
	ErrorDescription string `json:"error_description"` // Понятное описание ошибки
}
