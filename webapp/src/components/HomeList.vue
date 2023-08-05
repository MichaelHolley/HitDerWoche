<script setup lang="ts">
  import { reactive } from 'vue'
  import { type Track } from '../types'
  import { formatDuration } from '../utils/datetime'

  let tracks: Track[] = reactive([])

  fetch(`${import.meta.env.VITE_API_URL}/api/tracks`)
    .then((resp) => resp.json())
    .then((data) => {
      console.log(data.items[0].imageUrl)

      tracks.push(
        ...(data.items as Track[]).sort((t1, t2) => t2.playlist_position - t1.playlist_position)
      )
    })
</script>

<template>
  <div id="overview" class="min-h-screen">
    <div class="text-4xl font-bold text-center pt-10">Die letzten Songs</div>
    <div class="mt-8 xl:w-8/12 md:w-10/12 sm:w-11/12 w-full px-3 sm:px-0 mx-auto">
      <ul>
        <li v-for="track in tracks" :key="track.id" class="my-4 py-3 hover:bg-stone-900 rounded">
          <div class="flex flex-row relative">
            <div class="pl-3">
              <div class="h-28 w-28">
                <img :src="track.image_url" loading="lazy" class="image-full rounded" />
              </div>
            </div>
            <div class="w-full flex flex-col sm:flex-row justify-between px-6 z-10">
              <div>
                <div class="text-2xl font-medium text-primary">
                  {{ track.name }}
                </div>
                <div class="w-full text-lg">
                  {{ track.artists }}
                </div>
                <div class="text-stone-500">
                  {{ formatDuration(track.duration_ms) }}
                </div>
              </div>
            </div>
            <div class="text-5xl p-0.5 font-thin absolute bottom-0 right-6 rounded text-stone-800">
              {{ track.playlist_position }}
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>
