package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/exec"
	"src/Util"
)

const DefaultConfigPath = "./DefaultExampleConfig.yaml"
const AutoMihoyoBBSConfigPath = "./mihoyo/config/config.yaml"
const UserConfigFilePath = "./data/config.automihoyobbs"
const GitUserName = "github-actions"
const GitUserEmail = "github-actions@github.com"

func main() {
	if os.Args[1] == "pushuserdata" {
		// Push User Data
		SetUserName := exec.Command("git", "config", "user.name", GitUserName)
		err := SetUserName.Run()
		if err != nil {
			log.Fatal("Set Git User Name Error, ", err)
		}

		SetUserEmail := exec.Command("git", "config", "user.email", GitUserEmail)
		err = SetUserEmail.Run()
		if err != nil {
			log.Fatal("Set Git User E-Mail Error, ", err)
		}

		AddUserData := exec.Command("git", "add", UserConfigFilePath)
		err = AddUserData.Run()
		if err != nil {
			log.Fatal("Add User Config File Error, ", err)
		}

		AddCommit := exec.Command("git", "commit", "-m", "Update User Data")
		err = AddCommit.Run()
		if err != nil {
			log.Fatal("Commit Error, ", err)
		}

		Push := exec.Command("git", "push")
		err = Push.Run()
		if err != nil {
			log.Fatal("Push Error, ", err)
		}

		os.Exit(0)
	}

	AES := Util.NewAES()
	AES.SetNonce(
		Util.GetSha256(
			Util.GetSha512(
				Util.GetSha512(
					Util.GetSha512(
						[]byte(os.Args[1])))))[:11])
	AES.SetKey(
		Util.GetSha256(
			Util.GetSha512(
				Util.GetSha512(
					Util.GetSha512(
						[]byte(os.Args[1]))))))
	AES.SetToken(
		Util.GetSha512(
			Util.GetSha512(
				Util.GetSha512(
					[]byte(os.Args[1])))))

	_, err := os.Stat(UserConfigFilePath)
	if os.IsExist(err) {
		ConfigFile, err := AES.DecryptFromBase64(ReadConfig(UserConfigFilePath))
		if err != nil {
			log.Fatal("Decrypt User Config Error, ", err)
		}

		WriteConfig(AutoMihoyoBBSConfigPath, ConfigFile)

		os.Exit(0)
	} else {
		Config := ReadConfig(DefaultConfigPath)
		ConfigStruct, err := Util.Parse(Config)
		if err != nil {
			log.Fatal("Parse File Error, ", err)
		}
		fmt.Println(ConfigStruct)
		ConfigStruct.Account.Cookie = os.Args[1]
		str, _ := yaml.Marshal(ConfigStruct)
		WriteConfig(AutoMihoyoBBSConfigPath, str)

		Encrypted, err := AES.EncryptToBase64(str)
		if err != nil {
			log.Fatal("Encrypt User Data Error, ", err)
		}

		WriteConfig(UserConfigFilePath, []byte(Encrypted))

		os.Exit(0)
	}
}

func ReadConfig(Path string) string {
	Config, err := os.ReadFile(Path)
	if err != nil {
		log.Fatal("Read Config File Error, ", err)
	}
	return string(Config)
}

func WriteConfig(Path string, Config []byte) {
	ConfigFile, err := os.OpenFile(Path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Open Config File Error, ", err)
	}

	defer ConfigFile.Close()

	write := bufio.NewWriter(ConfigFile)
	for i := 0; i < 5; i++ {
		_, err := write.Write(Config)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = write.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
