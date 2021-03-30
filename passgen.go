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

func PassGen(enNums, enLCLs, enUCLs, enSyms bool, length int) (string, error) {
    if length < 8 {
        return "", fmt.Errorf("Length is less than 8")
    }

    numbers := []byte("0123456789")
    lowerCaseLetters := []byte("abcdefghijklmnopqrstuvwxyz")
    upperCaseLetters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    symbols := []byte("!#$%&*+-<=>?@^_{|}~")

    // Create and seed the generator
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    var numCount int
    var lowCount int
    var uppCount int
    var symCount int

    if enNums {
        numCount = r.Intn(len(numbers))+1
    }
    if enLCLs {
        lowCount = r.Intn(len(lowerCaseLetters))+1
    }
    if enUCLs {
        uppCount = r.Intn(len(upperCaseLetters))+1
    }
    if enSyms {
        symCount = r.Intn(len(symbols))+1
    }
    totalCount := numCount + lowCount + uppCount + symCount

    // normalize counts
    var maxCount int
    var pCount *int
    // numCount
    if numCount > 0 {
        numCount = numCount * length / totalCount + 1
        maxCount = numCount
        pCount = &numCount
    }
    // lowCount
    if lowCount > 0 {
        lowCount = lowCount * length / totalCount + 1
        if lowCount > maxCount {
            maxCount = lowCount
            pCount = &lowCount
        }
    }
    // uppCount
    if uppCount > 0 {
        uppCount = uppCount * length / totalCount + 1
        if uppCount > maxCount {
            maxCount = uppCount
            pCount = &uppCount
        }
    }
    // symCount
    if symCount > 0 {
        symCount = symCount * length / totalCount + 1
        if symCount > maxCount {
            maxCount = symCount
            pCount = &symCount
        }
    }
    if pCount == nil {
        return "", fmt.Errorf("No symbols enabled")
    }

    for (numCount+lowCount+uppCount+symCount) > length {
        *pCount--
    }
    for (numCount+lowCount+uppCount+symCount) < length {
        *pCount++
    }

    var pass []byte
    for i := 0; i < numCount; i++ {
        pass = append(pass, numbers[r.Intn(len(numbers)-1)])
    }
    for i := 0; i < lowCount; i++ {
        pass = append(pass, lowerCaseLetters[r.Intn(len(lowerCaseLetters)-1)])
    }
    for i := 0; i < uppCount; i++ {
        pass = append(pass, upperCaseLetters[r.Intn(len(upperCaseLetters)-1)])
    }
    for i := 0; i < symCount; i++ {
        pass = append(pass, symbols[r.Intn(len(symbols)-1)])
    }
    r.Shuffle(length, func (i, j int) {
        pass[i], pass[j] = pass[j], pass[i]
    })

    return string(pass), nil
}

func main() {
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

    pass, err := PassGen(enNums, enLCLs, enUCLs, enSyms, length)
    if err != nil {
        panic(err)
    }

    fmt.Println(pass)
}
