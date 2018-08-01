
package models

import (
  "crypto/rand"
	"fmt"
	"io"
	"time"

	"annie_go/models/mymongo"
  "github.com/astaxie/beego"
  "golang.org/x/crypto/scrypt"
	"gopkg.in/mgo.v2"
)

type User struct {
  ID       string       `bson:"_id"      json:"_id,omitempty"`
  Name     string       `bson:"name"     json:"name,omitempty"`
  Password string       `bson:"password" json:"password,omitempty"`
  Salt     string       `bson:"salt"     json:"salt,omitempty"`
  RegDate  time.Time    `bson:"reg_date" json:"reg_date,omitempty"`
}

const pwHashBytes = 64
/**
 * [generate salt]
 * @type {[type]}
 */
func generateSalt() (salt string, err error) {
  buf := make([]byte, pwHashBytes)
  // a empty buf, all 0
  beego.Info(buf)
  //read a random file to buf
  if _, err := io.ReadFull(rand.Reader, buf); err != nil {
    return "", err
  }
  //a random buf
  beego.Info(buf)
  return fmt.Sprintf("%x", buf), nil
}

func generatePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h), nil
}

func NewUser(r *RegisterForm, t time.Time) (u *User, err error) {
  salt, err := generateSalt()
  if err != nil {
    return nil, err
  }
  //salt is the random buf saved as Hexadecimal
  beego.Info(salt)
  //generate a hash, which length is pwHashBytes
  hash, err := generatePassHash(r.Password, salt)
  if err != nil {
    return nil, err
  }

  user := User{
    ID:       r.Phone,
    Name:     r.Name,
    Password: hash,
    Salt:     salt,
    RegDate:  t}

  return &user, nil
}

func (u *User) Insert() (code int, err error) {
  mConn := mymongo.Conn()
  defer mConn.Close()

  c := mConn.DB("").C("users")
  err = c.Insert(u)

  if err != nil {
    if mgo.IsDup(err) {
      code = ErrDupRows
    } else {
      code = ErrDatabase
    }
  } else {
    code = 0
  }
  return
}
