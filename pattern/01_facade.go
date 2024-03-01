package pattern

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"
)

/*
	Реализовать паттерн «фасад».
	Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// -------

/*
	Применимость: нужен простой интерфейс к сложной структуре классов
	Плюсы: Изоляция от сложных компонентов,
	Минус: Можно создать объект, привязанный ко всем классам программы
*/

/*
	Предположим задачу: необходимо разработать программу извелечения
		и конвертации аудио из видео.

	Пусть имеется класс(структура) VideoConverter, взаимодейтвующий с видео и у него есть функция извлечь аудио.
	Пусть имеется класс(структура) AudioConverter, преобразующий аудио между различными форматами.

	Фасадом будет выступать класс(структура), который предложит пользователю функцию преобразования видео в нужный формат аудио
*/

// Сложная структура программы

var (
	ErrFormatNotFound = errors.New("audio format not found")
)

type VideoConverter struct {
	video []byte
}

func NewVideoConverter(video []byte) *VideoConverter {
	return &VideoConverter{video}
}

func (v *VideoConverter) GetAudio() []byte {
	vid := slices.Clone(v.video)
	return append(vid, []byte(" to raw audio")...)
}

type AudioConverter struct {
	rawAudio []byte
}

func NewAudioConverter(audio []byte) *AudioConverter {
	return &AudioConverter{audio}
}

func (ac *AudioConverter) ToMp3() []byte {
	aud := slices.Clone(ac.rawAudio)
	return append(aud, []byte(" to mp3 format")...)
}

func (ac *AudioConverter) ToWav() []byte {
	aud := slices.Clone(ac.rawAudio)
	return append(aud, []byte(" to wav format")...)
}

// Фасад, прячущий ее

type AudioExtractor struct{}

func (*AudioExtractor) ConvertToAudio(video []byte, audioFormat string) ([]byte, error) {
	vidConv := NewVideoConverter(video)
	audConv := NewAudioConverter(vidConv.GetAudio())

	switch strings.ToLower(audioFormat) {
	case "mp3":
		return audConv.ToMp3(), nil
	case "wav":
		return audConv.ToWav(), nil
	}

	return nil, ErrFormatNotFound
}

func ExampleFacadePattern() {
	// соответсвенно есть видео
	video := []byte("some video bytes")
	// И нужен всего лишь один класс, чтобы достать аудио
	audExtr := &AudioExtractor{}

	res, err := audExtr.ConvertToAudio(video, "mp3")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(res))
}
