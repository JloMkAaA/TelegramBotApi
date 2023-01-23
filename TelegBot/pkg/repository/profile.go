package repository

type ProfileRepository interface {
	GetProfile(chat_id int64) (Profile, error)
	SavePhoto(chat_id int64, photoId string) error
	SaveAge(chat_id int64, age string) error
	SaveName(chat_id int64, name string) error
	CreateProfile(chat_id int64) (Current, error)
	CheckUserInDB(chat_id int64) (UserId, Current)
	SwitchCurrent(chat_id int64, num int8) error
}

type Profile struct {
	Chat_id int64
	Current int8
	Photo   string
	Name    string
	Age     int
}

type Current struct {
	Id int
}

type UserId struct {
	Id int64
}
