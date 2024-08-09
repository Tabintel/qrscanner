/*
package main

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"

    "github.com/tuotoo/qrcode"
)

type QRCodeResult struct {
    URL string `json:"url"`
}

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

func main() {
    // Path to your image containing QR codes
    imagePath := filepath.Join("images", "sample.png")

    qrCodes, err := DecodeImage(imagePath)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Convert the QR code results to JSON
    qrCodeJSON, err := json.Marshal(qrCodes)
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }

    fmt.Println("QR Code content as JSON:", string(qrCodeJSON))
}

*/