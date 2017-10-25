package main

import (
    "fmt"
    "os/exec"
    "syscall"
    "time"
)

func main() {
    cmd := exec.Command("./script/test.sh")
    cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true, Pgid: 0}
    cmd.Start()
    done := make(chan struct{})
    go func() {
        err := cmd.Wait()
        status := cmd.ProcessState.Sys().(syscall.WaitStatus)
        exitStatus := status.ExitStatus()
        signaled := status.Signaled()
        signal := status.Signal()
        fmt.Println("Error:", err)
        if signaled {
            fmt.Println("Signal:", signal)
        } else {
            fmt.Println("Status:", exitStatus)
        }
        close(done)
    }()
    pgid, _ := syscall.Getpgid(cmd.Process.Pid)
    fmt.Println(pgid, cmd.Process.Pid)
    go func(pgid int) {
        time.AfterFunc(20 * time.Second, func() {
            exec.Command("touch", "/tmp/test.tmp")
            syscall.Kill(-pgid, syscall.SIGKILL)
        })
    }(pgid)
    <-done
}
