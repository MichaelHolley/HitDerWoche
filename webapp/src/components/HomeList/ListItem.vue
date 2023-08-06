<script setup lang="ts">
  import { type Track } from '@/types';
  import { formatDuration } from '@/utils/datetime';
  import { PlayIcon } from '@heroicons/vue/24/outline';

  defineProps<{ track: Track }>();
</script>

<template>
  <div class="flex flex-row p-3 relative hover:bg-stone-900 rounded">
    <div>
      <div class="h-28 w-28">
        <img :src="track.image_url" loading="lazy" class="image-full rounded" />
      </div>
    </div>
    <div class="w-full px-6 z-10">
      <div class="text-2xl font-medium text-primary">
        {{ track.name }}
      </div>
      <div class="w-full text-lg">
        {{ track.artists }}
      </div>
      <div class="text-stone-500">
        {{ formatDuration(track.duration_ms) }}
      </div>
      <div
        class="text-stone-600 flex flex-row"
        v-if="!!track.preview_url && track.preview_url != ''"
      >
        <PlayIcon class="h-6" /> Vorschau verfügbar
      </div>
    </div>
    <div class="text-5xl p-0.5 font-thin absolute bottom-4 right-6 rounded text-stone-800">
      {{ track.playlist_position }}
    </div>
  </div>
</template>
