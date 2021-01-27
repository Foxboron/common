// SPDX-License-Identifier: Apache-2.0

// Copyright 2013-2018 Docker, Inc.

package seccomp

import (
	"golang.org/x/sys/unix"
)

func arches() []Architecture {
	return []Architecture{
		{
			Arch:      ArchX86_64,
			SubArches: []Arch{ArchX86, ArchX32},
		},
		{
			Arch:      ArchAARCH64,
			SubArches: []Arch{ArchARM},
		},
		{
			Arch:      ArchMIPS64,
			SubArches: []Arch{ArchMIPS, ArchMIPS64N32},
		},
		{
			Arch:      ArchMIPS64N32,
			SubArches: []Arch{ArchMIPS, ArchMIPS64},
		},
		{
			Arch:      ArchMIPSEL64,
			SubArches: []Arch{ArchMIPSEL, ArchMIPSEL64N32},
		},
		{
			Arch:      ArchMIPSEL64N32,
			SubArches: []Arch{ArchMIPSEL, ArchMIPSEL64},
		},
		{
			Arch:      ArchS390X,
			SubArches: []Arch{ArchS390},
		},
	}
}

// DefaultProfile defines the allowlist for the default seccomp profile.
func DefaultProfile() *Seccomp {
	einval := uint(unix.EINVAL)

	syscalls := []*Syscall{
		{
			Names: []string{
				"_llseek",
				"_newselect",
				"accept",
				"accept4",
				"access",
				"adjtimex",
				"alarm",
				"bind",
				"brk",
				"capget",
				"capset",
				"chdir",
				"chmod",
				"chown",
				"chown32",
				"clock_adjtime",
				"clock_adjtime64",
				"clock_getres",
				"clock_getres_time64",
				"clock_gettime",
				"clock_gettime64",
				"clock_nanosleep",
				"clock_nanosleep_time64",
				"clone",
				"close",
				"close_range",
				"connect",
				"copy_file_range",
				"creat",
				"dup",
				"dup2",
				"dup3",
				"epoll_create",
				"epoll_create1",
				"epoll_ctl",
				"epoll_ctl_old",
				"epoll_pwait",
				"epoll_wait",
				"epoll_wait_old",
				"eventfd",
				"eventfd2",
				"execve",
				"execveat",
				"exit",
				"exit_group",
				"faccessat",
				"faccessat2",
				"fadvise64",
				"fadvise64_64",
				"fallocate",
				"fanotify_mark",
				"fchdir",
				"fchmod",
				"fchmodat",
				"fchown",
				"fchown32",
				"fchownat",
				"fcntl",
				"fcntl64",
				"fdatasync",
				"fgetxattr",
				"flistxattr",
				"flock",
				"fork",
				"fremovexattr",
				"fsetxattr",
				"fstat",
				"fstat64",
				"fstatat64",
				"fstatfs",
				"fstatfs64",
				"fsync",
				"ftruncate",
				"ftruncate64",
				"futex",
				"futimesat",
				"get_robust_list",
				"get_thread_area",
				"getcpu",
				"getcwd",
				"getdents",
				"getdents64",
				"getegid",
				"getegid32",
				"geteuid",
				"geteuid32",
				"getgid",
				"getgid32",
				"getgroups",
				"getgroups32",
				"getitimer",
				"getpeername",
				"getpgid",
				"getpgrp",
				"getpid",
				"getppid",
				"getpriority",
				"getrandom",
				"getresgid",
				"getresgid32",
				"getresuid",
				"getresuid32",
				"getrlimit",
				"getrusage",
				"getsid",
				"getsockname",
				"getsockopt",
				"gettid",
				"gettimeofday",
				"getuid",
				"getuid32",
				"getxattr",
				"inotify_add_watch",
				"inotify_init",
				"inotify_init1",
				"inotify_rm_watch",
				"io_cancel",
				"io_destroy",
				"io_getevents",
				"io_setup",
				"io_submit",
				"ioctl",
				"ioprio_get",
				"ioprio_set",
				"ipc",
				"keyctl",
				"kill",
				"lchown",
				"lchown32",
				"lgetxattr",
				"link",
				"linkat",
				"listen",
				"listxattr",
				"llistxattr",
				"lremovexattr",
				"lseek",
				"lsetxattr",
				"lstat",
				"lstat64",
				"madvise",
				"memfd_create",
				"mincore",
				"mkdir",
				"mkdirat",
				"mknod",
				"mknodat",
				"mlock",
				"mlock2",
				"mlockall",
				"mmap",
				"mmap2",
				"mount",
				"mprotect",
				"mq_getsetattr",
				"mq_notify",
				"mq_open",
				"mq_timedreceive",
				"mq_timedsend",
				"mq_unlink",
				"mremap",
				"msgctl",
				"msgget",
				"msgrcv",
				"msgsnd",
				"msync",
				"munlock",
				"munlockall",
				"munmap",
				"name_to_handle_at",
				"nanosleep",
				"newfstatat",
				"open",
				"openat",
				"openat2",
				"pause",
				"pidfd_getfd",
				"pidfd_open",
				"pidfd_send_signal",
				"pipe",
				"pipe2",
				"pivot_root",
				"poll",
				"ppoll",
				"ppoll_time64",
				"prctl",
				"pread64",
				"preadv",
				"preadv2",
				"prlimit64",
				"pselect6",
				"pselect6_time64",
				"pwrite64",
				"pwritev",
				"pwritev2",
				"read",
				"readahead",
				"readlink",
				"readlinkat",
				"readv",
				"reboot",
				"recv",
				"recvfrom",
				"recvmmsg",
				"recvmsg",
				"remap_file_pages",
				"removexattr",
				"rename",
				"renameat",
				"renameat2",
				"restart_syscall",
				"rmdir",
				"rt_sigaction",
				"rt_sigpending",
				"rt_sigprocmask",
				"rt_sigqueueinfo",
				"rt_sigreturn",
				"rt_sigsuspend",
				"rt_sigtimedwait",
				"rt_tgsigqueueinfo",
				"sched_get_priority_max",
				"sched_get_priority_min",
				"sched_getaffinity",
				"sched_getattr",
				"sched_getparam",
				"sched_getscheduler",
				"sched_rr_get_interval",
				"sched_setaffinity",
				"sched_setattr",
				"sched_setparam",
				"sched_setscheduler",
				"sched_yield",
				"seccomp",
				"select",
				"semctl",
				"semget",
				"semop",
				"semtimedop",
				"send",
				"sendfile",
				"sendfile64",
				"sendmmsg",
				"sendmsg",
				"sendto",
				"set_robust_list",
				"set_thread_area",
				"set_tid_address",
				"setfsgid",
				"setfsgid32",
				"setfsuid",
				"setfsuid32",
				"setgid",
				"setgid32",
				"setgroups",
				"setgroups32",
				"setitimer",
				"setpgid",
				"setpriority",
				"setregid",
				"setregid32",
				"setresgid",
				"setresgid32",
				"setresuid",
				"setresuid32",
				"setreuid",
				"setreuid32",
				"setrlimit",
				"setsid",
				"setsockopt",
				"setuid",
				"setuid32",
				"setxattr",
				"shmat",
				"shmctl",
				"shmdt",
				"shmget",
				"shutdown",
				"sigaltstack",
				"signalfd",
				"signalfd4",
				"sigreturn",
				"socketcall",
				"socketpair",
				"splice",
				"stat",
				"stat64",
				"statfs",
				"statfs64",
				"statx",
				"symlink",
				"symlinkat",
				"sync",
				"sync_file_range",
				"syncfs",
				"sysinfo",
				"syslog",
				"tee",
				"tgkill",
				"time",
				"timer_create",
				"timer_delete",
				"timer_getoverrun",
				"timer_gettime",
				"timer_gettime64",
				"timer_settime",
				"timerfd_create",
				"timerfd_gettime",
				"timerfd_gettime64",
				"timerfd_settime",
				"timerfd_settime64",
				"times",
				"tkill",
				"truncate",
				"truncate64",
				"ugetrlimit",
				"umask",
				"umount",
				"umount2",
				"uname",
				"unlink",
				"unlinkat",
				"unshare",
				"utime",
				"utimensat",
				"utimensat_time64",
				"utimes",
				"vfork",
				"wait4",
				"waitid",
				"waitpid",
				"write",
				"writev",
			},
			Action: ActAllow,
			Args:   []*Arg{},
		},
		{
			Names:  []string{"personality"},
			Action: ActAllow,
			Args: []*Arg{
				{
					Index: 0,
					Value: 0x0,
					Op:    OpEqualTo,
				},
			},
		},
		{
			Names:  []string{"personality"},
			Action: ActAllow,
			Args: []*Arg{
				{
					Index: 0,
					Value: 0x0008,
					Op:    OpEqualTo,
				},
			},
		},
		{
			Names:  []string{"personality"},
			Action: ActAllow,
			Args: []*Arg{
				{
					Index: 0,
					Value: 0x20000,
					Op:    OpEqualTo,
				},
			},
		},
		{
			Names:  []string{"personality"},
			Action: ActAllow,
			Args: []*Arg{
				{
					Index: 0,
					Value: 0x20008,
					Op:    OpEqualTo,
				},
			},
		},
		{
			Names:  []string{"personality"},
			Action: ActAllow,
			Args: []*Arg{
				{
					Index: 0,
					Value: 0xffffffff,
					Op:    OpEqualTo,
				},
			},
		},
		{
			Names: []string{
				"sync_file_range2",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Arches: []string{"ppc64le"},
			},
		},
		{
			Names: []string{
				"arm_fadvise64_64",
				"arm_sync_file_range",
				"sync_file_range2",
				"breakpoint",
				"cacheflush",
				"set_tls",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Arches: []string{"arm", "arm64"},
			},
		},
		{
			Names: []string{
				"arch_prctl",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Arches: []string{"amd64", "x32"},
			},
		},
		{
			Names: []string{
				"modify_ldt",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Arches: []string{"amd64", "x32", "x86"},
			},
		},
		{
			Names: []string{
				"s390_pci_mmio_read",
				"s390_pci_mmio_write",
				"s390_runtime_instr",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Arches: []string{"s390", "s390x"},
			},
		},
		{
			Names: []string{
				"open_by_handle_at",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Caps: []string{"CAP_DAC_READ_SEARCH"},
			},
		},
		{
			Names: []string{
				"bpf",
				"fanotify_init",
				"lookup_dcookie",
				"perf_event_open",
				"quotactl",
				"setdomainname",
				"sethostname",
				"setns",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Caps: []string{"CAP_SYS_ADMIN"},
			},
		},
		{
			Names: []string{
				"chroot",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Caps: []string{"CAP_SYS_CHROOT"},
			},
		},
		{
			Names: []string{
				"delete_module",
				"init_module",
				"finit_module",
				"query_module",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Caps: []string{"CAP_SYS_MODULE"},
			},
		},
		{
			Names: []string{
				"get_mempolicy",
				"mbind",
				"set_mempolicy",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Caps: []string{"CAP_SYS_NICE"},
			},
		},
		{
			Names: []string{
				"acct",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Caps: []string{"CAP_SYS_PACCT"},
			},
		},
		{
			Names: []string{
				"kcmp",
				"process_vm_readv",
				"process_vm_writev",
				"ptrace",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Caps: []string{"CAP_SYS_PTRACE"},
			},
		},
		{
			Names: []string{
				"iopl",
				"ioperm",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Caps: []string{"CAP_SYS_RAWIO"},
			},
		},
		{
			Names: []string{
				"settimeofday",
				"stime",
				"clock_settime",
				"clock_settime64",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Caps: []string{"CAP_SYS_TIME"},
			},
		},
		{
			Names: []string{
				"vhangup",
			},
			Action: ActAllow,
			Args:   []*Arg{},
			Includes: Filter{
				Caps: []string{"CAP_SYS_TTY_CONFIG"},
			},
		},
		{
			Names: []string{
				"socket",
			},
			Action:   ActErrno,
			ErrnoRet: &einval,
			Args: []*Arg{
				{
					Index: 0,
					Value: unix.AF_NETLINK,
					Op:    OpEqualTo,
				},
				{
					Index: 2,
					Value: unix.NETLINK_AUDIT,
					Op:    OpEqualTo,
				},
			},
			Excludes: Filter{
				Caps: []string{"CAP_AUDIT_WRITE"},
			},
		},
		{
			Names: []string{
				"socket",
			},
			Action: ActAllow,
			Args: []*Arg{
				{
					Index: 2,
					Value: unix.NETLINK_AUDIT,
					Op:    OpNotEqual,
				},
			},
			Excludes: Filter{
				Caps: []string{"CAP_AUDIT_WRITE"},
			},
		},
		{
			Names: []string{
				"socket",
			},
			Action: ActAllow,
			Args: []*Arg{
				{
					Index: 0,
					Value: unix.AF_NETLINK,
					Op:    OpNotEqual,
				},
			},
			Excludes: Filter{
				Caps: []string{"CAP_AUDIT_WRITE"},
			},
		},
		{
			Names: []string{
				"socket",
			},
			Action: ActAllow,
			Args: []*Arg{
				{
					Index: 2,
					Value: unix.NETLINK_AUDIT,
					Op:    OpNotEqual,
				},
			},
			Excludes: Filter{
				Caps: []string{"CAP_AUDIT_WRITE"},
			},
		},
		{
			Names: []string{
				"socket",
			},
			Action: ActAllow,
			Includes: Filter{
				Caps: []string{"CAP_AUDIT_WRITE"},
			},
		},
	}

	return &Seccomp{
		DefaultAction: ActErrno,
		ArchMap:       arches(),
		Syscalls:      syscalls,
	}
}
