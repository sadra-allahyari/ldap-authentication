package extractor

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

// ExtractStrFromFile extracts a string value for a specified key from a config file.
// It reads the config file using Viper and returns the value associated with the provided key.
//
// Parameters:
// - source: The section or source in the config file.
// - key: The key for which the value needs to be extracted.
//
// Returns:
// - string: The value associated with the provided key.
func ExtractStrFromFile(source, key string) string {
	// Check if .env file exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Println(".env file does not exist, reading CONFIG_PATH from environment variables")
	} else {
		// Load environment file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Can't load .env file")
		}
	}

	// Read the config path from the environment variable
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set in environment variables")
	}

	// Configure Viper to read the config file
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Get the value associated with the provided key in the specified source
	value := viper.GetString(source + "." + key)

	// Check if the value is empty and log an error if it is
	if value == "" {
		log.Fatalf("%v not available in config.yml", key)
	}

	return value
}

// ExtractBoolFromFile extracts a boolean value for a specified key from a config file.
// It reads the config file using Viper and returns the value associated with the provided key.
//
// Parameters:
// - source: The section or source in the config file.
// - key: The key for which the value needs to be extracted.
//
// Returns:
// - boolean: The value associated with the provided key.
func ExtractBoolFromFile(source, key string) bool {
	// Check if .env file exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Println(".env file does not exist, reading CONFIG_PATH from environment variables")
	} else {
		// Load environment file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Can't load .env file")
		}
	}

	// Read the config path from the environment variable
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set in environment variables")
	}

	// Configure Viper to read the config file
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	// Attempt to read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Retrieve the boolean value from the config file
	value := viper.GetBool(source + "." + key)

	// Check if the key is present in the config file
	if !viper.IsSet(source + "." + key) {
		log.Fatalf("%v not available in config.yml", key)
	}

	return value
}

// ExtractIntFromFile extracts an integer value for a specified key from a config file.
// It reads the config file using Viper and returns the value associated with the provided key.
//
// Parameters:
// - source: The section or source in the config file.
// - key: The key for which the value needs to be extracted.
//
// Returns:
// - integer: The value associated with the provided key.
func ExtractIntFromFile(source, key string) int {
	// Check if .env file exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Println(".env file does not exist, reading CONFIG_PATH from environment variables")
	} else {
		// Load environment file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Can't load .env file")
		}
	}

	// Read the config path from the environment variable
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set in environment variables")
	}

	// Configure Viper to read the config file
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Retrieve the integer value from the specified source and key
	value := viper.GetInt(source + "." + key)

	// Check if the value is zero, indicating that the key is not available or
	// the value is not an integer in the config file
	if value == 0 {
		log.Fatalf("%v not available or not an integer in config.yml", key)
	}

	// Return the extracted integer value
	return value
}
