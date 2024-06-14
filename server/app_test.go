package main

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
)

const (
	url = "http://127.0.0.1:8080"
)

var user User
var machinesList []VendingMachine
var selectedMachine VendingMachine
var selectedProduct Product

func TestNewUser(t *testing.T) {
	fmt.Println("Testing /api/users/new endpoint")

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&user).
		Get(url + "/api/users/new")
	if err != nil || resp.StatusCode() != http.StatusOK {
		t.Errorf("Error in api call")
	}
}

func TestGetMachines(t *testing.T) {
	fmt.Println("Testing /api/machines/ endpoint")
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", user.Token).
		SetResult(&machinesList).
		Get(url + "/api/machines/")
	if err != nil || resp.StatusCode() != http.StatusOK {
		t.Errorf("Error in api call" + string(resp.Body()))
	} else {
		//result := gjson.Get(string(resp.Body()), "ok")

		if len(machinesList) == 0 {
			t.Errorf("Machines list empty")
		}
	}

}

func TestSelectDrink(t *testing.T) {
	machineID := selectedMachine.ID
	endpoint := "/api/machines/" + strconv.Itoa(machineID) + "/select_product/drink"
	fmt.Println("Testing " + endpoint + " endpoint")
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", user.Token).
		SetResult(&selectedProduct).
		Get(url + endpoint)
	if err != nil || resp.StatusCode() != http.StatusOK {
		t.Errorf("Error in api call" + string(resp.Body()))
	}

}
func TestInsertCoin(t *testing.T) {
	machineID := machinesList[0].ID
	endpoint := "/api/machines/" + strconv.Itoa(machineID) + "/insert_coin/" + strconv.Itoa(selectedProduct.Price)
	fmt.Println("Testing " + endpoint + " endpoint")
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", user.Token).
		Get(url + endpoint)
	if err != nil || resp.StatusCode() != http.StatusOK {
		t.Errorf("Error in api call" + string(resp.Body()))
	}
}

func TestSelectCoffee(t *testing.T) {
	machineID := selectedMachine.ID
	endpoint := "/api/machines/" + strconv.Itoa(machineID) + "/select_product/coffee"
	fmt.Println("Testing " + endpoint + " endpoint")
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", user.Token).
		SetResult(&selectedProduct).
		Get(url + endpoint)
	if err != nil || resp.StatusCode() != http.StatusOK {
		t.Errorf("Error in api call" + string(resp.Body()))
	}

}

func TestInsertCoffee(t *testing.T) {
	machineID := machinesList[0].ID
	endpoint := "/api/machines/" + strconv.Itoa(machineID) + "/insert_coin/" + strconv.Itoa(selectedProduct.Price)
	fmt.Println("Testing " + endpoint + " endpoint")
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", user.Token).
		Get(url + endpoint)
	if err != nil || resp.StatusCode() != http.StatusOK {
		t.Errorf("Error in api call" + string(resp.Body()))
	}
}
