package Util

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type ConfigFileStruct struct {
	Enable    bool
	Version   int
	Account   account
	Mihoyobbs mihoyobbs
	Games     games
}

type account struct {
	Cookie       string
	Login_ticket string
	Stuid        string
	Stoken       string
}

type mihoyobbs struct {
	Enable             bool
	Checkin            bool
	Checkin_multi      bool
	Checkin_multi_list []int
	Read_posts         bool
	Like_posts         bool
	Cancel_like_posts  bool
	Share_post         bool
}

type games struct {
	Cn struct {
		Enable    bool
		Useragent string
		Genshin   struct {
			Auto_checkin bool
			Black_list   []string
		}
		Hokai2 struct {
			Auto_checkin bool
			Black_list   []string
		}
		Tears_of_themis struct {
			Auto_checkin bool
			Black_list   []string
		}
	}
	Os struct {
		Enable  bool
		Cookie  string
		Genshin struct {
			Auto_checkin bool
			Black_list   []string
		}
	}
}

func ParseConfig(YamlData string) (ConfigFileStruct, error) {
	var configFileStruct ConfigFileStruct

	err := yaml.Unmarshal([]byte(YamlData), &configFileStruct)
	if err != nil {
		return ConfigFileStruct{}, err
	}

	return configFileStruct, nil
}

func ParseCookie() map[string]string {
	var CookieMap interface{}
	var Cookie map[string]string

	CookieJSON := os.Getenv("COOKIE")

	err := json.Unmarshal([]byte(CookieJSON), &CookieMap)
	if err != nil {
		log.Fatal("Parse Cookie Json Error, ", err)
	}

	m := CookieMap.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			Cookie[k] = vv
		}
	}

	return Cookie
}
