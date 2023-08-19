package itswizard_module_berlin_package

import (
	"net/http"
	"time"
)

type BerlinBsp struct {
	grant_type       string
	client_id        string
	client_secret    string
	access_token     string
	token_type       string
	client           *http.Client
	fqdn             string
	authentification bool
	createdTokenTime time.Time
}

type School struct {
	SchuleBezirkName     string   `json:"schuleBezirkName"`
	SchuleBezirkNummer   string   `json:"schuleBezirkNummer"`
	SchuleDtGeaendertAm  string   `json:"schuleDtGeaendertAm"`
	SchuleName           string   `json:"schuleName"`
	SchuleNummer         string   `json:"schuleNummer"`
	SchuleSchulformListe []string `json:"schuleSchulformListe"`
	SchuleTyp            string   `json:"schuleTyp"`
	SchuleUID            string   `json:"schuleUID"`
}

type User struct {
	BenutzerAktionCode        string `json:"benutzerAktionCode"`
	BenutzerBezugspersonListe []struct {
		BezugspersonPersonUID string `json:"bezugspersonPersonUID"`
		BezugspersonTyp       string `json:"bezugspersonTyp"`
	} `json:"benutzerBezugspersonListe"`
	BenutzerDtAktion     string `json:"benutzerDtAktion"`
	BenutzerGeburtsdatum string `json:"benutzerGeburtsdatum"`
	BenutzerGlobalRolle  string `json:"benutzerGlobalRolle"`
	BenutzerNachname     string `json:"benutzerNachname"`
	BenutzerSchuleListe  []struct {
		BenutzerKlasseListe []struct {
			BenutzerKlasseRolle string `json:"benutzerKlasseRolle"`
			KlasseUID           string `json:"klasseUID"`
		} `json:"benutzerKlasseListe"`
		BenutzerKursListe []struct {
			BenutzerKursRolle string `json:"benutzerKursRolle"`
			EnutzerKursUID    string `json:"enutzerKursUID"`
		} `json:"benutzerKursListe"`
		BenutzerSchuleRolle string `json:"benutzerSchuleRolle"`
		BenutzerStatus      string `json:"benutzerStatus"`
		SchuleUID           string `json:"schuleUID"`
	} `json:"benutzerSchuleListe"`
	BenutzerVorname string `json:"benutzerVorname"`
	PersonUID       string `json:"personUID"`
}

type Class struct {
	KlasseDtGeaendertAm string `json:"klasseDtGeaendertAm"`
	KlasseName          string `json:"klasseName"`
	KlasseUID           string `json:"klasseUID"`
	SchuleUID           string `json:"schuleUID"`
}

type Course struct {
	KursBezeichnung   string `json:"kursBezeichnung"`
	KursDtGeaendertAm string `json:"kursDtGeaendertAm"`
	KursFach          string `json:"kursFach"`
	KursJahrgang      string `json:"kursJahrgang"`
	KursSchulform     string `json:"kursSchulform"`
	KursUID           string `json:"kursUID"`
	SchuleUID         string `json:"schuleUID"`
}
