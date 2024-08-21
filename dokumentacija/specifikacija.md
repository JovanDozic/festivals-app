# Specifikacija diplomskog rada

## Tema: Sistem za prodaju i menadzment festivala

## TODO:

- [ ] Promeniti naslov da ima smisla, ne sviđa mi se `Prodaja i menadzment festivala`.
- [ ] Mozda napomenuti da se narukvice prodaju po dropovima? Ili da ostavimo da je on-demand znaci kako ko bude hteo da kupi da kupi - TJ. pretpostavljamo da festival nece biti rasprodat u nekoliko sati.

## Opis osnovne funkcionalnosti

1. Novi posetilac festivala (korisnik) kreira njegov nalog na app. Korisnik je u obavezi da verifikuje nalog putem linka koji je poslat na email koji je uneo.
2. Korisnik potom moze da se uloguje na app.
3. Home page aplikacije cemo tek da smislimo sta ce i kako da bude.
4. Korisnik moze da ode na stranicu za festival i da klikne dugme `kupi kartu`. Ukoliko je shop za taj festival otvoren (u smislu da su karte pustene u prodaju) korisnik je redirectovan na stranicu za kupovinu paketa.
5. Korisniku je prikazan portal u koji ima vise koraka. Na vrhu ekrana nalazi se neki counter ili progress bar ili slicno koji pokazuje koliko je korisnik odmakao u procesu kupovine.
   a. Korisnik prvo bira kontinent, drzavu i grad odakle zeli da departuje.
   b. Korisnik bira dostupnu vrstu prevoza za taj grad/drzavu (avion, autobus, voz).
   b2. Korisnik bira termin polaska.
   b3. Korisnik bira termin povratka.
   c. Korisnik bira vrstu smestaja (koji kamp).
   d. Korisniku ce biti prikazane sve opcije za kamp. Korisnik moze svaku od opcija da klikne i vidi vise detalja o njoj, ukljucujuci i slike i detaljniji opis.
   e. Potom, korisniku su ponuđene dve vrste festivalskih karata - regular i comfort tj. VIP. Korisnik bira jednu od njih.
   f. Korisniku je konacno prikazan pregled svih izabranih opcija, i potom moze da ide da "kupi" kartu.
   g. Korisnik unosi sve informacije o sebi, adresi, i nacinu placanja. Posle "placanja", korisnik dobija potvrdu o kupovini (na portalu kao i na email adresu).
   - Payment process nece biti podrzan u aplikaciji. Placanje ce biti simulirano.
   - U svakom od koraka, korisniku ce biti omoguceno da se vrati korak unazad.
   - [opciono] Proces kupovine je vremenski ogranicen an 15 minuta, gde ukoliko to vreme prođe, korisnik ce biti redirectovan nazad na Home Page - i bice mu onemoguceno da udje u shop za taj festival narednih sat vremena.
   - Korisnik nece moci da izađe iz aplikacije i posle nastavi proces.
   - Korisnik moze da kupi samo jednu kartu samo za sebe i svoju opciju kampovanja i prevoza.
6. Nakon uspesne kupovine, administratorskom sistemu ce stici order koji je napravljen.
7. Admini mogu da vide sve kupovine, i sve detalje vezane za njih.
8. Admini imaju opciju da za svaku kupovinu izdaju narukvicu. Svaka narukvica se posebno salje na korisnikovu adresu. Narukvica sadrzi jedinstveni PIN kod koji je zapisan prilikom izdavanja kartice u bazi podataka tako da bude vezan za porudzbinu. Narukvica pored PIN koda, ima i svoj ID kod koji je takođe jedinstven - i nije tajan.
9. Kada admini posalju narukvicu, shipping number ce biti dodeljen na order.
10. Kada narukvica stigne kod korisnika, on moze da je otpakuje i da ode na sajt i tamo je aktivira. Aktivacija se radi tako sto se PIN kod sa narukvice unese na sajtu. On se mora poklopiti sa PIN kodom koji je administrator uneo prilikom izdavanja.
11. Kada je korisnik aktivirao narukvicu, on ima mogucnost da pregleda stanje narukvice i slicno.
12. Korisnik moze da uplati pearls na narukvicu.
    a. 1 pearl = 1.82 EUR.
    b. Korisnik moze da uplati perle u iznosu od 20 do 2000 EUR.
    c. Korisnik unosi koliko eura uplacuje.
    d. Korisnik obavlja transakciju i dobija potvrdu o uplati.
    e. Korisnik moze da proveri stanje perli na narukvici.
13. Admini mogu da dodaju festivale, i da ih ažuriraju.
14. Admini mogu da dodaju pakete za festivale, i da ih ažuriraju.
15. Admini mogu da generisu izvestaje o prodaji, aktivaciji narukvica, i upotrebi perli.

- Svakakvi razni pregledi, izvestaji i slicno ce biti dostupno svim vrstama korisnika u kontekstima koji ti korisnici imaju (korisnik moze da vidi stanje na narukvici, admin moze da vidi sve porudzbine itd.).
- Postojace vrsta korisnika `superadmin`, on ce moci da dodeljuje ostale administratore, i uopsteno da upravlja sistemom i ostalim korisnicima. Međutim, superadmin nece moci da upravlja festivalima, paketima itd.
- Admin moze da upravlja jednim festivalom. Jedan festival moze da ima vise administratora, i jednog vlasnika. Vlasnik je admin koji je kreirao festival. Vlasnik moze da dodeli druge admine za taj festival. Vlasnik moze da kreira pakete, da ih menja, i da uređuje svaki aspekt festivala. Vlasnik moze i da doda festival currency ili da koristi EUR kao default.

## Upit za [čet](https://chatgpt.com/)

Super. Imas sledeci zadatak. Kao svoj diplomski rad, odlucio sam da napravim sistem koji ce omoguciti prodaju i menadzment festivalskih karata/paketa, slicno poput Tomorrowland Global Journey. Tvoj zadatak je da smislis funkcionalni zahtev koji ce opisati ono cime ce se moj diplomski baviti. Dakle, treba da pokrijemo prodaju tih paketa (biranje kontinenta, drzave, grada odakle se polece, vrstu prevoza i termin, vrstu smestaja (hotel/DreamVille-kamp), vrstu vrste smestaja (koji hotel/koji kamp tj. sator itd.), vrstu karte (regular ili comfort-vip)), potom treba pokriti proces kada se narukvice posalju posetiocima na kucne adrese, kada im narukvice stignu - treba napraviti aktivaciju narukvice sa izdatim PIN code na narukvici, onda moze da se uplacuju pare na narukvicu u vidu perli (1 perla je 1.82 EUR). Ovo je prvo sto je meni palo na pamet, ti slobodno dodaj jos nesto sto smatras da je kljucno za ovaj sistem i da ovaj sistem treba da pokrije. Napomena: ovo bi trebalo da bude omoguceno za vise festivala - dakle, ovo treba da moze da se implementira tako da imamo opcije za recimo Tomorrowland 2025, 2026, ili razlicite vikende istog festivala, ili razlicite destinacije. NAPOMENA 2: u prilogu ti saljem primer funkcionalnog zahteva diplomskog rada mog mentora, to pogledaj i pokusaj da napravis nesto slicno za ovaj sistem (iako je tema drasticno drugacija, primer ti saljem radi proforme i da imas u vidu kako to treba da izgleda i do koje mere to treba da bude razvijeno).
