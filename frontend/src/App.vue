<script setup lang="ts">
import { EventsOn } from 'wailsjs/runtime/runtime.js'
import { store } from './store'

EventsOn('addCommandResult', store.addCommandResult)
EventsOn('setCommandStatus', store.setCommandStatus)
</script>

<template>
  <main font-sans p="x-4 y-10" text="center gray-700 dark:gray-200">
    <RouterView />
    <TheFooter />
  </main>
</template>
