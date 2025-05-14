package main

import (
	"fmt"

	"github.com/L0PE/gator/internal/config"
)


func main() {
	configStruct, err := config.Read()
	if err != nil {
		fmt.Printf("Errow durind reading the config: %v", err)
		return
	}

	err = configStruct.SetUser("Liu")
	if err != nil {
		fmt.Printf("Errow durind setting the user: %v", err)
		return
	}
	
	configStruct, err = config.Read()
	if err != nil {
		fmt.Printf("Errow durind reading the config: %v", err)
		return
	}
	fmt.Printf("url: %s\n", configStruct.DB_url)
	fmt.Printf("user: %s\n", configStruct.Current_user_name)
}
