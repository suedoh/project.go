package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "os/exec"
)

const (
    projectPath = "~/code/"
)

func main() {
    // take care of Git logic
    // handleGit()
    git := &Git{}
    git.HandleGit()

    // create a Go module
    handleGoMod()
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
