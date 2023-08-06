<script lang="ts" setup>
  import type { Track } from '@/types';

  defineProps<{ currentTrack?: Track }>();
</script>

<template>
  <div>
    <span class="text-stone-600">{{ '#' + currentTrack?.playlist_position }}</span>
    <span class="mx-2 font-bold">{{ currentTrack?.name }}</span>
    <span class="text-stone-600">{{ currentTrack?.artists }}</span>
  </div>
</template>
