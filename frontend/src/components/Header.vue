<template>
  <header v-if="!isEditor && !isPreview">
    <div class="left">
      <button v-if="isMobile" @click="openSidebar" :aria-label="$t('buttons.toggleSidebar')" :title="$t('buttons.toggleSidebar')" class="-mobile action">
        <i class="material-icons">menu</i>
      </button>

      <template v-if="!isMobile">
        <img :src="logoURL" :alt="name">
        <h2>{{ name }}</h2>
      </template>
    </div>

    <!--
      <search v-if="isLogged"></search>
        <button v-show="!isSharing" @click="openSearch" :aria-label="$t('buttons.search')" :title="$t('buttons.search')" class="search-button action">
          <i class="material-icons">search</i>
        </button>
    -->

    <div class="right">
      <router-link class="action" to="/settings" :aria-label="$t('sidebar.settings')" :title="$t('sidebar.settings')">
        <i class="material-icons">settings</i>
      </router-link>

      <button v-if="authMethod == 'json'" @click="logout" class="action" id="logout" :aria-label="$t('sidebar.logout')" :title="$t('sidebar.logout')">
        <i class="material-icons">exit_to_app</i>
      </button>
    </div>
  </header>
</template>

<script>
//import Search from './Search'
import { mapGetters, mapState } from 'vuex'
import { name, logoURL } from '@/utils/constants'
import * as auth from '@/utils/auth'
import { authMethod } from '@/utils/constants'

export default {
  name: 'header-layout',
  components: {
//    Search,
  },
  data: () => { return {
    'width': window.innerWidth
  }},
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
    isMobile() {
      return this.width <= 736
    }
  },
  methods: {
    openSidebar () {
      this.$store.commit('showHover', 'sidebar')
    },
    /*
    openSearch () {
      this.$store.commit('showHover', 'search')
    },
    */
    logout: auth.logout
  }
}
</script>
