<template>
  <nav id="sidebar" :class="{active}">
    <template v-if="!isMobile">
      <router-link :class="'action' + (this.$store.state.route.name === 'Files' ? ' active':'')" to="/files/" :aria-label="$t('sidebar.home')" :title="$t('sidebar.home')">
        <i class="material-icons">cloud</i>
        <span>{{ $t('sidebar.home') }}</span>
      </router-link>

      <div id="hierarchy">
      </div>

      <div v-for="(drive, index) in this.user.drives" :key="index">
        <router-link class="action" :to="'/files/' + drive.name" :aria-label="drive.name" :title="drive.name">
          <i class="material-icons">{{ drive.icon }}</i>
          <span>{{ drive.name }}</span>
        </router-link>
      </div>

      <div class="separator"></div>

      <router-link :class="'action' + (this.$store.state.route.name === 'Shares' ? ' active':'')" to="/shares" :aria-label="$t('sidebar.shares')" :title="$t('sidebar.shares')">
        <i class="material-icons">folder_shared</i>
        <span>{{$t('sidebar.shares')}}</span>
      </router-link>

      <router-link v-if="user.perm.admin" :class="'action' + (this.$store.state.route.name === 'Users' ? ' active':'')" to="/users" :aria-label="$t('sidebar.users')" :title="$t('sidebar.users')">
        <i class="material-icons">people</i>
        <span>{{$t('sidebar.users')}}</span>
      </router-link>
    </template>
  </nav>
</template>

<script>
import { mapState, mapGetters } from 'vuex'
import { version, signup, disableExternal, noAuth, authMethod } from '@/utils/constants'

export default {
  name: 'sidebar',
  data: () => { return {
    'width': window.innerWidth
  }},
  created () {
    window.addEventListener('resize', () => {
      this.width = window.innerWidth
    })
  },
  computed: {
    ...mapState([ 'user' ]),
    ...mapGetters([
      'isLogged',
      'isSharing'
    ]),
    active () {
      return this.$store.state.show === 'sidebar'
    },
    signup: () => signup,
    version: () => version,
    disableExternal: () => disableExternal,
    noAuth: () => noAuth,
    authMethod: () => authMethod,
    isMobile() {
      return this.width <= 736
    }
  },
  methods: {

  }
}
</script>
