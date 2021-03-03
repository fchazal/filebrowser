<template>
	<div id="disk-usage">
    <div class="bar">
      <div class="fill" :style="width()"></div>
    </div>
    <div class="usage">{{ humanSize() }}</div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import filesize from 'filesize'

export default {
  name: 'disk-usage',
  computed: {
    ...mapState(['disk'])
  },
  methods: {
    width: function () {
      let percent = (this.disk.total) ? Math.round(100 * this.disk.usage / this.disk.total) : 0
      return 'width: ' + percent + '%'
    },
    humanSize: function () {
      return filesize(this.disk.usage) + ' / ' + filesize(this.disk.total)
    },
  }
}
</script>

<style scoped>
  #disk-usage {
    margin: 1em;
  }
  .bar {
    overflow: hidden;
    border-radius: .5em;
    height: .5em;
    background-color: var(--main-color-1);
  }
  .fill {
    height: 100%;
    background: var(--accent-color);
  }
  .usage {
    text-align: center;
    margin: .5em;
    font-size: .8em;
  }
</style>