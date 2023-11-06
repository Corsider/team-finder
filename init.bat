docker-compose up -d --build
docker cp tfb.sql postgres:/var/backups
docker exec postgres pg_restore -d team_finder /var/backups/tfb.sql --username postgres -c
pause