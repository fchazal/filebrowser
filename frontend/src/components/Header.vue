<template>
  <header v-if="!isEditor && !isPreview">
    <div>
      <button @click="openSidebar" :aria-label="$t('buttons.toggleSidebar')" :title="$t('buttons.toggleSidebar')" class="action">
        <i class="material-icons">menu</i>
      </button>
      <img :src="logoURL" alt="File Browser">
      <h2>{{ name }}</h2>
      <!--<search v-if="isLogged"></search>-->
    </div>
    <div>
      <template v-if="isLogged || isSharing">
        <button v-show="!isSharing" @click="openSearch" :aria-label="$t('buttons.search')" :title="$t('buttons.search')" class="search-button action">
          <i class="material-icons">search</i>
        </button>

        <button @click="openMore" id="more" :aria-label="$t('buttons.more')" :title="$t('buttons.more')" class="action">
          <i class="material-icons">more_vert</i>
        </button>

        <!-- Menu that shows on listing AND mobile when there are files selected -->
        <div id="file-selection" v-if="isMobile && isListing && !isSharing">
          <span v-if="selectedCount > 0">{{ selectedCount }} selected</span>
          <share-button v-show="showShareButton"></share-button>
          <rename-button v-show="showRenameButton"></rename-button>
          <copy-button v-show="showCopyButton"></copy-button>
          <move-button v-show="showMoveButton"></move-button>
          <delete-button v-show="showDeleteButton"></delete-button>
          <switch-button v-show="isListing"></switch-button>
          <info-button v-show="isFiles"></info-button>
        </div>

        <!-- This buttons are shown on a dropdown on mobile phones -->
        <div id="dropdown" :class="{ active: showMore }">
          <!--
          <div v-if="!isListing || !isMobile">
            <share-button v-show="showShareButton"></share-button>
            <rename-button v-show="showRenameButton"></rename-button>
            <copy-button v-show="showCopyButton"></copy-button>
            <move-button v-show="showMoveButton"></move-button>
            <delete-button v-show="showDeleteButton"></delete-button>
          </div>

          <switch-button v-show="isListing"></switch-button>
          <info-button v-show="isFiles"></info-button>
          -->


          <div>
          
            <router-link class="action" to="/settings" :aria-label="$t('sidebar.settings')" :title="$t('sidebar.settings')">
              <i class="material-icons">settings</i>
              <span>{{ $t('sidebar.settings') }}</span>
            </router-link>

            <button v-if="authMethod == 'json'" @click="logout" class="action" id="logout" :aria-label="$t('sidebar.logout')" :title="$t('sidebar.logout')">
              <i class="material-icons">exit_to_app</i>
              <span>{{ $t('sidebar.logout') }}</span>
            </button>
          </div>
        </div>
      </template>

      <div v-show="showOverlay" @click="resetPrompts" class="overlay"></div>
    </div>
  </header>
</template>

<script>
//import Search from './Search'
import {mapGetters, mapState} from 'vuex'
import { name, logoURL, enableExec } from '@/utils/constants'
import * as api from '@/api'
import buttons from '@/utils/buttons'
import * as auth from '@/utils/auth'
import { authMethod } from '@/utils/constants'

export default {
  name: 'header-layout',
  components: {
//    Search,
  },
  data: function () {
    return {
      width: window.innerWidth,
      pluginData: {
        api,
        buttons,
        'store': this.$store,
        'router': this.$router
      }
    }
  },
  created () {
    window.addEventListener('resize', () => {
      this.width = window.innerWidth
    })
  },
  computed: {
    ...mapGetters([
      'selectedCount',
      'isFiles',
      'isEditor',
      'isPreview',
      'isListing',
      'isLogged',
      'isSharing'
    ]),
    ...mapState([
      'req',
      'user',
      'loading',
      'reload',
      'multiple'
    ]),
    authMethod: () => authMethod,
    name: () => name,
    logoURL: () => logoURL,
    isExecEnabled: () => enableExec,
    isMobile () {
      return this.width <= 736
    },
    showMore () {
      return (this.isFiles || this.isSharing) && this.$store.state.show === 'more'
    },
    showOverlay () {
      return this.showMore
    }
  },
  methods: {
    openSidebar () {
      this.$store.commit('showHover', 'sidebar')
    },
    openMore () {
      this.$store.commit('showHover', 'more')
    },
    openSearch () {
      this.$store.commit('showHover', 'search')
    },
    toggleMultipleSelection () {
      this.$store.commit('multiple', !this.multiple)
      this.resetPrompts()
    },
    resetPrompts () {
      this.$store.commit('closeHovers')
    },
    logout: auth.logout
  }
}
</script>
