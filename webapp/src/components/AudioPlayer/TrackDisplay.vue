<script lang="ts" setup>
  import type { Track } from '@/types';

  defineProps<{ currentTrack?: Track }>();
</script>

<template>
  <span v-if="!currentTrack"></span>
  <div v-else class="flex flex-row justify-start items-center">
    <p>
      <span class="text-stone-600">{{ '#' + currentTrack.playlist_position }}</span>
      <span class="mx-2 font-bold">{{ currentTrack.name }}</span>
      <span class="text-stone-600">{{ currentTrack.artists }}</span>
    </p>
  </div>
</template>
