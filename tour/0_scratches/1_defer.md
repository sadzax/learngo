```
package main  
  
import "fmt"  
  
func main() {  
    fmt.Println("counting")  
  
    for i := 0; i < 10; i++ {  
       defer fmt.Println(i)  
    }  
  
    fmt.Println("done")  
}
```

В Go ключевое слово **defer** используется для отложенного выполнения функции. Это значит, что функция, объявленная с помощью defer, будет вызвана **только после завершения окружающей функции** (той, в которой находится defer).