package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/goura/yaku-cli/pkg/config"
	"github.com/goura/yaku-cli/pkg/translator"
	"github.com/urfave/cli/v2"
	"golang.org/x/text/language"
)

func main() {
	var engineTag, sourceLangID, targetLangID string

	app := &cli.App{
		Name:  "yaku",
		Usage: "Translate source language to target language",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "engine",
				Aliases:     []string{"e"},
				Value:       "deepl",
				Destination: &engineTag,
				Usage:       "translation engine to use",
			},
			&cli.StringFlag{
				Name:        "source-lang",
				Aliases:     []string{"s"},
				Destination: &sourceLangID,
				Required:    true,
				Usage:       "language to be translated from",
			},
			&cli.StringFlag{
				Name:        "target-lang",
				Aliases:     []string{"t"},
				Destination: &targetLangID,
				Required:    true,
				Usage:       "language to be translated into",
			},
		},
		Action: func(cCtx *cli.Context) error {
			// Read config
			conf := config.NewEnvConfig()

			// Convert language identifiers into language.Tag
			sourceLang, err := language.Parse(sourceLangID)
			if err != nil {
				return fmt.Errorf("Couldn't parse lang:%s", sourceLangID)
			}

			targetLang, err := language.Parse(targetLangID)
			if err != nil {
				return fmt.Errorf("Couldn't parse lang:%s", targetLangID)
			}

			// Instanciate a translator with the specified engine
			var instance translator.TranslatorInstance
			switch engineTag {
			case "dummyfortest":
				instance = translator.NewDummyTranslator()
			case "deepl":
				instance = translator.NewDeepLTranslator()
			default:
				return fmt.Errorf("engine:%s is not supported", engineTag)
			}

			// Read stdin
			src := ""
			for {
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				s := scanner.Text()
				src += s
				if s == "" {
					break
				}
			}

			// Execute translation
			s, err := instance.DoTranslation(cCtx.Context, conf, sourceLang, targetLang, src)
			if err != nil {
				return err
			}

			fmt.Println(s)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
