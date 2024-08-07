# QR Scanner

A Go application that decodes QR codes from image files and opens the URL in the default web browser.

## Setup

1. **Clone the Repository**

   ```sh
   git clone https://github/tabintel/qrscanner.git
   cd qrscanner
   ```

2. **Install Dependencies**

   Ensure you have Go installed, then run:

   ```sh
   go mod tidy
   ```

3. **Add Image File**

   Use the default image to test or place your image with the QR code in the `images/` directory and name it `sample.png` or update the `imagePath` in `main.go` to match your image file name.

4. **Run the Application**

   ```sh
   go run main.go
   ```

   The application will decode the QR code from the image and open the URL in your default web browser.


