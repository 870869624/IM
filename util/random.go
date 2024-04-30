package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	keys = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nums = "1234567890"
)

func RandomUserString() []string {
	var string []string
	key := strings.Split(keys, "")
	for i := 0; i < 8; i++ {
		k := rand.Intn(46)
		string = append(string, key[k])
	}
	fmt.Println(string)
	return string
}

func RandomNum() []string {
	var num []string
	key := strings.Split(nums, "")
	for i := 0; i < 8; i++ {
		k := rand.Intn(10)
		num = append(num, key[k])
	}
	fmt.Println(num)
	return num
}
