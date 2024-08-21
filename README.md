napisz program chaat room pozwalający komunikować sie miedzy użytkownikami.
Proam powinien wspierać pokoje i nazwy uzytkowników. 
innstrukcje obsługiwane przez program:
- name ustawia nazwe uzytkownikowi "/name John"
- join <nazwa pokoju> dołącz do pokoju "/join go_room"
- room_list wylistowanie wsztszystkich pokoi "/room_list"
- quit wyjscie z pokoju "/quit"
- msg wysłanie wiadomosci "/msg <tresc wiadomosci>"
Jeśli użytkownik wejdzie do pokoju wszyscy użytkownicy obecni w pokoju powinni otrzymać wiadomosc kto wchodzi do pokoju np John joined to room. Uzytkownik tez powinien otrzymać wiadomosc
"you joined to room: <nazwa>". Przy wyjsciu z pokoju wszyscy pozostali powinni otrzymać wiadomosc "<nik> has left room".Wszystkie instrukcje musza być poprzedzone "/". Jeśli użytkownik chce dołaczyc ddo pokoju który nie istnieje
powinien on zostac utworzony. Jeśli użytkownik chce wysłać wiadomosc a nie jest obecnie w pokoju musi otrzymać informacje "you must join to room first". Program powinien wspierac broadcast na wszystkie pokoje
w celu poinformowania wszystkich uzytkowników o zamknieciu programu.
