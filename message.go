package hangouts

type Update struct {
	Type      string  `json:"type"`
	EventTime string  `json:"eventTime"`
	Token     string  `json:"token"`
	User      User    `json:"user"`
	Message   Message `json:"message"`
	Space     Space   `json:"space"`
}

type Message struct {
	Name       string `json:"name"`
	CreateTime string `json:"createTime"`
	Text       string `json:"text"`
	Sender     User   `json:"sender"`
	Space      Space  `json:"space"`
}

type User struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	AvatarURL   string `json:"avatarUrl"`
	Email       string `json:"email"`
	Type        string `json:"type"`
}

type Space struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type SimpleMessage struct {
	Text string `json:"text"`
}
