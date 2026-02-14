package a

// Мини-заглушка под то, что нужно анализатору: zap.L().Info("...")
var zap zapPkg

type zapPkg struct{}

func (zapPkg) L() logger { return logger{} }

type logger struct{}

func (logger) Debug(string, ...any) {}
func (logger) Info(string, ...any)  {}
func (logger) Warn(string, ...any)  {}
func (logger) Error(string, ...any) {}
