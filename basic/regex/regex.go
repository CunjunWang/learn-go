package main

import (
	"fmt"
	"regexp"
)

const text = `My email is 
duckwcj@gmail.com
email is dd@qq.com   
email2 is    kkk@163.com  		
email3 is   pp@test.com.ca`

func main() {
	// re, err := regexp.Compile("duckwcj@gmail.com")
	//re := regexp.MustCompile("duckwcj@gmail.com")
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	// match := re.FindString(text)
	// match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}

}
