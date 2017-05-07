package main

import (
    "strings"
    "testing"
    "unicode"
)

func TestMakeMessage_mustStartWithHello(t *testing.T) {
    actual := MakeMessage()
    if !strings.HasPrefix(actual, "Hello") {
        t.Errorf("does not start with Hello: %v", actual)
    }
}

func TestMakeMessage_mustEndWithPunctuation(t *testing.T) {
    actual := MakeMessage()
    chars := []rune(actual)
    if !unicode.IsPunct(chars[len(chars) - 1]) {
        t.Errorf("does not end with a punctuation: %v", actual)
    }
}
