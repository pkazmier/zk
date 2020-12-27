package cmd

import "github.com/mickael-menu/zk/core/zk"

type Init struct {
	Directory string `arg optional name:"directory" default:"."`
}

func (cmd *Init) Run() error {
	return zk.Create(cmd.Directory)
}