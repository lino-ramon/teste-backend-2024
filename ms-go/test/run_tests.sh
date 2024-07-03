# Variaveis de ambiente para TEST
export MONGO_HOST=localhost
export MONGO_DATABASE=teste_backend
export MAIN_COLLECTION=products_test

# Executar a partir do diretório: ms-go/test
go test ./... -v