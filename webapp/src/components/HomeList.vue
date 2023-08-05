<script setup lang="ts">
  import { usePlayerStore } from '@/store/player';
  import { useTracksStore } from '@/store/tracks';
  import type { Track } from '@/types';
  import { formatDuration } from '@/utils/datetime';

  const trackStore = useTracksStore();
  const playerStore = usePlayerStore();

  const playTrack = (track: Track) => {
    if (isValidPreviewUrl(track.preview_url)) {
      playerStore.playTrack(track);
    }
  };

  const isValidPreviewUrl = (url: string | undefined) => {
    return !!url && url != '';
  };
</script>

<template>
  <div id="overview" class="min-h-screen">
    <div class="text-4xl font-bold text-center pt-10">Die letzten Songs</div>
    <div v-if="trackStore.isLoading" class="h-40 w-full">
      <span
        class="loading loading-infinity loading-lg text-primary scale-[2.5] h-full m-auto"
        style="display: table"
      ></span>
    </div>
    <div v-else class="mt-8 xl:w-8/12 md:w-10/12 sm:w-11/12 w-full px-3 sm:px-0 mx-auto">
      <ul>
        <li
          v-for="track in trackStore.tracks"
          :key="track.id"
          class="my-4 py-3 hover:bg-stone-900 rounded"
          :class="{ 'hover:cursor-pointer': isValidPreviewUrl(track.preview_url) }"
          @click="() => playTrack(track)"
        >
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
