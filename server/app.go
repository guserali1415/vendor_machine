package main

import (
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	MACHINES_COUNT = 4
)

// This method used to initilize machines
func initMachines() []VendingMachine {
	machines := []VendingMachine{}
	for i := range MACHINES_COUNT {
		initialProducts := []Product{}
		initialProducts = append(initialProducts, Product{Price: 2, Name: "coffee", ID: 1, Stock: 3})
		initialProducts = append(initialProducts, Product{Price: 2, Name: "drink", ID: 2, Stock: 50})
		machine := VendingMachine{ID: i, Name: "Machine " + strconv.Itoa(i), IDLE: true, AssignedUserID: -1, Products: initialProducts}
		machines = append(machines, machine)
	}
	return machines
}

// Mutex used to prevent concurrent modification to shared resources
var mutex sync.Mutex

var machines []VendingMachine

// Used to check out in stock product
func findAvailableProductIndex(machine VendingMachine, name string) int {
	for i := range len(machine.Products) {
		if machine.Products[i].Name == name && machine.Products[i].Stock > 0 {
			return i
		}
	}
	return NOT_FOUND
}

func main() {

	machines = initMachines()

	//In memory users
	users := []User{}

	//used gin as rest api server
	router := gin.Default()

	//A middleware to disable cors and check user authentication for api section
	router.Use(func(ctx *gin.Context) {
		defer func() {
			r := recover()
			if r != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"messsage": "Unexpected error, check input data "})
				return
			}
		}()
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if ctx.Request.Method == "OPTIONS" {
			ctx.Status(204)
			return
		}

		if !strings.HasPrefix(ctx.Request.URL.Path, "/api") {
			ctx.Next()
			return
		}
		if strings.HasPrefix(ctx.Request.URL.Path, "/api/users/new") {
			ctx.Next()
			return
		}

		user := User{ID: -1}

		token := ctx.Request.Header["Authorization"][0]
		for i := range len(users) {
			if users[i].Token == token {
				user = users[i]
			}
		}
		if user.ID == -1 {
			ctx.JSON(http.StatusForbidden, gin.H{"messsage": "Not authorized"})
			return
		}
		ctx.Keys = make(map[string]interface{})
		ctx.Keys["user"] = user
		ctx.Next()
	})

	//used to generate new user and token
	router.GET("/api/users/new", func(context *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()

		lastID := 0
		if len(users) > 0 {
			lastID = users[len(users)-1].ID
		}
		lastID++
		user := User{ID: lastID, Token: uuid.New().String()}

		users = append(users, user)
		context.JSON(http.StatusOK, user)
	})

	//Machine group api to organize machine related api
	rMachines := router.Group("/api/machines")
	{
		//Get list of all machines
		rMachines.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, machines)
		})

		rMachine := rMachines.Group("/:id")
		{

			//Used this endpoind to select product to prevent other users selecet simultaneously
			rMachine.GET("/select_product/:name", func(ctx *gin.Context) {
				user := ctx.Keys["user"].(User)

				id_str, err_1 := ctx.Params.Get("id")
				if !err_1 {
					ctx.JSON(http.StatusBadRequest, gin.H{"message": "id parameter is not set"})
					return
				}
				id, err := strconv.Atoi(id_str)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"message": "Number excepted as id paramater"})
					return
				}

				name_str, err_2 := ctx.Params.Get("name")
				if !err_2 {
					ctx.JSON(http.StatusBadRequest, gin.H{"message": "name parameter is not set"})
					return
				}
				if !machines[id].IDLE && machines[id].AssignedUserID != user.ID {
					ctx.JSON(http.StatusBadRequest, gin.H{"message": "This machine is busy now please try again later"})
					return
				}

				productIndex := findAvailableProductIndex(machines[id], name_str)

				if productIndex == NOT_FOUND {
					ctx.JSON(http.StatusBadRequest, gin.H{"message": "Product not found or finished"})
					return
				}

				mutex.Lock()
				defer mutex.Unlock()

				machines[id].IDLE = false
				machines[id].AssignedUserID = user.ID
				machines[id].SelectedProductIndex = productIndex

				ctx.JSON(http.StatusOK, machines[id].Products[productIndex])

			})

			//Used to purchase product
			rMachine.GET("/insert_coin/:coin_count", func(ctx *gin.Context) {
				user := ctx.Keys["user"].(User)

				id_str, err_1 := ctx.Params.Get("id")
				if !err_1 {
					ctx.JSON(http.StatusBadRequest, gin.H{"message": "id parameter is not set"})
					return
				}
				id, err := strconv.Atoi(id_str)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"message": "Number excepted as id paramater " + strconv.Itoa(id)})
					return
				}

				if machines[id].AssignedUserID != user.ID {
					ctx.JSON(http.StatusBadRequest, gin.H{"message": "This machine is not assigned with this user"})
					return
				}

				coin_count_str, err_2 := ctx.Params.Get("coin_count")
				if !err_2 {
					ctx.JSON(http.StatusBadRequest, gin.H{"message": "coin_count parameter is not set"})
					return
				}
				coin_count, err3 := strconv.Atoi(coin_count_str)
				if err3 != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"message": "Number excepted as coin_count paramater " + strconv.Itoa(id)})
					return
				}

				if machines[id].IDLE {
					ctx.JSON(http.StatusOK, gin.H{"message": "Please select a product first"})
				} else {
					mutex.Lock()
					defer mutex.Unlock()

					if machines[id].AssignedUserID != user.ID {
						ctx.JSON(http.StatusBadRequest, gin.H{"message": "This machine is not assigned with you"})
						return
					}

					if coin_count != machines[id].Products[machines[id].SelectedProductIndex].Price {
						ctx.JSON(http.StatusBadRequest, gin.H{"message": "Please insert exact coins, refunding ..."})
						return
					}
					machines[id].CoinsInventory += machines[id].Products[machines[id].SelectedProductIndex].Price
					machines[id].Products[machines[id].SelectedProductIndex].Stock -= 1
					machines[id].AssignedUserID = -1
					machines[id].SelectedProductIndex = NOT_FOUND
					machines[id].IDLE = true

					ctx.JSON(http.StatusOK, machines[id])
				}

			})

		}

	}

	//Used another server on porn 5000 to server web files

	router_web := gin.Default()
	router_web.Static("/", "./web")

	//Run it in goroutine to server independently
	//In production we should handle termination of goroutune to prevent memroy leaks
	go func(r *gin.Engine) {
		//Run web server on 5000
		r.Run(":5000")
	}(router_web)

	//Run api server on 8080
	router.Run(":8080")

}
