package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/ledongthuc/pdf"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type Dipendente struct {
	id    int
	nome  string
	email string
	cf    string
}

func dbConn() *sql.DB {
	godotenv.Load(".env")
	username := os.Getenv("username")
	password := os.Getenv("password")
	server := os.Getenv("server")
	dbName := os.Getenv("dbName")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", username, password, server, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func PdfSplitter(path string, dipendenti []Dipendente, db *sql.DB) {
	cfg := EmailConfig{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     587,
		Username: os.Getenv("SMTP_USER"),
		Password: os.Getenv("SMTP_PASS"),
		From:     os.Getenv("SMTP_FROM"),
	}

	f, r, err := pdf.Open(path)
	if err != nil {
		fmt.Printf("apertura %s: %s\n", path, err)
		return
	}
	defer f.Close()

	if r.NumPage() < 1 {
		fmt.Printf("%s non contiene pagine\n", path)
		return
	}

	conf := model.NewDefaultConfiguration()
	re := regexp.MustCompile(`(?i)(gennaio|febbraio|marzo|aprile|maggio|giugno|luglio|agosto|settembre|ottobre|novembre|dicembre)\s*/\s*(\d{4})`)

	for i := 0; i < r.NumPage(); i++ {
		page := r.Page(i + 1)
		text, err := page.GetPlainText(nil)
		if err != nil {
			fmt.Printf("errore estrazione testo pagina %d: %v\n", i+1, err)
			continue
		}

		for idx := range dipendenti {
			if !strings.Contains(strings.ToLower(text), strings.ToLower(dipendenti[idx].cf)) {
				continue
			}

			m := re.FindStringSubmatch(text)
			if m == nil {
				fmt.Printf("CF trovato ma mese/anno non trovato per %s\n", dipendenti[idx].nome)
				continue
			}

			nomeMese := strings.ToLower(m[1])
			anno := m[2]

			cartella := filepath.Join("buste", anno, nomeMese)
			if err := os.MkdirAll(cartella, 0755); err != nil {
				fmt.Println("errore creando cartelle:", err)
				continue
			}

			nomeOutput := filepath.Join(cartella, fmt.Sprintf("%s.pdf", dipendenti[idx].nome))

			err = api.TrimFile(path, nomeOutput, []string{strconv.Itoa(i + 1)}, conf)
			if err != nil {
				fmt.Printf("errore salvando pagina %d: %v\n", i+1, err)
				continue
			}

			fmt.Println("Ho trovato", dipendenti[idx].nome, "-> salvato", nomeOutput)

			if err = InviaBustaPaga(cfg, dipendenti[idx].email, dipendenti[idx].nome, nomeMese+" "+anno, nomeOutput); err != nil {
				fmt.Println("errore invio email:", err)
				continue
			}

			_, err = db.Exec("INSERT INTO buste(idDipendente, datainvio, percorso) VALUES (?,?,?)",
				dipendenti[idx].id, time.Now(), nomeOutput)
			if err != nil {
				fmt.Println("errore inserimento DB:", err)
				continue
			}
		}
	}
}

func main() {
	fmt.Println("Gestore buste")
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Inserisci il file.pdf")
		return
	}

	db := dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT id, nome, email, cf FROM dipendenti")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	dipendenti := []Dipendente{}
	for rows.Next() {
		var d Dipendente
		if err := rows.Scan(&d.id, &d.nome, &d.email, &d.cf); err != nil {
			fmt.Println("errore scan:", err)
			continue
		}
		dipendenti = append(dipendenti, d)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	PdfSplitter(args[1], dipendenti, db)
}
