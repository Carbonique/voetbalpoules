package voetbalpoules

import (
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/Carbonique/voetbalpoules/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type bericht struct {
	titel  string
	inhoud string
}

type Bot struct {
	*tgbotapi.BotAPI
	chat int64
}

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func NewBot(token string, chat int64) *Bot {
	tbot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	client := client.NewClient("https://www.voetbalpoules.nl/")
	tbot.Debug = false
	log.Printf("Authorized on account %s", tbot.Self.UserName)
	b := Bot{tbot, chat}
	return &b
}

func (bot *Bot) verzend(b *bericht) {
	//
	msg := tgbotapi.NewMessage(bot.chat, b.titel+"\n"+b.inhoud)
	msg.ParseMode = "Markdown"
	msg.DisableWebPagePreview = true
	bot.Send(msg)

}

func (bot *Bot) StuurVoorspelling(w client.Wedstrijd, titelAanvulling string, alsHyperlink bool) {

	bericht := newVoorspellingBericht(&w, titelAanvulling, alsHyperlink)
	bot.verzend(bericht)
}

func (bot *Bot) StuurUitslag(w client.Wedstrijd) {

	bericht := newUitslagBericht(&w)
	bot.verzend(bericht)
}

func (bot *Bot) StuurStand(p *Pool) {

	bericht := newStandBericht(p)
	bot.verzend(bericht)
}

func newStandBericht(p *Pool) *bericht {
	b := bericht{}
	dag := strconv.Itoa(time.Now().Day())
	maand := strconv.Itoa(int(time.Now().Month()))
	b.titel = "*Stand " + dag + "-" + maand + ":* \n"
	sorteerStand(p)
	for _, d := range p.deelnemers {
		p := strconv.Itoa(d.positie) + ". " + d.naam + " " + strconv.Itoa(d.puntenTotaal) + " (+" + strconv.Itoa(d.puntenRonde) + ") \n"
		b.inhoud = b.inhoud + p
	}
	return &b
}

func sorteerStand(p *Pool) {
	sort.Slice(p.deelnemers, func(i, j int) bool {
		return p.deelnemers[i].positie < p.deelnemers[j].positie
	})
}

func newVoorspellingBericht(w *Wedstrijd, titelAanvulling string, alsHyperlink bool) *bericht {
	b := bericht{}
	uur := strconv.Itoa(w.Datum.Hour())
	minuut := strconv.Itoa(w.Datum.Minute())

	if minuut == "0" {
		minuut = minuut + "0"
	}
	b.titel = "*Voorspelling " + titelAanvulling + ":\n" + w.ThuisTeam + " - " + w.UitTeam + " " + uur + ":" + minuut + "* \n"
	w.sorteerVoorspellingen()

	b.inhoud = voorspellingenToString(w, alsHyperlink)

	return &b
}

func newUitslagBericht(w *Wedstrijd) *bericht {
	b := bericht{}
	uur := strconv.Itoa(w.Datum.Hour())
	minuut := strconv.Itoa(w.Datum.Minute())

	if minuut == "0" {
		minuut = minuut + "0"
	}

	b.titel = "*Uitslag:\n" + w.ThuisTeam + " - " + w.UitTeam + " (" + w.Uitslag + ") " + uur + ":" + minuut + "*\n"
	w.Datum.Minute()
	if w.Wvdw {
		b.titel = b.titel + "(" + w.ThuisDoelpuntenMaker + " - " + w.UitDoelpuntenMaker + ") \n"
	}

	w.sorteerVoorspellingen()

	b.inhoud = voorspellingenToString(w, false)

	return &b
}

func (w *Wedstrijd) sorteerVoorspellingen() {
	var t []client.Voorspelling
	var u []Voorspelling
	var g []Voorspelling
	var n []Voorspelling

	for _, v := range w.Voorspellingen {
		if v.UitslagIngevuld {

			switch {
			case v.ThuisDoelpunten > v.UitDoelpunten:
				t = append(t, v)
			case v.UitDoelpunten > v.ThuisDoelpunten:
				u = append(u, v)
			case v.ThuisDoelpunten == v.UitDoelpunten:
				g = append(g, v)
			}
		} else {
			n = append(n, v)
		}
	}
	sorteerThuisVoorspellingen(t)
	sorteerUitVoorspellingen(u)
	sorteerGelijkspelVoorspellingen(g)

	w.Voorspellingen = append(append(append(t, u...), g...), n...)

}

func bepaalWitRegels(w *Wedstrijd) []int {
	var t int
	var u int
	var g int

	for _, v := range w.Voorspellingen {
		if v.UitslagIngevuld {

			switch {
			case v.ThuisDoelpunten > v.UitDoelpunten:
				t++
			case v.UitDoelpunten > v.ThuisDoelpunten:
				u++
			case v.ThuisDoelpunten == v.UitDoelpunten:
				g++
			}
		}
	}

	//Het betreft hier positie van witregels die NA een groep voorspellingen
	//Geplaatst wordt.
	thuisWitRegel := t - 1
	uitWitRegel := thuisWitRegel + u
	gelijkWitRegel := uitWitRegel + g

	return []int{thuisWitRegel, uitWitRegel, gelijkWitRegel}
}

func voorspellingenToString(w *Wedstrijd, gebruikerAlsHyperlink bool) string {
	var voorspellingenTekst string
	wr := bepaalWitRegels(w)

	for i, v := range w.Voorspellingen {

		uitslag := v.getUitslag()
		deelnemer := v.Deelnemer.naam

		if gebruikerAlsHyperlink {
			deelnemer = "[" + v.Deelnemer.naam + "](" + baseUrl + "deelnemer/" + strconv.Itoa(v.Deelnemer.id) + "/voorspellingen/" + w.Competitie + "/" + w.RondeSubDirectory + ")"
		}

		var doelpuntenMakers string
		if w.Wvdw {
			doelpuntenMakers = v.getDoelpuntenMakers()
		}

		voorspellingTekst := deelnemer + " " + uitslag + " " + doelpuntenMakers
		//Als er een witregel moet komen
		if contains(wr, i) {
			voorspellingTekst = voorspellingTekst + "\n"
		}
		voorspellingenTekst = voorspellingenTekst + voorspellingTekst + "\n"

	}
	return voorspellingenTekst

}

func (v Voorspelling) getDoelpuntenMakers() string {
	return "(" + v.ThuisDoelpuntenMaker + " - " + v.UitDoelpuntenMaker + ")"
}

func (v Voorspelling) getUitslag() string {
	if !v.UitslagIngevuld {
		return "(Niet ingevuld)"
	}
	return "(" + strconv.Itoa(v.ThuisDoelpunten) + " - " + strconv.Itoa(v.UitDoelpunten) + ")"
}

func contains(s []int, i int) bool {
	for _, v := range s {
		if i == v {
			return true
		}
	}
	return false
}

func sorteerThuisVoorspellingen(v []Voorspelling) {
	sort.Slice(v, func(i, j int) bool {
		return v[i].ThuisDoelpunten < v[j].ThuisDoelpunten
	})

	sort.Slice(v, func(i, j int) bool {
		return v[i].UitDoelpunten < v[j].UitDoelpunten
	})

}

func sorteerUitVoorspellingen(v []Voorspelling) {
	sort.Slice(v, func(i, j int) bool {
		return v[i].UitDoelpunten < v[j].UitDoelpunten
	})

	sort.Slice(v, func(i, j int) bool {
		return v[i].ThuisDoelpunten < v[j].ThuisDoelpunten
	})

}

func sorteerGelijkspelVoorspellingen(v []Voorspelling) {
	sort.Slice(v, func(i, j int) bool {
		return v[i].ThuisDoelpunten < v[j].ThuisDoelpunten
	})

}
