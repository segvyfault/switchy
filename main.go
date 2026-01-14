package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func expand_home(path string) string {
	if strings.HasPrefix(path, "~/") {
		home := os.ExpandEnv("$HOME");
		path = strings.ReplaceAll(path, "~", home)
	}
	return path
}

func read_file(path string) string {
	path = expand_home(path);
	bytes, err := os.ReadFile(path);

	if err != nil {
		if os.IsNotExist(err) {
			os.Create(path)
			return ""
		} else {
			log.Fatalf("Error occured while reading: %s", err);
		}

	}

	return string(bytes)
}

func write_new_bg(new string) {
	path := expand_home("~/.paper");
	os.WriteFile(path, []byte(new), 0666);
}

func parse_config(contents string) []string {
	papers := strings.Split(contents, "\n");
	if papers[len(papers) - 1] == "" {
		papers = papers[:len(papers) - 1]	
	}

	for i := range len(papers) {
		papers[i] = expand_home(papers[i])
	}

	return papers
}

func parse_actions() *[][]string {
	contents := read_file("~/.paperactions");
	if contents == "" { return nil }

	read := strings.Split(contents, "\n");

	if read[len(read) - 1] == "" {
		read = read[:len(read) - 1]	
	}

	var actions [][]string 

	for i := range len(read) {
		actions = append(actions, strings.Split(read[i], " "))	
	}

	return &actions
}

func execute(config []string, previous string, actions *[][]string) {
	next_index := 0;

	if previous != "" {
		for i := range len(config) {
			if strings.HasPrefix(config[i], previous) {
				next_index = i + 1;
				if next_index >= len(config) {
					next_index = 0
				}
				break;
			}
		}
	}

	text := strings.Split(config[next_index], " ");
	matugen_args := []string{"image"}

	for i := range len(text) {
		matugen_args = append(matugen_args, text[i])
	}

	matugen := exec.Command("matugen", matugen_args...);
	log.Println("Executing", matugen.Args);
	if err := matugen.Run(); err == nil {
		matugen.Process.Release()
	}

	if actions != nil {
		for i := range len(*actions) {
			action := (*actions)[i];
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

	if len(os.Args) <= 1 || (len(os.Args) > 1 && os.Args[1] != "no-write") {
		write_new_bg(text[0])
	}
}

func main() {
	var parsed []string

	if len(os.Args) > 1 && os.Args[1] != "" && os.Args[1] != "no-write" {
		var contents string

		if os.Args[1] == "-" {
			bytes, err := io.ReadAll(os.Stdin) // "echo "meow" | ./program" situation
			if err != nil {
				log.Fatalln("Error occured while readin stdin:", err)
			}
			contents = string(bytes)
		} else {
			contents = read_file(os.Args[1])
		}

		parsed = parse_config(contents)
	} else {
		parsed = parse_config(read_file("~/.papers"))
	}

	if len(parsed) == 0 {
		log.Fatalln("The parsed wallpapers file is empty")
	}

	var previous string = read_file("~/.paper");
	execute(parsed, previous, parse_actions())
}
