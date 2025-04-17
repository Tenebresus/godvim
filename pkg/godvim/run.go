package godvim

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Run() {

    args := os.Args

    if len(args) < 4 || len(args) > 4 {
        fmt.Println("This command should be run as ./godvim [IP] [port] [vim command]")
    }

    nvim_bin_location := args[1]
    ip := args[2]
    port := args[3]
    vim_commad := args[4]

    log.Println(fmt.Sprintf("Ip: %s, port: %s, vim_command: %s", ip, port, vim_commad))

    writeLog(fmt.Sprintf("Ip: %s, port: %s, vim_command: %s", ip, port, vim_commad))

    // nvim --server localhost:3456 --remote-send '<Esc><Space>d<C-w>w:e /tmp/test.go<CR>'
    exec := exec.Command(nvim_bin_location, "--server", ip + ":" + port, "--remote-send", vim_commad)

    log.Println("Configured Command")
    writeLog("Configured Command")

    err := exec.Run()

    writeLog("executed command")
    
    if err != nil {
        log.Println(err)
        writeLog(err.Error())
    }

}

func writeLog(log_line string) {

    file, err := os.OpenFile("/tmp/godvim_log.txt", os.O_WRONLY, os.ModeAppend)

    if err != nil {
        fmt.Println(err)
    }

    file.WriteString(log_line + "\n")
    file.Close()

}
