package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/aws/aws-lambda-go/lambda"
)

const bufSize = 1024 * 1024

func main() {
	root, ok := os.LookupEnv("LAMBDA_TASK_ROOT")
	if !ok {
		fatal(fmt.Errorf("missing required env 'LAMBDA_TASK_ROOT'"))
	}

	cmd := exec.Command(fmt.Sprintf("%s/wrapper-rs", root))

	i, err := cmd.StdinPipe()
	if err != nil {
		fatal(fmt.Errorf("stdin pipe: %v", err))
	}

	o, err := cmd.StdoutPipe()
	if err != nil {
		fatal(fmt.Errorf("stdout pipe: %v", err))
	}

	err = cmd.Start()
	if err != nil {
		fatal(fmt.Errorf("cmd start: %v", err))
	}

	out := bufio.NewReaderSize(o, bufSize)

	lambda.Start(func(line json.RawMessage) (json.RawMessage, error) {
		_, err := i.Write(append(line, '\n'))
		if err != nil {
			return nil, err
		}

		rsp, err := out.ReadBytes('\n')
		if err != nil {
			return nil, err
		}

		return []byte(fmt.Sprintf(`{"success": true, "message": %s}`, rsp)), nil
	})
}

func fatal(err error) {
	fmt.Println(err)
	os.Exit(1)
}
