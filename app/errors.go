package app

type appErr string

const (
	initErr = appErr("could not initialize")
)
