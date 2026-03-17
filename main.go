package main
import (
    "fmt"
	"log"
	 "github.com/Wayne_Francis/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = cfg.SetUser("Wayne Francis")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%+v\n", cfg)
}
