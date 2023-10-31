package main

import (
	"fmt"
	"testing"
)

func TestDeclaration(t *testing.T) {
	type NoKTP string

	var noKtpRama NoKTP = "192391293"
	fmt.Println(noKtpRama)
}
