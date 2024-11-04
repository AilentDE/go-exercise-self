package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/lpernett/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	images := []string{"./asset/image1.jpeg", "./asset/image2.jpeg", "./asset/image3.jpeg"}

	res, err := sendPostRequest(&images)
	if err != nil {
		fmt.Println("Error sending post request:", err)
		return
	}
	fmt.Println("Response:", res)
}

func transformImageToBase64(filePath string) (string, error) {
	fmt.Println("transformImageToBase64", filePath)

	// Read the image file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file")
		return "", err
	}
	defer file.Close()

	// Read the file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file")
		return "", err
	}

	// Encode the file content to base64
	imageBase64 := base64.StdEncoding.EncodeToString(fileContent)
	return imageBase64, nil
}

func sendPostRequest(images *[]string) (string, error) {
	// Transform the images to base64 concurrently
	imageBase64s := make([]string, len(*images))
	var wg sync.WaitGroup
	for index, image := range *images {
		wg.Add(1)
		go func(i int, img string) {
			defer wg.Done()
			imageBase64, err := transformImageToBase64(img)
			if err != nil {
				fmt.Println("Error transforming image to base64")
				return
			}
			imageBase64s[i] = imageBase64
		}(index, image)
	}
	wg.Wait()

	// Prepare form-data
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	defer writer.Close()

	// Add type field
	writer.WriteField("type", "TW_DRIVER_LICENSE_FRONT")

	// Add images as idCardImage fields
	for _, imageBase64 := range imageBase64s {
		writer.WriteField("idCardImage", imageBase64)
	}

	// Send the images to the server
	apiUrl := os.Getenv("API_URL") + "/api/v1/idcard/ocr"
	accessToken, err := getAccessToken()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiUrl, &buf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

func getAccessToken() (string, error) {
	apiUrl := os.Getenv("API_URL") + "/api/v1/api_key/auth"
	apiKey := os.Getenv("API_KEY")

	// Prepare x-www-form-urlencoded
	data := url.Values{}
	data.Set("apiKey", apiKey)
	data.Set("appId", "com.cyberlink.platform.faceme")
	data.Set("appSecret", "FaceMe#1")

	// Send the request
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
	var responseBody []byte
	if responseBody, err = io.ReadAll(resp.Body); err != nil {
		return "", err
	}

	var responseMap map[string]interface{}
	if err := json.Unmarshal(responseBody, &responseMap); err != nil {
		return "", err
	}

	token, ok := responseMap["token"].(string)
	if !ok {
		return "", fmt.Errorf("token not found in response")
	}
	return token, nil
}

func saveFile(filePath string, content *string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(*content)
	if err != nil {
		return err
	}

	return nil
}