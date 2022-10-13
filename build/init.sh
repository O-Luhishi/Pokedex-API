#!/bin/sh
echo 'Running migrations...'
/app/migrate up

echo 'Starting application...'
/app/pokedex_api