Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
defer выполняет свою функцию перед возвратом из функции.
Аргументы defer'ной функции расчитываются, когда обратывается сам defer
Выведет 2 и 1 так как во втором случае переменная x - локальная
А в первой функции x - именнованная переменная возврата
```
