<template>
  <div>
    <router-link :to="uri">
      <div
        :style="'padding-left: '+level*1.5+'rem'"
        :aria-label="name"
        :aria-level="level"
        :class="'item action' + (this.opened?' opened':'') + (this.selected?' selected':'')"
        @click="update">
        <i class="chevron material-icons">arrow_right</i>
        <i class="icon material-icons ">{{ icon }}</i>
        <p class="name">{{ name }}</p>
      </div>
    </router-link>

    <div
      :key="item.name" v-for="item in items"
      :data-url="item.url">
      <folder-tree :uri="item.url" :name="item.name" :level="level + 1"></folder-tree>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import { files } from '@/api'

export default {
  name: 'folder-tree',
  props: {
    'uri': String,
    'name': String,
    'icon': { 
      type: String,
      default: 'folder'
    },
    'level': {
      type: Number,
      default: 0
    }
  },
  data: function () {
    return {
      items: [],
      touches: {
        id: '',
        count: 0
      },
      opened: false,
      selected: false,
      current: window.location.pathname
    }
  },
  computed: {
    ...mapState([ 'req', 'user' ]),
    nav () {
      return decodeURIComponent(this.current)
    }
  },
  mounted () {
    //console.log("req" + this.req)
    //this.fillOptions(this.req)
  },
  methods: {
    update () {
      if (!this.opened) {
        files.fetch(this.uri)
          .then(this.fillOptions)
          .catch(this.$showError)

        this.opened = true
      }
      else {
        this.items = []
        this.opened = false
      }
    },
    fillOptions (req) {
      // Sets the current path and resets
      // the current items.
      this.current = req.url
      this.items = []

      this.$emit('update:selected', this.current)

      // If this folder is empty, finish here.
      if (req.items === null) return

      // Otherwise we add every directory to the
      // move options.
      for (let item of req.items) {
        if (!item.isDir) continue

        this.items.push({
          name: item.name,
          url: item.url
        })
      }
    },
    next: function (event) {
      // Retrieves the URL of the directory the user
      // just clicked in and fill the options with its
      // content.
      let uri = event.currentTarget.dataset.url

      files.fetch(uri)
        .then(this.fillOptions)
        .catch(this.$showError)
    },
    touchstart (event) {
      let url = event.currentTarget.dataset.url

      // In 300 milliseconds, we shall reset the count.
      setTimeout(() => {
        this.touches.count = 0
      }, 300)

      // If the element the user is touching
      // is different from the last one he touched,
      // reset the count.
      if (this.touches.id !== url) {
        this.touches.id = url
        this.touches.count = 1
        return
      }

      this.touches.count++

      // If there is more than one touch already,
      // open the next screen.
      if (this.touches.count > 1) {
        this.next(event)
      }
    },
    itemClick: function (event) {
      if (this.user.singleClick) this.next(event)
      else this.select(event)
    },
    select: function (event) {
      // If the element is already selected, unselect it.
      if (this.selected === event.currentTarget.dataset.url) {
        this.selected = null
        this.$emit('update:selected', this.current)
        return
      }

      // Otherwise select the element.
      this.selected = event.currentTarget.dataset.url
      this.$emit('update:selected', this.selected)
    }
  }
}
</script>
