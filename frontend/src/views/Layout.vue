<template>
  <div id="app" @contextmenu="ctxclick">
    <div id="progress">
      <div v-bind:style="{ width: this.progress + '%' }"></div>
    </div>
    
    <site-header v-if="!isSharing"></site-header>

    <div id="container">
      <sidebar v-if="!isSharing"></sidebar>
      <router-view></router-view>
    </div>

    <prompts></prompts>
  </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex'
import Sidebar from '@/components/Sidebar'
import Prompts from '@/components/prompts/Prompts'
import SiteHeader from '@/components/Header'
import { enableExec } from '@/utils/constants'

export default {
  name: 'layout',
  components: {
    Sidebar,
    SiteHeader,
    Prompts
  },
  computed: {
    ...mapGetters([
      'isLogged',
      'isSharing',
      'progress'
    ]),
    ...mapState([ 'user' ]),
    isExecEnabled: () => enableExec
  },
  watch: {
    '$route': function () {
      this.$store.commit('resetSelected')
      this.$store.commit('multiple', false)
      if (this.$store.state.show !== 'success') this.$store.commit('closeHovers')
    }
  },
  methods: {
    ctxclick (event) {
      event.preventDefault()
    },
  }
}
</script>
