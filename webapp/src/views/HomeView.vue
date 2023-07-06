<script setup lang="ts">
  import DownArrowIcon from '../components/icons/IconDownArrow.vue'
</script>

<template>
  <main>
    <div class="hero min-h-screen">
      <div class="hero-content text-center">
        <div class="max-w-md">
          <h1 class="text-5xl font-bold">Hit der Woche</h1>
          <p class="py-6">Das Party-Büro präsentiert die Hit-der-Woche Spotify-Playlist</p>
          <div class="flex flex-row justify-center">
            <a
              class="btn btn-primary mx-1"
              href="https://open.spotify.com/playlist/4RnajpiFvsWCxU3xqwXXrc"
              target="_blank"
              >Zu Spotify</a
            >
            <RouterLink to="preview" class="btn btn-outline btn-primary mx-1">Vorschau</RouterLink>
          </div>
          <div class="my-6 hover:text-primary">
            <a href="#overview" class="flex flex-col justify-center items-center">
              <p class="mr-3">Übersicht</p>
              <DownArrowIcon
            /></a>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
