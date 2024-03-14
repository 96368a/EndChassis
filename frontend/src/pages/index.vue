<script setup lang="ts" generic="T extends any, O extends any">
import { store } from '~/store'

defineOptions({
  name: 'IndexPage',
})

const name = ref('')

const router = useRouter()
function go() {
  if (name.value)
    router.push(`/hi/${encodeURIComponent(name.value)}`)
}

const command = ref('')
function execute() {
  store.addCommand(command.value)
  command.value = ''
}
</script>

<template>
  <div>
    <div i-carbon-campsite inline-block text-4xl />
    <p>
      <a rel="noreferrer" href="https://github.com/antfu/vitesse-lite" target="_blank">
        Vitesse Lite
      </a>
    </p>
    <p>
      <em text-sm op75>Opinionated Vite Starter Template</em>
    </p>

    <div py-4 />

    <TheInput v-model="name" placeholder="What's your name?" autocomplete="false" @keydown.enter="go" />

    <div>
      <button class="m-3 text-sm btn" :disabled="!name" @click="go">
        Go
      </button>
      <var-button @click="() => console.log(store.commandQueue)">
        说你好
      </var-button>
    </div>
    <div class="min-h-40 flex flex-col">
      <div class="flex-1 bg-gray-800 p-4 pb-10 text-white" rounded>
        <div v-for="(c, i) in store.commandQueue" :key="i">
          <div flex gap-1>
            <span class="text-green-400">guest@tailwindcss:</span>
            <span class="text-blue-400">~$</span>
            <span ml-2>{{ c.command }}</span>
          </div>
          <pre class="whitespace-pre-wrap text-left">{{ c.result }}</pre>
        </div>

        <div v-show="!store.isCommandRunning()">
          <div flex gap-1>
            <span class="text-green-400">guest@tailwindcss:</span>
            <span class="text-blue-400">~$</span>
            <input v-model="command" type="text" class="w-full border-none bg-transparent focus:outline-none" ml-2 @keydown.enter="execute">
          </div>
          <pre class="whitespace-pre-wrap text-left" />
        </div>
      </div>
    </div>
  </div>
</template>
