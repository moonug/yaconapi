package resources

const OrgResource = "organizations"

type Org struct {
	// Номер ревизии организации.
	// Каждое действие с организацией (добавление сотрудника, создание команды, изменение настроек) увеличивает номер ревизии на единицу.
	Revision uint `json:"revision,omitempty"`
	// Идентификатор организации.
	ID uint `json:"id,omitempty"`
	// Обозначение организации.
	// Допускается использовать только символы латинского алфавита, цифры, знаки подчеркивания и дефиса.
	Label string `json:"label,omitempty"`
	// Объект с информацией о доменах организации.
	Domains struct {
		// Имя основного домена.
		Display string `json:"display,omitempty"`
		// Имя технического домена.
		Master string   `json:"master,omitempty"`
		All    []string `json:"all,omitempty"`
	} `json:"domains,omitempty"`
	// Идентификатор администратора организации.
	AdminUID uint `json:"admin_uid,omitempty"`
	// Основная электронная почта организации.
	Email string `json:"email,omitempty"`
	// Массив объектов с информацией о подключенных сервисах.
	Services []struct {
		/*
			Обозначение сервиса:
			calendar — Календарь;
			disk — Яндекс.Диск;
			mail — Яндекс.Почта;
			portal — Администрирование;
			search — Поиск по организации;
			staff — Люди и команды;
			tracker — Яндекс.Трекер;
			wiki — Яндекс.Вики;
			yamb — Ямб.
		*/
		Slug string `json:"slug,omitempty"`
		// Состояние сервиса:
		// true — готов к работе;
		// false — не готов.
		Ready bool `json:"ready,omitempty"`
	} `json:"services,omitempty"`
	// Предел доступного пространства на Яндекс.Диске.
	DiskLimit uint `json:"disk_limit,omitempty"`
	// Тип подписки на Яндекс.Коннект:
	// free — бесплатная;
	// paid — платная.
	SubscriptionPlan string `json:"subscription_plan,omitempty"`
	// Обозначение страны организации по ISO 3166-1.
	Country string `json:"country,omitempty"`
	// Язык интерфейса для сотрудников организации:
	// ru — русский;
	// en — английский.
	Language string `json:"language,omitempty"`
	//  Название организации.
	Name string `json:"name,omitempty"`
	// Номер факса.
	Fax string `json:"fax,omitempty"`
	// Объем используемого пространства на Яндекс.Диске.
	DiskUsage uint `json:"disk_usage,omitempty"`
	// Номер телефона.
	PhoneNumber string `json:"phone_number,omitempty"`
}

func (o Org) isR() {}

func GetOrgs() (Org, error) {
	return Org{}, nil
}
