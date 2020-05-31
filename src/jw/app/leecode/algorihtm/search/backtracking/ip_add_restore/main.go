package main

import "fileget/util"

func ip_add_restore(input string) (ips []string) {
	ips = make([]string, 0)
	doRestore(0, "", &ips, input)
	return
}

func doRestore(k int, tempAddress string, ips *[]string, s string)  {
	if k == 4 || len(s) == 0 {
		if k == 4 && len(s) == 0 {
			*ips = append(*ips, tempAddress)
		}
	}

	rs := []rune(s)
	for i := 0; i < len(s) && i <= 2; i++ {
		if i > 0 && rs[0] == '0' {
			break
		}
		part := s[:i+1]
		if util.ValueOfString(part) <= 255 {
			if len(tempAddress) != 0 {
				part = "." + part
			}
			tp := tempAddress
			tempAddress += part
			doRestore(k + 1, tempAddress, ips, s[i+1:])
			tempAddress = tp
		}
	}

}

func main() {
	util.Lg.Debugf("output: %v", ip_add_restore("25525511135"))
	util.Lg.Debugf("output: %v", ip_add_restore("1223344"))
}
