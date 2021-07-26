package main

import "fmt"

type Coder interface {
	whatAreYouDoing()
}

type BackendDeveloper struct {
}

func (b BackendDeveloper) whatAreYouDoing() {
	fmt.Println("Coding bugs on System's Backends.")
}

type FrontendDeveloper struct {
}

func (c FrontendDeveloper) whatAreYouDoing() {
	fmt.Println("Coding bugs on System's Frontends.")
}

func main() {

	backDev := BackendDeveloper{}

	frontDev := FrontendDeveloper{}

	backDev.whatAreYouDoing()
	frontDev.whatAreYouDoing()

}
