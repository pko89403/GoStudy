package runner

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// Handle - 배시 쉘 스크립트를 스크립트명과 파라미터들을 넣어서 실행
func Handle(scriptName string, args ...string) {
	fmt.Println("***************")

	args = append(args, scriptName)

	scriptArgsSlice := []string{scriptName}
	scriptArgsSlice = append(scriptArgsSlice, args...)
	// args = strings.Join(args, "-")
	fmt.Println(args)
	fmt.Println("***************")

	cmd := exec.Command("/bin/bash", scriptArgsSlice...)
	stdout, err := cmd.StdoutPipe()
	checkErr(err)
	stderr, err := cmd.StderrPipe()
	checkErr(err)

	err = cmd.Start()
	checkErr(err)

	strOut, err := ioutil.ReadAll(stdout)
	checkErr(err)
	strErr, err := ioutil.ReadAll(stderr)
	checkErr(err)

	err = cmd.Wait()
	checkErr(err)

	fmt.Println("stdout - ", string(strOut))
	fmt.Println("stderr - ", string(strErr))
}
