package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"

	cowsay "github.com/Code-Hex/Neo-cowsay"
)

const cowsayCommand = "cowsay"

func createCowsayCommand() *model.Command {
	return &model.Command{
		Trigger:          cowsayCommand,
		AutoComplete:     true,
		AutoCompleteDesc: "Draw an ascii character saying something",
		AutoCompleteHint: "[character]",
		AutocompleteData: getAutocompleteData(),
	}
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	split := strings.Fields(args.Command)
	command := split[0]
	action := ""
	if len(split) > 1 {
		action = split[1]
	}

	if command != "/"+cowsayCommand {
		return &model.CommandResponse{}, nil
	}

	for _, character := range cowsay.Cows() {
		if character == "sodomized" || character == "head-in" {
			continue
		}
		if action == character {
			message, err := cowsay.Say(
				cowsay.Phrase(strings.TrimPrefix(args.Command, command+" "+action)),
				cowsay.Type(character),
			)
			if err != nil {
				return &model.CommandResponse{}, nil
			}
			p.API.CreatePost(&model.Post{
				Message:   "```\n" + message + "\n```",
				UserId:    args.UserId,
				ChannelId: args.ChannelId,
				ParentId:  args.ParentId,
				RootId:    args.RootId,
			})
			return &model.CommandResponse{}, nil
		}
	}
	return &model.CommandResponse{}, nil
}

func getAutocompleteData() *model.AutocompleteData {
	cowsayCmd := model.NewAutocompleteData("cowsay", "[character] [extra-text]", "Draw an ascii character saying something")

	for _, character := range cowsay.Cows() {
		if character == "sodomized" || character == "head-in" {
			continue
		}
		autocomp := model.NewAutocompleteData(character, "[extra-text]", fmt.Sprintf("Draw a %s saying somethign", character))
		cowsayCmd.AddCommand(autocomp)
	}
	return cowsayCmd
}
