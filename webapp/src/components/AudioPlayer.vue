<script setup lang="ts">
  import { usePlayerStore } from '@/store/player';
  import { storeToRefs } from 'pinia';
  import { onMounted, ref, watch } from 'vue';

  type AudioState = 'playing' | 'paused' | 'blocked';

  const player = ref<HTMLAudioElement>();
  const audioState = ref<AudioState>('blocked');

  const playerStore = storeToRefs(usePlayerStore());

  onMounted(() => {
    if (player.value) {
      player.value.volume = playerStore.volume.value;
    }
  });

  watch(playerStore.currentTrack, (newVal) => {
    if (!newVal) {
      audioState.value = 'blocked';
    } else {
      setTimeout(() => {
        player.value?.play().then(() => {
          audioState.value = 'playing';
        });
      }, 250);
    }
  });

  watch(playerStore.volume, (newVal) => {
    if (player.value) {
      player.value.volume = newVal;
    }
  });

  function onVolumeChange(event: Event) {
    if (event.target) {
      const targetVolume = Number.parseFloat((event.target as HTMLInputElement).value);
      playerStore.volume.value = targetVolume;
    }
  }

  function toggleAudioState() {
    if (audioState.value === 'playing') {
      player.value?.pause();
      audioState.value = 'paused';
      return;
    }

    if (audioState.value === 'paused') {
      player.value?.play();
      audioState.value = 'playing';
      return;
    }
  }
</script>

<template>
  <div class="sticky bottom-0 z-50 border-t-2 bg-stone-800 border-primary">
    <audio ref="player" hidden :src="playerStore.currentTrack.value?.preview_url"></audio>
    <div class="p-3 flex flex-row justify-between">
      <div>{{ playerStore.currentTrack.value?.name }}</div>
      <div>
        <div @click="toggleAudioState()">{{ audioState === 'playing' ? 'SPIELT' : 'STUMM' }}</div>
      </div>
      <div class="flex flex-row">
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
