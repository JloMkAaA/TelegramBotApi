package sqlite

import (
	"DotaFind/pkg/repository"
	"database/sql"
	"fmt"
	"log"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewStorageRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) GetProfile(chat_id int64) (repository.Profile, error) {
	var profile repository.Profile
	q := "SELECT * FROM profile WHERE chat_id = ?"

	rows, err := r.db.Query(q, chat_id)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&profile.Chat_id, &profile.Current, &profile.Photo, &profile.Name, &profile.Age)
		if err != nil {
			log.Fatal(err)
		}
	}

	return profile, nil
}

func (r *ProfileRepository) SavePhoto(chat_id int64, photoId string) error {
	q := "UPDATE profile SET photo = $1 WHERE chat_id = $2"

	_, err := r.db.Exec(q, photoId, chat_id)
	if err != nil {
		return fmt.Errorf("can't insert into table: %w", err)
	}

	return nil
}

func (r *ProfileRepository) SaveAge(chat_id int64, age string) error {
	q := "UPDATE profile SET age = $1 WHERE chat_id = $2"

	_, err := r.db.Exec(q, age, chat_id)
	if err != nil {
		return fmt.Errorf("can't insert into table: %w", err)
	}

	return nil
}

func (r *ProfileRepository) SaveName(chat_id int64, name string) error {
	q := "UPDATE profile SET name = $1 WHERE chat_id = $2"

	_, err := r.db.Exec(q, name, chat_id)
	if err != nil {
		return fmt.Errorf("can't insert into table: %w", err)
	}

	return nil
}

func (r *ProfileRepository) CreateProfile(chat_id int64) (repository.Current, error) {
	var current repository.Current
	current.Id = 1
	q := `INSERT INTO profile (chat_id, current) VALUES (?,?) `

	_, err := r.db.Exec(q, chat_id, current.Id)

	if err != nil {
		log.Fatal(err)
	}

	return current, nil
}

func (r *ProfileRepository) CheckUserInDB(chat_id int64) (repository.UserId, repository.Current) {
	var userId repository.UserId
	var curr repository.Current

	//err := r.db.QueryRow(q, chat_id).Scan(&userId.Id)

	q := "SELECT chat_id, current FROM profile WHERE chat_id = ?"

	rows, err := r.db.Query(q, chat_id)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&userId.Id, &curr.Id)
		if err != nil {
			log.Fatal(err)
		}
	}

	return userId, curr
}

func (r *ProfileRepository) SwitchCurrent(chat_id int64, num int8) error {
	q := "UPDATE profile SET current = $1 WHERE chat_id = $2"

	_, err := r.db.Exec(q, num, chat_id)

	return err
}

func (r *ProfileRepository) Init() error {
	q := `CREATE TABLE IF NOT EXISTS profile (chat_id int, current int, photo string, name string, age string)`

	_, err := r.db.Exec(q)
	if err != nil {
		return fmt.Errorf("Не удалось создать таблицу: %w", err)
	}

	return nil
}
