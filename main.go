package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func execute(config Config) {
	isNoWrite := len(os.Args) > 1 && os.Args[1] == "no-write";
	next_index := 0;

	if config.current != "" {
		for i := range len(config.papers) {
			if strings.HasPrefix(config.papers[i], config.current) {
				if isNoWrite {
					next_index = i;
				} else {
					next_index = i + 1;
					if next_index >= len(config.papers) {
						next_index = 0
					}
				}
				break;
			}
		}
	}

	text := strings.Split(config.papers[next_index], " ");
	matugen_args := []string{"image"}

	for i := range len(text) {
		matugen_args = append(matugen_args, text[i])
	}

	matugen := exec.Command("matugen", matugen_args...);
	log.Println("Executing", matugen.Args);
	if err := matugen.Run(); err == nil {
		matugen.Process.Release()
	}

	if len(os.Args) <= 1 || isNoWrite {
		config.write_new_bg(text[0])
	}
}

func main() {
	config := parse_all()
	execute(config)
}
