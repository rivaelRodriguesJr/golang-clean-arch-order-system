
#!/usr/bin/env bash

set -e

BASE_URL="${BASE_URL:-http://localhost:8000}"
ENDPOINT="${BASE_URL}/order"

echo "Iniciando carga de Orders em ${ENDPOINT}"

create_order() {
    local id="$1"
    local price="$2"
    local tax="$3"

    echo "Criando Order ${id}..."

    curl --fail --silent --show-error \
        -X POST "${ENDPOINT}" \
        -H "Content-Type: application/json" \
        -d "{
            \"id\": \"${id}\",
            \"price\": ${price},
            \"tax\": ${tax}
        }"

    echo
}

create_order "0001" 100.50 0.50
create_order "0002" 200.00 1.00
create_order "0003" 300.75 2.50
create_order "0004" 450.00 5.00
create_order "0005" 999.99 10.00

echo "Carga de Orders concluída!"
