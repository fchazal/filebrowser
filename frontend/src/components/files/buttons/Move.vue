<template>
  <button @click="show" :aria-label="$t('files.actions.move')" :title="$t('files.actions.move')" class="action" id="move-button">
    <i class="material-icons">subdirectory_arrow_right</i>
    <span>{{ $t('files.actions.move') }}</span>
  </button>
</template>

<script>
export default {
  name: 'move-button',
  methods: {
    show: function () {
      this.$store.commit('showHover', 'move')
    }
  }
}
</script>
