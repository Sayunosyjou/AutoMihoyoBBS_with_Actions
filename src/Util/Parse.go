package Util

import "gopkg.in/yaml.v3"

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

func Parse(YamlData string) (ConfigFileStruct, error) {
	var configFileStruct ConfigFileStruct

	err := yaml.Unmarshal([]byte(YamlData), &configFileStruct)
	if err != nil {
		return ConfigFileStruct{}, err
	}

	return configFileStruct, nil
}
