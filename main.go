package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

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

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // Handle preflight requests (OPTIONS method)
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

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
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}


func main() {
    http.HandleFunc("/upload", uploadHandler)
    fmt.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
