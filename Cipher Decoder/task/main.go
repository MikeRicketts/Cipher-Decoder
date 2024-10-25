package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func cCEncrypt(text string, shift int) string {
	shift = shift % 26
	var result strings.Builder
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			result.WriteRune('a' + (char-'a'+rune(shift))%26)
		} else if char >= 'A' && char <= 'Z' {
			result.WriteRune('A' + (char-'A'+rune(shift))%26)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func caesarCipherDecrypt(text string, shift int) string {
	return cCEncrypt(text, 26-shift)
}

func main() {
	var g, p, A int

	fmt.Scanf("g is %d and p is %d", &g, &p)
	fmt.Println("OK")

	b := 7
	B := new(big.Int).Exp(big.NewInt(int64(g)), big.NewInt(int64(b)), big.NewInt(int64(p)))

	fmt.Scanf("A is %d", &A)

	S := new(big.Int).Exp(big.NewInt(int64(A)), big.NewInt(int64(b)), big.NewInt(int64(p))).Int64()

	fmt.Printf("B is %d\n", B)

	encryptedMessage := cCEncrypt("Will you marry me?", int(S)%26)
	fmt.Println(encryptedMessage)

	reader := bufio.NewReader(os.Stdin)
	aliceResponse, _ := reader.ReadString('\n')
	aliceResponse = strings.TrimSpace(aliceResponse)

	decryptedResponse := caesarCipherDecrypt(aliceResponse, int(S)%26)

	if decryptedResponse == "Yeah, okay!" {
		fmt.Println(cCEncrypt("Great!", int(S)%26))
	} else if decryptedResponse == "Let's be friends." {
		fmt.Println(cCEncrypt("What a pity!", int(S)%26))
	}
}
