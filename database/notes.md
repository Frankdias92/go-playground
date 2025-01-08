### Estudo dos Tópicos

---

#### 1. **`go env CGO_ENABLED`**
- **O que é CGO?**
  - CGO permite integrar código C com Go. Quando habilitado (`CGO_ENABLED=1`), você pode usar bibliotecas escritas em C, como a implementação SQLite do pacote `github.com/mattn/go-sqlite3`.
  - Quando desabilitado (`CGO_ENABLED=0`), o Go compila código puro sem dependências de bibliotecas C.

- **Como verificar e configurar:**
  ```bash
  go env CGO_ENABLED
  ```
  - Retorna `1` se habilitado ou `0` se desabilitado.
  
  - Para desabilitar:
    ```bash
    export CGO_ENABLED=0
    ```

- **Por que desabilitar?**
  - Portabilidade: Desabilitar CGO torna o binário gerado mais leve e independente.
  - Evita erros em sistemas que não possuem GCC ou bibliotecas C instaladas.

---

#### 2. **`gcc -v`**
- **O que é GCC?**
  - GCC (GNU Compiler Collection) é um compilador para linguagens como C e C++ usado pelo CGO para compilar código C.

- **Por que verificar?**
  - Para usar bibliotecas como `github.com/mattn/go-sqlite3`, o GCC precisa estar instalado e configurado no sistema.

- **Como verificar:**
  ```bash
  gcc -v
  ```
  - Exibe a versão do GCC e informações sobre sua configuração.

- **Instalando GCC**:
  - Em sistemas baseados em Debian/Ubuntu:
    ```bash
    sudo apt update && sudo apt install build-essential
    ```

---

#### 3. **Instalando `github.com/mattn/go-sqlite3`**
- **Comando:**
  ```bash
  go get -u github.com/mattn/go-sqlite3
  ```

- **Explicação:**
  - `-u`: Atualiza pacotes e suas dependências.
  - Este comando baixa e compila a biblioteca SQLite com suporte a CGO.

- **Dicas:**
  - Certifique-se de que o CGO está habilitado.
  - Se encontrar erros relacionados ao GCC, instale ou configure corretamente o compilador.

---

#### 4. **Nunca usar placeholders como `%s`**
- **Problema com placeholders:**
  - Utilizar placeholders como `%s` em consultas SQL (via `fmt.Sprintf`) pode levar a injeções SQL, um risco de segurança.

- **Exemplo inseguro:**
  ```go
  input := "1; DROP TABLE foo;"
  deleteSql := fmt.Sprintf(`DELETE FROM foo WHERE id = %s;`, input)
  db.Exec(deleteSql) // PERIGO!
  ```

- **Solução segura:**
  - Utilize placeholders `?` e parâmetros ao invés de interpolação:
    ```go
    input := 1
    deleteSql := `DELETE FROM foo WHERE id = ?`
    db.Exec(deleteSql, input) // Seguro
    ```

- **Por que `?` é seguro?**
  - A biblioteca SQL prepara a consulta antes de executar, evitando que entradas maliciosas sejam interpretadas como comandos SQL.

---

#### 5. **SQL sem CGO**
- **Biblioteca alternativa:**
  - O pacote `modernc.org/sqlite` é uma implementação em puro Go que não depende de CGO.

- **Instalação:**
  ```bash
  go get -u modernc.org/sqlite
  ```

- **Vantagens:**
  - Funciona em ambientes onde CGO não está disponível ou habilitado.
  - Binários mais portáveis.

- **Exemplo de uso:**
  ```go
  import (
      "database/sql"
      _ "modernc.org/sqlite"
  )

  func main() {
      db, err := sql.Open("sqlite", "./foo.db")
      if err != nil {
          panic(err)
      }
      defer db.Close()

      _, err = db.Exec(`CREATE TABLE IF NOT EXISTS foo (id INTEGER PRIMARY KEY, name TEXT)`)
      if err != nil {
          panic(err)
      }
  }
  ```
