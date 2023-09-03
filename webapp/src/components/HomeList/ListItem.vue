<script setup lang="ts">
  import { type Track } from '@/types';
  import { formatDuration } from '@/utils/datetime';
  import { PlayIcon, ArrowUturnRightIcon } from '@heroicons/vue/24/outline';

  defineProps<{ track: Track; playTrack: (track: Track) => void }>();
</script>

<template>
  <div class="flex flex-row p-3 relative hover:bg-stone-900 rounded">
    <div>
      <div class="h-28 w-28">
        <img :src="track.image_url" loading="lazy" class="image-full rounded" />
      </div>
    </div>
    <div class="w-full px-6 z-10">
      <div class="text-2xl font-medium text-primary flex flex-row justify-between">
        <div>{{ track.name }}</div>
        <div class="flex flex-row">
          <div
            class="mr-6 tooltip hover:cursor-pointer"
            data-tip="Vorschau"
            :onclick="() => playTrack(track)"
          >
            <PlayIcon
              v-if="!!track.preview_url && track.preview_url != ''"
              class="h-6 hover:scale-105"
            />
          </div>
          <div class="tooltip" data-tip="Spotify">
            <a :href="track.uri" target="_blank">
              <ArrowUturnRightIcon class="h-6 hover:scale-105" />
            </a>
          </div>
        </div>
      </div>
      <div class="w-full text-lg">
        {{ track.artists }}
      </div>
      <div class="text-stone-500">
        {{ formatDuration(track.duration_ms) }}
      </div>
    </div>
    <div class="text-5xl p-0.5 font-thin absolute bottom-4 right-6 rounded text-stone-800">
      {{ track.playlist_position }}
    </div>
  </div>
</template>
