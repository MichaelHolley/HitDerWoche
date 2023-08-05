import type { Track } from '@/types';
import { defineStore } from 'pinia';

export const usePlayerStore = defineStore('player', {
  state: () => ({ currentTrack: <Track | undefined>undefined, volume: 0.15 }),
  actions: {
    playTrack(track: Track) {
      this.currentTrack = track;
    },
    setVolume(volume: number) {
      this.volume = volume;
    }
  }
});
