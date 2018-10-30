package resources

type Department struct {
	// Алиасы почтового ящика отдела.
	Aliases string
	// Текстовое название отдела, например, «Отдел разработки».
	Name string `json:"name,omitempty"`
	// Коллективный почтовый адрес отдела.
	Email string `json:"email,omitempty"`
	// Произвольный идентификатор, который вы можете задать при создании отдела.
	ExternalID string `json:"external_id,omitempty"`
	// Признак удаленного отдела:
	// true — отдел удален;
	// false — отдел действующий.
	Removed bool `json:"removed,omitempty"`
	// Идентификатор отдела.
	ID uint `json:"id,omitempty"`
	// Идентификатор родительского отдела.
	ParentID uint
	// Массив объектов с информацией о родительских отделах. Содержит информацию обо всех вышестоящих отделах
	Parents []struct {
		// Текстовое название отдела, например, «Отдел разработки».
		Name string `json:"name,omitempty"`
		// Коллективный почтовый адрес отдела.
		Email string `json:"email,omitempty"`
		// Произвольный идентификатор, который вы можете задать при создании отдела.
		ExternalID string `json:"external_id,omitempty"`
		// Признак удаленного отдела:
		// true — отдел удален;
		// false — отдел действующий.
		Removed bool `json:"removed,omitempty"`
		// Идентификатор отдела.
		ID uint `json:"id,omitempty"`
		// Идентификатор родительского отдела.
		ParentID uint `json:"parent_id,omitempty"`
		// Имя почтового ящика отдела.
		// Имя может состоять только из символов латинского алфавита, цифр, знаков минус и нижнего подчеркивания.
		// Например, адрес ящика с именем new-department будет new-department@<ваш-домен>.tld
		Label string `json:"label,omitempty"`
		// Дата и время создания отдела в формате:
		// YYYY-MM-DDThh:mm:ss.ssssssZ
		Created string `json:"created,omitempty"`
		// Объект с информацией о непосредственном родителе отдела.
		// Описание отдела.
		Description string `json:"description,omitempty"`
		// Число сотрудников отдела без учета вложенных отделов.
		MembersCount uint `json:"members_count,omitempty"`
		// Идентификатор сотрудника-руководителя отдела.
	} `json:"parents,omitempty"`
	// Имя почтового ящика отдела.
	// Имя может состоять только из символов латинского алфавита, цифр, знаков минус и нижнего подчеркивания.
	// Например, адрес ящика с именем new-department будет new-department@<ваш-домен>.tld
	Label string `json:"label,omitempty"`
	// Дата и время создания отдела в формате:
	// YYYY-MM-DDThh:mm:ss.ssssssZ
	Created string `json:"created,omitempty"`
	// Объект с информацией о непосредственном родителе отдела.
	parent struct {
		// Текстовое название отдела, например, «Отдел разработки».
		Name string `json:"name,omitempty"`
		// Идентификатор отдела.
		ID uint `json:"id,omitempty"`
		// Произвольный идентификатор, который вы можете задать при создании отдела.
		ExternalID string `json:"external_id,omitempty"`
		// Признак удаленного отдела:
		// true — отдел удален;
		// false — отдел действующий.
		Removed bool `json:"removed,omitempty"`
		// Идентификатор родительского отдела.
		ParentID uint `json:"parent_id,omitempty"`
	} `json:"parent,omitempty"`
	// Описание отдела.
	Description string `json:"description,omitempty"`
	// Число сотрудников отдела без учета вложенных отделов.
	MembersCount uint `json:"members_count,omitempty"`
	// Идентификатор сотрудника-руководителя отдела.
	Head uint `json:"head,omitempty"`
	// Идентификатор команды руководителя отдела.
	HeadsGroupID uint
}
