package golang

import "fmt"

type LegacyPrinter interface {
	Print(msg string) string
}

type LegacyPrinterImpl struct{}

func (p *LegacyPrinterImpl) Print(msg string) string {
	newMsg := fmt.Sprintf("Legacy printer msg: %s", msg)
	fmt.Println(newMsg)
	return newMsg
}

type ModernPrinter interface {
	PrintStored() string
}

type ModernPrinterImpl struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *ModernPrinterImpl) PrintStored() string {
	newMsg := fmt.Sprintf("Modern printer msg: %s", p.Msg)
	fmt.Println(newMsg)
	return newMsg
}

var printer = ModernPrinterImpl{
	OldPrinter: &LegacyPrinterImpl{},
	Msg:        "test_message",
}

func RunPrinter() {
	printer.OldPrinter.Print("Hello World!")
	printer.PrintStored()
}
