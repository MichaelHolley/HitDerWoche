<script setup lang="ts">
  import { usePlayerStore } from '@/store/player';
  import { storeToRefs } from 'pinia';
  import { ref } from 'vue';

  const player = ref<HTMLAudioElement>();

  const playerStore = storeToRefs(usePlayerStore());

  function onVolumeChange(event: Event) {
    if (event.target) {
      const targetVolume = Number.parseFloat((event.target as HTMLInputElement).value);
      playerStore.volume.value = targetVolume;
    }
  }
</script>

<template>
  <div class="sticky bottom-0 z-50 border-t-2 bg-stone-800 border-primary">
    <audio
      ref="player"
      :controls="false"
      hidden
      :src="playerStore.currentTrack.value?.preview_url"
    ></audio>
    <div class="p-3 flex flex-row justify-between">
      <div>{{ playerStore.currentTrack.value?.name }}</div>
      <div>
        <div>{{ playerStore.playing.value ? 'SPIELT' : 'STUMM' }}</div>
        <input type="checkbox" v-model="playerStore.playing.value" />
      </div>
      <div>
        <input
          type="range"
          min="0"
          max="1"
          :step="0.02"
          class="range range-xs range-primary"
          v-model="playerStore.volume.value"
          @input="onVolumeChange"
        />
      </div>
    </div>
  </div>
</template>
