<script setup lang="ts">
  import { reactive } from 'vue'

  const data = reactive([
    {
      id: '123',
      title: 'the very cool song title',
      artist: 'the artist',
      releasYear: 1999,
      image: 'https://cataas.com/cat/says/hello%20world!'
    }
  ])
</script>

<template>
  <div id="overview" class="min-h-screen">
    <div class="text-3xl font-bold text-center pt-10">Die letzten Songs</div>
    <div>
      <ul>
        <li
          v-for="song in data"
          :key="song.id"
          class="my-2 border-2 border-primary rounded-2xl py-3 px-6"
        >
          <div class="flex flex-row">
            <img :src="song.image" loading="lazy" class="h-32 image-full rounded mr-6" />
            <div>
              <div class="text-2xl">
                {{ song.title }}
              </div>
              <div class="text-lg">
                {{ song.artist }}
              </div>
              <div>
                {{ song.releasYear }}
              </div>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>
