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

    ip := args[1]
    port := args[2]
    vim_commad := args[3]

    // nvim --server localhost:3456 --remote-send '<Esc><Space>d<C-w>w:e /tmp/test.go<CR>'
    exec := exec.Command("nvim", "--server", ip + ":" + port, "--remote-send", "'" + vim_commad + "'")
    err := exec.Run()
    
    if err != nil {
        log.Println(err)
    }


}
