package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pw string) (string, error) { b, err := bcrypt.GenerateFromPassword([]byte(pw), 12); return string(b), err }
func CheckPassword(hash, pw string) error { return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw)) }

func main(){h,_:=bcrypt.GenerateFromPassword([]byte("adminadmin"),12);fmt.Println(string(h))}
