package voetbalpoulestelegram

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	voetbalpoules "github.com/Carbonique/voetbalpoules/pkg/voetbalpoules-client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

type bericht struct {
	titel  string
	inhoud string
}

type Bot struct {
	*tgbotapi.BotAPI
	chat int64
}

func NewBot(token string, chat int64) *Bot {
	tbot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	tbot.Debug = false
	log.Printf("Authorized on account %s", tbot.Self.UserName)
	b := Bot{tbot, chat}
	return &b
}

func (bot *Bot) verzend(b *bericht) {

	msg := tgbotapi.NewMessage(bot.chat, b.titel+"\n"+b.inhoud)
	msg.ParseMode = "Markdown"
	msg.DisableWebPagePreview = true
	bot.Send(msg)

}

func (bot *Bot) StuurStand(deelnemers []voetbalpoules.Deelnemer) {

	bericht := newStandBericht(deelnemers)
	bot.verzend(bericht)
}

func newStandBericht(deelnemers []voetbalpoules.Deelnemer) *bericht {
	b := bericht{}
	dag := strconv.Itoa(time.Now().Day())
	maand := strconv.Itoa(int(time.Now().Month()))
	b.titel = fmt.Sprintf("*Stand %s - %s :*\n", dag, maand)

	for i, d := range deelnemers {
		stand := fmt.Sprintf("%d. %s %d (+%d) \n", i+1, d.Naam, d.Punten, d.PuntenRonde)
		b.inhoud = b.inhoud + stand
	}
	return &b
}

func (bot *Bot) StuurUitslag(w voetbalpoules.Wedstrijd, v []voetbalpoules.Voorspelling) {

	bericht := newUitslagBericht(w, v)
	bot.verzend(bericht)
}

func newUitslagBericht(w voetbalpoules.Wedstrijd, v []voetbalpoules.Voorspelling) *bericht {
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
	for _, v2 := range v {
		if v2.DoelpuntenThuis != nil && v2.DoelpuntenUit != nil {
			fmt.Printf("voorspelling %d - %d \n", *v2.DoelpuntenThuis, *v2.DoelpuntenUit)
		}
	}
	gesorteerdeVoorspellingen := sorteerVoorspellingen(v)
	for _, v2 := range gesorteerdeVoorspellingen {
		if v2.DoelpuntenThuis != nil && v2.DoelpuntenUit != nil {
			fmt.Printf("voorspelling %d - %d \n", *v2.DoelpuntenThuis, *v2.DoelpuntenUit)
		}
	}
	//	b.inhoud = voorspellingenToString(w, false)

	return &b
}

//func voorspellingenToString(v []voetbalpoules.Voorspelling, gebruikerAlsHyperlink bool) string {
//	var voorspellingenTekst string
//	wr := bepaalWitRegels(w)
//
//	for i, v := range w.Voorspellingen {
//
//		uitslag := v.getUitslag()
//		deelnemer := v.Deelnemer.naam
//
//		if gebruikerAlsHyperlink {
//			deelnemer = "[" + v.Deelnemer.naam + "](" + baseUrl + "deelnemer/" + strconv.Itoa(v.Deelnemer.id) + "/voorspellingen/" + w.Competitie + "/" + w.RondeSubDirectory + ")"
//		}
//
//		var doelpuntenMakers string
//		if w.Wvdw {
//			doelpuntenMakers = v.getDoelpuntenMakers()
//		}
//
//		voorspellingTekst := deelnemer + " " + uitslag + " " + doelpuntenMakers
//		//Als er een witregel moet komen
//		if contains(wr, i) {
//			voorspellingTekst = voorspellingTekst + "\n"
//		}
//		voorspellingenTekst = voorspellingenTekst + voorspellingTekst + "\n"
//
//	}
//	return voorspellingenTekst
//}

func sorteerVoorspellingen(voorspellingen []voetbalpoules.Voorspelling) []voetbalpoules.Voorspelling {
	log.Info("sorteerVoorspellingen")

	var t []voetbalpoules.Voorspelling
	var u []voetbalpoules.Voorspelling
	var g []voetbalpoules.Voorspelling
	var n []voetbalpoules.Voorspelling

	for _, v := range voorspellingen {
		if v.DoelpuntenThuis != nil && v.DoelpuntenUit != nil {

			switch {
			case *v.DoelpuntenThuis > *v.DoelpuntenUit:
				t = append(t, v)
			case *v.DoelpuntenUit > *v.DoelpuntenThuis:
				u = append(u, v)
			case *v.DoelpuntenThuis == *v.DoelpuntenUit:
				g = append(g, v)
			}
		} else {
			n = append(n, v)
		}
	}
	sorteerThuisVoorspellingen(t)
	sorteerUitVoorspellingen(u)
	sorteerGelijkspelVoorspellingen(g)

	temp_v := append(append(append(t, u...), g...), n...)
	return temp_v

}

func sorteerThuisVoorspellingen(v []voetbalpoules.Voorspelling) {
	sort.Slice(v, func(i, j int) bool {
		return *v[i].DoelpuntenThuis < *v[j].DoelpuntenThuis
	})

	sort.Slice(v, func(i, j int) bool {
		return *v[i].DoelpuntenUit > *v[j].DoelpuntenUit
	})

}

func sorteerUitVoorspellingen(v []voetbalpoules.Voorspelling) {
	sort.Slice(v, func(i, j int) bool {
		return *v[i].DoelpuntenUit < *v[j].DoelpuntenUit
	})

	sort.Slice(v, func(i, j int) bool {
		return *v[i].DoelpuntenThuis > *v[j].DoelpuntenThuis
	})

}

func sorteerGelijkspelVoorspellingen(v []voetbalpoules.Voorspelling) {
	sort.Slice(v, func(i, j int) bool {
		return *v[i].DoelpuntenThuis < *v[j].DoelpuntenThuis
	})

}
