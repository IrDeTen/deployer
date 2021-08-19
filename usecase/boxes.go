package usecase

import (
	"github.com/Delta456/box-cli-maker"
	"github.com/IrDeTen/deployer/models"
)

var (
	dVBox models.VBox = models.VBox{
		Box: box.New(box.Config{
			Px:           1,
			Py:           1,
			Type:         "Round",
			Color:        "HiCyan",
			TitlePos:     "Bottom",
			ContentAlign: "Left",
		}),
		Title:    "Service Deployer (type '--exit' to interrupt deployment)",
		LineSize: 100,
	}

	sVBox models.VBox = models.VBox{
		Box: box.New(box.Config{
			Px:           1,
			Py:           1,
			Type:         "Classic",
			Color:        "HiGreen",
			TitlePos:     "Top",
			ContentAlign: "Left",
		}),
		Title:    "Service stats",
		LineSize: 100,
	}

	errVBox models.VBox = models.VBox{
		Box: box.New(box.Config{
			Px:           1,
			Py:           1,
			Type:         "Bold",
			Color:        "HiRed",
			TitlePos:     "Top",
			ContentAlign: "Left",
		}),
		Title:    "ERROR!!",
		LineSize: 60,
	}
)
