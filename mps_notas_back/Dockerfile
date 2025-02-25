# Build stage
FROM golang:1.24-alpine AS builder

# Instalar dependências de build
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copiar apenas os arquivos necessários primeiro, remover comentários caso tenha dependências, no momento não coloquei nenhuma
# COPY go.sum  ./
# COPY go.mod  ./
# RUN go mod download

# Copiar o resto do código
COPY . .

# Build com flags de otimização
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Adicionar certificados CA e timezone
RUN apk --no-cache add ca-certificates tzdata

# Copiar o binário compilado
COPY --from=builder /app/server .

# Configurar usuário não-root
RUN notas -D notas
USER notas

# Expor porta
EXPOSE 8080

# Configurar healthcheck
HEALTHCHECK --interval=30s --timeout=3s \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Comando para executar o servidor
CMD ["./server"]