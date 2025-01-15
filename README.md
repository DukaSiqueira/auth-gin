go mod init authGin

go  get \
    github.com/gin-gonic/gin \
    github.com/golang-jwt/jwt/v4 \
    github.com/joho/godotenv \
    golang.org/x/crypto \
    gorm.io/driver/postgres \
    gorm.io/gorm \
    github.com/githubnemo/CompileDaemon


O script acima instalou o framework Gin com as seguintes bibliotecas:

JWT : uma biblioteca Go para implementação de JSON Web Token.
DotEnv : uma dependência Go para gerenciar variáveis ​​de ambiente.
Gorm : um ORM (Mapeador Relacional de Objetos) para Golang.
Go Crypto : uma biblioteca de criptografia em Go que será usada para criptografar e descriptografar senhas.
CompileDaemon: um pacote Go que facilita o recarregamento automático do aplicativo ao salvar as alterações. 