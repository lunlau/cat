package models

type TUser struct {
	Id   int    `xorm:"not null pk autoincr INT"`
	Name string `xorm:"not null default '' VARCHAR(128)"`
}
