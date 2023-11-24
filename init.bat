docker-compose up -d --build
docker cp tfb_trigger.sql postgres:/var/backups
docker exec postgres pg_restore -d team_finder /var/backups/tfb_trigger.sql --username postgres -c
pause