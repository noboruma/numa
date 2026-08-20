package numa

import (
	"syscall"
	"unsafe"

	"github.com/klauspost/cpuid/v2"
	"golang.org/x/sys/unix"
)

var fastway = cpuid.CPU.Supports(cpuid.RDTSCP)

func getCPUFast() (int, int)

// GetCPUAndNode returns the node id and cpu id which current caller running on.
// https://man7.org/linux/man-pages/man2/getcpu.2.html
//
// equal:
//
// if fastway {
// 	call RDTSCP
//  The linux kernel will fill the node cpu id in the private data of each cpu.
//  arch/x86/kernel/vsyscall_64.c@vsyscall_set_cpu
// }
// call vdsoGetCPU
func GetCPUAndNode() (cpu int, node int) {
	if fastway {
		return getCPUFast()
	}

	_, _, errno := syscall.RawSyscall(
		unix.SYS_GETCPU,
		uintptr(unsafe.Pointer(&cpu)),
		uintptr(unsafe.Pointer(&node)),
		0,
	)
	if errno != 0 {
		return 0, 0
	}

	return cpu, node
}
