package logger

type UseCase interface {
	LogError(error)
	LogInfo(string)
	LogWarning(string)
}
