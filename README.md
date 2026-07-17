- [Descrizione](#descrizione)
- [Funzionalità](#funzionalità)
- [Tecnologie](#tecnologie)
- [Installazione](#installazione)
- [Configurazione env](#configurazione-env)
- [Database](#database)

## Descrizione
Divide PDF multi-pagina con buste paga e le invia automaticamente ai dipendenti. 
Strumento che prende un PDF multi-pagina contenente le buste paga, le divide automaticamente in file singoli organizzati per anno e mese, e le invia via email a ciascun destinatario corretto — eliminando il lavoro manuale e riducendo il rischio di errore.

## Funzionalità
* Lettura di PDF multi-pagina con buste paga
* Riconoscimento automatico del dipendente tramite codice fiscale
* Estrazione della pagina singola come PDF autonomo
* Archiviazione ordinata in cartelle anno/mese
* Invio email automatico al dipendente con allegato
* Tracciamento nel database dello stato di ogni invio
* Prevenzione degli invii doppi grazie al controllo di idempotenza

## Tecnologie
* Go — linguaggio principale
* MySQL — persistenza dati e log invii
* pdfcpu — manipolazione e split dei PDF
* ledongthuc/pdf — estrazione testo dai PDF
* SMTP — invio email

## Installazione
```bash
# Clona il repository
git clone [https://github.com/tuousername/gestore-buste.git](https://github.com/tuousername/gestore-buste.git)
cd gestore-buste

# Copia e configura le variabili d'ambiente
cp .env.example .env
# Modifica .env con le tue credenziali

# Scarica le dipendenze
go mod tidy

# Compila
go build -o gestore-buste
