package main

import (
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"
)

// defaults
const (
    defLength = 16
    defEnNums = true
    defEnLCLs = true
    defEnUCLs = true
    defEnSyms = true
)

func printUsage() {
    fmt.Printf("%s", `passgen [N] [opt]

    N                Password length (default is 16)
    --numbers, -n    Enable numbers (enabled by default)
    --lower, -l      Enable lower case letters (enabled by default)
    --upper, -u      Enable upper case letters (enabled by default)
    --sumbols, -s    Enable special characters (enabled by default)
    --help, -h       Print usage

    If any option is specified then others should be enabled.

    Usage example:
        passgen
        passgen 12
        passgen 10 -nlu
`)
}

func main() {
    numbers := []byte("0123456789")
    lowerCaseLetters := []byte("abcdefghijklmnopqrstuvwxyz")
    upperCaseLetters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    symbols := []byte("!#$%&*+-<=>?@^_{|}~")

    enOptions := false

    length := defLength
    enNums := defEnNums
    enLCLs := defEnLCLs
    enUCLs := defEnUCLs
    enSyms := defEnSyms

    if len(os.Args) > 1 {
        for _, arg := range os.Args[1:] {
            if arg[0] == '-' {
                if arg[1] == 'h' || arg[2:] == "help" {
                    printUsage()
                    return
                }
                if !enOptions {
                    enOptions = true
                    enNums = false
                    enLCLs = false
                    enUCLs = false
                    enSyms = false
                }

                arg = arg[1:]
                // Check long options
                if arg[0] == '-' {
                    switch arg[1:] {
                    case "numbers":
                        enNums = true
                    case "lower":
                        enLCLs = true
                    case "upper":
                        enUCLs = true
                    case "symbols":
                        enSyms = true
                    default:
                        printUsage()
                        panic("Unknown option")
                    }
                } else {
                    for i := 0; i < len(arg); i++ {
                        switch arg[i] {
                        case 'n':
                            enNums = true
                        case 'l':
                            enLCLs = true
                        case 'u':
                            enUCLs = true
                        case 's':
                            enSyms = true
                        default:
                            printUsage()
                            panic("Unknown option")
                        }
                    }
                }
            } else {
                l, err := strconv.Atoi(arg)
                if err != nil {
                    printUsage()
                    panic("The argument should be an option or length")
                }
                if l < 8 {
                    printUsage()
                    panic("Password length should be greater than 8")
                }
                length = l
            }
        }
    }

    var passString []byte

    if enNums {
        passString = append(passString, numbers...)
    }
    if enLCLs {
        passString = append(passString, lowerCaseLetters...)
    }
    if enUCLs {
        passString = append(passString, upperCaseLetters...)
    }
    if enSyms {
        passString = append(passString, symbols...)
    }


    // Create and seed the generator
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    var pass []byte
    //r.Shuffle(len(pass), func(i, j int) {
    //    pass[i], pass[j] = pass[j], pass[i]
    //})
    for i := 0; i < length; i++ {
        pass = append(pass, passString[r.Intn(len(passString)-1)])
    }
    fmt.Printf("%s\n", pass)
}
