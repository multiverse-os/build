package main

type Account struct {
	Username  string
	Password  string
	PublicKey key.Public
	APIKey    string
}
