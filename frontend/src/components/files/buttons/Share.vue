<template>
  <button @click="show" :aria-label="$t('files.actions.share')" :title="$t('files.actions.share')" class="action">
    <i class="material-icons">link</i>
    <span>{{ $t('files.actions.share') }}</span>
  </button>
</template>

<script>
export default {
  name: 'share-button',
  methods: {
    show () {
      this.$store.commit('showHover', 'share')
    }
  }
}
</script>
