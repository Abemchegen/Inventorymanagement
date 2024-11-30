package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl              string
	Port                     int
	DbName                   string
	UserCollection           string
	OrderCollection           string
	ProductCollection		string
	ContextTimeout           int
	AccessTokenExpiryHour    int
	RefreshTokenExpiryHour   int
	AccessTokenSecret        string
	RefreshTokenSecret       string
	SellerCollection			string
	StoreCollection			string
	ReportCollection			string
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	dbURL := os.Getenv("DATABASE_URL")
	portStr := os.Getenv("PORT")
	dbname := os.Getenv("DB_NAME")
	usercoll := os.Getenv("user_collection")
	orderColl := os.Getenv("order_collection")
	productColl := os.Getenv("product_collection")
	contextTimeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	accessTokenExpiryHourStr := os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR")
	refreshTokenExpiryHourStr := os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	refreshTokenSecret := os.Getenv("REFRESH_TOKEN_SECRET")

	port, err := strconv.Atoi(portStr)
	sellerColl := os.Getenv("Sellers_collection")
	storesColl := os.Getenv("Store_collection")
	reportColl := os.Getenv("Report_collection")

	if err != nil {
		log.Fatal("Invalid PORT value")
		return nil, err
	}

	contextTimeout, err := strconv.Atoi(contextTimeoutStr)
	if err != nil {
		log.Fatal("Invalid CONTEXT_TIMEOUT value")
		return nil, err
	}

	accessTokenExpiryHour, err := strconv.Atoi(accessTokenExpiryHourStr)
	if err != nil {
		log.Fatal("Invalid ACCESS_TOKEN_EXPIRY_HOUR value")
		return nil, err
	}

	refreshTokenExpiryHour, err := strconv.Atoi(refreshTokenExpiryHourStr)
	if err != nil {
		log.Fatal("Invalid REFRESH_TOKEN_EXPIRY_HOUR value")
		return nil, err
	}

	config := &Config{
		DatabaseUrl:            dbURL,
		Port:                   port,
		DbName:                 dbname,
		ContextTimeout:         contextTimeout,
		AccessTokenExpiryHour:  accessTokenExpiryHour,
		RefreshTokenExpiryHour: refreshTokenExpiryHour,
		AccessTokenSecret:      accessTokenSecret,
		RefreshTokenSecret:     refreshTokenSecret,
		
		UserCollection:         usercoll,
		OrderCollection:         orderColl,
		ProductCollection:		productColl,
		SellerCollection:		sellerColl,
		StoreCollection:		storesColl,
		ReportCollection:		reportColl,
	}

	return config, nil
}
