<template>
  <aside id="sidebar" :class="{active}">
    <div id="drive">
      <folder-tree uri="/files/" :name="$t('sidebar.home')" icon="cloud"></folder-tree>
    </div>
<!--
    <div id="trash">
      <folder-tree uri="/trash/" :name="$t('sidebar.trash')" icon="delete"></folder-tree>
    </div>
-->
    <div class="separator"></div>

    <router-link :class="'action' + (this.$store.state.route.name === 'Shares' ? ' active':'')" to="/shares" :aria-label="$t('sidebar.shares')" :title="$t('sidebar.shares')">
      <i class="material-icons">folder_shared</i>
      <span>{{$t('sidebar.shares')}}</span>
    </router-link>

    <router-link v-if="user.perm.admin" :class="'action' + (this.$store.state.route.name === 'Users' ? ' active':'')" to="/users" :aria-label="$t('sidebar.users')" :title="$t('sidebar.users')">
      <i class="material-icons">people</i>
      <span>{{$t('sidebar.users')}}</span>
    </router-link>

    <router-link v-if="user.perm.admin" :class="'action' + (this.$store.state.route.name === 'Settings' ? ' active':'')" to="/settings" :aria-label="$t('sidebar.settings')" :title="$t('sidebar.settings')">
      <i class="material-icons">settings</i>
      <span>{{$t('sidebar.settings')}}</span>
    </router-link>
  </aside>
</template>

<script>
import { mapState, mapGetters } from 'vuex'
import { version, signup, disableExternal, noAuth, authMethod } from '@/utils/constants'
import FolderTree from './files/FolderTree.vue'

export default {
  name: 'sidebar',
  data: () => { return {
    'width': window.innerWidth,
    current: window.location.pathname,
    dest: null
  }},
  components: {
    FolderTree
  },
  created () {
    window.addEventListener('resize', () => {
      this.width = window.innerWidth
    })
  },
  computed: {
    ...mapState([ 'req', 'selected', 'user' ]),
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
