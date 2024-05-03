package sberparser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sbercheck/internal/config"
	"slices"
	"strings"
)

type Cookie struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type container struct {
	path string
	name string
}

func LoadCookies(cookiePath string) (*Cookie, error) {
	var cookies []Cookie

	// log.Print("Starting parsing...")
	data, err := os.ReadFile(cookiePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &cookies); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return nil, err
	}
	idx := slices.IndexFunc(cookies, func(c Cookie) bool { return c.Name == "ecom_token" })
	return &cookies[idx], nil

}

// experiments.json must be reused instead of opening it everytime
// call without defer body.Close()?
func BuildRequest(cfg *config.ParserConfig, cookiePath string) (*http.Request, error) {
	expData, err := os.ReadFile(cfg.Experiments)
	if err != nil {
		return nil, err
	}
	var cookieString string
	switch {
	case strings.HasSuffix(cookiePath, ".txt"):
		cookie, err := os.ReadFile(cookiePath)
		if err != nil {
			return nil, err
		}
		cookieString = string(cookie)
	case strings.HasSuffix(cookiePath, ".json"):
		cookie, err := LoadCookies(cookiePath)
		if err != nil {
			log.Print("Error loading cookies from JSON: ", err)
			return nil, err
		}
		cookieString = fmt.Sprintf("%s=%s;", cookie.Name, cookie.Value)
	}

	experiments := bytes.NewBuffer(expData)

	req, err := http.NewRequest("POST", "https://megamarket.ru/api/mobile/v1/promotionService/promoCode/list", experiments)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", string(cookieString))

	return req, nil
}

func findContainers(root string) ([]container, error) {
	var containers []container
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(entry.Name()) != "" {
			containers = append(containers, container{path: "./" + path, name: entry.Name()})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func RegParse() {

	discountText := "Скидка зависит от суммы товаров в корзине: 1 000₽ на заказ от 4 000₽, 5 000₽ на заказ от 25 000₽, 10 000₽ на заказ от 50 000₽, 20 000₽ на заказ от 100 000₽"

	re := regexp.MustCompile(`(\d{1,3}\s*\d{0,3})₽ на заказ от (\d{1,3}\s*\d{0,3})₽`)
	matches := re.FindAllStringSubmatch(discountText, -1)

	for _, match := range matches {
		minOrder := match[1]
		discount := match[2]
		fmt.Printf("Min Order: %s, Discount: %s\n", minOrder, discount)
	}

}
