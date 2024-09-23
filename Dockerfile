# Use uma imagem base do Go
FROM golang:1.23-alpine

# Defina o diretório de trabalho
WORKDIR /app

# Copie go.mod e go.sum e instale as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copie o restante do código-fonte
COPY . .

# Compile o aplicativo
RUN go build -o main .

# Exponha a porta em que o aplicativo irá rodar
EXPOSE 3000

# Comando para rodar o aplicativo
CMD ["./main"]
