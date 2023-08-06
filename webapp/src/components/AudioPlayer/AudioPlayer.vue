<script setup lang="ts">
  import { usePlayerStore } from '@/store/player';
  import { storeToRefs } from 'pinia';
  import { onMounted, ref, watch } from 'vue';
  import TrackDisplay from './TrackDisplay.vue';
  import { PlayIcon, PauseIcon } from '@heroicons/vue/24/solid';

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
      player.value?.pause();
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
  <div
    class="sticky z-50 border-t-2 bg-stone-800 border-primary transition-all duration-700"
    :class="{
      '-bottom-40 opacity-0': !playerStore.currentTrack.value,
      ' bottom-0 opacity-100': playerStore.currentTrack.value
    }"
  >
    <audio ref="player" hidden :src="playerStore.currentTrack.value?.preview_url"></audio>
    <div class="p-3 flex flex-row justify-between flex-wrap">
      <TrackDisplay
        class="w-full sm:w-1/2"
        :currentTrack="playerStore.currentTrack.value"
      ></TrackDisplay>
      <div class="w-full sm:w-1/2 flex flex-row justify-between">
        <div class="w-1/2 sm:w-auto sm:mb-0 mt-2 sm:mt-0">
          <div
            class="w-auto hover:cursor-pointer flex flex-row justify-center sm:justify-start"
            @click="toggleAudioState()"
          >
            <PlayIcon class="w-8 h-8 p-1 border-2 rounded-full" v-if="audioState === 'paused'" />
            <PauseIcon
              class="w-8 h-8 p-1 border-2 rounded-full"
              :class="{ 'text-stone-600 border-stone-600': audioState === 'blocked' }"
              v-else-if="audioState === 'playing' || audioState === 'blocked'"
            />
          </div>
        </div>
        <div class="w-1/2 sm:w-auto flex justify-center items-center">
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
  </div>
</template>
