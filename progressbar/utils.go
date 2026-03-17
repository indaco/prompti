package progressbar

import (
	"charm.land/bubbles/v2/progress"
)

func newProgressBarModel(cfg *Config) progress.Model {
	if cfg.Styles.GradientFrom != nil && cfg.Styles.GradientTo != nil {
		return progress.New(
			progress.WithColors(cfg.Styles.GradientFrom, cfg.Styles.GradientTo),
			progress.WithScaled(true),
			progress.WithWidth(40),
			progress.WithoutPercentage(),
		)
	}
	return progress.New(
		progress.WithDefaultBlend(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
}
