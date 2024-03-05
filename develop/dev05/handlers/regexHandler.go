package handlers

import (
	"github.com/Wmuga/wildberries-l2/develop/dev05/flags"
	"regexp"
)

// regexHandler выполняет поиск по строкам, используя входные данные как паттерн
type regexHandler struct {
	*baseNextHandler
}

func (r *regexHandler) Handle(pattern string, flags flags.Flags, strings []string) ([]HandlerResult, error) {
	if flags.Fixed {
		return r.invokeNext(pattern, flags, strings)
	}

	// Флаг игнорирования регистра
	if flags.IgnoreCase {
		pattern = "(?i)" + pattern
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	// скан строк регуляркой
	res := make([]HandlerResult, len(strings))
	for i := range res {
		res[i].Line = strings[i]
		res[i].Include = re.MatchString(strings[i])
	}

	return res, nil
}

func NewRegexHandler() Handler {
	return &regexHandler{&baseNextHandler{}}
}
