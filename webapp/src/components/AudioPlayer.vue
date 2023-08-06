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

  function getFormatedTrackname() {
    let display = '[#' + playerStore.currentTrack.value?.playlist_position + '] ';
    display += playerStore.currentTrack.value?.name;

    return display;
  }
</script>

<template>
  <div
    class="sticky z-50 border-t-2 bg-stone-800 border-primary transition-all duration-700"
    :class="{
      '-bottom-40 opacity-0': !playerStore.currentTrack.value,
      ' bottom-0 opacity-100': playerStore.currentTrack.value
    }"
  >
    <audio ref="player" hidden :src="playerStore.currentTrack.value?.preview_url"></audio>
    <div class="p-3 flex flex-row justify-between">
      <div class="basis-3/12">{{ getFormatedTrackname() }}</div>
      <div class="basis-6/12 justify-center text-center">
        <div @click="toggleAudioState()">{{ audioState === 'playing' ? 'Play' : 'Pause' }}</div>
      </div>
      <div class="justify-end items-end align-bottom ml-auto">
        <input
          type="range"
          min="0"
          max="1"
          :step="0.02"
          class="range range-xs range-primary w-40"
          v-model="playerStore.volume.value"
          @input="onVolumeChange"
        />
      </div>
    </div>
  </div>
</template>
