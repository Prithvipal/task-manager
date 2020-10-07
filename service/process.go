package service

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// GetChildProcesses returns list of child processes of given parent pid
func GetChildProcesses(pid int) {
	gid := unix.Getgid()
	fmt.Println(gid)
}
