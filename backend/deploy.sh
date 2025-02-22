#!/bin/bash

# Atualizar dependências
bun install

# Reiniciar o processo PM2 (se existir)
if pm2 list | grep -q "mps-backend"; then
    pm2 restart mps-backend
else
    pm2 start dist/index.js --name mps-backend
fi