package main

import (
	"fmt"
	"time"
	"github.com/eiannone/keyboard"
)

func pt(bg time.Time){
	for{	
		fmt.Print("\r", time.Since(bg).String())
	}
}

func main(){
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	fmt.Println("Press enter")
	for {
		_, key, err := keyboard.GetKey() //char, key, err
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEnter {
			break
		}
	}
	bg := time.Now()
	go pt(bg)
	for {
		_, key, err := keyboard.GetKey() //char, key, err
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEnter {
			fmt.Println("\r",time.Since(bg).String())
		}
	}
}