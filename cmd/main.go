package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"unicode"

	"go.einride.tech/can/pkg/socketcan"
)

func main() {
	interfaceFlag := flag.String("interface", "can0", "a native can interface (default=can0)")
	moduleFlag := flag.String("module", "", "the letter of the DOBISS module that you want to filter by")
	outputFlag := flag.Int("output", 0, "the number of the output on your relay module that you want to filter by")

	flag.Parse()

	err := validateFlags(*moduleFlag, *outputFlag)
	if err != nil {
		slog.Error("something went wrong while validating flags",
			"details", err)
		os.Exit(1)
	}

	connection, err := socketcan.DialContext(context.Background(), "can", *interfaceFlag)
	if err != nil {
		slog.Error("something went wrong while connecting to CAN interface",
			"details", err)
		os.Exit(1)
	}

	receiver := socketcan.NewReceiver(connection)

	for receiver.Receive() {
		frame := receiver.Frame()
		frameData := frame.Data[:frame.Length]

		if frameData[0] == 0 &&
			(*moduleFlag == "" || frameData[1] == (*moduleFlag)[0]) &&
			(*outputFlag == 0 || frameData[2] == byte(*outputFlag-1)) {

			var builder strings.Builder

			for _, dataByte := range frameData {
				_, err := fmt.Fprintf(&builder, "%02x", dataByte)
				if err != nil {
					slog.Error("something went wrong while parsing CAN data",
						"details", err)
					os.Exit(1)
				}
			}

			if builder.Len() > 0 {
				slog.Info("received can frame",
					"module", string(frameData[1]),
					"output", frameData[2]+1,
					"id", frame.ID,
					"data", builder.String())
			}
		}
	}
}

func validateFlags(moduleFlag string, outputFlag int) error {
	if len(moduleFlag) > 1 || (len(moduleFlag) == 1 && !unicode.IsLetter(rune(moduleFlag[0]))) {
		return fmt.Errorf("module flag must be a single letter of the alphabet")
	}

	if moduleFlag == "" && outputFlag != 0 {
		return fmt.Errorf("the module flag is required when the output flag is specified")
	}

	if outputFlag != 0 && (outputFlag < 1 || outputFlag > 12) {
		return fmt.Errorf("invalid output flag value (1-12 is allowed)")
	}

	return nil
}
