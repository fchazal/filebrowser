<template>
  <form class="drives small">
    <div v-for="(drive, index) in drives" :key="index">
      <input
        @keypress.enter.prevent
        type="text"
        v-model="drive.icon"
        :placeholder="$t('settings.insertIcon')" />
      <input
        @keypress.enter.prevent
        type="text"
        v-model="drive.name"
        :placeholder="$t('settings.insertName')" />
      <input
        @keypress.enter.prevent
        type="text"
        v-model="drive.path"
        :placeholder="$t('settings.insertPath')" />

      <button class="button button--red" @click="remove($event, index)">-</button>
    </div>

    <div>
      <button class="button" @click="create" default="false">{{ $t('buttons.new') }}</button>
    </div>
  </form>
</template>

<script>
export default {
  name: 'drives-textarea',
  props: ['drives'],
  methods: {
    remove (event, index) {
      event.preventDefault()
      let drives = [ ...this.drives ]
      drives.splice(index, 1)
      this.$emit('update:drives', [ ...drives ])
    },
    create (event) {
      event.preventDefault()

      this.$emit('update:drives', [
        ...this.drives,
        {
          name: '',
          path: '',
        }
      ])
    }
  }
}
</script>
