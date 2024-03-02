package main

import (
	"errors"
	"strings"
	"unicode"
)

/*
	=== Задача на распаковку ===

	Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
		- "a4bc2d5e" => "aaaabccddddde"
		- "abcd" => "abcd"
		- "45" => "" (некорректная строка)
		- "" => ""
	Дополнительное задание: поддержка escape - последовательностей
		- qwe\4\5 => qwe45 (*)
		- qwe\45 => qwe44444 (*)
		- qwe\\5 => qwe\\\\\ (*)

	В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

	Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	ErrWrongStr = errors.New("wrong input string")
)

// Структура декодера
type RLEDecoder struct {
	builder  strings.Builder
	finished bool
	cur      rune
	count    int
	input    []rune
	index    int
	escape   bool
}

// Функция ввода новой строки для обработки
func (rle *RLEDecoder) NewString(input string) {
	// Подготовка к новой строке
	rle.builder.Reset()
	rle.input = []rune(input)
	rle.index = 0
	rle.cur = '\000'
	rle.finished = false
	rle.count = 0
	rle.escape = false
}

// Шаг прохода по входной строке
func (rle *RLEDecoder) Step() (bool, error) {
	// Закончили - нечего вызывать
	if rle.finished {
		return false, nil
	}
	// Продолжаем записывать предыдущий символ, пока не закончилось количество
	if rle.count > 0 {
		if rle.cur == '\000' {
			rle.finished = true
			return false, ErrWrongStr
		}
		rle.count--
		rle.builder.WriteRune(rle.cur)
		return true, nil
	}
	// Сброс до нулевого
	rle.cur = '\000'
	// Если дошли до конца - выходим
	if rle.index == len(rle.input) {
		rle.finished = true
		// Если остался "пустой" \ - ошибка
		if rle.escape {
			return false, ErrWrongStr
		}
		return false, nil
	}
	// Если не было подсчитано количество повторений и символ не был \ - задаем 1цу
	defer func() {
		if rle.count == 0 && !rle.escape {
			rle.count = 1
		}
	}()
	// Читаем следующий символ
	cur := rle.input[rle.index]
	rle.index++
	// Escape последовательность
	if cur == '\\' && !rle.escape {
		// Если первый \ - отметка escape
		rle.escape = true
		return true, nil
	}
	// Если сейчас число
	if unicode.IsDigit(cur) {
		// Если до этого был слеш, то надо записать как есть
		if rle.escape {
			rle.cur = cur
			rle.escape = false
		} else {
			// Иначе цифра - начало числа количества записей
			rle.count = int(cur - '0')
		}
	} else {
		// иначе отправляем символ как есть
		rle.cur = cur
		rle.escape = false
	}
	// Вытаскивание всех последующих цифр в одно число
	for ; rle.index < len(rle.input) && unicode.IsDigit(rle.input[rle.index]); rle.index++ {
		rle.count = rle.count*10 + int(rle.input[rle.index]-'0')
	}

	return true, nil
}

// Получение текущего записываемого символа
func (rle *RLEDecoder) Current() rune {
	return rle.cur
}

// Получение окончательного результата
func (rle *RLEDecoder) String() string {
	if !rle.finished {
		return ""
	}
	return rle.builder.String()
}

func UnpackString(input string) (string, error) {
	rle := &RLEDecoder{}
	rle.NewString(input)
	for {
		step, err := rle.Step()
		if err != nil {
			return "", err
		}
		if !step {
			break
		}
	}
	return rle.String(), nil
}

func main() {

}
