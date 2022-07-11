package voetbalpoules

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gocolly/colly/v2"
)

var serverIndexResponse = []byte("hello world\n")

func TestGetWedstrijden(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	t1 := time.Now().Local()
	t2 := t1.AddDate(0, 0, 1)

	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)
	comp := NewCompetitie("ek_vrouwen_2022")
	w := comp.GetWedstrijden(c, ts.URL, t1, t2)
	fmt.Println(w)

}

func newTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(serverIndexResponse)
	})

	mux.HandleFunc("/wedstrijd/index/ek_vrouwen_2022", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
		<!DOCTYPE html>
		<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="nl" lang="nl">
		<head>
				<title>EK Vrouwen 2022</title>
				<meta name="keywords" content="ek pool, ek poule, ek vrouwen 2022, oranje, voorspel, gratis, pooltje, voetbaltoto, uitslagen, standen, teams, voetbal, european cup" />
				<meta name="description" content="De leukste EK Poule 2020! Voorspel de uitslagen van het EK 2022, start je eigen EK pool en speel tegen je vrienden, familie of collega's" />
				<meta http-equiv="content-type" content="text/html; charset=ISO-8859-1" />
				<link rel="apple-touch-icon-precomposed" sizes="129x129" href="/Content/images/0/icon-mobile.png">
				<link rel="shortcut icon" href="/favicon.ico">
		
				<script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
				<script type="text/javascript">(window.jQuery || document.write('<script src="/Content/jquery-1.11.0.min.js"><\/script>'))</script>
				<link href="/csso?v=6KXZupFKUhW4DVxuqoZ3aAde5Y_1Audj8YZwhJ7BUik1" rel="stylesheet"/>
		
				<script src="/js?v=2ezxhG8wCbT4_2rWXakjlvcJHpXC9v9-x9XF6oOxqTs1"></script>
		
				<meta property="og:title" content="EK Vrouwen 2022" />
				<meta property="og:type" content="game" />
				<meta property="og:url" content="https://www.voetbalpoules.nl" />
				<meta property="og:image" content="https://www.voetbalpoules.nl/Content/images/84/ekvoorspellen-og.png" />
				<meta property="og:site_name" content="Voetbalpoules.nl" />
				<meta property="og:description" content="Voorspel het EK Vrouwen 2022 met vrienden, familie of collega&#39;s!" />
				<meta property="og:locale" content="nl_NL" />
				<meta property="fb:app_id" content="189009751163294" />
				<script type="text/javascript">
				$(function() {
					var d3 = { seedTime: 1657461220929, language: 'nl', format: '%a %d-%m-%Y | %H:%M:%S' };
					$('#clock').jclock(d3);
				});
				</script>
						<script>
								var msTag = {
										data: { nsc: false },
										site: "voetbalpoules",
										page: "ros"
								}
						</script>
						<script type="text/javascript" src="https://massariuscdn.com/pubs/voetbalpoules/voetbalpoules_hb_setup.js"></script>
		</head>
		<body>
						<div id="fb-root"></div>
						<script>(function(d, s, id) { var js, fjs = d.getElementsByTagName(s)[0]; if (d.getElementById(id)) return; js = d.createElement(s); js.id = id; js.src = "//connect.facebook.net/nl_NL/sdk.js#xfbml=1&version=v2.5&appId=189009751163294"; fjs.parentNode.insertBefore(js, fjs); }(document, 'script', 'facebook-jssdk'));</script>
				<div class="main">
						<div class="header-holder noprint">
								<div class="header">
										<div class="header-left wide">
												<a href="/" class="titel">Voetbalpoules.nl</a>
												<div class="subtitel">
														<span>Met de leukste EK Pool van Nederland!</span>
												</div>
												<div class="datumtekst"><span id="clock"></span> - Nog <span class="tijd">134</span> dagen tot het WK 2022 begint!</div>
										</div>
										<div class="header-right small">
		
		
			<h2 class="naam">Hoi <a class="wit" href="/deelnemer/880081/gegevens">Jupiler League</a></h2>
				<div class="rechts">
		<a class="wit" data-ajax="false" href="/uitloggen">Uitloggen</a>		</div>
			<div style="clear:both;"></div>
			<div class="box donker info">
				<div class="punten">0 punten (+0)</div>
				<div class="positie">Positie: - (was -)</div>
			</div>
		
										</div>
										<div style="clear: both;"></div>
								</div>
						</div>
						<div class="menu-holder noprint">
								<div class="menu">
										<div class="menu-links">
		
		
		<ul class="menu noprint">
				<li>
		<a href="/deelnemer/880081/voorspellingen/ek_vrouwen_2022">Voorspellen</a>					</li>
				<li>
						<a href="#">
							Mijn...
					</a>
									<ul>
								<li>
										<a href="/deelnemer/880081/poules">
											Pools
										</a>
								</li>
								<li>
										<a href="/deelnemer/880081/positie">
											Posities
										</a>
								</li>
								<li>
										<a href="/deelnemer/880081/gegevens">
											Gegevens
										</a>
								</li>
								<li>
										<a href="/deelnemer/880081/statistieken/ek_vrouwen_2022">
											Statistieken
										</a>
								</li>
								<li>
										<a href="/deelnemer/880081/competities">
											Competities
										</a>
								</li>
						</ul>
				</li>
				<li>
		<a href="/prikbord">Prikbord</a>					</li>
				<li>
						<a href="#">
							Speluitleg
					</a>
									<ul>
								<li>
										<a href="/spelregels">
											Spelregels
										</a>
								</li>
								<li>
										<a href="/puntentelling">
											Puntentelling
										</a>
								</li>
								<li>
										<a href="/faq">
											Vragen (FAQ)
										</a>
								</li>
								<li>
										<a href="/prijzen">
											Prijzen
										</a>
								</li>
						</ul>
				</li>
				<li>
						<a href="#">
							Ranglijsten
					</a>
									<ul>
								<li>
										<a href="/stand">
											Individueel
										</a>
								</li>
								<li>
										<a href="/stand/poules">
											Pools
										</a>
								</li>
						</ul>
				</li>
				<li>
						<a href="#">
							Scorebord
					</a>
									<ul>
								<li>
										<a href="/wedstrijd/index">
											Wedstrijden
										</a>
								</li>
								<li>
										<a href="/competitie">
											Standen
										</a>
								</li>
								<li>
										<a href="/team">
											Teams
										</a>
								</li>
						</ul>
				</li>
		</ul>
		 
										</div>
										<div class="menu-rechts">
												<form id="zoeken" action="/zoeken" class="zoek" method="GET">
														<input type="search" placeholder="Zoek pool of deelnemer" maxlength="20" name="query" />
														<div class="zoek" onclick="$('#zoeken').submit();"></div>
												</form>
												<a class="language" href="https://www.voetbalpoules.nl/wedstrijd/index/ek_vrouwen_2022?language=en">
														<img src="/Content/Desktop/images/en.png" title="to English" />
												</a>
										</div>
								</div>
						</div>
						<div class="content">
								<div class="content-holder">
												<div style="min-height:250px; height:250px;">
														<!-- 13436254/Voetbalpoules_ROS_Header -->
														<div id="div-gpt-ad-voetbalpoules_ros_header" style="display:block;">
																<script>
																		googletag.cmd.push(function () { googletag.display('div-gpt-ad-voetbalpoules_ros_header'); });
																</script>
														</div>
												</div>
																				<div class="content-left">
																		<h1>EK Vrouwen 2022</h1>
		
		
		
		<div class="selectie">
				<select id="competitieNaam" onchange="window.location.href = '/wedstrijd/index/_competitie_'.replace('_competitie_', this.value);">
						<option value="overall" >Alle competities</option>
						<option value="bundesliga">
							Bundesliga
						</option>
						<option value="champions_league">
							Champions League
						</option>
						<option value="conference_league">
							Conference League
						</option>
						<option value="derde_divisie_zaterdag">
							Derde divisie (za)
						</option>
						<option value="derde_divisie_zondag">
							Derde divisie (zo)
						</option>
						<option value="dfb_pokal">
							DFB-Pokal
						</option>
						<option value="eerste_divisie">
							Eerste Divisie
						</option>
						<option value="ek_vrouwen_2022" selected>
							EK Vrouwen 2022
						</option>
						<option value="eredivisie">
							Eredivisie
						</option>
						<option value="europa_league">
							Europa League
						</option>
						<option value="fa_cup">
							FA Cup
						</option>
						<option value="jupiler_pro_league">
							Jupiler Pro League
						</option>
						<option value="knvb_beker">
							KNVB Beker
						</option>
						<option value="ligue_1">
							Ligue 1
						</option>
						<option value="nations_league">
							Nations League
						</option>
						<option value="premier_league">
							Premier League
						</option>
						<option value="primera_division">
							Primera Division
						</option>
						<option value="serie_a">
							Serie A
						</option>
						<option value="tweede_divisie">
							Tweede divisie
						</option>
						<option value="wk_2022">
							WK 2022
						</option>
						<option value="wk_kwalificatie_2022">
							WK Kwalificatie 2022
						</option>
		
			</select>
		
			<div style="clear:both"></div>
		</div>
		
			<table class="wedstrijden">
				<colgroup>
					<col style="width: 24%;">
					<col style="width: 32%;">
					<col style="width: 32%;">
					<col style="width: 12%;">
				</colgroup>
				<tr>
					<th>Datum</th>
					<th>Thuisteam</th>
					<th>Uitteam</th>
					<th>Uitslag</th>
				</tr>
							<tr>
								<td colspan="4">
																<h2>
																		EK Vrouwen 2022
																				<span style="font-size:0.9em;">Groepsfase 1e ronde</span>
																																</h2>
														</td>
												</tr>
		
		<tr class="">
				<td class="nowrap">
						6 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Engeland" src="https://vp-logos.azureedge.net/214221" />
						<label class="vp-team" for="">Engeland</label>
				</td>
				<td title="">
								<img alt="Oostenrijk" src="https://vp-logos.azureedge.net/214222" />
						<label class="vp-team" for="">Oostenrijk</label>
				</td>
				<td class="right nowrap">
								 1 - 0
				</td>
		</tr>
		
		<tr class="wvdw">
				<td class="nowrap">
						7 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Noorwegen" src="https://vp-logos.azureedge.net/214223" />
						<label class="vp-team" for="">Noorwegen</label>
				</td>
				<td title="">
								<img alt="Noord Ierland" src="https://vp-logos.azureedge.net/214224" />
						<label class="vp-team" for="">Noord Ierland</label>
				</td>
				<td class="right nowrap">
								 4 - 1
				</td>
		</tr>
				<tr class="wvdw doelpunten">
				<td>1e doelpunt</td>
				<td>
					Blakstad
				</td>
				<td>
					Nelson
				</td>
				<td></td>
			</tr>
		
		<tr class="">
				<td class="nowrap">
						8 jul	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Spanje" src="https://vp-logos.azureedge.net/214236" />
						<label class="vp-team" for="">Spanje</label>
				</td>
				<td title="">
								<img alt="Finland" src="https://vp-logos.azureedge.net/214237" />
						<label class="vp-team" for="">Finland</label>
				</td>
				<td class="right nowrap">
								 4 - 1
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						8 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Duitsland" src="https://vp-logos.azureedge.net/214238" />
						<label class="vp-team" for="">Duitsland</label>
				</td>
				<td title="">
								<img alt="Denemarken" src="https://vp-logos.azureedge.net/214239" />
						<label class="vp-team" for="">Denemarken</label>
				</td>
				<td class="right nowrap">
								 4 - 0
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						Gisteren	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Portugal" src="https://vp-logos.azureedge.net/215720" />
						<label class="vp-team" for="">Portugal</label>
				</td>
				<td title="">
								<img alt="Zwitserland" src="https://vp-logos.azureedge.net/214241" />
						<label class="vp-team" for="">Zwitserland</label>
				</td>
				<td class="right nowrap">
								 2 - 2
				</td>
		</tr>
		
		<tr class="wvdw">
				<td class="nowrap">
						Gisteren	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Nederland" src="https://vp-logos.azureedge.net/214242" />
						<label class="vp-team" for="">Nederland</label>
				</td>
				<td title="">
								<img alt="Zweden" src="https://vp-logos.azureedge.net/214243" />
						<label class="vp-team" for="">Zweden</label>
				</td>
				<td class="right nowrap">
								 1 - 1
				</td>
		</tr>
				<tr class="wvdw doelpunten">
				<td>1e doelpunt</td>
				<td>
					Roord
				</td>
				<td>
					Andersson
				</td>
				<td></td>
			</tr>
		
		<tr class="">
				<td class="nowrap">
						Vandaag	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Belgi&#235;" src="https://vp-logos.azureedge.net/214244" />
						<label class="vp-team" for="">Belgi&#235;</label>
				</td>
				<td title="">
								<img alt="IJsland" src="https://vp-logos.azureedge.net/214245" />
						<label class="vp-team" for="">IJsland</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						Vandaag	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Frankrijk" src="https://vp-logos.azureedge.net/214246" />
						<label class="vp-team" for="">Frankrijk</label>
				</td>
				<td title="">
								<img alt="Itali&#235;" src="https://vp-logos.azureedge.net/214247" />
						<label class="vp-team" for="">Itali&#235;</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
							<tr>
								<td colspan="4">
																<h2>
																		EK Vrouwen 2022
																				<span style="font-size:0.9em;">Groepsfase 2e ronde</span>
																																</h2>
														</td>
												</tr>
		
		<tr class="">
				<td class="nowrap">
						Morgen	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Oostenrijk" src="https://vp-logos.azureedge.net/214222" />
						<label class="vp-team" for="">Oostenrijk</label>
				</td>
				<td title="">
								<img alt="Noord Ierland" src="https://vp-logos.azureedge.net/214224" />
						<label class="vp-team" for="">Noord Ierland</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						Morgen	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Engeland" src="https://vp-logos.azureedge.net/214221" />
						<label class="vp-team" for="">Engeland</label>
				</td>
				<td title="">
								<img alt="Noorwegen" src="https://vp-logos.azureedge.net/214223" />
						<label class="vp-team" for="">Noorwegen</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						12 jul	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Denemarken" src="https://vp-logos.azureedge.net/214239" />
						<label class="vp-team" for="">Denemarken</label>
				</td>
				<td title="">
								<img alt="Finland" src="https://vp-logos.azureedge.net/214237" />
						<label class="vp-team" for="">Finland</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="wvdw">
				<td class="nowrap">
						12 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Duitsland" src="https://vp-logos.azureedge.net/214238" />
						<label class="vp-team" for="">Duitsland</label>
				</td>
				<td title="">
								<img alt="Spanje" src="https://vp-logos.azureedge.net/214236" />
						<label class="vp-team" for="">Spanje</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						13 jul	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Zweden" src="https://vp-logos.azureedge.net/214243" />
						<label class="vp-team" for="">Zweden</label>
				</td>
				<td title="">
								<img alt="Zwitserland" src="https://vp-logos.azureedge.net/214241" />
						<label class="vp-team" for="">Zwitserland</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						13 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Nederland" src="https://vp-logos.azureedge.net/214242" />
						<label class="vp-team" for="">Nederland</label>
				</td>
				<td title="">
								<img alt="Portugal" src="https://vp-logos.azureedge.net/215720" />
						<label class="vp-team" for="">Portugal</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						14 jul	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Itali&#235;" src="https://vp-logos.azureedge.net/214247" />
						<label class="vp-team" for="">Itali&#235;</label>
				</td>
				<td title="">
								<img alt="IJsland" src="https://vp-logos.azureedge.net/214245" />
						<label class="vp-team" for="">IJsland</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="wvdw">
				<td class="nowrap">
						14 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Frankrijk" src="https://vp-logos.azureedge.net/214246" />
						<label class="vp-team" for="">Frankrijk</label>
				</td>
				<td title="">
								<img alt="Belgi&#235;" src="https://vp-logos.azureedge.net/214244" />
						<label class="vp-team" for="">Belgi&#235;</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
							<tr>
								<td colspan="4">
																<h2>
																		EK Vrouwen 2022
																				<span style="font-size:0.9em;">Groepsfase 3e ronde</span>
																																</h2>
														</td>
												</tr>
		
		<tr class="">
				<td class="nowrap">
						15 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Noord Ierland" src="https://vp-logos.azureedge.net/214224" />
						<label class="vp-team" for="">Noord Ierland</label>
				</td>
				<td title="">
								<img alt="Engeland" src="https://vp-logos.azureedge.net/214221" />
						<label class="vp-team" for="">Engeland</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="wvdw">
				<td class="nowrap">
						15 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Oostenrijk" src="https://vp-logos.azureedge.net/214222" />
						<label class="vp-team" for="">Oostenrijk</label>
				</td>
				<td title="">
								<img alt="Noorwegen" src="https://vp-logos.azureedge.net/214223" />
						<label class="vp-team" for="">Noorwegen</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						16 jul	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Finland" src="https://vp-logos.azureedge.net/214237" />
						<label class="vp-team" for="">Finland</label>
				</td>
				<td title="">
								<img alt="Duitsland" src="https://vp-logos.azureedge.net/214238" />
						<label class="vp-team" for="">Duitsland</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						16 jul	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Denemarken" src="https://vp-logos.azureedge.net/214239" />
						<label class="vp-team" for="">Denemarken</label>
				</td>
				<td title="">
								<img alt="Spanje" src="https://vp-logos.azureedge.net/214236" />
						<label class="vp-team" for="">Spanje</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						17 jul	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Zwitserland" src="https://vp-logos.azureedge.net/214241" />
						<label class="vp-team" for="">Zwitserland</label>
				</td>
				<td title="">
								<img alt="Nederland" src="https://vp-logos.azureedge.net/214242" />
						<label class="vp-team" for="">Nederland</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="wvdw">
				<td class="nowrap">
						17 jul	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Zweden" src="https://vp-logos.azureedge.net/214243" />
						<label class="vp-team" for="">Zweden</label>
				</td>
				<td title="">
								<img alt="Portugal" src="https://vp-logos.azureedge.net/215720" />
						<label class="vp-team" for="">Portugal</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						18 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="IJsland" src="https://vp-logos.azureedge.net/214245" />
						<label class="vp-team" for="">IJsland</label>
				</td>
				<td title="">
								<img alt="Frankrijk" src="https://vp-logos.azureedge.net/214246" />
						<label class="vp-team" for="">Frankrijk</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						18 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Itali&#235;" src="https://vp-logos.azureedge.net/214247" />
						<label class="vp-team" for="">Itali&#235;</label>
				</td>
				<td title="">
								<img alt="Belgi&#235;" src="https://vp-logos.azureedge.net/214244" />
						<label class="vp-team" for="">Belgi&#235;</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
							<tr>
								<td colspan="4">
																<h2>
																		EK Vrouwen 2022
																				<span style="font-size:0.9em;">Finaleronde</span>
																																				<span style="font-size:0.8em;">- Kwartfinale</span>
																</h2>
														</td>
												</tr>
		
		<tr class="">
				<td class="nowrap">
						20 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Winnaar groep A" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Winnaar groep A</label>
				</td>
				<td title="">
								<img alt="Nummer 2 groep B" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Nummer 2 groep B</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						21 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Winnaar groep B" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Winnaar groep B</label>
				</td>
				<td title="">
								<img alt="Nummer 2 groep A" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Nummer 2 groep A</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						22 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Winnaar groep C" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Winnaar groep C</label>
				</td>
				<td title="">
								<img alt="Nummer 2 groep D" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Nummer 2 groep D</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						23 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Winnaar groep D" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Winnaar groep D</label>
				</td>
				<td title="">
								<img alt="Nummer 2 groep C" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Nummer 2 groep C</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
							<tr>
								<td colspan="4">
																<h2>
																		EK Vrouwen 2022
																				<span style="font-size:0.9em;">Finaleronde</span>
																																				<span style="font-size:0.8em;">- Halve finale</span>
																</h2>
														</td>
												</tr>
		
		<tr class="">
				<td class="nowrap">
						26 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Winnaar kwartfinale 1" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Winnaar kwartfinale 1</label>
				</td>
				<td title="">
								<img alt="Winnaar kwartfinale 3" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Winnaar kwartfinale 3</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
		<tr class="">
				<td class="nowrap">
						27 jul	<div>21:00</div>
		
				</td>
				<td title="">
								<img alt="Winnaar kwartfinale 2" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Winnaar kwartfinale 2</label>
				</td>
				<td title="">
								<img alt="Winnaar kwartfinale 4" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Winnaar kwartfinale 4</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
							<tr>
								<td colspan="4">
																<h2>
																		EK Vrouwen 2022
																				<span style="font-size:0.9em;">Finaleronde</span>
																																				<span style="font-size:0.8em;">- Finale</span>
																</h2>
														</td>
												</tr>
		
		<tr class="">
				<td class="nowrap">
						31 jul	<div>18:00</div>
		
				</td>
				<td title="">
								<img alt="Winnaar halve finale 1" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Winnaar halve finale 1</label>
				</td>
				<td title="">
								<img alt="Winnaar halve finale 2" src="https://vp-logos.azureedge.net/214212" />
						<label class="vp-team" for="">Winnaar halve finale 2</label>
				</td>
				<td class="right nowrap">
									- 
				</td>
		</tr>
		
			</table>
			<h3>Statistieken</h3>
			<table>
				<tr>
					<th>Competitie</th>
					<th>Wedstrijden</th>
					<th>Doelpunten</th>
				</tr>
					<tr>
						<td>EK Vrouwen 2022</td>
						<td>31</td>
						<td>21</td>
					</tr>
			</table>
								<div style="padding: 15px;">
										<!-- 13436254/Voetbalpoules_ROS_Bottom -->
										<div id="div-gpt-ad-voetbalpoules_ros_bottom" style="display:block;">
												<script>
														googletag.cmd.push(function () { googletag.display('div-gpt-ad-voetbalpoules_ros_bottom'); });
												</script>
										</div>
								</div>
		
												</div>
												<div class="content-right noprint">
																<div class="social">
						<div style="width: 150px"><div class="fb-like" data-width="150" data-href="https://www.facebook.com/voetbalpoules" data-layout="button_count" data-action="like" data-show-faces="false" data-share="true"></div></div>
						<div style="float:right">
								<a title="Volg Voetbalpoules op Facebook" target="_blank" href="https://www.facebook.com/voetbalpoules"><img src="/foto/10570" style="height:20px;width:20px"></a>
						</div>
				</div><div class="box grijs deelnemers">
				<div>Aantal deelnemers</div>
				<div>56.137</div>
				<div style="clear: both;"></div>
		</div>
						<div class="box donker centered">
								<a href="/poule/nieuw">Start eigen pool</a>
						</div>
				<div class="box licht centered">
						<a href="/uitnodigen">Anderen uitnodigen</a>
				</div>
						<div class="box anders centered">
								<a href="/supporter">Word supporter!</a>
						</div>
				<!-- 13436254/Voetbalpoules_ROS_Right -->
				<div id="div-gpt-ad-voetbalpoules_ros_right" style="display:block;">
						<script>
								googletag.cmd.push(function () { googletag.display('div-gpt-ad-voetbalpoules_ros_right'); });
						</script>
				</div>
		<div class="box grijs">
		<h2>Laatste voetbalnieuws</h2>
			<ul class="nieuws">
				<li><a href="https://www.soccernews.nl/news/885868" target="_blank">'Feyenoord ruikt miljoenen en gaat voor verlenging'</a></li>		
				<li><a href="https://www.soccernews.nl/news/885866" target="_blank">'Groot Wijnaldum-nieuws bereikt PSV'</a></li>		
				<li><a href="https://www.soccernews.nl/news/885863" target="_blank">'Lisandro Martínez is furieus op Ajax'</a></li>		
				<li><a href="https://www.soccernews.nl/news/885850" target="_blank">'Conte wijst viertal de deur bij Tottenham'</a></li>		
				<li><a href="https://www.soccernews.nl/news/885843" target="_blank">'Toptarget zet alles op transfer naar Barça'</a></li>		
			</ul>
		
		
		</div>
		<div class="box grijs">
		
		<div id="pollcontent" style="position: relative">
			<div class="loadingDiv" id="loading"></div>
			<h2>Poll</h2>
				<p>Het EK Vrouwen staat voor de deur en onze oranje leeuwinnen hebben de voorbereiding inmiddels erop zitten.<br />Hoe staat &#39;t met de interesse naar dit toernooi want de aandacht bij de media is (opvallend) groot(s) te noemen!</p>
					<table class="nostyle" style="z-index: 500">
							<tr>
								<td>
									<input onclick="" type="radio" name="pollAntwoordID" id="antwoord11958" value="11958">
								</td>
								<td>
									<label for="antwoord11958">Ik ga &#39;t hele EK volgen en heb ook alle voorbereidingswedstrijden van onze leeuwinnen gezien</label>
								</td>
							</tr>
							<tr>
								<td>
									<input onclick="" type="radio" name="pollAntwoordID" id="antwoord11959" value="11959">
								</td>
								<td>
									<label for="antwoord11959">Ik ga wel &#39;t EK volgen maar de voorbereidingswedstrijden alleen maar in samenvattingen gezien</label>
								</td>
							</tr>
							<tr>
								<td>
									<input onclick="" type="radio" name="pollAntwoordID" id="antwoord11960" value="11960">
								</td>
								<td>
									<label for="antwoord11960">Als &#39;t EK gaat beginnen ga ik pas kijken naar onze leeuwinnen maar niet ook nog eens naar alle andere wedstrijden</label>
								</td>
							</tr>
							<tr>
								<td>
									<input onclick="" type="radio" name="pollAntwoordID" id="antwoord11961" value="11961">
								</td>
								<td>
									<label for="antwoord11961">Ik kijk hooguit de samenvattingen in de groepsfase en knock-out wedstrijden misschien wel live</label>
								</td>
							</tr>
							<tr>
								<td>
									<input onclick="" type="radio" name="pollAntwoordID" id="antwoord11962" value="11962">
								</td>
								<td>
									<label for="antwoord11962">Ik kijk niet en zie de uitslagen vanzelf wel ergens in de media verschijnen</label>
								</td>
							</tr>
					</table>
					<script>
						$('input[name=pollAntwoordID]').click(function() {
							$("#loading").css("display", "block");
							$("#pollcontent").addClass("loading");
							$.ajax({
								url: '/layout/pollstem',
								data: { userID: 880081, pollAntwoordID: $('input[name=pollAntwoordID]:checked').val() },
								headers: { 'RequestVerificationToken': 'U_MaRtj3uLegqNCFcLpL-ZU3Pj6zJTNUT7bHRSN-BSF9zdHNnpuGmXj__2dFFI5aqAu2dF8Oo1gG1NLA1cunxdFpKUWnGRNvxAfn5W-NX0U1:WuLSUeeZKDQ6FWXIK4SKbAQu0QwTcOK1vQUhnxtIX77v_L6kH1V8daS5RnS9fDSu2WTuxjcYP9uNxstoC4Uzi5Pg8ghXIvefSNBt_VJ5pKc1OasJFIhHeZ97Musd52e_De9Z8I1wg1rQQzkE7_r-rw2' },
								type: "POST"
							})
							.success(function(data) { $("#pollcontent").html(data); })
							.error(function (data) { alert('Oei, er gaat iets fout!')})
							.complete(function() {
								$("#loading").css("display", "none");
								$("#pollcontent").removeClass("loading");
							});
						});
					</script>
				<p>Er is <b>188</b> keer gestemd.</p>
		</div>
		
		
		</div>
		<div class="box grijs">
		<h2>Sitenieuws</h2>
			<ul class="nieuws">
					<li><a href="/nieuws/2128">Tricky Dicky wint WK Kwalificatie 2022</a></li>		
					<li><a href="/nieuws/2127">Nieuw seizoen → oude seizoen verwijderen</a></li>		
					<li><a href="/nieuws/2126">de topvoorspeller wint Champions League Seizoen 2021-&#39;22</a></li>		
					<li><a href="/nieuws/2125">AchtentachtigRevisited wint Ligue 1 Seizoen 2021-&#39;22</a></li>		
					<li><a href="/nieuws/2124">kittie1964 wint Derde divisie (zo) Seizoen 2021-&#39;22</a></li>		
			</ul>
			<div style="text-align:right">
				<a href="/nieuws">Meer nieuws</a>
			</div>
		
		
		</div>
				<div class="box grijs">
						<h2>Link partners</h2>
						<ul class="nieuws">
								<li>
										<a href="https://bynco.com/nl/occasions/" target="_blank" title="Occasion kopen">Occasion kopen</a>
								</li>
										<li>
												<a href="https://casinonieuws.nl/" target="_blank" title="Casinonieuws.nl">Casinonieuws.nl</a>
										</li>
														<li>
										<a href="https://bynco.com/nl" target="_blank" title="BYNCO">BYNCO</a>
								</li>
										<li>
												<a href="https://www.gokken.com/" target="_blank" title="Gokken.com">Gokken.com</a>
										</li>
														<li>
										<a href="https://justlease.nl/" target="_blank" title="Private Lease">Private Lease</a>
								</li>
																<li>
												<a href="https://casinowiki.nl/" target="_blank" title="Casinowiki.nl">Casinowiki.nl</a>
										</li>
																		</ul>
				</div>
		
		
		
												</div>
												<div style="clear: both"></div>
								</div>
								<div class="footer">
										<div class="beheerder noprint">
												<a href="/prikbord">Prikbord</a>
												<a href="/medewerkers">Medewerkers</a>
												<a href="/privacy">Privacy</a>
												<a href="#" onclick="showConsentManager()"><u>Cookieinstellingen</u></a>
																				</div>
										<div class="c">
												<a class="noprint" href="/switchmode?mode=Mobile&amp;returnUrl=%2Fwedstrijd%2Findex%2Fek_vrouwen_2022" id="tomobile">Mobiele site</a>	<script>
				$(document).on('click', 'a#tomobile', function() {
					event.preventDefault();
					window.location = $(this).attr("href");
				});
			</script>
		 &copy; 2002-2022 <a href="/medewerkers">Voetbalpoules.nl</a> | 7
										</div>
								</div>
						</div>
				</div>
				<script>
				(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
					(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
						m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
				})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
		
				ga('create', 'UA-5350835-1', 'Voetbalpoules.nl');
				
				ga('send', 'pageview');
				</script>
		</body>
		</html>
		`))
	})
	return httptest.NewServer(mux)
}

//
//func TestNewCompetitie(t *testing.T) {
//
//	if got != want {
//		t.Errorf("got %q, wanted %q", got, want)
//	}
//}
