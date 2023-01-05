package entity

type DBConnection struct {
	User string
	Pass string
	Port string //example: :5347
	Host string
	DB   string
}
