package adapter

import (
	"fmt"
	"testing"
)

func TestPrinterAdapter(t *testing.T) {
	msg := "test"

	adapterPrinter := Adapter{OldPrinter: &LegacyPrinter{}, Msg: msg}
	newMsg := adapterPrinter.PrintStored()

	oldMsg := fmt.Sprintf("Legacy Printer: Adapter: %s", msg)
	if newMsg != oldMsg {
		t.Errorf("Expected same messages, instead got \n%s\n%s\n", newMsg, oldMsg)
	}
}

func TestNilLegacyAdapter(t *testing.T) {
	msg := "test"

	adapterPrinter := Adapter{OldPrinter: nil, Msg: msg}
	newMsg := adapterPrinter.PrintStored()

	if newMsg != msg {
		t.Errorf("Expected same message, instead got \n%s\n%s\n", msg, newMsg)
	}
}
