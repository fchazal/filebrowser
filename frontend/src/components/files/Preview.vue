<template>
  <div id="previewer">
    <div class="bar">
      <button @click="back" class="action" :title="$t('files.preview.close')" :aria-label="$t('files.closePreview')" id="close">
        <i class="material-icons">arrow_back</i>
      </button>

      <div class="title">{{ this.name }}</div>

      <download-button :disabled="loading" v-if="user.perm.download"></download-button>

      <button @click="openMore" id="more" :aria-label="$t('buttons.more')" :title="$t('buttons.more')" class="action">
        <i class="material-icons">more_vert</i>
      </button>

      <div id="menu" v-show="false" :class="{ active : showMore }">
        <rename-button :disabled="loading" v-if="user.perm.rename"></rename-button>
        <delete-button :disabled="loading" v-if="user.perm.delete"></delete-button>
        <download-button :disabled="loading" v-if="user.perm.download"></download-button>
        <info-button :disabled="loading"></info-button>
      </div>
    </div>

    <div class="loading" v-if="loading">
      <svg class="spinner" viewBox="0 0 50 50">
        <circle class="path" cx="25" cy="25" r="20" fill="none" stroke-width="5"></circle>
      </svg>
    </div>

    <button class="action" @click="prev" v-show="hasPrevious" :aria-label="$t('buttons.previous')" :title="$t('buttons.previous')">
      <i class="material-icons">chevron_left</i>
    </button>
    <button class="action" @click="next" v-show="hasNext" :aria-label="$t('buttons.next')" :title="$t('buttons.next')">
      <i class="material-icons">chevron_right</i>
    </button>

    <template v-if="!loading">
      <div class="preview" @click="this.back">
        <ExtendedImage v-if="req.type == 'image'" :src="raw"></ExtendedImage>
        <audio v-else-if="req.type == 'audio'" :src="raw" autoplay controls></audio>
        <video v-else-if="req.type == 'video'" :src="raw" autoplay controls>
          <track
            kind="captions"
            v-for="(sub, index) in subtitles"
            :key="index"
            :src="sub"
            :label="'Subtitle ' + index" :default="index === 0">
          Sorry, your browser doesn't support embedded videos,
          but don't worry, you can <a :href="download">download it</a>
          and watch it with your favorite video player!
        </video>
        <embed v-else-if="req.extension.toLowerCase() == '.pdf'" class="pdf" :src="raw+'#toolbar=0&navpanes=0'"/>
        <embed v-else-if="req.type == 'text'" class="text" :src="raw" type="text/plain" />
  
       <a v-else :href="download">
          <h2 class="message">{{ $t('buttons.download') }} <i class="material-icons">file_download</i></h2>
        </a>
      </div>
    </template>

    <div v-show="showMore" @click="resetPrompts" class="overlay"></div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import url from '@/utils/url'
import { baseURL, resizePreview } from '@/utils/constants'
import { files as api } from '@/api'
import InfoButton from '@/components/files/buttons/Info'
import DeleteButton from '@/components/files/buttons/Delete'
import RenameButton from '@/components/files/buttons/Rename'
import DownloadButton from '@/components/files/buttons/Download'
import ExtendedImage from './ExtendedImage'
/*
const mediaTypes = [
  "image",
  "video",
  "audio",
  "blob"
]
*/
export default {
  name: 'preview',
  components: {
    InfoButton,
    DeleteButton,
    RenameButton,
    DownloadButton,
    ExtendedImage
  },
  data: function () {
    return {
      previousLink: '',
      nextLink: '',
      listing: null,
      name: '',
      subtitles: [],
      fullSize: true //preview is too slow on small servers
    }
  },
  computed: {
    ...mapState(['req', 'user', 'oldReq', 'jwt', 'loading', 'show']),
    hasPrevious () {
      return (this.previousLink !== '')
    },
    hasNext () {
      return (this.nextLink !== '')
    },
    download () {
      return `${baseURL}/api/raw${url.encodePath(this.req.path)}?auth=${this.jwt}`
    },
    previewUrl () {
      if (this.req.type === 'image' && !this.fullSize) {
        return `${baseURL}/api/preview/big${url.encodePath(this.req.path)}?auth=${this.jwt}`
      }
      return `${baseURL}/api/raw${url.encodePath(this.req.path)}?auth=${this.jwt}`
    },
    raw () {
      return `${this.previewUrl}&inline=true`
    },
    showMore () {
      return this.$store.state.show === 'more'
    },
    isResizeEnabled () {
      return resizePreview
    }
  },
  watch: {
    $route: function () {
      this.updatePreview()
    }
  },
  async mounted () {
    window.addEventListener('keydown', this.key)
    this.$store.commit('setPreviewMode', true)
    this.listing = this.oldReq.items
    this.$root.$on('preview-deleted', this.deleted)
    this.updatePreview()
  },
  beforeDestroy () {
    window.removeEventListener('keydown', this.key)
    this.$store.commit('setPreviewMode', false)
    this.$root.$off('preview-deleted', this.deleted)
  },
  methods: {
    deleted () {
      this.listing = this.listing.filter(item => item.name !== this.name)

      if (this.hasNext) {
        this.next()
      } else if (!this.hasPrevious && !this.hasNext) {
        this.back()
      } else {
        this.prev()
      }
    },
    back () {
      this.$store.commit('setPreviewMode', false)
      let uri = url.removeLastDir(this.$route.path) + '/'
      this.$router.push({ path: uri })
    },
    prev () {
      this.$router.push({ path: this.previousLink })
    },
    next () {
      this.$router.push({ path: this.nextLink })
    },
    key (event) {
      if (this.show !== null) {
        return
      }

      switch (event.which) {
        case 27: // escape
          this.back()
          break
        case 37: // left arrow
          if (this.hasPrevious) this.prev()
          break
        case 39: // right arrow
          if (this.hasNext) this.next()
          break
      }
    },
    async updatePreview () {
      if (this.req.subtitles) {
        this.subtitles = this.req.subtitles.map(sub => `${baseURL}/api/raw${sub}?auth=${this.jwt}&inline=true`)
      }

      let dirs = this.$route.fullPath.split("/")
      this.name = decodeURIComponent(dirs[dirs.length - 1])

      if (!this.listing) {
        try {
          const path = url.removeLastDir(this.$route.path)
          const res = await api.fetch(path)
          this.listing = res.items
        } catch (e) {
          this.$showError(e)
        }
      }

      this.previousLink = ''
      this.nextLink = ''

      for (let i = 0; i < this.listing.length; i++) {
        if (this.listing[i].name !== this.name) {
          continue
        }

        for (let j = i - 1; j >= 0; j--) {
          if (!this.listing[j].isDir) {//mediaTypes.includes(this.listing[j].type)) {
            this.previousLink = this.listing[j].url
            break
          }
        }

        for (let j = i + 1; j < this.listing.length; j++) {
          if (!this.listing[j].isDir) {//mediaTypes.includes(this.listing[j].type)) {
            this.nextLink = this.listing[j].url
            break
          }
        }

        return
      }
    },
    openMore () {
      this.$store.commit('showHover', 'more')
    },
    resetPrompts () {
      this.$store.commit('closeHovers')
    },
    toggleSize () {
      this.fullSize = !this.fullSize
    }
  }
}
</script>

<style scoped>
  .loading {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
</style>