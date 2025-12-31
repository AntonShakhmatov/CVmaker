package buildcv

import (
	"context"
	"net/http"
	"path/filepath"
	"time"

	"backend/internal/db"
	"backend/internal/services/pdfmaker"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	maker := pdfmaker.NewMaker(db.DB)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)

	docData, err := maker.LoadAll(ctx, 1)
	must(err)

	lastName, err := maker.LoadLastname(ctx, 1)

	// outName := strings.TrimSuffix(filepath.Base(pdfPath), filepath.Ext(pdfPath)) + "_NEW.pdf"
	outPath := filepath.Join("uploads", lastName.LastName+".pdf")

	must(maker.MakeCV(outPath, docData))

	cancel()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
