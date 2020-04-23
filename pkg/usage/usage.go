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
	arguments      []string
}

var usg usage

func init() {
	usg = usage{
		appDescription: "",
		commands:       make([]command, 0),
		parameters:     make(map[string]parameters.Parameter),
	}
}

// Initialize adds application name and description
func Initialize(name, description string) {
	usg.appName = name
	usg.appDescription = description
}

// AddCommand adds a command description
func AddCommand(name, description string, args ...string) {
	usg.commands = append(usg.commands, command{
		name:        name,
		description: description,
		arguments:   args,
	})
}

// AddParameters adds parameters description from parameters collection (flags, env, AWS Secrets Manager)
func AddParameters() {
	collection := parameters.GetCollection()
	if collection != nil {
		usg.parameters = collection
	}
}

// AddArguments adds arguments description to run without commands
func AddArguments(args ...string) {
	usg.arguments = args
}

// GetUsage returns a usage string
func GetUsage() string {
	return fmt.Sprintf("Usage: \n%v%v%v%v%v", getName(), getDescription(), getParameters(), getCommands(), getArguments())
}

func getName() string {
	return fmt.Sprintf("  Application: %v\n", usg.appName)
}

func getDescription() string {
	return fmt.Sprintf("  Description: %v\n", usg.appDescription)
}

func getCommands() string {
	str := ""

	if len(usg.commands) != 0 {
		buf := bytes.NewBufferString("  Commands:\n")

		for i, command := range usg.commands {
			buf.WriteString(fmt.Sprintf("\t%v", command.name))

			for _, argument := range command.arguments {
				buf.WriteString(fmt.Sprintf(" <%v>", argument))
			}

			if i == len(usg.commands)-1 {
				buf.WriteString(fmt.Sprintf("\n\t  %v\n", command.description))
			} else {
				buf.WriteString(fmt.Sprintf("\n\t  %v\n\n", command.description))
			}
		}

		str = buf.String()
	}

	return str
}

func getParameters() string {
	str := ""

	if len(usg.parameters) != 0 {
		buf := bytes.NewBufferString("  Parameters:\n")
		index := 0
		for _, param := range usg.parameters {
			name := param.Name

			if param.Required {
				name = name + "*"
			}

			buf.WriteString(fmt.Sprintf("\t%v", name))

			if index == len(usg.parameters)-1 {
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

func getArguments() string {
	str := ""

	if len(usg.arguments) != 0 {
		buf := bytes.NewBufferString("  Arguments:\n")
		index := 0
		for _, arg := range usg.arguments {

			buf.WriteString(fmt.Sprintf("\t<%v>", arg))

			index++
		}

		str = buf.String()
	}

	return str
}
