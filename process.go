package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func ProcessDiceRoll(dice *Dice, sleep int) (string, error) {
	id := generateId(6)
	fmt.Printf("Process #%s - Sleep Config: %d\n", id, sleep)
	time.Sleep(time.Duration(sleep) * time.Second)
	result := dice.Roll()
	fmt.Printf("Process #%s - Sleep Config: %d - Value: %s\n", id, sleep, result)
	return result, nil
}

func generateId(length int) string {
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	for i, b := range randomBytes {
		randomBytes[i] = charset[b%byte(len(charset))]
	}

	return string(randomBytes)
}
