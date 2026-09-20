from flask import Flask, request
from datetime import datetime
import os

app = Flask(__name__)

# the pictures are currently saved in a folder, this makes debugging in the beginning a lot easier
SAVE_FOLDER = "bilderrasp"
os.makedirs(SAVE_FOLDER, exist_ok=True)

@app.route("/upload", methods=["POST"])
def upload():

    bilddaten = request.data

    # zeit = datetime.now().strftime("%Y%m%d_%H%M%S")

    # dateiname = f"{SAVE_FOLDER}/bild_{zeit}.jpg"

    dateiname = f"{SAVE_FOLDER}/bild_aktuell.jpg"
    with open(dateiname, "wb") as f:
        f.write(bilddaten)

    print(f"Gespeichert: {dateiname}")

    return "Bild gespeichert", 200

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)