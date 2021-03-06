package cli

import (
	"os"
	"path"

	"github.com/urfave/cli"
)

func init() {
	cli.AppHelpTemplate = `Usage: {{.Name}} {{if .Flags}}[OPTIONS] {{end}}COMMAND [arg...]

{{.Usage}}

Version: {{.Version}}{{if or .Author .Email}}

Author:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}
{{if .Flags}}
Options:
  {{range .Flags}}{{.}}
  {{end}}{{end}}
Commands:
  {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}
Run '{{.Name}} COMMAND --help' for more information on a command.
`

	// See https://github.com/urfave/cli/pull/171/files
	cli.CommandHelpTemplate = `{{$DISCOVERY := or (eq .Name "manage") (eq .Name "join") (eq .Name "list")}}Usage: ` + path.Base(os.Args[0]) + ` {{.Name}}{{if .Flags}} [OPTIONS]{{end}} {{if $DISCOVERY}}<discovery>{{end}}

{{.Usage}}{{if $DISCOVERY}}

Arguments:
   <discovery>    discovery service to use [$SWARM_DISCOVERY]
                   * token://<token>
                   * consul://<ip>/<path>
                   * etcd://<ip1>,<ip2>/<path>
                   * file://path/to/file
                   * zk://<ip1>,<ip2>/<path>
                   * [nodes://]<ip1>,<ip2>{{end}}{{if .Flags}}

Options:
   {{range .Flags}}{{.}}
   {{end}}{{if (eq .Name "manage")}}{{printf "\t * swarm.overcommit=0.05\tovercommit to apply on resources"}}
                                    {{printf "\t * swarm.createretry=0\tcontainer create retry count after initial failure"}}
`

}
