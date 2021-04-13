package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func Test_single_sms(t *testing.T) {

	LoadEnv(".")
	siminn := new(SiminnSMS)
	siminn.Username = os.Getenv("USERNAME")
	siminn.Password = os.Getenv("PASSWORD")
	siminn.URL = "https://vasp.siminn.is/smap/"
	siminn.SendFrom = os.Getenv("FROM")
	err, ok := siminn.SendSMS2("hello from siminn sms service", os.Getenv("NUMBER"))
	if err != nil {
		log.Println(err)
		t.Fatal("Failed fetching token")
	}
	if !ok {
		t.Fatal("Could not send text")
	}
}

func LoadEnv(base string) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
}
