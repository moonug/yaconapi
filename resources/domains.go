package resources

type Domain struct {
	// Указывает ли mx-запись на север mx.yandex.net:
	// true — указывает;
	// false — не указывает.
	MX bool `json:"mx,omitempty"`
	// Признак домена, делегированного на серверы Яндекса:
	// true — делегирован;
	// false — не делегирован.
	Delegated bool `json:"delegated,omitempty"`
	// Признак технического домена (домен вида <ваш домен>.yaconnect.com)
	// true — технический домен;
	// false — обычный домен.
	Tech bool `json:"tech,omitempty"`
	// Доступ к почтовым ящикам домена по протоколу POP:
	// true — включен;
	// false — выключен.
	PopEnabled bool `json:"pop_enabled,omitempty"`
	// Признак основного домена:
	// true — основной домен;
	// false — домен-алиас.
	Master bool `json:"master,omitempty"`
	// Идентификатор почтового ящика по умолчанию.
	PostmasterUID uint `json:"postmaster_uid,omitempty"`
	// Признак подтвержденного домена:
	// true — домен подтвержден;
	// false — домен не подтвержден.
	Owned bool `json:"owned,omitempty"`
	// Страна домена.
	Country string `json:"country,omitempty"`
	// Полное доменное имя.
	Name string `json:"name,omitempty"`
	// Доступ к почтовым ящикам домена по протоколу IMAP:
	// true — включен;
	// false — выключен.
	ImapEnabled bool `json:"imap_enabled,omitempty"`
}

func AddDomain(name string) error {
	return nil
}
