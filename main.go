package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "os/exec"
)

func main() {
    // take care of Git logic
    handleGit()

    // create a Go module
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
            out, err := cmdStruct.Output()
            if err != nil {
                log.Fatalln("Error:", err)
            }
            fmt.Println(string(out))
            handlGitRemote()
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

func handlGitRemote() {
    fmt.Println("Would you like to add a remote url? (y/n)")
    
    reader := bufio.NewReader(os.Stdin)
    char, _, err := reader.ReadRune()
    if err != nil {
        log.Fatalln("Following error occurred:", err)
    }

    if string(char) != "y" {
        fmt.Println("y not provided, not adding remote")
        return
    }

    fmt.Println("Please provide the remote address: ")
    reader = bufio.NewReader(os.Stdin)
    remoteAddr, _, err := reader.ReadLine()
    if err != nil {
        log.Fatalln("Following error occurred:", err)
    }

    // add regex and string sanitizing
    address := string(remoteAddr)

    cmdStruct := exec.Command("git", "remote", "add", "origin", address)    
    out, err := cmdStruct.Output()
    if err != nil {
        log.Fatalln("Error:", err)
    }

    fmt.Println(out)
}

func handleGoMod()  {
    fmt.Println("Would you like to create a Go module? (y/n)")
    
    reader := bufio.NewReader(os.Stdin)
    char, _, err := reader.ReadRune()
    if err != nil {
        log.Fatalln("Following error occurred:", err)
    }

    if string(char) != "y" {
        fmt.Println("y not received, no module to be created.")
        return
    }

    fmt.Println("Please provide the name for your module:")
    reader = bufio.NewReader(os.Stdin)
    moduleName, _, err := reader.ReadLine()
    if err != nil {
        log.Fatalln("Following error occurred:", err)
    }

    // TODO: sanitize module name
    cmdStruct := exec.Command("go", "mod", "init", string(moduleName))    
    out, err := cmdStruct.Output()
    if err != nil {
        log.Fatalln("Error:", err)
    }

    fmt.Println(out)
}
