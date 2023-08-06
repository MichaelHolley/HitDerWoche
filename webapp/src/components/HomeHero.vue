<script setup lang="ts">
  import { ChevronDownIcon } from '@heroicons/vue/24/solid';
</script>

<template>
  <div class="hero min-h-screen" id="hero">
    <div class="hero-content text-center">
      <div class="max-w-md font-medium">
        <h1 class="text-5xl font-bold">Hit der Woche</h1>
        <p class="pt-6 text-xl">Das Party-Büro präsentiert</p>
        <p class="pb-6 text-2xl">die Hit-der-Woche Spotify-Playlist</p>
        <div class="flex flex-row justify-center">
          <a
            class="btn btn-primary mx-2"
            href="https://open.spotify.com/playlist/4RnajpiFvsWCxU3xqwXXrc"
            target="_blank"
            >Zu Spotify</a
          >
          <RouterLink to="preview" class="btn btn-outline btn-primary mx-2">Vorschau</RouterLink>
        </div>
        <div class="my-6 hover:text-primary">
          <a href="#overview" class="flex flex-col justify-center items-center text-center">
            <p class="text-xl">zur Übersicht</p>
            <ChevronDownIcon class="h-10" />
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
  #hero {
    position: relative;
  }

  #hero::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-size: cover;
    background-position: 50%;
    background-image: url(https://media.giphy.com/media/GNaAKmDFKwKEHEG064/giphy.gif);
    filter: brightness(0.5) blur(7px);
  }
</style>
