<script setup lang="ts">
  import { RouterView } from 'vue-router'
  import { useTracksStore } from '@/store/tracks'

  const tracksStore = useTracksStore()
  tracksStore.refreshTracks()
</script>

<template>
  <RouterView />
</template>
