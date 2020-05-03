package models

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"

	//	"crypto/sha1"
	//	"database/sql"
	"fmt"
	//pg is used to generate uuid
	"log"
	// _ "github.com/lib/pq"
)

//CreateUUID create a random UUID with from RFC 4122 adapted from http://github.com/nu7hatch/gouuid
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

//Encrypt is a hash plaintext with bcrypt
func Encrypt(plaintext string) (cryptext []byte, err error) {
	password := []byte(plaintext)
	cryptext, err = bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("Cannot generate password encryption", err)
	}
	return
}
