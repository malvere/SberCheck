package sberparser

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sbercheck/internal/config"
	"sbercheck/internal/model"
	"strings"
	"sync"
)

const (
	promoListPoint   = "https://megamarket.ru/api/mobile/v1/promotionService/promoCode/list"
	profileInfoPoint = "https://megamarket.ru/api/mobile/v1/securityService/profile/get"
	cartSearchPoint  = "https://megamarket.ru/api/mobile/v2/cartService/cart/search"
	getCartPoint     = "https://megamarket.ru/api/mobile/v2/cartService/cart/get"
	cncServicePoint  = "https://megamarket.ru/api/mobile/v2/cncService/location/get"
	orderCreatePoint = "https://megamarket.ru/api/mobile/v1/checkoutService/order/create"
)

type OrderData struct {
	CartID     model.CartIdentification
	LocationId string
	Profile    model.CustomerProfile
}

func Duplicate(cfg *config.Config) error {
	client := &http.Client{}

	cookiePaths, err := findContainers(cfg.Parser.CookieFolder)
	if err != nil {
		return err
	}

	// if err := parseAll(client, &cfg.Parser, cookiePaths); err != nil {
	// 	log.Fatal("Error launching: ", err)
	// }

	// cart, err := searchCart(client, &cfg.Parser, cookiePaths[0])
	// if err != nil {
	// 	log.Print("Error launching: ", err)
	// 	return err
	// }
	// log.Print(cart.Elements[0].Identification.ID)

	// location, err := getCart(client, &cfg.Parser, cookiePaths[0], cart.Elements[0].Identification.ID)
	// if err != nil {
	// 	return err
	// }

	// profile, err := getProfile(client, &cfg.Parser, cookiePaths[0])
	// if err != nil {
	// 	return err
	// }

	// order, err := createOrder(client, &cfg.Parser, cookiePaths[0], cart.Elements[0].Identification.ID, &profile.Profile)
	// if err != nil {
	// 	return err
	// }

	// log.Print(location.LocationID)
	// log.Print(profile.Profile)
	// log.Print(order)

	duplicateALL(client, &cfg.Parser, cookiePaths, "k5123ak2q48hdv")
	return nil
}

func duplicateALL(client *http.Client, cfg *config.ParserConfig, cookiePaths []container, promoCode string) {
	var prepareOrdersGroup sync.WaitGroup
	var makeOrdersGroup sync.WaitGroup

	prepareOrdersGroup.Add(len(cookiePaths))
	makeOrdersGroup.Add(len(cookiePaths))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	OrderDataSlice := make([]*OrderData, len(cookiePaths))
	OrderResultSlice := make([]*model.OrderCreateResponse, len(cookiePaths))

	log.Printf("Detected %d profiles", len(cookiePaths))
	prepareOrders := func(i int, cookiePath container) {
		defer prepareOrdersGroup.Done()

		select {
		case <-ctx.Done():
			log.Print("Error occured: ", ctx)
			return
		default:
		}

		cart, err := searchCart(client, cfg, cookiePath)
		if err != nil {
			log.Print("Error launching: ", err)
			cancel()
			return
		}

		location, err := getCart(client, cfg, cookiePaths[0], cart.Elements[0].Identification.ID)
		if err != nil {
			cancel()
			return
		}

		profile, err := getProfile(client, cfg, cookiePaths[0])
		if err != nil {
			cancel()
			return
		}
		OrderDataSlice[i] = &OrderData{
			CartID:     cart.Elements[0].Identification,
			LocationId: location.LocationID,
			Profile:    profile.Profile,
		}

	}

	for i, cookie := range cookiePaths {
		go prepareOrders(i, cookie)
	}

	prepareOrdersGroup.Wait()

	log.Print("Orders prepaired")

	makeOrders := func(i int, cookiePath container) {
		defer makeOrdersGroup.Done()

		select {
		case <-ctx.Done():
			log.Print("Error occured: ", ctx)
			return
		default:
		}

		order, err := createOrder(client, cfg, cookiePath, OrderDataSlice[i].CartID.ID, &OrderDataSlice[i].Profile, promoCode)
		if err != nil {
			log.Printf("Order Create failed in container %s. Error: %v", cookiePath.name, err)
		}

		OrderResultSlice[i] = order
	}

	for i, cookie := range cookiePaths {
		go makeOrders(i, cookie)
	}
	makeOrdersGroup.Wait()

	for i, order := range OrderResultSlice {
		if order.Success {
			log.Printf("Container %s. Status: success", cookiePaths[i].name)
		} else {
			log.Printf("Container %s. Status: FAILED", cookiePaths[i].name)
			log.Printf("Message: %v", order.Message)
			log.Print("Errors: ")
			for i, er := range order.Errors {
				log.Printf("ERROR #%d:::| %s: %s", i, er.Title, er.Detail)
			}

		}
	}

}

func MakeCookies(cookiePath string) (string, error) {
	switch {
	case strings.HasSuffix(cookiePath, ".txt"):
		cookie, err := os.ReadFile(cookiePath)
		if err != nil {
			return "", err
		}

		return string(cookie), nil
	case strings.HasSuffix(cookiePath, ".json"):
		cookie, err := LoadCookies(cookiePath)
		if err != nil {
			log.Print("Error loading cookies from JSON: ", err)
			return "", err
		}
		return fmt.Sprintf("%s=%s;", cookie.Name, cookie.Value), nil
	}
	return "", fmt.Errorf("cookie has unknown type")
}

func MakeRequest(cfg *config.ParserConfig, cookiePath string, endpoint string) (*http.Request, error) {
	expData, err := os.ReadFile(cfg.Experiments)
	if err != nil {
		return nil, err
	}
	// var cookieString string
	cookieString, err := MakeCookies(cookiePath)
	if err != nil {
		return nil, err
	}

	experiments := bytes.NewBuffer(expData)

	req, err := http.NewRequest("POST", endpoint, experiments)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("content-type", "application/json")
	req.Header.Set("Cookie", string(cookieString))
	return req, nil
}

func searchCart(client *http.Client, cfg *config.ParserConfig, cont container) (*model.CartSearchResponse, error) {
	req, err := MakeRequest(cfg, cont.path, cartSearchPoint)
	if err != nil {
		return nil, err
	}

	// defer resp.Body.Close() might be the case why reusing request does not work
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	response := &model.CartSearchResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Print("Error decoding JSON:", err)
		return nil, err
	}

	return response, nil
}

func getCart(client *http.Client, cfg *config.ParserConfig, cont container, cartID string) (*model.GetCartResponse, error) {

	auth := &model.AuthBody{}

	authData, err := os.ReadFile(cfg.Experiments)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(authData, &auth); err != nil {
		return nil, err
	}

	reqBody := &model.GetCartRequest{
		Identification: struct {
			ID string "json:\"id\""
		}{cartID},
		IsCartStateValidationRequired: true,
		IsSelectedItemGroupsOnly:      true,
		IsSkipPersonalDiscounts:       true,
		Auth:                          auth.Auth,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", getCartPoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	// log.Print(req)

	cookieString, err := MakeCookies(cont.path)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", string(cookieString))

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	response := &model.GetCartResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Print("Error decoding JSON:", err)
		return nil, err
	}

	return response, nil
}

func getProfile(client *http.Client, cfg *config.ParserConfig, cont container) (*model.GetProfileResponse, error) {
	req, err := MakeRequest(cfg, cont.path, profileInfoPoint)
	if err != nil {
		return nil, err
	}

	// defer resp.Body.Close() might be the case why reusing request does not work
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	response := &model.GetProfileResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Print("Error decoding JSON:", err)
		return nil, err
	}
	return response, nil
}

func createOrder(client *http.Client, cfg *config.ParserConfig, cont container, cartID string, profile *model.CustomerProfile, promo string) (*model.OrderCreateResponse, error) {
	auth := &model.AuthBody{}

	authData, err := os.ReadFile(cfg.Experiments)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(authData, &auth); err != nil {
		return nil, err
	}

	reqBody := &model.OrderCreateRequest{
		Customer: *profile,
		Identification: struct {
			ID string "json:\"id\""
		}{cartID},
		Discounts: []model.OrderDiscount{
			{
				Type:    "PROMO_CODE",
				Voucher: promo,
			},
		},
		IsSelectedCartItemGroupsOnly: true,
		PaymentType:                  "SBERPAY",
		Auth:                         auth.Auth,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", orderCreatePoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	// log.Print(req)
	cookieString, err := MakeCookies(cont.path)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", string(cookieString))

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	response := &model.OrderCreateResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Print("Error decoding JSON:", err)
		return nil, err
	}

	return response, nil
}
