# Hit-der-Woche

## Build

[![Go-Build](https://github.com/MichaelHolley/HitDerWoche/actions/workflows/go-build.yml/badge.svg)](https://github.com/MichaelHolley/HitDerWoche/actions/workflows/go-build.yml)
[![Frontend CI](https://github.com/MichaelHolley/HitDerWoche/actions/workflows/vue-build.yml/badge.svg)](https://github.com/MichaelHolley/HitDerWoche/actions/workflows/vue-build.yml)

## Deployment

[![Backend Docker CD](https://github.com/MichaelHolley/HitDerWoche/actions/workflows/docker-image.yml/badge.svg)](https://github.com/MichaelHolley/HitDerWoche/actions/workflows/docker-image.yml)
[![Deploy GH-pages](https://github.com/MichaelHolley/HitDerWoche/actions/workflows/gh-pages-deploy.yml/badge.svg)](https://github.com/MichaelHolley/HitDerWoche/actions/workflows/gh-pages-deploy.yml)

go-lang types were borrowed from this project <https://github.com/zmb3/spotify>

### Deploy with `docker-compose.yml`

´´´yml
version: '3'
services:
  backend:
    image: mpholley/hitderwoche-backend:latest
    ports:
      - 5000:5000
    environment:
      - PLAYLIST_ID=4RnajpiFvsWCxU3xqwXXrc
      - SPOTIFY_CLIENT_ID=
      - SPOTIFY_CLIENT_SECRET=
      - DB_PASSWORD=mysecretpassword
    restart: always
    depends_on:
      - db
  db:
    image: mysql:latest
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: mysecretpassword
      MYSQL_DATABASE: hitderwoche
    ports:
      - 3306:3306
    volumes:
      - data:/var/lib/mysql
volumes:
  data:
´´´
