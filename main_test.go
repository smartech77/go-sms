package siminn

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func Test_single_sms(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), 1500*time.Millisecond)

	LoadEnv(".")

	siminn := new(SiminnSMS)
	siminn.Username = os.Getenv("USERNAME")
	siminn.Password = os.Getenv("PASSWORD")
	siminn.URL = "https://vasp.siminn.is/smap/"
	siminn.SendFrom = os.Getenv("FROM")

	err, ok, code := siminn.SendSMS(ctx, "hello from siminn sms service", os.Getenv("NUMBER"))

	if code != 200 || err != nil || !ok {
		t.Fatal("Failed to send text", code, err, ok)
	}
}

func LoadEnv(base string) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
}
