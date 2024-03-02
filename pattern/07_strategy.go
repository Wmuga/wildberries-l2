package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	Применимость:
		Нужно использовать разные вариации одного алгоритма внутри одного объекта
		Много похожих классов, отличающихся поведением
		Нужно изолировать код, данные и зависимости от других объектов
		Различные варианты алгоритма реализованы в виде условного оператора
	Плюсы:
		Hot swap алгоритмов на лету
		Изолирует код и данные алгоритмов от других объектов
		Уход от наследования к делегированию
		Принцип открытости / закрытости
	Минусы:
		Усложнение дополнительными классами
		Необходимо знать разницу между стратениями, чтобы выбрать
*/

/*
	Предположим задача:
		Есть архиватор, пользователю необходима возможность выбрать алгоритм шифровки zip, rar, 7z
	Context - Archiver, Strategy - CompressionAlgorithm, ConcreteStrategy - ZipCompression, RarCompression, SevenZCompression
*/

type CompressionAlgorithm interface {
	Compress([]byte) []byte
}

type ZipCompression struct{}

func (*ZipCompression) Compress(in []byte) []byte {
	return append(in, []byte(" zipped")...)
}

type RarCompression struct{}

func (*RarCompression) Compress(in []byte) []byte {
	return append(in, []byte(" rarred")...)
}

type SevenZCompression struct{}

func (*SevenZCompression) Compress(in []byte) []byte {
	return append(in, []byte(" 7zipped")...)
}

type Archiver struct {
	Algorithm CompressionAlgorithm
}

func ExampleStrategyPattern() {
	arch := Archiver{}
	arch.Algorithm = &RarCompression{} // и его можно hot swap'пать
	data := []byte("data")
	data = arch.Algorithm.Compress(data)
	fmt.Println(string(data))
}
