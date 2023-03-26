package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "os/exec"
)

type Git struct {
    Init bool
    RemoteAddr string
}

func (g *Git) HandleGit() {
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
            g.Init = true
            cmdStruct := exec.Command("git", "init")    
            out, err := cmdStruct.Output()
            if err != nil {
                log.Fatalln("Error:", err)
            }
            fmt.Println(string(out))
            g.HandlGitRemote()
            return
        case "n":
            g.Init = false
            fmt.Println("No git init")
            return
        default:
            fmt.Println("Key stroke not understood")
            fmt.Println("Please enter either y or n")
            continue
        }
    }
}

func (g *Git) HandlGitRemote() {
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
    // address := string(remoteAddr)
    g.RemoteAddr = string(remoteAddr)

    cmdStruct := exec.Command("git", "remote", "add", "origin", g.RemoteAddr)    
    out, err := cmdStruct.Output()
    if err != nil {
        log.Fatalln("Error:", err)
    }

    fmt.Println(out)
}


