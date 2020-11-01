package golang

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(msg string) error
}

type WriterImpl struct {
	Msg string
}

func (w *WriterImpl) Write(msg []byte) (int, error) {
	n := len(msg)
	if n > 0 {
		w.Msg = string(msg)
		return n, nil
	}
	return 0, errors.New("empty string")
}

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Println(msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("writer not passed")
	}
	_, err := fmt.Fprintln(p.Writer, msg)
	return err
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (p *NormalPrinter) Print() error {
	return p.Printer.PrintMessage(p.Msg)
}

type AnotherPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (p *AnotherPrinter) Print() error {
	return p.Printer.PrintMessage(fmt.Sprintf("this is another printer, %s", p.Msg))
}
