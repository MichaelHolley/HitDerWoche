<script setup lang="ts">
  import { usePlayerStore } from '@/store/player';
  import { useTracksStore } from '@/store/tracks';
  import type { Track } from '@/types';
  import ListItem from './ListItem.vue';

  const trackStore = useTracksStore();
  const playerStore = usePlayerStore();

  const playTrack = (track: Track) => {
    if (!!track.preview_url && track.preview_url != '') {
      playerStore.playTrack(track);
    }
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
        <li v-for="track in trackStore.tracks" :key="track.id" class="my-4">
          <ListItem :track="track" :playTrack="(t: Track) => playTrack(t)"></ListItem>
        </li>
      </ul>
    </div>
  </div>
</template>
