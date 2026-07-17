# gestore-buste
Divide PDF multi-pagina con buste paga e le invia automaticamente ai dipendenti
Strumento che prende un PDF multi-pagina contenente le buste paga dei dipendenti, le divide automaticamente in file singoli organizzati per anno e mese, e le invia via email a ciascun destinatario corretto — eliminando il lavoro manuale e riducendo il rischio di errore.
Funzionalità
Lettura di PDF multi-pagina con buste paga
Riconoscimento automatico del dipendente tramite codice fiscale
Estrazione della pagina singola come PDF autonomo
Archiviazione ordinata in cartelle anno/mese
Invio email automatico al dipendente con allegato
Tracciamento nel database dello stato di ogni invio (pending / sent / failed)
Prevenzione degli invii doppi grazie al controllo di idempotenza
Tecnologie
Go — linguaggio principale
MySQL — persistenza dati e log invii
pdfcpu — manipolazione e split dei PDF
ledongthuc/pdf — estrazione testo dai PDF
SMTP — invio email
