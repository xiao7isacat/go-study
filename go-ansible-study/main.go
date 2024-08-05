package main

import (
	"context"
	"fmt"
	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/stdoutcallback/results"
	"time"

	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
)

func main() {
	remoteHost := "10.211.55.7"
	extraVars := make(map[string]interface{})
	extraVars["var1"] = "wangtangyu"
	ansiYamlPath := "test.yaml"
	if err := AnsiRunPlay(remoteHost, extraVars, ansiYamlPath); err != nil {
		fmt.Println(err)
	}
}

func AnsiRunPlay(remoteHost string, extraVars map[string]interface{}, ansiYamlPath string) error {
	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "smart",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: remoteHost + ",",
		ExtraVars: extraVars,
	}

	lplaybook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{ansiYamlPath},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		Exec: execute.NewDefaultExecute(
			execute.WithTransformers(
				results.Prepend("prome-shard"),
			),
		),

		//StdoutCallback: "json",
	}
	fmt.Println(lplaybook.String())

	err := lplaybook.Run(ctx)
	return err
}
