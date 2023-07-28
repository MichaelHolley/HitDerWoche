<script setup lang="ts">
  import { reactive } from 'vue'
  import { type Track } from '../types'
  import { formatDuration } from '../utils/datetime'

  let tracks: Track[] = reactive([])

  fetch('http://localhost:5000/tracks')
    .then((resp) => resp.json())
    .then((data) => {
      tracks.push(
        ...(data.items as Track[]).sort((t1, t2) => t1.playlistPosition - t2.playlistPosition)
      )
    })
</script>

<template>
  <div id="overview" class="min-h-screen">
    <div class="text-3xl font-bold text-center pt-10">Die letzten Songs</div>
    <div class="xl:w-8/12 md:w-10/12 sm:w-11/12 w-full px-3 sm:px-0 mx-auto">
      <ul>
        <li
          v-for="track in tracks"
          :key="track.id"
          class="my-4 rounded-2xl py-3 px-6 border-2 border-stone-700 bg-stone-900 hover:bg-stone-800 shadow-sm shadow-green-400 hover:cursor-pointer"
        >
          <div class="flex flex-row">
            <!-- <img :src="track.image" loading="lazy" class="h-28 image-full rounded mr-6" /> -->
            <div>
              <div class="text-2xl font-medium">
                {{ track.name }}
              </div>
              <div class="text-lg">
                {{ track.artists }}
              </div>
              <div>
                {{ formatDuration(track.duration_ms) }}
              </div>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>
