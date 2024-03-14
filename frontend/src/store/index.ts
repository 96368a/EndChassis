import { ExecCommand } from 'wailsjs/go/main/Utils.js'

export interface Command {
  command: string
  isRunning: boolean
  result: string
}

export const store = reactive({
  commandQueue: [] as Command[],
  addCommandResult(index: number, result: string) {
    // console.log('addCommandResult', index, result)
    store.commandQueue[index].result += result
  },
  setCommandStatus(index: number, status: boolean) {
    store.commandQueue[index].isRunning = status
  },
  addCommand(command: string) {
    this.commandQueue.push({
      command,
      isRunning: true,
      result: '',
    })
    ExecCommand(command, this.commandQueue.length - 1)
    // EventsEmit("", this.commandQueue.length - 1)
  },
  isCommandRunning() {
    return store.commandQueue.some(command => command.isRunning)
  },
})
