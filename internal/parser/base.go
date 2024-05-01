package sberparser

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sbercheck/internal/config"
	"sbercheck/internal/model"
	"strings"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

func Start(cfg *config.Config) error {
	client := &http.Client{}

	cookiePaths, err := findContainers(cfg.Parser.CookieFolder)
	if err != nil {
		return err
	}

	if err := parseAll(client, &cfg.Parser, cookiePaths); err != nil {
		log.Fatal("Error launching: ", err)
	}

	return nil
}

func parseAll(client *http.Client, cfg *config.ParserConfig, cookiePaths []container) error {
	var wg sync.WaitGroup
	wg.Add(len(cookiePaths))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Container", "Promo", "Value", "Expiry Date"})

	parseContainer := func(cookiePath container) {
		defer wg.Done()

		select {
		case <-ctx.Done():
			log.Print("Error occured: ", ctx)
			return
		default:
		}

		err := parse(client, cfg, cookiePath, t)
		if err != nil {
			log.Printf("Container %s failed with error: %v", cookiePath, err)
			cancel()
			return
		}
	}

	for _, cookie := range cookiePaths {
		go parseContainer(cookie)
	}

	wg.Wait()
	log.Printf("Finished, parsed %d containers", len(cookiePaths))
	log.Print(cookiePaths)

	t.SetColumnConfigs([]table.ColumnConfig{{Number: 1, AutoMerge: true}})
	t.SortBy([]table.SortBy{
		{Name: "Container", Mode: table.AscNumericAlpha},
	})
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true

	out, err := os.Create("filename.txt")
	if err != nil {
		log.Print("Error creating file: ", err)
	}
	defer out.Close()

	io.Copy(out, strings.NewReader(t.Render()))
	return nil
}

func parse(client *http.Client, cfg *config.ParserConfig, cont container, t table.Writer) error {
	req, err := BuildRequest(cfg, cont.path)
	if err != nil {
		return err
	}

	// defer resp.Body.Close() might be the case why reusing request does not work
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request: ", err)
		return err
	}
	defer resp.Body.Close()

	response := &model.ResponseModel{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Print("Error decoding JSON:", err)
		return err
	}

	currentCont := strings.Split(cont.name, ".")[0]
	switch len(response.PromoCodes) {
	case 0:
		t.AppendRow([]interface{}{
			currentCont,
			"-", 0,

			// promo.Description, // Add regular expresions (check utils)
		})
		// t.AppendSeparator()
	default:
		for _, promo := range response.PromoCodes {
			expirationDate := time.Date(
				promo.ExpireAt.Year,
				time.Month(promo.ExpireAt.Month),
				promo.ExpireAt.Day,
				promo.ExpireAt.Hours,
				promo.ExpireAt.Minutes,
				promo.ExpireAt.Seconds,
				promo.ExpireAt.Nanos,
				time.UTC,
			)
			t.AppendRow([]interface{}{
				currentCont,
				promo.Key,
				promo.Amount,
				expirationDate.Format("02-01-2006 15:04"),
				// promo.Description, // Add regular expresions (check utils)
			})
		}
		// t.AppendSeparator()
	}

	return nil
}
