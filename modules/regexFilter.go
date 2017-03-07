package modules

import "regexp"

func FilterDeviceName(name string)(match bool){
	match, _ = regexp.MatchString("([a-zA-Z]+[0-9]+)", name)
	return match
}
