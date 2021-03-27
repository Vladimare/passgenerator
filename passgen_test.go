package main

import (
    "strings"
    "testing"
)

func TestLength(t *testing.T) {
    length := 10
    pass, err := PassGen(true, true, true, true, length)
    if err != nil {
        t.Errorf("Length testing failed\nGot error: %v", err)
    }
    if length != len(pass) {
        t.Errorf("Length doesn't match\nGot: %d\nExpected: %d", len(pass), length)
    }
}

func TestNumbers(t *testing.T) {
    num := false
    pass, err := PassGen(true, true, true, true, 10)
    if err != nil {
        t.Errorf("Numbers testing failed\nGot error: %v", err)
    }
    for _, c := range []byte(pass) {
        if c >= '0' && c <= '9' {
            num = true
            break
        }
    }
    if num == false {
        t.Errorf("Numbers testing failed\nShould get at least 1 number")
    }
    pass, err = PassGen(true, false, false, false, 10)
    if err != nil {
        t.Errorf("Numbers testing failed\nGot error: %v", err)
    }
    num = true
    for _, c := range []byte(pass) {
        if c < '0' || c > '9' {
            num = false
            break
        }
    }
    if num == false {
        t.Errorf("Numbers testing failed\nShould get all numbers")
    }
}

func TestLowerCase(t *testing.T) {
    lower := false
    pass, err := PassGen(true, true, true, true, 10)
    if err != nil {
        t.Errorf("Lower case letters testing failed\nGot error: %v", err)
    }
    for _, c := range []byte(pass) {
        if c >= 'a' && c <= 'z' {
            lower = true
            break
        }
    }
    if lower == false {
        t.Errorf("Lower case letters testing failed\nShould get at least 1 letter")
    }
    pass, err = PassGen(false, true, false, false, 10)
    if err != nil {
        t.Errorf("Lower case letters testing failed\nGot error: %v", err)
    }
    lower = true
    for _, c := range []byte(pass) {
        if c < 'a' || c > 'z' {
            lower = false
            break
        }
    }
    if lower == false {
        t.Errorf("Lower case letters testing failed\nShould get all lower case letters")
    }
}

func TestUpperCase(t *testing.T) {
    upper := false
    pass, err := PassGen(true, true, true, true, 10)
    if err != nil {
        t.Errorf("Upper case letters testing failed\nGot error: %v", err)
    }
    for _, c := range []byte(pass) {
        if c >= 'A' && c <= 'Z' {
            upper = true
            break
        }
    }
    if upper == false {
        t.Errorf("Upper case letters testing failed\nShould get at least 1 letter")
    }
    pass, err = PassGen(false, false, true, false, 10)
    if err != nil {
        t.Errorf("Upper case letters testing failed\nGot error: %v", err)
    }
    upper = true
    for _, c := range []byte(pass) {
        if c < 'A' || c > 'Z' {
            upper = false
            break
        }
    }
    if upper == false {
        t.Errorf("Upper case letters testing failed\nShould get all upper case letters")
    }
}

func TestSymbols(t *testing.T) {
    symbols := string("!#$%&*+-<=>?@^_{|}~")
    symbol := false
    pass, err := PassGen(true, true, true, true, 10)
    if err != nil {
        t.Errorf("Symbols testing failed\nGot error: %v", err)
    }
    for _, c := range []byte(pass) {
        if strings.IndexByte(symbols, c) >= 0 {
            symbol = true
            break
        }
    }
    if symbol == false {
        t.Errorf("Symbols testing failed\nShould get at least 1 letter")
    }
    pass, err = PassGen(false, false, false, true, 10)
    if err != nil {
        t.Errorf("Symbols testing failed\nGot error: %v", err)
    }
    symbol = true
    for _, c := range []byte(pass) {
        if strings.IndexByte(symbols, c) == -1 {
            symbol = false
            break
        }
    }
    if symbol == false {
        t.Errorf("Symbols testing failed\nShould get all symbols")
    }
}

func TestWrongParam(t *testing.T) {
    _, err := PassGen(false, false, false, false, 10)
    if err == nil {
        t.Errorf("Wrong parameters (all false) testing failed\nGot nil error")
    }
    _, err = PassGen(true, true, true, true, 7)
    if err == nil {
        t.Errorf("Wrong parameters (length is less than 8) testing failed\nGot nil error")
    }
}
