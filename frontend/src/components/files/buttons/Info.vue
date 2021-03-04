<template>
  <button :title="$t('files.actions.info')" :aria-label="$t('files.actions.info')" class="action" @click="show">
    <i class="material-icons">info</i>
    <span>{{ $t('files.actions.info') }}</span>
  </button>
</template>

<script>
export default {
  name: 'info-button',
  methods: {
    show: function () {
      this.$store.commit('showHover', 'info')
    }
  }
}
</script>
