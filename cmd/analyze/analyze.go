package main

import (
	"Project-structure/internal/analyzer"
	"Project-structure/pkg/property"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal().Msg("Expected file as an argument")
	}

	filename := os.Args[1]
	analyzer := analyzer.New(property.FilenameAnalyzer{}, property.ContentLengthAnalyzer{})

	report, err := analyzer.Analyze(filename)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to analyze file")
	}

	for k, v := range report {
		fmt.Println("%s\t%v\n", k, v)
	}
}
