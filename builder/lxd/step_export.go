package lxd

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type stepExport struct{}

func (s *stepExport) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	config := state.Get("config").(*Config)
	imageFingerprint := state.Get("imageFingerprint").(string)

	ui.Say("Exporting image")
	_, err := LXDCommand("image", "export", imageFingerprint, config.OutputImage)
	if err != nil {
		err := fmt.Errorf("Error exporting container: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	ui.Say("Deleting image")
	_, err = LXDCommand("image", "delete", imageFingerprint)
	if err != nil {
		err := fmt.Errorf("Error deleting image: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *stepExport) Cleanup(state multistep.StateBag) {}
