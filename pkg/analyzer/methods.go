package analyzer

const (
	SlogPkgName = "log/slog"
	ZapPkgName  = "go.uber.org/zap"
)

var methods = map[string]map[string]bool{
	SlogPkgName: {
		"Debug":        true,
		"Info":         true,
		"Warn":         true,
		"Error":        true,
		"DebugContext": true,
		"InfoContext":  true,
		"WarnContext":  true,
		"ErrorContext": true,
	},
	ZapPkgName: {
		"Debug":  true,
		"Info":   true,
		"Warn":   true,
		"Error":  true,
		"DPanic": true,
		"Panic":  true,
		"Fatal":  true,
	},
}
