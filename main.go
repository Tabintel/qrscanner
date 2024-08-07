package main

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/tuotoo/qrcode"
    "github.com/pkg/browser"
)

// DecodeImage decodes an image file containing a QR code and returns the QR code content
func DecodeImage(imagePath string) (string, error) {
    // Open the image file
    file, err := os.Open(imagePath)
    if err != nil {
        return "", fmt.Errorf("error opening image file: %w", err)
    }
    defer file.Close()

    // Decode the QR code from the file
    qrCode, err := qrcode.Decode(file)
    if err != nil {
        return "", fmt.Errorf("error decoding QR code: %w", err)
    }

    return qrCode.Content, nil
}

func main() {
    // Image file path within the project directory
    imagePath := filepath.Join("images", "sample.png")

    // Decode the QR code from the image
    url, err := DecodeImage(imagePath)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Display the URL/link
    fmt.Println("QR Code content:", url)

    // Open the URL in the default web browser
    err = browser.OpenURL(url)
    if err != nil {
        fmt.Println("Error opening URL in default browser:", err)
        return
    }

    fmt.Println("URL opened in default browser successfully!")
}
