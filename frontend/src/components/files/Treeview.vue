<template>
  <div id="treeview">
    <div v-if="req.numDirs > 0" class="items folders">
      <item v-for="(item) in dirs"
        :key="base64(item.name)"
        v-bind:index="item.index"
        v-bind:name="item.name"
        v-bind:isDir="item.isDir"
        v-bind:url="item.url"
        v-bind:modified="item.modified"
        v-bind:type="item.type"
        v-bind:size="item.size">
      </item>
    </div>
  </div>
</template>

<script>
import { mapState, mapMutations } from 'vuex'
import Item from './ListingItem'

export default {
  name: 'treeview',
  components: { Item },
  computed: {
		...mapState(['req', 'selected', 'user', 'show']),
    items () {
      const dirs = []

      this.req.items.forEach((item) => {
        if (item.isDir)
          dirs.push(item)
      })

      return { dirs }
    },
    dirs () {
      return this.items.dirs.slice(0, this.showLimit)
    }
  },
  mounted: function () {
  },
  methods: {
    ...mapMutations([ 'updateUser', 'addSelected' ]),
    base64: function (name) {
      return window.btoa(unescape(encodeURIComponent(name)))
    },
  }
}
</script>
