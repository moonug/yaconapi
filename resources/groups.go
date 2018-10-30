package resources

type Group struct {
	// Название команды.
	Name string `json:"name,omitempty"`
	//	Адрес почтовой рассылки команды.
	Email string `json:"email,omitempty"`
	// Идентификатор команды.
	ID uint `json:"id,omitempty"`
	// Произвольный идентификатор, который вы можете задать при создании команды.
	ExternalID string `json:"external_id,omitempty"`
	// Массив объектов с информацией об участниках команды.
	Members []struct {
		// Тип участника команды:
		// user — сотрудник;
		// group — команда;
		// department — отдел.
		Type string `json:"type,omitempty"`
		// Объект с информацией об участнике.
		// Для разных типов участников набор полей объекта будет разным.
		Object struct {
		} `json:"object,omitempty"`
	} `json:"members,omitempty"`
	// Имя почтового ящика команды.
	// Имя может состоять только из символов латинского алфавита, цифр, знаков минус и нижнего подчеркивания.
	// Например, адрес ящика с именем new-group будет new-group@<ваш-домен>.tld.
	Label string `json:"label,omitempty"`
	// Дата и время создания команды в формате:
	// YYYY-MM-DDThh:mm:ss.ssssssZ
	Created string `json:"created,omitempty"`
	// Тип команды. Может принимать одно из значениий:
	// generic — любая команда, созданная одним из сотрудников;
	// organization_admin — команда администраторов организации;
	// robots — команда роботов;
	// department_head — команда руководителя отдела. Команды такого типа создаются автоматически для каждого отдела.
	Type string `json:"type,omitempty"`
	// Объект с информацией о сотруднике-создателе команды.
	Author User `json:"author,omitempty"`
	// Описаниие команды.
	Description string `json:"description,omitempty"`
	// Число участников команды.
	MembersCount string `json:"members_count,omitempty"`
	// Идентификатор создателя команды.
	AuthorID uint `json:"author_id,omitempty"`
}

type ObjectUser struct {
	// Перечень псевдонимов сотрудника.
	Aliases []struct {
	} `json:"aliases,omitempty"`
	// Идентификатор сотрудника.
	ID uint `json:"id,omitempty"`
	// Логин сотрудника.
	Nickname string `json:"nickname,omitempty"`
	// Идентификатор отдела, в котором состоит сотрудник.
	DepartmentID uint `json:"department_id,omitempty"`
	// Признак уволенного сотрудника:
	// true — уволенный;
	// false — действующий.
	IsDismissed bool `json:"is_dismissed,omitempty"`
	// Должность сотрудника
	Position string `json:"position,omitempty"`
	// Массив объектов с информацией о командах, в которых состоит сотрудник.
	Groups []struct {
	} `json:"groups,omitempty"`
	// Признак администратора организации:
	// true — администратор;
	// false — рядовой пользователь.
	IsAdmin bool `json:"is_admin,omitempty"`
	// Дата рождения сотрудника в формате
	// YYYY-MM-DD
	Birthday string `json:"birthday,omitempty"`
	// Основной адрес электронной почты сотрудника.
	Email string `json:"email,omitempty"`
	// Произвольный идентификатор, который вы можете задать при создании сотрудника.
	ExternalID string `json:"external_id,omitempty"`
	// Пол сотрудника:
	// male — мужской;
	// female — женский
	// null — пол не указан.
	Gender string `json:"gender,omitempty"`
	// Массив объектов с информацией о контактах сотрудника.
	Contacts []Contact `json:"contacts,omitempty"`
	// 	Объект с информацией об имени сотрудника
	Name struct {
	} `json:"name,omitempty"`
	// Содержимое поля О сотруднике.
	About string `json:"about,omitempty"`
}
