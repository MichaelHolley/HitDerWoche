import type { Track } from '@/types';
import { defineStore } from 'pinia';

export const usePlayerStore = defineStore('player', {
  state: () => ({ currentTrack: <Track | undefined>undefined, volume: 0.5, playing: false }),
  actions: {
    playTrack(track: Track) {
      this.currentTrack = track;
    },
    setVolume(volume: number) {
      this.volume = volume;
    },
    togglePlayState(playState?: boolean) {
      if (playState != undefined) {
        this.playing = playState;
      } else {
        this.playing = !this.playing;
      }
    }
  }
});
