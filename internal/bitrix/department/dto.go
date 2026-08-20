package department

// Структура отдела (с использованием FlexInt)
type BitrixDepartment struct {
	ID     FlexInt     `json:"ID"` // Теперь не упадет ни на "1", ни на 1
	Name   string      `json:"NAME"`
	Sort   FlexInt     `json:"SORT"` // И здесь тоже полная безопасность
	Parent interface{} `json:"PARENT"`
}

// Структура ответа, использует обновленный BitrixDepartment
type BitrixDepartmentResponse struct {
	Result           []BitrixDepartment `json:"result"`
	ErrorType        string             `json:"error"`
	ErrorDescription string             `json:"error_description"`
}
