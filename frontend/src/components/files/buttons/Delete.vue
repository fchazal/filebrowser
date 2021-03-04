<template>
  <button @click="show" :aria-label="$t('files.actions.delete')" :title="$t('files.actions.delete')" class="action" id="delete-button">
    <i class="material-icons">delete</i>
    <span>{{ $t('files.actions.delete') }}</span>
  </button>
</template>

<script>
export default {
  name: 'delete-button',
  methods: {
    show: function () {
      this.$store.commit('showHover', 'delete')
    }
  }
}
</script>
