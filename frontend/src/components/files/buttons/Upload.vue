<template>
  <button @click="upload" :aria-label="$t('files.actions.upload')" :title="$t('files.actions.upload')" class="action" id="upload-button">
    <i class="material-icons">file_upload</i>
    <span>{{ $t('files.actions.upload') }}</span>
  </button>
</template>

<script>
export default {
  name: 'upload-button',
  methods: {
    upload: function () {
      if (typeof(DataTransferItem.prototype.webkitGetAsEntry) !== 'undefined') {
        this.$store.commit('showHover', 'upload')
      } else {
        document.getElementById('upload-input').click();
      }
    }
  }
}
</script>
