package utils

import "syscall"

func GetUID() int {
	return syscall.Getuid()
}

func GetPID() int {
	return syscall.Getpid()
}
