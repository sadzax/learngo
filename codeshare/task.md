//Типизация и интерфейсы в Go

//1. Какой вывод программы? Как исправить, если необходимо?
var data interface{} = 42
data = "dynamic?"
num := data.(int)
fmt.Println(num)


//2. Какой вывод? Как исправить, если необходимо?
package main

import (
"fmt"
)

func main() {
var f interface{} = func() { fmt.Println("test") }
f()
}


//3. Какой вывод? Как исправить, если необходимо?

import "fmt"

func main() {

	err := returnsError()
	if err != nil { 
		fmt.Printf("Unexpected error %t", err==nil)
	}

}

type error interface {
Error() string
}

type MyError struct {
}

func (MyError) Error() string {
return "my error"
}

func returnsError() error {
var p *MyError = nil
return p
}

// Конкуренция



//1. Что происходит и какой вывод программы? Как починить?

package main

import "fmt"

func main() {
fmt.Println("main() started")

    c := make(chan string)
    c <- "John"

    fmt.Println("main() stopped")
}


//2. Что происходит и какой вывод программы? Как улучшить?

package main

import (
"fmt"
"time"
)

func Greet(dataChannel chan string) {
for data := range dataChannel {
fmt.Println("Hello:", data)
}
fmt.Println("Channel is closed. Goodbye!")
}

func main() {
dataChannel := make(chan string)

	go Greet(dataChannel)

	dataChannel <- "Dog"

	dataChannel <- "Cat"

	time.Sleep(1 * time.Second)

}



// ООП

1. Какой вывод?

package main

import "fmt"

type Release interface {
// IsPublic возвращает true, если релиз является публичным
IsPublic() bool

	// Version возвращает версию релиза
	Version() string

	// KubernetesVersion возвращает версию Kubernetes
	KubernetesVersion() string
}

type Version10 struct {
}

func (r Version10) IsPublic() bool { return true }

func (r Version10) Version() string { return "1.0.0" }

func (r Version10) KubernetesVersion() string { return "1.31.3" }

type Version100 struct {
Version10
}

func (r Version100) IsPublic() bool { return false }

func (r Version100) Version() string { return "1.0" }

type Version101 struct {
Version10
}

func (r Version101) Version() string { return "1.0.1" }

func (r Version101) IsPublic() bool { return false }

func (r Version101) KubernetesVersion() string { return "1.32.3" }

func main() {
versions := []Release{
Version100{},
Version101{},
}

	for _, v := range versions {
		fmt.Printf("%T: Version=%s, KubeVersion=%s, IsPublic=%t\n, ",
			v, v.Version(), v.KubernetesVersion(), v.IsPublic())
	}
}


