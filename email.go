package main

import (
	"fmt"

	mail "github.com/wneessen/go-mail"
)

// EmailConfig raccoglie i parametri del server SMTP. Da leggere da .env,
// MAI hardcodati nel sorgente.
type EmailConfig struct {
	Host     string // es: smtp.gmail.com
	Port     int    // es: 587
	Username string // es: aziende@tuodominio.it
	Password string
	From     string // mittente, es: aziende@tuodominio.it
}

// InviaBustaPaga invia il PDF allegato all'indirizzo email del dipendente.
func InviaBustaPaga(cfg EmailConfig, destinatario, nomeDipendente, mesePeriodo, pathPDF string) error {
	m := mail.NewMsg()

	if err := m.From(cfg.From); err != nil {
		return fmt.Errorf("indirizzo mittente non valido: %w", err)
	}
	if err := m.To(destinatario); err != nil {
		return fmt.Errorf("indirizzo destinatario non valido (%s): %w", destinatario, err)
	}

	m.Subject(fmt.Sprintf("Busta paga %s - %s", mesePeriodo, nomeDipendente))
	m.SetBodyString(mail.TypeTextPlain, fmt.Sprintf(
		"Ciao %s,\n\nin allegato trovi la busta paga relativa a %s.\n\nSaluti.",
		nomeDipendente, mesePeriodo,
	))

	m.AttachFile(pathPDF) // allega il PDF così com'è, nessuna conversione necessaria

	client, err := mail.NewClient(cfg.Host,
		mail.WithPort(cfg.Port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(cfg.Username),
		mail.WithPassword(cfg.Password),
		mail.WithTLSPolicy(mail.TLSMandatory),
	)
	if err != nil {
		return fmt.Errorf("creazione client SMTP: %w", err)
	}

	if err := client.DialAndSend(m); err != nil {
		return fmt.Errorf("invio email a %s: %w", destinatario, err)
	}
	return nil
}
