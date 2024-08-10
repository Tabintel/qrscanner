document.querySelector("form").addEventListener("submit", async function(event) {
    event.preventDefault();

    const formData = new FormData();
    const fileField = document.querySelector('input[type="file"]');

    formData.append("image", fileField.files[0]);

    try {
        const response = await fetch("https://decode-ohkt.onrender.com/upload", {
            method: "POST",
            body: formData,
        });

        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(errorText);
        }

        const result = await response.json();
        console.log("QR Code content as JSON:", result);

        // Display the result in the UI
        const resultDiv = document.querySelector("#result");
        resultDiv.innerHTML = ''; // Clear previous results

        result.forEach((qr, index) => {
            const linkElement = document.createElement('a');
            linkElement.href = qr.url;
            linkElement.textContent = `QR Code ${index + 1}: ${qr.url}`;
            linkElement.target = '_blank'; // Open the link in a new tab
            resultDiv.appendChild(linkElement);
            resultDiv.appendChild(document.createElement('br')); // Add a line break
        });
    } catch (error) {
        console.error("Error:", error);
        alert("An error occurred: " + error.message);
    }
});
