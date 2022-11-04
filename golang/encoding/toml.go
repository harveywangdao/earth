package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type People struct {
	Name string
	Age  int
}

func do1() {
	p := People{}
	data := []byte(`
Name = "xiaoming"
Age = 12
Name = "xiaoming2"
    `)
	err := toml.Unmarshal(data, &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}

/*
[Notify]
  AccessKey = "123"
  EnableV2Notification = true
  [Notify.TemplatesV2]
    file_comment_replied = "53"
    file_commented = "54"
    mentioned_in_doc = "55"
*/
type conf struct {
	Notify *NotifyConfig
}

type NotifyConfig struct {
	AccessKey            string
	EnableV2Notification bool
	TemplatesV2          map[string]string
}

func do2() {
	n := &NotifyConfig{
		AccessKey:            "123",
		EnableV2Notification: true,
	}
	t := make(map[string]string)
	t["file_comment_replied"] = "53"
	t["file_commented"] = "54"
	t["mentioned_in_doc"] = "55"
	n.TemplatesV2 = t

	c := &conf{
		Notify: n,
	}

	if err := toml.NewEncoder(os.Stdout).Encode(c); err != nil {
		fmt.Println("Error encoding TOML: %s", err)
	}
}

func main() {
	do2()
}
