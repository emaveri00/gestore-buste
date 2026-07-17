- [Descrizione](#descrizione)
- [Funzionalità](#funzionalità)
- [Tecnologie](#tecnologie)
- [Installazione](#installazione)
- [Configurazione](#configurazione)
- [Database](#database)


##Descrizione
Divide PDF multi-pagina con buste paga e le invia automaticamente ai dipendenti
Strumento che prende un PDF multi-pagina contenente le buste paga dei dipendenti, le divide automaticamente in file singoli organizzati per anno e mese, e le invia via email a ciascun destinatario corretto — eliminando il lavoro manuale e riducendo il rischio di errore.

##Funzionalità
Lettura di PDF multi-pagina con buste paga
Riconoscimento automatico del dipendente tramite codice fiscale
Estrazione della pagina singola come PDF autonomo
Archiviazione ordinata in cartelle anno/mese
Invio email automatico al dipendente con allegato
Tracciamento nel database dello stato di ogni invio (pending / sent / failed)
Prevenzione degli invii doppi grazie al controllo di idempotenza

##Tecnologie
Go — linguaggio principale
MySQL — persistenza dati e log invii
pdfcpu — manipolazione e split dei PDF
ledongthuc/pdf — estrazione testo dai PDF
SMTP — invio email

##Installazione

# Clona il repository
git clone https://github.com/tuousername/gestore-buste.git
cd gestore-buste

# Copia e configura le variabili d'ambiente
cp .env.example .env
# Modifica .env con le tue credenziali

# Scarica le dipendenze
go mod tidy

# Compila
go build -o gestore-buste


##Configurazione .env

# Database MySQL
username=root
password=tua_password
server=localhost
dbName=gestore_buste

# SMTP (es. Gmail)
SMTP_HOST=smtp.gmail.com
SMTP_USER=tua_email@gmail.com
SMTP_PASS=app_password_gmail
SMTP_FROM=noreply@tuaazienda.it

##Database:

CREATE DATABASE IF NOT EXISTS gestore_buste 
CHARACTER SET utf8mb4 
COLLATE utf8mb4_unicode_ci;

USE gestore_buste;

CREATE TABLE dipendenti (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    cf VARCHAR(16) NOT NULL UNIQUE
);

CREATE TABLE buste (
    idbuste INT AUTO_INCREMENT PRIMARY KEY,
    idDipendente INT NOT NULL,
    datainvio DATE NOT NULL,
    percorso VARCHAR(500) NOT NULL,
    FOREIGN KEY (idDipendente) REFERENCES dipendenti(id)
);
