package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what??")
	}
}

func run() {
	// Change first arg from run to child, send to exe for forking
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// duplicate hostname and pid-space
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	must(cmd.Run())
}

func child() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("Shallow")))
	// chroot
	must(syscall.Chroot("root"))
	must(os.Chdir("/"))

	// pid isolation
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(syscall.Mount("TempDir", "tmp", "tmpfs", 0, ""))
	defer func() { syscall.Unmount("tmp", 0) }()
	defer func() { syscall.Unmount("proc", 0) }()

	// run real command
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
