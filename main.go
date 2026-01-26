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

	if config.actions != nil {
		for i := range len(*config.actions) {
			action := (*config.actions)[i];
			program := action[0];
			args := action[1:];

			var dontWaitForEnd bool  = false;
			var err            error = nil;

			if args[len(args) - 1] == "nowait" {
				dontWaitForEnd = true;	
				args = action[:len(args) - 1]
			}

			cmd := exec.Command(program, args...);
			log.Println("Executing", cmd.Args);

			if dontWaitForEnd { 
				err = cmd.Start(); 
			} else { 
				err = cmd.Run();
			}

			if err == nil {
				cmd.Process.Release()
			}
		}
	}

	if len(os.Args) <= 1 || isNoWrite {
		config.write_new_bg(text[0])
	}
}

func main() {
	config := parse_all()
	execute(config)
}
