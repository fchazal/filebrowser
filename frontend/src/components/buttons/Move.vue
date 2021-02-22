<template>
  <button @click="show" :aria-label="$t('buttons.move')" :title="$t('buttons.move')" class="action" id="move-button">
    <i class="material-icons">subdirectory_arrow_right</i>
    <span>{{ $t('buttons.moveFile') }}</span>
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
