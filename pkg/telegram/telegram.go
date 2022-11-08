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

type deelnemerVoorspelling struct {
	deelnemer       voetbalpoules.Deelnemer
	doelpuntenThuis int
	doelpuntenUit   int
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

func (bot *Bot) verzend(b bericht) {

	msg := tgbotapi.NewMessage(bot.chat, b.titel+"\n"+b.inhoud)
	msg.ParseMode = "Markdown"
	msg.DisableWebPagePreview = true
	bot.Send(msg)

}

func (bot *Bot) StuurStand(deelnemers []voetbalpoules.Deelnemer) {

	bericht := newStandBericht(deelnemers)
	bot.verzend(bericht)
}

func newStandBericht(deelnemers []voetbalpoules.Deelnemer) bericht {
	b := bericht{}
	dag := strconv.Itoa(time.Now().Day())
	maand := strconv.Itoa(int(time.Now().Month()))
	b.titel = fmt.Sprintf("*Stand %s - %s:*\n", dag, maand)

	for i, d := range deelnemers {
		stand := fmt.Sprintf("%d. %s %d (+%d) \n", i+1, d.Naam, d.Punten, d.PuntenRonde)
		b.inhoud = b.inhoud + stand
	}
	return b
}

func (bot *Bot) StuurVoorspelling(vw voetbalpoules.VoorspeldeWedstrijd) {

	bericht := newVoorspellingBericht(vw)
	bot.verzend(bericht)
}

func newVoorspellingBericht(vw voetbalpoules.VoorspeldeWedstrijd) bericht {
	b := bericht{}

	b.titel = voorspellingBerichtTitel(vw.Wedstrijd)
	b.inhoud = voorspellingenBericht(vw)

	return b
}

func (bot *Bot) StuurUitslag(vw voetbalpoules.VoorspeldeWedstrijd) {

	bericht := newUitslagBericht(vw)
	bot.verzend(bericht)
}

func newUitslagBericht(vw voetbalpoules.VoorspeldeWedstrijd) bericht {
	b := bericht{}

	b.titel = uitslagBerichtTitel(vw.Wedstrijd)
	b.inhoud = voorspellingenBericht(vw)

	return b
}

func uitslagBerichtTitel(w voetbalpoules.Wedstrijd) string {
	uur := strconv.Itoa(w.Datum.Hour())
	minuut := strconv.Itoa(w.Datum.Minute())

	if minuut == "0" {
		minuut = minuut + "0"
	}

	titel := fmt.Sprintf("*Uitslag:\n %s - %s (%s) %s:%s*\n", w.ThuisTeam, w.UitTeam, w.Uitslag, uur, minuut)

	if w.Wvdw {
		titel = fmt.Sprintf("%s (%s - %s \n", titel, w.ThuisDoelpuntenMaker, w.UitDoelpuntenMaker)
	}
	return titel

}

func voorspellingBerichtTitel(w voetbalpoules.Wedstrijd) string {
	uur := strconv.Itoa(w.Datum.Hour())
	minuut := strconv.Itoa(w.Datum.Minute())

	if minuut == "0" {
		minuut = minuut + "0"
	}

	titel := fmt.Sprintf("*Voorspelling:\n %s - %s %s:%s*\n", w.ThuisTeam, w.UitTeam, uur, minuut)

	return titel

}

func voorspellingenBericht(vw voetbalpoules.VoorspeldeWedstrijd) string {
	gesorteerdeDeelnemers := sortDeelnemers(vw.DeelnemerVoorspellingen)
	wr := bepaalWitRegels(vw.DeelnemerVoorspellingen)

	inhoud := voorspellingenToString(vw, gesorteerdeDeelnemers, wr)

	return inhoud
}

func voorspellingenToString(vw voetbalpoules.VoorspeldeWedstrijd, volgorde []voetbalpoules.Deelnemer, witregels []int) string {
	var voorspellingenTekst string
	m := vw.DeelnemerVoorspellingen

	for i, d := range volgorde {
		var uitslag string
		if m[d].DoelpuntenThuis == nil && m[d].DoelpuntenUit == nil {
			uitslag = "(Niet ingevuld)"
		} else {
			uitslag = fmt.Sprintf("(%d - %d)", *m[d].DoelpuntenThuis, *m[d].DoelpuntenUit)
		}
		deelnemer := d.Naam

		var doelpuntenMakers string
		if vw.Wedstrijd.Wvdw {
			var thuisDoelpuntenMaker string
			var uitDoelpuntenMaker string
			if m[d].ThuisDoelpuntenMaker == nil && m[d].UitDoelpuntenMaker == nil {
				thuisDoelpuntenMaker = "Niet ingevuld"
				uitDoelpuntenMaker = "Niet ingevuld"
			} else {
				thuisDoelpuntenMaker = *m[d].ThuisDoelpuntenMaker
				uitDoelpuntenMaker = *m[d].UitDoelpuntenMaker
			}
			doelpuntenMakers = fmt.Sprintf("(%s - %s", thuisDoelpuntenMaker, uitDoelpuntenMaker)
		}

		voorspellingTekst := fmt.Sprintf("%s %s %s", deelnemer, uitslag, doelpuntenMakers)

		//Als er een witregel moet komen
		if contains(witregels, i) {
			voorspellingTekst = voorspellingTekst + "\n"
		}
		voorspellingenTekst = voorspellingenTekst + voorspellingTekst + "\n"
	}

	return voorspellingenTekst
}

func bepaalWitRegels(m map[voetbalpoules.Deelnemer]voetbalpoules.Voorspelling) []int {
	var t int
	var u int
	var g int

	for _, v := range m {
		if v.DoelpuntenThuis != nil && v.DoelpuntenUit != nil {

			switch {
			case *v.DoelpuntenThuis > *v.DoelpuntenUit:
				t++
			case *v.DoelpuntenUit > *v.DoelpuntenThuis:
				u++
			case *v.DoelpuntenThuis == *v.DoelpuntenUit:
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

func contains(s []int, i int) bool {
	for _, v := range s {
		if i == v {
			return true
		}
	}
	return false
}

func sortDeelnemers(m map[voetbalpoules.Deelnemer]voetbalpoules.Voorspelling) []voetbalpoules.Deelnemer {

	var t []deelnemerVoorspelling
	var u []deelnemerVoorspelling
	var g []deelnemerVoorspelling
	var n []deelnemerVoorspelling

	for d, v := range m {
		if v.DoelpuntenThuis != nil && v.DoelpuntenUit != nil {

			dv := deelnemerVoorspelling{
				deelnemer:       d,
				doelpuntenThuis: *v.DoelpuntenThuis,
				doelpuntenUit:   *v.DoelpuntenUit,
			}

			switch {
			case *v.DoelpuntenThuis > *v.DoelpuntenUit:
				t = append(t, dv)
			case *v.DoelpuntenUit > *v.DoelpuntenThuis:
				u = append(u, dv)
			case *v.DoelpuntenThuis == *v.DoelpuntenUit:
				g = append(g, dv)
			}
		} else {
			n = append(n, deelnemerVoorspelling{deelnemer: d})
		}
	}
	sorteerThuisVoorspellingen(t)
	sorteerUitVoorspellingen(u)
	sorteerGelijkspelVoorspellingen(g)

	deelnemerVoorspellingen := append(append(append(t, u...), g...), n...)
	var deelnemers []voetbalpoules.Deelnemer
	for _, dv := range deelnemerVoorspellingen {
		deelnemers = append(deelnemers, dv.deelnemer)
	}

	return deelnemers

}

func sorteerThuisVoorspellingen(dv []deelnemerVoorspelling) {
	sort.Slice(dv, func(i, j int) bool {
		return dv[i].doelpuntenThuis < dv[j].doelpuntenThuis
	})

	sort.Slice(dv, func(i, j int) bool {
		return dv[i].doelpuntenUit < dv[j].doelpuntenUit
	})

}

func sorteerUitVoorspellingen(dv []deelnemerVoorspelling) {
	sort.Slice(dv, func(i, j int) bool {
		return dv[i].doelpuntenUit < dv[j].doelpuntenUit
	})

	sort.Slice(dv, func(i, j int) bool {
		return dv[i].doelpuntenThuis < dv[j].doelpuntenThuis
	})

}

func sorteerGelijkspelVoorspellingen(dv []deelnemerVoorspelling) {
	sort.Slice(dv, func(i, j int) bool {
		return dv[i].doelpuntenThuis < dv[j].doelpuntenThuis
	})

}
