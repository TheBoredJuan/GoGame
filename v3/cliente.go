package cliente

import (
    "bufio"
    "fmt"
    "github.com/reiver/go-oi"
    "github.com/reiver/go-telnet"
    "log"
    "os"
    "os/exec"
)

type caller struct{}

func (c caller) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        oi.LongWrite(w, scanner.Bytes())
        oi.LongWrite(w, []byte(""))
    }
}

func cliente() {
    fmt.Printf("Dial to %s:%d\n", "localhost", 8080)
    err := telnet.DialToAndCall(fmt.Sprintf("%s:%d", "localhost", 8080), caller{})

    if err != nil {
        log.Fatal(err)
    }
}
