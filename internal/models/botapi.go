package models

type ControllerResponce struct {
	Answer    string   `json:"answer"`
	Delay     int      `json:"delay"`
	Keyboard  Keyboard `json:"keyboard"`
	IsKb      bool     `json:"isKb"`
	IsNextMsg bool     `json:"isNextMsg"`
	Id        int      `json:"id"`
}

type Result struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
	// MyChatMember    string  `json:"my_chat_member"`
	// ChatMember      string  `json:"chat_member"`
	// ChatJoinRequest string  `json:"chat_join_request"`
}

type Message struct {
	MessageID                     int                           `json:"message_id"`
	From                          User                          `json:"from"`
	Chat                          Chat                          `json:"chat"`
	Date                          int64                         `json:"date"`
	Text                          string                        `json:"text"`
	Entities                      []Entity                      `json:"entities"`
	MessageAutoDeleteTimerChanged MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	// ProximityAlertTriggered       ProximityAlertTriggered       `json:"proximity_alert_triggered"`
	// VoiceChatScheduled            string                        `json:"voice_chat_scheduled"`
	// VoiceChatStarted              string                        `json:"voice_chat_started"`
	// VoiceChatEnded                string                        `json:"voice_chat_ended"`
	// VoiceChatParticipantsInvited  string                        `json:"voice_chat_participants_invited"`
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
}

type User struct {
	ID int `json:"id"`
	// IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
	// IsPremium    bool   `json:"is_premium"`
}

type Chat struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	Photo     string `json:"photo"`
	Location  string `json:"location"`
}

type Entity struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}
