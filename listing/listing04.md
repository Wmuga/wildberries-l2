Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Выведет числа от 0 до 9 включительно и fatal'нется с deadlockом
Канал никогда не был закрыт - range все еще ждет данные. Данные никогда не придут - deadlock
```
