#!/bin/bash
# Inicializa o banco de dados e cria a tabela books se não existir

if [ ! -f /app/database.db ]; then
    sqlite3 /app/books.db < /app/init_db.sql
fi

# Executa a aplicação principal
/app/main
