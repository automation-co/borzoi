package lib_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGenerate(t *testing.T) {
	// making playground
	err := os.Mkdir("playground", 0777)
	if err != nil {
		t.Error("Playground couldn't be created")
	}
	initConfigMap := map[string]string{"tools\\borzoi": "https://github.com/automation-co/borzoi", "tools\\husky": "https://github.com/automation-co/husky"}
	jsonString, err := json.Marshal(initConfigMap)
	if err != nil {
		t.Error("Failed to convert test map to JSON")
	}
	err = ioutil.WriteFile("./playground/borzoi.json", jsonString, 0777)
	if err != nil {
		t.Error("Failed to create borzoi.json for Playground.")
	}
	// changing cd to playground
	err = os.Chdir("./playground")
	if err != nil {
		t.Error("Failed to change directory to playground")
	}
	// going to playground and running clone
	app := "borzoi"
	arg0 := "clone"
	cmd := exec.Command(app, arg0)
	err = cmd.Run()
	if err != nil {
		t.Error("Failed to run borzoi clone")
	}
	// deleting borzoi json
	err = os.Remove("borzoi.json")
	if err != nil {
		t.Error("Failed to delete initial borzoi json")
	}
	// generating borzoi json
	app = "borzoi"
	arg0 = "generate"
	cmd = exec.Command(app, arg0)
	err = cmd.Run()
	if err != nil {
		t.Error("Failed to generate from the test config borzoi json")
	}
	// Open our jsonFile
	jsonFile, err := os.Open("borzoi.json")
	if err != nil {
		t.Error("Failed to Unmarshal borzoi.json")
	}
	// reading json file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// closing json file
	jsonFile.Close()
	var result map[string]string
	// converting to map
	json.Unmarshal([]byte(byteValue), &result)
	testresult := cmp.Equal(result, initConfigMap)
	if testresult == false {
		t.Error("Test result and config file do not match!")
	}
	// coming back from playground
	err = os.Chdir("../")
	if err != nil {
		t.Error("Error while exiting directory Playground")
	}

	err = os.RemoveAll("./playground")
	if err != nil {
		t.Error("Error while removing Playground")
	}
}
