package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var ConntrackCountPaths = []string{
	"/proc/sys/net/netfilter/nf_conntrack_count",
	"/proc/sys/net/ipv4/netfilter/ip_conntrack_count",
}

var ConntrackMaxPaths = []string{
	"/proc/sys/net/nf_conntrack_max",
	"/proc/sys/net/netfilter/nf_conntrack_max",
	"/proc/sys/net/ipv4/ip_conntrack_max",
	"/proc/sys/net/ipv4/netfilter/ip_conntrack_max",
}

func Exists(f string) bool {
	_, err := os.Stat(f)
	return err == nil
}

func FindFile(paths []string) (f string, err error) {
	for _, f := range paths {
		if Exists(f) {
			return f, nil
		}
	}

	return "", fmt.Errorf("Cannot find files %s", paths)
}

func CurrentValue(paths []string) (n uint64, err error) {
	path, err := FindFile(paths)
	if err != nil {
		return 0, err
	}

	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	cnt := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cnt = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	n, err = strconv.ParseUint(cnt, 10, 64)
	return n, nil
}
