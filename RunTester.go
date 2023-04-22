package Tests

import (
	"fmt"
	"os/exec"
)

func RunTester(userFolder, fileName, url, user string) {

	avaiable_port := "12344"

	bash := "./create_sandbox.sh"
	cmd, err := exec.Command(bash, userFolder, fileName, url, user).Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output := string(cmd)
	fmt.Printf(output)

	start_docker := "./start_docker.sh"
	_, err = exec.Command(start_docker, userFolder, user, fileName).Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}

}
