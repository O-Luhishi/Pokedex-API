#!/bin/sh
echo 'Running migrations...'
/app/migrate up > /dev/null 2>&1 &

echo 'Starting application...'
/app/pokedex_api