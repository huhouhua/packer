package app

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"ruijie.com.cn/devops/packer/pkg/logx"
)

type App struct {
	name        string
	description string
	version     string
	runFunc     RunFunc
}
type RunFunc func() error

type Option func(*App)

func NewApp(name string, opts ...Option) *App {
	a := &App{
		name: name,
	}
	for _, o := range opts {
		o(a)
	}
	return a
}
func WithDescription(desc string) Option {
	return func(a *App) {
		a.description = desc
	}
}
func WithVersion(version string) Option {
	return func(a *App) {
		a.version = version
	}
}
func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.runFunc = run
	}
}

func (a *App) Run() {
	a.Print()
	if err := a.runFunc(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		logx.ErrorWithRed("Failed!")
		os.Exit(1)
	} else {
		logx.SuccessWithGreen("Successful")
		os.Exit(0)
	}
}

func (a *App) Print() {
	logx.BluePrintln("AppName:%s", a.name)
	logx.BluePrintln("Version:%s", a.version)
	logx.BluePrintln("Description:%s", a.description)
}
