package client

import (
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"time"
)

type Client struct {
	id           int64
	name         string
	secondName   string
	phone        int64
	email        string
	birthday     time.Time
	foto         string
	sex          int8
	club         *club.Club
	city         string
	comment      string
	registration time.Time
	bonusBalance int64
	cardId       int64
}

func NewClient(
	id int64,
	name string,
	secondName string,
	phone int64,
	email string,
	birthday time.Time,
	foto string,
	sex int8,
	club *club.Club,
	city string,
	comment string,
	registration time.Time,
	bonusBalance int64,
	cardId int64) *Client {
	return &Client{
		id:           id,
		name:         name,
		secondName:   secondName,
		phone:        phone,
		email:        email,
		birthday:     birthday,
		foto:         foto,
		sex:          sex,
		club:         club,
		city:         city,
		comment:      comment,
		registration: registration,
		bonusBalance: bonusBalance,
		cardId:       cardId,
	}
}

func (c Client) CardId() int64           { return c.cardId }
func (c Client) Id() int64               { return c.id }
func (c Client) Name() string            { return c.name }
func (c Client) SecondName() string      { return c.secondName }
func (c Client) Phone() int64            { return c.phone }
func (c Client) Email() string           { return c.email }
func (c Client) Birthday() time.Time     { return c.birthday }
func (c Client) Foto() string            { return c.foto }
func (c Client) Sex() int8               { return c.sex }
func (c Client) Club() *club.Club        { return c.club }
func (c Client) City() string            { return c.city }
func (c Client) Comment() string         { return c.comment }
func (c Client) Registration() time.Time { return c.registration }
func (c Client) BonusBalance() int64     { return c.bonusBalance }
