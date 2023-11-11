POSTGRES_SERVICE=${1}

until docker-compose exec "${POSTGRES_SERVICE}" pg_isready; do
 sleep 1
done
sleep 1