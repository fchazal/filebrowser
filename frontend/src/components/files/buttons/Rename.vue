<template>
  <button @click="show" :aria-label="$t('files.actions.rename')" :title="$t('files.actions.rename')" class="action" id="rename-button">
    <i class="material-icons">mode_edit</i>
    <span>{{ $t('files.actions.rename') }}</span>
  </button>
</template>

<script>
export default {
  name: 'rename-button',
  methods: {
    show: function () {
      this.$store.commit('showHover', 'rename')
    }
  }
}
</script>
