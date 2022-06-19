package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const SWD = "C:/Users/Andrey/Documents/Projects/OOP_course_work_modules/receiver/"

func createFlaskProcess(id string, port string) {
	cmd := exec.Command("./venv/Scripts/flask.exe", "run", "-p", port)
	cmd.Env = os.Environ()
	cmd.Dir = SWD + id
	cmd.Env = append(os.Environ(), "FLASK_APP=script.py")
	cmd.Run()
}

func newModule(mtype string, settings string) string {
	if mtype == "flask" {
		return newFlaskScript(settings)
	}
	return "0"
}

func newFlaskScript(settings string) string {
	var set map[string]interface{}
	json.Unmarshal([]byte(settings), &set)
	port := set["port"].(string)
	id := RandStringRunes(16)
	os.Mkdir(SWD+id, 777)
	cmdvenv := exec.Command(SWD+"flask/venv/Scripts/python.exe", "-m", "venv", "./venv")
	cmdvenv.Env = os.Environ()
	cmdvenv.Dir = SWD + id
	cmdvenv.Run()
	scriptbts, _ := os.ReadFile(SWD + "flask/script.py")
	script := string(scriptbts)
	script = strings.Replace(script, "$", id, 2)
	f, _ := os.Create(SWD + id + "/script.py")
	f.WriteString(script)
	f.Close()
	cmdvenv.Wait()
	cmdpip := exec.Command(SWD+id+"/venv/Scripts/pip.exe", "install", "requests", "flask")
	cmdpip.Env = os.Environ()
	cmdpip.Dir = SWD + id
	cmdpip.Run()
	go createFlaskProcess(id, port)
	return id
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
