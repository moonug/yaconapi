package resources

import "errors"

type r interface {
	isR()
}

type ResourceType int

/*
	Тип ресурса, к которому обращается API:
	users — сотрудники;
	departments — отделы
	groups — команды;
	domains — домены;
	organizations — организации.
*/
const (
	Users ResourceType = iota
	Departments
	Groups
	Domains
	Orgs
)

// Преобразование к строке
func (rt ResourceType) String() string {
	switch rt {
	case Users:
		return "users"
	case Departments:
		return "departments"
	case Groups:
		return "groups"
	case Domains:
		return "domains"
	case Orgs:
		return "organizations"
	}
	return ""
}

/*
200 OK
Запрос с методами GET или PATCH успешно выполнен.

201 Created
В результате выполнения запроса с методом POSTуспешно создан новый объект.

204 No Content
Запрос с методом DELETE успешно выполнен, объект удален.

400 Bad Request
Один из параметров запроса имеет недопустимое значение или формат данных.

403 Forbidden
У пользователя или приложения нет прав на доступ к ресурсу, запрос отклонен.

404 Not Found
Запрашиваемый ресурс не найден.

409 Conflict
Запрос не может быть выполнен по причине конфликта имен.

422 Unprocessable Entity
Ошибка валидации, запрос отклонен.

500 Internal Server Error
Внутренняя ошибка сервиса. Попробуйте повторно отправить запрос через некоторое время.

503 Service Unavailable
Сервис API временно недоступен.
*/

type Response struct {
	// Номер страницы ответа.
	Page uint
	// Общее число найденных сотрудников.
	Total uint
	// Максимальное число сотрудников, информация о которых содержится на одной странице.
	PerPage uint
	// Общее число страниц с результатами выполнения запроса.
	Pages uint
	// Объект, содержащий информацию о страницах ответа.
	Links struct {
		// Адрес следующей страницы ответа.
		Next string
		// Адрес предыдущей страницы ответа.
		Prev string
		// Адрес последней страницы ответа.
		Last string
		// Адрес первой страницы ответа.
		First string
	}
	// уникальный идентификатор ответа
	XRequestID string
	// Link: <https://api.directory.yandex.net/users/?page=3&per_page=10>; rel="next", <https://api.directory.yandex.net/users/?page=20&per_page=10>; rel="last"
	Link   string
	Result []r
}

type Request struct {
	// https://api.directory.yandex.net
	Host string
	// содержит OAuth-токен в формате OAuth <значение токена>
	Authorization string
	//  (только для запросов с методами POST и PATCH) — формат данных запроса.
	// Всегда должен иметь значение application/json; charset=utf-8.
	ContentType string
	//  (только для запросов с методом GET) — ожидаемый формат данных ответа.
	// Всегда должен иметь значение application/json.
	Accept string
	// идентификатор организации.
	// Заголовок необходимо передавать только если запрос выполняется от имени администратора нескольких организаций.
	XOrgID string
	Fileds []string
	/*
		Тип ресурса, к которому обращается API:
		users — сотрудники;
		departments — отделы
		groups — команды;
		domains — домены;
		organizations — организации.
	*/
	Resource ResourceType
	RID      uint
}

func NewRequest(rt ResourceType, token, method string, id uint) (*Request, error) {
	r := Request{
		Host:          ApiHost,
		Authorization: token,
	}
	r.RID = id
	switch method {
	case "GET":
		r.Accept = ApiAccept
	case "POST", "PATCH":
		r.ContentType = ApiContentType
	case "DELETE":
	default:
		return nil, errors.New("Api not support method " + method)
	}
	if rt.String() == "" {
		return nil, errors.New("Api not support " + rt.String() + " resource")
	}
	r.Resource = rt
	return &r, nil
}
