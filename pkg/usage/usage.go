package usage

import (
	"bytes"
	"fmt"
	"github.com/barchart/common-go/pkg/parameters"
)

type command struct {
	name        string
	description string
	arguments   []string
}

type usage struct {
	appName        string
	appDescription string
	commands       []command
	parameters     map[string]parameters.Parameter
}

var usg usage

func init() {
	usg = usage{
		appDescription: "",
		commands:       make([]command, 0),
		parameters:     make(map[string]parameters.Parameter),
	}
}

func Initialize(name, description string) {
	usg.appName = name
	usg.appDescription = description
}

func AddCommand(name, description string, args ...string) {
	usg.commands = append(usg.commands, command{
		name:        name,
		description: description,
		arguments:   args,
	})
}

func AddParameters(collection map[string]parameters.Parameter) {
	if collection != nil {
		usg.parameters = collection
	}
}

func (u usage) getName() string {
	return fmt.Sprintf("  Application: %v\n", u.appName)
}

func (u usage) getDescription() string {
	return fmt.Sprintf("  Description: %v\n", u.appDescription)
}

func (u usage) getCommands() string {
	str := ""

	if len(usg.commands) != 0 {
		buf := bytes.NewBufferString("  Commands:\n")

		for i, command := range u.commands {
			buf.WriteString(fmt.Sprintf("\t%v", command.name))

			for _, argument := range command.arguments {
				buf.WriteString(fmt.Sprintf(" <%v>", argument))
			}

			if i == len(u.commands)-1 {
				buf.WriteString(fmt.Sprintf("\n\t  %v", command.description))
			} else {
				buf.WriteString(fmt.Sprintf("\n\t  %v\n\n", command.description))
			}
		}

		str = buf.String()
	}

	return str
}

func (u usage) getParameters() string {
	str := ""

	if len(u.parameters) != 0 {
		buf := bytes.NewBufferString("  Parameters:\n")
		index := 0
		for _, param := range u.parameters {
			name := param.Name

			if param.Required {
				name = name + "*"
			}

			buf.WriteString(fmt.Sprintf("\t%v", name))

			if index == len(u.parameters)-1 {
				buf.WriteString(fmt.Sprintf("\n\t  %v (default %v)\n", param.Usage, param.DefaultValue))
			} else {
				buf.WriteString(fmt.Sprintf("\n\t  %v\n\n", param.Usage))
			}

			index++
		}

		str = buf.String()
	}

	return str
}

func GetUsage() string {
	return fmt.Sprintf("Usage: \n%v%v%v%v", usg.getName(), usg.getDescription(), usg.getParameters(), usg.getCommands())
}
