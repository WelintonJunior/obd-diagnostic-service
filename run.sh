#!/bin/bash

# Inicia o servidor gRPC em background
./bin/api_server initGrpcServer &
GRPC_PID=$!

# Espera até que a porta gRPC (ex: 3000) esteja pronta para conexões
echo "Aguardando o servidor gRPC iniciar..."
while ! nc -z localhost 3000; do
  sleep 0.5
done
echo "Servidor gRPC está pronto!"

# Agora inicia o CLI
./bin/api_server initCli &

# Espera todos os processos
wait $GRPC_PID
