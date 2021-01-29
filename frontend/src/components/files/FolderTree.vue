<template>
  <div>
    <router-link :to="uri">
      <div
        :style="'padding-left: '+level*1.5+'rem'"
        :aria-label="name"
        :aria-level="level"
        :class="'item action' + (this.opened?' opened':'') + (this.selected?' selected':'')"
        @dragover="dragOver"
        @drop="drop">
        <i @click="expand" class="chevron material-icons">arrow_right</i>
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
import * as upload  from '@/utils/upload'

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
      opened: false
    }
  },
  computed: {
    ...mapState([ 'req', 'user' ]),
    selected () {
      return this.req.url === this.uri
    }
  },
  methods: {
    expand (event) {
      event.preventDefault()

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
      this.items = []

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
    dragOver: function (event) {
      event.preventDefault()
      let el = event.target

      for (let i = 0; i < 5; i++) {
        if (!el.classList.contains('item')) {
          el = el.parentElement
        }
      }

      el.style.opacity = 1
    },
    drop: async function (event) {
      event.preventDefault()

      if (this.selectedCount === 0) return

      let el = event.target
      for (let i = 0; i < 5; i++) {
        if (el !== null && !el.classList.contains('item')) {
          el = el.parentElement
        }
      }

      let items = []

      for (let i of this.selected) {
        items.push({
          from: this.req.items[i].url,
          to: this.url + this.req.items[i].name,
          name: this.req.items[i].name
        })
      }      

      //let base = el.querySelector('.name').innerHTML + '/'
      let path = this.uri //$route.path + base
      let baseItems = (await files.fetch(path)).items

      let action = (overwrite, rename) => {
        files.move(items, overwrite, rename).then(() => {
          this.$store.commit('setReload', true)
        }).catch(this.$showError)
      }

      let conflict = upload.checkConflict(items, baseItems)

      let overwrite = false
      let rename = false

      if (conflict) {
        this.$store.commit('showHover', {
          prompt: 'replace-rename',
          confirm: (event, option) => {
            overwrite = option == 'overwrite'
            rename = option == 'rename'

            event.preventDefault()
            this.$store.commit('closeHovers')
            action(overwrite, rename)
          }
        })

        return
      }

      action(overwrite, rename)
    }
  }
}
</script>
