package main

import (
	"encoding/csv"
	"fmt"
	"github.com/ipdata/go"
	"os"
	"strconv"
	"time"
)

// Define a custom type for IP address information
type IP struct {
	Address   string
	Latitude  float64
	Longitude float64
	City      string
	Postal    string
	Region    string
	Country   string
}

func main() {
	// Define IP addresses and their associated data
	addresses := []string{
		"your_ip_address",
	}

	// Replace with your actual API key
	apiKey := "YOUR_API_KEY"

	// Create or open a CSV file for writing
	file, err := os.Create("ip_data-" + time.DateTime + ".csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the CSV header
	header := []string{"IP", "Latitude", "Longitude", "City", "Postal", "Region", "Country"}
	writer.Write(header)

	ipd, err := ipdata.NewClient(apiKey)
	if err != nil {
		fmt.Println("Error creating IPData client:", err)
		return
	}

	// Loop through IP addresses and write data to the CSV
	for _, ip := range addresses {
		data, err := ipd.Lookup(ip)
		if err != nil {
			fmt.Println("Error fetching data:", err)
			continue // Skip this IP address and proceed to the next one
		}

		// Create an instance of the IP type and populate it with data
		ipInfo := IP{
			Address:   ip,
			Latitude:  data.Latitude,
			Longitude: data.Longitude,
			City:      data.City,
			Postal:    data.Postal,
			Region:    data.Region,
			Country:   data.CountryName,
		}

		dataSlice := []string{
			ipInfo.Address,
			strconv.FormatFloat(ipInfo.Latitude, 'f', 6, 64),
			strconv.FormatFloat(ipInfo.Longitude, 'f', 6, 64),
			ipInfo.City,
			ipInfo.Postal,
			ipInfo.Region,
			ipInfo.Country,
		}

		// Write data to CSV
		writer.Write(dataSlice)
	}

	fmt.Println("CSV file has been created with the data.")
}
