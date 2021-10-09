package main

import (
	"encoding/json"
)
type User struct {
    Name string `json:"Name"`
    Email string `json:"Email"`
    Pass String  `json:"Password"`
}
type Post struct {
	Title string
	Caption string
	id string
}
type Password struct{
	pass char
	}
func main() {
    fmt.Println("Encryption Program v0.01")

    text := []byte("My Super Secret Code Stuff")
    key := []byte("passphrasewhichneedstobe32bytes!")

    c, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err)
 
