package ptime

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	list := []string{"d", "D", "j", "l", "N", "S", "w", "z", "W", "F", "m", "M", "n", "t", ":", "o", "y", "Y", "a", "A", "g", "G", "h", "H", "i", "s", "u", "U", "e", "E", "C", "O", "P", "T"}

	php := "/usr/local/opt/php@7.4/bin/php"
	path, _ := os.Getwd()
	host := "127.0.0.1:22719"
	cmd := exec.Command(php, "-S", host, path+"/gen.php")
	go func() { cmd.Run() }()
	time.Sleep(time.Second)
	log.Println("Server on process: ", cmd.Process.Pid)

	all := strings.Join(list, " ")
	list = append(list, all)

	for i := 0; i < 100; i++ {
		// src := time.Date(2021, 03, 21, 13, 31, 58, 123456, time.UTC)
		src := randate()
		for _, v := range list {
			resp, err := http.Get("http://" + host + "/?format=" + url.QueryEscape(v) + "&timestamp=" + fmt.Sprintf("%d", src.Unix()))
			if err != nil {
				log.Fatal(err)
			}
			bs, _ := ioutil.ReadAll(resp.Body)
			if string(bs) != Format(src, v) {
				t.Error("Format:", v, " PHP: '"+string(bs)+"'", " GoLang: '"+Format(src, v)+"'")
			}
		}
		log.Print(".")
	}
	cmd.Process.Kill()
}

func randate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).UTC()
}
