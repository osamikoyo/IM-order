package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct{
	AmqpUrl string
	Port string
	Host string
	MognoURl string
	DBname string
	Collname string
}

func Load() (*Config, error){
	err := godotenv.Load("dev.env")
	if err != nil{
		return nil, err
	}

	viper.AutomaticEnv()

	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("PORT", "50055")
	viper.SetDefault("AMQP_URL", "amqp://localhost:5672")
	viper.SetDefault("MONGO_URL", "mongodb://localhost:27017")
	viper.SetDefault("DB_NAME", "im")
	viper.SetDefault("COLL_NAME", "order")

	host := viper.GetString("HOST")
	port := viper.GetString("PORT")
	amqpUrl := viper.GetString("AMQP_URL")
	mongoUrl := viper.GetString("MONGO_URL")
	dbName := viper.GetString("DB_NAME")
	collName := viper.GetString("COLL_NAME")

	return &Config{
		Host: host,
		Port: port,
		AmqpUrl: amqpUrl,
		MognoURl: mongoUrl,
		DBname: dbName,
		Collname: collName,
	}, err
}