package utils_logger

import (
	"os"

	cli_styles "github.com/Uh-little-less-dum/dev-cli/internal/utils/constants"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var templateAppLogger = log.New(os.Stdout)

func setStylesTemplateAppLoggerStyles() {
	styles := log.DefaultStyles()
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("🔷").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color(cli_styles.UlldBlue))
	templateAppLogger.SetStyles(styles)
}

func LogTemplateAppMsg(msg string) {
	setStylesTemplateAppLoggerStyles()
	templateAppLogger.Info(msg)
}
