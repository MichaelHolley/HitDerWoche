import { defineStore } from 'pinia';
import { type Track } from '@/types';

export const useTracksStore = defineStore('tracks', {
  state: () => ({ tracks: <Track[]>[], isLoading: false }),
  actions: {
    refreshTracks() {
      this.isLoading = true;
      fetch(`${import.meta.env.VITE_API_URL}/api/tracks`)
        .then((resp) => {
          this.isLoading = false;
          return resp.json();
        })
        .then((data) => {
          this.tracks.push(
            ...(data.items as Track[]).sort((t1, t2) => t2.playlist_position - t1.playlist_position)
          );
          this.isLoading = false;
        });
    }
  }
});
