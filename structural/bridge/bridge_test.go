package bridge

import (
	"strings"
	"testing"
)

func TestPrintAPI1(t *testing.T) {
	printer := PrinterImpl1{}

	msg := "Hello"
	err := printer.PrintMessage(msg)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestPrintAPI2NonWriter(t *testing.T) {
	printer := PrinterImpl2{Writer: nil}

	msg := "Hello"
	err := printer.PrintMessage(msg)
	if err != nil {
		expectedErrMsg := "You need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrMsg) {
			t.Errorf("Error message was not correct.\nActual: %s\nExpected: %s\n", err.Error(), expectedErrMsg)
		}
	}
}

func TestPrintAPI2(t *testing.T) {
	writer := TestWriter{}
	printer := PrinterImpl2{Writer: &writer}

	msg := "Hello"
	err := printer.PrintMessage(msg)
	if err != nil {
		t.Errorf("Error using impl2: %s\n", err.Error())
	}

	if writer.Msg != msg {
		t.Errorf("Error message was not correct.\nActual: %s\nExpected: %s\n", writer.Msg, msg)
	}
}

func TestTestWriter(t *testing.T) {
	writer := TestWriter{}

	str := []byte("Hello")
	n, _ := writer.Write(str)
	if writer.Msg != string(str) {
		t.Errorf("Error on io.Writer msg, expected %s, got %s", str, writer.Msg)
	}
	if n != 5 {
		t.Errorf("Expected 5, instead got %d", n)
	}

	str = []byte("")
	_, err := writer.Write(str)
	if err == nil {
		expectedErr := "Content received on Writer was empty"
		t.Errorf("Expected: %s\nActual: %s\n", expectedErr, err.Error())
	}
}

func TestNormalPrintImpl1(t *testing.T) {
	msg := "Hello io.Writer"
	normal := NormalPrinter{
		Msg:     msg,
		Printer: &PrinterImpl1{},
	}

	err := normal.Print()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestNormalPrintImpl2(t *testing.T) {
	msg := "Hello io.Writer"

	normal := NormalPrinter{
		Msg: msg,
		Printer: &PrinterImpl2{
			Writer: &TestWriter{},
		},
	}

	err := normal.Print()
	if err != nil {
		t.Fatal(err.Error())
	}
	if msg != normal.Msg {
		t.Errorf("Expected: %s\nActual: %s\n", msg, normal.Msg)
	}
}

func TestPackPrinterImpl1(t *testing.T) {
	msg := "Hello"
	pack := PackPrinter{
		Msg:     msg,
		Printer: &PrinterImpl1{},
	}

	err := pack.Print()
	if err != nil {
		t.Fatal(err.Error())
	}
	if msg != pack.Msg {
		t.Errorf("Expected: %s\nActual: %s\n", msg, pack.Msg)
	}

}

func TestPackPrinterImpl2(t *testing.T) {
	msg := "Hello io.Writer"
	expectedMsg := "Message from Pack: Hello io.Writer"
	testWriter := TestWriter{}
	pack := PackPrinter{
		Msg: msg,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err := pack.Print()
	if err != nil {
		t.Fatal(err.Error())
	}
	if expectedMsg != testWriter.Msg {
		t.Errorf("Expected: %s\nActual: %s\n", expectedMsg, pack.Msg)
	}

}
