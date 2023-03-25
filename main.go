package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "os/exec"
)

func main() {
    handleGit()
}

func handleGit() {
    fmt.Println("Initialize a Git repo? (y/n)")

    reader := bufio.NewReader(os.Stdin)
    char, _, err := reader.ReadRune()
    if err != nil {
        log.Fatalln("Following error occurred:", err)
    }

    gitInput := string(char)
    for {
        switch gitInput {
        case "y":
            cmdStruct := exec.Command("git", "init")    
            out,err := cmdStruct.Output()
            if err != nil {
                log.Fatalln("Error:", err)
            }
            fmt.Println(string(out))
            return
        case "n":
            fmt.Println("No git init")
            return
        default:
            fmt.Println("Key stroke not understood")
            fmt.Println("Please enter either y or n")
            continue
        }
    }
}
