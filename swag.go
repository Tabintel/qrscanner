package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "github.com/tuotoo/qrcode"
    //_ "your_project/docs" // This is to import the Swagger docs
    "github.com/swaggo/http-swagger" // http-swagger middleware
)

// QRCodeResult represents the structure of the decoded QR code content
type QRCodeResult struct {
    URL string `json:"url"`
}

// DecodeImage decodes the QR code from the image file
func DecodeImage(imagePath string) ([]QRCodeResult, error) {
    file, err := os.Open(imagePath)
    if err != nil {
        return nil, fmt.Errorf("error opening image file: %w", err)
    }
    defer file.Close()

    // Decode the QR code directly from the file
    qrCode, err := qrcode.Decode(file)
    if err != nil {
        return nil, fmt.Errorf("error decoding QR code: %w", err)
    }

    qrCodes := []QRCodeResult{
        {URL: qrCode.Content},
    }

    return qrCodes, nil
}

// uploadHandler handles the image upload and QR code decoding
// @Summary Upload an image and decode QR codes
// @Description Upload an image file, decode any QR codes present, and return the decoded URLs
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image file with QR code"
// @Success 200 {array} QRCodeResult
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Server error"
// @Router /upload [post]
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    file, _, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "Error reading file: "+err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()

    tempFile, err := ioutil.TempFile("images", "upload-*.png")
    if err != nil {
        http.Error(w, "Error saving file: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer os.Remove(tempFile.Name())

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        http.Error(w, "Error reading file: "+err.Error(), http.StatusInternalServerError)
        return
    }

    if _, err = tempFile.Write(fileBytes); err != nil {
        http.Error(w, "Error writing file: "+err.Error(), http.StatusInternalServerError)
        return
    }

    qrCodes, err := DecodeImage(tempFile.Name())
    if err != nil {
        http.Error(w, "Error decoding QR code: "+err.Error(), http.StatusInternalServerError)
        return
    }

    jsonResponse, err := json.Marshal(qrCodes)
    if err != nil {
        http.Error(w, "Error encoding JSON: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

// @title QR Code Decoder API
// @version 1.0
// @description API for decoding QR codes from images
// @host localhost:8080
// @BasePath /
func main() {
    http.HandleFunc("/upload", uploadHandler)

    // Swagger route
    http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

    fmt.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
