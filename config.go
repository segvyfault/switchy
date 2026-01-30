package main

import (
	"fmt"
	"log"

	"io"
	"os"
	"strings"
)

const CONFIG_PREFIX = "~/.config/switchy/";

type Config struct {
	current string;   // current wallpaper
	papers []string;  // list of wallpapers
}

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

func (c *Config) parse_config(contents string) []string {
	papers := strings.Split(contents, "\n");
	if papers[len(papers) - 1] == "" {
		papers = papers[:len(papers) - 1]	
	}

	for i := range len(papers) {
		papers[i] = expand_home(papers[i])
	}

	return papers
}

func (c *Config) parse_wallpapers() error {
	var path     string = CONFIG_PREFIX + "papers"
	var parsed []string

	if len(os.Args) > 1 && os.Args[1] != "" && os.Args[1] != "no-write" {
		var contents string

		if os.Args[1] == "-" {
			path = "stdin"

			bytes, err := io.ReadAll(os.Stdin) // "echo "meow" | ./program" situation
			if err != nil {
				log.Fatalln("Error occured while readin stdin:", err)
			}
			contents = string(bytes)
		} else {
			path = os.Args[1];
			contents = read_file(path)
		}

		parsed = c.parse_config(contents)
	} else {
		parsed = c.parse_config(read_file(path))
	}

	// i know returning error here is non-sense
	// because i'll either way just panic
	// but this project is more of a learning experience
	// while writing THIS function i completely forgot 
	// how to link them to a struct instance
	if len(parsed) == 0 {
		return fmt.Errorf("Wallpapers file (%v) is empty", path)
	} else {
		c.papers = parsed;
		return nil
	}
}

func (c *Config) write_new_bg(new string) {
	path := expand_home(CONFIG_PREFIX + "paper");
	os.WriteFile(path, []byte(new), 0666);
}

func parse_all() Config {
	var c Config

	if err := c.parse_wallpapers(); err != nil {
		log.Fatalln(err.Error())
	}

	c.current = read_file(CONFIG_PREFIX + "paper")

	return c
}
