package sberparser

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sbercheck/internal/config"
	"sbercheck/internal/model"
	"strings"
	"sync"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
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

	re := regexp.MustCompile(`(-?\d{1,3}(?: \d{3})*)₽ (?:на заказ от|от) (\d{1,3}(?: \d{3})*)₽`)
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Container", "Bonuses", "Promo", "Value", "Expiry Date", "Discounts", "Discounts"}, table.RowConfig{AutoMerge: true})
	t.AppendHeader(table.Row{"Container", "Bonuses", "Promo", "Value", "Expiry Date", "Value", "Summ"})

	parseContainer := func(cookiePath container) {
		defer wg.Done()

		select {
		case <-ctx.Done():
			log.Print("Error occured: ", ctx)
			return
		default:
		}

		err := parse(client, cfg, cookiePath, t, re)
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

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: true},
		{Number: 2, AutoMerge: false},
		{Number: 3, AutoMerge: true},
		{Number: 4, AutoMerge: false},
		{Number: 5, AutoMerge: false},
		{Number: 6, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 7, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
	})
	t.SortBy([]table.SortBy{
		{Name: "Container", Mode: table.AscNumericAlpha},
		// {Name: "Promo", Mode: table.AscAlphaNumeric},
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

func parse(client *http.Client, cfg *config.ParserConfig, cont container, t table.Writer, re *regexp.Regexp) error {
	req, err := BuildRequest(cfg, cont.path)
	if err != nil {
		return err
	}

	profile, err := getProfile(client, cfg, cont)
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

	response := &model.PromoCodesList{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Print("Error decoding JSON:", err)
		return err
	}
	parsePromo(profile, response, cont, t, re)

	return nil
}
