package resources

const UserResourse = "users"

type User struct {
	// Признак служебных сотрудников-ботов:
	// true — бот;
	// false — человек.
	IsRobot bool `json:"is_robot,omitempty"`
	// Произвольный идентификатор, который вы можете задать при создании сотрудника.
	ExternalId string `json:"external_id,omitempty"`
	// Должность сотрудника
	Position string `json:"position,omitempty"`
	// Массив объектов с информацией об отделах, к которым относится сотрудник (включая вышестоящие отделы).
	Departments []struct {
		// Идентификатор отдела, к которому относится сотрудник.
		ID uint `json:"id,omitempty"`
	} `json:"departments,omitempty"`
	// Идентификатор организации, в которой состоит сотрудник.
	OrgID uint `json:"org_id,omitempty"`
	// Пол сотрудника:
	// male — мужской;
	// female — женский
	// null — пол не указан.
	Gender string `json:"gender,omitempty"`
	// Дата и время создания сотрудника в формате
	// YYYY-MM-DDThh:mm:ss.ssssssZ
	Created string `json:"created,omitempty"`
	// !
	// Объект с информацией об имени сотрудника
	Name struct {
		// Имя сотрудника
		First string `json:"first,omitempty"`
		// Фамилия сотрудника
		Last string `json:"last,omitempty"`
		// Отчество сотрудника (отчество != среднему имени)
		Middle string `json:"middle,omitempty"`
	} `json:"name,omitempty"`
	// Содержимое поля О сотруднике.
	About string `json:"about,omitempty"`
	// Логин сотрудника.
	Nickname string `json:"nickname,omitempty"`
	// !
	// Массив объектов с информацией о командах, в которых состоит сотрудник.
	Groups []struct {
		// Идентификатор команды, в которой состоит сотрудник.
		ID uint `json:"id,omitempty"`
	} `json:"groups,omitempty"`
	// Признак администратора организации:
	// true — администратор;
	// false — рядовой пользователь.
	IsAdmin bool `json:"is_admin,omitempty"`
	// Дата рождения сотрудника в формате
	// YYYY-MM-DD
	Birthday string `json:"birthday,omitempty"`
	// Идентификатор отдела, в котором состоит сотрудник.
	DepartmentID uint `json:"department_id,omitempty"`
	// Основной адрес электронной почты сотрудника.
	Email string `json:"email,omitempty"`
	// Объект с информацией об отделе, в котором состоит сотрудник.
	// По умолчанию содержит только поле id.
	// Список полей можно расширить, передав в запросе параметр
	// fields=department.<имя поля>.
	// Подробнее о полях отделов читайте в разделе Просмотреть параметры отдела.
	Department struct {
		// Идентификатор отдела.
		ID uint `json:"id,omitempty"`
		// Текстовое название отдела, например, «Отдел разработки».
		Name string `json:"name,omitempty"`
		// Описание отдела.
		Description string `json:"description,omitempty"`
		// Идентификатор сотрудника-руководителя отдела.
		HeadID uint `json:"head_id,omitempty"`
		// Имя почтового ящика отдела. Имя может состоять только из символов латинского алфавита, цифр, знаков минус и нижнего подчеркивания.
		// Например, адрес ящика с именем new-department будет new-department@<ваш-домен>.tld.
		Label string `json:"label,omitempty"`
		// Коллективный почтовый адрес отдела.
		Email string `json:"email,omitempty"`
		// Массив объектов с информацией о родительских отделах. Содержит информацию обо всех вышестоящих отделах
		Parents []interface{} `json:"parents,omitempty"`
	} `json:"department,omitempty"`
	// Массив объектов с информацией о контактах сотрудника.
	Contacts []Contact `json:"contacts,omitempty"`
	// Перечень псевдонимов сотрудника.
	Aliases []string `json:"aliases,omitempty"`
	// Идентификатор сотрудника.
	ID uint `json:"id,omitempty"`
	// Признак уволенного сотрудника:
	// true — уволенный;
	// false — действующий.
	IsDismissed bool `json:"is_dismissed,omitempty"`
}

type Contact struct {
	// Значение контакта.
	Value string `json:"value,omitempty"`
	// Тип контакта. Может принимать одно из значений:
	// email;
	// phone_extension;
	// phone;
	// site;
	// icq;
	// twitter;
	// facebook;
	// skype.
	Type string `json:"type,omitempty"`
	// Признак основного контакта:
	// true — основной;
	// false — альтернативный.
	// У сотрудников может быть только один основной контакт каждого типа.
	Main bool `json:"main,omitempty"`
	// Если у сотрудника есть псевдоним, для него автоматически создается контакт типа email:
	// true — контакт создан на основе псевдонима;
	// false — контакт создан вручную.
	Alias bool `json:"alias,omitempty"`
	// Признак автоматически созданного контакта:
	// true — контакт создан автоматически;
	// false — контакт создан вручную.
	Synthetic bool `json:"synthetic,omitempty"`
}

func (u User) isR() {}

type UserRequest struct {
	// Список полей сотрудников, которые вы хотите получить.
	// Подробнее о доступных параметрах сотрудников читайте в разделе Формат ответа.
	// Если параметр fields не указан, ответ содержит полные сведения о сотрудниках.
	Fields string
	// Идентификаторы сотрудников, о которых вы хотите получить информацию.
	ID uint
	// Логины сотрудников, о которых вы хотите получить информацию.
	Nickname string
	// Идентификаторы отдела, информацию о сотрудниках которых вы запрашиваете.
	// Ответ содержит сведения только сотрудниках самого отдела.
	// Сотрудники подотделов в ответе отсутствуют.
	DepartmentID uint
	// Идентификаторы отделов, информацию о сотрудниках которых вы запрашиваете.
	// Ответ содержит сведения о сотрудниках как самого отдела, так и всех его подотделов.
	RecursiveDepartmentID uint
	// Идентификаторы команд, информацию об участниках которых вы запрашиваете.
	// Ответ содержит сведения только об участниках самой команды.
	// Участники вложенных команд в ответе отсутствуют.
	GroupID uint
	// Идентификаторы команд, информацию об участниках которых вы запрашиваете.
	// Ответ содержит сведения об участниках как самой команды, так и всех вложенных команд.
	RecursiveGroupID uint
	// Включать ли в ответ уволенных сотрудников:
	// true — ответ содержит сведения только об уволенных сотрудниках;
	// false (значение по умолчанию) — ответ не содержит сведения об уволенных сотрудниках;
	// ignore — ответ содержит сведения обо всех сотрудниках;
	IsDismissed string
	// Номер страницы ответа. Значение по умолчанию — 1.
	Page uint
	// Число сотрудников на одной странице.
	// Максимальное значение — 1000. Значение по умолчанию — 20.
	PerPage uint
}

/*
400 Bad Request
Один из параметров запроса имеет недопустимое значение или формат данных.

403 Forbidden
У пользователя или приложения нет прав на доступ к ресурсу, запрос отклонен.

500 Internal Server Error
Внутренняя ошибка сервиса. Попробуйте повторно отправить запрос через некоторое время.

503 Service Unavailable
Сервис API временно недоступен.
*/

func GetUser(id uint, fields ...string) User {
	u := User{}
	return u
}

// В случае успешного выполнения запроса API возвращает ответ с кодом 201.
/* Если запрос не был успешно обработан, ответное сообщение содержит информацию о возникших ошибках:

HTTP-код ошибки	Описание ошибки
403 Forbidden
У пользователя или приложения нет прав на доступ к ресурсу, запрос отклонен.

409 Conflict
Запрос не может быть выполнен по причине конфликта имен.

422 Unprocessable Entity
Ошибка валидации, запрос отклонен.

500 Internal Server Error
Внутренняя ошибка сервиса. Попробуйте повторно отправить запрос через некоторое время.

503 Service Unavailable
Сервис API временно недоступен.
*/
func CreateUser(u User) {

}

/*
Тело запроса содержит новые значения параметров сотрудника. Запрос изменит значения только тех параметров, которые явно указаны в теле.

Совет.
Чтобы удалить значение параметра, укажите null в качестве нового значения.
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
func UpdateUser(u User) {

}

/*
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
func AddUserAlias(id uint, alias string) {

}
