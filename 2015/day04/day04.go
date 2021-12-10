package day04

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

const input = "yzbqklnj"

func Tasks() {
    result := mineAdventCoin(input, 5)

    fmt.Printf("Day 4 Task 1: %d\n", result)
    result2 := mineAdventCoin(input, 6)
    fmt.Printf("Day 4 Task 2: %d\n", result2)
}

func GetMD5Hash(text string) string {
   hash := md5.Sum([]byte(text))
   return hex.EncodeToString(hash[:])
}

func mineAdventCoin(input string, zeros int) int {
    search := strings.Repeat("0", zeros)
    for i := 0; ; i++ {
        text := fmt.Sprintf("%s%d", input, i)
        hash := GetMD5Hash(text)
        if strings.HasPrefix(hash, search) {
            return i 
        }
    }
}
