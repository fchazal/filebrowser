<template>
  <button @click="show" :aria-label="$t('files.actions.copy')" :title="$t('files.actions.copy')" class="action" id="copy-button">
    <i class="material-icons">content_copy</i>
    <span>{{ $t('files.actions.copy') }}</span>
  </button>
</template>

<script>
export default {
  name: 'copy-button',
  methods: {
    show: function () {
      this.$store.commit('showHover', 'copy')
    }
  }
}
</script>
