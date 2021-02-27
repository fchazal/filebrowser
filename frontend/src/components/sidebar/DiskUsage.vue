<template>
	<div id="disk-usage">
    <div class="bar">
      <div class="fill"></div>
    </div>
    <div class="usage">20.0 GB of 100.0 GB used</div>
  </div>
</template>

<script>
export default {
  name: 'disk-usage',
  methods: {}
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
    width: 20%;
    height: 100%;
    background: var(--accent-color);
  }
  .usage {
    text-align: center;
    margin: .5em;
    font-size: .8em;
  }
</style>