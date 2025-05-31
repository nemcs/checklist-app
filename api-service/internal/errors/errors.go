package errors

import "errors"

var (
	ErrURLBuildFailed      = errors.New("не удалось собрать URL")
	ErrDBRequestFailed     = errors.New("ошибка при запросе в db-service")
	ErrDBResponseBadStatus = errors.New("db-service вернул неожиданный статус")
	ErrJSONDecodeFailed    = errors.New("ошибка при декодировании JSON")
	ErrJSONEncodeFailed    = errors.New("ошибка при кодировании JSON")
)
